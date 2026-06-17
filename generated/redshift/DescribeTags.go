package redshift

// DescribeTags is generated as a reference stub.
// Executable command wiring lives under cmd/redshift.go.
//
// Returns a list of tags. You can return tags from a specific resource by
// specifying an ARN, or you can return all tags for a given type of resource, such
// as clusters, snapshots, and so on.
//
// The following are limitations for DescribeTags :
//
// - You cannot specify an ARN and a resource-type value together in the same
// request.
//
// - You cannot use the MaxRecords and Marker parameters together with the ARN
// parameter.
//
// - The MaxRecords parameter can be a range from 10 to 50 results to return in a
// request.
//
// If you specify both tag keys and tag values in the same request, Amazon
// Redshift returns all resources that match any combination of the specified keys
// and values. For example, if you have owner and environment for tag keys, and
// admin and test for tag values, all resources that have any combination of those
// values are returned.
//
// If both tag keys and values are omitted from the request, resources are
// returned regardless of whether they have tag keys or values associated with
// them.
