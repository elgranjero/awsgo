package iot

// CreateThing is generated as a reference stub.
// Executable command wiring lives under cmd/iot.go.
//
// Creates a thing record in the registry. If this call is made multiple times
// using the same thing name and configuration, the call will succeed. If this call
// is made with the same thing name but different configuration a
// ResourceAlreadyExistsException is thrown.
//
// This is a control plane operation. See [Authorization] for information about authorizing
// control plane actions.
//
// Requires permission to access the [CreateThing] action.
//
// [Authorization]: https://docs.aws.amazon.com/iot/latest/developerguide/iot-authorization.html
// [CreateThing]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
