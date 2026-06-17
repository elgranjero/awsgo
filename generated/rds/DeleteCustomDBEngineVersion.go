package rds

// DeleteCustomDBEngineVersion is generated as a reference stub.
// Executable command wiring lives under cmd/rds.go.
//
// Deletes a custom engine version. To run this command, make sure you meet the
// following prerequisites:
//
// - The CEV must not be the default for RDS Custom. If it is, change the
// default before running this command.
//
// - The CEV must not be associated with an RDS Custom DB instance, RDS Custom
// instance snapshot, or automated backup of your RDS Custom instance.
//
// Typically, deletion takes a few minutes.
//
// The MediaImport service that imports files from Amazon S3 to create CEVs isn't
// integrated with Amazon Web Services CloudTrail. If you turn on data logging for
// Amazon RDS in CloudTrail, calls to the DeleteCustomDbEngineVersion event aren't
// logged. However, you might see calls from the API gateway that accesses your
// Amazon S3 bucket. These calls originate from the MediaImport service for the
// DeleteCustomDbEngineVersion event.
//
// For more information, see [Deleting a CEV] in the Amazon RDS User Guide.
//
// [Deleting a CEV]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-cev.html#custom-cev.delete
