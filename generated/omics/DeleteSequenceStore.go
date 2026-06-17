package omics

// DeleteSequenceStore is generated as a reference stub.
// Executable command wiring lives under cmd/omics.go.
//
// Deletes a sequence store and returns a response with no body if the operation
// is successful. You can only delete a sequence store when it does not contain any
// read sets.
//
// Use the BatchDeleteReadSet API operation to ensure that all read sets in the
// sequence store are deleted. When a sequence store is deleted, all tags
// associated with the store are also deleted.
//
// For more information, see [Deleting HealthOmics reference and sequence stores] in the Amazon Web Services HealthOmics User Guide.
//
// [Deleting HealthOmics reference and sequence stores]: https://docs.aws.amazon.com/omics/latest/dev/deleting-reference-and-sequence-stores.html
