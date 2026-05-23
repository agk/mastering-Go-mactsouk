package main

import "fmt"

func main() {
	// байтовый срез
	b := make([]byte, 12)
	fmt.Println("Byte slice:", b)

	b = []byte("Byte slice €")
	fmt.Println("Byte slice:", b)

	// вывести содержимое байтового среза в виде текста
	fmt.Printf("Byte slice as text: %s\n", b)
	fmt.Println("Byte slice as text:", string(b))

	// длина b
	fmt.Println("Length of b:", len(b))
}
