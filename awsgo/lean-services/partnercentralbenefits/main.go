package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/partnercentralbenefits"
)

var fields_amend_benefit_application = []leanruntime.Field{
	{Name: "AmendmentReason", Flag: "amendment-reason", Type: "*string", Required: true},
	{Name: "Amendments", Flag: "amendments", Type: "[]types.Amendment", Required: true},
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "Revision", Flag: "revision", Type: "*string", Required: true},
}

var fields_associate_benefit_application_resource = []leanruntime.Field{
	{Name: "BenefitApplicationIdentifier", Flag: "benefit-application-identifier", Type: "*string", Required: true},
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_cancel_benefit_application = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "Reason", Flag: "reason", Type: "*string", Required: false},
}

var fields_create_benefit_application = []leanruntime.Field{
	{Name: "AssociatedResources", Flag: "associated-resources", Type: "[]string", Required: false},
	{Name: "BenefitApplicationDetails", Flag: "benefit-application-details", Type: "document.Interface", Required: false},
	{Name: "BenefitIdentifier", Flag: "benefit-identifier", Type: "*string", Required: true},
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "FileDetails", Flag: "file-details", Type: "[]types.FileInput", Required: false},
	{Name: "FulfillmentTypes", Flag: "fulfillment-types", Type: "[]types.FulfillmentType", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "PartnerContacts", Flag: "partner-contacts", Type: "[]types.Contact", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_disassociate_benefit_application_resource = []leanruntime.Field{
	{Name: "BenefitApplicationIdentifier", Flag: "benefit-application-identifier", Type: "*string", Required: true},
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_get_benefit = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_benefit_allocation = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_benefit_application = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_list_benefit_allocations = []leanruntime.Field{
	{Name: "BenefitApplicationIdentifiers", Flag: "benefit-application-identifiers", Type: "[]string", Required: false},
	{Name: "BenefitIdentifiers", Flag: "benefit-identifiers", Type: "[]string", Required: false},
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "FulfillmentTypes", Flag: "fulfillment-types", Type: "[]types.FulfillmentType", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "[]types.BenefitAllocationStatus", Required: false},
}

var fields_list_benefit_applications = []leanruntime.Field{
	{Name: "AssociatedResourceArns", Flag: "associated-resource-arns", Type: "[]string", Required: false},
	{Name: "AssociatedResources", Flag: "associated-resources", Type: "[]types.AssociatedResource", Required: false},
	{Name: "BenefitIdentifiers", Flag: "benefit-identifiers", Type: "[]string", Required: false},
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "FulfillmentTypes", Flag: "fulfillment-types", Type: "[]types.FulfillmentType", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Programs", Flag: "programs", Type: "[]string", Required: false},
	{Name: "Stages", Flag: "stages", Type: "[]string", Required: false},
	{Name: "Status", Flag: "status", Type: "[]types.BenefitApplicationStatus", Required: false},
}

var fields_list_benefits = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "FulfillmentTypes", Flag: "fulfillment-types", Type: "[]types.FulfillmentType", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Programs", Flag: "programs", Type: "[]string", Required: false},
	{Name: "Status", Flag: "status", Type: "[]types.BenefitStatus", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_recall_benefit_application = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "Reason", Flag: "reason", Type: "*string", Required: true},
}

var fields_submit_benefit_application = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_benefit_application = []leanruntime.Field{
	{Name: "BenefitApplicationDetails", Flag: "benefit-application-details", Type: "document.Interface", Required: false},
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "FileDetails", Flag: "file-details", Type: "[]types.FileInput", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "PartnerContacts", Flag: "partner-contacts", Type: "[]types.Contact", Required: false},
	{Name: "Revision", Flag: "revision", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"amend-benefit-application": {
			Name:   "amend-benefit-application",
			Fields: fields_amend_benefit_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AmendBenefitApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_amend_benefit_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AmendBenefitApplication(ctx, input)
			},
		},
		"associate-benefit-application-resource": {
			Name:   "associate-benefit-application-resource",
			Fields: fields_associate_benefit_application_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateBenefitApplicationResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_benefit_application_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateBenefitApplicationResource(ctx, input)
			},
		},
		"cancel-benefit-application": {
			Name:   "cancel-benefit-application",
			Fields: fields_cancel_benefit_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelBenefitApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_benefit_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelBenefitApplication(ctx, input)
			},
		},
		"create-benefit-application": {
			Name:   "create-benefit-application",
			Fields: fields_create_benefit_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBenefitApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_benefit_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBenefitApplication(ctx, input)
			},
		},
		"disassociate-benefit-application-resource": {
			Name:   "disassociate-benefit-application-resource",
			Fields: fields_disassociate_benefit_application_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateBenefitApplicationResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_benefit_application_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateBenefitApplicationResource(ctx, input)
			},
		},
		"get-benefit": {
			Name:   "get-benefit",
			Fields: fields_get_benefit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBenefitInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_benefit, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBenefit(ctx, input)
			},
		},
		"get-benefit-allocation": {
			Name:   "get-benefit-allocation",
			Fields: fields_get_benefit_allocation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBenefitAllocationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_benefit_allocation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBenefitAllocation(ctx, input)
			},
		},
		"get-benefit-application": {
			Name:   "get-benefit-application",
			Fields: fields_get_benefit_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBenefitApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_benefit_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBenefitApplication(ctx, input)
			},
		},
		"list-benefit-allocations": {
			Name:   "list-benefit-allocations",
			Fields: fields_list_benefit_allocations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBenefitAllocationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_benefit_allocations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBenefitAllocations(ctx, input)
				}
				var results []*svc.ListBenefitAllocationsOutput
				p := svc.NewListBenefitAllocationsPaginator(client, input)
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
		"list-benefit-applications": {
			Name:   "list-benefit-applications",
			Fields: fields_list_benefit_applications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBenefitApplicationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_benefit_applications, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBenefitApplications(ctx, input)
				}
				var results []*svc.ListBenefitApplicationsOutput
				p := svc.NewListBenefitApplicationsPaginator(client, input)
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
		"list-benefits": {
			Name:   "list-benefits",
			Fields: fields_list_benefits,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBenefitsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_benefits, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBenefits(ctx, input)
				}
				var results []*svc.ListBenefitsOutput
				p := svc.NewListBenefitsPaginator(client, input)
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
		"recall-benefit-application": {
			Name:   "recall-benefit-application",
			Fields: fields_recall_benefit_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RecallBenefitApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_recall_benefit_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RecallBenefitApplication(ctx, input)
			},
		},
		"submit-benefit-application": {
			Name:   "submit-benefit-application",
			Fields: fields_submit_benefit_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SubmitBenefitApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_submit_benefit_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SubmitBenefitApplication(ctx, input)
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
		"update-benefit-application": {
			Name:   "update-benefit-application",
			Fields: fields_update_benefit_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateBenefitApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_benefit_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateBenefitApplication(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("partnercentralbenefits", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
