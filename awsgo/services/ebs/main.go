package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/ebs/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"complete-snapshot", "get-snapshot-block", "list-changed-blocks", "list-snapshot-blocks", "put-snapshot-block", "start-snapshot"},
		OperationSet: map[string]bool{"complete-snapshot": true, "get-snapshot-block": true, "list-changed-blocks": true, "list-snapshot-blocks": true, "put-snapshot-block": true, "start-snapshot": true},
		OperationInputs: map[string][]string{
			"complete-snapshot":    {"ChangedBlocksCount", "Checksum", "ChecksumAggregationMethod", "ChecksumAlgorithm", "SnapshotId"},
			"get-snapshot-block":   {"BlockIndex", "BlockToken", "SnapshotId"},
			"list-changed-blocks":  {"FirstSnapshotId", "MaxResults", "NextToken", "SecondSnapshotId", "StartingBlockIndex"},
			"list-snapshot-blocks": {"MaxResults", "NextToken", "SnapshotId", "StartingBlockIndex"},
			"put-snapshot-block":   {"BlockData", "BlockIndex", "Checksum", "ChecksumAlgorithm", "DataLength", "Progress", "SnapshotId"},
			"start-snapshot":       {"ClientToken", "Description", "Encrypted", "KmsKeyArn", "ParentSnapshotId", "Tags", "Timeout", "VolumeSize"},
		},
		OperationInputTypes: map[string]map[string]string{
			"complete-snapshot":    {"ChangedBlocksCount": "*int32", "Checksum": "*string", "ChecksumAggregationMethod": "types.ChecksumAggregationMethod", "ChecksumAlgorithm": "types.ChecksumAlgorithm", "SnapshotId": "*string"},
			"get-snapshot-block":   {"BlockIndex": "*int32", "BlockToken": "*string", "SnapshotId": "*string"},
			"list-changed-blocks":  {"FirstSnapshotId": "*string", "MaxResults": "*int32", "NextToken": "*string", "SecondSnapshotId": "*string", "StartingBlockIndex": "*int32"},
			"list-snapshot-blocks": {"MaxResults": "*int32", "NextToken": "*string", "SnapshotId": "*string", "StartingBlockIndex": "*int32"},
			"put-snapshot-block":   {"BlockData": "io.Reader", "BlockIndex": "*int32", "Checksum": "*string", "ChecksumAlgorithm": "types.ChecksumAlgorithm", "DataLength": "*int32", "Progress": "*int32", "SnapshotId": "*string"},
			"start-snapshot":       {"ClientToken": "*string", "Description": "*string", "Encrypted": "*bool", "KmsKeyArn": "*string", "ParentSnapshotId": "*string", "Tags": "[]types.Tag", "Timeout": "*int32", "VolumeSize": "*int64"},
		},
		OperationInputRequired: map[string][]string{
			"complete-snapshot":    {"ChangedBlocksCount", "SnapshotId"},
			"get-snapshot-block":   {"BlockIndex", "BlockToken", "SnapshotId"},
			"list-changed-blocks":  {"SecondSnapshotId"},
			"list-snapshot-blocks": {"SnapshotId"},
			"put-snapshot-block":   {"BlockData", "BlockIndex", "Checksum", "ChecksumAlgorithm", "DataLength", "SnapshotId"},
			"start-snapshot":       {"VolumeSize"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("ebs", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
