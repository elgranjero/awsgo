package redshift

// DescribeClusterSnapshots is generated as a reference stub.
// Executable command wiring lives under cmd/redshift.go.
//
// Returns one or more snapshot objects, which contain metadata about your cluster
// snapshots. By default, this operation returns information about all snapshots of
// all clusters that are owned by your Amazon Web Services account. No information
// is returned for snapshots owned by inactive Amazon Web Services accounts.
//
// If you specify both tag keys and tag values in the same request, Amazon
// Redshift returns all snapshots that match any combination of the specified keys
// and values. For example, if you have owner and environment for tag keys, and
// admin and test for tag values, all snapshots that have any combination of those
// values are returned. Only snapshots that you own are returned in the response;
// shared snapshots are not returned with the tag key and tag value request
// parameters.
//
// If both tag keys and values are omitted from the request, snapshots are
// returned regardless of whether they have tag keys or values associated with
// them.
