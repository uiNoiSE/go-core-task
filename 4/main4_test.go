package main

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	tests := []struct {
		name           string
		input1         []string
		input2         []string
		expectedOutput []string
	}{
		{
			name:           "Уникальные значения в обоих слайсах",
			input1:         []string{"1", "2", "3", "4", "5"},
			input2:         []string{"6", "7", "8", "9", "0"},
			expectedOutput: []string{"1", "2", "3", "4", "5"},
		},
		{
			name:           "Идентичные слайсы",
			input1:         []string{"1", "2", "3", "4", "5"},
			input2:         []string{"1", "2", "3", "4", "5"},
			expectedOutput: []string{},
		},
		{
			name:           "\"Фильтрующий\" слайс пуст",
			input1:         []string{"1", "2", "3", "4", "5"},
			input2:         []string{},
			expectedOutput: []string{"1", "2", "3", "4", "5"},
		},
		{
			name:           "Исходный(фильтруемый) слайс пуст",
			input1:         []string{},
			input2:         []string{"1", "2", "3", "4", "5"},
			expectedOutput: []string{},
		},
		{
			name:           "Дубликаты в исходном слайсе",
			input1:         []string{"1", "1", "2", "2", "3", "4", "5"},
			input2:         []string{"1"},
			expectedOutput: []string{"2", "2", "3", "4", "5"},
		},
		{
			name:           "Разный регистр",
			input1:         []string{"A", "B", "C", "d", "e", "f"},
			input2:         []string{"a", "B"},
			expectedOutput: []string{"A", "C", "d", "e", "f"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Filter(test.input1, test.input2)

			if !reflect.DeepEqual(result, test.expectedOutput) {
				t.Errorf("Ожидали: %v\nПолучили: %v\n", test.expectedOutput, result)
			}
		})
	}
}
