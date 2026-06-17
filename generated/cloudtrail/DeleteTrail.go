package cloudtrail

// DeleteTrail is generated as a reference stub.
// Executable command wiring lives under cmd/cloudtrail.go.
//
// Deletes a trail. This operation must be called from the Region in which the
// trail was created. DeleteTrail cannot be called on the shadow trails
// (replicated trails in other Regions) of a trail that is enabled in all Regions.
//
// While deleting a CloudTrail trail is an irreversible action, CloudTrail does
// not delete log files in the Amazon S3 bucket for that trail, the Amazon S3
// bucket itself, or the CloudWatchlog group to which the trail delivers events.
// Deleting a multi-Region trail will stop logging of events in all Amazon Web
// Services Regions enabled in your Amazon Web Services account. Deleting a
// single-Region trail will stop logging of events in that Region only. It will not
// stop logging of events in other Regions even if the trails in those other
// Regions have identical names to the deleted trail.
//
// For information about account closure and deletion of CloudTrail trails, see [https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-account-closure.html].
//
// [https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-account-closure.html]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-account-closure.html
