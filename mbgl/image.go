package mbgl

/*
#include <mbgl.h>
*/
import "C"

import (
	"image/color"
	"image"
)

type Image struct {
	cptr uintptr
	data []byte
	Size Size
}

func newImage(ptr uintptr) *Image {
	
	csize := C.mbgl_image_get_size(*C.MbglImage(ptr))
	
	img := Image{
		cptr: ptr,
		data: C.GoBytes(C.mbgl_image_get_data(*C.MbglImage(ptr))),
		size: Size{csize.width, csize.height}
	}
	
	return &img
}

func (i *Image) ColorModel() color.Model {
	return color.RGBAModel
}

/*
func (i Image) Bounds() image.Rectangle {
	size := i.GetSize()
	return image.Rect(0,size.Height,0,size.Width)
}


func (i Image) At(x, y int) color.Color {
	
	return color.RGBA()
}

func (i Image) getIndex(x, y int) int {
	
}

func (i Image) GetSize() Size {
	
}
*/
