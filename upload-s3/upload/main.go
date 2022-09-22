package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"bytes"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
    "github.com/aws/aws-sdk-go/private/protocol/rest"
	"time"
	"strconv"
	"strings"
)


type ImageRequestBody struct {
	FileName string `json:"filename"`
	Body     string `json:"body"`
}

type ImageUploadResponse struct {
	FileName string `json:"filename"`
	Url string `json:"url"`
}

const (
    AWS_S3_REGION = "us-east-1"
    AWS_S3_BUCKET = "golang-s3-upload"
)

type ResponseToSend events.APIGatewayProxyResponse

var s3session *s3.S3

func imageUpload(bodyRequest *ImageRequestBody, data []byte) ImageUploadResponse {

	s3session := s3.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String(AWS_S3_REGION),
	  })))

	uniqueNumber := time.Now().UnixNano()/(1<<22)
		
	d := strconv.FormatInt(int64(uniqueNumber), 10)
	split_name := strings.Split(bodyRequest.FileName, ".")

	file_name := split_name[0] + "_" +  d + "." + split_name[1]
	
	result, _ := s3session.PutObject(&s3.PutObjectInput{
		Body: bytes.NewReader(data),
		Bucket: aws.String(AWS_S3_BUCKET),
		Key: aws.String(file_name),
	  })

	fmt.Println(result)

	res, _ := s3session.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(AWS_S3_BUCKET),
		Key:    aws.String(file_name),
	})
		rest.Build(res)
		uploadedResourceLocation := res.HTTPRequest.URL.String()

	responseData := ImageUploadResponse{
			FileName: file_name,
			Url: uploadedResourceLocation,
		}
		return responseData
}

func Handler(request events.APIGatewayProxyRequest) (ResponseToSend, error) {

	// Extract the request body
	bodyRequest := &ImageRequestBody{}
	err := json.Unmarshal([]byte(request.Body), &bodyRequest)

	if err != nil {
		return ResponseToSend{Body: err.Error(), StatusCode: 404}, nil
	}

	decoded, err := base64.StdEncoding.DecodeString(bodyRequest.Body)
	if err != nil {
		return ResponseToSend{Body: err.Error(), StatusCode: 404}, nil
	}

	resp := imageUpload(bodyRequest, decoded)

	response, err := json.Marshal(&resp)
	if err != nil {
		return ResponseToSend{Body: err.Error(), StatusCode: 404}, nil
	}

	respToSend := ResponseToSend{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(response),
	}

	return respToSend, nil
}


func main() {
	lambda.Start(Handler)
}