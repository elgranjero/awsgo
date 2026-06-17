package cognitosync

// ListDatasets is generated as a reference stub.
// Executable command wiring lives under cmd/cognitosync.go.
//
// Lists datasets for an identity. With Amazon Cognito Sync, each identity has
// access only to its own data. Thus, the credentials used to make this API call
// need to have access to the identity data.
//
// ListDatasets can be called with temporary user credentials provided by Cognito
// Identity or with developer credentials. You should use the Cognito Identity
// credentials to make this API call.
//
// ListDatasets The following examples have been edited for readability. POST /
// HTTP/1.1 CONTENT-TYPE: application/json X-AMZN-REQUESTID:
// 15225768-209f-4078-aaed-7494ace9f2db X-AMZ-TARGET:
// com.amazonaws.cognito.sync.model.AWSCognitoSyncService.ListDatasets HOST:
// cognito-sync.us-east-1.amazonaws.com:443 X-AMZ-DATE: 20141111T215640Z
// AUTHORIZATION: AWS4-HMAC-SHA256 Credential=,
// SignedHeaders=content-type;host;x-amz-date;x-amz-target;x-amzn-requestid,
// Signature= { "Operation": "com.amazonaws.cognito.sync.model#ListDatasets",
// "Service": "com.amazonaws.cognito.sync.model#AWSCognitoSyncService", "Input": {
// "IdentityPoolId": "IDENTITY_POOL_ID", "IdentityId": "IDENTITY_ID", "MaxResults":
// "3" } } 1.1 200 OK x-amzn-requestid: 15225768-209f-4078-aaed-7494ace9f2db,
// 15225768-209f-4078-aaed-7494ace9f2db content-type: application/json
// content-length: 355 date: Tue, 11 Nov 2014 21:56:40 GMT
//
// { "Output": { "__type":
// "com.amazonaws.cognito.sync.model#ListDatasetsResponse", "Count": 1, "Datasets":
// [ { "CreationDate": 1.412974057151E9, "DataStorage": 16, "DatasetName":
// "my_list", "IdentityId": "IDENTITY_ID", "LastModifiedBy": "123456789012",
// "LastModifiedDate": 1.412974057244E9, "NumRecords": 1 }], "NextToken": null },
// "Version": "1.0" }
