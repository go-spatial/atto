package mbgl

/*
#include <mbgl.h>
*/
import "C"

type HeadlessFrontend struct {
	cptr uintptr
}

func (h HeadlessFrontend) cPtr() uintptr {
	return h.cptr
}

func (h HeadlessFrontend) Destroy() {
	C.mbgl_headless_frontend_destroy(C.MbglHeadlessFrontend(h.cptr))
}

func (h HeadlessFrontend) Reset() {
	C.mbgl_headless_frontend_reset(C.MbglHeadlessFrontend(h.cptr))
}

func NewHeadlessFrontend(size Size, pixelRatio float32, fileSource FileSource, scheduler Scheduler) HeadlessFrontend {
	f:= C.mbgl_headless_frontend_new(
		C.MbglSize(size),
		C.float(pixelRatio),
		C.MbglFileSource(fileSource.cPtr()),
		C.MbglScheduler(scheduler.cPtr()))
	hf := HeadlessFrontend{uintptr(f)}
	return hf
}
