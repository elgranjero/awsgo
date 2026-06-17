package storagegateway

// StartCacheReport is generated as a reference stub.
// Executable command wiring lives under cmd/storagegateway.go.
//
// Starts generating a report of the file metadata currently cached by an S3 File
// Gateway for a specific file share. You can use this report to identify and
// resolve issues if you have files failing upload from your gateway to Amazon S3.
// The report is a CSV file containing a list of files which match the set of
// filter parameters you specify in the request.
//
// The Files Failing Upload flag is reset every 24 hours and during gateway
// reboot. If this report captures the files after the reset, but before they
// become flagged again, they will not be reported as Files Failing Upload.
//
// The following requirements must be met to successfully generate a cache report:
//
// - You must have s3:PutObject and s3:AbortMultipartUpload permissions for the
// Amazon S3 bucket where you want to store the cache report.
//
// - No other cache reports can currently be in-progress for the specified file
// share.
//
// - There must be fewer than 10 existing cache reports for the specified file
// share.
//
// - The gateway must be online and connected to Amazon Web Services.
//
// - The root disk must have at least 20GB of free space when report generation
// starts.
//
// - You must specify at least one value for InclusionFilters or ExclusionFilters
// in the request.
