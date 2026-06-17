package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mediapackage"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// mediapackageCmd represents the mediapackage command
var _mediapackageCmd = &cobra.Command{
	Use:   "mediapackage",
	Short: "AWS mediapackage CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := mediapackage.NewFromConfig(cfg)
		if _mediapackageConfigureLogs {
			mediapackage_ConfigureLogs(cfg, client)
			return
		}
		if _mediapackageCreateChannel {
			mediapackage_CreateChannel(cfg, client)
			return
		}
		if _mediapackageCreateHarvestJob {
			mediapackage_CreateHarvestJob(cfg, client)
			return
		}
		if _mediapackageCreateOriginEndpoint {
			mediapackage_CreateOriginEndpoint(cfg, client)
			return
		}
		if _mediapackageDeleteChannel {
			mediapackage_DeleteChannel(cfg, client)
			return
		}
		if _mediapackageDeleteOriginEndpoint {
			mediapackage_DeleteOriginEndpoint(cfg, client)
			return
		}
		if _mediapackageDescribeChannel {
			mediapackage_DescribeChannel(cfg, client)
			return
		}
		if _mediapackageDescribeHarvestJob {
			mediapackage_DescribeHarvestJob(cfg, client)
			return
		}
		if _mediapackageDescribeOriginEndpoint {
			mediapackage_DescribeOriginEndpoint(cfg, client)
			return
		}
		if _mediapackageListChannels {
			mediapackage_ListChannels(cfg, client)
			return
		}
		if _mediapackageListHarvestJobs {
			mediapackage_ListHarvestJobs(cfg, client)
			return
		}
		if _mediapackageListOriginEndpoints {
			mediapackage_ListOriginEndpoints(cfg, client)
			return
		}
		if _mediapackageListTagsForResource {
			mediapackage_ListTagsForResource(cfg, client)
			return
		}
		if _mediapackageRotateChannelCredentials {
			mediapackage_RotateChannelCredentials(cfg, client)
			return
		}
		if _mediapackageRotateIngestEndpointCredentials {
			mediapackage_RotateIngestEndpointCredentials(cfg, client)
			return
		}
		if _mediapackageTagResource {
			mediapackage_TagResource(cfg, client)
			return
		}
		if _mediapackageUntagResource {
			mediapackage_UntagResource(cfg, client)
			return
		}
		if _mediapackageUpdateChannel {
			mediapackage_UpdateChannel(cfg, client)
			return
		}
		if _mediapackageUpdateOriginEndpoint {
			mediapackage_UpdateOriginEndpoint(cfg, client)
			return
		}

	},
}

var (
	_mediapackageConfigureLogs                   bool
	_mediapackageCreateChannel                   bool
	_mediapackageCreateHarvestJob                bool
	_mediapackageCreateOriginEndpoint            bool
	_mediapackageDeleteChannel                   bool
	_mediapackageDeleteOriginEndpoint            bool
	_mediapackageDescribeChannel                 bool
	_mediapackageDescribeHarvestJob              bool
	_mediapackageDescribeOriginEndpoint          bool
	_mediapackageListChannels                    bool
	_mediapackageListHarvestJobs                 bool
	_mediapackageListOriginEndpoints             bool
	_mediapackageListTagsForResource             bool
	_mediapackageRotateChannelCredentials        bool
	_mediapackageRotateIngestEndpointCredentials bool
	_mediapackageTagResource                     bool
	_mediapackageUntagResource                   bool
	_mediapackageUpdateChannel                   bool
	_mediapackageUpdateOriginEndpoint            bool

	_mediapackageAuthorization          string
	_mediapackageChannelId              string
	_mediapackageCmafPackage            string
	_mediapackageDashPackage            string
	_mediapackageDescription            string
	_mediapackageEgressAccessLogs       string
	_mediapackageEndTime                string
	_mediapackageHlsPackage             string
	_mediapackageId                     string
	_mediapackageIncludeChannelId       string
	_mediapackageIncludeStatus          string
	_mediapackageIngestEndpointId       string
	_mediapackageIngressAccessLogs      string
	_mediapackageManifestName           string
	_mediapackageMaxResults             string
	_mediapackageMssPackage             string
	_mediapackageNextToken              string
	_mediapackageOriginEndpointId       string
	_mediapackageOrigination            string
	_mediapackageResourceArn            string
	_mediapackageS3Destination          string
	_mediapackageStartTime              string
	_mediapackageStartoverWindowSeconds string
	_mediapackageTagKeys                []string
	_mediapackageTags                   string
	_mediapackageTimeDelaySeconds       string
	_mediapackageWhitelist              []string
)

// Changes the Channel's properities to configure log subscription
func mediapackage_ConfigureLogs(cfg aws.Config, client *mediapackage.Client) {
	input := &mediapackage.ConfigureLogsInput{
		// Id: *string, // Required
	}

	if len(_mediapackageId) > 0 {
		input.Id = aws.String(_mediapackageId)
	}
	if len(_mediapackageEgressAccessLogs) > 0 {
		if err := assignInputField(input, "EgressAccessLogs", _mediapackageEgressAccessLogs); err != nil {
			log.Errorf("invalid --egress-access-logs: %s", err.Error())
			return
		}
	}
	if len(_mediapackageIngressAccessLogs) > 0 {
		if err := assignInputField(input, "IngressAccessLogs", _mediapackageIngressAccessLogs); err != nil {
			log.Errorf("invalid --ingress-access-logs: %s", err.Error())
			return
		}
	}

	if resp, err := client.ConfigureLogs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new Channel.
func mediapackage_CreateChannel(cfg aws.Config, client *mediapackage.Client) {
	input := &mediapackage.CreateChannelInput{
		// Id: *string, // Required
	}

	if len(_mediapackageId) > 0 {
		input.Id = aws.String(_mediapackageId)
	}
	if len(_mediapackageDescription) > 0 {
		input.Description = aws.String(_mediapackageDescription)
	}
	if len(_mediapackageTags) > 0 {
		if err := assignInputField(input, "Tags", _mediapackageTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new HarvestJob record.
func mediapackage_CreateHarvestJob(cfg aws.Config, client *mediapackage.Client) {
	input := &mediapackage.CreateHarvestJobInput{
		// EndTime: *string, // Required
		// Id: *string, // Required
		// OriginEndpointId: *string, // Required
		// S3Destination: *types.S3Destination, // Required
		// StartTime: *string, // Required
	}

	if len(_mediapackageEndTime) > 0 {
		input.EndTime = aws.String(_mediapackageEndTime)
	}
	if len(_mediapackageId) > 0 {
		input.Id = aws.String(_mediapackageId)
	}
	if len(_mediapackageOriginEndpointId) > 0 {
		input.OriginEndpointId = aws.String(_mediapackageOriginEndpointId)
	}
	if len(_mediapackageS3Destination) > 0 {
		if err := assignInputField(input, "S3Destination", _mediapackageS3Destination); err != nil {
			log.Errorf("invalid --s3-destination: %s", err.Error())
			return
		}
	}
	if len(_mediapackageStartTime) > 0 {
		input.StartTime = aws.String(_mediapackageStartTime)
	}

	if resp, err := client.CreateHarvestJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new OriginEndpoint record.
func mediapackage_CreateOriginEndpoint(cfg aws.Config, client *mediapackage.Client) {
	input := &mediapackage.CreateOriginEndpointInput{
		// ChannelId: *string, // Required
		// Id: *string, // Required
	}

	if len(_mediapackageChannelId) > 0 {
		input.ChannelId = aws.String(_mediapackageChannelId)
	}
	if len(_mediapackageId) > 0 {
		input.Id = aws.String(_mediapackageId)
	}
	if len(_mediapackageAuthorization) > 0 {
		if err := assignInputField(input, "Authorization", _mediapackageAuthorization); err != nil {
			log.Errorf("invalid --authorization: %s", err.Error())
			return
		}
	}
	if len(_mediapackageCmafPackage) > 0 {
		if err := assignInputField(input, "CmafPackage", _mediapackageCmafPackage); err != nil {
			log.Errorf("invalid --cmaf-package: %s", err.Error())
			return
		}
	}
	if len(_mediapackageDashPackage) > 0 {
		if err := assignInputField(input, "DashPackage", _mediapackageDashPackage); err != nil {
			log.Errorf("invalid --dash-package: %s", err.Error())
			return
		}
	}
	if len(_mediapackageDescription) > 0 {
		input.Description = aws.String(_mediapackageDescription)
	}
	if len(_mediapackageHlsPackage) > 0 {
		if err := assignInputField(input, "HlsPackage", _mediapackageHlsPackage); err != nil {
			log.Errorf("invalid --hls-package: %s", err.Error())
			return
		}
	}
	if len(_mediapackageManifestName) > 0 {
		input.ManifestName = aws.String(_mediapackageManifestName)
	}
	if len(_mediapackageMssPackage) > 0 {
		if err := assignInputField(input, "MssPackage", _mediapackageMssPackage); err != nil {
			log.Errorf("invalid --mss-package: %s", err.Error())
			return
		}
	}
	if len(_mediapackageOrigination) > 0 {
		if err := assignInputField(input, "Origination", _mediapackageOrigination); err != nil {
			log.Errorf("invalid --origination: %s", err.Error())
			return
		}
	}
	if len(_mediapackageStartoverWindowSeconds) > 0 {
		if err := assignInputField(input, "StartoverWindowSeconds", _mediapackageStartoverWindowSeconds); err != nil {
			log.Errorf("invalid --startover-window-seconds: %s", err.Error())
			return
		}
	}
	if len(_mediapackageTags) > 0 {
		if err := assignInputField(input, "Tags", _mediapackageTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_mediapackageTimeDelaySeconds) > 0 {
		if err := assignInputField(input, "TimeDelaySeconds", _mediapackageTimeDelaySeconds); err != nil {
			log.Errorf("invalid --time-delay-seconds: %s", err.Error())
			return
		}
	}
	if len(_mediapackageWhitelist) > 0 {
		input.Whitelist = append([]string(nil), _mediapackageWhitelist...)
	}

	if resp, err := client.CreateOriginEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing Channel.
func mediapackage_DeleteChannel(cfg aws.Config, client *mediapackage.Client) {
	input := &mediapackage.DeleteChannelInput{
		// Id: *string, // Required
	}

	if len(_mediapackageId) > 0 {
		input.Id = aws.String(_mediapackageId)
	}

	if resp, err := client.DeleteChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing OriginEndpoint.
func mediapackage_DeleteOriginEndpoint(cfg aws.Config, client *mediapackage.Client) {
	input := &mediapackage.DeleteOriginEndpointInput{
		// Id: *string, // Required
	}

	if len(_mediapackageId) > 0 {
		input.Id = aws.String(_mediapackageId)
	}

	if resp, err := client.DeleteOriginEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets details about a Channel.
func mediapackage_DescribeChannel(cfg aws.Config, client *mediapackage.Client) {
	input := &mediapackage.DescribeChannelInput{
		// Id: *string, // Required
	}

	if len(_mediapackageId) > 0 {
		input.Id = aws.String(_mediapackageId)
	}

	if resp, err := client.DescribeChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets details about an existing HarvestJob.
func mediapackage_DescribeHarvestJob(cfg aws.Config, client *mediapackage.Client) {
	input := &mediapackage.DescribeHarvestJobInput{
		// Id: *string, // Required
	}

	if len(_mediapackageId) > 0 {
		input.Id = aws.String(_mediapackageId)
	}

	if resp, err := client.DescribeHarvestJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets details about an existing OriginEndpoint.
func mediapackage_DescribeOriginEndpoint(cfg aws.Config, client *mediapackage.Client) {
	input := &mediapackage.DescribeOriginEndpointInput{
		// Id: *string, // Required
	}

	if len(_mediapackageId) > 0 {
		input.Id = aws.String(_mediapackageId)
	}

	if resp, err := client.DescribeOriginEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a collection of Channels.
func mediapackage_ListChannels(cfg aws.Config, client *mediapackage.Client) {
	input := &mediapackage.ListChannelsInput{}

	if len(_mediapackageMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mediapackageMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mediapackageNextToken) > 0 {
		input.NextToken = aws.String(_mediapackageNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListChannels(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mediapackage.ListChannelsOutput
	p := mediapackage.NewListChannelsPaginator(client, input)
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

// Returns a collection of HarvestJob records.
func mediapackage_ListHarvestJobs(cfg aws.Config, client *mediapackage.Client) {
	input := &mediapackage.ListHarvestJobsInput{}

	if len(_mediapackageIncludeChannelId) > 0 {
		input.IncludeChannelId = aws.String(_mediapackageIncludeChannelId)
	}
	if len(_mediapackageIncludeStatus) > 0 {
		input.IncludeStatus = aws.String(_mediapackageIncludeStatus)
	}
	if len(_mediapackageMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mediapackageMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mediapackageNextToken) > 0 {
		input.NextToken = aws.String(_mediapackageNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListHarvestJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mediapackage.ListHarvestJobsOutput
	p := mediapackage.NewListHarvestJobsPaginator(client, input)
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

// Returns a collection of OriginEndpoint records.
func mediapackage_ListOriginEndpoints(cfg aws.Config, client *mediapackage.Client) {
	input := &mediapackage.ListOriginEndpointsInput{}

	if len(_mediapackageChannelId) > 0 {
		input.ChannelId = aws.String(_mediapackageChannelId)
	}
	if len(_mediapackageMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mediapackageMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mediapackageNextToken) > 0 {
		input.NextToken = aws.String(_mediapackageNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListOriginEndpoints(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mediapackage.ListOriginEndpointsOutput
	p := mediapackage.NewListOriginEndpointsPaginator(client, input)
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

func mediapackage_ListTagsForResource(cfg aws.Config, client *mediapackage.Client) {
	input := &mediapackage.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_mediapackageResourceArn) > 0 {
		input.ResourceArn = aws.String(_mediapackageResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes the Channel's first IngestEndpoint's username and password. WARNING -
// This API is deprecated. Please use RotateIngestEndpointCredentials instead
//
// Deprecated: This API is deprecated. Please use RotateIngestEndpointCredentials
// instead
func mediapackage_RotateChannelCredentials(cfg aws.Config, client *mediapackage.Client) {
	input := &mediapackage.RotateChannelCredentialsInput{
		// Id: *string, // Required
	}

	if len(_mediapackageId) > 0 {
		input.Id = aws.String(_mediapackageId)
	}

	if resp, err := client.RotateChannelCredentials(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Rotate the IngestEndpoint's username and password, as specified by the
// IngestEndpoint's id.
func mediapackage_RotateIngestEndpointCredentials(cfg aws.Config, client *mediapackage.Client) {
	input := &mediapackage.RotateIngestEndpointCredentialsInput{
		// Id: *string, // Required
		// IngestEndpointId: *string, // Required
	}

	if len(_mediapackageId) > 0 {
		input.Id = aws.String(_mediapackageId)
	}
	if len(_mediapackageIngestEndpointId) > 0 {
		input.IngestEndpointId = aws.String(_mediapackageIngestEndpointId)
	}

	if resp, err := client.RotateIngestEndpointCredentials(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func mediapackage_TagResource(cfg aws.Config, client *mediapackage.Client) {
	input := &mediapackage.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_mediapackageResourceArn) > 0 {
		input.ResourceArn = aws.String(_mediapackageResourceArn)
	}
	if len(_mediapackageTags) > 0 {
		if err := assignInputField(input, "Tags", _mediapackageTags); err != nil {
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

func mediapackage_UntagResource(cfg aws.Config, client *mediapackage.Client) {
	input := &mediapackage.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_mediapackageResourceArn) > 0 {
		input.ResourceArn = aws.String(_mediapackageResourceArn)
	}
	if len(_mediapackageTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _mediapackageTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing Channel.
func mediapackage_UpdateChannel(cfg aws.Config, client *mediapackage.Client) {
	input := &mediapackage.UpdateChannelInput{
		// Id: *string, // Required
	}

	if len(_mediapackageId) > 0 {
		input.Id = aws.String(_mediapackageId)
	}
	if len(_mediapackageDescription) > 0 {
		input.Description = aws.String(_mediapackageDescription)
	}

	if resp, err := client.UpdateChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing OriginEndpoint.
func mediapackage_UpdateOriginEndpoint(cfg aws.Config, client *mediapackage.Client) {
	input := &mediapackage.UpdateOriginEndpointInput{
		// Id: *string, // Required
	}

	if len(_mediapackageId) > 0 {
		input.Id = aws.String(_mediapackageId)
	}
	if len(_mediapackageAuthorization) > 0 {
		if err := assignInputField(input, "Authorization", _mediapackageAuthorization); err != nil {
			log.Errorf("invalid --authorization: %s", err.Error())
			return
		}
	}
	if len(_mediapackageCmafPackage) > 0 {
		if err := assignInputField(input, "CmafPackage", _mediapackageCmafPackage); err != nil {
			log.Errorf("invalid --cmaf-package: %s", err.Error())
			return
		}
	}
	if len(_mediapackageDashPackage) > 0 {
		if err := assignInputField(input, "DashPackage", _mediapackageDashPackage); err != nil {
			log.Errorf("invalid --dash-package: %s", err.Error())
			return
		}
	}
	if len(_mediapackageDescription) > 0 {
		input.Description = aws.String(_mediapackageDescription)
	}
	if len(_mediapackageHlsPackage) > 0 {
		if err := assignInputField(input, "HlsPackage", _mediapackageHlsPackage); err != nil {
			log.Errorf("invalid --hls-package: %s", err.Error())
			return
		}
	}
	if len(_mediapackageManifestName) > 0 {
		input.ManifestName = aws.String(_mediapackageManifestName)
	}
	if len(_mediapackageMssPackage) > 0 {
		if err := assignInputField(input, "MssPackage", _mediapackageMssPackage); err != nil {
			log.Errorf("invalid --mss-package: %s", err.Error())
			return
		}
	}
	if len(_mediapackageOrigination) > 0 {
		if err := assignInputField(input, "Origination", _mediapackageOrigination); err != nil {
			log.Errorf("invalid --origination: %s", err.Error())
			return
		}
	}
	if len(_mediapackageStartoverWindowSeconds) > 0 {
		if err := assignInputField(input, "StartoverWindowSeconds", _mediapackageStartoverWindowSeconds); err != nil {
			log.Errorf("invalid --startover-window-seconds: %s", err.Error())
			return
		}
	}
	if len(_mediapackageTimeDelaySeconds) > 0 {
		if err := assignInputField(input, "TimeDelaySeconds", _mediapackageTimeDelaySeconds); err != nil {
			log.Errorf("invalid --time-delay-seconds: %s", err.Error())
			return
		}
	}
	if len(_mediapackageWhitelist) > 0 {
		input.Whitelist = append([]string(nil), _mediapackageWhitelist...)
	}

	if resp, err := client.UpdateOriginEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_mediapackageCmd)
	_mediapackageCmd.Flags().SortFlags = false

	_mediapackageCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_mediapackageCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_mediapackageCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_mediapackageCmd.Flags().StringVarP(&_mediapackageAuthorization, "authorization", "", "", "Authorization")
	_mediapackageCmd.Flags().StringVarP(&_mediapackageChannelId, "channel-id", "", "", "Channel ID")
	_mediapackageCmd.Flags().StringVarP(&_mediapackageCmafPackage, "cmaf-package", "", "", "Cmaf Package")
	_mediapackageCmd.Flags().StringVarP(&_mediapackageDashPackage, "dash-package", "", "", "Dash Package")
	_mediapackageCmd.Flags().StringVarP(&_mediapackageDescription, "description", "", "", "Description")
	_mediapackageCmd.Flags().StringVarP(&_mediapackageEgressAccessLogs, "egress-access-logs", "", "", "Egress Access Logs")
	_mediapackageCmd.Flags().StringVarP(&_mediapackageEndTime, "end-time", "", "", "End Time")
	_mediapackageCmd.Flags().StringVarP(&_mediapackageHlsPackage, "hls-package", "", "", "Hls Package")
	_mediapackageCmd.Flags().StringVarP(&_mediapackageId, "id", "", "", "ID")
	_mediapackageCmd.Flags().StringVarP(&_mediapackageIncludeChannelId, "include-channel-id", "", "", "Include Channel ID")
	_mediapackageCmd.Flags().StringVarP(&_mediapackageIncludeStatus, "include-status", "", "", "Include Status")
	_mediapackageCmd.Flags().StringVarP(&_mediapackageIngestEndpointId, "ingest-endpoint-id", "", "", "Ingest Endpoint ID")
	_mediapackageCmd.Flags().StringVarP(&_mediapackageIngressAccessLogs, "ingress-access-logs", "", "", "Ingress Access Logs")
	_mediapackageCmd.Flags().StringVarP(&_mediapackageManifestName, "manifest-name", "", "", "Manifest Name")
	_mediapackageCmd.Flags().StringVarP(&_mediapackageMaxResults, "max-results", "", "", "Max Results")
	_mediapackageCmd.Flags().StringVarP(&_mediapackageMssPackage, "mss-package", "", "", "Mss Package")
	_mediapackageCmd.Flags().StringVarP(&_mediapackageNextToken, "next-token", "", "", "Next Token")
	_mediapackageCmd.Flags().StringVarP(&_mediapackageOriginEndpointId, "origin-endpoint-id", "", "", "Origin Endpoint ID")
	_mediapackageCmd.Flags().StringVarP(&_mediapackageOrigination, "origination", "", "", "Origination")
	_mediapackageCmd.Flags().StringVarP(&_mediapackageResourceArn, "resource-arn", "", "", "Resource ARN")
	_mediapackageCmd.Flags().StringVarP(&_mediapackageS3Destination, "s3-destination", "", "", "S3 Destination")
	_mediapackageCmd.Flags().StringVarP(&_mediapackageStartTime, "start-time", "", "", "Start Time")
	_mediapackageCmd.Flags().StringVarP(&_mediapackageStartoverWindowSeconds, "startover-window-seconds", "", "", "Startover Window Seconds")
	_mediapackageCmd.Flags().StringSliceVarP(&_mediapackageTagKeys, "tag-keys", "", nil, "Tag Keys")
	_mediapackageCmd.Flags().StringVarP(&_mediapackageTags, "tags", "", "", "Tags")
	_mediapackageCmd.Flags().StringVarP(&_mediapackageTimeDelaySeconds, "time-delay-seconds", "", "", "Time Delay Seconds")
	_mediapackageCmd.Flags().StringSliceVarP(&_mediapackageWhitelist, "whitelist", "", nil, "Whitelist")

	_mediapackageCmd.Flags().BoolVarP(&_mediapackageConfigureLogs, "configure-logs", "", false, "Configure Logs")
	_mediapackageCmd.Flags().BoolVarP(&_mediapackageCreateChannel, "create-channel", "", false, "Create Channel")
	_mediapackageCmd.Flags().BoolVarP(&_mediapackageCreateHarvestJob, "create-harvest-job", "", false, "Create Harvest Job")
	_mediapackageCmd.Flags().BoolVarP(&_mediapackageCreateOriginEndpoint, "create-origin-endpoint", "", false, "Create Origin Endpoint")
	_mediapackageCmd.Flags().BoolVarP(&_mediapackageDeleteChannel, "delete-channel", "", false, "Delete Channel")
	_mediapackageCmd.Flags().BoolVarP(&_mediapackageDeleteOriginEndpoint, "delete-origin-endpoint", "", false, "Delete Origin Endpoint")
	_mediapackageCmd.Flags().BoolVarP(&_mediapackageDescribeChannel, "describe-channel", "", false, "Describe Channel")
	_mediapackageCmd.Flags().BoolVarP(&_mediapackageDescribeHarvestJob, "describe-harvest-job", "", false, "Describe Harvest Job")
	_mediapackageCmd.Flags().BoolVarP(&_mediapackageDescribeOriginEndpoint, "describe-origin-endpoint", "", false, "Describe Origin Endpoint")
	_mediapackageCmd.Flags().BoolVarP(&_mediapackageListChannels, "list-channels", "", false, "List Channels")
	_mediapackageCmd.Flags().BoolVarP(&_mediapackageListHarvestJobs, "list-harvest-jobs", "", false, "List Harvest Jobs")
	_mediapackageCmd.Flags().BoolVarP(&_mediapackageListOriginEndpoints, "list-origin-endpoints", "", false, "List Origin Endpoints")
	_mediapackageCmd.Flags().BoolVarP(&_mediapackageListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_mediapackageCmd.Flags().BoolVarP(&_mediapackageRotateChannelCredentials, "rotate-channel-credentials", "", false, "Rotate Channel Credentials")
	_mediapackageCmd.Flags().BoolVarP(&_mediapackageRotateIngestEndpointCredentials, "rotate-ingest-endpoint-credentials", "", false, "Rotate Ingest Endpoint Credentials")
	_mediapackageCmd.Flags().BoolVarP(&_mediapackageTagResource, "tag-resource", "", false, "Tag Resource")
	_mediapackageCmd.Flags().BoolVarP(&_mediapackageUntagResource, "untag-resource", "", false, "Untag Resource")
	_mediapackageCmd.Flags().BoolVarP(&_mediapackageUpdateChannel, "update-channel", "", false, "Update Channel")
	_mediapackageCmd.Flags().BoolVarP(&_mediapackageUpdateOriginEndpoint, "update-origin-endpoint", "", false, "Update Origin Endpoint")

}
