package cognitosync

// ListIdentityPoolUsage is generated as a reference stub.
// Executable command wiring lives under cmd/cognitosync.go.
//
// Gets a list of identity pools registered with Cognito.
//
// ListIdentityPoolUsage can only be called with developer credentials. You cannot
// make this API call with the temporary user credentials provided by Cognito
// Identity.
//
// ListIdentityPoolUsage The following examples have been edited for readability.
// POST / HTTP/1.1 CONTENT-TYPE: application/json X-AMZN-REQUESTID:
// 9be7c425-ef05-48c0-aef3-9f0ff2fe17d3 X-AMZ-TARGET:
// com.amazonaws.cognito.sync.model.AWSCognitoSyncService.ListIdentityPoolUsage
// HOST: cognito-sync.us-east-1.amazonaws.com:443 X-AMZ-DATE: 20141111T211414Z
// AUTHORIZATION: AWS4-HMAC-SHA256 Credential=,
// SignedHeaders=content-type;host;x-amz-date;x-amz-target;x-amzn-requestid,
// Signature= { "Operation":
// "com.amazonaws.cognito.sync.model#ListIdentityPoolUsage", "Service":
// "com.amazonaws.cognito.sync.model#AWSCognitoSyncService", "Input": {
// "MaxResults": "2" } } 1.1 200 OK x-amzn-requestid:
// 9be7c425-ef05-48c0-aef3-9f0ff2fe17d3 content-type: application/json
// content-length: 519 date: Tue, 11 Nov 2014 21:14:14 GMT
//
// { "Output": { "__type":
// "com.amazonaws.cognito.sync.model#ListIdentityPoolUsageResponse", "Count": 2,
// "IdentityPoolUsages": [ { "DataStorage": 0, "IdentityPoolId":
// "IDENTITY_POOL_ID", "LastModifiedDate": 1.413836234607E9, "SyncSessionsCount":
// null }, { "DataStorage": 0, "IdentityPoolId": "IDENTITY_POOL_ID",
// "LastModifiedDate": 1.410892165601E9, "SyncSessionsCount": null }],
// "MaxResults": 2, "NextToken":
// "dXMtZWFzdC0xOjBjMWJhMDUyLWUwOTgtNDFmYS1hNzZlLWVhYTJjMTI1Zjg2MQ==" }, "Version":
// "1.0" }
