package geometry

import "testing"

func TestGeom_Distance(t *testing.T) {
	tests := []struct {
		x1, y1, x2, y2 float64
		name           string
		want           float64
	}{
		{
			name: "same points",
			x1:   0,
			y1:   0,
			x2:   0,
			y2:   0,
			want: 0,
		},
		{
			name: "first point coordinates are greater",
			x1:   8,
			y1:   10,
			x2:   5,
			y2:   7,
			want: 4.242640687119285,
		},
		{
			name: "second point coordinates are greater",
			x1:   5,
			y1:   7,
			x2:   8,
			y2:   10,
			want: 4.242640687119285,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Distance(tt.x1, tt.y1, tt.x2, tt.y2); got != tt.want {
				t.Errorf("Distance() = %v, want %v", got, tt.want)
			}
		})
	}
}
