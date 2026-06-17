package directoryservice

// AddIpRoutes is generated as a reference stub.
// Executable command wiring lives under cmd/directoryservice.go.
//
// If the DNS server for your self-managed domain uses a publicly addressable IP
// address, you must add a CIDR address block to correctly route traffic to and
// from your Microsoft AD on Amazon Web Services. AddIpRoutes adds this address
// block. You can also use AddIpRoutes to facilitate routing traffic that uses
// public IP ranges from your Microsoft AD on Amazon Web Services to a peer VPC.
//
// Before you call AddIpRoutes, ensure that all of the required permissions have
// been explicitly granted through a policy. For details about what permissions are
// required to run the AddIpRoutes operation, see [Directory Service API Permissions: Actions, Resources, and Conditions Reference].
//
// [Directory Service API Permissions: Actions, Resources, and Conditions Reference]: http://docs.aws.amazon.com/directoryservice/latest/admin-guide/UsingWithDS_IAM_ResourcePermissions.html
