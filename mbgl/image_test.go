package mbgl

import (
	"os"
	"testing"
	"image/png"
)

func TestImage(t *testing.T) {

var pixelRatio float32 = 1
    
    loop := NewRunLoop()
    defer loop.Destroy()
    
	token := os.Getenv("MAPBOX_ACCESS_TOKEN")
	
    fileSource := NewDefaultFileSource("testdata/cache.sqlite", ".")
    fileSource.SetAccessToken(token)
    defer fileSource.Destroy()

    threadPool := NewThreadPool(4)
    defer threadPool.Destroy()
    
    size := Size{ 512, 512 }
    
    frontEnd := NewHeadlessFrontend(size, pixelRatio, fileSource, threadPool)
    defer frontEnd.Destroy()
    
    pmap := NewMap(frontEnd, size, pixelRatio, fileSource, threadPool, Static, HeightOnly, Default)
    defer pmap.Destroy()
    
    pmap.GetStyle().LoadURL("mapbox://styles/mapbox/satellite-v9")
    
    latLng := NewLatLng(0,0)
    defer latLng.Destroy()
    
    var zoom, bearing, pitch float32 = 0, 0, 0
    
    pmap.SetLatLngZoom(latLng, zoom)
    pmap.SetBearing(bearing)
    pmap.SetPitch(pitch)
    
    //frontEnd.RenderToFile(pmap, "out.png")
    
	image := frontEnd.Render(pmap)
	defer image.Destroy()

	f, err := os.Create("out.png")
	if err != nil {
		t.Fatal(err)
	}

	if err := png.Encode(f, image); err != nil {
		f.Close()
		t.Fatal(err)
	}

	if err := f.Close(); err != nil {
		t.Fatal(err)
	}
	
}
