package main

import "testing"

const expected = "Hello Testing"

func TestPrinter_printing(t *testing.T) {
	printer := Printer{}

	t.Run("Testing print", func(t *testing.T) {
		toPrint := "Testing"
		result := printer.print(toPrint)

		if result != expected {
			t.Errorf("Expected %s but got %s", expected, result)
		}
	})
}
