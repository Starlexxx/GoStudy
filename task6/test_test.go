package test

import (
	"sort"
	"testing"
)

func TestSort_Ints(t *testing.T) {
	data := []int{5, 2, 6, 3, 1, 4}
	sort.Ints(data)
	if !sort.IntsAreSorted(data) {
		t.Errorf("data is not sorted: %v", data)
	}
}

func TestSort_Strings(t *testing.T) {
	tests := []struct {
		name string
		data []string
		want []string
	}{
		{
			name: "normal",
			data: []string{"b", "a", "c"},
			want: []string{"a", "b", "c"},
		},
		{
			name: "empty",
			data: []string{},
			want: []string{},
		},
		{
			name: "one",
			data: []string{"a"},
			want: []string{"a"},
		},
		{
			name: "same",
			data: []string{"a", "a", "a"},
			want: []string{"a", "a", "a"},
		},
		{
			name: "reversed",
			data: []string{"c", "b", "a"},
			want: []string{"a", "b", "c"},
		},
		{
			name: "with empty",
			data: []string{"c", "", "a"},
			want: []string{"", "a", "c"},
		},
		{
			name: "with cyrillic",
			data: []string{"в", "а", "б"},
			want: []string{"а", "б", "в"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.Strings(tt.data)
			for s := range tt.data {
				if tt.data[s] != tt.want[s] {
					t.Errorf("data is not sorted: %v", tt.data)
				}
			}
		})
	}
}

func BenchmarkSort_Ints(b *testing.B) {
	unsortedArrays := generateIntArrays(1000, 1000)
	for i := 0; i < b.N; i++ {
		for _, arr := range unsortedArrays {
			sort.Ints(arr)
		}
	}
}

func generateIntArrays(n, m int) [][]int {
	arrays := make([][]int, n)
	for i := 0; i < n; i++ {
		arrays[i] = make([]int, m)
		for j := 0; j < m; j++ {
			arrays[i][j] = j
		}
	}

	return arrays
}

func BenchmarkSort_Floats(b *testing.B) {
	unsortedArrays := generateFloatArrays(1000, 1000)
	for i := 0; i < b.N; i++ {
		for _, arr := range unsortedArrays {
			sort.Float64s(arr)
		}
	}
}

func generateFloatArrays(n, m int) [][]float64 {
	arrays := make([][]float64, n)
	for i := 0; i < n; i++ {
		arrays[i] = make([]float64, m)
		for j := 0; j < m; j++ {
			arrays[i][j] = float64(j)
		}
	}

	return arrays
}

/*
Bencmark results:

Running tool: /opt/homebrew/bin/go test -benchmem -run=^$ -bench ^BenchmarkSort_Ints$ GoStudy/task6

goos: darwin
goarch: arm64
pkg: GoStudy/task6
=== RUN   BenchmarkSort_Ints
BenchmarkSort_Ints
BenchmarkSort_Ints-10                504           2340143 ns/op           40302 B/op       1001 allocs/op
PASS
ok      GoStudy/task6   2.167s


> Test run finished at 8/4/2023, 2:02:32 PM <

Running tool: /opt/homebrew/bin/go test -benchmem -run=^$ -bench ^BenchmarkSort_Floats$ GoStudy/task6

goos: darwin
goarch: arm64
pkg: GoStudy/task6
=== RUN   BenchmarkSort_Floats
BenchmarkSort_Floats
BenchmarkSort_Floats-10              506           2367257 ns/op           40238 B/op       1001 allocs/op
PASS
ok      GoStudy/task6   1.751s


> Test run finished at 8/4/2023, 2:03:01 PM <
*/
