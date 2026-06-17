package dynamodb

// PutResourcePolicy is generated as a reference stub.
// Executable command wiring lives under cmd/dynamodb.go.
//
// Attaches a resource-based policy document to the resource, which can be a table
// or stream. When you attach a resource-based policy using this API, the policy
// application is [eventually consistent].
//
// PutResourcePolicy is an idempotent operation; running it multiple times on the
// same resource using the same policy document will return the same revision ID.
// If you specify an ExpectedRevisionId that doesn't match the current policy's
// RevisionId , the PolicyNotFoundException will be returned.
//
// PutResourcePolicy is an asynchronous operation. If you issue a GetResourcePolicy
// request immediately after a PutResourcePolicy request, DynamoDB might return
// your previous policy, if there was one, or return the PolicyNotFoundException .
// This is because GetResourcePolicy uses an eventually consistent query, and the
// metadata for your policy or table might not be available at that moment. Wait
// for a few seconds, and then try the GetResourcePolicy request again.
//
// [eventually consistent]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadConsistency.html
