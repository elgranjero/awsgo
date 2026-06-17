package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/textract/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"analyze-document", "analyze-expense", "analyze-id", "create-adapter", "create-adapter-version", "delete-adapter", "delete-adapter-version", "detect-document-text", "get-adapter", "get-adapter-version", "get-document-analysis", "get-document-text-detection", "get-expense-analysis", "get-lending-analysis", "get-lending-analysis-summary", "list-adapter-versions", "list-adapters", "list-tags-for-resource", "start-document-analysis", "start-document-text-detection", "start-expense-analysis", "start-lending-analysis", "tag-resource", "untag-resource", "update-adapter"},
		OperationSet: map[string]bool{"analyze-document": true, "analyze-expense": true, "analyze-id": true, "create-adapter": true, "create-adapter-version": true, "delete-adapter": true, "delete-adapter-version": true, "detect-document-text": true, "get-adapter": true, "get-adapter-version": true, "get-document-analysis": true, "get-document-text-detection": true, "get-expense-analysis": true, "get-lending-analysis": true, "get-lending-analysis-summary": true, "list-adapter-versions": true, "list-adapters": true, "list-tags-for-resource": true, "start-document-analysis": true, "start-document-text-detection": true, "start-expense-analysis": true, "start-lending-analysis": true, "tag-resource": true, "untag-resource": true, "update-adapter": true},
		OperationInputs: map[string][]string{
			"analyze-document":              {"AdaptersConfig", "Document", "FeatureTypes", "HumanLoopConfig", "QueriesConfig"},
			"analyze-expense":               {"Document"},
			"analyze-id":                    {"DocumentPages"},
			"create-adapter":                {"AdapterName", "AutoUpdate", "ClientRequestToken", "Description", "FeatureTypes", "Tags"},
			"create-adapter-version":        {"AdapterId", "ClientRequestToken", "DatasetConfig", "KMSKeyId", "OutputConfig", "Tags"},
			"delete-adapter":                {"AdapterId"},
			"delete-adapter-version":        {"AdapterId", "AdapterVersion"},
			"detect-document-text":          {"Document"},
			"get-adapter":                   {"AdapterId"},
			"get-adapter-version":           {"AdapterId", "AdapterVersion"},
			"get-document-analysis":         {"JobId", "MaxResults", "NextToken"},
			"get-document-text-detection":   {"JobId", "MaxResults", "NextToken"},
			"get-expense-analysis":          {"JobId", "MaxResults", "NextToken"},
			"get-lending-analysis":          {"JobId", "MaxResults", "NextToken"},
			"get-lending-analysis-summary":  {"JobId"},
			"list-adapter-versions":         {"AdapterId", "AfterCreationTime", "BeforeCreationTime", "MaxResults", "NextToken"},
			"list-adapters":                 {"AfterCreationTime", "BeforeCreationTime", "MaxResults", "NextToken"},
			"list-tags-for-resource":        {"ResourceARN"},
			"start-document-analysis":       {"AdaptersConfig", "ClientRequestToken", "DocumentLocation", "FeatureTypes", "JobTag", "KMSKeyId", "NotificationChannel", "OutputConfig", "QueriesConfig"},
			"start-document-text-detection": {"ClientRequestToken", "DocumentLocation", "JobTag", "KMSKeyId", "NotificationChannel", "OutputConfig"},
			"start-expense-analysis":        {"ClientRequestToken", "DocumentLocation", "JobTag", "KMSKeyId", "NotificationChannel", "OutputConfig"},
			"start-lending-analysis":        {"ClientRequestToken", "DocumentLocation", "JobTag", "KMSKeyId", "NotificationChannel", "OutputConfig"},
			"tag-resource":                  {"ResourceARN", "Tags"},
			"untag-resource":                {"ResourceARN", "TagKeys"},
			"update-adapter":                {"AdapterId", "AdapterName", "AutoUpdate", "Description"},
		},
		OperationInputTypes: map[string]map[string]string{
			"analyze-document":              {"AdaptersConfig": "*types.AdaptersConfig", "Document": "*types.Document", "FeatureTypes": "[]types.FeatureType", "HumanLoopConfig": "*types.HumanLoopConfig", "QueriesConfig": "*types.QueriesConfig"},
			"analyze-expense":               {"Document": "*types.Document"},
			"analyze-id":                    {"DocumentPages": "[]types.Document"},
			"create-adapter":                {"AdapterName": "*string", "AutoUpdate": "types.AutoUpdate", "ClientRequestToken": "*string", "Description": "*string", "FeatureTypes": "[]types.FeatureType", "Tags": "map[string]string"},
			"create-adapter-version":        {"AdapterId": "*string", "ClientRequestToken": "*string", "DatasetConfig": "*types.AdapterVersionDatasetConfig", "KMSKeyId": "*string", "OutputConfig": "*types.OutputConfig", "Tags": "map[string]string"},
			"delete-adapter":                {"AdapterId": "*string"},
			"delete-adapter-version":        {"AdapterId": "*string", "AdapterVersion": "*string"},
			"detect-document-text":          {"Document": "*types.Document"},
			"get-adapter":                   {"AdapterId": "*string"},
			"get-adapter-version":           {"AdapterId": "*string", "AdapterVersion": "*string"},
			"get-document-analysis":         {"JobId": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"get-document-text-detection":   {"JobId": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"get-expense-analysis":          {"JobId": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"get-lending-analysis":          {"JobId": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"get-lending-analysis-summary":  {"JobId": "*string"},
			"list-adapter-versions":         {"AdapterId": "*string", "AfterCreationTime": "*time.Time", "BeforeCreationTime": "*time.Time", "MaxResults": "*int32", "NextToken": "*string"},
			"list-adapters":                 {"AfterCreationTime": "*time.Time", "BeforeCreationTime": "*time.Time", "MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":        {"ResourceARN": "*string"},
			"start-document-analysis":       {"AdaptersConfig": "*types.AdaptersConfig", "ClientRequestToken": "*string", "DocumentLocation": "*types.DocumentLocation", "FeatureTypes": "[]types.FeatureType", "JobTag": "*string", "KMSKeyId": "*string", "NotificationChannel": "*types.NotificationChannel", "OutputConfig": "*types.OutputConfig", "QueriesConfig": "*types.QueriesConfig"},
			"start-document-text-detection": {"ClientRequestToken": "*string", "DocumentLocation": "*types.DocumentLocation", "JobTag": "*string", "KMSKeyId": "*string", "NotificationChannel": "*types.NotificationChannel", "OutputConfig": "*types.OutputConfig"},
			"start-expense-analysis":        {"ClientRequestToken": "*string", "DocumentLocation": "*types.DocumentLocation", "JobTag": "*string", "KMSKeyId": "*string", "NotificationChannel": "*types.NotificationChannel", "OutputConfig": "*types.OutputConfig"},
			"start-lending-analysis":        {"ClientRequestToken": "*string", "DocumentLocation": "*types.DocumentLocation", "JobTag": "*string", "KMSKeyId": "*string", "NotificationChannel": "*types.NotificationChannel", "OutputConfig": "*types.OutputConfig"},
			"tag-resource":                  {"ResourceARN": "*string", "Tags": "map[string]string"},
			"untag-resource":                {"ResourceARN": "*string", "TagKeys": "[]string"},
			"update-adapter":                {"AdapterId": "*string", "AdapterName": "*string", "AutoUpdate": "types.AutoUpdate", "Description": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"analyze-document":              {"Document", "FeatureTypes"},
			"analyze-expense":               {"Document"},
			"analyze-id":                    {"DocumentPages"},
			"create-adapter":                {"AdapterName", "FeatureTypes"},
			"create-adapter-version":        {"AdapterId", "DatasetConfig", "OutputConfig"},
			"delete-adapter":                {"AdapterId"},
			"delete-adapter-version":        {"AdapterId", "AdapterVersion"},
			"detect-document-text":          {"Document"},
			"get-adapter":                   {"AdapterId"},
			"get-adapter-version":           {"AdapterId", "AdapterVersion"},
			"get-document-analysis":         {"JobId"},
			"get-document-text-detection":   {"JobId"},
			"get-expense-analysis":          {"JobId"},
			"get-lending-analysis":          {"JobId"},
			"get-lending-analysis-summary":  {"JobId"},
			"list-adapter-versions":         {},
			"list-adapters":                 {},
			"list-tags-for-resource":        {"ResourceARN"},
			"start-document-analysis":       {"DocumentLocation", "FeatureTypes"},
			"start-document-text-detection": {"DocumentLocation"},
			"start-expense-analysis":        {"DocumentLocation"},
			"start-lending-analysis":        {"DocumentLocation"},
			"tag-resource":                  {"ResourceARN", "Tags"},
			"untag-resource":                {"ResourceARN", "TagKeys"},
			"update-adapter":                {"AdapterId"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("textract", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
