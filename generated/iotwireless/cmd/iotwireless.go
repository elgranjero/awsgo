package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iotwireless"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// iotwirelessCmd represents the iotwireless command
var _iotwirelessCmd = &cobra.Command{
	Use:   "iotwireless",
	Short: "AWS iotwireless CLI",
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
		client := iotwireless.NewFromConfig(cfg)
		if _iotwirelessAssociateAwsAccountWithPartnerAccount {
			iotwireless_AssociateAwsAccountWithPartnerAccount(cfg, client)
			return
		}
		if _iotwirelessAssociateMulticastGroupWithFuotaTask {
			iotwireless_AssociateMulticastGroupWithFuotaTask(cfg, client)
			return
		}
		if _iotwirelessAssociateWirelessDeviceWithFuotaTask {
			iotwireless_AssociateWirelessDeviceWithFuotaTask(cfg, client)
			return
		}
		if _iotwirelessAssociateWirelessDeviceWithMulticastGroup {
			iotwireless_AssociateWirelessDeviceWithMulticastGroup(cfg, client)
			return
		}
		if _iotwirelessAssociateWirelessDeviceWithThing {
			iotwireless_AssociateWirelessDeviceWithThing(cfg, client)
			return
		}
		if _iotwirelessAssociateWirelessGatewayWithCertificate {
			iotwireless_AssociateWirelessGatewayWithCertificate(cfg, client)
			return
		}
		if _iotwirelessAssociateWirelessGatewayWithThing {
			iotwireless_AssociateWirelessGatewayWithThing(cfg, client)
			return
		}
		if _iotwirelessCancelMulticastGroupSession {
			iotwireless_CancelMulticastGroupSession(cfg, client)
			return
		}
		if _iotwirelessCreateDestination {
			iotwireless_CreateDestination(cfg, client)
			return
		}
		if _iotwirelessCreateDeviceProfile {
			iotwireless_CreateDeviceProfile(cfg, client)
			return
		}
		if _iotwirelessCreateFuotaTask {
			iotwireless_CreateFuotaTask(cfg, client)
			return
		}
		if _iotwirelessCreateMulticastGroup {
			iotwireless_CreateMulticastGroup(cfg, client)
			return
		}
		if _iotwirelessCreateNetworkAnalyzerConfiguration {
			iotwireless_CreateNetworkAnalyzerConfiguration(cfg, client)
			return
		}
		if _iotwirelessCreateServiceProfile {
			iotwireless_CreateServiceProfile(cfg, client)
			return
		}
		if _iotwirelessCreateWirelessDevice {
			iotwireless_CreateWirelessDevice(cfg, client)
			return
		}
		if _iotwirelessCreateWirelessGateway {
			iotwireless_CreateWirelessGateway(cfg, client)
			return
		}
		if _iotwirelessCreateWirelessGatewayTask {
			iotwireless_CreateWirelessGatewayTask(cfg, client)
			return
		}
		if _iotwirelessCreateWirelessGatewayTaskDefinition {
			iotwireless_CreateWirelessGatewayTaskDefinition(cfg, client)
			return
		}
		if _iotwirelessDeleteDestination {
			iotwireless_DeleteDestination(cfg, client)
			return
		}
		if _iotwirelessDeleteDeviceProfile {
			iotwireless_DeleteDeviceProfile(cfg, client)
			return
		}
		if _iotwirelessDeleteFuotaTask {
			iotwireless_DeleteFuotaTask(cfg, client)
			return
		}
		if _iotwirelessDeleteMulticastGroup {
			iotwireless_DeleteMulticastGroup(cfg, client)
			return
		}
		if _iotwirelessDeleteNetworkAnalyzerConfiguration {
			iotwireless_DeleteNetworkAnalyzerConfiguration(cfg, client)
			return
		}
		if _iotwirelessDeleteQueuedMessages {
			iotwireless_DeleteQueuedMessages(cfg, client)
			return
		}
		if _iotwirelessDeleteServiceProfile {
			iotwireless_DeleteServiceProfile(cfg, client)
			return
		}
		if _iotwirelessDeleteWirelessDevice {
			iotwireless_DeleteWirelessDevice(cfg, client)
			return
		}
		if _iotwirelessDeleteWirelessDeviceImportTask {
			iotwireless_DeleteWirelessDeviceImportTask(cfg, client)
			return
		}
		if _iotwirelessDeleteWirelessGateway {
			iotwireless_DeleteWirelessGateway(cfg, client)
			return
		}
		if _iotwirelessDeleteWirelessGatewayTask {
			iotwireless_DeleteWirelessGatewayTask(cfg, client)
			return
		}
		if _iotwirelessDeleteWirelessGatewayTaskDefinition {
			iotwireless_DeleteWirelessGatewayTaskDefinition(cfg, client)
			return
		}
		if _iotwirelessDeregisterWirelessDevice {
			iotwireless_DeregisterWirelessDevice(cfg, client)
			return
		}
		if _iotwirelessDisassociateAwsAccountFromPartnerAccount {
			iotwireless_DisassociateAwsAccountFromPartnerAccount(cfg, client)
			return
		}
		if _iotwirelessDisassociateMulticastGroupFromFuotaTask {
			iotwireless_DisassociateMulticastGroupFromFuotaTask(cfg, client)
			return
		}
		if _iotwirelessDisassociateWirelessDeviceFromFuotaTask {
			iotwireless_DisassociateWirelessDeviceFromFuotaTask(cfg, client)
			return
		}
		if _iotwirelessDisassociateWirelessDeviceFromMulticastGroup {
			iotwireless_DisassociateWirelessDeviceFromMulticastGroup(cfg, client)
			return
		}
		if _iotwirelessDisassociateWirelessDeviceFromThing {
			iotwireless_DisassociateWirelessDeviceFromThing(cfg, client)
			return
		}
		if _iotwirelessDisassociateWirelessGatewayFromCertificate {
			iotwireless_DisassociateWirelessGatewayFromCertificate(cfg, client)
			return
		}
		if _iotwirelessDisassociateWirelessGatewayFromThing {
			iotwireless_DisassociateWirelessGatewayFromThing(cfg, client)
			return
		}
		if _iotwirelessGetDestination {
			iotwireless_GetDestination(cfg, client)
			return
		}
		if _iotwirelessGetDeviceProfile {
			iotwireless_GetDeviceProfile(cfg, client)
			return
		}
		if _iotwirelessGetEventConfigurationByResourceTypes {
			iotwireless_GetEventConfigurationByResourceTypes(cfg, client)
			return
		}
		if _iotwirelessGetFuotaTask {
			iotwireless_GetFuotaTask(cfg, client)
			return
		}
		if _iotwirelessGetLogLevelsByResourceTypes {
			iotwireless_GetLogLevelsByResourceTypes(cfg, client)
			return
		}
		if _iotwirelessGetMetricConfiguration {
			iotwireless_GetMetricConfiguration(cfg, client)
			return
		}
		if _iotwirelessGetMetrics {
			iotwireless_GetMetrics(cfg, client)
			return
		}
		if _iotwirelessGetMulticastGroup {
			iotwireless_GetMulticastGroup(cfg, client)
			return
		}
		if _iotwirelessGetMulticastGroupSession {
			iotwireless_GetMulticastGroupSession(cfg, client)
			return
		}
		if _iotwirelessGetNetworkAnalyzerConfiguration {
			iotwireless_GetNetworkAnalyzerConfiguration(cfg, client)
			return
		}
		if _iotwirelessGetPartnerAccount {
			iotwireless_GetPartnerAccount(cfg, client)
			return
		}
		if _iotwirelessGetPosition {
			iotwireless_GetPosition(cfg, client)
			return
		}
		if _iotwirelessGetPositionConfiguration {
			iotwireless_GetPositionConfiguration(cfg, client)
			return
		}
		if _iotwirelessGetPositionEstimate {
			iotwireless_GetPositionEstimate(cfg, client)
			return
		}
		if _iotwirelessGetResourceEventConfiguration {
			iotwireless_GetResourceEventConfiguration(cfg, client)
			return
		}
		if _iotwirelessGetResourceLogLevel {
			iotwireless_GetResourceLogLevel(cfg, client)
			return
		}
		if _iotwirelessGetResourcePosition {
			iotwireless_GetResourcePosition(cfg, client)
			return
		}
		if _iotwirelessGetServiceEndpoint {
			iotwireless_GetServiceEndpoint(cfg, client)
			return
		}
		if _iotwirelessGetServiceProfile {
			iotwireless_GetServiceProfile(cfg, client)
			return
		}
		if _iotwirelessGetWirelessDevice {
			iotwireless_GetWirelessDevice(cfg, client)
			return
		}
		if _iotwirelessGetWirelessDeviceImportTask {
			iotwireless_GetWirelessDeviceImportTask(cfg, client)
			return
		}
		if _iotwirelessGetWirelessDeviceStatistics {
			iotwireless_GetWirelessDeviceStatistics(cfg, client)
			return
		}
		if _iotwirelessGetWirelessGateway {
			iotwireless_GetWirelessGateway(cfg, client)
			return
		}
		if _iotwirelessGetWirelessGatewayCertificate {
			iotwireless_GetWirelessGatewayCertificate(cfg, client)
			return
		}
		if _iotwirelessGetWirelessGatewayFirmwareInformation {
			iotwireless_GetWirelessGatewayFirmwareInformation(cfg, client)
			return
		}
		if _iotwirelessGetWirelessGatewayStatistics {
			iotwireless_GetWirelessGatewayStatistics(cfg, client)
			return
		}
		if _iotwirelessGetWirelessGatewayTask {
			iotwireless_GetWirelessGatewayTask(cfg, client)
			return
		}
		if _iotwirelessGetWirelessGatewayTaskDefinition {
			iotwireless_GetWirelessGatewayTaskDefinition(cfg, client)
			return
		}
		if _iotwirelessListDestinations {
			iotwireless_ListDestinations(cfg, client)
			return
		}
		if _iotwirelessListDeviceProfiles {
			iotwireless_ListDeviceProfiles(cfg, client)
			return
		}
		if _iotwirelessListDevicesForWirelessDeviceImportTask {
			iotwireless_ListDevicesForWirelessDeviceImportTask(cfg, client)
			return
		}
		if _iotwirelessListEventConfigurations {
			iotwireless_ListEventConfigurations(cfg, client)
			return
		}
		if _iotwirelessListFuotaTasks {
			iotwireless_ListFuotaTasks(cfg, client)
			return
		}
		if _iotwirelessListMulticastGroups {
			iotwireless_ListMulticastGroups(cfg, client)
			return
		}
		if _iotwirelessListMulticastGroupsByFuotaTask {
			iotwireless_ListMulticastGroupsByFuotaTask(cfg, client)
			return
		}
		if _iotwirelessListNetworkAnalyzerConfigurations {
			iotwireless_ListNetworkAnalyzerConfigurations(cfg, client)
			return
		}
		if _iotwirelessListPartnerAccounts {
			iotwireless_ListPartnerAccounts(cfg, client)
			return
		}
		if _iotwirelessListPositionConfigurations {
			iotwireless_ListPositionConfigurations(cfg, client)
			return
		}
		if _iotwirelessListQueuedMessages {
			iotwireless_ListQueuedMessages(cfg, client)
			return
		}
		if _iotwirelessListServiceProfiles {
			iotwireless_ListServiceProfiles(cfg, client)
			return
		}
		if _iotwirelessListTagsForResource {
			iotwireless_ListTagsForResource(cfg, client)
			return
		}
		if _iotwirelessListWirelessDeviceImportTasks {
			iotwireless_ListWirelessDeviceImportTasks(cfg, client)
			return
		}
		if _iotwirelessListWirelessDevices {
			iotwireless_ListWirelessDevices(cfg, client)
			return
		}
		if _iotwirelessListWirelessGatewayTaskDefinitions {
			iotwireless_ListWirelessGatewayTaskDefinitions(cfg, client)
			return
		}
		if _iotwirelessListWirelessGateways {
			iotwireless_ListWirelessGateways(cfg, client)
			return
		}
		if _iotwirelessPutPositionConfiguration {
			iotwireless_PutPositionConfiguration(cfg, client)
			return
		}
		if _iotwirelessPutResourceLogLevel {
			iotwireless_PutResourceLogLevel(cfg, client)
			return
		}
		if _iotwirelessResetAllResourceLogLevels {
			iotwireless_ResetAllResourceLogLevels(cfg, client)
			return
		}
		if _iotwirelessResetResourceLogLevel {
			iotwireless_ResetResourceLogLevel(cfg, client)
			return
		}
		if _iotwirelessSendDataToMulticastGroup {
			iotwireless_SendDataToMulticastGroup(cfg, client)
			return
		}
		if _iotwirelessSendDataToWirelessDevice {
			iotwireless_SendDataToWirelessDevice(cfg, client)
			return
		}
		if _iotwirelessStartBulkAssociateWirelessDeviceWithMulticastGroup {
			iotwireless_StartBulkAssociateWirelessDeviceWithMulticastGroup(cfg, client)
			return
		}
		if _iotwirelessStartBulkDisassociateWirelessDeviceFromMulticastGroup {
			iotwireless_StartBulkDisassociateWirelessDeviceFromMulticastGroup(cfg, client)
			return
		}
		if _iotwirelessStartFuotaTask {
			iotwireless_StartFuotaTask(cfg, client)
			return
		}
		if _iotwirelessStartMulticastGroupSession {
			iotwireless_StartMulticastGroupSession(cfg, client)
			return
		}
		if _iotwirelessStartSingleWirelessDeviceImportTask {
			iotwireless_StartSingleWirelessDeviceImportTask(cfg, client)
			return
		}
		if _iotwirelessStartWirelessDeviceImportTask {
			iotwireless_StartWirelessDeviceImportTask(cfg, client)
			return
		}
		if _iotwirelessTagResource {
			iotwireless_TagResource(cfg, client)
			return
		}
		if _iotwirelessTestWirelessDevice {
			iotwireless_TestWirelessDevice(cfg, client)
			return
		}
		if _iotwirelessUntagResource {
			iotwireless_UntagResource(cfg, client)
			return
		}
		if _iotwirelessUpdateDestination {
			iotwireless_UpdateDestination(cfg, client)
			return
		}
		if _iotwirelessUpdateEventConfigurationByResourceTypes {
			iotwireless_UpdateEventConfigurationByResourceTypes(cfg, client)
			return
		}
		if _iotwirelessUpdateFuotaTask {
			iotwireless_UpdateFuotaTask(cfg, client)
			return
		}
		if _iotwirelessUpdateLogLevelsByResourceTypes {
			iotwireless_UpdateLogLevelsByResourceTypes(cfg, client)
			return
		}
		if _iotwirelessUpdateMetricConfiguration {
			iotwireless_UpdateMetricConfiguration(cfg, client)
			return
		}
		if _iotwirelessUpdateMulticastGroup {
			iotwireless_UpdateMulticastGroup(cfg, client)
			return
		}
		if _iotwirelessUpdateNetworkAnalyzerConfiguration {
			iotwireless_UpdateNetworkAnalyzerConfiguration(cfg, client)
			return
		}
		if _iotwirelessUpdatePartnerAccount {
			iotwireless_UpdatePartnerAccount(cfg, client)
			return
		}
		if _iotwirelessUpdatePosition {
			iotwireless_UpdatePosition(cfg, client)
			return
		}
		if _iotwirelessUpdateResourceEventConfiguration {
			iotwireless_UpdateResourceEventConfiguration(cfg, client)
			return
		}
		if _iotwirelessUpdateResourcePosition {
			iotwireless_UpdateResourcePosition(cfg, client)
			return
		}
		if _iotwirelessUpdateWirelessDevice {
			iotwireless_UpdateWirelessDevice(cfg, client)
			return
		}
		if _iotwirelessUpdateWirelessDeviceImportTask {
			iotwireless_UpdateWirelessDeviceImportTask(cfg, client)
			return
		}
		if _iotwirelessUpdateWirelessGateway {
			iotwireless_UpdateWirelessGateway(cfg, client)
			return
		}

	},
}

var (
	_iotwirelessAssociateAwsAccountWithPartnerAccount                 bool
	_iotwirelessAssociateMulticastGroupWithFuotaTask                  bool
	_iotwirelessAssociateWirelessDeviceWithFuotaTask                  bool
	_iotwirelessAssociateWirelessDeviceWithMulticastGroup             bool
	_iotwirelessAssociateWirelessDeviceWithThing                      bool
	_iotwirelessAssociateWirelessGatewayWithCertificate               bool
	_iotwirelessAssociateWirelessGatewayWithThing                     bool
	_iotwirelessCancelMulticastGroupSession                           bool
	_iotwirelessCreateDestination                                     bool
	_iotwirelessCreateDeviceProfile                                   bool
	_iotwirelessCreateFuotaTask                                       bool
	_iotwirelessCreateMulticastGroup                                  bool
	_iotwirelessCreateNetworkAnalyzerConfiguration                    bool
	_iotwirelessCreateServiceProfile                                  bool
	_iotwirelessCreateWirelessDevice                                  bool
	_iotwirelessCreateWirelessGateway                                 bool
	_iotwirelessCreateWirelessGatewayTask                             bool
	_iotwirelessCreateWirelessGatewayTaskDefinition                   bool
	_iotwirelessDeleteDestination                                     bool
	_iotwirelessDeleteDeviceProfile                                   bool
	_iotwirelessDeleteFuotaTask                                       bool
	_iotwirelessDeleteMulticastGroup                                  bool
	_iotwirelessDeleteNetworkAnalyzerConfiguration                    bool
	_iotwirelessDeleteQueuedMessages                                  bool
	_iotwirelessDeleteServiceProfile                                  bool
	_iotwirelessDeleteWirelessDevice                                  bool
	_iotwirelessDeleteWirelessDeviceImportTask                        bool
	_iotwirelessDeleteWirelessGateway                                 bool
	_iotwirelessDeleteWirelessGatewayTask                             bool
	_iotwirelessDeleteWirelessGatewayTaskDefinition                   bool
	_iotwirelessDeregisterWirelessDevice                              bool
	_iotwirelessDisassociateAwsAccountFromPartnerAccount              bool
	_iotwirelessDisassociateMulticastGroupFromFuotaTask               bool
	_iotwirelessDisassociateWirelessDeviceFromFuotaTask               bool
	_iotwirelessDisassociateWirelessDeviceFromMulticastGroup          bool
	_iotwirelessDisassociateWirelessDeviceFromThing                   bool
	_iotwirelessDisassociateWirelessGatewayFromCertificate            bool
	_iotwirelessDisassociateWirelessGatewayFromThing                  bool
	_iotwirelessGetDestination                                        bool
	_iotwirelessGetDeviceProfile                                      bool
	_iotwirelessGetEventConfigurationByResourceTypes                  bool
	_iotwirelessGetFuotaTask                                          bool
	_iotwirelessGetLogLevelsByResourceTypes                           bool
	_iotwirelessGetMetricConfiguration                                bool
	_iotwirelessGetMetrics                                            bool
	_iotwirelessGetMulticastGroup                                     bool
	_iotwirelessGetMulticastGroupSession                              bool
	_iotwirelessGetNetworkAnalyzerConfiguration                       bool
	_iotwirelessGetPartnerAccount                                     bool
	_iotwirelessGetPosition                                           bool
	_iotwirelessGetPositionConfiguration                              bool
	_iotwirelessGetPositionEstimate                                   bool
	_iotwirelessGetResourceEventConfiguration                         bool
	_iotwirelessGetResourceLogLevel                                   bool
	_iotwirelessGetResourcePosition                                   bool
	_iotwirelessGetServiceEndpoint                                    bool
	_iotwirelessGetServiceProfile                                     bool
	_iotwirelessGetWirelessDevice                                     bool
	_iotwirelessGetWirelessDeviceImportTask                           bool
	_iotwirelessGetWirelessDeviceStatistics                           bool
	_iotwirelessGetWirelessGateway                                    bool
	_iotwirelessGetWirelessGatewayCertificate                         bool
	_iotwirelessGetWirelessGatewayFirmwareInformation                 bool
	_iotwirelessGetWirelessGatewayStatistics                          bool
	_iotwirelessGetWirelessGatewayTask                                bool
	_iotwirelessGetWirelessGatewayTaskDefinition                      bool
	_iotwirelessListDestinations                                      bool
	_iotwirelessListDeviceProfiles                                    bool
	_iotwirelessListDevicesForWirelessDeviceImportTask                bool
	_iotwirelessListEventConfigurations                               bool
	_iotwirelessListFuotaTasks                                        bool
	_iotwirelessListMulticastGroups                                   bool
	_iotwirelessListMulticastGroupsByFuotaTask                        bool
	_iotwirelessListNetworkAnalyzerConfigurations                     bool
	_iotwirelessListPartnerAccounts                                   bool
	_iotwirelessListPositionConfigurations                            bool
	_iotwirelessListQueuedMessages                                    bool
	_iotwirelessListServiceProfiles                                   bool
	_iotwirelessListTagsForResource                                   bool
	_iotwirelessListWirelessDeviceImportTasks                         bool
	_iotwirelessListWirelessDevices                                   bool
	_iotwirelessListWirelessGatewayTaskDefinitions                    bool
	_iotwirelessListWirelessGateways                                  bool
	_iotwirelessPutPositionConfiguration                              bool
	_iotwirelessPutResourceLogLevel                                   bool
	_iotwirelessResetAllResourceLogLevels                             bool
	_iotwirelessResetResourceLogLevel                                 bool
	_iotwirelessSendDataToMulticastGroup                              bool
	_iotwirelessSendDataToWirelessDevice                              bool
	_iotwirelessStartBulkAssociateWirelessDeviceWithMulticastGroup    bool
	_iotwirelessStartBulkDisassociateWirelessDeviceFromMulticastGroup bool
	_iotwirelessStartFuotaTask                                        bool
	_iotwirelessStartMulticastGroupSession                            bool
	_iotwirelessStartSingleWirelessDeviceImportTask                   bool
	_iotwirelessStartWirelessDeviceImportTask                         bool
	_iotwirelessTagResource                                           bool
	_iotwirelessTestWirelessDevice                                    bool
	_iotwirelessUntagResource                                         bool
	_iotwirelessUpdateDestination                                     bool
	_iotwirelessUpdateEventConfigurationByResourceTypes               bool
	_iotwirelessUpdateFuotaTask                                       bool
	_iotwirelessUpdateLogLevelsByResourceTypes                        bool
	_iotwirelessUpdateMetricConfiguration                             bool
	_iotwirelessUpdateMulticastGroup                                  bool
	_iotwirelessUpdateNetworkAnalyzerConfiguration                    bool
	_iotwirelessUpdatePartnerAccount                                  bool
	_iotwirelessUpdatePosition                                        bool
	_iotwirelessUpdateResourceEventConfiguration                      bool
	_iotwirelessUpdateResourcePosition                                bool
	_iotwirelessUpdateWirelessDevice                                  bool
	_iotwirelessUpdateWirelessDeviceImportTask                        bool
	_iotwirelessUpdateWirelessGateway                                 bool

	_iotwirelessAutoCreateTasks                 string
	_iotwirelessCellTowers                      string
	_iotwirelessClientRequestToken              string
	_iotwirelessConfigurationName               string
	_iotwirelessConnectionStatus                string
	_iotwirelessDefaultLogLevel                 string
	_iotwirelessDescription                     string
	_iotwirelessDescriptor                      string
	_iotwirelessDestination                     string
	_iotwirelessDestinationName                 string
	_iotwirelessDeviceName                      string
	_iotwirelessDeviceProfileId                 string
	_iotwirelessDeviceProfileType               string
	_iotwirelessDeviceRegistrationState         string
	_iotwirelessExpression                      string
	_iotwirelessExpressionType                  string
	_iotwirelessFirmwareUpdateImage             string
	_iotwirelessFirmwareUpdateRole              string
	_iotwirelessFragmentIntervalMS              string
	_iotwirelessFragmentSizeBytes               string
	_iotwirelessFuotaTaskId                     string
	_iotwirelessFuotaTaskLogOptions             string
	_iotwirelessGeoJsonPayload                  string
	_iotwirelessGnss                            string
	_iotwirelessId                              string
	_iotwirelessIdentifier                      string
	_iotwirelessIdentifierType                  string
	_iotwirelessIotCertificateId                string
	_iotwirelessIp                              string
	_iotwirelessJoin                            string
	_iotwirelessJoinEuiFilters                  string
	_iotwirelessLoRaWAN                         string
	_iotwirelessLogLevel                        string
	_iotwirelessMaxEirp                         string
	_iotwirelessMaxResults                      string
	_iotwirelessMessageDeliveryStatus           string
	_iotwirelessMessageId                       string
	_iotwirelessMulticastGroupId                string
	_iotwirelessMulticastGroups                 []string
	_iotwirelessMulticastGroupsToAdd            []string
	_iotwirelessMulticastGroupsToRemove         []string
	_iotwirelessName                            string
	_iotwirelessNetIdFilters                    []string
	_iotwirelessNextToken                       string
	_iotwirelessPartnerAccountId                string
	_iotwirelessPartnerType                     string
	_iotwirelessPayloadData                     string
	_iotwirelessPosition                        string
	_iotwirelessPositioning                     string
	_iotwirelessProximity                       string
	_iotwirelessQueryString                     string
	_iotwirelessRedundancyPercent               string
	_iotwirelessResourceArn                     string
	_iotwirelessResourceIdentifier              string
	_iotwirelessResourceType                    string
	_iotwirelessRoleArn                         string
	_iotwirelessServiceProfileId                string
	_iotwirelessServiceType                     string
	_iotwirelessSidewalk                        string
	_iotwirelessSolvers                         string
	_iotwirelessStatus                          string
	_iotwirelessSummaryMetric                   string
	_iotwirelessSummaryMetricQueries            string
	_iotwirelessTagKeys                         []string
	_iotwirelessTags                            string
	_iotwirelessTaskDefinitionType              string
	_iotwirelessThingArn                        string
	_iotwirelessTimestamp                       string
	_iotwirelessTraceContent                    string
	_iotwirelessTransmitMode                    string
	_iotwirelessType                            string
	_iotwirelessUpdate                          string
	_iotwirelessWiFiAccessPoints                string
	_iotwirelessWirelessDeviceId                string
	_iotwirelessWirelessDeviceLogOptions        string
	_iotwirelessWirelessDeviceType              string
	_iotwirelessWirelessDevices                 []string
	_iotwirelessWirelessDevicesToAdd            []string
	_iotwirelessWirelessDevicesToRemove         []string
	_iotwirelessWirelessGatewayId               string
	_iotwirelessWirelessGatewayLogOptions       string
	_iotwirelessWirelessGatewayTaskDefinitionId string
	_iotwirelessWirelessGateways                []string
	_iotwirelessWirelessGatewaysToAdd           []string
	_iotwirelessWirelessGatewaysToRemove        []string
	_iotwirelessWirelessMetadata                string
)

// Associates a partner account with your AWS account.
func iotwireless_AssociateAwsAccountWithPartnerAccount(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.AssociateAwsAccountWithPartnerAccountInput{
		// Sidewalk: *types.SidewalkAccountInfo, // Required
	}

	if len(_iotwirelessSidewalk) > 0 {
		if err := assignInputField(input, "Sidewalk", _iotwirelessSidewalk); err != nil {
			log.Errorf("invalid --sidewalk: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_iotwirelessClientRequestToken)
	}
	if len(_iotwirelessTags) > 0 {
		if err := assignInputField(input, "Tags", _iotwirelessTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.AssociateAwsAccountWithPartnerAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associate a multicast group with a FUOTA task.
func iotwireless_AssociateMulticastGroupWithFuotaTask(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.AssociateMulticastGroupWithFuotaTaskInput{
		// Id: *string, // Required
		// MulticastGroupId: *string, // Required
	}

	if len(_iotwirelessId) > 0 {
		input.Id = aws.String(_iotwirelessId)
	}
	if len(_iotwirelessMulticastGroupId) > 0 {
		input.MulticastGroupId = aws.String(_iotwirelessMulticastGroupId)
	}

	if resp, err := client.AssociateMulticastGroupWithFuotaTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associate a wireless device with a FUOTA task.
func iotwireless_AssociateWirelessDeviceWithFuotaTask(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.AssociateWirelessDeviceWithFuotaTaskInput{
		// Id: *string, // Required
		// WirelessDeviceId: *string, // Required
	}

	if len(_iotwirelessId) > 0 {
		input.Id = aws.String(_iotwirelessId)
	}
	if len(_iotwirelessWirelessDeviceId) > 0 {
		input.WirelessDeviceId = aws.String(_iotwirelessWirelessDeviceId)
	}

	if resp, err := client.AssociateWirelessDeviceWithFuotaTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a wireless device with a multicast group.
func iotwireless_AssociateWirelessDeviceWithMulticastGroup(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.AssociateWirelessDeviceWithMulticastGroupInput{
		// Id: *string, // Required
		// WirelessDeviceId: *string, // Required
	}

	if len(_iotwirelessId) > 0 {
		input.Id = aws.String(_iotwirelessId)
	}
	if len(_iotwirelessWirelessDeviceId) > 0 {
		input.WirelessDeviceId = aws.String(_iotwirelessWirelessDeviceId)
	}

	if resp, err := client.AssociateWirelessDeviceWithMulticastGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a wireless device with a thing.
func iotwireless_AssociateWirelessDeviceWithThing(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.AssociateWirelessDeviceWithThingInput{
		// Id: *string, // Required
		// ThingArn: *string, // Required
	}

	if len(_iotwirelessId) > 0 {
		input.Id = aws.String(_iotwirelessId)
	}
	if len(_iotwirelessThingArn) > 0 {
		input.ThingArn = aws.String(_iotwirelessThingArn)
	}

	if resp, err := client.AssociateWirelessDeviceWithThing(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a wireless gateway with a certificate.
func iotwireless_AssociateWirelessGatewayWithCertificate(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.AssociateWirelessGatewayWithCertificateInput{
		// Id: *string, // Required
		// IotCertificateId: *string, // Required
	}

	if len(_iotwirelessId) > 0 {
		input.Id = aws.String(_iotwirelessId)
	}
	if len(_iotwirelessIotCertificateId) > 0 {
		input.IotCertificateId = aws.String(_iotwirelessIotCertificateId)
	}

	if resp, err := client.AssociateWirelessGatewayWithCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a wireless gateway with a thing.
func iotwireless_AssociateWirelessGatewayWithThing(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.AssociateWirelessGatewayWithThingInput{
		// Id: *string, // Required
		// ThingArn: *string, // Required
	}

	if len(_iotwirelessId) > 0 {
		input.Id = aws.String(_iotwirelessId)
	}
	if len(_iotwirelessThingArn) > 0 {
		input.ThingArn = aws.String(_iotwirelessThingArn)
	}

	if resp, err := client.AssociateWirelessGatewayWithThing(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels an existing multicast group session.
func iotwireless_CancelMulticastGroupSession(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.CancelMulticastGroupSessionInput{
		// Id: *string, // Required
	}

	if len(_iotwirelessId) > 0 {
		input.Id = aws.String(_iotwirelessId)
	}

	if resp, err := client.CancelMulticastGroupSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new destination that maps a device message to an AWS IoT rule.
func iotwireless_CreateDestination(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.CreateDestinationInput{
		// Expression: *string, // Required
		// ExpressionType: types.ExpressionType, // Required
		// Name: *string, // Required
		// RoleArn: *string, // Required
	}

	if len(_iotwirelessExpression) > 0 {
		input.Expression = aws.String(_iotwirelessExpression)
	}
	if len(_iotwirelessExpressionType) > 0 {
		if err := assignInputField(input, "ExpressionType", _iotwirelessExpressionType); err != nil {
			log.Errorf("invalid --expression-type: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessName) > 0 {
		input.Name = aws.String(_iotwirelessName)
	}
	if len(_iotwirelessRoleArn) > 0 {
		input.RoleArn = aws.String(_iotwirelessRoleArn)
	}
	if len(_iotwirelessClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_iotwirelessClientRequestToken)
	}
	if len(_iotwirelessDescription) > 0 {
		input.Description = aws.String(_iotwirelessDescription)
	}
	if len(_iotwirelessTags) > 0 {
		if err := assignInputField(input, "Tags", _iotwirelessTags); err != nil {
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

// Creates a new device profile.
func iotwireless_CreateDeviceProfile(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.CreateDeviceProfileInput{}

	if len(_iotwirelessClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_iotwirelessClientRequestToken)
	}
	if len(_iotwirelessLoRaWAN) > 0 {
		if err := assignInputField(input, "LoRaWAN", _iotwirelessLoRaWAN); err != nil {
			log.Errorf("invalid --lo-ra-wan: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessName) > 0 {
		input.Name = aws.String(_iotwirelessName)
	}
	if len(_iotwirelessSidewalk) > 0 {
		if err := assignInputField(input, "Sidewalk", _iotwirelessSidewalk); err != nil {
			log.Errorf("invalid --sidewalk: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessTags) > 0 {
		if err := assignInputField(input, "Tags", _iotwirelessTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDeviceProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a FUOTA task.
func iotwireless_CreateFuotaTask(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.CreateFuotaTaskInput{
		// FirmwareUpdateImage: *string, // Required
		// FirmwareUpdateRole: *string, // Required
	}

	if len(_iotwirelessFirmwareUpdateImage) > 0 {
		input.FirmwareUpdateImage = aws.String(_iotwirelessFirmwareUpdateImage)
	}
	if len(_iotwirelessFirmwareUpdateRole) > 0 {
		input.FirmwareUpdateRole = aws.String(_iotwirelessFirmwareUpdateRole)
	}
	if len(_iotwirelessClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_iotwirelessClientRequestToken)
	}
	if len(_iotwirelessDescription) > 0 {
		input.Description = aws.String(_iotwirelessDescription)
	}
	if len(_iotwirelessDescriptor) > 0 {
		input.Descriptor = aws.String(_iotwirelessDescriptor)
	}
	if len(_iotwirelessFragmentIntervalMS) > 0 {
		if err := assignInputField(input, "FragmentIntervalMS", _iotwirelessFragmentIntervalMS); err != nil {
			log.Errorf("invalid --fragment-interval-ms: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessFragmentSizeBytes) > 0 {
		if err := assignInputField(input, "FragmentSizeBytes", _iotwirelessFragmentSizeBytes); err != nil {
			log.Errorf("invalid --fragment-size-bytes: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessLoRaWAN) > 0 {
		if err := assignInputField(input, "LoRaWAN", _iotwirelessLoRaWAN); err != nil {
			log.Errorf("invalid --lo-ra-wan: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessName) > 0 {
		input.Name = aws.String(_iotwirelessName)
	}
	if len(_iotwirelessRedundancyPercent) > 0 {
		if err := assignInputField(input, "RedundancyPercent", _iotwirelessRedundancyPercent); err != nil {
			log.Errorf("invalid --redundancy-percent: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessTags) > 0 {
		if err := assignInputField(input, "Tags", _iotwirelessTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateFuotaTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a multicast group.
func iotwireless_CreateMulticastGroup(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.CreateMulticastGroupInput{
		// LoRaWAN: *types.LoRaWANMulticast, // Required
	}

	if len(_iotwirelessLoRaWAN) > 0 {
		if err := assignInputField(input, "LoRaWAN", _iotwirelessLoRaWAN); err != nil {
			log.Errorf("invalid --lo-ra-wan: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_iotwirelessClientRequestToken)
	}
	if len(_iotwirelessDescription) > 0 {
		input.Description = aws.String(_iotwirelessDescription)
	}
	if len(_iotwirelessName) > 0 {
		input.Name = aws.String(_iotwirelessName)
	}
	if len(_iotwirelessTags) > 0 {
		if err := assignInputField(input, "Tags", _iotwirelessTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateMulticastGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new network analyzer configuration.
func iotwireless_CreateNetworkAnalyzerConfiguration(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.CreateNetworkAnalyzerConfigurationInput{
		// Name: *string, // Required
	}

	if len(_iotwirelessName) > 0 {
		input.Name = aws.String(_iotwirelessName)
	}
	if len(_iotwirelessClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_iotwirelessClientRequestToken)
	}
	if len(_iotwirelessDescription) > 0 {
		input.Description = aws.String(_iotwirelessDescription)
	}
	if len(_iotwirelessMulticastGroups) > 0 {
		input.MulticastGroups = append([]string(nil), _iotwirelessMulticastGroups...)
	}
	if len(_iotwirelessTags) > 0 {
		if err := assignInputField(input, "Tags", _iotwirelessTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessTraceContent) > 0 {
		if err := assignInputField(input, "TraceContent", _iotwirelessTraceContent); err != nil {
			log.Errorf("invalid --trace-content: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessWirelessDevices) > 0 {
		input.WirelessDevices = append([]string(nil), _iotwirelessWirelessDevices...)
	}
	if len(_iotwirelessWirelessGateways) > 0 {
		input.WirelessGateways = append([]string(nil), _iotwirelessWirelessGateways...)
	}

	if resp, err := client.CreateNetworkAnalyzerConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new service profile.
func iotwireless_CreateServiceProfile(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.CreateServiceProfileInput{}

	if len(_iotwirelessClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_iotwirelessClientRequestToken)
	}
	if len(_iotwirelessLoRaWAN) > 0 {
		if err := assignInputField(input, "LoRaWAN", _iotwirelessLoRaWAN); err != nil {
			log.Errorf("invalid --lo-ra-wan: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessName) > 0 {
		input.Name = aws.String(_iotwirelessName)
	}
	if len(_iotwirelessTags) > 0 {
		if err := assignInputField(input, "Tags", _iotwirelessTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateServiceProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provisions a wireless device.
func iotwireless_CreateWirelessDevice(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.CreateWirelessDeviceInput{
		// DestinationName: *string, // Required
		// Type: types.WirelessDeviceType, // Required
	}

	if len(_iotwirelessDestinationName) > 0 {
		input.DestinationName = aws.String(_iotwirelessDestinationName)
	}
	if len(_iotwirelessType) > 0 {
		if err := assignInputField(input, "Type", _iotwirelessType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_iotwirelessClientRequestToken)
	}
	if len(_iotwirelessDescription) > 0 {
		input.Description = aws.String(_iotwirelessDescription)
	}
	if len(_iotwirelessLoRaWAN) > 0 {
		if err := assignInputField(input, "LoRaWAN", _iotwirelessLoRaWAN); err != nil {
			log.Errorf("invalid --lo-ra-wan: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessName) > 0 {
		input.Name = aws.String(_iotwirelessName)
	}
	if len(_iotwirelessPositioning) > 0 {
		if err := assignInputField(input, "Positioning", _iotwirelessPositioning); err != nil {
			log.Errorf("invalid --positioning: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessSidewalk) > 0 {
		if err := assignInputField(input, "Sidewalk", _iotwirelessSidewalk); err != nil {
			log.Errorf("invalid --sidewalk: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessTags) > 0 {
		if err := assignInputField(input, "Tags", _iotwirelessTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateWirelessDevice(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provisions a wireless gateway.
// When provisioning a wireless gateway, you might run into duplication errors for
// the following reasons.
//
// - If you specify a GatewayEui value that already exists.
//
// - If you used a ClientRequestToken with the same parameters within the last 10
// minutes.
//
// To avoid this error, make sure that you use unique identifiers and parameters
// for each request within the specified time period.
func iotwireless_CreateWirelessGateway(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.CreateWirelessGatewayInput{
		// LoRaWAN: *types.LoRaWANGateway, // Required
	}

	if len(_iotwirelessLoRaWAN) > 0 {
		if err := assignInputField(input, "LoRaWAN", _iotwirelessLoRaWAN); err != nil {
			log.Errorf("invalid --lo-ra-wan: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_iotwirelessClientRequestToken)
	}
	if len(_iotwirelessDescription) > 0 {
		input.Description = aws.String(_iotwirelessDescription)
	}
	if len(_iotwirelessName) > 0 {
		input.Name = aws.String(_iotwirelessName)
	}
	if len(_iotwirelessTags) > 0 {
		if err := assignInputField(input, "Tags", _iotwirelessTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateWirelessGateway(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a task for a wireless gateway.
func iotwireless_CreateWirelessGatewayTask(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.CreateWirelessGatewayTaskInput{
		// Id: *string, // Required
		// WirelessGatewayTaskDefinitionId: *string, // Required
	}

	if len(_iotwirelessId) > 0 {
		input.Id = aws.String(_iotwirelessId)
	}
	if len(_iotwirelessWirelessGatewayTaskDefinitionId) > 0 {
		input.WirelessGatewayTaskDefinitionId = aws.String(_iotwirelessWirelessGatewayTaskDefinitionId)
	}

	if resp, err := client.CreateWirelessGatewayTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a gateway task definition.
func iotwireless_CreateWirelessGatewayTaskDefinition(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.CreateWirelessGatewayTaskDefinitionInput{
		// AutoCreateTasks: bool, // Required
	}

	if len(_iotwirelessAutoCreateTasks) > 0 {
		if err := assignInputField(input, "AutoCreateTasks", _iotwirelessAutoCreateTasks); err != nil {
			log.Errorf("invalid --auto-create-tasks: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_iotwirelessClientRequestToken)
	}
	if len(_iotwirelessName) > 0 {
		input.Name = aws.String(_iotwirelessName)
	}
	if len(_iotwirelessTags) > 0 {
		if err := assignInputField(input, "Tags", _iotwirelessTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessUpdate) > 0 {
		if err := assignInputField(input, "Update", _iotwirelessUpdate); err != nil {
			log.Errorf("invalid --update: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateWirelessGatewayTaskDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a destination.
func iotwireless_DeleteDestination(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.DeleteDestinationInput{
		// Name: *string, // Required
	}

	if len(_iotwirelessName) > 0 {
		input.Name = aws.String(_iotwirelessName)
	}

	if resp, err := client.DeleteDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a device profile.
func iotwireless_DeleteDeviceProfile(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.DeleteDeviceProfileInput{
		// Id: *string, // Required
	}

	if len(_iotwirelessId) > 0 {
		input.Id = aws.String(_iotwirelessId)
	}

	if resp, err := client.DeleteDeviceProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a FUOTA task.
func iotwireless_DeleteFuotaTask(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.DeleteFuotaTaskInput{
		// Id: *string, // Required
	}

	if len(_iotwirelessId) > 0 {
		input.Id = aws.String(_iotwirelessId)
	}

	if resp, err := client.DeleteFuotaTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a multicast group if it is not in use by a FUOTA task.
func iotwireless_DeleteMulticastGroup(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.DeleteMulticastGroupInput{
		// Id: *string, // Required
	}

	if len(_iotwirelessId) > 0 {
		input.Id = aws.String(_iotwirelessId)
	}

	if resp, err := client.DeleteMulticastGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a network analyzer configuration.
func iotwireless_DeleteNetworkAnalyzerConfiguration(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.DeleteNetworkAnalyzerConfigurationInput{
		// ConfigurationName: *string, // Required
	}

	if len(_iotwirelessConfigurationName) > 0 {
		input.ConfigurationName = aws.String(_iotwirelessConfigurationName)
	}

	if resp, err := client.DeleteNetworkAnalyzerConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Remove queued messages from the downlink queue.
func iotwireless_DeleteQueuedMessages(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.DeleteQueuedMessagesInput{
		// Id: *string, // Required
		// MessageId: *string, // Required
	}

	if len(_iotwirelessId) > 0 {
		input.Id = aws.String(_iotwirelessId)
	}
	if len(_iotwirelessMessageId) > 0 {
		input.MessageId = aws.String(_iotwirelessMessageId)
	}
	if len(_iotwirelessWirelessDeviceType) > 0 {
		if err := assignInputField(input, "WirelessDeviceType", _iotwirelessWirelessDeviceType); err != nil {
			log.Errorf("invalid --wireless-device-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteQueuedMessages(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a service profile.
func iotwireless_DeleteServiceProfile(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.DeleteServiceProfileInput{
		// Id: *string, // Required
	}

	if len(_iotwirelessId) > 0 {
		input.Id = aws.String(_iotwirelessId)
	}

	if resp, err := client.DeleteServiceProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a wireless device.
func iotwireless_DeleteWirelessDevice(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.DeleteWirelessDeviceInput{
		// Id: *string, // Required
	}

	if len(_iotwirelessId) > 0 {
		input.Id = aws.String(_iotwirelessId)
	}

	if resp, err := client.DeleteWirelessDevice(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete an import task.
func iotwireless_DeleteWirelessDeviceImportTask(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.DeleteWirelessDeviceImportTaskInput{
		// Id: *string, // Required
	}

	if len(_iotwirelessId) > 0 {
		input.Id = aws.String(_iotwirelessId)
	}

	if resp, err := client.DeleteWirelessDeviceImportTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a wireless gateway.
// When deleting a wireless gateway, you might run into duplication errors for the
// following reasons.
//
// - If you specify a GatewayEui value that already exists.
//
// - If you used a ClientRequestToken with the same parameters within the last 10
// minutes.
//
// To avoid this error, make sure that you use unique identifiers and parameters
// for each request within the specified time period.
func iotwireless_DeleteWirelessGateway(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.DeleteWirelessGatewayInput{
		// Id: *string, // Required
	}

	if len(_iotwirelessId) > 0 {
		input.Id = aws.String(_iotwirelessId)
	}

	if resp, err := client.DeleteWirelessGateway(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a wireless gateway task.
func iotwireless_DeleteWirelessGatewayTask(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.DeleteWirelessGatewayTaskInput{
		// Id: *string, // Required
	}

	if len(_iotwirelessId) > 0 {
		input.Id = aws.String(_iotwirelessId)
	}

	if resp, err := client.DeleteWirelessGatewayTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a wireless gateway task definition. Deleting this task definition does
// not affect tasks that are currently in progress.
func iotwireless_DeleteWirelessGatewayTaskDefinition(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.DeleteWirelessGatewayTaskDefinitionInput{
		// Id: *string, // Required
	}

	if len(_iotwirelessId) > 0 {
		input.Id = aws.String(_iotwirelessId)
	}

	if resp, err := client.DeleteWirelessGatewayTaskDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deregister a wireless device from AWS IoT Wireless.
func iotwireless_DeregisterWirelessDevice(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.DeregisterWirelessDeviceInput{
		// Identifier: *string, // Required
	}

	if len(_iotwirelessIdentifier) > 0 {
		input.Identifier = aws.String(_iotwirelessIdentifier)
	}
	if len(_iotwirelessWirelessDeviceType) > 0 {
		if err := assignInputField(input, "WirelessDeviceType", _iotwirelessWirelessDeviceType); err != nil {
			log.Errorf("invalid --wireless-device-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeregisterWirelessDevice(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates your AWS account from a partner account. If PartnerAccountId and
// PartnerType are null , disassociates your AWS account from all partner accounts.
func iotwireless_DisassociateAwsAccountFromPartnerAccount(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.DisassociateAwsAccountFromPartnerAccountInput{
		// PartnerAccountId: *string, // Required
		// PartnerType: types.PartnerType, // Required
	}

	if len(_iotwirelessPartnerAccountId) > 0 {
		input.PartnerAccountId = aws.String(_iotwirelessPartnerAccountId)
	}
	if len(_iotwirelessPartnerType) > 0 {
		if err := assignInputField(input, "PartnerType", _iotwirelessPartnerType); err != nil {
			log.Errorf("invalid --partner-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.DisassociateAwsAccountFromPartnerAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a multicast group from a FUOTA task.
func iotwireless_DisassociateMulticastGroupFromFuotaTask(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.DisassociateMulticastGroupFromFuotaTaskInput{
		// Id: *string, // Required
		// MulticastGroupId: *string, // Required
	}

	if len(_iotwirelessId) > 0 {
		input.Id = aws.String(_iotwirelessId)
	}
	if len(_iotwirelessMulticastGroupId) > 0 {
		input.MulticastGroupId = aws.String(_iotwirelessMulticastGroupId)
	}

	if resp, err := client.DisassociateMulticastGroupFromFuotaTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a wireless device from a FUOTA task.
func iotwireless_DisassociateWirelessDeviceFromFuotaTask(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.DisassociateWirelessDeviceFromFuotaTaskInput{
		// Id: *string, // Required
		// WirelessDeviceId: *string, // Required
	}

	if len(_iotwirelessId) > 0 {
		input.Id = aws.String(_iotwirelessId)
	}
	if len(_iotwirelessWirelessDeviceId) > 0 {
		input.WirelessDeviceId = aws.String(_iotwirelessWirelessDeviceId)
	}

	if resp, err := client.DisassociateWirelessDeviceFromFuotaTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a wireless device from a multicast group.
func iotwireless_DisassociateWirelessDeviceFromMulticastGroup(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.DisassociateWirelessDeviceFromMulticastGroupInput{
		// Id: *string, // Required
		// WirelessDeviceId: *string, // Required
	}

	if len(_iotwirelessId) > 0 {
		input.Id = aws.String(_iotwirelessId)
	}
	if len(_iotwirelessWirelessDeviceId) > 0 {
		input.WirelessDeviceId = aws.String(_iotwirelessWirelessDeviceId)
	}

	if resp, err := client.DisassociateWirelessDeviceFromMulticastGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a wireless device from its currently associated thing.
func iotwireless_DisassociateWirelessDeviceFromThing(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.DisassociateWirelessDeviceFromThingInput{
		// Id: *string, // Required
	}

	if len(_iotwirelessId) > 0 {
		input.Id = aws.String(_iotwirelessId)
	}

	if resp, err := client.DisassociateWirelessDeviceFromThing(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a wireless gateway from its currently associated certificate.
func iotwireless_DisassociateWirelessGatewayFromCertificate(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.DisassociateWirelessGatewayFromCertificateInput{
		// Id: *string, // Required
	}

	if len(_iotwirelessId) > 0 {
		input.Id = aws.String(_iotwirelessId)
	}

	if resp, err := client.DisassociateWirelessGatewayFromCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a wireless gateway from its currently associated thing.
func iotwireless_DisassociateWirelessGatewayFromThing(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.DisassociateWirelessGatewayFromThingInput{
		// Id: *string, // Required
	}

	if len(_iotwirelessId) > 0 {
		input.Id = aws.String(_iotwirelessId)
	}

	if resp, err := client.DisassociateWirelessGatewayFromThing(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a destination.
func iotwireless_GetDestination(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.GetDestinationInput{
		// Name: *string, // Required
	}

	if len(_iotwirelessName) > 0 {
		input.Name = aws.String(_iotwirelessName)
	}

	if resp, err := client.GetDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a device profile.
func iotwireless_GetDeviceProfile(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.GetDeviceProfileInput{
		// Id: *string, // Required
	}

	if len(_iotwirelessId) > 0 {
		input.Id = aws.String(_iotwirelessId)
	}

	if resp, err := client.GetDeviceProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the event configuration based on resource types.
func iotwireless_GetEventConfigurationByResourceTypes(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.GetEventConfigurationByResourceTypesInput{}

	if resp, err := client.GetEventConfigurationByResourceTypes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a FUOTA task.
func iotwireless_GetFuotaTask(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.GetFuotaTaskInput{
		// Id: *string, // Required
	}

	if len(_iotwirelessId) > 0 {
		input.Id = aws.String(_iotwirelessId)
	}

	if resp, err := client.GetFuotaTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns current default log levels or log levels by resource types. Based on
// the resource type, log levels can be returned for wireless device, wireless
// gateway, or FUOTA task log options.
func iotwireless_GetLogLevelsByResourceTypes(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.GetLogLevelsByResourceTypesInput{}

	if resp, err := client.GetLogLevelsByResourceTypes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the metric configuration status for this AWS account.
func iotwireless_GetMetricConfiguration(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.GetMetricConfigurationInput{}

	if resp, err := client.GetMetricConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the summary metrics for this AWS account.
func iotwireless_GetMetrics(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.GetMetricsInput{}

	if len(_iotwirelessSummaryMetricQueries) > 0 {
		if err := assignInputField(input, "SummaryMetricQueries", _iotwirelessSummaryMetricQueries); err != nil {
			log.Errorf("invalid --summary-metric-queries: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetMetrics(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a multicast group.
func iotwireless_GetMulticastGroup(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.GetMulticastGroupInput{
		// Id: *string, // Required
	}

	if len(_iotwirelessId) > 0 {
		input.Id = aws.String(_iotwirelessId)
	}

	if resp, err := client.GetMulticastGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a multicast group session.
func iotwireless_GetMulticastGroupSession(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.GetMulticastGroupSessionInput{
		// Id: *string, // Required
	}

	if len(_iotwirelessId) > 0 {
		input.Id = aws.String(_iotwirelessId)
	}

	if resp, err := client.GetMulticastGroupSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get network analyzer configuration.
func iotwireless_GetNetworkAnalyzerConfiguration(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.GetNetworkAnalyzerConfigurationInput{
		// ConfigurationName: *string, // Required
	}

	if len(_iotwirelessConfigurationName) > 0 {
		input.ConfigurationName = aws.String(_iotwirelessConfigurationName)
	}

	if resp, err := client.GetNetworkAnalyzerConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a partner account. If PartnerAccountId and PartnerType
// are null , returns all partner accounts.
func iotwireless_GetPartnerAccount(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.GetPartnerAccountInput{
		// PartnerAccountId: *string, // Required
		// PartnerType: types.PartnerType, // Required
	}

	if len(_iotwirelessPartnerAccountId) > 0 {
		input.PartnerAccountId = aws.String(_iotwirelessPartnerAccountId)
	}
	if len(_iotwirelessPartnerType) > 0 {
		if err := assignInputField(input, "PartnerType", _iotwirelessPartnerType); err != nil {
			log.Errorf("invalid --partner-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetPartnerAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the position information for a given resource.
// This action is no longer supported. Calls to retrieve the position information
// should use the [GetResourcePosition]API operation instead.
//
// Deprecated: This operation is no longer supported.
//
// [GetResourcePosition]: https://docs.aws.amazon.com/iot-wireless/latest/apireference/API_GetResourcePosition.html
func iotwireless_GetPosition(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.GetPositionInput{
		// ResourceIdentifier: *string, // Required
		// ResourceType: types.PositionResourceType, // Required
	}

	if len(_iotwirelessResourceIdentifier) > 0 {
		input.ResourceIdentifier = aws.String(_iotwirelessResourceIdentifier)
	}
	if len(_iotwirelessResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _iotwirelessResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetPosition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get position configuration for a given resource.
// This action is no longer supported. Calls to retrieve the position
// configuration should use the [GetResourcePosition]API operation instead.
//
// Deprecated: This operation is no longer supported.
//
// [GetResourcePosition]: https://docs.aws.amazon.com/iot-wireless/latest/apireference/API_GetResourcePosition.html
func iotwireless_GetPositionConfiguration(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.GetPositionConfigurationInput{
		// ResourceIdentifier: *string, // Required
		// ResourceType: types.PositionResourceType, // Required
	}

	if len(_iotwirelessResourceIdentifier) > 0 {
		input.ResourceIdentifier = aws.String(_iotwirelessResourceIdentifier)
	}
	if len(_iotwirelessResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _iotwirelessResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetPositionConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get estimated position information as a payload in GeoJSON format. The payload
// measurement data is resolved using solvers that are provided by third-party
// vendors.
func iotwireless_GetPositionEstimate(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.GetPositionEstimateInput{}

	if len(_iotwirelessCellTowers) > 0 {
		if err := assignInputField(input, "CellTowers", _iotwirelessCellTowers); err != nil {
			log.Errorf("invalid --cell-towers: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessGnss) > 0 {
		if err := assignInputField(input, "Gnss", _iotwirelessGnss); err != nil {
			log.Errorf("invalid --gnss: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessIp) > 0 {
		if err := assignInputField(input, "Ip", _iotwirelessIp); err != nil {
			log.Errorf("invalid --ip: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessTimestamp) > 0 {
		if err := assignInputField(input, "Timestamp", _iotwirelessTimestamp); err != nil {
			log.Errorf("invalid --timestamp: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessWiFiAccessPoints) > 0 {
		if err := assignInputField(input, "WiFiAccessPoints", _iotwirelessWiFiAccessPoints); err != nil {
			log.Errorf("invalid --wi-fi-access-points: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetPositionEstimate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the event configuration for a particular resource identifier.
func iotwireless_GetResourceEventConfiguration(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.GetResourceEventConfigurationInput{
		// Identifier: *string, // Required
		// IdentifierType: types.IdentifierType, // Required
	}

	if len(_iotwirelessIdentifier) > 0 {
		input.Identifier = aws.String(_iotwirelessIdentifier)
	}
	if len(_iotwirelessIdentifierType) > 0 {
		if err := assignInputField(input, "IdentifierType", _iotwirelessIdentifierType); err != nil {
			log.Errorf("invalid --identifier-type: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessPartnerType) > 0 {
		if err := assignInputField(input, "PartnerType", _iotwirelessPartnerType); err != nil {
			log.Errorf("invalid --partner-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetResourceEventConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Fetches the log-level override, if any, for a given resource ID and resource
// type..
func iotwireless_GetResourceLogLevel(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.GetResourceLogLevelInput{
		// ResourceIdentifier: *string, // Required
		// ResourceType: *string, // Required
	}

	if len(_iotwirelessResourceIdentifier) > 0 {
		input.ResourceIdentifier = aws.String(_iotwirelessResourceIdentifier)
	}
	if len(_iotwirelessResourceType) > 0 {
		input.ResourceType = aws.String(_iotwirelessResourceType)
	}

	if resp, err := client.GetResourceLogLevel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the position information for a given wireless device or a wireless gateway
// resource. The position information uses the [World Geodetic System (WGS84)].
//
// [World Geodetic System (WGS84)]: https://gisgeography.com/wgs84-world-geodetic-system/
func iotwireless_GetResourcePosition(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.GetResourcePositionInput{
		// ResourceIdentifier: *string, // Required
		// ResourceType: types.PositionResourceType, // Required
	}

	if len(_iotwirelessResourceIdentifier) > 0 {
		input.ResourceIdentifier = aws.String(_iotwirelessResourceIdentifier)
	}
	if len(_iotwirelessResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _iotwirelessResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetResourcePosition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the account-specific endpoint for Configuration and Update Server (CUPS)
// protocol or LoRaWAN Network Server (LNS) connections.
func iotwireless_GetServiceEndpoint(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.GetServiceEndpointInput{}

	if len(_iotwirelessServiceType) > 0 {
		if err := assignInputField(input, "ServiceType", _iotwirelessServiceType); err != nil {
			log.Errorf("invalid --service-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetServiceEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a service profile.
func iotwireless_GetServiceProfile(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.GetServiceProfileInput{
		// Id: *string, // Required
	}

	if len(_iotwirelessId) > 0 {
		input.Id = aws.String(_iotwirelessId)
	}

	if resp, err := client.GetServiceProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a wireless device.
func iotwireless_GetWirelessDevice(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.GetWirelessDeviceInput{
		// Identifier: *string, // Required
		// IdentifierType: types.WirelessDeviceIdType, // Required
	}

	if len(_iotwirelessIdentifier) > 0 {
		input.Identifier = aws.String(_iotwirelessIdentifier)
	}
	if len(_iotwirelessIdentifierType) > 0 {
		if err := assignInputField(input, "IdentifierType", _iotwirelessIdentifierType); err != nil {
			log.Errorf("invalid --identifier-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetWirelessDevice(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get information about an import task and count of device onboarding summary
// information for the import task.
func iotwireless_GetWirelessDeviceImportTask(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.GetWirelessDeviceImportTaskInput{
		// Id: *string, // Required
	}

	if len(_iotwirelessId) > 0 {
		input.Id = aws.String(_iotwirelessId)
	}

	if resp, err := client.GetWirelessDeviceImportTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets operating information about a wireless device.
func iotwireless_GetWirelessDeviceStatistics(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.GetWirelessDeviceStatisticsInput{
		// WirelessDeviceId: *string, // Required
	}

	if len(_iotwirelessWirelessDeviceId) > 0 {
		input.WirelessDeviceId = aws.String(_iotwirelessWirelessDeviceId)
	}

	if resp, err := client.GetWirelessDeviceStatistics(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a wireless gateway.
func iotwireless_GetWirelessGateway(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.GetWirelessGatewayInput{
		// Identifier: *string, // Required
		// IdentifierType: types.WirelessGatewayIdType, // Required
	}

	if len(_iotwirelessIdentifier) > 0 {
		input.Identifier = aws.String(_iotwirelessIdentifier)
	}
	if len(_iotwirelessIdentifierType) > 0 {
		if err := assignInputField(input, "IdentifierType", _iotwirelessIdentifierType); err != nil {
			log.Errorf("invalid --identifier-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetWirelessGateway(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the ID of the certificate that is currently associated with a wireless
// gateway.
func iotwireless_GetWirelessGatewayCertificate(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.GetWirelessGatewayCertificateInput{
		// Id: *string, // Required
	}

	if len(_iotwirelessId) > 0 {
		input.Id = aws.String(_iotwirelessId)
	}

	if resp, err := client.GetWirelessGatewayCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the firmware version and other information about a wireless gateway.
func iotwireless_GetWirelessGatewayFirmwareInformation(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.GetWirelessGatewayFirmwareInformationInput{
		// Id: *string, // Required
	}

	if len(_iotwirelessId) > 0 {
		input.Id = aws.String(_iotwirelessId)
	}

	if resp, err := client.GetWirelessGatewayFirmwareInformation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets operating information about a wireless gateway.
func iotwireless_GetWirelessGatewayStatistics(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.GetWirelessGatewayStatisticsInput{
		// WirelessGatewayId: *string, // Required
	}

	if len(_iotwirelessWirelessGatewayId) > 0 {
		input.WirelessGatewayId = aws.String(_iotwirelessWirelessGatewayId)
	}

	if resp, err := client.GetWirelessGatewayStatistics(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a wireless gateway task.
func iotwireless_GetWirelessGatewayTask(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.GetWirelessGatewayTaskInput{
		// Id: *string, // Required
	}

	if len(_iotwirelessId) > 0 {
		input.Id = aws.String(_iotwirelessId)
	}

	if resp, err := client.GetWirelessGatewayTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a wireless gateway task definition.
func iotwireless_GetWirelessGatewayTaskDefinition(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.GetWirelessGatewayTaskDefinitionInput{
		// Id: *string, // Required
	}

	if len(_iotwirelessId) > 0 {
		input.Id = aws.String(_iotwirelessId)
	}

	if resp, err := client.GetWirelessGatewayTaskDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the destinations registered to your AWS account.
func iotwireless_ListDestinations(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.ListDestinationsInput{}

	if len(_iotwirelessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotwirelessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessNextToken) > 0 {
		input.NextToken = aws.String(_iotwirelessNextToken)
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

	var results []*iotwireless.ListDestinationsOutput
	p := iotwireless.NewListDestinationsPaginator(client, input)
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

// Lists the device profiles registered to your AWS account.
func iotwireless_ListDeviceProfiles(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.ListDeviceProfilesInput{}

	if len(_iotwirelessDeviceProfileType) > 0 {
		if err := assignInputField(input, "DeviceProfileType", _iotwirelessDeviceProfileType); err != nil {
			log.Errorf("invalid --device-profile-type: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotwirelessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessNextToken) > 0 {
		input.NextToken = aws.String(_iotwirelessNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDeviceProfiles(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotwireless.ListDeviceProfilesOutput
	p := iotwireless.NewListDeviceProfilesPaginator(client, input)
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

// List the Sidewalk devices in an import task and their onboarding status.
func iotwireless_ListDevicesForWirelessDeviceImportTask(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.ListDevicesForWirelessDeviceImportTaskInput{
		// Id: *string, // Required
	}

	if len(_iotwirelessId) > 0 {
		input.Id = aws.String(_iotwirelessId)
	}
	if len(_iotwirelessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotwirelessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessNextToken) > 0 {
		input.NextToken = aws.String(_iotwirelessNextToken)
	}
	if len(_iotwirelessStatus) > 0 {
		if err := assignInputField(input, "Status", _iotwirelessStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListDevicesForWirelessDeviceImportTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List event configurations where at least one event topic has been enabled.
func iotwireless_ListEventConfigurations(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.ListEventConfigurationsInput{
		// ResourceType: types.EventNotificationResourceType, // Required
	}

	if len(_iotwirelessResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _iotwirelessResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotwirelessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessNextToken) > 0 {
		input.NextToken = aws.String(_iotwirelessNextToken)
	}

	if resp, err := client.ListEventConfigurations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the FUOTA tasks registered to your AWS account.
func iotwireless_ListFuotaTasks(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.ListFuotaTasksInput{}

	if len(_iotwirelessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotwirelessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessNextToken) > 0 {
		input.NextToken = aws.String(_iotwirelessNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListFuotaTasks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotwireless.ListFuotaTasksOutput
	p := iotwireless.NewListFuotaTasksPaginator(client, input)
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

// Lists the multicast groups registered to your AWS account.
func iotwireless_ListMulticastGroups(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.ListMulticastGroupsInput{}

	if len(_iotwirelessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotwirelessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessNextToken) > 0 {
		input.NextToken = aws.String(_iotwirelessNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListMulticastGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotwireless.ListMulticastGroupsOutput
	p := iotwireless.NewListMulticastGroupsPaginator(client, input)
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

// List all multicast groups associated with a FUOTA task.
func iotwireless_ListMulticastGroupsByFuotaTask(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.ListMulticastGroupsByFuotaTaskInput{
		// Id: *string, // Required
	}

	if len(_iotwirelessId) > 0 {
		input.Id = aws.String(_iotwirelessId)
	}
	if len(_iotwirelessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotwirelessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessNextToken) > 0 {
		input.NextToken = aws.String(_iotwirelessNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListMulticastGroupsByFuotaTask(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotwireless.ListMulticastGroupsByFuotaTaskOutput
	p := iotwireless.NewListMulticastGroupsByFuotaTaskPaginator(client, input)
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

// Lists the network analyzer configurations.
func iotwireless_ListNetworkAnalyzerConfigurations(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.ListNetworkAnalyzerConfigurationsInput{}

	if len(_iotwirelessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotwirelessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessNextToken) > 0 {
		input.NextToken = aws.String(_iotwirelessNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListNetworkAnalyzerConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotwireless.ListNetworkAnalyzerConfigurationsOutput
	p := iotwireless.NewListNetworkAnalyzerConfigurationsPaginator(client, input)
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

// Lists the partner accounts associated with your AWS account.
func iotwireless_ListPartnerAccounts(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.ListPartnerAccountsInput{}

	if len(_iotwirelessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotwirelessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessNextToken) > 0 {
		input.NextToken = aws.String(_iotwirelessNextToken)
	}

	if resp, err := client.ListPartnerAccounts(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List position configurations for a given resource, such as positioning solvers.
// This action is no longer supported. Calls to retrieve position information
// should use the [GetResourcePosition]API operation instead.
//
// Deprecated: This operation is no longer supported.
//
// [GetResourcePosition]: https://docs.aws.amazon.com/iot-wireless/latest/apireference/API_GetResourcePosition.html
func iotwireless_ListPositionConfigurations(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.ListPositionConfigurationsInput{}

	if len(_iotwirelessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotwirelessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessNextToken) > 0 {
		input.NextToken = aws.String(_iotwirelessNextToken)
	}
	if len(_iotwirelessResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _iotwirelessResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListPositionConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotwireless.ListPositionConfigurationsOutput
	p := iotwireless.NewListPositionConfigurationsPaginator(client, input)
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

// List queued messages in the downlink queue.
func iotwireless_ListQueuedMessages(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.ListQueuedMessagesInput{
		// Id: *string, // Required
	}

	if len(_iotwirelessId) > 0 {
		input.Id = aws.String(_iotwirelessId)
	}
	if len(_iotwirelessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotwirelessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessNextToken) > 0 {
		input.NextToken = aws.String(_iotwirelessNextToken)
	}
	if len(_iotwirelessWirelessDeviceType) > 0 {
		if err := assignInputField(input, "WirelessDeviceType", _iotwirelessWirelessDeviceType); err != nil {
			log.Errorf("invalid --wireless-device-type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListQueuedMessages(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotwireless.ListQueuedMessagesOutput
	p := iotwireless.NewListQueuedMessagesPaginator(client, input)
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

// Lists the service profiles registered to your AWS account.
func iotwireless_ListServiceProfiles(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.ListServiceProfilesInput{}

	if len(_iotwirelessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotwirelessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessNextToken) > 0 {
		input.NextToken = aws.String(_iotwirelessNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListServiceProfiles(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotwireless.ListServiceProfilesOutput
	p := iotwireless.NewListServiceProfilesPaginator(client, input)
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

// Lists the tags (metadata) you have assigned to the resource.
func iotwireless_ListTagsForResource(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_iotwirelessResourceArn) > 0 {
		input.ResourceArn = aws.String(_iotwirelessResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List of import tasks and summary information of onboarding status of devices in
// each import task.
func iotwireless_ListWirelessDeviceImportTasks(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.ListWirelessDeviceImportTasksInput{}

	if len(_iotwirelessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotwirelessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessNextToken) > 0 {
		input.NextToken = aws.String(_iotwirelessNextToken)
	}

	if resp, err := client.ListWirelessDeviceImportTasks(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the wireless devices registered to your AWS account.
func iotwireless_ListWirelessDevices(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.ListWirelessDevicesInput{}

	if len(_iotwirelessDestinationName) > 0 {
		input.DestinationName = aws.String(_iotwirelessDestinationName)
	}
	if len(_iotwirelessDeviceProfileId) > 0 {
		input.DeviceProfileId = aws.String(_iotwirelessDeviceProfileId)
	}
	if len(_iotwirelessFuotaTaskId) > 0 {
		input.FuotaTaskId = aws.String(_iotwirelessFuotaTaskId)
	}
	if len(_iotwirelessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotwirelessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessMulticastGroupId) > 0 {
		input.MulticastGroupId = aws.String(_iotwirelessMulticastGroupId)
	}
	if len(_iotwirelessNextToken) > 0 {
		input.NextToken = aws.String(_iotwirelessNextToken)
	}
	if len(_iotwirelessServiceProfileId) > 0 {
		input.ServiceProfileId = aws.String(_iotwirelessServiceProfileId)
	}
	if len(_iotwirelessWirelessDeviceType) > 0 {
		if err := assignInputField(input, "WirelessDeviceType", _iotwirelessWirelessDeviceType); err != nil {
			log.Errorf("invalid --wireless-device-type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListWirelessDevices(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotwireless.ListWirelessDevicesOutput
	p := iotwireless.NewListWirelessDevicesPaginator(client, input)
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

// List the wireless gateway tasks definitions registered to your AWS account.
func iotwireless_ListWirelessGatewayTaskDefinitions(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.ListWirelessGatewayTaskDefinitionsInput{}

	if len(_iotwirelessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotwirelessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessNextToken) > 0 {
		input.NextToken = aws.String(_iotwirelessNextToken)
	}
	if len(_iotwirelessTaskDefinitionType) > 0 {
		if err := assignInputField(input, "TaskDefinitionType", _iotwirelessTaskDefinitionType); err != nil {
			log.Errorf("invalid --task-definition-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListWirelessGatewayTaskDefinitions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the wireless gateways registered to your AWS account.
func iotwireless_ListWirelessGateways(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.ListWirelessGatewaysInput{}

	if len(_iotwirelessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotwirelessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessNextToken) > 0 {
		input.NextToken = aws.String(_iotwirelessNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListWirelessGateways(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotwireless.ListWirelessGatewaysOutput
	p := iotwireless.NewListWirelessGatewaysPaginator(client, input)
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

// Put position configuration for a given resource.
// This action is no longer supported. Calls to update the position configuration
// should use the [UpdateResourcePosition]API operation instead.
//
// Deprecated: This operation is no longer supported.
//
// [UpdateResourcePosition]: https://docs.aws.amazon.com/iot-wireless/latest/apireference/API_UpdateResourcePosition.html
func iotwireless_PutPositionConfiguration(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.PutPositionConfigurationInput{
		// ResourceIdentifier: *string, // Required
		// ResourceType: types.PositionResourceType, // Required
	}

	if len(_iotwirelessResourceIdentifier) > 0 {
		input.ResourceIdentifier = aws.String(_iotwirelessResourceIdentifier)
	}
	if len(_iotwirelessResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _iotwirelessResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessDestination) > 0 {
		input.Destination = aws.String(_iotwirelessDestination)
	}
	if len(_iotwirelessSolvers) > 0 {
		if err := assignInputField(input, "Solvers", _iotwirelessSolvers); err != nil {
			log.Errorf("invalid --solvers: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutPositionConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the log-level override for a resource ID and resource type. A limit of 200
// log level override can be set per account.
func iotwireless_PutResourceLogLevel(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.PutResourceLogLevelInput{
		// LogLevel: types.LogLevel, // Required
		// ResourceIdentifier: *string, // Required
		// ResourceType: *string, // Required
	}

	if len(_iotwirelessLogLevel) > 0 {
		if err := assignInputField(input, "LogLevel", _iotwirelessLogLevel); err != nil {
			log.Errorf("invalid --log-level: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessResourceIdentifier) > 0 {
		input.ResourceIdentifier = aws.String(_iotwirelessResourceIdentifier)
	}
	if len(_iotwirelessResourceType) > 0 {
		input.ResourceType = aws.String(_iotwirelessResourceType)
	}

	if resp, err := client.PutResourceLogLevel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the log-level overrides for all resources; wireless devices, wireless
// gateways, and FUOTA tasks.
func iotwireless_ResetAllResourceLogLevels(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.ResetAllResourceLogLevelsInput{}

	if resp, err := client.ResetAllResourceLogLevels(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the log-level override, if any, for a specific resource ID and resource
// type. It can be used for a wireless device, a wireless gateway, or a FUOTA task.
func iotwireless_ResetResourceLogLevel(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.ResetResourceLogLevelInput{
		// ResourceIdentifier: *string, // Required
		// ResourceType: *string, // Required
	}

	if len(_iotwirelessResourceIdentifier) > 0 {
		input.ResourceIdentifier = aws.String(_iotwirelessResourceIdentifier)
	}
	if len(_iotwirelessResourceType) > 0 {
		input.ResourceType = aws.String(_iotwirelessResourceType)
	}

	if resp, err := client.ResetResourceLogLevel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sends the specified data to a multicast group.
func iotwireless_SendDataToMulticastGroup(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.SendDataToMulticastGroupInput{
		// Id: *string, // Required
		// PayloadData: *string, // Required
		// WirelessMetadata: *types.MulticastWirelessMetadata, // Required
	}

	if len(_iotwirelessId) > 0 {
		input.Id = aws.String(_iotwirelessId)
	}
	if len(_iotwirelessPayloadData) > 0 {
		input.PayloadData = aws.String(_iotwirelessPayloadData)
	}
	if len(_iotwirelessWirelessMetadata) > 0 {
		if err := assignInputField(input, "WirelessMetadata", _iotwirelessWirelessMetadata); err != nil {
			log.Errorf("invalid --wireless-metadata: %s", err.Error())
			return
		}
	}

	if resp, err := client.SendDataToMulticastGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sends a decrypted application data frame to a device.
func iotwireless_SendDataToWirelessDevice(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.SendDataToWirelessDeviceInput{
		// Id: *string, // Required
		// PayloadData: *string, // Required
		// TransmitMode: *int32, // Required
	}

	if len(_iotwirelessId) > 0 {
		input.Id = aws.String(_iotwirelessId)
	}
	if len(_iotwirelessPayloadData) > 0 {
		input.PayloadData = aws.String(_iotwirelessPayloadData)
	}
	if len(_iotwirelessTransmitMode) > 0 {
		if err := assignInputField(input, "TransmitMode", _iotwirelessTransmitMode); err != nil {
			log.Errorf("invalid --transmit-mode: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessWirelessMetadata) > 0 {
		if err := assignInputField(input, "WirelessMetadata", _iotwirelessWirelessMetadata); err != nil {
			log.Errorf("invalid --wireless-metadata: %s", err.Error())
			return
		}
	}

	if resp, err := client.SendDataToWirelessDevice(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a bulk association of all qualifying wireless devices with a multicast
// group.
func iotwireless_StartBulkAssociateWirelessDeviceWithMulticastGroup(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.StartBulkAssociateWirelessDeviceWithMulticastGroupInput{
		// Id: *string, // Required
	}

	if len(_iotwirelessId) > 0 {
		input.Id = aws.String(_iotwirelessId)
	}
	if len(_iotwirelessQueryString) > 0 {
		input.QueryString = aws.String(_iotwirelessQueryString)
	}
	if len(_iotwirelessTags) > 0 {
		if err := assignInputField(input, "Tags", _iotwirelessTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartBulkAssociateWirelessDeviceWithMulticastGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a bulk disassociatin of all qualifying wireless devices from a multicast
// group.
func iotwireless_StartBulkDisassociateWirelessDeviceFromMulticastGroup(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.StartBulkDisassociateWirelessDeviceFromMulticastGroupInput{
		// Id: *string, // Required
	}

	if len(_iotwirelessId) > 0 {
		input.Id = aws.String(_iotwirelessId)
	}
	if len(_iotwirelessQueryString) > 0 {
		input.QueryString = aws.String(_iotwirelessQueryString)
	}
	if len(_iotwirelessTags) > 0 {
		if err := assignInputField(input, "Tags", _iotwirelessTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartBulkDisassociateWirelessDeviceFromMulticastGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a FUOTA task.
func iotwireless_StartFuotaTask(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.StartFuotaTaskInput{
		// Id: *string, // Required
	}

	if len(_iotwirelessId) > 0 {
		input.Id = aws.String(_iotwirelessId)
	}
	if len(_iotwirelessLoRaWAN) > 0 {
		if err := assignInputField(input, "LoRaWAN", _iotwirelessLoRaWAN); err != nil {
			log.Errorf("invalid --lo-ra-wan: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartFuotaTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a multicast group session.
func iotwireless_StartMulticastGroupSession(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.StartMulticastGroupSessionInput{
		// Id: *string, // Required
		// LoRaWAN: *types.LoRaWANMulticastSession, // Required
	}

	if len(_iotwirelessId) > 0 {
		input.Id = aws.String(_iotwirelessId)
	}
	if len(_iotwirelessLoRaWAN) > 0 {
		if err := assignInputField(input, "LoRaWAN", _iotwirelessLoRaWAN); err != nil {
			log.Errorf("invalid --lo-ra-wan: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartMulticastGroupSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Start import task for a single wireless device.
func iotwireless_StartSingleWirelessDeviceImportTask(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.StartSingleWirelessDeviceImportTaskInput{
		// DestinationName: *string, // Required
		// Sidewalk: *types.SidewalkSingleStartImportInfo, // Required
	}

	if len(_iotwirelessDestinationName) > 0 {
		input.DestinationName = aws.String(_iotwirelessDestinationName)
	}
	if len(_iotwirelessSidewalk) > 0 {
		if err := assignInputField(input, "Sidewalk", _iotwirelessSidewalk); err != nil {
			log.Errorf("invalid --sidewalk: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_iotwirelessClientRequestToken)
	}
	if len(_iotwirelessDeviceName) > 0 {
		input.DeviceName = aws.String(_iotwirelessDeviceName)
	}
	if len(_iotwirelessPositioning) > 0 {
		if err := assignInputField(input, "Positioning", _iotwirelessPositioning); err != nil {
			log.Errorf("invalid --positioning: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessTags) > 0 {
		if err := assignInputField(input, "Tags", _iotwirelessTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartSingleWirelessDeviceImportTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Start import task for provisioning Sidewalk devices in bulk using an S3 CSV
// file.
func iotwireless_StartWirelessDeviceImportTask(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.StartWirelessDeviceImportTaskInput{
		// DestinationName: *string, // Required
		// Sidewalk: *types.SidewalkStartImportInfo, // Required
	}

	if len(_iotwirelessDestinationName) > 0 {
		input.DestinationName = aws.String(_iotwirelessDestinationName)
	}
	if len(_iotwirelessSidewalk) > 0 {
		if err := assignInputField(input, "Sidewalk", _iotwirelessSidewalk); err != nil {
			log.Errorf("invalid --sidewalk: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_iotwirelessClientRequestToken)
	}
	if len(_iotwirelessPositioning) > 0 {
		if err := assignInputField(input, "Positioning", _iotwirelessPositioning); err != nil {
			log.Errorf("invalid --positioning: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessTags) > 0 {
		if err := assignInputField(input, "Tags", _iotwirelessTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartWirelessDeviceImportTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a tag to a resource.
func iotwireless_TagResource(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_iotwirelessResourceArn) > 0 {
		input.ResourceArn = aws.String(_iotwirelessResourceArn)
	}
	if len(_iotwirelessTags) > 0 {
		if err := assignInputField(input, "Tags", _iotwirelessTags); err != nil {
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

// Simulates a provisioned device by sending an uplink data payload of Hello .
func iotwireless_TestWirelessDevice(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.TestWirelessDeviceInput{
		// Id: *string, // Required
	}

	if len(_iotwirelessId) > 0 {
		input.Id = aws.String(_iotwirelessId)
	}

	if resp, err := client.TestWirelessDevice(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes one or more tags from a resource.
func iotwireless_UntagResource(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_iotwirelessResourceArn) > 0 {
		input.ResourceArn = aws.String(_iotwirelessResourceArn)
	}
	if len(_iotwirelessTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _iotwirelessTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates properties of a destination.
func iotwireless_UpdateDestination(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.UpdateDestinationInput{
		// Name: *string, // Required
	}

	if len(_iotwirelessName) > 0 {
		input.Name = aws.String(_iotwirelessName)
	}
	if len(_iotwirelessDescription) > 0 {
		input.Description = aws.String(_iotwirelessDescription)
	}
	if len(_iotwirelessExpression) > 0 {
		input.Expression = aws.String(_iotwirelessExpression)
	}
	if len(_iotwirelessExpressionType) > 0 {
		if err := assignInputField(input, "ExpressionType", _iotwirelessExpressionType); err != nil {
			log.Errorf("invalid --expression-type: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessRoleArn) > 0 {
		input.RoleArn = aws.String(_iotwirelessRoleArn)
	}

	if resp, err := client.UpdateDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update the event configuration based on resource types.
func iotwireless_UpdateEventConfigurationByResourceTypes(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.UpdateEventConfigurationByResourceTypesInput{}

	if len(_iotwirelessConnectionStatus) > 0 {
		if err := assignInputField(input, "ConnectionStatus", _iotwirelessConnectionStatus); err != nil {
			log.Errorf("invalid --connection-status: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessDeviceRegistrationState) > 0 {
		if err := assignInputField(input, "DeviceRegistrationState", _iotwirelessDeviceRegistrationState); err != nil {
			log.Errorf("invalid --device-registration-state: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessJoin) > 0 {
		if err := assignInputField(input, "Join", _iotwirelessJoin); err != nil {
			log.Errorf("invalid --join: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessMessageDeliveryStatus) > 0 {
		if err := assignInputField(input, "MessageDeliveryStatus", _iotwirelessMessageDeliveryStatus); err != nil {
			log.Errorf("invalid --message-delivery-status: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessProximity) > 0 {
		if err := assignInputField(input, "Proximity", _iotwirelessProximity); err != nil {
			log.Errorf("invalid --proximity: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateEventConfigurationByResourceTypes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates properties of a FUOTA task.
func iotwireless_UpdateFuotaTask(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.UpdateFuotaTaskInput{
		// Id: *string, // Required
	}

	if len(_iotwirelessId) > 0 {
		input.Id = aws.String(_iotwirelessId)
	}
	if len(_iotwirelessDescription) > 0 {
		input.Description = aws.String(_iotwirelessDescription)
	}
	if len(_iotwirelessDescriptor) > 0 {
		input.Descriptor = aws.String(_iotwirelessDescriptor)
	}
	if len(_iotwirelessFirmwareUpdateImage) > 0 {
		input.FirmwareUpdateImage = aws.String(_iotwirelessFirmwareUpdateImage)
	}
	if len(_iotwirelessFirmwareUpdateRole) > 0 {
		input.FirmwareUpdateRole = aws.String(_iotwirelessFirmwareUpdateRole)
	}
	if len(_iotwirelessFragmentIntervalMS) > 0 {
		if err := assignInputField(input, "FragmentIntervalMS", _iotwirelessFragmentIntervalMS); err != nil {
			log.Errorf("invalid --fragment-interval-ms: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessFragmentSizeBytes) > 0 {
		if err := assignInputField(input, "FragmentSizeBytes", _iotwirelessFragmentSizeBytes); err != nil {
			log.Errorf("invalid --fragment-size-bytes: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessLoRaWAN) > 0 {
		if err := assignInputField(input, "LoRaWAN", _iotwirelessLoRaWAN); err != nil {
			log.Errorf("invalid --lo-ra-wan: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessName) > 0 {
		input.Name = aws.String(_iotwirelessName)
	}
	if len(_iotwirelessRedundancyPercent) > 0 {
		if err := assignInputField(input, "RedundancyPercent", _iotwirelessRedundancyPercent); err != nil {
			log.Errorf("invalid --redundancy-percent: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateFuotaTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Set default log level, or log levels by resource types. This can be for
// wireless device, wireless gateway, or FUOTA task log options, and is used to
// control the log messages that'll be displayed in CloudWatch.
func iotwireless_UpdateLogLevelsByResourceTypes(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.UpdateLogLevelsByResourceTypesInput{}

	if len(_iotwirelessDefaultLogLevel) > 0 {
		if err := assignInputField(input, "DefaultLogLevel", _iotwirelessDefaultLogLevel); err != nil {
			log.Errorf("invalid --default-log-level: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessFuotaTaskLogOptions) > 0 {
		if err := assignInputField(input, "FuotaTaskLogOptions", _iotwirelessFuotaTaskLogOptions); err != nil {
			log.Errorf("invalid --fuota-task-log-options: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessWirelessDeviceLogOptions) > 0 {
		if err := assignInputField(input, "WirelessDeviceLogOptions", _iotwirelessWirelessDeviceLogOptions); err != nil {
			log.Errorf("invalid --wireless-device-log-options: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessWirelessGatewayLogOptions) > 0 {
		if err := assignInputField(input, "WirelessGatewayLogOptions", _iotwirelessWirelessGatewayLogOptions); err != nil {
			log.Errorf("invalid --wireless-gateway-log-options: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateLogLevelsByResourceTypes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update the summary metric configuration.
func iotwireless_UpdateMetricConfiguration(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.UpdateMetricConfigurationInput{}

	if len(_iotwirelessSummaryMetric) > 0 {
		if err := assignInputField(input, "SummaryMetric", _iotwirelessSummaryMetric); err != nil {
			log.Errorf("invalid --summary-metric: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateMetricConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates properties of a multicast group session.
func iotwireless_UpdateMulticastGroup(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.UpdateMulticastGroupInput{
		// Id: *string, // Required
	}

	if len(_iotwirelessId) > 0 {
		input.Id = aws.String(_iotwirelessId)
	}
	if len(_iotwirelessDescription) > 0 {
		input.Description = aws.String(_iotwirelessDescription)
	}
	if len(_iotwirelessLoRaWAN) > 0 {
		if err := assignInputField(input, "LoRaWAN", _iotwirelessLoRaWAN); err != nil {
			log.Errorf("invalid --lo-ra-wan: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessName) > 0 {
		input.Name = aws.String(_iotwirelessName)
	}

	if resp, err := client.UpdateMulticastGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update network analyzer configuration.
func iotwireless_UpdateNetworkAnalyzerConfiguration(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.UpdateNetworkAnalyzerConfigurationInput{
		// ConfigurationName: *string, // Required
	}

	if len(_iotwirelessConfigurationName) > 0 {
		input.ConfigurationName = aws.String(_iotwirelessConfigurationName)
	}
	if len(_iotwirelessDescription) > 0 {
		input.Description = aws.String(_iotwirelessDescription)
	}
	if len(_iotwirelessMulticastGroupsToAdd) > 0 {
		input.MulticastGroupsToAdd = append([]string(nil), _iotwirelessMulticastGroupsToAdd...)
	}
	if len(_iotwirelessMulticastGroupsToRemove) > 0 {
		input.MulticastGroupsToRemove = append([]string(nil), _iotwirelessMulticastGroupsToRemove...)
	}
	if len(_iotwirelessTraceContent) > 0 {
		if err := assignInputField(input, "TraceContent", _iotwirelessTraceContent); err != nil {
			log.Errorf("invalid --trace-content: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessWirelessDevicesToAdd) > 0 {
		input.WirelessDevicesToAdd = append([]string(nil), _iotwirelessWirelessDevicesToAdd...)
	}
	if len(_iotwirelessWirelessDevicesToRemove) > 0 {
		input.WirelessDevicesToRemove = append([]string(nil), _iotwirelessWirelessDevicesToRemove...)
	}
	if len(_iotwirelessWirelessGatewaysToAdd) > 0 {
		input.WirelessGatewaysToAdd = append([]string(nil), _iotwirelessWirelessGatewaysToAdd...)
	}
	if len(_iotwirelessWirelessGatewaysToRemove) > 0 {
		input.WirelessGatewaysToRemove = append([]string(nil), _iotwirelessWirelessGatewaysToRemove...)
	}

	if resp, err := client.UpdateNetworkAnalyzerConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates properties of a partner account.
func iotwireless_UpdatePartnerAccount(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.UpdatePartnerAccountInput{
		// PartnerAccountId: *string, // Required
		// PartnerType: types.PartnerType, // Required
		// Sidewalk: *types.SidewalkUpdateAccount, // Required
	}

	if len(_iotwirelessPartnerAccountId) > 0 {
		input.PartnerAccountId = aws.String(_iotwirelessPartnerAccountId)
	}
	if len(_iotwirelessPartnerType) > 0 {
		if err := assignInputField(input, "PartnerType", _iotwirelessPartnerType); err != nil {
			log.Errorf("invalid --partner-type: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessSidewalk) > 0 {
		if err := assignInputField(input, "Sidewalk", _iotwirelessSidewalk); err != nil {
			log.Errorf("invalid --sidewalk: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdatePartnerAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update the position information of a resource.
// This action is no longer supported. Calls to update the position information
// should use the [UpdateResourcePosition]API operation instead.
//
// Deprecated: This operation is no longer supported.
//
// [UpdateResourcePosition]: https://docs.aws.amazon.com/iot-wireless/latest/apireference/API_UpdateResourcePosition.html
func iotwireless_UpdatePosition(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.UpdatePositionInput{
		// Position: []float32, // Required
		// ResourceIdentifier: *string, // Required
		// ResourceType: types.PositionResourceType, // Required
	}

	if len(_iotwirelessPosition) > 0 {
		if err := assignInputField(input, "Position", _iotwirelessPosition); err != nil {
			log.Errorf("invalid --position: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessResourceIdentifier) > 0 {
		input.ResourceIdentifier = aws.String(_iotwirelessResourceIdentifier)
	}
	if len(_iotwirelessResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _iotwirelessResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdatePosition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update the event configuration for a particular resource identifier.
func iotwireless_UpdateResourceEventConfiguration(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.UpdateResourceEventConfigurationInput{
		// Identifier: *string, // Required
		// IdentifierType: types.IdentifierType, // Required
	}

	if len(_iotwirelessIdentifier) > 0 {
		input.Identifier = aws.String(_iotwirelessIdentifier)
	}
	if len(_iotwirelessIdentifierType) > 0 {
		if err := assignInputField(input, "IdentifierType", _iotwirelessIdentifierType); err != nil {
			log.Errorf("invalid --identifier-type: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessConnectionStatus) > 0 {
		if err := assignInputField(input, "ConnectionStatus", _iotwirelessConnectionStatus); err != nil {
			log.Errorf("invalid --connection-status: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessDeviceRegistrationState) > 0 {
		if err := assignInputField(input, "DeviceRegistrationState", _iotwirelessDeviceRegistrationState); err != nil {
			log.Errorf("invalid --device-registration-state: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessJoin) > 0 {
		if err := assignInputField(input, "Join", _iotwirelessJoin); err != nil {
			log.Errorf("invalid --join: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessMessageDeliveryStatus) > 0 {
		if err := assignInputField(input, "MessageDeliveryStatus", _iotwirelessMessageDeliveryStatus); err != nil {
			log.Errorf("invalid --message-delivery-status: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessPartnerType) > 0 {
		if err := assignInputField(input, "PartnerType", _iotwirelessPartnerType); err != nil {
			log.Errorf("invalid --partner-type: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessProximity) > 0 {
		if err := assignInputField(input, "Proximity", _iotwirelessProximity); err != nil {
			log.Errorf("invalid --proximity: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateResourceEventConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update the position information of a given wireless device or a wireless
// gateway resource. The position coordinates are based on the [World Geodetic System (WGS84)].
//
// [World Geodetic System (WGS84)]: https://gisgeography.com/wgs84-world-geodetic-system/
func iotwireless_UpdateResourcePosition(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.UpdateResourcePositionInput{
		// ResourceIdentifier: *string, // Required
		// ResourceType: types.PositionResourceType, // Required
	}

	if len(_iotwirelessResourceIdentifier) > 0 {
		input.ResourceIdentifier = aws.String(_iotwirelessResourceIdentifier)
	}
	if len(_iotwirelessResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _iotwirelessResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessGeoJsonPayload) > 0 {
		if err := assignInputField(input, "GeoJsonPayload", _iotwirelessGeoJsonPayload); err != nil {
			log.Errorf("invalid --geo-json-payload: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateResourcePosition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates properties of a wireless device.
func iotwireless_UpdateWirelessDevice(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.UpdateWirelessDeviceInput{
		// Id: *string, // Required
	}

	if len(_iotwirelessId) > 0 {
		input.Id = aws.String(_iotwirelessId)
	}
	if len(_iotwirelessDescription) > 0 {
		input.Description = aws.String(_iotwirelessDescription)
	}
	if len(_iotwirelessDestinationName) > 0 {
		input.DestinationName = aws.String(_iotwirelessDestinationName)
	}
	if len(_iotwirelessLoRaWAN) > 0 {
		if err := assignInputField(input, "LoRaWAN", _iotwirelessLoRaWAN); err != nil {
			log.Errorf("invalid --lo-ra-wan: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessName) > 0 {
		input.Name = aws.String(_iotwirelessName)
	}
	if len(_iotwirelessPositioning) > 0 {
		if err := assignInputField(input, "Positioning", _iotwirelessPositioning); err != nil {
			log.Errorf("invalid --positioning: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessSidewalk) > 0 {
		if err := assignInputField(input, "Sidewalk", _iotwirelessSidewalk); err != nil {
			log.Errorf("invalid --sidewalk: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateWirelessDevice(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update an import task to add more devices to the task.
func iotwireless_UpdateWirelessDeviceImportTask(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.UpdateWirelessDeviceImportTaskInput{
		// Id: *string, // Required
		// Sidewalk: *types.SidewalkUpdateImportInfo, // Required
	}

	if len(_iotwirelessId) > 0 {
		input.Id = aws.String(_iotwirelessId)
	}
	if len(_iotwirelessSidewalk) > 0 {
		if err := assignInputField(input, "Sidewalk", _iotwirelessSidewalk); err != nil {
			log.Errorf("invalid --sidewalk: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateWirelessDeviceImportTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates properties of a wireless gateway.
func iotwireless_UpdateWirelessGateway(cfg aws.Config, client *iotwireless.Client) {
	input := &iotwireless.UpdateWirelessGatewayInput{
		// Id: *string, // Required
	}

	if len(_iotwirelessId) > 0 {
		input.Id = aws.String(_iotwirelessId)
	}
	if len(_iotwirelessDescription) > 0 {
		input.Description = aws.String(_iotwirelessDescription)
	}
	if len(_iotwirelessJoinEuiFilters) > 0 {
		if err := assignInputField(input, "JoinEuiFilters", _iotwirelessJoinEuiFilters); err != nil {
			log.Errorf("invalid --join-eui-filters: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessMaxEirp) > 0 {
		if err := assignInputField(input, "MaxEirp", _iotwirelessMaxEirp); err != nil {
			log.Errorf("invalid --max-eirp: %s", err.Error())
			return
		}
	}
	if len(_iotwirelessName) > 0 {
		input.Name = aws.String(_iotwirelessName)
	}
	if len(_iotwirelessNetIdFilters) > 0 {
		input.NetIdFilters = append([]string(nil), _iotwirelessNetIdFilters...)
	}

	if resp, err := client.UpdateWirelessGateway(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_iotwirelessCmd)
	_iotwirelessCmd.Flags().SortFlags = false

	_iotwirelessCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_iotwirelessCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_iotwirelessCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessAutoCreateTasks, "auto-create-tasks", "", "", "Auto Create Tasks")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessCellTowers, "cell-towers", "", "", "Cell Towers")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessClientRequestToken, "client-request-token", "", "", "Client Request Token")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessConfigurationName, "configuration-name", "", "", "Configuration Name")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessConnectionStatus, "connection-status", "", "", "Connection Status")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessDefaultLogLevel, "default-log-level", "", "", "Default Log Level")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessDescription, "description", "", "", "Description")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessDescriptor, "descriptor", "", "", "Descriptor")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessDestination, "destination", "", "", "Destination")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessDestinationName, "destination-name", "", "", "Destination Name")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessDeviceName, "device-name", "", "", "Device Name")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessDeviceProfileId, "device-profile-id", "", "", "Device Profile ID")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessDeviceProfileType, "device-profile-type", "", "", "Device Profile Type")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessDeviceRegistrationState, "device-registration-state", "", "", "Device Registration State")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessExpression, "expression", "", "", "Expression")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessExpressionType, "expression-type", "", "", "Expression Type")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessFirmwareUpdateImage, "firmware-update-image", "", "", "Firmware Update Image")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessFirmwareUpdateRole, "firmware-update-role", "", "", "Firmware Update Role")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessFragmentIntervalMS, "fragment-interval-ms", "", "", "Fragment Interval Ms")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessFragmentSizeBytes, "fragment-size-bytes", "", "", "Fragment Size Bytes")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessFuotaTaskId, "fuota-task-id", "", "", "Fuota Task ID")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessFuotaTaskLogOptions, "fuota-task-log-options", "", "", "Fuota Task Log Options")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessGeoJsonPayload, "geo-json-payload", "", "", "Geo JSON Payload")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessGnss, "gnss", "", "", "Gnss")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessId, "id", "", "", "ID")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessIdentifier, "identifier", "", "", "Identifier")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessIdentifierType, "identifier-type", "", "", "Identifier Type")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessIotCertificateId, "iot-certificate-id", "", "", "Iot Certificate ID")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessIp, "ip", "", "", "IP")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessJoin, "join", "", "", "Join")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessJoinEuiFilters, "join-eui-filters", "", "", "Join Eui Filters")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessLoRaWAN, "lo-ra-wan", "", "", "Lo Ra Wan")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessLogLevel, "log-level", "", "", "Log Level")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessMaxEirp, "max-eirp", "", "", "Max Eirp")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessMaxResults, "max-results", "", "", "Max Results")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessMessageDeliveryStatus, "message-delivery-status", "", "", "Message Delivery Status")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessMessageId, "message-id", "", "", "Message ID")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessMulticastGroupId, "multicast-group-id", "", "", "Multicast Group ID")
	_iotwirelessCmd.Flags().StringSliceVarP(&_iotwirelessMulticastGroups, "multicast-groups", "", nil, "Multicast Groups")
	_iotwirelessCmd.Flags().StringSliceVarP(&_iotwirelessMulticastGroupsToAdd, "multicast-groups-to-add", "", nil, "Multicast Groups To Add")
	_iotwirelessCmd.Flags().StringSliceVarP(&_iotwirelessMulticastGroupsToRemove, "multicast-groups-to-remove", "", nil, "Multicast Groups To Remove")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessName, "name", "", "", "Name")
	_iotwirelessCmd.Flags().StringSliceVarP(&_iotwirelessNetIdFilters, "net-id-filters", "", nil, "Net ID Filters")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessNextToken, "next-token", "", "", "Next Token")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessPartnerAccountId, "partner-account-id", "", "", "Partner Account ID")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessPartnerType, "partner-type", "", "", "Partner Type")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessPayloadData, "payload-data", "", "", "Payload Data")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessPosition, "position", "", "", "Position")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessPositioning, "positioning", "", "", "Positioning")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessProximity, "proximity", "", "", "Proximity")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessQueryString, "query-string", "", "", "Query String")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessRedundancyPercent, "redundancy-percent", "", "", "Redundancy Percent")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessResourceArn, "resource-arn", "", "", "Resource ARN")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessResourceIdentifier, "resource-identifier", "", "", "Resource Identifier")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessResourceType, "resource-type", "", "", "Resource Type")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessRoleArn, "role-arn", "", "", "Role ARN")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessServiceProfileId, "service-profile-id", "", "", "Service Profile ID")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessServiceType, "service-type", "", "", "Service Type")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessSidewalk, "sidewalk", "", "", "Sidewalk")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessSolvers, "solvers", "", "", "Solvers")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessStatus, "status", "", "", "Status")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessSummaryMetric, "summary-metric", "", "", "Summary Metric")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessSummaryMetricQueries, "summary-metric-queries", "", "", "Summary Metric Queries")
	_iotwirelessCmd.Flags().StringSliceVarP(&_iotwirelessTagKeys, "tag-keys", "", nil, "Tag Keys")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessTags, "tags", "", "", "Tags")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessTaskDefinitionType, "task-definition-type", "", "", "Task Definition Type")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessThingArn, "thing-arn", "", "", "Thing ARN")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessTimestamp, "timestamp", "", "", "Timestamp")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessTraceContent, "trace-content", "", "", "Trace Content")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessTransmitMode, "transmit-mode", "", "", "Transmit Mode")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessType, "type", "", "", "Type")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessUpdate, "update", "", "", "Update")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessWiFiAccessPoints, "wi-fi-access-points", "", "", "Wi Fi Access Points")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessWirelessDeviceId, "wireless-device-id", "", "", "Wireless Device ID")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessWirelessDeviceLogOptions, "wireless-device-log-options", "", "", "Wireless Device Log Options")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessWirelessDeviceType, "wireless-device-type", "", "", "Wireless Device Type")
	_iotwirelessCmd.Flags().StringSliceVarP(&_iotwirelessWirelessDevices, "wireless-devices", "", nil, "Wireless Devices")
	_iotwirelessCmd.Flags().StringSliceVarP(&_iotwirelessWirelessDevicesToAdd, "wireless-devices-to-add", "", nil, "Wireless Devices To Add")
	_iotwirelessCmd.Flags().StringSliceVarP(&_iotwirelessWirelessDevicesToRemove, "wireless-devices-to-remove", "", nil, "Wireless Devices To Remove")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessWirelessGatewayId, "wireless-gateway-id", "", "", "Wireless Gateway ID")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessWirelessGatewayLogOptions, "wireless-gateway-log-options", "", "", "Wireless Gateway Log Options")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessWirelessGatewayTaskDefinitionId, "wireless-gateway-task-definition-id", "", "", "Wireless Gateway Task Definition ID")
	_iotwirelessCmd.Flags().StringSliceVarP(&_iotwirelessWirelessGateways, "wireless-gateways", "", nil, "Wireless Gateways")
	_iotwirelessCmd.Flags().StringSliceVarP(&_iotwirelessWirelessGatewaysToAdd, "wireless-gateways-to-add", "", nil, "Wireless Gateways To Add")
	_iotwirelessCmd.Flags().StringSliceVarP(&_iotwirelessWirelessGatewaysToRemove, "wireless-gateways-to-remove", "", nil, "Wireless Gateways To Remove")
	_iotwirelessCmd.Flags().StringVarP(&_iotwirelessWirelessMetadata, "wireless-metadata", "", "", "Wireless Metadata")

	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessAssociateAwsAccountWithPartnerAccount, "associate-aws-account-with-partner-account", "", false, "Associate AWS Account With Partner Account")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessAssociateMulticastGroupWithFuotaTask, "associate-multicast-group-with-fuota-task", "", false, "Associate Multicast Group With Fuota Task")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessAssociateWirelessDeviceWithFuotaTask, "associate-wireless-device-with-fuota-task", "", false, "Associate Wireless Device With Fuota Task")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessAssociateWirelessDeviceWithMulticastGroup, "associate-wireless-device-with-multicast-group", "", false, "Associate Wireless Device With Multicast Group")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessAssociateWirelessDeviceWithThing, "associate-wireless-device-with-thing", "", false, "Associate Wireless Device With Thing")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessAssociateWirelessGatewayWithCertificate, "associate-wireless-gateway-with-certificate", "", false, "Associate Wireless Gateway With Certificate")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessAssociateWirelessGatewayWithThing, "associate-wireless-gateway-with-thing", "", false, "Associate Wireless Gateway With Thing")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessCancelMulticastGroupSession, "cancel-multicast-group-session", "", false, "Cancel Multicast Group Session")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessCreateDestination, "create-destination", "", false, "Create Destination")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessCreateDeviceProfile, "create-device-profile", "", false, "Create Device Profile")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessCreateFuotaTask, "create-fuota-task", "", false, "Create Fuota Task")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessCreateMulticastGroup, "create-multicast-group", "", false, "Create Multicast Group")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessCreateNetworkAnalyzerConfiguration, "create-network-analyzer-configuration", "", false, "Create Network Analyzer Configuration")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessCreateServiceProfile, "create-service-profile", "", false, "Create Service Profile")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessCreateWirelessDevice, "create-wireless-device", "", false, "Create Wireless Device")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessCreateWirelessGateway, "create-wireless-gateway", "", false, "Create Wireless Gateway")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessCreateWirelessGatewayTask, "create-wireless-gateway-task", "", false, "Create Wireless Gateway Task")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessCreateWirelessGatewayTaskDefinition, "create-wireless-gateway-task-definition", "", false, "Create Wireless Gateway Task Definition")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessDeleteDestination, "delete-destination", "", false, "Delete Destination")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessDeleteDeviceProfile, "delete-device-profile", "", false, "Delete Device Profile")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessDeleteFuotaTask, "delete-fuota-task", "", false, "Delete Fuota Task")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessDeleteMulticastGroup, "delete-multicast-group", "", false, "Delete Multicast Group")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessDeleteNetworkAnalyzerConfiguration, "delete-network-analyzer-configuration", "", false, "Delete Network Analyzer Configuration")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessDeleteQueuedMessages, "delete-queued-messages", "", false, "Delete Queued Messages")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessDeleteServiceProfile, "delete-service-profile", "", false, "Delete Service Profile")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessDeleteWirelessDevice, "delete-wireless-device", "", false, "Delete Wireless Device")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessDeleteWirelessDeviceImportTask, "delete-wireless-device-import-task", "", false, "Delete Wireless Device Import Task")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessDeleteWirelessGateway, "delete-wireless-gateway", "", false, "Delete Wireless Gateway")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessDeleteWirelessGatewayTask, "delete-wireless-gateway-task", "", false, "Delete Wireless Gateway Task")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessDeleteWirelessGatewayTaskDefinition, "delete-wireless-gateway-task-definition", "", false, "Delete Wireless Gateway Task Definition")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessDeregisterWirelessDevice, "deregister-wireless-device", "", false, "Deregister Wireless Device")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessDisassociateAwsAccountFromPartnerAccount, "disassociate-aws-account-from-partner-account", "", false, "Disassociate AWS Account From Partner Account")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessDisassociateMulticastGroupFromFuotaTask, "disassociate-multicast-group-from-fuota-task", "", false, "Disassociate Multicast Group From Fuota Task")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessDisassociateWirelessDeviceFromFuotaTask, "disassociate-wireless-device-from-fuota-task", "", false, "Disassociate Wireless Device From Fuota Task")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessDisassociateWirelessDeviceFromMulticastGroup, "disassociate-wireless-device-from-multicast-group", "", false, "Disassociate Wireless Device From Multicast Group")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessDisassociateWirelessDeviceFromThing, "disassociate-wireless-device-from-thing", "", false, "Disassociate Wireless Device From Thing")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessDisassociateWirelessGatewayFromCertificate, "disassociate-wireless-gateway-from-certificate", "", false, "Disassociate Wireless Gateway From Certificate")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessDisassociateWirelessGatewayFromThing, "disassociate-wireless-gateway-from-thing", "", false, "Disassociate Wireless Gateway From Thing")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessGetDestination, "get-destination", "", false, "Get Destination")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessGetDeviceProfile, "get-device-profile", "", false, "Get Device Profile")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessGetEventConfigurationByResourceTypes, "get-event-configuration-by-resource-types", "", false, "Get Event Configuration By Resource Types")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessGetFuotaTask, "get-fuota-task", "", false, "Get Fuota Task")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessGetLogLevelsByResourceTypes, "get-log-levels-by-resource-types", "", false, "Get Log Levels By Resource Types")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessGetMetricConfiguration, "get-metric-configuration", "", false, "Get Metric Configuration")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessGetMetrics, "get-metrics", "", false, "Get Metrics")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessGetMulticastGroup, "get-multicast-group", "", false, "Get Multicast Group")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessGetMulticastGroupSession, "get-multicast-group-session", "", false, "Get Multicast Group Session")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessGetNetworkAnalyzerConfiguration, "get-network-analyzer-configuration", "", false, "Get Network Analyzer Configuration")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessGetPartnerAccount, "get-partner-account", "", false, "Get Partner Account")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessGetPosition, "get-position", "", false, "Get Position")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessGetPositionConfiguration, "get-position-configuration", "", false, "Get Position Configuration")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessGetPositionEstimate, "get-position-estimate", "", false, "Get Position Estimate")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessGetResourceEventConfiguration, "get-resource-event-configuration", "", false, "Get Resource Event Configuration")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessGetResourceLogLevel, "get-resource-log-level", "", false, "Get Resource Log Level")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessGetResourcePosition, "get-resource-position", "", false, "Get Resource Position")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessGetServiceEndpoint, "get-service-endpoint", "", false, "Get Service Endpoint")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessGetServiceProfile, "get-service-profile", "", false, "Get Service Profile")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessGetWirelessDevice, "get-wireless-device", "", false, "Get Wireless Device")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessGetWirelessDeviceImportTask, "get-wireless-device-import-task", "", false, "Get Wireless Device Import Task")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessGetWirelessDeviceStatistics, "get-wireless-device-statistics", "", false, "Get Wireless Device Statistics")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessGetWirelessGateway, "get-wireless-gateway", "", false, "Get Wireless Gateway")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessGetWirelessGatewayCertificate, "get-wireless-gateway-certificate", "", false, "Get Wireless Gateway Certificate")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessGetWirelessGatewayFirmwareInformation, "get-wireless-gateway-firmware-information", "", false, "Get Wireless Gateway Firmware Information")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessGetWirelessGatewayStatistics, "get-wireless-gateway-statistics", "", false, "Get Wireless Gateway Statistics")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessGetWirelessGatewayTask, "get-wireless-gateway-task", "", false, "Get Wireless Gateway Task")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessGetWirelessGatewayTaskDefinition, "get-wireless-gateway-task-definition", "", false, "Get Wireless Gateway Task Definition")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessListDestinations, "list-destinations", "", false, "List Destinations")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessListDeviceProfiles, "list-device-profiles", "", false, "List Device Profiles")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessListDevicesForWirelessDeviceImportTask, "list-devices-for-wireless-device-import-task", "", false, "List Devices For Wireless Device Import Task")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessListEventConfigurations, "list-event-configurations", "", false, "List Event Configurations")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessListFuotaTasks, "list-fuota-tasks", "", false, "List Fuota Tasks")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessListMulticastGroups, "list-multicast-groups", "", false, "List Multicast Groups")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessListMulticastGroupsByFuotaTask, "list-multicast-groups-by-fuota-task", "", false, "List Multicast Groups By Fuota Task")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessListNetworkAnalyzerConfigurations, "list-network-analyzer-configurations", "", false, "List Network Analyzer Configurations")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessListPartnerAccounts, "list-partner-accounts", "", false, "List Partner Accounts")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessListPositionConfigurations, "list-position-configurations", "", false, "List Position Configurations")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessListQueuedMessages, "list-queued-messages", "", false, "List Queued Messages")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessListServiceProfiles, "list-service-profiles", "", false, "List Service Profiles")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessListWirelessDeviceImportTasks, "list-wireless-device-import-tasks", "", false, "List Wireless Device Import Tasks")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessListWirelessDevices, "list-wireless-devices", "", false, "List Wireless Devices")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessListWirelessGatewayTaskDefinitions, "list-wireless-gateway-task-definitions", "", false, "List Wireless Gateway Task Definitions")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessListWirelessGateways, "list-wireless-gateways", "", false, "List Wireless Gateways")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessPutPositionConfiguration, "put-position-configuration", "", false, "Put Position Configuration")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessPutResourceLogLevel, "put-resource-log-level", "", false, "Put Resource Log Level")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessResetAllResourceLogLevels, "reset-all-resource-log-levels", "", false, "Reset All Resource Log Levels")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessResetResourceLogLevel, "reset-resource-log-level", "", false, "Reset Resource Log Level")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessSendDataToMulticastGroup, "send-data-to-multicast-group", "", false, "Send Data To Multicast Group")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessSendDataToWirelessDevice, "send-data-to-wireless-device", "", false, "Send Data To Wireless Device")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessStartBulkAssociateWirelessDeviceWithMulticastGroup, "start-bulk-associate-wireless-device-with-multicast-group", "", false, "Start Bulk Associate Wireless Device With Multicast Group")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessStartBulkDisassociateWirelessDeviceFromMulticastGroup, "start-bulk-disassociate-wireless-device-from-multicast-group", "", false, "Start Bulk Disassociate Wireless Device From Multicast Group")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessStartFuotaTask, "start-fuota-task", "", false, "Start Fuota Task")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessStartMulticastGroupSession, "start-multicast-group-session", "", false, "Start Multicast Group Session")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessStartSingleWirelessDeviceImportTask, "start-single-wireless-device-import-task", "", false, "Start Single Wireless Device Import Task")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessStartWirelessDeviceImportTask, "start-wireless-device-import-task", "", false, "Start Wireless Device Import Task")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessTagResource, "tag-resource", "", false, "Tag Resource")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessTestWirelessDevice, "test-wireless-device", "", false, "Test Wireless Device")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessUntagResource, "untag-resource", "", false, "Untag Resource")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessUpdateDestination, "update-destination", "", false, "Update Destination")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessUpdateEventConfigurationByResourceTypes, "update-event-configuration-by-resource-types", "", false, "Update Event Configuration By Resource Types")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessUpdateFuotaTask, "update-fuota-task", "", false, "Update Fuota Task")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessUpdateLogLevelsByResourceTypes, "update-log-levels-by-resource-types", "", false, "Update Log Levels By Resource Types")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessUpdateMetricConfiguration, "update-metric-configuration", "", false, "Update Metric Configuration")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessUpdateMulticastGroup, "update-multicast-group", "", false, "Update Multicast Group")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessUpdateNetworkAnalyzerConfiguration, "update-network-analyzer-configuration", "", false, "Update Network Analyzer Configuration")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessUpdatePartnerAccount, "update-partner-account", "", false, "Update Partner Account")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessUpdatePosition, "update-position", "", false, "Update Position")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessUpdateResourceEventConfiguration, "update-resource-event-configuration", "", false, "Update Resource Event Configuration")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessUpdateResourcePosition, "update-resource-position", "", false, "Update Resource Position")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessUpdateWirelessDevice, "update-wireless-device", "", false, "Update Wireless Device")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessUpdateWirelessDeviceImportTask, "update-wireless-device-import-task", "", false, "Update Wireless Device Import Task")
	_iotwirelessCmd.Flags().BoolVarP(&_iotwirelessUpdateWirelessGateway, "update-wireless-gateway", "", false, "Update Wireless Gateway")

}
