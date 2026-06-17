package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/identitystore/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-group", "create-group-membership", "create-user", "delete-group", "delete-group-membership", "delete-user", "describe-group", "describe-group-membership", "describe-user", "get-group-id", "get-group-membership-id", "get-user-id", "is-member-in-groups", "list-group-memberships", "list-group-memberships-for-member", "list-groups", "list-users", "update-group", "update-user"},
		OperationSet: map[string]bool{"create-group": true, "create-group-membership": true, "create-user": true, "delete-group": true, "delete-group-membership": true, "delete-user": true, "describe-group": true, "describe-group-membership": true, "describe-user": true, "get-group-id": true, "get-group-membership-id": true, "get-user-id": true, "is-member-in-groups": true, "list-group-memberships": true, "list-group-memberships-for-member": true, "list-groups": true, "list-users": true, "update-group": true, "update-user": true},
		OperationInputs: map[string][]string{
			"create-group":                      {"Description", "DisplayName", "IdentityStoreId"},
			"create-group-membership":           {"GroupId", "IdentityStoreId", "MemberId"},
			"create-user":                       {"Addresses", "Birthdate", "DisplayName", "Emails", "Extensions", "IdentityStoreId", "Locale", "Name", "NickName", "PhoneNumbers", "Photos", "PreferredLanguage", "ProfileUrl", "Roles", "Timezone", "Title", "UserName", "UserType", "Website"},
			"delete-group":                      {"GroupId", "IdentityStoreId"},
			"delete-group-membership":           {"IdentityStoreId", "MembershipId"},
			"delete-user":                       {"IdentityStoreId", "UserId"},
			"describe-group":                    {"GroupId", "IdentityStoreId"},
			"describe-group-membership":         {"IdentityStoreId", "MembershipId"},
			"describe-user":                     {"Extensions", "IdentityStoreId", "UserId"},
			"get-group-id":                      {"AlternateIdentifier", "IdentityStoreId"},
			"get-group-membership-id":           {"GroupId", "IdentityStoreId", "MemberId"},
			"get-user-id":                       {"AlternateIdentifier", "IdentityStoreId"},
			"is-member-in-groups":               {"GroupIds", "IdentityStoreId", "MemberId"},
			"list-group-memberships":            {"GroupId", "IdentityStoreId", "MaxResults", "NextToken"},
			"list-group-memberships-for-member": {"IdentityStoreId", "MaxResults", "MemberId", "NextToken"},
			"list-groups":                       {"Filters", "IdentityStoreId", "MaxResults", "NextToken"},
			"list-users":                        {"Extensions", "Filters", "IdentityStoreId", "MaxResults", "NextToken"},
			"update-group":                      {"GroupId", "IdentityStoreId", "Operations"},
			"update-user":                       {"IdentityStoreId", "Operations", "UserId"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-group":                      {"Description": "*string", "DisplayName": "*string", "IdentityStoreId": "*string"},
			"create-group-membership":           {"GroupId": "*string", "IdentityStoreId": "*string", "MemberId": "types.MemberId"},
			"create-user":                       {"Addresses": "[]types.Address", "Birthdate": "*string", "DisplayName": "*string", "Emails": "[]types.Email", "Extensions": "map[string]document.Interface", "IdentityStoreId": "*string", "Locale": "*string", "Name": "*types.Name", "NickName": "*string", "PhoneNumbers": "[]types.PhoneNumber", "Photos": "[]types.Photo", "PreferredLanguage": "*string", "ProfileUrl": "*string", "Roles": "[]types.Role", "Timezone": "*string", "Title": "*string", "UserName": "*string", "UserType": "*string", "Website": "*string"},
			"delete-group":                      {"GroupId": "*string", "IdentityStoreId": "*string"},
			"delete-group-membership":           {"IdentityStoreId": "*string", "MembershipId": "*string"},
			"delete-user":                       {"IdentityStoreId": "*string", "UserId": "*string"},
			"describe-group":                    {"GroupId": "*string", "IdentityStoreId": "*string"},
			"describe-group-membership":         {"IdentityStoreId": "*string", "MembershipId": "*string"},
			"describe-user":                     {"Extensions": "[]string", "IdentityStoreId": "*string", "UserId": "*string"},
			"get-group-id":                      {"AlternateIdentifier": "types.AlternateIdentifier", "IdentityStoreId": "*string"},
			"get-group-membership-id":           {"GroupId": "*string", "IdentityStoreId": "*string", "MemberId": "types.MemberId"},
			"get-user-id":                       {"AlternateIdentifier": "types.AlternateIdentifier", "IdentityStoreId": "*string"},
			"is-member-in-groups":               {"GroupIds": "[]string", "IdentityStoreId": "*string", "MemberId": "types.MemberId"},
			"list-group-memberships":            {"GroupId": "*string", "IdentityStoreId": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-group-memberships-for-member": {"IdentityStoreId": "*string", "MaxResults": "*int32", "MemberId": "types.MemberId", "NextToken": "*string"},
			"list-groups":                       {"Filters": "[]types.Filter", "IdentityStoreId": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-users":                        {"Extensions": "[]string", "Filters": "[]types.Filter", "IdentityStoreId": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"update-group":                      {"GroupId": "*string", "IdentityStoreId": "*string", "Operations": "[]types.AttributeOperation"},
			"update-user":                       {"IdentityStoreId": "*string", "Operations": "[]types.AttributeOperation", "UserId": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"create-group":                      {"IdentityStoreId"},
			"create-group-membership":           {"GroupId", "IdentityStoreId", "MemberId"},
			"create-user":                       {"IdentityStoreId"},
			"delete-group":                      {"GroupId", "IdentityStoreId"},
			"delete-group-membership":           {"IdentityStoreId", "MembershipId"},
			"delete-user":                       {"IdentityStoreId", "UserId"},
			"describe-group":                    {"GroupId", "IdentityStoreId"},
			"describe-group-membership":         {"IdentityStoreId", "MembershipId"},
			"describe-user":                     {"IdentityStoreId", "UserId"},
			"get-group-id":                      {"AlternateIdentifier", "IdentityStoreId"},
			"get-group-membership-id":           {"GroupId", "IdentityStoreId", "MemberId"},
			"get-user-id":                       {"AlternateIdentifier", "IdentityStoreId"},
			"is-member-in-groups":               {"GroupIds", "IdentityStoreId", "MemberId"},
			"list-group-memberships":            {"GroupId", "IdentityStoreId"},
			"list-group-memberships-for-member": {"IdentityStoreId", "MemberId"},
			"list-groups":                       {"IdentityStoreId"},
			"list-users":                        {"IdentityStoreId"},
			"update-group":                      {"GroupId", "IdentityStoreId", "Operations"},
			"update-user":                       {"IdentityStoreId", "Operations", "UserId"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("identitystore", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
