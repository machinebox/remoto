package generator

import (
	"io"

	"github.com/gobuffalo/plush"
	"github.com/machinebox/remoto/generator/definition"
	"github.com/pkg/errors"
)

// Render renders the tpl template with the Definition into w.
func Render(w io.Writer, templateName, tpl string, def definition.Definition) error {
	ctx := plush.NewContext()
	ctx.Set("def", def)
	AddTemplateHelpers(ctx)
	out, err := plush.Render(tpl, ctx)
	if err != nil {
		return errors.Wrapf(err, "plush.Render (%s)", templateName)
	}
	if _, err := io.WriteString(w, out); err != nil {
		return err
	}
	return nil
}
