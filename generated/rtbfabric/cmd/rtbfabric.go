package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rtbfabric"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// rtbfabricCmd represents the rtbfabric command
var _rtbfabricCmd = &cobra.Command{
	Use:   "rtbfabric",
	Short: "AWS rtbfabric CLI",
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
		client := rtbfabric.NewFromConfig(cfg)
		if _rtbfabricAcceptLink {
			rtbfabric_AcceptLink(cfg, client)
			return
		}
		if _rtbfabricCreateInboundExternalLink {
			rtbfabric_CreateInboundExternalLink(cfg, client)
			return
		}
		if _rtbfabricCreateLink {
			rtbfabric_CreateLink(cfg, client)
			return
		}
		if _rtbfabricCreateOutboundExternalLink {
			rtbfabric_CreateOutboundExternalLink(cfg, client)
			return
		}
		if _rtbfabricCreateRequesterGateway {
			rtbfabric_CreateRequesterGateway(cfg, client)
			return
		}
		if _rtbfabricCreateResponderGateway {
			rtbfabric_CreateResponderGateway(cfg, client)
			return
		}
		if _rtbfabricDeleteInboundExternalLink {
			rtbfabric_DeleteInboundExternalLink(cfg, client)
			return
		}
		if _rtbfabricDeleteLink {
			rtbfabric_DeleteLink(cfg, client)
			return
		}
		if _rtbfabricDeleteOutboundExternalLink {
			rtbfabric_DeleteOutboundExternalLink(cfg, client)
			return
		}
		if _rtbfabricDeleteRequesterGateway {
			rtbfabric_DeleteRequesterGateway(cfg, client)
			return
		}
		if _rtbfabricDeleteResponderGateway {
			rtbfabric_DeleteResponderGateway(cfg, client)
			return
		}
		if _rtbfabricGetInboundExternalLink {
			rtbfabric_GetInboundExternalLink(cfg, client)
			return
		}
		if _rtbfabricGetLink {
			rtbfabric_GetLink(cfg, client)
			return
		}
		if _rtbfabricGetOutboundExternalLink {
			rtbfabric_GetOutboundExternalLink(cfg, client)
			return
		}
		if _rtbfabricGetRequesterGateway {
			rtbfabric_GetRequesterGateway(cfg, client)
			return
		}
		if _rtbfabricGetResponderGateway {
			rtbfabric_GetResponderGateway(cfg, client)
			return
		}
		if _rtbfabricListLinks {
			rtbfabric_ListLinks(cfg, client)
			return
		}
		if _rtbfabricListRequesterGateways {
			rtbfabric_ListRequesterGateways(cfg, client)
			return
		}
		if _rtbfabricListResponderGateways {
			rtbfabric_ListResponderGateways(cfg, client)
			return
		}
		if _rtbfabricListTagsForResource {
			rtbfabric_ListTagsForResource(cfg, client)
			return
		}
		if _rtbfabricRejectLink {
			rtbfabric_RejectLink(cfg, client)
			return
		}
		if _rtbfabricTagResource {
			rtbfabric_TagResource(cfg, client)
			return
		}
		if _rtbfabricUntagResource {
			rtbfabric_UntagResource(cfg, client)
			return
		}
		if _rtbfabricUpdateLink {
			rtbfabric_UpdateLink(cfg, client)
			return
		}
		if _rtbfabricUpdateLinkModuleFlow {
			rtbfabric_UpdateLinkModuleFlow(cfg, client)
			return
		}
		if _rtbfabricUpdateRequesterGateway {
			rtbfabric_UpdateRequesterGateway(cfg, client)
			return
		}
		if _rtbfabricUpdateResponderGateway {
			rtbfabric_UpdateResponderGateway(cfg, client)
			return
		}

	},
}

var (
	_rtbfabricAcceptLink                 bool
	_rtbfabricCreateInboundExternalLink  bool
	_rtbfabricCreateLink                 bool
	_rtbfabricCreateOutboundExternalLink bool
	_rtbfabricCreateRequesterGateway     bool
	_rtbfabricCreateResponderGateway     bool
	_rtbfabricDeleteInboundExternalLink  bool
	_rtbfabricDeleteLink                 bool
	_rtbfabricDeleteOutboundExternalLink bool
	_rtbfabricDeleteRequesterGateway     bool
	_rtbfabricDeleteResponderGateway     bool
	_rtbfabricGetInboundExternalLink     bool
	_rtbfabricGetLink                    bool
	_rtbfabricGetOutboundExternalLink    bool
	_rtbfabricGetRequesterGateway        bool
	_rtbfabricGetResponderGateway        bool
	_rtbfabricListLinks                  bool
	_rtbfabricListRequesterGateways      bool
	_rtbfabricListResponderGateways      bool
	_rtbfabricListTagsForResource        bool
	_rtbfabricRejectLink                 bool
	_rtbfabricTagResource                bool
	_rtbfabricUntagResource              bool
	_rtbfabricUpdateLink                 bool
	_rtbfabricUpdateLinkModuleFlow       bool
	_rtbfabricUpdateRequesterGateway     bool
	_rtbfabricUpdateResponderGateway     bool

	_rtbfabricAttributes                   string
	_rtbfabricClientToken                  string
	_rtbfabricDescription                  string
	_rtbfabricDomainName                   string
	_rtbfabricGatewayId                    string
	_rtbfabricHttpResponderAllowed         string
	_rtbfabricLinkId                       string
	_rtbfabricLogSettings                  string
	_rtbfabricManagedEndpointConfiguration string
	_rtbfabricMaxResults                   string
	_rtbfabricModules                      string
	_rtbfabricNextToken                    string
	_rtbfabricPeerGatewayId                string
	_rtbfabricPort                         string
	_rtbfabricProtocol                     string
	_rtbfabricPublicEndpoint               string
	_rtbfabricResourceArn                  string
	_rtbfabricSecurityGroupIds             []string
	_rtbfabricSubnetIds                    []string
	_rtbfabricTagKeys                      []string
	_rtbfabricTags                         string
	_rtbfabricTrustStoreConfiguration      string
	_rtbfabricVpcId                        string
)

// Accepts a link request between gateways.
// When a requester gateway requests to link with a responder gateway, the
// responder can use this operation to accept the link request and establish the
// connection.
func rtbfabric_AcceptLink(cfg aws.Config, client *rtbfabric.Client) {
	input := &rtbfabric.AcceptLinkInput{
		// GatewayId: *string, // Required
		// LinkId: *string, // Required
		// LogSettings: *types.LinkLogSettings, // Required
	}

	if len(_rtbfabricGatewayId) > 0 {
		input.GatewayId = aws.String(_rtbfabricGatewayId)
	}
	if len(_rtbfabricLinkId) > 0 {
		input.LinkId = aws.String(_rtbfabricLinkId)
	}
	if len(_rtbfabricLogSettings) > 0 {
		if err := assignInputField(input, "LogSettings", _rtbfabricLogSettings); err != nil {
			log.Errorf("invalid --log-settings: %s", err.Error())
			return
		}
	}
	if len(_rtbfabricAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _rtbfabricAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}

	if resp, err := client.AcceptLink(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an inbound external link.
func rtbfabric_CreateInboundExternalLink(cfg aws.Config, client *rtbfabric.Client) {
	input := &rtbfabric.CreateInboundExternalLinkInput{
		// ClientToken: *string, // Required
		// GatewayId: *string, // Required
		// LogSettings: *types.LinkLogSettings, // Required
	}

	if len(_rtbfabricClientToken) > 0 {
		input.ClientToken = aws.String(_rtbfabricClientToken)
	}
	if len(_rtbfabricGatewayId) > 0 {
		input.GatewayId = aws.String(_rtbfabricGatewayId)
	}
	if len(_rtbfabricLogSettings) > 0 {
		if err := assignInputField(input, "LogSettings", _rtbfabricLogSettings); err != nil {
			log.Errorf("invalid --log-settings: %s", err.Error())
			return
		}
	}
	if len(_rtbfabricAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _rtbfabricAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_rtbfabricTags) > 0 {
		if err := assignInputField(input, "Tags", _rtbfabricTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateInboundExternalLink(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new link between gateways.
// Establishes a connection that allows gateways to communicate and exchange bid
// requests and responses.
func rtbfabric_CreateLink(cfg aws.Config, client *rtbfabric.Client) {
	input := &rtbfabric.CreateLinkInput{
		// GatewayId: *string, // Required
		// LogSettings: *types.LinkLogSettings, // Required
		// PeerGatewayId: *string, // Required
	}

	if len(_rtbfabricGatewayId) > 0 {
		input.GatewayId = aws.String(_rtbfabricGatewayId)
	}
	if len(_rtbfabricLogSettings) > 0 {
		if err := assignInputField(input, "LogSettings", _rtbfabricLogSettings); err != nil {
			log.Errorf("invalid --log-settings: %s", err.Error())
			return
		}
	}
	if len(_rtbfabricPeerGatewayId) > 0 {
		input.PeerGatewayId = aws.String(_rtbfabricPeerGatewayId)
	}
	if len(_rtbfabricAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _rtbfabricAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_rtbfabricHttpResponderAllowed) > 0 {
		if err := assignInputField(input, "HttpResponderAllowed", _rtbfabricHttpResponderAllowed); err != nil {
			log.Errorf("invalid --http-responder-allowed: %s", err.Error())
			return
		}
	}
	if len(_rtbfabricTags) > 0 {
		if err := assignInputField(input, "Tags", _rtbfabricTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateLink(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an outbound external link.
func rtbfabric_CreateOutboundExternalLink(cfg aws.Config, client *rtbfabric.Client) {
	input := &rtbfabric.CreateOutboundExternalLinkInput{
		// ClientToken: *string, // Required
		// GatewayId: *string, // Required
		// LogSettings: *types.LinkLogSettings, // Required
		// PublicEndpoint: *string, // Required
	}

	if len(_rtbfabricClientToken) > 0 {
		input.ClientToken = aws.String(_rtbfabricClientToken)
	}
	if len(_rtbfabricGatewayId) > 0 {
		input.GatewayId = aws.String(_rtbfabricGatewayId)
	}
	if len(_rtbfabricLogSettings) > 0 {
		if err := assignInputField(input, "LogSettings", _rtbfabricLogSettings); err != nil {
			log.Errorf("invalid --log-settings: %s", err.Error())
			return
		}
	}
	if len(_rtbfabricPublicEndpoint) > 0 {
		input.PublicEndpoint = aws.String(_rtbfabricPublicEndpoint)
	}
	if len(_rtbfabricAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _rtbfabricAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_rtbfabricTags) > 0 {
		if err := assignInputField(input, "Tags", _rtbfabricTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateOutboundExternalLink(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a requester gateway.
func rtbfabric_CreateRequesterGateway(cfg aws.Config, client *rtbfabric.Client) {
	input := &rtbfabric.CreateRequesterGatewayInput{
		// ClientToken: *string, // Required
		// SecurityGroupIds: []string, // Required
		// SubnetIds: []string, // Required
		// VpcId: *string, // Required
	}

	if len(_rtbfabricClientToken) > 0 {
		input.ClientToken = aws.String(_rtbfabricClientToken)
	}
	if len(_rtbfabricSecurityGroupIds) > 0 {
		input.SecurityGroupIds = append([]string(nil), _rtbfabricSecurityGroupIds...)
	}
	if len(_rtbfabricSubnetIds) > 0 {
		input.SubnetIds = append([]string(nil), _rtbfabricSubnetIds...)
	}
	if len(_rtbfabricVpcId) > 0 {
		input.VpcId = aws.String(_rtbfabricVpcId)
	}
	if len(_rtbfabricDescription) > 0 {
		input.Description = aws.String(_rtbfabricDescription)
	}
	if len(_rtbfabricTags) > 0 {
		if err := assignInputField(input, "Tags", _rtbfabricTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRequesterGateway(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a responder gateway.
// A domain name or managed endpoint is required.
func rtbfabric_CreateResponderGateway(cfg aws.Config, client *rtbfabric.Client) {
	input := &rtbfabric.CreateResponderGatewayInput{
		// ClientToken: *string, // Required
		// Port: *int32, // Required
		// Protocol: types.Protocol, // Required
		// SecurityGroupIds: []string, // Required
		// SubnetIds: []string, // Required
		// VpcId: *string, // Required
	}

	if len(_rtbfabricClientToken) > 0 {
		input.ClientToken = aws.String(_rtbfabricClientToken)
	}
	if len(_rtbfabricPort) > 0 {
		if err := assignInputField(input, "Port", _rtbfabricPort); err != nil {
			log.Errorf("invalid --port: %s", err.Error())
			return
		}
	}
	if len(_rtbfabricProtocol) > 0 {
		if err := assignInputField(input, "Protocol", _rtbfabricProtocol); err != nil {
			log.Errorf("invalid --protocol: %s", err.Error())
			return
		}
	}
	if len(_rtbfabricSecurityGroupIds) > 0 {
		input.SecurityGroupIds = append([]string(nil), _rtbfabricSecurityGroupIds...)
	}
	if len(_rtbfabricSubnetIds) > 0 {
		input.SubnetIds = append([]string(nil), _rtbfabricSubnetIds...)
	}
	if len(_rtbfabricVpcId) > 0 {
		input.VpcId = aws.String(_rtbfabricVpcId)
	}
	if len(_rtbfabricDescription) > 0 {
		input.Description = aws.String(_rtbfabricDescription)
	}
	if len(_rtbfabricDomainName) > 0 {
		input.DomainName = aws.String(_rtbfabricDomainName)
	}
	if len(_rtbfabricManagedEndpointConfiguration) > 0 {
		if err := assignInputField(input, "ManagedEndpointConfiguration", _rtbfabricManagedEndpointConfiguration); err != nil {
			log.Errorf("invalid --managed-endpoint-configuration: %s", err.Error())
			return
		}
	}
	if len(_rtbfabricTags) > 0 {
		if err := assignInputField(input, "Tags", _rtbfabricTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_rtbfabricTrustStoreConfiguration) > 0 {
		if err := assignInputField(input, "TrustStoreConfiguration", _rtbfabricTrustStoreConfiguration); err != nil {
			log.Errorf("invalid --trust-store-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateResponderGateway(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an inbound external link.
func rtbfabric_DeleteInboundExternalLink(cfg aws.Config, client *rtbfabric.Client) {
	input := &rtbfabric.DeleteInboundExternalLinkInput{
		// GatewayId: *string, // Required
		// LinkId: *string, // Required
	}

	if len(_rtbfabricGatewayId) > 0 {
		input.GatewayId = aws.String(_rtbfabricGatewayId)
	}
	if len(_rtbfabricLinkId) > 0 {
		input.LinkId = aws.String(_rtbfabricLinkId)
	}

	if resp, err := client.DeleteInboundExternalLink(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a link between gateways.
// Permanently removes the connection between gateways. This action cannot be
// undone.
func rtbfabric_DeleteLink(cfg aws.Config, client *rtbfabric.Client) {
	input := &rtbfabric.DeleteLinkInput{
		// GatewayId: *string, // Required
		// LinkId: *string, // Required
	}

	if len(_rtbfabricGatewayId) > 0 {
		input.GatewayId = aws.String(_rtbfabricGatewayId)
	}
	if len(_rtbfabricLinkId) > 0 {
		input.LinkId = aws.String(_rtbfabricLinkId)
	}

	if resp, err := client.DeleteLink(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an outbound external link.
func rtbfabric_DeleteOutboundExternalLink(cfg aws.Config, client *rtbfabric.Client) {
	input := &rtbfabric.DeleteOutboundExternalLinkInput{
		// GatewayId: *string, // Required
		// LinkId: *string, // Required
	}

	if len(_rtbfabricGatewayId) > 0 {
		input.GatewayId = aws.String(_rtbfabricGatewayId)
	}
	if len(_rtbfabricLinkId) > 0 {
		input.LinkId = aws.String(_rtbfabricLinkId)
	}

	if resp, err := client.DeleteOutboundExternalLink(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a requester gateway.
func rtbfabric_DeleteRequesterGateway(cfg aws.Config, client *rtbfabric.Client) {
	input := &rtbfabric.DeleteRequesterGatewayInput{
		// GatewayId: *string, // Required
	}

	if len(_rtbfabricGatewayId) > 0 {
		input.GatewayId = aws.String(_rtbfabricGatewayId)
	}

	if resp, err := client.DeleteRequesterGateway(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a responder gateway.
func rtbfabric_DeleteResponderGateway(cfg aws.Config, client *rtbfabric.Client) {
	input := &rtbfabric.DeleteResponderGatewayInput{
		// GatewayId: *string, // Required
	}

	if len(_rtbfabricGatewayId) > 0 {
		input.GatewayId = aws.String(_rtbfabricGatewayId)
	}

	if resp, err := client.DeleteResponderGateway(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about an inbound external link.
func rtbfabric_GetInboundExternalLink(cfg aws.Config, client *rtbfabric.Client) {
	input := &rtbfabric.GetInboundExternalLinkInput{
		// GatewayId: *string, // Required
		// LinkId: *string, // Required
	}

	if len(_rtbfabricGatewayId) > 0 {
		input.GatewayId = aws.String(_rtbfabricGatewayId)
	}
	if len(_rtbfabricLinkId) > 0 {
		input.LinkId = aws.String(_rtbfabricLinkId)
	}

	if resp, err := client.GetInboundExternalLink(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a link between gateways.
// Returns detailed information about the link configuration, status, and
// associated gateways.
func rtbfabric_GetLink(cfg aws.Config, client *rtbfabric.Client) {
	input := &rtbfabric.GetLinkInput{
		// GatewayId: *string, // Required
		// LinkId: *string, // Required
	}

	if len(_rtbfabricGatewayId) > 0 {
		input.GatewayId = aws.String(_rtbfabricGatewayId)
	}
	if len(_rtbfabricLinkId) > 0 {
		input.LinkId = aws.String(_rtbfabricLinkId)
	}

	if resp, err := client.GetLink(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about an outbound external link.
func rtbfabric_GetOutboundExternalLink(cfg aws.Config, client *rtbfabric.Client) {
	input := &rtbfabric.GetOutboundExternalLinkInput{
		// GatewayId: *string, // Required
		// LinkId: *string, // Required
	}

	if len(_rtbfabricGatewayId) > 0 {
		input.GatewayId = aws.String(_rtbfabricGatewayId)
	}
	if len(_rtbfabricLinkId) > 0 {
		input.LinkId = aws.String(_rtbfabricLinkId)
	}

	if resp, err := client.GetOutboundExternalLink(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a requester gateway.
func rtbfabric_GetRequesterGateway(cfg aws.Config, client *rtbfabric.Client) {
	input := &rtbfabric.GetRequesterGatewayInput{
		// GatewayId: *string, // Required
	}

	if len(_rtbfabricGatewayId) > 0 {
		input.GatewayId = aws.String(_rtbfabricGatewayId)
	}

	if resp, err := client.GetRequesterGateway(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a responder gateway.
func rtbfabric_GetResponderGateway(cfg aws.Config, client *rtbfabric.Client) {
	input := &rtbfabric.GetResponderGatewayInput{
		// GatewayId: *string, // Required
	}

	if len(_rtbfabricGatewayId) > 0 {
		input.GatewayId = aws.String(_rtbfabricGatewayId)
	}

	if resp, err := client.GetResponderGateway(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists links associated with gateways.
// Returns a list of all links for the specified gateways, including their status
// and configuration details.
func rtbfabric_ListLinks(cfg aws.Config, client *rtbfabric.Client) {
	input := &rtbfabric.ListLinksInput{
		// GatewayId: *string, // Required
	}

	if len(_rtbfabricGatewayId) > 0 {
		input.GatewayId = aws.String(_rtbfabricGatewayId)
	}
	if len(_rtbfabricMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _rtbfabricMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_rtbfabricNextToken) > 0 {
		input.NextToken = aws.String(_rtbfabricNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListLinks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rtbfabric.ListLinksOutput
	p := rtbfabric.NewListLinksPaginator(client, input)
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

// Lists requester gateways.
func rtbfabric_ListRequesterGateways(cfg aws.Config, client *rtbfabric.Client) {
	input := &rtbfabric.ListRequesterGatewaysInput{}

	if len(_rtbfabricMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _rtbfabricMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_rtbfabricNextToken) > 0 {
		input.NextToken = aws.String(_rtbfabricNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRequesterGateways(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rtbfabric.ListRequesterGatewaysOutput
	p := rtbfabric.NewListRequesterGatewaysPaginator(client, input)
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

// Lists reponder gateways.
func rtbfabric_ListResponderGateways(cfg aws.Config, client *rtbfabric.Client) {
	input := &rtbfabric.ListResponderGatewaysInput{}

	if len(_rtbfabricMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _rtbfabricMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_rtbfabricNextToken) > 0 {
		input.NextToken = aws.String(_rtbfabricNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListResponderGateways(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rtbfabric.ListResponderGatewaysOutput
	p := rtbfabric.NewListResponderGatewaysPaginator(client, input)
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

// Lists tags for a resource.
func rtbfabric_ListTagsForResource(cfg aws.Config, client *rtbfabric.Client) {
	input := &rtbfabric.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_rtbfabricResourceArn) > 0 {
		input.ResourceArn = aws.String(_rtbfabricResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Rejects a link request between gateways.
// When a requester gateway requests to link with a responder gateway, the
// responder can use this operation to decline the link request.
func rtbfabric_RejectLink(cfg aws.Config, client *rtbfabric.Client) {
	input := &rtbfabric.RejectLinkInput{
		// GatewayId: *string, // Required
		// LinkId: *string, // Required
	}

	if len(_rtbfabricGatewayId) > 0 {
		input.GatewayId = aws.String(_rtbfabricGatewayId)
	}
	if len(_rtbfabricLinkId) > 0 {
		input.LinkId = aws.String(_rtbfabricLinkId)
	}

	if resp, err := client.RejectLink(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assigns one or more tags (key-value pairs) to the specified resource.
func rtbfabric_TagResource(cfg aws.Config, client *rtbfabric.Client) {
	input := &rtbfabric.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_rtbfabricResourceArn) > 0 {
		input.ResourceArn = aws.String(_rtbfabricResourceArn)
	}
	if len(_rtbfabricTags) > 0 {
		if err := assignInputField(input, "Tags", _rtbfabricTags); err != nil {
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

// Removes a tag or tags from a resource.
func rtbfabric_UntagResource(cfg aws.Config, client *rtbfabric.Client) {
	input := &rtbfabric.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_rtbfabricResourceArn) > 0 {
		input.ResourceArn = aws.String(_rtbfabricResourceArn)
	}
	if len(_rtbfabricTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _rtbfabricTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration of a link between gateways.
// Allows you to modify settings and parameters for an existing link.
func rtbfabric_UpdateLink(cfg aws.Config, client *rtbfabric.Client) {
	input := &rtbfabric.UpdateLinkInput{
		// GatewayId: *string, // Required
		// LinkId: *string, // Required
	}

	if len(_rtbfabricGatewayId) > 0 {
		input.GatewayId = aws.String(_rtbfabricGatewayId)
	}
	if len(_rtbfabricLinkId) > 0 {
		input.LinkId = aws.String(_rtbfabricLinkId)
	}
	if len(_rtbfabricLogSettings) > 0 {
		if err := assignInputField(input, "LogSettings", _rtbfabricLogSettings); err != nil {
			log.Errorf("invalid --log-settings: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateLink(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a link module flow.
func rtbfabric_UpdateLinkModuleFlow(cfg aws.Config, client *rtbfabric.Client) {
	input := &rtbfabric.UpdateLinkModuleFlowInput{
		// ClientToken: *string, // Required
		// GatewayId: *string, // Required
		// LinkId: *string, // Required
		// Modules: []types.ModuleConfiguration, // Required
	}

	if len(_rtbfabricClientToken) > 0 {
		input.ClientToken = aws.String(_rtbfabricClientToken)
	}
	if len(_rtbfabricGatewayId) > 0 {
		input.GatewayId = aws.String(_rtbfabricGatewayId)
	}
	if len(_rtbfabricLinkId) > 0 {
		input.LinkId = aws.String(_rtbfabricLinkId)
	}
	if len(_rtbfabricModules) > 0 {
		if err := assignInputField(input, "Modules", _rtbfabricModules); err != nil {
			log.Errorf("invalid --modules: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateLinkModuleFlow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a requester gateway.
func rtbfabric_UpdateRequesterGateway(cfg aws.Config, client *rtbfabric.Client) {
	input := &rtbfabric.UpdateRequesterGatewayInput{
		// ClientToken: *string, // Required
		// GatewayId: *string, // Required
	}

	if len(_rtbfabricClientToken) > 0 {
		input.ClientToken = aws.String(_rtbfabricClientToken)
	}
	if len(_rtbfabricGatewayId) > 0 {
		input.GatewayId = aws.String(_rtbfabricGatewayId)
	}
	if len(_rtbfabricDescription) > 0 {
		input.Description = aws.String(_rtbfabricDescription)
	}

	if resp, err := client.UpdateRequesterGateway(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a responder gateway.
func rtbfabric_UpdateResponderGateway(cfg aws.Config, client *rtbfabric.Client) {
	input := &rtbfabric.UpdateResponderGatewayInput{
		// ClientToken: *string, // Required
		// GatewayId: *string, // Required
		// Port: *int32, // Required
		// Protocol: types.Protocol, // Required
	}

	if len(_rtbfabricClientToken) > 0 {
		input.ClientToken = aws.String(_rtbfabricClientToken)
	}
	if len(_rtbfabricGatewayId) > 0 {
		input.GatewayId = aws.String(_rtbfabricGatewayId)
	}
	if len(_rtbfabricPort) > 0 {
		if err := assignInputField(input, "Port", _rtbfabricPort); err != nil {
			log.Errorf("invalid --port: %s", err.Error())
			return
		}
	}
	if len(_rtbfabricProtocol) > 0 {
		if err := assignInputField(input, "Protocol", _rtbfabricProtocol); err != nil {
			log.Errorf("invalid --protocol: %s", err.Error())
			return
		}
	}
	if len(_rtbfabricDescription) > 0 {
		input.Description = aws.String(_rtbfabricDescription)
	}
	if len(_rtbfabricDomainName) > 0 {
		input.DomainName = aws.String(_rtbfabricDomainName)
	}
	if len(_rtbfabricManagedEndpointConfiguration) > 0 {
		if err := assignInputField(input, "ManagedEndpointConfiguration", _rtbfabricManagedEndpointConfiguration); err != nil {
			log.Errorf("invalid --managed-endpoint-configuration: %s", err.Error())
			return
		}
	}
	if len(_rtbfabricTrustStoreConfiguration) > 0 {
		if err := assignInputField(input, "TrustStoreConfiguration", _rtbfabricTrustStoreConfiguration); err != nil {
			log.Errorf("invalid --trust-store-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateResponderGateway(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_rtbfabricCmd)
	_rtbfabricCmd.Flags().SortFlags = false

	_rtbfabricCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_rtbfabricCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_rtbfabricCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_rtbfabricCmd.Flags().StringVarP(&_rtbfabricAttributes, "attributes", "", "", "Attributes")
	_rtbfabricCmd.Flags().StringVarP(&_rtbfabricClientToken, "client-token", "", "", "Client Token")
	_rtbfabricCmd.Flags().StringVarP(&_rtbfabricDescription, "description", "", "", "Description")
	_rtbfabricCmd.Flags().StringVarP(&_rtbfabricDomainName, "domain-name", "", "", "Domain Name")
	_rtbfabricCmd.Flags().StringVarP(&_rtbfabricGatewayId, "gateway-id", "", "", "Gateway ID")
	_rtbfabricCmd.Flags().StringVarP(&_rtbfabricHttpResponderAllowed, "http-responder-allowed", "", "", "HTTP Responder Allowed")
	_rtbfabricCmd.Flags().StringVarP(&_rtbfabricLinkId, "link-id", "", "", "Link ID")
	_rtbfabricCmd.Flags().StringVarP(&_rtbfabricLogSettings, "log-settings", "", "", "Log Settings")
	_rtbfabricCmd.Flags().StringVarP(&_rtbfabricManagedEndpointConfiguration, "managed-endpoint-configuration", "", "", "Managed Endpoint Configuration")
	_rtbfabricCmd.Flags().StringVarP(&_rtbfabricMaxResults, "max-results", "", "", "Max Results")
	_rtbfabricCmd.Flags().StringVarP(&_rtbfabricModules, "modules", "", "", "Modules")
	_rtbfabricCmd.Flags().StringVarP(&_rtbfabricNextToken, "next-token", "", "", "Next Token")
	_rtbfabricCmd.Flags().StringVarP(&_rtbfabricPeerGatewayId, "peer-gateway-id", "", "", "Peer Gateway ID")
	_rtbfabricCmd.Flags().StringVarP(&_rtbfabricPort, "port", "", "", "Port")
	_rtbfabricCmd.Flags().StringVarP(&_rtbfabricProtocol, "protocol", "", "", "Protocol")
	_rtbfabricCmd.Flags().StringVarP(&_rtbfabricPublicEndpoint, "public-endpoint", "", "", "Public Endpoint")
	_rtbfabricCmd.Flags().StringVarP(&_rtbfabricResourceArn, "resource-arn", "", "", "Resource ARN")
	_rtbfabricCmd.Flags().StringSliceVarP(&_rtbfabricSecurityGroupIds, "security-group-ids", "", nil, "Security Group Ids")
	_rtbfabricCmd.Flags().StringSliceVarP(&_rtbfabricSubnetIds, "subnet-ids", "", nil, "Subnet Ids")
	_rtbfabricCmd.Flags().StringSliceVarP(&_rtbfabricTagKeys, "tag-keys", "", nil, "Tag Keys")
	_rtbfabricCmd.Flags().StringVarP(&_rtbfabricTags, "tags", "", "", "Tags")
	_rtbfabricCmd.Flags().StringVarP(&_rtbfabricTrustStoreConfiguration, "trust-store-configuration", "", "", "Trust Store Configuration")
	_rtbfabricCmd.Flags().StringVarP(&_rtbfabricVpcId, "vpc-id", "", "", "VPC ID")

	_rtbfabricCmd.Flags().BoolVarP(&_rtbfabricAcceptLink, "accept-link", "", false, "Accept Link")
	_rtbfabricCmd.Flags().BoolVarP(&_rtbfabricCreateInboundExternalLink, "create-inbound-external-link", "", false, "Create Inbound External Link")
	_rtbfabricCmd.Flags().BoolVarP(&_rtbfabricCreateLink, "create-link", "", false, "Create Link")
	_rtbfabricCmd.Flags().BoolVarP(&_rtbfabricCreateOutboundExternalLink, "create-outbound-external-link", "", false, "Create Outbound External Link")
	_rtbfabricCmd.Flags().BoolVarP(&_rtbfabricCreateRequesterGateway, "create-requester-gateway", "", false, "Create Requester Gateway")
	_rtbfabricCmd.Flags().BoolVarP(&_rtbfabricCreateResponderGateway, "create-responder-gateway", "", false, "Create Responder Gateway")
	_rtbfabricCmd.Flags().BoolVarP(&_rtbfabricDeleteInboundExternalLink, "delete-inbound-external-link", "", false, "Delete Inbound External Link")
	_rtbfabricCmd.Flags().BoolVarP(&_rtbfabricDeleteLink, "delete-link", "", false, "Delete Link")
	_rtbfabricCmd.Flags().BoolVarP(&_rtbfabricDeleteOutboundExternalLink, "delete-outbound-external-link", "", false, "Delete Outbound External Link")
	_rtbfabricCmd.Flags().BoolVarP(&_rtbfabricDeleteRequesterGateway, "delete-requester-gateway", "", false, "Delete Requester Gateway")
	_rtbfabricCmd.Flags().BoolVarP(&_rtbfabricDeleteResponderGateway, "delete-responder-gateway", "", false, "Delete Responder Gateway")
	_rtbfabricCmd.Flags().BoolVarP(&_rtbfabricGetInboundExternalLink, "get-inbound-external-link", "", false, "Get Inbound External Link")
	_rtbfabricCmd.Flags().BoolVarP(&_rtbfabricGetLink, "get-link", "", false, "Get Link")
	_rtbfabricCmd.Flags().BoolVarP(&_rtbfabricGetOutboundExternalLink, "get-outbound-external-link", "", false, "Get Outbound External Link")
	_rtbfabricCmd.Flags().BoolVarP(&_rtbfabricGetRequesterGateway, "get-requester-gateway", "", false, "Get Requester Gateway")
	_rtbfabricCmd.Flags().BoolVarP(&_rtbfabricGetResponderGateway, "get-responder-gateway", "", false, "Get Responder Gateway")
	_rtbfabricCmd.Flags().BoolVarP(&_rtbfabricListLinks, "list-links", "", false, "List Links")
	_rtbfabricCmd.Flags().BoolVarP(&_rtbfabricListRequesterGateways, "list-requester-gateways", "", false, "List Requester Gateways")
	_rtbfabricCmd.Flags().BoolVarP(&_rtbfabricListResponderGateways, "list-responder-gateways", "", false, "List Responder Gateways")
	_rtbfabricCmd.Flags().BoolVarP(&_rtbfabricListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_rtbfabricCmd.Flags().BoolVarP(&_rtbfabricRejectLink, "reject-link", "", false, "Reject Link")
	_rtbfabricCmd.Flags().BoolVarP(&_rtbfabricTagResource, "tag-resource", "", false, "Tag Resource")
	_rtbfabricCmd.Flags().BoolVarP(&_rtbfabricUntagResource, "untag-resource", "", false, "Untag Resource")
	_rtbfabricCmd.Flags().BoolVarP(&_rtbfabricUpdateLink, "update-link", "", false, "Update Link")
	_rtbfabricCmd.Flags().BoolVarP(&_rtbfabricUpdateLinkModuleFlow, "update-link-module-flow", "", false, "Update Link Module Flow")
	_rtbfabricCmd.Flags().BoolVarP(&_rtbfabricUpdateRequesterGateway, "update-requester-gateway", "", false, "Update Requester Gateway")
	_rtbfabricCmd.Flags().BoolVarP(&_rtbfabricUpdateResponderGateway, "update-responder-gateway", "", false, "Update Responder Gateway")

}
