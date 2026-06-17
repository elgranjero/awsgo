package configservice

// PutConfigurationRecorder is generated as a reference stub.
// Executable command wiring lives under cmd/configservice.go.
//
// Creates or updates the customer managed configuration recorder.
//
// You can use this operation to create a new customer managed configuration
// recorder or to update the roleARN and the recordingGroup for an existing
// customer managed configuration recorder.
//
// To start the customer managed configuration recorder and begin recording
// configuration changes for the resource types you specify, use the [StartConfigurationRecorder]operation.
//
// For more information, see [Working with the Configuration Recorder] in the Config Developer Guide.
//
// # One customer managed configuration recorder per account per Region
//
// You can create only one customer managed configuration recorder for each
// account for each Amazon Web Services Region.
//
// Default is to record all supported resource types, excluding the global IAM
// resource types
//
// If you have not specified values for the recordingGroup field, the default for
// the customer managed configuration recorder is to record all supported resource
// types, excluding the global IAM resource types: AWS::IAM::Group ,
// AWS::IAM::Policy , AWS::IAM::Role , and AWS::IAM::User .
//
// # Tags are added at creation and cannot be updated
//
// PutConfigurationRecorder is an idempotent API. Subsequent requests won’t create
// a duplicate resource if one was already created. If a following request has
// different tags values, Config will ignore these differences and treat it as an
// idempotent request of the previous. In this case, tags will not be updated, even
// if they are different.
//
// Use [TagResource] and [UntagResource] to update tags after creation.
//
// [Working with the Configuration Recorder]: https://docs.aws.amazon.com/config/latest/developerguide/stop-start-recorder.html
// [TagResource]: https://docs.aws.amazon.com/config/latest/APIReference/API_TagResource.html
// [UntagResource]: https://docs.aws.amazon.com/config/latest/APIReference/API_UntagResource.html
// [StartConfigurationRecorder]: https://docs.aws.amazon.com/config/latest/APIReference/API_StartConfigurationRecorder.html
