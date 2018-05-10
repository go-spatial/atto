package mbgl
/*
#include <core.h>
*/
import "C"

type RunLoop C.MbglRunLoop

func NewRunLoop() *RunLoop {
	rl := RunLoop(*C.mbgl_run_loop_new())
	return &rl
}
