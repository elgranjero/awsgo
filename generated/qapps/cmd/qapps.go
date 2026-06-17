package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/qapps"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// qappsCmd represents the qapps command
var _qappsCmd = &cobra.Command{
	Use:   "qapps",
	Short: "AWS qapps CLI",
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
		client := qapps.NewFromConfig(cfg)
		if _qappsAssociateLibraryItemReview {
			qapps_AssociateLibraryItemReview(cfg, client)
			return
		}
		if _qappsAssociateQAppWithUser {
			qapps_AssociateQAppWithUser(cfg, client)
			return
		}
		if _qappsBatchCreateCategory {
			qapps_BatchCreateCategory(cfg, client)
			return
		}
		if _qappsBatchDeleteCategory {
			qapps_BatchDeleteCategory(cfg, client)
			return
		}
		if _qappsBatchUpdateCategory {
			qapps_BatchUpdateCategory(cfg, client)
			return
		}
		if _qappsCreateLibraryItem {
			qapps_CreateLibraryItem(cfg, client)
			return
		}
		if _qappsCreatePresignedUrl {
			qapps_CreatePresignedUrl(cfg, client)
			return
		}
		if _qappsCreateQApp {
			qapps_CreateQApp(cfg, client)
			return
		}
		if _qappsDeleteLibraryItem {
			qapps_DeleteLibraryItem(cfg, client)
			return
		}
		if _qappsDeleteQApp {
			qapps_DeleteQApp(cfg, client)
			return
		}
		if _qappsDescribeQAppPermissions {
			qapps_DescribeQAppPermissions(cfg, client)
			return
		}
		if _qappsDisassociateLibraryItemReview {
			qapps_DisassociateLibraryItemReview(cfg, client)
			return
		}
		if _qappsDisassociateQAppFromUser {
			qapps_DisassociateQAppFromUser(cfg, client)
			return
		}
		if _qappsExportQAppSessionData {
			qapps_ExportQAppSessionData(cfg, client)
			return
		}
		if _qappsGetLibraryItem {
			qapps_GetLibraryItem(cfg, client)
			return
		}
		if _qappsGetQApp {
			qapps_GetQApp(cfg, client)
			return
		}
		if _qappsGetQAppSession {
			qapps_GetQAppSession(cfg, client)
			return
		}
		if _qappsGetQAppSessionMetadata {
			qapps_GetQAppSessionMetadata(cfg, client)
			return
		}
		if _qappsImportDocument {
			qapps_ImportDocument(cfg, client)
			return
		}
		if _qappsListCategories {
			qapps_ListCategories(cfg, client)
			return
		}
		if _qappsListLibraryItems {
			qapps_ListLibraryItems(cfg, client)
			return
		}
		if _qappsListQAppSessionData {
			qapps_ListQAppSessionData(cfg, client)
			return
		}
		if _qappsListQApps {
			qapps_ListQApps(cfg, client)
			return
		}
		if _qappsListTagsForResource {
			qapps_ListTagsForResource(cfg, client)
			return
		}
		if _qappsPredictQApp {
			qapps_PredictQApp(cfg, client)
			return
		}
		if _qappsStartQAppSession {
			qapps_StartQAppSession(cfg, client)
			return
		}
		if _qappsStopQAppSession {
			qapps_StopQAppSession(cfg, client)
			return
		}
		if _qappsTagResource {
			qapps_TagResource(cfg, client)
			return
		}
		if _qappsUntagResource {
			qapps_UntagResource(cfg, client)
			return
		}
		if _qappsUpdateLibraryItem {
			qapps_UpdateLibraryItem(cfg, client)
			return
		}
		if _qappsUpdateLibraryItemMetadata {
			qapps_UpdateLibraryItemMetadata(cfg, client)
			return
		}
		if _qappsUpdateQApp {
			qapps_UpdateQApp(cfg, client)
			return
		}
		if _qappsUpdateQAppPermissions {
			qapps_UpdateQAppPermissions(cfg, client)
			return
		}
		if _qappsUpdateQAppSession {
			qapps_UpdateQAppSession(cfg, client)
			return
		}
		if _qappsUpdateQAppSessionMetadata {
			qapps_UpdateQAppSessionMetadata(cfg, client)
			return
		}

	},
}

var (
	_qappsAssociateLibraryItemReview    bool
	_qappsAssociateQAppWithUser         bool
	_qappsBatchCreateCategory           bool
	_qappsBatchDeleteCategory           bool
	_qappsBatchUpdateCategory           bool
	_qappsCreateLibraryItem             bool
	_qappsCreatePresignedUrl            bool
	_qappsCreateQApp                    bool
	_qappsDeleteLibraryItem             bool
	_qappsDeleteQApp                    bool
	_qappsDescribeQAppPermissions       bool
	_qappsDisassociateLibraryItemReview bool
	_qappsDisassociateQAppFromUser      bool
	_qappsExportQAppSessionData         bool
	_qappsGetLibraryItem                bool
	_qappsGetQApp                       bool
	_qappsGetQAppSession                bool
	_qappsGetQAppSessionMetadata        bool
	_qappsImportDocument                bool
	_qappsListCategories                bool
	_qappsListLibraryItems              bool
	_qappsListQAppSessionData           bool
	_qappsListQApps                     bool
	_qappsListTagsForResource           bool
	_qappsPredictQApp                   bool
	_qappsStartQAppSession              bool
	_qappsStopQAppSession               bool
	_qappsTagResource                   bool
	_qappsUntagResource                 bool
	_qappsUpdateLibraryItem             bool
	_qappsUpdateLibraryItemMetadata     bool
	_qappsUpdateQApp                    bool
	_qappsUpdateQAppPermissions         bool
	_qappsUpdateQAppSession             bool
	_qappsUpdateQAppSessionMetadata     bool

	_qappsAppDefinition        string
	_qappsAppId                string
	_qappsAppVersion           string
	_qappsCardId               string
	_qappsCategories           []string
	_qappsCategoryId           string
	_qappsDescription          string
	_qappsFileContentsBase64   string
	_qappsFileContentsSha256   string
	_qappsFileName             string
	_qappsGrantPermissions     string
	_qappsInitialValues        string
	_qappsInstanceId           string
	_qappsIsVerified           string
	_qappsLibraryItemId        string
	_qappsLimit                string
	_qappsNextToken            string
	_qappsOptions              string
	_qappsResourceARN          string
	_qappsRevokePermissions    string
	_qappsScope                string
	_qappsSessionId            string
	_qappsSessionName          string
	_qappsSharingConfiguration string
	_qappsStatus               string
	_qappsTagKeys              []string
	_qappsTags                 string
	_qappsTitle                string
	_qappsValues               string
)

// Associates a rating or review for a library item with the user submitting the
// request. This increments the rating count for the specified library item.
func qapps_AssociateLibraryItemReview(cfg aws.Config, client *qapps.Client) {
	input := &qapps.AssociateLibraryItemReviewInput{
		// InstanceId: *string, // Required
		// LibraryItemId: *string, // Required
	}

	if len(_qappsInstanceId) > 0 {
		input.InstanceId = aws.String(_qappsInstanceId)
	}
	if len(_qappsLibraryItemId) > 0 {
		input.LibraryItemId = aws.String(_qappsLibraryItemId)
	}

	if resp, err := client.AssociateLibraryItemReview(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation creates a link between the user's identity calling the operation
// and a specific Q App. This is useful to mark the Q App as a favorite for the
// user if the user doesn't own the Amazon Q App so they can still run it and see
// it in their inventory of Q Apps.
func qapps_AssociateQAppWithUser(cfg aws.Config, client *qapps.Client) {
	input := &qapps.AssociateQAppWithUserInput{
		// AppId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_qappsAppId) > 0 {
		input.AppId = aws.String(_qappsAppId)
	}
	if len(_qappsInstanceId) > 0 {
		input.InstanceId = aws.String(_qappsInstanceId)
	}

	if resp, err := client.AssociateQAppWithUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates Categories for the Amazon Q Business application environment instance.
// Web experience users use Categories to tag and filter library items. For more
// information, see [Custom labels for Amazon Q Apps].
//
// [Custom labels for Amazon Q Apps]: https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/qapps-custom-labels.html
func qapps_BatchCreateCategory(cfg aws.Config, client *qapps.Client) {
	input := &qapps.BatchCreateCategoryInput{
		// Categories: []types.BatchCreateCategoryInputCategory, // Required
		// InstanceId: *string, // Required
	}

	if len(_qappsCategories) > 0 {
		if err := assignInputField(input, "Categories", _qappsCategories[0]); err != nil {
			log.Errorf("invalid --categories: %s", err.Error())
			return
		}
	}
	if len(_qappsInstanceId) > 0 {
		input.InstanceId = aws.String(_qappsInstanceId)
	}

	if resp, err := client.BatchCreateCategory(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes Categories for the Amazon Q Business application environment instance.
// Web experience users use Categories to tag and filter library items. For more
// information, see [Custom labels for Amazon Q Apps].
//
// [Custom labels for Amazon Q Apps]: https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/qapps-custom-labels.html
func qapps_BatchDeleteCategory(cfg aws.Config, client *qapps.Client) {
	input := &qapps.BatchDeleteCategoryInput{
		// Categories: []string, // Required
		// InstanceId: *string, // Required
	}

	if len(_qappsCategories) > 0 {
		input.Categories = append([]string(nil), _qappsCategories...)
	}
	if len(_qappsInstanceId) > 0 {
		input.InstanceId = aws.String(_qappsInstanceId)
	}

	if resp, err := client.BatchDeleteCategory(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates Categories for the Amazon Q Business application environment instance.
// Web experience users use Categories to tag and filter library items. For more
// information, see [Custom labels for Amazon Q Apps].
//
// [Custom labels for Amazon Q Apps]: https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/qapps-custom-labels.html
func qapps_BatchUpdateCategory(cfg aws.Config, client *qapps.Client) {
	input := &qapps.BatchUpdateCategoryInput{
		// Categories: []types.CategoryInput, // Required
		// InstanceId: *string, // Required
	}

	if len(_qappsCategories) > 0 {
		if err := assignInputField(input, "Categories", _qappsCategories[0]); err != nil {
			log.Errorf("invalid --categories: %s", err.Error())
			return
		}
	}
	if len(_qappsInstanceId) > 0 {
		input.InstanceId = aws.String(_qappsInstanceId)
	}

	if resp, err := client.BatchUpdateCategory(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new library item for an Amazon Q App, allowing it to be discovered
// and used by other allowed users.
func qapps_CreateLibraryItem(cfg aws.Config, client *qapps.Client) {
	input := &qapps.CreateLibraryItemInput{
		// AppId: *string, // Required
		// AppVersion: *int32, // Required
		// Categories: []string, // Required
		// InstanceId: *string, // Required
	}

	if len(_qappsAppId) > 0 {
		input.AppId = aws.String(_qappsAppId)
	}
	if len(_qappsAppVersion) > 0 {
		if err := assignInputField(input, "AppVersion", _qappsAppVersion); err != nil {
			log.Errorf("invalid --app-version: %s", err.Error())
			return
		}
	}
	if len(_qappsCategories) > 0 {
		input.Categories = append([]string(nil), _qappsCategories...)
	}
	if len(_qappsInstanceId) > 0 {
		input.InstanceId = aws.String(_qappsInstanceId)
	}

	if resp, err := client.CreateLibraryItem(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a presigned URL for an S3 POST operation to upload a file. You can use
// this URL to set a default file for a FileUploadCard in a Q App definition or to
// provide a file for a single Q App run. The scope parameter determines how the
// file will be used, either at the app definition level or the app session level.
//
// The IAM permissions are derived from the qapps:ImportDocument action. For more
// information on the IAM policy for Amazon Q Apps, see [IAM permissions for using Amazon Q Apps].
//
// [IAM permissions for using Amazon Q Apps]: https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/deploy-q-apps-iam-permissions.html
func qapps_CreatePresignedUrl(cfg aws.Config, client *qapps.Client) {
	input := &qapps.CreatePresignedUrlInput{
		// AppId: *string, // Required
		// CardId: *string, // Required
		// FileContentsSha256: *string, // Required
		// FileName: *string, // Required
		// InstanceId: *string, // Required
		// Scope: types.DocumentScope, // Required
	}

	if len(_qappsAppId) > 0 {
		input.AppId = aws.String(_qappsAppId)
	}
	if len(_qappsCardId) > 0 {
		input.CardId = aws.String(_qappsCardId)
	}
	if len(_qappsFileContentsSha256) > 0 {
		input.FileContentsSha256 = aws.String(_qappsFileContentsSha256)
	}
	if len(_qappsFileName) > 0 {
		input.FileName = aws.String(_qappsFileName)
	}
	if len(_qappsInstanceId) > 0 {
		input.InstanceId = aws.String(_qappsInstanceId)
	}
	if len(_qappsScope) > 0 {
		if err := assignInputField(input, "Scope", _qappsScope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}
	if len(_qappsSessionId) > 0 {
		input.SessionId = aws.String(_qappsSessionId)
	}

	if resp, err := client.CreatePresignedUrl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new Amazon Q App based on the provided definition. The Q App
// definition specifies the cards and flow of the Q App. This operation also
// calculates the dependencies between the cards by inspecting the references in
// the prompts.
func qapps_CreateQApp(cfg aws.Config, client *qapps.Client) {
	input := &qapps.CreateQAppInput{
		// AppDefinition: *types.AppDefinitionInput, // Required
		// InstanceId: *string, // Required
		// Title: *string, // Required
	}

	if len(_qappsAppDefinition) > 0 {
		if err := assignInputField(input, "AppDefinition", _qappsAppDefinition); err != nil {
			log.Errorf("invalid --app-definition: %s", err.Error())
			return
		}
	}
	if len(_qappsInstanceId) > 0 {
		input.InstanceId = aws.String(_qappsInstanceId)
	}
	if len(_qappsTitle) > 0 {
		input.Title = aws.String(_qappsTitle)
	}
	if len(_qappsDescription) > 0 {
		input.Description = aws.String(_qappsDescription)
	}
	if len(_qappsTags) > 0 {
		if err := assignInputField(input, "Tags", _qappsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateQApp(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a library item for an Amazon Q App, removing it from the library so it
// can no longer be discovered or used by other users.
func qapps_DeleteLibraryItem(cfg aws.Config, client *qapps.Client) {
	input := &qapps.DeleteLibraryItemInput{
		// InstanceId: *string, // Required
		// LibraryItemId: *string, // Required
	}

	if len(_qappsInstanceId) > 0 {
		input.InstanceId = aws.String(_qappsInstanceId)
	}
	if len(_qappsLibraryItemId) > 0 {
		input.LibraryItemId = aws.String(_qappsLibraryItemId)
	}

	if resp, err := client.DeleteLibraryItem(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon Q App owned by the user. If the Q App was previously
// published to the library, it is also removed from the library.
func qapps_DeleteQApp(cfg aws.Config, client *qapps.Client) {
	input := &qapps.DeleteQAppInput{
		// AppId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_qappsAppId) > 0 {
		input.AppId = aws.String(_qappsAppId)
	}
	if len(_qappsInstanceId) > 0 {
		input.InstanceId = aws.String(_qappsInstanceId)
	}

	if resp, err := client.DeleteQApp(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes read permissions for a Amazon Q App in Amazon Q Business application
// environment instance.
func qapps_DescribeQAppPermissions(cfg aws.Config, client *qapps.Client) {
	input := &qapps.DescribeQAppPermissionsInput{
		// AppId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_qappsAppId) > 0 {
		input.AppId = aws.String(_qappsAppId)
	}
	if len(_qappsInstanceId) > 0 {
		input.InstanceId = aws.String(_qappsInstanceId)
	}

	if resp, err := client.DescribeQAppPermissions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a rating or review previously submitted by the user for a library item.
func qapps_DisassociateLibraryItemReview(cfg aws.Config, client *qapps.Client) {
	input := &qapps.DisassociateLibraryItemReviewInput{
		// InstanceId: *string, // Required
		// LibraryItemId: *string, // Required
	}

	if len(_qappsInstanceId) > 0 {
		input.InstanceId = aws.String(_qappsInstanceId)
	}
	if len(_qappsLibraryItemId) > 0 {
		input.LibraryItemId = aws.String(_qappsLibraryItemId)
	}

	if resp, err := client.DisassociateLibraryItemReview(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a Q App from a user removing the user's access to run the Q App.
func qapps_DisassociateQAppFromUser(cfg aws.Config, client *qapps.Client) {
	input := &qapps.DisassociateQAppFromUserInput{
		// AppId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_qappsAppId) > 0 {
		input.AppId = aws.String(_qappsAppId)
	}
	if len(_qappsInstanceId) > 0 {
		input.InstanceId = aws.String(_qappsInstanceId)
	}

	if resp, err := client.DisassociateQAppFromUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Exports the collected data of a Q App data collection session.
func qapps_ExportQAppSessionData(cfg aws.Config, client *qapps.Client) {
	input := &qapps.ExportQAppSessionDataInput{
		// InstanceId: *string, // Required
		// SessionId: *string, // Required
	}

	if len(_qappsInstanceId) > 0 {
		input.InstanceId = aws.String(_qappsInstanceId)
	}
	if len(_qappsSessionId) > 0 {
		input.SessionId = aws.String(_qappsSessionId)
	}

	if resp, err := client.ExportQAppSessionData(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details about a library item for an Amazon Q App, including its
// metadata, categories, ratings, and usage statistics.
func qapps_GetLibraryItem(cfg aws.Config, client *qapps.Client) {
	input := &qapps.GetLibraryItemInput{
		// InstanceId: *string, // Required
		// LibraryItemId: *string, // Required
	}

	if len(_qappsInstanceId) > 0 {
		input.InstanceId = aws.String(_qappsInstanceId)
	}
	if len(_qappsLibraryItemId) > 0 {
		input.LibraryItemId = aws.String(_qappsLibraryItemId)
	}
	if len(_qappsAppId) > 0 {
		input.AppId = aws.String(_qappsAppId)
	}

	if resp, err := client.GetLibraryItem(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the full details of an Q App, including its definition specifying the
// cards and flow.
func qapps_GetQApp(cfg aws.Config, client *qapps.Client) {
	input := &qapps.GetQAppInput{
		// AppId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_qappsAppId) > 0 {
		input.AppId = aws.String(_qappsAppId)
	}
	if len(_qappsInstanceId) > 0 {
		input.InstanceId = aws.String(_qappsInstanceId)
	}
	if len(_qappsAppVersion) > 0 {
		if err := assignInputField(input, "AppVersion", _qappsAppVersion); err != nil {
			log.Errorf("invalid --app-version: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetQApp(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the current state and results for an active session of an Amazon Q
// App.
func qapps_GetQAppSession(cfg aws.Config, client *qapps.Client) {
	input := &qapps.GetQAppSessionInput{
		// InstanceId: *string, // Required
		// SessionId: *string, // Required
	}

	if len(_qappsInstanceId) > 0 {
		input.InstanceId = aws.String(_qappsInstanceId)
	}
	if len(_qappsSessionId) > 0 {
		input.SessionId = aws.String(_qappsSessionId)
	}

	if resp, err := client.GetQAppSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the current configuration of a Q App session.
func qapps_GetQAppSessionMetadata(cfg aws.Config, client *qapps.Client) {
	input := &qapps.GetQAppSessionMetadataInput{
		// InstanceId: *string, // Required
		// SessionId: *string, // Required
	}

	if len(_qappsInstanceId) > 0 {
		input.InstanceId = aws.String(_qappsInstanceId)
	}
	if len(_qappsSessionId) > 0 {
		input.SessionId = aws.String(_qappsSessionId)
	}

	if resp, err := client.GetQAppSessionMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Uploads a file that can then be used either as a default in a FileUploadCard
// from Q App definition or as a file that is used inside a single Q App run. The
// purpose of the document is determined by a scope parameter that indicates
// whether it is at the app definition level or at the app session level.
func qapps_ImportDocument(cfg aws.Config, client *qapps.Client) {
	input := &qapps.ImportDocumentInput{
		// AppId: *string, // Required
		// CardId: *string, // Required
		// FileContentsBase64: *string, // Required
		// FileName: *string, // Required
		// InstanceId: *string, // Required
		// Scope: types.DocumentScope, // Required
	}

	if len(_qappsAppId) > 0 {
		input.AppId = aws.String(_qappsAppId)
	}
	if len(_qappsCardId) > 0 {
		input.CardId = aws.String(_qappsCardId)
	}
	if len(_qappsFileContentsBase64) > 0 {
		input.FileContentsBase64 = aws.String(_qappsFileContentsBase64)
	}
	if len(_qappsFileName) > 0 {
		input.FileName = aws.String(_qappsFileName)
	}
	if len(_qappsInstanceId) > 0 {
		input.InstanceId = aws.String(_qappsInstanceId)
	}
	if len(_qappsScope) > 0 {
		if err := assignInputField(input, "Scope", _qappsScope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}
	if len(_qappsSessionId) > 0 {
		input.SessionId = aws.String(_qappsSessionId)
	}

	if resp, err := client.ImportDocument(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the categories of a Amazon Q Business application environment instance.
// For more information, see [Custom labels for Amazon Q Apps].
//
// [Custom labels for Amazon Q Apps]: https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/qapps-custom-labels.html
func qapps_ListCategories(cfg aws.Config, client *qapps.Client) {
	input := &qapps.ListCategoriesInput{
		// InstanceId: *string, // Required
	}

	if len(_qappsInstanceId) > 0 {
		input.InstanceId = aws.String(_qappsInstanceId)
	}

	if resp, err := client.ListCategories(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the library items for Amazon Q Apps that are published and available for
// users in your Amazon Web Services account.
func qapps_ListLibraryItems(cfg aws.Config, client *qapps.Client) {
	input := &qapps.ListLibraryItemsInput{
		// InstanceId: *string, // Required
	}

	if len(_qappsInstanceId) > 0 {
		input.InstanceId = aws.String(_qappsInstanceId)
	}
	if len(_qappsCategoryId) > 0 {
		input.CategoryId = aws.String(_qappsCategoryId)
	}
	if len(_qappsLimit) > 0 {
		if err := assignInputField(input, "Limit", _qappsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_qappsNextToken) > 0 {
		input.NextToken = aws.String(_qappsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListLibraryItems(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*qapps.ListLibraryItemsOutput
	p := qapps.NewListLibraryItemsPaginator(client, input)
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

// Lists the collected data of a Q App data collection session.
func qapps_ListQAppSessionData(cfg aws.Config, client *qapps.Client) {
	input := &qapps.ListQAppSessionDataInput{
		// InstanceId: *string, // Required
		// SessionId: *string, // Required
	}

	if len(_qappsInstanceId) > 0 {
		input.InstanceId = aws.String(_qappsInstanceId)
	}
	if len(_qappsSessionId) > 0 {
		input.SessionId = aws.String(_qappsSessionId)
	}

	if resp, err := client.ListQAppSessionData(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the Amazon Q Apps owned by or associated with the user either because
// they created it or because they used it from the library in the past. The user
// identity is extracted from the credentials used to invoke this operation..
func qapps_ListQApps(cfg aws.Config, client *qapps.Client) {
	input := &qapps.ListQAppsInput{
		// InstanceId: *string, // Required
	}

	if len(_qappsInstanceId) > 0 {
		input.InstanceId = aws.String(_qappsInstanceId)
	}
	if len(_qappsLimit) > 0 {
		if err := assignInputField(input, "Limit", _qappsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_qappsNextToken) > 0 {
		input.NextToken = aws.String(_qappsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListQApps(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*qapps.ListQAppsOutput
	p := qapps.NewListQAppsPaginator(client, input)
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

// Lists the tags associated with an Amazon Q Apps resource.
func qapps_ListTagsForResource(cfg aws.Config, client *qapps.Client) {
	input := &qapps.ListTagsForResourceInput{
		// ResourceARN: *string, // Required
	}

	if len(_qappsResourceARN) > 0 {
		input.ResourceARN = aws.String(_qappsResourceARN)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Generates an Amazon Q App definition based on either a conversation or a
// problem statement provided as input.The resulting app definition can be used to
// call CreateQApp . This API doesn't create Amazon Q Apps directly.
func qapps_PredictQApp(cfg aws.Config, client *qapps.Client) {
	input := &qapps.PredictQAppInput{
		// InstanceId: *string, // Required
	}

	if len(_qappsInstanceId) > 0 {
		input.InstanceId = aws.String(_qappsInstanceId)
	}
	if len(_qappsOptions) > 0 {
		if err := assignInputField(input, "Options", _qappsOptions); err != nil {
			log.Errorf("invalid --options: %s", err.Error())
			return
		}
	}

	if resp, err := client.PredictQApp(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a new session for an Amazon Q App, allowing inputs to be provided and
// the app to be run.
//
// Each Q App session will be condensed into a single conversation in the web
// experience.
func qapps_StartQAppSession(cfg aws.Config, client *qapps.Client) {
	input := &qapps.StartQAppSessionInput{
		// AppId: *string, // Required
		// AppVersion: *int32, // Required
		// InstanceId: *string, // Required
	}

	if len(_qappsAppId) > 0 {
		input.AppId = aws.String(_qappsAppId)
	}
	if len(_qappsAppVersion) > 0 {
		if err := assignInputField(input, "AppVersion", _qappsAppVersion); err != nil {
			log.Errorf("invalid --app-version: %s", err.Error())
			return
		}
	}
	if len(_qappsInstanceId) > 0 {
		input.InstanceId = aws.String(_qappsInstanceId)
	}
	if len(_qappsInitialValues) > 0 {
		if err := assignInputField(input, "InitialValues", _qappsInitialValues); err != nil {
			log.Errorf("invalid --initial-values: %s", err.Error())
			return
		}
	}
	if len(_qappsSessionId) > 0 {
		input.SessionId = aws.String(_qappsSessionId)
	}
	if len(_qappsTags) > 0 {
		if err := assignInputField(input, "Tags", _qappsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartQAppSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops an active session for an Amazon Q App.This deletes all data related to
// the session and makes it invalid for future uses. The results of the session
// will be persisted as part of the conversation.
func qapps_StopQAppSession(cfg aws.Config, client *qapps.Client) {
	input := &qapps.StopQAppSessionInput{
		// InstanceId: *string, // Required
		// SessionId: *string, // Required
	}

	if len(_qappsInstanceId) > 0 {
		input.InstanceId = aws.String(_qappsInstanceId)
	}
	if len(_qappsSessionId) > 0 {
		input.SessionId = aws.String(_qappsSessionId)
	}

	if resp, err := client.StopQAppSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates tags with an Amazon Q Apps resource.
func qapps_TagResource(cfg aws.Config, client *qapps.Client) {
	input := &qapps.TagResourceInput{
		// ResourceARN: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_qappsResourceARN) > 0 {
		input.ResourceARN = aws.String(_qappsResourceARN)
	}
	if len(_qappsTags) > 0 {
		if err := assignInputField(input, "Tags", _qappsTags); err != nil {
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

// Disassociates tags from an Amazon Q Apps resource.
func qapps_UntagResource(cfg aws.Config, client *qapps.Client) {
	input := &qapps.UntagResourceInput{
		// ResourceARN: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_qappsResourceARN) > 0 {
		input.ResourceARN = aws.String(_qappsResourceARN)
	}
	if len(_qappsTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _qappsTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the library item for an Amazon Q App.
func qapps_UpdateLibraryItem(cfg aws.Config, client *qapps.Client) {
	input := &qapps.UpdateLibraryItemInput{
		// InstanceId: *string, // Required
		// LibraryItemId: *string, // Required
	}

	if len(_qappsInstanceId) > 0 {
		input.InstanceId = aws.String(_qappsInstanceId)
	}
	if len(_qappsLibraryItemId) > 0 {
		input.LibraryItemId = aws.String(_qappsLibraryItemId)
	}
	if len(_qappsCategories) > 0 {
		input.Categories = append([]string(nil), _qappsCategories...)
	}
	if len(_qappsStatus) > 0 {
		if err := assignInputField(input, "Status", _qappsStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateLibraryItem(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the verification status of a library item for an Amazon Q App.
func qapps_UpdateLibraryItemMetadata(cfg aws.Config, client *qapps.Client) {
	input := &qapps.UpdateLibraryItemMetadataInput{
		// InstanceId: *string, // Required
		// LibraryItemId: *string, // Required
	}

	if len(_qappsInstanceId) > 0 {
		input.InstanceId = aws.String(_qappsInstanceId)
	}
	if len(_qappsLibraryItemId) > 0 {
		input.LibraryItemId = aws.String(_qappsLibraryItemId)
	}
	if len(_qappsIsVerified) > 0 {
		if err := assignInputField(input, "IsVerified", _qappsIsVerified); err != nil {
			log.Errorf("invalid --is-verified: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateLibraryItemMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing Amazon Q App, allowing modifications to its title,
// description, and definition.
func qapps_UpdateQApp(cfg aws.Config, client *qapps.Client) {
	input := &qapps.UpdateQAppInput{
		// AppId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_qappsAppId) > 0 {
		input.AppId = aws.String(_qappsAppId)
	}
	if len(_qappsInstanceId) > 0 {
		input.InstanceId = aws.String(_qappsInstanceId)
	}
	if len(_qappsAppDefinition) > 0 {
		if err := assignInputField(input, "AppDefinition", _qappsAppDefinition); err != nil {
			log.Errorf("invalid --app-definition: %s", err.Error())
			return
		}
	}
	if len(_qappsDescription) > 0 {
		input.Description = aws.String(_qappsDescription)
	}
	if len(_qappsTitle) > 0 {
		input.Title = aws.String(_qappsTitle)
	}

	if resp, err := client.UpdateQApp(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates read permissions for a Amazon Q App in Amazon Q Business application
// environment instance.
func qapps_UpdateQAppPermissions(cfg aws.Config, client *qapps.Client) {
	input := &qapps.UpdateQAppPermissionsInput{
		// AppId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_qappsAppId) > 0 {
		input.AppId = aws.String(_qappsAppId)
	}
	if len(_qappsInstanceId) > 0 {
		input.InstanceId = aws.String(_qappsInstanceId)
	}
	if len(_qappsGrantPermissions) > 0 {
		if err := assignInputField(input, "GrantPermissions", _qappsGrantPermissions); err != nil {
			log.Errorf("invalid --grant-permissions: %s", err.Error())
			return
		}
	}
	if len(_qappsRevokePermissions) > 0 {
		if err := assignInputField(input, "RevokePermissions", _qappsRevokePermissions); err != nil {
			log.Errorf("invalid --revoke-permissions: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateQAppPermissions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the session for a given Q App sessionId . This is only valid when at
// least one card of the session is in the WAITING state. Data for each WAITING
// card can be provided as input. If inputs are not provided, the call will be
// accepted but session will not move forward. Inputs for cards that are not in the
// WAITING status will be ignored.
func qapps_UpdateQAppSession(cfg aws.Config, client *qapps.Client) {
	input := &qapps.UpdateQAppSessionInput{
		// InstanceId: *string, // Required
		// SessionId: *string, // Required
	}

	if len(_qappsInstanceId) > 0 {
		input.InstanceId = aws.String(_qappsInstanceId)
	}
	if len(_qappsSessionId) > 0 {
		input.SessionId = aws.String(_qappsSessionId)
	}
	if len(_qappsValues) > 0 {
		if err := assignInputField(input, "Values", _qappsValues); err != nil {
			log.Errorf("invalid --values: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateQAppSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration metadata of a session for a given Q App sessionId .
func qapps_UpdateQAppSessionMetadata(cfg aws.Config, client *qapps.Client) {
	input := &qapps.UpdateQAppSessionMetadataInput{
		// InstanceId: *string, // Required
		// SessionId: *string, // Required
		// SharingConfiguration: *types.SessionSharingConfiguration, // Required
	}

	if len(_qappsInstanceId) > 0 {
		input.InstanceId = aws.String(_qappsInstanceId)
	}
	if len(_qappsSessionId) > 0 {
		input.SessionId = aws.String(_qappsSessionId)
	}
	if len(_qappsSharingConfiguration) > 0 {
		if err := assignInputField(input, "SharingConfiguration", _qappsSharingConfiguration); err != nil {
			log.Errorf("invalid --sharing-configuration: %s", err.Error())
			return
		}
	}
	if len(_qappsSessionName) > 0 {
		input.SessionName = aws.String(_qappsSessionName)
	}

	if resp, err := client.UpdateQAppSessionMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_qappsCmd)
	_qappsCmd.Flags().SortFlags = false

	_qappsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_qappsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_qappsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_qappsCmd.Flags().StringVarP(&_qappsAppDefinition, "app-definition", "", "", "App Definition")
	_qappsCmd.Flags().StringVarP(&_qappsAppId, "app-id", "", "", "App ID")
	_qappsCmd.Flags().StringVarP(&_qappsAppVersion, "app-version", "", "", "App Version")
	_qappsCmd.Flags().StringVarP(&_qappsCardId, "card-id", "", "", "Card ID")
	_qappsCmd.Flags().StringSliceVarP(&_qappsCategories, "categories", "", nil, "Categories")
	_qappsCmd.Flags().StringVarP(&_qappsCategoryId, "category-id", "", "", "Category ID")
	_qappsCmd.Flags().StringVarP(&_qappsDescription, "description", "", "", "Description")
	_qappsCmd.Flags().StringVarP(&_qappsFileContentsBase64, "file-contents-base64", "", "", "File Contents Base64")
	_qappsCmd.Flags().StringVarP(&_qappsFileContentsSha256, "file-contents-sha256", "", "", "File Contents SHA256")
	_qappsCmd.Flags().StringVarP(&_qappsFileName, "file-name", "", "", "File Name")
	_qappsCmd.Flags().StringVarP(&_qappsGrantPermissions, "grant-permissions", "", "", "Grant Permissions")
	_qappsCmd.Flags().StringVarP(&_qappsInitialValues, "initial-values", "", "", "Initial Values")
	_qappsCmd.Flags().StringVarP(&_qappsInstanceId, "instance-id", "", "", "Instance ID")
	_qappsCmd.Flags().StringVarP(&_qappsIsVerified, "is-verified", "", "", "Is Verified")
	_qappsCmd.Flags().StringVarP(&_qappsLibraryItemId, "library-item-id", "", "", "Library Item ID")
	_qappsCmd.Flags().StringVarP(&_qappsLimit, "limit", "", "", "Limit")
	_qappsCmd.Flags().StringVarP(&_qappsNextToken, "next-token", "", "", "Next Token")
	_qappsCmd.Flags().StringVarP(&_qappsOptions, "options", "", "", "Options")
	_qappsCmd.Flags().StringVarP(&_qappsResourceARN, "resource-arn", "", "", "Resource ARN")
	_qappsCmd.Flags().StringVarP(&_qappsRevokePermissions, "revoke-permissions", "", "", "Revoke Permissions")
	_qappsCmd.Flags().StringVarP(&_qappsScope, "scope", "", "", "Scope")
	_qappsCmd.Flags().StringVarP(&_qappsSessionId, "session-id", "", "", "Session ID")
	_qappsCmd.Flags().StringVarP(&_qappsSessionName, "session-name", "", "", "Session Name")
	_qappsCmd.Flags().StringVarP(&_qappsSharingConfiguration, "sharing-configuration", "", "", "Sharing Configuration")
	_qappsCmd.Flags().StringVarP(&_qappsStatus, "status", "", "", "Status")
	_qappsCmd.Flags().StringSliceVarP(&_qappsTagKeys, "tag-keys", "", nil, "Tag Keys")
	_qappsCmd.Flags().StringVarP(&_qappsTags, "tags", "", "", "Tags")
	_qappsCmd.Flags().StringVarP(&_qappsTitle, "title", "", "", "Title")
	_qappsCmd.Flags().StringVarP(&_qappsValues, "values", "", "", "Values")

	_qappsCmd.Flags().BoolVarP(&_qappsAssociateLibraryItemReview, "associate-library-item-review", "", false, "Associate Library Item Review")
	_qappsCmd.Flags().BoolVarP(&_qappsAssociateQAppWithUser, "associate-qapp-with-user", "", false, "Associate Qapp With User")
	_qappsCmd.Flags().BoolVarP(&_qappsBatchCreateCategory, "batch-create-category", "", false, "Batch Create Category")
	_qappsCmd.Flags().BoolVarP(&_qappsBatchDeleteCategory, "batch-delete-category", "", false, "Batch Delete Category")
	_qappsCmd.Flags().BoolVarP(&_qappsBatchUpdateCategory, "batch-update-category", "", false, "Batch Update Category")
	_qappsCmd.Flags().BoolVarP(&_qappsCreateLibraryItem, "create-library-item", "", false, "Create Library Item")
	_qappsCmd.Flags().BoolVarP(&_qappsCreatePresignedUrl, "create-presigned-url", "", false, "Create Presigned URL")
	_qappsCmd.Flags().BoolVarP(&_qappsCreateQApp, "create-qapp", "", false, "Create Qapp")
	_qappsCmd.Flags().BoolVarP(&_qappsDeleteLibraryItem, "delete-library-item", "", false, "Delete Library Item")
	_qappsCmd.Flags().BoolVarP(&_qappsDeleteQApp, "delete-qapp", "", false, "Delete Qapp")
	_qappsCmd.Flags().BoolVarP(&_qappsDescribeQAppPermissions, "describe-qapp-permissions", "", false, "Describe Qapp Permissions")
	_qappsCmd.Flags().BoolVarP(&_qappsDisassociateLibraryItemReview, "disassociate-library-item-review", "", false, "Disassociate Library Item Review")
	_qappsCmd.Flags().BoolVarP(&_qappsDisassociateQAppFromUser, "disassociate-qapp-from-user", "", false, "Disassociate Qapp From User")
	_qappsCmd.Flags().BoolVarP(&_qappsExportQAppSessionData, "export-qapp-session-data", "", false, "Export Qapp Session Data")
	_qappsCmd.Flags().BoolVarP(&_qappsGetLibraryItem, "get-library-item", "", false, "Get Library Item")
	_qappsCmd.Flags().BoolVarP(&_qappsGetQApp, "get-qapp", "", false, "Get Qapp")
	_qappsCmd.Flags().BoolVarP(&_qappsGetQAppSession, "get-qapp-session", "", false, "Get Qapp Session")
	_qappsCmd.Flags().BoolVarP(&_qappsGetQAppSessionMetadata, "get-qapp-session-metadata", "", false, "Get Qapp Session Metadata")
	_qappsCmd.Flags().BoolVarP(&_qappsImportDocument, "import-document", "", false, "Import Document")
	_qappsCmd.Flags().BoolVarP(&_qappsListCategories, "list-categories", "", false, "List Categories")
	_qappsCmd.Flags().BoolVarP(&_qappsListLibraryItems, "list-library-items", "", false, "List Library Items")
	_qappsCmd.Flags().BoolVarP(&_qappsListQAppSessionData, "list-qapp-session-data", "", false, "List Qapp Session Data")
	_qappsCmd.Flags().BoolVarP(&_qappsListQApps, "list-qapps", "", false, "List Qapps")
	_qappsCmd.Flags().BoolVarP(&_qappsListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_qappsCmd.Flags().BoolVarP(&_qappsPredictQApp, "predict-qapp", "", false, "Predict Qapp")
	_qappsCmd.Flags().BoolVarP(&_qappsStartQAppSession, "start-qapp-session", "", false, "Start Qapp Session")
	_qappsCmd.Flags().BoolVarP(&_qappsStopQAppSession, "stop-qapp-session", "", false, "Stop Qapp Session")
	_qappsCmd.Flags().BoolVarP(&_qappsTagResource, "tag-resource", "", false, "Tag Resource")
	_qappsCmd.Flags().BoolVarP(&_qappsUntagResource, "untag-resource", "", false, "Untag Resource")
	_qappsCmd.Flags().BoolVarP(&_qappsUpdateLibraryItem, "update-library-item", "", false, "Update Library Item")
	_qappsCmd.Flags().BoolVarP(&_qappsUpdateLibraryItemMetadata, "update-library-item-metadata", "", false, "Update Library Item Metadata")
	_qappsCmd.Flags().BoolVarP(&_qappsUpdateQApp, "update-qapp", "", false, "Update Qapp")
	_qappsCmd.Flags().BoolVarP(&_qappsUpdateQAppPermissions, "update-qapp-permissions", "", false, "Update Qapp Permissions")
	_qappsCmd.Flags().BoolVarP(&_qappsUpdateQAppSession, "update-qapp-session", "", false, "Update Qapp Session")
	_qappsCmd.Flags().BoolVarP(&_qappsUpdateQAppSessionMetadata, "update-qapp-session-metadata", "", false, "Update Qapp Session Metadata")

}
