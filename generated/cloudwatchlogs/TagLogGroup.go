package cloudwatchlogs

// TagLogGroup is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatchlogs.go.
//
// The TagLogGroup operation is on the path to deprecation. We recommend that you
// use [TagResource]instead.
//
// Adds or updates the specified tags for the specified log group.
//
// To list the tags for a log group, use [ListTagsForResource]. To remove tags, use [UntagResource].
//
// For more information about tags, see [Tag Log Groups in Amazon CloudWatch Logs] in the Amazon CloudWatch Logs User Guide.
//
// CloudWatch Logs doesn't support IAM policies that prevent users from assigning
// specified tags to log groups using the aws:Resource/key-name  or aws:TagKeys
// condition keys. For more information about using tags to control access, see [Controlling access to Amazon Web Services resources using tags].
//
// Deprecated: Please use the generic tagging API TagResource
//
// [TagResource]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_TagResource.html
// [UntagResource]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_UntagResource.html
// [Tag Log Groups in Amazon CloudWatch Logs]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/Working-with-log-groups-and-streams.html#log-group-tagging
// [Controlling access to Amazon Web Services resources using tags]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_tags.html
// [ListTagsForResource]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_ListTagsForResource.html
