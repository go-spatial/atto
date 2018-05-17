package main

import (
	"os"
	//"github.com/jung-kurt/gofpdf"

	"github.com/go-spatial/atto/mbgl"
)

func main() {
    
    //debug := true
    
    var pixelRatio float32 = 1
    
    loop := mbgl.NewRunLoop()
    defer loop.Destroy()
    
	token := os.Getenv("MAPBOX_ACCESS_TOKEN")
	
    fileSource := mbgl.NewDefaultFileSource("mbgl/testdata/cache.sqlite", ".")
    fileSource.SetAccessToken(token)
    defer fileSource.Destroy()

    threadPool := mbgl.NewThreadPool(4)
    defer threadPool.Destroy()
    
    size := mbgl.Size{ 512, 512 }
    
    frontEnd := mbgl.NewHeadlessFrontend(size, pixelRatio, fileSource, threadPool)
    defer frontEnd.Destroy()
    
    pmap := mbgl.NewMap(frontEnd, size, pixelRatio, fileSource, threadPool, mbgl.Static, mbgl.HeightOnly, mbgl.Default)
    defer pmap.Destroy()
    
    pmap.GetStyle().LoadURL("mapbox://styles/mapbox/satellite-v9")
    
    latLng := mbgl.NewLatLng(0,0)
    defer latLng.Destroy()
    
    var zoom, bearing, pitch float32 = 0, 0, 0
    
    pmap.SetLatLngZoom(latLng, zoom)
    pmap.SetBearing(bearing)
    pmap.SetPitch(pitch)
    
	//if debug {
	//	_map.SetDebug()
    //}
    
    frontEnd.RenderToFile(pmap, "out.png")
    
    //image := EncodePNG(frontEnd.Render(pmap))
}
