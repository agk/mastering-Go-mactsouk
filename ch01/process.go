package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need one or more arguments!")
		return
	}

	var total, nInts, nFloats int
	invalid := make([]string, 0)
	for _, k := range arguments[1:] {
		// это целое число?
		_, err := strconv.Atoi(k)
		if err == nil {
			total++
			nInts++
			continue
		}
		// число в плавающей запятой float
		_, err = strconv.ParseFloat(k, 64)
		if err == nil {
			total++
			nFloats++
			continue
		}
		// осталось только недопустимое значение
		invalid = append(invalid, k)
	}

	fmt.Println("#read:", total, "#ints:", nInts, "#floats:", nFloats)
	if len(invalid) > total {
		fmt.Println("Too much invalid input:", len(invalid))
		for _, s := range invalid {
			fmt.Println(s)
		}
	}

}
