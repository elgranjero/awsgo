package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ivsrealtime"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ivsrealtimeCmd represents the ivsrealtime command
var _ivsrealtimeCmd = &cobra.Command{
	Use:   "ivsrealtime",
	Short: "AWS ivsrealtime CLI",
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
		client := ivsrealtime.NewFromConfig(cfg)
		if _ivsrealtimeCreateEncoderConfiguration {
			ivsrealtime_CreateEncoderConfiguration(cfg, client)
			return
		}
		if _ivsrealtimeCreateIngestConfiguration {
			ivsrealtime_CreateIngestConfiguration(cfg, client)
			return
		}
		if _ivsrealtimeCreateParticipantToken {
			ivsrealtime_CreateParticipantToken(cfg, client)
			return
		}
		if _ivsrealtimeCreateStage {
			ivsrealtime_CreateStage(cfg, client)
			return
		}
		if _ivsrealtimeCreateStorageConfiguration {
			ivsrealtime_CreateStorageConfiguration(cfg, client)
			return
		}
		if _ivsrealtimeDeleteEncoderConfiguration {
			ivsrealtime_DeleteEncoderConfiguration(cfg, client)
			return
		}
		if _ivsrealtimeDeleteIngestConfiguration {
			ivsrealtime_DeleteIngestConfiguration(cfg, client)
			return
		}
		if _ivsrealtimeDeletePublicKey {
			ivsrealtime_DeletePublicKey(cfg, client)
			return
		}
		if _ivsrealtimeDeleteStage {
			ivsrealtime_DeleteStage(cfg, client)
			return
		}
		if _ivsrealtimeDeleteStorageConfiguration {
			ivsrealtime_DeleteStorageConfiguration(cfg, client)
			return
		}
		if _ivsrealtimeDisconnectParticipant {
			ivsrealtime_DisconnectParticipant(cfg, client)
			return
		}
		if _ivsrealtimeGetComposition {
			ivsrealtime_GetComposition(cfg, client)
			return
		}
		if _ivsrealtimeGetEncoderConfiguration {
			ivsrealtime_GetEncoderConfiguration(cfg, client)
			return
		}
		if _ivsrealtimeGetIngestConfiguration {
			ivsrealtime_GetIngestConfiguration(cfg, client)
			return
		}
		if _ivsrealtimeGetParticipant {
			ivsrealtime_GetParticipant(cfg, client)
			return
		}
		if _ivsrealtimeGetPublicKey {
			ivsrealtime_GetPublicKey(cfg, client)
			return
		}
		if _ivsrealtimeGetStage {
			ivsrealtime_GetStage(cfg, client)
			return
		}
		if _ivsrealtimeGetStageSession {
			ivsrealtime_GetStageSession(cfg, client)
			return
		}
		if _ivsrealtimeGetStorageConfiguration {
			ivsrealtime_GetStorageConfiguration(cfg, client)
			return
		}
		if _ivsrealtimeImportPublicKey {
			ivsrealtime_ImportPublicKey(cfg, client)
			return
		}
		if _ivsrealtimeListCompositions {
			ivsrealtime_ListCompositions(cfg, client)
			return
		}
		if _ivsrealtimeListEncoderConfigurations {
			ivsrealtime_ListEncoderConfigurations(cfg, client)
			return
		}
		if _ivsrealtimeListIngestConfigurations {
			ivsrealtime_ListIngestConfigurations(cfg, client)
			return
		}
		if _ivsrealtimeListParticipantEvents {
			ivsrealtime_ListParticipantEvents(cfg, client)
			return
		}
		if _ivsrealtimeListParticipantReplicas {
			ivsrealtime_ListParticipantReplicas(cfg, client)
			return
		}
		if _ivsrealtimeListParticipants {
			ivsrealtime_ListParticipants(cfg, client)
			return
		}
		if _ivsrealtimeListPublicKeys {
			ivsrealtime_ListPublicKeys(cfg, client)
			return
		}
		if _ivsrealtimeListStageSessions {
			ivsrealtime_ListStageSessions(cfg, client)
			return
		}
		if _ivsrealtimeListStages {
			ivsrealtime_ListStages(cfg, client)
			return
		}
		if _ivsrealtimeListStorageConfigurations {
			ivsrealtime_ListStorageConfigurations(cfg, client)
			return
		}
		if _ivsrealtimeListTagsForResource {
			ivsrealtime_ListTagsForResource(cfg, client)
			return
		}
		if _ivsrealtimeStartComposition {
			ivsrealtime_StartComposition(cfg, client)
			return
		}
		if _ivsrealtimeStartParticipantReplication {
			ivsrealtime_StartParticipantReplication(cfg, client)
			return
		}
		if _ivsrealtimeStopComposition {
			ivsrealtime_StopComposition(cfg, client)
			return
		}
		if _ivsrealtimeStopParticipantReplication {
			ivsrealtime_StopParticipantReplication(cfg, client)
			return
		}
		if _ivsrealtimeTagResource {
			ivsrealtime_TagResource(cfg, client)
			return
		}
		if _ivsrealtimeUntagResource {
			ivsrealtime_UntagResource(cfg, client)
			return
		}
		if _ivsrealtimeUpdateIngestConfiguration {
			ivsrealtime_UpdateIngestConfiguration(cfg, client)
			return
		}
		if _ivsrealtimeUpdateStage {
			ivsrealtime_UpdateStage(cfg, client)
			return
		}

	},
}

var (
	_ivsrealtimeCreateEncoderConfiguration  bool
	_ivsrealtimeCreateIngestConfiguration   bool
	_ivsrealtimeCreateParticipantToken      bool
	_ivsrealtimeCreateStage                 bool
	_ivsrealtimeCreateStorageConfiguration  bool
	_ivsrealtimeDeleteEncoderConfiguration  bool
	_ivsrealtimeDeleteIngestConfiguration   bool
	_ivsrealtimeDeletePublicKey             bool
	_ivsrealtimeDeleteStage                 bool
	_ivsrealtimeDeleteStorageConfiguration  bool
	_ivsrealtimeDisconnectParticipant       bool
	_ivsrealtimeGetComposition              bool
	_ivsrealtimeGetEncoderConfiguration     bool
	_ivsrealtimeGetIngestConfiguration      bool
	_ivsrealtimeGetParticipant              bool
	_ivsrealtimeGetPublicKey                bool
	_ivsrealtimeGetStage                    bool
	_ivsrealtimeGetStageSession             bool
	_ivsrealtimeGetStorageConfiguration     bool
	_ivsrealtimeImportPublicKey             bool
	_ivsrealtimeListCompositions            bool
	_ivsrealtimeListEncoderConfigurations   bool
	_ivsrealtimeListIngestConfigurations    bool
	_ivsrealtimeListParticipantEvents       bool
	_ivsrealtimeListParticipantReplicas     bool
	_ivsrealtimeListParticipants            bool
	_ivsrealtimeListPublicKeys              bool
	_ivsrealtimeListStageSessions           bool
	_ivsrealtimeListStages                  bool
	_ivsrealtimeListStorageConfigurations   bool
	_ivsrealtimeListTagsForResource         bool
	_ivsrealtimeStartComposition            bool
	_ivsrealtimeStartParticipantReplication bool
	_ivsrealtimeStopComposition             bool
	_ivsrealtimeStopParticipantReplication  bool
	_ivsrealtimeTagResource                 bool
	_ivsrealtimeUntagResource               bool
	_ivsrealtimeUpdateIngestConfiguration   bool
	_ivsrealtimeUpdateStage                 bool

	_ivsrealtimeArn                                   string
	_ivsrealtimeAttributes                            string
	_ivsrealtimeAutoParticipantRecordingConfiguration string
	_ivsrealtimeCapabilities                          string
	_ivsrealtimeDestinationStageArn                   string
	_ivsrealtimeDestinations                          string
	_ivsrealtimeDuration                              string
	_ivsrealtimeFilterByEncoderConfigurationArn       string
	_ivsrealtimeFilterByPublished                     string
	_ivsrealtimeFilterByRecordingState                string
	_ivsrealtimeFilterByStageArn                      string
	_ivsrealtimeFilterByState                         string
	_ivsrealtimeFilterByUserId                        string
	_ivsrealtimeForce                                 string
	_ivsrealtimeIdempotencyToken                      string
	_ivsrealtimeIngestProtocol                        string
	_ivsrealtimeInsecureIngest                        string
	_ivsrealtimeLayout                                string
	_ivsrealtimeMaxResults                            string
	_ivsrealtimeName                                  string
	_ivsrealtimeNextToken                             string
	_ivsrealtimeParticipantId                         string
	_ivsrealtimeParticipantTokenConfigurations        string
	_ivsrealtimePublicKeyMaterial                     string
	_ivsrealtimeReason                                string
	_ivsrealtimeReconnectWindowSeconds                string
	_ivsrealtimeResourceArn                           string
	_ivsrealtimeS3                                    string
	_ivsrealtimeSessionId                             string
	_ivsrealtimeSourceStageArn                        string
	_ivsrealtimeStageArn                              string
	_ivsrealtimeTagKeys                               []string
	_ivsrealtimeTags                                  string
	_ivsrealtimeUserId                                string
	_ivsrealtimeVideo                                 string
)

// Creates an EncoderConfiguration object.
func ivsrealtime_CreateEncoderConfiguration(cfg aws.Config, client *ivsrealtime.Client) {
	input := &ivsrealtime.CreateEncoderConfigurationInput{}

	if len(_ivsrealtimeName) > 0 {
		input.Name = aws.String(_ivsrealtimeName)
	}
	if len(_ivsrealtimeTags) > 0 {
		if err := assignInputField(input, "Tags", _ivsrealtimeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_ivsrealtimeVideo) > 0 {
		if err := assignInputField(input, "Video", _ivsrealtimeVideo); err != nil {
			log.Errorf("invalid --video: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateEncoderConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new IngestConfiguration resource, used to specify the ingest protocol
// for a stage.
func ivsrealtime_CreateIngestConfiguration(cfg aws.Config, client *ivsrealtime.Client) {
	input := &ivsrealtime.CreateIngestConfigurationInput{
		// IngestProtocol: types.IngestProtocol, // Required
	}

	if len(_ivsrealtimeIngestProtocol) > 0 {
		if err := assignInputField(input, "IngestProtocol", _ivsrealtimeIngestProtocol); err != nil {
			log.Errorf("invalid --ingest-protocol: %s", err.Error())
			return
		}
	}
	if len(_ivsrealtimeAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _ivsrealtimeAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_ivsrealtimeInsecureIngest) > 0 {
		if err := assignInputField(input, "InsecureIngest", _ivsrealtimeInsecureIngest); err != nil {
			log.Errorf("invalid --insecure-ingest: %s", err.Error())
			return
		}
	}
	if len(_ivsrealtimeName) > 0 {
		input.Name = aws.String(_ivsrealtimeName)
	}
	if len(_ivsrealtimeStageArn) > 0 {
		input.StageArn = aws.String(_ivsrealtimeStageArn)
	}
	if len(_ivsrealtimeTags) > 0 {
		if err := assignInputField(input, "Tags", _ivsrealtimeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_ivsrealtimeUserId) > 0 {
		input.UserId = aws.String(_ivsrealtimeUserId)
	}

	if resp, err := client.CreateIngestConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an additional token for a specified stage. This can be done after stage
// creation or when tokens expire. Tokens always are scoped to the stage for which
// they are created.
//
// Encryption keys are owned by Amazon IVS and never used directly by your
// application.
func ivsrealtime_CreateParticipantToken(cfg aws.Config, client *ivsrealtime.Client) {
	input := &ivsrealtime.CreateParticipantTokenInput{
		// StageArn: *string, // Required
	}

	if len(_ivsrealtimeStageArn) > 0 {
		input.StageArn = aws.String(_ivsrealtimeStageArn)
	}
	if len(_ivsrealtimeAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _ivsrealtimeAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_ivsrealtimeCapabilities) > 0 {
		if err := assignInputField(input, "Capabilities", _ivsrealtimeCapabilities); err != nil {
			log.Errorf("invalid --capabilities: %s", err.Error())
			return
		}
	}
	if len(_ivsrealtimeDuration) > 0 {
		if err := assignInputField(input, "Duration", _ivsrealtimeDuration); err != nil {
			log.Errorf("invalid --duration: %s", err.Error())
			return
		}
	}
	if len(_ivsrealtimeUserId) > 0 {
		input.UserId = aws.String(_ivsrealtimeUserId)
	}

	if resp, err := client.CreateParticipantToken(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new stage (and optionally participant tokens).
func ivsrealtime_CreateStage(cfg aws.Config, client *ivsrealtime.Client) {
	input := &ivsrealtime.CreateStageInput{}

	if len(_ivsrealtimeAutoParticipantRecordingConfiguration) > 0 {
		if err := assignInputField(input, "AutoParticipantRecordingConfiguration", _ivsrealtimeAutoParticipantRecordingConfiguration); err != nil {
			log.Errorf("invalid --auto-participant-recording-configuration: %s", err.Error())
			return
		}
	}
	if len(_ivsrealtimeName) > 0 {
		input.Name = aws.String(_ivsrealtimeName)
	}
	if len(_ivsrealtimeParticipantTokenConfigurations) > 0 {
		if err := assignInputField(input, "ParticipantTokenConfigurations", _ivsrealtimeParticipantTokenConfigurations); err != nil {
			log.Errorf("invalid --participant-token-configurations: %s", err.Error())
			return
		}
	}
	if len(_ivsrealtimeTags) > 0 {
		if err := assignInputField(input, "Tags", _ivsrealtimeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateStage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new storage configuration, used to enable recording to Amazon S3.
// When a StorageConfiguration is created, IVS will modify the S3 bucketPolicy of
// the provided bucket. This will ensure that IVS has sufficient permissions to
// write content to the provided bucket.
func ivsrealtime_CreateStorageConfiguration(cfg aws.Config, client *ivsrealtime.Client) {
	input := &ivsrealtime.CreateStorageConfigurationInput{
		// S3: *types.S3StorageConfiguration, // Required
	}

	if len(_ivsrealtimeS3) > 0 {
		if err := assignInputField(input, "S3", _ivsrealtimeS3); err != nil {
			log.Errorf("invalid --s3: %s", err.Error())
			return
		}
	}
	if len(_ivsrealtimeName) > 0 {
		input.Name = aws.String(_ivsrealtimeName)
	}
	if len(_ivsrealtimeTags) > 0 {
		if err := assignInputField(input, "Tags", _ivsrealtimeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateStorageConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an EncoderConfiguration resource. Ensures that no Compositions are
// using this template; otherwise, returns an error.
func ivsrealtime_DeleteEncoderConfiguration(cfg aws.Config, client *ivsrealtime.Client) {
	input := &ivsrealtime.DeleteEncoderConfigurationInput{
		// Arn: *string, // Required
	}

	if len(_ivsrealtimeArn) > 0 {
		input.Arn = aws.String(_ivsrealtimeArn)
	}

	if resp, err := client.DeleteEncoderConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specified IngestConfiguration, so it can no longer be used to
// broadcast. An IngestConfiguration cannot be deleted if the publisher is actively
// streaming to a stage, unless force is set to true .
func ivsrealtime_DeleteIngestConfiguration(cfg aws.Config, client *ivsrealtime.Client) {
	input := &ivsrealtime.DeleteIngestConfigurationInput{
		// Arn: *string, // Required
	}

	if len(_ivsrealtimeArn) > 0 {
		input.Arn = aws.String(_ivsrealtimeArn)
	}
	if len(_ivsrealtimeForce) > 0 {
		if err := assignInputField(input, "Force", _ivsrealtimeForce); err != nil {
			log.Errorf("invalid --force: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteIngestConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified public key used to sign stage participant tokens. This
// invalidates future participant tokens generated using the key pair’s private
// key.
func ivsrealtime_DeletePublicKey(cfg aws.Config, client *ivsrealtime.Client) {
	input := &ivsrealtime.DeletePublicKeyInput{
		// Arn: *string, // Required
	}

	if len(_ivsrealtimeArn) > 0 {
		input.Arn = aws.String(_ivsrealtimeArn)
	}

	if resp, err := client.DeletePublicKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Shuts down and deletes the specified stage (disconnecting all participants).
// This operation also removes the stageArn from the associated IngestConfiguration, if there are
// participants using the IngestConfiguration to publish to the stage.
func ivsrealtime_DeleteStage(cfg aws.Config, client *ivsrealtime.Client) {
	input := &ivsrealtime.DeleteStageInput{
		// Arn: *string, // Required
	}

	if len(_ivsrealtimeArn) > 0 {
		input.Arn = aws.String(_ivsrealtimeArn)
	}

	if resp, err := client.DeleteStage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the storage configuration for the specified ARN.
// If you try to delete a storage configuration that is used by a Composition, you
// will get an error (409 ConflictException). To avoid this, for all Compositions
// that reference the storage configuration, first use StopCompositionand wait for it to
// complete, then use DeleteStorageConfiguration.
func ivsrealtime_DeleteStorageConfiguration(cfg aws.Config, client *ivsrealtime.Client) {
	input := &ivsrealtime.DeleteStorageConfigurationInput{
		// Arn: *string, // Required
	}

	if len(_ivsrealtimeArn) > 0 {
		input.Arn = aws.String(_ivsrealtimeArn)
	}

	if resp, err := client.DeleteStorageConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disconnects a specified participant from a specified stage. If the participant
// is publishing using an IngestConfiguration, DisconnectParticipant also updates the stageArn in the
// IngestConfiguration to be an empty string.
func ivsrealtime_DisconnectParticipant(cfg aws.Config, client *ivsrealtime.Client) {
	input := &ivsrealtime.DisconnectParticipantInput{
		// ParticipantId: *string, // Required
		// StageArn: *string, // Required
	}

	if len(_ivsrealtimeParticipantId) > 0 {
		input.ParticipantId = aws.String(_ivsrealtimeParticipantId)
	}
	if len(_ivsrealtimeStageArn) > 0 {
		input.StageArn = aws.String(_ivsrealtimeStageArn)
	}
	if len(_ivsrealtimeReason) > 0 {
		input.Reason = aws.String(_ivsrealtimeReason)
	}

	if resp, err := client.DisconnectParticipant(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get information about the specified Composition resource.
func ivsrealtime_GetComposition(cfg aws.Config, client *ivsrealtime.Client) {
	input := &ivsrealtime.GetCompositionInput{
		// Arn: *string, // Required
	}

	if len(_ivsrealtimeArn) > 0 {
		input.Arn = aws.String(_ivsrealtimeArn)
	}

	if resp, err := client.GetComposition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the specified EncoderConfiguration resource.
func ivsrealtime_GetEncoderConfiguration(cfg aws.Config, client *ivsrealtime.Client) {
	input := &ivsrealtime.GetEncoderConfigurationInput{
		// Arn: *string, // Required
	}

	if len(_ivsrealtimeArn) > 0 {
		input.Arn = aws.String(_ivsrealtimeArn)
	}

	if resp, err := client.GetEncoderConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the specified IngestConfiguration.
func ivsrealtime_GetIngestConfiguration(cfg aws.Config, client *ivsrealtime.Client) {
	input := &ivsrealtime.GetIngestConfigurationInput{
		// Arn: *string, // Required
	}

	if len(_ivsrealtimeArn) > 0 {
		input.Arn = aws.String(_ivsrealtimeArn)
	}

	if resp, err := client.GetIngestConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the specified participant token.
func ivsrealtime_GetParticipant(cfg aws.Config, client *ivsrealtime.Client) {
	input := &ivsrealtime.GetParticipantInput{
		// ParticipantId: *string, // Required
		// SessionId: *string, // Required
		// StageArn: *string, // Required
	}

	if len(_ivsrealtimeParticipantId) > 0 {
		input.ParticipantId = aws.String(_ivsrealtimeParticipantId)
	}
	if len(_ivsrealtimeSessionId) > 0 {
		input.SessionId = aws.String(_ivsrealtimeSessionId)
	}
	if len(_ivsrealtimeStageArn) > 0 {
		input.StageArn = aws.String(_ivsrealtimeStageArn)
	}

	if resp, err := client.GetParticipant(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information for the specified public key.
func ivsrealtime_GetPublicKey(cfg aws.Config, client *ivsrealtime.Client) {
	input := &ivsrealtime.GetPublicKeyInput{
		// Arn: *string, // Required
	}

	if len(_ivsrealtimeArn) > 0 {
		input.Arn = aws.String(_ivsrealtimeArn)
	}

	if resp, err := client.GetPublicKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information for the specified stage.
func ivsrealtime_GetStage(cfg aws.Config, client *ivsrealtime.Client) {
	input := &ivsrealtime.GetStageInput{
		// Arn: *string, // Required
	}

	if len(_ivsrealtimeArn) > 0 {
		input.Arn = aws.String(_ivsrealtimeArn)
	}

	if resp, err := client.GetStage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information for the specified stage session.
func ivsrealtime_GetStageSession(cfg aws.Config, client *ivsrealtime.Client) {
	input := &ivsrealtime.GetStageSessionInput{
		// SessionId: *string, // Required
		// StageArn: *string, // Required
	}

	if len(_ivsrealtimeSessionId) > 0 {
		input.SessionId = aws.String(_ivsrealtimeSessionId)
	}
	if len(_ivsrealtimeStageArn) > 0 {
		input.StageArn = aws.String(_ivsrealtimeStageArn)
	}

	if resp, err := client.GetStageSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the storage configuration for the specified ARN.
func ivsrealtime_GetStorageConfiguration(cfg aws.Config, client *ivsrealtime.Client) {
	input := &ivsrealtime.GetStorageConfigurationInput{
		// Arn: *string, // Required
	}

	if len(_ivsrealtimeArn) > 0 {
		input.Arn = aws.String(_ivsrealtimeArn)
	}

	if resp, err := client.GetStorageConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Import a public key to be used for signing stage participant tokens.
func ivsrealtime_ImportPublicKey(cfg aws.Config, client *ivsrealtime.Client) {
	input := &ivsrealtime.ImportPublicKeyInput{
		// PublicKeyMaterial: *string, // Required
	}

	if len(_ivsrealtimePublicKeyMaterial) > 0 {
		input.PublicKeyMaterial = aws.String(_ivsrealtimePublicKeyMaterial)
	}
	if len(_ivsrealtimeName) > 0 {
		input.Name = aws.String(_ivsrealtimeName)
	}
	if len(_ivsrealtimeTags) > 0 {
		if err := assignInputField(input, "Tags", _ivsrealtimeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.ImportPublicKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets summary information about all Compositions in your account, in the AWS
// region where the API request is processed.
func ivsrealtime_ListCompositions(cfg aws.Config, client *ivsrealtime.Client) {
	input := &ivsrealtime.ListCompositionsInput{}

	if len(_ivsrealtimeFilterByEncoderConfigurationArn) > 0 {
		input.FilterByEncoderConfigurationArn = aws.String(_ivsrealtimeFilterByEncoderConfigurationArn)
	}
	if len(_ivsrealtimeFilterByStageArn) > 0 {
		input.FilterByStageArn = aws.String(_ivsrealtimeFilterByStageArn)
	}
	if len(_ivsrealtimeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ivsrealtimeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ivsrealtimeNextToken) > 0 {
		input.NextToken = aws.String(_ivsrealtimeNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCompositions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ivsrealtime.ListCompositionsOutput
	p := ivsrealtime.NewListCompositionsPaginator(client, input)
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

// Gets summary information about all EncoderConfigurations in your account, in
// the AWS region where the API request is processed.
func ivsrealtime_ListEncoderConfigurations(cfg aws.Config, client *ivsrealtime.Client) {
	input := &ivsrealtime.ListEncoderConfigurationsInput{}

	if len(_ivsrealtimeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ivsrealtimeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ivsrealtimeNextToken) > 0 {
		input.NextToken = aws.String(_ivsrealtimeNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEncoderConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ivsrealtime.ListEncoderConfigurationsOutput
	p := ivsrealtime.NewListEncoderConfigurationsPaginator(client, input)
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

// Lists all IngestConfigurations in your account, in the AWS region where the API
// request is processed.
func ivsrealtime_ListIngestConfigurations(cfg aws.Config, client *ivsrealtime.Client) {
	input := &ivsrealtime.ListIngestConfigurationsInput{}

	if len(_ivsrealtimeFilterByStageArn) > 0 {
		input.FilterByStageArn = aws.String(_ivsrealtimeFilterByStageArn)
	}
	if len(_ivsrealtimeFilterByState) > 0 {
		if err := assignInputField(input, "FilterByState", _ivsrealtimeFilterByState); err != nil {
			log.Errorf("invalid --filter-by-state: %s", err.Error())
			return
		}
	}
	if len(_ivsrealtimeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ivsrealtimeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ivsrealtimeNextToken) > 0 {
		input.NextToken = aws.String(_ivsrealtimeNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListIngestConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ivsrealtime.ListIngestConfigurationsOutput
	p := ivsrealtime.NewListIngestConfigurationsPaginator(client, input)
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

// Lists events for a specified participant that occurred during a specified stage
// session.
func ivsrealtime_ListParticipantEvents(cfg aws.Config, client *ivsrealtime.Client) {
	input := &ivsrealtime.ListParticipantEventsInput{
		// ParticipantId: *string, // Required
		// SessionId: *string, // Required
		// StageArn: *string, // Required
	}

	if len(_ivsrealtimeParticipantId) > 0 {
		input.ParticipantId = aws.String(_ivsrealtimeParticipantId)
	}
	if len(_ivsrealtimeSessionId) > 0 {
		input.SessionId = aws.String(_ivsrealtimeSessionId)
	}
	if len(_ivsrealtimeStageArn) > 0 {
		input.StageArn = aws.String(_ivsrealtimeStageArn)
	}
	if len(_ivsrealtimeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ivsrealtimeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ivsrealtimeNextToken) > 0 {
		input.NextToken = aws.String(_ivsrealtimeNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListParticipantEvents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ivsrealtime.ListParticipantEventsOutput
	p := ivsrealtime.NewListParticipantEventsPaginator(client, input)
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

// Lists all the replicas for a participant from a source stage.
func ivsrealtime_ListParticipantReplicas(cfg aws.Config, client *ivsrealtime.Client) {
	input := &ivsrealtime.ListParticipantReplicasInput{
		// ParticipantId: *string, // Required
		// SourceStageArn: *string, // Required
	}

	if len(_ivsrealtimeParticipantId) > 0 {
		input.ParticipantId = aws.String(_ivsrealtimeParticipantId)
	}
	if len(_ivsrealtimeSourceStageArn) > 0 {
		input.SourceStageArn = aws.String(_ivsrealtimeSourceStageArn)
	}
	if len(_ivsrealtimeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ivsrealtimeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ivsrealtimeNextToken) > 0 {
		input.NextToken = aws.String(_ivsrealtimeNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListParticipantReplicas(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ivsrealtime.ListParticipantReplicasOutput
	p := ivsrealtime.NewListParticipantReplicasPaginator(client, input)
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

// Lists all participants in a specified stage session.
func ivsrealtime_ListParticipants(cfg aws.Config, client *ivsrealtime.Client) {
	input := &ivsrealtime.ListParticipantsInput{
		// SessionId: *string, // Required
		// StageArn: *string, // Required
	}

	if len(_ivsrealtimeSessionId) > 0 {
		input.SessionId = aws.String(_ivsrealtimeSessionId)
	}
	if len(_ivsrealtimeStageArn) > 0 {
		input.StageArn = aws.String(_ivsrealtimeStageArn)
	}
	if len(_ivsrealtimeFilterByPublished) > 0 {
		if err := assignInputField(input, "FilterByPublished", _ivsrealtimeFilterByPublished); err != nil {
			log.Errorf("invalid --filter-by-published: %s", err.Error())
			return
		}
	}
	if len(_ivsrealtimeFilterByRecordingState) > 0 {
		if err := assignInputField(input, "FilterByRecordingState", _ivsrealtimeFilterByRecordingState); err != nil {
			log.Errorf("invalid --filter-by-recording-state: %s", err.Error())
			return
		}
	}
	if len(_ivsrealtimeFilterByState) > 0 {
		if err := assignInputField(input, "FilterByState", _ivsrealtimeFilterByState); err != nil {
			log.Errorf("invalid --filter-by-state: %s", err.Error())
			return
		}
	}
	if len(_ivsrealtimeFilterByUserId) > 0 {
		input.FilterByUserId = aws.String(_ivsrealtimeFilterByUserId)
	}
	if len(_ivsrealtimeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ivsrealtimeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ivsrealtimeNextToken) > 0 {
		input.NextToken = aws.String(_ivsrealtimeNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListParticipants(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ivsrealtime.ListParticipantsOutput
	p := ivsrealtime.NewListParticipantsPaginator(client, input)
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

// Gets summary information about all public keys in your account, in the AWS
// region where the API request is processed.
func ivsrealtime_ListPublicKeys(cfg aws.Config, client *ivsrealtime.Client) {
	input := &ivsrealtime.ListPublicKeysInput{}

	if len(_ivsrealtimeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ivsrealtimeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ivsrealtimeNextToken) > 0 {
		input.NextToken = aws.String(_ivsrealtimeNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPublicKeys(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ivsrealtime.ListPublicKeysOutput
	p := ivsrealtime.NewListPublicKeysPaginator(client, input)
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

// Gets all sessions for a specified stage.
func ivsrealtime_ListStageSessions(cfg aws.Config, client *ivsrealtime.Client) {
	input := &ivsrealtime.ListStageSessionsInput{
		// StageArn: *string, // Required
	}

	if len(_ivsrealtimeStageArn) > 0 {
		input.StageArn = aws.String(_ivsrealtimeStageArn)
	}
	if len(_ivsrealtimeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ivsrealtimeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ivsrealtimeNextToken) > 0 {
		input.NextToken = aws.String(_ivsrealtimeNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListStageSessions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ivsrealtime.ListStageSessionsOutput
	p := ivsrealtime.NewListStageSessionsPaginator(client, input)
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

// Gets summary information about all stages in your account, in the AWS region
// where the API request is processed.
func ivsrealtime_ListStages(cfg aws.Config, client *ivsrealtime.Client) {
	input := &ivsrealtime.ListStagesInput{}

	if len(_ivsrealtimeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ivsrealtimeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ivsrealtimeNextToken) > 0 {
		input.NextToken = aws.String(_ivsrealtimeNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListStages(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ivsrealtime.ListStagesOutput
	p := ivsrealtime.NewListStagesPaginator(client, input)
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

// Gets summary information about all storage configurations in your account, in
// the AWS region where the API request is processed.
func ivsrealtime_ListStorageConfigurations(cfg aws.Config, client *ivsrealtime.Client) {
	input := &ivsrealtime.ListStorageConfigurationsInput{}

	if len(_ivsrealtimeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ivsrealtimeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ivsrealtimeNextToken) > 0 {
		input.NextToken = aws.String(_ivsrealtimeNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListStorageConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ivsrealtime.ListStorageConfigurationsOutput
	p := ivsrealtime.NewListStorageConfigurationsPaginator(client, input)
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

// Gets information about AWS tags for the specified ARN.
func ivsrealtime_ListTagsForResource(cfg aws.Config, client *ivsrealtime.Client) {
	input := &ivsrealtime.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_ivsrealtimeResourceArn) > 0 {
		input.ResourceArn = aws.String(_ivsrealtimeResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a Composition from a stage based on the configuration provided in the
// request.
//
// A Composition is an ephemeral resource that exists after this operation returns
// successfully. Composition stops and the resource is deleted:
//
// - When StopCompositionis called.
//
// - After a 1-minute timeout, when all participants are disconnected from the
// stage.
//
// - After a 1-minute timeout, if there are no participants in the stage when
// StartComposition is called.
//
// - When broadcasting to the IVS channel fails and all retries are exhausted.
//
// - When broadcasting is disconnected and all attempts to reconnect are
// exhausted.
func ivsrealtime_StartComposition(cfg aws.Config, client *ivsrealtime.Client) {
	input := &ivsrealtime.StartCompositionInput{
		// Destinations: []types.DestinationConfiguration, // Required
		// StageArn: *string, // Required
	}

	if len(_ivsrealtimeDestinations) > 0 {
		if err := assignInputField(input, "Destinations", _ivsrealtimeDestinations); err != nil {
			log.Errorf("invalid --destinations: %s", err.Error())
			return
		}
	}
	if len(_ivsrealtimeStageArn) > 0 {
		input.StageArn = aws.String(_ivsrealtimeStageArn)
	}
	if len(_ivsrealtimeIdempotencyToken) > 0 {
		input.IdempotencyToken = aws.String(_ivsrealtimeIdempotencyToken)
	}
	if len(_ivsrealtimeLayout) > 0 {
		if err := assignInputField(input, "Layout", _ivsrealtimeLayout); err != nil {
			log.Errorf("invalid --layout: %s", err.Error())
			return
		}
	}
	if len(_ivsrealtimeTags) > 0 {
		if err := assignInputField(input, "Tags", _ivsrealtimeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartComposition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts replicating a publishing participant from a source stage to a
// destination stage.
func ivsrealtime_StartParticipantReplication(cfg aws.Config, client *ivsrealtime.Client) {
	input := &ivsrealtime.StartParticipantReplicationInput{
		// DestinationStageArn: *string, // Required
		// ParticipantId: *string, // Required
		// SourceStageArn: *string, // Required
	}

	if len(_ivsrealtimeDestinationStageArn) > 0 {
		input.DestinationStageArn = aws.String(_ivsrealtimeDestinationStageArn)
	}
	if len(_ivsrealtimeParticipantId) > 0 {
		input.ParticipantId = aws.String(_ivsrealtimeParticipantId)
	}
	if len(_ivsrealtimeSourceStageArn) > 0 {
		input.SourceStageArn = aws.String(_ivsrealtimeSourceStageArn)
	}
	if len(_ivsrealtimeAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _ivsrealtimeAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_ivsrealtimeReconnectWindowSeconds) > 0 {
		if err := assignInputField(input, "ReconnectWindowSeconds", _ivsrealtimeReconnectWindowSeconds); err != nil {
			log.Errorf("invalid --reconnect-window-seconds: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartParticipantReplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops and deletes a Composition resource. Any broadcast from the Composition
// resource is stopped.
func ivsrealtime_StopComposition(cfg aws.Config, client *ivsrealtime.Client) {
	input := &ivsrealtime.StopCompositionInput{
		// Arn: *string, // Required
	}

	if len(_ivsrealtimeArn) > 0 {
		input.Arn = aws.String(_ivsrealtimeArn)
	}

	if resp, err := client.StopComposition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a replicated participant session.
func ivsrealtime_StopParticipantReplication(cfg aws.Config, client *ivsrealtime.Client) {
	input := &ivsrealtime.StopParticipantReplicationInput{
		// DestinationStageArn: *string, // Required
		// ParticipantId: *string, // Required
		// SourceStageArn: *string, // Required
	}

	if len(_ivsrealtimeDestinationStageArn) > 0 {
		input.DestinationStageArn = aws.String(_ivsrealtimeDestinationStageArn)
	}
	if len(_ivsrealtimeParticipantId) > 0 {
		input.ParticipantId = aws.String(_ivsrealtimeParticipantId)
	}
	if len(_ivsrealtimeSourceStageArn) > 0 {
		input.SourceStageArn = aws.String(_ivsrealtimeSourceStageArn)
	}

	if resp, err := client.StopParticipantReplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or updates tags for the AWS resource with the specified ARN.
func ivsrealtime_TagResource(cfg aws.Config, client *ivsrealtime.Client) {
	input := &ivsrealtime.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_ivsrealtimeResourceArn) > 0 {
		input.ResourceArn = aws.String(_ivsrealtimeResourceArn)
	}
	if len(_ivsrealtimeTags) > 0 {
		if err := assignInputField(input, "Tags", _ivsrealtimeTags); err != nil {
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

// Removes tags from the resource with the specified ARN.
func ivsrealtime_UntagResource(cfg aws.Config, client *ivsrealtime.Client) {
	input := &ivsrealtime.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_ivsrealtimeResourceArn) > 0 {
		input.ResourceArn = aws.String(_ivsrealtimeResourceArn)
	}
	if len(_ivsrealtimeTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _ivsrealtimeTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a specified IngestConfiguration. Only the stage ARN attached to the
// IngestConfiguration can be updated. An IngestConfiguration that is active cannot
// be updated.
func ivsrealtime_UpdateIngestConfiguration(cfg aws.Config, client *ivsrealtime.Client) {
	input := &ivsrealtime.UpdateIngestConfigurationInput{
		// Arn: *string, // Required
	}

	if len(_ivsrealtimeArn) > 0 {
		input.Arn = aws.String(_ivsrealtimeArn)
	}
	if len(_ivsrealtimeStageArn) > 0 {
		input.StageArn = aws.String(_ivsrealtimeStageArn)
	}

	if resp, err := client.UpdateIngestConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a stage’s configuration.
func ivsrealtime_UpdateStage(cfg aws.Config, client *ivsrealtime.Client) {
	input := &ivsrealtime.UpdateStageInput{
		// Arn: *string, // Required
	}

	if len(_ivsrealtimeArn) > 0 {
		input.Arn = aws.String(_ivsrealtimeArn)
	}
	if len(_ivsrealtimeAutoParticipantRecordingConfiguration) > 0 {
		if err := assignInputField(input, "AutoParticipantRecordingConfiguration", _ivsrealtimeAutoParticipantRecordingConfiguration); err != nil {
			log.Errorf("invalid --auto-participant-recording-configuration: %s", err.Error())
			return
		}
	}
	if len(_ivsrealtimeName) > 0 {
		input.Name = aws.String(_ivsrealtimeName)
	}

	if resp, err := client.UpdateStage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_ivsrealtimeCmd)
	_ivsrealtimeCmd.Flags().SortFlags = false

	_ivsrealtimeCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_ivsrealtimeCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_ivsrealtimeCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_ivsrealtimeCmd.Flags().StringVarP(&_ivsrealtimeArn, "arn", "", "", "ARN")
	_ivsrealtimeCmd.Flags().StringVarP(&_ivsrealtimeAttributes, "attributes", "", "", "Attributes")
	_ivsrealtimeCmd.Flags().StringVarP(&_ivsrealtimeAutoParticipantRecordingConfiguration, "auto-participant-recording-configuration", "", "", "Auto Participant Recording Configuration")
	_ivsrealtimeCmd.Flags().StringVarP(&_ivsrealtimeCapabilities, "capabilities", "", "", "Capabilities")
	_ivsrealtimeCmd.Flags().StringVarP(&_ivsrealtimeDestinationStageArn, "destination-stage-arn", "", "", "Destination Stage ARN")
	_ivsrealtimeCmd.Flags().StringVarP(&_ivsrealtimeDestinations, "destinations", "", "", "Destinations")
	_ivsrealtimeCmd.Flags().StringVarP(&_ivsrealtimeDuration, "duration", "", "", "Duration")
	_ivsrealtimeCmd.Flags().StringVarP(&_ivsrealtimeFilterByEncoderConfigurationArn, "filter-by-encoder-configuration-arn", "", "", "Filter By Encoder Configuration ARN")
	_ivsrealtimeCmd.Flags().StringVarP(&_ivsrealtimeFilterByPublished, "filter-by-published", "", "", "Filter By Published")
	_ivsrealtimeCmd.Flags().StringVarP(&_ivsrealtimeFilterByRecordingState, "filter-by-recording-state", "", "", "Filter By Recording State")
	_ivsrealtimeCmd.Flags().StringVarP(&_ivsrealtimeFilterByStageArn, "filter-by-stage-arn", "", "", "Filter By Stage ARN")
	_ivsrealtimeCmd.Flags().StringVarP(&_ivsrealtimeFilterByState, "filter-by-state", "", "", "Filter By State")
	_ivsrealtimeCmd.Flags().StringVarP(&_ivsrealtimeFilterByUserId, "filter-by-user-id", "", "", "Filter By User ID")
	_ivsrealtimeCmd.Flags().StringVarP(&_ivsrealtimeForce, "force", "", "", "Force")
	_ivsrealtimeCmd.Flags().StringVarP(&_ivsrealtimeIdempotencyToken, "idempotency-token", "", "", "Idempotency Token")
	_ivsrealtimeCmd.Flags().StringVarP(&_ivsrealtimeIngestProtocol, "ingest-protocol", "", "", "Ingest Protocol")
	_ivsrealtimeCmd.Flags().StringVarP(&_ivsrealtimeInsecureIngest, "insecure-ingest", "", "", "Insecure Ingest")
	_ivsrealtimeCmd.Flags().StringVarP(&_ivsrealtimeLayout, "layout", "", "", "Layout")
	_ivsrealtimeCmd.Flags().StringVarP(&_ivsrealtimeMaxResults, "max-results", "", "", "Max Results")
	_ivsrealtimeCmd.Flags().StringVarP(&_ivsrealtimeName, "name", "", "", "Name")
	_ivsrealtimeCmd.Flags().StringVarP(&_ivsrealtimeNextToken, "next-token", "", "", "Next Token")
	_ivsrealtimeCmd.Flags().StringVarP(&_ivsrealtimeParticipantId, "participant-id", "", "", "Participant ID")
	_ivsrealtimeCmd.Flags().StringVarP(&_ivsrealtimeParticipantTokenConfigurations, "participant-token-configurations", "", "", "Participant Token Configurations")
	_ivsrealtimeCmd.Flags().StringVarP(&_ivsrealtimePublicKeyMaterial, "public-key-material", "", "", "Public Key Material")
	_ivsrealtimeCmd.Flags().StringVarP(&_ivsrealtimeReason, "reason", "", "", "Reason")
	_ivsrealtimeCmd.Flags().StringVarP(&_ivsrealtimeReconnectWindowSeconds, "reconnect-window-seconds", "", "", "Reconnect Window Seconds")
	_ivsrealtimeCmd.Flags().StringVarP(&_ivsrealtimeResourceArn, "resource-arn", "", "", "Resource ARN")
	_ivsrealtimeCmd.Flags().StringVarP(&_ivsrealtimeS3, "s3", "", "", "S3")
	_ivsrealtimeCmd.Flags().StringVarP(&_ivsrealtimeSessionId, "session-id", "", "", "Session ID")
	_ivsrealtimeCmd.Flags().StringVarP(&_ivsrealtimeSourceStageArn, "source-stage-arn", "", "", "Source Stage ARN")
	_ivsrealtimeCmd.Flags().StringVarP(&_ivsrealtimeStageArn, "stage-arn", "", "", "Stage ARN")
	_ivsrealtimeCmd.Flags().StringSliceVarP(&_ivsrealtimeTagKeys, "tag-keys", "", nil, "Tag Keys")
	_ivsrealtimeCmd.Flags().StringVarP(&_ivsrealtimeTags, "tags", "", "", "Tags")
	_ivsrealtimeCmd.Flags().StringVarP(&_ivsrealtimeUserId, "user-id", "", "", "User ID")
	_ivsrealtimeCmd.Flags().StringVarP(&_ivsrealtimeVideo, "video", "", "", "Video")

	_ivsrealtimeCmd.Flags().BoolVarP(&_ivsrealtimeCreateEncoderConfiguration, "create-encoder-configuration", "", false, "Create Encoder Configuration")
	_ivsrealtimeCmd.Flags().BoolVarP(&_ivsrealtimeCreateIngestConfiguration, "create-ingest-configuration", "", false, "Create Ingest Configuration")
	_ivsrealtimeCmd.Flags().BoolVarP(&_ivsrealtimeCreateParticipantToken, "create-participant-token", "", false, "Create Participant Token")
	_ivsrealtimeCmd.Flags().BoolVarP(&_ivsrealtimeCreateStage, "create-stage", "", false, "Create Stage")
	_ivsrealtimeCmd.Flags().BoolVarP(&_ivsrealtimeCreateStorageConfiguration, "create-storage-configuration", "", false, "Create Storage Configuration")
	_ivsrealtimeCmd.Flags().BoolVarP(&_ivsrealtimeDeleteEncoderConfiguration, "delete-encoder-configuration", "", false, "Delete Encoder Configuration")
	_ivsrealtimeCmd.Flags().BoolVarP(&_ivsrealtimeDeleteIngestConfiguration, "delete-ingest-configuration", "", false, "Delete Ingest Configuration")
	_ivsrealtimeCmd.Flags().BoolVarP(&_ivsrealtimeDeletePublicKey, "delete-public-key", "", false, "Delete Public Key")
	_ivsrealtimeCmd.Flags().BoolVarP(&_ivsrealtimeDeleteStage, "delete-stage", "", false, "Delete Stage")
	_ivsrealtimeCmd.Flags().BoolVarP(&_ivsrealtimeDeleteStorageConfiguration, "delete-storage-configuration", "", false, "Delete Storage Configuration")
	_ivsrealtimeCmd.Flags().BoolVarP(&_ivsrealtimeDisconnectParticipant, "disconnect-participant", "", false, "Disconnect Participant")
	_ivsrealtimeCmd.Flags().BoolVarP(&_ivsrealtimeGetComposition, "get-composition", "", false, "Get Composition")
	_ivsrealtimeCmd.Flags().BoolVarP(&_ivsrealtimeGetEncoderConfiguration, "get-encoder-configuration", "", false, "Get Encoder Configuration")
	_ivsrealtimeCmd.Flags().BoolVarP(&_ivsrealtimeGetIngestConfiguration, "get-ingest-configuration", "", false, "Get Ingest Configuration")
	_ivsrealtimeCmd.Flags().BoolVarP(&_ivsrealtimeGetParticipant, "get-participant", "", false, "Get Participant")
	_ivsrealtimeCmd.Flags().BoolVarP(&_ivsrealtimeGetPublicKey, "get-public-key", "", false, "Get Public Key")
	_ivsrealtimeCmd.Flags().BoolVarP(&_ivsrealtimeGetStage, "get-stage", "", false, "Get Stage")
	_ivsrealtimeCmd.Flags().BoolVarP(&_ivsrealtimeGetStageSession, "get-stage-session", "", false, "Get Stage Session")
	_ivsrealtimeCmd.Flags().BoolVarP(&_ivsrealtimeGetStorageConfiguration, "get-storage-configuration", "", false, "Get Storage Configuration")
	_ivsrealtimeCmd.Flags().BoolVarP(&_ivsrealtimeImportPublicKey, "import-public-key", "", false, "Import Public Key")
	_ivsrealtimeCmd.Flags().BoolVarP(&_ivsrealtimeListCompositions, "list-compositions", "", false, "List Compositions")
	_ivsrealtimeCmd.Flags().BoolVarP(&_ivsrealtimeListEncoderConfigurations, "list-encoder-configurations", "", false, "List Encoder Configurations")
	_ivsrealtimeCmd.Flags().BoolVarP(&_ivsrealtimeListIngestConfigurations, "list-ingest-configurations", "", false, "List Ingest Configurations")
	_ivsrealtimeCmd.Flags().BoolVarP(&_ivsrealtimeListParticipantEvents, "list-participant-events", "", false, "List Participant Events")
	_ivsrealtimeCmd.Flags().BoolVarP(&_ivsrealtimeListParticipantReplicas, "list-participant-replicas", "", false, "List Participant Replicas")
	_ivsrealtimeCmd.Flags().BoolVarP(&_ivsrealtimeListParticipants, "list-participants", "", false, "List Participants")
	_ivsrealtimeCmd.Flags().BoolVarP(&_ivsrealtimeListPublicKeys, "list-public-keys", "", false, "List Public Keys")
	_ivsrealtimeCmd.Flags().BoolVarP(&_ivsrealtimeListStageSessions, "list-stage-sessions", "", false, "List Stage Sessions")
	_ivsrealtimeCmd.Flags().BoolVarP(&_ivsrealtimeListStages, "list-stages", "", false, "List Stages")
	_ivsrealtimeCmd.Flags().BoolVarP(&_ivsrealtimeListStorageConfigurations, "list-storage-configurations", "", false, "List Storage Configurations")
	_ivsrealtimeCmd.Flags().BoolVarP(&_ivsrealtimeListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_ivsrealtimeCmd.Flags().BoolVarP(&_ivsrealtimeStartComposition, "start-composition", "", false, "Start Composition")
	_ivsrealtimeCmd.Flags().BoolVarP(&_ivsrealtimeStartParticipantReplication, "start-participant-replication", "", false, "Start Participant Replication")
	_ivsrealtimeCmd.Flags().BoolVarP(&_ivsrealtimeStopComposition, "stop-composition", "", false, "Stop Composition")
	_ivsrealtimeCmd.Flags().BoolVarP(&_ivsrealtimeStopParticipantReplication, "stop-participant-replication", "", false, "Stop Participant Replication")
	_ivsrealtimeCmd.Flags().BoolVarP(&_ivsrealtimeTagResource, "tag-resource", "", false, "Tag Resource")
	_ivsrealtimeCmd.Flags().BoolVarP(&_ivsrealtimeUntagResource, "untag-resource", "", false, "Untag Resource")
	_ivsrealtimeCmd.Flags().BoolVarP(&_ivsrealtimeUpdateIngestConfiguration, "update-ingest-configuration", "", false, "Update Ingest Configuration")
	_ivsrealtimeCmd.Flags().BoolVarP(&_ivsrealtimeUpdateStage, "update-stage", "", false, "Update Stage")

}
