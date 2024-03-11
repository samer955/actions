package main

import "fmt"

var printer Printer

var something Printer

func main() {

	var toPrint = "World"
	finalString := printer.print(toPrint)
	fmt.Println(finalString)

}
