package main

import "fmt"

var printer Printer

func main() {

	var toPrint = "World"
	finalString := printer.print(toPrint)
	fmt.Println(finalString)

}
