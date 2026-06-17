package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/migrationhubrefactorspaces"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// migrationhubrefactorspacesCmd represents the migrationhubrefactorspaces command
var _migrationhubrefactorspacesCmd = &cobra.Command{
	Use:   "migrationhubrefactorspaces",
	Short: "AWS migrationhubrefactorspaces CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := migrationhubrefactorspaces.NewFromConfig(cfg)
		if _migrationhubrefactorspacesCreateApplication {
			migrationhubrefactorspaces_CreateApplication(cfg, client)
			return
		}
		if _migrationhubrefactorspacesCreateEnvironment {
			migrationhubrefactorspaces_CreateEnvironment(cfg, client)
			return
		}
		if _migrationhubrefactorspacesCreateRoute {
			migrationhubrefactorspaces_CreateRoute(cfg, client)
			return
		}
		if _migrationhubrefactorspacesCreateService {
			migrationhubrefactorspaces_CreateService(cfg, client)
			return
		}
		if _migrationhubrefactorspacesDeleteApplication {
			migrationhubrefactorspaces_DeleteApplication(cfg, client)
			return
		}
		if _migrationhubrefactorspacesDeleteEnvironment {
			migrationhubrefactorspaces_DeleteEnvironment(cfg, client)
			return
		}
		if _migrationhubrefactorspacesDeleteResourcePolicy {
			migrationhubrefactorspaces_DeleteResourcePolicy(cfg, client)
			return
		}
		if _migrationhubrefactorspacesDeleteRoute {
			migrationhubrefactorspaces_DeleteRoute(cfg, client)
			return
		}
		if _migrationhubrefactorspacesDeleteService {
			migrationhubrefactorspaces_DeleteService(cfg, client)
			return
		}
		if _migrationhubrefactorspacesGetApplication {
			migrationhubrefactorspaces_GetApplication(cfg, client)
			return
		}
		if _migrationhubrefactorspacesGetEnvironment {
			migrationhubrefactorspaces_GetEnvironment(cfg, client)
			return
		}
		if _migrationhubrefactorspacesGetResourcePolicy {
			migrationhubrefactorspaces_GetResourcePolicy(cfg, client)
			return
		}
		if _migrationhubrefactorspacesGetRoute {
			migrationhubrefactorspaces_GetRoute(cfg, client)
			return
		}
		if _migrationhubrefactorspacesGetService {
			migrationhubrefactorspaces_GetService(cfg, client)
			return
		}
		if _migrationhubrefactorspacesListApplications {
			migrationhubrefactorspaces_ListApplications(cfg, client)
			return
		}
		if _migrationhubrefactorspacesListEnvironmentVpcs {
			migrationhubrefactorspaces_ListEnvironmentVpcs(cfg, client)
			return
		}
		if _migrationhubrefactorspacesListEnvironments {
			migrationhubrefactorspaces_ListEnvironments(cfg, client)
			return
		}
		if _migrationhubrefactorspacesListRoutes {
			migrationhubrefactorspaces_ListRoutes(cfg, client)
			return
		}
		if _migrationhubrefactorspacesListServices {
			migrationhubrefactorspaces_ListServices(cfg, client)
			return
		}
		if _migrationhubrefactorspacesListTagsForResource {
			migrationhubrefactorspaces_ListTagsForResource(cfg, client)
			return
		}
		if _migrationhubrefactorspacesPutResourcePolicy {
			migrationhubrefactorspaces_PutResourcePolicy(cfg, client)
			return
		}
		if _migrationhubrefactorspacesTagResource {
			migrationhubrefactorspaces_TagResource(cfg, client)
			return
		}
		if _migrationhubrefactorspacesUntagResource {
			migrationhubrefactorspaces_UntagResource(cfg, client)
			return
		}
		if _migrationhubrefactorspacesUpdateRoute {
			migrationhubrefactorspaces_UpdateRoute(cfg, client)
			return
		}

	},
}

var (
	_migrationhubrefactorspacesCreateApplication    bool
	_migrationhubrefactorspacesCreateEnvironment    bool
	_migrationhubrefactorspacesCreateRoute          bool
	_migrationhubrefactorspacesCreateService        bool
	_migrationhubrefactorspacesDeleteApplication    bool
	_migrationhubrefactorspacesDeleteEnvironment    bool
	_migrationhubrefactorspacesDeleteResourcePolicy bool
	_migrationhubrefactorspacesDeleteRoute          bool
	_migrationhubrefactorspacesDeleteService        bool
	_migrationhubrefactorspacesGetApplication       bool
	_migrationhubrefactorspacesGetEnvironment       bool
	_migrationhubrefactorspacesGetResourcePolicy    bool
	_migrationhubrefactorspacesGetRoute             bool
	_migrationhubrefactorspacesGetService           bool
	_migrationhubrefactorspacesListApplications     bool
	_migrationhubrefactorspacesListEnvironmentVpcs  bool
	_migrationhubrefactorspacesListEnvironments     bool
	_migrationhubrefactorspacesListRoutes           bool
	_migrationhubrefactorspacesListServices         bool
	_migrationhubrefactorspacesListTagsForResource  bool
	_migrationhubrefactorspacesPutResourcePolicy    bool
	_migrationhubrefactorspacesTagResource          bool
	_migrationhubrefactorspacesUntagResource        bool
	_migrationhubrefactorspacesUpdateRoute          bool

	_migrationhubrefactorspacesActivationState       string
	_migrationhubrefactorspacesApiGatewayProxy       string
	_migrationhubrefactorspacesApplicationIdentifier string
	_migrationhubrefactorspacesClientToken           string
	_migrationhubrefactorspacesDefaultRoute          string
	_migrationhubrefactorspacesDescription           string
	_migrationhubrefactorspacesEndpointType          string
	_migrationhubrefactorspacesEnvironmentIdentifier string
	_migrationhubrefactorspacesIdentifier            string
	_migrationhubrefactorspacesLambdaEndpoint        string
	_migrationhubrefactorspacesMaxResults            string
	_migrationhubrefactorspacesName                  string
	_migrationhubrefactorspacesNetworkFabricType     string
	_migrationhubrefactorspacesNextToken             string
	_migrationhubrefactorspacesPolicy                string
	_migrationhubrefactorspacesProxyType             string
	_migrationhubrefactorspacesResourceArn           string
	_migrationhubrefactorspacesRouteIdentifier       string
	_migrationhubrefactorspacesRouteType             string
	_migrationhubrefactorspacesServiceIdentifier     string
	_migrationhubrefactorspacesTagKeys               []string
	_migrationhubrefactorspacesTags                  string
	_migrationhubrefactorspacesUriPathRoute          string
	_migrationhubrefactorspacesUrlEndpoint           string
	_migrationhubrefactorspacesVpcId                 string
)

// Creates an Amazon Web Services Migration Hub Refactor Spaces application. The
// account that owns the environment also owns the applications created inside the
// environment, regardless of the account that creates the application. Refactor
// Spaces provisions an Amazon API Gateway, API Gateway VPC link, and Network Load
// Balancer for the application proxy inside your account.
//
// In environments created with a [CreateEnvironment:NetworkFabricType] of NONE you need to configure [VPC to VPC connectivity] between your
// service VPC and the application proxy VPC to route traffic through the
// application proxy to a service with a private URL endpoint. For more
// information, see [Create an application]in the Refactor Spaces User Guide.
//
// [VPC to VPC connectivity]: https://docs.aws.amazon.com/whitepapers/latest/aws-vpc-connectivity-options/amazon-vpc-to-amazon-vpc-connectivity-options.html
// [Create an application]: https://docs.aws.amazon.com/migrationhub-refactor-spaces/latest/userguide/getting-started-create-application.html
// [CreateEnvironment:NetworkFabricType]: https://docs.aws.amazon.com/migrationhub-refactor-spaces/latest/APIReference/API_CreateEnvironment.html#migrationhubrefactorspaces-CreateEnvironment-request-NetworkFabricType
func migrationhubrefactorspaces_CreateApplication(cfg aws.Config, client *migrationhubrefactorspaces.Client) {
	input := &migrationhubrefactorspaces.CreateApplicationInput{
		// EnvironmentIdentifier: *string, // Required
		// Name: *string, // Required
		// ProxyType: types.ProxyType, // Required
		// VpcId: *string, // Required
	}

	if len(_migrationhubrefactorspacesEnvironmentIdentifier) > 0 {
		input.EnvironmentIdentifier = aws.String(_migrationhubrefactorspacesEnvironmentIdentifier)
	}
	if len(_migrationhubrefactorspacesName) > 0 {
		input.Name = aws.String(_migrationhubrefactorspacesName)
	}
	if len(_migrationhubrefactorspacesProxyType) > 0 {
		if err := assignInputField(input, "ProxyType", _migrationhubrefactorspacesProxyType); err != nil {
			log.Errorf("invalid --proxy-type: %s", err.Error())
			return
		}
	}
	if len(_migrationhubrefactorspacesVpcId) > 0 {
		input.VpcId = aws.String(_migrationhubrefactorspacesVpcId)
	}
	if len(_migrationhubrefactorspacesApiGatewayProxy) > 0 {
		if err := assignInputField(input, "ApiGatewayProxy", _migrationhubrefactorspacesApiGatewayProxy); err != nil {
			log.Errorf("invalid --api-gateway-proxy: %s", err.Error())
			return
		}
	}
	if len(_migrationhubrefactorspacesClientToken) > 0 {
		input.ClientToken = aws.String(_migrationhubrefactorspacesClientToken)
	}
	if len(_migrationhubrefactorspacesTags) > 0 {
		if err := assignInputField(input, "Tags", _migrationhubrefactorspacesTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon Web Services Migration Hub Refactor Spaces environment. The
// caller owns the environment resource, and all Refactor Spaces applications,
// services, and routes created within the environment. They are referred to as the
// environment owner. The environment owner has cross-account visibility and
// control of Refactor Spaces resources that are added to the environment by other
// accounts that the environment is shared with.
//
// When creating an environment with a [CreateEnvironment:NetworkFabricType] of TRANSIT_GATEWAY , Refactor Spaces
// provisions a transit gateway to enable services in VPCs to communicate directly
// across accounts. If [CreateEnvironment:NetworkFabricType]is NONE , Refactor Spaces does not create a transit gateway
// and you must use your network infrastructure to route traffic to services with
// private URL endpoints.
//
// [CreateEnvironment:NetworkFabricType]: https://docs.aws.amazon.com/migrationhub-refactor-spaces/latest/APIReference/API_CreateEnvironment.html#migrationhubrefactorspaces-CreateEnvironment-request-NetworkFabricType
func migrationhubrefactorspaces_CreateEnvironment(cfg aws.Config, client *migrationhubrefactorspaces.Client) {
	input := &migrationhubrefactorspaces.CreateEnvironmentInput{
		// Name: *string, // Required
		// NetworkFabricType: types.NetworkFabricType, // Required
	}

	if len(_migrationhubrefactorspacesName) > 0 {
		input.Name = aws.String(_migrationhubrefactorspacesName)
	}
	if len(_migrationhubrefactorspacesNetworkFabricType) > 0 {
		if err := assignInputField(input, "NetworkFabricType", _migrationhubrefactorspacesNetworkFabricType); err != nil {
			log.Errorf("invalid --network-fabric-type: %s", err.Error())
			return
		}
	}
	if len(_migrationhubrefactorspacesClientToken) > 0 {
		input.ClientToken = aws.String(_migrationhubrefactorspacesClientToken)
	}
	if len(_migrationhubrefactorspacesDescription) > 0 {
		input.Description = aws.String(_migrationhubrefactorspacesDescription)
	}
	if len(_migrationhubrefactorspacesTags) > 0 {
		if err := assignInputField(input, "Tags", _migrationhubrefactorspacesTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon Web Services Migration Hub Refactor Spaces route. The account
// owner of the service resource is always the environment owner, regardless of
// which account creates the route. Routes target a service in the application. If
// an application does not have any routes, then the first route must be created as
// a DEFAULT RouteType .
//
// When created, the default route defaults to an active state so state is not a
// required input. However, like all other state values the state of the default
// route can be updated after creation, but only when all other routes are also
// inactive. Conversely, no route can be active without the default route also
// being active.
//
// When you create a route, Refactor Spaces configures the Amazon API Gateway to
// send traffic to the target service as follows:
//
// - URL Endpoints
//
// # If the service has a URL endpoint, and the endpoint resolves to a private IP
//
// address, Refactor Spaces routes traffic using the API Gateway VPC link. If a
// service endpoint resolves to a public IP address, Refactor Spaces routes traffic
// over the public internet. Services can have HTTP or HTTPS URL endpoints. For
// HTTPS URLs, publicly-signed certificates are supported. Private Certificate
// Authorities (CAs) are permitted only if the CA's domain is also publicly
// resolvable.
//
// Refactor Spaces automatically resolves the public Domain Name System (DNS)
//
// names that are set in CreateService:UrlEndpoint when you create a service.
// The DNS names resolve when the DNS time-to-live (TTL) expires, or every 60
// seconds for TTLs less than 60 seconds. This periodic DNS resolution ensures that
// the route configuration remains up-to-date.
//
// # One-time health check
//
// # A one-time health check is performed on the service when either the route is
//
// updated from inactive to active, or when it is created with an active state. If
// the health check fails, the route transitions the route state to FAILED , an
// error code of SERVICE_ENDPOINT_HEALTH_CHECK_FAILURE is provided, and no
// traffic is sent to the service.
//
// # For private URLs, a target group is created on the Network Load Balancer and
//
// the load balancer target group runs default target health checks. By default,
// the health check is run against the service endpoint URL. Optionally, the health
// check can be performed against a different protocol, port, and/or path using the
// [CreateService:UrlEndpoint]parameter. All other health check settings for the load balancer use the
// default values described in the [Health checks for your target groups]in the Elastic Load Balancing guide. The
// health check is considered successful if at least one target within the target
// group transitions to a healthy state.
//
// - Lambda function endpoints
//
// # If the service has an Lambda function endpoint, then Refactor Spaces configures
//
// the Lambda function's resource policy to allow the application's API Gateway to
// invoke the function.
//
// The Lambda function state is checked. If the function is not active, the
//
// function configuration is updated so that Lambda resources are provisioned. If
// the Lambda state is Failed , then the route creation fails. For more
// information, see the [GetFunctionConfiguration's State response parameter]in the Lambda Developer Guide.
//
// # A check is performed to determine that a Lambda function with the specified ARN
//
// exists. If it does not exist, the health check fails. For public URLs, a
// connection is opened to the public endpoint. If the URL is not reachable, the
// health check fails.
//
// # Environments without a network bridge
//
// When you create environments without a network bridge ([CreateEnvironment:NetworkFabricType] is NONE) and you use
// your own networking infrastructure, you need to configure [VPC to VPC connectivity]between your network
// and the application proxy VPC. Route creation from the application proxy to
// service endpoints will fail if your network is not configured to connect to the
// application proxy VPC. For more information, see [Create a route]in the Refactor Spaces User
// Guide.
//
// [VPC to VPC connectivity]: https://docs.aws.amazon.com/whitepapers/latest/aws-vpc-connectivity-options/amazon-vpc-to-amazon-vpc-connectivity-options.html
// [Create a route]: https://docs.aws.amazon.com/migrationhub-refactor-spaces/latest/userguide/getting-started-create-role.html
// [CreateEnvironment:NetworkFabricType]: https://docs.aws.amazon.com/migrationhub-refactor-spaces/latest/APIReference/API_CreateEnvironment.html#migrationhubrefactorspaces-CreateEnvironment-request-NetworkFabricType
// [CreateService:UrlEndpoint]: https://docs.aws.amazon.com/migrationhub-refactor-spaces/latest/APIReference/API_CreateService.html#migrationhubrefactorspaces-CreateService-request-UrlEndpoint
// [Health checks for your target groups]: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/target-group-health-checks.html
// [GetFunctionConfiguration's State response parameter]: https://docs.aws.amazon.com/lambda/latest/dg/API_GetFunctionConfiguration.html#SSS-GetFunctionConfiguration-response-State
func migrationhubrefactorspaces_CreateRoute(cfg aws.Config, client *migrationhubrefactorspaces.Client) {
	input := &migrationhubrefactorspaces.CreateRouteInput{
		// ApplicationIdentifier: *string, // Required
		// EnvironmentIdentifier: *string, // Required
		// RouteType: types.RouteType, // Required
		// ServiceIdentifier: *string, // Required
	}

	if len(_migrationhubrefactorspacesApplicationIdentifier) > 0 {
		input.ApplicationIdentifier = aws.String(_migrationhubrefactorspacesApplicationIdentifier)
	}
	if len(_migrationhubrefactorspacesEnvironmentIdentifier) > 0 {
		input.EnvironmentIdentifier = aws.String(_migrationhubrefactorspacesEnvironmentIdentifier)
	}
	if len(_migrationhubrefactorspacesRouteType) > 0 {
		if err := assignInputField(input, "RouteType", _migrationhubrefactorspacesRouteType); err != nil {
			log.Errorf("invalid --route-type: %s", err.Error())
			return
		}
	}
	if len(_migrationhubrefactorspacesServiceIdentifier) > 0 {
		input.ServiceIdentifier = aws.String(_migrationhubrefactorspacesServiceIdentifier)
	}
	if len(_migrationhubrefactorspacesClientToken) > 0 {
		input.ClientToken = aws.String(_migrationhubrefactorspacesClientToken)
	}
	if len(_migrationhubrefactorspacesDefaultRoute) > 0 {
		if err := assignInputField(input, "DefaultRoute", _migrationhubrefactorspacesDefaultRoute); err != nil {
			log.Errorf("invalid --default-route: %s", err.Error())
			return
		}
	}
	if len(_migrationhubrefactorspacesTags) > 0 {
		if err := assignInputField(input, "Tags", _migrationhubrefactorspacesTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_migrationhubrefactorspacesUriPathRoute) > 0 {
		if err := assignInputField(input, "UriPathRoute", _migrationhubrefactorspacesUriPathRoute); err != nil {
			log.Errorf("invalid --uri-path-route: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRoute(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon Web Services Migration Hub Refactor Spaces service. The
// account owner of the service is always the environment owner, regardless of
// which account in the environment creates the service. Services have either a URL
// endpoint in a virtual private cloud (VPC), or a Lambda function endpoint.
//
// If an Amazon Web Services resource is launched in a service VPC, and you want
// it to be accessible to all of an environment’s services with VPCs and routes,
// apply the RefactorSpacesSecurityGroup to the resource. Alternatively, to add
// more cross-account constraints, apply your own security group.
func migrationhubrefactorspaces_CreateService(cfg aws.Config, client *migrationhubrefactorspaces.Client) {
	input := &migrationhubrefactorspaces.CreateServiceInput{
		// ApplicationIdentifier: *string, // Required
		// EndpointType: types.ServiceEndpointType, // Required
		// EnvironmentIdentifier: *string, // Required
		// Name: *string, // Required
	}

	if len(_migrationhubrefactorspacesApplicationIdentifier) > 0 {
		input.ApplicationIdentifier = aws.String(_migrationhubrefactorspacesApplicationIdentifier)
	}
	if len(_migrationhubrefactorspacesEndpointType) > 0 {
		if err := assignInputField(input, "EndpointType", _migrationhubrefactorspacesEndpointType); err != nil {
			log.Errorf("invalid --endpoint-type: %s", err.Error())
			return
		}
	}
	if len(_migrationhubrefactorspacesEnvironmentIdentifier) > 0 {
		input.EnvironmentIdentifier = aws.String(_migrationhubrefactorspacesEnvironmentIdentifier)
	}
	if len(_migrationhubrefactorspacesName) > 0 {
		input.Name = aws.String(_migrationhubrefactorspacesName)
	}
	if len(_migrationhubrefactorspacesClientToken) > 0 {
		input.ClientToken = aws.String(_migrationhubrefactorspacesClientToken)
	}
	if len(_migrationhubrefactorspacesDescription) > 0 {
		input.Description = aws.String(_migrationhubrefactorspacesDescription)
	}
	if len(_migrationhubrefactorspacesLambdaEndpoint) > 0 {
		if err := assignInputField(input, "LambdaEndpoint", _migrationhubrefactorspacesLambdaEndpoint); err != nil {
			log.Errorf("invalid --lambda-endpoint: %s", err.Error())
			return
		}
	}
	if len(_migrationhubrefactorspacesTags) > 0 {
		if err := assignInputField(input, "Tags", _migrationhubrefactorspacesTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_migrationhubrefactorspacesUrlEndpoint) > 0 {
		if err := assignInputField(input, "UrlEndpoint", _migrationhubrefactorspacesUrlEndpoint); err != nil {
			log.Errorf("invalid --url-endpoint: %s", err.Error())
			return
		}
	}
	if len(_migrationhubrefactorspacesVpcId) > 0 {
		input.VpcId = aws.String(_migrationhubrefactorspacesVpcId)
	}

	if resp, err := client.CreateService(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon Web Services Migration Hub Refactor Spaces application.
// Before you can delete an application, you must first delete any services or
// routes within the application.
func migrationhubrefactorspaces_DeleteApplication(cfg aws.Config, client *migrationhubrefactorspaces.Client) {
	input := &migrationhubrefactorspaces.DeleteApplicationInput{
		// ApplicationIdentifier: *string, // Required
		// EnvironmentIdentifier: *string, // Required
	}

	if len(_migrationhubrefactorspacesApplicationIdentifier) > 0 {
		input.ApplicationIdentifier = aws.String(_migrationhubrefactorspacesApplicationIdentifier)
	}
	if len(_migrationhubrefactorspacesEnvironmentIdentifier) > 0 {
		input.EnvironmentIdentifier = aws.String(_migrationhubrefactorspacesEnvironmentIdentifier)
	}

	if resp, err := client.DeleteApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon Web Services Migration Hub Refactor Spaces environment.
// Before you can delete an environment, you must first delete any applications and
// services within the environment.
func migrationhubrefactorspaces_DeleteEnvironment(cfg aws.Config, client *migrationhubrefactorspaces.Client) {
	input := &migrationhubrefactorspaces.DeleteEnvironmentInput{
		// EnvironmentIdentifier: *string, // Required
	}

	if len(_migrationhubrefactorspacesEnvironmentIdentifier) > 0 {
		input.EnvironmentIdentifier = aws.String(_migrationhubrefactorspacesEnvironmentIdentifier)
	}

	if resp, err := client.DeleteEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the resource policy set for the environment.
func migrationhubrefactorspaces_DeleteResourcePolicy(cfg aws.Config, client *migrationhubrefactorspaces.Client) {
	input := &migrationhubrefactorspaces.DeleteResourcePolicyInput{
		// Identifier: *string, // Required
	}

	if len(_migrationhubrefactorspacesIdentifier) > 0 {
		input.Identifier = aws.String(_migrationhubrefactorspacesIdentifier)
	}

	if resp, err := client.DeleteResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon Web Services Migration Hub Refactor Spaces route.
func migrationhubrefactorspaces_DeleteRoute(cfg aws.Config, client *migrationhubrefactorspaces.Client) {
	input := &migrationhubrefactorspaces.DeleteRouteInput{
		// ApplicationIdentifier: *string, // Required
		// EnvironmentIdentifier: *string, // Required
		// RouteIdentifier: *string, // Required
	}

	if len(_migrationhubrefactorspacesApplicationIdentifier) > 0 {
		input.ApplicationIdentifier = aws.String(_migrationhubrefactorspacesApplicationIdentifier)
	}
	if len(_migrationhubrefactorspacesEnvironmentIdentifier) > 0 {
		input.EnvironmentIdentifier = aws.String(_migrationhubrefactorspacesEnvironmentIdentifier)
	}
	if len(_migrationhubrefactorspacesRouteIdentifier) > 0 {
		input.RouteIdentifier = aws.String(_migrationhubrefactorspacesRouteIdentifier)
	}

	if resp, err := client.DeleteRoute(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon Web Services Migration Hub Refactor Spaces service.
func migrationhubrefactorspaces_DeleteService(cfg aws.Config, client *migrationhubrefactorspaces.Client) {
	input := &migrationhubrefactorspaces.DeleteServiceInput{
		// ApplicationIdentifier: *string, // Required
		// EnvironmentIdentifier: *string, // Required
		// ServiceIdentifier: *string, // Required
	}

	if len(_migrationhubrefactorspacesApplicationIdentifier) > 0 {
		input.ApplicationIdentifier = aws.String(_migrationhubrefactorspacesApplicationIdentifier)
	}
	if len(_migrationhubrefactorspacesEnvironmentIdentifier) > 0 {
		input.EnvironmentIdentifier = aws.String(_migrationhubrefactorspacesEnvironmentIdentifier)
	}
	if len(_migrationhubrefactorspacesServiceIdentifier) > 0 {
		input.ServiceIdentifier = aws.String(_migrationhubrefactorspacesServiceIdentifier)
	}

	if resp, err := client.DeleteService(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets an Amazon Web Services Migration Hub Refactor Spaces application.
func migrationhubrefactorspaces_GetApplication(cfg aws.Config, client *migrationhubrefactorspaces.Client) {
	input := &migrationhubrefactorspaces.GetApplicationInput{
		// ApplicationIdentifier: *string, // Required
		// EnvironmentIdentifier: *string, // Required
	}

	if len(_migrationhubrefactorspacesApplicationIdentifier) > 0 {
		input.ApplicationIdentifier = aws.String(_migrationhubrefactorspacesApplicationIdentifier)
	}
	if len(_migrationhubrefactorspacesEnvironmentIdentifier) > 0 {
		input.EnvironmentIdentifier = aws.String(_migrationhubrefactorspacesEnvironmentIdentifier)
	}

	if resp, err := client.GetApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets an Amazon Web Services Migration Hub Refactor Spaces environment.
func migrationhubrefactorspaces_GetEnvironment(cfg aws.Config, client *migrationhubrefactorspaces.Client) {
	input := &migrationhubrefactorspaces.GetEnvironmentInput{
		// EnvironmentIdentifier: *string, // Required
	}

	if len(_migrationhubrefactorspacesEnvironmentIdentifier) > 0 {
		input.EnvironmentIdentifier = aws.String(_migrationhubrefactorspacesEnvironmentIdentifier)
	}

	if resp, err := client.GetEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the resource-based permission policy that is set for the given
// environment.
func migrationhubrefactorspaces_GetResourcePolicy(cfg aws.Config, client *migrationhubrefactorspaces.Client) {
	input := &migrationhubrefactorspaces.GetResourcePolicyInput{
		// Identifier: *string, // Required
	}

	if len(_migrationhubrefactorspacesIdentifier) > 0 {
		input.Identifier = aws.String(_migrationhubrefactorspacesIdentifier)
	}

	if resp, err := client.GetResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets an Amazon Web Services Migration Hub Refactor Spaces route.
func migrationhubrefactorspaces_GetRoute(cfg aws.Config, client *migrationhubrefactorspaces.Client) {
	input := &migrationhubrefactorspaces.GetRouteInput{
		// ApplicationIdentifier: *string, // Required
		// EnvironmentIdentifier: *string, // Required
		// RouteIdentifier: *string, // Required
	}

	if len(_migrationhubrefactorspacesApplicationIdentifier) > 0 {
		input.ApplicationIdentifier = aws.String(_migrationhubrefactorspacesApplicationIdentifier)
	}
	if len(_migrationhubrefactorspacesEnvironmentIdentifier) > 0 {
		input.EnvironmentIdentifier = aws.String(_migrationhubrefactorspacesEnvironmentIdentifier)
	}
	if len(_migrationhubrefactorspacesRouteIdentifier) > 0 {
		input.RouteIdentifier = aws.String(_migrationhubrefactorspacesRouteIdentifier)
	}

	if resp, err := client.GetRoute(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets an Amazon Web Services Migration Hub Refactor Spaces service.
func migrationhubrefactorspaces_GetService(cfg aws.Config, client *migrationhubrefactorspaces.Client) {
	input := &migrationhubrefactorspaces.GetServiceInput{
		// ApplicationIdentifier: *string, // Required
		// EnvironmentIdentifier: *string, // Required
		// ServiceIdentifier: *string, // Required
	}

	if len(_migrationhubrefactorspacesApplicationIdentifier) > 0 {
		input.ApplicationIdentifier = aws.String(_migrationhubrefactorspacesApplicationIdentifier)
	}
	if len(_migrationhubrefactorspacesEnvironmentIdentifier) > 0 {
		input.EnvironmentIdentifier = aws.String(_migrationhubrefactorspacesEnvironmentIdentifier)
	}
	if len(_migrationhubrefactorspacesServiceIdentifier) > 0 {
		input.ServiceIdentifier = aws.String(_migrationhubrefactorspacesServiceIdentifier)
	}

	if resp, err := client.GetService(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all the Amazon Web Services Migration Hub Refactor Spaces applications
// within an environment.
func migrationhubrefactorspaces_ListApplications(cfg aws.Config, client *migrationhubrefactorspaces.Client) {
	input := &migrationhubrefactorspaces.ListApplicationsInput{
		// EnvironmentIdentifier: *string, // Required
	}

	if len(_migrationhubrefactorspacesEnvironmentIdentifier) > 0 {
		input.EnvironmentIdentifier = aws.String(_migrationhubrefactorspacesEnvironmentIdentifier)
	}
	if len(_migrationhubrefactorspacesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _migrationhubrefactorspacesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_migrationhubrefactorspacesNextToken) > 0 {
		input.NextToken = aws.String(_migrationhubrefactorspacesNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListApplications(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*migrationhubrefactorspaces.ListApplicationsOutput
	p := migrationhubrefactorspaces.NewListApplicationsPaginator(client, input)
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

// Lists all Amazon Web Services Migration Hub Refactor Spaces service virtual
// private clouds (VPCs) that are part of the environment.
func migrationhubrefactorspaces_ListEnvironmentVpcs(cfg aws.Config, client *migrationhubrefactorspaces.Client) {
	input := &migrationhubrefactorspaces.ListEnvironmentVpcsInput{
		// EnvironmentIdentifier: *string, // Required
	}

	if len(_migrationhubrefactorspacesEnvironmentIdentifier) > 0 {
		input.EnvironmentIdentifier = aws.String(_migrationhubrefactorspacesEnvironmentIdentifier)
	}
	if len(_migrationhubrefactorspacesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _migrationhubrefactorspacesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_migrationhubrefactorspacesNextToken) > 0 {
		input.NextToken = aws.String(_migrationhubrefactorspacesNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEnvironmentVpcs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*migrationhubrefactorspaces.ListEnvironmentVpcsOutput
	p := migrationhubrefactorspaces.NewListEnvironmentVpcsPaginator(client, input)
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

// Lists Amazon Web Services Migration Hub Refactor Spaces environments owned by a
// caller account or shared with the caller account.
func migrationhubrefactorspaces_ListEnvironments(cfg aws.Config, client *migrationhubrefactorspaces.Client) {
	input := &migrationhubrefactorspaces.ListEnvironmentsInput{}

	if len(_migrationhubrefactorspacesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _migrationhubrefactorspacesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_migrationhubrefactorspacesNextToken) > 0 {
		input.NextToken = aws.String(_migrationhubrefactorspacesNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEnvironments(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*migrationhubrefactorspaces.ListEnvironmentsOutput
	p := migrationhubrefactorspaces.NewListEnvironmentsPaginator(client, input)
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

// Lists all the Amazon Web Services Migration Hub Refactor Spaces routes within
// an application.
func migrationhubrefactorspaces_ListRoutes(cfg aws.Config, client *migrationhubrefactorspaces.Client) {
	input := &migrationhubrefactorspaces.ListRoutesInput{
		// ApplicationIdentifier: *string, // Required
		// EnvironmentIdentifier: *string, // Required
	}

	if len(_migrationhubrefactorspacesApplicationIdentifier) > 0 {
		input.ApplicationIdentifier = aws.String(_migrationhubrefactorspacesApplicationIdentifier)
	}
	if len(_migrationhubrefactorspacesEnvironmentIdentifier) > 0 {
		input.EnvironmentIdentifier = aws.String(_migrationhubrefactorspacesEnvironmentIdentifier)
	}
	if len(_migrationhubrefactorspacesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _migrationhubrefactorspacesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_migrationhubrefactorspacesNextToken) > 0 {
		input.NextToken = aws.String(_migrationhubrefactorspacesNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRoutes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*migrationhubrefactorspaces.ListRoutesOutput
	p := migrationhubrefactorspaces.NewListRoutesPaginator(client, input)
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

// Lists all the Amazon Web Services Migration Hub Refactor Spaces services within
// an application.
func migrationhubrefactorspaces_ListServices(cfg aws.Config, client *migrationhubrefactorspaces.Client) {
	input := &migrationhubrefactorspaces.ListServicesInput{
		// ApplicationIdentifier: *string, // Required
		// EnvironmentIdentifier: *string, // Required
	}

	if len(_migrationhubrefactorspacesApplicationIdentifier) > 0 {
		input.ApplicationIdentifier = aws.String(_migrationhubrefactorspacesApplicationIdentifier)
	}
	if len(_migrationhubrefactorspacesEnvironmentIdentifier) > 0 {
		input.EnvironmentIdentifier = aws.String(_migrationhubrefactorspacesEnvironmentIdentifier)
	}
	if len(_migrationhubrefactorspacesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _migrationhubrefactorspacesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_migrationhubrefactorspacesNextToken) > 0 {
		input.NextToken = aws.String(_migrationhubrefactorspacesNextToken)
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

	var results []*migrationhubrefactorspaces.ListServicesOutput
	p := migrationhubrefactorspaces.NewListServicesPaginator(client, input)
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

// Lists the tags of a resource. The caller account must be the same as the
// resource’s OwnerAccountId . Listing tags in other accounts is not supported.
func migrationhubrefactorspaces_ListTagsForResource(cfg aws.Config, client *migrationhubrefactorspaces.Client) {
	input := &migrationhubrefactorspaces.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_migrationhubrefactorspacesResourceArn) > 0 {
		input.ResourceArn = aws.String(_migrationhubrefactorspacesResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attaches a resource-based permission policy to the Amazon Web Services
// Migration Hub Refactor Spaces environment. The policy must contain the same
// actions and condition statements as the
// arn:aws:ram::aws:permission/AWSRAMDefaultPermissionRefactorSpacesEnvironment
// permission in Resource Access Manager. The policy must not contain new lines or
// blank lines.
func migrationhubrefactorspaces_PutResourcePolicy(cfg aws.Config, client *migrationhubrefactorspaces.Client) {
	input := &migrationhubrefactorspaces.PutResourcePolicyInput{
		// Policy: *string, // Required
		// ResourceArn: *string, // Required
	}

	if len(_migrationhubrefactorspacesPolicy) > 0 {
		input.Policy = aws.String(_migrationhubrefactorspacesPolicy)
	}
	if len(_migrationhubrefactorspacesResourceArn) > 0 {
		input.ResourceArn = aws.String(_migrationhubrefactorspacesResourceArn)
	}

	if resp, err := client.PutResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the tags of a given resource. Tags are metadata which can be used to
// manage a resource. To tag a resource, the caller account must be the same as the
// resource’s OwnerAccountId . Tagging resources in other accounts is not supported.
//
// Amazon Web Services Migration Hub Refactor Spaces does not propagate tags to
// orchestrated resources, such as an environment’s transit gateway.
func migrationhubrefactorspaces_TagResource(cfg aws.Config, client *migrationhubrefactorspaces.Client) {
	input := &migrationhubrefactorspaces.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_migrationhubrefactorspacesResourceArn) > 0 {
		input.ResourceArn = aws.String(_migrationhubrefactorspacesResourceArn)
	}
	if len(_migrationhubrefactorspacesTags) > 0 {
		if err := assignInputField(input, "Tags", _migrationhubrefactorspacesTags); err != nil {
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

// Adds to or modifies the tags of the given resource. Tags are metadata which can
// be used to manage a resource. To untag a resource, the caller account must be
// the same as the resource’s OwnerAccountId . Untagging resources across accounts
// is not supported.
func migrationhubrefactorspaces_UntagResource(cfg aws.Config, client *migrationhubrefactorspaces.Client) {
	input := &migrationhubrefactorspaces.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_migrationhubrefactorspacesResourceArn) > 0 {
		input.ResourceArn = aws.String(_migrationhubrefactorspacesResourceArn)
	}
	if len(_migrationhubrefactorspacesTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _migrationhubrefactorspacesTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an Amazon Web Services Migration Hub Refactor Spaces route.
func migrationhubrefactorspaces_UpdateRoute(cfg aws.Config, client *migrationhubrefactorspaces.Client) {
	input := &migrationhubrefactorspaces.UpdateRouteInput{
		// ActivationState: types.RouteActivationState, // Required
		// ApplicationIdentifier: *string, // Required
		// EnvironmentIdentifier: *string, // Required
		// RouteIdentifier: *string, // Required
	}

	if len(_migrationhubrefactorspacesActivationState) > 0 {
		if err := assignInputField(input, "ActivationState", _migrationhubrefactorspacesActivationState); err != nil {
			log.Errorf("invalid --activation-state: %s", err.Error())
			return
		}
	}
	if len(_migrationhubrefactorspacesApplicationIdentifier) > 0 {
		input.ApplicationIdentifier = aws.String(_migrationhubrefactorspacesApplicationIdentifier)
	}
	if len(_migrationhubrefactorspacesEnvironmentIdentifier) > 0 {
		input.EnvironmentIdentifier = aws.String(_migrationhubrefactorspacesEnvironmentIdentifier)
	}
	if len(_migrationhubrefactorspacesRouteIdentifier) > 0 {
		input.RouteIdentifier = aws.String(_migrationhubrefactorspacesRouteIdentifier)
	}

	if resp, err := client.UpdateRoute(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_migrationhubrefactorspacesCmd)
	_migrationhubrefactorspacesCmd.Flags().SortFlags = false

	_migrationhubrefactorspacesCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_migrationhubrefactorspacesCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_migrationhubrefactorspacesCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_migrationhubrefactorspacesCmd.Flags().StringVarP(&_migrationhubrefactorspacesActivationState, "activation-state", "", "", "Activation State")
	_migrationhubrefactorspacesCmd.Flags().StringVarP(&_migrationhubrefactorspacesApiGatewayProxy, "api-gateway-proxy", "", "", "API Gateway Proxy")
	_migrationhubrefactorspacesCmd.Flags().StringVarP(&_migrationhubrefactorspacesApplicationIdentifier, "application-identifier", "", "", "Application Identifier")
	_migrationhubrefactorspacesCmd.Flags().StringVarP(&_migrationhubrefactorspacesClientToken, "client-token", "", "", "Client Token")
	_migrationhubrefactorspacesCmd.Flags().StringVarP(&_migrationhubrefactorspacesDefaultRoute, "default-route", "", "", "Default Route")
	_migrationhubrefactorspacesCmd.Flags().StringVarP(&_migrationhubrefactorspacesDescription, "description", "", "", "Description")
	_migrationhubrefactorspacesCmd.Flags().StringVarP(&_migrationhubrefactorspacesEndpointType, "endpoint-type", "", "", "Endpoint Type")
	_migrationhubrefactorspacesCmd.Flags().StringVarP(&_migrationhubrefactorspacesEnvironmentIdentifier, "environment-identifier", "", "", "Environment Identifier")
	_migrationhubrefactorspacesCmd.Flags().StringVarP(&_migrationhubrefactorspacesIdentifier, "identifier", "", "", "Identifier")
	_migrationhubrefactorspacesCmd.Flags().StringVarP(&_migrationhubrefactorspacesLambdaEndpoint, "lambda-endpoint", "", "", "Lambda Endpoint")
	_migrationhubrefactorspacesCmd.Flags().StringVarP(&_migrationhubrefactorspacesMaxResults, "max-results", "", "", "Max Results")
	_migrationhubrefactorspacesCmd.Flags().StringVarP(&_migrationhubrefactorspacesName, "name", "", "", "Name")
	_migrationhubrefactorspacesCmd.Flags().StringVarP(&_migrationhubrefactorspacesNetworkFabricType, "network-fabric-type", "", "", "Network Fabric Type")
	_migrationhubrefactorspacesCmd.Flags().StringVarP(&_migrationhubrefactorspacesNextToken, "next-token", "", "", "Next Token")
	_migrationhubrefactorspacesCmd.Flags().StringVarP(&_migrationhubrefactorspacesPolicy, "policy", "", "", "Policy")
	_migrationhubrefactorspacesCmd.Flags().StringVarP(&_migrationhubrefactorspacesProxyType, "proxy-type", "", "", "Proxy Type")
	_migrationhubrefactorspacesCmd.Flags().StringVarP(&_migrationhubrefactorspacesResourceArn, "resource-arn", "", "", "Resource ARN")
	_migrationhubrefactorspacesCmd.Flags().StringVarP(&_migrationhubrefactorspacesRouteIdentifier, "route-identifier", "", "", "Route Identifier")
	_migrationhubrefactorspacesCmd.Flags().StringVarP(&_migrationhubrefactorspacesRouteType, "route-type", "", "", "Route Type")
	_migrationhubrefactorspacesCmd.Flags().StringVarP(&_migrationhubrefactorspacesServiceIdentifier, "service-identifier", "", "", "Service Identifier")
	_migrationhubrefactorspacesCmd.Flags().StringSliceVarP(&_migrationhubrefactorspacesTagKeys, "tag-keys", "", nil, "Tag Keys")
	_migrationhubrefactorspacesCmd.Flags().StringVarP(&_migrationhubrefactorspacesTags, "tags", "", "", "Tags")
	_migrationhubrefactorspacesCmd.Flags().StringVarP(&_migrationhubrefactorspacesUriPathRoute, "uri-path-route", "", "", "URI Path Route")
	_migrationhubrefactorspacesCmd.Flags().StringVarP(&_migrationhubrefactorspacesUrlEndpoint, "url-endpoint", "", "", "URL Endpoint")
	_migrationhubrefactorspacesCmd.Flags().StringVarP(&_migrationhubrefactorspacesVpcId, "vpc-id", "", "", "VPC ID")

	_migrationhubrefactorspacesCmd.Flags().BoolVarP(&_migrationhubrefactorspacesCreateApplication, "create-application", "", false, "Create Application")
	_migrationhubrefactorspacesCmd.Flags().BoolVarP(&_migrationhubrefactorspacesCreateEnvironment, "create-environment", "", false, "Create Environment")
	_migrationhubrefactorspacesCmd.Flags().BoolVarP(&_migrationhubrefactorspacesCreateRoute, "create-route", "", false, "Create Route")
	_migrationhubrefactorspacesCmd.Flags().BoolVarP(&_migrationhubrefactorspacesCreateService, "create-service", "", false, "Create Service")
	_migrationhubrefactorspacesCmd.Flags().BoolVarP(&_migrationhubrefactorspacesDeleteApplication, "delete-application", "", false, "Delete Application")
	_migrationhubrefactorspacesCmd.Flags().BoolVarP(&_migrationhubrefactorspacesDeleteEnvironment, "delete-environment", "", false, "Delete Environment")
	_migrationhubrefactorspacesCmd.Flags().BoolVarP(&_migrationhubrefactorspacesDeleteResourcePolicy, "delete-resource-policy", "", false, "Delete Resource Policy")
	_migrationhubrefactorspacesCmd.Flags().BoolVarP(&_migrationhubrefactorspacesDeleteRoute, "delete-route", "", false, "Delete Route")
	_migrationhubrefactorspacesCmd.Flags().BoolVarP(&_migrationhubrefactorspacesDeleteService, "delete-service", "", false, "Delete Service")
	_migrationhubrefactorspacesCmd.Flags().BoolVarP(&_migrationhubrefactorspacesGetApplication, "get-application", "", false, "Get Application")
	_migrationhubrefactorspacesCmd.Flags().BoolVarP(&_migrationhubrefactorspacesGetEnvironment, "get-environment", "", false, "Get Environment")
	_migrationhubrefactorspacesCmd.Flags().BoolVarP(&_migrationhubrefactorspacesGetResourcePolicy, "get-resource-policy", "", false, "Get Resource Policy")
	_migrationhubrefactorspacesCmd.Flags().BoolVarP(&_migrationhubrefactorspacesGetRoute, "get-route", "", false, "Get Route")
	_migrationhubrefactorspacesCmd.Flags().BoolVarP(&_migrationhubrefactorspacesGetService, "get-service", "", false, "Get Service")
	_migrationhubrefactorspacesCmd.Flags().BoolVarP(&_migrationhubrefactorspacesListApplications, "list-applications", "", false, "List Applications")
	_migrationhubrefactorspacesCmd.Flags().BoolVarP(&_migrationhubrefactorspacesListEnvironmentVpcs, "list-environment-vpcs", "", false, "List Environment Vpcs")
	_migrationhubrefactorspacesCmd.Flags().BoolVarP(&_migrationhubrefactorspacesListEnvironments, "list-environments", "", false, "List Environments")
	_migrationhubrefactorspacesCmd.Flags().BoolVarP(&_migrationhubrefactorspacesListRoutes, "list-routes", "", false, "List Routes")
	_migrationhubrefactorspacesCmd.Flags().BoolVarP(&_migrationhubrefactorspacesListServices, "list-services", "", false, "List Services")
	_migrationhubrefactorspacesCmd.Flags().BoolVarP(&_migrationhubrefactorspacesListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_migrationhubrefactorspacesCmd.Flags().BoolVarP(&_migrationhubrefactorspacesPutResourcePolicy, "put-resource-policy", "", false, "Put Resource Policy")
	_migrationhubrefactorspacesCmd.Flags().BoolVarP(&_migrationhubrefactorspacesTagResource, "tag-resource", "", false, "Tag Resource")
	_migrationhubrefactorspacesCmd.Flags().BoolVarP(&_migrationhubrefactorspacesUntagResource, "untag-resource", "", false, "Untag Resource")
	_migrationhubrefactorspacesCmd.Flags().BoolVarP(&_migrationhubrefactorspacesUpdateRoute, "update-route", "", false, "Update Route")

}
