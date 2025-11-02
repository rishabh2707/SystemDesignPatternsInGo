package main

import "fmt"

type HPPrinter struct {
}

func (h *HPPrinter) PrintFile() {
	fmt.Println("Printing by HP Printer")
}
