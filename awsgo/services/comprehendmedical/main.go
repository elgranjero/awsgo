package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/comprehendmedical/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"describe-entities-detection-v2-job", "describe-icd10cm-inference-job", "describe-phi-detection-job", "describe-rx-norm-inference-job", "describe-snomedct-inference-job", "detect-entities", "detect-entities-v2", "detect-phi", "infer-icd10cm", "infer-rx-norm", "infer-snomedct", "list-entities-detection-v2-jobs", "list-icd10cm-inference-jobs", "list-phi-detection-jobs", "list-rx-norm-inference-jobs", "list-snomedct-inference-jobs", "start-entities-detection-v2-job", "start-icd10cm-inference-job", "start-phi-detection-job", "start-rx-norm-inference-job", "start-snomedct-inference-job", "stop-entities-detection-v2-job", "stop-icd10cm-inference-job", "stop-phi-detection-job", "stop-rx-norm-inference-job", "stop-snomedct-inference-job"},
		OperationSet: map[string]bool{"describe-entities-detection-v2-job": true, "describe-icd10cm-inference-job": true, "describe-phi-detection-job": true, "describe-rx-norm-inference-job": true, "describe-snomedct-inference-job": true, "detect-entities": true, "detect-entities-v2": true, "detect-phi": true, "infer-icd10cm": true, "infer-rx-norm": true, "infer-snomedct": true, "list-entities-detection-v2-jobs": true, "list-icd10cm-inference-jobs": true, "list-phi-detection-jobs": true, "list-rx-norm-inference-jobs": true, "list-snomedct-inference-jobs": true, "start-entities-detection-v2-job": true, "start-icd10cm-inference-job": true, "start-phi-detection-job": true, "start-rx-norm-inference-job": true, "start-snomedct-inference-job": true, "stop-entities-detection-v2-job": true, "stop-icd10cm-inference-job": true, "stop-phi-detection-job": true, "stop-rx-norm-inference-job": true, "stop-snomedct-inference-job": true},
		OperationInputs: map[string][]string{
			"describe-entities-detection-v2-job": {"JobId"},
			"describe-icd10cm-inference-job":     {"JobId"},
			"describe-phi-detection-job":         {"JobId"},
			"describe-rx-norm-inference-job":     {"JobId"},
			"describe-snomedct-inference-job":    {"JobId"},
			"detect-entities":                    {"Text"},
			"detect-entities-v2":                 {"Text"},
			"detect-phi":                         {"Text"},
			"infer-icd10cm":                      {"Text"},
			"infer-rx-norm":                      {"Text"},
			"infer-snomedct":                     {"Text"},
			"list-entities-detection-v2-jobs":    {"Filter", "MaxResults", "NextToken"},
			"list-icd10cm-inference-jobs":        {"Filter", "MaxResults", "NextToken"},
			"list-phi-detection-jobs":            {"Filter", "MaxResults", "NextToken"},
			"list-rx-norm-inference-jobs":        {"Filter", "MaxResults", "NextToken"},
			"list-snomedct-inference-jobs":       {"Filter", "MaxResults", "NextToken"},
			"start-entities-detection-v2-job":    {"ClientRequestToken", "DataAccessRoleArn", "InputDataConfig", "JobName", "KMSKey", "LanguageCode", "OutputDataConfig"},
			"start-icd10cm-inference-job":        {"ClientRequestToken", "DataAccessRoleArn", "InputDataConfig", "JobName", "KMSKey", "LanguageCode", "OutputDataConfig"},
			"start-phi-detection-job":            {"ClientRequestToken", "DataAccessRoleArn", "InputDataConfig", "JobName", "KMSKey", "LanguageCode", "OutputDataConfig"},
			"start-rx-norm-inference-job":        {"ClientRequestToken", "DataAccessRoleArn", "InputDataConfig", "JobName", "KMSKey", "LanguageCode", "OutputDataConfig"},
			"start-snomedct-inference-job":       {"ClientRequestToken", "DataAccessRoleArn", "InputDataConfig", "JobName", "KMSKey", "LanguageCode", "OutputDataConfig"},
			"stop-entities-detection-v2-job":     {"JobId"},
			"stop-icd10cm-inference-job":         {"JobId"},
			"stop-phi-detection-job":             {"JobId"},
			"stop-rx-norm-inference-job":         {"JobId"},
			"stop-snomedct-inference-job":        {"JobId"},
		},
		OperationInputTypes: map[string]map[string]string{
			"describe-entities-detection-v2-job": {"JobId": "*string"},
			"describe-icd10cm-inference-job":     {"JobId": "*string"},
			"describe-phi-detection-job":         {"JobId": "*string"},
			"describe-rx-norm-inference-job":     {"JobId": "*string"},
			"describe-snomedct-inference-job":    {"JobId": "*string"},
			"detect-entities":                    {"Text": "*string"},
			"detect-entities-v2":                 {"Text": "*string"},
			"detect-phi":                         {"Text": "*string"},
			"infer-icd10cm":                      {"Text": "*string"},
			"infer-rx-norm":                      {"Text": "*string"},
			"infer-snomedct":                     {"Text": "*string"},
			"list-entities-detection-v2-jobs":    {"Filter": "*types.ComprehendMedicalAsyncJobFilter", "MaxResults": "*int32", "NextToken": "*string"},
			"list-icd10cm-inference-jobs":        {"Filter": "*types.ComprehendMedicalAsyncJobFilter", "MaxResults": "*int32", "NextToken": "*string"},
			"list-phi-detection-jobs":            {"Filter": "*types.ComprehendMedicalAsyncJobFilter", "MaxResults": "*int32", "NextToken": "*string"},
			"list-rx-norm-inference-jobs":        {"Filter": "*types.ComprehendMedicalAsyncJobFilter", "MaxResults": "*int32", "NextToken": "*string"},
			"list-snomedct-inference-jobs":       {"Filter": "*types.ComprehendMedicalAsyncJobFilter", "MaxResults": "*int32", "NextToken": "*string"},
			"start-entities-detection-v2-job":    {"ClientRequestToken": "*string", "DataAccessRoleArn": "*string", "InputDataConfig": "*types.InputDataConfig", "JobName": "*string", "KMSKey": "*string", "LanguageCode": "types.LanguageCode", "OutputDataConfig": "*types.OutputDataConfig"},
			"start-icd10cm-inference-job":        {"ClientRequestToken": "*string", "DataAccessRoleArn": "*string", "InputDataConfig": "*types.InputDataConfig", "JobName": "*string", "KMSKey": "*string", "LanguageCode": "types.LanguageCode", "OutputDataConfig": "*types.OutputDataConfig"},
			"start-phi-detection-job":            {"ClientRequestToken": "*string", "DataAccessRoleArn": "*string", "InputDataConfig": "*types.InputDataConfig", "JobName": "*string", "KMSKey": "*string", "LanguageCode": "types.LanguageCode", "OutputDataConfig": "*types.OutputDataConfig"},
			"start-rx-norm-inference-job":        {"ClientRequestToken": "*string", "DataAccessRoleArn": "*string", "InputDataConfig": "*types.InputDataConfig", "JobName": "*string", "KMSKey": "*string", "LanguageCode": "types.LanguageCode", "OutputDataConfig": "*types.OutputDataConfig"},
			"start-snomedct-inference-job":       {"ClientRequestToken": "*string", "DataAccessRoleArn": "*string", "InputDataConfig": "*types.InputDataConfig", "JobName": "*string", "KMSKey": "*string", "LanguageCode": "types.LanguageCode", "OutputDataConfig": "*types.OutputDataConfig"},
			"stop-entities-detection-v2-job":     {"JobId": "*string"},
			"stop-icd10cm-inference-job":         {"JobId": "*string"},
			"stop-phi-detection-job":             {"JobId": "*string"},
			"stop-rx-norm-inference-job":         {"JobId": "*string"},
			"stop-snomedct-inference-job":        {"JobId": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"describe-entities-detection-v2-job": {"JobId"},
			"describe-icd10cm-inference-job":     {"JobId"},
			"describe-phi-detection-job":         {"JobId"},
			"describe-rx-norm-inference-job":     {"JobId"},
			"describe-snomedct-inference-job":    {"JobId"},
			"detect-entities":                    {"Text"},
			"detect-entities-v2":                 {"Text"},
			"detect-phi":                         {"Text"},
			"infer-icd10cm":                      {"Text"},
			"infer-rx-norm":                      {"Text"},
			"infer-snomedct":                     {"Text"},
			"list-entities-detection-v2-jobs":    {},
			"list-icd10cm-inference-jobs":        {},
			"list-phi-detection-jobs":            {},
			"list-rx-norm-inference-jobs":        {},
			"list-snomedct-inference-jobs":       {},
			"start-entities-detection-v2-job":    {"DataAccessRoleArn", "InputDataConfig", "LanguageCode", "OutputDataConfig"},
			"start-icd10cm-inference-job":        {"DataAccessRoleArn", "InputDataConfig", "LanguageCode", "OutputDataConfig"},
			"start-phi-detection-job":            {"DataAccessRoleArn", "InputDataConfig", "LanguageCode", "OutputDataConfig"},
			"start-rx-norm-inference-job":        {"DataAccessRoleArn", "InputDataConfig", "LanguageCode", "OutputDataConfig"},
			"start-snomedct-inference-job":       {"DataAccessRoleArn", "InputDataConfig", "LanguageCode", "OutputDataConfig"},
			"stop-entities-detection-v2-job":     {"JobId"},
			"stop-icd10cm-inference-job":         {"JobId"},
			"stop-phi-detection-job":             {"JobId"},
			"stop-rx-norm-inference-job":         {"JobId"},
			"stop-snomedct-inference-job":        {"JobId"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("comprehendmedical", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
