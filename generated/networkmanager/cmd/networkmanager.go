package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/networkmanager"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// networkmanagerCmd represents the networkmanager command
var _networkmanagerCmd = &cobra.Command{
	Use:   "networkmanager",
	Short: "AWS networkmanager CLI",
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
		client := networkmanager.NewFromConfig(cfg)
		if _networkmanagerAcceptAttachment {
			networkmanager_AcceptAttachment(cfg, client)
			return
		}
		if _networkmanagerAssociateConnectPeer {
			networkmanager_AssociateConnectPeer(cfg, client)
			return
		}
		if _networkmanagerAssociateCustomerGateway {
			networkmanager_AssociateCustomerGateway(cfg, client)
			return
		}
		if _networkmanagerAssociateLink {
			networkmanager_AssociateLink(cfg, client)
			return
		}
		if _networkmanagerAssociateTransitGatewayConnectPeer {
			networkmanager_AssociateTransitGatewayConnectPeer(cfg, client)
			return
		}
		if _networkmanagerCreateConnectAttachment {
			networkmanager_CreateConnectAttachment(cfg, client)
			return
		}
		if _networkmanagerCreateConnectPeer {
			networkmanager_CreateConnectPeer(cfg, client)
			return
		}
		if _networkmanagerCreateConnection {
			networkmanager_CreateConnection(cfg, client)
			return
		}
		if _networkmanagerCreateCoreNetwork {
			networkmanager_CreateCoreNetwork(cfg, client)
			return
		}
		if _networkmanagerCreateCoreNetworkPrefixListAssociation {
			networkmanager_CreateCoreNetworkPrefixListAssociation(cfg, client)
			return
		}
		if _networkmanagerCreateDevice {
			networkmanager_CreateDevice(cfg, client)
			return
		}
		if _networkmanagerCreateDirectConnectGatewayAttachment {
			networkmanager_CreateDirectConnectGatewayAttachment(cfg, client)
			return
		}
		if _networkmanagerCreateGlobalNetwork {
			networkmanager_CreateGlobalNetwork(cfg, client)
			return
		}
		if _networkmanagerCreateLink {
			networkmanager_CreateLink(cfg, client)
			return
		}
		if _networkmanagerCreateSite {
			networkmanager_CreateSite(cfg, client)
			return
		}
		if _networkmanagerCreateSiteToSiteVpnAttachment {
			networkmanager_CreateSiteToSiteVpnAttachment(cfg, client)
			return
		}
		if _networkmanagerCreateTransitGatewayPeering {
			networkmanager_CreateTransitGatewayPeering(cfg, client)
			return
		}
		if _networkmanagerCreateTransitGatewayRouteTableAttachment {
			networkmanager_CreateTransitGatewayRouteTableAttachment(cfg, client)
			return
		}
		if _networkmanagerCreateVpcAttachment {
			networkmanager_CreateVpcAttachment(cfg, client)
			return
		}
		if _networkmanagerDeleteAttachment {
			networkmanager_DeleteAttachment(cfg, client)
			return
		}
		if _networkmanagerDeleteConnectPeer {
			networkmanager_DeleteConnectPeer(cfg, client)
			return
		}
		if _networkmanagerDeleteConnection {
			networkmanager_DeleteConnection(cfg, client)
			return
		}
		if _networkmanagerDeleteCoreNetwork {
			networkmanager_DeleteCoreNetwork(cfg, client)
			return
		}
		if _networkmanagerDeleteCoreNetworkPolicyVersion {
			networkmanager_DeleteCoreNetworkPolicyVersion(cfg, client)
			return
		}
		if _networkmanagerDeleteCoreNetworkPrefixListAssociation {
			networkmanager_DeleteCoreNetworkPrefixListAssociation(cfg, client)
			return
		}
		if _networkmanagerDeleteDevice {
			networkmanager_DeleteDevice(cfg, client)
			return
		}
		if _networkmanagerDeleteGlobalNetwork {
			networkmanager_DeleteGlobalNetwork(cfg, client)
			return
		}
		if _networkmanagerDeleteLink {
			networkmanager_DeleteLink(cfg, client)
			return
		}
		if _networkmanagerDeletePeering {
			networkmanager_DeletePeering(cfg, client)
			return
		}
		if _networkmanagerDeleteResourcePolicy {
			networkmanager_DeleteResourcePolicy(cfg, client)
			return
		}
		if _networkmanagerDeleteSite {
			networkmanager_DeleteSite(cfg, client)
			return
		}
		if _networkmanagerDeregisterTransitGateway {
			networkmanager_DeregisterTransitGateway(cfg, client)
			return
		}
		if _networkmanagerDescribeGlobalNetworks {
			networkmanager_DescribeGlobalNetworks(cfg, client)
			return
		}
		if _networkmanagerDisassociateConnectPeer {
			networkmanager_DisassociateConnectPeer(cfg, client)
			return
		}
		if _networkmanagerDisassociateCustomerGateway {
			networkmanager_DisassociateCustomerGateway(cfg, client)
			return
		}
		if _networkmanagerDisassociateLink {
			networkmanager_DisassociateLink(cfg, client)
			return
		}
		if _networkmanagerDisassociateTransitGatewayConnectPeer {
			networkmanager_DisassociateTransitGatewayConnectPeer(cfg, client)
			return
		}
		if _networkmanagerExecuteCoreNetworkChangeSet {
			networkmanager_ExecuteCoreNetworkChangeSet(cfg, client)
			return
		}
		if _networkmanagerGetConnectAttachment {
			networkmanager_GetConnectAttachment(cfg, client)
			return
		}
		if _networkmanagerGetConnectPeer {
			networkmanager_GetConnectPeer(cfg, client)
			return
		}
		if _networkmanagerGetConnectPeerAssociations {
			networkmanager_GetConnectPeerAssociations(cfg, client)
			return
		}
		if _networkmanagerGetConnections {
			networkmanager_GetConnections(cfg, client)
			return
		}
		if _networkmanagerGetCoreNetwork {
			networkmanager_GetCoreNetwork(cfg, client)
			return
		}
		if _networkmanagerGetCoreNetworkChangeEvents {
			networkmanager_GetCoreNetworkChangeEvents(cfg, client)
			return
		}
		if _networkmanagerGetCoreNetworkChangeSet {
			networkmanager_GetCoreNetworkChangeSet(cfg, client)
			return
		}
		if _networkmanagerGetCoreNetworkPolicy {
			networkmanager_GetCoreNetworkPolicy(cfg, client)
			return
		}
		if _networkmanagerGetCustomerGatewayAssociations {
			networkmanager_GetCustomerGatewayAssociations(cfg, client)
			return
		}
		if _networkmanagerGetDevices {
			networkmanager_GetDevices(cfg, client)
			return
		}
		if _networkmanagerGetDirectConnectGatewayAttachment {
			networkmanager_GetDirectConnectGatewayAttachment(cfg, client)
			return
		}
		if _networkmanagerGetLinkAssociations {
			networkmanager_GetLinkAssociations(cfg, client)
			return
		}
		if _networkmanagerGetLinks {
			networkmanager_GetLinks(cfg, client)
			return
		}
		if _networkmanagerGetNetworkResourceCounts {
			networkmanager_GetNetworkResourceCounts(cfg, client)
			return
		}
		if _networkmanagerGetNetworkResourceRelationships {
			networkmanager_GetNetworkResourceRelationships(cfg, client)
			return
		}
		if _networkmanagerGetNetworkResources {
			networkmanager_GetNetworkResources(cfg, client)
			return
		}
		if _networkmanagerGetNetworkRoutes {
			networkmanager_GetNetworkRoutes(cfg, client)
			return
		}
		if _networkmanagerGetNetworkTelemetry {
			networkmanager_GetNetworkTelemetry(cfg, client)
			return
		}
		if _networkmanagerGetResourcePolicy {
			networkmanager_GetResourcePolicy(cfg, client)
			return
		}
		if _networkmanagerGetRouteAnalysis {
			networkmanager_GetRouteAnalysis(cfg, client)
			return
		}
		if _networkmanagerGetSiteToSiteVpnAttachment {
			networkmanager_GetSiteToSiteVpnAttachment(cfg, client)
			return
		}
		if _networkmanagerGetSites {
			networkmanager_GetSites(cfg, client)
			return
		}
		if _networkmanagerGetTransitGatewayConnectPeerAssociations {
			networkmanager_GetTransitGatewayConnectPeerAssociations(cfg, client)
			return
		}
		if _networkmanagerGetTransitGatewayPeering {
			networkmanager_GetTransitGatewayPeering(cfg, client)
			return
		}
		if _networkmanagerGetTransitGatewayRegistrations {
			networkmanager_GetTransitGatewayRegistrations(cfg, client)
			return
		}
		if _networkmanagerGetTransitGatewayRouteTableAttachment {
			networkmanager_GetTransitGatewayRouteTableAttachment(cfg, client)
			return
		}
		if _networkmanagerGetVpcAttachment {
			networkmanager_GetVpcAttachment(cfg, client)
			return
		}
		if _networkmanagerListAttachmentRoutingPolicyAssociations {
			networkmanager_ListAttachmentRoutingPolicyAssociations(cfg, client)
			return
		}
		if _networkmanagerListAttachments {
			networkmanager_ListAttachments(cfg, client)
			return
		}
		if _networkmanagerListConnectPeers {
			networkmanager_ListConnectPeers(cfg, client)
			return
		}
		if _networkmanagerListCoreNetworkPolicyVersions {
			networkmanager_ListCoreNetworkPolicyVersions(cfg, client)
			return
		}
		if _networkmanagerListCoreNetworkPrefixListAssociations {
			networkmanager_ListCoreNetworkPrefixListAssociations(cfg, client)
			return
		}
		if _networkmanagerListCoreNetworkRoutingInformation {
			networkmanager_ListCoreNetworkRoutingInformation(cfg, client)
			return
		}
		if _networkmanagerListCoreNetworks {
			networkmanager_ListCoreNetworks(cfg, client)
			return
		}
		if _networkmanagerListOrganizationServiceAccessStatus {
			networkmanager_ListOrganizationServiceAccessStatus(cfg, client)
			return
		}
		if _networkmanagerListPeerings {
			networkmanager_ListPeerings(cfg, client)
			return
		}
		if _networkmanagerListTagsForResource {
			networkmanager_ListTagsForResource(cfg, client)
			return
		}
		if _networkmanagerPutAttachmentRoutingPolicyLabel {
			networkmanager_PutAttachmentRoutingPolicyLabel(cfg, client)
			return
		}
		if _networkmanagerPutCoreNetworkPolicy {
			networkmanager_PutCoreNetworkPolicy(cfg, client)
			return
		}
		if _networkmanagerPutResourcePolicy {
			networkmanager_PutResourcePolicy(cfg, client)
			return
		}
		if _networkmanagerRegisterTransitGateway {
			networkmanager_RegisterTransitGateway(cfg, client)
			return
		}
		if _networkmanagerRejectAttachment {
			networkmanager_RejectAttachment(cfg, client)
			return
		}
		if _networkmanagerRemoveAttachmentRoutingPolicyLabel {
			networkmanager_RemoveAttachmentRoutingPolicyLabel(cfg, client)
			return
		}
		if _networkmanagerRestoreCoreNetworkPolicyVersion {
			networkmanager_RestoreCoreNetworkPolicyVersion(cfg, client)
			return
		}
		if _networkmanagerStartOrganizationServiceAccessUpdate {
			networkmanager_StartOrganizationServiceAccessUpdate(cfg, client)
			return
		}
		if _networkmanagerStartRouteAnalysis {
			networkmanager_StartRouteAnalysis(cfg, client)
			return
		}
		if _networkmanagerTagResource {
			networkmanager_TagResource(cfg, client)
			return
		}
		if _networkmanagerUntagResource {
			networkmanager_UntagResource(cfg, client)
			return
		}
		if _networkmanagerUpdateConnection {
			networkmanager_UpdateConnection(cfg, client)
			return
		}
		if _networkmanagerUpdateCoreNetwork {
			networkmanager_UpdateCoreNetwork(cfg, client)
			return
		}
		if _networkmanagerUpdateDevice {
			networkmanager_UpdateDevice(cfg, client)
			return
		}
		if _networkmanagerUpdateDirectConnectGatewayAttachment {
			networkmanager_UpdateDirectConnectGatewayAttachment(cfg, client)
			return
		}
		if _networkmanagerUpdateGlobalNetwork {
			networkmanager_UpdateGlobalNetwork(cfg, client)
			return
		}
		if _networkmanagerUpdateLink {
			networkmanager_UpdateLink(cfg, client)
			return
		}
		if _networkmanagerUpdateNetworkResourceMetadata {
			networkmanager_UpdateNetworkResourceMetadata(cfg, client)
			return
		}
		if _networkmanagerUpdateSite {
			networkmanager_UpdateSite(cfg, client)
			return
		}
		if _networkmanagerUpdateVpcAttachment {
			networkmanager_UpdateVpcAttachment(cfg, client)
			return
		}

	},
}

var (
	_networkmanagerAcceptAttachment                         bool
	_networkmanagerAssociateConnectPeer                     bool
	_networkmanagerAssociateCustomerGateway                 bool
	_networkmanagerAssociateLink                            bool
	_networkmanagerAssociateTransitGatewayConnectPeer       bool
	_networkmanagerCreateConnectAttachment                  bool
	_networkmanagerCreateConnectPeer                        bool
	_networkmanagerCreateConnection                         bool
	_networkmanagerCreateCoreNetwork                        bool
	_networkmanagerCreateCoreNetworkPrefixListAssociation   bool
	_networkmanagerCreateDevice                             bool
	_networkmanagerCreateDirectConnectGatewayAttachment     bool
	_networkmanagerCreateGlobalNetwork                      bool
	_networkmanagerCreateLink                               bool
	_networkmanagerCreateSite                               bool
	_networkmanagerCreateSiteToSiteVpnAttachment            bool
	_networkmanagerCreateTransitGatewayPeering              bool
	_networkmanagerCreateTransitGatewayRouteTableAttachment bool
	_networkmanagerCreateVpcAttachment                      bool
	_networkmanagerDeleteAttachment                         bool
	_networkmanagerDeleteConnectPeer                        bool
	_networkmanagerDeleteConnection                         bool
	_networkmanagerDeleteCoreNetwork                        bool
	_networkmanagerDeleteCoreNetworkPolicyVersion           bool
	_networkmanagerDeleteCoreNetworkPrefixListAssociation   bool
	_networkmanagerDeleteDevice                             bool
	_networkmanagerDeleteGlobalNetwork                      bool
	_networkmanagerDeleteLink                               bool
	_networkmanagerDeletePeering                            bool
	_networkmanagerDeleteResourcePolicy                     bool
	_networkmanagerDeleteSite                               bool
	_networkmanagerDeregisterTransitGateway                 bool
	_networkmanagerDescribeGlobalNetworks                   bool
	_networkmanagerDisassociateConnectPeer                  bool
	_networkmanagerDisassociateCustomerGateway              bool
	_networkmanagerDisassociateLink                         bool
	_networkmanagerDisassociateTransitGatewayConnectPeer    bool
	_networkmanagerExecuteCoreNetworkChangeSet              bool
	_networkmanagerGetConnectAttachment                     bool
	_networkmanagerGetConnectPeer                           bool
	_networkmanagerGetConnectPeerAssociations               bool
	_networkmanagerGetConnections                           bool
	_networkmanagerGetCoreNetwork                           bool
	_networkmanagerGetCoreNetworkChangeEvents               bool
	_networkmanagerGetCoreNetworkChangeSet                  bool
	_networkmanagerGetCoreNetworkPolicy                     bool
	_networkmanagerGetCustomerGatewayAssociations           bool
	_networkmanagerGetDevices                               bool
	_networkmanagerGetDirectConnectGatewayAttachment        bool
	_networkmanagerGetLinkAssociations                      bool
	_networkmanagerGetLinks                                 bool
	_networkmanagerGetNetworkResourceCounts                 bool
	_networkmanagerGetNetworkResourceRelationships          bool
	_networkmanagerGetNetworkResources                      bool
	_networkmanagerGetNetworkRoutes                         bool
	_networkmanagerGetNetworkTelemetry                      bool
	_networkmanagerGetResourcePolicy                        bool
	_networkmanagerGetRouteAnalysis                         bool
	_networkmanagerGetSiteToSiteVpnAttachment               bool
	_networkmanagerGetSites                                 bool
	_networkmanagerGetTransitGatewayConnectPeerAssociations bool
	_networkmanagerGetTransitGatewayPeering                 bool
	_networkmanagerGetTransitGatewayRegistrations           bool
	_networkmanagerGetTransitGatewayRouteTableAttachment    bool
	_networkmanagerGetVpcAttachment                         bool
	_networkmanagerListAttachmentRoutingPolicyAssociations  bool
	_networkmanagerListAttachments                          bool
	_networkmanagerListConnectPeers                         bool
	_networkmanagerListCoreNetworkPolicyVersions            bool
	_networkmanagerListCoreNetworkPrefixListAssociations    bool
	_networkmanagerListCoreNetworkRoutingInformation        bool
	_networkmanagerListCoreNetworks                         bool
	_networkmanagerListOrganizationServiceAccessStatus      bool
	_networkmanagerListPeerings                             bool
	_networkmanagerListTagsForResource                      bool
	_networkmanagerPutAttachmentRoutingPolicyLabel          bool
	_networkmanagerPutCoreNetworkPolicy                     bool
	_networkmanagerPutResourcePolicy                        bool
	_networkmanagerRegisterTransitGateway                   bool
	_networkmanagerRejectAttachment                         bool
	_networkmanagerRemoveAttachmentRoutingPolicyLabel       bool
	_networkmanagerRestoreCoreNetworkPolicyVersion          bool
	_networkmanagerStartOrganizationServiceAccessUpdate     bool
	_networkmanagerStartRouteAnalysis                       bool
	_networkmanagerTagResource                              bool
	_networkmanagerUntagResource                            bool
	_networkmanagerUpdateConnection                         bool
	_networkmanagerUpdateCoreNetwork                        bool
	_networkmanagerUpdateDevice                             bool
	_networkmanagerUpdateDirectConnectGatewayAttachment     bool
	_networkmanagerUpdateGlobalNetwork                      bool
	_networkmanagerUpdateLink                               bool
	_networkmanagerUpdateNetworkResourceMetadata            bool
	_networkmanagerUpdateSite                               bool
	_networkmanagerUpdateVpcAttachment                      bool

	_networkmanagerAccountId                     string
	_networkmanagerAction                        string
	_networkmanagerAddSubnetArns                 []string
	_networkmanagerAlias                         string
	_networkmanagerAttachmentId                  string
	_networkmanagerAttachmentType                string
	_networkmanagerAWSLocation                   string
	_networkmanagerAwsRegion                     string
	_networkmanagerBandwidth                     string
	_networkmanagerBgpOptions                    string
	_networkmanagerClientToken                   string
	_networkmanagerCommunityMatches              []string
	_networkmanagerConnectAttachmentId           string
	_networkmanagerConnectPeerId                 string
	_networkmanagerConnectPeerIds                []string
	_networkmanagerConnectedDeviceId             string
	_networkmanagerConnectedLinkId               string
	_networkmanagerConnectionId                  string
	_networkmanagerConnectionIds                 []string
	_networkmanagerCoreNetworkAddress            string
	_networkmanagerCoreNetworkId                 string
	_networkmanagerCustomerGatewayArn            string
	_networkmanagerCustomerGatewayArns           []string
	_networkmanagerDescription                   string
	_networkmanagerDestination                   string
	_networkmanagerDestinationFilters            string
	_networkmanagerDeviceId                      string
	_networkmanagerDeviceIds                     []string
	_networkmanagerDirectConnectGatewayArn       string
	_networkmanagerEdgeLocation                  string
	_networkmanagerEdgeLocations                 []string
	_networkmanagerExactAsPathMatches            []string
	_networkmanagerExactCidrMatches              []string
	_networkmanagerGlobalNetworkId               string
	_networkmanagerGlobalNetworkIds              []string
	_networkmanagerIncludeReturnPath             string
	_networkmanagerInsideCidrBlocks              []string
	_networkmanagerLatestVersionId               string
	_networkmanagerLinkId                        string
	_networkmanagerLinkIds                       []string
	_networkmanagerLocalPreferenceMatches        []string
	_networkmanagerLocation                      string
	_networkmanagerLongestPrefixMatches          []string
	_networkmanagerMaxResults                    string
	_networkmanagerMedMatches                    []string
	_networkmanagerMetadata                      string
	_networkmanagerModel                         string
	_networkmanagerNextHopFilters                string
	_networkmanagerNextToken                     string
	_networkmanagerOptions                       string
	_networkmanagerPeerAddress                   string
	_networkmanagerPeeringId                     string
	_networkmanagerPeeringType                   string
	_networkmanagerPolicyDocument                string
	_networkmanagerPolicyVersionId               string
	_networkmanagerPrefixListAlias               string
	_networkmanagerPrefixListArn                 string
	_networkmanagerPrefixListIds                 []string
	_networkmanagerProvider                      string
	_networkmanagerRegisteredGatewayArn          string
	_networkmanagerRemoveSubnetArns              []string
	_networkmanagerResourceArn                   string
	_networkmanagerResourceType                  string
	_networkmanagerRouteAnalysisId               string
	_networkmanagerRouteTableIdentifier          string
	_networkmanagerRoutingPolicyLabel            string
	_networkmanagerSegmentName                   string
	_networkmanagerSerialNumber                  string
	_networkmanagerSiteId                        string
	_networkmanagerSiteIds                       []string
	_networkmanagerSource                        string
	_networkmanagerState                         string
	_networkmanagerStates                        string
	_networkmanagerSubnetArn                     string
	_networkmanagerSubnetArns                    []string
	_networkmanagerSubnetOfMatches               []string
	_networkmanagerSupernetOfMatches             []string
	_networkmanagerTagKeys                       []string
	_networkmanagerTags                          string
	_networkmanagerTransitGatewayArn             string
	_networkmanagerTransitGatewayArns            []string
	_networkmanagerTransitGatewayConnectPeerArn  string
	_networkmanagerTransitGatewayConnectPeerArns []string
	_networkmanagerTransitGatewayRouteTableArn   string
	_networkmanagerTransportAttachmentId         string
	_networkmanagerType                          string
	_networkmanagerTypes                         string
	_networkmanagerUseMiddleboxes                string
	_networkmanagerVendor                        string
	_networkmanagerVpcArn                        string
	_networkmanagerVpnConnectionArn              string
)

// Accepts a core network attachment request.
// Once the attachment request is accepted by a core network owner, the attachment
// is created and connected to a core network.
func networkmanager_AcceptAttachment(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.AcceptAttachmentInput{
		// AttachmentId: *string, // Required
	}

	if len(_networkmanagerAttachmentId) > 0 {
		input.AttachmentId = aws.String(_networkmanagerAttachmentId)
	}

	if resp, err := client.AcceptAttachment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a core network Connect peer with a device and optionally, with a
// link.
//
// If you specify a link, it must be associated with the specified device. You can
// only associate core network Connect peers that have been created on a core
// network Connect attachment on a core network.
func networkmanager_AssociateConnectPeer(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.AssociateConnectPeerInput{
		// ConnectPeerId: *string, // Required
		// DeviceId: *string, // Required
		// GlobalNetworkId: *string, // Required
	}

	if len(_networkmanagerConnectPeerId) > 0 {
		input.ConnectPeerId = aws.String(_networkmanagerConnectPeerId)
	}
	if len(_networkmanagerDeviceId) > 0 {
		input.DeviceId = aws.String(_networkmanagerDeviceId)
	}
	if len(_networkmanagerGlobalNetworkId) > 0 {
		input.GlobalNetworkId = aws.String(_networkmanagerGlobalNetworkId)
	}
	if len(_networkmanagerLinkId) > 0 {
		input.LinkId = aws.String(_networkmanagerLinkId)
	}

	if resp, err := client.AssociateConnectPeer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a customer gateway with a device and optionally, with a link. If you
// specify a link, it must be associated with the specified device.
//
// You can only associate customer gateways that are connected to a VPN attachment
// on a transit gateway or core network registered in your global network. When you
// register a transit gateway or core network, customer gateways that are connected
// to the transit gateway are automatically included in the global network. To list
// customer gateways that are connected to a transit gateway, use the [DescribeVpnConnections]EC2 API and
// filter by transit-gateway-id .
//
// You cannot associate a customer gateway with more than one device and link.
//
// [DescribeVpnConnections]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeVpnConnections.html
func networkmanager_AssociateCustomerGateway(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.AssociateCustomerGatewayInput{
		// CustomerGatewayArn: *string, // Required
		// DeviceId: *string, // Required
		// GlobalNetworkId: *string, // Required
	}

	if len(_networkmanagerCustomerGatewayArn) > 0 {
		input.CustomerGatewayArn = aws.String(_networkmanagerCustomerGatewayArn)
	}
	if len(_networkmanagerDeviceId) > 0 {
		input.DeviceId = aws.String(_networkmanagerDeviceId)
	}
	if len(_networkmanagerGlobalNetworkId) > 0 {
		input.GlobalNetworkId = aws.String(_networkmanagerGlobalNetworkId)
	}
	if len(_networkmanagerLinkId) > 0 {
		input.LinkId = aws.String(_networkmanagerLinkId)
	}

	if resp, err := client.AssociateCustomerGateway(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a link to a device. A device can be associated to multiple links and
// a link can be associated to multiple devices. The device and link must be in the
// same global network and the same site.
func networkmanager_AssociateLink(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.AssociateLinkInput{
		// DeviceId: *string, // Required
		// GlobalNetworkId: *string, // Required
		// LinkId: *string, // Required
	}

	if len(_networkmanagerDeviceId) > 0 {
		input.DeviceId = aws.String(_networkmanagerDeviceId)
	}
	if len(_networkmanagerGlobalNetworkId) > 0 {
		input.GlobalNetworkId = aws.String(_networkmanagerGlobalNetworkId)
	}
	if len(_networkmanagerLinkId) > 0 {
		input.LinkId = aws.String(_networkmanagerLinkId)
	}

	if resp, err := client.AssociateLink(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a transit gateway Connect peer with a device, and optionally, with a
// link. If you specify a link, it must be associated with the specified device.
//
// You can only associate transit gateway Connect peers that have been created on
// a transit gateway that's registered in your global network.
//
// You cannot associate a transit gateway Connect peer with more than one device
// and link.
func networkmanager_AssociateTransitGatewayConnectPeer(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.AssociateTransitGatewayConnectPeerInput{
		// DeviceId: *string, // Required
		// GlobalNetworkId: *string, // Required
		// TransitGatewayConnectPeerArn: *string, // Required
	}

	if len(_networkmanagerDeviceId) > 0 {
		input.DeviceId = aws.String(_networkmanagerDeviceId)
	}
	if len(_networkmanagerGlobalNetworkId) > 0 {
		input.GlobalNetworkId = aws.String(_networkmanagerGlobalNetworkId)
	}
	if len(_networkmanagerTransitGatewayConnectPeerArn) > 0 {
		input.TransitGatewayConnectPeerArn = aws.String(_networkmanagerTransitGatewayConnectPeerArn)
	}
	if len(_networkmanagerLinkId) > 0 {
		input.LinkId = aws.String(_networkmanagerLinkId)
	}

	if resp, err := client.AssociateTransitGatewayConnectPeer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a core network Connect attachment from a specified core network
// attachment.
//
// A core network Connect attachment is a GRE-based tunnel attachment that you can
// use to establish a connection between a core network and an appliance. A core
// network Connect attachment uses an existing VPC attachment as the underlying
// transport mechanism.
func networkmanager_CreateConnectAttachment(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.CreateConnectAttachmentInput{
		// CoreNetworkId: *string, // Required
		// EdgeLocation: *string, // Required
		// Options: *types.ConnectAttachmentOptions, // Required
		// TransportAttachmentId: *string, // Required
	}

	if len(_networkmanagerCoreNetworkId) > 0 {
		input.CoreNetworkId = aws.String(_networkmanagerCoreNetworkId)
	}
	if len(_networkmanagerEdgeLocation) > 0 {
		input.EdgeLocation = aws.String(_networkmanagerEdgeLocation)
	}
	if len(_networkmanagerOptions) > 0 {
		if err := assignInputField(input, "Options", _networkmanagerOptions); err != nil {
			log.Errorf("invalid --options: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerTransportAttachmentId) > 0 {
		input.TransportAttachmentId = aws.String(_networkmanagerTransportAttachmentId)
	}
	if len(_networkmanagerClientToken) > 0 {
		input.ClientToken = aws.String(_networkmanagerClientToken)
	}
	if len(_networkmanagerRoutingPolicyLabel) > 0 {
		input.RoutingPolicyLabel = aws.String(_networkmanagerRoutingPolicyLabel)
	}
	if len(_networkmanagerTags) > 0 {
		if err := assignInputField(input, "Tags", _networkmanagerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateConnectAttachment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a core network Connect peer for a specified core network connect
// attachment between a core network and an appliance. The peer address and transit
// gateway address must be the same IP address family (IPv4 or IPv6).
func networkmanager_CreateConnectPeer(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.CreateConnectPeerInput{
		// ConnectAttachmentId: *string, // Required
		// PeerAddress: *string, // Required
	}

	if len(_networkmanagerConnectAttachmentId) > 0 {
		input.ConnectAttachmentId = aws.String(_networkmanagerConnectAttachmentId)
	}
	if len(_networkmanagerPeerAddress) > 0 {
		input.PeerAddress = aws.String(_networkmanagerPeerAddress)
	}
	if len(_networkmanagerBgpOptions) > 0 {
		if err := assignInputField(input, "BgpOptions", _networkmanagerBgpOptions); err != nil {
			log.Errorf("invalid --bgp-options: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerClientToken) > 0 {
		input.ClientToken = aws.String(_networkmanagerClientToken)
	}
	if len(_networkmanagerCoreNetworkAddress) > 0 {
		input.CoreNetworkAddress = aws.String(_networkmanagerCoreNetworkAddress)
	}
	if len(_networkmanagerInsideCidrBlocks) > 0 {
		input.InsideCidrBlocks = append([]string(nil), _networkmanagerInsideCidrBlocks...)
	}
	if len(_networkmanagerSubnetArn) > 0 {
		input.SubnetArn = aws.String(_networkmanagerSubnetArn)
	}
	if len(_networkmanagerTags) > 0 {
		if err := assignInputField(input, "Tags", _networkmanagerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateConnectPeer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a connection between two devices. The devices can be a physical or
// virtual appliance that connects to a third-party appliance in a VPC, or a
// physical appliance that connects to another physical appliance in an on-premises
// network.
func networkmanager_CreateConnection(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.CreateConnectionInput{
		// ConnectedDeviceId: *string, // Required
		// DeviceId: *string, // Required
		// GlobalNetworkId: *string, // Required
	}

	if len(_networkmanagerConnectedDeviceId) > 0 {
		input.ConnectedDeviceId = aws.String(_networkmanagerConnectedDeviceId)
	}
	if len(_networkmanagerDeviceId) > 0 {
		input.DeviceId = aws.String(_networkmanagerDeviceId)
	}
	if len(_networkmanagerGlobalNetworkId) > 0 {
		input.GlobalNetworkId = aws.String(_networkmanagerGlobalNetworkId)
	}
	if len(_networkmanagerConnectedLinkId) > 0 {
		input.ConnectedLinkId = aws.String(_networkmanagerConnectedLinkId)
	}
	if len(_networkmanagerDescription) > 0 {
		input.Description = aws.String(_networkmanagerDescription)
	}
	if len(_networkmanagerLinkId) > 0 {
		input.LinkId = aws.String(_networkmanagerLinkId)
	}
	if len(_networkmanagerTags) > 0 {
		if err := assignInputField(input, "Tags", _networkmanagerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a core network as part of your global network, and optionally, with a
// core network policy.
func networkmanager_CreateCoreNetwork(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.CreateCoreNetworkInput{
		// GlobalNetworkId: *string, // Required
	}

	if len(_networkmanagerGlobalNetworkId) > 0 {
		input.GlobalNetworkId = aws.String(_networkmanagerGlobalNetworkId)
	}
	if len(_networkmanagerClientToken) > 0 {
		input.ClientToken = aws.String(_networkmanagerClientToken)
	}
	if len(_networkmanagerDescription) > 0 {
		input.Description = aws.String(_networkmanagerDescription)
	}
	if len(_networkmanagerPolicyDocument) > 0 {
		input.PolicyDocument = aws.String(_networkmanagerPolicyDocument)
	}
	if len(_networkmanagerTags) > 0 {
		if err := assignInputField(input, "Tags", _networkmanagerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCoreNetwork(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an association between a core network and a prefix list for routing
// control.
func networkmanager_CreateCoreNetworkPrefixListAssociation(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.CreateCoreNetworkPrefixListAssociationInput{
		// CoreNetworkId: *string, // Required
		// PrefixListAlias: *string, // Required
		// PrefixListArn: *string, // Required
	}

	if len(_networkmanagerCoreNetworkId) > 0 {
		input.CoreNetworkId = aws.String(_networkmanagerCoreNetworkId)
	}
	if len(_networkmanagerPrefixListAlias) > 0 {
		input.PrefixListAlias = aws.String(_networkmanagerPrefixListAlias)
	}
	if len(_networkmanagerPrefixListArn) > 0 {
		input.PrefixListArn = aws.String(_networkmanagerPrefixListArn)
	}
	if len(_networkmanagerClientToken) > 0 {
		input.ClientToken = aws.String(_networkmanagerClientToken)
	}

	if resp, err := client.CreateCoreNetworkPrefixListAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new device in a global network. If you specify both a site ID and a
// location, the location of the site is used for visualization in the Network
// Manager console.
func networkmanager_CreateDevice(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.CreateDeviceInput{
		// GlobalNetworkId: *string, // Required
	}

	if len(_networkmanagerGlobalNetworkId) > 0 {
		input.GlobalNetworkId = aws.String(_networkmanagerGlobalNetworkId)
	}
	if len(_networkmanagerAWSLocation) > 0 {
		if err := assignInputField(input, "AWSLocation", _networkmanagerAWSLocation); err != nil {
			log.Errorf("invalid --aws-location: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerDescription) > 0 {
		input.Description = aws.String(_networkmanagerDescription)
	}
	if len(_networkmanagerLocation) > 0 {
		if err := assignInputField(input, "Location", _networkmanagerLocation); err != nil {
			log.Errorf("invalid --location: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerModel) > 0 {
		input.Model = aws.String(_networkmanagerModel)
	}
	if len(_networkmanagerSerialNumber) > 0 {
		input.SerialNumber = aws.String(_networkmanagerSerialNumber)
	}
	if len(_networkmanagerSiteId) > 0 {
		input.SiteId = aws.String(_networkmanagerSiteId)
	}
	if len(_networkmanagerTags) > 0 {
		if err := assignInputField(input, "Tags", _networkmanagerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerType) > 0 {
		input.Type = aws.String(_networkmanagerType)
	}
	if len(_networkmanagerVendor) > 0 {
		input.Vendor = aws.String(_networkmanagerVendor)
	}

	if resp, err := client.CreateDevice(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon Web Services Direct Connect gateway attachment
func networkmanager_CreateDirectConnectGatewayAttachment(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.CreateDirectConnectGatewayAttachmentInput{
		// CoreNetworkId: *string, // Required
		// DirectConnectGatewayArn: *string, // Required
		// EdgeLocations: []string, // Required
	}

	if len(_networkmanagerCoreNetworkId) > 0 {
		input.CoreNetworkId = aws.String(_networkmanagerCoreNetworkId)
	}
	if len(_networkmanagerDirectConnectGatewayArn) > 0 {
		input.DirectConnectGatewayArn = aws.String(_networkmanagerDirectConnectGatewayArn)
	}
	if len(_networkmanagerEdgeLocations) > 0 {
		input.EdgeLocations = append([]string(nil), _networkmanagerEdgeLocations...)
	}
	if len(_networkmanagerClientToken) > 0 {
		input.ClientToken = aws.String(_networkmanagerClientToken)
	}
	if len(_networkmanagerRoutingPolicyLabel) > 0 {
		input.RoutingPolicyLabel = aws.String(_networkmanagerRoutingPolicyLabel)
	}
	if len(_networkmanagerTags) > 0 {
		if err := assignInputField(input, "Tags", _networkmanagerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDirectConnectGatewayAttachment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new, empty global network.
func networkmanager_CreateGlobalNetwork(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.CreateGlobalNetworkInput{}

	if len(_networkmanagerDescription) > 0 {
		input.Description = aws.String(_networkmanagerDescription)
	}
	if len(_networkmanagerTags) > 0 {
		if err := assignInputField(input, "Tags", _networkmanagerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateGlobalNetwork(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new link for a specified site.
func networkmanager_CreateLink(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.CreateLinkInput{
		// Bandwidth: *types.Bandwidth, // Required
		// GlobalNetworkId: *string, // Required
		// SiteId: *string, // Required
	}

	if len(_networkmanagerBandwidth) > 0 {
		if err := assignInputField(input, "Bandwidth", _networkmanagerBandwidth); err != nil {
			log.Errorf("invalid --bandwidth: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerGlobalNetworkId) > 0 {
		input.GlobalNetworkId = aws.String(_networkmanagerGlobalNetworkId)
	}
	if len(_networkmanagerSiteId) > 0 {
		input.SiteId = aws.String(_networkmanagerSiteId)
	}
	if len(_networkmanagerDescription) > 0 {
		input.Description = aws.String(_networkmanagerDescription)
	}
	if len(_networkmanagerProvider) > 0 {
		input.Provider = aws.String(_networkmanagerProvider)
	}
	if len(_networkmanagerTags) > 0 {
		if err := assignInputField(input, "Tags", _networkmanagerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerType) > 0 {
		input.Type = aws.String(_networkmanagerType)
	}

	if resp, err := client.CreateLink(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new site in a global network.
func networkmanager_CreateSite(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.CreateSiteInput{
		// GlobalNetworkId: *string, // Required
	}

	if len(_networkmanagerGlobalNetworkId) > 0 {
		input.GlobalNetworkId = aws.String(_networkmanagerGlobalNetworkId)
	}
	if len(_networkmanagerDescription) > 0 {
		input.Description = aws.String(_networkmanagerDescription)
	}
	if len(_networkmanagerLocation) > 0 {
		if err := assignInputField(input, "Location", _networkmanagerLocation); err != nil {
			log.Errorf("invalid --location: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerTags) > 0 {
		if err := assignInputField(input, "Tags", _networkmanagerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSite(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon Web Services site-to-site VPN attachment on an edge location
// of a core network.
func networkmanager_CreateSiteToSiteVpnAttachment(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.CreateSiteToSiteVpnAttachmentInput{
		// CoreNetworkId: *string, // Required
		// VpnConnectionArn: *string, // Required
	}

	if len(_networkmanagerCoreNetworkId) > 0 {
		input.CoreNetworkId = aws.String(_networkmanagerCoreNetworkId)
	}
	if len(_networkmanagerVpnConnectionArn) > 0 {
		input.VpnConnectionArn = aws.String(_networkmanagerVpnConnectionArn)
	}
	if len(_networkmanagerClientToken) > 0 {
		input.ClientToken = aws.String(_networkmanagerClientToken)
	}
	if len(_networkmanagerRoutingPolicyLabel) > 0 {
		input.RoutingPolicyLabel = aws.String(_networkmanagerRoutingPolicyLabel)
	}
	if len(_networkmanagerTags) > 0 {
		if err := assignInputField(input, "Tags", _networkmanagerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSiteToSiteVpnAttachment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a transit gateway peering connection.
func networkmanager_CreateTransitGatewayPeering(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.CreateTransitGatewayPeeringInput{
		// CoreNetworkId: *string, // Required
		// TransitGatewayArn: *string, // Required
	}

	if len(_networkmanagerCoreNetworkId) > 0 {
		input.CoreNetworkId = aws.String(_networkmanagerCoreNetworkId)
	}
	if len(_networkmanagerTransitGatewayArn) > 0 {
		input.TransitGatewayArn = aws.String(_networkmanagerTransitGatewayArn)
	}
	if len(_networkmanagerClientToken) > 0 {
		input.ClientToken = aws.String(_networkmanagerClientToken)
	}
	if len(_networkmanagerTags) > 0 {
		if err := assignInputField(input, "Tags", _networkmanagerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTransitGatewayPeering(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a transit gateway route table attachment.
func networkmanager_CreateTransitGatewayRouteTableAttachment(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.CreateTransitGatewayRouteTableAttachmentInput{
		// PeeringId: *string, // Required
		// TransitGatewayRouteTableArn: *string, // Required
	}

	if len(_networkmanagerPeeringId) > 0 {
		input.PeeringId = aws.String(_networkmanagerPeeringId)
	}
	if len(_networkmanagerTransitGatewayRouteTableArn) > 0 {
		input.TransitGatewayRouteTableArn = aws.String(_networkmanagerTransitGatewayRouteTableArn)
	}
	if len(_networkmanagerClientToken) > 0 {
		input.ClientToken = aws.String(_networkmanagerClientToken)
	}
	if len(_networkmanagerRoutingPolicyLabel) > 0 {
		input.RoutingPolicyLabel = aws.String(_networkmanagerRoutingPolicyLabel)
	}
	if len(_networkmanagerTags) > 0 {
		if err := assignInputField(input, "Tags", _networkmanagerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTransitGatewayRouteTableAttachment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a VPC attachment on an edge location of a core network.
func networkmanager_CreateVpcAttachment(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.CreateVpcAttachmentInput{
		// CoreNetworkId: *string, // Required
		// SubnetArns: []string, // Required
		// VpcArn: *string, // Required
	}

	if len(_networkmanagerCoreNetworkId) > 0 {
		input.CoreNetworkId = aws.String(_networkmanagerCoreNetworkId)
	}
	if len(_networkmanagerSubnetArns) > 0 {
		input.SubnetArns = append([]string(nil), _networkmanagerSubnetArns...)
	}
	if len(_networkmanagerVpcArn) > 0 {
		input.VpcArn = aws.String(_networkmanagerVpcArn)
	}
	if len(_networkmanagerClientToken) > 0 {
		input.ClientToken = aws.String(_networkmanagerClientToken)
	}
	if len(_networkmanagerOptions) > 0 {
		if err := assignInputField(input, "Options", _networkmanagerOptions); err != nil {
			log.Errorf("invalid --options: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerRoutingPolicyLabel) > 0 {
		input.RoutingPolicyLabel = aws.String(_networkmanagerRoutingPolicyLabel)
	}
	if len(_networkmanagerTags) > 0 {
		if err := assignInputField(input, "Tags", _networkmanagerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateVpcAttachment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an attachment. Supports all attachment types.
func networkmanager_DeleteAttachment(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.DeleteAttachmentInput{
		// AttachmentId: *string, // Required
	}

	if len(_networkmanagerAttachmentId) > 0 {
		input.AttachmentId = aws.String(_networkmanagerAttachmentId)
	}

	if resp, err := client.DeleteAttachment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Connect peer.
func networkmanager_DeleteConnectPeer(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.DeleteConnectPeerInput{
		// ConnectPeerId: *string, // Required
	}

	if len(_networkmanagerConnectPeerId) > 0 {
		input.ConnectPeerId = aws.String(_networkmanagerConnectPeerId)
	}

	if resp, err := client.DeleteConnectPeer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified connection in your global network.
func networkmanager_DeleteConnection(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.DeleteConnectionInput{
		// ConnectionId: *string, // Required
		// GlobalNetworkId: *string, // Required
	}

	if len(_networkmanagerConnectionId) > 0 {
		input.ConnectionId = aws.String(_networkmanagerConnectionId)
	}
	if len(_networkmanagerGlobalNetworkId) > 0 {
		input.GlobalNetworkId = aws.String(_networkmanagerGlobalNetworkId)
	}

	if resp, err := client.DeleteConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a core network along with all core network policies. This can only be
// done if there are no attachments on a core network.
func networkmanager_DeleteCoreNetwork(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.DeleteCoreNetworkInput{
		// CoreNetworkId: *string, // Required
	}

	if len(_networkmanagerCoreNetworkId) > 0 {
		input.CoreNetworkId = aws.String(_networkmanagerCoreNetworkId)
	}

	if resp, err := client.DeleteCoreNetwork(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a policy version from a core network. You can't delete the current LIVE
// policy.
func networkmanager_DeleteCoreNetworkPolicyVersion(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.DeleteCoreNetworkPolicyVersionInput{
		// CoreNetworkId: *string, // Required
		// PolicyVersionId: *int32, // Required
	}

	if len(_networkmanagerCoreNetworkId) > 0 {
		input.CoreNetworkId = aws.String(_networkmanagerCoreNetworkId)
	}
	if len(_networkmanagerPolicyVersionId) > 0 {
		if err := assignInputField(input, "PolicyVersionId", _networkmanagerPolicyVersionId); err != nil {
			log.Errorf("invalid --policy-version-id: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteCoreNetworkPolicyVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an association between a core network and a prefix list.
func networkmanager_DeleteCoreNetworkPrefixListAssociation(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.DeleteCoreNetworkPrefixListAssociationInput{
		// CoreNetworkId: *string, // Required
		// PrefixListArn: *string, // Required
	}

	if len(_networkmanagerCoreNetworkId) > 0 {
		input.CoreNetworkId = aws.String(_networkmanagerCoreNetworkId)
	}
	if len(_networkmanagerPrefixListArn) > 0 {
		input.PrefixListArn = aws.String(_networkmanagerPrefixListArn)
	}

	if resp, err := client.DeleteCoreNetworkPrefixListAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing device. You must first disassociate the device from any
// links and customer gateways.
func networkmanager_DeleteDevice(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.DeleteDeviceInput{
		// DeviceId: *string, // Required
		// GlobalNetworkId: *string, // Required
	}

	if len(_networkmanagerDeviceId) > 0 {
		input.DeviceId = aws.String(_networkmanagerDeviceId)
	}
	if len(_networkmanagerGlobalNetworkId) > 0 {
		input.GlobalNetworkId = aws.String(_networkmanagerGlobalNetworkId)
	}

	if resp, err := client.DeleteDevice(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing global network. You must first delete all global network
// objects (devices, links, and sites), deregister all transit gateways, and delete
// any core networks.
func networkmanager_DeleteGlobalNetwork(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.DeleteGlobalNetworkInput{
		// GlobalNetworkId: *string, // Required
	}

	if len(_networkmanagerGlobalNetworkId) > 0 {
		input.GlobalNetworkId = aws.String(_networkmanagerGlobalNetworkId)
	}

	if resp, err := client.DeleteGlobalNetwork(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing link. You must first disassociate the link from any devices
// and customer gateways.
func networkmanager_DeleteLink(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.DeleteLinkInput{
		// GlobalNetworkId: *string, // Required
		// LinkId: *string, // Required
	}

	if len(_networkmanagerGlobalNetworkId) > 0 {
		input.GlobalNetworkId = aws.String(_networkmanagerGlobalNetworkId)
	}
	if len(_networkmanagerLinkId) > 0 {
		input.LinkId = aws.String(_networkmanagerLinkId)
	}

	if resp, err := client.DeleteLink(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing peering connection.
func networkmanager_DeletePeering(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.DeletePeeringInput{
		// PeeringId: *string, // Required
	}

	if len(_networkmanagerPeeringId) > 0 {
		input.PeeringId = aws.String(_networkmanagerPeeringId)
	}

	if resp, err := client.DeletePeering(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a resource policy for the specified resource. This revokes the access
// of the principals specified in the resource policy.
func networkmanager_DeleteResourcePolicy(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.DeleteResourcePolicyInput{
		// ResourceArn: *string, // Required
	}

	if len(_networkmanagerResourceArn) > 0 {
		input.ResourceArn = aws.String(_networkmanagerResourceArn)
	}

	if resp, err := client.DeleteResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing site. The site cannot be associated with any device or link.
func networkmanager_DeleteSite(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.DeleteSiteInput{
		// GlobalNetworkId: *string, // Required
		// SiteId: *string, // Required
	}

	if len(_networkmanagerGlobalNetworkId) > 0 {
		input.GlobalNetworkId = aws.String(_networkmanagerGlobalNetworkId)
	}
	if len(_networkmanagerSiteId) > 0 {
		input.SiteId = aws.String(_networkmanagerSiteId)
	}

	if resp, err := client.DeleteSite(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deregisters a transit gateway from your global network. This action does not
// delete your transit gateway, or modify any of its attachments. This action
// removes any customer gateway associations.
func networkmanager_DeregisterTransitGateway(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.DeregisterTransitGatewayInput{
		// GlobalNetworkId: *string, // Required
		// TransitGatewayArn: *string, // Required
	}

	if len(_networkmanagerGlobalNetworkId) > 0 {
		input.GlobalNetworkId = aws.String(_networkmanagerGlobalNetworkId)
	}
	if len(_networkmanagerTransitGatewayArn) > 0 {
		input.TransitGatewayArn = aws.String(_networkmanagerTransitGatewayArn)
	}

	if resp, err := client.DeregisterTransitGateway(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes one or more global networks. By default, all global networks are
// described. To describe the objects in your global network, you must use the
// appropriate Get* action. For example, to list the transit gateways in your
// global network, use GetTransitGatewayRegistrations.
func networkmanager_DescribeGlobalNetworks(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.DescribeGlobalNetworksInput{}

	if len(_networkmanagerGlobalNetworkIds) > 0 {
		input.GlobalNetworkIds = append([]string(nil), _networkmanagerGlobalNetworkIds...)
	}
	if len(_networkmanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _networkmanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerNextToken) > 0 {
		input.NextToken = aws.String(_networkmanagerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeGlobalNetworks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*networkmanager.DescribeGlobalNetworksOutput
	p := networkmanager.NewDescribeGlobalNetworksPaginator(client, input)
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

// Disassociates a core network Connect peer from a device and a link.
func networkmanager_DisassociateConnectPeer(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.DisassociateConnectPeerInput{
		// ConnectPeerId: *string, // Required
		// GlobalNetworkId: *string, // Required
	}

	if len(_networkmanagerConnectPeerId) > 0 {
		input.ConnectPeerId = aws.String(_networkmanagerConnectPeerId)
	}
	if len(_networkmanagerGlobalNetworkId) > 0 {
		input.GlobalNetworkId = aws.String(_networkmanagerGlobalNetworkId)
	}

	if resp, err := client.DisassociateConnectPeer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a customer gateway from a device and a link.
func networkmanager_DisassociateCustomerGateway(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.DisassociateCustomerGatewayInput{
		// CustomerGatewayArn: *string, // Required
		// GlobalNetworkId: *string, // Required
	}

	if len(_networkmanagerCustomerGatewayArn) > 0 {
		input.CustomerGatewayArn = aws.String(_networkmanagerCustomerGatewayArn)
	}
	if len(_networkmanagerGlobalNetworkId) > 0 {
		input.GlobalNetworkId = aws.String(_networkmanagerGlobalNetworkId)
	}

	if resp, err := client.DisassociateCustomerGateway(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates an existing device from a link. You must first disassociate any
// customer gateways that are associated with the link.
func networkmanager_DisassociateLink(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.DisassociateLinkInput{
		// DeviceId: *string, // Required
		// GlobalNetworkId: *string, // Required
		// LinkId: *string, // Required
	}

	if len(_networkmanagerDeviceId) > 0 {
		input.DeviceId = aws.String(_networkmanagerDeviceId)
	}
	if len(_networkmanagerGlobalNetworkId) > 0 {
		input.GlobalNetworkId = aws.String(_networkmanagerGlobalNetworkId)
	}
	if len(_networkmanagerLinkId) > 0 {
		input.LinkId = aws.String(_networkmanagerLinkId)
	}

	if resp, err := client.DisassociateLink(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a transit gateway Connect peer from a device and link.
func networkmanager_DisassociateTransitGatewayConnectPeer(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.DisassociateTransitGatewayConnectPeerInput{
		// GlobalNetworkId: *string, // Required
		// TransitGatewayConnectPeerArn: *string, // Required
	}

	if len(_networkmanagerGlobalNetworkId) > 0 {
		input.GlobalNetworkId = aws.String(_networkmanagerGlobalNetworkId)
	}
	if len(_networkmanagerTransitGatewayConnectPeerArn) > 0 {
		input.TransitGatewayConnectPeerArn = aws.String(_networkmanagerTransitGatewayConnectPeerArn)
	}

	if resp, err := client.DisassociateTransitGatewayConnectPeer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Executes a change set on your core network. Deploys changes globally based on
// the policy submitted..
func networkmanager_ExecuteCoreNetworkChangeSet(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.ExecuteCoreNetworkChangeSetInput{
		// CoreNetworkId: *string, // Required
		// PolicyVersionId: *int32, // Required
	}

	if len(_networkmanagerCoreNetworkId) > 0 {
		input.CoreNetworkId = aws.String(_networkmanagerCoreNetworkId)
	}
	if len(_networkmanagerPolicyVersionId) > 0 {
		if err := assignInputField(input, "PolicyVersionId", _networkmanagerPolicyVersionId); err != nil {
			log.Errorf("invalid --policy-version-id: %s", err.Error())
			return
		}
	}

	if resp, err := client.ExecuteCoreNetworkChangeSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a core network Connect attachment.
func networkmanager_GetConnectAttachment(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.GetConnectAttachmentInput{
		// AttachmentId: *string, // Required
	}

	if len(_networkmanagerAttachmentId) > 0 {
		input.AttachmentId = aws.String(_networkmanagerAttachmentId)
	}

	if resp, err := client.GetConnectAttachment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a core network Connect peer.
func networkmanager_GetConnectPeer(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.GetConnectPeerInput{
		// ConnectPeerId: *string, // Required
	}

	if len(_networkmanagerConnectPeerId) > 0 {
		input.ConnectPeerId = aws.String(_networkmanagerConnectPeerId)
	}

	if resp, err := client.GetConnectPeer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a core network Connect peer associations.
func networkmanager_GetConnectPeerAssociations(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.GetConnectPeerAssociationsInput{
		// GlobalNetworkId: *string, // Required
	}

	if len(_networkmanagerGlobalNetworkId) > 0 {
		input.GlobalNetworkId = aws.String(_networkmanagerGlobalNetworkId)
	}
	if len(_networkmanagerConnectPeerIds) > 0 {
		input.ConnectPeerIds = append([]string(nil), _networkmanagerConnectPeerIds...)
	}
	if len(_networkmanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _networkmanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerNextToken) > 0 {
		input.NextToken = aws.String(_networkmanagerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetConnectPeerAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*networkmanager.GetConnectPeerAssociationsOutput
	p := networkmanager.NewGetConnectPeerAssociationsPaginator(client, input)
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

// Gets information about one or more of your connections in a global network.
func networkmanager_GetConnections(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.GetConnectionsInput{
		// GlobalNetworkId: *string, // Required
	}

	if len(_networkmanagerGlobalNetworkId) > 0 {
		input.GlobalNetworkId = aws.String(_networkmanagerGlobalNetworkId)
	}
	if len(_networkmanagerConnectionIds) > 0 {
		input.ConnectionIds = append([]string(nil), _networkmanagerConnectionIds...)
	}
	if len(_networkmanagerDeviceId) > 0 {
		input.DeviceId = aws.String(_networkmanagerDeviceId)
	}
	if len(_networkmanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _networkmanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerNextToken) > 0 {
		input.NextToken = aws.String(_networkmanagerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetConnections(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*networkmanager.GetConnectionsOutput
	p := networkmanager.NewGetConnectionsPaginator(client, input)
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

// Returns information about the LIVE policy for a core network.
func networkmanager_GetCoreNetwork(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.GetCoreNetworkInput{
		// CoreNetworkId: *string, // Required
	}

	if len(_networkmanagerCoreNetworkId) > 0 {
		input.CoreNetworkId = aws.String(_networkmanagerCoreNetworkId)
	}

	if resp, err := client.GetCoreNetwork(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a core network change event.
func networkmanager_GetCoreNetworkChangeEvents(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.GetCoreNetworkChangeEventsInput{
		// CoreNetworkId: *string, // Required
		// PolicyVersionId: *int32, // Required
	}

	if len(_networkmanagerCoreNetworkId) > 0 {
		input.CoreNetworkId = aws.String(_networkmanagerCoreNetworkId)
	}
	if len(_networkmanagerPolicyVersionId) > 0 {
		if err := assignInputField(input, "PolicyVersionId", _networkmanagerPolicyVersionId); err != nil {
			log.Errorf("invalid --policy-version-id: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _networkmanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerNextToken) > 0 {
		input.NextToken = aws.String(_networkmanagerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetCoreNetworkChangeEvents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*networkmanager.GetCoreNetworkChangeEventsOutput
	p := networkmanager.NewGetCoreNetworkChangeEventsPaginator(client, input)
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

// Returns a change set between the LIVE core network policy and a submitted
// policy.
func networkmanager_GetCoreNetworkChangeSet(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.GetCoreNetworkChangeSetInput{
		// CoreNetworkId: *string, // Required
		// PolicyVersionId: *int32, // Required
	}

	if len(_networkmanagerCoreNetworkId) > 0 {
		input.CoreNetworkId = aws.String(_networkmanagerCoreNetworkId)
	}
	if len(_networkmanagerPolicyVersionId) > 0 {
		if err := assignInputField(input, "PolicyVersionId", _networkmanagerPolicyVersionId); err != nil {
			log.Errorf("invalid --policy-version-id: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _networkmanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerNextToken) > 0 {
		input.NextToken = aws.String(_networkmanagerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetCoreNetworkChangeSet(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*networkmanager.GetCoreNetworkChangeSetOutput
	p := networkmanager.NewGetCoreNetworkChangeSetPaginator(client, input)
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

// Returns details about a core network policy. You can get details about your
// current live policy or any previous policy version.
func networkmanager_GetCoreNetworkPolicy(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.GetCoreNetworkPolicyInput{
		// CoreNetworkId: *string, // Required
	}

	if len(_networkmanagerCoreNetworkId) > 0 {
		input.CoreNetworkId = aws.String(_networkmanagerCoreNetworkId)
	}
	if len(_networkmanagerAlias) > 0 {
		if err := assignInputField(input, "Alias", _networkmanagerAlias); err != nil {
			log.Errorf("invalid --alias: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerPolicyVersionId) > 0 {
		if err := assignInputField(input, "PolicyVersionId", _networkmanagerPolicyVersionId); err != nil {
			log.Errorf("invalid --policy-version-id: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetCoreNetworkPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the association information for customer gateways that are associated with
// devices and links in your global network.
func networkmanager_GetCustomerGatewayAssociations(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.GetCustomerGatewayAssociationsInput{
		// GlobalNetworkId: *string, // Required
	}

	if len(_networkmanagerGlobalNetworkId) > 0 {
		input.GlobalNetworkId = aws.String(_networkmanagerGlobalNetworkId)
	}
	if len(_networkmanagerCustomerGatewayArns) > 0 {
		input.CustomerGatewayArns = append([]string(nil), _networkmanagerCustomerGatewayArns...)
	}
	if len(_networkmanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _networkmanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerNextToken) > 0 {
		input.NextToken = aws.String(_networkmanagerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetCustomerGatewayAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*networkmanager.GetCustomerGatewayAssociationsOutput
	p := networkmanager.NewGetCustomerGatewayAssociationsPaginator(client, input)
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

// Gets information about one or more of your devices in a global network.
func networkmanager_GetDevices(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.GetDevicesInput{
		// GlobalNetworkId: *string, // Required
	}

	if len(_networkmanagerGlobalNetworkId) > 0 {
		input.GlobalNetworkId = aws.String(_networkmanagerGlobalNetworkId)
	}
	if len(_networkmanagerDeviceIds) > 0 {
		input.DeviceIds = append([]string(nil), _networkmanagerDeviceIds...)
	}
	if len(_networkmanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _networkmanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerNextToken) > 0 {
		input.NextToken = aws.String(_networkmanagerNextToken)
	}
	if len(_networkmanagerSiteId) > 0 {
		input.SiteId = aws.String(_networkmanagerSiteId)
	}

	if disablePaginator() {
		if resp, err := client.GetDevices(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*networkmanager.GetDevicesOutput
	p := networkmanager.NewGetDevicesPaginator(client, input)
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

// Returns information about a specific Amazon Web Services Direct Connect gateway
// attachment.
func networkmanager_GetDirectConnectGatewayAttachment(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.GetDirectConnectGatewayAttachmentInput{
		// AttachmentId: *string, // Required
	}

	if len(_networkmanagerAttachmentId) > 0 {
		input.AttachmentId = aws.String(_networkmanagerAttachmentId)
	}

	if resp, err := client.GetDirectConnectGatewayAttachment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the link associations for a device or a link. Either the device ID or the
// link ID must be specified.
func networkmanager_GetLinkAssociations(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.GetLinkAssociationsInput{
		// GlobalNetworkId: *string, // Required
	}

	if len(_networkmanagerGlobalNetworkId) > 0 {
		input.GlobalNetworkId = aws.String(_networkmanagerGlobalNetworkId)
	}
	if len(_networkmanagerDeviceId) > 0 {
		input.DeviceId = aws.String(_networkmanagerDeviceId)
	}
	if len(_networkmanagerLinkId) > 0 {
		input.LinkId = aws.String(_networkmanagerLinkId)
	}
	if len(_networkmanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _networkmanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerNextToken) > 0 {
		input.NextToken = aws.String(_networkmanagerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetLinkAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*networkmanager.GetLinkAssociationsOutput
	p := networkmanager.NewGetLinkAssociationsPaginator(client, input)
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

// Gets information about one or more links in a specified global network.
// If you specify the site ID, you cannot specify the type or provider in the same
// request. You can specify the type and provider in the same request.
func networkmanager_GetLinks(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.GetLinksInput{
		// GlobalNetworkId: *string, // Required
	}

	if len(_networkmanagerGlobalNetworkId) > 0 {
		input.GlobalNetworkId = aws.String(_networkmanagerGlobalNetworkId)
	}
	if len(_networkmanagerLinkIds) > 0 {
		input.LinkIds = append([]string(nil), _networkmanagerLinkIds...)
	}
	if len(_networkmanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _networkmanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerNextToken) > 0 {
		input.NextToken = aws.String(_networkmanagerNextToken)
	}
	if len(_networkmanagerProvider) > 0 {
		input.Provider = aws.String(_networkmanagerProvider)
	}
	if len(_networkmanagerSiteId) > 0 {
		input.SiteId = aws.String(_networkmanagerSiteId)
	}
	if len(_networkmanagerType) > 0 {
		input.Type = aws.String(_networkmanagerType)
	}

	if disablePaginator() {
		if resp, err := client.GetLinks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*networkmanager.GetLinksOutput
	p := networkmanager.NewGetLinksPaginator(client, input)
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

// Gets the count of network resources, by resource type, for the specified global
// network.
func networkmanager_GetNetworkResourceCounts(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.GetNetworkResourceCountsInput{
		// GlobalNetworkId: *string, // Required
	}

	if len(_networkmanagerGlobalNetworkId) > 0 {
		input.GlobalNetworkId = aws.String(_networkmanagerGlobalNetworkId)
	}
	if len(_networkmanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _networkmanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerNextToken) > 0 {
		input.NextToken = aws.String(_networkmanagerNextToken)
	}
	if len(_networkmanagerResourceType) > 0 {
		input.ResourceType = aws.String(_networkmanagerResourceType)
	}

	if disablePaginator() {
		if resp, err := client.GetNetworkResourceCounts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*networkmanager.GetNetworkResourceCountsOutput
	p := networkmanager.NewGetNetworkResourceCountsPaginator(client, input)
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

// Gets the network resource relationships for the specified global network.
func networkmanager_GetNetworkResourceRelationships(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.GetNetworkResourceRelationshipsInput{
		// GlobalNetworkId: *string, // Required
	}

	if len(_networkmanagerGlobalNetworkId) > 0 {
		input.GlobalNetworkId = aws.String(_networkmanagerGlobalNetworkId)
	}
	if len(_networkmanagerAccountId) > 0 {
		input.AccountId = aws.String(_networkmanagerAccountId)
	}
	if len(_networkmanagerAwsRegion) > 0 {
		input.AwsRegion = aws.String(_networkmanagerAwsRegion)
	}
	if len(_networkmanagerCoreNetworkId) > 0 {
		input.CoreNetworkId = aws.String(_networkmanagerCoreNetworkId)
	}
	if len(_networkmanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _networkmanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerNextToken) > 0 {
		input.NextToken = aws.String(_networkmanagerNextToken)
	}
	if len(_networkmanagerRegisteredGatewayArn) > 0 {
		input.RegisteredGatewayArn = aws.String(_networkmanagerRegisteredGatewayArn)
	}
	if len(_networkmanagerResourceArn) > 0 {
		input.ResourceArn = aws.String(_networkmanagerResourceArn)
	}
	if len(_networkmanagerResourceType) > 0 {
		input.ResourceType = aws.String(_networkmanagerResourceType)
	}

	if disablePaginator() {
		if resp, err := client.GetNetworkResourceRelationships(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*networkmanager.GetNetworkResourceRelationshipsOutput
	p := networkmanager.NewGetNetworkResourceRelationshipsPaginator(client, input)
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

// Describes the network resources for the specified global network.
// The results include information from the corresponding Describe call for the
// resource, minus any sensitive information such as pre-shared keys.
func networkmanager_GetNetworkResources(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.GetNetworkResourcesInput{
		// GlobalNetworkId: *string, // Required
	}

	if len(_networkmanagerGlobalNetworkId) > 0 {
		input.GlobalNetworkId = aws.String(_networkmanagerGlobalNetworkId)
	}
	if len(_networkmanagerAccountId) > 0 {
		input.AccountId = aws.String(_networkmanagerAccountId)
	}
	if len(_networkmanagerAwsRegion) > 0 {
		input.AwsRegion = aws.String(_networkmanagerAwsRegion)
	}
	if len(_networkmanagerCoreNetworkId) > 0 {
		input.CoreNetworkId = aws.String(_networkmanagerCoreNetworkId)
	}
	if len(_networkmanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _networkmanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerNextToken) > 0 {
		input.NextToken = aws.String(_networkmanagerNextToken)
	}
	if len(_networkmanagerRegisteredGatewayArn) > 0 {
		input.RegisteredGatewayArn = aws.String(_networkmanagerRegisteredGatewayArn)
	}
	if len(_networkmanagerResourceArn) > 0 {
		input.ResourceArn = aws.String(_networkmanagerResourceArn)
	}
	if len(_networkmanagerResourceType) > 0 {
		input.ResourceType = aws.String(_networkmanagerResourceType)
	}

	if disablePaginator() {
		if resp, err := client.GetNetworkResources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*networkmanager.GetNetworkResourcesOutput
	p := networkmanager.NewGetNetworkResourcesPaginator(client, input)
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

// Gets the network routes of the specified global network.
func networkmanager_GetNetworkRoutes(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.GetNetworkRoutesInput{
		// GlobalNetworkId: *string, // Required
		// RouteTableIdentifier: *types.RouteTableIdentifier, // Required
	}

	if len(_networkmanagerGlobalNetworkId) > 0 {
		input.GlobalNetworkId = aws.String(_networkmanagerGlobalNetworkId)
	}
	if len(_networkmanagerRouteTableIdentifier) > 0 {
		if err := assignInputField(input, "RouteTableIdentifier", _networkmanagerRouteTableIdentifier); err != nil {
			log.Errorf("invalid --route-table-identifier: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerDestinationFilters) > 0 {
		if err := assignInputField(input, "DestinationFilters", _networkmanagerDestinationFilters); err != nil {
			log.Errorf("invalid --destination-filters: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerExactCidrMatches) > 0 {
		input.ExactCidrMatches = append([]string(nil), _networkmanagerExactCidrMatches...)
	}
	if len(_networkmanagerLongestPrefixMatches) > 0 {
		input.LongestPrefixMatches = append([]string(nil), _networkmanagerLongestPrefixMatches...)
	}
	if len(_networkmanagerPrefixListIds) > 0 {
		input.PrefixListIds = append([]string(nil), _networkmanagerPrefixListIds...)
	}
	if len(_networkmanagerStates) > 0 {
		if err := assignInputField(input, "States", _networkmanagerStates); err != nil {
			log.Errorf("invalid --states: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerSubnetOfMatches) > 0 {
		input.SubnetOfMatches = append([]string(nil), _networkmanagerSubnetOfMatches...)
	}
	if len(_networkmanagerSupernetOfMatches) > 0 {
		input.SupernetOfMatches = append([]string(nil), _networkmanagerSupernetOfMatches...)
	}
	if len(_networkmanagerTypes) > 0 {
		if err := assignInputField(input, "Types", _networkmanagerTypes); err != nil {
			log.Errorf("invalid --types: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetNetworkRoutes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the network telemetry of the specified global network.
func networkmanager_GetNetworkTelemetry(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.GetNetworkTelemetryInput{
		// GlobalNetworkId: *string, // Required
	}

	if len(_networkmanagerGlobalNetworkId) > 0 {
		input.GlobalNetworkId = aws.String(_networkmanagerGlobalNetworkId)
	}
	if len(_networkmanagerAccountId) > 0 {
		input.AccountId = aws.String(_networkmanagerAccountId)
	}
	if len(_networkmanagerAwsRegion) > 0 {
		input.AwsRegion = aws.String(_networkmanagerAwsRegion)
	}
	if len(_networkmanagerCoreNetworkId) > 0 {
		input.CoreNetworkId = aws.String(_networkmanagerCoreNetworkId)
	}
	if len(_networkmanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _networkmanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerNextToken) > 0 {
		input.NextToken = aws.String(_networkmanagerNextToken)
	}
	if len(_networkmanagerRegisteredGatewayArn) > 0 {
		input.RegisteredGatewayArn = aws.String(_networkmanagerRegisteredGatewayArn)
	}
	if len(_networkmanagerResourceArn) > 0 {
		input.ResourceArn = aws.String(_networkmanagerResourceArn)
	}
	if len(_networkmanagerResourceType) > 0 {
		input.ResourceType = aws.String(_networkmanagerResourceType)
	}

	if disablePaginator() {
		if resp, err := client.GetNetworkTelemetry(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*networkmanager.GetNetworkTelemetryOutput
	p := networkmanager.NewGetNetworkTelemetryPaginator(client, input)
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

// Returns information about a resource policy.
func networkmanager_GetResourcePolicy(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.GetResourcePolicyInput{
		// ResourceArn: *string, // Required
	}

	if len(_networkmanagerResourceArn) > 0 {
		input.ResourceArn = aws.String(_networkmanagerResourceArn)
	}

	if resp, err := client.GetResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the specified route analysis.
func networkmanager_GetRouteAnalysis(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.GetRouteAnalysisInput{
		// GlobalNetworkId: *string, // Required
		// RouteAnalysisId: *string, // Required
	}

	if len(_networkmanagerGlobalNetworkId) > 0 {
		input.GlobalNetworkId = aws.String(_networkmanagerGlobalNetworkId)
	}
	if len(_networkmanagerRouteAnalysisId) > 0 {
		input.RouteAnalysisId = aws.String(_networkmanagerRouteAnalysisId)
	}

	if resp, err := client.GetRouteAnalysis(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a site-to-site VPN attachment.
func networkmanager_GetSiteToSiteVpnAttachment(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.GetSiteToSiteVpnAttachmentInput{
		// AttachmentId: *string, // Required
	}

	if len(_networkmanagerAttachmentId) > 0 {
		input.AttachmentId = aws.String(_networkmanagerAttachmentId)
	}

	if resp, err := client.GetSiteToSiteVpnAttachment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about one or more of your sites in a global network.
func networkmanager_GetSites(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.GetSitesInput{
		// GlobalNetworkId: *string, // Required
	}

	if len(_networkmanagerGlobalNetworkId) > 0 {
		input.GlobalNetworkId = aws.String(_networkmanagerGlobalNetworkId)
	}
	if len(_networkmanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _networkmanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerNextToken) > 0 {
		input.NextToken = aws.String(_networkmanagerNextToken)
	}
	if len(_networkmanagerSiteIds) > 0 {
		input.SiteIds = append([]string(nil), _networkmanagerSiteIds...)
	}

	if disablePaginator() {
		if resp, err := client.GetSites(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*networkmanager.GetSitesOutput
	p := networkmanager.NewGetSitesPaginator(client, input)
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

// Gets information about one or more of your transit gateway Connect peer
// associations in a global network.
func networkmanager_GetTransitGatewayConnectPeerAssociations(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.GetTransitGatewayConnectPeerAssociationsInput{
		// GlobalNetworkId: *string, // Required
	}

	if len(_networkmanagerGlobalNetworkId) > 0 {
		input.GlobalNetworkId = aws.String(_networkmanagerGlobalNetworkId)
	}
	if len(_networkmanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _networkmanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerNextToken) > 0 {
		input.NextToken = aws.String(_networkmanagerNextToken)
	}
	if len(_networkmanagerTransitGatewayConnectPeerArns) > 0 {
		input.TransitGatewayConnectPeerArns = append([]string(nil), _networkmanagerTransitGatewayConnectPeerArns...)
	}

	if disablePaginator() {
		if resp, err := client.GetTransitGatewayConnectPeerAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*networkmanager.GetTransitGatewayConnectPeerAssociationsOutput
	p := networkmanager.NewGetTransitGatewayConnectPeerAssociationsPaginator(client, input)
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

// Returns information about a transit gateway peer.
func networkmanager_GetTransitGatewayPeering(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.GetTransitGatewayPeeringInput{
		// PeeringId: *string, // Required
	}

	if len(_networkmanagerPeeringId) > 0 {
		input.PeeringId = aws.String(_networkmanagerPeeringId)
	}

	if resp, err := client.GetTransitGatewayPeering(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the transit gateway registrations in a specified global
// network.
func networkmanager_GetTransitGatewayRegistrations(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.GetTransitGatewayRegistrationsInput{
		// GlobalNetworkId: *string, // Required
	}

	if len(_networkmanagerGlobalNetworkId) > 0 {
		input.GlobalNetworkId = aws.String(_networkmanagerGlobalNetworkId)
	}
	if len(_networkmanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _networkmanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerNextToken) > 0 {
		input.NextToken = aws.String(_networkmanagerNextToken)
	}
	if len(_networkmanagerTransitGatewayArns) > 0 {
		input.TransitGatewayArns = append([]string(nil), _networkmanagerTransitGatewayArns...)
	}

	if disablePaginator() {
		if resp, err := client.GetTransitGatewayRegistrations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*networkmanager.GetTransitGatewayRegistrationsOutput
	p := networkmanager.NewGetTransitGatewayRegistrationsPaginator(client, input)
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

// Returns information about a transit gateway route table attachment.
func networkmanager_GetTransitGatewayRouteTableAttachment(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.GetTransitGatewayRouteTableAttachmentInput{
		// AttachmentId: *string, // Required
	}

	if len(_networkmanagerAttachmentId) > 0 {
		input.AttachmentId = aws.String(_networkmanagerAttachmentId)
	}

	if resp, err := client.GetTransitGatewayRouteTableAttachment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a VPC attachment.
func networkmanager_GetVpcAttachment(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.GetVpcAttachmentInput{
		// AttachmentId: *string, // Required
	}

	if len(_networkmanagerAttachmentId) > 0 {
		input.AttachmentId = aws.String(_networkmanagerAttachmentId)
	}

	if resp, err := client.GetVpcAttachment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the routing policy associations for attachments in a core network.
func networkmanager_ListAttachmentRoutingPolicyAssociations(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.ListAttachmentRoutingPolicyAssociationsInput{
		// CoreNetworkId: *string, // Required
	}

	if len(_networkmanagerCoreNetworkId) > 0 {
		input.CoreNetworkId = aws.String(_networkmanagerCoreNetworkId)
	}
	if len(_networkmanagerAttachmentId) > 0 {
		input.AttachmentId = aws.String(_networkmanagerAttachmentId)
	}
	if len(_networkmanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _networkmanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerNextToken) > 0 {
		input.NextToken = aws.String(_networkmanagerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAttachmentRoutingPolicyAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*networkmanager.ListAttachmentRoutingPolicyAssociationsOutput
	p := networkmanager.NewListAttachmentRoutingPolicyAssociationsPaginator(client, input)
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

// Returns a list of core network attachments.
func networkmanager_ListAttachments(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.ListAttachmentsInput{}

	if len(_networkmanagerAttachmentType) > 0 {
		if err := assignInputField(input, "AttachmentType", _networkmanagerAttachmentType); err != nil {
			log.Errorf("invalid --attachment-type: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerCoreNetworkId) > 0 {
		input.CoreNetworkId = aws.String(_networkmanagerCoreNetworkId)
	}
	if len(_networkmanagerEdgeLocation) > 0 {
		input.EdgeLocation = aws.String(_networkmanagerEdgeLocation)
	}
	if len(_networkmanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _networkmanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerNextToken) > 0 {
		input.NextToken = aws.String(_networkmanagerNextToken)
	}
	if len(_networkmanagerState) > 0 {
		if err := assignInputField(input, "State", _networkmanagerState); err != nil {
			log.Errorf("invalid --state: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListAttachments(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*networkmanager.ListAttachmentsOutput
	p := networkmanager.NewListAttachmentsPaginator(client, input)
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

// Returns a list of core network Connect peers.
func networkmanager_ListConnectPeers(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.ListConnectPeersInput{}

	if len(_networkmanagerConnectAttachmentId) > 0 {
		input.ConnectAttachmentId = aws.String(_networkmanagerConnectAttachmentId)
	}
	if len(_networkmanagerCoreNetworkId) > 0 {
		input.CoreNetworkId = aws.String(_networkmanagerCoreNetworkId)
	}
	if len(_networkmanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _networkmanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerNextToken) > 0 {
		input.NextToken = aws.String(_networkmanagerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListConnectPeers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*networkmanager.ListConnectPeersOutput
	p := networkmanager.NewListConnectPeersPaginator(client, input)
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

// Returns a list of core network policy versions.
func networkmanager_ListCoreNetworkPolicyVersions(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.ListCoreNetworkPolicyVersionsInput{
		// CoreNetworkId: *string, // Required
	}

	if len(_networkmanagerCoreNetworkId) > 0 {
		input.CoreNetworkId = aws.String(_networkmanagerCoreNetworkId)
	}
	if len(_networkmanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _networkmanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerNextToken) > 0 {
		input.NextToken = aws.String(_networkmanagerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCoreNetworkPolicyVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*networkmanager.ListCoreNetworkPolicyVersionsOutput
	p := networkmanager.NewListCoreNetworkPolicyVersionsPaginator(client, input)
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

// Lists the prefix list associations for a core network.
func networkmanager_ListCoreNetworkPrefixListAssociations(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.ListCoreNetworkPrefixListAssociationsInput{
		// CoreNetworkId: *string, // Required
	}

	if len(_networkmanagerCoreNetworkId) > 0 {
		input.CoreNetworkId = aws.String(_networkmanagerCoreNetworkId)
	}
	if len(_networkmanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _networkmanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerNextToken) > 0 {
		input.NextToken = aws.String(_networkmanagerNextToken)
	}
	if len(_networkmanagerPrefixListArn) > 0 {
		input.PrefixListArn = aws.String(_networkmanagerPrefixListArn)
	}

	if disablePaginator() {
		if resp, err := client.ListCoreNetworkPrefixListAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*networkmanager.ListCoreNetworkPrefixListAssociationsOutput
	p := networkmanager.NewListCoreNetworkPrefixListAssociationsPaginator(client, input)
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

// Lists routing information for a core network, including routes and their
// attributes.
func networkmanager_ListCoreNetworkRoutingInformation(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.ListCoreNetworkRoutingInformationInput{
		// CoreNetworkId: *string, // Required
		// EdgeLocation: *string, // Required
		// SegmentName: *string, // Required
	}

	if len(_networkmanagerCoreNetworkId) > 0 {
		input.CoreNetworkId = aws.String(_networkmanagerCoreNetworkId)
	}
	if len(_networkmanagerEdgeLocation) > 0 {
		input.EdgeLocation = aws.String(_networkmanagerEdgeLocation)
	}
	if len(_networkmanagerSegmentName) > 0 {
		input.SegmentName = aws.String(_networkmanagerSegmentName)
	}
	if len(_networkmanagerCommunityMatches) > 0 {
		input.CommunityMatches = append([]string(nil), _networkmanagerCommunityMatches...)
	}
	if len(_networkmanagerExactAsPathMatches) > 0 {
		input.ExactAsPathMatches = append([]string(nil), _networkmanagerExactAsPathMatches...)
	}
	if len(_networkmanagerLocalPreferenceMatches) > 0 {
		input.LocalPreferenceMatches = append([]string(nil), _networkmanagerLocalPreferenceMatches...)
	}
	if len(_networkmanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _networkmanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerMedMatches) > 0 {
		input.MedMatches = append([]string(nil), _networkmanagerMedMatches...)
	}
	if len(_networkmanagerNextHopFilters) > 0 {
		if err := assignInputField(input, "NextHopFilters", _networkmanagerNextHopFilters); err != nil {
			log.Errorf("invalid --next-hop-filters: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerNextToken) > 0 {
		input.NextToken = aws.String(_networkmanagerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCoreNetworkRoutingInformation(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*networkmanager.ListCoreNetworkRoutingInformationOutput
	p := networkmanager.NewListCoreNetworkRoutingInformationPaginator(client, input)
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

// Returns a list of owned and shared core networks.
func networkmanager_ListCoreNetworks(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.ListCoreNetworksInput{}

	if len(_networkmanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _networkmanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerNextToken) > 0 {
		input.NextToken = aws.String(_networkmanagerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCoreNetworks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*networkmanager.ListCoreNetworksOutput
	p := networkmanager.NewListCoreNetworksPaginator(client, input)
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

// Gets the status of the Service Linked Role (SLR) deployment for the accounts in
// a given Amazon Web Services Organization.
func networkmanager_ListOrganizationServiceAccessStatus(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.ListOrganizationServiceAccessStatusInput{}

	if len(_networkmanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _networkmanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerNextToken) > 0 {
		input.NextToken = aws.String(_networkmanagerNextToken)
	}

	if resp, err := client.ListOrganizationServiceAccessStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the peerings for a core network.
func networkmanager_ListPeerings(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.ListPeeringsInput{}

	if len(_networkmanagerCoreNetworkId) > 0 {
		input.CoreNetworkId = aws.String(_networkmanagerCoreNetworkId)
	}
	if len(_networkmanagerEdgeLocation) > 0 {
		input.EdgeLocation = aws.String(_networkmanagerEdgeLocation)
	}
	if len(_networkmanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _networkmanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerNextToken) > 0 {
		input.NextToken = aws.String(_networkmanagerNextToken)
	}
	if len(_networkmanagerPeeringType) > 0 {
		if err := assignInputField(input, "PeeringType", _networkmanagerPeeringType); err != nil {
			log.Errorf("invalid --peering-type: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerState) > 0 {
		if err := assignInputField(input, "State", _networkmanagerState); err != nil {
			log.Errorf("invalid --state: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListPeerings(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*networkmanager.ListPeeringsOutput
	p := networkmanager.NewListPeeringsPaginator(client, input)
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
func networkmanager_ListTagsForResource(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_networkmanagerResourceArn) > 0 {
		input.ResourceArn = aws.String(_networkmanagerResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Applies a routing policy label to an attachment for traffic routing decisions.
func networkmanager_PutAttachmentRoutingPolicyLabel(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.PutAttachmentRoutingPolicyLabelInput{
		// AttachmentId: *string, // Required
		// CoreNetworkId: *string, // Required
		// RoutingPolicyLabel: *string, // Required
	}

	if len(_networkmanagerAttachmentId) > 0 {
		input.AttachmentId = aws.String(_networkmanagerAttachmentId)
	}
	if len(_networkmanagerCoreNetworkId) > 0 {
		input.CoreNetworkId = aws.String(_networkmanagerCoreNetworkId)
	}
	if len(_networkmanagerRoutingPolicyLabel) > 0 {
		input.RoutingPolicyLabel = aws.String(_networkmanagerRoutingPolicyLabel)
	}
	if len(_networkmanagerClientToken) > 0 {
		input.ClientToken = aws.String(_networkmanagerClientToken)
	}

	if resp, err := client.PutAttachmentRoutingPolicyLabel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new, immutable version of a core network policy. A subsequent change
// set is created showing the differences between the LIVE policy and the submitted
// policy.
func networkmanager_PutCoreNetworkPolicy(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.PutCoreNetworkPolicyInput{
		// CoreNetworkId: *string, // Required
		// PolicyDocument: *string, // Required
	}

	if len(_networkmanagerCoreNetworkId) > 0 {
		input.CoreNetworkId = aws.String(_networkmanagerCoreNetworkId)
	}
	if len(_networkmanagerPolicyDocument) > 0 {
		input.PolicyDocument = aws.String(_networkmanagerPolicyDocument)
	}
	if len(_networkmanagerClientToken) > 0 {
		input.ClientToken = aws.String(_networkmanagerClientToken)
	}
	if len(_networkmanagerDescription) > 0 {
		input.Description = aws.String(_networkmanagerDescription)
	}
	if len(_networkmanagerLatestVersionId) > 0 {
		if err := assignInputField(input, "LatestVersionId", _networkmanagerLatestVersionId); err != nil {
			log.Errorf("invalid --latest-version-id: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutCoreNetworkPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates a resource policy.
func networkmanager_PutResourcePolicy(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.PutResourcePolicyInput{
		// PolicyDocument: *string, // Required
		// ResourceArn: *string, // Required
	}

	if len(_networkmanagerPolicyDocument) > 0 {
		input.PolicyDocument = aws.String(_networkmanagerPolicyDocument)
	}
	if len(_networkmanagerResourceArn) > 0 {
		input.ResourceArn = aws.String(_networkmanagerResourceArn)
	}

	if resp, err := client.PutResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Registers a transit gateway in your global network. Not all Regions support
// transit gateways for global networks. For a list of the supported Regions, see [Region Availability]
// in the Amazon Web Services Transit Gateways for Global Networks User Guide. The
// transit gateway can be in any of the supported Amazon Web Services Regions, but
// it must be owned by the same Amazon Web Services account that owns the global
// network. You cannot register a transit gateway in more than one global network.
//
// [Region Availability]: https://docs.aws.amazon.com/network-manager/latest/tgwnm/what-are-global-networks.html#nm-available-regions
func networkmanager_RegisterTransitGateway(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.RegisterTransitGatewayInput{
		// GlobalNetworkId: *string, // Required
		// TransitGatewayArn: *string, // Required
	}

	if len(_networkmanagerGlobalNetworkId) > 0 {
		input.GlobalNetworkId = aws.String(_networkmanagerGlobalNetworkId)
	}
	if len(_networkmanagerTransitGatewayArn) > 0 {
		input.TransitGatewayArn = aws.String(_networkmanagerTransitGatewayArn)
	}

	if resp, err := client.RegisterTransitGateway(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Rejects a core network attachment request.
func networkmanager_RejectAttachment(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.RejectAttachmentInput{
		// AttachmentId: *string, // Required
	}

	if len(_networkmanagerAttachmentId) > 0 {
		input.AttachmentId = aws.String(_networkmanagerAttachmentId)
	}

	if resp, err := client.RejectAttachment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a routing policy label from an attachment.
func networkmanager_RemoveAttachmentRoutingPolicyLabel(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.RemoveAttachmentRoutingPolicyLabelInput{
		// AttachmentId: *string, // Required
		// CoreNetworkId: *string, // Required
	}

	if len(_networkmanagerAttachmentId) > 0 {
		input.AttachmentId = aws.String(_networkmanagerAttachmentId)
	}
	if len(_networkmanagerCoreNetworkId) > 0 {
		input.CoreNetworkId = aws.String(_networkmanagerCoreNetworkId)
	}

	if resp, err := client.RemoveAttachmentRoutingPolicyLabel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Restores a previous policy version as a new, immutable version of a core
// network policy. A subsequent change set is created showing the differences
// between the LIVE policy and restored policy.
func networkmanager_RestoreCoreNetworkPolicyVersion(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.RestoreCoreNetworkPolicyVersionInput{
		// CoreNetworkId: *string, // Required
		// PolicyVersionId: *int32, // Required
	}

	if len(_networkmanagerCoreNetworkId) > 0 {
		input.CoreNetworkId = aws.String(_networkmanagerCoreNetworkId)
	}
	if len(_networkmanagerPolicyVersionId) > 0 {
		if err := assignInputField(input, "PolicyVersionId", _networkmanagerPolicyVersionId); err != nil {
			log.Errorf("invalid --policy-version-id: %s", err.Error())
			return
		}
	}

	if resp, err := client.RestoreCoreNetworkPolicyVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables the Network Manager service for an Amazon Web Services Organization.
// This can only be called by a management account within the organization.
func networkmanager_StartOrganizationServiceAccessUpdate(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.StartOrganizationServiceAccessUpdateInput{
		// Action: *string, // Required
	}

	if len(_networkmanagerAction) > 0 {
		input.Action = aws.String(_networkmanagerAction)
	}

	if resp, err := client.StartOrganizationServiceAccessUpdate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts analyzing the routing path between the specified source and destination.
// For more information, see [Route Analyzer].
//
// [Route Analyzer]: https://docs.aws.amazon.com/vpc/latest/tgw/route-analyzer.html
func networkmanager_StartRouteAnalysis(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.StartRouteAnalysisInput{
		// Destination: *types.RouteAnalysisEndpointOptionsSpecification, // Required
		// GlobalNetworkId: *string, // Required
		// Source: *types.RouteAnalysisEndpointOptionsSpecification, // Required
	}

	if len(_networkmanagerDestination) > 0 {
		if err := assignInputField(input, "Destination", _networkmanagerDestination); err != nil {
			log.Errorf("invalid --destination: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerGlobalNetworkId) > 0 {
		input.GlobalNetworkId = aws.String(_networkmanagerGlobalNetworkId)
	}
	if len(_networkmanagerSource) > 0 {
		if err := assignInputField(input, "Source", _networkmanagerSource); err != nil {
			log.Errorf("invalid --source: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerIncludeReturnPath) > 0 {
		if err := assignInputField(input, "IncludeReturnPath", _networkmanagerIncludeReturnPath); err != nil {
			log.Errorf("invalid --include-return-path: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerUseMiddleboxes) > 0 {
		if err := assignInputField(input, "UseMiddleboxes", _networkmanagerUseMiddleboxes); err != nil {
			log.Errorf("invalid --use-middleboxes: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartRouteAnalysis(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Tags a specified resource.
func networkmanager_TagResource(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_networkmanagerResourceArn) > 0 {
		input.ResourceArn = aws.String(_networkmanagerResourceArn)
	}
	if len(_networkmanagerTags) > 0 {
		if err := assignInputField(input, "Tags", _networkmanagerTags); err != nil {
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
func networkmanager_UntagResource(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_networkmanagerResourceArn) > 0 {
		input.ResourceArn = aws.String(_networkmanagerResourceArn)
	}
	if len(_networkmanagerTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _networkmanagerTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the information for an existing connection. To remove information for
// any of the parameters, specify an empty string.
func networkmanager_UpdateConnection(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.UpdateConnectionInput{
		// ConnectionId: *string, // Required
		// GlobalNetworkId: *string, // Required
	}

	if len(_networkmanagerConnectionId) > 0 {
		input.ConnectionId = aws.String(_networkmanagerConnectionId)
	}
	if len(_networkmanagerGlobalNetworkId) > 0 {
		input.GlobalNetworkId = aws.String(_networkmanagerGlobalNetworkId)
	}
	if len(_networkmanagerConnectedLinkId) > 0 {
		input.ConnectedLinkId = aws.String(_networkmanagerConnectedLinkId)
	}
	if len(_networkmanagerDescription) > 0 {
		input.Description = aws.String(_networkmanagerDescription)
	}
	if len(_networkmanagerLinkId) > 0 {
		input.LinkId = aws.String(_networkmanagerLinkId)
	}

	if resp, err := client.UpdateConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the description of a core network.
func networkmanager_UpdateCoreNetwork(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.UpdateCoreNetworkInput{
		// CoreNetworkId: *string, // Required
	}

	if len(_networkmanagerCoreNetworkId) > 0 {
		input.CoreNetworkId = aws.String(_networkmanagerCoreNetworkId)
	}
	if len(_networkmanagerDescription) > 0 {
		input.Description = aws.String(_networkmanagerDescription)
	}

	if resp, err := client.UpdateCoreNetwork(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the details for an existing device. To remove information for any of
// the parameters, specify an empty string.
func networkmanager_UpdateDevice(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.UpdateDeviceInput{
		// DeviceId: *string, // Required
		// GlobalNetworkId: *string, // Required
	}

	if len(_networkmanagerDeviceId) > 0 {
		input.DeviceId = aws.String(_networkmanagerDeviceId)
	}
	if len(_networkmanagerGlobalNetworkId) > 0 {
		input.GlobalNetworkId = aws.String(_networkmanagerGlobalNetworkId)
	}
	if len(_networkmanagerAWSLocation) > 0 {
		if err := assignInputField(input, "AWSLocation", _networkmanagerAWSLocation); err != nil {
			log.Errorf("invalid --aws-location: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerDescription) > 0 {
		input.Description = aws.String(_networkmanagerDescription)
	}
	if len(_networkmanagerLocation) > 0 {
		if err := assignInputField(input, "Location", _networkmanagerLocation); err != nil {
			log.Errorf("invalid --location: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerModel) > 0 {
		input.Model = aws.String(_networkmanagerModel)
	}
	if len(_networkmanagerSerialNumber) > 0 {
		input.SerialNumber = aws.String(_networkmanagerSerialNumber)
	}
	if len(_networkmanagerSiteId) > 0 {
		input.SiteId = aws.String(_networkmanagerSiteId)
	}
	if len(_networkmanagerType) > 0 {
		input.Type = aws.String(_networkmanagerType)
	}
	if len(_networkmanagerVendor) > 0 {
		input.Vendor = aws.String(_networkmanagerVendor)
	}

	if resp, err := client.UpdateDevice(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the edge locations associated with an Amazon Web Services Direct
// Connect gateway attachment.
func networkmanager_UpdateDirectConnectGatewayAttachment(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.UpdateDirectConnectGatewayAttachmentInput{
		// AttachmentId: *string, // Required
	}

	if len(_networkmanagerAttachmentId) > 0 {
		input.AttachmentId = aws.String(_networkmanagerAttachmentId)
	}
	if len(_networkmanagerEdgeLocations) > 0 {
		input.EdgeLocations = append([]string(nil), _networkmanagerEdgeLocations...)
	}

	if resp, err := client.UpdateDirectConnectGatewayAttachment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing global network. To remove information for any of the
// parameters, specify an empty string.
func networkmanager_UpdateGlobalNetwork(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.UpdateGlobalNetworkInput{
		// GlobalNetworkId: *string, // Required
	}

	if len(_networkmanagerGlobalNetworkId) > 0 {
		input.GlobalNetworkId = aws.String(_networkmanagerGlobalNetworkId)
	}
	if len(_networkmanagerDescription) > 0 {
		input.Description = aws.String(_networkmanagerDescription)
	}

	if resp, err := client.UpdateGlobalNetwork(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the details for an existing link. To remove information for any of the
// parameters, specify an empty string.
func networkmanager_UpdateLink(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.UpdateLinkInput{
		// GlobalNetworkId: *string, // Required
		// LinkId: *string, // Required
	}

	if len(_networkmanagerGlobalNetworkId) > 0 {
		input.GlobalNetworkId = aws.String(_networkmanagerGlobalNetworkId)
	}
	if len(_networkmanagerLinkId) > 0 {
		input.LinkId = aws.String(_networkmanagerLinkId)
	}
	if len(_networkmanagerBandwidth) > 0 {
		if err := assignInputField(input, "Bandwidth", _networkmanagerBandwidth); err != nil {
			log.Errorf("invalid --bandwidth: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerDescription) > 0 {
		input.Description = aws.String(_networkmanagerDescription)
	}
	if len(_networkmanagerProvider) > 0 {
		input.Provider = aws.String(_networkmanagerProvider)
	}
	if len(_networkmanagerType) > 0 {
		input.Type = aws.String(_networkmanagerType)
	}

	if resp, err := client.UpdateLink(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the resource metadata for the specified global network.
func networkmanager_UpdateNetworkResourceMetadata(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.UpdateNetworkResourceMetadataInput{
		// GlobalNetworkId: *string, // Required
		// Metadata: map[string]string, // Required
		// ResourceArn: *string, // Required
	}

	if len(_networkmanagerGlobalNetworkId) > 0 {
		input.GlobalNetworkId = aws.String(_networkmanagerGlobalNetworkId)
	}
	if len(_networkmanagerMetadata) > 0 {
		if err := assignInputField(input, "Metadata", _networkmanagerMetadata); err != nil {
			log.Errorf("invalid --metadata: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerResourceArn) > 0 {
		input.ResourceArn = aws.String(_networkmanagerResourceArn)
	}

	if resp, err := client.UpdateNetworkResourceMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the information for an existing site. To remove information for any of
// the parameters, specify an empty string.
func networkmanager_UpdateSite(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.UpdateSiteInput{
		// GlobalNetworkId: *string, // Required
		// SiteId: *string, // Required
	}

	if len(_networkmanagerGlobalNetworkId) > 0 {
		input.GlobalNetworkId = aws.String(_networkmanagerGlobalNetworkId)
	}
	if len(_networkmanagerSiteId) > 0 {
		input.SiteId = aws.String(_networkmanagerSiteId)
	}
	if len(_networkmanagerDescription) > 0 {
		input.Description = aws.String(_networkmanagerDescription)
	}
	if len(_networkmanagerLocation) > 0 {
		if err := assignInputField(input, "Location", _networkmanagerLocation); err != nil {
			log.Errorf("invalid --location: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSite(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a VPC attachment.
func networkmanager_UpdateVpcAttachment(cfg aws.Config, client *networkmanager.Client) {
	input := &networkmanager.UpdateVpcAttachmentInput{
		// AttachmentId: *string, // Required
	}

	if len(_networkmanagerAttachmentId) > 0 {
		input.AttachmentId = aws.String(_networkmanagerAttachmentId)
	}
	if len(_networkmanagerAddSubnetArns) > 0 {
		input.AddSubnetArns = append([]string(nil), _networkmanagerAddSubnetArns...)
	}
	if len(_networkmanagerOptions) > 0 {
		if err := assignInputField(input, "Options", _networkmanagerOptions); err != nil {
			log.Errorf("invalid --options: %s", err.Error())
			return
		}
	}
	if len(_networkmanagerRemoveSubnetArns) > 0 {
		input.RemoveSubnetArns = append([]string(nil), _networkmanagerRemoveSubnetArns...)
	}

	if resp, err := client.UpdateVpcAttachment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_networkmanagerCmd)
	_networkmanagerCmd.Flags().SortFlags = false

	_networkmanagerCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_networkmanagerCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_networkmanagerCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerAccountId, "account-id", "", "", "Account ID")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerAction, "action", "", "", "Action")
	_networkmanagerCmd.Flags().StringSliceVarP(&_networkmanagerAddSubnetArns, "add-subnet-arns", "", nil, "Add Subnet Arns")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerAlias, "alias", "", "", "Alias")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerAttachmentId, "attachment-id", "", "", "Attachment ID")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerAttachmentType, "attachment-type", "", "", "Attachment Type")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerAWSLocation, "aws-location", "", "", "AWS Location")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerAwsRegion, "aws-region", "", "", "AWS Region")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerBandwidth, "bandwidth", "", "", "Bandwidth")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerBgpOptions, "bgp-options", "", "", "Bgp Options")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerClientToken, "client-token", "", "", "Client Token")
	_networkmanagerCmd.Flags().StringSliceVarP(&_networkmanagerCommunityMatches, "community-matches", "", nil, "Community Matches")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerConnectAttachmentId, "connect-attachment-id", "", "", "Connect Attachment ID")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerConnectPeerId, "connect-peer-id", "", "", "Connect Peer ID")
	_networkmanagerCmd.Flags().StringSliceVarP(&_networkmanagerConnectPeerIds, "connect-peer-ids", "", nil, "Connect Peer Ids")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerConnectedDeviceId, "connected-device-id", "", "", "Connected Device ID")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerConnectedLinkId, "connected-link-id", "", "", "Connected Link ID")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerConnectionId, "connection-id", "", "", "Connection ID")
	_networkmanagerCmd.Flags().StringSliceVarP(&_networkmanagerConnectionIds, "connection-ids", "", nil, "Connection Ids")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerCoreNetworkAddress, "core-network-address", "", "", "Core Network Address")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerCoreNetworkId, "core-network-id", "", "", "Core Network ID")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerCustomerGatewayArn, "customer-gateway-arn", "", "", "Customer Gateway ARN")
	_networkmanagerCmd.Flags().StringSliceVarP(&_networkmanagerCustomerGatewayArns, "customer-gateway-arns", "", nil, "Customer Gateway Arns")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerDescription, "description", "", "", "Description")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerDestination, "destination", "", "", "Destination")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerDestinationFilters, "destination-filters", "", "", "Destination Filters")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerDeviceId, "device-id", "", "", "Device ID")
	_networkmanagerCmd.Flags().StringSliceVarP(&_networkmanagerDeviceIds, "device-ids", "", nil, "Device Ids")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerDirectConnectGatewayArn, "direct-connect-gateway-arn", "", "", "Direct Connect Gateway ARN")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerEdgeLocation, "edge-location", "", "", "Edge Location")
	_networkmanagerCmd.Flags().StringSliceVarP(&_networkmanagerEdgeLocations, "edge-locations", "", nil, "Edge Locations")
	_networkmanagerCmd.Flags().StringSliceVarP(&_networkmanagerExactAsPathMatches, "exact-as-path-matches", "", nil, "Exact As Path Matches")
	_networkmanagerCmd.Flags().StringSliceVarP(&_networkmanagerExactCidrMatches, "exact-cidr-matches", "", nil, "Exact CIDR Matches")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerGlobalNetworkId, "global-network-id", "", "", "Global Network ID")
	_networkmanagerCmd.Flags().StringSliceVarP(&_networkmanagerGlobalNetworkIds, "global-network-ids", "", nil, "Global Network Ids")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerIncludeReturnPath, "include-return-path", "", "", "Include Return Path")
	_networkmanagerCmd.Flags().StringSliceVarP(&_networkmanagerInsideCidrBlocks, "inside-cidr-blocks", "", nil, "Inside CIDR Blocks")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerLatestVersionId, "latest-version-id", "", "", "Latest Version ID")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerLinkId, "link-id", "", "", "Link ID")
	_networkmanagerCmd.Flags().StringSliceVarP(&_networkmanagerLinkIds, "link-ids", "", nil, "Link Ids")
	_networkmanagerCmd.Flags().StringSliceVarP(&_networkmanagerLocalPreferenceMatches, "local-preference-matches", "", nil, "Local Preference Matches")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerLocation, "location", "", "", "Location")
	_networkmanagerCmd.Flags().StringSliceVarP(&_networkmanagerLongestPrefixMatches, "longest-prefix-matches", "", nil, "Longest Prefix Matches")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerMaxResults, "max-results", "", "", "Max Results")
	_networkmanagerCmd.Flags().StringSliceVarP(&_networkmanagerMedMatches, "med-matches", "", nil, "Med Matches")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerMetadata, "metadata", "", "", "Metadata")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerModel, "model", "", "", "Model")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerNextHopFilters, "next-hop-filters", "", "", "Next Hop Filters")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerNextToken, "next-token", "", "", "Next Token")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerOptions, "options", "", "", "Options")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerPeerAddress, "peer-address", "", "", "Peer Address")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerPeeringId, "peering-id", "", "", "Peering ID")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerPeeringType, "peering-type", "", "", "Peering Type")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerPolicyDocument, "policy-document", "", "", "Policy Document")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerPolicyVersionId, "policy-version-id", "", "", "Policy Version ID")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerPrefixListAlias, "prefix-list-alias", "", "", "Prefix List Alias")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerPrefixListArn, "prefix-list-arn", "", "", "Prefix List ARN")
	_networkmanagerCmd.Flags().StringSliceVarP(&_networkmanagerPrefixListIds, "prefix-list-ids", "", nil, "Prefix List Ids")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerProvider, "provider", "", "", "Provider")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerRegisteredGatewayArn, "registered-gateway-arn", "", "", "Registered Gateway ARN")
	_networkmanagerCmd.Flags().StringSliceVarP(&_networkmanagerRemoveSubnetArns, "remove-subnet-arns", "", nil, "Remove Subnet Arns")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerResourceArn, "resource-arn", "", "", "Resource ARN")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerResourceType, "resource-type", "", "", "Resource Type")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerRouteAnalysisId, "route-analysis-id", "", "", "Route Analysis ID")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerRouteTableIdentifier, "route-table-identifier", "", "", "Route Table Identifier")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerRoutingPolicyLabel, "routing-policy-label", "", "", "Routing Policy Label")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerSegmentName, "segment-name", "", "", "Segment Name")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerSerialNumber, "serial-number", "", "", "Serial Number")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerSiteId, "site-id", "", "", "Site ID")
	_networkmanagerCmd.Flags().StringSliceVarP(&_networkmanagerSiteIds, "site-ids", "", nil, "Site Ids")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerSource, "source", "", "", "Source")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerState, "state", "", "", "State")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerStates, "states", "", "", "States")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerSubnetArn, "subnet-arn", "", "", "Subnet ARN")
	_networkmanagerCmd.Flags().StringSliceVarP(&_networkmanagerSubnetArns, "subnet-arns", "", nil, "Subnet Arns")
	_networkmanagerCmd.Flags().StringSliceVarP(&_networkmanagerSubnetOfMatches, "subnet-of-matches", "", nil, "Subnet Of Matches")
	_networkmanagerCmd.Flags().StringSliceVarP(&_networkmanagerSupernetOfMatches, "supernet-of-matches", "", nil, "Supernet Of Matches")
	_networkmanagerCmd.Flags().StringSliceVarP(&_networkmanagerTagKeys, "tag-keys", "", nil, "Tag Keys")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerTags, "tags", "", "", "Tags")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerTransitGatewayArn, "transit-gateway-arn", "", "", "Transit Gateway ARN")
	_networkmanagerCmd.Flags().StringSliceVarP(&_networkmanagerTransitGatewayArns, "transit-gateway-arns", "", nil, "Transit Gateway Arns")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerTransitGatewayConnectPeerArn, "transit-gateway-connect-peer-arn", "", "", "Transit Gateway Connect Peer ARN")
	_networkmanagerCmd.Flags().StringSliceVarP(&_networkmanagerTransitGatewayConnectPeerArns, "transit-gateway-connect-peer-arns", "", nil, "Transit Gateway Connect Peer Arns")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerTransitGatewayRouteTableArn, "transit-gateway-route-table-arn", "", "", "Transit Gateway Route Table ARN")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerTransportAttachmentId, "transport-attachment-id", "", "", "Transport Attachment ID")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerType, "type", "", "", "Type")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerTypes, "types", "", "", "Types")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerUseMiddleboxes, "use-middleboxes", "", "", "Use Middleboxes")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerVendor, "vendor", "", "", "Vendor")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerVpcArn, "vpc-arn", "", "", "VPC ARN")
	_networkmanagerCmd.Flags().StringVarP(&_networkmanagerVpnConnectionArn, "vpn-connection-arn", "", "", "VPN Connection ARN")

	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerAcceptAttachment, "accept-attachment", "", false, "Accept Attachment")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerAssociateConnectPeer, "associate-connect-peer", "", false, "Associate Connect Peer")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerAssociateCustomerGateway, "associate-customer-gateway", "", false, "Associate Customer Gateway")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerAssociateLink, "associate-link", "", false, "Associate Link")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerAssociateTransitGatewayConnectPeer, "associate-transit-gateway-connect-peer", "", false, "Associate Transit Gateway Connect Peer")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerCreateConnectAttachment, "create-connect-attachment", "", false, "Create Connect Attachment")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerCreateConnectPeer, "create-connect-peer", "", false, "Create Connect Peer")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerCreateConnection, "create-connection", "", false, "Create Connection")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerCreateCoreNetwork, "create-core-network", "", false, "Create Core Network")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerCreateCoreNetworkPrefixListAssociation, "create-core-network-prefix-list-association", "", false, "Create Core Network Prefix List Association")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerCreateDevice, "create-device", "", false, "Create Device")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerCreateDirectConnectGatewayAttachment, "create-direct-connect-gateway-attachment", "", false, "Create Direct Connect Gateway Attachment")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerCreateGlobalNetwork, "create-global-network", "", false, "Create Global Network")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerCreateLink, "create-link", "", false, "Create Link")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerCreateSite, "create-site", "", false, "Create Site")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerCreateSiteToSiteVpnAttachment, "create-site-to-site-vpn-attachment", "", false, "Create Site To Site VPN Attachment")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerCreateTransitGatewayPeering, "create-transit-gateway-peering", "", false, "Create Transit Gateway Peering")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerCreateTransitGatewayRouteTableAttachment, "create-transit-gateway-route-table-attachment", "", false, "Create Transit Gateway Route Table Attachment")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerCreateVpcAttachment, "create-vpc-attachment", "", false, "Create VPC Attachment")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerDeleteAttachment, "delete-attachment", "", false, "Delete Attachment")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerDeleteConnectPeer, "delete-connect-peer", "", false, "Delete Connect Peer")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerDeleteConnection, "delete-connection", "", false, "Delete Connection")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerDeleteCoreNetwork, "delete-core-network", "", false, "Delete Core Network")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerDeleteCoreNetworkPolicyVersion, "delete-core-network-policy-version", "", false, "Delete Core Network Policy Version")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerDeleteCoreNetworkPrefixListAssociation, "delete-core-network-prefix-list-association", "", false, "Delete Core Network Prefix List Association")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerDeleteDevice, "delete-device", "", false, "Delete Device")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerDeleteGlobalNetwork, "delete-global-network", "", false, "Delete Global Network")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerDeleteLink, "delete-link", "", false, "Delete Link")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerDeletePeering, "delete-peering", "", false, "Delete Peering")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerDeleteResourcePolicy, "delete-resource-policy", "", false, "Delete Resource Policy")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerDeleteSite, "delete-site", "", false, "Delete Site")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerDeregisterTransitGateway, "deregister-transit-gateway", "", false, "Deregister Transit Gateway")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerDescribeGlobalNetworks, "describe-global-networks", "", false, "Describe Global Networks")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerDisassociateConnectPeer, "disassociate-connect-peer", "", false, "Disassociate Connect Peer")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerDisassociateCustomerGateway, "disassociate-customer-gateway", "", false, "Disassociate Customer Gateway")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerDisassociateLink, "disassociate-link", "", false, "Disassociate Link")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerDisassociateTransitGatewayConnectPeer, "disassociate-transit-gateway-connect-peer", "", false, "Disassociate Transit Gateway Connect Peer")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerExecuteCoreNetworkChangeSet, "execute-core-network-change-set", "", false, "Execute Core Network Change Set")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerGetConnectAttachment, "get-connect-attachment", "", false, "Get Connect Attachment")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerGetConnectPeer, "get-connect-peer", "", false, "Get Connect Peer")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerGetConnectPeerAssociations, "get-connect-peer-associations", "", false, "Get Connect Peer Associations")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerGetConnections, "get-connections", "", false, "Get Connections")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerGetCoreNetwork, "get-core-network", "", false, "Get Core Network")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerGetCoreNetworkChangeEvents, "get-core-network-change-events", "", false, "Get Core Network Change Events")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerGetCoreNetworkChangeSet, "get-core-network-change-set", "", false, "Get Core Network Change Set")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerGetCoreNetworkPolicy, "get-core-network-policy", "", false, "Get Core Network Policy")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerGetCustomerGatewayAssociations, "get-customer-gateway-associations", "", false, "Get Customer Gateway Associations")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerGetDevices, "get-devices", "", false, "Get Devices")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerGetDirectConnectGatewayAttachment, "get-direct-connect-gateway-attachment", "", false, "Get Direct Connect Gateway Attachment")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerGetLinkAssociations, "get-link-associations", "", false, "Get Link Associations")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerGetLinks, "get-links", "", false, "Get Links")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerGetNetworkResourceCounts, "get-network-resource-counts", "", false, "Get Network Resource Counts")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerGetNetworkResourceRelationships, "get-network-resource-relationships", "", false, "Get Network Resource Relationships")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerGetNetworkResources, "get-network-resources", "", false, "Get Network Resources")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerGetNetworkRoutes, "get-network-routes", "", false, "Get Network Routes")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerGetNetworkTelemetry, "get-network-telemetry", "", false, "Get Network Telemetry")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerGetResourcePolicy, "get-resource-policy", "", false, "Get Resource Policy")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerGetRouteAnalysis, "get-route-analysis", "", false, "Get Route Analysis")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerGetSiteToSiteVpnAttachment, "get-site-to-site-vpn-attachment", "", false, "Get Site To Site VPN Attachment")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerGetSites, "get-sites", "", false, "Get Sites")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerGetTransitGatewayConnectPeerAssociations, "get-transit-gateway-connect-peer-associations", "", false, "Get Transit Gateway Connect Peer Associations")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerGetTransitGatewayPeering, "get-transit-gateway-peering", "", false, "Get Transit Gateway Peering")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerGetTransitGatewayRegistrations, "get-transit-gateway-registrations", "", false, "Get Transit Gateway Registrations")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerGetTransitGatewayRouteTableAttachment, "get-transit-gateway-route-table-attachment", "", false, "Get Transit Gateway Route Table Attachment")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerGetVpcAttachment, "get-vpc-attachment", "", false, "Get VPC Attachment")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerListAttachmentRoutingPolicyAssociations, "list-attachment-routing-policy-associations", "", false, "List Attachment Routing Policy Associations")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerListAttachments, "list-attachments", "", false, "List Attachments")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerListConnectPeers, "list-connect-peers", "", false, "List Connect Peers")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerListCoreNetworkPolicyVersions, "list-core-network-policy-versions", "", false, "List Core Network Policy Versions")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerListCoreNetworkPrefixListAssociations, "list-core-network-prefix-list-associations", "", false, "List Core Network Prefix List Associations")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerListCoreNetworkRoutingInformation, "list-core-network-routing-information", "", false, "List Core Network Routing Information")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerListCoreNetworks, "list-core-networks", "", false, "List Core Networks")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerListOrganizationServiceAccessStatus, "list-organization-service-access-status", "", false, "List Organization Service Access Status")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerListPeerings, "list-peerings", "", false, "List Peerings")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerPutAttachmentRoutingPolicyLabel, "put-attachment-routing-policy-label", "", false, "Put Attachment Routing Policy Label")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerPutCoreNetworkPolicy, "put-core-network-policy", "", false, "Put Core Network Policy")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerPutResourcePolicy, "put-resource-policy", "", false, "Put Resource Policy")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerRegisterTransitGateway, "register-transit-gateway", "", false, "Register Transit Gateway")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerRejectAttachment, "reject-attachment", "", false, "Reject Attachment")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerRemoveAttachmentRoutingPolicyLabel, "remove-attachment-routing-policy-label", "", false, "Remove Attachment Routing Policy Label")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerRestoreCoreNetworkPolicyVersion, "restore-core-network-policy-version", "", false, "Restore Core Network Policy Version")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerStartOrganizationServiceAccessUpdate, "start-organization-service-access-update", "", false, "Start Organization Service Access Update")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerStartRouteAnalysis, "start-route-analysis", "", false, "Start Route Analysis")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerTagResource, "tag-resource", "", false, "Tag Resource")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerUntagResource, "untag-resource", "", false, "Untag Resource")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerUpdateConnection, "update-connection", "", false, "Update Connection")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerUpdateCoreNetwork, "update-core-network", "", false, "Update Core Network")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerUpdateDevice, "update-device", "", false, "Update Device")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerUpdateDirectConnectGatewayAttachment, "update-direct-connect-gateway-attachment", "", false, "Update Direct Connect Gateway Attachment")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerUpdateGlobalNetwork, "update-global-network", "", false, "Update Global Network")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerUpdateLink, "update-link", "", false, "Update Link")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerUpdateNetworkResourceMetadata, "update-network-resource-metadata", "", false, "Update Network Resource Metadata")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerUpdateSite, "update-site", "", false, "Update Site")
	_networkmanagerCmd.Flags().BoolVarP(&_networkmanagerUpdateVpcAttachment, "update-vpc-attachment", "", false, "Update VPC Attachment")

}
