package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// directoryserviceCmd represents the directoryservice command
var _directoryserviceCmd = &cobra.Command{
	Use:   "directoryservice",
	Short: "AWS directoryservice CLI",
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
		client := directoryservice.NewFromConfig(cfg)
		if _directoryserviceAcceptSharedDirectory {
			directoryservice_AcceptSharedDirectory(cfg, client)
			return
		}
		if _directoryserviceAddIpRoutes {
			directoryservice_AddIpRoutes(cfg, client)
			return
		}
		if _directoryserviceAddRegion {
			directoryservice_AddRegion(cfg, client)
			return
		}
		if _directoryserviceAddTagsToResource {
			directoryservice_AddTagsToResource(cfg, client)
			return
		}
		if _directoryserviceCancelSchemaExtension {
			directoryservice_CancelSchemaExtension(cfg, client)
			return
		}
		if _directoryserviceConnectDirectory {
			directoryservice_ConnectDirectory(cfg, client)
			return
		}
		if _directoryserviceCreateAlias {
			directoryservice_CreateAlias(cfg, client)
			return
		}
		if _directoryserviceCreateComputer {
			directoryservice_CreateComputer(cfg, client)
			return
		}
		if _directoryserviceCreateConditionalForwarder {
			directoryservice_CreateConditionalForwarder(cfg, client)
			return
		}
		if _directoryserviceCreateDirectory {
			directoryservice_CreateDirectory(cfg, client)
			return
		}
		if _directoryserviceCreateHybridAD {
			directoryservice_CreateHybridAD(cfg, client)
			return
		}
		if _directoryserviceCreateLogSubscription {
			directoryservice_CreateLogSubscription(cfg, client)
			return
		}
		if _directoryserviceCreateMicrosoftAD {
			directoryservice_CreateMicrosoftAD(cfg, client)
			return
		}
		if _directoryserviceCreateSnapshot {
			directoryservice_CreateSnapshot(cfg, client)
			return
		}
		if _directoryserviceCreateTrust {
			directoryservice_CreateTrust(cfg, client)
			return
		}
		if _directoryserviceDeleteADAssessment {
			directoryservice_DeleteADAssessment(cfg, client)
			return
		}
		if _directoryserviceDeleteConditionalForwarder {
			directoryservice_DeleteConditionalForwarder(cfg, client)
			return
		}
		if _directoryserviceDeleteDirectory {
			directoryservice_DeleteDirectory(cfg, client)
			return
		}
		if _directoryserviceDeleteLogSubscription {
			directoryservice_DeleteLogSubscription(cfg, client)
			return
		}
		if _directoryserviceDeleteSnapshot {
			directoryservice_DeleteSnapshot(cfg, client)
			return
		}
		if _directoryserviceDeleteTrust {
			directoryservice_DeleteTrust(cfg, client)
			return
		}
		if _directoryserviceDeregisterCertificate {
			directoryservice_DeregisterCertificate(cfg, client)
			return
		}
		if _directoryserviceDeregisterEventTopic {
			directoryservice_DeregisterEventTopic(cfg, client)
			return
		}
		if _directoryserviceDescribeADAssessment {
			directoryservice_DescribeADAssessment(cfg, client)
			return
		}
		if _directoryserviceDescribeCAEnrollmentPolicy {
			directoryservice_DescribeCAEnrollmentPolicy(cfg, client)
			return
		}
		if _directoryserviceDescribeCertificate {
			directoryservice_DescribeCertificate(cfg, client)
			return
		}
		if _directoryserviceDescribeClientAuthenticationSettings {
			directoryservice_DescribeClientAuthenticationSettings(cfg, client)
			return
		}
		if _directoryserviceDescribeConditionalForwarders {
			directoryservice_DescribeConditionalForwarders(cfg, client)
			return
		}
		if _directoryserviceDescribeDirectories {
			directoryservice_DescribeDirectories(cfg, client)
			return
		}
		if _directoryserviceDescribeDirectoryDataAccess {
			directoryservice_DescribeDirectoryDataAccess(cfg, client)
			return
		}
		if _directoryserviceDescribeDomainControllers {
			directoryservice_DescribeDomainControllers(cfg, client)
			return
		}
		if _directoryserviceDescribeEventTopics {
			directoryservice_DescribeEventTopics(cfg, client)
			return
		}
		if _directoryserviceDescribeHybridADUpdate {
			directoryservice_DescribeHybridADUpdate(cfg, client)
			return
		}
		if _directoryserviceDescribeLDAPSSettings {
			directoryservice_DescribeLDAPSSettings(cfg, client)
			return
		}
		if _directoryserviceDescribeRegions {
			directoryservice_DescribeRegions(cfg, client)
			return
		}
		if _directoryserviceDescribeSettings {
			directoryservice_DescribeSettings(cfg, client)
			return
		}
		if _directoryserviceDescribeSharedDirectories {
			directoryservice_DescribeSharedDirectories(cfg, client)
			return
		}
		if _directoryserviceDescribeSnapshots {
			directoryservice_DescribeSnapshots(cfg, client)
			return
		}
		if _directoryserviceDescribeTrusts {
			directoryservice_DescribeTrusts(cfg, client)
			return
		}
		if _directoryserviceDescribeUpdateDirectory {
			directoryservice_DescribeUpdateDirectory(cfg, client)
			return
		}
		if _directoryserviceDisableCAEnrollmentPolicy {
			directoryservice_DisableCAEnrollmentPolicy(cfg, client)
			return
		}
		if _directoryserviceDisableClientAuthentication {
			directoryservice_DisableClientAuthentication(cfg, client)
			return
		}
		if _directoryserviceDisableDirectoryDataAccess {
			directoryservice_DisableDirectoryDataAccess(cfg, client)
			return
		}
		if _directoryserviceDisableLDAPS {
			directoryservice_DisableLDAPS(cfg, client)
			return
		}
		if _directoryserviceDisableRadius {
			directoryservice_DisableRadius(cfg, client)
			return
		}
		if _directoryserviceDisableSso {
			directoryservice_DisableSso(cfg, client)
			return
		}
		if _directoryserviceEnableCAEnrollmentPolicy {
			directoryservice_EnableCAEnrollmentPolicy(cfg, client)
			return
		}
		if _directoryserviceEnableClientAuthentication {
			directoryservice_EnableClientAuthentication(cfg, client)
			return
		}
		if _directoryserviceEnableDirectoryDataAccess {
			directoryservice_EnableDirectoryDataAccess(cfg, client)
			return
		}
		if _directoryserviceEnableLDAPS {
			directoryservice_EnableLDAPS(cfg, client)
			return
		}
		if _directoryserviceEnableRadius {
			directoryservice_EnableRadius(cfg, client)
			return
		}
		if _directoryserviceEnableSso {
			directoryservice_EnableSso(cfg, client)
			return
		}
		if _directoryserviceGetDirectoryLimits {
			directoryservice_GetDirectoryLimits(cfg, client)
			return
		}
		if _directoryserviceGetSnapshotLimits {
			directoryservice_GetSnapshotLimits(cfg, client)
			return
		}
		if _directoryserviceListADAssessments {
			directoryservice_ListADAssessments(cfg, client)
			return
		}
		if _directoryserviceListCertificates {
			directoryservice_ListCertificates(cfg, client)
			return
		}
		if _directoryserviceListIpRoutes {
			directoryservice_ListIpRoutes(cfg, client)
			return
		}
		if _directoryserviceListLogSubscriptions {
			directoryservice_ListLogSubscriptions(cfg, client)
			return
		}
		if _directoryserviceListSchemaExtensions {
			directoryservice_ListSchemaExtensions(cfg, client)
			return
		}
		if _directoryserviceListTagsForResource {
			directoryservice_ListTagsForResource(cfg, client)
			return
		}
		if _directoryserviceRegisterCertificate {
			directoryservice_RegisterCertificate(cfg, client)
			return
		}
		if _directoryserviceRegisterEventTopic {
			directoryservice_RegisterEventTopic(cfg, client)
			return
		}
		if _directoryserviceRejectSharedDirectory {
			directoryservice_RejectSharedDirectory(cfg, client)
			return
		}
		if _directoryserviceRemoveIpRoutes {
			directoryservice_RemoveIpRoutes(cfg, client)
			return
		}
		if _directoryserviceRemoveRegion {
			directoryservice_RemoveRegion(cfg, client)
			return
		}
		if _directoryserviceRemoveTagsFromResource {
			directoryservice_RemoveTagsFromResource(cfg, client)
			return
		}
		if _directoryserviceResetUserPassword {
			directoryservice_ResetUserPassword(cfg, client)
			return
		}
		if _directoryserviceRestoreFromSnapshot {
			directoryservice_RestoreFromSnapshot(cfg, client)
			return
		}
		if _directoryserviceShareDirectory {
			directoryservice_ShareDirectory(cfg, client)
			return
		}
		if _directoryserviceStartADAssessment {
			directoryservice_StartADAssessment(cfg, client)
			return
		}
		if _directoryserviceStartSchemaExtension {
			directoryservice_StartSchemaExtension(cfg, client)
			return
		}
		if _directoryserviceUnshareDirectory {
			directoryservice_UnshareDirectory(cfg, client)
			return
		}
		if _directoryserviceUpdateConditionalForwarder {
			directoryservice_UpdateConditionalForwarder(cfg, client)
			return
		}
		if _directoryserviceUpdateDirectorySetup {
			directoryservice_UpdateDirectorySetup(cfg, client)
			return
		}
		if _directoryserviceUpdateHybridAD {
			directoryservice_UpdateHybridAD(cfg, client)
			return
		}
		if _directoryserviceUpdateNumberOfDomainControllers {
			directoryservice_UpdateNumberOfDomainControllers(cfg, client)
			return
		}
		if _directoryserviceUpdateRadius {
			directoryservice_UpdateRadius(cfg, client)
			return
		}
		if _directoryserviceUpdateSettings {
			directoryservice_UpdateSettings(cfg, client)
			return
		}
		if _directoryserviceUpdateTrust {
			directoryservice_UpdateTrust(cfg, client)
			return
		}
		if _directoryserviceVerifyTrust {
			directoryservice_VerifyTrust(cfg, client)
			return
		}

	},
}

var (
	_directoryserviceAcceptSharedDirectory                bool
	_directoryserviceAddIpRoutes                          bool
	_directoryserviceAddRegion                            bool
	_directoryserviceAddTagsToResource                    bool
	_directoryserviceCancelSchemaExtension                bool
	_directoryserviceConnectDirectory                     bool
	_directoryserviceCreateAlias                          bool
	_directoryserviceCreateComputer                       bool
	_directoryserviceCreateConditionalForwarder           bool
	_directoryserviceCreateDirectory                      bool
	_directoryserviceCreateHybridAD                       bool
	_directoryserviceCreateLogSubscription                bool
	_directoryserviceCreateMicrosoftAD                    bool
	_directoryserviceCreateSnapshot                       bool
	_directoryserviceCreateTrust                          bool
	_directoryserviceDeleteADAssessment                   bool
	_directoryserviceDeleteConditionalForwarder           bool
	_directoryserviceDeleteDirectory                      bool
	_directoryserviceDeleteLogSubscription                bool
	_directoryserviceDeleteSnapshot                       bool
	_directoryserviceDeleteTrust                          bool
	_directoryserviceDeregisterCertificate                bool
	_directoryserviceDeregisterEventTopic                 bool
	_directoryserviceDescribeADAssessment                 bool
	_directoryserviceDescribeCAEnrollmentPolicy           bool
	_directoryserviceDescribeCertificate                  bool
	_directoryserviceDescribeClientAuthenticationSettings bool
	_directoryserviceDescribeConditionalForwarders        bool
	_directoryserviceDescribeDirectories                  bool
	_directoryserviceDescribeDirectoryDataAccess          bool
	_directoryserviceDescribeDomainControllers            bool
	_directoryserviceDescribeEventTopics                  bool
	_directoryserviceDescribeHybridADUpdate               bool
	_directoryserviceDescribeLDAPSSettings                bool
	_directoryserviceDescribeRegions                      bool
	_directoryserviceDescribeSettings                     bool
	_directoryserviceDescribeSharedDirectories            bool
	_directoryserviceDescribeSnapshots                    bool
	_directoryserviceDescribeTrusts                       bool
	_directoryserviceDescribeUpdateDirectory              bool
	_directoryserviceDisableCAEnrollmentPolicy            bool
	_directoryserviceDisableClientAuthentication          bool
	_directoryserviceDisableDirectoryDataAccess           bool
	_directoryserviceDisableLDAPS                         bool
	_directoryserviceDisableRadius                        bool
	_directoryserviceDisableSso                           bool
	_directoryserviceEnableCAEnrollmentPolicy             bool
	_directoryserviceEnableClientAuthentication           bool
	_directoryserviceEnableDirectoryDataAccess            bool
	_directoryserviceEnableLDAPS                          bool
	_directoryserviceEnableRadius                         bool
	_directoryserviceEnableSso                            bool
	_directoryserviceGetDirectoryLimits                   bool
	_directoryserviceGetSnapshotLimits                    bool
	_directoryserviceListADAssessments                    bool
	_directoryserviceListCertificates                     bool
	_directoryserviceListIpRoutes                         bool
	_directoryserviceListLogSubscriptions                 bool
	_directoryserviceListSchemaExtensions                 bool
	_directoryserviceListTagsForResource                  bool
	_directoryserviceRegisterCertificate                  bool
	_directoryserviceRegisterEventTopic                   bool
	_directoryserviceRejectSharedDirectory                bool
	_directoryserviceRemoveIpRoutes                       bool
	_directoryserviceRemoveRegion                         bool
	_directoryserviceRemoveTagsFromResource               bool
	_directoryserviceResetUserPassword                    bool
	_directoryserviceRestoreFromSnapshot                  bool
	_directoryserviceShareDirectory                       bool
	_directoryserviceStartADAssessment                    bool
	_directoryserviceStartSchemaExtension                 bool
	_directoryserviceUnshareDirectory                     bool
	_directoryserviceUpdateConditionalForwarder           bool
	_directoryserviceUpdateDirectorySetup                 bool
	_directoryserviceUpdateHybridAD                       bool
	_directoryserviceUpdateNumberOfDomainControllers      bool
	_directoryserviceUpdateRadius                         bool
	_directoryserviceUpdateSettings                       bool
	_directoryserviceUpdateTrust                          bool
	_directoryserviceVerifyTrust                          bool

	_directoryserviceAlias                                      string
	_directoryserviceAssessmentConfiguration                    string
	_directoryserviceAssessmentId                               string
	_directoryserviceCertificateData                            string
	_directoryserviceCertificateId                              string
	_directoryserviceCidrIps                                    []string
	_directoryserviceCidrIpv6s                                  []string
	_directoryserviceClientCertAuthSettings                     string
	_directoryserviceComputerAttributes                         string
	_directoryserviceComputerName                               string
	_directoryserviceConditionalForwarderIpAddrs                []string
	_directoryserviceConditionalForwarderIpv6Addrs              []string
	_directoryserviceConnectSettings                            string
	_directoryserviceCreateSnapshotBeforeSchemaExtension        string
	_directoryserviceCreateSnapshotBeforeUpdate                 string
	_directoryserviceDeleteAssociatedConditionalForwarder       string
	_directoryserviceDescription                                string
	_directoryserviceDesiredNumber                              string
	_directoryserviceDirectoryId                                string
	_directoryserviceDirectoryIds                               []string
	_directoryserviceDirectorySizeUpdateSettings                string
	_directoryserviceDnsIpAddrs                                 []string
	_directoryserviceDnsIpv6Addrs                               []string
	_directoryserviceDomainControllerIds                        []string
	_directoryserviceEdition                                    string
	_directoryserviceHybridAdministratorAccountUpdate           string
	_directoryserviceIpRoutes                                   string
	_directoryserviceLdifContent                                string
	_directoryserviceLimit                                      string
	_directoryserviceLogGroupName                               string
	_directoryserviceName                                       string
	_directoryserviceNetworkType                                string
	_directoryserviceNetworkUpdateSettings                      string
	_directoryserviceNewPassword                                string
	_directoryserviceNextToken                                  string
	_directoryserviceOrganizationalUnitDistinguishedName        string
	_directoryserviceOSUpdateSettings                           string
	_directoryserviceOwnerDirectoryId                           string
	_directoryservicePassword                                   string
	_directoryservicePcaConnectorArn                            string
	_directoryserviceRadiusSettings                             string
	_directoryserviceRegionName                                 string
	_directoryserviceRemoteDomainName                           string
	_directoryserviceRemoteDomainNames                          []string
	_directoryserviceResourceId                                 string
	_directoryserviceSchemaExtensionId                          string
	_directoryserviceSecretArn                                  string
	_directoryserviceSelectiveAuth                              string
	_directoryserviceSelfManagedInstancesSettings               string
	_directoryserviceSettings                                   string
	_directoryserviceShareMethod                                string
	_directoryserviceShareNotes                                 string
	_directoryserviceShareTarget                                string
	_directoryserviceSharedDirectoryId                          string
	_directoryserviceSharedDirectoryIds                         []string
	_directoryserviceShortName                                  string
	_directoryserviceSize                                       string
	_directoryserviceSnapshotId                                 string
	_directoryserviceSnapshotIds                                []string
	_directoryserviceStatus                                     string
	_directoryserviceTagKeys                                    []string
	_directoryserviceTags                                       string
	_directoryserviceTopicName                                  string
	_directoryserviceTopicNames                                 []string
	_directoryserviceTrustDirection                             string
	_directoryserviceTrustId                                    string
	_directoryserviceTrustIds                                   []string
	_directoryserviceTrustPassword                              string
	_directoryserviceTrustType                                  string
	_directoryserviceType                                       string
	_directoryserviceUnshareTarget                              string
	_directoryserviceUpdateSecurityGroupForDirectoryControllers string
	_directoryserviceUpdateType                                 string
	_directoryserviceUserName                                   string
	_directoryserviceVPCSettings                                string
)

// Accepts a directory sharing request that was sent from the directory owner
// account.
func directoryservice_AcceptSharedDirectory(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.AcceptSharedDirectoryInput{
		// SharedDirectoryId: *string, // Required
	}

	if len(_directoryserviceSharedDirectoryId) > 0 {
		input.SharedDirectoryId = aws.String(_directoryserviceSharedDirectoryId)
	}

	if resp, err := client.AcceptSharedDirectory(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// If the DNS server for your self-managed domain uses a publicly addressable IP
// address, you must add a CIDR address block to correctly route traffic to and
// from your Microsoft AD on Amazon Web Services. AddIpRoutes adds this address
// block. You can also use AddIpRoutes to facilitate routing traffic that uses
// public IP ranges from your Microsoft AD on Amazon Web Services to a peer VPC.
//
// Before you call AddIpRoutes, ensure that all of the required permissions have
// been explicitly granted through a policy. For details about what permissions are
// required to run the AddIpRoutes operation, see [Directory Service API Permissions: Actions, Resources, and Conditions Reference].
//
// [Directory Service API Permissions: Actions, Resources, and Conditions Reference]: http://docs.aws.amazon.com/directoryservice/latest/admin-guide/UsingWithDS_IAM_ResourcePermissions.html
func directoryservice_AddIpRoutes(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.AddIpRoutesInput{
		// DirectoryId: *string, // Required
		// IpRoutes: []types.IpRoute, // Required
	}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}
	if len(_directoryserviceIpRoutes) > 0 {
		if err := assignInputField(input, "IpRoutes", _directoryserviceIpRoutes); err != nil {
			log.Errorf("invalid --ip-routes: %s", err.Error())
			return
		}
	}
	if len(_directoryserviceUpdateSecurityGroupForDirectoryControllers) > 0 {
		if err := assignInputField(input, "UpdateSecurityGroupForDirectoryControllers", _directoryserviceUpdateSecurityGroupForDirectoryControllers); err != nil {
			log.Errorf("invalid --update-security-group-for-directory-controllers: %s", err.Error())
			return
		}
	}

	if resp, err := client.AddIpRoutes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds two domain controllers in the specified Region for the specified directory.
func directoryservice_AddRegion(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.AddRegionInput{
		// DirectoryId: *string, // Required
		// RegionName: *string, // Required
		// VPCSettings: *types.DirectoryVpcSettings, // Required
	}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}
	if len(_directoryserviceRegionName) > 0 {
		input.RegionName = aws.String(_directoryserviceRegionName)
	}
	if len(_directoryserviceVPCSettings) > 0 {
		if err := assignInputField(input, "VPCSettings", _directoryserviceVPCSettings); err != nil {
			log.Errorf("invalid --vpc-settings: %s", err.Error())
			return
		}
	}

	if resp, err := client.AddRegion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or overwrites one or more tags for the specified directory. Each directory
// can have a maximum of 50 tags. Each tag consists of a key and optional value.
// Tag keys must be unique to each resource.
func directoryservice_AddTagsToResource(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.AddTagsToResourceInput{
		// ResourceId: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_directoryserviceResourceId) > 0 {
		input.ResourceId = aws.String(_directoryserviceResourceId)
	}
	if len(_directoryserviceTags) > 0 {
		if err := assignInputField(input, "Tags", _directoryserviceTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.AddTagsToResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels an in-progress schema extension to a Microsoft AD directory. Once a
// schema extension has started replicating to all domain controllers, the task can
// no longer be canceled. A schema extension can be canceled during any of the
// following states; Initializing , CreatingSnapshot , and UpdatingSchema .
func directoryservice_CancelSchemaExtension(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.CancelSchemaExtensionInput{
		// DirectoryId: *string, // Required
		// SchemaExtensionId: *string, // Required
	}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}
	if len(_directoryserviceSchemaExtensionId) > 0 {
		input.SchemaExtensionId = aws.String(_directoryserviceSchemaExtensionId)
	}

	if resp, err := client.CancelSchemaExtension(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an AD Connector to connect to a self-managed directory.
// Before you call ConnectDirectory , ensure that all of the required permissions
// have been explicitly granted through a policy. For details about what
// permissions are required to run the ConnectDirectory operation, see [Directory Service API Permissions: Actions, Resources, and Conditions Reference].
//
// [Directory Service API Permissions: Actions, Resources, and Conditions Reference]: http://docs.aws.amazon.com/directoryservice/latest/admin-guide/UsingWithDS_IAM_ResourcePermissions.html
func directoryservice_ConnectDirectory(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.ConnectDirectoryInput{
		// ConnectSettings: *types.DirectoryConnectSettings, // Required
		// Name: *string, // Required
		// Password: *string, // Required
		// Size: types.DirectorySize, // Required
	}

	if len(_directoryserviceConnectSettings) > 0 {
		if err := assignInputField(input, "ConnectSettings", _directoryserviceConnectSettings); err != nil {
			log.Errorf("invalid --connect-settings: %s", err.Error())
			return
		}
	}
	if len(_directoryserviceName) > 0 {
		input.Name = aws.String(_directoryserviceName)
	}
	if len(_directoryservicePassword) > 0 {
		input.Password = aws.String(_directoryservicePassword)
	}
	if len(_directoryserviceSize) > 0 {
		if err := assignInputField(input, "Size", _directoryserviceSize); err != nil {
			log.Errorf("invalid --size: %s", err.Error())
			return
		}
	}
	if len(_directoryserviceDescription) > 0 {
		input.Description = aws.String(_directoryserviceDescription)
	}
	if len(_directoryserviceNetworkType) > 0 {
		if err := assignInputField(input, "NetworkType", _directoryserviceNetworkType); err != nil {
			log.Errorf("invalid --network-type: %s", err.Error())
			return
		}
	}
	if len(_directoryserviceShortName) > 0 {
		input.ShortName = aws.String(_directoryserviceShortName)
	}
	if len(_directoryserviceTags) > 0 {
		if err := assignInputField(input, "Tags", _directoryserviceTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.ConnectDirectory(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an alias for a directory and assigns the alias to the directory. The
// alias is used to construct the access URL for the directory, such as
// http://.awsapps.com .
//
// After an alias has been created, it cannot be deleted or reused, so this
// operation should only be used when absolutely necessary.
func directoryservice_CreateAlias(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.CreateAliasInput{
		// Alias: *string, // Required
		// DirectoryId: *string, // Required
	}

	if len(_directoryserviceAlias) > 0 {
		input.Alias = aws.String(_directoryserviceAlias)
	}
	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}

	if resp, err := client.CreateAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Active Directory computer object in the specified directory.
func directoryservice_CreateComputer(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.CreateComputerInput{
		// ComputerName: *string, // Required
		// DirectoryId: *string, // Required
		// Password: *string, // Required
	}

	if len(_directoryserviceComputerName) > 0 {
		input.ComputerName = aws.String(_directoryserviceComputerName)
	}
	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}
	if len(_directoryservicePassword) > 0 {
		input.Password = aws.String(_directoryservicePassword)
	}
	if len(_directoryserviceComputerAttributes) > 0 {
		if err := assignInputField(input, "ComputerAttributes", _directoryserviceComputerAttributes); err != nil {
			log.Errorf("invalid --computer-attributes: %s", err.Error())
			return
		}
	}
	if len(_directoryserviceOrganizationalUnitDistinguishedName) > 0 {
		input.OrganizationalUnitDistinguishedName = aws.String(_directoryserviceOrganizationalUnitDistinguishedName)
	}

	if resp, err := client.CreateComputer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a conditional forwarder associated with your Amazon Web Services
// directory. Conditional forwarders are required in order to set up a trust
// relationship with another domain. The conditional forwarder points to the
// trusted domain.
func directoryservice_CreateConditionalForwarder(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.CreateConditionalForwarderInput{
		// DirectoryId: *string, // Required
		// RemoteDomainName: *string, // Required
	}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}
	if len(_directoryserviceRemoteDomainName) > 0 {
		input.RemoteDomainName = aws.String(_directoryserviceRemoteDomainName)
	}
	if len(_directoryserviceDnsIpAddrs) > 0 {
		input.DnsIpAddrs = append([]string(nil), _directoryserviceDnsIpAddrs...)
	}
	if len(_directoryserviceDnsIpv6Addrs) > 0 {
		input.DnsIpv6Addrs = append([]string(nil), _directoryserviceDnsIpv6Addrs...)
	}

	if resp, err := client.CreateConditionalForwarder(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a Simple AD directory. For more information, see [Simple Active Directory] in the Directory
// Service Admin Guide.
//
// Before you call CreateDirectory , ensure that all of the required permissions
// have been explicitly granted through a policy. For details about what
// permissions are required to run the CreateDirectory operation, see [Directory Service API Permissions: Actions, Resources, and Conditions Reference].
//
// [Directory Service API Permissions: Actions, Resources, and Conditions Reference]: http://docs.aws.amazon.com/directoryservice/latest/admin-guide/UsingWithDS_IAM_ResourcePermissions.html
// [Simple Active Directory]: https://docs.aws.amazon.com/directoryservice/latest/admin-guide/directory_simple_ad.html
func directoryservice_CreateDirectory(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.CreateDirectoryInput{
		// Name: *string, // Required
		// Password: *string, // Required
		// Size: types.DirectorySize, // Required
	}

	if len(_directoryserviceName) > 0 {
		input.Name = aws.String(_directoryserviceName)
	}
	if len(_directoryservicePassword) > 0 {
		input.Password = aws.String(_directoryservicePassword)
	}
	if len(_directoryserviceSize) > 0 {
		if err := assignInputField(input, "Size", _directoryserviceSize); err != nil {
			log.Errorf("invalid --size: %s", err.Error())
			return
		}
	}
	if len(_directoryserviceDescription) > 0 {
		input.Description = aws.String(_directoryserviceDescription)
	}
	if len(_directoryserviceNetworkType) > 0 {
		if err := assignInputField(input, "NetworkType", _directoryserviceNetworkType); err != nil {
			log.Errorf("invalid --network-type: %s", err.Error())
			return
		}
	}
	if len(_directoryserviceShortName) > 0 {
		input.ShortName = aws.String(_directoryserviceShortName)
	}
	if len(_directoryserviceTags) > 0 {
		if err := assignInputField(input, "Tags", _directoryserviceTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_directoryserviceVPCSettings) > 0 {
		if err := assignInputField(input, "VpcSettings", _directoryserviceVPCSettings); err != nil {
			log.Errorf("invalid --vpc-settings: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDirectory(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a hybrid directory that connects your self-managed Active Directory
// (AD) infrastructure and Amazon Web Services.
//
// You must have a successful directory assessment using StartADAssessment to validate your
// environment compatibility before you use this operation.
//
// Updates are applied asynchronously. Use DescribeDirectories to monitor the progress of directory
// creation.
func directoryservice_CreateHybridAD(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.CreateHybridADInput{
		// AssessmentId: *string, // Required
		// SecretArn: *string, // Required
	}

	if len(_directoryserviceAssessmentId) > 0 {
		input.AssessmentId = aws.String(_directoryserviceAssessmentId)
	}
	if len(_directoryserviceSecretArn) > 0 {
		input.SecretArn = aws.String(_directoryserviceSecretArn)
	}
	if len(_directoryserviceTags) > 0 {
		if err := assignInputField(input, "Tags", _directoryserviceTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateHybridAD(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a subscription to forward real-time Directory Service domain controller
// security logs to the specified Amazon CloudWatch log group in your Amazon Web
// Services account.
func directoryservice_CreateLogSubscription(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.CreateLogSubscriptionInput{
		// DirectoryId: *string, // Required
		// LogGroupName: *string, // Required
	}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}
	if len(_directoryserviceLogGroupName) > 0 {
		input.LogGroupName = aws.String(_directoryserviceLogGroupName)
	}

	if resp, err := client.CreateLogSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a Microsoft AD directory in the Amazon Web Services Cloud. For more
// information, see [Managed Microsoft AD]in the Directory Service Admin Guide.
//
// Before you call CreateMicrosoftAD, ensure that all of the required permissions
// have been explicitly granted through a policy. For details about what
// permissions are required to run the CreateMicrosoftAD operation, see [Directory Service API Permissions: Actions, Resources, and Conditions Reference].
//
// [Managed Microsoft AD]: https://docs.aws.amazon.com/directoryservice/latest/admin-guide/directory_microsoft_ad.html
// [Directory Service API Permissions: Actions, Resources, and Conditions Reference]: http://docs.aws.amazon.com/directoryservice/latest/admin-guide/UsingWithDS_IAM_ResourcePermissions.html
func directoryservice_CreateMicrosoftAD(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.CreateMicrosoftADInput{
		// Name: *string, // Required
		// Password: *string, // Required
		// VpcSettings: *types.DirectoryVpcSettings, // Required
	}

	if len(_directoryserviceName) > 0 {
		input.Name = aws.String(_directoryserviceName)
	}
	if len(_directoryservicePassword) > 0 {
		input.Password = aws.String(_directoryservicePassword)
	}
	if len(_directoryserviceVPCSettings) > 0 {
		if err := assignInputField(input, "VpcSettings", _directoryserviceVPCSettings); err != nil {
			log.Errorf("invalid --vpc-settings: %s", err.Error())
			return
		}
	}
	if len(_directoryserviceDescription) > 0 {
		input.Description = aws.String(_directoryserviceDescription)
	}
	if len(_directoryserviceEdition) > 0 {
		if err := assignInputField(input, "Edition", _directoryserviceEdition); err != nil {
			log.Errorf("invalid --edition: %s", err.Error())
			return
		}
	}
	if len(_directoryserviceNetworkType) > 0 {
		if err := assignInputField(input, "NetworkType", _directoryserviceNetworkType); err != nil {
			log.Errorf("invalid --network-type: %s", err.Error())
			return
		}
	}
	if len(_directoryserviceShortName) > 0 {
		input.ShortName = aws.String(_directoryserviceShortName)
	}
	if len(_directoryserviceTags) > 0 {
		if err := assignInputField(input, "Tags", _directoryserviceTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateMicrosoftAD(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a snapshot of a Simple AD or Microsoft AD directory in the Amazon Web
// Services cloud.
//
// You cannot take snapshots of AD Connector directories.
func directoryservice_CreateSnapshot(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.CreateSnapshotInput{
		// DirectoryId: *string, // Required
	}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}
	if len(_directoryserviceName) > 0 {
		input.Name = aws.String(_directoryserviceName)
	}

	if resp, err := client.CreateSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Directory Service for Microsoft Active Directory allows you to configure trust
// relationships. For example, you can establish a trust between your Managed
// Microsoft AD directory, and your existing self-managed Microsoft Active
// Directory. This would allow you to provide users and groups access to resources
// in either domain, with a single set of credentials.
//
// This action initiates the creation of the Amazon Web Services side of a trust
// relationship between an Managed Microsoft AD directory and an external domain.
// You can create either a forest trust or an external trust.
func directoryservice_CreateTrust(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.CreateTrustInput{
		// DirectoryId: *string, // Required
		// RemoteDomainName: *string, // Required
		// TrustDirection: types.TrustDirection, // Required
		// TrustPassword: *string, // Required
	}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}
	if len(_directoryserviceRemoteDomainName) > 0 {
		input.RemoteDomainName = aws.String(_directoryserviceRemoteDomainName)
	}
	if len(_directoryserviceTrustDirection) > 0 {
		if err := assignInputField(input, "TrustDirection", _directoryserviceTrustDirection); err != nil {
			log.Errorf("invalid --trust-direction: %s", err.Error())
			return
		}
	}
	if len(_directoryserviceTrustPassword) > 0 {
		input.TrustPassword = aws.String(_directoryserviceTrustPassword)
	}
	if len(_directoryserviceConditionalForwarderIpAddrs) > 0 {
		input.ConditionalForwarderIpAddrs = append([]string(nil), _directoryserviceConditionalForwarderIpAddrs...)
	}
	if len(_directoryserviceConditionalForwarderIpv6Addrs) > 0 {
		input.ConditionalForwarderIpv6Addrs = append([]string(nil), _directoryserviceConditionalForwarderIpv6Addrs...)
	}
	if len(_directoryserviceSelectiveAuth) > 0 {
		if err := assignInputField(input, "SelectiveAuth", _directoryserviceSelectiveAuth); err != nil {
			log.Errorf("invalid --selective-auth: %s", err.Error())
			return
		}
	}
	if len(_directoryserviceTrustType) > 0 {
		if err := assignInputField(input, "TrustType", _directoryserviceTrustType); err != nil {
			log.Errorf("invalid --trust-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTrust(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a directory assessment and all associated data. This operation
// permanently removes the assessment results, validation reports, and
// configuration information.
//
// You cannot delete system-initiated assessments. You can delete customer-created
// assessments even if they are in progress.
func directoryservice_DeleteADAssessment(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.DeleteADAssessmentInput{
		// AssessmentId: *string, // Required
	}

	if len(_directoryserviceAssessmentId) > 0 {
		input.AssessmentId = aws.String(_directoryserviceAssessmentId)
	}

	if resp, err := client.DeleteADAssessment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a conditional forwarder that has been set up for your Amazon Web
// Services directory.
func directoryservice_DeleteConditionalForwarder(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.DeleteConditionalForwarderInput{
		// DirectoryId: *string, // Required
		// RemoteDomainName: *string, // Required
	}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}
	if len(_directoryserviceRemoteDomainName) > 0 {
		input.RemoteDomainName = aws.String(_directoryserviceRemoteDomainName)
	}

	if resp, err := client.DeleteConditionalForwarder(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Directory Service directory.
// Before you call DeleteDirectory , ensure that all of the required permissions
// have been explicitly granted through a policy. For details about what
// permissions are required to run the DeleteDirectory operation, see [Directory Service API Permissions: Actions, Resources, and Conditions Reference].
//
// [Directory Service API Permissions: Actions, Resources, and Conditions Reference]: http://docs.aws.amazon.com/directoryservice/latest/admin-guide/UsingWithDS_IAM_ResourcePermissions.html
func directoryservice_DeleteDirectory(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.DeleteDirectoryInput{
		// DirectoryId: *string, // Required
	}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}

	if resp, err := client.DeleteDirectory(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified log subscription.
func directoryservice_DeleteLogSubscription(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.DeleteLogSubscriptionInput{
		// DirectoryId: *string, // Required
	}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}

	if resp, err := client.DeleteLogSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a directory snapshot.
func directoryservice_DeleteSnapshot(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.DeleteSnapshotInput{
		// SnapshotId: *string, // Required
	}

	if len(_directoryserviceSnapshotId) > 0 {
		input.SnapshotId = aws.String(_directoryserviceSnapshotId)
	}

	if resp, err := client.DeleteSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing trust relationship between your Managed Microsoft AD
// directory and an external domain.
func directoryservice_DeleteTrust(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.DeleteTrustInput{
		// TrustId: *string, // Required
	}

	if len(_directoryserviceTrustId) > 0 {
		input.TrustId = aws.String(_directoryserviceTrustId)
	}
	if len(_directoryserviceDeleteAssociatedConditionalForwarder) > 0 {
		if err := assignInputField(input, "DeleteAssociatedConditionalForwarder", _directoryserviceDeleteAssociatedConditionalForwarder); err != nil {
			log.Errorf("invalid --delete-associated-conditional-forwarder: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteTrust(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes from the system the certificate that was registered for secure LDAP or
// client certificate authentication.
func directoryservice_DeregisterCertificate(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.DeregisterCertificateInput{
		// CertificateId: *string, // Required
		// DirectoryId: *string, // Required
	}

	if len(_directoryserviceCertificateId) > 0 {
		input.CertificateId = aws.String(_directoryserviceCertificateId)
	}
	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}

	if resp, err := client.DeregisterCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified directory as a publisher to the specified Amazon SNS
// topic.
func directoryservice_DeregisterEventTopic(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.DeregisterEventTopicInput{
		// DirectoryId: *string, // Required
		// TopicName: *string, // Required
	}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}
	if len(_directoryserviceTopicName) > 0 {
		input.TopicName = aws.String(_directoryserviceTopicName)
	}

	if resp, err := client.DeregisterEventTopic(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves detailed information about a directory assessment, including its
// current status, validation results, and configuration details. Use this
// operation to monitor assessment progress and review results.
func directoryservice_DescribeADAssessment(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.DescribeADAssessmentInput{
		// AssessmentId: *string, // Required
	}

	if len(_directoryserviceAssessmentId) > 0 {
		input.AssessmentId = aws.String(_directoryserviceAssessmentId)
	}

	if resp, err := client.DescribeADAssessment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves detailed information about the certificate authority (CA) enrollment
// policy for the specified directory. This policy determines how client
// certificates are automatically enrolled and managed through Amazon Web Services
// Private Certificate Authority.
func directoryservice_DescribeCAEnrollmentPolicy(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.DescribeCAEnrollmentPolicyInput{
		// DirectoryId: *string, // Required
	}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}

	if resp, err := client.DescribeCAEnrollmentPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Displays information about the certificate registered for secure LDAP or client
// certificate authentication.
func directoryservice_DescribeCertificate(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.DescribeCertificateInput{
		// CertificateId: *string, // Required
		// DirectoryId: *string, // Required
	}

	if len(_directoryserviceCertificateId) > 0 {
		input.CertificateId = aws.String(_directoryserviceCertificateId)
	}
	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}

	if resp, err := client.DescribeCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the type of client authentication for the specified
// directory, if the type is specified. If no type is specified, information about
// all client authentication types that are supported for the specified directory
// is retrieved. Currently, only SmartCard is supported.
func directoryservice_DescribeClientAuthenticationSettings(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.DescribeClientAuthenticationSettingsInput{
		// DirectoryId: *string, // Required
	}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}
	if len(_directoryserviceLimit) > 0 {
		if err := assignInputField(input, "Limit", _directoryserviceLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_directoryserviceNextToken) > 0 {
		input.NextToken = aws.String(_directoryserviceNextToken)
	}
	if len(_directoryserviceType) > 0 {
		if err := assignInputField(input, "Type", _directoryserviceType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeClientAuthenticationSettings(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*directoryservice.DescribeClientAuthenticationSettingsOutput
	p := directoryservice.NewDescribeClientAuthenticationSettingsPaginator(client, input)
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

// Obtains information about the conditional forwarders for this account.
// If no input parameters are provided for RemoteDomainNames, this request
// describes all conditional forwarders for the specified directory ID.
func directoryservice_DescribeConditionalForwarders(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.DescribeConditionalForwardersInput{
		// DirectoryId: *string, // Required
	}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}
	if len(_directoryserviceRemoteDomainNames) > 0 {
		input.RemoteDomainNames = append([]string(nil), _directoryserviceRemoteDomainNames...)
	}

	if resp, err := client.DescribeConditionalForwarders(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Obtains information about the directories that belong to this account.
// You can retrieve information about specific directories by passing the
// directory identifiers in the DirectoryIds parameter. Otherwise, all directories
// that belong to the current account are returned.
//
// This operation supports pagination with the use of the NextToken request and
// response parameters. If more results are available, the
// DescribeDirectoriesResult.NextToken member contains a token that you pass in the
// next call to DescribeDirectoriesto retrieve the next set of items.
//
// You can also specify a maximum number of return results with the Limit
// parameter.
func directoryservice_DescribeDirectories(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.DescribeDirectoriesInput{}

	if len(_directoryserviceDirectoryIds) > 0 {
		input.DirectoryIds = append([]string(nil), _directoryserviceDirectoryIds...)
	}
	if len(_directoryserviceLimit) > 0 {
		if err := assignInputField(input, "Limit", _directoryserviceLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_directoryserviceNextToken) > 0 {
		input.NextToken = aws.String(_directoryserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeDirectories(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*directoryservice.DescribeDirectoriesOutput
	p := directoryservice.NewDescribeDirectoriesPaginator(client, input)
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

// Obtains status of directory data access enablement through the Directory
// Service Data API for the specified directory.
func directoryservice_DescribeDirectoryDataAccess(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.DescribeDirectoryDataAccessInput{
		// DirectoryId: *string, // Required
	}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}

	if resp, err := client.DescribeDirectoryDataAccess(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides information about any domain controllers in your directory.
func directoryservice_DescribeDomainControllers(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.DescribeDomainControllersInput{
		// DirectoryId: *string, // Required
	}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}
	if len(_directoryserviceDomainControllerIds) > 0 {
		input.DomainControllerIds = append([]string(nil), _directoryserviceDomainControllerIds...)
	}
	if len(_directoryserviceLimit) > 0 {
		if err := assignInputField(input, "Limit", _directoryserviceLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_directoryserviceNextToken) > 0 {
		input.NextToken = aws.String(_directoryserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeDomainControllers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*directoryservice.DescribeDomainControllersOutput
	p := directoryservice.NewDescribeDomainControllersPaginator(client, input)
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

// Obtains information about which Amazon SNS topics receive status messages from
// the specified directory.
//
// If no input parameters are provided, such as DirectoryId or TopicName, this
// request describes all of the associations in the account.
func directoryservice_DescribeEventTopics(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.DescribeEventTopicsInput{}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}
	if len(_directoryserviceTopicNames) > 0 {
		input.TopicNames = append([]string(nil), _directoryserviceTopicNames...)
	}

	if resp, err := client.DescribeEventTopics(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about update activities for a hybrid directory. This
// operation provides details about configuration changes, administrator account
// updates, and self-managed instance settings (IDs and DNS IPs).
func directoryservice_DescribeHybridADUpdate(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.DescribeHybridADUpdateInput{
		// DirectoryId: *string, // Required
	}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}
	if len(_directoryserviceNextToken) > 0 {
		input.NextToken = aws.String(_directoryserviceNextToken)
	}
	if len(_directoryserviceUpdateType) > 0 {
		if err := assignInputField(input, "UpdateType", _directoryserviceUpdateType); err != nil {
			log.Errorf("invalid --update-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeHybridADUpdate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the status of LDAP security for the specified directory.
func directoryservice_DescribeLDAPSSettings(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.DescribeLDAPSSettingsInput{
		// DirectoryId: *string, // Required
	}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}
	if len(_directoryserviceLimit) > 0 {
		if err := assignInputField(input, "Limit", _directoryserviceLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_directoryserviceNextToken) > 0 {
		input.NextToken = aws.String(_directoryserviceNextToken)
	}
	if len(_directoryserviceType) > 0 {
		if err := assignInputField(input, "Type", _directoryserviceType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeLDAPSSettings(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*directoryservice.DescribeLDAPSSettingsOutput
	p := directoryservice.NewDescribeLDAPSSettingsPaginator(client, input)
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

// Provides information about the Regions that are configured for multi-Region
// replication.
func directoryservice_DescribeRegions(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.DescribeRegionsInput{
		// DirectoryId: *string, // Required
	}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}
	if len(_directoryserviceNextToken) > 0 {
		input.NextToken = aws.String(_directoryserviceNextToken)
	}
	if len(_directoryserviceRegionName) > 0 {
		input.RegionName = aws.String(_directoryserviceRegionName)
	}

	if disablePaginator() {
		if resp, err := client.DescribeRegions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*directoryservice.DescribeRegionsOutput
	p := directoryservice.NewDescribeRegionsPaginator(client, input)
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

// Retrieves information about the configurable settings for the specified
// directory.
func directoryservice_DescribeSettings(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.DescribeSettingsInput{
		// DirectoryId: *string, // Required
	}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}
	if len(_directoryserviceNextToken) > 0 {
		input.NextToken = aws.String(_directoryserviceNextToken)
	}
	if len(_directoryserviceStatus) > 0 {
		if err := assignInputField(input, "Status", _directoryserviceStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the shared directories in your account.
func directoryservice_DescribeSharedDirectories(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.DescribeSharedDirectoriesInput{
		// OwnerDirectoryId: *string, // Required
	}

	if len(_directoryserviceOwnerDirectoryId) > 0 {
		input.OwnerDirectoryId = aws.String(_directoryserviceOwnerDirectoryId)
	}
	if len(_directoryserviceLimit) > 0 {
		if err := assignInputField(input, "Limit", _directoryserviceLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_directoryserviceNextToken) > 0 {
		input.NextToken = aws.String(_directoryserviceNextToken)
	}
	if len(_directoryserviceSharedDirectoryIds) > 0 {
		input.SharedDirectoryIds = append([]string(nil), _directoryserviceSharedDirectoryIds...)
	}

	if disablePaginator() {
		if resp, err := client.DescribeSharedDirectories(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*directoryservice.DescribeSharedDirectoriesOutput
	p := directoryservice.NewDescribeSharedDirectoriesPaginator(client, input)
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

// Obtains information about the directory snapshots that belong to this account.
// This operation supports pagination with the use of the NextToken request and
// response parameters. If more results are available, the
// DescribeSnapshots.NextToken member contains a token that you pass in the next
// call to DescribeSnapshotsto retrieve the next set of items.
//
// You can also specify a maximum number of return results with the Limit
// parameter.
func directoryservice_DescribeSnapshots(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.DescribeSnapshotsInput{}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}
	if len(_directoryserviceLimit) > 0 {
		if err := assignInputField(input, "Limit", _directoryserviceLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_directoryserviceNextToken) > 0 {
		input.NextToken = aws.String(_directoryserviceNextToken)
	}
	if len(_directoryserviceSnapshotIds) > 0 {
		input.SnapshotIds = append([]string(nil), _directoryserviceSnapshotIds...)
	}

	if disablePaginator() {
		if resp, err := client.DescribeSnapshots(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*directoryservice.DescribeSnapshotsOutput
	p := directoryservice.NewDescribeSnapshotsPaginator(client, input)
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

// Obtains information about the trust relationships for this account.
// If no input parameters are provided, such as DirectoryId or TrustIds, this
// request describes all the trust relationships belonging to the account.
func directoryservice_DescribeTrusts(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.DescribeTrustsInput{}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}
	if len(_directoryserviceLimit) > 0 {
		if err := assignInputField(input, "Limit", _directoryserviceLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_directoryserviceNextToken) > 0 {
		input.NextToken = aws.String(_directoryserviceNextToken)
	}
	if len(_directoryserviceTrustIds) > 0 {
		input.TrustIds = append([]string(nil), _directoryserviceTrustIds...)
	}

	if disablePaginator() {
		if resp, err := client.DescribeTrusts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*directoryservice.DescribeTrustsOutput
	p := directoryservice.NewDescribeTrustsPaginator(client, input)
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

// Describes the updates of a directory for a particular update type.
func directoryservice_DescribeUpdateDirectory(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.DescribeUpdateDirectoryInput{
		// DirectoryId: *string, // Required
		// UpdateType: types.UpdateType, // Required
	}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}
	if len(_directoryserviceUpdateType) > 0 {
		if err := assignInputField(input, "UpdateType", _directoryserviceUpdateType); err != nil {
			log.Errorf("invalid --update-type: %s", err.Error())
			return
		}
	}
	if len(_directoryserviceNextToken) > 0 {
		input.NextToken = aws.String(_directoryserviceNextToken)
	}
	if len(_directoryserviceRegionName) > 0 {
		input.RegionName = aws.String(_directoryserviceRegionName)
	}

	if disablePaginator() {
		if resp, err := client.DescribeUpdateDirectory(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*directoryservice.DescribeUpdateDirectoryOutput
	p := directoryservice.NewDescribeUpdateDirectoryPaginator(client, input)
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

// Disables the certificate authority (CA) enrollment policy for the specified
// directory. This stops automatic certificate enrollment and management for
// domain-joined clients, but does not affect existing certificates.
//
// Disabling the CA enrollment policy prevents new certificates from being
// automatically enrolled, but existing certificates remain valid and functional
// until they expire.
func directoryservice_DisableCAEnrollmentPolicy(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.DisableCAEnrollmentPolicyInput{
		// DirectoryId: *string, // Required
	}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}

	if resp, err := client.DisableCAEnrollmentPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables alternative client authentication methods for the specified directory.
func directoryservice_DisableClientAuthentication(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.DisableClientAuthenticationInput{
		// DirectoryId: *string, // Required
		// Type: types.ClientAuthenticationType, // Required
	}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}
	if len(_directoryserviceType) > 0 {
		if err := assignInputField(input, "Type", _directoryserviceType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.DisableClientAuthentication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deactivates access to directory data via the Directory Service Data API for the
// specified directory. For more information, see [Directory Service Data API Reference].
//
// [Directory Service Data API Reference]: https://docs.aws.amazon.com/directoryservicedata/latest/DirectoryServiceDataAPIReference/Welcome.html
func directoryservice_DisableDirectoryDataAccess(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.DisableDirectoryDataAccessInput{
		// DirectoryId: *string, // Required
	}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}

	if resp, err := client.DisableDirectoryDataAccess(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deactivates LDAP secure calls for the specified directory.
func directoryservice_DisableLDAPS(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.DisableLDAPSInput{
		// DirectoryId: *string, // Required
		// Type: types.LDAPSType, // Required
	}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}
	if len(_directoryserviceType) > 0 {
		if err := assignInputField(input, "Type", _directoryserviceType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.DisableLDAPS(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables multi-factor authentication (MFA) with the Remote Authentication Dial
// In User Service (RADIUS) server for an AD Connector or Microsoft AD directory.
func directoryservice_DisableRadius(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.DisableRadiusInput{
		// DirectoryId: *string, // Required
	}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}

	if resp, err := client.DisableRadius(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables single-sign on for a directory.
func directoryservice_DisableSso(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.DisableSsoInput{
		// DirectoryId: *string, // Required
	}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}
	if len(_directoryservicePassword) > 0 {
		input.Password = aws.String(_directoryservicePassword)
	}
	if len(_directoryserviceUserName) > 0 {
		input.UserName = aws.String(_directoryserviceUserName)
	}

	if resp, err := client.DisableSso(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables certificate authority (CA) enrollment policy for the specified
// directory. This allows domain-joined clients to automatically request and
// receive certificates from the specified Amazon Web Services Private Certificate
// Authority.
//
// Before enabling CA enrollment, ensure that the PCA connector is properly
// configured and accessible from the directory. The connector must be in an active
// state and have the necessary permissions.
func directoryservice_EnableCAEnrollmentPolicy(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.EnableCAEnrollmentPolicyInput{
		// DirectoryId: *string, // Required
		// PcaConnectorArn: *string, // Required
	}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}
	if len(_directoryservicePcaConnectorArn) > 0 {
		input.PcaConnectorArn = aws.String(_directoryservicePcaConnectorArn)
	}

	if resp, err := client.EnableCAEnrollmentPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables alternative client authentication methods for the specified directory.
func directoryservice_EnableClientAuthentication(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.EnableClientAuthenticationInput{
		// DirectoryId: *string, // Required
		// Type: types.ClientAuthenticationType, // Required
	}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}
	if len(_directoryserviceType) > 0 {
		if err := assignInputField(input, "Type", _directoryserviceType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.EnableClientAuthentication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables access to directory data via the Directory Service Data API for the
// specified directory. For more information, see [Directory Service Data API Reference].
//
// [Directory Service Data API Reference]: https://docs.aws.amazon.com/directoryservicedata/latest/DirectoryServiceDataAPIReference/Welcome.html
func directoryservice_EnableDirectoryDataAccess(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.EnableDirectoryDataAccessInput{
		// DirectoryId: *string, // Required
	}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}

	if resp, err := client.EnableDirectoryDataAccess(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Activates the switch for the specific directory to always use LDAP secure calls.
func directoryservice_EnableLDAPS(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.EnableLDAPSInput{
		// DirectoryId: *string, // Required
		// Type: types.LDAPSType, // Required
	}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}
	if len(_directoryserviceType) > 0 {
		if err := assignInputField(input, "Type", _directoryserviceType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.EnableLDAPS(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables multi-factor authentication (MFA) with the Remote Authentication Dial
// In User Service (RADIUS) server for an AD Connector or Microsoft AD directory.
func directoryservice_EnableRadius(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.EnableRadiusInput{
		// DirectoryId: *string, // Required
		// RadiusSettings: *types.RadiusSettings, // Required
	}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}
	if len(_directoryserviceRadiusSettings) > 0 {
		if err := assignInputField(input, "RadiusSettings", _directoryserviceRadiusSettings); err != nil {
			log.Errorf("invalid --radius-settings: %s", err.Error())
			return
		}
	}

	if resp, err := client.EnableRadius(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables single sign-on for a directory. Single sign-on allows users in your
// directory to access certain Amazon Web Services services from a computer joined
// to the directory without having to enter their credentials separately.
func directoryservice_EnableSso(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.EnableSsoInput{
		// DirectoryId: *string, // Required
	}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}
	if len(_directoryservicePassword) > 0 {
		input.Password = aws.String(_directoryservicePassword)
	}
	if len(_directoryserviceUserName) > 0 {
		input.UserName = aws.String(_directoryserviceUserName)
	}

	if resp, err := client.EnableSso(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Obtains directory limit information for the current Region.
func directoryservice_GetDirectoryLimits(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.GetDirectoryLimitsInput{}

	if resp, err := client.GetDirectoryLimits(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Obtains the manual snapshot limits for a directory.
func directoryservice_GetSnapshotLimits(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.GetSnapshotLimitsInput{
		// DirectoryId: *string, // Required
	}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}

	if resp, err := client.GetSnapshotLimits(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of directory assessments for the specified directory or all
// assessments in your account. Use this operation to monitor assessment status and
// manage multiple assessments.
func directoryservice_ListADAssessments(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.ListADAssessmentsInput{}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}
	if len(_directoryserviceLimit) > 0 {
		if err := assignInputField(input, "Limit", _directoryserviceLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_directoryserviceNextToken) > 0 {
		input.NextToken = aws.String(_directoryserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListADAssessments(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*directoryservice.ListADAssessmentsOutput
	p := directoryservice.NewListADAssessmentsPaginator(client, input)
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

// For the specified directory, lists all the certificates registered for a secure
// LDAP or client certificate authentication.
func directoryservice_ListCertificates(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.ListCertificatesInput{
		// DirectoryId: *string, // Required
	}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}
	if len(_directoryserviceLimit) > 0 {
		if err := assignInputField(input, "Limit", _directoryserviceLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_directoryserviceNextToken) > 0 {
		input.NextToken = aws.String(_directoryserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCertificates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*directoryservice.ListCertificatesOutput
	p := directoryservice.NewListCertificatesPaginator(client, input)
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

// Lists the address blocks that you have added to a directory.
func directoryservice_ListIpRoutes(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.ListIpRoutesInput{
		// DirectoryId: *string, // Required
	}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}
	if len(_directoryserviceLimit) > 0 {
		if err := assignInputField(input, "Limit", _directoryserviceLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_directoryserviceNextToken) > 0 {
		input.NextToken = aws.String(_directoryserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListIpRoutes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*directoryservice.ListIpRoutesOutput
	p := directoryservice.NewListIpRoutesPaginator(client, input)
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

// Lists the active log subscriptions for the Amazon Web Services account.
func directoryservice_ListLogSubscriptions(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.ListLogSubscriptionsInput{}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}
	if len(_directoryserviceLimit) > 0 {
		if err := assignInputField(input, "Limit", _directoryserviceLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_directoryserviceNextToken) > 0 {
		input.NextToken = aws.String(_directoryserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListLogSubscriptions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*directoryservice.ListLogSubscriptionsOutput
	p := directoryservice.NewListLogSubscriptionsPaginator(client, input)
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

// Lists all schema extensions applied to a Microsoft AD Directory.
func directoryservice_ListSchemaExtensions(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.ListSchemaExtensionsInput{
		// DirectoryId: *string, // Required
	}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}
	if len(_directoryserviceLimit) > 0 {
		if err := assignInputField(input, "Limit", _directoryserviceLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_directoryserviceNextToken) > 0 {
		input.NextToken = aws.String(_directoryserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSchemaExtensions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*directoryservice.ListSchemaExtensionsOutput
	p := directoryservice.NewListSchemaExtensionsPaginator(client, input)
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

// Lists all tags on a directory.
func directoryservice_ListTagsForResource(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.ListTagsForResourceInput{
		// ResourceId: *string, // Required
	}

	if len(_directoryserviceResourceId) > 0 {
		input.ResourceId = aws.String(_directoryserviceResourceId)
	}
	if len(_directoryserviceLimit) > 0 {
		if err := assignInputField(input, "Limit", _directoryserviceLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_directoryserviceNextToken) > 0 {
		input.NextToken = aws.String(_directoryserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*directoryservice.ListTagsForResourceOutput
	p := directoryservice.NewListTagsForResourcePaginator(client, input)
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

// Registers a certificate for a secure LDAP or client certificate authentication.
func directoryservice_RegisterCertificate(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.RegisterCertificateInput{
		// CertificateData: *string, // Required
		// DirectoryId: *string, // Required
	}

	if len(_directoryserviceCertificateData) > 0 {
		input.CertificateData = aws.String(_directoryserviceCertificateData)
	}
	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}
	if len(_directoryserviceClientCertAuthSettings) > 0 {
		if err := assignInputField(input, "ClientCertAuthSettings", _directoryserviceClientCertAuthSettings); err != nil {
			log.Errorf("invalid --client-cert-auth-settings: %s", err.Error())
			return
		}
	}
	if len(_directoryserviceType) > 0 {
		if err := assignInputField(input, "Type", _directoryserviceType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.RegisterCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a directory with an Amazon SNS topic. This establishes the directory
// as a publisher to the specified Amazon SNS topic. You can then receive email or
// text (SMS) messages when the status of your directory changes. You get notified
// if your directory goes from an Active status to an Impaired or Inoperable
// status. You also receive a notification when the directory returns to an Active
// status.
func directoryservice_RegisterEventTopic(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.RegisterEventTopicInput{
		// DirectoryId: *string, // Required
		// TopicName: *string, // Required
	}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}
	if len(_directoryserviceTopicName) > 0 {
		input.TopicName = aws.String(_directoryserviceTopicName)
	}

	if resp, err := client.RegisterEventTopic(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Rejects a directory sharing request that was sent from the directory owner
// account.
func directoryservice_RejectSharedDirectory(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.RejectSharedDirectoryInput{
		// SharedDirectoryId: *string, // Required
	}

	if len(_directoryserviceSharedDirectoryId) > 0 {
		input.SharedDirectoryId = aws.String(_directoryserviceSharedDirectoryId)
	}

	if resp, err := client.RejectSharedDirectory(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes IP address blocks from a directory.
func directoryservice_RemoveIpRoutes(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.RemoveIpRoutesInput{
		// DirectoryId: *string, // Required
	}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}
	if len(_directoryserviceCidrIps) > 0 {
		input.CidrIps = append([]string(nil), _directoryserviceCidrIps...)
	}
	if len(_directoryserviceCidrIpv6s) > 0 {
		input.CidrIpv6s = append([]string(nil), _directoryserviceCidrIpv6s...)
	}

	if resp, err := client.RemoveIpRoutes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops all replication and removes the domain controllers from the specified
// Region. You cannot remove the primary Region with this operation. Instead, use
// the DeleteDirectory API.
func directoryservice_RemoveRegion(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.RemoveRegionInput{
		// DirectoryId: *string, // Required
	}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}

	if resp, err := client.RemoveRegion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes tags from a directory.
func directoryservice_RemoveTagsFromResource(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.RemoveTagsFromResourceInput{
		// ResourceId: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_directoryserviceResourceId) > 0 {
		input.ResourceId = aws.String(_directoryserviceResourceId)
	}
	if len(_directoryserviceTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _directoryserviceTagKeys...)
	}

	if resp, err := client.RemoveTagsFromResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Resets the password for any user in your Managed Microsoft AD or Simple AD
// directory. Disabled users will become enabled and can be authenticated following
// the API call.
//
// You can reset the password for any user in your directory with the following
// exceptions:
//
// - For Simple AD, you cannot reset the password for any user that is a member
// of either the Domain Admins or Enterprise Admins group except for the
// administrator user.
//
// - For Managed Microsoft AD, you can only reset the password for a user that
// is in an OU based off of the NetBIOS name that you typed when you created your
// directory. For example, you cannot reset the password for a user in the Amazon
// Web Services Reserved OU. For more information about the OU structure for an
// Managed Microsoft AD directory, see [What Gets Created]in the Directory Service Administration
// Guide.
//
// [What Gets Created]: https://docs.aws.amazon.com/directoryservice/latest/admin-guide/ms_ad_getting_started_what_gets_created.html
func directoryservice_ResetUserPassword(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.ResetUserPasswordInput{
		// DirectoryId: *string, // Required
		// NewPassword: *string, // Required
		// UserName: *string, // Required
	}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}
	if len(_directoryserviceNewPassword) > 0 {
		input.NewPassword = aws.String(_directoryserviceNewPassword)
	}
	if len(_directoryserviceUserName) > 0 {
		input.UserName = aws.String(_directoryserviceUserName)
	}

	if resp, err := client.ResetUserPassword(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Restores a directory using an existing directory snapshot.
// When you restore a directory from a snapshot, any changes made to the directory
// after the snapshot date are overwritten.
//
// This action returns as soon as the restore operation is initiated. You can
// monitor the progress of the restore operation by calling the DescribeDirectoriesoperation with the
// directory identifier. When the DirectoryDescription.Stage value changes to
// Active , the restore operation is complete.
func directoryservice_RestoreFromSnapshot(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.RestoreFromSnapshotInput{
		// SnapshotId: *string, // Required
	}

	if len(_directoryserviceSnapshotId) > 0 {
		input.SnapshotId = aws.String(_directoryserviceSnapshotId)
	}

	if resp, err := client.RestoreFromSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Shares a specified directory ( DirectoryId ) in your Amazon Web Services account
// (directory owner) with another Amazon Web Services account (directory consumer).
// With this operation you can use your directory from any Amazon Web Services
// account and from any Amazon VPC within an Amazon Web Services Region.
//
// When you share your Managed Microsoft AD directory, Directory Service creates a
// shared directory in the directory consumer account. This shared directory
// contains the metadata to provide access to the directory within the directory
// owner account. The shared directory is visible in all VPCs in the directory
// consumer account.
//
// The ShareMethod parameter determines whether the specified directory can be
// shared between Amazon Web Services accounts inside the same Amazon Web Services
// organization ( ORGANIZATIONS ). It also determines whether you can share the
// directory with any other Amazon Web Services account either inside or outside of
// the organization ( HANDSHAKE ).
//
// The ShareNotes parameter is only used when HANDSHAKE is called, which sends a
// directory sharing request to the directory consumer.
func directoryservice_ShareDirectory(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.ShareDirectoryInput{
		// DirectoryId: *string, // Required
		// ShareMethod: types.ShareMethod, // Required
		// ShareTarget: *types.ShareTarget, // Required
	}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}
	if len(_directoryserviceShareMethod) > 0 {
		if err := assignInputField(input, "ShareMethod", _directoryserviceShareMethod); err != nil {
			log.Errorf("invalid --share-method: %s", err.Error())
			return
		}
	}
	if len(_directoryserviceShareTarget) > 0 {
		if err := assignInputField(input, "ShareTarget", _directoryserviceShareTarget); err != nil {
			log.Errorf("invalid --share-target: %s", err.Error())
			return
		}
	}
	if len(_directoryserviceShareNotes) > 0 {
		input.ShareNotes = aws.String(_directoryserviceShareNotes)
	}

	if resp, err := client.ShareDirectory(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initiates a directory assessment to validate your self-managed AD environment
// for hybrid domain join. The assessment checks compatibility and connectivity of
// the self-managed AD environment.
//
// A directory assessment is automatically created when you create a hybrid
// directory. There are two types of assessments: CUSTOMER and SYSTEM . Your Amazon
// Web Services account has a limit of 100 CUSTOMER directory assessments.
//
// The assessment process typically takes 30 minutes or more to complete. The
// assessment process is asynchronous and you can monitor it with
// DescribeADAssessment .
//
// The InstanceIds must have a one-to-one correspondence with CustomerDnsIps ,
// meaning that if the IP address for instance i-10243410 is 10.24.34.100 and the
// IP address for instance i-10243420 is 10.24.34.200, then the input arrays must
// maintain the same order relationship, either [10.24.34.100, 10.24.34.200] paired
// with [i-10243410, i-10243420] or [10.24.34.200, 10.24.34.100] paired with
// [i-10243420, i-10243410].
//
// Note: You must provide exactly one DirectoryId or AssessmentConfiguration .
func directoryservice_StartADAssessment(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.StartADAssessmentInput{}

	if len(_directoryserviceAssessmentConfiguration) > 0 {
		if err := assignInputField(input, "AssessmentConfiguration", _directoryserviceAssessmentConfiguration); err != nil {
			log.Errorf("invalid --assessment-configuration: %s", err.Error())
			return
		}
	}
	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}

	if resp, err := client.StartADAssessment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Applies a schema extension to a Microsoft AD directory.
func directoryservice_StartSchemaExtension(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.StartSchemaExtensionInput{
		// CreateSnapshotBeforeSchemaExtension: bool, // Required
		// Description: *string, // Required
		// DirectoryId: *string, // Required
		// LdifContent: *string, // Required
	}

	if len(_directoryserviceCreateSnapshotBeforeSchemaExtension) > 0 {
		if err := assignInputField(input, "CreateSnapshotBeforeSchemaExtension", _directoryserviceCreateSnapshotBeforeSchemaExtension); err != nil {
			log.Errorf("invalid --create-snapshot-before-schema-extension: %s", err.Error())
			return
		}
	}
	if len(_directoryserviceDescription) > 0 {
		input.Description = aws.String(_directoryserviceDescription)
	}
	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}
	if len(_directoryserviceLdifContent) > 0 {
		input.LdifContent = aws.String(_directoryserviceLdifContent)
	}

	if resp, err := client.StartSchemaExtension(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops the directory sharing between the directory owner and consumer accounts.
func directoryservice_UnshareDirectory(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.UnshareDirectoryInput{
		// DirectoryId: *string, // Required
		// UnshareTarget: *types.UnshareTarget, // Required
	}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}
	if len(_directoryserviceUnshareTarget) > 0 {
		if err := assignInputField(input, "UnshareTarget", _directoryserviceUnshareTarget); err != nil {
			log.Errorf("invalid --unshare-target: %s", err.Error())
			return
		}
	}

	if resp, err := client.UnshareDirectory(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a conditional forwarder that has been set up for your Amazon Web
// Services directory.
func directoryservice_UpdateConditionalForwarder(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.UpdateConditionalForwarderInput{
		// DirectoryId: *string, // Required
		// RemoteDomainName: *string, // Required
	}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}
	if len(_directoryserviceRemoteDomainName) > 0 {
		input.RemoteDomainName = aws.String(_directoryserviceRemoteDomainName)
	}
	if len(_directoryserviceDnsIpAddrs) > 0 {
		input.DnsIpAddrs = append([]string(nil), _directoryserviceDnsIpAddrs...)
	}
	if len(_directoryserviceDnsIpv6Addrs) > 0 {
		input.DnsIpv6Addrs = append([]string(nil), _directoryserviceDnsIpv6Addrs...)
	}

	if resp, err := client.UpdateConditionalForwarder(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates directory configuration for the specified update type.
func directoryservice_UpdateDirectorySetup(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.UpdateDirectorySetupInput{
		// DirectoryId: *string, // Required
		// UpdateType: types.UpdateType, // Required
	}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}
	if len(_directoryserviceUpdateType) > 0 {
		if err := assignInputField(input, "UpdateType", _directoryserviceUpdateType); err != nil {
			log.Errorf("invalid --update-type: %s", err.Error())
			return
		}
	}
	if len(_directoryserviceCreateSnapshotBeforeUpdate) > 0 {
		if err := assignInputField(input, "CreateSnapshotBeforeUpdate", _directoryserviceCreateSnapshotBeforeUpdate); err != nil {
			log.Errorf("invalid --create-snapshot-before-update: %s", err.Error())
			return
		}
	}
	if len(_directoryserviceDirectorySizeUpdateSettings) > 0 {
		if err := assignInputField(input, "DirectorySizeUpdateSettings", _directoryserviceDirectorySizeUpdateSettings); err != nil {
			log.Errorf("invalid --directory-size-update-settings: %s", err.Error())
			return
		}
	}
	if len(_directoryserviceNetworkUpdateSettings) > 0 {
		if err := assignInputField(input, "NetworkUpdateSettings", _directoryserviceNetworkUpdateSettings); err != nil {
			log.Errorf("invalid --network-update-settings: %s", err.Error())
			return
		}
	}
	if len(_directoryserviceOSUpdateSettings) > 0 {
		if err := assignInputField(input, "OSUpdateSettings", _directoryserviceOSUpdateSettings); err != nil {
			log.Errorf("invalid --os-update-settings: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDirectorySetup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration of an existing hybrid directory. You can recover
// hybrid directory administrator account or modify self-managed instance settings.
//
// Updates are applied asynchronously. Use DescribeHybridADUpdate to monitor the progress of
// configuration changes.
//
// The InstanceIds must have a one-to-one correspondence with CustomerDnsIps ,
// meaning that if the IP address for instance i-10243410 is 10.24.34.100 and the
// IP address for instance i-10243420 is 10.24.34.200, then the input arrays must
// maintain the same order relationship, either [10.24.34.100, 10.24.34.200] paired
// with [i-10243410, i-10243420] or [10.24.34.200, 10.24.34.100] paired with
// [i-10243420, i-10243410].
//
// You must provide at least one update to UpdateHybridADRequest$HybridAdministratorAccountUpdate or UpdateHybridADRequest$SelfManagedInstancesSettings.
func directoryservice_UpdateHybridAD(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.UpdateHybridADInput{
		// DirectoryId: *string, // Required
	}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}
	if len(_directoryserviceHybridAdministratorAccountUpdate) > 0 {
		if err := assignInputField(input, "HybridAdministratorAccountUpdate", _directoryserviceHybridAdministratorAccountUpdate); err != nil {
			log.Errorf("invalid --hybrid-administrator-account-update: %s", err.Error())
			return
		}
	}
	if len(_directoryserviceSelfManagedInstancesSettings) > 0 {
		if err := assignInputField(input, "SelfManagedInstancesSettings", _directoryserviceSelfManagedInstancesSettings); err != nil {
			log.Errorf("invalid --self-managed-instances-settings: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateHybridAD(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or removes domain controllers to or from the directory. Based on the
// difference between current value and new value (provided through this API call),
// domain controllers will be added or removed. It may take up to 45 minutes for
// any new domain controllers to become fully active once the requested number of
// domain controllers is updated. During this time, you cannot make another update
// request.
func directoryservice_UpdateNumberOfDomainControllers(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.UpdateNumberOfDomainControllersInput{
		// DesiredNumber: *int32, // Required
		// DirectoryId: *string, // Required
	}

	if len(_directoryserviceDesiredNumber) > 0 {
		if err := assignInputField(input, "DesiredNumber", _directoryserviceDesiredNumber); err != nil {
			log.Errorf("invalid --desired-number: %s", err.Error())
			return
		}
	}
	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}

	if resp, err := client.UpdateNumberOfDomainControllers(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the Remote Authentication Dial In User Service (RADIUS) server
// information for an AD Connector or Microsoft AD directory.
func directoryservice_UpdateRadius(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.UpdateRadiusInput{
		// DirectoryId: *string, // Required
		// RadiusSettings: *types.RadiusSettings, // Required
	}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}
	if len(_directoryserviceRadiusSettings) > 0 {
		if err := assignInputField(input, "RadiusSettings", _directoryserviceRadiusSettings); err != nil {
			log.Errorf("invalid --radius-settings: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateRadius(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configurable settings for the specified directory.
func directoryservice_UpdateSettings(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.UpdateSettingsInput{
		// DirectoryId: *string, // Required
		// Settings: []types.Setting, // Required
	}

	if len(_directoryserviceDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryserviceDirectoryId)
	}
	if len(_directoryserviceSettings) > 0 {
		if err := assignInputField(input, "Settings", _directoryserviceSettings); err != nil {
			log.Errorf("invalid --settings: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the trust that has been set up between your Managed Microsoft AD
// directory and an self-managed Active Directory.
func directoryservice_UpdateTrust(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.UpdateTrustInput{
		// TrustId: *string, // Required
	}

	if len(_directoryserviceTrustId) > 0 {
		input.TrustId = aws.String(_directoryserviceTrustId)
	}
	if len(_directoryserviceSelectiveAuth) > 0 {
		if err := assignInputField(input, "SelectiveAuth", _directoryserviceSelectiveAuth); err != nil {
			log.Errorf("invalid --selective-auth: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateTrust(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Directory Service for Microsoft Active Directory allows you to configure and
// verify trust relationships.
//
// This action verifies a trust relationship between your Managed Microsoft AD
// directory and an external domain.
func directoryservice_VerifyTrust(cfg aws.Config, client *directoryservice.Client) {
	input := &directoryservice.VerifyTrustInput{
		// TrustId: *string, // Required
	}

	if len(_directoryserviceTrustId) > 0 {
		input.TrustId = aws.String(_directoryserviceTrustId)
	}

	if resp, err := client.VerifyTrust(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_directoryserviceCmd)
	_directoryserviceCmd.Flags().SortFlags = false

	_directoryserviceCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_directoryserviceCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_directoryserviceCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceAlias, "alias", "", "", "Alias")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceAssessmentConfiguration, "assessment-configuration", "", "", "Assessment Configuration")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceAssessmentId, "assessment-id", "", "", "Assessment ID")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceCertificateData, "certificate-data", "", "", "Certificate Data")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceCertificateId, "certificate-id", "", "", "Certificate ID")
	_directoryserviceCmd.Flags().StringSliceVarP(&_directoryserviceCidrIps, "cidr-ips", "", nil, "CIDR Ips")
	_directoryserviceCmd.Flags().StringSliceVarP(&_directoryserviceCidrIpv6s, "cidr-ipv6s", "", nil, "CIDR Ipv6s")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceClientCertAuthSettings, "client-cert-auth-settings", "", "", "Client Cert Auth Settings")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceComputerAttributes, "computer-attributes", "", "", "Computer Attributes")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceComputerName, "computer-name", "", "", "Computer Name")
	_directoryserviceCmd.Flags().StringSliceVarP(&_directoryserviceConditionalForwarderIpAddrs, "conditional-forwarder-ip-addrs", "", nil, "Conditional Forwarder IP Addrs")
	_directoryserviceCmd.Flags().StringSliceVarP(&_directoryserviceConditionalForwarderIpv6Addrs, "conditional-forwarder-ipv6-addrs", "", nil, "Conditional Forwarder IPV6 Addrs")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceConnectSettings, "connect-settings", "", "", "Connect Settings")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceCreateSnapshotBeforeSchemaExtension, "create-snapshot-before-schema-extension", "", "", "Create Snapshot Before Schema Extension")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceCreateSnapshotBeforeUpdate, "create-snapshot-before-update", "", "", "Create Snapshot Before Update")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceDeleteAssociatedConditionalForwarder, "delete-associated-conditional-forwarder", "", "", "Delete Associated Conditional Forwarder")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceDescription, "description", "", "", "Description")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceDesiredNumber, "desired-number", "", "", "Desired Number")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceDirectoryId, "directory-id", "", "", "Directory ID")
	_directoryserviceCmd.Flags().StringSliceVarP(&_directoryserviceDirectoryIds, "directory-ids", "", nil, "Directory Ids")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceDirectorySizeUpdateSettings, "directory-size-update-settings", "", "", "Directory Size Update Settings")
	_directoryserviceCmd.Flags().StringSliceVarP(&_directoryserviceDnsIpAddrs, "dns-ip-addrs", "", nil, "DNS IP Addrs")
	_directoryserviceCmd.Flags().StringSliceVarP(&_directoryserviceDnsIpv6Addrs, "dns-ipv6-addrs", "", nil, "DNS IPV6 Addrs")
	_directoryserviceCmd.Flags().StringSliceVarP(&_directoryserviceDomainControllerIds, "domain-controller-ids", "", nil, "Domain Controller Ids")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceEdition, "edition", "", "", "Edition")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceHybridAdministratorAccountUpdate, "hybrid-administrator-account-update", "", "", "Hybrid Administrator Account Update")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceIpRoutes, "ip-routes", "", "", "IP Routes")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceLdifContent, "ldif-content", "", "", "Ldif Content")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceLimit, "limit", "", "", "Limit")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceLogGroupName, "log-group-name", "", "", "Log Group Name")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceName, "name", "", "", "Name")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceNetworkType, "network-type", "", "", "Network Type")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceNetworkUpdateSettings, "network-update-settings", "", "", "Network Update Settings")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceNewPassword, "new-password", "", "", "New Password")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceNextToken, "next-token", "", "", "Next Token")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceOrganizationalUnitDistinguishedName, "organizational-unit-distinguished-name", "", "", "Organizational Unit Distinguished Name")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceOSUpdateSettings, "os-update-settings", "", "", "OS Update Settings")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceOwnerDirectoryId, "owner-directory-id", "", "", "Owner Directory ID")
	_directoryserviceCmd.Flags().StringVarP(&_directoryservicePassword, "password", "", "", "Password")
	_directoryserviceCmd.Flags().StringVarP(&_directoryservicePcaConnectorArn, "pca-connector-arn", "", "", "Pca Connector ARN")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceRadiusSettings, "radius-settings", "", "", "Radius Settings")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceRegionName, "region-name", "", "", "Region Name")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceRemoteDomainName, "remote-domain-name", "", "", "Remote Domain Name")
	_directoryserviceCmd.Flags().StringSliceVarP(&_directoryserviceRemoteDomainNames, "remote-domain-names", "", nil, "Remote Domain Names")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceResourceId, "resource-id", "", "", "Resource ID")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceSchemaExtensionId, "schema-extension-id", "", "", "Schema Extension ID")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceSecretArn, "secret-arn", "", "", "Secret ARN")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceSelectiveAuth, "selective-auth", "", "", "Selective Auth")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceSelfManagedInstancesSettings, "self-managed-instances-settings", "", "", "Self Managed Instances Settings")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceSettings, "settings", "", "", "Settings")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceShareMethod, "share-method", "", "", "Share Method")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceShareNotes, "share-notes", "", "", "Share Notes")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceShareTarget, "share-target", "", "", "Share Target")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceSharedDirectoryId, "shared-directory-id", "", "", "Shared Directory ID")
	_directoryserviceCmd.Flags().StringSliceVarP(&_directoryserviceSharedDirectoryIds, "shared-directory-ids", "", nil, "Shared Directory Ids")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceShortName, "short-name", "", "", "Short Name")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceSize, "size", "", "", "Size")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceSnapshotId, "snapshot-id", "", "", "Snapshot ID")
	_directoryserviceCmd.Flags().StringSliceVarP(&_directoryserviceSnapshotIds, "snapshot-ids", "", nil, "Snapshot Ids")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceStatus, "status", "", "", "Status")
	_directoryserviceCmd.Flags().StringSliceVarP(&_directoryserviceTagKeys, "tag-keys", "", nil, "Tag Keys")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceTags, "tags", "", "", "Tags")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceTopicName, "topic-name", "", "", "Topic Name")
	_directoryserviceCmd.Flags().StringSliceVarP(&_directoryserviceTopicNames, "topic-names", "", nil, "Topic Names")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceTrustDirection, "trust-direction", "", "", "Trust Direction")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceTrustId, "trust-id", "", "", "Trust ID")
	_directoryserviceCmd.Flags().StringSliceVarP(&_directoryserviceTrustIds, "trust-ids", "", nil, "Trust Ids")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceTrustPassword, "trust-password", "", "", "Trust Password")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceTrustType, "trust-type", "", "", "Trust Type")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceType, "type", "", "", "Type")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceUnshareTarget, "unshare-target", "", "", "Unshare Target")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceUpdateSecurityGroupForDirectoryControllers, "update-security-group-for-directory-controllers", "", "", "Update Security Group For Directory Controllers")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceUpdateType, "update-type", "", "", "Update Type")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceUserName, "user-name", "", "", "User Name")
	_directoryserviceCmd.Flags().StringVarP(&_directoryserviceVPCSettings, "vpc-settings", "", "", "VPC Settings")

	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceAcceptSharedDirectory, "accept-shared-directory", "", false, "Accept Shared Directory")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceAddIpRoutes, "add-ip-routes", "", false, "Add IP Routes")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceAddRegion, "add-region", "", false, "Add Region")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceAddTagsToResource, "add-tags-to-resource", "", false, "Add Tags To Resource")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceCancelSchemaExtension, "cancel-schema-extension", "", false, "Cancel Schema Extension")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceConnectDirectory, "connect-directory", "", false, "Connect Directory")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceCreateAlias, "create-alias", "", false, "Create Alias")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceCreateComputer, "create-computer", "", false, "Create Computer")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceCreateConditionalForwarder, "create-conditional-forwarder", "", false, "Create Conditional Forwarder")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceCreateDirectory, "create-directory", "", false, "Create Directory")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceCreateHybridAD, "create-hybrid-ad", "", false, "Create Hybrid Ad")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceCreateLogSubscription, "create-log-subscription", "", false, "Create Log Subscription")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceCreateMicrosoftAD, "create-microsoft-ad", "", false, "Create Microsoft Ad")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceCreateSnapshot, "create-snapshot", "", false, "Create Snapshot")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceCreateTrust, "create-trust", "", false, "Create Trust")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceDeleteADAssessment, "delete-ad-assessment", "", false, "Delete Ad Assessment")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceDeleteConditionalForwarder, "delete-conditional-forwarder", "", false, "Delete Conditional Forwarder")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceDeleteDirectory, "delete-directory", "", false, "Delete Directory")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceDeleteLogSubscription, "delete-log-subscription", "", false, "Delete Log Subscription")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceDeleteSnapshot, "delete-snapshot", "", false, "Delete Snapshot")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceDeleteTrust, "delete-trust", "", false, "Delete Trust")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceDeregisterCertificate, "deregister-certificate", "", false, "Deregister Certificate")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceDeregisterEventTopic, "deregister-event-topic", "", false, "Deregister Event Topic")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceDescribeADAssessment, "describe-ad-assessment", "", false, "Describe Ad Assessment")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceDescribeCAEnrollmentPolicy, "describe-ca-enrollment-policy", "", false, "Describe Ca Enrollment Policy")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceDescribeCertificate, "describe-certificate", "", false, "Describe Certificate")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceDescribeClientAuthenticationSettings, "describe-client-authentication-settings", "", false, "Describe Client Authentication Settings")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceDescribeConditionalForwarders, "describe-conditional-forwarders", "", false, "Describe Conditional Forwarders")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceDescribeDirectories, "describe-directories", "", false, "Describe Directories")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceDescribeDirectoryDataAccess, "describe-directory-data-access", "", false, "Describe Directory Data Access")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceDescribeDomainControllers, "describe-domain-controllers", "", false, "Describe Domain Controllers")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceDescribeEventTopics, "describe-event-topics", "", false, "Describe Event Topics")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceDescribeHybridADUpdate, "describe-hybrid-ad-update", "", false, "Describe Hybrid Ad Update")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceDescribeLDAPSSettings, "describe-ldaps-settings", "", false, "Describe Ldaps Settings")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceDescribeRegions, "describe-regions", "", false, "Describe Regions")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceDescribeSettings, "describe-settings", "", false, "Describe Settings")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceDescribeSharedDirectories, "describe-shared-directories", "", false, "Describe Shared Directories")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceDescribeSnapshots, "describe-snapshots", "", false, "Describe Snapshots")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceDescribeTrusts, "describe-trusts", "", false, "Describe Trusts")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceDescribeUpdateDirectory, "describe-update-directory", "", false, "Describe Update Directory")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceDisableCAEnrollmentPolicy, "disable-ca-enrollment-policy", "", false, "Disable Ca Enrollment Policy")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceDisableClientAuthentication, "disable-client-authentication", "", false, "Disable Client Authentication")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceDisableDirectoryDataAccess, "disable-directory-data-access", "", false, "Disable Directory Data Access")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceDisableLDAPS, "disable-ldaps", "", false, "Disable Ldaps")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceDisableRadius, "disable-radius", "", false, "Disable Radius")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceDisableSso, "disable-sso", "", false, "Disable Sso")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceEnableCAEnrollmentPolicy, "enable-ca-enrollment-policy", "", false, "Enable Ca Enrollment Policy")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceEnableClientAuthentication, "enable-client-authentication", "", false, "Enable Client Authentication")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceEnableDirectoryDataAccess, "enable-directory-data-access", "", false, "Enable Directory Data Access")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceEnableLDAPS, "enable-ldaps", "", false, "Enable Ldaps")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceEnableRadius, "enable-radius", "", false, "Enable Radius")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceEnableSso, "enable-sso", "", false, "Enable Sso")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceGetDirectoryLimits, "get-directory-limits", "", false, "Get Directory Limits")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceGetSnapshotLimits, "get-snapshot-limits", "", false, "Get Snapshot Limits")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceListADAssessments, "list-ad-assessments", "", false, "List Ad Assessments")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceListCertificates, "list-certificates", "", false, "List Certificates")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceListIpRoutes, "list-ip-routes", "", false, "List IP Routes")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceListLogSubscriptions, "list-log-subscriptions", "", false, "List Log Subscriptions")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceListSchemaExtensions, "list-schema-extensions", "", false, "List Schema Extensions")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceRegisterCertificate, "register-certificate", "", false, "Register Certificate")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceRegisterEventTopic, "register-event-topic", "", false, "Register Event Topic")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceRejectSharedDirectory, "reject-shared-directory", "", false, "Reject Shared Directory")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceRemoveIpRoutes, "remove-ip-routes", "", false, "Remove IP Routes")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceRemoveRegion, "remove-region", "", false, "Remove Region")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceRemoveTagsFromResource, "remove-tags-from-resource", "", false, "Remove Tags From Resource")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceResetUserPassword, "reset-user-password", "", false, "Reset User Password")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceRestoreFromSnapshot, "restore-from-snapshot", "", false, "Restore From Snapshot")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceShareDirectory, "share-directory", "", false, "Share Directory")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceStartADAssessment, "start-ad-assessment", "", false, "Start Ad Assessment")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceStartSchemaExtension, "start-schema-extension", "", false, "Start Schema Extension")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceUnshareDirectory, "unshare-directory", "", false, "Unshare Directory")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceUpdateConditionalForwarder, "update-conditional-forwarder", "", false, "Update Conditional Forwarder")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceUpdateDirectorySetup, "update-directory-setup", "", false, "Update Directory Setup")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceUpdateHybridAD, "update-hybrid-ad", "", false, "Update Hybrid Ad")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceUpdateNumberOfDomainControllers, "update-number-of-domain-controllers", "", false, "Update Number Of Domain Controllers")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceUpdateRadius, "update-radius", "", false, "Update Radius")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceUpdateSettings, "update-settings", "", false, "Update Settings")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceUpdateTrust, "update-trust", "", false, "Update Trust")
	_directoryserviceCmd.Flags().BoolVarP(&_directoryserviceVerifyTrust, "verify-trust", "", false, "Verify Trust")

}
