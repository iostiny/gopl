package main

import (
	"fmt"
	"os"
)

func main() {
	for i, v := range os.Args[1:] {
		fmt.Printf("index is :%d, v is :%s\n", i, v)
	}
}
