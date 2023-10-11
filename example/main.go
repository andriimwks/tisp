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
