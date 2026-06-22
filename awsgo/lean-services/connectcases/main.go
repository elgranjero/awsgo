package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/connectcases"
)

var fields_batch_get_case_rule = []leanruntime.Field{
	{Name: "CaseRules", Flag: "case-rules", Type: "[]types.CaseRuleIdentifier", Required: true},
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
}

var fields_batch_get_field = []leanruntime.Field{
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "Fields", Flag: "fields", Type: "[]types.FieldIdentifier", Required: true},
}

var fields_batch_put_field_options = []leanruntime.Field{
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "FieldId", Flag: "field-id", Type: "*string", Required: true},
	{Name: "Options", Flag: "options", Type: "[]types.FieldOption", Required: true},
}

var fields_create_case = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "Fields", Flag: "fields", Type: "[]types.FieldValue", Required: true},
	{Name: "PerformedBy", Flag: "performed-by", Type: "types.UserUnion", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TemplateId", Flag: "template-id", Type: "*string", Required: true},
}

var fields_create_case_rule = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Rule", Flag: "rule", Type: "types.CaseRuleDetails", Required: true},
}

var fields_create_domain = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_create_field = []leanruntime.Field{
	{Name: "Attributes", Flag: "attributes", Type: "types.FieldAttributes", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.FieldType", Required: true},
}

var fields_create_layout = []leanruntime.Field{
	{Name: "Content", Flag: "content", Type: "types.LayoutContent", Required: true},
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_create_related_item = []leanruntime.Field{
	{Name: "CaseId", Flag: "case-id", Type: "*string", Required: true},
	{Name: "Content", Flag: "content", Type: "types.RelatedItemInputContent", Required: true},
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "PerformedBy", Flag: "performed-by", Type: "types.UserUnion", Required: false},
	{Name: "Type", Flag: "type", Type: "types.RelatedItemType", Required: true},
}

var fields_create_template = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "LayoutConfiguration", Flag: "layout-configuration", Type: "*types.LayoutConfiguration", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RequiredFields", Flag: "required-fields", Type: "[]types.RequiredField", Required: false},
	{Name: "Rules", Flag: "rules", Type: "[]types.TemplateRule", Required: false},
	{Name: "Status", Flag: "status", Type: "types.TemplateStatus", Required: false},
	{Name: "TagPropagationConfigurations", Flag: "tag-propagation-configurations", Type: "[]types.TagPropagationConfiguration", Required: false},
}

var fields_delete_case = []leanruntime.Field{
	{Name: "CaseId", Flag: "case-id", Type: "*string", Required: true},
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
}

var fields_delete_case_rule = []leanruntime.Field{
	{Name: "CaseRuleId", Flag: "case-rule-id", Type: "*string", Required: true},
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
}

var fields_delete_domain = []leanruntime.Field{
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
}

var fields_delete_field = []leanruntime.Field{
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "FieldId", Flag: "field-id", Type: "*string", Required: true},
}

var fields_delete_layout = []leanruntime.Field{
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "LayoutId", Flag: "layout-id", Type: "*string", Required: true},
}

var fields_delete_related_item = []leanruntime.Field{
	{Name: "CaseId", Flag: "case-id", Type: "*string", Required: true},
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "RelatedItemId", Flag: "related-item-id", Type: "*string", Required: true},
}

var fields_delete_template = []leanruntime.Field{
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "TemplateId", Flag: "template-id", Type: "*string", Required: true},
}

var fields_get_case = []leanruntime.Field{
	{Name: "CaseId", Flag: "case-id", Type: "*string", Required: true},
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "Fields", Flag: "fields", Type: "[]types.FieldIdentifier", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_case_audit_events = []leanruntime.Field{
	{Name: "CaseId", Flag: "case-id", Type: "*string", Required: true},
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_case_event_configuration = []leanruntime.Field{
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
}

var fields_get_domain = []leanruntime.Field{
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
}

var fields_get_layout = []leanruntime.Field{
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "LayoutId", Flag: "layout-id", Type: "*string", Required: true},
}

var fields_get_template = []leanruntime.Field{
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "TemplateId", Flag: "template-id", Type: "*string", Required: true},
}

var fields_list_case_rules = []leanruntime.Field{
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_cases_for_contact = []leanruntime.Field{
	{Name: "ContactArn", Flag: "contact-arn", Type: "*string", Required: true},
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_domains = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_field_options = []leanruntime.Field{
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "FieldId", Flag: "field-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Values", Flag: "values", Type: "[]string", Required: false},
}

var fields_list_fields = []leanruntime.Field{
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_layouts = []leanruntime.Field{
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_list_templates = []leanruntime.Field{
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "[]types.TemplateStatus", Required: false},
}

var fields_put_case_event_configuration = []leanruntime.Field{
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "EventBridge", Flag: "event-bridge", Type: "*types.EventBridgeConfiguration", Required: true},
}

var fields_search_all_related_items = []leanruntime.Field{
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.RelatedItemTypeFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Sorts", Flag: "sorts", Type: "[]types.SearchAllRelatedItemsSort", Required: false},
}

var fields_search_cases = []leanruntime.Field{
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "Fields", Flag: "fields", Type: "[]types.FieldIdentifier", Required: false},
	{Name: "Filter", Flag: "filter", Type: "types.CaseFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SearchTerm", Flag: "search-term", Type: "*string", Required: false},
	{Name: "Sorts", Flag: "sorts", Type: "[]types.Sort", Required: false},
}

var fields_search_related_items = []leanruntime.Field{
	{Name: "CaseId", Flag: "case-id", Type: "*string", Required: true},
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.RelatedItemTypeFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]*string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_case = []leanruntime.Field{
	{Name: "CaseId", Flag: "case-id", Type: "*string", Required: true},
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "Fields", Flag: "fields", Type: "[]types.FieldValue", Required: true},
	{Name: "PerformedBy", Flag: "performed-by", Type: "types.UserUnion", Required: false},
}

var fields_update_case_rule = []leanruntime.Field{
	{Name: "CaseRuleId", Flag: "case-rule-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Rule", Flag: "rule", Type: "types.CaseRuleDetails", Required: false},
}

var fields_update_field = []leanruntime.Field{
	{Name: "Attributes", Flag: "attributes", Type: "types.FieldAttributes", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "FieldId", Flag: "field-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_layout = []leanruntime.Field{
	{Name: "Content", Flag: "content", Type: "types.LayoutContent", Required: false},
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "LayoutId", Flag: "layout-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_template = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "LayoutConfiguration", Flag: "layout-configuration", Type: "*types.LayoutConfiguration", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "RequiredFields", Flag: "required-fields", Type: "[]types.RequiredField", Required: false},
	{Name: "Rules", Flag: "rules", Type: "[]types.TemplateRule", Required: false},
	{Name: "Status", Flag: "status", Type: "types.TemplateStatus", Required: false},
	{Name: "TagPropagationConfigurations", Flag: "tag-propagation-configurations", Type: "[]types.TagPropagationConfiguration", Required: false},
	{Name: "TemplateId", Flag: "template-id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"batch-get-case-rule": {
			Name:   "batch-get-case-rule",
			Fields: fields_batch_get_case_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetCaseRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_case_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetCaseRule(ctx, input)
			},
		},
		"batch-get-field": {
			Name:   "batch-get-field",
			Fields: fields_batch_get_field,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetFieldInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_field, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetField(ctx, input)
			},
		},
		"batch-put-field-options": {
			Name:   "batch-put-field-options",
			Fields: fields_batch_put_field_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchPutFieldOptionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_put_field_options, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchPutFieldOptions(ctx, input)
			},
		},
		"create-case": {
			Name:   "create-case",
			Fields: fields_create_case,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_case, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCase(ctx, input)
			},
		},
		"create-case-rule": {
			Name:   "create-case-rule",
			Fields: fields_create_case_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCaseRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_case_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCaseRule(ctx, input)
			},
		},
		"create-domain": {
			Name:   "create-domain",
			Fields: fields_create_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDomain(ctx, input)
			},
		},
		"create-field": {
			Name:   "create-field",
			Fields: fields_create_field,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFieldInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_field, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateField(ctx, input)
			},
		},
		"create-layout": {
			Name:   "create-layout",
			Fields: fields_create_layout,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLayoutInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_layout, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLayout(ctx, input)
			},
		},
		"create-related-item": {
			Name:   "create-related-item",
			Fields: fields_create_related_item,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRelatedItemInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_related_item, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRelatedItem(ctx, input)
			},
		},
		"create-template": {
			Name:   "create-template",
			Fields: fields_create_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTemplate(ctx, input)
			},
		},
		"delete-case": {
			Name:   "delete-case",
			Fields: fields_delete_case,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_case, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCase(ctx, input)
			},
		},
		"delete-case-rule": {
			Name:   "delete-case-rule",
			Fields: fields_delete_case_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCaseRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_case_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCaseRule(ctx, input)
			},
		},
		"delete-domain": {
			Name:   "delete-domain",
			Fields: fields_delete_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDomain(ctx, input)
			},
		},
		"delete-field": {
			Name:   "delete-field",
			Fields: fields_delete_field,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFieldInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_field, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteField(ctx, input)
			},
		},
		"delete-layout": {
			Name:   "delete-layout",
			Fields: fields_delete_layout,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLayoutInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_layout, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLayout(ctx, input)
			},
		},
		"delete-related-item": {
			Name:   "delete-related-item",
			Fields: fields_delete_related_item,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRelatedItemInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_related_item, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRelatedItem(ctx, input)
			},
		},
		"delete-template": {
			Name:   "delete-template",
			Fields: fields_delete_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTemplate(ctx, input)
			},
		},
		"get-case": {
			Name:   "get-case",
			Fields: fields_get_case,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCaseInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_case, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetCase(ctx, input)
				}
				var results []*svc.GetCaseOutput
				p := svc.NewGetCasePaginator(client, input)
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
		"get-case-audit-events": {
			Name:   "get-case-audit-events",
			Fields: fields_get_case_audit_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCaseAuditEventsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_case_audit_events, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetCaseAuditEvents(ctx, input)
				}
				var results []*svc.GetCaseAuditEventsOutput
				p := svc.NewGetCaseAuditEventsPaginator(client, input)
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
		"get-case-event-configuration": {
			Name:   "get-case-event-configuration",
			Fields: fields_get_case_event_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCaseEventConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_case_event_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCaseEventConfiguration(ctx, input)
			},
		},
		"get-domain": {
			Name:   "get-domain",
			Fields: fields_get_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDomain(ctx, input)
			},
		},
		"get-layout": {
			Name:   "get-layout",
			Fields: fields_get_layout,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLayoutInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_layout, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLayout(ctx, input)
			},
		},
		"get-template": {
			Name:   "get-template",
			Fields: fields_get_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTemplate(ctx, input)
			},
		},
		"list-case-rules": {
			Name:   "list-case-rules",
			Fields: fields_list_case_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCaseRulesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_case_rules, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCaseRules(ctx, input)
				}
				var results []*svc.ListCaseRulesOutput
				p := svc.NewListCaseRulesPaginator(client, input)
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
		"list-cases-for-contact": {
			Name:   "list-cases-for-contact",
			Fields: fields_list_cases_for_contact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCasesForContactInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_cases_for_contact, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCasesForContact(ctx, input)
				}
				var results []*svc.ListCasesForContactOutput
				p := svc.NewListCasesForContactPaginator(client, input)
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
		"list-domains": {
			Name:   "list-domains",
			Fields: fields_list_domains,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDomainsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_domains, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDomains(ctx, input)
				}
				var results []*svc.ListDomainsOutput
				p := svc.NewListDomainsPaginator(client, input)
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
		"list-field-options": {
			Name:   "list-field-options",
			Fields: fields_list_field_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFieldOptionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_field_options, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFieldOptions(ctx, input)
				}
				var results []*svc.ListFieldOptionsOutput
				p := svc.NewListFieldOptionsPaginator(client, input)
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
		"list-fields": {
			Name:   "list-fields",
			Fields: fields_list_fields,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFieldsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_fields, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFields(ctx, input)
				}
				var results []*svc.ListFieldsOutput
				p := svc.NewListFieldsPaginator(client, input)
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
		"list-layouts": {
			Name:   "list-layouts",
			Fields: fields_list_layouts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLayoutsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_layouts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListLayouts(ctx, input)
				}
				var results []*svc.ListLayoutsOutput
				p := svc.NewListLayoutsPaginator(client, input)
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
		"list-templates": {
			Name:   "list-templates",
			Fields: fields_list_templates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTemplatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_templates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTemplates(ctx, input)
				}
				var results []*svc.ListTemplatesOutput
				p := svc.NewListTemplatesPaginator(client, input)
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
		"put-case-event-configuration": {
			Name:   "put-case-event-configuration",
			Fields: fields_put_case_event_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutCaseEventConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_case_event_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutCaseEventConfiguration(ctx, input)
			},
		},
		"search-all-related-items": {
			Name:   "search-all-related-items",
			Fields: fields_search_all_related_items,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchAllRelatedItemsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_all_related_items, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchAllRelatedItems(ctx, input)
				}
				var results []*svc.SearchAllRelatedItemsOutput
				p := svc.NewSearchAllRelatedItemsPaginator(client, input)
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
		"search-cases": {
			Name:   "search-cases",
			Fields: fields_search_cases,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchCasesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_cases, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchCases(ctx, input)
				}
				var results []*svc.SearchCasesOutput
				p := svc.NewSearchCasesPaginator(client, input)
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
		"search-related-items": {
			Name:   "search-related-items",
			Fields: fields_search_related_items,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchRelatedItemsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_related_items, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchRelatedItems(ctx, input)
				}
				var results []*svc.SearchRelatedItemsOutput
				p := svc.NewSearchRelatedItemsPaginator(client, input)
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
		"update-case": {
			Name:   "update-case",
			Fields: fields_update_case,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_case, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCase(ctx, input)
			},
		},
		"update-case-rule": {
			Name:   "update-case-rule",
			Fields: fields_update_case_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCaseRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_case_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCaseRule(ctx, input)
			},
		},
		"update-field": {
			Name:   "update-field",
			Fields: fields_update_field,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFieldInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_field, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateField(ctx, input)
			},
		},
		"update-layout": {
			Name:   "update-layout",
			Fields: fields_update_layout,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateLayoutInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_layout, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateLayout(ctx, input)
			},
		},
		"update-template": {
			Name:   "update-template",
			Fields: fields_update_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTemplate(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("connectcases", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
