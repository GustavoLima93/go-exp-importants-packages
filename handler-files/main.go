package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("file.txt")

	if err != nil {
		panic(err)
	}

	fmt.Printf("File created successfully !! \n")

	f, error := file.WriteString("Hello, World! \n")

	if error != nil {
		panic(error)
	}

	fmt.Printf("File written successfully: %d \n", f)

	f, error = file.Write([]byte("Hello, World - 2! \n"))

	if error != nil {
		panic(error)
	}

	fmt.Printf("File written successfully: %d \n", f)

	file.Close()

	data, err := os.ReadFile("file.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

	newFile, err := os.Open("file.txt")

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(newFile)
	buffer := make([]byte, 5)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	newFile.Close()

	err = os.Remove("file.txt")

	if err != nil {
		panic(err)
	}

}
