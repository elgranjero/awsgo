package configservice

// GetResourceConfigHistory is generated as a reference stub.
// Executable command wiring lives under cmd/configservice.go.
//
// For accurate reporting on the compliance status, you must record the
// AWS::Config::ResourceCompliance resource type.
//
// For more information, see [Recording Amazon Web Services Resources] in the Config Resources Developer Guide.
//
// Returns a list of configurations items (CIs) for the specified resource.
//
// # Contents
//
// The list contains details about each state of the resource during the specified
// time interval. If you specified a retention period to retain your CIs between a
// minimum of 30 days and a maximum of 7 years (2557 days), Config returns the CIs
// for the specified retention period.
//
// # Pagination
//
// The response is paginated. By default, Config returns a limit of 10
// configuration items per page. You can customize this number with the limit
// parameter. The response includes a nextToken string. To get the next page of
// results, run the request again and specify the string for the nextToken
// parameter.
//
// Each call to the API is limited to span a duration of seven days. It is likely
// that the number of records returned is smaller than the specified limit . In
// such cases, you can make another call, using the nextToken .
//
// [Recording Amazon Web Services Resources]: https://docs.aws.amazon.com/config/latest/developerguide/select-resources.html
