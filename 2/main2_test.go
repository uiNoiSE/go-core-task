package main

import (
	"reflect"
	"testing"
)

func TestSliceExample(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Чётные + нечётные",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: []int{2, 4, 6, 8, 10},
		},
		{
			name:     "Только нечётные",
			input:    []int{1, 3, 5, 7, 9},
			expected: nil,
		},
		{
			name:     "Пустой слайс",
			input:    []int{},
			expected: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := SliceExample(test.input)

			if !reflect.DeepEqual(actual, test.expected) && !(len(actual) == 0 && len(test.expected) == 0) {
				t.Errorf("Ожидали: %v, получили: %v", test.expected, actual)
			}
		})
	}
}

func TestAddElements(t *testing.T) {
	input := []int{1, 2, 3}
	expected := []int{1, 2, 3, 4}

	actual := AddElements(input, 4)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Результат: %v, не совпал с ожидаемым: %v", actual, expected)
	}
}

func TestCopySlice(t *testing.T) {
	input := []int{10, 20, 30}
	copied := CopySlice(input)

	if !reflect.DeepEqual(copied, input) {
		t.Fatalf("Копия: %v, не совадает с оригиналом: %v", copied, input)
	}

	input[0] = 1111
	if copied[0] == 1111 {
		t.Errorf("Измение оригинала: %v повлияло на копию: %v", input, copied)
	}
}

func TestRemoveElement(t *testing.T) {
	input := []int{10, 15, 20, 30, 40}
	expected := []int{10, 20, 30, 40}

	actual := RemoveElement(input, 1)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Результат: %v, не совпал с ожидаемым: %v", actual, expected)
	}

	if len(input) != 5 || input[1] != 15 {
		t.Errorf("Исходный слайс был мутирован: %v", input)
	}
}
