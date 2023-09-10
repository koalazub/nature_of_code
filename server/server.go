package server

import (
	"html/template"
	"net/http"
	"os"

	"golang.org/x/exp/slog"
)

// New server
func New() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	// localhost/static/img.jpg
	slog.Info("Listening on: http://127.0.0.1:8008")
	http.HandleFunc("/static", handleTemplate)
	http.ListenAndServe("127.0.0.1:8008", nil)
	os.Exit(1)
}

func handleDumb(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("dumb.html")
	if err != nil {
		slog.Error("Couldn't find `html` file")
	}
	err = tmpl.Execute(w, nil)
}

func handleTemplate(w http.ResponseWriter, r *http.Request) {
	const t = `<h1>Here is some text</h1>`

	tmpl, err := template.New("example").Parse(t)
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

	err = tmpl.Execute(w, data)
	if err != nil {
		slog.Error("couldn't apply template", err)
	}
}
