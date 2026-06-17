package redshift

// DeleteClusterSnapshot is generated as a reference stub.
// Executable command wiring lives under cmd/redshift.go.
//
// Deletes the specified manual snapshot. The snapshot must be in the available
// state, with no other users authorized to access the snapshot.
//
// Unlike automated snapshots, manual snapshots are retained even after you delete
// your cluster. Amazon Redshift does not delete your manual snapshots. You must
// delete manual snapshot explicitly to avoid getting charged. If other accounts
// are authorized to access the snapshot, you must revoke all of the authorizations
// before you can delete the snapshot.
