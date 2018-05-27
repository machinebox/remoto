package generator

import (
	"io"

	"github.com/gobuffalo/plush"
	"github.com/pkg/errors"
)

// Render renders the tpl template with the Definition into w.
func Render(w io.Writer, tpl string, def Definition) error {
	ctx := plush.NewContext()
	ctx.Set("def", def)
	ctx.Set("unique_structures", uniqueStructures)
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
