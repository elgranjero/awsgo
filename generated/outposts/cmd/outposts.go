package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/outposts"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// outpostsCmd represents the outposts command
var _outpostsCmd = &cobra.Command{
	Use:   "outposts",
	Short: "AWS outposts CLI",
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
		client := outposts.NewFromConfig(cfg)
		if _outpostsCancelCapacityTask {
			outposts_CancelCapacityTask(cfg, client)
			return
		}
		if _outpostsCancelOrder {
			outposts_CancelOrder(cfg, client)
			return
		}
		if _outpostsCreateOrder {
			outposts_CreateOrder(cfg, client)
			return
		}
		if _outpostsCreateOutpost {
			outposts_CreateOutpost(cfg, client)
			return
		}
		if _outpostsCreateSite {
			outposts_CreateSite(cfg, client)
			return
		}
		if _outpostsDeleteOutpost {
			outposts_DeleteOutpost(cfg, client)
			return
		}
		if _outpostsDeleteSite {
			outposts_DeleteSite(cfg, client)
			return
		}
		if _outpostsGetCapacityTask {
			outposts_GetCapacityTask(cfg, client)
			return
		}
		if _outpostsGetCatalogItem {
			outposts_GetCatalogItem(cfg, client)
			return
		}
		if _outpostsGetConnection {
			outposts_GetConnection(cfg, client)
			return
		}
		if _outpostsGetOrder {
			outposts_GetOrder(cfg, client)
			return
		}
		if _outpostsGetOutpost {
			outposts_GetOutpost(cfg, client)
			return
		}
		if _outpostsGetOutpostBillingInformation {
			outposts_GetOutpostBillingInformation(cfg, client)
			return
		}
		if _outpostsGetOutpostInstanceTypes {
			outposts_GetOutpostInstanceTypes(cfg, client)
			return
		}
		if _outpostsGetOutpostSupportedInstanceTypes {
			outposts_GetOutpostSupportedInstanceTypes(cfg, client)
			return
		}
		if _outpostsGetSite {
			outposts_GetSite(cfg, client)
			return
		}
		if _outpostsGetSiteAddress {
			outposts_GetSiteAddress(cfg, client)
			return
		}
		if _outpostsListAssetInstances {
			outposts_ListAssetInstances(cfg, client)
			return
		}
		if _outpostsListAssets {
			outposts_ListAssets(cfg, client)
			return
		}
		if _outpostsListBlockingInstancesForCapacityTask {
			outposts_ListBlockingInstancesForCapacityTask(cfg, client)
			return
		}
		if _outpostsListCapacityTasks {
			outposts_ListCapacityTasks(cfg, client)
			return
		}
		if _outpostsListCatalogItems {
			outposts_ListCatalogItems(cfg, client)
			return
		}
		if _outpostsListOrders {
			outposts_ListOrders(cfg, client)
			return
		}
		if _outpostsListOutposts {
			outposts_ListOutposts(cfg, client)
			return
		}
		if _outpostsListSites {
			outposts_ListSites(cfg, client)
			return
		}
		if _outpostsListTagsForResource {
			outposts_ListTagsForResource(cfg, client)
			return
		}
		if _outpostsStartCapacityTask {
			outposts_StartCapacityTask(cfg, client)
			return
		}
		if _outpostsStartConnection {
			outposts_StartConnection(cfg, client)
			return
		}
		if _outpostsStartOutpostDecommission {
			outposts_StartOutpostDecommission(cfg, client)
			return
		}
		if _outpostsTagResource {
			outposts_TagResource(cfg, client)
			return
		}
		if _outpostsUntagResource {
			outposts_UntagResource(cfg, client)
			return
		}
		if _outpostsUpdateOutpost {
			outposts_UpdateOutpost(cfg, client)
			return
		}
		if _outpostsUpdateSite {
			outposts_UpdateSite(cfg, client)
			return
		}
		if _outpostsUpdateSiteAddress {
			outposts_UpdateSiteAddress(cfg, client)
			return
		}
		if _outpostsUpdateSiteRackPhysicalProperties {
			outposts_UpdateSiteRackPhysicalProperties(cfg, client)
			return
		}

	},
}

var (
	_outpostsCancelCapacityTask                   bool
	_outpostsCancelOrder                          bool
	_outpostsCreateOrder                          bool
	_outpostsCreateOutpost                        bool
	_outpostsCreateSite                           bool
	_outpostsDeleteOutpost                        bool
	_outpostsDeleteSite                           bool
	_outpostsGetCapacityTask                      bool
	_outpostsGetCatalogItem                       bool
	_outpostsGetConnection                        bool
	_outpostsGetOrder                             bool
	_outpostsGetOutpost                           bool
	_outpostsGetOutpostBillingInformation         bool
	_outpostsGetOutpostInstanceTypes              bool
	_outpostsGetOutpostSupportedInstanceTypes     bool
	_outpostsGetSite                              bool
	_outpostsGetSiteAddress                       bool
	_outpostsListAssetInstances                   bool
	_outpostsListAssets                           bool
	_outpostsListBlockingInstancesForCapacityTask bool
	_outpostsListCapacityTasks                    bool
	_outpostsListCatalogItems                     bool
	_outpostsListOrders                           bool
	_outpostsListOutposts                         bool
	_outpostsListSites                            bool
	_outpostsListTagsForResource                  bool
	_outpostsStartCapacityTask                    bool
	_outpostsStartConnection                      bool
	_outpostsStartOutpostDecommission             bool
	_outpostsTagResource                          bool
	_outpostsUntagResource                        bool
	_outpostsUpdateOutpost                        bool
	_outpostsUpdateSite                           bool
	_outpostsUpdateSiteAddress                    bool
	_outpostsUpdateSiteRackPhysicalProperties     bool

	_outpostsAccountIdFilter                     []string
	_outpostsAddress                             string
	_outpostsAddressType                         string
	_outpostsAssetId                             string
	_outpostsAssetIdFilter                       []string
	_outpostsAvailabilityZone                    string
	_outpostsAvailabilityZoneFilter              []string
	_outpostsAvailabilityZoneId                  string
	_outpostsAvailabilityZoneIdFilter            []string
	_outpostsAwsServiceFilter                    string
	_outpostsCapacityTaskId                      string
	_outpostsCapacityTaskStatusFilter            string
	_outpostsCatalogItemId                       string
	_outpostsClientPublicKey                     string
	_outpostsConnectionId                        string
	_outpostsDescription                         string
	_outpostsDeviceSerialNumber                  string
	_outpostsDryRun                              string
	_outpostsEC2FamilyFilter                     []string
	_outpostsFiberOpticCableType                 string
	_outpostsHostIdFilter                        []string
	_outpostsInstancePools                       string
	_outpostsInstanceTypeFilter                  []string
	_outpostsInstancesToExclude                  string
	_outpostsItemClassFilter                     string
	_outpostsLifeCycleStatusFilter               []string
	_outpostsLineItems                           string
	_outpostsMaxResults                          string
	_outpostsMaximumSupportedWeightLbs           string
	_outpostsName                                string
	_outpostsNetworkInterfaceDeviceIndex         string
	_outpostsNextToken                           string
	_outpostsNotes                               string
	_outpostsOperatingAddress                    string
	_outpostsOperatingAddressCityFilter          []string
	_outpostsOperatingAddressCountryCodeFilter   []string
	_outpostsOperatingAddressStateOrRegionFilter []string
	_outpostsOpticalStandard                     string
	_outpostsOrderId                             string
	_outpostsOutpostId                           string
	_outpostsOutpostIdentifier                   string
	_outpostsOutpostIdentifierFilter             string
	_outpostsPaymentOption                       string
	_outpostsPaymentTerm                         string
	_outpostsPowerConnector                      string
	_outpostsPowerDrawKva                        string
	_outpostsPowerFeedDrop                       string
	_outpostsPowerPhase                          string
	_outpostsRackPhysicalProperties              string
	_outpostsResourceArn                         string
	_outpostsShippingAddress                     string
	_outpostsSiteId                              string
	_outpostsStatusFilter                        string
	_outpostsSupportedHardwareType               string
	_outpostsSupportedStorageFilter              string
	_outpostsTagKeys                             []string
	_outpostsTags                                string
	_outpostsTaskActionOnBlockingInstances       string
	_outpostsUplinkCount                         string
	_outpostsUplinkGbps                          string
	_outpostsValidateOnly                        string
)

// Cancels the capacity task.
func outposts_CancelCapacityTask(cfg aws.Config, client *outposts.Client) {
	input := &outposts.CancelCapacityTaskInput{
		// CapacityTaskId: *string, // Required
		// OutpostIdentifier: *string, // Required
	}

	if len(_outpostsCapacityTaskId) > 0 {
		input.CapacityTaskId = aws.String(_outpostsCapacityTaskId)
	}
	if len(_outpostsOutpostIdentifier) > 0 {
		input.OutpostIdentifier = aws.String(_outpostsOutpostIdentifier)
	}

	if resp, err := client.CancelCapacityTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels the specified order for an Outpost.
func outposts_CancelOrder(cfg aws.Config, client *outposts.Client) {
	input := &outposts.CancelOrderInput{
		// OrderId: *string, // Required
	}

	if len(_outpostsOrderId) > 0 {
		input.OrderId = aws.String(_outpostsOrderId)
	}

	if resp, err := client.CancelOrder(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an order for an Outpost.
func outposts_CreateOrder(cfg aws.Config, client *outposts.Client) {
	input := &outposts.CreateOrderInput{
		// OutpostIdentifier: *string, // Required
		// PaymentOption: types.PaymentOption, // Required
	}

	if len(_outpostsOutpostIdentifier) > 0 {
		input.OutpostIdentifier = aws.String(_outpostsOutpostIdentifier)
	}
	if len(_outpostsPaymentOption) > 0 {
		if err := assignInputField(input, "PaymentOption", _outpostsPaymentOption); err != nil {
			log.Errorf("invalid --payment-option: %s", err.Error())
			return
		}
	}
	if len(_outpostsLineItems) > 0 {
		if err := assignInputField(input, "LineItems", _outpostsLineItems); err != nil {
			log.Errorf("invalid --line-items: %s", err.Error())
			return
		}
	}
	if len(_outpostsPaymentTerm) > 0 {
		if err := assignInputField(input, "PaymentTerm", _outpostsPaymentTerm); err != nil {
			log.Errorf("invalid --payment-term: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateOrder(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Outpost.
// You can specify either an Availability one or an AZ ID.
func outposts_CreateOutpost(cfg aws.Config, client *outposts.Client) {
	input := &outposts.CreateOutpostInput{
		// Name: *string, // Required
		// SiteId: *string, // Required
	}

	if len(_outpostsName) > 0 {
		input.Name = aws.String(_outpostsName)
	}
	if len(_outpostsSiteId) > 0 {
		input.SiteId = aws.String(_outpostsSiteId)
	}
	if len(_outpostsAvailabilityZone) > 0 {
		input.AvailabilityZone = aws.String(_outpostsAvailabilityZone)
	}
	if len(_outpostsAvailabilityZoneId) > 0 {
		input.AvailabilityZoneId = aws.String(_outpostsAvailabilityZoneId)
	}
	if len(_outpostsDescription) > 0 {
		input.Description = aws.String(_outpostsDescription)
	}
	if len(_outpostsSupportedHardwareType) > 0 {
		if err := assignInputField(input, "SupportedHardwareType", _outpostsSupportedHardwareType); err != nil {
			log.Errorf("invalid --supported-hardware-type: %s", err.Error())
			return
		}
	}
	if len(_outpostsTags) > 0 {
		if err := assignInputField(input, "Tags", _outpostsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateOutpost(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a site for an Outpost.
func outposts_CreateSite(cfg aws.Config, client *outposts.Client) {
	input := &outposts.CreateSiteInput{
		// Name: *string, // Required
	}

	if len(_outpostsName) > 0 {
		input.Name = aws.String(_outpostsName)
	}
	if len(_outpostsDescription) > 0 {
		input.Description = aws.String(_outpostsDescription)
	}
	if len(_outpostsNotes) > 0 {
		input.Notes = aws.String(_outpostsNotes)
	}
	if len(_outpostsOperatingAddress) > 0 {
		if err := assignInputField(input, "OperatingAddress", _outpostsOperatingAddress); err != nil {
			log.Errorf("invalid --operating-address: %s", err.Error())
			return
		}
	}
	if len(_outpostsRackPhysicalProperties) > 0 {
		if err := assignInputField(input, "RackPhysicalProperties", _outpostsRackPhysicalProperties); err != nil {
			log.Errorf("invalid --rack-physical-properties: %s", err.Error())
			return
		}
	}
	if len(_outpostsShippingAddress) > 0 {
		if err := assignInputField(input, "ShippingAddress", _outpostsShippingAddress); err != nil {
			log.Errorf("invalid --shipping-address: %s", err.Error())
			return
		}
	}
	if len(_outpostsTags) > 0 {
		if err := assignInputField(input, "Tags", _outpostsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSite(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified Outpost.
func outposts_DeleteOutpost(cfg aws.Config, client *outposts.Client) {
	input := &outposts.DeleteOutpostInput{
		// OutpostId: *string, // Required
	}

	if len(_outpostsOutpostId) > 0 {
		input.OutpostId = aws.String(_outpostsOutpostId)
	}

	if resp, err := client.DeleteOutpost(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified site.
func outposts_DeleteSite(cfg aws.Config, client *outposts.Client) {
	input := &outposts.DeleteSiteInput{
		// SiteId: *string, // Required
	}

	if len(_outpostsSiteId) > 0 {
		input.SiteId = aws.String(_outpostsSiteId)
	}

	if resp, err := client.DeleteSite(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets details of the specified capacity task.
func outposts_GetCapacityTask(cfg aws.Config, client *outposts.Client) {
	input := &outposts.GetCapacityTaskInput{
		// CapacityTaskId: *string, // Required
		// OutpostIdentifier: *string, // Required
	}

	if len(_outpostsCapacityTaskId) > 0 {
		input.CapacityTaskId = aws.String(_outpostsCapacityTaskId)
	}
	if len(_outpostsOutpostIdentifier) > 0 {
		input.OutpostIdentifier = aws.String(_outpostsOutpostIdentifier)
	}

	if resp, err := client.GetCapacityTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the specified catalog item.
func outposts_GetCatalogItem(cfg aws.Config, client *outposts.Client) {
	input := &outposts.GetCatalogItemInput{
		// CatalogItemId: *string, // Required
	}

	if len(_outpostsCatalogItemId) > 0 {
		input.CatalogItemId = aws.String(_outpostsCatalogItemId)
	}

	if resp, err := client.GetCatalogItem(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Amazon Web Services uses this action to install Outpost servers.
// Gets information about the specified connection.
//
// Use CloudTrail to monitor this action or Amazon Web Services managed policy for
// Amazon Web Services Outposts to secure it. For more information, see [Amazon Web Services managed policies for Amazon Web Services Outposts]and [Logging Amazon Web Services Outposts API calls with Amazon Web Services CloudTrail] in
// the Amazon Web Services Outposts User Guide.
//
// [Logging Amazon Web Services Outposts API calls with Amazon Web Services CloudTrail]: https://docs.aws.amazon.com/outposts/latest/userguide/logging-using-cloudtrail.html
// [Amazon Web Services managed policies for Amazon Web Services Outposts]: https://docs.aws.amazon.com/outposts/latest/userguide/security-iam-awsmanpol.html
func outposts_GetConnection(cfg aws.Config, client *outposts.Client) {
	input := &outposts.GetConnectionInput{
		// ConnectionId: *string, // Required
	}

	if len(_outpostsConnectionId) > 0 {
		input.ConnectionId = aws.String(_outpostsConnectionId)
	}

	if resp, err := client.GetConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the specified order.
func outposts_GetOrder(cfg aws.Config, client *outposts.Client) {
	input := &outposts.GetOrderInput{
		// OrderId: *string, // Required
	}

	if len(_outpostsOrderId) > 0 {
		input.OrderId = aws.String(_outpostsOrderId)
	}

	if resp, err := client.GetOrder(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the specified Outpost.
func outposts_GetOutpost(cfg aws.Config, client *outposts.Client) {
	input := &outposts.GetOutpostInput{
		// OutpostId: *string, // Required
	}

	if len(_outpostsOutpostId) > 0 {
		input.OutpostId = aws.String(_outpostsOutpostId)
	}

	if resp, err := client.GetOutpost(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets current and historical billing information about the specified Outpost.
func outposts_GetOutpostBillingInformation(cfg aws.Config, client *outposts.Client) {
	input := &outposts.GetOutpostBillingInformationInput{
		// OutpostIdentifier: *string, // Required
	}

	if len(_outpostsOutpostIdentifier) > 0 {
		input.OutpostIdentifier = aws.String(_outpostsOutpostIdentifier)
	}
	if len(_outpostsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _outpostsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_outpostsNextToken) > 0 {
		input.NextToken = aws.String(_outpostsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetOutpostBillingInformation(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*outposts.GetOutpostBillingInformationOutput
	p := outposts.NewGetOutpostBillingInformationPaginator(client, input)
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

// Gets the instance types for the specified Outpost.
func outposts_GetOutpostInstanceTypes(cfg aws.Config, client *outposts.Client) {
	input := &outposts.GetOutpostInstanceTypesInput{
		// OutpostId: *string, // Required
	}

	if len(_outpostsOutpostId) > 0 {
		input.OutpostId = aws.String(_outpostsOutpostId)
	}
	if len(_outpostsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _outpostsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_outpostsNextToken) > 0 {
		input.NextToken = aws.String(_outpostsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetOutpostInstanceTypes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*outposts.GetOutpostInstanceTypesOutput
	p := outposts.NewGetOutpostInstanceTypesPaginator(client, input)
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

// Gets the instance types that an Outpost can support in InstanceTypeCapacity .
// This will generally include instance types that are not currently configured and
// therefore cannot be launched with the current Outpost capacity configuration.
func outposts_GetOutpostSupportedInstanceTypes(cfg aws.Config, client *outposts.Client) {
	input := &outposts.GetOutpostSupportedInstanceTypesInput{
		// OutpostIdentifier: *string, // Required
	}

	if len(_outpostsOutpostIdentifier) > 0 {
		input.OutpostIdentifier = aws.String(_outpostsOutpostIdentifier)
	}
	if len(_outpostsAssetId) > 0 {
		input.AssetId = aws.String(_outpostsAssetId)
	}
	if len(_outpostsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _outpostsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_outpostsNextToken) > 0 {
		input.NextToken = aws.String(_outpostsNextToken)
	}
	if len(_outpostsOrderId) > 0 {
		input.OrderId = aws.String(_outpostsOrderId)
	}

	if disablePaginator() {
		if resp, err := client.GetOutpostSupportedInstanceTypes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*outposts.GetOutpostSupportedInstanceTypesOutput
	p := outposts.NewGetOutpostSupportedInstanceTypesPaginator(client, input)
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

// Gets information about the specified Outpost site.
func outposts_GetSite(cfg aws.Config, client *outposts.Client) {
	input := &outposts.GetSiteInput{
		// SiteId: *string, // Required
	}

	if len(_outpostsSiteId) > 0 {
		input.SiteId = aws.String(_outpostsSiteId)
	}

	if resp, err := client.GetSite(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the site address of the specified site.
func outposts_GetSiteAddress(cfg aws.Config, client *outposts.Client) {
	input := &outposts.GetSiteAddressInput{
		// AddressType: types.AddressType, // Required
		// SiteId: *string, // Required
	}

	if len(_outpostsAddressType) > 0 {
		if err := assignInputField(input, "AddressType", _outpostsAddressType); err != nil {
			log.Errorf("invalid --address-type: %s", err.Error())
			return
		}
	}
	if len(_outpostsSiteId) > 0 {
		input.SiteId = aws.String(_outpostsSiteId)
	}

	if resp, err := client.GetSiteAddress(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// A list of Amazon EC2 instances, belonging to all accounts, running on the
// specified Outpost. Does not include Amazon EBS or Amazon S3 instances.
func outposts_ListAssetInstances(cfg aws.Config, client *outposts.Client) {
	input := &outposts.ListAssetInstancesInput{
		// OutpostIdentifier: *string, // Required
	}

	if len(_outpostsOutpostIdentifier) > 0 {
		input.OutpostIdentifier = aws.String(_outpostsOutpostIdentifier)
	}
	if len(_outpostsAccountIdFilter) > 0 {
		input.AccountIdFilter = append([]string(nil), _outpostsAccountIdFilter...)
	}
	if len(_outpostsAssetIdFilter) > 0 {
		input.AssetIdFilter = append([]string(nil), _outpostsAssetIdFilter...)
	}
	if len(_outpostsAwsServiceFilter) > 0 {
		if err := assignInputField(input, "AwsServiceFilter", _outpostsAwsServiceFilter); err != nil {
			log.Errorf("invalid --aws-service-filter: %s", err.Error())
			return
		}
	}
	if len(_outpostsInstanceTypeFilter) > 0 {
		input.InstanceTypeFilter = append([]string(nil), _outpostsInstanceTypeFilter...)
	}
	if len(_outpostsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _outpostsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_outpostsNextToken) > 0 {
		input.NextToken = aws.String(_outpostsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAssetInstances(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*outposts.ListAssetInstancesOutput
	p := outposts.NewListAssetInstancesPaginator(client, input)
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

// Lists the hardware assets for the specified Outpost.
// Use filters to return specific results. If you specify multiple filters, the
// results include only the resources that match all of the specified filters. For
// a filter where you can specify multiple values, the results include items that
// match any of the values that you specify for the filter.
func outposts_ListAssets(cfg aws.Config, client *outposts.Client) {
	input := &outposts.ListAssetsInput{
		// OutpostIdentifier: *string, // Required
	}

	if len(_outpostsOutpostIdentifier) > 0 {
		input.OutpostIdentifier = aws.String(_outpostsOutpostIdentifier)
	}
	if len(_outpostsHostIdFilter) > 0 {
		input.HostIdFilter = append([]string(nil), _outpostsHostIdFilter...)
	}
	if len(_outpostsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _outpostsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_outpostsNextToken) > 0 {
		input.NextToken = aws.String(_outpostsNextToken)
	}
	if len(_outpostsStatusFilter) > 0 {
		if err := assignInputField(input, "StatusFilter", _outpostsStatusFilter); err != nil {
			log.Errorf("invalid --status-filter: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListAssets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*outposts.ListAssetsOutput
	p := outposts.NewListAssetsPaginator(client, input)
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

// A list of Amazon EC2 instances running on the Outpost and belonging to the
// account that initiated the capacity task. Use this list to specify the instances
// you cannot stop to free up capacity to run the capacity task.
func outposts_ListBlockingInstancesForCapacityTask(cfg aws.Config, client *outposts.Client) {
	input := &outposts.ListBlockingInstancesForCapacityTaskInput{
		// CapacityTaskId: *string, // Required
		// OutpostIdentifier: *string, // Required
	}

	if len(_outpostsCapacityTaskId) > 0 {
		input.CapacityTaskId = aws.String(_outpostsCapacityTaskId)
	}
	if len(_outpostsOutpostIdentifier) > 0 {
		input.OutpostIdentifier = aws.String(_outpostsOutpostIdentifier)
	}
	if len(_outpostsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _outpostsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_outpostsNextToken) > 0 {
		input.NextToken = aws.String(_outpostsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListBlockingInstancesForCapacityTask(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*outposts.ListBlockingInstancesForCapacityTaskOutput
	p := outposts.NewListBlockingInstancesForCapacityTaskPaginator(client, input)
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

// Lists the capacity tasks for your Amazon Web Services account.
// Use filters to return specific results. If you specify multiple filters, the
// results include only the resources that match all of the specified filters. For
// a filter where you can specify multiple values, the results include items that
// match any of the values that you specify for the filter.
func outposts_ListCapacityTasks(cfg aws.Config, client *outposts.Client) {
	input := &outposts.ListCapacityTasksInput{}

	if len(_outpostsCapacityTaskStatusFilter) > 0 {
		if err := assignInputField(input, "CapacityTaskStatusFilter", _outpostsCapacityTaskStatusFilter); err != nil {
			log.Errorf("invalid --capacity-task-status-filter: %s", err.Error())
			return
		}
	}
	if len(_outpostsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _outpostsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_outpostsNextToken) > 0 {
		input.NextToken = aws.String(_outpostsNextToken)
	}
	if len(_outpostsOutpostIdentifierFilter) > 0 {
		input.OutpostIdentifierFilter = aws.String(_outpostsOutpostIdentifierFilter)
	}

	if disablePaginator() {
		if resp, err := client.ListCapacityTasks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*outposts.ListCapacityTasksOutput
	p := outposts.NewListCapacityTasksPaginator(client, input)
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

// Lists the items in the catalog.
// Use filters to return specific results. If you specify multiple filters, the
// results include only the resources that match all of the specified filters. For
// a filter where you can specify multiple values, the results include items that
// match any of the values that you specify for the filter.
func outposts_ListCatalogItems(cfg aws.Config, client *outposts.Client) {
	input := &outposts.ListCatalogItemsInput{}

	if len(_outpostsEC2FamilyFilter) > 0 {
		input.EC2FamilyFilter = append([]string(nil), _outpostsEC2FamilyFilter...)
	}
	if len(_outpostsItemClassFilter) > 0 {
		if err := assignInputField(input, "ItemClassFilter", _outpostsItemClassFilter); err != nil {
			log.Errorf("invalid --item-class-filter: %s", err.Error())
			return
		}
	}
	if len(_outpostsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _outpostsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_outpostsNextToken) > 0 {
		input.NextToken = aws.String(_outpostsNextToken)
	}
	if len(_outpostsSupportedStorageFilter) > 0 {
		if err := assignInputField(input, "SupportedStorageFilter", _outpostsSupportedStorageFilter); err != nil {
			log.Errorf("invalid --supported-storage-filter: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListCatalogItems(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*outposts.ListCatalogItemsOutput
	p := outposts.NewListCatalogItemsPaginator(client, input)
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

// Lists the Outpost orders for your Amazon Web Services account.
func outposts_ListOrders(cfg aws.Config, client *outposts.Client) {
	input := &outposts.ListOrdersInput{}

	if len(_outpostsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _outpostsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_outpostsNextToken) > 0 {
		input.NextToken = aws.String(_outpostsNextToken)
	}
	if len(_outpostsOutpostIdentifierFilter) > 0 {
		input.OutpostIdentifierFilter = aws.String(_outpostsOutpostIdentifierFilter)
	}

	if disablePaginator() {
		if resp, err := client.ListOrders(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*outposts.ListOrdersOutput
	p := outposts.NewListOrdersPaginator(client, input)
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

// Lists the Outposts for your Amazon Web Services account.
// Use filters to return specific results. If you specify multiple filters, the
// results include only the resources that match all of the specified filters. For
// a filter where you can specify multiple values, the results include items that
// match any of the values that you specify for the filter.
func outposts_ListOutposts(cfg aws.Config, client *outposts.Client) {
	input := &outposts.ListOutpostsInput{}

	if len(_outpostsAvailabilityZoneFilter) > 0 {
		input.AvailabilityZoneFilter = append([]string(nil), _outpostsAvailabilityZoneFilter...)
	}
	if len(_outpostsAvailabilityZoneIdFilter) > 0 {
		input.AvailabilityZoneIdFilter = append([]string(nil), _outpostsAvailabilityZoneIdFilter...)
	}
	if len(_outpostsLifeCycleStatusFilter) > 0 {
		input.LifeCycleStatusFilter = append([]string(nil), _outpostsLifeCycleStatusFilter...)
	}
	if len(_outpostsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _outpostsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_outpostsNextToken) > 0 {
		input.NextToken = aws.String(_outpostsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListOutposts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*outposts.ListOutpostsOutput
	p := outposts.NewListOutpostsPaginator(client, input)
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

// Lists the Outpost sites for your Amazon Web Services account. Use filters to
// return specific results.
//
// Use filters to return specific results. If you specify multiple filters, the
// results include only the resources that match all of the specified filters. For
// a filter where you can specify multiple values, the results include items that
// match any of the values that you specify for the filter.
func outposts_ListSites(cfg aws.Config, client *outposts.Client) {
	input := &outposts.ListSitesInput{}

	if len(_outpostsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _outpostsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_outpostsNextToken) > 0 {
		input.NextToken = aws.String(_outpostsNextToken)
	}
	if len(_outpostsOperatingAddressCityFilter) > 0 {
		input.OperatingAddressCityFilter = append([]string(nil), _outpostsOperatingAddressCityFilter...)
	}
	if len(_outpostsOperatingAddressCountryCodeFilter) > 0 {
		input.OperatingAddressCountryCodeFilter = append([]string(nil), _outpostsOperatingAddressCountryCodeFilter...)
	}
	if len(_outpostsOperatingAddressStateOrRegionFilter) > 0 {
		input.OperatingAddressStateOrRegionFilter = append([]string(nil), _outpostsOperatingAddressStateOrRegionFilter...)
	}

	if disablePaginator() {
		if resp, err := client.ListSites(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*outposts.ListSitesOutput
	p := outposts.NewListSitesPaginator(client, input)
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

// Lists the tags for the specified resource.
func outposts_ListTagsForResource(cfg aws.Config, client *outposts.Client) {
	input := &outposts.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_outpostsResourceArn) > 0 {
		input.ResourceArn = aws.String(_outpostsResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts the specified capacity task. You can have one active capacity task for
// each order and each Outpost.
func outposts_StartCapacityTask(cfg aws.Config, client *outposts.Client) {
	input := &outposts.StartCapacityTaskInput{
		// InstancePools: []types.InstanceTypeCapacity, // Required
		// OutpostIdentifier: *string, // Required
	}

	if len(_outpostsInstancePools) > 0 {
		if err := assignInputField(input, "InstancePools", _outpostsInstancePools); err != nil {
			log.Errorf("invalid --instance-pools: %s", err.Error())
			return
		}
	}
	if len(_outpostsOutpostIdentifier) > 0 {
		input.OutpostIdentifier = aws.String(_outpostsOutpostIdentifier)
	}
	if len(_outpostsAssetId) > 0 {
		input.AssetId = aws.String(_outpostsAssetId)
	}
	if len(_outpostsDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _outpostsDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}
	if len(_outpostsInstancesToExclude) > 0 {
		if err := assignInputField(input, "InstancesToExclude", _outpostsInstancesToExclude); err != nil {
			log.Errorf("invalid --instances-to-exclude: %s", err.Error())
			return
		}
	}
	if len(_outpostsOrderId) > 0 {
		input.OrderId = aws.String(_outpostsOrderId)
	}
	if len(_outpostsTaskActionOnBlockingInstances) > 0 {
		if err := assignInputField(input, "TaskActionOnBlockingInstances", _outpostsTaskActionOnBlockingInstances); err != nil {
			log.Errorf("invalid --task-action-on-blocking-instances: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartCapacityTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Amazon Web Services uses this action to install Outpost servers.
// Starts the connection required for Outpost server installation.
//
// Use CloudTrail to monitor this action or Amazon Web Services managed policy for
// Amazon Web Services Outposts to secure it. For more information, see [Amazon Web Services managed policies for Amazon Web Services Outposts]and [Logging Amazon Web Services Outposts API calls with Amazon Web Services CloudTrail] in
// the Amazon Web Services Outposts User Guide.
//
// [Logging Amazon Web Services Outposts API calls with Amazon Web Services CloudTrail]: https://docs.aws.amazon.com/outposts/latest/userguide/logging-using-cloudtrail.html
// [Amazon Web Services managed policies for Amazon Web Services Outposts]: https://docs.aws.amazon.com/outposts/latest/userguide/security-iam-awsmanpol.html
func outposts_StartConnection(cfg aws.Config, client *outposts.Client) {
	input := &outposts.StartConnectionInput{
		// AssetId: *string, // Required
		// ClientPublicKey: *string, // Required
		// NetworkInterfaceDeviceIndex: int32, // Required
	}

	if len(_outpostsAssetId) > 0 {
		input.AssetId = aws.String(_outpostsAssetId)
	}
	if len(_outpostsClientPublicKey) > 0 {
		input.ClientPublicKey = aws.String(_outpostsClientPublicKey)
	}
	if len(_outpostsNetworkInterfaceDeviceIndex) > 0 {
		if err := assignInputField(input, "NetworkInterfaceDeviceIndex", _outpostsNetworkInterfaceDeviceIndex); err != nil {
			log.Errorf("invalid --network-interface-device-index: %s", err.Error())
			return
		}
	}
	if len(_outpostsDeviceSerialNumber) > 0 {
		input.DeviceSerialNumber = aws.String(_outpostsDeviceSerialNumber)
	}

	if resp, err := client.StartConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts the decommission process to return the Outposts racks or servers.
func outposts_StartOutpostDecommission(cfg aws.Config, client *outposts.Client) {
	input := &outposts.StartOutpostDecommissionInput{
		// OutpostIdentifier: *string, // Required
	}

	if len(_outpostsOutpostIdentifier) > 0 {
		input.OutpostIdentifier = aws.String(_outpostsOutpostIdentifier)
	}
	if len(_outpostsValidateOnly) > 0 {
		if err := assignInputField(input, "ValidateOnly", _outpostsValidateOnly); err != nil {
			log.Errorf("invalid --validate-only: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartOutpostDecommission(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds tags to the specified resource.
func outposts_TagResource(cfg aws.Config, client *outposts.Client) {
	input := &outposts.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_outpostsResourceArn) > 0 {
		input.ResourceArn = aws.String(_outpostsResourceArn)
	}
	if len(_outpostsTags) > 0 {
		if err := assignInputField(input, "Tags", _outpostsTags); err != nil {
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

// Removes tags from the specified resource.
func outposts_UntagResource(cfg aws.Config, client *outposts.Client) {
	input := &outposts.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_outpostsResourceArn) > 0 {
		input.ResourceArn = aws.String(_outpostsResourceArn)
	}
	if len(_outpostsTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _outpostsTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an Outpost.
func outposts_UpdateOutpost(cfg aws.Config, client *outposts.Client) {
	input := &outposts.UpdateOutpostInput{
		// OutpostId: *string, // Required
	}

	if len(_outpostsOutpostId) > 0 {
		input.OutpostId = aws.String(_outpostsOutpostId)
	}
	if len(_outpostsDescription) > 0 {
		input.Description = aws.String(_outpostsDescription)
	}
	if len(_outpostsName) > 0 {
		input.Name = aws.String(_outpostsName)
	}
	if len(_outpostsSupportedHardwareType) > 0 {
		if err := assignInputField(input, "SupportedHardwareType", _outpostsSupportedHardwareType); err != nil {
			log.Errorf("invalid --supported-hardware-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateOutpost(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified site.
func outposts_UpdateSite(cfg aws.Config, client *outposts.Client) {
	input := &outposts.UpdateSiteInput{
		// SiteId: *string, // Required
	}

	if len(_outpostsSiteId) > 0 {
		input.SiteId = aws.String(_outpostsSiteId)
	}
	if len(_outpostsDescription) > 0 {
		input.Description = aws.String(_outpostsDescription)
	}
	if len(_outpostsName) > 0 {
		input.Name = aws.String(_outpostsName)
	}
	if len(_outpostsNotes) > 0 {
		input.Notes = aws.String(_outpostsNotes)
	}

	if resp, err := client.UpdateSite(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the address of the specified site.
// You can't update a site address if there is an order in progress. You must wait
// for the order to complete or cancel the order.
//
// You can update the operating address before you place an order at the site, or
// after all Outposts that belong to the site have been deactivated.
func outposts_UpdateSiteAddress(cfg aws.Config, client *outposts.Client) {
	input := &outposts.UpdateSiteAddressInput{
		// Address: *types.Address, // Required
		// AddressType: types.AddressType, // Required
		// SiteId: *string, // Required
	}

	if len(_outpostsAddress) > 0 {
		if err := assignInputField(input, "Address", _outpostsAddress); err != nil {
			log.Errorf("invalid --address: %s", err.Error())
			return
		}
	}
	if len(_outpostsAddressType) > 0 {
		if err := assignInputField(input, "AddressType", _outpostsAddressType); err != nil {
			log.Errorf("invalid --address-type: %s", err.Error())
			return
		}
	}
	if len(_outpostsSiteId) > 0 {
		input.SiteId = aws.String(_outpostsSiteId)
	}

	if resp, err := client.UpdateSiteAddress(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update the physical and logistical details for a rack at a site. For more
// information about hardware requirements for racks, see [Network readiness checklist]in the Amazon Web
// Services Outposts User Guide.
//
// To update a rack at a site with an order of IN_PROGRESS , you must wait for the
// order to complete or cancel the order.
//
// [Network readiness checklist]: https://docs.aws.amazon.com/outposts/latest/userguide/outposts-requirements.html#checklist
func outposts_UpdateSiteRackPhysicalProperties(cfg aws.Config, client *outposts.Client) {
	input := &outposts.UpdateSiteRackPhysicalPropertiesInput{
		// SiteId: *string, // Required
	}

	if len(_outpostsSiteId) > 0 {
		input.SiteId = aws.String(_outpostsSiteId)
	}
	if len(_outpostsFiberOpticCableType) > 0 {
		if err := assignInputField(input, "FiberOpticCableType", _outpostsFiberOpticCableType); err != nil {
			log.Errorf("invalid --fiber-optic-cable-type: %s", err.Error())
			return
		}
	}
	if len(_outpostsMaximumSupportedWeightLbs) > 0 {
		if err := assignInputField(input, "MaximumSupportedWeightLbs", _outpostsMaximumSupportedWeightLbs); err != nil {
			log.Errorf("invalid --maximum-supported-weight-lbs: %s", err.Error())
			return
		}
	}
	if len(_outpostsOpticalStandard) > 0 {
		if err := assignInputField(input, "OpticalStandard", _outpostsOpticalStandard); err != nil {
			log.Errorf("invalid --optical-standard: %s", err.Error())
			return
		}
	}
	if len(_outpostsPowerConnector) > 0 {
		if err := assignInputField(input, "PowerConnector", _outpostsPowerConnector); err != nil {
			log.Errorf("invalid --power-connector: %s", err.Error())
			return
		}
	}
	if len(_outpostsPowerDrawKva) > 0 {
		if err := assignInputField(input, "PowerDrawKva", _outpostsPowerDrawKva); err != nil {
			log.Errorf("invalid --power-draw-kva: %s", err.Error())
			return
		}
	}
	if len(_outpostsPowerFeedDrop) > 0 {
		if err := assignInputField(input, "PowerFeedDrop", _outpostsPowerFeedDrop); err != nil {
			log.Errorf("invalid --power-feed-drop: %s", err.Error())
			return
		}
	}
	if len(_outpostsPowerPhase) > 0 {
		if err := assignInputField(input, "PowerPhase", _outpostsPowerPhase); err != nil {
			log.Errorf("invalid --power-phase: %s", err.Error())
			return
		}
	}
	if len(_outpostsUplinkCount) > 0 {
		if err := assignInputField(input, "UplinkCount", _outpostsUplinkCount); err != nil {
			log.Errorf("invalid --uplink-count: %s", err.Error())
			return
		}
	}
	if len(_outpostsUplinkGbps) > 0 {
		if err := assignInputField(input, "UplinkGbps", _outpostsUplinkGbps); err != nil {
			log.Errorf("invalid --uplink-gbps: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSiteRackPhysicalProperties(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_outpostsCmd)
	_outpostsCmd.Flags().SortFlags = false

	_outpostsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_outpostsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_outpostsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_outpostsCmd.Flags().StringSliceVarP(&_outpostsAccountIdFilter, "account-id-filter", "", nil, "Account ID Filter")
	_outpostsCmd.Flags().StringVarP(&_outpostsAddress, "address", "", "", "Address")
	_outpostsCmd.Flags().StringVarP(&_outpostsAddressType, "address-type", "", "", "Address Type")
	_outpostsCmd.Flags().StringVarP(&_outpostsAssetId, "asset-id", "", "", "Asset ID")
	_outpostsCmd.Flags().StringSliceVarP(&_outpostsAssetIdFilter, "asset-id-filter", "", nil, "Asset ID Filter")
	_outpostsCmd.Flags().StringVarP(&_outpostsAvailabilityZone, "availability-zone", "", "", "Availability Zone")
	_outpostsCmd.Flags().StringSliceVarP(&_outpostsAvailabilityZoneFilter, "availability-zone-filter", "", nil, "Availability Zone Filter")
	_outpostsCmd.Flags().StringVarP(&_outpostsAvailabilityZoneId, "availability-zone-id", "", "", "Availability Zone ID")
	_outpostsCmd.Flags().StringSliceVarP(&_outpostsAvailabilityZoneIdFilter, "availability-zone-id-filter", "", nil, "Availability Zone ID Filter")
	_outpostsCmd.Flags().StringVarP(&_outpostsAwsServiceFilter, "aws-service-filter", "", "", "AWS Service Filter")
	_outpostsCmd.Flags().StringVarP(&_outpostsCapacityTaskId, "capacity-task-id", "", "", "Capacity Task ID")
	_outpostsCmd.Flags().StringVarP(&_outpostsCapacityTaskStatusFilter, "capacity-task-status-filter", "", "", "Capacity Task Status Filter")
	_outpostsCmd.Flags().StringVarP(&_outpostsCatalogItemId, "catalog-item-id", "", "", "Catalog Item ID")
	_outpostsCmd.Flags().StringVarP(&_outpostsClientPublicKey, "client-public-key", "", "", "Client Public Key")
	_outpostsCmd.Flags().StringVarP(&_outpostsConnectionId, "connection-id", "", "", "Connection ID")
	_outpostsCmd.Flags().StringVarP(&_outpostsDescription, "description", "", "", "Description")
	_outpostsCmd.Flags().StringVarP(&_outpostsDeviceSerialNumber, "device-serial-number", "", "", "Device Serial Number")
	_outpostsCmd.Flags().StringVarP(&_outpostsDryRun, "dry-run", "", "", "Dry Run")
	_outpostsCmd.Flags().StringSliceVarP(&_outpostsEC2FamilyFilter, "ec2-family-filter", "", nil, "EC2 Family Filter")
	_outpostsCmd.Flags().StringVarP(&_outpostsFiberOpticCableType, "fiber-optic-cable-type", "", "", "Fiber Optic Cable Type")
	_outpostsCmd.Flags().StringSliceVarP(&_outpostsHostIdFilter, "host-id-filter", "", nil, "Host ID Filter")
	_outpostsCmd.Flags().StringVarP(&_outpostsInstancePools, "instance-pools", "", "", "Instance Pools")
	_outpostsCmd.Flags().StringSliceVarP(&_outpostsInstanceTypeFilter, "instance-type-filter", "", nil, "Instance Type Filter")
	_outpostsCmd.Flags().StringVarP(&_outpostsInstancesToExclude, "instances-to-exclude", "", "", "Instances To Exclude")
	_outpostsCmd.Flags().StringVarP(&_outpostsItemClassFilter, "item-class-filter", "", "", "Item Class Filter")
	_outpostsCmd.Flags().StringSliceVarP(&_outpostsLifeCycleStatusFilter, "life-cycle-status-filter", "", nil, "Life Cycle Status Filter")
	_outpostsCmd.Flags().StringVarP(&_outpostsLineItems, "line-items", "", "", "Line Items")
	_outpostsCmd.Flags().StringVarP(&_outpostsMaxResults, "max-results", "", "", "Max Results")
	_outpostsCmd.Flags().StringVarP(&_outpostsMaximumSupportedWeightLbs, "maximum-supported-weight-lbs", "", "", "Maximum Supported Weight Lbs")
	_outpostsCmd.Flags().StringVarP(&_outpostsName, "name", "", "", "Name")
	_outpostsCmd.Flags().StringVarP(&_outpostsNetworkInterfaceDeviceIndex, "network-interface-device-index", "", "", "Network Interface Device Index")
	_outpostsCmd.Flags().StringVarP(&_outpostsNextToken, "next-token", "", "", "Next Token")
	_outpostsCmd.Flags().StringVarP(&_outpostsNotes, "notes", "", "", "Notes")
	_outpostsCmd.Flags().StringVarP(&_outpostsOperatingAddress, "operating-address", "", "", "Operating Address")
	_outpostsCmd.Flags().StringSliceVarP(&_outpostsOperatingAddressCityFilter, "operating-address-city-filter", "", nil, "Operating Address City Filter")
	_outpostsCmd.Flags().StringSliceVarP(&_outpostsOperatingAddressCountryCodeFilter, "operating-address-country-code-filter", "", nil, "Operating Address Country Code Filter")
	_outpostsCmd.Flags().StringSliceVarP(&_outpostsOperatingAddressStateOrRegionFilter, "operating-address-state-or-region-filter", "", nil, "Operating Address State Or Region Filter")
	_outpostsCmd.Flags().StringVarP(&_outpostsOpticalStandard, "optical-standard", "", "", "Optical Standard")
	_outpostsCmd.Flags().StringVarP(&_outpostsOrderId, "order-id", "", "", "Order ID")
	_outpostsCmd.Flags().StringVarP(&_outpostsOutpostId, "outpost-id", "", "", "Outpost ID")
	_outpostsCmd.Flags().StringVarP(&_outpostsOutpostIdentifier, "outpost-identifier", "", "", "Outpost Identifier")
	_outpostsCmd.Flags().StringVarP(&_outpostsOutpostIdentifierFilter, "outpost-identifier-filter", "", "", "Outpost Identifier Filter")
	_outpostsCmd.Flags().StringVarP(&_outpostsPaymentOption, "payment-option", "", "", "Payment Option")
	_outpostsCmd.Flags().StringVarP(&_outpostsPaymentTerm, "payment-term", "", "", "Payment Term")
	_outpostsCmd.Flags().StringVarP(&_outpostsPowerConnector, "power-connector", "", "", "Power Connector")
	_outpostsCmd.Flags().StringVarP(&_outpostsPowerDrawKva, "power-draw-kva", "", "", "Power Draw Kva")
	_outpostsCmd.Flags().StringVarP(&_outpostsPowerFeedDrop, "power-feed-drop", "", "", "Power Feed Drop")
	_outpostsCmd.Flags().StringVarP(&_outpostsPowerPhase, "power-phase", "", "", "Power Phase")
	_outpostsCmd.Flags().StringVarP(&_outpostsRackPhysicalProperties, "rack-physical-properties", "", "", "Rack Physical Properties")
	_outpostsCmd.Flags().StringVarP(&_outpostsResourceArn, "resource-arn", "", "", "Resource ARN")
	_outpostsCmd.Flags().StringVarP(&_outpostsShippingAddress, "shipping-address", "", "", "Shipping Address")
	_outpostsCmd.Flags().StringVarP(&_outpostsSiteId, "site-id", "", "", "Site ID")
	_outpostsCmd.Flags().StringVarP(&_outpostsStatusFilter, "status-filter", "", "", "Status Filter")
	_outpostsCmd.Flags().StringVarP(&_outpostsSupportedHardwareType, "supported-hardware-type", "", "", "Supported Hardware Type")
	_outpostsCmd.Flags().StringVarP(&_outpostsSupportedStorageFilter, "supported-storage-filter", "", "", "Supported Storage Filter")
	_outpostsCmd.Flags().StringSliceVarP(&_outpostsTagKeys, "tag-keys", "", nil, "Tag Keys")
	_outpostsCmd.Flags().StringVarP(&_outpostsTags, "tags", "", "", "Tags")
	_outpostsCmd.Flags().StringVarP(&_outpostsTaskActionOnBlockingInstances, "task-action-on-blocking-instances", "", "", "Task Action On Blocking Instances")
	_outpostsCmd.Flags().StringVarP(&_outpostsUplinkCount, "uplink-count", "", "", "Uplink Count")
	_outpostsCmd.Flags().StringVarP(&_outpostsUplinkGbps, "uplink-gbps", "", "", "Uplink Gbps")
	_outpostsCmd.Flags().StringVarP(&_outpostsValidateOnly, "validate-only", "", "", "Validate Only")

	_outpostsCmd.Flags().BoolVarP(&_outpostsCancelCapacityTask, "cancel-capacity-task", "", false, "Cancel Capacity Task")
	_outpostsCmd.Flags().BoolVarP(&_outpostsCancelOrder, "cancel-order", "", false, "Cancel Order")
	_outpostsCmd.Flags().BoolVarP(&_outpostsCreateOrder, "create-order", "", false, "Create Order")
	_outpostsCmd.Flags().BoolVarP(&_outpostsCreateOutpost, "create-outpost", "", false, "Create Outpost")
	_outpostsCmd.Flags().BoolVarP(&_outpostsCreateSite, "create-site", "", false, "Create Site")
	_outpostsCmd.Flags().BoolVarP(&_outpostsDeleteOutpost, "delete-outpost", "", false, "Delete Outpost")
	_outpostsCmd.Flags().BoolVarP(&_outpostsDeleteSite, "delete-site", "", false, "Delete Site")
	_outpostsCmd.Flags().BoolVarP(&_outpostsGetCapacityTask, "get-capacity-task", "", false, "Get Capacity Task")
	_outpostsCmd.Flags().BoolVarP(&_outpostsGetCatalogItem, "get-catalog-item", "", false, "Get Catalog Item")
	_outpostsCmd.Flags().BoolVarP(&_outpostsGetConnection, "get-connection", "", false, "Get Connection")
	_outpostsCmd.Flags().BoolVarP(&_outpostsGetOrder, "get-order", "", false, "Get Order")
	_outpostsCmd.Flags().BoolVarP(&_outpostsGetOutpost, "get-outpost", "", false, "Get Outpost")
	_outpostsCmd.Flags().BoolVarP(&_outpostsGetOutpostBillingInformation, "get-outpost-billing-information", "", false, "Get Outpost Billing Information")
	_outpostsCmd.Flags().BoolVarP(&_outpostsGetOutpostInstanceTypes, "get-outpost-instance-types", "", false, "Get Outpost Instance Types")
	_outpostsCmd.Flags().BoolVarP(&_outpostsGetOutpostSupportedInstanceTypes, "get-outpost-supported-instance-types", "", false, "Get Outpost Supported Instance Types")
	_outpostsCmd.Flags().BoolVarP(&_outpostsGetSite, "get-site", "", false, "Get Site")
	_outpostsCmd.Flags().BoolVarP(&_outpostsGetSiteAddress, "get-site-address", "", false, "Get Site Address")
	_outpostsCmd.Flags().BoolVarP(&_outpostsListAssetInstances, "list-asset-instances", "", false, "List Asset Instances")
	_outpostsCmd.Flags().BoolVarP(&_outpostsListAssets, "list-assets", "", false, "List Assets")
	_outpostsCmd.Flags().BoolVarP(&_outpostsListBlockingInstancesForCapacityTask, "list-blocking-instances-for-capacity-task", "", false, "List Blocking Instances For Capacity Task")
	_outpostsCmd.Flags().BoolVarP(&_outpostsListCapacityTasks, "list-capacity-tasks", "", false, "List Capacity Tasks")
	_outpostsCmd.Flags().BoolVarP(&_outpostsListCatalogItems, "list-catalog-items", "", false, "List Catalog Items")
	_outpostsCmd.Flags().BoolVarP(&_outpostsListOrders, "list-orders", "", false, "List Orders")
	_outpostsCmd.Flags().BoolVarP(&_outpostsListOutposts, "list-outposts", "", false, "List Outposts")
	_outpostsCmd.Flags().BoolVarP(&_outpostsListSites, "list-sites", "", false, "List Sites")
	_outpostsCmd.Flags().BoolVarP(&_outpostsListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_outpostsCmd.Flags().BoolVarP(&_outpostsStartCapacityTask, "start-capacity-task", "", false, "Start Capacity Task")
	_outpostsCmd.Flags().BoolVarP(&_outpostsStartConnection, "start-connection", "", false, "Start Connection")
	_outpostsCmd.Flags().BoolVarP(&_outpostsStartOutpostDecommission, "start-outpost-decommission", "", false, "Start Outpost Decommission")
	_outpostsCmd.Flags().BoolVarP(&_outpostsTagResource, "tag-resource", "", false, "Tag Resource")
	_outpostsCmd.Flags().BoolVarP(&_outpostsUntagResource, "untag-resource", "", false, "Untag Resource")
	_outpostsCmd.Flags().BoolVarP(&_outpostsUpdateOutpost, "update-outpost", "", false, "Update Outpost")
	_outpostsCmd.Flags().BoolVarP(&_outpostsUpdateSite, "update-site", "", false, "Update Site")
	_outpostsCmd.Flags().BoolVarP(&_outpostsUpdateSiteAddress, "update-site-address", "", false, "Update Site Address")
	_outpostsCmd.Flags().BoolVarP(&_outpostsUpdateSiteRackPhysicalProperties, "update-site-rack-physical-properties", "", false, "Update Site Rack Physical Properties")

}
