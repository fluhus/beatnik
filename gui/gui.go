// Command gui provides a graphical web UI for the beatnik package.
package main

//go:generate go run gen.go

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/fluhus/beatnik"
)

const (
	mimeHTML = "text/html"
	mimeMIDI = "audio/midi"
	mimeText = "text/plain"
)

var (
	src       = flag.String("src", "", "Path of HTML source files, for development.")
	port      = flag.Uint("port", 8080, "Port to listen on.")
	indexFile = ""
)

func main() {
	flag.Parse()
	if *src != "" {
		indexFile = filepath.Join(*src, "index.html")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", mimeHTML)
		template.Must(indexPage(indexFile)).Execute(w, nil)
	})

	http.HandleFunc("/midi", func(w http.ResponseWriter, r *http.Request) {
		// Get source.
		r.ParseForm()
		src := r.FormValue("src")

		// Parse text.
		t, err := beatnik.ParseTrack(src)
		if err != nil {
			w.Header().Set("Content-Type", mimeText)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "Failed to parse source:", err)
			return
		}

		// Encode midi.
		midi, err := t.MarshalBinary()
		if err != nil {
			w.Header().Set("Content-Type", mimeText)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Failed to encode MIDI:", err)
			return
		}

		w.Header().Set("Content-Type", mimeMIDI)
		w.Header().Set("Content-Disposition",
			"attachment; filename=\"beat.mid\"")
		w.Write(midi)
	})

	fmt.Println("Listening")
	http.ListenAndServe(":"+fmt.Sprint(*port), nil)
}
