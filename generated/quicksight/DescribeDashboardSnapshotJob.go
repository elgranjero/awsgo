package quicksight

// DescribeDashboardSnapshotJob is generated as a reference stub.
// Executable command wiring lives under cmd/quicksight.go.
//
// Describes an existing snapshot job.
//
// Poll job descriptions after a job starts to know the status of the job. For
// information on available status codes, see JobStatus .
//
// # Registered user support
//
// This API can be called as before to get status of a job started by the same
// Quick Sight user.
//
// # Possible error scenarios
//
// Request will fail with an Access Denied error in the following scenarios:
//
// - The credentials have expired.
//
// - Job has been started by a different user.
//
// - Impersonated Quick Sight user doesn't have access to the specified
// dashboard in the job.
