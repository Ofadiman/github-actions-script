# Golang script

This repository contains examples of how to run [Golang](https://go.dev/) and [Deno](https://deno.land/) scripts in
GitHub Actions.

## Handling inputs in Golang

Inputs can be passed to scripts using command-line arguments.

```yaml
go run .github/scripts/main.go -message "${{ inputs.message }}"
```

The [flag](https://pkg.go.dev/flag) package can be used for parsing CLI arguments.

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

## Handling outputs in Golang

Outputs are stored in a file which name is available at `GITHUB_OUTPUT` environment variable. In order to store an
output, the script code has to append key/value pairs to the aforementioned file.

```go
package main

import (
	"fmt"
	"os"
)

func main() {
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
```

## Handling inputs in Deno

Inputs can be passed to scripts using command-line arguments.

```yaml
deno run --allow-env --allow-read --allow-write .github/scripts/deno.ts --message "${{ inputs.message }}"
```

The [flags](https://deno.land/std@0.178.0/flags/mod.ts) module can be used for parsing CLI arguments.

```typescript
import {parse} from 'https://deno.land/std@0.178.0/flags/mod.ts'

const parsedArgs = parse(Deno.args, {
  string: ['message'],
})
console.log(`parsed message flag: ${parsedArgs.message}`)
```

## Handling outputs in Deno

Outputs are stored in a file which name is available at `GITHUB_OUTPUT` environment variable. In order to store an
output, the script code has to append key/value pairs to the aforementioned file.

```typescript
const githubOutput = Deno.env.get('GITHUB_OUTPUT') as string
await Deno.writeFile(githubOutput, new TextEncoder().encode(`value=example\n`), {append: true})
```

## Resources

- [Command line arguments in Deno](https://examples.deno.land/command-line-arguments)
- [Creating GitHub actions in Go](https://jacobtomlinson.dev/posts/2019/creating-github-actions-in-go/)
- [Custom GitHub actions with Go](https://thedevelopercafe.com/articles/custom-github-action-with-go-29d9ce66e5a8)
- [Permissions in Deno](https://deno.land/manual@v1.31.1/basics/permissions)
- [Writing files in Deno](https://medium.com/deno-the-complete-reference/writing-to-files-in-deno-6b3317ccd946)
