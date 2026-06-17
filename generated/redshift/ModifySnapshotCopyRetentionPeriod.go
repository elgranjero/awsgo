package redshift

// ModifySnapshotCopyRetentionPeriod is generated as a reference stub.
// Executable command wiring lives under cmd/redshift.go.
//
// Modifies the number of days to retain snapshots in the destination Amazon Web
// Services Region after they are copied from the source Amazon Web Services
// Region. By default, this operation only changes the retention period of copied
// automated snapshots. The retention periods for both new and existing copied
// automated snapshots are updated with the new retention period. You can set the
// manual option to change only the retention periods of copied manual snapshots.
// If you set this option, only newly copied manual snapshots have the new
// retention period.
