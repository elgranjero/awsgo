package glacier

// DescribeJob is generated as a reference stub.
// Executable command wiring lives under cmd/glacier.go.
//
// This operation returns information about a job you previously initiated,
// including the job initiation date, the user who initiated the job, the job
// status code/message and the Amazon SNS topic to notify after Amazon Glacier
// (Glacier) completes the job. For more information about initiating a job, see InitiateJob.
//
// This operation enables you to check the status of your job. However, it is
// strongly recommended that you set up an Amazon SNS topic and specify it in your
// initiate job request so that Glacier can notify the topic after it completes the
// job.
//
// A job ID will not expire for at least 24 hours after Glacier completes the job.
//
// An AWS account has full permission to perform all operations (actions).
// However, AWS Identity and Access Management (IAM) users don't have any
// permissions by default. You must grant them explicit permission to perform
// specific actions. For more information, see [Access Control Using AWS Identity and Access Management (IAM)].
//
// For more information about using this operation, see the documentation for the
// underlying REST API [Describe Job]in the Amazon Glacier Developer Guide.
//
// [Describe Job]: https://docs.aws.amazon.com/amazonglacier/latest/dev/api-describe-job-get.html
// [Access Control Using AWS Identity and Access Management (IAM)]: https://docs.aws.amazon.com/amazonglacier/latest/dev/using-iam-with-amazon-glacier.html
