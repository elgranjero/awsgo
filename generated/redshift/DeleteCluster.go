package redshift

// DeleteCluster is generated as a reference stub.
// Executable command wiring lives under cmd/redshift.go.
//
// Deletes a previously provisioned cluster without its final snapshot being
// created. A successful response from the web service indicates that the request
// was received correctly. Use DescribeClustersto monitor the status of the deletion. The delete
// operation cannot be canceled or reverted once submitted. For more information
// about managing clusters, go to [Amazon Redshift Clusters]in the Amazon Redshift Cluster Management Guide.
//
// If you want to shut down the cluster and retain it for future use, set
// SkipFinalClusterSnapshot to false and specify a name for
// FinalClusterSnapshotIdentifier. You can later restore this snapshot to resume
// using the cluster. If a final cluster snapshot is requested, the status of the
// cluster will be "final-snapshot" while the snapshot is being taken, then it's
// "deleting" once Amazon Redshift begins deleting the cluster.
//
// For more information about managing clusters, go to [Amazon Redshift Clusters] in the Amazon Redshift
// Cluster Management Guide.
//
// [Amazon Redshift Clusters]: https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html
