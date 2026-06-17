package iam

// GetServiceLinkedRoleDeletionStatus is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Retrieves the status of your service-linked role deletion. After you use [DeleteServiceLinkedRole] to
// submit a service-linked role for deletion, you can use the DeletionTaskId
// parameter in GetServiceLinkedRoleDeletionStatus to check the status of the
// deletion. If the deletion fails, this operation returns the reason that it
// failed, if that information is returned by the service.
//
// [DeleteServiceLinkedRole]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_DeleteServiceLinkedRole.html
