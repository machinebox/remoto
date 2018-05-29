package generator

import (
	"fmt"
	"io"

	"github.com/gobuffalo/plush"
	"github.com/pkg/errors"
)

// Render renders the tpl template with the Definition into w.
func Render(w io.Writer, tpl string, def Definition) error {
	ctx := plush.NewContext()
	ctx.Set("def", def)
	ctx.Set("unique_structures", uniqueStructures)
	ctx.Set("has_field_type", hasFieldType)
	out, err := plush.Render(tpl, ctx)
	if err != nil {
		return errors.Wrap(err, "plush.Render")
	}
	if _, err := io.WriteString(w, out); err != nil {
		return err
	}
	return nil
}

// uniqueStructures gets all unique Structure types from all services.
// Structures with the same name are considered the same.
func uniqueStructures(def Definition) []Structure {
	structures := make(map[string]Structure)
	for _, service := range def.Services {
		for _, structure := range service.Structures {
			structures[structure.Name] = structure
		}
	}
	s := make([]Structure, 0, len(structures))
	for _, structure := range structures {
		s = append(s, structure)
	}
	return s
}

// hasFieldType checks the Structure to see if it has any specific
// types or not.
func hasFieldType(typ interface{}, typename string) bool {
	switch v := typ.(type) {
	case Structure:
		for _, field := range v.Fields {
			if field.Type.Name == "remototypes.File" {
				return true
			}
		}
	case Definition:
		for typ := range v.FieldTypes {
			if typ == typename {
				return true
			}
		}
	default:
		panic(fmt.Sprintf("has_field_type does not support %T", typ))
	}
	return false
}
