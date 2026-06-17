package deadline

// StartSessionsStatisticsAggregation is generated as a reference stub.
// Executable command wiring lives under cmd/deadline.go.
//
// Starts an asynchronous request for getting aggregated statistics about queues
// and farms. Get the statistics using the GetSessionsStatisticsAggregation
// operation. You can only have one running aggregation for your Deadline Cloud
// farm. Call the GetSessionsStatisticsAggregation operation and check the status
// field to see if an aggregation is running. Statistics are available for 1 hour
// after you call the StartSessionsStatisticsAggregation operation.
