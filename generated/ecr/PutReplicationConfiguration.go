package ecr

// PutReplicationConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/ecr.go.
//
// Creates or updates the replication configuration for a registry. The existing
// replication configuration for a repository can be retrieved with the DescribeRegistryAPI
// action. The first time the PutReplicationConfiguration API is called, a
// service-linked IAM role is created in your account for the replication process.
// For more information, see [Using service-linked roles for Amazon ECR]in the Amazon Elastic Container Registry User Guide.
// For more information on the custom role for replication, see [Creating an IAM role for replication].
//
// When configuring cross-account replication, the destination account must grant
// the source account permission to replicate. This permission is controlled using
// a registry permissions policy. For more information, see PutRegistryPolicy.
//
// [Creating an IAM role for replication]: https://docs.aws.amazon.com/AmazonECR/latest/userguide/replication-creation-templates.html#roles-creatingrole-user-console
// [Using service-linked roles for Amazon ECR]: https://docs.aws.amazon.com/AmazonECR/latest/userguide/using-service-linked-roles.html
