package utils

import (
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/tencentyun/cos-go-sdk-v5/debug"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// getCos 获取腾讯云cos client
func getCos(bucketName string) *cos.Client {
	appID := os.Getenv("APP_ID")
	region := os.Getenv("COS_REGION_B")
	secretID := os.Getenv("SECRET_ID")
	secretKey := os.Getenv("SECRET_KEY")
	l := fmt.Sprintf("https://%v-%v.%v", bucketName, appID, region)
	u, _ := url.Parse(l)

	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  secretID,
			SecretKey: secretKey,
			Transport: &debug.DebugRequestTransport{
				RequestHeader:  true,
				RequestBody:    false,
				ResponseHeader: true,
				ResponseBody:   true,
				Writer:         nil,
			},
		},
	})
	return client
}

func CheckErr(err error) bool {
	if err == nil {
		return true
	}
	if cos.IsNotFoundError(err) {
		log.Println("WARN: Resource is not existed")
		return false
	} else if e, ok := cos.IsCOSError(err); ok {
		log.Println(fmt.Sprintf("ERROR: Code: %v\n", e.Code))
		log.Println(fmt.Sprintf("ERROR: Message: %v\n", e.Message))
		log.Println(fmt.Sprintf("ERROR: Resource: %v\n", e.Resource))
		log.Println(fmt.Sprintf("ERROR: RequestID: %v\n", e.RequestID))
		return false
	} else {
		log.Println(fmt.Sprintf("ERROR: %v\n", err))
		return false
	}
}

func Upload(ctx context.Context, bucketName, fileName string, file *multipart.File) (string, error) {
	client := getCos(bucketName)
	_, err := client.Object.Put(ctx, fileName, *file, nil)
	if err != nil {
		return "", err
	}
	appID := os.Getenv("APP_ID")
	region := os.Getenv("COS_REGION_B")
	return fmt.Sprintf("https://%v-%v.%v/%v", bucketName, appID, region, fileName), nil

}

func Delete(cosUrl string) error {
	parts := strings.Split(cosUrl, "/")
	appID := os.Getenv("APP_ID")
	bucketParts := strings.Split(parts[2], "-"+appID)
	bucket := bucketParts[0]
	fileName := strings.Join(parts[3:], "/")
	c := getCos(bucket)
	_, err := c.Object.Delete(context.Background(), fileName, nil)
	if err != nil {
		return err
	}
	return nil
}
