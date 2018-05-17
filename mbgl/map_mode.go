package mbgl

/*
#include <mbgl.h>
*/
import "C"

type MapMode uint32

const (
	Continuous 	MapMode = C.CONTINUOUS
	Static		MapMode = C.STATIC
	Tile		MapMode = C.TILE
)

