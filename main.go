package main

import (
	"bufio"
	"io"
	"fmt"
	"os"
)

func main() {
	// Calling the count function to count the number of words
	// received from the Standard Input and printing it out.
	fmt.Println(count(os.Stdin))
}

func count(r io.Reader) int {
	// A scanner is used to read a text from a Reader (such as file).
	scanner := bufio.NewScanner(r)

	// define the scanner split type to words. (default is split by lines)
	scanner.Split(bufio.ScanWords)

	// Defining word counter
	wc := 0

	// for every word scanned, increment the counter
	for scanner.Scan() {
		wc++
	}

	return wc
} 