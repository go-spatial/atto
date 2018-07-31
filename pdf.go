package main

import (
	"bytes"

	"github.com/jung-kurt/gofpdf"
)

const (
	margin   = 10
	wd       = 210
	ht       = 297
	fontSize = 15
)


func pdf(image []byte, filename string) {
	
	msgStr := "Hello World"
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Helvetica", "", fontSize)
	ln := pdf.PointConvert(fontSize)
	pdf.MultiCell(wd-margin-margin, ln, msgStr, "", "L", false)

	infoPtr := pdf.RegisterImageOptionsReader("map", gofpdf.ImageOptions{ "PNG", true, false }, bytes.NewReader(image))
	
	if pdf.Ok() {
		imgWd, imgHt := infoPtr.Extent()
		pdf.Image("map", (wd-imgWd)/2.0, pdf.GetY()+ln,
			imgWd, imgHt, false, "image/png", 0, "")
	}
	
	pdf.OutputFileAndClose(filename)
	
}
