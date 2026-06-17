package iot

// GetPercentiles is generated as a reference stub.
// Executable command wiring lives under cmd/iot.go.
//
// Groups the aggregated values that match the query into percentile groupings.
// The default percentile groupings are: 1,5,25,50,75,95,99, although you can
// specify your own when you call GetPercentiles . This function returns a value
// for each percentile group specified (or the default percentile groupings). The
// percentile group "1" contains the aggregated field value that occurs in
// approximately one percent of the values that match the query. The percentile
// group "5" contains the aggregated field value that occurs in approximately five
// percent of the values that match the query, and so on. The result is an
// approximation, the more values that match the query, the more accurate the
// percentile values.
//
// Requires permission to access the [GetPercentiles] action.
//
// [GetPercentiles]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
