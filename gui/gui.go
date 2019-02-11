// Command gui provides a graphical web UI for the beatnik package.
package main

//go:generate go run gen.go

import (
	"flag"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"path/filepath"
	"time"

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
	rand.Seed(time.Now().UnixNano())

	// Main page handler.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", mimeHTML)
		template.Must(indexPage(indexFile)).Execute(w, nil)
	})

	// Compile handler.
	http.HandleFunc("/compile", func(w http.ResponseWriter, r *http.Request) {
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

		file := fmt.Sprintf("%v.mid", rand.Int63())
		putMIDIFile(file, midi)

		w.Write([]byte(file))
	})

	// MIDI handler.
	http.HandleFunc("/midi/", func(w http.ResponseWriter, r *http.Request) {
		file := r.URL.Path[len("/midi/"):]

		midi, ok := getMIDIFile(file)
		if !ok {
			w.Header().Set("Content-Type", mimeText)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "MIDI file not found:", file)
			return
		}

		w.Header().Set("Content-Type", mimeMIDI)
		w.Header().Set("Content-Disposition",
			"attachment; filename=\""+file+"\"")
		w.Write(midi)
	})

	fmt.Println("Listening")
	http.ListenAndServe(":"+fmt.Sprint(*port), nil)
}
