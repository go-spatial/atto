package mbgl

/*
#include <mbgl.h>
*/
import "C"

type RendererFrontend interface {
	cPtr() uintptr
	Reset()
	Destroy()
}

