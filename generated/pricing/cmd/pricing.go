package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/pricing"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// pricingCmd represents the pricing command
var _pricingCmd = &cobra.Command{
	Use:   "pricing",
	Short: "AWS pricing CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := pricing.NewFromConfig(cfg)
		if _pricingDescribeServices {
			pricing_DescribeServices(cfg, client)
			return
		}
		if _pricingGetAttributeValues {
			pricing_GetAttributeValues(cfg, client)
			return
		}
		if _pricingGetPriceListFileUrl {
			pricing_GetPriceListFileUrl(cfg, client)
			return
		}
		if _pricingGetProducts {
			pricing_GetProducts(cfg, client)
			return
		}
		if _pricingListPriceLists {
			pricing_ListPriceLists(cfg, client)
			return
		}

	},
}

var (
	_pricingDescribeServices    bool
	_pricingGetAttributeValues  bool
	_pricingGetPriceListFileUrl bool
	_pricingGetProducts         bool
	_pricingListPriceLists      bool

	_pricingAttributeName string
	_pricingCurrencyCode  string
	_pricingEffectiveDate string
	_pricingFileFormat    string
	_pricingFilters       string
	_pricingFormatVersion string
	_pricingMaxResults    string
	_pricingNextToken     string
	_pricingPriceListArn  string
	_pricingRegionCode    string
	_pricingServiceCode   string
)

// Returns the metadata for one service or a list of the metadata for all
// services. Use this without a service code to get the service codes for all
// services. Use it with a service code, such as AmazonEC2 , to get information
// specific to that service, such as the attribute names available for that
// service. For example, some of the attribute names available for EC2 are
// volumeType , maxIopsVolume , operation , locationType , and
// instanceCapacity10xlarge .
func pricing_DescribeServices(cfg aws.Config, client *pricing.Client) {
	input := &pricing.DescribeServicesInput{}

	if len(_pricingFormatVersion) > 0 {
		input.FormatVersion = aws.String(_pricingFormatVersion)
	}
	if len(_pricingMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _pricingMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_pricingNextToken) > 0 {
		input.NextToken = aws.String(_pricingNextToken)
	}
	if len(_pricingServiceCode) > 0 {
		input.ServiceCode = aws.String(_pricingServiceCode)
	}

	if disablePaginator() {
		if resp, err := client.DescribeServices(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*pricing.DescribeServicesOutput
	p := pricing.NewDescribeServicesPaginator(client, input)
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

// Returns a list of attribute values. Attributes are similar to the details in a
// Price List API offer file. For a list of available attributes, see [Offer File Definitions]in the [Billing and Cost Management User Guide].
//
// [Billing and Cost Management User Guide]: https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/billing-what-is.html
// [Offer File Definitions]: https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/reading-an-offer.html#pps-defs
func pricing_GetAttributeValues(cfg aws.Config, client *pricing.Client) {
	input := &pricing.GetAttributeValuesInput{
		// AttributeName: *string, // Required
		// ServiceCode: *string, // Required
	}

	if len(_pricingAttributeName) > 0 {
		input.AttributeName = aws.String(_pricingAttributeName)
	}
	if len(_pricingServiceCode) > 0 {
		input.ServiceCode = aws.String(_pricingServiceCode)
	}
	if len(_pricingMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _pricingMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_pricingNextToken) > 0 {
		input.NextToken = aws.String(_pricingNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetAttributeValues(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*pricing.GetAttributeValuesOutput
	p := pricing.NewGetAttributeValuesPaginator(client, input)
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

// This feature is in preview release and is subject to change. Your use of
// Amazon Web Services Price List API is subject to the Beta Service Participation
// terms of the [Amazon Web Services Service Terms](Section 1.10).
//
// This returns the URL that you can retrieve your Price List file from. This URL
// is based on the PriceListArn and FileFormat that you retrieve from the [ListPriceLists]
// response.
//
// [Amazon Web Services Service Terms]: https://aws.amazon.com/service-terms/
// [ListPriceLists]: https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_pricing_ListPriceLists.html
func pricing_GetPriceListFileUrl(cfg aws.Config, client *pricing.Client) {
	input := &pricing.GetPriceListFileUrlInput{
		// FileFormat: *string, // Required
		// PriceListArn: *string, // Required
	}

	if len(_pricingFileFormat) > 0 {
		input.FileFormat = aws.String(_pricingFileFormat)
	}
	if len(_pricingPriceListArn) > 0 {
		input.PriceListArn = aws.String(_pricingPriceListArn)
	}

	if resp, err := client.GetPriceListFileUrl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of all products that match the filter criteria.
func pricing_GetProducts(cfg aws.Config, client *pricing.Client) {
	input := &pricing.GetProductsInput{
		// ServiceCode: *string, // Required
	}

	if len(_pricingServiceCode) > 0 {
		input.ServiceCode = aws.String(_pricingServiceCode)
	}
	if len(_pricingFilters) > 0 {
		if err := assignInputField(input, "Filters", _pricingFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_pricingFormatVersion) > 0 {
		input.FormatVersion = aws.String(_pricingFormatVersion)
	}
	if len(_pricingMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _pricingMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_pricingNextToken) > 0 {
		input.NextToken = aws.String(_pricingNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetProducts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*pricing.GetProductsOutput
	p := pricing.NewGetProductsPaginator(client, input)
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

// This feature is in preview release and is subject to change. Your use of
// Amazon Web Services Price List API is subject to the Beta Service Participation
// terms of the [Amazon Web Services Service Terms](Section 1.10).
//
// This returns a list of Price List references that the requester if authorized
// to view, given a ServiceCode , CurrencyCode , and an EffectiveDate . Use without
// a RegionCode filter to list Price List references from all available Amazon Web
// Services Regions. Use with a RegionCode filter to get the Price List reference
// that's specific to a specific Amazon Web Services Region. You can use the
// PriceListArn from the response to get your preferred Price List files through
// the [GetPriceListFileUrl]API.
//
// [Amazon Web Services Service Terms]: https://aws.amazon.com/service-terms/
// [GetPriceListFileUrl]: https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_pricing_GetPriceListFileUrl.html
func pricing_ListPriceLists(cfg aws.Config, client *pricing.Client) {
	input := &pricing.ListPriceListsInput{
		// CurrencyCode: *string, // Required
		// EffectiveDate: *time.Time, // Required
		// ServiceCode: *string, // Required
	}

	if len(_pricingCurrencyCode) > 0 {
		input.CurrencyCode = aws.String(_pricingCurrencyCode)
	}
	if len(_pricingEffectiveDate) > 0 {
		if err := assignInputField(input, "EffectiveDate", _pricingEffectiveDate); err != nil {
			log.Errorf("invalid --effective-date: %s", err.Error())
			return
		}
	}
	if len(_pricingServiceCode) > 0 {
		input.ServiceCode = aws.String(_pricingServiceCode)
	}
	if len(_pricingMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _pricingMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_pricingNextToken) > 0 {
		input.NextToken = aws.String(_pricingNextToken)
	}
	if len(_pricingRegionCode) > 0 {
		input.RegionCode = aws.String(_pricingRegionCode)
	}

	if disablePaginator() {
		if resp, err := client.ListPriceLists(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*pricing.ListPriceListsOutput
	p := pricing.NewListPriceListsPaginator(client, input)
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
	_rootCmd.AddCommand(_pricingCmd)
	_pricingCmd.Flags().SortFlags = false

	_pricingCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_pricingCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_pricingCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_pricingCmd.Flags().StringVarP(&_pricingAttributeName, "attribute-name", "", "", "Attribute Name")
	_pricingCmd.Flags().StringVarP(&_pricingCurrencyCode, "currency-code", "", "", "Currency Code")
	_pricingCmd.Flags().StringVarP(&_pricingEffectiveDate, "effective-date", "", "", "Effective Date")
	_pricingCmd.Flags().StringVarP(&_pricingFileFormat, "file-format", "", "", "File Format")
	_pricingCmd.Flags().StringVarP(&_pricingFilters, "filters", "", "", "Filters")
	_pricingCmd.Flags().StringVarP(&_pricingFormatVersion, "format-version", "", "", "Format Version")
	_pricingCmd.Flags().StringVarP(&_pricingMaxResults, "max-results", "", "", "Max Results")
	_pricingCmd.Flags().StringVarP(&_pricingNextToken, "next-token", "", "", "Next Token")
	_pricingCmd.Flags().StringVarP(&_pricingPriceListArn, "price-list-arn", "", "", "Price List ARN")
	_pricingCmd.Flags().StringVarP(&_pricingRegionCode, "region-code", "", "", "Region Code")
	_pricingCmd.Flags().StringVarP(&_pricingServiceCode, "service-code", "", "", "Service Code")

	_pricingCmd.Flags().BoolVarP(&_pricingDescribeServices, "describe-services", "", false, "Describe Services")
	_pricingCmd.Flags().BoolVarP(&_pricingGetAttributeValues, "get-attribute-values", "", false, "Get Attribute Values")
	_pricingCmd.Flags().BoolVarP(&_pricingGetPriceListFileUrl, "get-price-list-file-url", "", false, "Get Price List File URL")
	_pricingCmd.Flags().BoolVarP(&_pricingGetProducts, "get-products", "", false, "Get Products")
	_pricingCmd.Flags().BoolVarP(&_pricingListPriceLists, "list-price-lists", "", false, "List Price Lists")

}
