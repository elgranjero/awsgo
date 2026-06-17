package rds

// StartExportTask is generated as a reference stub.
// Executable command wiring lives under cmd/rds.go.
//
// Starts an export of DB snapshot or DB cluster data to Amazon S3. The provided
// IAM role must have access to the S3 bucket.
//
// You can't export snapshot data from RDS Custom DB instances. For more
// information, see [Supported Regions and DB engines for exporting snapshots to S3 in Amazon RDS].
//
// For more information on exporting DB snapshot data, see [Exporting DB snapshot data to Amazon S3] in the Amazon RDS User
// Guide or [Exporting DB cluster snapshot data to Amazon S3]in the Amazon Aurora User Guide.
//
// For more information on exporting DB cluster data, see [Exporting DB cluster data to Amazon S3] in the Amazon Aurora
// User Guide.
//
// [Exporting DB cluster snapshot data to Amazon S3]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-export-snapshot.html
// [Exporting DB cluster data to Amazon S3]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/export-cluster-data.html
// [Exporting DB snapshot data to Amazon S3]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_ExportSnapshot.html
// [Supported Regions and DB engines for exporting snapshots to S3 in Amazon RDS]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.RDS_Fea_Regions_DB-eng.Feature.ExportSnapshotToS3.html
