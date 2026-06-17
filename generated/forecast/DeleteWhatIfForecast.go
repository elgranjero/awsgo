package forecast

// DeleteWhatIfForecast is generated as a reference stub.
// Executable command wiring lives under cmd/forecast.go.
//
// Deletes a what-if forecast created using the CreateWhatIfForecast operation. You can delete only
// what-if forecasts that have a status of ACTIVE or CREATE_FAILED . To get the
// status, use the DescribeWhatIfForecastoperation.
//
// You can't delete a what-if forecast while it is being exported. After a what-if
// forecast is deleted, you can no longer query the what-if analysis.
