package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	lines := flag.Bool("l", false, "Count lines")
	byteCount := flag.Bool("b", false, "Count bytes")
	filename := flag.String("file", "", "File to count words from")
	// Parsing the flags provided by the user
	flag.Parse()

	// Calling the count function to return the number of words
	// read from STDIN
	fmt.Println(count(*filename, os.Stdin, *lines, *byteCount))
}

func count(filename string, r io.Reader, countlines bool, countBytes bool) int {

	// io.Reader      (p []byte)     (n int, err error)
	// ioutil.ReadFile(fname string) ([]byte, error)

	// if reading from a file
	if filename != "" {

	}

	input, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Errorf("Error reading %s", filename)
		os.Exit(1)
	}

	s := bufio.NewScanner(input)

	scanner := bufio.NewScanner(r)

	if countlines {
		scanner.Split(bufio.ScanLines)
	} else if countBytes {
		scanner.Split(bufio.ScanBytes)
	} else {
		scanner.Split(bufio.ScanWords)
	}

	wc := 0

	for scanner.Scan() {
		wc++
	}

	return wc
}

type fileReader []byte

func (f *fileReader) Read(p []byte) (n int, err error) {

}
