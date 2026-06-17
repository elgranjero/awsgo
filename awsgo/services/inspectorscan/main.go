package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/inspectorscan/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"scan-sbom"},
		OperationSet: map[string]bool{"scan-sbom": true},
		OperationInputs: map[string][]string{
			"scan-sbom": {"OutputFormat", "Sbom"},
		},
		OperationInputTypes: map[string]map[string]string{
			"scan-sbom": {"OutputFormat": "types.OutputFormat", "Sbom": "document.Interface"},
		},
		OperationInputRequired: map[string][]string{
			"scan-sbom": {"Sbom"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("inspectorscan", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
