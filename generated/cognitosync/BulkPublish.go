package cognitosync

// BulkPublish is generated as a reference stub.
// Executable command wiring lives under cmd/cognitosync.go.
//
// Initiates a bulk publish of all existing datasets for an Identity Pool to the
// configured stream. Customers are limited to one successful bulk publish per 24
// hours. Bulk publish is an asynchronous request, customers can see the status of
// the request via the GetBulkPublishDetails operation.
//
// This API can only be called with developer credentials. You cannot call this
// API with the temporary user credentials provided by Cognito Identity.
