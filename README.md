# Konmai LZ77 Compressor & Decompressor
Compressor & Decompressor for various konmai resources such as files and network messages.

## Install
With a [correctly configured](https://go.dev/doc/install#testing) Go toolchain:
```shell
go get -u github.com/SirusDoma/klz77
```

## Usage
```go
package main

import(
    "io/ioutil"
	
    "github.com/SirusDoma/klz77"
)

func main() {
    // Read input file
    input, _ := ioutil.ReadFile("/some/path/to/file")

    // Write compressed data into a file
    compressed, _ := klz77.Compress(input)
    _ = ioutil.WriteFile("/some/path/to/compressed", compressed, 0644)

    // Re-decompress compressed data and write it into a file
    decompressed, _ := klz77.Decompress(compressed)
    _ = ioutil.WriteFile("/some/path/to/decompressed", decompressed, 0644)
}
```

## License
This is an open-sourced program licensed under the [MIT License](http://github.com/SirusDoma/klz77/blob/master/LICENSE).