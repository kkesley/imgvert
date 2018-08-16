package imgvert

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func GetS3(region string, bucket string, key string) (error, error) {
	config := aws.Config{
		Region: aws.String(region),
	}
	sess := session.Must(session.NewSession(&config))

	svc := s3.New(sess)

	s3Output, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(s3Output.Body)

	imgFormat := bytes.NewReader(buf.Bytes())
	imgReader := bytes.NewReader(buf.Bytes())
	_, format, err := image.DecodeConfig(imgFormat)
	if err != nil {
		return nil, err
	}
	fmt.Println(format)
	img, err := png.Decode(imgReader)
	if err != nil {
		return nil, err
	}
	fmt.Println(img)
	return nil, nil
}

func TestFormatToJPG(test *testing.T) {
	ImageFile := "test.png"
	GetS3("ap-southeast-1", "iteacloud", ImageFile)
	// img, err := png.Decode(output.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(img)

	// config := aws.Config{
	// 	Region: aws.String("ap-southeast-1"),
	// }
	// sess := session.Must(session.NewSession(&config))
	// buff := &aws.WriteAtBuffer{}
	// s3dl := s3manager.NewDownloader(sess)
	// _, err := s3dl.Download(buff, &s3.GetObjectInput{
	// 	Bucket: aws.String("iteacloud"),
	// 	Key:    aws.String(ImageFile),
	// })

	// if err != nil {
	// 	log.Printf("Could not download from S3: %v", err)
	// }

	// png.Decode()

	// output, err := ReadImage("ap-southeast-1", "iteacloud", ImageFile)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(output)
}
