package mbgl

/*
#include <core.h>
*/
import "C"

import "unsafe"

type DefaultFileSource C.MbglDefaultFileSource

func NewDefaultFileSource(cachePath string, assetRoot string) *DefaultFileSource {
	cpath := C.CString(cachePath)
	aroot := C.CString(assetRoot)
	defer C.free(unsafe.Pointer(cpath)); C.free(unsafe.Pointer(aroot))
	
	dfs := DefaultFileSource(*C.mbgl_default_file_source_new(cpath, aroot))
	return &dfs
}

