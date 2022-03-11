package image

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"

	"github.com/fogleman/gg"
	"github.com/imroc/req"
)

// LoadImageURL load image from url
func LoadImageURL(ctx context.Context, url string) (image.Image, error) {
	if url == "" {
		return nil, nil
	}

	var img image.Image
	var err error
	var httpclient = req.New()
	resp, err := httpclient.Get(url, ctx)
	if err != nil {
		return img, err
	}

	httpResp := resp.Response()
	if httpResp == nil {
		return img, fmt.Errorf("http response is nil")
	}
	defer httpResp.Body.Close()

	imageType := httpResp.Header.Get("Content-Type")

	switch imageType {
	case "image/png":
		img, err = png.Decode(httpResp.Body)
		if err != nil {
			return img, err
		}

	case "image/jpg", "image/jpeg":
		img, err = jpeg.Decode(httpResp.Body)
		if err != nil {
			return img, err
		}

	default:
		return img, fmt.Errorf("图片格式只支持png/jpg/jpeg")
	}

	return img, nil
}

// Draw draw images
func Draw(ctx context.Context, args ...interface{}) (io.Reader, error) {
	var images []DrawParam
	var canvas Canvas
	for _, arg := range args {
		switch t := arg.(type) {
		case []DrawParam:
			images = append(images, t...)
		case DrawParam:
			images = append(images, t)
		case Canvas:
			canvas = t
		}
	}
	buf := new(bytes.Buffer)
	if len(images) == 0 {
		return buf, fmt.Errorf("no images")
	}

	// init context
	var dc *gg.Context
	if canvas.Width != 0 || canvas.Height != 0 {
		dc = gg.NewContext(canvas.Width, canvas.Height)
	} else {
		dc = gg.NewContextForImage(images[0].Image)
		images = images[1:]
	}

	// draw images
	for i := range images {
		dc.DrawImage(images[i].Image, images[i].X, images[i].Y)
	}

	err := dc.EncodePNG(buf)
	return buf, err
}
