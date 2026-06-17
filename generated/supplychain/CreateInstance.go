package supplychain

// CreateInstance is generated as a reference stub.
// Executable command wiring lives under cmd/supplychain.go.
//
// Enables you to programmatically create an Amazon Web Services Supply Chain
// instance by applying KMS keys and relevant information associated with the API
// without using the Amazon Web Services console.
//
// This is an asynchronous operation. Upon receiving a CreateInstance request,
// Amazon Web Services Supply Chain immediately returns the instance resource,
// instance ID, and the initializing state while simultaneously creating all
// required Amazon Web Services resources for an instance creation. You can use
// GetInstance to check the status of the instance. If the instance results in an
// unhealthy state, you need to check the error message, delete the current
// instance, and recreate a new one based on the mitigation from the error message.
