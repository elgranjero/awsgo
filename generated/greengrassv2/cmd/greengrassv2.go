package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/greengrassv2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// greengrassv2Cmd represents the greengrassv2 command
var _greengrassv2Cmd = &cobra.Command{
	Use:   "greengrassv2",
	Short: "AWS greengrassv2 CLI",
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
		client := greengrassv2.NewFromConfig(cfg)
		if _greengrassv2AssociateServiceRoleToAccount {
			greengrassv2_AssociateServiceRoleToAccount(cfg, client)
			return
		}
		if _greengrassv2BatchAssociateClientDeviceWithCoreDevice {
			greengrassv2_BatchAssociateClientDeviceWithCoreDevice(cfg, client)
			return
		}
		if _greengrassv2BatchDisassociateClientDeviceFromCoreDevice {
			greengrassv2_BatchDisassociateClientDeviceFromCoreDevice(cfg, client)
			return
		}
		if _greengrassv2CancelDeployment {
			greengrassv2_CancelDeployment(cfg, client)
			return
		}
		if _greengrassv2CreateComponentVersion {
			greengrassv2_CreateComponentVersion(cfg, client)
			return
		}
		if _greengrassv2CreateDeployment {
			greengrassv2_CreateDeployment(cfg, client)
			return
		}
		if _greengrassv2DeleteComponent {
			greengrassv2_DeleteComponent(cfg, client)
			return
		}
		if _greengrassv2DeleteCoreDevice {
			greengrassv2_DeleteCoreDevice(cfg, client)
			return
		}
		if _greengrassv2DeleteDeployment {
			greengrassv2_DeleteDeployment(cfg, client)
			return
		}
		if _greengrassv2DescribeComponent {
			greengrassv2_DescribeComponent(cfg, client)
			return
		}
		if _greengrassv2DisassociateServiceRoleFromAccount {
			greengrassv2_DisassociateServiceRoleFromAccount(cfg, client)
			return
		}
		if _greengrassv2GetComponent {
			greengrassv2_GetComponent(cfg, client)
			return
		}
		if _greengrassv2GetComponentVersionArtifact {
			greengrassv2_GetComponentVersionArtifact(cfg, client)
			return
		}
		if _greengrassv2GetConnectivityInfo {
			greengrassv2_GetConnectivityInfo(cfg, client)
			return
		}
		if _greengrassv2GetCoreDevice {
			greengrassv2_GetCoreDevice(cfg, client)
			return
		}
		if _greengrassv2GetDeployment {
			greengrassv2_GetDeployment(cfg, client)
			return
		}
		if _greengrassv2GetServiceRoleForAccount {
			greengrassv2_GetServiceRoleForAccount(cfg, client)
			return
		}
		if _greengrassv2ListClientDevicesAssociatedWithCoreDevice {
			greengrassv2_ListClientDevicesAssociatedWithCoreDevice(cfg, client)
			return
		}
		if _greengrassv2ListComponentVersions {
			greengrassv2_ListComponentVersions(cfg, client)
			return
		}
		if _greengrassv2ListComponents {
			greengrassv2_ListComponents(cfg, client)
			return
		}
		if _greengrassv2ListCoreDevices {
			greengrassv2_ListCoreDevices(cfg, client)
			return
		}
		if _greengrassv2ListDeployments {
			greengrassv2_ListDeployments(cfg, client)
			return
		}
		if _greengrassv2ListEffectiveDeployments {
			greengrassv2_ListEffectiveDeployments(cfg, client)
			return
		}
		if _greengrassv2ListInstalledComponents {
			greengrassv2_ListInstalledComponents(cfg, client)
			return
		}
		if _greengrassv2ListTagsForResource {
			greengrassv2_ListTagsForResource(cfg, client)
			return
		}
		if _greengrassv2ResolveComponentCandidates {
			greengrassv2_ResolveComponentCandidates(cfg, client)
			return
		}
		if _greengrassv2TagResource {
			greengrassv2_TagResource(cfg, client)
			return
		}
		if _greengrassv2UntagResource {
			greengrassv2_UntagResource(cfg, client)
			return
		}
		if _greengrassv2UpdateConnectivityInfo {
			greengrassv2_UpdateConnectivityInfo(cfg, client)
			return
		}

	},
}

var (
	_greengrassv2AssociateServiceRoleToAccount               bool
	_greengrassv2BatchAssociateClientDeviceWithCoreDevice    bool
	_greengrassv2BatchDisassociateClientDeviceFromCoreDevice bool
	_greengrassv2CancelDeployment                            bool
	_greengrassv2CreateComponentVersion                      bool
	_greengrassv2CreateDeployment                            bool
	_greengrassv2DeleteComponent                             bool
	_greengrassv2DeleteCoreDevice                            bool
	_greengrassv2DeleteDeployment                            bool
	_greengrassv2DescribeComponent                           bool
	_greengrassv2DisassociateServiceRoleFromAccount          bool
	_greengrassv2GetComponent                                bool
	_greengrassv2GetComponentVersionArtifact                 bool
	_greengrassv2GetConnectivityInfo                         bool
	_greengrassv2GetCoreDevice                               bool
	_greengrassv2GetDeployment                               bool
	_greengrassv2GetServiceRoleForAccount                    bool
	_greengrassv2ListClientDevicesAssociatedWithCoreDevice   bool
	_greengrassv2ListComponentVersions                       bool
	_greengrassv2ListComponents                              bool
	_greengrassv2ListCoreDevices                             bool
	_greengrassv2ListDeployments                             bool
	_greengrassv2ListEffectiveDeployments                    bool
	_greengrassv2ListInstalledComponents                     bool
	_greengrassv2ListTagsForResource                         bool
	_greengrassv2ResolveComponentCandidates                  bool
	_greengrassv2TagResource                                 bool
	_greengrassv2UntagResource                               bool
	_greengrassv2UpdateConnectivityInfo                      bool

	_greengrassv2Arn                 string
	_greengrassv2ArtifactName        string
	_greengrassv2ClientToken         string
	_greengrassv2ComponentCandidates string
	_greengrassv2Components          string
	_greengrassv2ConnectivityInfo    string
	_greengrassv2CoreDeviceThingName string
	_greengrassv2DeploymentId        string
	_greengrassv2DeploymentName      string
	_greengrassv2DeploymentPolicies  string
	_greengrassv2Entries             string
	_greengrassv2HistoryFilter       string
	_greengrassv2InlineRecipe        string
	_greengrassv2IotEndpointType     string
	_greengrassv2IotJobConfiguration string
	_greengrassv2LambdaFunction      string
	_greengrassv2MaxResults          string
	_greengrassv2NextToken           string
	_greengrassv2ParentTargetArn     string
	_greengrassv2Platform            string
	_greengrassv2RecipeOutputFormat  string
	_greengrassv2ResourceArn         string
	_greengrassv2RoleArn             string
	_greengrassv2Runtime             string
	_greengrassv2S3EndpointType      string
	_greengrassv2Scope               string
	_greengrassv2Status              string
	_greengrassv2TagKeys             []string
	_greengrassv2Tags                string
	_greengrassv2TargetArn           string
	_greengrassv2ThingGroupArn       string
	_greengrassv2ThingName           string
	_greengrassv2TopologyFilter      string
)

// Associates a Greengrass service role with IoT Greengrass for your Amazon Web
// Services account in this Amazon Web Services Region. IoT Greengrass uses this
// role to verify the identity of client devices and manage core device
// connectivity information. The role must include the [AWSGreengrassResourceAccessRolePolicy]managed policy or a custom
// policy that defines equivalent permissions for the IoT Greengrass features that
// you use. For more information, see [Greengrass service role]in the IoT Greengrass Version 2 Developer
// Guide.
//
// [AWSGreengrassResourceAccessRolePolicy]: https://console.aws.amazon.com/iam/home#/policies/arn:awsiam::aws:policy/service-role/AWSGreengrassResourceAccessRolePolicy
// [Greengrass service role]: https://docs.aws.amazon.com/greengrass/v2/developerguide/greengrass-service-role.html
func greengrassv2_AssociateServiceRoleToAccount(cfg aws.Config, client *greengrassv2.Client) {
	input := &greengrassv2.AssociateServiceRoleToAccountInput{
		// RoleArn: *string, // Required
	}

	if len(_greengrassv2RoleArn) > 0 {
		input.RoleArn = aws.String(_greengrassv2RoleArn)
	}

	if resp, err := client.AssociateServiceRoleToAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a list of client devices with a core device. Use this API operation
// to specify which client devices can discover a core device through cloud
// discovery. With cloud discovery, client devices connect to IoT Greengrass to
// retrieve associated core devices' connectivity information and certificates. For
// more information, see [Configure cloud discovery]in the IoT Greengrass V2 Developer Guide.
//
// Client devices are local IoT devices that connect to and communicate with an
// IoT Greengrass core device over MQTT. You can connect client devices to a core
// device to sync MQTT messages and data to Amazon Web Services IoT Core and
// interact with client devices in Greengrass components. For more information, see
// [Interact with local IoT devices]in the IoT Greengrass V2 Developer Guide.
//
// [Configure cloud discovery]: https://docs.aws.amazon.com/greengrass/v2/developerguide/configure-cloud-discovery.html
// [Interact with local IoT devices]: https://docs.aws.amazon.com/greengrass/v2/developerguide/interact-with-local-iot-devices.html
func greengrassv2_BatchAssociateClientDeviceWithCoreDevice(cfg aws.Config, client *greengrassv2.Client) {
	input := &greengrassv2.BatchAssociateClientDeviceWithCoreDeviceInput{
		// CoreDeviceThingName: *string, // Required
	}

	if len(_greengrassv2CoreDeviceThingName) > 0 {
		input.CoreDeviceThingName = aws.String(_greengrassv2CoreDeviceThingName)
	}
	if len(_greengrassv2Entries) > 0 {
		if err := assignInputField(input, "Entries", _greengrassv2Entries); err != nil {
			log.Errorf("invalid --entries: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchAssociateClientDeviceWithCoreDevice(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a list of client devices from a core device. After you
// disassociate a client device from a core device, the client device won't be able
// to use cloud discovery to retrieve the core device's connectivity information
// and certificates.
func greengrassv2_BatchDisassociateClientDeviceFromCoreDevice(cfg aws.Config, client *greengrassv2.Client) {
	input := &greengrassv2.BatchDisassociateClientDeviceFromCoreDeviceInput{
		// CoreDeviceThingName: *string, // Required
	}

	if len(_greengrassv2CoreDeviceThingName) > 0 {
		input.CoreDeviceThingName = aws.String(_greengrassv2CoreDeviceThingName)
	}
	if len(_greengrassv2Entries) > 0 {
		if err := assignInputField(input, "Entries", _greengrassv2Entries); err != nil {
			log.Errorf("invalid --entries: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchDisassociateClientDeviceFromCoreDevice(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels a deployment. This operation cancels the deployment for devices that
// haven't yet received it. If a device already received the deployment, this
// operation doesn't change anything for that device.
func greengrassv2_CancelDeployment(cfg aws.Config, client *greengrassv2.Client) {
	input := &greengrassv2.CancelDeploymentInput{
		// DeploymentId: *string, // Required
	}

	if len(_greengrassv2DeploymentId) > 0 {
		input.DeploymentId = aws.String(_greengrassv2DeploymentId)
	}

	if resp, err := client.CancelDeployment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a component. Components are software that run on Greengrass core
// devices. After you develop and test a component on your core device, you can use
// this operation to upload your component to IoT Greengrass. Then, you can deploy
// the component to other core devices.
//
// You can use this operation to do the following:
//
// - Create components from recipes
//
// # Create a component from a recipe, which is a file that defines the component's
//
// metadata, parameters, dependencies, lifecycle, artifacts, and platform
// capability. For more information, see [IoT Greengrass component recipe reference]in the IoT Greengrass V2 Developer
// Guide.
//
// # To create a component from a recipe, specify inlineRecipe when you call this
//
// operation.
//
// - Create components from Lambda functions
//
// Create a component from an Lambda function that runs on IoT Greengrass. This
//
// creates a recipe and artifacts from the Lambda function's deployment package.
// You can use this operation to migrate Lambda functions from IoT Greengrass V1 to
// IoT Greengrass V2.
//
// This function accepts Lambda functions in all supported versions of Python,
//
// Node.js, and Java runtimes. IoT Greengrass doesn't apply any additional
// restrictions on deprecated Lambda runtime versions.
//
// # To create a component from a Lambda function, specify lambdaFunction when you
//
// call this operation.
//
// IoT Greengrass currently supports Lambda functions on only Linux core devices.
//
// [IoT Greengrass component recipe reference]: https://docs.aws.amazon.com/greengrass/v2/developerguide/component-recipe-reference.html
func greengrassv2_CreateComponentVersion(cfg aws.Config, client *greengrassv2.Client) {
	input := &greengrassv2.CreateComponentVersionInput{}

	if len(_greengrassv2ClientToken) > 0 {
		input.ClientToken = aws.String(_greengrassv2ClientToken)
	}
	if len(_greengrassv2InlineRecipe) > 0 {
		if err := assignInputField(input, "InlineRecipe", _greengrassv2InlineRecipe); err != nil {
			log.Errorf("invalid --inline-recipe: %s", err.Error())
			return
		}
	}
	if len(_greengrassv2LambdaFunction) > 0 {
		if err := assignInputField(input, "LambdaFunction", _greengrassv2LambdaFunction); err != nil {
			log.Errorf("invalid --lambda-function: %s", err.Error())
			return
		}
	}
	if len(_greengrassv2Tags) > 0 {
		if err := assignInputField(input, "Tags", _greengrassv2Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateComponentVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a continuous deployment for a target, which is a Greengrass core device
// or group of core devices. When you add a new core device to a group of core
// devices that has a deployment, IoT Greengrass deploys that group's deployment to
// the new device.
//
// You can define one deployment for each target. When you create a new deployment
// for a target that has an existing deployment, you replace the previous
// deployment. IoT Greengrass applies the new deployment to the target devices.
//
// Every deployment has a revision number that indicates how many deployment
// revisions you define for a target. Use this operation to create a new revision
// of an existing deployment.
//
// For more information, see the [Create deployments] in the IoT Greengrass V2 Developer Guide.
//
// [Create deployments]: https://docs.aws.amazon.com/greengrass/v2/developerguide/create-deployments.html
func greengrassv2_CreateDeployment(cfg aws.Config, client *greengrassv2.Client) {
	input := &greengrassv2.CreateDeploymentInput{
		// TargetArn: *string, // Required
	}

	if len(_greengrassv2TargetArn) > 0 {
		input.TargetArn = aws.String(_greengrassv2TargetArn)
	}
	if len(_greengrassv2ClientToken) > 0 {
		input.ClientToken = aws.String(_greengrassv2ClientToken)
	}
	if len(_greengrassv2Components) > 0 {
		if err := assignInputField(input, "Components", _greengrassv2Components); err != nil {
			log.Errorf("invalid --components: %s", err.Error())
			return
		}
	}
	if len(_greengrassv2DeploymentName) > 0 {
		input.DeploymentName = aws.String(_greengrassv2DeploymentName)
	}
	if len(_greengrassv2DeploymentPolicies) > 0 {
		if err := assignInputField(input, "DeploymentPolicies", _greengrassv2DeploymentPolicies); err != nil {
			log.Errorf("invalid --deployment-policies: %s", err.Error())
			return
		}
	}
	if len(_greengrassv2IotJobConfiguration) > 0 {
		if err := assignInputField(input, "IotJobConfiguration", _greengrassv2IotJobConfiguration); err != nil {
			log.Errorf("invalid --iot-job-configuration: %s", err.Error())
			return
		}
	}
	if len(_greengrassv2ParentTargetArn) > 0 {
		input.ParentTargetArn = aws.String(_greengrassv2ParentTargetArn)
	}
	if len(_greengrassv2Tags) > 0 {
		if err := assignInputField(input, "Tags", _greengrassv2Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDeployment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a version of a component from IoT Greengrass.
// This operation deletes the component's recipe and artifacts. As a result,
// deployments that refer to this component version will fail. If you have
// deployments that use this component version, you can remove the component from
// the deployment or update the deployment to use a valid version.
func greengrassv2_DeleteComponent(cfg aws.Config, client *greengrassv2.Client) {
	input := &greengrassv2.DeleteComponentInput{
		// Arn: *string, // Required
	}

	if len(_greengrassv2Arn) > 0 {
		input.Arn = aws.String(_greengrassv2Arn)
	}

	if resp, err := client.DeleteComponent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Greengrass core device, which is an IoT thing. This operation removes
// the core device from the list of core devices. This operation doesn't delete the
// IoT thing. For more information about how to delete the IoT thing, see [DeleteThing]in the
// IoT API Reference.
//
// [DeleteThing]: https://docs.aws.amazon.com/iot/latest/apireference/API_DeleteThing.html
func greengrassv2_DeleteCoreDevice(cfg aws.Config, client *greengrassv2.Client) {
	input := &greengrassv2.DeleteCoreDeviceInput{
		// CoreDeviceThingName: *string, // Required
	}

	if len(_greengrassv2CoreDeviceThingName) > 0 {
		input.CoreDeviceThingName = aws.String(_greengrassv2CoreDeviceThingName)
	}

	if resp, err := client.DeleteCoreDevice(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a deployment. To delete an active deployment, you must first cancel it.
// For more information, see [CancelDeployment].
//
// Deleting a deployment doesn't affect core devices that run that deployment,
// because core devices store the deployment's configuration on the device.
// Additionally, core devices can roll back to a previous deployment that has been
// deleted.
//
// [CancelDeployment]: https://docs.aws.amazon.com/iot/latest/apireference/API_CancelDeployment.html
func greengrassv2_DeleteDeployment(cfg aws.Config, client *greengrassv2.Client) {
	input := &greengrassv2.DeleteDeploymentInput{
		// DeploymentId: *string, // Required
	}

	if len(_greengrassv2DeploymentId) > 0 {
		input.DeploymentId = aws.String(_greengrassv2DeploymentId)
	}

	if resp, err := client.DeleteDeployment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves metadata for a version of a component.
func greengrassv2_DescribeComponent(cfg aws.Config, client *greengrassv2.Client) {
	input := &greengrassv2.DescribeComponentInput{
		// Arn: *string, // Required
	}

	if len(_greengrassv2Arn) > 0 {
		input.Arn = aws.String(_greengrassv2Arn)
	}

	if resp, err := client.DescribeComponent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates the Greengrass service role from IoT Greengrass for your Amazon
// Web Services account in this Amazon Web Services Region. Without a service role,
// IoT Greengrass can't verify the identity of client devices or manage core device
// connectivity information. For more information, see [Greengrass service role]in the IoT Greengrass
// Version 2 Developer Guide.
//
// [Greengrass service role]: https://docs.aws.amazon.com/greengrass/v2/developerguide/greengrass-service-role.html
func greengrassv2_DisassociateServiceRoleFromAccount(cfg aws.Config, client *greengrassv2.Client) {
	input := &greengrassv2.DisassociateServiceRoleFromAccountInput{}

	if resp, err := client.DisassociateServiceRoleFromAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the recipe for a version of a component.
func greengrassv2_GetComponent(cfg aws.Config, client *greengrassv2.Client) {
	input := &greengrassv2.GetComponentInput{
		// Arn: *string, // Required
	}

	if len(_greengrassv2Arn) > 0 {
		input.Arn = aws.String(_greengrassv2Arn)
	}
	if len(_greengrassv2RecipeOutputFormat) > 0 {
		if err := assignInputField(input, "RecipeOutputFormat", _greengrassv2RecipeOutputFormat); err != nil {
			log.Errorf("invalid --recipe-output-format: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetComponent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the pre-signed URL to download a public or a Lambda component artifact.
// Core devices call this operation to identify the URL that they can use to
// download an artifact to install.
func greengrassv2_GetComponentVersionArtifact(cfg aws.Config, client *greengrassv2.Client) {
	input := &greengrassv2.GetComponentVersionArtifactInput{
		// Arn: *string, // Required
		// ArtifactName: *string, // Required
	}

	if len(_greengrassv2Arn) > 0 {
		input.Arn = aws.String(_greengrassv2Arn)
	}
	if len(_greengrassv2ArtifactName) > 0 {
		input.ArtifactName = aws.String(_greengrassv2ArtifactName)
	}
	if len(_greengrassv2IotEndpointType) > 0 {
		if err := assignInputField(input, "IotEndpointType", _greengrassv2IotEndpointType); err != nil {
			log.Errorf("invalid --iot-endpoint-type: %s", err.Error())
			return
		}
	}
	if len(_greengrassv2S3EndpointType) > 0 {
		if err := assignInputField(input, "S3EndpointType", _greengrassv2S3EndpointType); err != nil {
			log.Errorf("invalid --s3-endpoint-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetComponentVersionArtifact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves connectivity information for a Greengrass core device.
// Connectivity information includes endpoints and ports where client devices can
// connect to an MQTT broker on the core device. When a client device calls the [IoT Greengrass discovery API],
// IoT Greengrass returns connectivity information for all of the core devices
// where the client device can connect. For more information, see [Connect client devices to core devices]in the IoT
// Greengrass Version 2 Developer Guide.
//
// [Connect client devices to core devices]: https://docs.aws.amazon.com/greengrass/v2/developerguide/connect-client-devices.html
// [IoT Greengrass discovery API]: https://docs.aws.amazon.com/greengrass/v2/developerguide/greengrass-discover-api.html
func greengrassv2_GetConnectivityInfo(cfg aws.Config, client *greengrassv2.Client) {
	input := &greengrassv2.GetConnectivityInfoInput{
		// ThingName: *string, // Required
	}

	if len(_greengrassv2ThingName) > 0 {
		input.ThingName = aws.String(_greengrassv2ThingName)
	}

	if resp, err := client.GetConnectivityInfo(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves metadata for a Greengrass core device.
// IoT Greengrass relies on individual devices to send status updates to the
// Amazon Web Services Cloud. If the IoT Greengrass Core software isn't running on
// the device, or if device isn't connected to the Amazon Web Services Cloud, then
// the reported status of that device might not reflect its current status. The
// status timestamp indicates when the device status was last updated.
//
// Core devices send status updates at the following times:
//
// - When the IoT Greengrass Core software starts
//
// - When the core device receives a deployment from the Amazon Web Services
// Cloud
//
// - When the status of any component on the core device becomes BROKEN
//
// - At a [regular interval that you can configure], which defaults to 24 hours
//
// - For IoT Greengrass Core v2.7.0, the core device sends status updates upon
// local deployment and cloud deployment
//
// [regular interval that you can configure]: https://docs.aws.amazon.com/greengrass/v2/developerguide/greengrass-nucleus-component.html#greengrass-nucleus-component-configuration-fss
func greengrassv2_GetCoreDevice(cfg aws.Config, client *greengrassv2.Client) {
	input := &greengrassv2.GetCoreDeviceInput{
		// CoreDeviceThingName: *string, // Required
	}

	if len(_greengrassv2CoreDeviceThingName) > 0 {
		input.CoreDeviceThingName = aws.String(_greengrassv2CoreDeviceThingName)
	}

	if resp, err := client.GetCoreDevice(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a deployment. Deployments define the components that run on Greengrass
// core devices.
func greengrassv2_GetDeployment(cfg aws.Config, client *greengrassv2.Client) {
	input := &greengrassv2.GetDeploymentInput{
		// DeploymentId: *string, // Required
	}

	if len(_greengrassv2DeploymentId) > 0 {
		input.DeploymentId = aws.String(_greengrassv2DeploymentId)
	}

	if resp, err := client.GetDeployment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the service role associated with IoT Greengrass for your Amazon Web
// Services account in this Amazon Web Services Region. IoT Greengrass uses this
// role to verify the identity of client devices and manage core device
// connectivity information. For more information, see [Greengrass service role]in the IoT Greengrass
// Version 2 Developer Guide.
//
// [Greengrass service role]: https://docs.aws.amazon.com/greengrass/v2/developerguide/greengrass-service-role.html
func greengrassv2_GetServiceRoleForAccount(cfg aws.Config, client *greengrassv2.Client) {
	input := &greengrassv2.GetServiceRoleForAccountInput{}

	if resp, err := client.GetServiceRoleForAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a paginated list of client devices that are associated with a core
// device.
func greengrassv2_ListClientDevicesAssociatedWithCoreDevice(cfg aws.Config, client *greengrassv2.Client) {
	input := &greengrassv2.ListClientDevicesAssociatedWithCoreDeviceInput{
		// CoreDeviceThingName: *string, // Required
	}

	if len(_greengrassv2CoreDeviceThingName) > 0 {
		input.CoreDeviceThingName = aws.String(_greengrassv2CoreDeviceThingName)
	}
	if len(_greengrassv2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _greengrassv2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_greengrassv2NextToken) > 0 {
		input.NextToken = aws.String(_greengrassv2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListClientDevicesAssociatedWithCoreDevice(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*greengrassv2.ListClientDevicesAssociatedWithCoreDeviceOutput
	p := greengrassv2.NewListClientDevicesAssociatedWithCoreDevicePaginator(client, input)
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

// Retrieves a paginated list of all versions for a component. Greater versions
// are listed first.
func greengrassv2_ListComponentVersions(cfg aws.Config, client *greengrassv2.Client) {
	input := &greengrassv2.ListComponentVersionsInput{
		// Arn: *string, // Required
	}

	if len(_greengrassv2Arn) > 0 {
		input.Arn = aws.String(_greengrassv2Arn)
	}
	if len(_greengrassv2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _greengrassv2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_greengrassv2NextToken) > 0 {
		input.NextToken = aws.String(_greengrassv2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListComponentVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*greengrassv2.ListComponentVersionsOutput
	p := greengrassv2.NewListComponentVersionsPaginator(client, input)
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

// Retrieves a paginated list of component summaries. This list includes
// components that you have permission to view.
func greengrassv2_ListComponents(cfg aws.Config, client *greengrassv2.Client) {
	input := &greengrassv2.ListComponentsInput{}

	if len(_greengrassv2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _greengrassv2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_greengrassv2NextToken) > 0 {
		input.NextToken = aws.String(_greengrassv2NextToken)
	}
	if len(_greengrassv2Scope) > 0 {
		if err := assignInputField(input, "Scope", _greengrassv2Scope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListComponents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*greengrassv2.ListComponentsOutput
	p := greengrassv2.NewListComponentsPaginator(client, input)
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

// Retrieves a paginated list of Greengrass core devices.
// IoT Greengrass relies on individual devices to send status updates to the
// Amazon Web Services Cloud. If the IoT Greengrass Core software isn't running on
// the device, or if device isn't connected to the Amazon Web Services Cloud, then
// the reported status of that device might not reflect its current status. The
// status timestamp indicates when the device status was last updated.
//
// Core devices send status updates at the following times:
//
// - When the IoT Greengrass Core software starts
//
// - When the core device receives a deployment from the Amazon Web Services
// Cloud
//
// - For Greengrass nucleus 2.12.2 and earlier, the core device sends status
// updates when the status of any component on the core device becomes ERRORED or
// BROKEN .
//
// - For Greengrass nucleus 2.12.3 and later, the core device sends status
// updates when the status of any component on the core device becomes ERRORED ,
// BROKEN , RUNNING , or FINISHED .
//
// - At a [regular interval that you can configure], which defaults to 24 hours
//
// - For IoT Greengrass Core v2.7.0, the core device sends status updates upon
// local deployment and cloud deployment
//
// [regular interval that you can configure]: https://docs.aws.amazon.com/greengrass/v2/developerguide/greengrass-nucleus-component.html#greengrass-nucleus-component-configuration-fss
func greengrassv2_ListCoreDevices(cfg aws.Config, client *greengrassv2.Client) {
	input := &greengrassv2.ListCoreDevicesInput{}

	if len(_greengrassv2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _greengrassv2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_greengrassv2NextToken) > 0 {
		input.NextToken = aws.String(_greengrassv2NextToken)
	}
	if len(_greengrassv2Runtime) > 0 {
		input.Runtime = aws.String(_greengrassv2Runtime)
	}
	if len(_greengrassv2Status) > 0 {
		if err := assignInputField(input, "Status", _greengrassv2Status); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_greengrassv2ThingGroupArn) > 0 {
		input.ThingGroupArn = aws.String(_greengrassv2ThingGroupArn)
	}

	if disablePaginator() {
		if resp, err := client.ListCoreDevices(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*greengrassv2.ListCoreDevicesOutput
	p := greengrassv2.NewListCoreDevicesPaginator(client, input)
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

// Retrieves a paginated list of deployments.
func greengrassv2_ListDeployments(cfg aws.Config, client *greengrassv2.Client) {
	input := &greengrassv2.ListDeploymentsInput{}

	if len(_greengrassv2HistoryFilter) > 0 {
		if err := assignInputField(input, "HistoryFilter", _greengrassv2HistoryFilter); err != nil {
			log.Errorf("invalid --history-filter: %s", err.Error())
			return
		}
	}
	if len(_greengrassv2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _greengrassv2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_greengrassv2NextToken) > 0 {
		input.NextToken = aws.String(_greengrassv2NextToken)
	}
	if len(_greengrassv2ParentTargetArn) > 0 {
		input.ParentTargetArn = aws.String(_greengrassv2ParentTargetArn)
	}
	if len(_greengrassv2TargetArn) > 0 {
		input.TargetArn = aws.String(_greengrassv2TargetArn)
	}

	if disablePaginator() {
		if resp, err := client.ListDeployments(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*greengrassv2.ListDeploymentsOutput
	p := greengrassv2.NewListDeploymentsPaginator(client, input)
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

// Retrieves a paginated list of deployment jobs that IoT Greengrass sends to
// Greengrass core devices.
func greengrassv2_ListEffectiveDeployments(cfg aws.Config, client *greengrassv2.Client) {
	input := &greengrassv2.ListEffectiveDeploymentsInput{
		// CoreDeviceThingName: *string, // Required
	}

	if len(_greengrassv2CoreDeviceThingName) > 0 {
		input.CoreDeviceThingName = aws.String(_greengrassv2CoreDeviceThingName)
	}
	if len(_greengrassv2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _greengrassv2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_greengrassv2NextToken) > 0 {
		input.NextToken = aws.String(_greengrassv2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEffectiveDeployments(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*greengrassv2.ListEffectiveDeploymentsOutput
	p := greengrassv2.NewListEffectiveDeploymentsPaginator(client, input)
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

// Retrieves a paginated list of the components that a Greengrass core device
// runs. By default, this list doesn't include components that are deployed as
// dependencies of other components. To include dependencies in the response, set
// the topologyFilter parameter to ALL .
//
// IoT Greengrass relies on individual devices to send status updates to the
// Amazon Web Services Cloud. If the IoT Greengrass Core software isn't running on
// the device, or if device isn't connected to the Amazon Web Services Cloud, then
// the reported status of that device might not reflect its current status. The
// status timestamp indicates when the device status was last updated.
//
// Core devices send status updates at the following times:
//
// - When the IoT Greengrass Core software starts
//
// - When the core device receives a deployment from the Amazon Web Services
// Cloud
//
// - When the status of any component on the core device becomes BROKEN
//
// - At a [regular interval that you can configure], which defaults to 24 hours
//
// - For IoT Greengrass Core v2.7.0, the core device sends status updates upon
// local deployment and cloud deployment
//
// [regular interval that you can configure]: https://docs.aws.amazon.com/greengrass/v2/developerguide/greengrass-nucleus-component.html#greengrass-nucleus-component-configuration-fss
func greengrassv2_ListInstalledComponents(cfg aws.Config, client *greengrassv2.Client) {
	input := &greengrassv2.ListInstalledComponentsInput{
		// CoreDeviceThingName: *string, // Required
	}

	if len(_greengrassv2CoreDeviceThingName) > 0 {
		input.CoreDeviceThingName = aws.String(_greengrassv2CoreDeviceThingName)
	}
	if len(_greengrassv2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _greengrassv2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_greengrassv2NextToken) > 0 {
		input.NextToken = aws.String(_greengrassv2NextToken)
	}
	if len(_greengrassv2TopologyFilter) > 0 {
		if err := assignInputField(input, "TopologyFilter", _greengrassv2TopologyFilter); err != nil {
			log.Errorf("invalid --topology-filter: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListInstalledComponents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*greengrassv2.ListInstalledComponentsOutput
	p := greengrassv2.NewListInstalledComponentsPaginator(client, input)
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

// Retrieves the list of tags for an IoT Greengrass resource.
func greengrassv2_ListTagsForResource(cfg aws.Config, client *greengrassv2.Client) {
	input := &greengrassv2.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_greengrassv2ResourceArn) > 0 {
		input.ResourceArn = aws.String(_greengrassv2ResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of components that meet the component, version, and platform
// requirements of a deployment. Greengrass core devices call this operation when
// they receive a deployment to identify the components to install.
//
// This operation identifies components that meet all dependency requirements for
// a deployment. If the requirements conflict, then this operation returns an error
// and the deployment fails. For example, this occurs if component A requires
// version >2.0.0 and component B requires version <2.0.0 of a component
// dependency.
//
// When you specify the component candidates to resolve, IoT Greengrass compares
// each component's digest from the core device with the component's digest in the
// Amazon Web Services Cloud. If the digests don't match, then IoT Greengrass
// specifies to use the version from the Amazon Web Services Cloud.
//
// To use this operation, you must use the data plane API endpoint and
// authenticate with an IoT device certificate. For more information, see [IoT Greengrass endpoints and quotas].
//
// [IoT Greengrass endpoints and quotas]: https://docs.aws.amazon.com/general/latest/gr/greengrass.html
func greengrassv2_ResolveComponentCandidates(cfg aws.Config, client *greengrassv2.Client) {
	input := &greengrassv2.ResolveComponentCandidatesInput{}

	if len(_greengrassv2ComponentCandidates) > 0 {
		if err := assignInputField(input, "ComponentCandidates", _greengrassv2ComponentCandidates); err != nil {
			log.Errorf("invalid --component-candidates: %s", err.Error())
			return
		}
	}
	if len(_greengrassv2Platform) > 0 {
		if err := assignInputField(input, "Platform", _greengrassv2Platform); err != nil {
			log.Errorf("invalid --platform: %s", err.Error())
			return
		}
	}

	if resp, err := client.ResolveComponentCandidates(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds tags to an IoT Greengrass resource. If a tag already exists for the
// resource, this operation updates the tag's value.
func greengrassv2_TagResource(cfg aws.Config, client *greengrassv2.Client) {
	input := &greengrassv2.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_greengrassv2ResourceArn) > 0 {
		input.ResourceArn = aws.String(_greengrassv2ResourceArn)
	}
	if len(_greengrassv2Tags) > 0 {
		if err := assignInputField(input, "Tags", _greengrassv2Tags); err != nil {
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

// Removes a tag from an IoT Greengrass resource.
func greengrassv2_UntagResource(cfg aws.Config, client *greengrassv2.Client) {
	input := &greengrassv2.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_greengrassv2ResourceArn) > 0 {
		input.ResourceArn = aws.String(_greengrassv2ResourceArn)
	}
	if len(_greengrassv2TagKeys) > 0 {
		input.TagKeys = append([]string(nil), _greengrassv2TagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates connectivity information for a Greengrass core device.
// Connectivity information includes endpoints and ports where client devices can
// connect to an MQTT broker on the core device. When a client device calls the [IoT Greengrass discovery API],
// IoT Greengrass returns connectivity information for all of the core devices
// where the client device can connect. For more information, see [Connect client devices to core devices]in the IoT
// Greengrass Version 2 Developer Guide.
//
// [Connect client devices to core devices]: https://docs.aws.amazon.com/greengrass/v2/developerguide/connect-client-devices.html
// [IoT Greengrass discovery API]: https://docs.aws.amazon.com/greengrass/v2/developerguide/greengrass-discover-api.html
func greengrassv2_UpdateConnectivityInfo(cfg aws.Config, client *greengrassv2.Client) {
	input := &greengrassv2.UpdateConnectivityInfoInput{
		// ConnectivityInfo: []types.ConnectivityInfo, // Required
		// ThingName: *string, // Required
	}

	if len(_greengrassv2ConnectivityInfo) > 0 {
		if err := assignInputField(input, "ConnectivityInfo", _greengrassv2ConnectivityInfo); err != nil {
			log.Errorf("invalid --connectivity-info: %s", err.Error())
			return
		}
	}
	if len(_greengrassv2ThingName) > 0 {
		input.ThingName = aws.String(_greengrassv2ThingName)
	}

	if resp, err := client.UpdateConnectivityInfo(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_greengrassv2Cmd)
	_greengrassv2Cmd.Flags().SortFlags = false

	_greengrassv2Cmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_greengrassv2Cmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_greengrassv2Cmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_greengrassv2Cmd.Flags().StringVarP(&_greengrassv2Arn, "arn", "", "", "ARN")
	_greengrassv2Cmd.Flags().StringVarP(&_greengrassv2ArtifactName, "artifact-name", "", "", "Artifact Name")
	_greengrassv2Cmd.Flags().StringVarP(&_greengrassv2ClientToken, "client-token", "", "", "Client Token")
	_greengrassv2Cmd.Flags().StringVarP(&_greengrassv2ComponentCandidates, "component-candidates", "", "", "Component Candidates")
	_greengrassv2Cmd.Flags().StringVarP(&_greengrassv2Components, "components", "", "", "Components")
	_greengrassv2Cmd.Flags().StringVarP(&_greengrassv2ConnectivityInfo, "connectivity-info", "", "", "Connectivity Info")
	_greengrassv2Cmd.Flags().StringVarP(&_greengrassv2CoreDeviceThingName, "core-device-thing-name", "", "", "Core Device Thing Name")
	_greengrassv2Cmd.Flags().StringVarP(&_greengrassv2DeploymentId, "deployment-id", "", "", "Deployment ID")
	_greengrassv2Cmd.Flags().StringVarP(&_greengrassv2DeploymentName, "deployment-name", "", "", "Deployment Name")
	_greengrassv2Cmd.Flags().StringVarP(&_greengrassv2DeploymentPolicies, "deployment-policies", "", "", "Deployment Policies")
	_greengrassv2Cmd.Flags().StringVarP(&_greengrassv2Entries, "entries", "", "", "Entries")
	_greengrassv2Cmd.Flags().StringVarP(&_greengrassv2HistoryFilter, "history-filter", "", "", "History Filter")
	_greengrassv2Cmd.Flags().StringVarP(&_greengrassv2InlineRecipe, "inline-recipe", "", "", "Inline Recipe")
	_greengrassv2Cmd.Flags().StringVarP(&_greengrassv2IotEndpointType, "iot-endpoint-type", "", "", "Iot Endpoint Type")
	_greengrassv2Cmd.Flags().StringVarP(&_greengrassv2IotJobConfiguration, "iot-job-configuration", "", "", "Iot Job Configuration")
	_greengrassv2Cmd.Flags().StringVarP(&_greengrassv2LambdaFunction, "lambda-function", "", "", "Lambda Function")
	_greengrassv2Cmd.Flags().StringVarP(&_greengrassv2MaxResults, "max-results", "", "", "Max Results")
	_greengrassv2Cmd.Flags().StringVarP(&_greengrassv2NextToken, "next-token", "", "", "Next Token")
	_greengrassv2Cmd.Flags().StringVarP(&_greengrassv2ParentTargetArn, "parent-target-arn", "", "", "Parent Target ARN")
	_greengrassv2Cmd.Flags().StringVarP(&_greengrassv2Platform, "platform", "", "", "Platform")
	_greengrassv2Cmd.Flags().StringVarP(&_greengrassv2RecipeOutputFormat, "recipe-output-format", "", "", "Recipe Output Format")
	_greengrassv2Cmd.Flags().StringVarP(&_greengrassv2ResourceArn, "resource-arn", "", "", "Resource ARN")
	_greengrassv2Cmd.Flags().StringVarP(&_greengrassv2RoleArn, "role-arn", "", "", "Role ARN")
	_greengrassv2Cmd.Flags().StringVarP(&_greengrassv2Runtime, "runtime", "", "", "Runtime")
	_greengrassv2Cmd.Flags().StringVarP(&_greengrassv2S3EndpointType, "s3-endpoint-type", "", "", "S3 Endpoint Type")
	_greengrassv2Cmd.Flags().StringVarP(&_greengrassv2Scope, "scope", "", "", "Scope")
	_greengrassv2Cmd.Flags().StringVarP(&_greengrassv2Status, "status", "", "", "Status")
	_greengrassv2Cmd.Flags().StringSliceVarP(&_greengrassv2TagKeys, "tag-keys", "", nil, "Tag Keys")
	_greengrassv2Cmd.Flags().StringVarP(&_greengrassv2Tags, "tags", "", "", "Tags")
	_greengrassv2Cmd.Flags().StringVarP(&_greengrassv2TargetArn, "target-arn", "", "", "Target ARN")
	_greengrassv2Cmd.Flags().StringVarP(&_greengrassv2ThingGroupArn, "thing-group-arn", "", "", "Thing Group ARN")
	_greengrassv2Cmd.Flags().StringVarP(&_greengrassv2ThingName, "thing-name", "", "", "Thing Name")
	_greengrassv2Cmd.Flags().StringVarP(&_greengrassv2TopologyFilter, "topology-filter", "", "", "Topology Filter")

	_greengrassv2Cmd.Flags().BoolVarP(&_greengrassv2AssociateServiceRoleToAccount, "associate-service-role-to-account", "", false, "Associate Service Role To Account")
	_greengrassv2Cmd.Flags().BoolVarP(&_greengrassv2BatchAssociateClientDeviceWithCoreDevice, "batch-associate-client-device-with-core-device", "", false, "Batch Associate Client Device With Core Device")
	_greengrassv2Cmd.Flags().BoolVarP(&_greengrassv2BatchDisassociateClientDeviceFromCoreDevice, "batch-disassociate-client-device-from-core-device", "", false, "Batch Disassociate Client Device From Core Device")
	_greengrassv2Cmd.Flags().BoolVarP(&_greengrassv2CancelDeployment, "cancel-deployment", "", false, "Cancel Deployment")
	_greengrassv2Cmd.Flags().BoolVarP(&_greengrassv2CreateComponentVersion, "create-component-version", "", false, "Create Component Version")
	_greengrassv2Cmd.Flags().BoolVarP(&_greengrassv2CreateDeployment, "create-deployment", "", false, "Create Deployment")
	_greengrassv2Cmd.Flags().BoolVarP(&_greengrassv2DeleteComponent, "delete-component", "", false, "Delete Component")
	_greengrassv2Cmd.Flags().BoolVarP(&_greengrassv2DeleteCoreDevice, "delete-core-device", "", false, "Delete Core Device")
	_greengrassv2Cmd.Flags().BoolVarP(&_greengrassv2DeleteDeployment, "delete-deployment", "", false, "Delete Deployment")
	_greengrassv2Cmd.Flags().BoolVarP(&_greengrassv2DescribeComponent, "describe-component", "", false, "Describe Component")
	_greengrassv2Cmd.Flags().BoolVarP(&_greengrassv2DisassociateServiceRoleFromAccount, "disassociate-service-role-from-account", "", false, "Disassociate Service Role From Account")
	_greengrassv2Cmd.Flags().BoolVarP(&_greengrassv2GetComponent, "get-component", "", false, "Get Component")
	_greengrassv2Cmd.Flags().BoolVarP(&_greengrassv2GetComponentVersionArtifact, "get-component-version-artifact", "", false, "Get Component Version Artifact")
	_greengrassv2Cmd.Flags().BoolVarP(&_greengrassv2GetConnectivityInfo, "get-connectivity-info", "", false, "Get Connectivity Info")
	_greengrassv2Cmd.Flags().BoolVarP(&_greengrassv2GetCoreDevice, "get-core-device", "", false, "Get Core Device")
	_greengrassv2Cmd.Flags().BoolVarP(&_greengrassv2GetDeployment, "get-deployment", "", false, "Get Deployment")
	_greengrassv2Cmd.Flags().BoolVarP(&_greengrassv2GetServiceRoleForAccount, "get-service-role-for-account", "", false, "Get Service Role For Account")
	_greengrassv2Cmd.Flags().BoolVarP(&_greengrassv2ListClientDevicesAssociatedWithCoreDevice, "list-client-devices-associated-with-core-device", "", false, "List Client Devices Associated With Core Device")
	_greengrassv2Cmd.Flags().BoolVarP(&_greengrassv2ListComponentVersions, "list-component-versions", "", false, "List Component Versions")
	_greengrassv2Cmd.Flags().BoolVarP(&_greengrassv2ListComponents, "list-components", "", false, "List Components")
	_greengrassv2Cmd.Flags().BoolVarP(&_greengrassv2ListCoreDevices, "list-core-devices", "", false, "List Core Devices")
	_greengrassv2Cmd.Flags().BoolVarP(&_greengrassv2ListDeployments, "list-deployments", "", false, "List Deployments")
	_greengrassv2Cmd.Flags().BoolVarP(&_greengrassv2ListEffectiveDeployments, "list-effective-deployments", "", false, "List Effective Deployments")
	_greengrassv2Cmd.Flags().BoolVarP(&_greengrassv2ListInstalledComponents, "list-installed-components", "", false, "List Installed Components")
	_greengrassv2Cmd.Flags().BoolVarP(&_greengrassv2ListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_greengrassv2Cmd.Flags().BoolVarP(&_greengrassv2ResolveComponentCandidates, "resolve-component-candidates", "", false, "Resolve Component Candidates")
	_greengrassv2Cmd.Flags().BoolVarP(&_greengrassv2TagResource, "tag-resource", "", false, "Tag Resource")
	_greengrassv2Cmd.Flags().BoolVarP(&_greengrassv2UntagResource, "untag-resource", "", false, "Untag Resource")
	_greengrassv2Cmd.Flags().BoolVarP(&_greengrassv2UpdateConnectivityInfo, "update-connectivity-info", "", false, "Update Connectivity Info")

}
