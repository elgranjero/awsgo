package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitosync"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// cognitosyncCmd represents the cognitosync command
var _cognitosyncCmd = &cobra.Command{
	Use:   "cognitosync",
	Short: "AWS cognitosync CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := cognitosync.NewFromConfig(cfg)
		if _cognitosyncBulkPublish {
			cognitosync_BulkPublish(cfg, client)
			return
		}
		if _cognitosyncDeleteDataset {
			cognitosync_DeleteDataset(cfg, client)
			return
		}
		if _cognitosyncDescribeDataset {
			cognitosync_DescribeDataset(cfg, client)
			return
		}
		if _cognitosyncDescribeIdentityPoolUsage {
			cognitosync_DescribeIdentityPoolUsage(cfg, client)
			return
		}
		if _cognitosyncDescribeIdentityUsage {
			cognitosync_DescribeIdentityUsage(cfg, client)
			return
		}
		if _cognitosyncGetBulkPublishDetails {
			cognitosync_GetBulkPublishDetails(cfg, client)
			return
		}
		if _cognitosyncGetCognitoEvents {
			cognitosync_GetCognitoEvents(cfg, client)
			return
		}
		if _cognitosyncGetIdentityPoolConfiguration {
			cognitosync_GetIdentityPoolConfiguration(cfg, client)
			return
		}
		if _cognitosyncListDatasets {
			cognitosync_ListDatasets(cfg, client)
			return
		}
		if _cognitosyncListIdentityPoolUsage {
			cognitosync_ListIdentityPoolUsage(cfg, client)
			return
		}
		if _cognitosyncListRecords {
			cognitosync_ListRecords(cfg, client)
			return
		}
		if _cognitosyncRegisterDevice {
			cognitosync_RegisterDevice(cfg, client)
			return
		}
		if _cognitosyncSetCognitoEvents {
			cognitosync_SetCognitoEvents(cfg, client)
			return
		}
		if _cognitosyncSetIdentityPoolConfiguration {
			cognitosync_SetIdentityPoolConfiguration(cfg, client)
			return
		}
		if _cognitosyncSubscribeToDataset {
			cognitosync_SubscribeToDataset(cfg, client)
			return
		}
		if _cognitosyncUnsubscribeFromDataset {
			cognitosync_UnsubscribeFromDataset(cfg, client)
			return
		}
		if _cognitosyncUpdateRecords {
			cognitosync_UpdateRecords(cfg, client)
			return
		}

	},
}

var (
	_cognitosyncBulkPublish                  bool
	_cognitosyncDeleteDataset                bool
	_cognitosyncDescribeDataset              bool
	_cognitosyncDescribeIdentityPoolUsage    bool
	_cognitosyncDescribeIdentityUsage        bool
	_cognitosyncGetBulkPublishDetails        bool
	_cognitosyncGetCognitoEvents             bool
	_cognitosyncGetIdentityPoolConfiguration bool
	_cognitosyncListDatasets                 bool
	_cognitosyncListIdentityPoolUsage        bool
	_cognitosyncListRecords                  bool
	_cognitosyncRegisterDevice               bool
	_cognitosyncSetCognitoEvents             bool
	_cognitosyncSetIdentityPoolConfiguration bool
	_cognitosyncSubscribeToDataset           bool
	_cognitosyncUnsubscribeFromDataset       bool
	_cognitosyncUpdateRecords                bool

	_cognitosyncClientContext    string
	_cognitosyncCognitoStreams   string
	_cognitosyncDatasetName      string
	_cognitosyncDeviceId         string
	_cognitosyncEvents           string
	_cognitosyncIdentityId       string
	_cognitosyncIdentityPoolId   string
	_cognitosyncLastSyncCount    string
	_cognitosyncMaxResults       string
	_cognitosyncNextToken        string
	_cognitosyncPlatform         string
	_cognitosyncPushSync         string
	_cognitosyncRecordPatches    string
	_cognitosyncSyncSessionToken string
	_cognitosyncToken            string
)

// Initiates a bulk publish of all existing datasets for an Identity Pool to the
// configured stream. Customers are limited to one successful bulk publish per 24
// hours. Bulk publish is an asynchronous request, customers can see the status of
// the request via the GetBulkPublishDetails operation.
//
// This API can only be called with developer credentials. You cannot call this
// API with the temporary user credentials provided by Cognito Identity.
func cognitosync_BulkPublish(cfg aws.Config, client *cognitosync.Client) {
	input := &cognitosync.BulkPublishInput{
		// IdentityPoolId: *string, // Required
	}

	if len(_cognitosyncIdentityPoolId) > 0 {
		input.IdentityPoolId = aws.String(_cognitosyncIdentityPoolId)
	}

	if resp, err := client.BulkPublish(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specific dataset. The dataset will be deleted permanently, and the
// action can't be undone. Datasets that this dataset was merged with will no
// longer report the merge. Any subsequent operation on this dataset will result in
// a ResourceNotFoundException.
//
// This API can be called with temporary user credentials provided by Cognito
// Identity or with developer credentials.
func cognitosync_DeleteDataset(cfg aws.Config, client *cognitosync.Client) {
	input := &cognitosync.DeleteDatasetInput{
		// DatasetName: *string, // Required
		// IdentityId: *string, // Required
		// IdentityPoolId: *string, // Required
	}

	if len(_cognitosyncDatasetName) > 0 {
		input.DatasetName = aws.String(_cognitosyncDatasetName)
	}
	if len(_cognitosyncIdentityId) > 0 {
		input.IdentityId = aws.String(_cognitosyncIdentityId)
	}
	if len(_cognitosyncIdentityPoolId) > 0 {
		input.IdentityPoolId = aws.String(_cognitosyncIdentityPoolId)
	}

	if resp, err := client.DeleteDataset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets meta data about a dataset by identity and dataset name. With Amazon
// Cognito Sync, each identity has access only to its own data. Thus, the
// credentials used to make this API call need to have access to the identity data.
//
// This API can be called with temporary user credentials provided by Cognito
// Identity or with developer credentials. You should use Cognito Identity
// credentials to make this API call.
func cognitosync_DescribeDataset(cfg aws.Config, client *cognitosync.Client) {
	input := &cognitosync.DescribeDatasetInput{
		// DatasetName: *string, // Required
		// IdentityId: *string, // Required
		// IdentityPoolId: *string, // Required
	}

	if len(_cognitosyncDatasetName) > 0 {
		input.DatasetName = aws.String(_cognitosyncDatasetName)
	}
	if len(_cognitosyncIdentityId) > 0 {
		input.IdentityId = aws.String(_cognitosyncIdentityId)
	}
	if len(_cognitosyncIdentityPoolId) > 0 {
		input.IdentityPoolId = aws.String(_cognitosyncIdentityPoolId)
	}

	if resp, err := client.DescribeDataset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets usage details (for example, data storage) about a particular identity pool.
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
func cognitosync_DescribeIdentityPoolUsage(cfg aws.Config, client *cognitosync.Client) {
	input := &cognitosync.DescribeIdentityPoolUsageInput{
		// IdentityPoolId: *string, // Required
	}

	if len(_cognitosyncIdentityPoolId) > 0 {
		input.IdentityPoolId = aws.String(_cognitosyncIdentityPoolId)
	}

	if resp, err := client.DescribeIdentityPoolUsage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

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
func cognitosync_DescribeIdentityUsage(cfg aws.Config, client *cognitosync.Client) {
	input := &cognitosync.DescribeIdentityUsageInput{
		// IdentityId: *string, // Required
		// IdentityPoolId: *string, // Required
	}

	if len(_cognitosyncIdentityId) > 0 {
		input.IdentityId = aws.String(_cognitosyncIdentityId)
	}
	if len(_cognitosyncIdentityPoolId) > 0 {
		input.IdentityPoolId = aws.String(_cognitosyncIdentityPoolId)
	}

	if resp, err := client.DescribeIdentityUsage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the status of the last BulkPublish operation for an identity pool.
// This API can only be called with developer credentials. You cannot call this
// API with the temporary user credentials provided by Cognito Identity.
func cognitosync_GetBulkPublishDetails(cfg aws.Config, client *cognitosync.Client) {
	input := &cognitosync.GetBulkPublishDetailsInput{
		// IdentityPoolId: *string, // Required
	}

	if len(_cognitosyncIdentityPoolId) > 0 {
		input.IdentityPoolId = aws.String(_cognitosyncIdentityPoolId)
	}

	if resp, err := client.GetBulkPublishDetails(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the events and the corresponding Lambda functions associated with an
// identity pool.
//
// This API can only be called with developer credentials. You cannot call this
// API with the temporary user credentials provided by Cognito Identity.
func cognitosync_GetCognitoEvents(cfg aws.Config, client *cognitosync.Client) {
	input := &cognitosync.GetCognitoEventsInput{
		// IdentityPoolId: *string, // Required
	}

	if len(_cognitosyncIdentityPoolId) > 0 {
		input.IdentityPoolId = aws.String(_cognitosyncIdentityPoolId)
	}

	if resp, err := client.GetCognitoEvents(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the configuration settings of an identity pool.
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
func cognitosync_GetIdentityPoolConfiguration(cfg aws.Config, client *cognitosync.Client) {
	input := &cognitosync.GetIdentityPoolConfigurationInput{
		// IdentityPoolId: *string, // Required
	}

	if len(_cognitosyncIdentityPoolId) > 0 {
		input.IdentityPoolId = aws.String(_cognitosyncIdentityPoolId)
	}

	if resp, err := client.GetIdentityPoolConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

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
func cognitosync_ListDatasets(cfg aws.Config, client *cognitosync.Client) {
	input := &cognitosync.ListDatasetsInput{
		// IdentityId: *string, // Required
		// IdentityPoolId: *string, // Required
	}

	if len(_cognitosyncIdentityId) > 0 {
		input.IdentityId = aws.String(_cognitosyncIdentityId)
	}
	if len(_cognitosyncIdentityPoolId) > 0 {
		input.IdentityPoolId = aws.String(_cognitosyncIdentityPoolId)
	}
	if len(_cognitosyncMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cognitosyncMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cognitosyncNextToken) > 0 {
		input.NextToken = aws.String(_cognitosyncNextToken)
	}

	if resp, err := client.ListDatasets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a list of identity pools registered with Cognito.
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
func cognitosync_ListIdentityPoolUsage(cfg aws.Config, client *cognitosync.Client) {
	input := &cognitosync.ListIdentityPoolUsageInput{}

	if len(_cognitosyncMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cognitosyncMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cognitosyncNextToken) > 0 {
		input.NextToken = aws.String(_cognitosyncNextToken)
	}

	if resp, err := client.ListIdentityPoolUsage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

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
func cognitosync_ListRecords(cfg aws.Config, client *cognitosync.Client) {
	input := &cognitosync.ListRecordsInput{
		// DatasetName: *string, // Required
		// IdentityId: *string, // Required
		// IdentityPoolId: *string, // Required
	}

	if len(_cognitosyncDatasetName) > 0 {
		input.DatasetName = aws.String(_cognitosyncDatasetName)
	}
	if len(_cognitosyncIdentityId) > 0 {
		input.IdentityId = aws.String(_cognitosyncIdentityId)
	}
	if len(_cognitosyncIdentityPoolId) > 0 {
		input.IdentityPoolId = aws.String(_cognitosyncIdentityPoolId)
	}
	if len(_cognitosyncLastSyncCount) > 0 {
		if err := assignInputField(input, "LastSyncCount", _cognitosyncLastSyncCount); err != nil {
			log.Errorf("invalid --last-sync-count: %s", err.Error())
			return
		}
	}
	if len(_cognitosyncMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cognitosyncMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cognitosyncNextToken) > 0 {
		input.NextToken = aws.String(_cognitosyncNextToken)
	}
	if len(_cognitosyncSyncSessionToken) > 0 {
		input.SyncSessionToken = aws.String(_cognitosyncSyncSessionToken)
	}

	if resp, err := client.ListRecords(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Registers a device to receive push sync notifications.
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
func cognitosync_RegisterDevice(cfg aws.Config, client *cognitosync.Client) {
	input := &cognitosync.RegisterDeviceInput{
		// IdentityId: *string, // Required
		// IdentityPoolId: *string, // Required
		// Platform: types.Platform, // Required
		// Token: *string, // Required
	}

	if len(_cognitosyncIdentityId) > 0 {
		input.IdentityId = aws.String(_cognitosyncIdentityId)
	}
	if len(_cognitosyncIdentityPoolId) > 0 {
		input.IdentityPoolId = aws.String(_cognitosyncIdentityPoolId)
	}
	if len(_cognitosyncPlatform) > 0 {
		if err := assignInputField(input, "Platform", _cognitosyncPlatform); err != nil {
			log.Errorf("invalid --platform: %s", err.Error())
			return
		}
	}
	if len(_cognitosyncToken) > 0 {
		input.Token = aws.String(_cognitosyncToken)
	}

	if resp, err := client.RegisterDevice(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the AWS Lambda function for a given event type for an identity pool. This
// request only updates the key/value pair specified. Other key/values pairs are
// not updated. To remove a key value pair, pass a empty value for the particular
// key.
//
// This API can only be called with developer credentials. You cannot call this
// API with the temporary user credentials provided by Cognito Identity.
func cognitosync_SetCognitoEvents(cfg aws.Config, client *cognitosync.Client) {
	input := &cognitosync.SetCognitoEventsInput{
		// Events: map[string]string, // Required
		// IdentityPoolId: *string, // Required
	}

	if len(_cognitosyncEvents) > 0 {
		if err := assignInputField(input, "Events", _cognitosyncEvents); err != nil {
			log.Errorf("invalid --events: %s", err.Error())
			return
		}
	}
	if len(_cognitosyncIdentityPoolId) > 0 {
		input.IdentityPoolId = aws.String(_cognitosyncIdentityPoolId)
	}

	if resp, err := client.SetCognitoEvents(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the necessary configuration for push sync.
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
func cognitosync_SetIdentityPoolConfiguration(cfg aws.Config, client *cognitosync.Client) {
	input := &cognitosync.SetIdentityPoolConfigurationInput{
		// IdentityPoolId: *string, // Required
	}

	if len(_cognitosyncIdentityPoolId) > 0 {
		input.IdentityPoolId = aws.String(_cognitosyncIdentityPoolId)
	}
	if len(_cognitosyncCognitoStreams) > 0 {
		if err := assignInputField(input, "CognitoStreams", _cognitosyncCognitoStreams); err != nil {
			log.Errorf("invalid --cognito-streams: %s", err.Error())
			return
		}
	}
	if len(_cognitosyncPushSync) > 0 {
		if err := assignInputField(input, "PushSync", _cognitosyncPushSync); err != nil {
			log.Errorf("invalid --push-sync: %s", err.Error())
			return
		}
	}

	if resp, err := client.SetIdentityPoolConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

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
func cognitosync_SubscribeToDataset(cfg aws.Config, client *cognitosync.Client) {
	input := &cognitosync.SubscribeToDatasetInput{
		// DatasetName: *string, // Required
		// DeviceId: *string, // Required
		// IdentityId: *string, // Required
		// IdentityPoolId: *string, // Required
	}

	if len(_cognitosyncDatasetName) > 0 {
		input.DatasetName = aws.String(_cognitosyncDatasetName)
	}
	if len(_cognitosyncDeviceId) > 0 {
		input.DeviceId = aws.String(_cognitosyncDeviceId)
	}
	if len(_cognitosyncIdentityId) > 0 {
		input.IdentityId = aws.String(_cognitosyncIdentityId)
	}
	if len(_cognitosyncIdentityPoolId) > 0 {
		input.IdentityPoolId = aws.String(_cognitosyncIdentityPoolId)
	}

	if resp, err := client.SubscribeToDataset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

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
func cognitosync_UnsubscribeFromDataset(cfg aws.Config, client *cognitosync.Client) {
	input := &cognitosync.UnsubscribeFromDatasetInput{
		// DatasetName: *string, // Required
		// DeviceId: *string, // Required
		// IdentityId: *string, // Required
		// IdentityPoolId: *string, // Required
	}

	if len(_cognitosyncDatasetName) > 0 {
		input.DatasetName = aws.String(_cognitosyncDatasetName)
	}
	if len(_cognitosyncDeviceId) > 0 {
		input.DeviceId = aws.String(_cognitosyncDeviceId)
	}
	if len(_cognitosyncIdentityId) > 0 {
		input.IdentityId = aws.String(_cognitosyncIdentityId)
	}
	if len(_cognitosyncIdentityPoolId) > 0 {
		input.IdentityPoolId = aws.String(_cognitosyncIdentityPoolId)
	}

	if resp, err := client.UnsubscribeFromDataset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Posts updates to records and adds and deletes records for a dataset and user.
// The sync count in the record patch is your last known sync count for that
// record. The server will reject an UpdateRecords request with a
// ResourceConflictException if you try to patch a record with a new value but a
// stale sync count.
//
// For example, if the sync count on the server is 5 for a key called highScore
// and you try and submit a new highScore with sync count of 4, the request will be
// rejected. To obtain the current sync count for a record, call ListRecords. On a
// successful update of the record, the response returns the new sync count for
// that record. You should present that sync count the next time you try to update
// that same record. When the record does not exist, specify the sync count as 0.
//
// This API can be called with temporary user credentials provided by Cognito
// Identity or with developer credentials.
func cognitosync_UpdateRecords(cfg aws.Config, client *cognitosync.Client) {
	input := &cognitosync.UpdateRecordsInput{
		// DatasetName: *string, // Required
		// IdentityId: *string, // Required
		// IdentityPoolId: *string, // Required
		// SyncSessionToken: *string, // Required
	}

	if len(_cognitosyncDatasetName) > 0 {
		input.DatasetName = aws.String(_cognitosyncDatasetName)
	}
	if len(_cognitosyncIdentityId) > 0 {
		input.IdentityId = aws.String(_cognitosyncIdentityId)
	}
	if len(_cognitosyncIdentityPoolId) > 0 {
		input.IdentityPoolId = aws.String(_cognitosyncIdentityPoolId)
	}
	if len(_cognitosyncSyncSessionToken) > 0 {
		input.SyncSessionToken = aws.String(_cognitosyncSyncSessionToken)
	}
	if len(_cognitosyncClientContext) > 0 {
		input.ClientContext = aws.String(_cognitosyncClientContext)
	}
	if len(_cognitosyncDeviceId) > 0 {
		input.DeviceId = aws.String(_cognitosyncDeviceId)
	}
	if len(_cognitosyncRecordPatches) > 0 {
		if err := assignInputField(input, "RecordPatches", _cognitosyncRecordPatches); err != nil {
			log.Errorf("invalid --record-patches: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateRecords(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_cognitosyncCmd)
	_cognitosyncCmd.Flags().SortFlags = false

	_cognitosyncCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_cognitosyncCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_cognitosyncCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_cognitosyncCmd.Flags().StringVarP(&_cognitosyncClientContext, "client-context", "", "", "Client Context")
	_cognitosyncCmd.Flags().StringVarP(&_cognitosyncCognitoStreams, "cognito-streams", "", "", "Cognito Streams")
	_cognitosyncCmd.Flags().StringVarP(&_cognitosyncDatasetName, "dataset-name", "", "", "Dataset Name")
	_cognitosyncCmd.Flags().StringVarP(&_cognitosyncDeviceId, "device-id", "", "", "Device ID")
	_cognitosyncCmd.Flags().StringVarP(&_cognitosyncEvents, "events", "", "", "Events")
	_cognitosyncCmd.Flags().StringVarP(&_cognitosyncIdentityId, "identity-id", "", "", "Identity ID")
	_cognitosyncCmd.Flags().StringVarP(&_cognitosyncIdentityPoolId, "identity-pool-id", "", "", "Identity Pool ID")
	_cognitosyncCmd.Flags().StringVarP(&_cognitosyncLastSyncCount, "last-sync-count", "", "", "Last Sync Count")
	_cognitosyncCmd.Flags().StringVarP(&_cognitosyncMaxResults, "max-results", "", "", "Max Results")
	_cognitosyncCmd.Flags().StringVarP(&_cognitosyncNextToken, "next-token", "", "", "Next Token")
	_cognitosyncCmd.Flags().StringVarP(&_cognitosyncPlatform, "platform", "", "", "Platform")
	_cognitosyncCmd.Flags().StringVarP(&_cognitosyncPushSync, "push-sync", "", "", "Push Sync")
	_cognitosyncCmd.Flags().StringVarP(&_cognitosyncRecordPatches, "record-patches", "", "", "Record Patches")
	_cognitosyncCmd.Flags().StringVarP(&_cognitosyncSyncSessionToken, "sync-session-token", "", "", "Sync Session Token")
	_cognitosyncCmd.Flags().StringVarP(&_cognitosyncToken, "token", "", "", "Token")

	_cognitosyncCmd.Flags().BoolVarP(&_cognitosyncBulkPublish, "bulk-publish", "", false, "Bulk Publish")
	_cognitosyncCmd.Flags().BoolVarP(&_cognitosyncDeleteDataset, "delete-dataset", "", false, "Delete Dataset")
	_cognitosyncCmd.Flags().BoolVarP(&_cognitosyncDescribeDataset, "describe-dataset", "", false, "Describe Dataset")
	_cognitosyncCmd.Flags().BoolVarP(&_cognitosyncDescribeIdentityPoolUsage, "describe-identity-pool-usage", "", false, "Describe Identity Pool Usage")
	_cognitosyncCmd.Flags().BoolVarP(&_cognitosyncDescribeIdentityUsage, "describe-identity-usage", "", false, "Describe Identity Usage")
	_cognitosyncCmd.Flags().BoolVarP(&_cognitosyncGetBulkPublishDetails, "get-bulk-publish-details", "", false, "Get Bulk Publish Details")
	_cognitosyncCmd.Flags().BoolVarP(&_cognitosyncGetCognitoEvents, "get-cognito-events", "", false, "Get Cognito Events")
	_cognitosyncCmd.Flags().BoolVarP(&_cognitosyncGetIdentityPoolConfiguration, "get-identity-pool-configuration", "", false, "Get Identity Pool Configuration")
	_cognitosyncCmd.Flags().BoolVarP(&_cognitosyncListDatasets, "list-datasets", "", false, "List Datasets")
	_cognitosyncCmd.Flags().BoolVarP(&_cognitosyncListIdentityPoolUsage, "list-identity-pool-usage", "", false, "List Identity Pool Usage")
	_cognitosyncCmd.Flags().BoolVarP(&_cognitosyncListRecords, "list-records", "", false, "List Records")
	_cognitosyncCmd.Flags().BoolVarP(&_cognitosyncRegisterDevice, "register-device", "", false, "Register Device")
	_cognitosyncCmd.Flags().BoolVarP(&_cognitosyncSetCognitoEvents, "set-cognito-events", "", false, "Set Cognito Events")
	_cognitosyncCmd.Flags().BoolVarP(&_cognitosyncSetIdentityPoolConfiguration, "set-identity-pool-configuration", "", false, "Set Identity Pool Configuration")
	_cognitosyncCmd.Flags().BoolVarP(&_cognitosyncSubscribeToDataset, "subscribe-to-dataset", "", false, "Subscribe To Dataset")
	_cognitosyncCmd.Flags().BoolVarP(&_cognitosyncUnsubscribeFromDataset, "unsubscribe-from-dataset", "", false, "Unsubscribe From Dataset")
	_cognitosyncCmd.Flags().BoolVarP(&_cognitosyncUpdateRecords, "update-records", "", false, "Update Records")

}
