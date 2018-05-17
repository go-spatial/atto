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

type ConstrainMode uint32

const (
	None			ConstrainMode = C.NONE
	HeightOnly		ConstrainMode = C.HEIGHT_ONLY
	WidthAndHeight	ConstrainMode = C.WIDTH_AND_HEIGHT
)

type ViewportMode uint32

const (
	Default 	ViewportMode = C.DEFAULT
	FlippedY	ViewportMode = C.FLIPPED_Y
)

type MapDebugOptions uint32

const (
	NoDebug		MapDebugOptions = C.NO_DEBUG
	TileBorders	MapDebugOptions = C.TILE_BORDERS
	ParseStatus	MapDebugOptions = C.PARSE_STATUS
	Timestamps	MapDebugOptions = C.TIMESTAMPS
	Collision	MapDebugOptions = C.COLLISION
	Overdraw	MapDebugOptions = C.OVERDRAW
)
