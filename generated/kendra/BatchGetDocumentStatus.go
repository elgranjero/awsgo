package kendra

// BatchGetDocumentStatus is generated as a reference stub.
// Executable command wiring lives under cmd/kendra.go.
//
// Returns the indexing status for one or more documents submitted with the [BatchPutDocument] API.
//
// When you use the BatchPutDocument API, documents are indexed asynchronously.
// You can use the BatchGetDocumentStatus API to get the current status of a list
// of documents so that you can determine if they have been successfully indexed.
//
// You can also use the BatchGetDocumentStatus API to check the status of the [BatchDeleteDocument]
// API. When a document is deleted from the index, Amazon Kendra returns NOT_FOUND
// as the status.
//
// [BatchPutDocument]: https://docs.aws.amazon.com/kendra/latest/dg/API_BatchPutDocument.html
// [BatchDeleteDocument]: https://docs.aws.amazon.com/kendra/latest/dg/API_BatchDeleteDocument.html
