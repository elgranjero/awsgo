package migrationhubrefactorspaces

// CreateApplication is generated as a reference stub.
// Executable command wiring lives under cmd/migrationhubrefactorspaces.go.
//
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
