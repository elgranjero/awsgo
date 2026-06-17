package iot

// DeleteCustomMetric is generated as a reference stub.
// Executable command wiring lives under cmd/iot.go.
//
// Deletes a Device Defender detect custom metric.
//
// Requires permission to access the [DeleteCustomMetric] action.
//
// Before you can delete a custom metric, you must first remove the custom metric
// from all security profiles it's a part of. The security profile associated with
// the custom metric can be found using the [ListSecurityProfiles]API with metricName set to your custom
// metric name.
//
// [DeleteCustomMetric]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
// [ListSecurityProfiles]: https://docs.aws.amazon.com/iot/latest/apireference/API_ListSecurityProfiles.html
