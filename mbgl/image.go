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
	
	img := Image{
		cptr: ptr,
		data: C.GoBytes(unsafe.Pointer(carray), C.int(length)),
		Size: Size{uint32(csize.width), uint32(csize.height)},
	}
	
	return &img
}

func (i *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i *Image) Bounds() image.Rectangle {
	return image.Rect(0,0,int(i.Size.Width),int(i.Size.Height))
}

func (i *Image) At(x, y int) color.Color {
	
	stride := i.Size.Width * 4
	
	row := y * int(stride)
	col := x * 4
	idx := row + col
	
	return color.RGBA{i.data[idx], i.data[idx+1], i.data[idx+2], i.data[idx+3]}
}

func (i *Image) Destroy() {
	C.mbgl_image_destroy(C.MbglImage(i.cptr))
}

/*
func EncodePNG(img *Image) []byte {
	var size C.size_t
	png := C.mbgl_encode_png(C.MbglPremultipliedImage(img.cptr), &size)
	
	slice := (*[1 << 30]byte)(unsafe.Pointer(&img))[:int(size):int(size)]
	
	return slice

}*/
