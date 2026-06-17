package forecast

// CreateWhatIfForecastExport is generated as a reference stub.
// Executable command wiring lives under cmd/forecast.go.
//
// Exports a forecast created by the CreateWhatIfForecast operation to your Amazon Simple Storage
// Service (Amazon S3) bucket. The forecast file name will match the following
// conventions:
//
// ≈__
//
// The component is in Java SimpleDateFormat (yyyy-MM-ddTHH-mm-ssZ).
//
// You must specify a DataDestination object that includes an Identity and Access Management
// (IAM) role that Amazon Forecast can assume to access the Amazon S3 bucket. For
// more information, see aws-forecast-iam-roles.
//
// For more information, see howitworks-forecast.
//
// To get a list of all your what-if forecast export jobs, use the ListWhatIfForecastExports operation.
//
// The Status of the forecast export job must be ACTIVE before you can access the
// forecast in your Amazon S3 bucket. To get the status, use the DescribeWhatIfForecastExportoperation.
