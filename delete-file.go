package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Remove("example.txt")
	if err != nil {
		//log.Fatal(err);
		fmt.Println("Error deleting file:", err)
		return 
	}

	fmt.Println("File deleted successfully");
}