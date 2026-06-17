package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/invoicing"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// invoicingCmd represents the invoicing command
var _invoicingCmd = &cobra.Command{
	Use:   "invoicing",
	Short: "AWS invoicing CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := invoicing.NewFromConfig(cfg)
		if _invoicingBatchGetInvoiceProfile {
			invoicing_BatchGetInvoiceProfile(cfg, client)
			return
		}
		if _invoicingCreateInvoiceUnit {
			invoicing_CreateInvoiceUnit(cfg, client)
			return
		}
		if _invoicingCreateProcurementPortalPreference {
			invoicing_CreateProcurementPortalPreference(cfg, client)
			return
		}
		if _invoicingDeleteInvoiceUnit {
			invoicing_DeleteInvoiceUnit(cfg, client)
			return
		}
		if _invoicingDeleteProcurementPortalPreference {
			invoicing_DeleteProcurementPortalPreference(cfg, client)
			return
		}
		if _invoicingGetInvoicePDF {
			invoicing_GetInvoicePDF(cfg, client)
			return
		}
		if _invoicingGetInvoiceUnit {
			invoicing_GetInvoiceUnit(cfg, client)
			return
		}
		if _invoicingGetProcurementPortalPreference {
			invoicing_GetProcurementPortalPreference(cfg, client)
			return
		}
		if _invoicingListInvoiceSummaries {
			invoicing_ListInvoiceSummaries(cfg, client)
			return
		}
		if _invoicingListInvoiceUnits {
			invoicing_ListInvoiceUnits(cfg, client)
			return
		}
		if _invoicingListProcurementPortalPreferences {
			invoicing_ListProcurementPortalPreferences(cfg, client)
			return
		}
		if _invoicingListTagsForResource {
			invoicing_ListTagsForResource(cfg, client)
			return
		}
		if _invoicingPutProcurementPortalPreference {
			invoicing_PutProcurementPortalPreference(cfg, client)
			return
		}
		if _invoicingTagResource {
			invoicing_TagResource(cfg, client)
			return
		}
		if _invoicingUntagResource {
			invoicing_UntagResource(cfg, client)
			return
		}
		if _invoicingUpdateInvoiceUnit {
			invoicing_UpdateInvoiceUnit(cfg, client)
			return
		}
		if _invoicingUpdateProcurementPortalPreferenceStatus {
			invoicing_UpdateProcurementPortalPreferenceStatus(cfg, client)
			return
		}

	},
}

var (
	_invoicingBatchGetInvoiceProfile                  bool
	_invoicingCreateInvoiceUnit                       bool
	_invoicingCreateProcurementPortalPreference       bool
	_invoicingDeleteInvoiceUnit                       bool
	_invoicingDeleteProcurementPortalPreference       bool
	_invoicingGetInvoicePDF                           bool
	_invoicingGetInvoiceUnit                          bool
	_invoicingGetProcurementPortalPreference          bool
	_invoicingListInvoiceSummaries                    bool
	_invoicingListInvoiceUnits                        bool
	_invoicingListProcurementPortalPreferences        bool
	_invoicingListTagsForResource                     bool
	_invoicingPutProcurementPortalPreference          bool
	_invoicingTagResource                             bool
	_invoicingUntagResource                           bool
	_invoicingUpdateInvoiceUnit                       bool
	_invoicingUpdateProcurementPortalPreferenceStatus bool

	_invoicingAccountIds                                   []string
	_invoicingAsOf                                         string
	_invoicingBuyerDomain                                  string
	_invoicingBuyerIdentifier                              string
	_invoicingClientToken                                  string
	_invoicingContacts                                     string
	_invoicingDescription                                  string
	_invoicingEinvoiceDeliveryEnabled                      string
	_invoicingEinvoiceDeliveryPreference                   string
	_invoicingEinvoiceDeliveryPreferenceStatus             string
	_invoicingEinvoiceDeliveryPreferenceStatusReason       string
	_invoicingFilter                                       string
	_invoicingFilters                                      string
	_invoicingInvoiceId                                    string
	_invoicingInvoiceReceiver                              string
	_invoicingInvoiceUnitArn                               string
	_invoicingMaxResults                                   string
	_invoicingName                                         string
	_invoicingNextToken                                    string
	_invoicingProcurementPortalInstanceEndpoint            string
	_invoicingProcurementPortalName                        string
	_invoicingProcurementPortalPreferenceArn               string
	_invoicingProcurementPortalSharedSecret                string
	_invoicingPurchaseOrderRetrievalEnabled                string
	_invoicingPurchaseOrderRetrievalPreferenceStatus       string
	_invoicingPurchaseOrderRetrievalPreferenceStatusReason string
	_invoicingResourceArn                                  string
	_invoicingResourceTagKeys                              []string
	_invoicingResourceTags                                 string
	_invoicingRule                                         string
	_invoicingSelector                                     string
	_invoicingSupplierDomain                               string
	_invoicingSupplierIdentifier                           string
	_invoicingTaxInheritanceDisabled                       string
	_invoicingTestEnvPreference                            string
)

// This gets the invoice profile associated with a set of accounts. The accounts
// must be linked accounts under the requester management account organization.
func invoicing_BatchGetInvoiceProfile(cfg aws.Config, client *invoicing.Client) {
	input := &invoicing.BatchGetInvoiceProfileInput{
		// AccountIds: []string, // Required
	}

	if len(_invoicingAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _invoicingAccountIds...)
	}

	if resp, err := client.BatchGetInvoiceProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This creates a new invoice unit with the provided definition.
func invoicing_CreateInvoiceUnit(cfg aws.Config, client *invoicing.Client) {
	input := &invoicing.CreateInvoiceUnitInput{
		// InvoiceReceiver: *string, // Required
		// Name: *string, // Required
		// Rule: *types.InvoiceUnitRule, // Required
	}

	if len(_invoicingInvoiceReceiver) > 0 {
		input.InvoiceReceiver = aws.String(_invoicingInvoiceReceiver)
	}
	if len(_invoicingName) > 0 {
		input.Name = aws.String(_invoicingName)
	}
	if len(_invoicingRule) > 0 {
		if err := assignInputField(input, "Rule", _invoicingRule); err != nil {
			log.Errorf("invalid --rule: %s", err.Error())
			return
		}
	}
	if len(_invoicingDescription) > 0 {
		input.Description = aws.String(_invoicingDescription)
	}
	if len(_invoicingResourceTags) > 0 {
		if err := assignInputField(input, "ResourceTags", _invoicingResourceTags); err != nil {
			log.Errorf("invalid --resource-tags: %s", err.Error())
			return
		}
	}
	if len(_invoicingTaxInheritanceDisabled) > 0 {
		if err := assignInputField(input, "TaxInheritanceDisabled", _invoicingTaxInheritanceDisabled); err != nil {
			log.Errorf("invalid --tax-inheritance-disabled: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateInvoiceUnit(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a procurement portal preference configuration for e-invoice delivery
// and purchase order retrieval. This preference defines how invoices are delivered
// to a procurement portal and how purchase orders are retrieved.
func invoicing_CreateProcurementPortalPreference(cfg aws.Config, client *invoicing.Client) {
	input := &invoicing.CreateProcurementPortalPreferenceInput{
		// BuyerDomain: types.BuyerDomain, // Required
		// BuyerIdentifier: *string, // Required
		// Contacts: []types.Contact, // Required
		// EinvoiceDeliveryEnabled: *bool, // Required
		// ProcurementPortalName: types.ProcurementPortalName, // Required
		// PurchaseOrderRetrievalEnabled: *bool, // Required
		// SupplierDomain: types.SupplierDomain, // Required
		// SupplierIdentifier: *string, // Required
	}

	if len(_invoicingBuyerDomain) > 0 {
		if err := assignInputField(input, "BuyerDomain", _invoicingBuyerDomain); err != nil {
			log.Errorf("invalid --buyer-domain: %s", err.Error())
			return
		}
	}
	if len(_invoicingBuyerIdentifier) > 0 {
		input.BuyerIdentifier = aws.String(_invoicingBuyerIdentifier)
	}
	if len(_invoicingContacts) > 0 {
		if err := assignInputField(input, "Contacts", _invoicingContacts); err != nil {
			log.Errorf("invalid --contacts: %s", err.Error())
			return
		}
	}
	if len(_invoicingEinvoiceDeliveryEnabled) > 0 {
		if err := assignInputField(input, "EinvoiceDeliveryEnabled", _invoicingEinvoiceDeliveryEnabled); err != nil {
			log.Errorf("invalid --einvoice-delivery-enabled: %s", err.Error())
			return
		}
	}
	if len(_invoicingProcurementPortalName) > 0 {
		if err := assignInputField(input, "ProcurementPortalName", _invoicingProcurementPortalName); err != nil {
			log.Errorf("invalid --procurement-portal-name: %s", err.Error())
			return
		}
	}
	if len(_invoicingPurchaseOrderRetrievalEnabled) > 0 {
		if err := assignInputField(input, "PurchaseOrderRetrievalEnabled", _invoicingPurchaseOrderRetrievalEnabled); err != nil {
			log.Errorf("invalid --purchase-order-retrieval-enabled: %s", err.Error())
			return
		}
	}
	if len(_invoicingSupplierDomain) > 0 {
		if err := assignInputField(input, "SupplierDomain", _invoicingSupplierDomain); err != nil {
			log.Errorf("invalid --supplier-domain: %s", err.Error())
			return
		}
	}
	if len(_invoicingSupplierIdentifier) > 0 {
		input.SupplierIdentifier = aws.String(_invoicingSupplierIdentifier)
	}
	if len(_invoicingClientToken) > 0 {
		input.ClientToken = aws.String(_invoicingClientToken)
	}
	if len(_invoicingEinvoiceDeliveryPreference) > 0 {
		if err := assignInputField(input, "EinvoiceDeliveryPreference", _invoicingEinvoiceDeliveryPreference); err != nil {
			log.Errorf("invalid --einvoice-delivery-preference: %s", err.Error())
			return
		}
	}
	if len(_invoicingProcurementPortalInstanceEndpoint) > 0 {
		input.ProcurementPortalInstanceEndpoint = aws.String(_invoicingProcurementPortalInstanceEndpoint)
	}
	if len(_invoicingProcurementPortalSharedSecret) > 0 {
		input.ProcurementPortalSharedSecret = aws.String(_invoicingProcurementPortalSharedSecret)
	}
	if len(_invoicingResourceTags) > 0 {
		if err := assignInputField(input, "ResourceTags", _invoicingResourceTags); err != nil {
			log.Errorf("invalid --resource-tags: %s", err.Error())
			return
		}
	}
	if len(_invoicingSelector) > 0 {
		if err := assignInputField(input, "Selector", _invoicingSelector); err != nil {
			log.Errorf("invalid --selector: %s", err.Error())
			return
		}
	}
	if len(_invoicingTestEnvPreference) > 0 {
		if err := assignInputField(input, "TestEnvPreference", _invoicingTestEnvPreference); err != nil {
			log.Errorf("invalid --test-env-preference: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateProcurementPortalPreference(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This deletes an invoice unit with the provided invoice unit ARN.
func invoicing_DeleteInvoiceUnit(cfg aws.Config, client *invoicing.Client) {
	input := &invoicing.DeleteInvoiceUnitInput{
		// InvoiceUnitArn: *string, // Required
	}

	if len(_invoicingInvoiceUnitArn) > 0 {
		input.InvoiceUnitArn = aws.String(_invoicingInvoiceUnitArn)
	}

	if resp, err := client.DeleteInvoiceUnit(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing procurement portal preference. This action cannot be
// undone. Active e-invoice delivery and PO retrieval configurations will be
// terminated.
func invoicing_DeleteProcurementPortalPreference(cfg aws.Config, client *invoicing.Client) {
	input := &invoicing.DeleteProcurementPortalPreferenceInput{
		// ProcurementPortalPreferenceArn: *string, // Required
	}

	if len(_invoicingProcurementPortalPreferenceArn) > 0 {
		input.ProcurementPortalPreferenceArn = aws.String(_invoicingProcurementPortalPreferenceArn)
	}

	if resp, err := client.DeleteProcurementPortalPreference(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a URL to download the invoice document and supplemental documents
// associated with an invoice. The URLs are pre-signed and have expiration time.
// For special cases like Brazil, where Amazon Web Services generated invoice
// identifiers and government provided identifiers do not match, use the Amazon Web
// Services generated invoice identifier when making API requests. To grant IAM
// permission to use this operation, the caller needs the invoicing:GetInvoicePDF
// policy action.
func invoicing_GetInvoicePDF(cfg aws.Config, client *invoicing.Client) {
	input := &invoicing.GetInvoicePDFInput{
		// InvoiceId: *string, // Required
	}

	if len(_invoicingInvoiceId) > 0 {
		input.InvoiceId = aws.String(_invoicingInvoiceId)
	}

	if resp, err := client.GetInvoicePDF(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This retrieves the invoice unit definition.
func invoicing_GetInvoiceUnit(cfg aws.Config, client *invoicing.Client) {
	input := &invoicing.GetInvoiceUnitInput{
		// InvoiceUnitArn: *string, // Required
	}

	if len(_invoicingInvoiceUnitArn) > 0 {
		input.InvoiceUnitArn = aws.String(_invoicingInvoiceUnitArn)
	}
	if len(_invoicingAsOf) > 0 {
		if err := assignInputField(input, "AsOf", _invoicingAsOf); err != nil {
			log.Errorf("invalid --as-of: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetInvoiceUnit(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the details of a specific procurement portal preference configuration.
func invoicing_GetProcurementPortalPreference(cfg aws.Config, client *invoicing.Client) {
	input := &invoicing.GetProcurementPortalPreferenceInput{
		// ProcurementPortalPreferenceArn: *string, // Required
	}

	if len(_invoicingProcurementPortalPreferenceArn) > 0 {
		input.ProcurementPortalPreferenceArn = aws.String(_invoicingProcurementPortalPreferenceArn)
	}

	if resp, err := client.GetProcurementPortalPreference(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves your invoice details programmatically, without line item details.
func invoicing_ListInvoiceSummaries(cfg aws.Config, client *invoicing.Client) {
	input := &invoicing.ListInvoiceSummariesInput{
		// Selector: *types.InvoiceSummariesSelector, // Required
	}

	if len(_invoicingSelector) > 0 {
		if err := assignInputField(input, "Selector", _invoicingSelector); err != nil {
			log.Errorf("invalid --selector: %s", err.Error())
			return
		}
	}
	if len(_invoicingFilter) > 0 {
		if err := assignInputField(input, "Filter", _invoicingFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_invoicingMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _invoicingMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_invoicingNextToken) > 0 {
		input.NextToken = aws.String(_invoicingNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListInvoiceSummaries(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*invoicing.ListInvoiceSummariesOutput
	p := invoicing.NewListInvoiceSummariesPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// This fetches a list of all invoice unit definitions for a given account, as of
// the provided AsOf date.
func invoicing_ListInvoiceUnits(cfg aws.Config, client *invoicing.Client) {
	input := &invoicing.ListInvoiceUnitsInput{}

	if len(_invoicingAsOf) > 0 {
		if err := assignInputField(input, "AsOf", _invoicingAsOf); err != nil {
			log.Errorf("invalid --as-of: %s", err.Error())
			return
		}
	}
	if len(_invoicingFilters) > 0 {
		if err := assignInputField(input, "Filters", _invoicingFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_invoicingMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _invoicingMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_invoicingNextToken) > 0 {
		input.NextToken = aws.String(_invoicingNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListInvoiceUnits(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*invoicing.ListInvoiceUnitsOutput
	p := invoicing.NewListInvoiceUnitsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Retrieves a list of procurement portal preferences associated with the Amazon
// Web Services account.
func invoicing_ListProcurementPortalPreferences(cfg aws.Config, client *invoicing.Client) {
	input := &invoicing.ListProcurementPortalPreferencesInput{}

	if len(_invoicingMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _invoicingMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_invoicingNextToken) > 0 {
		input.NextToken = aws.String(_invoicingNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListProcurementPortalPreferences(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*invoicing.ListProcurementPortalPreferencesOutput
	p := invoicing.NewListProcurementPortalPreferencesPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Lists the tags for a resource.
func invoicing_ListTagsForResource(cfg aws.Config, client *invoicing.Client) {
	input := &invoicing.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_invoicingResourceArn) > 0 {
		input.ResourceArn = aws.String(_invoicingResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing procurement portal preference configuration. This operation
// can modify settings for e-invoice delivery and purchase order retrieval.
func invoicing_PutProcurementPortalPreference(cfg aws.Config, client *invoicing.Client) {
	input := &invoicing.PutProcurementPortalPreferenceInput{
		// Contacts: []types.Contact, // Required
		// EinvoiceDeliveryEnabled: *bool, // Required
		// ProcurementPortalPreferenceArn: *string, // Required
		// PurchaseOrderRetrievalEnabled: *bool, // Required
	}

	if len(_invoicingContacts) > 0 {
		if err := assignInputField(input, "Contacts", _invoicingContacts); err != nil {
			log.Errorf("invalid --contacts: %s", err.Error())
			return
		}
	}
	if len(_invoicingEinvoiceDeliveryEnabled) > 0 {
		if err := assignInputField(input, "EinvoiceDeliveryEnabled", _invoicingEinvoiceDeliveryEnabled); err != nil {
			log.Errorf("invalid --einvoice-delivery-enabled: %s", err.Error())
			return
		}
	}
	if len(_invoicingProcurementPortalPreferenceArn) > 0 {
		input.ProcurementPortalPreferenceArn = aws.String(_invoicingProcurementPortalPreferenceArn)
	}
	if len(_invoicingPurchaseOrderRetrievalEnabled) > 0 {
		if err := assignInputField(input, "PurchaseOrderRetrievalEnabled", _invoicingPurchaseOrderRetrievalEnabled); err != nil {
			log.Errorf("invalid --purchase-order-retrieval-enabled: %s", err.Error())
			return
		}
	}
	if len(_invoicingEinvoiceDeliveryPreference) > 0 {
		if err := assignInputField(input, "EinvoiceDeliveryPreference", _invoicingEinvoiceDeliveryPreference); err != nil {
			log.Errorf("invalid --einvoice-delivery-preference: %s", err.Error())
			return
		}
	}
	if len(_invoicingProcurementPortalInstanceEndpoint) > 0 {
		input.ProcurementPortalInstanceEndpoint = aws.String(_invoicingProcurementPortalInstanceEndpoint)
	}
	if len(_invoicingProcurementPortalSharedSecret) > 0 {
		input.ProcurementPortalSharedSecret = aws.String(_invoicingProcurementPortalSharedSecret)
	}
	if len(_invoicingSelector) > 0 {
		if err := assignInputField(input, "Selector", _invoicingSelector); err != nil {
			log.Errorf("invalid --selector: %s", err.Error())
			return
		}
	}
	if len(_invoicingTestEnvPreference) > 0 {
		if err := assignInputField(input, "TestEnvPreference", _invoicingTestEnvPreference); err != nil {
			log.Errorf("invalid --test-env-preference: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutProcurementPortalPreference(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a tag to a resource.
func invoicing_TagResource(cfg aws.Config, client *invoicing.Client) {
	input := &invoicing.TagResourceInput{
		// ResourceArn: *string, // Required
		// ResourceTags: []types.ResourceTag, // Required
	}

	if len(_invoicingResourceArn) > 0 {
		input.ResourceArn = aws.String(_invoicingResourceArn)
	}
	if len(_invoicingResourceTags) > 0 {
		if err := assignInputField(input, "ResourceTags", _invoicingResourceTags); err != nil {
			log.Errorf("invalid --resource-tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.TagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a tag from a resource.
func invoicing_UntagResource(cfg aws.Config, client *invoicing.Client) {
	input := &invoicing.UntagResourceInput{
		// ResourceArn: *string, // Required
		// ResourceTagKeys: []string, // Required
	}

	if len(_invoicingResourceArn) > 0 {
		input.ResourceArn = aws.String(_invoicingResourceArn)
	}
	if len(_invoicingResourceTagKeys) > 0 {
		input.ResourceTagKeys = append([]string(nil), _invoicingResourceTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// You can update the invoice unit configuration at any time, and Amazon Web
// Services will use the latest configuration at the end of the month.
func invoicing_UpdateInvoiceUnit(cfg aws.Config, client *invoicing.Client) {
	input := &invoicing.UpdateInvoiceUnitInput{
		// InvoiceUnitArn: *string, // Required
	}

	if len(_invoicingInvoiceUnitArn) > 0 {
		input.InvoiceUnitArn = aws.String(_invoicingInvoiceUnitArn)
	}
	if len(_invoicingDescription) > 0 {
		input.Description = aws.String(_invoicingDescription)
	}
	if len(_invoicingRule) > 0 {
		if err := assignInputField(input, "Rule", _invoicingRule); err != nil {
			log.Errorf("invalid --rule: %s", err.Error())
			return
		}
	}
	if len(_invoicingTaxInheritanceDisabled) > 0 {
		if err := assignInputField(input, "TaxInheritanceDisabled", _invoicingTaxInheritanceDisabled); err != nil {
			log.Errorf("invalid --tax-inheritance-disabled: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateInvoiceUnit(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the status of a procurement portal preference, including the activation
// state of e-invoice delivery and purchase order retrieval features.
func invoicing_UpdateProcurementPortalPreferenceStatus(cfg aws.Config, client *invoicing.Client) {
	input := &invoicing.UpdateProcurementPortalPreferenceStatusInput{
		// ProcurementPortalPreferenceArn: *string, // Required
	}

	if len(_invoicingProcurementPortalPreferenceArn) > 0 {
		input.ProcurementPortalPreferenceArn = aws.String(_invoicingProcurementPortalPreferenceArn)
	}
	if len(_invoicingEinvoiceDeliveryPreferenceStatus) > 0 {
		if err := assignInputField(input, "EinvoiceDeliveryPreferenceStatus", _invoicingEinvoiceDeliveryPreferenceStatus); err != nil {
			log.Errorf("invalid --einvoice-delivery-preference-status: %s", err.Error())
			return
		}
	}
	if len(_invoicingEinvoiceDeliveryPreferenceStatusReason) > 0 {
		input.EinvoiceDeliveryPreferenceStatusReason = aws.String(_invoicingEinvoiceDeliveryPreferenceStatusReason)
	}
	if len(_invoicingPurchaseOrderRetrievalPreferenceStatus) > 0 {
		if err := assignInputField(input, "PurchaseOrderRetrievalPreferenceStatus", _invoicingPurchaseOrderRetrievalPreferenceStatus); err != nil {
			log.Errorf("invalid --purchase-order-retrieval-preference-status: %s", err.Error())
			return
		}
	}
	if len(_invoicingPurchaseOrderRetrievalPreferenceStatusReason) > 0 {
		input.PurchaseOrderRetrievalPreferenceStatusReason = aws.String(_invoicingPurchaseOrderRetrievalPreferenceStatusReason)
	}

	if resp, err := client.UpdateProcurementPortalPreferenceStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_invoicingCmd)
	_invoicingCmd.Flags().SortFlags = false

	_invoicingCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_invoicingCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_invoicingCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_invoicingCmd.Flags().StringSliceVarP(&_invoicingAccountIds, "account-ids", "", nil, "Account Ids")
	_invoicingCmd.Flags().StringVarP(&_invoicingAsOf, "as-of", "", "", "As Of")
	_invoicingCmd.Flags().StringVarP(&_invoicingBuyerDomain, "buyer-domain", "", "", "Buyer Domain")
	_invoicingCmd.Flags().StringVarP(&_invoicingBuyerIdentifier, "buyer-identifier", "", "", "Buyer Identifier")
	_invoicingCmd.Flags().StringVarP(&_invoicingClientToken, "client-token", "", "", "Client Token")
	_invoicingCmd.Flags().StringVarP(&_invoicingContacts, "contacts", "", "", "Contacts")
	_invoicingCmd.Flags().StringVarP(&_invoicingDescription, "description", "", "", "Description")
	_invoicingCmd.Flags().StringVarP(&_invoicingEinvoiceDeliveryEnabled, "einvoice-delivery-enabled", "", "", "Einvoice Delivery Enabled")
	_invoicingCmd.Flags().StringVarP(&_invoicingEinvoiceDeliveryPreference, "einvoice-delivery-preference", "", "", "Einvoice Delivery Preference")
	_invoicingCmd.Flags().StringVarP(&_invoicingEinvoiceDeliveryPreferenceStatus, "einvoice-delivery-preference-status", "", "", "Einvoice Delivery Preference Status")
	_invoicingCmd.Flags().StringVarP(&_invoicingEinvoiceDeliveryPreferenceStatusReason, "einvoice-delivery-preference-status-reason", "", "", "Einvoice Delivery Preference Status Reason")
	_invoicingCmd.Flags().StringVarP(&_invoicingFilter, "filter", "", "", "Filter")
	_invoicingCmd.Flags().StringVarP(&_invoicingFilters, "filters", "", "", "Filters")
	_invoicingCmd.Flags().StringVarP(&_invoicingInvoiceId, "invoice-id", "", "", "Invoice ID")
	_invoicingCmd.Flags().StringVarP(&_invoicingInvoiceReceiver, "invoice-receiver", "", "", "Invoice Receiver")
	_invoicingCmd.Flags().StringVarP(&_invoicingInvoiceUnitArn, "invoice-unit-arn", "", "", "Invoice Unit ARN")
	_invoicingCmd.Flags().StringVarP(&_invoicingMaxResults, "max-results", "", "", "Max Results")
	_invoicingCmd.Flags().StringVarP(&_invoicingName, "name", "", "", "Name")
	_invoicingCmd.Flags().StringVarP(&_invoicingNextToken, "next-token", "", "", "Next Token")
	_invoicingCmd.Flags().StringVarP(&_invoicingProcurementPortalInstanceEndpoint, "procurement-portal-instance-endpoint", "", "", "Procurement Portal Instance Endpoint")
	_invoicingCmd.Flags().StringVarP(&_invoicingProcurementPortalName, "procurement-portal-name", "", "", "Procurement Portal Name")
	_invoicingCmd.Flags().StringVarP(&_invoicingProcurementPortalPreferenceArn, "procurement-portal-preference-arn", "", "", "Procurement Portal Preference ARN")
	_invoicingCmd.Flags().StringVarP(&_invoicingProcurementPortalSharedSecret, "procurement-portal-shared-secret", "", "", "Procurement Portal Shared Secret")
	_invoicingCmd.Flags().StringVarP(&_invoicingPurchaseOrderRetrievalEnabled, "purchase-order-retrieval-enabled", "", "", "Purchase Order Retrieval Enabled")
	_invoicingCmd.Flags().StringVarP(&_invoicingPurchaseOrderRetrievalPreferenceStatus, "purchase-order-retrieval-preference-status", "", "", "Purchase Order Retrieval Preference Status")
	_invoicingCmd.Flags().StringVarP(&_invoicingPurchaseOrderRetrievalPreferenceStatusReason, "purchase-order-retrieval-preference-status-reason", "", "", "Purchase Order Retrieval Preference Status Reason")
	_invoicingCmd.Flags().StringVarP(&_invoicingResourceArn, "resource-arn", "", "", "Resource ARN")
	_invoicingCmd.Flags().StringSliceVarP(&_invoicingResourceTagKeys, "resource-tag-keys", "", nil, "Resource Tag Keys")
	_invoicingCmd.Flags().StringVarP(&_invoicingResourceTags, "resource-tags", "", "", "Resource Tags")
	_invoicingCmd.Flags().StringVarP(&_invoicingRule, "rule", "", "", "Rule")
	_invoicingCmd.Flags().StringVarP(&_invoicingSelector, "selector", "", "", "Selector")
	_invoicingCmd.Flags().StringVarP(&_invoicingSupplierDomain, "supplier-domain", "", "", "Supplier Domain")
	_invoicingCmd.Flags().StringVarP(&_invoicingSupplierIdentifier, "supplier-identifier", "", "", "Supplier Identifier")
	_invoicingCmd.Flags().StringVarP(&_invoicingTaxInheritanceDisabled, "tax-inheritance-disabled", "", "", "Tax Inheritance Disabled")
	_invoicingCmd.Flags().StringVarP(&_invoicingTestEnvPreference, "test-env-preference", "", "", "Test Env Preference")

	_invoicingCmd.Flags().BoolVarP(&_invoicingBatchGetInvoiceProfile, "batch-get-invoice-profile", "", false, "Batch Get Invoice Profile")
	_invoicingCmd.Flags().BoolVarP(&_invoicingCreateInvoiceUnit, "create-invoice-unit", "", false, "Create Invoice Unit")
	_invoicingCmd.Flags().BoolVarP(&_invoicingCreateProcurementPortalPreference, "create-procurement-portal-preference", "", false, "Create Procurement Portal Preference")
	_invoicingCmd.Flags().BoolVarP(&_invoicingDeleteInvoiceUnit, "delete-invoice-unit", "", false, "Delete Invoice Unit")
	_invoicingCmd.Flags().BoolVarP(&_invoicingDeleteProcurementPortalPreference, "delete-procurement-portal-preference", "", false, "Delete Procurement Portal Preference")
	_invoicingCmd.Flags().BoolVarP(&_invoicingGetInvoicePDF, "get-invoice-pdf", "", false, "Get Invoice Pdf")
	_invoicingCmd.Flags().BoolVarP(&_invoicingGetInvoiceUnit, "get-invoice-unit", "", false, "Get Invoice Unit")
	_invoicingCmd.Flags().BoolVarP(&_invoicingGetProcurementPortalPreference, "get-procurement-portal-preference", "", false, "Get Procurement Portal Preference")
	_invoicingCmd.Flags().BoolVarP(&_invoicingListInvoiceSummaries, "list-invoice-summaries", "", false, "List Invoice Summaries")
	_invoicingCmd.Flags().BoolVarP(&_invoicingListInvoiceUnits, "list-invoice-units", "", false, "List Invoice Units")
	_invoicingCmd.Flags().BoolVarP(&_invoicingListProcurementPortalPreferences, "list-procurement-portal-preferences", "", false, "List Procurement Portal Preferences")
	_invoicingCmd.Flags().BoolVarP(&_invoicingListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_invoicingCmd.Flags().BoolVarP(&_invoicingPutProcurementPortalPreference, "put-procurement-portal-preference", "", false, "Put Procurement Portal Preference")
	_invoicingCmd.Flags().BoolVarP(&_invoicingTagResource, "tag-resource", "", false, "Tag Resource")
	_invoicingCmd.Flags().BoolVarP(&_invoicingUntagResource, "untag-resource", "", false, "Untag Resource")
	_invoicingCmd.Flags().BoolVarP(&_invoicingUpdateInvoiceUnit, "update-invoice-unit", "", false, "Update Invoice Unit")
	_invoicingCmd.Flags().BoolVarP(&_invoicingUpdateProcurementPortalPreferenceStatus, "update-procurement-portal-preference-status", "", false, "Update Procurement Portal Preference Status")

}
