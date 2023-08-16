package writer

import (
	"bytes"
	"io"
	"testing"
)

func TestWrite(t *testing.T) {
	tests := []struct {
		name    string
		w       io.Writer
		strings []interface{}
		want    string
	}{
		{
			name:    "normal",
			w:       &bytes.Buffer{},
			strings: []interface{}{"a", "b", "c"},
			want:    "abc",
		},
		{
			name:    "empty",
			w:       &bytes.Buffer{},
			strings: []interface{}{},
			want:    "",
		},
		{
			name:    "nil",
			w:       &bytes.Buffer{},
			strings: nil,
			want:    "",
		},
		{
			name:    "with non-string",
			w:       &bytes.Buffer{},
			strings: []interface{}{"a", 1, "b", 2, "c", 3},
			want:    "abc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Write(tt.w, tt.strings)
			got := tt.w.(*bytes.Buffer).String()
			if got != tt.want {
				t.Errorf("Write() = %v, want %v", got, tt.want)
			}
		})
	}
}
