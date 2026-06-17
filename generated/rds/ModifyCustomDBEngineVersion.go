package rds

// ModifyCustomDBEngineVersion is generated as a reference stub.
// Executable command wiring lives under cmd/rds.go.
//
// Modifies the status of a custom engine version (CEV). You can find CEVs to
// modify by calling DescribeDBEngineVersions .
//
// The MediaImport service that imports files from Amazon S3 to create CEVs isn't
// integrated with Amazon Web Services CloudTrail. If you turn on data logging for
// Amazon RDS in CloudTrail, calls to the ModifyCustomDbEngineVersion event aren't
// logged. However, you might see calls from the API gateway that accesses your
// Amazon S3 bucket. These calls originate from the MediaImport service for the
// ModifyCustomDbEngineVersion event.
//
// For more information, see [Modifying CEV status] in the Amazon RDS User Guide.
//
// [Modifying CEV status]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-cev.html#custom-cev.modify
