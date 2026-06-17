package cloudtrail

// UpdateEventDataStore is generated as a reference stub.
// Executable command wiring lives under cmd/cloudtrail.go.
//
// Updates an event data store. The required EventDataStore value is an ARN or the
// ID portion of the ARN. Other parameters are optional, but at least one optional
// parameter must be specified, or CloudTrail throws an error. RetentionPeriod is
// in days, and valid values are integers between 7 and 3653 if the BillingMode is
// set to EXTENDABLE_RETENTION_PRICING , or between 7 and 2557 if BillingMode is
// set to FIXED_RETENTION_PRICING . By default, TerminationProtection is enabled.
//
// For event data stores for CloudTrail events, AdvancedEventSelectors includes or
// excludes management, data, or network activity events in your event data store.
// For more information about AdvancedEventSelectors , see [AdvancedEventSelectors].
//
// For event data stores for CloudTrail Insights events, Config configuration
// items, Audit Manager evidence, or non-Amazon Web Services events,
// AdvancedEventSelectors includes events of that type in your event data store.
//
// [AdvancedEventSelectors]: https://docs.aws.amazon.com/awscloudtrail/latest/APIReference/API_AdvancedEventSelector.html
