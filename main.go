package main

import (
	"bufio" // to read text
	"flag"  // for managing cmd options
	"fmt"
	"io"
	"os"
)

func count(r io.Reader, countLines bool) int {
	scanner := bufio.NewScanner(r) // read text

	if !countLines {
		scanner.Split(bufio.ScanWords) // split to words
	}

	wc := 0              // defining the counter
	for scanner.Scan() { // increment wc with every word
		wc++
	}
	return wc
}

func main() {

	lines := flag.Bool("l", false, "Count lines")
	flag.Parse()
	fmt.Println(count(os.Stdin, *lines))
}
