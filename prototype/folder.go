package prototype

import "fmt"

type Folder struct {
	Children []INode
	Name     string
}

func (folder *Folder) Print(ident string) {
	fmt.Println(ident + folder.Name)
	for _, child := range folder.Children {
		child.Print(ident + ident)
	}
}

func (folder *Folder) Clone() INode {

	var tempChild []INode
	folder_clone := &Folder{}
	folder_clone.Name = folder.Name + "_clone"
	for _, ch := range folder.Children {
		copy := ch.Clone()
		tempChild = append(tempChild, copy)
	}

	folder_clone.Children = tempChild
	return folder_clone
}
