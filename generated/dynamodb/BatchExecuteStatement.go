package dynamodb

// BatchExecuteStatement is generated as a reference stub.
// Executable command wiring lives under cmd/dynamodb.go.
//
// This operation allows you to perform batch reads or writes on data stored in
// DynamoDB, using PartiQL. Each read statement in a BatchExecuteStatement must
// specify an equality condition on all key attributes. This enforces that each
// SELECT statement in a batch returns at most a single item. For more information,
// see [Running batch operations with PartiQL for DynamoDB].
//
// The entire batch must consist of either read statements or write statements,
// you cannot mix both in one batch.
//
// A HTTP 200 response does not mean that all statements in the
// BatchExecuteStatement succeeded. Error details for individual statements can be
// found under the [Error]field of the BatchStatementResponse for each statement.
//
// [Error]: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_BatchStatementResponse.html#DDB-Type-BatchStatementResponse-Error
// [Running batch operations with PartiQL for DynamoDB]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ql-reference.multiplestatements.batching.html
