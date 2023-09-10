package renderer

import (
	"html/template"
	"os"

	"golang.org/x/exp/slog"
)

func Renderer() {

}

func HtmlTemplate() {
	tmpl, err := template.New("example").Parse(`
			<html>
			<head><title>{{.Title}}</head>
			<body>
			<h1>{{.Message}}</h1>
			</body>
		</html>
		`)
	if err != nil {
		slog.Error("Couldn't render template", err)
	}

	data := struct {
		Title   string
		Message string
	}{
		Title:   "Go templates examplwe",
		Message: "oooooh we movin",
	}

	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		slog.Error("couldn't apply template", err)
	}

}
