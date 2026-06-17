package redshift

// DescribeUsageLimits is generated as a reference stub.
// Executable command wiring lives under cmd/redshift.go.
//
// Shows usage limits on a cluster. Results are filtered based on the combination
// of input usage limit identifier, cluster identifier, and feature type
// parameters:
//
// - If usage limit identifier, cluster identifier, and feature type are not
// provided, then all usage limit objects for the current account in the current
// region are returned.
//
// - If usage limit identifier is provided, then the corresponding usage limit
// object is returned.
//
// - If cluster identifier is provided, then all usage limit objects for the
// specified cluster are returned.
//
// - If cluster identifier and feature type are provided, then all usage limit
// objects for the combination of cluster and feature are returned.
