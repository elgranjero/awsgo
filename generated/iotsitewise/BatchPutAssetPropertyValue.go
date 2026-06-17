package iotsitewise

// BatchPutAssetPropertyValue is generated as a reference stub.
// Executable command wiring lives under cmd/iotsitewise.go.
//
// Sends a list of asset property values to IoT SiteWise. Each value is a
// timestamp-quality-value (TQV) data point. For more information, see [Ingesting data using the API]in the IoT
// SiteWise User Guide.
//
// To identify an asset property, you must specify one of the following:
//
// - The assetId and propertyId of an asset property.
//
// - A propertyAlias , which is a data stream alias (for example,
// /company/windfarm/3/turbine/7/temperature ). To define an asset property's
// alias, see [UpdateAssetProperty].
//
// With respect to Unix epoch time, IoT SiteWise accepts only TQVs that have a
// timestamp of no more than 7 days in the past and no more than 10 minutes in the
// future. IoT SiteWise rejects timestamps outside of the inclusive range of [-7
// days, +10 minutes] and returns a TimestampOutOfRangeException error.
//
// For each asset property, IoT SiteWise overwrites TQVs with duplicate timestamps
// unless the newer TQV has a different quality. For example, if you store a TQV
// {T1, GOOD, V1} , then storing {T1, GOOD, V2} replaces the existing TQV.
//
// IoT SiteWise authorizes access to each BatchPutAssetPropertyValue entry
// individually. For more information, see [BatchPutAssetPropertyValue authorization]in the IoT SiteWise User Guide.
//
// [Ingesting data using the API]: https://docs.aws.amazon.com/iot-sitewise/latest/userguide/ingest-api.html
// [BatchPutAssetPropertyValue authorization]: https://docs.aws.amazon.com/iot-sitewise/latest/userguide/security_iam_service-with-iam.html#security_iam_service-with-iam-id-based-policies-batchputassetpropertyvalue-action
// [UpdateAssetProperty]: https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_UpdateAssetProperty.html
