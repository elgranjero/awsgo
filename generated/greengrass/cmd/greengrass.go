package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/greengrass"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// greengrassCmd represents the greengrass command
var _greengrassCmd = &cobra.Command{
	Use:   "greengrass",
	Short: "AWS greengrass CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := greengrass.NewFromConfig(cfg)
		if _greengrassAssociateRoleToGroup {
			greengrass_AssociateRoleToGroup(cfg, client)
			return
		}
		if _greengrassAssociateServiceRoleToAccount {
			greengrass_AssociateServiceRoleToAccount(cfg, client)
			return
		}
		if _greengrassCreateConnectorDefinition {
			greengrass_CreateConnectorDefinition(cfg, client)
			return
		}
		if _greengrassCreateConnectorDefinitionVersion {
			greengrass_CreateConnectorDefinitionVersion(cfg, client)
			return
		}
		if _greengrassCreateCoreDefinition {
			greengrass_CreateCoreDefinition(cfg, client)
			return
		}
		if _greengrassCreateCoreDefinitionVersion {
			greengrass_CreateCoreDefinitionVersion(cfg, client)
			return
		}
		if _greengrassCreateDeployment {
			greengrass_CreateDeployment(cfg, client)
			return
		}
		if _greengrassCreateDeviceDefinition {
			greengrass_CreateDeviceDefinition(cfg, client)
			return
		}
		if _greengrassCreateDeviceDefinitionVersion {
			greengrass_CreateDeviceDefinitionVersion(cfg, client)
			return
		}
		if _greengrassCreateFunctionDefinition {
			greengrass_CreateFunctionDefinition(cfg, client)
			return
		}
		if _greengrassCreateFunctionDefinitionVersion {
			greengrass_CreateFunctionDefinitionVersion(cfg, client)
			return
		}
		if _greengrassCreateGroup {
			greengrass_CreateGroup(cfg, client)
			return
		}
		if _greengrassCreateGroupCertificateAuthority {
			greengrass_CreateGroupCertificateAuthority(cfg, client)
			return
		}
		if _greengrassCreateGroupVersion {
			greengrass_CreateGroupVersion(cfg, client)
			return
		}
		if _greengrassCreateLoggerDefinition {
			greengrass_CreateLoggerDefinition(cfg, client)
			return
		}
		if _greengrassCreateLoggerDefinitionVersion {
			greengrass_CreateLoggerDefinitionVersion(cfg, client)
			return
		}
		if _greengrassCreateResourceDefinition {
			greengrass_CreateResourceDefinition(cfg, client)
			return
		}
		if _greengrassCreateResourceDefinitionVersion {
			greengrass_CreateResourceDefinitionVersion(cfg, client)
			return
		}
		if _greengrassCreateSoftwareUpdateJob {
			greengrass_CreateSoftwareUpdateJob(cfg, client)
			return
		}
		if _greengrassCreateSubscriptionDefinition {
			greengrass_CreateSubscriptionDefinition(cfg, client)
			return
		}
		if _greengrassCreateSubscriptionDefinitionVersion {
			greengrass_CreateSubscriptionDefinitionVersion(cfg, client)
			return
		}
		if _greengrassDeleteConnectorDefinition {
			greengrass_DeleteConnectorDefinition(cfg, client)
			return
		}
		if _greengrassDeleteCoreDefinition {
			greengrass_DeleteCoreDefinition(cfg, client)
			return
		}
		if _greengrassDeleteDeviceDefinition {
			greengrass_DeleteDeviceDefinition(cfg, client)
			return
		}
		if _greengrassDeleteFunctionDefinition {
			greengrass_DeleteFunctionDefinition(cfg, client)
			return
		}
		if _greengrassDeleteGroup {
			greengrass_DeleteGroup(cfg, client)
			return
		}
		if _greengrassDeleteLoggerDefinition {
			greengrass_DeleteLoggerDefinition(cfg, client)
			return
		}
		if _greengrassDeleteResourceDefinition {
			greengrass_DeleteResourceDefinition(cfg, client)
			return
		}
		if _greengrassDeleteSubscriptionDefinition {
			greengrass_DeleteSubscriptionDefinition(cfg, client)
			return
		}
		if _greengrassDisassociateRoleFromGroup {
			greengrass_DisassociateRoleFromGroup(cfg, client)
			return
		}
		if _greengrassDisassociateServiceRoleFromAccount {
			greengrass_DisassociateServiceRoleFromAccount(cfg, client)
			return
		}
		if _greengrassGetAssociatedRole {
			greengrass_GetAssociatedRole(cfg, client)
			return
		}
		if _greengrassGetBulkDeploymentStatus {
			greengrass_GetBulkDeploymentStatus(cfg, client)
			return
		}
		if _greengrassGetConnectivityInfo {
			greengrass_GetConnectivityInfo(cfg, client)
			return
		}
		if _greengrassGetConnectorDefinition {
			greengrass_GetConnectorDefinition(cfg, client)
			return
		}
		if _greengrassGetConnectorDefinitionVersion {
			greengrass_GetConnectorDefinitionVersion(cfg, client)
			return
		}
		if _greengrassGetCoreDefinition {
			greengrass_GetCoreDefinition(cfg, client)
			return
		}
		if _greengrassGetCoreDefinitionVersion {
			greengrass_GetCoreDefinitionVersion(cfg, client)
			return
		}
		if _greengrassGetDeploymentStatus {
			greengrass_GetDeploymentStatus(cfg, client)
			return
		}
		if _greengrassGetDeviceDefinition {
			greengrass_GetDeviceDefinition(cfg, client)
			return
		}
		if _greengrassGetDeviceDefinitionVersion {
			greengrass_GetDeviceDefinitionVersion(cfg, client)
			return
		}
		if _greengrassGetFunctionDefinition {
			greengrass_GetFunctionDefinition(cfg, client)
			return
		}
		if _greengrassGetFunctionDefinitionVersion {
			greengrass_GetFunctionDefinitionVersion(cfg, client)
			return
		}
		if _greengrassGetGroup {
			greengrass_GetGroup(cfg, client)
			return
		}
		if _greengrassGetGroupCertificateAuthority {
			greengrass_GetGroupCertificateAuthority(cfg, client)
			return
		}
		if _greengrassGetGroupCertificateConfiguration {
			greengrass_GetGroupCertificateConfiguration(cfg, client)
			return
		}
		if _greengrassGetGroupVersion {
			greengrass_GetGroupVersion(cfg, client)
			return
		}
		if _greengrassGetLoggerDefinition {
			greengrass_GetLoggerDefinition(cfg, client)
			return
		}
		if _greengrassGetLoggerDefinitionVersion {
			greengrass_GetLoggerDefinitionVersion(cfg, client)
			return
		}
		if _greengrassGetResourceDefinition {
			greengrass_GetResourceDefinition(cfg, client)
			return
		}
		if _greengrassGetResourceDefinitionVersion {
			greengrass_GetResourceDefinitionVersion(cfg, client)
			return
		}
		if _greengrassGetServiceRoleForAccount {
			greengrass_GetServiceRoleForAccount(cfg, client)
			return
		}
		if _greengrassGetSubscriptionDefinition {
			greengrass_GetSubscriptionDefinition(cfg, client)
			return
		}
		if _greengrassGetSubscriptionDefinitionVersion {
			greengrass_GetSubscriptionDefinitionVersion(cfg, client)
			return
		}
		if _greengrassGetThingRuntimeConfiguration {
			greengrass_GetThingRuntimeConfiguration(cfg, client)
			return
		}
		if _greengrassListBulkDeploymentDetailedReports {
			greengrass_ListBulkDeploymentDetailedReports(cfg, client)
			return
		}
		if _greengrassListBulkDeployments {
			greengrass_ListBulkDeployments(cfg, client)
			return
		}
		if _greengrassListConnectorDefinitionVersions {
			greengrass_ListConnectorDefinitionVersions(cfg, client)
			return
		}
		if _greengrassListConnectorDefinitions {
			greengrass_ListConnectorDefinitions(cfg, client)
			return
		}
		if _greengrassListCoreDefinitionVersions {
			greengrass_ListCoreDefinitionVersions(cfg, client)
			return
		}
		if _greengrassListCoreDefinitions {
			greengrass_ListCoreDefinitions(cfg, client)
			return
		}
		if _greengrassListDeployments {
			greengrass_ListDeployments(cfg, client)
			return
		}
		if _greengrassListDeviceDefinitionVersions {
			greengrass_ListDeviceDefinitionVersions(cfg, client)
			return
		}
		if _greengrassListDeviceDefinitions {
			greengrass_ListDeviceDefinitions(cfg, client)
			return
		}
		if _greengrassListFunctionDefinitionVersions {
			greengrass_ListFunctionDefinitionVersions(cfg, client)
			return
		}
		if _greengrassListFunctionDefinitions {
			greengrass_ListFunctionDefinitions(cfg, client)
			return
		}
		if _greengrassListGroupCertificateAuthorities {
			greengrass_ListGroupCertificateAuthorities(cfg, client)
			return
		}
		if _greengrassListGroupVersions {
			greengrass_ListGroupVersions(cfg, client)
			return
		}
		if _greengrassListGroups {
			greengrass_ListGroups(cfg, client)
			return
		}
		if _greengrassListLoggerDefinitionVersions {
			greengrass_ListLoggerDefinitionVersions(cfg, client)
			return
		}
		if _greengrassListLoggerDefinitions {
			greengrass_ListLoggerDefinitions(cfg, client)
			return
		}
		if _greengrassListResourceDefinitionVersions {
			greengrass_ListResourceDefinitionVersions(cfg, client)
			return
		}
		if _greengrassListResourceDefinitions {
			greengrass_ListResourceDefinitions(cfg, client)
			return
		}
		if _greengrassListSubscriptionDefinitionVersions {
			greengrass_ListSubscriptionDefinitionVersions(cfg, client)
			return
		}
		if _greengrassListSubscriptionDefinitions {
			greengrass_ListSubscriptionDefinitions(cfg, client)
			return
		}
		if _greengrassListTagsForResource {
			greengrass_ListTagsForResource(cfg, client)
			return
		}
		if _greengrassResetDeployments {
			greengrass_ResetDeployments(cfg, client)
			return
		}
		if _greengrassStartBulkDeployment {
			greengrass_StartBulkDeployment(cfg, client)
			return
		}
		if _greengrassStopBulkDeployment {
			greengrass_StopBulkDeployment(cfg, client)
			return
		}
		if _greengrassTagResource {
			greengrass_TagResource(cfg, client)
			return
		}
		if _greengrassUntagResource {
			greengrass_UntagResource(cfg, client)
			return
		}
		if _greengrassUpdateConnectivityInfo {
			greengrass_UpdateConnectivityInfo(cfg, client)
			return
		}
		if _greengrassUpdateConnectorDefinition {
			greengrass_UpdateConnectorDefinition(cfg, client)
			return
		}
		if _greengrassUpdateCoreDefinition {
			greengrass_UpdateCoreDefinition(cfg, client)
			return
		}
		if _greengrassUpdateDeviceDefinition {
			greengrass_UpdateDeviceDefinition(cfg, client)
			return
		}
		if _greengrassUpdateFunctionDefinition {
			greengrass_UpdateFunctionDefinition(cfg, client)
			return
		}
		if _greengrassUpdateGroup {
			greengrass_UpdateGroup(cfg, client)
			return
		}
		if _greengrassUpdateGroupCertificateConfiguration {
			greengrass_UpdateGroupCertificateConfiguration(cfg, client)
			return
		}
		if _greengrassUpdateLoggerDefinition {
			greengrass_UpdateLoggerDefinition(cfg, client)
			return
		}
		if _greengrassUpdateResourceDefinition {
			greengrass_UpdateResourceDefinition(cfg, client)
			return
		}
		if _greengrassUpdateSubscriptionDefinition {
			greengrass_UpdateSubscriptionDefinition(cfg, client)
			return
		}
		if _greengrassUpdateThingRuntimeConfiguration {
			greengrass_UpdateThingRuntimeConfiguration(cfg, client)
			return
		}

	},
}

var (
	_greengrassAssociateRoleToGroup                bool
	_greengrassAssociateServiceRoleToAccount       bool
	_greengrassCreateConnectorDefinition           bool
	_greengrassCreateConnectorDefinitionVersion    bool
	_greengrassCreateCoreDefinition                bool
	_greengrassCreateCoreDefinitionVersion         bool
	_greengrassCreateDeployment                    bool
	_greengrassCreateDeviceDefinition              bool
	_greengrassCreateDeviceDefinitionVersion       bool
	_greengrassCreateFunctionDefinition            bool
	_greengrassCreateFunctionDefinitionVersion     bool
	_greengrassCreateGroup                         bool
	_greengrassCreateGroupCertificateAuthority     bool
	_greengrassCreateGroupVersion                  bool
	_greengrassCreateLoggerDefinition              bool
	_greengrassCreateLoggerDefinitionVersion       bool
	_greengrassCreateResourceDefinition            bool
	_greengrassCreateResourceDefinitionVersion     bool
	_greengrassCreateSoftwareUpdateJob             bool
	_greengrassCreateSubscriptionDefinition        bool
	_greengrassCreateSubscriptionDefinitionVersion bool
	_greengrassDeleteConnectorDefinition           bool
	_greengrassDeleteCoreDefinition                bool
	_greengrassDeleteDeviceDefinition              bool
	_greengrassDeleteFunctionDefinition            bool
	_greengrassDeleteGroup                         bool
	_greengrassDeleteLoggerDefinition              bool
	_greengrassDeleteResourceDefinition            bool
	_greengrassDeleteSubscriptionDefinition        bool
	_greengrassDisassociateRoleFromGroup           bool
	_greengrassDisassociateServiceRoleFromAccount  bool
	_greengrassGetAssociatedRole                   bool
	_greengrassGetBulkDeploymentStatus             bool
	_greengrassGetConnectivityInfo                 bool
	_greengrassGetConnectorDefinition              bool
	_greengrassGetConnectorDefinitionVersion       bool
	_greengrassGetCoreDefinition                   bool
	_greengrassGetCoreDefinitionVersion            bool
	_greengrassGetDeploymentStatus                 bool
	_greengrassGetDeviceDefinition                 bool
	_greengrassGetDeviceDefinitionVersion          bool
	_greengrassGetFunctionDefinition               bool
	_greengrassGetFunctionDefinitionVersion        bool
	_greengrassGetGroup                            bool
	_greengrassGetGroupCertificateAuthority        bool
	_greengrassGetGroupCertificateConfiguration    bool
	_greengrassGetGroupVersion                     bool
	_greengrassGetLoggerDefinition                 bool
	_greengrassGetLoggerDefinitionVersion          bool
	_greengrassGetResourceDefinition               bool
	_greengrassGetResourceDefinitionVersion        bool
	_greengrassGetServiceRoleForAccount            bool
	_greengrassGetSubscriptionDefinition           bool
	_greengrassGetSubscriptionDefinitionVersion    bool
	_greengrassGetThingRuntimeConfiguration        bool
	_greengrassListBulkDeploymentDetailedReports   bool
	_greengrassListBulkDeployments                 bool
	_greengrassListConnectorDefinitionVersions     bool
	_greengrassListConnectorDefinitions            bool
	_greengrassListCoreDefinitionVersions          bool
	_greengrassListCoreDefinitions                 bool
	_greengrassListDeployments                     bool
	_greengrassListDeviceDefinitionVersions        bool
	_greengrassListDeviceDefinitions               bool
	_greengrassListFunctionDefinitionVersions      bool
	_greengrassListFunctionDefinitions             bool
	_greengrassListGroupCertificateAuthorities     bool
	_greengrassListGroupVersions                   bool
	_greengrassListGroups                          bool
	_greengrassListLoggerDefinitionVersions        bool
	_greengrassListLoggerDefinitions               bool
	_greengrassListResourceDefinitionVersions      bool
	_greengrassListResourceDefinitions             bool
	_greengrassListSubscriptionDefinitionVersions  bool
	_greengrassListSubscriptionDefinitions         bool
	_greengrassListTagsForResource                 bool
	_greengrassResetDeployments                    bool
	_greengrassStartBulkDeployment                 bool
	_greengrassStopBulkDeployment                  bool
	_greengrassTagResource                         bool
	_greengrassUntagResource                       bool
	_greengrassUpdateConnectivityInfo              bool
	_greengrassUpdateConnectorDefinition           bool
	_greengrassUpdateCoreDefinition                bool
	_greengrassUpdateDeviceDefinition              bool
	_greengrassUpdateFunctionDefinition            bool
	_greengrassUpdateGroup                         bool
	_greengrassUpdateGroupCertificateConfiguration bool
	_greengrassUpdateLoggerDefinition              bool
	_greengrassUpdateResourceDefinition            bool
	_greengrassUpdateSubscriptionDefinition        bool
	_greengrassUpdateThingRuntimeConfiguration     bool

	_greengrassAmznClientToken                  string
	_greengrassBulkDeploymentId                 string
	_greengrassCertificateAuthorityId           string
	_greengrassCertificateExpiryInMilliseconds  string
	_greengrassConnectivityInfo                 string
	_greengrassConnectorDefinitionId            string
	_greengrassConnectorDefinitionVersionArn    string
	_greengrassConnectorDefinitionVersionId     string
	_greengrassConnectors                       string
	_greengrassCoreDefinitionId                 string
	_greengrassCoreDefinitionVersionArn         string
	_greengrassCoreDefinitionVersionId          string
	_greengrassCores                            string
	_greengrassDefaultConfig                    string
	_greengrassDeploymentId                     string
	_greengrassDeploymentType                   string
	_greengrassDeviceDefinitionId               string
	_greengrassDeviceDefinitionVersionArn       string
	_greengrassDeviceDefinitionVersionId        string
	_greengrassDevices                          string
	_greengrassExecutionRoleArn                 string
	_greengrassForce                            string
	_greengrassFunctionDefinitionId             string
	_greengrassFunctionDefinitionVersionArn     string
	_greengrassFunctionDefinitionVersionId      string
	_greengrassFunctions                        string
	_greengrassGroupId                          string
	_greengrassGroupVersionId                   string
	_greengrassInitialVersion                   string
	_greengrassInputFileUri                     string
	_greengrassLoggerDefinitionId               string
	_greengrassLoggerDefinitionVersionArn       string
	_greengrassLoggerDefinitionVersionId        string
	_greengrassLoggers                          string
	_greengrassMaxResults                       string
	_greengrassName                             string
	_greengrassNextToken                        string
	_greengrassResourceArn                      string
	_greengrassResourceDefinitionId             string
	_greengrassResourceDefinitionVersionArn     string
	_greengrassResourceDefinitionVersionId      string
	_greengrassResources                        string
	_greengrassRoleArn                          string
	_greengrassS3UrlSignerRole                  string
	_greengrassSoftwareToUpdate                 string
	_greengrassSubscriptionDefinitionId         string
	_greengrassSubscriptionDefinitionVersionArn string
	_greengrassSubscriptionDefinitionVersionId  string
	_greengrassSubscriptions                    string
	_greengrassTagKeys                          []string
	_greengrassTags                             string
	_greengrassTelemetryConfiguration           string
	_greengrassThingName                        string
	_greengrassUpdateAgentLogLevel              string
	_greengrassUpdateTargets                    []string
	_greengrassUpdateTargetsArchitecture        string
	_greengrassUpdateTargetsOperatingSystem     string
)

// Associates a role with a group. Your Greengrass core will use the role to
// access AWS cloud services. The role's permissions should allow Greengrass core
// Lambda functions to perform actions against the cloud.
func greengrass_AssociateRoleToGroup(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.AssociateRoleToGroupInput{
		// GroupId: *string, // Required
		// RoleArn: *string, // Required
	}

	if len(_greengrassGroupId) > 0 {
		input.GroupId = aws.String(_greengrassGroupId)
	}
	if len(_greengrassRoleArn) > 0 {
		input.RoleArn = aws.String(_greengrassRoleArn)
	}

	if resp, err := client.AssociateRoleToGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a role with your account. AWS IoT Greengrass will use the role to
// access your Lambda functions and AWS IoT resources. This is necessary for
// deployments to succeed. The role must have at least minimum permissions in the
// policy ”AWSGreengrassResourceAccessRolePolicy”.
func greengrass_AssociateServiceRoleToAccount(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.AssociateServiceRoleToAccountInput{
		// RoleArn: *string, // Required
	}

	if len(_greengrassRoleArn) > 0 {
		input.RoleArn = aws.String(_greengrassRoleArn)
	}

	if resp, err := client.AssociateServiceRoleToAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a connector definition. You may provide the initial version of the
// connector definition now or use ”CreateConnectorDefinitionVersion” at a later
// time.
func greengrass_CreateConnectorDefinition(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.CreateConnectorDefinitionInput{}

	if len(_greengrassAmznClientToken) > 0 {
		input.AmznClientToken = aws.String(_greengrassAmznClientToken)
	}
	if len(_greengrassInitialVersion) > 0 {
		if err := assignInputField(input, "InitialVersion", _greengrassInitialVersion); err != nil {
			log.Errorf("invalid --initial-version: %s", err.Error())
			return
		}
	}
	if len(_greengrassName) > 0 {
		input.Name = aws.String(_greengrassName)
	}
	if len(_greengrassTags) > 0 {
		if err := assignInputField(input, "Tags", _greengrassTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateConnectorDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a version of a connector definition which has already been defined.
func greengrass_CreateConnectorDefinitionVersion(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.CreateConnectorDefinitionVersionInput{
		// ConnectorDefinitionId: *string, // Required
	}

	if len(_greengrassConnectorDefinitionId) > 0 {
		input.ConnectorDefinitionId = aws.String(_greengrassConnectorDefinitionId)
	}
	if len(_greengrassAmznClientToken) > 0 {
		input.AmznClientToken = aws.String(_greengrassAmznClientToken)
	}
	if len(_greengrassConnectors) > 0 {
		if err := assignInputField(input, "Connectors", _greengrassConnectors); err != nil {
			log.Errorf("invalid --connectors: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateConnectorDefinitionVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a core definition. You may provide the initial version of the core
// definition now or use ”CreateCoreDefinitionVersion” at a later time.
// Greengrass groups must each contain exactly one Greengrass core.
func greengrass_CreateCoreDefinition(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.CreateCoreDefinitionInput{}

	if len(_greengrassAmznClientToken) > 0 {
		input.AmznClientToken = aws.String(_greengrassAmznClientToken)
	}
	if len(_greengrassInitialVersion) > 0 {
		if err := assignInputField(input, "InitialVersion", _greengrassInitialVersion); err != nil {
			log.Errorf("invalid --initial-version: %s", err.Error())
			return
		}
	}
	if len(_greengrassName) > 0 {
		input.Name = aws.String(_greengrassName)
	}
	if len(_greengrassTags) > 0 {
		if err := assignInputField(input, "Tags", _greengrassTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCoreDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a version of a core definition that has already been defined.
// Greengrass groups must each contain exactly one Greengrass core.
func greengrass_CreateCoreDefinitionVersion(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.CreateCoreDefinitionVersionInput{
		// CoreDefinitionId: *string, // Required
	}

	if len(_greengrassCoreDefinitionId) > 0 {
		input.CoreDefinitionId = aws.String(_greengrassCoreDefinitionId)
	}
	if len(_greengrassAmznClientToken) > 0 {
		input.AmznClientToken = aws.String(_greengrassAmznClientToken)
	}
	if len(_greengrassCores) > 0 {
		if err := assignInputField(input, "Cores", _greengrassCores); err != nil {
			log.Errorf("invalid --cores: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCoreDefinitionVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a deployment. ”CreateDeployment” requests are idempotent with respect
// to the ”X-Amzn-Client-Token” token and the request parameters.
func greengrass_CreateDeployment(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.CreateDeploymentInput{
		// DeploymentType: types.DeploymentType, // Required
		// GroupId: *string, // Required
	}

	if len(_greengrassDeploymentType) > 0 {
		if err := assignInputField(input, "DeploymentType", _greengrassDeploymentType); err != nil {
			log.Errorf("invalid --deployment-type: %s", err.Error())
			return
		}
	}
	if len(_greengrassGroupId) > 0 {
		input.GroupId = aws.String(_greengrassGroupId)
	}
	if len(_greengrassAmznClientToken) > 0 {
		input.AmznClientToken = aws.String(_greengrassAmznClientToken)
	}
	if len(_greengrassDeploymentId) > 0 {
		input.DeploymentId = aws.String(_greengrassDeploymentId)
	}
	if len(_greengrassGroupVersionId) > 0 {
		input.GroupVersionId = aws.String(_greengrassGroupVersionId)
	}

	if resp, err := client.CreateDeployment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a device definition. You may provide the initial version of the device
// definition now or use ”CreateDeviceDefinitionVersion” at a later time.
func greengrass_CreateDeviceDefinition(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.CreateDeviceDefinitionInput{}

	if len(_greengrassAmznClientToken) > 0 {
		input.AmznClientToken = aws.String(_greengrassAmznClientToken)
	}
	if len(_greengrassInitialVersion) > 0 {
		if err := assignInputField(input, "InitialVersion", _greengrassInitialVersion); err != nil {
			log.Errorf("invalid --initial-version: %s", err.Error())
			return
		}
	}
	if len(_greengrassName) > 0 {
		input.Name = aws.String(_greengrassName)
	}
	if len(_greengrassTags) > 0 {
		if err := assignInputField(input, "Tags", _greengrassTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDeviceDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a version of a device definition that has already been defined.
func greengrass_CreateDeviceDefinitionVersion(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.CreateDeviceDefinitionVersionInput{
		// DeviceDefinitionId: *string, // Required
	}

	if len(_greengrassDeviceDefinitionId) > 0 {
		input.DeviceDefinitionId = aws.String(_greengrassDeviceDefinitionId)
	}
	if len(_greengrassAmznClientToken) > 0 {
		input.AmznClientToken = aws.String(_greengrassAmznClientToken)
	}
	if len(_greengrassDevices) > 0 {
		if err := assignInputField(input, "Devices", _greengrassDevices); err != nil {
			log.Errorf("invalid --devices: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDeviceDefinitionVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a Lambda function definition which contains a list of Lambda functions
// and their configurations to be used in a group. You can create an initial
// version of the definition by providing a list of Lambda functions and their
// configurations now, or use ”CreateFunctionDefinitionVersion” later.
func greengrass_CreateFunctionDefinition(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.CreateFunctionDefinitionInput{}

	if len(_greengrassAmznClientToken) > 0 {
		input.AmznClientToken = aws.String(_greengrassAmznClientToken)
	}
	if len(_greengrassInitialVersion) > 0 {
		if err := assignInputField(input, "InitialVersion", _greengrassInitialVersion); err != nil {
			log.Errorf("invalid --initial-version: %s", err.Error())
			return
		}
	}
	if len(_greengrassName) > 0 {
		input.Name = aws.String(_greengrassName)
	}
	if len(_greengrassTags) > 0 {
		if err := assignInputField(input, "Tags", _greengrassTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateFunctionDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a version of a Lambda function definition that has already been defined.
func greengrass_CreateFunctionDefinitionVersion(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.CreateFunctionDefinitionVersionInput{
		// FunctionDefinitionId: *string, // Required
	}

	if len(_greengrassFunctionDefinitionId) > 0 {
		input.FunctionDefinitionId = aws.String(_greengrassFunctionDefinitionId)
	}
	if len(_greengrassAmznClientToken) > 0 {
		input.AmznClientToken = aws.String(_greengrassAmznClientToken)
	}
	if len(_greengrassDefaultConfig) > 0 {
		if err := assignInputField(input, "DefaultConfig", _greengrassDefaultConfig); err != nil {
			log.Errorf("invalid --default-config: %s", err.Error())
			return
		}
	}
	if len(_greengrassFunctions) > 0 {
		if err := assignInputField(input, "Functions", _greengrassFunctions); err != nil {
			log.Errorf("invalid --functions: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateFunctionDefinitionVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a group. You may provide the initial version of the group or use
// ”CreateGroupVersion” at a later time. Tip: You can use the ”gg_group_setup”
// package (https://github.com/awslabs/aws-greengrass-group-setup) as a library or
// command-line application to create and deploy Greengrass groups.
func greengrass_CreateGroup(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.CreateGroupInput{
		// Name: *string, // Required
	}

	if len(_greengrassName) > 0 {
		input.Name = aws.String(_greengrassName)
	}
	if len(_greengrassAmznClientToken) > 0 {
		input.AmznClientToken = aws.String(_greengrassAmznClientToken)
	}
	if len(_greengrassInitialVersion) > 0 {
		if err := assignInputField(input, "InitialVersion", _greengrassInitialVersion); err != nil {
			log.Errorf("invalid --initial-version: %s", err.Error())
			return
		}
	}
	if len(_greengrassTags) > 0 {
		if err := assignInputField(input, "Tags", _greengrassTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a CA for the group. If a CA already exists, it will rotate the existing
// CA.
func greengrass_CreateGroupCertificateAuthority(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.CreateGroupCertificateAuthorityInput{
		// GroupId: *string, // Required
	}

	if len(_greengrassGroupId) > 0 {
		input.GroupId = aws.String(_greengrassGroupId)
	}
	if len(_greengrassAmznClientToken) > 0 {
		input.AmznClientToken = aws.String(_greengrassAmznClientToken)
	}

	if resp, err := client.CreateGroupCertificateAuthority(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a version of a group which has already been defined.
func greengrass_CreateGroupVersion(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.CreateGroupVersionInput{
		// GroupId: *string, // Required
	}

	if len(_greengrassGroupId) > 0 {
		input.GroupId = aws.String(_greengrassGroupId)
	}
	if len(_greengrassAmznClientToken) > 0 {
		input.AmznClientToken = aws.String(_greengrassAmznClientToken)
	}
	if len(_greengrassConnectorDefinitionVersionArn) > 0 {
		input.ConnectorDefinitionVersionArn = aws.String(_greengrassConnectorDefinitionVersionArn)
	}
	if len(_greengrassCoreDefinitionVersionArn) > 0 {
		input.CoreDefinitionVersionArn = aws.String(_greengrassCoreDefinitionVersionArn)
	}
	if len(_greengrassDeviceDefinitionVersionArn) > 0 {
		input.DeviceDefinitionVersionArn = aws.String(_greengrassDeviceDefinitionVersionArn)
	}
	if len(_greengrassFunctionDefinitionVersionArn) > 0 {
		input.FunctionDefinitionVersionArn = aws.String(_greengrassFunctionDefinitionVersionArn)
	}
	if len(_greengrassLoggerDefinitionVersionArn) > 0 {
		input.LoggerDefinitionVersionArn = aws.String(_greengrassLoggerDefinitionVersionArn)
	}
	if len(_greengrassResourceDefinitionVersionArn) > 0 {
		input.ResourceDefinitionVersionArn = aws.String(_greengrassResourceDefinitionVersionArn)
	}
	if len(_greengrassSubscriptionDefinitionVersionArn) > 0 {
		input.SubscriptionDefinitionVersionArn = aws.String(_greengrassSubscriptionDefinitionVersionArn)
	}

	if resp, err := client.CreateGroupVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a logger definition. You may provide the initial version of the logger
// definition now or use ”CreateLoggerDefinitionVersion” at a later time.
func greengrass_CreateLoggerDefinition(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.CreateLoggerDefinitionInput{}

	if len(_greengrassAmznClientToken) > 0 {
		input.AmznClientToken = aws.String(_greengrassAmznClientToken)
	}
	if len(_greengrassInitialVersion) > 0 {
		if err := assignInputField(input, "InitialVersion", _greengrassInitialVersion); err != nil {
			log.Errorf("invalid --initial-version: %s", err.Error())
			return
		}
	}
	if len(_greengrassName) > 0 {
		input.Name = aws.String(_greengrassName)
	}
	if len(_greengrassTags) > 0 {
		if err := assignInputField(input, "Tags", _greengrassTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateLoggerDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a version of a logger definition that has already been defined.
func greengrass_CreateLoggerDefinitionVersion(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.CreateLoggerDefinitionVersionInput{
		// LoggerDefinitionId: *string, // Required
	}

	if len(_greengrassLoggerDefinitionId) > 0 {
		input.LoggerDefinitionId = aws.String(_greengrassLoggerDefinitionId)
	}
	if len(_greengrassAmznClientToken) > 0 {
		input.AmznClientToken = aws.String(_greengrassAmznClientToken)
	}
	if len(_greengrassLoggers) > 0 {
		if err := assignInputField(input, "Loggers", _greengrassLoggers); err != nil {
			log.Errorf("invalid --loggers: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateLoggerDefinitionVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a resource definition which contains a list of resources to be used in
// a group. You can create an initial version of the definition by providing a list
// of resources now, or use ”CreateResourceDefinitionVersion” later.
func greengrass_CreateResourceDefinition(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.CreateResourceDefinitionInput{}

	if len(_greengrassAmznClientToken) > 0 {
		input.AmznClientToken = aws.String(_greengrassAmznClientToken)
	}
	if len(_greengrassInitialVersion) > 0 {
		if err := assignInputField(input, "InitialVersion", _greengrassInitialVersion); err != nil {
			log.Errorf("invalid --initial-version: %s", err.Error())
			return
		}
	}
	if len(_greengrassName) > 0 {
		input.Name = aws.String(_greengrassName)
	}
	if len(_greengrassTags) > 0 {
		if err := assignInputField(input, "Tags", _greengrassTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateResourceDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a version of a resource definition that has already been defined.
func greengrass_CreateResourceDefinitionVersion(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.CreateResourceDefinitionVersionInput{
		// ResourceDefinitionId: *string, // Required
	}

	if len(_greengrassResourceDefinitionId) > 0 {
		input.ResourceDefinitionId = aws.String(_greengrassResourceDefinitionId)
	}
	if len(_greengrassAmznClientToken) > 0 {
		input.AmznClientToken = aws.String(_greengrassAmznClientToken)
	}
	if len(_greengrassResources) > 0 {
		if err := assignInputField(input, "Resources", _greengrassResources); err != nil {
			log.Errorf("invalid --resources: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateResourceDefinitionVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a software update for a core or group of cores (specified as an IoT
// thing group.) Use this to update the OTA Agent as well as the Greengrass core
// software. It makes use of the IoT Jobs feature which provides additional
// commands to manage a Greengrass core software update job.
func greengrass_CreateSoftwareUpdateJob(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.CreateSoftwareUpdateJobInput{
		// S3UrlSignerRole: *string, // Required
		// SoftwareToUpdate: types.SoftwareToUpdate, // Required
		// UpdateTargets: []string, // Required
		// UpdateTargetsArchitecture: types.UpdateTargetsArchitecture, // Required
		// UpdateTargetsOperatingSystem: types.UpdateTargetsOperatingSystem, // Required
	}

	if len(_greengrassS3UrlSignerRole) > 0 {
		input.S3UrlSignerRole = aws.String(_greengrassS3UrlSignerRole)
	}
	if len(_greengrassSoftwareToUpdate) > 0 {
		if err := assignInputField(input, "SoftwareToUpdate", _greengrassSoftwareToUpdate); err != nil {
			log.Errorf("invalid --software-to-update: %s", err.Error())
			return
		}
	}
	if len(_greengrassUpdateTargets) > 0 {
		input.UpdateTargets = append([]string(nil), _greengrassUpdateTargets...)
	}
	if len(_greengrassUpdateTargetsArchitecture) > 0 {
		if err := assignInputField(input, "UpdateTargetsArchitecture", _greengrassUpdateTargetsArchitecture); err != nil {
			log.Errorf("invalid --update-targets-architecture: %s", err.Error())
			return
		}
	}
	if len(_greengrassUpdateTargetsOperatingSystem) > 0 {
		if err := assignInputField(input, "UpdateTargetsOperatingSystem", _greengrassUpdateTargetsOperatingSystem); err != nil {
			log.Errorf("invalid --update-targets-operating-system: %s", err.Error())
			return
		}
	}
	if len(_greengrassAmznClientToken) > 0 {
		input.AmznClientToken = aws.String(_greengrassAmznClientToken)
	}
	if len(_greengrassUpdateAgentLogLevel) > 0 {
		if err := assignInputField(input, "UpdateAgentLogLevel", _greengrassUpdateAgentLogLevel); err != nil {
			log.Errorf("invalid --update-agent-log-level: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSoftwareUpdateJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a subscription definition. You may provide the initial version of the
// subscription definition now or use ”CreateSubscriptionDefinitionVersion” at a
// later time.
func greengrass_CreateSubscriptionDefinition(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.CreateSubscriptionDefinitionInput{}

	if len(_greengrassAmznClientToken) > 0 {
		input.AmznClientToken = aws.String(_greengrassAmznClientToken)
	}
	if len(_greengrassInitialVersion) > 0 {
		if err := assignInputField(input, "InitialVersion", _greengrassInitialVersion); err != nil {
			log.Errorf("invalid --initial-version: %s", err.Error())
			return
		}
	}
	if len(_greengrassName) > 0 {
		input.Name = aws.String(_greengrassName)
	}
	if len(_greengrassTags) > 0 {
		if err := assignInputField(input, "Tags", _greengrassTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSubscriptionDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a version of a subscription definition which has already been defined.
func greengrass_CreateSubscriptionDefinitionVersion(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.CreateSubscriptionDefinitionVersionInput{
		// SubscriptionDefinitionId: *string, // Required
	}

	if len(_greengrassSubscriptionDefinitionId) > 0 {
		input.SubscriptionDefinitionId = aws.String(_greengrassSubscriptionDefinitionId)
	}
	if len(_greengrassAmznClientToken) > 0 {
		input.AmznClientToken = aws.String(_greengrassAmznClientToken)
	}
	if len(_greengrassSubscriptions) > 0 {
		if err := assignInputField(input, "Subscriptions", _greengrassSubscriptions); err != nil {
			log.Errorf("invalid --subscriptions: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSubscriptionDefinitionVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a connector definition.
func greengrass_DeleteConnectorDefinition(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.DeleteConnectorDefinitionInput{
		// ConnectorDefinitionId: *string, // Required
	}

	if len(_greengrassConnectorDefinitionId) > 0 {
		input.ConnectorDefinitionId = aws.String(_greengrassConnectorDefinitionId)
	}

	if resp, err := client.DeleteConnectorDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a core definition.
func greengrass_DeleteCoreDefinition(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.DeleteCoreDefinitionInput{
		// CoreDefinitionId: *string, // Required
	}

	if len(_greengrassCoreDefinitionId) > 0 {
		input.CoreDefinitionId = aws.String(_greengrassCoreDefinitionId)
	}

	if resp, err := client.DeleteCoreDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a device definition.
func greengrass_DeleteDeviceDefinition(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.DeleteDeviceDefinitionInput{
		// DeviceDefinitionId: *string, // Required
	}

	if len(_greengrassDeviceDefinitionId) > 0 {
		input.DeviceDefinitionId = aws.String(_greengrassDeviceDefinitionId)
	}

	if resp, err := client.DeleteDeviceDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Lambda function definition.
func greengrass_DeleteFunctionDefinition(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.DeleteFunctionDefinitionInput{
		// FunctionDefinitionId: *string, // Required
	}

	if len(_greengrassFunctionDefinitionId) > 0 {
		input.FunctionDefinitionId = aws.String(_greengrassFunctionDefinitionId)
	}

	if resp, err := client.DeleteFunctionDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a group.
func greengrass_DeleteGroup(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.DeleteGroupInput{
		// GroupId: *string, // Required
	}

	if len(_greengrassGroupId) > 0 {
		input.GroupId = aws.String(_greengrassGroupId)
	}

	if resp, err := client.DeleteGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a logger definition.
func greengrass_DeleteLoggerDefinition(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.DeleteLoggerDefinitionInput{
		// LoggerDefinitionId: *string, // Required
	}

	if len(_greengrassLoggerDefinitionId) > 0 {
		input.LoggerDefinitionId = aws.String(_greengrassLoggerDefinitionId)
	}

	if resp, err := client.DeleteLoggerDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a resource definition.
func greengrass_DeleteResourceDefinition(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.DeleteResourceDefinitionInput{
		// ResourceDefinitionId: *string, // Required
	}

	if len(_greengrassResourceDefinitionId) > 0 {
		input.ResourceDefinitionId = aws.String(_greengrassResourceDefinitionId)
	}

	if resp, err := client.DeleteResourceDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a subscription definition.
func greengrass_DeleteSubscriptionDefinition(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.DeleteSubscriptionDefinitionInput{
		// SubscriptionDefinitionId: *string, // Required
	}

	if len(_greengrassSubscriptionDefinitionId) > 0 {
		input.SubscriptionDefinitionId = aws.String(_greengrassSubscriptionDefinitionId)
	}

	if resp, err := client.DeleteSubscriptionDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates the role from a group.
func greengrass_DisassociateRoleFromGroup(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.DisassociateRoleFromGroupInput{
		// GroupId: *string, // Required
	}

	if len(_greengrassGroupId) > 0 {
		input.GroupId = aws.String(_greengrassGroupId)
	}

	if resp, err := client.DisassociateRoleFromGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates the service role from your account. Without a service role,
// deployments will not work.
func greengrass_DisassociateServiceRoleFromAccount(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.DisassociateServiceRoleFromAccountInput{}

	if resp, err := client.DisassociateServiceRoleFromAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the role associated with a particular group.
func greengrass_GetAssociatedRole(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.GetAssociatedRoleInput{
		// GroupId: *string, // Required
	}

	if len(_greengrassGroupId) > 0 {
		input.GroupId = aws.String(_greengrassGroupId)
	}

	if resp, err := client.GetAssociatedRole(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the status of a bulk deployment.
func greengrass_GetBulkDeploymentStatus(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.GetBulkDeploymentStatusInput{
		// BulkDeploymentId: *string, // Required
	}

	if len(_greengrassBulkDeploymentId) > 0 {
		input.BulkDeploymentId = aws.String(_greengrassBulkDeploymentId)
	}

	if resp, err := client.GetBulkDeploymentStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the connectivity information for a core.
func greengrass_GetConnectivityInfo(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.GetConnectivityInfoInput{
		// ThingName: *string, // Required
	}

	if len(_greengrassThingName) > 0 {
		input.ThingName = aws.String(_greengrassThingName)
	}

	if resp, err := client.GetConnectivityInfo(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a connector definition.
func greengrass_GetConnectorDefinition(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.GetConnectorDefinitionInput{
		// ConnectorDefinitionId: *string, // Required
	}

	if len(_greengrassConnectorDefinitionId) > 0 {
		input.ConnectorDefinitionId = aws.String(_greengrassConnectorDefinitionId)
	}

	if resp, err := client.GetConnectorDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a connector definition version, including the
// connectors that the version contains. Connectors are prebuilt modules that
// interact with local infrastructure, device protocols, AWS, and other cloud
// services.
func greengrass_GetConnectorDefinitionVersion(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.GetConnectorDefinitionVersionInput{
		// ConnectorDefinitionId: *string, // Required
		// ConnectorDefinitionVersionId: *string, // Required
	}

	if len(_greengrassConnectorDefinitionId) > 0 {
		input.ConnectorDefinitionId = aws.String(_greengrassConnectorDefinitionId)
	}
	if len(_greengrassConnectorDefinitionVersionId) > 0 {
		input.ConnectorDefinitionVersionId = aws.String(_greengrassConnectorDefinitionVersionId)
	}
	if len(_greengrassNextToken) > 0 {
		input.NextToken = aws.String(_greengrassNextToken)
	}

	if resp, err := client.GetConnectorDefinitionVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a core definition version.
func greengrass_GetCoreDefinition(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.GetCoreDefinitionInput{
		// CoreDefinitionId: *string, // Required
	}

	if len(_greengrassCoreDefinitionId) > 0 {
		input.CoreDefinitionId = aws.String(_greengrassCoreDefinitionId)
	}

	if resp, err := client.GetCoreDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a core definition version.
func greengrass_GetCoreDefinitionVersion(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.GetCoreDefinitionVersionInput{
		// CoreDefinitionId: *string, // Required
		// CoreDefinitionVersionId: *string, // Required
	}

	if len(_greengrassCoreDefinitionId) > 0 {
		input.CoreDefinitionId = aws.String(_greengrassCoreDefinitionId)
	}
	if len(_greengrassCoreDefinitionVersionId) > 0 {
		input.CoreDefinitionVersionId = aws.String(_greengrassCoreDefinitionVersionId)
	}

	if resp, err := client.GetCoreDefinitionVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the status of a deployment.
func greengrass_GetDeploymentStatus(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.GetDeploymentStatusInput{
		// DeploymentId: *string, // Required
		// GroupId: *string, // Required
	}

	if len(_greengrassDeploymentId) > 0 {
		input.DeploymentId = aws.String(_greengrassDeploymentId)
	}
	if len(_greengrassGroupId) > 0 {
		input.GroupId = aws.String(_greengrassGroupId)
	}

	if resp, err := client.GetDeploymentStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a device definition.
func greengrass_GetDeviceDefinition(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.GetDeviceDefinitionInput{
		// DeviceDefinitionId: *string, // Required
	}

	if len(_greengrassDeviceDefinitionId) > 0 {
		input.DeviceDefinitionId = aws.String(_greengrassDeviceDefinitionId)
	}

	if resp, err := client.GetDeviceDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a device definition version.
func greengrass_GetDeviceDefinitionVersion(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.GetDeviceDefinitionVersionInput{
		// DeviceDefinitionId: *string, // Required
		// DeviceDefinitionVersionId: *string, // Required
	}

	if len(_greengrassDeviceDefinitionId) > 0 {
		input.DeviceDefinitionId = aws.String(_greengrassDeviceDefinitionId)
	}
	if len(_greengrassDeviceDefinitionVersionId) > 0 {
		input.DeviceDefinitionVersionId = aws.String(_greengrassDeviceDefinitionVersionId)
	}
	if len(_greengrassNextToken) > 0 {
		input.NextToken = aws.String(_greengrassNextToken)
	}

	if resp, err := client.GetDeviceDefinitionVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a Lambda function definition, including its
// creation time and latest version.
func greengrass_GetFunctionDefinition(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.GetFunctionDefinitionInput{
		// FunctionDefinitionId: *string, // Required
	}

	if len(_greengrassFunctionDefinitionId) > 0 {
		input.FunctionDefinitionId = aws.String(_greengrassFunctionDefinitionId)
	}

	if resp, err := client.GetFunctionDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a Lambda function definition version, including
// which Lambda functions are included in the version and their configurations.
func greengrass_GetFunctionDefinitionVersion(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.GetFunctionDefinitionVersionInput{
		// FunctionDefinitionId: *string, // Required
		// FunctionDefinitionVersionId: *string, // Required
	}

	if len(_greengrassFunctionDefinitionId) > 0 {
		input.FunctionDefinitionId = aws.String(_greengrassFunctionDefinitionId)
	}
	if len(_greengrassFunctionDefinitionVersionId) > 0 {
		input.FunctionDefinitionVersionId = aws.String(_greengrassFunctionDefinitionVersionId)
	}
	if len(_greengrassNextToken) > 0 {
		input.NextToken = aws.String(_greengrassNextToken)
	}

	if resp, err := client.GetFunctionDefinitionVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a group.
func greengrass_GetGroup(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.GetGroupInput{
		// GroupId: *string, // Required
	}

	if len(_greengrassGroupId) > 0 {
		input.GroupId = aws.String(_greengrassGroupId)
	}

	if resp, err := client.GetGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retreives the CA associated with a group. Returns the public key of the CA.
func greengrass_GetGroupCertificateAuthority(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.GetGroupCertificateAuthorityInput{
		// CertificateAuthorityId: *string, // Required
		// GroupId: *string, // Required
	}

	if len(_greengrassCertificateAuthorityId) > 0 {
		input.CertificateAuthorityId = aws.String(_greengrassCertificateAuthorityId)
	}
	if len(_greengrassGroupId) > 0 {
		input.GroupId = aws.String(_greengrassGroupId)
	}

	if resp, err := client.GetGroupCertificateAuthority(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the current configuration for the CA used by the group.
func greengrass_GetGroupCertificateConfiguration(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.GetGroupCertificateConfigurationInput{
		// GroupId: *string, // Required
	}

	if len(_greengrassGroupId) > 0 {
		input.GroupId = aws.String(_greengrassGroupId)
	}

	if resp, err := client.GetGroupCertificateConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a group version.
func greengrass_GetGroupVersion(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.GetGroupVersionInput{
		// GroupId: *string, // Required
		// GroupVersionId: *string, // Required
	}

	if len(_greengrassGroupId) > 0 {
		input.GroupId = aws.String(_greengrassGroupId)
	}
	if len(_greengrassGroupVersionId) > 0 {
		input.GroupVersionId = aws.String(_greengrassGroupVersionId)
	}

	if resp, err := client.GetGroupVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a logger definition.
func greengrass_GetLoggerDefinition(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.GetLoggerDefinitionInput{
		// LoggerDefinitionId: *string, // Required
	}

	if len(_greengrassLoggerDefinitionId) > 0 {
		input.LoggerDefinitionId = aws.String(_greengrassLoggerDefinitionId)
	}

	if resp, err := client.GetLoggerDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a logger definition version.
func greengrass_GetLoggerDefinitionVersion(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.GetLoggerDefinitionVersionInput{
		// LoggerDefinitionId: *string, // Required
		// LoggerDefinitionVersionId: *string, // Required
	}

	if len(_greengrassLoggerDefinitionId) > 0 {
		input.LoggerDefinitionId = aws.String(_greengrassLoggerDefinitionId)
	}
	if len(_greengrassLoggerDefinitionVersionId) > 0 {
		input.LoggerDefinitionVersionId = aws.String(_greengrassLoggerDefinitionVersionId)
	}
	if len(_greengrassNextToken) > 0 {
		input.NextToken = aws.String(_greengrassNextToken)
	}

	if resp, err := client.GetLoggerDefinitionVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a resource definition, including its creation time
// and latest version.
func greengrass_GetResourceDefinition(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.GetResourceDefinitionInput{
		// ResourceDefinitionId: *string, // Required
	}

	if len(_greengrassResourceDefinitionId) > 0 {
		input.ResourceDefinitionId = aws.String(_greengrassResourceDefinitionId)
	}

	if resp, err := client.GetResourceDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a resource definition version, including which
// resources are included in the version.
func greengrass_GetResourceDefinitionVersion(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.GetResourceDefinitionVersionInput{
		// ResourceDefinitionId: *string, // Required
		// ResourceDefinitionVersionId: *string, // Required
	}

	if len(_greengrassResourceDefinitionId) > 0 {
		input.ResourceDefinitionId = aws.String(_greengrassResourceDefinitionId)
	}
	if len(_greengrassResourceDefinitionVersionId) > 0 {
		input.ResourceDefinitionVersionId = aws.String(_greengrassResourceDefinitionVersionId)
	}

	if resp, err := client.GetResourceDefinitionVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the service role that is attached to your account.
func greengrass_GetServiceRoleForAccount(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.GetServiceRoleForAccountInput{}

	if resp, err := client.GetServiceRoleForAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a subscription definition.
func greengrass_GetSubscriptionDefinition(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.GetSubscriptionDefinitionInput{
		// SubscriptionDefinitionId: *string, // Required
	}

	if len(_greengrassSubscriptionDefinitionId) > 0 {
		input.SubscriptionDefinitionId = aws.String(_greengrassSubscriptionDefinitionId)
	}

	if resp, err := client.GetSubscriptionDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a subscription definition version.
func greengrass_GetSubscriptionDefinitionVersion(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.GetSubscriptionDefinitionVersionInput{
		// SubscriptionDefinitionId: *string, // Required
		// SubscriptionDefinitionVersionId: *string, // Required
	}

	if len(_greengrassSubscriptionDefinitionId) > 0 {
		input.SubscriptionDefinitionId = aws.String(_greengrassSubscriptionDefinitionId)
	}
	if len(_greengrassSubscriptionDefinitionVersionId) > 0 {
		input.SubscriptionDefinitionVersionId = aws.String(_greengrassSubscriptionDefinitionVersionId)
	}
	if len(_greengrassNextToken) > 0 {
		input.NextToken = aws.String(_greengrassNextToken)
	}

	if resp, err := client.GetSubscriptionDefinitionVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the runtime configuration of a thing.
func greengrass_GetThingRuntimeConfiguration(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.GetThingRuntimeConfigurationInput{
		// ThingName: *string, // Required
	}

	if len(_greengrassThingName) > 0 {
		input.ThingName = aws.String(_greengrassThingName)
	}

	if resp, err := client.GetThingRuntimeConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a paginated list of the deployments that have been started in a bulk
// deployment operation, and their current deployment status.
func greengrass_ListBulkDeploymentDetailedReports(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.ListBulkDeploymentDetailedReportsInput{
		// BulkDeploymentId: *string, // Required
	}

	if len(_greengrassBulkDeploymentId) > 0 {
		input.BulkDeploymentId = aws.String(_greengrassBulkDeploymentId)
	}
	if len(_greengrassMaxResults) > 0 {
		input.MaxResults = aws.String(_greengrassMaxResults)
	}
	if len(_greengrassNextToken) > 0 {
		input.NextToken = aws.String(_greengrassNextToken)
	}

	if resp, err := client.ListBulkDeploymentDetailedReports(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of bulk deployments.
func greengrass_ListBulkDeployments(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.ListBulkDeploymentsInput{}

	if len(_greengrassMaxResults) > 0 {
		input.MaxResults = aws.String(_greengrassMaxResults)
	}
	if len(_greengrassNextToken) > 0 {
		input.NextToken = aws.String(_greengrassNextToken)
	}

	if resp, err := client.ListBulkDeployments(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the versions of a connector definition, which are containers for
// connectors. Connectors run on the Greengrass core and contain built-in
// integration with local infrastructure, device protocols, AWS, and other cloud
// services.
func greengrass_ListConnectorDefinitionVersions(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.ListConnectorDefinitionVersionsInput{
		// ConnectorDefinitionId: *string, // Required
	}

	if len(_greengrassConnectorDefinitionId) > 0 {
		input.ConnectorDefinitionId = aws.String(_greengrassConnectorDefinitionId)
	}
	if len(_greengrassMaxResults) > 0 {
		input.MaxResults = aws.String(_greengrassMaxResults)
	}
	if len(_greengrassNextToken) > 0 {
		input.NextToken = aws.String(_greengrassNextToken)
	}

	if resp, err := client.ListConnectorDefinitionVersions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of connector definitions.
func greengrass_ListConnectorDefinitions(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.ListConnectorDefinitionsInput{}

	if len(_greengrassMaxResults) > 0 {
		input.MaxResults = aws.String(_greengrassMaxResults)
	}
	if len(_greengrassNextToken) > 0 {
		input.NextToken = aws.String(_greengrassNextToken)
	}

	if resp, err := client.ListConnectorDefinitions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the versions of a core definition.
func greengrass_ListCoreDefinitionVersions(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.ListCoreDefinitionVersionsInput{
		// CoreDefinitionId: *string, // Required
	}

	if len(_greengrassCoreDefinitionId) > 0 {
		input.CoreDefinitionId = aws.String(_greengrassCoreDefinitionId)
	}
	if len(_greengrassMaxResults) > 0 {
		input.MaxResults = aws.String(_greengrassMaxResults)
	}
	if len(_greengrassNextToken) > 0 {
		input.NextToken = aws.String(_greengrassNextToken)
	}

	if resp, err := client.ListCoreDefinitionVersions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of core definitions.
func greengrass_ListCoreDefinitions(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.ListCoreDefinitionsInput{}

	if len(_greengrassMaxResults) > 0 {
		input.MaxResults = aws.String(_greengrassMaxResults)
	}
	if len(_greengrassNextToken) > 0 {
		input.NextToken = aws.String(_greengrassNextToken)
	}

	if resp, err := client.ListCoreDefinitions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a history of deployments for the group.
func greengrass_ListDeployments(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.ListDeploymentsInput{
		// GroupId: *string, // Required
	}

	if len(_greengrassGroupId) > 0 {
		input.GroupId = aws.String(_greengrassGroupId)
	}
	if len(_greengrassMaxResults) > 0 {
		input.MaxResults = aws.String(_greengrassMaxResults)
	}
	if len(_greengrassNextToken) > 0 {
		input.NextToken = aws.String(_greengrassNextToken)
	}

	if resp, err := client.ListDeployments(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the versions of a device definition.
func greengrass_ListDeviceDefinitionVersions(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.ListDeviceDefinitionVersionsInput{
		// DeviceDefinitionId: *string, // Required
	}

	if len(_greengrassDeviceDefinitionId) > 0 {
		input.DeviceDefinitionId = aws.String(_greengrassDeviceDefinitionId)
	}
	if len(_greengrassMaxResults) > 0 {
		input.MaxResults = aws.String(_greengrassMaxResults)
	}
	if len(_greengrassNextToken) > 0 {
		input.NextToken = aws.String(_greengrassNextToken)
	}

	if resp, err := client.ListDeviceDefinitionVersions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of device definitions.
func greengrass_ListDeviceDefinitions(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.ListDeviceDefinitionsInput{}

	if len(_greengrassMaxResults) > 0 {
		input.MaxResults = aws.String(_greengrassMaxResults)
	}
	if len(_greengrassNextToken) > 0 {
		input.NextToken = aws.String(_greengrassNextToken)
	}

	if resp, err := client.ListDeviceDefinitions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the versions of a Lambda function definition.
func greengrass_ListFunctionDefinitionVersions(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.ListFunctionDefinitionVersionsInput{
		// FunctionDefinitionId: *string, // Required
	}

	if len(_greengrassFunctionDefinitionId) > 0 {
		input.FunctionDefinitionId = aws.String(_greengrassFunctionDefinitionId)
	}
	if len(_greengrassMaxResults) > 0 {
		input.MaxResults = aws.String(_greengrassMaxResults)
	}
	if len(_greengrassNextToken) > 0 {
		input.NextToken = aws.String(_greengrassNextToken)
	}

	if resp, err := client.ListFunctionDefinitionVersions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of Lambda function definitions.
func greengrass_ListFunctionDefinitions(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.ListFunctionDefinitionsInput{}

	if len(_greengrassMaxResults) > 0 {
		input.MaxResults = aws.String(_greengrassMaxResults)
	}
	if len(_greengrassNextToken) > 0 {
		input.NextToken = aws.String(_greengrassNextToken)
	}

	if resp, err := client.ListFunctionDefinitions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the current CAs for a group.
func greengrass_ListGroupCertificateAuthorities(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.ListGroupCertificateAuthoritiesInput{
		// GroupId: *string, // Required
	}

	if len(_greengrassGroupId) > 0 {
		input.GroupId = aws.String(_greengrassGroupId)
	}

	if resp, err := client.ListGroupCertificateAuthorities(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the versions of a group.
func greengrass_ListGroupVersions(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.ListGroupVersionsInput{
		// GroupId: *string, // Required
	}

	if len(_greengrassGroupId) > 0 {
		input.GroupId = aws.String(_greengrassGroupId)
	}
	if len(_greengrassMaxResults) > 0 {
		input.MaxResults = aws.String(_greengrassMaxResults)
	}
	if len(_greengrassNextToken) > 0 {
		input.NextToken = aws.String(_greengrassNextToken)
	}

	if resp, err := client.ListGroupVersions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of groups.
func greengrass_ListGroups(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.ListGroupsInput{}

	if len(_greengrassMaxResults) > 0 {
		input.MaxResults = aws.String(_greengrassMaxResults)
	}
	if len(_greengrassNextToken) > 0 {
		input.NextToken = aws.String(_greengrassNextToken)
	}

	if resp, err := client.ListGroups(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the versions of a logger definition.
func greengrass_ListLoggerDefinitionVersions(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.ListLoggerDefinitionVersionsInput{
		// LoggerDefinitionId: *string, // Required
	}

	if len(_greengrassLoggerDefinitionId) > 0 {
		input.LoggerDefinitionId = aws.String(_greengrassLoggerDefinitionId)
	}
	if len(_greengrassMaxResults) > 0 {
		input.MaxResults = aws.String(_greengrassMaxResults)
	}
	if len(_greengrassNextToken) > 0 {
		input.NextToken = aws.String(_greengrassNextToken)
	}

	if resp, err := client.ListLoggerDefinitionVersions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of logger definitions.
func greengrass_ListLoggerDefinitions(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.ListLoggerDefinitionsInput{}

	if len(_greengrassMaxResults) > 0 {
		input.MaxResults = aws.String(_greengrassMaxResults)
	}
	if len(_greengrassNextToken) > 0 {
		input.NextToken = aws.String(_greengrassNextToken)
	}

	if resp, err := client.ListLoggerDefinitions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the versions of a resource definition.
func greengrass_ListResourceDefinitionVersions(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.ListResourceDefinitionVersionsInput{
		// ResourceDefinitionId: *string, // Required
	}

	if len(_greengrassResourceDefinitionId) > 0 {
		input.ResourceDefinitionId = aws.String(_greengrassResourceDefinitionId)
	}
	if len(_greengrassMaxResults) > 0 {
		input.MaxResults = aws.String(_greengrassMaxResults)
	}
	if len(_greengrassNextToken) > 0 {
		input.NextToken = aws.String(_greengrassNextToken)
	}

	if resp, err := client.ListResourceDefinitionVersions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of resource definitions.
func greengrass_ListResourceDefinitions(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.ListResourceDefinitionsInput{}

	if len(_greengrassMaxResults) > 0 {
		input.MaxResults = aws.String(_greengrassMaxResults)
	}
	if len(_greengrassNextToken) > 0 {
		input.NextToken = aws.String(_greengrassNextToken)
	}

	if resp, err := client.ListResourceDefinitions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the versions of a subscription definition.
func greengrass_ListSubscriptionDefinitionVersions(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.ListSubscriptionDefinitionVersionsInput{
		// SubscriptionDefinitionId: *string, // Required
	}

	if len(_greengrassSubscriptionDefinitionId) > 0 {
		input.SubscriptionDefinitionId = aws.String(_greengrassSubscriptionDefinitionId)
	}
	if len(_greengrassMaxResults) > 0 {
		input.MaxResults = aws.String(_greengrassMaxResults)
	}
	if len(_greengrassNextToken) > 0 {
		input.NextToken = aws.String(_greengrassNextToken)
	}

	if resp, err := client.ListSubscriptionDefinitionVersions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of subscription definitions.
func greengrass_ListSubscriptionDefinitions(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.ListSubscriptionDefinitionsInput{}

	if len(_greengrassMaxResults) > 0 {
		input.MaxResults = aws.String(_greengrassMaxResults)
	}
	if len(_greengrassNextToken) > 0 {
		input.NextToken = aws.String(_greengrassNextToken)
	}

	if resp, err := client.ListSubscriptionDefinitions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of resource tags for a resource arn.
func greengrass_ListTagsForResource(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_greengrassResourceArn) > 0 {
		input.ResourceArn = aws.String(_greengrassResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Resets a group's deployments.
func greengrass_ResetDeployments(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.ResetDeploymentsInput{
		// GroupId: *string, // Required
	}

	if len(_greengrassGroupId) > 0 {
		input.GroupId = aws.String(_greengrassGroupId)
	}
	if len(_greengrassAmznClientToken) > 0 {
		input.AmznClientToken = aws.String(_greengrassAmznClientToken)
	}
	if len(_greengrassForce) > 0 {
		if err := assignInputField(input, "Force", _greengrassForce); err != nil {
			log.Errorf("invalid --force: %s", err.Error())
			return
		}
	}

	if resp, err := client.ResetDeployments(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deploys multiple groups in one operation. This action starts the bulk
// deployment of a specified set of group versions. Each group version deployment
// will be triggered with an adaptive rate that has a fixed upper limit. We
// recommend that you include an ”X-Amzn-Client-Token” token in every
// ”StartBulkDeployment” request. These requests are idempotent with respect to
// the token and the request parameters.
func greengrass_StartBulkDeployment(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.StartBulkDeploymentInput{
		// ExecutionRoleArn: *string, // Required
		// InputFileUri: *string, // Required
	}

	if len(_greengrassExecutionRoleArn) > 0 {
		input.ExecutionRoleArn = aws.String(_greengrassExecutionRoleArn)
	}
	if len(_greengrassInputFileUri) > 0 {
		input.InputFileUri = aws.String(_greengrassInputFileUri)
	}
	if len(_greengrassAmznClientToken) > 0 {
		input.AmznClientToken = aws.String(_greengrassAmznClientToken)
	}
	if len(_greengrassTags) > 0 {
		if err := assignInputField(input, "Tags", _greengrassTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartBulkDeployment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops the execution of a bulk deployment. This action returns a status of
// ”Stopping” until the deployment is stopped. You cannot start a new bulk
// deployment while a previous deployment is in the ”Stopping” state. This action
// doesn't rollback completed deployments or cancel pending deployments.
func greengrass_StopBulkDeployment(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.StopBulkDeploymentInput{
		// BulkDeploymentId: *string, // Required
	}

	if len(_greengrassBulkDeploymentId) > 0 {
		input.BulkDeploymentId = aws.String(_greengrassBulkDeploymentId)
	}

	if resp, err := client.StopBulkDeployment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds tags to a Greengrass resource. Valid resources are 'Group',
// 'ConnectorDefinition', 'CoreDefinition', 'DeviceDefinition',
// 'FunctionDefinition', 'LoggerDefinition', 'SubscriptionDefinition',
// 'ResourceDefinition', and 'BulkDeployment'.
func greengrass_TagResource(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.TagResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_greengrassResourceArn) > 0 {
		input.ResourceArn = aws.String(_greengrassResourceArn)
	}
	if len(_greengrassTags) > 0 {
		if err := assignInputField(input, "Tags", _greengrassTags); err != nil {
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

// Remove resource tags from a Greengrass Resource.
func greengrass_UntagResource(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_greengrassResourceArn) > 0 {
		input.ResourceArn = aws.String(_greengrassResourceArn)
	}
	if len(_greengrassTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _greengrassTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the connectivity information for the core. Any devices that belong to
// the group which has this core will receive this information in order to find the
// location of the core and connect to it.
func greengrass_UpdateConnectivityInfo(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.UpdateConnectivityInfoInput{
		// ThingName: *string, // Required
	}

	if len(_greengrassThingName) > 0 {
		input.ThingName = aws.String(_greengrassThingName)
	}
	if len(_greengrassConnectivityInfo) > 0 {
		if err := assignInputField(input, "ConnectivityInfo", _greengrassConnectivityInfo); err != nil {
			log.Errorf("invalid --connectivity-info: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateConnectivityInfo(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a connector definition.
func greengrass_UpdateConnectorDefinition(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.UpdateConnectorDefinitionInput{
		// ConnectorDefinitionId: *string, // Required
	}

	if len(_greengrassConnectorDefinitionId) > 0 {
		input.ConnectorDefinitionId = aws.String(_greengrassConnectorDefinitionId)
	}
	if len(_greengrassName) > 0 {
		input.Name = aws.String(_greengrassName)
	}

	if resp, err := client.UpdateConnectorDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a core definition.
func greengrass_UpdateCoreDefinition(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.UpdateCoreDefinitionInput{
		// CoreDefinitionId: *string, // Required
	}

	if len(_greengrassCoreDefinitionId) > 0 {
		input.CoreDefinitionId = aws.String(_greengrassCoreDefinitionId)
	}
	if len(_greengrassName) > 0 {
		input.Name = aws.String(_greengrassName)
	}

	if resp, err := client.UpdateCoreDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a device definition.
func greengrass_UpdateDeviceDefinition(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.UpdateDeviceDefinitionInput{
		// DeviceDefinitionId: *string, // Required
	}

	if len(_greengrassDeviceDefinitionId) > 0 {
		input.DeviceDefinitionId = aws.String(_greengrassDeviceDefinitionId)
	}
	if len(_greengrassName) > 0 {
		input.Name = aws.String(_greengrassName)
	}

	if resp, err := client.UpdateDeviceDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a Lambda function definition.
func greengrass_UpdateFunctionDefinition(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.UpdateFunctionDefinitionInput{
		// FunctionDefinitionId: *string, // Required
	}

	if len(_greengrassFunctionDefinitionId) > 0 {
		input.FunctionDefinitionId = aws.String(_greengrassFunctionDefinitionId)
	}
	if len(_greengrassName) > 0 {
		input.Name = aws.String(_greengrassName)
	}

	if resp, err := client.UpdateFunctionDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a group.
func greengrass_UpdateGroup(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.UpdateGroupInput{
		// GroupId: *string, // Required
	}

	if len(_greengrassGroupId) > 0 {
		input.GroupId = aws.String(_greengrassGroupId)
	}
	if len(_greengrassName) > 0 {
		input.Name = aws.String(_greengrassName)
	}

	if resp, err := client.UpdateGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the Certificate expiry time for a group.
func greengrass_UpdateGroupCertificateConfiguration(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.UpdateGroupCertificateConfigurationInput{
		// GroupId: *string, // Required
	}

	if len(_greengrassGroupId) > 0 {
		input.GroupId = aws.String(_greengrassGroupId)
	}
	if len(_greengrassCertificateExpiryInMilliseconds) > 0 {
		input.CertificateExpiryInMilliseconds = aws.String(_greengrassCertificateExpiryInMilliseconds)
	}

	if resp, err := client.UpdateGroupCertificateConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a logger definition.
func greengrass_UpdateLoggerDefinition(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.UpdateLoggerDefinitionInput{
		// LoggerDefinitionId: *string, // Required
	}

	if len(_greengrassLoggerDefinitionId) > 0 {
		input.LoggerDefinitionId = aws.String(_greengrassLoggerDefinitionId)
	}
	if len(_greengrassName) > 0 {
		input.Name = aws.String(_greengrassName)
	}

	if resp, err := client.UpdateLoggerDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a resource definition.
func greengrass_UpdateResourceDefinition(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.UpdateResourceDefinitionInput{
		// ResourceDefinitionId: *string, // Required
	}

	if len(_greengrassResourceDefinitionId) > 0 {
		input.ResourceDefinitionId = aws.String(_greengrassResourceDefinitionId)
	}
	if len(_greengrassName) > 0 {
		input.Name = aws.String(_greengrassName)
	}

	if resp, err := client.UpdateResourceDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a subscription definition.
func greengrass_UpdateSubscriptionDefinition(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.UpdateSubscriptionDefinitionInput{
		// SubscriptionDefinitionId: *string, // Required
	}

	if len(_greengrassSubscriptionDefinitionId) > 0 {
		input.SubscriptionDefinitionId = aws.String(_greengrassSubscriptionDefinitionId)
	}
	if len(_greengrassName) > 0 {
		input.Name = aws.String(_greengrassName)
	}

	if resp, err := client.UpdateSubscriptionDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the runtime configuration of a thing.
func greengrass_UpdateThingRuntimeConfiguration(cfg aws.Config, client *greengrass.Client) {
	input := &greengrass.UpdateThingRuntimeConfigurationInput{
		// ThingName: *string, // Required
	}

	if len(_greengrassThingName) > 0 {
		input.ThingName = aws.String(_greengrassThingName)
	}
	if len(_greengrassTelemetryConfiguration) > 0 {
		if err := assignInputField(input, "TelemetryConfiguration", _greengrassTelemetryConfiguration); err != nil {
			log.Errorf("invalid --telemetry-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateThingRuntimeConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_greengrassCmd)
	_greengrassCmd.Flags().SortFlags = false

	_greengrassCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_greengrassCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_greengrassCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_greengrassCmd.Flags().StringVarP(&_greengrassAmznClientToken, "amzn-client-token", "", "", "Amzn Client Token")
	_greengrassCmd.Flags().StringVarP(&_greengrassBulkDeploymentId, "bulk-deployment-id", "", "", "Bulk Deployment ID")
	_greengrassCmd.Flags().StringVarP(&_greengrassCertificateAuthorityId, "certificate-authority-id", "", "", "Certificate Authority ID")
	_greengrassCmd.Flags().StringVarP(&_greengrassCertificateExpiryInMilliseconds, "certificate-expiry-in-milliseconds", "", "", "Certificate Expiry In Milliseconds")
	_greengrassCmd.Flags().StringVarP(&_greengrassConnectivityInfo, "connectivity-info", "", "", "Connectivity Info")
	_greengrassCmd.Flags().StringVarP(&_greengrassConnectorDefinitionId, "connector-definition-id", "", "", "Connector Definition ID")
	_greengrassCmd.Flags().StringVarP(&_greengrassConnectorDefinitionVersionArn, "connector-definition-version-arn", "", "", "Connector Definition Version ARN")
	_greengrassCmd.Flags().StringVarP(&_greengrassConnectorDefinitionVersionId, "connector-definition-version-id", "", "", "Connector Definition Version ID")
	_greengrassCmd.Flags().StringVarP(&_greengrassConnectors, "connectors", "", "", "Connectors")
	_greengrassCmd.Flags().StringVarP(&_greengrassCoreDefinitionId, "core-definition-id", "", "", "Core Definition ID")
	_greengrassCmd.Flags().StringVarP(&_greengrassCoreDefinitionVersionArn, "core-definition-version-arn", "", "", "Core Definition Version ARN")
	_greengrassCmd.Flags().StringVarP(&_greengrassCoreDefinitionVersionId, "core-definition-version-id", "", "", "Core Definition Version ID")
	_greengrassCmd.Flags().StringVarP(&_greengrassCores, "cores", "", "", "Cores")
	_greengrassCmd.Flags().StringVarP(&_greengrassDefaultConfig, "default-config", "", "", "Default Config")
	_greengrassCmd.Flags().StringVarP(&_greengrassDeploymentId, "deployment-id", "", "", "Deployment ID")
	_greengrassCmd.Flags().StringVarP(&_greengrassDeploymentType, "deployment-type", "", "", "Deployment Type")
	_greengrassCmd.Flags().StringVarP(&_greengrassDeviceDefinitionId, "device-definition-id", "", "", "Device Definition ID")
	_greengrassCmd.Flags().StringVarP(&_greengrassDeviceDefinitionVersionArn, "device-definition-version-arn", "", "", "Device Definition Version ARN")
	_greengrassCmd.Flags().StringVarP(&_greengrassDeviceDefinitionVersionId, "device-definition-version-id", "", "", "Device Definition Version ID")
	_greengrassCmd.Flags().StringVarP(&_greengrassDevices, "devices", "", "", "Devices")
	_greengrassCmd.Flags().StringVarP(&_greengrassExecutionRoleArn, "execution-role-arn", "", "", "Execution Role ARN")
	_greengrassCmd.Flags().StringVarP(&_greengrassForce, "force", "", "", "Force")
	_greengrassCmd.Flags().StringVarP(&_greengrassFunctionDefinitionId, "function-definition-id", "", "", "Function Definition ID")
	_greengrassCmd.Flags().StringVarP(&_greengrassFunctionDefinitionVersionArn, "function-definition-version-arn", "", "", "Function Definition Version ARN")
	_greengrassCmd.Flags().StringVarP(&_greengrassFunctionDefinitionVersionId, "function-definition-version-id", "", "", "Function Definition Version ID")
	_greengrassCmd.Flags().StringVarP(&_greengrassFunctions, "functions", "", "", "Functions")
	_greengrassCmd.Flags().StringVarP(&_greengrassGroupId, "group-id", "", "", "Group ID")
	_greengrassCmd.Flags().StringVarP(&_greengrassGroupVersionId, "group-version-id", "", "", "Group Version ID")
	_greengrassCmd.Flags().StringVarP(&_greengrassInitialVersion, "initial-version", "", "", "Initial Version")
	_greengrassCmd.Flags().StringVarP(&_greengrassInputFileUri, "input-file-uri", "", "", "Input File URI")
	_greengrassCmd.Flags().StringVarP(&_greengrassLoggerDefinitionId, "logger-definition-id", "", "", "Logger Definition ID")
	_greengrassCmd.Flags().StringVarP(&_greengrassLoggerDefinitionVersionArn, "logger-definition-version-arn", "", "", "Logger Definition Version ARN")
	_greengrassCmd.Flags().StringVarP(&_greengrassLoggerDefinitionVersionId, "logger-definition-version-id", "", "", "Logger Definition Version ID")
	_greengrassCmd.Flags().StringVarP(&_greengrassLoggers, "loggers", "", "", "Loggers")
	_greengrassCmd.Flags().StringVarP(&_greengrassMaxResults, "max-results", "", "", "Max Results")
	_greengrassCmd.Flags().StringVarP(&_greengrassName, "name", "", "", "Name")
	_greengrassCmd.Flags().StringVarP(&_greengrassNextToken, "next-token", "", "", "Next Token")
	_greengrassCmd.Flags().StringVarP(&_greengrassResourceArn, "resource-arn", "", "", "Resource ARN")
	_greengrassCmd.Flags().StringVarP(&_greengrassResourceDefinitionId, "resource-definition-id", "", "", "Resource Definition ID")
	_greengrassCmd.Flags().StringVarP(&_greengrassResourceDefinitionVersionArn, "resource-definition-version-arn", "", "", "Resource Definition Version ARN")
	_greengrassCmd.Flags().StringVarP(&_greengrassResourceDefinitionVersionId, "resource-definition-version-id", "", "", "Resource Definition Version ID")
	_greengrassCmd.Flags().StringVarP(&_greengrassResources, "resources", "", "", "Resources")
	_greengrassCmd.Flags().StringVarP(&_greengrassRoleArn, "role-arn", "", "", "Role ARN")
	_greengrassCmd.Flags().StringVarP(&_greengrassS3UrlSignerRole, "s3-url-signer-role", "", "", "S3 URL Signer Role")
	_greengrassCmd.Flags().StringVarP(&_greengrassSoftwareToUpdate, "software-to-update", "", "", "Software To Update")
	_greengrassCmd.Flags().StringVarP(&_greengrassSubscriptionDefinitionId, "subscription-definition-id", "", "", "Subscription Definition ID")
	_greengrassCmd.Flags().StringVarP(&_greengrassSubscriptionDefinitionVersionArn, "subscription-definition-version-arn", "", "", "Subscription Definition Version ARN")
	_greengrassCmd.Flags().StringVarP(&_greengrassSubscriptionDefinitionVersionId, "subscription-definition-version-id", "", "", "Subscription Definition Version ID")
	_greengrassCmd.Flags().StringVarP(&_greengrassSubscriptions, "subscriptions", "", "", "Subscriptions")
	_greengrassCmd.Flags().StringSliceVarP(&_greengrassTagKeys, "tag-keys", "", nil, "Tag Keys")
	_greengrassCmd.Flags().StringVarP(&_greengrassTags, "tags", "", "", "Tags")
	_greengrassCmd.Flags().StringVarP(&_greengrassTelemetryConfiguration, "telemetry-configuration", "", "", "Telemetry Configuration")
	_greengrassCmd.Flags().StringVarP(&_greengrassThingName, "thing-name", "", "", "Thing Name")
	_greengrassCmd.Flags().StringVarP(&_greengrassUpdateAgentLogLevel, "update-agent-log-level", "", "", "Update Agent Log Level")
	_greengrassCmd.Flags().StringSliceVarP(&_greengrassUpdateTargets, "update-targets", "", nil, "Update Targets")
	_greengrassCmd.Flags().StringVarP(&_greengrassUpdateTargetsArchitecture, "update-targets-architecture", "", "", "Update Targets Architecture")
	_greengrassCmd.Flags().StringVarP(&_greengrassUpdateTargetsOperatingSystem, "update-targets-operating-system", "", "", "Update Targets Operating System")

	_greengrassCmd.Flags().BoolVarP(&_greengrassAssociateRoleToGroup, "associate-role-to-group", "", false, "Associate Role To Group")
	_greengrassCmd.Flags().BoolVarP(&_greengrassAssociateServiceRoleToAccount, "associate-service-role-to-account", "", false, "Associate Service Role To Account")
	_greengrassCmd.Flags().BoolVarP(&_greengrassCreateConnectorDefinition, "create-connector-definition", "", false, "Create Connector Definition")
	_greengrassCmd.Flags().BoolVarP(&_greengrassCreateConnectorDefinitionVersion, "create-connector-definition-version", "", false, "Create Connector Definition Version")
	_greengrassCmd.Flags().BoolVarP(&_greengrassCreateCoreDefinition, "create-core-definition", "", false, "Create Core Definition")
	_greengrassCmd.Flags().BoolVarP(&_greengrassCreateCoreDefinitionVersion, "create-core-definition-version", "", false, "Create Core Definition Version")
	_greengrassCmd.Flags().BoolVarP(&_greengrassCreateDeployment, "create-deployment", "", false, "Create Deployment")
	_greengrassCmd.Flags().BoolVarP(&_greengrassCreateDeviceDefinition, "create-device-definition", "", false, "Create Device Definition")
	_greengrassCmd.Flags().BoolVarP(&_greengrassCreateDeviceDefinitionVersion, "create-device-definition-version", "", false, "Create Device Definition Version")
	_greengrassCmd.Flags().BoolVarP(&_greengrassCreateFunctionDefinition, "create-function-definition", "", false, "Create Function Definition")
	_greengrassCmd.Flags().BoolVarP(&_greengrassCreateFunctionDefinitionVersion, "create-function-definition-version", "", false, "Create Function Definition Version")
	_greengrassCmd.Flags().BoolVarP(&_greengrassCreateGroup, "create-group", "", false, "Create Group")
	_greengrassCmd.Flags().BoolVarP(&_greengrassCreateGroupCertificateAuthority, "create-group-certificate-authority", "", false, "Create Group Certificate Authority")
	_greengrassCmd.Flags().BoolVarP(&_greengrassCreateGroupVersion, "create-group-version", "", false, "Create Group Version")
	_greengrassCmd.Flags().BoolVarP(&_greengrassCreateLoggerDefinition, "create-logger-definition", "", false, "Create Logger Definition")
	_greengrassCmd.Flags().BoolVarP(&_greengrassCreateLoggerDefinitionVersion, "create-logger-definition-version", "", false, "Create Logger Definition Version")
	_greengrassCmd.Flags().BoolVarP(&_greengrassCreateResourceDefinition, "create-resource-definition", "", false, "Create Resource Definition")
	_greengrassCmd.Flags().BoolVarP(&_greengrassCreateResourceDefinitionVersion, "create-resource-definition-version", "", false, "Create Resource Definition Version")
	_greengrassCmd.Flags().BoolVarP(&_greengrassCreateSoftwareUpdateJob, "create-software-update-job", "", false, "Create Software Update Job")
	_greengrassCmd.Flags().BoolVarP(&_greengrassCreateSubscriptionDefinition, "create-subscription-definition", "", false, "Create Subscription Definition")
	_greengrassCmd.Flags().BoolVarP(&_greengrassCreateSubscriptionDefinitionVersion, "create-subscription-definition-version", "", false, "Create Subscription Definition Version")
	_greengrassCmd.Flags().BoolVarP(&_greengrassDeleteConnectorDefinition, "delete-connector-definition", "", false, "Delete Connector Definition")
	_greengrassCmd.Flags().BoolVarP(&_greengrassDeleteCoreDefinition, "delete-core-definition", "", false, "Delete Core Definition")
	_greengrassCmd.Flags().BoolVarP(&_greengrassDeleteDeviceDefinition, "delete-device-definition", "", false, "Delete Device Definition")
	_greengrassCmd.Flags().BoolVarP(&_greengrassDeleteFunctionDefinition, "delete-function-definition", "", false, "Delete Function Definition")
	_greengrassCmd.Flags().BoolVarP(&_greengrassDeleteGroup, "delete-group", "", false, "Delete Group")
	_greengrassCmd.Flags().BoolVarP(&_greengrassDeleteLoggerDefinition, "delete-logger-definition", "", false, "Delete Logger Definition")
	_greengrassCmd.Flags().BoolVarP(&_greengrassDeleteResourceDefinition, "delete-resource-definition", "", false, "Delete Resource Definition")
	_greengrassCmd.Flags().BoolVarP(&_greengrassDeleteSubscriptionDefinition, "delete-subscription-definition", "", false, "Delete Subscription Definition")
	_greengrassCmd.Flags().BoolVarP(&_greengrassDisassociateRoleFromGroup, "disassociate-role-from-group", "", false, "Disassociate Role From Group")
	_greengrassCmd.Flags().BoolVarP(&_greengrassDisassociateServiceRoleFromAccount, "disassociate-service-role-from-account", "", false, "Disassociate Service Role From Account")
	_greengrassCmd.Flags().BoolVarP(&_greengrassGetAssociatedRole, "get-associated-role", "", false, "Get Associated Role")
	_greengrassCmd.Flags().BoolVarP(&_greengrassGetBulkDeploymentStatus, "get-bulk-deployment-status", "", false, "Get Bulk Deployment Status")
	_greengrassCmd.Flags().BoolVarP(&_greengrassGetConnectivityInfo, "get-connectivity-info", "", false, "Get Connectivity Info")
	_greengrassCmd.Flags().BoolVarP(&_greengrassGetConnectorDefinition, "get-connector-definition", "", false, "Get Connector Definition")
	_greengrassCmd.Flags().BoolVarP(&_greengrassGetConnectorDefinitionVersion, "get-connector-definition-version", "", false, "Get Connector Definition Version")
	_greengrassCmd.Flags().BoolVarP(&_greengrassGetCoreDefinition, "get-core-definition", "", false, "Get Core Definition")
	_greengrassCmd.Flags().BoolVarP(&_greengrassGetCoreDefinitionVersion, "get-core-definition-version", "", false, "Get Core Definition Version")
	_greengrassCmd.Flags().BoolVarP(&_greengrassGetDeploymentStatus, "get-deployment-status", "", false, "Get Deployment Status")
	_greengrassCmd.Flags().BoolVarP(&_greengrassGetDeviceDefinition, "get-device-definition", "", false, "Get Device Definition")
	_greengrassCmd.Flags().BoolVarP(&_greengrassGetDeviceDefinitionVersion, "get-device-definition-version", "", false, "Get Device Definition Version")
	_greengrassCmd.Flags().BoolVarP(&_greengrassGetFunctionDefinition, "get-function-definition", "", false, "Get Function Definition")
	_greengrassCmd.Flags().BoolVarP(&_greengrassGetFunctionDefinitionVersion, "get-function-definition-version", "", false, "Get Function Definition Version")
	_greengrassCmd.Flags().BoolVarP(&_greengrassGetGroup, "get-group", "", false, "Get Group")
	_greengrassCmd.Flags().BoolVarP(&_greengrassGetGroupCertificateAuthority, "get-group-certificate-authority", "", false, "Get Group Certificate Authority")
	_greengrassCmd.Flags().BoolVarP(&_greengrassGetGroupCertificateConfiguration, "get-group-certificate-configuration", "", false, "Get Group Certificate Configuration")
	_greengrassCmd.Flags().BoolVarP(&_greengrassGetGroupVersion, "get-group-version", "", false, "Get Group Version")
	_greengrassCmd.Flags().BoolVarP(&_greengrassGetLoggerDefinition, "get-logger-definition", "", false, "Get Logger Definition")
	_greengrassCmd.Flags().BoolVarP(&_greengrassGetLoggerDefinitionVersion, "get-logger-definition-version", "", false, "Get Logger Definition Version")
	_greengrassCmd.Flags().BoolVarP(&_greengrassGetResourceDefinition, "get-resource-definition", "", false, "Get Resource Definition")
	_greengrassCmd.Flags().BoolVarP(&_greengrassGetResourceDefinitionVersion, "get-resource-definition-version", "", false, "Get Resource Definition Version")
	_greengrassCmd.Flags().BoolVarP(&_greengrassGetServiceRoleForAccount, "get-service-role-for-account", "", false, "Get Service Role For Account")
	_greengrassCmd.Flags().BoolVarP(&_greengrassGetSubscriptionDefinition, "get-subscription-definition", "", false, "Get Subscription Definition")
	_greengrassCmd.Flags().BoolVarP(&_greengrassGetSubscriptionDefinitionVersion, "get-subscription-definition-version", "", false, "Get Subscription Definition Version")
	_greengrassCmd.Flags().BoolVarP(&_greengrassGetThingRuntimeConfiguration, "get-thing-runtime-configuration", "", false, "Get Thing Runtime Configuration")
	_greengrassCmd.Flags().BoolVarP(&_greengrassListBulkDeploymentDetailedReports, "list-bulk-deployment-detailed-reports", "", false, "List Bulk Deployment Detailed Reports")
	_greengrassCmd.Flags().BoolVarP(&_greengrassListBulkDeployments, "list-bulk-deployments", "", false, "List Bulk Deployments")
	_greengrassCmd.Flags().BoolVarP(&_greengrassListConnectorDefinitionVersions, "list-connector-definition-versions", "", false, "List Connector Definition Versions")
	_greengrassCmd.Flags().BoolVarP(&_greengrassListConnectorDefinitions, "list-connector-definitions", "", false, "List Connector Definitions")
	_greengrassCmd.Flags().BoolVarP(&_greengrassListCoreDefinitionVersions, "list-core-definition-versions", "", false, "List Core Definition Versions")
	_greengrassCmd.Flags().BoolVarP(&_greengrassListCoreDefinitions, "list-core-definitions", "", false, "List Core Definitions")
	_greengrassCmd.Flags().BoolVarP(&_greengrassListDeployments, "list-deployments", "", false, "List Deployments")
	_greengrassCmd.Flags().BoolVarP(&_greengrassListDeviceDefinitionVersions, "list-device-definition-versions", "", false, "List Device Definition Versions")
	_greengrassCmd.Flags().BoolVarP(&_greengrassListDeviceDefinitions, "list-device-definitions", "", false, "List Device Definitions")
	_greengrassCmd.Flags().BoolVarP(&_greengrassListFunctionDefinitionVersions, "list-function-definition-versions", "", false, "List Function Definition Versions")
	_greengrassCmd.Flags().BoolVarP(&_greengrassListFunctionDefinitions, "list-function-definitions", "", false, "List Function Definitions")
	_greengrassCmd.Flags().BoolVarP(&_greengrassListGroupCertificateAuthorities, "list-group-certificate-authorities", "", false, "List Group Certificate Authorities")
	_greengrassCmd.Flags().BoolVarP(&_greengrassListGroupVersions, "list-group-versions", "", false, "List Group Versions")
	_greengrassCmd.Flags().BoolVarP(&_greengrassListGroups, "list-groups", "", false, "List Groups")
	_greengrassCmd.Flags().BoolVarP(&_greengrassListLoggerDefinitionVersions, "list-logger-definition-versions", "", false, "List Logger Definition Versions")
	_greengrassCmd.Flags().BoolVarP(&_greengrassListLoggerDefinitions, "list-logger-definitions", "", false, "List Logger Definitions")
	_greengrassCmd.Flags().BoolVarP(&_greengrassListResourceDefinitionVersions, "list-resource-definition-versions", "", false, "List Resource Definition Versions")
	_greengrassCmd.Flags().BoolVarP(&_greengrassListResourceDefinitions, "list-resource-definitions", "", false, "List Resource Definitions")
	_greengrassCmd.Flags().BoolVarP(&_greengrassListSubscriptionDefinitionVersions, "list-subscription-definition-versions", "", false, "List Subscription Definition Versions")
	_greengrassCmd.Flags().BoolVarP(&_greengrassListSubscriptionDefinitions, "list-subscription-definitions", "", false, "List Subscription Definitions")
	_greengrassCmd.Flags().BoolVarP(&_greengrassListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_greengrassCmd.Flags().BoolVarP(&_greengrassResetDeployments, "reset-deployments", "", false, "Reset Deployments")
	_greengrassCmd.Flags().BoolVarP(&_greengrassStartBulkDeployment, "start-bulk-deployment", "", false, "Start Bulk Deployment")
	_greengrassCmd.Flags().BoolVarP(&_greengrassStopBulkDeployment, "stop-bulk-deployment", "", false, "Stop Bulk Deployment")
	_greengrassCmd.Flags().BoolVarP(&_greengrassTagResource, "tag-resource", "", false, "Tag Resource")
	_greengrassCmd.Flags().BoolVarP(&_greengrassUntagResource, "untag-resource", "", false, "Untag Resource")
	_greengrassCmd.Flags().BoolVarP(&_greengrassUpdateConnectivityInfo, "update-connectivity-info", "", false, "Update Connectivity Info")
	_greengrassCmd.Flags().BoolVarP(&_greengrassUpdateConnectorDefinition, "update-connector-definition", "", false, "Update Connector Definition")
	_greengrassCmd.Flags().BoolVarP(&_greengrassUpdateCoreDefinition, "update-core-definition", "", false, "Update Core Definition")
	_greengrassCmd.Flags().BoolVarP(&_greengrassUpdateDeviceDefinition, "update-device-definition", "", false, "Update Device Definition")
	_greengrassCmd.Flags().BoolVarP(&_greengrassUpdateFunctionDefinition, "update-function-definition", "", false, "Update Function Definition")
	_greengrassCmd.Flags().BoolVarP(&_greengrassUpdateGroup, "update-group", "", false, "Update Group")
	_greengrassCmd.Flags().BoolVarP(&_greengrassUpdateGroupCertificateConfiguration, "update-group-certificate-configuration", "", false, "Update Group Certificate Configuration")
	_greengrassCmd.Flags().BoolVarP(&_greengrassUpdateLoggerDefinition, "update-logger-definition", "", false, "Update Logger Definition")
	_greengrassCmd.Flags().BoolVarP(&_greengrassUpdateResourceDefinition, "update-resource-definition", "", false, "Update Resource Definition")
	_greengrassCmd.Flags().BoolVarP(&_greengrassUpdateSubscriptionDefinition, "update-subscription-definition", "", false, "Update Subscription Definition")
	_greengrassCmd.Flags().BoolVarP(&_greengrassUpdateThingRuntimeConfiguration, "update-thing-runtime-configuration", "", false, "Update Thing Runtime Configuration")

}
