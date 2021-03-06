// Package beatnik defines data objects and a text language for encoding midi
// drum tracks.
package beatnik

// Type definitions.

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// TODO(amit): Encode tempo in a way that Reaper can recognize.
// TODO(amit): Add humanize feature.

// A Track is an entire drum track, with its drum data and metadata.
type Track struct {
	Hits []*Hit // Order of hits in this track.
	BPM  uint   // Track tempo.
}

// MarshalBinary returns a binary encoding of the track as a complete midi file.
func (t *Track) MarshalBinary() ([]byte, error) {
	if t.BPM == 0 {
		return nil, fmt.Errorf("cannot encode with bpm=0")
	}

	buf := bytes.NewBuffer(nil)
	buf.Write(t.encodeHeaderChunk())
	buf.Write(t.encodeMetaChunk())
	buf.Write(t.encodeHits())
	return buf.Bytes(), nil
}

// encodeHeaderChunk returns a binary encoding of the midi header track.
func (*Track) encodeHeaderChunk() []byte {
	buf := bytes.NewBuffer(nil)
	buf.Write([]byte("MThd"))
	binary.Write(buf, binary.BigEndian, uint32(6))
	binary.Write(buf, binary.BigEndian, uint16(1)) // File format (0/1/2).
	binary.Write(buf, binary.BigEndian, uint16(2)) // Number of tracks.
	binary.Write(buf, binary.BigEndian, uint16(96))

	return buf.Bytes()
}

// encodeMetaChunk returns a binary encoding of the midi first (metadata)
// track.
func (t *Track) encodeMetaChunk() []byte {
	// Extract us per beat from bpm.
	mpb := 1 / float64(t.BPM)
	uspb := uint32(mpb * 60 * 1000000)

	// Encode track.
	buf := bytes.NewBuffer(nil)
	buf.Write([]byte("MTrk"))

	buf2 := bytes.NewBuffer(nil)
	// TODO(amit): Extract meta events to functions.
	buf2.Write([]byte{0, 0xFF, 0x58, 4, 4, 2, 24, 8})
	buf2.Write([]byte{0, 0xFF, 0x51, 3})
	buf2.Write(bin(uspb)[1:])
	buf2.Write([]byte{0, 0xFF, 0x2F, 0})

	buf.Write(bin(uint32(buf2.Len())))
	return append(buf.Bytes(), buf2.Bytes()...)
}

// encodeHits returns a binary encoding of the drum hits in this track as a
// single midi track.
func (t *Track) encodeHits() []byte {
	buf := bytes.NewBuffer([]byte("MTrk"))
	buf2 := bytes.NewBuffer(nil)
	for _, h := range t.Hits {
		buf2.Write(h.encode())
	}
	buf2.Write([]byte{0, 0xFF, 0x2F, 0})
	buf.Write(bin(uint32(buf2.Len())))

	return append(buf.Bytes(), buf2.Bytes()...)
}

// A Hit is a set of drums being hit at the same time.
type Hit struct {
	Notes map[byte]Velocity // Notes to strike with their velocities.
	T     uint              // Number of ticks this hit lasts (96 is a quarter bar).
}

func (h *Hit) String() string {
	return fmt.Sprintf("{t=%v,%v}", h.T, h.Notes)
}

// copy returns a deep copy of a hit.
func (h *Hit) copy() *Hit {
	result := &Hit{
		map[byte]Velocity{},
		h.T,
	}
	for k, v := range h.Notes {
		result.Notes[k] = v
	}
	return result
}

// encode returns a binary encoding of the hit as midi events.
func (h *Hit) encode() []byte {
	buf := bytes.NewBuffer(nil)
	for n, v := range h.Notes {
		buf.Write([]byte{0, 0x99, n, byte(v)})
	}
	first := true
	for n := range h.Notes {
		if first {
			buf.Write(uvarint(h.T))
			first = false
		} else {
			buf.Write(uvarint(0))
		}
		buf.Write([]byte{0x89, n, 64})
	}
	return buf.Bytes()
}

// Velocity is a drum hit's volume.
type Velocity byte

// Predefined velocities.
const (
	PPP Velocity = 85  // Pianississimo
	PP           = 91  // Pianissimo
	P            = 97  // Piano
	MP           = 103 // Mezzo-piano
	MF           = 109 // Mezzo-forte
	F            = 115 // Forte
	FF           = 121 // Fortissimo
	FFF          = 127 // Fortississimo
)
