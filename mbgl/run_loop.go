package mbgl
/*
#include <mbgl-c/util/run_loop.h>
*/
import "C"

type RunLoop C.MbglRunLoop

func NewRunLoop() *RunLoop {
	rl := RunLoop(*C.mbgl_run_loop_new())
	return &rl
}
