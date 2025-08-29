package main

import (
	"fmt"
	"os"
	"path"
)

type Entry struct {
	Name    string
	Surname string
	Tel     string
}

var data = []Entry{}

func search(key string) *Entry {
	for i, v := range data {
		if v.Surname == key {
			return &data[i]
		}
	}
	return nil
}

func list() {
	for _, v := range data {
		fmt.Println(v)
	}
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		exe := path.Base(arguments[0])
		fmt.Printf("Usage: %s search|list <arguments>\n", exe)
		return
	}

	data = append(data, Entry{"Mihalis", "Tsoukalos", "2109416471"})
	data = append(data, Entry{"Mary", "Doe", "21094168771"})
	data = append(data, Entry{"John", "Black", "2109416123"})

	// различаем команды
	switch arguments[1] {
	// команда поиска
	case "search":
		if len(arguments) != 3 {
			fmt.Println("Usage: search Surname")
			return
		}
		result := search(arguments[2])
		if result == nil {
			fmt.Println("Entry not found: ", arguments[2])
			return
		}
		fmt.Println(*result)
	// команда списка
	case "list":
		list()
	// ответ на все остальное кроме search и list
	default:
		fmt.Println("Not a valid option")

	}
}
