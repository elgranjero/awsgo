package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apprunner"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// apprunnerCmd represents the apprunner command
var _apprunnerCmd = &cobra.Command{
	Use:   "apprunner",
	Short: "AWS apprunner CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := apprunner.NewFromConfig(cfg)
		if _apprunnerAssociateCustomDomain {
			apprunner_AssociateCustomDomain(cfg, client)
			return
		}
		if _apprunnerCreateAutoScalingConfiguration {
			apprunner_CreateAutoScalingConfiguration(cfg, client)
			return
		}
		if _apprunnerCreateConnection {
			apprunner_CreateConnection(cfg, client)
			return
		}
		if _apprunnerCreateObservabilityConfiguration {
			apprunner_CreateObservabilityConfiguration(cfg, client)
			return
		}
		if _apprunnerCreateService {
			apprunner_CreateService(cfg, client)
			return
		}
		if _apprunnerCreateVpcConnector {
			apprunner_CreateVpcConnector(cfg, client)
			return
		}
		if _apprunnerCreateVpcIngressConnection {
			apprunner_CreateVpcIngressConnection(cfg, client)
			return
		}
		if _apprunnerDeleteAutoScalingConfiguration {
			apprunner_DeleteAutoScalingConfiguration(cfg, client)
			return
		}
		if _apprunnerDeleteConnection {
			apprunner_DeleteConnection(cfg, client)
			return
		}
		if _apprunnerDeleteObservabilityConfiguration {
			apprunner_DeleteObservabilityConfiguration(cfg, client)
			return
		}
		if _apprunnerDeleteService {
			apprunner_DeleteService(cfg, client)
			return
		}
		if _apprunnerDeleteVpcConnector {
			apprunner_DeleteVpcConnector(cfg, client)
			return
		}
		if _apprunnerDeleteVpcIngressConnection {
			apprunner_DeleteVpcIngressConnection(cfg, client)
			return
		}
		if _apprunnerDescribeAutoScalingConfiguration {
			apprunner_DescribeAutoScalingConfiguration(cfg, client)
			return
		}
		if _apprunnerDescribeCustomDomains {
			apprunner_DescribeCustomDomains(cfg, client)
			return
		}
		if _apprunnerDescribeObservabilityConfiguration {
			apprunner_DescribeObservabilityConfiguration(cfg, client)
			return
		}
		if _apprunnerDescribeService {
			apprunner_DescribeService(cfg, client)
			return
		}
		if _apprunnerDescribeVpcConnector {
			apprunner_DescribeVpcConnector(cfg, client)
			return
		}
		if _apprunnerDescribeVpcIngressConnection {
			apprunner_DescribeVpcIngressConnection(cfg, client)
			return
		}
		if _apprunnerDisassociateCustomDomain {
			apprunner_DisassociateCustomDomain(cfg, client)
			return
		}
		if _apprunnerListAutoScalingConfigurations {
			apprunner_ListAutoScalingConfigurations(cfg, client)
			return
		}
		if _apprunnerListConnections {
			apprunner_ListConnections(cfg, client)
			return
		}
		if _apprunnerListObservabilityConfigurations {
			apprunner_ListObservabilityConfigurations(cfg, client)
			return
		}
		if _apprunnerListOperations {
			apprunner_ListOperations(cfg, client)
			return
		}
		if _apprunnerListServices {
			apprunner_ListServices(cfg, client)
			return
		}
		if _apprunnerListServicesForAutoScalingConfiguration {
			apprunner_ListServicesForAutoScalingConfiguration(cfg, client)
			return
		}
		if _apprunnerListTagsForResource {
			apprunner_ListTagsForResource(cfg, client)
			return
		}
		if _apprunnerListVpcConnectors {
			apprunner_ListVpcConnectors(cfg, client)
			return
		}
		if _apprunnerListVpcIngressConnections {
			apprunner_ListVpcIngressConnections(cfg, client)
			return
		}
		if _apprunnerPauseService {
			apprunner_PauseService(cfg, client)
			return
		}
		if _apprunnerResumeService {
			apprunner_ResumeService(cfg, client)
			return
		}
		if _apprunnerStartDeployment {
			apprunner_StartDeployment(cfg, client)
			return
		}
		if _apprunnerTagResource {
			apprunner_TagResource(cfg, client)
			return
		}
		if _apprunnerUntagResource {
			apprunner_UntagResource(cfg, client)
			return
		}
		if _apprunnerUpdateDefaultAutoScalingConfiguration {
			apprunner_UpdateDefaultAutoScalingConfiguration(cfg, client)
			return
		}
		if _apprunnerUpdateService {
			apprunner_UpdateService(cfg, client)
			return
		}
		if _apprunnerUpdateVpcIngressConnection {
			apprunner_UpdateVpcIngressConnection(cfg, client)
			return
		}

	},
}

var (
	_apprunnerAssociateCustomDomain                   bool
	_apprunnerCreateAutoScalingConfiguration          bool
	_apprunnerCreateConnection                        bool
	_apprunnerCreateObservabilityConfiguration        bool
	_apprunnerCreateService                           bool
	_apprunnerCreateVpcConnector                      bool
	_apprunnerCreateVpcIngressConnection              bool
	_apprunnerDeleteAutoScalingConfiguration          bool
	_apprunnerDeleteConnection                        bool
	_apprunnerDeleteObservabilityConfiguration        bool
	_apprunnerDeleteService                           bool
	_apprunnerDeleteVpcConnector                      bool
	_apprunnerDeleteVpcIngressConnection              bool
	_apprunnerDescribeAutoScalingConfiguration        bool
	_apprunnerDescribeCustomDomains                   bool
	_apprunnerDescribeObservabilityConfiguration      bool
	_apprunnerDescribeService                         bool
	_apprunnerDescribeVpcConnector                    bool
	_apprunnerDescribeVpcIngressConnection            bool
	_apprunnerDisassociateCustomDomain                bool
	_apprunnerListAutoScalingConfigurations           bool
	_apprunnerListConnections                         bool
	_apprunnerListObservabilityConfigurations         bool
	_apprunnerListOperations                          bool
	_apprunnerListServices                            bool
	_apprunnerListServicesForAutoScalingConfiguration bool
	_apprunnerListTagsForResource                     bool
	_apprunnerListVpcConnectors                       bool
	_apprunnerListVpcIngressConnections               bool
	_apprunnerPauseService                            bool
	_apprunnerResumeService                           bool
	_apprunnerStartDeployment                         bool
	_apprunnerTagResource                             bool
	_apprunnerUntagResource                           bool
	_apprunnerUpdateDefaultAutoScalingConfiguration   bool
	_apprunnerUpdateService                           bool
	_apprunnerUpdateVpcIngressConnection              bool

	_apprunnerAutoScalingConfigurationArn    string
	_apprunnerAutoScalingConfigurationName   string
	_apprunnerConnectionArn                  string
	_apprunnerConnectionName                 string
	_apprunnerDeleteAllRevisions             string
	_apprunnerDomainName                     string
	_apprunnerEnableWWWSubdomain             string
	_apprunnerEncryptionConfiguration        string
	_apprunnerFilter                         string
	_apprunnerHealthCheckConfiguration       string
	_apprunnerIngressVpcConfiguration        string
	_apprunnerInstanceConfiguration          string
	_apprunnerLatestOnly                     string
	_apprunnerMaxConcurrency                 string
	_apprunnerMaxResults                     string
	_apprunnerMaxSize                        string
	_apprunnerMinSize                        string
	_apprunnerNetworkConfiguration           string
	_apprunnerNextToken                      string
	_apprunnerObservabilityConfiguration     string
	_apprunnerObservabilityConfigurationArn  string
	_apprunnerObservabilityConfigurationName string
	_apprunnerProviderType                   string
	_apprunnerResourceArn                    string
	_apprunnerSecurityGroups                 []string
	_apprunnerServiceArn                     string
	_apprunnerServiceName                    string
	_apprunnerSourceConfiguration            string
	_apprunnerSubnets                        []string
	_apprunnerTagKeys                        []string
	_apprunnerTags                           string
	_apprunnerTraceConfiguration             string
	_apprunnerVpcConnectorArn                string
	_apprunnerVpcConnectorName               string
	_apprunnerVpcIngressConnectionArn        string
	_apprunnerVpcIngressConnectionName       string
)

// Associate your own domain name with the App Runner subdomain URL of your App
// Runner service.
//
// After you call AssociateCustomDomain and receive a successful response, use the
// information in the CustomDomainrecord that's returned to add CNAME records to your Domain
// Name System (DNS). For each mapped domain name, add a mapping to the target App
// Runner subdomain and one or more certificate validation records. App Runner then
// performs DNS validation to verify that you own or control the domain name that
// you associated. App Runner tracks domain validity in a certificate stored in [AWS Certificate Manager (ACM)].
//
// [AWS Certificate Manager (ACM)]: https://docs.aws.amazon.com/acm/latest/userguide
func apprunner_AssociateCustomDomain(cfg aws.Config, client *apprunner.Client) {
	input := &apprunner.AssociateCustomDomainInput{
		// DomainName: *string, // Required
		// ServiceArn: *string, // Required
	}

	if len(_apprunnerDomainName) > 0 {
		input.DomainName = aws.String(_apprunnerDomainName)
	}
	if len(_apprunnerServiceArn) > 0 {
		input.ServiceArn = aws.String(_apprunnerServiceArn)
	}
	if len(_apprunnerEnableWWWSubdomain) > 0 {
		if err := assignInputField(input, "EnableWWWSubdomain", _apprunnerEnableWWWSubdomain); err != nil {
			log.Errorf("invalid --enable-www-subdomain: %s", err.Error())
			return
		}
	}

	if resp, err := client.AssociateCustomDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create an App Runner automatic scaling configuration resource. App Runner
// requires this resource when you create or update App Runner services and you
// require non-default auto scaling settings. You can share an auto scaling
// configuration across multiple services.
//
// Create multiple revisions of a configuration by calling this action multiple
// times using the same AutoScalingConfigurationName . The call returns incremental
// AutoScalingConfigurationRevision values. When you create a service and configure
// an auto scaling configuration resource, the service uses the latest active
// revision of the auto scaling configuration by default. You can optionally
// configure the service to use a specific revision.
//
// Configure a higher MinSize to increase the spread of your App Runner service
// over more Availability Zones in the Amazon Web Services Region. The tradeoff is
// a higher minimal cost.
//
// Configure a lower MaxSize to control your cost. The tradeoff is lower
// responsiveness during peak demand.
func apprunner_CreateAutoScalingConfiguration(cfg aws.Config, client *apprunner.Client) {
	input := &apprunner.CreateAutoScalingConfigurationInput{
		// AutoScalingConfigurationName: *string, // Required
	}

	if len(_apprunnerAutoScalingConfigurationName) > 0 {
		input.AutoScalingConfigurationName = aws.String(_apprunnerAutoScalingConfigurationName)
	}
	if len(_apprunnerMaxConcurrency) > 0 {
		if err := assignInputField(input, "MaxConcurrency", _apprunnerMaxConcurrency); err != nil {
			log.Errorf("invalid --max-concurrency: %s", err.Error())
			return
		}
	}
	if len(_apprunnerMaxSize) > 0 {
		if err := assignInputField(input, "MaxSize", _apprunnerMaxSize); err != nil {
			log.Errorf("invalid --max-size: %s", err.Error())
			return
		}
	}
	if len(_apprunnerMinSize) > 0 {
		if err := assignInputField(input, "MinSize", _apprunnerMinSize); err != nil {
			log.Errorf("invalid --min-size: %s", err.Error())
			return
		}
	}
	if len(_apprunnerTags) > 0 {
		if err := assignInputField(input, "Tags", _apprunnerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAutoScalingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create an App Runner connection resource. App Runner requires a connection
// resource when you create App Runner services that access private repositories
// from certain third-party providers. You can share a connection across multiple
// services.
//
// A connection resource is needed to access GitHub and Bitbucket repositories.
// Both require a user interface approval process through the App Runner console
// before you can use the connection.
func apprunner_CreateConnection(cfg aws.Config, client *apprunner.Client) {
	input := &apprunner.CreateConnectionInput{
		// ConnectionName: *string, // Required
		// ProviderType: types.ProviderType, // Required
	}

	if len(_apprunnerConnectionName) > 0 {
		input.ConnectionName = aws.String(_apprunnerConnectionName)
	}
	if len(_apprunnerProviderType) > 0 {
		if err := assignInputField(input, "ProviderType", _apprunnerProviderType); err != nil {
			log.Errorf("invalid --provider-type: %s", err.Error())
			return
		}
	}
	if len(_apprunnerTags) > 0 {
		if err := assignInputField(input, "Tags", _apprunnerTags); err != nil {
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

// Create an App Runner observability configuration resource. App Runner requires
// this resource when you create or update App Runner services and you want to
// enable non-default observability features. You can share an observability
// configuration across multiple services.
//
// Create multiple revisions of a configuration by calling this action multiple
// times using the same ObservabilityConfigurationName . The call returns
// incremental ObservabilityConfigurationRevision values. When you create a
// service and configure an observability configuration resource, the service uses
// the latest active revision of the observability configuration by default. You
// can optionally configure the service to use a specific revision.
//
// The observability configuration resource is designed to configure multiple
// features (currently one feature, tracing). This action takes optional parameters
// that describe the configuration of these features (currently one parameter,
// TraceConfiguration ). If you don't specify a feature parameter, App Runner
// doesn't enable the feature.
func apprunner_CreateObservabilityConfiguration(cfg aws.Config, client *apprunner.Client) {
	input := &apprunner.CreateObservabilityConfigurationInput{
		// ObservabilityConfigurationName: *string, // Required
	}

	if len(_apprunnerObservabilityConfigurationName) > 0 {
		input.ObservabilityConfigurationName = aws.String(_apprunnerObservabilityConfigurationName)
	}
	if len(_apprunnerTags) > 0 {
		if err := assignInputField(input, "Tags", _apprunnerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_apprunnerTraceConfiguration) > 0 {
		if err := assignInputField(input, "TraceConfiguration", _apprunnerTraceConfiguration); err != nil {
			log.Errorf("invalid --trace-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateObservabilityConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create an App Runner service. After the service is created, the action also
// automatically starts a deployment.
//
// This is an asynchronous operation. On a successful call, you can use the
// returned OperationId and the [ListOperations] call to track the operation's progress.
//
// [ListOperations]: https://docs.aws.amazon.com/apprunner/latest/api/API_ListOperations.html
func apprunner_CreateService(cfg aws.Config, client *apprunner.Client) {
	input := &apprunner.CreateServiceInput{
		// ServiceName: *string, // Required
		// SourceConfiguration: *types.SourceConfiguration, // Required
	}

	if len(_apprunnerServiceName) > 0 {
		input.ServiceName = aws.String(_apprunnerServiceName)
	}
	if len(_apprunnerSourceConfiguration) > 0 {
		if err := assignInputField(input, "SourceConfiguration", _apprunnerSourceConfiguration); err != nil {
			log.Errorf("invalid --source-configuration: %s", err.Error())
			return
		}
	}
	if len(_apprunnerAutoScalingConfigurationArn) > 0 {
		input.AutoScalingConfigurationArn = aws.String(_apprunnerAutoScalingConfigurationArn)
	}
	if len(_apprunnerEncryptionConfiguration) > 0 {
		if err := assignInputField(input, "EncryptionConfiguration", _apprunnerEncryptionConfiguration); err != nil {
			log.Errorf("invalid --encryption-configuration: %s", err.Error())
			return
		}
	}
	if len(_apprunnerHealthCheckConfiguration) > 0 {
		if err := assignInputField(input, "HealthCheckConfiguration", _apprunnerHealthCheckConfiguration); err != nil {
			log.Errorf("invalid --health-check-configuration: %s", err.Error())
			return
		}
	}
	if len(_apprunnerInstanceConfiguration) > 0 {
		if err := assignInputField(input, "InstanceConfiguration", _apprunnerInstanceConfiguration); err != nil {
			log.Errorf("invalid --instance-configuration: %s", err.Error())
			return
		}
	}
	if len(_apprunnerNetworkConfiguration) > 0 {
		if err := assignInputField(input, "NetworkConfiguration", _apprunnerNetworkConfiguration); err != nil {
			log.Errorf("invalid --network-configuration: %s", err.Error())
			return
		}
	}
	if len(_apprunnerObservabilityConfiguration) > 0 {
		if err := assignInputField(input, "ObservabilityConfiguration", _apprunnerObservabilityConfiguration); err != nil {
			log.Errorf("invalid --observability-configuration: %s", err.Error())
			return
		}
	}
	if len(_apprunnerTags) > 0 {
		if err := assignInputField(input, "Tags", _apprunnerTags); err != nil {
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

// Create an App Runner VPC connector resource. App Runner requires this resource
// when you want to associate your App Runner service to a custom Amazon Virtual
// Private Cloud (Amazon VPC).
func apprunner_CreateVpcConnector(cfg aws.Config, client *apprunner.Client) {
	input := &apprunner.CreateVpcConnectorInput{
		// Subnets: []string, // Required
		// VpcConnectorName: *string, // Required
	}

	if len(_apprunnerSubnets) > 0 {
		input.Subnets = append([]string(nil), _apprunnerSubnets...)
	}
	if len(_apprunnerVpcConnectorName) > 0 {
		input.VpcConnectorName = aws.String(_apprunnerVpcConnectorName)
	}
	if len(_apprunnerSecurityGroups) > 0 {
		input.SecurityGroups = append([]string(nil), _apprunnerSecurityGroups...)
	}
	if len(_apprunnerTags) > 0 {
		if err := assignInputField(input, "Tags", _apprunnerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateVpcConnector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create an App Runner VPC Ingress Connection resource. App Runner requires this
// resource when you want to associate your App Runner service with an Amazon VPC
// endpoint.
func apprunner_CreateVpcIngressConnection(cfg aws.Config, client *apprunner.Client) {
	input := &apprunner.CreateVpcIngressConnectionInput{
		// IngressVpcConfiguration: *types.IngressVpcConfiguration, // Required
		// ServiceArn: *string, // Required
		// VpcIngressConnectionName: *string, // Required
	}

	if len(_apprunnerIngressVpcConfiguration) > 0 {
		if err := assignInputField(input, "IngressVpcConfiguration", _apprunnerIngressVpcConfiguration); err != nil {
			log.Errorf("invalid --ingress-vpc-configuration: %s", err.Error())
			return
		}
	}
	if len(_apprunnerServiceArn) > 0 {
		input.ServiceArn = aws.String(_apprunnerServiceArn)
	}
	if len(_apprunnerVpcIngressConnectionName) > 0 {
		input.VpcIngressConnectionName = aws.String(_apprunnerVpcIngressConnectionName)
	}
	if len(_apprunnerTags) > 0 {
		if err := assignInputField(input, "Tags", _apprunnerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateVpcIngressConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete an App Runner automatic scaling configuration resource. You can delete a
// top level auto scaling configuration, a specific revision of one, or all
// revisions associated with the top level configuration. You can't delete the
// default auto scaling configuration or a configuration that's used by one or more
// App Runner services.
func apprunner_DeleteAutoScalingConfiguration(cfg aws.Config, client *apprunner.Client) {
	input := &apprunner.DeleteAutoScalingConfigurationInput{
		// AutoScalingConfigurationArn: *string, // Required
	}

	if len(_apprunnerAutoScalingConfigurationArn) > 0 {
		input.AutoScalingConfigurationArn = aws.String(_apprunnerAutoScalingConfigurationArn)
	}
	if len(_apprunnerDeleteAllRevisions) > 0 {
		if err := assignInputField(input, "DeleteAllRevisions", _apprunnerDeleteAllRevisions); err != nil {
			log.Errorf("invalid --delete-all-revisions: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteAutoScalingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete an App Runner connection. You must first ensure that there are no
// running App Runner services that use this connection. If there are any, the
// DeleteConnection action fails.
func apprunner_DeleteConnection(cfg aws.Config, client *apprunner.Client) {
	input := &apprunner.DeleteConnectionInput{
		// ConnectionArn: *string, // Required
	}

	if len(_apprunnerConnectionArn) > 0 {
		input.ConnectionArn = aws.String(_apprunnerConnectionArn)
	}

	if resp, err := client.DeleteConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete an App Runner observability configuration resource. You can delete a
// specific revision or the latest active revision. You can't delete a
// configuration that's used by one or more App Runner services.
func apprunner_DeleteObservabilityConfiguration(cfg aws.Config, client *apprunner.Client) {
	input := &apprunner.DeleteObservabilityConfigurationInput{
		// ObservabilityConfigurationArn: *string, // Required
	}

	if len(_apprunnerObservabilityConfigurationArn) > 0 {
		input.ObservabilityConfigurationArn = aws.String(_apprunnerObservabilityConfigurationArn)
	}

	if resp, err := client.DeleteObservabilityConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete an App Runner service.
// This is an asynchronous operation. On a successful call, you can use the
// returned OperationId and the ListOperations call to track the operation's progress.
//
// Make sure that you don't have any active VPCIngressConnections associated with
// the service you want to delete.
func apprunner_DeleteService(cfg aws.Config, client *apprunner.Client) {
	input := &apprunner.DeleteServiceInput{
		// ServiceArn: *string, // Required
	}

	if len(_apprunnerServiceArn) > 0 {
		input.ServiceArn = aws.String(_apprunnerServiceArn)
	}

	if resp, err := client.DeleteService(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete an App Runner VPC connector resource. You can't delete a connector
// that's used by one or more App Runner services.
func apprunner_DeleteVpcConnector(cfg aws.Config, client *apprunner.Client) {
	input := &apprunner.DeleteVpcConnectorInput{
		// VpcConnectorArn: *string, // Required
	}

	if len(_apprunnerVpcConnectorArn) > 0 {
		input.VpcConnectorArn = aws.String(_apprunnerVpcConnectorArn)
	}

	if resp, err := client.DeleteVpcConnector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete an App Runner VPC Ingress Connection resource that's associated with an
// App Runner service. The VPC Ingress Connection must be in one of the following
// states to be deleted:
//
// - AVAILABLE
//
// - FAILED_CREATION
//
// - FAILED_UPDATE
//
// - FAILED_DELETION
func apprunner_DeleteVpcIngressConnection(cfg aws.Config, client *apprunner.Client) {
	input := &apprunner.DeleteVpcIngressConnectionInput{
		// VpcIngressConnectionArn: *string, // Required
	}

	if len(_apprunnerVpcIngressConnectionArn) > 0 {
		input.VpcIngressConnectionArn = aws.String(_apprunnerVpcIngressConnectionArn)
	}

	if resp, err := client.DeleteVpcIngressConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Return a full description of an App Runner automatic scaling configuration
// resource.
func apprunner_DescribeAutoScalingConfiguration(cfg aws.Config, client *apprunner.Client) {
	input := &apprunner.DescribeAutoScalingConfigurationInput{
		// AutoScalingConfigurationArn: *string, // Required
	}

	if len(_apprunnerAutoScalingConfigurationArn) > 0 {
		input.AutoScalingConfigurationArn = aws.String(_apprunnerAutoScalingConfigurationArn)
	}

	if resp, err := client.DescribeAutoScalingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Return a description of custom domain names that are associated with an App
// Runner service.
func apprunner_DescribeCustomDomains(cfg aws.Config, client *apprunner.Client) {
	input := &apprunner.DescribeCustomDomainsInput{
		// ServiceArn: *string, // Required
	}

	if len(_apprunnerServiceArn) > 0 {
		input.ServiceArn = aws.String(_apprunnerServiceArn)
	}
	if len(_apprunnerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _apprunnerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_apprunnerNextToken) > 0 {
		input.NextToken = aws.String(_apprunnerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeCustomDomains(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*apprunner.DescribeCustomDomainsOutput
	p := apprunner.NewDescribeCustomDomainsPaginator(client, input)
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

// Return a full description of an App Runner observability configuration resource.
func apprunner_DescribeObservabilityConfiguration(cfg aws.Config, client *apprunner.Client) {
	input := &apprunner.DescribeObservabilityConfigurationInput{
		// ObservabilityConfigurationArn: *string, // Required
	}

	if len(_apprunnerObservabilityConfigurationArn) > 0 {
		input.ObservabilityConfigurationArn = aws.String(_apprunnerObservabilityConfigurationArn)
	}

	if resp, err := client.DescribeObservabilityConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Return a full description of an App Runner service.
func apprunner_DescribeService(cfg aws.Config, client *apprunner.Client) {
	input := &apprunner.DescribeServiceInput{
		// ServiceArn: *string, // Required
	}

	if len(_apprunnerServiceArn) > 0 {
		input.ServiceArn = aws.String(_apprunnerServiceArn)
	}

	if resp, err := client.DescribeService(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Return a description of an App Runner VPC connector resource.
func apprunner_DescribeVpcConnector(cfg aws.Config, client *apprunner.Client) {
	input := &apprunner.DescribeVpcConnectorInput{
		// VpcConnectorArn: *string, // Required
	}

	if len(_apprunnerVpcConnectorArn) > 0 {
		input.VpcConnectorArn = aws.String(_apprunnerVpcConnectorArn)
	}

	if resp, err := client.DescribeVpcConnector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Return a full description of an App Runner VPC Ingress Connection resource.
func apprunner_DescribeVpcIngressConnection(cfg aws.Config, client *apprunner.Client) {
	input := &apprunner.DescribeVpcIngressConnectionInput{
		// VpcIngressConnectionArn: *string, // Required
	}

	if len(_apprunnerVpcIngressConnectionArn) > 0 {
		input.VpcIngressConnectionArn = aws.String(_apprunnerVpcIngressConnectionArn)
	}

	if resp, err := client.DescribeVpcIngressConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociate a custom domain name from an App Runner service.
// Certificates tracking domain validity are associated with a custom domain and
// are stored in [AWS Certificate Manager (ACM)]. These certificates aren't deleted as part of this action. App
// Runner delays certificate deletion for 30 days after a domain is disassociated
// from your service.
//
// [AWS Certificate Manager (ACM)]: https://docs.aws.amazon.com/acm/latest/userguide
func apprunner_DisassociateCustomDomain(cfg aws.Config, client *apprunner.Client) {
	input := &apprunner.DisassociateCustomDomainInput{
		// DomainName: *string, // Required
		// ServiceArn: *string, // Required
	}

	if len(_apprunnerDomainName) > 0 {
		input.DomainName = aws.String(_apprunnerDomainName)
	}
	if len(_apprunnerServiceArn) > 0 {
		input.ServiceArn = aws.String(_apprunnerServiceArn)
	}

	if resp, err := client.DisassociateCustomDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of active App Runner automatic scaling configurations in your
// Amazon Web Services account. You can query the revisions for a specific
// configuration name or the revisions for all active configurations in your
// account. You can optionally query only the latest revision of each requested
// name.
//
// To retrieve a full description of a particular configuration revision, call and
// provide one of the ARNs returned by ListAutoScalingConfigurations .
func apprunner_ListAutoScalingConfigurations(cfg aws.Config, client *apprunner.Client) {
	input := &apprunner.ListAutoScalingConfigurationsInput{}

	if len(_apprunnerAutoScalingConfigurationName) > 0 {
		input.AutoScalingConfigurationName = aws.String(_apprunnerAutoScalingConfigurationName)
	}
	if len(_apprunnerLatestOnly) > 0 {
		if err := assignInputField(input, "LatestOnly", _apprunnerLatestOnly); err != nil {
			log.Errorf("invalid --latest-only: %s", err.Error())
			return
		}
	}
	if len(_apprunnerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _apprunnerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_apprunnerNextToken) > 0 {
		input.NextToken = aws.String(_apprunnerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAutoScalingConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*apprunner.ListAutoScalingConfigurationsOutput
	p := apprunner.NewListAutoScalingConfigurationsPaginator(client, input)
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

// Returns a list of App Runner connections that are associated with your Amazon
// Web Services account.
func apprunner_ListConnections(cfg aws.Config, client *apprunner.Client) {
	input := &apprunner.ListConnectionsInput{}

	if len(_apprunnerConnectionName) > 0 {
		input.ConnectionName = aws.String(_apprunnerConnectionName)
	}
	if len(_apprunnerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _apprunnerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_apprunnerNextToken) > 0 {
		input.NextToken = aws.String(_apprunnerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListConnections(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*apprunner.ListConnectionsOutput
	p := apprunner.NewListConnectionsPaginator(client, input)
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

// Returns a list of active App Runner observability configurations in your Amazon
// Web Services account. You can query the revisions for a specific configuration
// name or the revisions for all active configurations in your account. You can
// optionally query only the latest revision of each requested name.
//
// To retrieve a full description of a particular configuration revision, call and
// provide one of the ARNs returned by ListObservabilityConfigurations .
func apprunner_ListObservabilityConfigurations(cfg aws.Config, client *apprunner.Client) {
	input := &apprunner.ListObservabilityConfigurationsInput{}

	if len(_apprunnerLatestOnly) > 0 {
		if err := assignInputField(input, "LatestOnly", _apprunnerLatestOnly); err != nil {
			log.Errorf("invalid --latest-only: %s", err.Error())
			return
		}
	}
	if len(_apprunnerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _apprunnerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_apprunnerNextToken) > 0 {
		input.NextToken = aws.String(_apprunnerNextToken)
	}
	if len(_apprunnerObservabilityConfigurationName) > 0 {
		input.ObservabilityConfigurationName = aws.String(_apprunnerObservabilityConfigurationName)
	}

	if disablePaginator() {
		if resp, err := client.ListObservabilityConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*apprunner.ListObservabilityConfigurationsOutput
	p := apprunner.NewListObservabilityConfigurationsPaginator(client, input)
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

// Return a list of operations that occurred on an App Runner service.
// The resulting list of OperationSummary objects is sorted in reverse chronological order. The
// first object on the list represents the last started operation.
func apprunner_ListOperations(cfg aws.Config, client *apprunner.Client) {
	input := &apprunner.ListOperationsInput{
		// ServiceArn: *string, // Required
	}

	if len(_apprunnerServiceArn) > 0 {
		input.ServiceArn = aws.String(_apprunnerServiceArn)
	}
	if len(_apprunnerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _apprunnerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_apprunnerNextToken) > 0 {
		input.NextToken = aws.String(_apprunnerNextToken)
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

	var results []*apprunner.ListOperationsOutput
	p := apprunner.NewListOperationsPaginator(client, input)
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

// Returns a list of running App Runner services in your Amazon Web Services
// account.
func apprunner_ListServices(cfg aws.Config, client *apprunner.Client) {
	input := &apprunner.ListServicesInput{}

	if len(_apprunnerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _apprunnerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_apprunnerNextToken) > 0 {
		input.NextToken = aws.String(_apprunnerNextToken)
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

	var results []*apprunner.ListServicesOutput
	p := apprunner.NewListServicesPaginator(client, input)
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

// Returns a list of the associated App Runner services using an auto scaling
// configuration.
func apprunner_ListServicesForAutoScalingConfiguration(cfg aws.Config, client *apprunner.Client) {
	input := &apprunner.ListServicesForAutoScalingConfigurationInput{
		// AutoScalingConfigurationArn: *string, // Required
	}

	if len(_apprunnerAutoScalingConfigurationArn) > 0 {
		input.AutoScalingConfigurationArn = aws.String(_apprunnerAutoScalingConfigurationArn)
	}
	if len(_apprunnerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _apprunnerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_apprunnerNextToken) > 0 {
		input.NextToken = aws.String(_apprunnerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListServicesForAutoScalingConfiguration(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*apprunner.ListServicesForAutoScalingConfigurationOutput
	p := apprunner.NewListServicesForAutoScalingConfigurationPaginator(client, input)
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

// List tags that are associated with for an App Runner resource. The response
// contains a list of tag key-value pairs.
func apprunner_ListTagsForResource(cfg aws.Config, client *apprunner.Client) {
	input := &apprunner.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_apprunnerResourceArn) > 0 {
		input.ResourceArn = aws.String(_apprunnerResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of App Runner VPC connectors in your Amazon Web Services account.
func apprunner_ListVpcConnectors(cfg aws.Config, client *apprunner.Client) {
	input := &apprunner.ListVpcConnectorsInput{}

	if len(_apprunnerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _apprunnerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_apprunnerNextToken) > 0 {
		input.NextToken = aws.String(_apprunnerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListVpcConnectors(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*apprunner.ListVpcConnectorsOutput
	p := apprunner.NewListVpcConnectorsPaginator(client, input)
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

// Return a list of App Runner VPC Ingress Connections in your Amazon Web Services
// account.
func apprunner_ListVpcIngressConnections(cfg aws.Config, client *apprunner.Client) {
	input := &apprunner.ListVpcIngressConnectionsInput{}

	if len(_apprunnerFilter) > 0 {
		if err := assignInputField(input, "Filter", _apprunnerFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_apprunnerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _apprunnerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_apprunnerNextToken) > 0 {
		input.NextToken = aws.String(_apprunnerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListVpcIngressConnections(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*apprunner.ListVpcIngressConnectionsOutput
	p := apprunner.NewListVpcIngressConnectionsPaginator(client, input)
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

// Pause an active App Runner service. App Runner reduces compute capacity for the
// service to zero and loses state (for example, ephemeral storage is removed).
//
// This is an asynchronous operation. On a successful call, you can use the
// returned OperationId and the ListOperations call to track the operation's progress.
func apprunner_PauseService(cfg aws.Config, client *apprunner.Client) {
	input := &apprunner.PauseServiceInput{
		// ServiceArn: *string, // Required
	}

	if len(_apprunnerServiceArn) > 0 {
		input.ServiceArn = aws.String(_apprunnerServiceArn)
	}

	if resp, err := client.PauseService(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Resume an active App Runner service. App Runner provisions compute capacity for
// the service.
//
// This is an asynchronous operation. On a successful call, you can use the
// returned OperationId and the ListOperations call to track the operation's progress.
func apprunner_ResumeService(cfg aws.Config, client *apprunner.Client) {
	input := &apprunner.ResumeServiceInput{
		// ServiceArn: *string, // Required
	}

	if len(_apprunnerServiceArn) > 0 {
		input.ServiceArn = aws.String(_apprunnerServiceArn)
	}

	if resp, err := client.ResumeService(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initiate a manual deployment of the latest commit in a source code repository
// or the latest image in a source image repository to an App Runner service.
//
// For a source code repository, App Runner retrieves the commit and builds a
// Docker image. For a source image repository, App Runner retrieves the latest
// Docker image. In both cases, App Runner then deploys the new image to your
// service and starts a new container instance.
//
// This is an asynchronous operation. On a successful call, you can use the
// returned OperationId and the ListOperations call to track the operation's progress.
func apprunner_StartDeployment(cfg aws.Config, client *apprunner.Client) {
	input := &apprunner.StartDeploymentInput{
		// ServiceArn: *string, // Required
	}

	if len(_apprunnerServiceArn) > 0 {
		input.ServiceArn = aws.String(_apprunnerServiceArn)
	}

	if resp, err := client.StartDeployment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Add tags to, or update the tag values of, an App Runner resource. A tag is a
// key-value pair.
func apprunner_TagResource(cfg aws.Config, client *apprunner.Client) {
	input := &apprunner.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_apprunnerResourceArn) > 0 {
		input.ResourceArn = aws.String(_apprunnerResourceArn)
	}
	if len(_apprunnerTags) > 0 {
		if err := assignInputField(input, "Tags", _apprunnerTags); err != nil {
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

// Remove tags from an App Runner resource.
func apprunner_UntagResource(cfg aws.Config, client *apprunner.Client) {
	input := &apprunner.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_apprunnerResourceArn) > 0 {
		input.ResourceArn = aws.String(_apprunnerResourceArn)
	}
	if len(_apprunnerTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _apprunnerTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update an auto scaling configuration to be the default. The existing default
// auto scaling configuration will be set to non-default automatically.
func apprunner_UpdateDefaultAutoScalingConfiguration(cfg aws.Config, client *apprunner.Client) {
	input := &apprunner.UpdateDefaultAutoScalingConfigurationInput{
		// AutoScalingConfigurationArn: *string, // Required
	}

	if len(_apprunnerAutoScalingConfigurationArn) > 0 {
		input.AutoScalingConfigurationArn = aws.String(_apprunnerAutoScalingConfigurationArn)
	}

	if resp, err := client.UpdateDefaultAutoScalingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update an App Runner service. You can update the source configuration and
// instance configuration of the service. You can also update the ARN of the auto
// scaling configuration resource that's associated with the service. However, you
// can't change the name or the encryption configuration of the service. These can
// be set only when you create the service.
//
// To update the tags applied to your service, use the separate actions TagResource and UntagResource.
//
// This is an asynchronous operation. On a successful call, you can use the
// returned OperationId and the ListOperations call to track the operation's progress.
func apprunner_UpdateService(cfg aws.Config, client *apprunner.Client) {
	input := &apprunner.UpdateServiceInput{
		// ServiceArn: *string, // Required
	}

	if len(_apprunnerServiceArn) > 0 {
		input.ServiceArn = aws.String(_apprunnerServiceArn)
	}
	if len(_apprunnerAutoScalingConfigurationArn) > 0 {
		input.AutoScalingConfigurationArn = aws.String(_apprunnerAutoScalingConfigurationArn)
	}
	if len(_apprunnerHealthCheckConfiguration) > 0 {
		if err := assignInputField(input, "HealthCheckConfiguration", _apprunnerHealthCheckConfiguration); err != nil {
			log.Errorf("invalid --health-check-configuration: %s", err.Error())
			return
		}
	}
	if len(_apprunnerInstanceConfiguration) > 0 {
		if err := assignInputField(input, "InstanceConfiguration", _apprunnerInstanceConfiguration); err != nil {
			log.Errorf("invalid --instance-configuration: %s", err.Error())
			return
		}
	}
	if len(_apprunnerNetworkConfiguration) > 0 {
		if err := assignInputField(input, "NetworkConfiguration", _apprunnerNetworkConfiguration); err != nil {
			log.Errorf("invalid --network-configuration: %s", err.Error())
			return
		}
	}
	if len(_apprunnerObservabilityConfiguration) > 0 {
		if err := assignInputField(input, "ObservabilityConfiguration", _apprunnerObservabilityConfiguration); err != nil {
			log.Errorf("invalid --observability-configuration: %s", err.Error())
			return
		}
	}
	if len(_apprunnerSourceConfiguration) > 0 {
		if err := assignInputField(input, "SourceConfiguration", _apprunnerSourceConfiguration); err != nil {
			log.Errorf("invalid --source-configuration: %s", err.Error())
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

// Update an existing App Runner VPC Ingress Connection resource. The VPC Ingress
// Connection must be in one of the following states to be updated:
//
// - AVAILABLE
//
// - FAILED_CREATION
//
// - FAILED_UPDATE
func apprunner_UpdateVpcIngressConnection(cfg aws.Config, client *apprunner.Client) {
	input := &apprunner.UpdateVpcIngressConnectionInput{
		// IngressVpcConfiguration: *types.IngressVpcConfiguration, // Required
		// VpcIngressConnectionArn: *string, // Required
	}

	if len(_apprunnerIngressVpcConfiguration) > 0 {
		if err := assignInputField(input, "IngressVpcConfiguration", _apprunnerIngressVpcConfiguration); err != nil {
			log.Errorf("invalid --ingress-vpc-configuration: %s", err.Error())
			return
		}
	}
	if len(_apprunnerVpcIngressConnectionArn) > 0 {
		input.VpcIngressConnectionArn = aws.String(_apprunnerVpcIngressConnectionArn)
	}

	if resp, err := client.UpdateVpcIngressConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_apprunnerCmd)
	_apprunnerCmd.Flags().SortFlags = false

	_apprunnerCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_apprunnerCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_apprunnerCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_apprunnerCmd.Flags().StringVarP(&_apprunnerAutoScalingConfigurationArn, "auto-scaling-configuration-arn", "", "", "Auto Scaling Configuration ARN")
	_apprunnerCmd.Flags().StringVarP(&_apprunnerAutoScalingConfigurationName, "auto-scaling-configuration-name", "", "", "Auto Scaling Configuration Name")
	_apprunnerCmd.Flags().StringVarP(&_apprunnerConnectionArn, "connection-arn", "", "", "Connection ARN")
	_apprunnerCmd.Flags().StringVarP(&_apprunnerConnectionName, "connection-name", "", "", "Connection Name")
	_apprunnerCmd.Flags().StringVarP(&_apprunnerDeleteAllRevisions, "delete-all-revisions", "", "", "Delete All Revisions")
	_apprunnerCmd.Flags().StringVarP(&_apprunnerDomainName, "domain-name", "", "", "Domain Name")
	_apprunnerCmd.Flags().StringVarP(&_apprunnerEnableWWWSubdomain, "enable-www-subdomain", "", "", "Enable Www Subdomain")
	_apprunnerCmd.Flags().StringVarP(&_apprunnerEncryptionConfiguration, "encryption-configuration", "", "", "Encryption Configuration")
	_apprunnerCmd.Flags().StringVarP(&_apprunnerFilter, "filter", "", "", "Filter")
	_apprunnerCmd.Flags().StringVarP(&_apprunnerHealthCheckConfiguration, "health-check-configuration", "", "", "Health Check Configuration")
	_apprunnerCmd.Flags().StringVarP(&_apprunnerIngressVpcConfiguration, "ingress-vpc-configuration", "", "", "Ingress VPC Configuration")
	_apprunnerCmd.Flags().StringVarP(&_apprunnerInstanceConfiguration, "instance-configuration", "", "", "Instance Configuration")
	_apprunnerCmd.Flags().StringVarP(&_apprunnerLatestOnly, "latest-only", "", "", "Latest Only")
	_apprunnerCmd.Flags().StringVarP(&_apprunnerMaxConcurrency, "max-concurrency", "", "", "Max Concurrency")
	_apprunnerCmd.Flags().StringVarP(&_apprunnerMaxResults, "max-results", "", "", "Max Results")
	_apprunnerCmd.Flags().StringVarP(&_apprunnerMaxSize, "max-size", "", "", "Max Size")
	_apprunnerCmd.Flags().StringVarP(&_apprunnerMinSize, "min-size", "", "", "Min Size")
	_apprunnerCmd.Flags().StringVarP(&_apprunnerNetworkConfiguration, "network-configuration", "", "", "Network Configuration")
	_apprunnerCmd.Flags().StringVarP(&_apprunnerNextToken, "next-token", "", "", "Next Token")
	_apprunnerCmd.Flags().StringVarP(&_apprunnerObservabilityConfiguration, "observability-configuration", "", "", "Observability Configuration")
	_apprunnerCmd.Flags().StringVarP(&_apprunnerObservabilityConfigurationArn, "observability-configuration-arn", "", "", "Observability Configuration ARN")
	_apprunnerCmd.Flags().StringVarP(&_apprunnerObservabilityConfigurationName, "observability-configuration-name", "", "", "Observability Configuration Name")
	_apprunnerCmd.Flags().StringVarP(&_apprunnerProviderType, "provider-type", "", "", "Provider Type")
	_apprunnerCmd.Flags().StringVarP(&_apprunnerResourceArn, "resource-arn", "", "", "Resource ARN")
	_apprunnerCmd.Flags().StringSliceVarP(&_apprunnerSecurityGroups, "security-groups", "", nil, "Security Groups")
	_apprunnerCmd.Flags().StringVarP(&_apprunnerServiceArn, "service-arn", "", "", "Service ARN")
	_apprunnerCmd.Flags().StringVarP(&_apprunnerServiceName, "service-name", "", "", "Service Name")
	_apprunnerCmd.Flags().StringVarP(&_apprunnerSourceConfiguration, "source-configuration", "", "", "Source Configuration")
	_apprunnerCmd.Flags().StringSliceVarP(&_apprunnerSubnets, "subnets", "", nil, "Subnets")
	_apprunnerCmd.Flags().StringSliceVarP(&_apprunnerTagKeys, "tag-keys", "", nil, "Tag Keys")
	_apprunnerCmd.Flags().StringVarP(&_apprunnerTags, "tags", "", "", "Tags")
	_apprunnerCmd.Flags().StringVarP(&_apprunnerTraceConfiguration, "trace-configuration", "", "", "Trace Configuration")
	_apprunnerCmd.Flags().StringVarP(&_apprunnerVpcConnectorArn, "vpc-connector-arn", "", "", "VPC Connector ARN")
	_apprunnerCmd.Flags().StringVarP(&_apprunnerVpcConnectorName, "vpc-connector-name", "", "", "VPC Connector Name")
	_apprunnerCmd.Flags().StringVarP(&_apprunnerVpcIngressConnectionArn, "vpc-ingress-connection-arn", "", "", "VPC Ingress Connection ARN")
	_apprunnerCmd.Flags().StringVarP(&_apprunnerVpcIngressConnectionName, "vpc-ingress-connection-name", "", "", "VPC Ingress Connection Name")

	_apprunnerCmd.Flags().BoolVarP(&_apprunnerAssociateCustomDomain, "associate-custom-domain", "", false, "Associate Custom Domain")
	_apprunnerCmd.Flags().BoolVarP(&_apprunnerCreateAutoScalingConfiguration, "create-auto-scaling-configuration", "", false, "Create Auto Scaling Configuration")
	_apprunnerCmd.Flags().BoolVarP(&_apprunnerCreateConnection, "create-connection", "", false, "Create Connection")
	_apprunnerCmd.Flags().BoolVarP(&_apprunnerCreateObservabilityConfiguration, "create-observability-configuration", "", false, "Create Observability Configuration")
	_apprunnerCmd.Flags().BoolVarP(&_apprunnerCreateService, "create-service", "", false, "Create Service")
	_apprunnerCmd.Flags().BoolVarP(&_apprunnerCreateVpcConnector, "create-vpc-connector", "", false, "Create VPC Connector")
	_apprunnerCmd.Flags().BoolVarP(&_apprunnerCreateVpcIngressConnection, "create-vpc-ingress-connection", "", false, "Create VPC Ingress Connection")
	_apprunnerCmd.Flags().BoolVarP(&_apprunnerDeleteAutoScalingConfiguration, "delete-auto-scaling-configuration", "", false, "Delete Auto Scaling Configuration")
	_apprunnerCmd.Flags().BoolVarP(&_apprunnerDeleteConnection, "delete-connection", "", false, "Delete Connection")
	_apprunnerCmd.Flags().BoolVarP(&_apprunnerDeleteObservabilityConfiguration, "delete-observability-configuration", "", false, "Delete Observability Configuration")
	_apprunnerCmd.Flags().BoolVarP(&_apprunnerDeleteService, "delete-service", "", false, "Delete Service")
	_apprunnerCmd.Flags().BoolVarP(&_apprunnerDeleteVpcConnector, "delete-vpc-connector", "", false, "Delete VPC Connector")
	_apprunnerCmd.Flags().BoolVarP(&_apprunnerDeleteVpcIngressConnection, "delete-vpc-ingress-connection", "", false, "Delete VPC Ingress Connection")
	_apprunnerCmd.Flags().BoolVarP(&_apprunnerDescribeAutoScalingConfiguration, "describe-auto-scaling-configuration", "", false, "Describe Auto Scaling Configuration")
	_apprunnerCmd.Flags().BoolVarP(&_apprunnerDescribeCustomDomains, "describe-custom-domains", "", false, "Describe Custom Domains")
	_apprunnerCmd.Flags().BoolVarP(&_apprunnerDescribeObservabilityConfiguration, "describe-observability-configuration", "", false, "Describe Observability Configuration")
	_apprunnerCmd.Flags().BoolVarP(&_apprunnerDescribeService, "describe-service", "", false, "Describe Service")
	_apprunnerCmd.Flags().BoolVarP(&_apprunnerDescribeVpcConnector, "describe-vpc-connector", "", false, "Describe VPC Connector")
	_apprunnerCmd.Flags().BoolVarP(&_apprunnerDescribeVpcIngressConnection, "describe-vpc-ingress-connection", "", false, "Describe VPC Ingress Connection")
	_apprunnerCmd.Flags().BoolVarP(&_apprunnerDisassociateCustomDomain, "disassociate-custom-domain", "", false, "Disassociate Custom Domain")
	_apprunnerCmd.Flags().BoolVarP(&_apprunnerListAutoScalingConfigurations, "list-auto-scaling-configurations", "", false, "List Auto Scaling Configurations")
	_apprunnerCmd.Flags().BoolVarP(&_apprunnerListConnections, "list-connections", "", false, "List Connections")
	_apprunnerCmd.Flags().BoolVarP(&_apprunnerListObservabilityConfigurations, "list-observability-configurations", "", false, "List Observability Configurations")
	_apprunnerCmd.Flags().BoolVarP(&_apprunnerListOperations, "list-operations", "", false, "List Operations")
	_apprunnerCmd.Flags().BoolVarP(&_apprunnerListServices, "list-services", "", false, "List Services")
	_apprunnerCmd.Flags().BoolVarP(&_apprunnerListServicesForAutoScalingConfiguration, "list-services-for-auto-scaling-configuration", "", false, "List Services For Auto Scaling Configuration")
	_apprunnerCmd.Flags().BoolVarP(&_apprunnerListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_apprunnerCmd.Flags().BoolVarP(&_apprunnerListVpcConnectors, "list-vpc-connectors", "", false, "List VPC Connectors")
	_apprunnerCmd.Flags().BoolVarP(&_apprunnerListVpcIngressConnections, "list-vpc-ingress-connections", "", false, "List VPC Ingress Connections")
	_apprunnerCmd.Flags().BoolVarP(&_apprunnerPauseService, "pause-service", "", false, "Pause Service")
	_apprunnerCmd.Flags().BoolVarP(&_apprunnerResumeService, "resume-service", "", false, "Resume Service")
	_apprunnerCmd.Flags().BoolVarP(&_apprunnerStartDeployment, "start-deployment", "", false, "Start Deployment")
	_apprunnerCmd.Flags().BoolVarP(&_apprunnerTagResource, "tag-resource", "", false, "Tag Resource")
	_apprunnerCmd.Flags().BoolVarP(&_apprunnerUntagResource, "untag-resource", "", false, "Untag Resource")
	_apprunnerCmd.Flags().BoolVarP(&_apprunnerUpdateDefaultAutoScalingConfiguration, "update-default-auto-scaling-configuration", "", false, "Update Default Auto Scaling Configuration")
	_apprunnerCmd.Flags().BoolVarP(&_apprunnerUpdateService, "update-service", "", false, "Update Service")
	_apprunnerCmd.Flags().BoolVarP(&_apprunnerUpdateVpcIngressConnection, "update-vpc-ingress-connection", "", false, "Update VPC Ingress Connection")

}
