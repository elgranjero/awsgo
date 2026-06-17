package cognitosync

// SetIdentityPoolConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/cognitosync.go.
//
// Sets the necessary configuration for push sync.
//
// This API can only be called with developer credentials. You cannot call this
// API with the temporary user credentials provided by Cognito Identity.
//
// SetIdentityPoolConfiguration The following examples have been edited for
// readability. POST / HTTP/1.1 CONTENT-TYPE: application/json X-AMZN-REQUESTID:
// a46db021-f5dd-45d6-af5b-7069fa4a211b X-AMZ-TARGET:
// com.amazonaws.cognito.sync.model.AWSCognitoSyncService.SetIdentityPoolConfiguration
// HOST: cognito-sync.us-east-1.amazonaws.com X-AMZ-DATE: 20141004T200006Z
// AUTHORIZATION: AWS4-HMAC-SHA256 Credential=,
// SignedHeaders=content-type;content-length;host;x-amz-date;x-amz-target,
// Signature= { "Operation":
// "com.amazonaws.cognito.sync.model#SetIdentityPoolConfiguration", "Service":
// "com.amazonaws.cognito.sync.model#AWSCognitoSyncService", "Input": {
// "IdentityPoolId": "ID_POOL_ID", "PushSync": { "ApplicationArns":
// ["PLATFORMARN1", "PLATFORMARN2"], "RoleArn": "ROLEARN" } } } 1.1 200 OK
// x-amzn-requestid: a46db021-f5dd-45d6-af5b-7069fa4a211b date: Sat, 04 Oct 2014
// 20:00:06 GMT content-type: application/json content-length: 332
//
// { "Output": { "__type":
// "com.amazonaws.cognito.sync.model#SetIdentityPoolConfigurationResponse",
// "IdentityPoolId": "ID_POOL_ID", "PushSync": { "ApplicationArns":
// ["PLATFORMARN1", "PLATFORMARN2"], "RoleArn": "ROLEARN" } }, "Version": "1.0" }
