package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroups"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// resourcegroupsCmd represents the resourcegroups command
var _resourcegroupsCmd = &cobra.Command{
	Use:   "resourcegroups",
	Short: "AWS resourcegroups CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := resourcegroups.NewFromConfig(cfg)
		if _resourcegroupsCancelTagSyncTask {
			resourcegroups_CancelTagSyncTask(cfg, client)
			return
		}
		if _resourcegroupsCreateGroup {
			resourcegroups_CreateGroup(cfg, client)
			return
		}
		if _resourcegroupsDeleteGroup {
			resourcegroups_DeleteGroup(cfg, client)
			return
		}
		if _resourcegroupsGetAccountSettings {
			resourcegroups_GetAccountSettings(cfg, client)
			return
		}
		if _resourcegroupsGetGroup {
			resourcegroups_GetGroup(cfg, client)
			return
		}
		if _resourcegroupsGetGroupConfiguration {
			resourcegroups_GetGroupConfiguration(cfg, client)
			return
		}
		if _resourcegroupsGetGroupQuery {
			resourcegroups_GetGroupQuery(cfg, client)
			return
		}
		if _resourcegroupsGetTagSyncTask {
			resourcegroups_GetTagSyncTask(cfg, client)
			return
		}
		if _resourcegroupsGetTags {
			resourcegroups_GetTags(cfg, client)
			return
		}
		if _resourcegroupsGroupResources {
			resourcegroups_GroupResources(cfg, client)
			return
		}
		if _resourcegroupsListGroupResources {
			resourcegroups_ListGroupResources(cfg, client)
			return
		}
		if _resourcegroupsListGroupingStatuses {
			resourcegroups_ListGroupingStatuses(cfg, client)
			return
		}
		if _resourcegroupsListGroups {
			resourcegroups_ListGroups(cfg, client)
			return
		}
		if _resourcegroupsListTagSyncTasks {
			resourcegroups_ListTagSyncTasks(cfg, client)
			return
		}
		if _resourcegroupsPutGroupConfiguration {
			resourcegroups_PutGroupConfiguration(cfg, client)
			return
		}
		if _resourcegroupsSearchResources {
			resourcegroups_SearchResources(cfg, client)
			return
		}
		if _resourcegroupsStartTagSyncTask {
			resourcegroups_StartTagSyncTask(cfg, client)
			return
		}
		if _resourcegroupsTag {
			resourcegroups_Tag(cfg, client)
			return
		}
		if _resourcegroupsUngroupResources {
			resourcegroups_UngroupResources(cfg, client)
			return
		}
		if _resourcegroupsUntag {
			resourcegroups_Untag(cfg, client)
			return
		}
		if _resourcegroupsUpdateAccountSettings {
			resourcegroups_UpdateAccountSettings(cfg, client)
			return
		}
		if _resourcegroupsUpdateGroup {
			resourcegroups_UpdateGroup(cfg, client)
			return
		}
		if _resourcegroupsUpdateGroupQuery {
			resourcegroups_UpdateGroupQuery(cfg, client)
			return
		}

	},
}

var (
	_resourcegroupsCancelTagSyncTask     bool
	_resourcegroupsCreateGroup           bool
	_resourcegroupsDeleteGroup           bool
	_resourcegroupsGetAccountSettings    bool
	_resourcegroupsGetGroup              bool
	_resourcegroupsGetGroupConfiguration bool
	_resourcegroupsGetGroupQuery         bool
	_resourcegroupsGetTagSyncTask        bool
	_resourcegroupsGetTags               bool
	_resourcegroupsGroupResources        bool
	_resourcegroupsListGroupResources    bool
	_resourcegroupsListGroupingStatuses  bool
	_resourcegroupsListGroups            bool
	_resourcegroupsListTagSyncTasks      bool
	_resourcegroupsPutGroupConfiguration bool
	_resourcegroupsSearchResources       bool
	_resourcegroupsStartTagSyncTask      bool
	_resourcegroupsTag                   bool
	_resourcegroupsUngroupResources      bool
	_resourcegroupsUntag                 bool
	_resourcegroupsUpdateAccountSettings bool
	_resourcegroupsUpdateGroup           bool
	_resourcegroupsUpdateGroupQuery      bool

	_resourcegroupsArn                               string
	_resourcegroupsConfiguration                     string
	_resourcegroupsCriticality                       string
	_resourcegroupsDescription                       string
	_resourcegroupsDisplayName                       string
	_resourcegroupsFilters                           string
	_resourcegroupsGroup                             string
	_resourcegroupsGroupLifecycleEventsDesiredStatus string
	_resourcegroupsGroupName                         string
	_resourcegroupsKeys                              []string
	_resourcegroupsMaxResults                        string
	_resourcegroupsName                              string
	_resourcegroupsNextToken                         string
	_resourcegroupsOwner                             string
	_resourcegroupsResourceArns                      []string
	_resourcegroupsResourceQuery                     string
	_resourcegroupsRoleArn                           string
	_resourcegroupsTagKey                            string
	_resourcegroupsTagValue                          string
	_resourcegroupsTags                              string
	_resourcegroupsTaskArn                           string
)

// Cancels the specified tag-sync task.
// # Minimum permissions
//
// To run this command, you must have the following permissions:
//
// - resource-groups:CancelTagSyncTask on the application group
//
// - resource-groups:DeleteGroup
func resourcegroups_CancelTagSyncTask(cfg aws.Config, client *resourcegroups.Client) {
	input := &resourcegroups.CancelTagSyncTaskInput{
		// TaskArn: *string, // Required
	}

	if len(_resourcegroupsTaskArn) > 0 {
		input.TaskArn = aws.String(_resourcegroupsTaskArn)
	}

	if resp, err := client.CancelTagSyncTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a resource group with the specified name and description. You can
// optionally include either a resource query or a service configuration. For more
// information about constructing a resource query, see [Build queries and groups in Resource Groups]in the Resource Groups
// User Guide. For more information about service-linked groups and service
// configurations, see [Service configurations for Resource Groups].
//
// # Minimum permissions
//
// To run this command, you must have the following permissions:
//
// - resource-groups:CreateGroup
//
// [Build queries and groups in Resource Groups]: https://docs.aws.amazon.com/ARG/latest/userguide/getting_started-query.html
// [Service configurations for Resource Groups]: https://docs.aws.amazon.com/ARG/latest/APIReference/about-slg.html
func resourcegroups_CreateGroup(cfg aws.Config, client *resourcegroups.Client) {
	input := &resourcegroups.CreateGroupInput{
		// Name: *string, // Required
	}

	if len(_resourcegroupsName) > 0 {
		input.Name = aws.String(_resourcegroupsName)
	}
	if len(_resourcegroupsConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _resourcegroupsConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_resourcegroupsCriticality) > 0 {
		if err := assignInputField(input, "Criticality", _resourcegroupsCriticality); err != nil {
			log.Errorf("invalid --criticality: %s", err.Error())
			return
		}
	}
	if len(_resourcegroupsDescription) > 0 {
		input.Description = aws.String(_resourcegroupsDescription)
	}
	if len(_resourcegroupsDisplayName) > 0 {
		input.DisplayName = aws.String(_resourcegroupsDisplayName)
	}
	if len(_resourcegroupsOwner) > 0 {
		input.Owner = aws.String(_resourcegroupsOwner)
	}
	if len(_resourcegroupsResourceQuery) > 0 {
		if err := assignInputField(input, "ResourceQuery", _resourcegroupsResourceQuery); err != nil {
			log.Errorf("invalid --resource-query: %s", err.Error())
			return
		}
	}
	if len(_resourcegroupsTags) > 0 {
		if err := assignInputField(input, "Tags", _resourcegroupsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified resource group. Deleting a resource group does not delete
// any resources that are members of the group; it only deletes the group
// structure.
//
// # Minimum permissions
//
// To run this command, you must have the following permissions:
//
// - resource-groups:DeleteGroup
func resourcegroups_DeleteGroup(cfg aws.Config, client *resourcegroups.Client) {
	input := &resourcegroups.DeleteGroupInput{}

	if len(_resourcegroupsGroup) > 0 {
		input.Group = aws.String(_resourcegroupsGroup)
	}
	if len(_resourcegroupsGroupName) > 0 {
		input.GroupName = aws.String(_resourcegroupsGroupName)
	}

	if resp, err := client.DeleteGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the current status of optional features in Resource Groups.
func resourcegroups_GetAccountSettings(cfg aws.Config, client *resourcegroups.Client) {
	input := &resourcegroups.GetAccountSettingsInput{}

	if resp, err := client.GetAccountSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a specified resource group.
// # Minimum permissions
//
// To run this command, you must have the following permissions:
//
// - resource-groups:GetGroup
func resourcegroups_GetGroup(cfg aws.Config, client *resourcegroups.Client) {
	input := &resourcegroups.GetGroupInput{}

	if len(_resourcegroupsGroup) > 0 {
		input.Group = aws.String(_resourcegroupsGroup)
	}
	if len(_resourcegroupsGroupName) > 0 {
		input.GroupName = aws.String(_resourcegroupsGroupName)
	}

	if resp, err := client.GetGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the service configuration associated with the specified resource
// group. For details about the service configuration syntax, see [Service configurations for Resource Groups].
//
// # Minimum permissions
//
// To run this command, you must have the following permissions:
//
// - resource-groups:GetGroupConfiguration
//
// [Service configurations for Resource Groups]: https://docs.aws.amazon.com/ARG/latest/APIReference/about-slg.html
func resourcegroups_GetGroupConfiguration(cfg aws.Config, client *resourcegroups.Client) {
	input := &resourcegroups.GetGroupConfigurationInput{}

	if len(_resourcegroupsGroup) > 0 {
		input.Group = aws.String(_resourcegroupsGroup)
	}

	if resp, err := client.GetGroupConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the resource query associated with the specified resource group. For
// more information about resource queries, see [Create a tag-based group in Resource Groups].
//
// # Minimum permissions
//
// To run this command, you must have the following permissions:
//
// - resource-groups:GetGroupQuery
//
// [Create a tag-based group in Resource Groups]: https://docs.aws.amazon.com/ARG/latest/userguide/gettingstarted-query.html#gettingstarted-query-cli-tag
func resourcegroups_GetGroupQuery(cfg aws.Config, client *resourcegroups.Client) {
	input := &resourcegroups.GetGroupQueryInput{}

	if len(_resourcegroupsGroup) > 0 {
		input.Group = aws.String(_resourcegroupsGroup)
	}
	if len(_resourcegroupsGroupName) > 0 {
		input.GroupName = aws.String(_resourcegroupsGroupName)
	}

	if resp, err := client.GetGroupQuery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a specified tag-sync task.
// # Minimum permissions
//
// To run this command, you must have the following permissions:
//
// - resource-groups:GetTagSyncTask on the application group
func resourcegroups_GetTagSyncTask(cfg aws.Config, client *resourcegroups.Client) {
	input := &resourcegroups.GetTagSyncTaskInput{
		// TaskArn: *string, // Required
	}

	if len(_resourcegroupsTaskArn) > 0 {
		input.TaskArn = aws.String(_resourcegroupsTaskArn)
	}

	if resp, err := client.GetTagSyncTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of tags that are associated with a resource group, specified by
// an Amazon resource name (ARN).
//
// # Minimum permissions
//
// To run this command, you must have the following permissions:
//
// - resource-groups:GetTags
func resourcegroups_GetTags(cfg aws.Config, client *resourcegroups.Client) {
	input := &resourcegroups.GetTagsInput{
		// Arn: *string, // Required
	}

	if len(_resourcegroupsArn) > 0 {
		input.Arn = aws.String(_resourcegroupsArn)
	}

	if resp, err := client.GetTags(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds the specified resources to the specified group.
// You can only use this operation with the following groups:
//
// - AWS::EC2::HostManagement
//
// - AWS::EC2::CapacityReservationPool
//
// - AWS::ResourceGroups::ApplicationGroup
//
// Other resource group types and resource types are not currently supported by
// this operation.
//
// # Minimum permissions
//
// To run this command, you must have the following permissions:
//
// - resource-groups:GroupResources
func resourcegroups_GroupResources(cfg aws.Config, client *resourcegroups.Client) {
	input := &resourcegroups.GroupResourcesInput{
		// Group: *string, // Required
		// ResourceArns: []string, // Required
	}

	if len(_resourcegroupsGroup) > 0 {
		input.Group = aws.String(_resourcegroupsGroup)
	}
	if len(_resourcegroupsResourceArns) > 0 {
		input.ResourceArns = append([]string(nil), _resourcegroupsResourceArns...)
	}

	if resp, err := client.GroupResources(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of Amazon resource names (ARNs) of the resources that are
// members of a specified resource group.
//
// # Minimum permissions
//
// To run this command, you must have the following permissions:
//
// - resource-groups:ListGroupResources
//
// - cloudformation:DescribeStacks
//
// - cloudformation:ListStackResources
//
// - tag:GetResources
func resourcegroups_ListGroupResources(cfg aws.Config, client *resourcegroups.Client) {
	input := &resourcegroups.ListGroupResourcesInput{}

	if len(_resourcegroupsFilters) > 0 {
		if err := assignInputField(input, "Filters", _resourcegroupsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_resourcegroupsGroup) > 0 {
		input.Group = aws.String(_resourcegroupsGroup)
	}
	if len(_resourcegroupsGroupName) > 0 {
		input.GroupName = aws.String(_resourcegroupsGroupName)
	}
	if len(_resourcegroupsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _resourcegroupsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_resourcegroupsNextToken) > 0 {
		input.NextToken = aws.String(_resourcegroupsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListGroupResources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*resourcegroups.ListGroupResourcesOutput
	p := resourcegroups.NewListGroupResourcesPaginator(client, input)
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

// Returns the status of the last grouping or ungrouping action for each resource
// in the specified application group.
func resourcegroups_ListGroupingStatuses(cfg aws.Config, client *resourcegroups.Client) {
	input := &resourcegroups.ListGroupingStatusesInput{
		// Group: *string, // Required
	}

	if len(_resourcegroupsGroup) > 0 {
		input.Group = aws.String(_resourcegroupsGroup)
	}
	if len(_resourcegroupsFilters) > 0 {
		if err := assignInputField(input, "Filters", _resourcegroupsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_resourcegroupsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _resourcegroupsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_resourcegroupsNextToken) > 0 {
		input.NextToken = aws.String(_resourcegroupsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListGroupingStatuses(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*resourcegroups.ListGroupingStatusesOutput
	p := resourcegroups.NewListGroupingStatusesPaginator(client, input)
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

// Returns a list of existing Resource Groups in your account.
// # Minimum permissions
//
// To run this command, you must have the following permissions:
//
// - resource-groups:ListGroups
func resourcegroups_ListGroups(cfg aws.Config, client *resourcegroups.Client) {
	input := &resourcegroups.ListGroupsInput{}

	if len(_resourcegroupsFilters) > 0 {
		if err := assignInputField(input, "Filters", _resourcegroupsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_resourcegroupsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _resourcegroupsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_resourcegroupsNextToken) > 0 {
		input.NextToken = aws.String(_resourcegroupsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*resourcegroups.ListGroupsOutput
	p := resourcegroups.NewListGroupsPaginator(client, input)
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

// Returns a list of tag-sync tasks.
// # Minimum permissions
//
// To run this command, you must have the following permissions:
//
// - resource-groups:ListTagSyncTasks with the group passed in the filters as the
// resource or * if using no filters
func resourcegroups_ListTagSyncTasks(cfg aws.Config, client *resourcegroups.Client) {
	input := &resourcegroups.ListTagSyncTasksInput{}

	if len(_resourcegroupsFilters) > 0 {
		if err := assignInputField(input, "Filters", _resourcegroupsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_resourcegroupsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _resourcegroupsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_resourcegroupsNextToken) > 0 {
		input.NextToken = aws.String(_resourcegroupsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTagSyncTasks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*resourcegroups.ListTagSyncTasksOutput
	p := resourcegroups.NewListTagSyncTasksPaginator(client, input)
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

// Attaches a service configuration to the specified group. This occurs
// asynchronously, and can take time to complete. You can use GetGroupConfigurationto check the status
// of the update.
//
// # Minimum permissions
//
// To run this command, you must have the following permissions:
//
// - resource-groups:PutGroupConfiguration
func resourcegroups_PutGroupConfiguration(cfg aws.Config, client *resourcegroups.Client) {
	input := &resourcegroups.PutGroupConfigurationInput{}

	if len(_resourcegroupsConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _resourcegroupsConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_resourcegroupsGroup) > 0 {
		input.Group = aws.String(_resourcegroupsGroup)
	}

	if resp, err := client.PutGroupConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of Amazon Web Services resource identifiers that matches the
// specified query. The query uses the same format as a resource query in a CreateGroupor UpdateGroupQuery
// operation.
//
// # Minimum permissions
//
// To run this command, you must have the following permissions:
//
// - resource-groups:SearchResources
//
// - cloudformation:DescribeStacks
//
// - cloudformation:ListStackResources
//
// - tag:GetResources
func resourcegroups_SearchResources(cfg aws.Config, client *resourcegroups.Client) {
	input := &resourcegroups.SearchResourcesInput{
		// ResourceQuery: *types.ResourceQuery, // Required
	}

	if len(_resourcegroupsResourceQuery) > 0 {
		if err := assignInputField(input, "ResourceQuery", _resourcegroupsResourceQuery); err != nil {
			log.Errorf("invalid --resource-query: %s", err.Error())
			return
		}
	}
	if len(_resourcegroupsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _resourcegroupsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_resourcegroupsNextToken) > 0 {
		input.NextToken = aws.String(_resourcegroupsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.SearchResources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*resourcegroups.SearchResourcesOutput
	p := resourcegroups.NewSearchResourcesPaginator(client, input)
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

// Creates a new tag-sync task to onboard and sync resources tagged with a
// specific tag key-value pair to an application. To start a tag-sync task, you
// need a [resource tagging role]. The resource tagging role grants permissions to tag and untag
// applications resources and must include a trust policy that allows Resource
// Groups to assume the role and perform resource tagging tasks on your behalf.
//
// For instructions on creating a tag-sync task, see [Create a tag-sync using the Resource Groups API] in the Amazon Web Services
// Service Catalog AppRegistry Administrator Guide.
//
// # Minimum permissions
//
// To run this command, you must have the following permissions:
//
// - resource-groups:StartTagSyncTask on the application group
//
// - resource-groups:CreateGroup
//
// - iam:PassRole on the role provided in the request
//
// [resource tagging role]: https://docs.aws.amazon.com/servicecatalog/latest/arguide/app-tag-sync.html#tag-sync-role
// [Create a tag-sync using the Resource Groups API]: https://docs.aws.amazon.com/servicecatalog/latest/arguide/app-tag-sync.html#create-tag-sync
func resourcegroups_StartTagSyncTask(cfg aws.Config, client *resourcegroups.Client) {
	input := &resourcegroups.StartTagSyncTaskInput{
		// Group: *string, // Required
		// RoleArn: *string, // Required
	}

	if len(_resourcegroupsGroup) > 0 {
		input.Group = aws.String(_resourcegroupsGroup)
	}
	if len(_resourcegroupsRoleArn) > 0 {
		input.RoleArn = aws.String(_resourcegroupsRoleArn)
	}
	if len(_resourcegroupsResourceQuery) > 0 {
		if err := assignInputField(input, "ResourceQuery", _resourcegroupsResourceQuery); err != nil {
			log.Errorf("invalid --resource-query: %s", err.Error())
			return
		}
	}
	if len(_resourcegroupsTagKey) > 0 {
		input.TagKey = aws.String(_resourcegroupsTagKey)
	}
	if len(_resourcegroupsTagValue) > 0 {
		input.TagValue = aws.String(_resourcegroupsTagValue)
	}

	if resp, err := client.StartTagSyncTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds tags to a resource group with the specified Amazon resource name (ARN).
// Existing tags on a resource group are not changed if they are not specified in
// the request parameters.
//
// Do not store personally identifiable information (PII) or other confidential or
// sensitive information in tags. We use tags to provide you with billing and
// administration services. Tags are not intended to be used for private or
// sensitive data.
//
// # Minimum permissions
//
// To run this command, you must have the following permissions:
//
// - resource-groups:Tag
func resourcegroups_Tag(cfg aws.Config, client *resourcegroups.Client) {
	input := &resourcegroups.TagInput{
		// Arn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_resourcegroupsArn) > 0 {
		input.Arn = aws.String(_resourcegroupsArn)
	}
	if len(_resourcegroupsTags) > 0 {
		if err := assignInputField(input, "Tags", _resourcegroupsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.Tag(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified resources from the specified group. This operation works
// only with static groups that you populated using the GroupResourcesoperation. It doesn't work
// with any resource groups that are automatically populated by tag-based or
// CloudFormation stack-based queries.
//
// # Minimum permissions
//
// To run this command, you must have the following permissions:
//
// - resource-groups:UngroupResources
func resourcegroups_UngroupResources(cfg aws.Config, client *resourcegroups.Client) {
	input := &resourcegroups.UngroupResourcesInput{
		// Group: *string, // Required
		// ResourceArns: []string, // Required
	}

	if len(_resourcegroupsGroup) > 0 {
		input.Group = aws.String(_resourcegroupsGroup)
	}
	if len(_resourcegroupsResourceArns) > 0 {
		input.ResourceArns = append([]string(nil), _resourcegroupsResourceArns...)
	}

	if resp, err := client.UngroupResources(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes tags from a specified resource group.
// # Minimum permissions
//
// To run this command, you must have the following permissions:
//
// - resource-groups:Untag
func resourcegroups_Untag(cfg aws.Config, client *resourcegroups.Client) {
	input := &resourcegroups.UntagInput{
		// Arn: *string, // Required
		// Keys: []string, // Required
	}

	if len(_resourcegroupsArn) > 0 {
		input.Arn = aws.String(_resourcegroupsArn)
	}
	if len(_resourcegroupsKeys) > 0 {
		input.Keys = append([]string(nil), _resourcegroupsKeys...)
	}

	if resp, err := client.Untag(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Turns on or turns off optional features in Resource Groups.
// The preceding example shows that the request to turn on group lifecycle events
// is IN_PROGRESS . You can call the GetAccountSettings operation to check for completion by looking
// for GroupLifecycleEventsStatus to change to ACTIVE .
func resourcegroups_UpdateAccountSettings(cfg aws.Config, client *resourcegroups.Client) {
	input := &resourcegroups.UpdateAccountSettingsInput{}

	if len(_resourcegroupsGroupLifecycleEventsDesiredStatus) > 0 {
		if err := assignInputField(input, "GroupLifecycleEventsDesiredStatus", _resourcegroupsGroupLifecycleEventsDesiredStatus); err != nil {
			log.Errorf("invalid --group-lifecycle-events-desired-status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAccountSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the description for an existing group. You cannot update the name of a
// resource group.
//
// # Minimum permissions
//
// To run this command, you must have the following permissions:
//
// - resource-groups:UpdateGroup
func resourcegroups_UpdateGroup(cfg aws.Config, client *resourcegroups.Client) {
	input := &resourcegroups.UpdateGroupInput{}

	if len(_resourcegroupsCriticality) > 0 {
		if err := assignInputField(input, "Criticality", _resourcegroupsCriticality); err != nil {
			log.Errorf("invalid --criticality: %s", err.Error())
			return
		}
	}
	if len(_resourcegroupsDescription) > 0 {
		input.Description = aws.String(_resourcegroupsDescription)
	}
	if len(_resourcegroupsDisplayName) > 0 {
		input.DisplayName = aws.String(_resourcegroupsDisplayName)
	}
	if len(_resourcegroupsGroup) > 0 {
		input.Group = aws.String(_resourcegroupsGroup)
	}
	if len(_resourcegroupsGroupName) > 0 {
		input.GroupName = aws.String(_resourcegroupsGroupName)
	}
	if len(_resourcegroupsOwner) > 0 {
		input.Owner = aws.String(_resourcegroupsOwner)
	}

	if resp, err := client.UpdateGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the resource query of a group. For more information about resource
// queries, see [Create a tag-based group in Resource Groups].
//
// # Minimum permissions
//
// To run this command, you must have the following permissions:
//
// - resource-groups:UpdateGroupQuery
//
// [Create a tag-based group in Resource Groups]: https://docs.aws.amazon.com/ARG/latest/userguide/gettingstarted-query.html#gettingstarted-query-cli-tag
func resourcegroups_UpdateGroupQuery(cfg aws.Config, client *resourcegroups.Client) {
	input := &resourcegroups.UpdateGroupQueryInput{
		// ResourceQuery: *types.ResourceQuery, // Required
	}

	if len(_resourcegroupsResourceQuery) > 0 {
		if err := assignInputField(input, "ResourceQuery", _resourcegroupsResourceQuery); err != nil {
			log.Errorf("invalid --resource-query: %s", err.Error())
			return
		}
	}
	if len(_resourcegroupsGroup) > 0 {
		input.Group = aws.String(_resourcegroupsGroup)
	}
	if len(_resourcegroupsGroupName) > 0 {
		input.GroupName = aws.String(_resourcegroupsGroupName)
	}

	if resp, err := client.UpdateGroupQuery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_resourcegroupsCmd)
	_resourcegroupsCmd.Flags().SortFlags = false

	_resourcegroupsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_resourcegroupsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_resourcegroupsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_resourcegroupsCmd.Flags().StringVarP(&_resourcegroupsArn, "arn", "", "", "ARN")
	_resourcegroupsCmd.Flags().StringVarP(&_resourcegroupsConfiguration, "configuration", "", "", "Configuration")
	_resourcegroupsCmd.Flags().StringVarP(&_resourcegroupsCriticality, "criticality", "", "", "Criticality")
	_resourcegroupsCmd.Flags().StringVarP(&_resourcegroupsDescription, "description", "", "", "Description")
	_resourcegroupsCmd.Flags().StringVarP(&_resourcegroupsDisplayName, "display-name", "", "", "Display Name")
	_resourcegroupsCmd.Flags().StringVarP(&_resourcegroupsFilters, "filters", "", "", "Filters")
	_resourcegroupsCmd.Flags().StringVarP(&_resourcegroupsGroup, "group", "", "", "Group")
	_resourcegroupsCmd.Flags().StringVarP(&_resourcegroupsGroupLifecycleEventsDesiredStatus, "group-lifecycle-events-desired-status", "", "", "Group Lifecycle Events Desired Status")
	_resourcegroupsCmd.Flags().StringVarP(&_resourcegroupsGroupName, "group-name", "", "", "Group Name")
	_resourcegroupsCmd.Flags().StringSliceVarP(&_resourcegroupsKeys, "keys", "", nil, "Keys")
	_resourcegroupsCmd.Flags().StringVarP(&_resourcegroupsMaxResults, "max-results", "", "", "Max Results")
	_resourcegroupsCmd.Flags().StringVarP(&_resourcegroupsName, "name", "", "", "Name")
	_resourcegroupsCmd.Flags().StringVarP(&_resourcegroupsNextToken, "next-token", "", "", "Next Token")
	_resourcegroupsCmd.Flags().StringVarP(&_resourcegroupsOwner, "owner", "", "", "Owner")
	_resourcegroupsCmd.Flags().StringSliceVarP(&_resourcegroupsResourceArns, "resource-arns", "", nil, "Resource Arns")
	_resourcegroupsCmd.Flags().StringVarP(&_resourcegroupsResourceQuery, "resource-query", "", "", "Resource Query")
	_resourcegroupsCmd.Flags().StringVarP(&_resourcegroupsRoleArn, "role-arn", "", "", "Role ARN")
	_resourcegroupsCmd.Flags().StringVarP(&_resourcegroupsTagKey, "tag-key", "", "", "Tag Key")
	_resourcegroupsCmd.Flags().StringVarP(&_resourcegroupsTagValue, "tag-value", "", "", "Tag Value")
	_resourcegroupsCmd.Flags().StringVarP(&_resourcegroupsTags, "tags", "", "", "Tags")
	_resourcegroupsCmd.Flags().StringVarP(&_resourcegroupsTaskArn, "task-arn", "", "", "Task ARN")

	_resourcegroupsCmd.Flags().BoolVarP(&_resourcegroupsCancelTagSyncTask, "cancel-tag-sync-task", "", false, "Cancel Tag Sync Task")
	_resourcegroupsCmd.Flags().BoolVarP(&_resourcegroupsCreateGroup, "create-group", "", false, "Create Group")
	_resourcegroupsCmd.Flags().BoolVarP(&_resourcegroupsDeleteGroup, "delete-group", "", false, "Delete Group")
	_resourcegroupsCmd.Flags().BoolVarP(&_resourcegroupsGetAccountSettings, "get-account-settings", "", false, "Get Account Settings")
	_resourcegroupsCmd.Flags().BoolVarP(&_resourcegroupsGetGroup, "get-group", "", false, "Get Group")
	_resourcegroupsCmd.Flags().BoolVarP(&_resourcegroupsGetGroupConfiguration, "get-group-configuration", "", false, "Get Group Configuration")
	_resourcegroupsCmd.Flags().BoolVarP(&_resourcegroupsGetGroupQuery, "get-group-query", "", false, "Get Group Query")
	_resourcegroupsCmd.Flags().BoolVarP(&_resourcegroupsGetTagSyncTask, "get-tag-sync-task", "", false, "Get Tag Sync Task")
	_resourcegroupsCmd.Flags().BoolVarP(&_resourcegroupsGetTags, "get-tags", "", false, "Get Tags")
	_resourcegroupsCmd.Flags().BoolVarP(&_resourcegroupsGroupResources, "group-resources", "", false, "Group Resources")
	_resourcegroupsCmd.Flags().BoolVarP(&_resourcegroupsListGroupResources, "list-group-resources", "", false, "List Group Resources")
	_resourcegroupsCmd.Flags().BoolVarP(&_resourcegroupsListGroupingStatuses, "list-grouping-statuses", "", false, "List Grouping Statuses")
	_resourcegroupsCmd.Flags().BoolVarP(&_resourcegroupsListGroups, "list-groups", "", false, "List Groups")
	_resourcegroupsCmd.Flags().BoolVarP(&_resourcegroupsListTagSyncTasks, "list-tag-sync-tasks", "", false, "List Tag Sync Tasks")
	_resourcegroupsCmd.Flags().BoolVarP(&_resourcegroupsPutGroupConfiguration, "put-group-configuration", "", false, "Put Group Configuration")
	_resourcegroupsCmd.Flags().BoolVarP(&_resourcegroupsSearchResources, "search-resources", "", false, "Search Resources")
	_resourcegroupsCmd.Flags().BoolVarP(&_resourcegroupsStartTagSyncTask, "start-tag-sync-task", "", false, "Start Tag Sync Task")
	_resourcegroupsCmd.Flags().BoolVarP(&_resourcegroupsTag, "tag", "", false, "Tag")
	_resourcegroupsCmd.Flags().BoolVarP(&_resourcegroupsUngroupResources, "ungroup-resources", "", false, "Ungroup Resources")
	_resourcegroupsCmd.Flags().BoolVarP(&_resourcegroupsUntag, "untag", "", false, "Untag")
	_resourcegroupsCmd.Flags().BoolVarP(&_resourcegroupsUpdateAccountSettings, "update-account-settings", "", false, "Update Account Settings")
	_resourcegroupsCmd.Flags().BoolVarP(&_resourcegroupsUpdateGroup, "update-group", "", false, "Update Group")
	_resourcegroupsCmd.Flags().BoolVarP(&_resourcegroupsUpdateGroupQuery, "update-group-query", "", false, "Update Group Query")

}
