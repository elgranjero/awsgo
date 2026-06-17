package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/medicalimaging/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"copy-image-set", "create-datastore", "delete-datastore", "delete-image-set", "get-datastore", "get-dicom-import-job", "get-image-frame", "get-image-set", "get-image-set-metadata", "list-datastores", "list-dicom-import-jobs", "list-image-set-versions", "list-tags-for-resource", "search-image-sets", "start-dicom-import-job", "tag-resource", "untag-resource", "update-image-set-metadata"},
		OperationSet: map[string]bool{"copy-image-set": true, "create-datastore": true, "delete-datastore": true, "delete-image-set": true, "get-datastore": true, "get-dicom-import-job": true, "get-image-frame": true, "get-image-set": true, "get-image-set-metadata": true, "list-datastores": true, "list-dicom-import-jobs": true, "list-image-set-versions": true, "list-tags-for-resource": true, "search-image-sets": true, "start-dicom-import-job": true, "tag-resource": true, "untag-resource": true, "update-image-set-metadata": true},
		OperationInputs: map[string][]string{
			"copy-image-set":            {"CopyImageSetInformation", "DatastoreId", "Force", "PromoteToPrimary", "SourceImageSetId"},
			"create-datastore":          {"ClientToken", "DatastoreName", "KmsKeyArn", "LambdaAuthorizerArn", "LosslessStorageFormat", "Tags"},
			"delete-datastore":          {"DatastoreId"},
			"delete-image-set":          {"DatastoreId", "ImageSetId"},
			"get-datastore":             {"DatastoreId"},
			"get-dicom-import-job":      {"DatastoreId", "JobId"},
			"get-image-frame":           {"DatastoreId", "ImageFrameInformation", "ImageSetId"},
			"get-image-set":             {"DatastoreId", "ImageSetId", "VersionId"},
			"get-image-set-metadata":    {"DatastoreId", "ImageSetId", "VersionId"},
			"list-datastores":           {"DatastoreStatus", "MaxResults", "NextToken"},
			"list-dicom-import-jobs":    {"DatastoreId", "JobStatus", "MaxResults", "NextToken"},
			"list-image-set-versions":   {"DatastoreId", "ImageSetId", "MaxResults", "NextToken"},
			"list-tags-for-resource":    {"ResourceArn"},
			"search-image-sets":         {"DatastoreId", "MaxResults", "NextToken", "SearchCriteria"},
			"start-dicom-import-job":    {"ClientToken", "DataAccessRoleArn", "DatastoreId", "InputOwnerAccountId", "InputS3Uri", "JobName", "OutputS3Uri"},
			"tag-resource":              {"ResourceArn", "Tags"},
			"untag-resource":            {"ResourceArn", "TagKeys"},
			"update-image-set-metadata": {"DatastoreId", "Force", "ImageSetId", "LatestVersionId", "UpdateImageSetMetadataUpdates"},
		},
		OperationInputTypes: map[string]map[string]string{
			"copy-image-set":            {"CopyImageSetInformation": "*types.CopyImageSetInformation", "DatastoreId": "*string", "Force": "*bool", "PromoteToPrimary": "*bool", "SourceImageSetId": "*string"},
			"create-datastore":          {"ClientToken": "*string", "DatastoreName": "*string", "KmsKeyArn": "*string", "LambdaAuthorizerArn": "*string", "LosslessStorageFormat": "types.LosslessStorageFormat", "Tags": "map[string]string"},
			"delete-datastore":          {"DatastoreId": "*string"},
			"delete-image-set":          {"DatastoreId": "*string", "ImageSetId": "*string"},
			"get-datastore":             {"DatastoreId": "*string"},
			"get-dicom-import-job":      {"DatastoreId": "*string", "JobId": "*string"},
			"get-image-frame":           {"DatastoreId": "*string", "ImageFrameInformation": "*types.ImageFrameInformation", "ImageSetId": "*string"},
			"get-image-set":             {"DatastoreId": "*string", "ImageSetId": "*string", "VersionId": "*string"},
			"get-image-set-metadata":    {"DatastoreId": "*string", "ImageSetId": "*string", "VersionId": "*string"},
			"list-datastores":           {"DatastoreStatus": "types.DatastoreStatus", "MaxResults": "*int32", "NextToken": "*string"},
			"list-dicom-import-jobs":    {"DatastoreId": "*string", "JobStatus": "types.JobStatus", "MaxResults": "*int32", "NextToken": "*string"},
			"list-image-set-versions":   {"DatastoreId": "*string", "ImageSetId": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":    {"ResourceArn": "*string"},
			"search-image-sets":         {"DatastoreId": "*string", "MaxResults": "*int32", "NextToken": "*string", "SearchCriteria": "*types.SearchCriteria"},
			"start-dicom-import-job":    {"ClientToken": "*string", "DataAccessRoleArn": "*string", "DatastoreId": "*string", "InputOwnerAccountId": "*string", "InputS3Uri": "*string", "JobName": "*string", "OutputS3Uri": "*string"},
			"tag-resource":              {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":            {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-image-set-metadata": {"DatastoreId": "*string", "Force": "*bool", "ImageSetId": "*string", "LatestVersionId": "*string", "UpdateImageSetMetadataUpdates": "types.MetadataUpdates"},
		},
		OperationInputRequired: map[string][]string{
			"copy-image-set":            {"CopyImageSetInformation", "DatastoreId", "SourceImageSetId"},
			"create-datastore":          {"ClientToken"},
			"delete-datastore":          {"DatastoreId"},
			"delete-image-set":          {"DatastoreId", "ImageSetId"},
			"get-datastore":             {"DatastoreId"},
			"get-dicom-import-job":      {"DatastoreId", "JobId"},
			"get-image-frame":           {"DatastoreId", "ImageFrameInformation", "ImageSetId"},
			"get-image-set":             {"DatastoreId", "ImageSetId"},
			"get-image-set-metadata":    {"DatastoreId", "ImageSetId"},
			"list-datastores":           {},
			"list-dicom-import-jobs":    {"DatastoreId"},
			"list-image-set-versions":   {"DatastoreId", "ImageSetId"},
			"list-tags-for-resource":    {"ResourceArn"},
			"search-image-sets":         {"DatastoreId"},
			"start-dicom-import-job":    {"ClientToken", "DataAccessRoleArn", "DatastoreId", "InputS3Uri", "OutputS3Uri"},
			"tag-resource":              {"ResourceArn", "Tags"},
			"untag-resource":            {"ResourceArn", "TagKeys"},
			"update-image-set-metadata": {"DatastoreId", "ImageSetId", "LatestVersionId", "UpdateImageSetMetadataUpdates"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("medicalimaging", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
