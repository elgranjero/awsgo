package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/s3outposts/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-endpoint", "delete-endpoint", "list-endpoints", "list-outposts-with-s3", "list-shared-endpoints"},
		OperationSet: map[string]bool{"create-endpoint": true, "delete-endpoint": true, "list-endpoints": true, "list-outposts-with-s3": true, "list-shared-endpoints": true},
		OperationInputs: map[string][]string{
			"create-endpoint":       {"AccessType", "CustomerOwnedIpv4Pool", "OutpostId", "SecurityGroupId", "SubnetId"},
			"delete-endpoint":       {"EndpointId", "OutpostId"},
			"list-endpoints":        {"MaxResults", "NextToken"},
			"list-outposts-with-s3": {"MaxResults", "NextToken"},
			"list-shared-endpoints": {"MaxResults", "NextToken", "OutpostId"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-endpoint":       {"AccessType": "types.EndpointAccessType", "CustomerOwnedIpv4Pool": "*string", "OutpostId": "*string", "SecurityGroupId": "*string", "SubnetId": "*string"},
			"delete-endpoint":       {"EndpointId": "*string", "OutpostId": "*string"},
			"list-endpoints":        {"MaxResults": "int32", "NextToken": "*string"},
			"list-outposts-with-s3": {"MaxResults": "int32", "NextToken": "*string"},
			"list-shared-endpoints": {"MaxResults": "int32", "NextToken": "*string", "OutpostId": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"create-endpoint":       {"OutpostId", "SecurityGroupId", "SubnetId"},
			"delete-endpoint":       {"EndpointId", "OutpostId"},
			"list-endpoints":        {},
			"list-outposts-with-s3": {},
			"list-shared-endpoints": {"OutpostId"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("s3outposts", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
