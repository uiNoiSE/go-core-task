package main

import (
	"testing"
	"time"
)

func TestNewRandomGenerator_EmitsValues(t *testing.T) {
	ch := NewRandomGenerator()

	for range 3 {
		select {
		case num := <-ch:
			_ = num
		case <-time.After(100 * time.Millisecond):
			t.Fatalf("Тест завис: Генератор не отдал число вовремя")
		}
	}
}

func TestNewRandomGenerator_Timeout(t *testing.T) {
	ch := NewRandomGenerator()

	select {
	case <-ch:
	case <-time.After(100 * time.Millisecond):
		t.Fatalf("Тест упал по таймауту")
	}
}

func TestNewRandomGenerator_IsRandom(t *testing.T) {
	ch := NewRandomGenerator()

	num1 := <-ch
	num2 := <-ch

	if num1 == num2 {
		t.Errorf("Генератор выдал два одинаковых числа подряд: %d %d", num1, num2)
	}
}
