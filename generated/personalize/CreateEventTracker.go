package personalize

// CreateEventTracker is generated as a reference stub.
// Executable command wiring lives under cmd/personalize.go.
//
// Creates an event tracker that you use when adding event data to a specified
// dataset group using the [PutEvents]API.
//
// Only one event tracker can be associated with a dataset group. You will get an
// error if you call CreateEventTracker using the same dataset group as an
// existing event tracker.
//
// When you create an event tracker, the response includes a tracking ID, which
// you pass as a parameter when you use the [PutEvents]operation. Amazon Personalize then
// appends the event data to the Item interactions dataset of the dataset group you
// specify in your event tracker.
//
// The event tracker can be in one of the following states:
//
// - CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE FAILED
//
// - DELETE PENDING > DELETE IN_PROGRESS
//
// To get the status of the event tracker, call [DescribeEventTracker].
//
// The event tracker must be in the ACTIVE state before using the tracking ID.
//
// # Related APIs
//
// [ListEventTrackers]
//
// [DescribeEventTracker]
//
// [DeleteEventTracker]
//
// [PutEvents]: https://docs.aws.amazon.com/personalize/latest/dg/API_UBS_PutEvents.html
// [DescribeEventTracker]: https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeEventTracker.html
// [ListEventTrackers]: https://docs.aws.amazon.com/personalize/latest/dg/API_ListEventTrackers.html
// [DeleteEventTracker]: https://docs.aws.amazon.com/personalize/latest/dg/API_DeleteEventTracker.html
