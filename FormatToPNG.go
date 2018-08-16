package imgvert

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

func Formatpng(img image.Image) {
	out, err := os.Create("converterdTOPNG.png")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = png.Encode(out, img)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(img, "\n success... \n ")

}
