package quicksight

// DescribeAssetBundleExportJob is generated as a reference stub.
// Executable command wiring lives under cmd/quicksight.go.
//
// Describes an existing export job.
//
// Poll job descriptions after a job starts to know the status of the job. When a
// job succeeds, a URL is provided to download the exported assets' data from.
// Download URLs are valid for five minutes after they are generated. You can call
// the DescribeAssetBundleExportJob API for a new download URL as needed.
//
// Job descriptions are available for 14 days after the job starts.
