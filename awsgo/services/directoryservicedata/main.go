package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/directoryservicedata/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"add-group-member", "create-group", "create-user", "delete-group", "delete-user", "describe-group", "describe-user", "disable-user", "list-group-members", "list-groups", "list-groups-for-member", "list-users", "remove-group-member", "search-groups", "search-users", "update-group", "update-user"},
		OperationSet: map[string]bool{"add-group-member": true, "create-group": true, "create-user": true, "delete-group": true, "delete-user": true, "describe-group": true, "describe-user": true, "disable-user": true, "list-group-members": true, "list-groups": true, "list-groups-for-member": true, "list-users": true, "remove-group-member": true, "search-groups": true, "search-users": true, "update-group": true, "update-user": true},
		OperationInputs: map[string][]string{
			"add-group-member":       {"ClientToken", "DirectoryId", "GroupName", "MemberName", "MemberRealm"},
			"create-group":           {"ClientToken", "DirectoryId", "GroupScope", "GroupType", "OtherAttributes", "SAMAccountName"},
			"create-user":            {"ClientToken", "DirectoryId", "EmailAddress", "GivenName", "OtherAttributes", "SAMAccountName", "Surname"},
			"delete-group":           {"ClientToken", "DirectoryId", "SAMAccountName"},
			"delete-user":            {"ClientToken", "DirectoryId", "SAMAccountName"},
			"describe-group":         {"DirectoryId", "OtherAttributes", "Realm", "SAMAccountName"},
			"describe-user":          {"DirectoryId", "OtherAttributes", "Realm", "SAMAccountName"},
			"disable-user":           {"ClientToken", "DirectoryId", "SAMAccountName"},
			"list-group-members":     {"DirectoryId", "MaxResults", "MemberRealm", "NextToken", "Realm", "SAMAccountName"},
			"list-groups":            {"DirectoryId", "MaxResults", "NextToken", "Realm"},
			"list-groups-for-member": {"DirectoryId", "MaxResults", "MemberRealm", "NextToken", "Realm", "SAMAccountName"},
			"list-users":             {"DirectoryId", "MaxResults", "NextToken", "Realm"},
			"remove-group-member":    {"ClientToken", "DirectoryId", "GroupName", "MemberName", "MemberRealm"},
			"search-groups":          {"DirectoryId", "MaxResults", "NextToken", "Realm", "SearchAttributes", "SearchString"},
			"search-users":           {"DirectoryId", "MaxResults", "NextToken", "Realm", "SearchAttributes", "SearchString"},
			"update-group":           {"ClientToken", "DirectoryId", "GroupScope", "GroupType", "OtherAttributes", "SAMAccountName", "UpdateType"},
			"update-user":            {"ClientToken", "DirectoryId", "EmailAddress", "GivenName", "OtherAttributes", "SAMAccountName", "Surname", "UpdateType"},
		},
		OperationInputTypes: map[string]map[string]string{
			"add-group-member":       {"ClientToken": "*string", "DirectoryId": "*string", "GroupName": "*string", "MemberName": "*string", "MemberRealm": "*string"},
			"create-group":           {"ClientToken": "*string", "DirectoryId": "*string", "GroupScope": "types.GroupScope", "GroupType": "types.GroupType", "OtherAttributes": "map[string]types.AttributeValue", "SAMAccountName": "*string"},
			"create-user":            {"ClientToken": "*string", "DirectoryId": "*string", "EmailAddress": "*string", "GivenName": "*string", "OtherAttributes": "map[string]types.AttributeValue", "SAMAccountName": "*string", "Surname": "*string"},
			"delete-group":           {"ClientToken": "*string", "DirectoryId": "*string", "SAMAccountName": "*string"},
			"delete-user":            {"ClientToken": "*string", "DirectoryId": "*string", "SAMAccountName": "*string"},
			"describe-group":         {"DirectoryId": "*string", "OtherAttributes": "[]string", "Realm": "*string", "SAMAccountName": "*string"},
			"describe-user":          {"DirectoryId": "*string", "OtherAttributes": "[]string", "Realm": "*string", "SAMAccountName": "*string"},
			"disable-user":           {"ClientToken": "*string", "DirectoryId": "*string", "SAMAccountName": "*string"},
			"list-group-members":     {"DirectoryId": "*string", "MaxResults": "*int32", "MemberRealm": "*string", "NextToken": "*string", "Realm": "*string", "SAMAccountName": "*string"},
			"list-groups":            {"DirectoryId": "*string", "MaxResults": "*int32", "NextToken": "*string", "Realm": "*string"},
			"list-groups-for-member": {"DirectoryId": "*string", "MaxResults": "*int32", "MemberRealm": "*string", "NextToken": "*string", "Realm": "*string", "SAMAccountName": "*string"},
			"list-users":             {"DirectoryId": "*string", "MaxResults": "*int32", "NextToken": "*string", "Realm": "*string"},
			"remove-group-member":    {"ClientToken": "*string", "DirectoryId": "*string", "GroupName": "*string", "MemberName": "*string", "MemberRealm": "*string"},
			"search-groups":          {"DirectoryId": "*string", "MaxResults": "*int32", "NextToken": "*string", "Realm": "*string", "SearchAttributes": "[]string", "SearchString": "*string"},
			"search-users":           {"DirectoryId": "*string", "MaxResults": "*int32", "NextToken": "*string", "Realm": "*string", "SearchAttributes": "[]string", "SearchString": "*string"},
			"update-group":           {"ClientToken": "*string", "DirectoryId": "*string", "GroupScope": "types.GroupScope", "GroupType": "types.GroupType", "OtherAttributes": "map[string]types.AttributeValue", "SAMAccountName": "*string", "UpdateType": "types.UpdateType"},
			"update-user":            {"ClientToken": "*string", "DirectoryId": "*string", "EmailAddress": "*string", "GivenName": "*string", "OtherAttributes": "map[string]types.AttributeValue", "SAMAccountName": "*string", "Surname": "*string", "UpdateType": "types.UpdateType"},
		},
		OperationInputRequired: map[string][]string{
			"add-group-member":       {"DirectoryId", "GroupName", "MemberName"},
			"create-group":           {"DirectoryId", "SAMAccountName"},
			"create-user":            {"DirectoryId", "SAMAccountName"},
			"delete-group":           {"DirectoryId", "SAMAccountName"},
			"delete-user":            {"DirectoryId", "SAMAccountName"},
			"describe-group":         {"DirectoryId", "SAMAccountName"},
			"describe-user":          {"DirectoryId", "SAMAccountName"},
			"disable-user":           {"DirectoryId", "SAMAccountName"},
			"list-group-members":     {"DirectoryId", "SAMAccountName"},
			"list-groups":            {"DirectoryId"},
			"list-groups-for-member": {"DirectoryId", "SAMAccountName"},
			"list-users":             {"DirectoryId"},
			"remove-group-member":    {"DirectoryId", "GroupName", "MemberName"},
			"search-groups":          {"DirectoryId", "SearchAttributes", "SearchString"},
			"search-users":           {"DirectoryId", "SearchAttributes", "SearchString"},
			"update-group":           {"DirectoryId", "SAMAccountName"},
			"update-user":            {"DirectoryId", "SAMAccountName"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("directoryservicedata", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
