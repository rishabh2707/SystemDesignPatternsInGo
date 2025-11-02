package main

import "fmt"

type EpsonPrinter struct {
}

func (e *EpsonPrinter) PrintFile() {
	fmt.Println("Printing by EPSON Printer")
}
