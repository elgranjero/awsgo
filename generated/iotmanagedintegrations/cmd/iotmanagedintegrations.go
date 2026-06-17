package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iotmanagedintegrations"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// iotmanagedintegrationsCmd represents the iotmanagedintegrations command
var _iotmanagedintegrationsCmd = &cobra.Command{
	Use:   "iotmanagedintegrations",
	Short: "AWS iotmanagedintegrations CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := iotmanagedintegrations.NewFromConfig(cfg)
		if _iotmanagedintegrationsCreateAccountAssociation {
			iotmanagedintegrations_CreateAccountAssociation(cfg, client)
			return
		}
		if _iotmanagedintegrationsCreateCloudConnector {
			iotmanagedintegrations_CreateCloudConnector(cfg, client)
			return
		}
		if _iotmanagedintegrationsCreateConnectorDestination {
			iotmanagedintegrations_CreateConnectorDestination(cfg, client)
			return
		}
		if _iotmanagedintegrationsCreateCredentialLocker {
			iotmanagedintegrations_CreateCredentialLocker(cfg, client)
			return
		}
		if _iotmanagedintegrationsCreateDestination {
			iotmanagedintegrations_CreateDestination(cfg, client)
			return
		}
		if _iotmanagedintegrationsCreateEventLogConfiguration {
			iotmanagedintegrations_CreateEventLogConfiguration(cfg, client)
			return
		}
		if _iotmanagedintegrationsCreateManagedThing {
			iotmanagedintegrations_CreateManagedThing(cfg, client)
			return
		}
		if _iotmanagedintegrationsCreateNotificationConfiguration {
			iotmanagedintegrations_CreateNotificationConfiguration(cfg, client)
			return
		}
		if _iotmanagedintegrationsCreateOtaTask {
			iotmanagedintegrations_CreateOtaTask(cfg, client)
			return
		}
		if _iotmanagedintegrationsCreateOtaTaskConfiguration {
			iotmanagedintegrations_CreateOtaTaskConfiguration(cfg, client)
			return
		}
		if _iotmanagedintegrationsCreateProvisioningProfile {
			iotmanagedintegrations_CreateProvisioningProfile(cfg, client)
			return
		}
		if _iotmanagedintegrationsDeleteAccountAssociation {
			iotmanagedintegrations_DeleteAccountAssociation(cfg, client)
			return
		}
		if _iotmanagedintegrationsDeleteCloudConnector {
			iotmanagedintegrations_DeleteCloudConnector(cfg, client)
			return
		}
		if _iotmanagedintegrationsDeleteConnectorDestination {
			iotmanagedintegrations_DeleteConnectorDestination(cfg, client)
			return
		}
		if _iotmanagedintegrationsDeleteCredentialLocker {
			iotmanagedintegrations_DeleteCredentialLocker(cfg, client)
			return
		}
		if _iotmanagedintegrationsDeleteDestination {
			iotmanagedintegrations_DeleteDestination(cfg, client)
			return
		}
		if _iotmanagedintegrationsDeleteEventLogConfiguration {
			iotmanagedintegrations_DeleteEventLogConfiguration(cfg, client)
			return
		}
		if _iotmanagedintegrationsDeleteManagedThing {
			iotmanagedintegrations_DeleteManagedThing(cfg, client)
			return
		}
		if _iotmanagedintegrationsDeleteNotificationConfiguration {
			iotmanagedintegrations_DeleteNotificationConfiguration(cfg, client)
			return
		}
		if _iotmanagedintegrationsDeleteOtaTask {
			iotmanagedintegrations_DeleteOtaTask(cfg, client)
			return
		}
		if _iotmanagedintegrationsDeleteOtaTaskConfiguration {
			iotmanagedintegrations_DeleteOtaTaskConfiguration(cfg, client)
			return
		}
		if _iotmanagedintegrationsDeleteProvisioningProfile {
			iotmanagedintegrations_DeleteProvisioningProfile(cfg, client)
			return
		}
		if _iotmanagedintegrationsDeregisterAccountAssociation {
			iotmanagedintegrations_DeregisterAccountAssociation(cfg, client)
			return
		}
		if _iotmanagedintegrationsGetAccountAssociation {
			iotmanagedintegrations_GetAccountAssociation(cfg, client)
			return
		}
		if _iotmanagedintegrationsGetCloudConnector {
			iotmanagedintegrations_GetCloudConnector(cfg, client)
			return
		}
		if _iotmanagedintegrationsGetConnectorDestination {
			iotmanagedintegrations_GetConnectorDestination(cfg, client)
			return
		}
		if _iotmanagedintegrationsGetCredentialLocker {
			iotmanagedintegrations_GetCredentialLocker(cfg, client)
			return
		}
		if _iotmanagedintegrationsGetCustomEndpoint {
			iotmanagedintegrations_GetCustomEndpoint(cfg, client)
			return
		}
		if _iotmanagedintegrationsGetDefaultEncryptionConfiguration {
			iotmanagedintegrations_GetDefaultEncryptionConfiguration(cfg, client)
			return
		}
		if _iotmanagedintegrationsGetDestination {
			iotmanagedintegrations_GetDestination(cfg, client)
			return
		}
		if _iotmanagedintegrationsGetDeviceDiscovery {
			iotmanagedintegrations_GetDeviceDiscovery(cfg, client)
			return
		}
		if _iotmanagedintegrationsGetEventLogConfiguration {
			iotmanagedintegrations_GetEventLogConfiguration(cfg, client)
			return
		}
		if _iotmanagedintegrationsGetHubConfiguration {
			iotmanagedintegrations_GetHubConfiguration(cfg, client)
			return
		}
		if _iotmanagedintegrationsGetManagedThing {
			iotmanagedintegrations_GetManagedThing(cfg, client)
			return
		}
		if _iotmanagedintegrationsGetManagedThingCapabilities {
			iotmanagedintegrations_GetManagedThingCapabilities(cfg, client)
			return
		}
		if _iotmanagedintegrationsGetManagedThingCertificate {
			iotmanagedintegrations_GetManagedThingCertificate(cfg, client)
			return
		}
		if _iotmanagedintegrationsGetManagedThingConnectivityData {
			iotmanagedintegrations_GetManagedThingConnectivityData(cfg, client)
			return
		}
		if _iotmanagedintegrationsGetManagedThingMetaData {
			iotmanagedintegrations_GetManagedThingMetaData(cfg, client)
			return
		}
		if _iotmanagedintegrationsGetManagedThingState {
			iotmanagedintegrations_GetManagedThingState(cfg, client)
			return
		}
		if _iotmanagedintegrationsGetNotificationConfiguration {
			iotmanagedintegrations_GetNotificationConfiguration(cfg, client)
			return
		}
		if _iotmanagedintegrationsGetOtaTask {
			iotmanagedintegrations_GetOtaTask(cfg, client)
			return
		}
		if _iotmanagedintegrationsGetOtaTaskConfiguration {
			iotmanagedintegrations_GetOtaTaskConfiguration(cfg, client)
			return
		}
		if _iotmanagedintegrationsGetProvisioningProfile {
			iotmanagedintegrations_GetProvisioningProfile(cfg, client)
			return
		}
		if _iotmanagedintegrationsGetRuntimeLogConfiguration {
			iotmanagedintegrations_GetRuntimeLogConfiguration(cfg, client)
			return
		}
		if _iotmanagedintegrationsGetSchemaVersion {
			iotmanagedintegrations_GetSchemaVersion(cfg, client)
			return
		}
		if _iotmanagedintegrationsListAccountAssociations {
			iotmanagedintegrations_ListAccountAssociations(cfg, client)
			return
		}
		if _iotmanagedintegrationsListCloudConnectors {
			iotmanagedintegrations_ListCloudConnectors(cfg, client)
			return
		}
		if _iotmanagedintegrationsListConnectorDestinations {
			iotmanagedintegrations_ListConnectorDestinations(cfg, client)
			return
		}
		if _iotmanagedintegrationsListCredentialLockers {
			iotmanagedintegrations_ListCredentialLockers(cfg, client)
			return
		}
		if _iotmanagedintegrationsListDestinations {
			iotmanagedintegrations_ListDestinations(cfg, client)
			return
		}
		if _iotmanagedintegrationsListDeviceDiscoveries {
			iotmanagedintegrations_ListDeviceDiscoveries(cfg, client)
			return
		}
		if _iotmanagedintegrationsListDiscoveredDevices {
			iotmanagedintegrations_ListDiscoveredDevices(cfg, client)
			return
		}
		if _iotmanagedintegrationsListEventLogConfigurations {
			iotmanagedintegrations_ListEventLogConfigurations(cfg, client)
			return
		}
		if _iotmanagedintegrationsListManagedThingAccountAssociations {
			iotmanagedintegrations_ListManagedThingAccountAssociations(cfg, client)
			return
		}
		if _iotmanagedintegrationsListManagedThingSchemas {
			iotmanagedintegrations_ListManagedThingSchemas(cfg, client)
			return
		}
		if _iotmanagedintegrationsListManagedThings {
			iotmanagedintegrations_ListManagedThings(cfg, client)
			return
		}
		if _iotmanagedintegrationsListNotificationConfigurations {
			iotmanagedintegrations_ListNotificationConfigurations(cfg, client)
			return
		}
		if _iotmanagedintegrationsListOtaTaskConfigurations {
			iotmanagedintegrations_ListOtaTaskConfigurations(cfg, client)
			return
		}
		if _iotmanagedintegrationsListOtaTaskExecutions {
			iotmanagedintegrations_ListOtaTaskExecutions(cfg, client)
			return
		}
		if _iotmanagedintegrationsListOtaTasks {
			iotmanagedintegrations_ListOtaTasks(cfg, client)
			return
		}
		if _iotmanagedintegrationsListProvisioningProfiles {
			iotmanagedintegrations_ListProvisioningProfiles(cfg, client)
			return
		}
		if _iotmanagedintegrationsListSchemaVersions {
			iotmanagedintegrations_ListSchemaVersions(cfg, client)
			return
		}
		if _iotmanagedintegrationsListTagsForResource {
			iotmanagedintegrations_ListTagsForResource(cfg, client)
			return
		}
		if _iotmanagedintegrationsPutDefaultEncryptionConfiguration {
			iotmanagedintegrations_PutDefaultEncryptionConfiguration(cfg, client)
			return
		}
		if _iotmanagedintegrationsPutHubConfiguration {
			iotmanagedintegrations_PutHubConfiguration(cfg, client)
			return
		}
		if _iotmanagedintegrationsPutRuntimeLogConfiguration {
			iotmanagedintegrations_PutRuntimeLogConfiguration(cfg, client)
			return
		}
		if _iotmanagedintegrationsRegisterAccountAssociation {
			iotmanagedintegrations_RegisterAccountAssociation(cfg, client)
			return
		}
		if _iotmanagedintegrationsRegisterCustomEndpoint {
			iotmanagedintegrations_RegisterCustomEndpoint(cfg, client)
			return
		}
		if _iotmanagedintegrationsResetRuntimeLogConfiguration {
			iotmanagedintegrations_ResetRuntimeLogConfiguration(cfg, client)
			return
		}
		if _iotmanagedintegrationsSendConnectorEvent {
			iotmanagedintegrations_SendConnectorEvent(cfg, client)
			return
		}
		if _iotmanagedintegrationsSendManagedThingCommand {
			iotmanagedintegrations_SendManagedThingCommand(cfg, client)
			return
		}
		if _iotmanagedintegrationsStartAccountAssociationRefresh {
			iotmanagedintegrations_StartAccountAssociationRefresh(cfg, client)
			return
		}
		if _iotmanagedintegrationsStartDeviceDiscovery {
			iotmanagedintegrations_StartDeviceDiscovery(cfg, client)
			return
		}
		if _iotmanagedintegrationsTagResource {
			iotmanagedintegrations_TagResource(cfg, client)
			return
		}
		if _iotmanagedintegrationsUntagResource {
			iotmanagedintegrations_UntagResource(cfg, client)
			return
		}
		if _iotmanagedintegrationsUpdateAccountAssociation {
			iotmanagedintegrations_UpdateAccountAssociation(cfg, client)
			return
		}
		if _iotmanagedintegrationsUpdateCloudConnector {
			iotmanagedintegrations_UpdateCloudConnector(cfg, client)
			return
		}
		if _iotmanagedintegrationsUpdateConnectorDestination {
			iotmanagedintegrations_UpdateConnectorDestination(cfg, client)
			return
		}
		if _iotmanagedintegrationsUpdateDestination {
			iotmanagedintegrations_UpdateDestination(cfg, client)
			return
		}
		if _iotmanagedintegrationsUpdateEventLogConfiguration {
			iotmanagedintegrations_UpdateEventLogConfiguration(cfg, client)
			return
		}
		if _iotmanagedintegrationsUpdateManagedThing {
			iotmanagedintegrations_UpdateManagedThing(cfg, client)
			return
		}
		if _iotmanagedintegrationsUpdateNotificationConfiguration {
			iotmanagedintegrations_UpdateNotificationConfiguration(cfg, client)
			return
		}
		if _iotmanagedintegrationsUpdateOtaTask {
			iotmanagedintegrations_UpdateOtaTask(cfg, client)
			return
		}

	},
}

var (
	_iotmanagedintegrationsCreateAccountAssociation            bool
	_iotmanagedintegrationsCreateCloudConnector                bool
	_iotmanagedintegrationsCreateConnectorDestination          bool
	_iotmanagedintegrationsCreateCredentialLocker              bool
	_iotmanagedintegrationsCreateDestination                   bool
	_iotmanagedintegrationsCreateEventLogConfiguration         bool
	_iotmanagedintegrationsCreateManagedThing                  bool
	_iotmanagedintegrationsCreateNotificationConfiguration     bool
	_iotmanagedintegrationsCreateOtaTask                       bool
	_iotmanagedintegrationsCreateOtaTaskConfiguration          bool
	_iotmanagedintegrationsCreateProvisioningProfile           bool
	_iotmanagedintegrationsDeleteAccountAssociation            bool
	_iotmanagedintegrationsDeleteCloudConnector                bool
	_iotmanagedintegrationsDeleteConnectorDestination          bool
	_iotmanagedintegrationsDeleteCredentialLocker              bool
	_iotmanagedintegrationsDeleteDestination                   bool
	_iotmanagedintegrationsDeleteEventLogConfiguration         bool
	_iotmanagedintegrationsDeleteManagedThing                  bool
	_iotmanagedintegrationsDeleteNotificationConfiguration     bool
	_iotmanagedintegrationsDeleteOtaTask                       bool
	_iotmanagedintegrationsDeleteOtaTaskConfiguration          bool
	_iotmanagedintegrationsDeleteProvisioningProfile           bool
	_iotmanagedintegrationsDeregisterAccountAssociation        bool
	_iotmanagedintegrationsGetAccountAssociation               bool
	_iotmanagedintegrationsGetCloudConnector                   bool
	_iotmanagedintegrationsGetConnectorDestination             bool
	_iotmanagedintegrationsGetCredentialLocker                 bool
	_iotmanagedintegrationsGetCustomEndpoint                   bool
	_iotmanagedintegrationsGetDefaultEncryptionConfiguration   bool
	_iotmanagedintegrationsGetDestination                      bool
	_iotmanagedintegrationsGetDeviceDiscovery                  bool
	_iotmanagedintegrationsGetEventLogConfiguration            bool
	_iotmanagedintegrationsGetHubConfiguration                 bool
	_iotmanagedintegrationsGetManagedThing                     bool
	_iotmanagedintegrationsGetManagedThingCapabilities         bool
	_iotmanagedintegrationsGetManagedThingCertificate          bool
	_iotmanagedintegrationsGetManagedThingConnectivityData     bool
	_iotmanagedintegrationsGetManagedThingMetaData             bool
	_iotmanagedintegrationsGetManagedThingState                bool
	_iotmanagedintegrationsGetNotificationConfiguration        bool
	_iotmanagedintegrationsGetOtaTask                          bool
	_iotmanagedintegrationsGetOtaTaskConfiguration             bool
	_iotmanagedintegrationsGetProvisioningProfile              bool
	_iotmanagedintegrationsGetRuntimeLogConfiguration          bool
	_iotmanagedintegrationsGetSchemaVersion                    bool
	_iotmanagedintegrationsListAccountAssociations             bool
	_iotmanagedintegrationsListCloudConnectors                 bool
	_iotmanagedintegrationsListConnectorDestinations           bool
	_iotmanagedintegrationsListCredentialLockers               bool
	_iotmanagedintegrationsListDestinations                    bool
	_iotmanagedintegrationsListDeviceDiscoveries               bool
	_iotmanagedintegrationsListDiscoveredDevices               bool
	_iotmanagedintegrationsListEventLogConfigurations          bool
	_iotmanagedintegrationsListManagedThingAccountAssociations bool
	_iotmanagedintegrationsListManagedThingSchemas             bool
	_iotmanagedintegrationsListManagedThings                   bool
	_iotmanagedintegrationsListNotificationConfigurations      bool
	_iotmanagedintegrationsListOtaTaskConfigurations           bool
	_iotmanagedintegrationsListOtaTaskExecutions               bool
	_iotmanagedintegrationsListOtaTasks                        bool
	_iotmanagedintegrationsListProvisioningProfiles            bool
	_iotmanagedintegrationsListSchemaVersions                  bool
	_iotmanagedintegrationsListTagsForResource                 bool
	_iotmanagedintegrationsPutDefaultEncryptionConfiguration   bool
	_iotmanagedintegrationsPutHubConfiguration                 bool
	_iotmanagedintegrationsPutRuntimeLogConfiguration          bool
	_iotmanagedintegrationsRegisterAccountAssociation          bool
	_iotmanagedintegrationsRegisterCustomEndpoint              bool
	_iotmanagedintegrationsResetRuntimeLogConfiguration        bool
	_iotmanagedintegrationsSendConnectorEvent                  bool
	_iotmanagedintegrationsSendManagedThingCommand             bool
	_iotmanagedintegrationsStartAccountAssociationRefresh      bool
	_iotmanagedintegrationsStartDeviceDiscovery                bool
	_iotmanagedintegrationsTagResource                         bool
	_iotmanagedintegrationsUntagResource                       bool
	_iotmanagedintegrationsUpdateAccountAssociation            bool
	_iotmanagedintegrationsUpdateCloudConnector                bool
	_iotmanagedintegrationsUpdateConnectorDestination          bool
	_iotmanagedintegrationsUpdateDestination                   bool
	_iotmanagedintegrationsUpdateEventLogConfiguration         bool
	_iotmanagedintegrationsUpdateManagedThing                  bool
	_iotmanagedintegrationsUpdateNotificationConfiguration     bool
	_iotmanagedintegrationsUpdateOtaTask                       bool

	_iotmanagedintegrationsAccountAssociationId                string
	_iotmanagedintegrationsAuthConfig                          string
	_iotmanagedintegrationsAuthType                            string
	_iotmanagedintegrationsAuthenticationMaterial              string
	_iotmanagedintegrationsAuthenticationMaterialType          string
	_iotmanagedintegrationsBrand                               string
	_iotmanagedintegrationsCaCertificate                       string
	_iotmanagedintegrationsCapabilities                        string
	_iotmanagedintegrationsCapabilityIdFilter                  string
	_iotmanagedintegrationsCapabilityReport                    string
	_iotmanagedintegrationsCapabilitySchemas                   string
	_iotmanagedintegrationsClaimCertificate                    string
	_iotmanagedintegrationsClassification                      string
	_iotmanagedintegrationsClientToken                         string
	_iotmanagedintegrationsCloudConnectorId                    string
	_iotmanagedintegrationsConnectorAssociationId              string
	_iotmanagedintegrationsConnectorAssociationIdentifier      string
	_iotmanagedintegrationsConnectorDestinationId              string
	_iotmanagedintegrationsConnectorDestinationIdFilter        string
	_iotmanagedintegrationsConnectorDeviceId                   string
	_iotmanagedintegrationsConnectorDeviceIdFilter             string
	_iotmanagedintegrationsConnectorDeviceIdList               []string
	_iotmanagedintegrationsConnectorId                         string
	_iotmanagedintegrationsConnectorPolicyIdFilter             string
	_iotmanagedintegrationsControllerIdentifier                string
	_iotmanagedintegrationsCredentialLockerFilter              string
	_iotmanagedintegrationsCredentialLockerId                  string
	_iotmanagedintegrationsCustomProtocolDetail                string
	_iotmanagedintegrationsDeliveryDestinationArn              string
	_iotmanagedintegrationsDeliveryDestinationType             string
	_iotmanagedintegrationsDescription                         string
	_iotmanagedintegrationsDestinationName                     string
	_iotmanagedintegrationsDeviceDiscoveryId                   string
	_iotmanagedintegrationsDevices                             string
	_iotmanagedintegrationsDiscoveryType                       string
	_iotmanagedintegrationsEncryptionType                      string
	_iotmanagedintegrationsEndDeviceIdentifier                 string
	_iotmanagedintegrationsEndpointConfig                      string
	_iotmanagedintegrationsEndpointIdFilter                    string
	_iotmanagedintegrationsEndpointType                        string
	_iotmanagedintegrationsEndpoints                           string
	_iotmanagedintegrationsEventLogLevel                       string
	_iotmanagedintegrationsEventType                           string
	_iotmanagedintegrationsForce                               string
	_iotmanagedintegrationsFormat                              string
	_iotmanagedintegrationsGeneralAuthorization                string
	_iotmanagedintegrationsHubNetworkMode                      string
	_iotmanagedintegrationsHubTokenTimerExpirySettingInSeconds string
	_iotmanagedintegrationsId                                  string
	_iotmanagedintegrationsIdentifier                          string
	_iotmanagedintegrationsKmsKeyArn                           string
	_iotmanagedintegrationsLambdaArn                           string
	_iotmanagedintegrationsManagedThingId                      string
	_iotmanagedintegrationsMatterEndpoint                      string
	_iotmanagedintegrationsMaxResults                          string
	_iotmanagedintegrationsMessage                             string
	_iotmanagedintegrationsMetaData                            string
	_iotmanagedintegrationsModel                               string
	_iotmanagedintegrationsName                                string
	_iotmanagedintegrationsNamespace                           string
	_iotmanagedintegrationsNextToken                           string
	_iotmanagedintegrationsOperation                           string
	_iotmanagedintegrationsOperationVersion                    string
	_iotmanagedintegrationsOtaMechanism                        string
	_iotmanagedintegrationsOtaSchedulingConfig                 string
	_iotmanagedintegrationsOtaTargetQueryString                string
	_iotmanagedintegrationsOtaTaskExecutionRetryConfig         string
	_iotmanagedintegrationsOtaType                             string
	_iotmanagedintegrationsOwner                               string
	_iotmanagedintegrationsOwnerFilter                         string
	_iotmanagedintegrationsParentControllerIdentifierFilter    string
	_iotmanagedintegrationsProtocol                            string
	_iotmanagedintegrationsProvisioningStatusFilter            string
	_iotmanagedintegrationsProvisioningType                    string
	_iotmanagedintegrationsPushConfig                          string
	_iotmanagedintegrationsResourceArn                         string
	_iotmanagedintegrationsResourceId                          string
	_iotmanagedintegrationsResourceType                        string
	_iotmanagedintegrationsRole                                string
	_iotmanagedintegrationsRoleArn                             string
	_iotmanagedintegrationsRoleFilter                          string
	_iotmanagedintegrationsRuntimeLogConfigurations            string
	_iotmanagedintegrationsS3Url                               string
	_iotmanagedintegrationsSchemaId                            string
	_iotmanagedintegrationsSchemaVersionedId                   string
	_iotmanagedintegrationsSecretsManager                      string
	_iotmanagedintegrationsSemanticVersion                     string
	_iotmanagedintegrationsSerialNumber                        string
	_iotmanagedintegrationsSerialNumberFilter                  string
	_iotmanagedintegrationsStatusCode                          string
	_iotmanagedintegrationsStatusFilter                        string
	_iotmanagedintegrationsTagKeys                             []string
	_iotmanagedintegrationsTags                                string
	_iotmanagedintegrationsTarget                              []string
	_iotmanagedintegrationsTaskConfigurationId                 string
	_iotmanagedintegrationsTraceId                             string
	_iotmanagedintegrationsType                                string
	_iotmanagedintegrationsTypeFilter                          string
	_iotmanagedintegrationsUserId                              string
	_iotmanagedintegrationsVisibility                          string
	_iotmanagedintegrationsWiFiSimpleSetupConfiguration        string
)

// Creates a new account association via the destination id.
func iotmanagedintegrations_CreateAccountAssociation(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.CreateAccountAssociationInput{
		// ConnectorDestinationId: *string, // Required
	}

	if len(_iotmanagedintegrationsConnectorDestinationId) > 0 {
		input.ConnectorDestinationId = aws.String(_iotmanagedintegrationsConnectorDestinationId)
	}
	if len(_iotmanagedintegrationsClientToken) > 0 {
		input.ClientToken = aws.String(_iotmanagedintegrationsClientToken)
	}
	if len(_iotmanagedintegrationsDescription) > 0 {
		input.Description = aws.String(_iotmanagedintegrationsDescription)
	}
	if len(_iotmanagedintegrationsGeneralAuthorization) > 0 {
		if err := assignInputField(input, "GeneralAuthorization", _iotmanagedintegrationsGeneralAuthorization); err != nil {
			log.Errorf("invalid --general-authorization: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsName) > 0 {
		input.Name = aws.String(_iotmanagedintegrationsName)
	}
	if len(_iotmanagedintegrationsTags) > 0 {
		if err := assignInputField(input, "Tags", _iotmanagedintegrationsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAccountAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a C2C (cloud-to-cloud) connector.
func iotmanagedintegrations_CreateCloudConnector(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.CreateCloudConnectorInput{
		// EndpointConfig: *types.EndpointConfig, // Required
		// Name: *string, // Required
	}

	if len(_iotmanagedintegrationsEndpointConfig) > 0 {
		if err := assignInputField(input, "EndpointConfig", _iotmanagedintegrationsEndpointConfig); err != nil {
			log.Errorf("invalid --endpoint-config: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsName) > 0 {
		input.Name = aws.String(_iotmanagedintegrationsName)
	}
	if len(_iotmanagedintegrationsClientToken) > 0 {
		input.ClientToken = aws.String(_iotmanagedintegrationsClientToken)
	}
	if len(_iotmanagedintegrationsDescription) > 0 {
		input.Description = aws.String(_iotmanagedintegrationsDescription)
	}
	if len(_iotmanagedintegrationsEndpointType) > 0 {
		if err := assignInputField(input, "EndpointType", _iotmanagedintegrationsEndpointType); err != nil {
			log.Errorf("invalid --endpoint-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCloudConnector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a connector destination for connecting a cloud-to-cloud (C2C) connector
// to the customer's Amazon Web Services account.
func iotmanagedintegrations_CreateConnectorDestination(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.CreateConnectorDestinationInput{
		// AuthConfig: *types.AuthConfig, // Required
		// CloudConnectorId: *string, // Required
	}

	if len(_iotmanagedintegrationsAuthConfig) > 0 {
		if err := assignInputField(input, "AuthConfig", _iotmanagedintegrationsAuthConfig); err != nil {
			log.Errorf("invalid --auth-config: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsCloudConnectorId) > 0 {
		input.CloudConnectorId = aws.String(_iotmanagedintegrationsCloudConnectorId)
	}
	if len(_iotmanagedintegrationsAuthType) > 0 {
		if err := assignInputField(input, "AuthType", _iotmanagedintegrationsAuthType); err != nil {
			log.Errorf("invalid --auth-type: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsClientToken) > 0 {
		input.ClientToken = aws.String(_iotmanagedintegrationsClientToken)
	}
	if len(_iotmanagedintegrationsDescription) > 0 {
		input.Description = aws.String(_iotmanagedintegrationsDescription)
	}
	if len(_iotmanagedintegrationsName) > 0 {
		input.Name = aws.String(_iotmanagedintegrationsName)
	}
	if len(_iotmanagedintegrationsSecretsManager) > 0 {
		if err := assignInputField(input, "SecretsManager", _iotmanagedintegrationsSecretsManager); err != nil {
			log.Errorf("invalid --secrets-manager: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateConnectorDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a credential locker.
// This operation will not trigger the creation of all the manufacturing resources.
func iotmanagedintegrations_CreateCredentialLocker(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.CreateCredentialLockerInput{}

	if len(_iotmanagedintegrationsClientToken) > 0 {
		input.ClientToken = aws.String(_iotmanagedintegrationsClientToken)
	}
	if len(_iotmanagedintegrationsName) > 0 {
		input.Name = aws.String(_iotmanagedintegrationsName)
	}
	if len(_iotmanagedintegrationsTags) > 0 {
		if err := assignInputField(input, "Tags", _iotmanagedintegrationsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCredentialLocker(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a notification destination such as Kinesis Data Streams that receive
// events and notifications from Managed integrations. Managed integrations uses
// the destination to determine where to deliver notifications.
func iotmanagedintegrations_CreateDestination(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.CreateDestinationInput{
		// DeliveryDestinationArn: *string, // Required
		// DeliveryDestinationType: types.DeliveryDestinationType, // Required
		// Name: *string, // Required
		// RoleArn: *string, // Required
	}

	if len(_iotmanagedintegrationsDeliveryDestinationArn) > 0 {
		input.DeliveryDestinationArn = aws.String(_iotmanagedintegrationsDeliveryDestinationArn)
	}
	if len(_iotmanagedintegrationsDeliveryDestinationType) > 0 {
		if err := assignInputField(input, "DeliveryDestinationType", _iotmanagedintegrationsDeliveryDestinationType); err != nil {
			log.Errorf("invalid --delivery-destination-type: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsName) > 0 {
		input.Name = aws.String(_iotmanagedintegrationsName)
	}
	if len(_iotmanagedintegrationsRoleArn) > 0 {
		input.RoleArn = aws.String(_iotmanagedintegrationsRoleArn)
	}
	if len(_iotmanagedintegrationsClientToken) > 0 {
		input.ClientToken = aws.String(_iotmanagedintegrationsClientToken)
	}
	if len(_iotmanagedintegrationsDescription) > 0 {
		input.Description = aws.String(_iotmanagedintegrationsDescription)
	}
	if len(_iotmanagedintegrationsTags) > 0 {
		if err := assignInputField(input, "Tags", _iotmanagedintegrationsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Set the event log configuration for the account, resource type, or specific
// resource.
func iotmanagedintegrations_CreateEventLogConfiguration(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.CreateEventLogConfigurationInput{
		// EventLogLevel: types.LogLevel, // Required
		// ResourceType: *string, // Required
	}

	if len(_iotmanagedintegrationsEventLogLevel) > 0 {
		if err := assignInputField(input, "EventLogLevel", _iotmanagedintegrationsEventLogLevel); err != nil {
			log.Errorf("invalid --event-log-level: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsResourceType) > 0 {
		input.ResourceType = aws.String(_iotmanagedintegrationsResourceType)
	}
	if len(_iotmanagedintegrationsClientToken) > 0 {
		input.ClientToken = aws.String(_iotmanagedintegrationsClientToken)
	}
	if len(_iotmanagedintegrationsResourceId) > 0 {
		input.ResourceId = aws.String(_iotmanagedintegrationsResourceId)
	}

	if resp, err := client.CreateEventLogConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a managed thing. A managed thing contains the device identifier,
// protocol supported, and capabilities of the device in a data model format
// defined by Managed integrations.
func iotmanagedintegrations_CreateManagedThing(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.CreateManagedThingInput{
		// AuthenticationMaterial: *string, // Required
		// AuthenticationMaterialType: types.AuthMaterialType, // Required
		// Role: types.Role, // Required
	}

	if len(_iotmanagedintegrationsAuthenticationMaterial) > 0 {
		input.AuthenticationMaterial = aws.String(_iotmanagedintegrationsAuthenticationMaterial)
	}
	if len(_iotmanagedintegrationsAuthenticationMaterialType) > 0 {
		if err := assignInputField(input, "AuthenticationMaterialType", _iotmanagedintegrationsAuthenticationMaterialType); err != nil {
			log.Errorf("invalid --authentication-material-type: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsRole) > 0 {
		if err := assignInputField(input, "Role", _iotmanagedintegrationsRole); err != nil {
			log.Errorf("invalid --role: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsBrand) > 0 {
		input.Brand = aws.String(_iotmanagedintegrationsBrand)
	}
	if len(_iotmanagedintegrationsCapabilities) > 0 {
		input.Capabilities = aws.String(_iotmanagedintegrationsCapabilities)
	}
	if len(_iotmanagedintegrationsCapabilityReport) > 0 {
		if err := assignInputField(input, "CapabilityReport", _iotmanagedintegrationsCapabilityReport); err != nil {
			log.Errorf("invalid --capability-report: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsCapabilitySchemas) > 0 {
		if err := assignInputField(input, "CapabilitySchemas", _iotmanagedintegrationsCapabilitySchemas); err != nil {
			log.Errorf("invalid --capability-schemas: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsClassification) > 0 {
		input.Classification = aws.String(_iotmanagedintegrationsClassification)
	}
	if len(_iotmanagedintegrationsClientToken) > 0 {
		input.ClientToken = aws.String(_iotmanagedintegrationsClientToken)
	}
	if len(_iotmanagedintegrationsCredentialLockerId) > 0 {
		input.CredentialLockerId = aws.String(_iotmanagedintegrationsCredentialLockerId)
	}
	if len(_iotmanagedintegrationsMetaData) > 0 {
		if err := assignInputField(input, "MetaData", _iotmanagedintegrationsMetaData); err != nil {
			log.Errorf("invalid --meta-data: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsModel) > 0 {
		input.Model = aws.String(_iotmanagedintegrationsModel)
	}
	if len(_iotmanagedintegrationsName) > 0 {
		input.Name = aws.String(_iotmanagedintegrationsName)
	}
	if len(_iotmanagedintegrationsOwner) > 0 {
		input.Owner = aws.String(_iotmanagedintegrationsOwner)
	}
	if len(_iotmanagedintegrationsSerialNumber) > 0 {
		input.SerialNumber = aws.String(_iotmanagedintegrationsSerialNumber)
	}
	if len(_iotmanagedintegrationsTags) > 0 {
		if err := assignInputField(input, "Tags", _iotmanagedintegrationsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsWiFiSimpleSetupConfiguration) > 0 {
		if err := assignInputField(input, "WiFiSimpleSetupConfiguration", _iotmanagedintegrationsWiFiSimpleSetupConfiguration); err != nil {
			log.Errorf("invalid --wi-fi-simple-setup-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateManagedThing(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a notification configuration. A configuration is a connection between
// an event type and a destination that you have already created.
func iotmanagedintegrations_CreateNotificationConfiguration(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.CreateNotificationConfigurationInput{
		// DestinationName: *string, // Required
		// EventType: types.EventType, // Required
	}

	if len(_iotmanagedintegrationsDestinationName) > 0 {
		input.DestinationName = aws.String(_iotmanagedintegrationsDestinationName)
	}
	if len(_iotmanagedintegrationsEventType) > 0 {
		if err := assignInputField(input, "EventType", _iotmanagedintegrationsEventType); err != nil {
			log.Errorf("invalid --event-type: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsClientToken) > 0 {
		input.ClientToken = aws.String(_iotmanagedintegrationsClientToken)
	}
	if len(_iotmanagedintegrationsTags) > 0 {
		if err := assignInputField(input, "Tags", _iotmanagedintegrationsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateNotificationConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create an over-the-air (OTA) task to target a device.
func iotmanagedintegrations_CreateOtaTask(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.CreateOtaTaskInput{
		// OtaType: types.OtaType, // Required
		// S3Url: *string, // Required
	}

	if len(_iotmanagedintegrationsOtaType) > 0 {
		if err := assignInputField(input, "OtaType", _iotmanagedintegrationsOtaType); err != nil {
			log.Errorf("invalid --ota-type: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsS3Url) > 0 {
		input.S3Url = aws.String(_iotmanagedintegrationsS3Url)
	}
	if len(_iotmanagedintegrationsClientToken) > 0 {
		input.ClientToken = aws.String(_iotmanagedintegrationsClientToken)
	}
	if len(_iotmanagedintegrationsDescription) > 0 {
		input.Description = aws.String(_iotmanagedintegrationsDescription)
	}
	if len(_iotmanagedintegrationsOtaMechanism) > 0 {
		if err := assignInputField(input, "OtaMechanism", _iotmanagedintegrationsOtaMechanism); err != nil {
			log.Errorf("invalid --ota-mechanism: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsOtaSchedulingConfig) > 0 {
		if err := assignInputField(input, "OtaSchedulingConfig", _iotmanagedintegrationsOtaSchedulingConfig); err != nil {
			log.Errorf("invalid --ota-scheduling-config: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsOtaTargetQueryString) > 0 {
		input.OtaTargetQueryString = aws.String(_iotmanagedintegrationsOtaTargetQueryString)
	}
	if len(_iotmanagedintegrationsOtaTaskExecutionRetryConfig) > 0 {
		if err := assignInputField(input, "OtaTaskExecutionRetryConfig", _iotmanagedintegrationsOtaTaskExecutionRetryConfig); err != nil {
			log.Errorf("invalid --ota-task-execution-retry-config: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsProtocol) > 0 {
		if err := assignInputField(input, "Protocol", _iotmanagedintegrationsProtocol); err != nil {
			log.Errorf("invalid --protocol: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsTags) > 0 {
		if err := assignInputField(input, "Tags", _iotmanagedintegrationsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsTarget) > 0 {
		input.Target = append([]string(nil), _iotmanagedintegrationsTarget...)
	}
	if len(_iotmanagedintegrationsTaskConfigurationId) > 0 {
		input.TaskConfigurationId = aws.String(_iotmanagedintegrationsTaskConfigurationId)
	}

	if resp, err := client.CreateOtaTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a configuraiton for the over-the-air (OTA) task.
func iotmanagedintegrations_CreateOtaTaskConfiguration(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.CreateOtaTaskConfigurationInput{}

	if len(_iotmanagedintegrationsClientToken) > 0 {
		input.ClientToken = aws.String(_iotmanagedintegrationsClientToken)
	}
	if len(_iotmanagedintegrationsDescription) > 0 {
		input.Description = aws.String(_iotmanagedintegrationsDescription)
	}
	if len(_iotmanagedintegrationsName) > 0 {
		input.Name = aws.String(_iotmanagedintegrationsName)
	}
	if len(_iotmanagedintegrationsPushConfig) > 0 {
		if err := assignInputField(input, "PushConfig", _iotmanagedintegrationsPushConfig); err != nil {
			log.Errorf("invalid --push-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateOtaTaskConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a provisioning profile for a device to execute the provisioning flows
// using a provisioning template. The provisioning template is a document that
// defines the set of resources and policies applied to a device during the
// provisioning process.
func iotmanagedintegrations_CreateProvisioningProfile(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.CreateProvisioningProfileInput{
		// ProvisioningType: types.ProvisioningType, // Required
	}

	if len(_iotmanagedintegrationsProvisioningType) > 0 {
		if err := assignInputField(input, "ProvisioningType", _iotmanagedintegrationsProvisioningType); err != nil {
			log.Errorf("invalid --provisioning-type: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsCaCertificate) > 0 {
		input.CaCertificate = aws.String(_iotmanagedintegrationsCaCertificate)
	}
	if len(_iotmanagedintegrationsClaimCertificate) > 0 {
		input.ClaimCertificate = aws.String(_iotmanagedintegrationsClaimCertificate)
	}
	if len(_iotmanagedintegrationsClientToken) > 0 {
		input.ClientToken = aws.String(_iotmanagedintegrationsClientToken)
	}
	if len(_iotmanagedintegrationsName) > 0 {
		input.Name = aws.String(_iotmanagedintegrationsName)
	}
	if len(_iotmanagedintegrationsTags) > 0 {
		if err := assignInputField(input, "Tags", _iotmanagedintegrationsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateProvisioningProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Remove a third-party account association for an end user.
// You must first call the DeregisterAccountAssociation to remove the connection
// between the managed thing and the third-party account before calling the
// DeleteAccountAssociation API.
func iotmanagedintegrations_DeleteAccountAssociation(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.DeleteAccountAssociationInput{
		// AccountAssociationId: *string, // Required
	}

	if len(_iotmanagedintegrationsAccountAssociationId) > 0 {
		input.AccountAssociationId = aws.String(_iotmanagedintegrationsAccountAssociationId)
	}

	if resp, err := client.DeleteAccountAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a cloud connector.
func iotmanagedintegrations_DeleteCloudConnector(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.DeleteCloudConnectorInput{
		// Identifier: *string, // Required
	}

	if len(_iotmanagedintegrationsIdentifier) > 0 {
		input.Identifier = aws.String(_iotmanagedintegrationsIdentifier)
	}

	if resp, err := client.DeleteCloudConnector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a connector destination linked to a cloud-to-cloud (C2C) connector.
// Deletion can't be done if the account association has used this connector
// destination.
func iotmanagedintegrations_DeleteConnectorDestination(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.DeleteConnectorDestinationInput{
		// Identifier: *string, // Required
	}

	if len(_iotmanagedintegrationsIdentifier) > 0 {
		input.Identifier = aws.String(_iotmanagedintegrationsIdentifier)
	}

	if resp, err := client.DeleteConnectorDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a credential locker.
// This operation can't be undone and any existing device won't be able to use IoT
// managed integrations.
func iotmanagedintegrations_DeleteCredentialLocker(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.DeleteCredentialLockerInput{
		// Identifier: *string, // Required
	}

	if len(_iotmanagedintegrationsIdentifier) > 0 {
		input.Identifier = aws.String(_iotmanagedintegrationsIdentifier)
	}

	if resp, err := client.DeleteCredentialLocker(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a notification destination specified by name.
func iotmanagedintegrations_DeleteDestination(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.DeleteDestinationInput{
		// Name: *string, // Required
	}

	if len(_iotmanagedintegrationsName) > 0 {
		input.Name = aws.String(_iotmanagedintegrationsName)
	}

	if resp, err := client.DeleteDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete an event log configuration.
func iotmanagedintegrations_DeleteEventLogConfiguration(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.DeleteEventLogConfigurationInput{
		// Id: *string, // Required
	}

	if len(_iotmanagedintegrationsId) > 0 {
		input.Id = aws.String(_iotmanagedintegrationsId)
	}

	if resp, err := client.DeleteEventLogConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a managed thing. For direct-connected and hub-connected devices
// connecting with Managed integrations via a controller, all of the devices
// connected to it will have their status changed to PENDING . It is not possible
// to remove a cloud-to-cloud device.
func iotmanagedintegrations_DeleteManagedThing(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.DeleteManagedThingInput{
		// Identifier: *string, // Required
	}

	if len(_iotmanagedintegrationsIdentifier) > 0 {
		input.Identifier = aws.String(_iotmanagedintegrationsIdentifier)
	}
	if len(_iotmanagedintegrationsForce) > 0 {
		if err := assignInputField(input, "Force", _iotmanagedintegrationsForce); err != nil {
			log.Errorf("invalid --force: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteManagedThing(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a notification configuration.
func iotmanagedintegrations_DeleteNotificationConfiguration(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.DeleteNotificationConfigurationInput{
		// EventType: types.EventType, // Required
	}

	if len(_iotmanagedintegrationsEventType) > 0 {
		if err := assignInputField(input, "EventType", _iotmanagedintegrationsEventType); err != nil {
			log.Errorf("invalid --event-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteNotificationConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete the over-the-air (OTA) task.
func iotmanagedintegrations_DeleteOtaTask(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.DeleteOtaTaskInput{
		// Identifier: *string, // Required
	}

	if len(_iotmanagedintegrationsIdentifier) > 0 {
		input.Identifier = aws.String(_iotmanagedintegrationsIdentifier)
	}

	if resp, err := client.DeleteOtaTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete the over-the-air (OTA) task configuration.
func iotmanagedintegrations_DeleteOtaTaskConfiguration(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.DeleteOtaTaskConfigurationInput{
		// Identifier: *string, // Required
	}

	if len(_iotmanagedintegrationsIdentifier) > 0 {
		input.Identifier = aws.String(_iotmanagedintegrationsIdentifier)
	}

	if resp, err := client.DeleteOtaTaskConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a provisioning profile.
func iotmanagedintegrations_DeleteProvisioningProfile(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.DeleteProvisioningProfileInput{
		// Identifier: *string, // Required
	}

	if len(_iotmanagedintegrationsIdentifier) > 0 {
		input.Identifier = aws.String(_iotmanagedintegrationsIdentifier)
	}

	if resp, err := client.DeleteProvisioningProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deregister an account association from a managed thing.
func iotmanagedintegrations_DeregisterAccountAssociation(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.DeregisterAccountAssociationInput{
		// AccountAssociationId: *string, // Required
		// ManagedThingId: *string, // Required
	}

	if len(_iotmanagedintegrationsAccountAssociationId) > 0 {
		input.AccountAssociationId = aws.String(_iotmanagedintegrationsAccountAssociationId)
	}
	if len(_iotmanagedintegrationsManagedThingId) > 0 {
		input.ManagedThingId = aws.String(_iotmanagedintegrationsManagedThingId)
	}

	if resp, err := client.DeregisterAccountAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get an account association for an Amazon Web Services account linked to a
// customer-managed destination.
func iotmanagedintegrations_GetAccountAssociation(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.GetAccountAssociationInput{
		// AccountAssociationId: *string, // Required
	}

	if len(_iotmanagedintegrationsAccountAssociationId) > 0 {
		input.AccountAssociationId = aws.String(_iotmanagedintegrationsAccountAssociationId)
	}

	if resp, err := client.GetAccountAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get configuration details for a cloud connector.
func iotmanagedintegrations_GetCloudConnector(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.GetCloudConnectorInput{
		// Identifier: *string, // Required
	}

	if len(_iotmanagedintegrationsIdentifier) > 0 {
		input.Identifier = aws.String(_iotmanagedintegrationsIdentifier)
	}

	if resp, err := client.GetCloudConnector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get connector destination details linked to a cloud-to-cloud (C2C) connector.
func iotmanagedintegrations_GetConnectorDestination(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.GetConnectorDestinationInput{
		// Identifier: *string, // Required
	}

	if len(_iotmanagedintegrationsIdentifier) > 0 {
		input.Identifier = aws.String(_iotmanagedintegrationsIdentifier)
	}

	if resp, err := client.GetConnectorDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get information on an existing credential locker
func iotmanagedintegrations_GetCredentialLocker(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.GetCredentialLockerInput{
		// Identifier: *string, // Required
	}

	if len(_iotmanagedintegrationsIdentifier) > 0 {
		input.Identifier = aws.String(_iotmanagedintegrationsIdentifier)
	}

	if resp, err := client.GetCredentialLocker(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the IoT managed integrations custom endpoint.
func iotmanagedintegrations_GetCustomEndpoint(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.GetCustomEndpointInput{}

	if resp, err := client.GetCustomEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the default encryption configuration for the
// Amazon Web Services account in the default or specified region. For more
// information, see [Key management]in the AWS IoT SiteWise User Guide.
//
// [Key management]: https://docs.aws.amazon.com/iot-sitewise/latest/userguide/key-management.html
func iotmanagedintegrations_GetDefaultEncryptionConfiguration(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.GetDefaultEncryptionConfigurationInput{}

	if resp, err := client.GetDefaultEncryptionConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a destination by name.
func iotmanagedintegrations_GetDestination(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.GetDestinationInput{
		// Name: *string, // Required
	}

	if len(_iotmanagedintegrationsName) > 0 {
		input.Name = aws.String(_iotmanagedintegrationsName)
	}

	if resp, err := client.GetDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the current state of a device discovery.
func iotmanagedintegrations_GetDeviceDiscovery(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.GetDeviceDiscoveryInput{
		// Identifier: *string, // Required
	}

	if len(_iotmanagedintegrationsIdentifier) > 0 {
		input.Identifier = aws.String(_iotmanagedintegrationsIdentifier)
	}

	if resp, err := client.GetDeviceDiscovery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get an event log configuration.
func iotmanagedintegrations_GetEventLogConfiguration(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.GetEventLogConfigurationInput{
		// Id: *string, // Required
	}

	if len(_iotmanagedintegrationsId) > 0 {
		input.Id = aws.String(_iotmanagedintegrationsId)
	}

	if resp, err := client.GetEventLogConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get a hub configuration.
func iotmanagedintegrations_GetHubConfiguration(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.GetHubConfigurationInput{}

	if resp, err := client.GetHubConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get details of a managed thing including its attributes and capabilities.
func iotmanagedintegrations_GetManagedThing(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.GetManagedThingInput{
		// Identifier: *string, // Required
	}

	if len(_iotmanagedintegrationsIdentifier) > 0 {
		input.Identifier = aws.String(_iotmanagedintegrationsIdentifier)
	}

	if resp, err := client.GetManagedThing(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the capabilities for a managed thing using the device ID.
func iotmanagedintegrations_GetManagedThingCapabilities(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.GetManagedThingCapabilitiesInput{
		// Identifier: *string, // Required
	}

	if len(_iotmanagedintegrationsIdentifier) > 0 {
		input.Identifier = aws.String(_iotmanagedintegrationsIdentifier)
	}

	if resp, err := client.GetManagedThingCapabilities(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the certificate PEM for a managed IoT thing.
func iotmanagedintegrations_GetManagedThingCertificate(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.GetManagedThingCertificateInput{
		// Identifier: *string, // Required
	}

	if len(_iotmanagedintegrationsIdentifier) > 0 {
		input.Identifier = aws.String(_iotmanagedintegrationsIdentifier)
	}

	if resp, err := client.GetManagedThingCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the connectivity status of a managed thing.
func iotmanagedintegrations_GetManagedThingConnectivityData(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.GetManagedThingConnectivityDataInput{
		// Identifier: *string, // Required
	}

	if len(_iotmanagedintegrationsIdentifier) > 0 {
		input.Identifier = aws.String(_iotmanagedintegrationsIdentifier)
	}

	if resp, err := client.GetManagedThingConnectivityData(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the metadata information for a managed thing.
// The managedThing metadata parameter is used for associating attributes with a
// managedThing that can be used for grouping over-the-air (OTA) tasks. Name value
// pairs in metadata can be used in the OtaTargetQueryString parameter for the
// CreateOtaTask API operation.
func iotmanagedintegrations_GetManagedThingMetaData(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.GetManagedThingMetaDataInput{
		// Identifier: *string, // Required
	}

	if len(_iotmanagedintegrationsIdentifier) > 0 {
		input.Identifier = aws.String(_iotmanagedintegrationsIdentifier)
	}

	if resp, err := client.GetManagedThingMetaData(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the managed thing state for the given device Id.
func iotmanagedintegrations_GetManagedThingState(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.GetManagedThingStateInput{
		// ManagedThingId: *string, // Required
	}

	if len(_iotmanagedintegrationsManagedThingId) > 0 {
		input.ManagedThingId = aws.String(_iotmanagedintegrationsManagedThingId)
	}

	if resp, err := client.GetManagedThingState(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get a notification configuration for a specified event type.
func iotmanagedintegrations_GetNotificationConfiguration(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.GetNotificationConfigurationInput{
		// EventType: types.EventType, // Required
	}

	if len(_iotmanagedintegrationsEventType) > 0 {
		if err := assignInputField(input, "EventType", _iotmanagedintegrationsEventType); err != nil {
			log.Errorf("invalid --event-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetNotificationConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get details of the over-the-air (OTA) task by its task id.
func iotmanagedintegrations_GetOtaTask(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.GetOtaTaskInput{
		// Identifier: *string, // Required
	}

	if len(_iotmanagedintegrationsIdentifier) > 0 {
		input.Identifier = aws.String(_iotmanagedintegrationsIdentifier)
	}

	if resp, err := client.GetOtaTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get a configuraiton for the over-the-air (OTA) task.
func iotmanagedintegrations_GetOtaTaskConfiguration(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.GetOtaTaskConfigurationInput{
		// Identifier: *string, // Required
	}

	if len(_iotmanagedintegrationsIdentifier) > 0 {
		input.Identifier = aws.String(_iotmanagedintegrationsIdentifier)
	}

	if resp, err := client.GetOtaTaskConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get a provisioning profile by template name.
func iotmanagedintegrations_GetProvisioningProfile(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.GetProvisioningProfileInput{
		// Identifier: *string, // Required
	}

	if len(_iotmanagedintegrationsIdentifier) > 0 {
		input.Identifier = aws.String(_iotmanagedintegrationsIdentifier)
	}

	if resp, err := client.GetProvisioningProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the runtime log configuration for a specific managed thing.
func iotmanagedintegrations_GetRuntimeLogConfiguration(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.GetRuntimeLogConfigurationInput{
		// ManagedThingId: *string, // Required
	}

	if len(_iotmanagedintegrationsManagedThingId) > 0 {
		input.ManagedThingId = aws.String(_iotmanagedintegrationsManagedThingId)
	}

	if resp, err := client.GetRuntimeLogConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a schema version with the provided information.
func iotmanagedintegrations_GetSchemaVersion(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.GetSchemaVersionInput{
		// SchemaVersionedId: *string, // Required
		// Type: types.SchemaVersionType, // Required
	}

	if len(_iotmanagedintegrationsSchemaVersionedId) > 0 {
		input.SchemaVersionedId = aws.String(_iotmanagedintegrationsSchemaVersionedId)
	}
	if len(_iotmanagedintegrationsType) > 0 {
		if err := assignInputField(input, "Type", _iotmanagedintegrationsType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsFormat) > 0 {
		if err := assignInputField(input, "Format", _iotmanagedintegrationsFormat); err != nil {
			log.Errorf("invalid --format: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetSchemaVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all account associations, with optional filtering by connector
// destination ID.
func iotmanagedintegrations_ListAccountAssociations(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.ListAccountAssociationsInput{}

	if len(_iotmanagedintegrationsConnectorDestinationId) > 0 {
		input.ConnectorDestinationId = aws.String(_iotmanagedintegrationsConnectorDestinationId)
	}
	if len(_iotmanagedintegrationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotmanagedintegrationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsNextToken) > 0 {
		input.NextToken = aws.String(_iotmanagedintegrationsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAccountAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotmanagedintegrations.ListAccountAssociationsOutput
	p := iotmanagedintegrations.NewListAccountAssociationsPaginator(client, input)
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

// Returns a list of connectors filtered by its Lambda Amazon Resource Name (ARN)
// and type .
func iotmanagedintegrations_ListCloudConnectors(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.ListCloudConnectorsInput{}

	if len(_iotmanagedintegrationsLambdaArn) > 0 {
		input.LambdaArn = aws.String(_iotmanagedintegrationsLambdaArn)
	}
	if len(_iotmanagedintegrationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotmanagedintegrationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsNextToken) > 0 {
		input.NextToken = aws.String(_iotmanagedintegrationsNextToken)
	}
	if len(_iotmanagedintegrationsType) > 0 {
		if err := assignInputField(input, "Type", _iotmanagedintegrationsType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListCloudConnectors(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotmanagedintegrations.ListCloudConnectorsOutput
	p := iotmanagedintegrations.NewListCloudConnectorsPaginator(client, input)
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

// Lists all connector destinations, with optional filtering by cloud connector ID.
func iotmanagedintegrations_ListConnectorDestinations(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.ListConnectorDestinationsInput{}

	if len(_iotmanagedintegrationsCloudConnectorId) > 0 {
		input.CloudConnectorId = aws.String(_iotmanagedintegrationsCloudConnectorId)
	}
	if len(_iotmanagedintegrationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotmanagedintegrationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsNextToken) > 0 {
		input.NextToken = aws.String(_iotmanagedintegrationsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListConnectorDestinations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotmanagedintegrations.ListConnectorDestinationsOutput
	p := iotmanagedintegrations.NewListConnectorDestinationsPaginator(client, input)
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

// List information on an existing credential locker.
func iotmanagedintegrations_ListCredentialLockers(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.ListCredentialLockersInput{}

	if len(_iotmanagedintegrationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotmanagedintegrationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsNextToken) > 0 {
		input.NextToken = aws.String(_iotmanagedintegrationsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCredentialLockers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotmanagedintegrations.ListCredentialLockersOutput
	p := iotmanagedintegrations.NewListCredentialLockersPaginator(client, input)
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

// List all notification destinations.
func iotmanagedintegrations_ListDestinations(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.ListDestinationsInput{}

	if len(_iotmanagedintegrationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotmanagedintegrationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsNextToken) > 0 {
		input.NextToken = aws.String(_iotmanagedintegrationsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDestinations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotmanagedintegrations.ListDestinationsOutput
	p := iotmanagedintegrations.NewListDestinationsPaginator(client, input)
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

// Lists all device discovery tasks, with optional filtering by type and status.
func iotmanagedintegrations_ListDeviceDiscoveries(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.ListDeviceDiscoveriesInput{}

	if len(_iotmanagedintegrationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotmanagedintegrationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsNextToken) > 0 {
		input.NextToken = aws.String(_iotmanagedintegrationsNextToken)
	}
	if len(_iotmanagedintegrationsStatusFilter) > 0 {
		if err := assignInputField(input, "StatusFilter", _iotmanagedintegrationsStatusFilter); err != nil {
			log.Errorf("invalid --status-filter: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsTypeFilter) > 0 {
		if err := assignInputField(input, "TypeFilter", _iotmanagedintegrationsTypeFilter); err != nil {
			log.Errorf("invalid --type-filter: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListDeviceDiscoveries(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotmanagedintegrations.ListDeviceDiscoveriesOutput
	p := iotmanagedintegrations.NewListDeviceDiscoveriesPaginator(client, input)
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

// Lists all devices discovered during a specific device discovery task.
func iotmanagedintegrations_ListDiscoveredDevices(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.ListDiscoveredDevicesInput{
		// Identifier: *string, // Required
	}

	if len(_iotmanagedintegrationsIdentifier) > 0 {
		input.Identifier = aws.String(_iotmanagedintegrationsIdentifier)
	}
	if len(_iotmanagedintegrationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotmanagedintegrationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsNextToken) > 0 {
		input.NextToken = aws.String(_iotmanagedintegrationsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDiscoveredDevices(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotmanagedintegrations.ListDiscoveredDevicesOutput
	p := iotmanagedintegrations.NewListDiscoveredDevicesPaginator(client, input)
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

// List all event log configurations for an account.
func iotmanagedintegrations_ListEventLogConfigurations(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.ListEventLogConfigurationsInput{}

	if len(_iotmanagedintegrationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotmanagedintegrationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsNextToken) > 0 {
		input.NextToken = aws.String(_iotmanagedintegrationsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEventLogConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotmanagedintegrations.ListEventLogConfigurationsOutput
	p := iotmanagedintegrations.NewListEventLogConfigurationsPaginator(client, input)
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

// Lists all account associations for a specific managed thing.
func iotmanagedintegrations_ListManagedThingAccountAssociations(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.ListManagedThingAccountAssociationsInput{}

	if len(_iotmanagedintegrationsAccountAssociationId) > 0 {
		input.AccountAssociationId = aws.String(_iotmanagedintegrationsAccountAssociationId)
	}
	if len(_iotmanagedintegrationsManagedThingId) > 0 {
		input.ManagedThingId = aws.String(_iotmanagedintegrationsManagedThingId)
	}
	if len(_iotmanagedintegrationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotmanagedintegrationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsNextToken) > 0 {
		input.NextToken = aws.String(_iotmanagedintegrationsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListManagedThingAccountAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotmanagedintegrations.ListManagedThingAccountAssociationsOutput
	p := iotmanagedintegrations.NewListManagedThingAccountAssociationsPaginator(client, input)
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

// List schemas associated with a managed thing.
func iotmanagedintegrations_ListManagedThingSchemas(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.ListManagedThingSchemasInput{
		// Identifier: *string, // Required
	}

	if len(_iotmanagedintegrationsIdentifier) > 0 {
		input.Identifier = aws.String(_iotmanagedintegrationsIdentifier)
	}
	if len(_iotmanagedintegrationsCapabilityIdFilter) > 0 {
		input.CapabilityIdFilter = aws.String(_iotmanagedintegrationsCapabilityIdFilter)
	}
	if len(_iotmanagedintegrationsEndpointIdFilter) > 0 {
		input.EndpointIdFilter = aws.String(_iotmanagedintegrationsEndpointIdFilter)
	}
	if len(_iotmanagedintegrationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotmanagedintegrationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsNextToken) > 0 {
		input.NextToken = aws.String(_iotmanagedintegrationsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListManagedThingSchemas(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotmanagedintegrations.ListManagedThingSchemasOutput
	p := iotmanagedintegrations.NewListManagedThingSchemasPaginator(client, input)
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

// Listing all managed things with provision for filters.
func iotmanagedintegrations_ListManagedThings(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.ListManagedThingsInput{}

	if len(_iotmanagedintegrationsConnectorDestinationIdFilter) > 0 {
		input.ConnectorDestinationIdFilter = aws.String(_iotmanagedintegrationsConnectorDestinationIdFilter)
	}
	if len(_iotmanagedintegrationsConnectorDeviceIdFilter) > 0 {
		input.ConnectorDeviceIdFilter = aws.String(_iotmanagedintegrationsConnectorDeviceIdFilter)
	}
	if len(_iotmanagedintegrationsConnectorPolicyIdFilter) > 0 {
		input.ConnectorPolicyIdFilter = aws.String(_iotmanagedintegrationsConnectorPolicyIdFilter)
	}
	if len(_iotmanagedintegrationsCredentialLockerFilter) > 0 {
		input.CredentialLockerFilter = aws.String(_iotmanagedintegrationsCredentialLockerFilter)
	}
	if len(_iotmanagedintegrationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotmanagedintegrationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsNextToken) > 0 {
		input.NextToken = aws.String(_iotmanagedintegrationsNextToken)
	}
	if len(_iotmanagedintegrationsOwnerFilter) > 0 {
		input.OwnerFilter = aws.String(_iotmanagedintegrationsOwnerFilter)
	}
	if len(_iotmanagedintegrationsParentControllerIdentifierFilter) > 0 {
		input.ParentControllerIdentifierFilter = aws.String(_iotmanagedintegrationsParentControllerIdentifierFilter)
	}
	if len(_iotmanagedintegrationsProvisioningStatusFilter) > 0 {
		if err := assignInputField(input, "ProvisioningStatusFilter", _iotmanagedintegrationsProvisioningStatusFilter); err != nil {
			log.Errorf("invalid --provisioning-status-filter: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsRoleFilter) > 0 {
		if err := assignInputField(input, "RoleFilter", _iotmanagedintegrationsRoleFilter); err != nil {
			log.Errorf("invalid --role-filter: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsSerialNumberFilter) > 0 {
		input.SerialNumberFilter = aws.String(_iotmanagedintegrationsSerialNumberFilter)
	}

	if disablePaginator() {
		if resp, err := client.ListManagedThings(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotmanagedintegrations.ListManagedThingsOutput
	p := iotmanagedintegrations.NewListManagedThingsPaginator(client, input)
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

// List all notification configurations.
func iotmanagedintegrations_ListNotificationConfigurations(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.ListNotificationConfigurationsInput{}

	if len(_iotmanagedintegrationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotmanagedintegrationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsNextToken) > 0 {
		input.NextToken = aws.String(_iotmanagedintegrationsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListNotificationConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotmanagedintegrations.ListNotificationConfigurationsOutput
	p := iotmanagedintegrations.NewListNotificationConfigurationsPaginator(client, input)
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

// List all of the over-the-air (OTA) task configurations.
func iotmanagedintegrations_ListOtaTaskConfigurations(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.ListOtaTaskConfigurationsInput{}

	if len(_iotmanagedintegrationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotmanagedintegrationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsNextToken) > 0 {
		input.NextToken = aws.String(_iotmanagedintegrationsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListOtaTaskConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotmanagedintegrations.ListOtaTaskConfigurationsOutput
	p := iotmanagedintegrations.NewListOtaTaskConfigurationsPaginator(client, input)
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

// List all of the over-the-air (OTA) task executions.
func iotmanagedintegrations_ListOtaTaskExecutions(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.ListOtaTaskExecutionsInput{
		// Identifier: *string, // Required
	}

	if len(_iotmanagedintegrationsIdentifier) > 0 {
		input.Identifier = aws.String(_iotmanagedintegrationsIdentifier)
	}
	if len(_iotmanagedintegrationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotmanagedintegrationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsNextToken) > 0 {
		input.NextToken = aws.String(_iotmanagedintegrationsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListOtaTaskExecutions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotmanagedintegrations.ListOtaTaskExecutionsOutput
	p := iotmanagedintegrations.NewListOtaTaskExecutionsPaginator(client, input)
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

// List all of the over-the-air (OTA) tasks.
func iotmanagedintegrations_ListOtaTasks(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.ListOtaTasksInput{}

	if len(_iotmanagedintegrationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotmanagedintegrationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsNextToken) > 0 {
		input.NextToken = aws.String(_iotmanagedintegrationsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListOtaTasks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotmanagedintegrations.ListOtaTasksOutput
	p := iotmanagedintegrations.NewListOtaTasksPaginator(client, input)
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

// List the provisioning profiles within the Amazon Web Services account.
func iotmanagedintegrations_ListProvisioningProfiles(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.ListProvisioningProfilesInput{}

	if len(_iotmanagedintegrationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotmanagedintegrationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsNextToken) > 0 {
		input.NextToken = aws.String(_iotmanagedintegrationsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListProvisioningProfiles(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotmanagedintegrations.ListProvisioningProfilesOutput
	p := iotmanagedintegrations.NewListProvisioningProfilesPaginator(client, input)
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

// Lists schema versions with the provided information.
func iotmanagedintegrations_ListSchemaVersions(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.ListSchemaVersionsInput{
		// Type: types.SchemaVersionType, // Required
	}

	if len(_iotmanagedintegrationsType) > 0 {
		if err := assignInputField(input, "Type", _iotmanagedintegrationsType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotmanagedintegrationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsNamespace) > 0 {
		input.Namespace = aws.String(_iotmanagedintegrationsNamespace)
	}
	if len(_iotmanagedintegrationsNextToken) > 0 {
		input.NextToken = aws.String(_iotmanagedintegrationsNextToken)
	}
	if len(_iotmanagedintegrationsSchemaId) > 0 {
		input.SchemaId = aws.String(_iotmanagedintegrationsSchemaId)
	}
	if len(_iotmanagedintegrationsSemanticVersion) > 0 {
		input.SemanticVersion = aws.String(_iotmanagedintegrationsSemanticVersion)
	}
	if len(_iotmanagedintegrationsVisibility) > 0 {
		if err := assignInputField(input, "Visibility", _iotmanagedintegrationsVisibility); err != nil {
			log.Errorf("invalid --visibility: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListSchemaVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotmanagedintegrations.ListSchemaVersionsOutput
	p := iotmanagedintegrations.NewListSchemaVersionsPaginator(client, input)
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

// Lists the tags for a specified resource.
func iotmanagedintegrations_ListTagsForResource(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_iotmanagedintegrationsResourceArn) > 0 {
		input.ResourceArn = aws.String(_iotmanagedintegrationsResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the default encryption configuration for the Amazon Web Services account.
// For more information, see [Key management]in the AWS IoT SiteWise User Guide.
//
// [Key management]: https://docs.aws.amazon.com/iot-sitewise/latest/userguide/key-management.html
func iotmanagedintegrations_PutDefaultEncryptionConfiguration(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.PutDefaultEncryptionConfigurationInput{
		// EncryptionType: types.EncryptionType, // Required
	}

	if len(_iotmanagedintegrationsEncryptionType) > 0 {
		if err := assignInputField(input, "EncryptionType", _iotmanagedintegrationsEncryptionType); err != nil {
			log.Errorf("invalid --encryption-type: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsKmsKeyArn) > 0 {
		input.KmsKeyArn = aws.String(_iotmanagedintegrationsKmsKeyArn)
	}

	if resp, err := client.PutDefaultEncryptionConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update a hub configuration.
func iotmanagedintegrations_PutHubConfiguration(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.PutHubConfigurationInput{
		// HubTokenTimerExpirySettingInSeconds: *int64, // Required
	}

	if len(_iotmanagedintegrationsHubTokenTimerExpirySettingInSeconds) > 0 {
		if err := assignInputField(input, "HubTokenTimerExpirySettingInSeconds", _iotmanagedintegrationsHubTokenTimerExpirySettingInSeconds); err != nil {
			log.Errorf("invalid --hub-token-timer-expiry-setting-in-seconds: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutHubConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Set the runtime log configuration for a specific managed thing.
func iotmanagedintegrations_PutRuntimeLogConfiguration(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.PutRuntimeLogConfigurationInput{
		// ManagedThingId: *string, // Required
		// RuntimeLogConfigurations: *types.RuntimeLogConfigurations, // Required
	}

	if len(_iotmanagedintegrationsManagedThingId) > 0 {
		input.ManagedThingId = aws.String(_iotmanagedintegrationsManagedThingId)
	}
	if len(_iotmanagedintegrationsRuntimeLogConfigurations) > 0 {
		if err := assignInputField(input, "RuntimeLogConfigurations", _iotmanagedintegrationsRuntimeLogConfigurations); err != nil {
			log.Errorf("invalid --runtime-log-configurations: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutRuntimeLogConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Registers an account association with a managed thing, establishing a
// connection between a device and a third-party account.
func iotmanagedintegrations_RegisterAccountAssociation(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.RegisterAccountAssociationInput{
		// AccountAssociationId: *string, // Required
		// DeviceDiscoveryId: *string, // Required
		// ManagedThingId: *string, // Required
	}

	if len(_iotmanagedintegrationsAccountAssociationId) > 0 {
		input.AccountAssociationId = aws.String(_iotmanagedintegrationsAccountAssociationId)
	}
	if len(_iotmanagedintegrationsDeviceDiscoveryId) > 0 {
		input.DeviceDiscoveryId = aws.String(_iotmanagedintegrationsDeviceDiscoveryId)
	}
	if len(_iotmanagedintegrationsManagedThingId) > 0 {
		input.ManagedThingId = aws.String(_iotmanagedintegrationsManagedThingId)
	}

	if resp, err := client.RegisterAccountAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Customers can request IoT managed integrations to manage the server trust for
// them or bring their own external server trusts for the custom domain. Returns an
// IoT managed integrations endpoint.
func iotmanagedintegrations_RegisterCustomEndpoint(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.RegisterCustomEndpointInput{}

	if resp, err := client.RegisterCustomEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Reset a runtime log configuration for a specific managed thing.
func iotmanagedintegrations_ResetRuntimeLogConfiguration(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.ResetRuntimeLogConfigurationInput{
		// ManagedThingId: *string, // Required
	}

	if len(_iotmanagedintegrationsManagedThingId) > 0 {
		input.ManagedThingId = aws.String(_iotmanagedintegrationsManagedThingId)
	}

	if resp, err := client.ResetRuntimeLogConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Relays third-party device events for a connector such as a new device or a
// device state change event.
func iotmanagedintegrations_SendConnectorEvent(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.SendConnectorEventInput{
		// ConnectorId: *string, // Required
		// Operation: types.ConnectorEventOperation, // Required
	}

	if len(_iotmanagedintegrationsConnectorId) > 0 {
		input.ConnectorId = aws.String(_iotmanagedintegrationsConnectorId)
	}
	if len(_iotmanagedintegrationsOperation) > 0 {
		if err := assignInputField(input, "Operation", _iotmanagedintegrationsOperation); err != nil {
			log.Errorf("invalid --operation: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsConnectorDeviceId) > 0 {
		input.ConnectorDeviceId = aws.String(_iotmanagedintegrationsConnectorDeviceId)
	}
	if len(_iotmanagedintegrationsDeviceDiscoveryId) > 0 {
		input.DeviceDiscoveryId = aws.String(_iotmanagedintegrationsDeviceDiscoveryId)
	}
	if len(_iotmanagedintegrationsDevices) > 0 {
		if err := assignInputField(input, "Devices", _iotmanagedintegrationsDevices); err != nil {
			log.Errorf("invalid --devices: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsMatterEndpoint) > 0 {
		if err := assignInputField(input, "MatterEndpoint", _iotmanagedintegrationsMatterEndpoint); err != nil {
			log.Errorf("invalid --matter-endpoint: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsMessage) > 0 {
		input.Message = aws.String(_iotmanagedintegrationsMessage)
	}
	if len(_iotmanagedintegrationsOperationVersion) > 0 {
		input.OperationVersion = aws.String(_iotmanagedintegrationsOperationVersion)
	}
	if len(_iotmanagedintegrationsStatusCode) > 0 {
		if err := assignInputField(input, "StatusCode", _iotmanagedintegrationsStatusCode); err != nil {
			log.Errorf("invalid --status-code: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsTraceId) > 0 {
		input.TraceId = aws.String(_iotmanagedintegrationsTraceId)
	}
	if len(_iotmanagedintegrationsUserId) > 0 {
		input.UserId = aws.String(_iotmanagedintegrationsUserId)
	}

	if resp, err := client.SendConnectorEvent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Send the command to the device represented by the managed thing.
func iotmanagedintegrations_SendManagedThingCommand(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.SendManagedThingCommandInput{
		// Endpoints: []types.CommandEndpoint, // Required
		// ManagedThingId: *string, // Required
	}

	if len(_iotmanagedintegrationsEndpoints) > 0 {
		if err := assignInputField(input, "Endpoints", _iotmanagedintegrationsEndpoints); err != nil {
			log.Errorf("invalid --endpoints: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsManagedThingId) > 0 {
		input.ManagedThingId = aws.String(_iotmanagedintegrationsManagedThingId)
	}
	if len(_iotmanagedintegrationsAccountAssociationId) > 0 {
		input.AccountAssociationId = aws.String(_iotmanagedintegrationsAccountAssociationId)
	}
	if len(_iotmanagedintegrationsConnectorAssociationId) > 0 {
		input.ConnectorAssociationId = aws.String(_iotmanagedintegrationsConnectorAssociationId)
	}

	if resp, err := client.SendManagedThingCommand(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initiates a refresh of an existing account association to update its
// authorization and connection status.
func iotmanagedintegrations_StartAccountAssociationRefresh(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.StartAccountAssociationRefreshInput{
		// AccountAssociationId: *string, // Required
	}

	if len(_iotmanagedintegrationsAccountAssociationId) > 0 {
		input.AccountAssociationId = aws.String(_iotmanagedintegrationsAccountAssociationId)
	}

	if resp, err := client.StartAccountAssociationRefresh(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API is used to start device discovery for hub-connected and
// third-party-connected devices. The authentication material (install code) is
// delivered as a message to the controller instructing it to start the discovery.
func iotmanagedintegrations_StartDeviceDiscovery(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.StartDeviceDiscoveryInput{
		// DiscoveryType: types.DiscoveryType, // Required
	}

	if len(_iotmanagedintegrationsDiscoveryType) > 0 {
		if err := assignInputField(input, "DiscoveryType", _iotmanagedintegrationsDiscoveryType); err != nil {
			log.Errorf("invalid --discovery-type: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsAccountAssociationId) > 0 {
		input.AccountAssociationId = aws.String(_iotmanagedintegrationsAccountAssociationId)
	}
	if len(_iotmanagedintegrationsAuthenticationMaterial) > 0 {
		input.AuthenticationMaterial = aws.String(_iotmanagedintegrationsAuthenticationMaterial)
	}
	if len(_iotmanagedintegrationsAuthenticationMaterialType) > 0 {
		if err := assignInputField(input, "AuthenticationMaterialType", _iotmanagedintegrationsAuthenticationMaterialType); err != nil {
			log.Errorf("invalid --authentication-material-type: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsClientToken) > 0 {
		input.ClientToken = aws.String(_iotmanagedintegrationsClientToken)
	}
	if len(_iotmanagedintegrationsConnectorAssociationIdentifier) > 0 {
		input.ConnectorAssociationIdentifier = aws.String(_iotmanagedintegrationsConnectorAssociationIdentifier)
	}
	if len(_iotmanagedintegrationsConnectorDeviceIdList) > 0 {
		input.ConnectorDeviceIdList = append([]string(nil), _iotmanagedintegrationsConnectorDeviceIdList...)
	}
	if len(_iotmanagedintegrationsControllerIdentifier) > 0 {
		input.ControllerIdentifier = aws.String(_iotmanagedintegrationsControllerIdentifier)
	}
	if len(_iotmanagedintegrationsCustomProtocolDetail) > 0 {
		if err := assignInputField(input, "CustomProtocolDetail", _iotmanagedintegrationsCustomProtocolDetail); err != nil {
			log.Errorf("invalid --custom-protocol-detail: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsEndDeviceIdentifier) > 0 {
		input.EndDeviceIdentifier = aws.String(_iotmanagedintegrationsEndDeviceIdentifier)
	}
	if len(_iotmanagedintegrationsProtocol) > 0 {
		if err := assignInputField(input, "Protocol", _iotmanagedintegrationsProtocol); err != nil {
			log.Errorf("invalid --protocol: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsTags) > 0 {
		if err := assignInputField(input, "Tags", _iotmanagedintegrationsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartDeviceDiscovery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds tags to a specified resource.
func iotmanagedintegrations_TagResource(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_iotmanagedintegrationsResourceArn) > 0 {
		input.ResourceArn = aws.String(_iotmanagedintegrationsResourceArn)
	}
	if len(_iotmanagedintegrationsTags) > 0 {
		if err := assignInputField(input, "Tags", _iotmanagedintegrationsTags); err != nil {
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

// Removes tags from a specified resource.
func iotmanagedintegrations_UntagResource(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_iotmanagedintegrationsResourceArn) > 0 {
		input.ResourceArn = aws.String(_iotmanagedintegrationsResourceArn)
	}
	if len(_iotmanagedintegrationsTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _iotmanagedintegrationsTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the properties of an existing account association.
func iotmanagedintegrations_UpdateAccountAssociation(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.UpdateAccountAssociationInput{
		// AccountAssociationId: *string, // Required
	}

	if len(_iotmanagedintegrationsAccountAssociationId) > 0 {
		input.AccountAssociationId = aws.String(_iotmanagedintegrationsAccountAssociationId)
	}
	if len(_iotmanagedintegrationsDescription) > 0 {
		input.Description = aws.String(_iotmanagedintegrationsDescription)
	}
	if len(_iotmanagedintegrationsName) > 0 {
		input.Name = aws.String(_iotmanagedintegrationsName)
	}

	if resp, err := client.UpdateAccountAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update an existing cloud connector.
func iotmanagedintegrations_UpdateCloudConnector(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.UpdateCloudConnectorInput{
		// Identifier: *string, // Required
	}

	if len(_iotmanagedintegrationsIdentifier) > 0 {
		input.Identifier = aws.String(_iotmanagedintegrationsIdentifier)
	}
	if len(_iotmanagedintegrationsDescription) > 0 {
		input.Description = aws.String(_iotmanagedintegrationsDescription)
	}
	if len(_iotmanagedintegrationsName) > 0 {
		input.Name = aws.String(_iotmanagedintegrationsName)
	}

	if resp, err := client.UpdateCloudConnector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the properties of an existing connector destination.
func iotmanagedintegrations_UpdateConnectorDestination(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.UpdateConnectorDestinationInput{
		// Identifier: *string, // Required
	}

	if len(_iotmanagedintegrationsIdentifier) > 0 {
		input.Identifier = aws.String(_iotmanagedintegrationsIdentifier)
	}
	if len(_iotmanagedintegrationsAuthConfig) > 0 {
		if err := assignInputField(input, "AuthConfig", _iotmanagedintegrationsAuthConfig); err != nil {
			log.Errorf("invalid --auth-config: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsAuthType) > 0 {
		if err := assignInputField(input, "AuthType", _iotmanagedintegrationsAuthType); err != nil {
			log.Errorf("invalid --auth-type: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsDescription) > 0 {
		input.Description = aws.String(_iotmanagedintegrationsDescription)
	}
	if len(_iotmanagedintegrationsName) > 0 {
		input.Name = aws.String(_iotmanagedintegrationsName)
	}
	if len(_iotmanagedintegrationsSecretsManager) > 0 {
		if err := assignInputField(input, "SecretsManager", _iotmanagedintegrationsSecretsManager); err != nil {
			log.Errorf("invalid --secrets-manager: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateConnectorDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update a destination specified by name.
func iotmanagedintegrations_UpdateDestination(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.UpdateDestinationInput{
		// Name: *string, // Required
	}

	if len(_iotmanagedintegrationsName) > 0 {
		input.Name = aws.String(_iotmanagedintegrationsName)
	}
	if len(_iotmanagedintegrationsDeliveryDestinationArn) > 0 {
		input.DeliveryDestinationArn = aws.String(_iotmanagedintegrationsDeliveryDestinationArn)
	}
	if len(_iotmanagedintegrationsDeliveryDestinationType) > 0 {
		if err := assignInputField(input, "DeliveryDestinationType", _iotmanagedintegrationsDeliveryDestinationType); err != nil {
			log.Errorf("invalid --delivery-destination-type: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsDescription) > 0 {
		input.Description = aws.String(_iotmanagedintegrationsDescription)
	}
	if len(_iotmanagedintegrationsRoleArn) > 0 {
		input.RoleArn = aws.String(_iotmanagedintegrationsRoleArn)
	}

	if resp, err := client.UpdateDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update an event log configuration by log configuration ID.
func iotmanagedintegrations_UpdateEventLogConfiguration(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.UpdateEventLogConfigurationInput{
		// EventLogLevel: types.LogLevel, // Required
		// Id: *string, // Required
	}

	if len(_iotmanagedintegrationsEventLogLevel) > 0 {
		if err := assignInputField(input, "EventLogLevel", _iotmanagedintegrationsEventLogLevel); err != nil {
			log.Errorf("invalid --event-log-level: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsId) > 0 {
		input.Id = aws.String(_iotmanagedintegrationsId)
	}

	if resp, err := client.UpdateEventLogConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update the attributes and capabilities associated with a managed thing.
func iotmanagedintegrations_UpdateManagedThing(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.UpdateManagedThingInput{
		// Identifier: *string, // Required
	}

	if len(_iotmanagedintegrationsIdentifier) > 0 {
		input.Identifier = aws.String(_iotmanagedintegrationsIdentifier)
	}
	if len(_iotmanagedintegrationsBrand) > 0 {
		input.Brand = aws.String(_iotmanagedintegrationsBrand)
	}
	if len(_iotmanagedintegrationsCapabilities) > 0 {
		input.Capabilities = aws.String(_iotmanagedintegrationsCapabilities)
	}
	if len(_iotmanagedintegrationsCapabilityReport) > 0 {
		if err := assignInputField(input, "CapabilityReport", _iotmanagedintegrationsCapabilityReport); err != nil {
			log.Errorf("invalid --capability-report: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsCapabilitySchemas) > 0 {
		if err := assignInputField(input, "CapabilitySchemas", _iotmanagedintegrationsCapabilitySchemas); err != nil {
			log.Errorf("invalid --capability-schemas: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsClassification) > 0 {
		input.Classification = aws.String(_iotmanagedintegrationsClassification)
	}
	if len(_iotmanagedintegrationsCredentialLockerId) > 0 {
		input.CredentialLockerId = aws.String(_iotmanagedintegrationsCredentialLockerId)
	}
	if len(_iotmanagedintegrationsHubNetworkMode) > 0 {
		if err := assignInputField(input, "HubNetworkMode", _iotmanagedintegrationsHubNetworkMode); err != nil {
			log.Errorf("invalid --hub-network-mode: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsMetaData) > 0 {
		if err := assignInputField(input, "MetaData", _iotmanagedintegrationsMetaData); err != nil {
			log.Errorf("invalid --meta-data: %s", err.Error())
			return
		}
	}
	if len(_iotmanagedintegrationsModel) > 0 {
		input.Model = aws.String(_iotmanagedintegrationsModel)
	}
	if len(_iotmanagedintegrationsName) > 0 {
		input.Name = aws.String(_iotmanagedintegrationsName)
	}
	if len(_iotmanagedintegrationsOwner) > 0 {
		input.Owner = aws.String(_iotmanagedintegrationsOwner)
	}
	if len(_iotmanagedintegrationsSerialNumber) > 0 {
		input.SerialNumber = aws.String(_iotmanagedintegrationsSerialNumber)
	}
	if len(_iotmanagedintegrationsWiFiSimpleSetupConfiguration) > 0 {
		if err := assignInputField(input, "WiFiSimpleSetupConfiguration", _iotmanagedintegrationsWiFiSimpleSetupConfiguration); err != nil {
			log.Errorf("invalid --wi-fi-simple-setup-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateManagedThing(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update a notification configuration.
func iotmanagedintegrations_UpdateNotificationConfiguration(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.UpdateNotificationConfigurationInput{
		// DestinationName: *string, // Required
		// EventType: types.EventType, // Required
	}

	if len(_iotmanagedintegrationsDestinationName) > 0 {
		input.DestinationName = aws.String(_iotmanagedintegrationsDestinationName)
	}
	if len(_iotmanagedintegrationsEventType) > 0 {
		if err := assignInputField(input, "EventType", _iotmanagedintegrationsEventType); err != nil {
			log.Errorf("invalid --event-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateNotificationConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update an over-the-air (OTA) task.
func iotmanagedintegrations_UpdateOtaTask(cfg aws.Config, client *iotmanagedintegrations.Client) {
	input := &iotmanagedintegrations.UpdateOtaTaskInput{
		// Identifier: *string, // Required
	}

	if len(_iotmanagedintegrationsIdentifier) > 0 {
		input.Identifier = aws.String(_iotmanagedintegrationsIdentifier)
	}
	if len(_iotmanagedintegrationsDescription) > 0 {
		input.Description = aws.String(_iotmanagedintegrationsDescription)
	}
	if len(_iotmanagedintegrationsTaskConfigurationId) > 0 {
		input.TaskConfigurationId = aws.String(_iotmanagedintegrationsTaskConfigurationId)
	}

	if resp, err := client.UpdateOtaTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_iotmanagedintegrationsCmd)
	_iotmanagedintegrationsCmd.Flags().SortFlags = false

	_iotmanagedintegrationsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_iotmanagedintegrationsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsAccountAssociationId, "account-association-id", "", "", "Account Association ID")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsAuthConfig, "auth-config", "", "", "Auth Config")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsAuthType, "auth-type", "", "", "Auth Type")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsAuthenticationMaterial, "authentication-material", "", "", "Authentication Material")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsAuthenticationMaterialType, "authentication-material-type", "", "", "Authentication Material Type")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsBrand, "brand", "", "", "Brand")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsCaCertificate, "ca-certificate", "", "", "Ca Certificate")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsCapabilities, "capabilities", "", "", "Capabilities")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsCapabilityIdFilter, "capability-id-filter", "", "", "Capability ID Filter")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsCapabilityReport, "capability-report", "", "", "Capability Report")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsCapabilitySchemas, "capability-schemas", "", "", "Capability Schemas")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsClaimCertificate, "claim-certificate", "", "", "Claim Certificate")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsClassification, "classification", "", "", "Classification")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsClientToken, "client-token", "", "", "Client Token")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsCloudConnectorId, "cloud-connector-id", "", "", "Cloud Connector ID")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsConnectorAssociationId, "connector-association-id", "", "", "Connector Association ID")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsConnectorAssociationIdentifier, "connector-association-identifier", "", "", "Connector Association Identifier")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsConnectorDestinationId, "connector-destination-id", "", "", "Connector Destination ID")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsConnectorDestinationIdFilter, "connector-destination-id-filter", "", "", "Connector Destination ID Filter")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsConnectorDeviceId, "connector-device-id", "", "", "Connector Device ID")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsConnectorDeviceIdFilter, "connector-device-id-filter", "", "", "Connector Device ID Filter")
	_iotmanagedintegrationsCmd.Flags().StringSliceVarP(&_iotmanagedintegrationsConnectorDeviceIdList, "connector-device-id-list", "", nil, "Connector Device ID List")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsConnectorId, "connector-id", "", "", "Connector ID")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsConnectorPolicyIdFilter, "connector-policy-id-filter", "", "", "Connector Policy ID Filter")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsControllerIdentifier, "controller-identifier", "", "", "Controller Identifier")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsCredentialLockerFilter, "credential-locker-filter", "", "", "Credential Locker Filter")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsCredentialLockerId, "credential-locker-id", "", "", "Credential Locker ID")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsCustomProtocolDetail, "custom-protocol-detail", "", "", "Custom Protocol Detail")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsDeliveryDestinationArn, "delivery-destination-arn", "", "", "Delivery Destination ARN")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsDeliveryDestinationType, "delivery-destination-type", "", "", "Delivery Destination Type")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsDescription, "description", "", "", "Description")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsDestinationName, "destination-name", "", "", "Destination Name")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsDeviceDiscoveryId, "device-discovery-id", "", "", "Device Discovery ID")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsDevices, "devices", "", "", "Devices")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsDiscoveryType, "discovery-type", "", "", "Discovery Type")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsEncryptionType, "encryption-type", "", "", "Encryption Type")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsEndDeviceIdentifier, "end-device-identifier", "", "", "End Device Identifier")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsEndpointConfig, "endpoint-config", "", "", "Endpoint Config")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsEndpointIdFilter, "endpoint-id-filter", "", "", "Endpoint ID Filter")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsEndpointType, "endpoint-type", "", "", "Endpoint Type")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsEndpoints, "endpoints", "", "", "Endpoints")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsEventLogLevel, "event-log-level", "", "", "Event Log Level")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsEventType, "event-type", "", "", "Event Type")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsForce, "force", "", "", "Force")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsFormat, "format", "", "", "Format")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsGeneralAuthorization, "general-authorization", "", "", "General Authorization")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsHubNetworkMode, "hub-network-mode", "", "", "Hub Network Mode")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsHubTokenTimerExpirySettingInSeconds, "hub-token-timer-expiry-setting-in-seconds", "", "", "Hub Token Timer Expiry Setting In Seconds")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsId, "id", "", "", "ID")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsIdentifier, "identifier", "", "", "Identifier")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsKmsKeyArn, "kms-key-arn", "", "", "KMS Key ARN")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsLambdaArn, "lambda-arn", "", "", "Lambda ARN")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsManagedThingId, "managed-thing-id", "", "", "Managed Thing ID")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsMatterEndpoint, "matter-endpoint", "", "", "Matter Endpoint")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsMaxResults, "max-results", "", "", "Max Results")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsMessage, "message", "", "", "Message")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsMetaData, "meta-data", "", "", "Meta Data")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsModel, "model", "", "", "Model")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsName, "name", "", "", "Name")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsNamespace, "namespace", "", "", "Namespace")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsNextToken, "next-token", "", "", "Next Token")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsOperation, "operation", "", "", "Operation")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsOperationVersion, "operation-version", "", "", "Operation Version")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsOtaMechanism, "ota-mechanism", "", "", "Ota Mechanism")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsOtaSchedulingConfig, "ota-scheduling-config", "", "", "Ota Scheduling Config")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsOtaTargetQueryString, "ota-target-query-string", "", "", "Ota Target Query String")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsOtaTaskExecutionRetryConfig, "ota-task-execution-retry-config", "", "", "Ota Task Execution Retry Config")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsOtaType, "ota-type", "", "", "Ota Type")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsOwner, "owner", "", "", "Owner")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsOwnerFilter, "owner-filter", "", "", "Owner Filter")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsParentControllerIdentifierFilter, "parent-controller-identifier-filter", "", "", "Parent Controller Identifier Filter")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsProtocol, "protocol", "", "", "Protocol")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsProvisioningStatusFilter, "provisioning-status-filter", "", "", "Provisioning Status Filter")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsProvisioningType, "provisioning-type", "", "", "Provisioning Type")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsPushConfig, "push-config", "", "", "Push Config")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsResourceArn, "resource-arn", "", "", "Resource ARN")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsResourceId, "resource-id", "", "", "Resource ID")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsResourceType, "resource-type", "", "", "Resource Type")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsRole, "role", "", "", "Role")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsRoleArn, "role-arn", "", "", "Role ARN")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsRoleFilter, "role-filter", "", "", "Role Filter")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsRuntimeLogConfigurations, "runtime-log-configurations", "", "", "Runtime Log Configurations")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsS3Url, "s3-url", "", "", "S3 URL")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsSchemaId, "schema-id", "", "", "Schema ID")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsSchemaVersionedId, "schema-versioned-id", "", "", "Schema Versioned ID")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsSecretsManager, "secrets-manager", "", "", "Secrets Manager")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsSemanticVersion, "semantic-version", "", "", "Semantic Version")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsSerialNumber, "serial-number", "", "", "Serial Number")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsSerialNumberFilter, "serial-number-filter", "", "", "Serial Number Filter")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsStatusCode, "status-code", "", "", "Status Code")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsStatusFilter, "status-filter", "", "", "Status Filter")
	_iotmanagedintegrationsCmd.Flags().StringSliceVarP(&_iotmanagedintegrationsTagKeys, "tag-keys", "", nil, "Tag Keys")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsTags, "tags", "", "", "Tags")
	_iotmanagedintegrationsCmd.Flags().StringSliceVarP(&_iotmanagedintegrationsTarget, "target", "", nil, "Target")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsTaskConfigurationId, "task-configuration-id", "", "", "Task Configuration ID")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsTraceId, "trace-id", "", "", "Trace ID")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsType, "type", "", "", "Type")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsTypeFilter, "type-filter", "", "", "Type Filter")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsUserId, "user-id", "", "", "User ID")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsVisibility, "visibility", "", "", "Visibility")
	_iotmanagedintegrationsCmd.Flags().StringVarP(&_iotmanagedintegrationsWiFiSimpleSetupConfiguration, "wi-fi-simple-setup-configuration", "", "", "Wi Fi Simple Setup Configuration")

	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsCreateAccountAssociation, "create-account-association", "", false, "Create Account Association")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsCreateCloudConnector, "create-cloud-connector", "", false, "Create Cloud Connector")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsCreateConnectorDestination, "create-connector-destination", "", false, "Create Connector Destination")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsCreateCredentialLocker, "create-credential-locker", "", false, "Create Credential Locker")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsCreateDestination, "create-destination", "", false, "Create Destination")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsCreateEventLogConfiguration, "create-event-log-configuration", "", false, "Create Event Log Configuration")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsCreateManagedThing, "create-managed-thing", "", false, "Create Managed Thing")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsCreateNotificationConfiguration, "create-notification-configuration", "", false, "Create Notification Configuration")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsCreateOtaTask, "create-ota-task", "", false, "Create Ota Task")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsCreateOtaTaskConfiguration, "create-ota-task-configuration", "", false, "Create Ota Task Configuration")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsCreateProvisioningProfile, "create-provisioning-profile", "", false, "Create Provisioning Profile")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsDeleteAccountAssociation, "delete-account-association", "", false, "Delete Account Association")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsDeleteCloudConnector, "delete-cloud-connector", "", false, "Delete Cloud Connector")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsDeleteConnectorDestination, "delete-connector-destination", "", false, "Delete Connector Destination")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsDeleteCredentialLocker, "delete-credential-locker", "", false, "Delete Credential Locker")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsDeleteDestination, "delete-destination", "", false, "Delete Destination")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsDeleteEventLogConfiguration, "delete-event-log-configuration", "", false, "Delete Event Log Configuration")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsDeleteManagedThing, "delete-managed-thing", "", false, "Delete Managed Thing")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsDeleteNotificationConfiguration, "delete-notification-configuration", "", false, "Delete Notification Configuration")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsDeleteOtaTask, "delete-ota-task", "", false, "Delete Ota Task")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsDeleteOtaTaskConfiguration, "delete-ota-task-configuration", "", false, "Delete Ota Task Configuration")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsDeleteProvisioningProfile, "delete-provisioning-profile", "", false, "Delete Provisioning Profile")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsDeregisterAccountAssociation, "deregister-account-association", "", false, "Deregister Account Association")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsGetAccountAssociation, "get-account-association", "", false, "Get Account Association")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsGetCloudConnector, "get-cloud-connector", "", false, "Get Cloud Connector")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsGetConnectorDestination, "get-connector-destination", "", false, "Get Connector Destination")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsGetCredentialLocker, "get-credential-locker", "", false, "Get Credential Locker")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsGetCustomEndpoint, "get-custom-endpoint", "", false, "Get Custom Endpoint")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsGetDefaultEncryptionConfiguration, "get-default-encryption-configuration", "", false, "Get Default Encryption Configuration")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsGetDestination, "get-destination", "", false, "Get Destination")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsGetDeviceDiscovery, "get-device-discovery", "", false, "Get Device Discovery")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsGetEventLogConfiguration, "get-event-log-configuration", "", false, "Get Event Log Configuration")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsGetHubConfiguration, "get-hub-configuration", "", false, "Get Hub Configuration")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsGetManagedThing, "get-managed-thing", "", false, "Get Managed Thing")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsGetManagedThingCapabilities, "get-managed-thing-capabilities", "", false, "Get Managed Thing Capabilities")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsGetManagedThingCertificate, "get-managed-thing-certificate", "", false, "Get Managed Thing Certificate")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsGetManagedThingConnectivityData, "get-managed-thing-connectivity-data", "", false, "Get Managed Thing Connectivity Data")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsGetManagedThingMetaData, "get-managed-thing-meta-data", "", false, "Get Managed Thing Meta Data")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsGetManagedThingState, "get-managed-thing-state", "", false, "Get Managed Thing State")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsGetNotificationConfiguration, "get-notification-configuration", "", false, "Get Notification Configuration")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsGetOtaTask, "get-ota-task", "", false, "Get Ota Task")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsGetOtaTaskConfiguration, "get-ota-task-configuration", "", false, "Get Ota Task Configuration")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsGetProvisioningProfile, "get-provisioning-profile", "", false, "Get Provisioning Profile")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsGetRuntimeLogConfiguration, "get-runtime-log-configuration", "", false, "Get Runtime Log Configuration")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsGetSchemaVersion, "get-schema-version", "", false, "Get Schema Version")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsListAccountAssociations, "list-account-associations", "", false, "List Account Associations")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsListCloudConnectors, "list-cloud-connectors", "", false, "List Cloud Connectors")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsListConnectorDestinations, "list-connector-destinations", "", false, "List Connector Destinations")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsListCredentialLockers, "list-credential-lockers", "", false, "List Credential Lockers")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsListDestinations, "list-destinations", "", false, "List Destinations")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsListDeviceDiscoveries, "list-device-discoveries", "", false, "List Device Discoveries")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsListDiscoveredDevices, "list-discovered-devices", "", false, "List Discovered Devices")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsListEventLogConfigurations, "list-event-log-configurations", "", false, "List Event Log Configurations")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsListManagedThingAccountAssociations, "list-managed-thing-account-associations", "", false, "List Managed Thing Account Associations")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsListManagedThingSchemas, "list-managed-thing-schemas", "", false, "List Managed Thing Schemas")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsListManagedThings, "list-managed-things", "", false, "List Managed Things")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsListNotificationConfigurations, "list-notification-configurations", "", false, "List Notification Configurations")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsListOtaTaskConfigurations, "list-ota-task-configurations", "", false, "List Ota Task Configurations")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsListOtaTaskExecutions, "list-ota-task-executions", "", false, "List Ota Task Executions")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsListOtaTasks, "list-ota-tasks", "", false, "List Ota Tasks")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsListProvisioningProfiles, "list-provisioning-profiles", "", false, "List Provisioning Profiles")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsListSchemaVersions, "list-schema-versions", "", false, "List Schema Versions")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsPutDefaultEncryptionConfiguration, "put-default-encryption-configuration", "", false, "Put Default Encryption Configuration")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsPutHubConfiguration, "put-hub-configuration", "", false, "Put Hub Configuration")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsPutRuntimeLogConfiguration, "put-runtime-log-configuration", "", false, "Put Runtime Log Configuration")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsRegisterAccountAssociation, "register-account-association", "", false, "Register Account Association")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsRegisterCustomEndpoint, "register-custom-endpoint", "", false, "Register Custom Endpoint")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsResetRuntimeLogConfiguration, "reset-runtime-log-configuration", "", false, "Reset Runtime Log Configuration")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsSendConnectorEvent, "send-connector-event", "", false, "Send Connector Event")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsSendManagedThingCommand, "send-managed-thing-command", "", false, "Send Managed Thing Command")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsStartAccountAssociationRefresh, "start-account-association-refresh", "", false, "Start Account Association Refresh")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsStartDeviceDiscovery, "start-device-discovery", "", false, "Start Device Discovery")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsTagResource, "tag-resource", "", false, "Tag Resource")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsUntagResource, "untag-resource", "", false, "Untag Resource")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsUpdateAccountAssociation, "update-account-association", "", false, "Update Account Association")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsUpdateCloudConnector, "update-cloud-connector", "", false, "Update Cloud Connector")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsUpdateConnectorDestination, "update-connector-destination", "", false, "Update Connector Destination")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsUpdateDestination, "update-destination", "", false, "Update Destination")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsUpdateEventLogConfiguration, "update-event-log-configuration", "", false, "Update Event Log Configuration")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsUpdateManagedThing, "update-managed-thing", "", false, "Update Managed Thing")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsUpdateNotificationConfiguration, "update-notification-configuration", "", false, "Update Notification Configuration")
	_iotmanagedintegrationsCmd.Flags().BoolVarP(&_iotmanagedintegrationsUpdateOtaTask, "update-ota-task", "", false, "Update Ota Task")

}
