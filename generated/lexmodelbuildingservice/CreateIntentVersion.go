package lexmodelbuildingservice

// CreateIntentVersion is generated as a reference stub.
// Executable command wiring lives under cmd/lexmodelbuildingservice.go.
//
// Creates a new version of an intent based on the $LATEST version of the intent.
// If the $LATEST version of this intent hasn't changed since you last updated it,
// Amazon Lex doesn't create a new version. It returns the last version you
// created.
//
// You can update only the $LATEST version of the intent. You can't update the
// numbered versions that you create with the CreateIntentVersion operation.
//
// When you create a version of an intent, Amazon Lex sets the version to 1.
// Subsequent versions increment by 1. For more information, see versioning-intro.
//
// This operation requires permissions to perform the lex:CreateIntentVersion
// action.
