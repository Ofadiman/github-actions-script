package main

import (
	"flag"
	"fmt"
	"os"
)

func output(key string, value string) {
	githubOutput := os.Getenv("GITHUB_OUTPUT")

	file, err := os.OpenFile(githubOutput, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	if _, err = file.WriteString(fmt.Sprintf("%s=%s\n", key, value)); err != nil {
		panic(err)
	}

	err = file.Close()
	if err != nil {
		panic(err)
	}
}

func main() {
	message := flag.String("message", "", "")
	foo := flag.String("foo", "", "")

	flag.Parse()

	fmt.Printf("parsed message flag: %s\n", *message)
	fmt.Printf("parsed foo flag: %s\n", *foo)

	output("value", "example output")
}
