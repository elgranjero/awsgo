package iotthingsgraph

// UploadEntityDefinitions is generated as a reference stub.
// Executable command wiring lives under cmd/iotthingsgraph.go.
//
// Asynchronously uploads one or more entity definitions to the user's namespace.
// The document parameter is required if syncWithPublicNamespace and
// deleteExistingEntites are false. If the syncWithPublicNamespace parameter is
// set to true , the user's namespace will synchronize with the latest version of
// the public namespace. If deprecateExistingEntities is set to true, all entities
// in the latest version will be deleted before the new DefinitionDocument is
// uploaded.
//
// When a user uploads entity definitions for the first time, the service creates
// a new namespace for the user. The new namespace tracks the public namespace.
// Currently users can have only one namespace. The namespace version increments
// whenever a user uploads entity definitions that are backwards-incompatible and
// whenever a user sets the syncWithPublicNamespace parameter or the
// deprecateExistingEntities parameter to true .
//
// The IDs for all of the entities should be in URN format. Each entity must be in
// the user's namespace. Users can't create entities in the public namespace, but
// entity definitions can refer to entities in the public namespace.
//
// Valid entities are Device , DeviceModel , Service , Capability , State , Action
// , Event , Property , Mapping , Enum .
//
// Deprecated: since: 2022-08-30
