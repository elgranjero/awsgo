package customerprofiles

// PutProfileObject is generated as a reference stub.
// Executable command wiring lives under cmd/customerprofiles.go.
//
// Adds additional objects to customer profiles of a given ObjectType.
//
// When adding a specific profile object, like a Contact Record, an inferred
// profile can get created if it is not mapped to an existing profile. The
// resulting profile will only have a phone number populated in the standard
// ProfileObject. Any additional Contact Records with the same phone number will be
// mapped to the same inferred profile.
//
// When a ProfileObject is created and if a ProfileObjectType already exists for
// the ProfileObject, it will provide data to a standard profile depending on the
// ProfileObjectType definition.
//
// PutProfileObject needs an ObjectType, which can be created using
// PutProfileObjectType.
