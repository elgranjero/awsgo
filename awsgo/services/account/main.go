package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/account/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"accept-primary-email-update", "delete-alternate-contact", "disable-region", "enable-region", "get-account-information", "get-alternate-contact", "get-contact-information", "get-gov-cloud-account-information", "get-primary-email", "get-region-opt-status", "list-regions", "put-account-name", "put-alternate-contact", "put-contact-information", "start-primary-email-update"},
		OperationSet: map[string]bool{"accept-primary-email-update": true, "delete-alternate-contact": true, "disable-region": true, "enable-region": true, "get-account-information": true, "get-alternate-contact": true, "get-contact-information": true, "get-gov-cloud-account-information": true, "get-primary-email": true, "get-region-opt-status": true, "list-regions": true, "put-account-name": true, "put-alternate-contact": true, "put-contact-information": true, "start-primary-email-update": true},
		OperationInputs: map[string][]string{
			"accept-primary-email-update":       {"AccountId", "Otp", "PrimaryEmail"},
			"delete-alternate-contact":          {"AccountId", "AlternateContactType"},
			"disable-region":                    {"AccountId", "RegionName"},
			"enable-region":                     {"AccountId", "RegionName"},
			"get-account-information":           {"AccountId"},
			"get-alternate-contact":             {"AccountId", "AlternateContactType"},
			"get-contact-information":           {"AccountId"},
			"get-gov-cloud-account-information": {"StandardAccountId"},
			"get-primary-email":                 {"AccountId"},
			"get-region-opt-status":             {"AccountId", "RegionName"},
			"list-regions":                      {"AccountId", "MaxResults", "NextToken", "RegionOptStatusContains"},
			"put-account-name":                  {"AccountId", "AccountName"},
			"put-alternate-contact":             {"AccountId", "AlternateContactType", "EmailAddress", "Name", "PhoneNumber", "Title"},
			"put-contact-information":           {"AccountId", "ContactInformation"},
			"start-primary-email-update":        {"AccountId", "PrimaryEmail"},
		},
		OperationInputTypes: map[string]map[string]string{
			"accept-primary-email-update":       {"AccountId": "*string", "Otp": "*string", "PrimaryEmail": "*string"},
			"delete-alternate-contact":          {"AccountId": "*string", "AlternateContactType": "types.AlternateContactType"},
			"disable-region":                    {"AccountId": "*string", "RegionName": "*string"},
			"enable-region":                     {"AccountId": "*string", "RegionName": "*string"},
			"get-account-information":           {"AccountId": "*string"},
			"get-alternate-contact":             {"AccountId": "*string", "AlternateContactType": "types.AlternateContactType"},
			"get-contact-information":           {"AccountId": "*string"},
			"get-gov-cloud-account-information": {"StandardAccountId": "*string"},
			"get-primary-email":                 {"AccountId": "*string"},
			"get-region-opt-status":             {"AccountId": "*string", "RegionName": "*string"},
			"list-regions":                      {"AccountId": "*string", "MaxResults": "*int32", "NextToken": "*string", "RegionOptStatusContains": "[]types.RegionOptStatus"},
			"put-account-name":                  {"AccountId": "*string", "AccountName": "*string"},
			"put-alternate-contact":             {"AccountId": "*string", "AlternateContactType": "types.AlternateContactType", "EmailAddress": "*string", "Name": "*string", "PhoneNumber": "*string", "Title": "*string"},
			"put-contact-information":           {"AccountId": "*string", "ContactInformation": "*types.ContactInformation"},
			"start-primary-email-update":        {"AccountId": "*string", "PrimaryEmail": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"accept-primary-email-update":       {"AccountId", "Otp", "PrimaryEmail"},
			"delete-alternate-contact":          {"AlternateContactType"},
			"disable-region":                    {"RegionName"},
			"enable-region":                     {"RegionName"},
			"get-account-information":           {},
			"get-alternate-contact":             {"AlternateContactType"},
			"get-contact-information":           {},
			"get-gov-cloud-account-information": {},
			"get-primary-email":                 {"AccountId"},
			"get-region-opt-status":             {"RegionName"},
			"list-regions":                      {},
			"put-account-name":                  {"AccountName"},
			"put-alternate-contact":             {"AlternateContactType", "EmailAddress", "Name", "PhoneNumber", "Title"},
			"put-contact-information":           {"ContactInformation"},
			"start-primary-email-update":        {"AccountId", "PrimaryEmail"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("account", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
