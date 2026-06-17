package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/supportapp/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-slack-channel-configuration", "delete-account-alias", "delete-slack-channel-configuration", "delete-slack-workspace-configuration", "get-account-alias", "list-slack-channel-configurations", "list-slack-workspace-configurations", "put-account-alias", "register-slack-workspace-for-organization", "update-slack-channel-configuration"},
		OperationSet: map[string]bool{"create-slack-channel-configuration": true, "delete-account-alias": true, "delete-slack-channel-configuration": true, "delete-slack-workspace-configuration": true, "get-account-alias": true, "list-slack-channel-configurations": true, "list-slack-workspace-configurations": true, "put-account-alias": true, "register-slack-workspace-for-organization": true, "update-slack-channel-configuration": true},
		OperationInputs: map[string][]string{
			"create-slack-channel-configuration":        {"ChannelId", "ChannelName", "ChannelRoleArn", "NotifyOnAddCorrespondenceToCase", "NotifyOnCaseSeverity", "NotifyOnCreateOrReopenCase", "NotifyOnResolveCase", "TeamId"},
			"delete-account-alias":                      {},
			"delete-slack-channel-configuration":        {"ChannelId", "TeamId"},
			"delete-slack-workspace-configuration":      {"TeamId"},
			"get-account-alias":                         {},
			"list-slack-channel-configurations":         {"NextToken"},
			"list-slack-workspace-configurations":       {"NextToken"},
			"put-account-alias":                         {"AccountAlias"},
			"register-slack-workspace-for-organization": {"TeamId"},
			"update-slack-channel-configuration":        {"ChannelId", "ChannelName", "ChannelRoleArn", "NotifyOnAddCorrespondenceToCase", "NotifyOnCaseSeverity", "NotifyOnCreateOrReopenCase", "NotifyOnResolveCase", "TeamId"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-slack-channel-configuration":        {"ChannelId": "*string", "ChannelName": "*string", "ChannelRoleArn": "*string", "NotifyOnAddCorrespondenceToCase": "*bool", "NotifyOnCaseSeverity": "types.NotificationSeverityLevel", "NotifyOnCreateOrReopenCase": "*bool", "NotifyOnResolveCase": "*bool", "TeamId": "*string"},
			"delete-account-alias":                      {},
			"delete-slack-channel-configuration":        {"ChannelId": "*string", "TeamId": "*string"},
			"delete-slack-workspace-configuration":      {"TeamId": "*string"},
			"get-account-alias":                         {},
			"list-slack-channel-configurations":         {"NextToken": "*string"},
			"list-slack-workspace-configurations":       {"NextToken": "*string"},
			"put-account-alias":                         {"AccountAlias": "*string"},
			"register-slack-workspace-for-organization": {"TeamId": "*string"},
			"update-slack-channel-configuration":        {"ChannelId": "*string", "ChannelName": "*string", "ChannelRoleArn": "*string", "NotifyOnAddCorrespondenceToCase": "*bool", "NotifyOnCaseSeverity": "types.NotificationSeverityLevel", "NotifyOnCreateOrReopenCase": "*bool", "NotifyOnResolveCase": "*bool", "TeamId": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"create-slack-channel-configuration":        {"ChannelId", "ChannelRoleArn", "NotifyOnCaseSeverity", "TeamId"},
			"delete-account-alias":                      {},
			"delete-slack-channel-configuration":        {"ChannelId", "TeamId"},
			"delete-slack-workspace-configuration":      {"TeamId"},
			"get-account-alias":                         {},
			"list-slack-channel-configurations":         {},
			"list-slack-workspace-configurations":       {},
			"put-account-alias":                         {"AccountAlias"},
			"register-slack-workspace-for-organization": {"TeamId"},
			"update-slack-channel-configuration":        {"ChannelId", "TeamId"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("supportapp", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
