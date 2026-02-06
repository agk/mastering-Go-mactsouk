package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	fmt.Println(os.TempDir())
	LOGFILE := path.Join(os.TempDir(), "mGo.log")
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	// вызов os.OpenFile() файл журнала для записи,
	// еслион не существует, или же открывает его для записи,
	// путем добавления новых данных в конце (os.O_APPEND)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	LstdFlags := log.Ldate | log.Lshortfile
	iLog := log.New(f, "iLog ", LstdFlags)
	iLog.Println("Hello there!")
	iLog.Println("Mastering Go 3rd edition")

	iLog.SetFlags(log.Lshortfile | log.LstdFlags)
	iLog.Println("another log entry!")

}
