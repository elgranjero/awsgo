package organizations

// EnableAWSServiceAccess is generated as a reference stub.
// Executable command wiring lives under cmd/organizations.go.
//
// Provides an Amazon Web Services service (the service that is specified by
// ServicePrincipal ) with permissions to view the structure of an organization,
// create a [service-linked role]in all the accounts in the organization, and allow the service to
// perform operations on behalf of the organization and its accounts. Establishing
// these permissions can be a first step in enabling the integration of an Amazon
// Web Services service with Organizations.
//
// We recommend that you enable integration between Organizations and the
// specified Amazon Web Services service by using the console or commands that are
// provided by the specified service. Doing so ensures that the service is aware
// that it can create the resources that are required for the integration. How the
// service creates those resources in the organization's accounts depends on that
// service. For more information, see the documentation for the other Amazon Web
// Services service.
//
// For more information about enabling services to integrate with Organizations,
// see [Using Organizations with other Amazon Web Services services]in the Organizations User Guide.
//
// You can only call this operation from the management account.
//
// [Using Organizations with other Amazon Web Services services]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services.html
// [service-linked role]: https://docs.aws.amazon.com/IAM/latest/UserGuide/using-service-linked-roles.html
