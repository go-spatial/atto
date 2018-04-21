package map_mode

/*
#include <mbgl-c/map/mode.h>
*/
import "C"

const (
	Continuous 	uint32 = C.CONTINUOUS
	Static		uint32 = C.STATIC
	Tile		uint32 = C.TILE
)
