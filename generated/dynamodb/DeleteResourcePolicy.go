package dynamodb

// DeleteResourcePolicy is generated as a reference stub.
// Executable command wiring lives under cmd/dynamodb.go.
//
// Deletes the resource-based policy attached to the resource, which can be a
// table or stream.
//
// DeleteResourcePolicy is an idempotent operation; running it multiple times on
// the same resource doesn't result in an error response, unless you specify an
// ExpectedRevisionId , which will then return a PolicyNotFoundException .
//
// To make sure that you don't inadvertently lock yourself out of your own
// resources, the root principal in your Amazon Web Services account can perform
// DeleteResourcePolicy requests, even if your resource-based policy explicitly
// denies the root principal's access.
//
// DeleteResourcePolicy is an asynchronous operation. If you issue a
// GetResourcePolicy request immediately after running the DeleteResourcePolicy
// request, DynamoDB might still return the deleted policy. This is because the
// policy for your resource might not have been deleted yet. Wait for a few
// seconds, and then try the GetResourcePolicy request again.
