package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/workspaces"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// workspacesCmd represents the workspaces command
var _workspacesCmd = &cobra.Command{
	Use:   "workspaces",
	Short: "AWS workspaces CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := workspaces.NewFromConfig(cfg)
		if _workspacesAcceptAccountLinkInvitation {
			workspaces_AcceptAccountLinkInvitation(cfg, client)
			return
		}
		if _workspacesAssociateConnectionAlias {
			workspaces_AssociateConnectionAlias(cfg, client)
			return
		}
		if _workspacesAssociateIpGroups {
			workspaces_AssociateIpGroups(cfg, client)
			return
		}
		if _workspacesAssociateWorkspaceApplication {
			workspaces_AssociateWorkspaceApplication(cfg, client)
			return
		}
		if _workspacesAuthorizeIpRules {
			workspaces_AuthorizeIpRules(cfg, client)
			return
		}
		if _workspacesCopyWorkspaceImage {
			workspaces_CopyWorkspaceImage(cfg, client)
			return
		}
		if _workspacesCreateAccountLinkInvitation {
			workspaces_CreateAccountLinkInvitation(cfg, client)
			return
		}
		if _workspacesCreateConnectClientAddIn {
			workspaces_CreateConnectClientAddIn(cfg, client)
			return
		}
		if _workspacesCreateConnectionAlias {
			workspaces_CreateConnectionAlias(cfg, client)
			return
		}
		if _workspacesCreateIpGroup {
			workspaces_CreateIpGroup(cfg, client)
			return
		}
		if _workspacesCreateStandbyWorkspaces {
			workspaces_CreateStandbyWorkspaces(cfg, client)
			return
		}
		if _workspacesCreateTags {
			workspaces_CreateTags(cfg, client)
			return
		}
		if _workspacesCreateUpdatedWorkspaceImage {
			workspaces_CreateUpdatedWorkspaceImage(cfg, client)
			return
		}
		if _workspacesCreateWorkspaceBundle {
			workspaces_CreateWorkspaceBundle(cfg, client)
			return
		}
		if _workspacesCreateWorkspaceImage {
			workspaces_CreateWorkspaceImage(cfg, client)
			return
		}
		if _workspacesCreateWorkspaces {
			workspaces_CreateWorkspaces(cfg, client)
			return
		}
		if _workspacesCreateWorkspacesPool {
			workspaces_CreateWorkspacesPool(cfg, client)
			return
		}
		if _workspacesDeleteAccountLinkInvitation {
			workspaces_DeleteAccountLinkInvitation(cfg, client)
			return
		}
		if _workspacesDeleteClientBranding {
			workspaces_DeleteClientBranding(cfg, client)
			return
		}
		if _workspacesDeleteConnectClientAddIn {
			workspaces_DeleteConnectClientAddIn(cfg, client)
			return
		}
		if _workspacesDeleteConnectionAlias {
			workspaces_DeleteConnectionAlias(cfg, client)
			return
		}
		if _workspacesDeleteIpGroup {
			workspaces_DeleteIpGroup(cfg, client)
			return
		}
		if _workspacesDeleteTags {
			workspaces_DeleteTags(cfg, client)
			return
		}
		if _workspacesDeleteWorkspaceBundle {
			workspaces_DeleteWorkspaceBundle(cfg, client)
			return
		}
		if _workspacesDeleteWorkspaceImage {
			workspaces_DeleteWorkspaceImage(cfg, client)
			return
		}
		if _workspacesDeployWorkspaceApplications {
			workspaces_DeployWorkspaceApplications(cfg, client)
			return
		}
		if _workspacesDeregisterWorkspaceDirectory {
			workspaces_DeregisterWorkspaceDirectory(cfg, client)
			return
		}
		if _workspacesDescribeAccount {
			workspaces_DescribeAccount(cfg, client)
			return
		}
		if _workspacesDescribeAccountModifications {
			workspaces_DescribeAccountModifications(cfg, client)
			return
		}
		if _workspacesDescribeApplicationAssociations {
			workspaces_DescribeApplicationAssociations(cfg, client)
			return
		}
		if _workspacesDescribeApplications {
			workspaces_DescribeApplications(cfg, client)
			return
		}
		if _workspacesDescribeBundleAssociations {
			workspaces_DescribeBundleAssociations(cfg, client)
			return
		}
		if _workspacesDescribeClientBranding {
			workspaces_DescribeClientBranding(cfg, client)
			return
		}
		if _workspacesDescribeClientProperties {
			workspaces_DescribeClientProperties(cfg, client)
			return
		}
		if _workspacesDescribeConnectClientAddIns {
			workspaces_DescribeConnectClientAddIns(cfg, client)
			return
		}
		if _workspacesDescribeConnectionAliasPermissions {
			workspaces_DescribeConnectionAliasPermissions(cfg, client)
			return
		}
		if _workspacesDescribeConnectionAliases {
			workspaces_DescribeConnectionAliases(cfg, client)
			return
		}
		if _workspacesDescribeCustomWorkspaceImageImport {
			workspaces_DescribeCustomWorkspaceImageImport(cfg, client)
			return
		}
		if _workspacesDescribeImageAssociations {
			workspaces_DescribeImageAssociations(cfg, client)
			return
		}
		if _workspacesDescribeIpGroups {
			workspaces_DescribeIpGroups(cfg, client)
			return
		}
		if _workspacesDescribeTags {
			workspaces_DescribeTags(cfg, client)
			return
		}
		if _workspacesDescribeWorkspaceAssociations {
			workspaces_DescribeWorkspaceAssociations(cfg, client)
			return
		}
		if _workspacesDescribeWorkspaceBundles {
			workspaces_DescribeWorkspaceBundles(cfg, client)
			return
		}
		if _workspacesDescribeWorkspaceDirectories {
			workspaces_DescribeWorkspaceDirectories(cfg, client)
			return
		}
		if _workspacesDescribeWorkspaceImagePermissions {
			workspaces_DescribeWorkspaceImagePermissions(cfg, client)
			return
		}
		if _workspacesDescribeWorkspaceImages {
			workspaces_DescribeWorkspaceImages(cfg, client)
			return
		}
		if _workspacesDescribeWorkspaceSnapshots {
			workspaces_DescribeWorkspaceSnapshots(cfg, client)
			return
		}
		if _workspacesDescribeWorkspaces {
			workspaces_DescribeWorkspaces(cfg, client)
			return
		}
		if _workspacesDescribeWorkspacesConnectionStatus {
			workspaces_DescribeWorkspacesConnectionStatus(cfg, client)
			return
		}
		if _workspacesDescribeWorkspacesPoolSessions {
			workspaces_DescribeWorkspacesPoolSessions(cfg, client)
			return
		}
		if _workspacesDescribeWorkspacesPools {
			workspaces_DescribeWorkspacesPools(cfg, client)
			return
		}
		if _workspacesDisassociateConnectionAlias {
			workspaces_DisassociateConnectionAlias(cfg, client)
			return
		}
		if _workspacesDisassociateIpGroups {
			workspaces_DisassociateIpGroups(cfg, client)
			return
		}
		if _workspacesDisassociateWorkspaceApplication {
			workspaces_DisassociateWorkspaceApplication(cfg, client)
			return
		}
		if _workspacesGetAccountLink {
			workspaces_GetAccountLink(cfg, client)
			return
		}
		if _workspacesImportClientBranding {
			workspaces_ImportClientBranding(cfg, client)
			return
		}
		if _workspacesImportCustomWorkspaceImage {
			workspaces_ImportCustomWorkspaceImage(cfg, client)
			return
		}
		if _workspacesImportWorkspaceImage {
			workspaces_ImportWorkspaceImage(cfg, client)
			return
		}
		if _workspacesListAccountLinks {
			workspaces_ListAccountLinks(cfg, client)
			return
		}
		if _workspacesListAvailableManagementCidrRanges {
			workspaces_ListAvailableManagementCidrRanges(cfg, client)
			return
		}
		if _workspacesMigrateWorkspace {
			workspaces_MigrateWorkspace(cfg, client)
			return
		}
		if _workspacesModifyAccount {
			workspaces_ModifyAccount(cfg, client)
			return
		}
		if _workspacesModifyCertificateBasedAuthProperties {
			workspaces_ModifyCertificateBasedAuthProperties(cfg, client)
			return
		}
		if _workspacesModifyClientProperties {
			workspaces_ModifyClientProperties(cfg, client)
			return
		}
		if _workspacesModifyEndpointEncryptionMode {
			workspaces_ModifyEndpointEncryptionMode(cfg, client)
			return
		}
		if _workspacesModifySamlProperties {
			workspaces_ModifySamlProperties(cfg, client)
			return
		}
		if _workspacesModifySelfservicePermissions {
			workspaces_ModifySelfservicePermissions(cfg, client)
			return
		}
		if _workspacesModifyStreamingProperties {
			workspaces_ModifyStreamingProperties(cfg, client)
			return
		}
		if _workspacesModifyWorkspaceAccessProperties {
			workspaces_ModifyWorkspaceAccessProperties(cfg, client)
			return
		}
		if _workspacesModifyWorkspaceCreationProperties {
			workspaces_ModifyWorkspaceCreationProperties(cfg, client)
			return
		}
		if _workspacesModifyWorkspaceProperties {
			workspaces_ModifyWorkspaceProperties(cfg, client)
			return
		}
		if _workspacesModifyWorkspaceState {
			workspaces_ModifyWorkspaceState(cfg, client)
			return
		}
		if _workspacesRebootWorkspaces {
			workspaces_RebootWorkspaces(cfg, client)
			return
		}
		if _workspacesRebuildWorkspaces {
			workspaces_RebuildWorkspaces(cfg, client)
			return
		}
		if _workspacesRegisterWorkspaceDirectory {
			workspaces_RegisterWorkspaceDirectory(cfg, client)
			return
		}
		if _workspacesRejectAccountLinkInvitation {
			workspaces_RejectAccountLinkInvitation(cfg, client)
			return
		}
		if _workspacesRestoreWorkspace {
			workspaces_RestoreWorkspace(cfg, client)
			return
		}
		if _workspacesRevokeIpRules {
			workspaces_RevokeIpRules(cfg, client)
			return
		}
		if _workspacesStartWorkspaces {
			workspaces_StartWorkspaces(cfg, client)
			return
		}
		if _workspacesStartWorkspacesPool {
			workspaces_StartWorkspacesPool(cfg, client)
			return
		}
		if _workspacesStopWorkspaces {
			workspaces_StopWorkspaces(cfg, client)
			return
		}
		if _workspacesStopWorkspacesPool {
			workspaces_StopWorkspacesPool(cfg, client)
			return
		}
		if _workspacesTerminateWorkspaces {
			workspaces_TerminateWorkspaces(cfg, client)
			return
		}
		if _workspacesTerminateWorkspacesPool {
			workspaces_TerminateWorkspacesPool(cfg, client)
			return
		}
		if _workspacesTerminateWorkspacesPoolSession {
			workspaces_TerminateWorkspacesPoolSession(cfg, client)
			return
		}
		if _workspacesUpdateConnectClientAddIn {
			workspaces_UpdateConnectClientAddIn(cfg, client)
			return
		}
		if _workspacesUpdateConnectionAliasPermission {
			workspaces_UpdateConnectionAliasPermission(cfg, client)
			return
		}
		if _workspacesUpdateRulesOfIpGroup {
			workspaces_UpdateRulesOfIpGroup(cfg, client)
			return
		}
		if _workspacesUpdateWorkspaceBundle {
			workspaces_UpdateWorkspaceBundle(cfg, client)
			return
		}
		if _workspacesUpdateWorkspaceImagePermission {
			workspaces_UpdateWorkspaceImagePermission(cfg, client)
			return
		}
		if _workspacesUpdateWorkspacesPool {
			workspaces_UpdateWorkspacesPool(cfg, client)
			return
		}

	},
}

var (
	_workspacesAcceptAccountLinkInvitation          bool
	_workspacesAssociateConnectionAlias             bool
	_workspacesAssociateIpGroups                    bool
	_workspacesAssociateWorkspaceApplication        bool
	_workspacesAuthorizeIpRules                     bool
	_workspacesCopyWorkspaceImage                   bool
	_workspacesCreateAccountLinkInvitation          bool
	_workspacesCreateConnectClientAddIn             bool
	_workspacesCreateConnectionAlias                bool
	_workspacesCreateIpGroup                        bool
	_workspacesCreateStandbyWorkspaces              bool
	_workspacesCreateTags                           bool
	_workspacesCreateUpdatedWorkspaceImage          bool
	_workspacesCreateWorkspaceBundle                bool
	_workspacesCreateWorkspaceImage                 bool
	_workspacesCreateWorkspaces                     bool
	_workspacesCreateWorkspacesPool                 bool
	_workspacesDeleteAccountLinkInvitation          bool
	_workspacesDeleteClientBranding                 bool
	_workspacesDeleteConnectClientAddIn             bool
	_workspacesDeleteConnectionAlias                bool
	_workspacesDeleteIpGroup                        bool
	_workspacesDeleteTags                           bool
	_workspacesDeleteWorkspaceBundle                bool
	_workspacesDeleteWorkspaceImage                 bool
	_workspacesDeployWorkspaceApplications          bool
	_workspacesDeregisterWorkspaceDirectory         bool
	_workspacesDescribeAccount                      bool
	_workspacesDescribeAccountModifications         bool
	_workspacesDescribeApplicationAssociations      bool
	_workspacesDescribeApplications                 bool
	_workspacesDescribeBundleAssociations           bool
	_workspacesDescribeClientBranding               bool
	_workspacesDescribeClientProperties             bool
	_workspacesDescribeConnectClientAddIns          bool
	_workspacesDescribeConnectionAliasPermissions   bool
	_workspacesDescribeConnectionAliases            bool
	_workspacesDescribeCustomWorkspaceImageImport   bool
	_workspacesDescribeImageAssociations            bool
	_workspacesDescribeIpGroups                     bool
	_workspacesDescribeTags                         bool
	_workspacesDescribeWorkspaceAssociations        bool
	_workspacesDescribeWorkspaceBundles             bool
	_workspacesDescribeWorkspaceDirectories         bool
	_workspacesDescribeWorkspaceImagePermissions    bool
	_workspacesDescribeWorkspaceImages              bool
	_workspacesDescribeWorkspaceSnapshots           bool
	_workspacesDescribeWorkspaces                   bool
	_workspacesDescribeWorkspacesConnectionStatus   bool
	_workspacesDescribeWorkspacesPoolSessions       bool
	_workspacesDescribeWorkspacesPools              bool
	_workspacesDisassociateConnectionAlias          bool
	_workspacesDisassociateIpGroups                 bool
	_workspacesDisassociateWorkspaceApplication     bool
	_workspacesGetAccountLink                       bool
	_workspacesImportClientBranding                 bool
	_workspacesImportCustomWorkspaceImage           bool
	_workspacesImportWorkspaceImage                 bool
	_workspacesListAccountLinks                     bool
	_workspacesListAvailableManagementCidrRanges    bool
	_workspacesMigrateWorkspace                     bool
	_workspacesModifyAccount                        bool
	_workspacesModifyCertificateBasedAuthProperties bool
	_workspacesModifyClientProperties               bool
	_workspacesModifyEndpointEncryptionMode         bool
	_workspacesModifySamlProperties                 bool
	_workspacesModifySelfservicePermissions         bool
	_workspacesModifyStreamingProperties            bool
	_workspacesModifyWorkspaceAccessProperties      bool
	_workspacesModifyWorkspaceCreationProperties    bool
	_workspacesModifyWorkspaceProperties            bool
	_workspacesModifyWorkspaceState                 bool
	_workspacesRebootWorkspaces                     bool
	_workspacesRebuildWorkspaces                    bool
	_workspacesRegisterWorkspaceDirectory           bool
	_workspacesRejectAccountLinkInvitation          bool
	_workspacesRestoreWorkspace                     bool
	_workspacesRevokeIpRules                        bool
	_workspacesStartWorkspaces                      bool
	_workspacesStartWorkspacesPool                  bool
	_workspacesStopWorkspaces                       bool
	_workspacesStopWorkspacesPool                   bool
	_workspacesTerminateWorkspaces                  bool
	_workspacesTerminateWorkspacesPool              bool
	_workspacesTerminateWorkspacesPoolSession       bool
	_workspacesUpdateConnectClientAddIn             bool
	_workspacesUpdateConnectionAliasPermission      bool
	_workspacesUpdateRulesOfIpGroup                 bool
	_workspacesUpdateWorkspaceBundle                bool
	_workspacesUpdateWorkspaceImagePermission       bool
	_workspacesUpdateWorkspacesPool                 bool

	_workspacesActiveDirectoryConfig               string
	_workspacesAddInId                             string
	_workspacesAliasId                             string
	_workspacesAliasIds                            []string
	_workspacesAllowCopyImage                      string
	_workspacesApplicationId                       string
	_workspacesApplicationIds                      []string
	_workspacesApplicationSettings                 string
	_workspacesApplications                        string
	_workspacesAssociatedResourceTypes             string
	_workspacesBundleDescription                   string
	_workspacesBundleId                            string
	_workspacesBundleIds                           []string
	_workspacesBundleName                          string
	_workspacesCapacity                            string
	_workspacesCertificateBasedAuthProperties      string
	_workspacesClientProperties                    string
	_workspacesClientToken                         string
	_workspacesComputeType                         string
	_workspacesComputeTypeNames                    string
	_workspacesConnectionAliasPermission           string
	_workspacesConnectionString                    string
	_workspacesDataReplication                     string
	_workspacesDedicatedTenancyManagementCidrRange string
	_workspacesDedicatedTenancySupport             string
	_workspacesDescription                         string
	_workspacesDeviceTypeAndroid                   string
	_workspacesDeviceTypeIos                       string
	_workspacesDeviceTypeLinux                     string
	_workspacesDeviceTypeOsx                       string
	_workspacesDeviceTypeWeb                       string
	_workspacesDeviceTypeWindows                   string
	_workspacesDirectoryId                         string
	_workspacesDirectoryIds                        []string
	_workspacesEc2ImageId                          string
	_workspacesEnableSelfService                   string
	_workspacesEndpointEncryptionMode              string
	_workspacesFilters                             string
	_workspacesForce                               string
	_workspacesGroupDesc                           string
	_workspacesGroupId                             string
	_workspacesGroupIds                            []string
	_workspacesGroupName                           string
	_workspacesIdcInstanceArn                      string
	_workspacesImageDescription                    string
	_workspacesImageId                             string
	_workspacesImageIds                            []string
	_workspacesImageName                           string
	_workspacesImageSource                         string
	_workspacesImageType                           string
	_workspacesInfrastructureConfigurationArn      string
	_workspacesIngestionProcess                    string
	_workspacesLicenseType                         string
	_workspacesLimit                               string
	_workspacesLinkId                              string
	_workspacesLinkStatusFilter                    string
	_workspacesLinkedAccountId                     string
	_workspacesManagementCidrRangeConstraint       string
	_workspacesMaxResults                          string
	_workspacesMicrosoftEntraConfig                string
	_workspacesName                                string
	_workspacesNextToken                           string
	_workspacesOperatingSystemNames                string
	_workspacesOsVersion                           string
	_workspacesOwner                               string
	_workspacesPlatform                            string
	_workspacesPlatforms                           string
	_workspacesPoolId                              string
	_workspacesPoolIds                             []string
	_workspacesPoolName                            string
	_workspacesPrimaryRegion                       string
	_workspacesPropertiesToDelete                  string
	_workspacesProtocol                            string
	_workspacesRebootWorkspaceRequests             string
	_workspacesRebuildWorkspaceRequests            string
	_workspacesResourceId                          string
	_workspacesResourceIds                         []string
	_workspacesRootStorage                         string
	_workspacesRunningMode                         string
	_workspacesSamlProperties                      string
	_workspacesSelfservicePermissions              string
	_workspacesSessionId                           string
	_workspacesSharedAccountId                     string
	_workspacesSourceImageId                       string
	_workspacesSourceRegion                        string
	_workspacesSourceWorkspaceId                   string
	_workspacesStandbyWorkspaces                   string
	_workspacesStartWorkspaceRequests              string
	_workspacesStopWorkspaceRequests               string
	_workspacesStreamingProperties                 string
	_workspacesSubnetIds                           []string
	_workspacesTagKeys                             []string
	_workspacesTags                                string
	_workspacesTargetAccountId                     string
	_workspacesTenancy                             string
	_workspacesTerminateWorkspaceRequests          string
	_workspacesTimeoutSettings                     string
	_workspacesURL                                 string
	_workspacesUserId                              string
	_workspacesUserIdentityType                    string
	_workspacesUserName                            string
	_workspacesUserRules                           string
	_workspacesUserStorage                         string
	_workspacesWorkspaceAccessProperties           string
	_workspacesWorkspaceCreationProperties         string
	_workspacesWorkspaceDirectoryDescription       string
	_workspacesWorkspaceDirectoryName              string
	_workspacesWorkspaceDirectoryNames             []string
	_workspacesWorkspaceId                         string
	_workspacesWorkspaceIds                        []string
	_workspacesWorkspaceName                       string
	_workspacesWorkspaceProperties                 string
	_workspacesWorkspaceState                      string
	_workspacesWorkspaceType                       string
	_workspacesWorkspaces                          string
)

// Accepts the account link invitation.
// There's currently no unlinking capability after you accept the account linking
// invitation.
func workspaces_AcceptAccountLinkInvitation(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.AcceptAccountLinkInvitationInput{
		// LinkId: *string, // Required
	}

	if len(_workspacesLinkId) > 0 {
		input.LinkId = aws.String(_workspacesLinkId)
	}
	if len(_workspacesClientToken) > 0 {
		input.ClientToken = aws.String(_workspacesClientToken)
	}

	if resp, err := client.AcceptAccountLinkInvitation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates the specified connection alias with the specified directory to
// enable cross-Region redirection. For more information, see [Cross-Region Redirection for Amazon WorkSpaces].
//
// Before performing this operation, call [DescribeConnectionAliases] to make sure that the current state of
// the connection alias is CREATED .
//
// [DescribeConnectionAliases]: https://docs.aws.amazon.com/workspaces/latest/api/API_DescribeConnectionAliases.html
// [Cross-Region Redirection for Amazon WorkSpaces]: https://docs.aws.amazon.com/workspaces/latest/adminguide/cross-region-redirection.html
func workspaces_AssociateConnectionAlias(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.AssociateConnectionAliasInput{
		// AliasId: *string, // Required
		// ResourceId: *string, // Required
	}

	if len(_workspacesAliasId) > 0 {
		input.AliasId = aws.String(_workspacesAliasId)
	}
	if len(_workspacesResourceId) > 0 {
		input.ResourceId = aws.String(_workspacesResourceId)
	}

	if resp, err := client.AssociateConnectionAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates the specified IP access control group with the specified directory.
func workspaces_AssociateIpGroups(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.AssociateIpGroupsInput{
		// DirectoryId: *string, // Required
		// GroupIds: []string, // Required
	}

	if len(_workspacesDirectoryId) > 0 {
		input.DirectoryId = aws.String(_workspacesDirectoryId)
	}
	if len(_workspacesGroupIds) > 0 {
		input.GroupIds = append([]string(nil), _workspacesGroupIds...)
	}

	if resp, err := client.AssociateIpGroups(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates the specified application to the specified WorkSpace.
func workspaces_AssociateWorkspaceApplication(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.AssociateWorkspaceApplicationInput{
		// ApplicationId: *string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_workspacesApplicationId) > 0 {
		input.ApplicationId = aws.String(_workspacesApplicationId)
	}
	if len(_workspacesWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_workspacesWorkspaceId)
	}

	if resp, err := client.AssociateWorkspaceApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds one or more rules to the specified IP access control group.
// This action gives users permission to access their WorkSpaces from the CIDR
// address ranges specified in the rules.
func workspaces_AuthorizeIpRules(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.AuthorizeIpRulesInput{
		// GroupId: *string, // Required
		// UserRules: []types.IpRuleItem, // Required
	}

	if len(_workspacesGroupId) > 0 {
		input.GroupId = aws.String(_workspacesGroupId)
	}
	if len(_workspacesUserRules) > 0 {
		if err := assignInputField(input, "UserRules", _workspacesUserRules); err != nil {
			log.Errorf("invalid --user-rules: %s", err.Error())
			return
		}
	}

	if resp, err := client.AuthorizeIpRules(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Copies the specified image from the specified Region to the current Region. For
// more information about copying images, see [Copy a Custom WorkSpaces Image].
//
// In the China (Ningxia) Region, you can copy images only within the same Region.
//
// In Amazon Web Services GovCloud (US), to copy images to and from other Regions,
// contact Amazon Web Services Support.
//
// Before copying a shared image, be sure to verify that it has been shared from
// the correct Amazon Web Services account. To determine if an image has been
// shared and to see the ID of the Amazon Web Services account that owns an image,
// use the [DescribeWorkSpaceImages]and [DescribeWorkspaceImagePermissions] API operations.
//
// [DescribeWorkspaceImagePermissions]: https://docs.aws.amazon.com/workspaces/latest/api/API_DescribeWorkspaceImagePermissions.html
// [DescribeWorkSpaceImages]: https://docs.aws.amazon.com/workspaces/latest/api/API_DescribeWorkspaceImages.html
// [Copy a Custom WorkSpaces Image]: https://docs.aws.amazon.com/workspaces/latest/adminguide/copy-custom-image.html
func workspaces_CopyWorkspaceImage(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.CopyWorkspaceImageInput{
		// Name: *string, // Required
		// SourceImageId: *string, // Required
		// SourceRegion: *string, // Required
	}

	if len(_workspacesName) > 0 {
		input.Name = aws.String(_workspacesName)
	}
	if len(_workspacesSourceImageId) > 0 {
		input.SourceImageId = aws.String(_workspacesSourceImageId)
	}
	if len(_workspacesSourceRegion) > 0 {
		input.SourceRegion = aws.String(_workspacesSourceRegion)
	}
	if len(_workspacesDescription) > 0 {
		input.Description = aws.String(_workspacesDescription)
	}
	if len(_workspacesTags) > 0 {
		if err := assignInputField(input, "Tags", _workspacesTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CopyWorkspaceImage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates the account link invitation.
func workspaces_CreateAccountLinkInvitation(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.CreateAccountLinkInvitationInput{
		// TargetAccountId: *string, // Required
	}

	if len(_workspacesTargetAccountId) > 0 {
		input.TargetAccountId = aws.String(_workspacesTargetAccountId)
	}
	if len(_workspacesClientToken) > 0 {
		input.ClientToken = aws.String(_workspacesClientToken)
	}

	if resp, err := client.CreateAccountLinkInvitation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a client-add-in for Amazon Connect within a directory. You can create
// only one Amazon Connect client add-in within a directory.
//
// This client add-in allows WorkSpaces users to seamlessly connect to Amazon
// Connect.
func workspaces_CreateConnectClientAddIn(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.CreateConnectClientAddInInput{
		// Name: *string, // Required
		// ResourceId: *string, // Required
		// URL: *string, // Required
	}

	if len(_workspacesName) > 0 {
		input.Name = aws.String(_workspacesName)
	}
	if len(_workspacesResourceId) > 0 {
		input.ResourceId = aws.String(_workspacesResourceId)
	}
	if len(_workspacesURL) > 0 {
		input.URL = aws.String(_workspacesURL)
	}

	if resp, err := client.CreateConnectClientAddIn(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates the specified connection alias for use with cross-Region redirection.
// For more information, see [Cross-Region Redirection for Amazon WorkSpaces].
//
// [Cross-Region Redirection for Amazon WorkSpaces]: https://docs.aws.amazon.com/workspaces/latest/adminguide/cross-region-redirection.html
func workspaces_CreateConnectionAlias(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.CreateConnectionAliasInput{
		// ConnectionString: *string, // Required
	}

	if len(_workspacesConnectionString) > 0 {
		input.ConnectionString = aws.String(_workspacesConnectionString)
	}
	if len(_workspacesTags) > 0 {
		if err := assignInputField(input, "Tags", _workspacesTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateConnectionAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an IP access control group.
// An IP access control group provides you with the ability to control the IP
// addresses from which users are allowed to access their WorkSpaces. To specify
// the CIDR address ranges, add rules to your IP access control group and then
// associate the group with your directory. You can add rules when you create the
// group or at any time using AuthorizeIpRules.
//
// There is a default IP access control group associated with your directory. If
// you don't associate an IP access control group with your directory, the default
// group is used. The default group includes a default rule that allows users to
// access their WorkSpaces from anywhere. You cannot modify the default IP access
// control group for your directory.
func workspaces_CreateIpGroup(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.CreateIpGroupInput{
		// GroupName: *string, // Required
	}

	if len(_workspacesGroupName) > 0 {
		input.GroupName = aws.String(_workspacesGroupName)
	}
	if len(_workspacesGroupDesc) > 0 {
		input.GroupDesc = aws.String(_workspacesGroupDesc)
	}
	if len(_workspacesTags) > 0 {
		if err := assignInputField(input, "Tags", _workspacesTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_workspacesUserRules) > 0 {
		if err := assignInputField(input, "UserRules", _workspacesUserRules); err != nil {
			log.Errorf("invalid --user-rules: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateIpGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a standby WorkSpace in a secondary Region.
func workspaces_CreateStandbyWorkspaces(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.CreateStandbyWorkspacesInput{
		// PrimaryRegion: *string, // Required
		// StandbyWorkspaces: []types.StandbyWorkspace, // Required
	}

	if len(_workspacesPrimaryRegion) > 0 {
		input.PrimaryRegion = aws.String(_workspacesPrimaryRegion)
	}
	if len(_workspacesStandbyWorkspaces) > 0 {
		if err := assignInputField(input, "StandbyWorkspaces", _workspacesStandbyWorkspaces); err != nil {
			log.Errorf("invalid --standby-workspaces: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateStandbyWorkspaces(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates the specified tags for the specified WorkSpaces resource.
func workspaces_CreateTags(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.CreateTagsInput{
		// ResourceId: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_workspacesResourceId) > 0 {
		input.ResourceId = aws.String(_workspacesResourceId)
	}
	if len(_workspacesTags) > 0 {
		if err := assignInputField(input, "Tags", _workspacesTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTags(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new updated WorkSpace image based on the specified source image. The
// new updated WorkSpace image has the latest drivers and other updates required by
// the Amazon WorkSpaces components.
//
// To determine which WorkSpace images need to be updated with the latest Amazon
// WorkSpaces requirements, use [DescribeWorkspaceImages].
//
// - Only Windows 10, Windows Server 2016, and Windows Server 2019 WorkSpace
// images can be programmatically updated at this time.
//
// - Microsoft Windows updates and other application updates are not included in
// the update process.
//
// - The source WorkSpace image is not deleted. You can delete the source image
// after you've verified your new updated image and created a new bundle.
//
// [DescribeWorkspaceImages]: https://docs.aws.amazon.com/workspaces/latest/api/API_DescribeWorkspaceImages.html
func workspaces_CreateUpdatedWorkspaceImage(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.CreateUpdatedWorkspaceImageInput{
		// Description: *string, // Required
		// Name: *string, // Required
		// SourceImageId: *string, // Required
	}

	if len(_workspacesDescription) > 0 {
		input.Description = aws.String(_workspacesDescription)
	}
	if len(_workspacesName) > 0 {
		input.Name = aws.String(_workspacesName)
	}
	if len(_workspacesSourceImageId) > 0 {
		input.SourceImageId = aws.String(_workspacesSourceImageId)
	}
	if len(_workspacesTags) > 0 {
		if err := assignInputField(input, "Tags", _workspacesTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateUpdatedWorkspaceImage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates the specified WorkSpace bundle. For more information about creating
// WorkSpace bundles, see [Create a Custom WorkSpaces Image and Bundle].
//
// [Create a Custom WorkSpaces Image and Bundle]: https://docs.aws.amazon.com/workspaces/latest/adminguide/create-custom-bundle.html
func workspaces_CreateWorkspaceBundle(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.CreateWorkspaceBundleInput{
		// BundleDescription: *string, // Required
		// BundleName: *string, // Required
		// ComputeType: *types.ComputeType, // Required
		// ImageId: *string, // Required
		// UserStorage: *types.UserStorage, // Required
	}

	if len(_workspacesBundleDescription) > 0 {
		input.BundleDescription = aws.String(_workspacesBundleDescription)
	}
	if len(_workspacesBundleName) > 0 {
		input.BundleName = aws.String(_workspacesBundleName)
	}
	if len(_workspacesComputeType) > 0 {
		if err := assignInputField(input, "ComputeType", _workspacesComputeType); err != nil {
			log.Errorf("invalid --compute-type: %s", err.Error())
			return
		}
	}
	if len(_workspacesImageId) > 0 {
		input.ImageId = aws.String(_workspacesImageId)
	}
	if len(_workspacesUserStorage) > 0 {
		if err := assignInputField(input, "UserStorage", _workspacesUserStorage); err != nil {
			log.Errorf("invalid --user-storage: %s", err.Error())
			return
		}
	}
	if len(_workspacesRootStorage) > 0 {
		if err := assignInputField(input, "RootStorage", _workspacesRootStorage); err != nil {
			log.Errorf("invalid --root-storage: %s", err.Error())
			return
		}
	}
	if len(_workspacesTags) > 0 {
		if err := assignInputField(input, "Tags", _workspacesTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateWorkspaceBundle(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new WorkSpace image from an existing WorkSpace.
func workspaces_CreateWorkspaceImage(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.CreateWorkspaceImageInput{
		// Description: *string, // Required
		// Name: *string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_workspacesDescription) > 0 {
		input.Description = aws.String(_workspacesDescription)
	}
	if len(_workspacesName) > 0 {
		input.Name = aws.String(_workspacesName)
	}
	if len(_workspacesWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_workspacesWorkspaceId)
	}
	if len(_workspacesTags) > 0 {
		if err := assignInputField(input, "Tags", _workspacesTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateWorkspaceImage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates one or more WorkSpaces.
// This operation is asynchronous and returns before the WorkSpaces are created.
//
// - The MANUAL running mode value is only supported by Amazon WorkSpaces Core.
// Contact your account team to be allow-listed to use this value. For more
// information, see [Amazon WorkSpaces Core].
//
// - You don't need to specify the PCOIP protocol for Linux bundles because DCV
// (formerly WSP) is the default protocol for those bundles.
//
// - User-decoupled WorkSpaces are only supported by Amazon WorkSpaces Core.
//
// - Review your running mode to ensure you are using one that is optimal for
// your needs and budget. For more information on switching running modes, see [Can I switch between hourly and monthly billing?]
//
// [Can I switch between hourly and monthly billing?]: http://aws.amazon.com/workspaces-family/workspaces/faqs/#:~:text=Can%20I%20switch%20between%20hourly%20and%20monthly%20billing%20on%20WorkSpaces%20Personal%3F
// [Amazon WorkSpaces Core]: http://aws.amazon.com/workspaces/core/
func workspaces_CreateWorkspaces(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.CreateWorkspacesInput{
		// Workspaces: []types.WorkspaceRequest, // Required
	}

	if len(_workspacesWorkspaces) > 0 {
		if err := assignInputField(input, "Workspaces", _workspacesWorkspaces); err != nil {
			log.Errorf("invalid --workspaces: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateWorkspaces(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a pool of WorkSpaces.
func workspaces_CreateWorkspacesPool(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.CreateWorkspacesPoolInput{
		// BundleId: *string, // Required
		// Capacity: *types.Capacity, // Required
		// Description: *string, // Required
		// DirectoryId: *string, // Required
		// PoolName: *string, // Required
	}

	if len(_workspacesBundleId) > 0 {
		input.BundleId = aws.String(_workspacesBundleId)
	}
	if len(_workspacesCapacity) > 0 {
		if err := assignInputField(input, "Capacity", _workspacesCapacity); err != nil {
			log.Errorf("invalid --capacity: %s", err.Error())
			return
		}
	}
	if len(_workspacesDescription) > 0 {
		input.Description = aws.String(_workspacesDescription)
	}
	if len(_workspacesDirectoryId) > 0 {
		input.DirectoryId = aws.String(_workspacesDirectoryId)
	}
	if len(_workspacesPoolName) > 0 {
		input.PoolName = aws.String(_workspacesPoolName)
	}
	if len(_workspacesApplicationSettings) > 0 {
		if err := assignInputField(input, "ApplicationSettings", _workspacesApplicationSettings); err != nil {
			log.Errorf("invalid --application-settings: %s", err.Error())
			return
		}
	}
	if len(_workspacesRunningMode) > 0 {
		if err := assignInputField(input, "RunningMode", _workspacesRunningMode); err != nil {
			log.Errorf("invalid --running-mode: %s", err.Error())
			return
		}
	}
	if len(_workspacesTags) > 0 {
		if err := assignInputField(input, "Tags", _workspacesTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_workspacesTimeoutSettings) > 0 {
		if err := assignInputField(input, "TimeoutSettings", _workspacesTimeoutSettings); err != nil {
			log.Errorf("invalid --timeout-settings: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateWorkspacesPool(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the account link invitation.
func workspaces_DeleteAccountLinkInvitation(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.DeleteAccountLinkInvitationInput{
		// LinkId: *string, // Required
	}

	if len(_workspacesLinkId) > 0 {
		input.LinkId = aws.String(_workspacesLinkId)
	}
	if len(_workspacesClientToken) > 0 {
		input.ClientToken = aws.String(_workspacesClientToken)
	}

	if resp, err := client.DeleteAccountLinkInvitation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes customized client branding. Client branding allows you to customize
// your WorkSpace's client login portal. You can tailor your login portal company
// logo, the support email address, support link, link to reset password, and a
// custom message for users trying to sign in.
//
// After you delete your customized client branding, your login portal reverts to
// the default client branding.
func workspaces_DeleteClientBranding(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.DeleteClientBrandingInput{
		// Platforms: []types.ClientDeviceType, // Required
		// ResourceId: *string, // Required
	}

	if len(_workspacesPlatforms) > 0 {
		if err := assignInputField(input, "Platforms", _workspacesPlatforms); err != nil {
			log.Errorf("invalid --platforms: %s", err.Error())
			return
		}
	}
	if len(_workspacesResourceId) > 0 {
		input.ResourceId = aws.String(_workspacesResourceId)
	}

	if resp, err := client.DeleteClientBranding(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a client-add-in for Amazon Connect that is configured within a
// directory.
func workspaces_DeleteConnectClientAddIn(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.DeleteConnectClientAddInInput{
		// AddInId: *string, // Required
		// ResourceId: *string, // Required
	}

	if len(_workspacesAddInId) > 0 {
		input.AddInId = aws.String(_workspacesAddInId)
	}
	if len(_workspacesResourceId) > 0 {
		input.ResourceId = aws.String(_workspacesResourceId)
	}

	if resp, err := client.DeleteConnectClientAddIn(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified connection alias. For more information, see [Cross-Region Redirection for Amazon WorkSpaces].
// If you will no longer be using a fully qualified domain name (FQDN) as the
// registration code for your WorkSpaces users, you must take certain precautions
// to prevent potential security issues. For more information, see [Security Considerations if You Stop Using Cross-Region Redirection].
//
// To delete a connection alias that has been shared, the shared account must
// first disassociate the connection alias from any directories it has been
// associated with. Then you must unshare the connection alias from the account it
// has been shared with. You can delete a connection alias only after it is no
// longer shared with any accounts or associated with any directories.
//
// [Security Considerations if You Stop Using Cross-Region Redirection]: https://docs.aws.amazon.com/workspaces/latest/adminguide/cross-region-redirection.html#cross-region-redirection-security-considerations
// [Cross-Region Redirection for Amazon WorkSpaces]: https://docs.aws.amazon.com/workspaces/latest/adminguide/cross-region-redirection.html
func workspaces_DeleteConnectionAlias(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.DeleteConnectionAliasInput{
		// AliasId: *string, // Required
	}

	if len(_workspacesAliasId) > 0 {
		input.AliasId = aws.String(_workspacesAliasId)
	}

	if resp, err := client.DeleteConnectionAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified IP access control group.
// You cannot delete an IP access control group that is associated with a
// directory.
func workspaces_DeleteIpGroup(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.DeleteIpGroupInput{
		// GroupId: *string, // Required
	}

	if len(_workspacesGroupId) > 0 {
		input.GroupId = aws.String(_workspacesGroupId)
	}

	if resp, err := client.DeleteIpGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified tags from the specified WorkSpaces resource.
func workspaces_DeleteTags(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.DeleteTagsInput{
		// ResourceId: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_workspacesResourceId) > 0 {
		input.ResourceId = aws.String(_workspacesResourceId)
	}
	if len(_workspacesTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _workspacesTagKeys...)
	}

	if resp, err := client.DeleteTags(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified WorkSpace bundle. For more information about deleting
// WorkSpace bundles, see [Delete a Custom WorkSpaces Bundle or Image].
//
// [Delete a Custom WorkSpaces Bundle or Image]: https://docs.aws.amazon.com/workspaces/latest/adminguide/delete_bundle.html
func workspaces_DeleteWorkspaceBundle(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.DeleteWorkspaceBundleInput{}

	if len(_workspacesBundleId) > 0 {
		input.BundleId = aws.String(_workspacesBundleId)
	}

	if resp, err := client.DeleteWorkspaceBundle(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified image from your account. To delete an image, you must
// first delete any bundles that are associated with the image and unshare the
// image if it is shared with other accounts.
func workspaces_DeleteWorkspaceImage(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.DeleteWorkspaceImageInput{
		// ImageId: *string, // Required
	}

	if len(_workspacesImageId) > 0 {
		input.ImageId = aws.String(_workspacesImageId)
	}

	if resp, err := client.DeleteWorkspaceImage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deploys associated applications to the specified WorkSpace
func workspaces_DeployWorkspaceApplications(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.DeployWorkspaceApplicationsInput{
		// WorkspaceId: *string, // Required
	}

	if len(_workspacesWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_workspacesWorkspaceId)
	}
	if len(_workspacesForce) > 0 {
		if err := assignInputField(input, "Force", _workspacesForce); err != nil {
			log.Errorf("invalid --force: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeployWorkspaceApplications(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deregisters the specified directory. This operation is asynchronous and returns
// before the WorkSpace directory is deregistered. If any WorkSpaces are registered
// to this directory, you must remove them before you can deregister the directory.
//
// Simple AD and AD Connector are made available to you free of charge to use with
// WorkSpaces. If there are no WorkSpaces being used with your Simple AD or AD
// Connector directory for 30 consecutive days, this directory will be
// automatically deregistered for use with Amazon WorkSpaces, and you will be
// charged for this directory as per the [Directory Service pricing terms].
//
// To delete empty directories, see [Delete the Directory for Your WorkSpaces]. If you delete your Simple AD or AD Connector
// directory, you can always create a new one when you want to start using
// WorkSpaces again.
//
// [Delete the Directory for Your WorkSpaces]: https://docs.aws.amazon.com/workspaces/latest/adminguide/delete-workspaces-directory.html
// [Directory Service pricing terms]: http://aws.amazon.com/directoryservice/pricing/
func workspaces_DeregisterWorkspaceDirectory(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.DeregisterWorkspaceDirectoryInput{
		// DirectoryId: *string, // Required
	}

	if len(_workspacesDirectoryId) > 0 {
		input.DirectoryId = aws.String(_workspacesDirectoryId)
	}

	if resp, err := client.DeregisterWorkspaceDirectory(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list that describes the configuration of Bring Your Own License
// (BYOL) for the specified account.
func workspaces_DescribeAccount(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.DescribeAccountInput{}

	if resp, err := client.DescribeAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list that describes modifications to the configuration of Bring
// Your Own License (BYOL) for the specified account.
func workspaces_DescribeAccountModifications(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.DescribeAccountModificationsInput{}

	if len(_workspacesNextToken) > 0 {
		input.NextToken = aws.String(_workspacesNextToken)
	}

	if resp, err := client.DescribeAccountModifications(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the associations between the application and the specified associated
// resources.
func workspaces_DescribeApplicationAssociations(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.DescribeApplicationAssociationsInput{
		// ApplicationId: *string, // Required
		// AssociatedResourceTypes: []types.ApplicationAssociatedResourceType, // Required
	}

	if len(_workspacesApplicationId) > 0 {
		input.ApplicationId = aws.String(_workspacesApplicationId)
	}
	if len(_workspacesAssociatedResourceTypes) > 0 {
		if err := assignInputField(input, "AssociatedResourceTypes", _workspacesAssociatedResourceTypes); err != nil {
			log.Errorf("invalid --associated-resource-types: %s", err.Error())
			return
		}
	}
	if len(_workspacesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _workspacesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_workspacesNextToken) > 0 {
		input.NextToken = aws.String(_workspacesNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeApplicationAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*workspaces.DescribeApplicationAssociationsOutput
	p := workspaces.NewDescribeApplicationAssociationsPaginator(client, input)
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

// Describes the specified applications by filtering based on their compute types,
// license availability, operating systems, and owners.
func workspaces_DescribeApplications(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.DescribeApplicationsInput{}

	if len(_workspacesApplicationIds) > 0 {
		input.ApplicationIds = append([]string(nil), _workspacesApplicationIds...)
	}
	if len(_workspacesComputeTypeNames) > 0 {
		if err := assignInputField(input, "ComputeTypeNames", _workspacesComputeTypeNames); err != nil {
			log.Errorf("invalid --compute-type-names: %s", err.Error())
			return
		}
	}
	if len(_workspacesLicenseType) > 0 {
		if err := assignInputField(input, "LicenseType", _workspacesLicenseType); err != nil {
			log.Errorf("invalid --license-type: %s", err.Error())
			return
		}
	}
	if len(_workspacesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _workspacesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_workspacesNextToken) > 0 {
		input.NextToken = aws.String(_workspacesNextToken)
	}
	if len(_workspacesOperatingSystemNames) > 0 {
		if err := assignInputField(input, "OperatingSystemNames", _workspacesOperatingSystemNames); err != nil {
			log.Errorf("invalid --operating-system-names: %s", err.Error())
			return
		}
	}
	if len(_workspacesOwner) > 0 {
		input.Owner = aws.String(_workspacesOwner)
	}

	if disablePaginator() {
		if resp, err := client.DescribeApplications(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*workspaces.DescribeApplicationsOutput
	p := workspaces.NewDescribeApplicationsPaginator(client, input)
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

// Describes the associations between the applications and the specified bundle.
func workspaces_DescribeBundleAssociations(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.DescribeBundleAssociationsInput{
		// AssociatedResourceTypes: []types.BundleAssociatedResourceType, // Required
		// BundleId: *string, // Required
	}

	if len(_workspacesAssociatedResourceTypes) > 0 {
		if err := assignInputField(input, "AssociatedResourceTypes", _workspacesAssociatedResourceTypes); err != nil {
			log.Errorf("invalid --associated-resource-types: %s", err.Error())
			return
		}
	}
	if len(_workspacesBundleId) > 0 {
		input.BundleId = aws.String(_workspacesBundleId)
	}

	if resp, err := client.DescribeBundleAssociations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the specified client branding. Client branding allows you to
// customize the log in page of various device types for your users. You can add
// your company logo, the support email address, support link, link to reset
// password, and a custom message for users trying to sign in.
//
// Only device types that have branding information configured will be shown in
// the response.
func workspaces_DescribeClientBranding(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.DescribeClientBrandingInput{
		// ResourceId: *string, // Required
	}

	if len(_workspacesResourceId) > 0 {
		input.ResourceId = aws.String(_workspacesResourceId)
	}

	if resp, err := client.DescribeClientBranding(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list that describes one or more specified Amazon WorkSpaces clients.
func workspaces_DescribeClientProperties(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.DescribeClientPropertiesInput{
		// ResourceIds: []string, // Required
	}

	if len(_workspacesResourceIds) > 0 {
		input.ResourceIds = append([]string(nil), _workspacesResourceIds...)
	}

	if resp, err := client.DescribeClientProperties(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of Amazon Connect client add-ins that have been created.
func workspaces_DescribeConnectClientAddIns(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.DescribeConnectClientAddInsInput{
		// ResourceId: *string, // Required
	}

	if len(_workspacesResourceId) > 0 {
		input.ResourceId = aws.String(_workspacesResourceId)
	}
	if len(_workspacesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _workspacesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_workspacesNextToken) > 0 {
		input.NextToken = aws.String(_workspacesNextToken)
	}

	if resp, err := client.DescribeConnectClientAddIns(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the permissions that the owner of a connection alias has granted to
// another Amazon Web Services account for the specified connection alias. For more
// information, see [Cross-Region Redirection for Amazon WorkSpaces].
//
// [Cross-Region Redirection for Amazon WorkSpaces]: https://docs.aws.amazon.com/workspaces/latest/adminguide/cross-region-redirection.html
func workspaces_DescribeConnectionAliasPermissions(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.DescribeConnectionAliasPermissionsInput{
		// AliasId: *string, // Required
	}

	if len(_workspacesAliasId) > 0 {
		input.AliasId = aws.String(_workspacesAliasId)
	}
	if len(_workspacesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _workspacesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_workspacesNextToken) > 0 {
		input.NextToken = aws.String(_workspacesNextToken)
	}

	if resp, err := client.DescribeConnectionAliasPermissions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list that describes the connection aliases used for cross-Region
// redirection. For more information, see [Cross-Region Redirection for Amazon WorkSpaces].
//
// [Cross-Region Redirection for Amazon WorkSpaces]: https://docs.aws.amazon.com/workspaces/latest/adminguide/cross-region-redirection.html
func workspaces_DescribeConnectionAliases(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.DescribeConnectionAliasesInput{}

	if len(_workspacesAliasIds) > 0 {
		input.AliasIds = append([]string(nil), _workspacesAliasIds...)
	}
	if len(_workspacesLimit) > 0 {
		if err := assignInputField(input, "Limit", _workspacesLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_workspacesNextToken) > 0 {
		input.NextToken = aws.String(_workspacesNextToken)
	}
	if len(_workspacesResourceId) > 0 {
		input.ResourceId = aws.String(_workspacesResourceId)
	}

	if resp, err := client.DescribeConnectionAliases(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a WorkSpace BYOL image being imported via
// ImportCustomWorkspaceImage.
func workspaces_DescribeCustomWorkspaceImageImport(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.DescribeCustomWorkspaceImageImportInput{
		// ImageId: *string, // Required
	}

	if len(_workspacesImageId) > 0 {
		input.ImageId = aws.String(_workspacesImageId)
	}

	if resp, err := client.DescribeCustomWorkspaceImageImport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the associations between the applications and the specified image.
func workspaces_DescribeImageAssociations(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.DescribeImageAssociationsInput{
		// AssociatedResourceTypes: []types.ImageAssociatedResourceType, // Required
		// ImageId: *string, // Required
	}

	if len(_workspacesAssociatedResourceTypes) > 0 {
		if err := assignInputField(input, "AssociatedResourceTypes", _workspacesAssociatedResourceTypes); err != nil {
			log.Errorf("invalid --associated-resource-types: %s", err.Error())
			return
		}
	}
	if len(_workspacesImageId) > 0 {
		input.ImageId = aws.String(_workspacesImageId)
	}

	if resp, err := client.DescribeImageAssociations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes one or more of your IP access control groups.
func workspaces_DescribeIpGroups(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.DescribeIpGroupsInput{}

	if len(_workspacesGroupIds) > 0 {
		input.GroupIds = append([]string(nil), _workspacesGroupIds...)
	}
	if len(_workspacesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _workspacesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_workspacesNextToken) > 0 {
		input.NextToken = aws.String(_workspacesNextToken)
	}

	if resp, err := client.DescribeIpGroups(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the specified tags for the specified WorkSpaces resource.
func workspaces_DescribeTags(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.DescribeTagsInput{
		// ResourceId: *string, // Required
	}

	if len(_workspacesResourceId) > 0 {
		input.ResourceId = aws.String(_workspacesResourceId)
	}

	if resp, err := client.DescribeTags(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the associations betweens applications and the specified WorkSpace.
func workspaces_DescribeWorkspaceAssociations(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.DescribeWorkspaceAssociationsInput{
		// AssociatedResourceTypes: []types.WorkSpaceAssociatedResourceType, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_workspacesAssociatedResourceTypes) > 0 {
		if err := assignInputField(input, "AssociatedResourceTypes", _workspacesAssociatedResourceTypes); err != nil {
			log.Errorf("invalid --associated-resource-types: %s", err.Error())
			return
		}
	}
	if len(_workspacesWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_workspacesWorkspaceId)
	}

	if resp, err := client.DescribeWorkspaceAssociations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list that describes the available WorkSpace bundles.
// You can filter the results using either bundle ID or owner, but not both.
func workspaces_DescribeWorkspaceBundles(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.DescribeWorkspaceBundlesInput{}

	if len(_workspacesBundleIds) > 0 {
		input.BundleIds = append([]string(nil), _workspacesBundleIds...)
	}
	if len(_workspacesNextToken) > 0 {
		input.NextToken = aws.String(_workspacesNextToken)
	}
	if len(_workspacesOwner) > 0 {
		input.Owner = aws.String(_workspacesOwner)
	}

	if disablePaginator() {
		if resp, err := client.DescribeWorkspaceBundles(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*workspaces.DescribeWorkspaceBundlesOutput
	p := workspaces.NewDescribeWorkspaceBundlesPaginator(client, input)
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

// Describes the available directories that are registered with Amazon WorkSpaces.
func workspaces_DescribeWorkspaceDirectories(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.DescribeWorkspaceDirectoriesInput{}

	if len(_workspacesDirectoryIds) > 0 {
		input.DirectoryIds = append([]string(nil), _workspacesDirectoryIds...)
	}
	if len(_workspacesFilters) > 0 {
		if err := assignInputField(input, "Filters", _workspacesFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_workspacesLimit) > 0 {
		if err := assignInputField(input, "Limit", _workspacesLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_workspacesNextToken) > 0 {
		input.NextToken = aws.String(_workspacesNextToken)
	}
	if len(_workspacesWorkspaceDirectoryNames) > 0 {
		input.WorkspaceDirectoryNames = append([]string(nil), _workspacesWorkspaceDirectoryNames...)
	}

	if disablePaginator() {
		if resp, err := client.DescribeWorkspaceDirectories(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*workspaces.DescribeWorkspaceDirectoriesOutput
	p := workspaces.NewDescribeWorkspaceDirectoriesPaginator(client, input)
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

// Describes the permissions that the owner of an image has granted to other
// Amazon Web Services accounts for an image.
func workspaces_DescribeWorkspaceImagePermissions(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.DescribeWorkspaceImagePermissionsInput{
		// ImageId: *string, // Required
	}

	if len(_workspacesImageId) > 0 {
		input.ImageId = aws.String(_workspacesImageId)
	}
	if len(_workspacesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _workspacesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_workspacesNextToken) > 0 {
		input.NextToken = aws.String(_workspacesNextToken)
	}

	if resp, err := client.DescribeWorkspaceImagePermissions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list that describes one or more specified images, if the image
// identifiers are provided. Otherwise, all images in the account are described.
func workspaces_DescribeWorkspaceImages(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.DescribeWorkspaceImagesInput{}

	if len(_workspacesImageIds) > 0 {
		input.ImageIds = append([]string(nil), _workspacesImageIds...)
	}
	if len(_workspacesImageType) > 0 {
		if err := assignInputField(input, "ImageType", _workspacesImageType); err != nil {
			log.Errorf("invalid --image-type: %s", err.Error())
			return
		}
	}
	if len(_workspacesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _workspacesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_workspacesNextToken) > 0 {
		input.NextToken = aws.String(_workspacesNextToken)
	}

	if resp, err := client.DescribeWorkspaceImages(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the snapshots for the specified WorkSpace.
func workspaces_DescribeWorkspaceSnapshots(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.DescribeWorkspaceSnapshotsInput{
		// WorkspaceId: *string, // Required
	}

	if len(_workspacesWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_workspacesWorkspaceId)
	}

	if resp, err := client.DescribeWorkspaceSnapshots(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the specified WorkSpaces.
// You can filter the results by using the bundle identifier, directory
// identifier, or owner, but you can specify only one filter at a time.
func workspaces_DescribeWorkspaces(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.DescribeWorkspacesInput{}

	if len(_workspacesBundleId) > 0 {
		input.BundleId = aws.String(_workspacesBundleId)
	}
	if len(_workspacesDirectoryId) > 0 {
		input.DirectoryId = aws.String(_workspacesDirectoryId)
	}
	if len(_workspacesLimit) > 0 {
		if err := assignInputField(input, "Limit", _workspacesLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_workspacesNextToken) > 0 {
		input.NextToken = aws.String(_workspacesNextToken)
	}
	if len(_workspacesUserName) > 0 {
		input.UserName = aws.String(_workspacesUserName)
	}
	if len(_workspacesWorkspaceIds) > 0 {
		input.WorkspaceIds = append([]string(nil), _workspacesWorkspaceIds...)
	}
	if len(_workspacesWorkspaceName) > 0 {
		input.WorkspaceName = aws.String(_workspacesWorkspaceName)
	}

	if disablePaginator() {
		if resp, err := client.DescribeWorkspaces(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*workspaces.DescribeWorkspacesOutput
	p := workspaces.NewDescribeWorkspacesPaginator(client, input)
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

// Describes the connection status of the specified WorkSpaces.
func workspaces_DescribeWorkspacesConnectionStatus(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.DescribeWorkspacesConnectionStatusInput{}

	if len(_workspacesNextToken) > 0 {
		input.NextToken = aws.String(_workspacesNextToken)
	}
	if len(_workspacesWorkspaceIds) > 0 {
		input.WorkspaceIds = append([]string(nil), _workspacesWorkspaceIds...)
	}

	if resp, err := client.DescribeWorkspacesConnectionStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list that describes the streaming sessions for a specified pool.
func workspaces_DescribeWorkspacesPoolSessions(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.DescribeWorkspacesPoolSessionsInput{
		// PoolId: *string, // Required
	}

	if len(_workspacesPoolId) > 0 {
		input.PoolId = aws.String(_workspacesPoolId)
	}
	if len(_workspacesLimit) > 0 {
		if err := assignInputField(input, "Limit", _workspacesLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_workspacesNextToken) > 0 {
		input.NextToken = aws.String(_workspacesNextToken)
	}
	if len(_workspacesUserId) > 0 {
		input.UserId = aws.String(_workspacesUserId)
	}

	if resp, err := client.DescribeWorkspacesPoolSessions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the specified WorkSpaces Pools.
func workspaces_DescribeWorkspacesPools(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.DescribeWorkspacesPoolsInput{}

	if len(_workspacesFilters) > 0 {
		if err := assignInputField(input, "Filters", _workspacesFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_workspacesLimit) > 0 {
		if err := assignInputField(input, "Limit", _workspacesLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_workspacesNextToken) > 0 {
		input.NextToken = aws.String(_workspacesNextToken)
	}
	if len(_workspacesPoolIds) > 0 {
		input.PoolIds = append([]string(nil), _workspacesPoolIds...)
	}

	if resp, err := client.DescribeWorkspacesPools(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a connection alias from a directory. Disassociating a connection
// alias disables cross-Region redirection between two directories in different
// Regions. For more information, see [Cross-Region Redirection for Amazon WorkSpaces].
//
// Before performing this operation, call [DescribeConnectionAliases] to make sure that the current state of
// the connection alias is CREATED .
//
// [DescribeConnectionAliases]: https://docs.aws.amazon.com/workspaces/latest/api/API_DescribeConnectionAliases.html
// [Cross-Region Redirection for Amazon WorkSpaces]: https://docs.aws.amazon.com/workspaces/latest/adminguide/cross-region-redirection.html
func workspaces_DisassociateConnectionAlias(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.DisassociateConnectionAliasInput{
		// AliasId: *string, // Required
	}

	if len(_workspacesAliasId) > 0 {
		input.AliasId = aws.String(_workspacesAliasId)
	}

	if resp, err := client.DisassociateConnectionAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates the specified IP access control group from the specified
// directory.
func workspaces_DisassociateIpGroups(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.DisassociateIpGroupsInput{
		// DirectoryId: *string, // Required
		// GroupIds: []string, // Required
	}

	if len(_workspacesDirectoryId) > 0 {
		input.DirectoryId = aws.String(_workspacesDirectoryId)
	}
	if len(_workspacesGroupIds) > 0 {
		input.GroupIds = append([]string(nil), _workspacesGroupIds...)
	}

	if resp, err := client.DisassociateIpGroups(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates the specified application from a WorkSpace.
func workspaces_DisassociateWorkspaceApplication(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.DisassociateWorkspaceApplicationInput{
		// ApplicationId: *string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_workspacesApplicationId) > 0 {
		input.ApplicationId = aws.String(_workspacesApplicationId)
	}
	if len(_workspacesWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_workspacesWorkspaceId)
	}

	if resp, err := client.DisassociateWorkspaceApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves account link information.
func workspaces_GetAccountLink(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.GetAccountLinkInput{}

	if len(_workspacesLinkId) > 0 {
		input.LinkId = aws.String(_workspacesLinkId)
	}
	if len(_workspacesLinkedAccountId) > 0 {
		input.LinkedAccountId = aws.String(_workspacesLinkedAccountId)
	}

	if resp, err := client.GetAccountLink(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Imports client branding. Client branding allows you to customize your
// WorkSpace's client login portal. You can tailor your login portal company logo,
// the support email address, support link, link to reset password, and a custom
// message for users trying to sign in.
//
// After you import client branding, the default branding experience for the
// specified platform type is replaced with the imported experience
//
// - You must specify at least one platform type when importing client branding.
//
// - You can import up to 6 MB of data with each request. If your request
// exceeds this limit, you can import client branding for different platform types
// using separate requests.
//
// - In each platform type, the SupportEmail and SupportLink parameters are
// mutually exclusive. You can specify only one parameter for each platform type,
// but not both.
//
// - Imported data can take up to a minute to appear in the WorkSpaces client.
func workspaces_ImportClientBranding(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.ImportClientBrandingInput{
		// ResourceId: *string, // Required
	}

	if len(_workspacesResourceId) > 0 {
		input.ResourceId = aws.String(_workspacesResourceId)
	}
	if len(_workspacesDeviceTypeAndroid) > 0 {
		if err := assignInputField(input, "DeviceTypeAndroid", _workspacesDeviceTypeAndroid); err != nil {
			log.Errorf("invalid --device-type-android: %s", err.Error())
			return
		}
	}
	if len(_workspacesDeviceTypeIos) > 0 {
		if err := assignInputField(input, "DeviceTypeIos", _workspacesDeviceTypeIos); err != nil {
			log.Errorf("invalid --device-type-ios: %s", err.Error())
			return
		}
	}
	if len(_workspacesDeviceTypeLinux) > 0 {
		if err := assignInputField(input, "DeviceTypeLinux", _workspacesDeviceTypeLinux); err != nil {
			log.Errorf("invalid --device-type-linux: %s", err.Error())
			return
		}
	}
	if len(_workspacesDeviceTypeOsx) > 0 {
		if err := assignInputField(input, "DeviceTypeOsx", _workspacesDeviceTypeOsx); err != nil {
			log.Errorf("invalid --device-type-osx: %s", err.Error())
			return
		}
	}
	if len(_workspacesDeviceTypeWeb) > 0 {
		if err := assignInputField(input, "DeviceTypeWeb", _workspacesDeviceTypeWeb); err != nil {
			log.Errorf("invalid --device-type-web: %s", err.Error())
			return
		}
	}
	if len(_workspacesDeviceTypeWindows) > 0 {
		if err := assignInputField(input, "DeviceTypeWindows", _workspacesDeviceTypeWindows); err != nil {
			log.Errorf("invalid --device-type-windows: %s", err.Error())
			return
		}
	}

	if resp, err := client.ImportClientBranding(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Imports the specified Windows 10 or 11 Bring Your Own License (BYOL) image into
// Amazon WorkSpaces using EC2 Image Builder. The image must be an already licensed
// image that is in your Amazon Web Services account, and you must own the image.
// For more information about creating BYOL images, see [Bring Your Own Windows Desktop Licenses].
//
// [Bring Your Own Windows Desktop Licenses]: https://docs.aws.amazon.com/workspaces/latest/adminguide/byol-windows-images.html
func workspaces_ImportCustomWorkspaceImage(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.ImportCustomWorkspaceImageInput{
		// ComputeType: types.ImageComputeType, // Required
		// ImageDescription: *string, // Required
		// ImageName: *string, // Required
		// ImageSource: types.ImageSourceIdentifier, // Required
		// InfrastructureConfigurationArn: *string, // Required
		// OsVersion: types.OSVersion, // Required
		// Platform: types.Platform, // Required
		// Protocol: types.CustomImageProtocol, // Required
	}

	if len(_workspacesComputeType) > 0 {
		if err := assignInputField(input, "ComputeType", _workspacesComputeType); err != nil {
			log.Errorf("invalid --compute-type: %s", err.Error())
			return
		}
	}
	if len(_workspacesImageDescription) > 0 {
		input.ImageDescription = aws.String(_workspacesImageDescription)
	}
	if len(_workspacesImageName) > 0 {
		input.ImageName = aws.String(_workspacesImageName)
	}
	if len(_workspacesImageSource) > 0 {
		if err := assignInputField(input, "ImageSource", _workspacesImageSource); err != nil {
			log.Errorf("invalid --image-source: %s", err.Error())
			return
		}
	}
	if len(_workspacesInfrastructureConfigurationArn) > 0 {
		input.InfrastructureConfigurationArn = aws.String(_workspacesInfrastructureConfigurationArn)
	}
	if len(_workspacesOsVersion) > 0 {
		if err := assignInputField(input, "OsVersion", _workspacesOsVersion); err != nil {
			log.Errorf("invalid --os-version: %s", err.Error())
			return
		}
	}
	if len(_workspacesPlatform) > 0 {
		if err := assignInputField(input, "Platform", _workspacesPlatform); err != nil {
			log.Errorf("invalid --platform: %s", err.Error())
			return
		}
	}
	if len(_workspacesProtocol) > 0 {
		if err := assignInputField(input, "Protocol", _workspacesProtocol); err != nil {
			log.Errorf("invalid --protocol: %s", err.Error())
			return
		}
	}
	if len(_workspacesTags) > 0 {
		if err := assignInputField(input, "Tags", _workspacesTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.ImportCustomWorkspaceImage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Imports the specified Windows 10 or 11 Bring Your Own License (BYOL) image into
// Amazon WorkSpaces. The image must be an already licensed Amazon EC2 image that
// is in your Amazon Web Services account, and you must own the image. For more
// information about creating BYOL images, see [Bring Your Own Windows Desktop Licenses].
//
// [Bring Your Own Windows Desktop Licenses]: https://docs.aws.amazon.com/workspaces/latest/adminguide/byol-windows-images.html
func workspaces_ImportWorkspaceImage(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.ImportWorkspaceImageInput{
		// Ec2ImageId: *string, // Required
		// ImageDescription: *string, // Required
		// ImageName: *string, // Required
		// IngestionProcess: types.WorkspaceImageIngestionProcess, // Required
	}

	if len(_workspacesEc2ImageId) > 0 {
		input.Ec2ImageId = aws.String(_workspacesEc2ImageId)
	}
	if len(_workspacesImageDescription) > 0 {
		input.ImageDescription = aws.String(_workspacesImageDescription)
	}
	if len(_workspacesImageName) > 0 {
		input.ImageName = aws.String(_workspacesImageName)
	}
	if len(_workspacesIngestionProcess) > 0 {
		if err := assignInputField(input, "IngestionProcess", _workspacesIngestionProcess); err != nil {
			log.Errorf("invalid --ingestion-process: %s", err.Error())
			return
		}
	}
	if len(_workspacesApplications) > 0 {
		if err := assignInputField(input, "Applications", _workspacesApplications); err != nil {
			log.Errorf("invalid --applications: %s", err.Error())
			return
		}
	}
	if len(_workspacesTags) > 0 {
		if err := assignInputField(input, "Tags", _workspacesTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.ImportWorkspaceImage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all account links.
func workspaces_ListAccountLinks(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.ListAccountLinksInput{}

	if len(_workspacesLinkStatusFilter) > 0 {
		if err := assignInputField(input, "LinkStatusFilter", _workspacesLinkStatusFilter); err != nil {
			log.Errorf("invalid --link-status-filter: %s", err.Error())
			return
		}
	}
	if len(_workspacesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _workspacesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_workspacesNextToken) > 0 {
		input.NextToken = aws.String(_workspacesNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAccountLinks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*workspaces.ListAccountLinksOutput
	p := workspaces.NewListAccountLinksPaginator(client, input)
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

// Retrieves a list of IP address ranges, specified as IPv4 CIDR blocks, that you
// can use for the network management interface when you enable Bring Your Own
// License (BYOL).
//
// This operation can be run only by Amazon Web Services accounts that are enabled
// for BYOL. If your account isn't enabled for BYOL, you'll receive an
// AccessDeniedException error.
//
// The management network interface is connected to a secure Amazon WorkSpaces
// management network. It is used for interactive streaming of the WorkSpace
// desktop to Amazon WorkSpaces clients, and to allow Amazon WorkSpaces to manage
// the WorkSpace.
func workspaces_ListAvailableManagementCidrRanges(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.ListAvailableManagementCidrRangesInput{
		// ManagementCidrRangeConstraint: *string, // Required
	}

	if len(_workspacesManagementCidrRangeConstraint) > 0 {
		input.ManagementCidrRangeConstraint = aws.String(_workspacesManagementCidrRangeConstraint)
	}
	if len(_workspacesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _workspacesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_workspacesNextToken) > 0 {
		input.NextToken = aws.String(_workspacesNextToken)
	}

	if resp, err := client.ListAvailableManagementCidrRanges(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Migrates a WorkSpace from one operating system or bundle type to another, while
// retaining the data on the user volume.
//
// The migration process recreates the WorkSpace by using a new root volume from
// the target bundle image and the user volume from the last available snapshot of
// the original WorkSpace. During migration, the original D:\Users\%USERNAME% user
// profile folder is renamed to D:\Users\%USERNAME%MMddyyTHHmmss%.NotMigrated . A
// new D:\Users\%USERNAME%\ folder is generated by the new OS. Certain files in
// the old user profile are moved to the new user profile.
//
// For available migration scenarios, details about what happens during migration,
// and best practices, see [Migrate a WorkSpace].
//
// [Migrate a WorkSpace]: https://docs.aws.amazon.com/workspaces/latest/adminguide/migrate-workspaces.html
func workspaces_MigrateWorkspace(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.MigrateWorkspaceInput{
		// BundleId: *string, // Required
		// SourceWorkspaceId: *string, // Required
	}

	if len(_workspacesBundleId) > 0 {
		input.BundleId = aws.String(_workspacesBundleId)
	}
	if len(_workspacesSourceWorkspaceId) > 0 {
		input.SourceWorkspaceId = aws.String(_workspacesSourceWorkspaceId)
	}

	if resp, err := client.MigrateWorkspace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the configuration of Bring Your Own License (BYOL) for the specified
// account.
func workspaces_ModifyAccount(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.ModifyAccountInput{}

	if len(_workspacesDedicatedTenancyManagementCidrRange) > 0 {
		input.DedicatedTenancyManagementCidrRange = aws.String(_workspacesDedicatedTenancyManagementCidrRange)
	}
	if len(_workspacesDedicatedTenancySupport) > 0 {
		if err := assignInputField(input, "DedicatedTenancySupport", _workspacesDedicatedTenancySupport); err != nil {
			log.Errorf("invalid --dedicated-tenancy-support: %s", err.Error())
			return
		}
	}

	if resp, err := client.ModifyAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the properties of the certificate-based authentication you want to use
// with your WorkSpaces.
func workspaces_ModifyCertificateBasedAuthProperties(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.ModifyCertificateBasedAuthPropertiesInput{
		// ResourceId: *string, // Required
	}

	if len(_workspacesResourceId) > 0 {
		input.ResourceId = aws.String(_workspacesResourceId)
	}
	if len(_workspacesCertificateBasedAuthProperties) > 0 {
		if err := assignInputField(input, "CertificateBasedAuthProperties", _workspacesCertificateBasedAuthProperties); err != nil {
			log.Errorf("invalid --certificate-based-auth-properties: %s", err.Error())
			return
		}
	}
	if len(_workspacesPropertiesToDelete) > 0 {
		if err := assignInputField(input, "PropertiesToDelete", _workspacesPropertiesToDelete); err != nil {
			log.Errorf("invalid --properties-to-delete: %s", err.Error())
			return
		}
	}

	if resp, err := client.ModifyCertificateBasedAuthProperties(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the properties of the specified Amazon WorkSpaces clients.
func workspaces_ModifyClientProperties(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.ModifyClientPropertiesInput{
		// ClientProperties: *types.ClientProperties, // Required
		// ResourceId: *string, // Required
	}

	if len(_workspacesClientProperties) > 0 {
		if err := assignInputField(input, "ClientProperties", _workspacesClientProperties); err != nil {
			log.Errorf("invalid --client-properties: %s", err.Error())
			return
		}
	}
	if len(_workspacesResourceId) > 0 {
		input.ResourceId = aws.String(_workspacesResourceId)
	}

	if resp, err := client.ModifyClientProperties(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the endpoint encryption mode that allows you to configure the
// specified directory between Standard TLS and FIPS 140-2 validated mode.
func workspaces_ModifyEndpointEncryptionMode(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.ModifyEndpointEncryptionModeInput{
		// DirectoryId: *string, // Required
		// EndpointEncryptionMode: types.EndpointEncryptionMode, // Required
	}

	if len(_workspacesDirectoryId) > 0 {
		input.DirectoryId = aws.String(_workspacesDirectoryId)
	}
	if len(_workspacesEndpointEncryptionMode) > 0 {
		if err := assignInputField(input, "EndpointEncryptionMode", _workspacesEndpointEncryptionMode); err != nil {
			log.Errorf("invalid --endpoint-encryption-mode: %s", err.Error())
			return
		}
	}

	if resp, err := client.ModifyEndpointEncryptionMode(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies multiple properties related to SAML 2.0 authentication, including the
// enablement status, user access URL, and relay state parameter name that are used
// for configuring federation with an SAML 2.0 identity provider.
func workspaces_ModifySamlProperties(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.ModifySamlPropertiesInput{
		// ResourceId: *string, // Required
	}

	if len(_workspacesResourceId) > 0 {
		input.ResourceId = aws.String(_workspacesResourceId)
	}
	if len(_workspacesPropertiesToDelete) > 0 {
		if err := assignInputField(input, "PropertiesToDelete", _workspacesPropertiesToDelete); err != nil {
			log.Errorf("invalid --properties-to-delete: %s", err.Error())
			return
		}
	}
	if len(_workspacesSamlProperties) > 0 {
		if err := assignInputField(input, "SamlProperties", _workspacesSamlProperties); err != nil {
			log.Errorf("invalid --saml-properties: %s", err.Error())
			return
		}
	}

	if resp, err := client.ModifySamlProperties(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the self-service WorkSpace management capabilities for your users. For
// more information, see [Enable Self-Service WorkSpace Management Capabilities for Your Users].
//
// [Enable Self-Service WorkSpace Management Capabilities for Your Users]: https://docs.aws.amazon.com/workspaces/latest/adminguide/enable-user-self-service-workspace-management.html
func workspaces_ModifySelfservicePermissions(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.ModifySelfservicePermissionsInput{
		// ResourceId: *string, // Required
		// SelfservicePermissions: *types.SelfservicePermissions, // Required
	}

	if len(_workspacesResourceId) > 0 {
		input.ResourceId = aws.String(_workspacesResourceId)
	}
	if len(_workspacesSelfservicePermissions) > 0 {
		if err := assignInputField(input, "SelfservicePermissions", _workspacesSelfservicePermissions); err != nil {
			log.Errorf("invalid --selfservice-permissions: %s", err.Error())
			return
		}
	}

	if resp, err := client.ModifySelfservicePermissions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the specified streaming properties.
func workspaces_ModifyStreamingProperties(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.ModifyStreamingPropertiesInput{
		// ResourceId: *string, // Required
	}

	if len(_workspacesResourceId) > 0 {
		input.ResourceId = aws.String(_workspacesResourceId)
	}
	if len(_workspacesStreamingProperties) > 0 {
		if err := assignInputField(input, "StreamingProperties", _workspacesStreamingProperties); err != nil {
			log.Errorf("invalid --streaming-properties: %s", err.Error())
			return
		}
	}

	if resp, err := client.ModifyStreamingProperties(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Specifies which devices and operating systems users can use to access their
// WorkSpaces. For more information, see [Control Device Access].
//
// [Control Device Access]: https://docs.aws.amazon.com/workspaces/latest/adminguide/update-directory-details.html#control-device-access
func workspaces_ModifyWorkspaceAccessProperties(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.ModifyWorkspaceAccessPropertiesInput{
		// ResourceId: *string, // Required
		// WorkspaceAccessProperties: *types.WorkspaceAccessProperties, // Required
	}

	if len(_workspacesResourceId) > 0 {
		input.ResourceId = aws.String(_workspacesResourceId)
	}
	if len(_workspacesWorkspaceAccessProperties) > 0 {
		if err := assignInputField(input, "WorkspaceAccessProperties", _workspacesWorkspaceAccessProperties); err != nil {
			log.Errorf("invalid --workspace-access-properties: %s", err.Error())
			return
		}
	}

	if resp, err := client.ModifyWorkspaceAccessProperties(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modify the default properties used to create WorkSpaces.
func workspaces_ModifyWorkspaceCreationProperties(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.ModifyWorkspaceCreationPropertiesInput{
		// ResourceId: *string, // Required
		// WorkspaceCreationProperties: *types.WorkspaceCreationProperties, // Required
	}

	if len(_workspacesResourceId) > 0 {
		input.ResourceId = aws.String(_workspacesResourceId)
	}
	if len(_workspacesWorkspaceCreationProperties) > 0 {
		if err := assignInputField(input, "WorkspaceCreationProperties", _workspacesWorkspaceCreationProperties); err != nil {
			log.Errorf("invalid --workspace-creation-properties: %s", err.Error())
			return
		}
	}

	if resp, err := client.ModifyWorkspaceCreationProperties(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the specified WorkSpace properties. For important information about
// how to modify the size of the root and user volumes, see [Modify a WorkSpace].
//
// The MANUAL running mode value is only supported by Amazon WorkSpaces Core.
// Contact your account team to be allow-listed to use this value. For more
// information, see [Amazon WorkSpaces Core].
//
// [Amazon WorkSpaces Core]: http://aws.amazon.com/workspaces/core/
// [Modify a WorkSpace]: https://docs.aws.amazon.com/workspaces/latest/adminguide/modify-workspaces.html
func workspaces_ModifyWorkspaceProperties(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.ModifyWorkspacePropertiesInput{
		// WorkspaceId: *string, // Required
	}

	if len(_workspacesWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_workspacesWorkspaceId)
	}
	if len(_workspacesDataReplication) > 0 {
		if err := assignInputField(input, "DataReplication", _workspacesDataReplication); err != nil {
			log.Errorf("invalid --data-replication: %s", err.Error())
			return
		}
	}
	if len(_workspacesWorkspaceProperties) > 0 {
		if err := assignInputField(input, "WorkspaceProperties", _workspacesWorkspaceProperties); err != nil {
			log.Errorf("invalid --workspace-properties: %s", err.Error())
			return
		}
	}

	if resp, err := client.ModifyWorkspaceProperties(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the state of the specified WorkSpace.
// To maintain a WorkSpace without being interrupted, set the WorkSpace state to
// ADMIN_MAINTENANCE . WorkSpaces in this state do not respond to requests to
// reboot, stop, start, rebuild, or restore. An AutoStop WorkSpace in this state is
// not stopped. Users cannot log into a WorkSpace in the ADMIN_MAINTENANCE state.
func workspaces_ModifyWorkspaceState(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.ModifyWorkspaceStateInput{
		// WorkspaceId: *string, // Required
		// WorkspaceState: types.TargetWorkspaceState, // Required
	}

	if len(_workspacesWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_workspacesWorkspaceId)
	}
	if len(_workspacesWorkspaceState) > 0 {
		if err := assignInputField(input, "WorkspaceState", _workspacesWorkspaceState); err != nil {
			log.Errorf("invalid --workspace-state: %s", err.Error())
			return
		}
	}

	if resp, err := client.ModifyWorkspaceState(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Reboots the specified WorkSpaces.
// You cannot reboot a WorkSpace unless its state is AVAILABLE , UNHEALTHY , or
// REBOOTING . Reboot a WorkSpace in the REBOOTING state only if your WorkSpace
// has been stuck in the REBOOTING state for over 20 minutes.
//
// This operation is asynchronous and returns before the WorkSpaces have rebooted.
func workspaces_RebootWorkspaces(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.RebootWorkspacesInput{
		// RebootWorkspaceRequests: []types.RebootRequest, // Required
	}

	if len(_workspacesRebootWorkspaceRequests) > 0 {
		if err := assignInputField(input, "RebootWorkspaceRequests", _workspacesRebootWorkspaceRequests); err != nil {
			log.Errorf("invalid --reboot-workspace-requests: %s", err.Error())
			return
		}
	}

	if resp, err := client.RebootWorkspaces(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Rebuilds the specified WorkSpace.
// You cannot rebuild a WorkSpace unless its state is AVAILABLE , ERROR , UNHEALTHY
// , STOPPED , or REBOOTING .
//
// Rebuilding a WorkSpace is a potentially destructive action that can result in
// the loss of data. For more information, see [Rebuild a WorkSpace].
//
// This operation is asynchronous and returns before the WorkSpaces have been
// completely rebuilt.
//
// [Rebuild a WorkSpace]: https://docs.aws.amazon.com/workspaces/latest/adminguide/reset-workspace.html
func workspaces_RebuildWorkspaces(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.RebuildWorkspacesInput{
		// RebuildWorkspaceRequests: []types.RebuildRequest, // Required
	}

	if len(_workspacesRebuildWorkspaceRequests) > 0 {
		if err := assignInputField(input, "RebuildWorkspaceRequests", _workspacesRebuildWorkspaceRequests); err != nil {
			log.Errorf("invalid --rebuild-workspace-requests: %s", err.Error())
			return
		}
	}

	if resp, err := client.RebuildWorkspaces(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Registers the specified directory. This operation is asynchronous and returns
// before the WorkSpace directory is registered. If this is the first time you are
// registering a directory, you will need to create the workspaces_DefaultRole role
// before you can register a directory. For more information, see [Creating the workspaces_DefaultRole Role].
//
// [Creating the workspaces_DefaultRole Role]: https://docs.aws.amazon.com/workspaces/latest/adminguide/workspaces-access-control.html#create-default-role
func workspaces_RegisterWorkspaceDirectory(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.RegisterWorkspaceDirectoryInput{}

	if len(_workspacesActiveDirectoryConfig) > 0 {
		if err := assignInputField(input, "ActiveDirectoryConfig", _workspacesActiveDirectoryConfig); err != nil {
			log.Errorf("invalid --active-directory-config: %s", err.Error())
			return
		}
	}
	if len(_workspacesDirectoryId) > 0 {
		input.DirectoryId = aws.String(_workspacesDirectoryId)
	}
	if len(_workspacesEnableSelfService) > 0 {
		if err := assignInputField(input, "EnableSelfService", _workspacesEnableSelfService); err != nil {
			log.Errorf("invalid --enable-self-service: %s", err.Error())
			return
		}
	}
	if len(_workspacesIdcInstanceArn) > 0 {
		input.IdcInstanceArn = aws.String(_workspacesIdcInstanceArn)
	}
	if len(_workspacesMicrosoftEntraConfig) > 0 {
		if err := assignInputField(input, "MicrosoftEntraConfig", _workspacesMicrosoftEntraConfig); err != nil {
			log.Errorf("invalid --microsoft-entra-config: %s", err.Error())
			return
		}
	}
	if len(_workspacesSubnetIds) > 0 {
		input.SubnetIds = append([]string(nil), _workspacesSubnetIds...)
	}
	if len(_workspacesTags) > 0 {
		if err := assignInputField(input, "Tags", _workspacesTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_workspacesTenancy) > 0 {
		if err := assignInputField(input, "Tenancy", _workspacesTenancy); err != nil {
			log.Errorf("invalid --tenancy: %s", err.Error())
			return
		}
	}
	if len(_workspacesUserIdentityType) > 0 {
		if err := assignInputField(input, "UserIdentityType", _workspacesUserIdentityType); err != nil {
			log.Errorf("invalid --user-identity-type: %s", err.Error())
			return
		}
	}
	if len(_workspacesWorkspaceDirectoryDescription) > 0 {
		input.WorkspaceDirectoryDescription = aws.String(_workspacesWorkspaceDirectoryDescription)
	}
	if len(_workspacesWorkspaceDirectoryName) > 0 {
		input.WorkspaceDirectoryName = aws.String(_workspacesWorkspaceDirectoryName)
	}
	if len(_workspacesWorkspaceType) > 0 {
		if err := assignInputField(input, "WorkspaceType", _workspacesWorkspaceType); err != nil {
			log.Errorf("invalid --workspace-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.RegisterWorkspaceDirectory(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Rejects the account link invitation.
func workspaces_RejectAccountLinkInvitation(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.RejectAccountLinkInvitationInput{
		// LinkId: *string, // Required
	}

	if len(_workspacesLinkId) > 0 {
		input.LinkId = aws.String(_workspacesLinkId)
	}
	if len(_workspacesClientToken) > 0 {
		input.ClientToken = aws.String(_workspacesClientToken)
	}

	if resp, err := client.RejectAccountLinkInvitation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Restores the specified WorkSpace to its last known healthy state.
// You cannot restore a WorkSpace unless its state is  AVAILABLE , ERROR ,
// UNHEALTHY , or STOPPED .
//
// Restoring a WorkSpace is a potentially destructive action that can result in
// the loss of data. For more information, see [Restore a WorkSpace].
//
// This operation is asynchronous and returns before the WorkSpace is completely
// restored.
//
// [Restore a WorkSpace]: https://docs.aws.amazon.com/workspaces/latest/adminguide/restore-workspace.html
func workspaces_RestoreWorkspace(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.RestoreWorkspaceInput{
		// WorkspaceId: *string, // Required
	}

	if len(_workspacesWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_workspacesWorkspaceId)
	}

	if resp, err := client.RestoreWorkspace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes one or more rules from the specified IP access control group.
func workspaces_RevokeIpRules(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.RevokeIpRulesInput{
		// GroupId: *string, // Required
		// UserRules: []string, // Required
	}

	if len(_workspacesGroupId) > 0 {
		input.GroupId = aws.String(_workspacesGroupId)
	}
	if len(_workspacesUserRules) > 0 {
		input.UserRules = []string{_workspacesUserRules}
	}

	if resp, err := client.RevokeIpRules(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts the specified WorkSpaces.
// You cannot start a WorkSpace unless it has a running mode of AutoStop or Manual
// and a state of STOPPED .
func workspaces_StartWorkspaces(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.StartWorkspacesInput{
		// StartWorkspaceRequests: []types.StartRequest, // Required
	}

	if len(_workspacesStartWorkspaceRequests) > 0 {
		if err := assignInputField(input, "StartWorkspaceRequests", _workspacesStartWorkspaceRequests); err != nil {
			log.Errorf("invalid --start-workspace-requests: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartWorkspaces(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts the specified pool.
// You cannot start a pool unless it has a running mode of AutoStop and a state of
// STOPPED .
func workspaces_StartWorkspacesPool(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.StartWorkspacesPoolInput{
		// PoolId: *string, // Required
	}

	if len(_workspacesPoolId) > 0 {
		input.PoolId = aws.String(_workspacesPoolId)
	}

	if resp, err := client.StartWorkspacesPool(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops the specified WorkSpaces.
// You cannot stop a WorkSpace unless it has a running mode of AutoStop or Manual
// and a state of AVAILABLE , IMPAIRED , UNHEALTHY , or ERROR .
func workspaces_StopWorkspaces(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.StopWorkspacesInput{
		// StopWorkspaceRequests: []types.StopRequest, // Required
	}

	if len(_workspacesStopWorkspaceRequests) > 0 {
		if err := assignInputField(input, "StopWorkspaceRequests", _workspacesStopWorkspaceRequests); err != nil {
			log.Errorf("invalid --stop-workspace-requests: %s", err.Error())
			return
		}
	}

	if resp, err := client.StopWorkspaces(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops the specified pool.
// You cannot stop a WorkSpace pool unless it has a running mode of AutoStop and a
// state of AVAILABLE , IMPAIRED , UNHEALTHY , or ERROR .
func workspaces_StopWorkspacesPool(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.StopWorkspacesPoolInput{
		// PoolId: *string, // Required
	}

	if len(_workspacesPoolId) > 0 {
		input.PoolId = aws.String(_workspacesPoolId)
	}

	if resp, err := client.StopWorkspacesPool(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Terminates the specified WorkSpaces.
// Terminating a WorkSpace is a permanent action and cannot be undone. The user's
// data is destroyed. If you need to archive any user data, contact Amazon Web
// Services Support before terminating the WorkSpace.
//
// You can terminate a WorkSpace that is in any state except SUSPENDED .
//
// This operation is asynchronous and returns before the WorkSpaces have been
// completely terminated. After a WorkSpace is terminated, the TERMINATED state is
// returned only briefly before the WorkSpace directory metadata is cleaned up, so
// this state is rarely returned. To confirm that a WorkSpace is terminated, check
// for the WorkSpace ID by using [DescribeWorkSpaces]. If the WorkSpace ID isn't returned, then the
// WorkSpace has been successfully terminated.
//
// Simple AD and AD Connector are made available to you free of charge to use with
// WorkSpaces. If there are no WorkSpaces being used with your Simple AD or AD
// Connector directory for 30 consecutive days, this directory will be
// automatically deregistered for use with Amazon WorkSpaces, and you will be
// charged for this directory as per the [Directory Service pricing terms].
//
// To delete empty directories, see [Delete the Directory for Your WorkSpaces]. If you delete your Simple AD or AD Connector
// directory, you can always create a new one when you want to start using
// WorkSpaces again.
//
// [DescribeWorkSpaces]: https://docs.aws.amazon.com/workspaces/latest/api/API_DescribeWorkspaces.html
// [Delete the Directory for Your WorkSpaces]: https://docs.aws.amazon.com/workspaces/latest/adminguide/delete-workspaces-directory.html
// [Directory Service pricing terms]: http://aws.amazon.com/directoryservice/pricing/
func workspaces_TerminateWorkspaces(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.TerminateWorkspacesInput{
		// TerminateWorkspaceRequests: []types.TerminateRequest, // Required
	}

	if len(_workspacesTerminateWorkspaceRequests) > 0 {
		if err := assignInputField(input, "TerminateWorkspaceRequests", _workspacesTerminateWorkspaceRequests); err != nil {
			log.Errorf("invalid --terminate-workspace-requests: %s", err.Error())
			return
		}
	}

	if resp, err := client.TerminateWorkspaces(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Terminates the specified pool.
func workspaces_TerminateWorkspacesPool(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.TerminateWorkspacesPoolInput{
		// PoolId: *string, // Required
	}

	if len(_workspacesPoolId) > 0 {
		input.PoolId = aws.String(_workspacesPoolId)
	}

	if resp, err := client.TerminateWorkspacesPool(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Terminates the pool session.
func workspaces_TerminateWorkspacesPoolSession(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.TerminateWorkspacesPoolSessionInput{
		// SessionId: *string, // Required
	}

	if len(_workspacesSessionId) > 0 {
		input.SessionId = aws.String(_workspacesSessionId)
	}

	if resp, err := client.TerminateWorkspacesPoolSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a Amazon Connect client add-in. Use this action to update the name and
// endpoint URL of a Amazon Connect client add-in.
func workspaces_UpdateConnectClientAddIn(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.UpdateConnectClientAddInInput{
		// AddInId: *string, // Required
		// ResourceId: *string, // Required
	}

	if len(_workspacesAddInId) > 0 {
		input.AddInId = aws.String(_workspacesAddInId)
	}
	if len(_workspacesResourceId) > 0 {
		input.ResourceId = aws.String(_workspacesResourceId)
	}
	if len(_workspacesName) > 0 {
		input.Name = aws.String(_workspacesName)
	}
	if len(_workspacesURL) > 0 {
		input.URL = aws.String(_workspacesURL)
	}

	if resp, err := client.UpdateConnectClientAddIn(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Shares or unshares a connection alias with one account by specifying whether
// that account has permission to associate the connection alias with a directory.
// If the association permission is granted, the connection alias is shared with
// that account. If the association permission is revoked, the connection alias is
// unshared with the account. For more information, see [Cross-Region Redirection for Amazon WorkSpaces].
//
// - Before performing this operation, call [DescribeConnectionAliases]to make sure that the current state
// of the connection alias is CREATED .
//
// - To delete a connection alias that has been shared, the shared account must
// first disassociate the connection alias from any directories it has been
// associated with. Then you must unshare the connection alias from the account it
// has been shared with. You can delete a connection alias only after it is no
// longer shared with any accounts or associated with any directories.
//
// [DescribeConnectionAliases]: https://docs.aws.amazon.com/workspaces/latest/api/API_DescribeConnectionAliases.html
// [Cross-Region Redirection for Amazon WorkSpaces]: https://docs.aws.amazon.com/workspaces/latest/adminguide/cross-region-redirection.html
func workspaces_UpdateConnectionAliasPermission(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.UpdateConnectionAliasPermissionInput{
		// AliasId: *string, // Required
		// ConnectionAliasPermission: *types.ConnectionAliasPermission, // Required
	}

	if len(_workspacesAliasId) > 0 {
		input.AliasId = aws.String(_workspacesAliasId)
	}
	if len(_workspacesConnectionAliasPermission) > 0 {
		if err := assignInputField(input, "ConnectionAliasPermission", _workspacesConnectionAliasPermission); err != nil {
			log.Errorf("invalid --connection-alias-permission: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateConnectionAliasPermission(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Replaces the current rules of the specified IP access control group with the
// specified rules.
func workspaces_UpdateRulesOfIpGroup(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.UpdateRulesOfIpGroupInput{
		// GroupId: *string, // Required
		// UserRules: []types.IpRuleItem, // Required
	}

	if len(_workspacesGroupId) > 0 {
		input.GroupId = aws.String(_workspacesGroupId)
	}
	if len(_workspacesUserRules) > 0 {
		if err := assignInputField(input, "UserRules", _workspacesUserRules); err != nil {
			log.Errorf("invalid --user-rules: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateRulesOfIpGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a WorkSpace bundle with a new image. For more information about
// updating WorkSpace bundles, see [Update a Custom WorkSpaces Bundle].
//
// Existing WorkSpaces aren't automatically updated when you update the bundle
// that they're based on. To update existing WorkSpaces that are based on a bundle
// that you've updated, you must either rebuild the WorkSpaces or delete and
// recreate them.
//
// [Update a Custom WorkSpaces Bundle]: https://docs.aws.amazon.com/workspaces/latest/adminguide/update-custom-bundle.html
func workspaces_UpdateWorkspaceBundle(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.UpdateWorkspaceBundleInput{}

	if len(_workspacesBundleId) > 0 {
		input.BundleId = aws.String(_workspacesBundleId)
	}
	if len(_workspacesImageId) > 0 {
		input.ImageId = aws.String(_workspacesImageId)
	}

	if resp, err := client.UpdateWorkspaceBundle(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Shares or unshares an image with one account in the same Amazon Web Services
// Region by specifying whether that account has permission to copy the image. If
// the copy image permission is granted, the image is shared with that account. If
// the copy image permission is revoked, the image is unshared with the account.
//
// After an image has been shared, the recipient account can copy the image to
// other Regions as needed.
//
// In the China (Ningxia) Region, you can copy images only within the same Region.
//
// In Amazon Web Services GovCloud (US), to copy images to and from other Regions,
// contact Amazon Web Services Support.
//
// For more information about sharing images, see [Share or Unshare a Custom WorkSpaces Image].
//
// - To delete an image that has been shared, you must unshare the image before
// you delete it.
//
// - Sharing Bring Your Own License (BYOL) images across Amazon Web Services
// accounts isn't supported at this time in Amazon Web Services GovCloud (US). To
// share BYOL images across accounts in Amazon Web Services GovCloud (US), contact
// Amazon Web Services Support.
//
// [Share or Unshare a Custom WorkSpaces Image]: https://docs.aws.amazon.com/workspaces/latest/adminguide/share-custom-image.html
func workspaces_UpdateWorkspaceImagePermission(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.UpdateWorkspaceImagePermissionInput{
		// AllowCopyImage: *bool, // Required
		// ImageId: *string, // Required
		// SharedAccountId: *string, // Required
	}

	if len(_workspacesAllowCopyImage) > 0 {
		if err := assignInputField(input, "AllowCopyImage", _workspacesAllowCopyImage); err != nil {
			log.Errorf("invalid --allow-copy-image: %s", err.Error())
			return
		}
	}
	if len(_workspacesImageId) > 0 {
		input.ImageId = aws.String(_workspacesImageId)
	}
	if len(_workspacesSharedAccountId) > 0 {
		input.SharedAccountId = aws.String(_workspacesSharedAccountId)
	}

	if resp, err := client.UpdateWorkspaceImagePermission(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified pool.
func workspaces_UpdateWorkspacesPool(cfg aws.Config, client *workspaces.Client) {
	input := &workspaces.UpdateWorkspacesPoolInput{
		// PoolId: *string, // Required
	}

	if len(_workspacesPoolId) > 0 {
		input.PoolId = aws.String(_workspacesPoolId)
	}
	if len(_workspacesApplicationSettings) > 0 {
		if err := assignInputField(input, "ApplicationSettings", _workspacesApplicationSettings); err != nil {
			log.Errorf("invalid --application-settings: %s", err.Error())
			return
		}
	}
	if len(_workspacesBundleId) > 0 {
		input.BundleId = aws.String(_workspacesBundleId)
	}
	if len(_workspacesCapacity) > 0 {
		if err := assignInputField(input, "Capacity", _workspacesCapacity); err != nil {
			log.Errorf("invalid --capacity: %s", err.Error())
			return
		}
	}
	if len(_workspacesDescription) > 0 {
		input.Description = aws.String(_workspacesDescription)
	}
	if len(_workspacesDirectoryId) > 0 {
		input.DirectoryId = aws.String(_workspacesDirectoryId)
	}
	if len(_workspacesRunningMode) > 0 {
		if err := assignInputField(input, "RunningMode", _workspacesRunningMode); err != nil {
			log.Errorf("invalid --running-mode: %s", err.Error())
			return
		}
	}
	if len(_workspacesTimeoutSettings) > 0 {
		if err := assignInputField(input, "TimeoutSettings", _workspacesTimeoutSettings); err != nil {
			log.Errorf("invalid --timeout-settings: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateWorkspacesPool(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_workspacesCmd)
	_workspacesCmd.Flags().SortFlags = false

	_workspacesCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_workspacesCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_workspacesCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_workspacesCmd.Flags().StringVarP(&_workspacesActiveDirectoryConfig, "active-directory-config", "", "", "Active Directory Config")
	_workspacesCmd.Flags().StringVarP(&_workspacesAddInId, "add-in-id", "", "", "Add In ID")
	_workspacesCmd.Flags().StringVarP(&_workspacesAliasId, "alias-id", "", "", "Alias ID")
	_workspacesCmd.Flags().StringSliceVarP(&_workspacesAliasIds, "alias-ids", "", nil, "Alias Ids")
	_workspacesCmd.Flags().StringVarP(&_workspacesAllowCopyImage, "allow-copy-image", "", "", "Allow Copy Image")
	_workspacesCmd.Flags().StringVarP(&_workspacesApplicationId, "application-id", "", "", "Application ID")
	_workspacesCmd.Flags().StringSliceVarP(&_workspacesApplicationIds, "application-ids", "", nil, "Application Ids")
	_workspacesCmd.Flags().StringVarP(&_workspacesApplicationSettings, "application-settings", "", "", "Application Settings")
	_workspacesCmd.Flags().StringVarP(&_workspacesApplications, "applications", "", "", "Applications")
	_workspacesCmd.Flags().StringVarP(&_workspacesAssociatedResourceTypes, "associated-resource-types", "", "", "Associated Resource Types")
	_workspacesCmd.Flags().StringVarP(&_workspacesBundleDescription, "bundle-description", "", "", "Bundle Description")
	_workspacesCmd.Flags().StringVarP(&_workspacesBundleId, "bundle-id", "", "", "Bundle ID")
	_workspacesCmd.Flags().StringSliceVarP(&_workspacesBundleIds, "bundle-ids", "", nil, "Bundle Ids")
	_workspacesCmd.Flags().StringVarP(&_workspacesBundleName, "bundle-name", "", "", "Bundle Name")
	_workspacesCmd.Flags().StringVarP(&_workspacesCapacity, "capacity", "", "", "Capacity")
	_workspacesCmd.Flags().StringVarP(&_workspacesCertificateBasedAuthProperties, "certificate-based-auth-properties", "", "", "Certificate Based Auth Properties")
	_workspacesCmd.Flags().StringVarP(&_workspacesClientProperties, "client-properties", "", "", "Client Properties")
	_workspacesCmd.Flags().StringVarP(&_workspacesClientToken, "client-token", "", "", "Client Token")
	_workspacesCmd.Flags().StringVarP(&_workspacesComputeType, "compute-type", "", "", "Compute Type")
	_workspacesCmd.Flags().StringVarP(&_workspacesComputeTypeNames, "compute-type-names", "", "", "Compute Type Names")
	_workspacesCmd.Flags().StringVarP(&_workspacesConnectionAliasPermission, "connection-alias-permission", "", "", "Connection Alias Permission")
	_workspacesCmd.Flags().StringVarP(&_workspacesConnectionString, "connection-string", "", "", "Connection String")
	_workspacesCmd.Flags().StringVarP(&_workspacesDataReplication, "data-replication", "", "", "Data Replication")
	_workspacesCmd.Flags().StringVarP(&_workspacesDedicatedTenancyManagementCidrRange, "dedicated-tenancy-management-cidr-range", "", "", "Dedicated Tenancy Management CIDR Range")
	_workspacesCmd.Flags().StringVarP(&_workspacesDedicatedTenancySupport, "dedicated-tenancy-support", "", "", "Dedicated Tenancy Support")
	_workspacesCmd.Flags().StringVarP(&_workspacesDescription, "description", "", "", "Description")
	_workspacesCmd.Flags().StringVarP(&_workspacesDeviceTypeAndroid, "device-type-android", "", "", "Device Type Android")
	_workspacesCmd.Flags().StringVarP(&_workspacesDeviceTypeIos, "device-type-ios", "", "", "Device Type Ios")
	_workspacesCmd.Flags().StringVarP(&_workspacesDeviceTypeLinux, "device-type-linux", "", "", "Device Type Linux")
	_workspacesCmd.Flags().StringVarP(&_workspacesDeviceTypeOsx, "device-type-osx", "", "", "Device Type Osx")
	_workspacesCmd.Flags().StringVarP(&_workspacesDeviceTypeWeb, "device-type-web", "", "", "Device Type Web")
	_workspacesCmd.Flags().StringVarP(&_workspacesDeviceTypeWindows, "device-type-windows", "", "", "Device Type Windows")
	_workspacesCmd.Flags().StringVarP(&_workspacesDirectoryId, "directory-id", "", "", "Directory ID")
	_workspacesCmd.Flags().StringSliceVarP(&_workspacesDirectoryIds, "directory-ids", "", nil, "Directory Ids")
	_workspacesCmd.Flags().StringVarP(&_workspacesEc2ImageId, "ec2-image-id", "", "", "EC2 Image ID")
	_workspacesCmd.Flags().StringVarP(&_workspacesEnableSelfService, "enable-self-service", "", "", "Enable Self Service")
	_workspacesCmd.Flags().StringVarP(&_workspacesEndpointEncryptionMode, "endpoint-encryption-mode", "", "", "Endpoint Encryption Mode")
	_workspacesCmd.Flags().StringVarP(&_workspacesFilters, "filters", "", "", "Filters")
	_workspacesCmd.Flags().StringVarP(&_workspacesForce, "force", "", "", "Force")
	_workspacesCmd.Flags().StringVarP(&_workspacesGroupDesc, "group-desc", "", "", "Group Desc")
	_workspacesCmd.Flags().StringVarP(&_workspacesGroupId, "group-id", "", "", "Group ID")
	_workspacesCmd.Flags().StringSliceVarP(&_workspacesGroupIds, "group-ids", "", nil, "Group Ids")
	_workspacesCmd.Flags().StringVarP(&_workspacesGroupName, "group-name", "", "", "Group Name")
	_workspacesCmd.Flags().StringVarP(&_workspacesIdcInstanceArn, "idc-instance-arn", "", "", "Idc Instance ARN")
	_workspacesCmd.Flags().StringVarP(&_workspacesImageDescription, "image-description", "", "", "Image Description")
	_workspacesCmd.Flags().StringVarP(&_workspacesImageId, "image-id", "", "", "Image ID")
	_workspacesCmd.Flags().StringSliceVarP(&_workspacesImageIds, "image-ids", "", nil, "Image Ids")
	_workspacesCmd.Flags().StringVarP(&_workspacesImageName, "image-name", "", "", "Image Name")
	_workspacesCmd.Flags().StringVarP(&_workspacesImageSource, "image-source", "", "", "Image Source")
	_workspacesCmd.Flags().StringVarP(&_workspacesImageType, "image-type", "", "", "Image Type")
	_workspacesCmd.Flags().StringVarP(&_workspacesInfrastructureConfigurationArn, "infrastructure-configuration-arn", "", "", "Infrastructure Configuration ARN")
	_workspacesCmd.Flags().StringVarP(&_workspacesIngestionProcess, "ingestion-process", "", "", "Ingestion Process")
	_workspacesCmd.Flags().StringVarP(&_workspacesLicenseType, "license-type", "", "", "License Type")
	_workspacesCmd.Flags().StringVarP(&_workspacesLimit, "limit", "", "", "Limit")
	_workspacesCmd.Flags().StringVarP(&_workspacesLinkId, "link-id", "", "", "Link ID")
	_workspacesCmd.Flags().StringVarP(&_workspacesLinkStatusFilter, "link-status-filter", "", "", "Link Status Filter")
	_workspacesCmd.Flags().StringVarP(&_workspacesLinkedAccountId, "linked-account-id", "", "", "Linked Account ID")
	_workspacesCmd.Flags().StringVarP(&_workspacesManagementCidrRangeConstraint, "management-cidr-range-constraint", "", "", "Management CIDR Range Constraint")
	_workspacesCmd.Flags().StringVarP(&_workspacesMaxResults, "max-results", "", "", "Max Results")
	_workspacesCmd.Flags().StringVarP(&_workspacesMicrosoftEntraConfig, "microsoft-entra-config", "", "", "Microsoft Entra Config")
	_workspacesCmd.Flags().StringVarP(&_workspacesName, "name", "", "", "Name")
	_workspacesCmd.Flags().StringVarP(&_workspacesNextToken, "next-token", "", "", "Next Token")
	_workspacesCmd.Flags().StringVarP(&_workspacesOperatingSystemNames, "operating-system-names", "", "", "Operating System Names")
	_workspacesCmd.Flags().StringVarP(&_workspacesOsVersion, "os-version", "", "", "OS Version")
	_workspacesCmd.Flags().StringVarP(&_workspacesOwner, "owner", "", "", "Owner")
	_workspacesCmd.Flags().StringVarP(&_workspacesPlatform, "platform", "", "", "Platform")
	_workspacesCmd.Flags().StringVarP(&_workspacesPlatforms, "platforms", "", "", "Platforms")
	_workspacesCmd.Flags().StringVarP(&_workspacesPoolId, "pool-id", "", "", "Pool ID")
	_workspacesCmd.Flags().StringSliceVarP(&_workspacesPoolIds, "pool-ids", "", nil, "Pool Ids")
	_workspacesCmd.Flags().StringVarP(&_workspacesPoolName, "pool-name", "", "", "Pool Name")
	_workspacesCmd.Flags().StringVarP(&_workspacesPrimaryRegion, "primary-region", "", "", "Primary Region")
	_workspacesCmd.Flags().StringVarP(&_workspacesPropertiesToDelete, "properties-to-delete", "", "", "Properties To Delete")
	_workspacesCmd.Flags().StringVarP(&_workspacesProtocol, "protocol", "", "", "Protocol")
	_workspacesCmd.Flags().StringVarP(&_workspacesRebootWorkspaceRequests, "reboot-workspace-requests", "", "", "Reboot Workspace Requests")
	_workspacesCmd.Flags().StringVarP(&_workspacesRebuildWorkspaceRequests, "rebuild-workspace-requests", "", "", "Rebuild Workspace Requests")
	_workspacesCmd.Flags().StringVarP(&_workspacesResourceId, "resource-id", "", "", "Resource ID")
	_workspacesCmd.Flags().StringSliceVarP(&_workspacesResourceIds, "resource-ids", "", nil, "Resource Ids")
	_workspacesCmd.Flags().StringVarP(&_workspacesRootStorage, "root-storage", "", "", "Root Storage")
	_workspacesCmd.Flags().StringVarP(&_workspacesRunningMode, "running-mode", "", "", "Running Mode")
	_workspacesCmd.Flags().StringVarP(&_workspacesSamlProperties, "saml-properties", "", "", "Saml Properties")
	_workspacesCmd.Flags().StringVarP(&_workspacesSelfservicePermissions, "selfservice-permissions", "", "", "Selfservice Permissions")
	_workspacesCmd.Flags().StringVarP(&_workspacesSessionId, "session-id", "", "", "Session ID")
	_workspacesCmd.Flags().StringVarP(&_workspacesSharedAccountId, "shared-account-id", "", "", "Shared Account ID")
	_workspacesCmd.Flags().StringVarP(&_workspacesSourceImageId, "source-image-id", "", "", "Source Image ID")
	_workspacesCmd.Flags().StringVarP(&_workspacesSourceRegion, "source-region", "", "", "Source Region")
	_workspacesCmd.Flags().StringVarP(&_workspacesSourceWorkspaceId, "source-workspace-id", "", "", "Source Workspace ID")
	_workspacesCmd.Flags().StringVarP(&_workspacesStandbyWorkspaces, "standby-workspaces", "", "", "Standby Workspaces")
	_workspacesCmd.Flags().StringVarP(&_workspacesStartWorkspaceRequests, "start-workspace-requests", "", "", "Start Workspace Requests")
	_workspacesCmd.Flags().StringVarP(&_workspacesStopWorkspaceRequests, "stop-workspace-requests", "", "", "Stop Workspace Requests")
	_workspacesCmd.Flags().StringVarP(&_workspacesStreamingProperties, "streaming-properties", "", "", "Streaming Properties")
	_workspacesCmd.Flags().StringSliceVarP(&_workspacesSubnetIds, "subnet-ids", "", nil, "Subnet Ids")
	_workspacesCmd.Flags().StringSliceVarP(&_workspacesTagKeys, "tag-keys", "", nil, "Tag Keys")
	_workspacesCmd.Flags().StringVarP(&_workspacesTags, "tags", "", "", "Tags")
	_workspacesCmd.Flags().StringVarP(&_workspacesTargetAccountId, "target-account-id", "", "", "Target Account ID")
	_workspacesCmd.Flags().StringVarP(&_workspacesTenancy, "tenancy", "", "", "Tenancy")
	_workspacesCmd.Flags().StringVarP(&_workspacesTerminateWorkspaceRequests, "terminate-workspace-requests", "", "", "Terminate Workspace Requests")
	_workspacesCmd.Flags().StringVarP(&_workspacesTimeoutSettings, "timeout-settings", "", "", "Timeout Settings")
	_workspacesCmd.Flags().StringVarP(&_workspacesURL, "url", "", "", "URL")
	_workspacesCmd.Flags().StringVarP(&_workspacesUserId, "user-id", "", "", "User ID")
	_workspacesCmd.Flags().StringVarP(&_workspacesUserIdentityType, "user-identity-type", "", "", "User Identity Type")
	_workspacesCmd.Flags().StringVarP(&_workspacesUserName, "user-name", "", "", "User Name")
	_workspacesCmd.Flags().StringVarP(&_workspacesUserRules, "user-rules", "", "", "User Rules")
	_workspacesCmd.Flags().StringVarP(&_workspacesUserStorage, "user-storage", "", "", "User Storage")
	_workspacesCmd.Flags().StringVarP(&_workspacesWorkspaceAccessProperties, "workspace-access-properties", "", "", "Workspace Access Properties")
	_workspacesCmd.Flags().StringVarP(&_workspacesWorkspaceCreationProperties, "workspace-creation-properties", "", "", "Workspace Creation Properties")
	_workspacesCmd.Flags().StringVarP(&_workspacesWorkspaceDirectoryDescription, "workspace-directory-description", "", "", "Workspace Directory Description")
	_workspacesCmd.Flags().StringVarP(&_workspacesWorkspaceDirectoryName, "workspace-directory-name", "", "", "Workspace Directory Name")
	_workspacesCmd.Flags().StringSliceVarP(&_workspacesWorkspaceDirectoryNames, "workspace-directory-names", "", nil, "Workspace Directory Names")
	_workspacesCmd.Flags().StringVarP(&_workspacesWorkspaceId, "workspace-id", "", "", "Workspace ID")
	_workspacesCmd.Flags().StringSliceVarP(&_workspacesWorkspaceIds, "workspace-ids", "", nil, "Workspace Ids")
	_workspacesCmd.Flags().StringVarP(&_workspacesWorkspaceName, "workspace-name", "", "", "Workspace Name")
	_workspacesCmd.Flags().StringVarP(&_workspacesWorkspaceProperties, "workspace-properties", "", "", "Workspace Properties")
	_workspacesCmd.Flags().StringVarP(&_workspacesWorkspaceState, "workspace-state", "", "", "Workspace State")
	_workspacesCmd.Flags().StringVarP(&_workspacesWorkspaceType, "workspace-type", "", "", "Workspace Type")
	_workspacesCmd.Flags().StringVarP(&_workspacesWorkspaces, "workspaces", "", "", "Workspaces")

	_workspacesCmd.Flags().BoolVarP(&_workspacesAcceptAccountLinkInvitation, "accept-account-link-invitation", "", false, "Accept Account Link Invitation")
	_workspacesCmd.Flags().BoolVarP(&_workspacesAssociateConnectionAlias, "associate-connection-alias", "", false, "Associate Connection Alias")
	_workspacesCmd.Flags().BoolVarP(&_workspacesAssociateIpGroups, "associate-ip-groups", "", false, "Associate IP Groups")
	_workspacesCmd.Flags().BoolVarP(&_workspacesAssociateWorkspaceApplication, "associate-workspace-application", "", false, "Associate Workspace Application")
	_workspacesCmd.Flags().BoolVarP(&_workspacesAuthorizeIpRules, "authorize-ip-rules", "", false, "Authorize IP Rules")
	_workspacesCmd.Flags().BoolVarP(&_workspacesCopyWorkspaceImage, "copy-workspace-image", "", false, "Copy Workspace Image")
	_workspacesCmd.Flags().BoolVarP(&_workspacesCreateAccountLinkInvitation, "create-account-link-invitation", "", false, "Create Account Link Invitation")
	_workspacesCmd.Flags().BoolVarP(&_workspacesCreateConnectClientAddIn, "create-connect-client-add-in", "", false, "Create Connect Client Add In")
	_workspacesCmd.Flags().BoolVarP(&_workspacesCreateConnectionAlias, "create-connection-alias", "", false, "Create Connection Alias")
	_workspacesCmd.Flags().BoolVarP(&_workspacesCreateIpGroup, "create-ip-group", "", false, "Create IP Group")
	_workspacesCmd.Flags().BoolVarP(&_workspacesCreateStandbyWorkspaces, "create-standby-workspaces", "", false, "Create Standby Workspaces")
	_workspacesCmd.Flags().BoolVarP(&_workspacesCreateTags, "create-tags", "", false, "Create Tags")
	_workspacesCmd.Flags().BoolVarP(&_workspacesCreateUpdatedWorkspaceImage, "create-updated-workspace-image", "", false, "Create Updated Workspace Image")
	_workspacesCmd.Flags().BoolVarP(&_workspacesCreateWorkspaceBundle, "create-workspace-bundle", "", false, "Create Workspace Bundle")
	_workspacesCmd.Flags().BoolVarP(&_workspacesCreateWorkspaceImage, "create-workspace-image", "", false, "Create Workspace Image")
	_workspacesCmd.Flags().BoolVarP(&_workspacesCreateWorkspaces, "create-workspaces", "", false, "Create Workspaces")
	_workspacesCmd.Flags().BoolVarP(&_workspacesCreateWorkspacesPool, "create-workspaces-pool", "", false, "Create Workspaces Pool")
	_workspacesCmd.Flags().BoolVarP(&_workspacesDeleteAccountLinkInvitation, "delete-account-link-invitation", "", false, "Delete Account Link Invitation")
	_workspacesCmd.Flags().BoolVarP(&_workspacesDeleteClientBranding, "delete-client-branding", "", false, "Delete Client Branding")
	_workspacesCmd.Flags().BoolVarP(&_workspacesDeleteConnectClientAddIn, "delete-connect-client-add-in", "", false, "Delete Connect Client Add In")
	_workspacesCmd.Flags().BoolVarP(&_workspacesDeleteConnectionAlias, "delete-connection-alias", "", false, "Delete Connection Alias")
	_workspacesCmd.Flags().BoolVarP(&_workspacesDeleteIpGroup, "delete-ip-group", "", false, "Delete IP Group")
	_workspacesCmd.Flags().BoolVarP(&_workspacesDeleteTags, "delete-tags", "", false, "Delete Tags")
	_workspacesCmd.Flags().BoolVarP(&_workspacesDeleteWorkspaceBundle, "delete-workspace-bundle", "", false, "Delete Workspace Bundle")
	_workspacesCmd.Flags().BoolVarP(&_workspacesDeleteWorkspaceImage, "delete-workspace-image", "", false, "Delete Workspace Image")
	_workspacesCmd.Flags().BoolVarP(&_workspacesDeployWorkspaceApplications, "deploy-workspace-applications", "", false, "Deploy Workspace Applications")
	_workspacesCmd.Flags().BoolVarP(&_workspacesDeregisterWorkspaceDirectory, "deregister-workspace-directory", "", false, "Deregister Workspace Directory")
	_workspacesCmd.Flags().BoolVarP(&_workspacesDescribeAccount, "describe-account", "", false, "Describe Account")
	_workspacesCmd.Flags().BoolVarP(&_workspacesDescribeAccountModifications, "describe-account-modifications", "", false, "Describe Account Modifications")
	_workspacesCmd.Flags().BoolVarP(&_workspacesDescribeApplicationAssociations, "describe-application-associations", "", false, "Describe Application Associations")
	_workspacesCmd.Flags().BoolVarP(&_workspacesDescribeApplications, "describe-applications", "", false, "Describe Applications")
	_workspacesCmd.Flags().BoolVarP(&_workspacesDescribeBundleAssociations, "describe-bundle-associations", "", false, "Describe Bundle Associations")
	_workspacesCmd.Flags().BoolVarP(&_workspacesDescribeClientBranding, "describe-client-branding", "", false, "Describe Client Branding")
	_workspacesCmd.Flags().BoolVarP(&_workspacesDescribeClientProperties, "describe-client-properties", "", false, "Describe Client Properties")
	_workspacesCmd.Flags().BoolVarP(&_workspacesDescribeConnectClientAddIns, "describe-connect-client-add-ins", "", false, "Describe Connect Client Add Ins")
	_workspacesCmd.Flags().BoolVarP(&_workspacesDescribeConnectionAliasPermissions, "describe-connection-alias-permissions", "", false, "Describe Connection Alias Permissions")
	_workspacesCmd.Flags().BoolVarP(&_workspacesDescribeConnectionAliases, "describe-connection-aliases", "", false, "Describe Connection Aliases")
	_workspacesCmd.Flags().BoolVarP(&_workspacesDescribeCustomWorkspaceImageImport, "describe-custom-workspace-image-import", "", false, "Describe Custom Workspace Image Import")
	_workspacesCmd.Flags().BoolVarP(&_workspacesDescribeImageAssociations, "describe-image-associations", "", false, "Describe Image Associations")
	_workspacesCmd.Flags().BoolVarP(&_workspacesDescribeIpGroups, "describe-ip-groups", "", false, "Describe IP Groups")
	_workspacesCmd.Flags().BoolVarP(&_workspacesDescribeTags, "describe-tags", "", false, "Describe Tags")
	_workspacesCmd.Flags().BoolVarP(&_workspacesDescribeWorkspaceAssociations, "describe-workspace-associations", "", false, "Describe Workspace Associations")
	_workspacesCmd.Flags().BoolVarP(&_workspacesDescribeWorkspaceBundles, "describe-workspace-bundles", "", false, "Describe Workspace Bundles")
	_workspacesCmd.Flags().BoolVarP(&_workspacesDescribeWorkspaceDirectories, "describe-workspace-directories", "", false, "Describe Workspace Directories")
	_workspacesCmd.Flags().BoolVarP(&_workspacesDescribeWorkspaceImagePermissions, "describe-workspace-image-permissions", "", false, "Describe Workspace Image Permissions")
	_workspacesCmd.Flags().BoolVarP(&_workspacesDescribeWorkspaceImages, "describe-workspace-images", "", false, "Describe Workspace Images")
	_workspacesCmd.Flags().BoolVarP(&_workspacesDescribeWorkspaceSnapshots, "describe-workspace-snapshots", "", false, "Describe Workspace Snapshots")
	_workspacesCmd.Flags().BoolVarP(&_workspacesDescribeWorkspaces, "describe-workspaces", "", false, "Describe Workspaces")
	_workspacesCmd.Flags().BoolVarP(&_workspacesDescribeWorkspacesConnectionStatus, "describe-workspaces-connection-status", "", false, "Describe Workspaces Connection Status")
	_workspacesCmd.Flags().BoolVarP(&_workspacesDescribeWorkspacesPoolSessions, "describe-workspaces-pool-sessions", "", false, "Describe Workspaces Pool Sessions")
	_workspacesCmd.Flags().BoolVarP(&_workspacesDescribeWorkspacesPools, "describe-workspaces-pools", "", false, "Describe Workspaces Pools")
	_workspacesCmd.Flags().BoolVarP(&_workspacesDisassociateConnectionAlias, "disassociate-connection-alias", "", false, "Disassociate Connection Alias")
	_workspacesCmd.Flags().BoolVarP(&_workspacesDisassociateIpGroups, "disassociate-ip-groups", "", false, "Disassociate IP Groups")
	_workspacesCmd.Flags().BoolVarP(&_workspacesDisassociateWorkspaceApplication, "disassociate-workspace-application", "", false, "Disassociate Workspace Application")
	_workspacesCmd.Flags().BoolVarP(&_workspacesGetAccountLink, "get-account-link", "", false, "Get Account Link")
	_workspacesCmd.Flags().BoolVarP(&_workspacesImportClientBranding, "import-client-branding", "", false, "Import Client Branding")
	_workspacesCmd.Flags().BoolVarP(&_workspacesImportCustomWorkspaceImage, "import-custom-workspace-image", "", false, "Import Custom Workspace Image")
	_workspacesCmd.Flags().BoolVarP(&_workspacesImportWorkspaceImage, "import-workspace-image", "", false, "Import Workspace Image")
	_workspacesCmd.Flags().BoolVarP(&_workspacesListAccountLinks, "list-account-links", "", false, "List Account Links")
	_workspacesCmd.Flags().BoolVarP(&_workspacesListAvailableManagementCidrRanges, "list-available-management-cidr-ranges", "", false, "List Available Management CIDR Ranges")
	_workspacesCmd.Flags().BoolVarP(&_workspacesMigrateWorkspace, "migrate-workspace", "", false, "Migrate Workspace")
	_workspacesCmd.Flags().BoolVarP(&_workspacesModifyAccount, "modify-account", "", false, "Modify Account")
	_workspacesCmd.Flags().BoolVarP(&_workspacesModifyCertificateBasedAuthProperties, "modify-certificate-based-auth-properties", "", false, "Modify Certificate Based Auth Properties")
	_workspacesCmd.Flags().BoolVarP(&_workspacesModifyClientProperties, "modify-client-properties", "", false, "Modify Client Properties")
	_workspacesCmd.Flags().BoolVarP(&_workspacesModifyEndpointEncryptionMode, "modify-endpoint-encryption-mode", "", false, "Modify Endpoint Encryption Mode")
	_workspacesCmd.Flags().BoolVarP(&_workspacesModifySamlProperties, "modify-saml-properties", "", false, "Modify Saml Properties")
	_workspacesCmd.Flags().BoolVarP(&_workspacesModifySelfservicePermissions, "modify-selfservice-permissions", "", false, "Modify Selfservice Permissions")
	_workspacesCmd.Flags().BoolVarP(&_workspacesModifyStreamingProperties, "modify-streaming-properties", "", false, "Modify Streaming Properties")
	_workspacesCmd.Flags().BoolVarP(&_workspacesModifyWorkspaceAccessProperties, "modify-workspace-access-properties", "", false, "Modify Workspace Access Properties")
	_workspacesCmd.Flags().BoolVarP(&_workspacesModifyWorkspaceCreationProperties, "modify-workspace-creation-properties", "", false, "Modify Workspace Creation Properties")
	_workspacesCmd.Flags().BoolVarP(&_workspacesModifyWorkspaceProperties, "modify-workspace-properties", "", false, "Modify Workspace Properties")
	_workspacesCmd.Flags().BoolVarP(&_workspacesModifyWorkspaceState, "modify-workspace-state", "", false, "Modify Workspace State")
	_workspacesCmd.Flags().BoolVarP(&_workspacesRebootWorkspaces, "reboot-workspaces", "", false, "Reboot Workspaces")
	_workspacesCmd.Flags().BoolVarP(&_workspacesRebuildWorkspaces, "rebuild-workspaces", "", false, "Rebuild Workspaces")
	_workspacesCmd.Flags().BoolVarP(&_workspacesRegisterWorkspaceDirectory, "register-workspace-directory", "", false, "Register Workspace Directory")
	_workspacesCmd.Flags().BoolVarP(&_workspacesRejectAccountLinkInvitation, "reject-account-link-invitation", "", false, "Reject Account Link Invitation")
	_workspacesCmd.Flags().BoolVarP(&_workspacesRestoreWorkspace, "restore-workspace", "", false, "Restore Workspace")
	_workspacesCmd.Flags().BoolVarP(&_workspacesRevokeIpRules, "revoke-ip-rules", "", false, "Revoke IP Rules")
	_workspacesCmd.Flags().BoolVarP(&_workspacesStartWorkspaces, "start-workspaces", "", false, "Start Workspaces")
	_workspacesCmd.Flags().BoolVarP(&_workspacesStartWorkspacesPool, "start-workspaces-pool", "", false, "Start Workspaces Pool")
	_workspacesCmd.Flags().BoolVarP(&_workspacesStopWorkspaces, "stop-workspaces", "", false, "Stop Workspaces")
	_workspacesCmd.Flags().BoolVarP(&_workspacesStopWorkspacesPool, "stop-workspaces-pool", "", false, "Stop Workspaces Pool")
	_workspacesCmd.Flags().BoolVarP(&_workspacesTerminateWorkspaces, "terminate-workspaces", "", false, "Terminate Workspaces")
	_workspacesCmd.Flags().BoolVarP(&_workspacesTerminateWorkspacesPool, "terminate-workspaces-pool", "", false, "Terminate Workspaces Pool")
	_workspacesCmd.Flags().BoolVarP(&_workspacesTerminateWorkspacesPoolSession, "terminate-workspaces-pool-session", "", false, "Terminate Workspaces Pool Session")
	_workspacesCmd.Flags().BoolVarP(&_workspacesUpdateConnectClientAddIn, "update-connect-client-add-in", "", false, "Update Connect Client Add In")
	_workspacesCmd.Flags().BoolVarP(&_workspacesUpdateConnectionAliasPermission, "update-connection-alias-permission", "", false, "Update Connection Alias Permission")
	_workspacesCmd.Flags().BoolVarP(&_workspacesUpdateRulesOfIpGroup, "update-rules-of-ip-group", "", false, "Update Rules Of IP Group")
	_workspacesCmd.Flags().BoolVarP(&_workspacesUpdateWorkspaceBundle, "update-workspace-bundle", "", false, "Update Workspace Bundle")
	_workspacesCmd.Flags().BoolVarP(&_workspacesUpdateWorkspaceImagePermission, "update-workspace-image-permission", "", false, "Update Workspace Image Permission")
	_workspacesCmd.Flags().BoolVarP(&_workspacesUpdateWorkspacesPool, "update-workspaces-pool", "", false, "Update Workspaces Pool")

}
