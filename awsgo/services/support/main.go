package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/support/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"add-attachments-to-set", "add-communication-to-case", "create-case", "describe-attachment", "describe-cases", "describe-communications", "describe-create-case-options", "describe-services", "describe-severity-levels", "describe-supported-languages", "describe-trusted-advisor-check-refresh-statuses", "describe-trusted-advisor-check-result", "describe-trusted-advisor-check-summaries", "describe-trusted-advisor-checks", "refresh-trusted-advisor-check", "resolve-case"},
		OperationSet: map[string]bool{"add-attachments-to-set": true, "add-communication-to-case": true, "create-case": true, "describe-attachment": true, "describe-cases": true, "describe-communications": true, "describe-create-case-options": true, "describe-services": true, "describe-severity-levels": true, "describe-supported-languages": true, "describe-trusted-advisor-check-refresh-statuses": true, "describe-trusted-advisor-check-result": true, "describe-trusted-advisor-check-summaries": true, "describe-trusted-advisor-checks": true, "refresh-trusted-advisor-check": true, "resolve-case": true},
		OperationInputs: map[string][]string{
			"add-attachments-to-set":                          {"AttachmentSetId", "Attachments"},
			"add-communication-to-case":                       {"AttachmentSetId", "CaseId", "CcEmailAddresses", "CommunicationBody"},
			"create-case":                                     {"AttachmentSetId", "CategoryCode", "CcEmailAddresses", "CommunicationBody", "IssueType", "Language", "ServiceCode", "SeverityCode", "Subject"},
			"describe-attachment":                             {"AttachmentId"},
			"describe-cases":                                  {"AfterTime", "BeforeTime", "CaseIdList", "DisplayId", "IncludeCommunications", "IncludeResolvedCases", "Language", "MaxResults", "NextToken"},
			"describe-communications":                         {"AfterTime", "BeforeTime", "CaseId", "MaxResults", "NextToken"},
			"describe-create-case-options":                    {"CategoryCode", "IssueType", "Language", "ServiceCode"},
			"describe-services":                               {"Language", "ServiceCodeList"},
			"describe-severity-levels":                        {"Language"},
			"describe-supported-languages":                    {"CategoryCode", "IssueType", "ServiceCode"},
			"describe-trusted-advisor-check-refresh-statuses": {"CheckIds"},
			"describe-trusted-advisor-check-result":           {"CheckId", "Language"},
			"describe-trusted-advisor-check-summaries":        {"CheckIds"},
			"describe-trusted-advisor-checks":                 {"Language"},
			"refresh-trusted-advisor-check":                   {"CheckId"},
			"resolve-case":                                    {"CaseId"},
		},
		OperationInputTypes: map[string]map[string]string{
			"add-attachments-to-set":                          {"AttachmentSetId": "*string", "Attachments": "[]types.Attachment"},
			"add-communication-to-case":                       {"AttachmentSetId": "*string", "CaseId": "*string", "CcEmailAddresses": "[]string", "CommunicationBody": "*string"},
			"create-case":                                     {"AttachmentSetId": "*string", "CategoryCode": "*string", "CcEmailAddresses": "[]string", "CommunicationBody": "*string", "IssueType": "*string", "Language": "*string", "ServiceCode": "*string", "SeverityCode": "*string", "Subject": "*string"},
			"describe-attachment":                             {"AttachmentId": "*string"},
			"describe-cases":                                  {"AfterTime": "*string", "BeforeTime": "*string", "CaseIdList": "[]string", "DisplayId": "*string", "IncludeCommunications": "*bool", "IncludeResolvedCases": "bool", "Language": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"describe-communications":                         {"AfterTime": "*string", "BeforeTime": "*string", "CaseId": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"describe-create-case-options":                    {"CategoryCode": "*string", "IssueType": "*string", "Language": "*string", "ServiceCode": "*string"},
			"describe-services":                               {"Language": "*string", "ServiceCodeList": "[]string"},
			"describe-severity-levels":                        {"Language": "*string"},
			"describe-supported-languages":                    {"CategoryCode": "*string", "IssueType": "*string", "ServiceCode": "*string"},
			"describe-trusted-advisor-check-refresh-statuses": {"CheckIds": "[]*string"},
			"describe-trusted-advisor-check-result":           {"CheckId": "*string", "Language": "*string"},
			"describe-trusted-advisor-check-summaries":        {"CheckIds": "[]*string"},
			"describe-trusted-advisor-checks":                 {"Language": "*string"},
			"refresh-trusted-advisor-check":                   {"CheckId": "*string"},
			"resolve-case":                                    {"CaseId": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"add-attachments-to-set":                          {"Attachments"},
			"add-communication-to-case":                       {"CommunicationBody"},
			"create-case":                                     {"CommunicationBody", "Subject"},
			"describe-attachment":                             {"AttachmentId"},
			"describe-cases":                                  {},
			"describe-communications":                         {"CaseId"},
			"describe-create-case-options":                    {"CategoryCode", "IssueType", "Language", "ServiceCode"},
			"describe-services":                               {},
			"describe-severity-levels":                        {},
			"describe-supported-languages":                    {"CategoryCode", "IssueType", "ServiceCode"},
			"describe-trusted-advisor-check-refresh-statuses": {"CheckIds"},
			"describe-trusted-advisor-check-result":           {"CheckId"},
			"describe-trusted-advisor-check-summaries":        {"CheckIds"},
			"describe-trusted-advisor-checks":                 {"Language"},
			"refresh-trusted-advisor-check":                   {"CheckId"},
			"resolve-case":                                    {},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("support", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
