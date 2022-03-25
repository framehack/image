package image

import (
	"context"
	"fmt"

	"github.com/davidbyttow/govips/v2/vips"
	"github.com/imroc/req"
)

// LoadImageURL load image from url
func LoadImageURL(ctx context.Context, url string) (*vips.ImageRef, error) {
	if url == "" {
		return nil, nil
	}

	var httpclient = req.New()
	resp, err := httpclient.Get(url, ctx)
	if err != nil {
		return nil, err
	}

	httpResp := resp.Response()
	if httpResp == nil {
		return nil, fmt.Errorf("http response is nil")
	}
	defer httpResp.Body.Close()
	
	return vips.NewImageFromReader(httpResp.Body)
}