package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/notificationscontacts/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"activate-email-contact", "create-email-contact", "delete-email-contact", "get-email-contact", "list-email-contacts", "list-tags-for-resource", "send-activation-code", "tag-resource", "untag-resource"},
		OperationSet: map[string]bool{"activate-email-contact": true, "create-email-contact": true, "delete-email-contact": true, "get-email-contact": true, "list-email-contacts": true, "list-tags-for-resource": true, "send-activation-code": true, "tag-resource": true, "untag-resource": true},
		OperationInputs: map[string][]string{
			"activate-email-contact": {"Arn", "Code"},
			"create-email-contact":   {"EmailAddress", "Name", "Tags"},
			"delete-email-contact":   {"Arn"},
			"get-email-contact":      {"Arn"},
			"list-email-contacts":    {"MaxResults", "NextToken"},
			"list-tags-for-resource": {"Arn"},
			"send-activation-code":   {"Arn"},
			"tag-resource":           {"Arn", "Tags"},
			"untag-resource":         {"Arn", "TagKeys"},
		},
		OperationInputTypes: map[string]map[string]string{
			"activate-email-contact": {"Arn": "*string", "Code": "*string"},
			"create-email-contact":   {"EmailAddress": "*string", "Name": "*string", "Tags": "map[string]string"},
			"delete-email-contact":   {"Arn": "*string"},
			"get-email-contact":      {"Arn": "*string"},
			"list-email-contacts":    {"MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource": {"Arn": "*string"},
			"send-activation-code":   {"Arn": "*string"},
			"tag-resource":           {"Arn": "*string", "Tags": "map[string]string"},
			"untag-resource":         {"Arn": "*string", "TagKeys": "[]string"},
		},
		OperationInputRequired: map[string][]string{
			"activate-email-contact": {"Arn", "Code"},
			"create-email-contact":   {"EmailAddress", "Name"},
			"delete-email-contact":   {"Arn"},
			"get-email-contact":      {"Arn"},
			"list-email-contacts":    {},
			"list-tags-for-resource": {"Arn"},
			"send-activation-code":   {"Arn"},
			"tag-resource":           {"Arn", "Tags"},
			"untag-resource":         {"Arn", "TagKeys"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("notificationscontacts", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
