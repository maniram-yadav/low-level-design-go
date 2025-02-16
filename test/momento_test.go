package test

import (
	"lld/momento"
	"testing"
)

func TestMomento(t *testing.T) {

	image := &momento.Image{Data: []byte{100, 150, 200, 50, 75, 100}} // Sample image data
	imageEditor := momento.NewImageEditor(image)

	grayFilter := momento.GrayFilter{}
	imageEditor.ApplyFilter(grayFilter)
	t.Log("Image Data after filter:", imageEditor.Image.Data)
	imageEditor.Undo()
	t.Log("Image Data after undo:", imageEditor.Image.Data)

}
