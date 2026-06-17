package s3control

// PutJobTagging is generated as a reference stub.
// Executable command wiring lives under cmd/s3control.go.
//
// Sets the supplied tag-set on an S3 Batch Operations job.
//
// A tag is a key-value pair. You can associate S3 Batch Operations tags with any
// job by sending a PUT request against the tagging subresource that is associated
// with the job. To modify the existing tag set, you can either replace the
// existing tag set entirely, or make changes within the existing tag set by
// retrieving the existing tag set using [GetJobTagging], modify that tag set, and use this
// operation to replace the tag set with the one you modified. For more
// information, see [Controlling access and labeling jobs using tags]in the Amazon S3 User Guide.
//
// - If you send this request with an empty tag set, Amazon S3 deletes the
// existing tag set on the Batch Operations job. If you use this method, you are
// charged for a Tier 1 Request (PUT). For more information, see [Amazon S3 pricing].
//
// - For deleting existing tags for your Batch Operations job, a [DeleteJobTagging]request is
// preferred because it achieves the same result without incurring charges.
//
// - A few things to consider about using tags:
//
// - Amazon S3 limits the maximum number of tags to 50 tags per job.
//
// - You can associate up to 50 tags with a job as long as they have unique tag
// keys.
//
// - A tag key can be up to 128 Unicode characters in length, and tag values can
// be up to 256 Unicode characters in length.
//
// - The key and values are case sensitive.
//
// - For tagging-related restrictions related to characters and encodings, see [User-Defined Tag Restrictions]
// in the Billing and Cost Management User Guide.
//
// Permissions To use the PutJobTagging operation, you must have permission to
// perform the s3:PutJobTagging action.
//
// Related actions include:
//
// [CreateJob]
//
// [GetJobTagging]
//
// [DeleteJobTagging]
//
// [DeleteJobTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteJobTagging.html
// [Controlling access and labeling jobs using tags]: https://docs.aws.amazon.com/AmazonS3/latest/dev/batch-ops-managing-jobs.html#batch-ops-job-tags
// [User-Defined Tag Restrictions]: https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/allocation-tag-restrictions.html
// [GetJobTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetJobTagging.html
// [CreateJob]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateJob.html
// [Amazon S3 pricing]: http://aws.amazon.com/s3/pricing/
