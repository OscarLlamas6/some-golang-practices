package main

import (
	"fmt"
)

type (
	component interface {
		search(string)
	}

	file struct {
		name string
	}

	folder struct {
		components []component
		name       string
	}
)

// File methods
func (f *file) search(keyword string) {
	fmt.Printf("We are looinkg for the keyword %s in the file %s\n", keyword, f.name)
}

func (f *file) getName() string {
	return f.name
}

// Folder methods
func (f *folder) search(keyword string) {
	fmt.Printf("Searching recursively for the keyword %s in the folder %s\n", keyword, f.name)
	for _, composite := range f.components {
		composite.search(keyword)
	}
}

func (f *folder) add(c component) {
	f.components = append(f.components, c)
}

//main

func main() {
	file1 := &file{name: "file1"}
	file2 := &file{name: "file2"}
	file3 := &file{name: "file3"}

	folder1 := &folder{name: "folder1"}
	folder1.add(file1)

	folder2 := &folder{name: "folder2"}
	folder2.add(file2)
	folder2.add(file3)

	folder2.add(folder1)

	folder2.search("jossvan")
}
