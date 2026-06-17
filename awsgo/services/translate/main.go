package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/translate/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-parallel-data", "delete-parallel-data", "delete-terminology", "describe-text-translation-job", "get-parallel-data", "get-terminology", "import-terminology", "list-languages", "list-parallel-data", "list-tags-for-resource", "list-terminologies", "list-text-translation-jobs", "start-text-translation-job", "stop-text-translation-job", "tag-resource", "translate-document", "translate-text", "untag-resource", "update-parallel-data"},
		OperationSet: map[string]bool{"create-parallel-data": true, "delete-parallel-data": true, "delete-terminology": true, "describe-text-translation-job": true, "get-parallel-data": true, "get-terminology": true, "import-terminology": true, "list-languages": true, "list-parallel-data": true, "list-tags-for-resource": true, "list-terminologies": true, "list-text-translation-jobs": true, "start-text-translation-job": true, "stop-text-translation-job": true, "tag-resource": true, "translate-document": true, "translate-text": true, "untag-resource": true, "update-parallel-data": true},
		OperationInputs: map[string][]string{
			"create-parallel-data":          {"ClientToken", "Description", "EncryptionKey", "Name", "ParallelDataConfig", "Tags"},
			"delete-parallel-data":          {"Name"},
			"delete-terminology":            {"Name"},
			"describe-text-translation-job": {"JobId"},
			"get-parallel-data":             {"Name"},
			"get-terminology":               {"Name", "TerminologyDataFormat"},
			"import-terminology":            {"Description", "EncryptionKey", "MergeStrategy", "Name", "Tags", "TerminologyData"},
			"list-languages":                {"DisplayLanguageCode", "MaxResults", "NextToken"},
			"list-parallel-data":            {"MaxResults", "NextToken"},
			"list-tags-for-resource":        {"ResourceArn"},
			"list-terminologies":            {"MaxResults", "NextToken"},
			"list-text-translation-jobs":    {"Filter", "MaxResults", "NextToken"},
			"start-text-translation-job":    {"ClientToken", "DataAccessRoleArn", "InputDataConfig", "JobName", "OutputDataConfig", "ParallelDataNames", "Settings", "SourceLanguageCode", "TargetLanguageCodes", "TerminologyNames"},
			"stop-text-translation-job":     {"JobId"},
			"tag-resource":                  {"ResourceArn", "Tags"},
			"translate-document":            {"Document", "Settings", "SourceLanguageCode", "TargetLanguageCode", "TerminologyNames"},
			"translate-text":                {"Settings", "SourceLanguageCode", "TargetLanguageCode", "TerminologyNames", "Text"},
			"untag-resource":                {"ResourceArn", "TagKeys"},
			"update-parallel-data":          {"ClientToken", "Description", "Name", "ParallelDataConfig"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-parallel-data":          {"ClientToken": "*string", "Description": "*string", "EncryptionKey": "*types.EncryptionKey", "Name": "*string", "ParallelDataConfig": "*types.ParallelDataConfig", "Tags": "[]types.Tag"},
			"delete-parallel-data":          {"Name": "*string"},
			"delete-terminology":            {"Name": "*string"},
			"describe-text-translation-job": {"JobId": "*string"},
			"get-parallel-data":             {"Name": "*string"},
			"get-terminology":               {"Name": "*string", "TerminologyDataFormat": "types.TerminologyDataFormat"},
			"import-terminology":            {"Description": "*string", "EncryptionKey": "*types.EncryptionKey", "MergeStrategy": "types.MergeStrategy", "Name": "*string", "Tags": "[]types.Tag", "TerminologyData": "*types.TerminologyData"},
			"list-languages":                {"DisplayLanguageCode": "types.DisplayLanguageCode", "MaxResults": "*int32", "NextToken": "*string"},
			"list-parallel-data":            {"MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":        {"ResourceArn": "*string"},
			"list-terminologies":            {"MaxResults": "*int32", "NextToken": "*string"},
			"list-text-translation-jobs":    {"Filter": "*types.TextTranslationJobFilter", "MaxResults": "*int32", "NextToken": "*string"},
			"start-text-translation-job":    {"ClientToken": "*string", "DataAccessRoleArn": "*string", "InputDataConfig": "*types.InputDataConfig", "JobName": "*string", "OutputDataConfig": "*types.OutputDataConfig", "ParallelDataNames": "[]string", "Settings": "*types.TranslationSettings", "SourceLanguageCode": "*string", "TargetLanguageCodes": "[]string", "TerminologyNames": "[]string"},
			"stop-text-translation-job":     {"JobId": "*string"},
			"tag-resource":                  {"ResourceArn": "*string", "Tags": "[]types.Tag"},
			"translate-document":            {"Document": "*types.Document", "Settings": "*types.TranslationSettings", "SourceLanguageCode": "*string", "TargetLanguageCode": "*string", "TerminologyNames": "[]string"},
			"translate-text":                {"Settings": "*types.TranslationSettings", "SourceLanguageCode": "*string", "TargetLanguageCode": "*string", "TerminologyNames": "[]string", "Text": "*string"},
			"untag-resource":                {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-parallel-data":          {"ClientToken": "*string", "Description": "*string", "Name": "*string", "ParallelDataConfig": "*types.ParallelDataConfig"},
		},
		OperationInputRequired: map[string][]string{
			"create-parallel-data":          {"ClientToken", "Name", "ParallelDataConfig"},
			"delete-parallel-data":          {"Name"},
			"delete-terminology":            {"Name"},
			"describe-text-translation-job": {"JobId"},
			"get-parallel-data":             {"Name"},
			"get-terminology":               {"Name"},
			"import-terminology":            {"MergeStrategy", "Name", "TerminologyData"},
			"list-languages":                {},
			"list-parallel-data":            {},
			"list-tags-for-resource":        {"ResourceArn"},
			"list-terminologies":            {},
			"list-text-translation-jobs":    {},
			"start-text-translation-job":    {"ClientToken", "DataAccessRoleArn", "InputDataConfig", "OutputDataConfig", "SourceLanguageCode", "TargetLanguageCodes"},
			"stop-text-translation-job":     {"JobId"},
			"tag-resource":                  {"ResourceArn", "Tags"},
			"translate-document":            {"Document", "SourceLanguageCode", "TargetLanguageCode"},
			"translate-text":                {"SourceLanguageCode", "TargetLanguageCode", "Text"},
			"untag-resource":                {"ResourceArn", "TagKeys"},
			"update-parallel-data":          {"ClientToken", "Name", "ParallelDataConfig"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("translate", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
