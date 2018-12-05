# go-die

Go-DIE is a go wrapper module for [Detect-It-Easy](https://github.com/horsicq/Detect-It-Easy/).

This module can only run with runtime `GOOS=windows` and `GOARCH=386` (because of dependency to the 32-bit diedll.dll).
## Usage

### In Code

```go
import "github.com/adamyordan/go-die/die"

res, err := die.DIEScan(os.Args[1], die.DIE_SHOWERRORS)
```

### CLI
```cmd
$ cd .\go-die\bin

$ .\go-die.exe
usage: go-die.exe [file_path]

$ .\go-die.exe .\msvcp100.dll
DIE Scan successful

Result:
PE: packer: UPX
PE: compiler: Microsoft Visual C/C++
PE: linker: Microsoft Linker
```

## Example

```go
package main

import (
	"fmt"
	"github.com/adamyordan/go-die/die"
	"log"
	"os"
)

func main()  {
	if len(os.Args) < 2 {
		log.Fatalf("usage: go-die.exe [file_path]")
	}
	res, err := die.DIEScan(os.Args[1], die.DIE_SHOWERRORS)
	if err != nil {
		log.Fatalf("error running die scan: %v", err)
	}
	fmt.Printf("DIE Scan successful.\n\nResult:\n%s\n", res)
}
```