package dynamodb

// GetItem is generated as a reference stub.
// Executable command wiring lives under cmd/dynamodb.go.
//
// The GetItem operation returns a set of attributes for the item with the given
// primary key. If there is no matching item, GetItem does not return any data and
// there will be no Item element in the response.
//
// GetItem provides an eventually consistent read by default. If your application
// requires a strongly consistent read, set ConsistentRead to true . Although a
// strongly consistent read might take more time than an eventually consistent
// read, it always returns the last updated value.
