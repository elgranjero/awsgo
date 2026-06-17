package detective

// StartMonitoringMember is generated as a reference stub.
// Executable command wiring lives under cmd/detective.go.
//
// Sends a request to enable data ingest for a member account that has a status of
// ACCEPTED_BUT_DISABLED .
//
// For valid member accounts, the status is updated as follows.
//
// - If Detective enabled the member account, then the new status is ENABLED .
//
// - If Detective cannot enable the member account, the status remains
// ACCEPTED_BUT_DISABLED .
