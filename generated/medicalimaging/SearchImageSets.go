package medicalimaging

// SearchImageSets is generated as a reference stub.
// Executable command wiring lives under cmd/medicalimaging.go.
//
// Search image sets based on defined input attributes.
//
// SearchImageSets accepts a single search query parameter and returns a paginated
// response of all image sets that have the matching criteria. All date range
// queries must be input as (lowerBound, upperBound) .
//
// By default, SearchImageSets uses the updatedAt field for sorting in descending
// order from newest to oldest.
