package neptunedata

// GetSparqlStream is generated as a reference stub.
// Executable command wiring lives under cmd/neptunedata.go.
//
// Gets a stream for an RDF graph.
//
// With the Neptune Streams feature, you can generate a complete sequence of
// change-log entries that record every change made to your graph data as it
// happens. GetSparqlStream lets you collect these change-log entries for an RDF
// graph.
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
// Note that the [neptune-db:QueryLanguage:Sparql] IAM condition key can be used in the policy document to restrict
// the use of SPARQL queries (see [Condition keys available in Neptune IAM data-access policy statements]).
//
// [Condition keys available in Neptune IAM data-access policy statements]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html
// [neptune-db:QueryLanguage:Sparql]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html#iam-neptune-condition-keys
// [neptune_streams]: https://docs.aws.amazon.com/neptune/latest/userguide/parameters.html#parameters-db-cluster-parameters-neptune_streams
// [Capturing graph changes in real time using Neptune streams]: https://docs.aws.amazon.com/neptune/latest/userguide/streams.html
// [neptune-db:GetStreamRecords]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#getstreamrecords
