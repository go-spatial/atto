package mbgl

/*
#include <mbgl.h>
*/
import "C"

type Scheduler interface {
	cPtr() uintptr
}

