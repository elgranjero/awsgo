package cognitosync

// DescribeIdentityUsage is generated as a reference stub.
// Executable command wiring lives under cmd/cognitosync.go.
//
// Gets usage information for an identity, including number of datasets and data
// usage.
//
// This API can be called with temporary user credentials provided by Cognito
// Identity or with developer credentials.
//
// DescribeIdentityUsage The following examples have been edited for readability.
// POST / HTTP/1.1 CONTENT-TYPE: application/json X-AMZN-REQUESTID:
// 33f9b4e4-a177-4aad-a3bb-6edb7980b283 X-AMZ-TARGET:
// com.amazonaws.cognito.sync.model.AWSCognitoSyncService.DescribeIdentityUsage
// HOST: cognito-sync.us-east-1.amazonaws.com:443 X-AMZ-DATE: 20141111T215129Z
// AUTHORIZATION: AWS4-HMAC-SHA256 Credential=,
// SignedHeaders=content-type;host;x-amz-date;x-amz-target;x-amzn-requestid,
// Signature= { "Operation":
// "com.amazonaws.cognito.sync.model#DescribeIdentityUsage", "Service":
// "com.amazonaws.cognito.sync.model#AWSCognitoSyncService", "Input": {
// "IdentityPoolId": "IDENTITY_POOL_ID", "IdentityId": "IDENTITY_ID" } } 1.1 200 OK
// x-amzn-requestid: 33f9b4e4-a177-4aad-a3bb-6edb7980b283 content-type:
// application/json content-length: 318 date: Tue, 11 Nov 2014 21:51:29 GMT
//
// { "Output": { "__type":
// "com.amazonaws.cognito.sync.model#DescribeIdentityUsageResponse",
// "IdentityUsage": { "DataStorage": 16, "DatasetCount": 1, "IdentityId":
// "IDENTITY_ID", "IdentityPoolId": "IDENTITY_POOL_ID", "LastModifiedDate":
// 1.412974081336E9 } }, "Version": "1.0" }
