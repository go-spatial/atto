package mbgl
/*
#include <mbgl-c/gl/headless_frontend.h>
#include <mbgl-c/util/size.h>
*/
import "C"

import "unsafe"

type HeadlessFrontend C.MbglHeadlessFrontend

func NewHeadlessFrontend(size Size, pixelRatio float32, fileSource FileSource, scheduler Scheduler) *HeadlessFrontend {
	hf := HeadlessFrontend(*C.mbgl_headless_frontend_new(
		_Ctype_struct_MbglSize(size),
		_Ctype_float(pixelRatio),
		unsafe.Pointer(&fileSource),
		unsafe.Pointer(&scheduler)))
	return &hf
}
