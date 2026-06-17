package iot

// AssociateTargetsWithJob is generated as a reference stub.
// Executable command wiring lives under cmd/iot.go.
//
// Associates a group with a continuous job. The following criteria must be met:
//
// - The job must have been created with the targetSelection field set to
// "CONTINUOUS".
//
// - The job status must currently be "IN_PROGRESS".
//
// - The total number of targets associated with a job must not exceed 100.
//
// Requires permission to access the [AssociateTargetsWithJob] action.
//
// [AssociateTargetsWithJob]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
