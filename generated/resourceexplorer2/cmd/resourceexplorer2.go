package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/resourceexplorer2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// resourceexplorer2Cmd represents the resourceexplorer2 command
var _resourceexplorer2Cmd = &cobra.Command{
	Use:   "resourceexplorer2",
	Short: "AWS resourceexplorer2 CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := resourceexplorer2.NewFromConfig(cfg)
		if _resourceexplorer2AssociateDefaultView {
			resourceexplorer2_AssociateDefaultView(cfg, client)
			return
		}
		if _resourceexplorer2BatchGetView {
			resourceexplorer2_BatchGetView(cfg, client)
			return
		}
		if _resourceexplorer2CreateIndex {
			resourceexplorer2_CreateIndex(cfg, client)
			return
		}
		if _resourceexplorer2CreateResourceExplorerSetup {
			resourceexplorer2_CreateResourceExplorerSetup(cfg, client)
			return
		}
		if _resourceexplorer2CreateView {
			resourceexplorer2_CreateView(cfg, client)
			return
		}
		if _resourceexplorer2DeleteIndex {
			resourceexplorer2_DeleteIndex(cfg, client)
			return
		}
		if _resourceexplorer2DeleteResourceExplorerSetup {
			resourceexplorer2_DeleteResourceExplorerSetup(cfg, client)
			return
		}
		if _resourceexplorer2DeleteView {
			resourceexplorer2_DeleteView(cfg, client)
			return
		}
		if _resourceexplorer2DisassociateDefaultView {
			resourceexplorer2_DisassociateDefaultView(cfg, client)
			return
		}
		if _resourceexplorer2GetAccountLevelServiceConfiguration {
			resourceexplorer2_GetAccountLevelServiceConfiguration(cfg, client)
			return
		}
		if _resourceexplorer2GetDefaultView {
			resourceexplorer2_GetDefaultView(cfg, client)
			return
		}
		if _resourceexplorer2GetIndex {
			resourceexplorer2_GetIndex(cfg, client)
			return
		}
		if _resourceexplorer2GetManagedView {
			resourceexplorer2_GetManagedView(cfg, client)
			return
		}
		if _resourceexplorer2GetResourceExplorerSetup {
			resourceexplorer2_GetResourceExplorerSetup(cfg, client)
			return
		}
		if _resourceexplorer2GetServiceIndex {
			resourceexplorer2_GetServiceIndex(cfg, client)
			return
		}
		if _resourceexplorer2GetServiceView {
			resourceexplorer2_GetServiceView(cfg, client)
			return
		}
		if _resourceexplorer2GetView {
			resourceexplorer2_GetView(cfg, client)
			return
		}
		if _resourceexplorer2ListIndexes {
			resourceexplorer2_ListIndexes(cfg, client)
			return
		}
		if _resourceexplorer2ListIndexesForMembers {
			resourceexplorer2_ListIndexesForMembers(cfg, client)
			return
		}
		if _resourceexplorer2ListManagedViews {
			resourceexplorer2_ListManagedViews(cfg, client)
			return
		}
		if _resourceexplorer2ListResources {
			resourceexplorer2_ListResources(cfg, client)
			return
		}
		if _resourceexplorer2ListServiceIndexes {
			resourceexplorer2_ListServiceIndexes(cfg, client)
			return
		}
		if _resourceexplorer2ListServiceViews {
			resourceexplorer2_ListServiceViews(cfg, client)
			return
		}
		if _resourceexplorer2ListStreamingAccessForServices {
			resourceexplorer2_ListStreamingAccessForServices(cfg, client)
			return
		}
		if _resourceexplorer2ListSupportedResourceTypes {
			resourceexplorer2_ListSupportedResourceTypes(cfg, client)
			return
		}
		if _resourceexplorer2ListTagsForResource {
			resourceexplorer2_ListTagsForResource(cfg, client)
			return
		}
		if _resourceexplorer2ListViews {
			resourceexplorer2_ListViews(cfg, client)
			return
		}
		if _resourceexplorer2Search {
			resourceexplorer2_Search(cfg, client)
			return
		}
		if _resourceexplorer2TagResource {
			resourceexplorer2_TagResource(cfg, client)
			return
		}
		if _resourceexplorer2UntagResource {
			resourceexplorer2_UntagResource(cfg, client)
			return
		}
		if _resourceexplorer2UpdateIndexType {
			resourceexplorer2_UpdateIndexType(cfg, client)
			return
		}
		if _resourceexplorer2UpdateView {
			resourceexplorer2_UpdateView(cfg, client)
			return
		}

	},
}

var (
	_resourceexplorer2AssociateDefaultView                bool
	_resourceexplorer2BatchGetView                        bool
	_resourceexplorer2CreateIndex                         bool
	_resourceexplorer2CreateResourceExplorerSetup         bool
	_resourceexplorer2CreateView                          bool
	_resourceexplorer2DeleteIndex                         bool
	_resourceexplorer2DeleteResourceExplorerSetup         bool
	_resourceexplorer2DeleteView                          bool
	_resourceexplorer2DisassociateDefaultView             bool
	_resourceexplorer2GetAccountLevelServiceConfiguration bool
	_resourceexplorer2GetDefaultView                      bool
	_resourceexplorer2GetIndex                            bool
	_resourceexplorer2GetManagedView                      bool
	_resourceexplorer2GetResourceExplorerSetup            bool
	_resourceexplorer2GetServiceIndex                     bool
	_resourceexplorer2GetServiceView                      bool
	_resourceexplorer2GetView                             bool
	_resourceexplorer2ListIndexes                         bool
	_resourceexplorer2ListIndexesForMembers               bool
	_resourceexplorer2ListManagedViews                    bool
	_resourceexplorer2ListResources                       bool
	_resourceexplorer2ListServiceIndexes                  bool
	_resourceexplorer2ListServiceViews                    bool
	_resourceexplorer2ListStreamingAccessForServices      bool
	_resourceexplorer2ListSupportedResourceTypes          bool
	_resourceexplorer2ListTagsForResource                 bool
	_resourceexplorer2ListViews                           bool
	_resourceexplorer2Search                              bool
	_resourceexplorer2TagResource                         bool
	_resourceexplorer2UntagResource                       bool
	_resourceexplorer2UpdateIndexType                     bool
	_resourceexplorer2UpdateView                          bool

	_resourceexplorer2AccountIdList      []string
	_resourceexplorer2AggregatorRegions  []string
	_resourceexplorer2Arn                string
	_resourceexplorer2ClientToken        string
	_resourceexplorer2DeleteInAllRegions string
	_resourceexplorer2Filters            string
	_resourceexplorer2IncludedProperties string
	_resourceexplorer2ManagedViewArn     string
	_resourceexplorer2MaxResults         string
	_resourceexplorer2NextToken          string
	_resourceexplorer2QueryString        string
	_resourceexplorer2RegionList         []string
	_resourceexplorer2Regions            []string
	_resourceexplorer2ResourceArn        string
	_resourceexplorer2Scope              string
	_resourceexplorer2ServicePrincipal   string
	_resourceexplorer2ServiceViewArn     string
	_resourceexplorer2TagKeys            []string
	_resourceexplorer2Tags               string
	_resourceexplorer2TaskId             string
	_resourceexplorer2Type               string
	_resourceexplorer2ViewArn            string
	_resourceexplorer2ViewArns           []string
	_resourceexplorer2ViewName           string
)

// Sets the specified view as the default for the Amazon Web Services Region in
// which you call this operation. When a user performs a Searchthat doesn't explicitly
// specify which view to use, then Amazon Web Services Resource Explorer
// automatically chooses this default view for searches performed in this Amazon
// Web Services Region.
//
// If an Amazon Web Services Region doesn't have a default view configured, then
// users must explicitly specify a view with every Search operation performed in
// that Region.
func resourceexplorer2_AssociateDefaultView(cfg aws.Config, client *resourceexplorer2.Client) {
	input := &resourceexplorer2.AssociateDefaultViewInput{
		// ViewArn: *string, // Required
	}

	if len(_resourceexplorer2ViewArn) > 0 {
		input.ViewArn = aws.String(_resourceexplorer2ViewArn)
	}

	if resp, err := client.AssociateDefaultView(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details about a list of views.
func resourceexplorer2_BatchGetView(cfg aws.Config, client *resourceexplorer2.Client) {
	input := &resourceexplorer2.BatchGetViewInput{}

	if len(_resourceexplorer2ViewArns) > 0 {
		input.ViewArns = append([]string(nil), _resourceexplorer2ViewArns...)
	}

	if resp, err := client.BatchGetView(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Turns on Amazon Web Services Resource Explorer in the Amazon Web Services
// Region in which you called this operation by creating an index. Resource
// Explorer begins discovering the resources in this Region and stores the details
// about the resources in the index so that they can be queried by using the Search
// operation. You can create only one index in a Region.
//
// This operation creates only a local index. To promote the local index in one
// Amazon Web Services Region into the aggregator index for the Amazon Web Services
// account, use the UpdateIndexTypeoperation. For more information, see [Turning on cross-Region search by creating an aggregator index] in the Amazon Web
// Services Resource Explorer User Guide.
//
// For more details about what happens when you turn on Resource Explorer in an
// Amazon Web Services Region, see [Turn on Resource Explorer to index your resources in an Amazon Web Services Region]in the Amazon Web Services Resource Explorer
// User Guide.
//
// If this is the first Amazon Web Services Region in which you've created an
// index for Resource Explorer, then this operation also [creates a service-linked role]in your Amazon Web
// Services account that allows Resource Explorer to enumerate your resources to
// populate the index.
//
// - Action: resource-explorer-2:CreateIndex
//
// Resource: The ARN of the index (as it will exist after the operation completes)
//
// in the Amazon Web Services Region and account in which you're trying to create
// the index. Use the wildcard character ( * ) at the end of the string to match
// the eventual UUID. For example, the following Resource element restricts the
// role or user to creating an index in only the us-east-2 Region of the
// specified account.
//
// "Resource": "arn:aws:resource-explorer-2:us-west-2:<account-id>:index/*"
//
// Alternatively, you can use "Resource": "*" to allow the role or user to create
//
// an index in any Region.
//
// - Action: iam:CreateServiceLinkedRole
//
// Resource: No specific resource (*).
//
// # This permission is required only the first time you create an index to turn on
//
// Resource Explorer in the account. Resource Explorer uses this to create the [service-linked role needed to index the resources in your account].
// Resource Explorer uses the same service-linked role for all additional indexes
// you create afterwards.
//
// [Turning on cross-Region search by creating an aggregator index]: https://docs.aws.amazon.com/resource-explorer/latest/userguide/manage-aggregator-region.html
// [creates a service-linked role]: https://docs.aws.amazon.com/resource-explorer/latest/userguide/security_iam_service-linked-roles.html
// [Turn on Resource Explorer to index your resources in an Amazon Web Services Region]: https://docs.aws.amazon.com/resource-explorer/latest/userguide/manage-service-activate.html
// [service-linked role needed to index the resources in your account]: https://docs.aws.amazon.com/resource-explorer/latest/userguide/security_iam_service-linked-roles.html
func resourceexplorer2_CreateIndex(cfg aws.Config, client *resourceexplorer2.Client) {
	input := &resourceexplorer2.CreateIndexInput{}

	if len(_resourceexplorer2ClientToken) > 0 {
		input.ClientToken = aws.String(_resourceexplorer2ClientToken)
	}
	if len(_resourceexplorer2Tags) > 0 {
		if err := assignInputField(input, "Tags", _resourceexplorer2Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateIndex(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a Resource Explorer setup configuration across multiple Amazon Web
// Services Regions. This operation sets up indexes and views in the specified
// Regions. This operation can also be used to set an aggregator Region for
// cross-Region resource search.
func resourceexplorer2_CreateResourceExplorerSetup(cfg aws.Config, client *resourceexplorer2.Client) {
	input := &resourceexplorer2.CreateResourceExplorerSetupInput{
		// RegionList: []string, // Required
		// ViewName: *string, // Required
	}

	if len(_resourceexplorer2RegionList) > 0 {
		input.RegionList = append([]string(nil), _resourceexplorer2RegionList...)
	}
	if len(_resourceexplorer2ViewName) > 0 {
		input.ViewName = aws.String(_resourceexplorer2ViewName)
	}
	if len(_resourceexplorer2AggregatorRegions) > 0 {
		input.AggregatorRegions = append([]string(nil), _resourceexplorer2AggregatorRegions...)
	}

	if resp, err := client.CreateResourceExplorerSetup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a view that users can query by using the Search operation. Results from
// queries that you make using this view include only resources that match the
// view's Filters . For more information about Amazon Web Services Resource
// Explorer views, see [Managing views]in the Amazon Web Services Resource Explorer User Guide.
//
// Only the principals with an IAM identity-based policy that grants Allow to the
// Search action on a Resource with the [Amazon resource name (ARN)] of this view can Search using views you create
// with this operation.
//
// [Managing views]: https://docs.aws.amazon.com/resource-explorer/latest/userguide/manage-views.html
// [Amazon resource name (ARN)]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
func resourceexplorer2_CreateView(cfg aws.Config, client *resourceexplorer2.Client) {
	input := &resourceexplorer2.CreateViewInput{
		// ViewName: *string, // Required
	}

	if len(_resourceexplorer2ViewName) > 0 {
		input.ViewName = aws.String(_resourceexplorer2ViewName)
	}
	if len(_resourceexplorer2ClientToken) > 0 {
		input.ClientToken = aws.String(_resourceexplorer2ClientToken)
	}
	if len(_resourceexplorer2Filters) > 0 {
		if err := assignInputField(input, "Filters", _resourceexplorer2Filters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_resourceexplorer2IncludedProperties) > 0 {
		if err := assignInputField(input, "IncludedProperties", _resourceexplorer2IncludedProperties); err != nil {
			log.Errorf("invalid --included-properties: %s", err.Error())
			return
		}
	}
	if len(_resourceexplorer2Scope) > 0 {
		input.Scope = aws.String(_resourceexplorer2Scope)
	}
	if len(_resourceexplorer2Tags) > 0 {
		if err := assignInputField(input, "Tags", _resourceexplorer2Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateView(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified index and turns off Amazon Web Services Resource Explorer
// in the specified Amazon Web Services Region. When you delete an index, Resource
// Explorer stops discovering and indexing resources in that Region. Resource
// Explorer also deletes all views in that Region. These actions occur as
// asynchronous background tasks. You can check to see when the actions are
// complete by using the GetIndexoperation and checking the Status response value.
//
// If the index you delete is the aggregator index for the Amazon Web Services
// account, you must wait 24 hours before you can promote another local index to be
// the aggregator index for the account. Users can't perform account-wide searches
// using Resource Explorer until another aggregator index is configured.
func resourceexplorer2_DeleteIndex(cfg aws.Config, client *resourceexplorer2.Client) {
	input := &resourceexplorer2.DeleteIndexInput{
		// Arn: *string, // Required
	}

	if len(_resourceexplorer2Arn) > 0 {
		input.Arn = aws.String(_resourceexplorer2Arn)
	}

	if resp, err := client.DeleteIndex(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Resource Explorer setup configuration. This operation removes indexes
// and views from the specified Regions or all Regions where Resource Explorer is
// configured.
func resourceexplorer2_DeleteResourceExplorerSetup(cfg aws.Config, client *resourceexplorer2.Client) {
	input := &resourceexplorer2.DeleteResourceExplorerSetupInput{}

	if len(_resourceexplorer2DeleteInAllRegions) > 0 {
		if err := assignInputField(input, "DeleteInAllRegions", _resourceexplorer2DeleteInAllRegions); err != nil {
			log.Errorf("invalid --delete-in-all-regions: %s", err.Error())
			return
		}
	}
	if len(_resourceexplorer2RegionList) > 0 {
		input.RegionList = append([]string(nil), _resourceexplorer2RegionList...)
	}

	if resp, err := client.DeleteResourceExplorerSetup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified view.
// If the specified view is the default view for its Amazon Web Services Region,
// then all Searchoperations in that Region must explicitly specify the view to use
// until you configure a new default by calling the AssociateDefaultViewoperation.
func resourceexplorer2_DeleteView(cfg aws.Config, client *resourceexplorer2.Client) {
	input := &resourceexplorer2.DeleteViewInput{
		// ViewArn: *string, // Required
	}

	if len(_resourceexplorer2ViewArn) > 0 {
		input.ViewArn = aws.String(_resourceexplorer2ViewArn)
	}

	if resp, err := client.DeleteView(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// After you call this operation, the affected Amazon Web Services Region no
// longer has a default view. All Searchoperations in that Region must explicitly
// specify a view or the operation fails. You can configure a new default by
// calling the AssociateDefaultViewoperation.
//
// If an Amazon Web Services Region doesn't have a default view configured, then
// users must explicitly specify a view with every Search operation performed in
// that Region.
func resourceexplorer2_DisassociateDefaultView(cfg aws.Config, client *resourceexplorer2.Client) {
	input := &resourceexplorer2.DisassociateDefaultViewInput{}

	if resp, err := client.DisassociateDefaultView(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the status of your account's Amazon Web Services service access, and
// validates the service linked role required to access the multi-account search
// feature. Only the management account can invoke this API call.
func resourceexplorer2_GetAccountLevelServiceConfiguration(cfg aws.Config, client *resourceexplorer2.Client) {
	input := &resourceexplorer2.GetAccountLevelServiceConfigurationInput{}

	if resp, err := client.GetAccountLevelServiceConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the Amazon Resource Name (ARN) of the view that is the default for
// the Amazon Web Services Region in which you call this operation. You can then
// call GetViewto retrieve the details of that view.
func resourceexplorer2_GetDefaultView(cfg aws.Config, client *resourceexplorer2.Client) {
	input := &resourceexplorer2.GetDefaultViewInput{}

	if resp, err := client.GetDefaultView(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details about the Amazon Web Services Resource Explorer index in the
// Amazon Web Services Region in which you invoked the operation.
func resourceexplorer2_GetIndex(cfg aws.Config, client *resourceexplorer2.Client) {
	input := &resourceexplorer2.GetIndexInput{}

	if resp, err := client.GetIndex(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details of the specified [Amazon Web Services-managed view].
//
// [Amazon Web Services-managed view]: https://docs.aws.amazon.com/resource-explorer/latest/userguide/aws-managed-views.html
func resourceexplorer2_GetManagedView(cfg aws.Config, client *resourceexplorer2.Client) {
	input := &resourceexplorer2.GetManagedViewInput{
		// ManagedViewArn: *string, // Required
	}

	if len(_resourceexplorer2ManagedViewArn) > 0 {
		input.ManagedViewArn = aws.String(_resourceexplorer2ManagedViewArn)
	}

	if resp, err := client.GetManagedView(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the status and details of a Resource Explorer setup operation. This
// operation returns information about the progress of creating or deleting
// Resource Explorer configurations across Regions.
func resourceexplorer2_GetResourceExplorerSetup(cfg aws.Config, client *resourceexplorer2.Client) {
	input := &resourceexplorer2.GetResourceExplorerSetupInput{
		// TaskId: *string, // Required
	}

	if len(_resourceexplorer2TaskId) > 0 {
		input.TaskId = aws.String(_resourceexplorer2TaskId)
	}
	if len(_resourceexplorer2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _resourceexplorer2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_resourceexplorer2NextToken) > 0 {
		input.NextToken = aws.String(_resourceexplorer2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetResourceExplorerSetup(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*resourceexplorer2.GetResourceExplorerSetupOutput
	p := resourceexplorer2.NewGetResourceExplorerSetupPaginator(client, input)
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

// Retrieves information about the Resource Explorer index in the current Amazon
// Web Services Region. This operation returns the ARN and type of the index if one
// exists.
func resourceexplorer2_GetServiceIndex(cfg aws.Config, client *resourceexplorer2.Client) {
	input := &resourceexplorer2.GetServiceIndexInput{}

	if resp, err := client.GetServiceIndex(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details about a specific Resource Explorer service view. This
// operation returns the configuration and properties of the specified view.
func resourceexplorer2_GetServiceView(cfg aws.Config, client *resourceexplorer2.Client) {
	input := &resourceexplorer2.GetServiceViewInput{
		// ServiceViewArn: *string, // Required
	}

	if len(_resourceexplorer2ServiceViewArn) > 0 {
		input.ServiceViewArn = aws.String(_resourceexplorer2ServiceViewArn)
	}

	if resp, err := client.GetServiceView(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details of the specified view.
func resourceexplorer2_GetView(cfg aws.Config, client *resourceexplorer2.Client) {
	input := &resourceexplorer2.GetViewInput{
		// ViewArn: *string, // Required
	}

	if len(_resourceexplorer2ViewArn) > 0 {
		input.ViewArn = aws.String(_resourceexplorer2ViewArn)
	}

	if resp, err := client.GetView(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of all of the indexes in Amazon Web Services Regions that are
// currently collecting resource information for Amazon Web Services Resource
// Explorer.
func resourceexplorer2_ListIndexes(cfg aws.Config, client *resourceexplorer2.Client) {
	input := &resourceexplorer2.ListIndexesInput{}

	if len(_resourceexplorer2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _resourceexplorer2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_resourceexplorer2NextToken) > 0 {
		input.NextToken = aws.String(_resourceexplorer2NextToken)
	}
	if len(_resourceexplorer2Regions) > 0 {
		input.Regions = append([]string(nil), _resourceexplorer2Regions...)
	}
	if len(_resourceexplorer2Type) > 0 {
		if err := assignInputField(input, "Type", _resourceexplorer2Type); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListIndexes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*resourceexplorer2.ListIndexesOutput
	p := resourceexplorer2.NewListIndexesPaginator(client, input)
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

// Retrieves a list of a member's indexes in all Amazon Web Services Regions that
// are currently collecting resource information for Amazon Web Services Resource
// Explorer. Only the management account or a delegated administrator with service
// access enabled can invoke this API call.
func resourceexplorer2_ListIndexesForMembers(cfg aws.Config, client *resourceexplorer2.Client) {
	input := &resourceexplorer2.ListIndexesForMembersInput{
		// AccountIdList: []string, // Required
	}

	if len(_resourceexplorer2AccountIdList) > 0 {
		input.AccountIdList = append([]string(nil), _resourceexplorer2AccountIdList...)
	}
	if len(_resourceexplorer2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _resourceexplorer2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_resourceexplorer2NextToken) > 0 {
		input.NextToken = aws.String(_resourceexplorer2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListIndexesForMembers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*resourceexplorer2.ListIndexesForMembersOutput
	p := resourceexplorer2.NewListIndexesForMembersPaginator(client, input)
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

// Lists the Amazon resource names (ARNs) of the [Amazon Web Services-managed views] available in the Amazon Web
// Services Region in which you call this operation.
//
// [Amazon Web Services-managed views]: https://docs.aws.amazon.com/resource-explorer/latest/userguide/aws-managed-views.html
func resourceexplorer2_ListManagedViews(cfg aws.Config, client *resourceexplorer2.Client) {
	input := &resourceexplorer2.ListManagedViewsInput{}

	if len(_resourceexplorer2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _resourceexplorer2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_resourceexplorer2NextToken) > 0 {
		input.NextToken = aws.String(_resourceexplorer2NextToken)
	}
	if len(_resourceexplorer2ServicePrincipal) > 0 {
		input.ServicePrincipal = aws.String(_resourceexplorer2ServicePrincipal)
	}

	if disablePaginator() {
		if resp, err := client.ListManagedViews(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*resourceexplorer2.ListManagedViewsOutput
	p := resourceexplorer2.NewListManagedViewsPaginator(client, input)
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

// Returns a list of resources and their details that match the specified
// criteria. This query must use a view. If you don’t explicitly specify a view,
// then Resource Explorer uses the default view for the Amazon Web Services Region
// in which you call this operation.
func resourceexplorer2_ListResources(cfg aws.Config, client *resourceexplorer2.Client) {
	input := &resourceexplorer2.ListResourcesInput{}

	if len(_resourceexplorer2Filters) > 0 {
		if err := assignInputField(input, "Filters", _resourceexplorer2Filters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_resourceexplorer2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _resourceexplorer2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_resourceexplorer2NextToken) > 0 {
		input.NextToken = aws.String(_resourceexplorer2NextToken)
	}
	if len(_resourceexplorer2ViewArn) > 0 {
		input.ViewArn = aws.String(_resourceexplorer2ViewArn)
	}

	if disablePaginator() {
		if resp, err := client.ListResources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*resourceexplorer2.ListResourcesOutput
	p := resourceexplorer2.NewListResourcesPaginator(client, input)
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

// Lists all Resource Explorer indexes across the specified Amazon Web Services
// Regions. This operation returns information about indexes including their ARNs,
// types, and Regions.
func resourceexplorer2_ListServiceIndexes(cfg aws.Config, client *resourceexplorer2.Client) {
	input := &resourceexplorer2.ListServiceIndexesInput{}

	if len(_resourceexplorer2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _resourceexplorer2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_resourceexplorer2NextToken) > 0 {
		input.NextToken = aws.String(_resourceexplorer2NextToken)
	}
	if len(_resourceexplorer2Regions) > 0 {
		input.Regions = append([]string(nil), _resourceexplorer2Regions...)
	}

	if disablePaginator() {
		if resp, err := client.ListServiceIndexes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*resourceexplorer2.ListServiceIndexesOutput
	p := resourceexplorer2.NewListServiceIndexesPaginator(client, input)
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

// Lists all Resource Explorer service views available in the current Amazon Web
// Services account. This operation returns the ARNs of available service views.
func resourceexplorer2_ListServiceViews(cfg aws.Config, client *resourceexplorer2.Client) {
	input := &resourceexplorer2.ListServiceViewsInput{}

	if len(_resourceexplorer2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _resourceexplorer2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_resourceexplorer2NextToken) > 0 {
		input.NextToken = aws.String(_resourceexplorer2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListServiceViews(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*resourceexplorer2.ListServiceViewsOutput
	p := resourceexplorer2.NewListServiceViewsPaginator(client, input)
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

// Returns a list of Amazon Web Services services that have been granted streaming
// access to your Resource Explorer data. Streaming access allows Amazon Web
// Services services to receive real-time updates about your resources as they are
// indexed by Resource Explorer.
func resourceexplorer2_ListStreamingAccessForServices(cfg aws.Config, client *resourceexplorer2.Client) {
	input := &resourceexplorer2.ListStreamingAccessForServicesInput{}

	if len(_resourceexplorer2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _resourceexplorer2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_resourceexplorer2NextToken) > 0 {
		input.NextToken = aws.String(_resourceexplorer2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListStreamingAccessForServices(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*resourceexplorer2.ListStreamingAccessForServicesOutput
	p := resourceexplorer2.NewListStreamingAccessForServicesPaginator(client, input)
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

// Retrieves a list of all resource types currently supported by Amazon Web
// Services Resource Explorer.
func resourceexplorer2_ListSupportedResourceTypes(cfg aws.Config, client *resourceexplorer2.Client) {
	input := &resourceexplorer2.ListSupportedResourceTypesInput{}

	if len(_resourceexplorer2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _resourceexplorer2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_resourceexplorer2NextToken) > 0 {
		input.NextToken = aws.String(_resourceexplorer2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSupportedResourceTypes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*resourceexplorer2.ListSupportedResourceTypesOutput
	p := resourceexplorer2.NewListSupportedResourceTypesPaginator(client, input)
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

// Lists the tags that are attached to the specified resource.
func resourceexplorer2_ListTagsForResource(cfg aws.Config, client *resourceexplorer2.Client) {
	input := &resourceexplorer2.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_resourceexplorer2ResourceArn) > 0 {
		input.ResourceArn = aws.String(_resourceexplorer2ResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the [Amazon resource names (ARNs)] of the views available in the Amazon Web Services Region in which
// you call this operation.
//
// Always check the NextToken response parameter for a null value when calling a
// paginated operation. These operations can occasionally return an empty set of
// results even when there are more results available. The NextToken response
// parameter value is null only when there are no more results to display.
//
// [Amazon resource names (ARNs)]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
func resourceexplorer2_ListViews(cfg aws.Config, client *resourceexplorer2.Client) {
	input := &resourceexplorer2.ListViewsInput{}

	if len(_resourceexplorer2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _resourceexplorer2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_resourceexplorer2NextToken) > 0 {
		input.NextToken = aws.String(_resourceexplorer2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListViews(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*resourceexplorer2.ListViewsOutput
	p := resourceexplorer2.NewListViewsPaginator(client, input)
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

// Searches for resources and displays details about all resources that match the
// specified criteria. You must specify a query string.
//
// All search queries must use a view. If you don't explicitly specify a view,
// then Amazon Web Services Resource Explorer uses the default view for the Amazon
// Web Services Region in which you call this operation. The results are the
// logical intersection of the results that match both the QueryString parameter
// supplied to this operation and the SearchFilter parameter attached to the view.
//
// For the complete syntax supported by the QueryString parameter, see [Search query syntax reference for Resource Explorer].
//
// If your search results are empty, or are missing results that you think should
// be there, see [Troubleshooting Resource Explorer search].
//
// [Troubleshooting Resource Explorer search]: https://docs.aws.amazon.com/resource-explorer/latest/userguide/troubleshooting_search.html
// [Search query syntax reference for Resource Explorer]: https://docs.aws.amazon.com/resource-explorer/latest/APIReference/about-query-syntax.html
func resourceexplorer2_Search(cfg aws.Config, client *resourceexplorer2.Client) {
	input := &resourceexplorer2.SearchInput{
		// QueryString: *string, // Required
	}

	if len(_resourceexplorer2QueryString) > 0 {
		input.QueryString = aws.String(_resourceexplorer2QueryString)
	}
	if len(_resourceexplorer2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _resourceexplorer2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_resourceexplorer2NextToken) > 0 {
		input.NextToken = aws.String(_resourceexplorer2NextToken)
	}
	if len(_resourceexplorer2ViewArn) > 0 {
		input.ViewArn = aws.String(_resourceexplorer2ViewArn)
	}

	if disablePaginator() {
		if resp, err := client.Search(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*resourceexplorer2.SearchOutput
	p := resourceexplorer2.NewSearchPaginator(client, input)
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

// Adds one or more tag key and value pairs to an Amazon Web Services Resource
// Explorer view or index.
func resourceexplorer2_TagResource(cfg aws.Config, client *resourceexplorer2.Client) {
	input := &resourceexplorer2.TagResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_resourceexplorer2ResourceArn) > 0 {
		input.ResourceArn = aws.String(_resourceexplorer2ResourceArn)
	}
	if len(_resourceexplorer2Tags) > 0 {
		if err := assignInputField(input, "Tags", _resourceexplorer2Tags); err != nil {
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

// Removes one or more tag key and value pairs from an Amazon Web Services
// Resource Explorer view or index.
func resourceexplorer2_UntagResource(cfg aws.Config, client *resourceexplorer2.Client) {
	input := &resourceexplorer2.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_resourceexplorer2ResourceArn) > 0 {
		input.ResourceArn = aws.String(_resourceexplorer2ResourceArn)
	}
	if len(_resourceexplorer2TagKeys) > 0 {
		input.TagKeys = append([]string(nil), _resourceexplorer2TagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes the type of the index from one of the following types to the other. For
// more information about indexes and the role they perform in Amazon Web Services
// Resource Explorer, see [Turning on cross-Region search by creating an aggregator index]in the Amazon Web Services Resource Explorer User Guide.
//
// - AGGREGATOR index type
//
// # The index contains information about resources from all Amazon Web Services
//
// Regions in the Amazon Web Services account in which you've created a Resource
// Explorer index. Resource information from all other Regions is replicated to
// this Region's index.
//
// # When you change the index type to AGGREGATOR , Resource Explorer turns on
//
// replication of all discovered resource information from the other Amazon Web
// Services Regions in your account to this index. You can then, from this Region
// only, perform resource search queries that span all Amazon Web Services Regions
// in the Amazon Web Services account. Turning on replication from all other
// Regions is performed by asynchronous background tasks. You can check the status
// of the asynchronous tasks by using the GetIndexoperation. When the asynchronous tasks
// complete, the Status response of that operation changes from UPDATING to
// ACTIVE . After that, you can start to see results from other Amazon Web
// Services Regions in query results. However, it can take several hours for
// replication from all other Regions to complete.
//
// You can have only one aggregator index per Amazon Web Services account. Before
//
// you can promote a different index to be the aggregator index for the account,
// you must first demote the existing aggregator index to type LOCAL .
//
// - LOCAL index type
//
// # The index contains information about resources in only the Amazon Web Services
//
// Region in which the index exists. If an aggregator index in another Region
// exists, then information in this local index is replicated to the aggregator
// index.
//
// # When you change the index type to LOCAL , Resource Explorer turns off the
//
// replication of resource information from all other Amazon Web Services Regions
// in the Amazon Web Services account to this Region. The aggregator index remains
// in the UPDATING state until all replication with other Regions successfully
// stops. You can check the status of the asynchronous task by using the GetIndex
// operation. When Resource Explorer successfully stops all replication with other
// Regions, the Status response of that operation changes from UPDATING to ACTIVE
// . Separately, the resource information from other Regions that was previously
// stored in the index is deleted within 30 days by another background task. Until
// that asynchronous task completes, some results from other Regions can continue
// to appear in search results.
//
// # After you demote an aggregator index to a local index, you must wait 24 hours
//
// before you can promote another index to be the new aggregator index for the
// account.
//
// [Turning on cross-Region search by creating an aggregator index]: https://docs.aws.amazon.com/resource-explorer/latest/userguide/manage-aggregator-region.html
func resourceexplorer2_UpdateIndexType(cfg aws.Config, client *resourceexplorer2.Client) {
	input := &resourceexplorer2.UpdateIndexTypeInput{
		// Arn: *string, // Required
		// Type: types.IndexType, // Required
	}

	if len(_resourceexplorer2Arn) > 0 {
		input.Arn = aws.String(_resourceexplorer2Arn)
	}
	if len(_resourceexplorer2Type) > 0 {
		if err := assignInputField(input, "Type", _resourceexplorer2Type); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateIndexType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies some of the details of a view. You can change the filter string and
// the list of included properties. You can't change the name of the view.
func resourceexplorer2_UpdateView(cfg aws.Config, client *resourceexplorer2.Client) {
	input := &resourceexplorer2.UpdateViewInput{
		// ViewArn: *string, // Required
	}

	if len(_resourceexplorer2ViewArn) > 0 {
		input.ViewArn = aws.String(_resourceexplorer2ViewArn)
	}
	if len(_resourceexplorer2Filters) > 0 {
		if err := assignInputField(input, "Filters", _resourceexplorer2Filters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_resourceexplorer2IncludedProperties) > 0 {
		if err := assignInputField(input, "IncludedProperties", _resourceexplorer2IncludedProperties); err != nil {
			log.Errorf("invalid --included-properties: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateView(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_resourceexplorer2Cmd)
	_resourceexplorer2Cmd.Flags().SortFlags = false

	_resourceexplorer2Cmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_resourceexplorer2Cmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_resourceexplorer2Cmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_resourceexplorer2Cmd.Flags().StringSliceVarP(&_resourceexplorer2AccountIdList, "account-id-list", "", nil, "Account ID List")
	_resourceexplorer2Cmd.Flags().StringSliceVarP(&_resourceexplorer2AggregatorRegions, "aggregator-regions", "", nil, "Aggregator Regions")
	_resourceexplorer2Cmd.Flags().StringVarP(&_resourceexplorer2Arn, "arn", "", "", "ARN")
	_resourceexplorer2Cmd.Flags().StringVarP(&_resourceexplorer2ClientToken, "client-token", "", "", "Client Token")
	_resourceexplorer2Cmd.Flags().StringVarP(&_resourceexplorer2DeleteInAllRegions, "delete-in-all-regions", "", "", "Delete In All Regions")
	_resourceexplorer2Cmd.Flags().StringVarP(&_resourceexplorer2Filters, "filters", "", "", "Filters")
	_resourceexplorer2Cmd.Flags().StringVarP(&_resourceexplorer2IncludedProperties, "included-properties", "", "", "Included Properties")
	_resourceexplorer2Cmd.Flags().StringVarP(&_resourceexplorer2ManagedViewArn, "managed-view-arn", "", "", "Managed View ARN")
	_resourceexplorer2Cmd.Flags().StringVarP(&_resourceexplorer2MaxResults, "max-results", "", "", "Max Results")
	_resourceexplorer2Cmd.Flags().StringVarP(&_resourceexplorer2NextToken, "next-token", "", "", "Next Token")
	_resourceexplorer2Cmd.Flags().StringVarP(&_resourceexplorer2QueryString, "query-string", "", "", "Query String")
	_resourceexplorer2Cmd.Flags().StringSliceVarP(&_resourceexplorer2RegionList, "region-list", "", nil, "Region List")
	_resourceexplorer2Cmd.Flags().StringSliceVarP(&_resourceexplorer2Regions, "regions", "", nil, "Regions")
	_resourceexplorer2Cmd.Flags().StringVarP(&_resourceexplorer2ResourceArn, "resource-arn", "", "", "Resource ARN")
	_resourceexplorer2Cmd.Flags().StringVarP(&_resourceexplorer2Scope, "scope", "", "", "Scope")
	_resourceexplorer2Cmd.Flags().StringVarP(&_resourceexplorer2ServicePrincipal, "service-principal", "", "", "Service Principal")
	_resourceexplorer2Cmd.Flags().StringVarP(&_resourceexplorer2ServiceViewArn, "service-view-arn", "", "", "Service View ARN")
	_resourceexplorer2Cmd.Flags().StringSliceVarP(&_resourceexplorer2TagKeys, "tag-keys", "", nil, "Tag Keys")
	_resourceexplorer2Cmd.Flags().StringVarP(&_resourceexplorer2Tags, "tags", "", "", "Tags")
	_resourceexplorer2Cmd.Flags().StringVarP(&_resourceexplorer2TaskId, "task-id", "", "", "Task ID")
	_resourceexplorer2Cmd.Flags().StringVarP(&_resourceexplorer2Type, "type", "", "", "Type")
	_resourceexplorer2Cmd.Flags().StringVarP(&_resourceexplorer2ViewArn, "view-arn", "", "", "View ARN")
	_resourceexplorer2Cmd.Flags().StringSliceVarP(&_resourceexplorer2ViewArns, "view-arns", "", nil, "View Arns")
	_resourceexplorer2Cmd.Flags().StringVarP(&_resourceexplorer2ViewName, "view-name", "", "", "View Name")

	_resourceexplorer2Cmd.Flags().BoolVarP(&_resourceexplorer2AssociateDefaultView, "associate-default-view", "", false, "Associate Default View")
	_resourceexplorer2Cmd.Flags().BoolVarP(&_resourceexplorer2BatchGetView, "batch-get-view", "", false, "Batch Get View")
	_resourceexplorer2Cmd.Flags().BoolVarP(&_resourceexplorer2CreateIndex, "create-index", "", false, "Create Index")
	_resourceexplorer2Cmd.Flags().BoolVarP(&_resourceexplorer2CreateResourceExplorerSetup, "create-resource-explorer-setup", "", false, "Create Resource Explorer Setup")
	_resourceexplorer2Cmd.Flags().BoolVarP(&_resourceexplorer2CreateView, "create-view", "", false, "Create View")
	_resourceexplorer2Cmd.Flags().BoolVarP(&_resourceexplorer2DeleteIndex, "delete-index", "", false, "Delete Index")
	_resourceexplorer2Cmd.Flags().BoolVarP(&_resourceexplorer2DeleteResourceExplorerSetup, "delete-resource-explorer-setup", "", false, "Delete Resource Explorer Setup")
	_resourceexplorer2Cmd.Flags().BoolVarP(&_resourceexplorer2DeleteView, "delete-view", "", false, "Delete View")
	_resourceexplorer2Cmd.Flags().BoolVarP(&_resourceexplorer2DisassociateDefaultView, "disassociate-default-view", "", false, "Disassociate Default View")
	_resourceexplorer2Cmd.Flags().BoolVarP(&_resourceexplorer2GetAccountLevelServiceConfiguration, "get-account-level-service-configuration", "", false, "Get Account Level Service Configuration")
	_resourceexplorer2Cmd.Flags().BoolVarP(&_resourceexplorer2GetDefaultView, "get-default-view", "", false, "Get Default View")
	_resourceexplorer2Cmd.Flags().BoolVarP(&_resourceexplorer2GetIndex, "get-index", "", false, "Get Index")
	_resourceexplorer2Cmd.Flags().BoolVarP(&_resourceexplorer2GetManagedView, "get-managed-view", "", false, "Get Managed View")
	_resourceexplorer2Cmd.Flags().BoolVarP(&_resourceexplorer2GetResourceExplorerSetup, "get-resource-explorer-setup", "", false, "Get Resource Explorer Setup")
	_resourceexplorer2Cmd.Flags().BoolVarP(&_resourceexplorer2GetServiceIndex, "get-service-index", "", false, "Get Service Index")
	_resourceexplorer2Cmd.Flags().BoolVarP(&_resourceexplorer2GetServiceView, "get-service-view", "", false, "Get Service View")
	_resourceexplorer2Cmd.Flags().BoolVarP(&_resourceexplorer2GetView, "get-view", "", false, "Get View")
	_resourceexplorer2Cmd.Flags().BoolVarP(&_resourceexplorer2ListIndexes, "list-indexes", "", false, "List Indexes")
	_resourceexplorer2Cmd.Flags().BoolVarP(&_resourceexplorer2ListIndexesForMembers, "list-indexes-for-members", "", false, "List Indexes For Members")
	_resourceexplorer2Cmd.Flags().BoolVarP(&_resourceexplorer2ListManagedViews, "list-managed-views", "", false, "List Managed Views")
	_resourceexplorer2Cmd.Flags().BoolVarP(&_resourceexplorer2ListResources, "list-resources", "", false, "List Resources")
	_resourceexplorer2Cmd.Flags().BoolVarP(&_resourceexplorer2ListServiceIndexes, "list-service-indexes", "", false, "List Service Indexes")
	_resourceexplorer2Cmd.Flags().BoolVarP(&_resourceexplorer2ListServiceViews, "list-service-views", "", false, "List Service Views")
	_resourceexplorer2Cmd.Flags().BoolVarP(&_resourceexplorer2ListStreamingAccessForServices, "list-streaming-access-for-services", "", false, "List Streaming Access For Services")
	_resourceexplorer2Cmd.Flags().BoolVarP(&_resourceexplorer2ListSupportedResourceTypes, "list-supported-resource-types", "", false, "List Supported Resource Types")
	_resourceexplorer2Cmd.Flags().BoolVarP(&_resourceexplorer2ListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_resourceexplorer2Cmd.Flags().BoolVarP(&_resourceexplorer2ListViews, "list-views", "", false, "List Views")
	_resourceexplorer2Cmd.Flags().BoolVarP(&_resourceexplorer2Search, "search", "", false, "Search")
	_resourceexplorer2Cmd.Flags().BoolVarP(&_resourceexplorer2TagResource, "tag-resource", "", false, "Tag Resource")
	_resourceexplorer2Cmd.Flags().BoolVarP(&_resourceexplorer2UntagResource, "untag-resource", "", false, "Untag Resource")
	_resourceexplorer2Cmd.Flags().BoolVarP(&_resourceexplorer2UpdateIndexType, "update-index-type", "", false, "Update Index Type")
	_resourceexplorer2Cmd.Flags().BoolVarP(&_resourceexplorer2UpdateView, "update-view", "", false, "Update View")

}
