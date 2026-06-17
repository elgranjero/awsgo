package iot

// DeleteThingType is generated as a reference stub.
// Executable command wiring lives under cmd/iot.go.
//
// Deletes the specified thing type. You cannot delete a thing type if it has
// things associated with it. To delete a thing type, first mark it as deprecated
// by calling DeprecateThingType, then remove any associated things by calling UpdateThing to change the thing
// type on any associated thing, and finally use DeleteThingTypeto delete the thing type.
//
// Requires permission to access the [DeleteThingType] action.
//
// [DeleteThingType]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
