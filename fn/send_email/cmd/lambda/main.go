package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/monitprod/send_email/pkg/handler"
)

func main() {
	lambda.Start(handler.HandleRequest)
}
