package definition

import (
	"testing"

	"github.com/matryer/is"
)

func TestStructure(t *testing.T) {
	is := is.New(t)

	struct1 := Structure{
		Name: "StructureOne",
	}
	struct2 := Structure{
		Name: "StructureTwo",
	}
	def := Definition{
		Services: []Service{
			{
				Name:       "One",
				Structures: []Structure{struct1},
			},
			{
				Name:       "Two",
				Structures: []Structure{struct2},
			},
		},
	}
	is.Equal(*def.Structure("StructureOne"), struct1)
	is.Equal(*def.Structure("StructureTwo"), struct2)
	is.True(def.Structure("Nope") == nil)
}
