package main

import (
	//"os"
	"runtime"
	
	//"github.com/jung-kurt/gofpdf"

	"github.com/go-spatial/atto/mbgl"
)

func main() {
    
    //debug := true
    
    var pixelRatio float32 = 5
    
    loop := mbgl.NewRunLoop()
    defer loop.Destroy()
    
	//token := os.Getenv("MAPBOX_ACCESS_TOKEN")
	
    //fileSource := mbgl.NewDefaultFileSource("mbgl/testdata/cache.sqlite", ".")
    //fileSource.SetAccessToken(token)
    fileSource := mbgl.NewOnlineFileSource()
    fileSource.SetAPIBaseUrl("https://osm.tegola.io/")
    defer fileSource.Destroy()

    threadPool := mbgl.NewThreadPool(runtime.NumCPU())
    defer threadPool.Destroy()
    
    size := mbgl.Size{ 1024, 1024 }
    
    frontEnd := mbgl.NewHeadlessFrontend(size, pixelRatio, fileSource, threadPool)
    defer frontEnd.Destroy()
    
    pmap := mbgl.NewMap(frontEnd, size, pixelRatio, fileSource, threadPool, mbgl.Static, mbgl.HeightOnly, mbgl.Default)
    defer pmap.Destroy()
    
    pmap.GetStyle().LoadURL("https://raw.githubusercontent.com/go-spatial/tegola-web-demo/master/styles/camo.json")
    
    latLng := mbgl.NewLatLng(39.153,-76.275)
    defer latLng.Destroy()
    
    var zoom, bearing, pitch float32 = 8, 0, 0
    
    pmap.SetLatLngZoom(latLng, zoom)
    pmap.SetBearing(bearing)
    pmap.SetPitch(pitch)
    
	//if debug {
	//	pmap.SetDebug()
    //}
    
    frontEnd.RenderToFile(pmap, "out.png")
    
    //image := EncodePNG(frontEnd.Render(pmap))
}
