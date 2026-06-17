package ivsrealtime

// StartComposition is generated as a reference stub.
// Executable command wiring lives under cmd/ivsrealtime.go.
//
// Starts a Composition from a stage based on the configuration provided in the
// request.
//
// A Composition is an ephemeral resource that exists after this operation returns
// successfully. Composition stops and the resource is deleted:
//
// - When StopCompositionis called.
//
// - After a 1-minute timeout, when all participants are disconnected from the
// stage.
//
// - After a 1-minute timeout, if there are no participants in the stage when
// StartComposition is called.
//
// - When broadcasting to the IVS channel fails and all retries are exhausted.
//
// - When broadcasting is disconnected and all attempts to reconnect are
// exhausted.
