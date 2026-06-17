package qbusiness

// CreateIndex is generated as a reference stub.
// Executable command wiring lives under cmd/qbusiness.go.
//
// Creates an Amazon Q Business index.
//
// To determine if index creation has completed, check the Status field returned
// from a call to DescribeIndex . The Status field is set to ACTIVE when the index
// is ready to use.
//
// Once the index is active, you can index your documents using the [BatchPutDocument]
// BatchPutDocument API or the [CreateDataSource]CreateDataSource API.
//
// [BatchPutDocument]: https://docs.aws.amazon.com/amazonq/latest/api-reference/API_BatchPutDocument.html
// [CreateDataSource]: https://docs.aws.amazon.com/amazonq/latest/api-reference/API_CreateDataSource.html
