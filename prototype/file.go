package prototype

import "fmt"

type File struct {
	Name string
}

func (file *File) Print(ident string) {
	fmt.Println(ident + file.Name)
}

func (file *File) Clone() INode {
	return &File{Name: file.Name + "_clone"}
}
