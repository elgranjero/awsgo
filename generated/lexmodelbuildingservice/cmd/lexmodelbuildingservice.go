package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// lexmodelbuildingserviceCmd represents the lexmodelbuildingservice command
var _lexmodelbuildingserviceCmd = &cobra.Command{
	Use:   "lexmodelbuildingservice",
	Short: "AWS lexmodelbuildingservice CLI",
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
		client := lexmodelbuildingservice.NewFromConfig(cfg)
		if _lexmodelbuildingserviceCreateBotVersion {
			lexmodelbuildingservice_CreateBotVersion(cfg, client)
			return
		}
		if _lexmodelbuildingserviceCreateIntentVersion {
			lexmodelbuildingservice_CreateIntentVersion(cfg, client)
			return
		}
		if _lexmodelbuildingserviceCreateSlotTypeVersion {
			lexmodelbuildingservice_CreateSlotTypeVersion(cfg, client)
			return
		}
		if _lexmodelbuildingserviceDeleteBot {
			lexmodelbuildingservice_DeleteBot(cfg, client)
			return
		}
		if _lexmodelbuildingserviceDeleteBotAlias {
			lexmodelbuildingservice_DeleteBotAlias(cfg, client)
			return
		}
		if _lexmodelbuildingserviceDeleteBotChannelAssociation {
			lexmodelbuildingservice_DeleteBotChannelAssociation(cfg, client)
			return
		}
		if _lexmodelbuildingserviceDeleteBotVersion {
			lexmodelbuildingservice_DeleteBotVersion(cfg, client)
			return
		}
		if _lexmodelbuildingserviceDeleteIntent {
			lexmodelbuildingservice_DeleteIntent(cfg, client)
			return
		}
		if _lexmodelbuildingserviceDeleteIntentVersion {
			lexmodelbuildingservice_DeleteIntentVersion(cfg, client)
			return
		}
		if _lexmodelbuildingserviceDeleteSlotType {
			lexmodelbuildingservice_DeleteSlotType(cfg, client)
			return
		}
		if _lexmodelbuildingserviceDeleteSlotTypeVersion {
			lexmodelbuildingservice_DeleteSlotTypeVersion(cfg, client)
			return
		}
		if _lexmodelbuildingserviceDeleteUtterances {
			lexmodelbuildingservice_DeleteUtterances(cfg, client)
			return
		}
		if _lexmodelbuildingserviceGetBot {
			lexmodelbuildingservice_GetBot(cfg, client)
			return
		}
		if _lexmodelbuildingserviceGetBotAlias {
			lexmodelbuildingservice_GetBotAlias(cfg, client)
			return
		}
		if _lexmodelbuildingserviceGetBotAliases {
			lexmodelbuildingservice_GetBotAliases(cfg, client)
			return
		}
		if _lexmodelbuildingserviceGetBotChannelAssociation {
			lexmodelbuildingservice_GetBotChannelAssociation(cfg, client)
			return
		}
		if _lexmodelbuildingserviceGetBotChannelAssociations {
			lexmodelbuildingservice_GetBotChannelAssociations(cfg, client)
			return
		}
		if _lexmodelbuildingserviceGetBotVersions {
			lexmodelbuildingservice_GetBotVersions(cfg, client)
			return
		}
		if _lexmodelbuildingserviceGetBots {
			lexmodelbuildingservice_GetBots(cfg, client)
			return
		}
		if _lexmodelbuildingserviceGetBuiltinIntent {
			lexmodelbuildingservice_GetBuiltinIntent(cfg, client)
			return
		}
		if _lexmodelbuildingserviceGetBuiltinIntents {
			lexmodelbuildingservice_GetBuiltinIntents(cfg, client)
			return
		}
		if _lexmodelbuildingserviceGetBuiltinSlotTypes {
			lexmodelbuildingservice_GetBuiltinSlotTypes(cfg, client)
			return
		}
		if _lexmodelbuildingserviceGetExport {
			lexmodelbuildingservice_GetExport(cfg, client)
			return
		}
		if _lexmodelbuildingserviceGetImport {
			lexmodelbuildingservice_GetImport(cfg, client)
			return
		}
		if _lexmodelbuildingserviceGetIntent {
			lexmodelbuildingservice_GetIntent(cfg, client)
			return
		}
		if _lexmodelbuildingserviceGetIntentVersions {
			lexmodelbuildingservice_GetIntentVersions(cfg, client)
			return
		}
		if _lexmodelbuildingserviceGetIntents {
			lexmodelbuildingservice_GetIntents(cfg, client)
			return
		}
		if _lexmodelbuildingserviceGetMigration {
			lexmodelbuildingservice_GetMigration(cfg, client)
			return
		}
		if _lexmodelbuildingserviceGetMigrations {
			lexmodelbuildingservice_GetMigrations(cfg, client)
			return
		}
		if _lexmodelbuildingserviceGetSlotType {
			lexmodelbuildingservice_GetSlotType(cfg, client)
			return
		}
		if _lexmodelbuildingserviceGetSlotTypeVersions {
			lexmodelbuildingservice_GetSlotTypeVersions(cfg, client)
			return
		}
		if _lexmodelbuildingserviceGetSlotTypes {
			lexmodelbuildingservice_GetSlotTypes(cfg, client)
			return
		}
		if _lexmodelbuildingserviceGetUtterancesView {
			lexmodelbuildingservice_GetUtterancesView(cfg, client)
			return
		}
		if _lexmodelbuildingserviceListTagsForResource {
			lexmodelbuildingservice_ListTagsForResource(cfg, client)
			return
		}
		if _lexmodelbuildingservicePutBot {
			lexmodelbuildingservice_PutBot(cfg, client)
			return
		}
		if _lexmodelbuildingservicePutBotAlias {
			lexmodelbuildingservice_PutBotAlias(cfg, client)
			return
		}
		if _lexmodelbuildingservicePutIntent {
			lexmodelbuildingservice_PutIntent(cfg, client)
			return
		}
		if _lexmodelbuildingservicePutSlotType {
			lexmodelbuildingservice_PutSlotType(cfg, client)
			return
		}
		if _lexmodelbuildingserviceStartImport {
			lexmodelbuildingservice_StartImport(cfg, client)
			return
		}
		if _lexmodelbuildingserviceStartMigration {
			lexmodelbuildingservice_StartMigration(cfg, client)
			return
		}
		if _lexmodelbuildingserviceTagResource {
			lexmodelbuildingservice_TagResource(cfg, client)
			return
		}
		if _lexmodelbuildingserviceUntagResource {
			lexmodelbuildingservice_UntagResource(cfg, client)
			return
		}

	},
}

var (
	_lexmodelbuildingserviceCreateBotVersion            bool
	_lexmodelbuildingserviceCreateIntentVersion         bool
	_lexmodelbuildingserviceCreateSlotTypeVersion       bool
	_lexmodelbuildingserviceDeleteBot                   bool
	_lexmodelbuildingserviceDeleteBotAlias              bool
	_lexmodelbuildingserviceDeleteBotChannelAssociation bool
	_lexmodelbuildingserviceDeleteBotVersion            bool
	_lexmodelbuildingserviceDeleteIntent                bool
	_lexmodelbuildingserviceDeleteIntentVersion         bool
	_lexmodelbuildingserviceDeleteSlotType              bool
	_lexmodelbuildingserviceDeleteSlotTypeVersion       bool
	_lexmodelbuildingserviceDeleteUtterances            bool
	_lexmodelbuildingserviceGetBot                      bool
	_lexmodelbuildingserviceGetBotAlias                 bool
	_lexmodelbuildingserviceGetBotAliases               bool
	_lexmodelbuildingserviceGetBotChannelAssociation    bool
	_lexmodelbuildingserviceGetBotChannelAssociations   bool
	_lexmodelbuildingserviceGetBotVersions              bool
	_lexmodelbuildingserviceGetBots                     bool
	_lexmodelbuildingserviceGetBuiltinIntent            bool
	_lexmodelbuildingserviceGetBuiltinIntents           bool
	_lexmodelbuildingserviceGetBuiltinSlotTypes         bool
	_lexmodelbuildingserviceGetExport                   bool
	_lexmodelbuildingserviceGetImport                   bool
	_lexmodelbuildingserviceGetIntent                   bool
	_lexmodelbuildingserviceGetIntentVersions           bool
	_lexmodelbuildingserviceGetIntents                  bool
	_lexmodelbuildingserviceGetMigration                bool
	_lexmodelbuildingserviceGetMigrations               bool
	_lexmodelbuildingserviceGetSlotType                 bool
	_lexmodelbuildingserviceGetSlotTypeVersions         bool
	_lexmodelbuildingserviceGetSlotTypes                bool
	_lexmodelbuildingserviceGetUtterancesView           bool
	_lexmodelbuildingserviceListTagsForResource         bool
	_lexmodelbuildingservicePutBot                      bool
	_lexmodelbuildingservicePutBotAlias                 bool
	_lexmodelbuildingservicePutIntent                   bool
	_lexmodelbuildingservicePutSlotType                 bool
	_lexmodelbuildingserviceStartImport                 bool
	_lexmodelbuildingserviceStartMigration              bool
	_lexmodelbuildingserviceTagResource                 bool
	_lexmodelbuildingserviceUntagResource               bool

	_lexmodelbuildingserviceAbortStatement               string
	_lexmodelbuildingserviceBotAlias                     string
	_lexmodelbuildingserviceBotName                      string
	_lexmodelbuildingserviceBotVersion                   string
	_lexmodelbuildingserviceBotVersions                  []string
	_lexmodelbuildingserviceChecksum                     string
	_lexmodelbuildingserviceChildDirected                string
	_lexmodelbuildingserviceClarificationPrompt          string
	_lexmodelbuildingserviceConclusionStatement          string
	_lexmodelbuildingserviceConfirmationPrompt           string
	_lexmodelbuildingserviceConversationLogs             string
	_lexmodelbuildingserviceCreateVersion                string
	_lexmodelbuildingserviceDescription                  string
	_lexmodelbuildingserviceDetectSentiment              string
	_lexmodelbuildingserviceDialogCodeHook               string
	_lexmodelbuildingserviceEnableModelImprovements      string
	_lexmodelbuildingserviceEnumerationValues            string
	_lexmodelbuildingserviceExportType                   string
	_lexmodelbuildingserviceFollowUpPrompt               string
	_lexmodelbuildingserviceFulfillmentActivity          string
	_lexmodelbuildingserviceIdleSessionTTLInSeconds      string
	_lexmodelbuildingserviceImportId                     string
	_lexmodelbuildingserviceInputContexts                string
	_lexmodelbuildingserviceIntents                      string
	_lexmodelbuildingserviceKendraConfiguration          string
	_lexmodelbuildingserviceLocale                       string
	_lexmodelbuildingserviceMaxResults                   string
	_lexmodelbuildingserviceMergeStrategy                string
	_lexmodelbuildingserviceMigrationId                  string
	_lexmodelbuildingserviceMigrationStatusEquals        string
	_lexmodelbuildingserviceMigrationStrategy            string
	_lexmodelbuildingserviceName                         string
	_lexmodelbuildingserviceNameContains                 string
	_lexmodelbuildingserviceNextToken                    string
	_lexmodelbuildingserviceNluIntentConfidenceThreshold string
	_lexmodelbuildingserviceOutputContexts               string
	_lexmodelbuildingserviceParentIntentSignature        string
	_lexmodelbuildingserviceParentSlotTypeSignature      string
	_lexmodelbuildingservicePayload                      string
	_lexmodelbuildingserviceProcessBehavior              string
	_lexmodelbuildingserviceRejectionStatement           string
	_lexmodelbuildingserviceResourceArn                  string
	_lexmodelbuildingserviceResourceType                 string
	_lexmodelbuildingserviceSampleUtterances             []string
	_lexmodelbuildingserviceSignature                    string
	_lexmodelbuildingserviceSignatureContains            string
	_lexmodelbuildingserviceSlotTypeConfigurations       string
	_lexmodelbuildingserviceSlots                        string
	_lexmodelbuildingserviceSortByAttribute              string
	_lexmodelbuildingserviceSortByOrder                  string
	_lexmodelbuildingserviceStatusType                   string
	_lexmodelbuildingserviceTagKeys                      []string
	_lexmodelbuildingserviceTags                         string
	_lexmodelbuildingserviceUserId                       string
	_lexmodelbuildingserviceV1BotName                    string
	_lexmodelbuildingserviceV1BotNameContains            string
	_lexmodelbuildingserviceV1BotVersion                 string
	_lexmodelbuildingserviceV2BotName                    string
	_lexmodelbuildingserviceV2BotRole                    string
	_lexmodelbuildingserviceValueSelectionStrategy       string
	_lexmodelbuildingserviceVersion                      string
	_lexmodelbuildingserviceVersionOrAlias               string
	_lexmodelbuildingserviceVoiceId                      string
)

// Creates a new version of the bot based on the $LATEST version. If the $LATEST
// version of this resource hasn't changed since you created the last version,
// Amazon Lex doesn't create a new version. It returns the last created version.
//
// You can update only the $LATEST version of the bot. You can't update the
// numbered versions that you create with the CreateBotVersion operation.
//
// When you create the first version of a bot, Amazon Lex sets the version to 1.
// Subsequent versions increment by 1. For more information, see versioning-intro.
//
// This operation requires permission for the lex:CreateBotVersion action.
func lexmodelbuildingservice_CreateBotVersion(cfg aws.Config, client *lexmodelbuildingservice.Client) {
	input := &lexmodelbuildingservice.CreateBotVersionInput{
		// Name: *string, // Required
	}

	if len(_lexmodelbuildingserviceName) > 0 {
		input.Name = aws.String(_lexmodelbuildingserviceName)
	}
	if len(_lexmodelbuildingserviceChecksum) > 0 {
		input.Checksum = aws.String(_lexmodelbuildingserviceChecksum)
	}

	if resp, err := client.CreateBotVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new version of an intent based on the $LATEST version of the intent.
// If the $LATEST version of this intent hasn't changed since you last updated it,
// Amazon Lex doesn't create a new version. It returns the last version you
// created.
//
// You can update only the $LATEST version of the intent. You can't update the
// numbered versions that you create with the CreateIntentVersion operation.
//
// When you create a version of an intent, Amazon Lex sets the version to 1.
// Subsequent versions increment by 1. For more information, see versioning-intro.
//
// This operation requires permissions to perform the lex:CreateIntentVersion
// action.
func lexmodelbuildingservice_CreateIntentVersion(cfg aws.Config, client *lexmodelbuildingservice.Client) {
	input := &lexmodelbuildingservice.CreateIntentVersionInput{
		// Name: *string, // Required
	}

	if len(_lexmodelbuildingserviceName) > 0 {
		input.Name = aws.String(_lexmodelbuildingserviceName)
	}
	if len(_lexmodelbuildingserviceChecksum) > 0 {
		input.Checksum = aws.String(_lexmodelbuildingserviceChecksum)
	}

	if resp, err := client.CreateIntentVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new version of a slot type based on the $LATEST version of the
// specified slot type. If the $LATEST version of this resource has not changed
// since the last version that you created, Amazon Lex doesn't create a new
// version. It returns the last version that you created.
//
// You can update only the $LATEST version of a slot type. You can't update the
// numbered versions that you create with the CreateSlotTypeVersion operation.
//
// When you create a version of a slot type, Amazon Lex sets the version to 1.
// Subsequent versions increment by 1. For more information, see versioning-intro.
//
// This operation requires permissions for the lex:CreateSlotTypeVersion action.
func lexmodelbuildingservice_CreateSlotTypeVersion(cfg aws.Config, client *lexmodelbuildingservice.Client) {
	input := &lexmodelbuildingservice.CreateSlotTypeVersionInput{
		// Name: *string, // Required
	}

	if len(_lexmodelbuildingserviceName) > 0 {
		input.Name = aws.String(_lexmodelbuildingserviceName)
	}
	if len(_lexmodelbuildingserviceChecksum) > 0 {
		input.Checksum = aws.String(_lexmodelbuildingserviceChecksum)
	}

	if resp, err := client.CreateSlotTypeVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes all versions of the bot, including the $LATEST version. To delete a
// specific version of the bot, use the DeleteBotVersionoperation. The DeleteBot operation doesn't
// immediately remove the bot schema. Instead, it is marked for deletion and
// removed later.
//
// Amazon Lex stores utterances indefinitely for improving the ability of your bot
// to respond to user inputs. These utterances are not removed when the bot is
// deleted. To remove the utterances, use the DeleteUtterancesoperation.
//
// If a bot has an alias, you can't delete it. Instead, the DeleteBot operation
// returns a ResourceInUseException exception that includes a reference to the
// alias that refers to the bot. To remove the reference to the bot, delete the
// alias. If you get the same exception again, delete the referring alias until the
// DeleteBot operation is successful.
//
// This operation requires permissions for the lex:DeleteBot action.
func lexmodelbuildingservice_DeleteBot(cfg aws.Config, client *lexmodelbuildingservice.Client) {
	input := &lexmodelbuildingservice.DeleteBotInput{
		// Name: *string, // Required
	}

	if len(_lexmodelbuildingserviceName) > 0 {
		input.Name = aws.String(_lexmodelbuildingserviceName)
	}

	if resp, err := client.DeleteBot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an alias for the specified bot.
// You can't delete an alias that is used in the association between a bot and a
// messaging channel. If an alias is used in a channel association, the DeleteBot
// operation returns a ResourceInUseException exception that includes a reference
// to the channel association that refers to the bot. You can remove the reference
// to the alias by deleting the channel association. If you get the same exception
// again, delete the referring association until the DeleteBotAlias operation is
// successful.
func lexmodelbuildingservice_DeleteBotAlias(cfg aws.Config, client *lexmodelbuildingservice.Client) {
	input := &lexmodelbuildingservice.DeleteBotAliasInput{
		// BotName: *string, // Required
		// Name: *string, // Required
	}

	if len(_lexmodelbuildingserviceBotName) > 0 {
		input.BotName = aws.String(_lexmodelbuildingserviceBotName)
	}
	if len(_lexmodelbuildingserviceName) > 0 {
		input.Name = aws.String(_lexmodelbuildingserviceName)
	}

	if resp, err := client.DeleteBotAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the association between an Amazon Lex bot and a messaging platform.
// This operation requires permission for the lex:DeleteBotChannelAssociation
// action.
func lexmodelbuildingservice_DeleteBotChannelAssociation(cfg aws.Config, client *lexmodelbuildingservice.Client) {
	input := &lexmodelbuildingservice.DeleteBotChannelAssociationInput{
		// BotAlias: *string, // Required
		// BotName: *string, // Required
		// Name: *string, // Required
	}

	if len(_lexmodelbuildingserviceBotAlias) > 0 {
		input.BotAlias = aws.String(_lexmodelbuildingserviceBotAlias)
	}
	if len(_lexmodelbuildingserviceBotName) > 0 {
		input.BotName = aws.String(_lexmodelbuildingserviceBotName)
	}
	if len(_lexmodelbuildingserviceName) > 0 {
		input.Name = aws.String(_lexmodelbuildingserviceName)
	}

	if resp, err := client.DeleteBotChannelAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specific version of a bot. To delete all versions of a bot, use the DeleteBot
// operation.
//
// This operation requires permissions for the lex:DeleteBotVersion action.
func lexmodelbuildingservice_DeleteBotVersion(cfg aws.Config, client *lexmodelbuildingservice.Client) {
	input := &lexmodelbuildingservice.DeleteBotVersionInput{
		// Name: *string, // Required
		// Version: *string, // Required
	}

	if len(_lexmodelbuildingserviceName) > 0 {
		input.Name = aws.String(_lexmodelbuildingserviceName)
	}
	if len(_lexmodelbuildingserviceVersion) > 0 {
		input.Version = aws.String(_lexmodelbuildingserviceVersion)
	}

	if resp, err := client.DeleteBotVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes all versions of the intent, including the $LATEST version. To delete a
// specific version of the intent, use the DeleteIntentVersionoperation.
//
// You can delete a version of an intent only if it is not referenced. To delete
// an intent that is referred to in one or more bots (see how-it-works), you must remove those
// references first.
//
// If you get the ResourceInUseException exception, it provides an example
// reference that shows where the intent is referenced. To remove the reference to
// the intent, either update the bot or delete it. If you get the same exception
// when you attempt to delete the intent again, repeat until the intent has no
// references and the call to DeleteIntent is successful.
//
// This operation requires permission for the lex:DeleteIntent action.
func lexmodelbuildingservice_DeleteIntent(cfg aws.Config, client *lexmodelbuildingservice.Client) {
	input := &lexmodelbuildingservice.DeleteIntentInput{
		// Name: *string, // Required
	}

	if len(_lexmodelbuildingserviceName) > 0 {
		input.Name = aws.String(_lexmodelbuildingserviceName)
	}

	if resp, err := client.DeleteIntent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specific version of an intent. To delete all versions of a intent,
// use the DeleteIntentoperation.
//
// This operation requires permissions for the lex:DeleteIntentVersion action.
func lexmodelbuildingservice_DeleteIntentVersion(cfg aws.Config, client *lexmodelbuildingservice.Client) {
	input := &lexmodelbuildingservice.DeleteIntentVersionInput{
		// Name: *string, // Required
		// Version: *string, // Required
	}

	if len(_lexmodelbuildingserviceName) > 0 {
		input.Name = aws.String(_lexmodelbuildingserviceName)
	}
	if len(_lexmodelbuildingserviceVersion) > 0 {
		input.Version = aws.String(_lexmodelbuildingserviceVersion)
	}

	if resp, err := client.DeleteIntentVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes all versions of the slot type, including the $LATEST version. To delete
// a specific version of the slot type, use the DeleteSlotTypeVersionoperation.
//
// You can delete a version of a slot type only if it is not referenced. To delete
// a slot type that is referred to in one or more intents, you must remove those
// references first.
//
// If you get the ResourceInUseException exception, the exception provides an
// example reference that shows the intent where the slot type is referenced. To
// remove the reference to the slot type, either update the intent or delete it. If
// you get the same exception when you attempt to delete the slot type again,
// repeat until the slot type has no references and the DeleteSlotType call is
// successful.
//
// This operation requires permission for the lex:DeleteSlotType action.
func lexmodelbuildingservice_DeleteSlotType(cfg aws.Config, client *lexmodelbuildingservice.Client) {
	input := &lexmodelbuildingservice.DeleteSlotTypeInput{
		// Name: *string, // Required
	}

	if len(_lexmodelbuildingserviceName) > 0 {
		input.Name = aws.String(_lexmodelbuildingserviceName)
	}

	if resp, err := client.DeleteSlotType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specific version of a slot type. To delete all versions of a slot
// type, use the DeleteSlotTypeoperation.
//
// This operation requires permissions for the lex:DeleteSlotTypeVersion action.
func lexmodelbuildingservice_DeleteSlotTypeVersion(cfg aws.Config, client *lexmodelbuildingservice.Client) {
	input := &lexmodelbuildingservice.DeleteSlotTypeVersionInput{
		// Name: *string, // Required
		// Version: *string, // Required
	}

	if len(_lexmodelbuildingserviceName) > 0 {
		input.Name = aws.String(_lexmodelbuildingserviceName)
	}
	if len(_lexmodelbuildingserviceVersion) > 0 {
		input.Version = aws.String(_lexmodelbuildingserviceVersion)
	}

	if resp, err := client.DeleteSlotTypeVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes stored utterances.
// Amazon Lex stores the utterances that users send to your bot. Utterances are
// stored for 15 days for use with the GetUtterancesViewoperation, and then stored indefinitely for
// use in improving the ability of your bot to respond to user input.
//
// Use the DeleteUtterances operation to manually delete stored utterances for a
// specific user. When you use the DeleteUtterances operation, utterances stored
// for improving your bot's ability to respond to user input are deleted
// immediately. Utterances stored for use with the GetUtterancesView operation are
// deleted after 15 days.
//
// This operation requires permissions for the lex:DeleteUtterances action.
func lexmodelbuildingservice_DeleteUtterances(cfg aws.Config, client *lexmodelbuildingservice.Client) {
	input := &lexmodelbuildingservice.DeleteUtterancesInput{
		// BotName: *string, // Required
		// UserId: *string, // Required
	}

	if len(_lexmodelbuildingserviceBotName) > 0 {
		input.BotName = aws.String(_lexmodelbuildingserviceBotName)
	}
	if len(_lexmodelbuildingserviceUserId) > 0 {
		input.UserId = aws.String(_lexmodelbuildingserviceUserId)
	}

	if resp, err := client.DeleteUtterances(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns metadata information for a specific bot. You must provide the bot name
// and the bot version or alias.
//
// This operation requires permissions for the lex:GetBot action.
func lexmodelbuildingservice_GetBot(cfg aws.Config, client *lexmodelbuildingservice.Client) {
	input := &lexmodelbuildingservice.GetBotInput{
		// Name: *string, // Required
		// VersionOrAlias: *string, // Required
	}

	if len(_lexmodelbuildingserviceName) > 0 {
		input.Name = aws.String(_lexmodelbuildingserviceName)
	}
	if len(_lexmodelbuildingserviceVersionOrAlias) > 0 {
		input.VersionOrAlias = aws.String(_lexmodelbuildingserviceVersionOrAlias)
	}

	if resp, err := client.GetBot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about an Amazon Lex bot alias. For more information about
// aliases, see versioning-aliases.
//
// This operation requires permissions for the lex:GetBotAlias action.
func lexmodelbuildingservice_GetBotAlias(cfg aws.Config, client *lexmodelbuildingservice.Client) {
	input := &lexmodelbuildingservice.GetBotAliasInput{
		// BotName: *string, // Required
		// Name: *string, // Required
	}

	if len(_lexmodelbuildingserviceBotName) > 0 {
		input.BotName = aws.String(_lexmodelbuildingserviceBotName)
	}
	if len(_lexmodelbuildingserviceName) > 0 {
		input.Name = aws.String(_lexmodelbuildingserviceName)
	}

	if resp, err := client.GetBotAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of aliases for a specified Amazon Lex bot.
// This operation requires permissions for the lex:GetBotAliases action.
func lexmodelbuildingservice_GetBotAliases(cfg aws.Config, client *lexmodelbuildingservice.Client) {
	input := &lexmodelbuildingservice.GetBotAliasesInput{
		// BotName: *string, // Required
	}

	if len(_lexmodelbuildingserviceBotName) > 0 {
		input.BotName = aws.String(_lexmodelbuildingserviceBotName)
	}
	if len(_lexmodelbuildingserviceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lexmodelbuildingserviceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lexmodelbuildingserviceNameContains) > 0 {
		input.NameContains = aws.String(_lexmodelbuildingserviceNameContains)
	}
	if len(_lexmodelbuildingserviceNextToken) > 0 {
		input.NextToken = aws.String(_lexmodelbuildingserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetBotAliases(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lexmodelbuildingservice.GetBotAliasesOutput
	p := lexmodelbuildingservice.NewGetBotAliasesPaginator(client, input)
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

// Returns information about the association between an Amazon Lex bot and a
// messaging platform.
//
// This operation requires permissions for the lex:GetBotChannelAssociation action.
func lexmodelbuildingservice_GetBotChannelAssociation(cfg aws.Config, client *lexmodelbuildingservice.Client) {
	input := &lexmodelbuildingservice.GetBotChannelAssociationInput{
		// BotAlias: *string, // Required
		// BotName: *string, // Required
		// Name: *string, // Required
	}

	if len(_lexmodelbuildingserviceBotAlias) > 0 {
		input.BotAlias = aws.String(_lexmodelbuildingserviceBotAlias)
	}
	if len(_lexmodelbuildingserviceBotName) > 0 {
		input.BotName = aws.String(_lexmodelbuildingserviceBotName)
	}
	if len(_lexmodelbuildingserviceName) > 0 {
		input.Name = aws.String(_lexmodelbuildingserviceName)
	}

	if resp, err := client.GetBotChannelAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of all of the channels associated with the specified bot.
// The GetBotChannelAssociations operation requires permissions for the
// lex:GetBotChannelAssociations action.
func lexmodelbuildingservice_GetBotChannelAssociations(cfg aws.Config, client *lexmodelbuildingservice.Client) {
	input := &lexmodelbuildingservice.GetBotChannelAssociationsInput{
		// BotAlias: *string, // Required
		// BotName: *string, // Required
	}

	if len(_lexmodelbuildingserviceBotAlias) > 0 {
		input.BotAlias = aws.String(_lexmodelbuildingserviceBotAlias)
	}
	if len(_lexmodelbuildingserviceBotName) > 0 {
		input.BotName = aws.String(_lexmodelbuildingserviceBotName)
	}
	if len(_lexmodelbuildingserviceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lexmodelbuildingserviceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lexmodelbuildingserviceNameContains) > 0 {
		input.NameContains = aws.String(_lexmodelbuildingserviceNameContains)
	}
	if len(_lexmodelbuildingserviceNextToken) > 0 {
		input.NextToken = aws.String(_lexmodelbuildingserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetBotChannelAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lexmodelbuildingservice.GetBotChannelAssociationsOutput
	p := lexmodelbuildingservice.NewGetBotChannelAssociationsPaginator(client, input)
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

// Gets information about all of the versions of a bot.
// The GetBotVersions operation returns a BotMetadata object for each version of a
// bot. For example, if a bot has three numbered versions, the GetBotVersions
// operation returns four BotMetadata objects in the response, one for each
// numbered version and one for the $LATEST version.
//
// The GetBotVersions operation always returns at least one version, the $LATEST
// version.
//
// This operation requires permissions for the lex:GetBotVersions action.
func lexmodelbuildingservice_GetBotVersions(cfg aws.Config, client *lexmodelbuildingservice.Client) {
	input := &lexmodelbuildingservice.GetBotVersionsInput{
		// Name: *string, // Required
	}

	if len(_lexmodelbuildingserviceName) > 0 {
		input.Name = aws.String(_lexmodelbuildingserviceName)
	}
	if len(_lexmodelbuildingserviceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lexmodelbuildingserviceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lexmodelbuildingserviceNextToken) > 0 {
		input.NextToken = aws.String(_lexmodelbuildingserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetBotVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lexmodelbuildingservice.GetBotVersionsOutput
	p := lexmodelbuildingservice.NewGetBotVersionsPaginator(client, input)
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

// Returns bot information as follows:
// - If you provide the nameContains field, the response includes information for
// the $LATEST version of all bots whose name contains the specified string.
//
// - If you don't specify the nameContains field, the operation returns
// information about the $LATEST version of all of your bots.
//
// This operation requires permission for the lex:GetBots action.
func lexmodelbuildingservice_GetBots(cfg aws.Config, client *lexmodelbuildingservice.Client) {
	input := &lexmodelbuildingservice.GetBotsInput{}

	if len(_lexmodelbuildingserviceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lexmodelbuildingserviceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lexmodelbuildingserviceNameContains) > 0 {
		input.NameContains = aws.String(_lexmodelbuildingserviceNameContains)
	}
	if len(_lexmodelbuildingserviceNextToken) > 0 {
		input.NextToken = aws.String(_lexmodelbuildingserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetBots(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lexmodelbuildingservice.GetBotsOutput
	p := lexmodelbuildingservice.NewGetBotsPaginator(client, input)
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

// Returns information about a built-in intent.
// This operation requires permission for the lex:GetBuiltinIntent action.
func lexmodelbuildingservice_GetBuiltinIntent(cfg aws.Config, client *lexmodelbuildingservice.Client) {
	input := &lexmodelbuildingservice.GetBuiltinIntentInput{
		// Signature: *string, // Required
	}

	if len(_lexmodelbuildingserviceSignature) > 0 {
		input.Signature = aws.String(_lexmodelbuildingserviceSignature)
	}

	if resp, err := client.GetBuiltinIntent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a list of built-in intents that meet the specified criteria.
// This operation requires permission for the lex:GetBuiltinIntents action.
func lexmodelbuildingservice_GetBuiltinIntents(cfg aws.Config, client *lexmodelbuildingservice.Client) {
	input := &lexmodelbuildingservice.GetBuiltinIntentsInput{}

	if len(_lexmodelbuildingserviceLocale) > 0 {
		if err := assignInputField(input, "Locale", _lexmodelbuildingserviceLocale); err != nil {
			log.Errorf("invalid --locale: %s", err.Error())
			return
		}
	}
	if len(_lexmodelbuildingserviceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lexmodelbuildingserviceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lexmodelbuildingserviceNextToken) > 0 {
		input.NextToken = aws.String(_lexmodelbuildingserviceNextToken)
	}
	if len(_lexmodelbuildingserviceSignatureContains) > 0 {
		input.SignatureContains = aws.String(_lexmodelbuildingserviceSignatureContains)
	}

	if disablePaginator() {
		if resp, err := client.GetBuiltinIntents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lexmodelbuildingservice.GetBuiltinIntentsOutput
	p := lexmodelbuildingservice.NewGetBuiltinIntentsPaginator(client, input)
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

// Gets a list of built-in slot types that meet the specified criteria.
// For a list of built-in slot types, see [Slot Type Reference] in the Alexa Skills Kit.
//
// This operation requires permission for the lex:GetBuiltInSlotTypes action.
//
// [Slot Type Reference]: https://developer.amazon.com/public/solutions/alexa/alexa-skills-kit/docs/built-in-intent-ref/slot-type-reference
func lexmodelbuildingservice_GetBuiltinSlotTypes(cfg aws.Config, client *lexmodelbuildingservice.Client) {
	input := &lexmodelbuildingservice.GetBuiltinSlotTypesInput{}

	if len(_lexmodelbuildingserviceLocale) > 0 {
		if err := assignInputField(input, "Locale", _lexmodelbuildingserviceLocale); err != nil {
			log.Errorf("invalid --locale: %s", err.Error())
			return
		}
	}
	if len(_lexmodelbuildingserviceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lexmodelbuildingserviceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lexmodelbuildingserviceNextToken) > 0 {
		input.NextToken = aws.String(_lexmodelbuildingserviceNextToken)
	}
	if len(_lexmodelbuildingserviceSignatureContains) > 0 {
		input.SignatureContains = aws.String(_lexmodelbuildingserviceSignatureContains)
	}

	if disablePaginator() {
		if resp, err := client.GetBuiltinSlotTypes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lexmodelbuildingservice.GetBuiltinSlotTypesOutput
	p := lexmodelbuildingservice.NewGetBuiltinSlotTypesPaginator(client, input)
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

// Exports the contents of a Amazon Lex resource in a specified format.
func lexmodelbuildingservice_GetExport(cfg aws.Config, client *lexmodelbuildingservice.Client) {
	input := &lexmodelbuildingservice.GetExportInput{
		// ExportType: types.ExportType, // Required
		// Name: *string, // Required
		// ResourceType: types.ResourceType, // Required
		// Version: *string, // Required
	}

	if len(_lexmodelbuildingserviceExportType) > 0 {
		if err := assignInputField(input, "ExportType", _lexmodelbuildingserviceExportType); err != nil {
			log.Errorf("invalid --export-type: %s", err.Error())
			return
		}
	}
	if len(_lexmodelbuildingserviceName) > 0 {
		input.Name = aws.String(_lexmodelbuildingserviceName)
	}
	if len(_lexmodelbuildingserviceResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _lexmodelbuildingserviceResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}
	if len(_lexmodelbuildingserviceVersion) > 0 {
		input.Version = aws.String(_lexmodelbuildingserviceVersion)
	}

	if resp, err := client.GetExport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about an import job started with the StartImport operation.
func lexmodelbuildingservice_GetImport(cfg aws.Config, client *lexmodelbuildingservice.Client) {
	input := &lexmodelbuildingservice.GetImportInput{
		// ImportId: *string, // Required
	}

	if len(_lexmodelbuildingserviceImportId) > 0 {
		input.ImportId = aws.String(_lexmodelbuildingserviceImportId)
	}

	if resp, err := client.GetImport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about an intent. In addition to the intent name, you must
// specify the intent version.
//
// This operation requires permissions to perform the lex:GetIntent action.
func lexmodelbuildingservice_GetIntent(cfg aws.Config, client *lexmodelbuildingservice.Client) {
	input := &lexmodelbuildingservice.GetIntentInput{
		// Name: *string, // Required
		// Version: *string, // Required
	}

	if len(_lexmodelbuildingserviceName) > 0 {
		input.Name = aws.String(_lexmodelbuildingserviceName)
	}
	if len(_lexmodelbuildingserviceVersion) > 0 {
		input.Version = aws.String(_lexmodelbuildingserviceVersion)
	}

	if resp, err := client.GetIntent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about all of the versions of an intent.
// The GetIntentVersions operation returns an IntentMetadata object for each
// version of an intent. For example, if an intent has three numbered versions, the
// GetIntentVersions operation returns four IntentMetadata objects in the
// response, one for each numbered version and one for the $LATEST version.
//
// The GetIntentVersions operation always returns at least one version, the
// $LATEST version.
//
// This operation requires permissions for the lex:GetIntentVersions action.
func lexmodelbuildingservice_GetIntentVersions(cfg aws.Config, client *lexmodelbuildingservice.Client) {
	input := &lexmodelbuildingservice.GetIntentVersionsInput{
		// Name: *string, // Required
	}

	if len(_lexmodelbuildingserviceName) > 0 {
		input.Name = aws.String(_lexmodelbuildingserviceName)
	}
	if len(_lexmodelbuildingserviceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lexmodelbuildingserviceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lexmodelbuildingserviceNextToken) > 0 {
		input.NextToken = aws.String(_lexmodelbuildingserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetIntentVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lexmodelbuildingservice.GetIntentVersionsOutput
	p := lexmodelbuildingservice.NewGetIntentVersionsPaginator(client, input)
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

// Returns intent information as follows:
// - If you specify the nameContains field, returns the $LATEST version of all
// intents that contain the specified string.
//
// - If you don't specify the nameContains field, returns information about the
// $LATEST version of all intents.
//
// The operation requires permission for the lex:GetIntents action.
func lexmodelbuildingservice_GetIntents(cfg aws.Config, client *lexmodelbuildingservice.Client) {
	input := &lexmodelbuildingservice.GetIntentsInput{}

	if len(_lexmodelbuildingserviceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lexmodelbuildingserviceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lexmodelbuildingserviceNameContains) > 0 {
		input.NameContains = aws.String(_lexmodelbuildingserviceNameContains)
	}
	if len(_lexmodelbuildingserviceNextToken) > 0 {
		input.NextToken = aws.String(_lexmodelbuildingserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetIntents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lexmodelbuildingservice.GetIntentsOutput
	p := lexmodelbuildingservice.NewGetIntentsPaginator(client, input)
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

// Provides details about an ongoing or complete migration from an Amazon Lex V1
// bot to an Amazon Lex V2 bot. Use this operation to view the migration alerts and
// warnings related to the migration.
func lexmodelbuildingservice_GetMigration(cfg aws.Config, client *lexmodelbuildingservice.Client) {
	input := &lexmodelbuildingservice.GetMigrationInput{
		// MigrationId: *string, // Required
	}

	if len(_lexmodelbuildingserviceMigrationId) > 0 {
		input.MigrationId = aws.String(_lexmodelbuildingserviceMigrationId)
	}

	if resp, err := client.GetMigration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a list of migrations between Amazon Lex V1 and Amazon Lex V2.
func lexmodelbuildingservice_GetMigrations(cfg aws.Config, client *lexmodelbuildingservice.Client) {
	input := &lexmodelbuildingservice.GetMigrationsInput{}

	if len(_lexmodelbuildingserviceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lexmodelbuildingserviceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lexmodelbuildingserviceMigrationStatusEquals) > 0 {
		if err := assignInputField(input, "MigrationStatusEquals", _lexmodelbuildingserviceMigrationStatusEquals); err != nil {
			log.Errorf("invalid --migration-status-equals: %s", err.Error())
			return
		}
	}
	if len(_lexmodelbuildingserviceNextToken) > 0 {
		input.NextToken = aws.String(_lexmodelbuildingserviceNextToken)
	}
	if len(_lexmodelbuildingserviceSortByAttribute) > 0 {
		if err := assignInputField(input, "SortByAttribute", _lexmodelbuildingserviceSortByAttribute); err != nil {
			log.Errorf("invalid --sort-by-attribute: %s", err.Error())
			return
		}
	}
	if len(_lexmodelbuildingserviceSortByOrder) > 0 {
		if err := assignInputField(input, "SortByOrder", _lexmodelbuildingserviceSortByOrder); err != nil {
			log.Errorf("invalid --sort-by-order: %s", err.Error())
			return
		}
	}
	if len(_lexmodelbuildingserviceV1BotNameContains) > 0 {
		input.V1BotNameContains = aws.String(_lexmodelbuildingserviceV1BotNameContains)
	}

	if disablePaginator() {
		if resp, err := client.GetMigrations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lexmodelbuildingservice.GetMigrationsOutput
	p := lexmodelbuildingservice.NewGetMigrationsPaginator(client, input)
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

// Returns information about a specific version of a slot type. In addition to
// specifying the slot type name, you must specify the slot type version.
//
// This operation requires permissions for the lex:GetSlotType action.
func lexmodelbuildingservice_GetSlotType(cfg aws.Config, client *lexmodelbuildingservice.Client) {
	input := &lexmodelbuildingservice.GetSlotTypeInput{
		// Name: *string, // Required
		// Version: *string, // Required
	}

	if len(_lexmodelbuildingserviceName) > 0 {
		input.Name = aws.String(_lexmodelbuildingserviceName)
	}
	if len(_lexmodelbuildingserviceVersion) > 0 {
		input.Version = aws.String(_lexmodelbuildingserviceVersion)
	}

	if resp, err := client.GetSlotType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about all versions of a slot type.
// The GetSlotTypeVersions operation returns a SlotTypeMetadata object for each
// version of a slot type. For example, if a slot type has three numbered versions,
// the GetSlotTypeVersions operation returns four SlotTypeMetadata objects in the
// response, one for each numbered version and one for the $LATEST version.
//
// The GetSlotTypeVersions operation always returns at least one version, the
// $LATEST version.
//
// This operation requires permissions for the lex:GetSlotTypeVersions action.
func lexmodelbuildingservice_GetSlotTypeVersions(cfg aws.Config, client *lexmodelbuildingservice.Client) {
	input := &lexmodelbuildingservice.GetSlotTypeVersionsInput{
		// Name: *string, // Required
	}

	if len(_lexmodelbuildingserviceName) > 0 {
		input.Name = aws.String(_lexmodelbuildingserviceName)
	}
	if len(_lexmodelbuildingserviceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lexmodelbuildingserviceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lexmodelbuildingserviceNextToken) > 0 {
		input.NextToken = aws.String(_lexmodelbuildingserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetSlotTypeVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lexmodelbuildingservice.GetSlotTypeVersionsOutput
	p := lexmodelbuildingservice.NewGetSlotTypeVersionsPaginator(client, input)
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

// Returns slot type information as follows:
// - If you specify the nameContains field, returns the $LATEST version of all
// slot types that contain the specified string.
//
// - If you don't specify the nameContains field, returns information about the
// $LATEST version of all slot types.
//
// The operation requires permission for the lex:GetSlotTypes action.
func lexmodelbuildingservice_GetSlotTypes(cfg aws.Config, client *lexmodelbuildingservice.Client) {
	input := &lexmodelbuildingservice.GetSlotTypesInput{}

	if len(_lexmodelbuildingserviceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lexmodelbuildingserviceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lexmodelbuildingserviceNameContains) > 0 {
		input.NameContains = aws.String(_lexmodelbuildingserviceNameContains)
	}
	if len(_lexmodelbuildingserviceNextToken) > 0 {
		input.NextToken = aws.String(_lexmodelbuildingserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetSlotTypes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lexmodelbuildingservice.GetSlotTypesOutput
	p := lexmodelbuildingservice.NewGetSlotTypesPaginator(client, input)
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

// Use the GetUtterancesView operation to get information about the utterances
// that your users have made to your bot. You can use this list to tune the
// utterances that your bot responds to.
//
// For example, say that you have created a bot to order flowers. After your users
// have used your bot for a while, use the GetUtterancesView operation to see the
// requests that they have made and whether they have been successful. You might
// find that the utterance "I want flowers" is not being recognized. You could add
// this utterance to the OrderFlowers intent so that your bot recognizes that
// utterance.
//
// After you publish a new version of a bot, you can get information about the old
// version and the new so that you can compare the performance across the two
// versions.
//
// Utterance statistics are generated once a day. Data is available for the last
// 15 days. You can request information for up to 5 versions of your bot in each
// request. Amazon Lex returns the most frequent utterances received by the bot in
// the last 15 days. The response contains information about a maximum of 100
// utterances for each version.
//
// If you set childDirected field to true when you created your bot, if you are
// using slot obfuscation with one or more slots, or if you opted out of
// participating in improving Amazon Lex, utterances are not available.
//
// This operation requires permissions for the lex:GetUtterancesView action.
func lexmodelbuildingservice_GetUtterancesView(cfg aws.Config, client *lexmodelbuildingservice.Client) {
	input := &lexmodelbuildingservice.GetUtterancesViewInput{
		// BotName: *string, // Required
		// BotVersions: []string, // Required
		// StatusType: types.StatusType, // Required
	}

	if len(_lexmodelbuildingserviceBotName) > 0 {
		input.BotName = aws.String(_lexmodelbuildingserviceBotName)
	}
	if len(_lexmodelbuildingserviceBotVersions) > 0 {
		input.BotVersions = append([]string(nil), _lexmodelbuildingserviceBotVersions...)
	}
	if len(_lexmodelbuildingserviceStatusType) > 0 {
		if err := assignInputField(input, "StatusType", _lexmodelbuildingserviceStatusType); err != nil {
			log.Errorf("invalid --status-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetUtterancesView(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a list of tags associated with the specified resource. Only bots, bot
// aliases, and bot channels can have tags associated with them.
func lexmodelbuildingservice_ListTagsForResource(cfg aws.Config, client *lexmodelbuildingservice.Client) {
	input := &lexmodelbuildingservice.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_lexmodelbuildingserviceResourceArn) > 0 {
		input.ResourceArn = aws.String(_lexmodelbuildingserviceResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon Lex conversational bot or replaces an existing bot. When you
// create or update a bot you are only required to specify a name, a locale, and
// whether the bot is directed toward children under age 13. You can use this to
// add intents later, or to remove intents from an existing bot. When you create a
// bot with the minimum information, the bot is created or updated but Amazon Lex
// returns the response FAILED . You can build the bot after you add one or more
// intents. For more information about Amazon Lex bots, see how-it-works.
//
// If you specify the name of an existing bot, the fields in the request replace
// the existing values in the $LATEST version of the bot. Amazon Lex removes any
// fields that you don't provide values for in the request, except for the
// idleTTLInSeconds and privacySettings fields, which are set to their default
// values. If you don't specify values for required fields, Amazon Lex throws an
// exception.
//
// This operation requires permissions for the lex:PutBot action. For more
// information, see security-iam.
func lexmodelbuildingservice_PutBot(cfg aws.Config, client *lexmodelbuildingservice.Client) {
	input := &lexmodelbuildingservice.PutBotInput{
		// ChildDirected: *bool, // Required
		// Locale: types.Locale, // Required
		// Name: *string, // Required
	}

	if len(_lexmodelbuildingserviceChildDirected) > 0 {
		if err := assignInputField(input, "ChildDirected", _lexmodelbuildingserviceChildDirected); err != nil {
			log.Errorf("invalid --child-directed: %s", err.Error())
			return
		}
	}
	if len(_lexmodelbuildingserviceLocale) > 0 {
		if err := assignInputField(input, "Locale", _lexmodelbuildingserviceLocale); err != nil {
			log.Errorf("invalid --locale: %s", err.Error())
			return
		}
	}
	if len(_lexmodelbuildingserviceName) > 0 {
		input.Name = aws.String(_lexmodelbuildingserviceName)
	}
	if len(_lexmodelbuildingserviceAbortStatement) > 0 {
		if err := assignInputField(input, "AbortStatement", _lexmodelbuildingserviceAbortStatement); err != nil {
			log.Errorf("invalid --abort-statement: %s", err.Error())
			return
		}
	}
	if len(_lexmodelbuildingserviceChecksum) > 0 {
		input.Checksum = aws.String(_lexmodelbuildingserviceChecksum)
	}
	if len(_lexmodelbuildingserviceClarificationPrompt) > 0 {
		if err := assignInputField(input, "ClarificationPrompt", _lexmodelbuildingserviceClarificationPrompt); err != nil {
			log.Errorf("invalid --clarification-prompt: %s", err.Error())
			return
		}
	}
	if len(_lexmodelbuildingserviceCreateVersion) > 0 {
		if err := assignInputField(input, "CreateVersion", _lexmodelbuildingserviceCreateVersion); err != nil {
			log.Errorf("invalid --create-version: %s", err.Error())
			return
		}
	}
	if len(_lexmodelbuildingserviceDescription) > 0 {
		input.Description = aws.String(_lexmodelbuildingserviceDescription)
	}
	if len(_lexmodelbuildingserviceDetectSentiment) > 0 {
		if err := assignInputField(input, "DetectSentiment", _lexmodelbuildingserviceDetectSentiment); err != nil {
			log.Errorf("invalid --detect-sentiment: %s", err.Error())
			return
		}
	}
	if len(_lexmodelbuildingserviceEnableModelImprovements) > 0 {
		if err := assignInputField(input, "EnableModelImprovements", _lexmodelbuildingserviceEnableModelImprovements); err != nil {
			log.Errorf("invalid --enable-model-improvements: %s", err.Error())
			return
		}
	}
	if len(_lexmodelbuildingserviceIdleSessionTTLInSeconds) > 0 {
		if err := assignInputField(input, "IdleSessionTTLInSeconds", _lexmodelbuildingserviceIdleSessionTTLInSeconds); err != nil {
			log.Errorf("invalid --idle-session-ttlin-seconds: %s", err.Error())
			return
		}
	}
	if len(_lexmodelbuildingserviceIntents) > 0 {
		if err := assignInputField(input, "Intents", _lexmodelbuildingserviceIntents); err != nil {
			log.Errorf("invalid --intents: %s", err.Error())
			return
		}
	}
	if len(_lexmodelbuildingserviceNluIntentConfidenceThreshold) > 0 {
		if err := assignInputField(input, "NluIntentConfidenceThreshold", _lexmodelbuildingserviceNluIntentConfidenceThreshold); err != nil {
			log.Errorf("invalid --nlu-intent-confidence-threshold: %s", err.Error())
			return
		}
	}
	if len(_lexmodelbuildingserviceProcessBehavior) > 0 {
		if err := assignInputField(input, "ProcessBehavior", _lexmodelbuildingserviceProcessBehavior); err != nil {
			log.Errorf("invalid --process-behavior: %s", err.Error())
			return
		}
	}
	if len(_lexmodelbuildingserviceTags) > 0 {
		if err := assignInputField(input, "Tags", _lexmodelbuildingserviceTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_lexmodelbuildingserviceVoiceId) > 0 {
		input.VoiceId = aws.String(_lexmodelbuildingserviceVoiceId)
	}

	if resp, err := client.PutBot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an alias for the specified version of the bot or replaces an alias for
// the specified bot. To change the version of the bot that the alias points to,
// replace the alias. For more information about aliases, see versioning-aliases.
//
// This operation requires permissions for the lex:PutBotAlias action.
func lexmodelbuildingservice_PutBotAlias(cfg aws.Config, client *lexmodelbuildingservice.Client) {
	input := &lexmodelbuildingservice.PutBotAliasInput{
		// BotName: *string, // Required
		// BotVersion: *string, // Required
		// Name: *string, // Required
	}

	if len(_lexmodelbuildingserviceBotName) > 0 {
		input.BotName = aws.String(_lexmodelbuildingserviceBotName)
	}
	if len(_lexmodelbuildingserviceBotVersion) > 0 {
		input.BotVersion = aws.String(_lexmodelbuildingserviceBotVersion)
	}
	if len(_lexmodelbuildingserviceName) > 0 {
		input.Name = aws.String(_lexmodelbuildingserviceName)
	}
	if len(_lexmodelbuildingserviceChecksum) > 0 {
		input.Checksum = aws.String(_lexmodelbuildingserviceChecksum)
	}
	if len(_lexmodelbuildingserviceConversationLogs) > 0 {
		if err := assignInputField(input, "ConversationLogs", _lexmodelbuildingserviceConversationLogs); err != nil {
			log.Errorf("invalid --conversation-logs: %s", err.Error())
			return
		}
	}
	if len(_lexmodelbuildingserviceDescription) > 0 {
		input.Description = aws.String(_lexmodelbuildingserviceDescription)
	}
	if len(_lexmodelbuildingserviceTags) > 0 {
		if err := assignInputField(input, "Tags", _lexmodelbuildingserviceTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutBotAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an intent or replaces an existing intent.
// To define the interaction between the user and your bot, you use one or more
// intents. For a pizza ordering bot, for example, you would create an OrderPizza
// intent.
//
// To create an intent or replace an existing intent, you must provide the
// following:
//
// - Intent name. For example, OrderPizza .
//
// - Sample utterances. For example, "Can I order a pizza, please." and "I want
// to order a pizza."
//
// - Information to be gathered. You specify slot types for the information that
// your bot will request from the user. You can specify standard slot types, such
// as a date or a time, or custom slot types such as the size and crust of a pizza.
//
// - How the intent will be fulfilled. You can provide a Lambda function or
// configure the intent to return the intent information to the client application.
// If you use a Lambda function, when all of the intent information is available,
// Amazon Lex invokes your Lambda function. If you configure your intent to return
// the intent information to the client application.
//
// You can specify other optional information in the request, such as:
//
// - A confirmation prompt to ask the user to confirm an intent. For example,
// "Shall I order your pizza?"
//
// - A conclusion statement to send to the user after the intent has been
// fulfilled. For example, "I placed your pizza order."
//
// - A follow-up prompt that asks the user for additional activity. For example,
// asking "Do you want to order a drink with your pizza?"
//
// If you specify an existing intent name to update the intent, Amazon Lex
// replaces the values in the $LATEST version of the intent with the values in the
// request. Amazon Lex removes fields that you don't provide in the request. If you
// don't specify the required fields, Amazon Lex throws an exception. When you
// update the $LATEST version of an intent, the status field of any bot that uses
// the $LATEST version of the intent is set to NOT_BUILT .
//
// For more information, see how-it-works.
//
// This operation requires permissions for the lex:PutIntent action.
func lexmodelbuildingservice_PutIntent(cfg aws.Config, client *lexmodelbuildingservice.Client) {
	input := &lexmodelbuildingservice.PutIntentInput{
		// Name: *string, // Required
	}

	if len(_lexmodelbuildingserviceName) > 0 {
		input.Name = aws.String(_lexmodelbuildingserviceName)
	}
	if len(_lexmodelbuildingserviceChecksum) > 0 {
		input.Checksum = aws.String(_lexmodelbuildingserviceChecksum)
	}
	if len(_lexmodelbuildingserviceConclusionStatement) > 0 {
		if err := assignInputField(input, "ConclusionStatement", _lexmodelbuildingserviceConclusionStatement); err != nil {
			log.Errorf("invalid --conclusion-statement: %s", err.Error())
			return
		}
	}
	if len(_lexmodelbuildingserviceConfirmationPrompt) > 0 {
		if err := assignInputField(input, "ConfirmationPrompt", _lexmodelbuildingserviceConfirmationPrompt); err != nil {
			log.Errorf("invalid --confirmation-prompt: %s", err.Error())
			return
		}
	}
	if len(_lexmodelbuildingserviceCreateVersion) > 0 {
		if err := assignInputField(input, "CreateVersion", _lexmodelbuildingserviceCreateVersion); err != nil {
			log.Errorf("invalid --create-version: %s", err.Error())
			return
		}
	}
	if len(_lexmodelbuildingserviceDescription) > 0 {
		input.Description = aws.String(_lexmodelbuildingserviceDescription)
	}
	if len(_lexmodelbuildingserviceDialogCodeHook) > 0 {
		if err := assignInputField(input, "DialogCodeHook", _lexmodelbuildingserviceDialogCodeHook); err != nil {
			log.Errorf("invalid --dialog-code-hook: %s", err.Error())
			return
		}
	}
	if len(_lexmodelbuildingserviceFollowUpPrompt) > 0 {
		if err := assignInputField(input, "FollowUpPrompt", _lexmodelbuildingserviceFollowUpPrompt); err != nil {
			log.Errorf("invalid --follow-up-prompt: %s", err.Error())
			return
		}
	}
	if len(_lexmodelbuildingserviceFulfillmentActivity) > 0 {
		if err := assignInputField(input, "FulfillmentActivity", _lexmodelbuildingserviceFulfillmentActivity); err != nil {
			log.Errorf("invalid --fulfillment-activity: %s", err.Error())
			return
		}
	}
	if len(_lexmodelbuildingserviceInputContexts) > 0 {
		if err := assignInputField(input, "InputContexts", _lexmodelbuildingserviceInputContexts); err != nil {
			log.Errorf("invalid --input-contexts: %s", err.Error())
			return
		}
	}
	if len(_lexmodelbuildingserviceKendraConfiguration) > 0 {
		if err := assignInputField(input, "KendraConfiguration", _lexmodelbuildingserviceKendraConfiguration); err != nil {
			log.Errorf("invalid --kendra-configuration: %s", err.Error())
			return
		}
	}
	if len(_lexmodelbuildingserviceOutputContexts) > 0 {
		if err := assignInputField(input, "OutputContexts", _lexmodelbuildingserviceOutputContexts); err != nil {
			log.Errorf("invalid --output-contexts: %s", err.Error())
			return
		}
	}
	if len(_lexmodelbuildingserviceParentIntentSignature) > 0 {
		input.ParentIntentSignature = aws.String(_lexmodelbuildingserviceParentIntentSignature)
	}
	if len(_lexmodelbuildingserviceRejectionStatement) > 0 {
		if err := assignInputField(input, "RejectionStatement", _lexmodelbuildingserviceRejectionStatement); err != nil {
			log.Errorf("invalid --rejection-statement: %s", err.Error())
			return
		}
	}
	if len(_lexmodelbuildingserviceSampleUtterances) > 0 {
		input.SampleUtterances = append([]string(nil), _lexmodelbuildingserviceSampleUtterances...)
	}
	if len(_lexmodelbuildingserviceSlots) > 0 {
		if err := assignInputField(input, "Slots", _lexmodelbuildingserviceSlots); err != nil {
			log.Errorf("invalid --slots: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutIntent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a custom slot type or replaces an existing custom slot type.
// To create a custom slot type, specify a name for the slot type and a set of
// enumeration values, which are the values that a slot of this type can assume.
// For more information, see how-it-works.
//
// If you specify the name of an existing slot type, the fields in the request
// replace the existing values in the $LATEST version of the slot type. Amazon Lex
// removes the fields that you don't provide in the request. If you don't specify
// required fields, Amazon Lex throws an exception. When you update the $LATEST
// version of a slot type, if a bot uses the $LATEST version of an intent that
// contains the slot type, the bot's status field is set to NOT_BUILT .
//
// This operation requires permissions for the lex:PutSlotType action.
func lexmodelbuildingservice_PutSlotType(cfg aws.Config, client *lexmodelbuildingservice.Client) {
	input := &lexmodelbuildingservice.PutSlotTypeInput{
		// Name: *string, // Required
	}

	if len(_lexmodelbuildingserviceName) > 0 {
		input.Name = aws.String(_lexmodelbuildingserviceName)
	}
	if len(_lexmodelbuildingserviceChecksum) > 0 {
		input.Checksum = aws.String(_lexmodelbuildingserviceChecksum)
	}
	if len(_lexmodelbuildingserviceCreateVersion) > 0 {
		if err := assignInputField(input, "CreateVersion", _lexmodelbuildingserviceCreateVersion); err != nil {
			log.Errorf("invalid --create-version: %s", err.Error())
			return
		}
	}
	if len(_lexmodelbuildingserviceDescription) > 0 {
		input.Description = aws.String(_lexmodelbuildingserviceDescription)
	}
	if len(_lexmodelbuildingserviceEnumerationValues) > 0 {
		if err := assignInputField(input, "EnumerationValues", _lexmodelbuildingserviceEnumerationValues); err != nil {
			log.Errorf("invalid --enumeration-values: %s", err.Error())
			return
		}
	}
	if len(_lexmodelbuildingserviceParentSlotTypeSignature) > 0 {
		input.ParentSlotTypeSignature = aws.String(_lexmodelbuildingserviceParentSlotTypeSignature)
	}
	if len(_lexmodelbuildingserviceSlotTypeConfigurations) > 0 {
		if err := assignInputField(input, "SlotTypeConfigurations", _lexmodelbuildingserviceSlotTypeConfigurations); err != nil {
			log.Errorf("invalid --slot-type-configurations: %s", err.Error())
			return
		}
	}
	if len(_lexmodelbuildingserviceValueSelectionStrategy) > 0 {
		if err := assignInputField(input, "ValueSelectionStrategy", _lexmodelbuildingserviceValueSelectionStrategy); err != nil {
			log.Errorf("invalid --value-selection-strategy: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutSlotType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a job to import a resource to Amazon Lex.
func lexmodelbuildingservice_StartImport(cfg aws.Config, client *lexmodelbuildingservice.Client) {
	input := &lexmodelbuildingservice.StartImportInput{
		// MergeStrategy: types.MergeStrategy, // Required
		// Payload: []byte, // Required
		// ResourceType: types.ResourceType, // Required
	}

	if len(_lexmodelbuildingserviceMergeStrategy) > 0 {
		if err := assignInputField(input, "MergeStrategy", _lexmodelbuildingserviceMergeStrategy); err != nil {
			log.Errorf("invalid --merge-strategy: %s", err.Error())
			return
		}
	}
	if len(_lexmodelbuildingservicePayload) > 0 {
		if err := assignInputField(input, "Payload", _lexmodelbuildingservicePayload); err != nil {
			log.Errorf("invalid --payload: %s", err.Error())
			return
		}
	}
	if len(_lexmodelbuildingserviceResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _lexmodelbuildingserviceResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}
	if len(_lexmodelbuildingserviceTags) > 0 {
		if err := assignInputField(input, "Tags", _lexmodelbuildingserviceTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartImport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts migrating a bot from Amazon Lex V1 to Amazon Lex V2. Migrate your bot
// when you want to take advantage of the new features of Amazon Lex V2.
//
// For more information, see [Migrating a bot] in the Amazon Lex developer guide.
//
// [Migrating a bot]: https://docs.aws.amazon.com/lex/latest/dg/migrate.html
func lexmodelbuildingservice_StartMigration(cfg aws.Config, client *lexmodelbuildingservice.Client) {
	input := &lexmodelbuildingservice.StartMigrationInput{
		// MigrationStrategy: types.MigrationStrategy, // Required
		// V1BotName: *string, // Required
		// V1BotVersion: *string, // Required
		// V2BotName: *string, // Required
		// V2BotRole: *string, // Required
	}

	if len(_lexmodelbuildingserviceMigrationStrategy) > 0 {
		if err := assignInputField(input, "MigrationStrategy", _lexmodelbuildingserviceMigrationStrategy); err != nil {
			log.Errorf("invalid --migration-strategy: %s", err.Error())
			return
		}
	}
	if len(_lexmodelbuildingserviceV1BotName) > 0 {
		input.V1BotName = aws.String(_lexmodelbuildingserviceV1BotName)
	}
	if len(_lexmodelbuildingserviceV1BotVersion) > 0 {
		input.V1BotVersion = aws.String(_lexmodelbuildingserviceV1BotVersion)
	}
	if len(_lexmodelbuildingserviceV2BotName) > 0 {
		input.V2BotName = aws.String(_lexmodelbuildingserviceV2BotName)
	}
	if len(_lexmodelbuildingserviceV2BotRole) > 0 {
		input.V2BotRole = aws.String(_lexmodelbuildingserviceV2BotRole)
	}

	if resp, err := client.StartMigration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds the specified tags to the specified resource. If a tag key already exists,
// the existing value is replaced with the new value.
func lexmodelbuildingservice_TagResource(cfg aws.Config, client *lexmodelbuildingservice.Client) {
	input := &lexmodelbuildingservice.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_lexmodelbuildingserviceResourceArn) > 0 {
		input.ResourceArn = aws.String(_lexmodelbuildingserviceResourceArn)
	}
	if len(_lexmodelbuildingserviceTags) > 0 {
		if err := assignInputField(input, "Tags", _lexmodelbuildingserviceTags); err != nil {
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

// Removes tags from a bot, bot alias or bot channel.
func lexmodelbuildingservice_UntagResource(cfg aws.Config, client *lexmodelbuildingservice.Client) {
	input := &lexmodelbuildingservice.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_lexmodelbuildingserviceResourceArn) > 0 {
		input.ResourceArn = aws.String(_lexmodelbuildingserviceResourceArn)
	}
	if len(_lexmodelbuildingserviceTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _lexmodelbuildingserviceTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_lexmodelbuildingserviceCmd)
	_lexmodelbuildingserviceCmd.Flags().SortFlags = false

	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceAbortStatement, "abort-statement", "", "", "Abort Statement")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceBotAlias, "bot-alias", "", "", "Bot Alias")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceBotName, "bot-name", "", "", "Bot Name")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceBotVersion, "bot-version", "", "", "Bot Version")
	_lexmodelbuildingserviceCmd.Flags().StringSliceVarP(&_lexmodelbuildingserviceBotVersions, "bot-versions", "", nil, "Bot Versions")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceChecksum, "checksum", "", "", "Checksum")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceChildDirected, "child-directed", "", "", "Child Directed")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceClarificationPrompt, "clarification-prompt", "", "", "Clarification Prompt")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceConclusionStatement, "conclusion-statement", "", "", "Conclusion Statement")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceConfirmationPrompt, "confirmation-prompt", "", "", "Confirmation Prompt")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceConversationLogs, "conversation-logs", "", "", "Conversation Logs")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceCreateVersion, "create-version", "", "", "Create Version")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceDescription, "description", "", "", "Description")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceDetectSentiment, "detect-sentiment", "", "", "Detect Sentiment")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceDialogCodeHook, "dialog-code-hook", "", "", "Dialog Code Hook")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceEnableModelImprovements, "enable-model-improvements", "", "", "Enable Model Improvements")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceEnumerationValues, "enumeration-values", "", "", "Enumeration Values")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceExportType, "export-type", "", "", "Export Type")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceFollowUpPrompt, "follow-up-prompt", "", "", "Follow Up Prompt")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceFulfillmentActivity, "fulfillment-activity", "", "", "Fulfillment Activity")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceIdleSessionTTLInSeconds, "idle-session-ttlin-seconds", "", "", "Idle Session Ttlin Seconds")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceImportId, "import-id", "", "", "Import ID")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceInputContexts, "input-contexts", "", "", "Input Contexts")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceIntents, "intents", "", "", "Intents")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceKendraConfiguration, "kendra-configuration", "", "", "Kendra Configuration")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceLocale, "locale", "", "", "Locale")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceMaxResults, "max-results", "", "", "Max Results")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceMergeStrategy, "merge-strategy", "", "", "Merge Strategy")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceMigrationId, "migration-id", "", "", "Migration ID")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceMigrationStatusEquals, "migration-status-equals", "", "", "Migration Status Equals")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceMigrationStrategy, "migration-strategy", "", "", "Migration Strategy")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceName, "name", "", "", "Name")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceNameContains, "name-contains", "", "", "Name Contains")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceNextToken, "next-token", "", "", "Next Token")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceNluIntentConfidenceThreshold, "nlu-intent-confidence-threshold", "", "", "Nlu Intent Confidence Threshold")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceOutputContexts, "output-contexts", "", "", "Output Contexts")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceParentIntentSignature, "parent-intent-signature", "", "", "Parent Intent Signature")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceParentSlotTypeSignature, "parent-slot-type-signature", "", "", "Parent Slot Type Signature")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingservicePayload, "payload", "", "", "Payload")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceProcessBehavior, "process-behavior", "", "", "Process Behavior")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceRejectionStatement, "rejection-statement", "", "", "Rejection Statement")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceResourceArn, "resource-arn", "", "", "Resource ARN")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceResourceType, "resource-type", "", "", "Resource Type")
	_lexmodelbuildingserviceCmd.Flags().StringSliceVarP(&_lexmodelbuildingserviceSampleUtterances, "sample-utterances", "", nil, "Sample Utterances")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceSignature, "signature", "", "", "Signature")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceSignatureContains, "signature-contains", "", "", "Signature Contains")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceSlotTypeConfigurations, "slot-type-configurations", "", "", "Slot Type Configurations")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceSlots, "slots", "", "", "Slots")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceSortByAttribute, "sort-by-attribute", "", "", "Sort By Attribute")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceSortByOrder, "sort-by-order", "", "", "Sort By Order")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceStatusType, "status-type", "", "", "Status Type")
	_lexmodelbuildingserviceCmd.Flags().StringSliceVarP(&_lexmodelbuildingserviceTagKeys, "tag-keys", "", nil, "Tag Keys")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceTags, "tags", "", "", "Tags")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceUserId, "user-id", "", "", "User ID")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceV1BotName, "v1-bot-name", "", "", "V1 Bot Name")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceV1BotNameContains, "v1-bot-name-contains", "", "", "V1 Bot Name Contains")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceV1BotVersion, "v1-bot-version", "", "", "V1 Bot Version")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceV2BotName, "v2-bot-name", "", "", "V2 Bot Name")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceV2BotRole, "v2-bot-role", "", "", "V2 Bot Role")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceValueSelectionStrategy, "value-selection-strategy", "", "", "Value Selection Strategy")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceVersion, "version", "", "", "Version")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceVersionOrAlias, "version-or-alias", "", "", "Version Or Alias")
	_lexmodelbuildingserviceCmd.Flags().StringVarP(&_lexmodelbuildingserviceVoiceId, "voice-id", "", "", "Voice ID")

	_lexmodelbuildingserviceCmd.Flags().BoolVarP(&_lexmodelbuildingserviceCreateBotVersion, "create-bot-version", "", false, "Create Bot Version")
	_lexmodelbuildingserviceCmd.Flags().BoolVarP(&_lexmodelbuildingserviceCreateIntentVersion, "create-intent-version", "", false, "Create Intent Version")
	_lexmodelbuildingserviceCmd.Flags().BoolVarP(&_lexmodelbuildingserviceCreateSlotTypeVersion, "create-slot-type-version", "", false, "Create Slot Type Version")
	_lexmodelbuildingserviceCmd.Flags().BoolVarP(&_lexmodelbuildingserviceDeleteBot, "delete-bot", "", false, "Delete Bot")
	_lexmodelbuildingserviceCmd.Flags().BoolVarP(&_lexmodelbuildingserviceDeleteBotAlias, "delete-bot-alias", "", false, "Delete Bot Alias")
	_lexmodelbuildingserviceCmd.Flags().BoolVarP(&_lexmodelbuildingserviceDeleteBotChannelAssociation, "delete-bot-channel-association", "", false, "Delete Bot Channel Association")
	_lexmodelbuildingserviceCmd.Flags().BoolVarP(&_lexmodelbuildingserviceDeleteBotVersion, "delete-bot-version", "", false, "Delete Bot Version")
	_lexmodelbuildingserviceCmd.Flags().BoolVarP(&_lexmodelbuildingserviceDeleteIntent, "delete-intent", "", false, "Delete Intent")
	_lexmodelbuildingserviceCmd.Flags().BoolVarP(&_lexmodelbuildingserviceDeleteIntentVersion, "delete-intent-version", "", false, "Delete Intent Version")
	_lexmodelbuildingserviceCmd.Flags().BoolVarP(&_lexmodelbuildingserviceDeleteSlotType, "delete-slot-type", "", false, "Delete Slot Type")
	_lexmodelbuildingserviceCmd.Flags().BoolVarP(&_lexmodelbuildingserviceDeleteSlotTypeVersion, "delete-slot-type-version", "", false, "Delete Slot Type Version")
	_lexmodelbuildingserviceCmd.Flags().BoolVarP(&_lexmodelbuildingserviceDeleteUtterances, "delete-utterances", "", false, "Delete Utterances")
	_lexmodelbuildingserviceCmd.Flags().BoolVarP(&_lexmodelbuildingserviceGetBot, "get-bot", "", false, "Get Bot")
	_lexmodelbuildingserviceCmd.Flags().BoolVarP(&_lexmodelbuildingserviceGetBotAlias, "get-bot-alias", "", false, "Get Bot Alias")
	_lexmodelbuildingserviceCmd.Flags().BoolVarP(&_lexmodelbuildingserviceGetBotAliases, "get-bot-aliases", "", false, "Get Bot Aliases")
	_lexmodelbuildingserviceCmd.Flags().BoolVarP(&_lexmodelbuildingserviceGetBotChannelAssociation, "get-bot-channel-association", "", false, "Get Bot Channel Association")
	_lexmodelbuildingserviceCmd.Flags().BoolVarP(&_lexmodelbuildingserviceGetBotChannelAssociations, "get-bot-channel-associations", "", false, "Get Bot Channel Associations")
	_lexmodelbuildingserviceCmd.Flags().BoolVarP(&_lexmodelbuildingserviceGetBotVersions, "get-bot-versions", "", false, "Get Bot Versions")
	_lexmodelbuildingserviceCmd.Flags().BoolVarP(&_lexmodelbuildingserviceGetBots, "get-bots", "", false, "Get Bots")
	_lexmodelbuildingserviceCmd.Flags().BoolVarP(&_lexmodelbuildingserviceGetBuiltinIntent, "get-builtin-intent", "", false, "Get Builtin Intent")
	_lexmodelbuildingserviceCmd.Flags().BoolVarP(&_lexmodelbuildingserviceGetBuiltinIntents, "get-builtin-intents", "", false, "Get Builtin Intents")
	_lexmodelbuildingserviceCmd.Flags().BoolVarP(&_lexmodelbuildingserviceGetBuiltinSlotTypes, "get-builtin-slot-types", "", false, "Get Builtin Slot Types")
	_lexmodelbuildingserviceCmd.Flags().BoolVarP(&_lexmodelbuildingserviceGetExport, "get-export", "", false, "Get Export")
	_lexmodelbuildingserviceCmd.Flags().BoolVarP(&_lexmodelbuildingserviceGetImport, "get-import", "", false, "Get Import")
	_lexmodelbuildingserviceCmd.Flags().BoolVarP(&_lexmodelbuildingserviceGetIntent, "get-intent", "", false, "Get Intent")
	_lexmodelbuildingserviceCmd.Flags().BoolVarP(&_lexmodelbuildingserviceGetIntentVersions, "get-intent-versions", "", false, "Get Intent Versions")
	_lexmodelbuildingserviceCmd.Flags().BoolVarP(&_lexmodelbuildingserviceGetIntents, "get-intents", "", false, "Get Intents")
	_lexmodelbuildingserviceCmd.Flags().BoolVarP(&_lexmodelbuildingserviceGetMigration, "get-migration", "", false, "Get Migration")
	_lexmodelbuildingserviceCmd.Flags().BoolVarP(&_lexmodelbuildingserviceGetMigrations, "get-migrations", "", false, "Get Migrations")
	_lexmodelbuildingserviceCmd.Flags().BoolVarP(&_lexmodelbuildingserviceGetSlotType, "get-slot-type", "", false, "Get Slot Type")
	_lexmodelbuildingserviceCmd.Flags().BoolVarP(&_lexmodelbuildingserviceGetSlotTypeVersions, "get-slot-type-versions", "", false, "Get Slot Type Versions")
	_lexmodelbuildingserviceCmd.Flags().BoolVarP(&_lexmodelbuildingserviceGetSlotTypes, "get-slot-types", "", false, "Get Slot Types")
	_lexmodelbuildingserviceCmd.Flags().BoolVarP(&_lexmodelbuildingserviceGetUtterancesView, "get-utterances-view", "", false, "Get Utterances View")
	_lexmodelbuildingserviceCmd.Flags().BoolVarP(&_lexmodelbuildingserviceListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_lexmodelbuildingserviceCmd.Flags().BoolVarP(&_lexmodelbuildingservicePutBot, "put-bot", "", false, "Put Bot")
	_lexmodelbuildingserviceCmd.Flags().BoolVarP(&_lexmodelbuildingservicePutBotAlias, "put-bot-alias", "", false, "Put Bot Alias")
	_lexmodelbuildingserviceCmd.Flags().BoolVarP(&_lexmodelbuildingservicePutIntent, "put-intent", "", false, "Put Intent")
	_lexmodelbuildingserviceCmd.Flags().BoolVarP(&_lexmodelbuildingservicePutSlotType, "put-slot-type", "", false, "Put Slot Type")
	_lexmodelbuildingserviceCmd.Flags().BoolVarP(&_lexmodelbuildingserviceStartImport, "start-import", "", false, "Start Import")
	_lexmodelbuildingserviceCmd.Flags().BoolVarP(&_lexmodelbuildingserviceStartMigration, "start-migration", "", false, "Start Migration")
	_lexmodelbuildingserviceCmd.Flags().BoolVarP(&_lexmodelbuildingserviceTagResource, "tag-resource", "", false, "Tag Resource")
	_lexmodelbuildingserviceCmd.Flags().BoolVarP(&_lexmodelbuildingserviceUntagResource, "untag-resource", "", false, "Untag Resource")

}
