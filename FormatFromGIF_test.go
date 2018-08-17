package imgvert

import (
	"fmt"
	"io/ioutil"
	"testing"

	parser "github.com/kkesley/s3-parser"
)

func TestFormatFromGIF(test *testing.T) {
	ImageFile := "test.gif"
	data, err := ioutil.ReadFile(ImageFile)
	if err != nil {
		fmt.Println(err)
		test.Fail()
	}
	img, err := ReadImage(data)
	if err != nil {
		fmt.Println(err)
		test.Fail()
	}
	jpg, err := FormatToJPG(img)
	if err != nil {
		fmt.Println(err)
		test.Fail()
	}
	if err := parser.PutS3DocumentDefault(parser.PutS3DocumentDefaultInput{
		Region:      "ap-southeast-1",
		Bucket:      "iteacloud",
		Key:         "from-gif-to-jpg.jpg",
		ContentType: "image/jpeg",
		Content:     jpg.Bytes(),
	}); err != nil {
		fmt.Println(err)
		test.Fail()
	}

	gif, err := FormatToGIF(img)
	if err != nil {
		fmt.Println(err)
		test.Fail()
	}
	if err := parser.PutS3DocumentDefault(parser.PutS3DocumentDefaultInput{
		Region:      "ap-southeast-1",
		Bucket:      "iteacloud",
		Key:         "from-gif-to-gif.gif",
		ContentType: "image/gif",
		Content:     gif.Bytes(),
	}); err != nil {
		fmt.Println(err)
		test.Fail()
	}

	png, err := FormatToPNG(img)
	if err != nil {
		fmt.Println(err)
		test.Fail()
	}
	if err := parser.PutS3DocumentDefault(parser.PutS3DocumentDefaultInput{
		Region:      "ap-southeast-1",
		Bucket:      "iteacloud",
		Key:         "from-gif-to-png.png",
		ContentType: "image/png",
		Content:     png.Bytes(),
	}); err != nil {
		fmt.Println(err)
		test.Fail()
	}
}
