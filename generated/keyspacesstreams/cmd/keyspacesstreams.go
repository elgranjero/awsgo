package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/keyspacesstreams"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// keyspacesstreamsCmd represents the keyspacesstreams command
var _keyspacesstreamsCmd = &cobra.Command{
	Use:   "keyspacesstreams",
	Short: "AWS keyspacesstreams CLI",
	Run: func(cmd *cobra.Command, args []string) {
		_awsOutput = resolveAWSOutput(_awsProfile, cmd.Flags().Changed("output"))
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := keyspacesstreams.NewFromConfig(cfg)
		if _keyspacesstreamsGetRecords {
			keyspacesstreams_GetRecords(cfg, client)
			return
		}
		if _keyspacesstreamsGetShardIterator {
			keyspacesstreams_GetShardIterator(cfg, client)
			return
		}
		if _keyspacesstreamsGetStream {
			keyspacesstreams_GetStream(cfg, client)
			return
		}
		if _keyspacesstreamsListStreams {
			keyspacesstreams_ListStreams(cfg, client)
			return
		}

	},
}

var (
	_keyspacesstreamsGetRecords       bool
	_keyspacesstreamsGetShardIterator bool
	_keyspacesstreamsGetStream        bool
	_keyspacesstreamsListStreams      bool

	_keyspacesstreamsKeyspaceName      string
	_keyspacesstreamsMaxResults        string
	_keyspacesstreamsNextToken         string
	_keyspacesstreamsSequenceNumber    string
	_keyspacesstreamsShardFilter       string
	_keyspacesstreamsShardId           string
	_keyspacesstreamsShardIterator     string
	_keyspacesstreamsShardIteratorType string
	_keyspacesstreamsStreamArn         string
	_keyspacesstreamsTableName         string
)

// Retrieves data records from a specified shard in an Amazon Keyspaces data
// stream. This operation returns a collection of data records from the shard,
// including the primary key columns and information about modifications made to
// the captured table data. Each record represents a single data modification in
// the Amazon Keyspaces table and includes metadata about when the change occurred.
func keyspacesstreams_GetRecords(cfg aws.Config, client *keyspacesstreams.Client) {
	input := &keyspacesstreams.GetRecordsInput{
		// ShardIterator: *string, // Required
	}

	if len(_keyspacesstreamsShardIterator) > 0 {
		input.ShardIterator = aws.String(_keyspacesstreamsShardIterator)
	}
	if len(_keyspacesstreamsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _keyspacesstreamsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
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

// Returns a shard iterator that serves as a bookmark for reading data from a
// specific position in an Amazon Keyspaces data stream's shard. The shard iterator
// specifies the shard position from which to start reading data records
// sequentially. You can specify whether to begin reading at the latest record, the
// oldest record, or at a particular sequence number within the shard.
func keyspacesstreams_GetShardIterator(cfg aws.Config, client *keyspacesstreams.Client) {
	input := &keyspacesstreams.GetShardIteratorInput{
		// ShardId: *string, // Required
		// ShardIteratorType: types.ShardIteratorType, // Required
		// StreamArn: *string, // Required
	}

	if len(_keyspacesstreamsShardId) > 0 {
		input.ShardId = aws.String(_keyspacesstreamsShardId)
	}
	if len(_keyspacesstreamsShardIteratorType) > 0 {
		if err := assignInputField(input, "ShardIteratorType", _keyspacesstreamsShardIteratorType); err != nil {
			log.Errorf("invalid --shard-iterator-type: %s", err.Error())
			return
		}
	}
	if len(_keyspacesstreamsStreamArn) > 0 {
		input.StreamArn = aws.String(_keyspacesstreamsStreamArn)
	}
	if len(_keyspacesstreamsSequenceNumber) > 0 {
		input.SequenceNumber = aws.String(_keyspacesstreamsSequenceNumber)
	}

	if resp, err := client.GetShardIterator(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns detailed information about a specific data capture stream for an Amazon
// Keyspaces table. The information includes the stream's Amazon Resource Name
// (ARN), creation time, current status, retention period, shard composition, and
// associated table details. This operation helps you monitor and manage the
// configuration of your Amazon Keyspaces data streams.
func keyspacesstreams_GetStream(cfg aws.Config, client *keyspacesstreams.Client) {
	input := &keyspacesstreams.GetStreamInput{
		// StreamArn: *string, // Required
	}

	if len(_keyspacesstreamsStreamArn) > 0 {
		input.StreamArn = aws.String(_keyspacesstreamsStreamArn)
	}
	if len(_keyspacesstreamsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _keyspacesstreamsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_keyspacesstreamsNextToken) > 0 {
		input.NextToken = aws.String(_keyspacesstreamsNextToken)
	}
	if len(_keyspacesstreamsShardFilter) > 0 {
		if err := assignInputField(input, "ShardFilter", _keyspacesstreamsShardFilter); err != nil {
			log.Errorf("invalid --shard-filter: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetStream(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*keyspacesstreams.GetStreamOutput
	p := keyspacesstreams.NewGetStreamPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns a list of all data capture streams associated with your Amazon
// Keyspaces account or for a specific keyspace or table. The response includes
// information such as stream ARNs, table associations, creation timestamps, and
// current status. This operation helps you discover and manage all active data
// streams in your Amazon Keyspaces environment.
func keyspacesstreams_ListStreams(cfg aws.Config, client *keyspacesstreams.Client) {
	input := &keyspacesstreams.ListStreamsInput{}

	if len(_keyspacesstreamsKeyspaceName) > 0 {
		input.KeyspaceName = aws.String(_keyspacesstreamsKeyspaceName)
	}
	if len(_keyspacesstreamsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _keyspacesstreamsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_keyspacesstreamsNextToken) > 0 {
		input.NextToken = aws.String(_keyspacesstreamsNextToken)
	}
	if len(_keyspacesstreamsTableName) > 0 {
		input.TableName = aws.String(_keyspacesstreamsTableName)
	}

	if disablePaginator() {
		if resp, err := client.ListStreams(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*keyspacesstreams.ListStreamsOutput
	p := keyspacesstreams.NewListStreamsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

func init() {
	_rootCmd.AddCommand(_keyspacesstreamsCmd)
	_keyspacesstreamsCmd.Flags().SortFlags = false

	_keyspacesstreamsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_keyspacesstreamsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_keyspacesstreamsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_keyspacesstreamsCmd.Flags().StringVarP(&_keyspacesstreamsKeyspaceName, "keyspace-name", "", "", "Keyspace Name")
	_keyspacesstreamsCmd.Flags().StringVarP(&_keyspacesstreamsMaxResults, "max-results", "", "", "Max Results")
	_keyspacesstreamsCmd.Flags().StringVarP(&_keyspacesstreamsNextToken, "next-token", "", "", "Next Token")
	_keyspacesstreamsCmd.Flags().StringVarP(&_keyspacesstreamsSequenceNumber, "sequence-number", "", "", "Sequence Number")
	_keyspacesstreamsCmd.Flags().StringVarP(&_keyspacesstreamsShardFilter, "shard-filter", "", "", "Shard Filter")
	_keyspacesstreamsCmd.Flags().StringVarP(&_keyspacesstreamsShardId, "shard-id", "", "", "Shard ID")
	_keyspacesstreamsCmd.Flags().StringVarP(&_keyspacesstreamsShardIterator, "shard-iterator", "", "", "Shard Iterator")
	_keyspacesstreamsCmd.Flags().StringVarP(&_keyspacesstreamsShardIteratorType, "shard-iterator-type", "", "", "Shard Iterator Type")
	_keyspacesstreamsCmd.Flags().StringVarP(&_keyspacesstreamsStreamArn, "stream-arn", "", "", "Stream ARN")
	_keyspacesstreamsCmd.Flags().StringVarP(&_keyspacesstreamsTableName, "table-name", "", "", "Table Name")

	_keyspacesstreamsCmd.Flags().BoolVarP(&_keyspacesstreamsGetRecords, "get-records", "", false, "Get Records")
	_keyspacesstreamsCmd.Flags().BoolVarP(&_keyspacesstreamsGetShardIterator, "get-shard-iterator", "", false, "Get Shard Iterator")
	_keyspacesstreamsCmd.Flags().BoolVarP(&_keyspacesstreamsGetStream, "get-stream", "", false, "Get Stream")
	_keyspacesstreamsCmd.Flags().BoolVarP(&_keyspacesstreamsListStreams, "list-streams", "", false, "List Streams")

}
