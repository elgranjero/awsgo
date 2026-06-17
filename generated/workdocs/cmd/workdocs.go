package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/workdocs"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// workdocsCmd represents the workdocs command
var _workdocsCmd = &cobra.Command{
	Use:   "workdocs",
	Short: "AWS workdocs CLI",
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
		client := workdocs.NewFromConfig(cfg)
		if _workdocsAbortDocumentVersionUpload {
			workdocs_AbortDocumentVersionUpload(cfg, client)
			return
		}
		if _workdocsActivateUser {
			workdocs_ActivateUser(cfg, client)
			return
		}
		if _workdocsAddResourcePermissions {
			workdocs_AddResourcePermissions(cfg, client)
			return
		}
		if _workdocsCreateComment {
			workdocs_CreateComment(cfg, client)
			return
		}
		if _workdocsCreateCustomMetadata {
			workdocs_CreateCustomMetadata(cfg, client)
			return
		}
		if _workdocsCreateFolder {
			workdocs_CreateFolder(cfg, client)
			return
		}
		if _workdocsCreateLabels {
			workdocs_CreateLabels(cfg, client)
			return
		}
		if _workdocsCreateNotificationSubscription {
			workdocs_CreateNotificationSubscription(cfg, client)
			return
		}
		if _workdocsCreateUser {
			workdocs_CreateUser(cfg, client)
			return
		}
		if _workdocsDeactivateUser {
			workdocs_DeactivateUser(cfg, client)
			return
		}
		if _workdocsDeleteComment {
			workdocs_DeleteComment(cfg, client)
			return
		}
		if _workdocsDeleteCustomMetadata {
			workdocs_DeleteCustomMetadata(cfg, client)
			return
		}
		if _workdocsDeleteDocument {
			workdocs_DeleteDocument(cfg, client)
			return
		}
		if _workdocsDeleteDocumentVersion {
			workdocs_DeleteDocumentVersion(cfg, client)
			return
		}
		if _workdocsDeleteFolder {
			workdocs_DeleteFolder(cfg, client)
			return
		}
		if _workdocsDeleteFolderContents {
			workdocs_DeleteFolderContents(cfg, client)
			return
		}
		if _workdocsDeleteLabels {
			workdocs_DeleteLabels(cfg, client)
			return
		}
		if _workdocsDeleteNotificationSubscription {
			workdocs_DeleteNotificationSubscription(cfg, client)
			return
		}
		if _workdocsDeleteUser {
			workdocs_DeleteUser(cfg, client)
			return
		}
		if _workdocsDescribeActivities {
			workdocs_DescribeActivities(cfg, client)
			return
		}
		if _workdocsDescribeComments {
			workdocs_DescribeComments(cfg, client)
			return
		}
		if _workdocsDescribeDocumentVersions {
			workdocs_DescribeDocumentVersions(cfg, client)
			return
		}
		if _workdocsDescribeFolderContents {
			workdocs_DescribeFolderContents(cfg, client)
			return
		}
		if _workdocsDescribeGroups {
			workdocs_DescribeGroups(cfg, client)
			return
		}
		if _workdocsDescribeNotificationSubscriptions {
			workdocs_DescribeNotificationSubscriptions(cfg, client)
			return
		}
		if _workdocsDescribeResourcePermissions {
			workdocs_DescribeResourcePermissions(cfg, client)
			return
		}
		if _workdocsDescribeRootFolders {
			workdocs_DescribeRootFolders(cfg, client)
			return
		}
		if _workdocsDescribeUsers {
			workdocs_DescribeUsers(cfg, client)
			return
		}
		if _workdocsGetCurrentUser {
			workdocs_GetCurrentUser(cfg, client)
			return
		}
		if _workdocsGetDocument {
			workdocs_GetDocument(cfg, client)
			return
		}
		if _workdocsGetDocumentPath {
			workdocs_GetDocumentPath(cfg, client)
			return
		}
		if _workdocsGetDocumentVersion {
			workdocs_GetDocumentVersion(cfg, client)
			return
		}
		if _workdocsGetFolder {
			workdocs_GetFolder(cfg, client)
			return
		}
		if _workdocsGetFolderPath {
			workdocs_GetFolderPath(cfg, client)
			return
		}
		if _workdocsGetResources {
			workdocs_GetResources(cfg, client)
			return
		}
		if _workdocsInitiateDocumentVersionUpload {
			workdocs_InitiateDocumentVersionUpload(cfg, client)
			return
		}
		if _workdocsRemoveAllResourcePermissions {
			workdocs_RemoveAllResourcePermissions(cfg, client)
			return
		}
		if _workdocsRemoveResourcePermission {
			workdocs_RemoveResourcePermission(cfg, client)
			return
		}
		if _workdocsRestoreDocumentVersions {
			workdocs_RestoreDocumentVersions(cfg, client)
			return
		}
		if _workdocsSearchResources {
			workdocs_SearchResources(cfg, client)
			return
		}
		if _workdocsUpdateDocument {
			workdocs_UpdateDocument(cfg, client)
			return
		}
		if _workdocsUpdateDocumentVersion {
			workdocs_UpdateDocumentVersion(cfg, client)
			return
		}
		if _workdocsUpdateFolder {
			workdocs_UpdateFolder(cfg, client)
			return
		}
		if _workdocsUpdateUser {
			workdocs_UpdateUser(cfg, client)
			return
		}

	},
}

var (
	_workdocsAbortDocumentVersionUpload        bool
	_workdocsActivateUser                      bool
	_workdocsAddResourcePermissions            bool
	_workdocsCreateComment                     bool
	_workdocsCreateCustomMetadata              bool
	_workdocsCreateFolder                      bool
	_workdocsCreateLabels                      bool
	_workdocsCreateNotificationSubscription    bool
	_workdocsCreateUser                        bool
	_workdocsDeactivateUser                    bool
	_workdocsDeleteComment                     bool
	_workdocsDeleteCustomMetadata              bool
	_workdocsDeleteDocument                    bool
	_workdocsDeleteDocumentVersion             bool
	_workdocsDeleteFolder                      bool
	_workdocsDeleteFolderContents              bool
	_workdocsDeleteLabels                      bool
	_workdocsDeleteNotificationSubscription    bool
	_workdocsDeleteUser                        bool
	_workdocsDescribeActivities                bool
	_workdocsDescribeComments                  bool
	_workdocsDescribeDocumentVersions          bool
	_workdocsDescribeFolderContents            bool
	_workdocsDescribeGroups                    bool
	_workdocsDescribeNotificationSubscriptions bool
	_workdocsDescribeResourcePermissions       bool
	_workdocsDescribeRootFolders               bool
	_workdocsDescribeUsers                     bool
	_workdocsGetCurrentUser                    bool
	_workdocsGetDocument                       bool
	_workdocsGetDocumentPath                   bool
	_workdocsGetDocumentVersion                bool
	_workdocsGetFolder                         bool
	_workdocsGetFolderPath                     bool
	_workdocsGetResources                      bool
	_workdocsInitiateDocumentVersionUpload     bool
	_workdocsRemoveAllResourcePermissions      bool
	_workdocsRemoveResourcePermission          bool
	_workdocsRestoreDocumentVersions           bool
	_workdocsSearchResources                   bool
	_workdocsUpdateDocument                    bool
	_workdocsUpdateDocumentVersion             bool
	_workdocsUpdateFolder                      bool
	_workdocsUpdateUser                        bool

	_workdocsActivityTypes             string
	_workdocsAdditionalResponseFields  string
	_workdocsAuthenticationToken       string
	_workdocsCollectionType            string
	_workdocsCommentId                 string
	_workdocsContentCreatedTimestamp   string
	_workdocsContentModifiedTimestamp  string
	_workdocsContentType               string
	_workdocsCustomMetadata            string
	_workdocsDeleteAll                 string
	_workdocsDeletePriorVersions       string
	_workdocsDocumentId                string
	_workdocsDocumentSizeInBytes       string
	_workdocsEmailAddress              string
	_workdocsEndTime                   string
	_workdocsEndpoint                  string
	_workdocsFields                    string
	_workdocsFilters                   string
	_workdocsFolderId                  string
	_workdocsGivenName                 string
	_workdocsGrantPoweruserPrivileges  string
	_workdocsId                        string
	_workdocsInclude                   string
	_workdocsIncludeCustomMetadata     string
	_workdocsIncludeIndirectActivities string
	_workdocsKeys                      []string
	_workdocsLabels                    []string
	_workdocsLimit                     string
	_workdocsLocale                    string
	_workdocsMarker                    string
	_workdocsName                      string
	_workdocsNotificationOptions       string
	_workdocsNotifyCollaborators       string
	_workdocsOrder                     string
	_workdocsOrderBy                   string
	_workdocsOrganizationId            string
	_workdocsParentFolderId            string
	_workdocsParentId                  string
	_workdocsPassword                  string
	_workdocsPrincipalId               string
	_workdocsPrincipalType             string
	_workdocsPrincipals                string
	_workdocsProtocol                  string
	_workdocsQuery                     string
	_workdocsQueryScopes               string
	_workdocsQueryText                 string
	_workdocsResourceId                string
	_workdocsResourceState             string
	_workdocsSearchQuery               string
	_workdocsSort                      string
	_workdocsStartTime                 string
	_workdocsStorageRule               string
	_workdocsSubscriptionId            string
	_workdocsSubscriptionType          string
	_workdocsSurname                   string
	_workdocsText                      string
	_workdocsThreadId                  string
	_workdocsTimeZoneId                string
	_workdocsType                      string
	_workdocsUserId                    string
	_workdocsUserIds                   string
	_workdocsUsername                  string
	_workdocsVersionId                 string
	_workdocsVersionStatus             string
	_workdocsVisibility                string
)

// Aborts the upload of the specified document version that was previously
// initiated by InitiateDocumentVersionUpload. The client should make this call only when it no longer intends
// to upload the document version, or fails to do so.
func workdocs_AbortDocumentVersionUpload(cfg aws.Config, client *workdocs.Client) {
	input := &workdocs.AbortDocumentVersionUploadInput{
		// DocumentId: *string, // Required
		// VersionId: *string, // Required
	}

	if len(_workdocsDocumentId) > 0 {
		input.DocumentId = aws.String(_workdocsDocumentId)
	}
	if len(_workdocsVersionId) > 0 {
		input.VersionId = aws.String(_workdocsVersionId)
	}
	if len(_workdocsAuthenticationToken) > 0 {
		input.AuthenticationToken = aws.String(_workdocsAuthenticationToken)
	}

	if resp, err := client.AbortDocumentVersionUpload(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Activates the specified user. Only active users can access Amazon WorkDocs.
func workdocs_ActivateUser(cfg aws.Config, client *workdocs.Client) {
	input := &workdocs.ActivateUserInput{
		// UserId: *string, // Required
	}

	if len(_workdocsUserId) > 0 {
		input.UserId = aws.String(_workdocsUserId)
	}
	if len(_workdocsAuthenticationToken) > 0 {
		input.AuthenticationToken = aws.String(_workdocsAuthenticationToken)
	}

	if resp, err := client.ActivateUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a set of permissions for the specified folder or document. The resource
// permissions are overwritten if the principals already have different
// permissions.
func workdocs_AddResourcePermissions(cfg aws.Config, client *workdocs.Client) {
	input := &workdocs.AddResourcePermissionsInput{
		// Principals: []types.SharePrincipal, // Required
		// ResourceId: *string, // Required
	}

	if len(_workdocsPrincipals) > 0 {
		if err := assignInputField(input, "Principals", _workdocsPrincipals); err != nil {
			log.Errorf("invalid --principals: %s", err.Error())
			return
		}
	}
	if len(_workdocsResourceId) > 0 {
		input.ResourceId = aws.String(_workdocsResourceId)
	}
	if len(_workdocsAuthenticationToken) > 0 {
		input.AuthenticationToken = aws.String(_workdocsAuthenticationToken)
	}
	if len(_workdocsNotificationOptions) > 0 {
		if err := assignInputField(input, "NotificationOptions", _workdocsNotificationOptions); err != nil {
			log.Errorf("invalid --notification-options: %s", err.Error())
			return
		}
	}

	if resp, err := client.AddResourcePermissions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a new comment to the specified document version.
func workdocs_CreateComment(cfg aws.Config, client *workdocs.Client) {
	input := &workdocs.CreateCommentInput{
		// DocumentId: *string, // Required
		// Text: *string, // Required
		// VersionId: *string, // Required
	}

	if len(_workdocsDocumentId) > 0 {
		input.DocumentId = aws.String(_workdocsDocumentId)
	}
	if len(_workdocsText) > 0 {
		input.Text = aws.String(_workdocsText)
	}
	if len(_workdocsVersionId) > 0 {
		input.VersionId = aws.String(_workdocsVersionId)
	}
	if len(_workdocsAuthenticationToken) > 0 {
		input.AuthenticationToken = aws.String(_workdocsAuthenticationToken)
	}
	if len(_workdocsNotifyCollaborators) > 0 {
		if err := assignInputField(input, "NotifyCollaborators", _workdocsNotifyCollaborators); err != nil {
			log.Errorf("invalid --notify-collaborators: %s", err.Error())
			return
		}
	}
	if len(_workdocsParentId) > 0 {
		input.ParentId = aws.String(_workdocsParentId)
	}
	if len(_workdocsThreadId) > 0 {
		input.ThreadId = aws.String(_workdocsThreadId)
	}
	if len(_workdocsVisibility) > 0 {
		if err := assignInputField(input, "Visibility", _workdocsVisibility); err != nil {
			log.Errorf("invalid --visibility: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateComment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds one or more custom properties to the specified resource (a folder,
// document, or version).
func workdocs_CreateCustomMetadata(cfg aws.Config, client *workdocs.Client) {
	input := &workdocs.CreateCustomMetadataInput{
		// CustomMetadata: map[string]string, // Required
		// ResourceId: *string, // Required
	}

	if len(_workdocsCustomMetadata) > 0 {
		if err := assignInputField(input, "CustomMetadata", _workdocsCustomMetadata); err != nil {
			log.Errorf("invalid --custom-metadata: %s", err.Error())
			return
		}
	}
	if len(_workdocsResourceId) > 0 {
		input.ResourceId = aws.String(_workdocsResourceId)
	}
	if len(_workdocsAuthenticationToken) > 0 {
		input.AuthenticationToken = aws.String(_workdocsAuthenticationToken)
	}
	if len(_workdocsVersionId) > 0 {
		input.VersionId = aws.String(_workdocsVersionId)
	}

	if resp, err := client.CreateCustomMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a folder with the specified name and parent folder.
func workdocs_CreateFolder(cfg aws.Config, client *workdocs.Client) {
	input := &workdocs.CreateFolderInput{
		// ParentFolderId: *string, // Required
	}

	if len(_workdocsParentFolderId) > 0 {
		input.ParentFolderId = aws.String(_workdocsParentFolderId)
	}
	if len(_workdocsAuthenticationToken) > 0 {
		input.AuthenticationToken = aws.String(_workdocsAuthenticationToken)
	}
	if len(_workdocsName) > 0 {
		input.Name = aws.String(_workdocsName)
	}

	if resp, err := client.CreateFolder(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds the specified list of labels to the given resource (a document or folder)
func workdocs_CreateLabels(cfg aws.Config, client *workdocs.Client) {
	input := &workdocs.CreateLabelsInput{
		// Labels: []string, // Required
		// ResourceId: *string, // Required
	}

	if len(_workdocsLabels) > 0 {
		input.Labels = append([]string(nil), _workdocsLabels...)
	}
	if len(_workdocsResourceId) > 0 {
		input.ResourceId = aws.String(_workdocsResourceId)
	}
	if len(_workdocsAuthenticationToken) > 0 {
		input.AuthenticationToken = aws.String(_workdocsAuthenticationToken)
	}

	if resp, err := client.CreateLabels(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Configure Amazon WorkDocs to use Amazon SNS notifications. The endpoint
// receives a confirmation message, and must confirm the subscription.
//
// For more information, see [Setting up notifications for an IAM user or role] in the Amazon WorkDocs Developer Guide.
//
// [Setting up notifications for an IAM user or role]: https://docs.aws.amazon.com/workdocs/latest/developerguide/manage-notifications.html
func workdocs_CreateNotificationSubscription(cfg aws.Config, client *workdocs.Client) {
	input := &workdocs.CreateNotificationSubscriptionInput{
		// Endpoint: *string, // Required
		// OrganizationId: *string, // Required
		// Protocol: types.SubscriptionProtocolType, // Required
		// SubscriptionType: types.SubscriptionType, // Required
	}

	if len(_workdocsEndpoint) > 0 {
		input.Endpoint = aws.String(_workdocsEndpoint)
	}
	if len(_workdocsOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workdocsOrganizationId)
	}
	if len(_workdocsProtocol) > 0 {
		if err := assignInputField(input, "Protocol", _workdocsProtocol); err != nil {
			log.Errorf("invalid --protocol: %s", err.Error())
			return
		}
	}
	if len(_workdocsSubscriptionType) > 0 {
		if err := assignInputField(input, "SubscriptionType", _workdocsSubscriptionType); err != nil {
			log.Errorf("invalid --subscription-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateNotificationSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a user in a Simple AD or Microsoft AD directory. The status of a newly
// created user is "ACTIVE". New users can access Amazon WorkDocs.
func workdocs_CreateUser(cfg aws.Config, client *workdocs.Client) {
	input := &workdocs.CreateUserInput{
		// GivenName: *string, // Required
		// Password: *string, // Required
		// Surname: *string, // Required
		// Username: *string, // Required
	}

	if len(_workdocsGivenName) > 0 {
		input.GivenName = aws.String(_workdocsGivenName)
	}
	if len(_workdocsPassword) > 0 {
		input.Password = aws.String(_workdocsPassword)
	}
	if len(_workdocsSurname) > 0 {
		input.Surname = aws.String(_workdocsSurname)
	}
	if len(_workdocsUsername) > 0 {
		input.Username = aws.String(_workdocsUsername)
	}
	if len(_workdocsAuthenticationToken) > 0 {
		input.AuthenticationToken = aws.String(_workdocsAuthenticationToken)
	}
	if len(_workdocsEmailAddress) > 0 {
		input.EmailAddress = aws.String(_workdocsEmailAddress)
	}
	if len(_workdocsOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workdocsOrganizationId)
	}
	if len(_workdocsStorageRule) > 0 {
		if err := assignInputField(input, "StorageRule", _workdocsStorageRule); err != nil {
			log.Errorf("invalid --storage-rule: %s", err.Error())
			return
		}
	}
	if len(_workdocsTimeZoneId) > 0 {
		input.TimeZoneId = aws.String(_workdocsTimeZoneId)
	}

	if resp, err := client.CreateUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deactivates the specified user, which revokes the user's access to Amazon
// WorkDocs.
func workdocs_DeactivateUser(cfg aws.Config, client *workdocs.Client) {
	input := &workdocs.DeactivateUserInput{
		// UserId: *string, // Required
	}

	if len(_workdocsUserId) > 0 {
		input.UserId = aws.String(_workdocsUserId)
	}
	if len(_workdocsAuthenticationToken) > 0 {
		input.AuthenticationToken = aws.String(_workdocsAuthenticationToken)
	}

	if resp, err := client.DeactivateUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified comment from the document version.
func workdocs_DeleteComment(cfg aws.Config, client *workdocs.Client) {
	input := &workdocs.DeleteCommentInput{
		// CommentId: *string, // Required
		// DocumentId: *string, // Required
		// VersionId: *string, // Required
	}

	if len(_workdocsCommentId) > 0 {
		input.CommentId = aws.String(_workdocsCommentId)
	}
	if len(_workdocsDocumentId) > 0 {
		input.DocumentId = aws.String(_workdocsDocumentId)
	}
	if len(_workdocsVersionId) > 0 {
		input.VersionId = aws.String(_workdocsVersionId)
	}
	if len(_workdocsAuthenticationToken) > 0 {
		input.AuthenticationToken = aws.String(_workdocsAuthenticationToken)
	}

	if resp, err := client.DeleteComment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes custom metadata from the specified resource.
func workdocs_DeleteCustomMetadata(cfg aws.Config, client *workdocs.Client) {
	input := &workdocs.DeleteCustomMetadataInput{
		// ResourceId: *string, // Required
	}

	if len(_workdocsResourceId) > 0 {
		input.ResourceId = aws.String(_workdocsResourceId)
	}
	if len(_workdocsAuthenticationToken) > 0 {
		input.AuthenticationToken = aws.String(_workdocsAuthenticationToken)
	}
	if len(_workdocsDeleteAll) > 0 {
		if err := assignInputField(input, "DeleteAll", _workdocsDeleteAll); err != nil {
			log.Errorf("invalid --delete-all: %s", err.Error())
			return
		}
	}
	if len(_workdocsKeys) > 0 {
		input.Keys = append([]string(nil), _workdocsKeys...)
	}
	if len(_workdocsVersionId) > 0 {
		input.VersionId = aws.String(_workdocsVersionId)
	}

	if resp, err := client.DeleteCustomMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Permanently deletes the specified document and its associated metadata.
func workdocs_DeleteDocument(cfg aws.Config, client *workdocs.Client) {
	input := &workdocs.DeleteDocumentInput{
		// DocumentId: *string, // Required
	}

	if len(_workdocsDocumentId) > 0 {
		input.DocumentId = aws.String(_workdocsDocumentId)
	}
	if len(_workdocsAuthenticationToken) > 0 {
		input.AuthenticationToken = aws.String(_workdocsAuthenticationToken)
	}

	if resp, err := client.DeleteDocument(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specific version of a document.
func workdocs_DeleteDocumentVersion(cfg aws.Config, client *workdocs.Client) {
	input := &workdocs.DeleteDocumentVersionInput{
		// DeletePriorVersions: bool, // Required
		// DocumentId: *string, // Required
		// VersionId: *string, // Required
	}

	if len(_workdocsDeletePriorVersions) > 0 {
		if err := assignInputField(input, "DeletePriorVersions", _workdocsDeletePriorVersions); err != nil {
			log.Errorf("invalid --delete-prior-versions: %s", err.Error())
			return
		}
	}
	if len(_workdocsDocumentId) > 0 {
		input.DocumentId = aws.String(_workdocsDocumentId)
	}
	if len(_workdocsVersionId) > 0 {
		input.VersionId = aws.String(_workdocsVersionId)
	}
	if len(_workdocsAuthenticationToken) > 0 {
		input.AuthenticationToken = aws.String(_workdocsAuthenticationToken)
	}

	if resp, err := client.DeleteDocumentVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Permanently deletes the specified folder and its contents.
func workdocs_DeleteFolder(cfg aws.Config, client *workdocs.Client) {
	input := &workdocs.DeleteFolderInput{
		// FolderId: *string, // Required
	}

	if len(_workdocsFolderId) > 0 {
		input.FolderId = aws.String(_workdocsFolderId)
	}
	if len(_workdocsAuthenticationToken) > 0 {
		input.AuthenticationToken = aws.String(_workdocsAuthenticationToken)
	}

	if resp, err := client.DeleteFolder(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the contents of the specified folder.
func workdocs_DeleteFolderContents(cfg aws.Config, client *workdocs.Client) {
	input := &workdocs.DeleteFolderContentsInput{
		// FolderId: *string, // Required
	}

	if len(_workdocsFolderId) > 0 {
		input.FolderId = aws.String(_workdocsFolderId)
	}
	if len(_workdocsAuthenticationToken) > 0 {
		input.AuthenticationToken = aws.String(_workdocsAuthenticationToken)
	}

	if resp, err := client.DeleteFolderContents(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified list of labels from a resource.
func workdocs_DeleteLabels(cfg aws.Config, client *workdocs.Client) {
	input := &workdocs.DeleteLabelsInput{
		// ResourceId: *string, // Required
	}

	if len(_workdocsResourceId) > 0 {
		input.ResourceId = aws.String(_workdocsResourceId)
	}
	if len(_workdocsAuthenticationToken) > 0 {
		input.AuthenticationToken = aws.String(_workdocsAuthenticationToken)
	}
	if len(_workdocsDeleteAll) > 0 {
		if err := assignInputField(input, "DeleteAll", _workdocsDeleteAll); err != nil {
			log.Errorf("invalid --delete-all: %s", err.Error())
			return
		}
	}
	if len(_workdocsLabels) > 0 {
		input.Labels = append([]string(nil), _workdocsLabels...)
	}

	if resp, err := client.DeleteLabels(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified subscription from the specified organization.
func workdocs_DeleteNotificationSubscription(cfg aws.Config, client *workdocs.Client) {
	input := &workdocs.DeleteNotificationSubscriptionInput{
		// OrganizationId: *string, // Required
		// SubscriptionId: *string, // Required
	}

	if len(_workdocsOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workdocsOrganizationId)
	}
	if len(_workdocsSubscriptionId) > 0 {
		input.SubscriptionId = aws.String(_workdocsSubscriptionId)
	}

	if resp, err := client.DeleteNotificationSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified user from a Simple AD or Microsoft AD directory.
// Deleting a user immediately and permanently deletes all content in that user's
// folder structure. Site retention policies do NOT apply to this type of deletion.
func workdocs_DeleteUser(cfg aws.Config, client *workdocs.Client) {
	input := &workdocs.DeleteUserInput{
		// UserId: *string, // Required
	}

	if len(_workdocsUserId) > 0 {
		input.UserId = aws.String(_workdocsUserId)
	}
	if len(_workdocsAuthenticationToken) > 0 {
		input.AuthenticationToken = aws.String(_workdocsAuthenticationToken)
	}

	if resp, err := client.DeleteUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the user activities in a specified time period.
func workdocs_DescribeActivities(cfg aws.Config, client *workdocs.Client) {
	input := &workdocs.DescribeActivitiesInput{}

	if len(_workdocsActivityTypes) > 0 {
		input.ActivityTypes = aws.String(_workdocsActivityTypes)
	}
	if len(_workdocsAuthenticationToken) > 0 {
		input.AuthenticationToken = aws.String(_workdocsAuthenticationToken)
	}
	if len(_workdocsEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _workdocsEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_workdocsIncludeIndirectActivities) > 0 {
		if err := assignInputField(input, "IncludeIndirectActivities", _workdocsIncludeIndirectActivities); err != nil {
			log.Errorf("invalid --include-indirect-activities: %s", err.Error())
			return
		}
	}
	if len(_workdocsLimit) > 0 {
		if err := assignInputField(input, "Limit", _workdocsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_workdocsMarker) > 0 {
		input.Marker = aws.String(_workdocsMarker)
	}
	if len(_workdocsOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workdocsOrganizationId)
	}
	if len(_workdocsResourceId) > 0 {
		input.ResourceId = aws.String(_workdocsResourceId)
	}
	if len(_workdocsStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _workdocsStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_workdocsUserId) > 0 {
		input.UserId = aws.String(_workdocsUserId)
	}

	if disablePaginator() {
		if resp, err := client.DescribeActivities(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*workdocs.DescribeActivitiesOutput
	p := workdocs.NewDescribeActivitiesPaginator(client, input)
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

// List all the comments for the specified document version.
func workdocs_DescribeComments(cfg aws.Config, client *workdocs.Client) {
	input := &workdocs.DescribeCommentsInput{
		// DocumentId: *string, // Required
		// VersionId: *string, // Required
	}

	if len(_workdocsDocumentId) > 0 {
		input.DocumentId = aws.String(_workdocsDocumentId)
	}
	if len(_workdocsVersionId) > 0 {
		input.VersionId = aws.String(_workdocsVersionId)
	}
	if len(_workdocsAuthenticationToken) > 0 {
		input.AuthenticationToken = aws.String(_workdocsAuthenticationToken)
	}
	if len(_workdocsLimit) > 0 {
		if err := assignInputField(input, "Limit", _workdocsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_workdocsMarker) > 0 {
		input.Marker = aws.String(_workdocsMarker)
	}

	if disablePaginator() {
		if resp, err := client.DescribeComments(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*workdocs.DescribeCommentsOutput
	p := workdocs.NewDescribeCommentsPaginator(client, input)
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

// Retrieves the document versions for the specified document.
// By default, only active versions are returned.
func workdocs_DescribeDocumentVersions(cfg aws.Config, client *workdocs.Client) {
	input := &workdocs.DescribeDocumentVersionsInput{
		// DocumentId: *string, // Required
	}

	if len(_workdocsDocumentId) > 0 {
		input.DocumentId = aws.String(_workdocsDocumentId)
	}
	if len(_workdocsAuthenticationToken) > 0 {
		input.AuthenticationToken = aws.String(_workdocsAuthenticationToken)
	}
	if len(_workdocsFields) > 0 {
		input.Fields = aws.String(_workdocsFields)
	}
	if len(_workdocsInclude) > 0 {
		input.Include = aws.String(_workdocsInclude)
	}
	if len(_workdocsLimit) > 0 {
		if err := assignInputField(input, "Limit", _workdocsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_workdocsMarker) > 0 {
		input.Marker = aws.String(_workdocsMarker)
	}

	if disablePaginator() {
		if resp, err := client.DescribeDocumentVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*workdocs.DescribeDocumentVersionsOutput
	p := workdocs.NewDescribeDocumentVersionsPaginator(client, input)
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

// Describes the contents of the specified folder, including its documents and
// subfolders.
//
// By default, Amazon WorkDocs returns the first 100 active document and folder
// metadata items. If there are more results, the response includes a marker that
// you can use to request the next set of results. You can also request initialized
// documents.
func workdocs_DescribeFolderContents(cfg aws.Config, client *workdocs.Client) {
	input := &workdocs.DescribeFolderContentsInput{
		// FolderId: *string, // Required
	}

	if len(_workdocsFolderId) > 0 {
		input.FolderId = aws.String(_workdocsFolderId)
	}
	if len(_workdocsAuthenticationToken) > 0 {
		input.AuthenticationToken = aws.String(_workdocsAuthenticationToken)
	}
	if len(_workdocsInclude) > 0 {
		input.Include = aws.String(_workdocsInclude)
	}
	if len(_workdocsLimit) > 0 {
		if err := assignInputField(input, "Limit", _workdocsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_workdocsMarker) > 0 {
		input.Marker = aws.String(_workdocsMarker)
	}
	if len(_workdocsOrder) > 0 {
		if err := assignInputField(input, "Order", _workdocsOrder); err != nil {
			log.Errorf("invalid --order: %s", err.Error())
			return
		}
	}
	if len(_workdocsSort) > 0 {
		if err := assignInputField(input, "Sort", _workdocsSort); err != nil {
			log.Errorf("invalid --sort: %s", err.Error())
			return
		}
	}
	if len(_workdocsType) > 0 {
		if err := assignInputField(input, "Type", _workdocsType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeFolderContents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*workdocs.DescribeFolderContentsOutput
	p := workdocs.NewDescribeFolderContentsPaginator(client, input)
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

// Describes the groups specified by the query. Groups are defined by the
// underlying Active Directory.
func workdocs_DescribeGroups(cfg aws.Config, client *workdocs.Client) {
	input := &workdocs.DescribeGroupsInput{
		// SearchQuery: *string, // Required
	}

	if len(_workdocsSearchQuery) > 0 {
		input.SearchQuery = aws.String(_workdocsSearchQuery)
	}
	if len(_workdocsAuthenticationToken) > 0 {
		input.AuthenticationToken = aws.String(_workdocsAuthenticationToken)
	}
	if len(_workdocsLimit) > 0 {
		if err := assignInputField(input, "Limit", _workdocsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_workdocsMarker) > 0 {
		input.Marker = aws.String(_workdocsMarker)
	}
	if len(_workdocsOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workdocsOrganizationId)
	}

	if disablePaginator() {
		if resp, err := client.DescribeGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*workdocs.DescribeGroupsOutput
	p := workdocs.NewDescribeGroupsPaginator(client, input)
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

// Lists the specified notification subscriptions.
func workdocs_DescribeNotificationSubscriptions(cfg aws.Config, client *workdocs.Client) {
	input := &workdocs.DescribeNotificationSubscriptionsInput{
		// OrganizationId: *string, // Required
	}

	if len(_workdocsOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workdocsOrganizationId)
	}
	if len(_workdocsLimit) > 0 {
		if err := assignInputField(input, "Limit", _workdocsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_workdocsMarker) > 0 {
		input.Marker = aws.String(_workdocsMarker)
	}

	if disablePaginator() {
		if resp, err := client.DescribeNotificationSubscriptions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*workdocs.DescribeNotificationSubscriptionsOutput
	p := workdocs.NewDescribeNotificationSubscriptionsPaginator(client, input)
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

// Describes the permissions of a specified resource.
func workdocs_DescribeResourcePermissions(cfg aws.Config, client *workdocs.Client) {
	input := &workdocs.DescribeResourcePermissionsInput{
		// ResourceId: *string, // Required
	}

	if len(_workdocsResourceId) > 0 {
		input.ResourceId = aws.String(_workdocsResourceId)
	}
	if len(_workdocsAuthenticationToken) > 0 {
		input.AuthenticationToken = aws.String(_workdocsAuthenticationToken)
	}
	if len(_workdocsLimit) > 0 {
		if err := assignInputField(input, "Limit", _workdocsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_workdocsMarker) > 0 {
		input.Marker = aws.String(_workdocsMarker)
	}
	if len(_workdocsPrincipalId) > 0 {
		input.PrincipalId = aws.String(_workdocsPrincipalId)
	}

	if disablePaginator() {
		if resp, err := client.DescribeResourcePermissions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*workdocs.DescribeResourcePermissionsOutput
	p := workdocs.NewDescribeResourcePermissionsPaginator(client, input)
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

// Describes the current user's special folders; the RootFolder and the RecycleBin
// . RootFolder is the root of user's files and folders and RecycleBin is the root
// of recycled items. This is not a valid action for SigV4 (administrative API)
// clients.
//
// This action requires an authentication token. To get an authentication token,
// register an application with Amazon WorkDocs. For more information, see [Authentication and Access Control for User Applications]in the
// Amazon WorkDocs Developer Guide.
//
// [Authentication and Access Control for User Applications]: https://docs.aws.amazon.com/workdocs/latest/developerguide/wd-auth-user.html
func workdocs_DescribeRootFolders(cfg aws.Config, client *workdocs.Client) {
	input := &workdocs.DescribeRootFoldersInput{
		// AuthenticationToken: *string, // Required
	}

	if len(_workdocsAuthenticationToken) > 0 {
		input.AuthenticationToken = aws.String(_workdocsAuthenticationToken)
	}
	if len(_workdocsLimit) > 0 {
		if err := assignInputField(input, "Limit", _workdocsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_workdocsMarker) > 0 {
		input.Marker = aws.String(_workdocsMarker)
	}

	if disablePaginator() {
		if resp, err := client.DescribeRootFolders(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*workdocs.DescribeRootFoldersOutput
	p := workdocs.NewDescribeRootFoldersPaginator(client, input)
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

// Describes the specified users. You can describe all users or filter the results
// (for example, by status or organization).
//
// By default, Amazon WorkDocs returns the first 24 active or pending users. If
// there are more results, the response includes a marker that you can use to
// request the next set of results.
func workdocs_DescribeUsers(cfg aws.Config, client *workdocs.Client) {
	input := &workdocs.DescribeUsersInput{}

	if len(_workdocsAuthenticationToken) > 0 {
		input.AuthenticationToken = aws.String(_workdocsAuthenticationToken)
	}
	if len(_workdocsFields) > 0 {
		input.Fields = aws.String(_workdocsFields)
	}
	if len(_workdocsInclude) > 0 {
		if err := assignInputField(input, "Include", _workdocsInclude); err != nil {
			log.Errorf("invalid --include: %s", err.Error())
			return
		}
	}
	if len(_workdocsLimit) > 0 {
		if err := assignInputField(input, "Limit", _workdocsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_workdocsMarker) > 0 {
		input.Marker = aws.String(_workdocsMarker)
	}
	if len(_workdocsOrder) > 0 {
		if err := assignInputField(input, "Order", _workdocsOrder); err != nil {
			log.Errorf("invalid --order: %s", err.Error())
			return
		}
	}
	if len(_workdocsOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workdocsOrganizationId)
	}
	if len(_workdocsQuery) > 0 {
		input.Query = aws.String(_workdocsQuery)
	}
	if len(_workdocsSort) > 0 {
		if err := assignInputField(input, "Sort", _workdocsSort); err != nil {
			log.Errorf("invalid --sort: %s", err.Error())
			return
		}
	}
	if len(_workdocsUserIds) > 0 {
		input.UserIds = aws.String(_workdocsUserIds)
	}

	if disablePaginator() {
		if resp, err := client.DescribeUsers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*workdocs.DescribeUsersOutput
	p := workdocs.NewDescribeUsersPaginator(client, input)
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

// Retrieves details of the current user for whom the authentication token was
// generated. This is not a valid action for SigV4 (administrative API) clients.
//
// This action requires an authentication token. To get an authentication token,
// register an application with Amazon WorkDocs. For more information, see [Authentication and Access Control for User Applications]in the
// Amazon WorkDocs Developer Guide.
//
// [Authentication and Access Control for User Applications]: https://docs.aws.amazon.com/workdocs/latest/developerguide/wd-auth-user.html
func workdocs_GetCurrentUser(cfg aws.Config, client *workdocs.Client) {
	input := &workdocs.GetCurrentUserInput{
		// AuthenticationToken: *string, // Required
	}

	if len(_workdocsAuthenticationToken) > 0 {
		input.AuthenticationToken = aws.String(_workdocsAuthenticationToken)
	}

	if resp, err := client.GetCurrentUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details of a document.
func workdocs_GetDocument(cfg aws.Config, client *workdocs.Client) {
	input := &workdocs.GetDocumentInput{
		// DocumentId: *string, // Required
	}

	if len(_workdocsDocumentId) > 0 {
		input.DocumentId = aws.String(_workdocsDocumentId)
	}
	if len(_workdocsAuthenticationToken) > 0 {
		input.AuthenticationToken = aws.String(_workdocsAuthenticationToken)
	}
	if len(_workdocsIncludeCustomMetadata) > 0 {
		if err := assignInputField(input, "IncludeCustomMetadata", _workdocsIncludeCustomMetadata); err != nil {
			log.Errorf("invalid --include-custom-metadata: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetDocument(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the path information (the hierarchy from the root folder) for the
// requested document.
//
// By default, Amazon WorkDocs returns a maximum of 100 levels upwards from the
// requested document and only includes the IDs of the parent folders in the path.
// You can limit the maximum number of levels. You can also request the names of
// the parent folders.
func workdocs_GetDocumentPath(cfg aws.Config, client *workdocs.Client) {
	input := &workdocs.GetDocumentPathInput{
		// DocumentId: *string, // Required
	}

	if len(_workdocsDocumentId) > 0 {
		input.DocumentId = aws.String(_workdocsDocumentId)
	}
	if len(_workdocsAuthenticationToken) > 0 {
		input.AuthenticationToken = aws.String(_workdocsAuthenticationToken)
	}
	if len(_workdocsFields) > 0 {
		input.Fields = aws.String(_workdocsFields)
	}
	if len(_workdocsLimit) > 0 {
		if err := assignInputField(input, "Limit", _workdocsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_workdocsMarker) > 0 {
		input.Marker = aws.String(_workdocsMarker)
	}

	if resp, err := client.GetDocumentPath(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves version metadata for the specified document.
func workdocs_GetDocumentVersion(cfg aws.Config, client *workdocs.Client) {
	input := &workdocs.GetDocumentVersionInput{
		// DocumentId: *string, // Required
		// VersionId: *string, // Required
	}

	if len(_workdocsDocumentId) > 0 {
		input.DocumentId = aws.String(_workdocsDocumentId)
	}
	if len(_workdocsVersionId) > 0 {
		input.VersionId = aws.String(_workdocsVersionId)
	}
	if len(_workdocsAuthenticationToken) > 0 {
		input.AuthenticationToken = aws.String(_workdocsAuthenticationToken)
	}
	if len(_workdocsFields) > 0 {
		input.Fields = aws.String(_workdocsFields)
	}
	if len(_workdocsIncludeCustomMetadata) > 0 {
		if err := assignInputField(input, "IncludeCustomMetadata", _workdocsIncludeCustomMetadata); err != nil {
			log.Errorf("invalid --include-custom-metadata: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetDocumentVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the metadata of the specified folder.
func workdocs_GetFolder(cfg aws.Config, client *workdocs.Client) {
	input := &workdocs.GetFolderInput{
		// FolderId: *string, // Required
	}

	if len(_workdocsFolderId) > 0 {
		input.FolderId = aws.String(_workdocsFolderId)
	}
	if len(_workdocsAuthenticationToken) > 0 {
		input.AuthenticationToken = aws.String(_workdocsAuthenticationToken)
	}
	if len(_workdocsIncludeCustomMetadata) > 0 {
		if err := assignInputField(input, "IncludeCustomMetadata", _workdocsIncludeCustomMetadata); err != nil {
			log.Errorf("invalid --include-custom-metadata: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetFolder(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the path information (the hierarchy from the root folder) for the
// specified folder.
//
// By default, Amazon WorkDocs returns a maximum of 100 levels upwards from the
// requested folder and only includes the IDs of the parent folders in the path.
// You can limit the maximum number of levels. You can also request the parent
// folder names.
func workdocs_GetFolderPath(cfg aws.Config, client *workdocs.Client) {
	input := &workdocs.GetFolderPathInput{
		// FolderId: *string, // Required
	}

	if len(_workdocsFolderId) > 0 {
		input.FolderId = aws.String(_workdocsFolderId)
	}
	if len(_workdocsAuthenticationToken) > 0 {
		input.AuthenticationToken = aws.String(_workdocsAuthenticationToken)
	}
	if len(_workdocsFields) > 0 {
		input.Fields = aws.String(_workdocsFields)
	}
	if len(_workdocsLimit) > 0 {
		if err := assignInputField(input, "Limit", _workdocsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_workdocsMarker) > 0 {
		input.Marker = aws.String(_workdocsMarker)
	}

	if resp, err := client.GetFolderPath(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a collection of resources, including folders and documents. The only
// CollectionType supported is SHARED_WITH_ME .
func workdocs_GetResources(cfg aws.Config, client *workdocs.Client) {
	input := &workdocs.GetResourcesInput{}

	if len(_workdocsAuthenticationToken) > 0 {
		input.AuthenticationToken = aws.String(_workdocsAuthenticationToken)
	}
	if len(_workdocsCollectionType) > 0 {
		if err := assignInputField(input, "CollectionType", _workdocsCollectionType); err != nil {
			log.Errorf("invalid --collection-type: %s", err.Error())
			return
		}
	}
	if len(_workdocsLimit) > 0 {
		if err := assignInputField(input, "Limit", _workdocsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_workdocsMarker) > 0 {
		input.Marker = aws.String(_workdocsMarker)
	}
	if len(_workdocsUserId) > 0 {
		input.UserId = aws.String(_workdocsUserId)
	}

	if resp, err := client.GetResources(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new document object and version object.
// The client specifies the parent folder ID and name of the document to upload.
// The ID is optionally specified when creating a new version of an existing
// document. This is the first step to upload a document. Next, upload the document
// to the URL returned from the call, and then call UpdateDocumentVersion.
//
// To cancel the document upload, call AbortDocumentVersionUpload.
func workdocs_InitiateDocumentVersionUpload(cfg aws.Config, client *workdocs.Client) {
	input := &workdocs.InitiateDocumentVersionUploadInput{}

	if len(_workdocsAuthenticationToken) > 0 {
		input.AuthenticationToken = aws.String(_workdocsAuthenticationToken)
	}
	if len(_workdocsContentCreatedTimestamp) > 0 {
		if err := assignInputField(input, "ContentCreatedTimestamp", _workdocsContentCreatedTimestamp); err != nil {
			log.Errorf("invalid --content-created-timestamp: %s", err.Error())
			return
		}
	}
	if len(_workdocsContentModifiedTimestamp) > 0 {
		if err := assignInputField(input, "ContentModifiedTimestamp", _workdocsContentModifiedTimestamp); err != nil {
			log.Errorf("invalid --content-modified-timestamp: %s", err.Error())
			return
		}
	}
	if len(_workdocsContentType) > 0 {
		input.ContentType = aws.String(_workdocsContentType)
	}
	if len(_workdocsDocumentSizeInBytes) > 0 {
		if err := assignInputField(input, "DocumentSizeInBytes", _workdocsDocumentSizeInBytes); err != nil {
			log.Errorf("invalid --document-size-in-bytes: %s", err.Error())
			return
		}
	}
	if len(_workdocsId) > 0 {
		input.Id = aws.String(_workdocsId)
	}
	if len(_workdocsName) > 0 {
		input.Name = aws.String(_workdocsName)
	}
	if len(_workdocsParentFolderId) > 0 {
		input.ParentFolderId = aws.String(_workdocsParentFolderId)
	}

	if resp, err := client.InitiateDocumentVersionUpload(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes all the permissions from the specified resource.
func workdocs_RemoveAllResourcePermissions(cfg aws.Config, client *workdocs.Client) {
	input := &workdocs.RemoveAllResourcePermissionsInput{
		// ResourceId: *string, // Required
	}

	if len(_workdocsResourceId) > 0 {
		input.ResourceId = aws.String(_workdocsResourceId)
	}
	if len(_workdocsAuthenticationToken) > 0 {
		input.AuthenticationToken = aws.String(_workdocsAuthenticationToken)
	}

	if resp, err := client.RemoveAllResourcePermissions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the permission for the specified principal from the specified resource.
func workdocs_RemoveResourcePermission(cfg aws.Config, client *workdocs.Client) {
	input := &workdocs.RemoveResourcePermissionInput{
		// PrincipalId: *string, // Required
		// ResourceId: *string, // Required
	}

	if len(_workdocsPrincipalId) > 0 {
		input.PrincipalId = aws.String(_workdocsPrincipalId)
	}
	if len(_workdocsResourceId) > 0 {
		input.ResourceId = aws.String(_workdocsResourceId)
	}
	if len(_workdocsAuthenticationToken) > 0 {
		input.AuthenticationToken = aws.String(_workdocsAuthenticationToken)
	}
	if len(_workdocsPrincipalType) > 0 {
		if err := assignInputField(input, "PrincipalType", _workdocsPrincipalType); err != nil {
			log.Errorf("invalid --principal-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.RemoveResourcePermission(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Recovers a deleted version of an Amazon WorkDocs document.
func workdocs_RestoreDocumentVersions(cfg aws.Config, client *workdocs.Client) {
	input := &workdocs.RestoreDocumentVersionsInput{
		// DocumentId: *string, // Required
	}

	if len(_workdocsDocumentId) > 0 {
		input.DocumentId = aws.String(_workdocsDocumentId)
	}
	if len(_workdocsAuthenticationToken) > 0 {
		input.AuthenticationToken = aws.String(_workdocsAuthenticationToken)
	}

	if resp, err := client.RestoreDocumentVersions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Searches metadata and the content of folders, documents, document versions, and
// comments.
func workdocs_SearchResources(cfg aws.Config, client *workdocs.Client) {
	input := &workdocs.SearchResourcesInput{}

	if len(_workdocsAdditionalResponseFields) > 0 {
		if err := assignInputField(input, "AdditionalResponseFields", _workdocsAdditionalResponseFields); err != nil {
			log.Errorf("invalid --additional-response-fields: %s", err.Error())
			return
		}
	}
	if len(_workdocsAuthenticationToken) > 0 {
		input.AuthenticationToken = aws.String(_workdocsAuthenticationToken)
	}
	if len(_workdocsFilters) > 0 {
		if err := assignInputField(input, "Filters", _workdocsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_workdocsLimit) > 0 {
		if err := assignInputField(input, "Limit", _workdocsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_workdocsMarker) > 0 {
		input.Marker = aws.String(_workdocsMarker)
	}
	if len(_workdocsOrderBy) > 0 {
		if err := assignInputField(input, "OrderBy", _workdocsOrderBy); err != nil {
			log.Errorf("invalid --order-by: %s", err.Error())
			return
		}
	}
	if len(_workdocsOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workdocsOrganizationId)
	}
	if len(_workdocsQueryScopes) > 0 {
		if err := assignInputField(input, "QueryScopes", _workdocsQueryScopes); err != nil {
			log.Errorf("invalid --query-scopes: %s", err.Error())
			return
		}
	}
	if len(_workdocsQueryText) > 0 {
		input.QueryText = aws.String(_workdocsQueryText)
	}

	if disablePaginator() {
		if resp, err := client.SearchResources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*workdocs.SearchResourcesOutput
	p := workdocs.NewSearchResourcesPaginator(client, input)
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

// Updates the specified attributes of a document. The user must have access to
// both the document and its parent folder, if applicable.
func workdocs_UpdateDocument(cfg aws.Config, client *workdocs.Client) {
	input := &workdocs.UpdateDocumentInput{
		// DocumentId: *string, // Required
	}

	if len(_workdocsDocumentId) > 0 {
		input.DocumentId = aws.String(_workdocsDocumentId)
	}
	if len(_workdocsAuthenticationToken) > 0 {
		input.AuthenticationToken = aws.String(_workdocsAuthenticationToken)
	}
	if len(_workdocsName) > 0 {
		input.Name = aws.String(_workdocsName)
	}
	if len(_workdocsParentFolderId) > 0 {
		input.ParentFolderId = aws.String(_workdocsParentFolderId)
	}
	if len(_workdocsResourceState) > 0 {
		if err := assignInputField(input, "ResourceState", _workdocsResourceState); err != nil {
			log.Errorf("invalid --resource-state: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDocument(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes the status of the document version to ACTIVE.
// Amazon WorkDocs also sets its document container to ACTIVE. This is the last
// step in a document upload, after the client uploads the document to an
// S3-presigned URL returned by InitiateDocumentVersionUpload.
func workdocs_UpdateDocumentVersion(cfg aws.Config, client *workdocs.Client) {
	input := &workdocs.UpdateDocumentVersionInput{
		// DocumentId: *string, // Required
		// VersionId: *string, // Required
	}

	if len(_workdocsDocumentId) > 0 {
		input.DocumentId = aws.String(_workdocsDocumentId)
	}
	if len(_workdocsVersionId) > 0 {
		input.VersionId = aws.String(_workdocsVersionId)
	}
	if len(_workdocsAuthenticationToken) > 0 {
		input.AuthenticationToken = aws.String(_workdocsAuthenticationToken)
	}
	if len(_workdocsVersionStatus) > 0 {
		if err := assignInputField(input, "VersionStatus", _workdocsVersionStatus); err != nil {
			log.Errorf("invalid --version-status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDocumentVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified attributes of the specified folder. The user must have
// access to both the folder and its parent folder, if applicable.
func workdocs_UpdateFolder(cfg aws.Config, client *workdocs.Client) {
	input := &workdocs.UpdateFolderInput{
		// FolderId: *string, // Required
	}

	if len(_workdocsFolderId) > 0 {
		input.FolderId = aws.String(_workdocsFolderId)
	}
	if len(_workdocsAuthenticationToken) > 0 {
		input.AuthenticationToken = aws.String(_workdocsAuthenticationToken)
	}
	if len(_workdocsName) > 0 {
		input.Name = aws.String(_workdocsName)
	}
	if len(_workdocsParentFolderId) > 0 {
		input.ParentFolderId = aws.String(_workdocsParentFolderId)
	}
	if len(_workdocsResourceState) > 0 {
		if err := assignInputField(input, "ResourceState", _workdocsResourceState); err != nil {
			log.Errorf("invalid --resource-state: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateFolder(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified attributes of the specified user, and grants or revokes
// administrative privileges to the Amazon WorkDocs site.
func workdocs_UpdateUser(cfg aws.Config, client *workdocs.Client) {
	input := &workdocs.UpdateUserInput{
		// UserId: *string, // Required
	}

	if len(_workdocsUserId) > 0 {
		input.UserId = aws.String(_workdocsUserId)
	}
	if len(_workdocsAuthenticationToken) > 0 {
		input.AuthenticationToken = aws.String(_workdocsAuthenticationToken)
	}
	if len(_workdocsGivenName) > 0 {
		input.GivenName = aws.String(_workdocsGivenName)
	}
	if len(_workdocsGrantPoweruserPrivileges) > 0 {
		if err := assignInputField(input, "GrantPoweruserPrivileges", _workdocsGrantPoweruserPrivileges); err != nil {
			log.Errorf("invalid --grant-poweruser-privileges: %s", err.Error())
			return
		}
	}
	if len(_workdocsLocale) > 0 {
		if err := assignInputField(input, "Locale", _workdocsLocale); err != nil {
			log.Errorf("invalid --locale: %s", err.Error())
			return
		}
	}
	if len(_workdocsStorageRule) > 0 {
		if err := assignInputField(input, "StorageRule", _workdocsStorageRule); err != nil {
			log.Errorf("invalid --storage-rule: %s", err.Error())
			return
		}
	}
	if len(_workdocsSurname) > 0 {
		input.Surname = aws.String(_workdocsSurname)
	}
	if len(_workdocsTimeZoneId) > 0 {
		input.TimeZoneId = aws.String(_workdocsTimeZoneId)
	}
	if len(_workdocsType) > 0 {
		if err := assignInputField(input, "Type", _workdocsType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_workdocsCmd)
	_workdocsCmd.Flags().SortFlags = false

	_workdocsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_workdocsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_workdocsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_workdocsCmd.Flags().StringVarP(&_workdocsActivityTypes, "activity-types", "", "", "Activity Types")
	_workdocsCmd.Flags().StringVarP(&_workdocsAdditionalResponseFields, "additional-response-fields", "", "", "Additional Response Fields")
	_workdocsCmd.Flags().StringVarP(&_workdocsAuthenticationToken, "authentication-token", "", "", "Authentication Token")
	_workdocsCmd.Flags().StringVarP(&_workdocsCollectionType, "collection-type", "", "", "Collection Type")
	_workdocsCmd.Flags().StringVarP(&_workdocsCommentId, "comment-id", "", "", "Comment ID")
	_workdocsCmd.Flags().StringVarP(&_workdocsContentCreatedTimestamp, "content-created-timestamp", "", "", "Content Created Timestamp")
	_workdocsCmd.Flags().StringVarP(&_workdocsContentModifiedTimestamp, "content-modified-timestamp", "", "", "Content Modified Timestamp")
	_workdocsCmd.Flags().StringVarP(&_workdocsContentType, "content-type", "", "", "Content Type")
	_workdocsCmd.Flags().StringVarP(&_workdocsCustomMetadata, "custom-metadata", "", "", "Custom Metadata")
	_workdocsCmd.Flags().StringVarP(&_workdocsDeleteAll, "delete-all", "", "", "Delete All")
	_workdocsCmd.Flags().StringVarP(&_workdocsDeletePriorVersions, "delete-prior-versions", "", "", "Delete Prior Versions")
	_workdocsCmd.Flags().StringVarP(&_workdocsDocumentId, "document-id", "", "", "Document ID")
	_workdocsCmd.Flags().StringVarP(&_workdocsDocumentSizeInBytes, "document-size-in-bytes", "", "", "Document Size In Bytes")
	_workdocsCmd.Flags().StringVarP(&_workdocsEmailAddress, "email-address", "", "", "Email Address")
	_workdocsCmd.Flags().StringVarP(&_workdocsEndTime, "end-time", "", "", "End Time")
	_workdocsCmd.Flags().StringVarP(&_workdocsEndpoint, "endpoint", "", "", "Endpoint")
	_workdocsCmd.Flags().StringVarP(&_workdocsFields, "fields", "", "", "Fields")
	_workdocsCmd.Flags().StringVarP(&_workdocsFilters, "filters", "", "", "Filters")
	_workdocsCmd.Flags().StringVarP(&_workdocsFolderId, "folder-id", "", "", "Folder ID")
	_workdocsCmd.Flags().StringVarP(&_workdocsGivenName, "given-name", "", "", "Given Name")
	_workdocsCmd.Flags().StringVarP(&_workdocsGrantPoweruserPrivileges, "grant-poweruser-privileges", "", "", "Grant Poweruser Privileges")
	_workdocsCmd.Flags().StringVarP(&_workdocsId, "id", "", "", "ID")
	_workdocsCmd.Flags().StringVarP(&_workdocsInclude, "include", "", "", "Include")
	_workdocsCmd.Flags().StringVarP(&_workdocsIncludeCustomMetadata, "include-custom-metadata", "", "", "Include Custom Metadata")
	_workdocsCmd.Flags().StringVarP(&_workdocsIncludeIndirectActivities, "include-indirect-activities", "", "", "Include Indirect Activities")
	_workdocsCmd.Flags().StringSliceVarP(&_workdocsKeys, "keys", "", nil, "Keys")
	_workdocsCmd.Flags().StringSliceVarP(&_workdocsLabels, "labels", "", nil, "Labels")
	_workdocsCmd.Flags().StringVarP(&_workdocsLimit, "limit", "", "", "Limit")
	_workdocsCmd.Flags().StringVarP(&_workdocsLocale, "locale", "", "", "Locale")
	_workdocsCmd.Flags().StringVarP(&_workdocsMarker, "marker", "", "", "Marker")
	_workdocsCmd.Flags().StringVarP(&_workdocsName, "name", "", "", "Name")
	_workdocsCmd.Flags().StringVarP(&_workdocsNotificationOptions, "notification-options", "", "", "Notification Options")
	_workdocsCmd.Flags().StringVarP(&_workdocsNotifyCollaborators, "notify-collaborators", "", "", "Notify Collaborators")
	_workdocsCmd.Flags().StringVarP(&_workdocsOrder, "order", "", "", "Order")
	_workdocsCmd.Flags().StringVarP(&_workdocsOrderBy, "order-by", "", "", "Order By")
	_workdocsCmd.Flags().StringVarP(&_workdocsOrganizationId, "organization-id", "", "", "Organization ID")
	_workdocsCmd.Flags().StringVarP(&_workdocsParentFolderId, "parent-folder-id", "", "", "Parent Folder ID")
	_workdocsCmd.Flags().StringVarP(&_workdocsParentId, "parent-id", "", "", "Parent ID")
	_workdocsCmd.Flags().StringVarP(&_workdocsPassword, "password", "", "", "Password")
	_workdocsCmd.Flags().StringVarP(&_workdocsPrincipalId, "principal-id", "", "", "Principal ID")
	_workdocsCmd.Flags().StringVarP(&_workdocsPrincipalType, "principal-type", "", "", "Principal Type")
	_workdocsCmd.Flags().StringVarP(&_workdocsPrincipals, "principals", "", "", "Principals")
	_workdocsCmd.Flags().StringVarP(&_workdocsProtocol, "protocol", "", "", "Protocol")
	_workdocsCmd.Flags().StringVarP(&_workdocsQuery, "query", "", "", "Query")
	_workdocsCmd.Flags().StringVarP(&_workdocsQueryScopes, "query-scopes", "", "", "Query Scopes")
	_workdocsCmd.Flags().StringVarP(&_workdocsQueryText, "query-text", "", "", "Query Text")
	_workdocsCmd.Flags().StringVarP(&_workdocsResourceId, "resource-id", "", "", "Resource ID")
	_workdocsCmd.Flags().StringVarP(&_workdocsResourceState, "resource-state", "", "", "Resource State")
	_workdocsCmd.Flags().StringVarP(&_workdocsSearchQuery, "search-query", "", "", "Search Query")
	_workdocsCmd.Flags().StringVarP(&_workdocsSort, "sort", "", "", "Sort")
	_workdocsCmd.Flags().StringVarP(&_workdocsStartTime, "start-time", "", "", "Start Time")
	_workdocsCmd.Flags().StringVarP(&_workdocsStorageRule, "storage-rule", "", "", "Storage Rule")
	_workdocsCmd.Flags().StringVarP(&_workdocsSubscriptionId, "subscription-id", "", "", "Subscription ID")
	_workdocsCmd.Flags().StringVarP(&_workdocsSubscriptionType, "subscription-type", "", "", "Subscription Type")
	_workdocsCmd.Flags().StringVarP(&_workdocsSurname, "surname", "", "", "Surname")
	_workdocsCmd.Flags().StringVarP(&_workdocsText, "text", "", "", "Text")
	_workdocsCmd.Flags().StringVarP(&_workdocsThreadId, "thread-id", "", "", "Thread ID")
	_workdocsCmd.Flags().StringVarP(&_workdocsTimeZoneId, "time-zone-id", "", "", "Time Zone ID")
	_workdocsCmd.Flags().StringVarP(&_workdocsType, "type", "", "", "Type")
	_workdocsCmd.Flags().StringVarP(&_workdocsUserId, "user-id", "", "", "User ID")
	_workdocsCmd.Flags().StringVarP(&_workdocsUserIds, "user-ids", "", "", "User Ids")
	_workdocsCmd.Flags().StringVarP(&_workdocsUsername, "username", "", "", "Username")
	_workdocsCmd.Flags().StringVarP(&_workdocsVersionId, "version-id", "", "", "Version ID")
	_workdocsCmd.Flags().StringVarP(&_workdocsVersionStatus, "version-status", "", "", "Version Status")
	_workdocsCmd.Flags().StringVarP(&_workdocsVisibility, "visibility", "", "", "Visibility")

	_workdocsCmd.Flags().BoolVarP(&_workdocsAbortDocumentVersionUpload, "abort-document-version-upload", "", false, "Abort Document Version Upload")
	_workdocsCmd.Flags().BoolVarP(&_workdocsActivateUser, "activate-user", "", false, "Activate User")
	_workdocsCmd.Flags().BoolVarP(&_workdocsAddResourcePermissions, "add-resource-permissions", "", false, "Add Resource Permissions")
	_workdocsCmd.Flags().BoolVarP(&_workdocsCreateComment, "create-comment", "", false, "Create Comment")
	_workdocsCmd.Flags().BoolVarP(&_workdocsCreateCustomMetadata, "create-custom-metadata", "", false, "Create Custom Metadata")
	_workdocsCmd.Flags().BoolVarP(&_workdocsCreateFolder, "create-folder", "", false, "Create Folder")
	_workdocsCmd.Flags().BoolVarP(&_workdocsCreateLabels, "create-labels", "", false, "Create Labels")
	_workdocsCmd.Flags().BoolVarP(&_workdocsCreateNotificationSubscription, "create-notification-subscription", "", false, "Create Notification Subscription")
	_workdocsCmd.Flags().BoolVarP(&_workdocsCreateUser, "create-user", "", false, "Create User")
	_workdocsCmd.Flags().BoolVarP(&_workdocsDeactivateUser, "deactivate-user", "", false, "Deactivate User")
	_workdocsCmd.Flags().BoolVarP(&_workdocsDeleteComment, "delete-comment", "", false, "Delete Comment")
	_workdocsCmd.Flags().BoolVarP(&_workdocsDeleteCustomMetadata, "delete-custom-metadata", "", false, "Delete Custom Metadata")
	_workdocsCmd.Flags().BoolVarP(&_workdocsDeleteDocument, "delete-document", "", false, "Delete Document")
	_workdocsCmd.Flags().BoolVarP(&_workdocsDeleteDocumentVersion, "delete-document-version", "", false, "Delete Document Version")
	_workdocsCmd.Flags().BoolVarP(&_workdocsDeleteFolder, "delete-folder", "", false, "Delete Folder")
	_workdocsCmd.Flags().BoolVarP(&_workdocsDeleteFolderContents, "delete-folder-contents", "", false, "Delete Folder Contents")
	_workdocsCmd.Flags().BoolVarP(&_workdocsDeleteLabels, "delete-labels", "", false, "Delete Labels")
	_workdocsCmd.Flags().BoolVarP(&_workdocsDeleteNotificationSubscription, "delete-notification-subscription", "", false, "Delete Notification Subscription")
	_workdocsCmd.Flags().BoolVarP(&_workdocsDeleteUser, "delete-user", "", false, "Delete User")
	_workdocsCmd.Flags().BoolVarP(&_workdocsDescribeActivities, "describe-activities", "", false, "Describe Activities")
	_workdocsCmd.Flags().BoolVarP(&_workdocsDescribeComments, "describe-comments", "", false, "Describe Comments")
	_workdocsCmd.Flags().BoolVarP(&_workdocsDescribeDocumentVersions, "describe-document-versions", "", false, "Describe Document Versions")
	_workdocsCmd.Flags().BoolVarP(&_workdocsDescribeFolderContents, "describe-folder-contents", "", false, "Describe Folder Contents")
	_workdocsCmd.Flags().BoolVarP(&_workdocsDescribeGroups, "describe-groups", "", false, "Describe Groups")
	_workdocsCmd.Flags().BoolVarP(&_workdocsDescribeNotificationSubscriptions, "describe-notification-subscriptions", "", false, "Describe Notification Subscriptions")
	_workdocsCmd.Flags().BoolVarP(&_workdocsDescribeResourcePermissions, "describe-resource-permissions", "", false, "Describe Resource Permissions")
	_workdocsCmd.Flags().BoolVarP(&_workdocsDescribeRootFolders, "describe-root-folders", "", false, "Describe Root Folders")
	_workdocsCmd.Flags().BoolVarP(&_workdocsDescribeUsers, "describe-users", "", false, "Describe Users")
	_workdocsCmd.Flags().BoolVarP(&_workdocsGetCurrentUser, "get-current-user", "", false, "Get Current User")
	_workdocsCmd.Flags().BoolVarP(&_workdocsGetDocument, "get-document", "", false, "Get Document")
	_workdocsCmd.Flags().BoolVarP(&_workdocsGetDocumentPath, "get-document-path", "", false, "Get Document Path")
	_workdocsCmd.Flags().BoolVarP(&_workdocsGetDocumentVersion, "get-document-version", "", false, "Get Document Version")
	_workdocsCmd.Flags().BoolVarP(&_workdocsGetFolder, "get-folder", "", false, "Get Folder")
	_workdocsCmd.Flags().BoolVarP(&_workdocsGetFolderPath, "get-folder-path", "", false, "Get Folder Path")
	_workdocsCmd.Flags().BoolVarP(&_workdocsGetResources, "get-resources", "", false, "Get Resources")
	_workdocsCmd.Flags().BoolVarP(&_workdocsInitiateDocumentVersionUpload, "initiate-document-version-upload", "", false, "Initiate Document Version Upload")
	_workdocsCmd.Flags().BoolVarP(&_workdocsRemoveAllResourcePermissions, "remove-all-resource-permissions", "", false, "Remove All Resource Permissions")
	_workdocsCmd.Flags().BoolVarP(&_workdocsRemoveResourcePermission, "remove-resource-permission", "", false, "Remove Resource Permission")
	_workdocsCmd.Flags().BoolVarP(&_workdocsRestoreDocumentVersions, "restore-document-versions", "", false, "Restore Document Versions")
	_workdocsCmd.Flags().BoolVarP(&_workdocsSearchResources, "search-resources", "", false, "Search Resources")
	_workdocsCmd.Flags().BoolVarP(&_workdocsUpdateDocument, "update-document", "", false, "Update Document")
	_workdocsCmd.Flags().BoolVarP(&_workdocsUpdateDocumentVersion, "update-document-version", "", false, "Update Document Version")
	_workdocsCmd.Flags().BoolVarP(&_workdocsUpdateFolder, "update-folder", "", false, "Update Folder")
	_workdocsCmd.Flags().BoolVarP(&_workdocsUpdateUser, "update-user", "", false, "Update User")

}
