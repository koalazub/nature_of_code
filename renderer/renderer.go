package renderer

import (
	"html/template"
	"net/http"

	"golang.org/x/exp/slog"
)

// we'll be drawing lines here
func Renderer() {

}

type Resolution struct {
	Height float64
	Width  float64
}

type data struct {
	Title   string
	Heading string
	Body    string
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {

	d := data{
		Title:   "Rendar",
		Heading: "Here is Rendar",
		Body:    "Where we try to learn maths and rendering in one hot hit",
	}

	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		slog.Error("There was an error reading the file", "err", err)
		return
	}

	err = tmpl.Execute(w, d)
	if err != nil {
		slog.Error("Couldn't render the page", err)
		return
	}
}

func HandleVectorCanvas(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("./templates/canvas.html")
	if err != nil {
		slog.Error("There was an error reading the file", "err", err)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		slog.Error("Couldn't render the page", err)
		return
	}
}

func HandleTemplate(w http.ResponseWriter, r *http.Request) {
	const t = `
<h1>
	<title> this title{{.Title}}</title>
</h1>

<body>{{.Body}}</body>`

	tmpl, err := template.New("example").Parse(t)
	if err != nil {
		slog.Error("Couldn't render template", err)
		return
	}

	d := data{
		Title:   "zub",
		Heading: "OOOOOOH",
		Body:    "we movin",
	}

	err = tmpl.Execute(w, d)
	if err != nil {
		slog.Error("couldn't apply template", err)
	}
}
