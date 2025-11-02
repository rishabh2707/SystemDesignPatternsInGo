package main

func main() {
	hpPrinter := &HPPrinter{}
	epsonPrinter := &EpsonPrinter{}

	macComputerWithHpPrinter := &Mac{
		printer: hpPrinter,
	}
	windowsComputerWithEpsonPrinter := &Windows{
		printer: epsonPrinter,
	}

	macComputerWithHpPrinter.Print()
	windowsComputerWithEpsonPrinter.Print()
}
