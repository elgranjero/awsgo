package forecast

// CreateAutoPredictor is generated as a reference stub.
// Executable command wiring lives under cmd/forecast.go.
//
// Creates an Amazon Forecast predictor.
//
// Amazon Forecast creates predictors with AutoPredictor, which involves applying
// the optimal combination of algorithms to each time series in your datasets. You
// can use CreateAutoPredictorto create new predictors or upgrade/retrain existing predictors.
//
// # Creating new predictors
//
// The following parameters are required when creating a new predictor:
//
// - PredictorName - A unique name for the predictor.
//
// - DatasetGroupArn - The ARN of the dataset group used to train the predictor.
//
// - ForecastFrequency - The granularity of your forecasts (hourly, daily,
// weekly, etc).
//
// - ForecastHorizon - The number of time-steps that the model predicts. The
// forecast horizon is also called the prediction length.
//
// When creating a new predictor, do not specify a value for ReferencePredictorArn .
//
// # Upgrading and retraining predictors
//
// The following parameters are required when retraining or upgrading a predictor:
//
// - PredictorName - A unique name for the predictor.
//
// - ReferencePredictorArn - The ARN of the predictor to retrain or upgrade.
//
// When upgrading or retraining a predictor, only specify values for the
// ReferencePredictorArn and PredictorName .
