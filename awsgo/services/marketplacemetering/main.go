package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/marketplacemetering/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"batch-meter-usage", "meter-usage", "register-usage", "resolve-customer"},
		OperationSet: map[string]bool{"batch-meter-usage": true, "meter-usage": true, "register-usage": true, "resolve-customer": true},
		OperationInputs: map[string][]string{
			"batch-meter-usage": {"ProductCode", "UsageRecords"},
			"meter-usage":       {"ClientToken", "DryRun", "ProductCode", "Timestamp", "UsageAllocations", "UsageDimension", "UsageQuantity"},
			"register-usage":    {"Nonce", "ProductCode", "PublicKeyVersion"},
			"resolve-customer":  {"RegistrationToken"},
		},
		OperationInputTypes: map[string]map[string]string{
			"batch-meter-usage": {"ProductCode": "*string", "UsageRecords": "[]types.UsageRecord"},
			"meter-usage":       {"ClientToken": "*string", "DryRun": "*bool", "ProductCode": "*string", "Timestamp": "*time.Time", "UsageAllocations": "[]types.UsageAllocation", "UsageDimension": "*string", "UsageQuantity": "*int32"},
			"register-usage":    {"Nonce": "*string", "ProductCode": "*string", "PublicKeyVersion": "*int32"},
			"resolve-customer":  {"RegistrationToken": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"batch-meter-usage": {"UsageRecords"},
			"meter-usage":       {"ProductCode", "Timestamp", "UsageDimension"},
			"register-usage":    {"ProductCode", "PublicKeyVersion"},
			"resolve-customer":  {"RegistrationToken"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("marketplacemetering", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
