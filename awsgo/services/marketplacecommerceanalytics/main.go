package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/marketplacecommerceanalytics/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"generate-data-set", "start-support-data-export"},
		OperationSet: map[string]bool{"generate-data-set": true, "start-support-data-export": true},
		OperationInputs: map[string][]string{
			"generate-data-set":         {"CustomerDefinedValues", "DataSetPublicationDate", "DataSetType", "DestinationS3BucketName", "DestinationS3Prefix", "RoleNameArn", "SnsTopicArn"},
			"start-support-data-export": {"CustomerDefinedValues", "DataSetType", "DestinationS3BucketName", "DestinationS3Prefix", "FromDate", "RoleNameArn", "SnsTopicArn"},
		},
		OperationInputTypes: map[string]map[string]string{
			"generate-data-set":         {"CustomerDefinedValues": "map[string]string", "DataSetPublicationDate": "*time.Time", "DataSetType": "types.DataSetType", "DestinationS3BucketName": "*string", "DestinationS3Prefix": "*string", "RoleNameArn": "*string", "SnsTopicArn": "*string"},
			"start-support-data-export": {"CustomerDefinedValues": "map[string]string", "DataSetType": "types.SupportDataSetType", "DestinationS3BucketName": "*string", "DestinationS3Prefix": "*string", "FromDate": "*time.Time", "RoleNameArn": "*string", "SnsTopicArn": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"generate-data-set":         {"DataSetPublicationDate", "DataSetType", "DestinationS3BucketName", "RoleNameArn", "SnsTopicArn"},
			"start-support-data-export": {"DataSetType", "DestinationS3BucketName", "FromDate", "RoleNameArn", "SnsTopicArn"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("marketplacecommerceanalytics", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
