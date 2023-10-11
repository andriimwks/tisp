# TiSP - A Tiny Serialization Protocol

## Installation
Install package using the following command:
```console
go get -u github.com/andriimwks/tisp
```

## Usage
This package uses `io.Reader` interface to transfer data, so you can use `net.Conn`, `bytes.Buffer` etc.
```go
package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/andriimwks/tisp"
)

func main() {
	buf := new(bytes.Buffer)

	err := tisp.Write(buf, "Hello", "World")
	if err != nil {
		log.Fatal(err)
	}

	values, err := tisp.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s, %s!\n", values[0], values[1]) // Hello, World!
}
```

## Supported types
- bool
- int/8/16/32/64
- uint/8/16/32/64
- float32/64
- string
- map[string]interface{}
- []interface{}
