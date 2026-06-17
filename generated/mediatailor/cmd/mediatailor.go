package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mediatailor"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// mediatailorCmd represents the mediatailor command
var _mediatailorCmd = &cobra.Command{
	Use:   "mediatailor",
	Short: "AWS mediatailor CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := mediatailor.NewFromConfig(cfg)
		if _mediatailorConfigureLogsForChannel {
			mediatailor_ConfigureLogsForChannel(cfg, client)
			return
		}
		if _mediatailorConfigureLogsForPlaybackConfiguration {
			mediatailor_ConfigureLogsForPlaybackConfiguration(cfg, client)
			return
		}
		if _mediatailorCreateChannel {
			mediatailor_CreateChannel(cfg, client)
			return
		}
		if _mediatailorCreateLiveSource {
			mediatailor_CreateLiveSource(cfg, client)
			return
		}
		if _mediatailorCreatePrefetchSchedule {
			mediatailor_CreatePrefetchSchedule(cfg, client)
			return
		}
		if _mediatailorCreateProgram {
			mediatailor_CreateProgram(cfg, client)
			return
		}
		if _mediatailorCreateSourceLocation {
			mediatailor_CreateSourceLocation(cfg, client)
			return
		}
		if _mediatailorCreateVodSource {
			mediatailor_CreateVodSource(cfg, client)
			return
		}
		if _mediatailorDeleteChannel {
			mediatailor_DeleteChannel(cfg, client)
			return
		}
		if _mediatailorDeleteChannelPolicy {
			mediatailor_DeleteChannelPolicy(cfg, client)
			return
		}
		if _mediatailorDeleteLiveSource {
			mediatailor_DeleteLiveSource(cfg, client)
			return
		}
		if _mediatailorDeletePlaybackConfiguration {
			mediatailor_DeletePlaybackConfiguration(cfg, client)
			return
		}
		if _mediatailorDeletePrefetchSchedule {
			mediatailor_DeletePrefetchSchedule(cfg, client)
			return
		}
		if _mediatailorDeleteProgram {
			mediatailor_DeleteProgram(cfg, client)
			return
		}
		if _mediatailorDeleteSourceLocation {
			mediatailor_DeleteSourceLocation(cfg, client)
			return
		}
		if _mediatailorDeleteVodSource {
			mediatailor_DeleteVodSource(cfg, client)
			return
		}
		if _mediatailorDescribeChannel {
			mediatailor_DescribeChannel(cfg, client)
			return
		}
		if _mediatailorDescribeLiveSource {
			mediatailor_DescribeLiveSource(cfg, client)
			return
		}
		if _mediatailorDescribeProgram {
			mediatailor_DescribeProgram(cfg, client)
			return
		}
		if _mediatailorDescribeSourceLocation {
			mediatailor_DescribeSourceLocation(cfg, client)
			return
		}
		if _mediatailorDescribeVodSource {
			mediatailor_DescribeVodSource(cfg, client)
			return
		}
		if _mediatailorGetChannelPolicy {
			mediatailor_GetChannelPolicy(cfg, client)
			return
		}
		if _mediatailorGetChannelSchedule {
			mediatailor_GetChannelSchedule(cfg, client)
			return
		}
		if _mediatailorGetPlaybackConfiguration {
			mediatailor_GetPlaybackConfiguration(cfg, client)
			return
		}
		if _mediatailorGetPrefetchSchedule {
			mediatailor_GetPrefetchSchedule(cfg, client)
			return
		}
		if _mediatailorListAlerts {
			mediatailor_ListAlerts(cfg, client)
			return
		}
		if _mediatailorListChannels {
			mediatailor_ListChannels(cfg, client)
			return
		}
		if _mediatailorListLiveSources {
			mediatailor_ListLiveSources(cfg, client)
			return
		}
		if _mediatailorListPlaybackConfigurations {
			mediatailor_ListPlaybackConfigurations(cfg, client)
			return
		}
		if _mediatailorListPrefetchSchedules {
			mediatailor_ListPrefetchSchedules(cfg, client)
			return
		}
		if _mediatailorListSourceLocations {
			mediatailor_ListSourceLocations(cfg, client)
			return
		}
		if _mediatailorListTagsForResource {
			mediatailor_ListTagsForResource(cfg, client)
			return
		}
		if _mediatailorListVodSources {
			mediatailor_ListVodSources(cfg, client)
			return
		}
		if _mediatailorPutChannelPolicy {
			mediatailor_PutChannelPolicy(cfg, client)
			return
		}
		if _mediatailorPutPlaybackConfiguration {
			mediatailor_PutPlaybackConfiguration(cfg, client)
			return
		}
		if _mediatailorStartChannel {
			mediatailor_StartChannel(cfg, client)
			return
		}
		if _mediatailorStopChannel {
			mediatailor_StopChannel(cfg, client)
			return
		}
		if _mediatailorTagResource {
			mediatailor_TagResource(cfg, client)
			return
		}
		if _mediatailorUntagResource {
			mediatailor_UntagResource(cfg, client)
			return
		}
		if _mediatailorUpdateChannel {
			mediatailor_UpdateChannel(cfg, client)
			return
		}
		if _mediatailorUpdateLiveSource {
			mediatailor_UpdateLiveSource(cfg, client)
			return
		}
		if _mediatailorUpdateProgram {
			mediatailor_UpdateProgram(cfg, client)
			return
		}
		if _mediatailorUpdateSourceLocation {
			mediatailor_UpdateSourceLocation(cfg, client)
			return
		}
		if _mediatailorUpdateVodSource {
			mediatailor_UpdateVodSource(cfg, client)
			return
		}

	},
}

var (
	_mediatailorConfigureLogsForChannel               bool
	_mediatailorConfigureLogsForPlaybackConfiguration bool
	_mediatailorCreateChannel                         bool
	_mediatailorCreateLiveSource                      bool
	_mediatailorCreatePrefetchSchedule                bool
	_mediatailorCreateProgram                         bool
	_mediatailorCreateSourceLocation                  bool
	_mediatailorCreateVodSource                       bool
	_mediatailorDeleteChannel                         bool
	_mediatailorDeleteChannelPolicy                   bool
	_mediatailorDeleteLiveSource                      bool
	_mediatailorDeletePlaybackConfiguration           bool
	_mediatailorDeletePrefetchSchedule                bool
	_mediatailorDeleteProgram                         bool
	_mediatailorDeleteSourceLocation                  bool
	_mediatailorDeleteVodSource                       bool
	_mediatailorDescribeChannel                       bool
	_mediatailorDescribeLiveSource                    bool
	_mediatailorDescribeProgram                       bool
	_mediatailorDescribeSourceLocation                bool
	_mediatailorDescribeVodSource                     bool
	_mediatailorGetChannelPolicy                      bool
	_mediatailorGetChannelSchedule                    bool
	_mediatailorGetPlaybackConfiguration              bool
	_mediatailorGetPrefetchSchedule                   bool
	_mediatailorListAlerts                            bool
	_mediatailorListChannels                          bool
	_mediatailorListLiveSources                       bool
	_mediatailorListPlaybackConfigurations            bool
	_mediatailorListPrefetchSchedules                 bool
	_mediatailorListSourceLocations                   bool
	_mediatailorListTagsForResource                   bool
	_mediatailorListVodSources                        bool
	_mediatailorPutChannelPolicy                      bool
	_mediatailorPutPlaybackConfiguration              bool
	_mediatailorStartChannel                          bool
	_mediatailorStopChannel                           bool
	_mediatailorTagResource                           bool
	_mediatailorUntagResource                         bool
	_mediatailorUpdateChannel                         bool
	_mediatailorUpdateLiveSource                      bool
	_mediatailorUpdateProgram                         bool
	_mediatailorUpdateSourceLocation                  bool
	_mediatailorUpdateVodSource                       bool

	_mediatailorAccessConfiguration                 string
	_mediatailorAdBreaks                            string
	_mediatailorAdConditioningConfiguration         string
	_mediatailorAdDecisionServerConfiguration       string
	_mediatailorAdDecisionServerUrl                 string
	_mediatailorAdsInteractionLog                   string
	_mediatailorAudience                            string
	_mediatailorAudienceMedia                       string
	_mediatailorAudiences                           []string
	_mediatailorAvailSuppression                    string
	_mediatailorBumper                              string
	_mediatailorCdnConfiguration                    string
	_mediatailorChannelName                         string
	_mediatailorConfigurationAliases                string
	_mediatailorConsumption                         string
	_mediatailorDashConfiguration                   string
	_mediatailorDefaultSegmentDeliveryConfiguration string
	_mediatailorDurationMinutes                     string
	_mediatailorEnabledLoggingStrategies            string
	_mediatailorFillerSlate                         string
	_mediatailorHttpConfiguration                   string
	_mediatailorHttpPackageConfigurations           string
	_mediatailorInsertionMode                       string
	_mediatailorLivePreRollConfiguration            string
	_mediatailorLiveSourceName                      string
	_mediatailorLogTypes                            string
	_mediatailorManifestProcessingRules             string
	_mediatailorManifestServiceInteractionLog       string
	_mediatailorMaxResults                          string
	_mediatailorName                                string
	_mediatailorNextToken                           string
	_mediatailorOutputs                             string
	_mediatailorPercentEnabled                      string
	_mediatailorPersonalizationThresholdSeconds     string
	_mediatailorPlaybackConfigurationName           string
	_mediatailorPlaybackMode                        string
	_mediatailorPolicy                              string
	_mediatailorProgramName                         string
	_mediatailorRecurringPrefetchConfiguration      string
	_mediatailorResourceArn                         string
	_mediatailorRetrieval                           string
	_mediatailorScheduleConfiguration               string
	_mediatailorScheduleType                        string
	_mediatailorSegmentDeliveryConfigurations       string
	_mediatailorSlateAdUrl                          string
	_mediatailorSourceLocationName                  string
	_mediatailorStreamId                            string
	_mediatailorTagKeys                             []string
	_mediatailorTags                                string
	_mediatailorTier                                string
	_mediatailorTimeShiftConfiguration              string
	_mediatailorTranscodeProfileName                string
	_mediatailorVideoContentSourceUrl               string
	_mediatailorVodSourceName                       string
)

// Configures Amazon CloudWatch log settings for a channel.
func mediatailor_ConfigureLogsForChannel(cfg aws.Config, client *mediatailor.Client) {
	input := &mediatailor.ConfigureLogsForChannelInput{
		// ChannelName: *string, // Required
		// LogTypes: []types.LogType, // Required
	}

	if len(_mediatailorChannelName) > 0 {
		input.ChannelName = aws.String(_mediatailorChannelName)
	}
	if len(_mediatailorLogTypes) > 0 {
		if err := assignInputField(input, "LogTypes", _mediatailorLogTypes); err != nil {
			log.Errorf("invalid --log-types: %s", err.Error())
			return
		}
	}

	if resp, err := client.ConfigureLogsForChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Defines where AWS Elemental MediaTailor sends logs for the playback
// configuration.
func mediatailor_ConfigureLogsForPlaybackConfiguration(cfg aws.Config, client *mediatailor.Client) {
	input := &mediatailor.ConfigureLogsForPlaybackConfigurationInput{
		// PercentEnabled: int32, // Required
		// PlaybackConfigurationName: *string, // Required
	}

	if len(_mediatailorPercentEnabled) > 0 {
		if err := assignInputField(input, "PercentEnabled", _mediatailorPercentEnabled); err != nil {
			log.Errorf("invalid --percent-enabled: %s", err.Error())
			return
		}
	}
	if len(_mediatailorPlaybackConfigurationName) > 0 {
		input.PlaybackConfigurationName = aws.String(_mediatailorPlaybackConfigurationName)
	}
	if len(_mediatailorAdsInteractionLog) > 0 {
		if err := assignInputField(input, "AdsInteractionLog", _mediatailorAdsInteractionLog); err != nil {
			log.Errorf("invalid --ads-interaction-log: %s", err.Error())
			return
		}
	}
	if len(_mediatailorEnabledLoggingStrategies) > 0 {
		if err := assignInputField(input, "EnabledLoggingStrategies", _mediatailorEnabledLoggingStrategies); err != nil {
			log.Errorf("invalid --enabled-logging-strategies: %s", err.Error())
			return
		}
	}
	if len(_mediatailorManifestServiceInteractionLog) > 0 {
		if err := assignInputField(input, "ManifestServiceInteractionLog", _mediatailorManifestServiceInteractionLog); err != nil {
			log.Errorf("invalid --manifest-service-interaction-log: %s", err.Error())
			return
		}
	}

	if resp, err := client.ConfigureLogsForPlaybackConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a channel. For information about MediaTailor channels, see [Working with channels] in the
// MediaTailor User Guide.
//
// [Working with channels]: https://docs.aws.amazon.com/mediatailor/latest/ug/channel-assembly-channels.html
func mediatailor_CreateChannel(cfg aws.Config, client *mediatailor.Client) {
	input := &mediatailor.CreateChannelInput{
		// ChannelName: *string, // Required
		// Outputs: []types.RequestOutputItem, // Required
		// PlaybackMode: types.PlaybackMode, // Required
	}

	if len(_mediatailorChannelName) > 0 {
		input.ChannelName = aws.String(_mediatailorChannelName)
	}
	if len(_mediatailorOutputs) > 0 {
		if err := assignInputField(input, "Outputs", _mediatailorOutputs); err != nil {
			log.Errorf("invalid --outputs: %s", err.Error())
			return
		}
	}
	if len(_mediatailorPlaybackMode) > 0 {
		if err := assignInputField(input, "PlaybackMode", _mediatailorPlaybackMode); err != nil {
			log.Errorf("invalid --playback-mode: %s", err.Error())
			return
		}
	}
	if len(_mediatailorAudiences) > 0 {
		input.Audiences = append([]string(nil), _mediatailorAudiences...)
	}
	if len(_mediatailorFillerSlate) > 0 {
		if err := assignInputField(input, "FillerSlate", _mediatailorFillerSlate); err != nil {
			log.Errorf("invalid --filler-slate: %s", err.Error())
			return
		}
	}
	if len(_mediatailorTags) > 0 {
		if err := assignInputField(input, "Tags", _mediatailorTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_mediatailorTier) > 0 {
		if err := assignInputField(input, "Tier", _mediatailorTier); err != nil {
			log.Errorf("invalid --tier: %s", err.Error())
			return
		}
	}
	if len(_mediatailorTimeShiftConfiguration) > 0 {
		if err := assignInputField(input, "TimeShiftConfiguration", _mediatailorTimeShiftConfiguration); err != nil {
			log.Errorf("invalid --time-shift-configuration: %s", err.Error())
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

// The live source configuration.
func mediatailor_CreateLiveSource(cfg aws.Config, client *mediatailor.Client) {
	input := &mediatailor.CreateLiveSourceInput{
		// HttpPackageConfigurations: []types.HttpPackageConfiguration, // Required
		// LiveSourceName: *string, // Required
		// SourceLocationName: *string, // Required
	}

	if len(_mediatailorHttpPackageConfigurations) > 0 {
		if err := assignInputField(input, "HttpPackageConfigurations", _mediatailorHttpPackageConfigurations); err != nil {
			log.Errorf("invalid --http-package-configurations: %s", err.Error())
			return
		}
	}
	if len(_mediatailorLiveSourceName) > 0 {
		input.LiveSourceName = aws.String(_mediatailorLiveSourceName)
	}
	if len(_mediatailorSourceLocationName) > 0 {
		input.SourceLocationName = aws.String(_mediatailorSourceLocationName)
	}
	if len(_mediatailorTags) > 0 {
		if err := assignInputField(input, "Tags", _mediatailorTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateLiveSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a prefetch schedule for a playback configuration. A prefetch schedule
// allows you to tell MediaTailor to fetch and prepare certain ads before an ad
// break happens. For more information about ad prefetching, see [Using ad prefetching]in the
// MediaTailor User Guide.
//
// [Using ad prefetching]: https://docs.aws.amazon.com/mediatailor/latest/ug/prefetching-ads.html
func mediatailor_CreatePrefetchSchedule(cfg aws.Config, client *mediatailor.Client) {
	input := &mediatailor.CreatePrefetchScheduleInput{
		// Name: *string, // Required
		// PlaybackConfigurationName: *string, // Required
	}

	if len(_mediatailorName) > 0 {
		input.Name = aws.String(_mediatailorName)
	}
	if len(_mediatailorPlaybackConfigurationName) > 0 {
		input.PlaybackConfigurationName = aws.String(_mediatailorPlaybackConfigurationName)
	}
	if len(_mediatailorConsumption) > 0 {
		if err := assignInputField(input, "Consumption", _mediatailorConsumption); err != nil {
			log.Errorf("invalid --consumption: %s", err.Error())
			return
		}
	}
	if len(_mediatailorRecurringPrefetchConfiguration) > 0 {
		if err := assignInputField(input, "RecurringPrefetchConfiguration", _mediatailorRecurringPrefetchConfiguration); err != nil {
			log.Errorf("invalid --recurring-prefetch-configuration: %s", err.Error())
			return
		}
	}
	if len(_mediatailorRetrieval) > 0 {
		if err := assignInputField(input, "Retrieval", _mediatailorRetrieval); err != nil {
			log.Errorf("invalid --retrieval: %s", err.Error())
			return
		}
	}
	if len(_mediatailorScheduleType) > 0 {
		if err := assignInputField(input, "ScheduleType", _mediatailorScheduleType); err != nil {
			log.Errorf("invalid --schedule-type: %s", err.Error())
			return
		}
	}
	if len(_mediatailorStreamId) > 0 {
		input.StreamId = aws.String(_mediatailorStreamId)
	}

	if resp, err := client.CreatePrefetchSchedule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a program within a channel. For information about programs, see [Working with programs] in the
// MediaTailor User Guide.
//
// [Working with programs]: https://docs.aws.amazon.com/mediatailor/latest/ug/channel-assembly-programs.html
func mediatailor_CreateProgram(cfg aws.Config, client *mediatailor.Client) {
	input := &mediatailor.CreateProgramInput{
		// ChannelName: *string, // Required
		// ProgramName: *string, // Required
		// ScheduleConfiguration: *types.ScheduleConfiguration, // Required
		// SourceLocationName: *string, // Required
	}

	if len(_mediatailorChannelName) > 0 {
		input.ChannelName = aws.String(_mediatailorChannelName)
	}
	if len(_mediatailorProgramName) > 0 {
		input.ProgramName = aws.String(_mediatailorProgramName)
	}
	if len(_mediatailorScheduleConfiguration) > 0 {
		if err := assignInputField(input, "ScheduleConfiguration", _mediatailorScheduleConfiguration); err != nil {
			log.Errorf("invalid --schedule-configuration: %s", err.Error())
			return
		}
	}
	if len(_mediatailorSourceLocationName) > 0 {
		input.SourceLocationName = aws.String(_mediatailorSourceLocationName)
	}
	if len(_mediatailorAdBreaks) > 0 {
		if err := assignInputField(input, "AdBreaks", _mediatailorAdBreaks); err != nil {
			log.Errorf("invalid --ad-breaks: %s", err.Error())
			return
		}
	}
	if len(_mediatailorAudienceMedia) > 0 {
		if err := assignInputField(input, "AudienceMedia", _mediatailorAudienceMedia); err != nil {
			log.Errorf("invalid --audience-media: %s", err.Error())
			return
		}
	}
	if len(_mediatailorLiveSourceName) > 0 {
		input.LiveSourceName = aws.String(_mediatailorLiveSourceName)
	}
	if len(_mediatailorVodSourceName) > 0 {
		input.VodSourceName = aws.String(_mediatailorVodSourceName)
	}

	if resp, err := client.CreateProgram(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a source location. A source location is a container for sources. For
// more information about source locations, see [Working with source locations]in the MediaTailor User Guide.
//
// [Working with source locations]: https://docs.aws.amazon.com/mediatailor/latest/ug/channel-assembly-source-locations.html
func mediatailor_CreateSourceLocation(cfg aws.Config, client *mediatailor.Client) {
	input := &mediatailor.CreateSourceLocationInput{
		// HttpConfiguration: *types.HttpConfiguration, // Required
		// SourceLocationName: *string, // Required
	}

	if len(_mediatailorHttpConfiguration) > 0 {
		if err := assignInputField(input, "HttpConfiguration", _mediatailorHttpConfiguration); err != nil {
			log.Errorf("invalid --http-configuration: %s", err.Error())
			return
		}
	}
	if len(_mediatailorSourceLocationName) > 0 {
		input.SourceLocationName = aws.String(_mediatailorSourceLocationName)
	}
	if len(_mediatailorAccessConfiguration) > 0 {
		if err := assignInputField(input, "AccessConfiguration", _mediatailorAccessConfiguration); err != nil {
			log.Errorf("invalid --access-configuration: %s", err.Error())
			return
		}
	}
	if len(_mediatailorDefaultSegmentDeliveryConfiguration) > 0 {
		if err := assignInputField(input, "DefaultSegmentDeliveryConfiguration", _mediatailorDefaultSegmentDeliveryConfiguration); err != nil {
			log.Errorf("invalid --default-segment-delivery-configuration: %s", err.Error())
			return
		}
	}
	if len(_mediatailorSegmentDeliveryConfigurations) > 0 {
		if err := assignInputField(input, "SegmentDeliveryConfigurations", _mediatailorSegmentDeliveryConfigurations); err != nil {
			log.Errorf("invalid --segment-delivery-configurations: %s", err.Error())
			return
		}
	}
	if len(_mediatailorTags) > 0 {
		if err := assignInputField(input, "Tags", _mediatailorTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSourceLocation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The VOD source configuration parameters.
func mediatailor_CreateVodSource(cfg aws.Config, client *mediatailor.Client) {
	input := &mediatailor.CreateVodSourceInput{
		// HttpPackageConfigurations: []types.HttpPackageConfiguration, // Required
		// SourceLocationName: *string, // Required
		// VodSourceName: *string, // Required
	}

	if len(_mediatailorHttpPackageConfigurations) > 0 {
		if err := assignInputField(input, "HttpPackageConfigurations", _mediatailorHttpPackageConfigurations); err != nil {
			log.Errorf("invalid --http-package-configurations: %s", err.Error())
			return
		}
	}
	if len(_mediatailorSourceLocationName) > 0 {
		input.SourceLocationName = aws.String(_mediatailorSourceLocationName)
	}
	if len(_mediatailorVodSourceName) > 0 {
		input.VodSourceName = aws.String(_mediatailorVodSourceName)
	}
	if len(_mediatailorTags) > 0 {
		if err := assignInputField(input, "Tags", _mediatailorTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateVodSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a channel. For information about MediaTailor channels, see [Working with channels] in the
// MediaTailor User Guide.
//
// [Working with channels]: https://docs.aws.amazon.com/mediatailor/latest/ug/channel-assembly-channels.html
func mediatailor_DeleteChannel(cfg aws.Config, client *mediatailor.Client) {
	input := &mediatailor.DeleteChannelInput{
		// ChannelName: *string, // Required
	}

	if len(_mediatailorChannelName) > 0 {
		input.ChannelName = aws.String(_mediatailorChannelName)
	}

	if resp, err := client.DeleteChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The channel policy to delete.
func mediatailor_DeleteChannelPolicy(cfg aws.Config, client *mediatailor.Client) {
	input := &mediatailor.DeleteChannelPolicyInput{
		// ChannelName: *string, // Required
	}

	if len(_mediatailorChannelName) > 0 {
		input.ChannelName = aws.String(_mediatailorChannelName)
	}

	if resp, err := client.DeleteChannelPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The live source to delete.
func mediatailor_DeleteLiveSource(cfg aws.Config, client *mediatailor.Client) {
	input := &mediatailor.DeleteLiveSourceInput{
		// LiveSourceName: *string, // Required
		// SourceLocationName: *string, // Required
	}

	if len(_mediatailorLiveSourceName) > 0 {
		input.LiveSourceName = aws.String(_mediatailorLiveSourceName)
	}
	if len(_mediatailorSourceLocationName) > 0 {
		input.SourceLocationName = aws.String(_mediatailorSourceLocationName)
	}

	if resp, err := client.DeleteLiveSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a playback configuration. For information about MediaTailor
// configurations, see [Working with configurations in AWS Elemental MediaTailor].
//
// [Working with configurations in AWS Elemental MediaTailor]: https://docs.aws.amazon.com/mediatailor/latest/ug/configurations.html
func mediatailor_DeletePlaybackConfiguration(cfg aws.Config, client *mediatailor.Client) {
	input := &mediatailor.DeletePlaybackConfigurationInput{
		// Name: *string, // Required
	}

	if len(_mediatailorName) > 0 {
		input.Name = aws.String(_mediatailorName)
	}

	if resp, err := client.DeletePlaybackConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a prefetch schedule for a specific playback configuration. If you call
// DeletePrefetchSchedule on an expired prefetch schedule, MediaTailor returns an
// HTTP 404 status code. For more information about ad prefetching, see [Using ad prefetching]in the
// MediaTailor User Guide.
//
// [Using ad prefetching]: https://docs.aws.amazon.com/mediatailor/latest/ug/prefetching-ads.html
func mediatailor_DeletePrefetchSchedule(cfg aws.Config, client *mediatailor.Client) {
	input := &mediatailor.DeletePrefetchScheduleInput{
		// Name: *string, // Required
		// PlaybackConfigurationName: *string, // Required
	}

	if len(_mediatailorName) > 0 {
		input.Name = aws.String(_mediatailorName)
	}
	if len(_mediatailorPlaybackConfigurationName) > 0 {
		input.PlaybackConfigurationName = aws.String(_mediatailorPlaybackConfigurationName)
	}

	if resp, err := client.DeletePrefetchSchedule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a program within a channel. For information about programs, see [Working with programs] in the
// MediaTailor User Guide.
//
// [Working with programs]: https://docs.aws.amazon.com/mediatailor/latest/ug/channel-assembly-programs.html
func mediatailor_DeleteProgram(cfg aws.Config, client *mediatailor.Client) {
	input := &mediatailor.DeleteProgramInput{
		// ChannelName: *string, // Required
		// ProgramName: *string, // Required
	}

	if len(_mediatailorChannelName) > 0 {
		input.ChannelName = aws.String(_mediatailorChannelName)
	}
	if len(_mediatailorProgramName) > 0 {
		input.ProgramName = aws.String(_mediatailorProgramName)
	}

	if resp, err := client.DeleteProgram(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a source location. A source location is a container for sources. For
// more information about source locations, see [Working with source locations]in the MediaTailor User Guide.
//
// [Working with source locations]: https://docs.aws.amazon.com/mediatailor/latest/ug/channel-assembly-source-locations.html
func mediatailor_DeleteSourceLocation(cfg aws.Config, client *mediatailor.Client) {
	input := &mediatailor.DeleteSourceLocationInput{
		// SourceLocationName: *string, // Required
	}

	if len(_mediatailorSourceLocationName) > 0 {
		input.SourceLocationName = aws.String(_mediatailorSourceLocationName)
	}

	if resp, err := client.DeleteSourceLocation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The video on demand (VOD) source to delete.
func mediatailor_DeleteVodSource(cfg aws.Config, client *mediatailor.Client) {
	input := &mediatailor.DeleteVodSourceInput{
		// SourceLocationName: *string, // Required
		// VodSourceName: *string, // Required
	}

	if len(_mediatailorSourceLocationName) > 0 {
		input.SourceLocationName = aws.String(_mediatailorSourceLocationName)
	}
	if len(_mediatailorVodSourceName) > 0 {
		input.VodSourceName = aws.String(_mediatailorVodSourceName)
	}

	if resp, err := client.DeleteVodSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a channel. For information about MediaTailor channels, see [Working with channels] in the
// MediaTailor User Guide.
//
// [Working with channels]: https://docs.aws.amazon.com/mediatailor/latest/ug/channel-assembly-channels.html
func mediatailor_DescribeChannel(cfg aws.Config, client *mediatailor.Client) {
	input := &mediatailor.DescribeChannelInput{
		// ChannelName: *string, // Required
	}

	if len(_mediatailorChannelName) > 0 {
		input.ChannelName = aws.String(_mediatailorChannelName)
	}

	if resp, err := client.DescribeChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The live source to describe.
func mediatailor_DescribeLiveSource(cfg aws.Config, client *mediatailor.Client) {
	input := &mediatailor.DescribeLiveSourceInput{
		// LiveSourceName: *string, // Required
		// SourceLocationName: *string, // Required
	}

	if len(_mediatailorLiveSourceName) > 0 {
		input.LiveSourceName = aws.String(_mediatailorLiveSourceName)
	}
	if len(_mediatailorSourceLocationName) > 0 {
		input.SourceLocationName = aws.String(_mediatailorSourceLocationName)
	}

	if resp, err := client.DescribeLiveSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a program within a channel. For information about programs, see [Working with programs] in
// the MediaTailor User Guide.
//
// [Working with programs]: https://docs.aws.amazon.com/mediatailor/latest/ug/channel-assembly-programs.html
func mediatailor_DescribeProgram(cfg aws.Config, client *mediatailor.Client) {
	input := &mediatailor.DescribeProgramInput{
		// ChannelName: *string, // Required
		// ProgramName: *string, // Required
	}

	if len(_mediatailorChannelName) > 0 {
		input.ChannelName = aws.String(_mediatailorChannelName)
	}
	if len(_mediatailorProgramName) > 0 {
		input.ProgramName = aws.String(_mediatailorProgramName)
	}

	if resp, err := client.DescribeProgram(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a source location. A source location is a container for sources. For
// more information about source locations, see [Working with source locations]in the MediaTailor User Guide.
//
// [Working with source locations]: https://docs.aws.amazon.com/mediatailor/latest/ug/channel-assembly-source-locations.html
func mediatailor_DescribeSourceLocation(cfg aws.Config, client *mediatailor.Client) {
	input := &mediatailor.DescribeSourceLocationInput{
		// SourceLocationName: *string, // Required
	}

	if len(_mediatailorSourceLocationName) > 0 {
		input.SourceLocationName = aws.String(_mediatailorSourceLocationName)
	}

	if resp, err := client.DescribeSourceLocation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides details about a specific video on demand (VOD) source in a specific
// source location.
func mediatailor_DescribeVodSource(cfg aws.Config, client *mediatailor.Client) {
	input := &mediatailor.DescribeVodSourceInput{
		// SourceLocationName: *string, // Required
		// VodSourceName: *string, // Required
	}

	if len(_mediatailorSourceLocationName) > 0 {
		input.SourceLocationName = aws.String(_mediatailorSourceLocationName)
	}
	if len(_mediatailorVodSourceName) > 0 {
		input.VodSourceName = aws.String(_mediatailorVodSourceName)
	}

	if resp, err := client.DescribeVodSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the channel's IAM policy. IAM policies are used to control access to
// your channel.
func mediatailor_GetChannelPolicy(cfg aws.Config, client *mediatailor.Client) {
	input := &mediatailor.GetChannelPolicyInput{
		// ChannelName: *string, // Required
	}

	if len(_mediatailorChannelName) > 0 {
		input.ChannelName = aws.String(_mediatailorChannelName)
	}

	if resp, err := client.GetChannelPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about your channel's schedule.
func mediatailor_GetChannelSchedule(cfg aws.Config, client *mediatailor.Client) {
	input := &mediatailor.GetChannelScheduleInput{
		// ChannelName: *string, // Required
	}

	if len(_mediatailorChannelName) > 0 {
		input.ChannelName = aws.String(_mediatailorChannelName)
	}
	if len(_mediatailorAudience) > 0 {
		input.Audience = aws.String(_mediatailorAudience)
	}
	if len(_mediatailorDurationMinutes) > 0 {
		input.DurationMinutes = aws.String(_mediatailorDurationMinutes)
	}
	if len(_mediatailorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mediatailorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mediatailorNextToken) > 0 {
		input.NextToken = aws.String(_mediatailorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetChannelSchedule(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mediatailor.GetChannelScheduleOutput
	p := mediatailor.NewGetChannelSchedulePaginator(client, input)
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

// Retrieves a playback configuration. For information about MediaTailor
// configurations, see [Working with configurations in AWS Elemental MediaTailor].
//
// [Working with configurations in AWS Elemental MediaTailor]: https://docs.aws.amazon.com/mediatailor/latest/ug/configurations.html
func mediatailor_GetPlaybackConfiguration(cfg aws.Config, client *mediatailor.Client) {
	input := &mediatailor.GetPlaybackConfigurationInput{
		// Name: *string, // Required
	}

	if len(_mediatailorName) > 0 {
		input.Name = aws.String(_mediatailorName)
	}

	if resp, err := client.GetPlaybackConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a prefetch schedule for a playback configuration. A prefetch schedule
// allows you to tell MediaTailor to fetch and prepare certain ads before an ad
// break happens. For more information about ad prefetching, see [Using ad prefetching]in the
// MediaTailor User Guide.
//
// [Using ad prefetching]: https://docs.aws.amazon.com/mediatailor/latest/ug/prefetching-ads.html
func mediatailor_GetPrefetchSchedule(cfg aws.Config, client *mediatailor.Client) {
	input := &mediatailor.GetPrefetchScheduleInput{
		// Name: *string, // Required
		// PlaybackConfigurationName: *string, // Required
	}

	if len(_mediatailorName) > 0 {
		input.Name = aws.String(_mediatailorName)
	}
	if len(_mediatailorPlaybackConfigurationName) > 0 {
		input.PlaybackConfigurationName = aws.String(_mediatailorPlaybackConfigurationName)
	}

	if resp, err := client.GetPrefetchSchedule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the alerts that are associated with a MediaTailor channel assembly
// resource.
func mediatailor_ListAlerts(cfg aws.Config, client *mediatailor.Client) {
	input := &mediatailor.ListAlertsInput{
		// ResourceArn: *string, // Required
	}

	if len(_mediatailorResourceArn) > 0 {
		input.ResourceArn = aws.String(_mediatailorResourceArn)
	}
	if len(_mediatailorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mediatailorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mediatailorNextToken) > 0 {
		input.NextToken = aws.String(_mediatailorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAlerts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mediatailor.ListAlertsOutput
	p := mediatailor.NewListAlertsPaginator(client, input)
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

// Retrieves information about the channels that are associated with the current
// AWS account.
func mediatailor_ListChannels(cfg aws.Config, client *mediatailor.Client) {
	input := &mediatailor.ListChannelsInput{}

	if len(_mediatailorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mediatailorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mediatailorNextToken) > 0 {
		input.NextToken = aws.String(_mediatailorNextToken)
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

	var results []*mediatailor.ListChannelsOutput
	p := mediatailor.NewListChannelsPaginator(client, input)
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

// Lists the live sources contained in a source location. A source represents a
// piece of content.
func mediatailor_ListLiveSources(cfg aws.Config, client *mediatailor.Client) {
	input := &mediatailor.ListLiveSourcesInput{
		// SourceLocationName: *string, // Required
	}

	if len(_mediatailorSourceLocationName) > 0 {
		input.SourceLocationName = aws.String(_mediatailorSourceLocationName)
	}
	if len(_mediatailorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mediatailorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mediatailorNextToken) > 0 {
		input.NextToken = aws.String(_mediatailorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListLiveSources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mediatailor.ListLiveSourcesOutput
	p := mediatailor.NewListLiveSourcesPaginator(client, input)
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

// Retrieves existing playback configurations. For information about MediaTailor
// configurations, see [Working with Configurations in AWS Elemental MediaTailor].
//
// [Working with Configurations in AWS Elemental MediaTailor]: https://docs.aws.amazon.com/mediatailor/latest/ug/configurations.html
func mediatailor_ListPlaybackConfigurations(cfg aws.Config, client *mediatailor.Client) {
	input := &mediatailor.ListPlaybackConfigurationsInput{}

	if len(_mediatailorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mediatailorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mediatailorNextToken) > 0 {
		input.NextToken = aws.String(_mediatailorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPlaybackConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mediatailor.ListPlaybackConfigurationsOutput
	p := mediatailor.NewListPlaybackConfigurationsPaginator(client, input)
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

// Lists the prefetch schedules for a playback configuration.
func mediatailor_ListPrefetchSchedules(cfg aws.Config, client *mediatailor.Client) {
	input := &mediatailor.ListPrefetchSchedulesInput{
		// PlaybackConfigurationName: *string, // Required
	}

	if len(_mediatailorPlaybackConfigurationName) > 0 {
		input.PlaybackConfigurationName = aws.String(_mediatailorPlaybackConfigurationName)
	}
	if len(_mediatailorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mediatailorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mediatailorNextToken) > 0 {
		input.NextToken = aws.String(_mediatailorNextToken)
	}
	if len(_mediatailorScheduleType) > 0 {
		if err := assignInputField(input, "ScheduleType", _mediatailorScheduleType); err != nil {
			log.Errorf("invalid --schedule-type: %s", err.Error())
			return
		}
	}
	if len(_mediatailorStreamId) > 0 {
		input.StreamId = aws.String(_mediatailorStreamId)
	}

	if disablePaginator() {
		if resp, err := client.ListPrefetchSchedules(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mediatailor.ListPrefetchSchedulesOutput
	p := mediatailor.NewListPrefetchSchedulesPaginator(client, input)
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

// Lists the source locations for a channel. A source location defines the host
// server URL, and contains a list of sources.
func mediatailor_ListSourceLocations(cfg aws.Config, client *mediatailor.Client) {
	input := &mediatailor.ListSourceLocationsInput{}

	if len(_mediatailorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mediatailorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mediatailorNextToken) > 0 {
		input.NextToken = aws.String(_mediatailorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSourceLocations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mediatailor.ListSourceLocationsOutput
	p := mediatailor.NewListSourceLocationsPaginator(client, input)
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

// A list of tags that are associated with this resource. Tags are key-value pairs
// that you can associate with Amazon resources to help with organization, access
// control, and cost tracking. For more information, see [Tagging AWS Elemental MediaTailor Resources].
//
// [Tagging AWS Elemental MediaTailor Resources]: https://docs.aws.amazon.com/mediatailor/latest/ug/tagging.html
func mediatailor_ListTagsForResource(cfg aws.Config, client *mediatailor.Client) {
	input := &mediatailor.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_mediatailorResourceArn) > 0 {
		input.ResourceArn = aws.String(_mediatailorResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the VOD sources contained in a source location. A source represents a
// piece of content.
func mediatailor_ListVodSources(cfg aws.Config, client *mediatailor.Client) {
	input := &mediatailor.ListVodSourcesInput{
		// SourceLocationName: *string, // Required
	}

	if len(_mediatailorSourceLocationName) > 0 {
		input.SourceLocationName = aws.String(_mediatailorSourceLocationName)
	}
	if len(_mediatailorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mediatailorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mediatailorNextToken) > 0 {
		input.NextToken = aws.String(_mediatailorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListVodSources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mediatailor.ListVodSourcesOutput
	p := mediatailor.NewListVodSourcesPaginator(client, input)
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

// Creates an IAM policy for the channel. IAM policies are used to control access
// to your channel.
func mediatailor_PutChannelPolicy(cfg aws.Config, client *mediatailor.Client) {
	input := &mediatailor.PutChannelPolicyInput{
		// ChannelName: *string, // Required
		// Policy: *string, // Required
	}

	if len(_mediatailorChannelName) > 0 {
		input.ChannelName = aws.String(_mediatailorChannelName)
	}
	if len(_mediatailorPolicy) > 0 {
		input.Policy = aws.String(_mediatailorPolicy)
	}

	if resp, err := client.PutChannelPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a playback configuration. For information about MediaTailor
// configurations, see [Working with configurations in AWS Elemental MediaTailor].
//
// [Working with configurations in AWS Elemental MediaTailor]: https://docs.aws.amazon.com/mediatailor/latest/ug/configurations.html
func mediatailor_PutPlaybackConfiguration(cfg aws.Config, client *mediatailor.Client) {
	input := &mediatailor.PutPlaybackConfigurationInput{
		// Name: *string, // Required
	}

	if len(_mediatailorName) > 0 {
		input.Name = aws.String(_mediatailorName)
	}
	if len(_mediatailorAdConditioningConfiguration) > 0 {
		if err := assignInputField(input, "AdConditioningConfiguration", _mediatailorAdConditioningConfiguration); err != nil {
			log.Errorf("invalid --ad-conditioning-configuration: %s", err.Error())
			return
		}
	}
	if len(_mediatailorAdDecisionServerConfiguration) > 0 {
		if err := assignInputField(input, "AdDecisionServerConfiguration", _mediatailorAdDecisionServerConfiguration); err != nil {
			log.Errorf("invalid --ad-decision-server-configuration: %s", err.Error())
			return
		}
	}
	if len(_mediatailorAdDecisionServerUrl) > 0 {
		input.AdDecisionServerUrl = aws.String(_mediatailorAdDecisionServerUrl)
	}
	if len(_mediatailorAvailSuppression) > 0 {
		if err := assignInputField(input, "AvailSuppression", _mediatailorAvailSuppression); err != nil {
			log.Errorf("invalid --avail-suppression: %s", err.Error())
			return
		}
	}
	if len(_mediatailorBumper) > 0 {
		if err := assignInputField(input, "Bumper", _mediatailorBumper); err != nil {
			log.Errorf("invalid --bumper: %s", err.Error())
			return
		}
	}
	if len(_mediatailorCdnConfiguration) > 0 {
		if err := assignInputField(input, "CdnConfiguration", _mediatailorCdnConfiguration); err != nil {
			log.Errorf("invalid --cdn-configuration: %s", err.Error())
			return
		}
	}
	if len(_mediatailorConfigurationAliases) > 0 {
		if err := assignInputField(input, "ConfigurationAliases", _mediatailorConfigurationAliases); err != nil {
			log.Errorf("invalid --configuration-aliases: %s", err.Error())
			return
		}
	}
	if len(_mediatailorDashConfiguration) > 0 {
		if err := assignInputField(input, "DashConfiguration", _mediatailorDashConfiguration); err != nil {
			log.Errorf("invalid --dash-configuration: %s", err.Error())
			return
		}
	}
	if len(_mediatailorInsertionMode) > 0 {
		if err := assignInputField(input, "InsertionMode", _mediatailorInsertionMode); err != nil {
			log.Errorf("invalid --insertion-mode: %s", err.Error())
			return
		}
	}
	if len(_mediatailorLivePreRollConfiguration) > 0 {
		if err := assignInputField(input, "LivePreRollConfiguration", _mediatailorLivePreRollConfiguration); err != nil {
			log.Errorf("invalid --live-pre-roll-configuration: %s", err.Error())
			return
		}
	}
	if len(_mediatailorManifestProcessingRules) > 0 {
		if err := assignInputField(input, "ManifestProcessingRules", _mediatailorManifestProcessingRules); err != nil {
			log.Errorf("invalid --manifest-processing-rules: %s", err.Error())
			return
		}
	}
	if len(_mediatailorPersonalizationThresholdSeconds) > 0 {
		if err := assignInputField(input, "PersonalizationThresholdSeconds", _mediatailorPersonalizationThresholdSeconds); err != nil {
			log.Errorf("invalid --personalization-threshold-seconds: %s", err.Error())
			return
		}
	}
	if len(_mediatailorSlateAdUrl) > 0 {
		input.SlateAdUrl = aws.String(_mediatailorSlateAdUrl)
	}
	if len(_mediatailorTags) > 0 {
		if err := assignInputField(input, "Tags", _mediatailorTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_mediatailorTranscodeProfileName) > 0 {
		input.TranscodeProfileName = aws.String(_mediatailorTranscodeProfileName)
	}
	if len(_mediatailorVideoContentSourceUrl) > 0 {
		input.VideoContentSourceUrl = aws.String(_mediatailorVideoContentSourceUrl)
	}

	if resp, err := client.PutPlaybackConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a channel. For information about MediaTailor channels, see [Working with channels] in the
// MediaTailor User Guide.
//
// [Working with channels]: https://docs.aws.amazon.com/mediatailor/latest/ug/channel-assembly-channels.html
func mediatailor_StartChannel(cfg aws.Config, client *mediatailor.Client) {
	input := &mediatailor.StartChannelInput{
		// ChannelName: *string, // Required
	}

	if len(_mediatailorChannelName) > 0 {
		input.ChannelName = aws.String(_mediatailorChannelName)
	}

	if resp, err := client.StartChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a channel. For information about MediaTailor channels, see [Working with channels] in the
// MediaTailor User Guide.
//
// [Working with channels]: https://docs.aws.amazon.com/mediatailor/latest/ug/channel-assembly-channels.html
func mediatailor_StopChannel(cfg aws.Config, client *mediatailor.Client) {
	input := &mediatailor.StopChannelInput{
		// ChannelName: *string, // Required
	}

	if len(_mediatailorChannelName) > 0 {
		input.ChannelName = aws.String(_mediatailorChannelName)
	}

	if resp, err := client.StopChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The resource to tag. Tags are key-value pairs that you can associate with
// Amazon resources to help with organization, access control, and cost tracking.
// For more information, see [Tagging AWS Elemental MediaTailor Resources].
//
// [Tagging AWS Elemental MediaTailor Resources]: https://docs.aws.amazon.com/mediatailor/latest/ug/tagging.html
func mediatailor_TagResource(cfg aws.Config, client *mediatailor.Client) {
	input := &mediatailor.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_mediatailorResourceArn) > 0 {
		input.ResourceArn = aws.String(_mediatailorResourceArn)
	}
	if len(_mediatailorTags) > 0 {
		if err := assignInputField(input, "Tags", _mediatailorTags); err != nil {
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

// The resource to untag.
func mediatailor_UntagResource(cfg aws.Config, client *mediatailor.Client) {
	input := &mediatailor.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_mediatailorResourceArn) > 0 {
		input.ResourceArn = aws.String(_mediatailorResourceArn)
	}
	if len(_mediatailorTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _mediatailorTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a channel. For information about MediaTailor channels, see [Working with channels] in the
// MediaTailor User Guide.
//
// [Working with channels]: https://docs.aws.amazon.com/mediatailor/latest/ug/channel-assembly-channels.html
func mediatailor_UpdateChannel(cfg aws.Config, client *mediatailor.Client) {
	input := &mediatailor.UpdateChannelInput{
		// ChannelName: *string, // Required
		// Outputs: []types.RequestOutputItem, // Required
	}

	if len(_mediatailorChannelName) > 0 {
		input.ChannelName = aws.String(_mediatailorChannelName)
	}
	if len(_mediatailorOutputs) > 0 {
		if err := assignInputField(input, "Outputs", _mediatailorOutputs); err != nil {
			log.Errorf("invalid --outputs: %s", err.Error())
			return
		}
	}
	if len(_mediatailorAudiences) > 0 {
		input.Audiences = append([]string(nil), _mediatailorAudiences...)
	}
	if len(_mediatailorFillerSlate) > 0 {
		if err := assignInputField(input, "FillerSlate", _mediatailorFillerSlate); err != nil {
			log.Errorf("invalid --filler-slate: %s", err.Error())
			return
		}
	}
	if len(_mediatailorTimeShiftConfiguration) > 0 {
		if err := assignInputField(input, "TimeShiftConfiguration", _mediatailorTimeShiftConfiguration); err != nil {
			log.Errorf("invalid --time-shift-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a live source's configuration.
func mediatailor_UpdateLiveSource(cfg aws.Config, client *mediatailor.Client) {
	input := &mediatailor.UpdateLiveSourceInput{
		// HttpPackageConfigurations: []types.HttpPackageConfiguration, // Required
		// LiveSourceName: *string, // Required
		// SourceLocationName: *string, // Required
	}

	if len(_mediatailorHttpPackageConfigurations) > 0 {
		if err := assignInputField(input, "HttpPackageConfigurations", _mediatailorHttpPackageConfigurations); err != nil {
			log.Errorf("invalid --http-package-configurations: %s", err.Error())
			return
		}
	}
	if len(_mediatailorLiveSourceName) > 0 {
		input.LiveSourceName = aws.String(_mediatailorLiveSourceName)
	}
	if len(_mediatailorSourceLocationName) > 0 {
		input.SourceLocationName = aws.String(_mediatailorSourceLocationName)
	}

	if resp, err := client.UpdateLiveSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a program within a channel.
func mediatailor_UpdateProgram(cfg aws.Config, client *mediatailor.Client) {
	input := &mediatailor.UpdateProgramInput{
		// ChannelName: *string, // Required
		// ProgramName: *string, // Required
		// ScheduleConfiguration: *types.UpdateProgramScheduleConfiguration, // Required
	}

	if len(_mediatailorChannelName) > 0 {
		input.ChannelName = aws.String(_mediatailorChannelName)
	}
	if len(_mediatailorProgramName) > 0 {
		input.ProgramName = aws.String(_mediatailorProgramName)
	}
	if len(_mediatailorScheduleConfiguration) > 0 {
		if err := assignInputField(input, "ScheduleConfiguration", _mediatailorScheduleConfiguration); err != nil {
			log.Errorf("invalid --schedule-configuration: %s", err.Error())
			return
		}
	}
	if len(_mediatailorAdBreaks) > 0 {
		if err := assignInputField(input, "AdBreaks", _mediatailorAdBreaks); err != nil {
			log.Errorf("invalid --ad-breaks: %s", err.Error())
			return
		}
	}
	if len(_mediatailorAudienceMedia) > 0 {
		if err := assignInputField(input, "AudienceMedia", _mediatailorAudienceMedia); err != nil {
			log.Errorf("invalid --audience-media: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateProgram(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a source location. A source location is a container for sources. For
// more information about source locations, see [Working with source locations]in the MediaTailor User Guide.
//
// [Working with source locations]: https://docs.aws.amazon.com/mediatailor/latest/ug/channel-assembly-source-locations.html
func mediatailor_UpdateSourceLocation(cfg aws.Config, client *mediatailor.Client) {
	input := &mediatailor.UpdateSourceLocationInput{
		// HttpConfiguration: *types.HttpConfiguration, // Required
		// SourceLocationName: *string, // Required
	}

	if len(_mediatailorHttpConfiguration) > 0 {
		if err := assignInputField(input, "HttpConfiguration", _mediatailorHttpConfiguration); err != nil {
			log.Errorf("invalid --http-configuration: %s", err.Error())
			return
		}
	}
	if len(_mediatailorSourceLocationName) > 0 {
		input.SourceLocationName = aws.String(_mediatailorSourceLocationName)
	}
	if len(_mediatailorAccessConfiguration) > 0 {
		if err := assignInputField(input, "AccessConfiguration", _mediatailorAccessConfiguration); err != nil {
			log.Errorf("invalid --access-configuration: %s", err.Error())
			return
		}
	}
	if len(_mediatailorDefaultSegmentDeliveryConfiguration) > 0 {
		if err := assignInputField(input, "DefaultSegmentDeliveryConfiguration", _mediatailorDefaultSegmentDeliveryConfiguration); err != nil {
			log.Errorf("invalid --default-segment-delivery-configuration: %s", err.Error())
			return
		}
	}
	if len(_mediatailorSegmentDeliveryConfigurations) > 0 {
		if err := assignInputField(input, "SegmentDeliveryConfigurations", _mediatailorSegmentDeliveryConfigurations); err != nil {
			log.Errorf("invalid --segment-delivery-configurations: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSourceLocation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a VOD source's configuration.
func mediatailor_UpdateVodSource(cfg aws.Config, client *mediatailor.Client) {
	input := &mediatailor.UpdateVodSourceInput{
		// HttpPackageConfigurations: []types.HttpPackageConfiguration, // Required
		// SourceLocationName: *string, // Required
		// VodSourceName: *string, // Required
	}

	if len(_mediatailorHttpPackageConfigurations) > 0 {
		if err := assignInputField(input, "HttpPackageConfigurations", _mediatailorHttpPackageConfigurations); err != nil {
			log.Errorf("invalid --http-package-configurations: %s", err.Error())
			return
		}
	}
	if len(_mediatailorSourceLocationName) > 0 {
		input.SourceLocationName = aws.String(_mediatailorSourceLocationName)
	}
	if len(_mediatailorVodSourceName) > 0 {
		input.VodSourceName = aws.String(_mediatailorVodSourceName)
	}

	if resp, err := client.UpdateVodSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_mediatailorCmd)
	_mediatailorCmd.Flags().SortFlags = false

	_mediatailorCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_mediatailorCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_mediatailorCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_mediatailorCmd.Flags().StringVarP(&_mediatailorAccessConfiguration, "access-configuration", "", "", "Access Configuration")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorAdBreaks, "ad-breaks", "", "", "Ad Breaks")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorAdConditioningConfiguration, "ad-conditioning-configuration", "", "", "Ad Conditioning Configuration")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorAdDecisionServerConfiguration, "ad-decision-server-configuration", "", "", "Ad Decision Server Configuration")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorAdDecisionServerUrl, "ad-decision-server-url", "", "", "Ad Decision Server URL")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorAdsInteractionLog, "ads-interaction-log", "", "", "Ads Interaction Log")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorAudience, "audience", "", "", "Audience")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorAudienceMedia, "audience-media", "", "", "Audience Media")
	_mediatailorCmd.Flags().StringSliceVarP(&_mediatailorAudiences, "audiences", "", nil, "Audiences")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorAvailSuppression, "avail-suppression", "", "", "Avail Suppression")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorBumper, "bumper", "", "", "Bumper")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorCdnConfiguration, "cdn-configuration", "", "", "Cdn Configuration")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorChannelName, "channel-name", "", "", "Channel Name")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorConfigurationAliases, "configuration-aliases", "", "", "Configuration Aliases")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorConsumption, "consumption", "", "", "Consumption")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorDashConfiguration, "dash-configuration", "", "", "Dash Configuration")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorDefaultSegmentDeliveryConfiguration, "default-segment-delivery-configuration", "", "", "Default Segment Delivery Configuration")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorDurationMinutes, "duration-minutes", "", "", "Duration Minutes")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorEnabledLoggingStrategies, "enabled-logging-strategies", "", "", "Enabled Logging Strategies")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorFillerSlate, "filler-slate", "", "", "Filler Slate")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorHttpConfiguration, "http-configuration", "", "", "HTTP Configuration")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorHttpPackageConfigurations, "http-package-configurations", "", "", "HTTP Package Configurations")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorInsertionMode, "insertion-mode", "", "", "Insertion Mode")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorLivePreRollConfiguration, "live-pre-roll-configuration", "", "", "Live Pre Roll Configuration")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorLiveSourceName, "live-source-name", "", "", "Live Source Name")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorLogTypes, "log-types", "", "", "Log Types")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorManifestProcessingRules, "manifest-processing-rules", "", "", "Manifest Processing Rules")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorManifestServiceInteractionLog, "manifest-service-interaction-log", "", "", "Manifest Service Interaction Log")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorMaxResults, "max-results", "", "", "Max Results")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorName, "name", "", "", "Name")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorNextToken, "next-token", "", "", "Next Token")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorOutputs, "outputs", "", "", "Outputs")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorPercentEnabled, "percent-enabled", "", "", "Percent Enabled")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorPersonalizationThresholdSeconds, "personalization-threshold-seconds", "", "", "Personalization Threshold Seconds")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorPlaybackConfigurationName, "playback-configuration-name", "", "", "Playback Configuration Name")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorPlaybackMode, "playback-mode", "", "", "Playback Mode")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorPolicy, "policy", "", "", "Policy")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorProgramName, "program-name", "", "", "Program Name")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorRecurringPrefetchConfiguration, "recurring-prefetch-configuration", "", "", "Recurring Prefetch Configuration")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorResourceArn, "resource-arn", "", "", "Resource ARN")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorRetrieval, "retrieval", "", "", "Retrieval")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorScheduleConfiguration, "schedule-configuration", "", "", "Schedule Configuration")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorScheduleType, "schedule-type", "", "", "Schedule Type")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorSegmentDeliveryConfigurations, "segment-delivery-configurations", "", "", "Segment Delivery Configurations")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorSlateAdUrl, "slate-ad-url", "", "", "Slate Ad URL")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorSourceLocationName, "source-location-name", "", "", "Source Location Name")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorStreamId, "stream-id", "", "", "Stream ID")
	_mediatailorCmd.Flags().StringSliceVarP(&_mediatailorTagKeys, "tag-keys", "", nil, "Tag Keys")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorTags, "tags", "", "", "Tags")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorTier, "tier", "", "", "Tier")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorTimeShiftConfiguration, "time-shift-configuration", "", "", "Time Shift Configuration")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorTranscodeProfileName, "transcode-profile-name", "", "", "Transcode Profile Name")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorVideoContentSourceUrl, "video-content-source-url", "", "", "Video Content Source URL")
	_mediatailorCmd.Flags().StringVarP(&_mediatailorVodSourceName, "vod-source-name", "", "", "Vod Source Name")

	_mediatailorCmd.Flags().BoolVarP(&_mediatailorConfigureLogsForChannel, "configure-logs-for-channel", "", false, "Configure Logs For Channel")
	_mediatailorCmd.Flags().BoolVarP(&_mediatailorConfigureLogsForPlaybackConfiguration, "configure-logs-for-playback-configuration", "", false, "Configure Logs For Playback Configuration")
	_mediatailorCmd.Flags().BoolVarP(&_mediatailorCreateChannel, "create-channel", "", false, "Create Channel")
	_mediatailorCmd.Flags().BoolVarP(&_mediatailorCreateLiveSource, "create-live-source", "", false, "Create Live Source")
	_mediatailorCmd.Flags().BoolVarP(&_mediatailorCreatePrefetchSchedule, "create-prefetch-schedule", "", false, "Create Prefetch Schedule")
	_mediatailorCmd.Flags().BoolVarP(&_mediatailorCreateProgram, "create-program", "", false, "Create Program")
	_mediatailorCmd.Flags().BoolVarP(&_mediatailorCreateSourceLocation, "create-source-location", "", false, "Create Source Location")
	_mediatailorCmd.Flags().BoolVarP(&_mediatailorCreateVodSource, "create-vod-source", "", false, "Create Vod Source")
	_mediatailorCmd.Flags().BoolVarP(&_mediatailorDeleteChannel, "delete-channel", "", false, "Delete Channel")
	_mediatailorCmd.Flags().BoolVarP(&_mediatailorDeleteChannelPolicy, "delete-channel-policy", "", false, "Delete Channel Policy")
	_mediatailorCmd.Flags().BoolVarP(&_mediatailorDeleteLiveSource, "delete-live-source", "", false, "Delete Live Source")
	_mediatailorCmd.Flags().BoolVarP(&_mediatailorDeletePlaybackConfiguration, "delete-playback-configuration", "", false, "Delete Playback Configuration")
	_mediatailorCmd.Flags().BoolVarP(&_mediatailorDeletePrefetchSchedule, "delete-prefetch-schedule", "", false, "Delete Prefetch Schedule")
	_mediatailorCmd.Flags().BoolVarP(&_mediatailorDeleteProgram, "delete-program", "", false, "Delete Program")
	_mediatailorCmd.Flags().BoolVarP(&_mediatailorDeleteSourceLocation, "delete-source-location", "", false, "Delete Source Location")
	_mediatailorCmd.Flags().BoolVarP(&_mediatailorDeleteVodSource, "delete-vod-source", "", false, "Delete Vod Source")
	_mediatailorCmd.Flags().BoolVarP(&_mediatailorDescribeChannel, "describe-channel", "", false, "Describe Channel")
	_mediatailorCmd.Flags().BoolVarP(&_mediatailorDescribeLiveSource, "describe-live-source", "", false, "Describe Live Source")
	_mediatailorCmd.Flags().BoolVarP(&_mediatailorDescribeProgram, "describe-program", "", false, "Describe Program")
	_mediatailorCmd.Flags().BoolVarP(&_mediatailorDescribeSourceLocation, "describe-source-location", "", false, "Describe Source Location")
	_mediatailorCmd.Flags().BoolVarP(&_mediatailorDescribeVodSource, "describe-vod-source", "", false, "Describe Vod Source")
	_mediatailorCmd.Flags().BoolVarP(&_mediatailorGetChannelPolicy, "get-channel-policy", "", false, "Get Channel Policy")
	_mediatailorCmd.Flags().BoolVarP(&_mediatailorGetChannelSchedule, "get-channel-schedule", "", false, "Get Channel Schedule")
	_mediatailorCmd.Flags().BoolVarP(&_mediatailorGetPlaybackConfiguration, "get-playback-configuration", "", false, "Get Playback Configuration")
	_mediatailorCmd.Flags().BoolVarP(&_mediatailorGetPrefetchSchedule, "get-prefetch-schedule", "", false, "Get Prefetch Schedule")
	_mediatailorCmd.Flags().BoolVarP(&_mediatailorListAlerts, "list-alerts", "", false, "List Alerts")
	_mediatailorCmd.Flags().BoolVarP(&_mediatailorListChannels, "list-channels", "", false, "List Channels")
	_mediatailorCmd.Flags().BoolVarP(&_mediatailorListLiveSources, "list-live-sources", "", false, "List Live Sources")
	_mediatailorCmd.Flags().BoolVarP(&_mediatailorListPlaybackConfigurations, "list-playback-configurations", "", false, "List Playback Configurations")
	_mediatailorCmd.Flags().BoolVarP(&_mediatailorListPrefetchSchedules, "list-prefetch-schedules", "", false, "List Prefetch Schedules")
	_mediatailorCmd.Flags().BoolVarP(&_mediatailorListSourceLocations, "list-source-locations", "", false, "List Source Locations")
	_mediatailorCmd.Flags().BoolVarP(&_mediatailorListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_mediatailorCmd.Flags().BoolVarP(&_mediatailorListVodSources, "list-vod-sources", "", false, "List Vod Sources")
	_mediatailorCmd.Flags().BoolVarP(&_mediatailorPutChannelPolicy, "put-channel-policy", "", false, "Put Channel Policy")
	_mediatailorCmd.Flags().BoolVarP(&_mediatailorPutPlaybackConfiguration, "put-playback-configuration", "", false, "Put Playback Configuration")
	_mediatailorCmd.Flags().BoolVarP(&_mediatailorStartChannel, "start-channel", "", false, "Start Channel")
	_mediatailorCmd.Flags().BoolVarP(&_mediatailorStopChannel, "stop-channel", "", false, "Stop Channel")
	_mediatailorCmd.Flags().BoolVarP(&_mediatailorTagResource, "tag-resource", "", false, "Tag Resource")
	_mediatailorCmd.Flags().BoolVarP(&_mediatailorUntagResource, "untag-resource", "", false, "Untag Resource")
	_mediatailorCmd.Flags().BoolVarP(&_mediatailorUpdateChannel, "update-channel", "", false, "Update Channel")
	_mediatailorCmd.Flags().BoolVarP(&_mediatailorUpdateLiveSource, "update-live-source", "", false, "Update Live Source")
	_mediatailorCmd.Flags().BoolVarP(&_mediatailorUpdateProgram, "update-program", "", false, "Update Program")
	_mediatailorCmd.Flags().BoolVarP(&_mediatailorUpdateSourceLocation, "update-source-location", "", false, "Update Source Location")
	_mediatailorCmd.Flags().BoolVarP(&_mediatailorUpdateVodSource, "update-vod-source", "", false, "Update Vod Source")

}
