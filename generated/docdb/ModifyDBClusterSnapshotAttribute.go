package docdb

// ModifyDBClusterSnapshotAttribute is generated as a reference stub.
// Executable command wiring lives under cmd/docdb.go.
//
// Adds an attribute and values to, or removes an attribute and values from, a
// manual cluster snapshot.
//
// To share a manual cluster snapshot with other Amazon Web Services accounts,
// specify restore as the AttributeName , and use the ValuesToAdd parameter to add
// a list of IDs of the Amazon Web Services accounts that are authorized to restore
// the manual cluster snapshot. Use the value all to make the manual cluster
// snapshot public, which means that it can be copied or restored by all Amazon Web
// Services accounts. Do not add the all value for any manual cluster snapshots
// that contain private information that you don't want available to all Amazon Web
// Services accounts. If a manual cluster snapshot is encrypted, it can be shared,
// but only by specifying a list of authorized Amazon Web Services account IDs for
// the ValuesToAdd parameter. You can't use all as a value for that parameter in
// this case.
