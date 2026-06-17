package cloudwatchevents

// DeactivateEventSource is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatchevents.go.
//
// You can use this operation to temporarily stop receiving events from the
// specified partner event source. The matching event bus is not deleted.
//
// When you deactivate a partner event source, the source goes into PENDING state.
// If it remains in PENDING state for more than two weeks, it is deleted.
//
// To activate a deactivated partner event source, use [ActivateEventSource].
//
// [ActivateEventSource]: https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_ActivateEventSource.html
