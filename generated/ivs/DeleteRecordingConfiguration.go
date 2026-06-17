package ivs

// DeleteRecordingConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/ivs.go.
//
// Deletes the recording configuration for the specified ARN.
//
// If you try to delete a recording configuration that is associated with a
// channel, you will get an error (409 ConflictException). To avoid this, for all
// channels that reference the recording configuration, first use UpdateChannelto set the
// recordingConfigurationArn field to an empty string, then use
// DeleteRecordingConfiguration.
