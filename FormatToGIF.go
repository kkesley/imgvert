package imgvert

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
)

//FormatToGIF format image to GIF
func FormatToGIF(img image.Image) (bytes.Buffer, error) {
	var w bytes.Buffer //Writer
	newImg := image.NewRGBA(img.Bounds())
	// we will use white background to replace PNG's transparent background
	// you can change it to whichever color you want with
	// a new color.RGBA{} and use image.NewUniform(color.RGBA{<fill in color>}) function

	draw.Draw(newImg, newImg.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)
	// paste PNG image OVER to newImage
	draw.Draw(newImg, newImg.Bounds(), img, img.Bounds().Min, draw.Over)
	var opt gif.Options
	opt.NumColors = 256
	if err := gif.Encode(&w, newImg, &opt); err != nil {
		return w, err
	}
	return w, nil
}
