package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	message := flag.String("message", "", "")
	flag.Parse()
	fmt.Printf("parsed message flag: %s\n", *message)

	githubOutput := os.Getenv("GITHUB_OUTPUT")
	file, err := os.OpenFile(githubOutput, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	if _, err = file.WriteString(fmt.Sprintf("value=example\n")); err != nil {
		panic(err)
	}
	err = file.Close()
	if err != nil {
		panic(err)
	}
}
