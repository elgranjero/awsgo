package cognitosync

// UnsubscribeFromDataset is generated as a reference stub.
// Executable command wiring lives under cmd/cognitosync.go.
//
// Unsubscribes from receiving notifications when a dataset is modified by another
// device.
//
// This API can only be called with temporary credentials provided by Cognito
// Identity. You cannot call this API with developer credentials.
//
// UnsubscribeFromDataset The following examples have been edited for readability.
// POST / HTTP/1.1 CONTENT-TYPE: application/json X-AMZ-REQUESTSUPERTRACE: true
// X-AMZN-REQUESTID: 676896d6-14ca-45b1-8029-6d36b10a077e X-AMZ-TARGET:
// com.amazonaws.cognito.sync.model.AWSCognitoSyncService.UnsubscribeFromDataset
// HOST: cognito-sync.us-east-1.amazonaws.com X-AMZ-DATE: 20141004T195446Z
// X-AMZ-SECURITY-TOKEN: AUTHORIZATION: AWS4-HMAC-SHA256 Credential=,
// SignedHeaders=content-type;content-length;host;x-amz-date;x-amz-target,
// Signature= { "Operation":
// "com.amazonaws.cognito.sync.model#UnsubscribeFromDataset", "Service":
// "com.amazonaws.cognito.sync.model#AWSCognitoSyncService", "Input": {
// "IdentityPoolId": "ID_POOL_ID", "IdentityId": "IDENTITY_ID", "DatasetName":
// "Rufus", "DeviceId": "5cd28fbe-dd83-47ab-9f83-19093a5fb014" } } 1.1 200 OK
// x-amzn-requestid: 676896d6-14ca-45b1-8029-6d36b10a077e date: Sat, 04 Oct 2014
// 19:54:46 GMT content-type: application/json content-length: 103
//
// { "Output": { "__type":
// "com.amazonaws.cognito.sync.model#UnsubscribeFromDatasetResponse" }, "Version":
// "1.0" }
