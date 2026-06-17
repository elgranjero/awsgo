package iotsitewise

// DeleteTimeSeries is generated as a reference stub.
// Executable command wiring lives under cmd/iotsitewise.go.
//
// Deletes a time series (data stream). If you delete a time series that's
// associated with an asset property, the asset property still exists, but the time
// series will no longer be associated with this asset property.
//
// To identify a time series, do one of the following:
//
// - If the time series isn't associated with an asset property, specify the
// alias of the time series.
//
// - If the time series is associated with an asset property, specify one of the
// following:
//
// - The alias of the time series.
//
// - The assetId and propertyId that identifies the asset property.
