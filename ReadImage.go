package imgvert

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
)

func ReadImage(data []byte) (image.Image, error) {
	// decode jpeg into image.Image
	formatReader := bytes.NewReader(data)
	_, format, err := image.DecodeConfig(formatReader)
	if err != nil {
		return nil, err
	}
	fmt.Println(format)
	imageReader := bytes.NewReader(data)
	if format == "png" {
		return png.Decode(imageReader)
	} else if format == "jpg" {
		return jpeg.Decode(imageReader)
	} else if format == "gif" {
		return gif.Decode(imageReader)
	}
	return nil, errors.New("Format not supported")
}
