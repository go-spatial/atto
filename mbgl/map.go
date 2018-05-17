package mbgl

/*
#include <mbgl.h>
*/
import "C"

type Map struct {
	cptr uintptr
}

func NewMap(
	renderer RendererFrontend,
	size Size,
	pixelRatio float32,
	source FileSource,
	scheduler Scheduler,
	mapMode MapMode,
	constrainMode ConstrainMode,
	viewportMode ViewportMode) Map {
	
	nmap := C.mbgl_map_new(
		C.MbglRendererFrontend(renderer.cPtr()),
		C.mbgl_map_observer_null_observer(),
		C.MbglSize(size),
		C.float(pixelRatio),
		C.MbglFileSource(source.cPtr()),
		C.MbglScheduler(scheduler.cPtr()),
		C.MbglMapMode(mapMode),
		C.MbglConstrainMode(constrainMode),
		C.MbglViewportMode(viewportMode))
	
	return Map{ uintptr(nmap) }
}

func (m Map) Destroy() {
	C.mbgl_map_destroy(C.MbglMap(m.cptr))
}

func (m Map) GetStyle() Style {
	return Style{ uintptr(C.mbgl_map_get_style(C.MbglMap(m.cptr))) }
}
