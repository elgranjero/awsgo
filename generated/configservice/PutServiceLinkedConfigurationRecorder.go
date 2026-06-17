package configservice

// PutServiceLinkedConfigurationRecorder is generated as a reference stub.
// Executable command wiring lives under cmd/configservice.go.
//
// Creates a service-linked configuration recorder that is linked to a specific
// Amazon Web Services service based on the ServicePrincipal you specify.
//
// The configuration recorder's name , recordingGroup , recordingMode , and
// recordingScope is set by the service that is linked to the configuration
// recorder.
//
// For more information and a list of supported services/service principals, see [Working with the Configuration Recorder]
// in the Config Developer Guide.
//
// This API creates a service-linked role AWSServiceRoleForConfig in your account.
// The service-linked role is created only when the role does not exist in your
// account.
//
// # The recording scope determines if you receive configuration items
//
// The recording scope is set by the service that is linked to the configuration
// recorder and determines whether you receive configuration items (CIs) in the
// delivery channel. If the recording scope is internal, you will not receive CIs
// in the delivery channel.
//
// # Tags are added at creation and cannot be updated with this operation
//
// Use [TagResource] and [UntagResource] to update tags after creation.
//
// [Working with the Configuration Recorder]: https://docs.aws.amazon.com/config/latest/developerguide/stop-start-recorder.html
// [TagResource]: https://docs.aws.amazon.com/config/latest/APIReference/API_TagResource.html
// [UntagResource]: https://docs.aws.amazon.com/config/latest/APIReference/API_UntagResource.html
