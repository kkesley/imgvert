package imgvert

import (
	"fmt"
	"io/ioutil"
	"testing"

	parser "github.com/kkesley/s3-parser"
)

func TestFormatToJPG(test *testing.T) {
	ImageFile := "test.png"
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
	fmt.Println(img)
	jpg, err := FormatToJPG(img)
	if err != nil {
		fmt.Println(err)
		test.Fail()
	}
	parser.PutS3DocumentDefault(parser.PutS3DocumentDefaultInput{
		Region:      "ap-southeast-1",
		Bucket:      "iteacloud",
		Key:         "hello.jpg",
		ContentType: "image/jpeg",
		Content:     jpg.Bytes(),
	})
}
