package main

type Printer struct{}

func (Printer) print(toPrint string) string {
	return "Hello " + toPrint
}
