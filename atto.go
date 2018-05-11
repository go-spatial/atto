package main

import (
	"github.com/jung-kurt/gofpdf"

	"github.com/go-spatial/atto/mbgl"
)

func main() {
    
    debug := true
    
    fileSource := NewOnlineFileSource()
    threadPool := NewThreadPool(4)
    frontEnd := NewHeadLessFrontend({ width, height }, pixelRatio, fileSource, threadPool)
    pmap := NewMap(frontEnd, frontEnd.GetSize(), pixelRatio, fileSource, threadPool, map_mode.Static)
    
    defer pmap.Destroy()
    
    pmap.GetStyle().LoadURL(style)
    pmap.SetLatLngZoom({ lat, lon }, zoom)
    pmap.SetBearing(bearing)
    pmap.SetPitch(pitch)
    
	if debug {
		_map.SetDebug()
    }
    
    image := EncodePNG(frontEnd.Render(_map))
}
