package cloudsearchdomain

// UploadDocuments is generated as a reference stub.
// Executable command wiring lives under cmd/cloudsearchdomain.go.
//
// Posts a batch of documents to a search domain for indexing. A document batch is
// a collection of add and delete operations that represent the documents you want
// to add, update, or delete from your domain. Batches can be described in either
// JSON or XML. Each item that you want Amazon CloudSearch to return as a search
// result (such as a product) is represented as a document. Every document has a
// unique ID and one or more fields that contain the data that you want to search
// and return in results. Individual documents cannot contain more than 1 MB of
// data. The entire batch cannot exceed 5 MB. To get the best possible upload
// performance, group add and delete operations in batches that are close the 5 MB
// limit. Submitting a large volume of single-document batches can overload a
// domain's document service.
//
// The endpoint for submitting UploadDocuments requests is domain-specific. To get
// the document endpoint for your domain, use the Amazon CloudSearch configuration
// service DescribeDomains action. A domain's endpoints are also displayed on the
// domain dashboard in the Amazon CloudSearch console.
//
// For more information about formatting your data for Amazon CloudSearch, see [Preparing Your Data] in
// the Amazon CloudSearch Developer Guide. For more information about uploading
// data for indexing, see [Uploading Data]in the Amazon CloudSearch Developer Guide.
//
// [Preparing Your Data]: http://docs.aws.amazon.com/cloudsearch/latest/developerguide/preparing-data.html
// [Uploading Data]: http://docs.aws.amazon.com/cloudsearch/latest/developerguide/uploading-data.html
