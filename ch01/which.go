package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide an argument!")
		return
	}

	file := arguments[1]
	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)

	for _, directory := range pathSplit {
		fullPath := filepath.Join(directory, file)
		// существует ли путь?
		fileInfo, err := os.Stat(fullPath)
		if err == nil {
			mode := fileInfo.Mode()
			// это обысный файл
			if mode.IsRegular() {
				// является ли он исполняемым?
				if mode&0111 != 0 {
					fmt.Println(fullPath)
				}
			}
		}
	}
}
