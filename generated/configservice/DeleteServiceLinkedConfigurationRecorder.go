package configservice

// DeleteServiceLinkedConfigurationRecorder is generated as a reference stub.
// Executable command wiring lives under cmd/configservice.go.
//
// Deletes an existing service-linked configuration recorder.
//
// This operation does not delete the configuration information that was
// previously recorded. You will be able to access the previously recorded
// information by using the [GetResourceConfigHistory]operation, but you will not be able to access this
// information in the Config console until you have created a new service-linked
// configuration recorder for the same service.
//
// # The recording scope determines if you receive configuration items
//
// The recording scope is set by the service that is linked to the configuration
// recorder and determines whether you receive configuration items (CIs) in the
// delivery channel. If the recording scope is internal, you will not receive CIs
// in the delivery channel.
//
// [GetResourceConfigHistory]: https://docs.aws.amazon.com/config/latest/APIReference/API_GetResourceConfigHistory.html
