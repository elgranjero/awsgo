package dynamodb

// PutItem is generated as a reference stub.
// Executable command wiring lives under cmd/dynamodb.go.
//
// Creates a new item, or replaces an old item with a new item. If an item that
// has the same primary key as the new item already exists in the specified table,
// the new item completely replaces the existing item. You can perform a
// conditional put operation (add a new item if one with the specified primary key
// doesn't exist), or replace an existing item if it has certain attribute values.
// You can return the item's attribute values in the same operation, using the
// ReturnValues parameter.
//
// When you add an item, the primary key attributes are the only required
// attributes.
//
// Empty String and Binary attribute values are allowed. Attribute values of type
// String and Binary must have a length greater than zero if the attribute is used
// as a key attribute for a table or index. Set type attributes cannot be empty.
//
// Invalid Requests with empty values will be rejected with a ValidationException
// exception.
//
// To prevent a new item from replacing an existing item, use a conditional
// expression that contains the attribute_not_exists function with the name of the
// attribute being used as the partition key for the table. Since every record must
// contain that attribute, the attribute_not_exists function will only succeed if
// no matching item exists.
//
// To determine whether PutItem overwrote an existing item, use ReturnValues set
// to ALL_OLD . If the response includes the Attributes element, an existing item
// was overwritten.
//
// For more information about PutItem , see [Working with Items] in the Amazon DynamoDB Developer
// Guide.
//
// [Working with Items]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithItems.html
