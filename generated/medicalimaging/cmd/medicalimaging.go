package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/medicalimaging"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// medicalimagingCmd represents the medicalimaging command
var _medicalimagingCmd = &cobra.Command{
	Use:   "medicalimaging",
	Short: "AWS medicalimaging CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := medicalimaging.NewFromConfig(cfg)
		if _medicalimagingCopyImageSet {
			medicalimaging_CopyImageSet(cfg, client)
			return
		}
		if _medicalimagingCreateDatastore {
			medicalimaging_CreateDatastore(cfg, client)
			return
		}
		if _medicalimagingDeleteDatastore {
			medicalimaging_DeleteDatastore(cfg, client)
			return
		}
		if _medicalimagingDeleteImageSet {
			medicalimaging_DeleteImageSet(cfg, client)
			return
		}
		if _medicalimagingGetDatastore {
			medicalimaging_GetDatastore(cfg, client)
			return
		}
		if _medicalimagingGetDICOMImportJob {
			medicalimaging_GetDICOMImportJob(cfg, client)
			return
		}
		if _medicalimagingGetImageFrame {
			medicalimaging_GetImageFrame(cfg, client)
			return
		}
		if _medicalimagingGetImageSet {
			medicalimaging_GetImageSet(cfg, client)
			return
		}
		if _medicalimagingGetImageSetMetadata {
			medicalimaging_GetImageSetMetadata(cfg, client)
			return
		}
		if _medicalimagingListDatastores {
			medicalimaging_ListDatastores(cfg, client)
			return
		}
		if _medicalimagingListDICOMImportJobs {
			medicalimaging_ListDICOMImportJobs(cfg, client)
			return
		}
		if _medicalimagingListImageSetVersions {
			medicalimaging_ListImageSetVersions(cfg, client)
			return
		}
		if _medicalimagingListTagsForResource {
			medicalimaging_ListTagsForResource(cfg, client)
			return
		}
		if _medicalimagingSearchImageSets {
			medicalimaging_SearchImageSets(cfg, client)
			return
		}
		if _medicalimagingStartDICOMImportJob {
			medicalimaging_StartDICOMImportJob(cfg, client)
			return
		}
		if _medicalimagingTagResource {
			medicalimaging_TagResource(cfg, client)
			return
		}
		if _medicalimagingUntagResource {
			medicalimaging_UntagResource(cfg, client)
			return
		}
		if _medicalimagingUpdateImageSetMetadata {
			medicalimaging_UpdateImageSetMetadata(cfg, client)
			return
		}

	},
}

var (
	_medicalimagingCopyImageSet           bool
	_medicalimagingCreateDatastore        bool
	_medicalimagingDeleteDatastore        bool
	_medicalimagingDeleteImageSet         bool
	_medicalimagingGetDatastore           bool
	_medicalimagingGetDICOMImportJob      bool
	_medicalimagingGetImageFrame          bool
	_medicalimagingGetImageSet            bool
	_medicalimagingGetImageSetMetadata    bool
	_medicalimagingListDatastores         bool
	_medicalimagingListDICOMImportJobs    bool
	_medicalimagingListImageSetVersions   bool
	_medicalimagingListTagsForResource    bool
	_medicalimagingSearchImageSets        bool
	_medicalimagingStartDICOMImportJob    bool
	_medicalimagingTagResource            bool
	_medicalimagingUntagResource          bool
	_medicalimagingUpdateImageSetMetadata bool

	_medicalimagingClientToken                   string
	_medicalimagingCopyImageSetInformation       string
	_medicalimagingDataAccessRoleArn             string
	_medicalimagingDatastoreId                   string
	_medicalimagingDatastoreName                 string
	_medicalimagingDatastoreStatus               string
	_medicalimagingForce                         string
	_medicalimagingImageFrameInformation         string
	_medicalimagingImageSetId                    string
	_medicalimagingInputOwnerAccountId           string
	_medicalimagingInputS3Uri                    string
	_medicalimagingJobId                         string
	_medicalimagingJobName                       string
	_medicalimagingJobStatus                     string
	_medicalimagingKmsKeyArn                     string
	_medicalimagingLambdaAuthorizerArn           string
	_medicalimagingLatestVersionId               string
	_medicalimagingLosslessStorageFormat         string
	_medicalimagingMaxResults                    string
	_medicalimagingNextToken                     string
	_medicalimagingOutputS3Uri                   string
	_medicalimagingPromoteToPrimary              string
	_medicalimagingResourceArn                   string
	_medicalimagingSearchCriteria                string
	_medicalimagingSourceImageSetId              string
	_medicalimagingTagKeys                       []string
	_medicalimagingTags                          string
	_medicalimagingUpdateImageSetMetadataUpdates string
	_medicalimagingVersionId                     string
)

// Copy an image set.
func medicalimaging_CopyImageSet(cfg aws.Config, client *medicalimaging.Client) {
	input := &medicalimaging.CopyImageSetInput{
		// CopyImageSetInformation: *types.CopyImageSetInformation, // Required
		// DatastoreId: *string, // Required
		// SourceImageSetId: *string, // Required
	}

	if len(_medicalimagingCopyImageSetInformation) > 0 {
		if err := assignInputField(input, "CopyImageSetInformation", _medicalimagingCopyImageSetInformation); err != nil {
			log.Errorf("invalid --copy-image-set-information: %s", err.Error())
			return
		}
	}
	if len(_medicalimagingDatastoreId) > 0 {
		input.DatastoreId = aws.String(_medicalimagingDatastoreId)
	}
	if len(_medicalimagingSourceImageSetId) > 0 {
		input.SourceImageSetId = aws.String(_medicalimagingSourceImageSetId)
	}
	if len(_medicalimagingForce) > 0 {
		if err := assignInputField(input, "Force", _medicalimagingForce); err != nil {
			log.Errorf("invalid --force: %s", err.Error())
			return
		}
	}
	if len(_medicalimagingPromoteToPrimary) > 0 {
		if err := assignInputField(input, "PromoteToPrimary", _medicalimagingPromoteToPrimary); err != nil {
			log.Errorf("invalid --promote-to-primary: %s", err.Error())
			return
		}
	}

	if resp, err := client.CopyImageSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a data store.
func medicalimaging_CreateDatastore(cfg aws.Config, client *medicalimaging.Client) {
	input := &medicalimaging.CreateDatastoreInput{
		// ClientToken: *string, // Required
	}

	if len(_medicalimagingClientToken) > 0 {
		input.ClientToken = aws.String(_medicalimagingClientToken)
	}
	if len(_medicalimagingDatastoreName) > 0 {
		input.DatastoreName = aws.String(_medicalimagingDatastoreName)
	}
	if len(_medicalimagingKmsKeyArn) > 0 {
		input.KmsKeyArn = aws.String(_medicalimagingKmsKeyArn)
	}
	if len(_medicalimagingLambdaAuthorizerArn) > 0 {
		input.LambdaAuthorizerArn = aws.String(_medicalimagingLambdaAuthorizerArn)
	}
	if len(_medicalimagingLosslessStorageFormat) > 0 {
		if err := assignInputField(input, "LosslessStorageFormat", _medicalimagingLosslessStorageFormat); err != nil {
			log.Errorf("invalid --lossless-storage-format: %s", err.Error())
			return
		}
	}
	if len(_medicalimagingTags) > 0 {
		if err := assignInputField(input, "Tags", _medicalimagingTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDatastore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a data store.
// Before a data store can be deleted, you must first delete all image sets within
// it.
func medicalimaging_DeleteDatastore(cfg aws.Config, client *medicalimaging.Client) {
	input := &medicalimaging.DeleteDatastoreInput{
		// DatastoreId: *string, // Required
	}

	if len(_medicalimagingDatastoreId) > 0 {
		input.DatastoreId = aws.String(_medicalimagingDatastoreId)
	}

	if resp, err := client.DeleteDatastore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete an image set.
func medicalimaging_DeleteImageSet(cfg aws.Config, client *medicalimaging.Client) {
	input := &medicalimaging.DeleteImageSetInput{
		// DatastoreId: *string, // Required
		// ImageSetId: *string, // Required
	}

	if len(_medicalimagingDatastoreId) > 0 {
		input.DatastoreId = aws.String(_medicalimagingDatastoreId)
	}
	if len(_medicalimagingImageSetId) > 0 {
		input.ImageSetId = aws.String(_medicalimagingImageSetId)
	}

	if resp, err := client.DeleteImageSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get data store properties.
func medicalimaging_GetDatastore(cfg aws.Config, client *medicalimaging.Client) {
	input := &medicalimaging.GetDatastoreInput{
		// DatastoreId: *string, // Required
	}

	if len(_medicalimagingDatastoreId) > 0 {
		input.DatastoreId = aws.String(_medicalimagingDatastoreId)
	}

	if resp, err := client.GetDatastore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the import job properties to learn more about the job or job progress.
// The jobStatus refers to the execution of the import job. Therefore, an import
// job can return a jobStatus as COMPLETED even if validation issues are
// discovered during the import process. If a jobStatus returns as COMPLETED , we
// still recommend you review the output manifests written to S3, as they provide
// details on the success or failure of individual P10 object imports.
func medicalimaging_GetDICOMImportJob(cfg aws.Config, client *medicalimaging.Client) {
	input := &medicalimaging.GetDICOMImportJobInput{
		// DatastoreId: *string, // Required
		// JobId: *string, // Required
	}

	if len(_medicalimagingDatastoreId) > 0 {
		input.DatastoreId = aws.String(_medicalimagingDatastoreId)
	}
	if len(_medicalimagingJobId) > 0 {
		input.JobId = aws.String(_medicalimagingJobId)
	}

	if resp, err := client.GetDICOMImportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get an image frame (pixel data) for an image set.
func medicalimaging_GetImageFrame(cfg aws.Config, client *medicalimaging.Client) {
	input := &medicalimaging.GetImageFrameInput{
		// DatastoreId: *string, // Required
		// ImageFrameInformation: *types.ImageFrameInformation, // Required
		// ImageSetId: *string, // Required
	}

	if len(_medicalimagingDatastoreId) > 0 {
		input.DatastoreId = aws.String(_medicalimagingDatastoreId)
	}
	if len(_medicalimagingImageFrameInformation) > 0 {
		if err := assignInputField(input, "ImageFrameInformation", _medicalimagingImageFrameInformation); err != nil {
			log.Errorf("invalid --image-frame-information: %s", err.Error())
			return
		}
	}
	if len(_medicalimagingImageSetId) > 0 {
		input.ImageSetId = aws.String(_medicalimagingImageSetId)
	}

	if resp, err := client.GetImageFrame(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get image set properties.
func medicalimaging_GetImageSet(cfg aws.Config, client *medicalimaging.Client) {
	input := &medicalimaging.GetImageSetInput{
		// DatastoreId: *string, // Required
		// ImageSetId: *string, // Required
	}

	if len(_medicalimagingDatastoreId) > 0 {
		input.DatastoreId = aws.String(_medicalimagingDatastoreId)
	}
	if len(_medicalimagingImageSetId) > 0 {
		input.ImageSetId = aws.String(_medicalimagingImageSetId)
	}
	if len(_medicalimagingVersionId) > 0 {
		input.VersionId = aws.String(_medicalimagingVersionId)
	}

	if resp, err := client.GetImageSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get metadata attributes for an image set.
func medicalimaging_GetImageSetMetadata(cfg aws.Config, client *medicalimaging.Client) {
	input := &medicalimaging.GetImageSetMetadataInput{
		// DatastoreId: *string, // Required
		// ImageSetId: *string, // Required
	}

	if len(_medicalimagingDatastoreId) > 0 {
		input.DatastoreId = aws.String(_medicalimagingDatastoreId)
	}
	if len(_medicalimagingImageSetId) > 0 {
		input.ImageSetId = aws.String(_medicalimagingImageSetId)
	}
	if len(_medicalimagingVersionId) > 0 {
		input.VersionId = aws.String(_medicalimagingVersionId)
	}

	if resp, err := client.GetImageSetMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List data stores.
func medicalimaging_ListDatastores(cfg aws.Config, client *medicalimaging.Client) {
	input := &medicalimaging.ListDatastoresInput{}

	if len(_medicalimagingDatastoreStatus) > 0 {
		if err := assignInputField(input, "DatastoreStatus", _medicalimagingDatastoreStatus); err != nil {
			log.Errorf("invalid --datastore-status: %s", err.Error())
			return
		}
	}
	if len(_medicalimagingMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _medicalimagingMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_medicalimagingNextToken) > 0 {
		input.NextToken = aws.String(_medicalimagingNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDatastores(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*medicalimaging.ListDatastoresOutput
	p := medicalimaging.NewListDatastoresPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// List import jobs created for a specific data store.
func medicalimaging_ListDICOMImportJobs(cfg aws.Config, client *medicalimaging.Client) {
	input := &medicalimaging.ListDICOMImportJobsInput{
		// DatastoreId: *string, // Required
	}

	if len(_medicalimagingDatastoreId) > 0 {
		input.DatastoreId = aws.String(_medicalimagingDatastoreId)
	}
	if len(_medicalimagingJobStatus) > 0 {
		if err := assignInputField(input, "JobStatus", _medicalimagingJobStatus); err != nil {
			log.Errorf("invalid --job-status: %s", err.Error())
			return
		}
	}
	if len(_medicalimagingMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _medicalimagingMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_medicalimagingNextToken) > 0 {
		input.NextToken = aws.String(_medicalimagingNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDICOMImportJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*medicalimaging.ListDICOMImportJobsOutput
	p := medicalimaging.NewListDICOMImportJobsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// List image set versions.
func medicalimaging_ListImageSetVersions(cfg aws.Config, client *medicalimaging.Client) {
	input := &medicalimaging.ListImageSetVersionsInput{
		// DatastoreId: *string, // Required
		// ImageSetId: *string, // Required
	}

	if len(_medicalimagingDatastoreId) > 0 {
		input.DatastoreId = aws.String(_medicalimagingDatastoreId)
	}
	if len(_medicalimagingImageSetId) > 0 {
		input.ImageSetId = aws.String(_medicalimagingImageSetId)
	}
	if len(_medicalimagingMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _medicalimagingMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_medicalimagingNextToken) > 0 {
		input.NextToken = aws.String(_medicalimagingNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListImageSetVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*medicalimaging.ListImageSetVersionsOutput
	p := medicalimaging.NewListImageSetVersionsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Lists all tags associated with a medical imaging resource.
func medicalimaging_ListTagsForResource(cfg aws.Config, client *medicalimaging.Client) {
	input := &medicalimaging.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_medicalimagingResourceArn) > 0 {
		input.ResourceArn = aws.String(_medicalimagingResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Search image sets based on defined input attributes.
// SearchImageSets accepts a single search query parameter and returns a paginated
// response of all image sets that have the matching criteria. All date range
// queries must be input as (lowerBound, upperBound) .
//
// By default, SearchImageSets uses the updatedAt field for sorting in descending
// order from newest to oldest.
func medicalimaging_SearchImageSets(cfg aws.Config, client *medicalimaging.Client) {
	input := &medicalimaging.SearchImageSetsInput{
		// DatastoreId: *string, // Required
	}

	if len(_medicalimagingDatastoreId) > 0 {
		input.DatastoreId = aws.String(_medicalimagingDatastoreId)
	}
	if len(_medicalimagingMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _medicalimagingMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_medicalimagingNextToken) > 0 {
		input.NextToken = aws.String(_medicalimagingNextToken)
	}
	if len(_medicalimagingSearchCriteria) > 0 {
		if err := assignInputField(input, "SearchCriteria", _medicalimagingSearchCriteria); err != nil {
			log.Errorf("invalid --search-criteria: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.SearchImageSets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*medicalimaging.SearchImageSetsOutput
	p := medicalimaging.NewSearchImageSetsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Start importing bulk data into an ACTIVE data store. The import job imports
// DICOM P10 files found in the S3 prefix specified by the inputS3Uri parameter.
// The import job stores processing results in the file specified by the
// outputS3Uri parameter.
func medicalimaging_StartDICOMImportJob(cfg aws.Config, client *medicalimaging.Client) {
	input := &medicalimaging.StartDICOMImportJobInput{
		// ClientToken: *string, // Required
		// DataAccessRoleArn: *string, // Required
		// DatastoreId: *string, // Required
		// InputS3Uri: *string, // Required
		// OutputS3Uri: *string, // Required
	}

	if len(_medicalimagingClientToken) > 0 {
		input.ClientToken = aws.String(_medicalimagingClientToken)
	}
	if len(_medicalimagingDataAccessRoleArn) > 0 {
		input.DataAccessRoleArn = aws.String(_medicalimagingDataAccessRoleArn)
	}
	if len(_medicalimagingDatastoreId) > 0 {
		input.DatastoreId = aws.String(_medicalimagingDatastoreId)
	}
	if len(_medicalimagingInputS3Uri) > 0 {
		input.InputS3Uri = aws.String(_medicalimagingInputS3Uri)
	}
	if len(_medicalimagingOutputS3Uri) > 0 {
		input.OutputS3Uri = aws.String(_medicalimagingOutputS3Uri)
	}
	if len(_medicalimagingInputOwnerAccountId) > 0 {
		input.InputOwnerAccountId = aws.String(_medicalimagingInputOwnerAccountId)
	}
	if len(_medicalimagingJobName) > 0 {
		input.JobName = aws.String(_medicalimagingJobName)
	}

	if resp, err := client.StartDICOMImportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a user-specifed key and value tag to a medical imaging resource.
func medicalimaging_TagResource(cfg aws.Config, client *medicalimaging.Client) {
	input := &medicalimaging.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_medicalimagingResourceArn) > 0 {
		input.ResourceArn = aws.String(_medicalimagingResourceArn)
	}
	if len(_medicalimagingTags) > 0 {
		if err := assignInputField(input, "Tags", _medicalimagingTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.TagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes tags from a medical imaging resource.
func medicalimaging_UntagResource(cfg aws.Config, client *medicalimaging.Client) {
	input := &medicalimaging.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_medicalimagingResourceArn) > 0 {
		input.ResourceArn = aws.String(_medicalimagingResourceArn)
	}
	if len(_medicalimagingTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _medicalimagingTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update image set metadata attributes.
func medicalimaging_UpdateImageSetMetadata(cfg aws.Config, client *medicalimaging.Client) {
	input := &medicalimaging.UpdateImageSetMetadataInput{
		// DatastoreId: *string, // Required
		// ImageSetId: *string, // Required
		// LatestVersionId: *string, // Required
		// UpdateImageSetMetadataUpdates: types.MetadataUpdates, // Required
	}

	if len(_medicalimagingDatastoreId) > 0 {
		input.DatastoreId = aws.String(_medicalimagingDatastoreId)
	}
	if len(_medicalimagingImageSetId) > 0 {
		input.ImageSetId = aws.String(_medicalimagingImageSetId)
	}
	if len(_medicalimagingLatestVersionId) > 0 {
		input.LatestVersionId = aws.String(_medicalimagingLatestVersionId)
	}
	if len(_medicalimagingUpdateImageSetMetadataUpdates) > 0 {
		if err := assignInputField(input, "UpdateImageSetMetadataUpdates", _medicalimagingUpdateImageSetMetadataUpdates); err != nil {
			log.Errorf("invalid --update-image-set-metadata-updates: %s", err.Error())
			return
		}
	}
	if len(_medicalimagingForce) > 0 {
		if err := assignInputField(input, "Force", _medicalimagingForce); err != nil {
			log.Errorf("invalid --force: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateImageSetMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_medicalimagingCmd)
	_medicalimagingCmd.Flags().SortFlags = false

	_medicalimagingCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_medicalimagingCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_medicalimagingCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_medicalimagingCmd.Flags().StringVarP(&_medicalimagingClientToken, "client-token", "", "", "Client Token")
	_medicalimagingCmd.Flags().StringVarP(&_medicalimagingCopyImageSetInformation, "copy-image-set-information", "", "", "Copy Image Set Information")
	_medicalimagingCmd.Flags().StringVarP(&_medicalimagingDataAccessRoleArn, "data-access-role-arn", "", "", "Data Access Role ARN")
	_medicalimagingCmd.Flags().StringVarP(&_medicalimagingDatastoreId, "datastore-id", "", "", "Datastore ID")
	_medicalimagingCmd.Flags().StringVarP(&_medicalimagingDatastoreName, "datastore-name", "", "", "Datastore Name")
	_medicalimagingCmd.Flags().StringVarP(&_medicalimagingDatastoreStatus, "datastore-status", "", "", "Datastore Status")
	_medicalimagingCmd.Flags().StringVarP(&_medicalimagingForce, "force", "", "", "Force")
	_medicalimagingCmd.Flags().StringVarP(&_medicalimagingImageFrameInformation, "image-frame-information", "", "", "Image Frame Information")
	_medicalimagingCmd.Flags().StringVarP(&_medicalimagingImageSetId, "image-set-id", "", "", "Image Set ID")
	_medicalimagingCmd.Flags().StringVarP(&_medicalimagingInputOwnerAccountId, "input-owner-account-id", "", "", "Input Owner Account ID")
	_medicalimagingCmd.Flags().StringVarP(&_medicalimagingInputS3Uri, "input-s3-uri", "", "", "Input S3 URI")
	_medicalimagingCmd.Flags().StringVarP(&_medicalimagingJobId, "job-id", "", "", "Job ID")
	_medicalimagingCmd.Flags().StringVarP(&_medicalimagingJobName, "job-name", "", "", "Job Name")
	_medicalimagingCmd.Flags().StringVarP(&_medicalimagingJobStatus, "job-status", "", "", "Job Status")
	_medicalimagingCmd.Flags().StringVarP(&_medicalimagingKmsKeyArn, "kms-key-arn", "", "", "KMS Key ARN")
	_medicalimagingCmd.Flags().StringVarP(&_medicalimagingLambdaAuthorizerArn, "lambda-authorizer-arn", "", "", "Lambda Authorizer ARN")
	_medicalimagingCmd.Flags().StringVarP(&_medicalimagingLatestVersionId, "latest-version-id", "", "", "Latest Version ID")
	_medicalimagingCmd.Flags().StringVarP(&_medicalimagingLosslessStorageFormat, "lossless-storage-format", "", "", "Lossless Storage Format")
	_medicalimagingCmd.Flags().StringVarP(&_medicalimagingMaxResults, "max-results", "", "", "Max Results")
	_medicalimagingCmd.Flags().StringVarP(&_medicalimagingNextToken, "next-token", "", "", "Next Token")
	_medicalimagingCmd.Flags().StringVarP(&_medicalimagingOutputS3Uri, "output-s3-uri", "", "", "Output S3 URI")
	_medicalimagingCmd.Flags().StringVarP(&_medicalimagingPromoteToPrimary, "promote-to-primary", "", "", "Promote To Primary")
	_medicalimagingCmd.Flags().StringVarP(&_medicalimagingResourceArn, "resource-arn", "", "", "Resource ARN")
	_medicalimagingCmd.Flags().StringVarP(&_medicalimagingSearchCriteria, "search-criteria", "", "", "Search Criteria")
	_medicalimagingCmd.Flags().StringVarP(&_medicalimagingSourceImageSetId, "source-image-set-id", "", "", "Source Image Set ID")
	_medicalimagingCmd.Flags().StringSliceVarP(&_medicalimagingTagKeys, "tag-keys", "", nil, "Tag Keys")
	_medicalimagingCmd.Flags().StringVarP(&_medicalimagingTags, "tags", "", "", "Tags")
	_medicalimagingCmd.Flags().StringVarP(&_medicalimagingUpdateImageSetMetadataUpdates, "update-image-set-metadata-updates", "", "", "Update Image Set Metadata Updates")
	_medicalimagingCmd.Flags().StringVarP(&_medicalimagingVersionId, "version-id", "", "", "Version ID")

	_medicalimagingCmd.Flags().BoolVarP(&_medicalimagingCopyImageSet, "copy-image-set", "", false, "Copy Image Set")
	_medicalimagingCmd.Flags().BoolVarP(&_medicalimagingCreateDatastore, "create-datastore", "", false, "Create Datastore")
	_medicalimagingCmd.Flags().BoolVarP(&_medicalimagingDeleteDatastore, "delete-datastore", "", false, "Delete Datastore")
	_medicalimagingCmd.Flags().BoolVarP(&_medicalimagingDeleteImageSet, "delete-image-set", "", false, "Delete Image Set")
	_medicalimagingCmd.Flags().BoolVarP(&_medicalimagingGetDatastore, "get-datastore", "", false, "Get Datastore")
	_medicalimagingCmd.Flags().BoolVarP(&_medicalimagingGetDICOMImportJob, "get-dicom-import-job", "", false, "Get Dicom Import Job")
	_medicalimagingCmd.Flags().BoolVarP(&_medicalimagingGetImageFrame, "get-image-frame", "", false, "Get Image Frame")
	_medicalimagingCmd.Flags().BoolVarP(&_medicalimagingGetImageSet, "get-image-set", "", false, "Get Image Set")
	_medicalimagingCmd.Flags().BoolVarP(&_medicalimagingGetImageSetMetadata, "get-image-set-metadata", "", false, "Get Image Set Metadata")
	_medicalimagingCmd.Flags().BoolVarP(&_medicalimagingListDatastores, "list-datastores", "", false, "List Datastores")
	_medicalimagingCmd.Flags().BoolVarP(&_medicalimagingListDICOMImportJobs, "list-dicom-import-jobs", "", false, "List Dicom Import Jobs")
	_medicalimagingCmd.Flags().BoolVarP(&_medicalimagingListImageSetVersions, "list-image-set-versions", "", false, "List Image Set Versions")
	_medicalimagingCmd.Flags().BoolVarP(&_medicalimagingListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_medicalimagingCmd.Flags().BoolVarP(&_medicalimagingSearchImageSets, "search-image-sets", "", false, "Search Image Sets")
	_medicalimagingCmd.Flags().BoolVarP(&_medicalimagingStartDICOMImportJob, "start-dicom-import-job", "", false, "Start Dicom Import Job")
	_medicalimagingCmd.Flags().BoolVarP(&_medicalimagingTagResource, "tag-resource", "", false, "Tag Resource")
	_medicalimagingCmd.Flags().BoolVarP(&_medicalimagingUntagResource, "untag-resource", "", false, "Untag Resource")
	_medicalimagingCmd.Flags().BoolVarP(&_medicalimagingUpdateImageSetMetadata, "update-image-set-metadata", "", false, "Update Image Set Metadata")

}
