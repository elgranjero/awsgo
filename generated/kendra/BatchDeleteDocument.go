package kendra

// BatchDeleteDocument is generated as a reference stub.
// Executable command wiring lives under cmd/kendra.go.
//
// Removes one or more documents from an index. The documents must have been added
// with the BatchPutDocument API.
//
// The documents are deleted asynchronously. You can see the progress of the
// deletion by using Amazon Web Services CloudWatch. Any error messages related to
// the processing of the batch are sent to your Amazon Web Services CloudWatch log.
// You can also use the BatchGetDocumentStatus API to monitor the progress of
// deleting your documents.
//
// Deleting documents from an index using BatchDeleteDocument could take up to an
// hour or more, depending on the number of documents you want to delete.
