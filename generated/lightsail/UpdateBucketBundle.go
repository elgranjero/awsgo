package lightsail

// UpdateBucketBundle is generated as a reference stub.
// Executable command wiring lives under cmd/lightsail.go.
//
// Updates the bundle, or storage plan, of an existing Amazon Lightsail bucket.
//
// A bucket bundle specifies the monthly cost, storage space, and data transfer
// quota for a bucket. You can update a bucket's bundle only one time within a
// monthly Amazon Web Services billing cycle. To determine if you can update a
// bucket's bundle, use the [GetBuckets]action. The ableToUpdateBundle parameter in the
// response will indicate whether you can currently update a bucket's bundle.
//
// Update a bucket's bundle if it's consistently going over its storage space or
// data transfer quota, or if a bucket's usage is consistently in the lower range
// of its storage space or data transfer quota. Due to the unpredictable usage
// fluctuations that a bucket might experience, we strongly recommend that you
// update a bucket's bundle only as a long-term strategy, instead of as a
// short-term, monthly cost-cutting measure. Choose a bucket bundle that will
// provide the bucket with ample storage space and data transfer for a long time to
// come.
//
// [GetBuckets]: https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_GetBuckets.html
