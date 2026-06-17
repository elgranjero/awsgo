package connectcases

// DeleteField is generated as a reference stub.
// Executable command wiring lives under cmd/connectcases.go.
//
// Deletes a field from a cases template.
//
// After a field is deleted:
//
// - You can still retrieve the field by calling BatchGetField .
//
// - You cannot update a deleted field by calling UpdateField ; it throws a
// ValidationException .
//
// - Deleted fields are not included in the ListFields response.
//
// - Calling CreateCase with a deleted field throws a ValidationException
// denoting which field identifiers in the request have been deleted.
//
// - Calling GetCase with a deleted field identifier returns the deleted field's
// value if one exists.
//
// - Calling UpdateCase with a deleted field ID throws a ValidationException if
// the case does not already contain a value for the deleted field. Otherwise it
// succeeds, allowing you to update or remove (using emptyValue: {} ) the field's
// value from the case.
//
// - GetTemplate does not return field IDs for deleted fields.
//
// - GetLayout does not return field IDs for deleted fields.
//
// - Calling SearchCases with the deleted field ID as a filter returns any cases
// that have a value for the deleted field that matches the filter criteria.
//
// - Calling SearchCases with a searchTerm value that matches a deleted field's
// value on a case returns the case in the response.
//
// - Calling BatchPutFieldOptions with a deleted field ID throw a
// ValidationException .
//
// - Calling GetCaseEventConfiguration does not return field IDs for deleted
// fields.
