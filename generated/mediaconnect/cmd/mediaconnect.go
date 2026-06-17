package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mediaconnect"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// mediaconnectCmd represents the mediaconnect command
var _mediaconnectCmd = &cobra.Command{
	Use:   "mediaconnect",
	Short: "AWS mediaconnect CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := mediaconnect.NewFromConfig(cfg)
		if _mediaconnectAddBridgeOutputs {
			mediaconnect_AddBridgeOutputs(cfg, client)
			return
		}
		if _mediaconnectAddBridgeSources {
			mediaconnect_AddBridgeSources(cfg, client)
			return
		}
		if _mediaconnectAddFlowMediaStreams {
			mediaconnect_AddFlowMediaStreams(cfg, client)
			return
		}
		if _mediaconnectAddFlowOutputs {
			mediaconnect_AddFlowOutputs(cfg, client)
			return
		}
		if _mediaconnectAddFlowSources {
			mediaconnect_AddFlowSources(cfg, client)
			return
		}
		if _mediaconnectAddFlowVpcInterfaces {
			mediaconnect_AddFlowVpcInterfaces(cfg, client)
			return
		}
		if _mediaconnectBatchGetRouterInput {
			mediaconnect_BatchGetRouterInput(cfg, client)
			return
		}
		if _mediaconnectBatchGetRouterNetworkInterface {
			mediaconnect_BatchGetRouterNetworkInterface(cfg, client)
			return
		}
		if _mediaconnectBatchGetRouterOutput {
			mediaconnect_BatchGetRouterOutput(cfg, client)
			return
		}
		if _mediaconnectCreateBridge {
			mediaconnect_CreateBridge(cfg, client)
			return
		}
		if _mediaconnectCreateFlow {
			mediaconnect_CreateFlow(cfg, client)
			return
		}
		if _mediaconnectCreateGateway {
			mediaconnect_CreateGateway(cfg, client)
			return
		}
		if _mediaconnectCreateRouterInput {
			mediaconnect_CreateRouterInput(cfg, client)
			return
		}
		if _mediaconnectCreateRouterNetworkInterface {
			mediaconnect_CreateRouterNetworkInterface(cfg, client)
			return
		}
		if _mediaconnectCreateRouterOutput {
			mediaconnect_CreateRouterOutput(cfg, client)
			return
		}
		if _mediaconnectDeleteBridge {
			mediaconnect_DeleteBridge(cfg, client)
			return
		}
		if _mediaconnectDeleteFlow {
			mediaconnect_DeleteFlow(cfg, client)
			return
		}
		if _mediaconnectDeleteGateway {
			mediaconnect_DeleteGateway(cfg, client)
			return
		}
		if _mediaconnectDeleteRouterInput {
			mediaconnect_DeleteRouterInput(cfg, client)
			return
		}
		if _mediaconnectDeleteRouterNetworkInterface {
			mediaconnect_DeleteRouterNetworkInterface(cfg, client)
			return
		}
		if _mediaconnectDeleteRouterOutput {
			mediaconnect_DeleteRouterOutput(cfg, client)
			return
		}
		if _mediaconnectDeregisterGatewayInstance {
			mediaconnect_DeregisterGatewayInstance(cfg, client)
			return
		}
		if _mediaconnectDescribeBridge {
			mediaconnect_DescribeBridge(cfg, client)
			return
		}
		if _mediaconnectDescribeFlow {
			mediaconnect_DescribeFlow(cfg, client)
			return
		}
		if _mediaconnectDescribeFlowSourceMetadata {
			mediaconnect_DescribeFlowSourceMetadata(cfg, client)
			return
		}
		if _mediaconnectDescribeFlowSourceThumbnail {
			mediaconnect_DescribeFlowSourceThumbnail(cfg, client)
			return
		}
		if _mediaconnectDescribeGateway {
			mediaconnect_DescribeGateway(cfg, client)
			return
		}
		if _mediaconnectDescribeGatewayInstance {
			mediaconnect_DescribeGatewayInstance(cfg, client)
			return
		}
		if _mediaconnectDescribeOffering {
			mediaconnect_DescribeOffering(cfg, client)
			return
		}
		if _mediaconnectDescribeReservation {
			mediaconnect_DescribeReservation(cfg, client)
			return
		}
		if _mediaconnectGetRouterInput {
			mediaconnect_GetRouterInput(cfg, client)
			return
		}
		if _mediaconnectGetRouterInputSourceMetadata {
			mediaconnect_GetRouterInputSourceMetadata(cfg, client)
			return
		}
		if _mediaconnectGetRouterInputThumbnail {
			mediaconnect_GetRouterInputThumbnail(cfg, client)
			return
		}
		if _mediaconnectGetRouterNetworkInterface {
			mediaconnect_GetRouterNetworkInterface(cfg, client)
			return
		}
		if _mediaconnectGetRouterOutput {
			mediaconnect_GetRouterOutput(cfg, client)
			return
		}
		if _mediaconnectGrantFlowEntitlements {
			mediaconnect_GrantFlowEntitlements(cfg, client)
			return
		}
		if _mediaconnectListBridges {
			mediaconnect_ListBridges(cfg, client)
			return
		}
		if _mediaconnectListEntitlements {
			mediaconnect_ListEntitlements(cfg, client)
			return
		}
		if _mediaconnectListFlows {
			mediaconnect_ListFlows(cfg, client)
			return
		}
		if _mediaconnectListGatewayInstances {
			mediaconnect_ListGatewayInstances(cfg, client)
			return
		}
		if _mediaconnectListGateways {
			mediaconnect_ListGateways(cfg, client)
			return
		}
		if _mediaconnectListOfferings {
			mediaconnect_ListOfferings(cfg, client)
			return
		}
		if _mediaconnectListReservations {
			mediaconnect_ListReservations(cfg, client)
			return
		}
		if _mediaconnectListRouterInputs {
			mediaconnect_ListRouterInputs(cfg, client)
			return
		}
		if _mediaconnectListRouterNetworkInterfaces {
			mediaconnect_ListRouterNetworkInterfaces(cfg, client)
			return
		}
		if _mediaconnectListRouterOutputs {
			mediaconnect_ListRouterOutputs(cfg, client)
			return
		}
		if _mediaconnectListTagsForGlobalResource {
			mediaconnect_ListTagsForGlobalResource(cfg, client)
			return
		}
		if _mediaconnectListTagsForResource {
			mediaconnect_ListTagsForResource(cfg, client)
			return
		}
		if _mediaconnectPurchaseOffering {
			mediaconnect_PurchaseOffering(cfg, client)
			return
		}
		if _mediaconnectRemoveBridgeOutput {
			mediaconnect_RemoveBridgeOutput(cfg, client)
			return
		}
		if _mediaconnectRemoveBridgeSource {
			mediaconnect_RemoveBridgeSource(cfg, client)
			return
		}
		if _mediaconnectRemoveFlowMediaStream {
			mediaconnect_RemoveFlowMediaStream(cfg, client)
			return
		}
		if _mediaconnectRemoveFlowOutput {
			mediaconnect_RemoveFlowOutput(cfg, client)
			return
		}
		if _mediaconnectRemoveFlowSource {
			mediaconnect_RemoveFlowSource(cfg, client)
			return
		}
		if _mediaconnectRemoveFlowVpcInterface {
			mediaconnect_RemoveFlowVpcInterface(cfg, client)
			return
		}
		if _mediaconnectRestartRouterInput {
			mediaconnect_RestartRouterInput(cfg, client)
			return
		}
		if _mediaconnectRestartRouterOutput {
			mediaconnect_RestartRouterOutput(cfg, client)
			return
		}
		if _mediaconnectRevokeFlowEntitlement {
			mediaconnect_RevokeFlowEntitlement(cfg, client)
			return
		}
		if _mediaconnectStartFlow {
			mediaconnect_StartFlow(cfg, client)
			return
		}
		if _mediaconnectStartRouterInput {
			mediaconnect_StartRouterInput(cfg, client)
			return
		}
		if _mediaconnectStartRouterOutput {
			mediaconnect_StartRouterOutput(cfg, client)
			return
		}
		if _mediaconnectStopFlow {
			mediaconnect_StopFlow(cfg, client)
			return
		}
		if _mediaconnectStopRouterInput {
			mediaconnect_StopRouterInput(cfg, client)
			return
		}
		if _mediaconnectStopRouterOutput {
			mediaconnect_StopRouterOutput(cfg, client)
			return
		}
		if _mediaconnectTagGlobalResource {
			mediaconnect_TagGlobalResource(cfg, client)
			return
		}
		if _mediaconnectTagResource {
			mediaconnect_TagResource(cfg, client)
			return
		}
		if _mediaconnectTakeRouterInput {
			mediaconnect_TakeRouterInput(cfg, client)
			return
		}
		if _mediaconnectUntagGlobalResource {
			mediaconnect_UntagGlobalResource(cfg, client)
			return
		}
		if _mediaconnectUntagResource {
			mediaconnect_UntagResource(cfg, client)
			return
		}
		if _mediaconnectUpdateBridge {
			mediaconnect_UpdateBridge(cfg, client)
			return
		}
		if _mediaconnectUpdateBridgeOutput {
			mediaconnect_UpdateBridgeOutput(cfg, client)
			return
		}
		if _mediaconnectUpdateBridgeSource {
			mediaconnect_UpdateBridgeSource(cfg, client)
			return
		}
		if _mediaconnectUpdateBridgeState {
			mediaconnect_UpdateBridgeState(cfg, client)
			return
		}
		if _mediaconnectUpdateFlow {
			mediaconnect_UpdateFlow(cfg, client)
			return
		}
		if _mediaconnectUpdateFlowEntitlement {
			mediaconnect_UpdateFlowEntitlement(cfg, client)
			return
		}
		if _mediaconnectUpdateFlowMediaStream {
			mediaconnect_UpdateFlowMediaStream(cfg, client)
			return
		}
		if _mediaconnectUpdateFlowOutput {
			mediaconnect_UpdateFlowOutput(cfg, client)
			return
		}
		if _mediaconnectUpdateFlowSource {
			mediaconnect_UpdateFlowSource(cfg, client)
			return
		}
		if _mediaconnectUpdateGatewayInstance {
			mediaconnect_UpdateGatewayInstance(cfg, client)
			return
		}
		if _mediaconnectUpdateRouterInput {
			mediaconnect_UpdateRouterInput(cfg, client)
			return
		}
		if _mediaconnectUpdateRouterNetworkInterface {
			mediaconnect_UpdateRouterNetworkInterface(cfg, client)
			return
		}
		if _mediaconnectUpdateRouterOutput {
			mediaconnect_UpdateRouterOutput(cfg, client)
			return
		}

	},
}

var (
	_mediaconnectAddBridgeOutputs               bool
	_mediaconnectAddBridgeSources               bool
	_mediaconnectAddFlowMediaStreams            bool
	_mediaconnectAddFlowOutputs                 bool
	_mediaconnectAddFlowSources                 bool
	_mediaconnectAddFlowVpcInterfaces           bool
	_mediaconnectBatchGetRouterInput            bool
	_mediaconnectBatchGetRouterNetworkInterface bool
	_mediaconnectBatchGetRouterOutput           bool
	_mediaconnectCreateBridge                   bool
	_mediaconnectCreateFlow                     bool
	_mediaconnectCreateGateway                  bool
	_mediaconnectCreateRouterInput              bool
	_mediaconnectCreateRouterNetworkInterface   bool
	_mediaconnectCreateRouterOutput             bool
	_mediaconnectDeleteBridge                   bool
	_mediaconnectDeleteFlow                     bool
	_mediaconnectDeleteGateway                  bool
	_mediaconnectDeleteRouterInput              bool
	_mediaconnectDeleteRouterNetworkInterface   bool
	_mediaconnectDeleteRouterOutput             bool
	_mediaconnectDeregisterGatewayInstance      bool
	_mediaconnectDescribeBridge                 bool
	_mediaconnectDescribeFlow                   bool
	_mediaconnectDescribeFlowSourceMetadata     bool
	_mediaconnectDescribeFlowSourceThumbnail    bool
	_mediaconnectDescribeGateway                bool
	_mediaconnectDescribeGatewayInstance        bool
	_mediaconnectDescribeOffering               bool
	_mediaconnectDescribeReservation            bool
	_mediaconnectGetRouterInput                 bool
	_mediaconnectGetRouterInputSourceMetadata   bool
	_mediaconnectGetRouterInputThumbnail        bool
	_mediaconnectGetRouterNetworkInterface      bool
	_mediaconnectGetRouterOutput                bool
	_mediaconnectGrantFlowEntitlements          bool
	_mediaconnectListBridges                    bool
	_mediaconnectListEntitlements               bool
	_mediaconnectListFlows                      bool
	_mediaconnectListGatewayInstances           bool
	_mediaconnectListGateways                   bool
	_mediaconnectListOfferings                  bool
	_mediaconnectListReservations               bool
	_mediaconnectListRouterInputs               bool
	_mediaconnectListRouterNetworkInterfaces    bool
	_mediaconnectListRouterOutputs              bool
	_mediaconnectListTagsForGlobalResource      bool
	_mediaconnectListTagsForResource            bool
	_mediaconnectPurchaseOffering               bool
	_mediaconnectRemoveBridgeOutput             bool
	_mediaconnectRemoveBridgeSource             bool
	_mediaconnectRemoveFlowMediaStream          bool
	_mediaconnectRemoveFlowOutput               bool
	_mediaconnectRemoveFlowSource               bool
	_mediaconnectRemoveFlowVpcInterface         bool
	_mediaconnectRestartRouterInput             bool
	_mediaconnectRestartRouterOutput            bool
	_mediaconnectRevokeFlowEntitlement          bool
	_mediaconnectStartFlow                      bool
	_mediaconnectStartRouterInput               bool
	_mediaconnectStartRouterOutput              bool
	_mediaconnectStopFlow                       bool
	_mediaconnectStopRouterInput                bool
	_mediaconnectStopRouterOutput               bool
	_mediaconnectTagGlobalResource              bool
	_mediaconnectTagResource                    bool
	_mediaconnectTakeRouterInput                bool
	_mediaconnectUntagGlobalResource            bool
	_mediaconnectUntagResource                  bool
	_mediaconnectUpdateBridge                   bool
	_mediaconnectUpdateBridgeOutput             bool
	_mediaconnectUpdateBridgeSource             bool
	_mediaconnectUpdateBridgeState              bool
	_mediaconnectUpdateFlow                     bool
	_mediaconnectUpdateFlowEntitlement          bool
	_mediaconnectUpdateFlowMediaStream          bool
	_mediaconnectUpdateFlowOutput               bool
	_mediaconnectUpdateFlowSource               bool
	_mediaconnectUpdateGatewayInstance          bool
	_mediaconnectUpdateRouterInput              bool
	_mediaconnectUpdateRouterNetworkInterface   bool
	_mediaconnectUpdateRouterOutput             bool

	_mediaconnectArn                                string
	_mediaconnectArns                               []string
	_mediaconnectAttributes                         string
	_mediaconnectAvailabilityZone                   string
	_mediaconnectBridgeArn                          string
	_mediaconnectBridgePlacement                    string
	_mediaconnectCidrAllowList                      []string
	_mediaconnectClientToken                        string
	_mediaconnectClockRate                          string
	_mediaconnectConfiguration                      string
	_mediaconnectDecryption                         string
	_mediaconnectDescription                        string
	_mediaconnectDesiredState                       string
	_mediaconnectDestination                        string
	_mediaconnectEgressCidrBlocks                   []string
	_mediaconnectEgressGatewayBridge                string
	_mediaconnectEncodingConfig                     string
	_mediaconnectEncryption                         string
	_mediaconnectEntitlementArn                     string
	_mediaconnectEntitlementStatus                  string
	_mediaconnectEntitlements                       string
	_mediaconnectFilterArn                          string
	_mediaconnectFilters                            string
	_mediaconnectFlowArn                            string
	_mediaconnectFlowSize                           string
	_mediaconnectFlowSource                         string
	_mediaconnectFlowTags                           string
	_mediaconnectForce                              string
	_mediaconnectGatewayArn                         string
	_mediaconnectGatewayBridgeSource                string
	_mediaconnectGatewayInstanceArn                 string
	_mediaconnectIngestPort                         string
	_mediaconnectIngressGatewayBridge               string
	_mediaconnectMaintenance                        string
	_mediaconnectMaintenanceConfiguration           string
	_mediaconnectMaxBitrate                         string
	_mediaconnectMaxLatency                         string
	_mediaconnectMaxResults                         string
	_mediaconnectMaxSyncBuffer                      string
	_mediaconnectMaximumBitrate                     string
	_mediaconnectMediaStreamName                    string
	_mediaconnectMediaStreamOutputConfigurations    string
	_mediaconnectMediaStreamSourceConfigurations    string
	_mediaconnectMediaStreamType                    string
	_mediaconnectMediaStreams                       string
	_mediaconnectMinLatency                         string
	_mediaconnectName                               string
	_mediaconnectNdiConfig                          string
	_mediaconnectNdiProgramName                     string
	_mediaconnectNdiSourceSettings                  string
	_mediaconnectNdiSpeedHqQuality                  string
	_mediaconnectNetworkOutput                      string
	_mediaconnectNetworkSource                      string
	_mediaconnectNetworks                           string
	_mediaconnectNextToken                          string
	_mediaconnectOfferingArn                        string
	_mediaconnectOutputArn                          string
	_mediaconnectOutputName                         string
	_mediaconnectOutputStatus                       string
	_mediaconnectOutputs                            string
	_mediaconnectPlacementArn                       string
	_mediaconnectPort                               string
	_mediaconnectProtocol                           string
	_mediaconnectRegionName                         string
	_mediaconnectRemoteId                           string
	_mediaconnectReservationArn                     string
	_mediaconnectReservationName                    string
	_mediaconnectResourceArn                        string
	_mediaconnectRouterInputArn                     string
	_mediaconnectRouterIntegrationState             string
	_mediaconnectRouterIntegrationTransitDecryption string
	_mediaconnectRouterIntegrationTransitEncryption string
	_mediaconnectRouterOutputArn                    string
	_mediaconnectRoutingScope                       string
	_mediaconnectSenderControlPort                  string
	_mediaconnectSenderIpAddress                    string
	_mediaconnectSmoothingLatency                   string
	_mediaconnectSource                             string
	_mediaconnectSourceArn                          string
	_mediaconnectSourceFailoverConfig               string
	_mediaconnectSourceListenerAddress              string
	_mediaconnectSourceListenerPort                 string
	_mediaconnectSourceMonitoringConfig             string
	_mediaconnectSourceName                         string
	_mediaconnectSources                            string
	_mediaconnectStart                              string
	_mediaconnectStreamId                           string
	_mediaconnectSubscribers                        []string
	_mediaconnectTagKeys                            []string
	_mediaconnectTags                               string
	_mediaconnectTier                               string
	_mediaconnectTransitEncryption                  string
	_mediaconnectVideoFormat                        string
	_mediaconnectVpcInterfaceAttachment             string
	_mediaconnectVpcInterfaceName                   string
	_mediaconnectVpcInterfaces                      string
	_mediaconnectWhitelistCidr                      string
)

// Adds outputs to an existing bridge.
func mediaconnect_AddBridgeOutputs(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.AddBridgeOutputsInput{
		// BridgeArn: *string, // Required
		// Outputs: []types.AddBridgeOutputRequest, // Required
	}

	if len(_mediaconnectBridgeArn) > 0 {
		input.BridgeArn = aws.String(_mediaconnectBridgeArn)
	}
	if len(_mediaconnectOutputs) > 0 {
		if err := assignInputField(input, "Outputs", _mediaconnectOutputs); err != nil {
			log.Errorf("invalid --outputs: %s", err.Error())
			return
		}
	}

	if resp, err := client.AddBridgeOutputs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds sources to an existing bridge.
func mediaconnect_AddBridgeSources(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.AddBridgeSourcesInput{
		// BridgeArn: *string, // Required
		// Sources: []types.AddBridgeSourceRequest, // Required
	}

	if len(_mediaconnectBridgeArn) > 0 {
		input.BridgeArn = aws.String(_mediaconnectBridgeArn)
	}
	if len(_mediaconnectSources) > 0 {
		if err := assignInputField(input, "Sources", _mediaconnectSources); err != nil {
			log.Errorf("invalid --sources: %s", err.Error())
			return
		}
	}

	if resp, err := client.AddBridgeSources(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds media streams to an existing flow. After you add a media stream to a
// flow, you can associate it with a source and/or an output that uses the ST 2110
// JPEG XS or CDI protocol.
func mediaconnect_AddFlowMediaStreams(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.AddFlowMediaStreamsInput{
		// FlowArn: *string, // Required
		// MediaStreams: []types.AddMediaStreamRequest, // Required
	}

	if len(_mediaconnectFlowArn) > 0 {
		input.FlowArn = aws.String(_mediaconnectFlowArn)
	}
	if len(_mediaconnectMediaStreams) > 0 {
		if err := assignInputField(input, "MediaStreams", _mediaconnectMediaStreams); err != nil {
			log.Errorf("invalid --media-streams: %s", err.Error())
			return
		}
	}

	if resp, err := client.AddFlowMediaStreams(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds outputs to an existing flow. You can create up to 50 outputs per flow.
func mediaconnect_AddFlowOutputs(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.AddFlowOutputsInput{
		// FlowArn: *string, // Required
		// Outputs: []types.AddOutputRequest, // Required
	}

	if len(_mediaconnectFlowArn) > 0 {
		input.FlowArn = aws.String(_mediaconnectFlowArn)
	}
	if len(_mediaconnectOutputs) > 0 {
		if err := assignInputField(input, "Outputs", _mediaconnectOutputs); err != nil {
			log.Errorf("invalid --outputs: %s", err.Error())
			return
		}
	}

	if resp, err := client.AddFlowOutputs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds sources to a flow.
func mediaconnect_AddFlowSources(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.AddFlowSourcesInput{
		// FlowArn: *string, // Required
		// Sources: []types.SetSourceRequest, // Required
	}

	if len(_mediaconnectFlowArn) > 0 {
		input.FlowArn = aws.String(_mediaconnectFlowArn)
	}
	if len(_mediaconnectSources) > 0 {
		if err := assignInputField(input, "Sources", _mediaconnectSources); err != nil {
			log.Errorf("invalid --sources: %s", err.Error())
			return
		}
	}

	if resp, err := client.AddFlowSources(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds VPC interfaces to a flow.
func mediaconnect_AddFlowVpcInterfaces(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.AddFlowVpcInterfacesInput{
		// FlowArn: *string, // Required
		// VpcInterfaces: []types.VpcInterfaceRequest, // Required
	}

	if len(_mediaconnectFlowArn) > 0 {
		input.FlowArn = aws.String(_mediaconnectFlowArn)
	}
	if len(_mediaconnectVpcInterfaces) > 0 {
		if err := assignInputField(input, "VpcInterfaces", _mediaconnectVpcInterfaces); err != nil {
			log.Errorf("invalid --vpc-interfaces: %s", err.Error())
			return
		}
	}

	if resp, err := client.AddFlowVpcInterfaces(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about multiple router inputs in AWS Elemental
// MediaConnect.
func mediaconnect_BatchGetRouterInput(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.BatchGetRouterInputInput{
		// Arns: []string, // Required
	}

	if len(_mediaconnectArns) > 0 {
		input.Arns = append([]string(nil), _mediaconnectArns...)
	}

	if resp, err := client.BatchGetRouterInput(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about multiple router network interfaces in AWS Elemental
// MediaConnect.
func mediaconnect_BatchGetRouterNetworkInterface(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.BatchGetRouterNetworkInterfaceInput{
		// Arns: []string, // Required
	}

	if len(_mediaconnectArns) > 0 {
		input.Arns = append([]string(nil), _mediaconnectArns...)
	}

	if resp, err := client.BatchGetRouterNetworkInterface(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about multiple router outputs in AWS Elemental
// MediaConnect.
func mediaconnect_BatchGetRouterOutput(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.BatchGetRouterOutputInput{
		// Arns: []string, // Required
	}

	if len(_mediaconnectArns) > 0 {
		input.Arns = append([]string(nil), _mediaconnectArns...)
	}

	if resp, err := client.BatchGetRouterOutput(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new bridge. The request must include one source.
func mediaconnect_CreateBridge(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.CreateBridgeInput{
		// Name: *string, // Required
		// PlacementArn: *string, // Required
		// Sources: []types.AddBridgeSourceRequest, // Required
	}

	if len(_mediaconnectName) > 0 {
		input.Name = aws.String(_mediaconnectName)
	}
	if len(_mediaconnectPlacementArn) > 0 {
		input.PlacementArn = aws.String(_mediaconnectPlacementArn)
	}
	if len(_mediaconnectSources) > 0 {
		if err := assignInputField(input, "Sources", _mediaconnectSources); err != nil {
			log.Errorf("invalid --sources: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectEgressGatewayBridge) > 0 {
		if err := assignInputField(input, "EgressGatewayBridge", _mediaconnectEgressGatewayBridge); err != nil {
			log.Errorf("invalid --egress-gateway-bridge: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectIngressGatewayBridge) > 0 {
		if err := assignInputField(input, "IngressGatewayBridge", _mediaconnectIngressGatewayBridge); err != nil {
			log.Errorf("invalid --ingress-gateway-bridge: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectOutputs) > 0 {
		if err := assignInputField(input, "Outputs", _mediaconnectOutputs); err != nil {
			log.Errorf("invalid --outputs: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectSourceFailoverConfig) > 0 {
		if err := assignInputField(input, "SourceFailoverConfig", _mediaconnectSourceFailoverConfig); err != nil {
			log.Errorf("invalid --source-failover-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateBridge(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new flow. The request must include one source. The request
// optionally can include outputs (up to 50) and entitlements (up to 50).
func mediaconnect_CreateFlow(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.CreateFlowInput{
		// Name: *string, // Required
	}

	if len(_mediaconnectName) > 0 {
		input.Name = aws.String(_mediaconnectName)
	}
	if len(_mediaconnectAvailabilityZone) > 0 {
		input.AvailabilityZone = aws.String(_mediaconnectAvailabilityZone)
	}
	if len(_mediaconnectEncodingConfig) > 0 {
		if err := assignInputField(input, "EncodingConfig", _mediaconnectEncodingConfig); err != nil {
			log.Errorf("invalid --encoding-config: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectEntitlements) > 0 {
		if err := assignInputField(input, "Entitlements", _mediaconnectEntitlements); err != nil {
			log.Errorf("invalid --entitlements: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectFlowSize) > 0 {
		if err := assignInputField(input, "FlowSize", _mediaconnectFlowSize); err != nil {
			log.Errorf("invalid --flow-size: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectFlowTags) > 0 {
		if err := assignInputField(input, "FlowTags", _mediaconnectFlowTags); err != nil {
			log.Errorf("invalid --flow-tags: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectMaintenance) > 0 {
		if err := assignInputField(input, "Maintenance", _mediaconnectMaintenance); err != nil {
			log.Errorf("invalid --maintenance: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectMediaStreams) > 0 {
		if err := assignInputField(input, "MediaStreams", _mediaconnectMediaStreams); err != nil {
			log.Errorf("invalid --media-streams: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectNdiConfig) > 0 {
		if err := assignInputField(input, "NdiConfig", _mediaconnectNdiConfig); err != nil {
			log.Errorf("invalid --ndi-config: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectOutputs) > 0 {
		if err := assignInputField(input, "Outputs", _mediaconnectOutputs); err != nil {
			log.Errorf("invalid --outputs: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectSource) > 0 {
		if err := assignInputField(input, "Source", _mediaconnectSource); err != nil {
			log.Errorf("invalid --source: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectSourceFailoverConfig) > 0 {
		if err := assignInputField(input, "SourceFailoverConfig", _mediaconnectSourceFailoverConfig); err != nil {
			log.Errorf("invalid --source-failover-config: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectSourceMonitoringConfig) > 0 {
		if err := assignInputField(input, "SourceMonitoringConfig", _mediaconnectSourceMonitoringConfig); err != nil {
			log.Errorf("invalid --source-monitoring-config: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectSources) > 0 {
		if err := assignInputField(input, "Sources", _mediaconnectSources); err != nil {
			log.Errorf("invalid --sources: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectVpcInterfaces) > 0 {
		if err := assignInputField(input, "VpcInterfaces", _mediaconnectVpcInterfaces); err != nil {
			log.Errorf("invalid --vpc-interfaces: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateFlow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new gateway. The request must include at least one network (up to
// four).
func mediaconnect_CreateGateway(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.CreateGatewayInput{
		// EgressCidrBlocks: []string, // Required
		// Name: *string, // Required
		// Networks: []types.GatewayNetwork, // Required
	}

	if len(_mediaconnectEgressCidrBlocks) > 0 {
		input.EgressCidrBlocks = append([]string(nil), _mediaconnectEgressCidrBlocks...)
	}
	if len(_mediaconnectName) > 0 {
		input.Name = aws.String(_mediaconnectName)
	}
	if len(_mediaconnectNetworks) > 0 {
		if err := assignInputField(input, "Networks", _mediaconnectNetworks); err != nil {
			log.Errorf("invalid --networks: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateGateway(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new router input in AWS Elemental MediaConnect.
func mediaconnect_CreateRouterInput(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.CreateRouterInputInput{
		// Configuration: types.RouterInputConfiguration, // Required
		// MaximumBitrate: *int64, // Required
		// Name: *string, // Required
		// RoutingScope: types.RoutingScope, // Required
		// Tier: types.RouterInputTier, // Required
	}

	if len(_mediaconnectConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _mediaconnectConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectMaximumBitrate) > 0 {
		if err := assignInputField(input, "MaximumBitrate", _mediaconnectMaximumBitrate); err != nil {
			log.Errorf("invalid --maximum-bitrate: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectName) > 0 {
		input.Name = aws.String(_mediaconnectName)
	}
	if len(_mediaconnectRoutingScope) > 0 {
		if err := assignInputField(input, "RoutingScope", _mediaconnectRoutingScope); err != nil {
			log.Errorf("invalid --routing-scope: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectTier) > 0 {
		if err := assignInputField(input, "Tier", _mediaconnectTier); err != nil {
			log.Errorf("invalid --tier: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectAvailabilityZone) > 0 {
		input.AvailabilityZone = aws.String(_mediaconnectAvailabilityZone)
	}
	if len(_mediaconnectClientToken) > 0 {
		input.ClientToken = aws.String(_mediaconnectClientToken)
	}
	if len(_mediaconnectMaintenanceConfiguration) > 0 {
		if err := assignInputField(input, "MaintenanceConfiguration", _mediaconnectMaintenanceConfiguration); err != nil {
			log.Errorf("invalid --maintenance-configuration: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectRegionName) > 0 {
		input.RegionName = aws.String(_mediaconnectRegionName)
	}
	if len(_mediaconnectTags) > 0 {
		if err := assignInputField(input, "Tags", _mediaconnectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectTransitEncryption) > 0 {
		if err := assignInputField(input, "TransitEncryption", _mediaconnectTransitEncryption); err != nil {
			log.Errorf("invalid --transit-encryption: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRouterInput(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new router network interface in AWS Elemental MediaConnect.
func mediaconnect_CreateRouterNetworkInterface(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.CreateRouterNetworkInterfaceInput{
		// Configuration: types.RouterNetworkInterfaceConfiguration, // Required
		// Name: *string, // Required
	}

	if len(_mediaconnectConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _mediaconnectConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectName) > 0 {
		input.Name = aws.String(_mediaconnectName)
	}
	if len(_mediaconnectClientToken) > 0 {
		input.ClientToken = aws.String(_mediaconnectClientToken)
	}
	if len(_mediaconnectRegionName) > 0 {
		input.RegionName = aws.String(_mediaconnectRegionName)
	}
	if len(_mediaconnectTags) > 0 {
		if err := assignInputField(input, "Tags", _mediaconnectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRouterNetworkInterface(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new router output in AWS Elemental MediaConnect.
func mediaconnect_CreateRouterOutput(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.CreateRouterOutputInput{
		// Configuration: types.RouterOutputConfiguration, // Required
		// MaximumBitrate: *int64, // Required
		// Name: *string, // Required
		// RoutingScope: types.RoutingScope, // Required
		// Tier: types.RouterOutputTier, // Required
	}

	if len(_mediaconnectConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _mediaconnectConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectMaximumBitrate) > 0 {
		if err := assignInputField(input, "MaximumBitrate", _mediaconnectMaximumBitrate); err != nil {
			log.Errorf("invalid --maximum-bitrate: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectName) > 0 {
		input.Name = aws.String(_mediaconnectName)
	}
	if len(_mediaconnectRoutingScope) > 0 {
		if err := assignInputField(input, "RoutingScope", _mediaconnectRoutingScope); err != nil {
			log.Errorf("invalid --routing-scope: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectTier) > 0 {
		if err := assignInputField(input, "Tier", _mediaconnectTier); err != nil {
			log.Errorf("invalid --tier: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectAvailabilityZone) > 0 {
		input.AvailabilityZone = aws.String(_mediaconnectAvailabilityZone)
	}
	if len(_mediaconnectClientToken) > 0 {
		input.ClientToken = aws.String(_mediaconnectClientToken)
	}
	if len(_mediaconnectMaintenanceConfiguration) > 0 {
		if err := assignInputField(input, "MaintenanceConfiguration", _mediaconnectMaintenanceConfiguration); err != nil {
			log.Errorf("invalid --maintenance-configuration: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectRegionName) > 0 {
		input.RegionName = aws.String(_mediaconnectRegionName)
	}
	if len(_mediaconnectTags) > 0 {
		if err := assignInputField(input, "Tags", _mediaconnectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRouterOutput(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a bridge. Before you can delete a bridge, you must stop the bridge.
func mediaconnect_DeleteBridge(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.DeleteBridgeInput{
		// BridgeArn: *string, // Required
	}

	if len(_mediaconnectBridgeArn) > 0 {
		input.BridgeArn = aws.String(_mediaconnectBridgeArn)
	}

	if resp, err := client.DeleteBridge(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a flow. Before you can delete a flow, you must stop the flow.
func mediaconnect_DeleteFlow(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.DeleteFlowInput{
		// FlowArn: *string, // Required
	}

	if len(_mediaconnectFlowArn) > 0 {
		input.FlowArn = aws.String(_mediaconnectFlowArn)
	}

	if resp, err := client.DeleteFlow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a gateway. Before you can delete a gateway, you must deregister its
// instances and delete its bridges.
func mediaconnect_DeleteGateway(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.DeleteGatewayInput{
		// GatewayArn: *string, // Required
	}

	if len(_mediaconnectGatewayArn) > 0 {
		input.GatewayArn = aws.String(_mediaconnectGatewayArn)
	}

	if resp, err := client.DeleteGateway(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a router input from AWS Elemental MediaConnect.
func mediaconnect_DeleteRouterInput(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.DeleteRouterInputInput{
		// Arn: *string, // Required
	}

	if len(_mediaconnectArn) > 0 {
		input.Arn = aws.String(_mediaconnectArn)
	}

	if resp, err := client.DeleteRouterInput(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a router network interface from AWS Elemental MediaConnect.
func mediaconnect_DeleteRouterNetworkInterface(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.DeleteRouterNetworkInterfaceInput{
		// Arn: *string, // Required
	}

	if len(_mediaconnectArn) > 0 {
		input.Arn = aws.String(_mediaconnectArn)
	}

	if resp, err := client.DeleteRouterNetworkInterface(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a router output from AWS Elemental MediaConnect.
func mediaconnect_DeleteRouterOutput(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.DeleteRouterOutputInput{
		// Arn: *string, // Required
	}

	if len(_mediaconnectArn) > 0 {
		input.Arn = aws.String(_mediaconnectArn)
	}

	if resp, err := client.DeleteRouterOutput(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deregisters an instance. Before you deregister an instance, all bridges
// running on the instance must be stopped. If you want to deregister an instance
// without stopping the bridges, you must use the --force option.
func mediaconnect_DeregisterGatewayInstance(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.DeregisterGatewayInstanceInput{
		// GatewayInstanceArn: *string, // Required
	}

	if len(_mediaconnectGatewayInstanceArn) > 0 {
		input.GatewayInstanceArn = aws.String(_mediaconnectGatewayInstanceArn)
	}
	if len(_mediaconnectForce) > 0 {
		if err := assignInputField(input, "Force", _mediaconnectForce); err != nil {
			log.Errorf("invalid --force: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeregisterGatewayInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Displays the details of a bridge.
func mediaconnect_DescribeBridge(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.DescribeBridgeInput{
		// BridgeArn: *string, // Required
	}

	if len(_mediaconnectBridgeArn) > 0 {
		input.BridgeArn = aws.String(_mediaconnectBridgeArn)
	}

	if resp, err := client.DescribeBridge(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Displays the details of a flow. The response includes the flow Amazon Resource
// Name (ARN), name, and Availability Zone, as well as details about the source,
// outputs, and entitlements.
func mediaconnect_DescribeFlow(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.DescribeFlowInput{
		// FlowArn: *string, // Required
	}

	if len(_mediaconnectFlowArn) > 0 {
		input.FlowArn = aws.String(_mediaconnectFlowArn)
	}

	if resp, err := client.DescribeFlow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The DescribeFlowSourceMetadata API is used to view information about the
// flow's source transport stream and programs. This API displays status messages
// about the flow's source as well as details about the program's video, audio, and
// other data.
func mediaconnect_DescribeFlowSourceMetadata(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.DescribeFlowSourceMetadataInput{
		// FlowArn: *string, // Required
	}

	if len(_mediaconnectFlowArn) > 0 {
		input.FlowArn = aws.String(_mediaconnectFlowArn)
	}

	if resp, err := client.DescribeFlowSourceMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the thumbnail for the flow source.
func mediaconnect_DescribeFlowSourceThumbnail(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.DescribeFlowSourceThumbnailInput{
		// FlowArn: *string, // Required
	}

	if len(_mediaconnectFlowArn) > 0 {
		input.FlowArn = aws.String(_mediaconnectFlowArn)
	}

	if resp, err := client.DescribeFlowSourceThumbnail(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Displays the details of a gateway. The response includes the gateway Amazon
// Resource Name (ARN), name, and CIDR blocks, as well as details about the
// networks.
func mediaconnect_DescribeGateway(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.DescribeGatewayInput{
		// GatewayArn: *string, // Required
	}

	if len(_mediaconnectGatewayArn) > 0 {
		input.GatewayArn = aws.String(_mediaconnectGatewayArn)
	}

	if resp, err := client.DescribeGateway(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Displays the details of an instance.
func mediaconnect_DescribeGatewayInstance(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.DescribeGatewayInstanceInput{
		// GatewayInstanceArn: *string, // Required
	}

	if len(_mediaconnectGatewayInstanceArn) > 0 {
		input.GatewayInstanceArn = aws.String(_mediaconnectGatewayInstanceArn)
	}

	if resp, err := client.DescribeGatewayInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Displays the details of an offering. The response includes the offering
// description, duration, outbound bandwidth, price, and Amazon Resource Name
// (ARN).
func mediaconnect_DescribeOffering(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.DescribeOfferingInput{
		// OfferingArn: *string, // Required
	}

	if len(_mediaconnectOfferingArn) > 0 {
		input.OfferingArn = aws.String(_mediaconnectOfferingArn)
	}

	if resp, err := client.DescribeOffering(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Displays the details of a reservation. The response includes the reservation
// name, state, start date and time, and the details of the offering that make up
// the rest of the reservation (such as price, duration, and outbound bandwidth).
func mediaconnect_DescribeReservation(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.DescribeReservationInput{
		// ReservationArn: *string, // Required
	}

	if len(_mediaconnectReservationArn) > 0 {
		input.ReservationArn = aws.String(_mediaconnectReservationArn)
	}

	if resp, err := client.DescribeReservation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a specific router input in AWS Elemental
// MediaConnect.
func mediaconnect_GetRouterInput(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.GetRouterInputInput{
		// Arn: *string, // Required
	}

	if len(_mediaconnectArn) > 0 {
		input.Arn = aws.String(_mediaconnectArn)
	}

	if resp, err := client.GetRouterInput(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves detailed metadata information about a specific router input source,
// including stream details and connection state.
func mediaconnect_GetRouterInputSourceMetadata(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.GetRouterInputSourceMetadataInput{
		// Arn: *string, // Required
	}

	if len(_mediaconnectArn) > 0 {
		input.Arn = aws.String(_mediaconnectArn)
	}

	if resp, err := client.GetRouterInputSourceMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the thumbnail for a router input in AWS Elemental MediaConnect.
func mediaconnect_GetRouterInputThumbnail(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.GetRouterInputThumbnailInput{
		// Arn: *string, // Required
	}

	if len(_mediaconnectArn) > 0 {
		input.Arn = aws.String(_mediaconnectArn)
	}

	if resp, err := client.GetRouterInputThumbnail(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a specific router network interface in AWS
// Elemental MediaConnect.
func mediaconnect_GetRouterNetworkInterface(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.GetRouterNetworkInterfaceInput{
		// Arn: *string, // Required
	}

	if len(_mediaconnectArn) > 0 {
		input.Arn = aws.String(_mediaconnectArn)
	}

	if resp, err := client.GetRouterNetworkInterface(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a specific router output in AWS Elemental
// MediaConnect.
func mediaconnect_GetRouterOutput(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.GetRouterOutputInput{
		// Arn: *string, // Required
	}

	if len(_mediaconnectArn) > 0 {
		input.Arn = aws.String(_mediaconnectArn)
	}

	if resp, err := client.GetRouterOutput(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Grants entitlements to an existing flow.
func mediaconnect_GrantFlowEntitlements(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.GrantFlowEntitlementsInput{
		// Entitlements: []types.GrantEntitlementRequest, // Required
		// FlowArn: *string, // Required
	}

	if len(_mediaconnectEntitlements) > 0 {
		if err := assignInputField(input, "Entitlements", _mediaconnectEntitlements); err != nil {
			log.Errorf("invalid --entitlements: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectFlowArn) > 0 {
		input.FlowArn = aws.String(_mediaconnectFlowArn)
	}

	if resp, err := client.GrantFlowEntitlements(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Displays a list of bridges that are associated with this account and an
// optionally specified Amazon Resource Name (ARN). This request returns a
// paginated result.
func mediaconnect_ListBridges(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.ListBridgesInput{}

	if len(_mediaconnectFilterArn) > 0 {
		input.FilterArn = aws.String(_mediaconnectFilterArn)
	}
	if len(_mediaconnectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mediaconnectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectNextToken) > 0 {
		input.NextToken = aws.String(_mediaconnectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListBridges(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mediaconnect.ListBridgesOutput
	p := mediaconnect.NewListBridgesPaginator(client, input)
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

// Displays a list of all entitlements that have been granted to this account.
// This request returns 20 results per page.
func mediaconnect_ListEntitlements(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.ListEntitlementsInput{}

	if len(_mediaconnectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mediaconnectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectNextToken) > 0 {
		input.NextToken = aws.String(_mediaconnectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEntitlements(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mediaconnect.ListEntitlementsOutput
	p := mediaconnect.NewListEntitlementsPaginator(client, input)
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

// Displays a list of flows that are associated with this account. This request
// returns a paginated result.
func mediaconnect_ListFlows(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.ListFlowsInput{}

	if len(_mediaconnectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mediaconnectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectNextToken) > 0 {
		input.NextToken = aws.String(_mediaconnectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListFlows(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mediaconnect.ListFlowsOutput
	p := mediaconnect.NewListFlowsPaginator(client, input)
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

// Displays a list of instances associated with the Amazon Web Services account.
// This request returns a paginated result. You can use the filterArn property to
// display only the instances associated with the selected Gateway Amazon Resource
// Name (ARN).
func mediaconnect_ListGatewayInstances(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.ListGatewayInstancesInput{}

	if len(_mediaconnectFilterArn) > 0 {
		input.FilterArn = aws.String(_mediaconnectFilterArn)
	}
	if len(_mediaconnectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mediaconnectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectNextToken) > 0 {
		input.NextToken = aws.String(_mediaconnectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListGatewayInstances(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mediaconnect.ListGatewayInstancesOutput
	p := mediaconnect.NewListGatewayInstancesPaginator(client, input)
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

// Displays a list of gateways that are associated with this account. This
// request returns a paginated result.
func mediaconnect_ListGateways(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.ListGatewaysInput{}

	if len(_mediaconnectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mediaconnectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectNextToken) > 0 {
		input.NextToken = aws.String(_mediaconnectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListGateways(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mediaconnect.ListGatewaysOutput
	p := mediaconnect.NewListGatewaysPaginator(client, input)
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

// Displays a list of all offerings that are available to this account in the
// current Amazon Web Services Region. If you have an active reservation (which
// means you've purchased an offering that has already started and hasn't expired
// yet), your account isn't eligible for other offerings.
func mediaconnect_ListOfferings(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.ListOfferingsInput{}

	if len(_mediaconnectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mediaconnectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectNextToken) > 0 {
		input.NextToken = aws.String(_mediaconnectNextToken)
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

	var results []*mediaconnect.ListOfferingsOutput
	p := mediaconnect.NewListOfferingsPaginator(client, input)
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

// Displays a list of all reservations that have been purchased by this account
// in the current Amazon Web Services Region. This list includes all reservations
// in all states (such as active and expired).
func mediaconnect_ListReservations(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.ListReservationsInput{}

	if len(_mediaconnectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mediaconnectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectNextToken) > 0 {
		input.NextToken = aws.String(_mediaconnectNextToken)
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

	var results []*mediaconnect.ListReservationsOutput
	p := mediaconnect.NewListReservationsPaginator(client, input)
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

// Retrieves a list of router inputs in AWS Elemental MediaConnect.
func mediaconnect_ListRouterInputs(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.ListRouterInputsInput{}

	if len(_mediaconnectFilters) > 0 {
		if err := assignInputField(input, "Filters", _mediaconnectFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mediaconnectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectNextToken) > 0 {
		input.NextToken = aws.String(_mediaconnectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRouterInputs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mediaconnect.ListRouterInputsOutput
	p := mediaconnect.NewListRouterInputsPaginator(client, input)
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

// Retrieves a list of router network interfaces in AWS Elemental MediaConnect.
func mediaconnect_ListRouterNetworkInterfaces(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.ListRouterNetworkInterfacesInput{}

	if len(_mediaconnectFilters) > 0 {
		if err := assignInputField(input, "Filters", _mediaconnectFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mediaconnectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectNextToken) > 0 {
		input.NextToken = aws.String(_mediaconnectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRouterNetworkInterfaces(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mediaconnect.ListRouterNetworkInterfacesOutput
	p := mediaconnect.NewListRouterNetworkInterfacesPaginator(client, input)
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

// Retrieves a list of router outputs in AWS Elemental MediaConnect.
func mediaconnect_ListRouterOutputs(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.ListRouterOutputsInput{}

	if len(_mediaconnectFilters) > 0 {
		if err := assignInputField(input, "Filters", _mediaconnectFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mediaconnectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectNextToken) > 0 {
		input.NextToken = aws.String(_mediaconnectNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRouterOutputs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mediaconnect.ListRouterOutputsOutput
	p := mediaconnect.NewListRouterOutputsPaginator(client, input)
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

// Lists the tags associated with a global resource in AWS Elemental MediaConnect.
// The API supports the following global resources: router inputs, router outputs
// and router network interfaces.
func mediaconnect_ListTagsForGlobalResource(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.ListTagsForGlobalResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_mediaconnectResourceArn) > 0 {
		input.ResourceArn = aws.String(_mediaconnectResourceArn)
	}

	if resp, err := client.ListTagsForGlobalResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List all tags on a MediaConnect resource in the current region.
func mediaconnect_ListTagsForResource(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_mediaconnectResourceArn) > 0 {
		input.ResourceArn = aws.String(_mediaconnectResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Submits a request to purchase an offering. If you already have an active
// reservation, you can't purchase another offering.
func mediaconnect_PurchaseOffering(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.PurchaseOfferingInput{
		// OfferingArn: *string, // Required
		// ReservationName: *string, // Required
		// Start: *string, // Required
	}

	if len(_mediaconnectOfferingArn) > 0 {
		input.OfferingArn = aws.String(_mediaconnectOfferingArn)
	}
	if len(_mediaconnectReservationName) > 0 {
		input.ReservationName = aws.String(_mediaconnectReservationName)
	}
	if len(_mediaconnectStart) > 0 {
		input.Start = aws.String(_mediaconnectStart)
	}

	if resp, err := client.PurchaseOffering(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes an output from a bridge.
func mediaconnect_RemoveBridgeOutput(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.RemoveBridgeOutputInput{
		// BridgeArn: *string, // Required
		// OutputName: *string, // Required
	}

	if len(_mediaconnectBridgeArn) > 0 {
		input.BridgeArn = aws.String(_mediaconnectBridgeArn)
	}
	if len(_mediaconnectOutputName) > 0 {
		input.OutputName = aws.String(_mediaconnectOutputName)
	}

	if resp, err := client.RemoveBridgeOutput(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a source from a bridge.
func mediaconnect_RemoveBridgeSource(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.RemoveBridgeSourceInput{
		// BridgeArn: *string, // Required
		// SourceName: *string, // Required
	}

	if len(_mediaconnectBridgeArn) > 0 {
		input.BridgeArn = aws.String(_mediaconnectBridgeArn)
	}
	if len(_mediaconnectSourceName) > 0 {
		input.SourceName = aws.String(_mediaconnectSourceName)
	}

	if resp, err := client.RemoveBridgeSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a media stream from a flow. This action is only available if the media
// stream is not associated with a source or output.
func mediaconnect_RemoveFlowMediaStream(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.RemoveFlowMediaStreamInput{
		// FlowArn: *string, // Required
		// MediaStreamName: *string, // Required
	}

	if len(_mediaconnectFlowArn) > 0 {
		input.FlowArn = aws.String(_mediaconnectFlowArn)
	}
	if len(_mediaconnectMediaStreamName) > 0 {
		input.MediaStreamName = aws.String(_mediaconnectMediaStreamName)
	}

	if resp, err := client.RemoveFlowMediaStream(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes an output from an existing flow. This request can be made only on an
// output that does not have an entitlement associated with it. If the output has
// an entitlement, you must revoke the entitlement instead. When an entitlement is
// revoked from a flow, the service automatically removes the associated output.
func mediaconnect_RemoveFlowOutput(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.RemoveFlowOutputInput{
		// FlowArn: *string, // Required
		// OutputArn: *string, // Required
	}

	if len(_mediaconnectFlowArn) > 0 {
		input.FlowArn = aws.String(_mediaconnectFlowArn)
	}
	if len(_mediaconnectOutputArn) > 0 {
		input.OutputArn = aws.String(_mediaconnectOutputArn)
	}

	if resp, err := client.RemoveFlowOutput(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a source from an existing flow. This request can be made only if there
// is more than one source on the flow.
func mediaconnect_RemoveFlowSource(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.RemoveFlowSourceInput{
		// FlowArn: *string, // Required
		// SourceArn: *string, // Required
	}

	if len(_mediaconnectFlowArn) > 0 {
		input.FlowArn = aws.String(_mediaconnectFlowArn)
	}
	if len(_mediaconnectSourceArn) > 0 {
		input.SourceArn = aws.String(_mediaconnectSourceArn)
	}

	if resp, err := client.RemoveFlowSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a VPC Interface from an existing flow. This request can be made only
// on a VPC interface that does not have a Source or Output associated with it. If
// the VPC interface is referenced by a Source or Output, you must first delete or
// update the Source or Output to no longer reference the VPC interface.
func mediaconnect_RemoveFlowVpcInterface(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.RemoveFlowVpcInterfaceInput{
		// FlowArn: *string, // Required
		// VpcInterfaceName: *string, // Required
	}

	if len(_mediaconnectFlowArn) > 0 {
		input.FlowArn = aws.String(_mediaconnectFlowArn)
	}
	if len(_mediaconnectVpcInterfaceName) > 0 {
		input.VpcInterfaceName = aws.String(_mediaconnectVpcInterfaceName)
	}

	if resp, err := client.RemoveFlowVpcInterface(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Restarts a router input. This operation can be used to recover from errors or
// refresh the input state.
func mediaconnect_RestartRouterInput(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.RestartRouterInputInput{
		// Arn: *string, // Required
	}

	if len(_mediaconnectArn) > 0 {
		input.Arn = aws.String(_mediaconnectArn)
	}

	if resp, err := client.RestartRouterInput(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Restarts a router output. This operation can be used to recover from errors or
// refresh the output state.
func mediaconnect_RestartRouterOutput(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.RestartRouterOutputInput{
		// Arn: *string, // Required
	}

	if len(_mediaconnectArn) > 0 {
		input.Arn = aws.String(_mediaconnectArn)
	}

	if resp, err := client.RestartRouterOutput(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Revokes an entitlement from a flow. Once an entitlement is revoked, the
// content becomes unavailable to the subscriber and the associated output is
// removed.
func mediaconnect_RevokeFlowEntitlement(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.RevokeFlowEntitlementInput{
		// EntitlementArn: *string, // Required
		// FlowArn: *string, // Required
	}

	if len(_mediaconnectEntitlementArn) > 0 {
		input.EntitlementArn = aws.String(_mediaconnectEntitlementArn)
	}
	if len(_mediaconnectFlowArn) > 0 {
		input.FlowArn = aws.String(_mediaconnectFlowArn)
	}

	if resp, err := client.RevokeFlowEntitlement(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a flow.
func mediaconnect_StartFlow(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.StartFlowInput{
		// FlowArn: *string, // Required
	}

	if len(_mediaconnectFlowArn) > 0 {
		input.FlowArn = aws.String(_mediaconnectFlowArn)
	}

	if resp, err := client.StartFlow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a router input in AWS Elemental MediaConnect.
func mediaconnect_StartRouterInput(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.StartRouterInputInput{
		// Arn: *string, // Required
	}

	if len(_mediaconnectArn) > 0 {
		input.Arn = aws.String(_mediaconnectArn)
	}

	if resp, err := client.StartRouterInput(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a router output in AWS Elemental MediaConnect.
func mediaconnect_StartRouterOutput(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.StartRouterOutputInput{
		// Arn: *string, // Required
	}

	if len(_mediaconnectArn) > 0 {
		input.Arn = aws.String(_mediaconnectArn)
	}

	if resp, err := client.StartRouterOutput(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a flow.
func mediaconnect_StopFlow(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.StopFlowInput{
		// FlowArn: *string, // Required
	}

	if len(_mediaconnectFlowArn) > 0 {
		input.FlowArn = aws.String(_mediaconnectFlowArn)
	}

	if resp, err := client.StopFlow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a router input in AWS Elemental MediaConnect.
func mediaconnect_StopRouterInput(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.StopRouterInputInput{
		// Arn: *string, // Required
	}

	if len(_mediaconnectArn) > 0 {
		input.Arn = aws.String(_mediaconnectArn)
	}

	if resp, err := client.StopRouterInput(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a router output in AWS Elemental MediaConnect.
func mediaconnect_StopRouterOutput(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.StopRouterOutputInput{
		// Arn: *string, // Required
	}

	if len(_mediaconnectArn) > 0 {
		input.Arn = aws.String(_mediaconnectArn)
	}

	if resp, err := client.StopRouterOutput(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds tags to a global resource in AWS Elemental MediaConnect. The API supports
// the following global resources: router inputs, router outputs and router network
// interfaces.
func mediaconnect_TagGlobalResource(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.TagGlobalResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_mediaconnectResourceArn) > 0 {
		input.ResourceArn = aws.String(_mediaconnectResourceArn)
	}
	if len(_mediaconnectTags) > 0 {
		if err := assignInputField(input, "Tags", _mediaconnectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.TagGlobalResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates the specified tags to a resource with the specified resourceArn in
// the current region. If existing tags on a resource are not specified in the
// request parameters, they are not changed. When a resource is deleted, the tags
// associated with that resource are deleted as well.
func mediaconnect_TagResource(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_mediaconnectResourceArn) > 0 {
		input.ResourceArn = aws.String(_mediaconnectResourceArn)
	}
	if len(_mediaconnectTags) > 0 {
		if err := assignInputField(input, "Tags", _mediaconnectTags); err != nil {
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

// Associates a router input with a router output in AWS Elemental MediaConnect.
func mediaconnect_TakeRouterInput(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.TakeRouterInputInput{
		// RouterOutputArn: *string, // Required
	}

	if len(_mediaconnectRouterOutputArn) > 0 {
		input.RouterOutputArn = aws.String(_mediaconnectRouterOutputArn)
	}
	if len(_mediaconnectRouterInputArn) > 0 {
		input.RouterInputArn = aws.String(_mediaconnectRouterInputArn)
	}

	if resp, err := client.TakeRouterInput(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes tags from a global resource in AWS Elemental MediaConnect. The API
// supports the following global resources: router inputs, router outputs and
// router network interfaces.
func mediaconnect_UntagGlobalResource(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.UntagGlobalResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_mediaconnectResourceArn) > 0 {
		input.ResourceArn = aws.String(_mediaconnectResourceArn)
	}
	if len(_mediaconnectTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _mediaconnectTagKeys...)
	}

	if resp, err := client.UntagGlobalResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes specified tags from a resource in the current region.
func mediaconnect_UntagResource(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_mediaconnectResourceArn) > 0 {
		input.ResourceArn = aws.String(_mediaconnectResourceArn)
	}
	if len(_mediaconnectTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _mediaconnectTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the bridge.
func mediaconnect_UpdateBridge(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.UpdateBridgeInput{
		// BridgeArn: *string, // Required
	}

	if len(_mediaconnectBridgeArn) > 0 {
		input.BridgeArn = aws.String(_mediaconnectBridgeArn)
	}
	if len(_mediaconnectEgressGatewayBridge) > 0 {
		if err := assignInputField(input, "EgressGatewayBridge", _mediaconnectEgressGatewayBridge); err != nil {
			log.Errorf("invalid --egress-gateway-bridge: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectIngressGatewayBridge) > 0 {
		if err := assignInputField(input, "IngressGatewayBridge", _mediaconnectIngressGatewayBridge); err != nil {
			log.Errorf("invalid --ingress-gateway-bridge: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectSourceFailoverConfig) > 0 {
		if err := assignInputField(input, "SourceFailoverConfig", _mediaconnectSourceFailoverConfig); err != nil {
			log.Errorf("invalid --source-failover-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateBridge(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing bridge output.
func mediaconnect_UpdateBridgeOutput(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.UpdateBridgeOutputInput{
		// BridgeArn: *string, // Required
		// OutputName: *string, // Required
	}

	if len(_mediaconnectBridgeArn) > 0 {
		input.BridgeArn = aws.String(_mediaconnectBridgeArn)
	}
	if len(_mediaconnectOutputName) > 0 {
		input.OutputName = aws.String(_mediaconnectOutputName)
	}
	if len(_mediaconnectNetworkOutput) > 0 {
		if err := assignInputField(input, "NetworkOutput", _mediaconnectNetworkOutput); err != nil {
			log.Errorf("invalid --network-output: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateBridgeOutput(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing bridge source.
func mediaconnect_UpdateBridgeSource(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.UpdateBridgeSourceInput{
		// BridgeArn: *string, // Required
		// SourceName: *string, // Required
	}

	if len(_mediaconnectBridgeArn) > 0 {
		input.BridgeArn = aws.String(_mediaconnectBridgeArn)
	}
	if len(_mediaconnectSourceName) > 0 {
		input.SourceName = aws.String(_mediaconnectSourceName)
	}
	if len(_mediaconnectFlowSource) > 0 {
		if err := assignInputField(input, "FlowSource", _mediaconnectFlowSource); err != nil {
			log.Errorf("invalid --flow-source: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectNetworkSource) > 0 {
		if err := assignInputField(input, "NetworkSource", _mediaconnectNetworkSource); err != nil {
			log.Errorf("invalid --network-source: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateBridgeSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the bridge state.
func mediaconnect_UpdateBridgeState(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.UpdateBridgeStateInput{
		// BridgeArn: *string, // Required
		// DesiredState: types.DesiredState, // Required
	}

	if len(_mediaconnectBridgeArn) > 0 {
		input.BridgeArn = aws.String(_mediaconnectBridgeArn)
	}
	if len(_mediaconnectDesiredState) > 0 {
		if err := assignInputField(input, "DesiredState", _mediaconnectDesiredState); err != nil {
			log.Errorf("invalid --desired-state: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateBridgeState(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing flow.
// Because UpdateFlowSources and UpdateFlow are separate operations, you can't
// change both the source type AND the flow size in a single request.
//
// - If you have a MEDIUM flow and you want to change the flow source to NDI®:
//
// - First, use the UpdateFlow operation to upgrade the flow size to LARGE .
//
// - After that, you can then use the UpdateFlowSource operation to configure the
// NDI source.
//
// - If you're switching from an NDI source to a transport stream (TS) source
// and want to downgrade the flow size:
//
// - First, use the UpdateFlowSource operation to change the flow source type.
//
// - After that, you can then use the UpdateFlow operation to downgrade the flow
// size to MEDIUM .
func mediaconnect_UpdateFlow(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.UpdateFlowInput{
		// FlowArn: *string, // Required
	}

	if len(_mediaconnectFlowArn) > 0 {
		input.FlowArn = aws.String(_mediaconnectFlowArn)
	}
	if len(_mediaconnectEncodingConfig) > 0 {
		if err := assignInputField(input, "EncodingConfig", _mediaconnectEncodingConfig); err != nil {
			log.Errorf("invalid --encoding-config: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectFlowSize) > 0 {
		if err := assignInputField(input, "FlowSize", _mediaconnectFlowSize); err != nil {
			log.Errorf("invalid --flow-size: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectMaintenance) > 0 {
		if err := assignInputField(input, "Maintenance", _mediaconnectMaintenance); err != nil {
			log.Errorf("invalid --maintenance: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectNdiConfig) > 0 {
		if err := assignInputField(input, "NdiConfig", _mediaconnectNdiConfig); err != nil {
			log.Errorf("invalid --ndi-config: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectSourceFailoverConfig) > 0 {
		if err := assignInputField(input, "SourceFailoverConfig", _mediaconnectSourceFailoverConfig); err != nil {
			log.Errorf("invalid --source-failover-config: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectSourceMonitoringConfig) > 0 {
		if err := assignInputField(input, "SourceMonitoringConfig", _mediaconnectSourceMonitoringConfig); err != nil {
			log.Errorf("invalid --source-monitoring-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateFlow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an entitlement. You can change an entitlement's description,
// subscribers, and encryption. If you change the subscribers, the service will
// remove the outputs that are are used by the subscribers that are removed.
func mediaconnect_UpdateFlowEntitlement(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.UpdateFlowEntitlementInput{
		// EntitlementArn: *string, // Required
		// FlowArn: *string, // Required
	}

	if len(_mediaconnectEntitlementArn) > 0 {
		input.EntitlementArn = aws.String(_mediaconnectEntitlementArn)
	}
	if len(_mediaconnectFlowArn) > 0 {
		input.FlowArn = aws.String(_mediaconnectFlowArn)
	}
	if len(_mediaconnectDescription) > 0 {
		input.Description = aws.String(_mediaconnectDescription)
	}
	if len(_mediaconnectEncryption) > 0 {
		if err := assignInputField(input, "Encryption", _mediaconnectEncryption); err != nil {
			log.Errorf("invalid --encryption: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectEntitlementStatus) > 0 {
		if err := assignInputField(input, "EntitlementStatus", _mediaconnectEntitlementStatus); err != nil {
			log.Errorf("invalid --entitlement-status: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectSubscribers) > 0 {
		input.Subscribers = append([]string(nil), _mediaconnectSubscribers...)
	}

	if resp, err := client.UpdateFlowEntitlement(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing media stream.
func mediaconnect_UpdateFlowMediaStream(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.UpdateFlowMediaStreamInput{
		// FlowArn: *string, // Required
		// MediaStreamName: *string, // Required
	}

	if len(_mediaconnectFlowArn) > 0 {
		input.FlowArn = aws.String(_mediaconnectFlowArn)
	}
	if len(_mediaconnectMediaStreamName) > 0 {
		input.MediaStreamName = aws.String(_mediaconnectMediaStreamName)
	}
	if len(_mediaconnectAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _mediaconnectAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectClockRate) > 0 {
		if err := assignInputField(input, "ClockRate", _mediaconnectClockRate); err != nil {
			log.Errorf("invalid --clock-rate: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectDescription) > 0 {
		input.Description = aws.String(_mediaconnectDescription)
	}
	if len(_mediaconnectMediaStreamType) > 0 {
		if err := assignInputField(input, "MediaStreamType", _mediaconnectMediaStreamType); err != nil {
			log.Errorf("invalid --media-stream-type: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectVideoFormat) > 0 {
		input.VideoFormat = aws.String(_mediaconnectVideoFormat)
	}

	if resp, err := client.UpdateFlowMediaStream(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing flow output.
func mediaconnect_UpdateFlowOutput(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.UpdateFlowOutputInput{
		// FlowArn: *string, // Required
		// OutputArn: *string, // Required
	}

	if len(_mediaconnectFlowArn) > 0 {
		input.FlowArn = aws.String(_mediaconnectFlowArn)
	}
	if len(_mediaconnectOutputArn) > 0 {
		input.OutputArn = aws.String(_mediaconnectOutputArn)
	}
	if len(_mediaconnectCidrAllowList) > 0 {
		input.CidrAllowList = append([]string(nil), _mediaconnectCidrAllowList...)
	}
	if len(_mediaconnectDescription) > 0 {
		input.Description = aws.String(_mediaconnectDescription)
	}
	if len(_mediaconnectDestination) > 0 {
		input.Destination = aws.String(_mediaconnectDestination)
	}
	if len(_mediaconnectEncryption) > 0 {
		if err := assignInputField(input, "Encryption", _mediaconnectEncryption); err != nil {
			log.Errorf("invalid --encryption: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectMaxLatency) > 0 {
		if err := assignInputField(input, "MaxLatency", _mediaconnectMaxLatency); err != nil {
			log.Errorf("invalid --max-latency: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectMediaStreamOutputConfigurations) > 0 {
		if err := assignInputField(input, "MediaStreamOutputConfigurations", _mediaconnectMediaStreamOutputConfigurations); err != nil {
			log.Errorf("invalid --media-stream-output-configurations: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectMinLatency) > 0 {
		if err := assignInputField(input, "MinLatency", _mediaconnectMinLatency); err != nil {
			log.Errorf("invalid --min-latency: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectNdiProgramName) > 0 {
		input.NdiProgramName = aws.String(_mediaconnectNdiProgramName)
	}
	if len(_mediaconnectNdiSpeedHqQuality) > 0 {
		if err := assignInputField(input, "NdiSpeedHqQuality", _mediaconnectNdiSpeedHqQuality); err != nil {
			log.Errorf("invalid --ndi-speed-hq-quality: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectOutputStatus) > 0 {
		if err := assignInputField(input, "OutputStatus", _mediaconnectOutputStatus); err != nil {
			log.Errorf("invalid --output-status: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectPort) > 0 {
		if err := assignInputField(input, "Port", _mediaconnectPort); err != nil {
			log.Errorf("invalid --port: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectProtocol) > 0 {
		if err := assignInputField(input, "Protocol", _mediaconnectProtocol); err != nil {
			log.Errorf("invalid --protocol: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectRemoteId) > 0 {
		input.RemoteId = aws.String(_mediaconnectRemoteId)
	}
	if len(_mediaconnectRouterIntegrationState) > 0 {
		if err := assignInputField(input, "RouterIntegrationState", _mediaconnectRouterIntegrationState); err != nil {
			log.Errorf("invalid --router-integration-state: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectRouterIntegrationTransitEncryption) > 0 {
		if err := assignInputField(input, "RouterIntegrationTransitEncryption", _mediaconnectRouterIntegrationTransitEncryption); err != nil {
			log.Errorf("invalid --router-integration-transit-encryption: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectSenderControlPort) > 0 {
		if err := assignInputField(input, "SenderControlPort", _mediaconnectSenderControlPort); err != nil {
			log.Errorf("invalid --sender-control-port: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectSenderIpAddress) > 0 {
		input.SenderIpAddress = aws.String(_mediaconnectSenderIpAddress)
	}
	if len(_mediaconnectSmoothingLatency) > 0 {
		if err := assignInputField(input, "SmoothingLatency", _mediaconnectSmoothingLatency); err != nil {
			log.Errorf("invalid --smoothing-latency: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectStreamId) > 0 {
		input.StreamId = aws.String(_mediaconnectStreamId)
	}
	if len(_mediaconnectVpcInterfaceAttachment) > 0 {
		if err := assignInputField(input, "VpcInterfaceAttachment", _mediaconnectVpcInterfaceAttachment); err != nil {
			log.Errorf("invalid --vpc-interface-attachment: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateFlowOutput(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the source of a flow.
// Because UpdateFlowSources and UpdateFlow are separate operations, you can't
// change both the source type AND the flow size in a single request.
//
// - If you have a MEDIUM flow and you want to change the flow source to NDI®:
//
// - First, use the UpdateFlow operation to upgrade the flow size to LARGE .
//
// - After that, you can then use the UpdateFlowSource operation to configure the
// NDI source.
//
// - If you're switching from an NDI source to a transport stream (TS) source
// and want to downgrade the flow size:
//
// - First, use the UpdateFlowSource operation to change the flow source type.
//
// - After that, you can then use the UpdateFlow operation to downgrade the flow
// size to MEDIUM .
func mediaconnect_UpdateFlowSource(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.UpdateFlowSourceInput{
		// FlowArn: *string, // Required
		// SourceArn: *string, // Required
	}

	if len(_mediaconnectFlowArn) > 0 {
		input.FlowArn = aws.String(_mediaconnectFlowArn)
	}
	if len(_mediaconnectSourceArn) > 0 {
		input.SourceArn = aws.String(_mediaconnectSourceArn)
	}
	if len(_mediaconnectDecryption) > 0 {
		if err := assignInputField(input, "Decryption", _mediaconnectDecryption); err != nil {
			log.Errorf("invalid --decryption: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectDescription) > 0 {
		input.Description = aws.String(_mediaconnectDescription)
	}
	if len(_mediaconnectEntitlementArn) > 0 {
		input.EntitlementArn = aws.String(_mediaconnectEntitlementArn)
	}
	if len(_mediaconnectGatewayBridgeSource) > 0 {
		if err := assignInputField(input, "GatewayBridgeSource", _mediaconnectGatewayBridgeSource); err != nil {
			log.Errorf("invalid --gateway-bridge-source: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectIngestPort) > 0 {
		if err := assignInputField(input, "IngestPort", _mediaconnectIngestPort); err != nil {
			log.Errorf("invalid --ingest-port: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectMaxBitrate) > 0 {
		if err := assignInputField(input, "MaxBitrate", _mediaconnectMaxBitrate); err != nil {
			log.Errorf("invalid --max-bitrate: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectMaxLatency) > 0 {
		if err := assignInputField(input, "MaxLatency", _mediaconnectMaxLatency); err != nil {
			log.Errorf("invalid --max-latency: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectMaxSyncBuffer) > 0 {
		if err := assignInputField(input, "MaxSyncBuffer", _mediaconnectMaxSyncBuffer); err != nil {
			log.Errorf("invalid --max-sync-buffer: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectMediaStreamSourceConfigurations) > 0 {
		if err := assignInputField(input, "MediaStreamSourceConfigurations", _mediaconnectMediaStreamSourceConfigurations); err != nil {
			log.Errorf("invalid --media-stream-source-configurations: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectMinLatency) > 0 {
		if err := assignInputField(input, "MinLatency", _mediaconnectMinLatency); err != nil {
			log.Errorf("invalid --min-latency: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectNdiSourceSettings) > 0 {
		if err := assignInputField(input, "NdiSourceSettings", _mediaconnectNdiSourceSettings); err != nil {
			log.Errorf("invalid --ndi-source-settings: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectProtocol) > 0 {
		if err := assignInputField(input, "Protocol", _mediaconnectProtocol); err != nil {
			log.Errorf("invalid --protocol: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectRouterIntegrationState) > 0 {
		if err := assignInputField(input, "RouterIntegrationState", _mediaconnectRouterIntegrationState); err != nil {
			log.Errorf("invalid --router-integration-state: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectRouterIntegrationTransitDecryption) > 0 {
		if err := assignInputField(input, "RouterIntegrationTransitDecryption", _mediaconnectRouterIntegrationTransitDecryption); err != nil {
			log.Errorf("invalid --router-integration-transit-decryption: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectSenderControlPort) > 0 {
		if err := assignInputField(input, "SenderControlPort", _mediaconnectSenderControlPort); err != nil {
			log.Errorf("invalid --sender-control-port: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectSenderIpAddress) > 0 {
		input.SenderIpAddress = aws.String(_mediaconnectSenderIpAddress)
	}
	if len(_mediaconnectSourceListenerAddress) > 0 {
		input.SourceListenerAddress = aws.String(_mediaconnectSourceListenerAddress)
	}
	if len(_mediaconnectSourceListenerPort) > 0 {
		if err := assignInputField(input, "SourceListenerPort", _mediaconnectSourceListenerPort); err != nil {
			log.Errorf("invalid --source-listener-port: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectStreamId) > 0 {
		input.StreamId = aws.String(_mediaconnectStreamId)
	}
	if len(_mediaconnectVpcInterfaceName) > 0 {
		input.VpcInterfaceName = aws.String(_mediaconnectVpcInterfaceName)
	}
	if len(_mediaconnectWhitelistCidr) > 0 {
		input.WhitelistCidr = aws.String(_mediaconnectWhitelistCidr)
	}

	if resp, err := client.UpdateFlowSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing gateway instance.
func mediaconnect_UpdateGatewayInstance(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.UpdateGatewayInstanceInput{
		// GatewayInstanceArn: *string, // Required
	}

	if len(_mediaconnectGatewayInstanceArn) > 0 {
		input.GatewayInstanceArn = aws.String(_mediaconnectGatewayInstanceArn)
	}
	if len(_mediaconnectBridgePlacement) > 0 {
		if err := assignInputField(input, "BridgePlacement", _mediaconnectBridgePlacement); err != nil {
			log.Errorf("invalid --bridge-placement: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateGatewayInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration of an existing router input in AWS Elemental
// MediaConnect.
func mediaconnect_UpdateRouterInput(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.UpdateRouterInputInput{
		// Arn: *string, // Required
	}

	if len(_mediaconnectArn) > 0 {
		input.Arn = aws.String(_mediaconnectArn)
	}
	if len(_mediaconnectConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _mediaconnectConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectMaintenanceConfiguration) > 0 {
		if err := assignInputField(input, "MaintenanceConfiguration", _mediaconnectMaintenanceConfiguration); err != nil {
			log.Errorf("invalid --maintenance-configuration: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectMaximumBitrate) > 0 {
		if err := assignInputField(input, "MaximumBitrate", _mediaconnectMaximumBitrate); err != nil {
			log.Errorf("invalid --maximum-bitrate: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectName) > 0 {
		input.Name = aws.String(_mediaconnectName)
	}
	if len(_mediaconnectRoutingScope) > 0 {
		if err := assignInputField(input, "RoutingScope", _mediaconnectRoutingScope); err != nil {
			log.Errorf("invalid --routing-scope: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectTier) > 0 {
		if err := assignInputField(input, "Tier", _mediaconnectTier); err != nil {
			log.Errorf("invalid --tier: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectTransitEncryption) > 0 {
		if err := assignInputField(input, "TransitEncryption", _mediaconnectTransitEncryption); err != nil {
			log.Errorf("invalid --transit-encryption: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateRouterInput(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration of an existing router network interface in AWS
// Elemental MediaConnect.
func mediaconnect_UpdateRouterNetworkInterface(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.UpdateRouterNetworkInterfaceInput{
		// Arn: *string, // Required
	}

	if len(_mediaconnectArn) > 0 {
		input.Arn = aws.String(_mediaconnectArn)
	}
	if len(_mediaconnectConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _mediaconnectConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectName) > 0 {
		input.Name = aws.String(_mediaconnectName)
	}

	if resp, err := client.UpdateRouterNetworkInterface(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration of an existing router output in AWS Elemental
// MediaConnect.
func mediaconnect_UpdateRouterOutput(cfg aws.Config, client *mediaconnect.Client) {
	input := &mediaconnect.UpdateRouterOutputInput{
		// Arn: *string, // Required
	}

	if len(_mediaconnectArn) > 0 {
		input.Arn = aws.String(_mediaconnectArn)
	}
	if len(_mediaconnectConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _mediaconnectConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectMaintenanceConfiguration) > 0 {
		if err := assignInputField(input, "MaintenanceConfiguration", _mediaconnectMaintenanceConfiguration); err != nil {
			log.Errorf("invalid --maintenance-configuration: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectMaximumBitrate) > 0 {
		if err := assignInputField(input, "MaximumBitrate", _mediaconnectMaximumBitrate); err != nil {
			log.Errorf("invalid --maximum-bitrate: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectName) > 0 {
		input.Name = aws.String(_mediaconnectName)
	}
	if len(_mediaconnectRoutingScope) > 0 {
		if err := assignInputField(input, "RoutingScope", _mediaconnectRoutingScope); err != nil {
			log.Errorf("invalid --routing-scope: %s", err.Error())
			return
		}
	}
	if len(_mediaconnectTier) > 0 {
		if err := assignInputField(input, "Tier", _mediaconnectTier); err != nil {
			log.Errorf("invalid --tier: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateRouterOutput(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_mediaconnectCmd)
	_mediaconnectCmd.Flags().SortFlags = false

	_mediaconnectCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_mediaconnectCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_mediaconnectCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectArn, "arn", "", "", "ARN")
	_mediaconnectCmd.Flags().StringSliceVarP(&_mediaconnectArns, "arns", "", nil, "Arns")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectAttributes, "attributes", "", "", "Attributes")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectAvailabilityZone, "availability-zone", "", "", "Availability Zone")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectBridgeArn, "bridge-arn", "", "", "Bridge ARN")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectBridgePlacement, "bridge-placement", "", "", "Bridge Placement")
	_mediaconnectCmd.Flags().StringSliceVarP(&_mediaconnectCidrAllowList, "cidr-allow-list", "", nil, "CIDR Allow List")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectClientToken, "client-token", "", "", "Client Token")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectClockRate, "clock-rate", "", "", "Clock Rate")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectConfiguration, "configuration", "", "", "Configuration")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectDecryption, "decryption", "", "", "Decryption")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectDescription, "description", "", "", "Description")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectDesiredState, "desired-state", "", "", "Desired State")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectDestination, "destination", "", "", "Destination")
	_mediaconnectCmd.Flags().StringSliceVarP(&_mediaconnectEgressCidrBlocks, "egress-cidr-blocks", "", nil, "Egress CIDR Blocks")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectEgressGatewayBridge, "egress-gateway-bridge", "", "", "Egress Gateway Bridge")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectEncodingConfig, "encoding-config", "", "", "Encoding Config")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectEncryption, "encryption", "", "", "Encryption")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectEntitlementArn, "entitlement-arn", "", "", "Entitlement ARN")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectEntitlementStatus, "entitlement-status", "", "", "Entitlement Status")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectEntitlements, "entitlements", "", "", "Entitlements")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectFilterArn, "filter-arn", "", "", "Filter ARN")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectFilters, "filters", "", "", "Filters")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectFlowArn, "flow-arn", "", "", "Flow ARN")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectFlowSize, "flow-size", "", "", "Flow Size")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectFlowSource, "flow-source", "", "", "Flow Source")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectFlowTags, "flow-tags", "", "", "Flow Tags")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectForce, "force", "", "", "Force")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectGatewayArn, "gateway-arn", "", "", "Gateway ARN")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectGatewayBridgeSource, "gateway-bridge-source", "", "", "Gateway Bridge Source")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectGatewayInstanceArn, "gateway-instance-arn", "", "", "Gateway Instance ARN")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectIngestPort, "ingest-port", "", "", "Ingest Port")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectIngressGatewayBridge, "ingress-gateway-bridge", "", "", "Ingress Gateway Bridge")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectMaintenance, "maintenance", "", "", "Maintenance")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectMaintenanceConfiguration, "maintenance-configuration", "", "", "Maintenance Configuration")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectMaxBitrate, "max-bitrate", "", "", "Max Bitrate")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectMaxLatency, "max-latency", "", "", "Max Latency")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectMaxResults, "max-results", "", "", "Max Results")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectMaxSyncBuffer, "max-sync-buffer", "", "", "Max Sync Buffer")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectMaximumBitrate, "maximum-bitrate", "", "", "Maximum Bitrate")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectMediaStreamName, "media-stream-name", "", "", "Media Stream Name")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectMediaStreamOutputConfigurations, "media-stream-output-configurations", "", "", "Media Stream Output Configurations")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectMediaStreamSourceConfigurations, "media-stream-source-configurations", "", "", "Media Stream Source Configurations")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectMediaStreamType, "media-stream-type", "", "", "Media Stream Type")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectMediaStreams, "media-streams", "", "", "Media Streams")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectMinLatency, "min-latency", "", "", "Min Latency")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectName, "name", "", "", "Name")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectNdiConfig, "ndi-config", "", "", "Ndi Config")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectNdiProgramName, "ndi-program-name", "", "", "Ndi Program Name")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectNdiSourceSettings, "ndi-source-settings", "", "", "Ndi Source Settings")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectNdiSpeedHqQuality, "ndi-speed-hq-quality", "", "", "Ndi Speed Hq Quality")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectNetworkOutput, "network-output", "", "", "Network Output")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectNetworkSource, "network-source", "", "", "Network Source")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectNetworks, "networks", "", "", "Networks")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectNextToken, "next-token", "", "", "Next Token")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectOfferingArn, "offering-arn", "", "", "Offering ARN")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectOutputArn, "output-arn", "", "", "Output ARN")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectOutputName, "output-name", "", "", "Output Name")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectOutputStatus, "output-status", "", "", "Output Status")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectOutputs, "outputs", "", "", "Outputs")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectPlacementArn, "placement-arn", "", "", "Placement ARN")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectPort, "port", "", "", "Port")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectProtocol, "protocol", "", "", "Protocol")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectRegionName, "region-name", "", "", "Region Name")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectRemoteId, "remote-id", "", "", "Remote ID")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectReservationArn, "reservation-arn", "", "", "Reservation ARN")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectReservationName, "reservation-name", "", "", "Reservation Name")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectResourceArn, "resource-arn", "", "", "Resource ARN")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectRouterInputArn, "router-input-arn", "", "", "Router Input ARN")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectRouterIntegrationState, "router-integration-state", "", "", "Router Integration State")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectRouterIntegrationTransitDecryption, "router-integration-transit-decryption", "", "", "Router Integration Transit Decryption")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectRouterIntegrationTransitEncryption, "router-integration-transit-encryption", "", "", "Router Integration Transit Encryption")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectRouterOutputArn, "router-output-arn", "", "", "Router Output ARN")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectRoutingScope, "routing-scope", "", "", "Routing Scope")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectSenderControlPort, "sender-control-port", "", "", "Sender Control Port")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectSenderIpAddress, "sender-ip-address", "", "", "Sender IP Address")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectSmoothingLatency, "smoothing-latency", "", "", "Smoothing Latency")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectSource, "source", "", "", "Source")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectSourceArn, "source-arn", "", "", "Source ARN")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectSourceFailoverConfig, "source-failover-config", "", "", "Source Failover Config")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectSourceListenerAddress, "source-listener-address", "", "", "Source Listener Address")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectSourceListenerPort, "source-listener-port", "", "", "Source Listener Port")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectSourceMonitoringConfig, "source-monitoring-config", "", "", "Source Monitoring Config")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectSourceName, "source-name", "", "", "Source Name")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectSources, "sources", "", "", "Sources")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectStart, "start", "", "", "Start")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectStreamId, "stream-id", "", "", "Stream ID")
	_mediaconnectCmd.Flags().StringSliceVarP(&_mediaconnectSubscribers, "subscribers", "", nil, "Subscribers")
	_mediaconnectCmd.Flags().StringSliceVarP(&_mediaconnectTagKeys, "tag-keys", "", nil, "Tag Keys")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectTags, "tags", "", "", "Tags")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectTier, "tier", "", "", "Tier")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectTransitEncryption, "transit-encryption", "", "", "Transit Encryption")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectVideoFormat, "video-format", "", "", "Video Format")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectVpcInterfaceAttachment, "vpc-interface-attachment", "", "", "VPC Interface Attachment")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectVpcInterfaceName, "vpc-interface-name", "", "", "VPC Interface Name")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectVpcInterfaces, "vpc-interfaces", "", "", "VPC Interfaces")
	_mediaconnectCmd.Flags().StringVarP(&_mediaconnectWhitelistCidr, "whitelist-cidr", "", "", "Whitelist CIDR")

	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectAddBridgeOutputs, "add-bridge-outputs", "", false, "Add Bridge Outputs")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectAddBridgeSources, "add-bridge-sources", "", false, "Add Bridge Sources")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectAddFlowMediaStreams, "add-flow-media-streams", "", false, "Add Flow Media Streams")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectAddFlowOutputs, "add-flow-outputs", "", false, "Add Flow Outputs")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectAddFlowSources, "add-flow-sources", "", false, "Add Flow Sources")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectAddFlowVpcInterfaces, "add-flow-vpc-interfaces", "", false, "Add Flow VPC Interfaces")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectBatchGetRouterInput, "batch-get-router-input", "", false, "Batch Get Router Input")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectBatchGetRouterNetworkInterface, "batch-get-router-network-interface", "", false, "Batch Get Router Network Interface")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectBatchGetRouterOutput, "batch-get-router-output", "", false, "Batch Get Router Output")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectCreateBridge, "create-bridge", "", false, "Create Bridge")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectCreateFlow, "create-flow", "", false, "Create Flow")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectCreateGateway, "create-gateway", "", false, "Create Gateway")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectCreateRouterInput, "create-router-input", "", false, "Create Router Input")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectCreateRouterNetworkInterface, "create-router-network-interface", "", false, "Create Router Network Interface")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectCreateRouterOutput, "create-router-output", "", false, "Create Router Output")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectDeleteBridge, "delete-bridge", "", false, "Delete Bridge")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectDeleteFlow, "delete-flow", "", false, "Delete Flow")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectDeleteGateway, "delete-gateway", "", false, "Delete Gateway")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectDeleteRouterInput, "delete-router-input", "", false, "Delete Router Input")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectDeleteRouterNetworkInterface, "delete-router-network-interface", "", false, "Delete Router Network Interface")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectDeleteRouterOutput, "delete-router-output", "", false, "Delete Router Output")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectDeregisterGatewayInstance, "deregister-gateway-instance", "", false, "Deregister Gateway Instance")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectDescribeBridge, "describe-bridge", "", false, "Describe Bridge")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectDescribeFlow, "describe-flow", "", false, "Describe Flow")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectDescribeFlowSourceMetadata, "describe-flow-source-metadata", "", false, "Describe Flow Source Metadata")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectDescribeFlowSourceThumbnail, "describe-flow-source-thumbnail", "", false, "Describe Flow Source Thumbnail")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectDescribeGateway, "describe-gateway", "", false, "Describe Gateway")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectDescribeGatewayInstance, "describe-gateway-instance", "", false, "Describe Gateway Instance")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectDescribeOffering, "describe-offering", "", false, "Describe Offering")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectDescribeReservation, "describe-reservation", "", false, "Describe Reservation")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectGetRouterInput, "get-router-input", "", false, "Get Router Input")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectGetRouterInputSourceMetadata, "get-router-input-source-metadata", "", false, "Get Router Input Source Metadata")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectGetRouterInputThumbnail, "get-router-input-thumbnail", "", false, "Get Router Input Thumbnail")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectGetRouterNetworkInterface, "get-router-network-interface", "", false, "Get Router Network Interface")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectGetRouterOutput, "get-router-output", "", false, "Get Router Output")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectGrantFlowEntitlements, "grant-flow-entitlements", "", false, "Grant Flow Entitlements")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectListBridges, "list-bridges", "", false, "List Bridges")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectListEntitlements, "list-entitlements", "", false, "List Entitlements")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectListFlows, "list-flows", "", false, "List Flows")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectListGatewayInstances, "list-gateway-instances", "", false, "List Gateway Instances")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectListGateways, "list-gateways", "", false, "List Gateways")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectListOfferings, "list-offerings", "", false, "List Offerings")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectListReservations, "list-reservations", "", false, "List Reservations")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectListRouterInputs, "list-router-inputs", "", false, "List Router Inputs")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectListRouterNetworkInterfaces, "list-router-network-interfaces", "", false, "List Router Network Interfaces")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectListRouterOutputs, "list-router-outputs", "", false, "List Router Outputs")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectListTagsForGlobalResource, "list-tags-for-global-resource", "", false, "List Tags For Global Resource")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectPurchaseOffering, "purchase-offering", "", false, "Purchase Offering")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectRemoveBridgeOutput, "remove-bridge-output", "", false, "Remove Bridge Output")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectRemoveBridgeSource, "remove-bridge-source", "", false, "Remove Bridge Source")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectRemoveFlowMediaStream, "remove-flow-media-stream", "", false, "Remove Flow Media Stream")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectRemoveFlowOutput, "remove-flow-output", "", false, "Remove Flow Output")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectRemoveFlowSource, "remove-flow-source", "", false, "Remove Flow Source")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectRemoveFlowVpcInterface, "remove-flow-vpc-interface", "", false, "Remove Flow VPC Interface")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectRestartRouterInput, "restart-router-input", "", false, "Restart Router Input")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectRestartRouterOutput, "restart-router-output", "", false, "Restart Router Output")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectRevokeFlowEntitlement, "revoke-flow-entitlement", "", false, "Revoke Flow Entitlement")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectStartFlow, "start-flow", "", false, "Start Flow")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectStartRouterInput, "start-router-input", "", false, "Start Router Input")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectStartRouterOutput, "start-router-output", "", false, "Start Router Output")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectStopFlow, "stop-flow", "", false, "Stop Flow")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectStopRouterInput, "stop-router-input", "", false, "Stop Router Input")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectStopRouterOutput, "stop-router-output", "", false, "Stop Router Output")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectTagGlobalResource, "tag-global-resource", "", false, "Tag Global Resource")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectTagResource, "tag-resource", "", false, "Tag Resource")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectTakeRouterInput, "take-router-input", "", false, "Take Router Input")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectUntagGlobalResource, "untag-global-resource", "", false, "Untag Global Resource")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectUntagResource, "untag-resource", "", false, "Untag Resource")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectUpdateBridge, "update-bridge", "", false, "Update Bridge")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectUpdateBridgeOutput, "update-bridge-output", "", false, "Update Bridge Output")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectUpdateBridgeSource, "update-bridge-source", "", false, "Update Bridge Source")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectUpdateBridgeState, "update-bridge-state", "", false, "Update Bridge State")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectUpdateFlow, "update-flow", "", false, "Update Flow")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectUpdateFlowEntitlement, "update-flow-entitlement", "", false, "Update Flow Entitlement")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectUpdateFlowMediaStream, "update-flow-media-stream", "", false, "Update Flow Media Stream")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectUpdateFlowOutput, "update-flow-output", "", false, "Update Flow Output")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectUpdateFlowSource, "update-flow-source", "", false, "Update Flow Source")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectUpdateGatewayInstance, "update-gateway-instance", "", false, "Update Gateway Instance")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectUpdateRouterInput, "update-router-input", "", false, "Update Router Input")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectUpdateRouterNetworkInterface, "update-router-network-interface", "", false, "Update Router Network Interface")
	_mediaconnectCmd.Flags().BoolVarP(&_mediaconnectUpdateRouterOutput, "update-router-output", "", false, "Update Router Output")

}
