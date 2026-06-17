package cloudwatchlogs

// PutResourcePolicy is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatchlogs.go.
//
// Creates or updates a resource policy allowing other Amazon Web Services
// services to put log events to this account, such as Amazon Route 53. This API
// has the following restrictions:
//
// - Supported actions - Policy only supports logs:PutLogEvents and
// logs:CreateLogStream actions
//
// - Supported principals - Policy only applies when operations are invoked by
// Amazon Web Services service principals (not IAM users, roles, or cross-account
// principals
//
// - Policy limits - An account can have a maximum of 10 policies without
// resourceARN and one per LogGroup resourceARN
//
// Resource policies with actions invoked by non-Amazon Web Services service
// principals (such as IAM users, roles, or other Amazon Web Services accounts)
// will not be enforced. For access control involving these principals, use the IAM
// policies.
