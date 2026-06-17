package kendra

// DeleteIndex is generated as a reference stub.
// Executable command wiring lives under cmd/kendra.go.
//
// Deletes an Amazon Kendra index. An exception is not thrown if the index is
// already being deleted. While the index is being deleted, the Status field
// returned by a call to the DescribeIndex API is set to DELETING .
