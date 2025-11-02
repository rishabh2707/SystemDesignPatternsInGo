package main

type Computer interface {
	Print()
	SetPrinter(printer Printer)
}
