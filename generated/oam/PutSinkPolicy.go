package oam

// PutSinkPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/oam.go.
//
// Creates or updates the resource policy that grants permissions to source
// accounts to link to the monitoring account sink. When you create a sink policy,
// you can grant permissions to all accounts in an organization or to individual
// accounts.
//
// You can also use a sink policy to limit the types of data that is shared. The
// six types of services with their respective resource types that you can allow or
// deny are:
//
// - Metrics - Specify with AWS::CloudWatch::Metric
//
// - Log groups - Specify with AWS::Logs::LogGroup
//
// - Traces - Specify with AWS::XRay::Trace
//
// - Application Insights - Applications - Specify with
// AWS::ApplicationInsights::Application
//
// - Internet Monitor - Specify with AWS::InternetMonitor::Monitor
//
// - Application Signals - Specify with AWS::ApplicationSignals::Service and
// AWS::ApplicationSignals::ServiceLevelObjective
//
// See the examples in this section to see how to specify permitted source
// accounts and data types.
