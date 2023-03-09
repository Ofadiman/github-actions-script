package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("golang script executed")

	message := flag.String("message", "default value", "description")
	flag.Parse()
	fmt.Printf("parsed cli flag: %v\n", *message)

	githubOutput := os.Getenv("GITHUB_OUTPUT")

	bytes, _ := os.ReadFile(githubOutput)
	fmt.Println("\"\"\"")
	fmt.Println(string(bytes))
	fmt.Println("\"\"\"")

	file, err := os.OpenFile(githubOutput, os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}

	if _, err = file.WriteString("value=foo\n"); err != nil {
		log.Println(err)
	}

	bytes, _ = os.ReadFile(githubOutput)
	fmt.Println("\"\"\"")
	fmt.Println(string(bytes))
	fmt.Println("\"\"\"")

	err = file.Close()
	if err != nil {
		panic(err)
	}
}
