package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

type LambdaHelper struct {
	client *lambda.Lambda
}

func NewLambdaHelper(env string) *LambdaHelper {
	var region string

	if env == "qa" {
		region = "ap-northeast-1"
	}
	if env == "prod" {
		region = "ap-northeast-2"
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	client := lambda.New(sess, &aws.Config{Region: aws.String(region)})

	return &LambdaHelper{
		client: client,
	}
}

func (lh *LambdaHelper) Invoke(functionName string, paylaod interface{}) {
	payload, err := json.Marshal(paylaod)
	if err != nil {
		fmt.Println("Error marshalling MyGetItemsFunction request")
		os.Exit(0)
	}

	result, err := lh.client.Invoke(&lambda.InvokeInput{FunctionName: aws.String(functionName), Payload: payload})
	if err != nil {
		fmt.Println("Error calling Function")
		os.Exit(0)
	}

	var resp interface{}
	err = json.Unmarshal(result.Payload, &resp)
	if err != nil {
		fmt.Println("Error unmarshalling function response")
		os.Exit(0)
	}

	fmt.Println(resp)
}
