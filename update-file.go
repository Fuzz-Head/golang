package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("example.txt", os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}

	defer file.Close()

	if _, err := file.WriteString("\n This was appended to the file"); err != nil {
		fmt.Println("File could not be appended successfully")
		return
	}

	fmt.Println("File updated successfully")
}
