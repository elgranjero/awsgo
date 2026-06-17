package cognitosync

// RegisterDevice is generated as a reference stub.
// Executable command wiring lives under cmd/cognitosync.go.
//
// Registers a device to receive push sync notifications.
//
// This API can only be called with temporary credentials provided by Cognito
// Identity. You cannot call this API with developer credentials.
//
// RegisterDevice The following examples have been edited for readability. POST /
// HTTP/1.1 CONTENT-TYPE: application/json X-AMZN-REQUESTID:
// 368f9200-3eca-449e-93b3-7b9c08d8e185 X-AMZ-TARGET:
// com.amazonaws.cognito.sync.model.AWSCognitoSyncService.RegisterDevice HOST:
// cognito-sync.us-east-1.amazonaws.com X-AMZ-DATE: 20141004T194643Z
// X-AMZ-SECURITY-TOKEN: AUTHORIZATION: AWS4-HMAC-SHA256 Credential=,
// SignedHeaders=content-type;content-length;host;x-amz-date;x-amz-target,
// Signature= { "Operation": "com.amazonaws.cognito.sync.model#RegisterDevice",
// "Service": "com.amazonaws.cognito.sync.model#AWSCognitoSyncService", "Input": {
// "IdentityPoolId": "ID_POOL_ID", "IdentityId": "IDENTITY_ID", "Platform": "GCM",
// "Token": "PUSH_TOKEN" } } 1.1 200 OK x-amzn-requestid:
// 368f9200-3eca-449e-93b3-7b9c08d8e185 date: Sat, 04 Oct 2014 19:46:44 GMT
// content-type: application/json content-length: 145
//
// { "Output": { "__type":
// "com.amazonaws.cognito.sync.model#RegisterDeviceResponse", "DeviceId":
// "5cd28fbe-dd83-47ab-9f83-19093a5fb014" }, "Version": "1.0" }
