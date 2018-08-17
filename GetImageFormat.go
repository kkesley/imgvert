package imgvert

import (
	"bytes"
	"image"
)

//GetImageFormat get a format of an image
func GetImageFormat(data []byte) (string, error) {
	formatReader := bytes.NewReader(data)
	_, format, err := image.DecodeConfig(formatReader)
	return format, err
}
