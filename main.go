package main

import "fmt"

var printer Printer
var extra Printer

func main() {

	var toPrint = "World"
	finalString := printer.print(toPrint)
	fmt.Println(finalString)

}
