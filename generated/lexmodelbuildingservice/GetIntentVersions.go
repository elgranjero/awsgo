package lexmodelbuildingservice

// GetIntentVersions is generated as a reference stub.
// Executable command wiring lives under cmd/lexmodelbuildingservice.go.
//
// Gets information about all of the versions of an intent.
//
// The GetIntentVersions operation returns an IntentMetadata object for each
// version of an intent. For example, if an intent has three numbered versions, the
// GetIntentVersions operation returns four IntentMetadata objects in the
// response, one for each numbered version and one for the $LATEST version.
//
// The GetIntentVersions operation always returns at least one version, the
// $LATEST version.
//
// This operation requires permissions for the lex:GetIntentVersions action.
