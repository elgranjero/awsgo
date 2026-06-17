package internetmonitor

// ListInternetEvents is generated as a reference stub.
// Executable command wiring lives under cmd/internetmonitor.go.
//
// Lists internet events that cause performance or availability issues for client
// locations. Amazon CloudWatch Internet Monitor displays information about recent
// global health events, called internet events, on a global outages map that is
// available to all Amazon Web Services customers.
//
// You can constrain the list of internet events returned by providing a start
// time and end time to define a total time frame for events you want to list. Both
// start time and end time specify the time when an event started. End time is
// optional. If you don't include it, the default end time is the current time.
//
// You can also limit the events returned to a specific status ( ACTIVE or RESOLVED
// ) or type ( PERFORMANCE or AVAILABILITY ).
