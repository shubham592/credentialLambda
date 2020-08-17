package shadow

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go/service/iotdataplane"
	"github.com/aws/aws-sdk-go/service/sts"
)

/*
Go function for Updating Credentials into Thing Shadow.
@Shubham Urkude
@04-05-2020
*/

type Data struct {
	State State `json:"state"`
}
type Desired struct {
	Credentials string `json:"Credentials"`
}
type State struct {
	Desired Desired `json:"desired"`
}

func UpdateThingShadow(output *sts.AssumeRoleOutput, session *iotdataplane.IoTDataPlane, ThingName *string) error {

	cred := output.Credentials.GoString()
	inputData := Data{State{Desired: Desired{Credentials: cred}}}
	data, _ := json.Marshal(inputData)

	_, err := session.UpdateThingShadow(&iotdataplane.UpdateThingShadowInput{
		Payload:   data,
		ThingName: ThingName,
	})
	if err != nil {
		return err
	}
	log.Printf("Thing Shadow Updated SuccessFully for ThingName:%v\n", ThingName)
	return nil
}
