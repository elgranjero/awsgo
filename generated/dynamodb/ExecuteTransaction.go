package dynamodb

// ExecuteTransaction is generated as a reference stub.
// Executable command wiring lives under cmd/dynamodb.go.
//
// This operation allows you to perform transactional reads or writes on data
// stored in DynamoDB, using PartiQL.
//
// The entire transaction must consist of either read statements or write
// statements, you cannot mix both in one transaction. The EXISTS function is an
// exception and can be used to check the condition of specific attributes of the
// item in a similar manner to ConditionCheck in the [TransactWriteItems] API.
//
// [TransactWriteItems]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/transaction-apis.html#transaction-apis-txwriteitems
