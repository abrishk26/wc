package main

import (
	"bufio"
	"io"
	"fmt"
	"os"
	"flag"
)

func main() {

	lines := flag.Bool("l", false, "Count line")

	// parsing flags provided by the user
	flag.Parse()

	// Calling the count function to count the number of words
	// received from the Standard Input and printing it out.
	fmt.Println(count(os.Stdin, *lines))
}

func count(r io.Reader, countLines bool) int {
	// A scanner is used to read a text from a Reader (such as file).
	scanner := bufio.NewScanner(r)

	if !countLines {
		// define the scanner split type to words. (default is split by lines)
		scanner.Split(bufio.ScanWords)
	}

	// Defining word counter
	wc := 0

	// for every word scanned, increment the counter
	for scanner.Scan() {
		wc++
	}

	return wc
} 