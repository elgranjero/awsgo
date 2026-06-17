package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/osis"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// osisCmd represents the osis command
var _osisCmd = &cobra.Command{
	Use:   "osis",
	Short: "AWS osis CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := osis.NewFromConfig(cfg)
		if _osisCreatePipeline {
			osis_CreatePipeline(cfg, client)
			return
		}
		if _osisCreatePipelineEndpoint {
			osis_CreatePipelineEndpoint(cfg, client)
			return
		}
		if _osisDeletePipeline {
			osis_DeletePipeline(cfg, client)
			return
		}
		if _osisDeletePipelineEndpoint {
			osis_DeletePipelineEndpoint(cfg, client)
			return
		}
		if _osisDeleteResourcePolicy {
			osis_DeleteResourcePolicy(cfg, client)
			return
		}
		if _osisGetPipeline {
			osis_GetPipeline(cfg, client)
			return
		}
		if _osisGetPipelineBlueprint {
			osis_GetPipelineBlueprint(cfg, client)
			return
		}
		if _osisGetPipelineChangeProgress {
			osis_GetPipelineChangeProgress(cfg, client)
			return
		}
		if _osisGetResourcePolicy {
			osis_GetResourcePolicy(cfg, client)
			return
		}
		if _osisListPipelineBlueprints {
			osis_ListPipelineBlueprints(cfg, client)
			return
		}
		if _osisListPipelineEndpointConnections {
			osis_ListPipelineEndpointConnections(cfg, client)
			return
		}
		if _osisListPipelineEndpoints {
			osis_ListPipelineEndpoints(cfg, client)
			return
		}
		if _osisListPipelines {
			osis_ListPipelines(cfg, client)
			return
		}
		if _osisListTagsForResource {
			osis_ListTagsForResource(cfg, client)
			return
		}
		if _osisPutResourcePolicy {
			osis_PutResourcePolicy(cfg, client)
			return
		}
		if _osisRevokePipelineEndpointConnections {
			osis_RevokePipelineEndpointConnections(cfg, client)
			return
		}
		if _osisStartPipeline {
			osis_StartPipeline(cfg, client)
			return
		}
		if _osisStopPipeline {
			osis_StopPipeline(cfg, client)
			return
		}
		if _osisTagResource {
			osis_TagResource(cfg, client)
			return
		}
		if _osisUntagResource {
			osis_UntagResource(cfg, client)
			return
		}
		if _osisUpdatePipeline {
			osis_UpdatePipeline(cfg, client)
			return
		}
		if _osisValidatePipeline {
			osis_ValidatePipeline(cfg, client)
			return
		}

	},
}

var (
	_osisCreatePipeline                    bool
	_osisCreatePipelineEndpoint            bool
	_osisDeletePipeline                    bool
	_osisDeletePipelineEndpoint            bool
	_osisDeleteResourcePolicy              bool
	_osisGetPipeline                       bool
	_osisGetPipelineBlueprint              bool
	_osisGetPipelineChangeProgress         bool
	_osisGetResourcePolicy                 bool
	_osisListPipelineBlueprints            bool
	_osisListPipelineEndpointConnections   bool
	_osisListPipelineEndpoints             bool
	_osisListPipelines                     bool
	_osisListTagsForResource               bool
	_osisPutResourcePolicy                 bool
	_osisRevokePipelineEndpointConnections bool
	_osisStartPipeline                     bool
	_osisStopPipeline                      bool
	_osisTagResource                       bool
	_osisUntagResource                     bool
	_osisUpdatePipeline                    bool
	_osisValidatePipeline                  bool

	_osisArn                       string
	_osisBlueprintName             string
	_osisBufferOptions             string
	_osisEncryptionAtRestOptions   string
	_osisEndpointId                string
	_osisEndpointIds               []string
	_osisFormat                    string
	_osisLogPublishingOptions      string
	_osisMaxResults                string
	_osisMaxUnits                  string
	_osisMinUnits                  string
	_osisNextToken                 string
	_osisPipelineArn               string
	_osisPipelineConfigurationBody string
	_osisPipelineName              string
	_osisPipelineRoleArn           string
	_osisPolicy                    string
	_osisResourceArn               string
	_osisTagKeys                   []string
	_osisTags                      string
	_osisVpcOptions                string
)

// Creates an OpenSearch Ingestion pipeline. For more information, see [Creating Amazon OpenSearch Ingestion pipelines].
//
// [Creating Amazon OpenSearch Ingestion pipelines]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/creating-pipeline.html
func osis_CreatePipeline(cfg aws.Config, client *osis.Client) {
	input := &osis.CreatePipelineInput{
		// MaxUnits: *int32, // Required
		// MinUnits: *int32, // Required
		// PipelineConfigurationBody: *string, // Required
		// PipelineName: *string, // Required
	}

	if len(_osisMaxUnits) > 0 {
		if err := assignInputField(input, "MaxUnits", _osisMaxUnits); err != nil {
			log.Errorf("invalid --max-units: %s", err.Error())
			return
		}
	}
	if len(_osisMinUnits) > 0 {
		if err := assignInputField(input, "MinUnits", _osisMinUnits); err != nil {
			log.Errorf("invalid --min-units: %s", err.Error())
			return
		}
	}
	if len(_osisPipelineConfigurationBody) > 0 {
		input.PipelineConfigurationBody = aws.String(_osisPipelineConfigurationBody)
	}
	if len(_osisPipelineName) > 0 {
		input.PipelineName = aws.String(_osisPipelineName)
	}
	if len(_osisBufferOptions) > 0 {
		if err := assignInputField(input, "BufferOptions", _osisBufferOptions); err != nil {
			log.Errorf("invalid --buffer-options: %s", err.Error())
			return
		}
	}
	if len(_osisEncryptionAtRestOptions) > 0 {
		if err := assignInputField(input, "EncryptionAtRestOptions", _osisEncryptionAtRestOptions); err != nil {
			log.Errorf("invalid --encryption-at-rest-options: %s", err.Error())
			return
		}
	}
	if len(_osisLogPublishingOptions) > 0 {
		if err := assignInputField(input, "LogPublishingOptions", _osisLogPublishingOptions); err != nil {
			log.Errorf("invalid --log-publishing-options: %s", err.Error())
			return
		}
	}
	if len(_osisPipelineRoleArn) > 0 {
		input.PipelineRoleArn = aws.String(_osisPipelineRoleArn)
	}
	if len(_osisTags) > 0 {
		if err := assignInputField(input, "Tags", _osisTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_osisVpcOptions) > 0 {
		if err := assignInputField(input, "VpcOptions", _osisVpcOptions); err != nil {
			log.Errorf("invalid --vpc-options: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePipeline(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a VPC endpoint for an OpenSearch Ingestion pipeline. Pipeline endpoints
// allow you to ingest data from your VPC into pipelines that you have access to.
func osis_CreatePipelineEndpoint(cfg aws.Config, client *osis.Client) {
	input := &osis.CreatePipelineEndpointInput{
		// PipelineArn: *string, // Required
		// VpcOptions: *types.PipelineEndpointVpcOptions, // Required
	}

	if len(_osisPipelineArn) > 0 {
		input.PipelineArn = aws.String(_osisPipelineArn)
	}
	if len(_osisVpcOptions) > 0 {
		if err := assignInputField(input, "VpcOptions", _osisVpcOptions); err != nil {
			log.Errorf("invalid --vpc-options: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePipelineEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an OpenSearch Ingestion pipeline. For more information, see [Deleting Amazon OpenSearch Ingestion pipelines].
//
// [Deleting Amazon OpenSearch Ingestion pipelines]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/delete-pipeline.html
func osis_DeletePipeline(cfg aws.Config, client *osis.Client) {
	input := &osis.DeletePipelineInput{
		// PipelineName: *string, // Required
	}

	if len(_osisPipelineName) > 0 {
		input.PipelineName = aws.String(_osisPipelineName)
	}

	if resp, err := client.DeletePipeline(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a VPC endpoint for an OpenSearch Ingestion pipeline.
func osis_DeletePipelineEndpoint(cfg aws.Config, client *osis.Client) {
	input := &osis.DeletePipelineEndpointInput{
		// EndpointId: *string, // Required
	}

	if len(_osisEndpointId) > 0 {
		input.EndpointId = aws.String(_osisEndpointId)
	}

	if resp, err := client.DeletePipelineEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a resource-based policy from an OpenSearch Ingestion resource.
func osis_DeleteResourcePolicy(cfg aws.Config, client *osis.Client) {
	input := &osis.DeleteResourcePolicyInput{
		// ResourceArn: *string, // Required
	}

	if len(_osisResourceArn) > 0 {
		input.ResourceArn = aws.String(_osisResourceArn)
	}

	if resp, err := client.DeleteResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about an OpenSearch Ingestion pipeline.
func osis_GetPipeline(cfg aws.Config, client *osis.Client) {
	input := &osis.GetPipelineInput{
		// PipelineName: *string, // Required
	}

	if len(_osisPipelineName) > 0 {
		input.PipelineName = aws.String(_osisPipelineName)
	}

	if resp, err := client.GetPipeline(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a specific blueprint for OpenSearch Ingestion.
// Blueprints are templates for the configuration needed for a CreatePipeline
// request. For more information, see [Using blueprints to create a pipeline].
//
// [Using blueprints to create a pipeline]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/creating-pipeline.html#pipeline-blueprint
func osis_GetPipelineBlueprint(cfg aws.Config, client *osis.Client) {
	input := &osis.GetPipelineBlueprintInput{
		// BlueprintName: *string, // Required
	}

	if len(_osisBlueprintName) > 0 {
		input.BlueprintName = aws.String(_osisBlueprintName)
	}
	if len(_osisFormat) > 0 {
		input.Format = aws.String(_osisFormat)
	}

	if resp, err := client.GetPipelineBlueprint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns progress information for the current change happening on an OpenSearch
// Ingestion pipeline. Currently, this operation only returns information when a
// pipeline is being created.
//
// For more information, see [Tracking the status of pipeline creation].
//
// [Tracking the status of pipeline creation]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/creating-pipeline.html#get-pipeline-progress
func osis_GetPipelineChangeProgress(cfg aws.Config, client *osis.Client) {
	input := &osis.GetPipelineChangeProgressInput{
		// PipelineName: *string, // Required
	}

	if len(_osisPipelineName) > 0 {
		input.PipelineName = aws.String(_osisPipelineName)
	}

	if resp, err := client.GetPipelineChangeProgress(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the resource-based policy attached to an OpenSearch Ingestion
// resource.
func osis_GetResourcePolicy(cfg aws.Config, client *osis.Client) {
	input := &osis.GetResourcePolicyInput{
		// ResourceArn: *string, // Required
	}

	if len(_osisResourceArn) > 0 {
		input.ResourceArn = aws.String(_osisResourceArn)
	}

	if resp, err := client.GetResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of all available blueprints for Data Prepper. For more
// information, see [Using blueprints to create a pipeline].
//
// [Using blueprints to create a pipeline]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/creating-pipeline.html#pipeline-blueprint
func osis_ListPipelineBlueprints(cfg aws.Config, client *osis.Client) {
	input := &osis.ListPipelineBlueprintsInput{}

	if resp, err := client.ListPipelineBlueprints(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the pipeline endpoints connected to pipelines in your account.
func osis_ListPipelineEndpointConnections(cfg aws.Config, client *osis.Client) {
	input := &osis.ListPipelineEndpointConnectionsInput{}

	if len(_osisMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _osisMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_osisNextToken) > 0 {
		input.NextToken = aws.String(_osisNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPipelineEndpointConnections(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*osis.ListPipelineEndpointConnectionsOutput
	p := osis.NewListPipelineEndpointConnectionsPaginator(client, input)
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

// Lists all pipeline endpoints in your account.
func osis_ListPipelineEndpoints(cfg aws.Config, client *osis.Client) {
	input := &osis.ListPipelineEndpointsInput{}

	if len(_osisMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _osisMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_osisNextToken) > 0 {
		input.NextToken = aws.String(_osisNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPipelineEndpoints(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*osis.ListPipelineEndpointsOutput
	p := osis.NewListPipelineEndpointsPaginator(client, input)
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

// Lists all OpenSearch Ingestion pipelines in the current Amazon Web Services
// account and Region. For more information, see [Viewing Amazon OpenSearch Ingestion pipelines].
//
// [Viewing Amazon OpenSearch Ingestion pipelines]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/list-pipeline.html
func osis_ListPipelines(cfg aws.Config, client *osis.Client) {
	input := &osis.ListPipelinesInput{}

	if len(_osisMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _osisMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_osisNextToken) > 0 {
		input.NextToken = aws.String(_osisNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPipelines(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*osis.ListPipelinesOutput
	p := osis.NewListPipelinesPaginator(client, input)
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

// Lists all resource tags associated with an OpenSearch Ingestion pipeline. For
// more information, see [Tagging Amazon OpenSearch Ingestion pipelines].
//
// [Tagging Amazon OpenSearch Ingestion pipelines]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/tag-pipeline.html
func osis_ListTagsForResource(cfg aws.Config, client *osis.Client) {
	input := &osis.ListTagsForResourceInput{
		// Arn: *string, // Required
	}

	if len(_osisArn) > 0 {
		input.Arn = aws.String(_osisArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attaches a resource-based policy to an OpenSearch Ingestion resource.
// Resource-based policies grant permissions to principals to perform actions on
// the resource.
func osis_PutResourcePolicy(cfg aws.Config, client *osis.Client) {
	input := &osis.PutResourcePolicyInput{
		// Policy: *string, // Required
		// ResourceArn: *string, // Required
	}

	if len(_osisPolicy) > 0 {
		input.Policy = aws.String(_osisPolicy)
	}
	if len(_osisResourceArn) > 0 {
		input.ResourceArn = aws.String(_osisResourceArn)
	}

	if resp, err := client.PutResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Revokes pipeline endpoints from specified endpoint IDs.
func osis_RevokePipelineEndpointConnections(cfg aws.Config, client *osis.Client) {
	input := &osis.RevokePipelineEndpointConnectionsInput{
		// EndpointIds: []string, // Required
		// PipelineArn: *string, // Required
	}

	if len(_osisEndpointIds) > 0 {
		input.EndpointIds = append([]string(nil), _osisEndpointIds...)
	}
	if len(_osisPipelineArn) > 0 {
		input.PipelineArn = aws.String(_osisPipelineArn)
	}

	if resp, err := client.RevokePipelineEndpointConnections(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts an OpenSearch Ingestion pipeline. For more information, see [Starting an OpenSearch Ingestion pipeline].
//
// [Starting an OpenSearch Ingestion pipeline]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/pipeline--stop-start.html#pipeline--start
func osis_StartPipeline(cfg aws.Config, client *osis.Client) {
	input := &osis.StartPipelineInput{
		// PipelineName: *string, // Required
	}

	if len(_osisPipelineName) > 0 {
		input.PipelineName = aws.String(_osisPipelineName)
	}

	if resp, err := client.StartPipeline(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops an OpenSearch Ingestion pipeline. For more information, see [Stopping an OpenSearch Ingestion pipeline].
//
// [Stopping an OpenSearch Ingestion pipeline]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/pipeline--stop-start.html#pipeline--stop
func osis_StopPipeline(cfg aws.Config, client *osis.Client) {
	input := &osis.StopPipelineInput{
		// PipelineName: *string, // Required
	}

	if len(_osisPipelineName) > 0 {
		input.PipelineName = aws.String(_osisPipelineName)
	}

	if resp, err := client.StopPipeline(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Tags an OpenSearch Ingestion pipeline. For more information, see [Tagging Amazon OpenSearch Ingestion pipelines].
//
// [Tagging Amazon OpenSearch Ingestion pipelines]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/tag-pipeline.html
func osis_TagResource(cfg aws.Config, client *osis.Client) {
	input := &osis.TagResourceInput{
		// Arn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_osisArn) > 0 {
		input.Arn = aws.String(_osisArn)
	}
	if len(_osisTags) > 0 {
		if err := assignInputField(input, "Tags", _osisTags); err != nil {
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

// Removes one or more tags from an OpenSearch Ingestion pipeline. For more
// information, see [Tagging Amazon OpenSearch Ingestion pipelines].
//
// [Tagging Amazon OpenSearch Ingestion pipelines]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/tag-pipeline.html
func osis_UntagResource(cfg aws.Config, client *osis.Client) {
	input := &osis.UntagResourceInput{
		// Arn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_osisArn) > 0 {
		input.Arn = aws.String(_osisArn)
	}
	if len(_osisTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _osisTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an OpenSearch Ingestion pipeline. For more information, see [Updating Amazon OpenSearch Ingestion pipelines].
//
// [Updating Amazon OpenSearch Ingestion pipelines]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/update-pipeline.html
func osis_UpdatePipeline(cfg aws.Config, client *osis.Client) {
	input := &osis.UpdatePipelineInput{
		// PipelineName: *string, // Required
	}

	if len(_osisPipelineName) > 0 {
		input.PipelineName = aws.String(_osisPipelineName)
	}
	if len(_osisBufferOptions) > 0 {
		if err := assignInputField(input, "BufferOptions", _osisBufferOptions); err != nil {
			log.Errorf("invalid --buffer-options: %s", err.Error())
			return
		}
	}
	if len(_osisEncryptionAtRestOptions) > 0 {
		if err := assignInputField(input, "EncryptionAtRestOptions", _osisEncryptionAtRestOptions); err != nil {
			log.Errorf("invalid --encryption-at-rest-options: %s", err.Error())
			return
		}
	}
	if len(_osisLogPublishingOptions) > 0 {
		if err := assignInputField(input, "LogPublishingOptions", _osisLogPublishingOptions); err != nil {
			log.Errorf("invalid --log-publishing-options: %s", err.Error())
			return
		}
	}
	if len(_osisMaxUnits) > 0 {
		if err := assignInputField(input, "MaxUnits", _osisMaxUnits); err != nil {
			log.Errorf("invalid --max-units: %s", err.Error())
			return
		}
	}
	if len(_osisMinUnits) > 0 {
		if err := assignInputField(input, "MinUnits", _osisMinUnits); err != nil {
			log.Errorf("invalid --min-units: %s", err.Error())
			return
		}
	}
	if len(_osisPipelineConfigurationBody) > 0 {
		input.PipelineConfigurationBody = aws.String(_osisPipelineConfigurationBody)
	}
	if len(_osisPipelineRoleArn) > 0 {
		input.PipelineRoleArn = aws.String(_osisPipelineRoleArn)
	}

	if resp, err := client.UpdatePipeline(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Checks whether an OpenSearch Ingestion pipeline configuration is valid prior to
// creation. For more information, see [Creating Amazon OpenSearch Ingestion pipelines].
//
// [Creating Amazon OpenSearch Ingestion pipelines]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/creating-pipeline.html
func osis_ValidatePipeline(cfg aws.Config, client *osis.Client) {
	input := &osis.ValidatePipelineInput{
		// PipelineConfigurationBody: *string, // Required
	}

	if len(_osisPipelineConfigurationBody) > 0 {
		input.PipelineConfigurationBody = aws.String(_osisPipelineConfigurationBody)
	}

	if resp, err := client.ValidatePipeline(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_osisCmd)
	_osisCmd.Flags().SortFlags = false

	_osisCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_osisCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_osisCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_osisCmd.Flags().StringVarP(&_osisArn, "arn", "", "", "ARN")
	_osisCmd.Flags().StringVarP(&_osisBlueprintName, "blueprint-name", "", "", "Blueprint Name")
	_osisCmd.Flags().StringVarP(&_osisBufferOptions, "buffer-options", "", "", "Buffer Options")
	_osisCmd.Flags().StringVarP(&_osisEncryptionAtRestOptions, "encryption-at-rest-options", "", "", "Encryption At Rest Options")
	_osisCmd.Flags().StringVarP(&_osisEndpointId, "endpoint-id", "", "", "Endpoint ID")
	_osisCmd.Flags().StringSliceVarP(&_osisEndpointIds, "endpoint-ids", "", nil, "Endpoint Ids")
	_osisCmd.Flags().StringVarP(&_osisFormat, "format", "", "", "Format")
	_osisCmd.Flags().StringVarP(&_osisLogPublishingOptions, "log-publishing-options", "", "", "Log Publishing Options")
	_osisCmd.Flags().StringVarP(&_osisMaxResults, "max-results", "", "", "Max Results")
	_osisCmd.Flags().StringVarP(&_osisMaxUnits, "max-units", "", "", "Max Units")
	_osisCmd.Flags().StringVarP(&_osisMinUnits, "min-units", "", "", "Min Units")
	_osisCmd.Flags().StringVarP(&_osisNextToken, "next-token", "", "", "Next Token")
	_osisCmd.Flags().StringVarP(&_osisPipelineArn, "pipeline-arn", "", "", "Pipeline ARN")
	_osisCmd.Flags().StringVarP(&_osisPipelineConfigurationBody, "pipeline-configuration-body", "", "", "Pipeline Configuration Body")
	_osisCmd.Flags().StringVarP(&_osisPipelineName, "pipeline-name", "", "", "Pipeline Name")
	_osisCmd.Flags().StringVarP(&_osisPipelineRoleArn, "pipeline-role-arn", "", "", "Pipeline Role ARN")
	_osisCmd.Flags().StringVarP(&_osisPolicy, "policy", "", "", "Policy")
	_osisCmd.Flags().StringVarP(&_osisResourceArn, "resource-arn", "", "", "Resource ARN")
	_osisCmd.Flags().StringSliceVarP(&_osisTagKeys, "tag-keys", "", nil, "Tag Keys")
	_osisCmd.Flags().StringVarP(&_osisTags, "tags", "", "", "Tags")
	_osisCmd.Flags().StringVarP(&_osisVpcOptions, "vpc-options", "", "", "VPC Options")

	_osisCmd.Flags().BoolVarP(&_osisCreatePipeline, "create-pipeline", "", false, "Create Pipeline")
	_osisCmd.Flags().BoolVarP(&_osisCreatePipelineEndpoint, "create-pipeline-endpoint", "", false, "Create Pipeline Endpoint")
	_osisCmd.Flags().BoolVarP(&_osisDeletePipeline, "delete-pipeline", "", false, "Delete Pipeline")
	_osisCmd.Flags().BoolVarP(&_osisDeletePipelineEndpoint, "delete-pipeline-endpoint", "", false, "Delete Pipeline Endpoint")
	_osisCmd.Flags().BoolVarP(&_osisDeleteResourcePolicy, "delete-resource-policy", "", false, "Delete Resource Policy")
	_osisCmd.Flags().BoolVarP(&_osisGetPipeline, "get-pipeline", "", false, "Get Pipeline")
	_osisCmd.Flags().BoolVarP(&_osisGetPipelineBlueprint, "get-pipeline-blueprint", "", false, "Get Pipeline Blueprint")
	_osisCmd.Flags().BoolVarP(&_osisGetPipelineChangeProgress, "get-pipeline-change-progress", "", false, "Get Pipeline Change Progress")
	_osisCmd.Flags().BoolVarP(&_osisGetResourcePolicy, "get-resource-policy", "", false, "Get Resource Policy")
	_osisCmd.Flags().BoolVarP(&_osisListPipelineBlueprints, "list-pipeline-blueprints", "", false, "List Pipeline Blueprints")
	_osisCmd.Flags().BoolVarP(&_osisListPipelineEndpointConnections, "list-pipeline-endpoint-connections", "", false, "List Pipeline Endpoint Connections")
	_osisCmd.Flags().BoolVarP(&_osisListPipelineEndpoints, "list-pipeline-endpoints", "", false, "List Pipeline Endpoints")
	_osisCmd.Flags().BoolVarP(&_osisListPipelines, "list-pipelines", "", false, "List Pipelines")
	_osisCmd.Flags().BoolVarP(&_osisListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_osisCmd.Flags().BoolVarP(&_osisPutResourcePolicy, "put-resource-policy", "", false, "Put Resource Policy")
	_osisCmd.Flags().BoolVarP(&_osisRevokePipelineEndpointConnections, "revoke-pipeline-endpoint-connections", "", false, "Revoke Pipeline Endpoint Connections")
	_osisCmd.Flags().BoolVarP(&_osisStartPipeline, "start-pipeline", "", false, "Start Pipeline")
	_osisCmd.Flags().BoolVarP(&_osisStopPipeline, "stop-pipeline", "", false, "Stop Pipeline")
	_osisCmd.Flags().BoolVarP(&_osisTagResource, "tag-resource", "", false, "Tag Resource")
	_osisCmd.Flags().BoolVarP(&_osisUntagResource, "untag-resource", "", false, "Untag Resource")
	_osisCmd.Flags().BoolVarP(&_osisUpdatePipeline, "update-pipeline", "", false, "Update Pipeline")
	_osisCmd.Flags().BoolVarP(&_osisValidatePipeline, "validate-pipeline", "", false, "Validate Pipeline")

}
