package redshift

// CopyClusterSnapshot is generated as a reference stub.
// Executable command wiring lives under cmd/redshift.go.
//
// Copies the specified automated cluster snapshot to a new manual cluster
// snapshot. The source must be an automated snapshot and it must be in the
// available state.
//
// When you delete a cluster, Amazon Redshift deletes any automated snapshots of
// the cluster. Also, when the retention period of the snapshot expires, Amazon
// Redshift automatically deletes it. If you want to keep an automated snapshot for
// a longer period, you can make a manual copy of the snapshot. Manual snapshots
// are retained until you delete them.
//
// For more information about working with snapshots, go to [Amazon Redshift Snapshots] in the Amazon
// Redshift Cluster Management Guide.
//
// [Amazon Redshift Snapshots]: https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-snapshots.html
