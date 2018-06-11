// Command gui provides a graphical web UI for the beatnik package.
package main

//go:generate go run gen.go

import (
	"flag"
	"fmt"
	"net/http"
	
	"github.com/fluhus/beatnik"
	//"github.com/fluhus/rpk"
)

var (
	mime = struct{ html, midi, text string }{"text/html", "audio/midi", "text/plain"}
	src  = flag.String("src", "", "Path of HTML source files, for development.")
)

func main() {
	flag.Parse()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", mime.html)
		indexPageTemplate.Execute(w, nil)
	})
	http.HandleFunc("/midi", func(w http.ResponseWriter, r *http.Request) {
	// Get source.
		r.ParseForm()
		src := r.FormValue("src")
		
		// Parse text.
		t, err :=beatnik.ParseTrack(src)
		if err != nil {
			w.Header().Set("Content-Type", mime.text)
			w.WriteHeader( http.StatusBadRequest)
			fmt.Fprintln(w, "Failed to parse source:", err)
			return
		}
		
		// Encode midi.
		midi, err := t.MarshalBinary()
		if err != nil {
			w.Header().Set("Content-Type", mime.text)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Failed to encode MIDI:", err)
			return
		}
		
		w.Header().Set("Content-Type", mime.midi)
		w.Write(midi)
	})

	fmt.Println("Listening")
	http.ListenAndServe(":8080", nil)
}
