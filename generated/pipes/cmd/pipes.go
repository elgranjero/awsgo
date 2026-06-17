package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/pipes"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// pipesCmd represents the pipes command
var _pipesCmd = &cobra.Command{
	Use:   "pipes",
	Short: "AWS pipes CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := pipes.NewFromConfig(cfg)
		if _pipesCreatePipe {
			pipes_CreatePipe(cfg, client)
			return
		}
		if _pipesDeletePipe {
			pipes_DeletePipe(cfg, client)
			return
		}
		if _pipesDescribePipe {
			pipes_DescribePipe(cfg, client)
			return
		}
		if _pipesListPipes {
			pipes_ListPipes(cfg, client)
			return
		}
		if _pipesListTagsForResource {
			pipes_ListTagsForResource(cfg, client)
			return
		}
		if _pipesStartPipe {
			pipes_StartPipe(cfg, client)
			return
		}
		if _pipesStopPipe {
			pipes_StopPipe(cfg, client)
			return
		}
		if _pipesTagResource {
			pipes_TagResource(cfg, client)
			return
		}
		if _pipesUntagResource {
			pipes_UntagResource(cfg, client)
			return
		}
		if _pipesUpdatePipe {
			pipes_UpdatePipe(cfg, client)
			return
		}

	},
}

var (
	_pipesCreatePipe          bool
	_pipesDeletePipe          bool
	_pipesDescribePipe        bool
	_pipesListPipes           bool
	_pipesListTagsForResource bool
	_pipesStartPipe           bool
	_pipesStopPipe            bool
	_pipesTagResource         bool
	_pipesUntagResource       bool
	_pipesUpdatePipe          bool

	_pipesCurrentState         string
	_pipesDescription          string
	_pipesDesiredState         string
	_pipesEnrichment           string
	_pipesEnrichmentParameters string
	_pipesKmsKeyIdentifier     string
	_pipesLimit                string
	_pipesLogConfiguration     string
	_pipesName                 string
	_pipesNamePrefix           string
	_pipesNextToken            string
	_pipesResourceArn          string
	_pipesRoleArn              string
	_pipesSource               string
	_pipesSourceParameters     string
	_pipesSourcePrefix         string
	_pipesTagKeys              []string
	_pipesTags                 string
	_pipesTarget               string
	_pipesTargetParameters     string
	_pipesTargetPrefix         string
)

// Create a pipe. Amazon EventBridge Pipes connect event sources to targets and
// reduces the need for specialized knowledge and integration code.
func pipes_CreatePipe(cfg aws.Config, client *pipes.Client) {
	input := &pipes.CreatePipeInput{
		// Name: *string, // Required
		// RoleArn: *string, // Required
		// Source: *string, // Required
		// Target: *string, // Required
	}

	if len(_pipesName) > 0 {
		input.Name = aws.String(_pipesName)
	}
	if len(_pipesRoleArn) > 0 {
		input.RoleArn = aws.String(_pipesRoleArn)
	}
	if len(_pipesSource) > 0 {
		input.Source = aws.String(_pipesSource)
	}
	if len(_pipesTarget) > 0 {
		input.Target = aws.String(_pipesTarget)
	}
	if len(_pipesDescription) > 0 {
		input.Description = aws.String(_pipesDescription)
	}
	if len(_pipesDesiredState) > 0 {
		if err := assignInputField(input, "DesiredState", _pipesDesiredState); err != nil {
			log.Errorf("invalid --desired-state: %s", err.Error())
			return
		}
	}
	if len(_pipesEnrichment) > 0 {
		input.Enrichment = aws.String(_pipesEnrichment)
	}
	if len(_pipesEnrichmentParameters) > 0 {
		if err := assignInputField(input, "EnrichmentParameters", _pipesEnrichmentParameters); err != nil {
			log.Errorf("invalid --enrichment-parameters: %s", err.Error())
			return
		}
	}
	if len(_pipesKmsKeyIdentifier) > 0 {
		input.KmsKeyIdentifier = aws.String(_pipesKmsKeyIdentifier)
	}
	if len(_pipesLogConfiguration) > 0 {
		if err := assignInputField(input, "LogConfiguration", _pipesLogConfiguration); err != nil {
			log.Errorf("invalid --log-configuration: %s", err.Error())
			return
		}
	}
	if len(_pipesSourceParameters) > 0 {
		if err := assignInputField(input, "SourceParameters", _pipesSourceParameters); err != nil {
			log.Errorf("invalid --source-parameters: %s", err.Error())
			return
		}
	}
	if len(_pipesTags) > 0 {
		if err := assignInputField(input, "Tags", _pipesTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_pipesTargetParameters) > 0 {
		if err := assignInputField(input, "TargetParameters", _pipesTargetParameters); err != nil {
			log.Errorf("invalid --target-parameters: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePipe(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete an existing pipe. For more information about pipes, see [Amazon EventBridge Pipes] in the Amazon
// EventBridge User Guide.
//
// [Amazon EventBridge Pipes]: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes.html
func pipes_DeletePipe(cfg aws.Config, client *pipes.Client) {
	input := &pipes.DeletePipeInput{
		// Name: *string, // Required
	}

	if len(_pipesName) > 0 {
		input.Name = aws.String(_pipesName)
	}

	if resp, err := client.DeletePipe(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the information about an existing pipe. For more information about pipes,
// see [Amazon EventBridge Pipes]in the Amazon EventBridge User Guide.
//
// [Amazon EventBridge Pipes]: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes.html
func pipes_DescribePipe(cfg aws.Config, client *pipes.Client) {
	input := &pipes.DescribePipeInput{
		// Name: *string, // Required
	}

	if len(_pipesName) > 0 {
		input.Name = aws.String(_pipesName)
	}

	if resp, err := client.DescribePipe(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the pipes associated with this account. For more information about pipes,
// see [Amazon EventBridge Pipes]in the Amazon EventBridge User Guide.
//
// [Amazon EventBridge Pipes]: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes.html
func pipes_ListPipes(cfg aws.Config, client *pipes.Client) {
	input := &pipes.ListPipesInput{}

	if len(_pipesCurrentState) > 0 {
		if err := assignInputField(input, "CurrentState", _pipesCurrentState); err != nil {
			log.Errorf("invalid --current-state: %s", err.Error())
			return
		}
	}
	if len(_pipesDesiredState) > 0 {
		if err := assignInputField(input, "DesiredState", _pipesDesiredState); err != nil {
			log.Errorf("invalid --desired-state: %s", err.Error())
			return
		}
	}
	if len(_pipesLimit) > 0 {
		if err := assignInputField(input, "Limit", _pipesLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_pipesNamePrefix) > 0 {
		input.NamePrefix = aws.String(_pipesNamePrefix)
	}
	if len(_pipesNextToken) > 0 {
		input.NextToken = aws.String(_pipesNextToken)
	}
	if len(_pipesSourcePrefix) > 0 {
		input.SourcePrefix = aws.String(_pipesSourcePrefix)
	}
	if len(_pipesTargetPrefix) > 0 {
		input.TargetPrefix = aws.String(_pipesTargetPrefix)
	}

	if disablePaginator() {
		if resp, err := client.ListPipes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*pipes.ListPipesOutput
	p := pipes.NewListPipesPaginator(client, input)
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

// Displays the tags associated with a pipe.
func pipes_ListTagsForResource(cfg aws.Config, client *pipes.Client) {
	input := &pipes.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_pipesResourceArn) > 0 {
		input.ResourceArn = aws.String(_pipesResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Start an existing pipe.
func pipes_StartPipe(cfg aws.Config, client *pipes.Client) {
	input := &pipes.StartPipeInput{
		// Name: *string, // Required
	}

	if len(_pipesName) > 0 {
		input.Name = aws.String(_pipesName)
	}

	if resp, err := client.StartPipe(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stop an existing pipe.
func pipes_StopPipe(cfg aws.Config, client *pipes.Client) {
	input := &pipes.StopPipeInput{
		// Name: *string, // Required
	}

	if len(_pipesName) > 0 {
		input.Name = aws.String(_pipesName)
	}

	if resp, err := client.StopPipe(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assigns one or more tags (key-value pairs) to the specified pipe. Tags can help
// you organize and categorize your resources. You can also use them to scope user
// permissions by granting a user permission to access or change only resources
// with certain tag values.
//
// Tags don't have any semantic meaning to Amazon Web Services and are interpreted
// strictly as strings of characters.
//
// You can use the TagResource action with a pipe that already has tags. If you
// specify a new tag key, this tag is appended to the list of tags associated with
// the pipe. If you specify a tag key that is already associated with the pipe, the
// new tag value that you specify replaces the previous value for that tag.
//
// You can associate as many as 50 tags with a pipe.
func pipes_TagResource(cfg aws.Config, client *pipes.Client) {
	input := &pipes.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_pipesResourceArn) > 0 {
		input.ResourceArn = aws.String(_pipesResourceArn)
	}
	if len(_pipesTags) > 0 {
		if err := assignInputField(input, "Tags", _pipesTags); err != nil {
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

// Removes one or more tags from the specified pipes.
func pipes_UntagResource(cfg aws.Config, client *pipes.Client) {
	input := &pipes.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_pipesResourceArn) > 0 {
		input.ResourceArn = aws.String(_pipesResourceArn)
	}
	if len(_pipesTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _pipesTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update an existing pipe. When you call UpdatePipe , EventBridge only the updates
// fields you have specified in the request; the rest remain unchanged. The
// exception to this is if you modify any Amazon Web Services-service specific
// fields in the SourceParameters , EnrichmentParameters , or TargetParameters
// objects. For example, DynamoDBStreamParameters or EventBridgeEventBusParameters
// . EventBridge updates the fields in these objects atomically as one and
// overrides existing values. This is by design, and means that if you don't
// specify an optional field in one of these Parameters objects, EventBridge sets
// that field to its system-default value during the update.
//
// For more information about pipes, see [Amazon EventBridge Pipes] in the Amazon EventBridge User Guide.
//
// [Amazon EventBridge Pipes]: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes.html
func pipes_UpdatePipe(cfg aws.Config, client *pipes.Client) {
	input := &pipes.UpdatePipeInput{
		// Name: *string, // Required
		// RoleArn: *string, // Required
	}

	if len(_pipesName) > 0 {
		input.Name = aws.String(_pipesName)
	}
	if len(_pipesRoleArn) > 0 {
		input.RoleArn = aws.String(_pipesRoleArn)
	}
	if len(_pipesDescription) > 0 {
		input.Description = aws.String(_pipesDescription)
	}
	if len(_pipesDesiredState) > 0 {
		if err := assignInputField(input, "DesiredState", _pipesDesiredState); err != nil {
			log.Errorf("invalid --desired-state: %s", err.Error())
			return
		}
	}
	if len(_pipesEnrichment) > 0 {
		input.Enrichment = aws.String(_pipesEnrichment)
	}
	if len(_pipesEnrichmentParameters) > 0 {
		if err := assignInputField(input, "EnrichmentParameters", _pipesEnrichmentParameters); err != nil {
			log.Errorf("invalid --enrichment-parameters: %s", err.Error())
			return
		}
	}
	if len(_pipesKmsKeyIdentifier) > 0 {
		input.KmsKeyIdentifier = aws.String(_pipesKmsKeyIdentifier)
	}
	if len(_pipesLogConfiguration) > 0 {
		if err := assignInputField(input, "LogConfiguration", _pipesLogConfiguration); err != nil {
			log.Errorf("invalid --log-configuration: %s", err.Error())
			return
		}
	}
	if len(_pipesSourceParameters) > 0 {
		if err := assignInputField(input, "SourceParameters", _pipesSourceParameters); err != nil {
			log.Errorf("invalid --source-parameters: %s", err.Error())
			return
		}
	}
	if len(_pipesTarget) > 0 {
		input.Target = aws.String(_pipesTarget)
	}
	if len(_pipesTargetParameters) > 0 {
		if err := assignInputField(input, "TargetParameters", _pipesTargetParameters); err != nil {
			log.Errorf("invalid --target-parameters: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdatePipe(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_pipesCmd)
	_pipesCmd.Flags().SortFlags = false

	_pipesCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_pipesCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_pipesCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_pipesCmd.Flags().StringVarP(&_pipesCurrentState, "current-state", "", "", "Current State")
	_pipesCmd.Flags().StringVarP(&_pipesDescription, "description", "", "", "Description")
	_pipesCmd.Flags().StringVarP(&_pipesDesiredState, "desired-state", "", "", "Desired State")
	_pipesCmd.Flags().StringVarP(&_pipesEnrichment, "enrichment", "", "", "Enrichment")
	_pipesCmd.Flags().StringVarP(&_pipesEnrichmentParameters, "enrichment-parameters", "", "", "Enrichment Parameters")
	_pipesCmd.Flags().StringVarP(&_pipesKmsKeyIdentifier, "kms-key-identifier", "", "", "KMS Key Identifier")
	_pipesCmd.Flags().StringVarP(&_pipesLimit, "limit", "", "", "Limit")
	_pipesCmd.Flags().StringVarP(&_pipesLogConfiguration, "log-configuration", "", "", "Log Configuration")
	_pipesCmd.Flags().StringVarP(&_pipesName, "name", "", "", "Name")
	_pipesCmd.Flags().StringVarP(&_pipesNamePrefix, "name-prefix", "", "", "Name Prefix")
	_pipesCmd.Flags().StringVarP(&_pipesNextToken, "next-token", "", "", "Next Token")
	_pipesCmd.Flags().StringVarP(&_pipesResourceArn, "resource-arn", "", "", "Resource ARN")
	_pipesCmd.Flags().StringVarP(&_pipesRoleArn, "role-arn", "", "", "Role ARN")
	_pipesCmd.Flags().StringVarP(&_pipesSource, "source", "", "", "Source")
	_pipesCmd.Flags().StringVarP(&_pipesSourceParameters, "source-parameters", "", "", "Source Parameters")
	_pipesCmd.Flags().StringVarP(&_pipesSourcePrefix, "source-prefix", "", "", "Source Prefix")
	_pipesCmd.Flags().StringSliceVarP(&_pipesTagKeys, "tag-keys", "", nil, "Tag Keys")
	_pipesCmd.Flags().StringVarP(&_pipesTags, "tags", "", "", "Tags")
	_pipesCmd.Flags().StringVarP(&_pipesTarget, "target", "", "", "Target")
	_pipesCmd.Flags().StringVarP(&_pipesTargetParameters, "target-parameters", "", "", "Target Parameters")
	_pipesCmd.Flags().StringVarP(&_pipesTargetPrefix, "target-prefix", "", "", "Target Prefix")

	_pipesCmd.Flags().BoolVarP(&_pipesCreatePipe, "create-pipe", "", false, "Create Pipe")
	_pipesCmd.Flags().BoolVarP(&_pipesDeletePipe, "delete-pipe", "", false, "Delete Pipe")
	_pipesCmd.Flags().BoolVarP(&_pipesDescribePipe, "describe-pipe", "", false, "Describe Pipe")
	_pipesCmd.Flags().BoolVarP(&_pipesListPipes, "list-pipes", "", false, "List Pipes")
	_pipesCmd.Flags().BoolVarP(&_pipesListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_pipesCmd.Flags().BoolVarP(&_pipesStartPipe, "start-pipe", "", false, "Start Pipe")
	_pipesCmd.Flags().BoolVarP(&_pipesStopPipe, "stop-pipe", "", false, "Stop Pipe")
	_pipesCmd.Flags().BoolVarP(&_pipesTagResource, "tag-resource", "", false, "Tag Resource")
	_pipesCmd.Flags().BoolVarP(&_pipesUntagResource, "untag-resource", "", false, "Untag Resource")
	_pipesCmd.Flags().BoolVarP(&_pipesUpdatePipe, "update-pipe", "", false, "Update Pipe")

}
