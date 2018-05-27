package main

import (
	"bufio"
	"fmt"
	"os"
)

func wordfreq() map[string]int {
	count := make(map[string]int) //Golang could declare here and use outside of function
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		count[word]++
	}
	if scanner.Err() != nil {
		fmt.Fprintln(os.Stderr, scanner.Err())
	}

	return count
}

func main() {
	freq := wordfreq()
	for word, n := range freq {
		fmt.Printf("%-30s %d\n", word, n)
	}
}
