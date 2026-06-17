package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/groundstation"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// groundstationCmd represents the groundstation command
var _groundstationCmd = &cobra.Command{
	Use:   "groundstation",
	Short: "AWS groundstation CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := groundstation.NewFromConfig(cfg)
		if _groundstationCancelContact {
			groundstation_CancelContact(cfg, client)
			return
		}
		if _groundstationCreateConfig {
			groundstation_CreateConfig(cfg, client)
			return
		}
		if _groundstationCreateDataflowEndpointGroup {
			groundstation_CreateDataflowEndpointGroup(cfg, client)
			return
		}
		if _groundstationCreateDataflowEndpointGroupV2 {
			groundstation_CreateDataflowEndpointGroupV2(cfg, client)
			return
		}
		if _groundstationCreateEphemeris {
			groundstation_CreateEphemeris(cfg, client)
			return
		}
		if _groundstationCreateMissionProfile {
			groundstation_CreateMissionProfile(cfg, client)
			return
		}
		if _groundstationDeleteConfig {
			groundstation_DeleteConfig(cfg, client)
			return
		}
		if _groundstationDeleteDataflowEndpointGroup {
			groundstation_DeleteDataflowEndpointGroup(cfg, client)
			return
		}
		if _groundstationDeleteEphemeris {
			groundstation_DeleteEphemeris(cfg, client)
			return
		}
		if _groundstationDeleteMissionProfile {
			groundstation_DeleteMissionProfile(cfg, client)
			return
		}
		if _groundstationDescribeContact {
			groundstation_DescribeContact(cfg, client)
			return
		}
		if _groundstationDescribeEphemeris {
			groundstation_DescribeEphemeris(cfg, client)
			return
		}
		if _groundstationGetAgentConfiguration {
			groundstation_GetAgentConfiguration(cfg, client)
			return
		}
		if _groundstationGetAgentTaskResponseUrl {
			groundstation_GetAgentTaskResponseUrl(cfg, client)
			return
		}
		if _groundstationGetConfig {
			groundstation_GetConfig(cfg, client)
			return
		}
		if _groundstationGetDataflowEndpointGroup {
			groundstation_GetDataflowEndpointGroup(cfg, client)
			return
		}
		if _groundstationGetMinuteUsage {
			groundstation_GetMinuteUsage(cfg, client)
			return
		}
		if _groundstationGetMissionProfile {
			groundstation_GetMissionProfile(cfg, client)
			return
		}
		if _groundstationGetSatellite {
			groundstation_GetSatellite(cfg, client)
			return
		}
		if _groundstationListConfigs {
			groundstation_ListConfigs(cfg, client)
			return
		}
		if _groundstationListContacts {
			groundstation_ListContacts(cfg, client)
			return
		}
		if _groundstationListDataflowEndpointGroups {
			groundstation_ListDataflowEndpointGroups(cfg, client)
			return
		}
		if _groundstationListEphemerides {
			groundstation_ListEphemerides(cfg, client)
			return
		}
		if _groundstationListGroundStations {
			groundstation_ListGroundStations(cfg, client)
			return
		}
		if _groundstationListMissionProfiles {
			groundstation_ListMissionProfiles(cfg, client)
			return
		}
		if _groundstationListSatellites {
			groundstation_ListSatellites(cfg, client)
			return
		}
		if _groundstationListTagsForResource {
			groundstation_ListTagsForResource(cfg, client)
			return
		}
		if _groundstationRegisterAgent {
			groundstation_RegisterAgent(cfg, client)
			return
		}
		if _groundstationReserveContact {
			groundstation_ReserveContact(cfg, client)
			return
		}
		if _groundstationTagResource {
			groundstation_TagResource(cfg, client)
			return
		}
		if _groundstationUntagResource {
			groundstation_UntagResource(cfg, client)
			return
		}
		if _groundstationUpdateAgentStatus {
			groundstation_UpdateAgentStatus(cfg, client)
			return
		}
		if _groundstationUpdateConfig {
			groundstation_UpdateConfig(cfg, client)
			return
		}
		if _groundstationUpdateEphemeris {
			groundstation_UpdateEphemeris(cfg, client)
			return
		}
		if _groundstationUpdateMissionProfile {
			groundstation_UpdateMissionProfile(cfg, client)
			return
		}

	},
}

var (
	_groundstationCancelContact                 bool
	_groundstationCreateConfig                  bool
	_groundstationCreateDataflowEndpointGroup   bool
	_groundstationCreateDataflowEndpointGroupV2 bool
	_groundstationCreateEphemeris               bool
	_groundstationCreateMissionProfile          bool
	_groundstationDeleteConfig                  bool
	_groundstationDeleteDataflowEndpointGroup   bool
	_groundstationDeleteEphemeris               bool
	_groundstationDeleteMissionProfile          bool
	_groundstationDescribeContact               bool
	_groundstationDescribeEphemeris             bool
	_groundstationGetAgentConfiguration         bool
	_groundstationGetAgentTaskResponseUrl       bool
	_groundstationGetConfig                     bool
	_groundstationGetDataflowEndpointGroup      bool
	_groundstationGetMinuteUsage                bool
	_groundstationGetMissionProfile             bool
	_groundstationGetSatellite                  bool
	_groundstationListConfigs                   bool
	_groundstationListContacts                  bool
	_groundstationListDataflowEndpointGroups    bool
	_groundstationListEphemerides               bool
	_groundstationListGroundStations            bool
	_groundstationListMissionProfiles           bool
	_groundstationListSatellites                bool
	_groundstationListTagsForResource           bool
	_groundstationRegisterAgent                 bool
	_groundstationReserveContact                bool
	_groundstationTagResource                   bool
	_groundstationUntagResource                 bool
	_groundstationUpdateAgentStatus             bool
	_groundstationUpdateConfig                  bool
	_groundstationUpdateEphemeris               bool
	_groundstationUpdateMissionProfile          bool

	_groundstationAgentDetails                        string
	_groundstationAgentId                             string
	_groundstationAggregateStatus                     string
	_groundstationComponentStatuses                   string
	_groundstationConfigData                          string
	_groundstationConfigId                            string
	_groundstationConfigType                          string
	_groundstationContactId                           string
	_groundstationContactPostPassDurationSeconds      string
	_groundstationContactPrePassDurationSeconds       string
	_groundstationDataflowEdges                       string
	_groundstationDataflowEndpointGroupId             string
	_groundstationDiscoveryData                       string
	_groundstationEnabled                             string
	_groundstationEndTime                             string
	_groundstationEndpointDetails                     string
	_groundstationEndpoints                           string
	_groundstationEphemeris                           string
	_groundstationEphemerisId                         string
	_groundstationEphemerisType                       string
	_groundstationExpirationTime                      string
	_groundstationGroundStation                       string
	_groundstationKmsKeyArn                           string
	_groundstationMaxResults                          string
	_groundstationMinimumViableContactDurationSeconds string
	_groundstationMissionProfileArn                   string
	_groundstationMissionProfileId                    string
	_groundstationMonth                               string
	_groundstationName                                string
	_groundstationNextToken                           string
	_groundstationPriority                            string
	_groundstationResourceArn                         string
	_groundstationSatelliteArn                        string
	_groundstationSatelliteId                         string
	_groundstationStartTime                           string
	_groundstationStatusList                          string
	_groundstationStreamsKmsKey                       string
	_groundstationStreamsKmsRole                      string
	_groundstationTagKeys                             []string
	_groundstationTags                                string
	_groundstationTaskId                              string
	_groundstationTelemetrySinkConfigArn              string
	_groundstationTrackingConfigArn                   string
	_groundstationTrackingOverrides                   string
	_groundstationYear                                string
)

// Cancels or stops a contact with a specified contact ID based on its position in
// the [contact lifecycle].
//
// For contacts that:
//
// - Have yet to start, the contact will be cancelled.
//
// - Have started but have yet to finish, the contact will be stopped.
//
// [contact lifecycle]: https://docs.aws.amazon.com/ground-station/latest/ug/contacts.lifecycle.html
func groundstation_CancelContact(cfg aws.Config, client *groundstation.Client) {
	input := &groundstation.CancelContactInput{
		// ContactId: *string, // Required
	}

	if len(_groundstationContactId) > 0 {
		input.ContactId = aws.String(_groundstationContactId)
	}

	if resp, err := client.CancelContact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a Config with the specified configData parameters.
// Only one type of configData can be specified.
func groundstation_CreateConfig(cfg aws.Config, client *groundstation.Client) {
	input := &groundstation.CreateConfigInput{
		// ConfigData: types.ConfigTypeData, // Required
		// Name: *string, // Required
	}

	if len(_groundstationConfigData) > 0 {
		if err := assignInputField(input, "ConfigData", _groundstationConfigData); err != nil {
			log.Errorf("invalid --config-data: %s", err.Error())
			return
		}
	}
	if len(_groundstationName) > 0 {
		input.Name = aws.String(_groundstationName)
	}
	if len(_groundstationTags) > 0 {
		if err := assignInputField(input, "Tags", _groundstationTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a DataflowEndpoint group containing the specified list of
// DataflowEndpoint objects.
//
// The name field in each endpoint is used in your mission profile
// DataflowEndpointConfig to specify which endpoints to use during a contact.
//
// When a contact uses multiple DataflowEndpointConfig objects, each  Config must
// match a DataflowEndpoint in the same group.
func groundstation_CreateDataflowEndpointGroup(cfg aws.Config, client *groundstation.Client) {
	input := &groundstation.CreateDataflowEndpointGroupInput{
		// EndpointDetails: []types.EndpointDetails, // Required
	}

	if len(_groundstationEndpointDetails) > 0 {
		if err := assignInputField(input, "EndpointDetails", _groundstationEndpointDetails); err != nil {
			log.Errorf("invalid --endpoint-details: %s", err.Error())
			return
		}
	}
	if len(_groundstationContactPostPassDurationSeconds) > 0 {
		if err := assignInputField(input, "ContactPostPassDurationSeconds", _groundstationContactPostPassDurationSeconds); err != nil {
			log.Errorf("invalid --contact-post-pass-duration-seconds: %s", err.Error())
			return
		}
	}
	if len(_groundstationContactPrePassDurationSeconds) > 0 {
		if err := assignInputField(input, "ContactPrePassDurationSeconds", _groundstationContactPrePassDurationSeconds); err != nil {
			log.Errorf("invalid --contact-pre-pass-duration-seconds: %s", err.Error())
			return
		}
	}
	if len(_groundstationTags) > 0 {
		if err := assignInputField(input, "Tags", _groundstationTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDataflowEndpointGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a DataflowEndpoint group containing the specified list of Ground
// Station Agent based endpoints.
//
// The name field in each endpoint is used in your mission profile
// DataflowEndpointConfig to specify which endpoints to use during a contact.
//
// When a contact uses multiple DataflowEndpointConfig objects, each  Config must
// match a DataflowEndpoint in the same group.
func groundstation_CreateDataflowEndpointGroupV2(cfg aws.Config, client *groundstation.Client) {
	input := &groundstation.CreateDataflowEndpointGroupV2Input{
		// Endpoints: []types.CreateEndpointDetails, // Required
	}

	if len(_groundstationEndpoints) > 0 {
		if err := assignInputField(input, "Endpoints", _groundstationEndpoints); err != nil {
			log.Errorf("invalid --endpoints: %s", err.Error())
			return
		}
	}
	if len(_groundstationContactPostPassDurationSeconds) > 0 {
		if err := assignInputField(input, "ContactPostPassDurationSeconds", _groundstationContactPostPassDurationSeconds); err != nil {
			log.Errorf("invalid --contact-post-pass-duration-seconds: %s", err.Error())
			return
		}
	}
	if len(_groundstationContactPrePassDurationSeconds) > 0 {
		if err := assignInputField(input, "ContactPrePassDurationSeconds", _groundstationContactPrePassDurationSeconds); err != nil {
			log.Errorf("invalid --contact-pre-pass-duration-seconds: %s", err.Error())
			return
		}
	}
	if len(_groundstationTags) > 0 {
		if err := assignInputField(input, "Tags", _groundstationTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDataflowEndpointGroupV2(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create an ephemeris with your specified EphemerisData.
func groundstation_CreateEphemeris(cfg aws.Config, client *groundstation.Client) {
	input := &groundstation.CreateEphemerisInput{
		// Name: *string, // Required
	}

	if len(_groundstationName) > 0 {
		input.Name = aws.String(_groundstationName)
	}
	if len(_groundstationEnabled) > 0 {
		if err := assignInputField(input, "Enabled", _groundstationEnabled); err != nil {
			log.Errorf("invalid --enabled: %s", err.Error())
			return
		}
	}
	if len(_groundstationEphemeris) > 0 {
		if err := assignInputField(input, "Ephemeris", _groundstationEphemeris); err != nil {
			log.Errorf("invalid --ephemeris: %s", err.Error())
			return
		}
	}
	if len(_groundstationExpirationTime) > 0 {
		if err := assignInputField(input, "ExpirationTime", _groundstationExpirationTime); err != nil {
			log.Errorf("invalid --expiration-time: %s", err.Error())
			return
		}
	}
	if len(_groundstationKmsKeyArn) > 0 {
		input.KmsKeyArn = aws.String(_groundstationKmsKeyArn)
	}
	if len(_groundstationPriority) > 0 {
		if err := assignInputField(input, "Priority", _groundstationPriority); err != nil {
			log.Errorf("invalid --priority: %s", err.Error())
			return
		}
	}
	if len(_groundstationSatelliteId) > 0 {
		input.SatelliteId = aws.String(_groundstationSatelliteId)
	}
	if len(_groundstationTags) > 0 {
		if err := assignInputField(input, "Tags", _groundstationTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateEphemeris(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a mission profile.
// dataflowEdges is a list of lists of strings. Each lower level list of strings
// has two elements: a from ARN and a to ARN.
func groundstation_CreateMissionProfile(cfg aws.Config, client *groundstation.Client) {
	input := &groundstation.CreateMissionProfileInput{
		// DataflowEdges: [][]string, // Required
		// MinimumViableContactDurationSeconds: *int32, // Required
		// Name: *string, // Required
		// TrackingConfigArn: *string, // Required
	}

	if len(_groundstationDataflowEdges) > 0 {
		if err := assignInputField(input, "DataflowEdges", _groundstationDataflowEdges); err != nil {
			log.Errorf("invalid --dataflow-edges: %s", err.Error())
			return
		}
	}
	if len(_groundstationMinimumViableContactDurationSeconds) > 0 {
		if err := assignInputField(input, "MinimumViableContactDurationSeconds", _groundstationMinimumViableContactDurationSeconds); err != nil {
			log.Errorf("invalid --minimum-viable-contact-duration-seconds: %s", err.Error())
			return
		}
	}
	if len(_groundstationName) > 0 {
		input.Name = aws.String(_groundstationName)
	}
	if len(_groundstationTrackingConfigArn) > 0 {
		input.TrackingConfigArn = aws.String(_groundstationTrackingConfigArn)
	}
	if len(_groundstationContactPostPassDurationSeconds) > 0 {
		if err := assignInputField(input, "ContactPostPassDurationSeconds", _groundstationContactPostPassDurationSeconds); err != nil {
			log.Errorf("invalid --contact-post-pass-duration-seconds: %s", err.Error())
			return
		}
	}
	if len(_groundstationContactPrePassDurationSeconds) > 0 {
		if err := assignInputField(input, "ContactPrePassDurationSeconds", _groundstationContactPrePassDurationSeconds); err != nil {
			log.Errorf("invalid --contact-pre-pass-duration-seconds: %s", err.Error())
			return
		}
	}
	if len(_groundstationStreamsKmsKey) > 0 {
		if err := assignInputField(input, "StreamsKmsKey", _groundstationStreamsKmsKey); err != nil {
			log.Errorf("invalid --streams-kms-key: %s", err.Error())
			return
		}
	}
	if len(_groundstationStreamsKmsRole) > 0 {
		input.StreamsKmsRole = aws.String(_groundstationStreamsKmsRole)
	}
	if len(_groundstationTags) > 0 {
		if err := assignInputField(input, "Tags", _groundstationTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_groundstationTelemetrySinkConfigArn) > 0 {
		input.TelemetrySinkConfigArn = aws.String(_groundstationTelemetrySinkConfigArn)
	}

	if resp, err := client.CreateMissionProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Config .
func groundstation_DeleteConfig(cfg aws.Config, client *groundstation.Client) {
	input := &groundstation.DeleteConfigInput{
		// ConfigId: *string, // Required
		// ConfigType: types.ConfigCapabilityType, // Required
	}

	if len(_groundstationConfigId) > 0 {
		input.ConfigId = aws.String(_groundstationConfigId)
	}
	if len(_groundstationConfigType) > 0 {
		if err := assignInputField(input, "ConfigType", _groundstationConfigType); err != nil {
			log.Errorf("invalid --config-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a dataflow endpoint group.
func groundstation_DeleteDataflowEndpointGroup(cfg aws.Config, client *groundstation.Client) {
	input := &groundstation.DeleteDataflowEndpointGroupInput{
		// DataflowEndpointGroupId: *string, // Required
	}

	if len(_groundstationDataflowEndpointGroupId) > 0 {
		input.DataflowEndpointGroupId = aws.String(_groundstationDataflowEndpointGroupId)
	}

	if resp, err := client.DeleteDataflowEndpointGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete an ephemeris.
func groundstation_DeleteEphemeris(cfg aws.Config, client *groundstation.Client) {
	input := &groundstation.DeleteEphemerisInput{
		// EphemerisId: *string, // Required
	}

	if len(_groundstationEphemerisId) > 0 {
		input.EphemerisId = aws.String(_groundstationEphemerisId)
	}

	if resp, err := client.DeleteEphemeris(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a mission profile.
func groundstation_DeleteMissionProfile(cfg aws.Config, client *groundstation.Client) {
	input := &groundstation.DeleteMissionProfileInput{
		// MissionProfileId: *string, // Required
	}

	if len(_groundstationMissionProfileId) > 0 {
		input.MissionProfileId = aws.String(_groundstationMissionProfileId)
	}

	if resp, err := client.DeleteMissionProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes an existing contact.
func groundstation_DescribeContact(cfg aws.Config, client *groundstation.Client) {
	input := &groundstation.DescribeContactInput{
		// ContactId: *string, // Required
	}

	if len(_groundstationContactId) > 0 {
		input.ContactId = aws.String(_groundstationContactId)
	}

	if resp, err := client.DescribeContact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieve information about an existing ephemeris.
func groundstation_DescribeEphemeris(cfg aws.Config, client *groundstation.Client) {
	input := &groundstation.DescribeEphemerisInput{
		// EphemerisId: *string, // Required
	}

	if len(_groundstationEphemerisId) > 0 {
		input.EphemerisId = aws.String(_groundstationEphemerisId)
	}

	if resp, err := client.DescribeEphemeris(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// For use by AWS Ground Station Agent and shouldn't be called directly.
// Gets the latest configuration information for a registered agent.
func groundstation_GetAgentConfiguration(cfg aws.Config, client *groundstation.Client) {
	input := &groundstation.GetAgentConfigurationInput{
		// AgentId: *string, // Required
	}

	if len(_groundstationAgentId) > 0 {
		input.AgentId = aws.String(_groundstationAgentId)
	}

	if resp, err := client.GetAgentConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// For use by AWS Ground Station Agent and shouldn't be called directly.
// Gets a presigned URL for uploading agent task response logs.
func groundstation_GetAgentTaskResponseUrl(cfg aws.Config, client *groundstation.Client) {
	input := &groundstation.GetAgentTaskResponseUrlInput{
		// AgentId: *string, // Required
		// TaskId: *string, // Required
	}

	if len(_groundstationAgentId) > 0 {
		input.AgentId = aws.String(_groundstationAgentId)
	}
	if len(_groundstationTaskId) > 0 {
		input.TaskId = aws.String(_groundstationTaskId)
	}

	if resp, err := client.GetAgentTaskResponseUrl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns Config information.
// Only one Config response can be returned.
func groundstation_GetConfig(cfg aws.Config, client *groundstation.Client) {
	input := &groundstation.GetConfigInput{
		// ConfigId: *string, // Required
		// ConfigType: types.ConfigCapabilityType, // Required
	}

	if len(_groundstationConfigId) > 0 {
		input.ConfigId = aws.String(_groundstationConfigId)
	}
	if len(_groundstationConfigType) > 0 {
		if err := assignInputField(input, "ConfigType", _groundstationConfigType); err != nil {
			log.Errorf("invalid --config-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the dataflow endpoint group.
func groundstation_GetDataflowEndpointGroup(cfg aws.Config, client *groundstation.Client) {
	input := &groundstation.GetDataflowEndpointGroupInput{
		// DataflowEndpointGroupId: *string, // Required
	}

	if len(_groundstationDataflowEndpointGroupId) > 0 {
		input.DataflowEndpointGroupId = aws.String(_groundstationDataflowEndpointGroupId)
	}

	if resp, err := client.GetDataflowEndpointGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the number of reserved minutes used by account.
func groundstation_GetMinuteUsage(cfg aws.Config, client *groundstation.Client) {
	input := &groundstation.GetMinuteUsageInput{
		// Month: *int32, // Required
		// Year: *int32, // Required
	}

	if len(_groundstationMonth) > 0 {
		if err := assignInputField(input, "Month", _groundstationMonth); err != nil {
			log.Errorf("invalid --month: %s", err.Error())
			return
		}
	}
	if len(_groundstationYear) > 0 {
		if err := assignInputField(input, "Year", _groundstationYear); err != nil {
			log.Errorf("invalid --year: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetMinuteUsage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a mission profile.
func groundstation_GetMissionProfile(cfg aws.Config, client *groundstation.Client) {
	input := &groundstation.GetMissionProfileInput{
		// MissionProfileId: *string, // Required
	}

	if len(_groundstationMissionProfileId) > 0 {
		input.MissionProfileId = aws.String(_groundstationMissionProfileId)
	}

	if resp, err := client.GetMissionProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a satellite.
func groundstation_GetSatellite(cfg aws.Config, client *groundstation.Client) {
	input := &groundstation.GetSatelliteInput{
		// SatelliteId: *string, // Required
	}

	if len(_groundstationSatelliteId) > 0 {
		input.SatelliteId = aws.String(_groundstationSatelliteId)
	}

	if resp, err := client.GetSatellite(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of Config objects.
func groundstation_ListConfigs(cfg aws.Config, client *groundstation.Client) {
	input := &groundstation.ListConfigsInput{}

	if len(_groundstationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _groundstationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_groundstationNextToken) > 0 {
		input.NextToken = aws.String(_groundstationNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListConfigs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*groundstation.ListConfigsOutput
	p := groundstation.NewListConfigsPaginator(client, input)
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

// Returns a list of contacts.
// If statusList contains AVAILABLE, the request must include  groundStation ,
// missionprofileArn , and satelliteArn .
func groundstation_ListContacts(cfg aws.Config, client *groundstation.Client) {
	input := &groundstation.ListContactsInput{
		// EndTime: *time.Time, // Required
		// StartTime: *time.Time, // Required
		// StatusList: []types.ContactStatus, // Required
	}

	if len(_groundstationEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _groundstationEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_groundstationStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _groundstationStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_groundstationStatusList) > 0 {
		if err := assignInputField(input, "StatusList", _groundstationStatusList); err != nil {
			log.Errorf("invalid --status-list: %s", err.Error())
			return
		}
	}
	if len(_groundstationEphemeris) > 0 {
		if err := assignInputField(input, "Ephemeris", _groundstationEphemeris); err != nil {
			log.Errorf("invalid --ephemeris: %s", err.Error())
			return
		}
	}
	if len(_groundstationGroundStation) > 0 {
		input.GroundStation = aws.String(_groundstationGroundStation)
	}
	if len(_groundstationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _groundstationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_groundstationMissionProfileArn) > 0 {
		input.MissionProfileArn = aws.String(_groundstationMissionProfileArn)
	}
	if len(_groundstationNextToken) > 0 {
		input.NextToken = aws.String(_groundstationNextToken)
	}
	if len(_groundstationSatelliteArn) > 0 {
		input.SatelliteArn = aws.String(_groundstationSatelliteArn)
	}

	if disablePaginator() {
		if resp, err := client.ListContacts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*groundstation.ListContactsOutput
	p := groundstation.NewListContactsPaginator(client, input)
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

// Returns a list of DataflowEndpoint groups.
func groundstation_ListDataflowEndpointGroups(cfg aws.Config, client *groundstation.Client) {
	input := &groundstation.ListDataflowEndpointGroupsInput{}

	if len(_groundstationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _groundstationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_groundstationNextToken) > 0 {
		input.NextToken = aws.String(_groundstationNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDataflowEndpointGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*groundstation.ListDataflowEndpointGroupsOutput
	p := groundstation.NewListDataflowEndpointGroupsPaginator(client, input)
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

// List your existing ephemerides.
func groundstation_ListEphemerides(cfg aws.Config, client *groundstation.Client) {
	input := &groundstation.ListEphemeridesInput{
		// EndTime: *time.Time, // Required
		// StartTime: *time.Time, // Required
	}

	if len(_groundstationEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _groundstationEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_groundstationStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _groundstationStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_groundstationEphemerisType) > 0 {
		if err := assignInputField(input, "EphemerisType", _groundstationEphemerisType); err != nil {
			log.Errorf("invalid --ephemeris-type: %s", err.Error())
			return
		}
	}
	if len(_groundstationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _groundstationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_groundstationNextToken) > 0 {
		input.NextToken = aws.String(_groundstationNextToken)
	}
	if len(_groundstationSatelliteId) > 0 {
		input.SatelliteId = aws.String(_groundstationSatelliteId)
	}
	if len(_groundstationStatusList) > 0 {
		if err := assignInputField(input, "StatusList", _groundstationStatusList); err != nil {
			log.Errorf("invalid --status-list: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListEphemerides(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*groundstation.ListEphemeridesOutput
	p := groundstation.NewListEphemeridesPaginator(client, input)
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

// Returns a list of ground stations.
func groundstation_ListGroundStations(cfg aws.Config, client *groundstation.Client) {
	input := &groundstation.ListGroundStationsInput{}

	if len(_groundstationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _groundstationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_groundstationNextToken) > 0 {
		input.NextToken = aws.String(_groundstationNextToken)
	}
	if len(_groundstationSatelliteId) > 0 {
		input.SatelliteId = aws.String(_groundstationSatelliteId)
	}

	if disablePaginator() {
		if resp, err := client.ListGroundStations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*groundstation.ListGroundStationsOutput
	p := groundstation.NewListGroundStationsPaginator(client, input)
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

// Returns a list of mission profiles.
func groundstation_ListMissionProfiles(cfg aws.Config, client *groundstation.Client) {
	input := &groundstation.ListMissionProfilesInput{}

	if len(_groundstationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _groundstationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_groundstationNextToken) > 0 {
		input.NextToken = aws.String(_groundstationNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListMissionProfiles(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*groundstation.ListMissionProfilesOutput
	p := groundstation.NewListMissionProfilesPaginator(client, input)
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

// Returns a list of satellites.
func groundstation_ListSatellites(cfg aws.Config, client *groundstation.Client) {
	input := &groundstation.ListSatellitesInput{}

	if len(_groundstationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _groundstationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_groundstationNextToken) > 0 {
		input.NextToken = aws.String(_groundstationNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSatellites(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*groundstation.ListSatellitesOutput
	p := groundstation.NewListSatellitesPaginator(client, input)
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

// Returns a list of tags for a specified resource.
func groundstation_ListTagsForResource(cfg aws.Config, client *groundstation.Client) {
	input := &groundstation.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_groundstationResourceArn) > 0 {
		input.ResourceArn = aws.String(_groundstationResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// For use by AWS Ground Station Agent and shouldn't be called directly.
// Registers a new agent with AWS Ground Station.
func groundstation_RegisterAgent(cfg aws.Config, client *groundstation.Client) {
	input := &groundstation.RegisterAgentInput{
		// AgentDetails: *types.AgentDetails, // Required
		// DiscoveryData: *types.DiscoveryData, // Required
	}

	if len(_groundstationAgentDetails) > 0 {
		if err := assignInputField(input, "AgentDetails", _groundstationAgentDetails); err != nil {
			log.Errorf("invalid --agent-details: %s", err.Error())
			return
		}
	}
	if len(_groundstationDiscoveryData) > 0 {
		if err := assignInputField(input, "DiscoveryData", _groundstationDiscoveryData); err != nil {
			log.Errorf("invalid --discovery-data: %s", err.Error())
			return
		}
	}
	if len(_groundstationTags) > 0 {
		if err := assignInputField(input, "Tags", _groundstationTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.RegisterAgent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Reserves a contact using specified parameters.
func groundstation_ReserveContact(cfg aws.Config, client *groundstation.Client) {
	input := &groundstation.ReserveContactInput{
		// EndTime: *time.Time, // Required
		// GroundStation: *string, // Required
		// MissionProfileArn: *string, // Required
		// StartTime: *time.Time, // Required
	}

	if len(_groundstationEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _groundstationEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_groundstationGroundStation) > 0 {
		input.GroundStation = aws.String(_groundstationGroundStation)
	}
	if len(_groundstationMissionProfileArn) > 0 {
		input.MissionProfileArn = aws.String(_groundstationMissionProfileArn)
	}
	if len(_groundstationStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _groundstationStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_groundstationSatelliteArn) > 0 {
		input.SatelliteArn = aws.String(_groundstationSatelliteArn)
	}
	if len(_groundstationTags) > 0 {
		if err := assignInputField(input, "Tags", _groundstationTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_groundstationTrackingOverrides) > 0 {
		if err := assignInputField(input, "TrackingOverrides", _groundstationTrackingOverrides); err != nil {
			log.Errorf("invalid --tracking-overrides: %s", err.Error())
			return
		}
	}

	if resp, err := client.ReserveContact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assigns a tag to a resource.
func groundstation_TagResource(cfg aws.Config, client *groundstation.Client) {
	input := &groundstation.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_groundstationResourceArn) > 0 {
		input.ResourceArn = aws.String(_groundstationResourceArn)
	}
	if len(_groundstationTags) > 0 {
		if err := assignInputField(input, "Tags", _groundstationTags); err != nil {
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

// Deassigns a resource tag.
func groundstation_UntagResource(cfg aws.Config, client *groundstation.Client) {
	input := &groundstation.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_groundstationResourceArn) > 0 {
		input.ResourceArn = aws.String(_groundstationResourceArn)
	}
	if len(_groundstationTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _groundstationTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// For use by AWS Ground Station Agent and shouldn't be called directly.
// Update the status of the agent.
func groundstation_UpdateAgentStatus(cfg aws.Config, client *groundstation.Client) {
	input := &groundstation.UpdateAgentStatusInput{
		// AgentId: *string, // Required
		// AggregateStatus: *types.AggregateStatus, // Required
		// ComponentStatuses: []types.ComponentStatusData, // Required
		// TaskId: *string, // Required
	}

	if len(_groundstationAgentId) > 0 {
		input.AgentId = aws.String(_groundstationAgentId)
	}
	if len(_groundstationAggregateStatus) > 0 {
		if err := assignInputField(input, "AggregateStatus", _groundstationAggregateStatus); err != nil {
			log.Errorf("invalid --aggregate-status: %s", err.Error())
			return
		}
	}
	if len(_groundstationComponentStatuses) > 0 {
		if err := assignInputField(input, "ComponentStatuses", _groundstationComponentStatuses); err != nil {
			log.Errorf("invalid --component-statuses: %s", err.Error())
			return
		}
	}
	if len(_groundstationTaskId) > 0 {
		input.TaskId = aws.String(_groundstationTaskId)
	}

	if resp, err := client.UpdateAgentStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the Config used when scheduling contacts.
// Updating a Config will not update the execution parameters for existing future
// contacts scheduled with this Config .
func groundstation_UpdateConfig(cfg aws.Config, client *groundstation.Client) {
	input := &groundstation.UpdateConfigInput{
		// ConfigData: types.ConfigTypeData, // Required
		// ConfigId: *string, // Required
		// ConfigType: types.ConfigCapabilityType, // Required
		// Name: *string, // Required
	}

	if len(_groundstationConfigData) > 0 {
		if err := assignInputField(input, "ConfigData", _groundstationConfigData); err != nil {
			log.Errorf("invalid --config-data: %s", err.Error())
			return
		}
	}
	if len(_groundstationConfigId) > 0 {
		input.ConfigId = aws.String(_groundstationConfigId)
	}
	if len(_groundstationConfigType) > 0 {
		if err := assignInputField(input, "ConfigType", _groundstationConfigType); err != nil {
			log.Errorf("invalid --config-type: %s", err.Error())
			return
		}
	}
	if len(_groundstationName) > 0 {
		input.Name = aws.String(_groundstationName)
	}

	if resp, err := client.UpdateConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update an existing ephemeris.
func groundstation_UpdateEphemeris(cfg aws.Config, client *groundstation.Client) {
	input := &groundstation.UpdateEphemerisInput{
		// Enabled: *bool, // Required
		// EphemerisId: *string, // Required
	}

	if len(_groundstationEnabled) > 0 {
		if err := assignInputField(input, "Enabled", _groundstationEnabled); err != nil {
			log.Errorf("invalid --enabled: %s", err.Error())
			return
		}
	}
	if len(_groundstationEphemerisId) > 0 {
		input.EphemerisId = aws.String(_groundstationEphemerisId)
	}
	if len(_groundstationName) > 0 {
		input.Name = aws.String(_groundstationName)
	}
	if len(_groundstationPriority) > 0 {
		if err := assignInputField(input, "Priority", _groundstationPriority); err != nil {
			log.Errorf("invalid --priority: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateEphemeris(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a mission profile.
// Updating a mission profile will not update the execution parameters for
// existing future contacts.
func groundstation_UpdateMissionProfile(cfg aws.Config, client *groundstation.Client) {
	input := &groundstation.UpdateMissionProfileInput{
		// MissionProfileId: *string, // Required
	}

	if len(_groundstationMissionProfileId) > 0 {
		input.MissionProfileId = aws.String(_groundstationMissionProfileId)
	}
	if len(_groundstationContactPostPassDurationSeconds) > 0 {
		if err := assignInputField(input, "ContactPostPassDurationSeconds", _groundstationContactPostPassDurationSeconds); err != nil {
			log.Errorf("invalid --contact-post-pass-duration-seconds: %s", err.Error())
			return
		}
	}
	if len(_groundstationContactPrePassDurationSeconds) > 0 {
		if err := assignInputField(input, "ContactPrePassDurationSeconds", _groundstationContactPrePassDurationSeconds); err != nil {
			log.Errorf("invalid --contact-pre-pass-duration-seconds: %s", err.Error())
			return
		}
	}
	if len(_groundstationDataflowEdges) > 0 {
		if err := assignInputField(input, "DataflowEdges", _groundstationDataflowEdges); err != nil {
			log.Errorf("invalid --dataflow-edges: %s", err.Error())
			return
		}
	}
	if len(_groundstationMinimumViableContactDurationSeconds) > 0 {
		if err := assignInputField(input, "MinimumViableContactDurationSeconds", _groundstationMinimumViableContactDurationSeconds); err != nil {
			log.Errorf("invalid --minimum-viable-contact-duration-seconds: %s", err.Error())
			return
		}
	}
	if len(_groundstationName) > 0 {
		input.Name = aws.String(_groundstationName)
	}
	if len(_groundstationStreamsKmsKey) > 0 {
		if err := assignInputField(input, "StreamsKmsKey", _groundstationStreamsKmsKey); err != nil {
			log.Errorf("invalid --streams-kms-key: %s", err.Error())
			return
		}
	}
	if len(_groundstationStreamsKmsRole) > 0 {
		input.StreamsKmsRole = aws.String(_groundstationStreamsKmsRole)
	}
	if len(_groundstationTelemetrySinkConfigArn) > 0 {
		input.TelemetrySinkConfigArn = aws.String(_groundstationTelemetrySinkConfigArn)
	}
	if len(_groundstationTrackingConfigArn) > 0 {
		input.TrackingConfigArn = aws.String(_groundstationTrackingConfigArn)
	}

	if resp, err := client.UpdateMissionProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_groundstationCmd)
	_groundstationCmd.Flags().SortFlags = false

	_groundstationCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_groundstationCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_groundstationCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_groundstationCmd.Flags().StringVarP(&_groundstationAgentDetails, "agent-details", "", "", "Agent Details")
	_groundstationCmd.Flags().StringVarP(&_groundstationAgentId, "agent-id", "", "", "Agent ID")
	_groundstationCmd.Flags().StringVarP(&_groundstationAggregateStatus, "aggregate-status", "", "", "Aggregate Status")
	_groundstationCmd.Flags().StringVarP(&_groundstationComponentStatuses, "component-statuses", "", "", "Component Statuses")
	_groundstationCmd.Flags().StringVarP(&_groundstationConfigData, "config-data", "", "", "Config Data")
	_groundstationCmd.Flags().StringVarP(&_groundstationConfigId, "config-id", "", "", "Config ID")
	_groundstationCmd.Flags().StringVarP(&_groundstationConfigType, "config-type", "", "", "Config Type")
	_groundstationCmd.Flags().StringVarP(&_groundstationContactId, "contact-id", "", "", "Contact ID")
	_groundstationCmd.Flags().StringVarP(&_groundstationContactPostPassDurationSeconds, "contact-post-pass-duration-seconds", "", "", "Contact Post Pass Duration Seconds")
	_groundstationCmd.Flags().StringVarP(&_groundstationContactPrePassDurationSeconds, "contact-pre-pass-duration-seconds", "", "", "Contact Pre Pass Duration Seconds")
	_groundstationCmd.Flags().StringVarP(&_groundstationDataflowEdges, "dataflow-edges", "", "", "Dataflow Edges")
	_groundstationCmd.Flags().StringVarP(&_groundstationDataflowEndpointGroupId, "dataflow-endpoint-group-id", "", "", "Dataflow Endpoint Group ID")
	_groundstationCmd.Flags().StringVarP(&_groundstationDiscoveryData, "discovery-data", "", "", "Discovery Data")
	_groundstationCmd.Flags().StringVarP(&_groundstationEnabled, "enabled", "", "", "Enabled")
	_groundstationCmd.Flags().StringVarP(&_groundstationEndTime, "end-time", "", "", "End Time")
	_groundstationCmd.Flags().StringVarP(&_groundstationEndpointDetails, "endpoint-details", "", "", "Endpoint Details")
	_groundstationCmd.Flags().StringVarP(&_groundstationEndpoints, "endpoints", "", "", "Endpoints")
	_groundstationCmd.Flags().StringVarP(&_groundstationEphemeris, "ephemeris", "", "", "Ephemeris")
	_groundstationCmd.Flags().StringVarP(&_groundstationEphemerisId, "ephemeris-id", "", "", "Ephemeris ID")
	_groundstationCmd.Flags().StringVarP(&_groundstationEphemerisType, "ephemeris-type", "", "", "Ephemeris Type")
	_groundstationCmd.Flags().StringVarP(&_groundstationExpirationTime, "expiration-time", "", "", "Expiration Time")
	_groundstationCmd.Flags().StringVarP(&_groundstationGroundStation, "ground-station", "", "", "Ground Station")
	_groundstationCmd.Flags().StringVarP(&_groundstationKmsKeyArn, "kms-key-arn", "", "", "KMS Key ARN")
	_groundstationCmd.Flags().StringVarP(&_groundstationMaxResults, "max-results", "", "", "Max Results")
	_groundstationCmd.Flags().StringVarP(&_groundstationMinimumViableContactDurationSeconds, "minimum-viable-contact-duration-seconds", "", "", "Minimum Viable Contact Duration Seconds")
	_groundstationCmd.Flags().StringVarP(&_groundstationMissionProfileArn, "mission-profile-arn", "", "", "Mission Profile ARN")
	_groundstationCmd.Flags().StringVarP(&_groundstationMissionProfileId, "mission-profile-id", "", "", "Mission Profile ID")
	_groundstationCmd.Flags().StringVarP(&_groundstationMonth, "month", "", "", "Month")
	_groundstationCmd.Flags().StringVarP(&_groundstationName, "name", "", "", "Name")
	_groundstationCmd.Flags().StringVarP(&_groundstationNextToken, "next-token", "", "", "Next Token")
	_groundstationCmd.Flags().StringVarP(&_groundstationPriority, "priority", "", "", "Priority")
	_groundstationCmd.Flags().StringVarP(&_groundstationResourceArn, "resource-arn", "", "", "Resource ARN")
	_groundstationCmd.Flags().StringVarP(&_groundstationSatelliteArn, "satellite-arn", "", "", "Satellite ARN")
	_groundstationCmd.Flags().StringVarP(&_groundstationSatelliteId, "satellite-id", "", "", "Satellite ID")
	_groundstationCmd.Flags().StringVarP(&_groundstationStartTime, "start-time", "", "", "Start Time")
	_groundstationCmd.Flags().StringVarP(&_groundstationStatusList, "status-list", "", "", "Status List")
	_groundstationCmd.Flags().StringVarP(&_groundstationStreamsKmsKey, "streams-kms-key", "", "", "Streams KMS Key")
	_groundstationCmd.Flags().StringVarP(&_groundstationStreamsKmsRole, "streams-kms-role", "", "", "Streams KMS Role")
	_groundstationCmd.Flags().StringSliceVarP(&_groundstationTagKeys, "tag-keys", "", nil, "Tag Keys")
	_groundstationCmd.Flags().StringVarP(&_groundstationTags, "tags", "", "", "Tags")
	_groundstationCmd.Flags().StringVarP(&_groundstationTaskId, "task-id", "", "", "Task ID")
	_groundstationCmd.Flags().StringVarP(&_groundstationTelemetrySinkConfigArn, "telemetry-sink-config-arn", "", "", "Telemetry Sink Config ARN")
	_groundstationCmd.Flags().StringVarP(&_groundstationTrackingConfigArn, "tracking-config-arn", "", "", "Tracking Config ARN")
	_groundstationCmd.Flags().StringVarP(&_groundstationTrackingOverrides, "tracking-overrides", "", "", "Tracking Overrides")
	_groundstationCmd.Flags().StringVarP(&_groundstationYear, "year", "", "", "Year")

	_groundstationCmd.Flags().BoolVarP(&_groundstationCancelContact, "cancel-contact", "", false, "Cancel Contact")
	_groundstationCmd.Flags().BoolVarP(&_groundstationCreateConfig, "create-config", "", false, "Create Config")
	_groundstationCmd.Flags().BoolVarP(&_groundstationCreateDataflowEndpointGroup, "create-dataflow-endpoint-group", "", false, "Create Dataflow Endpoint Group")
	_groundstationCmd.Flags().BoolVarP(&_groundstationCreateDataflowEndpointGroupV2, "create-dataflow-endpoint-group-v2", "", false, "Create Dataflow Endpoint Group V2")
	_groundstationCmd.Flags().BoolVarP(&_groundstationCreateEphemeris, "create-ephemeris", "", false, "Create Ephemeris")
	_groundstationCmd.Flags().BoolVarP(&_groundstationCreateMissionProfile, "create-mission-profile", "", false, "Create Mission Profile")
	_groundstationCmd.Flags().BoolVarP(&_groundstationDeleteConfig, "delete-config", "", false, "Delete Config")
	_groundstationCmd.Flags().BoolVarP(&_groundstationDeleteDataflowEndpointGroup, "delete-dataflow-endpoint-group", "", false, "Delete Dataflow Endpoint Group")
	_groundstationCmd.Flags().BoolVarP(&_groundstationDeleteEphemeris, "delete-ephemeris", "", false, "Delete Ephemeris")
	_groundstationCmd.Flags().BoolVarP(&_groundstationDeleteMissionProfile, "delete-mission-profile", "", false, "Delete Mission Profile")
	_groundstationCmd.Flags().BoolVarP(&_groundstationDescribeContact, "describe-contact", "", false, "Describe Contact")
	_groundstationCmd.Flags().BoolVarP(&_groundstationDescribeEphemeris, "describe-ephemeris", "", false, "Describe Ephemeris")
	_groundstationCmd.Flags().BoolVarP(&_groundstationGetAgentConfiguration, "get-agent-configuration", "", false, "Get Agent Configuration")
	_groundstationCmd.Flags().BoolVarP(&_groundstationGetAgentTaskResponseUrl, "get-agent-task-response-url", "", false, "Get Agent Task Response URL")
	_groundstationCmd.Flags().BoolVarP(&_groundstationGetConfig, "get-config", "", false, "Get Config")
	_groundstationCmd.Flags().BoolVarP(&_groundstationGetDataflowEndpointGroup, "get-dataflow-endpoint-group", "", false, "Get Dataflow Endpoint Group")
	_groundstationCmd.Flags().BoolVarP(&_groundstationGetMinuteUsage, "get-minute-usage", "", false, "Get Minute Usage")
	_groundstationCmd.Flags().BoolVarP(&_groundstationGetMissionProfile, "get-mission-profile", "", false, "Get Mission Profile")
	_groundstationCmd.Flags().BoolVarP(&_groundstationGetSatellite, "get-satellite", "", false, "Get Satellite")
	_groundstationCmd.Flags().BoolVarP(&_groundstationListConfigs, "list-configs", "", false, "List Configs")
	_groundstationCmd.Flags().BoolVarP(&_groundstationListContacts, "list-contacts", "", false, "List Contacts")
	_groundstationCmd.Flags().BoolVarP(&_groundstationListDataflowEndpointGroups, "list-dataflow-endpoint-groups", "", false, "List Dataflow Endpoint Groups")
	_groundstationCmd.Flags().BoolVarP(&_groundstationListEphemerides, "list-ephemerides", "", false, "List Ephemerides")
	_groundstationCmd.Flags().BoolVarP(&_groundstationListGroundStations, "list-ground-stations", "", false, "List Ground Stations")
	_groundstationCmd.Flags().BoolVarP(&_groundstationListMissionProfiles, "list-mission-profiles", "", false, "List Mission Profiles")
	_groundstationCmd.Flags().BoolVarP(&_groundstationListSatellites, "list-satellites", "", false, "List Satellites")
	_groundstationCmd.Flags().BoolVarP(&_groundstationListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_groundstationCmd.Flags().BoolVarP(&_groundstationRegisterAgent, "register-agent", "", false, "Register Agent")
	_groundstationCmd.Flags().BoolVarP(&_groundstationReserveContact, "reserve-contact", "", false, "Reserve Contact")
	_groundstationCmd.Flags().BoolVarP(&_groundstationTagResource, "tag-resource", "", false, "Tag Resource")
	_groundstationCmd.Flags().BoolVarP(&_groundstationUntagResource, "untag-resource", "", false, "Untag Resource")
	_groundstationCmd.Flags().BoolVarP(&_groundstationUpdateAgentStatus, "update-agent-status", "", false, "Update Agent Status")
	_groundstationCmd.Flags().BoolVarP(&_groundstationUpdateConfig, "update-config", "", false, "Update Config")
	_groundstationCmd.Flags().BoolVarP(&_groundstationUpdateEphemeris, "update-ephemeris", "", false, "Update Ephemeris")
	_groundstationCmd.Flags().BoolVarP(&_groundstationUpdateMissionProfile, "update-mission-profile", "", false, "Update Mission Profile")

}
