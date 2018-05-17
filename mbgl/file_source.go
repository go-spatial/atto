package mbgl

/*
#include <mbgl.h>
*/
import "C"

type FileSource interface {
	cPtr() uintptr	
}
