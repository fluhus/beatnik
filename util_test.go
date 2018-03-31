package beatnik

import (
	"reflect"
	"testing"
)

func TestUvarint(t *testing.T) {
	tests := []struct {
		in   uint
		want []byte
	}{
		{0, []byte{0}},
		{1, []byte{1}},
		{127, []byte{127}},
		{128, []byte{129, 0}},
		{129, []byte{129, 1}},
		{130, []byte{129, 2}},
	}

	for _, test := range tests {
		got := uvarint(test.in)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("uvarint(%v)=%v, want %v", test.in, got, test.want)
		}
	}
}
