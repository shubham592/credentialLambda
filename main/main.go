package main

import (
	"creds_lambda/handler"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

/*
Go Main package for lambda Handler
@Shubham Urkude
*/

func main() {
	log.Println("Lambda Request received..!!")
	lambda.Start(handler.HandleRequest)
	//handler.HandleRequest();
}