package ram

// EnableSharingWithAwsOrganization is generated as a reference stub.
// Executable command wiring lives under cmd/ram.go.
//
// Enables resource sharing within your organization in Organizations. This
// operation creates a service-linked role called
// AWSServiceRoleForResourceAccessManager that has the IAM managed policy named
// AWSResourceAccessManagerServiceRolePolicy attached. This role permits RAM to
// retrieve information about the organization and its structure. This lets you
// share resources with all of the accounts in the calling account's organization
// by specifying the organization ID, or all of the accounts in an organizational
// unit (OU) by specifying the OU ID. Until you enable sharing within the
// organization, you can specify only individual Amazon Web Services accounts, or
// for supported resource types, IAM roles and users.
//
// You must call this operation from an IAM role or user in the organization's
// management account.
