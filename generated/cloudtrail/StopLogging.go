package cloudtrail

// StopLogging is generated as a reference stub.
// Executable command wiring lives under cmd/cloudtrail.go.
//
// Suspends the recording of Amazon Web Services API calls and log file delivery
// for the specified trail. Under most circumstances, there is no need to use this
// action. You can update a trail without stopping it first. This action is the
// only way to stop recording. For a trail enabled in all Regions, this operation
// must be called from the Region in which the trail was created, or an
// InvalidHomeRegionException will occur. This operation cannot be called on the
// shadow trails (replicated trails in other Regions) of a trail enabled in all
// Regions.
