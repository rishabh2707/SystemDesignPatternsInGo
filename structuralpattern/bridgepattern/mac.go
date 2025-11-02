package main

type Mac struct {
	printer Printer
}

func (m *Mac) Print() {
	m.printer.PrintFile()
}

func (m *Mac) SetPrinter(printer Printer) {
	m.printer = printer
}
