package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iotdeviceadvisor"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// iotdeviceadvisorCmd represents the iotdeviceadvisor command
var _iotdeviceadvisorCmd = &cobra.Command{
	Use:   "iotdeviceadvisor",
	Short: "AWS iotdeviceadvisor CLI",
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
		client := iotdeviceadvisor.NewFromConfig(cfg)
		if _iotdeviceadvisorCreateSuiteDefinition {
			iotdeviceadvisor_CreateSuiteDefinition(cfg, client)
			return
		}
		if _iotdeviceadvisorDeleteSuiteDefinition {
			iotdeviceadvisor_DeleteSuiteDefinition(cfg, client)
			return
		}
		if _iotdeviceadvisorGetEndpoint {
			iotdeviceadvisor_GetEndpoint(cfg, client)
			return
		}
		if _iotdeviceadvisorGetSuiteDefinition {
			iotdeviceadvisor_GetSuiteDefinition(cfg, client)
			return
		}
		if _iotdeviceadvisorGetSuiteRun {
			iotdeviceadvisor_GetSuiteRun(cfg, client)
			return
		}
		if _iotdeviceadvisorGetSuiteRunReport {
			iotdeviceadvisor_GetSuiteRunReport(cfg, client)
			return
		}
		if _iotdeviceadvisorListSuiteDefinitions {
			iotdeviceadvisor_ListSuiteDefinitions(cfg, client)
			return
		}
		if _iotdeviceadvisorListSuiteRuns {
			iotdeviceadvisor_ListSuiteRuns(cfg, client)
			return
		}
		if _iotdeviceadvisorListTagsForResource {
			iotdeviceadvisor_ListTagsForResource(cfg, client)
			return
		}
		if _iotdeviceadvisorStartSuiteRun {
			iotdeviceadvisor_StartSuiteRun(cfg, client)
			return
		}
		if _iotdeviceadvisorStopSuiteRun {
			iotdeviceadvisor_StopSuiteRun(cfg, client)
			return
		}
		if _iotdeviceadvisorTagResource {
			iotdeviceadvisor_TagResource(cfg, client)
			return
		}
		if _iotdeviceadvisorUntagResource {
			iotdeviceadvisor_UntagResource(cfg, client)
			return
		}
		if _iotdeviceadvisorUpdateSuiteDefinition {
			iotdeviceadvisor_UpdateSuiteDefinition(cfg, client)
			return
		}

	},
}

var (
	_iotdeviceadvisorCreateSuiteDefinition bool
	_iotdeviceadvisorDeleteSuiteDefinition bool
	_iotdeviceadvisorGetEndpoint           bool
	_iotdeviceadvisorGetSuiteDefinition    bool
	_iotdeviceadvisorGetSuiteRun           bool
	_iotdeviceadvisorGetSuiteRunReport     bool
	_iotdeviceadvisorListSuiteDefinitions  bool
	_iotdeviceadvisorListSuiteRuns         bool
	_iotdeviceadvisorListTagsForResource   bool
	_iotdeviceadvisorStartSuiteRun         bool
	_iotdeviceadvisorStopSuiteRun          bool
	_iotdeviceadvisorTagResource           bool
	_iotdeviceadvisorUntagResource         bool
	_iotdeviceadvisorUpdateSuiteDefinition bool

	_iotdeviceadvisorAuthenticationMethod         string
	_iotdeviceadvisorCertificateArn               string
	_iotdeviceadvisorClientToken                  string
	_iotdeviceadvisorDeviceRoleArn                string
	_iotdeviceadvisorMaxResults                   string
	_iotdeviceadvisorNextToken                    string
	_iotdeviceadvisorResourceArn                  string
	_iotdeviceadvisorSuiteDefinitionConfiguration string
	_iotdeviceadvisorSuiteDefinitionId            string
	_iotdeviceadvisorSuiteDefinitionVersion       string
	_iotdeviceadvisorSuiteRunConfiguration        string
	_iotdeviceadvisorSuiteRunId                   string
	_iotdeviceadvisorTagKeys                      []string
	_iotdeviceadvisorTags                         string
	_iotdeviceadvisorThingArn                     string
)

// Creates a Device Advisor test suite.
// Requires permission to access the [CreateSuiteDefinition] action.
//
// [CreateSuiteDefinition]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iotdeviceadvisor_CreateSuiteDefinition(cfg aws.Config, client *iotdeviceadvisor.Client) {
	input := &iotdeviceadvisor.CreateSuiteDefinitionInput{
		// SuiteDefinitionConfiguration: *types.SuiteDefinitionConfiguration, // Required
	}

	if len(_iotdeviceadvisorSuiteDefinitionConfiguration) > 0 {
		if err := assignInputField(input, "SuiteDefinitionConfiguration", _iotdeviceadvisorSuiteDefinitionConfiguration); err != nil {
			log.Errorf("invalid --suite-definition-configuration: %s", err.Error())
			return
		}
	}
	if len(_iotdeviceadvisorClientToken) > 0 {
		input.ClientToken = aws.String(_iotdeviceadvisorClientToken)
	}
	if len(_iotdeviceadvisorTags) > 0 {
		if err := assignInputField(input, "Tags", _iotdeviceadvisorTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSuiteDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Device Advisor test suite.
// Requires permission to access the [DeleteSuiteDefinition] action.
//
// [DeleteSuiteDefinition]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iotdeviceadvisor_DeleteSuiteDefinition(cfg aws.Config, client *iotdeviceadvisor.Client) {
	input := &iotdeviceadvisor.DeleteSuiteDefinitionInput{
		// SuiteDefinitionId: *string, // Required
	}

	if len(_iotdeviceadvisorSuiteDefinitionId) > 0 {
		input.SuiteDefinitionId = aws.String(_iotdeviceadvisorSuiteDefinitionId)
	}

	if resp, err := client.DeleteSuiteDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about an Device Advisor endpoint.
func iotdeviceadvisor_GetEndpoint(cfg aws.Config, client *iotdeviceadvisor.Client) {
	input := &iotdeviceadvisor.GetEndpointInput{}

	if len(_iotdeviceadvisorAuthenticationMethod) > 0 {
		if err := assignInputField(input, "AuthenticationMethod", _iotdeviceadvisorAuthenticationMethod); err != nil {
			log.Errorf("invalid --authentication-method: %s", err.Error())
			return
		}
	}
	if len(_iotdeviceadvisorCertificateArn) > 0 {
		input.CertificateArn = aws.String(_iotdeviceadvisorCertificateArn)
	}
	if len(_iotdeviceadvisorDeviceRoleArn) > 0 {
		input.DeviceRoleArn = aws.String(_iotdeviceadvisorDeviceRoleArn)
	}
	if len(_iotdeviceadvisorThingArn) > 0 {
		input.ThingArn = aws.String(_iotdeviceadvisorThingArn)
	}

	if resp, err := client.GetEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a Device Advisor test suite.
// Requires permission to access the [GetSuiteDefinition] action.
//
// [GetSuiteDefinition]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iotdeviceadvisor_GetSuiteDefinition(cfg aws.Config, client *iotdeviceadvisor.Client) {
	input := &iotdeviceadvisor.GetSuiteDefinitionInput{
		// SuiteDefinitionId: *string, // Required
	}

	if len(_iotdeviceadvisorSuiteDefinitionId) > 0 {
		input.SuiteDefinitionId = aws.String(_iotdeviceadvisorSuiteDefinitionId)
	}
	if len(_iotdeviceadvisorSuiteDefinitionVersion) > 0 {
		input.SuiteDefinitionVersion = aws.String(_iotdeviceadvisorSuiteDefinitionVersion)
	}

	if resp, err := client.GetSuiteDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a Device Advisor test suite run.
// Requires permission to access the [GetSuiteRun] action.
//
// [GetSuiteRun]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iotdeviceadvisor_GetSuiteRun(cfg aws.Config, client *iotdeviceadvisor.Client) {
	input := &iotdeviceadvisor.GetSuiteRunInput{
		// SuiteDefinitionId: *string, // Required
		// SuiteRunId: *string, // Required
	}

	if len(_iotdeviceadvisorSuiteDefinitionId) > 0 {
		input.SuiteDefinitionId = aws.String(_iotdeviceadvisorSuiteDefinitionId)
	}
	if len(_iotdeviceadvisorSuiteRunId) > 0 {
		input.SuiteRunId = aws.String(_iotdeviceadvisorSuiteRunId)
	}

	if resp, err := client.GetSuiteRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a report download link for a successful Device Advisor qualifying test
// suite run.
//
// Requires permission to access the [GetSuiteRunReport] action.
//
// [GetSuiteRunReport]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iotdeviceadvisor_GetSuiteRunReport(cfg aws.Config, client *iotdeviceadvisor.Client) {
	input := &iotdeviceadvisor.GetSuiteRunReportInput{
		// SuiteDefinitionId: *string, // Required
		// SuiteRunId: *string, // Required
	}

	if len(_iotdeviceadvisorSuiteDefinitionId) > 0 {
		input.SuiteDefinitionId = aws.String(_iotdeviceadvisorSuiteDefinitionId)
	}
	if len(_iotdeviceadvisorSuiteRunId) > 0 {
		input.SuiteRunId = aws.String(_iotdeviceadvisorSuiteRunId)
	}

	if resp, err := client.GetSuiteRunReport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the Device Advisor test suites you have created.
// Requires permission to access the [ListSuiteDefinitions] action.
//
// [ListSuiteDefinitions]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iotdeviceadvisor_ListSuiteDefinitions(cfg aws.Config, client *iotdeviceadvisor.Client) {
	input := &iotdeviceadvisor.ListSuiteDefinitionsInput{}

	if len(_iotdeviceadvisorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotdeviceadvisorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotdeviceadvisorNextToken) > 0 {
		input.NextToken = aws.String(_iotdeviceadvisorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSuiteDefinitions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotdeviceadvisor.ListSuiteDefinitionsOutput
	p := iotdeviceadvisor.NewListSuiteDefinitionsPaginator(client, input)
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

// Lists runs of the specified Device Advisor test suite. You can list all runs of
// the test suite, or the runs of a specific version of the test suite.
//
// Requires permission to access the [ListSuiteRuns] action.
//
// [ListSuiteRuns]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iotdeviceadvisor_ListSuiteRuns(cfg aws.Config, client *iotdeviceadvisor.Client) {
	input := &iotdeviceadvisor.ListSuiteRunsInput{}

	if len(_iotdeviceadvisorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotdeviceadvisorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotdeviceadvisorNextToken) > 0 {
		input.NextToken = aws.String(_iotdeviceadvisorNextToken)
	}
	if len(_iotdeviceadvisorSuiteDefinitionId) > 0 {
		input.SuiteDefinitionId = aws.String(_iotdeviceadvisorSuiteDefinitionId)
	}
	if len(_iotdeviceadvisorSuiteDefinitionVersion) > 0 {
		input.SuiteDefinitionVersion = aws.String(_iotdeviceadvisorSuiteDefinitionVersion)
	}

	if disablePaginator() {
		if resp, err := client.ListSuiteRuns(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotdeviceadvisor.ListSuiteRunsOutput
	p := iotdeviceadvisor.NewListSuiteRunsPaginator(client, input)
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

// Lists the tags attached to an IoT Device Advisor resource.
// Requires permission to access the [ListTagsForResource] action.
//
// [ListTagsForResource]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iotdeviceadvisor_ListTagsForResource(cfg aws.Config, client *iotdeviceadvisor.Client) {
	input := &iotdeviceadvisor.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_iotdeviceadvisorResourceArn) > 0 {
		input.ResourceArn = aws.String(_iotdeviceadvisorResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a Device Advisor test suite run.
// Requires permission to access the [StartSuiteRun] action.
//
// [StartSuiteRun]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iotdeviceadvisor_StartSuiteRun(cfg aws.Config, client *iotdeviceadvisor.Client) {
	input := &iotdeviceadvisor.StartSuiteRunInput{
		// SuiteDefinitionId: *string, // Required
		// SuiteRunConfiguration: *types.SuiteRunConfiguration, // Required
	}

	if len(_iotdeviceadvisorSuiteDefinitionId) > 0 {
		input.SuiteDefinitionId = aws.String(_iotdeviceadvisorSuiteDefinitionId)
	}
	if len(_iotdeviceadvisorSuiteRunConfiguration) > 0 {
		if err := assignInputField(input, "SuiteRunConfiguration", _iotdeviceadvisorSuiteRunConfiguration); err != nil {
			log.Errorf("invalid --suite-run-configuration: %s", err.Error())
			return
		}
	}
	if len(_iotdeviceadvisorSuiteDefinitionVersion) > 0 {
		input.SuiteDefinitionVersion = aws.String(_iotdeviceadvisorSuiteDefinitionVersion)
	}
	if len(_iotdeviceadvisorTags) > 0 {
		if err := assignInputField(input, "Tags", _iotdeviceadvisorTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartSuiteRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a Device Advisor test suite run that is currently running.
// Requires permission to access the [StopSuiteRun] action.
//
// [StopSuiteRun]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iotdeviceadvisor_StopSuiteRun(cfg aws.Config, client *iotdeviceadvisor.Client) {
	input := &iotdeviceadvisor.StopSuiteRunInput{
		// SuiteDefinitionId: *string, // Required
		// SuiteRunId: *string, // Required
	}

	if len(_iotdeviceadvisorSuiteDefinitionId) > 0 {
		input.SuiteDefinitionId = aws.String(_iotdeviceadvisorSuiteDefinitionId)
	}
	if len(_iotdeviceadvisorSuiteRunId) > 0 {
		input.SuiteRunId = aws.String(_iotdeviceadvisorSuiteRunId)
	}

	if resp, err := client.StopSuiteRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds to and modifies existing tags of an IoT Device Advisor resource.
// Requires permission to access the [TagResource] action.
//
// [TagResource]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iotdeviceadvisor_TagResource(cfg aws.Config, client *iotdeviceadvisor.Client) {
	input := &iotdeviceadvisor.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_iotdeviceadvisorResourceArn) > 0 {
		input.ResourceArn = aws.String(_iotdeviceadvisorResourceArn)
	}
	if len(_iotdeviceadvisorTags) > 0 {
		if err := assignInputField(input, "Tags", _iotdeviceadvisorTags); err != nil {
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

// Removes tags from an IoT Device Advisor resource.
// Requires permission to access the [UntagResource] action.
//
// [UntagResource]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iotdeviceadvisor_UntagResource(cfg aws.Config, client *iotdeviceadvisor.Client) {
	input := &iotdeviceadvisor.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_iotdeviceadvisorResourceArn) > 0 {
		input.ResourceArn = aws.String(_iotdeviceadvisorResourceArn)
	}
	if len(_iotdeviceadvisorTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _iotdeviceadvisorTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a Device Advisor test suite.
// Requires permission to access the [UpdateSuiteDefinition] action.
//
// [UpdateSuiteDefinition]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iotdeviceadvisor_UpdateSuiteDefinition(cfg aws.Config, client *iotdeviceadvisor.Client) {
	input := &iotdeviceadvisor.UpdateSuiteDefinitionInput{
		// SuiteDefinitionConfiguration: *types.SuiteDefinitionConfiguration, // Required
		// SuiteDefinitionId: *string, // Required
	}

	if len(_iotdeviceadvisorSuiteDefinitionConfiguration) > 0 {
		if err := assignInputField(input, "SuiteDefinitionConfiguration", _iotdeviceadvisorSuiteDefinitionConfiguration); err != nil {
			log.Errorf("invalid --suite-definition-configuration: %s", err.Error())
			return
		}
	}
	if len(_iotdeviceadvisorSuiteDefinitionId) > 0 {
		input.SuiteDefinitionId = aws.String(_iotdeviceadvisorSuiteDefinitionId)
	}

	if resp, err := client.UpdateSuiteDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_iotdeviceadvisorCmd)
	_iotdeviceadvisorCmd.Flags().SortFlags = false

	_iotdeviceadvisorCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_iotdeviceadvisorCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_iotdeviceadvisorCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_iotdeviceadvisorCmd.Flags().StringVarP(&_iotdeviceadvisorAuthenticationMethod, "authentication-method", "", "", "Authentication Method")
	_iotdeviceadvisorCmd.Flags().StringVarP(&_iotdeviceadvisorCertificateArn, "certificate-arn", "", "", "Certificate ARN")
	_iotdeviceadvisorCmd.Flags().StringVarP(&_iotdeviceadvisorClientToken, "client-token", "", "", "Client Token")
	_iotdeviceadvisorCmd.Flags().StringVarP(&_iotdeviceadvisorDeviceRoleArn, "device-role-arn", "", "", "Device Role ARN")
	_iotdeviceadvisorCmd.Flags().StringVarP(&_iotdeviceadvisorMaxResults, "max-results", "", "", "Max Results")
	_iotdeviceadvisorCmd.Flags().StringVarP(&_iotdeviceadvisorNextToken, "next-token", "", "", "Next Token")
	_iotdeviceadvisorCmd.Flags().StringVarP(&_iotdeviceadvisorResourceArn, "resource-arn", "", "", "Resource ARN")
	_iotdeviceadvisorCmd.Flags().StringVarP(&_iotdeviceadvisorSuiteDefinitionConfiguration, "suite-definition-configuration", "", "", "Suite Definition Configuration")
	_iotdeviceadvisorCmd.Flags().StringVarP(&_iotdeviceadvisorSuiteDefinitionId, "suite-definition-id", "", "", "Suite Definition ID")
	_iotdeviceadvisorCmd.Flags().StringVarP(&_iotdeviceadvisorSuiteDefinitionVersion, "suite-definition-version", "", "", "Suite Definition Version")
	_iotdeviceadvisorCmd.Flags().StringVarP(&_iotdeviceadvisorSuiteRunConfiguration, "suite-run-configuration", "", "", "Suite Run Configuration")
	_iotdeviceadvisorCmd.Flags().StringVarP(&_iotdeviceadvisorSuiteRunId, "suite-run-id", "", "", "Suite Run ID")
	_iotdeviceadvisorCmd.Flags().StringSliceVarP(&_iotdeviceadvisorTagKeys, "tag-keys", "", nil, "Tag Keys")
	_iotdeviceadvisorCmd.Flags().StringVarP(&_iotdeviceadvisorTags, "tags", "", "", "Tags")
	_iotdeviceadvisorCmd.Flags().StringVarP(&_iotdeviceadvisorThingArn, "thing-arn", "", "", "Thing ARN")

	_iotdeviceadvisorCmd.Flags().BoolVarP(&_iotdeviceadvisorCreateSuiteDefinition, "create-suite-definition", "", false, "Create Suite Definition")
	_iotdeviceadvisorCmd.Flags().BoolVarP(&_iotdeviceadvisorDeleteSuiteDefinition, "delete-suite-definition", "", false, "Delete Suite Definition")
	_iotdeviceadvisorCmd.Flags().BoolVarP(&_iotdeviceadvisorGetEndpoint, "get-endpoint", "", false, "Get Endpoint")
	_iotdeviceadvisorCmd.Flags().BoolVarP(&_iotdeviceadvisorGetSuiteDefinition, "get-suite-definition", "", false, "Get Suite Definition")
	_iotdeviceadvisorCmd.Flags().BoolVarP(&_iotdeviceadvisorGetSuiteRun, "get-suite-run", "", false, "Get Suite Run")
	_iotdeviceadvisorCmd.Flags().BoolVarP(&_iotdeviceadvisorGetSuiteRunReport, "get-suite-run-report", "", false, "Get Suite Run Report")
	_iotdeviceadvisorCmd.Flags().BoolVarP(&_iotdeviceadvisorListSuiteDefinitions, "list-suite-definitions", "", false, "List Suite Definitions")
	_iotdeviceadvisorCmd.Flags().BoolVarP(&_iotdeviceadvisorListSuiteRuns, "list-suite-runs", "", false, "List Suite Runs")
	_iotdeviceadvisorCmd.Flags().BoolVarP(&_iotdeviceadvisorListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_iotdeviceadvisorCmd.Flags().BoolVarP(&_iotdeviceadvisorStartSuiteRun, "start-suite-run", "", false, "Start Suite Run")
	_iotdeviceadvisorCmd.Flags().BoolVarP(&_iotdeviceadvisorStopSuiteRun, "stop-suite-run", "", false, "Stop Suite Run")
	_iotdeviceadvisorCmd.Flags().BoolVarP(&_iotdeviceadvisorTagResource, "tag-resource", "", false, "Tag Resource")
	_iotdeviceadvisorCmd.Flags().BoolVarP(&_iotdeviceadvisorUntagResource, "untag-resource", "", false, "Untag Resource")
	_iotdeviceadvisorCmd.Flags().BoolVarP(&_iotdeviceadvisorUpdateSuiteDefinition, "update-suite-definition", "", false, "Update Suite Definition")

}
