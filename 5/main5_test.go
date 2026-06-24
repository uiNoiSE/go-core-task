package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestIntersection(t *testing.T) {
	tests := []struct {
		name         string
		input1       []int
		input2       []int
		wantedFound  bool
		wantedResult []int
	}{
		{
			name:         "Нет пересечений",
			input1:       []int{3, 2, 1, 5, 4},
			input2:       []int{9, 7, 8, 6, 0, 11},
			wantedFound:  false,
			wantedResult: []int{},
		},
		{
			name:         "Один из слайсов пуст",
			input1:       []int{},
			input2:       []int{6, 7, 8, 9, 0},
			wantedFound:  false,
			wantedResult: []int{},
		},
		{
			name:         "Слайсы идентичны",
			input1:       []int{1, 2, 3, 4, 5},
			input2:       []int{4, 5, 1, 3, 2},
			wantedFound:  true,
			wantedResult: []int{1, 2, 3, 4, 5},
		},
		{
			name:         "Отрицательные числа",
			input1:       []int{1, -2, -3, -4, -1, 0, -5},
			input2:       []int{3, 2, 1, 4, 5},
			wantedFound:  true,
			wantedResult: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFound, gotResult := Intersection(tt.input1, tt.input2)

			if gotFound != tt.wantedFound {
				t.Errorf("Intersection() = %v, want %v", gotFound, tt.wantedFound)
			}

			sort.Ints(gotResult)
			if !reflect.DeepEqual(gotResult, tt.wantedResult) {
				t.Errorf("Intersection() = %v, want %v", gotResult, tt.wantedResult)
			}
		})
	}
}
