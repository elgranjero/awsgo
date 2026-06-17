package ec2

// ModifyVpcEndpointServicePermissions is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Modifies the permissions for your VPC endpoint service. You can add or remove
// permissions for service consumers (Amazon Web Services accounts, users, and IAM
// roles) to connect to your endpoint service. Principal ARNs with path components
// aren't supported.
//
// If you grant permissions to all principals, the service is public. Any users
// who know the name of a public service can send a request to attach an endpoint.
// If the service does not require manual approval, attachments are automatically
// approved.
