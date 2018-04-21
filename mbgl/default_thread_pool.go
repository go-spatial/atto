package mbgl
/*
#include <mbgl-c/util/default_thread_pool.h>
*/
import "C"

type ThreadPool C.MbglThreadPool

func NewThreadPool(count int) *ThreadPool {
	
	tp := ThreadPool(*C.mbgl_thread_pool_new(C.size_t(count)));

	return &tp
}
