package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/connectcampaigns/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-campaign", "delete-campaign", "delete-connect-instance-config", "delete-instance-onboarding-job", "describe-campaign", "get-campaign-state", "get-campaign-state-batch", "get-connect-instance-config", "get-instance-onboarding-job-status", "list-campaigns", "list-tags-for-resource", "pause-campaign", "put-dial-request-batch", "resume-campaign", "start-campaign", "start-instance-onboarding-job", "stop-campaign", "tag-resource", "untag-resource", "update-campaign-dialer-config", "update-campaign-name", "update-campaign-outbound-call-config"},
		OperationSet: map[string]bool{"create-campaign": true, "delete-campaign": true, "delete-connect-instance-config": true, "delete-instance-onboarding-job": true, "describe-campaign": true, "get-campaign-state": true, "get-campaign-state-batch": true, "get-connect-instance-config": true, "get-instance-onboarding-job-status": true, "list-campaigns": true, "list-tags-for-resource": true, "pause-campaign": true, "put-dial-request-batch": true, "resume-campaign": true, "start-campaign": true, "start-instance-onboarding-job": true, "stop-campaign": true, "tag-resource": true, "untag-resource": true, "update-campaign-dialer-config": true, "update-campaign-name": true, "update-campaign-outbound-call-config": true},
		OperationInputs: map[string][]string{
			"create-campaign":                      {"ConnectInstanceId", "DialerConfig", "Name", "OutboundCallConfig", "Tags"},
			"delete-campaign":                      {"Id"},
			"delete-connect-instance-config":       {"ConnectInstanceId"},
			"delete-instance-onboarding-job":       {"ConnectInstanceId"},
			"describe-campaign":                    {"Id"},
			"get-campaign-state":                   {"Id"},
			"get-campaign-state-batch":             {"CampaignIds"},
			"get-connect-instance-config":          {"ConnectInstanceId"},
			"get-instance-onboarding-job-status":   {"ConnectInstanceId"},
			"list-campaigns":                       {"Filters", "MaxResults", "NextToken"},
			"list-tags-for-resource":               {"Arn"},
			"pause-campaign":                       {"Id"},
			"put-dial-request-batch":               {"DialRequests", "Id"},
			"resume-campaign":                      {"Id"},
			"start-campaign":                       {"Id"},
			"start-instance-onboarding-job":        {"ConnectInstanceId", "EncryptionConfig"},
			"stop-campaign":                        {"Id"},
			"tag-resource":                         {"Arn", "Tags"},
			"untag-resource":                       {"Arn", "TagKeys"},
			"update-campaign-dialer-config":        {"DialerConfig", "Id"},
			"update-campaign-name":                 {"Id", "Name"},
			"update-campaign-outbound-call-config": {"AnswerMachineDetectionConfig", "ConnectContactFlowId", "ConnectSourcePhoneNumber", "Id"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-campaign":                      {"ConnectInstanceId": "*string", "DialerConfig": "types.DialerConfig", "Name": "*string", "OutboundCallConfig": "*types.OutboundCallConfig", "Tags": "map[string]string"},
			"delete-campaign":                      {"Id": "*string"},
			"delete-connect-instance-config":       {"ConnectInstanceId": "*string"},
			"delete-instance-onboarding-job":       {"ConnectInstanceId": "*string"},
			"describe-campaign":                    {"Id": "*string"},
			"get-campaign-state":                   {"Id": "*string"},
			"get-campaign-state-batch":             {"CampaignIds": "[]string"},
			"get-connect-instance-config":          {"ConnectInstanceId": "*string"},
			"get-instance-onboarding-job-status":   {"ConnectInstanceId": "*string"},
			"list-campaigns":                       {"Filters": "*types.CampaignFilters", "MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":               {"Arn": "*string"},
			"pause-campaign":                       {"Id": "*string"},
			"put-dial-request-batch":               {"DialRequests": "[]types.DialRequest", "Id": "*string"},
			"resume-campaign":                      {"Id": "*string"},
			"start-campaign":                       {"Id": "*string"},
			"start-instance-onboarding-job":        {"ConnectInstanceId": "*string", "EncryptionConfig": "*types.EncryptionConfig"},
			"stop-campaign":                        {"Id": "*string"},
			"tag-resource":                         {"Arn": "*string", "Tags": "map[string]string"},
			"untag-resource":                       {"Arn": "*string", "TagKeys": "[]string"},
			"update-campaign-dialer-config":        {"DialerConfig": "types.DialerConfig", "Id": "*string"},
			"update-campaign-name":                 {"Id": "*string", "Name": "*string"},
			"update-campaign-outbound-call-config": {"AnswerMachineDetectionConfig": "*types.AnswerMachineDetectionConfig", "ConnectContactFlowId": "*string", "ConnectSourcePhoneNumber": "*string", "Id": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"create-campaign":                      {"ConnectInstanceId", "DialerConfig", "Name", "OutboundCallConfig"},
			"delete-campaign":                      {"Id"},
			"delete-connect-instance-config":       {"ConnectInstanceId"},
			"delete-instance-onboarding-job":       {"ConnectInstanceId"},
			"describe-campaign":                    {"Id"},
			"get-campaign-state":                   {"Id"},
			"get-campaign-state-batch":             {"CampaignIds"},
			"get-connect-instance-config":          {"ConnectInstanceId"},
			"get-instance-onboarding-job-status":   {"ConnectInstanceId"},
			"list-campaigns":                       {},
			"list-tags-for-resource":               {"Arn"},
			"pause-campaign":                       {"Id"},
			"put-dial-request-batch":               {"DialRequests", "Id"},
			"resume-campaign":                      {"Id"},
			"start-campaign":                       {"Id"},
			"start-instance-onboarding-job":        {"ConnectInstanceId", "EncryptionConfig"},
			"stop-campaign":                        {"Id"},
			"tag-resource":                         {"Arn", "Tags"},
			"untag-resource":                       {"Arn", "TagKeys"},
			"update-campaign-dialer-config":        {"DialerConfig", "Id"},
			"update-campaign-name":                 {"Id", "Name"},
			"update-campaign-outbound-call-config": {"Id"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("connectcampaigns", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
