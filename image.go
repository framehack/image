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
func Draw(ctx context.Context, images []DrawParam) (io.Reader, error) {
	buf := new(bytes.Buffer)
	if len(images) == 0 {
		return buf, fmt.Errorf("no images")
	}
	dc := gg.NewContextForImage(images[0].Image)
	for i := range images[1:] {
		dc.DrawImage(images[i].Image, images[i].X, images[i].Y)
	}

	err := dc.EncodePNG(buf)
	return buf, err
}
