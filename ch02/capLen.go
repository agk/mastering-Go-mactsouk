package main

import "fmt"

func main() {
	// определена только длина, Capacity = length
	a := make([]int, 4)
	fmt.Println("L:", len(a), "C:", cap(a))

	// инициализировать срез, Capacity = length
	b := []int{0, 1, 2, 3, 4}
	fmt.Println("L:", len(b), "C:", cap(b))

	// одинаковая длина и емкость
	aSlice := make([]int, 4, 4)
	fmt.Println(aSlice)
	// добавить элемент
	aSlice = append(aSlice, 5)
	fmt.Println(aSlice)
	// емкость увеличивается вдвое
	fmt.Println("L:", len(aSlice), "C:", cap(aSlice))

	// теперь добавим еще 4 элемента
	aSlice = append(aSlice, []int{-1, -2, -3, -4}...)
	// операция ... разворачивает []int{-1,-2,-3,-4} в несколько аргументов
	// и append() по отедельнсти добавляет каждый аргумент в aSlice
	fmt.Println(aSlice)
	// емкость увеличивается вдвое
	fmt.Println("L:", len(aSlice), "C:", cap(aSlice))

}
