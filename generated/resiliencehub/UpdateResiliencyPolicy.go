package resiliencehub

// UpdateResiliencyPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/resiliencehub.go.
//
// Updates a resiliency policy.
//
// Resilience Hub allows you to provide a value of zero for rtoInSecs and rpoInSecs
// of your resiliency policy. But, while assessing your application, the lowest
// possible assessment result is near zero. Hence, if you provide value zero for
// rtoInSecs and rpoInSecs , the estimated workload RTO and estimated workload RPO
// result will be near zero and the Compliance status for your application will be
// set to Policy breached.
