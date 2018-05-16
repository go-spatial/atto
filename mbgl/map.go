package mbgl

/*
#include <mbgl.h>
*/
import "C"

type Map C.MbglMap

func NewMap(
	renderer RendererFrontend,
	size Size,
	pixelRatio float32,
	source FileSource,
	scheduler Scheduler,
	mapMode uint32,
	constrainMode uint32,
	viewportMode uint32) *Map {
	
	nmap := Map(*C.mbgl_map_new(
		renderer.(*C.MbglRendererFrontend),
		C.mbgl_map_observer_null_observer(),
		_Ctype_struct_MbglSize(size),
		_Ctype_float(pixelRatio),
		source.(*C.MbglFileSource),
		scheduler.(*C.MbglScheduler),
		_Ctype_MbglMapMode(mapMode),
		_Ctype_MbglConstrainMode(constrainMode),
		_Ctype_MbglViewportMode(viewportMode)))
	
	return &nmap
}

func (m *Map) GetStyle() Style {
	cmap := _Ctype_struct_MbglMap(*m)
	return Style(*C.mbgl_map_get_style(&cmap))
}
