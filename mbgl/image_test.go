package mbgl

import (
	"testing"
	"bytes"
	"image/png"
)

func TestImage(t *testing.T) {

var pixelRatio float32 = 1
    
    loop := NewRunLoop()
    defer loop.Destroy()
	
    fileSource := NewDefaultFileSource("testdata/cache.sqlite", ".")
    defer fileSource.Destroy()

    threadPool := NewThreadPool(4)
    defer threadPool.Destroy()
    
    size := Size{ 512, 512 }
    
    frontEnd := NewHeadlessFrontend(size, pixelRatio, fileSource, threadPool)
    defer frontEnd.Destroy()
    
    pmap := NewMap(frontEnd, size, pixelRatio, fileSource, threadPool, Static, HeightOnly, Default)
    defer pmap.Destroy()
    
    pmap.GetStyle().LoadURL("https://osm.tegola.io/maps/osm/style.json")
    
    latLng := NewLatLng(0,0)
    defer latLng.Destroy()
    
    var zoom, bearing, pitch float32 = 0, 0, 0
    
    pmap.SetLatLngZoom(latLng, zoom)
    pmap.SetBearing(bearing)
    pmap.SetPitch(pitch)
    
    //frontEnd.RenderToFile(pmap, "out.png")
    
	image := frontEnd.Render(pmap)
	defer image.Destroy()

	var b bytes.Buffer

	if err := png.Encode(&b, image); err != nil {
		t.Fatal(err)
	}
	
	if b.Len() == 0 {
		t.Fatal("Image is 0 size")
	}
	
}
