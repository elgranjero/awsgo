package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cleanroomsml"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// cleanroomsmlCmd represents the cleanroomsml command
var _cleanroomsmlCmd = &cobra.Command{
	Use:   "cleanroomsml",
	Short: "AWS cleanroomsml CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := cleanroomsml.NewFromConfig(cfg)
		if _cleanroomsmlCancelTrainedModel {
			cleanroomsml_CancelTrainedModel(cfg, client)
			return
		}
		if _cleanroomsmlCancelTrainedModelInferenceJob {
			cleanroomsml_CancelTrainedModelInferenceJob(cfg, client)
			return
		}
		if _cleanroomsmlCreateAudienceModel {
			cleanroomsml_CreateAudienceModel(cfg, client)
			return
		}
		if _cleanroomsmlCreateConfiguredAudienceModel {
			cleanroomsml_CreateConfiguredAudienceModel(cfg, client)
			return
		}
		if _cleanroomsmlCreateConfiguredModelAlgorithm {
			cleanroomsml_CreateConfiguredModelAlgorithm(cfg, client)
			return
		}
		if _cleanroomsmlCreateConfiguredModelAlgorithmAssociation {
			cleanroomsml_CreateConfiguredModelAlgorithmAssociation(cfg, client)
			return
		}
		if _cleanroomsmlCreateMLInputChannel {
			cleanroomsml_CreateMLInputChannel(cfg, client)
			return
		}
		if _cleanroomsmlCreateTrainedModel {
			cleanroomsml_CreateTrainedModel(cfg, client)
			return
		}
		if _cleanroomsmlCreateTrainingDataset {
			cleanroomsml_CreateTrainingDataset(cfg, client)
			return
		}
		if _cleanroomsmlDeleteAudienceGenerationJob {
			cleanroomsml_DeleteAudienceGenerationJob(cfg, client)
			return
		}
		if _cleanroomsmlDeleteAudienceModel {
			cleanroomsml_DeleteAudienceModel(cfg, client)
			return
		}
		if _cleanroomsmlDeleteConfiguredAudienceModel {
			cleanroomsml_DeleteConfiguredAudienceModel(cfg, client)
			return
		}
		if _cleanroomsmlDeleteConfiguredAudienceModelPolicy {
			cleanroomsml_DeleteConfiguredAudienceModelPolicy(cfg, client)
			return
		}
		if _cleanroomsmlDeleteConfiguredModelAlgorithm {
			cleanroomsml_DeleteConfiguredModelAlgorithm(cfg, client)
			return
		}
		if _cleanroomsmlDeleteConfiguredModelAlgorithmAssociation {
			cleanroomsml_DeleteConfiguredModelAlgorithmAssociation(cfg, client)
			return
		}
		if _cleanroomsmlDeleteMLConfiguration {
			cleanroomsml_DeleteMLConfiguration(cfg, client)
			return
		}
		if _cleanroomsmlDeleteMLInputChannelData {
			cleanroomsml_DeleteMLInputChannelData(cfg, client)
			return
		}
		if _cleanroomsmlDeleteTrainedModelOutput {
			cleanroomsml_DeleteTrainedModelOutput(cfg, client)
			return
		}
		if _cleanroomsmlDeleteTrainingDataset {
			cleanroomsml_DeleteTrainingDataset(cfg, client)
			return
		}
		if _cleanroomsmlGetAudienceGenerationJob {
			cleanroomsml_GetAudienceGenerationJob(cfg, client)
			return
		}
		if _cleanroomsmlGetAudienceModel {
			cleanroomsml_GetAudienceModel(cfg, client)
			return
		}
		if _cleanroomsmlGetCollaborationConfiguredModelAlgorithmAssociation {
			cleanroomsml_GetCollaborationConfiguredModelAlgorithmAssociation(cfg, client)
			return
		}
		if _cleanroomsmlGetCollaborationMLInputChannel {
			cleanroomsml_GetCollaborationMLInputChannel(cfg, client)
			return
		}
		if _cleanroomsmlGetCollaborationTrainedModel {
			cleanroomsml_GetCollaborationTrainedModel(cfg, client)
			return
		}
		if _cleanroomsmlGetConfiguredAudienceModel {
			cleanroomsml_GetConfiguredAudienceModel(cfg, client)
			return
		}
		if _cleanroomsmlGetConfiguredAudienceModelPolicy {
			cleanroomsml_GetConfiguredAudienceModelPolicy(cfg, client)
			return
		}
		if _cleanroomsmlGetConfiguredModelAlgorithm {
			cleanroomsml_GetConfiguredModelAlgorithm(cfg, client)
			return
		}
		if _cleanroomsmlGetConfiguredModelAlgorithmAssociation {
			cleanroomsml_GetConfiguredModelAlgorithmAssociation(cfg, client)
			return
		}
		if _cleanroomsmlGetMLConfiguration {
			cleanroomsml_GetMLConfiguration(cfg, client)
			return
		}
		if _cleanroomsmlGetMLInputChannel {
			cleanroomsml_GetMLInputChannel(cfg, client)
			return
		}
		if _cleanroomsmlGetTrainedModel {
			cleanroomsml_GetTrainedModel(cfg, client)
			return
		}
		if _cleanroomsmlGetTrainedModelInferenceJob {
			cleanroomsml_GetTrainedModelInferenceJob(cfg, client)
			return
		}
		if _cleanroomsmlGetTrainingDataset {
			cleanroomsml_GetTrainingDataset(cfg, client)
			return
		}
		if _cleanroomsmlListAudienceExportJobs {
			cleanroomsml_ListAudienceExportJobs(cfg, client)
			return
		}
		if _cleanroomsmlListAudienceGenerationJobs {
			cleanroomsml_ListAudienceGenerationJobs(cfg, client)
			return
		}
		if _cleanroomsmlListAudienceModels {
			cleanroomsml_ListAudienceModels(cfg, client)
			return
		}
		if _cleanroomsmlListCollaborationConfiguredModelAlgorithmAssociations {
			cleanroomsml_ListCollaborationConfiguredModelAlgorithmAssociations(cfg, client)
			return
		}
		if _cleanroomsmlListCollaborationMLInputChannels {
			cleanroomsml_ListCollaborationMLInputChannels(cfg, client)
			return
		}
		if _cleanroomsmlListCollaborationTrainedModelExportJobs {
			cleanroomsml_ListCollaborationTrainedModelExportJobs(cfg, client)
			return
		}
		if _cleanroomsmlListCollaborationTrainedModelInferenceJobs {
			cleanroomsml_ListCollaborationTrainedModelInferenceJobs(cfg, client)
			return
		}
		if _cleanroomsmlListCollaborationTrainedModels {
			cleanroomsml_ListCollaborationTrainedModels(cfg, client)
			return
		}
		if _cleanroomsmlListConfiguredAudienceModels {
			cleanroomsml_ListConfiguredAudienceModels(cfg, client)
			return
		}
		if _cleanroomsmlListConfiguredModelAlgorithmAssociations {
			cleanroomsml_ListConfiguredModelAlgorithmAssociations(cfg, client)
			return
		}
		if _cleanroomsmlListConfiguredModelAlgorithms {
			cleanroomsml_ListConfiguredModelAlgorithms(cfg, client)
			return
		}
		if _cleanroomsmlListMLInputChannels {
			cleanroomsml_ListMLInputChannels(cfg, client)
			return
		}
		if _cleanroomsmlListTagsForResource {
			cleanroomsml_ListTagsForResource(cfg, client)
			return
		}
		if _cleanroomsmlListTrainedModelInferenceJobs {
			cleanroomsml_ListTrainedModelInferenceJobs(cfg, client)
			return
		}
		if _cleanroomsmlListTrainedModelVersions {
			cleanroomsml_ListTrainedModelVersions(cfg, client)
			return
		}
		if _cleanroomsmlListTrainedModels {
			cleanroomsml_ListTrainedModels(cfg, client)
			return
		}
		if _cleanroomsmlListTrainingDatasets {
			cleanroomsml_ListTrainingDatasets(cfg, client)
			return
		}
		if _cleanroomsmlPutConfiguredAudienceModelPolicy {
			cleanroomsml_PutConfiguredAudienceModelPolicy(cfg, client)
			return
		}
		if _cleanroomsmlPutMLConfiguration {
			cleanroomsml_PutMLConfiguration(cfg, client)
			return
		}
		if _cleanroomsmlStartAudienceExportJob {
			cleanroomsml_StartAudienceExportJob(cfg, client)
			return
		}
		if _cleanroomsmlStartAudienceGenerationJob {
			cleanroomsml_StartAudienceGenerationJob(cfg, client)
			return
		}
		if _cleanroomsmlStartTrainedModelExportJob {
			cleanroomsml_StartTrainedModelExportJob(cfg, client)
			return
		}
		if _cleanroomsmlStartTrainedModelInferenceJob {
			cleanroomsml_StartTrainedModelInferenceJob(cfg, client)
			return
		}
		if _cleanroomsmlTagResource {
			cleanroomsml_TagResource(cfg, client)
			return
		}
		if _cleanroomsmlUntagResource {
			cleanroomsml_UntagResource(cfg, client)
			return
		}
		if _cleanroomsmlUpdateConfiguredAudienceModel {
			cleanroomsml_UpdateConfiguredAudienceModel(cfg, client)
			return
		}

	},
}

var (
	_cleanroomsmlCancelTrainedModel                                    bool
	_cleanroomsmlCancelTrainedModelInferenceJob                        bool
	_cleanroomsmlCreateAudienceModel                                   bool
	_cleanroomsmlCreateConfiguredAudienceModel                         bool
	_cleanroomsmlCreateConfiguredModelAlgorithm                        bool
	_cleanroomsmlCreateConfiguredModelAlgorithmAssociation             bool
	_cleanroomsmlCreateMLInputChannel                                  bool
	_cleanroomsmlCreateTrainedModel                                    bool
	_cleanroomsmlCreateTrainingDataset                                 bool
	_cleanroomsmlDeleteAudienceGenerationJob                           bool
	_cleanroomsmlDeleteAudienceModel                                   bool
	_cleanroomsmlDeleteConfiguredAudienceModel                         bool
	_cleanroomsmlDeleteConfiguredAudienceModelPolicy                   bool
	_cleanroomsmlDeleteConfiguredModelAlgorithm                        bool
	_cleanroomsmlDeleteConfiguredModelAlgorithmAssociation             bool
	_cleanroomsmlDeleteMLConfiguration                                 bool
	_cleanroomsmlDeleteMLInputChannelData                              bool
	_cleanroomsmlDeleteTrainedModelOutput                              bool
	_cleanroomsmlDeleteTrainingDataset                                 bool
	_cleanroomsmlGetAudienceGenerationJob                              bool
	_cleanroomsmlGetAudienceModel                                      bool
	_cleanroomsmlGetCollaborationConfiguredModelAlgorithmAssociation   bool
	_cleanroomsmlGetCollaborationMLInputChannel                        bool
	_cleanroomsmlGetCollaborationTrainedModel                          bool
	_cleanroomsmlGetConfiguredAudienceModel                            bool
	_cleanroomsmlGetConfiguredAudienceModelPolicy                      bool
	_cleanroomsmlGetConfiguredModelAlgorithm                           bool
	_cleanroomsmlGetConfiguredModelAlgorithmAssociation                bool
	_cleanroomsmlGetMLConfiguration                                    bool
	_cleanroomsmlGetMLInputChannel                                     bool
	_cleanroomsmlGetTrainedModel                                       bool
	_cleanroomsmlGetTrainedModelInferenceJob                           bool
	_cleanroomsmlGetTrainingDataset                                    bool
	_cleanroomsmlListAudienceExportJobs                                bool
	_cleanroomsmlListAudienceGenerationJobs                            bool
	_cleanroomsmlListAudienceModels                                    bool
	_cleanroomsmlListCollaborationConfiguredModelAlgorithmAssociations bool
	_cleanroomsmlListCollaborationMLInputChannels                      bool
	_cleanroomsmlListCollaborationTrainedModelExportJobs               bool
	_cleanroomsmlListCollaborationTrainedModelInferenceJobs            bool
	_cleanroomsmlListCollaborationTrainedModels                        bool
	_cleanroomsmlListConfiguredAudienceModels                          bool
	_cleanroomsmlListConfiguredModelAlgorithmAssociations              bool
	_cleanroomsmlListConfiguredModelAlgorithms                         bool
	_cleanroomsmlListMLInputChannels                                   bool
	_cleanroomsmlListTagsForResource                                   bool
	_cleanroomsmlListTrainedModelInferenceJobs                         bool
	_cleanroomsmlListTrainedModelVersions                              bool
	_cleanroomsmlListTrainedModels                                     bool
	_cleanroomsmlListTrainingDatasets                                  bool
	_cleanroomsmlPutConfiguredAudienceModelPolicy                      bool
	_cleanroomsmlPutMLConfiguration                                    bool
	_cleanroomsmlStartAudienceExportJob                                bool
	_cleanroomsmlStartAudienceGenerationJob                            bool
	_cleanroomsmlStartTrainedModelExportJob                            bool
	_cleanroomsmlStartTrainedModelInferenceJob                         bool
	_cleanroomsmlTagResource                                           bool
	_cleanroomsmlUntagResource                                         bool
	_cleanroomsmlUpdateConfiguredAudienceModel                         bool

	_cleanroomsmlAudienceGenerationJobArn               string
	_cleanroomsmlAudienceModelArn                       string
	_cleanroomsmlAudienceSize                           string
	_cleanroomsmlAudienceSizeConfig                     string
	_cleanroomsmlChildResourceTagOnCreatePolicy         string
	_cleanroomsmlCollaborationId                        string
	_cleanroomsmlCollaborationIdentifier                string
	_cleanroomsmlConfiguredAudienceModelArn             string
	_cleanroomsmlConfiguredAudienceModelPolicy          string
	_cleanroomsmlConfiguredModelAlgorithmArn            string
	_cleanroomsmlConfiguredModelAlgorithmAssociationArn string
	_cleanroomsmlConfiguredModelAlgorithmAssociations   []string
	_cleanroomsmlContainerExecutionParameters           string
	_cleanroomsmlDataChannels                           string
	_cleanroomsmlDataSource                             string
	_cleanroomsmlDefaultOutputLocation                  string
	_cleanroomsmlDescription                            string
	_cleanroomsmlEnvironment                            string
	_cleanroomsmlHyperparameters                        string
	_cleanroomsmlIncludeSeedInOutput                    string
	_cleanroomsmlIncrementalTrainingDataChannels        string
	_cleanroomsmlInferenceContainerConfig               string
	_cleanroomsmlInputChannel                           string
	_cleanroomsmlKmsKeyArn                              string
	_cleanroomsmlMaxResults                             string
	_cleanroomsmlMembershipIdentifier                   string
	_cleanroomsmlMinMatchingSeedSize                    string
	_cleanroomsmlMlInputChannelArn                      string
	_cleanroomsmlName                                   string
	_cleanroomsmlNextToken                              string
	_cleanroomsmlOutputConfig                           string
	_cleanroomsmlOutputConfiguration                    string
	_cleanroomsmlPolicyExistenceCondition               string
	_cleanroomsmlPreviousPolicyHash                     string
	_cleanroomsmlPrivacyConfiguration                   string
	_cleanroomsmlResourceArn                            string
	_cleanroomsmlResourceConfig                         string
	_cleanroomsmlRetentionInDays                        string
	_cleanroomsmlRoleArn                                string
	_cleanroomsmlSeedAudience                           string
	_cleanroomsmlSharedAudienceMetrics                  string
	_cleanroomsmlStatus                                 string
	_cleanroomsmlStoppingCondition                      string
	_cleanroomsmlTagKeys                                []string
	_cleanroomsmlTags                                   string
	_cleanroomsmlTrainedModelArn                        string
	_cleanroomsmlTrainedModelInferenceJobArn            string
	_cleanroomsmlTrainedModelVersionIdentifier          string
	_cleanroomsmlTrainingContainerConfig                string
	_cleanroomsmlTrainingData                           string
	_cleanroomsmlTrainingDataEndTime                    string
	_cleanroomsmlTrainingDataStartTime                  string
	_cleanroomsmlTrainingDatasetArn                     string
	_cleanroomsmlTrainingInputMode                      string
	_cleanroomsmlVersionIdentifier                      string
)

// Submits a request to cancel the trained model job.
func cleanroomsml_CancelTrainedModel(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.CancelTrainedModelInput{
		// MembershipIdentifier: *string, // Required
		// TrainedModelArn: *string, // Required
	}

	if len(_cleanroomsmlMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsmlMembershipIdentifier)
	}
	if len(_cleanroomsmlTrainedModelArn) > 0 {
		input.TrainedModelArn = aws.String(_cleanroomsmlTrainedModelArn)
	}
	if len(_cleanroomsmlVersionIdentifier) > 0 {
		input.VersionIdentifier = aws.String(_cleanroomsmlVersionIdentifier)
	}

	if resp, err := client.CancelTrainedModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Submits a request to cancel a trained model inference job.
func cleanroomsml_CancelTrainedModelInferenceJob(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.CancelTrainedModelInferenceJobInput{
		// MembershipIdentifier: *string, // Required
		// TrainedModelInferenceJobArn: *string, // Required
	}

	if len(_cleanroomsmlMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsmlMembershipIdentifier)
	}
	if len(_cleanroomsmlTrainedModelInferenceJobArn) > 0 {
		input.TrainedModelInferenceJobArn = aws.String(_cleanroomsmlTrainedModelInferenceJobArn)
	}

	if resp, err := client.CancelTrainedModelInferenceJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Defines the information necessary to create an audience model. An audience
// model is a machine learning model that Clean Rooms ML trains to measure
// similarity between users. Clean Rooms ML manages training and storing the
// audience model. The audience model can be used in multiple calls to the StartAudienceGenerationJobAPI.
func cleanroomsml_CreateAudienceModel(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.CreateAudienceModelInput{
		// Name: *string, // Required
		// TrainingDatasetArn: *string, // Required
	}

	if len(_cleanroomsmlName) > 0 {
		input.Name = aws.String(_cleanroomsmlName)
	}
	if len(_cleanroomsmlTrainingDatasetArn) > 0 {
		input.TrainingDatasetArn = aws.String(_cleanroomsmlTrainingDatasetArn)
	}
	if len(_cleanroomsmlDescription) > 0 {
		input.Description = aws.String(_cleanroomsmlDescription)
	}
	if len(_cleanroomsmlKmsKeyArn) > 0 {
		input.KmsKeyArn = aws.String(_cleanroomsmlKmsKeyArn)
	}
	if len(_cleanroomsmlTags) > 0 {
		if err := assignInputField(input, "Tags", _cleanroomsmlTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlTrainingDataEndTime) > 0 {
		if err := assignInputField(input, "TrainingDataEndTime", _cleanroomsmlTrainingDataEndTime); err != nil {
			log.Errorf("invalid --training-data-end-time: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlTrainingDataStartTime) > 0 {
		if err := assignInputField(input, "TrainingDataStartTime", _cleanroomsmlTrainingDataStartTime); err != nil {
			log.Errorf("invalid --training-data-start-time: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAudienceModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Defines the information necessary to create a configured audience model.
func cleanroomsml_CreateConfiguredAudienceModel(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.CreateConfiguredAudienceModelInput{
		// AudienceModelArn: *string, // Required
		// Name: *string, // Required
		// OutputConfig: *types.ConfiguredAudienceModelOutputConfig, // Required
		// SharedAudienceMetrics: []types.SharedAudienceMetrics, // Required
	}

	if len(_cleanroomsmlAudienceModelArn) > 0 {
		input.AudienceModelArn = aws.String(_cleanroomsmlAudienceModelArn)
	}
	if len(_cleanroomsmlName) > 0 {
		input.Name = aws.String(_cleanroomsmlName)
	}
	if len(_cleanroomsmlOutputConfig) > 0 {
		if err := assignInputField(input, "OutputConfig", _cleanroomsmlOutputConfig); err != nil {
			log.Errorf("invalid --output-config: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlSharedAudienceMetrics) > 0 {
		if err := assignInputField(input, "SharedAudienceMetrics", _cleanroomsmlSharedAudienceMetrics); err != nil {
			log.Errorf("invalid --shared-audience-metrics: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlAudienceSizeConfig) > 0 {
		if err := assignInputField(input, "AudienceSizeConfig", _cleanroomsmlAudienceSizeConfig); err != nil {
			log.Errorf("invalid --audience-size-config: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlChildResourceTagOnCreatePolicy) > 0 {
		if err := assignInputField(input, "ChildResourceTagOnCreatePolicy", _cleanroomsmlChildResourceTagOnCreatePolicy); err != nil {
			log.Errorf("invalid --child-resource-tag-on-create-policy: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlDescription) > 0 {
		input.Description = aws.String(_cleanroomsmlDescription)
	}
	if len(_cleanroomsmlMinMatchingSeedSize) > 0 {
		if err := assignInputField(input, "MinMatchingSeedSize", _cleanroomsmlMinMatchingSeedSize); err != nil {
			log.Errorf("invalid --min-matching-seed-size: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlTags) > 0 {
		if err := assignInputField(input, "Tags", _cleanroomsmlTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateConfiguredAudienceModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a configured model algorithm using a container image stored in an ECR
// repository.
func cleanroomsml_CreateConfiguredModelAlgorithm(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.CreateConfiguredModelAlgorithmInput{
		// Name: *string, // Required
		// RoleArn: *string, // Required
	}

	if len(_cleanroomsmlName) > 0 {
		input.Name = aws.String(_cleanroomsmlName)
	}
	if len(_cleanroomsmlRoleArn) > 0 {
		input.RoleArn = aws.String(_cleanroomsmlRoleArn)
	}
	if len(_cleanroomsmlDescription) > 0 {
		input.Description = aws.String(_cleanroomsmlDescription)
	}
	if len(_cleanroomsmlInferenceContainerConfig) > 0 {
		if err := assignInputField(input, "InferenceContainerConfig", _cleanroomsmlInferenceContainerConfig); err != nil {
			log.Errorf("invalid --inference-container-config: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlKmsKeyArn) > 0 {
		input.KmsKeyArn = aws.String(_cleanroomsmlKmsKeyArn)
	}
	if len(_cleanroomsmlTags) > 0 {
		if err := assignInputField(input, "Tags", _cleanroomsmlTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlTrainingContainerConfig) > 0 {
		if err := assignInputField(input, "TrainingContainerConfig", _cleanroomsmlTrainingContainerConfig); err != nil {
			log.Errorf("invalid --training-container-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateConfiguredModelAlgorithm(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a configured model algorithm to a collaboration for use by any
// member of the collaboration.
func cleanroomsml_CreateConfiguredModelAlgorithmAssociation(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.CreateConfiguredModelAlgorithmAssociationInput{
		// ConfiguredModelAlgorithmArn: *string, // Required
		// MembershipIdentifier: *string, // Required
		// Name: *string, // Required
	}

	if len(_cleanroomsmlConfiguredModelAlgorithmArn) > 0 {
		input.ConfiguredModelAlgorithmArn = aws.String(_cleanroomsmlConfiguredModelAlgorithmArn)
	}
	if len(_cleanroomsmlMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsmlMembershipIdentifier)
	}
	if len(_cleanroomsmlName) > 0 {
		input.Name = aws.String(_cleanroomsmlName)
	}
	if len(_cleanroomsmlDescription) > 0 {
		input.Description = aws.String(_cleanroomsmlDescription)
	}
	if len(_cleanroomsmlPrivacyConfiguration) > 0 {
		if err := assignInputField(input, "PrivacyConfiguration", _cleanroomsmlPrivacyConfiguration); err != nil {
			log.Errorf("invalid --privacy-configuration: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlTags) > 0 {
		if err := assignInputField(input, "Tags", _cleanroomsmlTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateConfiguredModelAlgorithmAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides the information to create an ML input channel. An ML input channel is
// the result of a query that can be used for ML modeling.
func cleanroomsml_CreateMLInputChannel(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.CreateMLInputChannelInput{
		// ConfiguredModelAlgorithmAssociations: []string, // Required
		// InputChannel: *types.InputChannel, // Required
		// MembershipIdentifier: *string, // Required
		// Name: *string, // Required
		// RetentionInDays: *int32, // Required
	}

	if len(_cleanroomsmlConfiguredModelAlgorithmAssociations) > 0 {
		input.ConfiguredModelAlgorithmAssociations = append([]string(nil), _cleanroomsmlConfiguredModelAlgorithmAssociations...)
	}
	if len(_cleanroomsmlInputChannel) > 0 {
		if err := assignInputField(input, "InputChannel", _cleanroomsmlInputChannel); err != nil {
			log.Errorf("invalid --input-channel: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsmlMembershipIdentifier)
	}
	if len(_cleanroomsmlName) > 0 {
		input.Name = aws.String(_cleanroomsmlName)
	}
	if len(_cleanroomsmlRetentionInDays) > 0 {
		if err := assignInputField(input, "RetentionInDays", _cleanroomsmlRetentionInDays); err != nil {
			log.Errorf("invalid --retention-in-days: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlDescription) > 0 {
		input.Description = aws.String(_cleanroomsmlDescription)
	}
	if len(_cleanroomsmlKmsKeyArn) > 0 {
		input.KmsKeyArn = aws.String(_cleanroomsmlKmsKeyArn)
	}
	if len(_cleanroomsmlTags) > 0 {
		if err := assignInputField(input, "Tags", _cleanroomsmlTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateMLInputChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a trained model from an associated configured model algorithm using
// data from any member of the collaboration.
func cleanroomsml_CreateTrainedModel(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.CreateTrainedModelInput{
		// ConfiguredModelAlgorithmAssociationArn: *string, // Required
		// DataChannels: []types.ModelTrainingDataChannel, // Required
		// MembershipIdentifier: *string, // Required
		// Name: *string, // Required
		// ResourceConfig: *types.ResourceConfig, // Required
	}

	if len(_cleanroomsmlConfiguredModelAlgorithmAssociationArn) > 0 {
		input.ConfiguredModelAlgorithmAssociationArn = aws.String(_cleanroomsmlConfiguredModelAlgorithmAssociationArn)
	}
	if len(_cleanroomsmlDataChannels) > 0 {
		if err := assignInputField(input, "DataChannels", _cleanroomsmlDataChannels); err != nil {
			log.Errorf("invalid --data-channels: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsmlMembershipIdentifier)
	}
	if len(_cleanroomsmlName) > 0 {
		input.Name = aws.String(_cleanroomsmlName)
	}
	if len(_cleanroomsmlResourceConfig) > 0 {
		if err := assignInputField(input, "ResourceConfig", _cleanroomsmlResourceConfig); err != nil {
			log.Errorf("invalid --resource-config: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlDescription) > 0 {
		input.Description = aws.String(_cleanroomsmlDescription)
	}
	if len(_cleanroomsmlEnvironment) > 0 {
		if err := assignInputField(input, "Environment", _cleanroomsmlEnvironment); err != nil {
			log.Errorf("invalid --environment: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlHyperparameters) > 0 {
		if err := assignInputField(input, "Hyperparameters", _cleanroomsmlHyperparameters); err != nil {
			log.Errorf("invalid --hyperparameters: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlIncrementalTrainingDataChannels) > 0 {
		if err := assignInputField(input, "IncrementalTrainingDataChannels", _cleanroomsmlIncrementalTrainingDataChannels); err != nil {
			log.Errorf("invalid --incremental-training-data-channels: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlKmsKeyArn) > 0 {
		input.KmsKeyArn = aws.String(_cleanroomsmlKmsKeyArn)
	}
	if len(_cleanroomsmlStoppingCondition) > 0 {
		if err := assignInputField(input, "StoppingCondition", _cleanroomsmlStoppingCondition); err != nil {
			log.Errorf("invalid --stopping-condition: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlTags) > 0 {
		if err := assignInputField(input, "Tags", _cleanroomsmlTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlTrainingInputMode) > 0 {
		if err := assignInputField(input, "TrainingInputMode", _cleanroomsmlTrainingInputMode); err != nil {
			log.Errorf("invalid --training-input-mode: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTrainedModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Defines the information necessary to create a training dataset. In Clean Rooms
// ML, the TrainingDataset is metadata that points to a Glue table, which is read
// only during AudienceModel creation.
func cleanroomsml_CreateTrainingDataset(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.CreateTrainingDatasetInput{
		// Name: *string, // Required
		// RoleArn: *string, // Required
		// TrainingData: []types.Dataset, // Required
	}

	if len(_cleanroomsmlName) > 0 {
		input.Name = aws.String(_cleanroomsmlName)
	}
	if len(_cleanroomsmlRoleArn) > 0 {
		input.RoleArn = aws.String(_cleanroomsmlRoleArn)
	}
	if len(_cleanroomsmlTrainingData) > 0 {
		if err := assignInputField(input, "TrainingData", _cleanroomsmlTrainingData); err != nil {
			log.Errorf("invalid --training-data: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlDescription) > 0 {
		input.Description = aws.String(_cleanroomsmlDescription)
	}
	if len(_cleanroomsmlTags) > 0 {
		if err := assignInputField(input, "Tags", _cleanroomsmlTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTrainingDataset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified audience generation job, and removes all data associated
// with the job.
func cleanroomsml_DeleteAudienceGenerationJob(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.DeleteAudienceGenerationJobInput{
		// AudienceGenerationJobArn: *string, // Required
	}

	if len(_cleanroomsmlAudienceGenerationJobArn) > 0 {
		input.AudienceGenerationJobArn = aws.String(_cleanroomsmlAudienceGenerationJobArn)
	}

	if resp, err := client.DeleteAudienceGenerationJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Specifies an audience model that you want to delete. You can't delete an
// audience model if there are any configured audience models that depend on the
// audience model.
func cleanroomsml_DeleteAudienceModel(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.DeleteAudienceModelInput{
		// AudienceModelArn: *string, // Required
	}

	if len(_cleanroomsmlAudienceModelArn) > 0 {
		input.AudienceModelArn = aws.String(_cleanroomsmlAudienceModelArn)
	}

	if resp, err := client.DeleteAudienceModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified configured audience model. You can't delete a configured
// audience model if there are any lookalike models that use the configured
// audience model. If you delete a configured audience model, it will be removed
// from any collaborations that it is associated to.
func cleanroomsml_DeleteConfiguredAudienceModel(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.DeleteConfiguredAudienceModelInput{
		// ConfiguredAudienceModelArn: *string, // Required
	}

	if len(_cleanroomsmlConfiguredAudienceModelArn) > 0 {
		input.ConfiguredAudienceModelArn = aws.String(_cleanroomsmlConfiguredAudienceModelArn)
	}

	if resp, err := client.DeleteConfiguredAudienceModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified configured audience model policy.
func cleanroomsml_DeleteConfiguredAudienceModelPolicy(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.DeleteConfiguredAudienceModelPolicyInput{
		// ConfiguredAudienceModelArn: *string, // Required
	}

	if len(_cleanroomsmlConfiguredAudienceModelArn) > 0 {
		input.ConfiguredAudienceModelArn = aws.String(_cleanroomsmlConfiguredAudienceModelArn)
	}

	if resp, err := client.DeleteConfiguredAudienceModelPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a configured model algorithm.
func cleanroomsml_DeleteConfiguredModelAlgorithm(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.DeleteConfiguredModelAlgorithmInput{
		// ConfiguredModelAlgorithmArn: *string, // Required
	}

	if len(_cleanroomsmlConfiguredModelAlgorithmArn) > 0 {
		input.ConfiguredModelAlgorithmArn = aws.String(_cleanroomsmlConfiguredModelAlgorithmArn)
	}

	if resp, err := client.DeleteConfiguredModelAlgorithm(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a configured model algorithm association.
func cleanroomsml_DeleteConfiguredModelAlgorithmAssociation(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.DeleteConfiguredModelAlgorithmAssociationInput{
		// ConfiguredModelAlgorithmAssociationArn: *string, // Required
		// MembershipIdentifier: *string, // Required
	}

	if len(_cleanroomsmlConfiguredModelAlgorithmAssociationArn) > 0 {
		input.ConfiguredModelAlgorithmAssociationArn = aws.String(_cleanroomsmlConfiguredModelAlgorithmAssociationArn)
	}
	if len(_cleanroomsmlMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsmlMembershipIdentifier)
	}

	if resp, err := client.DeleteConfiguredModelAlgorithmAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a ML modeling configuration.
func cleanroomsml_DeleteMLConfiguration(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.DeleteMLConfigurationInput{
		// MembershipIdentifier: *string, // Required
	}

	if len(_cleanroomsmlMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsmlMembershipIdentifier)
	}

	if resp, err := client.DeleteMLConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides the information necessary to delete an ML input channel.
func cleanroomsml_DeleteMLInputChannelData(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.DeleteMLInputChannelDataInput{
		// MembershipIdentifier: *string, // Required
		// MlInputChannelArn: *string, // Required
	}

	if len(_cleanroomsmlMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsmlMembershipIdentifier)
	}
	if len(_cleanroomsmlMlInputChannelArn) > 0 {
		input.MlInputChannelArn = aws.String(_cleanroomsmlMlInputChannelArn)
	}

	if resp, err := client.DeleteMLInputChannelData(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the model artifacts stored by the service.
func cleanroomsml_DeleteTrainedModelOutput(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.DeleteTrainedModelOutputInput{
		// MembershipIdentifier: *string, // Required
		// TrainedModelArn: *string, // Required
	}

	if len(_cleanroomsmlMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsmlMembershipIdentifier)
	}
	if len(_cleanroomsmlTrainedModelArn) > 0 {
		input.TrainedModelArn = aws.String(_cleanroomsmlTrainedModelArn)
	}
	if len(_cleanroomsmlVersionIdentifier) > 0 {
		input.VersionIdentifier = aws.String(_cleanroomsmlVersionIdentifier)
	}

	if resp, err := client.DeleteTrainedModelOutput(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Specifies a training dataset that you want to delete. You can't delete a
// training dataset if there are any audience models that depend on the training
// dataset. In Clean Rooms ML, the TrainingDataset is metadata that points to a
// Glue table, which is read only during AudienceModel creation. This action
// deletes the metadata.
func cleanroomsml_DeleteTrainingDataset(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.DeleteTrainingDatasetInput{
		// TrainingDatasetArn: *string, // Required
	}

	if len(_cleanroomsmlTrainingDatasetArn) > 0 {
		input.TrainingDatasetArn = aws.String(_cleanroomsmlTrainingDatasetArn)
	}

	if resp, err := client.DeleteTrainingDataset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about an audience generation job.
func cleanroomsml_GetAudienceGenerationJob(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.GetAudienceGenerationJobInput{
		// AudienceGenerationJobArn: *string, // Required
	}

	if len(_cleanroomsmlAudienceGenerationJobArn) > 0 {
		input.AudienceGenerationJobArn = aws.String(_cleanroomsmlAudienceGenerationJobArn)
	}

	if resp, err := client.GetAudienceGenerationJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about an audience model
func cleanroomsml_GetAudienceModel(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.GetAudienceModelInput{
		// AudienceModelArn: *string, // Required
	}

	if len(_cleanroomsmlAudienceModelArn) > 0 {
		input.AudienceModelArn = aws.String(_cleanroomsmlAudienceModelArn)
	}

	if resp, err := client.GetAudienceModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the configured model algorithm association in a
// collaboration.
func cleanroomsml_GetCollaborationConfiguredModelAlgorithmAssociation(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.GetCollaborationConfiguredModelAlgorithmAssociationInput{
		// CollaborationIdentifier: *string, // Required
		// ConfiguredModelAlgorithmAssociationArn: *string, // Required
	}

	if len(_cleanroomsmlCollaborationIdentifier) > 0 {
		input.CollaborationIdentifier = aws.String(_cleanroomsmlCollaborationIdentifier)
	}
	if len(_cleanroomsmlConfiguredModelAlgorithmAssociationArn) > 0 {
		input.ConfiguredModelAlgorithmAssociationArn = aws.String(_cleanroomsmlConfiguredModelAlgorithmAssociationArn)
	}

	if resp, err := client.GetCollaborationConfiguredModelAlgorithmAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a specific ML input channel in a collaboration.
func cleanroomsml_GetCollaborationMLInputChannel(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.GetCollaborationMLInputChannelInput{
		// CollaborationIdentifier: *string, // Required
		// MlInputChannelArn: *string, // Required
	}

	if len(_cleanroomsmlCollaborationIdentifier) > 0 {
		input.CollaborationIdentifier = aws.String(_cleanroomsmlCollaborationIdentifier)
	}
	if len(_cleanroomsmlMlInputChannelArn) > 0 {
		input.MlInputChannelArn = aws.String(_cleanroomsmlMlInputChannelArn)
	}

	if resp, err := client.GetCollaborationMLInputChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a trained model in a collaboration.
func cleanroomsml_GetCollaborationTrainedModel(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.GetCollaborationTrainedModelInput{
		// CollaborationIdentifier: *string, // Required
		// TrainedModelArn: *string, // Required
	}

	if len(_cleanroomsmlCollaborationIdentifier) > 0 {
		input.CollaborationIdentifier = aws.String(_cleanroomsmlCollaborationIdentifier)
	}
	if len(_cleanroomsmlTrainedModelArn) > 0 {
		input.TrainedModelArn = aws.String(_cleanroomsmlTrainedModelArn)
	}
	if len(_cleanroomsmlVersionIdentifier) > 0 {
		input.VersionIdentifier = aws.String(_cleanroomsmlVersionIdentifier)
	}

	if resp, err := client.GetCollaborationTrainedModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a specified configured audience model.
func cleanroomsml_GetConfiguredAudienceModel(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.GetConfiguredAudienceModelInput{
		// ConfiguredAudienceModelArn: *string, // Required
	}

	if len(_cleanroomsmlConfiguredAudienceModelArn) > 0 {
		input.ConfiguredAudienceModelArn = aws.String(_cleanroomsmlConfiguredAudienceModelArn)
	}

	if resp, err := client.GetConfiguredAudienceModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a configured audience model policy.
func cleanroomsml_GetConfiguredAudienceModelPolicy(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.GetConfiguredAudienceModelPolicyInput{
		// ConfiguredAudienceModelArn: *string, // Required
	}

	if len(_cleanroomsmlConfiguredAudienceModelArn) > 0 {
		input.ConfiguredAudienceModelArn = aws.String(_cleanroomsmlConfiguredAudienceModelArn)
	}

	if resp, err := client.GetConfiguredAudienceModelPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a configured model algorithm.
func cleanroomsml_GetConfiguredModelAlgorithm(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.GetConfiguredModelAlgorithmInput{
		// ConfiguredModelAlgorithmArn: *string, // Required
	}

	if len(_cleanroomsmlConfiguredModelAlgorithmArn) > 0 {
		input.ConfiguredModelAlgorithmArn = aws.String(_cleanroomsmlConfiguredModelAlgorithmArn)
	}

	if resp, err := client.GetConfiguredModelAlgorithm(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a configured model algorithm association.
func cleanroomsml_GetConfiguredModelAlgorithmAssociation(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.GetConfiguredModelAlgorithmAssociationInput{
		// ConfiguredModelAlgorithmAssociationArn: *string, // Required
		// MembershipIdentifier: *string, // Required
	}

	if len(_cleanroomsmlConfiguredModelAlgorithmAssociationArn) > 0 {
		input.ConfiguredModelAlgorithmAssociationArn = aws.String(_cleanroomsmlConfiguredModelAlgorithmAssociationArn)
	}
	if len(_cleanroomsmlMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsmlMembershipIdentifier)
	}

	if resp, err := client.GetConfiguredModelAlgorithmAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a specific ML configuration.
func cleanroomsml_GetMLConfiguration(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.GetMLConfigurationInput{
		// MembershipIdentifier: *string, // Required
	}

	if len(_cleanroomsmlMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsmlMembershipIdentifier)
	}

	if resp, err := client.GetMLConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about an ML input channel.
func cleanroomsml_GetMLInputChannel(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.GetMLInputChannelInput{
		// MembershipIdentifier: *string, // Required
		// MlInputChannelArn: *string, // Required
	}

	if len(_cleanroomsmlMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsmlMembershipIdentifier)
	}
	if len(_cleanroomsmlMlInputChannelArn) > 0 {
		input.MlInputChannelArn = aws.String(_cleanroomsmlMlInputChannelArn)
	}

	if resp, err := client.GetMLInputChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a trained model.
func cleanroomsml_GetTrainedModel(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.GetTrainedModelInput{
		// MembershipIdentifier: *string, // Required
		// TrainedModelArn: *string, // Required
	}

	if len(_cleanroomsmlMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsmlMembershipIdentifier)
	}
	if len(_cleanroomsmlTrainedModelArn) > 0 {
		input.TrainedModelArn = aws.String(_cleanroomsmlTrainedModelArn)
	}
	if len(_cleanroomsmlVersionIdentifier) > 0 {
		input.VersionIdentifier = aws.String(_cleanroomsmlVersionIdentifier)
	}

	if resp, err := client.GetTrainedModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a trained model inference job.
func cleanroomsml_GetTrainedModelInferenceJob(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.GetTrainedModelInferenceJobInput{
		// MembershipIdentifier: *string, // Required
		// TrainedModelInferenceJobArn: *string, // Required
	}

	if len(_cleanroomsmlMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsmlMembershipIdentifier)
	}
	if len(_cleanroomsmlTrainedModelInferenceJobArn) > 0 {
		input.TrainedModelInferenceJobArn = aws.String(_cleanroomsmlTrainedModelInferenceJobArn)
	}

	if resp, err := client.GetTrainedModelInferenceJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a training dataset.
func cleanroomsml_GetTrainingDataset(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.GetTrainingDatasetInput{
		// TrainingDatasetArn: *string, // Required
	}

	if len(_cleanroomsmlTrainingDatasetArn) > 0 {
		input.TrainingDatasetArn = aws.String(_cleanroomsmlTrainingDatasetArn)
	}

	if resp, err := client.GetTrainingDataset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of the audience export jobs.
func cleanroomsml_ListAudienceExportJobs(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.ListAudienceExportJobsInput{}

	if len(_cleanroomsmlAudienceGenerationJobArn) > 0 {
		input.AudienceGenerationJobArn = aws.String(_cleanroomsmlAudienceGenerationJobArn)
	}
	if len(_cleanroomsmlMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cleanroomsmlMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlNextToken) > 0 {
		input.NextToken = aws.String(_cleanroomsmlNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAudienceExportJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cleanroomsml.ListAudienceExportJobsOutput
	p := cleanroomsml.NewListAudienceExportJobsPaginator(client, input)
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

// Returns a list of audience generation jobs.
func cleanroomsml_ListAudienceGenerationJobs(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.ListAudienceGenerationJobsInput{}

	if len(_cleanroomsmlCollaborationId) > 0 {
		input.CollaborationId = aws.String(_cleanroomsmlCollaborationId)
	}
	if len(_cleanroomsmlConfiguredAudienceModelArn) > 0 {
		input.ConfiguredAudienceModelArn = aws.String(_cleanroomsmlConfiguredAudienceModelArn)
	}
	if len(_cleanroomsmlMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cleanroomsmlMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlNextToken) > 0 {
		input.NextToken = aws.String(_cleanroomsmlNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAudienceGenerationJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cleanroomsml.ListAudienceGenerationJobsOutput
	p := cleanroomsml.NewListAudienceGenerationJobsPaginator(client, input)
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

// Returns a list of audience models.
func cleanroomsml_ListAudienceModels(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.ListAudienceModelsInput{}

	if len(_cleanroomsmlMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cleanroomsmlMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlNextToken) > 0 {
		input.NextToken = aws.String(_cleanroomsmlNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAudienceModels(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cleanroomsml.ListAudienceModelsOutput
	p := cleanroomsml.NewListAudienceModelsPaginator(client, input)
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

// Returns a list of the configured model algorithm associations in a
// collaboration.
func cleanroomsml_ListCollaborationConfiguredModelAlgorithmAssociations(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.ListCollaborationConfiguredModelAlgorithmAssociationsInput{
		// CollaborationIdentifier: *string, // Required
	}

	if len(_cleanroomsmlCollaborationIdentifier) > 0 {
		input.CollaborationIdentifier = aws.String(_cleanroomsmlCollaborationIdentifier)
	}
	if len(_cleanroomsmlMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cleanroomsmlMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlNextToken) > 0 {
		input.NextToken = aws.String(_cleanroomsmlNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCollaborationConfiguredModelAlgorithmAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cleanroomsml.ListCollaborationConfiguredModelAlgorithmAssociationsOutput
	p := cleanroomsml.NewListCollaborationConfiguredModelAlgorithmAssociationsPaginator(client, input)
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

// Returns a list of the ML input channels in a collaboration.
func cleanroomsml_ListCollaborationMLInputChannels(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.ListCollaborationMLInputChannelsInput{
		// CollaborationIdentifier: *string, // Required
	}

	if len(_cleanroomsmlCollaborationIdentifier) > 0 {
		input.CollaborationIdentifier = aws.String(_cleanroomsmlCollaborationIdentifier)
	}
	if len(_cleanroomsmlMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cleanroomsmlMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlNextToken) > 0 {
		input.NextToken = aws.String(_cleanroomsmlNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCollaborationMLInputChannels(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cleanroomsml.ListCollaborationMLInputChannelsOutput
	p := cleanroomsml.NewListCollaborationMLInputChannelsPaginator(client, input)
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

// Returns a list of the export jobs for a trained model in a collaboration.
func cleanroomsml_ListCollaborationTrainedModelExportJobs(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.ListCollaborationTrainedModelExportJobsInput{
		// CollaborationIdentifier: *string, // Required
		// TrainedModelArn: *string, // Required
	}

	if len(_cleanroomsmlCollaborationIdentifier) > 0 {
		input.CollaborationIdentifier = aws.String(_cleanroomsmlCollaborationIdentifier)
	}
	if len(_cleanroomsmlTrainedModelArn) > 0 {
		input.TrainedModelArn = aws.String(_cleanroomsmlTrainedModelArn)
	}
	if len(_cleanroomsmlMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cleanroomsmlMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlNextToken) > 0 {
		input.NextToken = aws.String(_cleanroomsmlNextToken)
	}
	if len(_cleanroomsmlTrainedModelVersionIdentifier) > 0 {
		input.TrainedModelVersionIdentifier = aws.String(_cleanroomsmlTrainedModelVersionIdentifier)
	}

	if disablePaginator() {
		if resp, err := client.ListCollaborationTrainedModelExportJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cleanroomsml.ListCollaborationTrainedModelExportJobsOutput
	p := cleanroomsml.NewListCollaborationTrainedModelExportJobsPaginator(client, input)
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

// Returns a list of trained model inference jobs in a specified collaboration.
func cleanroomsml_ListCollaborationTrainedModelInferenceJobs(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.ListCollaborationTrainedModelInferenceJobsInput{
		// CollaborationIdentifier: *string, // Required
	}

	if len(_cleanroomsmlCollaborationIdentifier) > 0 {
		input.CollaborationIdentifier = aws.String(_cleanroomsmlCollaborationIdentifier)
	}
	if len(_cleanroomsmlMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cleanroomsmlMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlNextToken) > 0 {
		input.NextToken = aws.String(_cleanroomsmlNextToken)
	}
	if len(_cleanroomsmlTrainedModelArn) > 0 {
		input.TrainedModelArn = aws.String(_cleanroomsmlTrainedModelArn)
	}
	if len(_cleanroomsmlTrainedModelVersionIdentifier) > 0 {
		input.TrainedModelVersionIdentifier = aws.String(_cleanroomsmlTrainedModelVersionIdentifier)
	}

	if disablePaginator() {
		if resp, err := client.ListCollaborationTrainedModelInferenceJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cleanroomsml.ListCollaborationTrainedModelInferenceJobsOutput
	p := cleanroomsml.NewListCollaborationTrainedModelInferenceJobsPaginator(client, input)
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

// Returns a list of the trained models in a collaboration.
func cleanroomsml_ListCollaborationTrainedModels(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.ListCollaborationTrainedModelsInput{
		// CollaborationIdentifier: *string, // Required
	}

	if len(_cleanroomsmlCollaborationIdentifier) > 0 {
		input.CollaborationIdentifier = aws.String(_cleanroomsmlCollaborationIdentifier)
	}
	if len(_cleanroomsmlMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cleanroomsmlMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlNextToken) > 0 {
		input.NextToken = aws.String(_cleanroomsmlNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCollaborationTrainedModels(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cleanroomsml.ListCollaborationTrainedModelsOutput
	p := cleanroomsml.NewListCollaborationTrainedModelsPaginator(client, input)
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

// Returns a list of the configured audience models.
func cleanroomsml_ListConfiguredAudienceModels(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.ListConfiguredAudienceModelsInput{}

	if len(_cleanroomsmlMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cleanroomsmlMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlNextToken) > 0 {
		input.NextToken = aws.String(_cleanroomsmlNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListConfiguredAudienceModels(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cleanroomsml.ListConfiguredAudienceModelsOutput
	p := cleanroomsml.NewListConfiguredAudienceModelsPaginator(client, input)
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

// Returns a list of configured model algorithm associations.
func cleanroomsml_ListConfiguredModelAlgorithmAssociations(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.ListConfiguredModelAlgorithmAssociationsInput{
		// MembershipIdentifier: *string, // Required
	}

	if len(_cleanroomsmlMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsmlMembershipIdentifier)
	}
	if len(_cleanroomsmlMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cleanroomsmlMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlNextToken) > 0 {
		input.NextToken = aws.String(_cleanroomsmlNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListConfiguredModelAlgorithmAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cleanroomsml.ListConfiguredModelAlgorithmAssociationsOutput
	p := cleanroomsml.NewListConfiguredModelAlgorithmAssociationsPaginator(client, input)
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

// Returns a list of configured model algorithms.
func cleanroomsml_ListConfiguredModelAlgorithms(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.ListConfiguredModelAlgorithmsInput{}

	if len(_cleanroomsmlMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cleanroomsmlMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlNextToken) > 0 {
		input.NextToken = aws.String(_cleanroomsmlNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListConfiguredModelAlgorithms(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cleanroomsml.ListConfiguredModelAlgorithmsOutput
	p := cleanroomsml.NewListConfiguredModelAlgorithmsPaginator(client, input)
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

// Returns a list of ML input channels.
func cleanroomsml_ListMLInputChannels(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.ListMLInputChannelsInput{
		// MembershipIdentifier: *string, // Required
	}

	if len(_cleanroomsmlMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsmlMembershipIdentifier)
	}
	if len(_cleanroomsmlMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cleanroomsmlMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlNextToken) > 0 {
		input.NextToken = aws.String(_cleanroomsmlNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListMLInputChannels(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cleanroomsml.ListMLInputChannelsOutput
	p := cleanroomsml.NewListMLInputChannelsPaginator(client, input)
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

// Returns a list of tags for a provided resource.
func cleanroomsml_ListTagsForResource(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_cleanroomsmlResourceArn) > 0 {
		input.ResourceArn = aws.String(_cleanroomsmlResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of trained model inference jobs that match the request
// parameters.
func cleanroomsml_ListTrainedModelInferenceJobs(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.ListTrainedModelInferenceJobsInput{
		// MembershipIdentifier: *string, // Required
	}

	if len(_cleanroomsmlMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsmlMembershipIdentifier)
	}
	if len(_cleanroomsmlMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cleanroomsmlMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlNextToken) > 0 {
		input.NextToken = aws.String(_cleanroomsmlNextToken)
	}
	if len(_cleanroomsmlTrainedModelArn) > 0 {
		input.TrainedModelArn = aws.String(_cleanroomsmlTrainedModelArn)
	}
	if len(_cleanroomsmlTrainedModelVersionIdentifier) > 0 {
		input.TrainedModelVersionIdentifier = aws.String(_cleanroomsmlTrainedModelVersionIdentifier)
	}

	if disablePaginator() {
		if resp, err := client.ListTrainedModelInferenceJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cleanroomsml.ListTrainedModelInferenceJobsOutput
	p := cleanroomsml.NewListTrainedModelInferenceJobsPaginator(client, input)
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

// Returns a list of trained model versions for a specified trained model. This
// operation allows you to view all versions of a trained model, including
// information about their status and creation details. You can use this to track
// the evolution of your trained models and select specific versions for inference
// or further training.
func cleanroomsml_ListTrainedModelVersions(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.ListTrainedModelVersionsInput{
		// MembershipIdentifier: *string, // Required
		// TrainedModelArn: *string, // Required
	}

	if len(_cleanroomsmlMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsmlMembershipIdentifier)
	}
	if len(_cleanroomsmlTrainedModelArn) > 0 {
		input.TrainedModelArn = aws.String(_cleanroomsmlTrainedModelArn)
	}
	if len(_cleanroomsmlMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cleanroomsmlMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlNextToken) > 0 {
		input.NextToken = aws.String(_cleanroomsmlNextToken)
	}
	if len(_cleanroomsmlStatus) > 0 {
		if err := assignInputField(input, "Status", _cleanroomsmlStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListTrainedModelVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cleanroomsml.ListTrainedModelVersionsOutput
	p := cleanroomsml.NewListTrainedModelVersionsPaginator(client, input)
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

// Returns a list of trained models.
func cleanroomsml_ListTrainedModels(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.ListTrainedModelsInput{
		// MembershipIdentifier: *string, // Required
	}

	if len(_cleanroomsmlMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsmlMembershipIdentifier)
	}
	if len(_cleanroomsmlMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cleanroomsmlMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlNextToken) > 0 {
		input.NextToken = aws.String(_cleanroomsmlNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTrainedModels(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cleanroomsml.ListTrainedModelsOutput
	p := cleanroomsml.NewListTrainedModelsPaginator(client, input)
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

// Returns a list of training datasets.
func cleanroomsml_ListTrainingDatasets(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.ListTrainingDatasetsInput{}

	if len(_cleanroomsmlMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cleanroomsmlMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlNextToken) > 0 {
		input.NextToken = aws.String(_cleanroomsmlNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTrainingDatasets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cleanroomsml.ListTrainingDatasetsOutput
	p := cleanroomsml.NewListTrainingDatasetsPaginator(client, input)
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

// Create or update the resource policy for a configured audience model.
func cleanroomsml_PutConfiguredAudienceModelPolicy(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.PutConfiguredAudienceModelPolicyInput{
		// ConfiguredAudienceModelArn: *string, // Required
		// ConfiguredAudienceModelPolicy: *string, // Required
	}

	if len(_cleanroomsmlConfiguredAudienceModelArn) > 0 {
		input.ConfiguredAudienceModelArn = aws.String(_cleanroomsmlConfiguredAudienceModelArn)
	}
	if len(_cleanroomsmlConfiguredAudienceModelPolicy) > 0 {
		input.ConfiguredAudienceModelPolicy = aws.String(_cleanroomsmlConfiguredAudienceModelPolicy)
	}
	if len(_cleanroomsmlPolicyExistenceCondition) > 0 {
		if err := assignInputField(input, "PolicyExistenceCondition", _cleanroomsmlPolicyExistenceCondition); err != nil {
			log.Errorf("invalid --policy-existence-condition: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlPreviousPolicyHash) > 0 {
		input.PreviousPolicyHash = aws.String(_cleanroomsmlPreviousPolicyHash)
	}

	if resp, err := client.PutConfiguredAudienceModelPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assigns information about an ML configuration.
func cleanroomsml_PutMLConfiguration(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.PutMLConfigurationInput{
		// DefaultOutputLocation: *types.MLOutputConfiguration, // Required
		// MembershipIdentifier: *string, // Required
	}

	if len(_cleanroomsmlDefaultOutputLocation) > 0 {
		if err := assignInputField(input, "DefaultOutputLocation", _cleanroomsmlDefaultOutputLocation); err != nil {
			log.Errorf("invalid --default-output-location: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsmlMembershipIdentifier)
	}

	if resp, err := client.PutMLConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Export an audience of a specified size after you have generated an audience.
func cleanroomsml_StartAudienceExportJob(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.StartAudienceExportJobInput{
		// AudienceGenerationJobArn: *string, // Required
		// AudienceSize: *types.AudienceSize, // Required
		// Name: *string, // Required
	}

	if len(_cleanroomsmlAudienceGenerationJobArn) > 0 {
		input.AudienceGenerationJobArn = aws.String(_cleanroomsmlAudienceGenerationJobArn)
	}
	if len(_cleanroomsmlAudienceSize) > 0 {
		if err := assignInputField(input, "AudienceSize", _cleanroomsmlAudienceSize); err != nil {
			log.Errorf("invalid --audience-size: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlName) > 0 {
		input.Name = aws.String(_cleanroomsmlName)
	}
	if len(_cleanroomsmlDescription) > 0 {
		input.Description = aws.String(_cleanroomsmlDescription)
	}

	if resp, err := client.StartAudienceExportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Information necessary to start the audience generation job.
func cleanroomsml_StartAudienceGenerationJob(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.StartAudienceGenerationJobInput{
		// ConfiguredAudienceModelArn: *string, // Required
		// Name: *string, // Required
		// SeedAudience: *types.AudienceGenerationJobDataSource, // Required
	}

	if len(_cleanroomsmlConfiguredAudienceModelArn) > 0 {
		input.ConfiguredAudienceModelArn = aws.String(_cleanroomsmlConfiguredAudienceModelArn)
	}
	if len(_cleanroomsmlName) > 0 {
		input.Name = aws.String(_cleanroomsmlName)
	}
	if len(_cleanroomsmlSeedAudience) > 0 {
		if err := assignInputField(input, "SeedAudience", _cleanroomsmlSeedAudience); err != nil {
			log.Errorf("invalid --seed-audience: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlCollaborationId) > 0 {
		input.CollaborationId = aws.String(_cleanroomsmlCollaborationId)
	}
	if len(_cleanroomsmlDescription) > 0 {
		input.Description = aws.String(_cleanroomsmlDescription)
	}
	if len(_cleanroomsmlIncludeSeedInOutput) > 0 {
		if err := assignInputField(input, "IncludeSeedInOutput", _cleanroomsmlIncludeSeedInOutput); err != nil {
			log.Errorf("invalid --include-seed-in-output: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlTags) > 0 {
		if err := assignInputField(input, "Tags", _cleanroomsmlTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartAudienceGenerationJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides the information necessary to start a trained model export job.
func cleanroomsml_StartTrainedModelExportJob(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.StartTrainedModelExportJobInput{
		// MembershipIdentifier: *string, // Required
		// Name: *string, // Required
		// OutputConfiguration: *types.TrainedModelExportOutputConfiguration, // Required
		// TrainedModelArn: *string, // Required
	}

	if len(_cleanroomsmlMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsmlMembershipIdentifier)
	}
	if len(_cleanroomsmlName) > 0 {
		input.Name = aws.String(_cleanroomsmlName)
	}
	if len(_cleanroomsmlOutputConfiguration) > 0 {
		if err := assignInputField(input, "OutputConfiguration", _cleanroomsmlOutputConfiguration); err != nil {
			log.Errorf("invalid --output-configuration: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlTrainedModelArn) > 0 {
		input.TrainedModelArn = aws.String(_cleanroomsmlTrainedModelArn)
	}
	if len(_cleanroomsmlDescription) > 0 {
		input.Description = aws.String(_cleanroomsmlDescription)
	}
	if len(_cleanroomsmlTrainedModelVersionIdentifier) > 0 {
		input.TrainedModelVersionIdentifier = aws.String(_cleanroomsmlTrainedModelVersionIdentifier)
	}

	if resp, err := client.StartTrainedModelExportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Defines the information necessary to begin a trained model inference job.
func cleanroomsml_StartTrainedModelInferenceJob(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.StartTrainedModelInferenceJobInput{
		// DataSource: *types.ModelInferenceDataSource, // Required
		// MembershipIdentifier: *string, // Required
		// Name: *string, // Required
		// OutputConfiguration: *types.InferenceOutputConfiguration, // Required
		// ResourceConfig: *types.InferenceResourceConfig, // Required
		// TrainedModelArn: *string, // Required
	}

	if len(_cleanroomsmlDataSource) > 0 {
		if err := assignInputField(input, "DataSource", _cleanroomsmlDataSource); err != nil {
			log.Errorf("invalid --data-source: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsmlMembershipIdentifier)
	}
	if len(_cleanroomsmlName) > 0 {
		input.Name = aws.String(_cleanroomsmlName)
	}
	if len(_cleanroomsmlOutputConfiguration) > 0 {
		if err := assignInputField(input, "OutputConfiguration", _cleanroomsmlOutputConfiguration); err != nil {
			log.Errorf("invalid --output-configuration: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlResourceConfig) > 0 {
		if err := assignInputField(input, "ResourceConfig", _cleanroomsmlResourceConfig); err != nil {
			log.Errorf("invalid --resource-config: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlTrainedModelArn) > 0 {
		input.TrainedModelArn = aws.String(_cleanroomsmlTrainedModelArn)
	}
	if len(_cleanroomsmlConfiguredModelAlgorithmAssociationArn) > 0 {
		input.ConfiguredModelAlgorithmAssociationArn = aws.String(_cleanroomsmlConfiguredModelAlgorithmAssociationArn)
	}
	if len(_cleanroomsmlContainerExecutionParameters) > 0 {
		if err := assignInputField(input, "ContainerExecutionParameters", _cleanroomsmlContainerExecutionParameters); err != nil {
			log.Errorf("invalid --container-execution-parameters: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlDescription) > 0 {
		input.Description = aws.String(_cleanroomsmlDescription)
	}
	if len(_cleanroomsmlEnvironment) > 0 {
		if err := assignInputField(input, "Environment", _cleanroomsmlEnvironment); err != nil {
			log.Errorf("invalid --environment: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlKmsKeyArn) > 0 {
		input.KmsKeyArn = aws.String(_cleanroomsmlKmsKeyArn)
	}
	if len(_cleanroomsmlTags) > 0 {
		if err := assignInputField(input, "Tags", _cleanroomsmlTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlTrainedModelVersionIdentifier) > 0 {
		input.TrainedModelVersionIdentifier = aws.String(_cleanroomsmlTrainedModelVersionIdentifier)
	}

	if resp, err := client.StartTrainedModelInferenceJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds metadata tags to a specified resource.
func cleanroomsml_TagResource(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_cleanroomsmlResourceArn) > 0 {
		input.ResourceArn = aws.String(_cleanroomsmlResourceArn)
	}
	if len(_cleanroomsmlTags) > 0 {
		if err := assignInputField(input, "Tags", _cleanroomsmlTags); err != nil {
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

// Removes metadata tags from a specified resource.
func cleanroomsml_UntagResource(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_cleanroomsmlResourceArn) > 0 {
		input.ResourceArn = aws.String(_cleanroomsmlResourceArn)
	}
	if len(_cleanroomsmlTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _cleanroomsmlTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides the information necessary to update a configured audience model.
// Updates that impact audience generation jobs take effect when a new job starts,
// but do not impact currently running jobs.
func cleanroomsml_UpdateConfiguredAudienceModel(cfg aws.Config, client *cleanroomsml.Client) {
	input := &cleanroomsml.UpdateConfiguredAudienceModelInput{
		// ConfiguredAudienceModelArn: *string, // Required
	}

	if len(_cleanroomsmlConfiguredAudienceModelArn) > 0 {
		input.ConfiguredAudienceModelArn = aws.String(_cleanroomsmlConfiguredAudienceModelArn)
	}
	if len(_cleanroomsmlAudienceModelArn) > 0 {
		input.AudienceModelArn = aws.String(_cleanroomsmlAudienceModelArn)
	}
	if len(_cleanroomsmlAudienceSizeConfig) > 0 {
		if err := assignInputField(input, "AudienceSizeConfig", _cleanroomsmlAudienceSizeConfig); err != nil {
			log.Errorf("invalid --audience-size-config: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlDescription) > 0 {
		input.Description = aws.String(_cleanroomsmlDescription)
	}
	if len(_cleanroomsmlMinMatchingSeedSize) > 0 {
		if err := assignInputField(input, "MinMatchingSeedSize", _cleanroomsmlMinMatchingSeedSize); err != nil {
			log.Errorf("invalid --min-matching-seed-size: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlOutputConfig) > 0 {
		if err := assignInputField(input, "OutputConfig", _cleanroomsmlOutputConfig); err != nil {
			log.Errorf("invalid --output-config: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsmlSharedAudienceMetrics) > 0 {
		if err := assignInputField(input, "SharedAudienceMetrics", _cleanroomsmlSharedAudienceMetrics); err != nil {
			log.Errorf("invalid --shared-audience-metrics: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateConfiguredAudienceModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_cleanroomsmlCmd)
	_cleanroomsmlCmd.Flags().SortFlags = false

	_cleanroomsmlCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_cleanroomsmlCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_cleanroomsmlCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlAudienceGenerationJobArn, "audience-generation-job-arn", "", "", "Audience Generation Job ARN")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlAudienceModelArn, "audience-model-arn", "", "", "Audience Model ARN")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlAudienceSize, "audience-size", "", "", "Audience Size")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlAudienceSizeConfig, "audience-size-config", "", "", "Audience Size Config")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlChildResourceTagOnCreatePolicy, "child-resource-tag-on-create-policy", "", "", "Child Resource Tag On Create Policy")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlCollaborationId, "collaboration-id", "", "", "Collaboration ID")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlCollaborationIdentifier, "collaboration-identifier", "", "", "Collaboration Identifier")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlConfiguredAudienceModelArn, "configured-audience-model-arn", "", "", "Configured Audience Model ARN")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlConfiguredAudienceModelPolicy, "configured-audience-model-policy", "", "", "Configured Audience Model Policy")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlConfiguredModelAlgorithmArn, "configured-model-algorithm-arn", "", "", "Configured Model Algorithm ARN")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlConfiguredModelAlgorithmAssociationArn, "configured-model-algorithm-association-arn", "", "", "Configured Model Algorithm Association ARN")
	_cleanroomsmlCmd.Flags().StringSliceVarP(&_cleanroomsmlConfiguredModelAlgorithmAssociations, "configured-model-algorithm-associations", "", nil, "Configured Model Algorithm Associations")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlContainerExecutionParameters, "container-execution-parameters", "", "", "Container Execution Parameters")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlDataChannels, "data-channels", "", "", "Data Channels")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlDataSource, "data-source", "", "", "Data Source")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlDefaultOutputLocation, "default-output-location", "", "", "Default Output Location")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlDescription, "description", "", "", "Description")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlEnvironment, "environment", "", "", "Environment")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlHyperparameters, "hyperparameters", "", "", "Hyperparameters")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlIncludeSeedInOutput, "include-seed-in-output", "", "", "Include Seed In Output")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlIncrementalTrainingDataChannels, "incremental-training-data-channels", "", "", "Incremental Training Data Channels")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlInferenceContainerConfig, "inference-container-config", "", "", "Inference Container Config")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlInputChannel, "input-channel", "", "", "Input Channel")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlKmsKeyArn, "kms-key-arn", "", "", "KMS Key ARN")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlMaxResults, "max-results", "", "", "Max Results")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlMembershipIdentifier, "membership-identifier", "", "", "Membership Identifier")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlMinMatchingSeedSize, "min-matching-seed-size", "", "", "Min Matching Seed Size")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlMlInputChannelArn, "ml-input-channel-arn", "", "", "Ml Input Channel ARN")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlName, "name", "", "", "Name")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlNextToken, "next-token", "", "", "Next Token")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlOutputConfig, "output-config", "", "", "Output Config")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlOutputConfiguration, "output-configuration", "", "", "Output Configuration")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlPolicyExistenceCondition, "policy-existence-condition", "", "", "Policy Existence Condition")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlPreviousPolicyHash, "previous-policy-hash", "", "", "Previous Policy Hash")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlPrivacyConfiguration, "privacy-configuration", "", "", "Privacy Configuration")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlResourceArn, "resource-arn", "", "", "Resource ARN")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlResourceConfig, "resource-config", "", "", "Resource Config")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlRetentionInDays, "retention-in-days", "", "", "Retention In Days")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlRoleArn, "role-arn", "", "", "Role ARN")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlSeedAudience, "seed-audience", "", "", "Seed Audience")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlSharedAudienceMetrics, "shared-audience-metrics", "", "", "Shared Audience Metrics")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlStatus, "status", "", "", "Status")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlStoppingCondition, "stopping-condition", "", "", "Stopping Condition")
	_cleanroomsmlCmd.Flags().StringSliceVarP(&_cleanroomsmlTagKeys, "tag-keys", "", nil, "Tag Keys")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlTags, "tags", "", "", "Tags")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlTrainedModelArn, "trained-model-arn", "", "", "Trained Model ARN")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlTrainedModelInferenceJobArn, "trained-model-inference-job-arn", "", "", "Trained Model Inference Job ARN")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlTrainedModelVersionIdentifier, "trained-model-version-identifier", "", "", "Trained Model Version Identifier")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlTrainingContainerConfig, "training-container-config", "", "", "Training Container Config")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlTrainingData, "training-data", "", "", "Training Data")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlTrainingDataEndTime, "training-data-end-time", "", "", "Training Data End Time")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlTrainingDataStartTime, "training-data-start-time", "", "", "Training Data Start Time")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlTrainingDatasetArn, "training-dataset-arn", "", "", "Training Dataset ARN")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlTrainingInputMode, "training-input-mode", "", "", "Training Input Mode")
	_cleanroomsmlCmd.Flags().StringVarP(&_cleanroomsmlVersionIdentifier, "version-identifier", "", "", "Version Identifier")

	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlCancelTrainedModel, "cancel-trained-model", "", false, "Cancel Trained Model")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlCancelTrainedModelInferenceJob, "cancel-trained-model-inference-job", "", false, "Cancel Trained Model Inference Job")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlCreateAudienceModel, "create-audience-model", "", false, "Create Audience Model")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlCreateConfiguredAudienceModel, "create-configured-audience-model", "", false, "Create Configured Audience Model")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlCreateConfiguredModelAlgorithm, "create-configured-model-algorithm", "", false, "Create Configured Model Algorithm")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlCreateConfiguredModelAlgorithmAssociation, "create-configured-model-algorithm-association", "", false, "Create Configured Model Algorithm Association")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlCreateMLInputChannel, "create-ml-input-channel", "", false, "Create Ml Input Channel")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlCreateTrainedModel, "create-trained-model", "", false, "Create Trained Model")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlCreateTrainingDataset, "create-training-dataset", "", false, "Create Training Dataset")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlDeleteAudienceGenerationJob, "delete-audience-generation-job", "", false, "Delete Audience Generation Job")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlDeleteAudienceModel, "delete-audience-model", "", false, "Delete Audience Model")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlDeleteConfiguredAudienceModel, "delete-configured-audience-model", "", false, "Delete Configured Audience Model")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlDeleteConfiguredAudienceModelPolicy, "delete-configured-audience-model-policy", "", false, "Delete Configured Audience Model Policy")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlDeleteConfiguredModelAlgorithm, "delete-configured-model-algorithm", "", false, "Delete Configured Model Algorithm")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlDeleteConfiguredModelAlgorithmAssociation, "delete-configured-model-algorithm-association", "", false, "Delete Configured Model Algorithm Association")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlDeleteMLConfiguration, "delete-ml-configuration", "", false, "Delete Ml Configuration")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlDeleteMLInputChannelData, "delete-ml-input-channel-data", "", false, "Delete Ml Input Channel Data")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlDeleteTrainedModelOutput, "delete-trained-model-output", "", false, "Delete Trained Model Output")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlDeleteTrainingDataset, "delete-training-dataset", "", false, "Delete Training Dataset")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlGetAudienceGenerationJob, "get-audience-generation-job", "", false, "Get Audience Generation Job")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlGetAudienceModel, "get-audience-model", "", false, "Get Audience Model")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlGetCollaborationConfiguredModelAlgorithmAssociation, "get-collaboration-configured-model-algorithm-association", "", false, "Get Collaboration Configured Model Algorithm Association")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlGetCollaborationMLInputChannel, "get-collaboration-ml-input-channel", "", false, "Get Collaboration Ml Input Channel")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlGetCollaborationTrainedModel, "get-collaboration-trained-model", "", false, "Get Collaboration Trained Model")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlGetConfiguredAudienceModel, "get-configured-audience-model", "", false, "Get Configured Audience Model")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlGetConfiguredAudienceModelPolicy, "get-configured-audience-model-policy", "", false, "Get Configured Audience Model Policy")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlGetConfiguredModelAlgorithm, "get-configured-model-algorithm", "", false, "Get Configured Model Algorithm")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlGetConfiguredModelAlgorithmAssociation, "get-configured-model-algorithm-association", "", false, "Get Configured Model Algorithm Association")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlGetMLConfiguration, "get-ml-configuration", "", false, "Get Ml Configuration")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlGetMLInputChannel, "get-ml-input-channel", "", false, "Get Ml Input Channel")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlGetTrainedModel, "get-trained-model", "", false, "Get Trained Model")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlGetTrainedModelInferenceJob, "get-trained-model-inference-job", "", false, "Get Trained Model Inference Job")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlGetTrainingDataset, "get-training-dataset", "", false, "Get Training Dataset")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlListAudienceExportJobs, "list-audience-export-jobs", "", false, "List Audience Export Jobs")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlListAudienceGenerationJobs, "list-audience-generation-jobs", "", false, "List Audience Generation Jobs")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlListAudienceModels, "list-audience-models", "", false, "List Audience Models")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlListCollaborationConfiguredModelAlgorithmAssociations, "list-collaboration-configured-model-algorithm-associations", "", false, "List Collaboration Configured Model Algorithm Associations")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlListCollaborationMLInputChannels, "list-collaboration-ml-input-channels", "", false, "List Collaboration Ml Input Channels")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlListCollaborationTrainedModelExportJobs, "list-collaboration-trained-model-export-jobs", "", false, "List Collaboration Trained Model Export Jobs")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlListCollaborationTrainedModelInferenceJobs, "list-collaboration-trained-model-inference-jobs", "", false, "List Collaboration Trained Model Inference Jobs")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlListCollaborationTrainedModels, "list-collaboration-trained-models", "", false, "List Collaboration Trained Models")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlListConfiguredAudienceModels, "list-configured-audience-models", "", false, "List Configured Audience Models")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlListConfiguredModelAlgorithmAssociations, "list-configured-model-algorithm-associations", "", false, "List Configured Model Algorithm Associations")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlListConfiguredModelAlgorithms, "list-configured-model-algorithms", "", false, "List Configured Model Algorithms")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlListMLInputChannels, "list-ml-input-channels", "", false, "List Ml Input Channels")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlListTrainedModelInferenceJobs, "list-trained-model-inference-jobs", "", false, "List Trained Model Inference Jobs")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlListTrainedModelVersions, "list-trained-model-versions", "", false, "List Trained Model Versions")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlListTrainedModels, "list-trained-models", "", false, "List Trained Models")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlListTrainingDatasets, "list-training-datasets", "", false, "List Training Datasets")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlPutConfiguredAudienceModelPolicy, "put-configured-audience-model-policy", "", false, "Put Configured Audience Model Policy")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlPutMLConfiguration, "put-ml-configuration", "", false, "Put Ml Configuration")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlStartAudienceExportJob, "start-audience-export-job", "", false, "Start Audience Export Job")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlStartAudienceGenerationJob, "start-audience-generation-job", "", false, "Start Audience Generation Job")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlStartTrainedModelExportJob, "start-trained-model-export-job", "", false, "Start Trained Model Export Job")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlStartTrainedModelInferenceJob, "start-trained-model-inference-job", "", false, "Start Trained Model Inference Job")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlTagResource, "tag-resource", "", false, "Tag Resource")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlUntagResource, "untag-resource", "", false, "Untag Resource")
	_cleanroomsmlCmd.Flags().BoolVarP(&_cleanroomsmlUpdateConfiguredAudienceModel, "update-configured-audience-model", "", false, "Update Configured Audience Model")

}
