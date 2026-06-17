package eventbridge

// StartReplay is generated as a reference stub.
// Executable command wiring lives under cmd/eventbridge.go.
//
// Starts the specified replay. Events are not necessarily replayed in the exact
// same order that they were added to the archive. A replay processes events to
// replay based on the time in the event, and replays them using 1 minute
// intervals. If you specify an EventStartTime and an EventEndTime that covers a
// 20 minute time range, the events are replayed from the first minute of that 20
// minute range first. Then the events from the second minute are replayed. You can
// use DescribeReplay to determine the progress of a replay. The value returned
// for EventLastReplayedTime indicates the time within the specified time range
// associated with the last event replayed.
