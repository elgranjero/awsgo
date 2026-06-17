package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mediapackagevod"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// mediapackagevodCmd represents the mediapackagevod command
var _mediapackagevodCmd = &cobra.Command{
	Use:   "mediapackagevod",
	Short: "AWS mediapackagevod CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := mediapackagevod.NewFromConfig(cfg)
		if _mediapackagevodConfigureLogs {
			mediapackagevod_ConfigureLogs(cfg, client)
			return
		}
		if _mediapackagevodCreateAsset {
			mediapackagevod_CreateAsset(cfg, client)
			return
		}
		if _mediapackagevodCreatePackagingConfiguration {
			mediapackagevod_CreatePackagingConfiguration(cfg, client)
			return
		}
		if _mediapackagevodCreatePackagingGroup {
			mediapackagevod_CreatePackagingGroup(cfg, client)
			return
		}
		if _mediapackagevodDeleteAsset {
			mediapackagevod_DeleteAsset(cfg, client)
			return
		}
		if _mediapackagevodDeletePackagingConfiguration {
			mediapackagevod_DeletePackagingConfiguration(cfg, client)
			return
		}
		if _mediapackagevodDeletePackagingGroup {
			mediapackagevod_DeletePackagingGroup(cfg, client)
			return
		}
		if _mediapackagevodDescribeAsset {
			mediapackagevod_DescribeAsset(cfg, client)
			return
		}
		if _mediapackagevodDescribePackagingConfiguration {
			mediapackagevod_DescribePackagingConfiguration(cfg, client)
			return
		}
		if _mediapackagevodDescribePackagingGroup {
			mediapackagevod_DescribePackagingGroup(cfg, client)
			return
		}
		if _mediapackagevodListAssets {
			mediapackagevod_ListAssets(cfg, client)
			return
		}
		if _mediapackagevodListPackagingConfigurations {
			mediapackagevod_ListPackagingConfigurations(cfg, client)
			return
		}
		if _mediapackagevodListPackagingGroups {
			mediapackagevod_ListPackagingGroups(cfg, client)
			return
		}
		if _mediapackagevodListTagsForResource {
			mediapackagevod_ListTagsForResource(cfg, client)
			return
		}
		if _mediapackagevodTagResource {
			mediapackagevod_TagResource(cfg, client)
			return
		}
		if _mediapackagevodUntagResource {
			mediapackagevod_UntagResource(cfg, client)
			return
		}
		if _mediapackagevodUpdatePackagingGroup {
			mediapackagevod_UpdatePackagingGroup(cfg, client)
			return
		}

	},
}

var (
	_mediapackagevodConfigureLogs                  bool
	_mediapackagevodCreateAsset                    bool
	_mediapackagevodCreatePackagingConfiguration   bool
	_mediapackagevodCreatePackagingGroup           bool
	_mediapackagevodDeleteAsset                    bool
	_mediapackagevodDeletePackagingConfiguration   bool
	_mediapackagevodDeletePackagingGroup           bool
	_mediapackagevodDescribeAsset                  bool
	_mediapackagevodDescribePackagingConfiguration bool
	_mediapackagevodDescribePackagingGroup         bool
	_mediapackagevodListAssets                     bool
	_mediapackagevodListPackagingConfigurations    bool
	_mediapackagevodListPackagingGroups            bool
	_mediapackagevodListTagsForResource            bool
	_mediapackagevodTagResource                    bool
	_mediapackagevodUntagResource                  bool
	_mediapackagevodUpdatePackagingGroup           bool

	_mediapackagevodAuthorization    string
	_mediapackagevodCmafPackage      string
	_mediapackagevodDashPackage      string
	_mediapackagevodEgressAccessLogs string
	_mediapackagevodHlsPackage       string
	_mediapackagevodId               string
	_mediapackagevodMaxResults       string
	_mediapackagevodMssPackage       string
	_mediapackagevodNextToken        string
	_mediapackagevodPackagingGroupId string
	_mediapackagevodResourceArn      string
	_mediapackagevodResourceId       string
	_mediapackagevodSourceArn        string
	_mediapackagevodSourceRoleArn    string
	_mediapackagevodTagKeys          []string
	_mediapackagevodTags             string
)

// Changes the packaging group's properities to configure log subscription
func mediapackagevod_ConfigureLogs(cfg aws.Config, client *mediapackagevod.Client) {
	input := &mediapackagevod.ConfigureLogsInput{
		// Id: *string, // Required
	}

	if len(_mediapackagevodId) > 0 {
		input.Id = aws.String(_mediapackagevodId)
	}
	if len(_mediapackagevodEgressAccessLogs) > 0 {
		if err := assignInputField(input, "EgressAccessLogs", _mediapackagevodEgressAccessLogs); err != nil {
			log.Errorf("invalid --egress-access-logs: %s", err.Error())
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

// Creates a new MediaPackage VOD Asset resource.
func mediapackagevod_CreateAsset(cfg aws.Config, client *mediapackagevod.Client) {
	input := &mediapackagevod.CreateAssetInput{
		// Id: *string, // Required
		// PackagingGroupId: *string, // Required
		// SourceArn: *string, // Required
		// SourceRoleArn: *string, // Required
	}

	if len(_mediapackagevodId) > 0 {
		input.Id = aws.String(_mediapackagevodId)
	}
	if len(_mediapackagevodPackagingGroupId) > 0 {
		input.PackagingGroupId = aws.String(_mediapackagevodPackagingGroupId)
	}
	if len(_mediapackagevodSourceArn) > 0 {
		input.SourceArn = aws.String(_mediapackagevodSourceArn)
	}
	if len(_mediapackagevodSourceRoleArn) > 0 {
		input.SourceRoleArn = aws.String(_mediapackagevodSourceRoleArn)
	}
	if len(_mediapackagevodResourceId) > 0 {
		input.ResourceId = aws.String(_mediapackagevodResourceId)
	}
	if len(_mediapackagevodTags) > 0 {
		if err := assignInputField(input, "Tags", _mediapackagevodTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAsset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new MediaPackage VOD PackagingConfiguration resource.
func mediapackagevod_CreatePackagingConfiguration(cfg aws.Config, client *mediapackagevod.Client) {
	input := &mediapackagevod.CreatePackagingConfigurationInput{
		// Id: *string, // Required
		// PackagingGroupId: *string, // Required
	}

	if len(_mediapackagevodId) > 0 {
		input.Id = aws.String(_mediapackagevodId)
	}
	if len(_mediapackagevodPackagingGroupId) > 0 {
		input.PackagingGroupId = aws.String(_mediapackagevodPackagingGroupId)
	}
	if len(_mediapackagevodCmafPackage) > 0 {
		if err := assignInputField(input, "CmafPackage", _mediapackagevodCmafPackage); err != nil {
			log.Errorf("invalid --cmaf-package: %s", err.Error())
			return
		}
	}
	if len(_mediapackagevodDashPackage) > 0 {
		if err := assignInputField(input, "DashPackage", _mediapackagevodDashPackage); err != nil {
			log.Errorf("invalid --dash-package: %s", err.Error())
			return
		}
	}
	if len(_mediapackagevodHlsPackage) > 0 {
		if err := assignInputField(input, "HlsPackage", _mediapackagevodHlsPackage); err != nil {
			log.Errorf("invalid --hls-package: %s", err.Error())
			return
		}
	}
	if len(_mediapackagevodMssPackage) > 0 {
		if err := assignInputField(input, "MssPackage", _mediapackagevodMssPackage); err != nil {
			log.Errorf("invalid --mss-package: %s", err.Error())
			return
		}
	}
	if len(_mediapackagevodTags) > 0 {
		if err := assignInputField(input, "Tags", _mediapackagevodTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePackagingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new MediaPackage VOD PackagingGroup resource.
func mediapackagevod_CreatePackagingGroup(cfg aws.Config, client *mediapackagevod.Client) {
	input := &mediapackagevod.CreatePackagingGroupInput{
		// Id: *string, // Required
	}

	if len(_mediapackagevodId) > 0 {
		input.Id = aws.String(_mediapackagevodId)
	}
	if len(_mediapackagevodAuthorization) > 0 {
		if err := assignInputField(input, "Authorization", _mediapackagevodAuthorization); err != nil {
			log.Errorf("invalid --authorization: %s", err.Error())
			return
		}
	}
	if len(_mediapackagevodEgressAccessLogs) > 0 {
		if err := assignInputField(input, "EgressAccessLogs", _mediapackagevodEgressAccessLogs); err != nil {
			log.Errorf("invalid --egress-access-logs: %s", err.Error())
			return
		}
	}
	if len(_mediapackagevodTags) > 0 {
		if err := assignInputField(input, "Tags", _mediapackagevodTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePackagingGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing MediaPackage VOD Asset resource.
func mediapackagevod_DeleteAsset(cfg aws.Config, client *mediapackagevod.Client) {
	input := &mediapackagevod.DeleteAssetInput{
		// Id: *string, // Required
	}

	if len(_mediapackagevodId) > 0 {
		input.Id = aws.String(_mediapackagevodId)
	}

	if resp, err := client.DeleteAsset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a MediaPackage VOD PackagingConfiguration resource.
func mediapackagevod_DeletePackagingConfiguration(cfg aws.Config, client *mediapackagevod.Client) {
	input := &mediapackagevod.DeletePackagingConfigurationInput{
		// Id: *string, // Required
	}

	if len(_mediapackagevodId) > 0 {
		input.Id = aws.String(_mediapackagevodId)
	}

	if resp, err := client.DeletePackagingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a MediaPackage VOD PackagingGroup resource.
func mediapackagevod_DeletePackagingGroup(cfg aws.Config, client *mediapackagevod.Client) {
	input := &mediapackagevod.DeletePackagingGroupInput{
		// Id: *string, // Required
	}

	if len(_mediapackagevodId) > 0 {
		input.Id = aws.String(_mediapackagevodId)
	}

	if resp, err := client.DeletePackagingGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a description of a MediaPackage VOD Asset resource.
func mediapackagevod_DescribeAsset(cfg aws.Config, client *mediapackagevod.Client) {
	input := &mediapackagevod.DescribeAssetInput{
		// Id: *string, // Required
	}

	if len(_mediapackagevodId) > 0 {
		input.Id = aws.String(_mediapackagevodId)
	}

	if resp, err := client.DescribeAsset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a description of a MediaPackage VOD PackagingConfiguration resource.
func mediapackagevod_DescribePackagingConfiguration(cfg aws.Config, client *mediapackagevod.Client) {
	input := &mediapackagevod.DescribePackagingConfigurationInput{
		// Id: *string, // Required
	}

	if len(_mediapackagevodId) > 0 {
		input.Id = aws.String(_mediapackagevodId)
	}

	if resp, err := client.DescribePackagingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a description of a MediaPackage VOD PackagingGroup resource.
func mediapackagevod_DescribePackagingGroup(cfg aws.Config, client *mediapackagevod.Client) {
	input := &mediapackagevod.DescribePackagingGroupInput{
		// Id: *string, // Required
	}

	if len(_mediapackagevodId) > 0 {
		input.Id = aws.String(_mediapackagevodId)
	}

	if resp, err := client.DescribePackagingGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a collection of MediaPackage VOD Asset resources.
func mediapackagevod_ListAssets(cfg aws.Config, client *mediapackagevod.Client) {
	input := &mediapackagevod.ListAssetsInput{}

	if len(_mediapackagevodMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mediapackagevodMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mediapackagevodNextToken) > 0 {
		input.NextToken = aws.String(_mediapackagevodNextToken)
	}
	if len(_mediapackagevodPackagingGroupId) > 0 {
		input.PackagingGroupId = aws.String(_mediapackagevodPackagingGroupId)
	}

	if disablePaginator() {
		if resp, err := client.ListAssets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mediapackagevod.ListAssetsOutput
	p := mediapackagevod.NewListAssetsPaginator(client, input)
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

// Returns a collection of MediaPackage VOD PackagingConfiguration resources.
func mediapackagevod_ListPackagingConfigurations(cfg aws.Config, client *mediapackagevod.Client) {
	input := &mediapackagevod.ListPackagingConfigurationsInput{}

	if len(_mediapackagevodMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mediapackagevodMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mediapackagevodNextToken) > 0 {
		input.NextToken = aws.String(_mediapackagevodNextToken)
	}
	if len(_mediapackagevodPackagingGroupId) > 0 {
		input.PackagingGroupId = aws.String(_mediapackagevodPackagingGroupId)
	}

	if disablePaginator() {
		if resp, err := client.ListPackagingConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mediapackagevod.ListPackagingConfigurationsOutput
	p := mediapackagevod.NewListPackagingConfigurationsPaginator(client, input)
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

// Returns a collection of MediaPackage VOD PackagingGroup resources.
func mediapackagevod_ListPackagingGroups(cfg aws.Config, client *mediapackagevod.Client) {
	input := &mediapackagevod.ListPackagingGroupsInput{}

	if len(_mediapackagevodMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mediapackagevodMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mediapackagevodNextToken) > 0 {
		input.NextToken = aws.String(_mediapackagevodNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPackagingGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mediapackagevod.ListPackagingGroupsOutput
	p := mediapackagevod.NewListPackagingGroupsPaginator(client, input)
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

// Returns a list of the tags assigned to the specified resource.
func mediapackagevod_ListTagsForResource(cfg aws.Config, client *mediapackagevod.Client) {
	input := &mediapackagevod.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_mediapackagevodResourceArn) > 0 {
		input.ResourceArn = aws.String(_mediapackagevodResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds tags to the specified resource. You can specify one or more tags to add.
func mediapackagevod_TagResource(cfg aws.Config, client *mediapackagevod.Client) {
	input := &mediapackagevod.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_mediapackagevodResourceArn) > 0 {
		input.ResourceArn = aws.String(_mediapackagevodResourceArn)
	}
	if len(_mediapackagevodTags) > 0 {
		if err := assignInputField(input, "Tags", _mediapackagevodTags); err != nil {
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

// Removes tags from the specified resource. You can specify one or more tags to
// remove.
func mediapackagevod_UntagResource(cfg aws.Config, client *mediapackagevod.Client) {
	input := &mediapackagevod.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_mediapackagevodResourceArn) > 0 {
		input.ResourceArn = aws.String(_mediapackagevodResourceArn)
	}
	if len(_mediapackagevodTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _mediapackagevodTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a specific packaging group. You can't change the id attribute or any
// other system-generated attributes.
func mediapackagevod_UpdatePackagingGroup(cfg aws.Config, client *mediapackagevod.Client) {
	input := &mediapackagevod.UpdatePackagingGroupInput{
		// Id: *string, // Required
	}

	if len(_mediapackagevodId) > 0 {
		input.Id = aws.String(_mediapackagevodId)
	}
	if len(_mediapackagevodAuthorization) > 0 {
		if err := assignInputField(input, "Authorization", _mediapackagevodAuthorization); err != nil {
			log.Errorf("invalid --authorization: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdatePackagingGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_mediapackagevodCmd)
	_mediapackagevodCmd.Flags().SortFlags = false

	_mediapackagevodCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_mediapackagevodCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_mediapackagevodCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_mediapackagevodCmd.Flags().StringVarP(&_mediapackagevodAuthorization, "authorization", "", "", "Authorization")
	_mediapackagevodCmd.Flags().StringVarP(&_mediapackagevodCmafPackage, "cmaf-package", "", "", "Cmaf Package")
	_mediapackagevodCmd.Flags().StringVarP(&_mediapackagevodDashPackage, "dash-package", "", "", "Dash Package")
	_mediapackagevodCmd.Flags().StringVarP(&_mediapackagevodEgressAccessLogs, "egress-access-logs", "", "", "Egress Access Logs")
	_mediapackagevodCmd.Flags().StringVarP(&_mediapackagevodHlsPackage, "hls-package", "", "", "Hls Package")
	_mediapackagevodCmd.Flags().StringVarP(&_mediapackagevodId, "id", "", "", "ID")
	_mediapackagevodCmd.Flags().StringVarP(&_mediapackagevodMaxResults, "max-results", "", "", "Max Results")
	_mediapackagevodCmd.Flags().StringVarP(&_mediapackagevodMssPackage, "mss-package", "", "", "Mss Package")
	_mediapackagevodCmd.Flags().StringVarP(&_mediapackagevodNextToken, "next-token", "", "", "Next Token")
	_mediapackagevodCmd.Flags().StringVarP(&_mediapackagevodPackagingGroupId, "packaging-group-id", "", "", "Packaging Group ID")
	_mediapackagevodCmd.Flags().StringVarP(&_mediapackagevodResourceArn, "resource-arn", "", "", "Resource ARN")
	_mediapackagevodCmd.Flags().StringVarP(&_mediapackagevodResourceId, "resource-id", "", "", "Resource ID")
	_mediapackagevodCmd.Flags().StringVarP(&_mediapackagevodSourceArn, "source-arn", "", "", "Source ARN")
	_mediapackagevodCmd.Flags().StringVarP(&_mediapackagevodSourceRoleArn, "source-role-arn", "", "", "Source Role ARN")
	_mediapackagevodCmd.Flags().StringSliceVarP(&_mediapackagevodTagKeys, "tag-keys", "", nil, "Tag Keys")
	_mediapackagevodCmd.Flags().StringVarP(&_mediapackagevodTags, "tags", "", "", "Tags")

	_mediapackagevodCmd.Flags().BoolVarP(&_mediapackagevodConfigureLogs, "configure-logs", "", false, "Configure Logs")
	_mediapackagevodCmd.Flags().BoolVarP(&_mediapackagevodCreateAsset, "create-asset", "", false, "Create Asset")
	_mediapackagevodCmd.Flags().BoolVarP(&_mediapackagevodCreatePackagingConfiguration, "create-packaging-configuration", "", false, "Create Packaging Configuration")
	_mediapackagevodCmd.Flags().BoolVarP(&_mediapackagevodCreatePackagingGroup, "create-packaging-group", "", false, "Create Packaging Group")
	_mediapackagevodCmd.Flags().BoolVarP(&_mediapackagevodDeleteAsset, "delete-asset", "", false, "Delete Asset")
	_mediapackagevodCmd.Flags().BoolVarP(&_mediapackagevodDeletePackagingConfiguration, "delete-packaging-configuration", "", false, "Delete Packaging Configuration")
	_mediapackagevodCmd.Flags().BoolVarP(&_mediapackagevodDeletePackagingGroup, "delete-packaging-group", "", false, "Delete Packaging Group")
	_mediapackagevodCmd.Flags().BoolVarP(&_mediapackagevodDescribeAsset, "describe-asset", "", false, "Describe Asset")
	_mediapackagevodCmd.Flags().BoolVarP(&_mediapackagevodDescribePackagingConfiguration, "describe-packaging-configuration", "", false, "Describe Packaging Configuration")
	_mediapackagevodCmd.Flags().BoolVarP(&_mediapackagevodDescribePackagingGroup, "describe-packaging-group", "", false, "Describe Packaging Group")
	_mediapackagevodCmd.Flags().BoolVarP(&_mediapackagevodListAssets, "list-assets", "", false, "List Assets")
	_mediapackagevodCmd.Flags().BoolVarP(&_mediapackagevodListPackagingConfigurations, "list-packaging-configurations", "", false, "List Packaging Configurations")
	_mediapackagevodCmd.Flags().BoolVarP(&_mediapackagevodListPackagingGroups, "list-packaging-groups", "", false, "List Packaging Groups")
	_mediapackagevodCmd.Flags().BoolVarP(&_mediapackagevodListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_mediapackagevodCmd.Flags().BoolVarP(&_mediapackagevodTagResource, "tag-resource", "", false, "Tag Resource")
	_mediapackagevodCmd.Flags().BoolVarP(&_mediapackagevodUntagResource, "untag-resource", "", false, "Untag Resource")
	_mediapackagevodCmd.Flags().BoolVarP(&_mediapackagevodUpdatePackagingGroup, "update-packaging-group", "", false, "Update Packaging Group")

}
