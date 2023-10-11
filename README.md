# TiSP - A Tiny Serialization Protocol

## Installation
Install package using the following command:
```console
go get -u github.com/andriimwks/tisp
```

## Usage
This package uses `io.Reader` interface to transfer data, so you can use `net.Conn`, `bytes.Buffer` etc.
```go
buf := new(bytes.Buffer)

err := tisp.Write(buf, "hello", "world")
if err != nil {
    log.Fatal(err)
}

values, err := tisp.Read(buf)
if err != nil {
    log.Fatal(err)
}

fmt.Println(values[0].(string)) // hello
```

## Supported types
- bool
- int/8/16/32/64
- uint/8/16/32/64
- float32/64
- string
- map[string]interface{}
- []interface{}
