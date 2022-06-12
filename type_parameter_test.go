package Go_lang_Generic

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestLength(t *testing.T) {
	var result string = Length[string]("Eko")
	fmt.Println(result)

	var resultNumber int = Length[int](100)
	fmt.Print(resultNumber)
}

func kali(a int, b int, c int) int {
	hasil := a * b * c
	return hasil
}

func TestSample(t *testing.T) {
	hasil := kali(10, 3, 9)
	assert.Equal(t, 30, hasil)

	hasil = kali(40, 2, 8)
	assert.Equal(t, 40, hasil)
	assert.True(t, true)
}

func tambah(a int, b int, c int) int {
	hasil := a + b + c
	return hasil
}

func TestSample1(t *testing.T) {
	hasil := tambah(10, 10, 2)
	assert.Equal(t, 30, hasil)
}

func kurang(a int, b int, c int) int {
	hasil := a - b - c
	return hasil
}

func TestSample2(t *testing.T) {
	hasil := kurang(10, 4, 9)
	assert.Equal(t, 20, hasil)
}

func bagi(a int, b int, c int) int {
	hasil := a / b / c
	return hasil
}

func TestSample3(t *testing.T) {
	hasil := bagi(10, 10, 2)
	assert.Equal(t, 10, hasil)

	// type parameter
	var result1 = Length[string]("Roni")
	assert.Equal(t, "Roni", result1)

	var resultNumber = Length[int](100)
	assert.Equal(t, 100, resultNumber)

	var resultBool = Length[bool](true)
	assert.True(t, resultBool)

	var resultFloat = Length[float64](100.30)
	assert.Equal(t, 300.30, resultFloat)

}
