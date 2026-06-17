package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// servicediscoveryCmd represents the servicediscovery command
var _servicediscoveryCmd = &cobra.Command{
	Use:   "servicediscovery",
	Short: "AWS servicediscovery CLI",
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
		client := servicediscovery.NewFromConfig(cfg)
		if _servicediscoveryCreateHttpNamespace {
			servicediscovery_CreateHttpNamespace(cfg, client)
			return
		}
		if _servicediscoveryCreatePrivateDnsNamespace {
			servicediscovery_CreatePrivateDnsNamespace(cfg, client)
			return
		}
		if _servicediscoveryCreatePublicDnsNamespace {
			servicediscovery_CreatePublicDnsNamespace(cfg, client)
			return
		}
		if _servicediscoveryCreateService {
			servicediscovery_CreateService(cfg, client)
			return
		}
		if _servicediscoveryDeleteNamespace {
			servicediscovery_DeleteNamespace(cfg, client)
			return
		}
		if _servicediscoveryDeleteService {
			servicediscovery_DeleteService(cfg, client)
			return
		}
		if _servicediscoveryDeleteServiceAttributes {
			servicediscovery_DeleteServiceAttributes(cfg, client)
			return
		}
		if _servicediscoveryDeregisterInstance {
			servicediscovery_DeregisterInstance(cfg, client)
			return
		}
		if _servicediscoveryDiscoverInstances {
			servicediscovery_DiscoverInstances(cfg, client)
			return
		}
		if _servicediscoveryDiscoverInstancesRevision {
			servicediscovery_DiscoverInstancesRevision(cfg, client)
			return
		}
		if _servicediscoveryGetInstance {
			servicediscovery_GetInstance(cfg, client)
			return
		}
		if _servicediscoveryGetInstancesHealthStatus {
			servicediscovery_GetInstancesHealthStatus(cfg, client)
			return
		}
		if _servicediscoveryGetNamespace {
			servicediscovery_GetNamespace(cfg, client)
			return
		}
		if _servicediscoveryGetOperation {
			servicediscovery_GetOperation(cfg, client)
			return
		}
		if _servicediscoveryGetService {
			servicediscovery_GetService(cfg, client)
			return
		}
		if _servicediscoveryGetServiceAttributes {
			servicediscovery_GetServiceAttributes(cfg, client)
			return
		}
		if _servicediscoveryListInstances {
			servicediscovery_ListInstances(cfg, client)
			return
		}
		if _servicediscoveryListNamespaces {
			servicediscovery_ListNamespaces(cfg, client)
			return
		}
		if _servicediscoveryListOperations {
			servicediscovery_ListOperations(cfg, client)
			return
		}
		if _servicediscoveryListServices {
			servicediscovery_ListServices(cfg, client)
			return
		}
		if _servicediscoveryListTagsForResource {
			servicediscovery_ListTagsForResource(cfg, client)
			return
		}
		if _servicediscoveryRegisterInstance {
			servicediscovery_RegisterInstance(cfg, client)
			return
		}
		if _servicediscoveryTagResource {
			servicediscovery_TagResource(cfg, client)
			return
		}
		if _servicediscoveryUntagResource {
			servicediscovery_UntagResource(cfg, client)
			return
		}
		if _servicediscoveryUpdateHttpNamespace {
			servicediscovery_UpdateHttpNamespace(cfg, client)
			return
		}
		if _servicediscoveryUpdateInstanceCustomHealthStatus {
			servicediscovery_UpdateInstanceCustomHealthStatus(cfg, client)
			return
		}
		if _servicediscoveryUpdatePrivateDnsNamespace {
			servicediscovery_UpdatePrivateDnsNamespace(cfg, client)
			return
		}
		if _servicediscoveryUpdatePublicDnsNamespace {
			servicediscovery_UpdatePublicDnsNamespace(cfg, client)
			return
		}
		if _servicediscoveryUpdateService {
			servicediscovery_UpdateService(cfg, client)
			return
		}
		if _servicediscoveryUpdateServiceAttributes {
			servicediscovery_UpdateServiceAttributes(cfg, client)
			return
		}

	},
}

var (
	_servicediscoveryCreateHttpNamespace              bool
	_servicediscoveryCreatePrivateDnsNamespace        bool
	_servicediscoveryCreatePublicDnsNamespace         bool
	_servicediscoveryCreateService                    bool
	_servicediscoveryDeleteNamespace                  bool
	_servicediscoveryDeleteService                    bool
	_servicediscoveryDeleteServiceAttributes          bool
	_servicediscoveryDeregisterInstance               bool
	_servicediscoveryDiscoverInstances                bool
	_servicediscoveryDiscoverInstancesRevision        bool
	_servicediscoveryGetInstance                      bool
	_servicediscoveryGetInstancesHealthStatus         bool
	_servicediscoveryGetNamespace                     bool
	_servicediscoveryGetOperation                     bool
	_servicediscoveryGetService                       bool
	_servicediscoveryGetServiceAttributes             bool
	_servicediscoveryListInstances                    bool
	_servicediscoveryListNamespaces                   bool
	_servicediscoveryListOperations                   bool
	_servicediscoveryListServices                     bool
	_servicediscoveryListTagsForResource              bool
	_servicediscoveryRegisterInstance                 bool
	_servicediscoveryTagResource                      bool
	_servicediscoveryUntagResource                    bool
	_servicediscoveryUpdateHttpNamespace              bool
	_servicediscoveryUpdateInstanceCustomHealthStatus bool
	_servicediscoveryUpdatePrivateDnsNamespace        bool
	_servicediscoveryUpdatePublicDnsNamespace         bool
	_servicediscoveryUpdateService                    bool
	_servicediscoveryUpdateServiceAttributes          bool

	_servicediscoveryAttributes              string
	_servicediscoveryCreatorRequestId        string
	_servicediscoveryDescription             string
	_servicediscoveryDnsConfig               string
	_servicediscoveryFilters                 string
	_servicediscoveryHealthCheckConfig       string
	_servicediscoveryHealthCheckCustomConfig string
	_servicediscoveryHealthStatus            string
	_servicediscoveryId                      string
	_servicediscoveryInstanceId              string
	_servicediscoveryInstances               []string
	_servicediscoveryMaxResults              string
	_servicediscoveryName                    string
	_servicediscoveryNamespace               string
	_servicediscoveryNamespaceId             string
	_servicediscoveryNamespaceName           string
	_servicediscoveryNextToken               string
	_servicediscoveryOperationId             string
	_servicediscoveryOptionalParameters      string
	_servicediscoveryOwnerAccount            string
	_servicediscoveryProperties              string
	_servicediscoveryQueryParameters         string
	_servicediscoveryResourceARN             string
	_servicediscoveryService                 string
	_servicediscoveryServiceId               string
	_servicediscoveryServiceName             string
	_servicediscoveryStatus                  string
	_servicediscoveryTagKeys                 []string
	_servicediscoveryTags                    string
	_servicediscoveryType                    string
	_servicediscoveryUpdaterRequestId        string
	_servicediscoveryVpc                     string
)

// Creates an HTTP namespace. Service instances registered using an HTTP namespace
// can be discovered using a DiscoverInstances request but can't be discovered
// using DNS.
//
// For the current quota on the number of namespaces that you can create using the
// same Amazon Web Services account, see [Cloud Map quotas]in the Cloud Map Developer Guide.
//
// [Cloud Map quotas]: https://docs.aws.amazon.com/cloud-map/latest/dg/cloud-map-limits.html
func servicediscovery_CreateHttpNamespace(cfg aws.Config, client *servicediscovery.Client) {
	input := &servicediscovery.CreateHttpNamespaceInput{
		// Name: *string, // Required
	}

	if len(_servicediscoveryName) > 0 {
		input.Name = aws.String(_servicediscoveryName)
	}
	if len(_servicediscoveryCreatorRequestId) > 0 {
		input.CreatorRequestId = aws.String(_servicediscoveryCreatorRequestId)
	}
	if len(_servicediscoveryDescription) > 0 {
		input.Description = aws.String(_servicediscoveryDescription)
	}
	if len(_servicediscoveryTags) > 0 {
		if err := assignInputField(input, "Tags", _servicediscoveryTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateHttpNamespace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a private namespace based on DNS, which is visible only inside a
// specified Amazon VPC. The namespace defines your service naming scheme. For
// example, if you name your namespace example.com and name your service backend ,
// the resulting DNS name for the service is backend.example.com . Service
// instances that are registered using a private DNS namespace can be discovered
// using either a DiscoverInstances request or using DNS. For the current quota on
// the number of namespaces that you can create using the same Amazon Web Services
// account, see [Cloud Map quotas]in the Cloud Map Developer Guide.
//
// [Cloud Map quotas]: https://docs.aws.amazon.com/cloud-map/latest/dg/cloud-map-limits.html
func servicediscovery_CreatePrivateDnsNamespace(cfg aws.Config, client *servicediscovery.Client) {
	input := &servicediscovery.CreatePrivateDnsNamespaceInput{
		// Name: *string, // Required
		// Vpc: *string, // Required
	}

	if len(_servicediscoveryName) > 0 {
		input.Name = aws.String(_servicediscoveryName)
	}
	if len(_servicediscoveryVpc) > 0 {
		input.Vpc = aws.String(_servicediscoveryVpc)
	}
	if len(_servicediscoveryCreatorRequestId) > 0 {
		input.CreatorRequestId = aws.String(_servicediscoveryCreatorRequestId)
	}
	if len(_servicediscoveryDescription) > 0 {
		input.Description = aws.String(_servicediscoveryDescription)
	}
	if len(_servicediscoveryProperties) > 0 {
		if err := assignInputField(input, "Properties", _servicediscoveryProperties); err != nil {
			log.Errorf("invalid --properties: %s", err.Error())
			return
		}
	}
	if len(_servicediscoveryTags) > 0 {
		if err := assignInputField(input, "Tags", _servicediscoveryTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePrivateDnsNamespace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a public namespace based on DNS, which is visible on the internet. The
// namespace defines your service naming scheme. For example, if you name your
// namespace example.com and name your service backend , the resulting DNS name for
// the service is backend.example.com . You can discover instances that were
// registered with a public DNS namespace by using either a DiscoverInstances
// request or using DNS. For the current quota on the number of namespaces that you
// can create using the same Amazon Web Services account, see [Cloud Map quotas]in the Cloud Map
// Developer Guide.
//
// The CreatePublicDnsNamespace API operation is not supported in the Amazon Web
// Services GovCloud (US) Regions.
//
// [Cloud Map quotas]: https://docs.aws.amazon.com/cloud-map/latest/dg/cloud-map-limits.html
func servicediscovery_CreatePublicDnsNamespace(cfg aws.Config, client *servicediscovery.Client) {
	input := &servicediscovery.CreatePublicDnsNamespaceInput{
		// Name: *string, // Required
	}

	if len(_servicediscoveryName) > 0 {
		input.Name = aws.String(_servicediscoveryName)
	}
	if len(_servicediscoveryCreatorRequestId) > 0 {
		input.CreatorRequestId = aws.String(_servicediscoveryCreatorRequestId)
	}
	if len(_servicediscoveryDescription) > 0 {
		input.Description = aws.String(_servicediscoveryDescription)
	}
	if len(_servicediscoveryProperties) > 0 {
		if err := assignInputField(input, "Properties", _servicediscoveryProperties); err != nil {
			log.Errorf("invalid --properties: %s", err.Error())
			return
		}
	}
	if len(_servicediscoveryTags) > 0 {
		if err := assignInputField(input, "Tags", _servicediscoveryTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePublicDnsNamespace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a service. This action defines the configuration for the following
// entities:
//
// - For public and private DNS namespaces, one of the following combinations of
// DNS records in Amazon Route 53:
//
// - A
//
// - AAAA
//
// - A and AAAA
//
// - SRV
//
// - CNAME
//
// - Optionally, a health check
//
// After you create the service, you can submit a [RegisterInstance] request, and Cloud Map uses the
// values in the configuration to create the specified entities.
//
// For the current quota on the number of instances that you can register using
// the same namespace and using the same service, see [Cloud Map quotas]in the Cloud Map Developer
// Guide.
//
// [Cloud Map quotas]: https://docs.aws.amazon.com/cloud-map/latest/dg/cloud-map-limits.html
// [RegisterInstance]: https://docs.aws.amazon.com/cloud-map/latest/api/API_RegisterInstance.html
func servicediscovery_CreateService(cfg aws.Config, client *servicediscovery.Client) {
	input := &servicediscovery.CreateServiceInput{
		// Name: *string, // Required
	}

	if len(_servicediscoveryName) > 0 {
		input.Name = aws.String(_servicediscoveryName)
	}
	if len(_servicediscoveryCreatorRequestId) > 0 {
		input.CreatorRequestId = aws.String(_servicediscoveryCreatorRequestId)
	}
	if len(_servicediscoveryDescription) > 0 {
		input.Description = aws.String(_servicediscoveryDescription)
	}
	if len(_servicediscoveryDnsConfig) > 0 {
		if err := assignInputField(input, "DnsConfig", _servicediscoveryDnsConfig); err != nil {
			log.Errorf("invalid --dns-config: %s", err.Error())
			return
		}
	}
	if len(_servicediscoveryHealthCheckConfig) > 0 {
		if err := assignInputField(input, "HealthCheckConfig", _servicediscoveryHealthCheckConfig); err != nil {
			log.Errorf("invalid --health-check-config: %s", err.Error())
			return
		}
	}
	if len(_servicediscoveryHealthCheckCustomConfig) > 0 {
		if err := assignInputField(input, "HealthCheckCustomConfig", _servicediscoveryHealthCheckCustomConfig); err != nil {
			log.Errorf("invalid --health-check-custom-config: %s", err.Error())
			return
		}
	}
	if len(_servicediscoveryNamespaceId) > 0 {
		input.NamespaceId = aws.String(_servicediscoveryNamespaceId)
	}
	if len(_servicediscoveryTags) > 0 {
		if err := assignInputField(input, "Tags", _servicediscoveryTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_servicediscoveryType) > 0 {
		if err := assignInputField(input, "Type", _servicediscoveryType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
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

// Deletes a namespace from the current account. If the namespace still contains
// one or more services, the request fails.
func servicediscovery_DeleteNamespace(cfg aws.Config, client *servicediscovery.Client) {
	input := &servicediscovery.DeleteNamespaceInput{
		// Id: *string, // Required
	}

	if len(_servicediscoveryId) > 0 {
		input.Id = aws.String(_servicediscoveryId)
	}

	if resp, err := client.DeleteNamespace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specified service and all associated service attributes. If the
// service still contains one or more registered instances, the request fails.
func servicediscovery_DeleteService(cfg aws.Config, client *servicediscovery.Client) {
	input := &servicediscovery.DeleteServiceInput{
		// Id: *string, // Required
	}

	if len(_servicediscoveryId) > 0 {
		input.Id = aws.String(_servicediscoveryId)
	}

	if resp, err := client.DeleteService(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes specific attributes associated with a service.
func servicediscovery_DeleteServiceAttributes(cfg aws.Config, client *servicediscovery.Client) {
	input := &servicediscovery.DeleteServiceAttributesInput{
		// Attributes: []string, // Required
		// ServiceId: *string, // Required
	}

	if len(_servicediscoveryAttributes) > 0 {
		input.Attributes = []string{_servicediscoveryAttributes}
	}
	if len(_servicediscoveryServiceId) > 0 {
		input.ServiceId = aws.String(_servicediscoveryServiceId)
	}

	if resp, err := client.DeleteServiceAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the Amazon Route 53 DNS records and health check, if any, that Cloud
// Map created for the specified instance.
func servicediscovery_DeregisterInstance(cfg aws.Config, client *servicediscovery.Client) {
	input := &servicediscovery.DeregisterInstanceInput{
		// InstanceId: *string, // Required
		// ServiceId: *string, // Required
	}

	if len(_servicediscoveryInstanceId) > 0 {
		input.InstanceId = aws.String(_servicediscoveryInstanceId)
	}
	if len(_servicediscoveryServiceId) > 0 {
		input.ServiceId = aws.String(_servicediscoveryServiceId)
	}

	if resp, err := client.DeregisterInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Discovers registered instances for a specified namespace and service. You can
// use DiscoverInstances to discover instances for any type of namespace.
// DiscoverInstances returns a randomized list of instances allowing customers to
// distribute traffic evenly across instances. For public and private DNS
// namespaces, you can also use DNS queries to discover instances.
func servicediscovery_DiscoverInstances(cfg aws.Config, client *servicediscovery.Client) {
	input := &servicediscovery.DiscoverInstancesInput{
		// NamespaceName: *string, // Required
		// ServiceName: *string, // Required
	}

	if len(_servicediscoveryNamespaceName) > 0 {
		input.NamespaceName = aws.String(_servicediscoveryNamespaceName)
	}
	if len(_servicediscoveryServiceName) > 0 {
		input.ServiceName = aws.String(_servicediscoveryServiceName)
	}
	if len(_servicediscoveryHealthStatus) > 0 {
		if err := assignInputField(input, "HealthStatus", _servicediscoveryHealthStatus); err != nil {
			log.Errorf("invalid --health-status: %s", err.Error())
			return
		}
	}
	if len(_servicediscoveryMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _servicediscoveryMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_servicediscoveryOptionalParameters) > 0 {
		if err := assignInputField(input, "OptionalParameters", _servicediscoveryOptionalParameters); err != nil {
			log.Errorf("invalid --optional-parameters: %s", err.Error())
			return
		}
	}
	if len(_servicediscoveryOwnerAccount) > 0 {
		input.OwnerAccount = aws.String(_servicediscoveryOwnerAccount)
	}
	if len(_servicediscoveryQueryParameters) > 0 {
		if err := assignInputField(input, "QueryParameters", _servicediscoveryQueryParameters); err != nil {
			log.Errorf("invalid --query-parameters: %s", err.Error())
			return
		}
	}

	if resp, err := client.DiscoverInstances(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Discovers the increasing revision associated with an instance.
func servicediscovery_DiscoverInstancesRevision(cfg aws.Config, client *servicediscovery.Client) {
	input := &servicediscovery.DiscoverInstancesRevisionInput{
		// NamespaceName: *string, // Required
		// ServiceName: *string, // Required
	}

	if len(_servicediscoveryNamespaceName) > 0 {
		input.NamespaceName = aws.String(_servicediscoveryNamespaceName)
	}
	if len(_servicediscoveryServiceName) > 0 {
		input.ServiceName = aws.String(_servicediscoveryServiceName)
	}
	if len(_servicediscoveryOwnerAccount) > 0 {
		input.OwnerAccount = aws.String(_servicediscoveryOwnerAccount)
	}

	if resp, err := client.DiscoverInstancesRevision(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a specified instance.
func servicediscovery_GetInstance(cfg aws.Config, client *servicediscovery.Client) {
	input := &servicediscovery.GetInstanceInput{
		// InstanceId: *string, // Required
		// ServiceId: *string, // Required
	}

	if len(_servicediscoveryInstanceId) > 0 {
		input.InstanceId = aws.String(_servicediscoveryInstanceId)
	}
	if len(_servicediscoveryServiceId) > 0 {
		input.ServiceId = aws.String(_servicediscoveryServiceId)
	}

	if resp, err := client.GetInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the current health status ( Healthy , Unhealthy , or Unknown ) of one or
// more instances that are associated with a specified service.
//
// There's a brief delay between when you register an instance and when the health
// status for the instance is available.
func servicediscovery_GetInstancesHealthStatus(cfg aws.Config, client *servicediscovery.Client) {
	input := &servicediscovery.GetInstancesHealthStatusInput{
		// ServiceId: *string, // Required
	}

	if len(_servicediscoveryServiceId) > 0 {
		input.ServiceId = aws.String(_servicediscoveryServiceId)
	}
	if len(_servicediscoveryInstances) > 0 {
		input.Instances = append([]string(nil), _servicediscoveryInstances...)
	}
	if len(_servicediscoveryMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _servicediscoveryMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_servicediscoveryNextToken) > 0 {
		input.NextToken = aws.String(_servicediscoveryNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetInstancesHealthStatus(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*servicediscovery.GetInstancesHealthStatusOutput
	p := servicediscovery.NewGetInstancesHealthStatusPaginator(client, input)
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

// Gets information about a namespace.
func servicediscovery_GetNamespace(cfg aws.Config, client *servicediscovery.Client) {
	input := &servicediscovery.GetNamespaceInput{
		// Id: *string, // Required
	}

	if len(_servicediscoveryId) > 0 {
		input.Id = aws.String(_servicediscoveryId)
	}

	if resp, err := client.GetNamespace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about any operation that returns an operation ID in the
// response, such as a CreateHttpNamespace request.
//
// To get a list of operations that match specified criteria, see [ListOperations].
//
// [ListOperations]: https://docs.aws.amazon.com/cloud-map/latest/api/API_ListOperations.html
func servicediscovery_GetOperation(cfg aws.Config, client *servicediscovery.Client) {
	input := &servicediscovery.GetOperationInput{
		// OperationId: *string, // Required
	}

	if len(_servicediscoveryOperationId) > 0 {
		input.OperationId = aws.String(_servicediscoveryOperationId)
	}
	if len(_servicediscoveryOwnerAccount) > 0 {
		input.OwnerAccount = aws.String(_servicediscoveryOwnerAccount)
	}

	if resp, err := client.GetOperation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the settings for a specified service.
func servicediscovery_GetService(cfg aws.Config, client *servicediscovery.Client) {
	input := &servicediscovery.GetServiceInput{
		// Id: *string, // Required
	}

	if len(_servicediscoveryId) > 0 {
		input.Id = aws.String(_servicediscoveryId)
	}

	if resp, err := client.GetService(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the attributes associated with a specified service.
func servicediscovery_GetServiceAttributes(cfg aws.Config, client *servicediscovery.Client) {
	input := &servicediscovery.GetServiceAttributesInput{
		// ServiceId: *string, // Required
	}

	if len(_servicediscoveryServiceId) > 0 {
		input.ServiceId = aws.String(_servicediscoveryServiceId)
	}

	if resp, err := client.GetServiceAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists summary information about the instances that you registered by using a
// specified service.
func servicediscovery_ListInstances(cfg aws.Config, client *servicediscovery.Client) {
	input := &servicediscovery.ListInstancesInput{
		// ServiceId: *string, // Required
	}

	if len(_servicediscoveryServiceId) > 0 {
		input.ServiceId = aws.String(_servicediscoveryServiceId)
	}
	if len(_servicediscoveryMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _servicediscoveryMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_servicediscoveryNextToken) > 0 {
		input.NextToken = aws.String(_servicediscoveryNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListInstances(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*servicediscovery.ListInstancesOutput
	p := servicediscovery.NewListInstancesPaginator(client, input)
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

// Lists summary information about the namespaces that were created by the current
// Amazon Web Services account and shared with the current Amazon Web Services
// account.
func servicediscovery_ListNamespaces(cfg aws.Config, client *servicediscovery.Client) {
	input := &servicediscovery.ListNamespacesInput{}

	if len(_servicediscoveryFilters) > 0 {
		if err := assignInputField(input, "Filters", _servicediscoveryFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_servicediscoveryMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _servicediscoveryMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_servicediscoveryNextToken) > 0 {
		input.NextToken = aws.String(_servicediscoveryNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListNamespaces(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*servicediscovery.ListNamespacesOutput
	p := servicediscovery.NewListNamespacesPaginator(client, input)
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

// Lists operations that match the criteria that you specify.
func servicediscovery_ListOperations(cfg aws.Config, client *servicediscovery.Client) {
	input := &servicediscovery.ListOperationsInput{}

	if len(_servicediscoveryFilters) > 0 {
		if err := assignInputField(input, "Filters", _servicediscoveryFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_servicediscoveryMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _servicediscoveryMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_servicediscoveryNextToken) > 0 {
		input.NextToken = aws.String(_servicediscoveryNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListOperations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*servicediscovery.ListOperationsOutput
	p := servicediscovery.NewListOperationsPaginator(client, input)
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

// Lists summary information for all the services that are associated with one or
// more namespaces.
func servicediscovery_ListServices(cfg aws.Config, client *servicediscovery.Client) {
	input := &servicediscovery.ListServicesInput{}

	if len(_servicediscoveryFilters) > 0 {
		if err := assignInputField(input, "Filters", _servicediscoveryFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_servicediscoveryMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _servicediscoveryMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_servicediscoveryNextToken) > 0 {
		input.NextToken = aws.String(_servicediscoveryNextToken)
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

	var results []*servicediscovery.ListServicesOutput
	p := servicediscovery.NewListServicesPaginator(client, input)
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

// Lists tags for the specified resource.
func servicediscovery_ListTagsForResource(cfg aws.Config, client *servicediscovery.Client) {
	input := &servicediscovery.ListTagsForResourceInput{
		// ResourceARN: *string, // Required
	}

	if len(_servicediscoveryResourceARN) > 0 {
		input.ResourceARN = aws.String(_servicediscoveryResourceARN)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates one or more records and, optionally, creates a health check
// based on the settings in a specified service. When you submit a RegisterInstance
// request, the following occurs:
//
// - For each DNS record that you define in the service that's specified by
// ServiceId , a record is created or updated in the hosted zone that's
// associated with the corresponding namespace.
//
// - If the service includes HealthCheckConfig , a health check is created based
// on the settings in the health check configuration.
//
// - The health check, if any, is associated with each of the new or updated
// records.
//
// One RegisterInstance request must complete before you can submit another
// request and specify the same service ID and instance ID.
//
// For more information, see [CreateService].
//
// When Cloud Map receives a DNS query for the specified DNS name, it returns the
// applicable value:
//
// - If the health check is healthy: returns all the records
//
// - If the health check is unhealthy: returns the applicable value for the last
// healthy instance
//
// - If you didn't specify a health check configuration: returns all the records
//
// For the current quota on the number of instances that you can register using
// the same namespace and using the same service, see [Cloud Map quotas]in the Cloud Map Developer
// Guide.
//
// [CreateService]: https://docs.aws.amazon.com/cloud-map/latest/api/API_CreateService.html
// [Cloud Map quotas]: https://docs.aws.amazon.com/cloud-map/latest/dg/cloud-map-limits.html
func servicediscovery_RegisterInstance(cfg aws.Config, client *servicediscovery.Client) {
	input := &servicediscovery.RegisterInstanceInput{
		// Attributes: map[string]string, // Required
		// InstanceId: *string, // Required
		// ServiceId: *string, // Required
	}

	if len(_servicediscoveryAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _servicediscoveryAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_servicediscoveryInstanceId) > 0 {
		input.InstanceId = aws.String(_servicediscoveryInstanceId)
	}
	if len(_servicediscoveryServiceId) > 0 {
		input.ServiceId = aws.String(_servicediscoveryServiceId)
	}
	if len(_servicediscoveryCreatorRequestId) > 0 {
		input.CreatorRequestId = aws.String(_servicediscoveryCreatorRequestId)
	}

	if resp, err := client.RegisterInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds one or more tags to the specified resource.
func servicediscovery_TagResource(cfg aws.Config, client *servicediscovery.Client) {
	input := &servicediscovery.TagResourceInput{
		// ResourceARN: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_servicediscoveryResourceARN) > 0 {
		input.ResourceARN = aws.String(_servicediscoveryResourceARN)
	}
	if len(_servicediscoveryTags) > 0 {
		if err := assignInputField(input, "Tags", _servicediscoveryTags); err != nil {
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

// Removes one or more tags from the specified resource.
func servicediscovery_UntagResource(cfg aws.Config, client *servicediscovery.Client) {
	input := &servicediscovery.UntagResourceInput{
		// ResourceARN: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_servicediscoveryResourceARN) > 0 {
		input.ResourceARN = aws.String(_servicediscoveryResourceARN)
	}
	if len(_servicediscoveryTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _servicediscoveryTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an HTTP namespace.
func servicediscovery_UpdateHttpNamespace(cfg aws.Config, client *servicediscovery.Client) {
	input := &servicediscovery.UpdateHttpNamespaceInput{
		// Id: *string, // Required
		// Namespace: *types.HttpNamespaceChange, // Required
	}

	if len(_servicediscoveryId) > 0 {
		input.Id = aws.String(_servicediscoveryId)
	}
	if len(_servicediscoveryNamespace) > 0 {
		if err := assignInputField(input, "Namespace", _servicediscoveryNamespace); err != nil {
			log.Errorf("invalid --namespace: %s", err.Error())
			return
		}
	}
	if len(_servicediscoveryUpdaterRequestId) > 0 {
		input.UpdaterRequestId = aws.String(_servicediscoveryUpdaterRequestId)
	}

	if resp, err := client.UpdateHttpNamespace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Submits a request to change the health status of a custom health check to
// healthy or unhealthy.
//
// You can use UpdateInstanceCustomHealthStatus to change the status only for
// custom health checks, which you define using HealthCheckCustomConfig when you
// create a service. You can't use it to change the status for Route 53 health
// checks, which you define using HealthCheckConfig .
//
// For more information, see [HealthCheckCustomConfig].
//
// [HealthCheckCustomConfig]: https://docs.aws.amazon.com/cloud-map/latest/api/API_HealthCheckCustomConfig.html
func servicediscovery_UpdateInstanceCustomHealthStatus(cfg aws.Config, client *servicediscovery.Client) {
	input := &servicediscovery.UpdateInstanceCustomHealthStatusInput{
		// InstanceId: *string, // Required
		// ServiceId: *string, // Required
		// Status: types.CustomHealthStatus, // Required
	}

	if len(_servicediscoveryInstanceId) > 0 {
		input.InstanceId = aws.String(_servicediscoveryInstanceId)
	}
	if len(_servicediscoveryServiceId) > 0 {
		input.ServiceId = aws.String(_servicediscoveryServiceId)
	}
	if len(_servicediscoveryStatus) > 0 {
		if err := assignInputField(input, "Status", _servicediscoveryStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateInstanceCustomHealthStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a private DNS namespace.
func servicediscovery_UpdatePrivateDnsNamespace(cfg aws.Config, client *servicediscovery.Client) {
	input := &servicediscovery.UpdatePrivateDnsNamespaceInput{
		// Id: *string, // Required
		// Namespace: *types.PrivateDnsNamespaceChange, // Required
	}

	if len(_servicediscoveryId) > 0 {
		input.Id = aws.String(_servicediscoveryId)
	}
	if len(_servicediscoveryNamespace) > 0 {
		if err := assignInputField(input, "Namespace", _servicediscoveryNamespace); err != nil {
			log.Errorf("invalid --namespace: %s", err.Error())
			return
		}
	}
	if len(_servicediscoveryUpdaterRequestId) > 0 {
		input.UpdaterRequestId = aws.String(_servicediscoveryUpdaterRequestId)
	}

	if resp, err := client.UpdatePrivateDnsNamespace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a public DNS namespace.
func servicediscovery_UpdatePublicDnsNamespace(cfg aws.Config, client *servicediscovery.Client) {
	input := &servicediscovery.UpdatePublicDnsNamespaceInput{
		// Id: *string, // Required
		// Namespace: *types.PublicDnsNamespaceChange, // Required
	}

	if len(_servicediscoveryId) > 0 {
		input.Id = aws.String(_servicediscoveryId)
	}
	if len(_servicediscoveryNamespace) > 0 {
		if err := assignInputField(input, "Namespace", _servicediscoveryNamespace); err != nil {
			log.Errorf("invalid --namespace: %s", err.Error())
			return
		}
	}
	if len(_servicediscoveryUpdaterRequestId) > 0 {
		input.UpdaterRequestId = aws.String(_servicediscoveryUpdaterRequestId)
	}

	if resp, err := client.UpdatePublicDnsNamespace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Submits a request to perform the following operations:
// - Update the TTL setting for existing DnsRecords configurations
//
// - Add, update, or delete HealthCheckConfig for a specified service
//
// You can't add, update, or delete a HealthCheckCustomConfig configuration.
//
// For public and private DNS namespaces, note the following:
//
// - If you omit any existing DnsRecords or HealthCheckConfig configurations from
// an UpdateService request, the configurations are deleted from the service.
//
// - If you omit an existing HealthCheckCustomConfig configuration from an
// UpdateService request, the configuration isn't deleted from the service.
//
// You can't call UpdateService and update settings in the following scenarios:
//
// - When the service is associated with an HTTP namespace
//
// - When the service is associated with a shared namespace and contains
// instances that were registered by Amazon Web Services accounts other than the
// account making the UpdateService call
//
// When you update settings for a service, Cloud Map also updates the
// corresponding settings in all the records and health checks that were created by
// using the specified service.
func servicediscovery_UpdateService(cfg aws.Config, client *servicediscovery.Client) {
	input := &servicediscovery.UpdateServiceInput{
		// Id: *string, // Required
		// Service: *types.ServiceChange, // Required
	}

	if len(_servicediscoveryId) > 0 {
		input.Id = aws.String(_servicediscoveryId)
	}
	if len(_servicediscoveryService) > 0 {
		if err := assignInputField(input, "Service", _servicediscoveryService); err != nil {
			log.Errorf("invalid --service: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateService(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Submits a request to update a specified service to add service-level attributes.
func servicediscovery_UpdateServiceAttributes(cfg aws.Config, client *servicediscovery.Client) {
	input := &servicediscovery.UpdateServiceAttributesInput{
		// Attributes: map[string]string, // Required
		// ServiceId: *string, // Required
	}

	if len(_servicediscoveryAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _servicediscoveryAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_servicediscoveryServiceId) > 0 {
		input.ServiceId = aws.String(_servicediscoveryServiceId)
	}

	if resp, err := client.UpdateServiceAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_servicediscoveryCmd)
	_servicediscoveryCmd.Flags().SortFlags = false

	_servicediscoveryCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_servicediscoveryCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_servicediscoveryCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_servicediscoveryCmd.Flags().StringVarP(&_servicediscoveryAttributes, "attributes", "", "", "Attributes")
	_servicediscoveryCmd.Flags().StringVarP(&_servicediscoveryCreatorRequestId, "creator-request-id", "", "", "Creator Request ID")
	_servicediscoveryCmd.Flags().StringVarP(&_servicediscoveryDescription, "description", "", "", "Description")
	_servicediscoveryCmd.Flags().StringVarP(&_servicediscoveryDnsConfig, "dns-config", "", "", "DNS Config")
	_servicediscoveryCmd.Flags().StringVarP(&_servicediscoveryFilters, "filters", "", "", "Filters")
	_servicediscoveryCmd.Flags().StringVarP(&_servicediscoveryHealthCheckConfig, "health-check-config", "", "", "Health Check Config")
	_servicediscoveryCmd.Flags().StringVarP(&_servicediscoveryHealthCheckCustomConfig, "health-check-custom-config", "", "", "Health Check Custom Config")
	_servicediscoveryCmd.Flags().StringVarP(&_servicediscoveryHealthStatus, "health-status", "", "", "Health Status")
	_servicediscoveryCmd.Flags().StringVarP(&_servicediscoveryId, "id", "", "", "ID")
	_servicediscoveryCmd.Flags().StringVarP(&_servicediscoveryInstanceId, "instance-id", "", "", "Instance ID")
	_servicediscoveryCmd.Flags().StringSliceVarP(&_servicediscoveryInstances, "instances", "", nil, "Instances")
	_servicediscoveryCmd.Flags().StringVarP(&_servicediscoveryMaxResults, "max-results", "", "", "Max Results")
	_servicediscoveryCmd.Flags().StringVarP(&_servicediscoveryName, "name", "", "", "Name")
	_servicediscoveryCmd.Flags().StringVarP(&_servicediscoveryNamespace, "namespace", "", "", "Namespace")
	_servicediscoveryCmd.Flags().StringVarP(&_servicediscoveryNamespaceId, "namespace-id", "", "", "Namespace ID")
	_servicediscoveryCmd.Flags().StringVarP(&_servicediscoveryNamespaceName, "namespace-name", "", "", "Namespace Name")
	_servicediscoveryCmd.Flags().StringVarP(&_servicediscoveryNextToken, "next-token", "", "", "Next Token")
	_servicediscoveryCmd.Flags().StringVarP(&_servicediscoveryOperationId, "operation-id", "", "", "Operation ID")
	_servicediscoveryCmd.Flags().StringVarP(&_servicediscoveryOptionalParameters, "optional-parameters", "", "", "Optional Parameters")
	_servicediscoveryCmd.Flags().StringVarP(&_servicediscoveryOwnerAccount, "owner-account", "", "", "Owner Account")
	_servicediscoveryCmd.Flags().StringVarP(&_servicediscoveryProperties, "properties", "", "", "Properties")
	_servicediscoveryCmd.Flags().StringVarP(&_servicediscoveryQueryParameters, "query-parameters", "", "", "Query Parameters")
	_servicediscoveryCmd.Flags().StringVarP(&_servicediscoveryResourceARN, "resource-arn", "", "", "Resource ARN")
	_servicediscoveryCmd.Flags().StringVarP(&_servicediscoveryService, "service", "", "", "Service")
	_servicediscoveryCmd.Flags().StringVarP(&_servicediscoveryServiceId, "service-id", "", "", "Service ID")
	_servicediscoveryCmd.Flags().StringVarP(&_servicediscoveryServiceName, "service-name", "", "", "Service Name")
	_servicediscoveryCmd.Flags().StringVarP(&_servicediscoveryStatus, "status", "", "", "Status")
	_servicediscoveryCmd.Flags().StringSliceVarP(&_servicediscoveryTagKeys, "tag-keys", "", nil, "Tag Keys")
	_servicediscoveryCmd.Flags().StringVarP(&_servicediscoveryTags, "tags", "", "", "Tags")
	_servicediscoveryCmd.Flags().StringVarP(&_servicediscoveryType, "type", "", "", "Type")
	_servicediscoveryCmd.Flags().StringVarP(&_servicediscoveryUpdaterRequestId, "updater-request-id", "", "", "Updater Request ID")
	_servicediscoveryCmd.Flags().StringVarP(&_servicediscoveryVpc, "vpc", "", "", "VPC")

	_servicediscoveryCmd.Flags().BoolVarP(&_servicediscoveryCreateHttpNamespace, "create-http-namespace", "", false, "Create HTTP Namespace")
	_servicediscoveryCmd.Flags().BoolVarP(&_servicediscoveryCreatePrivateDnsNamespace, "create-private-dns-namespace", "", false, "Create Private DNS Namespace")
	_servicediscoveryCmd.Flags().BoolVarP(&_servicediscoveryCreatePublicDnsNamespace, "create-public-dns-namespace", "", false, "Create Public DNS Namespace")
	_servicediscoveryCmd.Flags().BoolVarP(&_servicediscoveryCreateService, "create-service", "", false, "Create Service")
	_servicediscoveryCmd.Flags().BoolVarP(&_servicediscoveryDeleteNamespace, "delete-namespace", "", false, "Delete Namespace")
	_servicediscoveryCmd.Flags().BoolVarP(&_servicediscoveryDeleteService, "delete-service", "", false, "Delete Service")
	_servicediscoveryCmd.Flags().BoolVarP(&_servicediscoveryDeleteServiceAttributes, "delete-service-attributes", "", false, "Delete Service Attributes")
	_servicediscoveryCmd.Flags().BoolVarP(&_servicediscoveryDeregisterInstance, "deregister-instance", "", false, "Deregister Instance")
	_servicediscoveryCmd.Flags().BoolVarP(&_servicediscoveryDiscoverInstances, "discover-instances", "", false, "Discover Instances")
	_servicediscoveryCmd.Flags().BoolVarP(&_servicediscoveryDiscoverInstancesRevision, "discover-instances-revision", "", false, "Discover Instances Revision")
	_servicediscoveryCmd.Flags().BoolVarP(&_servicediscoveryGetInstance, "get-instance", "", false, "Get Instance")
	_servicediscoveryCmd.Flags().BoolVarP(&_servicediscoveryGetInstancesHealthStatus, "get-instances-health-status", "", false, "Get Instances Health Status")
	_servicediscoveryCmd.Flags().BoolVarP(&_servicediscoveryGetNamespace, "get-namespace", "", false, "Get Namespace")
	_servicediscoveryCmd.Flags().BoolVarP(&_servicediscoveryGetOperation, "get-operation", "", false, "Get Operation")
	_servicediscoveryCmd.Flags().BoolVarP(&_servicediscoveryGetService, "get-service", "", false, "Get Service")
	_servicediscoveryCmd.Flags().BoolVarP(&_servicediscoveryGetServiceAttributes, "get-service-attributes", "", false, "Get Service Attributes")
	_servicediscoveryCmd.Flags().BoolVarP(&_servicediscoveryListInstances, "list-instances", "", false, "List Instances")
	_servicediscoveryCmd.Flags().BoolVarP(&_servicediscoveryListNamespaces, "list-namespaces", "", false, "List Namespaces")
	_servicediscoveryCmd.Flags().BoolVarP(&_servicediscoveryListOperations, "list-operations", "", false, "List Operations")
	_servicediscoveryCmd.Flags().BoolVarP(&_servicediscoveryListServices, "list-services", "", false, "List Services")
	_servicediscoveryCmd.Flags().BoolVarP(&_servicediscoveryListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_servicediscoveryCmd.Flags().BoolVarP(&_servicediscoveryRegisterInstance, "register-instance", "", false, "Register Instance")
	_servicediscoveryCmd.Flags().BoolVarP(&_servicediscoveryTagResource, "tag-resource", "", false, "Tag Resource")
	_servicediscoveryCmd.Flags().BoolVarP(&_servicediscoveryUntagResource, "untag-resource", "", false, "Untag Resource")
	_servicediscoveryCmd.Flags().BoolVarP(&_servicediscoveryUpdateHttpNamespace, "update-http-namespace", "", false, "Update HTTP Namespace")
	_servicediscoveryCmd.Flags().BoolVarP(&_servicediscoveryUpdateInstanceCustomHealthStatus, "update-instance-custom-health-status", "", false, "Update Instance Custom Health Status")
	_servicediscoveryCmd.Flags().BoolVarP(&_servicediscoveryUpdatePrivateDnsNamespace, "update-private-dns-namespace", "", false, "Update Private DNS Namespace")
	_servicediscoveryCmd.Flags().BoolVarP(&_servicediscoveryUpdatePublicDnsNamespace, "update-public-dns-namespace", "", false, "Update Public DNS Namespace")
	_servicediscoveryCmd.Flags().BoolVarP(&_servicediscoveryUpdateService, "update-service", "", false, "Update Service")
	_servicediscoveryCmd.Flags().BoolVarP(&_servicediscoveryUpdateServiceAttributes, "update-service-attributes", "", false, "Update Service Attributes")

}
