package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/billing/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"associate-source-views", "create-billing-view", "delete-billing-view", "disassociate-source-views", "get-billing-view", "get-resource-policy", "list-billing-views", "list-source-views-for-billing-view", "list-tags-for-resource", "tag-resource", "untag-resource", "update-billing-view"},
		OperationSet: map[string]bool{"associate-source-views": true, "create-billing-view": true, "delete-billing-view": true, "disassociate-source-views": true, "get-billing-view": true, "get-resource-policy": true, "list-billing-views": true, "list-source-views-for-billing-view": true, "list-tags-for-resource": true, "tag-resource": true, "untag-resource": true, "update-billing-view": true},
		OperationInputs: map[string][]string{
			"associate-source-views":             {"Arn", "SourceViews"},
			"create-billing-view":                {"ClientToken", "DataFilterExpression", "Description", "Name", "ResourceTags", "SourceViews"},
			"delete-billing-view":                {"Arn", "Force"},
			"disassociate-source-views":          {"Arn", "SourceViews"},
			"get-billing-view":                   {"Arn"},
			"get-resource-policy":                {"ResourceArn"},
			"list-billing-views":                 {"ActiveTimeRange", "Arns", "BillingViewTypes", "MaxResults", "Names", "NextToken", "OwnerAccountId", "SourceAccountId"},
			"list-source-views-for-billing-view": {"Arn", "MaxResults", "NextToken"},
			"list-tags-for-resource":             {"ResourceArn"},
			"tag-resource":                       {"ResourceArn", "ResourceTags"},
			"untag-resource":                     {"ResourceArn", "ResourceTagKeys"},
			"update-billing-view":                {"Arn", "DataFilterExpression", "Description", "Name"},
		},
		OperationInputTypes: map[string]map[string]string{
			"associate-source-views":             {"Arn": "*string", "SourceViews": "[]string"},
			"create-billing-view":                {"ClientToken": "*string", "DataFilterExpression": "*types.Expression", "Description": "*string", "Name": "*string", "ResourceTags": "[]types.ResourceTag", "SourceViews": "[]string"},
			"delete-billing-view":                {"Arn": "*string", "Force": "bool"},
			"disassociate-source-views":          {"Arn": "*string", "SourceViews": "[]string"},
			"get-billing-view":                   {"Arn": "*string"},
			"get-resource-policy":                {"ResourceArn": "*string"},
			"list-billing-views":                 {"ActiveTimeRange": "*types.ActiveTimeRange", "Arns": "[]string", "BillingViewTypes": "[]types.BillingViewType", "MaxResults": "*int32", "Names": "[]types.StringSearch", "NextToken": "*string", "OwnerAccountId": "*string", "SourceAccountId": "*string"},
			"list-source-views-for-billing-view": {"Arn": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":             {"ResourceArn": "*string"},
			"tag-resource":                       {"ResourceArn": "*string", "ResourceTags": "[]types.ResourceTag"},
			"untag-resource":                     {"ResourceArn": "*string", "ResourceTagKeys": "[]string"},
			"update-billing-view":                {"Arn": "*string", "DataFilterExpression": "*types.Expression", "Description": "*string", "Name": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"associate-source-views":             {"Arn", "SourceViews"},
			"create-billing-view":                {"Name", "SourceViews"},
			"delete-billing-view":                {"Arn"},
			"disassociate-source-views":          {"Arn", "SourceViews"},
			"get-billing-view":                   {"Arn"},
			"get-resource-policy":                {"ResourceArn"},
			"list-billing-views":                 {},
			"list-source-views-for-billing-view": {"Arn"},
			"list-tags-for-resource":             {"ResourceArn"},
			"tag-resource":                       {"ResourceArn", "ResourceTags"},
			"untag-resource":                     {"ResourceArn", "ResourceTagKeys"},
			"update-billing-view":                {"Arn"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("billing", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
