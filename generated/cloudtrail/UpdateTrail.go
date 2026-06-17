package cloudtrail

// UpdateTrail is generated as a reference stub.
// Executable command wiring lives under cmd/cloudtrail.go.
//
// Updates trail settings that control what events you are logging, and how to
// handle log files. Changes to a trail do not require stopping the CloudTrail
// service. Use this action to designate an existing bucket for log delivery. If
// the existing bucket has previously been a target for CloudTrail log files, an
// IAM policy exists for the bucket. UpdateTrail must be called from the Region in
// which the trail was created; otherwise, an InvalidHomeRegionException is thrown.
