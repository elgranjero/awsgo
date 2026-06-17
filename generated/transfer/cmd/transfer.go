package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/transfer"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// transferCmd represents the transfer command
var _transferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "AWS transfer CLI",
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
		client := transfer.NewFromConfig(cfg)
		if _transferCreateAccess {
			transfer_CreateAccess(cfg, client)
			return
		}
		if _transferCreateAgreement {
			transfer_CreateAgreement(cfg, client)
			return
		}
		if _transferCreateConnector {
			transfer_CreateConnector(cfg, client)
			return
		}
		if _transferCreateProfile {
			transfer_CreateProfile(cfg, client)
			return
		}
		if _transferCreateServer {
			transfer_CreateServer(cfg, client)
			return
		}
		if _transferCreateUser {
			transfer_CreateUser(cfg, client)
			return
		}
		if _transferCreateWebApp {
			transfer_CreateWebApp(cfg, client)
			return
		}
		if _transferCreateWorkflow {
			transfer_CreateWorkflow(cfg, client)
			return
		}
		if _transferDeleteAccess {
			transfer_DeleteAccess(cfg, client)
			return
		}
		if _transferDeleteAgreement {
			transfer_DeleteAgreement(cfg, client)
			return
		}
		if _transferDeleteCertificate {
			transfer_DeleteCertificate(cfg, client)
			return
		}
		if _transferDeleteConnector {
			transfer_DeleteConnector(cfg, client)
			return
		}
		if _transferDeleteHostKey {
			transfer_DeleteHostKey(cfg, client)
			return
		}
		if _transferDeleteProfile {
			transfer_DeleteProfile(cfg, client)
			return
		}
		if _transferDeleteServer {
			transfer_DeleteServer(cfg, client)
			return
		}
		if _transferDeleteSshPublicKey {
			transfer_DeleteSshPublicKey(cfg, client)
			return
		}
		if _transferDeleteUser {
			transfer_DeleteUser(cfg, client)
			return
		}
		if _transferDeleteWebApp {
			transfer_DeleteWebApp(cfg, client)
			return
		}
		if _transferDeleteWebAppCustomization {
			transfer_DeleteWebAppCustomization(cfg, client)
			return
		}
		if _transferDeleteWorkflow {
			transfer_DeleteWorkflow(cfg, client)
			return
		}
		if _transferDescribeAccess {
			transfer_DescribeAccess(cfg, client)
			return
		}
		if _transferDescribeAgreement {
			transfer_DescribeAgreement(cfg, client)
			return
		}
		if _transferDescribeCertificate {
			transfer_DescribeCertificate(cfg, client)
			return
		}
		if _transferDescribeConnector {
			transfer_DescribeConnector(cfg, client)
			return
		}
		if _transferDescribeExecution {
			transfer_DescribeExecution(cfg, client)
			return
		}
		if _transferDescribeHostKey {
			transfer_DescribeHostKey(cfg, client)
			return
		}
		if _transferDescribeProfile {
			transfer_DescribeProfile(cfg, client)
			return
		}
		if _transferDescribeSecurityPolicy {
			transfer_DescribeSecurityPolicy(cfg, client)
			return
		}
		if _transferDescribeServer {
			transfer_DescribeServer(cfg, client)
			return
		}
		if _transferDescribeUser {
			transfer_DescribeUser(cfg, client)
			return
		}
		if _transferDescribeWebApp {
			transfer_DescribeWebApp(cfg, client)
			return
		}
		if _transferDescribeWebAppCustomization {
			transfer_DescribeWebAppCustomization(cfg, client)
			return
		}
		if _transferDescribeWorkflow {
			transfer_DescribeWorkflow(cfg, client)
			return
		}
		if _transferImportCertificate {
			transfer_ImportCertificate(cfg, client)
			return
		}
		if _transferImportHostKey {
			transfer_ImportHostKey(cfg, client)
			return
		}
		if _transferImportSshPublicKey {
			transfer_ImportSshPublicKey(cfg, client)
			return
		}
		if _transferListAccesses {
			transfer_ListAccesses(cfg, client)
			return
		}
		if _transferListAgreements {
			transfer_ListAgreements(cfg, client)
			return
		}
		if _transferListCertificates {
			transfer_ListCertificates(cfg, client)
			return
		}
		if _transferListConnectors {
			transfer_ListConnectors(cfg, client)
			return
		}
		if _transferListExecutions {
			transfer_ListExecutions(cfg, client)
			return
		}
		if _transferListFileTransferResults {
			transfer_ListFileTransferResults(cfg, client)
			return
		}
		if _transferListHostKeys {
			transfer_ListHostKeys(cfg, client)
			return
		}
		if _transferListProfiles {
			transfer_ListProfiles(cfg, client)
			return
		}
		if _transferListSecurityPolicies {
			transfer_ListSecurityPolicies(cfg, client)
			return
		}
		if _transferListServers {
			transfer_ListServers(cfg, client)
			return
		}
		if _transferListTagsForResource {
			transfer_ListTagsForResource(cfg, client)
			return
		}
		if _transferListUsers {
			transfer_ListUsers(cfg, client)
			return
		}
		if _transferListWebApps {
			transfer_ListWebApps(cfg, client)
			return
		}
		if _transferListWorkflows {
			transfer_ListWorkflows(cfg, client)
			return
		}
		if _transferSendWorkflowStepState {
			transfer_SendWorkflowStepState(cfg, client)
			return
		}
		if _transferStartDirectoryListing {
			transfer_StartDirectoryListing(cfg, client)
			return
		}
		if _transferStartFileTransfer {
			transfer_StartFileTransfer(cfg, client)
			return
		}
		if _transferStartRemoteDelete {
			transfer_StartRemoteDelete(cfg, client)
			return
		}
		if _transferStartRemoteMove {
			transfer_StartRemoteMove(cfg, client)
			return
		}
		if _transferStartServer {
			transfer_StartServer(cfg, client)
			return
		}
		if _transferStopServer {
			transfer_StopServer(cfg, client)
			return
		}
		if _transferTagResource {
			transfer_TagResource(cfg, client)
			return
		}
		if _transferTestConnection {
			transfer_TestConnection(cfg, client)
			return
		}
		if _transferTestIdentityProvider {
			transfer_TestIdentityProvider(cfg, client)
			return
		}
		if _transferUntagResource {
			transfer_UntagResource(cfg, client)
			return
		}
		if _transferUpdateAccess {
			transfer_UpdateAccess(cfg, client)
			return
		}
		if _transferUpdateAgreement {
			transfer_UpdateAgreement(cfg, client)
			return
		}
		if _transferUpdateCertificate {
			transfer_UpdateCertificate(cfg, client)
			return
		}
		if _transferUpdateConnector {
			transfer_UpdateConnector(cfg, client)
			return
		}
		if _transferUpdateHostKey {
			transfer_UpdateHostKey(cfg, client)
			return
		}
		if _transferUpdateProfile {
			transfer_UpdateProfile(cfg, client)
			return
		}
		if _transferUpdateServer {
			transfer_UpdateServer(cfg, client)
			return
		}
		if _transferUpdateUser {
			transfer_UpdateUser(cfg, client)
			return
		}
		if _transferUpdateWebApp {
			transfer_UpdateWebApp(cfg, client)
			return
		}
		if _transferUpdateWebAppCustomization {
			transfer_UpdateWebAppCustomization(cfg, client)
			return
		}

	},
}

var (
	_transferCreateAccess                bool
	_transferCreateAgreement             bool
	_transferCreateConnector             bool
	_transferCreateProfile               bool
	_transferCreateServer                bool
	_transferCreateUser                  bool
	_transferCreateWebApp                bool
	_transferCreateWorkflow              bool
	_transferDeleteAccess                bool
	_transferDeleteAgreement             bool
	_transferDeleteCertificate           bool
	_transferDeleteConnector             bool
	_transferDeleteHostKey               bool
	_transferDeleteProfile               bool
	_transferDeleteServer                bool
	_transferDeleteSshPublicKey          bool
	_transferDeleteUser                  bool
	_transferDeleteWebApp                bool
	_transferDeleteWebAppCustomization   bool
	_transferDeleteWorkflow              bool
	_transferDescribeAccess              bool
	_transferDescribeAgreement           bool
	_transferDescribeCertificate         bool
	_transferDescribeConnector           bool
	_transferDescribeExecution           bool
	_transferDescribeHostKey             bool
	_transferDescribeProfile             bool
	_transferDescribeSecurityPolicy      bool
	_transferDescribeServer              bool
	_transferDescribeUser                bool
	_transferDescribeWebApp              bool
	_transferDescribeWebAppCustomization bool
	_transferDescribeWorkflow            bool
	_transferImportCertificate           bool
	_transferImportHostKey               bool
	_transferImportSshPublicKey          bool
	_transferListAccesses                bool
	_transferListAgreements              bool
	_transferListCertificates            bool
	_transferListConnectors              bool
	_transferListExecutions              bool
	_transferListFileTransferResults     bool
	_transferListHostKeys                bool
	_transferListProfiles                bool
	_transferListSecurityPolicies        bool
	_transferListServers                 bool
	_transferListTagsForResource         bool
	_transferListUsers                   bool
	_transferListWebApps                 bool
	_transferListWorkflows               bool
	_transferSendWorkflowStepState       bool
	_transferStartDirectoryListing       bool
	_transferStartFileTransfer           bool
	_transferStartRemoteDelete           bool
	_transferStartRemoteMove             bool
	_transferStartServer                 bool
	_transferStopServer                  bool
	_transferTagResource                 bool
	_transferTestConnection              bool
	_transferTestIdentityProvider        bool
	_transferUntagResource               bool
	_transferUpdateAccess                bool
	_transferUpdateAgreement             bool
	_transferUpdateCertificate           bool
	_transferUpdateConnector             bool
	_transferUpdateHostKey               bool
	_transferUpdateProfile               bool
	_transferUpdateServer                bool
	_transferUpdateUser                  bool
	_transferUpdateWebApp                bool
	_transferUpdateWebAppCustomization   bool

	_transferAccessEndpoint                string
	_transferAccessRole                    string
	_transferActiveDate                    string
	_transferAgreementId                   string
	_transferArn                           string
	_transferAs2Config                     string
	_transferAs2Id                         string
	_transferBaseDirectory                 string
	_transferCertificate                   string
	_transferCertificateChain              string
	_transferCertificateId                 string
	_transferCertificateIds                []string
	_transferConnectorId                   string
	_transferCustomDirectories             string
	_transferCustomHttpHeaders             string
	_transferDeletePath                    string
	_transferDescription                   string
	_transferDomain                        string
	_transferEgressConfig                  string
	_transferEndpointDetails               string
	_transferEndpointType                  string
	_transferEnforceMessageSigning         string
	_transferExecutionId                   string
	_transferExternalId                    string
	_transferFaviconFile                   string
	_transferHomeDirectory                 string
	_transferHomeDirectoryMappings         string
	_transferHomeDirectoryType             string
	_transferHostKey                       string
	_transferHostKeyBody                   string
	_transferHostKeyId                     string
	_transferIdentityProviderDetails       string
	_transferIdentityProviderType          string
	_transferInactiveDate                  string
	_transferIpAddressType                 string
	_transferLocalDirectoryPath            string
	_transferLocalProfileId                string
	_transferLoggingRole                   string
	_transferLogoFile                      string
	_transferMaxItems                      string
	_transferMaxResults                    string
	_transferNextToken                     string
	_transferOnExceptionSteps              string
	_transferOutputDirectoryPath           string
	_transferPartnerProfileId              string
	_transferPolicy                        string
	_transferPosixProfile                  string
	_transferPostAuthenticationLoginBanner string
	_transferPreAuthenticationLoginBanner  string
	_transferPreserveFilename              string
	_transferPrivateKey                    string
	_transferProfileId                     string
	_transferProfileType                   string
	_transferProtocolDetails               string
	_transferProtocols                     string
	_transferRemoteDirectoryPath           string
	_transferRetrieveFilePaths             []string
	_transferRole                          string
	_transferS3StorageOptions              string
	_transferSecurityPolicyName            string
	_transferSendFilePaths                 []string
	_transferServerId                      string
	_transferServerProtocol                string
	_transferSftpConfig                    string
	_transferSourceIp                      string
	_transferSourcePath                    string
	_transferSshPublicKeyBody              string
	_transferSshPublicKeyId                string
	_transferStatus                        string
	_transferSteps                         string
	_transferStructuredLogDestinations     []string
	_transferTagKeys                       []string
	_transferTags                          string
	_transferTargetPath                    string
	_transferTitle                         string
	_transferToken                         string
	_transferTransferId                    string
	_transferUrl                           string
	_transferUsage                         string
	_transferUserName                      string
	_transferUserPassword                  string
	_transferWebAppEndpointPolicy          string
	_transferWebAppId                      string
	_transferWebAppUnits                   string
	_transferWorkflowDetails               string
	_transferWorkflowId                    string
)

// Used by administrators to choose which groups in the directory should have
// access to upload and download files over the enabled protocols using Transfer
// Family. For example, a Microsoft Active Directory might contain 50,000 users,
// but only a small fraction might need the ability to transfer files to the
// server. An administrator can use CreateAccess to limit the access to the
// correct set of users who need this ability.
func transfer_CreateAccess(cfg aws.Config, client *transfer.Client) {
	input := &transfer.CreateAccessInput{
		// ExternalId: *string, // Required
		// Role: *string, // Required
		// ServerId: *string, // Required
	}

	if len(_transferExternalId) > 0 {
		input.ExternalId = aws.String(_transferExternalId)
	}
	if len(_transferRole) > 0 {
		input.Role = aws.String(_transferRole)
	}
	if len(_transferServerId) > 0 {
		input.ServerId = aws.String(_transferServerId)
	}
	if len(_transferHomeDirectory) > 0 {
		input.HomeDirectory = aws.String(_transferHomeDirectory)
	}
	if len(_transferHomeDirectoryMappings) > 0 {
		if err := assignInputField(input, "HomeDirectoryMappings", _transferHomeDirectoryMappings); err != nil {
			log.Errorf("invalid --home-directory-mappings: %s", err.Error())
			return
		}
	}
	if len(_transferHomeDirectoryType) > 0 {
		if err := assignInputField(input, "HomeDirectoryType", _transferHomeDirectoryType); err != nil {
			log.Errorf("invalid --home-directory-type: %s", err.Error())
			return
		}
	}
	if len(_transferPolicy) > 0 {
		input.Policy = aws.String(_transferPolicy)
	}
	if len(_transferPosixProfile) > 0 {
		if err := assignInputField(input, "PosixProfile", _transferPosixProfile); err != nil {
			log.Errorf("invalid --posix-profile: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAccess(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an agreement. An agreement is a bilateral trading partner agreement, or
// partnership, between an Transfer Family server and an AS2 process. The agreement
// defines the file and message transfer relationship between the server and the
// AS2 process. To define an agreement, Transfer Family combines a server, local
// profile, partner profile, certificate, and other attributes.
//
// The partner is identified with the PartnerProfileId , and the AS2 process is
// identified with the LocalProfileId .
//
// Specify either BaseDirectory or CustomDirectories , but not both. Specifying
// both causes the command to fail.
func transfer_CreateAgreement(cfg aws.Config, client *transfer.Client) {
	input := &transfer.CreateAgreementInput{
		// AccessRole: *string, // Required
		// LocalProfileId: *string, // Required
		// PartnerProfileId: *string, // Required
		// ServerId: *string, // Required
	}

	if len(_transferAccessRole) > 0 {
		input.AccessRole = aws.String(_transferAccessRole)
	}
	if len(_transferLocalProfileId) > 0 {
		input.LocalProfileId = aws.String(_transferLocalProfileId)
	}
	if len(_transferPartnerProfileId) > 0 {
		input.PartnerProfileId = aws.String(_transferPartnerProfileId)
	}
	if len(_transferServerId) > 0 {
		input.ServerId = aws.String(_transferServerId)
	}
	if len(_transferBaseDirectory) > 0 {
		input.BaseDirectory = aws.String(_transferBaseDirectory)
	}
	if len(_transferCustomDirectories) > 0 {
		if err := assignInputField(input, "CustomDirectories", _transferCustomDirectories); err != nil {
			log.Errorf("invalid --custom-directories: %s", err.Error())
			return
		}
	}
	if len(_transferDescription) > 0 {
		input.Description = aws.String(_transferDescription)
	}
	if len(_transferEnforceMessageSigning) > 0 {
		if err := assignInputField(input, "EnforceMessageSigning", _transferEnforceMessageSigning); err != nil {
			log.Errorf("invalid --enforce-message-signing: %s", err.Error())
			return
		}
	}
	if len(_transferPreserveFilename) > 0 {
		if err := assignInputField(input, "PreserveFilename", _transferPreserveFilename); err != nil {
			log.Errorf("invalid --preserve-filename: %s", err.Error())
			return
		}
	}
	if len(_transferStatus) > 0 {
		if err := assignInputField(input, "Status", _transferStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_transferTags) > 0 {
		if err := assignInputField(input, "Tags", _transferTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAgreement(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates the connector, which captures the parameters for a connection for the
// AS2 or SFTP protocol. For AS2, the connector is required for sending files to an
// externally hosted AS2 server. For SFTP, the connector is required when sending
// files to an SFTP server or receiving files from an SFTP server. For more details
// about connectors, see [Configure AS2 connectors]and [Create SFTP connectors].
//
// You must specify exactly one configuration object: either for AS2 ( As2Config )
// or SFTP ( SftpConfig ).
//
// [Configure AS2 connectors]: https://docs.aws.amazon.com/transfer/latest/userguide/configure-as2-connector.html
// [Create SFTP connectors]: https://docs.aws.amazon.com/transfer/latest/userguide/configure-sftp-connector.html
func transfer_CreateConnector(cfg aws.Config, client *transfer.Client) {
	input := &transfer.CreateConnectorInput{
		// AccessRole: *string, // Required
	}

	if len(_transferAccessRole) > 0 {
		input.AccessRole = aws.String(_transferAccessRole)
	}
	if len(_transferAs2Config) > 0 {
		if err := assignInputField(input, "As2Config", _transferAs2Config); err != nil {
			log.Errorf("invalid --as2-config: %s", err.Error())
			return
		}
	}
	if len(_transferEgressConfig) > 0 {
		if err := assignInputField(input, "EgressConfig", _transferEgressConfig); err != nil {
			log.Errorf("invalid --egress-config: %s", err.Error())
			return
		}
	}
	if len(_transferLoggingRole) > 0 {
		input.LoggingRole = aws.String(_transferLoggingRole)
	}
	if len(_transferSecurityPolicyName) > 0 {
		input.SecurityPolicyName = aws.String(_transferSecurityPolicyName)
	}
	if len(_transferSftpConfig) > 0 {
		if err := assignInputField(input, "SftpConfig", _transferSftpConfig); err != nil {
			log.Errorf("invalid --sftp-config: %s", err.Error())
			return
		}
	}
	if len(_transferTags) > 0 {
		if err := assignInputField(input, "Tags", _transferTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_transferUrl) > 0 {
		input.Url = aws.String(_transferUrl)
	}

	if resp, err := client.CreateConnector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates the local or partner profile to use for AS2 transfers.
func transfer_CreateProfile(cfg aws.Config, client *transfer.Client) {
	input := &transfer.CreateProfileInput{
		// As2Id: *string, // Required
		// ProfileType: types.ProfileType, // Required
	}

	if len(_transferAs2Id) > 0 {
		input.As2Id = aws.String(_transferAs2Id)
	}
	if len(_transferProfileType) > 0 {
		if err := assignInputField(input, "ProfileType", _transferProfileType); err != nil {
			log.Errorf("invalid --profile-type: %s", err.Error())
			return
		}
	}
	if len(_transferCertificateIds) > 0 {
		input.CertificateIds = append([]string(nil), _transferCertificateIds...)
	}
	if len(_transferTags) > 0 {
		if err := assignInputField(input, "Tags", _transferTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Instantiates an auto-scaling virtual server based on the selected file transfer
// protocol in Amazon Web Services. When you make updates to your file transfer
// protocol-enabled server or when you work with users, use the service-generated
// ServerId property that is assigned to the newly created server.
func transfer_CreateServer(cfg aws.Config, client *transfer.Client) {
	input := &transfer.CreateServerInput{}

	if len(_transferCertificate) > 0 {
		input.Certificate = aws.String(_transferCertificate)
	}
	if len(_transferDomain) > 0 {
		if err := assignInputField(input, "Domain", _transferDomain); err != nil {
			log.Errorf("invalid --domain: %s", err.Error())
			return
		}
	}
	if len(_transferEndpointDetails) > 0 {
		if err := assignInputField(input, "EndpointDetails", _transferEndpointDetails); err != nil {
			log.Errorf("invalid --endpoint-details: %s", err.Error())
			return
		}
	}
	if len(_transferEndpointType) > 0 {
		if err := assignInputField(input, "EndpointType", _transferEndpointType); err != nil {
			log.Errorf("invalid --endpoint-type: %s", err.Error())
			return
		}
	}
	if len(_transferHostKey) > 0 {
		input.HostKey = aws.String(_transferHostKey)
	}
	if len(_transferIdentityProviderDetails) > 0 {
		if err := assignInputField(input, "IdentityProviderDetails", _transferIdentityProviderDetails); err != nil {
			log.Errorf("invalid --identity-provider-details: %s", err.Error())
			return
		}
	}
	if len(_transferIdentityProviderType) > 0 {
		if err := assignInputField(input, "IdentityProviderType", _transferIdentityProviderType); err != nil {
			log.Errorf("invalid --identity-provider-type: %s", err.Error())
			return
		}
	}
	if len(_transferIpAddressType) > 0 {
		if err := assignInputField(input, "IpAddressType", _transferIpAddressType); err != nil {
			log.Errorf("invalid --ip-address-type: %s", err.Error())
			return
		}
	}
	if len(_transferLoggingRole) > 0 {
		input.LoggingRole = aws.String(_transferLoggingRole)
	}
	if len(_transferPostAuthenticationLoginBanner) > 0 {
		input.PostAuthenticationLoginBanner = aws.String(_transferPostAuthenticationLoginBanner)
	}
	if len(_transferPreAuthenticationLoginBanner) > 0 {
		input.PreAuthenticationLoginBanner = aws.String(_transferPreAuthenticationLoginBanner)
	}
	if len(_transferProtocolDetails) > 0 {
		if err := assignInputField(input, "ProtocolDetails", _transferProtocolDetails); err != nil {
			log.Errorf("invalid --protocol-details: %s", err.Error())
			return
		}
	}
	if len(_transferProtocols) > 0 {
		if err := assignInputField(input, "Protocols", _transferProtocols); err != nil {
			log.Errorf("invalid --protocols: %s", err.Error())
			return
		}
	}
	if len(_transferS3StorageOptions) > 0 {
		if err := assignInputField(input, "S3StorageOptions", _transferS3StorageOptions); err != nil {
			log.Errorf("invalid --s3-storage-options: %s", err.Error())
			return
		}
	}
	if len(_transferSecurityPolicyName) > 0 {
		input.SecurityPolicyName = aws.String(_transferSecurityPolicyName)
	}
	if len(_transferStructuredLogDestinations) > 0 {
		input.StructuredLogDestinations = append([]string(nil), _transferStructuredLogDestinations...)
	}
	if len(_transferTags) > 0 {
		if err := assignInputField(input, "Tags", _transferTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_transferWorkflowDetails) > 0 {
		if err := assignInputField(input, "WorkflowDetails", _transferWorkflowDetails); err != nil {
			log.Errorf("invalid --workflow-details: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateServer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a user and associates them with an existing file transfer
// protocol-enabled server. You can only create and associate users with servers
// that have the IdentityProviderType set to SERVICE_MANAGED . Using parameters for
// CreateUser , you can specify the user name, set the home directory, store the
// user's public key, and assign the user's Identity and Access Management (IAM)
// role. You can also optionally add a session policy, and assign metadata with
// tags that can be used to group and search for users.
func transfer_CreateUser(cfg aws.Config, client *transfer.Client) {
	input := &transfer.CreateUserInput{
		// Role: *string, // Required
		// ServerId: *string, // Required
		// UserName: *string, // Required
	}

	if len(_transferRole) > 0 {
		input.Role = aws.String(_transferRole)
	}
	if len(_transferServerId) > 0 {
		input.ServerId = aws.String(_transferServerId)
	}
	if len(_transferUserName) > 0 {
		input.UserName = aws.String(_transferUserName)
	}
	if len(_transferHomeDirectory) > 0 {
		input.HomeDirectory = aws.String(_transferHomeDirectory)
	}
	if len(_transferHomeDirectoryMappings) > 0 {
		if err := assignInputField(input, "HomeDirectoryMappings", _transferHomeDirectoryMappings); err != nil {
			log.Errorf("invalid --home-directory-mappings: %s", err.Error())
			return
		}
	}
	if len(_transferHomeDirectoryType) > 0 {
		if err := assignInputField(input, "HomeDirectoryType", _transferHomeDirectoryType); err != nil {
			log.Errorf("invalid --home-directory-type: %s", err.Error())
			return
		}
	}
	if len(_transferPolicy) > 0 {
		input.Policy = aws.String(_transferPolicy)
	}
	if len(_transferPosixProfile) > 0 {
		if err := assignInputField(input, "PosixProfile", _transferPosixProfile); err != nil {
			log.Errorf("invalid --posix-profile: %s", err.Error())
			return
		}
	}
	if len(_transferSshPublicKeyBody) > 0 {
		input.SshPublicKeyBody = aws.String(_transferSshPublicKeyBody)
	}
	if len(_transferTags) > 0 {
		if err := assignInputField(input, "Tags", _transferTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
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

// Creates a web app based on specified parameters, and returns the ID for the new
// web app. You can configure the web app to be publicly accessible or hosted
// within a VPC.
//
// For more information about using VPC endpoints with Transfer Family, see [Create a Transfer Family web app in a VPC].
//
// [Create a Transfer Family web app in a VPC]: https://docs.aws.amazon.com/transfer/latest/userguide/create-webapp-in-vpc.html
func transfer_CreateWebApp(cfg aws.Config, client *transfer.Client) {
	input := &transfer.CreateWebAppInput{
		// IdentityProviderDetails: types.WebAppIdentityProviderDetails, // Required
	}

	if len(_transferIdentityProviderDetails) > 0 {
		if err := assignInputField(input, "IdentityProviderDetails", _transferIdentityProviderDetails); err != nil {
			log.Errorf("invalid --identity-provider-details: %s", err.Error())
			return
		}
	}
	if len(_transferAccessEndpoint) > 0 {
		input.AccessEndpoint = aws.String(_transferAccessEndpoint)
	}
	if len(_transferEndpointDetails) > 0 {
		if err := assignInputField(input, "EndpointDetails", _transferEndpointDetails); err != nil {
			log.Errorf("invalid --endpoint-details: %s", err.Error())
			return
		}
	}
	if len(_transferTags) > 0 {
		if err := assignInputField(input, "Tags", _transferTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_transferWebAppEndpointPolicy) > 0 {
		if err := assignInputField(input, "WebAppEndpointPolicy", _transferWebAppEndpointPolicy); err != nil {
			log.Errorf("invalid --web-app-endpoint-policy: %s", err.Error())
			return
		}
	}
	if len(_transferWebAppUnits) > 0 {
		if err := assignInputField(input, "WebAppUnits", _transferWebAppUnits); err != nil {
			log.Errorf("invalid --web-app-units: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateWebApp(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Allows you to create a workflow with specified steps and step details the
// workflow invokes after file transfer completes. After creating a workflow, you
// can associate the workflow created with any transfer servers by specifying the
// workflow-details field in CreateServer and UpdateServer operations.
func transfer_CreateWorkflow(cfg aws.Config, client *transfer.Client) {
	input := &transfer.CreateWorkflowInput{
		// Steps: []types.WorkflowStep, // Required
	}

	if len(_transferSteps) > 0 {
		if err := assignInputField(input, "Steps", _transferSteps); err != nil {
			log.Errorf("invalid --steps: %s", err.Error())
			return
		}
	}
	if len(_transferDescription) > 0 {
		input.Description = aws.String(_transferDescription)
	}
	if len(_transferOnExceptionSteps) > 0 {
		if err := assignInputField(input, "OnExceptionSteps", _transferOnExceptionSteps); err != nil {
			log.Errorf("invalid --on-exception-steps: %s", err.Error())
			return
		}
	}
	if len(_transferTags) > 0 {
		if err := assignInputField(input, "Tags", _transferTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateWorkflow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Allows you to delete the access specified in the ServerID and ExternalID
// parameters.
func transfer_DeleteAccess(cfg aws.Config, client *transfer.Client) {
	input := &transfer.DeleteAccessInput{
		// ExternalId: *string, // Required
		// ServerId: *string, // Required
	}

	if len(_transferExternalId) > 0 {
		input.ExternalId = aws.String(_transferExternalId)
	}
	if len(_transferServerId) > 0 {
		input.ServerId = aws.String(_transferServerId)
	}

	if resp, err := client.DeleteAccess(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete the agreement that's specified in the provided AgreementId .
func transfer_DeleteAgreement(cfg aws.Config, client *transfer.Client) {
	input := &transfer.DeleteAgreementInput{
		// AgreementId: *string, // Required
		// ServerId: *string, // Required
	}

	if len(_transferAgreementId) > 0 {
		input.AgreementId = aws.String(_transferAgreementId)
	}
	if len(_transferServerId) > 0 {
		input.ServerId = aws.String(_transferServerId)
	}

	if resp, err := client.DeleteAgreement(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the certificate that's specified in the CertificateId parameter.
func transfer_DeleteCertificate(cfg aws.Config, client *transfer.Client) {
	input := &transfer.DeleteCertificateInput{
		// CertificateId: *string, // Required
	}

	if len(_transferCertificateId) > 0 {
		input.CertificateId = aws.String(_transferCertificateId)
	}

	if resp, err := client.DeleteCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the connector that's specified in the provided ConnectorId .
func transfer_DeleteConnector(cfg aws.Config, client *transfer.Client) {
	input := &transfer.DeleteConnectorInput{
		// ConnectorId: *string, // Required
	}

	if len(_transferConnectorId) > 0 {
		input.ConnectorId = aws.String(_transferConnectorId)
	}

	if resp, err := client.DeleteConnector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the host key that's specified in the HostKeyId parameter.
func transfer_DeleteHostKey(cfg aws.Config, client *transfer.Client) {
	input := &transfer.DeleteHostKeyInput{
		// HostKeyId: *string, // Required
		// ServerId: *string, // Required
	}

	if len(_transferHostKeyId) > 0 {
		input.HostKeyId = aws.String(_transferHostKeyId)
	}
	if len(_transferServerId) > 0 {
		input.ServerId = aws.String(_transferServerId)
	}

	if resp, err := client.DeleteHostKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the profile that's specified in the ProfileId parameter.
func transfer_DeleteProfile(cfg aws.Config, client *transfer.Client) {
	input := &transfer.DeleteProfileInput{
		// ProfileId: *string, // Required
	}

	if len(_transferProfileId) > 0 {
		input.ProfileId = aws.String(_transferProfileId)
	}

	if resp, err := client.DeleteProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the file transfer protocol-enabled server that you specify.
// No response returns from this operation.
func transfer_DeleteServer(cfg aws.Config, client *transfer.Client) {
	input := &transfer.DeleteServerInput{
		// ServerId: *string, // Required
	}

	if len(_transferServerId) > 0 {
		input.ServerId = aws.String(_transferServerId)
	}

	if resp, err := client.DeleteServer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a user's Secure Shell (SSH) public key.
func transfer_DeleteSshPublicKey(cfg aws.Config, client *transfer.Client) {
	input := &transfer.DeleteSshPublicKeyInput{
		// ServerId: *string, // Required
		// SshPublicKeyId: *string, // Required
		// UserName: *string, // Required
	}

	if len(_transferServerId) > 0 {
		input.ServerId = aws.String(_transferServerId)
	}
	if len(_transferSshPublicKeyId) > 0 {
		input.SshPublicKeyId = aws.String(_transferSshPublicKeyId)
	}
	if len(_transferUserName) > 0 {
		input.UserName = aws.String(_transferUserName)
	}

	if resp, err := client.DeleteSshPublicKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the user belonging to a file transfer protocol-enabled server you
// specify.
//
// No response returns from this operation.
//
// When you delete a user from a server, the user's information is lost.
func transfer_DeleteUser(cfg aws.Config, client *transfer.Client) {
	input := &transfer.DeleteUserInput{
		// ServerId: *string, // Required
		// UserName: *string, // Required
	}

	if len(_transferServerId) > 0 {
		input.ServerId = aws.String(_transferServerId)
	}
	if len(_transferUserName) > 0 {
		input.UserName = aws.String(_transferUserName)
	}

	if resp, err := client.DeleteUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified web app.
func transfer_DeleteWebApp(cfg aws.Config, client *transfer.Client) {
	input := &transfer.DeleteWebAppInput{
		// WebAppId: *string, // Required
	}

	if len(_transferWebAppId) > 0 {
		input.WebAppId = aws.String(_transferWebAppId)
	}

	if resp, err := client.DeleteWebApp(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the WebAppCustomization object that corresponds to the web app ID
// specified.
func transfer_DeleteWebAppCustomization(cfg aws.Config, client *transfer.Client) {
	input := &transfer.DeleteWebAppCustomizationInput{
		// WebAppId: *string, // Required
	}

	if len(_transferWebAppId) > 0 {
		input.WebAppId = aws.String(_transferWebAppId)
	}

	if resp, err := client.DeleteWebAppCustomization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified workflow.
func transfer_DeleteWorkflow(cfg aws.Config, client *transfer.Client) {
	input := &transfer.DeleteWorkflowInput{
		// WorkflowId: *string, // Required
	}

	if len(_transferWorkflowId) > 0 {
		input.WorkflowId = aws.String(_transferWorkflowId)
	}

	if resp, err := client.DeleteWorkflow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the access that is assigned to the specific file transfer
// protocol-enabled server, as identified by its ServerId property and its
// ExternalId .
//
// The response from this call returns the properties of the access that is
// associated with the ServerId value that was specified.
func transfer_DescribeAccess(cfg aws.Config, client *transfer.Client) {
	input := &transfer.DescribeAccessInput{
		// ExternalId: *string, // Required
		// ServerId: *string, // Required
	}

	if len(_transferExternalId) > 0 {
		input.ExternalId = aws.String(_transferExternalId)
	}
	if len(_transferServerId) > 0 {
		input.ServerId = aws.String(_transferServerId)
	}

	if resp, err := client.DescribeAccess(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the agreement that's identified by the AgreementId .
func transfer_DescribeAgreement(cfg aws.Config, client *transfer.Client) {
	input := &transfer.DescribeAgreementInput{
		// AgreementId: *string, // Required
		// ServerId: *string, // Required
	}

	if len(_transferAgreementId) > 0 {
		input.AgreementId = aws.String(_transferAgreementId)
	}
	if len(_transferServerId) > 0 {
		input.ServerId = aws.String(_transferServerId)
	}

	if resp, err := client.DescribeAgreement(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the certificate that's identified by the CertificateId .
// Transfer Family automatically publishes a Amazon CloudWatch metric called
// DaysUntilExpiry for imported certificates. This metric tracks the number of days
// until the certificate expires based on the InactiveDate . The metric is
// available in the AWS/Transfer namespace and includes the CertificateId as a
// dimension.
func transfer_DescribeCertificate(cfg aws.Config, client *transfer.Client) {
	input := &transfer.DescribeCertificateInput{
		// CertificateId: *string, // Required
	}

	if len(_transferCertificateId) > 0 {
		input.CertificateId = aws.String(_transferCertificateId)
	}

	if resp, err := client.DescribeCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the connector that's identified by the ConnectorId.
func transfer_DescribeConnector(cfg aws.Config, client *transfer.Client) {
	input := &transfer.DescribeConnectorInput{
		// ConnectorId: *string, // Required
	}

	if len(_transferConnectorId) > 0 {
		input.ConnectorId = aws.String(_transferConnectorId)
	}

	if resp, err := client.DescribeConnector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// You can use DescribeExecution to check the details of the execution of the
// specified workflow.
//
// This API call only returns details for in-progress workflows.
//
// If you provide an ID for an execution that is not in progress, or if the
// execution doesn't match the specified workflow ID, you receive a
// ResourceNotFound exception.
func transfer_DescribeExecution(cfg aws.Config, client *transfer.Client) {
	input := &transfer.DescribeExecutionInput{
		// ExecutionId: *string, // Required
		// WorkflowId: *string, // Required
	}

	if len(_transferExecutionId) > 0 {
		input.ExecutionId = aws.String(_transferExecutionId)
	}
	if len(_transferWorkflowId) > 0 {
		input.WorkflowId = aws.String(_transferWorkflowId)
	}

	if resp, err := client.DescribeExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the details of the host key that's specified by the HostKeyId and
// ServerId .
func transfer_DescribeHostKey(cfg aws.Config, client *transfer.Client) {
	input := &transfer.DescribeHostKeyInput{
		// HostKeyId: *string, // Required
		// ServerId: *string, // Required
	}

	if len(_transferHostKeyId) > 0 {
		input.HostKeyId = aws.String(_transferHostKeyId)
	}
	if len(_transferServerId) > 0 {
		input.ServerId = aws.String(_transferServerId)
	}

	if resp, err := client.DescribeHostKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the details of the profile that's specified by the ProfileId .
func transfer_DescribeProfile(cfg aws.Config, client *transfer.Client) {
	input := &transfer.DescribeProfileInput{
		// ProfileId: *string, // Required
	}

	if len(_transferProfileId) > 0 {
		input.ProfileId = aws.String(_transferProfileId)
	}

	if resp, err := client.DescribeProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the security policy that is attached to your server or SFTP
// connector. The response contains a description of the security policy's
// properties. For more information about security policies, see [Working with security policies for servers]or [Working with security policies for SFTP connectors].
//
// [Working with security policies for SFTP connectors]: https://docs.aws.amazon.com/transfer/latest/userguide/security-policies-connectors.html
// [Working with security policies for servers]: https://docs.aws.amazon.com/transfer/latest/userguide/security-policies.html
func transfer_DescribeSecurityPolicy(cfg aws.Config, client *transfer.Client) {
	input := &transfer.DescribeSecurityPolicyInput{
		// SecurityPolicyName: *string, // Required
	}

	if len(_transferSecurityPolicyName) > 0 {
		input.SecurityPolicyName = aws.String(_transferSecurityPolicyName)
	}

	if resp, err := client.DescribeSecurityPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a file transfer protocol-enabled server that you specify by passing
// the ServerId parameter.
//
// The response contains a description of a server's properties. When you set
// EndpointType to VPC, the response will contain the EndpointDetails .
func transfer_DescribeServer(cfg aws.Config, client *transfer.Client) {
	input := &transfer.DescribeServerInput{
		// ServerId: *string, // Required
	}

	if len(_transferServerId) > 0 {
		input.ServerId = aws.String(_transferServerId)
	}

	if resp, err := client.DescribeServer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the user assigned to the specific file transfer protocol-enabled
// server, as identified by its ServerId property.
//
// The response from this call returns the properties of the user associated with
// the ServerId value that was specified.
func transfer_DescribeUser(cfg aws.Config, client *transfer.Client) {
	input := &transfer.DescribeUserInput{
		// ServerId: *string, // Required
		// UserName: *string, // Required
	}

	if len(_transferServerId) > 0 {
		input.ServerId = aws.String(_transferServerId)
	}
	if len(_transferUserName) > 0 {
		input.UserName = aws.String(_transferUserName)
	}

	if resp, err := client.DescribeUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the web app that's identified by WebAppId . The response includes
// endpoint configuration details such as whether the web app is publicly
// accessible or VPC hosted.
//
// For more information about using VPC endpoints with Transfer Family, see [Create a Transfer Family web app in a VPC].
//
// [Create a Transfer Family web app in a VPC]: https://docs.aws.amazon.com/transfer/latest/userguide/create-webapp-in-vpc.html
func transfer_DescribeWebApp(cfg aws.Config, client *transfer.Client) {
	input := &transfer.DescribeWebAppInput{
		// WebAppId: *string, // Required
	}

	if len(_transferWebAppId) > 0 {
		input.WebAppId = aws.String(_transferWebAppId)
	}

	if resp, err := client.DescribeWebApp(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the web app customization object that's identified by WebAppId .
func transfer_DescribeWebAppCustomization(cfg aws.Config, client *transfer.Client) {
	input := &transfer.DescribeWebAppCustomizationInput{
		// WebAppId: *string, // Required
	}

	if len(_transferWebAppId) > 0 {
		input.WebAppId = aws.String(_transferWebAppId)
	}

	if resp, err := client.DescribeWebAppCustomization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the specified workflow.
func transfer_DescribeWorkflow(cfg aws.Config, client *transfer.Client) {
	input := &transfer.DescribeWorkflowInput{
		// WorkflowId: *string, // Required
	}

	if len(_transferWorkflowId) > 0 {
		input.WorkflowId = aws.String(_transferWorkflowId)
	}

	if resp, err := client.DescribeWorkflow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Imports the signing and encryption certificates that you need to create local
// (AS2) profiles and partner profiles.
//
// You can import both the certificate and its chain in the Certificate parameter.
//
// After importing a certificate, Transfer Family automatically creates a Amazon
// CloudWatch metric called DaysUntilExpiry that tracks the number of days until
// the certificate expires. The metric is based on the InactiveDate parameter and
// is published daily in the AWS/Transfer namespace.
//
// It can take up to a full day after importing a certificate for Transfer Family
// to emit the DaysUntilExpiry metric to your account.
//
// If you use the Certificate parameter to upload both the certificate and its
// chain, don't use the CertificateChain parameter.
//
// # CloudWatch monitoring
//
// The DaysUntilExpiry metric includes the following specifications:
//
// - Units: Count (days)
//
// - Dimensions: CertificateId (always present), Description (if provided during
// certificate import)
//
// - Statistics: Minimum, Maximum, Average
//
// - Frequency: Published daily
func transfer_ImportCertificate(cfg aws.Config, client *transfer.Client) {
	input := &transfer.ImportCertificateInput{
		// Certificate: *string, // Required
		// Usage: types.CertificateUsageType, // Required
	}

	if len(_transferCertificate) > 0 {
		input.Certificate = aws.String(_transferCertificate)
	}
	if len(_transferUsage) > 0 {
		if err := assignInputField(input, "Usage", _transferUsage); err != nil {
			log.Errorf("invalid --usage: %s", err.Error())
			return
		}
	}
	if len(_transferActiveDate) > 0 {
		if err := assignInputField(input, "ActiveDate", _transferActiveDate); err != nil {
			log.Errorf("invalid --active-date: %s", err.Error())
			return
		}
	}
	if len(_transferCertificateChain) > 0 {
		input.CertificateChain = aws.String(_transferCertificateChain)
	}
	if len(_transferDescription) > 0 {
		input.Description = aws.String(_transferDescription)
	}
	if len(_transferInactiveDate) > 0 {
		if err := assignInputField(input, "InactiveDate", _transferInactiveDate); err != nil {
			log.Errorf("invalid --inactive-date: %s", err.Error())
			return
		}
	}
	if len(_transferPrivateKey) > 0 {
		input.PrivateKey = aws.String(_transferPrivateKey)
	}
	if len(_transferTags) > 0 {
		if err := assignInputField(input, "Tags", _transferTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.ImportCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a host key to the server that's specified by the ServerId parameter.
func transfer_ImportHostKey(cfg aws.Config, client *transfer.Client) {
	input := &transfer.ImportHostKeyInput{
		// HostKeyBody: *string, // Required
		// ServerId: *string, // Required
	}

	if len(_transferHostKeyBody) > 0 {
		input.HostKeyBody = aws.String(_transferHostKeyBody)
	}
	if len(_transferServerId) > 0 {
		input.ServerId = aws.String(_transferServerId)
	}
	if len(_transferDescription) > 0 {
		input.Description = aws.String(_transferDescription)
	}
	if len(_transferTags) > 0 {
		if err := assignInputField(input, "Tags", _transferTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.ImportHostKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a Secure Shell (SSH) public key to a Transfer Family user identified by a
// UserName value assigned to the specific file transfer protocol-enabled server,
// identified by ServerId .
//
// The response returns the UserName value, the ServerId value, and the name of
// the SshPublicKeyId .
func transfer_ImportSshPublicKey(cfg aws.Config, client *transfer.Client) {
	input := &transfer.ImportSshPublicKeyInput{
		// ServerId: *string, // Required
		// SshPublicKeyBody: *string, // Required
		// UserName: *string, // Required
	}

	if len(_transferServerId) > 0 {
		input.ServerId = aws.String(_transferServerId)
	}
	if len(_transferSshPublicKeyBody) > 0 {
		input.SshPublicKeyBody = aws.String(_transferSshPublicKeyBody)
	}
	if len(_transferUserName) > 0 {
		input.UserName = aws.String(_transferUserName)
	}

	if resp, err := client.ImportSshPublicKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the details for all the accesses you have on your server.
func transfer_ListAccesses(cfg aws.Config, client *transfer.Client) {
	input := &transfer.ListAccessesInput{
		// ServerId: *string, // Required
	}

	if len(_transferServerId) > 0 {
		input.ServerId = aws.String(_transferServerId)
	}
	if len(_transferMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _transferMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_transferNextToken) > 0 {
		input.NextToken = aws.String(_transferNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAccesses(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*transfer.ListAccessesOutput
	p := transfer.NewListAccessesPaginator(client, input)
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

// Returns a list of the agreements for the server that's identified by the
// ServerId that you supply. If you want to limit the results to a certain number,
// supply a value for the MaxResults parameter. If you ran the command previously
// and received a value for NextToken , you can supply that value to continue
// listing agreements from where you left off.
func transfer_ListAgreements(cfg aws.Config, client *transfer.Client) {
	input := &transfer.ListAgreementsInput{
		// ServerId: *string, // Required
	}

	if len(_transferServerId) > 0 {
		input.ServerId = aws.String(_transferServerId)
	}
	if len(_transferMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _transferMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_transferNextToken) > 0 {
		input.NextToken = aws.String(_transferNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAgreements(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*transfer.ListAgreementsOutput
	p := transfer.NewListAgreementsPaginator(client, input)
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

// Returns a list of the current certificates that have been imported into
// Transfer Family. If you want to limit the results to a certain number, supply a
// value for the MaxResults parameter. If you ran the command previously and
// received a value for the NextToken parameter, you can supply that value to
// continue listing certificates from where you left off.
func transfer_ListCertificates(cfg aws.Config, client *transfer.Client) {
	input := &transfer.ListCertificatesInput{}

	if len(_transferMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _transferMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_transferNextToken) > 0 {
		input.NextToken = aws.String(_transferNextToken)
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

	var results []*transfer.ListCertificatesOutput
	p := transfer.NewListCertificatesPaginator(client, input)
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

// Lists the connectors for the specified Region.
func transfer_ListConnectors(cfg aws.Config, client *transfer.Client) {
	input := &transfer.ListConnectorsInput{}

	if len(_transferMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _transferMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_transferNextToken) > 0 {
		input.NextToken = aws.String(_transferNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListConnectors(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*transfer.ListConnectorsOutput
	p := transfer.NewListConnectorsPaginator(client, input)
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

// Lists all in-progress executions for the specified workflow.
// If the specified workflow ID cannot be found, ListExecutions returns a
// ResourceNotFound exception.
func transfer_ListExecutions(cfg aws.Config, client *transfer.Client) {
	input := &transfer.ListExecutionsInput{
		// WorkflowId: *string, // Required
	}

	if len(_transferWorkflowId) > 0 {
		input.WorkflowId = aws.String(_transferWorkflowId)
	}
	if len(_transferMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _transferMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_transferNextToken) > 0 {
		input.NextToken = aws.String(_transferNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListExecutions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*transfer.ListExecutionsOutput
	p := transfer.NewListExecutionsPaginator(client, input)
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

// Returns real-time updates and detailed information on the status of each
// individual file being transferred in a specific file transfer operation. You
// specify the file transfer by providing its ConnectorId and its TransferId .
//
// File transfer results are available up to 7 days after an operation has been
// requested.
func transfer_ListFileTransferResults(cfg aws.Config, client *transfer.Client) {
	input := &transfer.ListFileTransferResultsInput{
		// ConnectorId: *string, // Required
		// TransferId: *string, // Required
	}

	if len(_transferConnectorId) > 0 {
		input.ConnectorId = aws.String(_transferConnectorId)
	}
	if len(_transferTransferId) > 0 {
		input.TransferId = aws.String(_transferTransferId)
	}
	if len(_transferMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _transferMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_transferNextToken) > 0 {
		input.NextToken = aws.String(_transferNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListFileTransferResults(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*transfer.ListFileTransferResultsOutput
	p := transfer.NewListFileTransferResultsPaginator(client, input)
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

// Returns a list of host keys for the server that's specified by the ServerId
// parameter.
func transfer_ListHostKeys(cfg aws.Config, client *transfer.Client) {
	input := &transfer.ListHostKeysInput{
		// ServerId: *string, // Required
	}

	if len(_transferServerId) > 0 {
		input.ServerId = aws.String(_transferServerId)
	}
	if len(_transferMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _transferMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_transferNextToken) > 0 {
		input.NextToken = aws.String(_transferNextToken)
	}

	if resp, err := client.ListHostKeys(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of the profiles for your system. If you want to limit the
// results to a certain number, supply a value for the MaxResults parameter. If
// you ran the command previously and received a value for NextToken , you can
// supply that value to continue listing profiles from where you left off.
func transfer_ListProfiles(cfg aws.Config, client *transfer.Client) {
	input := &transfer.ListProfilesInput{}

	if len(_transferMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _transferMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_transferNextToken) > 0 {
		input.NextToken = aws.String(_transferNextToken)
	}
	if len(_transferProfileType) > 0 {
		if err := assignInputField(input, "ProfileType", _transferProfileType); err != nil {
			log.Errorf("invalid --profile-type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListProfiles(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*transfer.ListProfilesOutput
	p := transfer.NewListProfilesPaginator(client, input)
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

// Lists the security policies that are attached to your servers and SFTP
// connectors. For more information about security policies, see [Working with security policies for servers]or [Working with security policies for SFTP connectors].
//
// [Working with security policies for SFTP connectors]: https://docs.aws.amazon.com/transfer/latest/userguide/security-policies-connectors.html
// [Working with security policies for servers]: https://docs.aws.amazon.com/transfer/latest/userguide/security-policies.html
func transfer_ListSecurityPolicies(cfg aws.Config, client *transfer.Client) {
	input := &transfer.ListSecurityPoliciesInput{}

	if len(_transferMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _transferMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_transferNextToken) > 0 {
		input.NextToken = aws.String(_transferNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSecurityPolicies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*transfer.ListSecurityPoliciesOutput
	p := transfer.NewListSecurityPoliciesPaginator(client, input)
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

// Lists the file transfer protocol-enabled servers that are associated with your
// Amazon Web Services account.
func transfer_ListServers(cfg aws.Config, client *transfer.Client) {
	input := &transfer.ListServersInput{}

	if len(_transferMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _transferMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_transferNextToken) > 0 {
		input.NextToken = aws.String(_transferNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListServers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*transfer.ListServersOutput
	p := transfer.NewListServersPaginator(client, input)
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

// Lists all of the tags associated with the Amazon Resource Name (ARN) that you
// specify. The resource can be a user, server, or role.
func transfer_ListTagsForResource(cfg aws.Config, client *transfer.Client) {
	input := &transfer.ListTagsForResourceInput{
		// Arn: *string, // Required
	}

	if len(_transferArn) > 0 {
		input.Arn = aws.String(_transferArn)
	}
	if len(_transferMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _transferMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_transferNextToken) > 0 {
		input.NextToken = aws.String(_transferNextToken)
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

	var results []*transfer.ListTagsForResourceOutput
	p := transfer.NewListTagsForResourcePaginator(client, input)
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

// Lists the users for a file transfer protocol-enabled server that you specify by
// passing the ServerId parameter.
func transfer_ListUsers(cfg aws.Config, client *transfer.Client) {
	input := &transfer.ListUsersInput{
		// ServerId: *string, // Required
	}

	if len(_transferServerId) > 0 {
		input.ServerId = aws.String(_transferServerId)
	}
	if len(_transferMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _transferMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_transferNextToken) > 0 {
		input.NextToken = aws.String(_transferNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListUsers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*transfer.ListUsersOutput
	p := transfer.NewListUsersPaginator(client, input)
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

// Lists all web apps associated with your Amazon Web Services account for your
// current region. The response includes the endpoint type for each web app,
// showing whether it is publicly accessible or VPC hosted.
//
// For more information about using VPC endpoints with Transfer Family, see [Create a Transfer Family web app in a VPC].
//
// [Create a Transfer Family web app in a VPC]: https://docs.aws.amazon.com/transfer/latest/userguide/create-webapp-in-vpc.html
func transfer_ListWebApps(cfg aws.Config, client *transfer.Client) {
	input := &transfer.ListWebAppsInput{}

	if len(_transferMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _transferMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_transferNextToken) > 0 {
		input.NextToken = aws.String(_transferNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListWebApps(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*transfer.ListWebAppsOutput
	p := transfer.NewListWebAppsPaginator(client, input)
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

// Lists all workflows associated with your Amazon Web Services account for your
// current region.
func transfer_ListWorkflows(cfg aws.Config, client *transfer.Client) {
	input := &transfer.ListWorkflowsInput{}

	if len(_transferMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _transferMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_transferNextToken) > 0 {
		input.NextToken = aws.String(_transferNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListWorkflows(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*transfer.ListWorkflowsOutput
	p := transfer.NewListWorkflowsPaginator(client, input)
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

// Sends a callback for asynchronous custom steps.
// The ExecutionId , WorkflowId , and Token are passed to the target resource
// during execution of a custom step of a workflow. You must include those with
// their callback as well as providing a status.
func transfer_SendWorkflowStepState(cfg aws.Config, client *transfer.Client) {
	input := &transfer.SendWorkflowStepStateInput{
		// ExecutionId: *string, // Required
		// Status: types.CustomStepStatus, // Required
		// Token: *string, // Required
		// WorkflowId: *string, // Required
	}

	if len(_transferExecutionId) > 0 {
		input.ExecutionId = aws.String(_transferExecutionId)
	}
	if len(_transferStatus) > 0 {
		if err := assignInputField(input, "Status", _transferStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_transferToken) > 0 {
		input.Token = aws.String(_transferToken)
	}
	if len(_transferWorkflowId) > 0 {
		input.WorkflowId = aws.String(_transferWorkflowId)
	}

	if resp, err := client.SendWorkflowStepState(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of the contents of a directory from a remote SFTP server. You
// specify the connector ID, the output path, and the remote directory path. You
// can also specify the optional MaxItems value to control the maximum number of
// items that are listed from the remote directory. This API returns a list of all
// files and directories in the remote directory (up to the maximum value), but
// does not return files or folders in sub-directories. That is, it only returns a
// list of files and directories one-level deep.
//
// After you receive the listing file, you can provide the files that you want to
// transfer to the RetrieveFilePaths parameter of the StartFileTransfer API call.
//
// The naming convention for the output file is  connector-ID-listing-ID.json . The
// output file contains the following information:
//
// - filePath : the complete path of a remote file, relative to the directory of
// the listing request for your SFTP connector on the remote server.
//
// - modifiedTimestamp : the last time the file was modified, in UTC time format.
// This field is optional. If the remote file attributes don't contain a timestamp,
// it is omitted from the file listing.
//
// - size : the size of the file, in bytes. This field is optional. If the remote
// file attributes don't contain a file size, it is omitted from the file listing.
//
// - path : the complete path of a remote directory, relative to the directory of
// the listing request for your SFTP connector on the remote server.
//
// - truncated : a flag indicating whether the list output contains all of the
// items contained in the remote directory or not. If your Truncated output value
// is true, you can increase the value provided in the optional max-items input
// attribute to be able to list more items (up to the maximum allowed list size of
// 10,000 items).
func transfer_StartDirectoryListing(cfg aws.Config, client *transfer.Client) {
	input := &transfer.StartDirectoryListingInput{
		// ConnectorId: *string, // Required
		// OutputDirectoryPath: *string, // Required
		// RemoteDirectoryPath: *string, // Required
	}

	if len(_transferConnectorId) > 0 {
		input.ConnectorId = aws.String(_transferConnectorId)
	}
	if len(_transferOutputDirectoryPath) > 0 {
		input.OutputDirectoryPath = aws.String(_transferOutputDirectoryPath)
	}
	if len(_transferRemoteDirectoryPath) > 0 {
		input.RemoteDirectoryPath = aws.String(_transferRemoteDirectoryPath)
	}
	if len(_transferMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _transferMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartDirectoryListing(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Begins a file transfer between local Amazon Web Services storage and a remote
// AS2 or SFTP server.
//
// - For an AS2 connector, you specify the ConnectorId and one or more
// SendFilePaths to identify the files you want to transfer.
//
// - For an SFTP connector, the file transfer can be either outbound or inbound.
// In both cases, you specify the ConnectorId . Depending on the direction of the
// transfer, you also specify the following items:
//
// - If you are transferring file from a partner's SFTP server to Amazon Web
// Services storage, you specify one or more RetrieveFilePaths to identify the
// files you want to transfer, and a LocalDirectoryPath to specify the
// destination folder.
//
// - If you are transferring file to a partner's SFTP server from Amazon Web
// Services storage, you specify one or more SendFilePaths to identify the files
// you want to transfer, and a RemoteDirectoryPath to specify the destination
// folder.
func transfer_StartFileTransfer(cfg aws.Config, client *transfer.Client) {
	input := &transfer.StartFileTransferInput{
		// ConnectorId: *string, // Required
	}

	if len(_transferConnectorId) > 0 {
		input.ConnectorId = aws.String(_transferConnectorId)
	}
	if len(_transferCustomHttpHeaders) > 0 {
		if err := assignInputField(input, "CustomHttpHeaders", _transferCustomHttpHeaders); err != nil {
			log.Errorf("invalid --custom-http-headers: %s", err.Error())
			return
		}
	}
	if len(_transferLocalDirectoryPath) > 0 {
		input.LocalDirectoryPath = aws.String(_transferLocalDirectoryPath)
	}
	if len(_transferRemoteDirectoryPath) > 0 {
		input.RemoteDirectoryPath = aws.String(_transferRemoteDirectoryPath)
	}
	if len(_transferRetrieveFilePaths) > 0 {
		input.RetrieveFilePaths = append([]string(nil), _transferRetrieveFilePaths...)
	}
	if len(_transferSendFilePaths) > 0 {
		input.SendFilePaths = append([]string(nil), _transferSendFilePaths...)
	}

	if resp, err := client.StartFileTransfer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a file or directory on the remote SFTP server.
func transfer_StartRemoteDelete(cfg aws.Config, client *transfer.Client) {
	input := &transfer.StartRemoteDeleteInput{
		// ConnectorId: *string, // Required
		// DeletePath: *string, // Required
	}

	if len(_transferConnectorId) > 0 {
		input.ConnectorId = aws.String(_transferConnectorId)
	}
	if len(_transferDeletePath) > 0 {
		input.DeletePath = aws.String(_transferDeletePath)
	}

	if resp, err := client.StartRemoteDelete(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Moves or renames a file or directory on the remote SFTP server.
func transfer_StartRemoteMove(cfg aws.Config, client *transfer.Client) {
	input := &transfer.StartRemoteMoveInput{
		// ConnectorId: *string, // Required
		// SourcePath: *string, // Required
		// TargetPath: *string, // Required
	}

	if len(_transferConnectorId) > 0 {
		input.ConnectorId = aws.String(_transferConnectorId)
	}
	if len(_transferSourcePath) > 0 {
		input.SourcePath = aws.String(_transferSourcePath)
	}
	if len(_transferTargetPath) > 0 {
		input.TargetPath = aws.String(_transferTargetPath)
	}

	if resp, err := client.StartRemoteMove(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes the state of a file transfer protocol-enabled server from OFFLINE to
// ONLINE . It has no impact on a server that is already ONLINE . An ONLINE server
// can accept and process file transfer jobs.
//
// The state of STARTING indicates that the server is in an intermediate state,
// either not fully able to respond, or not fully online. The values of
// START_FAILED can indicate an error condition.
//
// No response is returned from this call.
func transfer_StartServer(cfg aws.Config, client *transfer.Client) {
	input := &transfer.StartServerInput{
		// ServerId: *string, // Required
	}

	if len(_transferServerId) > 0 {
		input.ServerId = aws.String(_transferServerId)
	}

	if resp, err := client.StartServer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes the state of a file transfer protocol-enabled server from ONLINE to
// OFFLINE . An OFFLINE server cannot accept and process file transfer jobs.
// Information tied to your server, such as server and user properties, are not
// affected by stopping your server.
//
// Stopping the server does not reduce or impact your file transfer protocol
// endpoint billing; you must delete the server to stop being billed.
//
// The state of STOPPING indicates that the server is in an intermediate state,
// either not fully able to respond, or not fully offline. The values of
// STOP_FAILED can indicate an error condition.
//
// No response is returned from this call.
func transfer_StopServer(cfg aws.Config, client *transfer.Client) {
	input := &transfer.StopServerInput{
		// ServerId: *string, // Required
	}

	if len(_transferServerId) > 0 {
		input.ServerId = aws.String(_transferServerId)
	}

	if resp, err := client.StopServer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attaches a key-value pair to a resource, as identified by its Amazon Resource
// Name (ARN). Resources are users, servers, roles, and other entities.
//
// There is no response returned from this call.
func transfer_TagResource(cfg aws.Config, client *transfer.Client) {
	input := &transfer.TagResourceInput{
		// Arn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_transferArn) > 0 {
		input.Arn = aws.String(_transferArn)
	}
	if len(_transferTags) > 0 {
		if err := assignInputField(input, "Tags", _transferTags); err != nil {
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

// Tests whether your SFTP connector is set up successfully. We highly recommend
// that you call this operation to test your ability to transfer files between
// local Amazon Web Services storage and a trading partner's SFTP server.
func transfer_TestConnection(cfg aws.Config, client *transfer.Client) {
	input := &transfer.TestConnectionInput{
		// ConnectorId: *string, // Required
	}

	if len(_transferConnectorId) > 0 {
		input.ConnectorId = aws.String(_transferConnectorId)
	}

	if resp, err := client.TestConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// If the IdentityProviderType of a file transfer protocol-enabled server is
// AWS_DIRECTORY_SERVICE or API_Gateway , tests whether your identity provider is
// set up successfully. We highly recommend that you call this operation to test
// your authentication method as soon as you create your server. By doing so, you
// can troubleshoot issues with the identity provider integration to ensure that
// your users can successfully use the service.
//
// The ServerId and UserName parameters are required. The ServerProtocol , SourceIp
// , and UserPassword are all optional.
//
// Note the following:
//
// - You cannot use TestIdentityProvider if the IdentityProviderType of your
// server is SERVICE_MANAGED .
//
// - TestIdentityProvider does not work with keys: it only accepts passwords.
//
// - TestIdentityProvider can test the password operation for a custom Identity
// Provider that handles keys and passwords.
//
// - If you provide any incorrect values for any parameters, the Response field
// is empty.
//
// - If you provide a server ID for a server that uses service-managed users,
// you get an error:
//
// # An error occurred (InvalidRequestException) when calling the
//
// TestIdentityProvider operation: s-server-ID not configured for external auth
//
// - If you enter a Server ID for the --server-id parameter that does not
// identify an actual Transfer server, you receive the following error:
//
// # An error occurred (ResourceNotFoundException) when calling the
//
// TestIdentityProvider operation: Unknown server .
//
// It is possible your sever is in a different region. You can specify a region by
//
// adding the following: --region region-code , such as --region us-east-2 to
// specify a server in US East (Ohio).
func transfer_TestIdentityProvider(cfg aws.Config, client *transfer.Client) {
	input := &transfer.TestIdentityProviderInput{
		// ServerId: *string, // Required
		// UserName: *string, // Required
	}

	if len(_transferServerId) > 0 {
		input.ServerId = aws.String(_transferServerId)
	}
	if len(_transferUserName) > 0 {
		input.UserName = aws.String(_transferUserName)
	}
	if len(_transferServerProtocol) > 0 {
		if err := assignInputField(input, "ServerProtocol", _transferServerProtocol); err != nil {
			log.Errorf("invalid --server-protocol: %s", err.Error())
			return
		}
	}
	if len(_transferSourceIp) > 0 {
		input.SourceIp = aws.String(_transferSourceIp)
	}
	if len(_transferUserPassword) > 0 {
		input.UserPassword = aws.String(_transferUserPassword)
	}

	if resp, err := client.TestIdentityProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Detaches a key-value pair from a resource, as identified by its Amazon Resource
// Name (ARN). Resources are users, servers, roles, and other entities.
//
// No response is returned from this call.
func transfer_UntagResource(cfg aws.Config, client *transfer.Client) {
	input := &transfer.UntagResourceInput{
		// Arn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_transferArn) > 0 {
		input.Arn = aws.String(_transferArn)
	}
	if len(_transferTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _transferTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Allows you to update parameters for the access specified in the ServerID and
// ExternalID parameters.
func transfer_UpdateAccess(cfg aws.Config, client *transfer.Client) {
	input := &transfer.UpdateAccessInput{
		// ExternalId: *string, // Required
		// ServerId: *string, // Required
	}

	if len(_transferExternalId) > 0 {
		input.ExternalId = aws.String(_transferExternalId)
	}
	if len(_transferServerId) > 0 {
		input.ServerId = aws.String(_transferServerId)
	}
	if len(_transferHomeDirectory) > 0 {
		input.HomeDirectory = aws.String(_transferHomeDirectory)
	}
	if len(_transferHomeDirectoryMappings) > 0 {
		if err := assignInputField(input, "HomeDirectoryMappings", _transferHomeDirectoryMappings); err != nil {
			log.Errorf("invalid --home-directory-mappings: %s", err.Error())
			return
		}
	}
	if len(_transferHomeDirectoryType) > 0 {
		if err := assignInputField(input, "HomeDirectoryType", _transferHomeDirectoryType); err != nil {
			log.Errorf("invalid --home-directory-type: %s", err.Error())
			return
		}
	}
	if len(_transferPolicy) > 0 {
		input.Policy = aws.String(_transferPolicy)
	}
	if len(_transferPosixProfile) > 0 {
		if err := assignInputField(input, "PosixProfile", _transferPosixProfile); err != nil {
			log.Errorf("invalid --posix-profile: %s", err.Error())
			return
		}
	}
	if len(_transferRole) > 0 {
		input.Role = aws.String(_transferRole)
	}

	if resp, err := client.UpdateAccess(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates some of the parameters for an existing agreement. Provide the
// AgreementId and the ServerId for the agreement that you want to update, along
// with the new values for the parameters to update.
//
// Specify either BaseDirectory or CustomDirectories , but not both. Specifying
// both causes the command to fail.
//
// If you update an agreement from using base directory to custom directories, the
// base directory is no longer used. Similarly, if you change from custom
// directories to a base directory, the custom directories are no longer used.
func transfer_UpdateAgreement(cfg aws.Config, client *transfer.Client) {
	input := &transfer.UpdateAgreementInput{
		// AgreementId: *string, // Required
		// ServerId: *string, // Required
	}

	if len(_transferAgreementId) > 0 {
		input.AgreementId = aws.String(_transferAgreementId)
	}
	if len(_transferServerId) > 0 {
		input.ServerId = aws.String(_transferServerId)
	}
	if len(_transferAccessRole) > 0 {
		input.AccessRole = aws.String(_transferAccessRole)
	}
	if len(_transferBaseDirectory) > 0 {
		input.BaseDirectory = aws.String(_transferBaseDirectory)
	}
	if len(_transferCustomDirectories) > 0 {
		if err := assignInputField(input, "CustomDirectories", _transferCustomDirectories); err != nil {
			log.Errorf("invalid --custom-directories: %s", err.Error())
			return
		}
	}
	if len(_transferDescription) > 0 {
		input.Description = aws.String(_transferDescription)
	}
	if len(_transferEnforceMessageSigning) > 0 {
		if err := assignInputField(input, "EnforceMessageSigning", _transferEnforceMessageSigning); err != nil {
			log.Errorf("invalid --enforce-message-signing: %s", err.Error())
			return
		}
	}
	if len(_transferLocalProfileId) > 0 {
		input.LocalProfileId = aws.String(_transferLocalProfileId)
	}
	if len(_transferPartnerProfileId) > 0 {
		input.PartnerProfileId = aws.String(_transferPartnerProfileId)
	}
	if len(_transferPreserveFilename) > 0 {
		if err := assignInputField(input, "PreserveFilename", _transferPreserveFilename); err != nil {
			log.Errorf("invalid --preserve-filename: %s", err.Error())
			return
		}
	}
	if len(_transferStatus) > 0 {
		if err := assignInputField(input, "Status", _transferStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAgreement(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the active and inactive dates for a certificate.
func transfer_UpdateCertificate(cfg aws.Config, client *transfer.Client) {
	input := &transfer.UpdateCertificateInput{
		// CertificateId: *string, // Required
	}

	if len(_transferCertificateId) > 0 {
		input.CertificateId = aws.String(_transferCertificateId)
	}
	if len(_transferActiveDate) > 0 {
		if err := assignInputField(input, "ActiveDate", _transferActiveDate); err != nil {
			log.Errorf("invalid --active-date: %s", err.Error())
			return
		}
	}
	if len(_transferDescription) > 0 {
		input.Description = aws.String(_transferDescription)
	}
	if len(_transferInactiveDate) > 0 {
		if err := assignInputField(input, "InactiveDate", _transferInactiveDate); err != nil {
			log.Errorf("invalid --inactive-date: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates some of the parameters for an existing connector. Provide the
// ConnectorId for the connector that you want to update, along with the new values
// for the parameters to update.
func transfer_UpdateConnector(cfg aws.Config, client *transfer.Client) {
	input := &transfer.UpdateConnectorInput{
		// ConnectorId: *string, // Required
	}

	if len(_transferConnectorId) > 0 {
		input.ConnectorId = aws.String(_transferConnectorId)
	}
	if len(_transferAccessRole) > 0 {
		input.AccessRole = aws.String(_transferAccessRole)
	}
	if len(_transferAs2Config) > 0 {
		if err := assignInputField(input, "As2Config", _transferAs2Config); err != nil {
			log.Errorf("invalid --as2-config: %s", err.Error())
			return
		}
	}
	if len(_transferEgressConfig) > 0 {
		if err := assignInputField(input, "EgressConfig", _transferEgressConfig); err != nil {
			log.Errorf("invalid --egress-config: %s", err.Error())
			return
		}
	}
	if len(_transferLoggingRole) > 0 {
		input.LoggingRole = aws.String(_transferLoggingRole)
	}
	if len(_transferSecurityPolicyName) > 0 {
		input.SecurityPolicyName = aws.String(_transferSecurityPolicyName)
	}
	if len(_transferSftpConfig) > 0 {
		if err := assignInputField(input, "SftpConfig", _transferSftpConfig); err != nil {
			log.Errorf("invalid --sftp-config: %s", err.Error())
			return
		}
	}
	if len(_transferUrl) > 0 {
		input.Url = aws.String(_transferUrl)
	}

	if resp, err := client.UpdateConnector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the description for the host key that's specified by the ServerId and
// HostKeyId parameters.
func transfer_UpdateHostKey(cfg aws.Config, client *transfer.Client) {
	input := &transfer.UpdateHostKeyInput{
		// Description: *string, // Required
		// HostKeyId: *string, // Required
		// ServerId: *string, // Required
	}

	if len(_transferDescription) > 0 {
		input.Description = aws.String(_transferDescription)
	}
	if len(_transferHostKeyId) > 0 {
		input.HostKeyId = aws.String(_transferHostKeyId)
	}
	if len(_transferServerId) > 0 {
		input.ServerId = aws.String(_transferServerId)
	}

	if resp, err := client.UpdateHostKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates some of the parameters for an existing profile. Provide the ProfileId
// for the profile that you want to update, along with the new values for the
// parameters to update.
func transfer_UpdateProfile(cfg aws.Config, client *transfer.Client) {
	input := &transfer.UpdateProfileInput{
		// ProfileId: *string, // Required
	}

	if len(_transferProfileId) > 0 {
		input.ProfileId = aws.String(_transferProfileId)
	}
	if len(_transferCertificateIds) > 0 {
		input.CertificateIds = append([]string(nil), _transferCertificateIds...)
	}

	if resp, err := client.UpdateProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the file transfer protocol-enabled server's properties after that
// server has been created.
//
// The UpdateServer call returns the ServerId of the server you updated.
func transfer_UpdateServer(cfg aws.Config, client *transfer.Client) {
	input := &transfer.UpdateServerInput{
		// ServerId: *string, // Required
	}

	if len(_transferServerId) > 0 {
		input.ServerId = aws.String(_transferServerId)
	}
	if len(_transferCertificate) > 0 {
		input.Certificate = aws.String(_transferCertificate)
	}
	if len(_transferEndpointDetails) > 0 {
		if err := assignInputField(input, "EndpointDetails", _transferEndpointDetails); err != nil {
			log.Errorf("invalid --endpoint-details: %s", err.Error())
			return
		}
	}
	if len(_transferEndpointType) > 0 {
		if err := assignInputField(input, "EndpointType", _transferEndpointType); err != nil {
			log.Errorf("invalid --endpoint-type: %s", err.Error())
			return
		}
	}
	if len(_transferHostKey) > 0 {
		input.HostKey = aws.String(_transferHostKey)
	}
	if len(_transferIdentityProviderDetails) > 0 {
		if err := assignInputField(input, "IdentityProviderDetails", _transferIdentityProviderDetails); err != nil {
			log.Errorf("invalid --identity-provider-details: %s", err.Error())
			return
		}
	}
	if len(_transferIdentityProviderType) > 0 {
		if err := assignInputField(input, "IdentityProviderType", _transferIdentityProviderType); err != nil {
			log.Errorf("invalid --identity-provider-type: %s", err.Error())
			return
		}
	}
	if len(_transferIpAddressType) > 0 {
		if err := assignInputField(input, "IpAddressType", _transferIpAddressType); err != nil {
			log.Errorf("invalid --ip-address-type: %s", err.Error())
			return
		}
	}
	if len(_transferLoggingRole) > 0 {
		input.LoggingRole = aws.String(_transferLoggingRole)
	}
	if len(_transferPostAuthenticationLoginBanner) > 0 {
		input.PostAuthenticationLoginBanner = aws.String(_transferPostAuthenticationLoginBanner)
	}
	if len(_transferPreAuthenticationLoginBanner) > 0 {
		input.PreAuthenticationLoginBanner = aws.String(_transferPreAuthenticationLoginBanner)
	}
	if len(_transferProtocolDetails) > 0 {
		if err := assignInputField(input, "ProtocolDetails", _transferProtocolDetails); err != nil {
			log.Errorf("invalid --protocol-details: %s", err.Error())
			return
		}
	}
	if len(_transferProtocols) > 0 {
		if err := assignInputField(input, "Protocols", _transferProtocols); err != nil {
			log.Errorf("invalid --protocols: %s", err.Error())
			return
		}
	}
	if len(_transferS3StorageOptions) > 0 {
		if err := assignInputField(input, "S3StorageOptions", _transferS3StorageOptions); err != nil {
			log.Errorf("invalid --s3-storage-options: %s", err.Error())
			return
		}
	}
	if len(_transferSecurityPolicyName) > 0 {
		input.SecurityPolicyName = aws.String(_transferSecurityPolicyName)
	}
	if len(_transferStructuredLogDestinations) > 0 {
		input.StructuredLogDestinations = append([]string(nil), _transferStructuredLogDestinations...)
	}
	if len(_transferWorkflowDetails) > 0 {
		if err := assignInputField(input, "WorkflowDetails", _transferWorkflowDetails); err != nil {
			log.Errorf("invalid --workflow-details: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateServer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assigns new properties to a user. Parameters you pass modify any or all of the
// following: the home directory, role, and policy for the UserName and ServerId
// you specify.
//
// The response returns the ServerId and the UserName for the updated user.
//
// In the console, you can select Restricted when you create or update a user.
// This ensures that the user can't access anything outside of their home
// directory. The programmatic way to configure this behavior is to update the
// user. Set their HomeDirectoryType to LOGICAL , and specify HomeDirectoryMappings
// with Entry as root ( / ) and Target as their home directory.
//
// For example, if the user's home directory is /test/admin-user , the following
// command updates the user so that their configuration in the console shows the
// Restricted flag as selected.
//
// aws transfer update-user --server-id <server-id> --user-name admin-user
// --home-directory-type LOGICAL --home-directory-mappings "[{\"Entry\":\"/\",
// \"Target\":\"/test/admin-user\"}]"
func transfer_UpdateUser(cfg aws.Config, client *transfer.Client) {
	input := &transfer.UpdateUserInput{
		// ServerId: *string, // Required
		// UserName: *string, // Required
	}

	if len(_transferServerId) > 0 {
		input.ServerId = aws.String(_transferServerId)
	}
	if len(_transferUserName) > 0 {
		input.UserName = aws.String(_transferUserName)
	}
	if len(_transferHomeDirectory) > 0 {
		input.HomeDirectory = aws.String(_transferHomeDirectory)
	}
	if len(_transferHomeDirectoryMappings) > 0 {
		if err := assignInputField(input, "HomeDirectoryMappings", _transferHomeDirectoryMappings); err != nil {
			log.Errorf("invalid --home-directory-mappings: %s", err.Error())
			return
		}
	}
	if len(_transferHomeDirectoryType) > 0 {
		if err := assignInputField(input, "HomeDirectoryType", _transferHomeDirectoryType); err != nil {
			log.Errorf("invalid --home-directory-type: %s", err.Error())
			return
		}
	}
	if len(_transferPolicy) > 0 {
		input.Policy = aws.String(_transferPolicy)
	}
	if len(_transferPosixProfile) > 0 {
		if err := assignInputField(input, "PosixProfile", _transferPosixProfile); err != nil {
			log.Errorf("invalid --posix-profile: %s", err.Error())
			return
		}
	}
	if len(_transferRole) > 0 {
		input.Role = aws.String(_transferRole)
	}

	if resp, err := client.UpdateUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assigns new properties to a web app. You can modify the access point, identity
// provider details, endpoint configuration, and the web app units.
//
// For more information about using VPC endpoints with Transfer Family, see [Create a Transfer Family web app in a VPC].
//
// [Create a Transfer Family web app in a VPC]: https://docs.aws.amazon.com/transfer/latest/userguide/create-webapp-in-vpc.html
func transfer_UpdateWebApp(cfg aws.Config, client *transfer.Client) {
	input := &transfer.UpdateWebAppInput{
		// WebAppId: *string, // Required
	}

	if len(_transferWebAppId) > 0 {
		input.WebAppId = aws.String(_transferWebAppId)
	}
	if len(_transferAccessEndpoint) > 0 {
		input.AccessEndpoint = aws.String(_transferAccessEndpoint)
	}
	if len(_transferEndpointDetails) > 0 {
		if err := assignInputField(input, "EndpointDetails", _transferEndpointDetails); err != nil {
			log.Errorf("invalid --endpoint-details: %s", err.Error())
			return
		}
	}
	if len(_transferIdentityProviderDetails) > 0 {
		if err := assignInputField(input, "IdentityProviderDetails", _transferIdentityProviderDetails); err != nil {
			log.Errorf("invalid --identity-provider-details: %s", err.Error())
			return
		}
	}
	if len(_transferWebAppUnits) > 0 {
		if err := assignInputField(input, "WebAppUnits", _transferWebAppUnits); err != nil {
			log.Errorf("invalid --web-app-units: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateWebApp(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assigns new customization properties to a web app. You can modify the icon
// file, logo file, and title.
func transfer_UpdateWebAppCustomization(cfg aws.Config, client *transfer.Client) {
	input := &transfer.UpdateWebAppCustomizationInput{
		// WebAppId: *string, // Required
	}

	if len(_transferWebAppId) > 0 {
		input.WebAppId = aws.String(_transferWebAppId)
	}
	if len(_transferFaviconFile) > 0 {
		if err := assignInputField(input, "FaviconFile", _transferFaviconFile); err != nil {
			log.Errorf("invalid --favicon-file: %s", err.Error())
			return
		}
	}
	if len(_transferLogoFile) > 0 {
		if err := assignInputField(input, "LogoFile", _transferLogoFile); err != nil {
			log.Errorf("invalid --logo-file: %s", err.Error())
			return
		}
	}
	if len(_transferTitle) > 0 {
		input.Title = aws.String(_transferTitle)
	}

	if resp, err := client.UpdateWebAppCustomization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_transferCmd)
	_transferCmd.Flags().SortFlags = false

	_transferCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_transferCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_transferCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_transferCmd.Flags().StringVarP(&_transferAccessEndpoint, "access-endpoint", "", "", "Access Endpoint")
	_transferCmd.Flags().StringVarP(&_transferAccessRole, "access-role", "", "", "Access Role")
	_transferCmd.Flags().StringVarP(&_transferActiveDate, "active-date", "", "", "Active Date")
	_transferCmd.Flags().StringVarP(&_transferAgreementId, "agreement-id", "", "", "Agreement ID")
	_transferCmd.Flags().StringVarP(&_transferArn, "arn", "", "", "ARN")
	_transferCmd.Flags().StringVarP(&_transferAs2Config, "as2-config", "", "", "As2 Config")
	_transferCmd.Flags().StringVarP(&_transferAs2Id, "as2-id", "", "", "As2 ID")
	_transferCmd.Flags().StringVarP(&_transferBaseDirectory, "base-directory", "", "", "Base Directory")
	_transferCmd.Flags().StringVarP(&_transferCertificate, "certificate", "", "", "Certificate")
	_transferCmd.Flags().StringVarP(&_transferCertificateChain, "certificate-chain", "", "", "Certificate Chain")
	_transferCmd.Flags().StringVarP(&_transferCertificateId, "certificate-id", "", "", "Certificate ID")
	_transferCmd.Flags().StringSliceVarP(&_transferCertificateIds, "certificate-ids", "", nil, "Certificate Ids")
	_transferCmd.Flags().StringVarP(&_transferConnectorId, "connector-id", "", "", "Connector ID")
	_transferCmd.Flags().StringVarP(&_transferCustomDirectories, "custom-directories", "", "", "Custom Directories")
	_transferCmd.Flags().StringVarP(&_transferCustomHttpHeaders, "custom-http-headers", "", "", "Custom HTTP Headers")
	_transferCmd.Flags().StringVarP(&_transferDeletePath, "delete-path", "", "", "Delete Path")
	_transferCmd.Flags().StringVarP(&_transferDescription, "description", "", "", "Description")
	_transferCmd.Flags().StringVarP(&_transferDomain, "domain", "", "", "Domain")
	_transferCmd.Flags().StringVarP(&_transferEgressConfig, "egress-config", "", "", "Egress Config")
	_transferCmd.Flags().StringVarP(&_transferEndpointDetails, "endpoint-details", "", "", "Endpoint Details")
	_transferCmd.Flags().StringVarP(&_transferEndpointType, "endpoint-type", "", "", "Endpoint Type")
	_transferCmd.Flags().StringVarP(&_transferEnforceMessageSigning, "enforce-message-signing", "", "", "Enforce Message Signing")
	_transferCmd.Flags().StringVarP(&_transferExecutionId, "execution-id", "", "", "Execution ID")
	_transferCmd.Flags().StringVarP(&_transferExternalId, "external-id", "", "", "External ID")
	_transferCmd.Flags().StringVarP(&_transferFaviconFile, "favicon-file", "", "", "Favicon File")
	_transferCmd.Flags().StringVarP(&_transferHomeDirectory, "home-directory", "", "", "Home Directory")
	_transferCmd.Flags().StringVarP(&_transferHomeDirectoryMappings, "home-directory-mappings", "", "", "Home Directory Mappings")
	_transferCmd.Flags().StringVarP(&_transferHomeDirectoryType, "home-directory-type", "", "", "Home Directory Type")
	_transferCmd.Flags().StringVarP(&_transferHostKey, "host-key", "", "", "Host Key")
	_transferCmd.Flags().StringVarP(&_transferHostKeyBody, "host-key-body", "", "", "Host Key Body")
	_transferCmd.Flags().StringVarP(&_transferHostKeyId, "host-key-id", "", "", "Host Key ID")
	_transferCmd.Flags().StringVarP(&_transferIdentityProviderDetails, "identity-provider-details", "", "", "Identity Provider Details")
	_transferCmd.Flags().StringVarP(&_transferIdentityProviderType, "identity-provider-type", "", "", "Identity Provider Type")
	_transferCmd.Flags().StringVarP(&_transferInactiveDate, "inactive-date", "", "", "Inactive Date")
	_transferCmd.Flags().StringVarP(&_transferIpAddressType, "ip-address-type", "", "", "IP Address Type")
	_transferCmd.Flags().StringVarP(&_transferLocalDirectoryPath, "local-directory-path", "", "", "Local Directory Path")
	_transferCmd.Flags().StringVarP(&_transferLocalProfileId, "local-profile-id", "", "", "Local Profile ID")
	_transferCmd.Flags().StringVarP(&_transferLoggingRole, "logging-role", "", "", "Logging Role")
	_transferCmd.Flags().StringVarP(&_transferLogoFile, "logo-file", "", "", "Logo File")
	_transferCmd.Flags().StringVarP(&_transferMaxItems, "max-items", "", "", "Max Items")
	_transferCmd.Flags().StringVarP(&_transferMaxResults, "max-results", "", "", "Max Results")
	_transferCmd.Flags().StringVarP(&_transferNextToken, "next-token", "", "", "Next Token")
	_transferCmd.Flags().StringVarP(&_transferOnExceptionSteps, "on-exception-steps", "", "", "On Exception Steps")
	_transferCmd.Flags().StringVarP(&_transferOutputDirectoryPath, "output-directory-path", "", "", "Output Directory Path")
	_transferCmd.Flags().StringVarP(&_transferPartnerProfileId, "partner-profile-id", "", "", "Partner Profile ID")
	_transferCmd.Flags().StringVarP(&_transferPolicy, "policy", "", "", "Policy")
	_transferCmd.Flags().StringVarP(&_transferPosixProfile, "posix-profile", "", "", "Posix Profile")
	_transferCmd.Flags().StringVarP(&_transferPostAuthenticationLoginBanner, "post-authentication-login-banner", "", "", "Post Authentication Login Banner")
	_transferCmd.Flags().StringVarP(&_transferPreAuthenticationLoginBanner, "pre-authentication-login-banner", "", "", "Pre Authentication Login Banner")
	_transferCmd.Flags().StringVarP(&_transferPreserveFilename, "preserve-filename", "", "", "Preserve Filename")
	_transferCmd.Flags().StringVarP(&_transferPrivateKey, "private-key", "", "", "Private Key")
	_transferCmd.Flags().StringVarP(&_transferProfileId, "profile-id", "", "", "Profile ID")
	_transferCmd.Flags().StringVarP(&_transferProfileType, "profile-type", "", "", "Profile Type")
	_transferCmd.Flags().StringVarP(&_transferProtocolDetails, "protocol-details", "", "", "Protocol Details")
	_transferCmd.Flags().StringVarP(&_transferProtocols, "protocols", "", "", "Protocols")
	_transferCmd.Flags().StringVarP(&_transferRemoteDirectoryPath, "remote-directory-path", "", "", "Remote Directory Path")
	_transferCmd.Flags().StringSliceVarP(&_transferRetrieveFilePaths, "retrieve-file-paths", "", nil, "Retrieve File Paths")
	_transferCmd.Flags().StringVarP(&_transferRole, "role", "", "", "Role")
	_transferCmd.Flags().StringVarP(&_transferS3StorageOptions, "s3-storage-options", "", "", "S3 Storage Options")
	_transferCmd.Flags().StringVarP(&_transferSecurityPolicyName, "security-policy-name", "", "", "Security Policy Name")
	_transferCmd.Flags().StringSliceVarP(&_transferSendFilePaths, "send-file-paths", "", nil, "Send File Paths")
	_transferCmd.Flags().StringVarP(&_transferServerId, "server-id", "", "", "Server ID")
	_transferCmd.Flags().StringVarP(&_transferServerProtocol, "server-protocol", "", "", "Server Protocol")
	_transferCmd.Flags().StringVarP(&_transferSftpConfig, "sftp-config", "", "", "Sftp Config")
	_transferCmd.Flags().StringVarP(&_transferSourceIp, "source-ip", "", "", "Source IP")
	_transferCmd.Flags().StringVarP(&_transferSourcePath, "source-path", "", "", "Source Path")
	_transferCmd.Flags().StringVarP(&_transferSshPublicKeyBody, "ssh-public-key-body", "", "", "SSH Public Key Body")
	_transferCmd.Flags().StringVarP(&_transferSshPublicKeyId, "ssh-public-key-id", "", "", "SSH Public Key ID")
	_transferCmd.Flags().StringVarP(&_transferStatus, "status", "", "", "Status")
	_transferCmd.Flags().StringVarP(&_transferSteps, "steps", "", "", "Steps")
	_transferCmd.Flags().StringSliceVarP(&_transferStructuredLogDestinations, "structured-log-destinations", "", nil, "Structured Log Destinations")
	_transferCmd.Flags().StringSliceVarP(&_transferTagKeys, "tag-keys", "", nil, "Tag Keys")
	_transferCmd.Flags().StringVarP(&_transferTags, "tags", "", "", "Tags")
	_transferCmd.Flags().StringVarP(&_transferTargetPath, "target-path", "", "", "Target Path")
	_transferCmd.Flags().StringVarP(&_transferTitle, "title", "", "", "Title")
	_transferCmd.Flags().StringVarP(&_transferToken, "token", "", "", "Token")
	_transferCmd.Flags().StringVarP(&_transferTransferId, "transfer-id", "", "", "Transfer ID")
	_transferCmd.Flags().StringVarP(&_transferUrl, "url", "", "", "URL")
	_transferCmd.Flags().StringVarP(&_transferUsage, "usage", "", "", "Usage")
	_transferCmd.Flags().StringVarP(&_transferUserName, "user-name", "", "", "User Name")
	_transferCmd.Flags().StringVarP(&_transferUserPassword, "user-password", "", "", "User Password")
	_transferCmd.Flags().StringVarP(&_transferWebAppEndpointPolicy, "web-app-endpoint-policy", "", "", "Web App Endpoint Policy")
	_transferCmd.Flags().StringVarP(&_transferWebAppId, "web-app-id", "", "", "Web App ID")
	_transferCmd.Flags().StringVarP(&_transferWebAppUnits, "web-app-units", "", "", "Web App Units")
	_transferCmd.Flags().StringVarP(&_transferWorkflowDetails, "workflow-details", "", "", "Workflow Details")
	_transferCmd.Flags().StringVarP(&_transferWorkflowId, "workflow-id", "", "", "Workflow ID")

	_transferCmd.Flags().BoolVarP(&_transferCreateAccess, "create-access", "", false, "Create Access")
	_transferCmd.Flags().BoolVarP(&_transferCreateAgreement, "create-agreement", "", false, "Create Agreement")
	_transferCmd.Flags().BoolVarP(&_transferCreateConnector, "create-connector", "", false, "Create Connector")
	_transferCmd.Flags().BoolVarP(&_transferCreateProfile, "create-profile", "", false, "Create Profile")
	_transferCmd.Flags().BoolVarP(&_transferCreateServer, "create-server", "", false, "Create Server")
	_transferCmd.Flags().BoolVarP(&_transferCreateUser, "create-user", "", false, "Create User")
	_transferCmd.Flags().BoolVarP(&_transferCreateWebApp, "create-web-app", "", false, "Create Web App")
	_transferCmd.Flags().BoolVarP(&_transferCreateWorkflow, "create-workflow", "", false, "Create Workflow")
	_transferCmd.Flags().BoolVarP(&_transferDeleteAccess, "delete-access", "", false, "Delete Access")
	_transferCmd.Flags().BoolVarP(&_transferDeleteAgreement, "delete-agreement", "", false, "Delete Agreement")
	_transferCmd.Flags().BoolVarP(&_transferDeleteCertificate, "delete-certificate", "", false, "Delete Certificate")
	_transferCmd.Flags().BoolVarP(&_transferDeleteConnector, "delete-connector", "", false, "Delete Connector")
	_transferCmd.Flags().BoolVarP(&_transferDeleteHostKey, "delete-host-key", "", false, "Delete Host Key")
	_transferCmd.Flags().BoolVarP(&_transferDeleteProfile, "delete-profile", "", false, "Delete Profile")
	_transferCmd.Flags().BoolVarP(&_transferDeleteServer, "delete-server", "", false, "Delete Server")
	_transferCmd.Flags().BoolVarP(&_transferDeleteSshPublicKey, "delete-ssh-public-key", "", false, "Delete SSH Public Key")
	_transferCmd.Flags().BoolVarP(&_transferDeleteUser, "delete-user", "", false, "Delete User")
	_transferCmd.Flags().BoolVarP(&_transferDeleteWebApp, "delete-web-app", "", false, "Delete Web App")
	_transferCmd.Flags().BoolVarP(&_transferDeleteWebAppCustomization, "delete-web-app-customization", "", false, "Delete Web App Customization")
	_transferCmd.Flags().BoolVarP(&_transferDeleteWorkflow, "delete-workflow", "", false, "Delete Workflow")
	_transferCmd.Flags().BoolVarP(&_transferDescribeAccess, "describe-access", "", false, "Describe Access")
	_transferCmd.Flags().BoolVarP(&_transferDescribeAgreement, "describe-agreement", "", false, "Describe Agreement")
	_transferCmd.Flags().BoolVarP(&_transferDescribeCertificate, "describe-certificate", "", false, "Describe Certificate")
	_transferCmd.Flags().BoolVarP(&_transferDescribeConnector, "describe-connector", "", false, "Describe Connector")
	_transferCmd.Flags().BoolVarP(&_transferDescribeExecution, "describe-execution", "", false, "Describe Execution")
	_transferCmd.Flags().BoolVarP(&_transferDescribeHostKey, "describe-host-key", "", false, "Describe Host Key")
	_transferCmd.Flags().BoolVarP(&_transferDescribeProfile, "describe-profile", "", false, "Describe Profile")
	_transferCmd.Flags().BoolVarP(&_transferDescribeSecurityPolicy, "describe-security-policy", "", false, "Describe Security Policy")
	_transferCmd.Flags().BoolVarP(&_transferDescribeServer, "describe-server", "", false, "Describe Server")
	_transferCmd.Flags().BoolVarP(&_transferDescribeUser, "describe-user", "", false, "Describe User")
	_transferCmd.Flags().BoolVarP(&_transferDescribeWebApp, "describe-web-app", "", false, "Describe Web App")
	_transferCmd.Flags().BoolVarP(&_transferDescribeWebAppCustomization, "describe-web-app-customization", "", false, "Describe Web App Customization")
	_transferCmd.Flags().BoolVarP(&_transferDescribeWorkflow, "describe-workflow", "", false, "Describe Workflow")
	_transferCmd.Flags().BoolVarP(&_transferImportCertificate, "import-certificate", "", false, "Import Certificate")
	_transferCmd.Flags().BoolVarP(&_transferImportHostKey, "import-host-key", "", false, "Import Host Key")
	_transferCmd.Flags().BoolVarP(&_transferImportSshPublicKey, "import-ssh-public-key", "", false, "Import SSH Public Key")
	_transferCmd.Flags().BoolVarP(&_transferListAccesses, "list-accesses", "", false, "List Accesses")
	_transferCmd.Flags().BoolVarP(&_transferListAgreements, "list-agreements", "", false, "List Agreements")
	_transferCmd.Flags().BoolVarP(&_transferListCertificates, "list-certificates", "", false, "List Certificates")
	_transferCmd.Flags().BoolVarP(&_transferListConnectors, "list-connectors", "", false, "List Connectors")
	_transferCmd.Flags().BoolVarP(&_transferListExecutions, "list-executions", "", false, "List Executions")
	_transferCmd.Flags().BoolVarP(&_transferListFileTransferResults, "list-file-transfer-results", "", false, "List File Transfer Results")
	_transferCmd.Flags().BoolVarP(&_transferListHostKeys, "list-host-keys", "", false, "List Host Keys")
	_transferCmd.Flags().BoolVarP(&_transferListProfiles, "list-profiles", "", false, "List Profiles")
	_transferCmd.Flags().BoolVarP(&_transferListSecurityPolicies, "list-security-policies", "", false, "List Security Policies")
	_transferCmd.Flags().BoolVarP(&_transferListServers, "list-servers", "", false, "List Servers")
	_transferCmd.Flags().BoolVarP(&_transferListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_transferCmd.Flags().BoolVarP(&_transferListUsers, "list-users", "", false, "List Users")
	_transferCmd.Flags().BoolVarP(&_transferListWebApps, "list-web-apps", "", false, "List Web Apps")
	_transferCmd.Flags().BoolVarP(&_transferListWorkflows, "list-workflows", "", false, "List Workflows")
	_transferCmd.Flags().BoolVarP(&_transferSendWorkflowStepState, "send-workflow-step-state", "", false, "Send Workflow Step State")
	_transferCmd.Flags().BoolVarP(&_transferStartDirectoryListing, "start-directory-listing", "", false, "Start Directory Listing")
	_transferCmd.Flags().BoolVarP(&_transferStartFileTransfer, "start-file-transfer", "", false, "Start File Transfer")
	_transferCmd.Flags().BoolVarP(&_transferStartRemoteDelete, "start-remote-delete", "", false, "Start Remote Delete")
	_transferCmd.Flags().BoolVarP(&_transferStartRemoteMove, "start-remote-move", "", false, "Start Remote Move")
	_transferCmd.Flags().BoolVarP(&_transferStartServer, "start-server", "", false, "Start Server")
	_transferCmd.Flags().BoolVarP(&_transferStopServer, "stop-server", "", false, "Stop Server")
	_transferCmd.Flags().BoolVarP(&_transferTagResource, "tag-resource", "", false, "Tag Resource")
	_transferCmd.Flags().BoolVarP(&_transferTestConnection, "test-connection", "", false, "Test Connection")
	_transferCmd.Flags().BoolVarP(&_transferTestIdentityProvider, "test-identity-provider", "", false, "Test Identity Provider")
	_transferCmd.Flags().BoolVarP(&_transferUntagResource, "untag-resource", "", false, "Untag Resource")
	_transferCmd.Flags().BoolVarP(&_transferUpdateAccess, "update-access", "", false, "Update Access")
	_transferCmd.Flags().BoolVarP(&_transferUpdateAgreement, "update-agreement", "", false, "Update Agreement")
	_transferCmd.Flags().BoolVarP(&_transferUpdateCertificate, "update-certificate", "", false, "Update Certificate")
	_transferCmd.Flags().BoolVarP(&_transferUpdateConnector, "update-connector", "", false, "Update Connector")
	_transferCmd.Flags().BoolVarP(&_transferUpdateHostKey, "update-host-key", "", false, "Update Host Key")
	_transferCmd.Flags().BoolVarP(&_transferUpdateProfile, "update-profile", "", false, "Update Profile")
	_transferCmd.Flags().BoolVarP(&_transferUpdateServer, "update-server", "", false, "Update Server")
	_transferCmd.Flags().BoolVarP(&_transferUpdateUser, "update-user", "", false, "Update User")
	_transferCmd.Flags().BoolVarP(&_transferUpdateWebApp, "update-web-app", "", false, "Update Web App")
	_transferCmd.Flags().BoolVarP(&_transferUpdateWebAppCustomization, "update-web-app-customization", "", false, "Update Web App Customization")

}
