// Charcount computes couns of unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)    // counts of unicode characters
	var utflen [utf8.UTFMax + 1]int // counts of lengths of UTF-u encodings
	invalid := 0                    // count of invalid UTF-u characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // return runes, nbytes, error
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		counts[r]++
		utflen[n]++
	}

	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("nlen\tcount\n")
	for i, v := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, v)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalide UTF-8 characters\n", invalid)
	}
}
