package image

import "image"

// DrawParam is the param for Draw
type DrawParam struct {
	Image image.Image // image to Draw
	X     int         // x position
	Y     int         // y position
}
