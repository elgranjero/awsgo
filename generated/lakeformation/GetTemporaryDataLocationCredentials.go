package lakeformation

// GetTemporaryDataLocationCredentials is generated as a reference stub.
// Executable command wiring lives under cmd/lakeformation.go.
//
// Allows a user or application in a secure environment to access data in a
// specific Amazon S3 location registered with Lake Formation by providing
// temporary scoped credentials that are limited to the requested data location and
// the caller's authorized access level.
//
// The API operation returns an error in the following scenarios:
//
// - The data location is not registered with Lake Formation.
//
// - No Glue table is associated with the data location.
//
// - The caller doesn't have required permissions on the associated table. The
// caller must have SELECT or SUPER permissions on the associated table, and
// credential vending for full table access must be enabled in the data lake
// settings.
//
// For more information, see [Application integration for full table access].
//
// - The data location is in a different Amazon Web Services Region. Lake
// Formation doesn't support cross-Region access when vending credentials for a
// data location. Lake Formation only supports Amazon S3 paths registered within
// the same Region as the API call.
//
// [Application integration for full table access]: https://docs.aws.amazon.com/lake-formation/latest/dg/full-table-credential-vending.html
