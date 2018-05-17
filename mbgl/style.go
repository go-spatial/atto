package mbgl

/*
#include <mbgl.h>
*/
import "C"

type Style struct {
	cptr uintptr
}

func (s Style) cPtr() uintptr {
	return s.cptr
}
