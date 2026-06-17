package cognitosync

// GetIdentityPoolConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/cognitosync.go.
//
// Gets the configuration settings of an identity pool.
//
// This API can only be called with developer credentials. You cannot call this
// API with the temporary user credentials provided by Cognito Identity.
//
// GetIdentityPoolConfiguration The following examples have been edited for
// readability. POST / HTTP/1.1 CONTENT-TYPE: application/json X-AMZN-REQUESTID:
// b1cfdd4b-f620-4fe4-be0f-02024a1d33da X-AMZ-TARGET:
// com.amazonaws.cognito.sync.model.AWSCognitoSyncService.GetIdentityPoolConfiguration
// HOST: cognito-sync.us-east-1.amazonaws.com X-AMZ-DATE: 20141004T195722Z
// AUTHORIZATION: AWS4-HMAC-SHA256 Credential=,
// SignedHeaders=content-type;content-length;host;x-amz-date;x-amz-target,
// Signature= { "Operation":
// "com.amazonaws.cognito.sync.model#GetIdentityPoolConfiguration", "Service":
// "com.amazonaws.cognito.sync.model#AWSCognitoSyncService", "Input": {
// "IdentityPoolId": "ID_POOL_ID" } }
//
// 1.1 200 OK x-amzn-requestid: b1cfdd4b-f620-4fe4-be0f-02024a1d33da date: Sat, 04
// Oct 2014 19:57:22 GMT content-type: application/json content-length: 332
//
// { "Output": { "__type":
// "com.amazonaws.cognito.sync.model#GetIdentityPoolConfigurationResponse",
// "IdentityPoolId": "ID_POOL_ID", "PushSync": { "ApplicationArns":
// ["PLATFORMARN1", "PLATFORMARN2"], "RoleArn": "ROLEARN" } }, "Version": "1.0" }
