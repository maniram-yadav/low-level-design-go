package momento

type Image struct {
	Data []byte
}

type Filter interface {
	Apply(image *Image) *Image
}

type ImageEditor struct {
	Image   *Image
	history []*Image
}

func NewImageEditor(image *Image) *ImageEditor {
	imageEditor := &ImageEditor{Image: image, history: make([]*Image, 0)}
	return imageEditor
}

func (ie *ImageEditor) ApplyFilter(filter Filter) {
	ie.history = append(ie.history, ie.Image)
	ie.Image = filter.Apply(ie.Image)
}

func (ie *ImageEditor) Undo() {
	if len(ie.history) > 0 {
		ie.Image = ie.history[len(ie.history)-1]
		ie.history = ie.history[:len(ie.history)-1]
	}
}

type GrayFilter struct{}

func (gf GrayFilter) Apply(image *Image) *Image {

	newImageData := make([]byte, len(image.Data))
	for i := range image.Data {
		newImageData[i] = image.Data[i] / 3
	}
	return &Image{Data: newImageData}

}
