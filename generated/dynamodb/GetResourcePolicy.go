package dynamodb

// GetResourcePolicy is generated as a reference stub.
// Executable command wiring lives under cmd/dynamodb.go.
//
// Returns the resource-based policy document attached to the resource, which can
// be a table or stream, in JSON format.
//
// GetResourcePolicy follows an [eventually consistent] model. The following list describes the outcomes
// when you issue the GetResourcePolicy request immediately after issuing another
// request:
//
// - If you issue a GetResourcePolicy request immediately after a
// PutResourcePolicy request, DynamoDB might return a PolicyNotFoundException .
//
// - If you issue a GetResourcePolicy request immediately after a
// DeleteResourcePolicy request, DynamoDB might return the policy that was
// present before the deletion request.
//
// - If you issue a GetResourcePolicy request immediately after a CreateTable
// request, which includes a resource-based policy, DynamoDB might return a
// ResourceNotFoundException or a PolicyNotFoundException .
//
// Because GetResourcePolicy uses an eventually consistent query, the metadata for
// your policy or table might not be available at that moment. Wait for a few
// seconds, and then retry the GetResourcePolicy request.
//
// After a GetResourcePolicy request returns a policy created using the
// PutResourcePolicy request, the policy will be applied in the authorization of
// requests to the resource. Because this process is eventually consistent, it will
// take some time to apply the policy to all requests to a resource. Policies that
// you attach while creating a table using the CreateTable request will always be
// applied to all requests for that table.
//
// [eventually consistent]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadConsistency.html
