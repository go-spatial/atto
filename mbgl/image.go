package mbgl

/*
#include <mbgl.h>
*/
import "C"

import (
	"image"
	"image/color"
	"unsafe"
)

type Image struct {
	cptr uintptr
	data []byte
	Size Size
}

func newImage(ptr uintptr) *Image {
	
	csize := C.mbgl_image_get_size(C.MbglImage(ptr))
	length := C.mbgl_image_get_bytes(C.MbglImage(ptr))
	carray := C.mbgl_image_get_data(C.MbglImage(ptr))
	
	slice := (*[1 << 30]byte)(unsafe.Pointer(carray))[:length:length]
	
	img := Image{
		cptr: ptr,
		data: slice,
		Size: Size{uint32(csize.width), uint32(csize.height)},
	}
	
	return &img
}

func (i *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i *Image) Bounds() image.Rectangle {
	size := i.Size
	return image.Rect(0,int(size.Height),0,int(size.Width))
}

/*
func (i Image) At(x, y int) color.Color {
	
	return color.RGBA()
}


func (i Image) getIndex(x, y int) int {
	
}

func (i Image) GetSize() Size {
	
}
*/

func EncodePNG(img *Image) []byte {
	var size C.size_t
	image = C.mbgl_encode_png(C.MbglPremultipliedImage(img.cptr), &size)
	
	slice := (*[1 << 30]byte)(unsafe.Pointer(&image))[:int(size):int(size)]
	
	return slice

}
