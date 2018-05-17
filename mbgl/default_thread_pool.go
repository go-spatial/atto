package mbgl

/*
#include <mbgl.h>
*/
import "C"

type ThreadPool struct {
	cptr uintptr
}

func (tp ThreadPool) cPtr() uintptr {
	return tp.cptr
}

func NewThreadPool(count int) ThreadPool {
	return ThreadPool{ uintptr(C.mbgl_thread_pool_new(C.size_t(count))) };
}
