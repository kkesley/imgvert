package imgvert

import (
	"bytes"
	"image"
	"image/png"
)

//FormatToPNG format image to PNG
func FormatToPNG(img image.Image) (bytes.Buffer, error) {
	var w bytes.Buffer //Writer
	enc := png.Encoder{
		CompressionLevel: png.NoCompression,
	}
	if err := enc.Encode(&w, img); err != nil {
		return w, err
	}
	return w, nil
}
