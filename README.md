# Vembed

Vembed simplifies embedding a version into a binary.

## Example

This example show how to use vembed.

```
package main

import (
	"fmt"

	"github.com/NoUseFreak/go-vembed"
)

func main() {
	fmt.Printf(
		"%s, build %s",
		vembed.Version.GetGitSummary(),
		vembed.Version.GetGitCommit(),
	)
}
```

## Build

```
go get github.com/NoUseFreak/go-vembed/vembed
go build -ldflags="`vembed`" main.go
```
