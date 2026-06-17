package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodbstreams"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// dynamodbstreamsCmd represents the dynamodbstreams command
var _dynamodbstreamsCmd = &cobra.Command{
	Use:   "dynamodbstreams",
	Short: "AWS dynamodbstreams CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := dynamodbstreams.NewFromConfig(cfg)
		if _dynamodbstreamsDescribeStream {
			dynamodbstreams_DescribeStream(cfg, client)
			return
		}
		if _dynamodbstreamsGetRecords {
			dynamodbstreams_GetRecords(cfg, client)
			return
		}
		if _dynamodbstreamsGetShardIterator {
			dynamodbstreams_GetShardIterator(cfg, client)
			return
		}
		if _dynamodbstreamsListStreams {
			dynamodbstreams_ListStreams(cfg, client)
			return
		}

	},
}

var (
	_dynamodbstreamsDescribeStream   bool
	_dynamodbstreamsGetRecords       bool
	_dynamodbstreamsGetShardIterator bool
	_dynamodbstreamsListStreams      bool

	_dynamodbstreamsExclusiveStartShardId   string
	_dynamodbstreamsExclusiveStartStreamArn string
	_dynamodbstreamsLimit                   string
	_dynamodbstreamsSequenceNumber          string
	_dynamodbstreamsShardFilter             string
	_dynamodbstreamsShardId                 string
	_dynamodbstreamsShardIterator           string
	_dynamodbstreamsShardIteratorType       string
	_dynamodbstreamsStreamArn               string
	_dynamodbstreamsTableName               string
)

// Returns information about a stream, including the current status of the stream,
// its Amazon Resource Name (ARN), the composition of its shards, and its
// corresponding DynamoDB table.
//
// You can call DescribeStream at a maximum rate of 10 times per second.
//
// Each shard in the stream has a SequenceNumberRange associated with it. If the
// SequenceNumberRange has a StartingSequenceNumber but no EndingSequenceNumber ,
// then the shard is still open (able to receive more stream records). If both
// StartingSequenceNumber and EndingSequenceNumber are present, then that shard is
// closed and can no longer receive more data.
func dynamodbstreams_DescribeStream(cfg aws.Config, client *dynamodbstreams.Client) {
	input := &dynamodbstreams.DescribeStreamInput{
		// StreamArn: *string, // Required
	}

	if len(_dynamodbstreamsStreamArn) > 0 {
		input.StreamArn = aws.String(_dynamodbstreamsStreamArn)
	}
	if len(_dynamodbstreamsExclusiveStartShardId) > 0 {
		input.ExclusiveStartShardId = aws.String(_dynamodbstreamsExclusiveStartShardId)
	}
	if len(_dynamodbstreamsLimit) > 0 {
		if err := assignInputField(input, "Limit", _dynamodbstreamsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_dynamodbstreamsShardFilter) > 0 {
		if err := assignInputField(input, "ShardFilter", _dynamodbstreamsShardFilter); err != nil {
			log.Errorf("invalid --shard-filter: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeStream(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the stream records from a given shard.
// Specify a shard iterator using the ShardIterator parameter. The shard iterator
// specifies the position in the shard from which you want to start reading stream
// records sequentially. If there are no stream records available in the portion of
// the shard that the iterator points to, GetRecords returns an empty list. Note
// that it might take multiple calls to get to a portion of the shard that contains
// stream records.
//
// GetRecords can retrieve a maximum of 1 MB of data or 1000 stream records,
// whichever comes first.
func dynamodbstreams_GetRecords(cfg aws.Config, client *dynamodbstreams.Client) {
	input := &dynamodbstreams.GetRecordsInput{
		// ShardIterator: *string, // Required
	}

	if len(_dynamodbstreamsShardIterator) > 0 {
		input.ShardIterator = aws.String(_dynamodbstreamsShardIterator)
	}
	if len(_dynamodbstreamsLimit) > 0 {
		if err := assignInputField(input, "Limit", _dynamodbstreamsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetRecords(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a shard iterator. A shard iterator provides information about how to
// retrieve the stream records from within a shard. Use the shard iterator in a
// subsequent GetRecords request to read the stream records from the shard.
//
// A shard iterator expires 15 minutes after it is returned to the requester.
func dynamodbstreams_GetShardIterator(cfg aws.Config, client *dynamodbstreams.Client) {
	input := &dynamodbstreams.GetShardIteratorInput{
		// ShardId: *string, // Required
		// ShardIteratorType: types.ShardIteratorType, // Required
		// StreamArn: *string, // Required
	}

	if len(_dynamodbstreamsShardId) > 0 {
		input.ShardId = aws.String(_dynamodbstreamsShardId)
	}
	if len(_dynamodbstreamsShardIteratorType) > 0 {
		if err := assignInputField(input, "ShardIteratorType", _dynamodbstreamsShardIteratorType); err != nil {
			log.Errorf("invalid --shard-iterator-type: %s", err.Error())
			return
		}
	}
	if len(_dynamodbstreamsStreamArn) > 0 {
		input.StreamArn = aws.String(_dynamodbstreamsStreamArn)
	}
	if len(_dynamodbstreamsSequenceNumber) > 0 {
		input.SequenceNumber = aws.String(_dynamodbstreamsSequenceNumber)
	}

	if resp, err := client.GetShardIterator(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns an array of stream ARNs associated with the current account and
// endpoint. If the TableName parameter is present, then ListStreams will return
// only the streams ARNs for that table.
//
// You can call ListStreams at a maximum rate of 5 times per second.
func dynamodbstreams_ListStreams(cfg aws.Config, client *dynamodbstreams.Client) {
	input := &dynamodbstreams.ListStreamsInput{}

	if len(_dynamodbstreamsExclusiveStartStreamArn) > 0 {
		input.ExclusiveStartStreamArn = aws.String(_dynamodbstreamsExclusiveStartStreamArn)
	}
	if len(_dynamodbstreamsLimit) > 0 {
		if err := assignInputField(input, "Limit", _dynamodbstreamsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_dynamodbstreamsTableName) > 0 {
		input.TableName = aws.String(_dynamodbstreamsTableName)
	}

	if resp, err := client.ListStreams(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_dynamodbstreamsCmd)
	_dynamodbstreamsCmd.Flags().SortFlags = false

	_dynamodbstreamsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_dynamodbstreamsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_dynamodbstreamsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_dynamodbstreamsCmd.Flags().StringVarP(&_dynamodbstreamsExclusiveStartShardId, "exclusive-start-shard-id", "", "", "Exclusive Start Shard ID")
	_dynamodbstreamsCmd.Flags().StringVarP(&_dynamodbstreamsExclusiveStartStreamArn, "exclusive-start-stream-arn", "", "", "Exclusive Start Stream ARN")
	_dynamodbstreamsCmd.Flags().StringVarP(&_dynamodbstreamsLimit, "limit", "", "", "Limit")
	_dynamodbstreamsCmd.Flags().StringVarP(&_dynamodbstreamsSequenceNumber, "sequence-number", "", "", "Sequence Number")
	_dynamodbstreamsCmd.Flags().StringVarP(&_dynamodbstreamsShardFilter, "shard-filter", "", "", "Shard Filter")
	_dynamodbstreamsCmd.Flags().StringVarP(&_dynamodbstreamsShardId, "shard-id", "", "", "Shard ID")
	_dynamodbstreamsCmd.Flags().StringVarP(&_dynamodbstreamsShardIterator, "shard-iterator", "", "", "Shard Iterator")
	_dynamodbstreamsCmd.Flags().StringVarP(&_dynamodbstreamsShardIteratorType, "shard-iterator-type", "", "", "Shard Iterator Type")
	_dynamodbstreamsCmd.Flags().StringVarP(&_dynamodbstreamsStreamArn, "stream-arn", "", "", "Stream ARN")
	_dynamodbstreamsCmd.Flags().StringVarP(&_dynamodbstreamsTableName, "table-name", "", "", "Table Name")

	_dynamodbstreamsCmd.Flags().BoolVarP(&_dynamodbstreamsDescribeStream, "describe-stream", "", false, "Describe Stream")
	_dynamodbstreamsCmd.Flags().BoolVarP(&_dynamodbstreamsGetRecords, "get-records", "", false, "Get Records")
	_dynamodbstreamsCmd.Flags().BoolVarP(&_dynamodbstreamsGetShardIterator, "get-shard-iterator", "", false, "Get Shard Iterator")
	_dynamodbstreamsCmd.Flags().BoolVarP(&_dynamodbstreamsListStreams, "list-streams", "", false, "List Streams")

}
