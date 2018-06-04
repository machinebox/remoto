package generator

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

// Definition is the definition of one or more services.
type Definition struct {
	Services       []Service
	PackageName    string
	PackageComment string
	FieldTypes     map[string]Type

	comments map[string]string
}

func (d Definition) String() string {
	s := "package " + d.PackageName + "\n\n"
	for i := range d.Services {
		s += d.Services[i].String()
	}
	return s
}

// Service describes a logically grouped set of endpoints.
type Service struct {
	Name       string
	Comment    string
	Methods    []Method
	Structures []Structure
}

// ensureStructure adds the Structure to the service if it isn't
// already there.
func (s *Service) ensureStructure(structure Structure) {
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
		str += "// " + s.Comment + "\n"
	}
	str += "type " + s.Name + " interface {\n"
	for i := range s.Methods {
		str += "\t" + s.Methods[i].String()
	}
	str += "}\n\n"
	for i := range s.Structures {
		str += s.Structures[i].String()
	}
	return str
}

// Method is a single method.
type Method struct {
	Name         string
	Comment      string
	RequestType  Structure
	ResponseType Structure
}

func (m Method) String() string {
	var str string
	if m.Comment != "" {
		str += "// " + m.Comment + "/n"
	}
	str += m.Name + "(context.Context, *" + m.RequestType.Name + ") (*" + m.ResponseType.Name + ", error)\n"
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
		str += "// " + s.Comment + "\n"
	}
	str += "type " + s.Name + " struct {\n"
	for i := range s.Fields {
		str += "\t" + s.Fields[i].String() + "\n"
	}
	str += "}\n\n"
	return str
}

// Field describes a structure field.
type Field struct {
	Name    string
	Comment string
	Type    Type
}

func (f Field) String() string {
	return fmt.Sprintf("%s %s", f.Name, f.Type.code())
}

// Type describes the type of a Field.
type Type struct {
	Name       string
	IsMultiple bool
	IsStruct   bool
}

func (t Type) code() string {
	str := t.Name
	if t.IsMultiple {
		str = "[]" + str
	}
	return str
}

// Parse parses a package of .remoto.go files.
func Parse(dir string) (Definition, error) {
	var def Definition
	def.comments = make(map[string]string)
	def.FieldTypes = make(map[string]Type)
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, func(info os.FileInfo) bool {
		return strings.HasSuffix(info.Name(), ".remoto.go")
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
	def.PackageName = pkgNames[0]
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
			// TODO(matryer): use a technique that can get comments for methods too.
			pos := comment.Pos()
			trimmedComment := strings.TrimSpace(comment.Text())
			name := strings.Split(trimmedComment, " ")[0]
			inner := pkg.Scope().Innermost(pos)
			if _, obj := inner.LookupParent(name, pos); obj != nil {
				def.comments[obj.Name()] = trimmedComment
			}
		}
	}
	for _, name := range pkg.Scope().Names() {
		obj := pkg.Scope().Lookup(name)
		switch v := obj.Type().Underlying().(type) {
		case *types.Interface:
			service, err := parseService(fset, pkg.Scope(), &def, obj, v)
			if err != nil {
				for sub, tip := range tips {
					if strings.Contains(err.Error(), sub) {
						err = errors.New(err.Error() + ": " + tip)
						break
					}
				}
				return def, err
			}
			def.Services = append(def.Services, service)
		}
	}
	return def, nil
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
	if !m.Exported() {
		return method, newErr(fset, m.Pos(), "method "+m.Name()+": must be exported")
	}
	sig := m.Type().(*types.Signature)
	if sig.Variadic() {
		return method, newErr(fset, m.Pos(), "service methods must have signature (context.Context, *Request) (*Response, error)")
	}
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
	if !strings.HasSuffix(requestStructure.Name, "Request") {
		return method, newErr(fset, m.Pos(), "request object should end with \"Request\"")
	}
	method.RequestType = requestStructure
	srv.ensureStructure(requestStructure)
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
	if !strings.HasSuffix(responseStructure.Name, "Response") {
		return method, newErr(fset, m.Pos(), "response object should end with \"Response\"")
	}
	addDefaultResponseFields(&responseStructure)
	method.ResponseType = responseStructure
	srv.ensureStructure(responseStructure)
	return method, nil
}

// addDefaultResponseFields adds the built-in remoto fields to the
// response structure.
func addDefaultResponseFields(structure *Structure) {
	structure.Fields = append(structure.Fields, Field{
		Comment: "Error is an error message if one occurred.",
		Name:    "Error",
		Type: Type{
			Name: "string",
		},
	})
}

func parseStructureFromParam(fset *token.FileSet, scope *types.Scope, def *Definition, srv *Service, structureKind string, v *types.Var) (Structure, error) {
	resolver := func(other *types.Package) string {
		if other.Name() != def.PackageName {
			return other.Name()
		}
		return ""
	}
	var structure Structure
	p, ok := v.Type().(*types.Pointer)
	if !ok {
		return structure, newErr(fset, v.Pos(), structureKind+" object must be a pointer to a struct")
	}
	st, ok := p.Elem().Underlying().(*types.Struct)
	if !ok {
		return structure, newErr(fset, v.Pos(), structureKind+" object must be a pointer to a struct")
	}
	structure.Name = types.TypeString(v.Type(), resolver)[1:]
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
		return structure, newErr(fset, obj.Pos(), obj.Type().String()+" field must be a pointer to a struct")
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
	if !v.Exported() {
		return field, newErr(fset, v.Pos(), "field "+v.Name()+": must be exported")
	}
	typ, err := resolveTypeName(def, v.Type())
	if err != nil {
		return field, newErr(fset, v.Pos(), err.Error())
	}
	def.FieldTypes[typ.Name] = typ
	field.Name = v.Name()
	field.Type = typ
	if typ.IsStruct {
		obj := scope.Lookup(typ.Name)
		structure, err := parseStructure(fset, scope, def, srv, obj)
		if err != nil {
			return field, err
		}
		srv.ensureStructure(structure)
	}
	return field, nil
}

func newErr(fset *token.FileSet, pos token.Pos, err string) error {
	position := fset.Position(pos)
	return errors.New(position.String() + ": " + err)
}

func resolveTypeName(def *Definition, typ types.Type) (Type, error) {
	resolver := func(other *types.Package) string {
		if other.Name() != def.PackageName {
			return other.Name()
		}
		return ""
	}
	var ty Type
	slice, ok := typ.(*types.Slice)
	if ok {
		ty.IsMultiple = true
		typ = slice.Elem()
	}
	ty.Name = types.TypeString(typ, resolver)
	if _, ok := typ.Underlying().(*types.Struct); ok {
		ty.IsStruct = true
		return ty, nil
	}
	switch ty.Name {
	case "string", "float64", "int32", "int64", "bool",
		"remototypes.File":
		return ty, nil
	}
	return ty, errors.New("type " + ty.Name + " not supported")
}

var tips = map[string]string{
	" int ": "use explicitly sized types int32 or int64",
}
