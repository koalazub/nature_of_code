package server

import (
	"net/http"
	"os"

	r "github.com/koalazub/nature_of_code/renderer"
	"golang.org/x/exp/slog"
)

// New server
func New() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	// localhost/static/img.jpg
	addr := "127.0.0.1:8008"
	slog.Info("Listening on", "address", addr)

	http.HandleFunc("/static", r.HandleTemplate)
	http.HandleFunc("/", r.HandleIndex)
	http.HandleFunc("/vector", r.HandleVectorCanvas)
	http.ListenAndServe(addr, nil)
	os.Exit(1)
}
