package main

import "com.systemdesignpattern.go.compositepattern/object"

func main() {
	folder := &object.Folder{
		Name:     "root",
		Children: make([]object.Inode, 0),
	}

	movies := &object.File{
		Name: "KGF2",
		Size: 100,
	}

	folder.Add(movies)

	childFolder := &object.Folder{
		Name:     "HorrorMovies",
		Children: make([]object.Inode, 0),
	}

	conjuringMovie := &object.File{
		Name: "Conjuring",
		Size: 200,
	}

	childFolder.Add(conjuringMovie)

	folder.Add(childFolder)

	folder.GetDetails()

}
