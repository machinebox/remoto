package generator

import (
	"testing"

	"github.com/machinebox/remoto/generator/definition"
	"github.com/matryer/is"
)

type testSetter map[string]interface{}

func (s *testSetter) Set(name string, value interface{}) {
	map[string]interface{}(*s)[name] = value
}

func TestAddHelpers(t *testing.T) {
	is := is.New(t)
	s := &testSetter{}
	AddTemplateHelpers(s)
	m := map[string]interface{}(*s)
	is.Equal(m["unique_structures"], uniqueStructures)
	is.Equal(m["print_comment"], printComment)
	is.Equal(m["go_type_string"], goTypeString)
}

func TestHelperComment(t *testing.T) {
	is := is.New(t)

	is.Equal(printComment(""), ``)
	is.Equal(printComment("Something"), `// Something`+"\n")

	is.Equal(printComment("one\ntwo\nthree"), `// one
// two
// three
`)

}

func TestHelperUniqueStructures(t *testing.T) {
	is := is.New(t)
	var def definition.Definition
	var srv1 definition.Service
	srv1.Structures = append(srv1.Structures, definition.Structure{
		Name: "s1",
	})
	srv1.Structures = append(srv1.Structures, definition.Structure{
		Name:       "s2",
		IsImported: true,
	})
	srv1.Structures = append(srv1.Structures, definition.Structure{
		Name: "s1",
	})
	srv1.Structures = append(srv1.Structures, definition.Structure{
		Name: "s3",
	})
	var srv2 definition.Service
	srv2.Structures = append(srv2.Structures, definition.Structure{
		Name: "s1",
	})
	srv2.Structures = append(srv2.Structures, definition.Structure{
		Name:       "s2",
		IsImported: true,
	})
	srv2.Structures = append(srv2.Structures, definition.Structure{
		Name:       "s1",
		IsImported: true,
	})
	srv2.Structures = append(srv2.Structures, definition.Structure{
		Name: "s3",
	})
	def.Services = append(def.Services, srv1, srv2)
	structs := uniqueStructures(def)
	is.Equal(len(structs), 2)
}

func TestGoTypeString(t *testing.T) {
	is := is.New(t)
	typ := definition.Type{
		Name:       "string",
		IsMultiple: false,
		IsStruct:   false,
	}
	is.Equal(goTypeString(typ), "string")
	typ = definition.Type{
		Name:       "string",
		IsMultiple: true,
		IsStruct:   false,
	}
	is.Equal(goTypeString(typ), "[]string")
}

func TestUnderscore(t *testing.T) {
	is := is.New(t)
	is.Equal(underscore("hello there"), `hello_there`)
	is.Equal(underscore("Hello There"), `hello_there`)
	is.Equal(underscore("SomethingElse"), `something_else`)
	is.Equal(underscore("ModelID"), `model_id`)
	is.Equal(underscore("ID"), `id`)
}

func TestCamelizeDownFirst(t *testing.T) {
	is := is.New(t)
	is.Equal(camelizeDownFirst("hello there"), `helloThere`)
	is.Equal(camelizeDownFirst("Hello There"), `helloThere`)
	is.Equal(camelizeDownFirst("SomethingElse"), `somethingElse`)
	is.Equal(camelizeDownFirst("ModelID"), `modelID`)
	is.Equal(camelizeDownFirst("ID"), `id`)
}

func TestReplace(t *testing.T) {
	is := is.New(t)
	is.Equal(replace("one two three", "two", "2"), "one 2 three")
}
