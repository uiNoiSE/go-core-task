package main

import (
	"fmt"
	"maps"
)

type StringIntMap struct {
	data map[string]int
}

func main() {
	first := NewStringIntMap()

	first.Add("apple", 1)
	first.Add("banana", 2)
	fmt.Printf("Add()	|	Структура(first) после двух Add():\n%v\n", first)

	first.Remove("apple")
	fmt.Printf("\nRemove(\"apple\")	|	Структура(first) после удаления:\n%v\n", first)

	second := first.Copy()
	fmt.Printf("\nCopy()	|	Структура(second) после копирования структуры(first): \n%v\n", second)

	second["tomato"] = 3
	second["potato"] = 4
	fmt.Printf("Структура(second) после ручного добавления элементов: \n%v\n", second)

	fmt.Println("\n--- Проверка иммутабельности ---")
	fmt.Printf("Exists(\"tomato\")	|	Есть ли ключ в оригинале (first)? %v\n", first.Exists("tomato"))
	fmt.Printf("Оригинальная структура (first) осталась прежней:\n%v\n", first)

	val, ok := first.Get("banana")
	fmt.Printf("Get(\"banana\")	|	Значение - %v, наличие ключа - %v\n", val, ok)
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{
		data: make(map[string]int),
	}
}

func (s *StringIntMap) Add(key string, val int) {
	s.data[key] = val
}

func (s *StringIntMap) Remove(key string) {
	delete(s.data, key)
}

func (s *StringIntMap) Copy() map[string]int {
	copied := make(map[string]int)
	maps.Copy(copied, s.data)

	return copied
}

func (s *StringIntMap) Exists(key string) bool {
	_, isExists := s.data[key]
	return isExists
}

func (s *StringIntMap) Get(key string) (int, bool) {
	val, ok := s.data[key]
	return val, ok
}
