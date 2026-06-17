package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// devicefarmCmd represents the devicefarm command
var _devicefarmCmd = &cobra.Command{
	Use:   "devicefarm",
	Short: "AWS devicefarm CLI",
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
		client := devicefarm.NewFromConfig(cfg)
		if _devicefarmCreateDevicePool {
			devicefarm_CreateDevicePool(cfg, client)
			return
		}
		if _devicefarmCreateInstanceProfile {
			devicefarm_CreateInstanceProfile(cfg, client)
			return
		}
		if _devicefarmCreateNetworkProfile {
			devicefarm_CreateNetworkProfile(cfg, client)
			return
		}
		if _devicefarmCreateProject {
			devicefarm_CreateProject(cfg, client)
			return
		}
		if _devicefarmCreateRemoteAccessSession {
			devicefarm_CreateRemoteAccessSession(cfg, client)
			return
		}
		if _devicefarmCreateTestGridProject {
			devicefarm_CreateTestGridProject(cfg, client)
			return
		}
		if _devicefarmCreateTestGridUrl {
			devicefarm_CreateTestGridUrl(cfg, client)
			return
		}
		if _devicefarmCreateUpload {
			devicefarm_CreateUpload(cfg, client)
			return
		}
		if _devicefarmCreateVPCEConfiguration {
			devicefarm_CreateVPCEConfiguration(cfg, client)
			return
		}
		if _devicefarmDeleteDevicePool {
			devicefarm_DeleteDevicePool(cfg, client)
			return
		}
		if _devicefarmDeleteInstanceProfile {
			devicefarm_DeleteInstanceProfile(cfg, client)
			return
		}
		if _devicefarmDeleteNetworkProfile {
			devicefarm_DeleteNetworkProfile(cfg, client)
			return
		}
		if _devicefarmDeleteProject {
			devicefarm_DeleteProject(cfg, client)
			return
		}
		if _devicefarmDeleteRemoteAccessSession {
			devicefarm_DeleteRemoteAccessSession(cfg, client)
			return
		}
		if _devicefarmDeleteRun {
			devicefarm_DeleteRun(cfg, client)
			return
		}
		if _devicefarmDeleteTestGridProject {
			devicefarm_DeleteTestGridProject(cfg, client)
			return
		}
		if _devicefarmDeleteUpload {
			devicefarm_DeleteUpload(cfg, client)
			return
		}
		if _devicefarmDeleteVPCEConfiguration {
			devicefarm_DeleteVPCEConfiguration(cfg, client)
			return
		}
		if _devicefarmGetAccountSettings {
			devicefarm_GetAccountSettings(cfg, client)
			return
		}
		if _devicefarmGetDevice {
			devicefarm_GetDevice(cfg, client)
			return
		}
		if _devicefarmGetDeviceInstance {
			devicefarm_GetDeviceInstance(cfg, client)
			return
		}
		if _devicefarmGetDevicePool {
			devicefarm_GetDevicePool(cfg, client)
			return
		}
		if _devicefarmGetDevicePoolCompatibility {
			devicefarm_GetDevicePoolCompatibility(cfg, client)
			return
		}
		if _devicefarmGetInstanceProfile {
			devicefarm_GetInstanceProfile(cfg, client)
			return
		}
		if _devicefarmGetJob {
			devicefarm_GetJob(cfg, client)
			return
		}
		if _devicefarmGetNetworkProfile {
			devicefarm_GetNetworkProfile(cfg, client)
			return
		}
		if _devicefarmGetOfferingStatus {
			devicefarm_GetOfferingStatus(cfg, client)
			return
		}
		if _devicefarmGetProject {
			devicefarm_GetProject(cfg, client)
			return
		}
		if _devicefarmGetRemoteAccessSession {
			devicefarm_GetRemoteAccessSession(cfg, client)
			return
		}
		if _devicefarmGetRun {
			devicefarm_GetRun(cfg, client)
			return
		}
		if _devicefarmGetSuite {
			devicefarm_GetSuite(cfg, client)
			return
		}
		if _devicefarmGetTest {
			devicefarm_GetTest(cfg, client)
			return
		}
		if _devicefarmGetTestGridProject {
			devicefarm_GetTestGridProject(cfg, client)
			return
		}
		if _devicefarmGetTestGridSession {
			devicefarm_GetTestGridSession(cfg, client)
			return
		}
		if _devicefarmGetUpload {
			devicefarm_GetUpload(cfg, client)
			return
		}
		if _devicefarmGetVPCEConfiguration {
			devicefarm_GetVPCEConfiguration(cfg, client)
			return
		}
		if _devicefarmInstallToRemoteAccessSession {
			devicefarm_InstallToRemoteAccessSession(cfg, client)
			return
		}
		if _devicefarmListArtifacts {
			devicefarm_ListArtifacts(cfg, client)
			return
		}
		if _devicefarmListDeviceInstances {
			devicefarm_ListDeviceInstances(cfg, client)
			return
		}
		if _devicefarmListDevicePools {
			devicefarm_ListDevicePools(cfg, client)
			return
		}
		if _devicefarmListDevices {
			devicefarm_ListDevices(cfg, client)
			return
		}
		if _devicefarmListInstanceProfiles {
			devicefarm_ListInstanceProfiles(cfg, client)
			return
		}
		if _devicefarmListJobs {
			devicefarm_ListJobs(cfg, client)
			return
		}
		if _devicefarmListNetworkProfiles {
			devicefarm_ListNetworkProfiles(cfg, client)
			return
		}
		if _devicefarmListOfferingPromotions {
			devicefarm_ListOfferingPromotions(cfg, client)
			return
		}
		if _devicefarmListOfferingTransactions {
			devicefarm_ListOfferingTransactions(cfg, client)
			return
		}
		if _devicefarmListOfferings {
			devicefarm_ListOfferings(cfg, client)
			return
		}
		if _devicefarmListProjects {
			devicefarm_ListProjects(cfg, client)
			return
		}
		if _devicefarmListRemoteAccessSessions {
			devicefarm_ListRemoteAccessSessions(cfg, client)
			return
		}
		if _devicefarmListRuns {
			devicefarm_ListRuns(cfg, client)
			return
		}
		if _devicefarmListSamples {
			devicefarm_ListSamples(cfg, client)
			return
		}
		if _devicefarmListSuites {
			devicefarm_ListSuites(cfg, client)
			return
		}
		if _devicefarmListTagsForResource {
			devicefarm_ListTagsForResource(cfg, client)
			return
		}
		if _devicefarmListTestGridProjects {
			devicefarm_ListTestGridProjects(cfg, client)
			return
		}
		if _devicefarmListTestGridSessionActions {
			devicefarm_ListTestGridSessionActions(cfg, client)
			return
		}
		if _devicefarmListTestGridSessionArtifacts {
			devicefarm_ListTestGridSessionArtifacts(cfg, client)
			return
		}
		if _devicefarmListTestGridSessions {
			devicefarm_ListTestGridSessions(cfg, client)
			return
		}
		if _devicefarmListTests {
			devicefarm_ListTests(cfg, client)
			return
		}
		if _devicefarmListUniqueProblems {
			devicefarm_ListUniqueProblems(cfg, client)
			return
		}
		if _devicefarmListUploads {
			devicefarm_ListUploads(cfg, client)
			return
		}
		if _devicefarmListVPCEConfigurations {
			devicefarm_ListVPCEConfigurations(cfg, client)
			return
		}
		if _devicefarmPurchaseOffering {
			devicefarm_PurchaseOffering(cfg, client)
			return
		}
		if _devicefarmRenewOffering {
			devicefarm_RenewOffering(cfg, client)
			return
		}
		if _devicefarmScheduleRun {
			devicefarm_ScheduleRun(cfg, client)
			return
		}
		if _devicefarmStopJob {
			devicefarm_StopJob(cfg, client)
			return
		}
		if _devicefarmStopRemoteAccessSession {
			devicefarm_StopRemoteAccessSession(cfg, client)
			return
		}
		if _devicefarmStopRun {
			devicefarm_StopRun(cfg, client)
			return
		}
		if _devicefarmTagResource {
			devicefarm_TagResource(cfg, client)
			return
		}
		if _devicefarmUntagResource {
			devicefarm_UntagResource(cfg, client)
			return
		}
		if _devicefarmUpdateDeviceInstance {
			devicefarm_UpdateDeviceInstance(cfg, client)
			return
		}
		if _devicefarmUpdateDevicePool {
			devicefarm_UpdateDevicePool(cfg, client)
			return
		}
		if _devicefarmUpdateInstanceProfile {
			devicefarm_UpdateInstanceProfile(cfg, client)
			return
		}
		if _devicefarmUpdateNetworkProfile {
			devicefarm_UpdateNetworkProfile(cfg, client)
			return
		}
		if _devicefarmUpdateProject {
			devicefarm_UpdateProject(cfg, client)
			return
		}
		if _devicefarmUpdateTestGridProject {
			devicefarm_UpdateTestGridProject(cfg, client)
			return
		}
		if _devicefarmUpdateUpload {
			devicefarm_UpdateUpload(cfg, client)
			return
		}
		if _devicefarmUpdateVPCEConfiguration {
			devicefarm_UpdateVPCEConfiguration(cfg, client)
			return
		}

	},
}

var (
	_devicefarmCreateDevicePool             bool
	_devicefarmCreateInstanceProfile        bool
	_devicefarmCreateNetworkProfile         bool
	_devicefarmCreateProject                bool
	_devicefarmCreateRemoteAccessSession    bool
	_devicefarmCreateTestGridProject        bool
	_devicefarmCreateTestGridUrl            bool
	_devicefarmCreateUpload                 bool
	_devicefarmCreateVPCEConfiguration      bool
	_devicefarmDeleteDevicePool             bool
	_devicefarmDeleteInstanceProfile        bool
	_devicefarmDeleteNetworkProfile         bool
	_devicefarmDeleteProject                bool
	_devicefarmDeleteRemoteAccessSession    bool
	_devicefarmDeleteRun                    bool
	_devicefarmDeleteTestGridProject        bool
	_devicefarmDeleteUpload                 bool
	_devicefarmDeleteVPCEConfiguration      bool
	_devicefarmGetAccountSettings           bool
	_devicefarmGetDevice                    bool
	_devicefarmGetDeviceInstance            bool
	_devicefarmGetDevicePool                bool
	_devicefarmGetDevicePoolCompatibility   bool
	_devicefarmGetInstanceProfile           bool
	_devicefarmGetJob                       bool
	_devicefarmGetNetworkProfile            bool
	_devicefarmGetOfferingStatus            bool
	_devicefarmGetProject                   bool
	_devicefarmGetRemoteAccessSession       bool
	_devicefarmGetRun                       bool
	_devicefarmGetSuite                     bool
	_devicefarmGetTest                      bool
	_devicefarmGetTestGridProject           bool
	_devicefarmGetTestGridSession           bool
	_devicefarmGetUpload                    bool
	_devicefarmGetVPCEConfiguration         bool
	_devicefarmInstallToRemoteAccessSession bool
	_devicefarmListArtifacts                bool
	_devicefarmListDeviceInstances          bool
	_devicefarmListDevicePools              bool
	_devicefarmListDevices                  bool
	_devicefarmListInstanceProfiles         bool
	_devicefarmListJobs                     bool
	_devicefarmListNetworkProfiles          bool
	_devicefarmListOfferingPromotions       bool
	_devicefarmListOfferingTransactions     bool
	_devicefarmListOfferings                bool
	_devicefarmListProjects                 bool
	_devicefarmListRemoteAccessSessions     bool
	_devicefarmListRuns                     bool
	_devicefarmListSamples                  bool
	_devicefarmListSuites                   bool
	_devicefarmListTagsForResource          bool
	_devicefarmListTestGridProjects         bool
	_devicefarmListTestGridSessionActions   bool
	_devicefarmListTestGridSessionArtifacts bool
	_devicefarmListTestGridSessions         bool
	_devicefarmListTests                    bool
	_devicefarmListUniqueProblems           bool
	_devicefarmListUploads                  bool
	_devicefarmListVPCEConfigurations       bool
	_devicefarmPurchaseOffering             bool
	_devicefarmRenewOffering                bool
	_devicefarmScheduleRun                  bool
	_devicefarmStopJob                      bool
	_devicefarmStopRemoteAccessSession      bool
	_devicefarmStopRun                      bool
	_devicefarmTagResource                  bool
	_devicefarmUntagResource                bool
	_devicefarmUpdateDeviceInstance         bool
	_devicefarmUpdateDevicePool             bool
	_devicefarmUpdateInstanceProfile        bool
	_devicefarmUpdateNetworkProfile         bool
	_devicefarmUpdateProject                bool
	_devicefarmUpdateTestGridProject        bool
	_devicefarmUpdateUpload                 bool
	_devicefarmUpdateVPCEConfiguration      bool

	_devicefarmAppArn                        string
	_devicefarmArn                           string
	_devicefarmClearMaxDevices               string
	_devicefarmConfiguration                 string
	_devicefarmContentType                   string
	_devicefarmCreationTimeAfter             string
	_devicefarmCreationTimeBefore            string
	_devicefarmDefaultJobTimeoutMinutes      string
	_devicefarmDescription                   string
	_devicefarmDeviceArn                     string
	_devicefarmDevicePoolArn                 string
	_devicefarmDeviceSelectionConfiguration  string
	_devicefarmDownlinkBandwidthBits         string
	_devicefarmDownlinkDelayMs               string
	_devicefarmDownlinkJitterMs              string
	_devicefarmDownlinkLossPercent           string
	_devicefarmEditContent                   string
	_devicefarmEndTimeAfter                  string
	_devicefarmEndTimeBefore                 string
	_devicefarmEnvironmentVariables          string
	_devicefarmExcludeAppPackagesFromCleanup []string
	_devicefarmExecutionConfiguration        string
	_devicefarmExecutionRoleArn              string
	_devicefarmExpiresInSeconds              string
	_devicefarmFilters                       string
	_devicefarmInstanceArn                   string
	_devicefarmInteractionMode               string
	_devicefarmLabels                        []string
	_devicefarmMaxDevices                    string
	_devicefarmMaxResult                     string
	_devicefarmMaxResults                    string
	_devicefarmName                          string
	_devicefarmNextToken                     string
	_devicefarmOfferingId                    string
	_devicefarmOfferingPromotionId           string
	_devicefarmPackageCleanup                string
	_devicefarmProfileArn                    string
	_devicefarmProjectArn                    string
	_devicefarmQuantity                      string
	_devicefarmRebootAfterUse                string
	_devicefarmRemoteAccessSessionArn        string
	_devicefarmResourceARN                   string
	_devicefarmRules                         string
	_devicefarmServiceDnsName                string
	_devicefarmSessionArn                    string
	_devicefarmSessionId                     string
	_devicefarmSkipAppResign                 string
	_devicefarmStatus                        string
	_devicefarmTagKeys                       []string
	_devicefarmTags                          string
	_devicefarmTest                          string
	_devicefarmTestType                      string
	_devicefarmType                          string
	_devicefarmUplinkBandwidthBits           string
	_devicefarmUplinkDelayMs                 string
	_devicefarmUplinkJitterMs                string
	_devicefarmUplinkLossPercent             string
	_devicefarmVpcConfig                     string
	_devicefarmVpceConfigurationDescription  string
	_devicefarmVpceConfigurationName         string
	_devicefarmVpceServiceName               string
)

// Creates a device pool.
func devicefarm_CreateDevicePool(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.CreateDevicePoolInput{
		// Name: *string, // Required
		// ProjectArn: *string, // Required
		// Rules: []types.Rule, // Required
	}

	if len(_devicefarmName) > 0 {
		input.Name = aws.String(_devicefarmName)
	}
	if len(_devicefarmProjectArn) > 0 {
		input.ProjectArn = aws.String(_devicefarmProjectArn)
	}
	if len(_devicefarmRules) > 0 {
		if err := assignInputField(input, "Rules", _devicefarmRules); err != nil {
			log.Errorf("invalid --rules: %s", err.Error())
			return
		}
	}
	if len(_devicefarmDescription) > 0 {
		input.Description = aws.String(_devicefarmDescription)
	}
	if len(_devicefarmMaxDevices) > 0 {
		if err := assignInputField(input, "MaxDevices", _devicefarmMaxDevices); err != nil {
			log.Errorf("invalid --max-devices: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDevicePool(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a profile that can be applied to one or more private fleet device
// instances.
func devicefarm_CreateInstanceProfile(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.CreateInstanceProfileInput{
		// Name: *string, // Required
	}

	if len(_devicefarmName) > 0 {
		input.Name = aws.String(_devicefarmName)
	}
	if len(_devicefarmDescription) > 0 {
		input.Description = aws.String(_devicefarmDescription)
	}
	if len(_devicefarmExcludeAppPackagesFromCleanup) > 0 {
		input.ExcludeAppPackagesFromCleanup = append([]string(nil), _devicefarmExcludeAppPackagesFromCleanup...)
	}
	if len(_devicefarmPackageCleanup) > 0 {
		if err := assignInputField(input, "PackageCleanup", _devicefarmPackageCleanup); err != nil {
			log.Errorf("invalid --package-cleanup: %s", err.Error())
			return
		}
	}
	if len(_devicefarmRebootAfterUse) > 0 {
		if err := assignInputField(input, "RebootAfterUse", _devicefarmRebootAfterUse); err != nil {
			log.Errorf("invalid --reboot-after-use: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateInstanceProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a network profile.
func devicefarm_CreateNetworkProfile(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.CreateNetworkProfileInput{
		// Name: *string, // Required
		// ProjectArn: *string, // Required
	}

	if len(_devicefarmName) > 0 {
		input.Name = aws.String(_devicefarmName)
	}
	if len(_devicefarmProjectArn) > 0 {
		input.ProjectArn = aws.String(_devicefarmProjectArn)
	}
	if len(_devicefarmDescription) > 0 {
		input.Description = aws.String(_devicefarmDescription)
	}
	if len(_devicefarmDownlinkBandwidthBits) > 0 {
		if err := assignInputField(input, "DownlinkBandwidthBits", _devicefarmDownlinkBandwidthBits); err != nil {
			log.Errorf("invalid --downlink-bandwidth-bits: %s", err.Error())
			return
		}
	}
	if len(_devicefarmDownlinkDelayMs) > 0 {
		if err := assignInputField(input, "DownlinkDelayMs", _devicefarmDownlinkDelayMs); err != nil {
			log.Errorf("invalid --downlink-delay-ms: %s", err.Error())
			return
		}
	}
	if len(_devicefarmDownlinkJitterMs) > 0 {
		if err := assignInputField(input, "DownlinkJitterMs", _devicefarmDownlinkJitterMs); err != nil {
			log.Errorf("invalid --downlink-jitter-ms: %s", err.Error())
			return
		}
	}
	if len(_devicefarmDownlinkLossPercent) > 0 {
		if err := assignInputField(input, "DownlinkLossPercent", _devicefarmDownlinkLossPercent); err != nil {
			log.Errorf("invalid --downlink-loss-percent: %s", err.Error())
			return
		}
	}
	if len(_devicefarmType) > 0 {
		if err := assignInputField(input, "Type", _devicefarmType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_devicefarmUplinkBandwidthBits) > 0 {
		if err := assignInputField(input, "UplinkBandwidthBits", _devicefarmUplinkBandwidthBits); err != nil {
			log.Errorf("invalid --uplink-bandwidth-bits: %s", err.Error())
			return
		}
	}
	if len(_devicefarmUplinkDelayMs) > 0 {
		if err := assignInputField(input, "UplinkDelayMs", _devicefarmUplinkDelayMs); err != nil {
			log.Errorf("invalid --uplink-delay-ms: %s", err.Error())
			return
		}
	}
	if len(_devicefarmUplinkJitterMs) > 0 {
		if err := assignInputField(input, "UplinkJitterMs", _devicefarmUplinkJitterMs); err != nil {
			log.Errorf("invalid --uplink-jitter-ms: %s", err.Error())
			return
		}
	}
	if len(_devicefarmUplinkLossPercent) > 0 {
		if err := assignInputField(input, "UplinkLossPercent", _devicefarmUplinkLossPercent); err != nil {
			log.Errorf("invalid --uplink-loss-percent: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateNetworkProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a project.
func devicefarm_CreateProject(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.CreateProjectInput{
		// Name: *string, // Required
	}

	if len(_devicefarmName) > 0 {
		input.Name = aws.String(_devicefarmName)
	}
	if len(_devicefarmDefaultJobTimeoutMinutes) > 0 {
		if err := assignInputField(input, "DefaultJobTimeoutMinutes", _devicefarmDefaultJobTimeoutMinutes); err != nil {
			log.Errorf("invalid --default-job-timeout-minutes: %s", err.Error())
			return
		}
	}
	if len(_devicefarmEnvironmentVariables) > 0 {
		if err := assignInputField(input, "EnvironmentVariables", _devicefarmEnvironmentVariables); err != nil {
			log.Errorf("invalid --environment-variables: %s", err.Error())
			return
		}
	}
	if len(_devicefarmExecutionRoleArn) > 0 {
		input.ExecutionRoleArn = aws.String(_devicefarmExecutionRoleArn)
	}
	if len(_devicefarmVpcConfig) > 0 {
		if err := assignInputField(input, "VpcConfig", _devicefarmVpcConfig); err != nil {
			log.Errorf("invalid --vpc-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateProject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Specifies and starts a remote access session.
func devicefarm_CreateRemoteAccessSession(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.CreateRemoteAccessSessionInput{
		// DeviceArn: *string, // Required
		// ProjectArn: *string, // Required
	}

	if len(_devicefarmDeviceArn) > 0 {
		input.DeviceArn = aws.String(_devicefarmDeviceArn)
	}
	if len(_devicefarmProjectArn) > 0 {
		input.ProjectArn = aws.String(_devicefarmProjectArn)
	}
	if len(_devicefarmAppArn) > 0 {
		input.AppArn = aws.String(_devicefarmAppArn)
	}
	if len(_devicefarmConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _devicefarmConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_devicefarmInstanceArn) > 0 {
		input.InstanceArn = aws.String(_devicefarmInstanceArn)
	}
	if len(_devicefarmInteractionMode) > 0 {
		if err := assignInputField(input, "InteractionMode", _devicefarmInteractionMode); err != nil {
			log.Errorf("invalid --interaction-mode: %s", err.Error())
			return
		}
	}
	if len(_devicefarmName) > 0 {
		input.Name = aws.String(_devicefarmName)
	}
	if len(_devicefarmSkipAppResign) > 0 {
		if err := assignInputField(input, "SkipAppResign", _devicefarmSkipAppResign); err != nil {
			log.Errorf("invalid --skip-app-resign: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRemoteAccessSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a Selenium testing project. Projects are used to track TestGridSession instances.
func devicefarm_CreateTestGridProject(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.CreateTestGridProjectInput{
		// Name: *string, // Required
	}

	if len(_devicefarmName) > 0 {
		input.Name = aws.String(_devicefarmName)
	}
	if len(_devicefarmDescription) > 0 {
		input.Description = aws.String(_devicefarmDescription)
	}
	if len(_devicefarmVpcConfig) > 0 {
		if err := assignInputField(input, "VpcConfig", _devicefarmVpcConfig); err != nil {
			log.Errorf("invalid --vpc-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTestGridProject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a signed, short-term URL that can be passed to a Selenium
// RemoteWebDriver constructor.
func devicefarm_CreateTestGridUrl(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.CreateTestGridUrlInput{
		// ExpiresInSeconds: *int32, // Required
		// ProjectArn: *string, // Required
	}

	if len(_devicefarmExpiresInSeconds) > 0 {
		if err := assignInputField(input, "ExpiresInSeconds", _devicefarmExpiresInSeconds); err != nil {
			log.Errorf("invalid --expires-in-seconds: %s", err.Error())
			return
		}
	}
	if len(_devicefarmProjectArn) > 0 {
		input.ProjectArn = aws.String(_devicefarmProjectArn)
	}

	if resp, err := client.CreateTestGridUrl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Uploads an app or test scripts.
func devicefarm_CreateUpload(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.CreateUploadInput{
		// Name: *string, // Required
		// ProjectArn: *string, // Required
		// Type: types.UploadType, // Required
	}

	if len(_devicefarmName) > 0 {
		input.Name = aws.String(_devicefarmName)
	}
	if len(_devicefarmProjectArn) > 0 {
		input.ProjectArn = aws.String(_devicefarmProjectArn)
	}
	if len(_devicefarmType) > 0 {
		if err := assignInputField(input, "Type", _devicefarmType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_devicefarmContentType) > 0 {
		input.ContentType = aws.String(_devicefarmContentType)
	}

	if resp, err := client.CreateUpload(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a configuration record in Device Farm for your Amazon Virtual Private
// Cloud (VPC) endpoint.
func devicefarm_CreateVPCEConfiguration(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.CreateVPCEConfigurationInput{
		// ServiceDnsName: *string, // Required
		// VpceConfigurationName: *string, // Required
		// VpceServiceName: *string, // Required
	}

	if len(_devicefarmServiceDnsName) > 0 {
		input.ServiceDnsName = aws.String(_devicefarmServiceDnsName)
	}
	if len(_devicefarmVpceConfigurationName) > 0 {
		input.VpceConfigurationName = aws.String(_devicefarmVpceConfigurationName)
	}
	if len(_devicefarmVpceServiceName) > 0 {
		input.VpceServiceName = aws.String(_devicefarmVpceServiceName)
	}
	if len(_devicefarmVpceConfigurationDescription) > 0 {
		input.VpceConfigurationDescription = aws.String(_devicefarmVpceConfigurationDescription)
	}

	if resp, err := client.CreateVPCEConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a device pool given the pool ARN. Does not allow deletion of curated
// pools owned by the system.
func devicefarm_DeleteDevicePool(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.DeleteDevicePoolInput{
		// Arn: *string, // Required
	}

	if len(_devicefarmArn) > 0 {
		input.Arn = aws.String(_devicefarmArn)
	}

	if resp, err := client.DeleteDevicePool(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a profile that can be applied to one or more private device instances.
func devicefarm_DeleteInstanceProfile(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.DeleteInstanceProfileInput{
		// Arn: *string, // Required
	}

	if len(_devicefarmArn) > 0 {
		input.Arn = aws.String(_devicefarmArn)
	}

	if resp, err := client.DeleteInstanceProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a network profile.
func devicefarm_DeleteNetworkProfile(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.DeleteNetworkProfileInput{
		// Arn: *string, // Required
	}

	if len(_devicefarmArn) > 0 {
		input.Arn = aws.String(_devicefarmArn)
	}

	if resp, err := client.DeleteNetworkProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an AWS Device Farm project, given the project ARN. You cannot delete a
// project if it has an active run or session.
//
// You cannot undo this operation.
func devicefarm_DeleteProject(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.DeleteProjectInput{
		// Arn: *string, // Required
	}

	if len(_devicefarmArn) > 0 {
		input.Arn = aws.String(_devicefarmArn)
	}

	if resp, err := client.DeleteProject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a completed remote access session and its results. You cannot delete a
// remote access session if it is still active.
//
// You cannot undo this operation.
func devicefarm_DeleteRemoteAccessSession(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.DeleteRemoteAccessSessionInput{
		// Arn: *string, // Required
	}

	if len(_devicefarmArn) > 0 {
		input.Arn = aws.String(_devicefarmArn)
	}

	if resp, err := client.DeleteRemoteAccessSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the run, given the run ARN. You cannot delete a run if it is still
// active.
//
// You cannot undo this operation.
func devicefarm_DeleteRun(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.DeleteRunInput{
		// Arn: *string, // Required
	}

	if len(_devicefarmArn) > 0 {
		input.Arn = aws.String(_devicefarmArn)
	}

	if resp, err := client.DeleteRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Selenium testing project and all content generated under it. You
// cannot delete a project if it has active sessions.
//
// You cannot undo this operation.
func devicefarm_DeleteTestGridProject(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.DeleteTestGridProjectInput{
		// ProjectArn: *string, // Required
	}

	if len(_devicefarmProjectArn) > 0 {
		input.ProjectArn = aws.String(_devicefarmProjectArn)
	}

	if resp, err := client.DeleteTestGridProject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an upload given the upload ARN.
func devicefarm_DeleteUpload(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.DeleteUploadInput{
		// Arn: *string, // Required
	}

	if len(_devicefarmArn) > 0 {
		input.Arn = aws.String(_devicefarmArn)
	}

	if resp, err := client.DeleteUpload(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a configuration for your Amazon Virtual Private Cloud (VPC) endpoint.
func devicefarm_DeleteVPCEConfiguration(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.DeleteVPCEConfigurationInput{
		// Arn: *string, // Required
	}

	if len(_devicefarmArn) > 0 {
		input.Arn = aws.String(_devicefarmArn)
	}

	if resp, err := client.DeleteVPCEConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the number of unmetered iOS or unmetered Android devices that have been
// purchased by the account.
func devicefarm_GetAccountSettings(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.GetAccountSettingsInput{}

	if resp, err := client.GetAccountSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a unique device type.
func devicefarm_GetDevice(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.GetDeviceInput{
		// Arn: *string, // Required
	}

	if len(_devicefarmArn) > 0 {
		input.Arn = aws.String(_devicefarmArn)
	}

	if resp, err := client.GetDevice(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a device instance that belongs to a private device
// fleet.
func devicefarm_GetDeviceInstance(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.GetDeviceInstanceInput{
		// Arn: *string, // Required
	}

	if len(_devicefarmArn) > 0 {
		input.Arn = aws.String(_devicefarmArn)
	}

	if resp, err := client.GetDeviceInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a device pool.
func devicefarm_GetDevicePool(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.GetDevicePoolInput{
		// Arn: *string, // Required
	}

	if len(_devicefarmArn) > 0 {
		input.Arn = aws.String(_devicefarmArn)
	}

	if resp, err := client.GetDevicePool(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about compatibility with a device pool.
func devicefarm_GetDevicePoolCompatibility(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.GetDevicePoolCompatibilityInput{
		// DevicePoolArn: *string, // Required
	}

	if len(_devicefarmDevicePoolArn) > 0 {
		input.DevicePoolArn = aws.String(_devicefarmDevicePoolArn)
	}
	if len(_devicefarmAppArn) > 0 {
		input.AppArn = aws.String(_devicefarmAppArn)
	}
	if len(_devicefarmConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _devicefarmConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_devicefarmProjectArn) > 0 {
		input.ProjectArn = aws.String(_devicefarmProjectArn)
	}
	if len(_devicefarmTest) > 0 {
		if err := assignInputField(input, "Test", _devicefarmTest); err != nil {
			log.Errorf("invalid --test: %s", err.Error())
			return
		}
	}
	if len(_devicefarmTestType) > 0 {
		if err := assignInputField(input, "TestType", _devicefarmTestType); err != nil {
			log.Errorf("invalid --test-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetDevicePoolCompatibility(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the specified instance profile.
func devicefarm_GetInstanceProfile(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.GetInstanceProfileInput{
		// Arn: *string, // Required
	}

	if len(_devicefarmArn) > 0 {
		input.Arn = aws.String(_devicefarmArn)
	}

	if resp, err := client.GetInstanceProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a job.
func devicefarm_GetJob(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.GetJobInput{
		// Arn: *string, // Required
	}

	if len(_devicefarmArn) > 0 {
		input.Arn = aws.String(_devicefarmArn)
	}

	if resp, err := client.GetJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a network profile.
func devicefarm_GetNetworkProfile(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.GetNetworkProfileInput{
		// Arn: *string, // Required
	}

	if len(_devicefarmArn) > 0 {
		input.Arn = aws.String(_devicefarmArn)
	}

	if resp, err := client.GetNetworkProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the current status and future status of all offerings purchased by an AWS
// account. The response indicates how many offerings are currently available and
// the offerings that will be available in the next period. The API returns a
// NotEligible error if the user is not permitted to invoke the operation. If you
// must be able to invoke this operation, contact aws-devicefarm-support(at)amazon.com.
func devicefarm_GetOfferingStatus(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.GetOfferingStatusInput{}

	if len(_devicefarmNextToken) > 0 {
		input.NextToken = aws.String(_devicefarmNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetOfferingStatus(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*devicefarm.GetOfferingStatusOutput
	p := devicefarm.NewGetOfferingStatusPaginator(client, input)
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

// Gets information about a project.
func devicefarm_GetProject(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.GetProjectInput{
		// Arn: *string, // Required
	}

	if len(_devicefarmArn) > 0 {
		input.Arn = aws.String(_devicefarmArn)
	}

	if resp, err := client.GetProject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a link to a currently running remote access session.
func devicefarm_GetRemoteAccessSession(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.GetRemoteAccessSessionInput{
		// Arn: *string, // Required
	}

	if len(_devicefarmArn) > 0 {
		input.Arn = aws.String(_devicefarmArn)
	}

	if resp, err := client.GetRemoteAccessSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a run.
func devicefarm_GetRun(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.GetRunInput{
		// Arn: *string, // Required
	}

	if len(_devicefarmArn) > 0 {
		input.Arn = aws.String(_devicefarmArn)
	}

	if resp, err := client.GetRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a suite.
func devicefarm_GetSuite(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.GetSuiteInput{
		// Arn: *string, // Required
	}

	if len(_devicefarmArn) > 0 {
		input.Arn = aws.String(_devicefarmArn)
	}

	if resp, err := client.GetSuite(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a test.
func devicefarm_GetTest(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.GetTestInput{
		// Arn: *string, // Required
	}

	if len(_devicefarmArn) > 0 {
		input.Arn = aws.String(_devicefarmArn)
	}

	if resp, err := client.GetTest(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a Selenium testing project.
func devicefarm_GetTestGridProject(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.GetTestGridProjectInput{
		// ProjectArn: *string, // Required
	}

	if len(_devicefarmProjectArn) > 0 {
		input.ProjectArn = aws.String(_devicefarmProjectArn)
	}

	if resp, err := client.GetTestGridProject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// A session is an instance of a browser created through a RemoteWebDriver with
// the URL from CreateTestGridUrlResult$url. You can use the following to look up sessions:
//
// - The session ARN (GetTestGridSessionRequest$sessionArn ).
//
// - The project ARN and a session ID (GetTestGridSessionRequest$projectArn and GetTestGridSessionRequest$sessionId).
func devicefarm_GetTestGridSession(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.GetTestGridSessionInput{}

	if len(_devicefarmProjectArn) > 0 {
		input.ProjectArn = aws.String(_devicefarmProjectArn)
	}
	if len(_devicefarmSessionArn) > 0 {
		input.SessionArn = aws.String(_devicefarmSessionArn)
	}
	if len(_devicefarmSessionId) > 0 {
		input.SessionId = aws.String(_devicefarmSessionId)
	}

	if resp, err := client.GetTestGridSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about an upload.
func devicefarm_GetUpload(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.GetUploadInput{
		// Arn: *string, // Required
	}

	if len(_devicefarmArn) > 0 {
		input.Arn = aws.String(_devicefarmArn)
	}

	if resp, err := client.GetUpload(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the configuration settings for your Amazon Virtual
// Private Cloud (VPC) endpoint.
func devicefarm_GetVPCEConfiguration(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.GetVPCEConfigurationInput{
		// Arn: *string, // Required
	}

	if len(_devicefarmArn) > 0 {
		input.Arn = aws.String(_devicefarmArn)
	}

	if resp, err := client.GetVPCEConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Installs an application to the device in a remote access session. For Android
// applications, the file must be in .apk format. For iOS applications, the file
// must be in .ipa format.
func devicefarm_InstallToRemoteAccessSession(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.InstallToRemoteAccessSessionInput{
		// AppArn: *string, // Required
		// RemoteAccessSessionArn: *string, // Required
	}

	if len(_devicefarmAppArn) > 0 {
		input.AppArn = aws.String(_devicefarmAppArn)
	}
	if len(_devicefarmRemoteAccessSessionArn) > 0 {
		input.RemoteAccessSessionArn = aws.String(_devicefarmRemoteAccessSessionArn)
	}

	if resp, err := client.InstallToRemoteAccessSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about artifacts.
func devicefarm_ListArtifacts(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.ListArtifactsInput{
		// Arn: *string, // Required
		// Type: types.ArtifactCategory, // Required
	}

	if len(_devicefarmArn) > 0 {
		input.Arn = aws.String(_devicefarmArn)
	}
	if len(_devicefarmType) > 0 {
		if err := assignInputField(input, "Type", _devicefarmType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_devicefarmNextToken) > 0 {
		input.NextToken = aws.String(_devicefarmNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListArtifacts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*devicefarm.ListArtifactsOutput
	p := devicefarm.NewListArtifactsPaginator(client, input)
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

// Returns information about the private device instances associated with one or
// more AWS accounts.
func devicefarm_ListDeviceInstances(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.ListDeviceInstancesInput{}

	if len(_devicefarmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _devicefarmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_devicefarmNextToken) > 0 {
		input.NextToken = aws.String(_devicefarmNextToken)
	}

	if resp, err := client.ListDeviceInstances(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about device pools.
func devicefarm_ListDevicePools(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.ListDevicePoolsInput{
		// Arn: *string, // Required
	}

	if len(_devicefarmArn) > 0 {
		input.Arn = aws.String(_devicefarmArn)
	}
	if len(_devicefarmNextToken) > 0 {
		input.NextToken = aws.String(_devicefarmNextToken)
	}
	if len(_devicefarmType) > 0 {
		if err := assignInputField(input, "Type", _devicefarmType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListDevicePools(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*devicefarm.ListDevicePoolsOutput
	p := devicefarm.NewListDevicePoolsPaginator(client, input)
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

// Gets information about unique device types.
func devicefarm_ListDevices(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.ListDevicesInput{}

	if len(_devicefarmArn) > 0 {
		input.Arn = aws.String(_devicefarmArn)
	}
	if len(_devicefarmFilters) > 0 {
		if err := assignInputField(input, "Filters", _devicefarmFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_devicefarmNextToken) > 0 {
		input.NextToken = aws.String(_devicefarmNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDevices(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*devicefarm.ListDevicesOutput
	p := devicefarm.NewListDevicesPaginator(client, input)
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

// Returns information about all the instance profiles in an AWS account.
func devicefarm_ListInstanceProfiles(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.ListInstanceProfilesInput{}

	if len(_devicefarmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _devicefarmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_devicefarmNextToken) > 0 {
		input.NextToken = aws.String(_devicefarmNextToken)
	}

	if resp, err := client.ListInstanceProfiles(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about jobs for a given test run.
func devicefarm_ListJobs(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.ListJobsInput{
		// Arn: *string, // Required
	}

	if len(_devicefarmArn) > 0 {
		input.Arn = aws.String(_devicefarmArn)
	}
	if len(_devicefarmNextToken) > 0 {
		input.NextToken = aws.String(_devicefarmNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*devicefarm.ListJobsOutput
	p := devicefarm.NewListJobsPaginator(client, input)
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

// Returns the list of available network profiles.
func devicefarm_ListNetworkProfiles(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.ListNetworkProfilesInput{
		// Arn: *string, // Required
	}

	if len(_devicefarmArn) > 0 {
		input.Arn = aws.String(_devicefarmArn)
	}
	if len(_devicefarmNextToken) > 0 {
		input.NextToken = aws.String(_devicefarmNextToken)
	}
	if len(_devicefarmType) > 0 {
		if err := assignInputField(input, "Type", _devicefarmType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListNetworkProfiles(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of offering promotions. Each offering promotion record contains
// the ID and description of the promotion. The API returns a NotEligible error if
// the caller is not permitted to invoke the operation. Contact aws-devicefarm-support(at)amazon.comif you must be
// able to invoke this operation.
func devicefarm_ListOfferingPromotions(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.ListOfferingPromotionsInput{}

	if len(_devicefarmNextToken) > 0 {
		input.NextToken = aws.String(_devicefarmNextToken)
	}

	if resp, err := client.ListOfferingPromotions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of all historical purchases, renewals, and system renewal
// transactions for an AWS account. The list is paginated and ordered by a
// descending timestamp (most recent transactions are first). The API returns a
// NotEligible error if the user is not permitted to invoke the operation. If you
// must be able to invoke this operation, contact aws-devicefarm-support(at)amazon.com.
func devicefarm_ListOfferingTransactions(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.ListOfferingTransactionsInput{}

	if len(_devicefarmNextToken) > 0 {
		input.NextToken = aws.String(_devicefarmNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListOfferingTransactions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*devicefarm.ListOfferingTransactionsOutput
	p := devicefarm.NewListOfferingTransactionsPaginator(client, input)
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

// Returns a list of products or offerings that the user can manage through the
// API. Each offering record indicates the recurring price per unit and the
// frequency for that offering. The API returns a NotEligible error if the user is
// not permitted to invoke the operation. If you must be able to invoke this
// operation, contact aws-devicefarm-support(at)amazon.com.
func devicefarm_ListOfferings(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.ListOfferingsInput{}

	if len(_devicefarmNextToken) > 0 {
		input.NextToken = aws.String(_devicefarmNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListOfferings(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*devicefarm.ListOfferingsOutput
	p := devicefarm.NewListOfferingsPaginator(client, input)
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

// Gets information about projects.
func devicefarm_ListProjects(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.ListProjectsInput{}

	if len(_devicefarmArn) > 0 {
		input.Arn = aws.String(_devicefarmArn)
	}
	if len(_devicefarmNextToken) > 0 {
		input.NextToken = aws.String(_devicefarmNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListProjects(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*devicefarm.ListProjectsOutput
	p := devicefarm.NewListProjectsPaginator(client, input)
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

// Returns a list of all currently running remote access sessions.
func devicefarm_ListRemoteAccessSessions(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.ListRemoteAccessSessionsInput{
		// Arn: *string, // Required
	}

	if len(_devicefarmArn) > 0 {
		input.Arn = aws.String(_devicefarmArn)
	}
	if len(_devicefarmNextToken) > 0 {
		input.NextToken = aws.String(_devicefarmNextToken)
	}

	if resp, err := client.ListRemoteAccessSessions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about runs, given an AWS Device Farm project ARN.
func devicefarm_ListRuns(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.ListRunsInput{
		// Arn: *string, // Required
	}

	if len(_devicefarmArn) > 0 {
		input.Arn = aws.String(_devicefarmArn)
	}
	if len(_devicefarmNextToken) > 0 {
		input.NextToken = aws.String(_devicefarmNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRuns(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*devicefarm.ListRunsOutput
	p := devicefarm.NewListRunsPaginator(client, input)
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

// Gets information about samples, given an AWS Device Farm job ARN.
func devicefarm_ListSamples(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.ListSamplesInput{
		// Arn: *string, // Required
	}

	if len(_devicefarmArn) > 0 {
		input.Arn = aws.String(_devicefarmArn)
	}
	if len(_devicefarmNextToken) > 0 {
		input.NextToken = aws.String(_devicefarmNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSamples(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*devicefarm.ListSamplesOutput
	p := devicefarm.NewListSamplesPaginator(client, input)
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

// Gets information about test suites for a given job.
func devicefarm_ListSuites(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.ListSuitesInput{
		// Arn: *string, // Required
	}

	if len(_devicefarmArn) > 0 {
		input.Arn = aws.String(_devicefarmArn)
	}
	if len(_devicefarmNextToken) > 0 {
		input.NextToken = aws.String(_devicefarmNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSuites(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*devicefarm.ListSuitesOutput
	p := devicefarm.NewListSuitesPaginator(client, input)
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

// List the tags for an AWS Device Farm resource.
func devicefarm_ListTagsForResource(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.ListTagsForResourceInput{
		// ResourceARN: *string, // Required
	}

	if len(_devicefarmResourceARN) > 0 {
		input.ResourceARN = aws.String(_devicefarmResourceARN)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a list of all Selenium testing projects in your account.
func devicefarm_ListTestGridProjects(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.ListTestGridProjectsInput{}

	if len(_devicefarmMaxResult) > 0 {
		if err := assignInputField(input, "MaxResult", _devicefarmMaxResult); err != nil {
			log.Errorf("invalid --max-result: %s", err.Error())
			return
		}
	}
	if len(_devicefarmNextToken) > 0 {
		input.NextToken = aws.String(_devicefarmNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTestGridProjects(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*devicefarm.ListTestGridProjectsOutput
	p := devicefarm.NewListTestGridProjectsPaginator(client, input)
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

// Returns a list of the actions taken in a TestGridSession.
func devicefarm_ListTestGridSessionActions(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.ListTestGridSessionActionsInput{
		// SessionArn: *string, // Required
	}

	if len(_devicefarmSessionArn) > 0 {
		input.SessionArn = aws.String(_devicefarmSessionArn)
	}
	if len(_devicefarmMaxResult) > 0 {
		if err := assignInputField(input, "MaxResult", _devicefarmMaxResult); err != nil {
			log.Errorf("invalid --max-result: %s", err.Error())
			return
		}
	}
	if len(_devicefarmNextToken) > 0 {
		input.NextToken = aws.String(_devicefarmNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTestGridSessionActions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*devicefarm.ListTestGridSessionActionsOutput
	p := devicefarm.NewListTestGridSessionActionsPaginator(client, input)
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

// Retrieves a list of artifacts created during the session.
func devicefarm_ListTestGridSessionArtifacts(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.ListTestGridSessionArtifactsInput{
		// SessionArn: *string, // Required
	}

	if len(_devicefarmSessionArn) > 0 {
		input.SessionArn = aws.String(_devicefarmSessionArn)
	}
	if len(_devicefarmMaxResult) > 0 {
		if err := assignInputField(input, "MaxResult", _devicefarmMaxResult); err != nil {
			log.Errorf("invalid --max-result: %s", err.Error())
			return
		}
	}
	if len(_devicefarmNextToken) > 0 {
		input.NextToken = aws.String(_devicefarmNextToken)
	}
	if len(_devicefarmType) > 0 {
		if err := assignInputField(input, "Type", _devicefarmType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListTestGridSessionArtifacts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*devicefarm.ListTestGridSessionArtifactsOutput
	p := devicefarm.NewListTestGridSessionArtifactsPaginator(client, input)
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

// Retrieves a list of sessions for a TestGridProject.
func devicefarm_ListTestGridSessions(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.ListTestGridSessionsInput{
		// ProjectArn: *string, // Required
	}

	if len(_devicefarmProjectArn) > 0 {
		input.ProjectArn = aws.String(_devicefarmProjectArn)
	}
	if len(_devicefarmCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _devicefarmCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_devicefarmCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _devicefarmCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_devicefarmEndTimeAfter) > 0 {
		if err := assignInputField(input, "EndTimeAfter", _devicefarmEndTimeAfter); err != nil {
			log.Errorf("invalid --end-time-after: %s", err.Error())
			return
		}
	}
	if len(_devicefarmEndTimeBefore) > 0 {
		if err := assignInputField(input, "EndTimeBefore", _devicefarmEndTimeBefore); err != nil {
			log.Errorf("invalid --end-time-before: %s", err.Error())
			return
		}
	}
	if len(_devicefarmMaxResult) > 0 {
		if err := assignInputField(input, "MaxResult", _devicefarmMaxResult); err != nil {
			log.Errorf("invalid --max-result: %s", err.Error())
			return
		}
	}
	if len(_devicefarmNextToken) > 0 {
		input.NextToken = aws.String(_devicefarmNextToken)
	}
	if len(_devicefarmStatus) > 0 {
		if err := assignInputField(input, "Status", _devicefarmStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListTestGridSessions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*devicefarm.ListTestGridSessionsOutput
	p := devicefarm.NewListTestGridSessionsPaginator(client, input)
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

// Gets information about tests in a given test suite.
func devicefarm_ListTests(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.ListTestsInput{
		// Arn: *string, // Required
	}

	if len(_devicefarmArn) > 0 {
		input.Arn = aws.String(_devicefarmArn)
	}
	if len(_devicefarmNextToken) > 0 {
		input.NextToken = aws.String(_devicefarmNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTests(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*devicefarm.ListTestsOutput
	p := devicefarm.NewListTestsPaginator(client, input)
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

// Gets information about unique problems, such as exceptions or crashes.
// Unique problems are defined as a single instance of an error across a run, job,
// or suite. For example, if a call in your application consistently raises an
// exception ( OutOfBoundsException in MyActivity.java:386 ), ListUniqueProblems
// returns a single entry instead of many individual entries for that exception.
func devicefarm_ListUniqueProblems(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.ListUniqueProblemsInput{
		// Arn: *string, // Required
	}

	if len(_devicefarmArn) > 0 {
		input.Arn = aws.String(_devicefarmArn)
	}
	if len(_devicefarmNextToken) > 0 {
		input.NextToken = aws.String(_devicefarmNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListUniqueProblems(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*devicefarm.ListUniqueProblemsOutput
	p := devicefarm.NewListUniqueProblemsPaginator(client, input)
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

// Gets information about uploads, given an AWS Device Farm project ARN.
func devicefarm_ListUploads(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.ListUploadsInput{
		// Arn: *string, // Required
	}

	if len(_devicefarmArn) > 0 {
		input.Arn = aws.String(_devicefarmArn)
	}
	if len(_devicefarmNextToken) > 0 {
		input.NextToken = aws.String(_devicefarmNextToken)
	}
	if len(_devicefarmType) > 0 {
		if err := assignInputField(input, "Type", _devicefarmType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListUploads(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*devicefarm.ListUploadsOutput
	p := devicefarm.NewListUploadsPaginator(client, input)
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

// Returns information about all Amazon Virtual Private Cloud (VPC) endpoint
// configurations in the AWS account.
func devicefarm_ListVPCEConfigurations(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.ListVPCEConfigurationsInput{}

	if len(_devicefarmMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _devicefarmMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_devicefarmNextToken) > 0 {
		input.NextToken = aws.String(_devicefarmNextToken)
	}

	if resp, err := client.ListVPCEConfigurations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Immediately purchases offerings for an AWS account. Offerings renew with the
// latest total purchased quantity for an offering, unless the renewal was
// overridden. The API returns a NotEligible error if the user is not permitted to
// invoke the operation. If you must be able to invoke this operation, contact aws-devicefarm-support(at)amazon.com.
func devicefarm_PurchaseOffering(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.PurchaseOfferingInput{
		// OfferingId: *string, // Required
		// Quantity: *int32, // Required
	}

	if len(_devicefarmOfferingId) > 0 {
		input.OfferingId = aws.String(_devicefarmOfferingId)
	}
	if len(_devicefarmQuantity) > 0 {
		if err := assignInputField(input, "Quantity", _devicefarmQuantity); err != nil {
			log.Errorf("invalid --quantity: %s", err.Error())
			return
		}
	}
	if len(_devicefarmOfferingPromotionId) > 0 {
		input.OfferingPromotionId = aws.String(_devicefarmOfferingPromotionId)
	}

	if resp, err := client.PurchaseOffering(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Explicitly sets the quantity of devices to renew for an offering, starting from
// the effectiveDate of the next period. The API returns a NotEligible error if
// the user is not permitted to invoke the operation. If you must be able to invoke
// this operation, contact aws-devicefarm-support(at)amazon.com.
func devicefarm_RenewOffering(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.RenewOfferingInput{
		// OfferingId: *string, // Required
		// Quantity: *int32, // Required
	}

	if len(_devicefarmOfferingId) > 0 {
		input.OfferingId = aws.String(_devicefarmOfferingId)
	}
	if len(_devicefarmQuantity) > 0 {
		if err := assignInputField(input, "Quantity", _devicefarmQuantity); err != nil {
			log.Errorf("invalid --quantity: %s", err.Error())
			return
		}
	}

	if resp, err := client.RenewOffering(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Schedules a run.
func devicefarm_ScheduleRun(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.ScheduleRunInput{
		// ProjectArn: *string, // Required
		// Test: *types.ScheduleRunTest, // Required
	}

	if len(_devicefarmProjectArn) > 0 {
		input.ProjectArn = aws.String(_devicefarmProjectArn)
	}
	if len(_devicefarmTest) > 0 {
		if err := assignInputField(input, "Test", _devicefarmTest); err != nil {
			log.Errorf("invalid --test: %s", err.Error())
			return
		}
	}
	if len(_devicefarmAppArn) > 0 {
		input.AppArn = aws.String(_devicefarmAppArn)
	}
	if len(_devicefarmConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _devicefarmConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_devicefarmDevicePoolArn) > 0 {
		input.DevicePoolArn = aws.String(_devicefarmDevicePoolArn)
	}
	if len(_devicefarmDeviceSelectionConfiguration) > 0 {
		if err := assignInputField(input, "DeviceSelectionConfiguration", _devicefarmDeviceSelectionConfiguration); err != nil {
			log.Errorf("invalid --device-selection-configuration: %s", err.Error())
			return
		}
	}
	if len(_devicefarmExecutionConfiguration) > 0 {
		if err := assignInputField(input, "ExecutionConfiguration", _devicefarmExecutionConfiguration); err != nil {
			log.Errorf("invalid --execution-configuration: %s", err.Error())
			return
		}
	}
	if len(_devicefarmName) > 0 {
		input.Name = aws.String(_devicefarmName)
	}

	if resp, err := client.ScheduleRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initiates a stop request for the current job. AWS Device Farm immediately stops
// the job on the device where tests have not started. You are not billed for this
// device. On the device where tests have started, setup suite and teardown suite
// tests run to completion on the device. You are billed for setup, teardown, and
// any tests that were in progress or already completed.
func devicefarm_StopJob(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.StopJobInput{
		// Arn: *string, // Required
	}

	if len(_devicefarmArn) > 0 {
		input.Arn = aws.String(_devicefarmArn)
	}

	if resp, err := client.StopJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Ends a specified remote access session.
func devicefarm_StopRemoteAccessSession(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.StopRemoteAccessSessionInput{
		// Arn: *string, // Required
	}

	if len(_devicefarmArn) > 0 {
		input.Arn = aws.String(_devicefarmArn)
	}

	if resp, err := client.StopRemoteAccessSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initiates a stop request for the current test run. AWS Device Farm immediately
// stops the run on devices where tests have not started. You are not billed for
// these devices. On devices where tests have started executing, setup suite and
// teardown suite tests run to completion on those devices. You are billed for
// setup, teardown, and any tests that were in progress or already completed.
func devicefarm_StopRun(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.StopRunInput{
		// Arn: *string, // Required
	}

	if len(_devicefarmArn) > 0 {
		input.Arn = aws.String(_devicefarmArn)
	}

	if resp, err := client.StopRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates the specified tags to a resource with the specified resourceArn . If
// existing tags on a resource are not specified in the request parameters, they
// are not changed. When a resource is deleted, the tags associated with that
// resource are also deleted.
func devicefarm_TagResource(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.TagResourceInput{
		// ResourceARN: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_devicefarmResourceARN) > 0 {
		input.ResourceARN = aws.String(_devicefarmResourceARN)
	}
	if len(_devicefarmTags) > 0 {
		if err := assignInputField(input, "Tags", _devicefarmTags); err != nil {
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

// Deletes the specified tags from a resource.
func devicefarm_UntagResource(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.UntagResourceInput{
		// ResourceARN: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_devicefarmResourceARN) > 0 {
		input.ResourceARN = aws.String(_devicefarmResourceARN)
	}
	if len(_devicefarmTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _devicefarmTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates information about a private device instance.
func devicefarm_UpdateDeviceInstance(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.UpdateDeviceInstanceInput{
		// Arn: *string, // Required
	}

	if len(_devicefarmArn) > 0 {
		input.Arn = aws.String(_devicefarmArn)
	}
	if len(_devicefarmLabels) > 0 {
		input.Labels = append([]string(nil), _devicefarmLabels...)
	}
	if len(_devicefarmProfileArn) > 0 {
		input.ProfileArn = aws.String(_devicefarmProfileArn)
	}

	if resp, err := client.UpdateDeviceInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the name, description, and rules in a device pool given the attributes
// and the pool ARN. Rule updates are all-or-nothing, meaning they can only be
// updated as a whole (or not at all).
func devicefarm_UpdateDevicePool(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.UpdateDevicePoolInput{
		// Arn: *string, // Required
	}

	if len(_devicefarmArn) > 0 {
		input.Arn = aws.String(_devicefarmArn)
	}
	if len(_devicefarmClearMaxDevices) > 0 {
		if err := assignInputField(input, "ClearMaxDevices", _devicefarmClearMaxDevices); err != nil {
			log.Errorf("invalid --clear-max-devices: %s", err.Error())
			return
		}
	}
	if len(_devicefarmDescription) > 0 {
		input.Description = aws.String(_devicefarmDescription)
	}
	if len(_devicefarmMaxDevices) > 0 {
		if err := assignInputField(input, "MaxDevices", _devicefarmMaxDevices); err != nil {
			log.Errorf("invalid --max-devices: %s", err.Error())
			return
		}
	}
	if len(_devicefarmName) > 0 {
		input.Name = aws.String(_devicefarmName)
	}
	if len(_devicefarmRules) > 0 {
		if err := assignInputField(input, "Rules", _devicefarmRules); err != nil {
			log.Errorf("invalid --rules: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDevicePool(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates information about an existing private device instance profile.
func devicefarm_UpdateInstanceProfile(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.UpdateInstanceProfileInput{
		// Arn: *string, // Required
	}

	if len(_devicefarmArn) > 0 {
		input.Arn = aws.String(_devicefarmArn)
	}
	if len(_devicefarmDescription) > 0 {
		input.Description = aws.String(_devicefarmDescription)
	}
	if len(_devicefarmExcludeAppPackagesFromCleanup) > 0 {
		input.ExcludeAppPackagesFromCleanup = append([]string(nil), _devicefarmExcludeAppPackagesFromCleanup...)
	}
	if len(_devicefarmName) > 0 {
		input.Name = aws.String(_devicefarmName)
	}
	if len(_devicefarmPackageCleanup) > 0 {
		if err := assignInputField(input, "PackageCleanup", _devicefarmPackageCleanup); err != nil {
			log.Errorf("invalid --package-cleanup: %s", err.Error())
			return
		}
	}
	if len(_devicefarmRebootAfterUse) > 0 {
		if err := assignInputField(input, "RebootAfterUse", _devicefarmRebootAfterUse); err != nil {
			log.Errorf("invalid --reboot-after-use: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateInstanceProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the network profile.
func devicefarm_UpdateNetworkProfile(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.UpdateNetworkProfileInput{
		// Arn: *string, // Required
	}

	if len(_devicefarmArn) > 0 {
		input.Arn = aws.String(_devicefarmArn)
	}
	if len(_devicefarmDescription) > 0 {
		input.Description = aws.String(_devicefarmDescription)
	}
	if len(_devicefarmDownlinkBandwidthBits) > 0 {
		if err := assignInputField(input, "DownlinkBandwidthBits", _devicefarmDownlinkBandwidthBits); err != nil {
			log.Errorf("invalid --downlink-bandwidth-bits: %s", err.Error())
			return
		}
	}
	if len(_devicefarmDownlinkDelayMs) > 0 {
		if err := assignInputField(input, "DownlinkDelayMs", _devicefarmDownlinkDelayMs); err != nil {
			log.Errorf("invalid --downlink-delay-ms: %s", err.Error())
			return
		}
	}
	if len(_devicefarmDownlinkJitterMs) > 0 {
		if err := assignInputField(input, "DownlinkJitterMs", _devicefarmDownlinkJitterMs); err != nil {
			log.Errorf("invalid --downlink-jitter-ms: %s", err.Error())
			return
		}
	}
	if len(_devicefarmDownlinkLossPercent) > 0 {
		if err := assignInputField(input, "DownlinkLossPercent", _devicefarmDownlinkLossPercent); err != nil {
			log.Errorf("invalid --downlink-loss-percent: %s", err.Error())
			return
		}
	}
	if len(_devicefarmName) > 0 {
		input.Name = aws.String(_devicefarmName)
	}
	if len(_devicefarmType) > 0 {
		if err := assignInputField(input, "Type", _devicefarmType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_devicefarmUplinkBandwidthBits) > 0 {
		if err := assignInputField(input, "UplinkBandwidthBits", _devicefarmUplinkBandwidthBits); err != nil {
			log.Errorf("invalid --uplink-bandwidth-bits: %s", err.Error())
			return
		}
	}
	if len(_devicefarmUplinkDelayMs) > 0 {
		if err := assignInputField(input, "UplinkDelayMs", _devicefarmUplinkDelayMs); err != nil {
			log.Errorf("invalid --uplink-delay-ms: %s", err.Error())
			return
		}
	}
	if len(_devicefarmUplinkJitterMs) > 0 {
		if err := assignInputField(input, "UplinkJitterMs", _devicefarmUplinkJitterMs); err != nil {
			log.Errorf("invalid --uplink-jitter-ms: %s", err.Error())
			return
		}
	}
	if len(_devicefarmUplinkLossPercent) > 0 {
		if err := assignInputField(input, "UplinkLossPercent", _devicefarmUplinkLossPercent); err != nil {
			log.Errorf("invalid --uplink-loss-percent: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateNetworkProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the specified project name, given the project ARN and a new name.
func devicefarm_UpdateProject(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.UpdateProjectInput{
		// Arn: *string, // Required
	}

	if len(_devicefarmArn) > 0 {
		input.Arn = aws.String(_devicefarmArn)
	}
	if len(_devicefarmDefaultJobTimeoutMinutes) > 0 {
		if err := assignInputField(input, "DefaultJobTimeoutMinutes", _devicefarmDefaultJobTimeoutMinutes); err != nil {
			log.Errorf("invalid --default-job-timeout-minutes: %s", err.Error())
			return
		}
	}
	if len(_devicefarmEnvironmentVariables) > 0 {
		if err := assignInputField(input, "EnvironmentVariables", _devicefarmEnvironmentVariables); err != nil {
			log.Errorf("invalid --environment-variables: %s", err.Error())
			return
		}
	}
	if len(_devicefarmExecutionRoleArn) > 0 {
		input.ExecutionRoleArn = aws.String(_devicefarmExecutionRoleArn)
	}
	if len(_devicefarmName) > 0 {
		input.Name = aws.String(_devicefarmName)
	}
	if len(_devicefarmVpcConfig) > 0 {
		if err := assignInputField(input, "VpcConfig", _devicefarmVpcConfig); err != nil {
			log.Errorf("invalid --vpc-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateProject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Change details of a project.
func devicefarm_UpdateTestGridProject(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.UpdateTestGridProjectInput{
		// ProjectArn: *string, // Required
	}

	if len(_devicefarmProjectArn) > 0 {
		input.ProjectArn = aws.String(_devicefarmProjectArn)
	}
	if len(_devicefarmDescription) > 0 {
		input.Description = aws.String(_devicefarmDescription)
	}
	if len(_devicefarmName) > 0 {
		input.Name = aws.String(_devicefarmName)
	}
	if len(_devicefarmVpcConfig) > 0 {
		if err := assignInputField(input, "VpcConfig", _devicefarmVpcConfig); err != nil {
			log.Errorf("invalid --vpc-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateTestGridProject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an uploaded test spec.
func devicefarm_UpdateUpload(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.UpdateUploadInput{
		// Arn: *string, // Required
	}

	if len(_devicefarmArn) > 0 {
		input.Arn = aws.String(_devicefarmArn)
	}
	if len(_devicefarmContentType) > 0 {
		input.ContentType = aws.String(_devicefarmContentType)
	}
	if len(_devicefarmEditContent) > 0 {
		if err := assignInputField(input, "EditContent", _devicefarmEditContent); err != nil {
			log.Errorf("invalid --edit-content: %s", err.Error())
			return
		}
	}
	if len(_devicefarmName) > 0 {
		input.Name = aws.String(_devicefarmName)
	}

	if resp, err := client.UpdateUpload(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates information about an Amazon Virtual Private Cloud (VPC) endpoint
// configuration.
func devicefarm_UpdateVPCEConfiguration(cfg aws.Config, client *devicefarm.Client) {
	input := &devicefarm.UpdateVPCEConfigurationInput{
		// Arn: *string, // Required
	}

	if len(_devicefarmArn) > 0 {
		input.Arn = aws.String(_devicefarmArn)
	}
	if len(_devicefarmServiceDnsName) > 0 {
		input.ServiceDnsName = aws.String(_devicefarmServiceDnsName)
	}
	if len(_devicefarmVpceConfigurationDescription) > 0 {
		input.VpceConfigurationDescription = aws.String(_devicefarmVpceConfigurationDescription)
	}
	if len(_devicefarmVpceConfigurationName) > 0 {
		input.VpceConfigurationName = aws.String(_devicefarmVpceConfigurationName)
	}
	if len(_devicefarmVpceServiceName) > 0 {
		input.VpceServiceName = aws.String(_devicefarmVpceServiceName)
	}

	if resp, err := client.UpdateVPCEConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_devicefarmCmd)
	_devicefarmCmd.Flags().SortFlags = false

	_devicefarmCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_devicefarmCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_devicefarmCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_devicefarmCmd.Flags().StringVarP(&_devicefarmAppArn, "app-arn", "", "", "App ARN")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmArn, "arn", "", "", "ARN")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmClearMaxDevices, "clear-max-devices", "", "", "Clear Max Devices")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmConfiguration, "configuration", "", "", "Configuration")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmContentType, "content-type", "", "", "Content Type")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmCreationTimeAfter, "creation-time-after", "", "", "Creation Time After")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmCreationTimeBefore, "creation-time-before", "", "", "Creation Time Before")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmDefaultJobTimeoutMinutes, "default-job-timeout-minutes", "", "", "Default Job Timeout Minutes")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmDescription, "description", "", "", "Description")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmDeviceArn, "device-arn", "", "", "Device ARN")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmDevicePoolArn, "device-pool-arn", "", "", "Device Pool ARN")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmDeviceSelectionConfiguration, "device-selection-configuration", "", "", "Device Selection Configuration")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmDownlinkBandwidthBits, "downlink-bandwidth-bits", "", "", "Downlink Bandwidth Bits")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmDownlinkDelayMs, "downlink-delay-ms", "", "", "Downlink Delay Ms")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmDownlinkJitterMs, "downlink-jitter-ms", "", "", "Downlink Jitter Ms")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmDownlinkLossPercent, "downlink-loss-percent", "", "", "Downlink Loss Percent")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmEditContent, "edit-content", "", "", "Edit Content")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmEndTimeAfter, "end-time-after", "", "", "End Time After")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmEndTimeBefore, "end-time-before", "", "", "End Time Before")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmEnvironmentVariables, "environment-variables", "", "", "Environment Variables")
	_devicefarmCmd.Flags().StringSliceVarP(&_devicefarmExcludeAppPackagesFromCleanup, "exclude-app-packages-from-cleanup", "", nil, "Exclude App Packages From Cleanup")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmExecutionConfiguration, "execution-configuration", "", "", "Execution Configuration")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmExecutionRoleArn, "execution-role-arn", "", "", "Execution Role ARN")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmExpiresInSeconds, "expires-in-seconds", "", "", "Expires In Seconds")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmFilters, "filters", "", "", "Filters")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmInstanceArn, "instance-arn", "", "", "Instance ARN")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmInteractionMode, "interaction-mode", "", "", "Interaction Mode")
	_devicefarmCmd.Flags().StringSliceVarP(&_devicefarmLabels, "labels", "", nil, "Labels")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmMaxDevices, "max-devices", "", "", "Max Devices")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmMaxResult, "max-result", "", "", "Max Result")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmMaxResults, "max-results", "", "", "Max Results")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmName, "name", "", "", "Name")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmNextToken, "next-token", "", "", "Next Token")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmOfferingId, "offering-id", "", "", "Offering ID")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmOfferingPromotionId, "offering-promotion-id", "", "", "Offering Promotion ID")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmPackageCleanup, "package-cleanup", "", "", "Package Cleanup")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmProfileArn, "profile-arn", "", "", "Profile ARN")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmProjectArn, "project-arn", "", "", "Project ARN")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmQuantity, "quantity", "", "", "Quantity")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmRebootAfterUse, "reboot-after-use", "", "", "Reboot After Use")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmRemoteAccessSessionArn, "remote-access-session-arn", "", "", "Remote Access Session ARN")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmResourceARN, "resource-arn", "", "", "Resource ARN")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmRules, "rules", "", "", "Rules")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmServiceDnsName, "service-dns-name", "", "", "Service DNS Name")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmSessionArn, "session-arn", "", "", "Session ARN")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmSessionId, "session-id", "", "", "Session ID")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmSkipAppResign, "skip-app-resign", "", "", "Skip App Resign")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmStatus, "status", "", "", "Status")
	_devicefarmCmd.Flags().StringSliceVarP(&_devicefarmTagKeys, "tag-keys", "", nil, "Tag Keys")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmTags, "tags", "", "", "Tags")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmTest, "test", "", "", "Test")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmTestType, "test-type", "", "", "Test Type")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmType, "type", "", "", "Type")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmUplinkBandwidthBits, "uplink-bandwidth-bits", "", "", "Uplink Bandwidth Bits")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmUplinkDelayMs, "uplink-delay-ms", "", "", "Uplink Delay Ms")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmUplinkJitterMs, "uplink-jitter-ms", "", "", "Uplink Jitter Ms")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmUplinkLossPercent, "uplink-loss-percent", "", "", "Uplink Loss Percent")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmVpcConfig, "vpc-config", "", "", "VPC Config")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmVpceConfigurationDescription, "vpce-configuration-description", "", "", "Vpce Configuration Description")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmVpceConfigurationName, "vpce-configuration-name", "", "", "Vpce Configuration Name")
	_devicefarmCmd.Flags().StringVarP(&_devicefarmVpceServiceName, "vpce-service-name", "", "", "Vpce Service Name")

	_devicefarmCmd.Flags().BoolVarP(&_devicefarmCreateDevicePool, "create-device-pool", "", false, "Create Device Pool")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmCreateInstanceProfile, "create-instance-profile", "", false, "Create Instance Profile")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmCreateNetworkProfile, "create-network-profile", "", false, "Create Network Profile")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmCreateProject, "create-project", "", false, "Create Project")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmCreateRemoteAccessSession, "create-remote-access-session", "", false, "Create Remote Access Session")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmCreateTestGridProject, "create-test-grid-project", "", false, "Create Test Grid Project")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmCreateTestGridUrl, "create-test-grid-url", "", false, "Create Test Grid URL")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmCreateUpload, "create-upload", "", false, "Create Upload")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmCreateVPCEConfiguration, "create-vpce-configuration", "", false, "Create Vpce Configuration")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmDeleteDevicePool, "delete-device-pool", "", false, "Delete Device Pool")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmDeleteInstanceProfile, "delete-instance-profile", "", false, "Delete Instance Profile")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmDeleteNetworkProfile, "delete-network-profile", "", false, "Delete Network Profile")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmDeleteProject, "delete-project", "", false, "Delete Project")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmDeleteRemoteAccessSession, "delete-remote-access-session", "", false, "Delete Remote Access Session")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmDeleteRun, "delete-run", "", false, "Delete Run")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmDeleteTestGridProject, "delete-test-grid-project", "", false, "Delete Test Grid Project")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmDeleteUpload, "delete-upload", "", false, "Delete Upload")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmDeleteVPCEConfiguration, "delete-vpce-configuration", "", false, "Delete Vpce Configuration")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmGetAccountSettings, "get-account-settings", "", false, "Get Account Settings")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmGetDevice, "get-device", "", false, "Get Device")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmGetDeviceInstance, "get-device-instance", "", false, "Get Device Instance")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmGetDevicePool, "get-device-pool", "", false, "Get Device Pool")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmGetDevicePoolCompatibility, "get-device-pool-compatibility", "", false, "Get Device Pool Compatibility")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmGetInstanceProfile, "get-instance-profile", "", false, "Get Instance Profile")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmGetJob, "get-job", "", false, "Get Job")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmGetNetworkProfile, "get-network-profile", "", false, "Get Network Profile")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmGetOfferingStatus, "get-offering-status", "", false, "Get Offering Status")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmGetProject, "get-project", "", false, "Get Project")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmGetRemoteAccessSession, "get-remote-access-session", "", false, "Get Remote Access Session")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmGetRun, "get-run", "", false, "Get Run")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmGetSuite, "get-suite", "", false, "Get Suite")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmGetTest, "get-test", "", false, "Get Test")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmGetTestGridProject, "get-test-grid-project", "", false, "Get Test Grid Project")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmGetTestGridSession, "get-test-grid-session", "", false, "Get Test Grid Session")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmGetUpload, "get-upload", "", false, "Get Upload")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmGetVPCEConfiguration, "get-vpce-configuration", "", false, "Get Vpce Configuration")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmInstallToRemoteAccessSession, "install-to-remote-access-session", "", false, "Install To Remote Access Session")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmListArtifacts, "list-artifacts", "", false, "List Artifacts")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmListDeviceInstances, "list-device-instances", "", false, "List Device Instances")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmListDevicePools, "list-device-pools", "", false, "List Device Pools")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmListDevices, "list-devices", "", false, "List Devices")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmListInstanceProfiles, "list-instance-profiles", "", false, "List Instance Profiles")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmListJobs, "list-jobs", "", false, "List Jobs")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmListNetworkProfiles, "list-network-profiles", "", false, "List Network Profiles")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmListOfferingPromotions, "list-offering-promotions", "", false, "List Offering Promotions")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmListOfferingTransactions, "list-offering-transactions", "", false, "List Offering Transactions")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmListOfferings, "list-offerings", "", false, "List Offerings")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmListProjects, "list-projects", "", false, "List Projects")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmListRemoteAccessSessions, "list-remote-access-sessions", "", false, "List Remote Access Sessions")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmListRuns, "list-runs", "", false, "List Runs")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmListSamples, "list-samples", "", false, "List Samples")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmListSuites, "list-suites", "", false, "List Suites")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmListTestGridProjects, "list-test-grid-projects", "", false, "List Test Grid Projects")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmListTestGridSessionActions, "list-test-grid-session-actions", "", false, "List Test Grid Session Actions")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmListTestGridSessionArtifacts, "list-test-grid-session-artifacts", "", false, "List Test Grid Session Artifacts")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmListTestGridSessions, "list-test-grid-sessions", "", false, "List Test Grid Sessions")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmListTests, "list-tests", "", false, "List Tests")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmListUniqueProblems, "list-unique-problems", "", false, "List Unique Problems")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmListUploads, "list-uploads", "", false, "List Uploads")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmListVPCEConfigurations, "list-vpce-configurations", "", false, "List Vpce Configurations")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmPurchaseOffering, "purchase-offering", "", false, "Purchase Offering")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmRenewOffering, "renew-offering", "", false, "Renew Offering")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmScheduleRun, "schedule-run", "", false, "Schedule Run")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmStopJob, "stop-job", "", false, "Stop Job")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmStopRemoteAccessSession, "stop-remote-access-session", "", false, "Stop Remote Access Session")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmStopRun, "stop-run", "", false, "Stop Run")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmTagResource, "tag-resource", "", false, "Tag Resource")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmUntagResource, "untag-resource", "", false, "Untag Resource")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmUpdateDeviceInstance, "update-device-instance", "", false, "Update Device Instance")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmUpdateDevicePool, "update-device-pool", "", false, "Update Device Pool")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmUpdateInstanceProfile, "update-instance-profile", "", false, "Update Instance Profile")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmUpdateNetworkProfile, "update-network-profile", "", false, "Update Network Profile")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmUpdateProject, "update-project", "", false, "Update Project")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmUpdateTestGridProject, "update-test-grid-project", "", false, "Update Test Grid Project")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmUpdateUpload, "update-upload", "", false, "Update Upload")
	_devicefarmCmd.Flags().BoolVarP(&_devicefarmUpdateVPCEConfiguration, "update-vpce-configuration", "", false, "Update Vpce Configuration")

}
