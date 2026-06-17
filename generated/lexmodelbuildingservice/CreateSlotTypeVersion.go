package lexmodelbuildingservice

// CreateSlotTypeVersion is generated as a reference stub.
// Executable command wiring lives under cmd/lexmodelbuildingservice.go.
//
// Creates a new version of a slot type based on the $LATEST version of the
// specified slot type. If the $LATEST version of this resource has not changed
// since the last version that you created, Amazon Lex doesn't create a new
// version. It returns the last version that you created.
//
// You can update only the $LATEST version of a slot type. You can't update the
// numbered versions that you create with the CreateSlotTypeVersion operation.
//
// When you create a version of a slot type, Amazon Lex sets the version to 1.
// Subsequent versions increment by 1. For more information, see versioning-intro.
//
// This operation requires permissions for the lex:CreateSlotTypeVersion action.
