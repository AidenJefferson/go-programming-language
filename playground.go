// Package tempconv performs Celsius and Fahrenheit conversions.
package main

// This program uses the package tempconv
import (
	"fmt"
	"os"
	"strconv"

	"github.com/AidenJefferson/go-programming-language/Packages/popcount"
	"github.com/AidenJefferson/go-programming-language/Packages/tempconv"
)

func main() {

	if len(os.Args) > 1 {
		fmt.Println("Converted temperature as follows:")
		for _, temp := range os.Args[1:] {
			_temp, err := strconv.ParseFloat(temp, 64)
			if err != nil {
				fmt.Printf("%v\n", err)
			}
			fmt.Printf("\t|%.02fC = %.02fF|\t|%.02fC = %.02fK|\t", _temp, tempconv.CToF(tempconv.Celsius(_temp)), _temp, tempconv.CToK(tempconv.Celsius(_temp)))
			fmt.Printf("\t|%.02fF = %.02fC|\t|%.02fF = %.02fK|\t", _temp, tempconv.FToC(tempconv.Fahrenheit(_temp)), _temp, tempconv.FToK(tempconv.Fahrenheit(_temp)))
			fmt.Printf("\t|%.02fK = %.02fC|\t|%.02fK = %.02fF|\n", _temp, tempconv.KToC(tempconv.Kelvin(_temp)), _temp, tempconv.KToF(tempconv.Kelvin(_temp)))
		}

	} else {
		fmt.Println("No temperature given.")
	}
}
