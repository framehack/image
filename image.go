package image

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/davidbyttow/govips/v2/vips"
)

// Service vips service
type Service struct {
}

// NewService create vips service
func NewService() *Service {
	vips.Startup(nil)
	return &Service{}
}

// Shutdown shutdown vips
func (s *Service) Shutdown() {
	vips.Shutdown()
}

// DrawWhiteCanvas draw white canvas
func (s *Service) DrawWhiteCanvas(width, height int) (*vips.ImageRef, error) {
	bg, err := vips.Black(1, 1)
	if err != nil {
		return bg, err
	}
	bg.DrawRect(vips.ColorRGBA{
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	}, 0, 0, 1, 1, true)

	err = bg.Embed(0, 0, width, height, vips.ExtendWhite)
	return bg, err
}

// Draw draw images
func (s *Service) Draw(ctx context.Context, args ...interface{}) (io.Reader, error) {
	var images []DrawParam
	var canvas Canvas
	var format = "jpeg"
	for _, arg := range args {
		switch t := arg.(type) {
		case []DrawParam:
			images = append(images, t...)
		case DrawParam:
			images = append(images, t)
		case Canvas:
			canvas = t
		case OutputPNG:
			format = "png"
		}
	}
	buf := new(bytes.Buffer)
	if len(images) == 0 {
		return buf, fmt.Errorf("no images")
	}
	var bg *vips.ImageRef
	var err error
	if canvas.Width != 0 && canvas.Height != 0 {
		bg, err = s.DrawWhiteCanvas(canvas.Width, canvas.Height)
		if err != nil {
			return buf, err
		}
	} else {
		bg = (images[0].Image)
		images = images[1:]
	}

	// draw images
	for i := range images {
		bg.Composite(images[i].Image, vips.BlendModeOver, images[i].X, images[i].Y)
	}
	var ep *vips.ExportParams
	switch format {
	case "png":
		ep = vips.NewDefaultPNGExportParams()
	default:
		ep = vips.NewDefaultJPEGExportParams()
	}

	bytes, _, err := bg.Export(ep)
	if err != nil {
		return buf, err
	}
	_, err = buf.Write(bytes)
	return buf, err
}
