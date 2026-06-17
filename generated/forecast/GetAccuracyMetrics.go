package forecast

// GetAccuracyMetrics is generated as a reference stub.
// Executable command wiring lives under cmd/forecast.go.
//
// Provides metrics on the accuracy of the models that were trained by the CreatePredictor
// operation. Use metrics to see how well the model performed and to decide whether
// to use the predictor to generate a forecast. For more information, see [Predictor Metrics].
//
// This operation generates metrics for each backtest window that was evaluated.
// The number of backtest windows ( NumberOfBacktestWindows ) is specified using
// the EvaluationParametersobject, which is optionally included in the CreatePredictor request. If
// NumberOfBacktestWindows isn't specified, the number defaults to one.
//
// The parameters of the filling method determine which items contribute to the
// metrics. If you want all items to contribute, specify zero . If you want only
// those items that have complete data in the range being evaluated to contribute,
// specify nan . For more information, see FeaturizationMethod.
//
// Before you can get accuracy metrics, the Status of the predictor must be ACTIVE
// , signifying that training has completed. To get the status, use the DescribePredictoroperation.
//
// [Predictor Metrics]: https://docs.aws.amazon.com/forecast/latest/dg/metrics.html
