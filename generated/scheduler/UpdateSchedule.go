package scheduler

// UpdateSchedule is generated as a reference stub.
// Executable command wiring lives under cmd/scheduler.go.
//
// Updates the specified schedule. When you call UpdateSchedule , EventBridge
//
// Scheduler uses all values, including empty values, specified in the request and
// overrides the existing schedule. This is by design. This means that if you do
// not set an optional field in your request, that field will be set to its
// system-default value after the update.
//
// Before calling this operation, we recommend that you call the GetSchedule API
// operation and make a note of all optional parameters for your UpdateSchedule
// call.
