package cloudtrail

// CancelQuery is generated as a reference stub.
// Executable command wiring lives under cmd/cloudtrail.go.
//
// Cancels a query if the query is not in a terminated state, such as CANCELLED ,
// FAILED , TIMED_OUT , or FINISHED . You must specify an ARN value for
// EventDataStore . The ID of the query that you want to cancel is also required.
// When you run CancelQuery , the query status might show as CANCELLED even if the
// operation is not yet finished.
