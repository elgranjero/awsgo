package synthetics

// DeleteCanary is generated as a reference stub.
// Executable command wiring lives under cmd/synthetics.go.
//
// Permanently deletes the specified canary.
//
// If the canary's ProvisionedResourceCleanup field is set to AUTOMATIC or you
// specify DeleteLambda in this operation as true , CloudWatch Synthetics also
// deletes the Lambda functions and layers that are used by the canary.
//
// Other resources used and created by the canary are not automatically deleted.
// After you delete a canary, you should also delete the following:
//
// - The CloudWatch alarms created for this canary. These alarms have a name of
// Synthetics-Alarm-first-198-characters-of-canary-name-canaryId-alarm number
//
// - Amazon S3 objects and buckets, such as the canary's artifact location.
//
// - IAM roles created for the canary. If they were created in the console,
// these roles have the name
// role/service-role/CloudWatchSyntheticsRole-First-21-Characters-of-CanaryName
//
// - CloudWatch Logs log groups created for the canary. These logs groups have
// the name /aws/lambda/cwsyn-First-21-Characters-of-CanaryName
//
// Before you delete a canary, you might want to use GetCanary to display the
// information about this canary. Make note of the information returned by this
// operation so that you can delete these resources after you delete the canary.
