package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/servicequotas/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"associate-service-quota-template", "create-support-case", "delete-service-quota-increase-request-from-template", "disassociate-service-quota-template", "get-association-for-service-quota-template", "get-auto-management-configuration", "get-aws-default-service-quota", "get-quota-utilization-report", "get-requested-service-quota-change", "get-service-quota", "get-service-quota-increase-request-from-template", "list-aws-default-service-quotas", "list-requested-service-quota-change-history", "list-requested-service-quota-change-history-by-quota", "list-service-quota-increase-requests-in-template", "list-service-quotas", "list-services", "list-tags-for-resource", "put-service-quota-increase-request-into-template", "request-service-quota-increase", "start-auto-management", "start-quota-utilization-report", "stop-auto-management", "tag-resource", "untag-resource", "update-auto-management"},
		OperationSet: map[string]bool{"associate-service-quota-template": true, "create-support-case": true, "delete-service-quota-increase-request-from-template": true, "disassociate-service-quota-template": true, "get-association-for-service-quota-template": true, "get-auto-management-configuration": true, "get-aws-default-service-quota": true, "get-quota-utilization-report": true, "get-requested-service-quota-change": true, "get-service-quota": true, "get-service-quota-increase-request-from-template": true, "list-aws-default-service-quotas": true, "list-requested-service-quota-change-history": true, "list-requested-service-quota-change-history-by-quota": true, "list-service-quota-increase-requests-in-template": true, "list-service-quotas": true, "list-services": true, "list-tags-for-resource": true, "put-service-quota-increase-request-into-template": true, "request-service-quota-increase": true, "start-auto-management": true, "start-quota-utilization-report": true, "stop-auto-management": true, "tag-resource": true, "untag-resource": true, "update-auto-management": true},
		OperationInputs: map[string][]string{
			"associate-service-quota-template":                     {},
			"create-support-case":                                  {"RequestId"},
			"delete-service-quota-increase-request-from-template":  {"AwsRegion", "QuotaCode", "ServiceCode"},
			"disassociate-service-quota-template":                  {},
			"get-association-for-service-quota-template":           {},
			"get-auto-management-configuration":                    {},
			"get-aws-default-service-quota":                        {"QuotaCode", "ServiceCode"},
			"get-quota-utilization-report":                         {"MaxResults", "NextToken", "ReportId"},
			"get-requested-service-quota-change":                   {"RequestId"},
			"get-service-quota":                                    {"ContextId", "QuotaCode", "ServiceCode"},
			"get-service-quota-increase-request-from-template":     {"AwsRegion", "QuotaCode", "ServiceCode"},
			"list-aws-default-service-quotas":                      {"MaxResults", "NextToken", "ServiceCode"},
			"list-requested-service-quota-change-history":          {"MaxResults", "NextToken", "QuotaRequestedAtLevel", "ServiceCode", "Status"},
			"list-requested-service-quota-change-history-by-quota": {"MaxResults", "NextToken", "QuotaCode", "QuotaRequestedAtLevel", "ServiceCode", "Status"},
			"list-service-quota-increase-requests-in-template":     {"AwsRegion", "MaxResults", "NextToken", "ServiceCode"},
			"list-service-quotas":                                  {"MaxResults", "NextToken", "QuotaAppliedAtLevel", "QuotaCode", "ServiceCode"},
			"list-services":                                        {"MaxResults", "NextToken"},
			"list-tags-for-resource":                               {"ResourceARN"},
			"put-service-quota-increase-request-into-template":     {"AwsRegion", "DesiredValue", "QuotaCode", "ServiceCode"},
			"request-service-quota-increase":                       {"ContextId", "DesiredValue", "QuotaCode", "ServiceCode", "SupportCaseAllowed"},
			"start-auto-management":                                {"ExclusionList", "NotificationArn", "OptInLevel", "OptInType"},
			"start-quota-utilization-report":                       {},
			"stop-auto-management":                                 {},
			"tag-resource":                                         {"ResourceARN", "Tags"},
			"untag-resource":                                       {"ResourceARN", "TagKeys"},
			"update-auto-management":                               {"ExclusionList", "NotificationArn", "OptInType"},
		},
		OperationInputTypes: map[string]map[string]string{
			"associate-service-quota-template":                     {},
			"create-support-case":                                  {"RequestId": "*string"},
			"delete-service-quota-increase-request-from-template":  {"AwsRegion": "*string", "QuotaCode": "*string", "ServiceCode": "*string"},
			"disassociate-service-quota-template":                  {},
			"get-association-for-service-quota-template":           {},
			"get-auto-management-configuration":                    {},
			"get-aws-default-service-quota":                        {"QuotaCode": "*string", "ServiceCode": "*string"},
			"get-quota-utilization-report":                         {"MaxResults": "*int32", "NextToken": "*string", "ReportId": "*string"},
			"get-requested-service-quota-change":                   {"RequestId": "*string"},
			"get-service-quota":                                    {"ContextId": "*string", "QuotaCode": "*string", "ServiceCode": "*string"},
			"get-service-quota-increase-request-from-template":     {"AwsRegion": "*string", "QuotaCode": "*string", "ServiceCode": "*string"},
			"list-aws-default-service-quotas":                      {"MaxResults": "*int32", "NextToken": "*string", "ServiceCode": "*string"},
			"list-requested-service-quota-change-history":          {"MaxResults": "*int32", "NextToken": "*string", "QuotaRequestedAtLevel": "types.AppliedLevelEnum", "ServiceCode": "*string", "Status": "types.RequestStatus"},
			"list-requested-service-quota-change-history-by-quota": {"MaxResults": "*int32", "NextToken": "*string", "QuotaCode": "*string", "QuotaRequestedAtLevel": "types.AppliedLevelEnum", "ServiceCode": "*string", "Status": "types.RequestStatus"},
			"list-service-quota-increase-requests-in-template":     {"AwsRegion": "*string", "MaxResults": "*int32", "NextToken": "*string", "ServiceCode": "*string"},
			"list-service-quotas":                                  {"MaxResults": "*int32", "NextToken": "*string", "QuotaAppliedAtLevel": "types.AppliedLevelEnum", "QuotaCode": "*string", "ServiceCode": "*string"},
			"list-services":                                        {"MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":                               {"ResourceARN": "*string"},
			"put-service-quota-increase-request-into-template":     {"AwsRegion": "*string", "DesiredValue": "*float64", "QuotaCode": "*string", "ServiceCode": "*string"},
			"request-service-quota-increase":                       {"ContextId": "*string", "DesiredValue": "*float64", "QuotaCode": "*string", "ServiceCode": "*string", "SupportCaseAllowed": "*bool"},
			"start-auto-management":                                {"ExclusionList": "map[string][]string", "NotificationArn": "*string", "OptInLevel": "types.OptInLevel", "OptInType": "types.OptInType"},
			"start-quota-utilization-report":                       {},
			"stop-auto-management":                                 {},
			"tag-resource":                                         {"ResourceARN": "*string", "Tags": "[]types.Tag"},
			"untag-resource":                                       {"ResourceARN": "*string", "TagKeys": "[]string"},
			"update-auto-management":                               {"ExclusionList": "map[string][]string", "NotificationArn": "*string", "OptInType": "types.OptInType"},
		},
		OperationInputRequired: map[string][]string{
			"associate-service-quota-template":                     {},
			"create-support-case":                                  {"RequestId"},
			"delete-service-quota-increase-request-from-template":  {"AwsRegion", "QuotaCode", "ServiceCode"},
			"disassociate-service-quota-template":                  {},
			"get-association-for-service-quota-template":           {},
			"get-auto-management-configuration":                    {},
			"get-aws-default-service-quota":                        {"QuotaCode", "ServiceCode"},
			"get-quota-utilization-report":                         {"ReportId"},
			"get-requested-service-quota-change":                   {"RequestId"},
			"get-service-quota":                                    {"QuotaCode", "ServiceCode"},
			"get-service-quota-increase-request-from-template":     {"AwsRegion", "QuotaCode", "ServiceCode"},
			"list-aws-default-service-quotas":                      {"ServiceCode"},
			"list-requested-service-quota-change-history":          {},
			"list-requested-service-quota-change-history-by-quota": {"QuotaCode", "ServiceCode"},
			"list-service-quota-increase-requests-in-template":     {},
			"list-service-quotas":                                  {"ServiceCode"},
			"list-services":                                        {},
			"list-tags-for-resource":                               {"ResourceARN"},
			"put-service-quota-increase-request-into-template":     {"AwsRegion", "DesiredValue", "QuotaCode", "ServiceCode"},
			"request-service-quota-increase":                       {"DesiredValue", "QuotaCode", "ServiceCode"},
			"start-auto-management":                                {"OptInLevel", "OptInType"},
			"start-quota-utilization-report":                       {},
			"stop-auto-management":                                 {},
			"tag-resource":                                         {"ResourceARN", "Tags"},
			"untag-resource":                                       {"ResourceARN", "TagKeys"},
			"update-auto-management":                               {},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("servicequotas", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
