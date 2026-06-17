package ssm

// UpdateMaintenanceWindow is generated as a reference stub.
// Executable command wiring lives under cmd/ssm.go.
//
// Updates an existing maintenance window. Only specified parameters are modified.
//
// The value you specify for Duration determines the specific end time for the
// maintenance window based on the time it begins. No maintenance window tasks are
// permitted to start after the resulting endtime minus the number of hours you
// specify for Cutoff . For example, if the maintenance window starts at 3 PM, the
// duration is three hours, and the value you specify for Cutoff is one hour, no
// maintenance window tasks can start after 5 PM.
