package applicationautoscaling

// GetPredictiveScalingForecast is generated as a reference stub.
// Executable command wiring lives under cmd/applicationautoscaling.go.
//
// Retrieves the forecast data for a predictive scaling policy.
//
// Load forecasts are predictions of the hourly load values using historical load
// data from CloudWatch and an analysis of historical trends. Capacity forecasts
// are represented as predicted values for the minimum capacity that is needed on
// an hourly basis, based on the hourly load forecast.
//
// A minimum of 24 hours of data is required to create the initial forecasts.
// However, having a full 14 days of historical data results in more accurate
// forecasts.
