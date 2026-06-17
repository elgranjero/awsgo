package costexplorer

// GetSavingsPlansUtilizationDetails is generated as a reference stub.
// Executable command wiring lives under cmd/costexplorer.go.
//
// Retrieves attribute data along with aggregate utilization and savings data for
// a given time period. This doesn't support granular or grouped data
// (daily/monthly) in response. You can't retrieve data by dates in a single
// response similar to GetSavingsPlanUtilization , but you have the option to make
// multiple calls to GetSavingsPlanUtilizationDetails by providing individual
// dates. You can use GetDimensionValues in SAVINGS_PLANS to determine the
// possible dimension values.
//
// GetSavingsPlanUtilizationDetails internally groups data by SavingsPlansArn .
