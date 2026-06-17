package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appstream"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// appstreamCmd represents the appstream command
var _appstreamCmd = &cobra.Command{
	Use:   "appstream",
	Short: "AWS appstream CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := appstream.NewFromConfig(cfg)
		if _appstreamAssociateAppBlockBuilderAppBlock {
			appstream_AssociateAppBlockBuilderAppBlock(cfg, client)
			return
		}
		if _appstreamAssociateApplicationFleet {
			appstream_AssociateApplicationFleet(cfg, client)
			return
		}
		if _appstreamAssociateApplicationToEntitlement {
			appstream_AssociateApplicationToEntitlement(cfg, client)
			return
		}
		if _appstreamAssociateFleet {
			appstream_AssociateFleet(cfg, client)
			return
		}
		if _appstreamAssociateSoftwareToImageBuilder {
			appstream_AssociateSoftwareToImageBuilder(cfg, client)
			return
		}
		if _appstreamBatchAssociateUserStack {
			appstream_BatchAssociateUserStack(cfg, client)
			return
		}
		if _appstreamBatchDisassociateUserStack {
			appstream_BatchDisassociateUserStack(cfg, client)
			return
		}
		if _appstreamCopyImage {
			appstream_CopyImage(cfg, client)
			return
		}
		if _appstreamCreateAppBlock {
			appstream_CreateAppBlock(cfg, client)
			return
		}
		if _appstreamCreateAppBlockBuilder {
			appstream_CreateAppBlockBuilder(cfg, client)
			return
		}
		if _appstreamCreateAppBlockBuilderStreamingURL {
			appstream_CreateAppBlockBuilderStreamingURL(cfg, client)
			return
		}
		if _appstreamCreateApplication {
			appstream_CreateApplication(cfg, client)
			return
		}
		if _appstreamCreateDirectoryConfig {
			appstream_CreateDirectoryConfig(cfg, client)
			return
		}
		if _appstreamCreateEntitlement {
			appstream_CreateEntitlement(cfg, client)
			return
		}
		if _appstreamCreateExportImageTask {
			appstream_CreateExportImageTask(cfg, client)
			return
		}
		if _appstreamCreateFleet {
			appstream_CreateFleet(cfg, client)
			return
		}
		if _appstreamCreateImageBuilder {
			appstream_CreateImageBuilder(cfg, client)
			return
		}
		if _appstreamCreateImageBuilderStreamingURL {
			appstream_CreateImageBuilderStreamingURL(cfg, client)
			return
		}
		if _appstreamCreateImportedImage {
			appstream_CreateImportedImage(cfg, client)
			return
		}
		if _appstreamCreateStack {
			appstream_CreateStack(cfg, client)
			return
		}
		if _appstreamCreateStreamingURL {
			appstream_CreateStreamingURL(cfg, client)
			return
		}
		if _appstreamCreateThemeForStack {
			appstream_CreateThemeForStack(cfg, client)
			return
		}
		if _appstreamCreateUpdatedImage {
			appstream_CreateUpdatedImage(cfg, client)
			return
		}
		if _appstreamCreateUsageReportSubscription {
			appstream_CreateUsageReportSubscription(cfg, client)
			return
		}
		if _appstreamCreateUser {
			appstream_CreateUser(cfg, client)
			return
		}
		if _appstreamDeleteAppBlock {
			appstream_DeleteAppBlock(cfg, client)
			return
		}
		if _appstreamDeleteAppBlockBuilder {
			appstream_DeleteAppBlockBuilder(cfg, client)
			return
		}
		if _appstreamDeleteApplication {
			appstream_DeleteApplication(cfg, client)
			return
		}
		if _appstreamDeleteDirectoryConfig {
			appstream_DeleteDirectoryConfig(cfg, client)
			return
		}
		if _appstreamDeleteEntitlement {
			appstream_DeleteEntitlement(cfg, client)
			return
		}
		if _appstreamDeleteFleet {
			appstream_DeleteFleet(cfg, client)
			return
		}
		if _appstreamDeleteImage {
			appstream_DeleteImage(cfg, client)
			return
		}
		if _appstreamDeleteImageBuilder {
			appstream_DeleteImageBuilder(cfg, client)
			return
		}
		if _appstreamDeleteImagePermissions {
			appstream_DeleteImagePermissions(cfg, client)
			return
		}
		if _appstreamDeleteStack {
			appstream_DeleteStack(cfg, client)
			return
		}
		if _appstreamDeleteThemeForStack {
			appstream_DeleteThemeForStack(cfg, client)
			return
		}
		if _appstreamDeleteUsageReportSubscription {
			appstream_DeleteUsageReportSubscription(cfg, client)
			return
		}
		if _appstreamDeleteUser {
			appstream_DeleteUser(cfg, client)
			return
		}
		if _appstreamDescribeAppBlockBuilderAppBlockAssociations {
			appstream_DescribeAppBlockBuilderAppBlockAssociations(cfg, client)
			return
		}
		if _appstreamDescribeAppBlockBuilders {
			appstream_DescribeAppBlockBuilders(cfg, client)
			return
		}
		if _appstreamDescribeAppBlocks {
			appstream_DescribeAppBlocks(cfg, client)
			return
		}
		if _appstreamDescribeAppLicenseUsage {
			appstream_DescribeAppLicenseUsage(cfg, client)
			return
		}
		if _appstreamDescribeApplicationFleetAssociations {
			appstream_DescribeApplicationFleetAssociations(cfg, client)
			return
		}
		if _appstreamDescribeApplications {
			appstream_DescribeApplications(cfg, client)
			return
		}
		if _appstreamDescribeDirectoryConfigs {
			appstream_DescribeDirectoryConfigs(cfg, client)
			return
		}
		if _appstreamDescribeEntitlements {
			appstream_DescribeEntitlements(cfg, client)
			return
		}
		if _appstreamDescribeFleets {
			appstream_DescribeFleets(cfg, client)
			return
		}
		if _appstreamDescribeImageBuilders {
			appstream_DescribeImageBuilders(cfg, client)
			return
		}
		if _appstreamDescribeImagePermissions {
			appstream_DescribeImagePermissions(cfg, client)
			return
		}
		if _appstreamDescribeImages {
			appstream_DescribeImages(cfg, client)
			return
		}
		if _appstreamDescribeSessions {
			appstream_DescribeSessions(cfg, client)
			return
		}
		if _appstreamDescribeSoftwareAssociations {
			appstream_DescribeSoftwareAssociations(cfg, client)
			return
		}
		if _appstreamDescribeStacks {
			appstream_DescribeStacks(cfg, client)
			return
		}
		if _appstreamDescribeThemeForStack {
			appstream_DescribeThemeForStack(cfg, client)
			return
		}
		if _appstreamDescribeUsageReportSubscriptions {
			appstream_DescribeUsageReportSubscriptions(cfg, client)
			return
		}
		if _appstreamDescribeUserStackAssociations {
			appstream_DescribeUserStackAssociations(cfg, client)
			return
		}
		if _appstreamDescribeUsers {
			appstream_DescribeUsers(cfg, client)
			return
		}
		if _appstreamDisableUser {
			appstream_DisableUser(cfg, client)
			return
		}
		if _appstreamDisassociateAppBlockBuilderAppBlock {
			appstream_DisassociateAppBlockBuilderAppBlock(cfg, client)
			return
		}
		if _appstreamDisassociateApplicationFleet {
			appstream_DisassociateApplicationFleet(cfg, client)
			return
		}
		if _appstreamDisassociateApplicationFromEntitlement {
			appstream_DisassociateApplicationFromEntitlement(cfg, client)
			return
		}
		if _appstreamDisassociateFleet {
			appstream_DisassociateFleet(cfg, client)
			return
		}
		if _appstreamDisassociateSoftwareFromImageBuilder {
			appstream_DisassociateSoftwareFromImageBuilder(cfg, client)
			return
		}
		if _appstreamEnableUser {
			appstream_EnableUser(cfg, client)
			return
		}
		if _appstreamExpireSession {
			appstream_ExpireSession(cfg, client)
			return
		}
		if _appstreamGetExportImageTask {
			appstream_GetExportImageTask(cfg, client)
			return
		}
		if _appstreamListAssociatedFleets {
			appstream_ListAssociatedFleets(cfg, client)
			return
		}
		if _appstreamListAssociatedStacks {
			appstream_ListAssociatedStacks(cfg, client)
			return
		}
		if _appstreamListEntitledApplications {
			appstream_ListEntitledApplications(cfg, client)
			return
		}
		if _appstreamListExportImageTasks {
			appstream_ListExportImageTasks(cfg, client)
			return
		}
		if _appstreamListTagsForResource {
			appstream_ListTagsForResource(cfg, client)
			return
		}
		if _appstreamStartAppBlockBuilder {
			appstream_StartAppBlockBuilder(cfg, client)
			return
		}
		if _appstreamStartFleet {
			appstream_StartFleet(cfg, client)
			return
		}
		if _appstreamStartImageBuilder {
			appstream_StartImageBuilder(cfg, client)
			return
		}
		if _appstreamStartSoftwareDeploymentToImageBuilder {
			appstream_StartSoftwareDeploymentToImageBuilder(cfg, client)
			return
		}
		if _appstreamStopAppBlockBuilder {
			appstream_StopAppBlockBuilder(cfg, client)
			return
		}
		if _appstreamStopFleet {
			appstream_StopFleet(cfg, client)
			return
		}
		if _appstreamStopImageBuilder {
			appstream_StopImageBuilder(cfg, client)
			return
		}
		if _appstreamTagResource {
			appstream_TagResource(cfg, client)
			return
		}
		if _appstreamUntagResource {
			appstream_UntagResource(cfg, client)
			return
		}
		if _appstreamUpdateAppBlockBuilder {
			appstream_UpdateAppBlockBuilder(cfg, client)
			return
		}
		if _appstreamUpdateApplication {
			appstream_UpdateApplication(cfg, client)
			return
		}
		if _appstreamUpdateDirectoryConfig {
			appstream_UpdateDirectoryConfig(cfg, client)
			return
		}
		if _appstreamUpdateEntitlement {
			appstream_UpdateEntitlement(cfg, client)
			return
		}
		if _appstreamUpdateFleet {
			appstream_UpdateFleet(cfg, client)
			return
		}
		if _appstreamUpdateImagePermissions {
			appstream_UpdateImagePermissions(cfg, client)
			return
		}
		if _appstreamUpdateStack {
			appstream_UpdateStack(cfg, client)
			return
		}
		if _appstreamUpdateThemeForStack {
			appstream_UpdateThemeForStack(cfg, client)
			return
		}

	},
}

var (
	_appstreamAssociateAppBlockBuilderAppBlock            bool
	_appstreamAssociateApplicationFleet                   bool
	_appstreamAssociateApplicationToEntitlement           bool
	_appstreamAssociateFleet                              bool
	_appstreamAssociateSoftwareToImageBuilder             bool
	_appstreamBatchAssociateUserStack                     bool
	_appstreamBatchDisassociateUserStack                  bool
	_appstreamCopyImage                                   bool
	_appstreamCreateAppBlock                              bool
	_appstreamCreateAppBlockBuilder                       bool
	_appstreamCreateAppBlockBuilderStreamingURL           bool
	_appstreamCreateApplication                           bool
	_appstreamCreateDirectoryConfig                       bool
	_appstreamCreateEntitlement                           bool
	_appstreamCreateExportImageTask                       bool
	_appstreamCreateFleet                                 bool
	_appstreamCreateImageBuilder                          bool
	_appstreamCreateImageBuilderStreamingURL              bool
	_appstreamCreateImportedImage                         bool
	_appstreamCreateStack                                 bool
	_appstreamCreateStreamingURL                          bool
	_appstreamCreateThemeForStack                         bool
	_appstreamCreateUpdatedImage                          bool
	_appstreamCreateUsageReportSubscription               bool
	_appstreamCreateUser                                  bool
	_appstreamDeleteAppBlock                              bool
	_appstreamDeleteAppBlockBuilder                       bool
	_appstreamDeleteApplication                           bool
	_appstreamDeleteDirectoryConfig                       bool
	_appstreamDeleteEntitlement                           bool
	_appstreamDeleteFleet                                 bool
	_appstreamDeleteImage                                 bool
	_appstreamDeleteImageBuilder                          bool
	_appstreamDeleteImagePermissions                      bool
	_appstreamDeleteStack                                 bool
	_appstreamDeleteThemeForStack                         bool
	_appstreamDeleteUsageReportSubscription               bool
	_appstreamDeleteUser                                  bool
	_appstreamDescribeAppBlockBuilderAppBlockAssociations bool
	_appstreamDescribeAppBlockBuilders                    bool
	_appstreamDescribeAppBlocks                           bool
	_appstreamDescribeAppLicenseUsage                     bool
	_appstreamDescribeApplicationFleetAssociations        bool
	_appstreamDescribeApplications                        bool
	_appstreamDescribeDirectoryConfigs                    bool
	_appstreamDescribeEntitlements                        bool
	_appstreamDescribeFleets                              bool
	_appstreamDescribeImageBuilders                       bool
	_appstreamDescribeImagePermissions                    bool
	_appstreamDescribeImages                              bool
	_appstreamDescribeSessions                            bool
	_appstreamDescribeSoftwareAssociations                bool
	_appstreamDescribeStacks                              bool
	_appstreamDescribeThemeForStack                       bool
	_appstreamDescribeUsageReportSubscriptions            bool
	_appstreamDescribeUserStackAssociations               bool
	_appstreamDescribeUsers                               bool
	_appstreamDisableUser                                 bool
	_appstreamDisassociateAppBlockBuilderAppBlock         bool
	_appstreamDisassociateApplicationFleet                bool
	_appstreamDisassociateApplicationFromEntitlement      bool
	_appstreamDisassociateFleet                           bool
	_appstreamDisassociateSoftwareFromImageBuilder        bool
	_appstreamEnableUser                                  bool
	_appstreamExpireSession                               bool
	_appstreamGetExportImageTask                          bool
	_appstreamListAssociatedFleets                        bool
	_appstreamListAssociatedStacks                        bool
	_appstreamListEntitledApplications                    bool
	_appstreamListExportImageTasks                        bool
	_appstreamListTagsForResource                         bool
	_appstreamStartAppBlockBuilder                        bool
	_appstreamStartFleet                                  bool
	_appstreamStartImageBuilder                           bool
	_appstreamStartSoftwareDeploymentToImageBuilder       bool
	_appstreamStopAppBlockBuilder                         bool
	_appstreamStopFleet                                   bool
	_appstreamStopImageBuilder                            bool
	_appstreamTagResource                                 bool
	_appstreamUntagResource                               bool
	_appstreamUpdateAppBlockBuilder                       bool
	_appstreamUpdateApplication                           bool
	_appstreamUpdateDirectoryConfig                       bool
	_appstreamUpdateEntitlement                           bool
	_appstreamUpdateFleet                                 bool
	_appstreamUpdateImagePermissions                      bool
	_appstreamUpdateStack                                 bool
	_appstreamUpdateThemeForStack                         bool

	_appstreamAccessEndpoints                      string
	_appstreamAgentSoftwareVersion                 string
	_appstreamAmiDescription                       string
	_appstreamAmiName                              string
	_appstreamAppBlockArn                          string
	_appstreamAppBlockBuilderName                  string
	_appstreamAppCatalogConfig                     string
	_appstreamAppVisibility                        string
	_appstreamApplicationArn                       string
	_appstreamApplicationId                        string
	_appstreamApplicationIdentifier                string
	_appstreamApplicationSettings                  string
	_appstreamAppstreamAgentVersion                string
	_appstreamArns                                 []string
	_appstreamAssociatedResource                   string
	_appstreamAttributes                           string
	_appstreamAttributesToDelete                   string
	_appstreamAuthenticationType                   string
	_appstreamBillingPeriod                        string
	_appstreamCertificateBasedAuthProperties       string
	_appstreamComputeCapacity                      string
	_appstreamDeleteStorageConnectors              string
	_appstreamDeleteVpcConfig                      string
	_appstreamDescription                          string
	_appstreamDestinationImageDescription          string
	_appstreamDestinationImageName                 string
	_appstreamDestinationRegion                    string
	_appstreamDirectoryName                        string
	_appstreamDirectoryNames                       []string
	_appstreamDisableIMDSV1                        string
	_appstreamDisconnectTimeoutInSeconds           string
	_appstreamDisplayName                          string
	_appstreamDomainJoinInfo                       string
	_appstreamDryRun                               string
	_appstreamEmbedHostDomains                     []string
	_appstreamEnableDefaultInternetAccess          string
	_appstreamEntitlementName                      string
	_appstreamExistingImageName                    string
	_appstreamFaviconS3Location                    string
	_appstreamFeedbackURL                          string
	_appstreamFilters                              string
	_appstreamFirstName                            string
	_appstreamFleetName                            string
	_appstreamFleetType                            string
	_appstreamFooterLinks                          string
	_appstreamIamRoleArn                           string
	_appstreamIconS3Location                       string
	_appstreamIdleDisconnectTimeoutInSeconds       string
	_appstreamImageArn                             string
	_appstreamImageBuilderName                     string
	_appstreamImageName                            string
	_appstreamImagePermissions                     string
	_appstreamInstanceFamilies                     []string
	_appstreamInstanceId                           string
	_appstreamInstanceType                         string
	_appstreamLastName                             string
	_appstreamLaunchParameters                     string
	_appstreamLaunchPath                           string
	_appstreamLimit                                string
	_appstreamMaxConcurrentSessions                string
	_appstreamMaxResults                           string
	_appstreamMaxSessionsPerInstance               string
	_appstreamMaxUserDurationInSeconds             string
	_appstreamMessageAction                        string
	_appstreamName                                 string
	_appstreamNames                                []string
	_appstreamNewImageDescription                  string
	_appstreamNewImageDisplayName                  string
	_appstreamNewImageName                         string
	_appstreamNewImageTags                         string
	_appstreamNextToken                            string
	_appstreamOrganizationLogoS3Location           string
	_appstreamOrganizationalUnitDistinguishedNames []string
	_appstreamPackagingType                        string
	_appstreamPlatform                             string
	_appstreamPlatforms                            string
	_appstreamPostSetupScriptDetails               string
	_appstreamRedirectURL                          string
	_appstreamResourceArn                          string
	_appstreamRetryFailedDeployments               string
	_appstreamRootVolumeConfig                     string
	_appstreamRuntimeValidationConfig              string
	_appstreamServiceAccountCredentials            string
	_appstreamSessionContext                       string
	_appstreamSessionId                            string
	_appstreamSessionScriptS3Location              string
	_appstreamSetupScriptDetails                   string
	_appstreamSharedAccountId                      string
	_appstreamSharedAwsAccountIds                  []string
	_appstreamSoftwareNames                        []string
	_appstreamSoftwaresToInstall                   []string
	_appstreamSoftwaresToUninstall                 []string
	_appstreamSourceAmiId                          string
	_appstreamSourceImageName                      string
	_appstreamSourceS3Location                     string
	_appstreamStackName                            string
	_appstreamState                                string
	_appstreamStorageConnectors                    string
	_appstreamStreamView                           string
	_appstreamStreamingExperienceSettings          string
	_appstreamTagKeys                              []string
	_appstreamTagSpecifications                    string
	_appstreamTags                                 string
	_appstreamTaskId                               string
	_appstreamThemeStyling                         string
	_appstreamTitleText                            string
	_appstreamType                                 string
	_appstreamUsbDeviceFilterStrings               []string
	_appstreamUserId                               string
	_appstreamUserName                             string
	_appstreamUserSettings                         string
	_appstreamUserStackAssociations                string
	_appstreamValidity                             string
	_appstreamVpcConfig                            string
	_appstreamWorkingDirectory                     string
)

// Associates the specified app block builder with the specified app block.
func appstream_AssociateAppBlockBuilderAppBlock(cfg aws.Config, client *appstream.Client) {
	input := &appstream.AssociateAppBlockBuilderAppBlockInput{
		// AppBlockArn: *string, // Required
		// AppBlockBuilderName: *string, // Required
	}

	if len(_appstreamAppBlockArn) > 0 {
		input.AppBlockArn = aws.String(_appstreamAppBlockArn)
	}
	if len(_appstreamAppBlockBuilderName) > 0 {
		input.AppBlockBuilderName = aws.String(_appstreamAppBlockBuilderName)
	}

	if resp, err := client.AssociateAppBlockBuilderAppBlock(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates the specified application with the specified fleet. This is only
// supported for Elastic fleets.
func appstream_AssociateApplicationFleet(cfg aws.Config, client *appstream.Client) {
	input := &appstream.AssociateApplicationFleetInput{
		// ApplicationArn: *string, // Required
		// FleetName: *string, // Required
	}

	if len(_appstreamApplicationArn) > 0 {
		input.ApplicationArn = aws.String(_appstreamApplicationArn)
	}
	if len(_appstreamFleetName) > 0 {
		input.FleetName = aws.String(_appstreamFleetName)
	}

	if resp, err := client.AssociateApplicationFleet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates an application to entitle.
func appstream_AssociateApplicationToEntitlement(cfg aws.Config, client *appstream.Client) {
	input := &appstream.AssociateApplicationToEntitlementInput{
		// ApplicationIdentifier: *string, // Required
		// EntitlementName: *string, // Required
		// StackName: *string, // Required
	}

	if len(_appstreamApplicationIdentifier) > 0 {
		input.ApplicationIdentifier = aws.String(_appstreamApplicationIdentifier)
	}
	if len(_appstreamEntitlementName) > 0 {
		input.EntitlementName = aws.String(_appstreamEntitlementName)
	}
	if len(_appstreamStackName) > 0 {
		input.StackName = aws.String(_appstreamStackName)
	}

	if resp, err := client.AssociateApplicationToEntitlement(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates the specified fleet with the specified stack.
func appstream_AssociateFleet(cfg aws.Config, client *appstream.Client) {
	input := &appstream.AssociateFleetInput{
		// FleetName: *string, // Required
		// StackName: *string, // Required
	}

	if len(_appstreamFleetName) > 0 {
		input.FleetName = aws.String(_appstreamFleetName)
	}
	if len(_appstreamStackName) > 0 {
		input.StackName = aws.String(_appstreamStackName)
	}

	if resp, err := client.AssociateFleet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates license included application(s) with an existing image builder
// instance.
func appstream_AssociateSoftwareToImageBuilder(cfg aws.Config, client *appstream.Client) {
	input := &appstream.AssociateSoftwareToImageBuilderInput{
		// ImageBuilderName: *string, // Required
		// SoftwareNames: []string, // Required
	}

	if len(_appstreamImageBuilderName) > 0 {
		input.ImageBuilderName = aws.String(_appstreamImageBuilderName)
	}
	if len(_appstreamSoftwareNames) > 0 {
		input.SoftwareNames = append([]string(nil), _appstreamSoftwareNames...)
	}

	if resp, err := client.AssociateSoftwareToImageBuilder(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates the specified users with the specified stacks. Users in a user pool
// cannot be assigned to stacks with fleets that are joined to an Active Directory
// domain.
func appstream_BatchAssociateUserStack(cfg aws.Config, client *appstream.Client) {
	input := &appstream.BatchAssociateUserStackInput{
		// UserStackAssociations: []types.UserStackAssociation, // Required
	}

	if len(_appstreamUserStackAssociations) > 0 {
		if err := assignInputField(input, "UserStackAssociations", _appstreamUserStackAssociations); err != nil {
			log.Errorf("invalid --user-stack-associations: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchAssociateUserStack(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates the specified users from the specified stacks.
func appstream_BatchDisassociateUserStack(cfg aws.Config, client *appstream.Client) {
	input := &appstream.BatchDisassociateUserStackInput{
		// UserStackAssociations: []types.UserStackAssociation, // Required
	}

	if len(_appstreamUserStackAssociations) > 0 {
		if err := assignInputField(input, "UserStackAssociations", _appstreamUserStackAssociations); err != nil {
			log.Errorf("invalid --user-stack-associations: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchDisassociateUserStack(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Copies the image within the same region or to a new region within the same AWS
// account. Note that any tags you added to the image will not be copied.
func appstream_CopyImage(cfg aws.Config, client *appstream.Client) {
	input := &appstream.CopyImageInput{
		// DestinationImageName: *string, // Required
		// DestinationRegion: *string, // Required
		// SourceImageName: *string, // Required
	}

	if len(_appstreamDestinationImageName) > 0 {
		input.DestinationImageName = aws.String(_appstreamDestinationImageName)
	}
	if len(_appstreamDestinationRegion) > 0 {
		input.DestinationRegion = aws.String(_appstreamDestinationRegion)
	}
	if len(_appstreamSourceImageName) > 0 {
		input.SourceImageName = aws.String(_appstreamSourceImageName)
	}
	if len(_appstreamDestinationImageDescription) > 0 {
		input.DestinationImageDescription = aws.String(_appstreamDestinationImageDescription)
	}

	if resp, err := client.CopyImage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an app block.
// App blocks are a WorkSpaces Applications resource that stores the details about
// the virtual hard disk in an S3 bucket. It also stores the setup script with
// details about how to mount the virtual hard disk. The virtual hard disk includes
// the application binaries and other files necessary to launch your applications.
// Multiple applications can be assigned to a single app block.
//
// This is only supported for Elastic fleets.
func appstream_CreateAppBlock(cfg aws.Config, client *appstream.Client) {
	input := &appstream.CreateAppBlockInput{
		// Name: *string, // Required
		// SourceS3Location: *types.S3Location, // Required
	}

	if len(_appstreamName) > 0 {
		input.Name = aws.String(_appstreamName)
	}
	if len(_appstreamSourceS3Location) > 0 {
		if err := assignInputField(input, "SourceS3Location", _appstreamSourceS3Location); err != nil {
			log.Errorf("invalid --source-s3-location: %s", err.Error())
			return
		}
	}
	if len(_appstreamDescription) > 0 {
		input.Description = aws.String(_appstreamDescription)
	}
	if len(_appstreamDisplayName) > 0 {
		input.DisplayName = aws.String(_appstreamDisplayName)
	}
	if len(_appstreamPackagingType) > 0 {
		if err := assignInputField(input, "PackagingType", _appstreamPackagingType); err != nil {
			log.Errorf("invalid --packaging-type: %s", err.Error())
			return
		}
	}
	if len(_appstreamPostSetupScriptDetails) > 0 {
		if err := assignInputField(input, "PostSetupScriptDetails", _appstreamPostSetupScriptDetails); err != nil {
			log.Errorf("invalid --post-setup-script-details: %s", err.Error())
			return
		}
	}
	if len(_appstreamSetupScriptDetails) > 0 {
		if err := assignInputField(input, "SetupScriptDetails", _appstreamSetupScriptDetails); err != nil {
			log.Errorf("invalid --setup-script-details: %s", err.Error())
			return
		}
	}
	if len(_appstreamTags) > 0 {
		if err := assignInputField(input, "Tags", _appstreamTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAppBlock(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an app block builder.
func appstream_CreateAppBlockBuilder(cfg aws.Config, client *appstream.Client) {
	input := &appstream.CreateAppBlockBuilderInput{
		// InstanceType: *string, // Required
		// Name: *string, // Required
		// Platform: types.AppBlockBuilderPlatformType, // Required
		// VpcConfig: *types.VpcConfig, // Required
	}

	if len(_appstreamInstanceType) > 0 {
		input.InstanceType = aws.String(_appstreamInstanceType)
	}
	if len(_appstreamName) > 0 {
		input.Name = aws.String(_appstreamName)
	}
	if len(_appstreamPlatform) > 0 {
		if err := assignInputField(input, "Platform", _appstreamPlatform); err != nil {
			log.Errorf("invalid --platform: %s", err.Error())
			return
		}
	}
	if len(_appstreamVpcConfig) > 0 {
		if err := assignInputField(input, "VpcConfig", _appstreamVpcConfig); err != nil {
			log.Errorf("invalid --vpc-config: %s", err.Error())
			return
		}
	}
	if len(_appstreamAccessEndpoints) > 0 {
		if err := assignInputField(input, "AccessEndpoints", _appstreamAccessEndpoints); err != nil {
			log.Errorf("invalid --access-endpoints: %s", err.Error())
			return
		}
	}
	if len(_appstreamDescription) > 0 {
		input.Description = aws.String(_appstreamDescription)
	}
	if len(_appstreamDisableIMDSV1) > 0 {
		if err := assignInputField(input, "DisableIMDSV1", _appstreamDisableIMDSV1); err != nil {
			log.Errorf("invalid --disable-imdsv1: %s", err.Error())
			return
		}
	}
	if len(_appstreamDisplayName) > 0 {
		input.DisplayName = aws.String(_appstreamDisplayName)
	}
	if len(_appstreamEnableDefaultInternetAccess) > 0 {
		if err := assignInputField(input, "EnableDefaultInternetAccess", _appstreamEnableDefaultInternetAccess); err != nil {
			log.Errorf("invalid --enable-default-internet-access: %s", err.Error())
			return
		}
	}
	if len(_appstreamIamRoleArn) > 0 {
		input.IamRoleArn = aws.String(_appstreamIamRoleArn)
	}
	if len(_appstreamTags) > 0 {
		if err := assignInputField(input, "Tags", _appstreamTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAppBlockBuilder(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a URL to start a create app block builder streaming session.
func appstream_CreateAppBlockBuilderStreamingURL(cfg aws.Config, client *appstream.Client) {
	input := &appstream.CreateAppBlockBuilderStreamingURLInput{
		// AppBlockBuilderName: *string, // Required
	}

	if len(_appstreamAppBlockBuilderName) > 0 {
		input.AppBlockBuilderName = aws.String(_appstreamAppBlockBuilderName)
	}
	if len(_appstreamValidity) > 0 {
		if err := assignInputField(input, "Validity", _appstreamValidity); err != nil {
			log.Errorf("invalid --validity: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAppBlockBuilderStreamingURL(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an application.
// Applications are a WorkSpaces Applications resource that stores the details
// about how to launch applications on Elastic fleet streaming instances. An
// application consists of the launch details, icon, and display name. Applications
// are associated with an app block that contains the application binaries and
// other files. The applications assigned to an Elastic fleet are the applications
// users can launch.
//
// This is only supported for Elastic fleets.
func appstream_CreateApplication(cfg aws.Config, client *appstream.Client) {
	input := &appstream.CreateApplicationInput{
		// AppBlockArn: *string, // Required
		// IconS3Location: *types.S3Location, // Required
		// InstanceFamilies: []string, // Required
		// LaunchPath: *string, // Required
		// Name: *string, // Required
		// Platforms: []types.PlatformType, // Required
	}

	if len(_appstreamAppBlockArn) > 0 {
		input.AppBlockArn = aws.String(_appstreamAppBlockArn)
	}
	if len(_appstreamIconS3Location) > 0 {
		if err := assignInputField(input, "IconS3Location", _appstreamIconS3Location); err != nil {
			log.Errorf("invalid --icon-s3-location: %s", err.Error())
			return
		}
	}
	if len(_appstreamInstanceFamilies) > 0 {
		input.InstanceFamilies = append([]string(nil), _appstreamInstanceFamilies...)
	}
	if len(_appstreamLaunchPath) > 0 {
		input.LaunchPath = aws.String(_appstreamLaunchPath)
	}
	if len(_appstreamName) > 0 {
		input.Name = aws.String(_appstreamName)
	}
	if len(_appstreamPlatforms) > 0 {
		if err := assignInputField(input, "Platforms", _appstreamPlatforms); err != nil {
			log.Errorf("invalid --platforms: %s", err.Error())
			return
		}
	}
	if len(_appstreamDescription) > 0 {
		input.Description = aws.String(_appstreamDescription)
	}
	if len(_appstreamDisplayName) > 0 {
		input.DisplayName = aws.String(_appstreamDisplayName)
	}
	if len(_appstreamLaunchParameters) > 0 {
		input.LaunchParameters = aws.String(_appstreamLaunchParameters)
	}
	if len(_appstreamTags) > 0 {
		if err := assignInputField(input, "Tags", _appstreamTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_appstreamWorkingDirectory) > 0 {
		input.WorkingDirectory = aws.String(_appstreamWorkingDirectory)
	}

	if resp, err := client.CreateApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a Directory Config object in WorkSpaces Applications. This object
// includes the configuration information required to join fleets and image
// builders to Microsoft Active Directory domains.
func appstream_CreateDirectoryConfig(cfg aws.Config, client *appstream.Client) {
	input := &appstream.CreateDirectoryConfigInput{
		// DirectoryName: *string, // Required
		// OrganizationalUnitDistinguishedNames: []string, // Required
	}

	if len(_appstreamDirectoryName) > 0 {
		input.DirectoryName = aws.String(_appstreamDirectoryName)
	}
	if len(_appstreamOrganizationalUnitDistinguishedNames) > 0 {
		input.OrganizationalUnitDistinguishedNames = append([]string(nil), _appstreamOrganizationalUnitDistinguishedNames...)
	}
	if len(_appstreamCertificateBasedAuthProperties) > 0 {
		if err := assignInputField(input, "CertificateBasedAuthProperties", _appstreamCertificateBasedAuthProperties); err != nil {
			log.Errorf("invalid --certificate-based-auth-properties: %s", err.Error())
			return
		}
	}
	if len(_appstreamServiceAccountCredentials) > 0 {
		if err := assignInputField(input, "ServiceAccountCredentials", _appstreamServiceAccountCredentials); err != nil {
			log.Errorf("invalid --service-account-credentials: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDirectoryConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new entitlement. Entitlements control access to specific applications
// within a stack, based on user attributes. Entitlements apply to SAML 2.0
// federated user identities. WorkSpaces Applications user pool and streaming URL
// users are entitled to all applications in a stack. Entitlements don't apply to
// the desktop stream view application, or to applications managed by a dynamic app
// provider using the Dynamic Application Framework.
func appstream_CreateEntitlement(cfg aws.Config, client *appstream.Client) {
	input := &appstream.CreateEntitlementInput{
		// AppVisibility: types.AppVisibility, // Required
		// Attributes: []types.EntitlementAttribute, // Required
		// Name: *string, // Required
		// StackName: *string, // Required
	}

	if len(_appstreamAppVisibility) > 0 {
		if err := assignInputField(input, "AppVisibility", _appstreamAppVisibility); err != nil {
			log.Errorf("invalid --app-visibility: %s", err.Error())
			return
		}
	}
	if len(_appstreamAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _appstreamAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_appstreamName) > 0 {
		input.Name = aws.String(_appstreamName)
	}
	if len(_appstreamStackName) > 0 {
		input.StackName = aws.String(_appstreamStackName)
	}
	if len(_appstreamDescription) > 0 {
		input.Description = aws.String(_appstreamDescription)
	}

	if resp, err := client.CreateEntitlement(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a task to export a WorkSpaces Applications image to an EC2 AMI. This
// allows you to use your customized WorkSpaces Applications images with other AWS
// services or for backup purposes.
func appstream_CreateExportImageTask(cfg aws.Config, client *appstream.Client) {
	input := &appstream.CreateExportImageTaskInput{
		// AmiName: *string, // Required
		// IamRoleArn: *string, // Required
		// ImageName: *string, // Required
	}

	if len(_appstreamAmiName) > 0 {
		input.AmiName = aws.String(_appstreamAmiName)
	}
	if len(_appstreamIamRoleArn) > 0 {
		input.IamRoleArn = aws.String(_appstreamIamRoleArn)
	}
	if len(_appstreamImageName) > 0 {
		input.ImageName = aws.String(_appstreamImageName)
	}
	if len(_appstreamAmiDescription) > 0 {
		input.AmiDescription = aws.String(_appstreamAmiDescription)
	}
	if len(_appstreamTagSpecifications) > 0 {
		if err := assignInputField(input, "TagSpecifications", _appstreamTagSpecifications); err != nil {
			log.Errorf("invalid --tag-specifications: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateExportImageTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a fleet. A fleet consists of streaming instances that your users access
// for their applications and desktops.
func appstream_CreateFleet(cfg aws.Config, client *appstream.Client) {
	input := &appstream.CreateFleetInput{
		// InstanceType: *string, // Required
		// Name: *string, // Required
	}

	if len(_appstreamInstanceType) > 0 {
		input.InstanceType = aws.String(_appstreamInstanceType)
	}
	if len(_appstreamName) > 0 {
		input.Name = aws.String(_appstreamName)
	}
	if len(_appstreamComputeCapacity) > 0 {
		if err := assignInputField(input, "ComputeCapacity", _appstreamComputeCapacity); err != nil {
			log.Errorf("invalid --compute-capacity: %s", err.Error())
			return
		}
	}
	if len(_appstreamDescription) > 0 {
		input.Description = aws.String(_appstreamDescription)
	}
	if len(_appstreamDisableIMDSV1) > 0 {
		if err := assignInputField(input, "DisableIMDSV1", _appstreamDisableIMDSV1); err != nil {
			log.Errorf("invalid --disable-imdsv1: %s", err.Error())
			return
		}
	}
	if len(_appstreamDisconnectTimeoutInSeconds) > 0 {
		if err := assignInputField(input, "DisconnectTimeoutInSeconds", _appstreamDisconnectTimeoutInSeconds); err != nil {
			log.Errorf("invalid --disconnect-timeout-in-seconds: %s", err.Error())
			return
		}
	}
	if len(_appstreamDisplayName) > 0 {
		input.DisplayName = aws.String(_appstreamDisplayName)
	}
	if len(_appstreamDomainJoinInfo) > 0 {
		if err := assignInputField(input, "DomainJoinInfo", _appstreamDomainJoinInfo); err != nil {
			log.Errorf("invalid --domain-join-info: %s", err.Error())
			return
		}
	}
	if len(_appstreamEnableDefaultInternetAccess) > 0 {
		if err := assignInputField(input, "EnableDefaultInternetAccess", _appstreamEnableDefaultInternetAccess); err != nil {
			log.Errorf("invalid --enable-default-internet-access: %s", err.Error())
			return
		}
	}
	if len(_appstreamFleetType) > 0 {
		if err := assignInputField(input, "FleetType", _appstreamFleetType); err != nil {
			log.Errorf("invalid --fleet-type: %s", err.Error())
			return
		}
	}
	if len(_appstreamIamRoleArn) > 0 {
		input.IamRoleArn = aws.String(_appstreamIamRoleArn)
	}
	if len(_appstreamIdleDisconnectTimeoutInSeconds) > 0 {
		if err := assignInputField(input, "IdleDisconnectTimeoutInSeconds", _appstreamIdleDisconnectTimeoutInSeconds); err != nil {
			log.Errorf("invalid --idle-disconnect-timeout-in-seconds: %s", err.Error())
			return
		}
	}
	if len(_appstreamImageArn) > 0 {
		input.ImageArn = aws.String(_appstreamImageArn)
	}
	if len(_appstreamImageName) > 0 {
		input.ImageName = aws.String(_appstreamImageName)
	}
	if len(_appstreamMaxConcurrentSessions) > 0 {
		if err := assignInputField(input, "MaxConcurrentSessions", _appstreamMaxConcurrentSessions); err != nil {
			log.Errorf("invalid --max-concurrent-sessions: %s", err.Error())
			return
		}
	}
	if len(_appstreamMaxSessionsPerInstance) > 0 {
		if err := assignInputField(input, "MaxSessionsPerInstance", _appstreamMaxSessionsPerInstance); err != nil {
			log.Errorf("invalid --max-sessions-per-instance: %s", err.Error())
			return
		}
	}
	if len(_appstreamMaxUserDurationInSeconds) > 0 {
		if err := assignInputField(input, "MaxUserDurationInSeconds", _appstreamMaxUserDurationInSeconds); err != nil {
			log.Errorf("invalid --max-user-duration-in-seconds: %s", err.Error())
			return
		}
	}
	if len(_appstreamPlatform) > 0 {
		if err := assignInputField(input, "Platform", _appstreamPlatform); err != nil {
			log.Errorf("invalid --platform: %s", err.Error())
			return
		}
	}
	if len(_appstreamRootVolumeConfig) > 0 {
		if err := assignInputField(input, "RootVolumeConfig", _appstreamRootVolumeConfig); err != nil {
			log.Errorf("invalid --root-volume-config: %s", err.Error())
			return
		}
	}
	if len(_appstreamSessionScriptS3Location) > 0 {
		if err := assignInputField(input, "SessionScriptS3Location", _appstreamSessionScriptS3Location); err != nil {
			log.Errorf("invalid --session-script-s3-location: %s", err.Error())
			return
		}
	}
	if len(_appstreamStreamView) > 0 {
		if err := assignInputField(input, "StreamView", _appstreamStreamView); err != nil {
			log.Errorf("invalid --stream-view: %s", err.Error())
			return
		}
	}
	if len(_appstreamTags) > 0 {
		if err := assignInputField(input, "Tags", _appstreamTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_appstreamUsbDeviceFilterStrings) > 0 {
		input.UsbDeviceFilterStrings = append([]string(nil), _appstreamUsbDeviceFilterStrings...)
	}
	if len(_appstreamVpcConfig) > 0 {
		if err := assignInputField(input, "VpcConfig", _appstreamVpcConfig); err != nil {
			log.Errorf("invalid --vpc-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateFleet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an image builder. An image builder is a virtual machine that is used to
// create an image.
//
// The initial state of the builder is PENDING . When it is ready, the state is
// RUNNING .
func appstream_CreateImageBuilder(cfg aws.Config, client *appstream.Client) {
	input := &appstream.CreateImageBuilderInput{
		// InstanceType: *string, // Required
		// Name: *string, // Required
	}

	if len(_appstreamInstanceType) > 0 {
		input.InstanceType = aws.String(_appstreamInstanceType)
	}
	if len(_appstreamName) > 0 {
		input.Name = aws.String(_appstreamName)
	}
	if len(_appstreamAccessEndpoints) > 0 {
		if err := assignInputField(input, "AccessEndpoints", _appstreamAccessEndpoints); err != nil {
			log.Errorf("invalid --access-endpoints: %s", err.Error())
			return
		}
	}
	if len(_appstreamAppstreamAgentVersion) > 0 {
		input.AppstreamAgentVersion = aws.String(_appstreamAppstreamAgentVersion)
	}
	if len(_appstreamDescription) > 0 {
		input.Description = aws.String(_appstreamDescription)
	}
	if len(_appstreamDisableIMDSV1) > 0 {
		if err := assignInputField(input, "DisableIMDSV1", _appstreamDisableIMDSV1); err != nil {
			log.Errorf("invalid --disable-imdsv1: %s", err.Error())
			return
		}
	}
	if len(_appstreamDisplayName) > 0 {
		input.DisplayName = aws.String(_appstreamDisplayName)
	}
	if len(_appstreamDomainJoinInfo) > 0 {
		if err := assignInputField(input, "DomainJoinInfo", _appstreamDomainJoinInfo); err != nil {
			log.Errorf("invalid --domain-join-info: %s", err.Error())
			return
		}
	}
	if len(_appstreamEnableDefaultInternetAccess) > 0 {
		if err := assignInputField(input, "EnableDefaultInternetAccess", _appstreamEnableDefaultInternetAccess); err != nil {
			log.Errorf("invalid --enable-default-internet-access: %s", err.Error())
			return
		}
	}
	if len(_appstreamIamRoleArn) > 0 {
		input.IamRoleArn = aws.String(_appstreamIamRoleArn)
	}
	if len(_appstreamImageArn) > 0 {
		input.ImageArn = aws.String(_appstreamImageArn)
	}
	if len(_appstreamImageName) > 0 {
		input.ImageName = aws.String(_appstreamImageName)
	}
	if len(_appstreamRootVolumeConfig) > 0 {
		if err := assignInputField(input, "RootVolumeConfig", _appstreamRootVolumeConfig); err != nil {
			log.Errorf("invalid --root-volume-config: %s", err.Error())
			return
		}
	}
	if len(_appstreamSoftwaresToInstall) > 0 {
		input.SoftwaresToInstall = append([]string(nil), _appstreamSoftwaresToInstall...)
	}
	if len(_appstreamSoftwaresToUninstall) > 0 {
		input.SoftwaresToUninstall = append([]string(nil), _appstreamSoftwaresToUninstall...)
	}
	if len(_appstreamTags) > 0 {
		if err := assignInputField(input, "Tags", _appstreamTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_appstreamVpcConfig) > 0 {
		if err := assignInputField(input, "VpcConfig", _appstreamVpcConfig); err != nil {
			log.Errorf("invalid --vpc-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateImageBuilder(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a URL to start an image builder streaming session.
func appstream_CreateImageBuilderStreamingURL(cfg aws.Config, client *appstream.Client) {
	input := &appstream.CreateImageBuilderStreamingURLInput{
		// Name: *string, // Required
	}

	if len(_appstreamName) > 0 {
		input.Name = aws.String(_appstreamName)
	}
	if len(_appstreamValidity) > 0 {
		if err := assignInputField(input, "Validity", _appstreamValidity); err != nil {
			log.Errorf("invalid --validity: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateImageBuilderStreamingURL(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a custom WorkSpaces Applications image by importing an EC2 AMI. This
// allows you to use your own customized AMI to create WorkSpaces Applications
// images that support additional instance types beyond the standard stream.*
// instances.
func appstream_CreateImportedImage(cfg aws.Config, client *appstream.Client) {
	input := &appstream.CreateImportedImageInput{
		// IamRoleArn: *string, // Required
		// Name: *string, // Required
		// SourceAmiId: *string, // Required
	}

	if len(_appstreamIamRoleArn) > 0 {
		input.IamRoleArn = aws.String(_appstreamIamRoleArn)
	}
	if len(_appstreamName) > 0 {
		input.Name = aws.String(_appstreamName)
	}
	if len(_appstreamSourceAmiId) > 0 {
		input.SourceAmiId = aws.String(_appstreamSourceAmiId)
	}
	if len(_appstreamAgentSoftwareVersion) > 0 {
		if err := assignInputField(input, "AgentSoftwareVersion", _appstreamAgentSoftwareVersion); err != nil {
			log.Errorf("invalid --agent-software-version: %s", err.Error())
			return
		}
	}
	if len(_appstreamAppCatalogConfig) > 0 {
		if err := assignInputField(input, "AppCatalogConfig", _appstreamAppCatalogConfig); err != nil {
			log.Errorf("invalid --app-catalog-config: %s", err.Error())
			return
		}
	}
	if len(_appstreamDescription) > 0 {
		input.Description = aws.String(_appstreamDescription)
	}
	if len(_appstreamDisplayName) > 0 {
		input.DisplayName = aws.String(_appstreamDisplayName)
	}
	if len(_appstreamDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _appstreamDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}
	if len(_appstreamRuntimeValidationConfig) > 0 {
		if err := assignInputField(input, "RuntimeValidationConfig", _appstreamRuntimeValidationConfig); err != nil {
			log.Errorf("invalid --runtime-validation-config: %s", err.Error())
			return
		}
	}
	if len(_appstreamTags) > 0 {
		if err := assignInputField(input, "Tags", _appstreamTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateImportedImage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a stack to start streaming applications to users. A stack consists of
// an associated fleet, user access policies, and storage configurations.
func appstream_CreateStack(cfg aws.Config, client *appstream.Client) {
	input := &appstream.CreateStackInput{
		// Name: *string, // Required
	}

	if len(_appstreamName) > 0 {
		input.Name = aws.String(_appstreamName)
	}
	if len(_appstreamAccessEndpoints) > 0 {
		if err := assignInputField(input, "AccessEndpoints", _appstreamAccessEndpoints); err != nil {
			log.Errorf("invalid --access-endpoints: %s", err.Error())
			return
		}
	}
	if len(_appstreamApplicationSettings) > 0 {
		if err := assignInputField(input, "ApplicationSettings", _appstreamApplicationSettings); err != nil {
			log.Errorf("invalid --application-settings: %s", err.Error())
			return
		}
	}
	if len(_appstreamDescription) > 0 {
		input.Description = aws.String(_appstreamDescription)
	}
	if len(_appstreamDisplayName) > 0 {
		input.DisplayName = aws.String(_appstreamDisplayName)
	}
	if len(_appstreamEmbedHostDomains) > 0 {
		input.EmbedHostDomains = append([]string(nil), _appstreamEmbedHostDomains...)
	}
	if len(_appstreamFeedbackURL) > 0 {
		input.FeedbackURL = aws.String(_appstreamFeedbackURL)
	}
	if len(_appstreamRedirectURL) > 0 {
		input.RedirectURL = aws.String(_appstreamRedirectURL)
	}
	if len(_appstreamStorageConnectors) > 0 {
		if err := assignInputField(input, "StorageConnectors", _appstreamStorageConnectors); err != nil {
			log.Errorf("invalid --storage-connectors: %s", err.Error())
			return
		}
	}
	if len(_appstreamStreamingExperienceSettings) > 0 {
		if err := assignInputField(input, "StreamingExperienceSettings", _appstreamStreamingExperienceSettings); err != nil {
			log.Errorf("invalid --streaming-experience-settings: %s", err.Error())
			return
		}
	}
	if len(_appstreamTags) > 0 {
		if err := assignInputField(input, "Tags", _appstreamTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_appstreamUserSettings) > 0 {
		if err := assignInputField(input, "UserSettings", _appstreamUserSettings); err != nil {
			log.Errorf("invalid --user-settings: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateStack(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a temporary URL to start an WorkSpaces Applications streaming session
// for the specified user. A streaming URL enables application streaming to be
// tested without user setup.
func appstream_CreateStreamingURL(cfg aws.Config, client *appstream.Client) {
	input := &appstream.CreateStreamingURLInput{
		// FleetName: *string, // Required
		// StackName: *string, // Required
		// UserId: *string, // Required
	}

	if len(_appstreamFleetName) > 0 {
		input.FleetName = aws.String(_appstreamFleetName)
	}
	if len(_appstreamStackName) > 0 {
		input.StackName = aws.String(_appstreamStackName)
	}
	if len(_appstreamUserId) > 0 {
		input.UserId = aws.String(_appstreamUserId)
	}
	if len(_appstreamApplicationId) > 0 {
		input.ApplicationId = aws.String(_appstreamApplicationId)
	}
	if len(_appstreamSessionContext) > 0 {
		input.SessionContext = aws.String(_appstreamSessionContext)
	}
	if len(_appstreamValidity) > 0 {
		if err := assignInputField(input, "Validity", _appstreamValidity); err != nil {
			log.Errorf("invalid --validity: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateStreamingURL(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates custom branding that customizes the appearance of the streaming
// application catalog page.
func appstream_CreateThemeForStack(cfg aws.Config, client *appstream.Client) {
	input := &appstream.CreateThemeForStackInput{
		// FaviconS3Location: *types.S3Location, // Required
		// OrganizationLogoS3Location: *types.S3Location, // Required
		// StackName: *string, // Required
		// ThemeStyling: types.ThemeStyling, // Required
		// TitleText: *string, // Required
	}

	if len(_appstreamFaviconS3Location) > 0 {
		if err := assignInputField(input, "FaviconS3Location", _appstreamFaviconS3Location); err != nil {
			log.Errorf("invalid --favicon-s3-location: %s", err.Error())
			return
		}
	}
	if len(_appstreamOrganizationLogoS3Location) > 0 {
		if err := assignInputField(input, "OrganizationLogoS3Location", _appstreamOrganizationLogoS3Location); err != nil {
			log.Errorf("invalid --organization-logo-s3-location: %s", err.Error())
			return
		}
	}
	if len(_appstreamStackName) > 0 {
		input.StackName = aws.String(_appstreamStackName)
	}
	if len(_appstreamThemeStyling) > 0 {
		if err := assignInputField(input, "ThemeStyling", _appstreamThemeStyling); err != nil {
			log.Errorf("invalid --theme-styling: %s", err.Error())
			return
		}
	}
	if len(_appstreamTitleText) > 0 {
		input.TitleText = aws.String(_appstreamTitleText)
	}
	if len(_appstreamFooterLinks) > 0 {
		if err := assignInputField(input, "FooterLinks", _appstreamFooterLinks); err != nil {
			log.Errorf("invalid --footer-links: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateThemeForStack(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new image with the latest Windows operating system updates, driver
// updates, and WorkSpaces Applications agent software.
//
// For more information, see the "Update an Image by Using Managed WorkSpaces
// Applications Image Updates" section in [Administer Your WorkSpaces Applications Images], in the Amazon WorkSpaces Applications
// Administration Guide.
//
// [Administer Your WorkSpaces Applications Images]: https://docs.aws.amazon.com/appstream2/latest/developerguide/administer-images.html
func appstream_CreateUpdatedImage(cfg aws.Config, client *appstream.Client) {
	input := &appstream.CreateUpdatedImageInput{
		// ExistingImageName: *string, // Required
		// NewImageName: *string, // Required
	}

	if len(_appstreamExistingImageName) > 0 {
		input.ExistingImageName = aws.String(_appstreamExistingImageName)
	}
	if len(_appstreamNewImageName) > 0 {
		input.NewImageName = aws.String(_appstreamNewImageName)
	}
	if len(_appstreamDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _appstreamDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}
	if len(_appstreamNewImageDescription) > 0 {
		input.NewImageDescription = aws.String(_appstreamNewImageDescription)
	}
	if len(_appstreamNewImageDisplayName) > 0 {
		input.NewImageDisplayName = aws.String(_appstreamNewImageDisplayName)
	}
	if len(_appstreamNewImageTags) > 0 {
		if err := assignInputField(input, "NewImageTags", _appstreamNewImageTags); err != nil {
			log.Errorf("invalid --new-image-tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateUpdatedImage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a usage report subscription. Usage reports are generated daily.
func appstream_CreateUsageReportSubscription(cfg aws.Config, client *appstream.Client) {
	input := &appstream.CreateUsageReportSubscriptionInput{}

	if resp, err := client.CreateUsageReportSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new user in the user pool.
func appstream_CreateUser(cfg aws.Config, client *appstream.Client) {
	input := &appstream.CreateUserInput{
		// AuthenticationType: types.AuthenticationType, // Required
		// UserName: *string, // Required
	}

	if len(_appstreamAuthenticationType) > 0 {
		if err := assignInputField(input, "AuthenticationType", _appstreamAuthenticationType); err != nil {
			log.Errorf("invalid --authentication-type: %s", err.Error())
			return
		}
	}
	if len(_appstreamUserName) > 0 {
		input.UserName = aws.String(_appstreamUserName)
	}
	if len(_appstreamFirstName) > 0 {
		input.FirstName = aws.String(_appstreamFirstName)
	}
	if len(_appstreamLastName) > 0 {
		input.LastName = aws.String(_appstreamLastName)
	}
	if len(_appstreamMessageAction) > 0 {
		if err := assignInputField(input, "MessageAction", _appstreamMessageAction); err != nil {
			log.Errorf("invalid --message-action: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an app block.
func appstream_DeleteAppBlock(cfg aws.Config, client *appstream.Client) {
	input := &appstream.DeleteAppBlockInput{
		// Name: *string, // Required
	}

	if len(_appstreamName) > 0 {
		input.Name = aws.String(_appstreamName)
	}

	if resp, err := client.DeleteAppBlock(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an app block builder.
// An app block builder can only be deleted when it has no association with an app
// block.
func appstream_DeleteAppBlockBuilder(cfg aws.Config, client *appstream.Client) {
	input := &appstream.DeleteAppBlockBuilderInput{
		// Name: *string, // Required
	}

	if len(_appstreamName) > 0 {
		input.Name = aws.String(_appstreamName)
	}

	if resp, err := client.DeleteAppBlockBuilder(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an application.
func appstream_DeleteApplication(cfg aws.Config, client *appstream.Client) {
	input := &appstream.DeleteApplicationInput{
		// Name: *string, // Required
	}

	if len(_appstreamName) > 0 {
		input.Name = aws.String(_appstreamName)
	}

	if resp, err := client.DeleteApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified Directory Config object from WorkSpaces Applications.
// This object includes the information required to join streaming instances to an
// Active Directory domain.
func appstream_DeleteDirectoryConfig(cfg aws.Config, client *appstream.Client) {
	input := &appstream.DeleteDirectoryConfigInput{
		// DirectoryName: *string, // Required
	}

	if len(_appstreamDirectoryName) > 0 {
		input.DirectoryName = aws.String(_appstreamDirectoryName)
	}

	if resp, err := client.DeleteDirectoryConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified entitlement.
func appstream_DeleteEntitlement(cfg aws.Config, client *appstream.Client) {
	input := &appstream.DeleteEntitlementInput{
		// Name: *string, // Required
		// StackName: *string, // Required
	}

	if len(_appstreamName) > 0 {
		input.Name = aws.String(_appstreamName)
	}
	if len(_appstreamStackName) > 0 {
		input.StackName = aws.String(_appstreamStackName)
	}

	if resp, err := client.DeleteEntitlement(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified fleet.
func appstream_DeleteFleet(cfg aws.Config, client *appstream.Client) {
	input := &appstream.DeleteFleetInput{
		// Name: *string, // Required
	}

	if len(_appstreamName) > 0 {
		input.Name = aws.String(_appstreamName)
	}

	if resp, err := client.DeleteFleet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified image. You cannot delete an image when it is in use.
// After you delete an image, you cannot provision new capacity using the image.
func appstream_DeleteImage(cfg aws.Config, client *appstream.Client) {
	input := &appstream.DeleteImageInput{
		// Name: *string, // Required
	}

	if len(_appstreamName) > 0 {
		input.Name = aws.String(_appstreamName)
	}

	if resp, err := client.DeleteImage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified image builder and releases the capacity.
func appstream_DeleteImageBuilder(cfg aws.Config, client *appstream.Client) {
	input := &appstream.DeleteImageBuilderInput{
		// Name: *string, // Required
	}

	if len(_appstreamName) > 0 {
		input.Name = aws.String(_appstreamName)
	}

	if resp, err := client.DeleteImageBuilder(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes permissions for the specified private image. After you delete
// permissions for an image, AWS accounts to which you previously granted these
// permissions can no longer use the image.
func appstream_DeleteImagePermissions(cfg aws.Config, client *appstream.Client) {
	input := &appstream.DeleteImagePermissionsInput{
		// Name: *string, // Required
		// SharedAccountId: *string, // Required
	}

	if len(_appstreamName) > 0 {
		input.Name = aws.String(_appstreamName)
	}
	if len(_appstreamSharedAccountId) > 0 {
		input.SharedAccountId = aws.String(_appstreamSharedAccountId)
	}

	if resp, err := client.DeleteImagePermissions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified stack. After the stack is deleted, the application
// streaming environment provided by the stack is no longer available to users.
// Also, any reservations made for application streaming sessions for the stack are
// released.
func appstream_DeleteStack(cfg aws.Config, client *appstream.Client) {
	input := &appstream.DeleteStackInput{
		// Name: *string, // Required
	}

	if len(_appstreamName) > 0 {
		input.Name = aws.String(_appstreamName)
	}

	if resp, err := client.DeleteStack(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes custom branding that customizes the appearance of the streaming
// application catalog page.
func appstream_DeleteThemeForStack(cfg aws.Config, client *appstream.Client) {
	input := &appstream.DeleteThemeForStackInput{
		// StackName: *string, // Required
	}

	if len(_appstreamStackName) > 0 {
		input.StackName = aws.String(_appstreamStackName)
	}

	if resp, err := client.DeleteThemeForStack(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables usage report generation.
func appstream_DeleteUsageReportSubscription(cfg aws.Config, client *appstream.Client) {
	input := &appstream.DeleteUsageReportSubscriptionInput{}

	if resp, err := client.DeleteUsageReportSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a user from the user pool.
func appstream_DeleteUser(cfg aws.Config, client *appstream.Client) {
	input := &appstream.DeleteUserInput{
		// AuthenticationType: types.AuthenticationType, // Required
		// UserName: *string, // Required
	}

	if len(_appstreamAuthenticationType) > 0 {
		if err := assignInputField(input, "AuthenticationType", _appstreamAuthenticationType); err != nil {
			log.Errorf("invalid --authentication-type: %s", err.Error())
			return
		}
	}
	if len(_appstreamUserName) > 0 {
		input.UserName = aws.String(_appstreamUserName)
	}

	if resp, err := client.DeleteUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list that describes one or more app block builder associations.
func appstream_DescribeAppBlockBuilderAppBlockAssociations(cfg aws.Config, client *appstream.Client) {
	input := &appstream.DescribeAppBlockBuilderAppBlockAssociationsInput{}

	if len(_appstreamAppBlockArn) > 0 {
		input.AppBlockArn = aws.String(_appstreamAppBlockArn)
	}
	if len(_appstreamAppBlockBuilderName) > 0 {
		input.AppBlockBuilderName = aws.String(_appstreamAppBlockBuilderName)
	}
	if len(_appstreamMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appstreamMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appstreamNextToken) > 0 {
		input.NextToken = aws.String(_appstreamNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeAppBlockBuilderAppBlockAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*appstream.DescribeAppBlockBuilderAppBlockAssociationsOutput
	p := appstream.NewDescribeAppBlockBuilderAppBlockAssociationsPaginator(client, input)
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

// Retrieves a list that describes one or more app block builders.
func appstream_DescribeAppBlockBuilders(cfg aws.Config, client *appstream.Client) {
	input := &appstream.DescribeAppBlockBuildersInput{}

	if len(_appstreamMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appstreamMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appstreamNames) > 0 {
		input.Names = append([]string(nil), _appstreamNames...)
	}
	if len(_appstreamNextToken) > 0 {
		input.NextToken = aws.String(_appstreamNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeAppBlockBuilders(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*appstream.DescribeAppBlockBuildersOutput
	p := appstream.NewDescribeAppBlockBuildersPaginator(client, input)
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

// Retrieves a list that describes one or more app blocks.
func appstream_DescribeAppBlocks(cfg aws.Config, client *appstream.Client) {
	input := &appstream.DescribeAppBlocksInput{}

	if len(_appstreamArns) > 0 {
		input.Arns = append([]string(nil), _appstreamArns...)
	}
	if len(_appstreamMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appstreamMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appstreamNextToken) > 0 {
		input.NextToken = aws.String(_appstreamNextToken)
	}

	if resp, err := client.DescribeAppBlocks(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves license included application usage information.
func appstream_DescribeAppLicenseUsage(cfg aws.Config, client *appstream.Client) {
	input := &appstream.DescribeAppLicenseUsageInput{
		// BillingPeriod: *string, // Required
	}

	if len(_appstreamBillingPeriod) > 0 {
		input.BillingPeriod = aws.String(_appstreamBillingPeriod)
	}
	if len(_appstreamMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appstreamMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appstreamNextToken) > 0 {
		input.NextToken = aws.String(_appstreamNextToken)
	}

	if resp, err := client.DescribeAppLicenseUsage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list that describes one or more application fleet associations.
// Either ApplicationArn or FleetName must be specified.
func appstream_DescribeApplicationFleetAssociations(cfg aws.Config, client *appstream.Client) {
	input := &appstream.DescribeApplicationFleetAssociationsInput{}

	if len(_appstreamApplicationArn) > 0 {
		input.ApplicationArn = aws.String(_appstreamApplicationArn)
	}
	if len(_appstreamFleetName) > 0 {
		input.FleetName = aws.String(_appstreamFleetName)
	}
	if len(_appstreamMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appstreamMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appstreamNextToken) > 0 {
		input.NextToken = aws.String(_appstreamNextToken)
	}

	if resp, err := client.DescribeApplicationFleetAssociations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list that describes one or more applications.
func appstream_DescribeApplications(cfg aws.Config, client *appstream.Client) {
	input := &appstream.DescribeApplicationsInput{}

	if len(_appstreamArns) > 0 {
		input.Arns = append([]string(nil), _appstreamArns...)
	}
	if len(_appstreamMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appstreamMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appstreamNextToken) > 0 {
		input.NextToken = aws.String(_appstreamNextToken)
	}

	if resp, err := client.DescribeApplications(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list that describes one or more specified Directory Config objects
// for WorkSpaces Applications, if the names for these objects are provided.
// Otherwise, all Directory Config objects in the account are described. These
// objects include the configuration information required to join fleets and image
// builders to Microsoft Active Directory domains.
//
// Although the response syntax in this topic includes the account password, this
// password is not returned in the actual response.
func appstream_DescribeDirectoryConfigs(cfg aws.Config, client *appstream.Client) {
	input := &appstream.DescribeDirectoryConfigsInput{}

	if len(_appstreamDirectoryNames) > 0 {
		input.DirectoryNames = append([]string(nil), _appstreamDirectoryNames...)
	}
	if len(_appstreamMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appstreamMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appstreamNextToken) > 0 {
		input.NextToken = aws.String(_appstreamNextToken)
	}

	if resp, err := client.DescribeDirectoryConfigs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list that describes one of more entitlements.
func appstream_DescribeEntitlements(cfg aws.Config, client *appstream.Client) {
	input := &appstream.DescribeEntitlementsInput{
		// StackName: *string, // Required
	}

	if len(_appstreamStackName) > 0 {
		input.StackName = aws.String(_appstreamStackName)
	}
	if len(_appstreamMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appstreamMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appstreamName) > 0 {
		input.Name = aws.String(_appstreamName)
	}
	if len(_appstreamNextToken) > 0 {
		input.NextToken = aws.String(_appstreamNextToken)
	}

	if resp, err := client.DescribeEntitlements(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list that describes one or more specified fleets, if the fleet
// names are provided. Otherwise, all fleets in the account are described.
func appstream_DescribeFleets(cfg aws.Config, client *appstream.Client) {
	input := &appstream.DescribeFleetsInput{}

	if len(_appstreamNames) > 0 {
		input.Names = append([]string(nil), _appstreamNames...)
	}
	if len(_appstreamNextToken) > 0 {
		input.NextToken = aws.String(_appstreamNextToken)
	}

	if resp, err := client.DescribeFleets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list that describes one or more specified image builders, if the
// image builder names are provided. Otherwise, all image builders in the account
// are described.
func appstream_DescribeImageBuilders(cfg aws.Config, client *appstream.Client) {
	input := &appstream.DescribeImageBuildersInput{}

	if len(_appstreamMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appstreamMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appstreamNames) > 0 {
		input.Names = append([]string(nil), _appstreamNames...)
	}
	if len(_appstreamNextToken) > 0 {
		input.NextToken = aws.String(_appstreamNextToken)
	}

	if resp, err := client.DescribeImageBuilders(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list that describes the permissions for shared AWS account IDs on a
// private image that you own.
func appstream_DescribeImagePermissions(cfg aws.Config, client *appstream.Client) {
	input := &appstream.DescribeImagePermissionsInput{
		// Name: *string, // Required
	}

	if len(_appstreamName) > 0 {
		input.Name = aws.String(_appstreamName)
	}
	if len(_appstreamMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appstreamMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appstreamNextToken) > 0 {
		input.NextToken = aws.String(_appstreamNextToken)
	}
	if len(_appstreamSharedAwsAccountIds) > 0 {
		input.SharedAwsAccountIds = append([]string(nil), _appstreamSharedAwsAccountIds...)
	}

	if disablePaginator() {
		if resp, err := client.DescribeImagePermissions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*appstream.DescribeImagePermissionsOutput
	p := appstream.NewDescribeImagePermissionsPaginator(client, input)
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

// Retrieves a list that describes one or more specified images, if the image
// names or image ARNs are provided. Otherwise, all images in the account are
// described.
func appstream_DescribeImages(cfg aws.Config, client *appstream.Client) {
	input := &appstream.DescribeImagesInput{}

	if len(_appstreamArns) > 0 {
		input.Arns = append([]string(nil), _appstreamArns...)
	}
	if len(_appstreamMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appstreamMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appstreamNames) > 0 {
		input.Names = append([]string(nil), _appstreamNames...)
	}
	if len(_appstreamNextToken) > 0 {
		input.NextToken = aws.String(_appstreamNextToken)
	}
	if len(_appstreamType) > 0 {
		if err := assignInputField(input, "Type", _appstreamType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeImages(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*appstream.DescribeImagesOutput
	p := appstream.NewDescribeImagesPaginator(client, input)
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

// Retrieves a list that describes the streaming sessions for a specified stack
// and fleet. If a UserId is provided for the stack and fleet, only streaming
// sessions for that user are described. If an authentication type is not provided,
// the default is to authenticate users using a streaming URL.
func appstream_DescribeSessions(cfg aws.Config, client *appstream.Client) {
	input := &appstream.DescribeSessionsInput{
		// FleetName: *string, // Required
		// StackName: *string, // Required
	}

	if len(_appstreamFleetName) > 0 {
		input.FleetName = aws.String(_appstreamFleetName)
	}
	if len(_appstreamStackName) > 0 {
		input.StackName = aws.String(_appstreamStackName)
	}
	if len(_appstreamAuthenticationType) > 0 {
		if err := assignInputField(input, "AuthenticationType", _appstreamAuthenticationType); err != nil {
			log.Errorf("invalid --authentication-type: %s", err.Error())
			return
		}
	}
	if len(_appstreamInstanceId) > 0 {
		input.InstanceId = aws.String(_appstreamInstanceId)
	}
	if len(_appstreamLimit) > 0 {
		if err := assignInputField(input, "Limit", _appstreamLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_appstreamNextToken) > 0 {
		input.NextToken = aws.String(_appstreamNextToken)
	}
	if len(_appstreamUserId) > 0 {
		input.UserId = aws.String(_appstreamUserId)
	}

	if resp, err := client.DescribeSessions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves license included application associations for a specified resource.
func appstream_DescribeSoftwareAssociations(cfg aws.Config, client *appstream.Client) {
	input := &appstream.DescribeSoftwareAssociationsInput{
		// AssociatedResource: *string, // Required
	}

	if len(_appstreamAssociatedResource) > 0 {
		input.AssociatedResource = aws.String(_appstreamAssociatedResource)
	}
	if len(_appstreamMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appstreamMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appstreamNextToken) > 0 {
		input.NextToken = aws.String(_appstreamNextToken)
	}

	if resp, err := client.DescribeSoftwareAssociations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list that describes one or more specified stacks, if the stack
// names are provided. Otherwise, all stacks in the account are described.
func appstream_DescribeStacks(cfg aws.Config, client *appstream.Client) {
	input := &appstream.DescribeStacksInput{}

	if len(_appstreamNames) > 0 {
		input.Names = append([]string(nil), _appstreamNames...)
	}
	if len(_appstreamNextToken) > 0 {
		input.NextToken = aws.String(_appstreamNextToken)
	}

	if resp, err := client.DescribeStacks(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list that describes the theme for a specified stack. A theme is
// custom branding that customizes the appearance of the streaming application
// catalog page.
func appstream_DescribeThemeForStack(cfg aws.Config, client *appstream.Client) {
	input := &appstream.DescribeThemeForStackInput{
		// StackName: *string, // Required
	}

	if len(_appstreamStackName) > 0 {
		input.StackName = aws.String(_appstreamStackName)
	}

	if resp, err := client.DescribeThemeForStack(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list that describes one or more usage report subscriptions.
func appstream_DescribeUsageReportSubscriptions(cfg aws.Config, client *appstream.Client) {
	input := &appstream.DescribeUsageReportSubscriptionsInput{}

	if len(_appstreamMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appstreamMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appstreamNextToken) > 0 {
		input.NextToken = aws.String(_appstreamNextToken)
	}

	if resp, err := client.DescribeUsageReportSubscriptions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list that describes the UserStackAssociation objects. You must
// specify either or both of the following:
//
// - The stack name
//
// - The user name (email address of the user associated with the stack) and the
// authentication type for the user
func appstream_DescribeUserStackAssociations(cfg aws.Config, client *appstream.Client) {
	input := &appstream.DescribeUserStackAssociationsInput{}

	if len(_appstreamAuthenticationType) > 0 {
		if err := assignInputField(input, "AuthenticationType", _appstreamAuthenticationType); err != nil {
			log.Errorf("invalid --authentication-type: %s", err.Error())
			return
		}
	}
	if len(_appstreamMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appstreamMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appstreamNextToken) > 0 {
		input.NextToken = aws.String(_appstreamNextToken)
	}
	if len(_appstreamStackName) > 0 {
		input.StackName = aws.String(_appstreamStackName)
	}
	if len(_appstreamUserName) > 0 {
		input.UserName = aws.String(_appstreamUserName)
	}

	if resp, err := client.DescribeUserStackAssociations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list that describes one or more specified users in the user pool.
func appstream_DescribeUsers(cfg aws.Config, client *appstream.Client) {
	input := &appstream.DescribeUsersInput{
		// AuthenticationType: types.AuthenticationType, // Required
	}

	if len(_appstreamAuthenticationType) > 0 {
		if err := assignInputField(input, "AuthenticationType", _appstreamAuthenticationType); err != nil {
			log.Errorf("invalid --authentication-type: %s", err.Error())
			return
		}
	}
	if len(_appstreamMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appstreamMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appstreamNextToken) > 0 {
		input.NextToken = aws.String(_appstreamNextToken)
	}

	if resp, err := client.DescribeUsers(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables the specified user in the user pool. Users can't sign in to WorkSpaces
// Applications until they are re-enabled. This action does not delete the user.
func appstream_DisableUser(cfg aws.Config, client *appstream.Client) {
	input := &appstream.DisableUserInput{
		// AuthenticationType: types.AuthenticationType, // Required
		// UserName: *string, // Required
	}

	if len(_appstreamAuthenticationType) > 0 {
		if err := assignInputField(input, "AuthenticationType", _appstreamAuthenticationType); err != nil {
			log.Errorf("invalid --authentication-type: %s", err.Error())
			return
		}
	}
	if len(_appstreamUserName) > 0 {
		input.UserName = aws.String(_appstreamUserName)
	}

	if resp, err := client.DisableUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a specified app block builder from a specified app block.
func appstream_DisassociateAppBlockBuilderAppBlock(cfg aws.Config, client *appstream.Client) {
	input := &appstream.DisassociateAppBlockBuilderAppBlockInput{
		// AppBlockArn: *string, // Required
		// AppBlockBuilderName: *string, // Required
	}

	if len(_appstreamAppBlockArn) > 0 {
		input.AppBlockArn = aws.String(_appstreamAppBlockArn)
	}
	if len(_appstreamAppBlockBuilderName) > 0 {
		input.AppBlockBuilderName = aws.String(_appstreamAppBlockBuilderName)
	}

	if resp, err := client.DisassociateAppBlockBuilderAppBlock(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates the specified application from the fleet.
func appstream_DisassociateApplicationFleet(cfg aws.Config, client *appstream.Client) {
	input := &appstream.DisassociateApplicationFleetInput{
		// ApplicationArn: *string, // Required
		// FleetName: *string, // Required
	}

	if len(_appstreamApplicationArn) > 0 {
		input.ApplicationArn = aws.String(_appstreamApplicationArn)
	}
	if len(_appstreamFleetName) > 0 {
		input.FleetName = aws.String(_appstreamFleetName)
	}

	if resp, err := client.DisassociateApplicationFleet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified application from the specified entitlement.
func appstream_DisassociateApplicationFromEntitlement(cfg aws.Config, client *appstream.Client) {
	input := &appstream.DisassociateApplicationFromEntitlementInput{
		// ApplicationIdentifier: *string, // Required
		// EntitlementName: *string, // Required
		// StackName: *string, // Required
	}

	if len(_appstreamApplicationIdentifier) > 0 {
		input.ApplicationIdentifier = aws.String(_appstreamApplicationIdentifier)
	}
	if len(_appstreamEntitlementName) > 0 {
		input.EntitlementName = aws.String(_appstreamEntitlementName)
	}
	if len(_appstreamStackName) > 0 {
		input.StackName = aws.String(_appstreamStackName)
	}

	if resp, err := client.DisassociateApplicationFromEntitlement(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates the specified fleet from the specified stack.
func appstream_DisassociateFleet(cfg aws.Config, client *appstream.Client) {
	input := &appstream.DisassociateFleetInput{
		// FleetName: *string, // Required
		// StackName: *string, // Required
	}

	if len(_appstreamFleetName) > 0 {
		input.FleetName = aws.String(_appstreamFleetName)
	}
	if len(_appstreamStackName) > 0 {
		input.StackName = aws.String(_appstreamStackName)
	}

	if resp, err := client.DisassociateFleet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes license included application(s) association(s) from an image builder
// instance.
func appstream_DisassociateSoftwareFromImageBuilder(cfg aws.Config, client *appstream.Client) {
	input := &appstream.DisassociateSoftwareFromImageBuilderInput{
		// ImageBuilderName: *string, // Required
		// SoftwareNames: []string, // Required
	}

	if len(_appstreamImageBuilderName) > 0 {
		input.ImageBuilderName = aws.String(_appstreamImageBuilderName)
	}
	if len(_appstreamSoftwareNames) > 0 {
		input.SoftwareNames = append([]string(nil), _appstreamSoftwareNames...)
	}

	if resp, err := client.DisassociateSoftwareFromImageBuilder(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables a user in the user pool. After being enabled, users can sign in to
// WorkSpaces Applications and open applications from the stacks to which they are
// assigned.
func appstream_EnableUser(cfg aws.Config, client *appstream.Client) {
	input := &appstream.EnableUserInput{
		// AuthenticationType: types.AuthenticationType, // Required
		// UserName: *string, // Required
	}

	if len(_appstreamAuthenticationType) > 0 {
		if err := assignInputField(input, "AuthenticationType", _appstreamAuthenticationType); err != nil {
			log.Errorf("invalid --authentication-type: %s", err.Error())
			return
		}
	}
	if len(_appstreamUserName) > 0 {
		input.UserName = aws.String(_appstreamUserName)
	}

	if resp, err := client.EnableUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Immediately stops the specified streaming session.
func appstream_ExpireSession(cfg aws.Config, client *appstream.Client) {
	input := &appstream.ExpireSessionInput{
		// SessionId: *string, // Required
	}

	if len(_appstreamSessionId) > 0 {
		input.SessionId = aws.String(_appstreamSessionId)
	}

	if resp, err := client.ExpireSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about an export image task, including its current state,
// progress, and any error details.
func appstream_GetExportImageTask(cfg aws.Config, client *appstream.Client) {
	input := &appstream.GetExportImageTaskInput{}

	if len(_appstreamTaskId) > 0 {
		input.TaskId = aws.String(_appstreamTaskId)
	}

	if resp, err := client.GetExportImageTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the name of the fleet that is associated with the specified stack.
func appstream_ListAssociatedFleets(cfg aws.Config, client *appstream.Client) {
	input := &appstream.ListAssociatedFleetsInput{
		// StackName: *string, // Required
	}

	if len(_appstreamStackName) > 0 {
		input.StackName = aws.String(_appstreamStackName)
	}
	if len(_appstreamNextToken) > 0 {
		input.NextToken = aws.String(_appstreamNextToken)
	}

	if resp, err := client.ListAssociatedFleets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the name of the stack with which the specified fleet is associated.
func appstream_ListAssociatedStacks(cfg aws.Config, client *appstream.Client) {
	input := &appstream.ListAssociatedStacksInput{
		// FleetName: *string, // Required
	}

	if len(_appstreamFleetName) > 0 {
		input.FleetName = aws.String(_appstreamFleetName)
	}
	if len(_appstreamNextToken) > 0 {
		input.NextToken = aws.String(_appstreamNextToken)
	}

	if resp, err := client.ListAssociatedStacks(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of entitled applications.
func appstream_ListEntitledApplications(cfg aws.Config, client *appstream.Client) {
	input := &appstream.ListEntitledApplicationsInput{
		// EntitlementName: *string, // Required
		// StackName: *string, // Required
	}

	if len(_appstreamEntitlementName) > 0 {
		input.EntitlementName = aws.String(_appstreamEntitlementName)
	}
	if len(_appstreamStackName) > 0 {
		input.StackName = aws.String(_appstreamStackName)
	}
	if len(_appstreamMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appstreamMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appstreamNextToken) > 0 {
		input.NextToken = aws.String(_appstreamNextToken)
	}

	if resp, err := client.ListEntitledApplications(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists export image tasks, with optional filtering and pagination. Use this
// operation to monitor the status of multiple export operations.
func appstream_ListExportImageTasks(cfg aws.Config, client *appstream.Client) {
	input := &appstream.ListExportImageTasksInput{}

	if len(_appstreamFilters) > 0 {
		if err := assignInputField(input, "Filters", _appstreamFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_appstreamMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appstreamMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appstreamNextToken) > 0 {
		input.NextToken = aws.String(_appstreamNextToken)
	}

	if resp, err := client.ListExportImageTasks(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of all tags for the specified WorkSpaces Applications
// resource. You can tag WorkSpaces Applications image builders, images, fleets,
// and stacks.
//
// For more information about tags, see [Tagging Your Resources] in the Amazon WorkSpaces Applications
// Administration Guide.
//
// [Tagging Your Resources]: https://docs.aws.amazon.com/appstream2/latest/developerguide/tagging-basic.html
func appstream_ListTagsForResource(cfg aws.Config, client *appstream.Client) {
	input := &appstream.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_appstreamResourceArn) > 0 {
		input.ResourceArn = aws.String(_appstreamResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts an app block builder.
// An app block builder can only be started when it's associated with an app block.
//
// Starting an app block builder starts a new instance, which is equivalent to an
// elastic fleet instance with application builder assistance functionality.
func appstream_StartAppBlockBuilder(cfg aws.Config, client *appstream.Client) {
	input := &appstream.StartAppBlockBuilderInput{
		// Name: *string, // Required
	}

	if len(_appstreamName) > 0 {
		input.Name = aws.String(_appstreamName)
	}

	if resp, err := client.StartAppBlockBuilder(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts the specified fleet.
func appstream_StartFleet(cfg aws.Config, client *appstream.Client) {
	input := &appstream.StartFleetInput{
		// Name: *string, // Required
	}

	if len(_appstreamName) > 0 {
		input.Name = aws.String(_appstreamName)
	}

	if resp, err := client.StartFleet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts the specified image builder.
func appstream_StartImageBuilder(cfg aws.Config, client *appstream.Client) {
	input := &appstream.StartImageBuilderInput{
		// Name: *string, // Required
	}

	if len(_appstreamName) > 0 {
		input.Name = aws.String(_appstreamName)
	}
	if len(_appstreamAppstreamAgentVersion) > 0 {
		input.AppstreamAgentVersion = aws.String(_appstreamAppstreamAgentVersion)
	}

	if resp, err := client.StartImageBuilder(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initiates license included applications deployment to an image builder instance.
func appstream_StartSoftwareDeploymentToImageBuilder(cfg aws.Config, client *appstream.Client) {
	input := &appstream.StartSoftwareDeploymentToImageBuilderInput{
		// ImageBuilderName: *string, // Required
	}

	if len(_appstreamImageBuilderName) > 0 {
		input.ImageBuilderName = aws.String(_appstreamImageBuilderName)
	}
	if len(_appstreamRetryFailedDeployments) > 0 {
		if err := assignInputField(input, "RetryFailedDeployments", _appstreamRetryFailedDeployments); err != nil {
			log.Errorf("invalid --retry-failed-deployments: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartSoftwareDeploymentToImageBuilder(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops an app block builder.
// Stopping an app block builder terminates the instance, and the instance state
// is not persisted.
func appstream_StopAppBlockBuilder(cfg aws.Config, client *appstream.Client) {
	input := &appstream.StopAppBlockBuilderInput{
		// Name: *string, // Required
	}

	if len(_appstreamName) > 0 {
		input.Name = aws.String(_appstreamName)
	}

	if resp, err := client.StopAppBlockBuilder(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops the specified fleet.
func appstream_StopFleet(cfg aws.Config, client *appstream.Client) {
	input := &appstream.StopFleetInput{
		// Name: *string, // Required
	}

	if len(_appstreamName) > 0 {
		input.Name = aws.String(_appstreamName)
	}

	if resp, err := client.StopFleet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops the specified image builder.
func appstream_StopImageBuilder(cfg aws.Config, client *appstream.Client) {
	input := &appstream.StopImageBuilderInput{
		// Name: *string, // Required
	}

	if len(_appstreamName) > 0 {
		input.Name = aws.String(_appstreamName)
	}

	if resp, err := client.StopImageBuilder(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or overwrites one or more tags for the specified WorkSpaces Applications
// resource. You can tag WorkSpaces Applications image builders, images, fleets,
// and stacks.
//
// Each tag consists of a key and an optional value. If a resource already has a
// tag with the same key, this operation updates its value.
//
// To list the current tags for your resources, use ListTagsForResource. To disassociate tags from
// your resources, use UntagResource.
//
// For more information about tags, see [Tagging Your Resources] in the Amazon WorkSpaces Applications
// Administration Guide.
//
// [Tagging Your Resources]: https://docs.aws.amazon.com/appstream2/latest/developerguide/tagging-basic.html
func appstream_TagResource(cfg aws.Config, client *appstream.Client) {
	input := &appstream.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_appstreamResourceArn) > 0 {
		input.ResourceArn = aws.String(_appstreamResourceArn)
	}
	if len(_appstreamTags) > 0 {
		if err := assignInputField(input, "Tags", _appstreamTags); err != nil {
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

// Disassociates one or more specified tags from the specified WorkSpaces
// Applications resource.
//
// To list the current tags for your resources, use ListTagsForResource.
//
// For more information about tags, see [Tagging Your Resources] in the Amazon WorkSpaces Applications
// Administration Guide.
//
// [Tagging Your Resources]: https://docs.aws.amazon.com/appstream2/latest/developerguide/tagging-basic.html
func appstream_UntagResource(cfg aws.Config, client *appstream.Client) {
	input := &appstream.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_appstreamResourceArn) > 0 {
		input.ResourceArn = aws.String(_appstreamResourceArn)
	}
	if len(_appstreamTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _appstreamTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an app block builder.
// If the app block builder is in the STARTING or STOPPING state, you can't update
// it. If the app block builder is in the RUNNING state, you can only update the
// DisplayName and Description. If the app block builder is in the STOPPED state,
// you can update any attribute except the Name.
func appstream_UpdateAppBlockBuilder(cfg aws.Config, client *appstream.Client) {
	input := &appstream.UpdateAppBlockBuilderInput{
		// Name: *string, // Required
	}

	if len(_appstreamName) > 0 {
		input.Name = aws.String(_appstreamName)
	}
	if len(_appstreamAccessEndpoints) > 0 {
		if err := assignInputField(input, "AccessEndpoints", _appstreamAccessEndpoints); err != nil {
			log.Errorf("invalid --access-endpoints: %s", err.Error())
			return
		}
	}
	if len(_appstreamAttributesToDelete) > 0 {
		if err := assignInputField(input, "AttributesToDelete", _appstreamAttributesToDelete); err != nil {
			log.Errorf("invalid --attributes-to-delete: %s", err.Error())
			return
		}
	}
	if len(_appstreamDescription) > 0 {
		input.Description = aws.String(_appstreamDescription)
	}
	if len(_appstreamDisableIMDSV1) > 0 {
		if err := assignInputField(input, "DisableIMDSV1", _appstreamDisableIMDSV1); err != nil {
			log.Errorf("invalid --disable-imdsv1: %s", err.Error())
			return
		}
	}
	if len(_appstreamDisplayName) > 0 {
		input.DisplayName = aws.String(_appstreamDisplayName)
	}
	if len(_appstreamEnableDefaultInternetAccess) > 0 {
		if err := assignInputField(input, "EnableDefaultInternetAccess", _appstreamEnableDefaultInternetAccess); err != nil {
			log.Errorf("invalid --enable-default-internet-access: %s", err.Error())
			return
		}
	}
	if len(_appstreamIamRoleArn) > 0 {
		input.IamRoleArn = aws.String(_appstreamIamRoleArn)
	}
	if len(_appstreamInstanceType) > 0 {
		input.InstanceType = aws.String(_appstreamInstanceType)
	}
	if len(_appstreamPlatform) > 0 {
		if err := assignInputField(input, "Platform", _appstreamPlatform); err != nil {
			log.Errorf("invalid --platform: %s", err.Error())
			return
		}
	}
	if len(_appstreamVpcConfig) > 0 {
		if err := assignInputField(input, "VpcConfig", _appstreamVpcConfig); err != nil {
			log.Errorf("invalid --vpc-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAppBlockBuilder(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified application.
func appstream_UpdateApplication(cfg aws.Config, client *appstream.Client) {
	input := &appstream.UpdateApplicationInput{
		// Name: *string, // Required
	}

	if len(_appstreamName) > 0 {
		input.Name = aws.String(_appstreamName)
	}
	if len(_appstreamAppBlockArn) > 0 {
		input.AppBlockArn = aws.String(_appstreamAppBlockArn)
	}
	if len(_appstreamAttributesToDelete) > 0 {
		if err := assignInputField(input, "AttributesToDelete", _appstreamAttributesToDelete); err != nil {
			log.Errorf("invalid --attributes-to-delete: %s", err.Error())
			return
		}
	}
	if len(_appstreamDescription) > 0 {
		input.Description = aws.String(_appstreamDescription)
	}
	if len(_appstreamDisplayName) > 0 {
		input.DisplayName = aws.String(_appstreamDisplayName)
	}
	if len(_appstreamIconS3Location) > 0 {
		if err := assignInputField(input, "IconS3Location", _appstreamIconS3Location); err != nil {
			log.Errorf("invalid --icon-s3-location: %s", err.Error())
			return
		}
	}
	if len(_appstreamLaunchParameters) > 0 {
		input.LaunchParameters = aws.String(_appstreamLaunchParameters)
	}
	if len(_appstreamLaunchPath) > 0 {
		input.LaunchPath = aws.String(_appstreamLaunchPath)
	}
	if len(_appstreamWorkingDirectory) > 0 {
		input.WorkingDirectory = aws.String(_appstreamWorkingDirectory)
	}

	if resp, err := client.UpdateApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified Directory Config object in WorkSpaces Applications. This
// object includes the configuration information required to join fleets and image
// builders to Microsoft Active Directory domains.
func appstream_UpdateDirectoryConfig(cfg aws.Config, client *appstream.Client) {
	input := &appstream.UpdateDirectoryConfigInput{
		// DirectoryName: *string, // Required
	}

	if len(_appstreamDirectoryName) > 0 {
		input.DirectoryName = aws.String(_appstreamDirectoryName)
	}
	if len(_appstreamCertificateBasedAuthProperties) > 0 {
		if err := assignInputField(input, "CertificateBasedAuthProperties", _appstreamCertificateBasedAuthProperties); err != nil {
			log.Errorf("invalid --certificate-based-auth-properties: %s", err.Error())
			return
		}
	}
	if len(_appstreamOrganizationalUnitDistinguishedNames) > 0 {
		input.OrganizationalUnitDistinguishedNames = append([]string(nil), _appstreamOrganizationalUnitDistinguishedNames...)
	}
	if len(_appstreamServiceAccountCredentials) > 0 {
		if err := assignInputField(input, "ServiceAccountCredentials", _appstreamServiceAccountCredentials); err != nil {
			log.Errorf("invalid --service-account-credentials: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDirectoryConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified entitlement.
func appstream_UpdateEntitlement(cfg aws.Config, client *appstream.Client) {
	input := &appstream.UpdateEntitlementInput{
		// Name: *string, // Required
		// StackName: *string, // Required
	}

	if len(_appstreamName) > 0 {
		input.Name = aws.String(_appstreamName)
	}
	if len(_appstreamStackName) > 0 {
		input.StackName = aws.String(_appstreamStackName)
	}
	if len(_appstreamAppVisibility) > 0 {
		if err := assignInputField(input, "AppVisibility", _appstreamAppVisibility); err != nil {
			log.Errorf("invalid --app-visibility: %s", err.Error())
			return
		}
	}
	if len(_appstreamAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _appstreamAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_appstreamDescription) > 0 {
		input.Description = aws.String(_appstreamDescription)
	}

	if resp, err := client.UpdateEntitlement(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified fleet.
// If the fleet is in the STOPPED state, you can update any attribute except the
// fleet name.
//
// If the fleet is in the RUNNING state, you can update the following based on the
// fleet type:
//
// - Always-On and On-Demand fleet types
//
// You can update the DisplayName , ComputeCapacity , ImageARN , ImageName ,
//
// IdleDisconnectTimeoutInSeconds , and DisconnectTimeoutInSeconds attributes.
//
// - Elastic fleet type
//
// You can update the DisplayName , IdleDisconnectTimeoutInSeconds ,
//
// DisconnectTimeoutInSeconds , MaxConcurrentSessions , SessionScriptS3Location
// and UsbDeviceFilterStrings attributes.
//
// If the fleet is in the STARTING or STOPPED state, you can't update it.
func appstream_UpdateFleet(cfg aws.Config, client *appstream.Client) {
	input := &appstream.UpdateFleetInput{}

	if len(_appstreamAttributesToDelete) > 0 {
		if err := assignInputField(input, "AttributesToDelete", _appstreamAttributesToDelete); err != nil {
			log.Errorf("invalid --attributes-to-delete: %s", err.Error())
			return
		}
	}
	if len(_appstreamComputeCapacity) > 0 {
		if err := assignInputField(input, "ComputeCapacity", _appstreamComputeCapacity); err != nil {
			log.Errorf("invalid --compute-capacity: %s", err.Error())
			return
		}
	}
	if len(_appstreamDeleteVpcConfig) > 0 {
		if err := assignInputField(input, "DeleteVpcConfig", _appstreamDeleteVpcConfig); err != nil {
			log.Errorf("invalid --delete-vpc-config: %s", err.Error())
			return
		}
	}
	if len(_appstreamDescription) > 0 {
		input.Description = aws.String(_appstreamDescription)
	}
	if len(_appstreamDisableIMDSV1) > 0 {
		if err := assignInputField(input, "DisableIMDSV1", _appstreamDisableIMDSV1); err != nil {
			log.Errorf("invalid --disable-imdsv1: %s", err.Error())
			return
		}
	}
	if len(_appstreamDisconnectTimeoutInSeconds) > 0 {
		if err := assignInputField(input, "DisconnectTimeoutInSeconds", _appstreamDisconnectTimeoutInSeconds); err != nil {
			log.Errorf("invalid --disconnect-timeout-in-seconds: %s", err.Error())
			return
		}
	}
	if len(_appstreamDisplayName) > 0 {
		input.DisplayName = aws.String(_appstreamDisplayName)
	}
	if len(_appstreamDomainJoinInfo) > 0 {
		if err := assignInputField(input, "DomainJoinInfo", _appstreamDomainJoinInfo); err != nil {
			log.Errorf("invalid --domain-join-info: %s", err.Error())
			return
		}
	}
	if len(_appstreamEnableDefaultInternetAccess) > 0 {
		if err := assignInputField(input, "EnableDefaultInternetAccess", _appstreamEnableDefaultInternetAccess); err != nil {
			log.Errorf("invalid --enable-default-internet-access: %s", err.Error())
			return
		}
	}
	if len(_appstreamIamRoleArn) > 0 {
		input.IamRoleArn = aws.String(_appstreamIamRoleArn)
	}
	if len(_appstreamIdleDisconnectTimeoutInSeconds) > 0 {
		if err := assignInputField(input, "IdleDisconnectTimeoutInSeconds", _appstreamIdleDisconnectTimeoutInSeconds); err != nil {
			log.Errorf("invalid --idle-disconnect-timeout-in-seconds: %s", err.Error())
			return
		}
	}
	if len(_appstreamImageArn) > 0 {
		input.ImageArn = aws.String(_appstreamImageArn)
	}
	if len(_appstreamImageName) > 0 {
		input.ImageName = aws.String(_appstreamImageName)
	}
	if len(_appstreamInstanceType) > 0 {
		input.InstanceType = aws.String(_appstreamInstanceType)
	}
	if len(_appstreamMaxConcurrentSessions) > 0 {
		if err := assignInputField(input, "MaxConcurrentSessions", _appstreamMaxConcurrentSessions); err != nil {
			log.Errorf("invalid --max-concurrent-sessions: %s", err.Error())
			return
		}
	}
	if len(_appstreamMaxSessionsPerInstance) > 0 {
		if err := assignInputField(input, "MaxSessionsPerInstance", _appstreamMaxSessionsPerInstance); err != nil {
			log.Errorf("invalid --max-sessions-per-instance: %s", err.Error())
			return
		}
	}
	if len(_appstreamMaxUserDurationInSeconds) > 0 {
		if err := assignInputField(input, "MaxUserDurationInSeconds", _appstreamMaxUserDurationInSeconds); err != nil {
			log.Errorf("invalid --max-user-duration-in-seconds: %s", err.Error())
			return
		}
	}
	if len(_appstreamName) > 0 {
		input.Name = aws.String(_appstreamName)
	}
	if len(_appstreamPlatform) > 0 {
		if err := assignInputField(input, "Platform", _appstreamPlatform); err != nil {
			log.Errorf("invalid --platform: %s", err.Error())
			return
		}
	}
	if len(_appstreamRootVolumeConfig) > 0 {
		if err := assignInputField(input, "RootVolumeConfig", _appstreamRootVolumeConfig); err != nil {
			log.Errorf("invalid --root-volume-config: %s", err.Error())
			return
		}
	}
	if len(_appstreamSessionScriptS3Location) > 0 {
		if err := assignInputField(input, "SessionScriptS3Location", _appstreamSessionScriptS3Location); err != nil {
			log.Errorf("invalid --session-script-s3-location: %s", err.Error())
			return
		}
	}
	if len(_appstreamStreamView) > 0 {
		if err := assignInputField(input, "StreamView", _appstreamStreamView); err != nil {
			log.Errorf("invalid --stream-view: %s", err.Error())
			return
		}
	}
	if len(_appstreamUsbDeviceFilterStrings) > 0 {
		input.UsbDeviceFilterStrings = append([]string(nil), _appstreamUsbDeviceFilterStrings...)
	}
	if len(_appstreamVpcConfig) > 0 {
		if err := assignInputField(input, "VpcConfig", _appstreamVpcConfig); err != nil {
			log.Errorf("invalid --vpc-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateFleet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or updates permissions for the specified private image.
func appstream_UpdateImagePermissions(cfg aws.Config, client *appstream.Client) {
	input := &appstream.UpdateImagePermissionsInput{
		// ImagePermissions: *types.ImagePermissions, // Required
		// Name: *string, // Required
		// SharedAccountId: *string, // Required
	}

	if len(_appstreamImagePermissions) > 0 {
		if err := assignInputField(input, "ImagePermissions", _appstreamImagePermissions); err != nil {
			log.Errorf("invalid --image-permissions: %s", err.Error())
			return
		}
	}
	if len(_appstreamName) > 0 {
		input.Name = aws.String(_appstreamName)
	}
	if len(_appstreamSharedAccountId) > 0 {
		input.SharedAccountId = aws.String(_appstreamSharedAccountId)
	}

	if resp, err := client.UpdateImagePermissions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified fields for the specified stack.
func appstream_UpdateStack(cfg aws.Config, client *appstream.Client) {
	input := &appstream.UpdateStackInput{
		// Name: *string, // Required
	}

	if len(_appstreamName) > 0 {
		input.Name = aws.String(_appstreamName)
	}
	if len(_appstreamAccessEndpoints) > 0 {
		if err := assignInputField(input, "AccessEndpoints", _appstreamAccessEndpoints); err != nil {
			log.Errorf("invalid --access-endpoints: %s", err.Error())
			return
		}
	}
	if len(_appstreamApplicationSettings) > 0 {
		if err := assignInputField(input, "ApplicationSettings", _appstreamApplicationSettings); err != nil {
			log.Errorf("invalid --application-settings: %s", err.Error())
			return
		}
	}
	if len(_appstreamAttributesToDelete) > 0 {
		if err := assignInputField(input, "AttributesToDelete", _appstreamAttributesToDelete); err != nil {
			log.Errorf("invalid --attributes-to-delete: %s", err.Error())
			return
		}
	}
	if len(_appstreamDeleteStorageConnectors) > 0 {
		if err := assignInputField(input, "DeleteStorageConnectors", _appstreamDeleteStorageConnectors); err != nil {
			log.Errorf("invalid --delete-storage-connectors: %s", err.Error())
			return
		}
	}
	if len(_appstreamDescription) > 0 {
		input.Description = aws.String(_appstreamDescription)
	}
	if len(_appstreamDisplayName) > 0 {
		input.DisplayName = aws.String(_appstreamDisplayName)
	}
	if len(_appstreamEmbedHostDomains) > 0 {
		input.EmbedHostDomains = append([]string(nil), _appstreamEmbedHostDomains...)
	}
	if len(_appstreamFeedbackURL) > 0 {
		input.FeedbackURL = aws.String(_appstreamFeedbackURL)
	}
	if len(_appstreamRedirectURL) > 0 {
		input.RedirectURL = aws.String(_appstreamRedirectURL)
	}
	if len(_appstreamStorageConnectors) > 0 {
		if err := assignInputField(input, "StorageConnectors", _appstreamStorageConnectors); err != nil {
			log.Errorf("invalid --storage-connectors: %s", err.Error())
			return
		}
	}
	if len(_appstreamStreamingExperienceSettings) > 0 {
		if err := assignInputField(input, "StreamingExperienceSettings", _appstreamStreamingExperienceSettings); err != nil {
			log.Errorf("invalid --streaming-experience-settings: %s", err.Error())
			return
		}
	}
	if len(_appstreamUserSettings) > 0 {
		if err := assignInputField(input, "UserSettings", _appstreamUserSettings); err != nil {
			log.Errorf("invalid --user-settings: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateStack(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates custom branding that customizes the appearance of the streaming
// application catalog page.
func appstream_UpdateThemeForStack(cfg aws.Config, client *appstream.Client) {
	input := &appstream.UpdateThemeForStackInput{
		// StackName: *string, // Required
	}

	if len(_appstreamStackName) > 0 {
		input.StackName = aws.String(_appstreamStackName)
	}
	if len(_appstreamAttributesToDelete) > 0 {
		if err := assignInputField(input, "AttributesToDelete", _appstreamAttributesToDelete); err != nil {
			log.Errorf("invalid --attributes-to-delete: %s", err.Error())
			return
		}
	}
	if len(_appstreamFaviconS3Location) > 0 {
		if err := assignInputField(input, "FaviconS3Location", _appstreamFaviconS3Location); err != nil {
			log.Errorf("invalid --favicon-s3-location: %s", err.Error())
			return
		}
	}
	if len(_appstreamFooterLinks) > 0 {
		if err := assignInputField(input, "FooterLinks", _appstreamFooterLinks); err != nil {
			log.Errorf("invalid --footer-links: %s", err.Error())
			return
		}
	}
	if len(_appstreamOrganizationLogoS3Location) > 0 {
		if err := assignInputField(input, "OrganizationLogoS3Location", _appstreamOrganizationLogoS3Location); err != nil {
			log.Errorf("invalid --organization-logo-s3-location: %s", err.Error())
			return
		}
	}
	if len(_appstreamState) > 0 {
		if err := assignInputField(input, "State", _appstreamState); err != nil {
			log.Errorf("invalid --state: %s", err.Error())
			return
		}
	}
	if len(_appstreamThemeStyling) > 0 {
		if err := assignInputField(input, "ThemeStyling", _appstreamThemeStyling); err != nil {
			log.Errorf("invalid --theme-styling: %s", err.Error())
			return
		}
	}
	if len(_appstreamTitleText) > 0 {
		input.TitleText = aws.String(_appstreamTitleText)
	}

	if resp, err := client.UpdateThemeForStack(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_appstreamCmd)
	_appstreamCmd.Flags().SortFlags = false

	_appstreamCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_appstreamCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_appstreamCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_appstreamCmd.Flags().StringVarP(&_appstreamAccessEndpoints, "access-endpoints", "", "", "Access Endpoints")
	_appstreamCmd.Flags().StringVarP(&_appstreamAgentSoftwareVersion, "agent-software-version", "", "", "Agent Software Version")
	_appstreamCmd.Flags().StringVarP(&_appstreamAmiDescription, "ami-description", "", "", "AMI Description")
	_appstreamCmd.Flags().StringVarP(&_appstreamAmiName, "ami-name", "", "", "AMI Name")
	_appstreamCmd.Flags().StringVarP(&_appstreamAppBlockArn, "app-block-arn", "", "", "App Block ARN")
	_appstreamCmd.Flags().StringVarP(&_appstreamAppBlockBuilderName, "app-block-builder-name", "", "", "App Block Builder Name")
	_appstreamCmd.Flags().StringVarP(&_appstreamAppCatalogConfig, "app-catalog-config", "", "", "App Catalog Config")
	_appstreamCmd.Flags().StringVarP(&_appstreamAppVisibility, "app-visibility", "", "", "App Visibility")
	_appstreamCmd.Flags().StringVarP(&_appstreamApplicationArn, "application-arn", "", "", "Application ARN")
	_appstreamCmd.Flags().StringVarP(&_appstreamApplicationId, "application-id", "", "", "Application ID")
	_appstreamCmd.Flags().StringVarP(&_appstreamApplicationIdentifier, "application-identifier", "", "", "Application Identifier")
	_appstreamCmd.Flags().StringVarP(&_appstreamApplicationSettings, "application-settings", "", "", "Application Settings")
	_appstreamCmd.Flags().StringVarP(&_appstreamAppstreamAgentVersion, "appstream-agent-version", "", "", "Appstream Agent Version")
	_appstreamCmd.Flags().StringSliceVarP(&_appstreamArns, "arns", "", nil, "Arns")
	_appstreamCmd.Flags().StringVarP(&_appstreamAssociatedResource, "associated-resource", "", "", "Associated Resource")
	_appstreamCmd.Flags().StringVarP(&_appstreamAttributes, "attributes", "", "", "Attributes")
	_appstreamCmd.Flags().StringVarP(&_appstreamAttributesToDelete, "attributes-to-delete", "", "", "Attributes To Delete")
	_appstreamCmd.Flags().StringVarP(&_appstreamAuthenticationType, "authentication-type", "", "", "Authentication Type")
	_appstreamCmd.Flags().StringVarP(&_appstreamBillingPeriod, "billing-period", "", "", "Billing Period")
	_appstreamCmd.Flags().StringVarP(&_appstreamCertificateBasedAuthProperties, "certificate-based-auth-properties", "", "", "Certificate Based Auth Properties")
	_appstreamCmd.Flags().StringVarP(&_appstreamComputeCapacity, "compute-capacity", "", "", "Compute Capacity")
	_appstreamCmd.Flags().StringVarP(&_appstreamDeleteStorageConnectors, "delete-storage-connectors", "", "", "Delete Storage Connectors")
	_appstreamCmd.Flags().StringVarP(&_appstreamDeleteVpcConfig, "delete-vpc-config", "", "", "Delete VPC Config")
	_appstreamCmd.Flags().StringVarP(&_appstreamDescription, "description", "", "", "Description")
	_appstreamCmd.Flags().StringVarP(&_appstreamDestinationImageDescription, "destination-image-description", "", "", "Destination Image Description")
	_appstreamCmd.Flags().StringVarP(&_appstreamDestinationImageName, "destination-image-name", "", "", "Destination Image Name")
	_appstreamCmd.Flags().StringVarP(&_appstreamDestinationRegion, "destination-region", "", "", "Destination Region")
	_appstreamCmd.Flags().StringVarP(&_appstreamDirectoryName, "directory-name", "", "", "Directory Name")
	_appstreamCmd.Flags().StringSliceVarP(&_appstreamDirectoryNames, "directory-names", "", nil, "Directory Names")
	_appstreamCmd.Flags().StringVarP(&_appstreamDisableIMDSV1, "disable-imdsv1", "", "", "Disable Imdsv1")
	_appstreamCmd.Flags().StringVarP(&_appstreamDisconnectTimeoutInSeconds, "disconnect-timeout-in-seconds", "", "", "Disconnect Timeout In Seconds")
	_appstreamCmd.Flags().StringVarP(&_appstreamDisplayName, "display-name", "", "", "Display Name")
	_appstreamCmd.Flags().StringVarP(&_appstreamDomainJoinInfo, "domain-join-info", "", "", "Domain Join Info")
	_appstreamCmd.Flags().StringVarP(&_appstreamDryRun, "dry-run", "", "", "Dry Run")
	_appstreamCmd.Flags().StringSliceVarP(&_appstreamEmbedHostDomains, "embed-host-domains", "", nil, "Embed Host Domains")
	_appstreamCmd.Flags().StringVarP(&_appstreamEnableDefaultInternetAccess, "enable-default-internet-access", "", "", "Enable Default Internet Access")
	_appstreamCmd.Flags().StringVarP(&_appstreamEntitlementName, "entitlement-name", "", "", "Entitlement Name")
	_appstreamCmd.Flags().StringVarP(&_appstreamExistingImageName, "existing-image-name", "", "", "Existing Image Name")
	_appstreamCmd.Flags().StringVarP(&_appstreamFaviconS3Location, "favicon-s3-location", "", "", "Favicon S3 Location")
	_appstreamCmd.Flags().StringVarP(&_appstreamFeedbackURL, "feedback-url", "", "", "Feedback URL")
	_appstreamCmd.Flags().StringVarP(&_appstreamFilters, "filters", "", "", "Filters")
	_appstreamCmd.Flags().StringVarP(&_appstreamFirstName, "first-name", "", "", "First Name")
	_appstreamCmd.Flags().StringVarP(&_appstreamFleetName, "fleet-name", "", "", "Fleet Name")
	_appstreamCmd.Flags().StringVarP(&_appstreamFleetType, "fleet-type", "", "", "Fleet Type")
	_appstreamCmd.Flags().StringVarP(&_appstreamFooterLinks, "footer-links", "", "", "Footer Links")
	_appstreamCmd.Flags().StringVarP(&_appstreamIamRoleArn, "iam-role-arn", "", "", "IAM Role ARN")
	_appstreamCmd.Flags().StringVarP(&_appstreamIconS3Location, "icon-s3-location", "", "", "Icon S3 Location")
	_appstreamCmd.Flags().StringVarP(&_appstreamIdleDisconnectTimeoutInSeconds, "idle-disconnect-timeout-in-seconds", "", "", "Idle Disconnect Timeout In Seconds")
	_appstreamCmd.Flags().StringVarP(&_appstreamImageArn, "image-arn", "", "", "Image ARN")
	_appstreamCmd.Flags().StringVarP(&_appstreamImageBuilderName, "image-builder-name", "", "", "Image Builder Name")
	_appstreamCmd.Flags().StringVarP(&_appstreamImageName, "image-name", "", "", "Image Name")
	_appstreamCmd.Flags().StringVarP(&_appstreamImagePermissions, "image-permissions", "", "", "Image Permissions")
	_appstreamCmd.Flags().StringSliceVarP(&_appstreamInstanceFamilies, "instance-families", "", nil, "Instance Families")
	_appstreamCmd.Flags().StringVarP(&_appstreamInstanceId, "instance-id", "", "", "Instance ID")
	_appstreamCmd.Flags().StringVarP(&_appstreamInstanceType, "instance-type", "", "", "Instance Type")
	_appstreamCmd.Flags().StringVarP(&_appstreamLastName, "last-name", "", "", "Last Name")
	_appstreamCmd.Flags().StringVarP(&_appstreamLaunchParameters, "launch-parameters", "", "", "Launch Parameters")
	_appstreamCmd.Flags().StringVarP(&_appstreamLaunchPath, "launch-path", "", "", "Launch Path")
	_appstreamCmd.Flags().StringVarP(&_appstreamLimit, "limit", "", "", "Limit")
	_appstreamCmd.Flags().StringVarP(&_appstreamMaxConcurrentSessions, "max-concurrent-sessions", "", "", "Max Concurrent Sessions")
	_appstreamCmd.Flags().StringVarP(&_appstreamMaxResults, "max-results", "", "", "Max Results")
	_appstreamCmd.Flags().StringVarP(&_appstreamMaxSessionsPerInstance, "max-sessions-per-instance", "", "", "Max Sessions Per Instance")
	_appstreamCmd.Flags().StringVarP(&_appstreamMaxUserDurationInSeconds, "max-user-duration-in-seconds", "", "", "Max User Duration In Seconds")
	_appstreamCmd.Flags().StringVarP(&_appstreamMessageAction, "message-action", "", "", "Message Action")
	_appstreamCmd.Flags().StringVarP(&_appstreamName, "name", "", "", "Name")
	_appstreamCmd.Flags().StringSliceVarP(&_appstreamNames, "names", "", nil, "Names")
	_appstreamCmd.Flags().StringVarP(&_appstreamNewImageDescription, "new-image-description", "", "", "New Image Description")
	_appstreamCmd.Flags().StringVarP(&_appstreamNewImageDisplayName, "new-image-display-name", "", "", "New Image Display Name")
	_appstreamCmd.Flags().StringVarP(&_appstreamNewImageName, "new-image-name", "", "", "New Image Name")
	_appstreamCmd.Flags().StringVarP(&_appstreamNewImageTags, "new-image-tags", "", "", "New Image Tags")
	_appstreamCmd.Flags().StringVarP(&_appstreamNextToken, "next-token", "", "", "Next Token")
	_appstreamCmd.Flags().StringVarP(&_appstreamOrganizationLogoS3Location, "organization-logo-s3-location", "", "", "Organization Logo S3 Location")
	_appstreamCmd.Flags().StringSliceVarP(&_appstreamOrganizationalUnitDistinguishedNames, "organizational-unit-distinguished-names", "", nil, "Organizational Unit Distinguished Names")
	_appstreamCmd.Flags().StringVarP(&_appstreamPackagingType, "packaging-type", "", "", "Packaging Type")
	_appstreamCmd.Flags().StringVarP(&_appstreamPlatform, "platform", "", "", "Platform")
	_appstreamCmd.Flags().StringVarP(&_appstreamPlatforms, "platforms", "", "", "Platforms")
	_appstreamCmd.Flags().StringVarP(&_appstreamPostSetupScriptDetails, "post-setup-script-details", "", "", "Post Setup Script Details")
	_appstreamCmd.Flags().StringVarP(&_appstreamRedirectURL, "redirect-url", "", "", "Redirect URL")
	_appstreamCmd.Flags().StringVarP(&_appstreamResourceArn, "resource-arn", "", "", "Resource ARN")
	_appstreamCmd.Flags().StringVarP(&_appstreamRetryFailedDeployments, "retry-failed-deployments", "", "", "Retry Failed Deployments")
	_appstreamCmd.Flags().StringVarP(&_appstreamRootVolumeConfig, "root-volume-config", "", "", "Root Volume Config")
	_appstreamCmd.Flags().StringVarP(&_appstreamRuntimeValidationConfig, "runtime-validation-config", "", "", "Runtime Validation Config")
	_appstreamCmd.Flags().StringVarP(&_appstreamServiceAccountCredentials, "service-account-credentials", "", "", "Service Account Credentials")
	_appstreamCmd.Flags().StringVarP(&_appstreamSessionContext, "session-context", "", "", "Session Context")
	_appstreamCmd.Flags().StringVarP(&_appstreamSessionId, "session-id", "", "", "Session ID")
	_appstreamCmd.Flags().StringVarP(&_appstreamSessionScriptS3Location, "session-script-s3-location", "", "", "Session Script S3 Location")
	_appstreamCmd.Flags().StringVarP(&_appstreamSetupScriptDetails, "setup-script-details", "", "", "Setup Script Details")
	_appstreamCmd.Flags().StringVarP(&_appstreamSharedAccountId, "shared-account-id", "", "", "Shared Account ID")
	_appstreamCmd.Flags().StringSliceVarP(&_appstreamSharedAwsAccountIds, "shared-aws-account-ids", "", nil, "Shared AWS Account Ids")
	_appstreamCmd.Flags().StringSliceVarP(&_appstreamSoftwareNames, "software-names", "", nil, "Software Names")
	_appstreamCmd.Flags().StringSliceVarP(&_appstreamSoftwaresToInstall, "softwares-to-install", "", nil, "Softwares To Install")
	_appstreamCmd.Flags().StringSliceVarP(&_appstreamSoftwaresToUninstall, "softwares-to-uninstall", "", nil, "Softwares To Uninstall")
	_appstreamCmd.Flags().StringVarP(&_appstreamSourceAmiId, "source-ami-id", "", "", "Source AMI ID")
	_appstreamCmd.Flags().StringVarP(&_appstreamSourceImageName, "source-image-name", "", "", "Source Image Name")
	_appstreamCmd.Flags().StringVarP(&_appstreamSourceS3Location, "source-s3-location", "", "", "Source S3 Location")
	_appstreamCmd.Flags().StringVarP(&_appstreamStackName, "stack-name", "", "", "Stack Name")
	_appstreamCmd.Flags().StringVarP(&_appstreamState, "state", "", "", "State")
	_appstreamCmd.Flags().StringVarP(&_appstreamStorageConnectors, "storage-connectors", "", "", "Storage Connectors")
	_appstreamCmd.Flags().StringVarP(&_appstreamStreamView, "stream-view", "", "", "Stream View")
	_appstreamCmd.Flags().StringVarP(&_appstreamStreamingExperienceSettings, "streaming-experience-settings", "", "", "Streaming Experience Settings")
	_appstreamCmd.Flags().StringSliceVarP(&_appstreamTagKeys, "tag-keys", "", nil, "Tag Keys")
	_appstreamCmd.Flags().StringVarP(&_appstreamTagSpecifications, "tag-specifications", "", "", "Tag Specifications")
	_appstreamCmd.Flags().StringVarP(&_appstreamTags, "tags", "", "", "Tags")
	_appstreamCmd.Flags().StringVarP(&_appstreamTaskId, "task-id", "", "", "Task ID")
	_appstreamCmd.Flags().StringVarP(&_appstreamThemeStyling, "theme-styling", "", "", "Theme Styling")
	_appstreamCmd.Flags().StringVarP(&_appstreamTitleText, "title-text", "", "", "Title Text")
	_appstreamCmd.Flags().StringVarP(&_appstreamType, "type", "", "", "Type")
	_appstreamCmd.Flags().StringSliceVarP(&_appstreamUsbDeviceFilterStrings, "usb-device-filter-strings", "", nil, "Usb Device Filter Strings")
	_appstreamCmd.Flags().StringVarP(&_appstreamUserId, "user-id", "", "", "User ID")
	_appstreamCmd.Flags().StringVarP(&_appstreamUserName, "user-name", "", "", "User Name")
	_appstreamCmd.Flags().StringVarP(&_appstreamUserSettings, "user-settings", "", "", "User Settings")
	_appstreamCmd.Flags().StringVarP(&_appstreamUserStackAssociations, "user-stack-associations", "", "", "User Stack Associations")
	_appstreamCmd.Flags().StringVarP(&_appstreamValidity, "validity", "", "", "Validity")
	_appstreamCmd.Flags().StringVarP(&_appstreamVpcConfig, "vpc-config", "", "", "VPC Config")
	_appstreamCmd.Flags().StringVarP(&_appstreamWorkingDirectory, "working-directory", "", "", "Working Directory")

	_appstreamCmd.Flags().BoolVarP(&_appstreamAssociateAppBlockBuilderAppBlock, "associate-app-block-builder-app-block", "", false, "Associate App Block Builder App Block")
	_appstreamCmd.Flags().BoolVarP(&_appstreamAssociateApplicationFleet, "associate-application-fleet", "", false, "Associate Application Fleet")
	_appstreamCmd.Flags().BoolVarP(&_appstreamAssociateApplicationToEntitlement, "associate-application-to-entitlement", "", false, "Associate Application To Entitlement")
	_appstreamCmd.Flags().BoolVarP(&_appstreamAssociateFleet, "associate-fleet", "", false, "Associate Fleet")
	_appstreamCmd.Flags().BoolVarP(&_appstreamAssociateSoftwareToImageBuilder, "associate-software-to-image-builder", "", false, "Associate Software To Image Builder")
	_appstreamCmd.Flags().BoolVarP(&_appstreamBatchAssociateUserStack, "batch-associate-user-stack", "", false, "Batch Associate User Stack")
	_appstreamCmd.Flags().BoolVarP(&_appstreamBatchDisassociateUserStack, "batch-disassociate-user-stack", "", false, "Batch Disassociate User Stack")
	_appstreamCmd.Flags().BoolVarP(&_appstreamCopyImage, "copy-image", "", false, "Copy Image")
	_appstreamCmd.Flags().BoolVarP(&_appstreamCreateAppBlock, "create-app-block", "", false, "Create App Block")
	_appstreamCmd.Flags().BoolVarP(&_appstreamCreateAppBlockBuilder, "create-app-block-builder", "", false, "Create App Block Builder")
	_appstreamCmd.Flags().BoolVarP(&_appstreamCreateAppBlockBuilderStreamingURL, "create-app-block-builder-streaming-url", "", false, "Create App Block Builder Streaming URL")
	_appstreamCmd.Flags().BoolVarP(&_appstreamCreateApplication, "create-application", "", false, "Create Application")
	_appstreamCmd.Flags().BoolVarP(&_appstreamCreateDirectoryConfig, "create-directory-config", "", false, "Create Directory Config")
	_appstreamCmd.Flags().BoolVarP(&_appstreamCreateEntitlement, "create-entitlement", "", false, "Create Entitlement")
	_appstreamCmd.Flags().BoolVarP(&_appstreamCreateExportImageTask, "create-export-image-task", "", false, "Create Export Image Task")
	_appstreamCmd.Flags().BoolVarP(&_appstreamCreateFleet, "create-fleet", "", false, "Create Fleet")
	_appstreamCmd.Flags().BoolVarP(&_appstreamCreateImageBuilder, "create-image-builder", "", false, "Create Image Builder")
	_appstreamCmd.Flags().BoolVarP(&_appstreamCreateImageBuilderStreamingURL, "create-image-builder-streaming-url", "", false, "Create Image Builder Streaming URL")
	_appstreamCmd.Flags().BoolVarP(&_appstreamCreateImportedImage, "create-imported-image", "", false, "Create Imported Image")
	_appstreamCmd.Flags().BoolVarP(&_appstreamCreateStack, "create-stack", "", false, "Create Stack")
	_appstreamCmd.Flags().BoolVarP(&_appstreamCreateStreamingURL, "create-streaming-url", "", false, "Create Streaming URL")
	_appstreamCmd.Flags().BoolVarP(&_appstreamCreateThemeForStack, "create-theme-for-stack", "", false, "Create Theme For Stack")
	_appstreamCmd.Flags().BoolVarP(&_appstreamCreateUpdatedImage, "create-updated-image", "", false, "Create Updated Image")
	_appstreamCmd.Flags().BoolVarP(&_appstreamCreateUsageReportSubscription, "create-usage-report-subscription", "", false, "Create Usage Report Subscription")
	_appstreamCmd.Flags().BoolVarP(&_appstreamCreateUser, "create-user", "", false, "Create User")
	_appstreamCmd.Flags().BoolVarP(&_appstreamDeleteAppBlock, "delete-app-block", "", false, "Delete App Block")
	_appstreamCmd.Flags().BoolVarP(&_appstreamDeleteAppBlockBuilder, "delete-app-block-builder", "", false, "Delete App Block Builder")
	_appstreamCmd.Flags().BoolVarP(&_appstreamDeleteApplication, "delete-application", "", false, "Delete Application")
	_appstreamCmd.Flags().BoolVarP(&_appstreamDeleteDirectoryConfig, "delete-directory-config", "", false, "Delete Directory Config")
	_appstreamCmd.Flags().BoolVarP(&_appstreamDeleteEntitlement, "delete-entitlement", "", false, "Delete Entitlement")
	_appstreamCmd.Flags().BoolVarP(&_appstreamDeleteFleet, "delete-fleet", "", false, "Delete Fleet")
	_appstreamCmd.Flags().BoolVarP(&_appstreamDeleteImage, "delete-image", "", false, "Delete Image")
	_appstreamCmd.Flags().BoolVarP(&_appstreamDeleteImageBuilder, "delete-image-builder", "", false, "Delete Image Builder")
	_appstreamCmd.Flags().BoolVarP(&_appstreamDeleteImagePermissions, "delete-image-permissions", "", false, "Delete Image Permissions")
	_appstreamCmd.Flags().BoolVarP(&_appstreamDeleteStack, "delete-stack", "", false, "Delete Stack")
	_appstreamCmd.Flags().BoolVarP(&_appstreamDeleteThemeForStack, "delete-theme-for-stack", "", false, "Delete Theme For Stack")
	_appstreamCmd.Flags().BoolVarP(&_appstreamDeleteUsageReportSubscription, "delete-usage-report-subscription", "", false, "Delete Usage Report Subscription")
	_appstreamCmd.Flags().BoolVarP(&_appstreamDeleteUser, "delete-user", "", false, "Delete User")
	_appstreamCmd.Flags().BoolVarP(&_appstreamDescribeAppBlockBuilderAppBlockAssociations, "describe-app-block-builder-app-block-associations", "", false, "Describe App Block Builder App Block Associations")
	_appstreamCmd.Flags().BoolVarP(&_appstreamDescribeAppBlockBuilders, "describe-app-block-builders", "", false, "Describe App Block Builders")
	_appstreamCmd.Flags().BoolVarP(&_appstreamDescribeAppBlocks, "describe-app-blocks", "", false, "Describe App Blocks")
	_appstreamCmd.Flags().BoolVarP(&_appstreamDescribeAppLicenseUsage, "describe-app-license-usage", "", false, "Describe App License Usage")
	_appstreamCmd.Flags().BoolVarP(&_appstreamDescribeApplicationFleetAssociations, "describe-application-fleet-associations", "", false, "Describe Application Fleet Associations")
	_appstreamCmd.Flags().BoolVarP(&_appstreamDescribeApplications, "describe-applications", "", false, "Describe Applications")
	_appstreamCmd.Flags().BoolVarP(&_appstreamDescribeDirectoryConfigs, "describe-directory-configs", "", false, "Describe Directory Configs")
	_appstreamCmd.Flags().BoolVarP(&_appstreamDescribeEntitlements, "describe-entitlements", "", false, "Describe Entitlements")
	_appstreamCmd.Flags().BoolVarP(&_appstreamDescribeFleets, "describe-fleets", "", false, "Describe Fleets")
	_appstreamCmd.Flags().BoolVarP(&_appstreamDescribeImageBuilders, "describe-image-builders", "", false, "Describe Image Builders")
	_appstreamCmd.Flags().BoolVarP(&_appstreamDescribeImagePermissions, "describe-image-permissions", "", false, "Describe Image Permissions")
	_appstreamCmd.Flags().BoolVarP(&_appstreamDescribeImages, "describe-images", "", false, "Describe Images")
	_appstreamCmd.Flags().BoolVarP(&_appstreamDescribeSessions, "describe-sessions", "", false, "Describe Sessions")
	_appstreamCmd.Flags().BoolVarP(&_appstreamDescribeSoftwareAssociations, "describe-software-associations", "", false, "Describe Software Associations")
	_appstreamCmd.Flags().BoolVarP(&_appstreamDescribeStacks, "describe-stacks", "", false, "Describe Stacks")
	_appstreamCmd.Flags().BoolVarP(&_appstreamDescribeThemeForStack, "describe-theme-for-stack", "", false, "Describe Theme For Stack")
	_appstreamCmd.Flags().BoolVarP(&_appstreamDescribeUsageReportSubscriptions, "describe-usage-report-subscriptions", "", false, "Describe Usage Report Subscriptions")
	_appstreamCmd.Flags().BoolVarP(&_appstreamDescribeUserStackAssociations, "describe-user-stack-associations", "", false, "Describe User Stack Associations")
	_appstreamCmd.Flags().BoolVarP(&_appstreamDescribeUsers, "describe-users", "", false, "Describe Users")
	_appstreamCmd.Flags().BoolVarP(&_appstreamDisableUser, "disable-user", "", false, "Disable User")
	_appstreamCmd.Flags().BoolVarP(&_appstreamDisassociateAppBlockBuilderAppBlock, "disassociate-app-block-builder-app-block", "", false, "Disassociate App Block Builder App Block")
	_appstreamCmd.Flags().BoolVarP(&_appstreamDisassociateApplicationFleet, "disassociate-application-fleet", "", false, "Disassociate Application Fleet")
	_appstreamCmd.Flags().BoolVarP(&_appstreamDisassociateApplicationFromEntitlement, "disassociate-application-from-entitlement", "", false, "Disassociate Application From Entitlement")
	_appstreamCmd.Flags().BoolVarP(&_appstreamDisassociateFleet, "disassociate-fleet", "", false, "Disassociate Fleet")
	_appstreamCmd.Flags().BoolVarP(&_appstreamDisassociateSoftwareFromImageBuilder, "disassociate-software-from-image-builder", "", false, "Disassociate Software From Image Builder")
	_appstreamCmd.Flags().BoolVarP(&_appstreamEnableUser, "enable-user", "", false, "Enable User")
	_appstreamCmd.Flags().BoolVarP(&_appstreamExpireSession, "expire-session", "", false, "Expire Session")
	_appstreamCmd.Flags().BoolVarP(&_appstreamGetExportImageTask, "get-export-image-task", "", false, "Get Export Image Task")
	_appstreamCmd.Flags().BoolVarP(&_appstreamListAssociatedFleets, "list-associated-fleets", "", false, "List Associated Fleets")
	_appstreamCmd.Flags().BoolVarP(&_appstreamListAssociatedStacks, "list-associated-stacks", "", false, "List Associated Stacks")
	_appstreamCmd.Flags().BoolVarP(&_appstreamListEntitledApplications, "list-entitled-applications", "", false, "List Entitled Applications")
	_appstreamCmd.Flags().BoolVarP(&_appstreamListExportImageTasks, "list-export-image-tasks", "", false, "List Export Image Tasks")
	_appstreamCmd.Flags().BoolVarP(&_appstreamListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_appstreamCmd.Flags().BoolVarP(&_appstreamStartAppBlockBuilder, "start-app-block-builder", "", false, "Start App Block Builder")
	_appstreamCmd.Flags().BoolVarP(&_appstreamStartFleet, "start-fleet", "", false, "Start Fleet")
	_appstreamCmd.Flags().BoolVarP(&_appstreamStartImageBuilder, "start-image-builder", "", false, "Start Image Builder")
	_appstreamCmd.Flags().BoolVarP(&_appstreamStartSoftwareDeploymentToImageBuilder, "start-software-deployment-to-image-builder", "", false, "Start Software Deployment To Image Builder")
	_appstreamCmd.Flags().BoolVarP(&_appstreamStopAppBlockBuilder, "stop-app-block-builder", "", false, "Stop App Block Builder")
	_appstreamCmd.Flags().BoolVarP(&_appstreamStopFleet, "stop-fleet", "", false, "Stop Fleet")
	_appstreamCmd.Flags().BoolVarP(&_appstreamStopImageBuilder, "stop-image-builder", "", false, "Stop Image Builder")
	_appstreamCmd.Flags().BoolVarP(&_appstreamTagResource, "tag-resource", "", false, "Tag Resource")
	_appstreamCmd.Flags().BoolVarP(&_appstreamUntagResource, "untag-resource", "", false, "Untag Resource")
	_appstreamCmd.Flags().BoolVarP(&_appstreamUpdateAppBlockBuilder, "update-app-block-builder", "", false, "Update App Block Builder")
	_appstreamCmd.Flags().BoolVarP(&_appstreamUpdateApplication, "update-application", "", false, "Update Application")
	_appstreamCmd.Flags().BoolVarP(&_appstreamUpdateDirectoryConfig, "update-directory-config", "", false, "Update Directory Config")
	_appstreamCmd.Flags().BoolVarP(&_appstreamUpdateEntitlement, "update-entitlement", "", false, "Update Entitlement")
	_appstreamCmd.Flags().BoolVarP(&_appstreamUpdateFleet, "update-fleet", "", false, "Update Fleet")
	_appstreamCmd.Flags().BoolVarP(&_appstreamUpdateImagePermissions, "update-image-permissions", "", false, "Update Image Permissions")
	_appstreamCmd.Flags().BoolVarP(&_appstreamUpdateStack, "update-stack", "", false, "Update Stack")
	_appstreamCmd.Flags().BoolVarP(&_appstreamUpdateThemeForStack, "update-theme-for-stack", "", false, "Update Theme For Stack")

}
