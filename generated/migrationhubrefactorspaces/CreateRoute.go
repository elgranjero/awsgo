package migrationhubrefactorspaces

// CreateRoute is generated as a reference stub.
// Executable command wiring lives under cmd/migrationhubrefactorspaces.go.
//
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
// If the service has a URL endpoint, and the endpoint resolves to a private IP
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
// A one-time health check is performed on the service when either the route is
//
// updated from inactive to active, or when it is created with an active state. If
// the health check fails, the route transitions the route state to FAILED , an
// error code of SERVICE_ENDPOINT_HEALTH_CHECK_FAILURE is provided, and no
// traffic is sent to the service.
//
// For private URLs, a target group is created on the Network Load Balancer and
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
// If the service has an Lambda function endpoint, then Refactor Spaces configures
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
// A check is performed to determine that a Lambda function with the specified ARN
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
//
// [CreateService:UrlEndpoint]: https://docs.aws.amazon.com/migrationhub-refactor-spaces/latest/APIReference/API_CreateService.html#migrationhubrefactorspaces-CreateService-request-UrlEndpoint
// [Health checks for your target groups]: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/target-group-health-checks.html
// [GetFunctionConfiguration's State response parameter]: https://docs.aws.amazon.com/lambda/latest/dg/API_GetFunctionConfiguration.html#SSS-GetFunctionConfiguration-response-State
