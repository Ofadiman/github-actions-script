package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("golang script executed")

	input := os.Getenv("INPUT_SCRIPT")
	fmt.Printf("selected file: %s\n", input)

	githubOutput := os.Getenv("GITHUB_OUTPUT")
	file, err := os.OpenFile(githubOutput, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	if _, err := file.WriteString("value=foo\n"); err != nil {
		log.Println(err)
	}

	err = file.Close()
	if err != nil {
		panic(err)
	}
}
