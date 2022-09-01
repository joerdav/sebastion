// Code generated by templ@v0.2.184 DO NOT EDIT.

package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func LogComponent(text string) templ.Component {
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
		// Element (standard)
		_, err = templBuffer.WriteString("<textarea")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" class=\"textarea is-black\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" readonly=\"true\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" id=\"action-logs\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// StringExpression
		_, err = templBuffer.WriteString(templ.EscapeString(text))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</textarea>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

func Log(outputId string, text string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = new(bytes.Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_2 := templ.GetChildren(ctx)
		if var_2 == nil {
			var_2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		// TemplElement
		var_3 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
			// TemplElement
			err = LogInit(outputId, text).Render(ctx, templBuffer)
			if err != nil {
				return err
			}
			return err
		})
		err = Layout().Render(templ.WithChildren(ctx, var_3), templBuffer)
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

func LogInit(outputId string, text string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = new(bytes.Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_4 := templ.GetChildren(ctx)
		if var_4 == nil {
			var_4 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		// Element (standard)
		_, err = templBuffer.WriteString("<turbo-frame")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" id=\"log-frame\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// TemplElement
		err = LogComponent(text).Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		// Element (void)
		_, err = templBuffer.WriteString("<input")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" type=\"hidden\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" id=\"outputid\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" value=")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(outputId))
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
// RawElement
		_, err = templBuffer.WriteString("<script>")
		if err != nil {
			return err
		}
// Text
var_5 := `
			var id = document.querySelector('#outputid').value
			console.log(id)
			Turbo.connectStreamSource(new WebSocket(` + "`" + `ws://${window.location.host}/output/${id}/ws` + "`" + `));
		`
_, err = templBuffer.WriteString(var_5)
if err != nil {
	return err
}
		_, err = templBuffer.WriteString("</script>")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</turbo-frame>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

func LogStream(text string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = new(bytes.Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_6 := templ.GetChildren(ctx)
		if var_6 == nil {
			var_6 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		// Element (standard)
		_, err = templBuffer.WriteString("<turbo-stream")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" action=\"replace\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" targets=\"#action-logs\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<template>")
		if err != nil {
			return err
		}
		// TemplElement
		err = LogComponent(text).Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</template>")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</turbo-stream>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}
