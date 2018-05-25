package main

import (
	"crypto/sha512"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var width = flag.Int("w", 384, "hash width (384 or 512)")

func main() {
	flag.Parse()
	var function func(b []byte) []byte
	switch *width {
	case 384:
		function = func(b []byte) []byte {
			h := sha512.Sum384(b)
			return h[:]
		}
	case 512:
		function = func(b []byte) []byte {
			h := sha512.Sum512(b)
			return h[:]
		}
	default:
		log.Fatal("Unexpected width specified.")
	}

	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%x\n", function(b))
}
