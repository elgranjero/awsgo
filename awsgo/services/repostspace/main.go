package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/repostspace/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"batch-add-channel-role-to-accessors", "batch-add-role", "batch-remove-channel-role-from-accessors", "batch-remove-role", "create-channel", "create-space", "delete-space", "deregister-admin", "get-channel", "get-space", "list-channels", "list-spaces", "list-tags-for-resource", "register-admin", "send-invites", "tag-resource", "untag-resource", "update-channel", "update-space"},
		OperationSet: map[string]bool{"batch-add-channel-role-to-accessors": true, "batch-add-role": true, "batch-remove-channel-role-from-accessors": true, "batch-remove-role": true, "create-channel": true, "create-space": true, "delete-space": true, "deregister-admin": true, "get-channel": true, "get-space": true, "list-channels": true, "list-spaces": true, "list-tags-for-resource": true, "register-admin": true, "send-invites": true, "tag-resource": true, "untag-resource": true, "update-channel": true, "update-space": true},
		OperationInputs: map[string][]string{
			"batch-add-channel-role-to-accessors":      {"AccessorIds", "ChannelId", "ChannelRole", "SpaceId"},
			"batch-add-role":                           {"AccessorIds", "Role", "SpaceId"},
			"batch-remove-channel-role-from-accessors": {"AccessorIds", "ChannelId", "ChannelRole", "SpaceId"},
			"batch-remove-role":                        {"AccessorIds", "Role", "SpaceId"},
			"create-channel":                           {"ChannelDescription", "ChannelName", "SpaceId"},
			"create-space":                             {"Description", "Name", "RoleArn", "Subdomain", "SupportedEmailDomains", "Tags", "Tier", "UserKMSKey"},
			"delete-space":                             {"SpaceId"},
			"deregister-admin":                         {"AdminId", "SpaceId"},
			"get-channel":                              {"ChannelId", "SpaceId"},
			"get-space":                                {"SpaceId"},
			"list-channels":                            {"MaxResults", "NextToken", "SpaceId"},
			"list-spaces":                              {"MaxResults", "NextToken"},
			"list-tags-for-resource":                   {"ResourceArn"},
			"register-admin":                           {"AdminId", "SpaceId"},
			"send-invites":                             {"AccessorIds", "Body", "SpaceId", "Title"},
			"tag-resource":                             {"ResourceArn", "Tags"},
			"untag-resource":                           {"ResourceArn", "TagKeys"},
			"update-channel":                           {"ChannelDescription", "ChannelId", "ChannelName", "SpaceId"},
			"update-space":                             {"Description", "RoleArn", "SpaceId", "SupportedEmailDomains", "Tier"},
		},
		OperationInputTypes: map[string]map[string]string{
			"batch-add-channel-role-to-accessors":      {"AccessorIds": "[]string", "ChannelId": "*string", "ChannelRole": "types.ChannelRole", "SpaceId": "*string"},
			"batch-add-role":                           {"AccessorIds": "[]string", "Role": "types.Role", "SpaceId": "*string"},
			"batch-remove-channel-role-from-accessors": {"AccessorIds": "[]string", "ChannelId": "*string", "ChannelRole": "types.ChannelRole", "SpaceId": "*string"},
			"batch-remove-role":                        {"AccessorIds": "[]string", "Role": "types.Role", "SpaceId": "*string"},
			"create-channel":                           {"ChannelDescription": "*string", "ChannelName": "*string", "SpaceId": "*string"},
			"create-space":                             {"Description": "*string", "Name": "*string", "RoleArn": "*string", "Subdomain": "*string", "SupportedEmailDomains": "*types.SupportedEmailDomainsParameters", "Tags": "map[string]string", "Tier": "types.TierLevel", "UserKMSKey": "*string"},
			"delete-space":                             {"SpaceId": "*string"},
			"deregister-admin":                         {"AdminId": "*string", "SpaceId": "*string"},
			"get-channel":                              {"ChannelId": "*string", "SpaceId": "*string"},
			"get-space":                                {"SpaceId": "*string"},
			"list-channels":                            {"MaxResults": "*int32", "NextToken": "*string", "SpaceId": "*string"},
			"list-spaces":                              {"MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":                   {"ResourceArn": "*string"},
			"register-admin":                           {"AdminId": "*string", "SpaceId": "*string"},
			"send-invites":                             {"AccessorIds": "[]string", "Body": "*string", "SpaceId": "*string", "Title": "*string"},
			"tag-resource":                             {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":                           {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-channel":                           {"ChannelDescription": "*string", "ChannelId": "*string", "ChannelName": "*string", "SpaceId": "*string"},
			"update-space":                             {"Description": "*string", "RoleArn": "*string", "SpaceId": "*string", "SupportedEmailDomains": "*types.SupportedEmailDomainsParameters", "Tier": "types.TierLevel"},
		},
		OperationInputRequired: map[string][]string{
			"batch-add-channel-role-to-accessors":      {"AccessorIds", "ChannelId", "ChannelRole", "SpaceId"},
			"batch-add-role":                           {"AccessorIds", "Role", "SpaceId"},
			"batch-remove-channel-role-from-accessors": {"AccessorIds", "ChannelId", "ChannelRole", "SpaceId"},
			"batch-remove-role":                        {"AccessorIds", "Role", "SpaceId"},
			"create-channel":                           {"ChannelName", "SpaceId"},
			"create-space":                             {"Name", "Subdomain", "Tier"},
			"delete-space":                             {"SpaceId"},
			"deregister-admin":                         {"AdminId", "SpaceId"},
			"get-channel":                              {"ChannelId", "SpaceId"},
			"get-space":                                {"SpaceId"},
			"list-channels":                            {"SpaceId"},
			"list-spaces":                              {},
			"list-tags-for-resource":                   {"ResourceArn"},
			"register-admin":                           {"AdminId", "SpaceId"},
			"send-invites":                             {"AccessorIds", "Body", "SpaceId", "Title"},
			"tag-resource":                             {"ResourceArn", "Tags"},
			"untag-resource":                           {"ResourceArn", "TagKeys"},
			"update-channel":                           {"ChannelId", "ChannelName", "SpaceId"},
			"update-space":                             {"SpaceId"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("repostspace", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
