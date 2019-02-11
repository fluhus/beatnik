package main

// A synchronous map that holds compiled MIDI files.

import (
	"sync"
	"time"
)

const (
	// How long to keep a compiled MIDI before removing it.
	midiFileTTL = time.Minute
)

var (
	midiFiles     = map[string][]byte{}
	midiFilesLock sync.Mutex
)

// getMIDIFile returns a MIDI file from the map.
// ok reflects if the file was present in the map.
func getMIDIFile(name string) (midi []byte, ok bool) {
	midiFilesLock.Lock()
	defer midiFilesLock.Unlock()
	midi, ok = midiFiles[name]
	return
}

// deleteMIDIFile deletes a MIDI file from the map.
func deleteMIDIFile(name string) {
	midiFilesLock.Lock()
	defer midiFilesLock.Unlock()
	delete(midiFiles, name)
}

// putMIDIFile adds a MIDI file to the map and deletes if after 1 minute.
func putMIDIFile(name string, midi []byte) {
	midiFilesLock.Lock()
	defer midiFilesLock.Unlock()
	midiFiles[name] = midi
	go func() {
		time.Sleep(midiFileTTL)
		deleteMIDIFile(name)
	}()
}
