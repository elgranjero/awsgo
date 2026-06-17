package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/signerdata/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"get-revocation-status"},
		OperationSet: map[string]bool{"get-revocation-status": true},
		OperationInputs: map[string][]string{
			"get-revocation-status": {"CertificateHashes", "JobArn", "PlatformId", "ProfileVersionArn", "SignatureTimestamp"},
		},
		OperationInputTypes: map[string]map[string]string{
			"get-revocation-status": {"CertificateHashes": "[]string", "JobArn": "*string", "PlatformId": "*string", "ProfileVersionArn": "*string", "SignatureTimestamp": "*time.Time"},
		},
		OperationInputRequired: map[string][]string{
			"get-revocation-status": {"CertificateHashes", "JobArn", "PlatformId", "ProfileVersionArn", "SignatureTimestamp"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("signerdata", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
