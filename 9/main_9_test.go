package main

import (
	"testing"
)

func TestCubePipeline(t *testing.T) {
	tests := []struct {
		name  string
		input []uint8
		want  []float64
	}{
		{"Нулевые значения", []uint8{0}, []float64{0}},
		{"Легальные uint8 значения", []uint8{2, 4, 16}, []float64{8, 64, 4096}},
		{"Максимальный uint8(255)", []uint8{255}, []float64{16581375}},
		{"Пустой слайс (проверка на пустоту)", []uint8{}, []float64{}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			in := make(chan uint8)
			out := make(chan float64)

			go CubePipeline(in, out)

			go func() {
				for _, v := range tc.input {
					in <- v
				}
				close(in)
			}()

			var results []float64
			for res := range out {
				results = append(results, res)
			}

			if len(results) != len(tc.want) {
				t.Fatalf("Ожидали :%d результатов, получили: %d", len(tc.want), len(results))
			}

			for i, res := range results {
				if res != tc.want[i] {
					t.Errorf("Для %d ожидали куб = %f, но получили: %f", tc.input[i], tc.want[i], res)
				}
			}
		})
	}
}
