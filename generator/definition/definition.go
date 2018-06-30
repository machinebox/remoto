package definition

import (
	"errors"
	"fmt"
	"strings"
)

// Definition is the definition of one or more services.
// In templates, it is usually accessed via the `def` variable.
//  Package name is <%= def.PackageName %>
type Definition struct {
	Services       []Service `json:"services"`
	PackageName    string    `json:"packageName"`
	PackageComment string    `json:"packageComment"`
}

// Source gets the Remoto source for this definition.
func (d Definition) Source() string {
	s := printComments(d.PackageComment)
	s += "package " + d.PackageName + "\n\n"
	for i := range d.Services {
		s += d.Services[i].String()
	}
	return s
}

func (d Definition) String() string {
	return d.Source()
}

// Valid gets whether this Definition is valid or not.
func (d Definition) Valid() error {
	if len(d.Services) == 0 {
		return errors.New("must provide at least one service")
	}
	for _, service := range d.Services {
		if len(service.Methods) == 0 {
			return errors.New("service " + service.Name + " must have at least one method")
		}
	}
	return nil
}

// Structure gets a Structure by name.
func (d Definition) Structure(name string) *Structure {
	for _, service := range d.Services {
		for _, structure := range service.Structures {
			if structure.Name == name {
				return &structure
			}
		}
	}
	return nil
}

// Service describes a logically grouped set of endpoints.
type Service struct {
	Name       string      `json:"name"`
	Comment    string      `json:"comment"`
	Methods    []Method    `json:"methods"`
	Structures []Structure `json:"structures"`
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
	str := printComments(s.Comment)
	str += "type " + s.Name + " interface {\n"
	for i := range s.Methods {
		str += indent(1, s.Methods[i].String())
	}
	str += "}\n\n"
	for i := range s.Structures {
		str += s.Structures[i].String()
	}
	return str
}

// Method is a single method.
type Method struct {
	Name              string    `json:"name"`
	Comment           string    `json:"comment"`
	RequestStructure  Structure `json:"requestStructure"`
	ResponseStructure Structure `json:"responseStructure"`
}

func (m Method) String() string {
	str := printComments(m.Comment)
	str += m.Name + "(" + m.RequestStructure.Name + ") " + m.ResponseStructure.Name
	return str
}

// Structure describes a data structure.
type Structure struct {
	Name       string  `json:"name"`
	Comment    string  `json:"comment"`
	Fields     []Field `json:"fields"`
	IsImported bool    `json:"isImported"`

	IsRequestObject  bool `json:"isRequestObject"`
	IsResponseObject bool `json:"isResponseObject"`
}

func (s Structure) String() string {
	str := printComments(s.Comment)
	str += "type " + s.Name + " struct {\n"
	for i := range s.Fields {
		str += indent(1, s.Fields[i].String())
	}
	str += "}\n\n"
	return str
}

// HasFields gets whether the Structure has any fields or not.
func (s Structure) HasFields() bool {
	return len(s.Fields) > 0
} // TODO: test

// HasField gets whether the Structure has a specific field or not.
func (s Structure) HasField(field string) bool {
	for _, f := range s.Fields {
		if f.Name == field {
			return true
		}
	}
	return false
} // TODO: test

// FieldsOfType gets all Field objects that have a specific type.
func (s Structure) FieldsOfType(typename string) []Field {
	var fields []Field
	for _, field := range s.Fields {
		if field.Type.Name == typename {
			fields = append(fields, field)
		}
	}
	return fields
}

// Field describes a structure field.
type Field struct {
	Name    string `json:"name"`
	Comment string `json:"comment"`
	Type    Type   `json:"type"`
}

func (f Field) String() string {
	return fmt.Sprintf("%s%s %s", printComments(f.Comment), f.Name, f.Type.code())
}

// IsExported gets whether the field is exported or not.
func (f Field) IsExported() bool {
	return strings.ToUpper(f.Name[0:1]) == f.Name[0:1]
}

// Type describes the type of a Field.
type Type struct {
	Name       string `json:"name"`
	IsMultiple bool   `json:"isMultiple"`
	IsStruct   bool   `json:"isStruct"`
	IsImported bool   `json:"isImported"`
}

func (t Type) code() string {
	str := t.Name
	if t.IsMultiple {
		str = "[]" + str
	}
	return str
}

func printComments(comment string) string {
	if comment == "" {
		return ""
	}
	var out string
	for _, line := range strings.Split(comment, "\n") {
		out += "// " + line + "\n"
	}
	return out
}

func indent(n int, s string) string {
	var out string
	prefix := strings.Repeat("\t", n)
	for _, line := range strings.Split(s, "\n") {
		out += prefix + line + "\n"
	}
	return out
}
