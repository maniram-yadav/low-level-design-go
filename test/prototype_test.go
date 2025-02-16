package test

import (
	"fmt"
	"lld/prototype"
	"testing"
)

func TestPrototype(t *testing.T) {
	file1 := &prototype.File{Name: "file1"}
	file2 := &prototype.File{Name: "file2"}
	file3 := &prototype.File{Name: "file3"}
	folder1 := &prototype.Folder{
		Children: []prototype.INode{file1},
		Name:     "folder1"}
	folder2 := &prototype.Folder{
		Children: []prototype.INode{folder1, file2, file3},
		Name:     "folder1"}
	fmt.Printf("\nPrint heirarchy for Folder %s\n", folder2.Name)
	folder2.Print("    ")
	clonedFolder := folder2.Clone()
	name := clonedFolder.(*prototype.Folder).Name
	fmt.Printf("\nPrining heirarchy for Folder %s\n", name)
	clonedFolder.Print("    ")

}
