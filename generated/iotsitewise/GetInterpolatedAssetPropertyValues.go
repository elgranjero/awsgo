package iotsitewise

// GetInterpolatedAssetPropertyValues is generated as a reference stub.
// Executable command wiring lives under cmd/iotsitewise.go.
//
// Get interpolated values for an asset property for a specified time interval,
// during a period of time. If your time series is missing data points during the
// specified time interval, you can use interpolation to estimate the missing data.
//
// For example, you can use this operation to return the interpolated temperature
// values for a wind turbine every 24 hours over a duration of 7 days.
//
// To identify an asset property, you must specify one of the following:
//
// - The assetId and propertyId of an asset property.
//
// - A propertyAlias , which is a data stream alias (for example,
// /company/windfarm/3/turbine/7/temperature ). To define an asset property's
// alias, see [UpdateAssetProperty].
//
// [UpdateAssetProperty]: https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_UpdateAssetProperty.html
