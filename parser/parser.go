package parser

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"strings"

	"github.com/pkg/errors"
)

// Parse parses a package of .rpc.go files.
func Parse(dir string) (Definition, error) {
	var def Definition
	def.comments = make(map[string]string)
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, func(info os.FileInfo) bool {
		return strings.HasSuffix(info.Name(), ".rpc.go")
	}, parser.ParseComments)
	if err != nil {
		return def, errors.Wrap(err, "parser.ParseDir")
	}
	pkgNames := make([]string, 0, len(pkgs))
	for pkg := range pkgs {
		pkgNames = append(pkgNames, pkg)
	}
	if len(pkgNames) == 0 {
		return def, errors.New("no packages found")
	}
	if len(pkgNames) > 1 {
		return def, errors.New("multiple packages found: " + strings.Join(pkgNames, ", "))
	}
	firstPkg := pkgs[pkgNames[0]]
	files := make([]*ast.File, 0, len(firstPkg.Files))
	for _, file := range firstPkg.Files {
		files = append(files, file)
	}
	info := &types.Info{}
	conf := types.Config{Importer: importer.Default()}
	pkg, err := conf.Check(dir, fset, files, info)
	if err != nil {
		return def, errors.Wrap(err, "conf.Check")
	}
	for _, f := range files {
		for _, comment := range f.Comments {
			pos := comment.Pos()
			name := strings.TrimSpace(comment.Text())
			name = strings.Split(name, " ")[0]
			inner := pkg.Scope().Innermost(pos)
			if _, obj := inner.LookupParent(name, pos); obj != nil {
				def.comments[obj.Name()] = comment.Text()
			}
		}
	}
	for _, name := range pkg.Scope().Names() {
		obj := pkg.Scope().Lookup(name)
		switch v := obj.Type().Underlying().(type) {
		case *types.Interface:
			service, err := parseService(fset, pkg.Scope(), &def, obj, v)
			if err != nil {
				return def, err
			}
			def.Services = append(def.Services, service)
		}
	}
	return def, nil
}

// Definition is the definition of one or more services.
type Definition struct {
	Services []Service
	comments map[string]string
}

// Service describes a logically grouped set of endpoints.
type Service struct {
	Name       string
	Comment    string
	Methods    []Method
	Structures []Structure
}

// EnsureStructure adds the Structure to the service if it isn't
// already there.
func (s *Service) EnsureStructure(structure Structure) {
	for i := range s.Structures {
		if s.Structures[i].Name == structure.Name {
			return
		}
	}
	s.Structures = append(s.Structures, structure)
}

func (s Service) String() string {
	var str string
	if s.Comment != "" {
		str += "// " + s.Comment
	}
	str += "type " + s.Name + " struct {\n"
	for i := range s.Methods {
		str += "\t" + s.Methods[i].String() + "\n"
	}
	str += "}\n"
	for i := range s.Structures {
		str += s.Structures[i].String()
	}
	return str
}

// Method is a single method.
type Method struct {
	Name     string
	Comment  string
	Request  string
	Response string
}

func (m Method) String() string {
	var str string
	if m.Comment != "" {
		str += "// " + m.Comment
	}
	str += m.Name + "(context.Context, *" + m.Request + ") (*" + m.Response + ", error)\n"
	return str
}

// Structure describes a data structure.
type Structure struct {
	Name    string
	Comment string
	Fields  []Field
}

func (s Structure) String() string {
	var str string
	if s.Comment != "" {
		str += "// " + s.Comment
	}
	str += "type " + s.Name + " struct {\n"
	for i := range s.Fields {
		str += "\t" + s.Fields[i].String() + "\n"
	}
	str += "}\n"
	return str
}

// Field describes a structure field.
type Field struct {
	Name string
	Type string
}

func (f Field) String() string {
	return fmt.Sprintf("%s %s", f.Name, f.Type)
}

func parseService(fset *token.FileSet, scope *types.Scope, def *Definition, obj types.Object, v *types.Interface) (Service, error) {
	srv := Service{
		Name:    obj.Name(),
		Comment: def.comments[obj.Name()],
	}
	for i := 0; i < v.NumMethods(); i++ {
		m := v.Method(i)
		method, err := parseMethod(fset, scope, def, &srv, m)
		if err != nil {
			return srv, err
		}
		srv.Methods = append(srv.Methods, method)
	}
	return srv, nil
}

func parseMethod(fset *token.FileSet, scope *types.Scope, def *Definition, srv *Service, m *types.Func) (Method, error) {
	method := Method{
		Name:    m.Name(),
		Comment: def.comments[m.Name()],
	}
	sig := m.Type().(*types.Signature)
	params := sig.Params()
	// process input arguments
	if params.Len() != 2 || params.At(0).Type().String() != "context.Context" {
		return method, newErr(fset, m.Pos(), "service methods must have signature (context.Context, *Request) (*Response, error)")
	}
	requestParam := params.At(1)
	requestStructure, err := parseStructureFromParam(fset, scope, def, srv, "request", requestParam)
	if err != nil {
		return method, err
	}
	method.Request = requestStructure.Name
	srv.EnsureStructure(requestStructure)
	// process return arguments
	returns := sig.Results()
	if returns.Len() != 2 || returns.At(1).Type().String() != "error" {
		return method, newErr(fset, m.Pos(), "service methods must have signature (context.Context, *Request) (*Response, error)")
	}
	responseParam := returns.At(0)
	responseStructure, err := parseStructureFromParam(fset, scope, def, srv, "response", responseParam)
	if err != nil {
		return method, err
	}
	method.Response = responseStructure.Name
	srv.EnsureStructure(responseStructure)
	return method, nil
}

func parseStructureFromParam(fset *token.FileSet, scope *types.Scope, def *Definition, srv *Service, structureKind string, v *types.Var) (Structure, error) {
	var structure Structure
	p, ok := v.Type().(*types.Pointer)
	if !ok {
		return structure, newErr(fset, v.Pos(), structureKind+" object must be a pointer to a struct")
	}
	st, ok := p.Elem().Underlying().(*types.Struct)
	if !ok {
		return structure, newErr(fset, v.Pos(), structureKind+" object must be a pointer to a struct")
	}
	structure.Name = types.TypeString(v.Type(), nakedType)[1:]
	structure.Comment = def.comments[structure.Name]
	for i := 0; i < st.NumFields(); i++ {
		field, err := parseField(fset, scope, def, srv, st.Field(i))
		if err != nil {
			return structure, err
		}
		structure.Fields = append(structure.Fields, field)
	}
	return structure, nil
}

func parseStructure(fset *token.FileSet, scope *types.Scope, def *Definition, srv *Service, obj types.Object) (Structure, error) {
	structure := Structure{
		Name: obj.Name(),
	}
	structure.Comment = def.comments[structure.Name]
	st, ok := obj.Type().Underlying().(*types.Struct)
	if !ok {
		return structure, newErr(fset, obj.Pos(), obj.Type().String()+" object must be a pointer to a struct")
	}
	for i := 0; i < st.NumFields(); i++ {
		field, err := parseField(fset, scope, def, srv, st.Field(i))
		if err != nil {
			return structure, err
		}
		structure.Fields = append(structure.Fields, field)
	}
	return structure, nil
}

func parseField(fset *token.FileSet, scope *types.Scope, def *Definition, srv *Service, v *types.Var) (Field, error) {
	var field Field
	if !v.IsField() {
		return field, newErr(fset, v.Pos(), v.Name()+" not a field")
	}
	typeName, isStruct, err := resolveTypeName(v.Type())
	if err != nil {
		return field, newErr(fset, v.Pos(), err.Error())
	}
	field.Name = v.Name()
	field.Type = typeName
	if isStruct {
		obj := scope.Lookup(typeName)
		structure, err := parseStructure(fset, scope, def, srv, obj)
		if err != nil {
			return field, err
		}
		srv.EnsureStructure(structure)
	}
	return field, nil
}

func newErr(fset *token.FileSet, pos token.Pos, err string) error {
	position := fset.Position(pos)
	return errors.New(position.String() + ": " + err)
}

// nakedType doesn't prefix the type name, regardless.
func nakedType(other *types.Package) string {
	return ""
}

func resolveTypeName(typ types.Type) (string, bool, error) {
	typeName := types.TypeString(typ, nakedType)
	if _, ok := typ.Underlying().(*types.Struct); ok {
		return typeName, true, nil
	}
	switch typeName {
	case "string", "float64", "int64", "bool":
		return typeName, false, nil
	}
	return "", false, errors.New("type " + typeName + " not supported")
}
