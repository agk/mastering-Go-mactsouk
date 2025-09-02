package main

import (
	"fmt"
)

func main() {
	aString := "Hello World! $"
	fmt.Println("Fist character", string(aString[0]))

	// руны
	// руна
	r := '$'

	// преобразование рун в текст
	fmt.Printf("As a string: %s and as a character: %c\n", r, r)

	// вывести существующую строку в виде рун
	for _, v := range aString {
		fmt.Printf("%x", v)
	}
	fmt.Println()

	// вывести существующую строку в виде символов
	for _, v := range aString {
		fmt.Printf("%c", v)
	}
	fmt.Println()

}
