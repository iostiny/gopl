// ex1.4 prints the count and text of lines that appear more than once in the
// input. It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	foundin := make(map[string][]string) //key is `lines that appear more than once`, value is a filename array
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, foundin)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, foundin)
			f.Close()
		}
	}

	for lines, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, foundin[lines], lines)
		}
	}
}

func in(needle string, strings []string) bool {
	for _, s := range strings {
		if needle == s {
			return true
		}
	}
	return false
}

func countLines(f *os.File, counts map[string]int, foundin map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		if !in(f.Name(), foundin[line]) {
			foundin[line] = append(foundin[line], f.Name())
		}
	}
	//Note: ignoring potential errors from input.Err()
}
