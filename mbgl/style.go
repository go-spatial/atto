package mbgl

/*
#include <mbgl.h>
*/
import "C"

import "unsafe"

type Style struct {
	cptr uintptr
}

func (s *Style) cPtr() uintptr {
	return s.cptr
}

func (s *Style) Destroy() {
	C.mbgl_style_destroy(C.MbglStyle(s.cptr))
}

func (s Style) LoadURL(url string) {
	curl := C.CString(url)
	defer C.free(unsafe.Pointer(curl))

	C.mbgl_style_load_url(C.MbglStyle(s.cptr), curl)
}
