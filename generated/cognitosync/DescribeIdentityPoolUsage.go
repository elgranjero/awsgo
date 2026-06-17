package cognitosync

// DescribeIdentityPoolUsage is generated as a reference stub.
// Executable command wiring lives under cmd/cognitosync.go.
//
// Gets usage details (for example, data storage) about a particular identity pool.
//
// This API can only be called with developer credentials. You cannot call this
// API with the temporary user credentials provided by Cognito Identity.
//
// DescribeIdentityPoolUsage The following examples have been edited for
// readability. POST / HTTP/1.1 CONTENT-TYPE: application/json X-AMZN-REQUESTID:
// 8dc0e749-c8cd-48bd-8520-da6be00d528b X-AMZ-TARGET:
// com.amazonaws.cognito.sync.model.AWSCognitoSyncService.DescribeIdentityPoolUsage
// HOST: cognito-sync.us-east-1.amazonaws.com:443 X-AMZ-DATE: 20141111T205737Z
// AUTHORIZATION: AWS4-HMAC-SHA256 Credential=,
// SignedHeaders=content-type;host;x-amz-date;x-amz-target;x-amzn-requestid,
// Signature=
//
// { "Operation": "com.amazonaws.cognito.sync.model#DescribeIdentityPoolUsage",
// "Service": "com.amazonaws.cognito.sync.model#AWSCognitoSyncService", "Input": {
// "IdentityPoolId": "IDENTITY_POOL_ID" } } 1.1 200 OK x-amzn-requestid:
// 8dc0e749-c8cd-48bd-8520-da6be00d528b content-type: application/json
// content-length: 271 date: Tue, 11 Nov 2014 20:57:37 GMT
//
// { "Output": { "__type":
// "com.amazonaws.cognito.sync.model#DescribeIdentityPoolUsageResponse",
// "IdentityPoolUsage": { "DataStorage": 0, "IdentityPoolId": "IDENTITY_POOL_ID",
// "LastModifiedDate": 1.413231134115E9, "SyncSessionsCount": null } }, "Version":
// "1.0" }
