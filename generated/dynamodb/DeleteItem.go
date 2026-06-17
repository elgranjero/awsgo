package dynamodb

// DeleteItem is generated as a reference stub.
// Executable command wiring lives under cmd/dynamodb.go.
//
// Deletes a single item in a table by primary key. You can perform a conditional
// delete operation that deletes the item if it exists, or if it has an expected
// attribute value.
//
// In addition to deleting an item, you can also return the item's attribute
// values in the same operation, using the ReturnValues parameter.
//
// Unless you specify conditions, the DeleteItem is an idempotent operation;
// running it multiple times on the same item or attribute does not result in an
// error response.
//
// Conditional deletes are useful for deleting items only if specific conditions
// are met. If those conditions are met, DynamoDB performs the delete. Otherwise,
// the item is not deleted.
