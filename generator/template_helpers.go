package generator

import (
	"sort"
	"strings"

	"github.com/machinebox/remoto/generator/definition"
	"github.com/markbates/inflect"
)

var defaultRuleset = inflect.NewDefaultRuleset()

// Setter may have data set on it, usually a plush context.
type Setter interface {
	Set(name string, value interface{})
}

// AddTemplateHelpers adds all the built-in template helpers.
func AddTemplateHelpers(s Setter) {
	s.Set("unique_structures", uniqueStructures)
	s.Set("print_comment", printComment)
	s.Set("go_type_string", goTypeString)
	s.Set("underscore", underscore)
	s.Set("camelize_down", camelizeDownFirst)

	// experimental (undocumented)
	s.Set("replace", replace)
}

// underscore converts a type name or other string into an underscored
// version. For example, "ModelID" becomes "model_id".
func underscore(s string) string {
	return defaultRuleset.Underscore(s)
}

// camelizeDownFirst converts a name or other string into a camel case
// version with the first letter lowercase. "ModelID" becomes "modelID".
func camelizeDownFirst(s string) string {
	if s == "ID" {
		return "id"
		// note: not sure why I need this, there's a lot that deals with
		// accronyms in the dependency packages but they don't seem to behave
		// as expected in this case.
	}
	return defaultRuleset.CamelizeDownFirst(s)
}

// uniqueStructures gets all unique Structure types from all services.
// Structures with the same name are considered the same.
// Use unique_structures(def) in templates.
func uniqueStructures(def definition.Definition) []definition.Structure {
	structures := make(map[string]definition.Structure)
	for _, service := range def.Services {
		for _, structure := range service.Structures {
			if structure.IsImported {
				continue
			}
			structures[structure.Name] = structure
		}
	}
	s := make([]definition.Structure, 0, len(structures))
	for _, structure := range structures {
		s = append(s, structure)
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i].Name < s[j].Name
	})
	return s
}

// printComment prints a comment with // prefix, unless the comment
// is empty.
// Use print_comment(s) in templates.
func printComment(comment string) string {
	if comment == "" {
		return ""
	}
	var out string
	for _, line := range strings.Split(comment, "\n") {
		out += "// " + line + "\n"
	}
	return out
}

// goTypeString gets the Type as a Go string.
// Use go_type_string(type) in templates.
func goTypeString(typ definition.Type) string {
	if typ.IsMultiple {
		return "[]" + typ.Name
	}
	return typ.Name
}

// replace is a string replacement function.
func replace(s, old, new string) string {
	return strings.Replace(s, old, new, -1)
}
