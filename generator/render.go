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
	out, err := plush.Render(tpl, ctx)
	if err != nil {
		return errors.Wrap(err, "plush.Render")
	}
	if _, err := io.WriteString(w, out); err != nil {
		return err
	}
	return nil
}
