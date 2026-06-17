package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/medialive"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// medialiveCmd represents the medialive command
var _medialiveCmd = &cobra.Command{
	Use:   "medialive",
	Short: "AWS medialive CLI",
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
		client := medialive.NewFromConfig(cfg)
		if _medialiveAcceptInputDeviceTransfer {
			medialive_AcceptInputDeviceTransfer(cfg, client)
			return
		}
		if _medialiveBatchDelete {
			medialive_BatchDelete(cfg, client)
			return
		}
		if _medialiveBatchStart {
			medialive_BatchStart(cfg, client)
			return
		}
		if _medialiveBatchStop {
			medialive_BatchStop(cfg, client)
			return
		}
		if _medialiveBatchUpdateSchedule {
			medialive_BatchUpdateSchedule(cfg, client)
			return
		}
		if _medialiveCancelInputDeviceTransfer {
			medialive_CancelInputDeviceTransfer(cfg, client)
			return
		}
		if _medialiveClaimDevice {
			medialive_ClaimDevice(cfg, client)
			return
		}
		if _medialiveCreateChannel {
			medialive_CreateChannel(cfg, client)
			return
		}
		if _medialiveCreateChannelPlacementGroup {
			medialive_CreateChannelPlacementGroup(cfg, client)
			return
		}
		if _medialiveCreateCloudWatchAlarmTemplate {
			medialive_CreateCloudWatchAlarmTemplate(cfg, client)
			return
		}
		if _medialiveCreateCloudWatchAlarmTemplateGroup {
			medialive_CreateCloudWatchAlarmTemplateGroup(cfg, client)
			return
		}
		if _medialiveCreateCluster {
			medialive_CreateCluster(cfg, client)
			return
		}
		if _medialiveCreateEventBridgeRuleTemplate {
			medialive_CreateEventBridgeRuleTemplate(cfg, client)
			return
		}
		if _medialiveCreateEventBridgeRuleTemplateGroup {
			medialive_CreateEventBridgeRuleTemplateGroup(cfg, client)
			return
		}
		if _medialiveCreateInput {
			medialive_CreateInput(cfg, client)
			return
		}
		if _medialiveCreateInputSecurityGroup {
			medialive_CreateInputSecurityGroup(cfg, client)
			return
		}
		if _medialiveCreateMultiplex {
			medialive_CreateMultiplex(cfg, client)
			return
		}
		if _medialiveCreateMultiplexProgram {
			medialive_CreateMultiplexProgram(cfg, client)
			return
		}
		if _medialiveCreateNetwork {
			medialive_CreateNetwork(cfg, client)
			return
		}
		if _medialiveCreateNode {
			medialive_CreateNode(cfg, client)
			return
		}
		if _medialiveCreateNodeRegistrationScript {
			medialive_CreateNodeRegistrationScript(cfg, client)
			return
		}
		if _medialiveCreatePartnerInput {
			medialive_CreatePartnerInput(cfg, client)
			return
		}
		if _medialiveCreateSdiSource {
			medialive_CreateSdiSource(cfg, client)
			return
		}
		if _medialiveCreateSignalMap {
			medialive_CreateSignalMap(cfg, client)
			return
		}
		if _medialiveCreateTags {
			medialive_CreateTags(cfg, client)
			return
		}
		if _medialiveDeleteChannel {
			medialive_DeleteChannel(cfg, client)
			return
		}
		if _medialiveDeleteChannelPlacementGroup {
			medialive_DeleteChannelPlacementGroup(cfg, client)
			return
		}
		if _medialiveDeleteCloudWatchAlarmTemplate {
			medialive_DeleteCloudWatchAlarmTemplate(cfg, client)
			return
		}
		if _medialiveDeleteCloudWatchAlarmTemplateGroup {
			medialive_DeleteCloudWatchAlarmTemplateGroup(cfg, client)
			return
		}
		if _medialiveDeleteCluster {
			medialive_DeleteCluster(cfg, client)
			return
		}
		if _medialiveDeleteEventBridgeRuleTemplate {
			medialive_DeleteEventBridgeRuleTemplate(cfg, client)
			return
		}
		if _medialiveDeleteEventBridgeRuleTemplateGroup {
			medialive_DeleteEventBridgeRuleTemplateGroup(cfg, client)
			return
		}
		if _medialiveDeleteInput {
			medialive_DeleteInput(cfg, client)
			return
		}
		if _medialiveDeleteInputSecurityGroup {
			medialive_DeleteInputSecurityGroup(cfg, client)
			return
		}
		if _medialiveDeleteMultiplex {
			medialive_DeleteMultiplex(cfg, client)
			return
		}
		if _medialiveDeleteMultiplexProgram {
			medialive_DeleteMultiplexProgram(cfg, client)
			return
		}
		if _medialiveDeleteNetwork {
			medialive_DeleteNetwork(cfg, client)
			return
		}
		if _medialiveDeleteNode {
			medialive_DeleteNode(cfg, client)
			return
		}
		if _medialiveDeleteReservation {
			medialive_DeleteReservation(cfg, client)
			return
		}
		if _medialiveDeleteSchedule {
			medialive_DeleteSchedule(cfg, client)
			return
		}
		if _medialiveDeleteSdiSource {
			medialive_DeleteSdiSource(cfg, client)
			return
		}
		if _medialiveDeleteSignalMap {
			medialive_DeleteSignalMap(cfg, client)
			return
		}
		if _medialiveDeleteTags {
			medialive_DeleteTags(cfg, client)
			return
		}
		if _medialiveDescribeAccountConfiguration {
			medialive_DescribeAccountConfiguration(cfg, client)
			return
		}
		if _medialiveDescribeChannel {
			medialive_DescribeChannel(cfg, client)
			return
		}
		if _medialiveDescribeChannelPlacementGroup {
			medialive_DescribeChannelPlacementGroup(cfg, client)
			return
		}
		if _medialiveDescribeCluster {
			medialive_DescribeCluster(cfg, client)
			return
		}
		if _medialiveDescribeInput {
			medialive_DescribeInput(cfg, client)
			return
		}
		if _medialiveDescribeInputDevice {
			medialive_DescribeInputDevice(cfg, client)
			return
		}
		if _medialiveDescribeInputDeviceThumbnail {
			medialive_DescribeInputDeviceThumbnail(cfg, client)
			return
		}
		if _medialiveDescribeInputSecurityGroup {
			medialive_DescribeInputSecurityGroup(cfg, client)
			return
		}
		if _medialiveDescribeMultiplex {
			medialive_DescribeMultiplex(cfg, client)
			return
		}
		if _medialiveDescribeMultiplexProgram {
			medialive_DescribeMultiplexProgram(cfg, client)
			return
		}
		if _medialiveDescribeNetwork {
			medialive_DescribeNetwork(cfg, client)
			return
		}
		if _medialiveDescribeNode {
			medialive_DescribeNode(cfg, client)
			return
		}
		if _medialiveDescribeOffering {
			medialive_DescribeOffering(cfg, client)
			return
		}
		if _medialiveDescribeReservation {
			medialive_DescribeReservation(cfg, client)
			return
		}
		if _medialiveDescribeSchedule {
			medialive_DescribeSchedule(cfg, client)
			return
		}
		if _medialiveDescribeSdiSource {
			medialive_DescribeSdiSource(cfg, client)
			return
		}
		if _medialiveDescribeThumbnails {
			medialive_DescribeThumbnails(cfg, client)
			return
		}
		if _medialiveGetCloudWatchAlarmTemplate {
			medialive_GetCloudWatchAlarmTemplate(cfg, client)
			return
		}
		if _medialiveGetCloudWatchAlarmTemplateGroup {
			medialive_GetCloudWatchAlarmTemplateGroup(cfg, client)
			return
		}
		if _medialiveGetEventBridgeRuleTemplate {
			medialive_GetEventBridgeRuleTemplate(cfg, client)
			return
		}
		if _medialiveGetEventBridgeRuleTemplateGroup {
			medialive_GetEventBridgeRuleTemplateGroup(cfg, client)
			return
		}
		if _medialiveGetSignalMap {
			medialive_GetSignalMap(cfg, client)
			return
		}
		if _medialiveListAlerts {
			medialive_ListAlerts(cfg, client)
			return
		}
		if _medialiveListChannelPlacementGroups {
			medialive_ListChannelPlacementGroups(cfg, client)
			return
		}
		if _medialiveListChannels {
			medialive_ListChannels(cfg, client)
			return
		}
		if _medialiveListCloudWatchAlarmTemplateGroups {
			medialive_ListCloudWatchAlarmTemplateGroups(cfg, client)
			return
		}
		if _medialiveListCloudWatchAlarmTemplates {
			medialive_ListCloudWatchAlarmTemplates(cfg, client)
			return
		}
		if _medialiveListClusterAlerts {
			medialive_ListClusterAlerts(cfg, client)
			return
		}
		if _medialiveListClusters {
			medialive_ListClusters(cfg, client)
			return
		}
		if _medialiveListEventBridgeRuleTemplateGroups {
			medialive_ListEventBridgeRuleTemplateGroups(cfg, client)
			return
		}
		if _medialiveListEventBridgeRuleTemplates {
			medialive_ListEventBridgeRuleTemplates(cfg, client)
			return
		}
		if _medialiveListInputDeviceTransfers {
			medialive_ListInputDeviceTransfers(cfg, client)
			return
		}
		if _medialiveListInputDevices {
			medialive_ListInputDevices(cfg, client)
			return
		}
		if _medialiveListInputSecurityGroups {
			medialive_ListInputSecurityGroups(cfg, client)
			return
		}
		if _medialiveListInputs {
			medialive_ListInputs(cfg, client)
			return
		}
		if _medialiveListMultiplexAlerts {
			medialive_ListMultiplexAlerts(cfg, client)
			return
		}
		if _medialiveListMultiplexPrograms {
			medialive_ListMultiplexPrograms(cfg, client)
			return
		}
		if _medialiveListMultiplexes {
			medialive_ListMultiplexes(cfg, client)
			return
		}
		if _medialiveListNetworks {
			medialive_ListNetworks(cfg, client)
			return
		}
		if _medialiveListNodes {
			medialive_ListNodes(cfg, client)
			return
		}
		if _medialiveListOfferings {
			medialive_ListOfferings(cfg, client)
			return
		}
		if _medialiveListReservations {
			medialive_ListReservations(cfg, client)
			return
		}
		if _medialiveListSdiSources {
			medialive_ListSdiSources(cfg, client)
			return
		}
		if _medialiveListSignalMaps {
			medialive_ListSignalMaps(cfg, client)
			return
		}
		if _medialiveListTagsForResource {
			medialive_ListTagsForResource(cfg, client)
			return
		}
		if _medialiveListVersions {
			medialive_ListVersions(cfg, client)
			return
		}
		if _medialivePurchaseOffering {
			medialive_PurchaseOffering(cfg, client)
			return
		}
		if _medialiveRebootInputDevice {
			medialive_RebootInputDevice(cfg, client)
			return
		}
		if _medialiveRejectInputDeviceTransfer {
			medialive_RejectInputDeviceTransfer(cfg, client)
			return
		}
		if _medialiveRestartChannelPipelines {
			medialive_RestartChannelPipelines(cfg, client)
			return
		}
		if _medialiveStartChannel {
			medialive_StartChannel(cfg, client)
			return
		}
		if _medialiveStartDeleteMonitorDeployment {
			medialive_StartDeleteMonitorDeployment(cfg, client)
			return
		}
		if _medialiveStartInputDevice {
			medialive_StartInputDevice(cfg, client)
			return
		}
		if _medialiveStartInputDeviceMaintenanceWindow {
			medialive_StartInputDeviceMaintenanceWindow(cfg, client)
			return
		}
		if _medialiveStartMonitorDeployment {
			medialive_StartMonitorDeployment(cfg, client)
			return
		}
		if _medialiveStartMultiplex {
			medialive_StartMultiplex(cfg, client)
			return
		}
		if _medialiveStartUpdateSignalMap {
			medialive_StartUpdateSignalMap(cfg, client)
			return
		}
		if _medialiveStopChannel {
			medialive_StopChannel(cfg, client)
			return
		}
		if _medialiveStopInputDevice {
			medialive_StopInputDevice(cfg, client)
			return
		}
		if _medialiveStopMultiplex {
			medialive_StopMultiplex(cfg, client)
			return
		}
		if _medialiveTransferInputDevice {
			medialive_TransferInputDevice(cfg, client)
			return
		}
		if _medialiveUpdateAccountConfiguration {
			medialive_UpdateAccountConfiguration(cfg, client)
			return
		}
		if _medialiveUpdateChannel {
			medialive_UpdateChannel(cfg, client)
			return
		}
		if _medialiveUpdateChannelClass {
			medialive_UpdateChannelClass(cfg, client)
			return
		}
		if _medialiveUpdateChannelPlacementGroup {
			medialive_UpdateChannelPlacementGroup(cfg, client)
			return
		}
		if _medialiveUpdateCloudWatchAlarmTemplate {
			medialive_UpdateCloudWatchAlarmTemplate(cfg, client)
			return
		}
		if _medialiveUpdateCloudWatchAlarmTemplateGroup {
			medialive_UpdateCloudWatchAlarmTemplateGroup(cfg, client)
			return
		}
		if _medialiveUpdateCluster {
			medialive_UpdateCluster(cfg, client)
			return
		}
		if _medialiveUpdateEventBridgeRuleTemplate {
			medialive_UpdateEventBridgeRuleTemplate(cfg, client)
			return
		}
		if _medialiveUpdateEventBridgeRuleTemplateGroup {
			medialive_UpdateEventBridgeRuleTemplateGroup(cfg, client)
			return
		}
		if _medialiveUpdateInput {
			medialive_UpdateInput(cfg, client)
			return
		}
		if _medialiveUpdateInputDevice {
			medialive_UpdateInputDevice(cfg, client)
			return
		}
		if _medialiveUpdateInputSecurityGroup {
			medialive_UpdateInputSecurityGroup(cfg, client)
			return
		}
		if _medialiveUpdateMultiplex {
			medialive_UpdateMultiplex(cfg, client)
			return
		}
		if _medialiveUpdateMultiplexProgram {
			medialive_UpdateMultiplexProgram(cfg, client)
			return
		}
		if _medialiveUpdateNetwork {
			medialive_UpdateNetwork(cfg, client)
			return
		}
		if _medialiveUpdateNode {
			medialive_UpdateNode(cfg, client)
			return
		}
		if _medialiveUpdateNodeState {
			medialive_UpdateNodeState(cfg, client)
			return
		}
		if _medialiveUpdateReservation {
			medialive_UpdateReservation(cfg, client)
			return
		}
		if _medialiveUpdateSdiSource {
			medialive_UpdateSdiSource(cfg, client)
			return
		}

	},
}

var (
	_medialiveAcceptInputDeviceTransfer          bool
	_medialiveBatchDelete                        bool
	_medialiveBatchStart                         bool
	_medialiveBatchStop                          bool
	_medialiveBatchUpdateSchedule                bool
	_medialiveCancelInputDeviceTransfer          bool
	_medialiveClaimDevice                        bool
	_medialiveCreateChannel                      bool
	_medialiveCreateChannelPlacementGroup        bool
	_medialiveCreateCloudWatchAlarmTemplate      bool
	_medialiveCreateCloudWatchAlarmTemplateGroup bool
	_medialiveCreateCluster                      bool
	_medialiveCreateEventBridgeRuleTemplate      bool
	_medialiveCreateEventBridgeRuleTemplateGroup bool
	_medialiveCreateInput                        bool
	_medialiveCreateInputSecurityGroup           bool
	_medialiveCreateMultiplex                    bool
	_medialiveCreateMultiplexProgram             bool
	_medialiveCreateNetwork                      bool
	_medialiveCreateNode                         bool
	_medialiveCreateNodeRegistrationScript       bool
	_medialiveCreatePartnerInput                 bool
	_medialiveCreateSdiSource                    bool
	_medialiveCreateSignalMap                    bool
	_medialiveCreateTags                         bool
	_medialiveDeleteChannel                      bool
	_medialiveDeleteChannelPlacementGroup        bool
	_medialiveDeleteCloudWatchAlarmTemplate      bool
	_medialiveDeleteCloudWatchAlarmTemplateGroup bool
	_medialiveDeleteCluster                      bool
	_medialiveDeleteEventBridgeRuleTemplate      bool
	_medialiveDeleteEventBridgeRuleTemplateGroup bool
	_medialiveDeleteInput                        bool
	_medialiveDeleteInputSecurityGroup           bool
	_medialiveDeleteMultiplex                    bool
	_medialiveDeleteMultiplexProgram             bool
	_medialiveDeleteNetwork                      bool
	_medialiveDeleteNode                         bool
	_medialiveDeleteReservation                  bool
	_medialiveDeleteSchedule                     bool
	_medialiveDeleteSdiSource                    bool
	_medialiveDeleteSignalMap                    bool
	_medialiveDeleteTags                         bool
	_medialiveDescribeAccountConfiguration       bool
	_medialiveDescribeChannel                    bool
	_medialiveDescribeChannelPlacementGroup      bool
	_medialiveDescribeCluster                    bool
	_medialiveDescribeInput                      bool
	_medialiveDescribeInputDevice                bool
	_medialiveDescribeInputDeviceThumbnail       bool
	_medialiveDescribeInputSecurityGroup         bool
	_medialiveDescribeMultiplex                  bool
	_medialiveDescribeMultiplexProgram           bool
	_medialiveDescribeNetwork                    bool
	_medialiveDescribeNode                       bool
	_medialiveDescribeOffering                   bool
	_medialiveDescribeReservation                bool
	_medialiveDescribeSchedule                   bool
	_medialiveDescribeSdiSource                  bool
	_medialiveDescribeThumbnails                 bool
	_medialiveGetCloudWatchAlarmTemplate         bool
	_medialiveGetCloudWatchAlarmTemplateGroup    bool
	_medialiveGetEventBridgeRuleTemplate         bool
	_medialiveGetEventBridgeRuleTemplateGroup    bool
	_medialiveGetSignalMap                       bool
	_medialiveListAlerts                         bool
	_medialiveListChannelPlacementGroups         bool
	_medialiveListChannels                       bool
	_medialiveListCloudWatchAlarmTemplateGroups  bool
	_medialiveListCloudWatchAlarmTemplates       bool
	_medialiveListClusterAlerts                  bool
	_medialiveListClusters                       bool
	_medialiveListEventBridgeRuleTemplateGroups  bool
	_medialiveListEventBridgeRuleTemplates       bool
	_medialiveListInputDeviceTransfers           bool
	_medialiveListInputDevices                   bool
	_medialiveListInputSecurityGroups            bool
	_medialiveListInputs                         bool
	_medialiveListMultiplexAlerts                bool
	_medialiveListMultiplexPrograms              bool
	_medialiveListMultiplexes                    bool
	_medialiveListNetworks                       bool
	_medialiveListNodes                          bool
	_medialiveListOfferings                      bool
	_medialiveListReservations                   bool
	_medialiveListSdiSources                     bool
	_medialiveListSignalMaps                     bool
	_medialiveListTagsForResource                bool
	_medialiveListVersions                       bool
	_medialivePurchaseOffering                   bool
	_medialiveRebootInputDevice                  bool
	_medialiveRejectInputDeviceTransfer          bool
	_medialiveRestartChannelPipelines            bool
	_medialiveStartChannel                       bool
	_medialiveStartDeleteMonitorDeployment       bool
	_medialiveStartInputDevice                   bool
	_medialiveStartInputDeviceMaintenanceWindow  bool
	_medialiveStartMonitorDeployment             bool
	_medialiveStartMultiplex                     bool
	_medialiveStartUpdateSignalMap               bool
	_medialiveStopChannel                        bool
	_medialiveStopInputDevice                    bool
	_medialiveStopMultiplex                      bool
	_medialiveTransferInputDevice                bool
	_medialiveUpdateAccountConfiguration         bool
	_medialiveUpdateChannel                      bool
	_medialiveUpdateChannelClass                 bool
	_medialiveUpdateChannelPlacementGroup        bool
	_medialiveUpdateCloudWatchAlarmTemplate      bool
	_medialiveUpdateCloudWatchAlarmTemplateGroup bool
	_medialiveUpdateCluster                      bool
	_medialiveUpdateEventBridgeRuleTemplate      bool
	_medialiveUpdateEventBridgeRuleTemplateGroup bool
	_medialiveUpdateInput                        bool
	_medialiveUpdateInputDevice                  bool
	_medialiveUpdateInputSecurityGroup           bool
	_medialiveUpdateMultiplex                    bool
	_medialiveUpdateMultiplexProgram             bool
	_medialiveUpdateNetwork                      bool
	_medialiveUpdateNode                         bool
	_medialiveUpdateNodeState                    bool
	_medialiveUpdateReservation                  bool
	_medialiveUpdateSdiSource                    bool

	_medialiveAccept                                  string
	_medialiveAccountConfiguration                    string
	_medialiveAnywhereSettings                        string
	_medialiveAvailabilityZone                        string
	_medialiveAvailabilityZones                       []string
	_medialiveCdiInputSpecification                   string
	_medialiveChannelClass                            string
	_medialiveChannelConfiguration                    string
	_medialiveChannelEngineVersion                    string
	_medialiveChannelId                               string
	_medialiveChannelIds                              []string
	_medialiveChannelPlacementGroupId                 string
	_medialiveChannelSecurityGroups                   []string
	_medialiveCloudWatchAlarmTemplateGroupIdentifier  string
	_medialiveCloudWatchAlarmTemplateGroupIdentifiers []string
	_medialiveClusterId                               string
	_medialiveClusterType                             string
	_medialiveCodec                                   string
	_medialiveComparisonOperator                      string
	_medialiveCount                                   string
	_medialiveCreates                                 string
	_medialiveDatapointsToAlarm                       string
	_medialiveDeletes                                 string
	_medialiveDescription                             string
	_medialiveDestinations                            string
	_medialiveDiscoveryEntryPointArn                  string
	_medialiveDryRun                                  string
	_medialiveDuration                                string
	_medialiveEncoderSettings                         string
	_medialiveEvaluationPeriods                       string
	_medialiveEventBridgeRuleTemplateGroupIdentifier  string
	_medialiveEventBridgeRuleTemplateGroupIdentifiers []string
	_medialiveEventTargets                            string
	_medialiveEventType                               string
	_medialiveForce                                   string
	_medialiveForceRediscovery                        string
	_medialiveGroupIdentifier                         string
	_medialiveHdDeviceSettings                        string
	_medialiveId                                      string
	_medialiveIdentifier                              string
	_medialiveInferenceSettings                       string
	_medialiveInputAttachments                        string
	_medialiveInputDeviceId                           string
	_medialiveInputDevices                            string
	_medialiveInputId                                 string
	_medialiveInputIds                                []string
	_medialiveInputNetworkLocation                    string
	_medialiveInputSecurityGroupId                    string
	_medialiveInputSecurityGroupIds                   []string
	_medialiveInputSecurityGroups                     []string
	_medialiveInputSpecification                      string
	_medialiveInstanceRoleArn                         string
	_medialiveIpPools                                 string
	_medialiveLinkedChannelSettings                   string
	_medialiveLogLevel                                string
	_medialiveMaintenance                             string
	_medialiveMaxResults                              string
	_medialiveMaximumBitrate                          string
	_medialiveMaximumFramerate                        string
	_medialiveMediaConnectFlows                       string
	_medialiveMetricName                              string
	_medialiveMode                                    string
	_medialiveMulticastSettings                       string
	_medialiveMultiplexId                             string
	_medialiveMultiplexIds                            []string
	_medialiveMultiplexProgramSettings                string
	_medialiveMultiplexSettings                       string
	_medialiveName                                    string
	_medialiveNetworkId                               string
	_medialiveNetworkSettings                         string
	_medialiveNextToken                               string
	_medialiveNodeId                                  string
	_medialiveNodeInterfaceMappings                   string
	_medialiveNodes                                   []string
	_medialiveOfferingId                              string
	_medialivePacketIdentifiersMapping                string
	_medialivePeriod                                  string
	_medialivePipelineId                              string
	_medialivePipelineIds                             string
	_medialiveProgramName                             string
	_medialiveRenewalSettings                         string
	_medialiveRequestId                               string
	_medialiveReservationId                           string
	_medialiveReserved                                string
	_medialiveResolution                              string
	_medialiveResourceArn                             string
	_medialiveResourceType                            string
	_medialiveRole                                    string
	_medialiveRoleArn                                 string
	_medialiveRouterSettings                          string
	_medialiveRoutes                                  string
	_medialiveScope                                   string
	_medialiveSdiSourceId                             string
	_medialiveSdiSourceMappings                       string
	_medialiveSdiSources                              []string
	_medialiveSignalMapIdentifier                     string
	_medialiveSmpte2110ReceiverGroupSettings          string
	_medialiveSources                                 string
	_medialiveSpecialFeature                          string
	_medialiveSpecialRouterSettings                   string
	_medialiveSrtSettings                             string
	_medialiveStart                                   string
	_medialiveState                                   string
	_medialiveStateFilter                             string
	_medialiveStatistic                               string
	_medialiveTagKeys                                 []string
	_medialiveTags                                    string
	_medialiveTargetCustomerId                        string
	_medialiveTargetRegion                            string
	_medialiveTargetResourceType                      string
	_medialiveThreshold                               string
	_medialiveThumbnailType                           string
	_medialiveTransferMessage                         string
	_medialiveTransferType                            string
	_medialiveTreatMissingData                        string
	_medialiveType                                    string
	_medialiveUhdDeviceSettings                       string
	_medialiveVideoQuality                            string
	_medialiveVpc                                     string
	_medialiveWhitelistRules                          string
)

// Accept an incoming input device transfer. The ownership of the device will
// transfer to your AWS account.
func medialive_AcceptInputDeviceTransfer(cfg aws.Config, client *medialive.Client) {
	input := &medialive.AcceptInputDeviceTransferInput{
		// InputDeviceId: *string, // Required
	}

	if len(_medialiveInputDeviceId) > 0 {
		input.InputDeviceId = aws.String(_medialiveInputDeviceId)
	}

	if resp, err := client.AcceptInputDeviceTransfer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts delete of resources.
func medialive_BatchDelete(cfg aws.Config, client *medialive.Client) {
	input := &medialive.BatchDeleteInput{}

	if len(_medialiveChannelIds) > 0 {
		input.ChannelIds = append([]string(nil), _medialiveChannelIds...)
	}
	if len(_medialiveInputIds) > 0 {
		input.InputIds = append([]string(nil), _medialiveInputIds...)
	}
	if len(_medialiveInputSecurityGroupIds) > 0 {
		input.InputSecurityGroupIds = append([]string(nil), _medialiveInputSecurityGroupIds...)
	}
	if len(_medialiveMultiplexIds) > 0 {
		input.MultiplexIds = append([]string(nil), _medialiveMultiplexIds...)
	}

	if resp, err := client.BatchDelete(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts existing resources
func medialive_BatchStart(cfg aws.Config, client *medialive.Client) {
	input := &medialive.BatchStartInput{}

	if len(_medialiveChannelIds) > 0 {
		input.ChannelIds = append([]string(nil), _medialiveChannelIds...)
	}
	if len(_medialiveMultiplexIds) > 0 {
		input.MultiplexIds = append([]string(nil), _medialiveMultiplexIds...)
	}

	if resp, err := client.BatchStart(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops running resources
func medialive_BatchStop(cfg aws.Config, client *medialive.Client) {
	input := &medialive.BatchStopInput{}

	if len(_medialiveChannelIds) > 0 {
		input.ChannelIds = append([]string(nil), _medialiveChannelIds...)
	}
	if len(_medialiveMultiplexIds) > 0 {
		input.MultiplexIds = append([]string(nil), _medialiveMultiplexIds...)
	}

	if resp, err := client.BatchStop(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update a channel schedule
func medialive_BatchUpdateSchedule(cfg aws.Config, client *medialive.Client) {
	input := &medialive.BatchUpdateScheduleInput{
		// ChannelId: *string, // Required
	}

	if len(_medialiveChannelId) > 0 {
		input.ChannelId = aws.String(_medialiveChannelId)
	}
	if len(_medialiveCreates) > 0 {
		if err := assignInputField(input, "Creates", _medialiveCreates); err != nil {
			log.Errorf("invalid --creates: %s", err.Error())
			return
		}
	}
	if len(_medialiveDeletes) > 0 {
		if err := assignInputField(input, "Deletes", _medialiveDeletes); err != nil {
			log.Errorf("invalid --deletes: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchUpdateSchedule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancel an input device transfer that you have requested.
func medialive_CancelInputDeviceTransfer(cfg aws.Config, client *medialive.Client) {
	input := &medialive.CancelInputDeviceTransferInput{
		// InputDeviceId: *string, // Required
	}

	if len(_medialiveInputDeviceId) > 0 {
		input.InputDeviceId = aws.String(_medialiveInputDeviceId)
	}

	if resp, err := client.CancelInputDeviceTransfer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Send a request to claim an AWS Elemental device that you have purchased from a
// third-party vendor. After the request succeeds, you will own the device.
func medialive_ClaimDevice(cfg aws.Config, client *medialive.Client) {
	input := &medialive.ClaimDeviceInput{}

	if len(_medialiveId) > 0 {
		input.Id = aws.String(_medialiveId)
	}

	if resp, err := client.ClaimDevice(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new channel
func medialive_CreateChannel(cfg aws.Config, client *medialive.Client) {
	input := &medialive.CreateChannelInput{}

	if len(_medialiveAnywhereSettings) > 0 {
		if err := assignInputField(input, "AnywhereSettings", _medialiveAnywhereSettings); err != nil {
			log.Errorf("invalid --anywhere-settings: %s", err.Error())
			return
		}
	}
	if len(_medialiveCdiInputSpecification) > 0 {
		if err := assignInputField(input, "CdiInputSpecification", _medialiveCdiInputSpecification); err != nil {
			log.Errorf("invalid --cdi-input-specification: %s", err.Error())
			return
		}
	}
	if len(_medialiveChannelClass) > 0 {
		if err := assignInputField(input, "ChannelClass", _medialiveChannelClass); err != nil {
			log.Errorf("invalid --channel-class: %s", err.Error())
			return
		}
	}
	if len(_medialiveChannelEngineVersion) > 0 {
		if err := assignInputField(input, "ChannelEngineVersion", _medialiveChannelEngineVersion); err != nil {
			log.Errorf("invalid --channel-engine-version: %s", err.Error())
			return
		}
	}
	if len(_medialiveChannelSecurityGroups) > 0 {
		input.ChannelSecurityGroups = append([]string(nil), _medialiveChannelSecurityGroups...)
	}
	if len(_medialiveDestinations) > 0 {
		if err := assignInputField(input, "Destinations", _medialiveDestinations); err != nil {
			log.Errorf("invalid --destinations: %s", err.Error())
			return
		}
	}
	if len(_medialiveDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _medialiveDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}
	if len(_medialiveEncoderSettings) > 0 {
		if err := assignInputField(input, "EncoderSettings", _medialiveEncoderSettings); err != nil {
			log.Errorf("invalid --encoder-settings: %s", err.Error())
			return
		}
	}
	if len(_medialiveInferenceSettings) > 0 {
		if err := assignInputField(input, "InferenceSettings", _medialiveInferenceSettings); err != nil {
			log.Errorf("invalid --inference-settings: %s", err.Error())
			return
		}
	}
	if len(_medialiveInputAttachments) > 0 {
		if err := assignInputField(input, "InputAttachments", _medialiveInputAttachments); err != nil {
			log.Errorf("invalid --input-attachments: %s", err.Error())
			return
		}
	}
	if len(_medialiveInputSpecification) > 0 {
		if err := assignInputField(input, "InputSpecification", _medialiveInputSpecification); err != nil {
			log.Errorf("invalid --input-specification: %s", err.Error())
			return
		}
	}
	if len(_medialiveLinkedChannelSettings) > 0 {
		if err := assignInputField(input, "LinkedChannelSettings", _medialiveLinkedChannelSettings); err != nil {
			log.Errorf("invalid --linked-channel-settings: %s", err.Error())
			return
		}
	}
	if len(_medialiveLogLevel) > 0 {
		if err := assignInputField(input, "LogLevel", _medialiveLogLevel); err != nil {
			log.Errorf("invalid --log-level: %s", err.Error())
			return
		}
	}
	if len(_medialiveMaintenance) > 0 {
		if err := assignInputField(input, "Maintenance", _medialiveMaintenance); err != nil {
			log.Errorf("invalid --maintenance: %s", err.Error())
			return
		}
	}
	if len(_medialiveName) > 0 {
		input.Name = aws.String(_medialiveName)
	}
	if len(_medialiveRequestId) > 0 {
		input.RequestId = aws.String(_medialiveRequestId)
	}
	if len(_medialiveReserved) > 0 {
		input.Reserved = aws.String(_medialiveReserved)
	}
	if len(_medialiveRoleArn) > 0 {
		input.RoleArn = aws.String(_medialiveRoleArn)
	}
	if len(_medialiveTags) > 0 {
		if err := assignInputField(input, "Tags", _medialiveTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_medialiveVpc) > 0 {
		if err := assignInputField(input, "Vpc", _medialiveVpc); err != nil {
			log.Errorf("invalid --vpc: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a ChannelPlacementGroup in the specified Cluster. As part of the create
// operation, you specify the Nodes to attach the group to.After you create a
// ChannelPlacementGroup, you add Channels to the group (you do this by modifying
// the Channels to add them to a specific group). You now have an association of
// Channels to ChannelPlacementGroup, and ChannelPlacementGroup to Nodes. This
// association means that all the Channels in the group are able to run on any of
// the Nodes associated with the group.
func medialive_CreateChannelPlacementGroup(cfg aws.Config, client *medialive.Client) {
	input := &medialive.CreateChannelPlacementGroupInput{
		// ClusterId: *string, // Required
	}

	if len(_medialiveClusterId) > 0 {
		input.ClusterId = aws.String(_medialiveClusterId)
	}
	if len(_medialiveName) > 0 {
		input.Name = aws.String(_medialiveName)
	}
	if len(_medialiveNodes) > 0 {
		input.Nodes = append([]string(nil), _medialiveNodes...)
	}
	if len(_medialiveRequestId) > 0 {
		input.RequestId = aws.String(_medialiveRequestId)
	}
	if len(_medialiveTags) > 0 {
		if err := assignInputField(input, "Tags", _medialiveTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateChannelPlacementGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a cloudwatch alarm template to dynamically generate cloudwatch metric
// alarms on targeted resource types.
func medialive_CreateCloudWatchAlarmTemplate(cfg aws.Config, client *medialive.Client) {
	input := &medialive.CreateCloudWatchAlarmTemplateInput{
		// ComparisonOperator: types.CloudWatchAlarmTemplateComparisonOperator, // Required
		// EvaluationPeriods: *int32, // Required
		// GroupIdentifier: *string, // Required
		// MetricName: *string, // Required
		// Name: *string, // Required
		// Period: *int32, // Required
		// Statistic: types.CloudWatchAlarmTemplateStatistic, // Required
		// TargetResourceType: types.CloudWatchAlarmTemplateTargetResourceType, // Required
		// Threshold: *float64, // Required
		// TreatMissingData: types.CloudWatchAlarmTemplateTreatMissingData, // Required
	}

	if len(_medialiveComparisonOperator) > 0 {
		if err := assignInputField(input, "ComparisonOperator", _medialiveComparisonOperator); err != nil {
			log.Errorf("invalid --comparison-operator: %s", err.Error())
			return
		}
	}
	if len(_medialiveEvaluationPeriods) > 0 {
		if err := assignInputField(input, "EvaluationPeriods", _medialiveEvaluationPeriods); err != nil {
			log.Errorf("invalid --evaluation-periods: %s", err.Error())
			return
		}
	}
	if len(_medialiveGroupIdentifier) > 0 {
		input.GroupIdentifier = aws.String(_medialiveGroupIdentifier)
	}
	if len(_medialiveMetricName) > 0 {
		input.MetricName = aws.String(_medialiveMetricName)
	}
	if len(_medialiveName) > 0 {
		input.Name = aws.String(_medialiveName)
	}
	if len(_medialivePeriod) > 0 {
		if err := assignInputField(input, "Period", _medialivePeriod); err != nil {
			log.Errorf("invalid --period: %s", err.Error())
			return
		}
	}
	if len(_medialiveStatistic) > 0 {
		if err := assignInputField(input, "Statistic", _medialiveStatistic); err != nil {
			log.Errorf("invalid --statistic: %s", err.Error())
			return
		}
	}
	if len(_medialiveTargetResourceType) > 0 {
		if err := assignInputField(input, "TargetResourceType", _medialiveTargetResourceType); err != nil {
			log.Errorf("invalid --target-resource-type: %s", err.Error())
			return
		}
	}
	if len(_medialiveThreshold) > 0 {
		if err := assignInputField(input, "Threshold", _medialiveThreshold); err != nil {
			log.Errorf("invalid --threshold: %s", err.Error())
			return
		}
	}
	if len(_medialiveTreatMissingData) > 0 {
		if err := assignInputField(input, "TreatMissingData", _medialiveTreatMissingData); err != nil {
			log.Errorf("invalid --treat-missing-data: %s", err.Error())
			return
		}
	}
	if len(_medialiveDatapointsToAlarm) > 0 {
		if err := assignInputField(input, "DatapointsToAlarm", _medialiveDatapointsToAlarm); err != nil {
			log.Errorf("invalid --datapoints-to-alarm: %s", err.Error())
			return
		}
	}
	if len(_medialiveDescription) > 0 {
		input.Description = aws.String(_medialiveDescription)
	}
	if len(_medialiveRequestId) > 0 {
		input.RequestId = aws.String(_medialiveRequestId)
	}
	if len(_medialiveTags) > 0 {
		if err := assignInputField(input, "Tags", _medialiveTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCloudWatchAlarmTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a cloudwatch alarm template group to group your cloudwatch alarm
// templates and to attach to signal maps for dynamically creating alarms.
func medialive_CreateCloudWatchAlarmTemplateGroup(cfg aws.Config, client *medialive.Client) {
	input := &medialive.CreateCloudWatchAlarmTemplateGroupInput{
		// Name: *string, // Required
	}

	if len(_medialiveName) > 0 {
		input.Name = aws.String(_medialiveName)
	}
	if len(_medialiveDescription) > 0 {
		input.Description = aws.String(_medialiveDescription)
	}
	if len(_medialiveRequestId) > 0 {
		input.RequestId = aws.String(_medialiveRequestId)
	}
	if len(_medialiveTags) > 0 {
		if err := assignInputField(input, "Tags", _medialiveTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCloudWatchAlarmTemplateGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a new Cluster.
func medialive_CreateCluster(cfg aws.Config, client *medialive.Client) {
	input := &medialive.CreateClusterInput{}

	if len(_medialiveClusterType) > 0 {
		if err := assignInputField(input, "ClusterType", _medialiveClusterType); err != nil {
			log.Errorf("invalid --cluster-type: %s", err.Error())
			return
		}
	}
	if len(_medialiveInstanceRoleArn) > 0 {
		input.InstanceRoleArn = aws.String(_medialiveInstanceRoleArn)
	}
	if len(_medialiveName) > 0 {
		input.Name = aws.String(_medialiveName)
	}
	if len(_medialiveNetworkSettings) > 0 {
		if err := assignInputField(input, "NetworkSettings", _medialiveNetworkSettings); err != nil {
			log.Errorf("invalid --network-settings: %s", err.Error())
			return
		}
	}
	if len(_medialiveRequestId) > 0 {
		input.RequestId = aws.String(_medialiveRequestId)
	}
	if len(_medialiveTags) > 0 {
		if err := assignInputField(input, "Tags", _medialiveTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an eventbridge rule template to monitor events and send notifications
// to your targeted resources.
func medialive_CreateEventBridgeRuleTemplate(cfg aws.Config, client *medialive.Client) {
	input := &medialive.CreateEventBridgeRuleTemplateInput{
		// EventType: types.EventBridgeRuleTemplateEventType, // Required
		// GroupIdentifier: *string, // Required
		// Name: *string, // Required
	}

	if len(_medialiveEventType) > 0 {
		if err := assignInputField(input, "EventType", _medialiveEventType); err != nil {
			log.Errorf("invalid --event-type: %s", err.Error())
			return
		}
	}
	if len(_medialiveGroupIdentifier) > 0 {
		input.GroupIdentifier = aws.String(_medialiveGroupIdentifier)
	}
	if len(_medialiveName) > 0 {
		input.Name = aws.String(_medialiveName)
	}
	if len(_medialiveDescription) > 0 {
		input.Description = aws.String(_medialiveDescription)
	}
	if len(_medialiveEventTargets) > 0 {
		if err := assignInputField(input, "EventTargets", _medialiveEventTargets); err != nil {
			log.Errorf("invalid --event-targets: %s", err.Error())
			return
		}
	}
	if len(_medialiveRequestId) > 0 {
		input.RequestId = aws.String(_medialiveRequestId)
	}
	if len(_medialiveTags) > 0 {
		if err := assignInputField(input, "Tags", _medialiveTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateEventBridgeRuleTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an eventbridge rule template group to group your eventbridge rule
// templates and to attach to signal maps for dynamically creating notification
// rules.
func medialive_CreateEventBridgeRuleTemplateGroup(cfg aws.Config, client *medialive.Client) {
	input := &medialive.CreateEventBridgeRuleTemplateGroupInput{
		// Name: *string, // Required
	}

	if len(_medialiveName) > 0 {
		input.Name = aws.String(_medialiveName)
	}
	if len(_medialiveDescription) > 0 {
		input.Description = aws.String(_medialiveDescription)
	}
	if len(_medialiveRequestId) > 0 {
		input.RequestId = aws.String(_medialiveRequestId)
	}
	if len(_medialiveTags) > 0 {
		if err := assignInputField(input, "Tags", _medialiveTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateEventBridgeRuleTemplateGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create an input
func medialive_CreateInput(cfg aws.Config, client *medialive.Client) {
	input := &medialive.CreateInputInput{}

	if len(_medialiveDestinations) > 0 {
		if err := assignInputField(input, "Destinations", _medialiveDestinations); err != nil {
			log.Errorf("invalid --destinations: %s", err.Error())
			return
		}
	}
	if len(_medialiveInputDevices) > 0 {
		if err := assignInputField(input, "InputDevices", _medialiveInputDevices); err != nil {
			log.Errorf("invalid --input-devices: %s", err.Error())
			return
		}
	}
	if len(_medialiveInputNetworkLocation) > 0 {
		if err := assignInputField(input, "InputNetworkLocation", _medialiveInputNetworkLocation); err != nil {
			log.Errorf("invalid --input-network-location: %s", err.Error())
			return
		}
	}
	if len(_medialiveInputSecurityGroups) > 0 {
		input.InputSecurityGroups = append([]string(nil), _medialiveInputSecurityGroups...)
	}
	if len(_medialiveMediaConnectFlows) > 0 {
		if err := assignInputField(input, "MediaConnectFlows", _medialiveMediaConnectFlows); err != nil {
			log.Errorf("invalid --media-connect-flows: %s", err.Error())
			return
		}
	}
	if len(_medialiveMulticastSettings) > 0 {
		if err := assignInputField(input, "MulticastSettings", _medialiveMulticastSettings); err != nil {
			log.Errorf("invalid --multicast-settings: %s", err.Error())
			return
		}
	}
	if len(_medialiveName) > 0 {
		input.Name = aws.String(_medialiveName)
	}
	if len(_medialiveRequestId) > 0 {
		input.RequestId = aws.String(_medialiveRequestId)
	}
	if len(_medialiveRoleArn) > 0 {
		input.RoleArn = aws.String(_medialiveRoleArn)
	}
	if len(_medialiveRouterSettings) > 0 {
		if err := assignInputField(input, "RouterSettings", _medialiveRouterSettings); err != nil {
			log.Errorf("invalid --router-settings: %s", err.Error())
			return
		}
	}
	if len(_medialiveSdiSources) > 0 {
		input.SdiSources = append([]string(nil), _medialiveSdiSources...)
	}
	if len(_medialiveSmpte2110ReceiverGroupSettings) > 0 {
		if err := assignInputField(input, "Smpte2110ReceiverGroupSettings", _medialiveSmpte2110ReceiverGroupSettings); err != nil {
			log.Errorf("invalid --smpte2110-receiver-group-settings: %s", err.Error())
			return
		}
	}
	if len(_medialiveSources) > 0 {
		if err := assignInputField(input, "Sources", _medialiveSources); err != nil {
			log.Errorf("invalid --sources: %s", err.Error())
			return
		}
	}
	if len(_medialiveSrtSettings) > 0 {
		if err := assignInputField(input, "SrtSettings", _medialiveSrtSettings); err != nil {
			log.Errorf("invalid --srt-settings: %s", err.Error())
			return
		}
	}
	if len(_medialiveTags) > 0 {
		if err := assignInputField(input, "Tags", _medialiveTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_medialiveType) > 0 {
		if err := assignInputField(input, "Type", _medialiveType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_medialiveVpc) > 0 {
		if err := assignInputField(input, "Vpc", _medialiveVpc); err != nil {
			log.Errorf("invalid --vpc: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateInput(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a Input Security Group
func medialive_CreateInputSecurityGroup(cfg aws.Config, client *medialive.Client) {
	input := &medialive.CreateInputSecurityGroupInput{}

	if len(_medialiveTags) > 0 {
		if err := assignInputField(input, "Tags", _medialiveTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_medialiveWhitelistRules) > 0 {
		if err := assignInputField(input, "WhitelistRules", _medialiveWhitelistRules); err != nil {
			log.Errorf("invalid --whitelist-rules: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateInputSecurityGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a new multiplex.
func medialive_CreateMultiplex(cfg aws.Config, client *medialive.Client) {
	input := &medialive.CreateMultiplexInput{
		// AvailabilityZones: []string, // Required
		// MultiplexSettings: *types.MultiplexSettings, // Required
		// Name: *string, // Required
		// RequestId: *string, // Required
	}

	if len(_medialiveAvailabilityZones) > 0 {
		input.AvailabilityZones = append([]string(nil), _medialiveAvailabilityZones...)
	}
	if len(_medialiveMultiplexSettings) > 0 {
		if err := assignInputField(input, "MultiplexSettings", _medialiveMultiplexSettings); err != nil {
			log.Errorf("invalid --multiplex-settings: %s", err.Error())
			return
		}
	}
	if len(_medialiveName) > 0 {
		input.Name = aws.String(_medialiveName)
	}
	if len(_medialiveRequestId) > 0 {
		input.RequestId = aws.String(_medialiveRequestId)
	}
	if len(_medialiveTags) > 0 {
		if err := assignInputField(input, "Tags", _medialiveTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateMultiplex(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a new program in the multiplex.
func medialive_CreateMultiplexProgram(cfg aws.Config, client *medialive.Client) {
	input := &medialive.CreateMultiplexProgramInput{
		// MultiplexId: *string, // Required
		// MultiplexProgramSettings: *types.MultiplexProgramSettings, // Required
		// ProgramName: *string, // Required
		// RequestId: *string, // Required
	}

	if len(_medialiveMultiplexId) > 0 {
		input.MultiplexId = aws.String(_medialiveMultiplexId)
	}
	if len(_medialiveMultiplexProgramSettings) > 0 {
		if err := assignInputField(input, "MultiplexProgramSettings", _medialiveMultiplexProgramSettings); err != nil {
			log.Errorf("invalid --multiplex-program-settings: %s", err.Error())
			return
		}
	}
	if len(_medialiveProgramName) > 0 {
		input.ProgramName = aws.String(_medialiveProgramName)
	}
	if len(_medialiveRequestId) > 0 {
		input.RequestId = aws.String(_medialiveRequestId)
	}

	if resp, err := client.CreateMultiplexProgram(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create as many Networks as you need. You will associate one or more Clusters
// with each Network.Each Network provides MediaLive Anywhere with required
// information about the network in your organization that you are using for video
// encoding using MediaLive.
func medialive_CreateNetwork(cfg aws.Config, client *medialive.Client) {
	input := &medialive.CreateNetworkInput{}

	if len(_medialiveIpPools) > 0 {
		if err := assignInputField(input, "IpPools", _medialiveIpPools); err != nil {
			log.Errorf("invalid --ip-pools: %s", err.Error())
			return
		}
	}
	if len(_medialiveName) > 0 {
		input.Name = aws.String(_medialiveName)
	}
	if len(_medialiveRequestId) > 0 {
		input.RequestId = aws.String(_medialiveRequestId)
	}
	if len(_medialiveRoutes) > 0 {
		if err := assignInputField(input, "Routes", _medialiveRoutes); err != nil {
			log.Errorf("invalid --routes: %s", err.Error())
			return
		}
	}
	if len(_medialiveTags) > 0 {
		if err := assignInputField(input, "Tags", _medialiveTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateNetwork(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a Node in the specified Cluster. You can also create Nodes using the
// CreateNodeRegistrationScript. Note that you can't move a Node to another
// Cluster.
func medialive_CreateNode(cfg aws.Config, client *medialive.Client) {
	input := &medialive.CreateNodeInput{
		// ClusterId: *string, // Required
	}

	if len(_medialiveClusterId) > 0 {
		input.ClusterId = aws.String(_medialiveClusterId)
	}
	if len(_medialiveName) > 0 {
		input.Name = aws.String(_medialiveName)
	}
	if len(_medialiveNodeInterfaceMappings) > 0 {
		if err := assignInputField(input, "NodeInterfaceMappings", _medialiveNodeInterfaceMappings); err != nil {
			log.Errorf("invalid --node-interface-mappings: %s", err.Error())
			return
		}
	}
	if len(_medialiveRequestId) > 0 {
		input.RequestId = aws.String(_medialiveRequestId)
	}
	if len(_medialiveRole) > 0 {
		if err := assignInputField(input, "Role", _medialiveRole); err != nil {
			log.Errorf("invalid --role: %s", err.Error())
			return
		}
	}
	if len(_medialiveTags) > 0 {
		if err := assignInputField(input, "Tags", _medialiveTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateNode(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create the Register Node script for all the nodes intended for a specific
// Cluster. You will then run the script on each hardware unit that is intended for
// that Cluster. The script creates a Node in the specified Cluster. It then binds
// the Node to this hardware unit, and activates the node hardware for use with
// MediaLive Anywhere.
func medialive_CreateNodeRegistrationScript(cfg aws.Config, client *medialive.Client) {
	input := &medialive.CreateNodeRegistrationScriptInput{
		// ClusterId: *string, // Required
	}

	if len(_medialiveClusterId) > 0 {
		input.ClusterId = aws.String(_medialiveClusterId)
	}
	if len(_medialiveId) > 0 {
		input.Id = aws.String(_medialiveId)
	}
	if len(_medialiveName) > 0 {
		input.Name = aws.String(_medialiveName)
	}
	if len(_medialiveNodeInterfaceMappings) > 0 {
		if err := assignInputField(input, "NodeInterfaceMappings", _medialiveNodeInterfaceMappings); err != nil {
			log.Errorf("invalid --node-interface-mappings: %s", err.Error())
			return
		}
	}
	if len(_medialiveRequestId) > 0 {
		input.RequestId = aws.String(_medialiveRequestId)
	}
	if len(_medialiveRole) > 0 {
		if err := assignInputField(input, "Role", _medialiveRole); err != nil {
			log.Errorf("invalid --role: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateNodeRegistrationScript(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a partner input
func medialive_CreatePartnerInput(cfg aws.Config, client *medialive.Client) {
	input := &medialive.CreatePartnerInputInput{
		// InputId: *string, // Required
	}

	if len(_medialiveInputId) > 0 {
		input.InputId = aws.String(_medialiveInputId)
	}
	if len(_medialiveRequestId) > 0 {
		input.RequestId = aws.String(_medialiveRequestId)
	}
	if len(_medialiveTags) > 0 {
		if err := assignInputField(input, "Tags", _medialiveTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePartnerInput(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create an SdiSource for each video source that uses the SDI protocol. You will
// reference the SdiSource when you create an SDI input in MediaLive. You will also
// reference it in an SdiSourceMapping, in order to create a connection between the
// logical SdiSource and the physical SDI card and port that the physical SDI
// source uses.
func medialive_CreateSdiSource(cfg aws.Config, client *medialive.Client) {
	input := &medialive.CreateSdiSourceInput{}

	if len(_medialiveMode) > 0 {
		if err := assignInputField(input, "Mode", _medialiveMode); err != nil {
			log.Errorf("invalid --mode: %s", err.Error())
			return
		}
	}
	if len(_medialiveName) > 0 {
		input.Name = aws.String(_medialiveName)
	}
	if len(_medialiveRequestId) > 0 {
		input.RequestId = aws.String(_medialiveRequestId)
	}
	if len(_medialiveTags) > 0 {
		if err := assignInputField(input, "Tags", _medialiveTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_medialiveType) > 0 {
		if err := assignInputField(input, "Type", _medialiveType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSdiSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initiates the creation of a new signal map. Will discover a new
// mediaResourceMap based on the provided discoveryEntryPointArn.
func medialive_CreateSignalMap(cfg aws.Config, client *medialive.Client) {
	input := &medialive.CreateSignalMapInput{
		// DiscoveryEntryPointArn: *string, // Required
		// Name: *string, // Required
	}

	if len(_medialiveDiscoveryEntryPointArn) > 0 {
		input.DiscoveryEntryPointArn = aws.String(_medialiveDiscoveryEntryPointArn)
	}
	if len(_medialiveName) > 0 {
		input.Name = aws.String(_medialiveName)
	}
	if len(_medialiveCloudWatchAlarmTemplateGroupIdentifiers) > 0 {
		input.CloudWatchAlarmTemplateGroupIdentifiers = append([]string(nil), _medialiveCloudWatchAlarmTemplateGroupIdentifiers...)
	}
	if len(_medialiveDescription) > 0 {
		input.Description = aws.String(_medialiveDescription)
	}
	if len(_medialiveEventBridgeRuleTemplateGroupIdentifiers) > 0 {
		input.EventBridgeRuleTemplateGroupIdentifiers = append([]string(nil), _medialiveEventBridgeRuleTemplateGroupIdentifiers...)
	}
	if len(_medialiveRequestId) > 0 {
		input.RequestId = aws.String(_medialiveRequestId)
	}
	if len(_medialiveTags) > 0 {
		if err := assignInputField(input, "Tags", _medialiveTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSignalMap(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create tags for a resource
func medialive_CreateTags(cfg aws.Config, client *medialive.Client) {
	input := &medialive.CreateTagsInput{
		// ResourceArn: *string, // Required
	}

	if len(_medialiveResourceArn) > 0 {
		input.ResourceArn = aws.String(_medialiveResourceArn)
	}
	if len(_medialiveTags) > 0 {
		if err := assignInputField(input, "Tags", _medialiveTags); err != nil {
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

// Starts deletion of channel. The associated outputs are also deleted.
func medialive_DeleteChannel(cfg aws.Config, client *medialive.Client) {
	input := &medialive.DeleteChannelInput{
		// ChannelId: *string, // Required
	}

	if len(_medialiveChannelId) > 0 {
		input.ChannelId = aws.String(_medialiveChannelId)
	}

	if resp, err := client.DeleteChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete the specified ChannelPlacementGroup that exists in the specified Cluster.
func medialive_DeleteChannelPlacementGroup(cfg aws.Config, client *medialive.Client) {
	input := &medialive.DeleteChannelPlacementGroupInput{
		// ChannelPlacementGroupId: *string, // Required
		// ClusterId: *string, // Required
	}

	if len(_medialiveChannelPlacementGroupId) > 0 {
		input.ChannelPlacementGroupId = aws.String(_medialiveChannelPlacementGroupId)
	}
	if len(_medialiveClusterId) > 0 {
		input.ClusterId = aws.String(_medialiveClusterId)
	}

	if resp, err := client.DeleteChannelPlacementGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a cloudwatch alarm template.
func medialive_DeleteCloudWatchAlarmTemplate(cfg aws.Config, client *medialive.Client) {
	input := &medialive.DeleteCloudWatchAlarmTemplateInput{
		// Identifier: *string, // Required
	}

	if len(_medialiveIdentifier) > 0 {
		input.Identifier = aws.String(_medialiveIdentifier)
	}

	if resp, err := client.DeleteCloudWatchAlarmTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a cloudwatch alarm template group. You must detach this group from all
// signal maps and ensure its existing templates are moved to another group or
// deleted.
func medialive_DeleteCloudWatchAlarmTemplateGroup(cfg aws.Config, client *medialive.Client) {
	input := &medialive.DeleteCloudWatchAlarmTemplateGroupInput{
		// Identifier: *string, // Required
	}

	if len(_medialiveIdentifier) > 0 {
		input.Identifier = aws.String(_medialiveIdentifier)
	}

	if resp, err := client.DeleteCloudWatchAlarmTemplateGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a Cluster. The Cluster must be idle.
func medialive_DeleteCluster(cfg aws.Config, client *medialive.Client) {
	input := &medialive.DeleteClusterInput{
		// ClusterId: *string, // Required
	}

	if len(_medialiveClusterId) > 0 {
		input.ClusterId = aws.String(_medialiveClusterId)
	}

	if resp, err := client.DeleteCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an eventbridge rule template.
func medialive_DeleteEventBridgeRuleTemplate(cfg aws.Config, client *medialive.Client) {
	input := &medialive.DeleteEventBridgeRuleTemplateInput{
		// Identifier: *string, // Required
	}

	if len(_medialiveIdentifier) > 0 {
		input.Identifier = aws.String(_medialiveIdentifier)
	}

	if resp, err := client.DeleteEventBridgeRuleTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an eventbridge rule template group. You must detach this group from all
// signal maps and ensure its existing templates are moved to another group or
// deleted.
func medialive_DeleteEventBridgeRuleTemplateGroup(cfg aws.Config, client *medialive.Client) {
	input := &medialive.DeleteEventBridgeRuleTemplateGroupInput{
		// Identifier: *string, // Required
	}

	if len(_medialiveIdentifier) > 0 {
		input.Identifier = aws.String(_medialiveIdentifier)
	}

	if resp, err := client.DeleteEventBridgeRuleTemplateGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the input end point
func medialive_DeleteInput(cfg aws.Config, client *medialive.Client) {
	input := &medialive.DeleteInputInput{
		// InputId: *string, // Required
	}

	if len(_medialiveInputId) > 0 {
		input.InputId = aws.String(_medialiveInputId)
	}

	if resp, err := client.DeleteInput(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Input Security Group
func medialive_DeleteInputSecurityGroup(cfg aws.Config, client *medialive.Client) {
	input := &medialive.DeleteInputSecurityGroupInput{
		// InputSecurityGroupId: *string, // Required
	}

	if len(_medialiveInputSecurityGroupId) > 0 {
		input.InputSecurityGroupId = aws.String(_medialiveInputSecurityGroupId)
	}

	if resp, err := client.DeleteInputSecurityGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a multiplex. The multiplex must be idle.
func medialive_DeleteMultiplex(cfg aws.Config, client *medialive.Client) {
	input := &medialive.DeleteMultiplexInput{
		// MultiplexId: *string, // Required
	}

	if len(_medialiveMultiplexId) > 0 {
		input.MultiplexId = aws.String(_medialiveMultiplexId)
	}

	if resp, err := client.DeleteMultiplex(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a program from a multiplex.
func medialive_DeleteMultiplexProgram(cfg aws.Config, client *medialive.Client) {
	input := &medialive.DeleteMultiplexProgramInput{
		// MultiplexId: *string, // Required
		// ProgramName: *string, // Required
	}

	if len(_medialiveMultiplexId) > 0 {
		input.MultiplexId = aws.String(_medialiveMultiplexId)
	}
	if len(_medialiveProgramName) > 0 {
		input.ProgramName = aws.String(_medialiveProgramName)
	}

	if resp, err := client.DeleteMultiplexProgram(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a Network. The Network must have no resources associated with it.
func medialive_DeleteNetwork(cfg aws.Config, client *medialive.Client) {
	input := &medialive.DeleteNetworkInput{
		// NetworkId: *string, // Required
	}

	if len(_medialiveNetworkId) > 0 {
		input.NetworkId = aws.String(_medialiveNetworkId)
	}

	if resp, err := client.DeleteNetwork(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a Node. The Node must be IDLE.
func medialive_DeleteNode(cfg aws.Config, client *medialive.Client) {
	input := &medialive.DeleteNodeInput{
		// ClusterId: *string, // Required
		// NodeId: *string, // Required
	}

	if len(_medialiveClusterId) > 0 {
		input.ClusterId = aws.String(_medialiveClusterId)
	}
	if len(_medialiveNodeId) > 0 {
		input.NodeId = aws.String(_medialiveNodeId)
	}

	if resp, err := client.DeleteNode(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete an expired reservation.
func medialive_DeleteReservation(cfg aws.Config, client *medialive.Client) {
	input := &medialive.DeleteReservationInput{
		// ReservationId: *string, // Required
	}

	if len(_medialiveReservationId) > 0 {
		input.ReservationId = aws.String(_medialiveReservationId)
	}

	if resp, err := client.DeleteReservation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete all schedule actions on a channel.
func medialive_DeleteSchedule(cfg aws.Config, client *medialive.Client) {
	input := &medialive.DeleteScheduleInput{
		// ChannelId: *string, // Required
	}

	if len(_medialiveChannelId) > 0 {
		input.ChannelId = aws.String(_medialiveChannelId)
	}

	if resp, err := client.DeleteSchedule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete an SdiSource. The SdiSource must not be part of any SidSourceMapping and
// must not be attached to any input.
func medialive_DeleteSdiSource(cfg aws.Config, client *medialive.Client) {
	input := &medialive.DeleteSdiSourceInput{
		// SdiSourceId: *string, // Required
	}

	if len(_medialiveSdiSourceId) > 0 {
		input.SdiSourceId = aws.String(_medialiveSdiSourceId)
	}

	if resp, err := client.DeleteSdiSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified signal map.
func medialive_DeleteSignalMap(cfg aws.Config, client *medialive.Client) {
	input := &medialive.DeleteSignalMapInput{
		// Identifier: *string, // Required
	}

	if len(_medialiveIdentifier) > 0 {
		input.Identifier = aws.String(_medialiveIdentifier)
	}

	if resp, err := client.DeleteSignalMap(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes tags for a resource
func medialive_DeleteTags(cfg aws.Config, client *medialive.Client) {
	input := &medialive.DeleteTagsInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_medialiveResourceArn) > 0 {
		input.ResourceArn = aws.String(_medialiveResourceArn)
	}
	if len(_medialiveTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _medialiveTagKeys...)
	}

	if resp, err := client.DeleteTags(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describe account configuration
func medialive_DescribeAccountConfiguration(cfg aws.Config, client *medialive.Client) {
	input := &medialive.DescribeAccountConfigurationInput{}

	if resp, err := client.DescribeAccountConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets details about a channel
func medialive_DescribeChannel(cfg aws.Config, client *medialive.Client) {
	input := &medialive.DescribeChannelInput{
		// ChannelId: *string, // Required
	}

	if len(_medialiveChannelId) > 0 {
		input.ChannelId = aws.String(_medialiveChannelId)
	}

	if resp, err := client.DescribeChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get details about a ChannelPlacementGroup.
func medialive_DescribeChannelPlacementGroup(cfg aws.Config, client *medialive.Client) {
	input := &medialive.DescribeChannelPlacementGroupInput{
		// ChannelPlacementGroupId: *string, // Required
		// ClusterId: *string, // Required
	}

	if len(_medialiveChannelPlacementGroupId) > 0 {
		input.ChannelPlacementGroupId = aws.String(_medialiveChannelPlacementGroupId)
	}
	if len(_medialiveClusterId) > 0 {
		input.ClusterId = aws.String(_medialiveClusterId)
	}

	if resp, err := client.DescribeChannelPlacementGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get details about a Cluster.
func medialive_DescribeCluster(cfg aws.Config, client *medialive.Client) {
	input := &medialive.DescribeClusterInput{
		// ClusterId: *string, // Required
	}

	if len(_medialiveClusterId) > 0 {
		input.ClusterId = aws.String(_medialiveClusterId)
	}

	if resp, err := client.DescribeCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Produces details about an input
func medialive_DescribeInput(cfg aws.Config, client *medialive.Client) {
	input := &medialive.DescribeInputInput{
		// InputId: *string, // Required
	}

	if len(_medialiveInputId) > 0 {
		input.InputId = aws.String(_medialiveInputId)
	}

	if resp, err := client.DescribeInput(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the details for the input device
func medialive_DescribeInputDevice(cfg aws.Config, client *medialive.Client) {
	input := &medialive.DescribeInputDeviceInput{
		// InputDeviceId: *string, // Required
	}

	if len(_medialiveInputDeviceId) > 0 {
		input.InputDeviceId = aws.String(_medialiveInputDeviceId)
	}

	if resp, err := client.DescribeInputDevice(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the latest thumbnail data for the input device.
func medialive_DescribeInputDeviceThumbnail(cfg aws.Config, client *medialive.Client) {
	input := &medialive.DescribeInputDeviceThumbnailInput{
		// Accept: types.AcceptHeader, // Required
		// InputDeviceId: *string, // Required
	}

	if len(_medialiveAccept) > 0 {
		if err := assignInputField(input, "Accept", _medialiveAccept); err != nil {
			log.Errorf("invalid --accept: %s", err.Error())
			return
		}
	}
	if len(_medialiveInputDeviceId) > 0 {
		input.InputDeviceId = aws.String(_medialiveInputDeviceId)
	}

	if resp, err := client.DescribeInputDeviceThumbnail(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Produces a summary of an Input Security Group
func medialive_DescribeInputSecurityGroup(cfg aws.Config, client *medialive.Client) {
	input := &medialive.DescribeInputSecurityGroupInput{
		// InputSecurityGroupId: *string, // Required
	}

	if len(_medialiveInputSecurityGroupId) > 0 {
		input.InputSecurityGroupId = aws.String(_medialiveInputSecurityGroupId)
	}

	if resp, err := client.DescribeInputSecurityGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets details about a multiplex.
func medialive_DescribeMultiplex(cfg aws.Config, client *medialive.Client) {
	input := &medialive.DescribeMultiplexInput{
		// MultiplexId: *string, // Required
	}

	if len(_medialiveMultiplexId) > 0 {
		input.MultiplexId = aws.String(_medialiveMultiplexId)
	}

	if resp, err := client.DescribeMultiplex(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the details for a program in a multiplex.
func medialive_DescribeMultiplexProgram(cfg aws.Config, client *medialive.Client) {
	input := &medialive.DescribeMultiplexProgramInput{
		// MultiplexId: *string, // Required
		// ProgramName: *string, // Required
	}

	if len(_medialiveMultiplexId) > 0 {
		input.MultiplexId = aws.String(_medialiveMultiplexId)
	}
	if len(_medialiveProgramName) > 0 {
		input.ProgramName = aws.String(_medialiveProgramName)
	}

	if resp, err := client.DescribeMultiplexProgram(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get details about a Network.
func medialive_DescribeNetwork(cfg aws.Config, client *medialive.Client) {
	input := &medialive.DescribeNetworkInput{
		// NetworkId: *string, // Required
	}

	if len(_medialiveNetworkId) > 0 {
		input.NetworkId = aws.String(_medialiveNetworkId)
	}

	if resp, err := client.DescribeNetwork(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get details about a Node in the specified Cluster.
func medialive_DescribeNode(cfg aws.Config, client *medialive.Client) {
	input := &medialive.DescribeNodeInput{
		// ClusterId: *string, // Required
		// NodeId: *string, // Required
	}

	if len(_medialiveClusterId) > 0 {
		input.ClusterId = aws.String(_medialiveClusterId)
	}
	if len(_medialiveNodeId) > 0 {
		input.NodeId = aws.String(_medialiveNodeId)
	}

	if resp, err := client.DescribeNode(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get details for an offering.
func medialive_DescribeOffering(cfg aws.Config, client *medialive.Client) {
	input := &medialive.DescribeOfferingInput{
		// OfferingId: *string, // Required
	}

	if len(_medialiveOfferingId) > 0 {
		input.OfferingId = aws.String(_medialiveOfferingId)
	}

	if resp, err := client.DescribeOffering(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get details for a reservation.
func medialive_DescribeReservation(cfg aws.Config, client *medialive.Client) {
	input := &medialive.DescribeReservationInput{
		// ReservationId: *string, // Required
	}

	if len(_medialiveReservationId) > 0 {
		input.ReservationId = aws.String(_medialiveReservationId)
	}

	if resp, err := client.DescribeReservation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get a channel schedule
func medialive_DescribeSchedule(cfg aws.Config, client *medialive.Client) {
	input := &medialive.DescribeScheduleInput{
		// ChannelId: *string, // Required
	}

	if len(_medialiveChannelId) > 0 {
		input.ChannelId = aws.String(_medialiveChannelId)
	}
	if len(_medialiveMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _medialiveMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_medialiveNextToken) > 0 {
		input.NextToken = aws.String(_medialiveNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeSchedule(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*medialive.DescribeScheduleOutput
	p := medialive.NewDescribeSchedulePaginator(client, input)
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

// Gets details about a SdiSource.
func medialive_DescribeSdiSource(cfg aws.Config, client *medialive.Client) {
	input := &medialive.DescribeSdiSourceInput{
		// SdiSourceId: *string, // Required
	}

	if len(_medialiveSdiSourceId) > 0 {
		input.SdiSourceId = aws.String(_medialiveSdiSourceId)
	}

	if resp, err := client.DescribeSdiSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describe the latest thumbnails data.
func medialive_DescribeThumbnails(cfg aws.Config, client *medialive.Client) {
	input := &medialive.DescribeThumbnailsInput{
		// ChannelId: *string, // Required
		// PipelineId: *string, // Required
		// ThumbnailType: *string, // Required
	}

	if len(_medialiveChannelId) > 0 {
		input.ChannelId = aws.String(_medialiveChannelId)
	}
	if len(_medialivePipelineId) > 0 {
		input.PipelineId = aws.String(_medialivePipelineId)
	}
	if len(_medialiveThumbnailType) > 0 {
		input.ThumbnailType = aws.String(_medialiveThumbnailType)
	}

	if resp, err := client.DescribeThumbnails(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the specified cloudwatch alarm template.
func medialive_GetCloudWatchAlarmTemplate(cfg aws.Config, client *medialive.Client) {
	input := &medialive.GetCloudWatchAlarmTemplateInput{
		// Identifier: *string, // Required
	}

	if len(_medialiveIdentifier) > 0 {
		input.Identifier = aws.String(_medialiveIdentifier)
	}

	if resp, err := client.GetCloudWatchAlarmTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the specified cloudwatch alarm template group.
func medialive_GetCloudWatchAlarmTemplateGroup(cfg aws.Config, client *medialive.Client) {
	input := &medialive.GetCloudWatchAlarmTemplateGroupInput{
		// Identifier: *string, // Required
	}

	if len(_medialiveIdentifier) > 0 {
		input.Identifier = aws.String(_medialiveIdentifier)
	}

	if resp, err := client.GetCloudWatchAlarmTemplateGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the specified eventbridge rule template.
func medialive_GetEventBridgeRuleTemplate(cfg aws.Config, client *medialive.Client) {
	input := &medialive.GetEventBridgeRuleTemplateInput{
		// Identifier: *string, // Required
	}

	if len(_medialiveIdentifier) > 0 {
		input.Identifier = aws.String(_medialiveIdentifier)
	}

	if resp, err := client.GetEventBridgeRuleTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the specified eventbridge rule template group.
func medialive_GetEventBridgeRuleTemplateGroup(cfg aws.Config, client *medialive.Client) {
	input := &medialive.GetEventBridgeRuleTemplateGroupInput{
		// Identifier: *string, // Required
	}

	if len(_medialiveIdentifier) > 0 {
		input.Identifier = aws.String(_medialiveIdentifier)
	}

	if resp, err := client.GetEventBridgeRuleTemplateGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the specified signal map.
func medialive_GetSignalMap(cfg aws.Config, client *medialive.Client) {
	input := &medialive.GetSignalMapInput{
		// Identifier: *string, // Required
	}

	if len(_medialiveIdentifier) > 0 {
		input.Identifier = aws.String(_medialiveIdentifier)
	}

	if resp, err := client.GetSignalMap(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List the alerts for a channel with optional filtering based on alert state.
func medialive_ListAlerts(cfg aws.Config, client *medialive.Client) {
	input := &medialive.ListAlertsInput{
		// ChannelId: *string, // Required
	}

	if len(_medialiveChannelId) > 0 {
		input.ChannelId = aws.String(_medialiveChannelId)
	}
	if len(_medialiveMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _medialiveMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_medialiveNextToken) > 0 {
		input.NextToken = aws.String(_medialiveNextToken)
	}
	if len(_medialiveStateFilter) > 0 {
		input.StateFilter = aws.String(_medialiveStateFilter)
	}

	if disablePaginator() {
		if resp, err := client.ListAlerts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*medialive.ListAlertsOutput
	p := medialive.NewListAlertsPaginator(client, input)
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

// Retrieve the list of ChannelPlacementGroups in the specified Cluster.
func medialive_ListChannelPlacementGroups(cfg aws.Config, client *medialive.Client) {
	input := &medialive.ListChannelPlacementGroupsInput{
		// ClusterId: *string, // Required
	}

	if len(_medialiveClusterId) > 0 {
		input.ClusterId = aws.String(_medialiveClusterId)
	}
	if len(_medialiveMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _medialiveMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_medialiveNextToken) > 0 {
		input.NextToken = aws.String(_medialiveNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListChannelPlacementGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*medialive.ListChannelPlacementGroupsOutput
	p := medialive.NewListChannelPlacementGroupsPaginator(client, input)
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

// Produces list of channels that have been created
func medialive_ListChannels(cfg aws.Config, client *medialive.Client) {
	input := &medialive.ListChannelsInput{}

	if len(_medialiveMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _medialiveMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_medialiveNextToken) > 0 {
		input.NextToken = aws.String(_medialiveNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListChannels(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*medialive.ListChannelsOutput
	p := medialive.NewListChannelsPaginator(client, input)
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

// Lists cloudwatch alarm template groups.
func medialive_ListCloudWatchAlarmTemplateGroups(cfg aws.Config, client *medialive.Client) {
	input := &medialive.ListCloudWatchAlarmTemplateGroupsInput{}

	if len(_medialiveMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _medialiveMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_medialiveNextToken) > 0 {
		input.NextToken = aws.String(_medialiveNextToken)
	}
	if len(_medialiveScope) > 0 {
		input.Scope = aws.String(_medialiveScope)
	}
	if len(_medialiveSignalMapIdentifier) > 0 {
		input.SignalMapIdentifier = aws.String(_medialiveSignalMapIdentifier)
	}

	if disablePaginator() {
		if resp, err := client.ListCloudWatchAlarmTemplateGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*medialive.ListCloudWatchAlarmTemplateGroupsOutput
	p := medialive.NewListCloudWatchAlarmTemplateGroupsPaginator(client, input)
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

// Lists cloudwatch alarm templates.
func medialive_ListCloudWatchAlarmTemplates(cfg aws.Config, client *medialive.Client) {
	input := &medialive.ListCloudWatchAlarmTemplatesInput{}

	if len(_medialiveGroupIdentifier) > 0 {
		input.GroupIdentifier = aws.String(_medialiveGroupIdentifier)
	}
	if len(_medialiveMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _medialiveMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_medialiveNextToken) > 0 {
		input.NextToken = aws.String(_medialiveNextToken)
	}
	if len(_medialiveScope) > 0 {
		input.Scope = aws.String(_medialiveScope)
	}
	if len(_medialiveSignalMapIdentifier) > 0 {
		input.SignalMapIdentifier = aws.String(_medialiveSignalMapIdentifier)
	}

	if disablePaginator() {
		if resp, err := client.ListCloudWatchAlarmTemplates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*medialive.ListCloudWatchAlarmTemplatesOutput
	p := medialive.NewListCloudWatchAlarmTemplatesPaginator(client, input)
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

// List the alerts for a cluster with optional filtering based on alert state.
func medialive_ListClusterAlerts(cfg aws.Config, client *medialive.Client) {
	input := &medialive.ListClusterAlertsInput{
		// ClusterId: *string, // Required
	}

	if len(_medialiveClusterId) > 0 {
		input.ClusterId = aws.String(_medialiveClusterId)
	}
	if len(_medialiveMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _medialiveMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_medialiveNextToken) > 0 {
		input.NextToken = aws.String(_medialiveNextToken)
	}
	if len(_medialiveStateFilter) > 0 {
		input.StateFilter = aws.String(_medialiveStateFilter)
	}

	if disablePaginator() {
		if resp, err := client.ListClusterAlerts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*medialive.ListClusterAlertsOutput
	p := medialive.NewListClusterAlertsPaginator(client, input)
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

// Retrieve the list of Clusters.
func medialive_ListClusters(cfg aws.Config, client *medialive.Client) {
	input := &medialive.ListClustersInput{}

	if len(_medialiveMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _medialiveMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_medialiveNextToken) > 0 {
		input.NextToken = aws.String(_medialiveNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListClusters(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*medialive.ListClustersOutput
	p := medialive.NewListClustersPaginator(client, input)
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

// Lists eventbridge rule template groups.
func medialive_ListEventBridgeRuleTemplateGroups(cfg aws.Config, client *medialive.Client) {
	input := &medialive.ListEventBridgeRuleTemplateGroupsInput{}

	if len(_medialiveMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _medialiveMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_medialiveNextToken) > 0 {
		input.NextToken = aws.String(_medialiveNextToken)
	}
	if len(_medialiveSignalMapIdentifier) > 0 {
		input.SignalMapIdentifier = aws.String(_medialiveSignalMapIdentifier)
	}

	if disablePaginator() {
		if resp, err := client.ListEventBridgeRuleTemplateGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*medialive.ListEventBridgeRuleTemplateGroupsOutput
	p := medialive.NewListEventBridgeRuleTemplateGroupsPaginator(client, input)
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

// Lists eventbridge rule templates.
func medialive_ListEventBridgeRuleTemplates(cfg aws.Config, client *medialive.Client) {
	input := &medialive.ListEventBridgeRuleTemplatesInput{}

	if len(_medialiveGroupIdentifier) > 0 {
		input.GroupIdentifier = aws.String(_medialiveGroupIdentifier)
	}
	if len(_medialiveMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _medialiveMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_medialiveNextToken) > 0 {
		input.NextToken = aws.String(_medialiveNextToken)
	}
	if len(_medialiveSignalMapIdentifier) > 0 {
		input.SignalMapIdentifier = aws.String(_medialiveSignalMapIdentifier)
	}

	if disablePaginator() {
		if resp, err := client.ListEventBridgeRuleTemplates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*medialive.ListEventBridgeRuleTemplatesOutput
	p := medialive.NewListEventBridgeRuleTemplatesPaginator(client, input)
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

// List input devices that are currently being transferred. List input devices
// that you are transferring from your AWS account or input devices that another
// AWS account is transferring to you.
func medialive_ListInputDeviceTransfers(cfg aws.Config, client *medialive.Client) {
	input := &medialive.ListInputDeviceTransfersInput{
		// TransferType: *string, // Required
	}

	if len(_medialiveTransferType) > 0 {
		input.TransferType = aws.String(_medialiveTransferType)
	}
	if len(_medialiveMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _medialiveMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_medialiveNextToken) > 0 {
		input.NextToken = aws.String(_medialiveNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListInputDeviceTransfers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*medialive.ListInputDeviceTransfersOutput
	p := medialive.NewListInputDeviceTransfersPaginator(client, input)
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

// List input devices
func medialive_ListInputDevices(cfg aws.Config, client *medialive.Client) {
	input := &medialive.ListInputDevicesInput{}

	if len(_medialiveMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _medialiveMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_medialiveNextToken) > 0 {
		input.NextToken = aws.String(_medialiveNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListInputDevices(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*medialive.ListInputDevicesOutput
	p := medialive.NewListInputDevicesPaginator(client, input)
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

// Produces a list of Input Security Groups for an account
func medialive_ListInputSecurityGroups(cfg aws.Config, client *medialive.Client) {
	input := &medialive.ListInputSecurityGroupsInput{}

	if len(_medialiveMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _medialiveMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_medialiveNextToken) > 0 {
		input.NextToken = aws.String(_medialiveNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListInputSecurityGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*medialive.ListInputSecurityGroupsOutput
	p := medialive.NewListInputSecurityGroupsPaginator(client, input)
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

// Produces list of inputs that have been created
func medialive_ListInputs(cfg aws.Config, client *medialive.Client) {
	input := &medialive.ListInputsInput{}

	if len(_medialiveMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _medialiveMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_medialiveNextToken) > 0 {
		input.NextToken = aws.String(_medialiveNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListInputs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*medialive.ListInputsOutput
	p := medialive.NewListInputsPaginator(client, input)
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

// List the alerts for a multiplex with optional filtering based on alert state.
func medialive_ListMultiplexAlerts(cfg aws.Config, client *medialive.Client) {
	input := &medialive.ListMultiplexAlertsInput{
		// MultiplexId: *string, // Required
	}

	if len(_medialiveMultiplexId) > 0 {
		input.MultiplexId = aws.String(_medialiveMultiplexId)
	}
	if len(_medialiveMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _medialiveMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_medialiveNextToken) > 0 {
		input.NextToken = aws.String(_medialiveNextToken)
	}
	if len(_medialiveStateFilter) > 0 {
		input.StateFilter = aws.String(_medialiveStateFilter)
	}

	if disablePaginator() {
		if resp, err := client.ListMultiplexAlerts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*medialive.ListMultiplexAlertsOutput
	p := medialive.NewListMultiplexAlertsPaginator(client, input)
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

// List the programs that currently exist for a specific multiplex.
func medialive_ListMultiplexPrograms(cfg aws.Config, client *medialive.Client) {
	input := &medialive.ListMultiplexProgramsInput{
		// MultiplexId: *string, // Required
	}

	if len(_medialiveMultiplexId) > 0 {
		input.MultiplexId = aws.String(_medialiveMultiplexId)
	}
	if len(_medialiveMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _medialiveMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_medialiveNextToken) > 0 {
		input.NextToken = aws.String(_medialiveNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListMultiplexPrograms(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*medialive.ListMultiplexProgramsOutput
	p := medialive.NewListMultiplexProgramsPaginator(client, input)
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

// Retrieve a list of the existing multiplexes.
func medialive_ListMultiplexes(cfg aws.Config, client *medialive.Client) {
	input := &medialive.ListMultiplexesInput{}

	if len(_medialiveMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _medialiveMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_medialiveNextToken) > 0 {
		input.NextToken = aws.String(_medialiveNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListMultiplexes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*medialive.ListMultiplexesOutput
	p := medialive.NewListMultiplexesPaginator(client, input)
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

// Retrieve the list of Networks.
func medialive_ListNetworks(cfg aws.Config, client *medialive.Client) {
	input := &medialive.ListNetworksInput{}

	if len(_medialiveMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _medialiveMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_medialiveNextToken) > 0 {
		input.NextToken = aws.String(_medialiveNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListNetworks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*medialive.ListNetworksOutput
	p := medialive.NewListNetworksPaginator(client, input)
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

// Retrieve the list of Nodes.
func medialive_ListNodes(cfg aws.Config, client *medialive.Client) {
	input := &medialive.ListNodesInput{
		// ClusterId: *string, // Required
	}

	if len(_medialiveClusterId) > 0 {
		input.ClusterId = aws.String(_medialiveClusterId)
	}
	if len(_medialiveMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _medialiveMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_medialiveNextToken) > 0 {
		input.NextToken = aws.String(_medialiveNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListNodes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*medialive.ListNodesOutput
	p := medialive.NewListNodesPaginator(client, input)
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

// List offerings available for purchase.
func medialive_ListOfferings(cfg aws.Config, client *medialive.Client) {
	input := &medialive.ListOfferingsInput{}

	if len(_medialiveChannelClass) > 0 {
		input.ChannelClass = aws.String(_medialiveChannelClass)
	}
	if len(_medialiveChannelConfiguration) > 0 {
		input.ChannelConfiguration = aws.String(_medialiveChannelConfiguration)
	}
	if len(_medialiveCodec) > 0 {
		input.Codec = aws.String(_medialiveCodec)
	}
	if len(_medialiveDuration) > 0 {
		input.Duration = aws.String(_medialiveDuration)
	}
	if len(_medialiveMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _medialiveMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_medialiveMaximumBitrate) > 0 {
		input.MaximumBitrate = aws.String(_medialiveMaximumBitrate)
	}
	if len(_medialiveMaximumFramerate) > 0 {
		input.MaximumFramerate = aws.String(_medialiveMaximumFramerate)
	}
	if len(_medialiveNextToken) > 0 {
		input.NextToken = aws.String(_medialiveNextToken)
	}
	if len(_medialiveResolution) > 0 {
		input.Resolution = aws.String(_medialiveResolution)
	}
	if len(_medialiveResourceType) > 0 {
		input.ResourceType = aws.String(_medialiveResourceType)
	}
	if len(_medialiveSpecialFeature) > 0 {
		input.SpecialFeature = aws.String(_medialiveSpecialFeature)
	}
	if len(_medialiveVideoQuality) > 0 {
		input.VideoQuality = aws.String(_medialiveVideoQuality)
	}

	if disablePaginator() {
		if resp, err := client.ListOfferings(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*medialive.ListOfferingsOutput
	p := medialive.NewListOfferingsPaginator(client, input)
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

// List purchased reservations.
func medialive_ListReservations(cfg aws.Config, client *medialive.Client) {
	input := &medialive.ListReservationsInput{}

	if len(_medialiveChannelClass) > 0 {
		input.ChannelClass = aws.String(_medialiveChannelClass)
	}
	if len(_medialiveCodec) > 0 {
		input.Codec = aws.String(_medialiveCodec)
	}
	if len(_medialiveMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _medialiveMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_medialiveMaximumBitrate) > 0 {
		input.MaximumBitrate = aws.String(_medialiveMaximumBitrate)
	}
	if len(_medialiveMaximumFramerate) > 0 {
		input.MaximumFramerate = aws.String(_medialiveMaximumFramerate)
	}
	if len(_medialiveNextToken) > 0 {
		input.NextToken = aws.String(_medialiveNextToken)
	}
	if len(_medialiveResolution) > 0 {
		input.Resolution = aws.String(_medialiveResolution)
	}
	if len(_medialiveResourceType) > 0 {
		input.ResourceType = aws.String(_medialiveResourceType)
	}
	if len(_medialiveSpecialFeature) > 0 {
		input.SpecialFeature = aws.String(_medialiveSpecialFeature)
	}
	if len(_medialiveVideoQuality) > 0 {
		input.VideoQuality = aws.String(_medialiveVideoQuality)
	}

	if disablePaginator() {
		if resp, err := client.ListReservations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*medialive.ListReservationsOutput
	p := medialive.NewListReservationsPaginator(client, input)
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

// List all the SdiSources in the AWS account.
func medialive_ListSdiSources(cfg aws.Config, client *medialive.Client) {
	input := &medialive.ListSdiSourcesInput{}

	if len(_medialiveMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _medialiveMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_medialiveNextToken) > 0 {
		input.NextToken = aws.String(_medialiveNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSdiSources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*medialive.ListSdiSourcesOutput
	p := medialive.NewListSdiSourcesPaginator(client, input)
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

// Lists signal maps.
func medialive_ListSignalMaps(cfg aws.Config, client *medialive.Client) {
	input := &medialive.ListSignalMapsInput{}

	if len(_medialiveCloudWatchAlarmTemplateGroupIdentifier) > 0 {
		input.CloudWatchAlarmTemplateGroupIdentifier = aws.String(_medialiveCloudWatchAlarmTemplateGroupIdentifier)
	}
	if len(_medialiveEventBridgeRuleTemplateGroupIdentifier) > 0 {
		input.EventBridgeRuleTemplateGroupIdentifier = aws.String(_medialiveEventBridgeRuleTemplateGroupIdentifier)
	}
	if len(_medialiveMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _medialiveMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_medialiveNextToken) > 0 {
		input.NextToken = aws.String(_medialiveNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSignalMaps(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*medialive.ListSignalMapsOutput
	p := medialive.NewListSignalMapsPaginator(client, input)
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

// Produces list of tags that have been created for a resource
func medialive_ListTagsForResource(cfg aws.Config, client *medialive.Client) {
	input := &medialive.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_medialiveResourceArn) > 0 {
		input.ResourceArn = aws.String(_medialiveResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves an array of all the encoder engine versions that are available in
// this AWS account.
func medialive_ListVersions(cfg aws.Config, client *medialive.Client) {
	input := &medialive.ListVersionsInput{}

	if resp, err := client.ListVersions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Purchase an offering and create a reservation.
func medialive_PurchaseOffering(cfg aws.Config, client *medialive.Client) {
	input := &medialive.PurchaseOfferingInput{
		// Count: *int32, // Required
		// OfferingId: *string, // Required
	}

	if len(_medialiveCount) > 0 {
		if err := assignInputField(input, "Count", _medialiveCount); err != nil {
			log.Errorf("invalid --count: %s", err.Error())
			return
		}
	}
	if len(_medialiveOfferingId) > 0 {
		input.OfferingId = aws.String(_medialiveOfferingId)
	}
	if len(_medialiveName) > 0 {
		input.Name = aws.String(_medialiveName)
	}
	if len(_medialiveRenewalSettings) > 0 {
		if err := assignInputField(input, "RenewalSettings", _medialiveRenewalSettings); err != nil {
			log.Errorf("invalid --renewal-settings: %s", err.Error())
			return
		}
	}
	if len(_medialiveRequestId) > 0 {
		input.RequestId = aws.String(_medialiveRequestId)
	}
	if len(_medialiveStart) > 0 {
		input.Start = aws.String(_medialiveStart)
	}
	if len(_medialiveTags) > 0 {
		if err := assignInputField(input, "Tags", _medialiveTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.PurchaseOffering(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Send a reboot command to the specified input device. The device will begin
// rebooting within a few seconds of sending the command. When the reboot is
// complete, the device’s connection status will change to connected.
func medialive_RebootInputDevice(cfg aws.Config, client *medialive.Client) {
	input := &medialive.RebootInputDeviceInput{
		// InputDeviceId: *string, // Required
	}

	if len(_medialiveInputDeviceId) > 0 {
		input.InputDeviceId = aws.String(_medialiveInputDeviceId)
	}
	if len(_medialiveForce) > 0 {
		if err := assignInputField(input, "Force", _medialiveForce); err != nil {
			log.Errorf("invalid --force: %s", err.Error())
			return
		}
	}

	if resp, err := client.RebootInputDevice(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Reject the transfer of the specified input device to your AWS account.
func medialive_RejectInputDeviceTransfer(cfg aws.Config, client *medialive.Client) {
	input := &medialive.RejectInputDeviceTransferInput{
		// InputDeviceId: *string, // Required
	}

	if len(_medialiveInputDeviceId) > 0 {
		input.InputDeviceId = aws.String(_medialiveInputDeviceId)
	}

	if resp, err := client.RejectInputDeviceTransfer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Restart pipelines in one channel that is currently running.
func medialive_RestartChannelPipelines(cfg aws.Config, client *medialive.Client) {
	input := &medialive.RestartChannelPipelinesInput{
		// ChannelId: *string, // Required
	}

	if len(_medialiveChannelId) > 0 {
		input.ChannelId = aws.String(_medialiveChannelId)
	}
	if len(_medialivePipelineIds) > 0 {
		if err := assignInputField(input, "PipelineIds", _medialivePipelineIds); err != nil {
			log.Errorf("invalid --pipeline-ids: %s", err.Error())
			return
		}
	}

	if resp, err := client.RestartChannelPipelines(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts an existing channel
func medialive_StartChannel(cfg aws.Config, client *medialive.Client) {
	input := &medialive.StartChannelInput{
		// ChannelId: *string, // Required
	}

	if len(_medialiveChannelId) > 0 {
		input.ChannelId = aws.String(_medialiveChannelId)
	}

	if resp, err := client.StartChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initiates a deployment to delete the monitor of the specified signal map.
func medialive_StartDeleteMonitorDeployment(cfg aws.Config, client *medialive.Client) {
	input := &medialive.StartDeleteMonitorDeploymentInput{
		// Identifier: *string, // Required
	}

	if len(_medialiveIdentifier) > 0 {
		input.Identifier = aws.String(_medialiveIdentifier)
	}

	if resp, err := client.StartDeleteMonitorDeployment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Start an input device that is attached to a MediaConnect flow. (There is no
// need to start a device that is attached to a MediaLive input; MediaLive starts
// the device when the channel starts.)
func medialive_StartInputDevice(cfg aws.Config, client *medialive.Client) {
	input := &medialive.StartInputDeviceInput{
		// InputDeviceId: *string, // Required
	}

	if len(_medialiveInputDeviceId) > 0 {
		input.InputDeviceId = aws.String(_medialiveInputDeviceId)
	}

	if resp, err := client.StartInputDevice(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Start a maintenance window for the specified input device. Starting a
// maintenance window will give the device up to two hours to install software. If
// the device was streaming prior to the maintenance, it will resume streaming when
// the software is fully installed. Devices automatically install updates while
// they are powered on and their MediaLive channels are stopped. A maintenance
// window allows you to update a device without having to stop MediaLive channels
// that use the device. The device must remain powered on and connected to the
// internet for the duration of the maintenance.
func medialive_StartInputDeviceMaintenanceWindow(cfg aws.Config, client *medialive.Client) {
	input := &medialive.StartInputDeviceMaintenanceWindowInput{
		// InputDeviceId: *string, // Required
	}

	if len(_medialiveInputDeviceId) > 0 {
		input.InputDeviceId = aws.String(_medialiveInputDeviceId)
	}

	if resp, err := client.StartInputDeviceMaintenanceWindow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initiates a deployment to deploy the latest monitor of the specified signal map.
func medialive_StartMonitorDeployment(cfg aws.Config, client *medialive.Client) {
	input := &medialive.StartMonitorDeploymentInput{
		// Identifier: *string, // Required
	}

	if len(_medialiveIdentifier) > 0 {
		input.Identifier = aws.String(_medialiveIdentifier)
	}
	if len(_medialiveDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _medialiveDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartMonitorDeployment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Start (run) the multiplex. Starting the multiplex does not start the channels.
// You must explicitly start each channel.
func medialive_StartMultiplex(cfg aws.Config, client *medialive.Client) {
	input := &medialive.StartMultiplexInput{
		// MultiplexId: *string, // Required
	}

	if len(_medialiveMultiplexId) > 0 {
		input.MultiplexId = aws.String(_medialiveMultiplexId)
	}

	if resp, err := client.StartMultiplex(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initiates an update for the specified signal map. Will discover a new signal
// map if a changed discoveryEntryPointArn is provided.
func medialive_StartUpdateSignalMap(cfg aws.Config, client *medialive.Client) {
	input := &medialive.StartUpdateSignalMapInput{
		// Identifier: *string, // Required
	}

	if len(_medialiveIdentifier) > 0 {
		input.Identifier = aws.String(_medialiveIdentifier)
	}
	if len(_medialiveCloudWatchAlarmTemplateGroupIdentifiers) > 0 {
		input.CloudWatchAlarmTemplateGroupIdentifiers = append([]string(nil), _medialiveCloudWatchAlarmTemplateGroupIdentifiers...)
	}
	if len(_medialiveDescription) > 0 {
		input.Description = aws.String(_medialiveDescription)
	}
	if len(_medialiveDiscoveryEntryPointArn) > 0 {
		input.DiscoveryEntryPointArn = aws.String(_medialiveDiscoveryEntryPointArn)
	}
	if len(_medialiveEventBridgeRuleTemplateGroupIdentifiers) > 0 {
		input.EventBridgeRuleTemplateGroupIdentifiers = append([]string(nil), _medialiveEventBridgeRuleTemplateGroupIdentifiers...)
	}
	if len(_medialiveForceRediscovery) > 0 {
		if err := assignInputField(input, "ForceRediscovery", _medialiveForceRediscovery); err != nil {
			log.Errorf("invalid --force-rediscovery: %s", err.Error())
			return
		}
	}
	if len(_medialiveName) > 0 {
		input.Name = aws.String(_medialiveName)
	}

	if resp, err := client.StartUpdateSignalMap(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a running channel
func medialive_StopChannel(cfg aws.Config, client *medialive.Client) {
	input := &medialive.StopChannelInput{
		// ChannelId: *string, // Required
	}

	if len(_medialiveChannelId) > 0 {
		input.ChannelId = aws.String(_medialiveChannelId)
	}

	if resp, err := client.StopChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stop an input device that is attached to a MediaConnect flow. (There is no need
// to stop a device that is attached to a MediaLive input; MediaLive automatically
// stops the device when the channel stops.)
func medialive_StopInputDevice(cfg aws.Config, client *medialive.Client) {
	input := &medialive.StopInputDeviceInput{
		// InputDeviceId: *string, // Required
	}

	if len(_medialiveInputDeviceId) > 0 {
		input.InputDeviceId = aws.String(_medialiveInputDeviceId)
	}

	if resp, err := client.StopInputDevice(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a running multiplex. If the multiplex isn't running, this action has no
// effect.
func medialive_StopMultiplex(cfg aws.Config, client *medialive.Client) {
	input := &medialive.StopMultiplexInput{
		// MultiplexId: *string, // Required
	}

	if len(_medialiveMultiplexId) > 0 {
		input.MultiplexId = aws.String(_medialiveMultiplexId)
	}

	if resp, err := client.StopMultiplex(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Start an input device transfer to another AWS account. After you make the
// request, the other account must accept or reject the transfer.
func medialive_TransferInputDevice(cfg aws.Config, client *medialive.Client) {
	input := &medialive.TransferInputDeviceInput{
		// InputDeviceId: *string, // Required
	}

	if len(_medialiveInputDeviceId) > 0 {
		input.InputDeviceId = aws.String(_medialiveInputDeviceId)
	}
	if len(_medialiveTargetCustomerId) > 0 {
		input.TargetCustomerId = aws.String(_medialiveTargetCustomerId)
	}
	if len(_medialiveTargetRegion) > 0 {
		input.TargetRegion = aws.String(_medialiveTargetRegion)
	}
	if len(_medialiveTransferMessage) > 0 {
		input.TransferMessage = aws.String(_medialiveTransferMessage)
	}

	if resp, err := client.TransferInputDevice(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update account configuration
func medialive_UpdateAccountConfiguration(cfg aws.Config, client *medialive.Client) {
	input := &medialive.UpdateAccountConfigurationInput{}

	if len(_medialiveAccountConfiguration) > 0 {
		if err := assignInputField(input, "AccountConfiguration", _medialiveAccountConfiguration); err != nil {
			log.Errorf("invalid --account-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAccountConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a channel.
func medialive_UpdateChannel(cfg aws.Config, client *medialive.Client) {
	input := &medialive.UpdateChannelInput{
		// ChannelId: *string, // Required
	}

	if len(_medialiveChannelId) > 0 {
		input.ChannelId = aws.String(_medialiveChannelId)
	}
	if len(_medialiveAnywhereSettings) > 0 {
		if err := assignInputField(input, "AnywhereSettings", _medialiveAnywhereSettings); err != nil {
			log.Errorf("invalid --anywhere-settings: %s", err.Error())
			return
		}
	}
	if len(_medialiveCdiInputSpecification) > 0 {
		if err := assignInputField(input, "CdiInputSpecification", _medialiveCdiInputSpecification); err != nil {
			log.Errorf("invalid --cdi-input-specification: %s", err.Error())
			return
		}
	}
	if len(_medialiveChannelEngineVersion) > 0 {
		if err := assignInputField(input, "ChannelEngineVersion", _medialiveChannelEngineVersion); err != nil {
			log.Errorf("invalid --channel-engine-version: %s", err.Error())
			return
		}
	}
	if len(_medialiveChannelSecurityGroups) > 0 {
		input.ChannelSecurityGroups = append([]string(nil), _medialiveChannelSecurityGroups...)
	}
	if len(_medialiveDestinations) > 0 {
		if err := assignInputField(input, "Destinations", _medialiveDestinations); err != nil {
			log.Errorf("invalid --destinations: %s", err.Error())
			return
		}
	}
	if len(_medialiveDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _medialiveDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}
	if len(_medialiveEncoderSettings) > 0 {
		if err := assignInputField(input, "EncoderSettings", _medialiveEncoderSettings); err != nil {
			log.Errorf("invalid --encoder-settings: %s", err.Error())
			return
		}
	}
	if len(_medialiveInferenceSettings) > 0 {
		if err := assignInputField(input, "InferenceSettings", _medialiveInferenceSettings); err != nil {
			log.Errorf("invalid --inference-settings: %s", err.Error())
			return
		}
	}
	if len(_medialiveInputAttachments) > 0 {
		if err := assignInputField(input, "InputAttachments", _medialiveInputAttachments); err != nil {
			log.Errorf("invalid --input-attachments: %s", err.Error())
			return
		}
	}
	if len(_medialiveInputSpecification) > 0 {
		if err := assignInputField(input, "InputSpecification", _medialiveInputSpecification); err != nil {
			log.Errorf("invalid --input-specification: %s", err.Error())
			return
		}
	}
	if len(_medialiveLinkedChannelSettings) > 0 {
		if err := assignInputField(input, "LinkedChannelSettings", _medialiveLinkedChannelSettings); err != nil {
			log.Errorf("invalid --linked-channel-settings: %s", err.Error())
			return
		}
	}
	if len(_medialiveLogLevel) > 0 {
		if err := assignInputField(input, "LogLevel", _medialiveLogLevel); err != nil {
			log.Errorf("invalid --log-level: %s", err.Error())
			return
		}
	}
	if len(_medialiveMaintenance) > 0 {
		if err := assignInputField(input, "Maintenance", _medialiveMaintenance); err != nil {
			log.Errorf("invalid --maintenance: %s", err.Error())
			return
		}
	}
	if len(_medialiveName) > 0 {
		input.Name = aws.String(_medialiveName)
	}
	if len(_medialiveRoleArn) > 0 {
		input.RoleArn = aws.String(_medialiveRoleArn)
	}

	if resp, err := client.UpdateChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes the class of the channel.
func medialive_UpdateChannelClass(cfg aws.Config, client *medialive.Client) {
	input := &medialive.UpdateChannelClassInput{
		// ChannelClass: types.ChannelClass, // Required
		// ChannelId: *string, // Required
	}

	if len(_medialiveChannelClass) > 0 {
		if err := assignInputField(input, "ChannelClass", _medialiveChannelClass); err != nil {
			log.Errorf("invalid --channel-class: %s", err.Error())
			return
		}
	}
	if len(_medialiveChannelId) > 0 {
		input.ChannelId = aws.String(_medialiveChannelId)
	}
	if len(_medialiveDestinations) > 0 {
		if err := assignInputField(input, "Destinations", _medialiveDestinations); err != nil {
			log.Errorf("invalid --destinations: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateChannelClass(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Change the settings for a ChannelPlacementGroup.
func medialive_UpdateChannelPlacementGroup(cfg aws.Config, client *medialive.Client) {
	input := &medialive.UpdateChannelPlacementGroupInput{
		// ChannelPlacementGroupId: *string, // Required
		// ClusterId: *string, // Required
	}

	if len(_medialiveChannelPlacementGroupId) > 0 {
		input.ChannelPlacementGroupId = aws.String(_medialiveChannelPlacementGroupId)
	}
	if len(_medialiveClusterId) > 0 {
		input.ClusterId = aws.String(_medialiveClusterId)
	}
	if len(_medialiveName) > 0 {
		input.Name = aws.String(_medialiveName)
	}
	if len(_medialiveNodes) > 0 {
		input.Nodes = append([]string(nil), _medialiveNodes...)
	}

	if resp, err := client.UpdateChannelPlacementGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified cloudwatch alarm template.
func medialive_UpdateCloudWatchAlarmTemplate(cfg aws.Config, client *medialive.Client) {
	input := &medialive.UpdateCloudWatchAlarmTemplateInput{
		// Identifier: *string, // Required
	}

	if len(_medialiveIdentifier) > 0 {
		input.Identifier = aws.String(_medialiveIdentifier)
	}
	if len(_medialiveComparisonOperator) > 0 {
		if err := assignInputField(input, "ComparisonOperator", _medialiveComparisonOperator); err != nil {
			log.Errorf("invalid --comparison-operator: %s", err.Error())
			return
		}
	}
	if len(_medialiveDatapointsToAlarm) > 0 {
		if err := assignInputField(input, "DatapointsToAlarm", _medialiveDatapointsToAlarm); err != nil {
			log.Errorf("invalid --datapoints-to-alarm: %s", err.Error())
			return
		}
	}
	if len(_medialiveDescription) > 0 {
		input.Description = aws.String(_medialiveDescription)
	}
	if len(_medialiveEvaluationPeriods) > 0 {
		if err := assignInputField(input, "EvaluationPeriods", _medialiveEvaluationPeriods); err != nil {
			log.Errorf("invalid --evaluation-periods: %s", err.Error())
			return
		}
	}
	if len(_medialiveGroupIdentifier) > 0 {
		input.GroupIdentifier = aws.String(_medialiveGroupIdentifier)
	}
	if len(_medialiveMetricName) > 0 {
		input.MetricName = aws.String(_medialiveMetricName)
	}
	if len(_medialiveName) > 0 {
		input.Name = aws.String(_medialiveName)
	}
	if len(_medialivePeriod) > 0 {
		if err := assignInputField(input, "Period", _medialivePeriod); err != nil {
			log.Errorf("invalid --period: %s", err.Error())
			return
		}
	}
	if len(_medialiveStatistic) > 0 {
		if err := assignInputField(input, "Statistic", _medialiveStatistic); err != nil {
			log.Errorf("invalid --statistic: %s", err.Error())
			return
		}
	}
	if len(_medialiveTargetResourceType) > 0 {
		if err := assignInputField(input, "TargetResourceType", _medialiveTargetResourceType); err != nil {
			log.Errorf("invalid --target-resource-type: %s", err.Error())
			return
		}
	}
	if len(_medialiveThreshold) > 0 {
		if err := assignInputField(input, "Threshold", _medialiveThreshold); err != nil {
			log.Errorf("invalid --threshold: %s", err.Error())
			return
		}
	}
	if len(_medialiveTreatMissingData) > 0 {
		if err := assignInputField(input, "TreatMissingData", _medialiveTreatMissingData); err != nil {
			log.Errorf("invalid --treat-missing-data: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateCloudWatchAlarmTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified cloudwatch alarm template group.
func medialive_UpdateCloudWatchAlarmTemplateGroup(cfg aws.Config, client *medialive.Client) {
	input := &medialive.UpdateCloudWatchAlarmTemplateGroupInput{
		// Identifier: *string, // Required
	}

	if len(_medialiveIdentifier) > 0 {
		input.Identifier = aws.String(_medialiveIdentifier)
	}
	if len(_medialiveDescription) > 0 {
		input.Description = aws.String(_medialiveDescription)
	}

	if resp, err := client.UpdateCloudWatchAlarmTemplateGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Change the settings for a Cluster.
func medialive_UpdateCluster(cfg aws.Config, client *medialive.Client) {
	input := &medialive.UpdateClusterInput{
		// ClusterId: *string, // Required
	}

	if len(_medialiveClusterId) > 0 {
		input.ClusterId = aws.String(_medialiveClusterId)
	}
	if len(_medialiveName) > 0 {
		input.Name = aws.String(_medialiveName)
	}
	if len(_medialiveNetworkSettings) > 0 {
		if err := assignInputField(input, "NetworkSettings", _medialiveNetworkSettings); err != nil {
			log.Errorf("invalid --network-settings: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified eventbridge rule template.
func medialive_UpdateEventBridgeRuleTemplate(cfg aws.Config, client *medialive.Client) {
	input := &medialive.UpdateEventBridgeRuleTemplateInput{
		// Identifier: *string, // Required
	}

	if len(_medialiveIdentifier) > 0 {
		input.Identifier = aws.String(_medialiveIdentifier)
	}
	if len(_medialiveDescription) > 0 {
		input.Description = aws.String(_medialiveDescription)
	}
	if len(_medialiveEventTargets) > 0 {
		if err := assignInputField(input, "EventTargets", _medialiveEventTargets); err != nil {
			log.Errorf("invalid --event-targets: %s", err.Error())
			return
		}
	}
	if len(_medialiveEventType) > 0 {
		if err := assignInputField(input, "EventType", _medialiveEventType); err != nil {
			log.Errorf("invalid --event-type: %s", err.Error())
			return
		}
	}
	if len(_medialiveGroupIdentifier) > 0 {
		input.GroupIdentifier = aws.String(_medialiveGroupIdentifier)
	}
	if len(_medialiveName) > 0 {
		input.Name = aws.String(_medialiveName)
	}

	if resp, err := client.UpdateEventBridgeRuleTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified eventbridge rule template group.
func medialive_UpdateEventBridgeRuleTemplateGroup(cfg aws.Config, client *medialive.Client) {
	input := &medialive.UpdateEventBridgeRuleTemplateGroupInput{
		// Identifier: *string, // Required
	}

	if len(_medialiveIdentifier) > 0 {
		input.Identifier = aws.String(_medialiveIdentifier)
	}
	if len(_medialiveDescription) > 0 {
		input.Description = aws.String(_medialiveDescription)
	}

	if resp, err := client.UpdateEventBridgeRuleTemplateGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an input.
func medialive_UpdateInput(cfg aws.Config, client *medialive.Client) {
	input := &medialive.UpdateInputInput{
		// InputId: *string, // Required
	}

	if len(_medialiveInputId) > 0 {
		input.InputId = aws.String(_medialiveInputId)
	}
	if len(_medialiveDestinations) > 0 {
		if err := assignInputField(input, "Destinations", _medialiveDestinations); err != nil {
			log.Errorf("invalid --destinations: %s", err.Error())
			return
		}
	}
	if len(_medialiveInputDevices) > 0 {
		if err := assignInputField(input, "InputDevices", _medialiveInputDevices); err != nil {
			log.Errorf("invalid --input-devices: %s", err.Error())
			return
		}
	}
	if len(_medialiveInputSecurityGroups) > 0 {
		input.InputSecurityGroups = append([]string(nil), _medialiveInputSecurityGroups...)
	}
	if len(_medialiveMediaConnectFlows) > 0 {
		if err := assignInputField(input, "MediaConnectFlows", _medialiveMediaConnectFlows); err != nil {
			log.Errorf("invalid --media-connect-flows: %s", err.Error())
			return
		}
	}
	if len(_medialiveMulticastSettings) > 0 {
		if err := assignInputField(input, "MulticastSettings", _medialiveMulticastSettings); err != nil {
			log.Errorf("invalid --multicast-settings: %s", err.Error())
			return
		}
	}
	if len(_medialiveName) > 0 {
		input.Name = aws.String(_medialiveName)
	}
	if len(_medialiveRoleArn) > 0 {
		input.RoleArn = aws.String(_medialiveRoleArn)
	}
	if len(_medialiveSdiSources) > 0 {
		input.SdiSources = append([]string(nil), _medialiveSdiSources...)
	}
	if len(_medialiveSmpte2110ReceiverGroupSettings) > 0 {
		if err := assignInputField(input, "Smpte2110ReceiverGroupSettings", _medialiveSmpte2110ReceiverGroupSettings); err != nil {
			log.Errorf("invalid --smpte2110-receiver-group-settings: %s", err.Error())
			return
		}
	}
	if len(_medialiveSources) > 0 {
		if err := assignInputField(input, "Sources", _medialiveSources); err != nil {
			log.Errorf("invalid --sources: %s", err.Error())
			return
		}
	}
	if len(_medialiveSpecialRouterSettings) > 0 {
		if err := assignInputField(input, "SpecialRouterSettings", _medialiveSpecialRouterSettings); err != nil {
			log.Errorf("invalid --special-router-settings: %s", err.Error())
			return
		}
	}
	if len(_medialiveSrtSettings) > 0 {
		if err := assignInputField(input, "SrtSettings", _medialiveSrtSettings); err != nil {
			log.Errorf("invalid --srt-settings: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateInput(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the parameters for the input device.
func medialive_UpdateInputDevice(cfg aws.Config, client *medialive.Client) {
	input := &medialive.UpdateInputDeviceInput{
		// InputDeviceId: *string, // Required
	}

	if len(_medialiveInputDeviceId) > 0 {
		input.InputDeviceId = aws.String(_medialiveInputDeviceId)
	}
	if len(_medialiveAvailabilityZone) > 0 {
		input.AvailabilityZone = aws.String(_medialiveAvailabilityZone)
	}
	if len(_medialiveHdDeviceSettings) > 0 {
		if err := assignInputField(input, "HdDeviceSettings", _medialiveHdDeviceSettings); err != nil {
			log.Errorf("invalid --hd-device-settings: %s", err.Error())
			return
		}
	}
	if len(_medialiveName) > 0 {
		input.Name = aws.String(_medialiveName)
	}
	if len(_medialiveUhdDeviceSettings) > 0 {
		if err := assignInputField(input, "UhdDeviceSettings", _medialiveUhdDeviceSettings); err != nil {
			log.Errorf("invalid --uhd-device-settings: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateInputDevice(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update an Input Security Group's Whilelists.
func medialive_UpdateInputSecurityGroup(cfg aws.Config, client *medialive.Client) {
	input := &medialive.UpdateInputSecurityGroupInput{
		// InputSecurityGroupId: *string, // Required
	}

	if len(_medialiveInputSecurityGroupId) > 0 {
		input.InputSecurityGroupId = aws.String(_medialiveInputSecurityGroupId)
	}
	if len(_medialiveTags) > 0 {
		if err := assignInputField(input, "Tags", _medialiveTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_medialiveWhitelistRules) > 0 {
		if err := assignInputField(input, "WhitelistRules", _medialiveWhitelistRules); err != nil {
			log.Errorf("invalid --whitelist-rules: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateInputSecurityGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a multiplex.
func medialive_UpdateMultiplex(cfg aws.Config, client *medialive.Client) {
	input := &medialive.UpdateMultiplexInput{
		// MultiplexId: *string, // Required
	}

	if len(_medialiveMultiplexId) > 0 {
		input.MultiplexId = aws.String(_medialiveMultiplexId)
	}
	if len(_medialiveMultiplexSettings) > 0 {
		if err := assignInputField(input, "MultiplexSettings", _medialiveMultiplexSettings); err != nil {
			log.Errorf("invalid --multiplex-settings: %s", err.Error())
			return
		}
	}
	if len(_medialiveName) > 0 {
		input.Name = aws.String(_medialiveName)
	}
	if len(_medialivePacketIdentifiersMapping) > 0 {
		if err := assignInputField(input, "PacketIdentifiersMapping", _medialivePacketIdentifiersMapping); err != nil {
			log.Errorf("invalid --packet-identifiers-mapping: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateMultiplex(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update a program in a multiplex.
func medialive_UpdateMultiplexProgram(cfg aws.Config, client *medialive.Client) {
	input := &medialive.UpdateMultiplexProgramInput{
		// MultiplexId: *string, // Required
		// ProgramName: *string, // Required
	}

	if len(_medialiveMultiplexId) > 0 {
		input.MultiplexId = aws.String(_medialiveMultiplexId)
	}
	if len(_medialiveProgramName) > 0 {
		input.ProgramName = aws.String(_medialiveProgramName)
	}
	if len(_medialiveMultiplexProgramSettings) > 0 {
		if err := assignInputField(input, "MultiplexProgramSettings", _medialiveMultiplexProgramSettings); err != nil {
			log.Errorf("invalid --multiplex-program-settings: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateMultiplexProgram(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Change the settings for a Network.
func medialive_UpdateNetwork(cfg aws.Config, client *medialive.Client) {
	input := &medialive.UpdateNetworkInput{
		// NetworkId: *string, // Required
	}

	if len(_medialiveNetworkId) > 0 {
		input.NetworkId = aws.String(_medialiveNetworkId)
	}
	if len(_medialiveIpPools) > 0 {
		if err := assignInputField(input, "IpPools", _medialiveIpPools); err != nil {
			log.Errorf("invalid --ip-pools: %s", err.Error())
			return
		}
	}
	if len(_medialiveName) > 0 {
		input.Name = aws.String(_medialiveName)
	}
	if len(_medialiveRoutes) > 0 {
		if err := assignInputField(input, "Routes", _medialiveRoutes); err != nil {
			log.Errorf("invalid --routes: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateNetwork(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Change the settings for a Node.
func medialive_UpdateNode(cfg aws.Config, client *medialive.Client) {
	input := &medialive.UpdateNodeInput{
		// ClusterId: *string, // Required
		// NodeId: *string, // Required
	}

	if len(_medialiveClusterId) > 0 {
		input.ClusterId = aws.String(_medialiveClusterId)
	}
	if len(_medialiveNodeId) > 0 {
		input.NodeId = aws.String(_medialiveNodeId)
	}
	if len(_medialiveName) > 0 {
		input.Name = aws.String(_medialiveName)
	}
	if len(_medialiveRole) > 0 {
		if err := assignInputField(input, "Role", _medialiveRole); err != nil {
			log.Errorf("invalid --role: %s", err.Error())
			return
		}
	}
	if len(_medialiveSdiSourceMappings) > 0 {
		if err := assignInputField(input, "SdiSourceMappings", _medialiveSdiSourceMappings); err != nil {
			log.Errorf("invalid --sdi-source-mappings: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateNode(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update the state of a node.
func medialive_UpdateNodeState(cfg aws.Config, client *medialive.Client) {
	input := &medialive.UpdateNodeStateInput{
		// ClusterId: *string, // Required
		// NodeId: *string, // Required
	}

	if len(_medialiveClusterId) > 0 {
		input.ClusterId = aws.String(_medialiveClusterId)
	}
	if len(_medialiveNodeId) > 0 {
		input.NodeId = aws.String(_medialiveNodeId)
	}
	if len(_medialiveState) > 0 {
		if err := assignInputField(input, "State", _medialiveState); err != nil {
			log.Errorf("invalid --state: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateNodeState(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update reservation.
func medialive_UpdateReservation(cfg aws.Config, client *medialive.Client) {
	input := &medialive.UpdateReservationInput{
		// ReservationId: *string, // Required
	}

	if len(_medialiveReservationId) > 0 {
		input.ReservationId = aws.String(_medialiveReservationId)
	}
	if len(_medialiveName) > 0 {
		input.Name = aws.String(_medialiveName)
	}
	if len(_medialiveRenewalSettings) > 0 {
		if err := assignInputField(input, "RenewalSettings", _medialiveRenewalSettings); err != nil {
			log.Errorf("invalid --renewal-settings: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateReservation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Change some of the settings in an SdiSource.
func medialive_UpdateSdiSource(cfg aws.Config, client *medialive.Client) {
	input := &medialive.UpdateSdiSourceInput{
		// SdiSourceId: *string, // Required
	}

	if len(_medialiveSdiSourceId) > 0 {
		input.SdiSourceId = aws.String(_medialiveSdiSourceId)
	}
	if len(_medialiveMode) > 0 {
		if err := assignInputField(input, "Mode", _medialiveMode); err != nil {
			log.Errorf("invalid --mode: %s", err.Error())
			return
		}
	}
	if len(_medialiveName) > 0 {
		input.Name = aws.String(_medialiveName)
	}
	if len(_medialiveType) > 0 {
		if err := assignInputField(input, "Type", _medialiveType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSdiSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_medialiveCmd)
	_medialiveCmd.Flags().SortFlags = false

	_medialiveCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_medialiveCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_medialiveCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_medialiveCmd.Flags().StringVarP(&_medialiveAccept, "accept", "", "", "Accept")
	_medialiveCmd.Flags().StringVarP(&_medialiveAccountConfiguration, "account-configuration", "", "", "Account Configuration")
	_medialiveCmd.Flags().StringVarP(&_medialiveAnywhereSettings, "anywhere-settings", "", "", "Anywhere Settings")
	_medialiveCmd.Flags().StringVarP(&_medialiveAvailabilityZone, "availability-zone", "", "", "Availability Zone")
	_medialiveCmd.Flags().StringSliceVarP(&_medialiveAvailabilityZones, "availability-zones", "", nil, "Availability Zones")
	_medialiveCmd.Flags().StringVarP(&_medialiveCdiInputSpecification, "cdi-input-specification", "", "", "Cdi Input Specification")
	_medialiveCmd.Flags().StringVarP(&_medialiveChannelClass, "channel-class", "", "", "Channel Class")
	_medialiveCmd.Flags().StringVarP(&_medialiveChannelConfiguration, "channel-configuration", "", "", "Channel Configuration")
	_medialiveCmd.Flags().StringVarP(&_medialiveChannelEngineVersion, "channel-engine-version", "", "", "Channel Engine Version")
	_medialiveCmd.Flags().StringVarP(&_medialiveChannelId, "channel-id", "", "", "Channel ID")
	_medialiveCmd.Flags().StringSliceVarP(&_medialiveChannelIds, "channel-ids", "", nil, "Channel Ids")
	_medialiveCmd.Flags().StringVarP(&_medialiveChannelPlacementGroupId, "channel-placement-group-id", "", "", "Channel Placement Group ID")
	_medialiveCmd.Flags().StringSliceVarP(&_medialiveChannelSecurityGroups, "channel-security-groups", "", nil, "Channel Security Groups")
	_medialiveCmd.Flags().StringVarP(&_medialiveCloudWatchAlarmTemplateGroupIdentifier, "cloud-watch-alarm-template-group-identifier", "", "", "Cloud Watch Alarm Template Group Identifier")
	_medialiveCmd.Flags().StringSliceVarP(&_medialiveCloudWatchAlarmTemplateGroupIdentifiers, "cloud-watch-alarm-template-group-identifiers", "", nil, "Cloud Watch Alarm Template Group Identifiers")
	_medialiveCmd.Flags().StringVarP(&_medialiveClusterId, "cluster-id", "", "", "Cluster ID")
	_medialiveCmd.Flags().StringVarP(&_medialiveClusterType, "cluster-type", "", "", "Cluster Type")
	_medialiveCmd.Flags().StringVarP(&_medialiveCodec, "codec", "", "", "Codec")
	_medialiveCmd.Flags().StringVarP(&_medialiveComparisonOperator, "comparison-operator", "", "", "Comparison Operator")
	_medialiveCmd.Flags().StringVarP(&_medialiveCount, "count", "", "", "Count")
	_medialiveCmd.Flags().StringVarP(&_medialiveCreates, "creates", "", "", "Creates")
	_medialiveCmd.Flags().StringVarP(&_medialiveDatapointsToAlarm, "datapoints-to-alarm", "", "", "Datapoints To Alarm")
	_medialiveCmd.Flags().StringVarP(&_medialiveDeletes, "deletes", "", "", "Deletes")
	_medialiveCmd.Flags().StringVarP(&_medialiveDescription, "description", "", "", "Description")
	_medialiveCmd.Flags().StringVarP(&_medialiveDestinations, "destinations", "", "", "Destinations")
	_medialiveCmd.Flags().StringVarP(&_medialiveDiscoveryEntryPointArn, "discovery-entry-point-arn", "", "", "Discovery Entry Point ARN")
	_medialiveCmd.Flags().StringVarP(&_medialiveDryRun, "dry-run", "", "", "Dry Run")
	_medialiveCmd.Flags().StringVarP(&_medialiveDuration, "duration", "", "", "Duration")
	_medialiveCmd.Flags().StringVarP(&_medialiveEncoderSettings, "encoder-settings", "", "", "Encoder Settings")
	_medialiveCmd.Flags().StringVarP(&_medialiveEvaluationPeriods, "evaluation-periods", "", "", "Evaluation Periods")
	_medialiveCmd.Flags().StringVarP(&_medialiveEventBridgeRuleTemplateGroupIdentifier, "event-bridge-rule-template-group-identifier", "", "", "Event Bridge Rule Template Group Identifier")
	_medialiveCmd.Flags().StringSliceVarP(&_medialiveEventBridgeRuleTemplateGroupIdentifiers, "event-bridge-rule-template-group-identifiers", "", nil, "Event Bridge Rule Template Group Identifiers")
	_medialiveCmd.Flags().StringVarP(&_medialiveEventTargets, "event-targets", "", "", "Event Targets")
	_medialiveCmd.Flags().StringVarP(&_medialiveEventType, "event-type", "", "", "Event Type")
	_medialiveCmd.Flags().StringVarP(&_medialiveForce, "force", "", "", "Force")
	_medialiveCmd.Flags().StringVarP(&_medialiveForceRediscovery, "force-rediscovery", "", "", "Force Rediscovery")
	_medialiveCmd.Flags().StringVarP(&_medialiveGroupIdentifier, "group-identifier", "", "", "Group Identifier")
	_medialiveCmd.Flags().StringVarP(&_medialiveHdDeviceSettings, "hd-device-settings", "", "", "Hd Device Settings")
	_medialiveCmd.Flags().StringVarP(&_medialiveId, "id", "", "", "ID")
	_medialiveCmd.Flags().StringVarP(&_medialiveIdentifier, "identifier", "", "", "Identifier")
	_medialiveCmd.Flags().StringVarP(&_medialiveInferenceSettings, "inference-settings", "", "", "Inference Settings")
	_medialiveCmd.Flags().StringVarP(&_medialiveInputAttachments, "input-attachments", "", "", "Input Attachments")
	_medialiveCmd.Flags().StringVarP(&_medialiveInputDeviceId, "input-device-id", "", "", "Input Device ID")
	_medialiveCmd.Flags().StringVarP(&_medialiveInputDevices, "input-devices", "", "", "Input Devices")
	_medialiveCmd.Flags().StringVarP(&_medialiveInputId, "input-id", "", "", "Input ID")
	_medialiveCmd.Flags().StringSliceVarP(&_medialiveInputIds, "input-ids", "", nil, "Input Ids")
	_medialiveCmd.Flags().StringVarP(&_medialiveInputNetworkLocation, "input-network-location", "", "", "Input Network Location")
	_medialiveCmd.Flags().StringVarP(&_medialiveInputSecurityGroupId, "input-security-group-id", "", "", "Input Security Group ID")
	_medialiveCmd.Flags().StringSliceVarP(&_medialiveInputSecurityGroupIds, "input-security-group-ids", "", nil, "Input Security Group Ids")
	_medialiveCmd.Flags().StringSliceVarP(&_medialiveInputSecurityGroups, "input-security-groups", "", nil, "Input Security Groups")
	_medialiveCmd.Flags().StringVarP(&_medialiveInputSpecification, "input-specification", "", "", "Input Specification")
	_medialiveCmd.Flags().StringVarP(&_medialiveInstanceRoleArn, "instance-role-arn", "", "", "Instance Role ARN")
	_medialiveCmd.Flags().StringVarP(&_medialiveIpPools, "ip-pools", "", "", "IP Pools")
	_medialiveCmd.Flags().StringVarP(&_medialiveLinkedChannelSettings, "linked-channel-settings", "", "", "Linked Channel Settings")
	_medialiveCmd.Flags().StringVarP(&_medialiveLogLevel, "log-level", "", "", "Log Level")
	_medialiveCmd.Flags().StringVarP(&_medialiveMaintenance, "maintenance", "", "", "Maintenance")
	_medialiveCmd.Flags().StringVarP(&_medialiveMaxResults, "max-results", "", "", "Max Results")
	_medialiveCmd.Flags().StringVarP(&_medialiveMaximumBitrate, "maximum-bitrate", "", "", "Maximum Bitrate")
	_medialiveCmd.Flags().StringVarP(&_medialiveMaximumFramerate, "maximum-framerate", "", "", "Maximum Framerate")
	_medialiveCmd.Flags().StringVarP(&_medialiveMediaConnectFlows, "media-connect-flows", "", "", "Media Connect Flows")
	_medialiveCmd.Flags().StringVarP(&_medialiveMetricName, "metric-name", "", "", "Metric Name")
	_medialiveCmd.Flags().StringVarP(&_medialiveMode, "mode", "", "", "Mode")
	_medialiveCmd.Flags().StringVarP(&_medialiveMulticastSettings, "multicast-settings", "", "", "Multicast Settings")
	_medialiveCmd.Flags().StringVarP(&_medialiveMultiplexId, "multiplex-id", "", "", "Multiplex ID")
	_medialiveCmd.Flags().StringSliceVarP(&_medialiveMultiplexIds, "multiplex-ids", "", nil, "Multiplex Ids")
	_medialiveCmd.Flags().StringVarP(&_medialiveMultiplexProgramSettings, "multiplex-program-settings", "", "", "Multiplex Program Settings")
	_medialiveCmd.Flags().StringVarP(&_medialiveMultiplexSettings, "multiplex-settings", "", "", "Multiplex Settings")
	_medialiveCmd.Flags().StringVarP(&_medialiveName, "name", "", "", "Name")
	_medialiveCmd.Flags().StringVarP(&_medialiveNetworkId, "network-id", "", "", "Network ID")
	_medialiveCmd.Flags().StringVarP(&_medialiveNetworkSettings, "network-settings", "", "", "Network Settings")
	_medialiveCmd.Flags().StringVarP(&_medialiveNextToken, "next-token", "", "", "Next Token")
	_medialiveCmd.Flags().StringVarP(&_medialiveNodeId, "node-id", "", "", "Node ID")
	_medialiveCmd.Flags().StringVarP(&_medialiveNodeInterfaceMappings, "node-interface-mappings", "", "", "Node Interface Mappings")
	_medialiveCmd.Flags().StringSliceVarP(&_medialiveNodes, "nodes", "", nil, "Nodes")
	_medialiveCmd.Flags().StringVarP(&_medialiveOfferingId, "offering-id", "", "", "Offering ID")
	_medialiveCmd.Flags().StringVarP(&_medialivePacketIdentifiersMapping, "packet-identifiers-mapping", "", "", "Packet Identifiers Mapping")
	_medialiveCmd.Flags().StringVarP(&_medialivePeriod, "period", "", "", "Period")
	_medialiveCmd.Flags().StringVarP(&_medialivePipelineId, "pipeline-id", "", "", "Pipeline ID")
	_medialiveCmd.Flags().StringVarP(&_medialivePipelineIds, "pipeline-ids", "", "", "Pipeline Ids")
	_medialiveCmd.Flags().StringVarP(&_medialiveProgramName, "program-name", "", "", "Program Name")
	_medialiveCmd.Flags().StringVarP(&_medialiveRenewalSettings, "renewal-settings", "", "", "Renewal Settings")
	_medialiveCmd.Flags().StringVarP(&_medialiveRequestId, "request-id", "", "", "Request ID")
	_medialiveCmd.Flags().StringVarP(&_medialiveReservationId, "reservation-id", "", "", "Reservation ID")
	_medialiveCmd.Flags().StringVarP(&_medialiveReserved, "reserved", "", "", "Reserved")
	_medialiveCmd.Flags().StringVarP(&_medialiveResolution, "resolution", "", "", "Resolution")
	_medialiveCmd.Flags().StringVarP(&_medialiveResourceArn, "resource-arn", "", "", "Resource ARN")
	_medialiveCmd.Flags().StringVarP(&_medialiveResourceType, "resource-type", "", "", "Resource Type")
	_medialiveCmd.Flags().StringVarP(&_medialiveRole, "role", "", "", "Role")
	_medialiveCmd.Flags().StringVarP(&_medialiveRoleArn, "role-arn", "", "", "Role ARN")
	_medialiveCmd.Flags().StringVarP(&_medialiveRouterSettings, "router-settings", "", "", "Router Settings")
	_medialiveCmd.Flags().StringVarP(&_medialiveRoutes, "routes", "", "", "Routes")
	_medialiveCmd.Flags().StringVarP(&_medialiveScope, "scope", "", "", "Scope")
	_medialiveCmd.Flags().StringVarP(&_medialiveSdiSourceId, "sdi-source-id", "", "", "Sdi Source ID")
	_medialiveCmd.Flags().StringVarP(&_medialiveSdiSourceMappings, "sdi-source-mappings", "", "", "Sdi Source Mappings")
	_medialiveCmd.Flags().StringSliceVarP(&_medialiveSdiSources, "sdi-sources", "", nil, "Sdi Sources")
	_medialiveCmd.Flags().StringVarP(&_medialiveSignalMapIdentifier, "signal-map-identifier", "", "", "Signal Map Identifier")
	_medialiveCmd.Flags().StringVarP(&_medialiveSmpte2110ReceiverGroupSettings, "smpte2110-receiver-group-settings", "", "", "Smpte2110 Receiver Group Settings")
	_medialiveCmd.Flags().StringVarP(&_medialiveSources, "sources", "", "", "Sources")
	_medialiveCmd.Flags().StringVarP(&_medialiveSpecialFeature, "special-feature", "", "", "Special Feature")
	_medialiveCmd.Flags().StringVarP(&_medialiveSpecialRouterSettings, "special-router-settings", "", "", "Special Router Settings")
	_medialiveCmd.Flags().StringVarP(&_medialiveSrtSettings, "srt-settings", "", "", "Srt Settings")
	_medialiveCmd.Flags().StringVarP(&_medialiveStart, "start", "", "", "Start")
	_medialiveCmd.Flags().StringVarP(&_medialiveState, "state", "", "", "State")
	_medialiveCmd.Flags().StringVarP(&_medialiveStateFilter, "state-filter", "", "", "State Filter")
	_medialiveCmd.Flags().StringVarP(&_medialiveStatistic, "statistic", "", "", "Statistic")
	_medialiveCmd.Flags().StringSliceVarP(&_medialiveTagKeys, "tag-keys", "", nil, "Tag Keys")
	_medialiveCmd.Flags().StringVarP(&_medialiveTags, "tags", "", "", "Tags")
	_medialiveCmd.Flags().StringVarP(&_medialiveTargetCustomerId, "target-customer-id", "", "", "Target Customer ID")
	_medialiveCmd.Flags().StringVarP(&_medialiveTargetRegion, "target-region", "", "", "Target Region")
	_medialiveCmd.Flags().StringVarP(&_medialiveTargetResourceType, "target-resource-type", "", "", "Target Resource Type")
	_medialiveCmd.Flags().StringVarP(&_medialiveThreshold, "threshold", "", "", "Threshold")
	_medialiveCmd.Flags().StringVarP(&_medialiveThumbnailType, "thumbnail-type", "", "", "Thumbnail Type")
	_medialiveCmd.Flags().StringVarP(&_medialiveTransferMessage, "transfer-message", "", "", "Transfer Message")
	_medialiveCmd.Flags().StringVarP(&_medialiveTransferType, "transfer-type", "", "", "Transfer Type")
	_medialiveCmd.Flags().StringVarP(&_medialiveTreatMissingData, "treat-missing-data", "", "", "Treat Missing Data")
	_medialiveCmd.Flags().StringVarP(&_medialiveType, "type", "", "", "Type")
	_medialiveCmd.Flags().StringVarP(&_medialiveUhdDeviceSettings, "uhd-device-settings", "", "", "Uhd Device Settings")
	_medialiveCmd.Flags().StringVarP(&_medialiveVideoQuality, "video-quality", "", "", "Video Quality")
	_medialiveCmd.Flags().StringVarP(&_medialiveVpc, "vpc", "", "", "VPC")
	_medialiveCmd.Flags().StringVarP(&_medialiveWhitelistRules, "whitelist-rules", "", "", "Whitelist Rules")

	_medialiveCmd.Flags().BoolVarP(&_medialiveAcceptInputDeviceTransfer, "accept-input-device-transfer", "", false, "Accept Input Device Transfer")
	_medialiveCmd.Flags().BoolVarP(&_medialiveBatchDelete, "batch-delete", "", false, "Batch Delete")
	_medialiveCmd.Flags().BoolVarP(&_medialiveBatchStart, "batch-start", "", false, "Batch Start")
	_medialiveCmd.Flags().BoolVarP(&_medialiveBatchStop, "batch-stop", "", false, "Batch Stop")
	_medialiveCmd.Flags().BoolVarP(&_medialiveBatchUpdateSchedule, "batch-update-schedule", "", false, "Batch Update Schedule")
	_medialiveCmd.Flags().BoolVarP(&_medialiveCancelInputDeviceTransfer, "cancel-input-device-transfer", "", false, "Cancel Input Device Transfer")
	_medialiveCmd.Flags().BoolVarP(&_medialiveClaimDevice, "claim-device", "", false, "Claim Device")
	_medialiveCmd.Flags().BoolVarP(&_medialiveCreateChannel, "create-channel", "", false, "Create Channel")
	_medialiveCmd.Flags().BoolVarP(&_medialiveCreateChannelPlacementGroup, "create-channel-placement-group", "", false, "Create Channel Placement Group")
	_medialiveCmd.Flags().BoolVarP(&_medialiveCreateCloudWatchAlarmTemplate, "create-cloud-watch-alarm-template", "", false, "Create Cloud Watch Alarm Template")
	_medialiveCmd.Flags().BoolVarP(&_medialiveCreateCloudWatchAlarmTemplateGroup, "create-cloud-watch-alarm-template-group", "", false, "Create Cloud Watch Alarm Template Group")
	_medialiveCmd.Flags().BoolVarP(&_medialiveCreateCluster, "create-cluster", "", false, "Create Cluster")
	_medialiveCmd.Flags().BoolVarP(&_medialiveCreateEventBridgeRuleTemplate, "create-event-bridge-rule-template", "", false, "Create Event Bridge Rule Template")
	_medialiveCmd.Flags().BoolVarP(&_medialiveCreateEventBridgeRuleTemplateGroup, "create-event-bridge-rule-template-group", "", false, "Create Event Bridge Rule Template Group")
	_medialiveCmd.Flags().BoolVarP(&_medialiveCreateInput, "create-input", "", false, "Create Input")
	_medialiveCmd.Flags().BoolVarP(&_medialiveCreateInputSecurityGroup, "create-input-security-group", "", false, "Create Input Security Group")
	_medialiveCmd.Flags().BoolVarP(&_medialiveCreateMultiplex, "create-multiplex", "", false, "Create Multiplex")
	_medialiveCmd.Flags().BoolVarP(&_medialiveCreateMultiplexProgram, "create-multiplex-program", "", false, "Create Multiplex Program")
	_medialiveCmd.Flags().BoolVarP(&_medialiveCreateNetwork, "create-network", "", false, "Create Network")
	_medialiveCmd.Flags().BoolVarP(&_medialiveCreateNode, "create-node", "", false, "Create Node")
	_medialiveCmd.Flags().BoolVarP(&_medialiveCreateNodeRegistrationScript, "create-node-registration-script", "", false, "Create Node Registration Script")
	_medialiveCmd.Flags().BoolVarP(&_medialiveCreatePartnerInput, "create-partner-input", "", false, "Create Partner Input")
	_medialiveCmd.Flags().BoolVarP(&_medialiveCreateSdiSource, "create-sdi-source", "", false, "Create Sdi Source")
	_medialiveCmd.Flags().BoolVarP(&_medialiveCreateSignalMap, "create-signal-map", "", false, "Create Signal Map")
	_medialiveCmd.Flags().BoolVarP(&_medialiveCreateTags, "create-tags", "", false, "Create Tags")
	_medialiveCmd.Flags().BoolVarP(&_medialiveDeleteChannel, "delete-channel", "", false, "Delete Channel")
	_medialiveCmd.Flags().BoolVarP(&_medialiveDeleteChannelPlacementGroup, "delete-channel-placement-group", "", false, "Delete Channel Placement Group")
	_medialiveCmd.Flags().BoolVarP(&_medialiveDeleteCloudWatchAlarmTemplate, "delete-cloud-watch-alarm-template", "", false, "Delete Cloud Watch Alarm Template")
	_medialiveCmd.Flags().BoolVarP(&_medialiveDeleteCloudWatchAlarmTemplateGroup, "delete-cloud-watch-alarm-template-group", "", false, "Delete Cloud Watch Alarm Template Group")
	_medialiveCmd.Flags().BoolVarP(&_medialiveDeleteCluster, "delete-cluster", "", false, "Delete Cluster")
	_medialiveCmd.Flags().BoolVarP(&_medialiveDeleteEventBridgeRuleTemplate, "delete-event-bridge-rule-template", "", false, "Delete Event Bridge Rule Template")
	_medialiveCmd.Flags().BoolVarP(&_medialiveDeleteEventBridgeRuleTemplateGroup, "delete-event-bridge-rule-template-group", "", false, "Delete Event Bridge Rule Template Group")
	_medialiveCmd.Flags().BoolVarP(&_medialiveDeleteInput, "delete-input", "", false, "Delete Input")
	_medialiveCmd.Flags().BoolVarP(&_medialiveDeleteInputSecurityGroup, "delete-input-security-group", "", false, "Delete Input Security Group")
	_medialiveCmd.Flags().BoolVarP(&_medialiveDeleteMultiplex, "delete-multiplex", "", false, "Delete Multiplex")
	_medialiveCmd.Flags().BoolVarP(&_medialiveDeleteMultiplexProgram, "delete-multiplex-program", "", false, "Delete Multiplex Program")
	_medialiveCmd.Flags().BoolVarP(&_medialiveDeleteNetwork, "delete-network", "", false, "Delete Network")
	_medialiveCmd.Flags().BoolVarP(&_medialiveDeleteNode, "delete-node", "", false, "Delete Node")
	_medialiveCmd.Flags().BoolVarP(&_medialiveDeleteReservation, "delete-reservation", "", false, "Delete Reservation")
	_medialiveCmd.Flags().BoolVarP(&_medialiveDeleteSchedule, "delete-schedule", "", false, "Delete Schedule")
	_medialiveCmd.Flags().BoolVarP(&_medialiveDeleteSdiSource, "delete-sdi-source", "", false, "Delete Sdi Source")
	_medialiveCmd.Flags().BoolVarP(&_medialiveDeleteSignalMap, "delete-signal-map", "", false, "Delete Signal Map")
	_medialiveCmd.Flags().BoolVarP(&_medialiveDeleteTags, "delete-tags", "", false, "Delete Tags")
	_medialiveCmd.Flags().BoolVarP(&_medialiveDescribeAccountConfiguration, "describe-account-configuration", "", false, "Describe Account Configuration")
	_medialiveCmd.Flags().BoolVarP(&_medialiveDescribeChannel, "describe-channel", "", false, "Describe Channel")
	_medialiveCmd.Flags().BoolVarP(&_medialiveDescribeChannelPlacementGroup, "describe-channel-placement-group", "", false, "Describe Channel Placement Group")
	_medialiveCmd.Flags().BoolVarP(&_medialiveDescribeCluster, "describe-cluster", "", false, "Describe Cluster")
	_medialiveCmd.Flags().BoolVarP(&_medialiveDescribeInput, "describe-input", "", false, "Describe Input")
	_medialiveCmd.Flags().BoolVarP(&_medialiveDescribeInputDevice, "describe-input-device", "", false, "Describe Input Device")
	_medialiveCmd.Flags().BoolVarP(&_medialiveDescribeInputDeviceThumbnail, "describe-input-device-thumbnail", "", false, "Describe Input Device Thumbnail")
	_medialiveCmd.Flags().BoolVarP(&_medialiveDescribeInputSecurityGroup, "describe-input-security-group", "", false, "Describe Input Security Group")
	_medialiveCmd.Flags().BoolVarP(&_medialiveDescribeMultiplex, "describe-multiplex", "", false, "Describe Multiplex")
	_medialiveCmd.Flags().BoolVarP(&_medialiveDescribeMultiplexProgram, "describe-multiplex-program", "", false, "Describe Multiplex Program")
	_medialiveCmd.Flags().BoolVarP(&_medialiveDescribeNetwork, "describe-network", "", false, "Describe Network")
	_medialiveCmd.Flags().BoolVarP(&_medialiveDescribeNode, "describe-node", "", false, "Describe Node")
	_medialiveCmd.Flags().BoolVarP(&_medialiveDescribeOffering, "describe-offering", "", false, "Describe Offering")
	_medialiveCmd.Flags().BoolVarP(&_medialiveDescribeReservation, "describe-reservation", "", false, "Describe Reservation")
	_medialiveCmd.Flags().BoolVarP(&_medialiveDescribeSchedule, "describe-schedule", "", false, "Describe Schedule")
	_medialiveCmd.Flags().BoolVarP(&_medialiveDescribeSdiSource, "describe-sdi-source", "", false, "Describe Sdi Source")
	_medialiveCmd.Flags().BoolVarP(&_medialiveDescribeThumbnails, "describe-thumbnails", "", false, "Describe Thumbnails")
	_medialiveCmd.Flags().BoolVarP(&_medialiveGetCloudWatchAlarmTemplate, "get-cloud-watch-alarm-template", "", false, "Get Cloud Watch Alarm Template")
	_medialiveCmd.Flags().BoolVarP(&_medialiveGetCloudWatchAlarmTemplateGroup, "get-cloud-watch-alarm-template-group", "", false, "Get Cloud Watch Alarm Template Group")
	_medialiveCmd.Flags().BoolVarP(&_medialiveGetEventBridgeRuleTemplate, "get-event-bridge-rule-template", "", false, "Get Event Bridge Rule Template")
	_medialiveCmd.Flags().BoolVarP(&_medialiveGetEventBridgeRuleTemplateGroup, "get-event-bridge-rule-template-group", "", false, "Get Event Bridge Rule Template Group")
	_medialiveCmd.Flags().BoolVarP(&_medialiveGetSignalMap, "get-signal-map", "", false, "Get Signal Map")
	_medialiveCmd.Flags().BoolVarP(&_medialiveListAlerts, "list-alerts", "", false, "List Alerts")
	_medialiveCmd.Flags().BoolVarP(&_medialiveListChannelPlacementGroups, "list-channel-placement-groups", "", false, "List Channel Placement Groups")
	_medialiveCmd.Flags().BoolVarP(&_medialiveListChannels, "list-channels", "", false, "List Channels")
	_medialiveCmd.Flags().BoolVarP(&_medialiveListCloudWatchAlarmTemplateGroups, "list-cloud-watch-alarm-template-groups", "", false, "List Cloud Watch Alarm Template Groups")
	_medialiveCmd.Flags().BoolVarP(&_medialiveListCloudWatchAlarmTemplates, "list-cloud-watch-alarm-templates", "", false, "List Cloud Watch Alarm Templates")
	_medialiveCmd.Flags().BoolVarP(&_medialiveListClusterAlerts, "list-cluster-alerts", "", false, "List Cluster Alerts")
	_medialiveCmd.Flags().BoolVarP(&_medialiveListClusters, "list-clusters", "", false, "List Clusters")
	_medialiveCmd.Flags().BoolVarP(&_medialiveListEventBridgeRuleTemplateGroups, "list-event-bridge-rule-template-groups", "", false, "List Event Bridge Rule Template Groups")
	_medialiveCmd.Flags().BoolVarP(&_medialiveListEventBridgeRuleTemplates, "list-event-bridge-rule-templates", "", false, "List Event Bridge Rule Templates")
	_medialiveCmd.Flags().BoolVarP(&_medialiveListInputDeviceTransfers, "list-input-device-transfers", "", false, "List Input Device Transfers")
	_medialiveCmd.Flags().BoolVarP(&_medialiveListInputDevices, "list-input-devices", "", false, "List Input Devices")
	_medialiveCmd.Flags().BoolVarP(&_medialiveListInputSecurityGroups, "list-input-security-groups", "", false, "List Input Security Groups")
	_medialiveCmd.Flags().BoolVarP(&_medialiveListInputs, "list-inputs", "", false, "List Inputs")
	_medialiveCmd.Flags().BoolVarP(&_medialiveListMultiplexAlerts, "list-multiplex-alerts", "", false, "List Multiplex Alerts")
	_medialiveCmd.Flags().BoolVarP(&_medialiveListMultiplexPrograms, "list-multiplex-programs", "", false, "List Multiplex Programs")
	_medialiveCmd.Flags().BoolVarP(&_medialiveListMultiplexes, "list-multiplexes", "", false, "List Multiplexes")
	_medialiveCmd.Flags().BoolVarP(&_medialiveListNetworks, "list-networks", "", false, "List Networks")
	_medialiveCmd.Flags().BoolVarP(&_medialiveListNodes, "list-nodes", "", false, "List Nodes")
	_medialiveCmd.Flags().BoolVarP(&_medialiveListOfferings, "list-offerings", "", false, "List Offerings")
	_medialiveCmd.Flags().BoolVarP(&_medialiveListReservations, "list-reservations", "", false, "List Reservations")
	_medialiveCmd.Flags().BoolVarP(&_medialiveListSdiSources, "list-sdi-sources", "", false, "List Sdi Sources")
	_medialiveCmd.Flags().BoolVarP(&_medialiveListSignalMaps, "list-signal-maps", "", false, "List Signal Maps")
	_medialiveCmd.Flags().BoolVarP(&_medialiveListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_medialiveCmd.Flags().BoolVarP(&_medialiveListVersions, "list-versions", "", false, "List Versions")
	_medialiveCmd.Flags().BoolVarP(&_medialivePurchaseOffering, "purchase-offering", "", false, "Purchase Offering")
	_medialiveCmd.Flags().BoolVarP(&_medialiveRebootInputDevice, "reboot-input-device", "", false, "Reboot Input Device")
	_medialiveCmd.Flags().BoolVarP(&_medialiveRejectInputDeviceTransfer, "reject-input-device-transfer", "", false, "Reject Input Device Transfer")
	_medialiveCmd.Flags().BoolVarP(&_medialiveRestartChannelPipelines, "restart-channel-pipelines", "", false, "Restart Channel Pipelines")
	_medialiveCmd.Flags().BoolVarP(&_medialiveStartChannel, "start-channel", "", false, "Start Channel")
	_medialiveCmd.Flags().BoolVarP(&_medialiveStartDeleteMonitorDeployment, "start-delete-monitor-deployment", "", false, "Start Delete Monitor Deployment")
	_medialiveCmd.Flags().BoolVarP(&_medialiveStartInputDevice, "start-input-device", "", false, "Start Input Device")
	_medialiveCmd.Flags().BoolVarP(&_medialiveStartInputDeviceMaintenanceWindow, "start-input-device-maintenance-window", "", false, "Start Input Device Maintenance Window")
	_medialiveCmd.Flags().BoolVarP(&_medialiveStartMonitorDeployment, "start-monitor-deployment", "", false, "Start Monitor Deployment")
	_medialiveCmd.Flags().BoolVarP(&_medialiveStartMultiplex, "start-multiplex", "", false, "Start Multiplex")
	_medialiveCmd.Flags().BoolVarP(&_medialiveStartUpdateSignalMap, "start-update-signal-map", "", false, "Start Update Signal Map")
	_medialiveCmd.Flags().BoolVarP(&_medialiveStopChannel, "stop-channel", "", false, "Stop Channel")
	_medialiveCmd.Flags().BoolVarP(&_medialiveStopInputDevice, "stop-input-device", "", false, "Stop Input Device")
	_medialiveCmd.Flags().BoolVarP(&_medialiveStopMultiplex, "stop-multiplex", "", false, "Stop Multiplex")
	_medialiveCmd.Flags().BoolVarP(&_medialiveTransferInputDevice, "transfer-input-device", "", false, "Transfer Input Device")
	_medialiveCmd.Flags().BoolVarP(&_medialiveUpdateAccountConfiguration, "update-account-configuration", "", false, "Update Account Configuration")
	_medialiveCmd.Flags().BoolVarP(&_medialiveUpdateChannel, "update-channel", "", false, "Update Channel")
	_medialiveCmd.Flags().BoolVarP(&_medialiveUpdateChannelClass, "update-channel-class", "", false, "Update Channel Class")
	_medialiveCmd.Flags().BoolVarP(&_medialiveUpdateChannelPlacementGroup, "update-channel-placement-group", "", false, "Update Channel Placement Group")
	_medialiveCmd.Flags().BoolVarP(&_medialiveUpdateCloudWatchAlarmTemplate, "update-cloud-watch-alarm-template", "", false, "Update Cloud Watch Alarm Template")
	_medialiveCmd.Flags().BoolVarP(&_medialiveUpdateCloudWatchAlarmTemplateGroup, "update-cloud-watch-alarm-template-group", "", false, "Update Cloud Watch Alarm Template Group")
	_medialiveCmd.Flags().BoolVarP(&_medialiveUpdateCluster, "update-cluster", "", false, "Update Cluster")
	_medialiveCmd.Flags().BoolVarP(&_medialiveUpdateEventBridgeRuleTemplate, "update-event-bridge-rule-template", "", false, "Update Event Bridge Rule Template")
	_medialiveCmd.Flags().BoolVarP(&_medialiveUpdateEventBridgeRuleTemplateGroup, "update-event-bridge-rule-template-group", "", false, "Update Event Bridge Rule Template Group")
	_medialiveCmd.Flags().BoolVarP(&_medialiveUpdateInput, "update-input", "", false, "Update Input")
	_medialiveCmd.Flags().BoolVarP(&_medialiveUpdateInputDevice, "update-input-device", "", false, "Update Input Device")
	_medialiveCmd.Flags().BoolVarP(&_medialiveUpdateInputSecurityGroup, "update-input-security-group", "", false, "Update Input Security Group")
	_medialiveCmd.Flags().BoolVarP(&_medialiveUpdateMultiplex, "update-multiplex", "", false, "Update Multiplex")
	_medialiveCmd.Flags().BoolVarP(&_medialiveUpdateMultiplexProgram, "update-multiplex-program", "", false, "Update Multiplex Program")
	_medialiveCmd.Flags().BoolVarP(&_medialiveUpdateNetwork, "update-network", "", false, "Update Network")
	_medialiveCmd.Flags().BoolVarP(&_medialiveUpdateNode, "update-node", "", false, "Update Node")
	_medialiveCmd.Flags().BoolVarP(&_medialiveUpdateNodeState, "update-node-state", "", false, "Update Node State")
	_medialiveCmd.Flags().BoolVarP(&_medialiveUpdateReservation, "update-reservation", "", false, "Update Reservation")
	_medialiveCmd.Flags().BoolVarP(&_medialiveUpdateSdiSource, "update-sdi-source", "", false, "Update Sdi Source")

}
