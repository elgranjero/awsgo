package supplychain

// DeleteInstance is generated as a reference stub.
// Executable command wiring lives under cmd/supplychain.go.
//
// Enables you to programmatically delete an Amazon Web Services Supply Chain
// instance by deleting the KMS keys and relevant information associated with the
// API without using the Amazon Web Services console.
//
// This is an asynchronous operation. Upon receiving a DeleteInstance request,
// Amazon Web Services Supply Chain immediately returns a response with the
// instance resource, delete state while cleaning up all Amazon Web Services
// resources created during the instance creation process. You can use the
// GetInstance action to check the instance status.
