package image

import (
	"github.com/davidbyttow/govips/v2/vips"
)

// DrawParam is the param for Draw
type DrawParam struct {
	Image *vips.ImageRef // image to Draw
	X     int         // x position
	Y     int         // y position
}

// Canvas is the canvas for Draw
type Canvas struct {
	Width  int
	Height int
}

// OutputJPEG is the output format for Draw, default format
type OutputJPEG struct {}
// OutputPNG output png format
type OutputPNG struct {}
