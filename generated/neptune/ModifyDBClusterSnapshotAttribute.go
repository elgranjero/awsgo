package neptune

// ModifyDBClusterSnapshotAttribute is generated as a reference stub.
// Executable command wiring lives under cmd/neptune.go.
//
// Adds an attribute and values to, or removes an attribute and values from, a
// manual DB cluster snapshot.
//
// To share a manual DB cluster snapshot with other Amazon accounts, specify
// restore as the AttributeName and use the ValuesToAdd parameter to add a list of
// IDs of the Amazon accounts that are authorized to restore the manual DB cluster
// snapshot. Use the value all to make the manual DB cluster snapshot public,
// which means that it can be copied or restored by all Amazon accounts. Do not add
// the all value for any manual DB cluster snapshots that contain private
// information that you don't want available to all Amazon accounts. If a manual DB
// cluster snapshot is encrypted, it can be shared, but only by specifying a list
// of authorized Amazon account IDs for the ValuesToAdd parameter. You can't use
// all as a value for that parameter in this case.
//
// To view which Amazon accounts have access to copy or restore a manual DB
// cluster snapshot, or whether a manual DB cluster snapshot public or private, use
// the DescribeDBClusterSnapshotAttributesAPI action.
