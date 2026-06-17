package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkvoice"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// chimesdkvoiceCmd represents the chimesdkvoice command
var _chimesdkvoiceCmd = &cobra.Command{
	Use:   "chimesdkvoice",
	Short: "AWS chimesdkvoice CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := chimesdkvoice.NewFromConfig(cfg)
		if _chimesdkvoiceAssociatePhoneNumbersWithVoiceConnector {
			chimesdkvoice_AssociatePhoneNumbersWithVoiceConnector(cfg, client)
			return
		}
		if _chimesdkvoiceAssociatePhoneNumbersWithVoiceConnectorGroup {
			chimesdkvoice_AssociatePhoneNumbersWithVoiceConnectorGroup(cfg, client)
			return
		}
		if _chimesdkvoiceBatchDeletePhoneNumber {
			chimesdkvoice_BatchDeletePhoneNumber(cfg, client)
			return
		}
		if _chimesdkvoiceBatchUpdatePhoneNumber {
			chimesdkvoice_BatchUpdatePhoneNumber(cfg, client)
			return
		}
		if _chimesdkvoiceCreatePhoneNumberOrder {
			chimesdkvoice_CreatePhoneNumberOrder(cfg, client)
			return
		}
		if _chimesdkvoiceCreateProxySession {
			chimesdkvoice_CreateProxySession(cfg, client)
			return
		}
		if _chimesdkvoiceCreateSipMediaApplication {
			chimesdkvoice_CreateSipMediaApplication(cfg, client)
			return
		}
		if _chimesdkvoiceCreateSipMediaApplicationCall {
			chimesdkvoice_CreateSipMediaApplicationCall(cfg, client)
			return
		}
		if _chimesdkvoiceCreateSipRule {
			chimesdkvoice_CreateSipRule(cfg, client)
			return
		}
		if _chimesdkvoiceCreateVoiceConnector {
			chimesdkvoice_CreateVoiceConnector(cfg, client)
			return
		}
		if _chimesdkvoiceCreateVoiceConnectorGroup {
			chimesdkvoice_CreateVoiceConnectorGroup(cfg, client)
			return
		}
		if _chimesdkvoiceCreateVoiceProfile {
			chimesdkvoice_CreateVoiceProfile(cfg, client)
			return
		}
		if _chimesdkvoiceCreateVoiceProfileDomain {
			chimesdkvoice_CreateVoiceProfileDomain(cfg, client)
			return
		}
		if _chimesdkvoiceDeletePhoneNumber {
			chimesdkvoice_DeletePhoneNumber(cfg, client)
			return
		}
		if _chimesdkvoiceDeleteProxySession {
			chimesdkvoice_DeleteProxySession(cfg, client)
			return
		}
		if _chimesdkvoiceDeleteSipMediaApplication {
			chimesdkvoice_DeleteSipMediaApplication(cfg, client)
			return
		}
		if _chimesdkvoiceDeleteSipRule {
			chimesdkvoice_DeleteSipRule(cfg, client)
			return
		}
		if _chimesdkvoiceDeleteVoiceConnector {
			chimesdkvoice_DeleteVoiceConnector(cfg, client)
			return
		}
		if _chimesdkvoiceDeleteVoiceConnectorEmergencyCallingConfiguration {
			chimesdkvoice_DeleteVoiceConnectorEmergencyCallingConfiguration(cfg, client)
			return
		}
		if _chimesdkvoiceDeleteVoiceConnectorExternalSystemsConfiguration {
			chimesdkvoice_DeleteVoiceConnectorExternalSystemsConfiguration(cfg, client)
			return
		}
		if _chimesdkvoiceDeleteVoiceConnectorGroup {
			chimesdkvoice_DeleteVoiceConnectorGroup(cfg, client)
			return
		}
		if _chimesdkvoiceDeleteVoiceConnectorOrigination {
			chimesdkvoice_DeleteVoiceConnectorOrigination(cfg, client)
			return
		}
		if _chimesdkvoiceDeleteVoiceConnectorProxy {
			chimesdkvoice_DeleteVoiceConnectorProxy(cfg, client)
			return
		}
		if _chimesdkvoiceDeleteVoiceConnectorStreamingConfiguration {
			chimesdkvoice_DeleteVoiceConnectorStreamingConfiguration(cfg, client)
			return
		}
		if _chimesdkvoiceDeleteVoiceConnectorTermination {
			chimesdkvoice_DeleteVoiceConnectorTermination(cfg, client)
			return
		}
		if _chimesdkvoiceDeleteVoiceConnectorTerminationCredentials {
			chimesdkvoice_DeleteVoiceConnectorTerminationCredentials(cfg, client)
			return
		}
		if _chimesdkvoiceDeleteVoiceProfile {
			chimesdkvoice_DeleteVoiceProfile(cfg, client)
			return
		}
		if _chimesdkvoiceDeleteVoiceProfileDomain {
			chimesdkvoice_DeleteVoiceProfileDomain(cfg, client)
			return
		}
		if _chimesdkvoiceDisassociatePhoneNumbersFromVoiceConnector {
			chimesdkvoice_DisassociatePhoneNumbersFromVoiceConnector(cfg, client)
			return
		}
		if _chimesdkvoiceDisassociatePhoneNumbersFromVoiceConnectorGroup {
			chimesdkvoice_DisassociatePhoneNumbersFromVoiceConnectorGroup(cfg, client)
			return
		}
		if _chimesdkvoiceGetGlobalSettings {
			chimesdkvoice_GetGlobalSettings(cfg, client)
			return
		}
		if _chimesdkvoiceGetPhoneNumber {
			chimesdkvoice_GetPhoneNumber(cfg, client)
			return
		}
		if _chimesdkvoiceGetPhoneNumberOrder {
			chimesdkvoice_GetPhoneNumberOrder(cfg, client)
			return
		}
		if _chimesdkvoiceGetPhoneNumberSettings {
			chimesdkvoice_GetPhoneNumberSettings(cfg, client)
			return
		}
		if _chimesdkvoiceGetProxySession {
			chimesdkvoice_GetProxySession(cfg, client)
			return
		}
		if _chimesdkvoiceGetSipMediaApplication {
			chimesdkvoice_GetSipMediaApplication(cfg, client)
			return
		}
		if _chimesdkvoiceGetSipMediaApplicationAlexaSkillConfiguration {
			chimesdkvoice_GetSipMediaApplicationAlexaSkillConfiguration(cfg, client)
			return
		}
		if _chimesdkvoiceGetSipMediaApplicationLoggingConfiguration {
			chimesdkvoice_GetSipMediaApplicationLoggingConfiguration(cfg, client)
			return
		}
		if _chimesdkvoiceGetSipRule {
			chimesdkvoice_GetSipRule(cfg, client)
			return
		}
		if _chimesdkvoiceGetSpeakerSearchTask {
			chimesdkvoice_GetSpeakerSearchTask(cfg, client)
			return
		}
		if _chimesdkvoiceGetVoiceConnector {
			chimesdkvoice_GetVoiceConnector(cfg, client)
			return
		}
		if _chimesdkvoiceGetVoiceConnectorEmergencyCallingConfiguration {
			chimesdkvoice_GetVoiceConnectorEmergencyCallingConfiguration(cfg, client)
			return
		}
		if _chimesdkvoiceGetVoiceConnectorExternalSystemsConfiguration {
			chimesdkvoice_GetVoiceConnectorExternalSystemsConfiguration(cfg, client)
			return
		}
		if _chimesdkvoiceGetVoiceConnectorGroup {
			chimesdkvoice_GetVoiceConnectorGroup(cfg, client)
			return
		}
		if _chimesdkvoiceGetVoiceConnectorLoggingConfiguration {
			chimesdkvoice_GetVoiceConnectorLoggingConfiguration(cfg, client)
			return
		}
		if _chimesdkvoiceGetVoiceConnectorOrigination {
			chimesdkvoice_GetVoiceConnectorOrigination(cfg, client)
			return
		}
		if _chimesdkvoiceGetVoiceConnectorProxy {
			chimesdkvoice_GetVoiceConnectorProxy(cfg, client)
			return
		}
		if _chimesdkvoiceGetVoiceConnectorStreamingConfiguration {
			chimesdkvoice_GetVoiceConnectorStreamingConfiguration(cfg, client)
			return
		}
		if _chimesdkvoiceGetVoiceConnectorTermination {
			chimesdkvoice_GetVoiceConnectorTermination(cfg, client)
			return
		}
		if _chimesdkvoiceGetVoiceConnectorTerminationHealth {
			chimesdkvoice_GetVoiceConnectorTerminationHealth(cfg, client)
			return
		}
		if _chimesdkvoiceGetVoiceProfile {
			chimesdkvoice_GetVoiceProfile(cfg, client)
			return
		}
		if _chimesdkvoiceGetVoiceProfileDomain {
			chimesdkvoice_GetVoiceProfileDomain(cfg, client)
			return
		}
		if _chimesdkvoiceGetVoiceToneAnalysisTask {
			chimesdkvoice_GetVoiceToneAnalysisTask(cfg, client)
			return
		}
		if _chimesdkvoiceListAvailableVoiceConnectorRegions {
			chimesdkvoice_ListAvailableVoiceConnectorRegions(cfg, client)
			return
		}
		if _chimesdkvoiceListPhoneNumberOrders {
			chimesdkvoice_ListPhoneNumberOrders(cfg, client)
			return
		}
		if _chimesdkvoiceListPhoneNumbers {
			chimesdkvoice_ListPhoneNumbers(cfg, client)
			return
		}
		if _chimesdkvoiceListProxySessions {
			chimesdkvoice_ListProxySessions(cfg, client)
			return
		}
		if _chimesdkvoiceListSipMediaApplications {
			chimesdkvoice_ListSipMediaApplications(cfg, client)
			return
		}
		if _chimesdkvoiceListSipRules {
			chimesdkvoice_ListSipRules(cfg, client)
			return
		}
		if _chimesdkvoiceListSupportedPhoneNumberCountries {
			chimesdkvoice_ListSupportedPhoneNumberCountries(cfg, client)
			return
		}
		if _chimesdkvoiceListTagsForResource {
			chimesdkvoice_ListTagsForResource(cfg, client)
			return
		}
		if _chimesdkvoiceListVoiceConnectorGroups {
			chimesdkvoice_ListVoiceConnectorGroups(cfg, client)
			return
		}
		if _chimesdkvoiceListVoiceConnectorTerminationCredentials {
			chimesdkvoice_ListVoiceConnectorTerminationCredentials(cfg, client)
			return
		}
		if _chimesdkvoiceListVoiceConnectors {
			chimesdkvoice_ListVoiceConnectors(cfg, client)
			return
		}
		if _chimesdkvoiceListVoiceProfileDomains {
			chimesdkvoice_ListVoiceProfileDomains(cfg, client)
			return
		}
		if _chimesdkvoiceListVoiceProfiles {
			chimesdkvoice_ListVoiceProfiles(cfg, client)
			return
		}
		if _chimesdkvoicePutSipMediaApplicationAlexaSkillConfiguration {
			chimesdkvoice_PutSipMediaApplicationAlexaSkillConfiguration(cfg, client)
			return
		}
		if _chimesdkvoicePutSipMediaApplicationLoggingConfiguration {
			chimesdkvoice_PutSipMediaApplicationLoggingConfiguration(cfg, client)
			return
		}
		if _chimesdkvoicePutVoiceConnectorEmergencyCallingConfiguration {
			chimesdkvoice_PutVoiceConnectorEmergencyCallingConfiguration(cfg, client)
			return
		}
		if _chimesdkvoicePutVoiceConnectorExternalSystemsConfiguration {
			chimesdkvoice_PutVoiceConnectorExternalSystemsConfiguration(cfg, client)
			return
		}
		if _chimesdkvoicePutVoiceConnectorLoggingConfiguration {
			chimesdkvoice_PutVoiceConnectorLoggingConfiguration(cfg, client)
			return
		}
		if _chimesdkvoicePutVoiceConnectorOrigination {
			chimesdkvoice_PutVoiceConnectorOrigination(cfg, client)
			return
		}
		if _chimesdkvoicePutVoiceConnectorProxy {
			chimesdkvoice_PutVoiceConnectorProxy(cfg, client)
			return
		}
		if _chimesdkvoicePutVoiceConnectorStreamingConfiguration {
			chimesdkvoice_PutVoiceConnectorStreamingConfiguration(cfg, client)
			return
		}
		if _chimesdkvoicePutVoiceConnectorTermination {
			chimesdkvoice_PutVoiceConnectorTermination(cfg, client)
			return
		}
		if _chimesdkvoicePutVoiceConnectorTerminationCredentials {
			chimesdkvoice_PutVoiceConnectorTerminationCredentials(cfg, client)
			return
		}
		if _chimesdkvoiceRestorePhoneNumber {
			chimesdkvoice_RestorePhoneNumber(cfg, client)
			return
		}
		if _chimesdkvoiceSearchAvailablePhoneNumbers {
			chimesdkvoice_SearchAvailablePhoneNumbers(cfg, client)
			return
		}
		if _chimesdkvoiceStartSpeakerSearchTask {
			chimesdkvoice_StartSpeakerSearchTask(cfg, client)
			return
		}
		if _chimesdkvoiceStartVoiceToneAnalysisTask {
			chimesdkvoice_StartVoiceToneAnalysisTask(cfg, client)
			return
		}
		if _chimesdkvoiceStopSpeakerSearchTask {
			chimesdkvoice_StopSpeakerSearchTask(cfg, client)
			return
		}
		if _chimesdkvoiceStopVoiceToneAnalysisTask {
			chimesdkvoice_StopVoiceToneAnalysisTask(cfg, client)
			return
		}
		if _chimesdkvoiceTagResource {
			chimesdkvoice_TagResource(cfg, client)
			return
		}
		if _chimesdkvoiceUntagResource {
			chimesdkvoice_UntagResource(cfg, client)
			return
		}
		if _chimesdkvoiceUpdateGlobalSettings {
			chimesdkvoice_UpdateGlobalSettings(cfg, client)
			return
		}
		if _chimesdkvoiceUpdatePhoneNumber {
			chimesdkvoice_UpdatePhoneNumber(cfg, client)
			return
		}
		if _chimesdkvoiceUpdatePhoneNumberSettings {
			chimesdkvoice_UpdatePhoneNumberSettings(cfg, client)
			return
		}
		if _chimesdkvoiceUpdateProxySession {
			chimesdkvoice_UpdateProxySession(cfg, client)
			return
		}
		if _chimesdkvoiceUpdateSipMediaApplication {
			chimesdkvoice_UpdateSipMediaApplication(cfg, client)
			return
		}
		if _chimesdkvoiceUpdateSipMediaApplicationCall {
			chimesdkvoice_UpdateSipMediaApplicationCall(cfg, client)
			return
		}
		if _chimesdkvoiceUpdateSipRule {
			chimesdkvoice_UpdateSipRule(cfg, client)
			return
		}
		if _chimesdkvoiceUpdateVoiceConnector {
			chimesdkvoice_UpdateVoiceConnector(cfg, client)
			return
		}
		if _chimesdkvoiceUpdateVoiceConnectorGroup {
			chimesdkvoice_UpdateVoiceConnectorGroup(cfg, client)
			return
		}
		if _chimesdkvoiceUpdateVoiceProfile {
			chimesdkvoice_UpdateVoiceProfile(cfg, client)
			return
		}
		if _chimesdkvoiceUpdateVoiceProfileDomain {
			chimesdkvoice_UpdateVoiceProfileDomain(cfg, client)
			return
		}
		if _chimesdkvoiceValidateE911Address {
			chimesdkvoice_ValidateE911Address(cfg, client)
			return
		}

	},
}

var (
	_chimesdkvoiceAssociatePhoneNumbersWithVoiceConnector           bool
	_chimesdkvoiceAssociatePhoneNumbersWithVoiceConnectorGroup      bool
	_chimesdkvoiceBatchDeletePhoneNumber                            bool
	_chimesdkvoiceBatchUpdatePhoneNumber                            bool
	_chimesdkvoiceCreatePhoneNumberOrder                            bool
	_chimesdkvoiceCreateProxySession                                bool
	_chimesdkvoiceCreateSipMediaApplication                         bool
	_chimesdkvoiceCreateSipMediaApplicationCall                     bool
	_chimesdkvoiceCreateSipRule                                     bool
	_chimesdkvoiceCreateVoiceConnector                              bool
	_chimesdkvoiceCreateVoiceConnectorGroup                         bool
	_chimesdkvoiceCreateVoiceProfile                                bool
	_chimesdkvoiceCreateVoiceProfileDomain                          bool
	_chimesdkvoiceDeletePhoneNumber                                 bool
	_chimesdkvoiceDeleteProxySession                                bool
	_chimesdkvoiceDeleteSipMediaApplication                         bool
	_chimesdkvoiceDeleteSipRule                                     bool
	_chimesdkvoiceDeleteVoiceConnector                              bool
	_chimesdkvoiceDeleteVoiceConnectorEmergencyCallingConfiguration bool
	_chimesdkvoiceDeleteVoiceConnectorExternalSystemsConfiguration  bool
	_chimesdkvoiceDeleteVoiceConnectorGroup                         bool
	_chimesdkvoiceDeleteVoiceConnectorOrigination                   bool
	_chimesdkvoiceDeleteVoiceConnectorProxy                         bool
	_chimesdkvoiceDeleteVoiceConnectorStreamingConfiguration        bool
	_chimesdkvoiceDeleteVoiceConnectorTermination                   bool
	_chimesdkvoiceDeleteVoiceConnectorTerminationCredentials        bool
	_chimesdkvoiceDeleteVoiceProfile                                bool
	_chimesdkvoiceDeleteVoiceProfileDomain                          bool
	_chimesdkvoiceDisassociatePhoneNumbersFromVoiceConnector        bool
	_chimesdkvoiceDisassociatePhoneNumbersFromVoiceConnectorGroup   bool
	_chimesdkvoiceGetGlobalSettings                                 bool
	_chimesdkvoiceGetPhoneNumber                                    bool
	_chimesdkvoiceGetPhoneNumberOrder                               bool
	_chimesdkvoiceGetPhoneNumberSettings                            bool
	_chimesdkvoiceGetProxySession                                   bool
	_chimesdkvoiceGetSipMediaApplication                            bool
	_chimesdkvoiceGetSipMediaApplicationAlexaSkillConfiguration     bool
	_chimesdkvoiceGetSipMediaApplicationLoggingConfiguration        bool
	_chimesdkvoiceGetSipRule                                        bool
	_chimesdkvoiceGetSpeakerSearchTask                              bool
	_chimesdkvoiceGetVoiceConnector                                 bool
	_chimesdkvoiceGetVoiceConnectorEmergencyCallingConfiguration    bool
	_chimesdkvoiceGetVoiceConnectorExternalSystemsConfiguration     bool
	_chimesdkvoiceGetVoiceConnectorGroup                            bool
	_chimesdkvoiceGetVoiceConnectorLoggingConfiguration             bool
	_chimesdkvoiceGetVoiceConnectorOrigination                      bool
	_chimesdkvoiceGetVoiceConnectorProxy                            bool
	_chimesdkvoiceGetVoiceConnectorStreamingConfiguration           bool
	_chimesdkvoiceGetVoiceConnectorTermination                      bool
	_chimesdkvoiceGetVoiceConnectorTerminationHealth                bool
	_chimesdkvoiceGetVoiceProfile                                   bool
	_chimesdkvoiceGetVoiceProfileDomain                             bool
	_chimesdkvoiceGetVoiceToneAnalysisTask                          bool
	_chimesdkvoiceListAvailableVoiceConnectorRegions                bool
	_chimesdkvoiceListPhoneNumberOrders                             bool
	_chimesdkvoiceListPhoneNumbers                                  bool
	_chimesdkvoiceListProxySessions                                 bool
	_chimesdkvoiceListSipMediaApplications                          bool
	_chimesdkvoiceListSipRules                                      bool
	_chimesdkvoiceListSupportedPhoneNumberCountries                 bool
	_chimesdkvoiceListTagsForResource                               bool
	_chimesdkvoiceListVoiceConnectorGroups                          bool
	_chimesdkvoiceListVoiceConnectorTerminationCredentials          bool
	_chimesdkvoiceListVoiceConnectors                               bool
	_chimesdkvoiceListVoiceProfileDomains                           bool
	_chimesdkvoiceListVoiceProfiles                                 bool
	_chimesdkvoicePutSipMediaApplicationAlexaSkillConfiguration     bool
	_chimesdkvoicePutSipMediaApplicationLoggingConfiguration        bool
	_chimesdkvoicePutVoiceConnectorEmergencyCallingConfiguration    bool
	_chimesdkvoicePutVoiceConnectorExternalSystemsConfiguration     bool
	_chimesdkvoicePutVoiceConnectorLoggingConfiguration             bool
	_chimesdkvoicePutVoiceConnectorOrigination                      bool
	_chimesdkvoicePutVoiceConnectorProxy                            bool
	_chimesdkvoicePutVoiceConnectorStreamingConfiguration           bool
	_chimesdkvoicePutVoiceConnectorTermination                      bool
	_chimesdkvoicePutVoiceConnectorTerminationCredentials           bool
	_chimesdkvoiceRestorePhoneNumber                                bool
	_chimesdkvoiceSearchAvailablePhoneNumbers                       bool
	_chimesdkvoiceStartSpeakerSearchTask                            bool
	_chimesdkvoiceStartVoiceToneAnalysisTask                        bool
	_chimesdkvoiceStopSpeakerSearchTask                             bool
	_chimesdkvoiceStopVoiceToneAnalysisTask                         bool
	_chimesdkvoiceTagResource                                       bool
	_chimesdkvoiceUntagResource                                     bool
	_chimesdkvoiceUpdateGlobalSettings                              bool
	_chimesdkvoiceUpdatePhoneNumber                                 bool
	_chimesdkvoiceUpdatePhoneNumberSettings                         bool
	_chimesdkvoiceUpdateProxySession                                bool
	_chimesdkvoiceUpdateSipMediaApplication                         bool
	_chimesdkvoiceUpdateSipMediaApplicationCall                     bool
	_chimesdkvoiceUpdateSipRule                                     bool
	_chimesdkvoiceUpdateVoiceConnector                              bool
	_chimesdkvoiceUpdateVoiceConnectorGroup                         bool
	_chimesdkvoiceUpdateVoiceProfile                                bool
	_chimesdkvoiceUpdateVoiceProfileDomain                          bool
	_chimesdkvoiceValidateE911Address                               bool

	_chimesdkvoiceAreaCode                                   string
	_chimesdkvoiceArguments                                  string
	_chimesdkvoiceArgumentsMap                               string
	_chimesdkvoiceAwsAccountId                               string
	_chimesdkvoiceAwsRegion                                  string
	_chimesdkvoiceCallLeg                                    string
	_chimesdkvoiceCallingName                                string
	_chimesdkvoiceCapabilities                               string
	_chimesdkvoiceCity                                       string
	_chimesdkvoiceClientRequestToken                         string
	_chimesdkvoiceContactCenterSystemTypes                   string
	_chimesdkvoiceCountry                                    string
	_chimesdkvoiceCredentials                                string
	_chimesdkvoiceDefaultSessionExpiryMinutes                string
	_chimesdkvoiceDescription                                string
	_chimesdkvoiceDisabled                                   string
	_chimesdkvoiceE164PhoneNumbers                           []string
	_chimesdkvoiceEmergencyCallingConfiguration              string
	_chimesdkvoiceEndpoints                                  string
	_chimesdkvoiceExpiryMinutes                              string
	_chimesdkvoiceFallBackPhoneNumber                        string
	_chimesdkvoiceFilterName                                 string
	_chimesdkvoiceFilterValue                                string
	_chimesdkvoiceForceAssociate                             string
	_chimesdkvoiceFromPhoneNumber                            string
	_chimesdkvoiceGeoMatchLevel                              string
	_chimesdkvoiceGeoMatchParams                             string
	_chimesdkvoiceIntegrationType                            string
	_chimesdkvoiceIsCaller                                   string
	_chimesdkvoiceLanguageCode                               string
	_chimesdkvoiceLoggingConfiguration                       string
	_chimesdkvoiceMaxResults                                 string
	_chimesdkvoiceName                                       string
	_chimesdkvoiceNetworkType                                string
	_chimesdkvoiceNextToken                                  string
	_chimesdkvoiceNumberSelectionBehavior                    string
	_chimesdkvoiceOrigination                                string
	_chimesdkvoiceParticipantPhoneNumbers                    []string
	_chimesdkvoicePhoneNumberId                              string
	_chimesdkvoicePhoneNumberIds                             []string
	_chimesdkvoicePhoneNumberOrderId                         string
	_chimesdkvoicePhoneNumberPoolCountries                   []string
	_chimesdkvoicePhoneNumberType                            string
	_chimesdkvoicePostalCode                                 string
	_chimesdkvoiceProductType                                string
	_chimesdkvoiceProxySessionId                             string
	_chimesdkvoiceRequireEncryption                          string
	_chimesdkvoiceResourceARN                                string
	_chimesdkvoiceServerSideEncryptionConfiguration          string
	_chimesdkvoiceSessionBorderControllerTypes               string
	_chimesdkvoiceSipHeaders                                 string
	_chimesdkvoiceSipMediaApplicationAlexaSkillConfiguration string
	_chimesdkvoiceSipMediaApplicationId                      string
	_chimesdkvoiceSipMediaApplicationLoggingConfiguration    string
	_chimesdkvoiceSipRuleId                                  string
	_chimesdkvoiceSpeakerSearchTaskId                        string
	_chimesdkvoiceState                                      string
	_chimesdkvoiceStatus                                     string
	_chimesdkvoiceStreamingConfiguration                     string
	_chimesdkvoiceStreetInfo                                 string
	_chimesdkvoiceStreetNumber                               string
	_chimesdkvoiceTagKeys                                    []string
	_chimesdkvoiceTags                                       string
	_chimesdkvoiceTargetApplications                         string
	_chimesdkvoiceTermination                                string
	_chimesdkvoiceToPhoneNumber                              string
	_chimesdkvoiceTollFreePrefix                             string
	_chimesdkvoiceTransactionId                              string
	_chimesdkvoiceTriggerType                                string
	_chimesdkvoiceTriggerValue                               string
	_chimesdkvoiceUpdatePhoneNumberRequestItems              string
	_chimesdkvoiceUsernames                                  []string
	_chimesdkvoiceVoiceConnector                             string
	_chimesdkvoiceVoiceConnectorGroupId                      string
	_chimesdkvoiceVoiceConnectorId                           string
	_chimesdkvoiceVoiceConnectorItems                        string
	_chimesdkvoiceVoiceProfileDomainId                       string
	_chimesdkvoiceVoiceProfileId                             string
	_chimesdkvoiceVoiceToneAnalysisTaskId                    string
)

// Associates phone numbers with the specified Amazon Chime SDK Voice Connector.
func chimesdkvoice_AssociatePhoneNumbersWithVoiceConnector(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.AssociatePhoneNumbersWithVoiceConnectorInput{
		// E164PhoneNumbers: []string, // Required
		// VoiceConnectorId: *string, // Required
	}

	if len(_chimesdkvoiceE164PhoneNumbers) > 0 {
		input.E164PhoneNumbers = append([]string(nil), _chimesdkvoiceE164PhoneNumbers...)
	}
	if len(_chimesdkvoiceVoiceConnectorId) > 0 {
		input.VoiceConnectorId = aws.String(_chimesdkvoiceVoiceConnectorId)
	}
	if len(_chimesdkvoiceForceAssociate) > 0 {
		if err := assignInputField(input, "ForceAssociate", _chimesdkvoiceForceAssociate); err != nil {
			log.Errorf("invalid --force-associate: %s", err.Error())
			return
		}
	}

	if resp, err := client.AssociatePhoneNumbersWithVoiceConnector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates phone numbers with the specified Amazon Chime SDK Voice Connector
// group.
func chimesdkvoice_AssociatePhoneNumbersWithVoiceConnectorGroup(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.AssociatePhoneNumbersWithVoiceConnectorGroupInput{
		// E164PhoneNumbers: []string, // Required
		// VoiceConnectorGroupId: *string, // Required
	}

	if len(_chimesdkvoiceE164PhoneNumbers) > 0 {
		input.E164PhoneNumbers = append([]string(nil), _chimesdkvoiceE164PhoneNumbers...)
	}
	if len(_chimesdkvoiceVoiceConnectorGroupId) > 0 {
		input.VoiceConnectorGroupId = aws.String(_chimesdkvoiceVoiceConnectorGroupId)
	}
	if len(_chimesdkvoiceForceAssociate) > 0 {
		if err := assignInputField(input, "ForceAssociate", _chimesdkvoiceForceAssociate); err != nil {
			log.Errorf("invalid --force-associate: %s", err.Error())
			return
		}
	}

	if resp, err := client.AssociatePhoneNumbersWithVoiceConnectorGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Moves phone numbers into the Deletion queue. Phone numbers must be
// disassociated from any users or Amazon Chime SDK Voice Connectors before they
// can be deleted.
//
// Phone numbers remain in the Deletion queue for 7 days before they are deleted
// permanently.
func chimesdkvoice_BatchDeletePhoneNumber(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.BatchDeletePhoneNumberInput{
		// PhoneNumberIds: []string, // Required
	}

	if len(_chimesdkvoicePhoneNumberIds) > 0 {
		input.PhoneNumberIds = append([]string(nil), _chimesdkvoicePhoneNumberIds...)
	}

	if resp, err := client.BatchDeletePhoneNumber(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates phone number product types, calling names, or phone number names. You
// can update one attribute at a time for each UpdatePhoneNumberRequestItem . For
// example, you can update the product type, the calling name, or phone name.
//
// You cannot have a duplicate phoneNumberId in a request.
func chimesdkvoice_BatchUpdatePhoneNumber(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.BatchUpdatePhoneNumberInput{
		// UpdatePhoneNumberRequestItems: []types.UpdatePhoneNumberRequestItem, // Required
	}

	if len(_chimesdkvoiceUpdatePhoneNumberRequestItems) > 0 {
		if err := assignInputField(input, "UpdatePhoneNumberRequestItems", _chimesdkvoiceUpdatePhoneNumberRequestItems); err != nil {
			log.Errorf("invalid --update-phone-number-request-items: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchUpdatePhoneNumber(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an order for phone numbers to be provisioned. For numbers outside the
// U.S., you must use the Amazon Chime SDK SIP media application dial-in product
// type.
func chimesdkvoice_CreatePhoneNumberOrder(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.CreatePhoneNumberOrderInput{
		// E164PhoneNumbers: []string, // Required
		// ProductType: types.PhoneNumberProductType, // Required
	}

	if len(_chimesdkvoiceE164PhoneNumbers) > 0 {
		input.E164PhoneNumbers = append([]string(nil), _chimesdkvoiceE164PhoneNumbers...)
	}
	if len(_chimesdkvoiceProductType) > 0 {
		if err := assignInputField(input, "ProductType", _chimesdkvoiceProductType); err != nil {
			log.Errorf("invalid --product-type: %s", err.Error())
			return
		}
	}
	if len(_chimesdkvoiceName) > 0 {
		input.Name = aws.String(_chimesdkvoiceName)
	}

	if resp, err := client.CreatePhoneNumberOrder(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a proxy session for the specified Amazon Chime SDK Voice Connector for
// the specified participant phone numbers.
func chimesdkvoice_CreateProxySession(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.CreateProxySessionInput{
		// Capabilities: []types.Capability, // Required
		// ParticipantPhoneNumbers: []string, // Required
		// VoiceConnectorId: *string, // Required
	}

	if len(_chimesdkvoiceCapabilities) > 0 {
		if err := assignInputField(input, "Capabilities", _chimesdkvoiceCapabilities); err != nil {
			log.Errorf("invalid --capabilities: %s", err.Error())
			return
		}
	}
	if len(_chimesdkvoiceParticipantPhoneNumbers) > 0 {
		input.ParticipantPhoneNumbers = append([]string(nil), _chimesdkvoiceParticipantPhoneNumbers...)
	}
	if len(_chimesdkvoiceVoiceConnectorId) > 0 {
		input.VoiceConnectorId = aws.String(_chimesdkvoiceVoiceConnectorId)
	}
	if len(_chimesdkvoiceExpiryMinutes) > 0 {
		if err := assignInputField(input, "ExpiryMinutes", _chimesdkvoiceExpiryMinutes); err != nil {
			log.Errorf("invalid --expiry-minutes: %s", err.Error())
			return
		}
	}
	if len(_chimesdkvoiceGeoMatchLevel) > 0 {
		if err := assignInputField(input, "GeoMatchLevel", _chimesdkvoiceGeoMatchLevel); err != nil {
			log.Errorf("invalid --geo-match-level: %s", err.Error())
			return
		}
	}
	if len(_chimesdkvoiceGeoMatchParams) > 0 {
		if err := assignInputField(input, "GeoMatchParams", _chimesdkvoiceGeoMatchParams); err != nil {
			log.Errorf("invalid --geo-match-params: %s", err.Error())
			return
		}
	}
	if len(_chimesdkvoiceName) > 0 {
		input.Name = aws.String(_chimesdkvoiceName)
	}
	if len(_chimesdkvoiceNumberSelectionBehavior) > 0 {
		if err := assignInputField(input, "NumberSelectionBehavior", _chimesdkvoiceNumberSelectionBehavior); err != nil {
			log.Errorf("invalid --number-selection-behavior: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateProxySession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a SIP media application. For more information about SIP media
// applications, see [Managing SIP media applications and rules]in the Amazon Chime SDK Administrator Guide.
//
// [Managing SIP media applications and rules]: https://docs.aws.amazon.com/chime-sdk/latest/ag/manage-sip-applications.html
func chimesdkvoice_CreateSipMediaApplication(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.CreateSipMediaApplicationInput{
		// AwsRegion: *string, // Required
		// Endpoints: []types.SipMediaApplicationEndpoint, // Required
		// Name: *string, // Required
	}

	if len(_chimesdkvoiceAwsRegion) > 0 {
		input.AwsRegion = aws.String(_chimesdkvoiceAwsRegion)
	}
	if len(_chimesdkvoiceEndpoints) > 0 {
		if err := assignInputField(input, "Endpoints", _chimesdkvoiceEndpoints); err != nil {
			log.Errorf("invalid --endpoints: %s", err.Error())
			return
		}
	}
	if len(_chimesdkvoiceName) > 0 {
		input.Name = aws.String(_chimesdkvoiceName)
	}
	if len(_chimesdkvoiceTags) > 0 {
		if err := assignInputField(input, "Tags", _chimesdkvoiceTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSipMediaApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an outbound call to a phone number from the phone number specified in
// the request, and it invokes the endpoint of the specified sipMediaApplicationId .
func chimesdkvoice_CreateSipMediaApplicationCall(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.CreateSipMediaApplicationCallInput{
		// FromPhoneNumber: *string, // Required
		// SipMediaApplicationId: *string, // Required
		// ToPhoneNumber: *string, // Required
	}

	if len(_chimesdkvoiceFromPhoneNumber) > 0 {
		input.FromPhoneNumber = aws.String(_chimesdkvoiceFromPhoneNumber)
	}
	if len(_chimesdkvoiceSipMediaApplicationId) > 0 {
		input.SipMediaApplicationId = aws.String(_chimesdkvoiceSipMediaApplicationId)
	}
	if len(_chimesdkvoiceToPhoneNumber) > 0 {
		input.ToPhoneNumber = aws.String(_chimesdkvoiceToPhoneNumber)
	}
	if len(_chimesdkvoiceArgumentsMap) > 0 {
		if err := assignInputField(input, "ArgumentsMap", _chimesdkvoiceArgumentsMap); err != nil {
			log.Errorf("invalid --arguments-map: %s", err.Error())
			return
		}
	}
	if len(_chimesdkvoiceSipHeaders) > 0 {
		if err := assignInputField(input, "SipHeaders", _chimesdkvoiceSipHeaders); err != nil {
			log.Errorf("invalid --sip-headers: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSipMediaApplicationCall(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a SIP rule, which can be used to run a SIP media application as a
// target for a specific trigger type. For more information about SIP rules, see [Managing SIP media applications and rules]
// in the Amazon Chime SDK Administrator Guide.
//
// [Managing SIP media applications and rules]: https://docs.aws.amazon.com/chime-sdk/latest/ag/manage-sip-applications.html
func chimesdkvoice_CreateSipRule(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.CreateSipRuleInput{
		// Name: *string, // Required
		// TriggerType: types.SipRuleTriggerType, // Required
		// TriggerValue: *string, // Required
	}

	if len(_chimesdkvoiceName) > 0 {
		input.Name = aws.String(_chimesdkvoiceName)
	}
	if len(_chimesdkvoiceTriggerType) > 0 {
		if err := assignInputField(input, "TriggerType", _chimesdkvoiceTriggerType); err != nil {
			log.Errorf("invalid --trigger-type: %s", err.Error())
			return
		}
	}
	if len(_chimesdkvoiceTriggerValue) > 0 {
		input.TriggerValue = aws.String(_chimesdkvoiceTriggerValue)
	}
	if len(_chimesdkvoiceDisabled) > 0 {
		if err := assignInputField(input, "Disabled", _chimesdkvoiceDisabled); err != nil {
			log.Errorf("invalid --disabled: %s", err.Error())
			return
		}
	}
	if len(_chimesdkvoiceTargetApplications) > 0 {
		if err := assignInputField(input, "TargetApplications", _chimesdkvoiceTargetApplications); err != nil {
			log.Errorf("invalid --target-applications: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSipRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon Chime SDK Voice Connector. For more information about Voice
// Connectors, see [Managing Amazon Chime SDK Voice Connector groups]in the Amazon Chime SDK Administrator Guide.
//
// [Managing Amazon Chime SDK Voice Connector groups]: https://docs.aws.amazon.com/chime-sdk/latest/ag/voice-connector-groups.html
func chimesdkvoice_CreateVoiceConnector(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.CreateVoiceConnectorInput{
		// Name: *string, // Required
		// RequireEncryption: *bool, // Required
	}

	if len(_chimesdkvoiceName) > 0 {
		input.Name = aws.String(_chimesdkvoiceName)
	}
	if len(_chimesdkvoiceRequireEncryption) > 0 {
		if err := assignInputField(input, "RequireEncryption", _chimesdkvoiceRequireEncryption); err != nil {
			log.Errorf("invalid --require-encryption: %s", err.Error())
			return
		}
	}
	if len(_chimesdkvoiceAwsRegion) > 0 {
		if err := assignInputField(input, "AwsRegion", _chimesdkvoiceAwsRegion); err != nil {
			log.Errorf("invalid --aws-region: %s", err.Error())
			return
		}
	}
	if len(_chimesdkvoiceIntegrationType) > 0 {
		if err := assignInputField(input, "IntegrationType", _chimesdkvoiceIntegrationType); err != nil {
			log.Errorf("invalid --integration-type: %s", err.Error())
			return
		}
	}
	if len(_chimesdkvoiceNetworkType) > 0 {
		if err := assignInputField(input, "NetworkType", _chimesdkvoiceNetworkType); err != nil {
			log.Errorf("invalid --network-type: %s", err.Error())
			return
		}
	}
	if len(_chimesdkvoiceTags) > 0 {
		if err := assignInputField(input, "Tags", _chimesdkvoiceTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateVoiceConnector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon Chime SDK Voice Connector group under the administrator's AWS
// account. You can associate Amazon Chime SDK Voice Connectors with the Voice
// Connector group by including VoiceConnectorItems in the request.
//
// You can include Voice Connectors from different AWS Regions in your group. This
// creates a fault tolerant mechanism for fallback in case of availability events.
func chimesdkvoice_CreateVoiceConnectorGroup(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.CreateVoiceConnectorGroupInput{
		// Name: *string, // Required
	}

	if len(_chimesdkvoiceName) > 0 {
		input.Name = aws.String(_chimesdkvoiceName)
	}
	if len(_chimesdkvoiceVoiceConnectorItems) > 0 {
		if err := assignInputField(input, "VoiceConnectorItems", _chimesdkvoiceVoiceConnectorItems); err != nil {
			log.Errorf("invalid --voice-connector-items: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateVoiceConnectorGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a voice profile, which consists of an enrolled user and their latest
// voice print.
//
// Before creating any voice profiles, you must provide all notices and obtain all
// consents from the speaker as required under applicable privacy and biometrics
// laws, and as required under the [AWS service terms]for the Amazon Chime SDK.
//
// For more information about voice profiles and voice analytics, see [Using Amazon Chime SDK Voice Analytics] in the
// Amazon Chime SDK Developer Guide.
//
// [AWS service terms]: https://aws.amazon.com/service-terms/
// [Using Amazon Chime SDK Voice Analytics]: https://docs.aws.amazon.com/chime-sdk/latest/dg/pstn-voice-analytics.html
func chimesdkvoice_CreateVoiceProfile(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.CreateVoiceProfileInput{
		// SpeakerSearchTaskId: *string, // Required
	}

	if len(_chimesdkvoiceSpeakerSearchTaskId) > 0 {
		input.SpeakerSearchTaskId = aws.String(_chimesdkvoiceSpeakerSearchTaskId)
	}

	if resp, err := client.CreateVoiceProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a voice profile domain, a collection of voice profiles, their voice
// prints, and encrypted enrollment audio.
//
// Before creating any voice profiles, you must provide all notices and obtain all
// consents from the speaker as required under applicable privacy and biometrics
// laws, and as required under the [AWS service terms]for the Amazon Chime SDK.
//
// For more information about voice profile domains, see [Using Amazon Chime SDK Voice Analytics] in the Amazon Chime SDK
// Developer Guide.
//
// [AWS service terms]: https://aws.amazon.com/service-terms/
// [Using Amazon Chime SDK Voice Analytics]: https://docs.aws.amazon.com/chime-sdk/latest/dg/pstn-voice-analytics.html
func chimesdkvoice_CreateVoiceProfileDomain(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.CreateVoiceProfileDomainInput{
		// Name: *string, // Required
		// ServerSideEncryptionConfiguration: *types.ServerSideEncryptionConfiguration, // Required
	}

	if len(_chimesdkvoiceName) > 0 {
		input.Name = aws.String(_chimesdkvoiceName)
	}
	if len(_chimesdkvoiceServerSideEncryptionConfiguration) > 0 {
		if err := assignInputField(input, "ServerSideEncryptionConfiguration", _chimesdkvoiceServerSideEncryptionConfiguration); err != nil {
			log.Errorf("invalid --server-side-encryption-configuration: %s", err.Error())
			return
		}
	}
	if len(_chimesdkvoiceClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_chimesdkvoiceClientRequestToken)
	}
	if len(_chimesdkvoiceDescription) > 0 {
		input.Description = aws.String(_chimesdkvoiceDescription)
	}
	if len(_chimesdkvoiceTags) > 0 {
		if err := assignInputField(input, "Tags", _chimesdkvoiceTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateVoiceProfileDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Moves the specified phone number into the Deletion queue. A phone number must
// be disassociated from any users or Amazon Chime SDK Voice Connectors before it
// can be deleted.
//
// Deleted phone numbers remain in the Deletion queue queue for 7 days before they
// are deleted permanently.
func chimesdkvoice_DeletePhoneNumber(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.DeletePhoneNumberInput{
		// PhoneNumberId: *string, // Required
	}

	if len(_chimesdkvoicePhoneNumberId) > 0 {
		input.PhoneNumberId = aws.String(_chimesdkvoicePhoneNumberId)
	}

	if resp, err := client.DeletePhoneNumber(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified proxy session from the specified Amazon Chime SDK Voice
// Connector.
func chimesdkvoice_DeleteProxySession(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.DeleteProxySessionInput{
		// ProxySessionId: *string, // Required
		// VoiceConnectorId: *string, // Required
	}

	if len(_chimesdkvoiceProxySessionId) > 0 {
		input.ProxySessionId = aws.String(_chimesdkvoiceProxySessionId)
	}
	if len(_chimesdkvoiceVoiceConnectorId) > 0 {
		input.VoiceConnectorId = aws.String(_chimesdkvoiceVoiceConnectorId)
	}

	if resp, err := client.DeleteProxySession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a SIP media application.
func chimesdkvoice_DeleteSipMediaApplication(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.DeleteSipMediaApplicationInput{
		// SipMediaApplicationId: *string, // Required
	}

	if len(_chimesdkvoiceSipMediaApplicationId) > 0 {
		input.SipMediaApplicationId = aws.String(_chimesdkvoiceSipMediaApplicationId)
	}

	if resp, err := client.DeleteSipMediaApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a SIP rule.
func chimesdkvoice_DeleteSipRule(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.DeleteSipRuleInput{
		// SipRuleId: *string, // Required
	}

	if len(_chimesdkvoiceSipRuleId) > 0 {
		input.SipRuleId = aws.String(_chimesdkvoiceSipRuleId)
	}

	if resp, err := client.DeleteSipRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon Chime SDK Voice Connector. Any phone numbers associated with
// the Amazon Chime SDK Voice Connector must be disassociated from it before it can
// be deleted.
func chimesdkvoice_DeleteVoiceConnector(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.DeleteVoiceConnectorInput{
		// VoiceConnectorId: *string, // Required
	}

	if len(_chimesdkvoiceVoiceConnectorId) > 0 {
		input.VoiceConnectorId = aws.String(_chimesdkvoiceVoiceConnectorId)
	}

	if resp, err := client.DeleteVoiceConnector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the emergency calling details from the specified Amazon Chime SDK Voice
// Connector.
func chimesdkvoice_DeleteVoiceConnectorEmergencyCallingConfiguration(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.DeleteVoiceConnectorEmergencyCallingConfigurationInput{
		// VoiceConnectorId: *string, // Required
	}

	if len(_chimesdkvoiceVoiceConnectorId) > 0 {
		input.VoiceConnectorId = aws.String(_chimesdkvoiceVoiceConnectorId)
	}

	if resp, err := client.DeleteVoiceConnectorEmergencyCallingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the external systems configuration for a Voice Connector.
func chimesdkvoice_DeleteVoiceConnectorExternalSystemsConfiguration(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.DeleteVoiceConnectorExternalSystemsConfigurationInput{
		// VoiceConnectorId: *string, // Required
	}

	if len(_chimesdkvoiceVoiceConnectorId) > 0 {
		input.VoiceConnectorId = aws.String(_chimesdkvoiceVoiceConnectorId)
	}

	if resp, err := client.DeleteVoiceConnectorExternalSystemsConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon Chime SDK Voice Connector group. Any VoiceConnectorItems and
// phone numbers associated with the group must be removed before it can be
// deleted.
func chimesdkvoice_DeleteVoiceConnectorGroup(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.DeleteVoiceConnectorGroupInput{
		// VoiceConnectorGroupId: *string, // Required
	}

	if len(_chimesdkvoiceVoiceConnectorGroupId) > 0 {
		input.VoiceConnectorGroupId = aws.String(_chimesdkvoiceVoiceConnectorGroupId)
	}

	if resp, err := client.DeleteVoiceConnectorGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the origination settings for the specified Amazon Chime SDK Voice
// Connector.
//
// If emergency calling is configured for the Voice Connector, it must be deleted
// prior to deleting the origination settings.
func chimesdkvoice_DeleteVoiceConnectorOrigination(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.DeleteVoiceConnectorOriginationInput{
		// VoiceConnectorId: *string, // Required
	}

	if len(_chimesdkvoiceVoiceConnectorId) > 0 {
		input.VoiceConnectorId = aws.String(_chimesdkvoiceVoiceConnectorId)
	}

	if resp, err := client.DeleteVoiceConnectorOrigination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the proxy configuration from the specified Amazon Chime SDK Voice
// Connector.
func chimesdkvoice_DeleteVoiceConnectorProxy(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.DeleteVoiceConnectorProxyInput{
		// VoiceConnectorId: *string, // Required
	}

	if len(_chimesdkvoiceVoiceConnectorId) > 0 {
		input.VoiceConnectorId = aws.String(_chimesdkvoiceVoiceConnectorId)
	}

	if resp, err := client.DeleteVoiceConnectorProxy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Voice Connector's streaming configuration.
func chimesdkvoice_DeleteVoiceConnectorStreamingConfiguration(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.DeleteVoiceConnectorStreamingConfigurationInput{
		// VoiceConnectorId: *string, // Required
	}

	if len(_chimesdkvoiceVoiceConnectorId) > 0 {
		input.VoiceConnectorId = aws.String(_chimesdkvoiceVoiceConnectorId)
	}

	if resp, err := client.DeleteVoiceConnectorStreamingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the termination settings for the specified Amazon Chime SDK Voice
// Connector.
//
// If emergency calling is configured for the Voice Connector, it must be deleted
// prior to deleting the termination settings.
func chimesdkvoice_DeleteVoiceConnectorTermination(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.DeleteVoiceConnectorTerminationInput{
		// VoiceConnectorId: *string, // Required
	}

	if len(_chimesdkvoiceVoiceConnectorId) > 0 {
		input.VoiceConnectorId = aws.String(_chimesdkvoiceVoiceConnectorId)
	}

	if resp, err := client.DeleteVoiceConnectorTermination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified SIP credentials used by your equipment to authenticate
// during call termination.
func chimesdkvoice_DeleteVoiceConnectorTerminationCredentials(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.DeleteVoiceConnectorTerminationCredentialsInput{
		// Usernames: []string, // Required
		// VoiceConnectorId: *string, // Required
	}

	if len(_chimesdkvoiceUsernames) > 0 {
		input.Usernames = append([]string(nil), _chimesdkvoiceUsernames...)
	}
	if len(_chimesdkvoiceVoiceConnectorId) > 0 {
		input.VoiceConnectorId = aws.String(_chimesdkvoiceVoiceConnectorId)
	}

	if resp, err := client.DeleteVoiceConnectorTerminationCredentials(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a voice profile, including its voice print and enrollment data.
// WARNING: This action is not reversible.
func chimesdkvoice_DeleteVoiceProfile(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.DeleteVoiceProfileInput{
		// VoiceProfileId: *string, // Required
	}

	if len(_chimesdkvoiceVoiceProfileId) > 0 {
		input.VoiceProfileId = aws.String(_chimesdkvoiceVoiceProfileId)
	}

	if resp, err := client.DeleteVoiceProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes all voice profiles in the domain. WARNING: This action is not
// reversible.
func chimesdkvoice_DeleteVoiceProfileDomain(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.DeleteVoiceProfileDomainInput{
		// VoiceProfileDomainId: *string, // Required
	}

	if len(_chimesdkvoiceVoiceProfileDomainId) > 0 {
		input.VoiceProfileDomainId = aws.String(_chimesdkvoiceVoiceProfileDomainId)
	}

	if resp, err := client.DeleteVoiceProfileDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates the specified phone numbers from the specified Amazon Chime SDK
// Voice Connector.
func chimesdkvoice_DisassociatePhoneNumbersFromVoiceConnector(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.DisassociatePhoneNumbersFromVoiceConnectorInput{
		// E164PhoneNumbers: []string, // Required
		// VoiceConnectorId: *string, // Required
	}

	if len(_chimesdkvoiceE164PhoneNumbers) > 0 {
		input.E164PhoneNumbers = append([]string(nil), _chimesdkvoiceE164PhoneNumbers...)
	}
	if len(_chimesdkvoiceVoiceConnectorId) > 0 {
		input.VoiceConnectorId = aws.String(_chimesdkvoiceVoiceConnectorId)
	}

	if resp, err := client.DisassociatePhoneNumbersFromVoiceConnector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates the specified phone numbers from the specified Amazon Chime SDK
// Voice Connector group.
func chimesdkvoice_DisassociatePhoneNumbersFromVoiceConnectorGroup(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.DisassociatePhoneNumbersFromVoiceConnectorGroupInput{
		// E164PhoneNumbers: []string, // Required
		// VoiceConnectorGroupId: *string, // Required
	}

	if len(_chimesdkvoiceE164PhoneNumbers) > 0 {
		input.E164PhoneNumbers = append([]string(nil), _chimesdkvoiceE164PhoneNumbers...)
	}
	if len(_chimesdkvoiceVoiceConnectorGroupId) > 0 {
		input.VoiceConnectorGroupId = aws.String(_chimesdkvoiceVoiceConnectorGroupId)
	}

	if resp, err := client.DisassociatePhoneNumbersFromVoiceConnectorGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the global settings for the Amazon Chime SDK Voice Connectors in an
// AWS account.
func chimesdkvoice_GetGlobalSettings(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.GetGlobalSettingsInput{}

	if resp, err := client.GetGlobalSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details for the specified phone number ID, such as associations,
// capabilities, and product type.
func chimesdkvoice_GetPhoneNumber(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.GetPhoneNumberInput{
		// PhoneNumberId: *string, // Required
	}

	if len(_chimesdkvoicePhoneNumberId) > 0 {
		input.PhoneNumberId = aws.String(_chimesdkvoicePhoneNumberId)
	}

	if resp, err := client.GetPhoneNumber(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details for the specified phone number order, such as the order
// creation timestamp, phone numbers in E.164 format, product type, and order
// status.
func chimesdkvoice_GetPhoneNumberOrder(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.GetPhoneNumberOrderInput{
		// PhoneNumberOrderId: *string, // Required
	}

	if len(_chimesdkvoicePhoneNumberOrderId) > 0 {
		input.PhoneNumberOrderId = aws.String(_chimesdkvoicePhoneNumberOrderId)
	}

	if resp, err := client.GetPhoneNumberOrder(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the phone number settings for the administrator's AWS account, such
// as the default outbound calling name.
func chimesdkvoice_GetPhoneNumberSettings(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.GetPhoneNumberSettingsInput{}

	if resp, err := client.GetPhoneNumberSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the specified proxy session details for the specified Amazon Chime
// SDK Voice Connector.
func chimesdkvoice_GetProxySession(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.GetProxySessionInput{
		// ProxySessionId: *string, // Required
		// VoiceConnectorId: *string, // Required
	}

	if len(_chimesdkvoiceProxySessionId) > 0 {
		input.ProxySessionId = aws.String(_chimesdkvoiceProxySessionId)
	}
	if len(_chimesdkvoiceVoiceConnectorId) > 0 {
		input.VoiceConnectorId = aws.String(_chimesdkvoiceVoiceConnectorId)
	}

	if resp, err := client.GetProxySession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the information for a SIP media application, including name, AWS
// Region, and endpoints.
func chimesdkvoice_GetSipMediaApplication(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.GetSipMediaApplicationInput{
		// SipMediaApplicationId: *string, // Required
	}

	if len(_chimesdkvoiceSipMediaApplicationId) > 0 {
		input.SipMediaApplicationId = aws.String(_chimesdkvoiceSipMediaApplicationId)
	}

	if resp, err := client.GetSipMediaApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the Alexa Skill configuration for the SIP media application.
// Due to changes made by the Amazon Alexa service, this API is no longer
// available for use. For more information, refer to the [Alexa Smart Properties]page.
//
// Deprecated: Due to changes made by the Amazon Alexa service, this API is no
// longer available for use. For more information, refer to the Alexa Smart
// Properties page(https://developer.amazon.com/en-US/alexa/alexasmartproperties).
//
// [Alexa Smart Properties]: https://developer.amazon.com/en-US/alexa/alexasmartproperties
func chimesdkvoice_GetSipMediaApplicationAlexaSkillConfiguration(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.GetSipMediaApplicationAlexaSkillConfigurationInput{
		// SipMediaApplicationId: *string, // Required
	}

	if len(_chimesdkvoiceSipMediaApplicationId) > 0 {
		input.SipMediaApplicationId = aws.String(_chimesdkvoiceSipMediaApplicationId)
	}

	if resp, err := client.GetSipMediaApplicationAlexaSkillConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the logging configuration for the specified SIP media application.
func chimesdkvoice_GetSipMediaApplicationLoggingConfiguration(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.GetSipMediaApplicationLoggingConfigurationInput{
		// SipMediaApplicationId: *string, // Required
	}

	if len(_chimesdkvoiceSipMediaApplicationId) > 0 {
		input.SipMediaApplicationId = aws.String(_chimesdkvoiceSipMediaApplicationId)
	}

	if resp, err := client.GetSipMediaApplicationLoggingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the details of a SIP rule, such as the rule ID, name, triggers, and
// target endpoints.
func chimesdkvoice_GetSipRule(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.GetSipRuleInput{
		// SipRuleId: *string, // Required
	}

	if len(_chimesdkvoiceSipRuleId) > 0 {
		input.SipRuleId = aws.String(_chimesdkvoiceSipRuleId)
	}

	if resp, err := client.GetSipRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the details of the specified speaker search task.
func chimesdkvoice_GetSpeakerSearchTask(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.GetSpeakerSearchTaskInput{
		// SpeakerSearchTaskId: *string, // Required
		// VoiceConnectorId: *string, // Required
	}

	if len(_chimesdkvoiceSpeakerSearchTaskId) > 0 {
		input.SpeakerSearchTaskId = aws.String(_chimesdkvoiceSpeakerSearchTaskId)
	}
	if len(_chimesdkvoiceVoiceConnectorId) > 0 {
		input.VoiceConnectorId = aws.String(_chimesdkvoiceVoiceConnectorId)
	}

	if resp, err := client.GetSpeakerSearchTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details for the specified Amazon Chime SDK Voice Connector, such as
// timestamps,name, outbound host, and encryption requirements.
func chimesdkvoice_GetVoiceConnector(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.GetVoiceConnectorInput{
		// VoiceConnectorId: *string, // Required
	}

	if len(_chimesdkvoiceVoiceConnectorId) > 0 {
		input.VoiceConnectorId = aws.String(_chimesdkvoiceVoiceConnectorId)
	}

	if resp, err := client.GetVoiceConnector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the emergency calling configuration details for the specified Voice
// Connector.
func chimesdkvoice_GetVoiceConnectorEmergencyCallingConfiguration(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.GetVoiceConnectorEmergencyCallingConfigurationInput{
		// VoiceConnectorId: *string, // Required
	}

	if len(_chimesdkvoiceVoiceConnectorId) > 0 {
		input.VoiceConnectorId = aws.String(_chimesdkvoiceVoiceConnectorId)
	}

	if resp, err := client.GetVoiceConnectorEmergencyCallingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about an external systems configuration for a Voice Connector.
func chimesdkvoice_GetVoiceConnectorExternalSystemsConfiguration(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.GetVoiceConnectorExternalSystemsConfigurationInput{
		// VoiceConnectorId: *string, // Required
	}

	if len(_chimesdkvoiceVoiceConnectorId) > 0 {
		input.VoiceConnectorId = aws.String(_chimesdkvoiceVoiceConnectorId)
	}

	if resp, err := client.GetVoiceConnectorExternalSystemsConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details for the specified Amazon Chime SDK Voice Connector group,
// such as timestamps,name, and associated VoiceConnectorItems .
func chimesdkvoice_GetVoiceConnectorGroup(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.GetVoiceConnectorGroupInput{
		// VoiceConnectorGroupId: *string, // Required
	}

	if len(_chimesdkvoiceVoiceConnectorGroupId) > 0 {
		input.VoiceConnectorGroupId = aws.String(_chimesdkvoiceVoiceConnectorGroupId)
	}

	if resp, err := client.GetVoiceConnectorGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the logging configuration settings for the specified Voice Connector.
// Shows whether SIP message logs are enabled for sending to Amazon CloudWatch
// Logs.
func chimesdkvoice_GetVoiceConnectorLoggingConfiguration(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.GetVoiceConnectorLoggingConfigurationInput{
		// VoiceConnectorId: *string, // Required
	}

	if len(_chimesdkvoiceVoiceConnectorId) > 0 {
		input.VoiceConnectorId = aws.String(_chimesdkvoiceVoiceConnectorId)
	}

	if resp, err := client.GetVoiceConnectorLoggingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the origination settings for the specified Voice Connector.
func chimesdkvoice_GetVoiceConnectorOrigination(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.GetVoiceConnectorOriginationInput{
		// VoiceConnectorId: *string, // Required
	}

	if len(_chimesdkvoiceVoiceConnectorId) > 0 {
		input.VoiceConnectorId = aws.String(_chimesdkvoiceVoiceConnectorId)
	}

	if resp, err := client.GetVoiceConnectorOrigination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the proxy configuration details for the specified Amazon Chime SDK
// Voice Connector.
func chimesdkvoice_GetVoiceConnectorProxy(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.GetVoiceConnectorProxyInput{
		// VoiceConnectorId: *string, // Required
	}

	if len(_chimesdkvoiceVoiceConnectorId) > 0 {
		input.VoiceConnectorId = aws.String(_chimesdkvoiceVoiceConnectorId)
	}

	if resp, err := client.GetVoiceConnectorProxy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the streaming configuration details for the specified Amazon Chime
// SDK Voice Connector. Shows whether media streaming is enabled for sending to
// Amazon Kinesis. It also shows the retention period, in hours, for the Amazon
// Kinesis data.
func chimesdkvoice_GetVoiceConnectorStreamingConfiguration(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.GetVoiceConnectorStreamingConfigurationInput{
		// VoiceConnectorId: *string, // Required
	}

	if len(_chimesdkvoiceVoiceConnectorId) > 0 {
		input.VoiceConnectorId = aws.String(_chimesdkvoiceVoiceConnectorId)
	}

	if resp, err := client.GetVoiceConnectorStreamingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the termination setting details for the specified Voice Connector.
func chimesdkvoice_GetVoiceConnectorTermination(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.GetVoiceConnectorTerminationInput{
		// VoiceConnectorId: *string, // Required
	}

	if len(_chimesdkvoiceVoiceConnectorId) > 0 {
		input.VoiceConnectorId = aws.String(_chimesdkvoiceVoiceConnectorId)
	}

	if resp, err := client.GetVoiceConnectorTermination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the last time a SIP OPTIONS ping was received from
// your SIP infrastructure for the specified Amazon Chime SDK Voice Connector.
func chimesdkvoice_GetVoiceConnectorTerminationHealth(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.GetVoiceConnectorTerminationHealthInput{
		// VoiceConnectorId: *string, // Required
	}

	if len(_chimesdkvoiceVoiceConnectorId) > 0 {
		input.VoiceConnectorId = aws.String(_chimesdkvoiceVoiceConnectorId)
	}

	if resp, err := client.GetVoiceConnectorTerminationHealth(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the details of the specified voice profile.
func chimesdkvoice_GetVoiceProfile(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.GetVoiceProfileInput{
		// VoiceProfileId: *string, // Required
	}

	if len(_chimesdkvoiceVoiceProfileId) > 0 {
		input.VoiceProfileId = aws.String(_chimesdkvoiceVoiceProfileId)
	}

	if resp, err := client.GetVoiceProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the details of the specified voice profile domain.
func chimesdkvoice_GetVoiceProfileDomain(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.GetVoiceProfileDomainInput{
		// VoiceProfileDomainId: *string, // Required
	}

	if len(_chimesdkvoiceVoiceProfileDomainId) > 0 {
		input.VoiceProfileDomainId = aws.String(_chimesdkvoiceVoiceProfileDomainId)
	}

	if resp, err := client.GetVoiceProfileDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the details of a voice tone analysis task.
func chimesdkvoice_GetVoiceToneAnalysisTask(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.GetVoiceToneAnalysisTaskInput{
		// IsCaller: *bool, // Required
		// VoiceConnectorId: *string, // Required
		// VoiceToneAnalysisTaskId: *string, // Required
	}

	if len(_chimesdkvoiceIsCaller) > 0 {
		if err := assignInputField(input, "IsCaller", _chimesdkvoiceIsCaller); err != nil {
			log.Errorf("invalid --is-caller: %s", err.Error())
			return
		}
	}
	if len(_chimesdkvoiceVoiceConnectorId) > 0 {
		input.VoiceConnectorId = aws.String(_chimesdkvoiceVoiceConnectorId)
	}
	if len(_chimesdkvoiceVoiceToneAnalysisTaskId) > 0 {
		input.VoiceToneAnalysisTaskId = aws.String(_chimesdkvoiceVoiceToneAnalysisTaskId)
	}

	if resp, err := client.GetVoiceToneAnalysisTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the available AWS Regions in which you can create an Amazon Chime SDK
// Voice Connector.
func chimesdkvoice_ListAvailableVoiceConnectorRegions(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.ListAvailableVoiceConnectorRegionsInput{}

	if resp, err := client.ListAvailableVoiceConnectorRegions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the phone numbers for an administrator's Amazon Chime SDK account.
func chimesdkvoice_ListPhoneNumberOrders(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.ListPhoneNumberOrdersInput{}

	if len(_chimesdkvoiceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _chimesdkvoiceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_chimesdkvoiceNextToken) > 0 {
		input.NextToken = aws.String(_chimesdkvoiceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPhoneNumberOrders(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*chimesdkvoice.ListPhoneNumberOrdersOutput
	p := chimesdkvoice.NewListPhoneNumberOrdersPaginator(client, input)
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

// Lists the phone numbers for the specified Amazon Chime SDK account, Amazon
// Chime SDK user, Amazon Chime SDK Voice Connector, or Amazon Chime SDK Voice
// Connector group.
func chimesdkvoice_ListPhoneNumbers(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.ListPhoneNumbersInput{}

	if len(_chimesdkvoiceFilterName) > 0 {
		if err := assignInputField(input, "FilterName", _chimesdkvoiceFilterName); err != nil {
			log.Errorf("invalid --filter-name: %s", err.Error())
			return
		}
	}
	if len(_chimesdkvoiceFilterValue) > 0 {
		input.FilterValue = aws.String(_chimesdkvoiceFilterValue)
	}
	if len(_chimesdkvoiceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _chimesdkvoiceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_chimesdkvoiceNextToken) > 0 {
		input.NextToken = aws.String(_chimesdkvoiceNextToken)
	}
	if len(_chimesdkvoiceProductType) > 0 {
		if err := assignInputField(input, "ProductType", _chimesdkvoiceProductType); err != nil {
			log.Errorf("invalid --product-type: %s", err.Error())
			return
		}
	}
	if len(_chimesdkvoiceStatus) > 0 {
		input.Status = aws.String(_chimesdkvoiceStatus)
	}

	if disablePaginator() {
		if resp, err := client.ListPhoneNumbers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*chimesdkvoice.ListPhoneNumbersOutput
	p := chimesdkvoice.NewListPhoneNumbersPaginator(client, input)
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

// Lists the proxy sessions for the specified Amazon Chime SDK Voice Connector.
func chimesdkvoice_ListProxySessions(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.ListProxySessionsInput{
		// VoiceConnectorId: *string, // Required
	}

	if len(_chimesdkvoiceVoiceConnectorId) > 0 {
		input.VoiceConnectorId = aws.String(_chimesdkvoiceVoiceConnectorId)
	}
	if len(_chimesdkvoiceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _chimesdkvoiceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_chimesdkvoiceNextToken) > 0 {
		input.NextToken = aws.String(_chimesdkvoiceNextToken)
	}
	if len(_chimesdkvoiceStatus) > 0 {
		if err := assignInputField(input, "Status", _chimesdkvoiceStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListProxySessions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*chimesdkvoice.ListProxySessionsOutput
	p := chimesdkvoice.NewListProxySessionsPaginator(client, input)
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

// Lists the SIP media applications under the administrator's AWS account.
func chimesdkvoice_ListSipMediaApplications(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.ListSipMediaApplicationsInput{}

	if len(_chimesdkvoiceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _chimesdkvoiceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_chimesdkvoiceNextToken) > 0 {
		input.NextToken = aws.String(_chimesdkvoiceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSipMediaApplications(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*chimesdkvoice.ListSipMediaApplicationsOutput
	p := chimesdkvoice.NewListSipMediaApplicationsPaginator(client, input)
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

// Lists the SIP rules under the administrator's AWS account.
func chimesdkvoice_ListSipRules(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.ListSipRulesInput{}

	if len(_chimesdkvoiceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _chimesdkvoiceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_chimesdkvoiceNextToken) > 0 {
		input.NextToken = aws.String(_chimesdkvoiceNextToken)
	}
	if len(_chimesdkvoiceSipMediaApplicationId) > 0 {
		input.SipMediaApplicationId = aws.String(_chimesdkvoiceSipMediaApplicationId)
	}

	if disablePaginator() {
		if resp, err := client.ListSipRules(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*chimesdkvoice.ListSipRulesOutput
	p := chimesdkvoice.NewListSipRulesPaginator(client, input)
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

// Lists the countries that you can order phone numbers from.
func chimesdkvoice_ListSupportedPhoneNumberCountries(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.ListSupportedPhoneNumberCountriesInput{
		// ProductType: types.PhoneNumberProductType, // Required
	}

	if len(_chimesdkvoiceProductType) > 0 {
		if err := assignInputField(input, "ProductType", _chimesdkvoiceProductType); err != nil {
			log.Errorf("invalid --product-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListSupportedPhoneNumberCountries(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of the tags in a given resource.
func chimesdkvoice_ListTagsForResource(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.ListTagsForResourceInput{
		// ResourceARN: *string, // Required
	}

	if len(_chimesdkvoiceResourceARN) > 0 {
		input.ResourceARN = aws.String(_chimesdkvoiceResourceARN)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the Amazon Chime SDK Voice Connector groups in the administrator's AWS
// account.
func chimesdkvoice_ListVoiceConnectorGroups(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.ListVoiceConnectorGroupsInput{}

	if len(_chimesdkvoiceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _chimesdkvoiceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_chimesdkvoiceNextToken) > 0 {
		input.NextToken = aws.String(_chimesdkvoiceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListVoiceConnectorGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*chimesdkvoice.ListVoiceConnectorGroupsOutput
	p := chimesdkvoice.NewListVoiceConnectorGroupsPaginator(client, input)
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

// Lists the SIP credentials for the specified Amazon Chime SDK Voice Connector.
func chimesdkvoice_ListVoiceConnectorTerminationCredentials(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.ListVoiceConnectorTerminationCredentialsInput{
		// VoiceConnectorId: *string, // Required
	}

	if len(_chimesdkvoiceVoiceConnectorId) > 0 {
		input.VoiceConnectorId = aws.String(_chimesdkvoiceVoiceConnectorId)
	}

	if resp, err := client.ListVoiceConnectorTerminationCredentials(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the Amazon Chime SDK Voice Connectors in the administrators AWS account.
func chimesdkvoice_ListVoiceConnectors(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.ListVoiceConnectorsInput{}

	if len(_chimesdkvoiceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _chimesdkvoiceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_chimesdkvoiceNextToken) > 0 {
		input.NextToken = aws.String(_chimesdkvoiceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListVoiceConnectors(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*chimesdkvoice.ListVoiceConnectorsOutput
	p := chimesdkvoice.NewListVoiceConnectorsPaginator(client, input)
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

// Lists the specified voice profile domains in the administrator's AWS account.
func chimesdkvoice_ListVoiceProfileDomains(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.ListVoiceProfileDomainsInput{}

	if len(_chimesdkvoiceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _chimesdkvoiceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_chimesdkvoiceNextToken) > 0 {
		input.NextToken = aws.String(_chimesdkvoiceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListVoiceProfileDomains(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*chimesdkvoice.ListVoiceProfileDomainsOutput
	p := chimesdkvoice.NewListVoiceProfileDomainsPaginator(client, input)
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

// Lists the voice profiles in a voice profile domain.
func chimesdkvoice_ListVoiceProfiles(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.ListVoiceProfilesInput{
		// VoiceProfileDomainId: *string, // Required
	}

	if len(_chimesdkvoiceVoiceProfileDomainId) > 0 {
		input.VoiceProfileDomainId = aws.String(_chimesdkvoiceVoiceProfileDomainId)
	}
	if len(_chimesdkvoiceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _chimesdkvoiceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_chimesdkvoiceNextToken) > 0 {
		input.NextToken = aws.String(_chimesdkvoiceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListVoiceProfiles(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*chimesdkvoice.ListVoiceProfilesOutput
	p := chimesdkvoice.NewListVoiceProfilesPaginator(client, input)
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

// Updates the Alexa Skill configuration for the SIP media application.
// Due to changes made by the Amazon Alexa service, this API is no longer
// available for use. For more information, refer to the [Alexa Smart Properties]page.
//
// Deprecated: Due to changes made by the Amazon Alexa service, this API is no
// longer available for use. For more information, refer to the Alexa Smart
// Properties page(https://developer.amazon.com/en-US/alexa/alexasmartproperties).
//
// [Alexa Smart Properties]: https://developer.amazon.com/en-US/alexa/alexasmartproperties
func chimesdkvoice_PutSipMediaApplicationAlexaSkillConfiguration(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.PutSipMediaApplicationAlexaSkillConfigurationInput{
		// SipMediaApplicationId: *string, // Required
	}

	if len(_chimesdkvoiceSipMediaApplicationId) > 0 {
		input.SipMediaApplicationId = aws.String(_chimesdkvoiceSipMediaApplicationId)
	}
	if len(_chimesdkvoiceSipMediaApplicationAlexaSkillConfiguration) > 0 {
		if err := assignInputField(input, "SipMediaApplicationAlexaSkillConfiguration", _chimesdkvoiceSipMediaApplicationAlexaSkillConfiguration); err != nil {
			log.Errorf("invalid --sip-media-application-alexa-skill-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutSipMediaApplicationAlexaSkillConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the logging configuration for the specified SIP media application.
func chimesdkvoice_PutSipMediaApplicationLoggingConfiguration(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.PutSipMediaApplicationLoggingConfigurationInput{
		// SipMediaApplicationId: *string, // Required
	}

	if len(_chimesdkvoiceSipMediaApplicationId) > 0 {
		input.SipMediaApplicationId = aws.String(_chimesdkvoiceSipMediaApplicationId)
	}
	if len(_chimesdkvoiceSipMediaApplicationLoggingConfiguration) > 0 {
		if err := assignInputField(input, "SipMediaApplicationLoggingConfiguration", _chimesdkvoiceSipMediaApplicationLoggingConfiguration); err != nil {
			log.Errorf("invalid --sip-media-application-logging-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutSipMediaApplicationLoggingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a Voice Connector's emergency calling configuration.
func chimesdkvoice_PutVoiceConnectorEmergencyCallingConfiguration(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.PutVoiceConnectorEmergencyCallingConfigurationInput{
		// EmergencyCallingConfiguration: *types.EmergencyCallingConfiguration, // Required
		// VoiceConnectorId: *string, // Required
	}

	if len(_chimesdkvoiceEmergencyCallingConfiguration) > 0 {
		if err := assignInputField(input, "EmergencyCallingConfiguration", _chimesdkvoiceEmergencyCallingConfiguration); err != nil {
			log.Errorf("invalid --emergency-calling-configuration: %s", err.Error())
			return
		}
	}
	if len(_chimesdkvoiceVoiceConnectorId) > 0 {
		input.VoiceConnectorId = aws.String(_chimesdkvoiceVoiceConnectorId)
	}

	if resp, err := client.PutVoiceConnectorEmergencyCallingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds an external systems configuration to a Voice Connector.
func chimesdkvoice_PutVoiceConnectorExternalSystemsConfiguration(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.PutVoiceConnectorExternalSystemsConfigurationInput{
		// VoiceConnectorId: *string, // Required
	}

	if len(_chimesdkvoiceVoiceConnectorId) > 0 {
		input.VoiceConnectorId = aws.String(_chimesdkvoiceVoiceConnectorId)
	}
	if len(_chimesdkvoiceContactCenterSystemTypes) > 0 {
		if err := assignInputField(input, "ContactCenterSystemTypes", _chimesdkvoiceContactCenterSystemTypes); err != nil {
			log.Errorf("invalid --contact-center-system-types: %s", err.Error())
			return
		}
	}
	if len(_chimesdkvoiceSessionBorderControllerTypes) > 0 {
		if err := assignInputField(input, "SessionBorderControllerTypes", _chimesdkvoiceSessionBorderControllerTypes); err != nil {
			log.Errorf("invalid --session-border-controller-types: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutVoiceConnectorExternalSystemsConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a Voice Connector's logging configuration.
func chimesdkvoice_PutVoiceConnectorLoggingConfiguration(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.PutVoiceConnectorLoggingConfigurationInput{
		// LoggingConfiguration: *types.LoggingConfiguration, // Required
		// VoiceConnectorId: *string, // Required
	}

	if len(_chimesdkvoiceLoggingConfiguration) > 0 {
		if err := assignInputField(input, "LoggingConfiguration", _chimesdkvoiceLoggingConfiguration); err != nil {
			log.Errorf("invalid --logging-configuration: %s", err.Error())
			return
		}
	}
	if len(_chimesdkvoiceVoiceConnectorId) > 0 {
		input.VoiceConnectorId = aws.String(_chimesdkvoiceVoiceConnectorId)
	}

	if resp, err := client.PutVoiceConnectorLoggingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a Voice Connector's origination settings.
func chimesdkvoice_PutVoiceConnectorOrigination(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.PutVoiceConnectorOriginationInput{
		// Origination: *types.Origination, // Required
		// VoiceConnectorId: *string, // Required
	}

	if len(_chimesdkvoiceOrigination) > 0 {
		if err := assignInputField(input, "Origination", _chimesdkvoiceOrigination); err != nil {
			log.Errorf("invalid --origination: %s", err.Error())
			return
		}
	}
	if len(_chimesdkvoiceVoiceConnectorId) > 0 {
		input.VoiceConnectorId = aws.String(_chimesdkvoiceVoiceConnectorId)
	}

	if resp, err := client.PutVoiceConnectorOrigination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Puts the specified proxy configuration to the specified Amazon Chime SDK Voice
// Connector.
func chimesdkvoice_PutVoiceConnectorProxy(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.PutVoiceConnectorProxyInput{
		// DefaultSessionExpiryMinutes: *int32, // Required
		// PhoneNumberPoolCountries: []string, // Required
		// VoiceConnectorId: *string, // Required
	}

	if len(_chimesdkvoiceDefaultSessionExpiryMinutes) > 0 {
		if err := assignInputField(input, "DefaultSessionExpiryMinutes", _chimesdkvoiceDefaultSessionExpiryMinutes); err != nil {
			log.Errorf("invalid --default-session-expiry-minutes: %s", err.Error())
			return
		}
	}
	if len(_chimesdkvoicePhoneNumberPoolCountries) > 0 {
		input.PhoneNumberPoolCountries = append([]string(nil), _chimesdkvoicePhoneNumberPoolCountries...)
	}
	if len(_chimesdkvoiceVoiceConnectorId) > 0 {
		input.VoiceConnectorId = aws.String(_chimesdkvoiceVoiceConnectorId)
	}
	if len(_chimesdkvoiceDisabled) > 0 {
		if err := assignInputField(input, "Disabled", _chimesdkvoiceDisabled); err != nil {
			log.Errorf("invalid --disabled: %s", err.Error())
			return
		}
	}
	if len(_chimesdkvoiceFallBackPhoneNumber) > 0 {
		input.FallBackPhoneNumber = aws.String(_chimesdkvoiceFallBackPhoneNumber)
	}

	if resp, err := client.PutVoiceConnectorProxy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a Voice Connector's streaming configuration settings.
func chimesdkvoice_PutVoiceConnectorStreamingConfiguration(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.PutVoiceConnectorStreamingConfigurationInput{
		// StreamingConfiguration: *types.StreamingConfiguration, // Required
		// VoiceConnectorId: *string, // Required
	}

	if len(_chimesdkvoiceStreamingConfiguration) > 0 {
		if err := assignInputField(input, "StreamingConfiguration", _chimesdkvoiceStreamingConfiguration); err != nil {
			log.Errorf("invalid --streaming-configuration: %s", err.Error())
			return
		}
	}
	if len(_chimesdkvoiceVoiceConnectorId) > 0 {
		input.VoiceConnectorId = aws.String(_chimesdkvoiceVoiceConnectorId)
	}

	if resp, err := client.PutVoiceConnectorStreamingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a Voice Connector's termination settings.
func chimesdkvoice_PutVoiceConnectorTermination(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.PutVoiceConnectorTerminationInput{
		// Termination: *types.Termination, // Required
		// VoiceConnectorId: *string, // Required
	}

	if len(_chimesdkvoiceTermination) > 0 {
		if err := assignInputField(input, "Termination", _chimesdkvoiceTermination); err != nil {
			log.Errorf("invalid --termination: %s", err.Error())
			return
		}
	}
	if len(_chimesdkvoiceVoiceConnectorId) > 0 {
		input.VoiceConnectorId = aws.String(_chimesdkvoiceVoiceConnectorId)
	}

	if resp, err := client.PutVoiceConnectorTermination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a Voice Connector's termination credentials.
func chimesdkvoice_PutVoiceConnectorTerminationCredentials(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.PutVoiceConnectorTerminationCredentialsInput{
		// VoiceConnectorId: *string, // Required
	}

	if len(_chimesdkvoiceVoiceConnectorId) > 0 {
		input.VoiceConnectorId = aws.String(_chimesdkvoiceVoiceConnectorId)
	}
	if len(_chimesdkvoiceCredentials) > 0 {
		if err := assignInputField(input, "Credentials", _chimesdkvoiceCredentials); err != nil {
			log.Errorf("invalid --credentials: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutVoiceConnectorTerminationCredentials(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Restores a deleted phone number.
func chimesdkvoice_RestorePhoneNumber(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.RestorePhoneNumberInput{
		// PhoneNumberId: *string, // Required
	}

	if len(_chimesdkvoicePhoneNumberId) > 0 {
		input.PhoneNumberId = aws.String(_chimesdkvoicePhoneNumberId)
	}

	if resp, err := client.RestorePhoneNumber(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Searches the provisioned phone numbers in an organization.
func chimesdkvoice_SearchAvailablePhoneNumbers(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.SearchAvailablePhoneNumbersInput{}

	if len(_chimesdkvoiceAreaCode) > 0 {
		input.AreaCode = aws.String(_chimesdkvoiceAreaCode)
	}
	if len(_chimesdkvoiceCity) > 0 {
		input.City = aws.String(_chimesdkvoiceCity)
	}
	if len(_chimesdkvoiceCountry) > 0 {
		input.Country = aws.String(_chimesdkvoiceCountry)
	}
	if len(_chimesdkvoiceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _chimesdkvoiceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_chimesdkvoiceNextToken) > 0 {
		input.NextToken = aws.String(_chimesdkvoiceNextToken)
	}
	if len(_chimesdkvoicePhoneNumberType) > 0 {
		if err := assignInputField(input, "PhoneNumberType", _chimesdkvoicePhoneNumberType); err != nil {
			log.Errorf("invalid --phone-number-type: %s", err.Error())
			return
		}
	}
	if len(_chimesdkvoiceState) > 0 {
		input.State = aws.String(_chimesdkvoiceState)
	}
	if len(_chimesdkvoiceTollFreePrefix) > 0 {
		input.TollFreePrefix = aws.String(_chimesdkvoiceTollFreePrefix)
	}

	if disablePaginator() {
		if resp, err := client.SearchAvailablePhoneNumbers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*chimesdkvoice.SearchAvailablePhoneNumbersOutput
	p := chimesdkvoice.NewSearchAvailablePhoneNumbersPaginator(client, input)
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

// Starts a speaker search task.
// Before starting any speaker search tasks, you must provide all notices and
// obtain all consents from the speaker as required under applicable privacy and
// biometrics laws, and as required under the [AWS service terms]for the Amazon Chime SDK.
//
// [AWS service terms]: https://aws.amazon.com/service-terms/
func chimesdkvoice_StartSpeakerSearchTask(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.StartSpeakerSearchTaskInput{
		// TransactionId: *string, // Required
		// VoiceConnectorId: *string, // Required
		// VoiceProfileDomainId: *string, // Required
	}

	if len(_chimesdkvoiceTransactionId) > 0 {
		input.TransactionId = aws.String(_chimesdkvoiceTransactionId)
	}
	if len(_chimesdkvoiceVoiceConnectorId) > 0 {
		input.VoiceConnectorId = aws.String(_chimesdkvoiceVoiceConnectorId)
	}
	if len(_chimesdkvoiceVoiceProfileDomainId) > 0 {
		input.VoiceProfileDomainId = aws.String(_chimesdkvoiceVoiceProfileDomainId)
	}
	if len(_chimesdkvoiceCallLeg) > 0 {
		if err := assignInputField(input, "CallLeg", _chimesdkvoiceCallLeg); err != nil {
			log.Errorf("invalid --call-leg: %s", err.Error())
			return
		}
	}
	if len(_chimesdkvoiceClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_chimesdkvoiceClientRequestToken)
	}

	if resp, err := client.StartSpeakerSearchTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a voice tone analysis task. For more information about voice tone
// analysis, see [Using Amazon Chime SDK voice analytics]in the Amazon Chime SDK Developer Guide.
//
// Before starting any voice tone analysis tasks, you must provide all notices and
// obtain all consents from the speaker as required under applicable privacy and
// biometrics laws, and as required under the [AWS service terms]for the Amazon Chime SDK.
//
// [Using Amazon Chime SDK voice analytics]: https://docs.aws.amazon.com/chime-sdk/latest/dg/pstn-voice-analytics.html
// [AWS service terms]: https://aws.amazon.com/service-terms/
func chimesdkvoice_StartVoiceToneAnalysisTask(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.StartVoiceToneAnalysisTaskInput{
		// LanguageCode: types.LanguageCode, // Required
		// TransactionId: *string, // Required
		// VoiceConnectorId: *string, // Required
	}

	if len(_chimesdkvoiceLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _chimesdkvoiceLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}
	if len(_chimesdkvoiceTransactionId) > 0 {
		input.TransactionId = aws.String(_chimesdkvoiceTransactionId)
	}
	if len(_chimesdkvoiceVoiceConnectorId) > 0 {
		input.VoiceConnectorId = aws.String(_chimesdkvoiceVoiceConnectorId)
	}
	if len(_chimesdkvoiceClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_chimesdkvoiceClientRequestToken)
	}

	if resp, err := client.StartVoiceToneAnalysisTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a speaker search task.
func chimesdkvoice_StopSpeakerSearchTask(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.StopSpeakerSearchTaskInput{
		// SpeakerSearchTaskId: *string, // Required
		// VoiceConnectorId: *string, // Required
	}

	if len(_chimesdkvoiceSpeakerSearchTaskId) > 0 {
		input.SpeakerSearchTaskId = aws.String(_chimesdkvoiceSpeakerSearchTaskId)
	}
	if len(_chimesdkvoiceVoiceConnectorId) > 0 {
		input.VoiceConnectorId = aws.String(_chimesdkvoiceVoiceConnectorId)
	}

	if resp, err := client.StopSpeakerSearchTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a voice tone analysis task.
func chimesdkvoice_StopVoiceToneAnalysisTask(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.StopVoiceToneAnalysisTaskInput{
		// VoiceConnectorId: *string, // Required
		// VoiceToneAnalysisTaskId: *string, // Required
	}

	if len(_chimesdkvoiceVoiceConnectorId) > 0 {
		input.VoiceConnectorId = aws.String(_chimesdkvoiceVoiceConnectorId)
	}
	if len(_chimesdkvoiceVoiceToneAnalysisTaskId) > 0 {
		input.VoiceToneAnalysisTaskId = aws.String(_chimesdkvoiceVoiceToneAnalysisTaskId)
	}

	if resp, err := client.StopVoiceToneAnalysisTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a tag to the specified resource.
func chimesdkvoice_TagResource(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.TagResourceInput{
		// ResourceARN: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_chimesdkvoiceResourceARN) > 0 {
		input.ResourceARN = aws.String(_chimesdkvoiceResourceARN)
	}
	if len(_chimesdkvoiceTags) > 0 {
		if err := assignInputField(input, "Tags", _chimesdkvoiceTags); err != nil {
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

// Removes tags from a resource.
func chimesdkvoice_UntagResource(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.UntagResourceInput{
		// ResourceARN: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_chimesdkvoiceResourceARN) > 0 {
		input.ResourceARN = aws.String(_chimesdkvoiceResourceARN)
	}
	if len(_chimesdkvoiceTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _chimesdkvoiceTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates global settings for the Amazon Chime SDK Voice Connectors in an AWS
// account.
func chimesdkvoice_UpdateGlobalSettings(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.UpdateGlobalSettingsInput{}

	if len(_chimesdkvoiceVoiceConnector) > 0 {
		if err := assignInputField(input, "VoiceConnector", _chimesdkvoiceVoiceConnector); err != nil {
			log.Errorf("invalid --voice-connector: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateGlobalSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates phone number details, such as product type, calling name, or phone
// number name for the specified phone number ID. You can update one phone number
// detail at a time. For example, you can update either the product type, calling
// name, or phone number name in one action.
//
// For numbers outside the U.S., you must use the Amazon Chime SDK SIP Media
// Application Dial-In product type.
//
// Updates to outbound calling names can take 72 hours to complete. Pending
// updates to outbound calling names must be complete before you can request
// another update.
func chimesdkvoice_UpdatePhoneNumber(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.UpdatePhoneNumberInput{
		// PhoneNumberId: *string, // Required
	}

	if len(_chimesdkvoicePhoneNumberId) > 0 {
		input.PhoneNumberId = aws.String(_chimesdkvoicePhoneNumberId)
	}
	if len(_chimesdkvoiceCallingName) > 0 {
		input.CallingName = aws.String(_chimesdkvoiceCallingName)
	}
	if len(_chimesdkvoiceName) > 0 {
		input.Name = aws.String(_chimesdkvoiceName)
	}
	if len(_chimesdkvoiceProductType) > 0 {
		if err := assignInputField(input, "ProductType", _chimesdkvoiceProductType); err != nil {
			log.Errorf("invalid --product-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdatePhoneNumber(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the phone number settings for the administrator's AWS account, such as
// the default outbound calling name. You can update the default outbound calling
// name once every seven days. Outbound calling names can take up to 72 hours to
// update.
func chimesdkvoice_UpdatePhoneNumberSettings(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.UpdatePhoneNumberSettingsInput{
		// CallingName: *string, // Required
	}

	if len(_chimesdkvoiceCallingName) > 0 {
		input.CallingName = aws.String(_chimesdkvoiceCallingName)
	}

	if resp, err := client.UpdatePhoneNumberSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified proxy session details, such as voice or SMS capabilities.
func chimesdkvoice_UpdateProxySession(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.UpdateProxySessionInput{
		// Capabilities: []types.Capability, // Required
		// ProxySessionId: *string, // Required
		// VoiceConnectorId: *string, // Required
	}

	if len(_chimesdkvoiceCapabilities) > 0 {
		if err := assignInputField(input, "Capabilities", _chimesdkvoiceCapabilities); err != nil {
			log.Errorf("invalid --capabilities: %s", err.Error())
			return
		}
	}
	if len(_chimesdkvoiceProxySessionId) > 0 {
		input.ProxySessionId = aws.String(_chimesdkvoiceProxySessionId)
	}
	if len(_chimesdkvoiceVoiceConnectorId) > 0 {
		input.VoiceConnectorId = aws.String(_chimesdkvoiceVoiceConnectorId)
	}
	if len(_chimesdkvoiceExpiryMinutes) > 0 {
		if err := assignInputField(input, "ExpiryMinutes", _chimesdkvoiceExpiryMinutes); err != nil {
			log.Errorf("invalid --expiry-minutes: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateProxySession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the details of the specified SIP media application.
func chimesdkvoice_UpdateSipMediaApplication(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.UpdateSipMediaApplicationInput{
		// SipMediaApplicationId: *string, // Required
	}

	if len(_chimesdkvoiceSipMediaApplicationId) > 0 {
		input.SipMediaApplicationId = aws.String(_chimesdkvoiceSipMediaApplicationId)
	}
	if len(_chimesdkvoiceEndpoints) > 0 {
		if err := assignInputField(input, "Endpoints", _chimesdkvoiceEndpoints); err != nil {
			log.Errorf("invalid --endpoints: %s", err.Error())
			return
		}
	}
	if len(_chimesdkvoiceName) > 0 {
		input.Name = aws.String(_chimesdkvoiceName)
	}

	if resp, err := client.UpdateSipMediaApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Invokes the AWS Lambda function associated with the SIP media application and
// transaction ID in an update request. The Lambda function can then return a new
// set of actions.
func chimesdkvoice_UpdateSipMediaApplicationCall(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.UpdateSipMediaApplicationCallInput{
		// Arguments: map[string]string, // Required
		// SipMediaApplicationId: *string, // Required
		// TransactionId: *string, // Required
	}

	if len(_chimesdkvoiceArguments) > 0 {
		if err := assignInputField(input, "Arguments", _chimesdkvoiceArguments); err != nil {
			log.Errorf("invalid --arguments: %s", err.Error())
			return
		}
	}
	if len(_chimesdkvoiceSipMediaApplicationId) > 0 {
		input.SipMediaApplicationId = aws.String(_chimesdkvoiceSipMediaApplicationId)
	}
	if len(_chimesdkvoiceTransactionId) > 0 {
		input.TransactionId = aws.String(_chimesdkvoiceTransactionId)
	}

	if resp, err := client.UpdateSipMediaApplicationCall(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the details of the specified SIP rule.
func chimesdkvoice_UpdateSipRule(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.UpdateSipRuleInput{
		// Name: *string, // Required
		// SipRuleId: *string, // Required
	}

	if len(_chimesdkvoiceName) > 0 {
		input.Name = aws.String(_chimesdkvoiceName)
	}
	if len(_chimesdkvoiceSipRuleId) > 0 {
		input.SipRuleId = aws.String(_chimesdkvoiceSipRuleId)
	}
	if len(_chimesdkvoiceDisabled) > 0 {
		if err := assignInputField(input, "Disabled", _chimesdkvoiceDisabled); err != nil {
			log.Errorf("invalid --disabled: %s", err.Error())
			return
		}
	}
	if len(_chimesdkvoiceTargetApplications) > 0 {
		if err := assignInputField(input, "TargetApplications", _chimesdkvoiceTargetApplications); err != nil {
			log.Errorf("invalid --target-applications: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSipRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the details for the specified Amazon Chime SDK Voice Connector.
func chimesdkvoice_UpdateVoiceConnector(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.UpdateVoiceConnectorInput{
		// Name: *string, // Required
		// RequireEncryption: *bool, // Required
		// VoiceConnectorId: *string, // Required
	}

	if len(_chimesdkvoiceName) > 0 {
		input.Name = aws.String(_chimesdkvoiceName)
	}
	if len(_chimesdkvoiceRequireEncryption) > 0 {
		if err := assignInputField(input, "RequireEncryption", _chimesdkvoiceRequireEncryption); err != nil {
			log.Errorf("invalid --require-encryption: %s", err.Error())
			return
		}
	}
	if len(_chimesdkvoiceVoiceConnectorId) > 0 {
		input.VoiceConnectorId = aws.String(_chimesdkvoiceVoiceConnectorId)
	}

	if resp, err := client.UpdateVoiceConnector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the settings for the specified Amazon Chime SDK Voice Connector group.
func chimesdkvoice_UpdateVoiceConnectorGroup(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.UpdateVoiceConnectorGroupInput{
		// Name: *string, // Required
		// VoiceConnectorGroupId: *string, // Required
		// VoiceConnectorItems: []types.VoiceConnectorItem, // Required
	}

	if len(_chimesdkvoiceName) > 0 {
		input.Name = aws.String(_chimesdkvoiceName)
	}
	if len(_chimesdkvoiceVoiceConnectorGroupId) > 0 {
		input.VoiceConnectorGroupId = aws.String(_chimesdkvoiceVoiceConnectorGroupId)
	}
	if len(_chimesdkvoiceVoiceConnectorItems) > 0 {
		if err := assignInputField(input, "VoiceConnectorItems", _chimesdkvoiceVoiceConnectorItems); err != nil {
			log.Errorf("invalid --voice-connector-items: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateVoiceConnectorGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified voice profile’s voice print and refreshes its expiration
// timestamp.
//
// As a condition of using this feature, you acknowledge that the collection, use,
// storage, and retention of your caller’s biometric identifiers and biometric
// information (“biometric data”) in the form of a digital voiceprint requires the
// caller’s informed consent via a written release. Such consent is required under
// various state laws, including biometrics laws in Illinois, Texas, Washington and
// other state privacy laws.
//
// You must provide a written release to each caller through a process that
// clearly reflects each caller’s informed consent before using Amazon Chime SDK
// Voice Insights service, as required under the terms of your agreement with AWS
// governing your use of the service.
func chimesdkvoice_UpdateVoiceProfile(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.UpdateVoiceProfileInput{
		// SpeakerSearchTaskId: *string, // Required
		// VoiceProfileId: *string, // Required
	}

	if len(_chimesdkvoiceSpeakerSearchTaskId) > 0 {
		input.SpeakerSearchTaskId = aws.String(_chimesdkvoiceSpeakerSearchTaskId)
	}
	if len(_chimesdkvoiceVoiceProfileId) > 0 {
		input.VoiceProfileId = aws.String(_chimesdkvoiceVoiceProfileId)
	}

	if resp, err := client.UpdateVoiceProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the settings for the specified voice profile domain.
func chimesdkvoice_UpdateVoiceProfileDomain(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.UpdateVoiceProfileDomainInput{
		// VoiceProfileDomainId: *string, // Required
	}

	if len(_chimesdkvoiceVoiceProfileDomainId) > 0 {
		input.VoiceProfileDomainId = aws.String(_chimesdkvoiceVoiceProfileDomainId)
	}
	if len(_chimesdkvoiceDescription) > 0 {
		input.Description = aws.String(_chimesdkvoiceDescription)
	}
	if len(_chimesdkvoiceName) > 0 {
		input.Name = aws.String(_chimesdkvoiceName)
	}

	if resp, err := client.UpdateVoiceProfileDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Validates an address to be used for 911 calls made with Amazon Chime SDK Voice
// Connectors. You can use validated addresses in a Presence Information Data
// Format Location Object file that you include in SIP requests. That helps ensure
// that addresses are routed to the appropriate Public Safety Answering Point.
func chimesdkvoice_ValidateE911Address(cfg aws.Config, client *chimesdkvoice.Client) {
	input := &chimesdkvoice.ValidateE911AddressInput{
		// AwsAccountId: *string, // Required
		// City: *string, // Required
		// Country: *string, // Required
		// PostalCode: *string, // Required
		// State: *string, // Required
		// StreetInfo: *string, // Required
		// StreetNumber: *string, // Required
	}

	if len(_chimesdkvoiceAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_chimesdkvoiceAwsAccountId)
	}
	if len(_chimesdkvoiceCity) > 0 {
		input.City = aws.String(_chimesdkvoiceCity)
	}
	if len(_chimesdkvoiceCountry) > 0 {
		input.Country = aws.String(_chimesdkvoiceCountry)
	}
	if len(_chimesdkvoicePostalCode) > 0 {
		input.PostalCode = aws.String(_chimesdkvoicePostalCode)
	}
	if len(_chimesdkvoiceState) > 0 {
		input.State = aws.String(_chimesdkvoiceState)
	}
	if len(_chimesdkvoiceStreetInfo) > 0 {
		input.StreetInfo = aws.String(_chimesdkvoiceStreetInfo)
	}
	if len(_chimesdkvoiceStreetNumber) > 0 {
		input.StreetNumber = aws.String(_chimesdkvoiceStreetNumber)
	}

	if resp, err := client.ValidateE911Address(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_chimesdkvoiceCmd)
	_chimesdkvoiceCmd.Flags().SortFlags = false

	_chimesdkvoiceCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_chimesdkvoiceCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_chimesdkvoiceCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceAreaCode, "area-code", "", "", "Area Code")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceArguments, "arguments", "", "", "Arguments")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceArgumentsMap, "arguments-map", "", "", "Arguments Map")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceAwsAccountId, "aws-account-id", "", "", "AWS Account ID")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceAwsRegion, "aws-region", "", "", "AWS Region")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceCallLeg, "call-leg", "", "", "Call Leg")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceCallingName, "calling-name", "", "", "Calling Name")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceCapabilities, "capabilities", "", "", "Capabilities")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceCity, "city", "", "", "City")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceClientRequestToken, "client-request-token", "", "", "Client Request Token")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceContactCenterSystemTypes, "contact-center-system-types", "", "", "Contact Center System Types")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceCountry, "country", "", "", "Country")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceCredentials, "credentials", "", "", "Credentials")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceDefaultSessionExpiryMinutes, "default-session-expiry-minutes", "", "", "Default Session Expiry Minutes")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceDescription, "description", "", "", "Description")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceDisabled, "disabled", "", "", "Disabled")
	_chimesdkvoiceCmd.Flags().StringSliceVarP(&_chimesdkvoiceE164PhoneNumbers, "e164-phone-numbers", "", nil, "E164 Phone Numbers")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceEmergencyCallingConfiguration, "emergency-calling-configuration", "", "", "Emergency Calling Configuration")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceEndpoints, "endpoints", "", "", "Endpoints")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceExpiryMinutes, "expiry-minutes", "", "", "Expiry Minutes")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceFallBackPhoneNumber, "fall-back-phone-number", "", "", "Fall Back Phone Number")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceFilterName, "filter-name", "", "", "Filter Name")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceFilterValue, "filter-value", "", "", "Filter Value")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceForceAssociate, "force-associate", "", "", "Force Associate")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceFromPhoneNumber, "from-phone-number", "", "", "From Phone Number")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceGeoMatchLevel, "geo-match-level", "", "", "Geo Match Level")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceGeoMatchParams, "geo-match-params", "", "", "Geo Match Params")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceIntegrationType, "integration-type", "", "", "Integration Type")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceIsCaller, "is-caller", "", "", "Is Caller")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceLanguageCode, "language-code", "", "", "Language Code")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceLoggingConfiguration, "logging-configuration", "", "", "Logging Configuration")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceMaxResults, "max-results", "", "", "Max Results")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceName, "name", "", "", "Name")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceNetworkType, "network-type", "", "", "Network Type")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceNextToken, "next-token", "", "", "Next Token")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceNumberSelectionBehavior, "number-selection-behavior", "", "", "Number Selection Behavior")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceOrigination, "origination", "", "", "Origination")
	_chimesdkvoiceCmd.Flags().StringSliceVarP(&_chimesdkvoiceParticipantPhoneNumbers, "participant-phone-numbers", "", nil, "Participant Phone Numbers")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoicePhoneNumberId, "phone-number-id", "", "", "Phone Number ID")
	_chimesdkvoiceCmd.Flags().StringSliceVarP(&_chimesdkvoicePhoneNumberIds, "phone-number-ids", "", nil, "Phone Number Ids")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoicePhoneNumberOrderId, "phone-number-order-id", "", "", "Phone Number Order ID")
	_chimesdkvoiceCmd.Flags().StringSliceVarP(&_chimesdkvoicePhoneNumberPoolCountries, "phone-number-pool-countries", "", nil, "Phone Number Pool Countries")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoicePhoneNumberType, "phone-number-type", "", "", "Phone Number Type")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoicePostalCode, "postal-code", "", "", "Postal Code")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceProductType, "product-type", "", "", "Product Type")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceProxySessionId, "proxy-session-id", "", "", "Proxy Session ID")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceRequireEncryption, "require-encryption", "", "", "Require Encryption")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceResourceARN, "resource-arn", "", "", "Resource ARN")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceServerSideEncryptionConfiguration, "server-side-encryption-configuration", "", "", "Server Side Encryption Configuration")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceSessionBorderControllerTypes, "session-border-controller-types", "", "", "Session Border Controller Types")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceSipHeaders, "sip-headers", "", "", "Sip Headers")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceSipMediaApplicationAlexaSkillConfiguration, "sip-media-application-alexa-skill-configuration", "", "", "Sip Media Application Alexa Skill Configuration")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceSipMediaApplicationId, "sip-media-application-id", "", "", "Sip Media Application ID")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceSipMediaApplicationLoggingConfiguration, "sip-media-application-logging-configuration", "", "", "Sip Media Application Logging Configuration")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceSipRuleId, "sip-rule-id", "", "", "Sip Rule ID")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceSpeakerSearchTaskId, "speaker-search-task-id", "", "", "Speaker Search Task ID")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceState, "state", "", "", "State")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceStatus, "status", "", "", "Status")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceStreamingConfiguration, "streaming-configuration", "", "", "Streaming Configuration")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceStreetInfo, "street-info", "", "", "Street Info")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceStreetNumber, "street-number", "", "", "Street Number")
	_chimesdkvoiceCmd.Flags().StringSliceVarP(&_chimesdkvoiceTagKeys, "tag-keys", "", nil, "Tag Keys")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceTags, "tags", "", "", "Tags")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceTargetApplications, "target-applications", "", "", "Target Applications")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceTermination, "termination", "", "", "Termination")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceToPhoneNumber, "to-phone-number", "", "", "To Phone Number")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceTollFreePrefix, "toll-free-prefix", "", "", "Toll Free Prefix")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceTransactionId, "transaction-id", "", "", "Transaction ID")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceTriggerType, "trigger-type", "", "", "Trigger Type")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceTriggerValue, "trigger-value", "", "", "Trigger Value")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceUpdatePhoneNumberRequestItems, "update-phone-number-request-items", "", "", "Update Phone Number Request Items")
	_chimesdkvoiceCmd.Flags().StringSliceVarP(&_chimesdkvoiceUsernames, "usernames", "", nil, "Usernames")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceVoiceConnector, "voice-connector", "", "", "Voice Connector")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceVoiceConnectorGroupId, "voice-connector-group-id", "", "", "Voice Connector Group ID")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceVoiceConnectorId, "voice-connector-id", "", "", "Voice Connector ID")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceVoiceConnectorItems, "voice-connector-items", "", "", "Voice Connector Items")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceVoiceProfileDomainId, "voice-profile-domain-id", "", "", "Voice Profile Domain ID")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceVoiceProfileId, "voice-profile-id", "", "", "Voice Profile ID")
	_chimesdkvoiceCmd.Flags().StringVarP(&_chimesdkvoiceVoiceToneAnalysisTaskId, "voice-tone-analysis-task-id", "", "", "Voice Tone Analysis Task ID")

	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceAssociatePhoneNumbersWithVoiceConnector, "associate-phone-numbers-with-voice-connector", "", false, "Associate Phone Numbers With Voice Connector")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceAssociatePhoneNumbersWithVoiceConnectorGroup, "associate-phone-numbers-with-voice-connector-group", "", false, "Associate Phone Numbers With Voice Connector Group")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceBatchDeletePhoneNumber, "batch-delete-phone-number", "", false, "Batch Delete Phone Number")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceBatchUpdatePhoneNumber, "batch-update-phone-number", "", false, "Batch Update Phone Number")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceCreatePhoneNumberOrder, "create-phone-number-order", "", false, "Create Phone Number Order")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceCreateProxySession, "create-proxy-session", "", false, "Create Proxy Session")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceCreateSipMediaApplication, "create-sip-media-application", "", false, "Create Sip Media Application")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceCreateSipMediaApplicationCall, "create-sip-media-application-call", "", false, "Create Sip Media Application Call")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceCreateSipRule, "create-sip-rule", "", false, "Create Sip Rule")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceCreateVoiceConnector, "create-voice-connector", "", false, "Create Voice Connector")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceCreateVoiceConnectorGroup, "create-voice-connector-group", "", false, "Create Voice Connector Group")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceCreateVoiceProfile, "create-voice-profile", "", false, "Create Voice Profile")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceCreateVoiceProfileDomain, "create-voice-profile-domain", "", false, "Create Voice Profile Domain")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceDeletePhoneNumber, "delete-phone-number", "", false, "Delete Phone Number")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceDeleteProxySession, "delete-proxy-session", "", false, "Delete Proxy Session")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceDeleteSipMediaApplication, "delete-sip-media-application", "", false, "Delete Sip Media Application")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceDeleteSipRule, "delete-sip-rule", "", false, "Delete Sip Rule")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceDeleteVoiceConnector, "delete-voice-connector", "", false, "Delete Voice Connector")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceDeleteVoiceConnectorEmergencyCallingConfiguration, "delete-voice-connector-emergency-calling-configuration", "", false, "Delete Voice Connector Emergency Calling Configuration")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceDeleteVoiceConnectorExternalSystemsConfiguration, "delete-voice-connector-external-systems-configuration", "", false, "Delete Voice Connector External Systems Configuration")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceDeleteVoiceConnectorGroup, "delete-voice-connector-group", "", false, "Delete Voice Connector Group")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceDeleteVoiceConnectorOrigination, "delete-voice-connector-origination", "", false, "Delete Voice Connector Origination")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceDeleteVoiceConnectorProxy, "delete-voice-connector-proxy", "", false, "Delete Voice Connector Proxy")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceDeleteVoiceConnectorStreamingConfiguration, "delete-voice-connector-streaming-configuration", "", false, "Delete Voice Connector Streaming Configuration")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceDeleteVoiceConnectorTermination, "delete-voice-connector-termination", "", false, "Delete Voice Connector Termination")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceDeleteVoiceConnectorTerminationCredentials, "delete-voice-connector-termination-credentials", "", false, "Delete Voice Connector Termination Credentials")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceDeleteVoiceProfile, "delete-voice-profile", "", false, "Delete Voice Profile")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceDeleteVoiceProfileDomain, "delete-voice-profile-domain", "", false, "Delete Voice Profile Domain")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceDisassociatePhoneNumbersFromVoiceConnector, "disassociate-phone-numbers-from-voice-connector", "", false, "Disassociate Phone Numbers From Voice Connector")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceDisassociatePhoneNumbersFromVoiceConnectorGroup, "disassociate-phone-numbers-from-voice-connector-group", "", false, "Disassociate Phone Numbers From Voice Connector Group")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceGetGlobalSettings, "get-global-settings", "", false, "Get Global Settings")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceGetPhoneNumber, "get-phone-number", "", false, "Get Phone Number")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceGetPhoneNumberOrder, "get-phone-number-order", "", false, "Get Phone Number Order")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceGetPhoneNumberSettings, "get-phone-number-settings", "", false, "Get Phone Number Settings")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceGetProxySession, "get-proxy-session", "", false, "Get Proxy Session")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceGetSipMediaApplication, "get-sip-media-application", "", false, "Get Sip Media Application")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceGetSipMediaApplicationAlexaSkillConfiguration, "get-sip-media-application-alexa-skill-configuration", "", false, "Get Sip Media Application Alexa Skill Configuration")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceGetSipMediaApplicationLoggingConfiguration, "get-sip-media-application-logging-configuration", "", false, "Get Sip Media Application Logging Configuration")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceGetSipRule, "get-sip-rule", "", false, "Get Sip Rule")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceGetSpeakerSearchTask, "get-speaker-search-task", "", false, "Get Speaker Search Task")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceGetVoiceConnector, "get-voice-connector", "", false, "Get Voice Connector")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceGetVoiceConnectorEmergencyCallingConfiguration, "get-voice-connector-emergency-calling-configuration", "", false, "Get Voice Connector Emergency Calling Configuration")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceGetVoiceConnectorExternalSystemsConfiguration, "get-voice-connector-external-systems-configuration", "", false, "Get Voice Connector External Systems Configuration")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceGetVoiceConnectorGroup, "get-voice-connector-group", "", false, "Get Voice Connector Group")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceGetVoiceConnectorLoggingConfiguration, "get-voice-connector-logging-configuration", "", false, "Get Voice Connector Logging Configuration")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceGetVoiceConnectorOrigination, "get-voice-connector-origination", "", false, "Get Voice Connector Origination")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceGetVoiceConnectorProxy, "get-voice-connector-proxy", "", false, "Get Voice Connector Proxy")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceGetVoiceConnectorStreamingConfiguration, "get-voice-connector-streaming-configuration", "", false, "Get Voice Connector Streaming Configuration")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceGetVoiceConnectorTermination, "get-voice-connector-termination", "", false, "Get Voice Connector Termination")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceGetVoiceConnectorTerminationHealth, "get-voice-connector-termination-health", "", false, "Get Voice Connector Termination Health")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceGetVoiceProfile, "get-voice-profile", "", false, "Get Voice Profile")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceGetVoiceProfileDomain, "get-voice-profile-domain", "", false, "Get Voice Profile Domain")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceGetVoiceToneAnalysisTask, "get-voice-tone-analysis-task", "", false, "Get Voice Tone Analysis Task")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceListAvailableVoiceConnectorRegions, "list-available-voice-connector-regions", "", false, "List Available Voice Connector Regions")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceListPhoneNumberOrders, "list-phone-number-orders", "", false, "List Phone Number Orders")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceListPhoneNumbers, "list-phone-numbers", "", false, "List Phone Numbers")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceListProxySessions, "list-proxy-sessions", "", false, "List Proxy Sessions")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceListSipMediaApplications, "list-sip-media-applications", "", false, "List Sip Media Applications")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceListSipRules, "list-sip-rules", "", false, "List Sip Rules")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceListSupportedPhoneNumberCountries, "list-supported-phone-number-countries", "", false, "List Supported Phone Number Countries")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceListVoiceConnectorGroups, "list-voice-connector-groups", "", false, "List Voice Connector Groups")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceListVoiceConnectorTerminationCredentials, "list-voice-connector-termination-credentials", "", false, "List Voice Connector Termination Credentials")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceListVoiceConnectors, "list-voice-connectors", "", false, "List Voice Connectors")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceListVoiceProfileDomains, "list-voice-profile-domains", "", false, "List Voice Profile Domains")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceListVoiceProfiles, "list-voice-profiles", "", false, "List Voice Profiles")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoicePutSipMediaApplicationAlexaSkillConfiguration, "put-sip-media-application-alexa-skill-configuration", "", false, "Put Sip Media Application Alexa Skill Configuration")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoicePutSipMediaApplicationLoggingConfiguration, "put-sip-media-application-logging-configuration", "", false, "Put Sip Media Application Logging Configuration")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoicePutVoiceConnectorEmergencyCallingConfiguration, "put-voice-connector-emergency-calling-configuration", "", false, "Put Voice Connector Emergency Calling Configuration")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoicePutVoiceConnectorExternalSystemsConfiguration, "put-voice-connector-external-systems-configuration", "", false, "Put Voice Connector External Systems Configuration")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoicePutVoiceConnectorLoggingConfiguration, "put-voice-connector-logging-configuration", "", false, "Put Voice Connector Logging Configuration")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoicePutVoiceConnectorOrigination, "put-voice-connector-origination", "", false, "Put Voice Connector Origination")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoicePutVoiceConnectorProxy, "put-voice-connector-proxy", "", false, "Put Voice Connector Proxy")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoicePutVoiceConnectorStreamingConfiguration, "put-voice-connector-streaming-configuration", "", false, "Put Voice Connector Streaming Configuration")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoicePutVoiceConnectorTermination, "put-voice-connector-termination", "", false, "Put Voice Connector Termination")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoicePutVoiceConnectorTerminationCredentials, "put-voice-connector-termination-credentials", "", false, "Put Voice Connector Termination Credentials")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceRestorePhoneNumber, "restore-phone-number", "", false, "Restore Phone Number")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceSearchAvailablePhoneNumbers, "search-available-phone-numbers", "", false, "Search Available Phone Numbers")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceStartSpeakerSearchTask, "start-speaker-search-task", "", false, "Start Speaker Search Task")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceStartVoiceToneAnalysisTask, "start-voice-tone-analysis-task", "", false, "Start Voice Tone Analysis Task")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceStopSpeakerSearchTask, "stop-speaker-search-task", "", false, "Stop Speaker Search Task")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceStopVoiceToneAnalysisTask, "stop-voice-tone-analysis-task", "", false, "Stop Voice Tone Analysis Task")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceTagResource, "tag-resource", "", false, "Tag Resource")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceUntagResource, "untag-resource", "", false, "Untag Resource")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceUpdateGlobalSettings, "update-global-settings", "", false, "Update Global Settings")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceUpdatePhoneNumber, "update-phone-number", "", false, "Update Phone Number")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceUpdatePhoneNumberSettings, "update-phone-number-settings", "", false, "Update Phone Number Settings")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceUpdateProxySession, "update-proxy-session", "", false, "Update Proxy Session")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceUpdateSipMediaApplication, "update-sip-media-application", "", false, "Update Sip Media Application")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceUpdateSipMediaApplicationCall, "update-sip-media-application-call", "", false, "Update Sip Media Application Call")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceUpdateSipRule, "update-sip-rule", "", false, "Update Sip Rule")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceUpdateVoiceConnector, "update-voice-connector", "", false, "Update Voice Connector")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceUpdateVoiceConnectorGroup, "update-voice-connector-group", "", false, "Update Voice Connector Group")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceUpdateVoiceProfile, "update-voice-profile", "", false, "Update Voice Profile")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceUpdateVoiceProfileDomain, "update-voice-profile-domain", "", false, "Update Voice Profile Domain")
	_chimesdkvoiceCmd.Flags().BoolVarP(&_chimesdkvoiceValidateE911Address, "validate-e911-address", "", false, "Validate E911 Address")

}
