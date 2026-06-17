package forecast

// CreatePredictor is generated as a reference stub.
// Executable command wiring lives under cmd/forecast.go.
//
// This operation creates a legacy predictor that does not include all the
//
// predictor functionalities provided by Amazon Forecast. To create a predictor
// that is compatible with all aspects of Forecast, use CreateAutoPredictor.
//
// Creates an Amazon Forecast predictor.
//
// In the request, provide a dataset group and either specify an algorithm or let
// Amazon Forecast choose an algorithm for you using AutoML. If you specify an
// algorithm, you also can override algorithm-specific hyperparameters.
//
// Amazon Forecast uses the algorithm to train a predictor using the latest
// version of the datasets in the specified dataset group. You can then generate a
// forecast using the CreateForecastoperation.
//
// To see the evaluation metrics, use the GetAccuracyMetrics operation.
//
// You can specify a featurization configuration to fill and aggregate the data
// fields in the TARGET_TIME_SERIES dataset to improve model training. For more
// information, see FeaturizationConfig.
//
// For RELATED_TIME_SERIES datasets, CreatePredictor verifies that the
// DataFrequency specified when the dataset was created matches the
// ForecastFrequency . TARGET_TIME_SERIES datasets don't have this restriction.
// Amazon Forecast also verifies the delimiter and timestamp format. For more
// information, see howitworks-datasets-groups.
//
// By default, predictors are trained and evaluated at the 0.1 (P10), 0.5 (P50),
// and 0.9 (P90) quantiles. You can choose custom forecast types to train and
// evaluate your predictor by setting the ForecastTypes .
//
// # AutoML
//
// If you want Amazon Forecast to evaluate each algorithm and choose the one that
// minimizes the objective function , set PerformAutoML to true . The objective
// function is defined as the mean of the weighted losses over the forecast types.
// By default, these are the p10, p50, and p90 quantile losses. For more
// information, see EvaluationResult.
//
// When AutoML is enabled, the following properties are disallowed:
//
// - AlgorithmArn
//
// - HPOConfig
//
// - PerformHPO
//
// - TrainingParameters
//
// To get a list of all of your predictors, use the ListPredictors operation.
//
// Before you can use the predictor to create a forecast, the Status of the
// predictor must be ACTIVE , signifying that training has completed. To get the
// status, use the DescribePredictoroperation.
