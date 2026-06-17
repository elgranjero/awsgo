package configservice

// PutAggregationAuthorization is generated as a reference stub.
// Executable command wiring lives under cmd/configservice.go.
//
// Authorizes the aggregator account and region to collect data from the source
// account and region.
//
// # Tags are added at creation and cannot be updated with this operation
//
// PutAggregationAuthorization is an idempotent API. Subsequent requests won’t
// create a duplicate resource if one was already created. If a following request
// has different tags values, Config will ignore these differences and treat it as
// an idempotent request of the previous. In this case, tags will not be updated,
// even if they are different.
//
// Use [TagResource] and [UntagResource] to update tags after creation.
//
// [TagResource]: https://docs.aws.amazon.com/config/latest/APIReference/API_TagResource.html
// [UntagResource]: https://docs.aws.amazon.com/config/latest/APIReference/API_UntagResource.html
