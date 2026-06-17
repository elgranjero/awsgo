package cloudfront

// ListDistributionsByRealtimeLogConfig is generated as a reference stub.
// Executable command wiring lives under cmd/cloudfront.go.
//
// Gets a list of distributions that have a cache behavior that's associated with
// the specified real-time log configuration.
//
// You can specify the real-time log configuration by its name or its Amazon
// Resource Name (ARN). You must provide at least one. If you provide both,
// CloudFront uses the name to identify the real-time log configuration to list
// distributions for.
//
// You can optionally specify the maximum number of items to receive in the
// response. If the total number of items in the list exceeds the maximum that you
// specify, or the default maximum, the response is paginated. To get the next page
// of items, send a subsequent request that specifies the NextMarker value from
// the current response as the Marker value in the subsequent request.
