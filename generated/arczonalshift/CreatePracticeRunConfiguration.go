package arczonalshift

// CreatePracticeRunConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/arczonalshift.go.
//
// A practice run configuration for zonal autoshift is required when you enable
// zonal autoshift. A practice run configuration includes specifications for
// blocked dates and blocked time windows, and for Amazon CloudWatch alarms that
// you create to use with practice runs. The alarms that you specify are an outcome
// alarm, to monitor application health during practice runs and, optionally, a
// blocking alarm, to block practice runs from starting.
//
// When a resource has a practice run configuration, ARC starts zonal shifts for
// the resource weekly, to shift traffic for practice runs. Practice runs help you
// to ensure that shifting away traffic from an Availability Zone during an
// autoshift is safe for your application.
//
// For more information, see [Considerations when you configure zonal autoshift] in the Amazon Application Recovery Controller
// Developer Guide.
//
// [Considerations when you configure zonal autoshift]: https://docs.aws.amazon.com/r53recovery/latest/dg/arc-zonal-autoshift.considerations.html
