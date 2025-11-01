package object

import "fmt"

type Folder struct {
	Name     string
	Children []Inode
}

func (f *Folder) GetDetails() {
	fmt.Printf("Folder: %s\n", f.Name)
	fmt.Println("Children:")
	for _, child := range f.Children {
		fmt.Print("  - ")
		child.GetDetails()
	}
}

func (f *Folder) Add(inode Inode) {
	f.Children = append(f.Children, inode)
}
