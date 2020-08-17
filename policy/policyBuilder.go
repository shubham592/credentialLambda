package policy

/*
Go File to Session Policy
@Shubham Urkude
@06-05-2020
*/

type PolicyDocument struct {
	Version   string           `json:"Version"`
	Statement []StatementEntry `json:"Statement"`
}

type StatementEntry struct {
	Effect   string   `json:"Effect"`
	Action   []string `json:"Action"`
	Resource string   `json:"Resource"`
}

func BuildPolicy(thingName string) PolicyDocument {
	//thingARN_connect :=
	//"arn:aws:iot:us-west-2:617260006156:client/9baff45c-1441-407a-80fb-5addf7867436"

	policy := PolicyDocument{
		Version: "2012-10-17",
		Statement: []StatementEntry{
			StatementEntry{
				Effect: "Allow",
				Action: []string{
					"iot:Connect",
				},
				Resource: "arn:aws:iot:us-west-2:434851320724:client/"+thingName,
			},
			StatementEntry{
				Effect: "Allow",
				Action: []string{
					"iot:Publish",
				},
				Resource: "arn:aws:iot:us-west-2:434851320724:topic/$aws/rules/daas_device_events_rule_usdevms/"+thingName,
			},
			/*StatementEntry{
				Effect: "Allow",
				Action: []string{
					"iot:UpdateJobExecution",
				},
				Resource: thingARN_job,
			},
			StatementEntry{
				Effect: "Allow",
				Action: []string{
					"iot:StartNextPendingJobExecution",
				},
				Resource: thingARN_job,
			},*/
		},
	}
	return policy
}
