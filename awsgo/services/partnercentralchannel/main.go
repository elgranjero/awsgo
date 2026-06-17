package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/partnercentralchannel/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"accept-channel-handshake", "cancel-channel-handshake", "create-channel-handshake", "create-program-management-account", "create-relationship", "delete-program-management-account", "delete-relationship", "get-relationship", "list-channel-handshakes", "list-program-management-accounts", "list-relationships", "list-tags-for-resource", "reject-channel-handshake", "tag-resource", "untag-resource", "update-program-management-account", "update-relationship"},
		OperationSet: map[string]bool{"accept-channel-handshake": true, "cancel-channel-handshake": true, "create-channel-handshake": true, "create-program-management-account": true, "create-relationship": true, "delete-program-management-account": true, "delete-relationship": true, "get-relationship": true, "list-channel-handshakes": true, "list-program-management-accounts": true, "list-relationships": true, "list-tags-for-resource": true, "reject-channel-handshake": true, "tag-resource": true, "untag-resource": true, "update-program-management-account": true, "update-relationship": true},
		OperationInputs: map[string][]string{
			"accept-channel-handshake":          {"Catalog", "Identifier"},
			"cancel-channel-handshake":          {"Catalog", "Identifier"},
			"create-channel-handshake":          {"AssociatedResourceIdentifier", "Catalog", "ClientToken", "HandshakeType", "Payload", "Tags"},
			"create-program-management-account": {"AccountId", "Catalog", "ClientToken", "DisplayName", "Program", "Tags"},
			"create-relationship":               {"AssociatedAccountId", "AssociationType", "Catalog", "ClientToken", "DisplayName", "ProgramManagementAccountIdentifier", "RequestedSupportPlan", "ResaleAccountModel", "Sector", "Tags"},
			"delete-program-management-account": {"Catalog", "ClientToken", "Identifier"},
			"delete-relationship":               {"Catalog", "ClientToken", "Identifier", "ProgramManagementAccountIdentifier"},
			"get-relationship":                  {"Catalog", "Identifier", "ProgramManagementAccountIdentifier"},
			"list-channel-handshakes":           {"AssociatedResourceIdentifiers", "Catalog", "HandshakeType", "HandshakeTypeFilters", "HandshakeTypeSort", "MaxResults", "NextToken", "ParticipantType", "Statuses"},
			"list-program-management-accounts":  {"AccountIds", "Catalog", "DisplayNames", "MaxResults", "NextToken", "Programs", "Sort", "Statuses"},
			"list-relationships":                {"AssociatedAccountIds", "AssociationTypes", "Catalog", "DisplayNames", "MaxResults", "NextToken", "ProgramManagementAccountIdentifiers", "Sort"},
			"list-tags-for-resource":            {"ResourceArn"},
			"reject-channel-handshake":          {"Catalog", "Identifier"},
			"tag-resource":                      {"ResourceArn", "Tags"},
			"untag-resource":                    {"ResourceArn", "TagKeys"},
			"update-program-management-account": {"Catalog", "DisplayName", "Identifier", "Revision"},
			"update-relationship":               {"Catalog", "DisplayName", "Identifier", "ProgramManagementAccountIdentifier", "RequestedSupportPlan", "Revision"},
		},
		OperationInputTypes: map[string]map[string]string{
			"accept-channel-handshake":          {"Catalog": "*string", "Identifier": "*string"},
			"cancel-channel-handshake":          {"Catalog": "*string", "Identifier": "*string"},
			"create-channel-handshake":          {"AssociatedResourceIdentifier": "*string", "Catalog": "*string", "ClientToken": "*string", "HandshakeType": "types.HandshakeType", "Payload": "types.ChannelHandshakePayload", "Tags": "[]types.Tag"},
			"create-program-management-account": {"AccountId": "*string", "Catalog": "*string", "ClientToken": "*string", "DisplayName": "*string", "Program": "types.Program", "Tags": "[]types.Tag"},
			"create-relationship":               {"AssociatedAccountId": "*string", "AssociationType": "types.AssociationType", "Catalog": "*string", "ClientToken": "*string", "DisplayName": "*string", "ProgramManagementAccountIdentifier": "*string", "RequestedSupportPlan": "types.SupportPlan", "ResaleAccountModel": "types.ResaleAccountModel", "Sector": "types.Sector", "Tags": "[]types.Tag"},
			"delete-program-management-account": {"Catalog": "*string", "ClientToken": "*string", "Identifier": "*string"},
			"delete-relationship":               {"Catalog": "*string", "ClientToken": "*string", "Identifier": "*string", "ProgramManagementAccountIdentifier": "*string"},
			"get-relationship":                  {"Catalog": "*string", "Identifier": "*string", "ProgramManagementAccountIdentifier": "*string"},
			"list-channel-handshakes":           {"AssociatedResourceIdentifiers": "[]string", "Catalog": "*string", "HandshakeType": "types.HandshakeType", "HandshakeTypeFilters": "types.ListChannelHandshakesTypeFilters", "HandshakeTypeSort": "types.ListChannelHandshakesTypeSort", "MaxResults": "*int32", "NextToken": "*string", "ParticipantType": "types.ParticipantType", "Statuses": "[]types.HandshakeStatus"},
			"list-program-management-accounts":  {"AccountIds": "[]string", "Catalog": "*string", "DisplayNames": "[]string", "MaxResults": "*int32", "NextToken": "*string", "Programs": "[]types.Program", "Sort": "*types.ListProgramManagementAccountsSortBase", "Statuses": "[]types.ProgramManagementAccountStatus"},
			"list-relationships":                {"AssociatedAccountIds": "[]string", "AssociationTypes": "[]types.AssociationType", "Catalog": "*string", "DisplayNames": "[]string", "MaxResults": "*int32", "NextToken": "*string", "ProgramManagementAccountIdentifiers": "[]string", "Sort": "*types.ListRelationshipsSortBase"},
			"list-tags-for-resource":            {"ResourceArn": "*string"},
			"reject-channel-handshake":          {"Catalog": "*string", "Identifier": "*string"},
			"tag-resource":                      {"ResourceArn": "*string", "Tags": "[]types.Tag"},
			"untag-resource":                    {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-program-management-account": {"Catalog": "*string", "DisplayName": "*string", "Identifier": "*string", "Revision": "*string"},
			"update-relationship":               {"Catalog": "*string", "DisplayName": "*string", "Identifier": "*string", "ProgramManagementAccountIdentifier": "*string", "RequestedSupportPlan": "types.SupportPlan", "Revision": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"accept-channel-handshake":          {"Catalog", "Identifier"},
			"cancel-channel-handshake":          {"Catalog", "Identifier"},
			"create-channel-handshake":          {"AssociatedResourceIdentifier", "Catalog", "HandshakeType"},
			"create-program-management-account": {"AccountId", "Catalog", "DisplayName", "Program"},
			"create-relationship":               {"AssociatedAccountId", "AssociationType", "Catalog", "DisplayName", "ProgramManagementAccountIdentifier", "Sector"},
			"delete-program-management-account": {"Catalog", "Identifier"},
			"delete-relationship":               {"Catalog", "Identifier", "ProgramManagementAccountIdentifier"},
			"get-relationship":                  {"Catalog", "Identifier", "ProgramManagementAccountIdentifier"},
			"list-channel-handshakes":           {"Catalog", "HandshakeType", "ParticipantType"},
			"list-program-management-accounts":  {"Catalog"},
			"list-relationships":                {"Catalog"},
			"list-tags-for-resource":            {"ResourceArn"},
			"reject-channel-handshake":          {"Catalog", "Identifier"},
			"tag-resource":                      {"ResourceArn", "Tags"},
			"untag-resource":                    {"ResourceArn", "TagKeys"},
			"update-program-management-account": {"Catalog", "Identifier"},
			"update-relationship":               {"Catalog", "Identifier", "ProgramManagementAccountIdentifier"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("partnercentralchannel", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
