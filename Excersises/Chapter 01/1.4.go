// prints the lines of one or more files that appear more than once.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Println("Enter input:")
		countLines(os.Stdin, "os.Stdin", counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				return
			}
			countLines(f, arg, counts)
			f.Close()
		}
	}

	// Ignoring error input from input.Scan()
	fmt.Println("Now printing output:")
	for line, filenameMap := range counts {

		fmt.Printf("\"%s\" found in:\n", line)

		for filename, count := range filenameMap {
			fmt.Printf("\t%s, %d times\n", filename, count)
		}
	}
}

func countLines(f *os.File, filename string, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if counts[input.Text()] == nil {
			counts[input.Text()] = make(map[string]int)
		}
		counts[input.Text()][filename]++
	}
}
