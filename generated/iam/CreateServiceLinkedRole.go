package iam

// CreateServiceLinkedRole is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Creates an IAM role that is linked to a specific Amazon Web Services service.
// The service controls the attached policies and when the role can be deleted.
// This helps ensure that the service is not broken by an unexpectedly changed or
// deleted role, which could put your Amazon Web Services resources into an unknown
// state. Allowing the service to control the role helps improve service stability
// and proper cleanup when a service and its role are no longer needed. For more
// information, see [Using service-linked roles]in the IAM User Guide.
//
// To attach a policy to this service-linked role, you must make the request using
// the Amazon Web Services service that depends on this role.
//
// [Using service-linked roles]: https://docs.aws.amazon.com/IAM/latest/UserGuide/using-service-linked-roles.html
