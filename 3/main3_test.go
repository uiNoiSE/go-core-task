package main

import (
	"maps"
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	m := NewStringIntMap()

	m.Add("apple", 1)

	val, ok := m.data["apple"]
	if !ok {
		t.Errorf("Add() не создал запись для ключа 'apple'")
	}
	if val != 1 {
		t.Errorf("Add() записал %d, а ожидали 1", val)
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		name    string
		input   map[string]int
		key     string
		wantVal int
	}{
		{
			name:    "Успешное удаление",
			input:   map[string]int{"apple": 1, "banana": 2},
			key:     "banana",
			wantVal: 0,
		},
		{
			name:    "Удаление по несуществующему ключу",
			input:   map[string]int{"apple": 1, "banana": 2},
			key:     "orange",
			wantVal: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := NewStringIntMap()

			maps.Copy(m.data, test.input)

			m.Remove(test.key)

			_, exists := m.data[test.key]
			if exists {
				t.Errorf("Remove() не удалил значение по ключу: %v", test.key)
			}
		})
	}
}

func TestCopy(t *testing.T) {
	m := NewStringIntMap()
	m.data["banana"] = 1

	actualCopy := m.Copy()
	expectedCopy := map[string]int{"banana": 1}

	if !reflect.DeepEqual(actualCopy, expectedCopy) {
		t.Fatalf("Copy() вернул некорректную структуру. Ожидали: %v, получили: %v", expectedCopy, actualCopy)
	}

	actualCopy["apple"] = 2
	_, ok := m.data["apple"]
	if ok {
		t.Errorf("Copy() нарушает иммутабельность оригинальной структуры")
	}
}

func TestExists(t *testing.T) {
	m := NewStringIntMap()
	m.data["apple"] = 1

	if !m.Exists("apple") {
		t.Errorf("Exists() вернул false для существующего ключа")
	}

	if m.Exists("not_exists") {
		t.Errorf("Exists() вернул true для несуществующего ключа")
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		name    string
		input   map[string]int
		key     string
		wantVal int
		wantOk  bool
	}{
		{
			name:    "Успешное получение",
			input:   map[string]int{"banana": 2},
			key:     "banana",
			wantVal: 2,
			wantOk:  true,
		},
		{
			name:    "Ключ не найден",
			input:   map[string]int{"banana": 2},
			key:     "orange",
			wantVal: 0,
			wantOk:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewStringIntMap()

			maps.Copy(m.data, tt.input)

			gotVal, gotOk := m.Get(tt.key)

			if gotOk != tt.wantOk || gotVal != tt.wantVal {
				t.Errorf("Из Get() получили: (%v, %v), ожидали: (%v, %v)", gotVal, gotOk, tt.wantVal, tt.wantOk)
			}
		})
	}
}
