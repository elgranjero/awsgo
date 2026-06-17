package health

// DisableHealthServiceAccessForOrganization is generated as a reference stub.
// Executable command wiring lives under cmd/health.go.
//
// Disables Health from working with Organizations. To call this operation, you
// must sign in to the organization's management account. For more information, see
// [Aggregating Health events]in the Health User Guide.
//
// This operation doesn't remove the service-linked role from the management
// account in your organization. You must use the IAM console, API, or Command Line
// Interface (CLI) to remove the service-linked role. For more information, see [Deleting a Service-Linked Role]in
// the IAM User Guide.
//
// You can also disable the organizational feature by using the Organizations [DisableAWSServiceAccess] API
// operation. After you call this operation, Health stops aggregating events for
// all other Amazon Web Services accounts in your organization. If you call the
// Health API operations for organizational view, Health returns an error. Health
// continues to aggregate health events for your Amazon Web Services account.
//
// [DisableAWSServiceAccess]: https://docs.aws.amazon.com/organizations/latest/APIReference/API_DisableAWSServiceAccess.html
// [Aggregating Health events]: https://docs.aws.amazon.com/health/latest/ug/aggregate-events.html
// [Deleting a Service-Linked Role]: https://docs.aws.amazon.com/IAM/latest/UserGuide/using-service-linked-roles.html#delete-service-linked-role
