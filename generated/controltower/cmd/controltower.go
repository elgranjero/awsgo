package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/controltower"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// controltowerCmd represents the controltower command
var _controltowerCmd = &cobra.Command{
	Use:   "controltower",
	Short: "AWS controltower CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := controltower.NewFromConfig(cfg)
		if _controltowerCreateLandingZone {
			controltower_CreateLandingZone(cfg, client)
			return
		}
		if _controltowerDeleteLandingZone {
			controltower_DeleteLandingZone(cfg, client)
			return
		}
		if _controltowerDisableBaseline {
			controltower_DisableBaseline(cfg, client)
			return
		}
		if _controltowerDisableControl {
			controltower_DisableControl(cfg, client)
			return
		}
		if _controltowerEnableBaseline {
			controltower_EnableBaseline(cfg, client)
			return
		}
		if _controltowerEnableControl {
			controltower_EnableControl(cfg, client)
			return
		}
		if _controltowerGetBaseline {
			controltower_GetBaseline(cfg, client)
			return
		}
		if _controltowerGetBaselineOperation {
			controltower_GetBaselineOperation(cfg, client)
			return
		}
		if _controltowerGetControlOperation {
			controltower_GetControlOperation(cfg, client)
			return
		}
		if _controltowerGetEnabledBaseline {
			controltower_GetEnabledBaseline(cfg, client)
			return
		}
		if _controltowerGetEnabledControl {
			controltower_GetEnabledControl(cfg, client)
			return
		}
		if _controltowerGetLandingZone {
			controltower_GetLandingZone(cfg, client)
			return
		}
		if _controltowerGetLandingZoneOperation {
			controltower_GetLandingZoneOperation(cfg, client)
			return
		}
		if _controltowerListBaselines {
			controltower_ListBaselines(cfg, client)
			return
		}
		if _controltowerListControlOperations {
			controltower_ListControlOperations(cfg, client)
			return
		}
		if _controltowerListEnabledBaselines {
			controltower_ListEnabledBaselines(cfg, client)
			return
		}
		if _controltowerListEnabledControls {
			controltower_ListEnabledControls(cfg, client)
			return
		}
		if _controltowerListLandingZoneOperations {
			controltower_ListLandingZoneOperations(cfg, client)
			return
		}
		if _controltowerListLandingZones {
			controltower_ListLandingZones(cfg, client)
			return
		}
		if _controltowerListTagsForResource {
			controltower_ListTagsForResource(cfg, client)
			return
		}
		if _controltowerResetEnabledBaseline {
			controltower_ResetEnabledBaseline(cfg, client)
			return
		}
		if _controltowerResetEnabledControl {
			controltower_ResetEnabledControl(cfg, client)
			return
		}
		if _controltowerResetLandingZone {
			controltower_ResetLandingZone(cfg, client)
			return
		}
		if _controltowerTagResource {
			controltower_TagResource(cfg, client)
			return
		}
		if _controltowerUntagResource {
			controltower_UntagResource(cfg, client)
			return
		}
		if _controltowerUpdateEnabledBaseline {
			controltower_UpdateEnabledBaseline(cfg, client)
			return
		}
		if _controltowerUpdateEnabledControl {
			controltower_UpdateEnabledControl(cfg, client)
			return
		}
		if _controltowerUpdateLandingZone {
			controltower_UpdateLandingZone(cfg, client)
			return
		}

	},
}

var (
	_controltowerCreateLandingZone         bool
	_controltowerDeleteLandingZone         bool
	_controltowerDisableBaseline           bool
	_controltowerDisableControl            bool
	_controltowerEnableBaseline            bool
	_controltowerEnableControl             bool
	_controltowerGetBaseline               bool
	_controltowerGetBaselineOperation      bool
	_controltowerGetControlOperation       bool
	_controltowerGetEnabledBaseline        bool
	_controltowerGetEnabledControl         bool
	_controltowerGetLandingZone            bool
	_controltowerGetLandingZoneOperation   bool
	_controltowerListBaselines             bool
	_controltowerListControlOperations     bool
	_controltowerListEnabledBaselines      bool
	_controltowerListEnabledControls       bool
	_controltowerListLandingZoneOperations bool
	_controltowerListLandingZones          bool
	_controltowerListTagsForResource       bool
	_controltowerResetEnabledBaseline      bool
	_controltowerResetEnabledControl       bool
	_controltowerResetLandingZone          bool
	_controltowerTagResource               bool
	_controltowerUntagResource             bool
	_controltowerUpdateEnabledBaseline     bool
	_controltowerUpdateEnabledControl      bool
	_controltowerUpdateLandingZone         bool

	_controltowerBaselineIdentifier        string
	_controltowerBaselineVersion           string
	_controltowerControlIdentifier         string
	_controltowerEnabledBaselineIdentifier string
	_controltowerEnabledControlIdentifier  string
	_controltowerFilter                    string
	_controltowerIncludeChildren           string
	_controltowerLandingZoneIdentifier     string
	_controltowerManifest                  string
	_controltowerMaxResults                string
	_controltowerNextToken                 string
	_controltowerOperationIdentifier       string
	_controltowerParameters                string
	_controltowerRemediationTypes          string
	_controltowerResourceArn               string
	_controltowerTagKeys                   []string
	_controltowerTags                      string
	_controltowerTargetIdentifier          string
	_controltowerVersion                   string
)

// Creates a new landing zone. This API call starts an asynchronous operation that
// creates and configures a landing zone, based on the parameters specified in the
// manifest JSON file.
func controltower_CreateLandingZone(cfg aws.Config, client *controltower.Client) {
	input := &controltower.CreateLandingZoneInput{
		// Version: *string, // Required
	}

	if len(_controltowerVersion) > 0 {
		input.Version = aws.String(_controltowerVersion)
	}
	if len(_controltowerManifest) > 0 {
		if err := assignInputField(input, "Manifest", _controltowerManifest); err != nil {
			log.Errorf("invalid --manifest: %s", err.Error())
			return
		}
	}
	if len(_controltowerRemediationTypes) > 0 {
		if err := assignInputField(input, "RemediationTypes", _controltowerRemediationTypes); err != nil {
			log.Errorf("invalid --remediation-types: %s", err.Error())
			return
		}
	}
	if len(_controltowerTags) > 0 {
		if err := assignInputField(input, "Tags", _controltowerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateLandingZone(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Decommissions a landing zone. This API call starts an asynchronous operation
// that deletes Amazon Web Services Control Tower resources deployed in accounts
// managed by Amazon Web Services Control Tower.
//
// Decommissioning a landing zone is a process with significant consequences, and
// it cannot be undone. We strongly recommend that you perform this decommissioning
// process only if you intend to stop using your landing zone.
func controltower_DeleteLandingZone(cfg aws.Config, client *controltower.Client) {
	input := &controltower.DeleteLandingZoneInput{
		// LandingZoneIdentifier: *string, // Required
	}

	if len(_controltowerLandingZoneIdentifier) > 0 {
		input.LandingZoneIdentifier = aws.String(_controltowerLandingZoneIdentifier)
	}

	if resp, err := client.DeleteLandingZone(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disable an EnabledBaseline resource on the specified Target. This API starts an
// asynchronous operation to remove all resources deployed as part of the baseline
// enablement. The resource will vary depending on the enabled baseline. For usage
// examples, see [the Amazon Web Services Control Tower User Guide].
//
// [the Amazon Web Services Control Tower User Guide]: https://docs.aws.amazon.com/controltower/latest/userguide/baseline-api-examples.html
func controltower_DisableBaseline(cfg aws.Config, client *controltower.Client) {
	input := &controltower.DisableBaselineInput{
		// EnabledBaselineIdentifier: *string, // Required
	}

	if len(_controltowerEnabledBaselineIdentifier) > 0 {
		input.EnabledBaselineIdentifier = aws.String(_controltowerEnabledBaselineIdentifier)
	}

	if resp, err := client.DisableBaseline(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API call turns off a control. It starts an asynchronous operation that
// deletes Amazon Web Services resources on the specified organizational unit and
// the accounts it contains. The resources will vary according to the control that
// you specify. For usage examples, see the [Controls Reference Guide].
//
// [Controls Reference Guide]: https://docs.aws.amazon.com/controltower/latest/controlreference/control-api-examples-short.html
func controltower_DisableControl(cfg aws.Config, client *controltower.Client) {
	input := &controltower.DisableControlInput{}

	if len(_controltowerControlIdentifier) > 0 {
		input.ControlIdentifier = aws.String(_controltowerControlIdentifier)
	}
	if len(_controltowerEnabledControlIdentifier) > 0 {
		input.EnabledControlIdentifier = aws.String(_controltowerEnabledControlIdentifier)
	}
	if len(_controltowerTargetIdentifier) > 0 {
		input.TargetIdentifier = aws.String(_controltowerTargetIdentifier)
	}

	if resp, err := client.DisableControl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enable (apply) a Baseline to a Target. This API starts an asynchronous
// operation to deploy resources specified by the Baseline to the specified
// Target. For usage examples, see [the Amazon Web Services Control Tower User Guide].
//
// [the Amazon Web Services Control Tower User Guide]: https://docs.aws.amazon.com/controltower/latest/userguide/baseline-api-examples.html
func controltower_EnableBaseline(cfg aws.Config, client *controltower.Client) {
	input := &controltower.EnableBaselineInput{
		// BaselineIdentifier: *string, // Required
		// BaselineVersion: *string, // Required
		// TargetIdentifier: *string, // Required
	}

	if len(_controltowerBaselineIdentifier) > 0 {
		input.BaselineIdentifier = aws.String(_controltowerBaselineIdentifier)
	}
	if len(_controltowerBaselineVersion) > 0 {
		input.BaselineVersion = aws.String(_controltowerBaselineVersion)
	}
	if len(_controltowerTargetIdentifier) > 0 {
		input.TargetIdentifier = aws.String(_controltowerTargetIdentifier)
	}
	if len(_controltowerParameters) > 0 {
		if err := assignInputField(input, "Parameters", _controltowerParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_controltowerTags) > 0 {
		if err := assignInputField(input, "Tags", _controltowerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.EnableBaseline(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API call activates a control. It starts an asynchronous operation that
// creates Amazon Web Services resources on the specified organizational unit and
// the accounts it contains. The resources created will vary according to the
// control that you specify. For usage examples, see the [Controls Reference Guide].
//
// [Controls Reference Guide]: https://docs.aws.amazon.com/controltower/latest/controlreference/control-api-examples-short.html
func controltower_EnableControl(cfg aws.Config, client *controltower.Client) {
	input := &controltower.EnableControlInput{
		// ControlIdentifier: *string, // Required
		// TargetIdentifier: *string, // Required
	}

	if len(_controltowerControlIdentifier) > 0 {
		input.ControlIdentifier = aws.String(_controltowerControlIdentifier)
	}
	if len(_controltowerTargetIdentifier) > 0 {
		input.TargetIdentifier = aws.String(_controltowerTargetIdentifier)
	}
	if len(_controltowerParameters) > 0 {
		if err := assignInputField(input, "Parameters", _controltowerParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_controltowerTags) > 0 {
		if err := assignInputField(input, "Tags", _controltowerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.EnableControl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieve details about an existing Baseline resource by specifying its
// identifier. For usage examples, see [the Amazon Web Services Control Tower User Guide].
//
// [the Amazon Web Services Control Tower User Guide]: https://docs.aws.amazon.com/controltower/latest/userguide/baseline-api-examples.html
func controltower_GetBaseline(cfg aws.Config, client *controltower.Client) {
	input := &controltower.GetBaselineInput{
		// BaselineIdentifier: *string, // Required
	}

	if len(_controltowerBaselineIdentifier) > 0 {
		input.BaselineIdentifier = aws.String(_controltowerBaselineIdentifier)
	}

	if resp, err := client.GetBaseline(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the details of an asynchronous baseline operation, as initiated by any
// of these APIs: EnableBaseline , DisableBaseline , UpdateEnabledBaseline ,
// ResetEnabledBaseline . A status message is displayed in case of operation
// failure. For usage examples, see [the Amazon Web Services Control Tower User Guide].
//
// [the Amazon Web Services Control Tower User Guide]: https://docs.aws.amazon.com/controltower/latest/userguide/baseline-api-examples.html
func controltower_GetBaselineOperation(cfg aws.Config, client *controltower.Client) {
	input := &controltower.GetBaselineOperationInput{
		// OperationIdentifier: *string, // Required
	}

	if len(_controltowerOperationIdentifier) > 0 {
		input.OperationIdentifier = aws.String(_controltowerOperationIdentifier)
	}

	if resp, err := client.GetBaselineOperation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the status of a particular EnableControl or DisableControl operation.
// Displays a message in case of error. Details for an operation are available for
// 90 days. For usage examples, see the [Controls Reference Guide].
//
// [Controls Reference Guide]: https://docs.aws.amazon.com/controltower/latest/controlreference/control-api-examples-short.html
func controltower_GetControlOperation(cfg aws.Config, client *controltower.Client) {
	input := &controltower.GetControlOperationInput{
		// OperationIdentifier: *string, // Required
	}

	if len(_controltowerOperationIdentifier) > 0 {
		input.OperationIdentifier = aws.String(_controltowerOperationIdentifier)
	}

	if resp, err := client.GetControlOperation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieve details of an EnabledBaseline resource by specifying its identifier.
func controltower_GetEnabledBaseline(cfg aws.Config, client *controltower.Client) {
	input := &controltower.GetEnabledBaselineInput{
		// EnabledBaselineIdentifier: *string, // Required
	}

	if len(_controltowerEnabledBaselineIdentifier) > 0 {
		input.EnabledBaselineIdentifier = aws.String(_controltowerEnabledBaselineIdentifier)
	}

	if resp, err := client.GetEnabledBaseline(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details about an enabled control. For usage examples, see the [Controls Reference Guide].
//
// [Controls Reference Guide]: https://docs.aws.amazon.com/controltower/latest/controlreference/control-api-examples-short.html
func controltower_GetEnabledControl(cfg aws.Config, client *controltower.Client) {
	input := &controltower.GetEnabledControlInput{
		// EnabledControlIdentifier: *string, // Required
	}

	if len(_controltowerEnabledControlIdentifier) > 0 {
		input.EnabledControlIdentifier = aws.String(_controltowerEnabledControlIdentifier)
	}

	if resp, err := client.GetEnabledControl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns details about the landing zone. Displays a message in case of error.
func controltower_GetLandingZone(cfg aws.Config, client *controltower.Client) {
	input := &controltower.GetLandingZoneInput{
		// LandingZoneIdentifier: *string, // Required
	}

	if len(_controltowerLandingZoneIdentifier) > 0 {
		input.LandingZoneIdentifier = aws.String(_controltowerLandingZoneIdentifier)
	}

	if resp, err := client.GetLandingZone(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the status of the specified landing zone operation. Details for an
// operation are available for 90 days.
func controltower_GetLandingZoneOperation(cfg aws.Config, client *controltower.Client) {
	input := &controltower.GetLandingZoneOperationInput{
		// OperationIdentifier: *string, // Required
	}

	if len(_controltowerOperationIdentifier) > 0 {
		input.OperationIdentifier = aws.String(_controltowerOperationIdentifier)
	}

	if resp, err := client.GetLandingZoneOperation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a summary list of all available baselines. For usage examples, see [the Amazon Web Services Control Tower User Guide].
//
// [the Amazon Web Services Control Tower User Guide]: https://docs.aws.amazon.com/controltower/latest/userguide/baseline-api-examples.html
func controltower_ListBaselines(cfg aws.Config, client *controltower.Client) {
	input := &controltower.ListBaselinesInput{}

	if len(_controltowerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _controltowerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_controltowerNextToken) > 0 {
		input.NextToken = aws.String(_controltowerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListBaselines(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*controltower.ListBaselinesOutput
	p := controltower.NewListBaselinesPaginator(client, input)
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

// Provides a list of operations in progress or queued. For usage examples, see [ListControlOperation examples].
//
// [ListControlOperation examples]: https://docs.aws.amazon.com/controltower/latest/controlreference/control-api-examples-short.html#list-control-operations-api-examples
func controltower_ListControlOperations(cfg aws.Config, client *controltower.Client) {
	input := &controltower.ListControlOperationsInput{}

	if len(_controltowerFilter) > 0 {
		if err := assignInputField(input, "Filter", _controltowerFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_controltowerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _controltowerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_controltowerNextToken) > 0 {
		input.NextToken = aws.String(_controltowerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListControlOperations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*controltower.ListControlOperationsOutput
	p := controltower.NewListControlOperationsPaginator(client, input)
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

// Returns a list of summaries describing EnabledBaseline resources. You can
// filter the list by the corresponding Baseline or Target of the EnabledBaseline
// resources. For usage examples, see [the Amazon Web Services Control Tower User Guide].
//
// [the Amazon Web Services Control Tower User Guide]: https://docs.aws.amazon.com/controltower/latest/userguide/baseline-api-examples.html
func controltower_ListEnabledBaselines(cfg aws.Config, client *controltower.Client) {
	input := &controltower.ListEnabledBaselinesInput{}

	if len(_controltowerFilter) > 0 {
		if err := assignInputField(input, "Filter", _controltowerFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_controltowerIncludeChildren) > 0 {
		if err := assignInputField(input, "IncludeChildren", _controltowerIncludeChildren); err != nil {
			log.Errorf("invalid --include-children: %s", err.Error())
			return
		}
	}
	if len(_controltowerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _controltowerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_controltowerNextToken) > 0 {
		input.NextToken = aws.String(_controltowerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEnabledBaselines(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*controltower.ListEnabledBaselinesOutput
	p := controltower.NewListEnabledBaselinesPaginator(client, input)
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

// Lists the controls enabled by Amazon Web Services Control Tower on the
// specified organizational unit and the accounts it contains. For usage examples,
// see the [Controls Reference Guide].
//
// [Controls Reference Guide]: https://docs.aws.amazon.com/controltower/latest/controlreference/control-api-examples-short.html
func controltower_ListEnabledControls(cfg aws.Config, client *controltower.Client) {
	input := &controltower.ListEnabledControlsInput{}

	if len(_controltowerFilter) > 0 {
		if err := assignInputField(input, "Filter", _controltowerFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_controltowerIncludeChildren) > 0 {
		if err := assignInputField(input, "IncludeChildren", _controltowerIncludeChildren); err != nil {
			log.Errorf("invalid --include-children: %s", err.Error())
			return
		}
	}
	if len(_controltowerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _controltowerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_controltowerNextToken) > 0 {
		input.NextToken = aws.String(_controltowerNextToken)
	}
	if len(_controltowerTargetIdentifier) > 0 {
		input.TargetIdentifier = aws.String(_controltowerTargetIdentifier)
	}

	if disablePaginator() {
		if resp, err := client.ListEnabledControls(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*controltower.ListEnabledControlsOutput
	p := controltower.NewListEnabledControlsPaginator(client, input)
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

// Lists all landing zone operations from the past 90 days. Results are sorted by
// time, with the most recent operation first.
func controltower_ListLandingZoneOperations(cfg aws.Config, client *controltower.Client) {
	input := &controltower.ListLandingZoneOperationsInput{}

	if len(_controltowerFilter) > 0 {
		if err := assignInputField(input, "Filter", _controltowerFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_controltowerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _controltowerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_controltowerNextToken) > 0 {
		input.NextToken = aws.String(_controltowerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListLandingZoneOperations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*controltower.ListLandingZoneOperationsOutput
	p := controltower.NewListLandingZoneOperationsPaginator(client, input)
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

// Returns the landing zone ARN for the landing zone deployed in your managed
// account. This API also creates an ARN for existing accounts that do not yet have
// a landing zone ARN.
//
// Returns one landing zone ARN.
func controltower_ListLandingZones(cfg aws.Config, client *controltower.Client) {
	input := &controltower.ListLandingZonesInput{}

	if len(_controltowerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _controltowerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_controltowerNextToken) > 0 {
		input.NextToken = aws.String(_controltowerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListLandingZones(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*controltower.ListLandingZonesOutput
	p := controltower.NewListLandingZonesPaginator(client, input)
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

// Returns a list of tags associated with the resource. For usage examples, see
// the [Controls Reference Guide].
//
// [Controls Reference Guide]: https://docs.aws.amazon.com/controltower/latest/controlreference/control-api-examples-short.html
func controltower_ListTagsForResource(cfg aws.Config, client *controltower.Client) {
	input := &controltower.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_controltowerResourceArn) > 0 {
		input.ResourceArn = aws.String(_controltowerResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Re-enables an EnabledBaseline resource. For example, this API can re-apply the
// existing Baseline after a new member account is moved to the target OU. For
// usage examples, see [the Amazon Web Services Control Tower User Guide].
//
// [the Amazon Web Services Control Tower User Guide]: https://docs.aws.amazon.com/controltower/latest/userguide/baseline-api-examples.html
func controltower_ResetEnabledBaseline(cfg aws.Config, client *controltower.Client) {
	input := &controltower.ResetEnabledBaselineInput{
		// EnabledBaselineIdentifier: *string, // Required
	}

	if len(_controltowerEnabledBaselineIdentifier) > 0 {
		input.EnabledBaselineIdentifier = aws.String(_controltowerEnabledBaselineIdentifier)
	}

	if resp, err := client.ResetEnabledBaseline(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Resets an enabled control. Does not work for controls implemented with SCPs.
func controltower_ResetEnabledControl(cfg aws.Config, client *controltower.Client) {
	input := &controltower.ResetEnabledControlInput{
		// EnabledControlIdentifier: *string, // Required
	}

	if len(_controltowerEnabledControlIdentifier) > 0 {
		input.EnabledControlIdentifier = aws.String(_controltowerEnabledControlIdentifier)
	}

	if resp, err := client.ResetEnabledControl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API call resets a landing zone. It starts an asynchronous operation that
// resets the landing zone to the parameters specified in the original
// configuration, which you specified in the manifest file. Nothing in the manifest
// file's original landing zone configuration is changed during the reset process,
// by default. This API is not the same as a rollback of a landing zone version,
// which is not a supported operation.
func controltower_ResetLandingZone(cfg aws.Config, client *controltower.Client) {
	input := &controltower.ResetLandingZoneInput{
		// LandingZoneIdentifier: *string, // Required
	}

	if len(_controltowerLandingZoneIdentifier) > 0 {
		input.LandingZoneIdentifier = aws.String(_controltowerLandingZoneIdentifier)
	}

	if resp, err := client.ResetLandingZone(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Applies tags to a resource. For usage examples, see the [Controls Reference Guide].
//
// [Controls Reference Guide]: https://docs.aws.amazon.com/controltower/latest/controlreference/control-api-examples-short.html
func controltower_TagResource(cfg aws.Config, client *controltower.Client) {
	input := &controltower.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_controltowerResourceArn) > 0 {
		input.ResourceArn = aws.String(_controltowerResourceArn)
	}
	if len(_controltowerTags) > 0 {
		if err := assignInputField(input, "Tags", _controltowerTags); err != nil {
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

// Removes tags from a resource. For usage examples, see the [Controls Reference Guide].
//
// [Controls Reference Guide]: https://docs.aws.amazon.com/controltower/latest/controlreference/control-api-examples-short.html
func controltower_UntagResource(cfg aws.Config, client *controltower.Client) {
	input := &controltower.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_controltowerResourceArn) > 0 {
		input.ResourceArn = aws.String(_controltowerResourceArn)
	}
	if len(_controltowerTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _controltowerTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an EnabledBaseline resource's applied parameters or version. For usage
// examples, see [the Amazon Web Services Control Tower User Guide].
//
// [the Amazon Web Services Control Tower User Guide]: https://docs.aws.amazon.com/controltower/latest/userguide/baseline-api-examples.html
func controltower_UpdateEnabledBaseline(cfg aws.Config, client *controltower.Client) {
	input := &controltower.UpdateEnabledBaselineInput{
		// BaselineVersion: *string, // Required
		// EnabledBaselineIdentifier: *string, // Required
	}

	if len(_controltowerBaselineVersion) > 0 {
		input.BaselineVersion = aws.String(_controltowerBaselineVersion)
	}
	if len(_controltowerEnabledBaselineIdentifier) > 0 {
		input.EnabledBaselineIdentifier = aws.String(_controltowerEnabledBaselineIdentifier)
	}
	if len(_controltowerParameters) > 0 {
		if err := assignInputField(input, "Parameters", _controltowerParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateEnabledBaseline(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration of an already enabled control.
// If the enabled control shows an EnablementStatus of SUCCEEDED, supply
// parameters that are different from the currently configured parameters.
// Otherwise, Amazon Web Services Control Tower will not accept the request.
//
// If the enabled control shows an EnablementStatus of FAILED, Amazon Web Services
// Control Tower updates the control to match any valid parameters that you supply.
//
// If the DriftSummary status for the control shows as DRIFTED , you cannot call
// this API. Instead, you can update the control by calling the ResetEnabledControl
// API. Alternatively, you can call DisableControl and then call EnableControl
// again. Also, you can run an extending governance operation to repair drift. For
// usage examples, see the [Controls Reference Guide].
//
// [Controls Reference Guide]: https://docs.aws.amazon.com/controltower/latest/controlreference/control-api-examples-short.html
func controltower_UpdateEnabledControl(cfg aws.Config, client *controltower.Client) {
	input := &controltower.UpdateEnabledControlInput{
		// EnabledControlIdentifier: *string, // Required
		// Parameters: []types.EnabledControlParameter, // Required
	}

	if len(_controltowerEnabledControlIdentifier) > 0 {
		input.EnabledControlIdentifier = aws.String(_controltowerEnabledControlIdentifier)
	}
	if len(_controltowerParameters) > 0 {
		if err := assignInputField(input, "Parameters", _controltowerParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateEnabledControl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API call updates the landing zone. It starts an asynchronous operation
// that updates the landing zone based on the new landing zone version, or on the
// changed parameters specified in the updated manifest file.
func controltower_UpdateLandingZone(cfg aws.Config, client *controltower.Client) {
	input := &controltower.UpdateLandingZoneInput{
		// LandingZoneIdentifier: *string, // Required
		// Version: *string, // Required
	}

	if len(_controltowerLandingZoneIdentifier) > 0 {
		input.LandingZoneIdentifier = aws.String(_controltowerLandingZoneIdentifier)
	}
	if len(_controltowerVersion) > 0 {
		input.Version = aws.String(_controltowerVersion)
	}
	if len(_controltowerManifest) > 0 {
		if err := assignInputField(input, "Manifest", _controltowerManifest); err != nil {
			log.Errorf("invalid --manifest: %s", err.Error())
			return
		}
	}
	if len(_controltowerRemediationTypes) > 0 {
		if err := assignInputField(input, "RemediationTypes", _controltowerRemediationTypes); err != nil {
			log.Errorf("invalid --remediation-types: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateLandingZone(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_controltowerCmd)
	_controltowerCmd.Flags().SortFlags = false

	_controltowerCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_controltowerCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_controltowerCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_controltowerCmd.Flags().StringVarP(&_controltowerBaselineIdentifier, "baseline-identifier", "", "", "Baseline Identifier")
	_controltowerCmd.Flags().StringVarP(&_controltowerBaselineVersion, "baseline-version", "", "", "Baseline Version")
	_controltowerCmd.Flags().StringVarP(&_controltowerControlIdentifier, "control-identifier", "", "", "Control Identifier")
	_controltowerCmd.Flags().StringVarP(&_controltowerEnabledBaselineIdentifier, "enabled-baseline-identifier", "", "", "Enabled Baseline Identifier")
	_controltowerCmd.Flags().StringVarP(&_controltowerEnabledControlIdentifier, "enabled-control-identifier", "", "", "Enabled Control Identifier")
	_controltowerCmd.Flags().StringVarP(&_controltowerFilter, "filter", "", "", "Filter")
	_controltowerCmd.Flags().StringVarP(&_controltowerIncludeChildren, "include-children", "", "", "Include Children")
	_controltowerCmd.Flags().StringVarP(&_controltowerLandingZoneIdentifier, "landing-zone-identifier", "", "", "Landing Zone Identifier")
	_controltowerCmd.Flags().StringVarP(&_controltowerManifest, "manifest", "", "", "Manifest")
	_controltowerCmd.Flags().StringVarP(&_controltowerMaxResults, "max-results", "", "", "Max Results")
	_controltowerCmd.Flags().StringVarP(&_controltowerNextToken, "next-token", "", "", "Next Token")
	_controltowerCmd.Flags().StringVarP(&_controltowerOperationIdentifier, "operation-identifier", "", "", "Operation Identifier")
	_controltowerCmd.Flags().StringVarP(&_controltowerParameters, "parameters", "", "", "Parameters")
	_controltowerCmd.Flags().StringVarP(&_controltowerRemediationTypes, "remediation-types", "", "", "Remediation Types")
	_controltowerCmd.Flags().StringVarP(&_controltowerResourceArn, "resource-arn", "", "", "Resource ARN")
	_controltowerCmd.Flags().StringSliceVarP(&_controltowerTagKeys, "tag-keys", "", nil, "Tag Keys")
	_controltowerCmd.Flags().StringVarP(&_controltowerTags, "tags", "", "", "Tags")
	_controltowerCmd.Flags().StringVarP(&_controltowerTargetIdentifier, "target-identifier", "", "", "Target Identifier")
	_controltowerCmd.Flags().StringVarP(&_controltowerVersion, "version", "", "", "Version")

	_controltowerCmd.Flags().BoolVarP(&_controltowerCreateLandingZone, "create-landing-zone", "", false, "Create Landing Zone")
	_controltowerCmd.Flags().BoolVarP(&_controltowerDeleteLandingZone, "delete-landing-zone", "", false, "Delete Landing Zone")
	_controltowerCmd.Flags().BoolVarP(&_controltowerDisableBaseline, "disable-baseline", "", false, "Disable Baseline")
	_controltowerCmd.Flags().BoolVarP(&_controltowerDisableControl, "disable-control", "", false, "Disable Control")
	_controltowerCmd.Flags().BoolVarP(&_controltowerEnableBaseline, "enable-baseline", "", false, "Enable Baseline")
	_controltowerCmd.Flags().BoolVarP(&_controltowerEnableControl, "enable-control", "", false, "Enable Control")
	_controltowerCmd.Flags().BoolVarP(&_controltowerGetBaseline, "get-baseline", "", false, "Get Baseline")
	_controltowerCmd.Flags().BoolVarP(&_controltowerGetBaselineOperation, "get-baseline-operation", "", false, "Get Baseline Operation")
	_controltowerCmd.Flags().BoolVarP(&_controltowerGetControlOperation, "get-control-operation", "", false, "Get Control Operation")
	_controltowerCmd.Flags().BoolVarP(&_controltowerGetEnabledBaseline, "get-enabled-baseline", "", false, "Get Enabled Baseline")
	_controltowerCmd.Flags().BoolVarP(&_controltowerGetEnabledControl, "get-enabled-control", "", false, "Get Enabled Control")
	_controltowerCmd.Flags().BoolVarP(&_controltowerGetLandingZone, "get-landing-zone", "", false, "Get Landing Zone")
	_controltowerCmd.Flags().BoolVarP(&_controltowerGetLandingZoneOperation, "get-landing-zone-operation", "", false, "Get Landing Zone Operation")
	_controltowerCmd.Flags().BoolVarP(&_controltowerListBaselines, "list-baselines", "", false, "List Baselines")
	_controltowerCmd.Flags().BoolVarP(&_controltowerListControlOperations, "list-control-operations", "", false, "List Control Operations")
	_controltowerCmd.Flags().BoolVarP(&_controltowerListEnabledBaselines, "list-enabled-baselines", "", false, "List Enabled Baselines")
	_controltowerCmd.Flags().BoolVarP(&_controltowerListEnabledControls, "list-enabled-controls", "", false, "List Enabled Controls")
	_controltowerCmd.Flags().BoolVarP(&_controltowerListLandingZoneOperations, "list-landing-zone-operations", "", false, "List Landing Zone Operations")
	_controltowerCmd.Flags().BoolVarP(&_controltowerListLandingZones, "list-landing-zones", "", false, "List Landing Zones")
	_controltowerCmd.Flags().BoolVarP(&_controltowerListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_controltowerCmd.Flags().BoolVarP(&_controltowerResetEnabledBaseline, "reset-enabled-baseline", "", false, "Reset Enabled Baseline")
	_controltowerCmd.Flags().BoolVarP(&_controltowerResetEnabledControl, "reset-enabled-control", "", false, "Reset Enabled Control")
	_controltowerCmd.Flags().BoolVarP(&_controltowerResetLandingZone, "reset-landing-zone", "", false, "Reset Landing Zone")
	_controltowerCmd.Flags().BoolVarP(&_controltowerTagResource, "tag-resource", "", false, "Tag Resource")
	_controltowerCmd.Flags().BoolVarP(&_controltowerUntagResource, "untag-resource", "", false, "Untag Resource")
	_controltowerCmd.Flags().BoolVarP(&_controltowerUpdateEnabledBaseline, "update-enabled-baseline", "", false, "Update Enabled Baseline")
	_controltowerCmd.Flags().BoolVarP(&_controltowerUpdateEnabledControl, "update-enabled-control", "", false, "Update Enabled Control")
	_controltowerCmd.Flags().BoolVarP(&_controltowerUpdateLandingZone, "update-landing-zone", "", false, "Update Landing Zone")

}
