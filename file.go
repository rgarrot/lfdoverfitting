package main

import (
	"fmt"
	"os"
	"strconv"
)

var path = "C:\\Users\\Renan\\Desktop\\base.csv"

func writeBase(b Base) {
	createFile()
	writeFile(b)
}

func createFile() {
	// detect if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		checkError(err)
		defer file.Close()
	}
}

func writeFile(b Base) {
	// open file using READ & WRITE permission
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	checkError(err)
	defer file.Close()
	for i := 0; i < len(b.Y); i++ {
		_, err = file.WriteString(strconv.FormatFloat(b.X[i], 'f', -1, 64))
		checkError(err)
		_, err = file.WriteString(",")
		checkError(err)
		_, err = file.WriteString(strconv.FormatFloat(b.Y[i], 'f', -1, 64))
		checkError(err)
		_, err = file.WriteString("\n")
		checkError(err)
	}

	// write some text to file
	// _, err = file.WriteString("halo\n")
	// checkError(err)
	// _, err = file.WriteString("mari belajar golang\n")
	// checkError(err)

	// save changes
	err = file.Sync()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}
