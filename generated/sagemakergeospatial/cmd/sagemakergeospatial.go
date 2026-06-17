package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sagemakergeospatial"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// sagemakergeospatialCmd represents the sagemakergeospatial command
var _sagemakergeospatialCmd = &cobra.Command{
	Use:   "sagemakergeospatial",
	Short: "AWS sagemakergeospatial CLI",
	Run: func(cmd *cobra.Command, args []string) {
		_awsOutput = resolveAWSOutput(_awsProfile, cmd.Flags().Changed("output"))
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := sagemakergeospatial.NewFromConfig(cfg)
		if _sagemakergeospatialDeleteEarthObservationJob {
			sagemakergeospatial_DeleteEarthObservationJob(cfg, client)
			return
		}
		if _sagemakergeospatialDeleteVectorEnrichmentJob {
			sagemakergeospatial_DeleteVectorEnrichmentJob(cfg, client)
			return
		}
		if _sagemakergeospatialExportEarthObservationJob {
			sagemakergeospatial_ExportEarthObservationJob(cfg, client)
			return
		}
		if _sagemakergeospatialExportVectorEnrichmentJob {
			sagemakergeospatial_ExportVectorEnrichmentJob(cfg, client)
			return
		}
		if _sagemakergeospatialGetEarthObservationJob {
			sagemakergeospatial_GetEarthObservationJob(cfg, client)
			return
		}
		if _sagemakergeospatialGetRasterDataCollection {
			sagemakergeospatial_GetRasterDataCollection(cfg, client)
			return
		}
		if _sagemakergeospatialGetTile {
			sagemakergeospatial_GetTile(cfg, client)
			return
		}
		if _sagemakergeospatialGetVectorEnrichmentJob {
			sagemakergeospatial_GetVectorEnrichmentJob(cfg, client)
			return
		}
		if _sagemakergeospatialListEarthObservationJobs {
			sagemakergeospatial_ListEarthObservationJobs(cfg, client)
			return
		}
		if _sagemakergeospatialListRasterDataCollections {
			sagemakergeospatial_ListRasterDataCollections(cfg, client)
			return
		}
		if _sagemakergeospatialListTagsForResource {
			sagemakergeospatial_ListTagsForResource(cfg, client)
			return
		}
		if _sagemakergeospatialListVectorEnrichmentJobs {
			sagemakergeospatial_ListVectorEnrichmentJobs(cfg, client)
			return
		}
		if _sagemakergeospatialSearchRasterDataCollection {
			sagemakergeospatial_SearchRasterDataCollection(cfg, client)
			return
		}
		if _sagemakergeospatialStartEarthObservationJob {
			sagemakergeospatial_StartEarthObservationJob(cfg, client)
			return
		}
		if _sagemakergeospatialStartVectorEnrichmentJob {
			sagemakergeospatial_StartVectorEnrichmentJob(cfg, client)
			return
		}
		if _sagemakergeospatialStopEarthObservationJob {
			sagemakergeospatial_StopEarthObservationJob(cfg, client)
			return
		}
		if _sagemakergeospatialStopVectorEnrichmentJob {
			sagemakergeospatial_StopVectorEnrichmentJob(cfg, client)
			return
		}
		if _sagemakergeospatialTagResource {
			sagemakergeospatial_TagResource(cfg, client)
			return
		}
		if _sagemakergeospatialUntagResource {
			sagemakergeospatial_UntagResource(cfg, client)
			return
		}

	},
}

var (
	_sagemakergeospatialDeleteEarthObservationJob  bool
	_sagemakergeospatialDeleteVectorEnrichmentJob  bool
	_sagemakergeospatialExportEarthObservationJob  bool
	_sagemakergeospatialExportVectorEnrichmentJob  bool
	_sagemakergeospatialGetEarthObservationJob     bool
	_sagemakergeospatialGetRasterDataCollection    bool
	_sagemakergeospatialGetTile                    bool
	_sagemakergeospatialGetVectorEnrichmentJob     bool
	_sagemakergeospatialListEarthObservationJobs   bool
	_sagemakergeospatialListRasterDataCollections  bool
	_sagemakergeospatialListTagsForResource        bool
	_sagemakergeospatialListVectorEnrichmentJobs   bool
	_sagemakergeospatialSearchRasterDataCollection bool
	_sagemakergeospatialStartEarthObservationJob   bool
	_sagemakergeospatialStartVectorEnrichmentJob   bool
	_sagemakergeospatialStopEarthObservationJob    bool
	_sagemakergeospatialStopVectorEnrichmentJob    bool
	_sagemakergeospatialTagResource                bool
	_sagemakergeospatialUntagResource              bool

	_sagemakergeospatialArn                       string
	_sagemakergeospatialClientToken               string
	_sagemakergeospatialExecutionRoleArn          string
	_sagemakergeospatialExportSourceImages        string
	_sagemakergeospatialImageAssets               []string
	_sagemakergeospatialImageMask                 string
	_sagemakergeospatialInputConfig               string
	_sagemakergeospatialJobConfig                 string
	_sagemakergeospatialKmsKeyId                  string
	_sagemakergeospatialMaxResults                string
	_sagemakergeospatialName                      string
	_sagemakergeospatialNextToken                 string
	_sagemakergeospatialOutputConfig              string
	_sagemakergeospatialOutputDataType            string
	_sagemakergeospatialOutputFormat              string
	_sagemakergeospatialPropertyFilters           string
	_sagemakergeospatialRasterDataCollectionQuery string
	_sagemakergeospatialResourceArn               string
	_sagemakergeospatialSortBy                    string
	_sagemakergeospatialSortOrder                 string
	_sagemakergeospatialStatusEquals              string
	_sagemakergeospatialTagKeys                   []string
	_sagemakergeospatialTags                      string
	_sagemakergeospatialTarget                    string
	_sagemakergeospatialTimeRangeFilter           string
	_sagemakergeospatialX                         string
	_sagemakergeospatialY                         string
	_sagemakergeospatialZ                         string
)

// Use this operation to delete an Earth Observation job.
func sagemakergeospatial_DeleteEarthObservationJob(cfg aws.Config, client *sagemakergeospatial.Client) {
	input := &sagemakergeospatial.DeleteEarthObservationJobInput{
		// Arn: *string, // Required
	}

	if len(_sagemakergeospatialArn) > 0 {
		input.Arn = aws.String(_sagemakergeospatialArn)
	}

	if resp, err := client.DeleteEarthObservationJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use this operation to delete a Vector Enrichment job.
func sagemakergeospatial_DeleteVectorEnrichmentJob(cfg aws.Config, client *sagemakergeospatial.Client) {
	input := &sagemakergeospatial.DeleteVectorEnrichmentJobInput{
		// Arn: *string, // Required
	}

	if len(_sagemakergeospatialArn) > 0 {
		input.Arn = aws.String(_sagemakergeospatialArn)
	}

	if resp, err := client.DeleteVectorEnrichmentJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use this operation to export results of an Earth Observation job and optionally
// source images used as input to the EOJ to an Amazon S3 location.
func sagemakergeospatial_ExportEarthObservationJob(cfg aws.Config, client *sagemakergeospatial.Client) {
	input := &sagemakergeospatial.ExportEarthObservationJobInput{
		// Arn: *string, // Required
		// ExecutionRoleArn: *string, // Required
		// OutputConfig: *types.OutputConfigInput, // Required
	}

	if len(_sagemakergeospatialArn) > 0 {
		input.Arn = aws.String(_sagemakergeospatialArn)
	}
	if len(_sagemakergeospatialExecutionRoleArn) > 0 {
		input.ExecutionRoleArn = aws.String(_sagemakergeospatialExecutionRoleArn)
	}
	if len(_sagemakergeospatialOutputConfig) > 0 {
		if err := assignInputField(input, "OutputConfig", _sagemakergeospatialOutputConfig); err != nil {
			log.Errorf("invalid --output-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakergeospatialClientToken) > 0 {
		input.ClientToken = aws.String(_sagemakergeospatialClientToken)
	}
	if len(_sagemakergeospatialExportSourceImages) > 0 {
		if err := assignInputField(input, "ExportSourceImages", _sagemakergeospatialExportSourceImages); err != nil {
			log.Errorf("invalid --export-source-images: %s", err.Error())
			return
		}
	}

	if resp, err := client.ExportEarthObservationJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use this operation to copy results of a Vector Enrichment job to an Amazon S3
// location.
func sagemakergeospatial_ExportVectorEnrichmentJob(cfg aws.Config, client *sagemakergeospatial.Client) {
	input := &sagemakergeospatial.ExportVectorEnrichmentJobInput{
		// Arn: *string, // Required
		// ExecutionRoleArn: *string, // Required
		// OutputConfig: *types.ExportVectorEnrichmentJobOutputConfig, // Required
	}

	if len(_sagemakergeospatialArn) > 0 {
		input.Arn = aws.String(_sagemakergeospatialArn)
	}
	if len(_sagemakergeospatialExecutionRoleArn) > 0 {
		input.ExecutionRoleArn = aws.String(_sagemakergeospatialExecutionRoleArn)
	}
	if len(_sagemakergeospatialOutputConfig) > 0 {
		if err := assignInputField(input, "OutputConfig", _sagemakergeospatialOutputConfig); err != nil {
			log.Errorf("invalid --output-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakergeospatialClientToken) > 0 {
		input.ClientToken = aws.String(_sagemakergeospatialClientToken)
	}

	if resp, err := client.ExportVectorEnrichmentJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the details for a previously initiated Earth Observation job.
func sagemakergeospatial_GetEarthObservationJob(cfg aws.Config, client *sagemakergeospatial.Client) {
	input := &sagemakergeospatial.GetEarthObservationJobInput{
		// Arn: *string, // Required
	}

	if len(_sagemakergeospatialArn) > 0 {
		input.Arn = aws.String(_sagemakergeospatialArn)
	}

	if resp, err := client.GetEarthObservationJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use this operation to get details of a specific raster data collection.
func sagemakergeospatial_GetRasterDataCollection(cfg aws.Config, client *sagemakergeospatial.Client) {
	input := &sagemakergeospatial.GetRasterDataCollectionInput{
		// Arn: *string, // Required
	}

	if len(_sagemakergeospatialArn) > 0 {
		input.Arn = aws.String(_sagemakergeospatialArn)
	}

	if resp, err := client.GetRasterDataCollection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a web mercator tile for the given Earth Observation job.
func sagemakergeospatial_GetTile(cfg aws.Config, client *sagemakergeospatial.Client) {
	input := &sagemakergeospatial.GetTileInput{
		// Arn: *string, // Required
		// ImageAssets: []string, // Required
		// Target: types.TargetOptions, // Required
		// X: *int32, // Required
		// Y: *int32, // Required
		// Z: *int32, // Required
	}

	if len(_sagemakergeospatialArn) > 0 {
		input.Arn = aws.String(_sagemakergeospatialArn)
	}
	if len(_sagemakergeospatialImageAssets) > 0 {
		input.ImageAssets = append([]string(nil), _sagemakergeospatialImageAssets...)
	}
	if len(_sagemakergeospatialTarget) > 0 {
		if err := assignInputField(input, "Target", _sagemakergeospatialTarget); err != nil {
			log.Errorf("invalid --target: %s", err.Error())
			return
		}
	}
	if len(_sagemakergeospatialX) > 0 {
		if err := assignInputField(input, "X", _sagemakergeospatialX); err != nil {
			log.Errorf("invalid --x: %s", err.Error())
			return
		}
	}
	if len(_sagemakergeospatialY) > 0 {
		if err := assignInputField(input, "Y", _sagemakergeospatialY); err != nil {
			log.Errorf("invalid --y: %s", err.Error())
			return
		}
	}
	if len(_sagemakergeospatialZ) > 0 {
		if err := assignInputField(input, "Z", _sagemakergeospatialZ); err != nil {
			log.Errorf("invalid --z: %s", err.Error())
			return
		}
	}
	if len(_sagemakergeospatialExecutionRoleArn) > 0 {
		input.ExecutionRoleArn = aws.String(_sagemakergeospatialExecutionRoleArn)
	}
	if len(_sagemakergeospatialImageMask) > 0 {
		if err := assignInputField(input, "ImageMask", _sagemakergeospatialImageMask); err != nil {
			log.Errorf("invalid --image-mask: %s", err.Error())
			return
		}
	}
	if len(_sagemakergeospatialOutputDataType) > 0 {
		if err := assignInputField(input, "OutputDataType", _sagemakergeospatialOutputDataType); err != nil {
			log.Errorf("invalid --output-data-type: %s", err.Error())
			return
		}
	}
	if len(_sagemakergeospatialOutputFormat) > 0 {
		input.OutputFormat = aws.String(_sagemakergeospatialOutputFormat)
	}
	if len(_sagemakergeospatialPropertyFilters) > 0 {
		input.PropertyFilters = aws.String(_sagemakergeospatialPropertyFilters)
	}
	if len(_sagemakergeospatialTimeRangeFilter) > 0 {
		input.TimeRangeFilter = aws.String(_sagemakergeospatialTimeRangeFilter)
	}

	if resp, err := client.GetTile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details of a Vector Enrichment Job for a given job Amazon Resource
// Name (ARN).
func sagemakergeospatial_GetVectorEnrichmentJob(cfg aws.Config, client *sagemakergeospatial.Client) {
	input := &sagemakergeospatial.GetVectorEnrichmentJobInput{
		// Arn: *string, // Required
	}

	if len(_sagemakergeospatialArn) > 0 {
		input.Arn = aws.String(_sagemakergeospatialArn)
	}

	if resp, err := client.GetVectorEnrichmentJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use this operation to get a list of the Earth Observation jobs associated with
// the calling Amazon Web Services account.
func sagemakergeospatial_ListEarthObservationJobs(cfg aws.Config, client *sagemakergeospatial.Client) {
	input := &sagemakergeospatial.ListEarthObservationJobsInput{}

	if len(_sagemakergeospatialMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakergeospatialMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakergeospatialNextToken) > 0 {
		input.NextToken = aws.String(_sagemakergeospatialNextToken)
	}
	if len(_sagemakergeospatialSortBy) > 0 {
		input.SortBy = aws.String(_sagemakergeospatialSortBy)
	}
	if len(_sagemakergeospatialSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakergeospatialSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_sagemakergeospatialStatusEquals) > 0 {
		if err := assignInputField(input, "StatusEquals", _sagemakergeospatialStatusEquals); err != nil {
			log.Errorf("invalid --status-equals: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListEarthObservationJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemakergeospatial.ListEarthObservationJobsOutput
	p := sagemakergeospatial.NewListEarthObservationJobsPaginator(client, input)
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

// Use this operation to get raster data collections.
func sagemakergeospatial_ListRasterDataCollections(cfg aws.Config, client *sagemakergeospatial.Client) {
	input := &sagemakergeospatial.ListRasterDataCollectionsInput{}

	if len(_sagemakergeospatialMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakergeospatialMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakergeospatialNextToken) > 0 {
		input.NextToken = aws.String(_sagemakergeospatialNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRasterDataCollections(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemakergeospatial.ListRasterDataCollectionsOutput
	p := sagemakergeospatial.NewListRasterDataCollectionsPaginator(client, input)
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

// Lists the tags attached to the resource.
func sagemakergeospatial_ListTagsForResource(cfg aws.Config, client *sagemakergeospatial.Client) {
	input := &sagemakergeospatial.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_sagemakergeospatialResourceArn) > 0 {
		input.ResourceArn = aws.String(_sagemakergeospatialResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of vector enrichment jobs.
func sagemakergeospatial_ListVectorEnrichmentJobs(cfg aws.Config, client *sagemakergeospatial.Client) {
	input := &sagemakergeospatial.ListVectorEnrichmentJobsInput{}

	if len(_sagemakergeospatialMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakergeospatialMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakergeospatialNextToken) > 0 {
		input.NextToken = aws.String(_sagemakergeospatialNextToken)
	}
	if len(_sagemakergeospatialSortBy) > 0 {
		input.SortBy = aws.String(_sagemakergeospatialSortBy)
	}
	if len(_sagemakergeospatialSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakergeospatialSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_sagemakergeospatialStatusEquals) > 0 {
		input.StatusEquals = aws.String(_sagemakergeospatialStatusEquals)
	}

	if disablePaginator() {
		if resp, err := client.ListVectorEnrichmentJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemakergeospatial.ListVectorEnrichmentJobsOutput
	p := sagemakergeospatial.NewListVectorEnrichmentJobsPaginator(client, input)
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

// Allows you run image query on a specific raster data collection to get a list
// of the satellite imagery matching the selected filters.
func sagemakergeospatial_SearchRasterDataCollection(cfg aws.Config, client *sagemakergeospatial.Client) {
	input := &sagemakergeospatial.SearchRasterDataCollectionInput{
		// Arn: *string, // Required
		// RasterDataCollectionQuery: *types.RasterDataCollectionQueryWithBandFilterInput, // Required
	}

	if len(_sagemakergeospatialArn) > 0 {
		input.Arn = aws.String(_sagemakergeospatialArn)
	}
	if len(_sagemakergeospatialRasterDataCollectionQuery) > 0 {
		if err := assignInputField(input, "RasterDataCollectionQuery", _sagemakergeospatialRasterDataCollectionQuery); err != nil {
			log.Errorf("invalid --raster-data-collection-query: %s", err.Error())
			return
		}
	}
	if len(_sagemakergeospatialNextToken) > 0 {
		input.NextToken = aws.String(_sagemakergeospatialNextToken)
	}

	if disablePaginator() {
		if resp, err := client.SearchRasterDataCollection(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemakergeospatial.SearchRasterDataCollectionOutput
	p := sagemakergeospatial.NewSearchRasterDataCollectionPaginator(client, input)
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

// Use this operation to create an Earth observation job.
func sagemakergeospatial_StartEarthObservationJob(cfg aws.Config, client *sagemakergeospatial.Client) {
	input := &sagemakergeospatial.StartEarthObservationJobInput{
		// ExecutionRoleArn: *string, // Required
		// InputConfig: *types.InputConfigInput, // Required
		// JobConfig: types.JobConfigInput, // Required
		// Name: *string, // Required
	}

	if len(_sagemakergeospatialExecutionRoleArn) > 0 {
		input.ExecutionRoleArn = aws.String(_sagemakergeospatialExecutionRoleArn)
	}
	if len(_sagemakergeospatialInputConfig) > 0 {
		if err := assignInputField(input, "InputConfig", _sagemakergeospatialInputConfig); err != nil {
			log.Errorf("invalid --input-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakergeospatialJobConfig) > 0 {
		if err := assignInputField(input, "JobConfig", _sagemakergeospatialJobConfig); err != nil {
			log.Errorf("invalid --job-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakergeospatialName) > 0 {
		input.Name = aws.String(_sagemakergeospatialName)
	}
	if len(_sagemakergeospatialClientToken) > 0 {
		input.ClientToken = aws.String(_sagemakergeospatialClientToken)
	}
	if len(_sagemakergeospatialKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_sagemakergeospatialKmsKeyId)
	}
	if len(_sagemakergeospatialTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakergeospatialTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartEarthObservationJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a Vector Enrichment job for the supplied job type. Currently, there are
// two supported job types: reverse geocoding and map matching.
func sagemakergeospatial_StartVectorEnrichmentJob(cfg aws.Config, client *sagemakergeospatial.Client) {
	input := &sagemakergeospatial.StartVectorEnrichmentJobInput{
		// ExecutionRoleArn: *string, // Required
		// InputConfig: *types.VectorEnrichmentJobInputConfig, // Required
		// JobConfig: types.VectorEnrichmentJobConfig, // Required
		// Name: *string, // Required
	}

	if len(_sagemakergeospatialExecutionRoleArn) > 0 {
		input.ExecutionRoleArn = aws.String(_sagemakergeospatialExecutionRoleArn)
	}
	if len(_sagemakergeospatialInputConfig) > 0 {
		if err := assignInputField(input, "InputConfig", _sagemakergeospatialInputConfig); err != nil {
			log.Errorf("invalid --input-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakergeospatialJobConfig) > 0 {
		if err := assignInputField(input, "JobConfig", _sagemakergeospatialJobConfig); err != nil {
			log.Errorf("invalid --job-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakergeospatialName) > 0 {
		input.Name = aws.String(_sagemakergeospatialName)
	}
	if len(_sagemakergeospatialClientToken) > 0 {
		input.ClientToken = aws.String(_sagemakergeospatialClientToken)
	}
	if len(_sagemakergeospatialKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_sagemakergeospatialKmsKeyId)
	}
	if len(_sagemakergeospatialTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakergeospatialTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartVectorEnrichmentJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use this operation to stop an existing earth observation job.
func sagemakergeospatial_StopEarthObservationJob(cfg aws.Config, client *sagemakergeospatial.Client) {
	input := &sagemakergeospatial.StopEarthObservationJobInput{
		// Arn: *string, // Required
	}

	if len(_sagemakergeospatialArn) > 0 {
		input.Arn = aws.String(_sagemakergeospatialArn)
	}

	if resp, err := client.StopEarthObservationJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops the Vector Enrichment job for a given job ARN.
func sagemakergeospatial_StopVectorEnrichmentJob(cfg aws.Config, client *sagemakergeospatial.Client) {
	input := &sagemakergeospatial.StopVectorEnrichmentJobInput{
		// Arn: *string, // Required
	}

	if len(_sagemakergeospatialArn) > 0 {
		input.Arn = aws.String(_sagemakergeospatialArn)
	}

	if resp, err := client.StopVectorEnrichmentJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The resource you want to tag.
func sagemakergeospatial_TagResource(cfg aws.Config, client *sagemakergeospatial.Client) {
	input := &sagemakergeospatial.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_sagemakergeospatialResourceArn) > 0 {
		input.ResourceArn = aws.String(_sagemakergeospatialResourceArn)
	}
	if len(_sagemakergeospatialTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakergeospatialTags); err != nil {
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

// The resource you want to untag.
func sagemakergeospatial_UntagResource(cfg aws.Config, client *sagemakergeospatial.Client) {
	input := &sagemakergeospatial.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_sagemakergeospatialResourceArn) > 0 {
		input.ResourceArn = aws.String(_sagemakergeospatialResourceArn)
	}
	if len(_sagemakergeospatialTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _sagemakergeospatialTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_sagemakergeospatialCmd)
	_sagemakergeospatialCmd.Flags().SortFlags = false

	_sagemakergeospatialCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_sagemakergeospatialCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_sagemakergeospatialCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_sagemakergeospatialCmd.Flags().StringVarP(&_sagemakergeospatialArn, "arn", "", "", "ARN")
	_sagemakergeospatialCmd.Flags().StringVarP(&_sagemakergeospatialClientToken, "client-token", "", "", "Client Token")
	_sagemakergeospatialCmd.Flags().StringVarP(&_sagemakergeospatialExecutionRoleArn, "execution-role-arn", "", "", "Execution Role ARN")
	_sagemakergeospatialCmd.Flags().StringVarP(&_sagemakergeospatialExportSourceImages, "export-source-images", "", "", "Export Source Images")
	_sagemakergeospatialCmd.Flags().StringSliceVarP(&_sagemakergeospatialImageAssets, "image-assets", "", nil, "Image Assets")
	_sagemakergeospatialCmd.Flags().StringVarP(&_sagemakergeospatialImageMask, "image-mask", "", "", "Image Mask")
	_sagemakergeospatialCmd.Flags().StringVarP(&_sagemakergeospatialInputConfig, "input-config", "", "", "Input Config")
	_sagemakergeospatialCmd.Flags().StringVarP(&_sagemakergeospatialJobConfig, "job-config", "", "", "Job Config")
	_sagemakergeospatialCmd.Flags().StringVarP(&_sagemakergeospatialKmsKeyId, "kms-key-id", "", "", "KMS Key ID")
	_sagemakergeospatialCmd.Flags().StringVarP(&_sagemakergeospatialMaxResults, "max-results", "", "", "Max Results")
	_sagemakergeospatialCmd.Flags().StringVarP(&_sagemakergeospatialName, "name", "", "", "Name")
	_sagemakergeospatialCmd.Flags().StringVarP(&_sagemakergeospatialNextToken, "next-token", "", "", "Next Token")
	_sagemakergeospatialCmd.Flags().StringVarP(&_sagemakergeospatialOutputConfig, "output-config", "", "", "Output Config")
	_sagemakergeospatialCmd.Flags().StringVarP(&_sagemakergeospatialOutputDataType, "output-data-type", "", "", "Output Data Type")
	_sagemakergeospatialCmd.Flags().StringVarP(&_sagemakergeospatialOutputFormat, "output-format", "", "", "Output Format")
	_sagemakergeospatialCmd.Flags().StringVarP(&_sagemakergeospatialPropertyFilters, "property-filters", "", "", "Property Filters")
	_sagemakergeospatialCmd.Flags().StringVarP(&_sagemakergeospatialRasterDataCollectionQuery, "raster-data-collection-query", "", "", "Raster Data Collection Query")
	_sagemakergeospatialCmd.Flags().StringVarP(&_sagemakergeospatialResourceArn, "resource-arn", "", "", "Resource ARN")
	_sagemakergeospatialCmd.Flags().StringVarP(&_sagemakergeospatialSortBy, "sort-by", "", "", "Sort By")
	_sagemakergeospatialCmd.Flags().StringVarP(&_sagemakergeospatialSortOrder, "sort-order", "", "", "Sort Order")
	_sagemakergeospatialCmd.Flags().StringVarP(&_sagemakergeospatialStatusEquals, "status-equals", "", "", "Status Equals")
	_sagemakergeospatialCmd.Flags().StringSliceVarP(&_sagemakergeospatialTagKeys, "tag-keys", "", nil, "Tag Keys")
	_sagemakergeospatialCmd.Flags().StringVarP(&_sagemakergeospatialTags, "tags", "", "", "Tags")
	_sagemakergeospatialCmd.Flags().StringVarP(&_sagemakergeospatialTarget, "target", "", "", "Target")
	_sagemakergeospatialCmd.Flags().StringVarP(&_sagemakergeospatialTimeRangeFilter, "time-range-filter", "", "", "Time Range Filter")
	_sagemakergeospatialCmd.Flags().StringVarP(&_sagemakergeospatialX, "x", "", "", "X")
	_sagemakergeospatialCmd.Flags().StringVarP(&_sagemakergeospatialY, "y", "", "", "Y")
	_sagemakergeospatialCmd.Flags().StringVarP(&_sagemakergeospatialZ, "z", "", "", "Z")

	_sagemakergeospatialCmd.Flags().BoolVarP(&_sagemakergeospatialDeleteEarthObservationJob, "delete-earth-observation-job", "", false, "Delete Earth Observation Job")
	_sagemakergeospatialCmd.Flags().BoolVarP(&_sagemakergeospatialDeleteVectorEnrichmentJob, "delete-vector-enrichment-job", "", false, "Delete Vector Enrichment Job")
	_sagemakergeospatialCmd.Flags().BoolVarP(&_sagemakergeospatialExportEarthObservationJob, "export-earth-observation-job", "", false, "Export Earth Observation Job")
	_sagemakergeospatialCmd.Flags().BoolVarP(&_sagemakergeospatialExportVectorEnrichmentJob, "export-vector-enrichment-job", "", false, "Export Vector Enrichment Job")
	_sagemakergeospatialCmd.Flags().BoolVarP(&_sagemakergeospatialGetEarthObservationJob, "get-earth-observation-job", "", false, "Get Earth Observation Job")
	_sagemakergeospatialCmd.Flags().BoolVarP(&_sagemakergeospatialGetRasterDataCollection, "get-raster-data-collection", "", false, "Get Raster Data Collection")
	_sagemakergeospatialCmd.Flags().BoolVarP(&_sagemakergeospatialGetTile, "get-tile", "", false, "Get Tile")
	_sagemakergeospatialCmd.Flags().BoolVarP(&_sagemakergeospatialGetVectorEnrichmentJob, "get-vector-enrichment-job", "", false, "Get Vector Enrichment Job")
	_sagemakergeospatialCmd.Flags().BoolVarP(&_sagemakergeospatialListEarthObservationJobs, "list-earth-observation-jobs", "", false, "List Earth Observation Jobs")
	_sagemakergeospatialCmd.Flags().BoolVarP(&_sagemakergeospatialListRasterDataCollections, "list-raster-data-collections", "", false, "List Raster Data Collections")
	_sagemakergeospatialCmd.Flags().BoolVarP(&_sagemakergeospatialListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_sagemakergeospatialCmd.Flags().BoolVarP(&_sagemakergeospatialListVectorEnrichmentJobs, "list-vector-enrichment-jobs", "", false, "List Vector Enrichment Jobs")
	_sagemakergeospatialCmd.Flags().BoolVarP(&_sagemakergeospatialSearchRasterDataCollection, "search-raster-data-collection", "", false, "Search Raster Data Collection")
	_sagemakergeospatialCmd.Flags().BoolVarP(&_sagemakergeospatialStartEarthObservationJob, "start-earth-observation-job", "", false, "Start Earth Observation Job")
	_sagemakergeospatialCmd.Flags().BoolVarP(&_sagemakergeospatialStartVectorEnrichmentJob, "start-vector-enrichment-job", "", false, "Start Vector Enrichment Job")
	_sagemakergeospatialCmd.Flags().BoolVarP(&_sagemakergeospatialStopEarthObservationJob, "stop-earth-observation-job", "", false, "Stop Earth Observation Job")
	_sagemakergeospatialCmd.Flags().BoolVarP(&_sagemakergeospatialStopVectorEnrichmentJob, "stop-vector-enrichment-job", "", false, "Stop Vector Enrichment Job")
	_sagemakergeospatialCmd.Flags().BoolVarP(&_sagemakergeospatialTagResource, "tag-resource", "", false, "Tag Resource")
	_sagemakergeospatialCmd.Flags().BoolVarP(&_sagemakergeospatialUntagResource, "untag-resource", "", false, "Untag Resource")

}
