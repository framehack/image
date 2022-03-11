package image

import (
	"context"
	"io/ioutil"
	"testing"
)

func TestDraw(t *testing.T) {
	ctx := context.Background()
	img1, err := LoadImageURL(ctx, "http://examples-1251000004.cos.ap-shanghai.myqcloud.com/sample.jpeg")
	if err != nil {
		t.Fatal(err)
	}
	img2, err := LoadImageURL(ctx, "https://image-demo.oss-cn-hangzhou.aliyuncs.com/example.jpg")
	if err != nil {
		t.Fatal(err)
	}
	images := make([]DrawParam, 2)
	images[0] = DrawParam{
		Image: img1,
		X:     0,
		Y:     0,
	}
	images[1] = DrawParam{
		Image: img2,
		X:     0,
		Y:     0,
	}
	reader, err := Draw(ctx, images)
	if err != nil {
		t.Fatal(err)
	}
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		t.Fatal(err)
	}
	err = ioutil.WriteFile("test.png", b, 0644)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDrawWithCanvas(t *testing.T) {
	ctx := context.Background()
	canvas := Canvas{
		Width:  100,
		Height: 100,
	}
	img1, err := LoadImageURL(ctx, "http://examples-1251000004.cos.ap-shanghai.myqcloud.com/sample.jpeg")
	if err != nil {
		t.Fatal(err)
	}
	img2, err := LoadImageURL(ctx, "https://image-demo.oss-cn-hangzhou.aliyuncs.com/example.jpg")
	if err != nil {
		t.Fatal(err)
	}
	images := make([]DrawParam, 2)
	images[0] = DrawParam{
		Image: img1,
		X:     0,
		Y:     0,
	}
	images[1] = DrawParam{
		Image: img2,
		X:     0,
		Y:     0,
	}
	reader, err := Draw(ctx, images, canvas)
	if err != nil {
		t.Fatal(err)
	}
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		t.Fatal(err)
	}
	err = ioutil.WriteFile("test.png", b, 0644)
	if err != nil {
		t.Fatal(err)
	}
}
