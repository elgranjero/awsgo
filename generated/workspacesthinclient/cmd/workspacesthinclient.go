package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/workspacesthinclient"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// workspacesthinclientCmd represents the workspacesthinclient command
var _workspacesthinclientCmd = &cobra.Command{
	Use:   "workspacesthinclient",
	Short: "AWS workspacesthinclient CLI",
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
		client := workspacesthinclient.NewFromConfig(cfg)
		if _workspacesthinclientCreateEnvironment {
			workspacesthinclient_CreateEnvironment(cfg, client)
			return
		}
		if _workspacesthinclientDeleteDevice {
			workspacesthinclient_DeleteDevice(cfg, client)
			return
		}
		if _workspacesthinclientDeleteEnvironment {
			workspacesthinclient_DeleteEnvironment(cfg, client)
			return
		}
		if _workspacesthinclientDeregisterDevice {
			workspacesthinclient_DeregisterDevice(cfg, client)
			return
		}
		if _workspacesthinclientGetDevice {
			workspacesthinclient_GetDevice(cfg, client)
			return
		}
		if _workspacesthinclientGetEnvironment {
			workspacesthinclient_GetEnvironment(cfg, client)
			return
		}
		if _workspacesthinclientGetSoftwareSet {
			workspacesthinclient_GetSoftwareSet(cfg, client)
			return
		}
		if _workspacesthinclientListDevices {
			workspacesthinclient_ListDevices(cfg, client)
			return
		}
		if _workspacesthinclientListEnvironments {
			workspacesthinclient_ListEnvironments(cfg, client)
			return
		}
		if _workspacesthinclientListSoftwareSets {
			workspacesthinclient_ListSoftwareSets(cfg, client)
			return
		}
		if _workspacesthinclientListTagsForResource {
			workspacesthinclient_ListTagsForResource(cfg, client)
			return
		}
		if _workspacesthinclientTagResource {
			workspacesthinclient_TagResource(cfg, client)
			return
		}
		if _workspacesthinclientUntagResource {
			workspacesthinclient_UntagResource(cfg, client)
			return
		}
		if _workspacesthinclientUpdateDevice {
			workspacesthinclient_UpdateDevice(cfg, client)
			return
		}
		if _workspacesthinclientUpdateEnvironment {
			workspacesthinclient_UpdateEnvironment(cfg, client)
			return
		}
		if _workspacesthinclientUpdateSoftwareSet {
			workspacesthinclient_UpdateSoftwareSet(cfg, client)
			return
		}

	},
}

var (
	_workspacesthinclientCreateEnvironment   bool
	_workspacesthinclientDeleteDevice        bool
	_workspacesthinclientDeleteEnvironment   bool
	_workspacesthinclientDeregisterDevice    bool
	_workspacesthinclientGetDevice           bool
	_workspacesthinclientGetEnvironment      bool
	_workspacesthinclientGetSoftwareSet      bool
	_workspacesthinclientListDevices         bool
	_workspacesthinclientListEnvironments    bool
	_workspacesthinclientListSoftwareSets    bool
	_workspacesthinclientListTagsForResource bool
	_workspacesthinclientTagResource         bool
	_workspacesthinclientUntagResource       bool
	_workspacesthinclientUpdateDevice        bool
	_workspacesthinclientUpdateEnvironment   bool
	_workspacesthinclientUpdateSoftwareSet   bool

	_workspacesthinclientClientToken               string
	_workspacesthinclientDesiredSoftwareSetId      string
	_workspacesthinclientDesktopArn                string
	_workspacesthinclientDesktopEndpoint           string
	_workspacesthinclientDeviceCreationTags        string
	_workspacesthinclientId                        string
	_workspacesthinclientKmsKeyArn                 string
	_workspacesthinclientMaintenanceWindow         string
	_workspacesthinclientMaxResults                string
	_workspacesthinclientName                      string
	_workspacesthinclientNextToken                 string
	_workspacesthinclientResourceArn               string
	_workspacesthinclientSoftwareSetUpdateMode     string
	_workspacesthinclientSoftwareSetUpdateSchedule string
	_workspacesthinclientTagKeys                   []string
	_workspacesthinclientTags                      string
	_workspacesthinclientTargetDeviceStatus        string
	_workspacesthinclientValidationStatus          string
)

// Creates an environment for your thin client devices.
func workspacesthinclient_CreateEnvironment(cfg aws.Config, client *workspacesthinclient.Client) {
	input := &workspacesthinclient.CreateEnvironmentInput{
		// DesktopArn: *string, // Required
	}

	if len(_workspacesthinclientDesktopArn) > 0 {
		input.DesktopArn = aws.String(_workspacesthinclientDesktopArn)
	}
	if len(_workspacesthinclientClientToken) > 0 {
		input.ClientToken = aws.String(_workspacesthinclientClientToken)
	}
	if len(_workspacesthinclientDesiredSoftwareSetId) > 0 {
		input.DesiredSoftwareSetId = aws.String(_workspacesthinclientDesiredSoftwareSetId)
	}
	if len(_workspacesthinclientDesktopEndpoint) > 0 {
		input.DesktopEndpoint = aws.String(_workspacesthinclientDesktopEndpoint)
	}
	if len(_workspacesthinclientDeviceCreationTags) > 0 {
		if err := assignInputField(input, "DeviceCreationTags", _workspacesthinclientDeviceCreationTags); err != nil {
			log.Errorf("invalid --device-creation-tags: %s", err.Error())
			return
		}
	}
	if len(_workspacesthinclientKmsKeyArn) > 0 {
		input.KmsKeyArn = aws.String(_workspacesthinclientKmsKeyArn)
	}
	if len(_workspacesthinclientMaintenanceWindow) > 0 {
		if err := assignInputField(input, "MaintenanceWindow", _workspacesthinclientMaintenanceWindow); err != nil {
			log.Errorf("invalid --maintenance-window: %s", err.Error())
			return
		}
	}
	if len(_workspacesthinclientName) > 0 {
		input.Name = aws.String(_workspacesthinclientName)
	}
	if len(_workspacesthinclientSoftwareSetUpdateMode) > 0 {
		if err := assignInputField(input, "SoftwareSetUpdateMode", _workspacesthinclientSoftwareSetUpdateMode); err != nil {
			log.Errorf("invalid --software-set-update-mode: %s", err.Error())
			return
		}
	}
	if len(_workspacesthinclientSoftwareSetUpdateSchedule) > 0 {
		if err := assignInputField(input, "SoftwareSetUpdateSchedule", _workspacesthinclientSoftwareSetUpdateSchedule); err != nil {
			log.Errorf("invalid --software-set-update-schedule: %s", err.Error())
			return
		}
	}
	if len(_workspacesthinclientTags) > 0 {
		if err := assignInputField(input, "Tags", _workspacesthinclientTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a thin client device.
func workspacesthinclient_DeleteDevice(cfg aws.Config, client *workspacesthinclient.Client) {
	input := &workspacesthinclient.DeleteDeviceInput{
		// Id: *string, // Required
	}

	if len(_workspacesthinclientId) > 0 {
		input.Id = aws.String(_workspacesthinclientId)
	}
	if len(_workspacesthinclientClientToken) > 0 {
		input.ClientToken = aws.String(_workspacesthinclientClientToken)
	}

	if resp, err := client.DeleteDevice(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an environment.
func workspacesthinclient_DeleteEnvironment(cfg aws.Config, client *workspacesthinclient.Client) {
	input := &workspacesthinclient.DeleteEnvironmentInput{
		// Id: *string, // Required
	}

	if len(_workspacesthinclientId) > 0 {
		input.Id = aws.String(_workspacesthinclientId)
	}
	if len(_workspacesthinclientClientToken) > 0 {
		input.ClientToken = aws.String(_workspacesthinclientClientToken)
	}

	if resp, err := client.DeleteEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deregisters a thin client device.
func workspacesthinclient_DeregisterDevice(cfg aws.Config, client *workspacesthinclient.Client) {
	input := &workspacesthinclient.DeregisterDeviceInput{
		// Id: *string, // Required
	}

	if len(_workspacesthinclientId) > 0 {
		input.Id = aws.String(_workspacesthinclientId)
	}
	if len(_workspacesthinclientClientToken) > 0 {
		input.ClientToken = aws.String(_workspacesthinclientClientToken)
	}
	if len(_workspacesthinclientTargetDeviceStatus) > 0 {
		if err := assignInputField(input, "TargetDeviceStatus", _workspacesthinclientTargetDeviceStatus); err != nil {
			log.Errorf("invalid --target-device-status: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeregisterDevice(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information for a thin client device.
func workspacesthinclient_GetDevice(cfg aws.Config, client *workspacesthinclient.Client) {
	input := &workspacesthinclient.GetDeviceInput{
		// Id: *string, // Required
	}

	if len(_workspacesthinclientId) > 0 {
		input.Id = aws.String(_workspacesthinclientId)
	}

	if resp, err := client.GetDevice(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information for an environment.
func workspacesthinclient_GetEnvironment(cfg aws.Config, client *workspacesthinclient.Client) {
	input := &workspacesthinclient.GetEnvironmentInput{
		// Id: *string, // Required
	}

	if len(_workspacesthinclientId) > 0 {
		input.Id = aws.String(_workspacesthinclientId)
	}

	if resp, err := client.GetEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information for a software set.
func workspacesthinclient_GetSoftwareSet(cfg aws.Config, client *workspacesthinclient.Client) {
	input := &workspacesthinclient.GetSoftwareSetInput{
		// Id: *string, // Required
	}

	if len(_workspacesthinclientId) > 0 {
		input.Id = aws.String(_workspacesthinclientId)
	}

	if resp, err := client.GetSoftwareSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of thin client devices.
func workspacesthinclient_ListDevices(cfg aws.Config, client *workspacesthinclient.Client) {
	input := &workspacesthinclient.ListDevicesInput{}

	if len(_workspacesthinclientMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _workspacesthinclientMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_workspacesthinclientNextToken) > 0 {
		input.NextToken = aws.String(_workspacesthinclientNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDevices(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*workspacesthinclient.ListDevicesOutput
	p := workspacesthinclient.NewListDevicesPaginator(client, input)
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

// Returns a list of environments.
func workspacesthinclient_ListEnvironments(cfg aws.Config, client *workspacesthinclient.Client) {
	input := &workspacesthinclient.ListEnvironmentsInput{}

	if len(_workspacesthinclientMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _workspacesthinclientMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_workspacesthinclientNextToken) > 0 {
		input.NextToken = aws.String(_workspacesthinclientNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEnvironments(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*workspacesthinclient.ListEnvironmentsOutput
	p := workspacesthinclient.NewListEnvironmentsPaginator(client, input)
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

// Returns a list of software sets.
func workspacesthinclient_ListSoftwareSets(cfg aws.Config, client *workspacesthinclient.Client) {
	input := &workspacesthinclient.ListSoftwareSetsInput{}

	if len(_workspacesthinclientMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _workspacesthinclientMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_workspacesthinclientNextToken) > 0 {
		input.NextToken = aws.String(_workspacesthinclientNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSoftwareSets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*workspacesthinclient.ListSoftwareSetsOutput
	p := workspacesthinclient.NewListSoftwareSetsPaginator(client, input)
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

// Returns a list of tags for a resource.
func workspacesthinclient_ListTagsForResource(cfg aws.Config, client *workspacesthinclient.Client) {
	input := &workspacesthinclient.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_workspacesthinclientResourceArn) > 0 {
		input.ResourceArn = aws.String(_workspacesthinclientResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assigns one or more tags (key-value pairs) to the specified resource.
func workspacesthinclient_TagResource(cfg aws.Config, client *workspacesthinclient.Client) {
	input := &workspacesthinclient.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_workspacesthinclientResourceArn) > 0 {
		input.ResourceArn = aws.String(_workspacesthinclientResourceArn)
	}
	if len(_workspacesthinclientTags) > 0 {
		if err := assignInputField(input, "Tags", _workspacesthinclientTags); err != nil {
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

// Removes a tag or tags from a resource.
func workspacesthinclient_UntagResource(cfg aws.Config, client *workspacesthinclient.Client) {
	input := &workspacesthinclient.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_workspacesthinclientResourceArn) > 0 {
		input.ResourceArn = aws.String(_workspacesthinclientResourceArn)
	}
	if len(_workspacesthinclientTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _workspacesthinclientTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a thin client device.
func workspacesthinclient_UpdateDevice(cfg aws.Config, client *workspacesthinclient.Client) {
	input := &workspacesthinclient.UpdateDeviceInput{
		// Id: *string, // Required
	}

	if len(_workspacesthinclientId) > 0 {
		input.Id = aws.String(_workspacesthinclientId)
	}
	if len(_workspacesthinclientDesiredSoftwareSetId) > 0 {
		input.DesiredSoftwareSetId = aws.String(_workspacesthinclientDesiredSoftwareSetId)
	}
	if len(_workspacesthinclientName) > 0 {
		input.Name = aws.String(_workspacesthinclientName)
	}
	if len(_workspacesthinclientSoftwareSetUpdateSchedule) > 0 {
		if err := assignInputField(input, "SoftwareSetUpdateSchedule", _workspacesthinclientSoftwareSetUpdateSchedule); err != nil {
			log.Errorf("invalid --software-set-update-schedule: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDevice(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an environment.
func workspacesthinclient_UpdateEnvironment(cfg aws.Config, client *workspacesthinclient.Client) {
	input := &workspacesthinclient.UpdateEnvironmentInput{
		// Id: *string, // Required
	}

	if len(_workspacesthinclientId) > 0 {
		input.Id = aws.String(_workspacesthinclientId)
	}
	if len(_workspacesthinclientDesiredSoftwareSetId) > 0 {
		input.DesiredSoftwareSetId = aws.String(_workspacesthinclientDesiredSoftwareSetId)
	}
	if len(_workspacesthinclientDesktopArn) > 0 {
		input.DesktopArn = aws.String(_workspacesthinclientDesktopArn)
	}
	if len(_workspacesthinclientDesktopEndpoint) > 0 {
		input.DesktopEndpoint = aws.String(_workspacesthinclientDesktopEndpoint)
	}
	if len(_workspacesthinclientDeviceCreationTags) > 0 {
		if err := assignInputField(input, "DeviceCreationTags", _workspacesthinclientDeviceCreationTags); err != nil {
			log.Errorf("invalid --device-creation-tags: %s", err.Error())
			return
		}
	}
	if len(_workspacesthinclientMaintenanceWindow) > 0 {
		if err := assignInputField(input, "MaintenanceWindow", _workspacesthinclientMaintenanceWindow); err != nil {
			log.Errorf("invalid --maintenance-window: %s", err.Error())
			return
		}
	}
	if len(_workspacesthinclientName) > 0 {
		input.Name = aws.String(_workspacesthinclientName)
	}
	if len(_workspacesthinclientSoftwareSetUpdateMode) > 0 {
		if err := assignInputField(input, "SoftwareSetUpdateMode", _workspacesthinclientSoftwareSetUpdateMode); err != nil {
			log.Errorf("invalid --software-set-update-mode: %s", err.Error())
			return
		}
	}
	if len(_workspacesthinclientSoftwareSetUpdateSchedule) > 0 {
		if err := assignInputField(input, "SoftwareSetUpdateSchedule", _workspacesthinclientSoftwareSetUpdateSchedule); err != nil {
			log.Errorf("invalid --software-set-update-schedule: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a software set.
func workspacesthinclient_UpdateSoftwareSet(cfg aws.Config, client *workspacesthinclient.Client) {
	input := &workspacesthinclient.UpdateSoftwareSetInput{
		// Id: *string, // Required
		// ValidationStatus: types.SoftwareSetValidationStatus, // Required
	}

	if len(_workspacesthinclientId) > 0 {
		input.Id = aws.String(_workspacesthinclientId)
	}
	if len(_workspacesthinclientValidationStatus) > 0 {
		if err := assignInputField(input, "ValidationStatus", _workspacesthinclientValidationStatus); err != nil {
			log.Errorf("invalid --validation-status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSoftwareSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_workspacesthinclientCmd)
	_workspacesthinclientCmd.Flags().SortFlags = false

	_workspacesthinclientCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_workspacesthinclientCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_workspacesthinclientCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_workspacesthinclientCmd.Flags().StringVarP(&_workspacesthinclientClientToken, "client-token", "", "", "Client Token")
	_workspacesthinclientCmd.Flags().StringVarP(&_workspacesthinclientDesiredSoftwareSetId, "desired-software-set-id", "", "", "Desired Software Set ID")
	_workspacesthinclientCmd.Flags().StringVarP(&_workspacesthinclientDesktopArn, "desktop-arn", "", "", "Desktop ARN")
	_workspacesthinclientCmd.Flags().StringVarP(&_workspacesthinclientDesktopEndpoint, "desktop-endpoint", "", "", "Desktop Endpoint")
	_workspacesthinclientCmd.Flags().StringVarP(&_workspacesthinclientDeviceCreationTags, "device-creation-tags", "", "", "Device Creation Tags")
	_workspacesthinclientCmd.Flags().StringVarP(&_workspacesthinclientId, "id", "", "", "ID")
	_workspacesthinclientCmd.Flags().StringVarP(&_workspacesthinclientKmsKeyArn, "kms-key-arn", "", "", "KMS Key ARN")
	_workspacesthinclientCmd.Flags().StringVarP(&_workspacesthinclientMaintenanceWindow, "maintenance-window", "", "", "Maintenance Window")
	_workspacesthinclientCmd.Flags().StringVarP(&_workspacesthinclientMaxResults, "max-results", "", "", "Max Results")
	_workspacesthinclientCmd.Flags().StringVarP(&_workspacesthinclientName, "name", "", "", "Name")
	_workspacesthinclientCmd.Flags().StringVarP(&_workspacesthinclientNextToken, "next-token", "", "", "Next Token")
	_workspacesthinclientCmd.Flags().StringVarP(&_workspacesthinclientResourceArn, "resource-arn", "", "", "Resource ARN")
	_workspacesthinclientCmd.Flags().StringVarP(&_workspacesthinclientSoftwareSetUpdateMode, "software-set-update-mode", "", "", "Software Set Update Mode")
	_workspacesthinclientCmd.Flags().StringVarP(&_workspacesthinclientSoftwareSetUpdateSchedule, "software-set-update-schedule", "", "", "Software Set Update Schedule")
	_workspacesthinclientCmd.Flags().StringSliceVarP(&_workspacesthinclientTagKeys, "tag-keys", "", nil, "Tag Keys")
	_workspacesthinclientCmd.Flags().StringVarP(&_workspacesthinclientTags, "tags", "", "", "Tags")
	_workspacesthinclientCmd.Flags().StringVarP(&_workspacesthinclientTargetDeviceStatus, "target-device-status", "", "", "Target Device Status")
	_workspacesthinclientCmd.Flags().StringVarP(&_workspacesthinclientValidationStatus, "validation-status", "", "", "Validation Status")

	_workspacesthinclientCmd.Flags().BoolVarP(&_workspacesthinclientCreateEnvironment, "create-environment", "", false, "Create Environment")
	_workspacesthinclientCmd.Flags().BoolVarP(&_workspacesthinclientDeleteDevice, "delete-device", "", false, "Delete Device")
	_workspacesthinclientCmd.Flags().BoolVarP(&_workspacesthinclientDeleteEnvironment, "delete-environment", "", false, "Delete Environment")
	_workspacesthinclientCmd.Flags().BoolVarP(&_workspacesthinclientDeregisterDevice, "deregister-device", "", false, "Deregister Device")
	_workspacesthinclientCmd.Flags().BoolVarP(&_workspacesthinclientGetDevice, "get-device", "", false, "Get Device")
	_workspacesthinclientCmd.Flags().BoolVarP(&_workspacesthinclientGetEnvironment, "get-environment", "", false, "Get Environment")
	_workspacesthinclientCmd.Flags().BoolVarP(&_workspacesthinclientGetSoftwareSet, "get-software-set", "", false, "Get Software Set")
	_workspacesthinclientCmd.Flags().BoolVarP(&_workspacesthinclientListDevices, "list-devices", "", false, "List Devices")
	_workspacesthinclientCmd.Flags().BoolVarP(&_workspacesthinclientListEnvironments, "list-environments", "", false, "List Environments")
	_workspacesthinclientCmd.Flags().BoolVarP(&_workspacesthinclientListSoftwareSets, "list-software-sets", "", false, "List Software Sets")
	_workspacesthinclientCmd.Flags().BoolVarP(&_workspacesthinclientListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_workspacesthinclientCmd.Flags().BoolVarP(&_workspacesthinclientTagResource, "tag-resource", "", false, "Tag Resource")
	_workspacesthinclientCmd.Flags().BoolVarP(&_workspacesthinclientUntagResource, "untag-resource", "", false, "Untag Resource")
	_workspacesthinclientCmd.Flags().BoolVarP(&_workspacesthinclientUpdateDevice, "update-device", "", false, "Update Device")
	_workspacesthinclientCmd.Flags().BoolVarP(&_workspacesthinclientUpdateEnvironment, "update-environment", "", false, "Update Environment")
	_workspacesthinclientCmd.Flags().BoolVarP(&_workspacesthinclientUpdateSoftwareSet, "update-software-set", "", false, "Update Software Set")

}
