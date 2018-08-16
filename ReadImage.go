package imgvert

import (
	"fmt"
	"image"
	"image/png"

	parser "github.com/kkesley/s3-parser"
)

func ReadImage(region string, bucket string, key string) (*image.Image, error) {
	// decode jpeg into image.Image
	output, err := parser.GetS3DocumentDefaultRaw(region, bucket, key)
	if err != nil {
		return nil, err
	}
	_, format, err := image.DecodeConfig(output.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(format)
	img, err := png.Decode(output.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(img)

	return &img, nil
}
