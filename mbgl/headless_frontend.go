package mbgl
/*
#include <mbgl.h>
*/
import "C"

type HeadlessFrontend C.MbglHeadlessFrontend

func NewHeadlessFrontend(size Size, pixelRatio float32, fileSource FileSource, scheduler Scheduler) *HeadlessFrontend {
	hf := HeadlessFrontend(*C.mbgl_headless_frontend_new(
		_Ctype_struct_MbglSize(size),
		_Ctype_float(pixelRatio),
		fileSource,
		scheduler))
	return &hf
}
