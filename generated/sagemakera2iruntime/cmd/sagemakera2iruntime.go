package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sagemakera2iruntime"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// sagemakera2iruntimeCmd represents the sagemakera2iruntime command
var _sagemakera2iruntimeCmd = &cobra.Command{
	Use:   "sagemakera2iruntime",
	Short: "AWS sagemakera2iruntime CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := sagemakera2iruntime.NewFromConfig(cfg)
		if _sagemakera2iruntimeDeleteHumanLoop {
			sagemakera2iruntime_DeleteHumanLoop(cfg, client)
			return
		}
		if _sagemakera2iruntimeDescribeHumanLoop {
			sagemakera2iruntime_DescribeHumanLoop(cfg, client)
			return
		}
		if _sagemakera2iruntimeListHumanLoops {
			sagemakera2iruntime_ListHumanLoops(cfg, client)
			return
		}
		if _sagemakera2iruntimeStartHumanLoop {
			sagemakera2iruntime_StartHumanLoop(cfg, client)
			return
		}
		if _sagemakera2iruntimeStopHumanLoop {
			sagemakera2iruntime_StopHumanLoop(cfg, client)
			return
		}

	},
}

var (
	_sagemakera2iruntimeDeleteHumanLoop   bool
	_sagemakera2iruntimeDescribeHumanLoop bool
	_sagemakera2iruntimeListHumanLoops    bool
	_sagemakera2iruntimeStartHumanLoop    bool
	_sagemakera2iruntimeStopHumanLoop     bool

	_sagemakera2iruntimeCreationTimeAfter  string
	_sagemakera2iruntimeCreationTimeBefore string
	_sagemakera2iruntimeDataAttributes     string
	_sagemakera2iruntimeFlowDefinitionArn  string
	_sagemakera2iruntimeHumanLoopInput     string
	_sagemakera2iruntimeHumanLoopName      string
	_sagemakera2iruntimeMaxResults         string
	_sagemakera2iruntimeNextToken          string
	_sagemakera2iruntimeSortOrder          string
)

// Deletes the specified human loop for a flow definition.
// If the human loop was deleted, this operation will return a
// ResourceNotFoundException .
func sagemakera2iruntime_DeleteHumanLoop(cfg aws.Config, client *sagemakera2iruntime.Client) {
	input := &sagemakera2iruntime.DeleteHumanLoopInput{
		// HumanLoopName: *string, // Required
	}

	if len(_sagemakera2iruntimeHumanLoopName) > 0 {
		input.HumanLoopName = aws.String(_sagemakera2iruntimeHumanLoopName)
	}

	if resp, err := client.DeleteHumanLoop(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the specified human loop. If the human loop was
// deleted, this operation will return a ResourceNotFoundException error.
func sagemakera2iruntime_DescribeHumanLoop(cfg aws.Config, client *sagemakera2iruntime.Client) {
	input := &sagemakera2iruntime.DescribeHumanLoopInput{
		// HumanLoopName: *string, // Required
	}

	if len(_sagemakera2iruntimeHumanLoopName) > 0 {
		input.HumanLoopName = aws.String(_sagemakera2iruntimeHumanLoopName)
	}

	if resp, err := client.DescribeHumanLoop(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about human loops, given the specified parameters. If a
// human loop was deleted, it will not be included.
func sagemakera2iruntime_ListHumanLoops(cfg aws.Config, client *sagemakera2iruntime.Client) {
	input := &sagemakera2iruntime.ListHumanLoopsInput{
		// FlowDefinitionArn: *string, // Required
	}

	if len(_sagemakera2iruntimeFlowDefinitionArn) > 0 {
		input.FlowDefinitionArn = aws.String(_sagemakera2iruntimeFlowDefinitionArn)
	}
	if len(_sagemakera2iruntimeCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _sagemakera2iruntimeCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakera2iruntimeCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _sagemakera2iruntimeCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakera2iruntimeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakera2iruntimeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakera2iruntimeNextToken) > 0 {
		input.NextToken = aws.String(_sagemakera2iruntimeNextToken)
	}
	if len(_sagemakera2iruntimeSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakera2iruntimeSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListHumanLoops(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemakera2iruntime.ListHumanLoopsOutput
	p := sagemakera2iruntime.NewListHumanLoopsPaginator(client, input)
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

// Starts a human loop, provided that at least one activation condition is met.
func sagemakera2iruntime_StartHumanLoop(cfg aws.Config, client *sagemakera2iruntime.Client) {
	input := &sagemakera2iruntime.StartHumanLoopInput{
		// FlowDefinitionArn: *string, // Required
		// HumanLoopInput: *types.HumanLoopInput, // Required
		// HumanLoopName: *string, // Required
	}

	if len(_sagemakera2iruntimeFlowDefinitionArn) > 0 {
		input.FlowDefinitionArn = aws.String(_sagemakera2iruntimeFlowDefinitionArn)
	}
	if len(_sagemakera2iruntimeHumanLoopInput) > 0 {
		if err := assignInputField(input, "HumanLoopInput", _sagemakera2iruntimeHumanLoopInput); err != nil {
			log.Errorf("invalid --human-loop-input: %s", err.Error())
			return
		}
	}
	if len(_sagemakera2iruntimeHumanLoopName) > 0 {
		input.HumanLoopName = aws.String(_sagemakera2iruntimeHumanLoopName)
	}
	if len(_sagemakera2iruntimeDataAttributes) > 0 {
		if err := assignInputField(input, "DataAttributes", _sagemakera2iruntimeDataAttributes); err != nil {
			log.Errorf("invalid --data-attributes: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartHumanLoop(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops the specified human loop.
func sagemakera2iruntime_StopHumanLoop(cfg aws.Config, client *sagemakera2iruntime.Client) {
	input := &sagemakera2iruntime.StopHumanLoopInput{
		// HumanLoopName: *string, // Required
	}

	if len(_sagemakera2iruntimeHumanLoopName) > 0 {
		input.HumanLoopName = aws.String(_sagemakera2iruntimeHumanLoopName)
	}

	if resp, err := client.StopHumanLoop(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_sagemakera2iruntimeCmd)
	_sagemakera2iruntimeCmd.Flags().SortFlags = false

	_sagemakera2iruntimeCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_sagemakera2iruntimeCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_sagemakera2iruntimeCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_sagemakera2iruntimeCmd.Flags().StringVarP(&_sagemakera2iruntimeCreationTimeAfter, "creation-time-after", "", "", "Creation Time After")
	_sagemakera2iruntimeCmd.Flags().StringVarP(&_sagemakera2iruntimeCreationTimeBefore, "creation-time-before", "", "", "Creation Time Before")
	_sagemakera2iruntimeCmd.Flags().StringVarP(&_sagemakera2iruntimeDataAttributes, "data-attributes", "", "", "Data Attributes")
	_sagemakera2iruntimeCmd.Flags().StringVarP(&_sagemakera2iruntimeFlowDefinitionArn, "flow-definition-arn", "", "", "Flow Definition ARN")
	_sagemakera2iruntimeCmd.Flags().StringVarP(&_sagemakera2iruntimeHumanLoopInput, "human-loop-input", "", "", "Human Loop Input")
	_sagemakera2iruntimeCmd.Flags().StringVarP(&_sagemakera2iruntimeHumanLoopName, "human-loop-name", "", "", "Human Loop Name")
	_sagemakera2iruntimeCmd.Flags().StringVarP(&_sagemakera2iruntimeMaxResults, "max-results", "", "", "Max Results")
	_sagemakera2iruntimeCmd.Flags().StringVarP(&_sagemakera2iruntimeNextToken, "next-token", "", "", "Next Token")
	_sagemakera2iruntimeCmd.Flags().StringVarP(&_sagemakera2iruntimeSortOrder, "sort-order", "", "", "Sort Order")

	_sagemakera2iruntimeCmd.Flags().BoolVarP(&_sagemakera2iruntimeDeleteHumanLoop, "delete-human-loop", "", false, "Delete Human Loop")
	_sagemakera2iruntimeCmd.Flags().BoolVarP(&_sagemakera2iruntimeDescribeHumanLoop, "describe-human-loop", "", false, "Describe Human Loop")
	_sagemakera2iruntimeCmd.Flags().BoolVarP(&_sagemakera2iruntimeListHumanLoops, "list-human-loops", "", false, "List Human Loops")
	_sagemakera2iruntimeCmd.Flags().BoolVarP(&_sagemakera2iruntimeStartHumanLoop, "start-human-loop", "", false, "Start Human Loop")
	_sagemakera2iruntimeCmd.Flags().BoolVarP(&_sagemakera2iruntimeStopHumanLoop, "stop-human-loop", "", false, "Stop Human Loop")

}
