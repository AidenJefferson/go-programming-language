// prints the lines of a file that appear more than once.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				return
			}
			countLines(f, counts)
			f.Close()
		}
	}

	// Ignoring error input from input.Scan()
	fmt.Println("Now printing output:")
	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%s. Appears %d\n", line, count)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
