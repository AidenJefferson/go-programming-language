// This program prints the content found at a URL (given as an argument).
// Before doing so it will check for the prefix "http://" and insert if missing.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	prefix := "http://"
	for _, url := range os.Args[1:] {
		if strings.HasPrefix(url, prefix) == false {
			url = prefix + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		noBytes, err := io.Copy(os.Stdout, resp.Body)
		fmt.Printf("Coppied %d bytes from http\n", noBytes)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
