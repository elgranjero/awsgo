package configservice

// DisassociateResourceTypes is generated as a reference stub.
// Executable command wiring lives under cmd/configservice.go.
//
// Removes all resource types specified in the ResourceTypes list from the [RecordingGroup] of
// configuration recorder and excludes these resource types when recording.
//
// For this operation, the configuration recorder must use a [RecordingStrategy] that is either
// INCLUSION_BY_RESOURCE_TYPES or EXCLUSION_BY_RESOURCE_TYPES .
//
// [RecordingStrategy]: https://docs.aws.amazon.com/config/latest/APIReference/API_RecordingStrategy.html
// [RecordingGroup]: https://docs.aws.amazon.com/config/latest/APIReference/API_RecordingGroup.html
