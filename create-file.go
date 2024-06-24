package main

import (
	"log"
	"os"
)

func main (){

	fileName := "example.txt"

	file, err := os.Create(fileName)
	if err != nil{
		log.Fatal(err)
	}

	defer file.Close();

	text := "This will be my file contents";
	_, err = file.WriteString(text)
	if err!= nil{
		log.Fatal(err)
	}

	log.Println("File created and edited");
}