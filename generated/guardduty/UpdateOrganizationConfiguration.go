package guardduty

// UpdateOrganizationConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/guardduty.go.
//
// Configures the delegated administrator account with the provided values. You
// must provide a value for either autoEnableOrganizationMembers or autoEnable ,
// but not both.
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
