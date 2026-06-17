package lexmodelbuildingservice

// PutSlotType is generated as a reference stub.
// Executable command wiring lives under cmd/lexmodelbuildingservice.go.
//
// Creates a custom slot type or replaces an existing custom slot type.
//
// To create a custom slot type, specify a name for the slot type and a set of
// enumeration values, which are the values that a slot of this type can assume.
// For more information, see how-it-works.
//
// If you specify the name of an existing slot type, the fields in the request
// replace the existing values in the $LATEST version of the slot type. Amazon Lex
// removes the fields that you don't provide in the request. If you don't specify
// required fields, Amazon Lex throws an exception. When you update the $LATEST
// version of a slot type, if a bot uses the $LATEST version of an intent that
// contains the slot type, the bot's status field is set to NOT_BUILT .
//
// This operation requires permissions for the lex:PutSlotType action.
