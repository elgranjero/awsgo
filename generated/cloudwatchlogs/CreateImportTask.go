package cloudwatchlogs

// CreateImportTask is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatchlogs.go.
//
// Starts an import from a data source to CloudWatch Log and creates a managed log
// group as the destination for the imported data. Currently, [CloudTrail Event Data Store]is the only
// supported data source.
//
// The import task must satisfy the following constraints:
//
// - The specified source must be in an ACTIVE state.
//
// - The API caller must have permissions to access the data in the provided
// source and to perform iam:PassRole on the provided import role which has the
// same permissions, as described below.
//
// - The provided IAM role must trust the "cloudtrail.amazonaws.com" principal
// and have the following permissions:
//
// - cloudtrail:GetEventDataStoreData
//
// - logs:CreateLogGroup
//
// - logs:CreateLogStream
//
// - logs:PutResourcePolicy
//
// - (If source has an associated Amazon Web Services KMS Key) kms:Decrypt
//
// - (If source has an associated Amazon Web Services KMS Key)
// kms:GenerateDataKey
//
// Example IAM policy for provided import role:
//
// [ { "Effect": "Allow", "Action": "iam:PassRole", "Resource":
//
// "arn:aws:iam::123456789012:role/apiCallerCredentials", "Condition": {
// "StringLike": { "iam:AssociatedResourceARN":
// "arn:aws:logs:us-east-1:123456789012:log-group:aws/cloudtrail/f1d45bff-d0e3-4868-b5d9-2eb678aa32fb:*"
// } } }, { "Effect": "Allow", "Action": [ "cloudtrail:GetEventDataStoreData" ],
// "Resource": [
// "arn:aws:cloudtrail:us-east-1:123456789012:eventdatastore/f1d45bff-d0e3-4868-b5d9-2eb678aa32fb"
// ] }, { "Effect": "Allow", "Action": [ "logs:CreateImportTask",
// "logs:CreateLogGroup", "logs:CreateLogStream", "logs:PutResourcePolicy" ],
// "Resource": [ "arn:aws:logs:us-east-1:123456789012:log-group:/aws/cloudtrail/*"
// ] }, { "Effect": "Allow", "Action": [ "kms:Decrypt", "kms:GenerateDataKey" ],
// "Resource": [
// "arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012" ]
// } ]
//
// - If the import source has a customer managed key, the
// "cloudtrail.amazonaws.com" principal needs permissions to perform kms:Decrypt
// and kms:GenerateDataKey.
//
// - There can be no more than 3 active imports per account at a given time.
//
// - The startEventTime must be less than or equal to endEventTime.
//
// - The data being imported must be within the specified source's retention
// period.
//
// [CloudTrail Event Data Store]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/query-event-data-store.html
