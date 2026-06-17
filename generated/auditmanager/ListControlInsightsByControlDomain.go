package auditmanager

// ListControlInsightsByControlDomain is generated as a reference stub.
// Executable command wiring lives under cmd/auditmanager.go.
//
// Lists the latest analytics data for controls within a specific control domain
// across all active assessments.
//
// Control insights are listed only if the control belongs to the control domain
// that was specified and the control collected evidence on the lastUpdated date
// of controlInsightsMetadata . If neither of these conditions are met, no data is
// listed for that control.
