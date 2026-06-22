package main

import (
	"crypto/sha256"
	"fmt"
)

type Variables struct {
	numDecimal int
	numOctal   int
	numHex     int
	pi         float64
	name       string
	isActive   bool
	complexNum complex64
}

func main() {
	data := Variables{
		numDecimal: 42,
		numOctal:   052,
		numHex:     0x2A,
		pi:         3.14,
		name:       "Golang",
		isActive:   true,
		complexNum: 1 + 2i,
	}

	GetVarsTypes(data)
	PrintSaltedHash(data)

}

func GetVarsTypes(d Variables) map[string]string {
	vars := []struct {
		label string
		value any
	}{
		{"numDecimal", d.numDecimal},
		{"numOctal", d.numOctal},
		{"numHex", d.numHex},
		{"pi", d.pi},
		{"name", d.name},
		{"isActive", d.isActive},
		{"complexNum", d.complexNum},
	}

	typesMap := make(map[string]string)

	fmt.Printf("%-12s	|	%-12s	|	%s\n", "ПЕРЕМЕННАЯ", "ТИП", "ЗНАЧЕНИЕ")
	fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - - - - - - - -")

	for _, v := range vars {
		typesMap[v.label] = fmt.Sprintf("%T", v.value)
		fmt.Printf("%-8s	|	%-8T	|	%v\n", v.label, v.value, v.value)
	}

	return typesMap
}

func PrintSaltedHash(d Variables) string {
	origin := fmt.Sprintf("%v%v%v%v%v%v%v", d.numDecimal, d.numOctal, d.numHex, d.pi, d.name, d.isActive, d.complexNum)
	runedOrigin := []rune(origin)
	runedSalt := []rune("go-2024")

	middleIndex := len(runedOrigin) / 2
	firstHalf := runedOrigin[:middleIndex]
	secondHalf := runedOrigin[middleIndex:]

	saltedOrigin := append(firstHalf, append(runedSalt, secondHalf...)...)
	hash := sha256.Sum256([]byte(string(saltedOrigin)))
	hashStr := fmt.Sprintf("%x", hash)

	fmt.Printf("\n\"Солённый\" хэш: %s\n", hashStr)

	return hashStr
}
