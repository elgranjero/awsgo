package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/healthlake/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-fhir-datastore", "delete-fhir-datastore", "describe-fhir-datastore", "describe-fhir-export-job", "describe-fhir-import-job", "list-fhir-datastores", "list-fhir-export-jobs", "list-fhir-import-jobs", "list-tags-for-resource", "start-fhir-export-job", "start-fhir-import-job", "tag-resource", "untag-resource"},
		OperationSet: map[string]bool{"create-fhir-datastore": true, "delete-fhir-datastore": true, "describe-fhir-datastore": true, "describe-fhir-export-job": true, "describe-fhir-import-job": true, "list-fhir-datastores": true, "list-fhir-export-jobs": true, "list-fhir-import-jobs": true, "list-tags-for-resource": true, "start-fhir-export-job": true, "start-fhir-import-job": true, "tag-resource": true, "untag-resource": true},
		OperationInputs: map[string][]string{
			"create-fhir-datastore":    {"ClientToken", "DatastoreName", "DatastoreTypeVersion", "IdentityProviderConfiguration", "PreloadDataConfig", "SseConfiguration", "Tags"},
			"delete-fhir-datastore":    {"DatastoreId"},
			"describe-fhir-datastore":  {"DatastoreId"},
			"describe-fhir-export-job": {"DatastoreId", "JobId"},
			"describe-fhir-import-job": {"DatastoreId", "JobId"},
			"list-fhir-datastores":     {"Filter", "MaxResults", "NextToken"},
			"list-fhir-export-jobs":    {"DatastoreId", "JobName", "JobStatus", "MaxResults", "NextToken", "SubmittedAfter", "SubmittedBefore"},
			"list-fhir-import-jobs":    {"DatastoreId", "JobName", "JobStatus", "MaxResults", "NextToken", "SubmittedAfter", "SubmittedBefore"},
			"list-tags-for-resource":   {"ResourceARN"},
			"start-fhir-export-job":    {"ClientToken", "DataAccessRoleArn", "DatastoreId", "JobName", "OutputDataConfig"},
			"start-fhir-import-job":    {"ClientToken", "DataAccessRoleArn", "DatastoreId", "InputDataConfig", "JobName", "JobOutputDataConfig", "ValidationLevel"},
			"tag-resource":             {"ResourceARN", "Tags"},
			"untag-resource":           {"ResourceARN", "TagKeys"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-fhir-datastore":    {"ClientToken": "*string", "DatastoreName": "*string", "DatastoreTypeVersion": "types.FHIRVersion", "IdentityProviderConfiguration": "*types.IdentityProviderConfiguration", "PreloadDataConfig": "*types.PreloadDataConfig", "SseConfiguration": "*types.SseConfiguration", "Tags": "[]types.Tag"},
			"delete-fhir-datastore":    {"DatastoreId": "*string"},
			"describe-fhir-datastore":  {"DatastoreId": "*string"},
			"describe-fhir-export-job": {"DatastoreId": "*string", "JobId": "*string"},
			"describe-fhir-import-job": {"DatastoreId": "*string", "JobId": "*string"},
			"list-fhir-datastores":     {"Filter": "*types.DatastoreFilter", "MaxResults": "*int32", "NextToken": "*string"},
			"list-fhir-export-jobs":    {"DatastoreId": "*string", "JobName": "*string", "JobStatus": "types.JobStatus", "MaxResults": "*int32", "NextToken": "*string", "SubmittedAfter": "*time.Time", "SubmittedBefore": "*time.Time"},
			"list-fhir-import-jobs":    {"DatastoreId": "*string", "JobName": "*string", "JobStatus": "types.JobStatus", "MaxResults": "*int32", "NextToken": "*string", "SubmittedAfter": "*time.Time", "SubmittedBefore": "*time.Time"},
			"list-tags-for-resource":   {"ResourceARN": "*string"},
			"start-fhir-export-job":    {"ClientToken": "*string", "DataAccessRoleArn": "*string", "DatastoreId": "*string", "JobName": "*string", "OutputDataConfig": "types.OutputDataConfig"},
			"start-fhir-import-job":    {"ClientToken": "*string", "DataAccessRoleArn": "*string", "DatastoreId": "*string", "InputDataConfig": "types.InputDataConfig", "JobName": "*string", "JobOutputDataConfig": "types.OutputDataConfig", "ValidationLevel": "types.ValidationLevel"},
			"tag-resource":             {"ResourceARN": "*string", "Tags": "[]types.Tag"},
			"untag-resource":           {"ResourceARN": "*string", "TagKeys": "[]string"},
		},
		OperationInputRequired: map[string][]string{
			"create-fhir-datastore":    {"DatastoreTypeVersion"},
			"delete-fhir-datastore":    {"DatastoreId"},
			"describe-fhir-datastore":  {"DatastoreId"},
			"describe-fhir-export-job": {"DatastoreId", "JobId"},
			"describe-fhir-import-job": {"DatastoreId", "JobId"},
			"list-fhir-datastores":     {},
			"list-fhir-export-jobs":    {"DatastoreId"},
			"list-fhir-import-jobs":    {"DatastoreId"},
			"list-tags-for-resource":   {"ResourceARN"},
			"start-fhir-export-job":    {"DataAccessRoleArn", "DatastoreId", "OutputDataConfig"},
			"start-fhir-import-job":    {"DataAccessRoleArn", "DatastoreId", "InputDataConfig", "JobOutputDataConfig"},
			"tag-resource":             {"ResourceARN", "Tags"},
			"untag-resource":           {"ResourceARN", "TagKeys"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("healthlake", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
