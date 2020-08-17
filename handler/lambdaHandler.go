package handler

import (
	"context"
	"creds_lambda/service"
	"creds_lambda/shadow"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iot"
	"github.com/aws/aws-sdk-go/service/iotdataplane"
	"github.com/aws/aws-sdk-go/service/sts"
	"log"
)

/*
Lambda handler to handle generate Credentials Request
@Shubham Urkude
*/

type MyEvent struct {
	ThingName string `json:"thingName"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	thingName := name.ThingName
	//thingName := "9baff45c-1441-407a-80fb-5addf7867436"
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})

	ses, err := session.NewSession(&aws.Config{
		Endpoint: iotEndpoint,
		Region:   aws.String("us-west-2"),
	})

	svc := sts.New(sess)
	svcShadow := iotdataplane.New(ses)
	output, err2 := service.GenerateCredentials(svc, thingName)
	if err2 != nil {
		return "nil", err2
	}
	err = shadow.UpdateThingShadow(output, svcShadow, &thingName)
	if err != nil {
		return "nil", err
	}
	return "Shadow Updated Successfully..!!", nil
}

var AWSClient *iot.IoT

func GetClientInstance() (*iot.IoT, error) {
	if AWSClient != nil {
		return AWSClient, nil
	} else {
		sess, err := session.NewSession(&aws.Config{
			Region: aws.String("us-west-2"),
		})
		if err != nil {
			return nil, err
		}
		AWSClient = iot.New(sess)
		return AWSClient, err
	}
}

var iotEndpoint *string

func init() {
	se, _ := GetClientInstance()
	var endpointType = "iot:Data-ATS"
	endpointReq, error := se.DescribeEndpoint(&iot.DescribeEndpointInput{EndpointType: &endpointType})

	if error != nil {
		log.Println("Error While Getting IoT Data EndPoint:", error)
	}
	iotEndpoint = endpointReq.EndpointAddress
}
