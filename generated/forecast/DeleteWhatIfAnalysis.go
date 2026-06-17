package forecast

// DeleteWhatIfAnalysis is generated as a reference stub.
// Executable command wiring lives under cmd/forecast.go.
//
// Deletes a what-if analysis created using the CreateWhatIfAnalysis operation. You can delete only
// what-if analyses that have a status of ACTIVE or CREATE_FAILED . To get the
// status, use the DescribeWhatIfAnalysisoperation.
//
// You can't delete a what-if analysis while any of its forecasts are being
// exported.
