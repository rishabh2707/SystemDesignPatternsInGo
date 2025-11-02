package main

type Windows struct {
	printer Printer
}

func (w *Windows) Print() {
	w.printer.PrintFile()
}

func (w *Windows) SetPrinter(printer Printer) {
	w.printer = printer
}
