package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/marketplacecatalog"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// marketplacecatalogCmd represents the marketplacecatalog command
var _marketplacecatalogCmd = &cobra.Command{
	Use:   "marketplacecatalog",
	Short: "AWS marketplacecatalog CLI",
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
		client := marketplacecatalog.NewFromConfig(cfg)
		if _marketplacecatalogBatchDescribeEntities {
			marketplacecatalog_BatchDescribeEntities(cfg, client)
			return
		}
		if _marketplacecatalogCancelChangeSet {
			marketplacecatalog_CancelChangeSet(cfg, client)
			return
		}
		if _marketplacecatalogDeleteResourcePolicy {
			marketplacecatalog_DeleteResourcePolicy(cfg, client)
			return
		}
		if _marketplacecatalogDescribeChangeSet {
			marketplacecatalog_DescribeChangeSet(cfg, client)
			return
		}
		if _marketplacecatalogDescribeEntity {
			marketplacecatalog_DescribeEntity(cfg, client)
			return
		}
		if _marketplacecatalogGetResourcePolicy {
			marketplacecatalog_GetResourcePolicy(cfg, client)
			return
		}
		if _marketplacecatalogListChangeSets {
			marketplacecatalog_ListChangeSets(cfg, client)
			return
		}
		if _marketplacecatalogListEntities {
			marketplacecatalog_ListEntities(cfg, client)
			return
		}
		if _marketplacecatalogListTagsForResource {
			marketplacecatalog_ListTagsForResource(cfg, client)
			return
		}
		if _marketplacecatalogPutResourcePolicy {
			marketplacecatalog_PutResourcePolicy(cfg, client)
			return
		}
		if _marketplacecatalogStartChangeSet {
			marketplacecatalog_StartChangeSet(cfg, client)
			return
		}
		if _marketplacecatalogTagResource {
			marketplacecatalog_TagResource(cfg, client)
			return
		}
		if _marketplacecatalogUntagResource {
			marketplacecatalog_UntagResource(cfg, client)
			return
		}

	},
}

var (
	_marketplacecatalogBatchDescribeEntities bool
	_marketplacecatalogCancelChangeSet       bool
	_marketplacecatalogDeleteResourcePolicy  bool
	_marketplacecatalogDescribeChangeSet     bool
	_marketplacecatalogDescribeEntity        bool
	_marketplacecatalogGetResourcePolicy     bool
	_marketplacecatalogListChangeSets        bool
	_marketplacecatalogListEntities          bool
	_marketplacecatalogListTagsForResource   bool
	_marketplacecatalogPutResourcePolicy     bool
	_marketplacecatalogStartChangeSet        bool
	_marketplacecatalogTagResource           bool
	_marketplacecatalogUntagResource         bool

	_marketplacecatalogCatalog            string
	_marketplacecatalogChangeSet          string
	_marketplacecatalogChangeSetId        string
	_marketplacecatalogChangeSetName      string
	_marketplacecatalogChangeSetTags      string
	_marketplacecatalogClientRequestToken string
	_marketplacecatalogEntityId           string
	_marketplacecatalogEntityRequestList  string
	_marketplacecatalogEntityType         string
	_marketplacecatalogEntityTypeFilters  string
	_marketplacecatalogEntityTypeSort     string
	_marketplacecatalogFilterList         string
	_marketplacecatalogIntent             string
	_marketplacecatalogMaxResults         string
	_marketplacecatalogNextToken          string
	_marketplacecatalogOwnershipType      string
	_marketplacecatalogPolicy             string
	_marketplacecatalogResourceArn        string
	_marketplacecatalogSort               string
	_marketplacecatalogTagKeys            []string
	_marketplacecatalogTags               string
)

// Returns metadata and content for multiple entities. This is the Batch version
// of the DescribeEntity API and uses the same IAM permission action as
// DescribeEntity API.
func marketplacecatalog_BatchDescribeEntities(cfg aws.Config, client *marketplacecatalog.Client) {
	input := &marketplacecatalog.BatchDescribeEntitiesInput{
		// EntityRequestList: []types.EntityRequest, // Required
	}

	if len(_marketplacecatalogEntityRequestList) > 0 {
		if err := assignInputField(input, "EntityRequestList", _marketplacecatalogEntityRequestList); err != nil {
			log.Errorf("invalid --entity-request-list: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchDescribeEntities(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Used to cancel an open change request. Must be sent before the status of the
// request changes to APPLYING , the final stage of completing your change request.
// You can describe a change during the 60-day request history retention period for
// API calls.
func marketplacecatalog_CancelChangeSet(cfg aws.Config, client *marketplacecatalog.Client) {
	input := &marketplacecatalog.CancelChangeSetInput{
		// Catalog: *string, // Required
		// ChangeSetId: *string, // Required
	}

	if len(_marketplacecatalogCatalog) > 0 {
		input.Catalog = aws.String(_marketplacecatalogCatalog)
	}
	if len(_marketplacecatalogChangeSetId) > 0 {
		input.ChangeSetId = aws.String(_marketplacecatalogChangeSetId)
	}

	if resp, err := client.CancelChangeSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a resource-based policy on an entity that is identified by its resource
// ARN.
func marketplacecatalog_DeleteResourcePolicy(cfg aws.Config, client *marketplacecatalog.Client) {
	input := &marketplacecatalog.DeleteResourcePolicyInput{
		// ResourceArn: *string, // Required
	}

	if len(_marketplacecatalogResourceArn) > 0 {
		input.ResourceArn = aws.String(_marketplacecatalogResourceArn)
	}

	if resp, err := client.DeleteResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides information about a given change set.
func marketplacecatalog_DescribeChangeSet(cfg aws.Config, client *marketplacecatalog.Client) {
	input := &marketplacecatalog.DescribeChangeSetInput{
		// Catalog: *string, // Required
		// ChangeSetId: *string, // Required
	}

	if len(_marketplacecatalogCatalog) > 0 {
		input.Catalog = aws.String(_marketplacecatalogCatalog)
	}
	if len(_marketplacecatalogChangeSetId) > 0 {
		input.ChangeSetId = aws.String(_marketplacecatalogChangeSetId)
	}

	if resp, err := client.DescribeChangeSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the metadata and content of the entity.
func marketplacecatalog_DescribeEntity(cfg aws.Config, client *marketplacecatalog.Client) {
	input := &marketplacecatalog.DescribeEntityInput{
		// Catalog: *string, // Required
		// EntityId: *string, // Required
	}

	if len(_marketplacecatalogCatalog) > 0 {
		input.Catalog = aws.String(_marketplacecatalogCatalog)
	}
	if len(_marketplacecatalogEntityId) > 0 {
		input.EntityId = aws.String(_marketplacecatalogEntityId)
	}

	if resp, err := client.DescribeEntity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a resource-based policy of an entity that is identified by its resource
// ARN.
func marketplacecatalog_GetResourcePolicy(cfg aws.Config, client *marketplacecatalog.Client) {
	input := &marketplacecatalog.GetResourcePolicyInput{
		// ResourceArn: *string, // Required
	}

	if len(_marketplacecatalogResourceArn) > 0 {
		input.ResourceArn = aws.String(_marketplacecatalogResourceArn)
	}

	if resp, err := client.GetResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the list of change sets owned by the account being used to make the
// call. You can filter this list by providing any combination of entityId ,
// ChangeSetName , and status. If you provide more than one filter, the API
// operation applies a logical AND between the filters.
//
// You can describe a change during the 60-day request history retention period
// for API calls.
func marketplacecatalog_ListChangeSets(cfg aws.Config, client *marketplacecatalog.Client) {
	input := &marketplacecatalog.ListChangeSetsInput{
		// Catalog: *string, // Required
	}

	if len(_marketplacecatalogCatalog) > 0 {
		input.Catalog = aws.String(_marketplacecatalogCatalog)
	}
	if len(_marketplacecatalogFilterList) > 0 {
		if err := assignInputField(input, "FilterList", _marketplacecatalogFilterList); err != nil {
			log.Errorf("invalid --filter-list: %s", err.Error())
			return
		}
	}
	if len(_marketplacecatalogMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _marketplacecatalogMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_marketplacecatalogNextToken) > 0 {
		input.NextToken = aws.String(_marketplacecatalogNextToken)
	}
	if len(_marketplacecatalogSort) > 0 {
		if err := assignInputField(input, "Sort", _marketplacecatalogSort); err != nil {
			log.Errorf("invalid --sort: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListChangeSets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*marketplacecatalog.ListChangeSetsOutput
	p := marketplacecatalog.NewListChangeSetsPaginator(client, input)
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

// Provides the list of entities of a given type.
func marketplacecatalog_ListEntities(cfg aws.Config, client *marketplacecatalog.Client) {
	input := &marketplacecatalog.ListEntitiesInput{
		// Catalog: *string, // Required
		// EntityType: *string, // Required
	}

	if len(_marketplacecatalogCatalog) > 0 {
		input.Catalog = aws.String(_marketplacecatalogCatalog)
	}
	if len(_marketplacecatalogEntityType) > 0 {
		input.EntityType = aws.String(_marketplacecatalogEntityType)
	}
	if len(_marketplacecatalogEntityTypeFilters) > 0 {
		if err := assignInputField(input, "EntityTypeFilters", _marketplacecatalogEntityTypeFilters); err != nil {
			log.Errorf("invalid --entity-type-filters: %s", err.Error())
			return
		}
	}
	if len(_marketplacecatalogEntityTypeSort) > 0 {
		if err := assignInputField(input, "EntityTypeSort", _marketplacecatalogEntityTypeSort); err != nil {
			log.Errorf("invalid --entity-type-sort: %s", err.Error())
			return
		}
	}
	if len(_marketplacecatalogFilterList) > 0 {
		if err := assignInputField(input, "FilterList", _marketplacecatalogFilterList); err != nil {
			log.Errorf("invalid --filter-list: %s", err.Error())
			return
		}
	}
	if len(_marketplacecatalogMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _marketplacecatalogMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_marketplacecatalogNextToken) > 0 {
		input.NextToken = aws.String(_marketplacecatalogNextToken)
	}
	if len(_marketplacecatalogOwnershipType) > 0 {
		if err := assignInputField(input, "OwnershipType", _marketplacecatalogOwnershipType); err != nil {
			log.Errorf("invalid --ownership-type: %s", err.Error())
			return
		}
	}
	if len(_marketplacecatalogSort) > 0 {
		if err := assignInputField(input, "Sort", _marketplacecatalogSort); err != nil {
			log.Errorf("invalid --sort: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListEntities(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*marketplacecatalog.ListEntitiesOutput
	p := marketplacecatalog.NewListEntitiesPaginator(client, input)
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

// Lists all tags that have been added to a resource (either an [entity] or [change set]).
//
// [change set]: https://docs.aws.amazon.com/marketplace-catalog/latest/api-reference/welcome.html#working-with-change-sets
// [entity]: https://docs.aws.amazon.com/marketplace-catalog/latest/api-reference/welcome.html#catalog-api-entities
func marketplacecatalog_ListTagsForResource(cfg aws.Config, client *marketplacecatalog.Client) {
	input := &marketplacecatalog.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_marketplacecatalogResourceArn) > 0 {
		input.ResourceArn = aws.String(_marketplacecatalogResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attaches a resource-based policy to an entity. Examples of an entity include:
// AmiProduct and ContainerProduct .
func marketplacecatalog_PutResourcePolicy(cfg aws.Config, client *marketplacecatalog.Client) {
	input := &marketplacecatalog.PutResourcePolicyInput{
		// Policy: *string, // Required
		// ResourceArn: *string, // Required
	}

	if len(_marketplacecatalogPolicy) > 0 {
		input.Policy = aws.String(_marketplacecatalogPolicy)
	}
	if len(_marketplacecatalogResourceArn) > 0 {
		input.ResourceArn = aws.String(_marketplacecatalogResourceArn)
	}

	if resp, err := client.PutResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Allows you to request changes for your entities. Within a single ChangeSet , you
// can't start the same change type against the same entity multiple times.
// Additionally, when a ChangeSet is running, all the entities targeted by the
// different changes are locked until the change set has completed (either
// succeeded, cancelled, or failed). If you try to start a change set containing a
// change against an entity that is already locked, you will receive a
// ResourceInUseException error.
//
// For example, you can't start the ChangeSet described in the [example] later in this
// topic because it contains two changes to run the same change type ( AddRevisions
// ) against the same entity ( entity-id(at)1 ).
//
// For more information about working with change sets, see [Working with change sets]. For information
// about change types for single-AMI products, see [Working with single-AMI products]. Also, for more information
// about change types available for container-based products, see [Working with container products].
//
// To download "DetailsDocument" shapes, see [Python] and [Java] shapes on GitHub.
//
// [Java]: https://github.com/awslabs/aws-marketplace-catalog-api-shapes-for-java/tree/main
// [Working with single-AMI products]: https://docs.aws.amazon.com/marketplace-catalog/latest/api-reference/ami-products.html#working-with-single-AMI-products
// [Working with change sets]: https://docs.aws.amazon.com/marketplace-catalog/latest/api-reference/welcome.html#working-with-change-sets
// [Working with container products]: https://docs.aws.amazon.com/marketplace-catalog/latest/api-reference/container-products.html#working-with-container-products
// [example]: https://docs.aws.amazon.com/marketplace-catalog/latest/api-reference/API_StartChangeSet.html#API_StartChangeSet_Examples
// [Python]: https://github.com/awslabs/aws-marketplace-catalog-api-shapes-for-python
func marketplacecatalog_StartChangeSet(cfg aws.Config, client *marketplacecatalog.Client) {
	input := &marketplacecatalog.StartChangeSetInput{
		// Catalog: *string, // Required
		// ChangeSet: []types.Change, // Required
	}

	if len(_marketplacecatalogCatalog) > 0 {
		input.Catalog = aws.String(_marketplacecatalogCatalog)
	}
	if len(_marketplacecatalogChangeSet) > 0 {
		if err := assignInputField(input, "ChangeSet", _marketplacecatalogChangeSet); err != nil {
			log.Errorf("invalid --change-set: %s", err.Error())
			return
		}
	}
	if len(_marketplacecatalogChangeSetName) > 0 {
		input.ChangeSetName = aws.String(_marketplacecatalogChangeSetName)
	}
	if len(_marketplacecatalogChangeSetTags) > 0 {
		if err := assignInputField(input, "ChangeSetTags", _marketplacecatalogChangeSetTags); err != nil {
			log.Errorf("invalid --change-set-tags: %s", err.Error())
			return
		}
	}
	if len(_marketplacecatalogClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_marketplacecatalogClientRequestToken)
	}
	if len(_marketplacecatalogIntent) > 0 {
		if err := assignInputField(input, "Intent", _marketplacecatalogIntent); err != nil {
			log.Errorf("invalid --intent: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartChangeSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Tags a resource (either an [entity] or [change set]).
//
// [change set]: https://docs.aws.amazon.com/marketplace-catalog/latest/api-reference/welcome.html#working-with-change-sets
// [entity]: https://docs.aws.amazon.com/marketplace-catalog/latest/api-reference/welcome.html#catalog-api-entities
func marketplacecatalog_TagResource(cfg aws.Config, client *marketplacecatalog.Client) {
	input := &marketplacecatalog.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_marketplacecatalogResourceArn) > 0 {
		input.ResourceArn = aws.String(_marketplacecatalogResourceArn)
	}
	if len(_marketplacecatalogTags) > 0 {
		if err := assignInputField(input, "Tags", _marketplacecatalogTags); err != nil {
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

// Removes a tag or list of tags from a resource (either an [entity] or [change set]).
//
// [change set]: https://docs.aws.amazon.com/marketplace-catalog/latest/api-reference/welcome.html#working-with-change-sets
// [entity]: https://docs.aws.amazon.com/marketplace-catalog/latest/api-reference/welcome.html#catalog-api-entities
func marketplacecatalog_UntagResource(cfg aws.Config, client *marketplacecatalog.Client) {
	input := &marketplacecatalog.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_marketplacecatalogResourceArn) > 0 {
		input.ResourceArn = aws.String(_marketplacecatalogResourceArn)
	}
	if len(_marketplacecatalogTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _marketplacecatalogTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_marketplacecatalogCmd)
	_marketplacecatalogCmd.Flags().SortFlags = false

	_marketplacecatalogCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_marketplacecatalogCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_marketplacecatalogCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_marketplacecatalogCmd.Flags().StringVarP(&_marketplacecatalogCatalog, "catalog", "", "", "Catalog")
	_marketplacecatalogCmd.Flags().StringVarP(&_marketplacecatalogChangeSet, "change-set", "", "", "Change Set")
	_marketplacecatalogCmd.Flags().StringVarP(&_marketplacecatalogChangeSetId, "change-set-id", "", "", "Change Set ID")
	_marketplacecatalogCmd.Flags().StringVarP(&_marketplacecatalogChangeSetName, "change-set-name", "", "", "Change Set Name")
	_marketplacecatalogCmd.Flags().StringVarP(&_marketplacecatalogChangeSetTags, "change-set-tags", "", "", "Change Set Tags")
	_marketplacecatalogCmd.Flags().StringVarP(&_marketplacecatalogClientRequestToken, "client-request-token", "", "", "Client Request Token")
	_marketplacecatalogCmd.Flags().StringVarP(&_marketplacecatalogEntityId, "entity-id", "", "", "Entity ID")
	_marketplacecatalogCmd.Flags().StringVarP(&_marketplacecatalogEntityRequestList, "entity-request-list", "", "", "Entity Request List")
	_marketplacecatalogCmd.Flags().StringVarP(&_marketplacecatalogEntityType, "entity-type", "", "", "Entity Type")
	_marketplacecatalogCmd.Flags().StringVarP(&_marketplacecatalogEntityTypeFilters, "entity-type-filters", "", "", "Entity Type Filters")
	_marketplacecatalogCmd.Flags().StringVarP(&_marketplacecatalogEntityTypeSort, "entity-type-sort", "", "", "Entity Type Sort")
	_marketplacecatalogCmd.Flags().StringVarP(&_marketplacecatalogFilterList, "filter-list", "", "", "Filter List")
	_marketplacecatalogCmd.Flags().StringVarP(&_marketplacecatalogIntent, "intent", "", "", "Intent")
	_marketplacecatalogCmd.Flags().StringVarP(&_marketplacecatalogMaxResults, "max-results", "", "", "Max Results")
	_marketplacecatalogCmd.Flags().StringVarP(&_marketplacecatalogNextToken, "next-token", "", "", "Next Token")
	_marketplacecatalogCmd.Flags().StringVarP(&_marketplacecatalogOwnershipType, "ownership-type", "", "", "Ownership Type")
	_marketplacecatalogCmd.Flags().StringVarP(&_marketplacecatalogPolicy, "policy", "", "", "Policy")
	_marketplacecatalogCmd.Flags().StringVarP(&_marketplacecatalogResourceArn, "resource-arn", "", "", "Resource ARN")
	_marketplacecatalogCmd.Flags().StringVarP(&_marketplacecatalogSort, "sort", "", "", "Sort")
	_marketplacecatalogCmd.Flags().StringSliceVarP(&_marketplacecatalogTagKeys, "tag-keys", "", nil, "Tag Keys")
	_marketplacecatalogCmd.Flags().StringVarP(&_marketplacecatalogTags, "tags", "", "", "Tags")

	_marketplacecatalogCmd.Flags().BoolVarP(&_marketplacecatalogBatchDescribeEntities, "batch-describe-entities", "", false, "Batch Describe Entities")
	_marketplacecatalogCmd.Flags().BoolVarP(&_marketplacecatalogCancelChangeSet, "cancel-change-set", "", false, "Cancel Change Set")
	_marketplacecatalogCmd.Flags().BoolVarP(&_marketplacecatalogDeleteResourcePolicy, "delete-resource-policy", "", false, "Delete Resource Policy")
	_marketplacecatalogCmd.Flags().BoolVarP(&_marketplacecatalogDescribeChangeSet, "describe-change-set", "", false, "Describe Change Set")
	_marketplacecatalogCmd.Flags().BoolVarP(&_marketplacecatalogDescribeEntity, "describe-entity", "", false, "Describe Entity")
	_marketplacecatalogCmd.Flags().BoolVarP(&_marketplacecatalogGetResourcePolicy, "get-resource-policy", "", false, "Get Resource Policy")
	_marketplacecatalogCmd.Flags().BoolVarP(&_marketplacecatalogListChangeSets, "list-change-sets", "", false, "List Change Sets")
	_marketplacecatalogCmd.Flags().BoolVarP(&_marketplacecatalogListEntities, "list-entities", "", false, "List Entities")
	_marketplacecatalogCmd.Flags().BoolVarP(&_marketplacecatalogListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_marketplacecatalogCmd.Flags().BoolVarP(&_marketplacecatalogPutResourcePolicy, "put-resource-policy", "", false, "Put Resource Policy")
	_marketplacecatalogCmd.Flags().BoolVarP(&_marketplacecatalogStartChangeSet, "start-change-set", "", false, "Start Change Set")
	_marketplacecatalogCmd.Flags().BoolVarP(&_marketplacecatalogTagResource, "tag-resource", "", false, "Tag Resource")
	_marketplacecatalogCmd.Flags().BoolVarP(&_marketplacecatalogUntagResource, "untag-resource", "", false, "Untag Resource")

}
