package storagegateway

// UpdateMaintenanceStartTime is generated as a reference stub.
// Executable command wiring lives under cmd/storagegateway.go.
//
// Updates a gateway's maintenance window schedule, with settings for monthly or
// weekly cadence, specific day and time to begin maintenance, and which types of
// updates to apply. Time configuration uses the gateway's time zone. You can pass
// values for a complete maintenance schedule, or update policy, or both. Previous
// values will persist for whichever setting you choose not to modify. If an
// incomplete or invalid maintenance schedule is passed, the entire request will be
// rejected with an error and no changes will occur.
//
// A complete maintenance schedule must include values for both MinuteOfHour and
// HourOfDay , and either DayOfMonth or DayOfWeek .
//
// We recommend keeping maintenance updates turned on, except in specific use
// cases where the brief disruptions caused by updating the gateway could
// critically impact your deployment.
