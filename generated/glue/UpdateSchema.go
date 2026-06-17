package glue

// UpdateSchema is generated as a reference stub.
// Executable command wiring lives under cmd/glue.go.
//
// Updates the description, compatibility setting, or version checkpoint for a
// schema set.
//
// For updating the compatibility setting, the call will not validate
// compatibility for the entire set of schema versions with the new compatibility
// setting. If the value for Compatibility is provided, the VersionNumber (a
// checkpoint) is also required. The API will validate the checkpoint version
// number for consistency.
//
// If the value for the VersionNumber (checkpoint) is provided, Compatibility is
// optional and this can be used to set/reset a checkpoint for the schema.
//
// This update will happen only if the schema is in the AVAILABLE state.
