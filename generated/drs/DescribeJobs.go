package drs

// DescribeJobs is generated as a reference stub.
// Executable command wiring lives under cmd/drs.go.
//
// Returns a list of Jobs. Use the JobsID and fromDate and toDate filters to limit
// which jobs are returned. The response is sorted by creationDataTime - latest
// date first. Jobs are created by the StartRecovery, TerminateRecoveryInstances
// and StartFailbackLaunch APIs. Jobs are also created by DiagnosticLaunch and
// TerminateDiagnosticInstances, which are APIs available only to *Support* and
// only used in response to relevant support tickets.
