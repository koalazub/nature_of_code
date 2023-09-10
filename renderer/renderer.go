package renderer

import (
	"html/template"
	"net/http"

	"golang.org/x/exp/slog"
)

// we'll be drawing lines here
func Renderer() {

}

/*
	tmpl, err := template.New("example").Parse(`
			<html>
			<head><title>{{.Title}}</head>
			<body>
			<h1>{{.Message}}</h1>
			</body>
		</html>
		`)
*/

func HandleTemplate(w http.ResponseWriter, r *http.Request) {
	const t = `
		<h1><title>{{.Title}}</title></h1>
		<body>{{.Message}}</body>`

	tmpl, err := template.New("example").Parse(t)
	if err != nil {
		slog.Error("Couldn't render template", err)
	}

	data := struct {
		Title   string
		Message string
	}{
		Title:   "zub",
		Message: "oooooh we movin",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		slog.Error("couldn't apply template", err)
	}
}
