Image
==

**Install**

> go get github.com/framehack/image

**Usage**

- load image from url

	``` go
		img, size, err := image.LoadImageURL(ctx, "image/to/load")
	```

- draw multi images

	```go
		reader, err := image.Draw(ctx, imgList)
	```