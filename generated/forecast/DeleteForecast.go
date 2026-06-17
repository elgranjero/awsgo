package forecast

// DeleteForecast is generated as a reference stub.
// Executable command wiring lives under cmd/forecast.go.
//
// Deletes a forecast created using the CreateForecast operation. You can delete only forecasts
// that have a status of ACTIVE or CREATE_FAILED . To get the status, use the DescribeForecast
// operation.
//
// You can't delete a forecast while it is being exported. After a forecast is
// deleted, you can no longer query the forecast.
