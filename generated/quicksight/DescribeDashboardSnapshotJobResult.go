package quicksight

// DescribeDashboardSnapshotJobResult is generated as a reference stub.
// Executable command wiring lives under cmd/quicksight.go.
//
// Describes the result of an existing snapshot job that has finished running.
//
// A finished snapshot job will return a COMPLETED or FAILED status when you poll
// the job with a DescribeDashboardSnapshotJob API call.
//
// If the job has not finished running, this operation returns a message that says
// Dashboard Snapshot Job with id has not reached a terminal state. .
//
// # Registered user support
//
// This API can be called as before to get the result of a job started by the same
// Quick Sight user. The result for the user will be returned in RegisteredUsers
// response attribute. The attribute will contain a list with at most one object in
// it.
//
// # Possible error scenarios
//
// The request fails with an Access Denied error in the following scenarios:
//
// - The credentials have expired.
//
// - The job was started by a different user.
//
// - The registered user doesn't have access to the specified dashboard.
//
// The request succeeds but the job fails in the following scenarios:
//
// - DASHBOARD_ACCESS_DENIED - The registered user lost access to the dashboard.
//
// - CAPABILITY_RESTRICTED - The registered user is restricted from exporting
// data in all selected formats.
//
// The request succeeds but the response contains an error code in the following
// scenarios:
//
// - CAPABILITY_RESTRICTED - The registered user is restricted from exporting
// data in some selected formats.
//
// - RLS_CHANGED - Row-level security settings have changed. Re-run the job with
// current settings.
//
// - CLS_CHANGED - Column-level security settings have changed. Re-run the job
// with current settings.
//
// - DATASET_DELETED - The dataset has been deleted. Verify the dataset exists
// before re-running the job.
