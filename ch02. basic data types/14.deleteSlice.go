package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need an integer value.")
		return
	}

	index := arguments[1]
	i, err := strconv.Atoi(index)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Using index", i)
	aSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("Original slice:", aSlice)
	// удалить элемент с индексом i
	if i > len(aSlice)-1 {
		fmt.Println("Cannot delete element", i)
		return
	}

	// Операция ... автоматически развертывает aSlice[i+1:] так,
	// что его элементы можно добавить к aSlice[:i] последовательно
	aSlice = append(aSlice[:i], aSlice[i+1:]...)
	fmt.Println("After 1st deletion:", aSlice)
	// Здесь мы логически разделяем исходный срез на два.
	// Два среза разделяются по индексу элемента, который необходимо удалить.
	// После этого мы объединяем эти два среза с помощью операции ....

	// удалить элемент с индексом i
	if i > len(aSlice)-1 {
		fmt.Println("Cannot delete element", i)
		return
	}
	// заменить элемент с индексом i на последний
	aSlice[i] = aSlice[len(aSlice)-1]
	// удалить последний элемент
	aSlice = aSlice[:len(aSlice)-1]
	fmt.Println("After 2nd deletion:", aSlice)

}
