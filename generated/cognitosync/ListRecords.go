package cognitosync

// ListRecords is generated as a reference stub.
// Executable command wiring lives under cmd/cognitosync.go.
//
// Gets paginated records, optionally changed after a particular sync count for a
// dataset and identity. With Amazon Cognito Sync, each identity has access only to
// its own data. Thus, the credentials used to make this API call need to have
// access to the identity data.
//
// ListRecords can be called with temporary user credentials provided by Cognito
// Identity or with developer credentials. You should use Cognito Identity
// credentials to make this API call.
//
// ListRecords The following examples have been edited for readability. POST /
// HTTP/1.1 CONTENT-TYPE: application/json X-AMZN-REQUESTID:
// b3d2e31e-d6b7-4612-8e84-c9ba288dab5d X-AMZ-TARGET:
// com.amazonaws.cognito.sync.model.AWSCognitoSyncService.ListRecords HOST:
// cognito-sync.us-east-1.amazonaws.com:443 X-AMZ-DATE: 20141111T183230Z
// AUTHORIZATION: AWS4-HMAC-SHA256 Credential=,
// SignedHeaders=content-type;host;x-amz-date;x-amz-target;x-amzn-requestid,
// Signature= { "Operation": "com.amazonaws.cognito.sync.model#ListRecords",
// "Service": "com.amazonaws.cognito.sync.model#AWSCognitoSyncService", "Input": {
// "IdentityPoolId": "IDENTITY_POOL_ID", "IdentityId": "IDENTITY_ID",
// "DatasetName": "newDataSet" } } 1.1 200 OK x-amzn-requestid:
// b3d2e31e-d6b7-4612-8e84-c9ba288dab5d content-type: application/json
// content-length: 623 date: Tue, 11 Nov 2014 18:32:30 GMT
//
// { "Output": { "__type": "com.amazonaws.cognito.sync.model#ListRecordsResponse",
// "Count": 0, "DatasetDeletedAfterRequestedSyncCount": false, "DatasetExists":
// false, "DatasetSyncCount": 0, "LastModifiedBy": null, "MergedDatasetNames":
// null, "NextToken": null, "Records": [], "SyncSessionToken": "SYNC_SESSION_TOKEN"
// }, "Version": "1.0" }
