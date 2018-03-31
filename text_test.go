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
		{"42", NewHit(96*4, F, 42)},
		{"36,49,57..+", NewHit(96, FF, 49, 57, 36)},
		{"36,49,57....----", NewHit(24, PP, 49, 57, 36)},
	}

	for _, test := range tests {
		got, err := parseHit(test.in)
		if err != nil {
			t.Errorf("parseHit(%v), want success: %v", test.in, err)
			continue
		}
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("parseHit(%v)=%v, want %v", test.in, got, test.want)
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
		"42+.",
		"11",
		"0",
		"127",
	}

	for _, test := range tests {
		if got, err := parseHit(test); err == nil {
			t.Errorf("parseHit(%v)=%v, want failure", test, got)
		}
	}
}

func TestParseTrack(t *testing.T) {
	in := "bpm:123  42   \t\n  36,49,57..+\n\n\t\t  \t\t\n\n36,49,57....----"
	want := &Track{
		Hits: []*Hit{
			NewHit(96*4, F, 42),
			NewHit(96, FF, 49, 57, 36),
			NewHit(24, PP, 49, 57, 36),
		},
		BPM: 123,
	}
	got, err := ParseTrack(in)
	if err != nil {
		t.Fatalf("parseHits(%v) should succeed, but failed: %v", in, err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("parseHits(%v)=%v, want %v", in, got, want)
	}
}
