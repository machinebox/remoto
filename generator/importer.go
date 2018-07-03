package generator

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"io/ioutil"
	"path"
	"strings"
)

type vendorImporter struct {
	imported      map[string]*types.Package
	base          types.Importer
	skipTestFiles bool
}

func newVendorImporter(importer types.Importer) *vendorImporter {
	return &vendorImporter{
		imported:      make(map[string]*types.Package),
		base:          importer,
		skipTestFiles: true,
	}
}

func (i *vendorImporter) Import(p string) (*types.Package, error) {
	if pkg, err := i.fsPkg(p); err == nil {
		return pkg, nil
	}
	p = "./" + path.Join("vendor", p)
	pkg, err := i.fsPkg(p)
	if err != nil {
		return nil, err
	}
	i.imported[p] = pkg
	return pkg, nil
}

func (i *vendorImporter) fsPkg(pkg string) (*types.Package, error) {
	if pkg, ok := i.imported[pkg]; ok {
		return pkg, nil
	}
	dirFiles, err := ioutil.ReadDir(pkg)
	if err != nil {
		return importOrErr(i.base, pkg, err)
	}

	fset := token.NewFileSet()
	files := make([]*ast.File, 0, len(dirFiles))
	for _, fileInfo := range dirFiles {
		if fileInfo.IsDir() {
			continue
		}
		n := fileInfo.Name()
		if path.Ext(fileInfo.Name()) != ".go" {
			continue
		}
		if i.skipTestFiles && strings.Contains(fileInfo.Name(), "_test.go") {
			continue
		}
		file := path.Join(pkg, n)
		src, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, err
		}
		f, err := parser.ParseFile(fset, file, src, 0)
		if err != nil {
			return nil, err
		}
		files = append(files, f)
	}
	conf := types.Config{
		Importer: i,
	}
	p, err := conf.Check(pkg, fset, files, nil)

	if err != nil {
		return importOrErr(i.base, pkg, err)
	}
	return p, nil
}

func importOrErr(base types.Importer, pkg string, err error) (*types.Package, error) {
	p, impErr := base.Import(pkg)
	if impErr != nil {
		return nil, err
	}
	return p, nil
}

// Default returns an importer that will try to import code from gopath before using go/importer.Default and skipping test files
func Default() types.Importer {
	return &vendorImporter{
		imported:      make(map[string]*types.Package),
		base:          importer.Default(),
		skipTestFiles: true,
	}
}

// DefaultWithTestFiles same as Default but it parses test files too
func DefaultWithTestFiles() types.Importer {
	return &vendorImporter{
		imported:      make(map[string]*types.Package),
		base:          importer.Default(),
		skipTestFiles: false,
	}
}
