// Code generated by templ@v0.2.184 DO NOT EDIT.

package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

// GoExpression
import (
	"github.com/joerdav/sebastion"
)

func Index(actions []sebastion.Action) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = new(bytes.Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_1 := templ.GetChildren(ctx)
		if var_1 == nil {
			var_1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		// TemplElement
		var_2 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
			// TemplElement
			err = title("Actions").Render(ctx, templBuffer)
			if err != nil {
				return err
			}
			// Whitespace (normalised)
			_, err = templBuffer.WriteString(` `)
			if err != nil {
				return err
			}
			// TemplElement
			err = table(actions).Render(ctx, templBuffer)
			if err != nil {
				return err
			}
			return err
		})
		err = Layout().Render(templ.WithChildren(ctx, var_2), templBuffer)
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

func table(actions []sebastion.Action) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = new(bytes.Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_3 := templ.GetChildren(ctx)
		if var_3 == nil {
			var_3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		// TemplElement
		var_4 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
			// Element (standard)
			_, err = templBuffer.WriteString("<table")
			if err != nil {
				return err
			}
			// Element Attributes
			_, err = templBuffer.WriteString(" class=\"table\"")
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString(">")
			if err != nil {
				return err
			}
			// Element (standard)
			_, err = templBuffer.WriteString("<thead>")
			if err != nil {
				return err
			}
			// Element (standard)
			_, err = templBuffer.WriteString("<tr>")
			if err != nil {
				return err
			}
			// Element (standard)
			_, err = templBuffer.WriteString("<th>")
			if err != nil {
				return err
			}
			// Text
			var_5 := `Title`
			_, err = templBuffer.WriteString(var_5)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</th>")
			if err != nil {
				return err
			}
			// Element (standard)
			_, err = templBuffer.WriteString("<th>")
			if err != nil {
				return err
			}
			// Text
			var_6 := `Description`
			_, err = templBuffer.WriteString(var_6)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</th>")
			if err != nil {
				return err
			}
			// Element (standard)
			_, err = templBuffer.WriteString("<th>")
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</th>")
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</tr>")
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</thead>")
			if err != nil {
				return err
			}
			// Element (standard)
			_, err = templBuffer.WriteString("<tbody>")
			if err != nil {
				return err
			}
			// For
			for _, a := range actions {
				// Element (standard)
				_, err = templBuffer.WriteString("<tr>")
				if err != nil {
					return err
				}
				// Element (standard)
				_, err = templBuffer.WriteString("<td>")
				if err != nil {
					return err
				}
				// StringExpression
				_, err = templBuffer.WriteString(templ.EscapeString(a.Details().Name))
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString("</td>")
				if err != nil {
					return err
				}
				// Element (standard)
				_, err = templBuffer.WriteString("<td>")
				if err != nil {
					return err
				}
				// StringExpression
				_, err = templBuffer.WriteString(templ.EscapeString(a.Details().Description))
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString("</td>")
				if err != nil {
					return err
				}
				// Element (standard)
				_, err = templBuffer.WriteString("<td>")
				if err != nil {
					return err
				}
				// Element (standard)
				_, err = templBuffer.WriteString("<a")
				if err != nil {
					return err
				}
				// Element Attributes
				_, err = templBuffer.WriteString(" class=\"button is-black is-radiusless\"")
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString(" href=")
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString("\"")
				if err != nil {
					return err
				}
				var var_7 templ.SafeURL = templ.SafeURL(actionUrl(a))
				_, err = templBuffer.WriteString(templ.EscapeString(string(var_7)))
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString("\"")
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString(">")
				if err != nil {
					return err
				}
				// Text
				var_8 := `Run`
				_, err = templBuffer.WriteString(var_8)
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString("</a>")
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString("</td>")
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString("</tr>")
				if err != nil {
					return err
				}
			}
			_, err = templBuffer.WriteString("</tbody>")
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</table>")
			if err != nil {
				return err
			}
			return err
		})
		err = card().Render(templ.WithChildren(ctx, var_4), templBuffer)
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

