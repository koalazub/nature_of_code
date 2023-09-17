package renderer

import (
	"html/template"
	"net/http"
	"strconv"

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

	tmpl, err := template.ParseFiles("./index.html")
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

func HandleTemplate(w http.ResponseWriter, r *http.Request) {
	const t = `
		<h1><title> this title{{.Title}}</title></h1>
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

func init() {

	page = template.New("main")

	page = template.Must(page.Parse(`<!DOCTYPE html> 
	<html>
	<head>
		<script src="https://unpkg.com/htmx.org@1.8.0"></script>
		<link rel="stylesheet" href="https://the.missing.style"/>
		<title>Template Fragment Example</title>
	</head>
	<body>
		<h1>Template Fragment Example</h1>
		
		<p>This page demonstrates how to create and serve 
		<a href="https://htmx.org/essays/template-fragments/">template fragments</a> 
		using the <a href="https://pkg.go.dev/text/template">built-in template package</a> in Go.</p>
		
		<p>This is accomplished by using the "block" action in the template, which lets you
		define and execute a sub-template in a single step.</p>
		<!-- Here's the fragment.  We can target it by executing the "buttonOnly" template. -->
		{{block "buttonOnly" .}}
			<button hx-get="/?counter={{.next}}&template=buttonOnly" hx-swap="outerHTML">
				This Button Has Been Clicked {{.counter}} Times
			</button>
		{{end}}
	</body>
	</html>`))
}

var page *template.Template

func HandleFragment(w http.ResponseWriter, r *http.Request) {
	counter, _ := strconv.Atoi(r.URL.Query().Get("counter"))
	templateName := r.URL.Query().Get("template")
	if templateName == "" {
		templateName = "main"
	}

	data := make(map[string]int)
	data["counter"] = counter
	data["next"] = counter + 1
	if err := page.ExecuteTemplate(w, templateName, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
