package iam

// GetOrganizationsAccessReport is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Retrieves the service last accessed data report for Organizations that was
// previously generated using the [GenerateOrganizationsAccessReport]operation. This operation retrieves the status
// of your report job and the report contents.
//
// Depending on the parameters that you passed when you generated the report, the
// data returned could include different information. For details, see [GenerateOrganizationsAccessReport].
//
// To call this operation, you must be signed in to the management account in your
// organization. SCPs must be enabled for your organization root. You must have
// permissions to perform this operation. For more information, see [Refining permissions using service last accessed data]in the IAM
// User Guide.
//
// For each service that principals in an account (root user, IAM users, or IAM
// roles) could access using SCPs, the operation returns details about the most
// recent access attempt. If there was no attempt, the service is listed without
// details about the most recent attempt to access the service. If the operation
// fails, it returns the reason that it failed.
//
// By default, the list is sorted by service namespace.
//
// [GenerateOrganizationsAccessReport]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GenerateOrganizationsAccessReport.html
// [Refining permissions using service last accessed data]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_access-advisor.html
