package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println(os.Environ())

	fmt.Println("golang script executed")

	input := os.Getenv("INPUT_MESSAGE")
	fmt.Printf("selected file: %v\n", input)

	githubOutput := os.Getenv("GITHUB_OUTPUT")

	bytes, _ := os.ReadFile(githubOutput)
	fmt.Println(string(bytes))

	file, err := os.OpenFile(githubOutput, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	if _, err := file.WriteString("value=foo\n"); err != nil {
		log.Println(err)
	}

	bytes, _ = os.ReadFile(githubOutput)
	fmt.Println(string(bytes))

	err = file.Close()
	if err != nil {
		panic(err)
	}
}
