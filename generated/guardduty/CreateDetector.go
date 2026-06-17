package guardduty

// CreateDetector is generated as a reference stub.
// Executable command wiring lives under cmd/guardduty.go.
//
// Creates a single GuardDuty detector. A detector is a resource that represents
// the GuardDuty service. To start using GuardDuty, you must create a detector in
// each Region where you enable the service. You can have only one detector per
// account per Region. All data sources are enabled in a new detector by default.
//
// - When you don't specify any features , with an exception to
// RUNTIME_MONITORING , all the optional features are enabled by default.
//
// - When you specify some of the features , any feature that is not specified in
// the API call gets enabled by default, with an exception to RUNTIME_MONITORING
// .
//
// Specifying both EKS Runtime Monitoring ( EKS_RUNTIME_MONITORING ) and Runtime
// Monitoring ( RUNTIME_MONITORING ) will cause an error. You can add only one of
// these two features because Runtime Monitoring already includes the threat
// detection for Amazon EKS resources. For more information, see [Runtime Monitoring].
//
// There might be regional differences because some data sources might not be
// available in all the Amazon Web Services Regions where GuardDuty is presently
// supported. For more information, see [Regions and endpoints].
//
// [Regions and endpoints]: https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_regions.html
// [Runtime Monitoring]: https://docs.aws.amazon.com/guardduty/latest/ug/runtime-monitoring.html
