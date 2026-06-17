package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// globalacceleratorCmd represents the globalaccelerator command
var _globalacceleratorCmd = &cobra.Command{
	Use:   "globalaccelerator",
	Short: "AWS globalaccelerator CLI",
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
		client := globalaccelerator.NewFromConfig(cfg)
		if _globalacceleratorAddCustomRoutingEndpoints {
			globalaccelerator_AddCustomRoutingEndpoints(cfg, client)
			return
		}
		if _globalacceleratorAddEndpoints {
			globalaccelerator_AddEndpoints(cfg, client)
			return
		}
		if _globalacceleratorAdvertiseByoipCidr {
			globalaccelerator_AdvertiseByoipCidr(cfg, client)
			return
		}
		if _globalacceleratorAllowCustomRoutingTraffic {
			globalaccelerator_AllowCustomRoutingTraffic(cfg, client)
			return
		}
		if _globalacceleratorCreateAccelerator {
			globalaccelerator_CreateAccelerator(cfg, client)
			return
		}
		if _globalacceleratorCreateCrossAccountAttachment {
			globalaccelerator_CreateCrossAccountAttachment(cfg, client)
			return
		}
		if _globalacceleratorCreateCustomRoutingAccelerator {
			globalaccelerator_CreateCustomRoutingAccelerator(cfg, client)
			return
		}
		if _globalacceleratorCreateCustomRoutingEndpointGroup {
			globalaccelerator_CreateCustomRoutingEndpointGroup(cfg, client)
			return
		}
		if _globalacceleratorCreateCustomRoutingListener {
			globalaccelerator_CreateCustomRoutingListener(cfg, client)
			return
		}
		if _globalacceleratorCreateEndpointGroup {
			globalaccelerator_CreateEndpointGroup(cfg, client)
			return
		}
		if _globalacceleratorCreateListener {
			globalaccelerator_CreateListener(cfg, client)
			return
		}
		if _globalacceleratorDeleteAccelerator {
			globalaccelerator_DeleteAccelerator(cfg, client)
			return
		}
		if _globalacceleratorDeleteCrossAccountAttachment {
			globalaccelerator_DeleteCrossAccountAttachment(cfg, client)
			return
		}
		if _globalacceleratorDeleteCustomRoutingAccelerator {
			globalaccelerator_DeleteCustomRoutingAccelerator(cfg, client)
			return
		}
		if _globalacceleratorDeleteCustomRoutingEndpointGroup {
			globalaccelerator_DeleteCustomRoutingEndpointGroup(cfg, client)
			return
		}
		if _globalacceleratorDeleteCustomRoutingListener {
			globalaccelerator_DeleteCustomRoutingListener(cfg, client)
			return
		}
		if _globalacceleratorDeleteEndpointGroup {
			globalaccelerator_DeleteEndpointGroup(cfg, client)
			return
		}
		if _globalacceleratorDeleteListener {
			globalaccelerator_DeleteListener(cfg, client)
			return
		}
		if _globalacceleratorDenyCustomRoutingTraffic {
			globalaccelerator_DenyCustomRoutingTraffic(cfg, client)
			return
		}
		if _globalacceleratorDeprovisionByoipCidr {
			globalaccelerator_DeprovisionByoipCidr(cfg, client)
			return
		}
		if _globalacceleratorDescribeAccelerator {
			globalaccelerator_DescribeAccelerator(cfg, client)
			return
		}
		if _globalacceleratorDescribeAcceleratorAttributes {
			globalaccelerator_DescribeAcceleratorAttributes(cfg, client)
			return
		}
		if _globalacceleratorDescribeCrossAccountAttachment {
			globalaccelerator_DescribeCrossAccountAttachment(cfg, client)
			return
		}
		if _globalacceleratorDescribeCustomRoutingAccelerator {
			globalaccelerator_DescribeCustomRoutingAccelerator(cfg, client)
			return
		}
		if _globalacceleratorDescribeCustomRoutingAcceleratorAttributes {
			globalaccelerator_DescribeCustomRoutingAcceleratorAttributes(cfg, client)
			return
		}
		if _globalacceleratorDescribeCustomRoutingEndpointGroup {
			globalaccelerator_DescribeCustomRoutingEndpointGroup(cfg, client)
			return
		}
		if _globalacceleratorDescribeCustomRoutingListener {
			globalaccelerator_DescribeCustomRoutingListener(cfg, client)
			return
		}
		if _globalacceleratorDescribeEndpointGroup {
			globalaccelerator_DescribeEndpointGroup(cfg, client)
			return
		}
		if _globalacceleratorDescribeListener {
			globalaccelerator_DescribeListener(cfg, client)
			return
		}
		if _globalacceleratorListAccelerators {
			globalaccelerator_ListAccelerators(cfg, client)
			return
		}
		if _globalacceleratorListByoipCidrs {
			globalaccelerator_ListByoipCidrs(cfg, client)
			return
		}
		if _globalacceleratorListCrossAccountAttachments {
			globalaccelerator_ListCrossAccountAttachments(cfg, client)
			return
		}
		if _globalacceleratorListCrossAccountResourceAccounts {
			globalaccelerator_ListCrossAccountResourceAccounts(cfg, client)
			return
		}
		if _globalacceleratorListCrossAccountResources {
			globalaccelerator_ListCrossAccountResources(cfg, client)
			return
		}
		if _globalacceleratorListCustomRoutingAccelerators {
			globalaccelerator_ListCustomRoutingAccelerators(cfg, client)
			return
		}
		if _globalacceleratorListCustomRoutingEndpointGroups {
			globalaccelerator_ListCustomRoutingEndpointGroups(cfg, client)
			return
		}
		if _globalacceleratorListCustomRoutingListeners {
			globalaccelerator_ListCustomRoutingListeners(cfg, client)
			return
		}
		if _globalacceleratorListCustomRoutingPortMappings {
			globalaccelerator_ListCustomRoutingPortMappings(cfg, client)
			return
		}
		if _globalacceleratorListCustomRoutingPortMappingsByDestination {
			globalaccelerator_ListCustomRoutingPortMappingsByDestination(cfg, client)
			return
		}
		if _globalacceleratorListEndpointGroups {
			globalaccelerator_ListEndpointGroups(cfg, client)
			return
		}
		if _globalacceleratorListListeners {
			globalaccelerator_ListListeners(cfg, client)
			return
		}
		if _globalacceleratorListTagsForResource {
			globalaccelerator_ListTagsForResource(cfg, client)
			return
		}
		if _globalacceleratorProvisionByoipCidr {
			globalaccelerator_ProvisionByoipCidr(cfg, client)
			return
		}
		if _globalacceleratorRemoveCustomRoutingEndpoints {
			globalaccelerator_RemoveCustomRoutingEndpoints(cfg, client)
			return
		}
		if _globalacceleratorRemoveEndpoints {
			globalaccelerator_RemoveEndpoints(cfg, client)
			return
		}
		if _globalacceleratorTagResource {
			globalaccelerator_TagResource(cfg, client)
			return
		}
		if _globalacceleratorUntagResource {
			globalaccelerator_UntagResource(cfg, client)
			return
		}
		if _globalacceleratorUpdateAccelerator {
			globalaccelerator_UpdateAccelerator(cfg, client)
			return
		}
		if _globalacceleratorUpdateAcceleratorAttributes {
			globalaccelerator_UpdateAcceleratorAttributes(cfg, client)
			return
		}
		if _globalacceleratorUpdateCrossAccountAttachment {
			globalaccelerator_UpdateCrossAccountAttachment(cfg, client)
			return
		}
		if _globalacceleratorUpdateCustomRoutingAccelerator {
			globalaccelerator_UpdateCustomRoutingAccelerator(cfg, client)
			return
		}
		if _globalacceleratorUpdateCustomRoutingAcceleratorAttributes {
			globalaccelerator_UpdateCustomRoutingAcceleratorAttributes(cfg, client)
			return
		}
		if _globalacceleratorUpdateCustomRoutingListener {
			globalaccelerator_UpdateCustomRoutingListener(cfg, client)
			return
		}
		if _globalacceleratorUpdateEndpointGroup {
			globalaccelerator_UpdateEndpointGroup(cfg, client)
			return
		}
		if _globalacceleratorUpdateListener {
			globalaccelerator_UpdateListener(cfg, client)
			return
		}
		if _globalacceleratorWithdrawByoipCidr {
			globalaccelerator_WithdrawByoipCidr(cfg, client)
			return
		}

	},
}

var (
	_globalacceleratorAddCustomRoutingEndpoints                  bool
	_globalacceleratorAddEndpoints                               bool
	_globalacceleratorAdvertiseByoipCidr                         bool
	_globalacceleratorAllowCustomRoutingTraffic                  bool
	_globalacceleratorCreateAccelerator                          bool
	_globalacceleratorCreateCrossAccountAttachment               bool
	_globalacceleratorCreateCustomRoutingAccelerator             bool
	_globalacceleratorCreateCustomRoutingEndpointGroup           bool
	_globalacceleratorCreateCustomRoutingListener                bool
	_globalacceleratorCreateEndpointGroup                        bool
	_globalacceleratorCreateListener                             bool
	_globalacceleratorDeleteAccelerator                          bool
	_globalacceleratorDeleteCrossAccountAttachment               bool
	_globalacceleratorDeleteCustomRoutingAccelerator             bool
	_globalacceleratorDeleteCustomRoutingEndpointGroup           bool
	_globalacceleratorDeleteCustomRoutingListener                bool
	_globalacceleratorDeleteEndpointGroup                        bool
	_globalacceleratorDeleteListener                             bool
	_globalacceleratorDenyCustomRoutingTraffic                   bool
	_globalacceleratorDeprovisionByoipCidr                       bool
	_globalacceleratorDescribeAccelerator                        bool
	_globalacceleratorDescribeAcceleratorAttributes              bool
	_globalacceleratorDescribeCrossAccountAttachment             bool
	_globalacceleratorDescribeCustomRoutingAccelerator           bool
	_globalacceleratorDescribeCustomRoutingAcceleratorAttributes bool
	_globalacceleratorDescribeCustomRoutingEndpointGroup         bool
	_globalacceleratorDescribeCustomRoutingListener              bool
	_globalacceleratorDescribeEndpointGroup                      bool
	_globalacceleratorDescribeListener                           bool
	_globalacceleratorListAccelerators                           bool
	_globalacceleratorListByoipCidrs                             bool
	_globalacceleratorListCrossAccountAttachments                bool
	_globalacceleratorListCrossAccountResourceAccounts           bool
	_globalacceleratorListCrossAccountResources                  bool
	_globalacceleratorListCustomRoutingAccelerators              bool
	_globalacceleratorListCustomRoutingEndpointGroups            bool
	_globalacceleratorListCustomRoutingListeners                 bool
	_globalacceleratorListCustomRoutingPortMappings              bool
	_globalacceleratorListCustomRoutingPortMappingsByDestination bool
	_globalacceleratorListEndpointGroups                         bool
	_globalacceleratorListListeners                              bool
	_globalacceleratorListTagsForResource                        bool
	_globalacceleratorProvisionByoipCidr                         bool
	_globalacceleratorRemoveCustomRoutingEndpoints               bool
	_globalacceleratorRemoveEndpoints                            bool
	_globalacceleratorTagResource                                bool
	_globalacceleratorUntagResource                              bool
	_globalacceleratorUpdateAccelerator                          bool
	_globalacceleratorUpdateAcceleratorAttributes                bool
	_globalacceleratorUpdateCrossAccountAttachment               bool
	_globalacceleratorUpdateCustomRoutingAccelerator             bool
	_globalacceleratorUpdateCustomRoutingAcceleratorAttributes   bool
	_globalacceleratorUpdateCustomRoutingListener                bool
	_globalacceleratorUpdateEndpointGroup                        bool
	_globalacceleratorUpdateListener                             bool
	_globalacceleratorWithdrawByoipCidr                          bool

	_globalacceleratorAcceleratorArn             string
	_globalacceleratorAddPrincipals              []string
	_globalacceleratorAddResources               string
	_globalacceleratorAllowAllTrafficToEndpoint  string
	_globalacceleratorAttachmentArn              string
	_globalacceleratorCidr                       string
	_globalacceleratorCidrAuthorizationContext   string
	_globalacceleratorClientAffinity             string
	_globalacceleratorDenyAllTrafficToEndpoint   string
	_globalacceleratorDestinationAddress         string
	_globalacceleratorDestinationAddresses       []string
	_globalacceleratorDestinationConfigurations  string
	_globalacceleratorDestinationPorts           string
	_globalacceleratorEnabled                    string
	_globalacceleratorEndpointConfigurations     string
	_globalacceleratorEndpointGroupArn           string
	_globalacceleratorEndpointGroupRegion        string
	_globalacceleratorEndpointId                 string
	_globalacceleratorEndpointIdentifiers        string
	_globalacceleratorEndpointIds                []string
	_globalacceleratorFlowLogsEnabled            string
	_globalacceleratorFlowLogsS3Bucket           string
	_globalacceleratorFlowLogsS3Prefix           string
	_globalacceleratorHealthCheckIntervalSeconds string
	_globalacceleratorHealthCheckPath            string
	_globalacceleratorHealthCheckPort            string
	_globalacceleratorHealthCheckProtocol        string
	_globalacceleratorIdempotencyToken           string
	_globalacceleratorIpAddressType              string
	_globalacceleratorIpAddresses                []string
	_globalacceleratorListenerArn                string
	_globalacceleratorMaxResults                 string
	_globalacceleratorName                       string
	_globalacceleratorNextToken                  string
	_globalacceleratorPortOverrides              string
	_globalacceleratorPortRanges                 string
	_globalacceleratorPrincipals                 []string
	_globalacceleratorProtocol                   string
	_globalacceleratorRemovePrincipals           []string
	_globalacceleratorRemoveResources            string
	_globalacceleratorResourceArn                string
	_globalacceleratorResourceOwnerAwsAccountId  string
	_globalacceleratorResources                  string
	_globalacceleratorTagKeys                    []string
	_globalacceleratorTags                       string
	_globalacceleratorThresholdCount             string
	_globalacceleratorTrafficDialPercentage      string
)

// Associate a virtual private cloud (VPC) subnet endpoint with your custom
// routing accelerator.
//
// The listener port range must be large enough to support the number of IP
// addresses that can be specified in your subnet. The number of ports required is:
// subnet size times the number of ports per destination EC2 instances. For
// example, a subnet defined as /24 requires a listener port range of at least 255
// ports.
//
// Note: You must have enough remaining listener ports available to map to the
// subnet ports, or the call will fail with a LimitExceededException.
//
// By default, all destinations in a subnet in a custom routing accelerator cannot
// receive traffic. To enable all destinations to receive traffic, or to specify
// individual port mappings that can receive traffic, see the [AllowCustomRoutingTraffic]operation.
//
// [AllowCustomRoutingTraffic]: https://docs.aws.amazon.com/global-accelerator/latest/api/API_AllowCustomRoutingTraffic.html
func globalaccelerator_AddCustomRoutingEndpoints(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.AddCustomRoutingEndpointsInput{
		// EndpointConfigurations: []types.CustomRoutingEndpointConfiguration, // Required
		// EndpointGroupArn: *string, // Required
	}

	if len(_globalacceleratorEndpointConfigurations) > 0 {
		if err := assignInputField(input, "EndpointConfigurations", _globalacceleratorEndpointConfigurations); err != nil {
			log.Errorf("invalid --endpoint-configurations: %s", err.Error())
			return
		}
	}
	if len(_globalacceleratorEndpointGroupArn) > 0 {
		input.EndpointGroupArn = aws.String(_globalacceleratorEndpointGroupArn)
	}

	if resp, err := client.AddCustomRoutingEndpoints(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Add endpoints to an endpoint group. The AddEndpoints API operation is the
// recommended option for adding endpoints. The alternative options are to add
// endpoints when you create an endpoint group (with the [CreateEndpointGroup]API) or when you update
// an endpoint group (with the [UpdateEndpointGroup]API).
//
// There are two advantages to using AddEndpoints to add endpoints in Global
// Accelerator:
//
// - It's faster, because Global Accelerator only has to resolve the new
// endpoints that you're adding, rather than resolving new and existing endpoints.
//
// - It's more convenient, because you don't need to specify the current
// endpoints that are already in the endpoint group, in addition to the new
// endpoints that you want to add.
//
// For information about endpoint types and requirements for endpoints that you
// can add to Global Accelerator, see [Endpoints for standard accelerators]in the Global Accelerator Developer Guide.
//
// [Endpoints for standard accelerators]: https://docs.aws.amazon.com/global-accelerator/latest/dg/about-endpoints.html
// [UpdateEndpointGroup]: https://docs.aws.amazon.com/global-accelerator/latest/api/API_UpdateEndpointGroup.html
// [CreateEndpointGroup]: https://docs.aws.amazon.com/global-accelerator/latest/api/API_CreateEndpointGroup.html
func globalaccelerator_AddEndpoints(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.AddEndpointsInput{
		// EndpointConfigurations: []types.EndpointConfiguration, // Required
		// EndpointGroupArn: *string, // Required
	}

	if len(_globalacceleratorEndpointConfigurations) > 0 {
		if err := assignInputField(input, "EndpointConfigurations", _globalacceleratorEndpointConfigurations); err != nil {
			log.Errorf("invalid --endpoint-configurations: %s", err.Error())
			return
		}
	}
	if len(_globalacceleratorEndpointGroupArn) > 0 {
		input.EndpointGroupArn = aws.String(_globalacceleratorEndpointGroupArn)
	}

	if resp, err := client.AddEndpoints(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Advertises an IPv4 address range that is provisioned for use with your Amazon
// Web Services resources through bring your own IP addresses (BYOIP). It can take
// a few minutes before traffic to the specified addresses starts routing to Amazon
// Web Services because of propagation delays.
//
// To stop advertising the BYOIP address range, use [WithdrawByoipCidr].
//
// For more information, see [Bring your own IP addresses (BYOIP)] in the Global Accelerator Developer Guide.
//
// [Bring your own IP addresses (BYOIP)]: https://docs.aws.amazon.com/global-accelerator/latest/dg/using-byoip.html
// [WithdrawByoipCidr]: https://docs.aws.amazon.com/global-accelerator/latest/api/WithdrawByoipCidr.html
func globalaccelerator_AdvertiseByoipCidr(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.AdvertiseByoipCidrInput{
		// Cidr: *string, // Required
	}

	if len(_globalacceleratorCidr) > 0 {
		input.Cidr = aws.String(_globalacceleratorCidr)
	}

	if resp, err := client.AdvertiseByoipCidr(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Specify the Amazon EC2 instance (destination) IP addresses and ports for a VPC
// subnet endpoint that can receive traffic for a custom routing accelerator. You
// can allow traffic to all destinations in the subnet endpoint, or allow traffic
// to a specified list of destination IP addresses and ports in the subnet. Note
// that you cannot specify IP addresses or ports outside of the range that you
// configured for the endpoint group.
//
// After you make changes, you can verify that the updates are complete by
// checking the status of your accelerator: the status changes from IN_PROGRESS to
// DEPLOYED.
func globalaccelerator_AllowCustomRoutingTraffic(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.AllowCustomRoutingTrafficInput{
		// EndpointGroupArn: *string, // Required
		// EndpointId: *string, // Required
	}

	if len(_globalacceleratorEndpointGroupArn) > 0 {
		input.EndpointGroupArn = aws.String(_globalacceleratorEndpointGroupArn)
	}
	if len(_globalacceleratorEndpointId) > 0 {
		input.EndpointId = aws.String(_globalacceleratorEndpointId)
	}
	if len(_globalacceleratorAllowAllTrafficToEndpoint) > 0 {
		if err := assignInputField(input, "AllowAllTrafficToEndpoint", _globalacceleratorAllowAllTrafficToEndpoint); err != nil {
			log.Errorf("invalid --allow-all-traffic-to-endpoint: %s", err.Error())
			return
		}
	}
	if len(_globalacceleratorDestinationAddresses) > 0 {
		input.DestinationAddresses = append([]string(nil), _globalacceleratorDestinationAddresses...)
	}
	if len(_globalacceleratorDestinationPorts) > 0 {
		if err := assignInputField(input, "DestinationPorts", _globalacceleratorDestinationPorts); err != nil {
			log.Errorf("invalid --destination-ports: %s", err.Error())
			return
		}
	}

	if resp, err := client.AllowCustomRoutingTraffic(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create an accelerator. An accelerator includes one or more listeners that
// process inbound connections and direct traffic to one or more endpoint groups,
// each of which includes endpoints, such as Network Load Balancers.
//
// Global Accelerator is a global service that supports endpoints in multiple
// Amazon Web Services Regions but you must specify the US West (Oregon) Region to
// create, update, or otherwise work with accelerators. That is, for example,
// specify --region us-west-2 on Amazon Web Services CLI commands.
func globalaccelerator_CreateAccelerator(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.CreateAcceleratorInput{
		// IdempotencyToken: *string, // Required
		// Name: *string, // Required
	}

	if len(_globalacceleratorIdempotencyToken) > 0 {
		input.IdempotencyToken = aws.String(_globalacceleratorIdempotencyToken)
	}
	if len(_globalacceleratorName) > 0 {
		input.Name = aws.String(_globalacceleratorName)
	}
	if len(_globalacceleratorEnabled) > 0 {
		if err := assignInputField(input, "Enabled", _globalacceleratorEnabled); err != nil {
			log.Errorf("invalid --enabled: %s", err.Error())
			return
		}
	}
	if len(_globalacceleratorIpAddressType) > 0 {
		if err := assignInputField(input, "IpAddressType", _globalacceleratorIpAddressType); err != nil {
			log.Errorf("invalid --ip-address-type: %s", err.Error())
			return
		}
	}
	if len(_globalacceleratorIpAddresses) > 0 {
		input.IpAddresses = append([]string(nil), _globalacceleratorIpAddresses...)
	}
	if len(_globalacceleratorTags) > 0 {
		if err := assignInputField(input, "Tags", _globalacceleratorTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAccelerator(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a cross-account attachment in Global Accelerator. You create a
// cross-account attachment to specify the principals who have permission to work
// with resources in accelerators in their own account. You specify, in the same
// attachment, the resources that are shared.
//
// A principal can be an Amazon Web Services account number or the Amazon Resource
// Name (ARN) for an accelerator. For account numbers that are listed as
// principals, to work with a resource listed in the attachment, you must sign in
// to an account specified as a principal. Then, you can work with resources that
// are listed, with any of your accelerators. If an accelerator ARN is listed in
// the cross-account attachment as a principal, anyone with permission to make
// updates to the accelerator can work with resources that are listed in the
// attachment.
//
// Specify each principal and resource separately. To specify two CIDR address
// pools, list them individually under Resources , and so on. For a command line
// operation, for example, you might use a statement like the following:
//
// "Resources": [{"Cidr": "169.254.60.0/24"},{"Cidr": "169.254.59.0/24"}]
//
// For more information, see [Working with cross-account attachments and resources in Global Accelerator] in the Global Accelerator Developer Guide.
//
// [Working with cross-account attachments and resources in Global Accelerator]: https://docs.aws.amazon.com/global-accelerator/latest/dg/cross-account-resources.html
func globalaccelerator_CreateCrossAccountAttachment(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.CreateCrossAccountAttachmentInput{
		// IdempotencyToken: *string, // Required
		// Name: *string, // Required
	}

	if len(_globalacceleratorIdempotencyToken) > 0 {
		input.IdempotencyToken = aws.String(_globalacceleratorIdempotencyToken)
	}
	if len(_globalacceleratorName) > 0 {
		input.Name = aws.String(_globalacceleratorName)
	}
	if len(_globalacceleratorPrincipals) > 0 {
		input.Principals = append([]string(nil), _globalacceleratorPrincipals...)
	}
	if len(_globalacceleratorResources) > 0 {
		if err := assignInputField(input, "Resources", _globalacceleratorResources); err != nil {
			log.Errorf("invalid --resources: %s", err.Error())
			return
		}
	}
	if len(_globalacceleratorTags) > 0 {
		if err := assignInputField(input, "Tags", _globalacceleratorTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCrossAccountAttachment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a custom routing accelerator. A custom routing accelerator directs
// traffic to one of possibly thousands of Amazon EC2 instance destinations running
// in a single or multiple virtual private clouds (VPC) subnet endpoints.
//
// Be aware that, by default, all destination EC2 instances in a VPC subnet
// endpoint cannot receive traffic. To enable all destinations to receive traffic,
// or to specify individual port mappings that can receive traffic, see the [AllowCustomRoutingTraffic]
// operation.
//
// Global Accelerator is a global service that supports endpoints in multiple
// Amazon Web Services Regions but you must specify the US West (Oregon) Region to
// create, update, or otherwise work with accelerators. That is, for example,
// specify --region us-west-2 on Amazon Web Services CLI commands.
//
// [AllowCustomRoutingTraffic]: https://docs.aws.amazon.com/global-accelerator/latest/api/API_AllowCustomRoutingTraffic.html
func globalaccelerator_CreateCustomRoutingAccelerator(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.CreateCustomRoutingAcceleratorInput{
		// IdempotencyToken: *string, // Required
		// Name: *string, // Required
	}

	if len(_globalacceleratorIdempotencyToken) > 0 {
		input.IdempotencyToken = aws.String(_globalacceleratorIdempotencyToken)
	}
	if len(_globalacceleratorName) > 0 {
		input.Name = aws.String(_globalacceleratorName)
	}
	if len(_globalacceleratorEnabled) > 0 {
		if err := assignInputField(input, "Enabled", _globalacceleratorEnabled); err != nil {
			log.Errorf("invalid --enabled: %s", err.Error())
			return
		}
	}
	if len(_globalacceleratorIpAddressType) > 0 {
		if err := assignInputField(input, "IpAddressType", _globalacceleratorIpAddressType); err != nil {
			log.Errorf("invalid --ip-address-type: %s", err.Error())
			return
		}
	}
	if len(_globalacceleratorIpAddresses) > 0 {
		input.IpAddresses = append([]string(nil), _globalacceleratorIpAddresses...)
	}
	if len(_globalacceleratorTags) > 0 {
		if err := assignInputField(input, "Tags", _globalacceleratorTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCustomRoutingAccelerator(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create an endpoint group for the specified listener for a custom routing
// accelerator. An endpoint group is a collection of endpoints in one Amazon Web
// Services Region.
func globalaccelerator_CreateCustomRoutingEndpointGroup(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.CreateCustomRoutingEndpointGroupInput{
		// DestinationConfigurations: []types.CustomRoutingDestinationConfiguration, // Required
		// EndpointGroupRegion: *string, // Required
		// IdempotencyToken: *string, // Required
		// ListenerArn: *string, // Required
	}

	if len(_globalacceleratorDestinationConfigurations) > 0 {
		if err := assignInputField(input, "DestinationConfigurations", _globalacceleratorDestinationConfigurations); err != nil {
			log.Errorf("invalid --destination-configurations: %s", err.Error())
			return
		}
	}
	if len(_globalacceleratorEndpointGroupRegion) > 0 {
		input.EndpointGroupRegion = aws.String(_globalacceleratorEndpointGroupRegion)
	}
	if len(_globalacceleratorIdempotencyToken) > 0 {
		input.IdempotencyToken = aws.String(_globalacceleratorIdempotencyToken)
	}
	if len(_globalacceleratorListenerArn) > 0 {
		input.ListenerArn = aws.String(_globalacceleratorListenerArn)
	}

	if resp, err := client.CreateCustomRoutingEndpointGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a listener to process inbound connections from clients to a custom
// routing accelerator. Connections arrive to assigned static IP addresses on the
// port range that you specify.
func globalaccelerator_CreateCustomRoutingListener(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.CreateCustomRoutingListenerInput{
		// AcceleratorArn: *string, // Required
		// IdempotencyToken: *string, // Required
		// PortRanges: []types.PortRange, // Required
	}

	if len(_globalacceleratorAcceleratorArn) > 0 {
		input.AcceleratorArn = aws.String(_globalacceleratorAcceleratorArn)
	}
	if len(_globalacceleratorIdempotencyToken) > 0 {
		input.IdempotencyToken = aws.String(_globalacceleratorIdempotencyToken)
	}
	if len(_globalacceleratorPortRanges) > 0 {
		if err := assignInputField(input, "PortRanges", _globalacceleratorPortRanges); err != nil {
			log.Errorf("invalid --port-ranges: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCustomRoutingListener(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create an endpoint group for the specified listener. An endpoint group is a
// collection of endpoints in one Amazon Web Services Region. A resource must be
// valid and active when you add it as an endpoint.
//
// For more information about endpoint types and requirements for endpoints that
// you can add to Global Accelerator, see [Endpoints for standard accelerators]in the Global Accelerator Developer
// Guide.
//
// [Endpoints for standard accelerators]: https://docs.aws.amazon.com/global-accelerator/latest/dg/about-endpoints.html
func globalaccelerator_CreateEndpointGroup(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.CreateEndpointGroupInput{
		// EndpointGroupRegion: *string, // Required
		// IdempotencyToken: *string, // Required
		// ListenerArn: *string, // Required
	}

	if len(_globalacceleratorEndpointGroupRegion) > 0 {
		input.EndpointGroupRegion = aws.String(_globalacceleratorEndpointGroupRegion)
	}
	if len(_globalacceleratorIdempotencyToken) > 0 {
		input.IdempotencyToken = aws.String(_globalacceleratorIdempotencyToken)
	}
	if len(_globalacceleratorListenerArn) > 0 {
		input.ListenerArn = aws.String(_globalacceleratorListenerArn)
	}
	if len(_globalacceleratorEndpointConfigurations) > 0 {
		if err := assignInputField(input, "EndpointConfigurations", _globalacceleratorEndpointConfigurations); err != nil {
			log.Errorf("invalid --endpoint-configurations: %s", err.Error())
			return
		}
	}
	if len(_globalacceleratorHealthCheckIntervalSeconds) > 0 {
		if err := assignInputField(input, "HealthCheckIntervalSeconds", _globalacceleratorHealthCheckIntervalSeconds); err != nil {
			log.Errorf("invalid --health-check-interval-seconds: %s", err.Error())
			return
		}
	}
	if len(_globalacceleratorHealthCheckPath) > 0 {
		input.HealthCheckPath = aws.String(_globalacceleratorHealthCheckPath)
	}
	if len(_globalacceleratorHealthCheckPort) > 0 {
		if err := assignInputField(input, "HealthCheckPort", _globalacceleratorHealthCheckPort); err != nil {
			log.Errorf("invalid --health-check-port: %s", err.Error())
			return
		}
	}
	if len(_globalacceleratorHealthCheckProtocol) > 0 {
		if err := assignInputField(input, "HealthCheckProtocol", _globalacceleratorHealthCheckProtocol); err != nil {
			log.Errorf("invalid --health-check-protocol: %s", err.Error())
			return
		}
	}
	if len(_globalacceleratorPortOverrides) > 0 {
		if err := assignInputField(input, "PortOverrides", _globalacceleratorPortOverrides); err != nil {
			log.Errorf("invalid --port-overrides: %s", err.Error())
			return
		}
	}
	if len(_globalacceleratorThresholdCount) > 0 {
		if err := assignInputField(input, "ThresholdCount", _globalacceleratorThresholdCount); err != nil {
			log.Errorf("invalid --threshold-count: %s", err.Error())
			return
		}
	}
	if len(_globalacceleratorTrafficDialPercentage) > 0 {
		if err := assignInputField(input, "TrafficDialPercentage", _globalacceleratorTrafficDialPercentage); err != nil {
			log.Errorf("invalid --traffic-dial-percentage: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateEndpointGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a listener to process inbound connections from clients to an
// accelerator. Connections arrive to assigned static IP addresses on a port, port
// range, or list of port ranges that you specify.
func globalaccelerator_CreateListener(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.CreateListenerInput{
		// AcceleratorArn: *string, // Required
		// IdempotencyToken: *string, // Required
		// PortRanges: []types.PortRange, // Required
		// Protocol: types.Protocol, // Required
	}

	if len(_globalacceleratorAcceleratorArn) > 0 {
		input.AcceleratorArn = aws.String(_globalacceleratorAcceleratorArn)
	}
	if len(_globalacceleratorIdempotencyToken) > 0 {
		input.IdempotencyToken = aws.String(_globalacceleratorIdempotencyToken)
	}
	if len(_globalacceleratorPortRanges) > 0 {
		if err := assignInputField(input, "PortRanges", _globalacceleratorPortRanges); err != nil {
			log.Errorf("invalid --port-ranges: %s", err.Error())
			return
		}
	}
	if len(_globalacceleratorProtocol) > 0 {
		if err := assignInputField(input, "Protocol", _globalacceleratorProtocol); err != nil {
			log.Errorf("invalid --protocol: %s", err.Error())
			return
		}
	}
	if len(_globalacceleratorClientAffinity) > 0 {
		if err := assignInputField(input, "ClientAffinity", _globalacceleratorClientAffinity); err != nil {
			log.Errorf("invalid --client-affinity: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateListener(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete an accelerator. Before you can delete an accelerator, you must disable
// it and remove all dependent resources (listeners and endpoint groups). To
// disable the accelerator, update the accelerator to set Enabled to false.
//
// When you create an accelerator, by default, Global Accelerator provides you
// with a set of two static IP addresses. Alternatively, you can bring your own IP
// address ranges to Global Accelerator and assign IP addresses from those ranges.
//
// The IP addresses are assigned to your accelerator for as long as it exists,
// even if you disable the accelerator and it no longer accepts or routes traffic.
// However, when you delete an accelerator, you lose the static IP addresses that
// are assigned to the accelerator, so you can no longer route traffic by using
// them. As a best practice, ensure that you have permissions in place to avoid
// inadvertently deleting accelerators. You can use IAM policies with Global
// Accelerator to limit the users who have permissions to delete an accelerator.
// For more information, see [Identity and access management]in the Global Accelerator Developer Guide.
//
// [Identity and access management]: https://docs.aws.amazon.com/global-accelerator/latest/dg/auth-and-access-control.html
func globalaccelerator_DeleteAccelerator(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.DeleteAcceleratorInput{
		// AcceleratorArn: *string, // Required
	}

	if len(_globalacceleratorAcceleratorArn) > 0 {
		input.AcceleratorArn = aws.String(_globalacceleratorAcceleratorArn)
	}

	if resp, err := client.DeleteAccelerator(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a cross-account attachment. When you delete an attachment, Global
// Accelerator revokes the permission to use the resources in the attachment from
// all principals in the list of principals. Global Accelerator revokes the
// permission for specific resources.
//
// For more information, see [Working with cross-account attachments and resources in Global Accelerator] in the Global Accelerator Developer Guide.
//
// [Working with cross-account attachments and resources in Global Accelerator]: https://docs.aws.amazon.com/global-accelerator/latest/dg/cross-account-resources.html
func globalaccelerator_DeleteCrossAccountAttachment(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.DeleteCrossAccountAttachmentInput{
		// AttachmentArn: *string, // Required
	}

	if len(_globalacceleratorAttachmentArn) > 0 {
		input.AttachmentArn = aws.String(_globalacceleratorAttachmentArn)
	}

	if resp, err := client.DeleteCrossAccountAttachment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a custom routing accelerator. Before you can delete an accelerator, you
// must disable it and remove all dependent resources (listeners and endpoint
// groups). To disable the accelerator, update the accelerator to set Enabled to
// false.
//
// When you create a custom routing accelerator, by default, Global Accelerator
// provides you with a set of two static IP addresses.
//
// The IP addresses are assigned to your accelerator for as long as it exists,
// even if you disable the accelerator and it no longer accepts or routes traffic.
// However, when you delete an accelerator, you lose the static IP addresses that
// are assigned to the accelerator, so you can no longer route traffic by using
// them. As a best practice, ensure that you have permissions in place to avoid
// inadvertently deleting accelerators. You can use IAM policies with Global
// Accelerator to limit the users who have permissions to delete an accelerator.
// For more information, see [Identity and access management]in the Global Accelerator Developer Guide.
//
// [Identity and access management]: https://docs.aws.amazon.com/global-accelerator/latest/dg/auth-and-access-control.html
func globalaccelerator_DeleteCustomRoutingAccelerator(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.DeleteCustomRoutingAcceleratorInput{
		// AcceleratorArn: *string, // Required
	}

	if len(_globalacceleratorAcceleratorArn) > 0 {
		input.AcceleratorArn = aws.String(_globalacceleratorAcceleratorArn)
	}

	if resp, err := client.DeleteCustomRoutingAccelerator(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete an endpoint group from a listener for a custom routing accelerator.
func globalaccelerator_DeleteCustomRoutingEndpointGroup(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.DeleteCustomRoutingEndpointGroupInput{
		// EndpointGroupArn: *string, // Required
	}

	if len(_globalacceleratorEndpointGroupArn) > 0 {
		input.EndpointGroupArn = aws.String(_globalacceleratorEndpointGroupArn)
	}

	if resp, err := client.DeleteCustomRoutingEndpointGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a listener for a custom routing accelerator.
func globalaccelerator_DeleteCustomRoutingListener(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.DeleteCustomRoutingListenerInput{
		// ListenerArn: *string, // Required
	}

	if len(_globalacceleratorListenerArn) > 0 {
		input.ListenerArn = aws.String(_globalacceleratorListenerArn)
	}

	if resp, err := client.DeleteCustomRoutingListener(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete an endpoint group from a listener.
func globalaccelerator_DeleteEndpointGroup(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.DeleteEndpointGroupInput{
		// EndpointGroupArn: *string, // Required
	}

	if len(_globalacceleratorEndpointGroupArn) > 0 {
		input.EndpointGroupArn = aws.String(_globalacceleratorEndpointGroupArn)
	}

	if resp, err := client.DeleteEndpointGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a listener from an accelerator.
func globalaccelerator_DeleteListener(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.DeleteListenerInput{
		// ListenerArn: *string, // Required
	}

	if len(_globalacceleratorListenerArn) > 0 {
		input.ListenerArn = aws.String(_globalacceleratorListenerArn)
	}

	if resp, err := client.DeleteListener(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Specify the Amazon EC2 instance (destination) IP addresses and ports for a VPC
// subnet endpoint that cannot receive traffic for a custom routing accelerator.
// You can deny traffic to all destinations in the VPC endpoint, or deny traffic to
// a specified list of destination IP addresses and ports. Note that you cannot
// specify IP addresses or ports outside of the range that you configured for the
// endpoint group.
//
// After you make changes, you can verify that the updates are complete by
// checking the status of your accelerator: the status changes from IN_PROGRESS to
// DEPLOYED.
func globalaccelerator_DenyCustomRoutingTraffic(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.DenyCustomRoutingTrafficInput{
		// EndpointGroupArn: *string, // Required
		// EndpointId: *string, // Required
	}

	if len(_globalacceleratorEndpointGroupArn) > 0 {
		input.EndpointGroupArn = aws.String(_globalacceleratorEndpointGroupArn)
	}
	if len(_globalacceleratorEndpointId) > 0 {
		input.EndpointId = aws.String(_globalacceleratorEndpointId)
	}
	if len(_globalacceleratorDenyAllTrafficToEndpoint) > 0 {
		if err := assignInputField(input, "DenyAllTrafficToEndpoint", _globalacceleratorDenyAllTrafficToEndpoint); err != nil {
			log.Errorf("invalid --deny-all-traffic-to-endpoint: %s", err.Error())
			return
		}
	}
	if len(_globalacceleratorDestinationAddresses) > 0 {
		input.DestinationAddresses = append([]string(nil), _globalacceleratorDestinationAddresses...)
	}
	if len(_globalacceleratorDestinationPorts) > 0 {
		if err := assignInputField(input, "DestinationPorts", _globalacceleratorDestinationPorts); err != nil {
			log.Errorf("invalid --destination-ports: %s", err.Error())
			return
		}
	}

	if resp, err := client.DenyCustomRoutingTraffic(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Releases the specified address range that you provisioned to use with your
// Amazon Web Services resources through bring your own IP addresses (BYOIP) and
// deletes the corresponding address pool.
//
// Before you can release an address range, you must stop advertising it by using [WithdrawByoipCidr]
// and you must not have any accelerators that are using static IP addresses
// allocated from its address range.
//
// For more information, see [Bring your own IP addresses (BYOIP)] in the Global Accelerator Developer Guide.
//
// [Bring your own IP addresses (BYOIP)]: https://docs.aws.amazon.com/global-accelerator/latest/dg/using-byoip.html
// [WithdrawByoipCidr]: https://docs.aws.amazon.com/global-accelerator/latest/api/WithdrawByoipCidr.html
func globalaccelerator_DeprovisionByoipCidr(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.DeprovisionByoipCidrInput{
		// Cidr: *string, // Required
	}

	if len(_globalacceleratorCidr) > 0 {
		input.Cidr = aws.String(_globalacceleratorCidr)
	}

	if resp, err := client.DeprovisionByoipCidr(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describe an accelerator.
func globalaccelerator_DescribeAccelerator(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.DescribeAcceleratorInput{
		// AcceleratorArn: *string, // Required
	}

	if len(_globalacceleratorAcceleratorArn) > 0 {
		input.AcceleratorArn = aws.String(_globalacceleratorAcceleratorArn)
	}

	if resp, err := client.DescribeAccelerator(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describe the attributes of an accelerator.
func globalaccelerator_DescribeAcceleratorAttributes(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.DescribeAcceleratorAttributesInput{
		// AcceleratorArn: *string, // Required
	}

	if len(_globalacceleratorAcceleratorArn) > 0 {
		input.AcceleratorArn = aws.String(_globalacceleratorAcceleratorArn)
	}

	if resp, err := client.DescribeAcceleratorAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets configuration information about a cross-account attachment.
func globalaccelerator_DescribeCrossAccountAttachment(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.DescribeCrossAccountAttachmentInput{
		// AttachmentArn: *string, // Required
	}

	if len(_globalacceleratorAttachmentArn) > 0 {
		input.AttachmentArn = aws.String(_globalacceleratorAttachmentArn)
	}

	if resp, err := client.DescribeCrossAccountAttachment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describe a custom routing accelerator.
func globalaccelerator_DescribeCustomRoutingAccelerator(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.DescribeCustomRoutingAcceleratorInput{
		// AcceleratorArn: *string, // Required
	}

	if len(_globalacceleratorAcceleratorArn) > 0 {
		input.AcceleratorArn = aws.String(_globalacceleratorAcceleratorArn)
	}

	if resp, err := client.DescribeCustomRoutingAccelerator(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describe the attributes of a custom routing accelerator.
func globalaccelerator_DescribeCustomRoutingAcceleratorAttributes(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.DescribeCustomRoutingAcceleratorAttributesInput{
		// AcceleratorArn: *string, // Required
	}

	if len(_globalacceleratorAcceleratorArn) > 0 {
		input.AcceleratorArn = aws.String(_globalacceleratorAcceleratorArn)
	}

	if resp, err := client.DescribeCustomRoutingAcceleratorAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describe an endpoint group for a custom routing accelerator.
func globalaccelerator_DescribeCustomRoutingEndpointGroup(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.DescribeCustomRoutingEndpointGroupInput{
		// EndpointGroupArn: *string, // Required
	}

	if len(_globalacceleratorEndpointGroupArn) > 0 {
		input.EndpointGroupArn = aws.String(_globalacceleratorEndpointGroupArn)
	}

	if resp, err := client.DescribeCustomRoutingEndpointGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The description of a listener for a custom routing accelerator.
func globalaccelerator_DescribeCustomRoutingListener(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.DescribeCustomRoutingListenerInput{
		// ListenerArn: *string, // Required
	}

	if len(_globalacceleratorListenerArn) > 0 {
		input.ListenerArn = aws.String(_globalacceleratorListenerArn)
	}

	if resp, err := client.DescribeCustomRoutingListener(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describe an endpoint group.
func globalaccelerator_DescribeEndpointGroup(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.DescribeEndpointGroupInput{
		// EndpointGroupArn: *string, // Required
	}

	if len(_globalacceleratorEndpointGroupArn) > 0 {
		input.EndpointGroupArn = aws.String(_globalacceleratorEndpointGroupArn)
	}

	if resp, err := client.DescribeEndpointGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describe a listener.
func globalaccelerator_DescribeListener(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.DescribeListenerInput{
		// ListenerArn: *string, // Required
	}

	if len(_globalacceleratorListenerArn) > 0 {
		input.ListenerArn = aws.String(_globalacceleratorListenerArn)
	}

	if resp, err := client.DescribeListener(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List the accelerators for an Amazon Web Services account.
func globalaccelerator_ListAccelerators(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.ListAcceleratorsInput{}

	if len(_globalacceleratorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _globalacceleratorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_globalacceleratorNextToken) > 0 {
		input.NextToken = aws.String(_globalacceleratorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAccelerators(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*globalaccelerator.ListAcceleratorsOutput
	p := globalaccelerator.NewListAcceleratorsPaginator(client, input)
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

// Lists the IP address ranges that were specified in calls to [ProvisionByoipCidr], including the
// current state and a history of state changes.
//
// [ProvisionByoipCidr]: https://docs.aws.amazon.com/global-accelerator/latest/api/ProvisionByoipCidr.html
func globalaccelerator_ListByoipCidrs(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.ListByoipCidrsInput{}

	if len(_globalacceleratorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _globalacceleratorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_globalacceleratorNextToken) > 0 {
		input.NextToken = aws.String(_globalacceleratorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListByoipCidrs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*globalaccelerator.ListByoipCidrsOutput
	p := globalaccelerator.NewListByoipCidrsPaginator(client, input)
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

// List the cross-account attachments that have been created in Global Accelerator.
func globalaccelerator_ListCrossAccountAttachments(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.ListCrossAccountAttachmentsInput{}

	if len(_globalacceleratorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _globalacceleratorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_globalacceleratorNextToken) > 0 {
		input.NextToken = aws.String(_globalacceleratorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCrossAccountAttachments(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*globalaccelerator.ListCrossAccountAttachmentsOutput
	p := globalaccelerator.NewListCrossAccountAttachmentsPaginator(client, input)
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

// List the accounts that have cross-account resources.
// For more information, see [Working with cross-account attachments and resources in Global Accelerator] in the Global Accelerator Developer Guide.
//
// [Working with cross-account attachments and resources in Global Accelerator]: https://docs.aws.amazon.com/global-accelerator/latest/dg/cross-account-resources.html
func globalaccelerator_ListCrossAccountResourceAccounts(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.ListCrossAccountResourceAccountsInput{}

	if resp, err := client.ListCrossAccountResourceAccounts(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List the cross-account resources available to work with.
func globalaccelerator_ListCrossAccountResources(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.ListCrossAccountResourcesInput{
		// ResourceOwnerAwsAccountId: *string, // Required
	}

	if len(_globalacceleratorResourceOwnerAwsAccountId) > 0 {
		input.ResourceOwnerAwsAccountId = aws.String(_globalacceleratorResourceOwnerAwsAccountId)
	}
	if len(_globalacceleratorAcceleratorArn) > 0 {
		input.AcceleratorArn = aws.String(_globalacceleratorAcceleratorArn)
	}
	if len(_globalacceleratorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _globalacceleratorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_globalacceleratorNextToken) > 0 {
		input.NextToken = aws.String(_globalacceleratorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCrossAccountResources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*globalaccelerator.ListCrossAccountResourcesOutput
	p := globalaccelerator.NewListCrossAccountResourcesPaginator(client, input)
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

// List the custom routing accelerators for an Amazon Web Services account.
func globalaccelerator_ListCustomRoutingAccelerators(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.ListCustomRoutingAcceleratorsInput{}

	if len(_globalacceleratorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _globalacceleratorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_globalacceleratorNextToken) > 0 {
		input.NextToken = aws.String(_globalacceleratorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCustomRoutingAccelerators(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*globalaccelerator.ListCustomRoutingAcceleratorsOutput
	p := globalaccelerator.NewListCustomRoutingAcceleratorsPaginator(client, input)
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

// List the endpoint groups that are associated with a listener for a custom
// routing accelerator.
func globalaccelerator_ListCustomRoutingEndpointGroups(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.ListCustomRoutingEndpointGroupsInput{
		// ListenerArn: *string, // Required
	}

	if len(_globalacceleratorListenerArn) > 0 {
		input.ListenerArn = aws.String(_globalacceleratorListenerArn)
	}
	if len(_globalacceleratorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _globalacceleratorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_globalacceleratorNextToken) > 0 {
		input.NextToken = aws.String(_globalacceleratorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCustomRoutingEndpointGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*globalaccelerator.ListCustomRoutingEndpointGroupsOutput
	p := globalaccelerator.NewListCustomRoutingEndpointGroupsPaginator(client, input)
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

// List the listeners for a custom routing accelerator.
func globalaccelerator_ListCustomRoutingListeners(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.ListCustomRoutingListenersInput{
		// AcceleratorArn: *string, // Required
	}

	if len(_globalacceleratorAcceleratorArn) > 0 {
		input.AcceleratorArn = aws.String(_globalacceleratorAcceleratorArn)
	}
	if len(_globalacceleratorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _globalacceleratorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_globalacceleratorNextToken) > 0 {
		input.NextToken = aws.String(_globalacceleratorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCustomRoutingListeners(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*globalaccelerator.ListCustomRoutingListenersOutput
	p := globalaccelerator.NewListCustomRoutingListenersPaginator(client, input)
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

// Provides a complete mapping from the public accelerator IP address and port to
// destination EC2 instance IP addresses and ports in the virtual public cloud
// (VPC) subnet endpoint for a custom routing accelerator. For each subnet endpoint
// that you add, Global Accelerator creates a new static port mapping for the
// accelerator. The port mappings don't change after Global Accelerator generates
// them, so you can retrieve and cache the full mapping on your servers.
//
// If you remove a subnet from your accelerator, Global Accelerator removes
// (reclaims) the port mappings. If you add a subnet to your accelerator, Global
// Accelerator creates new port mappings (the existing ones don't change). If you
// add or remove EC2 instances in your subnet, the port mappings don't change,
// because the mappings are created when you add the subnet to Global Accelerator.
//
// The mappings also include a flag for each destination denoting which
// destination IP addresses and ports are allowed or denied traffic.
func globalaccelerator_ListCustomRoutingPortMappings(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.ListCustomRoutingPortMappingsInput{
		// AcceleratorArn: *string, // Required
	}

	if len(_globalacceleratorAcceleratorArn) > 0 {
		input.AcceleratorArn = aws.String(_globalacceleratorAcceleratorArn)
	}
	if len(_globalacceleratorEndpointGroupArn) > 0 {
		input.EndpointGroupArn = aws.String(_globalacceleratorEndpointGroupArn)
	}
	if len(_globalacceleratorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _globalacceleratorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_globalacceleratorNextToken) > 0 {
		input.NextToken = aws.String(_globalacceleratorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCustomRoutingPortMappings(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*globalaccelerator.ListCustomRoutingPortMappingsOutput
	p := globalaccelerator.NewListCustomRoutingPortMappingsPaginator(client, input)
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

// List the port mappings for a specific EC2 instance (destination) in a VPC
// subnet endpoint. The response is the mappings for one destination IP address.
// This is useful when your subnet endpoint has mappings that span multiple custom
// routing accelerators in your account, or for scenarios where you only want to
// list the port mappings for a specific destination instance.
func globalaccelerator_ListCustomRoutingPortMappingsByDestination(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.ListCustomRoutingPortMappingsByDestinationInput{
		// DestinationAddress: *string, // Required
		// EndpointId: *string, // Required
	}

	if len(_globalacceleratorDestinationAddress) > 0 {
		input.DestinationAddress = aws.String(_globalacceleratorDestinationAddress)
	}
	if len(_globalacceleratorEndpointId) > 0 {
		input.EndpointId = aws.String(_globalacceleratorEndpointId)
	}
	if len(_globalacceleratorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _globalacceleratorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_globalacceleratorNextToken) > 0 {
		input.NextToken = aws.String(_globalacceleratorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCustomRoutingPortMappingsByDestination(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*globalaccelerator.ListCustomRoutingPortMappingsByDestinationOutput
	p := globalaccelerator.NewListCustomRoutingPortMappingsByDestinationPaginator(client, input)
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

// List the endpoint groups that are associated with a listener.
func globalaccelerator_ListEndpointGroups(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.ListEndpointGroupsInput{
		// ListenerArn: *string, // Required
	}

	if len(_globalacceleratorListenerArn) > 0 {
		input.ListenerArn = aws.String(_globalacceleratorListenerArn)
	}
	if len(_globalacceleratorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _globalacceleratorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_globalacceleratorNextToken) > 0 {
		input.NextToken = aws.String(_globalacceleratorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEndpointGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*globalaccelerator.ListEndpointGroupsOutput
	p := globalaccelerator.NewListEndpointGroupsPaginator(client, input)
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

// List the listeners for an accelerator.
func globalaccelerator_ListListeners(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.ListListenersInput{
		// AcceleratorArn: *string, // Required
	}

	if len(_globalacceleratorAcceleratorArn) > 0 {
		input.AcceleratorArn = aws.String(_globalacceleratorAcceleratorArn)
	}
	if len(_globalacceleratorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _globalacceleratorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_globalacceleratorNextToken) > 0 {
		input.NextToken = aws.String(_globalacceleratorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListListeners(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*globalaccelerator.ListListenersOutput
	p := globalaccelerator.NewListListenersPaginator(client, input)
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

// List all tags for an accelerator.
// For more information, see [Tagging in Global Accelerator] in the Global Accelerator Developer Guide.
//
// [Tagging in Global Accelerator]: https://docs.aws.amazon.com/global-accelerator/latest/dg/tagging-in-global-accelerator.html
func globalaccelerator_ListTagsForResource(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_globalacceleratorResourceArn) > 0 {
		input.ResourceArn = aws.String(_globalacceleratorResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provisions an IP address range to use with your Amazon Web Services resources
// through bring your own IP addresses (BYOIP) and creates a corresponding address
// pool. After the address range is provisioned, it is ready to be advertised using
// [AdvertiseByoipCidr].
//
// For more information, see [Bring your own IP addresses (BYOIP)] in the Global Accelerator Developer Guide.
//
// [AdvertiseByoipCidr]: https://docs.aws.amazon.com/global-accelerator/latest/api/AdvertiseByoipCidr.html
// [Bring your own IP addresses (BYOIP)]: https://docs.aws.amazon.com/global-accelerator/latest/dg/using-byoip.html
func globalaccelerator_ProvisionByoipCidr(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.ProvisionByoipCidrInput{
		// Cidr: *string, // Required
		// CidrAuthorizationContext: *types.CidrAuthorizationContext, // Required
	}

	if len(_globalacceleratorCidr) > 0 {
		input.Cidr = aws.String(_globalacceleratorCidr)
	}
	if len(_globalacceleratorCidrAuthorizationContext) > 0 {
		if err := assignInputField(input, "CidrAuthorizationContext", _globalacceleratorCidrAuthorizationContext); err != nil {
			log.Errorf("invalid --cidr-authorization-context: %s", err.Error())
			return
		}
	}

	if resp, err := client.ProvisionByoipCidr(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Remove endpoints from a custom routing accelerator.
func globalaccelerator_RemoveCustomRoutingEndpoints(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.RemoveCustomRoutingEndpointsInput{
		// EndpointGroupArn: *string, // Required
		// EndpointIds: []string, // Required
	}

	if len(_globalacceleratorEndpointGroupArn) > 0 {
		input.EndpointGroupArn = aws.String(_globalacceleratorEndpointGroupArn)
	}
	if len(_globalacceleratorEndpointIds) > 0 {
		input.EndpointIds = append([]string(nil), _globalacceleratorEndpointIds...)
	}

	if resp, err := client.RemoveCustomRoutingEndpoints(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Remove endpoints from an endpoint group.
// The RemoveEndpoints API operation is the recommended option for removing
// endpoints. The alternative is to remove endpoints by updating an endpoint group
// by using the [UpdateEndpointGroup]API operation. There are two advantages to using AddEndpoints to
// remove endpoints instead:
//
// - It's more convenient, because you only need to specify the endpoints that
// you want to remove. With the UpdateEndpointGroup API operation, you must
// specify all of the endpoints in the endpoint group except the ones that you want
// to remove from the group.
//
// - It's faster, because Global Accelerator doesn't need to resolve any
// endpoints. With the UpdateEndpointGroup API operation, Global Accelerator must
// resolve all of the endpoints that remain in the group.
//
// [UpdateEndpointGroup]: https://docs.aws.amazon.com/global-accelerator/latest/api/API_UpdateEndpointGroup.html
func globalaccelerator_RemoveEndpoints(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.RemoveEndpointsInput{
		// EndpointGroupArn: *string, // Required
		// EndpointIdentifiers: []types.EndpointIdentifier, // Required
	}

	if len(_globalacceleratorEndpointGroupArn) > 0 {
		input.EndpointGroupArn = aws.String(_globalacceleratorEndpointGroupArn)
	}
	if len(_globalacceleratorEndpointIdentifiers) > 0 {
		if err := assignInputField(input, "EndpointIdentifiers", _globalacceleratorEndpointIdentifiers); err != nil {
			log.Errorf("invalid --endpoint-identifiers: %s", err.Error())
			return
		}
	}

	if resp, err := client.RemoveEndpoints(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Add tags to an accelerator resource.
// For more information, see [Tagging in Global Accelerator] in the Global Accelerator Developer Guide.
//
// [Tagging in Global Accelerator]: https://docs.aws.amazon.com/global-accelerator/latest/dg/tagging-in-global-accelerator.html
func globalaccelerator_TagResource(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_globalacceleratorResourceArn) > 0 {
		input.ResourceArn = aws.String(_globalacceleratorResourceArn)
	}
	if len(_globalacceleratorTags) > 0 {
		if err := assignInputField(input, "Tags", _globalacceleratorTags); err != nil {
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

// Remove tags from a Global Accelerator resource. When you specify a tag key, the
// action removes both that key and its associated value. The operation succeeds
// even if you attempt to remove tags from an accelerator that was already removed.
//
// For more information, see [Tagging in Global Accelerator] in the Global Accelerator Developer Guide.
//
// [Tagging in Global Accelerator]: https://docs.aws.amazon.com/global-accelerator/latest/dg/tagging-in-global-accelerator.html
func globalaccelerator_UntagResource(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_globalacceleratorResourceArn) > 0 {
		input.ResourceArn = aws.String(_globalacceleratorResourceArn)
	}
	if len(_globalacceleratorTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _globalacceleratorTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update an accelerator to make changes, such as the following:
// - Change the name of the accelerator.
//
// - Disable the accelerator so that it no longer accepts or routes traffic, or
// so that you can delete it.
//
// - Enable the accelerator, if it is disabled.
//
// - Change the IP address type to dual-stack if it is IPv4, or change the IP
// address type to IPv4 if it's dual-stack.
//
// Be aware that static IP addresses remain assigned to your accelerator for as
// long as it exists, even if you disable the accelerator and it no longer accepts
// or routes traffic. However, when you delete the accelerator, you lose the static
// IP addresses that are assigned to it, so you can no longer route traffic by
// using them.
//
// Global Accelerator is a global service that supports endpoints in multiple
// Amazon Web Services Regions but you must specify the US West (Oregon) Region to
// create, update, or otherwise work with accelerators. That is, for example,
// specify --region us-west-2 on Amazon Web Services CLI commands.
func globalaccelerator_UpdateAccelerator(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.UpdateAcceleratorInput{
		// AcceleratorArn: *string, // Required
	}

	if len(_globalacceleratorAcceleratorArn) > 0 {
		input.AcceleratorArn = aws.String(_globalacceleratorAcceleratorArn)
	}
	if len(_globalacceleratorEnabled) > 0 {
		if err := assignInputField(input, "Enabled", _globalacceleratorEnabled); err != nil {
			log.Errorf("invalid --enabled: %s", err.Error())
			return
		}
	}
	if len(_globalacceleratorIpAddressType) > 0 {
		if err := assignInputField(input, "IpAddressType", _globalacceleratorIpAddressType); err != nil {
			log.Errorf("invalid --ip-address-type: %s", err.Error())
			return
		}
	}
	if len(_globalacceleratorIpAddresses) > 0 {
		input.IpAddresses = append([]string(nil), _globalacceleratorIpAddresses...)
	}
	if len(_globalacceleratorName) > 0 {
		input.Name = aws.String(_globalacceleratorName)
	}

	if resp, err := client.UpdateAccelerator(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update the attributes for an accelerator.
func globalaccelerator_UpdateAcceleratorAttributes(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.UpdateAcceleratorAttributesInput{
		// AcceleratorArn: *string, // Required
	}

	if len(_globalacceleratorAcceleratorArn) > 0 {
		input.AcceleratorArn = aws.String(_globalacceleratorAcceleratorArn)
	}
	if len(_globalacceleratorFlowLogsEnabled) > 0 {
		if err := assignInputField(input, "FlowLogsEnabled", _globalacceleratorFlowLogsEnabled); err != nil {
			log.Errorf("invalid --flow-logs-enabled: %s", err.Error())
			return
		}
	}
	if len(_globalacceleratorFlowLogsS3Bucket) > 0 {
		input.FlowLogsS3Bucket = aws.String(_globalacceleratorFlowLogsS3Bucket)
	}
	if len(_globalacceleratorFlowLogsS3Prefix) > 0 {
		input.FlowLogsS3Prefix = aws.String(_globalacceleratorFlowLogsS3Prefix)
	}

	if resp, err := client.UpdateAcceleratorAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update a cross-account attachment to add or remove principals or resources.
// When you update an attachment to remove a principal (account ID or accelerator)
// or a resource, Global Accelerator revokes the permission for specific resources.
//
// For more information, see [Working with cross-account attachments and resources in Global Accelerator] in the Global Accelerator Developer Guide.
//
// [Working with cross-account attachments and resources in Global Accelerator]: https://docs.aws.amazon.com/global-accelerator/latest/dg/cross-account-resources.html
func globalaccelerator_UpdateCrossAccountAttachment(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.UpdateCrossAccountAttachmentInput{
		// AttachmentArn: *string, // Required
	}

	if len(_globalacceleratorAttachmentArn) > 0 {
		input.AttachmentArn = aws.String(_globalacceleratorAttachmentArn)
	}
	if len(_globalacceleratorAddPrincipals) > 0 {
		input.AddPrincipals = append([]string(nil), _globalacceleratorAddPrincipals...)
	}
	if len(_globalacceleratorAddResources) > 0 {
		if err := assignInputField(input, "AddResources", _globalacceleratorAddResources); err != nil {
			log.Errorf("invalid --add-resources: %s", err.Error())
			return
		}
	}
	if len(_globalacceleratorName) > 0 {
		input.Name = aws.String(_globalacceleratorName)
	}
	if len(_globalacceleratorRemovePrincipals) > 0 {
		input.RemovePrincipals = append([]string(nil), _globalacceleratorRemovePrincipals...)
	}
	if len(_globalacceleratorRemoveResources) > 0 {
		if err := assignInputField(input, "RemoveResources", _globalacceleratorRemoveResources); err != nil {
			log.Errorf("invalid --remove-resources: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateCrossAccountAttachment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update a custom routing accelerator.
func globalaccelerator_UpdateCustomRoutingAccelerator(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.UpdateCustomRoutingAcceleratorInput{
		// AcceleratorArn: *string, // Required
	}

	if len(_globalacceleratorAcceleratorArn) > 0 {
		input.AcceleratorArn = aws.String(_globalacceleratorAcceleratorArn)
	}
	if len(_globalacceleratorEnabled) > 0 {
		if err := assignInputField(input, "Enabled", _globalacceleratorEnabled); err != nil {
			log.Errorf("invalid --enabled: %s", err.Error())
			return
		}
	}
	if len(_globalacceleratorIpAddressType) > 0 {
		if err := assignInputField(input, "IpAddressType", _globalacceleratorIpAddressType); err != nil {
			log.Errorf("invalid --ip-address-type: %s", err.Error())
			return
		}
	}
	if len(_globalacceleratorIpAddresses) > 0 {
		input.IpAddresses = append([]string(nil), _globalacceleratorIpAddresses...)
	}
	if len(_globalacceleratorName) > 0 {
		input.Name = aws.String(_globalacceleratorName)
	}

	if resp, err := client.UpdateCustomRoutingAccelerator(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update the attributes for a custom routing accelerator.
func globalaccelerator_UpdateCustomRoutingAcceleratorAttributes(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.UpdateCustomRoutingAcceleratorAttributesInput{
		// AcceleratorArn: *string, // Required
	}

	if len(_globalacceleratorAcceleratorArn) > 0 {
		input.AcceleratorArn = aws.String(_globalacceleratorAcceleratorArn)
	}
	if len(_globalacceleratorFlowLogsEnabled) > 0 {
		if err := assignInputField(input, "FlowLogsEnabled", _globalacceleratorFlowLogsEnabled); err != nil {
			log.Errorf("invalid --flow-logs-enabled: %s", err.Error())
			return
		}
	}
	if len(_globalacceleratorFlowLogsS3Bucket) > 0 {
		input.FlowLogsS3Bucket = aws.String(_globalacceleratorFlowLogsS3Bucket)
	}
	if len(_globalacceleratorFlowLogsS3Prefix) > 0 {
		input.FlowLogsS3Prefix = aws.String(_globalacceleratorFlowLogsS3Prefix)
	}

	if resp, err := client.UpdateCustomRoutingAcceleratorAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update a listener for a custom routing accelerator.
func globalaccelerator_UpdateCustomRoutingListener(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.UpdateCustomRoutingListenerInput{
		// ListenerArn: *string, // Required
		// PortRanges: []types.PortRange, // Required
	}

	if len(_globalacceleratorListenerArn) > 0 {
		input.ListenerArn = aws.String(_globalacceleratorListenerArn)
	}
	if len(_globalacceleratorPortRanges) > 0 {
		if err := assignInputField(input, "PortRanges", _globalacceleratorPortRanges); err != nil {
			log.Errorf("invalid --port-ranges: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateCustomRoutingListener(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update an endpoint group. A resource must be valid and active when you add it
// as an endpoint.
func globalaccelerator_UpdateEndpointGroup(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.UpdateEndpointGroupInput{
		// EndpointGroupArn: *string, // Required
	}

	if len(_globalacceleratorEndpointGroupArn) > 0 {
		input.EndpointGroupArn = aws.String(_globalacceleratorEndpointGroupArn)
	}
	if len(_globalacceleratorEndpointConfigurations) > 0 {
		if err := assignInputField(input, "EndpointConfigurations", _globalacceleratorEndpointConfigurations); err != nil {
			log.Errorf("invalid --endpoint-configurations: %s", err.Error())
			return
		}
	}
	if len(_globalacceleratorHealthCheckIntervalSeconds) > 0 {
		if err := assignInputField(input, "HealthCheckIntervalSeconds", _globalacceleratorHealthCheckIntervalSeconds); err != nil {
			log.Errorf("invalid --health-check-interval-seconds: %s", err.Error())
			return
		}
	}
	if len(_globalacceleratorHealthCheckPath) > 0 {
		input.HealthCheckPath = aws.String(_globalacceleratorHealthCheckPath)
	}
	if len(_globalacceleratorHealthCheckPort) > 0 {
		if err := assignInputField(input, "HealthCheckPort", _globalacceleratorHealthCheckPort); err != nil {
			log.Errorf("invalid --health-check-port: %s", err.Error())
			return
		}
	}
	if len(_globalacceleratorHealthCheckProtocol) > 0 {
		if err := assignInputField(input, "HealthCheckProtocol", _globalacceleratorHealthCheckProtocol); err != nil {
			log.Errorf("invalid --health-check-protocol: %s", err.Error())
			return
		}
	}
	if len(_globalacceleratorPortOverrides) > 0 {
		if err := assignInputField(input, "PortOverrides", _globalacceleratorPortOverrides); err != nil {
			log.Errorf("invalid --port-overrides: %s", err.Error())
			return
		}
	}
	if len(_globalacceleratorThresholdCount) > 0 {
		if err := assignInputField(input, "ThresholdCount", _globalacceleratorThresholdCount); err != nil {
			log.Errorf("invalid --threshold-count: %s", err.Error())
			return
		}
	}
	if len(_globalacceleratorTrafficDialPercentage) > 0 {
		if err := assignInputField(input, "TrafficDialPercentage", _globalacceleratorTrafficDialPercentage); err != nil {
			log.Errorf("invalid --traffic-dial-percentage: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateEndpointGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update a listener.
func globalaccelerator_UpdateListener(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.UpdateListenerInput{
		// ListenerArn: *string, // Required
	}

	if len(_globalacceleratorListenerArn) > 0 {
		input.ListenerArn = aws.String(_globalacceleratorListenerArn)
	}
	if len(_globalacceleratorClientAffinity) > 0 {
		if err := assignInputField(input, "ClientAffinity", _globalacceleratorClientAffinity); err != nil {
			log.Errorf("invalid --client-affinity: %s", err.Error())
			return
		}
	}
	if len(_globalacceleratorPortRanges) > 0 {
		if err := assignInputField(input, "PortRanges", _globalacceleratorPortRanges); err != nil {
			log.Errorf("invalid --port-ranges: %s", err.Error())
			return
		}
	}
	if len(_globalacceleratorProtocol) > 0 {
		if err := assignInputField(input, "Protocol", _globalacceleratorProtocol); err != nil {
			log.Errorf("invalid --protocol: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateListener(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops advertising an address range that is provisioned as an address pool. You
// can perform this operation at most once every 10 seconds, even if you specify
// different address ranges each time.
//
// It can take a few minutes before traffic to the specified addresses stops
// routing to Amazon Web Services because of propagation delays.
//
// For more information, see [Bring your own IP addresses (BYOIP)] in the Global Accelerator Developer Guide.
//
// [Bring your own IP addresses (BYOIP)]: https://docs.aws.amazon.com/global-accelerator/latest/dg/using-byoip.html
func globalaccelerator_WithdrawByoipCidr(cfg aws.Config, client *globalaccelerator.Client) {
	input := &globalaccelerator.WithdrawByoipCidrInput{
		// Cidr: *string, // Required
	}

	if len(_globalacceleratorCidr) > 0 {
		input.Cidr = aws.String(_globalacceleratorCidr)
	}

	if resp, err := client.WithdrawByoipCidr(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_globalacceleratorCmd)
	_globalacceleratorCmd.Flags().SortFlags = false

	_globalacceleratorCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_globalacceleratorCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_globalacceleratorCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_globalacceleratorCmd.Flags().StringVarP(&_globalacceleratorAcceleratorArn, "accelerator-arn", "", "", "Accelerator ARN")
	_globalacceleratorCmd.Flags().StringSliceVarP(&_globalacceleratorAddPrincipals, "add-principals", "", nil, "Add Principals")
	_globalacceleratorCmd.Flags().StringVarP(&_globalacceleratorAddResources, "add-resources", "", "", "Add Resources")
	_globalacceleratorCmd.Flags().StringVarP(&_globalacceleratorAllowAllTrafficToEndpoint, "allow-all-traffic-to-endpoint", "", "", "Allow All Traffic To Endpoint")
	_globalacceleratorCmd.Flags().StringVarP(&_globalacceleratorAttachmentArn, "attachment-arn", "", "", "Attachment ARN")
	_globalacceleratorCmd.Flags().StringVarP(&_globalacceleratorCidr, "cidr", "", "", "CIDR")
	_globalacceleratorCmd.Flags().StringVarP(&_globalacceleratorCidrAuthorizationContext, "cidr-authorization-context", "", "", "CIDR Authorization Context")
	_globalacceleratorCmd.Flags().StringVarP(&_globalacceleratorClientAffinity, "client-affinity", "", "", "Client Affinity")
	_globalacceleratorCmd.Flags().StringVarP(&_globalacceleratorDenyAllTrafficToEndpoint, "deny-all-traffic-to-endpoint", "", "", "Deny All Traffic To Endpoint")
	_globalacceleratorCmd.Flags().StringVarP(&_globalacceleratorDestinationAddress, "destination-address", "", "", "Destination Address")
	_globalacceleratorCmd.Flags().StringSliceVarP(&_globalacceleratorDestinationAddresses, "destination-addresses", "", nil, "Destination Addresses")
	_globalacceleratorCmd.Flags().StringVarP(&_globalacceleratorDestinationConfigurations, "destination-configurations", "", "", "Destination Configurations")
	_globalacceleratorCmd.Flags().StringVarP(&_globalacceleratorDestinationPorts, "destination-ports", "", "", "Destination Ports")
	_globalacceleratorCmd.Flags().StringVarP(&_globalacceleratorEnabled, "enabled", "", "", "Enabled")
	_globalacceleratorCmd.Flags().StringVarP(&_globalacceleratorEndpointConfigurations, "endpoint-configurations", "", "", "Endpoint Configurations")
	_globalacceleratorCmd.Flags().StringVarP(&_globalacceleratorEndpointGroupArn, "endpoint-group-arn", "", "", "Endpoint Group ARN")
	_globalacceleratorCmd.Flags().StringVarP(&_globalacceleratorEndpointGroupRegion, "endpoint-group-region", "", "", "Endpoint Group Region")
	_globalacceleratorCmd.Flags().StringVarP(&_globalacceleratorEndpointId, "endpoint-id", "", "", "Endpoint ID")
	_globalacceleratorCmd.Flags().StringVarP(&_globalacceleratorEndpointIdentifiers, "endpoint-identifiers", "", "", "Endpoint Identifiers")
	_globalacceleratorCmd.Flags().StringSliceVarP(&_globalacceleratorEndpointIds, "endpoint-ids", "", nil, "Endpoint Ids")
	_globalacceleratorCmd.Flags().StringVarP(&_globalacceleratorFlowLogsEnabled, "flow-logs-enabled", "", "", "Flow Logs Enabled")
	_globalacceleratorCmd.Flags().StringVarP(&_globalacceleratorFlowLogsS3Bucket, "flow-logs-s3-bucket", "", "", "Flow Logs S3 Bucket")
	_globalacceleratorCmd.Flags().StringVarP(&_globalacceleratorFlowLogsS3Prefix, "flow-logs-s3-prefix", "", "", "Flow Logs S3 Prefix")
	_globalacceleratorCmd.Flags().StringVarP(&_globalacceleratorHealthCheckIntervalSeconds, "health-check-interval-seconds", "", "", "Health Check Interval Seconds")
	_globalacceleratorCmd.Flags().StringVarP(&_globalacceleratorHealthCheckPath, "health-check-path", "", "", "Health Check Path")
	_globalacceleratorCmd.Flags().StringVarP(&_globalacceleratorHealthCheckPort, "health-check-port", "", "", "Health Check Port")
	_globalacceleratorCmd.Flags().StringVarP(&_globalacceleratorHealthCheckProtocol, "health-check-protocol", "", "", "Health Check Protocol")
	_globalacceleratorCmd.Flags().StringVarP(&_globalacceleratorIdempotencyToken, "idempotency-token", "", "", "Idempotency Token")
	_globalacceleratorCmd.Flags().StringVarP(&_globalacceleratorIpAddressType, "ip-address-type", "", "", "IP Address Type")
	_globalacceleratorCmd.Flags().StringSliceVarP(&_globalacceleratorIpAddresses, "ip-addresses", "", nil, "IP Addresses")
	_globalacceleratorCmd.Flags().StringVarP(&_globalacceleratorListenerArn, "listener-arn", "", "", "Listener ARN")
	_globalacceleratorCmd.Flags().StringVarP(&_globalacceleratorMaxResults, "max-results", "", "", "Max Results")
	_globalacceleratorCmd.Flags().StringVarP(&_globalacceleratorName, "name", "", "", "Name")
	_globalacceleratorCmd.Flags().StringVarP(&_globalacceleratorNextToken, "next-token", "", "", "Next Token")
	_globalacceleratorCmd.Flags().StringVarP(&_globalacceleratorPortOverrides, "port-overrides", "", "", "Port Overrides")
	_globalacceleratorCmd.Flags().StringVarP(&_globalacceleratorPortRanges, "port-ranges", "", "", "Port Ranges")
	_globalacceleratorCmd.Flags().StringSliceVarP(&_globalacceleratorPrincipals, "principals", "", nil, "Principals")
	_globalacceleratorCmd.Flags().StringVarP(&_globalacceleratorProtocol, "protocol", "", "", "Protocol")
	_globalacceleratorCmd.Flags().StringSliceVarP(&_globalacceleratorRemovePrincipals, "remove-principals", "", nil, "Remove Principals")
	_globalacceleratorCmd.Flags().StringVarP(&_globalacceleratorRemoveResources, "remove-resources", "", "", "Remove Resources")
	_globalacceleratorCmd.Flags().StringVarP(&_globalacceleratorResourceArn, "resource-arn", "", "", "Resource ARN")
	_globalacceleratorCmd.Flags().StringVarP(&_globalacceleratorResourceOwnerAwsAccountId, "resource-owner-aws-account-id", "", "", "Resource Owner AWS Account ID")
	_globalacceleratorCmd.Flags().StringVarP(&_globalacceleratorResources, "resources", "", "", "Resources")
	_globalacceleratorCmd.Flags().StringSliceVarP(&_globalacceleratorTagKeys, "tag-keys", "", nil, "Tag Keys")
	_globalacceleratorCmd.Flags().StringVarP(&_globalacceleratorTags, "tags", "", "", "Tags")
	_globalacceleratorCmd.Flags().StringVarP(&_globalacceleratorThresholdCount, "threshold-count", "", "", "Threshold Count")
	_globalacceleratorCmd.Flags().StringVarP(&_globalacceleratorTrafficDialPercentage, "traffic-dial-percentage", "", "", "Traffic Dial Percentage")

	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorAddCustomRoutingEndpoints, "add-custom-routing-endpoints", "", false, "Add Custom Routing Endpoints")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorAddEndpoints, "add-endpoints", "", false, "Add Endpoints")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorAdvertiseByoipCidr, "advertise-byoip-cidr", "", false, "Advertise Byoip CIDR")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorAllowCustomRoutingTraffic, "allow-custom-routing-traffic", "", false, "Allow Custom Routing Traffic")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorCreateAccelerator, "create-accelerator", "", false, "Create Accelerator")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorCreateCrossAccountAttachment, "create-cross-account-attachment", "", false, "Create Cross Account Attachment")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorCreateCustomRoutingAccelerator, "create-custom-routing-accelerator", "", false, "Create Custom Routing Accelerator")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorCreateCustomRoutingEndpointGroup, "create-custom-routing-endpoint-group", "", false, "Create Custom Routing Endpoint Group")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorCreateCustomRoutingListener, "create-custom-routing-listener", "", false, "Create Custom Routing Listener")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorCreateEndpointGroup, "create-endpoint-group", "", false, "Create Endpoint Group")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorCreateListener, "create-listener", "", false, "Create Listener")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorDeleteAccelerator, "delete-accelerator", "", false, "Delete Accelerator")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorDeleteCrossAccountAttachment, "delete-cross-account-attachment", "", false, "Delete Cross Account Attachment")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorDeleteCustomRoutingAccelerator, "delete-custom-routing-accelerator", "", false, "Delete Custom Routing Accelerator")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorDeleteCustomRoutingEndpointGroup, "delete-custom-routing-endpoint-group", "", false, "Delete Custom Routing Endpoint Group")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorDeleteCustomRoutingListener, "delete-custom-routing-listener", "", false, "Delete Custom Routing Listener")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorDeleteEndpointGroup, "delete-endpoint-group", "", false, "Delete Endpoint Group")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorDeleteListener, "delete-listener", "", false, "Delete Listener")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorDenyCustomRoutingTraffic, "deny-custom-routing-traffic", "", false, "Deny Custom Routing Traffic")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorDeprovisionByoipCidr, "deprovision-byoip-cidr", "", false, "Deprovision Byoip CIDR")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorDescribeAccelerator, "describe-accelerator", "", false, "Describe Accelerator")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorDescribeAcceleratorAttributes, "describe-accelerator-attributes", "", false, "Describe Accelerator Attributes")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorDescribeCrossAccountAttachment, "describe-cross-account-attachment", "", false, "Describe Cross Account Attachment")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorDescribeCustomRoutingAccelerator, "describe-custom-routing-accelerator", "", false, "Describe Custom Routing Accelerator")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorDescribeCustomRoutingAcceleratorAttributes, "describe-custom-routing-accelerator-attributes", "", false, "Describe Custom Routing Accelerator Attributes")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorDescribeCustomRoutingEndpointGroup, "describe-custom-routing-endpoint-group", "", false, "Describe Custom Routing Endpoint Group")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorDescribeCustomRoutingListener, "describe-custom-routing-listener", "", false, "Describe Custom Routing Listener")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorDescribeEndpointGroup, "describe-endpoint-group", "", false, "Describe Endpoint Group")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorDescribeListener, "describe-listener", "", false, "Describe Listener")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorListAccelerators, "list-accelerators", "", false, "List Accelerators")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorListByoipCidrs, "list-byoip-cidrs", "", false, "List Byoip Cidrs")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorListCrossAccountAttachments, "list-cross-account-attachments", "", false, "List Cross Account Attachments")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorListCrossAccountResourceAccounts, "list-cross-account-resource-accounts", "", false, "List Cross Account Resource Accounts")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorListCrossAccountResources, "list-cross-account-resources", "", false, "List Cross Account Resources")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorListCustomRoutingAccelerators, "list-custom-routing-accelerators", "", false, "List Custom Routing Accelerators")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorListCustomRoutingEndpointGroups, "list-custom-routing-endpoint-groups", "", false, "List Custom Routing Endpoint Groups")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorListCustomRoutingListeners, "list-custom-routing-listeners", "", false, "List Custom Routing Listeners")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorListCustomRoutingPortMappings, "list-custom-routing-port-mappings", "", false, "List Custom Routing Port Mappings")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorListCustomRoutingPortMappingsByDestination, "list-custom-routing-port-mappings-by-destination", "", false, "List Custom Routing Port Mappings By Destination")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorListEndpointGroups, "list-endpoint-groups", "", false, "List Endpoint Groups")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorListListeners, "list-listeners", "", false, "List Listeners")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorProvisionByoipCidr, "provision-byoip-cidr", "", false, "Provision Byoip CIDR")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorRemoveCustomRoutingEndpoints, "remove-custom-routing-endpoints", "", false, "Remove Custom Routing Endpoints")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorRemoveEndpoints, "remove-endpoints", "", false, "Remove Endpoints")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorTagResource, "tag-resource", "", false, "Tag Resource")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorUntagResource, "untag-resource", "", false, "Untag Resource")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorUpdateAccelerator, "update-accelerator", "", false, "Update Accelerator")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorUpdateAcceleratorAttributes, "update-accelerator-attributes", "", false, "Update Accelerator Attributes")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorUpdateCrossAccountAttachment, "update-cross-account-attachment", "", false, "Update Cross Account Attachment")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorUpdateCustomRoutingAccelerator, "update-custom-routing-accelerator", "", false, "Update Custom Routing Accelerator")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorUpdateCustomRoutingAcceleratorAttributes, "update-custom-routing-accelerator-attributes", "", false, "Update Custom Routing Accelerator Attributes")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorUpdateCustomRoutingListener, "update-custom-routing-listener", "", false, "Update Custom Routing Listener")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorUpdateEndpointGroup, "update-endpoint-group", "", false, "Update Endpoint Group")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorUpdateListener, "update-listener", "", false, "Update Listener")
	_globalacceleratorCmd.Flags().BoolVarP(&_globalacceleratorWithdrawByoipCidr, "withdraw-byoip-cidr", "", false, "Withdraw Byoip CIDR")

}
