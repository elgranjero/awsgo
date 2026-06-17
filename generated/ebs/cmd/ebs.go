package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ebs"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ebsCmd represents the ebs command
var _ebsCmd = &cobra.Command{
	Use:   "ebs",
	Short: "AWS ebs CLI",
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
		client := ebs.NewFromConfig(cfg)
		if _ebsCompleteSnapshot {
			ebs_CompleteSnapshot(cfg, client)
			return
		}
		if _ebsGetSnapshotBlock {
			ebs_GetSnapshotBlock(cfg, client)
			return
		}
		if _ebsListChangedBlocks {
			ebs_ListChangedBlocks(cfg, client)
			return
		}
		if _ebsListSnapshotBlocks {
			ebs_ListSnapshotBlocks(cfg, client)
			return
		}
		if _ebsPutSnapshotBlock {
			ebs_PutSnapshotBlock(cfg, client)
			return
		}
		if _ebsStartSnapshot {
			ebs_StartSnapshot(cfg, client)
			return
		}

	},
}

var (
	_ebsCompleteSnapshot   bool
	_ebsGetSnapshotBlock   bool
	_ebsListChangedBlocks  bool
	_ebsListSnapshotBlocks bool
	_ebsPutSnapshotBlock   bool
	_ebsStartSnapshot      bool

	_ebsBlockData                 string
	_ebsBlockIndex                string
	_ebsBlockToken                string
	_ebsChangedBlocksCount        string
	_ebsChecksum                  string
	_ebsChecksumAggregationMethod string
	_ebsChecksumAlgorithm         string
	_ebsClientToken               string
	_ebsDataLength                string
	_ebsDescription               string
	_ebsEncrypted                 string
	_ebsFirstSnapshotId           string
	_ebsKmsKeyArn                 string
	_ebsMaxResults                string
	_ebsNextToken                 string
	_ebsParentSnapshotId          string
	_ebsProgress                  string
	_ebsSecondSnapshotId          string
	_ebsSnapshotId                string
	_ebsStartingBlockIndex        string
	_ebsTags                      string
	_ebsTimeout                   string
	_ebsVolumeSize                string
)

// Seals and completes the snapshot after all of the required blocks of data have
// been written to it. Completing the snapshot changes the status to completed .
// You cannot write new blocks to a snapshot after it has been completed.
//
// You should always retry requests that receive server ( 5xx ) error responses,
// and ThrottlingException and RequestThrottledException client error responses.
// For more information see [Error retries]in the Amazon Elastic Compute Cloud User Guide.
//
// [Error retries]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/error-retries.html
func ebs_CompleteSnapshot(cfg aws.Config, client *ebs.Client) {
	input := &ebs.CompleteSnapshotInput{
		// ChangedBlocksCount: *int32, // Required
		// SnapshotId: *string, // Required
	}

	if len(_ebsChangedBlocksCount) > 0 {
		if err := assignInputField(input, "ChangedBlocksCount", _ebsChangedBlocksCount); err != nil {
			log.Errorf("invalid --changed-blocks-count: %s", err.Error())
			return
		}
	}
	if len(_ebsSnapshotId) > 0 {
		input.SnapshotId = aws.String(_ebsSnapshotId)
	}
	if len(_ebsChecksum) > 0 {
		input.Checksum = aws.String(_ebsChecksum)
	}
	if len(_ebsChecksumAggregationMethod) > 0 {
		if err := assignInputField(input, "ChecksumAggregationMethod", _ebsChecksumAggregationMethod); err != nil {
			log.Errorf("invalid --checksum-aggregation-method: %s", err.Error())
			return
		}
	}
	if len(_ebsChecksumAlgorithm) > 0 {
		if err := assignInputField(input, "ChecksumAlgorithm", _ebsChecksumAlgorithm); err != nil {
			log.Errorf("invalid --checksum-algorithm: %s", err.Error())
			return
		}
	}

	if resp, err := client.CompleteSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the data in a block in an Amazon Elastic Block Store snapshot.
// You should always retry requests that receive server ( 5xx ) error responses,
// and ThrottlingException and RequestThrottledException client error responses.
// For more information see [Error retries]in the Amazon Elastic Compute Cloud User Guide.
//
// [Error retries]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/error-retries.html
func ebs_GetSnapshotBlock(cfg aws.Config, client *ebs.Client) {
	input := &ebs.GetSnapshotBlockInput{
		// BlockIndex: *int32, // Required
		// BlockToken: *string, // Required
		// SnapshotId: *string, // Required
	}

	if len(_ebsBlockIndex) > 0 {
		if err := assignInputField(input, "BlockIndex", _ebsBlockIndex); err != nil {
			log.Errorf("invalid --block-index: %s", err.Error())
			return
		}
	}
	if len(_ebsBlockToken) > 0 {
		input.BlockToken = aws.String(_ebsBlockToken)
	}
	if len(_ebsSnapshotId) > 0 {
		input.SnapshotId = aws.String(_ebsSnapshotId)
	}

	if resp, err := client.GetSnapshotBlock(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the blocks that are different between two Amazon
// Elastic Block Store snapshots of the same volume/snapshot lineage.
//
// You should always retry requests that receive server ( 5xx ) error responses,
// and ThrottlingException and RequestThrottledException client error responses.
// For more information see [Error retries]in the Amazon Elastic Compute Cloud User Guide.
//
// [Error retries]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/error-retries.html
func ebs_ListChangedBlocks(cfg aws.Config, client *ebs.Client) {
	input := &ebs.ListChangedBlocksInput{
		// SecondSnapshotId: *string, // Required
	}

	if len(_ebsSecondSnapshotId) > 0 {
		input.SecondSnapshotId = aws.String(_ebsSecondSnapshotId)
	}
	if len(_ebsFirstSnapshotId) > 0 {
		input.FirstSnapshotId = aws.String(_ebsFirstSnapshotId)
	}
	if len(_ebsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ebsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ebsNextToken) > 0 {
		input.NextToken = aws.String(_ebsNextToken)
	}
	if len(_ebsStartingBlockIndex) > 0 {
		if err := assignInputField(input, "StartingBlockIndex", _ebsStartingBlockIndex); err != nil {
			log.Errorf("invalid --starting-block-index: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListChangedBlocks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ebs.ListChangedBlocksOutput
	p := ebs.NewListChangedBlocksPaginator(client, input)
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

// Returns information about the blocks in an Amazon Elastic Block Store snapshot.
// You should always retry requests that receive server ( 5xx ) error responses,
// and ThrottlingException and RequestThrottledException client error responses.
// For more information see [Error retries]in the Amazon Elastic Compute Cloud User Guide.
//
// [Error retries]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/error-retries.html
func ebs_ListSnapshotBlocks(cfg aws.Config, client *ebs.Client) {
	input := &ebs.ListSnapshotBlocksInput{
		// SnapshotId: *string, // Required
	}

	if len(_ebsSnapshotId) > 0 {
		input.SnapshotId = aws.String(_ebsSnapshotId)
	}
	if len(_ebsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ebsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ebsNextToken) > 0 {
		input.NextToken = aws.String(_ebsNextToken)
	}
	if len(_ebsStartingBlockIndex) > 0 {
		if err := assignInputField(input, "StartingBlockIndex", _ebsStartingBlockIndex); err != nil {
			log.Errorf("invalid --starting-block-index: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListSnapshotBlocks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ebs.ListSnapshotBlocksOutput
	p := ebs.NewListSnapshotBlocksPaginator(client, input)
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

// Writes a block of data to a snapshot. If the specified block contains data, the
// existing data is overwritten. The target snapshot must be in the pending state.
//
// Data written to a snapshot must be aligned with 512-KiB sectors.
//
// You should always retry requests that receive server ( 5xx ) error responses,
// and ThrottlingException and RequestThrottledException client error responses.
// For more information see [Error retries]in the Amazon Elastic Compute Cloud User Guide.
//
// [Error retries]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/error-retries.html
func ebs_PutSnapshotBlock(cfg aws.Config, client *ebs.Client) {
	input := &ebs.PutSnapshotBlockInput{
		// BlockData: io.Reader, // Required
		// BlockIndex: *int32, // Required
		// Checksum: *string, // Required
		// ChecksumAlgorithm: types.ChecksumAlgorithm, // Required
		// DataLength: *int32, // Required
		// SnapshotId: *string, // Required
	}

	if len(_ebsBlockData) > 0 {
		if err := assignInputField(input, "BlockData", _ebsBlockData); err != nil {
			log.Errorf("invalid --block-data: %s", err.Error())
			return
		}
	}
	if len(_ebsBlockIndex) > 0 {
		if err := assignInputField(input, "BlockIndex", _ebsBlockIndex); err != nil {
			log.Errorf("invalid --block-index: %s", err.Error())
			return
		}
	}
	if len(_ebsChecksum) > 0 {
		input.Checksum = aws.String(_ebsChecksum)
	}
	if len(_ebsChecksumAlgorithm) > 0 {
		if err := assignInputField(input, "ChecksumAlgorithm", _ebsChecksumAlgorithm); err != nil {
			log.Errorf("invalid --checksum-algorithm: %s", err.Error())
			return
		}
	}
	if len(_ebsDataLength) > 0 {
		if err := assignInputField(input, "DataLength", _ebsDataLength); err != nil {
			log.Errorf("invalid --data-length: %s", err.Error())
			return
		}
	}
	if len(_ebsSnapshotId) > 0 {
		input.SnapshotId = aws.String(_ebsSnapshotId)
	}
	if len(_ebsProgress) > 0 {
		if err := assignInputField(input, "Progress", _ebsProgress); err != nil {
			log.Errorf("invalid --progress: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutSnapshotBlock(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new Amazon EBS snapshot. The new snapshot enters the pending state
// after the request completes.
//
// After creating the snapshot, use [PutSnapshotBlock] to write blocks of data to the snapshot.
//
// You should always retry requests that receive server ( 5xx ) error responses,
// and ThrottlingException and RequestThrottledException client error responses.
// For more information see [Error retries]in the Amazon Elastic Compute Cloud User Guide.
//
// [PutSnapshotBlock]: https://docs.aws.amazon.com/ebs/latest/APIReference/API_PutSnapshotBlock.html
// [Error retries]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/error-retries.html
func ebs_StartSnapshot(cfg aws.Config, client *ebs.Client) {
	input := &ebs.StartSnapshotInput{
		// VolumeSize: *int64, // Required
	}

	if len(_ebsVolumeSize) > 0 {
		if err := assignInputField(input, "VolumeSize", _ebsVolumeSize); err != nil {
			log.Errorf("invalid --volume-size: %s", err.Error())
			return
		}
	}
	if len(_ebsClientToken) > 0 {
		input.ClientToken = aws.String(_ebsClientToken)
	}
	if len(_ebsDescription) > 0 {
		input.Description = aws.String(_ebsDescription)
	}
	if len(_ebsEncrypted) > 0 {
		if err := assignInputField(input, "Encrypted", _ebsEncrypted); err != nil {
			log.Errorf("invalid --encrypted: %s", err.Error())
			return
		}
	}
	if len(_ebsKmsKeyArn) > 0 {
		input.KmsKeyArn = aws.String(_ebsKmsKeyArn)
	}
	if len(_ebsParentSnapshotId) > 0 {
		input.ParentSnapshotId = aws.String(_ebsParentSnapshotId)
	}
	if len(_ebsTags) > 0 {
		if err := assignInputField(input, "Tags", _ebsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_ebsTimeout) > 0 {
		if err := assignInputField(input, "Timeout", _ebsTimeout); err != nil {
			log.Errorf("invalid --timeout: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_ebsCmd)
	_ebsCmd.Flags().SortFlags = false

	_ebsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_ebsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_ebsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_ebsCmd.Flags().StringVarP(&_ebsBlockData, "block-data", "", "", "Block Data")
	_ebsCmd.Flags().StringVarP(&_ebsBlockIndex, "block-index", "", "", "Block Index")
	_ebsCmd.Flags().StringVarP(&_ebsBlockToken, "block-token", "", "", "Block Token")
	_ebsCmd.Flags().StringVarP(&_ebsChangedBlocksCount, "changed-blocks-count", "", "", "Changed Blocks Count")
	_ebsCmd.Flags().StringVarP(&_ebsChecksum, "checksum", "", "", "Checksum")
	_ebsCmd.Flags().StringVarP(&_ebsChecksumAggregationMethod, "checksum-aggregation-method", "", "", "Checksum Aggregation Method")
	_ebsCmd.Flags().StringVarP(&_ebsChecksumAlgorithm, "checksum-algorithm", "", "", "Checksum Algorithm")
	_ebsCmd.Flags().StringVarP(&_ebsClientToken, "client-token", "", "", "Client Token")
	_ebsCmd.Flags().StringVarP(&_ebsDataLength, "data-length", "", "", "Data Length")
	_ebsCmd.Flags().StringVarP(&_ebsDescription, "description", "", "", "Description")
	_ebsCmd.Flags().StringVarP(&_ebsEncrypted, "encrypted", "", "", "Encrypted")
	_ebsCmd.Flags().StringVarP(&_ebsFirstSnapshotId, "first-snapshot-id", "", "", "First Snapshot ID")
	_ebsCmd.Flags().StringVarP(&_ebsKmsKeyArn, "kms-key-arn", "", "", "KMS Key ARN")
	_ebsCmd.Flags().StringVarP(&_ebsMaxResults, "max-results", "", "", "Max Results")
	_ebsCmd.Flags().StringVarP(&_ebsNextToken, "next-token", "", "", "Next Token")
	_ebsCmd.Flags().StringVarP(&_ebsParentSnapshotId, "parent-snapshot-id", "", "", "Parent Snapshot ID")
	_ebsCmd.Flags().StringVarP(&_ebsProgress, "progress", "", "", "Progress")
	_ebsCmd.Flags().StringVarP(&_ebsSecondSnapshotId, "second-snapshot-id", "", "", "Second Snapshot ID")
	_ebsCmd.Flags().StringVarP(&_ebsSnapshotId, "snapshot-id", "", "", "Snapshot ID")
	_ebsCmd.Flags().StringVarP(&_ebsStartingBlockIndex, "starting-block-index", "", "", "Starting Block Index")
	_ebsCmd.Flags().StringVarP(&_ebsTags, "tags", "", "", "Tags")
	_ebsCmd.Flags().StringVarP(&_ebsTimeout, "timeout", "", "", "Timeout")
	_ebsCmd.Flags().StringVarP(&_ebsVolumeSize, "volume-size", "", "", "Volume Size")

	_ebsCmd.Flags().BoolVarP(&_ebsCompleteSnapshot, "complete-snapshot", "", false, "Complete Snapshot")
	_ebsCmd.Flags().BoolVarP(&_ebsGetSnapshotBlock, "get-snapshot-block", "", false, "Get Snapshot Block")
	_ebsCmd.Flags().BoolVarP(&_ebsListChangedBlocks, "list-changed-blocks", "", false, "List Changed Blocks")
	_ebsCmd.Flags().BoolVarP(&_ebsListSnapshotBlocks, "list-snapshot-blocks", "", false, "List Snapshot Blocks")
	_ebsCmd.Flags().BoolVarP(&_ebsPutSnapshotBlock, "put-snapshot-block", "", false, "Put Snapshot Block")
	_ebsCmd.Flags().BoolVarP(&_ebsStartSnapshot, "start-snapshot", "", false, "Start Snapshot")

}
