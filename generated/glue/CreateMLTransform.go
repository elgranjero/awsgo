package glue

// CreateMLTransform is generated as a reference stub.
// Executable command wiring lives under cmd/glue.go.
//
// Creates an Glue machine learning transform. This operation creates the
// transform and all the necessary parameters to train it.
//
// Call this operation as the first step in the process of using a machine
// learning transform (such as the FindMatches transform) for deduplicating data.
// You can provide an optional Description , in addition to the parameters that you
// want to use for your algorithm.
//
// You must also specify certain parameters for the tasks that Glue runs on your
// behalf as part of learning from your data and creating a high-quality machine
// learning transform. These parameters include Role , and optionally,
// AllocatedCapacity , Timeout , and MaxRetries . For more information, see [Jobs].
//
// [Jobs]: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-jobs-job.html
