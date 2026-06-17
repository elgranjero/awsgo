package ec2

// DescribeStoreImageTasks is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Describes the progress of the AMI store tasks. You can describe the store tasks
// for specified AMIs. If you don't specify the AMIs, you get a paginated list of
// store tasks from the last 31 days.
//
// For each AMI task, the response indicates if the task is InProgress , Completed
// , or Failed . For tasks InProgress , the response shows the estimated progress
// as a percentage.
//
// Tasks are listed in reverse chronological order. Currently, only tasks from the
// past 31 days can be viewed.
//
// To use this API, you must have the required permissions. For more information,
// see [Permissions for storing and restoring AMIs using S3]in the Amazon EC2 User Guide.
//
// For more information, see [Store and restore an AMI using S3] in the Amazon EC2 User Guide.
//
// [Store and restore an AMI using S3]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ami-store-restore.html
// [Permissions for storing and restoring AMIs using S3]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/work-with-ami-store-restore.html#ami-s3-permissions
