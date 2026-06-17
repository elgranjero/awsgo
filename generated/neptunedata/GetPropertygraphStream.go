package neptunedata

// GetPropertygraphStream is generated as a reference stub.
// Executable command wiring lives under cmd/neptunedata.go.
//
// Gets a stream for a property graph.
//
// With the Neptune Streams feature, you can generate a complete sequence of
// change-log entries that record every change made to your graph data as it
// happens. GetPropertygraphStream lets you collect these change-log entries for a
// property graph.
//
// The Neptune streams feature needs to be enabled on your Neptune DBcluster. To
// enable streams, set the [neptune_streams]DB cluster parameter to 1 .
//
// See [Capturing graph changes in real time using Neptune streams].
//
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:GetStreamRecords]IAM action in that cluster.
//
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that enables one of the following IAM actions, depending on the query:
//
// Note that you can restrict property-graph queries using the following IAM
// context keys:
//
// [neptune-db:QueryLanguage:Gremlin]
//
// [neptune-db:QueryLanguage:OpenCypher]
//
// See [Condition keys available in Neptune IAM data-access policy statements]).
//
// [Condition keys available in Neptune IAM data-access policy statements]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html
// [neptune_streams]: https://docs.aws.amazon.com/neptune/latest/userguide/parameters.html#parameters-db-cluster-parameters-neptune_streams
// [Capturing graph changes in real time using Neptune streams]: https://docs.aws.amazon.com/neptune/latest/userguide/streams.html
// [neptune-db:GetStreamRecords]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#getstreamrecords
// [neptune-db:QueryLanguage:Gremlin]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html#iam-neptune-condition-keys
// [neptune-db:QueryLanguage:OpenCypher]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html#iam-neptune-condition-keys
