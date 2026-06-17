package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/finspacedata"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// finspacedataCmd represents the finspacedata command
var _finspacedataCmd = &cobra.Command{
	Use:   "finspacedata",
	Short: "AWS finspacedata CLI",
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
		client := finspacedata.NewFromConfig(cfg)
		if _finspacedataAssociateUserToPermissionGroup {
			finspacedata_AssociateUserToPermissionGroup(cfg, client)
			return
		}
		if _finspacedataCreateChangeset {
			finspacedata_CreateChangeset(cfg, client)
			return
		}
		if _finspacedataCreateDataView {
			finspacedata_CreateDataView(cfg, client)
			return
		}
		if _finspacedataCreateDataset {
			finspacedata_CreateDataset(cfg, client)
			return
		}
		if _finspacedataCreatePermissionGroup {
			finspacedata_CreatePermissionGroup(cfg, client)
			return
		}
		if _finspacedataCreateUser {
			finspacedata_CreateUser(cfg, client)
			return
		}
		if _finspacedataDeleteDataset {
			finspacedata_DeleteDataset(cfg, client)
			return
		}
		if _finspacedataDeletePermissionGroup {
			finspacedata_DeletePermissionGroup(cfg, client)
			return
		}
		if _finspacedataDisableUser {
			finspacedata_DisableUser(cfg, client)
			return
		}
		if _finspacedataDisassociateUserFromPermissionGroup {
			finspacedata_DisassociateUserFromPermissionGroup(cfg, client)
			return
		}
		if _finspacedataEnableUser {
			finspacedata_EnableUser(cfg, client)
			return
		}
		if _finspacedataGetChangeset {
			finspacedata_GetChangeset(cfg, client)
			return
		}
		if _finspacedataGetDataView {
			finspacedata_GetDataView(cfg, client)
			return
		}
		if _finspacedataGetDataset {
			finspacedata_GetDataset(cfg, client)
			return
		}
		if _finspacedataGetExternalDataViewAccessDetails {
			finspacedata_GetExternalDataViewAccessDetails(cfg, client)
			return
		}
		if _finspacedataGetPermissionGroup {
			finspacedata_GetPermissionGroup(cfg, client)
			return
		}
		if _finspacedataGetProgrammaticAccessCredentials {
			finspacedata_GetProgrammaticAccessCredentials(cfg, client)
			return
		}
		if _finspacedataGetUser {
			finspacedata_GetUser(cfg, client)
			return
		}
		if _finspacedataGetWorkingLocation {
			finspacedata_GetWorkingLocation(cfg, client)
			return
		}
		if _finspacedataListChangesets {
			finspacedata_ListChangesets(cfg, client)
			return
		}
		if _finspacedataListDataViews {
			finspacedata_ListDataViews(cfg, client)
			return
		}
		if _finspacedataListDatasets {
			finspacedata_ListDatasets(cfg, client)
			return
		}
		if _finspacedataListPermissionGroups {
			finspacedata_ListPermissionGroups(cfg, client)
			return
		}
		if _finspacedataListPermissionGroupsByUser {
			finspacedata_ListPermissionGroupsByUser(cfg, client)
			return
		}
		if _finspacedataListUsers {
			finspacedata_ListUsers(cfg, client)
			return
		}
		if _finspacedataListUsersByPermissionGroup {
			finspacedata_ListUsersByPermissionGroup(cfg, client)
			return
		}
		if _finspacedataResetUserPassword {
			finspacedata_ResetUserPassword(cfg, client)
			return
		}
		if _finspacedataUpdateChangeset {
			finspacedata_UpdateChangeset(cfg, client)
			return
		}
		if _finspacedataUpdateDataset {
			finspacedata_UpdateDataset(cfg, client)
			return
		}
		if _finspacedataUpdatePermissionGroup {
			finspacedata_UpdatePermissionGroup(cfg, client)
			return
		}
		if _finspacedataUpdateUser {
			finspacedata_UpdateUser(cfg, client)
			return
		}

	},
}

var (
	_finspacedataAssociateUserToPermissionGroup      bool
	_finspacedataCreateChangeset                     bool
	_finspacedataCreateDataView                      bool
	_finspacedataCreateDataset                       bool
	_finspacedataCreatePermissionGroup               bool
	_finspacedataCreateUser                          bool
	_finspacedataDeleteDataset                       bool
	_finspacedataDeletePermissionGroup               bool
	_finspacedataDisableUser                         bool
	_finspacedataDisassociateUserFromPermissionGroup bool
	_finspacedataEnableUser                          bool
	_finspacedataGetChangeset                        bool
	_finspacedataGetDataView                         bool
	_finspacedataGetDataset                          bool
	_finspacedataGetExternalDataViewAccessDetails    bool
	_finspacedataGetPermissionGroup                  bool
	_finspacedataGetProgrammaticAccessCredentials    bool
	_finspacedataGetUser                             bool
	_finspacedataGetWorkingLocation                  bool
	_finspacedataListChangesets                      bool
	_finspacedataListDataViews                       bool
	_finspacedataListDatasets                        bool
	_finspacedataListPermissionGroups                bool
	_finspacedataListPermissionGroupsByUser          bool
	_finspacedataListUsers                           bool
	_finspacedataListUsersByPermissionGroup          bool
	_finspacedataResetUserPassword                   bool
	_finspacedataUpdateChangeset                     bool
	_finspacedataUpdateDataset                       bool
	_finspacedataUpdatePermissionGroup               bool
	_finspacedataUpdateUser                          bool

	_finspacedataAlias                  string
	_finspacedataApiAccess              string
	_finspacedataApiAccessPrincipalArn  string
	_finspacedataApplicationPermissions string
	_finspacedataAsOfTimestamp          string
	_finspacedataAutoUpdate             string
	_finspacedataChangeType             string
	_finspacedataChangesetId            string
	_finspacedataClientToken            string
	_finspacedataDataViewId             string
	_finspacedataDatasetDescription     string
	_finspacedataDatasetId              string
	_finspacedataDatasetTitle           string
	_finspacedataDescription            string
	_finspacedataDestinationTypeParams  string
	_finspacedataDurationInMinutes      string
	_finspacedataEmailAddress           string
	_finspacedataEnvironmentId          string
	_finspacedataFirstName              string
	_finspacedataFormatParams           string
	_finspacedataKind                   string
	_finspacedataLastName               string
	_finspacedataLocationType           string
	_finspacedataMaxResults             string
	_finspacedataName                   string
	_finspacedataNextToken              string
	_finspacedataOwnerInfo              string
	_finspacedataPartitionColumns       []string
	_finspacedataPermissionGroupId      string
	_finspacedataPermissionGroupParams  string
	_finspacedataSchemaDefinition       string
	_finspacedataSortColumns            []string
	_finspacedataSourceParams           string
	_finspacedataType                   string
	_finspacedataUserId                 string
)

// Adds a user to a permission group to grant permissions for actions a user can
// perform in FinSpace.
//
// Deprecated: This method will be discontinued.
func finspacedata_AssociateUserToPermissionGroup(cfg aws.Config, client *finspacedata.Client) {
	input := &finspacedata.AssociateUserToPermissionGroupInput{
		// PermissionGroupId: *string, // Required
		// UserId: *string, // Required
	}

	if len(_finspacedataPermissionGroupId) > 0 {
		input.PermissionGroupId = aws.String(_finspacedataPermissionGroupId)
	}
	if len(_finspacedataUserId) > 0 {
		input.UserId = aws.String(_finspacedataUserId)
	}
	if len(_finspacedataClientToken) > 0 {
		input.ClientToken = aws.String(_finspacedataClientToken)
	}

	if resp, err := client.AssociateUserToPermissionGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new Changeset in a FinSpace Dataset.
// Deprecated: This method will be discontinued.
func finspacedata_CreateChangeset(cfg aws.Config, client *finspacedata.Client) {
	input := &finspacedata.CreateChangesetInput{
		// ChangeType: types.ChangeType, // Required
		// DatasetId: *string, // Required
		// FormatParams: map[string]string, // Required
		// SourceParams: map[string]string, // Required
	}

	if len(_finspacedataChangeType) > 0 {
		if err := assignInputField(input, "ChangeType", _finspacedataChangeType); err != nil {
			log.Errorf("invalid --change-type: %s", err.Error())
			return
		}
	}
	if len(_finspacedataDatasetId) > 0 {
		input.DatasetId = aws.String(_finspacedataDatasetId)
	}
	if len(_finspacedataFormatParams) > 0 {
		if err := assignInputField(input, "FormatParams", _finspacedataFormatParams); err != nil {
			log.Errorf("invalid --format-params: %s", err.Error())
			return
		}
	}
	if len(_finspacedataSourceParams) > 0 {
		if err := assignInputField(input, "SourceParams", _finspacedataSourceParams); err != nil {
			log.Errorf("invalid --source-params: %s", err.Error())
			return
		}
	}
	if len(_finspacedataClientToken) > 0 {
		input.ClientToken = aws.String(_finspacedataClientToken)
	}

	if resp, err := client.CreateChangeset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a Dataview for a Dataset.
// Deprecated: This method will be discontinued.
func finspacedata_CreateDataView(cfg aws.Config, client *finspacedata.Client) {
	input := &finspacedata.CreateDataViewInput{
		// DatasetId: *string, // Required
		// DestinationTypeParams: *types.DataViewDestinationTypeParams, // Required
	}

	if len(_finspacedataDatasetId) > 0 {
		input.DatasetId = aws.String(_finspacedataDatasetId)
	}
	if len(_finspacedataDestinationTypeParams) > 0 {
		if err := assignInputField(input, "DestinationTypeParams", _finspacedataDestinationTypeParams); err != nil {
			log.Errorf("invalid --destination-type-params: %s", err.Error())
			return
		}
	}
	if len(_finspacedataAsOfTimestamp) > 0 {
		if err := assignInputField(input, "AsOfTimestamp", _finspacedataAsOfTimestamp); err != nil {
			log.Errorf("invalid --as-of-timestamp: %s", err.Error())
			return
		}
	}
	if len(_finspacedataAutoUpdate) > 0 {
		if err := assignInputField(input, "AutoUpdate", _finspacedataAutoUpdate); err != nil {
			log.Errorf("invalid --auto-update: %s", err.Error())
			return
		}
	}
	if len(_finspacedataClientToken) > 0 {
		input.ClientToken = aws.String(_finspacedataClientToken)
	}
	if len(_finspacedataPartitionColumns) > 0 {
		input.PartitionColumns = append([]string(nil), _finspacedataPartitionColumns...)
	}
	if len(_finspacedataSortColumns) > 0 {
		input.SortColumns = append([]string(nil), _finspacedataSortColumns...)
	}

	if resp, err := client.CreateDataView(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new FinSpace Dataset.
// Deprecated: This method will be discontinued.
func finspacedata_CreateDataset(cfg aws.Config, client *finspacedata.Client) {
	input := &finspacedata.CreateDatasetInput{
		// DatasetTitle: *string, // Required
		// Kind: types.DatasetKind, // Required
		// PermissionGroupParams: *types.PermissionGroupParams, // Required
	}

	if len(_finspacedataDatasetTitle) > 0 {
		input.DatasetTitle = aws.String(_finspacedataDatasetTitle)
	}
	if len(_finspacedataKind) > 0 {
		if err := assignInputField(input, "Kind", _finspacedataKind); err != nil {
			log.Errorf("invalid --kind: %s", err.Error())
			return
		}
	}
	if len(_finspacedataPermissionGroupParams) > 0 {
		if err := assignInputField(input, "PermissionGroupParams", _finspacedataPermissionGroupParams); err != nil {
			log.Errorf("invalid --permission-group-params: %s", err.Error())
			return
		}
	}
	if len(_finspacedataAlias) > 0 {
		input.Alias = aws.String(_finspacedataAlias)
	}
	if len(_finspacedataClientToken) > 0 {
		input.ClientToken = aws.String(_finspacedataClientToken)
	}
	if len(_finspacedataDatasetDescription) > 0 {
		input.DatasetDescription = aws.String(_finspacedataDatasetDescription)
	}
	if len(_finspacedataOwnerInfo) > 0 {
		if err := assignInputField(input, "OwnerInfo", _finspacedataOwnerInfo); err != nil {
			log.Errorf("invalid --owner-info: %s", err.Error())
			return
		}
	}
	if len(_finspacedataSchemaDefinition) > 0 {
		if err := assignInputField(input, "SchemaDefinition", _finspacedataSchemaDefinition); err != nil {
			log.Errorf("invalid --schema-definition: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDataset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a group of permissions for various actions that a user can perform in
// FinSpace.
//
// Deprecated: This method will be discontinued.
func finspacedata_CreatePermissionGroup(cfg aws.Config, client *finspacedata.Client) {
	input := &finspacedata.CreatePermissionGroupInput{
		// ApplicationPermissions: []types.ApplicationPermission, // Required
		// Name: *string, // Required
	}

	if len(_finspacedataApplicationPermissions) > 0 {
		if err := assignInputField(input, "ApplicationPermissions", _finspacedataApplicationPermissions); err != nil {
			log.Errorf("invalid --application-permissions: %s", err.Error())
			return
		}
	}
	if len(_finspacedataName) > 0 {
		input.Name = aws.String(_finspacedataName)
	}
	if len(_finspacedataClientToken) > 0 {
		input.ClientToken = aws.String(_finspacedataClientToken)
	}
	if len(_finspacedataDescription) > 0 {
		input.Description = aws.String(_finspacedataDescription)
	}

	if resp, err := client.CreatePermissionGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new user in FinSpace.
// Deprecated: This method will be discontinued.
func finspacedata_CreateUser(cfg aws.Config, client *finspacedata.Client) {
	input := &finspacedata.CreateUserInput{
		// EmailAddress: *string, // Required
		// Type: types.UserType, // Required
	}

	if len(_finspacedataEmailAddress) > 0 {
		input.EmailAddress = aws.String(_finspacedataEmailAddress)
	}
	if len(_finspacedataType) > 0 {
		if err := assignInputField(input, "Type", _finspacedataType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_finspacedataApiAccess) > 0 {
		if err := assignInputField(input, "ApiAccess", _finspacedataApiAccess); err != nil {
			log.Errorf("invalid --api-access: %s", err.Error())
			return
		}
	}
	if len(_finspacedataApiAccessPrincipalArn) > 0 {
		input.ApiAccessPrincipalArn = aws.String(_finspacedataApiAccessPrincipalArn)
	}
	if len(_finspacedataClientToken) > 0 {
		input.ClientToken = aws.String(_finspacedataClientToken)
	}
	if len(_finspacedataFirstName) > 0 {
		input.FirstName = aws.String(_finspacedataFirstName)
	}
	if len(_finspacedataLastName) > 0 {
		input.LastName = aws.String(_finspacedataLastName)
	}

	if resp, err := client.CreateUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a FinSpace Dataset.
// Deprecated: This method will be discontinued.
func finspacedata_DeleteDataset(cfg aws.Config, client *finspacedata.Client) {
	input := &finspacedata.DeleteDatasetInput{
		// DatasetId: *string, // Required
	}

	if len(_finspacedataDatasetId) > 0 {
		input.DatasetId = aws.String(_finspacedataDatasetId)
	}
	if len(_finspacedataClientToken) > 0 {
		input.ClientToken = aws.String(_finspacedataClientToken)
	}

	if resp, err := client.DeleteDataset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a permission group. This action is irreversible.
// Deprecated: This method will be discontinued.
func finspacedata_DeletePermissionGroup(cfg aws.Config, client *finspacedata.Client) {
	input := &finspacedata.DeletePermissionGroupInput{
		// PermissionGroupId: *string, // Required
	}

	if len(_finspacedataPermissionGroupId) > 0 {
		input.PermissionGroupId = aws.String(_finspacedataPermissionGroupId)
	}
	if len(_finspacedataClientToken) > 0 {
		input.ClientToken = aws.String(_finspacedataClientToken)
	}

	if resp, err := client.DeletePermissionGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Denies access to the FinSpace web application and API for the specified user.
// Deprecated: This method will be discontinued.
func finspacedata_DisableUser(cfg aws.Config, client *finspacedata.Client) {
	input := &finspacedata.DisableUserInput{
		// UserId: *string, // Required
	}

	if len(_finspacedataUserId) > 0 {
		input.UserId = aws.String(_finspacedataUserId)
	}
	if len(_finspacedataClientToken) > 0 {
		input.ClientToken = aws.String(_finspacedataClientToken)
	}

	if resp, err := client.DisableUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a user from a permission group.
// Deprecated: This method will be discontinued.
func finspacedata_DisassociateUserFromPermissionGroup(cfg aws.Config, client *finspacedata.Client) {
	input := &finspacedata.DisassociateUserFromPermissionGroupInput{
		// PermissionGroupId: *string, // Required
		// UserId: *string, // Required
	}

	if len(_finspacedataPermissionGroupId) > 0 {
		input.PermissionGroupId = aws.String(_finspacedataPermissionGroupId)
	}
	if len(_finspacedataUserId) > 0 {
		input.UserId = aws.String(_finspacedataUserId)
	}
	if len(_finspacedataClientToken) > 0 {
		input.ClientToken = aws.String(_finspacedataClientToken)
	}

	if resp, err := client.DisassociateUserFromPermissionGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Allows the specified user to access the FinSpace web application and API.
// Deprecated: This method will be discontinued.
func finspacedata_EnableUser(cfg aws.Config, client *finspacedata.Client) {
	input := &finspacedata.EnableUserInput{
		// UserId: *string, // Required
	}

	if len(_finspacedataUserId) > 0 {
		input.UserId = aws.String(_finspacedataUserId)
	}
	if len(_finspacedataClientToken) > 0 {
		input.ClientToken = aws.String(_finspacedataClientToken)
	}

	if resp, err := client.EnableUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get information about a Changeset.
// Deprecated: This method will be discontinued.
func finspacedata_GetChangeset(cfg aws.Config, client *finspacedata.Client) {
	input := &finspacedata.GetChangesetInput{
		// ChangesetId: *string, // Required
		// DatasetId: *string, // Required
	}

	if len(_finspacedataChangesetId) > 0 {
		input.ChangesetId = aws.String(_finspacedataChangesetId)
	}
	if len(_finspacedataDatasetId) > 0 {
		input.DatasetId = aws.String(_finspacedataDatasetId)
	}

	if resp, err := client.GetChangeset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a Dataview.
// Deprecated: This method will be discontinued.
func finspacedata_GetDataView(cfg aws.Config, client *finspacedata.Client) {
	input := &finspacedata.GetDataViewInput{
		// DataViewId: *string, // Required
		// DatasetId: *string, // Required
	}

	if len(_finspacedataDataViewId) > 0 {
		input.DataViewId = aws.String(_finspacedataDataViewId)
	}
	if len(_finspacedataDatasetId) > 0 {
		input.DatasetId = aws.String(_finspacedataDatasetId)
	}

	if resp, err := client.GetDataView(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a Dataset.
// Deprecated: This method will be discontinued.
func finspacedata_GetDataset(cfg aws.Config, client *finspacedata.Client) {
	input := &finspacedata.GetDatasetInput{
		// DatasetId: *string, // Required
	}

	if len(_finspacedataDatasetId) > 0 {
		input.DatasetId = aws.String(_finspacedataDatasetId)
	}

	if resp, err := client.GetDataset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the credentials to access the external Dataview from an S3 location. To
// call this API:
//
// - You must retrieve the programmatic credentials.
//
// - You must be a member of a FinSpace user group, where the dataset that you
// want to access has Read Dataset Data permissions.
//
// Deprecated: This method will be discontinued.
func finspacedata_GetExternalDataViewAccessDetails(cfg aws.Config, client *finspacedata.Client) {
	input := &finspacedata.GetExternalDataViewAccessDetailsInput{
		// DataViewId: *string, // Required
		// DatasetId: *string, // Required
	}

	if len(_finspacedataDataViewId) > 0 {
		input.DataViewId = aws.String(_finspacedataDataViewId)
	}
	if len(_finspacedataDatasetId) > 0 {
		input.DatasetId = aws.String(_finspacedataDatasetId)
	}

	if resp, err := client.GetExternalDataViewAccessDetails(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the details of a specific permission group.
// Deprecated: This method will be discontinued.
func finspacedata_GetPermissionGroup(cfg aws.Config, client *finspacedata.Client) {
	input := &finspacedata.GetPermissionGroupInput{
		// PermissionGroupId: *string, // Required
	}

	if len(_finspacedataPermissionGroupId) > 0 {
		input.PermissionGroupId = aws.String(_finspacedataPermissionGroupId)
	}

	if resp, err := client.GetPermissionGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Request programmatic credentials to use with FinSpace SDK. For more
// information, see [Step 2. Access credentials programmatically using IAM access key id and secret access key].
//
// Deprecated: This method will be discontinued.
//
// [Step 2. Access credentials programmatically using IAM access key id and secret access key]: https://docs.aws.amazon.com/finspace/latest/data-api/fs-using-the-finspace-api.html#accessing-credentials
func finspacedata_GetProgrammaticAccessCredentials(cfg aws.Config, client *finspacedata.Client) {
	input := &finspacedata.GetProgrammaticAccessCredentialsInput{
		// EnvironmentId: *string, // Required
	}

	if len(_finspacedataEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_finspacedataEnvironmentId)
	}
	if len(_finspacedataDurationInMinutes) > 0 {
		if err := assignInputField(input, "DurationInMinutes", _finspacedataDurationInMinutes); err != nil {
			log.Errorf("invalid --duration-in-minutes: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetProgrammaticAccessCredentials(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details for a specific user.
// Deprecated: This method will be discontinued.
func finspacedata_GetUser(cfg aws.Config, client *finspacedata.Client) {
	input := &finspacedata.GetUserInput{
		// UserId: *string, // Required
	}

	if len(_finspacedataUserId) > 0 {
		input.UserId = aws.String(_finspacedataUserId)
	}

	if resp, err := client.GetUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// A temporary Amazon S3 location, where you can copy your files from a source
// location to stage or use as a scratch space in FinSpace notebook.
//
// Deprecated: This method will be discontinued.
func finspacedata_GetWorkingLocation(cfg aws.Config, client *finspacedata.Client) {
	input := &finspacedata.GetWorkingLocationInput{}

	if len(_finspacedataLocationType) > 0 {
		if err := assignInputField(input, "LocationType", _finspacedataLocationType); err != nil {
			log.Errorf("invalid --location-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetWorkingLocation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the FinSpace Changesets for a Dataset.
// Deprecated: This method will be discontinued.
func finspacedata_ListChangesets(cfg aws.Config, client *finspacedata.Client) {
	input := &finspacedata.ListChangesetsInput{
		// DatasetId: *string, // Required
	}

	if len(_finspacedataDatasetId) > 0 {
		input.DatasetId = aws.String(_finspacedataDatasetId)
	}
	if len(_finspacedataMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _finspacedataMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_finspacedataNextToken) > 0 {
		input.NextToken = aws.String(_finspacedataNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListChangesets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*finspacedata.ListChangesetsOutput
	p := finspacedata.NewListChangesetsPaginator(client, input)
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

// Lists all available Dataviews for a Dataset.
// Deprecated: This method will be discontinued.
func finspacedata_ListDataViews(cfg aws.Config, client *finspacedata.Client) {
	input := &finspacedata.ListDataViewsInput{
		// DatasetId: *string, // Required
	}

	if len(_finspacedataDatasetId) > 0 {
		input.DatasetId = aws.String(_finspacedataDatasetId)
	}
	if len(_finspacedataMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _finspacedataMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_finspacedataNextToken) > 0 {
		input.NextToken = aws.String(_finspacedataNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDataViews(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*finspacedata.ListDataViewsOutput
	p := finspacedata.NewListDataViewsPaginator(client, input)
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

// Lists all of the active Datasets that a user has access to.
// Deprecated: This method will be discontinued.
func finspacedata_ListDatasets(cfg aws.Config, client *finspacedata.Client) {
	input := &finspacedata.ListDatasetsInput{}

	if len(_finspacedataMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _finspacedataMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_finspacedataNextToken) > 0 {
		input.NextToken = aws.String(_finspacedataNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDatasets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*finspacedata.ListDatasetsOutput
	p := finspacedata.NewListDatasetsPaginator(client, input)
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

// Lists all available permission groups in FinSpace.
// Deprecated: This method will be discontinued.
func finspacedata_ListPermissionGroups(cfg aws.Config, client *finspacedata.Client) {
	input := &finspacedata.ListPermissionGroupsInput{
		// MaxResults: *int32, // Required
	}

	if len(_finspacedataMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _finspacedataMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_finspacedataNextToken) > 0 {
		input.NextToken = aws.String(_finspacedataNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPermissionGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*finspacedata.ListPermissionGroupsOutput
	p := finspacedata.NewListPermissionGroupsPaginator(client, input)
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

// Lists all the permission groups that are associated with a specific user.
// Deprecated: This method will be discontinued.
func finspacedata_ListPermissionGroupsByUser(cfg aws.Config, client *finspacedata.Client) {
	input := &finspacedata.ListPermissionGroupsByUserInput{
		// MaxResults: *int32, // Required
		// UserId: *string, // Required
	}

	if len(_finspacedataMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _finspacedataMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_finspacedataUserId) > 0 {
		input.UserId = aws.String(_finspacedataUserId)
	}
	if len(_finspacedataNextToken) > 0 {
		input.NextToken = aws.String(_finspacedataNextToken)
	}

	if resp, err := client.ListPermissionGroupsByUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all available users in FinSpace.
// Deprecated: This method will be discontinued.
func finspacedata_ListUsers(cfg aws.Config, client *finspacedata.Client) {
	input := &finspacedata.ListUsersInput{
		// MaxResults: *int32, // Required
	}

	if len(_finspacedataMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _finspacedataMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_finspacedataNextToken) > 0 {
		input.NextToken = aws.String(_finspacedataNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListUsers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*finspacedata.ListUsersOutput
	p := finspacedata.NewListUsersPaginator(client, input)
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

// Lists details of all the users in a specific permission group.
// Deprecated: This method will be discontinued.
func finspacedata_ListUsersByPermissionGroup(cfg aws.Config, client *finspacedata.Client) {
	input := &finspacedata.ListUsersByPermissionGroupInput{
		// MaxResults: *int32, // Required
		// PermissionGroupId: *string, // Required
	}

	if len(_finspacedataMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _finspacedataMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_finspacedataPermissionGroupId) > 0 {
		input.PermissionGroupId = aws.String(_finspacedataPermissionGroupId)
	}
	if len(_finspacedataNextToken) > 0 {
		input.NextToken = aws.String(_finspacedataNextToken)
	}

	if resp, err := client.ListUsersByPermissionGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Resets the password for a specified user ID and generates a temporary one. Only
// a superuser can reset password for other users. Resetting the password
// immediately invalidates the previous password associated with the user.
//
// Deprecated: This method will be discontinued.
func finspacedata_ResetUserPassword(cfg aws.Config, client *finspacedata.Client) {
	input := &finspacedata.ResetUserPasswordInput{
		// UserId: *string, // Required
	}

	if len(_finspacedataUserId) > 0 {
		input.UserId = aws.String(_finspacedataUserId)
	}
	if len(_finspacedataClientToken) > 0 {
		input.ClientToken = aws.String(_finspacedataClientToken)
	}

	if resp, err := client.ResetUserPassword(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a FinSpace Changeset.
// Deprecated: This method will be discontinued.
func finspacedata_UpdateChangeset(cfg aws.Config, client *finspacedata.Client) {
	input := &finspacedata.UpdateChangesetInput{
		// ChangesetId: *string, // Required
		// DatasetId: *string, // Required
		// FormatParams: map[string]string, // Required
		// SourceParams: map[string]string, // Required
	}

	if len(_finspacedataChangesetId) > 0 {
		input.ChangesetId = aws.String(_finspacedataChangesetId)
	}
	if len(_finspacedataDatasetId) > 0 {
		input.DatasetId = aws.String(_finspacedataDatasetId)
	}
	if len(_finspacedataFormatParams) > 0 {
		if err := assignInputField(input, "FormatParams", _finspacedataFormatParams); err != nil {
			log.Errorf("invalid --format-params: %s", err.Error())
			return
		}
	}
	if len(_finspacedataSourceParams) > 0 {
		if err := assignInputField(input, "SourceParams", _finspacedataSourceParams); err != nil {
			log.Errorf("invalid --source-params: %s", err.Error())
			return
		}
	}
	if len(_finspacedataClientToken) > 0 {
		input.ClientToken = aws.String(_finspacedataClientToken)
	}

	if resp, err := client.UpdateChangeset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a FinSpace Dataset.
// Deprecated: This method will be discontinued.
func finspacedata_UpdateDataset(cfg aws.Config, client *finspacedata.Client) {
	input := &finspacedata.UpdateDatasetInput{
		// DatasetId: *string, // Required
		// DatasetTitle: *string, // Required
		// Kind: types.DatasetKind, // Required
	}

	if len(_finspacedataDatasetId) > 0 {
		input.DatasetId = aws.String(_finspacedataDatasetId)
	}
	if len(_finspacedataDatasetTitle) > 0 {
		input.DatasetTitle = aws.String(_finspacedataDatasetTitle)
	}
	if len(_finspacedataKind) > 0 {
		if err := assignInputField(input, "Kind", _finspacedataKind); err != nil {
			log.Errorf("invalid --kind: %s", err.Error())
			return
		}
	}
	if len(_finspacedataAlias) > 0 {
		input.Alias = aws.String(_finspacedataAlias)
	}
	if len(_finspacedataClientToken) > 0 {
		input.ClientToken = aws.String(_finspacedataClientToken)
	}
	if len(_finspacedataDatasetDescription) > 0 {
		input.DatasetDescription = aws.String(_finspacedataDatasetDescription)
	}
	if len(_finspacedataSchemaDefinition) > 0 {
		if err := assignInputField(input, "SchemaDefinition", _finspacedataSchemaDefinition); err != nil {
			log.Errorf("invalid --schema-definition: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDataset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the details of a permission group. You cannot modify a
// permissionGroupID .
//
// Deprecated: This method will be discontinued.
func finspacedata_UpdatePermissionGroup(cfg aws.Config, client *finspacedata.Client) {
	input := &finspacedata.UpdatePermissionGroupInput{
		// PermissionGroupId: *string, // Required
	}

	if len(_finspacedataPermissionGroupId) > 0 {
		input.PermissionGroupId = aws.String(_finspacedataPermissionGroupId)
	}
	if len(_finspacedataApplicationPermissions) > 0 {
		if err := assignInputField(input, "ApplicationPermissions", _finspacedataApplicationPermissions); err != nil {
			log.Errorf("invalid --application-permissions: %s", err.Error())
			return
		}
	}
	if len(_finspacedataClientToken) > 0 {
		input.ClientToken = aws.String(_finspacedataClientToken)
	}
	if len(_finspacedataDescription) > 0 {
		input.Description = aws.String(_finspacedataDescription)
	}
	if len(_finspacedataName) > 0 {
		input.Name = aws.String(_finspacedataName)
	}

	if resp, err := client.UpdatePermissionGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the details of the specified user. You cannot update the userId for a
// user.
//
// Deprecated: This method will be discontinued.
func finspacedata_UpdateUser(cfg aws.Config, client *finspacedata.Client) {
	input := &finspacedata.UpdateUserInput{
		// UserId: *string, // Required
	}

	if len(_finspacedataUserId) > 0 {
		input.UserId = aws.String(_finspacedataUserId)
	}
	if len(_finspacedataApiAccess) > 0 {
		if err := assignInputField(input, "ApiAccess", _finspacedataApiAccess); err != nil {
			log.Errorf("invalid --api-access: %s", err.Error())
			return
		}
	}
	if len(_finspacedataApiAccessPrincipalArn) > 0 {
		input.ApiAccessPrincipalArn = aws.String(_finspacedataApiAccessPrincipalArn)
	}
	if len(_finspacedataClientToken) > 0 {
		input.ClientToken = aws.String(_finspacedataClientToken)
	}
	if len(_finspacedataFirstName) > 0 {
		input.FirstName = aws.String(_finspacedataFirstName)
	}
	if len(_finspacedataLastName) > 0 {
		input.LastName = aws.String(_finspacedataLastName)
	}
	if len(_finspacedataType) > 0 {
		if err := assignInputField(input, "Type", _finspacedataType); err != nil {
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
	_rootCmd.AddCommand(_finspacedataCmd)
	_finspacedataCmd.Flags().SortFlags = false

	_finspacedataCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_finspacedataCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_finspacedataCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_finspacedataCmd.Flags().StringVarP(&_finspacedataAlias, "alias", "", "", "Alias")
	_finspacedataCmd.Flags().StringVarP(&_finspacedataApiAccess, "api-access", "", "", "API Access")
	_finspacedataCmd.Flags().StringVarP(&_finspacedataApiAccessPrincipalArn, "api-access-principal-arn", "", "", "API Access Principal ARN")
	_finspacedataCmd.Flags().StringVarP(&_finspacedataApplicationPermissions, "application-permissions", "", "", "Application Permissions")
	_finspacedataCmd.Flags().StringVarP(&_finspacedataAsOfTimestamp, "as-of-timestamp", "", "", "As Of Timestamp")
	_finspacedataCmd.Flags().StringVarP(&_finspacedataAutoUpdate, "auto-update", "", "", "Auto Update")
	_finspacedataCmd.Flags().StringVarP(&_finspacedataChangeType, "change-type", "", "", "Change Type")
	_finspacedataCmd.Flags().StringVarP(&_finspacedataChangesetId, "changeset-id", "", "", "Changeset ID")
	_finspacedataCmd.Flags().StringVarP(&_finspacedataClientToken, "client-token", "", "", "Client Token")
	_finspacedataCmd.Flags().StringVarP(&_finspacedataDataViewId, "data-view-id", "", "", "Data View ID")
	_finspacedataCmd.Flags().StringVarP(&_finspacedataDatasetDescription, "dataset-description", "", "", "Dataset Description")
	_finspacedataCmd.Flags().StringVarP(&_finspacedataDatasetId, "dataset-id", "", "", "Dataset ID")
	_finspacedataCmd.Flags().StringVarP(&_finspacedataDatasetTitle, "dataset-title", "", "", "Dataset Title")
	_finspacedataCmd.Flags().StringVarP(&_finspacedataDescription, "description", "", "", "Description")
	_finspacedataCmd.Flags().StringVarP(&_finspacedataDestinationTypeParams, "destination-type-params", "", "", "Destination Type Params")
	_finspacedataCmd.Flags().StringVarP(&_finspacedataDurationInMinutes, "duration-in-minutes", "", "", "Duration In Minutes")
	_finspacedataCmd.Flags().StringVarP(&_finspacedataEmailAddress, "email-address", "", "", "Email Address")
	_finspacedataCmd.Flags().StringVarP(&_finspacedataEnvironmentId, "environment-id", "", "", "Environment ID")
	_finspacedataCmd.Flags().StringVarP(&_finspacedataFirstName, "first-name", "", "", "First Name")
	_finspacedataCmd.Flags().StringVarP(&_finspacedataFormatParams, "format-params", "", "", "Format Params")
	_finspacedataCmd.Flags().StringVarP(&_finspacedataKind, "kind", "", "", "Kind")
	_finspacedataCmd.Flags().StringVarP(&_finspacedataLastName, "last-name", "", "", "Last Name")
	_finspacedataCmd.Flags().StringVarP(&_finspacedataLocationType, "location-type", "", "", "Location Type")
	_finspacedataCmd.Flags().StringVarP(&_finspacedataMaxResults, "max-results", "", "", "Max Results")
	_finspacedataCmd.Flags().StringVarP(&_finspacedataName, "name", "", "", "Name")
	_finspacedataCmd.Flags().StringVarP(&_finspacedataNextToken, "next-token", "", "", "Next Token")
	_finspacedataCmd.Flags().StringVarP(&_finspacedataOwnerInfo, "owner-info", "", "", "Owner Info")
	_finspacedataCmd.Flags().StringSliceVarP(&_finspacedataPartitionColumns, "partition-columns", "", nil, "Partition Columns")
	_finspacedataCmd.Flags().StringVarP(&_finspacedataPermissionGroupId, "permission-group-id", "", "", "Permission Group ID")
	_finspacedataCmd.Flags().StringVarP(&_finspacedataPermissionGroupParams, "permission-group-params", "", "", "Permission Group Params")
	_finspacedataCmd.Flags().StringVarP(&_finspacedataSchemaDefinition, "schema-definition", "", "", "Schema Definition")
	_finspacedataCmd.Flags().StringSliceVarP(&_finspacedataSortColumns, "sort-columns", "", nil, "Sort Columns")
	_finspacedataCmd.Flags().StringVarP(&_finspacedataSourceParams, "source-params", "", "", "Source Params")
	_finspacedataCmd.Flags().StringVarP(&_finspacedataType, "type", "", "", "Type")
	_finspacedataCmd.Flags().StringVarP(&_finspacedataUserId, "user-id", "", "", "User ID")

	_finspacedataCmd.Flags().BoolVarP(&_finspacedataAssociateUserToPermissionGroup, "associate-user-to-permission-group", "", false, "Associate User To Permission Group")
	_finspacedataCmd.Flags().BoolVarP(&_finspacedataCreateChangeset, "create-changeset", "", false, "Create Changeset")
	_finspacedataCmd.Flags().BoolVarP(&_finspacedataCreateDataView, "create-data-view", "", false, "Create Data View")
	_finspacedataCmd.Flags().BoolVarP(&_finspacedataCreateDataset, "create-dataset", "", false, "Create Dataset")
	_finspacedataCmd.Flags().BoolVarP(&_finspacedataCreatePermissionGroup, "create-permission-group", "", false, "Create Permission Group")
	_finspacedataCmd.Flags().BoolVarP(&_finspacedataCreateUser, "create-user", "", false, "Create User")
	_finspacedataCmd.Flags().BoolVarP(&_finspacedataDeleteDataset, "delete-dataset", "", false, "Delete Dataset")
	_finspacedataCmd.Flags().BoolVarP(&_finspacedataDeletePermissionGroup, "delete-permission-group", "", false, "Delete Permission Group")
	_finspacedataCmd.Flags().BoolVarP(&_finspacedataDisableUser, "disable-user", "", false, "Disable User")
	_finspacedataCmd.Flags().BoolVarP(&_finspacedataDisassociateUserFromPermissionGroup, "disassociate-user-from-permission-group", "", false, "Disassociate User From Permission Group")
	_finspacedataCmd.Flags().BoolVarP(&_finspacedataEnableUser, "enable-user", "", false, "Enable User")
	_finspacedataCmd.Flags().BoolVarP(&_finspacedataGetChangeset, "get-changeset", "", false, "Get Changeset")
	_finspacedataCmd.Flags().BoolVarP(&_finspacedataGetDataView, "get-data-view", "", false, "Get Data View")
	_finspacedataCmd.Flags().BoolVarP(&_finspacedataGetDataset, "get-dataset", "", false, "Get Dataset")
	_finspacedataCmd.Flags().BoolVarP(&_finspacedataGetExternalDataViewAccessDetails, "get-external-data-view-access-details", "", false, "Get External Data View Access Details")
	_finspacedataCmd.Flags().BoolVarP(&_finspacedataGetPermissionGroup, "get-permission-group", "", false, "Get Permission Group")
	_finspacedataCmd.Flags().BoolVarP(&_finspacedataGetProgrammaticAccessCredentials, "get-programmatic-access-credentials", "", false, "Get Programmatic Access Credentials")
	_finspacedataCmd.Flags().BoolVarP(&_finspacedataGetUser, "get-user", "", false, "Get User")
	_finspacedataCmd.Flags().BoolVarP(&_finspacedataGetWorkingLocation, "get-working-location", "", false, "Get Working Location")
	_finspacedataCmd.Flags().BoolVarP(&_finspacedataListChangesets, "list-changesets", "", false, "List Changesets")
	_finspacedataCmd.Flags().BoolVarP(&_finspacedataListDataViews, "list-data-views", "", false, "List Data Views")
	_finspacedataCmd.Flags().BoolVarP(&_finspacedataListDatasets, "list-datasets", "", false, "List Datasets")
	_finspacedataCmd.Flags().BoolVarP(&_finspacedataListPermissionGroups, "list-permission-groups", "", false, "List Permission Groups")
	_finspacedataCmd.Flags().BoolVarP(&_finspacedataListPermissionGroupsByUser, "list-permission-groups-by-user", "", false, "List Permission Groups By User")
	_finspacedataCmd.Flags().BoolVarP(&_finspacedataListUsers, "list-users", "", false, "List Users")
	_finspacedataCmd.Flags().BoolVarP(&_finspacedataListUsersByPermissionGroup, "list-users-by-permission-group", "", false, "List Users By Permission Group")
	_finspacedataCmd.Flags().BoolVarP(&_finspacedataResetUserPassword, "reset-user-password", "", false, "Reset User Password")
	_finspacedataCmd.Flags().BoolVarP(&_finspacedataUpdateChangeset, "update-changeset", "", false, "Update Changeset")
	_finspacedataCmd.Flags().BoolVarP(&_finspacedataUpdateDataset, "update-dataset", "", false, "Update Dataset")
	_finspacedataCmd.Flags().BoolVarP(&_finspacedataUpdatePermissionGroup, "update-permission-group", "", false, "Update Permission Group")
	_finspacedataCmd.Flags().BoolVarP(&_finspacedataUpdateUser, "update-user", "", false, "Update User")

}
