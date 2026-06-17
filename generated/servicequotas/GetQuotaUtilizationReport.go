package servicequotas

// GetQuotaUtilizationReport is generated as a reference stub.
// Executable command wiring lives under cmd/servicequotas.go.
//
// Retrieves the quota utilization report for your Amazon Web Services account.
// This operation returns paginated results showing your quota usage across all
// Amazon Web Services services, sorted by utilization percentage in descending
// order (highest utilization first).
//
// You must first initiate a report using the StartQuotaUtilizationReport
// operation. The report generation process is asynchronous and may take several
// seconds to complete. Poll this operation periodically to check the status and
// retrieve results when the report is ready.
//
// Each report contains up to 1,000 quota records per page. Use the NextToken
// parameter to retrieve additional pages of results. Reports are automatically
// deleted after 15 minutes.
