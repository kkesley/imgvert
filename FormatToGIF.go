package imgvert

import (
	"fmt"
	"image"
	"image/gif"
	"os"
)

func FormatToGIF(img image.Image) {
	out, err := os.Create("converterdTOgif.gif")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var opt gif.Options
	opt.NumColors = 256
	err = gif.Encode(out, img, &opt)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(img, "\n success... \n ")
}
