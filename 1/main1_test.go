package main

import (
	"testing"
)

func TestGetVarsTypes(t *testing.T) {
	testData := Variables{
		numDecimal: 42,
		numOctal:   052,
		numHex:     0x2A,
		pi:         3.14,
		name:       "Golang",
		isActive:   true,
		complexNum: 1 + 2i,
	}

	actualTypes := GetVarsTypes(testData)

	expectedTypes := map[string]string{
		"numDecimal": "int",
		"numOctal":   "int",
		"numHex":     "int",
		"pi":         "float64",
		"name":       "string",
		"isActive":   "bool",
		"complexNum": "complex64",
	}

	for key, expectedType := range expectedTypes {
		actualType, exists := actualTypes[key]

		if !exists {

			continue
		}

		if actualType != expectedType {
			t.Errorf("Для поля \"%s\" ожидали тип: %s, получили: %s", key, expectedType, actualType)
		}
	}
}

func TestGetSaltedHash(t *testing.T) {
	testData := Variables{
		numDecimal: 42,
		numOctal:   052,
		numHex:     0x2A,
		pi:         3.14,
		name:       "Golang",
		isActive:   true,
		complexNum: 1 + 2i,
	}

	expectedHash := "53f2f60ac6c41389d3ed3d84d88d8c2860bf8981c677be18243a6f35a6b6a1b3"
	actualHash := PrintSaltedHash(testData)

	if actualHash != expectedHash {
		t.Errorf("Хэш не совпал.\n Ожидали: %s\nПолучили: %s", expectedHash, actualHash)
	}
}
