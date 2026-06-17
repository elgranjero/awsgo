package s3control

// SubmitMultiRegionAccessPointRoutes is generated as a reference stub.
// Executable command wiring lives under cmd/s3control.go.
//
// This operation is not supported by directory buckets.
//
// Submits an updated route configuration for a Multi-Region Access Point. This
// API operation updates the routing status for the specified Regions from active
// to passive, or from passive to active. A value of 0 indicates a passive status,
// which means that traffic won't be routed to the specified Region. A value of 100
// indicates an active status, which means that traffic will be routed to the
// specified Region. At least one Region must be active at all times.
//
// When the routing configuration is changed, any in-progress operations (uploads,
// copies, deletes, and so on) to formerly active Regions will continue to run to
// their final completion state (success or failure). The routing configurations of
// any Regions that aren’t specified remain unchanged.
//
// Updated routing configurations might not be immediately applied. It can take up
// to 2 minutes for your changes to take effect.
//
// To submit routing control changes and failover requests, use the Amazon S3
// failover control infrastructure endpoints in these five Amazon Web Services
// Regions:
//
// - us-east-1
//
// - us-west-2
//
// - ap-southeast-2
//
// - ap-northeast-1
//
// - eu-west-1
