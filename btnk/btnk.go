// Command btnk is a command line interface for the beatnik library.
package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fluhus/beatnik"
)

func main() {
	exitCode := 0
	for _, f := range os.Args[1:] {
		fmt.Printf("reading %q\n", f)
		d, err := ioutil.ReadFile(f)
		if err != nil {
			fmt.Printf("failed to read %q: %v\n", f, err)
			exitCode = 2
			continue
		}
		t, err := beatnik.ParseTrack(string(d))
		if err != nil {
			fmt.Printf("failed to parse %q: %v\n", f, err)
			exitCode = 2
			continue
		}
		b, err := t.MarshalBinary()
		if err != nil {
			fmt.Printf("failed to encode: %v\n", err)
			exitCode = 2
			continue
		}
		err = ioutil.WriteFile(f+".mid", b, 0600)
		if err != nil {
			fmt.Printf("failed to write %q: %v\n", f+".mid", err)
			exitCode = 2
			continue
		}
	}
	os.Exit(exitCode)
}
