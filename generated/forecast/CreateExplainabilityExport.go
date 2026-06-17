package forecast

// CreateExplainabilityExport is generated as a reference stub.
// Executable command wiring lives under cmd/forecast.go.
//
// Exports an Explainability resource created by the CreateExplainability operation. Exported files
// are exported to an Amazon Simple Storage Service (Amazon S3) bucket.
//
// You must specify a DataDestination object that includes an Amazon S3 bucket and an Identity
// and Access Management (IAM) role that Amazon Forecast can assume to access the
// Amazon S3 bucket. For more information, see aws-forecast-iam-roles.
//
// The Status of the export job must be ACTIVE before you can access the export in
// your Amazon S3 bucket. To get the status, use the DescribeExplainabilityExportoperation.
