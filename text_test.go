package beatnik

import (
	"reflect"
	"testing"
)

// TODO(amit): Add comments to test texts.

func TestParseHit(t *testing.T) {
	tests := []struct {
		in   string
		want *Hit
	}{
		{"42~~", &Hit{map[byte]Velocity{42: F}, 96 * 4}},
		{"38-..", &Hit{map[byte]Velocity{38: MF}, 96 / 4}},
		{"36+,49,57+", &Hit{map[byte]Velocity{49: F, 57: FF, 36: FF}, 96}},
		{"36----,49---,57++..", &Hit{map[byte]Velocity{49: P, 57: FFF, 36: PP}, 24}},
	}

	for i, test := range tests {
		got, err := parseHit(test.in)
		if err != nil {
			t.Errorf("#%v/%v parseHit(%v), want success: %v",
				i+1, len(tests), test.in, err)
			continue
		}
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("#%v/%v parseHit(%v)=%v, want %v",
				i+1, len(tests), test.in, got, test.want)
		}
	}
}

func TestParseHit_badInput(t *testing.T) {
	tests := []string{
		"",
		"ananas",
		"42,",
		"42+-",
		"42-------------",
		"42.............",
		"42~~~~~~~~~~~",
		"42.+",
		"0",
		".",
	}

	for i, test := range tests {
		if got, err := parseHit(test); err == nil {
			t.Errorf("#%v/%v parseHit(%v)=%v, want failure",
				i+1, len(tests), test, got)
		}
	}
}

func TestParseTrack(t *testing.T) {
	want := &Track{
		Hits: []*Hit{
			&Hit{map[byte]Velocity{42: F}, 96 * 4},
			&Hit{map[byte]Velocity{49: FF, 57: FF, 36: FF}, 96},
			&Hit{map[byte]Velocity{49: PP, 57: PP, 36: PP}, 24},
			&Hit{map[byte]Velocity{46: F, 36: F}, 96 + 48 + 24},
			&Hit{map[byte]Velocity{38: F, 42: F}, 96},
		},
		BPM: 123,
	}
	got, err := ParseTrack(testTrack)
	if err != nil {
		t.Fatalf("parseHits(%v) should succeed, but failed: %v", testTrack, err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("parseHits(%v)=%v, want %v", testTrack, got, want)
	}
}

var testTrack = `bpm:123  # Track tempo.
42~~#~  	
  36+,49+,57+
  
# Just a comment.
		36----,49----,57----.. 36,46 . .. 38,42
`
