package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/vpclattice"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// vpclatticeCmd represents the vpclattice command
var _vpclatticeCmd = &cobra.Command{
	Use:   "vpclattice",
	Short: "AWS vpclattice CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := vpclattice.NewFromConfig(cfg)
		if _vpclatticeBatchUpdateRule {
			vpclattice_BatchUpdateRule(cfg, client)
			return
		}
		if _vpclatticeCreateAccessLogSubscription {
			vpclattice_CreateAccessLogSubscription(cfg, client)
			return
		}
		if _vpclatticeCreateListener {
			vpclattice_CreateListener(cfg, client)
			return
		}
		if _vpclatticeCreateResourceConfiguration {
			vpclattice_CreateResourceConfiguration(cfg, client)
			return
		}
		if _vpclatticeCreateResourceGateway {
			vpclattice_CreateResourceGateway(cfg, client)
			return
		}
		if _vpclatticeCreateRule {
			vpclattice_CreateRule(cfg, client)
			return
		}
		if _vpclatticeCreateService {
			vpclattice_CreateService(cfg, client)
			return
		}
		if _vpclatticeCreateServiceNetwork {
			vpclattice_CreateServiceNetwork(cfg, client)
			return
		}
		if _vpclatticeCreateServiceNetworkResourceAssociation {
			vpclattice_CreateServiceNetworkResourceAssociation(cfg, client)
			return
		}
		if _vpclatticeCreateServiceNetworkServiceAssociation {
			vpclattice_CreateServiceNetworkServiceAssociation(cfg, client)
			return
		}
		if _vpclatticeCreateServiceNetworkVpcAssociation {
			vpclattice_CreateServiceNetworkVpcAssociation(cfg, client)
			return
		}
		if _vpclatticeCreateTargetGroup {
			vpclattice_CreateTargetGroup(cfg, client)
			return
		}
		if _vpclatticeDeleteAccessLogSubscription {
			vpclattice_DeleteAccessLogSubscription(cfg, client)
			return
		}
		if _vpclatticeDeleteAuthPolicy {
			vpclattice_DeleteAuthPolicy(cfg, client)
			return
		}
		if _vpclatticeDeleteDomainVerification {
			vpclattice_DeleteDomainVerification(cfg, client)
			return
		}
		if _vpclatticeDeleteListener {
			vpclattice_DeleteListener(cfg, client)
			return
		}
		if _vpclatticeDeleteResourceConfiguration {
			vpclattice_DeleteResourceConfiguration(cfg, client)
			return
		}
		if _vpclatticeDeleteResourceEndpointAssociation {
			vpclattice_DeleteResourceEndpointAssociation(cfg, client)
			return
		}
		if _vpclatticeDeleteResourceGateway {
			vpclattice_DeleteResourceGateway(cfg, client)
			return
		}
		if _vpclatticeDeleteResourcePolicy {
			vpclattice_DeleteResourcePolicy(cfg, client)
			return
		}
		if _vpclatticeDeleteRule {
			vpclattice_DeleteRule(cfg, client)
			return
		}
		if _vpclatticeDeleteService {
			vpclattice_DeleteService(cfg, client)
			return
		}
		if _vpclatticeDeleteServiceNetwork {
			vpclattice_DeleteServiceNetwork(cfg, client)
			return
		}
		if _vpclatticeDeleteServiceNetworkResourceAssociation {
			vpclattice_DeleteServiceNetworkResourceAssociation(cfg, client)
			return
		}
		if _vpclatticeDeleteServiceNetworkServiceAssociation {
			vpclattice_DeleteServiceNetworkServiceAssociation(cfg, client)
			return
		}
		if _vpclatticeDeleteServiceNetworkVpcAssociation {
			vpclattice_DeleteServiceNetworkVpcAssociation(cfg, client)
			return
		}
		if _vpclatticeDeleteTargetGroup {
			vpclattice_DeleteTargetGroup(cfg, client)
			return
		}
		if _vpclatticeDeregisterTargets {
			vpclattice_DeregisterTargets(cfg, client)
			return
		}
		if _vpclatticeGetAccessLogSubscription {
			vpclattice_GetAccessLogSubscription(cfg, client)
			return
		}
		if _vpclatticeGetAuthPolicy {
			vpclattice_GetAuthPolicy(cfg, client)
			return
		}
		if _vpclatticeGetDomainVerification {
			vpclattice_GetDomainVerification(cfg, client)
			return
		}
		if _vpclatticeGetListener {
			vpclattice_GetListener(cfg, client)
			return
		}
		if _vpclatticeGetResourceConfiguration {
			vpclattice_GetResourceConfiguration(cfg, client)
			return
		}
		if _vpclatticeGetResourceGateway {
			vpclattice_GetResourceGateway(cfg, client)
			return
		}
		if _vpclatticeGetResourcePolicy {
			vpclattice_GetResourcePolicy(cfg, client)
			return
		}
		if _vpclatticeGetRule {
			vpclattice_GetRule(cfg, client)
			return
		}
		if _vpclatticeGetService {
			vpclattice_GetService(cfg, client)
			return
		}
		if _vpclatticeGetServiceNetwork {
			vpclattice_GetServiceNetwork(cfg, client)
			return
		}
		if _vpclatticeGetServiceNetworkResourceAssociation {
			vpclattice_GetServiceNetworkResourceAssociation(cfg, client)
			return
		}
		if _vpclatticeGetServiceNetworkServiceAssociation {
			vpclattice_GetServiceNetworkServiceAssociation(cfg, client)
			return
		}
		if _vpclatticeGetServiceNetworkVpcAssociation {
			vpclattice_GetServiceNetworkVpcAssociation(cfg, client)
			return
		}
		if _vpclatticeGetTargetGroup {
			vpclattice_GetTargetGroup(cfg, client)
			return
		}
		if _vpclatticeListAccessLogSubscriptions {
			vpclattice_ListAccessLogSubscriptions(cfg, client)
			return
		}
		if _vpclatticeListDomainVerifications {
			vpclattice_ListDomainVerifications(cfg, client)
			return
		}
		if _vpclatticeListListeners {
			vpclattice_ListListeners(cfg, client)
			return
		}
		if _vpclatticeListResourceConfigurations {
			vpclattice_ListResourceConfigurations(cfg, client)
			return
		}
		if _vpclatticeListResourceEndpointAssociations {
			vpclattice_ListResourceEndpointAssociations(cfg, client)
			return
		}
		if _vpclatticeListResourceGateways {
			vpclattice_ListResourceGateways(cfg, client)
			return
		}
		if _vpclatticeListRules {
			vpclattice_ListRules(cfg, client)
			return
		}
		if _vpclatticeListServiceNetworkResourceAssociations {
			vpclattice_ListServiceNetworkResourceAssociations(cfg, client)
			return
		}
		if _vpclatticeListServiceNetworkServiceAssociations {
			vpclattice_ListServiceNetworkServiceAssociations(cfg, client)
			return
		}
		if _vpclatticeListServiceNetworkVpcAssociations {
			vpclattice_ListServiceNetworkVpcAssociations(cfg, client)
			return
		}
		if _vpclatticeListServiceNetworkVpcEndpointAssociations {
			vpclattice_ListServiceNetworkVpcEndpointAssociations(cfg, client)
			return
		}
		if _vpclatticeListServiceNetworks {
			vpclattice_ListServiceNetworks(cfg, client)
			return
		}
		if _vpclatticeListServices {
			vpclattice_ListServices(cfg, client)
			return
		}
		if _vpclatticeListTagsForResource {
			vpclattice_ListTagsForResource(cfg, client)
			return
		}
		if _vpclatticeListTargetGroups {
			vpclattice_ListTargetGroups(cfg, client)
			return
		}
		if _vpclatticeListTargets {
			vpclattice_ListTargets(cfg, client)
			return
		}
		if _vpclatticePutAuthPolicy {
			vpclattice_PutAuthPolicy(cfg, client)
			return
		}
		if _vpclatticePutResourcePolicy {
			vpclattice_PutResourcePolicy(cfg, client)
			return
		}
		if _vpclatticeRegisterTargets {
			vpclattice_RegisterTargets(cfg, client)
			return
		}
		if _vpclatticeStartDomainVerification {
			vpclattice_StartDomainVerification(cfg, client)
			return
		}
		if _vpclatticeTagResource {
			vpclattice_TagResource(cfg, client)
			return
		}
		if _vpclatticeUntagResource {
			vpclattice_UntagResource(cfg, client)
			return
		}
		if _vpclatticeUpdateAccessLogSubscription {
			vpclattice_UpdateAccessLogSubscription(cfg, client)
			return
		}
		if _vpclatticeUpdateListener {
			vpclattice_UpdateListener(cfg, client)
			return
		}
		if _vpclatticeUpdateResourceConfiguration {
			vpclattice_UpdateResourceConfiguration(cfg, client)
			return
		}
		if _vpclatticeUpdateResourceGateway {
			vpclattice_UpdateResourceGateway(cfg, client)
			return
		}
		if _vpclatticeUpdateRule {
			vpclattice_UpdateRule(cfg, client)
			return
		}
		if _vpclatticeUpdateService {
			vpclattice_UpdateService(cfg, client)
			return
		}
		if _vpclatticeUpdateServiceNetwork {
			vpclattice_UpdateServiceNetwork(cfg, client)
			return
		}
		if _vpclatticeUpdateServiceNetworkVpcAssociation {
			vpclattice_UpdateServiceNetworkVpcAssociation(cfg, client)
			return
		}
		if _vpclatticeUpdateTargetGroup {
			vpclattice_UpdateTargetGroup(cfg, client)
			return
		}

	},
}

var (
	_vpclatticeBatchUpdateRule                           bool
	_vpclatticeCreateAccessLogSubscription               bool
	_vpclatticeCreateListener                            bool
	_vpclatticeCreateResourceConfiguration               bool
	_vpclatticeCreateResourceGateway                     bool
	_vpclatticeCreateRule                                bool
	_vpclatticeCreateService                             bool
	_vpclatticeCreateServiceNetwork                      bool
	_vpclatticeCreateServiceNetworkResourceAssociation   bool
	_vpclatticeCreateServiceNetworkServiceAssociation    bool
	_vpclatticeCreateServiceNetworkVpcAssociation        bool
	_vpclatticeCreateTargetGroup                         bool
	_vpclatticeDeleteAccessLogSubscription               bool
	_vpclatticeDeleteAuthPolicy                          bool
	_vpclatticeDeleteDomainVerification                  bool
	_vpclatticeDeleteListener                            bool
	_vpclatticeDeleteResourceConfiguration               bool
	_vpclatticeDeleteResourceEndpointAssociation         bool
	_vpclatticeDeleteResourceGateway                     bool
	_vpclatticeDeleteResourcePolicy                      bool
	_vpclatticeDeleteRule                                bool
	_vpclatticeDeleteService                             bool
	_vpclatticeDeleteServiceNetwork                      bool
	_vpclatticeDeleteServiceNetworkResourceAssociation   bool
	_vpclatticeDeleteServiceNetworkServiceAssociation    bool
	_vpclatticeDeleteServiceNetworkVpcAssociation        bool
	_vpclatticeDeleteTargetGroup                         bool
	_vpclatticeDeregisterTargets                         bool
	_vpclatticeGetAccessLogSubscription                  bool
	_vpclatticeGetAuthPolicy                             bool
	_vpclatticeGetDomainVerification                     bool
	_vpclatticeGetListener                               bool
	_vpclatticeGetResourceConfiguration                  bool
	_vpclatticeGetResourceGateway                        bool
	_vpclatticeGetResourcePolicy                         bool
	_vpclatticeGetRule                                   bool
	_vpclatticeGetService                                bool
	_vpclatticeGetServiceNetwork                         bool
	_vpclatticeGetServiceNetworkResourceAssociation      bool
	_vpclatticeGetServiceNetworkServiceAssociation       bool
	_vpclatticeGetServiceNetworkVpcAssociation           bool
	_vpclatticeGetTargetGroup                            bool
	_vpclatticeListAccessLogSubscriptions                bool
	_vpclatticeListDomainVerifications                   bool
	_vpclatticeListListeners                             bool
	_vpclatticeListResourceConfigurations                bool
	_vpclatticeListResourceEndpointAssociations          bool
	_vpclatticeListResourceGateways                      bool
	_vpclatticeListRules                                 bool
	_vpclatticeListServiceNetworkResourceAssociations    bool
	_vpclatticeListServiceNetworkServiceAssociations     bool
	_vpclatticeListServiceNetworkVpcAssociations         bool
	_vpclatticeListServiceNetworkVpcEndpointAssociations bool
	_vpclatticeListServiceNetworks                       bool
	_vpclatticeListServices                              bool
	_vpclatticeListTagsForResource                       bool
	_vpclatticeListTargetGroups                          bool
	_vpclatticeListTargets                               bool
	_vpclatticePutAuthPolicy                             bool
	_vpclatticePutResourcePolicy                         bool
	_vpclatticeRegisterTargets                           bool
	_vpclatticeStartDomainVerification                   bool
	_vpclatticeTagResource                               bool
	_vpclatticeUntagResource                             bool
	_vpclatticeUpdateAccessLogSubscription               bool
	_vpclatticeUpdateListener                            bool
	_vpclatticeUpdateResourceConfiguration               bool
	_vpclatticeUpdateResourceGateway                     bool
	_vpclatticeUpdateRule                                bool
	_vpclatticeUpdateService                             bool
	_vpclatticeUpdateServiceNetwork                      bool
	_vpclatticeUpdateServiceNetworkVpcAssociation        bool
	_vpclatticeUpdateTargetGroup                         bool

	_vpclatticeAccessLogSubscriptionIdentifier             string
	_vpclatticeAction                                      string
	_vpclatticeAllowAssociationToShareableServiceNetwork   string
	_vpclatticeAuthType                                    string
	_vpclatticeCertificateArn                              string
	_vpclatticeClientToken                                 string
	_vpclatticeConfig                                      string
	_vpclatticeCustomDomainName                            string
	_vpclatticeDefaultAction                               string
	_vpclatticeDestinationArn                              string
	_vpclatticeDnsOptions                                  string
	_vpclatticeDomainName                                  string
	_vpclatticeDomainVerificationIdentifier                string
	_vpclatticeGroupDomain                                 string
	_vpclatticeHealthCheck                                 string
	_vpclatticeIncludeChildren                             string
	_vpclatticeIpAddressType                               string
	_vpclatticeIpv4AddressesPerEni                         string
	_vpclatticeListenerIdentifier                          string
	_vpclatticeMatch                                       string
	_vpclatticeMaxResults                                  string
	_vpclatticeName                                        string
	_vpclatticeNextToken                                   string
	_vpclatticePolicy                                      string
	_vpclatticePort                                        string
	_vpclatticePortRanges                                  []string
	_vpclatticePriority                                    string
	_vpclatticePrivateDnsEnabled                           string
	_vpclatticeProtocol                                    string
	_vpclatticeResourceArn                                 string
	_vpclatticeResourceConfigurationDefinition             string
	_vpclatticeResourceConfigurationGroupIdentifier        string
	_vpclatticeResourceConfigurationIdentifier             string
	_vpclatticeResourceEndpointAssociationIdentifier       string
	_vpclatticeResourceGatewayIdentifier                   string
	_vpclatticeResourceIdentifier                          string
	_vpclatticeRuleIdentifier                              string
	_vpclatticeRules                                       string
	_vpclatticeSecurityGroupIds                            []string
	_vpclatticeServiceIdentifier                           string
	_vpclatticeServiceNetworkIdentifier                    string
	_vpclatticeServiceNetworkLogType                       string
	_vpclatticeServiceNetworkResourceAssociationIdentifier string
	_vpclatticeServiceNetworkServiceAssociationIdentifier  string
	_vpclatticeServiceNetworkVpcAssociationIdentifier      string
	_vpclatticeSharingConfig                               string
	_vpclatticeSubnetIds                                   []string
	_vpclatticeTagKeys                                     []string
	_vpclatticeTags                                        string
	_vpclatticeTargetGroupIdentifier                       string
	_vpclatticeTargetGroupType                             string
	_vpclatticeTargets                                     string
	_vpclatticeType                                        string
	_vpclatticeVpcEndpointId                               string
	_vpclatticeVpcEndpointOwner                            string
	_vpclatticeVpcIdentifier                               string
)

// Updates the listener rules in a batch. You can use this operation to change the
// priority of listener rules. This can be useful when bulk updating or swapping
// rule priority.
//
// Required permissions: vpc-lattice:UpdateRule
//
// For more information, see [How Amazon VPC Lattice works with IAM] in the Amazon VPC Lattice User Guide.
//
// [How Amazon VPC Lattice works with IAM]: https://docs.aws.amazon.com/vpc-lattice/latest/ug/security_iam_service-with-iam.html
func vpclattice_BatchUpdateRule(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.BatchUpdateRuleInput{
		// ListenerIdentifier: *string, // Required
		// Rules: []types.RuleUpdate, // Required
		// ServiceIdentifier: *string, // Required
	}

	if len(_vpclatticeListenerIdentifier) > 0 {
		input.ListenerIdentifier = aws.String(_vpclatticeListenerIdentifier)
	}
	if len(_vpclatticeRules) > 0 {
		if err := assignInputField(input, "Rules", _vpclatticeRules); err != nil {
			log.Errorf("invalid --rules: %s", err.Error())
			return
		}
	}
	if len(_vpclatticeServiceIdentifier) > 0 {
		input.ServiceIdentifier = aws.String(_vpclatticeServiceIdentifier)
	}

	if resp, err := client.BatchUpdateRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables access logs to be sent to Amazon CloudWatch, Amazon S3, and Amazon
// Kinesis Data Firehose. The service network owner can use the access logs to
// audit the services in the network. The service network owner can only see access
// logs from clients and services that are associated with their service network.
// Access log entries represent traffic originated from VPCs associated with that
// network. For more information, see [Access logs]in the Amazon VPC Lattice User Guide.
//
// [Access logs]: https://docs.aws.amazon.com/vpc-lattice/latest/ug/monitoring-access-logs.html
func vpclattice_CreateAccessLogSubscription(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.CreateAccessLogSubscriptionInput{
		// DestinationArn: *string, // Required
		// ResourceIdentifier: *string, // Required
	}

	if len(_vpclatticeDestinationArn) > 0 {
		input.DestinationArn = aws.String(_vpclatticeDestinationArn)
	}
	if len(_vpclatticeResourceIdentifier) > 0 {
		input.ResourceIdentifier = aws.String(_vpclatticeResourceIdentifier)
	}
	if len(_vpclatticeClientToken) > 0 {
		input.ClientToken = aws.String(_vpclatticeClientToken)
	}
	if len(_vpclatticeServiceNetworkLogType) > 0 {
		if err := assignInputField(input, "ServiceNetworkLogType", _vpclatticeServiceNetworkLogType); err != nil {
			log.Errorf("invalid --service-network-log-type: %s", err.Error())
			return
		}
	}
	if len(_vpclatticeTags) > 0 {
		if err := assignInputField(input, "Tags", _vpclatticeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAccessLogSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a listener for a service. Before you start using your Amazon VPC
// Lattice service, you must add one or more listeners. A listener is a process
// that checks for connection requests to your services. For more information, see [Listeners]
// in the Amazon VPC Lattice User Guide.
//
// [Listeners]: https://docs.aws.amazon.com/vpc-lattice/latest/ug/listeners.html
func vpclattice_CreateListener(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.CreateListenerInput{
		// DefaultAction: types.RuleAction, // Required
		// Name: *string, // Required
		// Protocol: types.ListenerProtocol, // Required
		// ServiceIdentifier: *string, // Required
	}

	if len(_vpclatticeDefaultAction) > 0 {
		if err := assignInputField(input, "DefaultAction", _vpclatticeDefaultAction); err != nil {
			log.Errorf("invalid --default-action: %s", err.Error())
			return
		}
	}
	if len(_vpclatticeName) > 0 {
		input.Name = aws.String(_vpclatticeName)
	}
	if len(_vpclatticeProtocol) > 0 {
		if err := assignInputField(input, "Protocol", _vpclatticeProtocol); err != nil {
			log.Errorf("invalid --protocol: %s", err.Error())
			return
		}
	}
	if len(_vpclatticeServiceIdentifier) > 0 {
		input.ServiceIdentifier = aws.String(_vpclatticeServiceIdentifier)
	}
	if len(_vpclatticeClientToken) > 0 {
		input.ClientToken = aws.String(_vpclatticeClientToken)
	}
	if len(_vpclatticePort) > 0 {
		if err := assignInputField(input, "Port", _vpclatticePort); err != nil {
			log.Errorf("invalid --port: %s", err.Error())
			return
		}
	}
	if len(_vpclatticeTags) > 0 {
		if err := assignInputField(input, "Tags", _vpclatticeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
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

// Creates a resource configuration. A resource configuration defines a specific
// resource. You can associate a resource configuration with a service network or a
// VPC endpoint.
func vpclattice_CreateResourceConfiguration(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.CreateResourceConfigurationInput{
		// Name: *string, // Required
		// Type: types.ResourceConfigurationType, // Required
	}

	if len(_vpclatticeName) > 0 {
		input.Name = aws.String(_vpclatticeName)
	}
	if len(_vpclatticeType) > 0 {
		if err := assignInputField(input, "Type", _vpclatticeType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_vpclatticeAllowAssociationToShareableServiceNetwork) > 0 {
		if err := assignInputField(input, "AllowAssociationToShareableServiceNetwork", _vpclatticeAllowAssociationToShareableServiceNetwork); err != nil {
			log.Errorf("invalid --allow-association-to-shareable-service-network: %s", err.Error())
			return
		}
	}
	if len(_vpclatticeClientToken) > 0 {
		input.ClientToken = aws.String(_vpclatticeClientToken)
	}
	if len(_vpclatticeCustomDomainName) > 0 {
		input.CustomDomainName = aws.String(_vpclatticeCustomDomainName)
	}
	if len(_vpclatticeDomainVerificationIdentifier) > 0 {
		input.DomainVerificationIdentifier = aws.String(_vpclatticeDomainVerificationIdentifier)
	}
	if len(_vpclatticeGroupDomain) > 0 {
		input.GroupDomain = aws.String(_vpclatticeGroupDomain)
	}
	if len(_vpclatticePortRanges) > 0 {
		input.PortRanges = append([]string(nil), _vpclatticePortRanges...)
	}
	if len(_vpclatticeProtocol) > 0 {
		if err := assignInputField(input, "Protocol", _vpclatticeProtocol); err != nil {
			log.Errorf("invalid --protocol: %s", err.Error())
			return
		}
	}
	if len(_vpclatticeResourceConfigurationDefinition) > 0 {
		if err := assignInputField(input, "ResourceConfigurationDefinition", _vpclatticeResourceConfigurationDefinition); err != nil {
			log.Errorf("invalid --resource-configuration-definition: %s", err.Error())
			return
		}
	}
	if len(_vpclatticeResourceConfigurationGroupIdentifier) > 0 {
		input.ResourceConfigurationGroupIdentifier = aws.String(_vpclatticeResourceConfigurationGroupIdentifier)
	}
	if len(_vpclatticeResourceGatewayIdentifier) > 0 {
		input.ResourceGatewayIdentifier = aws.String(_vpclatticeResourceGatewayIdentifier)
	}
	if len(_vpclatticeTags) > 0 {
		if err := assignInputField(input, "Tags", _vpclatticeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateResourceConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// A resource gateway is a point of ingress into the VPC where a resource resides.
// It spans multiple Availability Zones. For your resource to be accessible from
// all Availability Zones, you should create your resource gateways to span as many
// Availability Zones as possible. A VPC can have multiple resource gateways.
func vpclattice_CreateResourceGateway(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.CreateResourceGatewayInput{
		// Name: *string, // Required
	}

	if len(_vpclatticeName) > 0 {
		input.Name = aws.String(_vpclatticeName)
	}
	if len(_vpclatticeClientToken) > 0 {
		input.ClientToken = aws.String(_vpclatticeClientToken)
	}
	if len(_vpclatticeIpAddressType) > 0 {
		if err := assignInputField(input, "IpAddressType", _vpclatticeIpAddressType); err != nil {
			log.Errorf("invalid --ip-address-type: %s", err.Error())
			return
		}
	}
	if len(_vpclatticeIpv4AddressesPerEni) > 0 {
		if err := assignInputField(input, "Ipv4AddressesPerEni", _vpclatticeIpv4AddressesPerEni); err != nil {
			log.Errorf("invalid --ipv4-addresses-per-eni: %s", err.Error())
			return
		}
	}
	if len(_vpclatticeSecurityGroupIds) > 0 {
		input.SecurityGroupIds = append([]string(nil), _vpclatticeSecurityGroupIds...)
	}
	if len(_vpclatticeSubnetIds) > 0 {
		input.SubnetIds = append([]string(nil), _vpclatticeSubnetIds...)
	}
	if len(_vpclatticeTags) > 0 {
		if err := assignInputField(input, "Tags", _vpclatticeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_vpclatticeVpcIdentifier) > 0 {
		input.VpcIdentifier = aws.String(_vpclatticeVpcIdentifier)
	}

	if resp, err := client.CreateResourceGateway(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a listener rule. Each listener has a default rule for checking
// connection requests, but you can define additional rules. Each rule consists of
// a priority, one or more actions, and one or more conditions. For more
// information, see [Listener rules]in the Amazon VPC Lattice User Guide.
//
// [Listener rules]: https://docs.aws.amazon.com/vpc-lattice/latest/ug/listeners.html#listener-rules
func vpclattice_CreateRule(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.CreateRuleInput{
		// Action: types.RuleAction, // Required
		// ListenerIdentifier: *string, // Required
		// Match: types.RuleMatch, // Required
		// Name: *string, // Required
		// Priority: *int32, // Required
		// ServiceIdentifier: *string, // Required
	}

	if len(_vpclatticeAction) > 0 {
		if err := assignInputField(input, "Action", _vpclatticeAction); err != nil {
			log.Errorf("invalid --action: %s", err.Error())
			return
		}
	}
	if len(_vpclatticeListenerIdentifier) > 0 {
		input.ListenerIdentifier = aws.String(_vpclatticeListenerIdentifier)
	}
	if len(_vpclatticeMatch) > 0 {
		if err := assignInputField(input, "Match", _vpclatticeMatch); err != nil {
			log.Errorf("invalid --match: %s", err.Error())
			return
		}
	}
	if len(_vpclatticeName) > 0 {
		input.Name = aws.String(_vpclatticeName)
	}
	if len(_vpclatticePriority) > 0 {
		if err := assignInputField(input, "Priority", _vpclatticePriority); err != nil {
			log.Errorf("invalid --priority: %s", err.Error())
			return
		}
	}
	if len(_vpclatticeServiceIdentifier) > 0 {
		input.ServiceIdentifier = aws.String(_vpclatticeServiceIdentifier)
	}
	if len(_vpclatticeClientToken) > 0 {
		input.ClientToken = aws.String(_vpclatticeClientToken)
	}
	if len(_vpclatticeTags) > 0 {
		if err := assignInputField(input, "Tags", _vpclatticeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a service. A service is any software application that can run on
// instances containers, or serverless functions within an account or virtual
// private cloud (VPC).
//
// For more information, see [Services] in the Amazon VPC Lattice User Guide.
//
// [Services]: https://docs.aws.amazon.com/vpc-lattice/latest/ug/services.html
func vpclattice_CreateService(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.CreateServiceInput{
		// Name: *string, // Required
	}

	if len(_vpclatticeName) > 0 {
		input.Name = aws.String(_vpclatticeName)
	}
	if len(_vpclatticeAuthType) > 0 {
		if err := assignInputField(input, "AuthType", _vpclatticeAuthType); err != nil {
			log.Errorf("invalid --auth-type: %s", err.Error())
			return
		}
	}
	if len(_vpclatticeCertificateArn) > 0 {
		input.CertificateArn = aws.String(_vpclatticeCertificateArn)
	}
	if len(_vpclatticeClientToken) > 0 {
		input.ClientToken = aws.String(_vpclatticeClientToken)
	}
	if len(_vpclatticeCustomDomainName) > 0 {
		input.CustomDomainName = aws.String(_vpclatticeCustomDomainName)
	}
	if len(_vpclatticeTags) > 0 {
		if err := assignInputField(input, "Tags", _vpclatticeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateService(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a service network. A service network is a logical boundary for a
// collection of services. You can associate services and VPCs with a service
// network.
//
// For more information, see [Service networks] in the Amazon VPC Lattice User Guide.
//
// [Service networks]: https://docs.aws.amazon.com/vpc-lattice/latest/ug/service-networks.html
func vpclattice_CreateServiceNetwork(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.CreateServiceNetworkInput{
		// Name: *string, // Required
	}

	if len(_vpclatticeName) > 0 {
		input.Name = aws.String(_vpclatticeName)
	}
	if len(_vpclatticeAuthType) > 0 {
		if err := assignInputField(input, "AuthType", _vpclatticeAuthType); err != nil {
			log.Errorf("invalid --auth-type: %s", err.Error())
			return
		}
	}
	if len(_vpclatticeClientToken) > 0 {
		input.ClientToken = aws.String(_vpclatticeClientToken)
	}
	if len(_vpclatticeSharingConfig) > 0 {
		if err := assignInputField(input, "SharingConfig", _vpclatticeSharingConfig); err != nil {
			log.Errorf("invalid --sharing-config: %s", err.Error())
			return
		}
	}
	if len(_vpclatticeTags) > 0 {
		if err := assignInputField(input, "Tags", _vpclatticeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateServiceNetwork(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates the specified service network with the specified resource
// configuration. This allows the resource configuration to receive connections
// through the service network, including through a service network VPC endpoint.
func vpclattice_CreateServiceNetworkResourceAssociation(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.CreateServiceNetworkResourceAssociationInput{
		// ResourceConfigurationIdentifier: *string, // Required
		// ServiceNetworkIdentifier: *string, // Required
	}

	if len(_vpclatticeResourceConfigurationIdentifier) > 0 {
		input.ResourceConfigurationIdentifier = aws.String(_vpclatticeResourceConfigurationIdentifier)
	}
	if len(_vpclatticeServiceNetworkIdentifier) > 0 {
		input.ServiceNetworkIdentifier = aws.String(_vpclatticeServiceNetworkIdentifier)
	}
	if len(_vpclatticeClientToken) > 0 {
		input.ClientToken = aws.String(_vpclatticeClientToken)
	}
	if len(_vpclatticePrivateDnsEnabled) > 0 {
		if err := assignInputField(input, "PrivateDnsEnabled", _vpclatticePrivateDnsEnabled); err != nil {
			log.Errorf("invalid --private-dns-enabled: %s", err.Error())
			return
		}
	}
	if len(_vpclatticeTags) > 0 {
		if err := assignInputField(input, "Tags", _vpclatticeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateServiceNetworkResourceAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates the specified service with the specified service network. For more
// information, see [Manage service associations]in the Amazon VPC Lattice User Guide.
//
// You can't use this operation if the service and service network are already
// associated or if there is a disassociation or deletion in progress. If the
// association fails, you can retry the operation by deleting the association and
// recreating it.
//
// You cannot associate a service and service network that are shared with a
// caller. The caller must own either the service or the service network.
//
// As a result of this operation, the association is created in the service
// network account and the association owner account.
//
// [Manage service associations]: https://docs.aws.amazon.com/vpc-lattice/latest/ug/service-network-associations.html#service-network-service-associations
func vpclattice_CreateServiceNetworkServiceAssociation(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.CreateServiceNetworkServiceAssociationInput{
		// ServiceIdentifier: *string, // Required
		// ServiceNetworkIdentifier: *string, // Required
	}

	if len(_vpclatticeServiceIdentifier) > 0 {
		input.ServiceIdentifier = aws.String(_vpclatticeServiceIdentifier)
	}
	if len(_vpclatticeServiceNetworkIdentifier) > 0 {
		input.ServiceNetworkIdentifier = aws.String(_vpclatticeServiceNetworkIdentifier)
	}
	if len(_vpclatticeClientToken) > 0 {
		input.ClientToken = aws.String(_vpclatticeClientToken)
	}
	if len(_vpclatticeTags) > 0 {
		if err := assignInputField(input, "Tags", _vpclatticeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateServiceNetworkServiceAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a VPC with a service network. When you associate a VPC with the
// service network, it enables all the resources within that VPC to be clients and
// communicate with other services in the service network. For more information,
// see [Manage VPC associations]in the Amazon VPC Lattice User Guide.
//
// You can't use this operation if there is a disassociation in progress. If the
// association fails, retry by deleting the association and recreating it.
//
// As a result of this operation, the association gets created in the service
// network account and the VPC owner account.
//
// If you add a security group to the service network and VPC association, the
// association must continue to always have at least one security group. You can
// add or edit security groups at any time. However, to remove all security groups,
// you must first delete the association and recreate it without security groups.
//
// [Manage VPC associations]: https://docs.aws.amazon.com/vpc-lattice/latest/ug/service-network-associations.html#service-network-vpc-associations
func vpclattice_CreateServiceNetworkVpcAssociation(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.CreateServiceNetworkVpcAssociationInput{
		// ServiceNetworkIdentifier: *string, // Required
		// VpcIdentifier: *string, // Required
	}

	if len(_vpclatticeServiceNetworkIdentifier) > 0 {
		input.ServiceNetworkIdentifier = aws.String(_vpclatticeServiceNetworkIdentifier)
	}
	if len(_vpclatticeVpcIdentifier) > 0 {
		input.VpcIdentifier = aws.String(_vpclatticeVpcIdentifier)
	}
	if len(_vpclatticeClientToken) > 0 {
		input.ClientToken = aws.String(_vpclatticeClientToken)
	}
	if len(_vpclatticeDnsOptions) > 0 {
		if err := assignInputField(input, "DnsOptions", _vpclatticeDnsOptions); err != nil {
			log.Errorf("invalid --dns-options: %s", err.Error())
			return
		}
	}
	if len(_vpclatticePrivateDnsEnabled) > 0 {
		if err := assignInputField(input, "PrivateDnsEnabled", _vpclatticePrivateDnsEnabled); err != nil {
			log.Errorf("invalid --private-dns-enabled: %s", err.Error())
			return
		}
	}
	if len(_vpclatticeSecurityGroupIds) > 0 {
		input.SecurityGroupIds = append([]string(nil), _vpclatticeSecurityGroupIds...)
	}
	if len(_vpclatticeTags) > 0 {
		if err := assignInputField(input, "Tags", _vpclatticeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateServiceNetworkVpcAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a target group. A target group is a collection of targets, or compute
// resources, that run your application or service. A target group can only be used
// by a single service.
//
// For more information, see [Target groups] in the Amazon VPC Lattice User Guide.
//
// [Target groups]: https://docs.aws.amazon.com/vpc-lattice/latest/ug/target-groups.html
func vpclattice_CreateTargetGroup(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.CreateTargetGroupInput{
		// Name: *string, // Required
		// Type: types.TargetGroupType, // Required
	}

	if len(_vpclatticeName) > 0 {
		input.Name = aws.String(_vpclatticeName)
	}
	if len(_vpclatticeType) > 0 {
		if err := assignInputField(input, "Type", _vpclatticeType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_vpclatticeClientToken) > 0 {
		input.ClientToken = aws.String(_vpclatticeClientToken)
	}
	if len(_vpclatticeConfig) > 0 {
		if err := assignInputField(input, "Config", _vpclatticeConfig); err != nil {
			log.Errorf("invalid --config: %s", err.Error())
			return
		}
	}
	if len(_vpclatticeTags) > 0 {
		if err := assignInputField(input, "Tags", _vpclatticeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTargetGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified access log subscription.
func vpclattice_DeleteAccessLogSubscription(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.DeleteAccessLogSubscriptionInput{
		// AccessLogSubscriptionIdentifier: *string, // Required
	}

	if len(_vpclatticeAccessLogSubscriptionIdentifier) > 0 {
		input.AccessLogSubscriptionIdentifier = aws.String(_vpclatticeAccessLogSubscriptionIdentifier)
	}

	if resp, err := client.DeleteAccessLogSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified auth policy. If an auth is set to AWS_IAM and the auth
// policy is deleted, all requests are denied. If you are trying to remove the auth
// policy completely, you must set the auth type to NONE . If auth is enabled on
// the resource, but no auth policy is set, all requests are denied.
func vpclattice_DeleteAuthPolicy(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.DeleteAuthPolicyInput{
		// ResourceIdentifier: *string, // Required
	}

	if len(_vpclatticeResourceIdentifier) > 0 {
		input.ResourceIdentifier = aws.String(_vpclatticeResourceIdentifier)
	}

	if resp, err := client.DeleteAuthPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified domain verification.
func vpclattice_DeleteDomainVerification(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.DeleteDomainVerificationInput{
		// DomainVerificationIdentifier: *string, // Required
	}

	if len(_vpclatticeDomainVerificationIdentifier) > 0 {
		input.DomainVerificationIdentifier = aws.String(_vpclatticeDomainVerificationIdentifier)
	}

	if resp, err := client.DeleteDomainVerification(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified listener.
func vpclattice_DeleteListener(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.DeleteListenerInput{
		// ListenerIdentifier: *string, // Required
		// ServiceIdentifier: *string, // Required
	}

	if len(_vpclatticeListenerIdentifier) > 0 {
		input.ListenerIdentifier = aws.String(_vpclatticeListenerIdentifier)
	}
	if len(_vpclatticeServiceIdentifier) > 0 {
		input.ServiceIdentifier = aws.String(_vpclatticeServiceIdentifier)
	}

	if resp, err := client.DeleteListener(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified resource configuration.
func vpclattice_DeleteResourceConfiguration(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.DeleteResourceConfigurationInput{
		// ResourceConfigurationIdentifier: *string, // Required
	}

	if len(_vpclatticeResourceConfigurationIdentifier) > 0 {
		input.ResourceConfigurationIdentifier = aws.String(_vpclatticeResourceConfigurationIdentifier)
	}

	if resp, err := client.DeleteResourceConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates the resource configuration from the resource VPC endpoint.
func vpclattice_DeleteResourceEndpointAssociation(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.DeleteResourceEndpointAssociationInput{
		// ResourceEndpointAssociationIdentifier: *string, // Required
	}

	if len(_vpclatticeResourceEndpointAssociationIdentifier) > 0 {
		input.ResourceEndpointAssociationIdentifier = aws.String(_vpclatticeResourceEndpointAssociationIdentifier)
	}

	if resp, err := client.DeleteResourceEndpointAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified resource gateway.
func vpclattice_DeleteResourceGateway(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.DeleteResourceGatewayInput{
		// ResourceGatewayIdentifier: *string, // Required
	}

	if len(_vpclatticeResourceGatewayIdentifier) > 0 {
		input.ResourceGatewayIdentifier = aws.String(_vpclatticeResourceGatewayIdentifier)
	}

	if resp, err := client.DeleteResourceGateway(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified resource policy.
func vpclattice_DeleteResourcePolicy(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.DeleteResourcePolicyInput{
		// ResourceArn: *string, // Required
	}

	if len(_vpclatticeResourceArn) > 0 {
		input.ResourceArn = aws.String(_vpclatticeResourceArn)
	}

	if resp, err := client.DeleteResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a listener rule. Each listener has a default rule for checking
// connection requests, but you can define additional rules. Each rule consists of
// a priority, one or more actions, and one or more conditions. You can delete
// additional listener rules, but you cannot delete the default rule.
//
// For more information, see [Listener rules] in the Amazon VPC Lattice User Guide.
//
// [Listener rules]: https://docs.aws.amazon.com/vpc-lattice/latest/ug/listeners.html#listener-rules
func vpclattice_DeleteRule(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.DeleteRuleInput{
		// ListenerIdentifier: *string, // Required
		// RuleIdentifier: *string, // Required
		// ServiceIdentifier: *string, // Required
	}

	if len(_vpclatticeListenerIdentifier) > 0 {
		input.ListenerIdentifier = aws.String(_vpclatticeListenerIdentifier)
	}
	if len(_vpclatticeRuleIdentifier) > 0 {
		input.RuleIdentifier = aws.String(_vpclatticeRuleIdentifier)
	}
	if len(_vpclatticeServiceIdentifier) > 0 {
		input.ServiceIdentifier = aws.String(_vpclatticeServiceIdentifier)
	}

	if resp, err := client.DeleteRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a service. A service can't be deleted if it's associated with a service
// network. If you delete a service, all resources related to the service, such as
// the resource policy, auth policy, listeners, listener rules, and access log
// subscriptions, are also deleted. For more information, see [Delete a service]in the Amazon VPC
// Lattice User Guide.
//
// [Delete a service]: https://docs.aws.amazon.com/vpc-lattice/latest/ug/services.html#delete-service
func vpclattice_DeleteService(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.DeleteServiceInput{
		// ServiceIdentifier: *string, // Required
	}

	if len(_vpclatticeServiceIdentifier) > 0 {
		input.ServiceIdentifier = aws.String(_vpclatticeServiceIdentifier)
	}

	if resp, err := client.DeleteService(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a service network. You can only delete the service network if there is
// no service or VPC associated with it. If you delete a service network, all
// resources related to the service network, such as the resource policy, auth
// policy, and access log subscriptions, are also deleted. For more information,
// see [Delete a service network]in the Amazon VPC Lattice User Guide.
//
// [Delete a service network]: https://docs.aws.amazon.com/vpc-lattice/latest/ug/service-networks.html#delete-service-network
func vpclattice_DeleteServiceNetwork(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.DeleteServiceNetworkInput{
		// ServiceNetworkIdentifier: *string, // Required
	}

	if len(_vpclatticeServiceNetworkIdentifier) > 0 {
		input.ServiceNetworkIdentifier = aws.String(_vpclatticeServiceNetworkIdentifier)
	}

	if resp, err := client.DeleteServiceNetwork(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the association between a service network and a resource configuration.
func vpclattice_DeleteServiceNetworkResourceAssociation(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.DeleteServiceNetworkResourceAssociationInput{
		// ServiceNetworkResourceAssociationIdentifier: *string, // Required
	}

	if len(_vpclatticeServiceNetworkResourceAssociationIdentifier) > 0 {
		input.ServiceNetworkResourceAssociationIdentifier = aws.String(_vpclatticeServiceNetworkResourceAssociationIdentifier)
	}

	if resp, err := client.DeleteServiceNetworkResourceAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the association between a service and a service network. This operation
// fails if an association is still in progress.
func vpclattice_DeleteServiceNetworkServiceAssociation(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.DeleteServiceNetworkServiceAssociationInput{
		// ServiceNetworkServiceAssociationIdentifier: *string, // Required
	}

	if len(_vpclatticeServiceNetworkServiceAssociationIdentifier) > 0 {
		input.ServiceNetworkServiceAssociationIdentifier = aws.String(_vpclatticeServiceNetworkServiceAssociationIdentifier)
	}

	if resp, err := client.DeleteServiceNetworkServiceAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates the VPC from the service network. You can't disassociate the VPC
// if there is a create or update association in progress.
func vpclattice_DeleteServiceNetworkVpcAssociation(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.DeleteServiceNetworkVpcAssociationInput{
		// ServiceNetworkVpcAssociationIdentifier: *string, // Required
	}

	if len(_vpclatticeServiceNetworkVpcAssociationIdentifier) > 0 {
		input.ServiceNetworkVpcAssociationIdentifier = aws.String(_vpclatticeServiceNetworkVpcAssociationIdentifier)
	}

	if resp, err := client.DeleteServiceNetworkVpcAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a target group. You can't delete a target group if it is used in a
// listener rule or if the target group creation is in progress.
func vpclattice_DeleteTargetGroup(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.DeleteTargetGroupInput{
		// TargetGroupIdentifier: *string, // Required
	}

	if len(_vpclatticeTargetGroupIdentifier) > 0 {
		input.TargetGroupIdentifier = aws.String(_vpclatticeTargetGroupIdentifier)
	}

	if resp, err := client.DeleteTargetGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deregisters the specified targets from the specified target group.
func vpclattice_DeregisterTargets(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.DeregisterTargetsInput{
		// TargetGroupIdentifier: *string, // Required
		// Targets: []types.Target, // Required
	}

	if len(_vpclatticeTargetGroupIdentifier) > 0 {
		input.TargetGroupIdentifier = aws.String(_vpclatticeTargetGroupIdentifier)
	}
	if len(_vpclatticeTargets) > 0 {
		if err := assignInputField(input, "Targets", _vpclatticeTargets); err != nil {
			log.Errorf("invalid --targets: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeregisterTargets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the specified access log subscription.
func vpclattice_GetAccessLogSubscription(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.GetAccessLogSubscriptionInput{
		// AccessLogSubscriptionIdentifier: *string, // Required
	}

	if len(_vpclatticeAccessLogSubscriptionIdentifier) > 0 {
		input.AccessLogSubscriptionIdentifier = aws.String(_vpclatticeAccessLogSubscriptionIdentifier)
	}

	if resp, err := client.GetAccessLogSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the auth policy for the specified service or
// service network.
func vpclattice_GetAuthPolicy(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.GetAuthPolicyInput{
		// ResourceIdentifier: *string, // Required
	}

	if len(_vpclatticeResourceIdentifier) > 0 {
		input.ResourceIdentifier = aws.String(_vpclatticeResourceIdentifier)
	}

	if resp, err := client.GetAuthPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a domain verification.ß
func vpclattice_GetDomainVerification(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.GetDomainVerificationInput{
		// DomainVerificationIdentifier: *string, // Required
	}

	if len(_vpclatticeDomainVerificationIdentifier) > 0 {
		input.DomainVerificationIdentifier = aws.String(_vpclatticeDomainVerificationIdentifier)
	}

	if resp, err := client.GetDomainVerification(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the specified listener for the specified service.
func vpclattice_GetListener(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.GetListenerInput{
		// ListenerIdentifier: *string, // Required
		// ServiceIdentifier: *string, // Required
	}

	if len(_vpclatticeListenerIdentifier) > 0 {
		input.ListenerIdentifier = aws.String(_vpclatticeListenerIdentifier)
	}
	if len(_vpclatticeServiceIdentifier) > 0 {
		input.ServiceIdentifier = aws.String(_vpclatticeServiceIdentifier)
	}

	if resp, err := client.GetListener(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the specified resource configuration.
func vpclattice_GetResourceConfiguration(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.GetResourceConfigurationInput{
		// ResourceConfigurationIdentifier: *string, // Required
	}

	if len(_vpclatticeResourceConfigurationIdentifier) > 0 {
		input.ResourceConfigurationIdentifier = aws.String(_vpclatticeResourceConfigurationIdentifier)
	}

	if resp, err := client.GetResourceConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the specified resource gateway.
func vpclattice_GetResourceGateway(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.GetResourceGatewayInput{
		// ResourceGatewayIdentifier: *string, // Required
	}

	if len(_vpclatticeResourceGatewayIdentifier) > 0 {
		input.ResourceGatewayIdentifier = aws.String(_vpclatticeResourceGatewayIdentifier)
	}

	if resp, err := client.GetResourceGateway(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the specified resource policy. The resource policy
// is an IAM policy created on behalf of the resource owner when they share a
// resource.
func vpclattice_GetResourcePolicy(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.GetResourcePolicyInput{
		// ResourceArn: *string, // Required
	}

	if len(_vpclatticeResourceArn) > 0 {
		input.ResourceArn = aws.String(_vpclatticeResourceArn)
	}

	if resp, err := client.GetResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the specified listener rules. You can also retrieve
// information about the default listener rule. For more information, see [Listener rules]in the
// Amazon VPC Lattice User Guide.
//
// [Listener rules]: https://docs.aws.amazon.com/vpc-lattice/latest/ug/listeners.html#listener-rules
func vpclattice_GetRule(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.GetRuleInput{
		// ListenerIdentifier: *string, // Required
		// RuleIdentifier: *string, // Required
		// ServiceIdentifier: *string, // Required
	}

	if len(_vpclatticeListenerIdentifier) > 0 {
		input.ListenerIdentifier = aws.String(_vpclatticeListenerIdentifier)
	}
	if len(_vpclatticeRuleIdentifier) > 0 {
		input.RuleIdentifier = aws.String(_vpclatticeRuleIdentifier)
	}
	if len(_vpclatticeServiceIdentifier) > 0 {
		input.ServiceIdentifier = aws.String(_vpclatticeServiceIdentifier)
	}

	if resp, err := client.GetRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the specified service.
func vpclattice_GetService(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.GetServiceInput{
		// ServiceIdentifier: *string, // Required
	}

	if len(_vpclatticeServiceIdentifier) > 0 {
		input.ServiceIdentifier = aws.String(_vpclatticeServiceIdentifier)
	}

	if resp, err := client.GetService(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the specified service network.
func vpclattice_GetServiceNetwork(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.GetServiceNetworkInput{
		// ServiceNetworkIdentifier: *string, // Required
	}

	if len(_vpclatticeServiceNetworkIdentifier) > 0 {
		input.ServiceNetworkIdentifier = aws.String(_vpclatticeServiceNetworkIdentifier)
	}

	if resp, err := client.GetServiceNetwork(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the specified association between a service network
// and a resource configuration.
func vpclattice_GetServiceNetworkResourceAssociation(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.GetServiceNetworkResourceAssociationInput{
		// ServiceNetworkResourceAssociationIdentifier: *string, // Required
	}

	if len(_vpclatticeServiceNetworkResourceAssociationIdentifier) > 0 {
		input.ServiceNetworkResourceAssociationIdentifier = aws.String(_vpclatticeServiceNetworkResourceAssociationIdentifier)
	}

	if resp, err := client.GetServiceNetworkResourceAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the specified association between a service network
// and a service.
func vpclattice_GetServiceNetworkServiceAssociation(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.GetServiceNetworkServiceAssociationInput{
		// ServiceNetworkServiceAssociationIdentifier: *string, // Required
	}

	if len(_vpclatticeServiceNetworkServiceAssociationIdentifier) > 0 {
		input.ServiceNetworkServiceAssociationIdentifier = aws.String(_vpclatticeServiceNetworkServiceAssociationIdentifier)
	}

	if resp, err := client.GetServiceNetworkServiceAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the specified association between a service network
// and a VPC.
func vpclattice_GetServiceNetworkVpcAssociation(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.GetServiceNetworkVpcAssociationInput{
		// ServiceNetworkVpcAssociationIdentifier: *string, // Required
	}

	if len(_vpclatticeServiceNetworkVpcAssociationIdentifier) > 0 {
		input.ServiceNetworkVpcAssociationIdentifier = aws.String(_vpclatticeServiceNetworkVpcAssociationIdentifier)
	}

	if resp, err := client.GetServiceNetworkVpcAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the specified target group.
func vpclattice_GetTargetGroup(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.GetTargetGroupInput{
		// TargetGroupIdentifier: *string, // Required
	}

	if len(_vpclatticeTargetGroupIdentifier) > 0 {
		input.TargetGroupIdentifier = aws.String(_vpclatticeTargetGroupIdentifier)
	}

	if resp, err := client.GetTargetGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the access log subscriptions for the specified service network or service.
func vpclattice_ListAccessLogSubscriptions(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.ListAccessLogSubscriptionsInput{
		// ResourceIdentifier: *string, // Required
	}

	if len(_vpclatticeResourceIdentifier) > 0 {
		input.ResourceIdentifier = aws.String(_vpclatticeResourceIdentifier)
	}
	if len(_vpclatticeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _vpclatticeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_vpclatticeNextToken) > 0 {
		input.NextToken = aws.String(_vpclatticeNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAccessLogSubscriptions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*vpclattice.ListAccessLogSubscriptionsOutput
	p := vpclattice.NewListAccessLogSubscriptionsPaginator(client, input)
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

// Lists the domain verifications.
func vpclattice_ListDomainVerifications(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.ListDomainVerificationsInput{}

	if len(_vpclatticeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _vpclatticeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_vpclatticeNextToken) > 0 {
		input.NextToken = aws.String(_vpclatticeNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDomainVerifications(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*vpclattice.ListDomainVerificationsOutput
	p := vpclattice.NewListDomainVerificationsPaginator(client, input)
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

// Lists the listeners for the specified service.
func vpclattice_ListListeners(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.ListListenersInput{
		// ServiceIdentifier: *string, // Required
	}

	if len(_vpclatticeServiceIdentifier) > 0 {
		input.ServiceIdentifier = aws.String(_vpclatticeServiceIdentifier)
	}
	if len(_vpclatticeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _vpclatticeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_vpclatticeNextToken) > 0 {
		input.NextToken = aws.String(_vpclatticeNextToken)
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

	var results []*vpclattice.ListListenersOutput
	p := vpclattice.NewListListenersPaginator(client, input)
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

// Lists the resource configurations owned by or shared with this account.
func vpclattice_ListResourceConfigurations(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.ListResourceConfigurationsInput{}

	if len(_vpclatticeDomainVerificationIdentifier) > 0 {
		input.DomainVerificationIdentifier = aws.String(_vpclatticeDomainVerificationIdentifier)
	}
	if len(_vpclatticeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _vpclatticeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_vpclatticeNextToken) > 0 {
		input.NextToken = aws.String(_vpclatticeNextToken)
	}
	if len(_vpclatticeResourceConfigurationGroupIdentifier) > 0 {
		input.ResourceConfigurationGroupIdentifier = aws.String(_vpclatticeResourceConfigurationGroupIdentifier)
	}
	if len(_vpclatticeResourceGatewayIdentifier) > 0 {
		input.ResourceGatewayIdentifier = aws.String(_vpclatticeResourceGatewayIdentifier)
	}

	if disablePaginator() {
		if resp, err := client.ListResourceConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*vpclattice.ListResourceConfigurationsOutput
	p := vpclattice.NewListResourceConfigurationsPaginator(client, input)
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

// Lists the associations for the specified VPC endpoint.
func vpclattice_ListResourceEndpointAssociations(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.ListResourceEndpointAssociationsInput{
		// ResourceConfigurationIdentifier: *string, // Required
	}

	if len(_vpclatticeResourceConfigurationIdentifier) > 0 {
		input.ResourceConfigurationIdentifier = aws.String(_vpclatticeResourceConfigurationIdentifier)
	}
	if len(_vpclatticeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _vpclatticeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_vpclatticeNextToken) > 0 {
		input.NextToken = aws.String(_vpclatticeNextToken)
	}
	if len(_vpclatticeResourceEndpointAssociationIdentifier) > 0 {
		input.ResourceEndpointAssociationIdentifier = aws.String(_vpclatticeResourceEndpointAssociationIdentifier)
	}
	if len(_vpclatticeVpcEndpointId) > 0 {
		input.VpcEndpointId = aws.String(_vpclatticeVpcEndpointId)
	}
	if len(_vpclatticeVpcEndpointOwner) > 0 {
		input.VpcEndpointOwner = aws.String(_vpclatticeVpcEndpointOwner)
	}

	if disablePaginator() {
		if resp, err := client.ListResourceEndpointAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*vpclattice.ListResourceEndpointAssociationsOutput
	p := vpclattice.NewListResourceEndpointAssociationsPaginator(client, input)
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

// Lists the resource gateways that you own or that were shared with you.
func vpclattice_ListResourceGateways(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.ListResourceGatewaysInput{}

	if len(_vpclatticeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _vpclatticeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_vpclatticeNextToken) > 0 {
		input.NextToken = aws.String(_vpclatticeNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListResourceGateways(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*vpclattice.ListResourceGatewaysOutput
	p := vpclattice.NewListResourceGatewaysPaginator(client, input)
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

// Lists the rules for the specified listener.
func vpclattice_ListRules(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.ListRulesInput{
		// ListenerIdentifier: *string, // Required
		// ServiceIdentifier: *string, // Required
	}

	if len(_vpclatticeListenerIdentifier) > 0 {
		input.ListenerIdentifier = aws.String(_vpclatticeListenerIdentifier)
	}
	if len(_vpclatticeServiceIdentifier) > 0 {
		input.ServiceIdentifier = aws.String(_vpclatticeServiceIdentifier)
	}
	if len(_vpclatticeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _vpclatticeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_vpclatticeNextToken) > 0 {
		input.NextToken = aws.String(_vpclatticeNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRules(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*vpclattice.ListRulesOutput
	p := vpclattice.NewListRulesPaginator(client, input)
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

// Lists the associations between a service network and a resource configuration.
func vpclattice_ListServiceNetworkResourceAssociations(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.ListServiceNetworkResourceAssociationsInput{}

	if len(_vpclatticeIncludeChildren) > 0 {
		if err := assignInputField(input, "IncludeChildren", _vpclatticeIncludeChildren); err != nil {
			log.Errorf("invalid --include-children: %s", err.Error())
			return
		}
	}
	if len(_vpclatticeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _vpclatticeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_vpclatticeNextToken) > 0 {
		input.NextToken = aws.String(_vpclatticeNextToken)
	}
	if len(_vpclatticeResourceConfigurationIdentifier) > 0 {
		input.ResourceConfigurationIdentifier = aws.String(_vpclatticeResourceConfigurationIdentifier)
	}
	if len(_vpclatticeServiceNetworkIdentifier) > 0 {
		input.ServiceNetworkIdentifier = aws.String(_vpclatticeServiceNetworkIdentifier)
	}

	if disablePaginator() {
		if resp, err := client.ListServiceNetworkResourceAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*vpclattice.ListServiceNetworkResourceAssociationsOutput
	p := vpclattice.NewListServiceNetworkResourceAssociationsPaginator(client, input)
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

// Lists the associations between a service network and a service. You can filter
// the list either by service or service network. You must provide either the
// service network identifier or the service identifier.
//
// Every association in Amazon VPC Lattice has a unique Amazon Resource Name
// (ARN), such as when a service network is associated with a VPC or when a service
// is associated with a service network. If the association is for a resource is
// shared with another account, the association includes the local account ID as
// the prefix in the ARN.
func vpclattice_ListServiceNetworkServiceAssociations(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.ListServiceNetworkServiceAssociationsInput{}

	if len(_vpclatticeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _vpclatticeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_vpclatticeNextToken) > 0 {
		input.NextToken = aws.String(_vpclatticeNextToken)
	}
	if len(_vpclatticeServiceIdentifier) > 0 {
		input.ServiceIdentifier = aws.String(_vpclatticeServiceIdentifier)
	}
	if len(_vpclatticeServiceNetworkIdentifier) > 0 {
		input.ServiceNetworkIdentifier = aws.String(_vpclatticeServiceNetworkIdentifier)
	}

	if disablePaginator() {
		if resp, err := client.ListServiceNetworkServiceAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*vpclattice.ListServiceNetworkServiceAssociationsOutput
	p := vpclattice.NewListServiceNetworkServiceAssociationsPaginator(client, input)
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

// Lists the associations between a service network and a VPC. You can filter the
// list either by VPC or service network. You must provide either the ID of the
// service network identifier or the ID of the VPC.
func vpclattice_ListServiceNetworkVpcAssociations(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.ListServiceNetworkVpcAssociationsInput{}

	if len(_vpclatticeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _vpclatticeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_vpclatticeNextToken) > 0 {
		input.NextToken = aws.String(_vpclatticeNextToken)
	}
	if len(_vpclatticeServiceNetworkIdentifier) > 0 {
		input.ServiceNetworkIdentifier = aws.String(_vpclatticeServiceNetworkIdentifier)
	}
	if len(_vpclatticeVpcIdentifier) > 0 {
		input.VpcIdentifier = aws.String(_vpclatticeVpcIdentifier)
	}

	if disablePaginator() {
		if resp, err := client.ListServiceNetworkVpcAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*vpclattice.ListServiceNetworkVpcAssociationsOutput
	p := vpclattice.NewListServiceNetworkVpcAssociationsPaginator(client, input)
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

// Lists the associations between a service network and a VPC endpoint.
func vpclattice_ListServiceNetworkVpcEndpointAssociations(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.ListServiceNetworkVpcEndpointAssociationsInput{
		// ServiceNetworkIdentifier: *string, // Required
	}

	if len(_vpclatticeServiceNetworkIdentifier) > 0 {
		input.ServiceNetworkIdentifier = aws.String(_vpclatticeServiceNetworkIdentifier)
	}
	if len(_vpclatticeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _vpclatticeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_vpclatticeNextToken) > 0 {
		input.NextToken = aws.String(_vpclatticeNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListServiceNetworkVpcEndpointAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*vpclattice.ListServiceNetworkVpcEndpointAssociationsOutput
	p := vpclattice.NewListServiceNetworkVpcEndpointAssociationsPaginator(client, input)
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

// Lists the service networks owned by or shared with this account. The account ID
// in the ARN shows which account owns the service network.
func vpclattice_ListServiceNetworks(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.ListServiceNetworksInput{}

	if len(_vpclatticeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _vpclatticeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_vpclatticeNextToken) > 0 {
		input.NextToken = aws.String(_vpclatticeNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListServiceNetworks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*vpclattice.ListServiceNetworksOutput
	p := vpclattice.NewListServiceNetworksPaginator(client, input)
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

// Lists the services owned by the caller account or shared with the caller
// account.
func vpclattice_ListServices(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.ListServicesInput{}

	if len(_vpclatticeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _vpclatticeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_vpclatticeNextToken) > 0 {
		input.NextToken = aws.String(_vpclatticeNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListServices(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*vpclattice.ListServicesOutput
	p := vpclattice.NewListServicesPaginator(client, input)
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

// Lists the tags for the specified resource.
func vpclattice_ListTagsForResource(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_vpclatticeResourceArn) > 0 {
		input.ResourceArn = aws.String(_vpclatticeResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists your target groups. You can narrow your search by using the filters below
// in your request.
func vpclattice_ListTargetGroups(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.ListTargetGroupsInput{}

	if len(_vpclatticeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _vpclatticeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_vpclatticeNextToken) > 0 {
		input.NextToken = aws.String(_vpclatticeNextToken)
	}
	if len(_vpclatticeTargetGroupType) > 0 {
		if err := assignInputField(input, "TargetGroupType", _vpclatticeTargetGroupType); err != nil {
			log.Errorf("invalid --target-group-type: %s", err.Error())
			return
		}
	}
	if len(_vpclatticeVpcIdentifier) > 0 {
		input.VpcIdentifier = aws.String(_vpclatticeVpcIdentifier)
	}

	if disablePaginator() {
		if resp, err := client.ListTargetGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*vpclattice.ListTargetGroupsOutput
	p := vpclattice.NewListTargetGroupsPaginator(client, input)
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

// Lists the targets for the target group. By default, all targets are included.
// You can use this API to check the health status of targets. You can also ﬁlter
// the results by target.
func vpclattice_ListTargets(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.ListTargetsInput{
		// TargetGroupIdentifier: *string, // Required
	}

	if len(_vpclatticeTargetGroupIdentifier) > 0 {
		input.TargetGroupIdentifier = aws.String(_vpclatticeTargetGroupIdentifier)
	}
	if len(_vpclatticeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _vpclatticeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_vpclatticeNextToken) > 0 {
		input.NextToken = aws.String(_vpclatticeNextToken)
	}
	if len(_vpclatticeTargets) > 0 {
		if err := assignInputField(input, "Targets", _vpclatticeTargets); err != nil {
			log.Errorf("invalid --targets: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListTargets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*vpclattice.ListTargetsOutput
	p := vpclattice.NewListTargetsPaginator(client, input)
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

// Creates or updates the auth policy. The policy string in JSON must not contain
// newlines or blank lines.
//
// For more information, see [Auth policies] in the Amazon VPC Lattice User Guide.
//
// [Auth policies]: https://docs.aws.amazon.com/vpc-lattice/latest/ug/auth-policies.html
func vpclattice_PutAuthPolicy(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.PutAuthPolicyInput{
		// Policy: *string, // Required
		// ResourceIdentifier: *string, // Required
	}

	if len(_vpclatticePolicy) > 0 {
		input.Policy = aws.String(_vpclatticePolicy)
	}
	if len(_vpclatticeResourceIdentifier) > 0 {
		input.ResourceIdentifier = aws.String(_vpclatticeResourceIdentifier)
	}

	if resp, err := client.PutAuthPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attaches a resource-based permission policy to a service or service network.
// The policy must contain the same actions and condition statements as the Amazon
// Web Services Resource Access Manager permission for sharing services and service
// networks.
func vpclattice_PutResourcePolicy(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.PutResourcePolicyInput{
		// Policy: *string, // Required
		// ResourceArn: *string, // Required
	}

	if len(_vpclatticePolicy) > 0 {
		input.Policy = aws.String(_vpclatticePolicy)
	}
	if len(_vpclatticeResourceArn) > 0 {
		input.ResourceArn = aws.String(_vpclatticeResourceArn)
	}

	if resp, err := client.PutResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Registers the targets with the target group. If it's a Lambda target, you can
// only have one target in a target group.
func vpclattice_RegisterTargets(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.RegisterTargetsInput{
		// TargetGroupIdentifier: *string, // Required
		// Targets: []types.Target, // Required
	}

	if len(_vpclatticeTargetGroupIdentifier) > 0 {
		input.TargetGroupIdentifier = aws.String(_vpclatticeTargetGroupIdentifier)
	}
	if len(_vpclatticeTargets) > 0 {
		if err := assignInputField(input, "Targets", _vpclatticeTargets); err != nil {
			log.Errorf("invalid --targets: %s", err.Error())
			return
		}
	}

	if resp, err := client.RegisterTargets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts the domain verification process for a custom domain name.
func vpclattice_StartDomainVerification(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.StartDomainVerificationInput{
		// DomainName: *string, // Required
	}

	if len(_vpclatticeDomainName) > 0 {
		input.DomainName = aws.String(_vpclatticeDomainName)
	}
	if len(_vpclatticeClientToken) > 0 {
		input.ClientToken = aws.String(_vpclatticeClientToken)
	}
	if len(_vpclatticeTags) > 0 {
		if err := assignInputField(input, "Tags", _vpclatticeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartDomainVerification(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds the specified tags to the specified resource.
func vpclattice_TagResource(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_vpclatticeResourceArn) > 0 {
		input.ResourceArn = aws.String(_vpclatticeResourceArn)
	}
	if len(_vpclatticeTags) > 0 {
		if err := assignInputField(input, "Tags", _vpclatticeTags); err != nil {
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

// Removes the specified tags from the specified resource.
func vpclattice_UntagResource(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_vpclatticeResourceArn) > 0 {
		input.ResourceArn = aws.String(_vpclatticeResourceArn)
	}
	if len(_vpclatticeTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _vpclatticeTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified access log subscription.
func vpclattice_UpdateAccessLogSubscription(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.UpdateAccessLogSubscriptionInput{
		// AccessLogSubscriptionIdentifier: *string, // Required
		// DestinationArn: *string, // Required
	}

	if len(_vpclatticeAccessLogSubscriptionIdentifier) > 0 {
		input.AccessLogSubscriptionIdentifier = aws.String(_vpclatticeAccessLogSubscriptionIdentifier)
	}
	if len(_vpclatticeDestinationArn) > 0 {
		input.DestinationArn = aws.String(_vpclatticeDestinationArn)
	}

	if resp, err := client.UpdateAccessLogSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified listener for the specified service.
func vpclattice_UpdateListener(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.UpdateListenerInput{
		// DefaultAction: types.RuleAction, // Required
		// ListenerIdentifier: *string, // Required
		// ServiceIdentifier: *string, // Required
	}

	if len(_vpclatticeDefaultAction) > 0 {
		if err := assignInputField(input, "DefaultAction", _vpclatticeDefaultAction); err != nil {
			log.Errorf("invalid --default-action: %s", err.Error())
			return
		}
	}
	if len(_vpclatticeListenerIdentifier) > 0 {
		input.ListenerIdentifier = aws.String(_vpclatticeListenerIdentifier)
	}
	if len(_vpclatticeServiceIdentifier) > 0 {
		input.ServiceIdentifier = aws.String(_vpclatticeServiceIdentifier)
	}

	if resp, err := client.UpdateListener(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified resource configuration.
func vpclattice_UpdateResourceConfiguration(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.UpdateResourceConfigurationInput{
		// ResourceConfigurationIdentifier: *string, // Required
	}

	if len(_vpclatticeResourceConfigurationIdentifier) > 0 {
		input.ResourceConfigurationIdentifier = aws.String(_vpclatticeResourceConfigurationIdentifier)
	}
	if len(_vpclatticeAllowAssociationToShareableServiceNetwork) > 0 {
		if err := assignInputField(input, "AllowAssociationToShareableServiceNetwork", _vpclatticeAllowAssociationToShareableServiceNetwork); err != nil {
			log.Errorf("invalid --allow-association-to-shareable-service-network: %s", err.Error())
			return
		}
	}
	if len(_vpclatticePortRanges) > 0 {
		input.PortRanges = append([]string(nil), _vpclatticePortRanges...)
	}
	if len(_vpclatticeResourceConfigurationDefinition) > 0 {
		if err := assignInputField(input, "ResourceConfigurationDefinition", _vpclatticeResourceConfigurationDefinition); err != nil {
			log.Errorf("invalid --resource-configuration-definition: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateResourceConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified resource gateway.
func vpclattice_UpdateResourceGateway(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.UpdateResourceGatewayInput{
		// ResourceGatewayIdentifier: *string, // Required
	}

	if len(_vpclatticeResourceGatewayIdentifier) > 0 {
		input.ResourceGatewayIdentifier = aws.String(_vpclatticeResourceGatewayIdentifier)
	}
	if len(_vpclatticeSecurityGroupIds) > 0 {
		input.SecurityGroupIds = append([]string(nil), _vpclatticeSecurityGroupIds...)
	}

	if resp, err := client.UpdateResourceGateway(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a specified rule for the listener. You can't modify a default listener
// rule. To modify a default listener rule, use UpdateListener .
func vpclattice_UpdateRule(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.UpdateRuleInput{
		// ListenerIdentifier: *string, // Required
		// RuleIdentifier: *string, // Required
		// ServiceIdentifier: *string, // Required
	}

	if len(_vpclatticeListenerIdentifier) > 0 {
		input.ListenerIdentifier = aws.String(_vpclatticeListenerIdentifier)
	}
	if len(_vpclatticeRuleIdentifier) > 0 {
		input.RuleIdentifier = aws.String(_vpclatticeRuleIdentifier)
	}
	if len(_vpclatticeServiceIdentifier) > 0 {
		input.ServiceIdentifier = aws.String(_vpclatticeServiceIdentifier)
	}
	if len(_vpclatticeAction) > 0 {
		if err := assignInputField(input, "Action", _vpclatticeAction); err != nil {
			log.Errorf("invalid --action: %s", err.Error())
			return
		}
	}
	if len(_vpclatticeMatch) > 0 {
		if err := assignInputField(input, "Match", _vpclatticeMatch); err != nil {
			log.Errorf("invalid --match: %s", err.Error())
			return
		}
	}
	if len(_vpclatticePriority) > 0 {
		if err := assignInputField(input, "Priority", _vpclatticePriority); err != nil {
			log.Errorf("invalid --priority: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified service.
func vpclattice_UpdateService(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.UpdateServiceInput{
		// ServiceIdentifier: *string, // Required
	}

	if len(_vpclatticeServiceIdentifier) > 0 {
		input.ServiceIdentifier = aws.String(_vpclatticeServiceIdentifier)
	}
	if len(_vpclatticeAuthType) > 0 {
		if err := assignInputField(input, "AuthType", _vpclatticeAuthType); err != nil {
			log.Errorf("invalid --auth-type: %s", err.Error())
			return
		}
	}
	if len(_vpclatticeCertificateArn) > 0 {
		input.CertificateArn = aws.String(_vpclatticeCertificateArn)
	}

	if resp, err := client.UpdateService(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified service network.
func vpclattice_UpdateServiceNetwork(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.UpdateServiceNetworkInput{
		// AuthType: types.AuthType, // Required
		// ServiceNetworkIdentifier: *string, // Required
	}

	if len(_vpclatticeAuthType) > 0 {
		if err := assignInputField(input, "AuthType", _vpclatticeAuthType); err != nil {
			log.Errorf("invalid --auth-type: %s", err.Error())
			return
		}
	}
	if len(_vpclatticeServiceNetworkIdentifier) > 0 {
		input.ServiceNetworkIdentifier = aws.String(_vpclatticeServiceNetworkIdentifier)
	}

	if resp, err := client.UpdateServiceNetwork(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the service network and VPC association. If you add a security group to
// the service network and VPC association, the association must continue to have
// at least one security group. You can add or edit security groups at any time.
// However, to remove all security groups, you must first delete the association
// and then recreate it without security groups.
func vpclattice_UpdateServiceNetworkVpcAssociation(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.UpdateServiceNetworkVpcAssociationInput{
		// SecurityGroupIds: []string, // Required
		// ServiceNetworkVpcAssociationIdentifier: *string, // Required
	}

	if len(_vpclatticeSecurityGroupIds) > 0 {
		input.SecurityGroupIds = append([]string(nil), _vpclatticeSecurityGroupIds...)
	}
	if len(_vpclatticeServiceNetworkVpcAssociationIdentifier) > 0 {
		input.ServiceNetworkVpcAssociationIdentifier = aws.String(_vpclatticeServiceNetworkVpcAssociationIdentifier)
	}

	if resp, err := client.UpdateServiceNetworkVpcAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified target group.
func vpclattice_UpdateTargetGroup(cfg aws.Config, client *vpclattice.Client) {
	input := &vpclattice.UpdateTargetGroupInput{
		// HealthCheck: *types.HealthCheckConfig, // Required
		// TargetGroupIdentifier: *string, // Required
	}

	if len(_vpclatticeHealthCheck) > 0 {
		if err := assignInputField(input, "HealthCheck", _vpclatticeHealthCheck); err != nil {
			log.Errorf("invalid --health-check: %s", err.Error())
			return
		}
	}
	if len(_vpclatticeTargetGroupIdentifier) > 0 {
		input.TargetGroupIdentifier = aws.String(_vpclatticeTargetGroupIdentifier)
	}

	if resp, err := client.UpdateTargetGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_vpclatticeCmd)
	_vpclatticeCmd.Flags().SortFlags = false

	_vpclatticeCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_vpclatticeCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_vpclatticeCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_vpclatticeCmd.Flags().StringVarP(&_vpclatticeAccessLogSubscriptionIdentifier, "access-log-subscription-identifier", "", "", "Access Log Subscription Identifier")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticeAction, "action", "", "", "Action")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticeAllowAssociationToShareableServiceNetwork, "allow-association-to-shareable-service-network", "", "", "Allow Association To Shareable Service Network")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticeAuthType, "auth-type", "", "", "Auth Type")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticeCertificateArn, "certificate-arn", "", "", "Certificate ARN")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticeClientToken, "client-token", "", "", "Client Token")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticeConfig, "config", "", "", "Config")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticeCustomDomainName, "custom-domain-name", "", "", "Custom Domain Name")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticeDefaultAction, "default-action", "", "", "Default Action")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticeDestinationArn, "destination-arn", "", "", "Destination ARN")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticeDnsOptions, "dns-options", "", "", "DNS Options")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticeDomainName, "domain-name", "", "", "Domain Name")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticeDomainVerificationIdentifier, "domain-verification-identifier", "", "", "Domain Verification Identifier")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticeGroupDomain, "group-domain", "", "", "Group Domain")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticeHealthCheck, "health-check", "", "", "Health Check")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticeIncludeChildren, "include-children", "", "", "Include Children")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticeIpAddressType, "ip-address-type", "", "", "IP Address Type")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticeIpv4AddressesPerEni, "ipv4-addresses-per-eni", "", "", "IPV4 Addresses Per Eni")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticeListenerIdentifier, "listener-identifier", "", "", "Listener Identifier")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticeMatch, "match", "", "", "Match")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticeMaxResults, "max-results", "", "", "Max Results")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticeName, "name", "", "", "Name")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticeNextToken, "next-token", "", "", "Next Token")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticePolicy, "policy", "", "", "Policy")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticePort, "port", "", "", "Port")
	_vpclatticeCmd.Flags().StringSliceVarP(&_vpclatticePortRanges, "port-ranges", "", nil, "Port Ranges")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticePriority, "priority", "", "", "Priority")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticePrivateDnsEnabled, "private-dns-enabled", "", "", "Private DNS Enabled")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticeProtocol, "protocol", "", "", "Protocol")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticeResourceArn, "resource-arn", "", "", "Resource ARN")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticeResourceConfigurationDefinition, "resource-configuration-definition", "", "", "Resource Configuration Definition")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticeResourceConfigurationGroupIdentifier, "resource-configuration-group-identifier", "", "", "Resource Configuration Group Identifier")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticeResourceConfigurationIdentifier, "resource-configuration-identifier", "", "", "Resource Configuration Identifier")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticeResourceEndpointAssociationIdentifier, "resource-endpoint-association-identifier", "", "", "Resource Endpoint Association Identifier")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticeResourceGatewayIdentifier, "resource-gateway-identifier", "", "", "Resource Gateway Identifier")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticeResourceIdentifier, "resource-identifier", "", "", "Resource Identifier")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticeRuleIdentifier, "rule-identifier", "", "", "Rule Identifier")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticeRules, "rules", "", "", "Rules")
	_vpclatticeCmd.Flags().StringSliceVarP(&_vpclatticeSecurityGroupIds, "security-group-ids", "", nil, "Security Group Ids")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticeServiceIdentifier, "service-identifier", "", "", "Service Identifier")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticeServiceNetworkIdentifier, "service-network-identifier", "", "", "Service Network Identifier")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticeServiceNetworkLogType, "service-network-log-type", "", "", "Service Network Log Type")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticeServiceNetworkResourceAssociationIdentifier, "service-network-resource-association-identifier", "", "", "Service Network Resource Association Identifier")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticeServiceNetworkServiceAssociationIdentifier, "service-network-service-association-identifier", "", "", "Service Network Service Association Identifier")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticeServiceNetworkVpcAssociationIdentifier, "service-network-vpc-association-identifier", "", "", "Service Network VPC Association Identifier")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticeSharingConfig, "sharing-config", "", "", "Sharing Config")
	_vpclatticeCmd.Flags().StringSliceVarP(&_vpclatticeSubnetIds, "subnet-ids", "", nil, "Subnet Ids")
	_vpclatticeCmd.Flags().StringSliceVarP(&_vpclatticeTagKeys, "tag-keys", "", nil, "Tag Keys")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticeTags, "tags", "", "", "Tags")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticeTargetGroupIdentifier, "target-group-identifier", "", "", "Target Group Identifier")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticeTargetGroupType, "target-group-type", "", "", "Target Group Type")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticeTargets, "targets", "", "", "Targets")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticeType, "type", "", "", "Type")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticeVpcEndpointId, "vpc-endpoint-id", "", "", "VPC Endpoint ID")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticeVpcEndpointOwner, "vpc-endpoint-owner", "", "", "VPC Endpoint Owner")
	_vpclatticeCmd.Flags().StringVarP(&_vpclatticeVpcIdentifier, "vpc-identifier", "", "", "VPC Identifier")

	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeBatchUpdateRule, "batch-update-rule", "", false, "Batch Update Rule")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeCreateAccessLogSubscription, "create-access-log-subscription", "", false, "Create Access Log Subscription")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeCreateListener, "create-listener", "", false, "Create Listener")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeCreateResourceConfiguration, "create-resource-configuration", "", false, "Create Resource Configuration")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeCreateResourceGateway, "create-resource-gateway", "", false, "Create Resource Gateway")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeCreateRule, "create-rule", "", false, "Create Rule")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeCreateService, "create-service", "", false, "Create Service")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeCreateServiceNetwork, "create-service-network", "", false, "Create Service Network")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeCreateServiceNetworkResourceAssociation, "create-service-network-resource-association", "", false, "Create Service Network Resource Association")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeCreateServiceNetworkServiceAssociation, "create-service-network-service-association", "", false, "Create Service Network Service Association")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeCreateServiceNetworkVpcAssociation, "create-service-network-vpc-association", "", false, "Create Service Network VPC Association")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeCreateTargetGroup, "create-target-group", "", false, "Create Target Group")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeDeleteAccessLogSubscription, "delete-access-log-subscription", "", false, "Delete Access Log Subscription")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeDeleteAuthPolicy, "delete-auth-policy", "", false, "Delete Auth Policy")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeDeleteDomainVerification, "delete-domain-verification", "", false, "Delete Domain Verification")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeDeleteListener, "delete-listener", "", false, "Delete Listener")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeDeleteResourceConfiguration, "delete-resource-configuration", "", false, "Delete Resource Configuration")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeDeleteResourceEndpointAssociation, "delete-resource-endpoint-association", "", false, "Delete Resource Endpoint Association")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeDeleteResourceGateway, "delete-resource-gateway", "", false, "Delete Resource Gateway")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeDeleteResourcePolicy, "delete-resource-policy", "", false, "Delete Resource Policy")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeDeleteRule, "delete-rule", "", false, "Delete Rule")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeDeleteService, "delete-service", "", false, "Delete Service")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeDeleteServiceNetwork, "delete-service-network", "", false, "Delete Service Network")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeDeleteServiceNetworkResourceAssociation, "delete-service-network-resource-association", "", false, "Delete Service Network Resource Association")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeDeleteServiceNetworkServiceAssociation, "delete-service-network-service-association", "", false, "Delete Service Network Service Association")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeDeleteServiceNetworkVpcAssociation, "delete-service-network-vpc-association", "", false, "Delete Service Network VPC Association")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeDeleteTargetGroup, "delete-target-group", "", false, "Delete Target Group")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeDeregisterTargets, "deregister-targets", "", false, "Deregister Targets")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeGetAccessLogSubscription, "get-access-log-subscription", "", false, "Get Access Log Subscription")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeGetAuthPolicy, "get-auth-policy", "", false, "Get Auth Policy")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeGetDomainVerification, "get-domain-verification", "", false, "Get Domain Verification")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeGetListener, "get-listener", "", false, "Get Listener")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeGetResourceConfiguration, "get-resource-configuration", "", false, "Get Resource Configuration")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeGetResourceGateway, "get-resource-gateway", "", false, "Get Resource Gateway")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeGetResourcePolicy, "get-resource-policy", "", false, "Get Resource Policy")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeGetRule, "get-rule", "", false, "Get Rule")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeGetService, "get-service", "", false, "Get Service")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeGetServiceNetwork, "get-service-network", "", false, "Get Service Network")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeGetServiceNetworkResourceAssociation, "get-service-network-resource-association", "", false, "Get Service Network Resource Association")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeGetServiceNetworkServiceAssociation, "get-service-network-service-association", "", false, "Get Service Network Service Association")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeGetServiceNetworkVpcAssociation, "get-service-network-vpc-association", "", false, "Get Service Network VPC Association")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeGetTargetGroup, "get-target-group", "", false, "Get Target Group")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeListAccessLogSubscriptions, "list-access-log-subscriptions", "", false, "List Access Log Subscriptions")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeListDomainVerifications, "list-domain-verifications", "", false, "List Domain Verifications")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeListListeners, "list-listeners", "", false, "List Listeners")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeListResourceConfigurations, "list-resource-configurations", "", false, "List Resource Configurations")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeListResourceEndpointAssociations, "list-resource-endpoint-associations", "", false, "List Resource Endpoint Associations")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeListResourceGateways, "list-resource-gateways", "", false, "List Resource Gateways")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeListRules, "list-rules", "", false, "List Rules")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeListServiceNetworkResourceAssociations, "list-service-network-resource-associations", "", false, "List Service Network Resource Associations")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeListServiceNetworkServiceAssociations, "list-service-network-service-associations", "", false, "List Service Network Service Associations")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeListServiceNetworkVpcAssociations, "list-service-network-vpc-associations", "", false, "List Service Network VPC Associations")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeListServiceNetworkVpcEndpointAssociations, "list-service-network-vpc-endpoint-associations", "", false, "List Service Network VPC Endpoint Associations")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeListServiceNetworks, "list-service-networks", "", false, "List Service Networks")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeListServices, "list-services", "", false, "List Services")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeListTargetGroups, "list-target-groups", "", false, "List Target Groups")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeListTargets, "list-targets", "", false, "List Targets")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticePutAuthPolicy, "put-auth-policy", "", false, "Put Auth Policy")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticePutResourcePolicy, "put-resource-policy", "", false, "Put Resource Policy")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeRegisterTargets, "register-targets", "", false, "Register Targets")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeStartDomainVerification, "start-domain-verification", "", false, "Start Domain Verification")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeTagResource, "tag-resource", "", false, "Tag Resource")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeUntagResource, "untag-resource", "", false, "Untag Resource")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeUpdateAccessLogSubscription, "update-access-log-subscription", "", false, "Update Access Log Subscription")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeUpdateListener, "update-listener", "", false, "Update Listener")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeUpdateResourceConfiguration, "update-resource-configuration", "", false, "Update Resource Configuration")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeUpdateResourceGateway, "update-resource-gateway", "", false, "Update Resource Gateway")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeUpdateRule, "update-rule", "", false, "Update Rule")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeUpdateService, "update-service", "", false, "Update Service")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeUpdateServiceNetwork, "update-service-network", "", false, "Update Service Network")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeUpdateServiceNetworkVpcAssociation, "update-service-network-vpc-association", "", false, "Update Service Network VPC Association")
	_vpclatticeCmd.Flags().BoolVarP(&_vpclatticeUpdateTargetGroup, "update-target-group", "", false, "Update Target Group")

}
