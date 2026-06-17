package ssm

// GetCalendarState is generated as a reference stub.
// Executable command wiring lives under cmd/ssm.go.
//
// Gets the state of a Amazon Web Services Systems Manager change calendar at the
// current time or a specified time. If you specify a time, GetCalendarState
// returns the state of the calendar at that specific time, and returns the next
// time that the change calendar state will transition. If you don't specify a
// time, GetCalendarState uses the current time. Change Calendar entries have two
// possible states: OPEN or CLOSED .
//
// If you specify more than one calendar in a request, the command returns the
// status of OPEN only if all calendars in the request are open. If one or more
// calendars in the request are closed, the status returned is CLOSED .
//
// For more information about Change Calendar, a tool in Amazon Web Services
// Systems Manager, see [Amazon Web Services Systems Manager Change Calendar]in the Amazon Web Services Systems Manager User Guide.
//
// [Amazon Web Services Systems Manager Change Calendar]: https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-change-calendar.html
