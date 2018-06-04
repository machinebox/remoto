package generator

import (
	"testing"

	"github.com/matryer/is"
)

type testSetter map[string]interface{}

func (s *testSetter) Set(name string, value interface{}) {
	map[string]interface{}(*s)[name] = value
}

func TestAddHelpers(t *testing.T) {
	is := is.New(t)
	s := &testSetter{}
	addHelpers(s)
	m := map[string]interface{}(*s)
	is.Equal(m["unique_structures"], uniqueStructures)
}

func TestHelperComment(t *testing.T) {
	is := is.New(t)

	is.Equal(printComment(""), ``)
	is.Equal(printComment("Something"), `// Something`+"\n")

}

func TestHelperUniqueStructures(t *testing.T) {
	is := is.New(t)
	var def Definition
	var srv1 Service
	srv1.Structures = append(srv1.Structures, Structure{
		Name: "s1",
	})
	srv1.Structures = append(srv1.Structures, Structure{
		Name: "s2",
	})
	srv1.Structures = append(srv1.Structures, Structure{
		Name: "s1",
	})
	srv1.Structures = append(srv1.Structures, Structure{
		Name: "s3",
	})
	var srv2 Service
	srv2.Structures = append(srv2.Structures, Structure{
		Name: "s1",
	})
	srv2.Structures = append(srv2.Structures, Structure{
		Name: "s2",
	})
	srv2.Structures = append(srv2.Structures, Structure{
		Name: "s1",
	})
	srv2.Structures = append(srv2.Structures, Structure{
		Name: "s3",
	})
	def.Services = append(def.Services, srv1, srv2)
	structs := uniqueStructures(def)
	is.Equal(len(structs), 3)
	is.Equal(structs[0].Name, "s1")
	is.Equal(structs[1].Name, "s2")
	is.Equal(structs[2].Name, "s3")
}
