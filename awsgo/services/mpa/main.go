package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/mpa/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"cancel-session", "create-approval-team", "create-identity-source", "delete-identity-source", "delete-inactive-approval-team-version", "get-approval-team", "get-identity-source", "get-policy-version", "get-resource-policy", "get-session", "list-approval-teams", "list-identity-sources", "list-policies", "list-policy-versions", "list-resource-policies", "list-sessions", "list-tags-for-resource", "start-active-approval-team-deletion", "tag-resource", "untag-resource", "update-approval-team"},
		OperationSet: map[string]bool{"cancel-session": true, "create-approval-team": true, "create-identity-source": true, "delete-identity-source": true, "delete-inactive-approval-team-version": true, "get-approval-team": true, "get-identity-source": true, "get-policy-version": true, "get-resource-policy": true, "get-session": true, "list-approval-teams": true, "list-identity-sources": true, "list-policies": true, "list-policy-versions": true, "list-resource-policies": true, "list-sessions": true, "list-tags-for-resource": true, "start-active-approval-team-deletion": true, "tag-resource": true, "untag-resource": true, "update-approval-team": true},
		OperationInputs: map[string][]string{
			"cancel-session":                        {"SessionArn"},
			"create-approval-team":                  {"ApprovalStrategy", "Approvers", "ClientToken", "Description", "Name", "Policies", "Tags"},
			"create-identity-source":                {"ClientToken", "IdentitySourceParameters", "Tags"},
			"delete-identity-source":                {"IdentitySourceArn"},
			"delete-inactive-approval-team-version": {"Arn", "VersionId"},
			"get-approval-team":                     {"Arn"},
			"get-identity-source":                   {"IdentitySourceArn"},
			"get-policy-version":                    {"PolicyVersionArn"},
			"get-resource-policy":                   {"PolicyName", "PolicyType", "ResourceArn"},
			"get-session":                           {"SessionArn"},
			"list-approval-teams":                   {"MaxResults", "NextToken"},
			"list-identity-sources":                 {"MaxResults", "NextToken"},
			"list-policies":                         {"MaxResults", "NextToken"},
			"list-policy-versions":                  {"MaxResults", "NextToken", "PolicyArn"},
			"list-resource-policies":                {"MaxResults", "NextToken", "ResourceArn"},
			"list-sessions":                         {"ApprovalTeamArn", "Filters", "MaxResults", "NextToken"},
			"list-tags-for-resource":                {"ResourceArn"},
			"start-active-approval-team-deletion":   {"Arn", "PendingWindowDays"},
			"tag-resource":                          {"ResourceArn", "Tags"},
			"untag-resource":                        {"ResourceArn", "TagKeys"},
			"update-approval-team":                  {"ApprovalStrategy", "Approvers", "Arn", "Description", "UpdateActions"},
		},
		OperationInputTypes: map[string]map[string]string{
			"cancel-session":                        {"SessionArn": "*string"},
			"create-approval-team":                  {"ApprovalStrategy": "types.ApprovalStrategy", "Approvers": "[]types.ApprovalTeamRequestApprover", "ClientToken": "*string", "Description": "*string", "Name": "*string", "Policies": "[]types.PolicyReference", "Tags": "map[string]string"},
			"create-identity-source":                {"ClientToken": "*string", "IdentitySourceParameters": "*types.IdentitySourceParameters", "Tags": "map[string]string"},
			"delete-identity-source":                {"IdentitySourceArn": "*string"},
			"delete-inactive-approval-team-version": {"Arn": "*string", "VersionId": "*string"},
			"get-approval-team":                     {"Arn": "*string"},
			"get-identity-source":                   {"IdentitySourceArn": "*string"},
			"get-policy-version":                    {"PolicyVersionArn": "*string"},
			"get-resource-policy":                   {"PolicyName": "*string", "PolicyType": "types.PolicyType", "ResourceArn": "*string"},
			"get-session":                           {"SessionArn": "*string"},
			"list-approval-teams":                   {"MaxResults": "*int32", "NextToken": "*string"},
			"list-identity-sources":                 {"MaxResults": "*int32", "NextToken": "*string"},
			"list-policies":                         {"MaxResults": "*int32", "NextToken": "*string"},
			"list-policy-versions":                  {"MaxResults": "*int32", "NextToken": "*string", "PolicyArn": "*string"},
			"list-resource-policies":                {"MaxResults": "*int32", "NextToken": "*string", "ResourceArn": "*string"},
			"list-sessions":                         {"ApprovalTeamArn": "*string", "Filters": "[]types.Filter", "MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":                {"ResourceArn": "*string"},
			"start-active-approval-team-deletion":   {"Arn": "*string", "PendingWindowDays": "*int32"},
			"tag-resource":                          {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":                        {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-approval-team":                  {"ApprovalStrategy": "types.ApprovalStrategy", "Approvers": "[]types.ApprovalTeamRequestApprover", "Arn": "*string", "Description": "*string", "UpdateActions": "[]types.UpdateAction"},
		},
		OperationInputRequired: map[string][]string{
			"cancel-session":                        {"SessionArn"},
			"create-approval-team":                  {"ApprovalStrategy", "Approvers", "Description", "Name", "Policies"},
			"create-identity-source":                {"IdentitySourceParameters"},
			"delete-identity-source":                {"IdentitySourceArn"},
			"delete-inactive-approval-team-version": {"Arn", "VersionId"},
			"get-approval-team":                     {"Arn"},
			"get-identity-source":                   {"IdentitySourceArn"},
			"get-policy-version":                    {"PolicyVersionArn"},
			"get-resource-policy":                   {"PolicyName", "PolicyType", "ResourceArn"},
			"get-session":                           {"SessionArn"},
			"list-approval-teams":                   {},
			"list-identity-sources":                 {},
			"list-policies":                         {},
			"list-policy-versions":                  {"PolicyArn"},
			"list-resource-policies":                {"ResourceArn"},
			"list-sessions":                         {"ApprovalTeamArn"},
			"list-tags-for-resource":                {"ResourceArn"},
			"start-active-approval-team-deletion":   {"Arn"},
			"tag-resource":                          {"ResourceArn", "Tags"},
			"untag-resource":                        {"ResourceArn", "TagKeys"},
			"update-approval-team":                  {"Arn"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("mpa", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
