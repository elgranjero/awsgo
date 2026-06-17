package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/cloudsearchdomain/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"search", "suggest", "upload-documents"},
		OperationSet: map[string]bool{"search": true, "suggest": true, "upload-documents": true},
		OperationInputs: map[string][]string{
			"search":           {"Cursor", "Expr", "Facet", "FilterQuery", "Highlight", "Partial", "Query", "QueryOptions", "QueryParser", "Return", "Size", "Sort", "Start", "Stats"},
			"suggest":          {"Query", "Size", "Suggester"},
			"upload-documents": {"ContentType", "Documents"},
		},
		OperationInputTypes: map[string]map[string]string{
			"search":           {"Cursor": "*string", "Expr": "*string", "Facet": "*string", "FilterQuery": "*string", "Highlight": "*string", "Partial": "bool", "Query": "*string", "QueryOptions": "*string", "QueryParser": "types.QueryParser", "Return": "*string", "Size": "int64", "Sort": "*string", "Start": "int64", "Stats": "*string"},
			"suggest":          {"Query": "*string", "Size": "int64", "Suggester": "*string"},
			"upload-documents": {"ContentType": "types.ContentType", "Documents": "io.Reader"},
		},
		OperationInputRequired: map[string][]string{
			"search":           {"Query"},
			"suggest":          {"Query", "Suggester"},
			"upload-documents": {"ContentType", "Documents"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("cloudsearchdomain", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
