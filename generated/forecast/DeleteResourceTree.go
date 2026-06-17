package forecast

// DeleteResourceTree is generated as a reference stub.
// Executable command wiring lives under cmd/forecast.go.
//
// Deletes an entire resource tree. This operation will delete the parent resource
// and its child resources.
//
// Child resources are resources that were created from another resource. For
// example, when a forecast is generated from a predictor, the forecast is the
// child resource and the predictor is the parent resource.
//
// Amazon Forecast resources possess the following parent-child resource
// hierarchies:
//
// - Dataset: dataset import jobs
//
// - Dataset Group: predictors, predictor backtest export jobs, forecasts,
// forecast export jobs
//
// - Predictor: predictor backtest export jobs, forecasts, forecast export jobs
//
// - Forecast: forecast export jobs
//
// DeleteResourceTree will only delete Amazon Forecast resources, and will not
// delete datasets or exported files stored in Amazon S3.
