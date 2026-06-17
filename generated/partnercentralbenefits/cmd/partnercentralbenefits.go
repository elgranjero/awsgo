package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/partnercentralbenefits"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// partnercentralbenefitsCmd represents the partnercentralbenefits command
var _partnercentralbenefitsCmd = &cobra.Command{
	Use:   "partnercentralbenefits",
	Short: "AWS partnercentralbenefits CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := partnercentralbenefits.NewFromConfig(cfg)
		if _partnercentralbenefitsAmendBenefitApplication {
			partnercentralbenefits_AmendBenefitApplication(cfg, client)
			return
		}
		if _partnercentralbenefitsAssociateBenefitApplicationResource {
			partnercentralbenefits_AssociateBenefitApplicationResource(cfg, client)
			return
		}
		if _partnercentralbenefitsCancelBenefitApplication {
			partnercentralbenefits_CancelBenefitApplication(cfg, client)
			return
		}
		if _partnercentralbenefitsCreateBenefitApplication {
			partnercentralbenefits_CreateBenefitApplication(cfg, client)
			return
		}
		if _partnercentralbenefitsDisassociateBenefitApplicationResource {
			partnercentralbenefits_DisassociateBenefitApplicationResource(cfg, client)
			return
		}
		if _partnercentralbenefitsGetBenefit {
			partnercentralbenefits_GetBenefit(cfg, client)
			return
		}
		if _partnercentralbenefitsGetBenefitAllocation {
			partnercentralbenefits_GetBenefitAllocation(cfg, client)
			return
		}
		if _partnercentralbenefitsGetBenefitApplication {
			partnercentralbenefits_GetBenefitApplication(cfg, client)
			return
		}
		if _partnercentralbenefitsListBenefitAllocations {
			partnercentralbenefits_ListBenefitAllocations(cfg, client)
			return
		}
		if _partnercentralbenefitsListBenefitApplications {
			partnercentralbenefits_ListBenefitApplications(cfg, client)
			return
		}
		if _partnercentralbenefitsListBenefits {
			partnercentralbenefits_ListBenefits(cfg, client)
			return
		}
		if _partnercentralbenefitsListTagsForResource {
			partnercentralbenefits_ListTagsForResource(cfg, client)
			return
		}
		if _partnercentralbenefitsRecallBenefitApplication {
			partnercentralbenefits_RecallBenefitApplication(cfg, client)
			return
		}
		if _partnercentralbenefitsSubmitBenefitApplication {
			partnercentralbenefits_SubmitBenefitApplication(cfg, client)
			return
		}
		if _partnercentralbenefitsTagResource {
			partnercentralbenefits_TagResource(cfg, client)
			return
		}
		if _partnercentralbenefitsUntagResource {
			partnercentralbenefits_UntagResource(cfg, client)
			return
		}
		if _partnercentralbenefitsUpdateBenefitApplication {
			partnercentralbenefits_UpdateBenefitApplication(cfg, client)
			return
		}

	},
}

var (
	_partnercentralbenefitsAmendBenefitApplication                bool
	_partnercentralbenefitsAssociateBenefitApplicationResource    bool
	_partnercentralbenefitsCancelBenefitApplication               bool
	_partnercentralbenefitsCreateBenefitApplication               bool
	_partnercentralbenefitsDisassociateBenefitApplicationResource bool
	_partnercentralbenefitsGetBenefit                             bool
	_partnercentralbenefitsGetBenefitAllocation                   bool
	_partnercentralbenefitsGetBenefitApplication                  bool
	_partnercentralbenefitsListBenefitAllocations                 bool
	_partnercentralbenefitsListBenefitApplications                bool
	_partnercentralbenefitsListBenefits                           bool
	_partnercentralbenefitsListTagsForResource                    bool
	_partnercentralbenefitsRecallBenefitApplication               bool
	_partnercentralbenefitsSubmitBenefitApplication               bool
	_partnercentralbenefitsTagResource                            bool
	_partnercentralbenefitsUntagResource                          bool
	_partnercentralbenefitsUpdateBenefitApplication               bool

	_partnercentralbenefitsAmendmentReason               string
	_partnercentralbenefitsAmendments                    string
	_partnercentralbenefitsAssociatedResourceArns        []string
	_partnercentralbenefitsAssociatedResources           string
	_partnercentralbenefitsBenefitApplicationDetails     string
	_partnercentralbenefitsBenefitApplicationIdentifier  string
	_partnercentralbenefitsBenefitApplicationIdentifiers []string
	_partnercentralbenefitsBenefitIdentifier             string
	_partnercentralbenefitsBenefitIdentifiers            []string
	_partnercentralbenefitsCatalog                       string
	_partnercentralbenefitsClientToken                   string
	_partnercentralbenefitsDescription                   string
	_partnercentralbenefitsFileDetails                   string
	_partnercentralbenefitsFulfillmentTypes              string
	_partnercentralbenefitsIdentifier                    string
	_partnercentralbenefitsMaxResults                    string
	_partnercentralbenefitsName                          string
	_partnercentralbenefitsNextToken                     string
	_partnercentralbenefitsPartnerContacts               string
	_partnercentralbenefitsPrograms                      []string
	_partnercentralbenefitsReason                        string
	_partnercentralbenefitsResourceArn                   string
	_partnercentralbenefitsRevision                      string
	_partnercentralbenefitsStages                        []string
	_partnercentralbenefitsStatus                        string
	_partnercentralbenefitsTagKeys                       []string
	_partnercentralbenefitsTags                          string
)

// Modifies an existing benefit application by applying amendments to specific
// fields while maintaining revision control.
func partnercentralbenefits_AmendBenefitApplication(cfg aws.Config, client *partnercentralbenefits.Client) {
	input := &partnercentralbenefits.AmendBenefitApplicationInput{
		// AmendmentReason: *string, // Required
		// Amendments: []types.Amendment, // Required
		// Catalog: *string, // Required
		// ClientToken: *string, // Required
		// Identifier: *string, // Required
		// Revision: *string, // Required
	}

	if len(_partnercentralbenefitsAmendmentReason) > 0 {
		input.AmendmentReason = aws.String(_partnercentralbenefitsAmendmentReason)
	}
	if len(_partnercentralbenefitsAmendments) > 0 {
		if err := assignInputField(input, "Amendments", _partnercentralbenefitsAmendments); err != nil {
			log.Errorf("invalid --amendments: %s", err.Error())
			return
		}
	}
	if len(_partnercentralbenefitsCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralbenefitsCatalog)
	}
	if len(_partnercentralbenefitsClientToken) > 0 {
		input.ClientToken = aws.String(_partnercentralbenefitsClientToken)
	}
	if len(_partnercentralbenefitsIdentifier) > 0 {
		input.Identifier = aws.String(_partnercentralbenefitsIdentifier)
	}
	if len(_partnercentralbenefitsRevision) > 0 {
		input.Revision = aws.String(_partnercentralbenefitsRevision)
	}

	if resp, err := client.AmendBenefitApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Links an AWS resource to an existing benefit application for tracking and
// management purposes.
func partnercentralbenefits_AssociateBenefitApplicationResource(cfg aws.Config, client *partnercentralbenefits.Client) {
	input := &partnercentralbenefits.AssociateBenefitApplicationResourceInput{
		// BenefitApplicationIdentifier: *string, // Required
		// Catalog: *string, // Required
		// ResourceArn: *string, // Required
	}

	if len(_partnercentralbenefitsBenefitApplicationIdentifier) > 0 {
		input.BenefitApplicationIdentifier = aws.String(_partnercentralbenefitsBenefitApplicationIdentifier)
	}
	if len(_partnercentralbenefitsCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralbenefitsCatalog)
	}
	if len(_partnercentralbenefitsResourceArn) > 0 {
		input.ResourceArn = aws.String(_partnercentralbenefitsResourceArn)
	}

	if resp, err := client.AssociateBenefitApplicationResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels a benefit application that is currently in progress, preventing further
// processing.
func partnercentralbenefits_CancelBenefitApplication(cfg aws.Config, client *partnercentralbenefits.Client) {
	input := &partnercentralbenefits.CancelBenefitApplicationInput{
		// Catalog: *string, // Required
		// ClientToken: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_partnercentralbenefitsCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralbenefitsCatalog)
	}
	if len(_partnercentralbenefitsClientToken) > 0 {
		input.ClientToken = aws.String(_partnercentralbenefitsClientToken)
	}
	if len(_partnercentralbenefitsIdentifier) > 0 {
		input.Identifier = aws.String(_partnercentralbenefitsIdentifier)
	}
	if len(_partnercentralbenefitsReason) > 0 {
		input.Reason = aws.String(_partnercentralbenefitsReason)
	}

	if resp, err := client.CancelBenefitApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new benefit application for a partner to request access to AWS
// benefits and programs.
func partnercentralbenefits_CreateBenefitApplication(cfg aws.Config, client *partnercentralbenefits.Client) {
	input := &partnercentralbenefits.CreateBenefitApplicationInput{
		// BenefitIdentifier: *string, // Required
		// Catalog: *string, // Required
		// ClientToken: *string, // Required
	}

	if len(_partnercentralbenefitsBenefitIdentifier) > 0 {
		input.BenefitIdentifier = aws.String(_partnercentralbenefitsBenefitIdentifier)
	}
	if len(_partnercentralbenefitsCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralbenefitsCatalog)
	}
	if len(_partnercentralbenefitsClientToken) > 0 {
		input.ClientToken = aws.String(_partnercentralbenefitsClientToken)
	}
	if len(_partnercentralbenefitsAssociatedResources) > 0 {
		input.AssociatedResources = []string{_partnercentralbenefitsAssociatedResources}
	}
	if len(_partnercentralbenefitsBenefitApplicationDetails) > 0 {
		if err := assignInputField(input, "BenefitApplicationDetails", _partnercentralbenefitsBenefitApplicationDetails); err != nil {
			log.Errorf("invalid --benefit-application-details: %s", err.Error())
			return
		}
	}
	if len(_partnercentralbenefitsDescription) > 0 {
		input.Description = aws.String(_partnercentralbenefitsDescription)
	}
	if len(_partnercentralbenefitsFileDetails) > 0 {
		if err := assignInputField(input, "FileDetails", _partnercentralbenefitsFileDetails); err != nil {
			log.Errorf("invalid --file-details: %s", err.Error())
			return
		}
	}
	if len(_partnercentralbenefitsFulfillmentTypes) > 0 {
		if err := assignInputField(input, "FulfillmentTypes", _partnercentralbenefitsFulfillmentTypes); err != nil {
			log.Errorf("invalid --fulfillment-types: %s", err.Error())
			return
		}
	}
	if len(_partnercentralbenefitsName) > 0 {
		input.Name = aws.String(_partnercentralbenefitsName)
	}
	if len(_partnercentralbenefitsPartnerContacts) > 0 {
		if err := assignInputField(input, "PartnerContacts", _partnercentralbenefitsPartnerContacts); err != nil {
			log.Errorf("invalid --partner-contacts: %s", err.Error())
			return
		}
	}
	if len(_partnercentralbenefitsTags) > 0 {
		if err := assignInputField(input, "Tags", _partnercentralbenefitsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateBenefitApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the association between an AWS resource and a benefit application.
func partnercentralbenefits_DisassociateBenefitApplicationResource(cfg aws.Config, client *partnercentralbenefits.Client) {
	input := &partnercentralbenefits.DisassociateBenefitApplicationResourceInput{
		// BenefitApplicationIdentifier: *string, // Required
		// Catalog: *string, // Required
		// ResourceArn: *string, // Required
	}

	if len(_partnercentralbenefitsBenefitApplicationIdentifier) > 0 {
		input.BenefitApplicationIdentifier = aws.String(_partnercentralbenefitsBenefitApplicationIdentifier)
	}
	if len(_partnercentralbenefitsCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralbenefitsCatalog)
	}
	if len(_partnercentralbenefitsResourceArn) > 0 {
		input.ResourceArn = aws.String(_partnercentralbenefitsResourceArn)
	}

	if resp, err := client.DisassociateBenefitApplicationResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves detailed information about a specific benefit available in the
// partner catalog.
func partnercentralbenefits_GetBenefit(cfg aws.Config, client *partnercentralbenefits.Client) {
	input := &partnercentralbenefits.GetBenefitInput{
		// Catalog: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_partnercentralbenefitsCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralbenefitsCatalog)
	}
	if len(_partnercentralbenefitsIdentifier) > 0 {
		input.Identifier = aws.String(_partnercentralbenefitsIdentifier)
	}

	if resp, err := client.GetBenefit(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves detailed information about a specific benefit allocation that has
// been granted to a partner.
func partnercentralbenefits_GetBenefitAllocation(cfg aws.Config, client *partnercentralbenefits.Client) {
	input := &partnercentralbenefits.GetBenefitAllocationInput{
		// Catalog: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_partnercentralbenefitsCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralbenefitsCatalog)
	}
	if len(_partnercentralbenefitsIdentifier) > 0 {
		input.Identifier = aws.String(_partnercentralbenefitsIdentifier)
	}

	if resp, err := client.GetBenefitAllocation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves detailed information about a specific benefit application.
func partnercentralbenefits_GetBenefitApplication(cfg aws.Config, client *partnercentralbenefits.Client) {
	input := &partnercentralbenefits.GetBenefitApplicationInput{
		// Catalog: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_partnercentralbenefitsCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralbenefitsCatalog)
	}
	if len(_partnercentralbenefitsIdentifier) > 0 {
		input.Identifier = aws.String(_partnercentralbenefitsIdentifier)
	}

	if resp, err := client.GetBenefitApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a paginated list of benefit allocations based on specified filter
// criteria.
func partnercentralbenefits_ListBenefitAllocations(cfg aws.Config, client *partnercentralbenefits.Client) {
	input := &partnercentralbenefits.ListBenefitAllocationsInput{
		// Catalog: *string, // Required
	}

	if len(_partnercentralbenefitsCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralbenefitsCatalog)
	}
	if len(_partnercentralbenefitsBenefitApplicationIdentifiers) > 0 {
		input.BenefitApplicationIdentifiers = append([]string(nil), _partnercentralbenefitsBenefitApplicationIdentifiers...)
	}
	if len(_partnercentralbenefitsBenefitIdentifiers) > 0 {
		input.BenefitIdentifiers = append([]string(nil), _partnercentralbenefitsBenefitIdentifiers...)
	}
	if len(_partnercentralbenefitsFulfillmentTypes) > 0 {
		if err := assignInputField(input, "FulfillmentTypes", _partnercentralbenefitsFulfillmentTypes); err != nil {
			log.Errorf("invalid --fulfillment-types: %s", err.Error())
			return
		}
	}
	if len(_partnercentralbenefitsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _partnercentralbenefitsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_partnercentralbenefitsNextToken) > 0 {
		input.NextToken = aws.String(_partnercentralbenefitsNextToken)
	}
	if len(_partnercentralbenefitsStatus) > 0 {
		if err := assignInputField(input, "Status", _partnercentralbenefitsStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListBenefitAllocations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*partnercentralbenefits.ListBenefitAllocationsOutput
	p := partnercentralbenefits.NewListBenefitAllocationsPaginator(client, input)
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

// Retrieves a paginated list of benefit applications based on specified filter
// criteria.
func partnercentralbenefits_ListBenefitApplications(cfg aws.Config, client *partnercentralbenefits.Client) {
	input := &partnercentralbenefits.ListBenefitApplicationsInput{
		// Catalog: *string, // Required
	}

	if len(_partnercentralbenefitsCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralbenefitsCatalog)
	}
	if len(_partnercentralbenefitsAssociatedResourceArns) > 0 {
		input.AssociatedResourceArns = append([]string(nil), _partnercentralbenefitsAssociatedResourceArns...)
	}
	if len(_partnercentralbenefitsAssociatedResources) > 0 {
		if err := assignInputField(input, "AssociatedResources", _partnercentralbenefitsAssociatedResources); err != nil {
			log.Errorf("invalid --associated-resources: %s", err.Error())
			return
		}
	}
	if len(_partnercentralbenefitsBenefitIdentifiers) > 0 {
		input.BenefitIdentifiers = append([]string(nil), _partnercentralbenefitsBenefitIdentifiers...)
	}
	if len(_partnercentralbenefitsFulfillmentTypes) > 0 {
		if err := assignInputField(input, "FulfillmentTypes", _partnercentralbenefitsFulfillmentTypes); err != nil {
			log.Errorf("invalid --fulfillment-types: %s", err.Error())
			return
		}
	}
	if len(_partnercentralbenefitsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _partnercentralbenefitsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_partnercentralbenefitsNextToken) > 0 {
		input.NextToken = aws.String(_partnercentralbenefitsNextToken)
	}
	if len(_partnercentralbenefitsPrograms) > 0 {
		input.Programs = append([]string(nil), _partnercentralbenefitsPrograms...)
	}
	if len(_partnercentralbenefitsStages) > 0 {
		input.Stages = append([]string(nil), _partnercentralbenefitsStages...)
	}
	if len(_partnercentralbenefitsStatus) > 0 {
		if err := assignInputField(input, "Status", _partnercentralbenefitsStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListBenefitApplications(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*partnercentralbenefits.ListBenefitApplicationsOutput
	p := partnercentralbenefits.NewListBenefitApplicationsPaginator(client, input)
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

// Retrieves a paginated list of available benefits based on specified filter
// criteria.
func partnercentralbenefits_ListBenefits(cfg aws.Config, client *partnercentralbenefits.Client) {
	input := &partnercentralbenefits.ListBenefitsInput{
		// Catalog: *string, // Required
	}

	if len(_partnercentralbenefitsCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralbenefitsCatalog)
	}
	if len(_partnercentralbenefitsFulfillmentTypes) > 0 {
		if err := assignInputField(input, "FulfillmentTypes", _partnercentralbenefitsFulfillmentTypes); err != nil {
			log.Errorf("invalid --fulfillment-types: %s", err.Error())
			return
		}
	}
	if len(_partnercentralbenefitsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _partnercentralbenefitsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_partnercentralbenefitsNextToken) > 0 {
		input.NextToken = aws.String(_partnercentralbenefitsNextToken)
	}
	if len(_partnercentralbenefitsPrograms) > 0 {
		input.Programs = append([]string(nil), _partnercentralbenefitsPrograms...)
	}
	if len(_partnercentralbenefitsStatus) > 0 {
		if err := assignInputField(input, "Status", _partnercentralbenefitsStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListBenefits(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*partnercentralbenefits.ListBenefitsOutput
	p := partnercentralbenefits.NewListBenefitsPaginator(client, input)
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

// Retrieves all tags associated with a specific resource.
func partnercentralbenefits_ListTagsForResource(cfg aws.Config, client *partnercentralbenefits.Client) {
	input := &partnercentralbenefits.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_partnercentralbenefitsResourceArn) > 0 {
		input.ResourceArn = aws.String(_partnercentralbenefitsResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Recalls a submitted benefit application, returning it to draft status for
// further modifications.
func partnercentralbenefits_RecallBenefitApplication(cfg aws.Config, client *partnercentralbenefits.Client) {
	input := &partnercentralbenefits.RecallBenefitApplicationInput{
		// Catalog: *string, // Required
		// Identifier: *string, // Required
		// Reason: *string, // Required
	}

	if len(_partnercentralbenefitsCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralbenefitsCatalog)
	}
	if len(_partnercentralbenefitsIdentifier) > 0 {
		input.Identifier = aws.String(_partnercentralbenefitsIdentifier)
	}
	if len(_partnercentralbenefitsReason) > 0 {
		input.Reason = aws.String(_partnercentralbenefitsReason)
	}
	if len(_partnercentralbenefitsClientToken) > 0 {
		input.ClientToken = aws.String(_partnercentralbenefitsClientToken)
	}

	if resp, err := client.RecallBenefitApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Submits a benefit application for review and processing by AWS.
func partnercentralbenefits_SubmitBenefitApplication(cfg aws.Config, client *partnercentralbenefits.Client) {
	input := &partnercentralbenefits.SubmitBenefitApplicationInput{
		// Catalog: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_partnercentralbenefitsCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralbenefitsCatalog)
	}
	if len(_partnercentralbenefitsIdentifier) > 0 {
		input.Identifier = aws.String(_partnercentralbenefitsIdentifier)
	}

	if resp, err := client.SubmitBenefitApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or updates tags for a specified resource.
func partnercentralbenefits_TagResource(cfg aws.Config, client *partnercentralbenefits.Client) {
	input := &partnercentralbenefits.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_partnercentralbenefitsResourceArn) > 0 {
		input.ResourceArn = aws.String(_partnercentralbenefitsResourceArn)
	}
	if len(_partnercentralbenefitsTags) > 0 {
		if err := assignInputField(input, "Tags", _partnercentralbenefitsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
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

// Removes specified tags from a resource.
func partnercentralbenefits_UntagResource(cfg aws.Config, client *partnercentralbenefits.Client) {
	input := &partnercentralbenefits.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_partnercentralbenefitsResourceArn) > 0 {
		input.ResourceArn = aws.String(_partnercentralbenefitsResourceArn)
	}
	if len(_partnercentralbenefitsTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _partnercentralbenefitsTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing benefit application with new information while maintaining
// revision control.
func partnercentralbenefits_UpdateBenefitApplication(cfg aws.Config, client *partnercentralbenefits.Client) {
	input := &partnercentralbenefits.UpdateBenefitApplicationInput{
		// Catalog: *string, // Required
		// ClientToken: *string, // Required
		// Identifier: *string, // Required
		// Revision: *string, // Required
	}

	if len(_partnercentralbenefitsCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralbenefitsCatalog)
	}
	if len(_partnercentralbenefitsClientToken) > 0 {
		input.ClientToken = aws.String(_partnercentralbenefitsClientToken)
	}
	if len(_partnercentralbenefitsIdentifier) > 0 {
		input.Identifier = aws.String(_partnercentralbenefitsIdentifier)
	}
	if len(_partnercentralbenefitsRevision) > 0 {
		input.Revision = aws.String(_partnercentralbenefitsRevision)
	}
	if len(_partnercentralbenefitsBenefitApplicationDetails) > 0 {
		if err := assignInputField(input, "BenefitApplicationDetails", _partnercentralbenefitsBenefitApplicationDetails); err != nil {
			log.Errorf("invalid --benefit-application-details: %s", err.Error())
			return
		}
	}
	if len(_partnercentralbenefitsDescription) > 0 {
		input.Description = aws.String(_partnercentralbenefitsDescription)
	}
	if len(_partnercentralbenefitsFileDetails) > 0 {
		if err := assignInputField(input, "FileDetails", _partnercentralbenefitsFileDetails); err != nil {
			log.Errorf("invalid --file-details: %s", err.Error())
			return
		}
	}
	if len(_partnercentralbenefitsName) > 0 {
		input.Name = aws.String(_partnercentralbenefitsName)
	}
	if len(_partnercentralbenefitsPartnerContacts) > 0 {
		if err := assignInputField(input, "PartnerContacts", _partnercentralbenefitsPartnerContacts); err != nil {
			log.Errorf("invalid --partner-contacts: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateBenefitApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_partnercentralbenefitsCmd)
	_partnercentralbenefitsCmd.Flags().SortFlags = false

	_partnercentralbenefitsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_partnercentralbenefitsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_partnercentralbenefitsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_partnercentralbenefitsCmd.Flags().StringVarP(&_partnercentralbenefitsAmendmentReason, "amendment-reason", "", "", "Amendment Reason")
	_partnercentralbenefitsCmd.Flags().StringVarP(&_partnercentralbenefitsAmendments, "amendments", "", "", "Amendments")
	_partnercentralbenefitsCmd.Flags().StringSliceVarP(&_partnercentralbenefitsAssociatedResourceArns, "associated-resource-arns", "", nil, "Associated Resource Arns")
	_partnercentralbenefitsCmd.Flags().StringVarP(&_partnercentralbenefitsAssociatedResources, "associated-resources", "", "", "Associated Resources")
	_partnercentralbenefitsCmd.Flags().StringVarP(&_partnercentralbenefitsBenefitApplicationDetails, "benefit-application-details", "", "", "Benefit Application Details")
	_partnercentralbenefitsCmd.Flags().StringVarP(&_partnercentralbenefitsBenefitApplicationIdentifier, "benefit-application-identifier", "", "", "Benefit Application Identifier")
	_partnercentralbenefitsCmd.Flags().StringSliceVarP(&_partnercentralbenefitsBenefitApplicationIdentifiers, "benefit-application-identifiers", "", nil, "Benefit Application Identifiers")
	_partnercentralbenefitsCmd.Flags().StringVarP(&_partnercentralbenefitsBenefitIdentifier, "benefit-identifier", "", "", "Benefit Identifier")
	_partnercentralbenefitsCmd.Flags().StringSliceVarP(&_partnercentralbenefitsBenefitIdentifiers, "benefit-identifiers", "", nil, "Benefit Identifiers")
	_partnercentralbenefitsCmd.Flags().StringVarP(&_partnercentralbenefitsCatalog, "catalog", "", "", "Catalog")
	_partnercentralbenefitsCmd.Flags().StringVarP(&_partnercentralbenefitsClientToken, "client-token", "", "", "Client Token")
	_partnercentralbenefitsCmd.Flags().StringVarP(&_partnercentralbenefitsDescription, "description", "", "", "Description")
	_partnercentralbenefitsCmd.Flags().StringVarP(&_partnercentralbenefitsFileDetails, "file-details", "", "", "File Details")
	_partnercentralbenefitsCmd.Flags().StringVarP(&_partnercentralbenefitsFulfillmentTypes, "fulfillment-types", "", "", "Fulfillment Types")
	_partnercentralbenefitsCmd.Flags().StringVarP(&_partnercentralbenefitsIdentifier, "identifier", "", "", "Identifier")
	_partnercentralbenefitsCmd.Flags().StringVarP(&_partnercentralbenefitsMaxResults, "max-results", "", "", "Max Results")
	_partnercentralbenefitsCmd.Flags().StringVarP(&_partnercentralbenefitsName, "name", "", "", "Name")
	_partnercentralbenefitsCmd.Flags().StringVarP(&_partnercentralbenefitsNextToken, "next-token", "", "", "Next Token")
	_partnercentralbenefitsCmd.Flags().StringVarP(&_partnercentralbenefitsPartnerContacts, "partner-contacts", "", "", "Partner Contacts")
	_partnercentralbenefitsCmd.Flags().StringSliceVarP(&_partnercentralbenefitsPrograms, "programs", "", nil, "Programs")
	_partnercentralbenefitsCmd.Flags().StringVarP(&_partnercentralbenefitsReason, "reason", "", "", "Reason")
	_partnercentralbenefitsCmd.Flags().StringVarP(&_partnercentralbenefitsResourceArn, "resource-arn", "", "", "Resource ARN")
	_partnercentralbenefitsCmd.Flags().StringVarP(&_partnercentralbenefitsRevision, "revision", "", "", "Revision")
	_partnercentralbenefitsCmd.Flags().StringSliceVarP(&_partnercentralbenefitsStages, "stages", "", nil, "Stages")
	_partnercentralbenefitsCmd.Flags().StringVarP(&_partnercentralbenefitsStatus, "status", "", "", "Status")
	_partnercentralbenefitsCmd.Flags().StringSliceVarP(&_partnercentralbenefitsTagKeys, "tag-keys", "", nil, "Tag Keys")
	_partnercentralbenefitsCmd.Flags().StringVarP(&_partnercentralbenefitsTags, "tags", "", "", "Tags")

	_partnercentralbenefitsCmd.Flags().BoolVarP(&_partnercentralbenefitsAmendBenefitApplication, "amend-benefit-application", "", false, "Amend Benefit Application")
	_partnercentralbenefitsCmd.Flags().BoolVarP(&_partnercentralbenefitsAssociateBenefitApplicationResource, "associate-benefit-application-resource", "", false, "Associate Benefit Application Resource")
	_partnercentralbenefitsCmd.Flags().BoolVarP(&_partnercentralbenefitsCancelBenefitApplication, "cancel-benefit-application", "", false, "Cancel Benefit Application")
	_partnercentralbenefitsCmd.Flags().BoolVarP(&_partnercentralbenefitsCreateBenefitApplication, "create-benefit-application", "", false, "Create Benefit Application")
	_partnercentralbenefitsCmd.Flags().BoolVarP(&_partnercentralbenefitsDisassociateBenefitApplicationResource, "disassociate-benefit-application-resource", "", false, "Disassociate Benefit Application Resource")
	_partnercentralbenefitsCmd.Flags().BoolVarP(&_partnercentralbenefitsGetBenefit, "get-benefit", "", false, "Get Benefit")
	_partnercentralbenefitsCmd.Flags().BoolVarP(&_partnercentralbenefitsGetBenefitAllocation, "get-benefit-allocation", "", false, "Get Benefit Allocation")
	_partnercentralbenefitsCmd.Flags().BoolVarP(&_partnercentralbenefitsGetBenefitApplication, "get-benefit-application", "", false, "Get Benefit Application")
	_partnercentralbenefitsCmd.Flags().BoolVarP(&_partnercentralbenefitsListBenefitAllocations, "list-benefit-allocations", "", false, "List Benefit Allocations")
	_partnercentralbenefitsCmd.Flags().BoolVarP(&_partnercentralbenefitsListBenefitApplications, "list-benefit-applications", "", false, "List Benefit Applications")
	_partnercentralbenefitsCmd.Flags().BoolVarP(&_partnercentralbenefitsListBenefits, "list-benefits", "", false, "List Benefits")
	_partnercentralbenefitsCmd.Flags().BoolVarP(&_partnercentralbenefitsListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_partnercentralbenefitsCmd.Flags().BoolVarP(&_partnercentralbenefitsRecallBenefitApplication, "recall-benefit-application", "", false, "Recall Benefit Application")
	_partnercentralbenefitsCmd.Flags().BoolVarP(&_partnercentralbenefitsSubmitBenefitApplication, "submit-benefit-application", "", false, "Submit Benefit Application")
	_partnercentralbenefitsCmd.Flags().BoolVarP(&_partnercentralbenefitsTagResource, "tag-resource", "", false, "Tag Resource")
	_partnercentralbenefitsCmd.Flags().BoolVarP(&_partnercentralbenefitsUntagResource, "untag-resource", "", false, "Untag Resource")
	_partnercentralbenefitsCmd.Flags().BoolVarP(&_partnercentralbenefitsUpdateBenefitApplication, "update-benefit-application", "", false, "Update Benefit Application")

}
