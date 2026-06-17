package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/invoicing/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"batch-get-invoice-profile", "create-invoice-unit", "create-procurement-portal-preference", "delete-invoice-unit", "delete-procurement-portal-preference", "get-invoice-pdf", "get-invoice-unit", "get-procurement-portal-preference", "list-invoice-summaries", "list-invoice-units", "list-procurement-portal-preferences", "list-tags-for-resource", "put-procurement-portal-preference", "tag-resource", "untag-resource", "update-invoice-unit", "update-procurement-portal-preference-status"},
		OperationSet: map[string]bool{"batch-get-invoice-profile": true, "create-invoice-unit": true, "create-procurement-portal-preference": true, "delete-invoice-unit": true, "delete-procurement-portal-preference": true, "get-invoice-pdf": true, "get-invoice-unit": true, "get-procurement-portal-preference": true, "list-invoice-summaries": true, "list-invoice-units": true, "list-procurement-portal-preferences": true, "list-tags-for-resource": true, "put-procurement-portal-preference": true, "tag-resource": true, "untag-resource": true, "update-invoice-unit": true, "update-procurement-portal-preference-status": true},
		OperationInputs: map[string][]string{
			"batch-get-invoice-profile":                   {"AccountIds"},
			"create-invoice-unit":                         {"Description", "InvoiceReceiver", "Name", "ResourceTags", "Rule", "TaxInheritanceDisabled"},
			"create-procurement-portal-preference":        {"BuyerDomain", "BuyerIdentifier", "ClientToken", "Contacts", "EinvoiceDeliveryEnabled", "EinvoiceDeliveryPreference", "ProcurementPortalInstanceEndpoint", "ProcurementPortalName", "ProcurementPortalSharedSecret", "PurchaseOrderRetrievalEnabled", "ResourceTags", "Selector", "SupplierDomain", "SupplierIdentifier", "TestEnvPreference"},
			"delete-invoice-unit":                         {"InvoiceUnitArn"},
			"delete-procurement-portal-preference":        {"ProcurementPortalPreferenceArn"},
			"get-invoice-pdf":                             {"InvoiceId"},
			"get-invoice-unit":                            {"AsOf", "InvoiceUnitArn"},
			"get-procurement-portal-preference":           {"ProcurementPortalPreferenceArn"},
			"list-invoice-summaries":                      {"Filter", "MaxResults", "NextToken", "Selector"},
			"list-invoice-units":                          {"AsOf", "Filters", "MaxResults", "NextToken"},
			"list-procurement-portal-preferences":         {"MaxResults", "NextToken"},
			"list-tags-for-resource":                      {"ResourceArn"},
			"put-procurement-portal-preference":           {"Contacts", "EinvoiceDeliveryEnabled", "EinvoiceDeliveryPreference", "ProcurementPortalInstanceEndpoint", "ProcurementPortalPreferenceArn", "ProcurementPortalSharedSecret", "PurchaseOrderRetrievalEnabled", "Selector", "TestEnvPreference"},
			"tag-resource":                                {"ResourceArn", "ResourceTags"},
			"untag-resource":                              {"ResourceArn", "ResourceTagKeys"},
			"update-invoice-unit":                         {"Description", "InvoiceUnitArn", "Rule", "TaxInheritanceDisabled"},
			"update-procurement-portal-preference-status": {"EinvoiceDeliveryPreferenceStatus", "EinvoiceDeliveryPreferenceStatusReason", "ProcurementPortalPreferenceArn", "PurchaseOrderRetrievalPreferenceStatus", "PurchaseOrderRetrievalPreferenceStatusReason"},
		},
		OperationInputTypes: map[string]map[string]string{
			"batch-get-invoice-profile":                   {"AccountIds": "[]string"},
			"create-invoice-unit":                         {"Description": "*string", "InvoiceReceiver": "*string", "Name": "*string", "ResourceTags": "[]types.ResourceTag", "Rule": "*types.InvoiceUnitRule", "TaxInheritanceDisabled": "bool"},
			"create-procurement-portal-preference":        {"BuyerDomain": "types.BuyerDomain", "BuyerIdentifier": "*string", "ClientToken": "*string", "Contacts": "[]types.Contact", "EinvoiceDeliveryEnabled": "*bool", "EinvoiceDeliveryPreference": "*types.EinvoiceDeliveryPreference", "ProcurementPortalInstanceEndpoint": "*string", "ProcurementPortalName": "types.ProcurementPortalName", "ProcurementPortalSharedSecret": "*string", "PurchaseOrderRetrievalEnabled": "*bool", "ResourceTags": "[]types.ResourceTag", "Selector": "*types.ProcurementPortalPreferenceSelector", "SupplierDomain": "types.SupplierDomain", "SupplierIdentifier": "*string", "TestEnvPreference": "*types.TestEnvPreferenceInput"},
			"delete-invoice-unit":                         {"InvoiceUnitArn": "*string"},
			"delete-procurement-portal-preference":        {"ProcurementPortalPreferenceArn": "*string"},
			"get-invoice-pdf":                             {"InvoiceId": "*string"},
			"get-invoice-unit":                            {"AsOf": "*time.Time", "InvoiceUnitArn": "*string"},
			"get-procurement-portal-preference":           {"ProcurementPortalPreferenceArn": "*string"},
			"list-invoice-summaries":                      {"Filter": "*types.InvoiceSummariesFilter", "MaxResults": "*int32", "NextToken": "*string", "Selector": "*types.InvoiceSummariesSelector"},
			"list-invoice-units":                          {"AsOf": "*time.Time", "Filters": "*types.Filters", "MaxResults": "*int32", "NextToken": "*string"},
			"list-procurement-portal-preferences":         {"MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":                      {"ResourceArn": "*string"},
			"put-procurement-portal-preference":           {"Contacts": "[]types.Contact", "EinvoiceDeliveryEnabled": "*bool", "EinvoiceDeliveryPreference": "*types.EinvoiceDeliveryPreference", "ProcurementPortalInstanceEndpoint": "*string", "ProcurementPortalPreferenceArn": "*string", "ProcurementPortalSharedSecret": "*string", "PurchaseOrderRetrievalEnabled": "*bool", "Selector": "*types.ProcurementPortalPreferenceSelector", "TestEnvPreference": "*types.TestEnvPreferenceInput"},
			"tag-resource":                                {"ResourceArn": "*string", "ResourceTags": "[]types.ResourceTag"},
			"untag-resource":                              {"ResourceArn": "*string", "ResourceTagKeys": "[]string"},
			"update-invoice-unit":                         {"Description": "*string", "InvoiceUnitArn": "*string", "Rule": "*types.InvoiceUnitRule", "TaxInheritanceDisabled": "*bool"},
			"update-procurement-portal-preference-status": {"EinvoiceDeliveryPreferenceStatus": "types.ProcurementPortalPreferenceStatus", "EinvoiceDeliveryPreferenceStatusReason": "*string", "ProcurementPortalPreferenceArn": "*string", "PurchaseOrderRetrievalPreferenceStatus": "types.ProcurementPortalPreferenceStatus", "PurchaseOrderRetrievalPreferenceStatusReason": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"batch-get-invoice-profile":                   {"AccountIds"},
			"create-invoice-unit":                         {"InvoiceReceiver", "Name", "Rule"},
			"create-procurement-portal-preference":        {"BuyerDomain", "BuyerIdentifier", "Contacts", "EinvoiceDeliveryEnabled", "ProcurementPortalName", "PurchaseOrderRetrievalEnabled", "SupplierDomain", "SupplierIdentifier"},
			"delete-invoice-unit":                         {"InvoiceUnitArn"},
			"delete-procurement-portal-preference":        {"ProcurementPortalPreferenceArn"},
			"get-invoice-pdf":                             {"InvoiceId"},
			"get-invoice-unit":                            {"InvoiceUnitArn"},
			"get-procurement-portal-preference":           {"ProcurementPortalPreferenceArn"},
			"list-invoice-summaries":                      {"Selector"},
			"list-invoice-units":                          {},
			"list-procurement-portal-preferences":         {},
			"list-tags-for-resource":                      {"ResourceArn"},
			"put-procurement-portal-preference":           {"Contacts", "EinvoiceDeliveryEnabled", "ProcurementPortalPreferenceArn", "PurchaseOrderRetrievalEnabled"},
			"tag-resource":                                {"ResourceArn", "ResourceTags"},
			"untag-resource":                              {"ResourceArn", "ResourceTagKeys"},
			"update-invoice-unit":                         {"InvoiceUnitArn"},
			"update-procurement-portal-preference-status": {"ProcurementPortalPreferenceArn"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("invoicing", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
