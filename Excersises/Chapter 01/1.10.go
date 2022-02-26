// This program fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	output, err := os.Create("1.10_output.txt")
	if err != nil {
		fmt.Printf("%v:\n", err) // receive from channel ch
	}
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	output.WriteString("Time:\tBytes:\tURL:\n")
	for range os.Args[1:] {
		output.WriteString(<-ch)
	}
	output.WriteString(fmt.Sprintf("\n%fs total elapsed\n", time.Since(start).Seconds()))
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v\n", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs\t%7d\t%s\n", secs, nbytes, url)
}
