package omics

// BatchDeleteReadSet is generated as a reference stub.
// Executable command wiring lives under cmd/omics.go.
//
// Deletes one or more read sets. If the operation is successful, it returns a
// response with no body. If there is an error with deleting one of the read sets,
// the operation returns an error list. If the operation successfully deletes only
// a subset of files, it will return an error list for the remaining files that
// fail to be deleted. There is a limit of 100 read sets that can be deleted in
// each BatchDeleteReadSet API call.
