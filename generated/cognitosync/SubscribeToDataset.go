package cognitosync

// SubscribeToDataset is generated as a reference stub.
// Executable command wiring lives under cmd/cognitosync.go.
//
// Subscribes to receive notifications when a dataset is modified by another
// device.
//
// This API can only be called with temporary credentials provided by Cognito
// Identity. You cannot call this API with developer credentials.
//
// SubscribeToDataset The following examples have been edited for readability.
// POST / HTTP/1.1 CONTENT-TYPE: application/json X-AMZN-REQUESTID:
// 8b9932b7-201d-4418-a960-0a470e11de9f X-AMZ-TARGET:
// com.amazonaws.cognito.sync.model.AWSCognitoSyncService.SubscribeToDataset HOST:
// cognito-sync.us-east-1.amazonaws.com X-AMZ-DATE: 20141004T195350Z
// X-AMZ-SECURITY-TOKEN: AUTHORIZATION: AWS4-HMAC-SHA256 Credential=,
// SignedHeaders=content-type;content-length;host;x-amz-date;x-amz-target,
// Signature= { "Operation":
// "com.amazonaws.cognito.sync.model#SubscribeToDataset", "Service":
// "com.amazonaws.cognito.sync.model#AWSCognitoSyncService", "Input": {
// "IdentityPoolId": "ID_POOL_ID", "IdentityId": "IDENTITY_ID", "DatasetName":
// "Rufus", "DeviceId": "5cd28fbe-dd83-47ab-9f83-19093a5fb014" } } 1.1 200 OK
// x-amzn-requestid: 8b9932b7-201d-4418-a960-0a470e11de9f date: Sat, 04 Oct 2014
// 19:53:50 GMT content-type: application/json content-length: 99
//
// { "Output": { "__type":
// "com.amazonaws.cognito.sync.model#SubscribeToDatasetResponse" }, "Version":
// "1.0" }
