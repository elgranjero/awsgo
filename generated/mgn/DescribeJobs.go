package mgn

// DescribeJobs is generated as a reference stub.
// Executable command wiring lives under cmd/mgn.go.
//
// Returns a list of Jobs. Use the JobsID and fromDate and toData filters to limit
// which jobs are returned. The response is sorted by creationDataTime - latest
// date first. Jobs are normally created by the StartTest, StartCutover, and
// TerminateTargetInstances APIs. Jobs are also created by DiagnosticLaunch and
// TerminateDiagnosticInstances, which are APIs available only to *Support* and
// only used in response to relevant support tickets.
