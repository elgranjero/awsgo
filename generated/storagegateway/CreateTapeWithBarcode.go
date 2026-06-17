package storagegateway

// CreateTapeWithBarcode is generated as a reference stub.
// Executable command wiring lives under cmd/storagegateway.go.
//
// Creates a virtual tape by using your own barcode. You write data to the virtual
// tape and then archive the tape. A barcode is unique and cannot be reused if it
// has already been used on a tape. This applies to barcodes used on deleted tapes.
// This operation is only supported in the tape gateway type.
//
// Cache storage must be allocated to the gateway before you can create a virtual
// tape. Use the AddCacheoperation to add cache storage to a gateway.
