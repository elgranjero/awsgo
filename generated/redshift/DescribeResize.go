package redshift

// DescribeResize is generated as a reference stub.
// Executable command wiring lives under cmd/redshift.go.
//
// Returns information about the last resize operation for the specified cluster.
// If no resize operation has ever been initiated for the specified cluster, a
// HTTP 404 error is returned. If a resize operation was initiated and completed,
// the status of the resize remains as SUCCEEDED until the next resize.
//
// A resize operation can be requested using ModifyCluster and specifying a different number or
// type of nodes for the cluster.
