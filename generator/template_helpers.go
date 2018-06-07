package generator

import (
	"github.com/markbates/inflect"
)

type setter interface {
	Set(name string, value interface{})
}

// addHelpers adds all the built-in template helpers.
func addHelpers(s setter) {
	s.Set("unique_structures", uniqueStructures)
	s.Set("print_comment", printComment)
	s.Set("go_type_string", goTypeString)
	s.Set("underscore", underscore)
}

// underscore converts a type name or other string into an underscored
// version. For example, "ModelID" becomes "model_id".
func underscore(s string) string {
	s = inflect.Underscore(s)
	return s
}

// uniqueStructures gets all unique Structure types from all services.
// Structures with the same name are considered the same.
// Use unique_structures(def) in templates.
func uniqueStructures(def Definition) []Structure {
	structures := make(map[string]Structure)
	for _, service := range def.Services {
		for _, structure := range service.Structures {
			if structure.IsImported {
				continue
			}
			structures[structure.Name] = structure
		}
	}
	s := make([]Structure, 0, len(structures))
	for _, structure := range structures {
		s = append(s, structure)
	}
	return s
}

// printComment prints a comment with // prefix, unless the comment
// is empty.
// Use print_comment(s) in templates.
func printComment(comment string) string {
	if comment == "" {
		return ""
	}
	return "// " + comment + "\n"
}

// goTypeString gets the Type as a Go string.
// Use go_type_string(type) in templates.
func goTypeString(typ Type) string {
	if typ.IsMultiple {
		return "[]" + typ.Name
	}
	return typ.Name
}
