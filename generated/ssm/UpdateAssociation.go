package ssm

// UpdateAssociation is generated as a reference stub.
// Executable command wiring lives under cmd/ssm.go.
//
// Updates an association. You can update the association name and version, the
// document version, schedule, parameters, and Amazon Simple Storage Service
// (Amazon S3) output. When you call UpdateAssociation , the system removes all
// optional parameters from the request and overwrites the association with null
// values for those parameters. This is by design. You must specify all optional
// parameters in the call, even if you are not changing the parameters. This
// includes the Name parameter. Before calling this API action, we recommend that
// you call the DescribeAssociationAPI operation and make a note of all optional parameters required
// for your UpdateAssociation call.
//
// In order to call this API operation, a user, group, or role must be granted
// permission to call the DescribeAssociationAPI operation. If you don't have permission to call
// DescribeAssociation , then you receive the following error: An error occurred
// (AccessDeniedException) when calling the UpdateAssociation operation: User:
// isn't authorized to perform: ssm:DescribeAssociation on resource:
//
// When you update an association, the association immediately runs against the
// specified targets. You can add the ApplyOnlyAtCronInterval parameter to run the
// association during the next schedule run.
