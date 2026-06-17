package storagegateway

// RetrieveTapeRecoveryPoint is generated as a reference stub.
// Executable command wiring lives under cmd/storagegateway.go.
//
// Retrieves the recovery point for the specified virtual tape. This operation is
// only supported in the tape gateway type.
//
// A recovery point is a point in time view of a virtual tape at which all the
// data on the tape is consistent. If your gateway crashes, virtual tapes that have
// recovery points can be recovered to a new gateway.
//
// The virtual tape can be retrieved to only one gateway. The retrieved tape is
// read-only. The virtual tape can be retrieved to only a tape gateway. There is no
// charge for retrieving recovery points.
