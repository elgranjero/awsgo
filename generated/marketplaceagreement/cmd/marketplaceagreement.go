package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/marketplaceagreement"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// marketplaceagreementCmd represents the marketplaceagreement command
var _marketplaceagreementCmd = &cobra.Command{
	Use:   "marketplaceagreement",
	Short: "AWS marketplaceagreement CLI",
	Run: func(cmd *cobra.Command, args []string) {
		_awsOutput = resolveAWSOutput(_awsProfile, cmd.Flags().Changed("output"))
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := marketplaceagreement.NewFromConfig(cfg)
		if _marketplaceagreementDescribeAgreement {
			marketplaceagreement_DescribeAgreement(cfg, client)
			return
		}
		if _marketplaceagreementGetAgreementTerms {
			marketplaceagreement_GetAgreementTerms(cfg, client)
			return
		}
		if _marketplaceagreementSearchAgreements {
			marketplaceagreement_SearchAgreements(cfg, client)
			return
		}

	},
}

var (
	_marketplaceagreementDescribeAgreement bool
	_marketplaceagreementGetAgreementTerms bool
	_marketplaceagreementSearchAgreements  bool

	_marketplaceagreementAgreementId string
	_marketplaceagreementCatalog     string
	_marketplaceagreementFilters     string
	_marketplaceagreementMaxResults  string
	_marketplaceagreementNextToken   string
	_marketplaceagreementSort        string
)

// Provides details about an agreement, such as the proposer, acceptor, start
// date, and end date.
func marketplaceagreement_DescribeAgreement(cfg aws.Config, client *marketplaceagreement.Client) {
	input := &marketplaceagreement.DescribeAgreementInput{
		// AgreementId: *string, // Required
	}

	if len(_marketplaceagreementAgreementId) > 0 {
		input.AgreementId = aws.String(_marketplaceagreementAgreementId)
	}

	if resp, err := client.DescribeAgreement(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Obtains details about the terms in an agreement that you participated in as
// proposer or acceptor.
//
// The details include:
//
// - TermType – The type of term, such as LegalTerm , RenewalTerm , or
// ConfigurableUpfrontPricingTerm .
//
// - TermID – The ID of the particular term, which is common between offer and
// agreement.
//
// - TermPayload – The key information contained in the term, such as the EULA
// for LegalTerm or pricing and dimensions for various pricing terms, such as
// ConfigurableUpfrontPricingTerm or UsageBasedPricingTerm .
//
// - Configuration – The buyer/acceptor's selection at the time of agreement
// creation, such as the number of units purchased for a dimension or setting the
// EnableAutoRenew flag.
func marketplaceagreement_GetAgreementTerms(cfg aws.Config, client *marketplaceagreement.Client) {
	input := &marketplaceagreement.GetAgreementTermsInput{
		// AgreementId: *string, // Required
	}

	if len(_marketplaceagreementAgreementId) > 0 {
		input.AgreementId = aws.String(_marketplaceagreementAgreementId)
	}
	if len(_marketplaceagreementMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _marketplaceagreementMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_marketplaceagreementNextToken) > 0 {
		input.NextToken = aws.String(_marketplaceagreementNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetAgreementTerms(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*marketplaceagreement.GetAgreementTermsOutput
	p := marketplaceagreement.NewGetAgreementTermsPaginator(client, input)
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

// Searches across all agreements that a proposer has in AWS Marketplace. The
// search returns a list of agreements with basic agreement information.
//
// The following filter combinations are supported when the PartyType is Proposer :
//
// - AgreementType
//
// - AgreementType + EndTime
//
// - AgreementType + ResourceType
//
// - AgreementType + ResourceType + EndTime
//
// - AgreementType + ResourceType + Status
//
// - AgreementType + ResourceType + Status + EndTime
//
// - AgreementType + ResourceId
//
// - AgreementType + ResourceId + EndTime
//
// - AgreementType + ResourceId + Status
//
// - AgreementType + ResourceId + Status + EndTime
//
// - AgreementType + AcceptorAccountId
//
// - AgreementType + AcceptorAccountId + EndTime
//
// - AgreementType + AcceptorAccountId + Status
//
// - AgreementType + AcceptorAccountId + Status + EndTime
//
// - AgreementType + AcceptorAccountId + OfferId
//
// - AgreementType + AcceptorAccountId + OfferId + Status
//
// - AgreementType + AcceptorAccountId + OfferId + EndTime
//
// - AgreementType + AcceptorAccountId + OfferId + Status + EndTime
//
// - AgreementType + AcceptorAccountId + ResourceId
//
// - AgreementType + AcceptorAccountId + ResourceId + Status
//
// - AgreementType + AcceptorAccountId + ResourceId + EndTime
//
// - AgreementType + AcceptorAccountId + ResourceId + Status + EndTime
//
// - AgreementType + AcceptorAccountId + ResourceType
//
// - AgreementType + AcceptorAccountId + ResourceType + EndTime
//
// - AgreementType + AcceptorAccountId + ResourceType + Status
//
// - AgreementType + AcceptorAccountId + ResourceType + Status + EndTime
//
// - AgreementType + Status
//
// - AgreementType + Status + EndTime
//
// - AgreementType + OfferId
//
// - AgreementType + OfferId + EndTime
//
// - AgreementType + OfferId + Status
//
// - AgreementType + OfferId + Status + EndTime
//
// - AgreementType + OfferSetId
//
// - AgreementType + OfferSetId + EndTime
//
// - AgreementType + OfferSetId + Status
//
// - AgreementType + OfferSetId + Status + EndTime
//
// To filter by EndTime , you can use either BeforeEndTime or AfterEndTime . Only
// EndTime is supported for sorting.
func marketplaceagreement_SearchAgreements(cfg aws.Config, client *marketplaceagreement.Client) {
	input := &marketplaceagreement.SearchAgreementsInput{}

	if len(_marketplaceagreementCatalog) > 0 {
		input.Catalog = aws.String(_marketplaceagreementCatalog)
	}
	if len(_marketplaceagreementFilters) > 0 {
		if err := assignInputField(input, "Filters", _marketplaceagreementFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_marketplaceagreementMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _marketplaceagreementMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_marketplaceagreementNextToken) > 0 {
		input.NextToken = aws.String(_marketplaceagreementNextToken)
	}
	if len(_marketplaceagreementSort) > 0 {
		if err := assignInputField(input, "Sort", _marketplaceagreementSort); err != nil {
			log.Errorf("invalid --sort: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.SearchAgreements(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*marketplaceagreement.SearchAgreementsOutput
	p := marketplaceagreement.NewSearchAgreementsPaginator(client, input)
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

func init() {
	_rootCmd.AddCommand(_marketplaceagreementCmd)
	_marketplaceagreementCmd.Flags().SortFlags = false

	_marketplaceagreementCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_marketplaceagreementCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_marketplaceagreementCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_marketplaceagreementCmd.Flags().StringVarP(&_marketplaceagreementAgreementId, "agreement-id", "", "", "Agreement ID")
	_marketplaceagreementCmd.Flags().StringVarP(&_marketplaceagreementCatalog, "catalog", "", "", "Catalog")
	_marketplaceagreementCmd.Flags().StringVarP(&_marketplaceagreementFilters, "filters", "", "", "Filters")
	_marketplaceagreementCmd.Flags().StringVarP(&_marketplaceagreementMaxResults, "max-results", "", "", "Max Results")
	_marketplaceagreementCmd.Flags().StringVarP(&_marketplaceagreementNextToken, "next-token", "", "", "Next Token")
	_marketplaceagreementCmd.Flags().StringVarP(&_marketplaceagreementSort, "sort", "", "", "Sort")

	_marketplaceagreementCmd.Flags().BoolVarP(&_marketplaceagreementDescribeAgreement, "describe-agreement", "", false, "Describe Agreement")
	_marketplaceagreementCmd.Flags().BoolVarP(&_marketplaceagreementGetAgreementTerms, "get-agreement-terms", "", false, "Get Agreement Terms")
	_marketplaceagreementCmd.Flags().BoolVarP(&_marketplaceagreementSearchAgreements, "search-agreements", "", false, "Search Agreements")

}
