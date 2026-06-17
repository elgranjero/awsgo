package dynamodb

// TransactWriteItems is generated as a reference stub.
// Executable command wiring lives under cmd/dynamodb.go.
//
// TransactWriteItems is a synchronous write operation that groups up to 100
// action requests. These actions can target items in different tables, but not in
// different Amazon Web Services accounts or Regions, and no two actions can target
// the same item. For example, you cannot both ConditionCheck and Update the same
// item. The aggregate size of the items in the transaction cannot exceed 4 MB.
//
// The actions are completed atomically so that either all of them succeed, or all
// of them fail. They are defined by the following objects:
//
// - Put — Initiates a PutItem operation to write a new item. This structure
// specifies the primary key of the item to be written, the name of the table to
// write it in, an optional condition expression that must be satisfied for the
// write to succeed, a list of the item's attributes, and a field indicating
// whether to retrieve the item's attributes if the condition is not met.
//
// - Update — Initiates an UpdateItem operation to update an existing item. This
// structure specifies the primary key of the item to be updated, the name of the
// table where it resides, an optional condition expression that must be satisfied
// for the update to succeed, an expression that defines one or more attributes to
// be updated, and a field indicating whether to retrieve the item's attributes if
// the condition is not met.
//
// - Delete — Initiates a DeleteItem operation to delete an existing item. This
// structure specifies the primary key of the item to be deleted, the name of the
// table where it resides, an optional condition expression that must be satisfied
// for the deletion to succeed, and a field indicating whether to retrieve the
// item's attributes if the condition is not met.
//
// - ConditionCheck — Applies a condition to an item that is not being modified
// by the transaction. This structure specifies the primary key of the item to be
// checked, the name of the table where it resides, a condition expression that
// must be satisfied for the transaction to succeed, and a field indicating whether
// to retrieve the item's attributes if the condition is not met.
//
// DynamoDB rejects the entire TransactWriteItems request if any of the following
// is true:
//
// - A condition in one of the condition expressions is not met.
//
// - An ongoing operation is in the process of updating the same item.
//
// - There is insufficient provisioned capacity for the transaction to be
// completed.
//
// - An item size becomes too large (bigger than 400 KB), a local secondary
// index (LSI) becomes too large, or a similar validation error occurs because of
// changes made by the transaction.
//
// - The aggregate size of the items in the transaction exceeds 4 MB.
//
// - There is a user error, such as an invalid data format.
