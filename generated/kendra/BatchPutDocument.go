package kendra

// BatchPutDocument is generated as a reference stub.
// Executable command wiring lives under cmd/kendra.go.
//
// Adds one or more documents to an index.
//
// The BatchPutDocument API enables you to ingest inline documents or a set of
// documents stored in an Amazon S3 bucket. Use this API to ingest your text and
// unstructured text into an index, add custom attributes to the documents, and to
// attach an access control list to the documents added to the index.
//
// The documents are indexed asynchronously. You can see the progress of the batch
// using Amazon Web Services CloudWatch. Any error messages related to processing
// the batch are sent to your Amazon Web Services CloudWatch log. You can also use
// the BatchGetDocumentStatus API to monitor the progress of indexing your
// documents.
//
// For an example of ingesting inline documents using Python and Java SDKs, see [Adding files directly to an index].
//
// [Adding files directly to an index]: https://docs.aws.amazon.com/kendra/latest/dg/in-adding-binary-doc.html
