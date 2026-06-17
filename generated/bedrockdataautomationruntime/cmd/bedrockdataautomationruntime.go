package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/bedrockdataautomationruntime"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// bedrockdataautomationruntimeCmd represents the bedrockdataautomationruntime command
var _bedrockdataautomationruntimeCmd = &cobra.Command{
	Use:   "bedrockdataautomationruntime",
	Short: "AWS bedrockdataautomationruntime CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := bedrockdataautomationruntime.NewFromConfig(cfg)
		if _bedrockdataautomationruntimeGetDataAutomationStatus {
			bedrockdataautomationruntime_GetDataAutomationStatus(cfg, client)
			return
		}
		if _bedrockdataautomationruntimeInvokeDataAutomation {
			bedrockdataautomationruntime_InvokeDataAutomation(cfg, client)
			return
		}
		if _bedrockdataautomationruntimeInvokeDataAutomationAsync {
			bedrockdataautomationruntime_InvokeDataAutomationAsync(cfg, client)
			return
		}
		if _bedrockdataautomationruntimeListTagsForResource {
			bedrockdataautomationruntime_ListTagsForResource(cfg, client)
			return
		}
		if _bedrockdataautomationruntimeTagResource {
			bedrockdataautomationruntime_TagResource(cfg, client)
			return
		}
		if _bedrockdataautomationruntimeUntagResource {
			bedrockdataautomationruntime_UntagResource(cfg, client)
			return
		}

	},
}

var (
	_bedrockdataautomationruntimeGetDataAutomationStatus   bool
	_bedrockdataautomationruntimeInvokeDataAutomation      bool
	_bedrockdataautomationruntimeInvokeDataAutomationAsync bool
	_bedrockdataautomationruntimeListTagsForResource       bool
	_bedrockdataautomationruntimeTagResource               bool
	_bedrockdataautomationruntimeUntagResource             bool

	_bedrockdataautomationruntimeBlueprints                  string
	_bedrockdataautomationruntimeClientToken                 string
	_bedrockdataautomationruntimeDataAutomationConfiguration string
	_bedrockdataautomationruntimeDataAutomationProfileArn    string
	_bedrockdataautomationruntimeEncryptionConfiguration     string
	_bedrockdataautomationruntimeInputConfiguration          string
	_bedrockdataautomationruntimeInvocationArn               string
	_bedrockdataautomationruntimeNotificationConfiguration   string
	_bedrockdataautomationruntimeOutputConfiguration         string
	_bedrockdataautomationruntimeResourceARN                 string
	_bedrockdataautomationruntimeTagKeys                     []string
	_bedrockdataautomationruntimeTags                        string
)

// API used to get data automation status.
func bedrockdataautomationruntime_GetDataAutomationStatus(cfg aws.Config, client *bedrockdataautomationruntime.Client) {
	input := &bedrockdataautomationruntime.GetDataAutomationStatusInput{
		// InvocationArn: *string, // Required
	}

	if len(_bedrockdataautomationruntimeInvocationArn) > 0 {
		input.InvocationArn = aws.String(_bedrockdataautomationruntimeInvocationArn)
	}

	if resp, err := client.GetDataAutomationStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sync API: Invoke data automation.
func bedrockdataautomationruntime_InvokeDataAutomation(cfg aws.Config, client *bedrockdataautomationruntime.Client) {
	input := &bedrockdataautomationruntime.InvokeDataAutomationInput{
		// DataAutomationProfileArn: *string, // Required
		// InputConfiguration: *types.SyncInputConfiguration, // Required
	}

	if len(_bedrockdataautomationruntimeDataAutomationProfileArn) > 0 {
		input.DataAutomationProfileArn = aws.String(_bedrockdataautomationruntimeDataAutomationProfileArn)
	}
	if len(_bedrockdataautomationruntimeInputConfiguration) > 0 {
		if err := assignInputField(input, "InputConfiguration", _bedrockdataautomationruntimeInputConfiguration); err != nil {
			log.Errorf("invalid --input-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockdataautomationruntimeBlueprints) > 0 {
		if err := assignInputField(input, "Blueprints", _bedrockdataautomationruntimeBlueprints); err != nil {
			log.Errorf("invalid --blueprints: %s", err.Error())
			return
		}
	}
	if len(_bedrockdataautomationruntimeDataAutomationConfiguration) > 0 {
		if err := assignInputField(input, "DataAutomationConfiguration", _bedrockdataautomationruntimeDataAutomationConfiguration); err != nil {
			log.Errorf("invalid --data-automation-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockdataautomationruntimeEncryptionConfiguration) > 0 {
		if err := assignInputField(input, "EncryptionConfiguration", _bedrockdataautomationruntimeEncryptionConfiguration); err != nil {
			log.Errorf("invalid --encryption-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockdataautomationruntimeOutputConfiguration) > 0 {
		if err := assignInputField(input, "OutputConfiguration", _bedrockdataautomationruntimeOutputConfiguration); err != nil {
			log.Errorf("invalid --output-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.InvokeDataAutomation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Async API: Invoke data automation.
func bedrockdataautomationruntime_InvokeDataAutomationAsync(cfg aws.Config, client *bedrockdataautomationruntime.Client) {
	input := &bedrockdataautomationruntime.InvokeDataAutomationAsyncInput{
		// DataAutomationProfileArn: *string, // Required
		// InputConfiguration: *types.InputConfiguration, // Required
		// OutputConfiguration: *types.OutputConfiguration, // Required
	}

	if len(_bedrockdataautomationruntimeDataAutomationProfileArn) > 0 {
		input.DataAutomationProfileArn = aws.String(_bedrockdataautomationruntimeDataAutomationProfileArn)
	}
	if len(_bedrockdataautomationruntimeInputConfiguration) > 0 {
		if err := assignInputField(input, "InputConfiguration", _bedrockdataautomationruntimeInputConfiguration); err != nil {
			log.Errorf("invalid --input-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockdataautomationruntimeOutputConfiguration) > 0 {
		if err := assignInputField(input, "OutputConfiguration", _bedrockdataautomationruntimeOutputConfiguration); err != nil {
			log.Errorf("invalid --output-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockdataautomationruntimeBlueprints) > 0 {
		if err := assignInputField(input, "Blueprints", _bedrockdataautomationruntimeBlueprints); err != nil {
			log.Errorf("invalid --blueprints: %s", err.Error())
			return
		}
	}
	if len(_bedrockdataautomationruntimeClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockdataautomationruntimeClientToken)
	}
	if len(_bedrockdataautomationruntimeDataAutomationConfiguration) > 0 {
		if err := assignInputField(input, "DataAutomationConfiguration", _bedrockdataautomationruntimeDataAutomationConfiguration); err != nil {
			log.Errorf("invalid --data-automation-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockdataautomationruntimeEncryptionConfiguration) > 0 {
		if err := assignInputField(input, "EncryptionConfiguration", _bedrockdataautomationruntimeEncryptionConfiguration); err != nil {
			log.Errorf("invalid --encryption-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockdataautomationruntimeNotificationConfiguration) > 0 {
		if err := assignInputField(input, "NotificationConfiguration", _bedrockdataautomationruntimeNotificationConfiguration); err != nil {
			log.Errorf("invalid --notification-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockdataautomationruntimeTags) > 0 {
		if err := assignInputField(input, "Tags", _bedrockdataautomationruntimeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.InvokeDataAutomationAsync(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List tags for an Amazon Bedrock Data Automation resource
func bedrockdataautomationruntime_ListTagsForResource(cfg aws.Config, client *bedrockdataautomationruntime.Client) {
	input := &bedrockdataautomationruntime.ListTagsForResourceInput{
		// ResourceARN: *string, // Required
	}

	if len(_bedrockdataautomationruntimeResourceARN) > 0 {
		input.ResourceARN = aws.String(_bedrockdataautomationruntimeResourceARN)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Tag an Amazon Bedrock Data Automation resource
func bedrockdataautomationruntime_TagResource(cfg aws.Config, client *bedrockdataautomationruntime.Client) {
	input := &bedrockdataautomationruntime.TagResourceInput{
		// ResourceARN: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_bedrockdataautomationruntimeResourceARN) > 0 {
		input.ResourceARN = aws.String(_bedrockdataautomationruntimeResourceARN)
	}
	if len(_bedrockdataautomationruntimeTags) > 0 {
		if err := assignInputField(input, "Tags", _bedrockdataautomationruntimeTags); err != nil {
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

// Untag an Amazon Bedrock Data Automation resource
func bedrockdataautomationruntime_UntagResource(cfg aws.Config, client *bedrockdataautomationruntime.Client) {
	input := &bedrockdataautomationruntime.UntagResourceInput{
		// ResourceARN: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_bedrockdataautomationruntimeResourceARN) > 0 {
		input.ResourceARN = aws.String(_bedrockdataautomationruntimeResourceARN)
	}
	if len(_bedrockdataautomationruntimeTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _bedrockdataautomationruntimeTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_bedrockdataautomationruntimeCmd)
	_bedrockdataautomationruntimeCmd.Flags().SortFlags = false

	_bedrockdataautomationruntimeCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_bedrockdataautomationruntimeCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_bedrockdataautomationruntimeCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_bedrockdataautomationruntimeCmd.Flags().StringVarP(&_bedrockdataautomationruntimeBlueprints, "blueprints", "", "", "Blueprints")
	_bedrockdataautomationruntimeCmd.Flags().StringVarP(&_bedrockdataautomationruntimeClientToken, "client-token", "", "", "Client Token")
	_bedrockdataautomationruntimeCmd.Flags().StringVarP(&_bedrockdataautomationruntimeDataAutomationConfiguration, "data-automation-configuration", "", "", "Data Automation Configuration")
	_bedrockdataautomationruntimeCmd.Flags().StringVarP(&_bedrockdataautomationruntimeDataAutomationProfileArn, "data-automation-profile-arn", "", "", "Data Automation Profile ARN")
	_bedrockdataautomationruntimeCmd.Flags().StringVarP(&_bedrockdataautomationruntimeEncryptionConfiguration, "encryption-configuration", "", "", "Encryption Configuration")
	_bedrockdataautomationruntimeCmd.Flags().StringVarP(&_bedrockdataautomationruntimeInputConfiguration, "input-configuration", "", "", "Input Configuration")
	_bedrockdataautomationruntimeCmd.Flags().StringVarP(&_bedrockdataautomationruntimeInvocationArn, "invocation-arn", "", "", "Invocation ARN")
	_bedrockdataautomationruntimeCmd.Flags().StringVarP(&_bedrockdataautomationruntimeNotificationConfiguration, "notification-configuration", "", "", "Notification Configuration")
	_bedrockdataautomationruntimeCmd.Flags().StringVarP(&_bedrockdataautomationruntimeOutputConfiguration, "output-configuration", "", "", "Output Configuration")
	_bedrockdataautomationruntimeCmd.Flags().StringVarP(&_bedrockdataautomationruntimeResourceARN, "resource-arn", "", "", "Resource ARN")
	_bedrockdataautomationruntimeCmd.Flags().StringSliceVarP(&_bedrockdataautomationruntimeTagKeys, "tag-keys", "", nil, "Tag Keys")
	_bedrockdataautomationruntimeCmd.Flags().StringVarP(&_bedrockdataautomationruntimeTags, "tags", "", "", "Tags")

	_bedrockdataautomationruntimeCmd.Flags().BoolVarP(&_bedrockdataautomationruntimeGetDataAutomationStatus, "get-data-automation-status", "", false, "Get Data Automation Status")
	_bedrockdataautomationruntimeCmd.Flags().BoolVarP(&_bedrockdataautomationruntimeInvokeDataAutomation, "invoke-data-automation", "", false, "Invoke Data Automation")
	_bedrockdataautomationruntimeCmd.Flags().BoolVarP(&_bedrockdataautomationruntimeInvokeDataAutomationAsync, "invoke-data-automation-async", "", false, "Invoke Data Automation Async")
	_bedrockdataautomationruntimeCmd.Flags().BoolVarP(&_bedrockdataautomationruntimeListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_bedrockdataautomationruntimeCmd.Flags().BoolVarP(&_bedrockdataautomationruntimeTagResource, "tag-resource", "", false, "Tag Resource")
	_bedrockdataautomationruntimeCmd.Flags().BoolVarP(&_bedrockdataautomationruntimeUntagResource, "untag-resource", "", false, "Untag Resource")

}
