package auditmanager

// ListAssessmentControlInsightsByControlDomain is generated as a reference stub.
// Executable command wiring lives under cmd/auditmanager.go.
//
// Lists the latest analytics data for controls within a specific control domain
// and a specific active assessment.
//
// Control insights are listed only if the control belongs to the control domain
// and assessment that was specified. Moreover, the control must have collected
// evidence on the lastUpdated date of controlInsightsByAssessment . If neither of
// these conditions are met, no data is listed for that control.
