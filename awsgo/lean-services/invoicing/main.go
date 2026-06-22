package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/invoicing"
)

var fields_batch_get_invoice_profile = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: true},
}

var fields_create_invoice_unit = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InvoiceReceiver", Flag: "invoice-receiver", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ResourceTags", Flag: "resource-tags", Type: "[]types.ResourceTag", Required: false},
	{Name: "Rule", Flag: "rule", Type: "*types.InvoiceUnitRule", Required: true},
	{Name: "TaxInheritanceDisabled", Flag: "tax-inheritance-disabled", Type: "bool", Required: false},
}

var fields_create_procurement_portal_preference = []leanruntime.Field{
	{Name: "BuyerDomain", Flag: "buyer-domain", Type: "types.BuyerDomain", Required: true},
	{Name: "BuyerIdentifier", Flag: "buyer-identifier", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Contacts", Flag: "contacts", Type: "[]types.Contact", Required: true},
	{Name: "EinvoiceDeliveryEnabled", Flag: "einvoice-delivery-enabled", Type: "*bool", Required: true},
	{Name: "EinvoiceDeliveryPreference", Flag: "einvoice-delivery-preference", Type: "*types.EinvoiceDeliveryPreference", Required: false},
	{Name: "ProcurementPortalInstanceEndpoint", Flag: "procurement-portal-instance-endpoint", Type: "*string", Required: false},
	{Name: "ProcurementPortalName", Flag: "procurement-portal-name", Type: "types.ProcurementPortalName", Required: true},
	{Name: "ProcurementPortalSharedSecret", Flag: "procurement-portal-shared-secret", Type: "*string", Required: false},
	{Name: "PurchaseOrderRetrievalEnabled", Flag: "purchase-order-retrieval-enabled", Type: "*bool", Required: true},
	{Name: "ResourceTags", Flag: "resource-tags", Type: "[]types.ResourceTag", Required: false},
	{Name: "Selector", Flag: "selector", Type: "*types.ProcurementPortalPreferenceSelector", Required: false},
	{Name: "SupplierDomain", Flag: "supplier-domain", Type: "types.SupplierDomain", Required: true},
	{Name: "SupplierIdentifier", Flag: "supplier-identifier", Type: "*string", Required: true},
	{Name: "TestEnvPreference", Flag: "test-env-preference", Type: "*types.TestEnvPreferenceInput", Required: false},
}

var fields_delete_invoice_unit = []leanruntime.Field{
	{Name: "InvoiceUnitArn", Flag: "invoice-unit-arn", Type: "*string", Required: true},
}

var fields_delete_procurement_portal_preference = []leanruntime.Field{
	{Name: "ProcurementPortalPreferenceArn", Flag: "procurement-portal-preference-arn", Type: "*string", Required: true},
}

var fields_get_invoice_pdf = []leanruntime.Field{
	{Name: "InvoiceId", Flag: "invoice-id", Type: "*string", Required: true},
}

var fields_get_invoice_unit = []leanruntime.Field{
	{Name: "AsOf", Flag: "as-of", Type: "*time.Time", Required: false},
	{Name: "InvoiceUnitArn", Flag: "invoice-unit-arn", Type: "*string", Required: true},
}

var fields_get_procurement_portal_preference = []leanruntime.Field{
	{Name: "ProcurementPortalPreferenceArn", Flag: "procurement-portal-preference-arn", Type: "*string", Required: true},
}

var fields_list_invoice_summaries = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.InvoiceSummariesFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Selector", Flag: "selector", Type: "*types.InvoiceSummariesSelector", Required: true},
}

var fields_list_invoice_units = []leanruntime.Field{
	{Name: "AsOf", Flag: "as-of", Type: "*time.Time", Required: false},
	{Name: "Filters", Flag: "filters", Type: "*types.Filters", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_procurement_portal_preferences = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_put_procurement_portal_preference = []leanruntime.Field{
	{Name: "Contacts", Flag: "contacts", Type: "[]types.Contact", Required: true},
	{Name: "EinvoiceDeliveryEnabled", Flag: "einvoice-delivery-enabled", Type: "*bool", Required: true},
	{Name: "EinvoiceDeliveryPreference", Flag: "einvoice-delivery-preference", Type: "*types.EinvoiceDeliveryPreference", Required: false},
	{Name: "ProcurementPortalInstanceEndpoint", Flag: "procurement-portal-instance-endpoint", Type: "*string", Required: false},
	{Name: "ProcurementPortalPreferenceArn", Flag: "procurement-portal-preference-arn", Type: "*string", Required: true},
	{Name: "ProcurementPortalSharedSecret", Flag: "procurement-portal-shared-secret", Type: "*string", Required: false},
	{Name: "PurchaseOrderRetrievalEnabled", Flag: "purchase-order-retrieval-enabled", Type: "*bool", Required: true},
	{Name: "Selector", Flag: "selector", Type: "*types.ProcurementPortalPreferenceSelector", Required: false},
	{Name: "TestEnvPreference", Flag: "test-env-preference", Type: "*types.TestEnvPreferenceInput", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "ResourceTags", Flag: "resource-tags", Type: "[]types.ResourceTag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "ResourceTagKeys", Flag: "resource-tag-keys", Type: "[]string", Required: true},
}

var fields_update_invoice_unit = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InvoiceUnitArn", Flag: "invoice-unit-arn", Type: "*string", Required: true},
	{Name: "Rule", Flag: "rule", Type: "*types.InvoiceUnitRule", Required: false},
	{Name: "TaxInheritanceDisabled", Flag: "tax-inheritance-disabled", Type: "*bool", Required: false},
}

var fields_update_procurement_portal_preference_status = []leanruntime.Field{
	{Name: "EinvoiceDeliveryPreferenceStatus", Flag: "einvoice-delivery-preference-status", Type: "types.ProcurementPortalPreferenceStatus", Required: false},
	{Name: "EinvoiceDeliveryPreferenceStatusReason", Flag: "einvoice-delivery-preference-status-reason", Type: "*string", Required: false},
	{Name: "ProcurementPortalPreferenceArn", Flag: "procurement-portal-preference-arn", Type: "*string", Required: true},
	{Name: "PurchaseOrderRetrievalPreferenceStatus", Flag: "purchase-order-retrieval-preference-status", Type: "types.ProcurementPortalPreferenceStatus", Required: false},
	{Name: "PurchaseOrderRetrievalPreferenceStatusReason", Flag: "purchase-order-retrieval-preference-status-reason", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"batch-get-invoice-profile": {
			Name:   "batch-get-invoice-profile",
			Fields: fields_batch_get_invoice_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetInvoiceProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_invoice_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetInvoiceProfile(ctx, input)
			},
		},
		"create-invoice-unit": {
			Name:   "create-invoice-unit",
			Fields: fields_create_invoice_unit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateInvoiceUnitInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_invoice_unit, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateInvoiceUnit(ctx, input)
			},
		},
		"create-procurement-portal-preference": {
			Name:   "create-procurement-portal-preference",
			Fields: fields_create_procurement_portal_preference,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateProcurementPortalPreferenceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_procurement_portal_preference, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateProcurementPortalPreference(ctx, input)
			},
		},
		"delete-invoice-unit": {
			Name:   "delete-invoice-unit",
			Fields: fields_delete_invoice_unit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteInvoiceUnitInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_invoice_unit, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteInvoiceUnit(ctx, input)
			},
		},
		"delete-procurement-portal-preference": {
			Name:   "delete-procurement-portal-preference",
			Fields: fields_delete_procurement_portal_preference,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteProcurementPortalPreferenceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_procurement_portal_preference, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteProcurementPortalPreference(ctx, input)
			},
		},
		"get-invoice-pdf": {
			Name:   "get-invoice-pdf",
			Fields: fields_get_invoice_pdf,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetInvoicePDFInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_invoice_pdf, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetInvoicePDF(ctx, input)
			},
		},
		"get-invoice-unit": {
			Name:   "get-invoice-unit",
			Fields: fields_get_invoice_unit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetInvoiceUnitInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_invoice_unit, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetInvoiceUnit(ctx, input)
			},
		},
		"get-procurement-portal-preference": {
			Name:   "get-procurement-portal-preference",
			Fields: fields_get_procurement_portal_preference,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetProcurementPortalPreferenceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_procurement_portal_preference, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetProcurementPortalPreference(ctx, input)
			},
		},
		"list-invoice-summaries": {
			Name:   "list-invoice-summaries",
			Fields: fields_list_invoice_summaries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListInvoiceSummariesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_invoice_summaries, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListInvoiceSummaries(ctx, input)
				}
				var results []*svc.ListInvoiceSummariesOutput
				p := svc.NewListInvoiceSummariesPaginator(client, input)
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
		"list-invoice-units": {
			Name:   "list-invoice-units",
			Fields: fields_list_invoice_units,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListInvoiceUnitsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_invoice_units, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListInvoiceUnits(ctx, input)
				}
				var results []*svc.ListInvoiceUnitsOutput
				p := svc.NewListInvoiceUnitsPaginator(client, input)
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
		"list-procurement-portal-preferences": {
			Name:   "list-procurement-portal-preferences",
			Fields: fields_list_procurement_portal_preferences,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProcurementPortalPreferencesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_procurement_portal_preferences, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListProcurementPortalPreferences(ctx, input)
				}
				var results []*svc.ListProcurementPortalPreferencesOutput
				p := svc.NewListProcurementPortalPreferencesPaginator(client, input)
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
		"put-procurement-portal-preference": {
			Name:   "put-procurement-portal-preference",
			Fields: fields_put_procurement_portal_preference,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutProcurementPortalPreferenceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_procurement_portal_preference, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutProcurementPortalPreference(ctx, input)
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
		"update-invoice-unit": {
			Name:   "update-invoice-unit",
			Fields: fields_update_invoice_unit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateInvoiceUnitInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_invoice_unit, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateInvoiceUnit(ctx, input)
			},
		},
		"update-procurement-portal-preference-status": {
			Name:   "update-procurement-portal-preference-status",
			Fields: fields_update_procurement_portal_preference_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateProcurementPortalPreferenceStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_procurement_portal_preference_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateProcurementPortalPreferenceStatus(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("invoicing", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
