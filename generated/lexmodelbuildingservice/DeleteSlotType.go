package lexmodelbuildingservice

// DeleteSlotType is generated as a reference stub.
// Executable command wiring lives under cmd/lexmodelbuildingservice.go.
//
// Deletes all versions of the slot type, including the $LATEST version. To delete
// a specific version of the slot type, use the DeleteSlotTypeVersionoperation.
//
// You can delete a version of a slot type only if it is not referenced. To delete
// a slot type that is referred to in one or more intents, you must remove those
// references first.
//
// If you get the ResourceInUseException exception, the exception provides an
// example reference that shows the intent where the slot type is referenced. To
// remove the reference to the slot type, either update the intent or delete it. If
// you get the same exception when you attempt to delete the slot type again,
// repeat until the slot type has no references and the DeleteSlotType call is
// successful.
//
// This operation requires permission for the lex:DeleteSlotType action.
