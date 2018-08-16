package imgvert

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"
)

func FormatToJPG(img image.Image) {
	out, err := os.Create("convertToJPG.jpg")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	newImg := image.NewRGBA(img.Bounds())

	// we will use white background to replace PNG's transparent background
	// you can change it to whichever color you want with
	// a new color.RGBA{} and use image.NewUniform(color.RGBA{<fill in color>}) function

	draw.Draw(newImg, newImg.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)

	// paste PNG image OVER to newImage
	draw.Draw(newImg, newImg.Bounds(), img, img.Bounds().Min, draw.Over)
	err = jpeg.Encode(out, newImg, &jpeg.Options{
		Quality: 80,
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(img, "\n success... \n ")

}
