package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/servicequotas"
)

var fields_associate_service_quota_template = []leanruntime.Field{}

var fields_create_support_case = []leanruntime.Field{
	{Name: "RequestId", Flag: "request-id", Type: "*string", Required: true},
}

var fields_delete_service_quota_increase_request_from_template = []leanruntime.Field{
	{Name: "AwsRegion", Flag: "aws-region", Type: "*string", Required: true},
	{Name: "QuotaCode", Flag: "quota-code", Type: "*string", Required: true},
	{Name: "ServiceCode", Flag: "service-code", Type: "*string", Required: true},
}

var fields_disassociate_service_quota_template = []leanruntime.Field{}

var fields_get_association_for_service_quota_template = []leanruntime.Field{}

var fields_get_auto_management_configuration = []leanruntime.Field{}

var fields_get_aws_default_service_quota = []leanruntime.Field{
	{Name: "QuotaCode", Flag: "quota-code", Type: "*string", Required: true},
	{Name: "ServiceCode", Flag: "service-code", Type: "*string", Required: true},
}

var fields_get_quota_utilization_report = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ReportId", Flag: "report-id", Type: "*string", Required: true},
}

var fields_get_requested_service_quota_change = []leanruntime.Field{
	{Name: "RequestId", Flag: "request-id", Type: "*string", Required: true},
}

var fields_get_service_quota = []leanruntime.Field{
	{Name: "ContextId", Flag: "context-id", Type: "*string", Required: false},
	{Name: "QuotaCode", Flag: "quota-code", Type: "*string", Required: true},
	{Name: "ServiceCode", Flag: "service-code", Type: "*string", Required: true},
}

var fields_get_service_quota_increase_request_from_template = []leanruntime.Field{
	{Name: "AwsRegion", Flag: "aws-region", Type: "*string", Required: true},
	{Name: "QuotaCode", Flag: "quota-code", Type: "*string", Required: true},
	{Name: "ServiceCode", Flag: "service-code", Type: "*string", Required: true},
}

var fields_list_aws_default_service_quotas = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ServiceCode", Flag: "service-code", Type: "*string", Required: true},
}

var fields_list_requested_service_quota_change_history = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "QuotaRequestedAtLevel", Flag: "quota-requested-at-level", Type: "types.AppliedLevelEnum", Required: false},
	{Name: "ServiceCode", Flag: "service-code", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.RequestStatus", Required: false},
}

var fields_list_requested_service_quota_change_history_by_quota = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "QuotaCode", Flag: "quota-code", Type: "*string", Required: true},
	{Name: "QuotaRequestedAtLevel", Flag: "quota-requested-at-level", Type: "types.AppliedLevelEnum", Required: false},
	{Name: "ServiceCode", Flag: "service-code", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "types.RequestStatus", Required: false},
}

var fields_list_service_quota_increase_requests_in_template = []leanruntime.Field{
	{Name: "AwsRegion", Flag: "aws-region", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ServiceCode", Flag: "service-code", Type: "*string", Required: false},
}

var fields_list_service_quotas = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "QuotaAppliedAtLevel", Flag: "quota-applied-at-level", Type: "types.AppliedLevelEnum", Required: false},
	{Name: "QuotaCode", Flag: "quota-code", Type: "*string", Required: false},
	{Name: "ServiceCode", Flag: "service-code", Type: "*string", Required: true},
}

var fields_list_services = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_put_service_quota_increase_request_into_template = []leanruntime.Field{
	{Name: "AwsRegion", Flag: "aws-region", Type: "*string", Required: true},
	{Name: "DesiredValue", Flag: "desired-value", Type: "*float64", Required: true},
	{Name: "QuotaCode", Flag: "quota-code", Type: "*string", Required: true},
	{Name: "ServiceCode", Flag: "service-code", Type: "*string", Required: true},
}

var fields_request_service_quota_increase = []leanruntime.Field{
	{Name: "ContextId", Flag: "context-id", Type: "*string", Required: false},
	{Name: "DesiredValue", Flag: "desired-value", Type: "*float64", Required: true},
	{Name: "QuotaCode", Flag: "quota-code", Type: "*string", Required: true},
	{Name: "ServiceCode", Flag: "service-code", Type: "*string", Required: true},
	{Name: "SupportCaseAllowed", Flag: "support-case-allowed", Type: "*bool", Required: false},
}

var fields_start_auto_management = []leanruntime.Field{
	{Name: "ExclusionList", Flag: "exclusion-list", Type: "map[string][]string", Required: false},
	{Name: "NotificationArn", Flag: "notification-arn", Type: "*string", Required: false},
	{Name: "OptInLevel", Flag: "opt-in-level", Type: "types.OptInLevel", Required: true},
	{Name: "OptInType", Flag: "opt-in-type", Type: "types.OptInType", Required: true},
}

var fields_start_quota_utilization_report = []leanruntime.Field{}

var fields_stop_auto_management = []leanruntime.Field{}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_auto_management = []leanruntime.Field{
	{Name: "ExclusionList", Flag: "exclusion-list", Type: "map[string][]string", Required: false},
	{Name: "NotificationArn", Flag: "notification-arn", Type: "*string", Required: false},
	{Name: "OptInType", Flag: "opt-in-type", Type: "types.OptInType", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-service-quota-template": {
			Name:   "associate-service-quota-template",
			Fields: fields_associate_service_quota_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateServiceQuotaTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_service_quota_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateServiceQuotaTemplate(ctx, input)
			},
		},
		"create-support-case": {
			Name:   "create-support-case",
			Fields: fields_create_support_case,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSupportCaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_support_case, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSupportCase(ctx, input)
			},
		},
		"delete-service-quota-increase-request-from-template": {
			Name:   "delete-service-quota-increase-request-from-template",
			Fields: fields_delete_service_quota_increase_request_from_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteServiceQuotaIncreaseRequestFromTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_service_quota_increase_request_from_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteServiceQuotaIncreaseRequestFromTemplate(ctx, input)
			},
		},
		"disassociate-service-quota-template": {
			Name:   "disassociate-service-quota-template",
			Fields: fields_disassociate_service_quota_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateServiceQuotaTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_service_quota_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateServiceQuotaTemplate(ctx, input)
			},
		},
		"get-association-for-service-quota-template": {
			Name:   "get-association-for-service-quota-template",
			Fields: fields_get_association_for_service_quota_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAssociationForServiceQuotaTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_association_for_service_quota_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAssociationForServiceQuotaTemplate(ctx, input)
			},
		},
		"get-auto-management-configuration": {
			Name:   "get-auto-management-configuration",
			Fields: fields_get_auto_management_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAutoManagementConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_auto_management_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAutoManagementConfiguration(ctx, input)
			},
		},
		"get-aws-default-service-quota": {
			Name:   "get-aws-default-service-quota",
			Fields: fields_get_aws_default_service_quota,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAWSDefaultServiceQuotaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_aws_default_service_quota, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAWSDefaultServiceQuota(ctx, input)
			},
		},
		"get-quota-utilization-report": {
			Name:   "get-quota-utilization-report",
			Fields: fields_get_quota_utilization_report,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetQuotaUtilizationReportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_quota_utilization_report, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetQuotaUtilizationReport(ctx, input)
			},
		},
		"get-requested-service-quota-change": {
			Name:   "get-requested-service-quota-change",
			Fields: fields_get_requested_service_quota_change,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRequestedServiceQuotaChangeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_requested_service_quota_change, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRequestedServiceQuotaChange(ctx, input)
			},
		},
		"get-service-quota": {
			Name:   "get-service-quota",
			Fields: fields_get_service_quota,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetServiceQuotaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_service_quota, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetServiceQuota(ctx, input)
			},
		},
		"get-service-quota-increase-request-from-template": {
			Name:   "get-service-quota-increase-request-from-template",
			Fields: fields_get_service_quota_increase_request_from_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetServiceQuotaIncreaseRequestFromTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_service_quota_increase_request_from_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetServiceQuotaIncreaseRequestFromTemplate(ctx, input)
			},
		},
		"list-aws-default-service-quotas": {
			Name:   "list-aws-default-service-quotas",
			Fields: fields_list_aws_default_service_quotas,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAWSDefaultServiceQuotasInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_aws_default_service_quotas, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAWSDefaultServiceQuotas(ctx, input)
				}
				var results []*svc.ListAWSDefaultServiceQuotasOutput
				p := svc.NewListAWSDefaultServiceQuotasPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-requested-service-quota-change-history": {
			Name:   "list-requested-service-quota-change-history",
			Fields: fields_list_requested_service_quota_change_history,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRequestedServiceQuotaChangeHistoryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_requested_service_quota_change_history, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRequestedServiceQuotaChangeHistory(ctx, input)
				}
				var results []*svc.ListRequestedServiceQuotaChangeHistoryOutput
				p := svc.NewListRequestedServiceQuotaChangeHistoryPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-requested-service-quota-change-history-by-quota": {
			Name:   "list-requested-service-quota-change-history-by-quota",
			Fields: fields_list_requested_service_quota_change_history_by_quota,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRequestedServiceQuotaChangeHistoryByQuotaInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_requested_service_quota_change_history_by_quota, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRequestedServiceQuotaChangeHistoryByQuota(ctx, input)
				}
				var results []*svc.ListRequestedServiceQuotaChangeHistoryByQuotaOutput
				p := svc.NewListRequestedServiceQuotaChangeHistoryByQuotaPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-service-quota-increase-requests-in-template": {
			Name:   "list-service-quota-increase-requests-in-template",
			Fields: fields_list_service_quota_increase_requests_in_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListServiceQuotaIncreaseRequestsInTemplateInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_service_quota_increase_requests_in_template, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListServiceQuotaIncreaseRequestsInTemplate(ctx, input)
				}
				var results []*svc.ListServiceQuotaIncreaseRequestsInTemplateOutput
				p := svc.NewListServiceQuotaIncreaseRequestsInTemplatePaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-service-quotas": {
			Name:   "list-service-quotas",
			Fields: fields_list_service_quotas,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListServiceQuotasInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_service_quotas, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListServiceQuotas(ctx, input)
				}
				var results []*svc.ListServiceQuotasOutput
				p := svc.NewListServiceQuotasPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-services": {
			Name:   "list-services",
			Fields: fields_list_services,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListServicesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_services, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListServices(ctx, input)
				}
				var results []*svc.ListServicesOutput
				p := svc.NewListServicesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-tags-for-resource": {
			Name:   "list-tags-for-resource",
			Fields: fields_list_tags_for_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTagsForResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_tags_for_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTagsForResource(ctx, input)
			},
		},
		"put-service-quota-increase-request-into-template": {
			Name:   "put-service-quota-increase-request-into-template",
			Fields: fields_put_service_quota_increase_request_into_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutServiceQuotaIncreaseRequestIntoTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_service_quota_increase_request_into_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutServiceQuotaIncreaseRequestIntoTemplate(ctx, input)
			},
		},
		"request-service-quota-increase": {
			Name:   "request-service-quota-increase",
			Fields: fields_request_service_quota_increase,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RequestServiceQuotaIncreaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_request_service_quota_increase, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RequestServiceQuotaIncrease(ctx, input)
			},
		},
		"start-auto-management": {
			Name:   "start-auto-management",
			Fields: fields_start_auto_management,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartAutoManagementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_auto_management, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartAutoManagement(ctx, input)
			},
		},
		"start-quota-utilization-report": {
			Name:   "start-quota-utilization-report",
			Fields: fields_start_quota_utilization_report,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartQuotaUtilizationReportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_quota_utilization_report, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartQuotaUtilizationReport(ctx, input)
			},
		},
		"stop-auto-management": {
			Name:   "stop-auto-management",
			Fields: fields_stop_auto_management,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopAutoManagementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_auto_management, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopAutoManagement(ctx, input)
			},
		},
		"tag-resource": {
			Name:   "tag-resource",
			Fields: fields_tag_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TagResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_tag_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TagResource(ctx, input)
			},
		},
		"untag-resource": {
			Name:   "untag-resource",
			Fields: fields_untag_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UntagResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_untag_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UntagResource(ctx, input)
			},
		},
		"update-auto-management": {
			Name:   "update-auto-management",
			Fields: fields_update_auto_management,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAutoManagementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_auto_management, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAutoManagement(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("servicequotas", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
