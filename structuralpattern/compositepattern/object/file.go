package object

import "fmt"

type File struct {
	Name string
	Size int
}

func (f *File) GetDetails() {
	fmt.Printf("File: %s, Size: %d\n", f.Name, f.Size)
}
