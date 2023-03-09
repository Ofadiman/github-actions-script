# Golang script

This repository contains a simple example of how to run Golang scripts in GitHub Actions.

## Handling inputs

Inputs can be passed to golang scripts using command-line arguments.

```yaml
go run .github/scripts/main.go -message "${{ inputs.message }}" -foo bar
```

The script must have logic that allows you to read the passed arguments. The [flag](https://pkg.go.dev/flag) package can
be used for this purpose.

```go
package main

import (
	"flag"
	"fmt"
)

func main() {
	message := flag.String("message", "", "")
	
	flag.Parse()
	
	fmt.Println(*message)
}
```

## Handling outputs

Outputs are stored in a file that is available under the path available in the environment variable `GITHUB_OUTPUT`.

```go
package main

import (
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
	output("foo", "bar")
}
```

## Resources

- [Creating GitHub actions in Go](https://jacobtomlinson.dev/posts/2019/creating-github-actions-in-go/)
- [Custom GitHub actions with Go](https://thedevelopercafe.com/articles/custom-github-action-with-go-29d9ce66e5a8)
