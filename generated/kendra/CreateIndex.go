package kendra

// CreateIndex is generated as a reference stub.
// Executable command wiring lives under cmd/kendra.go.
//
// Creates an Amazon Kendra index. Index creation is an asynchronous API. To
// determine if index creation has completed, check the Status field returned from
// a call to DescribeIndex . The Status field is set to ACTIVE when the index is
// ready to use.
//
// Once the index is active, you can index your documents using the
// BatchPutDocument API or using one of the supported [data sources].
//
// For an example of creating an index and data source using the Python SDK, see [Getting started with Python SDK].
// For an example of creating an index and data source using the Java SDK, see [Getting started with Java SDK].
//
// [data sources]: https://docs.aws.amazon.com/kendra/latest/dg/data-sources.html
// [Getting started with Python SDK]: https://docs.aws.amazon.com/kendra/latest/dg/gs-python.html
// [Getting started with Java SDK]: https://docs.aws.amazon.com/kendra/latest/dg/gs-java.html
