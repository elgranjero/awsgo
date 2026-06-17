package forecast

// CreateExplainability is generated as a reference stub.
// Executable command wiring lives under cmd/forecast.go.
//
// Explainability is only available for Forecasts and Predictors generated from an
// AutoPredictor (CreateAutoPredictor )
//
// Creates an Amazon Forecast Explainability.
//
// Explainability helps you better understand how the attributes in your datasets
// impact forecast. Amazon Forecast uses a metric called Impact scores to quantify
// the relative impact of each attribute and determine whether they increase or
// decrease forecast values.
//
// To enable Forecast Explainability, your predictor must include at least one of
// the following: related time series, item metadata, or additional datasets like
// Holidays and the Weather Index.
//
// CreateExplainability accepts either a Predictor ARN or Forecast ARN. To receive
// aggregated Impact scores for all time series and time points in your datasets,
// provide a Predictor ARN. To receive Impact scores for specific time series and
// time points, provide a Forecast ARN.
//
// # CreateExplainability with a Predictor ARN
//
// You can only have one Explainability resource per predictor. If you already
// enabled ExplainPredictor in CreateAutoPredictor, that predictor already has an Explainability
// resource.
//
// The following parameters are required when providing a Predictor ARN:
//
// - ExplainabilityName - A unique name for the Explainability.
//
// - ResourceArn - The Arn of the predictor.
//
// - TimePointGranularity - Must be set to “ALL”.
//
// - TimeSeriesGranularity - Must be set to “ALL”.
//
// Do not specify a value for the following parameters:
//
// - DataSource - Only valid when TimeSeriesGranularity is “SPECIFIC”.
//
// - Schema - Only valid when TimeSeriesGranularity is “SPECIFIC”.
//
// - StartDateTime - Only valid when TimePointGranularity is “SPECIFIC”.
//
// - EndDateTime - Only valid when TimePointGranularity is “SPECIFIC”.
//
// # CreateExplainability with a Forecast ARN
//
// You can specify a maximum of 50 time series and 500 time points.
//
// The following parameters are required when providing a Predictor ARN:
//
// - ExplainabilityName - A unique name for the Explainability.
//
// - ResourceArn - The Arn of the forecast.
//
// - TimePointGranularity - Either “ALL” or “SPECIFIC”.
//
// - TimeSeriesGranularity - Either “ALL” or “SPECIFIC”.
//
// If you set TimeSeriesGranularity to “SPECIFIC”, you must also provide the
// following:
//
// - DataSource - The S3 location of the CSV file specifying your time series.
//
// - Schema - The Schema defines the attributes and attribute types listed in the
// Data Source.
//
// If you set TimePointGranularity to “SPECIFIC”, you must also provide the
// following:
//
// - StartDateTime - The first timestamp in the range of time points.
//
// - EndDateTime - The last timestamp in the range of time points.
