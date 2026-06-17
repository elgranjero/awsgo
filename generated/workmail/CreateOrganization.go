package workmail

// CreateOrganization is generated as a reference stub.
// Executable command wiring lives under cmd/workmail.go.
//
// Creates a new WorkMail organization. Optionally, you can choose to associate an
// existing AWS Directory Service directory with your organization. If an AWS
// Directory Service directory ID is specified, the organization alias must match
// the directory alias. If you choose not to associate an existing directory with
// your organization, then we create a new WorkMail directory for you. For more
// information, see [Adding an organization]in the WorkMail Administrator Guide.
//
// You can associate multiple email domains with an organization, then choose your
// default email domain from the WorkMail console. You can also associate a domain
// that is managed in an Amazon Route 53 public hosted zone. For more information,
// see [Adding a domain]and [Choosing the default domain] in the WorkMail Administrator Guide.
//
// Optionally, you can use a customer managed key from AWS Key Management Service
// (AWS KMS) to encrypt email for your organization. If you don't associate an AWS
// KMS key, WorkMail creates a default, AWS managed key for you.
//
// [Adding an organization]: https://docs.aws.amazon.com/workmail/latest/adminguide/add_new_organization.html
// [Adding a domain]: https://docs.aws.amazon.com/workmail/latest/adminguide/add_domain.html
// [Choosing the default domain]: https://docs.aws.amazon.com/workmail/latest/adminguide/default_domain.html
