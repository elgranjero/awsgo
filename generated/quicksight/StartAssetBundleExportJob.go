package quicksight

// StartAssetBundleExportJob is generated as a reference stub.
// Executable command wiring lives under cmd/quicksight.go.
//
// Starts an Asset Bundle export job.
//
// An Asset Bundle export job exports specified Amazon Quick Sight assets. You can
// also choose to export any asset dependencies in the same job. Export jobs run
// asynchronously and can be polled with a DescribeAssetBundleExportJob API call.
// When a job is successfully completed, a download URL that contains the exported
// assets is returned. The URL is valid for 5 minutes and can be refreshed with a
// DescribeAssetBundleExportJob API call. Each Amazon Quick Sight account can run
// up to 5 export jobs concurrently.
//
// The API caller must have the necessary permissions in their IAM role to access
// each resource before the resources can be exported.
