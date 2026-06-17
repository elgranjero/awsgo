package keyspaces

// UpdateKeyspace is generated as a reference stub.
// Executable command wiring lives under cmd/keyspaces.go.
//
// Adds a new Amazon Web Services Region to the keyspace. You can add a new
//
// Region to a keyspace that is either a single or a multi-Region keyspace. Amazon
// Keyspaces is going to replicate all tables in the keyspace to the new Region. To
// successfully replicate all tables to the new Region, they must use client-side
// timestamps for conflict resolution. To enable client-side timestamps, specify
// clientSideTimestamps.status = enabled when invoking the API. For more
// information about client-side timestamps, see [Client-side timestamps in Amazon Keyspaces]in the Amazon Keyspaces Developer
// Guide.
//
// To add a Region to a keyspace using the UpdateKeyspace API, the IAM principal
// needs permissions for the following IAM actions:
//
// - cassandra:Alter
//
// - cassandra:AlterMultiRegionResource
//
// - cassandra:Create
//
// - cassandra:CreateMultiRegionResource
//
// - cassandra:Select
//
// - cassandra:SelectMultiRegionResource
//
// - cassandra:Modify
//
// - cassandra:ModifyMultiRegionResource
//
// If the keyspace contains a table that is configured in provisioned mode with
// auto scaling enabled, the following additional IAM actions need to be allowed.
//
// - application-autoscaling:RegisterScalableTarget
//
// - application-autoscaling:DeregisterScalableTarget
//
// - application-autoscaling:DescribeScalableTargets
//
// - application-autoscaling:PutScalingPolicy
//
// - application-autoscaling:DescribeScalingPolicies
//
// To use the UpdateKeyspace API, the IAM principal also needs permissions to
// create a service-linked role with the following elements:
//
// - iam:CreateServiceLinkedRole - The action the principal can perform.
//
// -
// arn:aws:iam::*:role/aws-service-role/replication.cassandra.amazonaws.com/AWSServiceRoleForKeyspacesReplication
//
// - The resource that the action can be performed on.
//
// - iam:AWSServiceName: replication.cassandra.amazonaws.com - The only Amazon
// Web Services service that this role can be attached to is Amazon Keyspaces.
//
// For more information, see [Configure the IAM permissions required to add an Amazon Web Services Region to a keyspace] in the Amazon Keyspaces Developer Guide.
//
// [Client-side timestamps in Amazon Keyspaces]: https://docs.aws.amazon.com/keyspaces/latest/devguide/client-side-timestamps.html
// [Configure the IAM permissions required to add an Amazon Web Services Region to a keyspace]: https://docs.aws.amazon.com/keyspaces/latest/devguide/howitworks_replication_permissions_addReplica.html
