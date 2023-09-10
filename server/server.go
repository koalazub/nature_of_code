package server

import (
	"html/template"
	"net/http"
	"os"

	r "github.com/koalazub/nature_of_code/renderer"
	"golang.org/x/exp/slog"
)

// New server
func New() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	// localhost/static/img.jpg
	slog.Info("Listening on: http://127.0.0.1:8008")
	http.HandleFunc("/static", r.HandleTemplate)
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
