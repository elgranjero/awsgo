package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// directconnectCmd represents the directconnect command
var _directconnectCmd = &cobra.Command{
	Use:   "directconnect",
	Short: "AWS directconnect CLI",
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
		client := directconnect.NewFromConfig(cfg)
		if _directconnectAcceptDirectConnectGatewayAssociationProposal {
			directconnect_AcceptDirectConnectGatewayAssociationProposal(cfg, client)
			return
		}
		if _directconnectAllocateConnectionOnInterconnect {
			directconnect_AllocateConnectionOnInterconnect(cfg, client)
			return
		}
		if _directconnectAllocateHostedConnection {
			directconnect_AllocateHostedConnection(cfg, client)
			return
		}
		if _directconnectAllocatePrivateVirtualInterface {
			directconnect_AllocatePrivateVirtualInterface(cfg, client)
			return
		}
		if _directconnectAllocatePublicVirtualInterface {
			directconnect_AllocatePublicVirtualInterface(cfg, client)
			return
		}
		if _directconnectAllocateTransitVirtualInterface {
			directconnect_AllocateTransitVirtualInterface(cfg, client)
			return
		}
		if _directconnectAssociateConnectionWithLag {
			directconnect_AssociateConnectionWithLag(cfg, client)
			return
		}
		if _directconnectAssociateHostedConnection {
			directconnect_AssociateHostedConnection(cfg, client)
			return
		}
		if _directconnectAssociateMacSecKey {
			directconnect_AssociateMacSecKey(cfg, client)
			return
		}
		if _directconnectAssociateVirtualInterface {
			directconnect_AssociateVirtualInterface(cfg, client)
			return
		}
		if _directconnectConfirmConnection {
			directconnect_ConfirmConnection(cfg, client)
			return
		}
		if _directconnectConfirmCustomerAgreement {
			directconnect_ConfirmCustomerAgreement(cfg, client)
			return
		}
		if _directconnectConfirmPrivateVirtualInterface {
			directconnect_ConfirmPrivateVirtualInterface(cfg, client)
			return
		}
		if _directconnectConfirmPublicVirtualInterface {
			directconnect_ConfirmPublicVirtualInterface(cfg, client)
			return
		}
		if _directconnectConfirmTransitVirtualInterface {
			directconnect_ConfirmTransitVirtualInterface(cfg, client)
			return
		}
		if _directconnectCreateBGPPeer {
			directconnect_CreateBGPPeer(cfg, client)
			return
		}
		if _directconnectCreateConnection {
			directconnect_CreateConnection(cfg, client)
			return
		}
		if _directconnectCreateDirectConnectGateway {
			directconnect_CreateDirectConnectGateway(cfg, client)
			return
		}
		if _directconnectCreateDirectConnectGatewayAssociation {
			directconnect_CreateDirectConnectGatewayAssociation(cfg, client)
			return
		}
		if _directconnectCreateDirectConnectGatewayAssociationProposal {
			directconnect_CreateDirectConnectGatewayAssociationProposal(cfg, client)
			return
		}
		if _directconnectCreateInterconnect {
			directconnect_CreateInterconnect(cfg, client)
			return
		}
		if _directconnectCreateLag {
			directconnect_CreateLag(cfg, client)
			return
		}
		if _directconnectCreatePrivateVirtualInterface {
			directconnect_CreatePrivateVirtualInterface(cfg, client)
			return
		}
		if _directconnectCreatePublicVirtualInterface {
			directconnect_CreatePublicVirtualInterface(cfg, client)
			return
		}
		if _directconnectCreateTransitVirtualInterface {
			directconnect_CreateTransitVirtualInterface(cfg, client)
			return
		}
		if _directconnectDeleteBGPPeer {
			directconnect_DeleteBGPPeer(cfg, client)
			return
		}
		if _directconnectDeleteConnection {
			directconnect_DeleteConnection(cfg, client)
			return
		}
		if _directconnectDeleteDirectConnectGateway {
			directconnect_DeleteDirectConnectGateway(cfg, client)
			return
		}
		if _directconnectDeleteDirectConnectGatewayAssociation {
			directconnect_DeleteDirectConnectGatewayAssociation(cfg, client)
			return
		}
		if _directconnectDeleteDirectConnectGatewayAssociationProposal {
			directconnect_DeleteDirectConnectGatewayAssociationProposal(cfg, client)
			return
		}
		if _directconnectDeleteInterconnect {
			directconnect_DeleteInterconnect(cfg, client)
			return
		}
		if _directconnectDeleteLag {
			directconnect_DeleteLag(cfg, client)
			return
		}
		if _directconnectDeleteVirtualInterface {
			directconnect_DeleteVirtualInterface(cfg, client)
			return
		}
		if _directconnectDescribeConnectionLoa {
			directconnect_DescribeConnectionLoa(cfg, client)
			return
		}
		if _directconnectDescribeConnections {
			directconnect_DescribeConnections(cfg, client)
			return
		}
		if _directconnectDescribeConnectionsOnInterconnect {
			directconnect_DescribeConnectionsOnInterconnect(cfg, client)
			return
		}
		if _directconnectDescribeCustomerMetadata {
			directconnect_DescribeCustomerMetadata(cfg, client)
			return
		}
		if _directconnectDescribeDirectConnectGatewayAssociationProposals {
			directconnect_DescribeDirectConnectGatewayAssociationProposals(cfg, client)
			return
		}
		if _directconnectDescribeDirectConnectGatewayAssociations {
			directconnect_DescribeDirectConnectGatewayAssociations(cfg, client)
			return
		}
		if _directconnectDescribeDirectConnectGatewayAttachments {
			directconnect_DescribeDirectConnectGatewayAttachments(cfg, client)
			return
		}
		if _directconnectDescribeDirectConnectGateways {
			directconnect_DescribeDirectConnectGateways(cfg, client)
			return
		}
		if _directconnectDescribeHostedConnections {
			directconnect_DescribeHostedConnections(cfg, client)
			return
		}
		if _directconnectDescribeInterconnectLoa {
			directconnect_DescribeInterconnectLoa(cfg, client)
			return
		}
		if _directconnectDescribeInterconnects {
			directconnect_DescribeInterconnects(cfg, client)
			return
		}
		if _directconnectDescribeLags {
			directconnect_DescribeLags(cfg, client)
			return
		}
		if _directconnectDescribeLoa {
			directconnect_DescribeLoa(cfg, client)
			return
		}
		if _directconnectDescribeLocations {
			directconnect_DescribeLocations(cfg, client)
			return
		}
		if _directconnectDescribeRouterConfiguration {
			directconnect_DescribeRouterConfiguration(cfg, client)
			return
		}
		if _directconnectDescribeTags {
			directconnect_DescribeTags(cfg, client)
			return
		}
		if _directconnectDescribeVirtualGateways {
			directconnect_DescribeVirtualGateways(cfg, client)
			return
		}
		if _directconnectDescribeVirtualInterfaces {
			directconnect_DescribeVirtualInterfaces(cfg, client)
			return
		}
		if _directconnectDisassociateConnectionFromLag {
			directconnect_DisassociateConnectionFromLag(cfg, client)
			return
		}
		if _directconnectDisassociateMacSecKey {
			directconnect_DisassociateMacSecKey(cfg, client)
			return
		}
		if _directconnectListVirtualInterfaceTestHistory {
			directconnect_ListVirtualInterfaceTestHistory(cfg, client)
			return
		}
		if _directconnectStartBgpFailoverTest {
			directconnect_StartBgpFailoverTest(cfg, client)
			return
		}
		if _directconnectStopBgpFailoverTest {
			directconnect_StopBgpFailoverTest(cfg, client)
			return
		}
		if _directconnectTagResource {
			directconnect_TagResource(cfg, client)
			return
		}
		if _directconnectUntagResource {
			directconnect_UntagResource(cfg, client)
			return
		}
		if _directconnectUpdateConnection {
			directconnect_UpdateConnection(cfg, client)
			return
		}
		if _directconnectUpdateDirectConnectGateway {
			directconnect_UpdateDirectConnectGateway(cfg, client)
			return
		}
		if _directconnectUpdateDirectConnectGatewayAssociation {
			directconnect_UpdateDirectConnectGatewayAssociation(cfg, client)
			return
		}
		if _directconnectUpdateLag {
			directconnect_UpdateLag(cfg, client)
			return
		}
		if _directconnectUpdateVirtualInterfaceAttributes {
			directconnect_UpdateVirtualInterfaceAttributes(cfg, client)
			return
		}

	},
}

var (
	_directconnectAcceptDirectConnectGatewayAssociationProposal    bool
	_directconnectAllocateConnectionOnInterconnect                 bool
	_directconnectAllocateHostedConnection                         bool
	_directconnectAllocatePrivateVirtualInterface                  bool
	_directconnectAllocatePublicVirtualInterface                   bool
	_directconnectAllocateTransitVirtualInterface                  bool
	_directconnectAssociateConnectionWithLag                       bool
	_directconnectAssociateHostedConnection                        bool
	_directconnectAssociateMacSecKey                               bool
	_directconnectAssociateVirtualInterface                        bool
	_directconnectConfirmConnection                                bool
	_directconnectConfirmCustomerAgreement                         bool
	_directconnectConfirmPrivateVirtualInterface                   bool
	_directconnectConfirmPublicVirtualInterface                    bool
	_directconnectConfirmTransitVirtualInterface                   bool
	_directconnectCreateBGPPeer                                    bool
	_directconnectCreateConnection                                 bool
	_directconnectCreateDirectConnectGateway                       bool
	_directconnectCreateDirectConnectGatewayAssociation            bool
	_directconnectCreateDirectConnectGatewayAssociationProposal    bool
	_directconnectCreateInterconnect                               bool
	_directconnectCreateLag                                        bool
	_directconnectCreatePrivateVirtualInterface                    bool
	_directconnectCreatePublicVirtualInterface                     bool
	_directconnectCreateTransitVirtualInterface                    bool
	_directconnectDeleteBGPPeer                                    bool
	_directconnectDeleteConnection                                 bool
	_directconnectDeleteDirectConnectGateway                       bool
	_directconnectDeleteDirectConnectGatewayAssociation            bool
	_directconnectDeleteDirectConnectGatewayAssociationProposal    bool
	_directconnectDeleteInterconnect                               bool
	_directconnectDeleteLag                                        bool
	_directconnectDeleteVirtualInterface                           bool
	_directconnectDescribeConnectionLoa                            bool
	_directconnectDescribeConnections                              bool
	_directconnectDescribeConnectionsOnInterconnect                bool
	_directconnectDescribeCustomerMetadata                         bool
	_directconnectDescribeDirectConnectGatewayAssociationProposals bool
	_directconnectDescribeDirectConnectGatewayAssociations         bool
	_directconnectDescribeDirectConnectGatewayAttachments          bool
	_directconnectDescribeDirectConnectGateways                    bool
	_directconnectDescribeHostedConnections                        bool
	_directconnectDescribeInterconnectLoa                          bool
	_directconnectDescribeInterconnects                            bool
	_directconnectDescribeLags                                     bool
	_directconnectDescribeLoa                                      bool
	_directconnectDescribeLocations                                bool
	_directconnectDescribeRouterConfiguration                      bool
	_directconnectDescribeTags                                     bool
	_directconnectDescribeVirtualGateways                          bool
	_directconnectDescribeVirtualInterfaces                        bool
	_directconnectDisassociateConnectionFromLag                    bool
	_directconnectDisassociateMacSecKey                            bool
	_directconnectListVirtualInterfaceTestHistory                  bool
	_directconnectStartBgpFailoverTest                             bool
	_directconnectStopBgpFailoverTest                              bool
	_directconnectTagResource                                      bool
	_directconnectUntagResource                                    bool
	_directconnectUpdateConnection                                 bool
	_directconnectUpdateDirectConnectGateway                       bool
	_directconnectUpdateDirectConnectGatewayAssociation            bool
	_directconnectUpdateLag                                        bool
	_directconnectUpdateVirtualInterfaceAttributes                 bool

	_directconnectAddAllowedPrefixesToDirectConnectGateway      string
	_directconnectAgreementName                                 string
	_directconnectAmazonSideAsn                                 string
	_directconnectAsn                                           string
	_directconnectAsnLong                                       string
	_directconnectAssociatedGatewayId                           string
	_directconnectAssociatedGatewayOwnerAccount                 string
	_directconnectAssociationId                                 string
	_directconnectBandwidth                                     string
	_directconnectBgpPeerId                                     string
	_directconnectBgpPeers                                      []string
	_directconnectCak                                           string
	_directconnectChildConnectionTags                           string
	_directconnectCkn                                           string
	_directconnectConnectionId                                  string
	_directconnectConnectionName                                string
	_directconnectConnectionsBandwidth                          string
	_directconnectCustomerAddress                               string
	_directconnectDirectConnectGatewayId                        string
	_directconnectDirectConnectGatewayName                      string
	_directconnectDirectConnectGatewayOwnerAccount              string
	_directconnectEnableSiteLink                                string
	_directconnectEncryptionMode                                string
	_directconnectGatewayId                                     string
	_directconnectInterconnectId                                string
	_directconnectInterconnectName                              string
	_directconnectLagId                                         string
	_directconnectLagName                                       string
	_directconnectLoaContentType                                string
	_directconnectLocation                                      string
	_directconnectMaxResults                                    string
	_directconnectMinimumLinks                                  string
	_directconnectMtu                                           string
	_directconnectNewBGPPeer                                    string
	_directconnectNewDirectConnectGatewayName                   string
	_directconnectNewPrivateVirtualInterface                    string
	_directconnectNewPrivateVirtualInterfaceAllocation          string
	_directconnectNewPublicVirtualInterface                     string
	_directconnectNewPublicVirtualInterfaceAllocation           string
	_directconnectNewTransitVirtualInterface                    string
	_directconnectNewTransitVirtualInterfaceAllocation          string
	_directconnectNextToken                                     string
	_directconnectNumberOfConnections                           string
	_directconnectOverrideAllowedPrefixesToDirectConnectGateway string
	_directconnectOwnerAccount                                  string
	_directconnectParentConnectionId                            string
	_directconnectProposalId                                    string
	_directconnectProviderName                                  string
	_directconnectRemoveAllowedPrefixesToDirectConnectGateway   string
	_directconnectRequestMACSec                                 string
	_directconnectResourceArn                                   string
	_directconnectResourceArns                                  []string
	_directconnectRouterTypeIdentifier                          string
	_directconnectSecretARN                                     string
	_directconnectStatus                                        string
	_directconnectTagKeys                                       []string
	_directconnectTags                                          string
	_directconnectTestDurationInMinutes                         string
	_directconnectTestId                                        string
	_directconnectVirtualGatewayId                              string
	_directconnectVirtualInterfaceId                            string
	_directconnectVirtualInterfaceName                          string
	_directconnectVlan                                          string
)

// Accepts a proposal request to attach a virtual private gateway or transit
// gateway to a Direct Connect gateway.
func directconnect_AcceptDirectConnectGatewayAssociationProposal(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.AcceptDirectConnectGatewayAssociationProposalInput{
		// AssociatedGatewayOwnerAccount: *string, // Required
		// DirectConnectGatewayId: *string, // Required
		// ProposalId: *string, // Required
	}

	if len(_directconnectAssociatedGatewayOwnerAccount) > 0 {
		input.AssociatedGatewayOwnerAccount = aws.String(_directconnectAssociatedGatewayOwnerAccount)
	}
	if len(_directconnectDirectConnectGatewayId) > 0 {
		input.DirectConnectGatewayId = aws.String(_directconnectDirectConnectGatewayId)
	}
	if len(_directconnectProposalId) > 0 {
		input.ProposalId = aws.String(_directconnectProposalId)
	}
	if len(_directconnectOverrideAllowedPrefixesToDirectConnectGateway) > 0 {
		if err := assignInputField(input, "OverrideAllowedPrefixesToDirectConnectGateway", _directconnectOverrideAllowedPrefixesToDirectConnectGateway); err != nil {
			log.Errorf("invalid --override-allowed-prefixes-to-direct-connect-gateway: %s", err.Error())
			return
		}
	}

	if resp, err := client.AcceptDirectConnectGatewayAssociationProposal(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deprecated. Use AllocateHostedConnection instead.
// Creates a hosted connection on an interconnect.
//
// Allocates a VLAN number and a specified amount of bandwidth for use by a hosted
// connection on the specified interconnect.
//
// Intended for use by Direct Connect Partners only.
//
// Deprecated: This operation has been deprecated.
func directconnect_AllocateConnectionOnInterconnect(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.AllocateConnectionOnInterconnectInput{
		// Bandwidth: *string, // Required
		// ConnectionName: *string, // Required
		// InterconnectId: *string, // Required
		// OwnerAccount: *string, // Required
		// Vlan: int32, // Required
	}

	if len(_directconnectBandwidth) > 0 {
		input.Bandwidth = aws.String(_directconnectBandwidth)
	}
	if len(_directconnectConnectionName) > 0 {
		input.ConnectionName = aws.String(_directconnectConnectionName)
	}
	if len(_directconnectInterconnectId) > 0 {
		input.InterconnectId = aws.String(_directconnectInterconnectId)
	}
	if len(_directconnectOwnerAccount) > 0 {
		input.OwnerAccount = aws.String(_directconnectOwnerAccount)
	}
	if len(_directconnectVlan) > 0 {
		if err := assignInputField(input, "Vlan", _directconnectVlan); err != nil {
			log.Errorf("invalid --vlan: %s", err.Error())
			return
		}
	}

	if resp, err := client.AllocateConnectionOnInterconnect(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a hosted connection on the specified interconnect or a link aggregation
// group (LAG) of interconnects.
//
// Allocates a VLAN number and a specified amount of capacity (bandwidth) for use
// by a hosted connection on the specified interconnect or LAG of interconnects.
// Amazon Web Services polices the hosted connection for the specified capacity and
// the Direct Connect Partner must also police the hosted connection for the
// specified capacity.
//
// Intended for use by Direct Connect Partners only.
func directconnect_AllocateHostedConnection(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.AllocateHostedConnectionInput{
		// Bandwidth: *string, // Required
		// ConnectionId: *string, // Required
		// ConnectionName: *string, // Required
		// OwnerAccount: *string, // Required
		// Vlan: int32, // Required
	}

	if len(_directconnectBandwidth) > 0 {
		input.Bandwidth = aws.String(_directconnectBandwidth)
	}
	if len(_directconnectConnectionId) > 0 {
		input.ConnectionId = aws.String(_directconnectConnectionId)
	}
	if len(_directconnectConnectionName) > 0 {
		input.ConnectionName = aws.String(_directconnectConnectionName)
	}
	if len(_directconnectOwnerAccount) > 0 {
		input.OwnerAccount = aws.String(_directconnectOwnerAccount)
	}
	if len(_directconnectVlan) > 0 {
		if err := assignInputField(input, "Vlan", _directconnectVlan); err != nil {
			log.Errorf("invalid --vlan: %s", err.Error())
			return
		}
	}
	if len(_directconnectTags) > 0 {
		if err := assignInputField(input, "Tags", _directconnectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.AllocateHostedConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provisions a private virtual interface to be owned by the specified Amazon Web
// Services account.
//
// Virtual interfaces created using this action must be confirmed by the owner
// using ConfirmPrivateVirtualInterface. Until then, the virtual interface is in the Confirming state and is not
// available to handle traffic.
func directconnect_AllocatePrivateVirtualInterface(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.AllocatePrivateVirtualInterfaceInput{
		// ConnectionId: *string, // Required
		// NewPrivateVirtualInterfaceAllocation: *types.NewPrivateVirtualInterfaceAllocation, // Required
		// OwnerAccount: *string, // Required
	}

	if len(_directconnectConnectionId) > 0 {
		input.ConnectionId = aws.String(_directconnectConnectionId)
	}
	if len(_directconnectNewPrivateVirtualInterfaceAllocation) > 0 {
		if err := assignInputField(input, "NewPrivateVirtualInterfaceAllocation", _directconnectNewPrivateVirtualInterfaceAllocation); err != nil {
			log.Errorf("invalid --new-private-virtual-interface-allocation: %s", err.Error())
			return
		}
	}
	if len(_directconnectOwnerAccount) > 0 {
		input.OwnerAccount = aws.String(_directconnectOwnerAccount)
	}

	if resp, err := client.AllocatePrivateVirtualInterface(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provisions a public virtual interface to be owned by the specified Amazon Web
// Services account.
//
// The owner of a connection calls this function to provision a public virtual
// interface to be owned by the specified Amazon Web Services account.
//
// Virtual interfaces created using this function must be confirmed by the owner
// using ConfirmPublicVirtualInterface. Until this step has been completed, the virtual interface is in the
// confirming state and is not available to handle traffic.
//
// When creating an IPv6 public virtual interface, omit the Amazon address and
// customer address. IPv6 addresses are automatically assigned from the Amazon pool
// of IPv6 addresses; you cannot specify custom IPv6 addresses.
func directconnect_AllocatePublicVirtualInterface(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.AllocatePublicVirtualInterfaceInput{
		// ConnectionId: *string, // Required
		// NewPublicVirtualInterfaceAllocation: *types.NewPublicVirtualInterfaceAllocation, // Required
		// OwnerAccount: *string, // Required
	}

	if len(_directconnectConnectionId) > 0 {
		input.ConnectionId = aws.String(_directconnectConnectionId)
	}
	if len(_directconnectNewPublicVirtualInterfaceAllocation) > 0 {
		if err := assignInputField(input, "NewPublicVirtualInterfaceAllocation", _directconnectNewPublicVirtualInterfaceAllocation); err != nil {
			log.Errorf("invalid --new-public-virtual-interface-allocation: %s", err.Error())
			return
		}
	}
	if len(_directconnectOwnerAccount) > 0 {
		input.OwnerAccount = aws.String(_directconnectOwnerAccount)
	}

	if resp, err := client.AllocatePublicVirtualInterface(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provisions a transit virtual interface to be owned by the specified Amazon Web
// Services account. Use this type of interface to connect a transit gateway to
// your Direct Connect gateway.
//
// The owner of a connection provisions a transit virtual interface to be owned by
// the specified Amazon Web Services account.
//
// After you create a transit virtual interface, it must be confirmed by the owner
// using ConfirmTransitVirtualInterface. Until this step has been completed, the transit virtual interface is in
// the requested state and is not available to handle traffic.
func directconnect_AllocateTransitVirtualInterface(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.AllocateTransitVirtualInterfaceInput{
		// ConnectionId: *string, // Required
		// NewTransitVirtualInterfaceAllocation: *types.NewTransitVirtualInterfaceAllocation, // Required
		// OwnerAccount: *string, // Required
	}

	if len(_directconnectConnectionId) > 0 {
		input.ConnectionId = aws.String(_directconnectConnectionId)
	}
	if len(_directconnectNewTransitVirtualInterfaceAllocation) > 0 {
		if err := assignInputField(input, "NewTransitVirtualInterfaceAllocation", _directconnectNewTransitVirtualInterfaceAllocation); err != nil {
			log.Errorf("invalid --new-transit-virtual-interface-allocation: %s", err.Error())
			return
		}
	}
	if len(_directconnectOwnerAccount) > 0 {
		input.OwnerAccount = aws.String(_directconnectOwnerAccount)
	}

	if resp, err := client.AllocateTransitVirtualInterface(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates an existing connection with a link aggregation group (LAG). The
// connection is interrupted and re-established as a member of the LAG
// (connectivity to Amazon Web Services is interrupted). The connection must be
// hosted on the same Direct Connect endpoint as the LAG, and its bandwidth must
// match the bandwidth for the LAG. You can re-associate a connection that's
// currently associated with a different LAG; however, if removing the connection
// would cause the original LAG to fall below its setting for minimum number of
// operational connections, the request fails.
//
// Any virtual interfaces that are directly associated with the connection are
// automatically re-associated with the LAG. If the connection was originally
// associated with a different LAG, the virtual interfaces remain associated with
// the original LAG.
//
// For interconnects, any hosted connections are automatically re-associated with
// the LAG. If the interconnect was originally associated with a different LAG, the
// hosted connections remain associated with the original LAG.
func directconnect_AssociateConnectionWithLag(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.AssociateConnectionWithLagInput{
		// ConnectionId: *string, // Required
		// LagId: *string, // Required
	}

	if len(_directconnectConnectionId) > 0 {
		input.ConnectionId = aws.String(_directconnectConnectionId)
	}
	if len(_directconnectLagId) > 0 {
		input.LagId = aws.String(_directconnectLagId)
	}

	if resp, err := client.AssociateConnectionWithLag(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a hosted connection and its virtual interfaces with a link
// aggregation group (LAG) or interconnect. If the target interconnect or LAG has
// an existing hosted connection with a conflicting VLAN number or IP address, the
// operation fails. This action temporarily interrupts the hosted connection's
// connectivity to Amazon Web Services as it is being migrated.
//
// Intended for use by Direct Connect Partners only.
func directconnect_AssociateHostedConnection(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.AssociateHostedConnectionInput{
		// ConnectionId: *string, // Required
		// ParentConnectionId: *string, // Required
	}

	if len(_directconnectConnectionId) > 0 {
		input.ConnectionId = aws.String(_directconnectConnectionId)
	}
	if len(_directconnectParentConnectionId) > 0 {
		input.ParentConnectionId = aws.String(_directconnectParentConnectionId)
	}

	if resp, err := client.AssociateHostedConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a MAC Security (MACsec) Connection Key Name (CKN)/ Connectivity
// Association Key (CAK) pair with a Direct Connect connection.
//
// You must supply either the secretARN, or the CKN/CAK ( ckn and cak ) pair in the
// request.
//
// For information about MAC Security (MACsec) key considerations, see [MACsec pre-shared CKN/CAK key considerations] in the
// Direct Connect User Guide.
//
// [MACsec pre-shared CKN/CAK key considerations]: https://docs.aws.amazon.com/directconnect/latest/UserGuide/direct-connect-mac-sec-getting-started.html#mac-sec-key-consideration
func directconnect_AssociateMacSecKey(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.AssociateMacSecKeyInput{
		// ConnectionId: *string, // Required
	}

	if len(_directconnectConnectionId) > 0 {
		input.ConnectionId = aws.String(_directconnectConnectionId)
	}
	if len(_directconnectCak) > 0 {
		input.Cak = aws.String(_directconnectCak)
	}
	if len(_directconnectCkn) > 0 {
		input.Ckn = aws.String(_directconnectCkn)
	}
	if len(_directconnectSecretARN) > 0 {
		input.SecretARN = aws.String(_directconnectSecretARN)
	}

	if resp, err := client.AssociateMacSecKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a virtual interface with a specified link aggregation group (LAG) or
// connection. Connectivity to Amazon Web Services is temporarily interrupted as
// the virtual interface is being migrated. If the target connection or LAG has an
// associated virtual interface with a conflicting VLAN number or a conflicting IP
// address, the operation fails.
//
// Virtual interfaces associated with a hosted connection cannot be associated
// with a LAG; hosted connections must be migrated along with their virtual
// interfaces using AssociateHostedConnection.
//
// To reassociate a virtual interface to a new connection or LAG, the requester
// must own either the virtual interface itself or the connection to which the
// virtual interface is currently associated. Additionally, the requester must own
// the connection or LAG for the association.
func directconnect_AssociateVirtualInterface(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.AssociateVirtualInterfaceInput{
		// ConnectionId: *string, // Required
		// VirtualInterfaceId: *string, // Required
	}

	if len(_directconnectConnectionId) > 0 {
		input.ConnectionId = aws.String(_directconnectConnectionId)
	}
	if len(_directconnectVirtualInterfaceId) > 0 {
		input.VirtualInterfaceId = aws.String(_directconnectVirtualInterfaceId)
	}

	if resp, err := client.AssociateVirtualInterface(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Confirms the creation of the specified hosted connection on an interconnect.
// Upon creation, the hosted connection is initially in the Ordering state, and
// remains in this state until the owner confirms creation of the hosted
// connection.
func directconnect_ConfirmConnection(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.ConfirmConnectionInput{
		// ConnectionId: *string, // Required
	}

	if len(_directconnectConnectionId) > 0 {
		input.ConnectionId = aws.String(_directconnectConnectionId)
	}

	if resp, err := client.ConfirmConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The confirmation of the terms of agreement when creating the connection/link
// aggregation group (LAG).
func directconnect_ConfirmCustomerAgreement(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.ConfirmCustomerAgreementInput{}

	if len(_directconnectAgreementName) > 0 {
		input.AgreementName = aws.String(_directconnectAgreementName)
	}

	if resp, err := client.ConfirmCustomerAgreement(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Accepts ownership of a private virtual interface created by another Amazon Web
// Services account.
//
// After the virtual interface owner makes this call, the virtual interface is
// created and attached to the specified virtual private gateway or Direct Connect
// gateway, and is made available to handle traffic.
func directconnect_ConfirmPrivateVirtualInterface(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.ConfirmPrivateVirtualInterfaceInput{
		// VirtualInterfaceId: *string, // Required
	}

	if len(_directconnectVirtualInterfaceId) > 0 {
		input.VirtualInterfaceId = aws.String(_directconnectVirtualInterfaceId)
	}
	if len(_directconnectDirectConnectGatewayId) > 0 {
		input.DirectConnectGatewayId = aws.String(_directconnectDirectConnectGatewayId)
	}
	if len(_directconnectVirtualGatewayId) > 0 {
		input.VirtualGatewayId = aws.String(_directconnectVirtualGatewayId)
	}

	if resp, err := client.ConfirmPrivateVirtualInterface(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Accepts ownership of a public virtual interface created by another Amazon Web
// Services account.
//
// After the virtual interface owner makes this call, the specified virtual
// interface is created and made available to handle traffic.
func directconnect_ConfirmPublicVirtualInterface(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.ConfirmPublicVirtualInterfaceInput{
		// VirtualInterfaceId: *string, // Required
	}

	if len(_directconnectVirtualInterfaceId) > 0 {
		input.VirtualInterfaceId = aws.String(_directconnectVirtualInterfaceId)
	}

	if resp, err := client.ConfirmPublicVirtualInterface(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Accepts ownership of a transit virtual interface created by another Amazon Web
// Services account.
//
// After the owner of the transit virtual interface makes this call, the specified
// transit virtual interface is created and made available to handle traffic.
func directconnect_ConfirmTransitVirtualInterface(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.ConfirmTransitVirtualInterfaceInput{
		// DirectConnectGatewayId: *string, // Required
		// VirtualInterfaceId: *string, // Required
	}

	if len(_directconnectDirectConnectGatewayId) > 0 {
		input.DirectConnectGatewayId = aws.String(_directconnectDirectConnectGatewayId)
	}
	if len(_directconnectVirtualInterfaceId) > 0 {
		input.VirtualInterfaceId = aws.String(_directconnectVirtualInterfaceId)
	}

	if resp, err := client.ConfirmTransitVirtualInterface(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a BGP peer on the specified virtual interface.
// You must create a BGP peer for the corresponding address family (IPv4/IPv6) in
// order to access Amazon Web Services resources that also use that address family.
//
// If logical redundancy is not supported by the connection, interconnect, or LAG,
// the BGP peer cannot be in the same address family as an existing BGP peer on the
// virtual interface.
//
// When creating a IPv6 BGP peer, omit the Amazon address and customer address.
// IPv6 addresses are automatically assigned from the Amazon pool of IPv6
// addresses; you cannot specify custom IPv6 addresses.
//
// If you let Amazon Web Services auto-assign IPv4 addresses, a /30 CIDR will be
// allocated from 169.254.0.0/16. Amazon Web Services does not recommend this
// option if you intend to use the customer router peer IP address as the source
// and destination for traffic. Instead you should use RFC 1918 or other
// addressing, and specify the address yourself. For more information about RFC
// 1918 see [Address Allocation for Private Internets].
//
// For a public virtual interface, the Autonomous System Number (ASN) must be
// private or already on the allow list for the virtual interface.
//
// [Address Allocation for Private Internets]: https://datatracker.ietf.org/doc/html/rfc1918
func directconnect_CreateBGPPeer(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.CreateBGPPeerInput{}

	if len(_directconnectNewBGPPeer) > 0 {
		if err := assignInputField(input, "NewBGPPeer", _directconnectNewBGPPeer); err != nil {
			log.Errorf("invalid --new-bgp-peer: %s", err.Error())
			return
		}
	}
	if len(_directconnectVirtualInterfaceId) > 0 {
		input.VirtualInterfaceId = aws.String(_directconnectVirtualInterfaceId)
	}

	if resp, err := client.CreateBGPPeer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a connection between a customer network and a specific Direct Connect
// location.
//
// A connection links your internal network to an Direct Connect location over a
// standard Ethernet fiber-optic cable. One end of the cable is connected to your
// router, the other to an Direct Connect router.
//
// To find the locations for your Region, use DescribeLocations.
//
// You can automatically add the new connection to a link aggregation group (LAG)
// by specifying a LAG ID in the request. This ensures that the new connection is
// allocated on the same Direct Connect endpoint that hosts the specified LAG. If
// there are no available ports on the endpoint, the request fails and no
// connection is created.
func directconnect_CreateConnection(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.CreateConnectionInput{
		// Bandwidth: *string, // Required
		// ConnectionName: *string, // Required
		// Location: *string, // Required
	}

	if len(_directconnectBandwidth) > 0 {
		input.Bandwidth = aws.String(_directconnectBandwidth)
	}
	if len(_directconnectConnectionName) > 0 {
		input.ConnectionName = aws.String(_directconnectConnectionName)
	}
	if len(_directconnectLocation) > 0 {
		input.Location = aws.String(_directconnectLocation)
	}
	if len(_directconnectLagId) > 0 {
		input.LagId = aws.String(_directconnectLagId)
	}
	if len(_directconnectProviderName) > 0 {
		input.ProviderName = aws.String(_directconnectProviderName)
	}
	if len(_directconnectRequestMACSec) > 0 {
		if err := assignInputField(input, "RequestMACSec", _directconnectRequestMACSec); err != nil {
			log.Errorf("invalid --request-mac-sec: %s", err.Error())
			return
		}
	}
	if len(_directconnectTags) > 0 {
		if err := assignInputField(input, "Tags", _directconnectTags); err != nil {
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

// Creates a Direct Connect gateway, which is an intermediate object that enables
// you to connect a set of virtual interfaces and virtual private gateways. A
// Direct Connect gateway is global and visible in any Amazon Web Services Region
// after it is created. The virtual interfaces and virtual private gateways that
// are connected through a Direct Connect gateway can be in different Amazon Web
// Services Regions. This enables you to connect to a VPC in any Region, regardless
// of the Region in which the virtual interfaces are located, and pass traffic
// between them.
func directconnect_CreateDirectConnectGateway(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.CreateDirectConnectGatewayInput{
		// DirectConnectGatewayName: *string, // Required
	}

	if len(_directconnectDirectConnectGatewayName) > 0 {
		input.DirectConnectGatewayName = aws.String(_directconnectDirectConnectGatewayName)
	}
	if len(_directconnectAmazonSideAsn) > 0 {
		if err := assignInputField(input, "AmazonSideAsn", _directconnectAmazonSideAsn); err != nil {
			log.Errorf("invalid --amazon-side-asn: %s", err.Error())
			return
		}
	}
	if len(_directconnectTags) > 0 {
		if err := assignInputField(input, "Tags", _directconnectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDirectConnectGateway(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an association between a Direct Connect gateway and a virtual private
// gateway. The virtual private gateway must be attached to a VPC and must not be
// associated with another Direct Connect gateway.
func directconnect_CreateDirectConnectGatewayAssociation(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.CreateDirectConnectGatewayAssociationInput{
		// DirectConnectGatewayId: *string, // Required
	}

	if len(_directconnectDirectConnectGatewayId) > 0 {
		input.DirectConnectGatewayId = aws.String(_directconnectDirectConnectGatewayId)
	}
	if len(_directconnectAddAllowedPrefixesToDirectConnectGateway) > 0 {
		if err := assignInputField(input, "AddAllowedPrefixesToDirectConnectGateway", _directconnectAddAllowedPrefixesToDirectConnectGateway); err != nil {
			log.Errorf("invalid --add-allowed-prefixes-to-direct-connect-gateway: %s", err.Error())
			return
		}
	}
	if len(_directconnectGatewayId) > 0 {
		input.GatewayId = aws.String(_directconnectGatewayId)
	}
	if len(_directconnectVirtualGatewayId) > 0 {
		input.VirtualGatewayId = aws.String(_directconnectVirtualGatewayId)
	}

	if resp, err := client.CreateDirectConnectGatewayAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a proposal to associate the specified virtual private gateway or
// transit gateway with the specified Direct Connect gateway.
//
// You can associate a Direct Connect gateway and virtual private gateway or
// transit gateway that is owned by any Amazon Web Services account.
func directconnect_CreateDirectConnectGatewayAssociationProposal(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.CreateDirectConnectGatewayAssociationProposalInput{
		// DirectConnectGatewayId: *string, // Required
		// DirectConnectGatewayOwnerAccount: *string, // Required
		// GatewayId: *string, // Required
	}

	if len(_directconnectDirectConnectGatewayId) > 0 {
		input.DirectConnectGatewayId = aws.String(_directconnectDirectConnectGatewayId)
	}
	if len(_directconnectDirectConnectGatewayOwnerAccount) > 0 {
		input.DirectConnectGatewayOwnerAccount = aws.String(_directconnectDirectConnectGatewayOwnerAccount)
	}
	if len(_directconnectGatewayId) > 0 {
		input.GatewayId = aws.String(_directconnectGatewayId)
	}
	if len(_directconnectAddAllowedPrefixesToDirectConnectGateway) > 0 {
		if err := assignInputField(input, "AddAllowedPrefixesToDirectConnectGateway", _directconnectAddAllowedPrefixesToDirectConnectGateway); err != nil {
			log.Errorf("invalid --add-allowed-prefixes-to-direct-connect-gateway: %s", err.Error())
			return
		}
	}
	if len(_directconnectRemoveAllowedPrefixesToDirectConnectGateway) > 0 {
		if err := assignInputField(input, "RemoveAllowedPrefixesToDirectConnectGateway", _directconnectRemoveAllowedPrefixesToDirectConnectGateway); err != nil {
			log.Errorf("invalid --remove-allowed-prefixes-to-direct-connect-gateway: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDirectConnectGatewayAssociationProposal(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an interconnect between an Direct Connect Partner's network and a
// specific Direct Connect location.
//
// An interconnect is a connection that is capable of hosting other connections.
// The Direct Connect Partner can use an interconnect to provide Direct Connect
// hosted connections to customers through their own network services. Like a
// standard connection, an interconnect links the partner's network to an Direct
// Connect location over a standard Ethernet fiber-optic cable. One end is
// connected to the partner's router, the other to an Direct Connect router.
//
// You can automatically add the new interconnect to a link aggregation group
// (LAG) by specifying a LAG ID in the request. This ensures that the new
// interconnect is allocated on the same Direct Connect endpoint that hosts the
// specified LAG. If there are no available ports on the endpoint, the request
// fails and no interconnect is created.
//
// For each end customer, the Direct Connect Partner provisions a connection on
// their interconnect by calling AllocateHostedConnection. The end customer can then connect to Amazon Web
// Services resources by creating a virtual interface on their connection, using
// the VLAN assigned to them by the Direct Connect Partner.
//
// Intended for use by Direct Connect Partners only.
func directconnect_CreateInterconnect(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.CreateInterconnectInput{
		// Bandwidth: *string, // Required
		// InterconnectName: *string, // Required
		// Location: *string, // Required
	}

	if len(_directconnectBandwidth) > 0 {
		input.Bandwidth = aws.String(_directconnectBandwidth)
	}
	if len(_directconnectInterconnectName) > 0 {
		input.InterconnectName = aws.String(_directconnectInterconnectName)
	}
	if len(_directconnectLocation) > 0 {
		input.Location = aws.String(_directconnectLocation)
	}
	if len(_directconnectLagId) > 0 {
		input.LagId = aws.String(_directconnectLagId)
	}
	if len(_directconnectProviderName) > 0 {
		input.ProviderName = aws.String(_directconnectProviderName)
	}
	if len(_directconnectRequestMACSec) > 0 {
		if err := assignInputField(input, "RequestMACSec", _directconnectRequestMACSec); err != nil {
			log.Errorf("invalid --request-mac-sec: %s", err.Error())
			return
		}
	}
	if len(_directconnectTags) > 0 {
		if err := assignInputField(input, "Tags", _directconnectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateInterconnect(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a link aggregation group (LAG) with the specified number of bundled
// physical dedicated connections between the customer network and a specific
// Direct Connect location. A LAG is a logical interface that uses the Link
// Aggregation Control Protocol (LACP) to aggregate multiple interfaces, enabling
// you to treat them as a single interface.
//
// All connections in a LAG must use the same bandwidth (either 1Gbps, 10Gbps,
// 100Gbps, or 400Gbps) and must terminate at the same Direct Connect endpoint.
//
// You can have up to 10 dedicated connections per location. Regardless of this
// limit, if you request more connections for the LAG than Direct Connect can
// allocate on a single endpoint, no LAG is created..
//
// You can specify an existing physical dedicated connection or interconnect to
// include in the LAG (which counts towards the total number of connections). Doing
// so interrupts the current physical dedicated connection, and re-establishes them
// as a member of the LAG. The LAG will be created on the same Direct Connect
// endpoint to which the dedicated connection terminates. Any virtual interfaces
// associated with the dedicated connection are automatically disassociated and
// re-associated with the LAG. The connection ID does not change.
//
// If the Amazon Web Services account used to create a LAG is a registered Direct
// Connect Partner, the LAG is automatically enabled to host sub-connections. For a
// LAG owned by a partner, any associated virtual interfaces cannot be directly
// configured.
func directconnect_CreateLag(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.CreateLagInput{
		// ConnectionsBandwidth: *string, // Required
		// LagName: *string, // Required
		// Location: *string, // Required
		// NumberOfConnections: int32, // Required
	}

	if len(_directconnectConnectionsBandwidth) > 0 {
		input.ConnectionsBandwidth = aws.String(_directconnectConnectionsBandwidth)
	}
	if len(_directconnectLagName) > 0 {
		input.LagName = aws.String(_directconnectLagName)
	}
	if len(_directconnectLocation) > 0 {
		input.Location = aws.String(_directconnectLocation)
	}
	if len(_directconnectNumberOfConnections) > 0 {
		if err := assignInputField(input, "NumberOfConnections", _directconnectNumberOfConnections); err != nil {
			log.Errorf("invalid --number-of-connections: %s", err.Error())
			return
		}
	}
	if len(_directconnectChildConnectionTags) > 0 {
		if err := assignInputField(input, "ChildConnectionTags", _directconnectChildConnectionTags); err != nil {
			log.Errorf("invalid --child-connection-tags: %s", err.Error())
			return
		}
	}
	if len(_directconnectConnectionId) > 0 {
		input.ConnectionId = aws.String(_directconnectConnectionId)
	}
	if len(_directconnectProviderName) > 0 {
		input.ProviderName = aws.String(_directconnectProviderName)
	}
	if len(_directconnectRequestMACSec) > 0 {
		if err := assignInputField(input, "RequestMACSec", _directconnectRequestMACSec); err != nil {
			log.Errorf("invalid --request-mac-sec: %s", err.Error())
			return
		}
	}
	if len(_directconnectTags) > 0 {
		if err := assignInputField(input, "Tags", _directconnectTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateLag(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a private virtual interface. A virtual interface is the VLAN that
// transports Direct Connect traffic. A private virtual interface can be connected
// to either a Direct Connect gateway or a Virtual Private Gateway (VGW).
// Connecting the private virtual interface to a Direct Connect gateway enables the
// possibility for connecting to multiple VPCs, including VPCs in different Amazon
// Web Services Regions. Connecting the private virtual interface to a VGW only
// provides access to a single VPC within the same Region.
//
// Setting the MTU of a virtual interface to 8500 (jumbo frames) can cause an
// update to the underlying physical connection if it wasn't updated to support
// jumbo frames. Updating the connection disrupts network connectivity for all
// virtual interfaces associated with the connection for up to 30 seconds. To check
// whether your connection supports jumbo frames, call DescribeConnections. To check whether your
// virtual interface supports jumbo frames, call DescribeVirtualInterfaces.
func directconnect_CreatePrivateVirtualInterface(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.CreatePrivateVirtualInterfaceInput{
		// ConnectionId: *string, // Required
		// NewPrivateVirtualInterface: *types.NewPrivateVirtualInterface, // Required
	}

	if len(_directconnectConnectionId) > 0 {
		input.ConnectionId = aws.String(_directconnectConnectionId)
	}
	if len(_directconnectNewPrivateVirtualInterface) > 0 {
		if err := assignInputField(input, "NewPrivateVirtualInterface", _directconnectNewPrivateVirtualInterface); err != nil {
			log.Errorf("invalid --new-private-virtual-interface: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePrivateVirtualInterface(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a public virtual interface. A virtual interface is the VLAN that
// transports Direct Connect traffic. A public virtual interface supports sending
// traffic to public services of Amazon Web Services such as Amazon S3.
//
// When creating an IPv6 public virtual interface ( addressFamily is ipv6 ), leave
// the customer and amazon address fields blank to use auto-assigned IPv6 space.
// Custom IPv6 addresses are not supported.
func directconnect_CreatePublicVirtualInterface(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.CreatePublicVirtualInterfaceInput{
		// ConnectionId: *string, // Required
		// NewPublicVirtualInterface: *types.NewPublicVirtualInterface, // Required
	}

	if len(_directconnectConnectionId) > 0 {
		input.ConnectionId = aws.String(_directconnectConnectionId)
	}
	if len(_directconnectNewPublicVirtualInterface) > 0 {
		if err := assignInputField(input, "NewPublicVirtualInterface", _directconnectNewPublicVirtualInterface); err != nil {
			log.Errorf("invalid --new-public-virtual-interface: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePublicVirtualInterface(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a transit virtual interface. A transit virtual interface should be used
// to access one or more transit gateways associated with Direct Connect gateways.
// A transit virtual interface enables the connection of multiple VPCs attached to
// a transit gateway to a Direct Connect gateway.
//
// If you associate your transit gateway with one or more Direct Connect gateways,
// the Autonomous System Number (ASN) used by the transit gateway and the Direct
// Connect gateway must be different. For example, if you use the default ASN 64512
// for both your the transit gateway and Direct Connect gateway, the association
// request fails.
//
// A jumbo MTU value must be either 1500 or 8500. No other values will be
// accepted. Setting the MTU of a virtual interface to 8500 (jumbo frames) can
// cause an update to the underlying physical connection if it wasn't updated to
// support jumbo frames. Updating the connection disrupts network connectivity for
// all virtual interfaces associated with the connection for up to 30 seconds. To
// check whether your connection supports jumbo frames, call DescribeConnections. To check whether
// your virtual interface supports jumbo frames, call DescribeVirtualInterfaces.
func directconnect_CreateTransitVirtualInterface(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.CreateTransitVirtualInterfaceInput{
		// ConnectionId: *string, // Required
		// NewTransitVirtualInterface: *types.NewTransitVirtualInterface, // Required
	}

	if len(_directconnectConnectionId) > 0 {
		input.ConnectionId = aws.String(_directconnectConnectionId)
	}
	if len(_directconnectNewTransitVirtualInterface) > 0 {
		if err := assignInputField(input, "NewTransitVirtualInterface", _directconnectNewTransitVirtualInterface); err != nil {
			log.Errorf("invalid --new-transit-virtual-interface: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTransitVirtualInterface(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified BGP peer on the specified virtual interface with the
// specified customer address and ASN.
//
// You cannot delete the last BGP peer from a virtual interface.
func directconnect_DeleteBGPPeer(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.DeleteBGPPeerInput{}

	if len(_directconnectAsn) > 0 {
		if err := assignInputField(input, "Asn", _directconnectAsn); err != nil {
			log.Errorf("invalid --asn: %s", err.Error())
			return
		}
	}
	if len(_directconnectAsnLong) > 0 {
		if err := assignInputField(input, "AsnLong", _directconnectAsnLong); err != nil {
			log.Errorf("invalid --asn-long: %s", err.Error())
			return
		}
	}
	if len(_directconnectBgpPeerId) > 0 {
		input.BgpPeerId = aws.String(_directconnectBgpPeerId)
	}
	if len(_directconnectCustomerAddress) > 0 {
		input.CustomerAddress = aws.String(_directconnectCustomerAddress)
	}
	if len(_directconnectVirtualInterfaceId) > 0 {
		input.VirtualInterfaceId = aws.String(_directconnectVirtualInterfaceId)
	}

	if resp, err := client.DeleteBGPPeer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified connection.
// Deleting a connection only stops the Direct Connect port hour and data transfer
// charges. If you are partnering with any third parties to connect with the Direct
// Connect location, you must cancel your service with them separately.
func directconnect_DeleteConnection(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.DeleteConnectionInput{
		// ConnectionId: *string, // Required
	}

	if len(_directconnectConnectionId) > 0 {
		input.ConnectionId = aws.String(_directconnectConnectionId)
	}

	if resp, err := client.DeleteConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified Direct Connect gateway. You must first delete all virtual
// interfaces that are attached to the Direct Connect gateway and disassociate all
// virtual private gateways associated with the Direct Connect gateway.
func directconnect_DeleteDirectConnectGateway(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.DeleteDirectConnectGatewayInput{
		// DirectConnectGatewayId: *string, // Required
	}

	if len(_directconnectDirectConnectGatewayId) > 0 {
		input.DirectConnectGatewayId = aws.String(_directconnectDirectConnectGatewayId)
	}

	if resp, err := client.DeleteDirectConnectGateway(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the association between the specified Direct Connect gateway and
// virtual private gateway.
//
// We recommend that you specify the associationID to delete the association.
// Alternatively, if you own virtual gateway and a Direct Connect gateway
// association, you can specify the virtualGatewayId and directConnectGatewayId to
// delete an association.
func directconnect_DeleteDirectConnectGatewayAssociation(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.DeleteDirectConnectGatewayAssociationInput{}

	if len(_directconnectAssociationId) > 0 {
		input.AssociationId = aws.String(_directconnectAssociationId)
	}
	if len(_directconnectDirectConnectGatewayId) > 0 {
		input.DirectConnectGatewayId = aws.String(_directconnectDirectConnectGatewayId)
	}
	if len(_directconnectVirtualGatewayId) > 0 {
		input.VirtualGatewayId = aws.String(_directconnectVirtualGatewayId)
	}

	if resp, err := client.DeleteDirectConnectGatewayAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the association proposal request between the specified Direct Connect
// gateway and virtual private gateway or transit gateway.
func directconnect_DeleteDirectConnectGatewayAssociationProposal(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.DeleteDirectConnectGatewayAssociationProposalInput{
		// ProposalId: *string, // Required
	}

	if len(_directconnectProposalId) > 0 {
		input.ProposalId = aws.String(_directconnectProposalId)
	}

	if resp, err := client.DeleteDirectConnectGatewayAssociationProposal(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified interconnect.
// Intended for use by Direct Connect Partners only.
func directconnect_DeleteInterconnect(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.DeleteInterconnectInput{
		// InterconnectId: *string, // Required
	}

	if len(_directconnectInterconnectId) > 0 {
		input.InterconnectId = aws.String(_directconnectInterconnectId)
	}

	if resp, err := client.DeleteInterconnect(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified link aggregation group (LAG). You cannot delete a LAG if
// it has active virtual interfaces or hosted connections.
func directconnect_DeleteLag(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.DeleteLagInput{
		// LagId: *string, // Required
	}

	if len(_directconnectLagId) > 0 {
		input.LagId = aws.String(_directconnectLagId)
	}

	if resp, err := client.DeleteLag(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a virtual interface.
func directconnect_DeleteVirtualInterface(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.DeleteVirtualInterfaceInput{
		// VirtualInterfaceId: *string, // Required
	}

	if len(_directconnectVirtualInterfaceId) > 0 {
		input.VirtualInterfaceId = aws.String(_directconnectVirtualInterfaceId)
	}

	if resp, err := client.DeleteVirtualInterface(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deprecated. Use DescribeLoa instead.
// Gets the LOA-CFA for a connection.
//
// The Letter of Authorization - Connecting Facility Assignment (LOA-CFA) is a
// document that your APN partner or service provider uses when establishing your
// cross connect to Amazon Web Services at the colocation facility. For more
// information, see [Requesting Cross Connects at Direct Connect Locations]in the Direct Connect User Guide.
//
// Deprecated: This operation has been deprecated.
//
// [Requesting Cross Connects at Direct Connect Locations]: https://docs.aws.amazon.com/directconnect/latest/UserGuide/Colocation.html
func directconnect_DescribeConnectionLoa(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.DescribeConnectionLoaInput{
		// ConnectionId: *string, // Required
	}

	if len(_directconnectConnectionId) > 0 {
		input.ConnectionId = aws.String(_directconnectConnectionId)
	}
	if len(_directconnectLoaContentType) > 0 {
		if err := assignInputField(input, "LoaContentType", _directconnectLoaContentType); err != nil {
			log.Errorf("invalid --loa-content-type: %s", err.Error())
			return
		}
	}
	if len(_directconnectProviderName) > 0 {
		input.ProviderName = aws.String(_directconnectProviderName)
	}

	if resp, err := client.DescribeConnectionLoa(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Displays the specified connection or all connections in this Region.
func directconnect_DescribeConnections(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.DescribeConnectionsInput{}

	if len(_directconnectConnectionId) > 0 {
		input.ConnectionId = aws.String(_directconnectConnectionId)
	}
	if len(_directconnectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _directconnectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_directconnectNextToken) > 0 {
		input.NextToken = aws.String(_directconnectNextToken)
	}

	if resp, err := client.DescribeConnections(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deprecated. Use DescribeHostedConnections instead.
// Lists the connections that have been provisioned on the specified interconnect.
//
// Intended for use by Direct Connect Partners only.
//
// Deprecated: This operation has been deprecated.
func directconnect_DescribeConnectionsOnInterconnect(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.DescribeConnectionsOnInterconnectInput{
		// InterconnectId: *string, // Required
	}

	if len(_directconnectInterconnectId) > 0 {
		input.InterconnectId = aws.String(_directconnectInterconnectId)
	}

	if resp, err := client.DescribeConnectionsOnInterconnect(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get and view a list of customer agreements, along with their signed status and
// whether the customer is an NNIPartner, NNIPartnerV2, or a nonPartner.
func directconnect_DescribeCustomerMetadata(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.DescribeCustomerMetadataInput{}

	if resp, err := client.DescribeCustomerMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes one or more association proposals for connection between a virtual
// private gateway or transit gateway and a Direct Connect gateway.
func directconnect_DescribeDirectConnectGatewayAssociationProposals(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.DescribeDirectConnectGatewayAssociationProposalsInput{}

	if len(_directconnectAssociatedGatewayId) > 0 {
		input.AssociatedGatewayId = aws.String(_directconnectAssociatedGatewayId)
	}
	if len(_directconnectDirectConnectGatewayId) > 0 {
		input.DirectConnectGatewayId = aws.String(_directconnectDirectConnectGatewayId)
	}
	if len(_directconnectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _directconnectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_directconnectNextToken) > 0 {
		input.NextToken = aws.String(_directconnectNextToken)
	}
	if len(_directconnectProposalId) > 0 {
		input.ProposalId = aws.String(_directconnectProposalId)
	}

	if resp, err := client.DescribeDirectConnectGatewayAssociationProposals(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the associations between your Direct Connect gateways and virtual private
// gateways and transit gateways. You must specify one of the following:
//
// - A Direct Connect gateway
//
// # The response contains all virtual private gateways and transit gateways
//
// associated with the Direct Connect gateway.
//
// - A virtual private gateway
//
// The response contains the Direct Connect gateway.
//
// - A transit gateway
//
// The response contains the Direct Connect gateway.
//
// - A Direct Connect gateway and a virtual private gateway
//
// # The response contains the association between the Direct Connect gateway and
//
// virtual private gateway.
//
// - A Direct Connect gateway and a transit gateway
//
// # The response contains the association between the Direct Connect gateway and
//
// transit gateway.
//
// - A Direct Connect gateway and a virtual private gateway
//
// # The response contains the association between the Direct Connect gateway and
//
// virtual private gateway.
//
// - A Direct Connect gateway association to a Cloud WAN core network
//
// # The response contains the Cloud WAN core network ID that the Direct Connect
//
// gateway is associated to.
func directconnect_DescribeDirectConnectGatewayAssociations(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.DescribeDirectConnectGatewayAssociationsInput{}

	if len(_directconnectAssociatedGatewayId) > 0 {
		input.AssociatedGatewayId = aws.String(_directconnectAssociatedGatewayId)
	}
	if len(_directconnectAssociationId) > 0 {
		input.AssociationId = aws.String(_directconnectAssociationId)
	}
	if len(_directconnectDirectConnectGatewayId) > 0 {
		input.DirectConnectGatewayId = aws.String(_directconnectDirectConnectGatewayId)
	}
	if len(_directconnectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _directconnectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_directconnectNextToken) > 0 {
		input.NextToken = aws.String(_directconnectNextToken)
	}
	if len(_directconnectVirtualGatewayId) > 0 {
		input.VirtualGatewayId = aws.String(_directconnectVirtualGatewayId)
	}

	if resp, err := client.DescribeDirectConnectGatewayAssociations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the attachments between your Direct Connect gateways and virtual
// interfaces. You must specify a Direct Connect gateway, a virtual interface, or
// both. If you specify a Direct Connect gateway, the response contains all virtual
// interfaces attached to the Direct Connect gateway. If you specify a virtual
// interface, the response contains all Direct Connect gateways attached to the
// virtual interface. If you specify both, the response contains the attachment
// between the Direct Connect gateway and the virtual interface.
func directconnect_DescribeDirectConnectGatewayAttachments(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.DescribeDirectConnectGatewayAttachmentsInput{}

	if len(_directconnectDirectConnectGatewayId) > 0 {
		input.DirectConnectGatewayId = aws.String(_directconnectDirectConnectGatewayId)
	}
	if len(_directconnectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _directconnectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_directconnectNextToken) > 0 {
		input.NextToken = aws.String(_directconnectNextToken)
	}
	if len(_directconnectVirtualInterfaceId) > 0 {
		input.VirtualInterfaceId = aws.String(_directconnectVirtualInterfaceId)
	}

	if resp, err := client.DescribeDirectConnectGatewayAttachments(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all your Direct Connect gateways or only the specified Direct Connect
// gateway. Deleted Direct Connect gateways are not returned.
func directconnect_DescribeDirectConnectGateways(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.DescribeDirectConnectGatewaysInput{}

	if len(_directconnectDirectConnectGatewayId) > 0 {
		input.DirectConnectGatewayId = aws.String(_directconnectDirectConnectGatewayId)
	}
	if len(_directconnectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _directconnectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_directconnectNextToken) > 0 {
		input.NextToken = aws.String(_directconnectNextToken)
	}

	if resp, err := client.DescribeDirectConnectGateways(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the hosted connections that have been provisioned on the specified
// interconnect or link aggregation group (LAG).
//
// Intended for use by Direct Connect Partners only.
func directconnect_DescribeHostedConnections(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.DescribeHostedConnectionsInput{
		// ConnectionId: *string, // Required
	}

	if len(_directconnectConnectionId) > 0 {
		input.ConnectionId = aws.String(_directconnectConnectionId)
	}
	if len(_directconnectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _directconnectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_directconnectNextToken) > 0 {
		input.NextToken = aws.String(_directconnectNextToken)
	}

	if resp, err := client.DescribeHostedConnections(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deprecated. Use DescribeLoa instead.
// Gets the LOA-CFA for the specified interconnect.
//
// The Letter of Authorization - Connecting Facility Assignment (LOA-CFA) is a
// document that is used when establishing your cross connect to Amazon Web
// Services at the colocation facility. For more information, see [Requesting Cross Connects at Direct Connect Locations]in the Direct
// Connect User Guide.
//
// Deprecated: This operation has been deprecated.
//
// [Requesting Cross Connects at Direct Connect Locations]: https://docs.aws.amazon.com/directconnect/latest/UserGuide/Colocation.html
func directconnect_DescribeInterconnectLoa(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.DescribeInterconnectLoaInput{
		// InterconnectId: *string, // Required
	}

	if len(_directconnectInterconnectId) > 0 {
		input.InterconnectId = aws.String(_directconnectInterconnectId)
	}
	if len(_directconnectLoaContentType) > 0 {
		if err := assignInputField(input, "LoaContentType", _directconnectLoaContentType); err != nil {
			log.Errorf("invalid --loa-content-type: %s", err.Error())
			return
		}
	}
	if len(_directconnectProviderName) > 0 {
		input.ProviderName = aws.String(_directconnectProviderName)
	}

	if resp, err := client.DescribeInterconnectLoa(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the interconnects owned by the Amazon Web Services account or only the
// specified interconnect.
func directconnect_DescribeInterconnects(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.DescribeInterconnectsInput{}

	if len(_directconnectInterconnectId) > 0 {
		input.InterconnectId = aws.String(_directconnectInterconnectId)
	}
	if len(_directconnectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _directconnectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_directconnectNextToken) > 0 {
		input.NextToken = aws.String(_directconnectNextToken)
	}

	if resp, err := client.DescribeInterconnects(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes all your link aggregation groups (LAG) or the specified LAG.
func directconnect_DescribeLags(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.DescribeLagsInput{}

	if len(_directconnectLagId) > 0 {
		input.LagId = aws.String(_directconnectLagId)
	}
	if len(_directconnectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _directconnectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_directconnectNextToken) > 0 {
		input.NextToken = aws.String(_directconnectNextToken)
	}

	if resp, err := client.DescribeLags(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the LOA-CFA for a connection, interconnect, or link aggregation group
// (LAG).
//
// The Letter of Authorization - Connecting Facility Assignment (LOA-CFA) is a
// document that is used when establishing your cross connect to Amazon Web
// Services at the colocation facility. For more information, see [Requesting Cross Connects at Direct Connect Locations]in the Direct
// Connect User Guide.
//
// [Requesting Cross Connects at Direct Connect Locations]: https://docs.aws.amazon.com/directconnect/latest/UserGuide/Colocation.html
func directconnect_DescribeLoa(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.DescribeLoaInput{
		// ConnectionId: *string, // Required
	}

	if len(_directconnectConnectionId) > 0 {
		input.ConnectionId = aws.String(_directconnectConnectionId)
	}
	if len(_directconnectLoaContentType) > 0 {
		if err := assignInputField(input, "LoaContentType", _directconnectLoaContentType); err != nil {
			log.Errorf("invalid --loa-content-type: %s", err.Error())
			return
		}
	}
	if len(_directconnectProviderName) > 0 {
		input.ProviderName = aws.String(_directconnectProviderName)
	}

	if resp, err := client.DescribeLoa(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the Direct Connect locations in the current Amazon Web Services Region.
// These are the locations that can be selected when calling CreateConnectionor CreateInterconnect.
func directconnect_DescribeLocations(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.DescribeLocationsInput{}

	if resp, err := client.DescribeLocations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Details about the router.
func directconnect_DescribeRouterConfiguration(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.DescribeRouterConfigurationInput{
		// VirtualInterfaceId: *string, // Required
	}

	if len(_directconnectVirtualInterfaceId) > 0 {
		input.VirtualInterfaceId = aws.String(_directconnectVirtualInterfaceId)
	}
	if len(_directconnectRouterTypeIdentifier) > 0 {
		input.RouterTypeIdentifier = aws.String(_directconnectRouterTypeIdentifier)
	}

	if resp, err := client.DescribeRouterConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the tags associated with the specified Direct Connect resources.
func directconnect_DescribeTags(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.DescribeTagsInput{
		// ResourceArns: []string, // Required
	}

	if len(_directconnectResourceArns) > 0 {
		input.ResourceArns = append([]string(nil), _directconnectResourceArns...)
	}

	if resp, err := client.DescribeTags(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deprecated. Use DescribeVpnGateways instead. See [DescribeVPNGateways] in the Amazon Elastic Compute
// Cloud API Reference.
//
// Lists the virtual private gateways owned by the Amazon Web Services account.
//
// You can create one or more Direct Connect private virtual interfaces linked to
// a virtual private gateway.
//
// [DescribeVPNGateways]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeVpnGateways.html
func directconnect_DescribeVirtualGateways(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.DescribeVirtualGatewaysInput{}

	if resp, err := client.DescribeVirtualGateways(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Displays all virtual interfaces for an Amazon Web Services account. Virtual
// interfaces deleted fewer than 15 minutes before you make the request are also
// returned. If you specify a connection ID, only the virtual interfaces associated
// with the connection are returned. If you specify a virtual interface ID, then
// only a single virtual interface is returned.
//
// A virtual interface (VLAN) transmits the traffic between the Direct Connect
// location and the customer network.
//
// - If you're using an asn , the response includes ASN value in both the asn and
// asnLong fields.
//
// - If you're using asnLong , the response returns a value of 0 (zero) for the
// asn attribute because it exceeds the highest ASN value of 2,147,483,647 that
// it can support
func directconnect_DescribeVirtualInterfaces(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.DescribeVirtualInterfacesInput{}

	if len(_directconnectConnectionId) > 0 {
		input.ConnectionId = aws.String(_directconnectConnectionId)
	}
	if len(_directconnectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _directconnectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_directconnectNextToken) > 0 {
		input.NextToken = aws.String(_directconnectNextToken)
	}
	if len(_directconnectVirtualInterfaceId) > 0 {
		input.VirtualInterfaceId = aws.String(_directconnectVirtualInterfaceId)
	}

	if resp, err := client.DescribeVirtualInterfaces(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a connection from a link aggregation group (LAG). The connection
// is interrupted and re-established as a standalone connection (the connection is
// not deleted; to delete the connection, use the DeleteConnectionrequest). If the LAG has
// associated virtual interfaces or hosted connections, they remain associated with
// the LAG. A disassociated connection owned by an Direct Connect Partner is
// automatically converted to an interconnect.
//
// If disassociating the connection would cause the LAG to fall below its setting
// for minimum number of operational connections, the request fails, except when
// it's the last member of the LAG. If all connections are disassociated, the LAG
// continues to exist as an empty LAG with no physical connections.
func directconnect_DisassociateConnectionFromLag(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.DisassociateConnectionFromLagInput{
		// ConnectionId: *string, // Required
		// LagId: *string, // Required
	}

	if len(_directconnectConnectionId) > 0 {
		input.ConnectionId = aws.String(_directconnectConnectionId)
	}
	if len(_directconnectLagId) > 0 {
		input.LagId = aws.String(_directconnectLagId)
	}

	if resp, err := client.DisassociateConnectionFromLag(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the association between a MAC Security (MACsec) security key and a
// Direct Connect connection.
func directconnect_DisassociateMacSecKey(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.DisassociateMacSecKeyInput{
		// ConnectionId: *string, // Required
		// SecretARN: *string, // Required
	}

	if len(_directconnectConnectionId) > 0 {
		input.ConnectionId = aws.String(_directconnectConnectionId)
	}
	if len(_directconnectSecretARN) > 0 {
		input.SecretARN = aws.String(_directconnectSecretARN)
	}

	if resp, err := client.DisassociateMacSecKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the virtual interface failover test history.
func directconnect_ListVirtualInterfaceTestHistory(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.ListVirtualInterfaceTestHistoryInput{}

	if len(_directconnectBgpPeers) > 0 {
		input.BgpPeers = append([]string(nil), _directconnectBgpPeers...)
	}
	if len(_directconnectMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _directconnectMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_directconnectNextToken) > 0 {
		input.NextToken = aws.String(_directconnectNextToken)
	}
	if len(_directconnectStatus) > 0 {
		input.Status = aws.String(_directconnectStatus)
	}
	if len(_directconnectTestId) > 0 {
		input.TestId = aws.String(_directconnectTestId)
	}
	if len(_directconnectVirtualInterfaceId) > 0 {
		input.VirtualInterfaceId = aws.String(_directconnectVirtualInterfaceId)
	}

	if resp, err := client.ListVirtualInterfaceTestHistory(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts the virtual interface failover test that verifies your configuration
// meets your resiliency requirements by placing the BGP peering session in the
// DOWN state. You can then send traffic to verify that there are no outages.
//
// You can run the test on public, private, transit, and hosted virtual interfaces.
//
// You can use [ListVirtualInterfaceTestHistory] to view the virtual interface test history.
//
// If you need to stop the test before the test interval completes, use [StopBgpFailoverTest].
//
// [ListVirtualInterfaceTestHistory]: https://docs.aws.amazon.com/directconnect/latest/APIReference/API_ListVirtualInterfaceTestHistory.html
// [StopBgpFailoverTest]: https://docs.aws.amazon.com/directconnect/latest/APIReference/API_StopBgpFailoverTest.html
func directconnect_StartBgpFailoverTest(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.StartBgpFailoverTestInput{
		// VirtualInterfaceId: *string, // Required
	}

	if len(_directconnectVirtualInterfaceId) > 0 {
		input.VirtualInterfaceId = aws.String(_directconnectVirtualInterfaceId)
	}
	if len(_directconnectBgpPeers) > 0 {
		input.BgpPeers = append([]string(nil), _directconnectBgpPeers...)
	}
	if len(_directconnectTestDurationInMinutes) > 0 {
		if err := assignInputField(input, "TestDurationInMinutes", _directconnectTestDurationInMinutes); err != nil {
			log.Errorf("invalid --test-duration-in-minutes: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartBgpFailoverTest(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops the virtual interface failover test.
func directconnect_StopBgpFailoverTest(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.StopBgpFailoverTestInput{
		// VirtualInterfaceId: *string, // Required
	}

	if len(_directconnectVirtualInterfaceId) > 0 {
		input.VirtualInterfaceId = aws.String(_directconnectVirtualInterfaceId)
	}

	if resp, err := client.StopBgpFailoverTest(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds the specified tags to the specified Direct Connect resource. Each resource
// can have a maximum of 50 tags.
//
// Each tag consists of a key and an optional value. If a tag with the same key is
// already associated with the resource, this action updates its value.
func directconnect_TagResource(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_directconnectResourceArn) > 0 {
		input.ResourceArn = aws.String(_directconnectResourceArn)
	}
	if len(_directconnectTags) > 0 {
		if err := assignInputField(input, "Tags", _directconnectTags); err != nil {
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

// Removes one or more tags from the specified Direct Connect resource.
func directconnect_UntagResource(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_directconnectResourceArn) > 0 {
		input.ResourceArn = aws.String(_directconnectResourceArn)
	}
	if len(_directconnectTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _directconnectTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the Direct Connect connection configuration.
// You can update the following parameters for a connection:
//
// - The connection name
//
// - The connection's MAC Security (MACsec) encryption mode.
func directconnect_UpdateConnection(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.UpdateConnectionInput{
		// ConnectionId: *string, // Required
	}

	if len(_directconnectConnectionId) > 0 {
		input.ConnectionId = aws.String(_directconnectConnectionId)
	}
	if len(_directconnectConnectionName) > 0 {
		input.ConnectionName = aws.String(_directconnectConnectionName)
	}
	if len(_directconnectEncryptionMode) > 0 {
		input.EncryptionMode = aws.String(_directconnectEncryptionMode)
	}

	if resp, err := client.UpdateConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the name of a current Direct Connect gateway.
func directconnect_UpdateDirectConnectGateway(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.UpdateDirectConnectGatewayInput{
		// DirectConnectGatewayId: *string, // Required
		// NewDirectConnectGatewayName: *string, // Required
	}

	if len(_directconnectDirectConnectGatewayId) > 0 {
		input.DirectConnectGatewayId = aws.String(_directconnectDirectConnectGatewayId)
	}
	if len(_directconnectNewDirectConnectGatewayName) > 0 {
		input.NewDirectConnectGatewayName = aws.String(_directconnectNewDirectConnectGatewayName)
	}

	if resp, err := client.UpdateDirectConnectGateway(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified attributes of the Direct Connect gateway association.
// Add or remove prefixes from the association.
func directconnect_UpdateDirectConnectGatewayAssociation(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.UpdateDirectConnectGatewayAssociationInput{}

	if len(_directconnectAddAllowedPrefixesToDirectConnectGateway) > 0 {
		if err := assignInputField(input, "AddAllowedPrefixesToDirectConnectGateway", _directconnectAddAllowedPrefixesToDirectConnectGateway); err != nil {
			log.Errorf("invalid --add-allowed-prefixes-to-direct-connect-gateway: %s", err.Error())
			return
		}
	}
	if len(_directconnectAssociationId) > 0 {
		input.AssociationId = aws.String(_directconnectAssociationId)
	}
	if len(_directconnectRemoveAllowedPrefixesToDirectConnectGateway) > 0 {
		if err := assignInputField(input, "RemoveAllowedPrefixesToDirectConnectGateway", _directconnectRemoveAllowedPrefixesToDirectConnectGateway); err != nil {
			log.Errorf("invalid --remove-allowed-prefixes-to-direct-connect-gateway: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDirectConnectGatewayAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the attributes of the specified link aggregation group (LAG).
// You can update the following LAG attributes:
//
// - The name of the LAG.
//
// - The value for the minimum number of connections that must be operational
// for the LAG itself to be operational.
//
// - The LAG's MACsec encryption mode.
//
// # Amazon Web Services assigns this value to each connection which is part of the
//
// LAG.
//
// - The tags
//
// If you adjust the threshold value for the minimum number of operational
// connections, ensure that the new value does not cause the LAG to fall below the
// threshold and become non-operational.
func directconnect_UpdateLag(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.UpdateLagInput{
		// LagId: *string, // Required
	}

	if len(_directconnectLagId) > 0 {
		input.LagId = aws.String(_directconnectLagId)
	}
	if len(_directconnectEncryptionMode) > 0 {
		input.EncryptionMode = aws.String(_directconnectEncryptionMode)
	}
	if len(_directconnectLagName) > 0 {
		input.LagName = aws.String(_directconnectLagName)
	}
	if len(_directconnectMinimumLinks) > 0 {
		if err := assignInputField(input, "MinimumLinks", _directconnectMinimumLinks); err != nil {
			log.Errorf("invalid --minimum-links: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateLag(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified attributes of the specified virtual private interface.
// Setting the MTU of a virtual interface to 8500 (jumbo frames) can cause an
// update to the underlying physical connection if it wasn't updated to support
// jumbo frames. Updating the connection disrupts network connectivity for all
// virtual interfaces associated with the connection for up to 30 seconds. To check
// whether your connection supports jumbo frames, call DescribeConnections. To check whether your
// virtual interface supports jumbo frames, call DescribeVirtualInterfaces.
func directconnect_UpdateVirtualInterfaceAttributes(cfg aws.Config, client *directconnect.Client) {
	input := &directconnect.UpdateVirtualInterfaceAttributesInput{
		// VirtualInterfaceId: *string, // Required
	}

	if len(_directconnectVirtualInterfaceId) > 0 {
		input.VirtualInterfaceId = aws.String(_directconnectVirtualInterfaceId)
	}
	if len(_directconnectEnableSiteLink) > 0 {
		if err := assignInputField(input, "EnableSiteLink", _directconnectEnableSiteLink); err != nil {
			log.Errorf("invalid --enable-site-link: %s", err.Error())
			return
		}
	}
	if len(_directconnectMtu) > 0 {
		if err := assignInputField(input, "Mtu", _directconnectMtu); err != nil {
			log.Errorf("invalid --mtu: %s", err.Error())
			return
		}
	}
	if len(_directconnectVirtualInterfaceName) > 0 {
		input.VirtualInterfaceName = aws.String(_directconnectVirtualInterfaceName)
	}

	if resp, err := client.UpdateVirtualInterfaceAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_directconnectCmd)
	_directconnectCmd.Flags().SortFlags = false

	_directconnectCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_directconnectCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_directconnectCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_directconnectCmd.Flags().StringVarP(&_directconnectAddAllowedPrefixesToDirectConnectGateway, "add-allowed-prefixes-to-direct-connect-gateway", "", "", "Add Allowed Prefixes To Direct Connect Gateway")
	_directconnectCmd.Flags().StringVarP(&_directconnectAgreementName, "agreement-name", "", "", "Agreement Name")
	_directconnectCmd.Flags().StringVarP(&_directconnectAmazonSideAsn, "amazon-side-asn", "", "", "Amazon Side Asn")
	_directconnectCmd.Flags().StringVarP(&_directconnectAsn, "asn", "", "", "Asn")
	_directconnectCmd.Flags().StringVarP(&_directconnectAsnLong, "asn-long", "", "", "Asn Long")
	_directconnectCmd.Flags().StringVarP(&_directconnectAssociatedGatewayId, "associated-gateway-id", "", "", "Associated Gateway ID")
	_directconnectCmd.Flags().StringVarP(&_directconnectAssociatedGatewayOwnerAccount, "associated-gateway-owner-account", "", "", "Associated Gateway Owner Account")
	_directconnectCmd.Flags().StringVarP(&_directconnectAssociationId, "association-id", "", "", "Association ID")
	_directconnectCmd.Flags().StringVarP(&_directconnectBandwidth, "bandwidth", "", "", "Bandwidth")
	_directconnectCmd.Flags().StringVarP(&_directconnectBgpPeerId, "bgp-peer-id", "", "", "Bgp Peer ID")
	_directconnectCmd.Flags().StringSliceVarP(&_directconnectBgpPeers, "bgp-peers", "", nil, "Bgp Peers")
	_directconnectCmd.Flags().StringVarP(&_directconnectCak, "cak", "", "", "Cak")
	_directconnectCmd.Flags().StringVarP(&_directconnectChildConnectionTags, "child-connection-tags", "", "", "Child Connection Tags")
	_directconnectCmd.Flags().StringVarP(&_directconnectCkn, "ckn", "", "", "Ckn")
	_directconnectCmd.Flags().StringVarP(&_directconnectConnectionId, "connection-id", "", "", "Connection ID")
	_directconnectCmd.Flags().StringVarP(&_directconnectConnectionName, "connection-name", "", "", "Connection Name")
	_directconnectCmd.Flags().StringVarP(&_directconnectConnectionsBandwidth, "connections-bandwidth", "", "", "Connections Bandwidth")
	_directconnectCmd.Flags().StringVarP(&_directconnectCustomerAddress, "customer-address", "", "", "Customer Address")
	_directconnectCmd.Flags().StringVarP(&_directconnectDirectConnectGatewayId, "direct-connect-gateway-id", "", "", "Direct Connect Gateway ID")
	_directconnectCmd.Flags().StringVarP(&_directconnectDirectConnectGatewayName, "direct-connect-gateway-name", "", "", "Direct Connect Gateway Name")
	_directconnectCmd.Flags().StringVarP(&_directconnectDirectConnectGatewayOwnerAccount, "direct-connect-gateway-owner-account", "", "", "Direct Connect Gateway Owner Account")
	_directconnectCmd.Flags().StringVarP(&_directconnectEnableSiteLink, "enable-site-link", "", "", "Enable Site Link")
	_directconnectCmd.Flags().StringVarP(&_directconnectEncryptionMode, "encryption-mode", "", "", "Encryption Mode")
	_directconnectCmd.Flags().StringVarP(&_directconnectGatewayId, "gateway-id", "", "", "Gateway ID")
	_directconnectCmd.Flags().StringVarP(&_directconnectInterconnectId, "interconnect-id", "", "", "Interconnect ID")
	_directconnectCmd.Flags().StringVarP(&_directconnectInterconnectName, "interconnect-name", "", "", "Interconnect Name")
	_directconnectCmd.Flags().StringVarP(&_directconnectLagId, "lag-id", "", "", "Lag ID")
	_directconnectCmd.Flags().StringVarP(&_directconnectLagName, "lag-name", "", "", "Lag Name")
	_directconnectCmd.Flags().StringVarP(&_directconnectLoaContentType, "loa-content-type", "", "", "Loa Content Type")
	_directconnectCmd.Flags().StringVarP(&_directconnectLocation, "location", "", "", "Location")
	_directconnectCmd.Flags().StringVarP(&_directconnectMaxResults, "max-results", "", "", "Max Results")
	_directconnectCmd.Flags().StringVarP(&_directconnectMinimumLinks, "minimum-links", "", "", "Minimum Links")
	_directconnectCmd.Flags().StringVarP(&_directconnectMtu, "mtu", "", "", "Mtu")
	_directconnectCmd.Flags().StringVarP(&_directconnectNewBGPPeer, "new-bgp-peer", "", "", "New Bgp Peer")
	_directconnectCmd.Flags().StringVarP(&_directconnectNewDirectConnectGatewayName, "new-direct-connect-gateway-name", "", "", "New Direct Connect Gateway Name")
	_directconnectCmd.Flags().StringVarP(&_directconnectNewPrivateVirtualInterface, "new-private-virtual-interface", "", "", "New Private Virtual Interface")
	_directconnectCmd.Flags().StringVarP(&_directconnectNewPrivateVirtualInterfaceAllocation, "new-private-virtual-interface-allocation", "", "", "New Private Virtual Interface Allocation")
	_directconnectCmd.Flags().StringVarP(&_directconnectNewPublicVirtualInterface, "new-public-virtual-interface", "", "", "New Public Virtual Interface")
	_directconnectCmd.Flags().StringVarP(&_directconnectNewPublicVirtualInterfaceAllocation, "new-public-virtual-interface-allocation", "", "", "New Public Virtual Interface Allocation")
	_directconnectCmd.Flags().StringVarP(&_directconnectNewTransitVirtualInterface, "new-transit-virtual-interface", "", "", "New Transit Virtual Interface")
	_directconnectCmd.Flags().StringVarP(&_directconnectNewTransitVirtualInterfaceAllocation, "new-transit-virtual-interface-allocation", "", "", "New Transit Virtual Interface Allocation")
	_directconnectCmd.Flags().StringVarP(&_directconnectNextToken, "next-token", "", "", "Next Token")
	_directconnectCmd.Flags().StringVarP(&_directconnectNumberOfConnections, "number-of-connections", "", "", "Number Of Connections")
	_directconnectCmd.Flags().StringVarP(&_directconnectOverrideAllowedPrefixesToDirectConnectGateway, "override-allowed-prefixes-to-direct-connect-gateway", "", "", "Override Allowed Prefixes To Direct Connect Gateway")
	_directconnectCmd.Flags().StringVarP(&_directconnectOwnerAccount, "owner-account", "", "", "Owner Account")
	_directconnectCmd.Flags().StringVarP(&_directconnectParentConnectionId, "parent-connection-id", "", "", "Parent Connection ID")
	_directconnectCmd.Flags().StringVarP(&_directconnectProposalId, "proposal-id", "", "", "Proposal ID")
	_directconnectCmd.Flags().StringVarP(&_directconnectProviderName, "provider-name", "", "", "Provider Name")
	_directconnectCmd.Flags().StringVarP(&_directconnectRemoveAllowedPrefixesToDirectConnectGateway, "remove-allowed-prefixes-to-direct-connect-gateway", "", "", "Remove Allowed Prefixes To Direct Connect Gateway")
	_directconnectCmd.Flags().StringVarP(&_directconnectRequestMACSec, "request-mac-sec", "", "", "Request Mac Sec")
	_directconnectCmd.Flags().StringVarP(&_directconnectResourceArn, "resource-arn", "", "", "Resource ARN")
	_directconnectCmd.Flags().StringSliceVarP(&_directconnectResourceArns, "resource-arns", "", nil, "Resource Arns")
	_directconnectCmd.Flags().StringVarP(&_directconnectRouterTypeIdentifier, "router-type-identifier", "", "", "Router Type Identifier")
	_directconnectCmd.Flags().StringVarP(&_directconnectSecretARN, "secret-arn", "", "", "Secret ARN")
	_directconnectCmd.Flags().StringVarP(&_directconnectStatus, "status", "", "", "Status")
	_directconnectCmd.Flags().StringSliceVarP(&_directconnectTagKeys, "tag-keys", "", nil, "Tag Keys")
	_directconnectCmd.Flags().StringVarP(&_directconnectTags, "tags", "", "", "Tags")
	_directconnectCmd.Flags().StringVarP(&_directconnectTestDurationInMinutes, "test-duration-in-minutes", "", "", "Test Duration In Minutes")
	_directconnectCmd.Flags().StringVarP(&_directconnectTestId, "test-id", "", "", "Test ID")
	_directconnectCmd.Flags().StringVarP(&_directconnectVirtualGatewayId, "virtual-gateway-id", "", "", "Virtual Gateway ID")
	_directconnectCmd.Flags().StringVarP(&_directconnectVirtualInterfaceId, "virtual-interface-id", "", "", "Virtual Interface ID")
	_directconnectCmd.Flags().StringVarP(&_directconnectVirtualInterfaceName, "virtual-interface-name", "", "", "Virtual Interface Name")
	_directconnectCmd.Flags().StringVarP(&_directconnectVlan, "vlan", "", "", "Vlan")

	_directconnectCmd.Flags().BoolVarP(&_directconnectAcceptDirectConnectGatewayAssociationProposal, "accept-direct-connect-gateway-association-proposal", "", false, "Accept Direct Connect Gateway Association Proposal")
	_directconnectCmd.Flags().BoolVarP(&_directconnectAllocateConnectionOnInterconnect, "allocate-connection-on-interconnect", "", false, "Allocate Connection On Interconnect")
	_directconnectCmd.Flags().BoolVarP(&_directconnectAllocateHostedConnection, "allocate-hosted-connection", "", false, "Allocate Hosted Connection")
	_directconnectCmd.Flags().BoolVarP(&_directconnectAllocatePrivateVirtualInterface, "allocate-private-virtual-interface", "", false, "Allocate Private Virtual Interface")
	_directconnectCmd.Flags().BoolVarP(&_directconnectAllocatePublicVirtualInterface, "allocate-public-virtual-interface", "", false, "Allocate Public Virtual Interface")
	_directconnectCmd.Flags().BoolVarP(&_directconnectAllocateTransitVirtualInterface, "allocate-transit-virtual-interface", "", false, "Allocate Transit Virtual Interface")
	_directconnectCmd.Flags().BoolVarP(&_directconnectAssociateConnectionWithLag, "associate-connection-with-lag", "", false, "Associate Connection With Lag")
	_directconnectCmd.Flags().BoolVarP(&_directconnectAssociateHostedConnection, "associate-hosted-connection", "", false, "Associate Hosted Connection")
	_directconnectCmd.Flags().BoolVarP(&_directconnectAssociateMacSecKey, "associate-mac-sec-key", "", false, "Associate Mac Sec Key")
	_directconnectCmd.Flags().BoolVarP(&_directconnectAssociateVirtualInterface, "associate-virtual-interface", "", false, "Associate Virtual Interface")
	_directconnectCmd.Flags().BoolVarP(&_directconnectConfirmConnection, "confirm-connection", "", false, "Confirm Connection")
	_directconnectCmd.Flags().BoolVarP(&_directconnectConfirmCustomerAgreement, "confirm-customer-agreement", "", false, "Confirm Customer Agreement")
	_directconnectCmd.Flags().BoolVarP(&_directconnectConfirmPrivateVirtualInterface, "confirm-private-virtual-interface", "", false, "Confirm Private Virtual Interface")
	_directconnectCmd.Flags().BoolVarP(&_directconnectConfirmPublicVirtualInterface, "confirm-public-virtual-interface", "", false, "Confirm Public Virtual Interface")
	_directconnectCmd.Flags().BoolVarP(&_directconnectConfirmTransitVirtualInterface, "confirm-transit-virtual-interface", "", false, "Confirm Transit Virtual Interface")
	_directconnectCmd.Flags().BoolVarP(&_directconnectCreateBGPPeer, "create-bgp-peer", "", false, "Create Bgp Peer")
	_directconnectCmd.Flags().BoolVarP(&_directconnectCreateConnection, "create-connection", "", false, "Create Connection")
	_directconnectCmd.Flags().BoolVarP(&_directconnectCreateDirectConnectGateway, "create-direct-connect-gateway", "", false, "Create Direct Connect Gateway")
	_directconnectCmd.Flags().BoolVarP(&_directconnectCreateDirectConnectGatewayAssociation, "create-direct-connect-gateway-association", "", false, "Create Direct Connect Gateway Association")
	_directconnectCmd.Flags().BoolVarP(&_directconnectCreateDirectConnectGatewayAssociationProposal, "create-direct-connect-gateway-association-proposal", "", false, "Create Direct Connect Gateway Association Proposal")
	_directconnectCmd.Flags().BoolVarP(&_directconnectCreateInterconnect, "create-interconnect", "", false, "Create Interconnect")
	_directconnectCmd.Flags().BoolVarP(&_directconnectCreateLag, "create-lag", "", false, "Create Lag")
	_directconnectCmd.Flags().BoolVarP(&_directconnectCreatePrivateVirtualInterface, "create-private-virtual-interface", "", false, "Create Private Virtual Interface")
	_directconnectCmd.Flags().BoolVarP(&_directconnectCreatePublicVirtualInterface, "create-public-virtual-interface", "", false, "Create Public Virtual Interface")
	_directconnectCmd.Flags().BoolVarP(&_directconnectCreateTransitVirtualInterface, "create-transit-virtual-interface", "", false, "Create Transit Virtual Interface")
	_directconnectCmd.Flags().BoolVarP(&_directconnectDeleteBGPPeer, "delete-bgp-peer", "", false, "Delete Bgp Peer")
	_directconnectCmd.Flags().BoolVarP(&_directconnectDeleteConnection, "delete-connection", "", false, "Delete Connection")
	_directconnectCmd.Flags().BoolVarP(&_directconnectDeleteDirectConnectGateway, "delete-direct-connect-gateway", "", false, "Delete Direct Connect Gateway")
	_directconnectCmd.Flags().BoolVarP(&_directconnectDeleteDirectConnectGatewayAssociation, "delete-direct-connect-gateway-association", "", false, "Delete Direct Connect Gateway Association")
	_directconnectCmd.Flags().BoolVarP(&_directconnectDeleteDirectConnectGatewayAssociationProposal, "delete-direct-connect-gateway-association-proposal", "", false, "Delete Direct Connect Gateway Association Proposal")
	_directconnectCmd.Flags().BoolVarP(&_directconnectDeleteInterconnect, "delete-interconnect", "", false, "Delete Interconnect")
	_directconnectCmd.Flags().BoolVarP(&_directconnectDeleteLag, "delete-lag", "", false, "Delete Lag")
	_directconnectCmd.Flags().BoolVarP(&_directconnectDeleteVirtualInterface, "delete-virtual-interface", "", false, "Delete Virtual Interface")
	_directconnectCmd.Flags().BoolVarP(&_directconnectDescribeConnectionLoa, "describe-connection-loa", "", false, "Describe Connection Loa")
	_directconnectCmd.Flags().BoolVarP(&_directconnectDescribeConnections, "describe-connections", "", false, "Describe Connections")
	_directconnectCmd.Flags().BoolVarP(&_directconnectDescribeConnectionsOnInterconnect, "describe-connections-on-interconnect", "", false, "Describe Connections On Interconnect")
	_directconnectCmd.Flags().BoolVarP(&_directconnectDescribeCustomerMetadata, "describe-customer-metadata", "", false, "Describe Customer Metadata")
	_directconnectCmd.Flags().BoolVarP(&_directconnectDescribeDirectConnectGatewayAssociationProposals, "describe-direct-connect-gateway-association-proposals", "", false, "Describe Direct Connect Gateway Association Proposals")
	_directconnectCmd.Flags().BoolVarP(&_directconnectDescribeDirectConnectGatewayAssociations, "describe-direct-connect-gateway-associations", "", false, "Describe Direct Connect Gateway Associations")
	_directconnectCmd.Flags().BoolVarP(&_directconnectDescribeDirectConnectGatewayAttachments, "describe-direct-connect-gateway-attachments", "", false, "Describe Direct Connect Gateway Attachments")
	_directconnectCmd.Flags().BoolVarP(&_directconnectDescribeDirectConnectGateways, "describe-direct-connect-gateways", "", false, "Describe Direct Connect Gateways")
	_directconnectCmd.Flags().BoolVarP(&_directconnectDescribeHostedConnections, "describe-hosted-connections", "", false, "Describe Hosted Connections")
	_directconnectCmd.Flags().BoolVarP(&_directconnectDescribeInterconnectLoa, "describe-interconnect-loa", "", false, "Describe Interconnect Loa")
	_directconnectCmd.Flags().BoolVarP(&_directconnectDescribeInterconnects, "describe-interconnects", "", false, "Describe Interconnects")
	_directconnectCmd.Flags().BoolVarP(&_directconnectDescribeLags, "describe-lags", "", false, "Describe Lags")
	_directconnectCmd.Flags().BoolVarP(&_directconnectDescribeLoa, "describe-loa", "", false, "Describe Loa")
	_directconnectCmd.Flags().BoolVarP(&_directconnectDescribeLocations, "describe-locations", "", false, "Describe Locations")
	_directconnectCmd.Flags().BoolVarP(&_directconnectDescribeRouterConfiguration, "describe-router-configuration", "", false, "Describe Router Configuration")
	_directconnectCmd.Flags().BoolVarP(&_directconnectDescribeTags, "describe-tags", "", false, "Describe Tags")
	_directconnectCmd.Flags().BoolVarP(&_directconnectDescribeVirtualGateways, "describe-virtual-gateways", "", false, "Describe Virtual Gateways")
	_directconnectCmd.Flags().BoolVarP(&_directconnectDescribeVirtualInterfaces, "describe-virtual-interfaces", "", false, "Describe Virtual Interfaces")
	_directconnectCmd.Flags().BoolVarP(&_directconnectDisassociateConnectionFromLag, "disassociate-connection-from-lag", "", false, "Disassociate Connection From Lag")
	_directconnectCmd.Flags().BoolVarP(&_directconnectDisassociateMacSecKey, "disassociate-mac-sec-key", "", false, "Disassociate Mac Sec Key")
	_directconnectCmd.Flags().BoolVarP(&_directconnectListVirtualInterfaceTestHistory, "list-virtual-interface-test-history", "", false, "List Virtual Interface Test History")
	_directconnectCmd.Flags().BoolVarP(&_directconnectStartBgpFailoverTest, "start-bgp-failover-test", "", false, "Start Bgp Failover Test")
	_directconnectCmd.Flags().BoolVarP(&_directconnectStopBgpFailoverTest, "stop-bgp-failover-test", "", false, "Stop Bgp Failover Test")
	_directconnectCmd.Flags().BoolVarP(&_directconnectTagResource, "tag-resource", "", false, "Tag Resource")
	_directconnectCmd.Flags().BoolVarP(&_directconnectUntagResource, "untag-resource", "", false, "Untag Resource")
	_directconnectCmd.Flags().BoolVarP(&_directconnectUpdateConnection, "update-connection", "", false, "Update Connection")
	_directconnectCmd.Flags().BoolVarP(&_directconnectUpdateDirectConnectGateway, "update-direct-connect-gateway", "", false, "Update Direct Connect Gateway")
	_directconnectCmd.Flags().BoolVarP(&_directconnectUpdateDirectConnectGatewayAssociation, "update-direct-connect-gateway-association", "", false, "Update Direct Connect Gateway Association")
	_directconnectCmd.Flags().BoolVarP(&_directconnectUpdateLag, "update-lag", "", false, "Update Lag")
	_directconnectCmd.Flags().BoolVarP(&_directconnectUpdateVirtualInterfaceAttributes, "update-virtual-interface-attributes", "", false, "Update Virtual Interface Attributes")

}
