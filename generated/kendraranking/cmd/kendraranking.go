package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kendraranking"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// kendrarankingCmd represents the kendraranking command
var _kendrarankingCmd = &cobra.Command{
	Use:   "kendraranking",
	Short: "AWS kendraranking CLI",
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
		client := kendraranking.NewFromConfig(cfg)
		if _kendrarankingCreateRescoreExecutionPlan {
			kendraranking_CreateRescoreExecutionPlan(cfg, client)
			return
		}
		if _kendrarankingDeleteRescoreExecutionPlan {
			kendraranking_DeleteRescoreExecutionPlan(cfg, client)
			return
		}
		if _kendrarankingDescribeRescoreExecutionPlan {
			kendraranking_DescribeRescoreExecutionPlan(cfg, client)
			return
		}
		if _kendrarankingListRescoreExecutionPlans {
			kendraranking_ListRescoreExecutionPlans(cfg, client)
			return
		}
		if _kendrarankingListTagsForResource {
			kendraranking_ListTagsForResource(cfg, client)
			return
		}
		if _kendrarankingRescore {
			kendraranking_Rescore(cfg, client)
			return
		}
		if _kendrarankingTagResource {
			kendraranking_TagResource(cfg, client)
			return
		}
		if _kendrarankingUntagResource {
			kendraranking_UntagResource(cfg, client)
			return
		}
		if _kendrarankingUpdateRescoreExecutionPlan {
			kendraranking_UpdateRescoreExecutionPlan(cfg, client)
			return
		}

	},
}

var (
	_kendrarankingCreateRescoreExecutionPlan   bool
	_kendrarankingDeleteRescoreExecutionPlan   bool
	_kendrarankingDescribeRescoreExecutionPlan bool
	_kendrarankingListRescoreExecutionPlans    bool
	_kendrarankingListTagsForResource          bool
	_kendrarankingRescore                      bool
	_kendrarankingTagResource                  bool
	_kendrarankingUntagResource                bool
	_kendrarankingUpdateRescoreExecutionPlan   bool

	_kendrarankingCapacityUnits          string
	_kendrarankingClientToken            string
	_kendrarankingDescription            string
	_kendrarankingDocuments              string
	_kendrarankingId                     string
	_kendrarankingMaxResults             string
	_kendrarankingName                   string
	_kendrarankingNextToken              string
	_kendrarankingRescoreExecutionPlanId string
	_kendrarankingResourceARN            string
	_kendrarankingSearchQuery            string
	_kendrarankingTagKeys                []string
	_kendrarankingTags                   string
)

// Creates a rescore execution plan. A rescore execution plan is an Amazon Kendra
// Intelligent Ranking resource used for provisioning the Rescore API. You set the
// number of capacity units that you require for Amazon Kendra Intelligent Ranking
// to rescore or re-rank a search service's results.
//
// For an example of using the CreateRescoreExecutionPlan API, including using the
// Python and Java SDKs, see [Semantically ranking a search service's results].
//
// [Semantically ranking a search service's results]: https://docs.aws.amazon.com/kendra/latest/dg/search-service-rerank.html
func kendraranking_CreateRescoreExecutionPlan(cfg aws.Config, client *kendraranking.Client) {
	input := &kendraranking.CreateRescoreExecutionPlanInput{
		// Name: *string, // Required
	}

	if len(_kendrarankingName) > 0 {
		input.Name = aws.String(_kendrarankingName)
	}
	if len(_kendrarankingCapacityUnits) > 0 {
		if err := assignInputField(input, "CapacityUnits", _kendrarankingCapacityUnits); err != nil {
			log.Errorf("invalid --capacity-units: %s", err.Error())
			return
		}
	}
	if len(_kendrarankingClientToken) > 0 {
		input.ClientToken = aws.String(_kendrarankingClientToken)
	}
	if len(_kendrarankingDescription) > 0 {
		input.Description = aws.String(_kendrarankingDescription)
	}
	if len(_kendrarankingTags) > 0 {
		if err := assignInputField(input, "Tags", _kendrarankingTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRescoreExecutionPlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a rescore execution plan. A rescore execution plan is an Amazon Kendra
// Intelligent Ranking resource used for provisioning the Rescore API.
func kendraranking_DeleteRescoreExecutionPlan(cfg aws.Config, client *kendraranking.Client) {
	input := &kendraranking.DeleteRescoreExecutionPlanInput{
		// Id: *string, // Required
	}

	if len(_kendrarankingId) > 0 {
		input.Id = aws.String(_kendrarankingId)
	}

	if resp, err := client.DeleteRescoreExecutionPlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a rescore execution plan. A rescore execution plan is an
// Amazon Kendra Intelligent Ranking resource used for provisioning the Rescore
// API.
func kendraranking_DescribeRescoreExecutionPlan(cfg aws.Config, client *kendraranking.Client) {
	input := &kendraranking.DescribeRescoreExecutionPlanInput{
		// Id: *string, // Required
	}

	if len(_kendrarankingId) > 0 {
		input.Id = aws.String(_kendrarankingId)
	}

	if resp, err := client.DescribeRescoreExecutionPlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists your rescore execution plans. A rescore execution plan is an Amazon
// Kendra Intelligent Ranking resource used for provisioning the Rescore API.
func kendraranking_ListRescoreExecutionPlans(cfg aws.Config, client *kendraranking.Client) {
	input := &kendraranking.ListRescoreExecutionPlansInput{}

	if len(_kendrarankingMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _kendrarankingMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_kendrarankingNextToken) > 0 {
		input.NextToken = aws.String(_kendrarankingNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRescoreExecutionPlans(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*kendraranking.ListRescoreExecutionPlansOutput
	p := kendraranking.NewListRescoreExecutionPlansPaginator(client, input)
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

// Gets a list of tags associated with a specified resource. A rescore execution
// plan is an example of a resource that can have tags associated with it.
func kendraranking_ListTagsForResource(cfg aws.Config, client *kendraranking.Client) {
	input := &kendraranking.ListTagsForResourceInput{
		// ResourceARN: *string, // Required
	}

	if len(_kendrarankingResourceARN) > 0 {
		input.ResourceARN = aws.String(_kendrarankingResourceARN)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Rescores or re-ranks search results from a search service such as OpenSearch
// (self managed). You use the semantic search capabilities of Amazon Kendra
// Intelligent Ranking to improve the search service's results.
func kendraranking_Rescore(cfg aws.Config, client *kendraranking.Client) {
	input := &kendraranking.RescoreInput{
		// Documents: []types.Document, // Required
		// RescoreExecutionPlanId: *string, // Required
		// SearchQuery: *string, // Required
	}

	if len(_kendrarankingDocuments) > 0 {
		if err := assignInputField(input, "Documents", _kendrarankingDocuments); err != nil {
			log.Errorf("invalid --documents: %s", err.Error())
			return
		}
	}
	if len(_kendrarankingRescoreExecutionPlanId) > 0 {
		input.RescoreExecutionPlanId = aws.String(_kendrarankingRescoreExecutionPlanId)
	}
	if len(_kendrarankingSearchQuery) > 0 {
		input.SearchQuery = aws.String(_kendrarankingSearchQuery)
	}

	if resp, err := client.Rescore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a specified tag to a specified rescore execution plan. A rescore execution
// plan is an Amazon Kendra Intelligent Ranking resource used for provisioning the
// Rescore API. If the tag already exists, the existing value is replaced with the
// new value.
func kendraranking_TagResource(cfg aws.Config, client *kendraranking.Client) {
	input := &kendraranking.TagResourceInput{
		// ResourceARN: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_kendrarankingResourceARN) > 0 {
		input.ResourceARN = aws.String(_kendrarankingResourceARN)
	}
	if len(_kendrarankingTags) > 0 {
		if err := assignInputField(input, "Tags", _kendrarankingTags); err != nil {
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

// Removes a tag from a rescore execution plan. A rescore execution plan is an
// Amazon Kendra Intelligent Ranking resource used for provisioning the Rescore
// operation.
func kendraranking_UntagResource(cfg aws.Config, client *kendraranking.Client) {
	input := &kendraranking.UntagResourceInput{
		// ResourceARN: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_kendrarankingResourceARN) > 0 {
		input.ResourceARN = aws.String(_kendrarankingResourceARN)
	}
	if len(_kendrarankingTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _kendrarankingTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a rescore execution plan. A rescore execution plan is an Amazon Kendra
// Intelligent Ranking resource used for provisioning the Rescore API. You can
// update the number of capacity units you require for Amazon Kendra Intelligent
// Ranking to rescore or re-rank a search service's results.
func kendraranking_UpdateRescoreExecutionPlan(cfg aws.Config, client *kendraranking.Client) {
	input := &kendraranking.UpdateRescoreExecutionPlanInput{
		// Id: *string, // Required
	}

	if len(_kendrarankingId) > 0 {
		input.Id = aws.String(_kendrarankingId)
	}
	if len(_kendrarankingCapacityUnits) > 0 {
		if err := assignInputField(input, "CapacityUnits", _kendrarankingCapacityUnits); err != nil {
			log.Errorf("invalid --capacity-units: %s", err.Error())
			return
		}
	}
	if len(_kendrarankingDescription) > 0 {
		input.Description = aws.String(_kendrarankingDescription)
	}
	if len(_kendrarankingName) > 0 {
		input.Name = aws.String(_kendrarankingName)
	}

	if resp, err := client.UpdateRescoreExecutionPlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_kendrarankingCmd)
	_kendrarankingCmd.Flags().SortFlags = false

	_kendrarankingCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_kendrarankingCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_kendrarankingCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_kendrarankingCmd.Flags().StringVarP(&_kendrarankingCapacityUnits, "capacity-units", "", "", "Capacity Units")
	_kendrarankingCmd.Flags().StringVarP(&_kendrarankingClientToken, "client-token", "", "", "Client Token")
	_kendrarankingCmd.Flags().StringVarP(&_kendrarankingDescription, "description", "", "", "Description")
	_kendrarankingCmd.Flags().StringVarP(&_kendrarankingDocuments, "documents", "", "", "Documents")
	_kendrarankingCmd.Flags().StringVarP(&_kendrarankingId, "id", "", "", "ID")
	_kendrarankingCmd.Flags().StringVarP(&_kendrarankingMaxResults, "max-results", "", "", "Max Results")
	_kendrarankingCmd.Flags().StringVarP(&_kendrarankingName, "name", "", "", "Name")
	_kendrarankingCmd.Flags().StringVarP(&_kendrarankingNextToken, "next-token", "", "", "Next Token")
	_kendrarankingCmd.Flags().StringVarP(&_kendrarankingRescoreExecutionPlanId, "rescore-execution-plan-id", "", "", "Rescore Execution Plan ID")
	_kendrarankingCmd.Flags().StringVarP(&_kendrarankingResourceARN, "resource-arn", "", "", "Resource ARN")
	_kendrarankingCmd.Flags().StringVarP(&_kendrarankingSearchQuery, "search-query", "", "", "Search Query")
	_kendrarankingCmd.Flags().StringSliceVarP(&_kendrarankingTagKeys, "tag-keys", "", nil, "Tag Keys")
	_kendrarankingCmd.Flags().StringVarP(&_kendrarankingTags, "tags", "", "", "Tags")

	_kendrarankingCmd.Flags().BoolVarP(&_kendrarankingCreateRescoreExecutionPlan, "create-rescore-execution-plan", "", false, "Create Rescore Execution Plan")
	_kendrarankingCmd.Flags().BoolVarP(&_kendrarankingDeleteRescoreExecutionPlan, "delete-rescore-execution-plan", "", false, "Delete Rescore Execution Plan")
	_kendrarankingCmd.Flags().BoolVarP(&_kendrarankingDescribeRescoreExecutionPlan, "describe-rescore-execution-plan", "", false, "Describe Rescore Execution Plan")
	_kendrarankingCmd.Flags().BoolVarP(&_kendrarankingListRescoreExecutionPlans, "list-rescore-execution-plans", "", false, "List Rescore Execution Plans")
	_kendrarankingCmd.Flags().BoolVarP(&_kendrarankingListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_kendrarankingCmd.Flags().BoolVarP(&_kendrarankingRescore, "rescore", "", false, "Rescore")
	_kendrarankingCmd.Flags().BoolVarP(&_kendrarankingTagResource, "tag-resource", "", false, "Tag Resource")
	_kendrarankingCmd.Flags().BoolVarP(&_kendrarankingUntagResource, "untag-resource", "", false, "Untag Resource")
	_kendrarankingCmd.Flags().BoolVarP(&_kendrarankingUpdateRescoreExecutionPlan, "update-rescore-execution-plan", "", false, "Update Rescore Execution Plan")

}
