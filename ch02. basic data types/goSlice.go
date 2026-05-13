package main

import "fmt"

func main() {
	// создать пустой срез
	aSlice := []float64{}
	// и длина, и емкость равны 0, поскольку aSlice пуст
	fmt.Println(aSlice, len(aSlice), cap(aSlice))

	// добавить элементы в срез
	aSlice = append(aSlice, 1234.56)
	aSlice = append(aSlice, -34.0)
	fmt.Println(aSlice, "with length", len(aSlice))

	// срез длинной 4
	t := make([]int, 4)
	t[0] = -1
	t[1] = -2
	t[2] = -3
	t[3] = -4
	// теперь нужно использовать append
	t = append(t, -5)
	fmt.Println(t)

	// двумерный срез
	// у вас может быть столько измерений, сколько необходимо
	twoD := [][]int{{1, 2, 3}, {4, 5, 6}}
	// просмотр всех элементов 2D-среза
	// с помощью двойного цикла for
	for _, i := range twoD {
		for _, k := range i {
			fmt.Println(k, " ")
		}
		fmt.Println()
	}

	// показано как создать переменную 2D-среза
	// и одновременно ее инициализировать
	make2D := make([][]int, 2)
	fmt.Println(make2D)
	make2D[0] = []int{1, 2, 3, 4}
	make2D[1] = []int{-1, -2, -3, -4}
	fmt.Println(make2D)

	fmt.Println(twoD)
}
