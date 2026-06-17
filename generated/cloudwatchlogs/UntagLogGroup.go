package cloudwatchlogs

// UntagLogGroup is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatchlogs.go.
//
// The UntagLogGroup operation is on the path to deprecation. We recommend that
// you use [UntagResource]instead.
//
// Removes the specified tags from the specified log group.
//
// To list the tags for a log group, use [ListTagsForResource]. To add tags, use [TagResource].
//
// When using IAM policies to control tag management for CloudWatch Logs log
// groups, the condition keys aws:Resource/key-name and aws:TagKeys cannot be used
// to restrict which tags users can assign.
//
// Deprecated: Please use the generic tagging API UntagResource
//
// [TagResource]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_TagResource.html
// [UntagResource]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_UntagResource.html
// [ListTagsForResource]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_ListTagsForResource.html
