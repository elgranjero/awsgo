package directoryservice

// CancelSchemaExtension is generated as a reference stub.
// Executable command wiring lives under cmd/directoryservice.go.
//
// Cancels an in-progress schema extension to a Microsoft AD directory. Once a
// schema extension has started replicating to all domain controllers, the task can
// no longer be canceled. A schema extension can be canceled during any of the
// following states; Initializing , CreatingSnapshot , and UpdatingSchema .
