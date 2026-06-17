package iam

// DeleteServiceLinkedRole is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Submits a service-linked role deletion request and returns a DeletionTaskId ,
// which you can use to check the status of the deletion. Before you call this
// operation, confirm that the role has no active sessions and that any resources
// used by the role in the linked service are deleted. If you call this operation
// more than once for the same service-linked role and an earlier deletion task is
// not complete, then the DeletionTaskId of the earlier request is returned.
//
// If you submit a deletion request for a service-linked role whose linked service
// is still accessing a resource, then the deletion task fails. If it fails, the [GetServiceLinkedRoleDeletionStatus]
// operation returns the reason for the failure, usually including the resources
// that must be deleted. To delete the service-linked role, you must first remove
// those resources from the linked service and then submit the deletion request
// again. Resources are specific to the service that is linked to the role. For
// more information about removing resources from a service, see the [Amazon Web Services documentation]for your
// service.
//
// For more information about service-linked roles, see [Roles terms and concepts: Amazon Web Services service-linked role] in the IAM User Guide.
//
// [Roles terms and concepts: Amazon Web Services service-linked role]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_terms-and-concepts.html#iam-term-service-linked-role
// [Amazon Web Services documentation]: http://docs.aws.amazon.com/
// [GetServiceLinkedRoleDeletionStatus]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetServiceLinkedRoleDeletionStatus.html
