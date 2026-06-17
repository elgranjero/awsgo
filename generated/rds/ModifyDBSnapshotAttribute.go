package rds

// ModifyDBSnapshotAttribute is generated as a reference stub.
// Executable command wiring lives under cmd/rds.go.
//
// Adds an attribute and values to, or removes an attribute and values from, a
// manual DB snapshot.
//
// To share a manual DB snapshot with other Amazon Web Services accounts, specify
// restore as the AttributeName and use the ValuesToAdd parameter to add a list of
// IDs of the Amazon Web Services accounts that are authorized to restore the
// manual DB snapshot. Uses the value all to make the manual DB snapshot public,
// which means it can be copied or restored by all Amazon Web Services accounts.
//
// Don't add the all value for any manual DB snapshots that contain private
// information that you don't want available to all Amazon Web Services accounts.
//
// If the manual DB snapshot is encrypted, it can be shared, but only by
// specifying a list of authorized Amazon Web Services account IDs for the
// ValuesToAdd parameter. You can't use all as a value for that parameter in this
// case.
//
// To view which Amazon Web Services accounts have access to copy or restore a
// manual DB snapshot, or whether a manual DB snapshot public or private, use the DescribeDBSnapshotAttributes
// API operation. The accounts are returned as values for the restore attribute.
