package mbgl

/*
#cgo LDFLAGS: -lstdc++
#cgo LDFLAGS: -L../deps/lib
#cgo LDFLAGS: -lmbgl-c -lmbgl-filesource -lmbgl-loop-uv -lmbgl-core
#cgo LDFLAGS: -luv -lrt -lpthread -lnsl -ldl -lsqlite3 -lcurl -lGL -lX11 -lnu -lpng16 -lz -lm -ljpeg -lwebp -licuuc -ldl
#cgo LDFLAGS: -L/usr/lib/x86_64-linux-gnu
#cgo CXXFLAGS: -std=c++14 -g -std=gnu++14
#cgo CFLAGS: -DRAPIDJSON_HAS_STDSTRING=1 -D_GLIBCXX_USE_CXX11_ABI=1 -DUCHAR_TYPE=char16_t
#cgo CFLAGS: -I../../mbgl-c/include
#cgo CFLAGS: -fPIC
#include <mbgl-c/map/map.h>
#include <mbgl-c/map/map_observer.h>
*/
import "C"

import "unsafe"

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
		unsafe.Pointer(renderer),
		C.mbgl_map_observer_null_observer(),
		_Ctype_struct_MbglSize(size),
		_Ctype_float(pixelRatio),
		unsafe.Pointer(source),
		unsafe.Pointer(scheduler),
		_Ctype_MbglMapMode(mapMode),
		_Ctype_MbglConstrainMode(constrainMode),
		_Ctype_MbglViewportMode(viewportMode)))
	
	return &nmap
}

func (m *Map) GetStyle() Style {
	cmap := _Ctype_struct_MbglMap(*m)
	return Style(*C.mbgl_map_get_style(&cmap))
}
