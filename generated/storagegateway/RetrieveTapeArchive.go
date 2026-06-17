package storagegateway

// RetrieveTapeArchive is generated as a reference stub.
// Executable command wiring lives under cmd/storagegateway.go.
//
// Retrieves an archived virtual tape from the virtual tape shelf (VTS) to a tape
// gateway. Virtual tapes archived in the VTS are not associated with any gateway.
// However after a tape is retrieved, it is associated with a gateway, even though
// it is also listed in the VTS, that is, archive. This operation is only supported
// in the tape gateway type.
//
// Once a tape is successfully retrieved to a gateway, it cannot be retrieved
// again to another gateway. You must archive the tape again before you can
// retrieve it to another gateway. This operation is only supported in the tape
// gateway type.
