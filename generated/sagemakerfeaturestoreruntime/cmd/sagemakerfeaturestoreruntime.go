package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sagemakerfeaturestoreruntime"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// sagemakerfeaturestoreruntimeCmd represents the sagemakerfeaturestoreruntime command
var _sagemakerfeaturestoreruntimeCmd = &cobra.Command{
	Use:   "sagemakerfeaturestoreruntime",
	Short: "AWS sagemakerfeaturestoreruntime CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := sagemakerfeaturestoreruntime.NewFromConfig(cfg)
		if _sagemakerfeaturestoreruntimeBatchGetRecord {
			sagemakerfeaturestoreruntime_BatchGetRecord(cfg, client)
			return
		}
		if _sagemakerfeaturestoreruntimeDeleteRecord {
			sagemakerfeaturestoreruntime_DeleteRecord(cfg, client)
			return
		}
		if _sagemakerfeaturestoreruntimeGetRecord {
			sagemakerfeaturestoreruntime_GetRecord(cfg, client)
			return
		}
		if _sagemakerfeaturestoreruntimePutRecord {
			sagemakerfeaturestoreruntime_PutRecord(cfg, client)
			return
		}

	},
}

var (
	_sagemakerfeaturestoreruntimeBatchGetRecord bool
	_sagemakerfeaturestoreruntimeDeleteRecord   bool
	_sagemakerfeaturestoreruntimeGetRecord      bool
	_sagemakerfeaturestoreruntimePutRecord      bool

	_sagemakerfeaturestoreruntimeDeletionMode                  string
	_sagemakerfeaturestoreruntimeEventTime                     string
	_sagemakerfeaturestoreruntimeExpirationTimeResponse        string
	_sagemakerfeaturestoreruntimeFeatureGroupName              string
	_sagemakerfeaturestoreruntimeFeatureNames                  []string
	_sagemakerfeaturestoreruntimeIdentifiers                   string
	_sagemakerfeaturestoreruntimeRecord                        string
	_sagemakerfeaturestoreruntimeRecordIdentifierValueAsString string
	_sagemakerfeaturestoreruntimeTargetStores                  string
	_sagemakerfeaturestoreruntimeTtlDuration                   string
)

// Retrieves a batch of Records from a FeatureGroup .
func sagemakerfeaturestoreruntime_BatchGetRecord(cfg aws.Config, client *sagemakerfeaturestoreruntime.Client) {
	input := &sagemakerfeaturestoreruntime.BatchGetRecordInput{
		// Identifiers: []types.BatchGetRecordIdentifier, // Required
	}

	if len(_sagemakerfeaturestoreruntimeIdentifiers) > 0 {
		if err := assignInputField(input, "Identifiers", _sagemakerfeaturestoreruntimeIdentifiers); err != nil {
			log.Errorf("invalid --identifiers: %s", err.Error())
			return
		}
	}
	if len(_sagemakerfeaturestoreruntimeExpirationTimeResponse) > 0 {
		if err := assignInputField(input, "ExpirationTimeResponse", _sagemakerfeaturestoreruntimeExpirationTimeResponse); err != nil {
			log.Errorf("invalid --expiration-time-response: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchGetRecord(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Record from a FeatureGroup in the OnlineStore . Feature Store supports
// both SoftDelete and HardDelete . For SoftDelete (default), feature columns are
// set to null and the record is no longer retrievable by GetRecord or
// BatchGetRecord . For HardDelete , the complete Record is removed from the
// OnlineStore . In both cases, Feature Store appends the deleted record marker to
// the OfflineStore . The deleted record marker is a record with the same
// RecordIdentifer as the original, but with is_deleted value set to True ,
// EventTime set to the delete input EventTime , and other feature values set to
// null .
//
// Note that the EventTime specified in DeleteRecord should be set later than the
// EventTime of the existing record in the OnlineStore for that RecordIdentifer .
// If it is not, the deletion does not occur:
//
// - For SoftDelete , the existing (not deleted) record remains in the
// OnlineStore , though the delete record marker is still written to the
// OfflineStore .
//
// - HardDelete returns EventTime : 400 ValidationException to indicate that the
// delete operation failed. No delete record marker is written to the
// OfflineStore .
//
// When a record is deleted from the OnlineStore , the deleted record marker is
// appended to the OfflineStore . If you have the Iceberg table format enabled for
// your OfflineStore , you can remove all history of a record from the OfflineStore
// using Amazon Athena or Apache Spark. For information on how to hard delete a
// record from the OfflineStore with the Iceberg table format enabled, see [Delete records from the offline store].
//
// [Delete records from the offline store]: https://docs.aws.amazon.com/sagemaker/latest/dg/feature-store-delete-records-offline-store.html#feature-store-delete-records-offline-store
func sagemakerfeaturestoreruntime_DeleteRecord(cfg aws.Config, client *sagemakerfeaturestoreruntime.Client) {
	input := &sagemakerfeaturestoreruntime.DeleteRecordInput{
		// EventTime: *string, // Required
		// FeatureGroupName: *string, // Required
		// RecordIdentifierValueAsString: *string, // Required
	}

	if len(_sagemakerfeaturestoreruntimeEventTime) > 0 {
		input.EventTime = aws.String(_sagemakerfeaturestoreruntimeEventTime)
	}
	if len(_sagemakerfeaturestoreruntimeFeatureGroupName) > 0 {
		input.FeatureGroupName = aws.String(_sagemakerfeaturestoreruntimeFeatureGroupName)
	}
	if len(_sagemakerfeaturestoreruntimeRecordIdentifierValueAsString) > 0 {
		input.RecordIdentifierValueAsString = aws.String(_sagemakerfeaturestoreruntimeRecordIdentifierValueAsString)
	}
	if len(_sagemakerfeaturestoreruntimeDeletionMode) > 0 {
		if err := assignInputField(input, "DeletionMode", _sagemakerfeaturestoreruntimeDeletionMode); err != nil {
			log.Errorf("invalid --deletion-mode: %s", err.Error())
			return
		}
	}
	if len(_sagemakerfeaturestoreruntimeTargetStores) > 0 {
		if err := assignInputField(input, "TargetStores", _sagemakerfeaturestoreruntimeTargetStores); err != nil {
			log.Errorf("invalid --target-stores: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteRecord(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use for OnlineStore serving from a FeatureStore . Only the latest records stored
// in the OnlineStore can be retrieved. If no Record with RecordIdentifierValue is
// found, then an empty result is returned.
func sagemakerfeaturestoreruntime_GetRecord(cfg aws.Config, client *sagemakerfeaturestoreruntime.Client) {
	input := &sagemakerfeaturestoreruntime.GetRecordInput{
		// FeatureGroupName: *string, // Required
		// RecordIdentifierValueAsString: *string, // Required
	}

	if len(_sagemakerfeaturestoreruntimeFeatureGroupName) > 0 {
		input.FeatureGroupName = aws.String(_sagemakerfeaturestoreruntimeFeatureGroupName)
	}
	if len(_sagemakerfeaturestoreruntimeRecordIdentifierValueAsString) > 0 {
		input.RecordIdentifierValueAsString = aws.String(_sagemakerfeaturestoreruntimeRecordIdentifierValueAsString)
	}
	if len(_sagemakerfeaturestoreruntimeExpirationTimeResponse) > 0 {
		if err := assignInputField(input, "ExpirationTimeResponse", _sagemakerfeaturestoreruntimeExpirationTimeResponse); err != nil {
			log.Errorf("invalid --expiration-time-response: %s", err.Error())
			return
		}
	}
	if len(_sagemakerfeaturestoreruntimeFeatureNames) > 0 {
		input.FeatureNames = append([]string(nil), _sagemakerfeaturestoreruntimeFeatureNames...)
	}

	if resp, err := client.GetRecord(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The PutRecord API is used to ingest a list of Records into your feature group.
// If a new record’s EventTime is greater, the new record is written to both the
// OnlineStore and OfflineStore . Otherwise, the record is a historic record and it
// is written only to the OfflineStore .
//
// You can specify the ingestion to be applied to the OnlineStore , OfflineStore ,
// or both by using the TargetStores request parameter.
//
// You can set the ingested record to expire at a given time to live (TTL)
// duration after the record’s event time, ExpiresAt = EventTime + TtlDuration , by
// specifying the TtlDuration parameter. A record level TtlDuration is set when
// specifying the TtlDuration parameter using the PutRecord API call. If the input
// TtlDuration is null or unspecified, TtlDuration is set to the default feature
// group level TtlDuration . A record level TtlDuration supersedes the group level
// TtlDuration .
func sagemakerfeaturestoreruntime_PutRecord(cfg aws.Config, client *sagemakerfeaturestoreruntime.Client) {
	input := &sagemakerfeaturestoreruntime.PutRecordInput{
		// FeatureGroupName: *string, // Required
		// Record: []types.FeatureValue, // Required
	}

	if len(_sagemakerfeaturestoreruntimeFeatureGroupName) > 0 {
		input.FeatureGroupName = aws.String(_sagemakerfeaturestoreruntimeFeatureGroupName)
	}
	if len(_sagemakerfeaturestoreruntimeRecord) > 0 {
		if err := assignInputField(input, "Record", _sagemakerfeaturestoreruntimeRecord); err != nil {
			log.Errorf("invalid --record: %s", err.Error())
			return
		}
	}
	if len(_sagemakerfeaturestoreruntimeTargetStores) > 0 {
		if err := assignInputField(input, "TargetStores", _sagemakerfeaturestoreruntimeTargetStores); err != nil {
			log.Errorf("invalid --target-stores: %s", err.Error())
			return
		}
	}
	if len(_sagemakerfeaturestoreruntimeTtlDuration) > 0 {
		if err := assignInputField(input, "TtlDuration", _sagemakerfeaturestoreruntimeTtlDuration); err != nil {
			log.Errorf("invalid --ttl-duration: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutRecord(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_sagemakerfeaturestoreruntimeCmd)
	_sagemakerfeaturestoreruntimeCmd.Flags().SortFlags = false

	_sagemakerfeaturestoreruntimeCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_sagemakerfeaturestoreruntimeCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_sagemakerfeaturestoreruntimeCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_sagemakerfeaturestoreruntimeCmd.Flags().StringVarP(&_sagemakerfeaturestoreruntimeDeletionMode, "deletion-mode", "", "", "Deletion Mode")
	_sagemakerfeaturestoreruntimeCmd.Flags().StringVarP(&_sagemakerfeaturestoreruntimeEventTime, "event-time", "", "", "Event Time")
	_sagemakerfeaturestoreruntimeCmd.Flags().StringVarP(&_sagemakerfeaturestoreruntimeExpirationTimeResponse, "expiration-time-response", "", "", "Expiration Time Response")
	_sagemakerfeaturestoreruntimeCmd.Flags().StringVarP(&_sagemakerfeaturestoreruntimeFeatureGroupName, "feature-group-name", "", "", "Feature Group Name")
	_sagemakerfeaturestoreruntimeCmd.Flags().StringSliceVarP(&_sagemakerfeaturestoreruntimeFeatureNames, "feature-names", "", nil, "Feature Names")
	_sagemakerfeaturestoreruntimeCmd.Flags().StringVarP(&_sagemakerfeaturestoreruntimeIdentifiers, "identifiers", "", "", "Identifiers")
	_sagemakerfeaturestoreruntimeCmd.Flags().StringVarP(&_sagemakerfeaturestoreruntimeRecord, "record", "", "", "Record")
	_sagemakerfeaturestoreruntimeCmd.Flags().StringVarP(&_sagemakerfeaturestoreruntimeRecordIdentifierValueAsString, "record-identifier-value-as-string", "", "", "Record Identifier Value As String")
	_sagemakerfeaturestoreruntimeCmd.Flags().StringVarP(&_sagemakerfeaturestoreruntimeTargetStores, "target-stores", "", "", "Target Stores")
	_sagemakerfeaturestoreruntimeCmd.Flags().StringVarP(&_sagemakerfeaturestoreruntimeTtlDuration, "ttl-duration", "", "", "TTL Duration")

	_sagemakerfeaturestoreruntimeCmd.Flags().BoolVarP(&_sagemakerfeaturestoreruntimeBatchGetRecord, "batch-get-record", "", false, "Batch Get Record")
	_sagemakerfeaturestoreruntimeCmd.Flags().BoolVarP(&_sagemakerfeaturestoreruntimeDeleteRecord, "delete-record", "", false, "Delete Record")
	_sagemakerfeaturestoreruntimeCmd.Flags().BoolVarP(&_sagemakerfeaturestoreruntimeGetRecord, "get-record", "", false, "Get Record")
	_sagemakerfeaturestoreruntimeCmd.Flags().BoolVarP(&_sagemakerfeaturestoreruntimePutRecord, "put-record", "", false, "Put Record")

}
