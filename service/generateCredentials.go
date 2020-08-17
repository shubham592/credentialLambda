package service

import (
	"creds_lambda/policy"
	"encoding/json"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/aws/aws-sdk-go/service/sts/stsiface"
	"log"
)

/*
Function to generate session token
@Shubham Urkude
*/

func GenerateCredentials(svc stsiface.STSAPI, thingName string) (*sts.AssumeRoleOutput, error) {
	roleARN :="arn:aws:iam::617260006156:role/IoTCredTestRole"
		//"arn:aws:iam::434851320724:role/IoTDaasDeviceRole.usdevms"


	roleSessionName := "IoTDaasDeviceRole"
	
	var duration int64= 43200

	policyDoc:=policy.BuildPolicy(thingName)
	policyJson, er :=json.Marshal(policyDoc)
	sessionPolicy:=string(policyJson)
	if er != nil{
		return nil,er
	}

	input := sts.AssumeRoleInput{
		DurationSeconds: &duration,
		Policy:          &sessionPolicy,
		RoleArn:         &roleARN,
		RoleSessionName: &roleSessionName,
	}

	output, error := svc.AssumeRole(&input)
	if error != nil {
		log.Println("error:", error.Error())
		return nil, error
	}
	//credentials := *output.Credentials
	log.Println("Credentials:", *output.Credentials)
	return output, nil
}