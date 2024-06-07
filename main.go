package main

import (
	"fmt"
	"os"
)

func main() {
	read := readFile("C:\\Users\\Cristina\\Desktop\\create-read-files\\text.txt")
	createFile(read)
	fmt.Println("Texto transferido com sucesso!")
}

func readFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func createFile(content string) {
	file, err := os.Create("C:\\Users\\Cristina\\Desktop\\create-read-files\\result.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data := []byte(content)
	nb, err := file.Write(data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Foram escritos %d bytes\n", nb)
}
