package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkidentity"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// chimesdkidentityCmd represents the chimesdkidentity command
var _chimesdkidentityCmd = &cobra.Command{
	Use:   "chimesdkidentity",
	Short: "AWS chimesdkidentity CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := chimesdkidentity.NewFromConfig(cfg)
		if _chimesdkidentityCreateAppInstance {
			chimesdkidentity_CreateAppInstance(cfg, client)
			return
		}
		if _chimesdkidentityCreateAppInstanceAdmin {
			chimesdkidentity_CreateAppInstanceAdmin(cfg, client)
			return
		}
		if _chimesdkidentityCreateAppInstanceBot {
			chimesdkidentity_CreateAppInstanceBot(cfg, client)
			return
		}
		if _chimesdkidentityCreateAppInstanceUser {
			chimesdkidentity_CreateAppInstanceUser(cfg, client)
			return
		}
		if _chimesdkidentityDeleteAppInstance {
			chimesdkidentity_DeleteAppInstance(cfg, client)
			return
		}
		if _chimesdkidentityDeleteAppInstanceAdmin {
			chimesdkidentity_DeleteAppInstanceAdmin(cfg, client)
			return
		}
		if _chimesdkidentityDeleteAppInstanceBot {
			chimesdkidentity_DeleteAppInstanceBot(cfg, client)
			return
		}
		if _chimesdkidentityDeleteAppInstanceUser {
			chimesdkidentity_DeleteAppInstanceUser(cfg, client)
			return
		}
		if _chimesdkidentityDeregisterAppInstanceUserEndpoint {
			chimesdkidentity_DeregisterAppInstanceUserEndpoint(cfg, client)
			return
		}
		if _chimesdkidentityDescribeAppInstance {
			chimesdkidentity_DescribeAppInstance(cfg, client)
			return
		}
		if _chimesdkidentityDescribeAppInstanceAdmin {
			chimesdkidentity_DescribeAppInstanceAdmin(cfg, client)
			return
		}
		if _chimesdkidentityDescribeAppInstanceBot {
			chimesdkidentity_DescribeAppInstanceBot(cfg, client)
			return
		}
		if _chimesdkidentityDescribeAppInstanceUser {
			chimesdkidentity_DescribeAppInstanceUser(cfg, client)
			return
		}
		if _chimesdkidentityDescribeAppInstanceUserEndpoint {
			chimesdkidentity_DescribeAppInstanceUserEndpoint(cfg, client)
			return
		}
		if _chimesdkidentityGetAppInstanceRetentionSettings {
			chimesdkidentity_GetAppInstanceRetentionSettings(cfg, client)
			return
		}
		if _chimesdkidentityListAppInstanceAdmins {
			chimesdkidentity_ListAppInstanceAdmins(cfg, client)
			return
		}
		if _chimesdkidentityListAppInstanceBots {
			chimesdkidentity_ListAppInstanceBots(cfg, client)
			return
		}
		if _chimesdkidentityListAppInstanceUserEndpoints {
			chimesdkidentity_ListAppInstanceUserEndpoints(cfg, client)
			return
		}
		if _chimesdkidentityListAppInstanceUsers {
			chimesdkidentity_ListAppInstanceUsers(cfg, client)
			return
		}
		if _chimesdkidentityListAppInstances {
			chimesdkidentity_ListAppInstances(cfg, client)
			return
		}
		if _chimesdkidentityListTagsForResource {
			chimesdkidentity_ListTagsForResource(cfg, client)
			return
		}
		if _chimesdkidentityPutAppInstanceRetentionSettings {
			chimesdkidentity_PutAppInstanceRetentionSettings(cfg, client)
			return
		}
		if _chimesdkidentityPutAppInstanceUserExpirationSettings {
			chimesdkidentity_PutAppInstanceUserExpirationSettings(cfg, client)
			return
		}
		if _chimesdkidentityRegisterAppInstanceUserEndpoint {
			chimesdkidentity_RegisterAppInstanceUserEndpoint(cfg, client)
			return
		}
		if _chimesdkidentityTagResource {
			chimesdkidentity_TagResource(cfg, client)
			return
		}
		if _chimesdkidentityUntagResource {
			chimesdkidentity_UntagResource(cfg, client)
			return
		}
		if _chimesdkidentityUpdateAppInstance {
			chimesdkidentity_UpdateAppInstance(cfg, client)
			return
		}
		if _chimesdkidentityUpdateAppInstanceBot {
			chimesdkidentity_UpdateAppInstanceBot(cfg, client)
			return
		}
		if _chimesdkidentityUpdateAppInstanceUser {
			chimesdkidentity_UpdateAppInstanceUser(cfg, client)
			return
		}
		if _chimesdkidentityUpdateAppInstanceUserEndpoint {
			chimesdkidentity_UpdateAppInstanceUserEndpoint(cfg, client)
			return
		}

	},
}

var (
	_chimesdkidentityCreateAppInstance                    bool
	_chimesdkidentityCreateAppInstanceAdmin               bool
	_chimesdkidentityCreateAppInstanceBot                 bool
	_chimesdkidentityCreateAppInstanceUser                bool
	_chimesdkidentityDeleteAppInstance                    bool
	_chimesdkidentityDeleteAppInstanceAdmin               bool
	_chimesdkidentityDeleteAppInstanceBot                 bool
	_chimesdkidentityDeleteAppInstanceUser                bool
	_chimesdkidentityDeregisterAppInstanceUserEndpoint    bool
	_chimesdkidentityDescribeAppInstance                  bool
	_chimesdkidentityDescribeAppInstanceAdmin             bool
	_chimesdkidentityDescribeAppInstanceBot               bool
	_chimesdkidentityDescribeAppInstanceUser              bool
	_chimesdkidentityDescribeAppInstanceUserEndpoint      bool
	_chimesdkidentityGetAppInstanceRetentionSettings      bool
	_chimesdkidentityListAppInstanceAdmins                bool
	_chimesdkidentityListAppInstanceBots                  bool
	_chimesdkidentityListAppInstanceUserEndpoints         bool
	_chimesdkidentityListAppInstanceUsers                 bool
	_chimesdkidentityListAppInstances                     bool
	_chimesdkidentityListTagsForResource                  bool
	_chimesdkidentityPutAppInstanceRetentionSettings      bool
	_chimesdkidentityPutAppInstanceUserExpirationSettings bool
	_chimesdkidentityRegisterAppInstanceUserEndpoint      bool
	_chimesdkidentityTagResource                          bool
	_chimesdkidentityUntagResource                        bool
	_chimesdkidentityUpdateAppInstance                    bool
	_chimesdkidentityUpdateAppInstanceBot                 bool
	_chimesdkidentityUpdateAppInstanceUser                bool
	_chimesdkidentityUpdateAppInstanceUserEndpoint        bool

	_chimesdkidentityAllowMessages                string
	_chimesdkidentityAppInstanceAdminArn          string
	_chimesdkidentityAppInstanceArn               string
	_chimesdkidentityAppInstanceBotArn            string
	_chimesdkidentityAppInstanceRetentionSettings string
	_chimesdkidentityAppInstanceUserArn           string
	_chimesdkidentityAppInstanceUserId            string
	_chimesdkidentityClientRequestToken           string
	_chimesdkidentityConfiguration                string
	_chimesdkidentityEndpointAttributes           string
	_chimesdkidentityEndpointId                   string
	_chimesdkidentityExpirationSettings           string
	_chimesdkidentityMaxResults                   string
	_chimesdkidentityMetadata                     string
	_chimesdkidentityName                         string
	_chimesdkidentityNextToken                    string
	_chimesdkidentityResourceARN                  string
	_chimesdkidentityTagKeys                      []string
	_chimesdkidentityTags                         string
	_chimesdkidentityType                         string
)

// Creates an Amazon Chime SDK messaging AppInstance under an AWS account. Only
// SDK messaging customers use this API. CreateAppInstance supports idempotency
// behavior as described in the AWS API Standard.
//
// identity
func chimesdkidentity_CreateAppInstance(cfg aws.Config, client *chimesdkidentity.Client) {
	input := &chimesdkidentity.CreateAppInstanceInput{
		// ClientRequestToken: *string, // Required
		// Name: *string, // Required
	}

	if len(_chimesdkidentityClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_chimesdkidentityClientRequestToken)
	}
	if len(_chimesdkidentityName) > 0 {
		input.Name = aws.String(_chimesdkidentityName)
	}
	if len(_chimesdkidentityMetadata) > 0 {
		input.Metadata = aws.String(_chimesdkidentityMetadata)
	}
	if len(_chimesdkidentityTags) > 0 {
		if err := assignInputField(input, "Tags", _chimesdkidentityTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAppInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Promotes an AppInstanceUser or AppInstanceBot to an AppInstanceAdmin . The
// promoted entity can perform the following actions.
//
// - ChannelModerator actions across all channels in the AppInstance .
//
// - DeleteChannelMessage actions.
//
// Only an AppInstanceUser and AppInstanceBot can be promoted to an
// AppInstanceAdmin role.
func chimesdkidentity_CreateAppInstanceAdmin(cfg aws.Config, client *chimesdkidentity.Client) {
	input := &chimesdkidentity.CreateAppInstanceAdminInput{
		// AppInstanceAdminArn: *string, // Required
		// AppInstanceArn: *string, // Required
	}

	if len(_chimesdkidentityAppInstanceAdminArn) > 0 {
		input.AppInstanceAdminArn = aws.String(_chimesdkidentityAppInstanceAdminArn)
	}
	if len(_chimesdkidentityAppInstanceArn) > 0 {
		input.AppInstanceArn = aws.String(_chimesdkidentityAppInstanceArn)
	}

	if resp, err := client.CreateAppInstanceAdmin(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a bot under an Amazon Chime AppInstance . The request consists of a
// unique Configuration and Name for that bot.
func chimesdkidentity_CreateAppInstanceBot(cfg aws.Config, client *chimesdkidentity.Client) {
	input := &chimesdkidentity.CreateAppInstanceBotInput{
		// AppInstanceArn: *string, // Required
		// ClientRequestToken: *string, // Required
		// Configuration: *types.Configuration, // Required
	}

	if len(_chimesdkidentityAppInstanceArn) > 0 {
		input.AppInstanceArn = aws.String(_chimesdkidentityAppInstanceArn)
	}
	if len(_chimesdkidentityClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_chimesdkidentityClientRequestToken)
	}
	if len(_chimesdkidentityConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _chimesdkidentityConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_chimesdkidentityMetadata) > 0 {
		input.Metadata = aws.String(_chimesdkidentityMetadata)
	}
	if len(_chimesdkidentityName) > 0 {
		input.Name = aws.String(_chimesdkidentityName)
	}
	if len(_chimesdkidentityTags) > 0 {
		if err := assignInputField(input, "Tags", _chimesdkidentityTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAppInstanceBot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a user under an Amazon Chime AppInstance . The request consists of a
// unique appInstanceUserId and Name for that user.
func chimesdkidentity_CreateAppInstanceUser(cfg aws.Config, client *chimesdkidentity.Client) {
	input := &chimesdkidentity.CreateAppInstanceUserInput{
		// AppInstanceArn: *string, // Required
		// AppInstanceUserId: *string, // Required
		// ClientRequestToken: *string, // Required
		// Name: *string, // Required
	}

	if len(_chimesdkidentityAppInstanceArn) > 0 {
		input.AppInstanceArn = aws.String(_chimesdkidentityAppInstanceArn)
	}
	if len(_chimesdkidentityAppInstanceUserId) > 0 {
		input.AppInstanceUserId = aws.String(_chimesdkidentityAppInstanceUserId)
	}
	if len(_chimesdkidentityClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_chimesdkidentityClientRequestToken)
	}
	if len(_chimesdkidentityName) > 0 {
		input.Name = aws.String(_chimesdkidentityName)
	}
	if len(_chimesdkidentityExpirationSettings) > 0 {
		if err := assignInputField(input, "ExpirationSettings", _chimesdkidentityExpirationSettings); err != nil {
			log.Errorf("invalid --expiration-settings: %s", err.Error())
			return
		}
	}
	if len(_chimesdkidentityMetadata) > 0 {
		input.Metadata = aws.String(_chimesdkidentityMetadata)
	}
	if len(_chimesdkidentityTags) > 0 {
		if err := assignInputField(input, "Tags", _chimesdkidentityTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAppInstanceUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an AppInstance and all associated data asynchronously.
func chimesdkidentity_DeleteAppInstance(cfg aws.Config, client *chimesdkidentity.Client) {
	input := &chimesdkidentity.DeleteAppInstanceInput{
		// AppInstanceArn: *string, // Required
	}

	if len(_chimesdkidentityAppInstanceArn) > 0 {
		input.AppInstanceArn = aws.String(_chimesdkidentityAppInstanceArn)
	}

	if resp, err := client.DeleteAppInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Demotes an AppInstanceAdmin to an AppInstanceUser or AppInstanceBot . This
// action does not delete the user.
func chimesdkidentity_DeleteAppInstanceAdmin(cfg aws.Config, client *chimesdkidentity.Client) {
	input := &chimesdkidentity.DeleteAppInstanceAdminInput{
		// AppInstanceAdminArn: *string, // Required
		// AppInstanceArn: *string, // Required
	}

	if len(_chimesdkidentityAppInstanceAdminArn) > 0 {
		input.AppInstanceAdminArn = aws.String(_chimesdkidentityAppInstanceAdminArn)
	}
	if len(_chimesdkidentityAppInstanceArn) > 0 {
		input.AppInstanceArn = aws.String(_chimesdkidentityAppInstanceArn)
	}

	if resp, err := client.DeleteAppInstanceAdmin(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an AppInstanceBot .
func chimesdkidentity_DeleteAppInstanceBot(cfg aws.Config, client *chimesdkidentity.Client) {
	input := &chimesdkidentity.DeleteAppInstanceBotInput{
		// AppInstanceBotArn: *string, // Required
	}

	if len(_chimesdkidentityAppInstanceBotArn) > 0 {
		input.AppInstanceBotArn = aws.String(_chimesdkidentityAppInstanceBotArn)
	}

	if resp, err := client.DeleteAppInstanceBot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an AppInstanceUser .
func chimesdkidentity_DeleteAppInstanceUser(cfg aws.Config, client *chimesdkidentity.Client) {
	input := &chimesdkidentity.DeleteAppInstanceUserInput{
		// AppInstanceUserArn: *string, // Required
	}

	if len(_chimesdkidentityAppInstanceUserArn) > 0 {
		input.AppInstanceUserArn = aws.String(_chimesdkidentityAppInstanceUserArn)
	}

	if resp, err := client.DeleteAppInstanceUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deregisters an AppInstanceUserEndpoint .
func chimesdkidentity_DeregisterAppInstanceUserEndpoint(cfg aws.Config, client *chimesdkidentity.Client) {
	input := &chimesdkidentity.DeregisterAppInstanceUserEndpointInput{
		// AppInstanceUserArn: *string, // Required
		// EndpointId: *string, // Required
	}

	if len(_chimesdkidentityAppInstanceUserArn) > 0 {
		input.AppInstanceUserArn = aws.String(_chimesdkidentityAppInstanceUserArn)
	}
	if len(_chimesdkidentityEndpointId) > 0 {
		input.EndpointId = aws.String(_chimesdkidentityEndpointId)
	}

	if resp, err := client.DeregisterAppInstanceUserEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the full details of an AppInstance .
func chimesdkidentity_DescribeAppInstance(cfg aws.Config, client *chimesdkidentity.Client) {
	input := &chimesdkidentity.DescribeAppInstanceInput{
		// AppInstanceArn: *string, // Required
	}

	if len(_chimesdkidentityAppInstanceArn) > 0 {
		input.AppInstanceArn = aws.String(_chimesdkidentityAppInstanceArn)
	}

	if resp, err := client.DescribeAppInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the full details of an AppInstanceAdmin .
func chimesdkidentity_DescribeAppInstanceAdmin(cfg aws.Config, client *chimesdkidentity.Client) {
	input := &chimesdkidentity.DescribeAppInstanceAdminInput{
		// AppInstanceAdminArn: *string, // Required
		// AppInstanceArn: *string, // Required
	}

	if len(_chimesdkidentityAppInstanceAdminArn) > 0 {
		input.AppInstanceAdminArn = aws.String(_chimesdkidentityAppInstanceAdminArn)
	}
	if len(_chimesdkidentityAppInstanceArn) > 0 {
		input.AppInstanceArn = aws.String(_chimesdkidentityAppInstanceArn)
	}

	if resp, err := client.DescribeAppInstanceAdmin(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The AppInstanceBot's information.
func chimesdkidentity_DescribeAppInstanceBot(cfg aws.Config, client *chimesdkidentity.Client) {
	input := &chimesdkidentity.DescribeAppInstanceBotInput{
		// AppInstanceBotArn: *string, // Required
	}

	if len(_chimesdkidentityAppInstanceBotArn) > 0 {
		input.AppInstanceBotArn = aws.String(_chimesdkidentityAppInstanceBotArn)
	}

	if resp, err := client.DescribeAppInstanceBot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the full details of an AppInstanceUser .
func chimesdkidentity_DescribeAppInstanceUser(cfg aws.Config, client *chimesdkidentity.Client) {
	input := &chimesdkidentity.DescribeAppInstanceUserInput{
		// AppInstanceUserArn: *string, // Required
	}

	if len(_chimesdkidentityAppInstanceUserArn) > 0 {
		input.AppInstanceUserArn = aws.String(_chimesdkidentityAppInstanceUserArn)
	}

	if resp, err := client.DescribeAppInstanceUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the full details of an AppInstanceUserEndpoint .
func chimesdkidentity_DescribeAppInstanceUserEndpoint(cfg aws.Config, client *chimesdkidentity.Client) {
	input := &chimesdkidentity.DescribeAppInstanceUserEndpointInput{
		// AppInstanceUserArn: *string, // Required
		// EndpointId: *string, // Required
	}

	if len(_chimesdkidentityAppInstanceUserArn) > 0 {
		input.AppInstanceUserArn = aws.String(_chimesdkidentityAppInstanceUserArn)
	}
	if len(_chimesdkidentityEndpointId) > 0 {
		input.EndpointId = aws.String(_chimesdkidentityEndpointId)
	}

	if resp, err := client.DescribeAppInstanceUserEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the retention settings for an AppInstance .
func chimesdkidentity_GetAppInstanceRetentionSettings(cfg aws.Config, client *chimesdkidentity.Client) {
	input := &chimesdkidentity.GetAppInstanceRetentionSettingsInput{
		// AppInstanceArn: *string, // Required
	}

	if len(_chimesdkidentityAppInstanceArn) > 0 {
		input.AppInstanceArn = aws.String(_chimesdkidentityAppInstanceArn)
	}

	if resp, err := client.GetAppInstanceRetentionSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of the administrators in the AppInstance .
func chimesdkidentity_ListAppInstanceAdmins(cfg aws.Config, client *chimesdkidentity.Client) {
	input := &chimesdkidentity.ListAppInstanceAdminsInput{
		// AppInstanceArn: *string, // Required
	}

	if len(_chimesdkidentityAppInstanceArn) > 0 {
		input.AppInstanceArn = aws.String(_chimesdkidentityAppInstanceArn)
	}
	if len(_chimesdkidentityMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _chimesdkidentityMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_chimesdkidentityNextToken) > 0 {
		input.NextToken = aws.String(_chimesdkidentityNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAppInstanceAdmins(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*chimesdkidentity.ListAppInstanceAdminsOutput
	p := chimesdkidentity.NewListAppInstanceAdminsPaginator(client, input)
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

// Lists all AppInstanceBots created under a single AppInstance .
func chimesdkidentity_ListAppInstanceBots(cfg aws.Config, client *chimesdkidentity.Client) {
	input := &chimesdkidentity.ListAppInstanceBotsInput{
		// AppInstanceArn: *string, // Required
	}

	if len(_chimesdkidentityAppInstanceArn) > 0 {
		input.AppInstanceArn = aws.String(_chimesdkidentityAppInstanceArn)
	}
	if len(_chimesdkidentityMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _chimesdkidentityMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_chimesdkidentityNextToken) > 0 {
		input.NextToken = aws.String(_chimesdkidentityNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAppInstanceBots(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*chimesdkidentity.ListAppInstanceBotsOutput
	p := chimesdkidentity.NewListAppInstanceBotsPaginator(client, input)
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

// Lists all the AppInstanceUserEndpoints created under a single AppInstanceUser .
func chimesdkidentity_ListAppInstanceUserEndpoints(cfg aws.Config, client *chimesdkidentity.Client) {
	input := &chimesdkidentity.ListAppInstanceUserEndpointsInput{
		// AppInstanceUserArn: *string, // Required
	}

	if len(_chimesdkidentityAppInstanceUserArn) > 0 {
		input.AppInstanceUserArn = aws.String(_chimesdkidentityAppInstanceUserArn)
	}
	if len(_chimesdkidentityMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _chimesdkidentityMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_chimesdkidentityNextToken) > 0 {
		input.NextToken = aws.String(_chimesdkidentityNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAppInstanceUserEndpoints(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*chimesdkidentity.ListAppInstanceUserEndpointsOutput
	p := chimesdkidentity.NewListAppInstanceUserEndpointsPaginator(client, input)
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

// List all AppInstanceUsers created under a single AppInstance .
func chimesdkidentity_ListAppInstanceUsers(cfg aws.Config, client *chimesdkidentity.Client) {
	input := &chimesdkidentity.ListAppInstanceUsersInput{
		// AppInstanceArn: *string, // Required
	}

	if len(_chimesdkidentityAppInstanceArn) > 0 {
		input.AppInstanceArn = aws.String(_chimesdkidentityAppInstanceArn)
	}
	if len(_chimesdkidentityMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _chimesdkidentityMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_chimesdkidentityNextToken) > 0 {
		input.NextToken = aws.String(_chimesdkidentityNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAppInstanceUsers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*chimesdkidentity.ListAppInstanceUsersOutput
	p := chimesdkidentity.NewListAppInstanceUsersPaginator(client, input)
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

// Lists all Amazon Chime AppInstance s created under a single AWS account.
func chimesdkidentity_ListAppInstances(cfg aws.Config, client *chimesdkidentity.Client) {
	input := &chimesdkidentity.ListAppInstancesInput{}

	if len(_chimesdkidentityMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _chimesdkidentityMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_chimesdkidentityNextToken) > 0 {
		input.NextToken = aws.String(_chimesdkidentityNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAppInstances(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*chimesdkidentity.ListAppInstancesOutput
	p := chimesdkidentity.NewListAppInstancesPaginator(client, input)
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

// Lists the tags applied to an Amazon Chime SDK identity resource.
func chimesdkidentity_ListTagsForResource(cfg aws.Config, client *chimesdkidentity.Client) {
	input := &chimesdkidentity.ListTagsForResourceInput{
		// ResourceARN: *string, // Required
	}

	if len(_chimesdkidentityResourceARN) > 0 {
		input.ResourceARN = aws.String(_chimesdkidentityResourceARN)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the amount of time in days that a given AppInstance retains data.
func chimesdkidentity_PutAppInstanceRetentionSettings(cfg aws.Config, client *chimesdkidentity.Client) {
	input := &chimesdkidentity.PutAppInstanceRetentionSettingsInput{
		// AppInstanceArn: *string, // Required
		// AppInstanceRetentionSettings: *types.AppInstanceRetentionSettings, // Required
	}

	if len(_chimesdkidentityAppInstanceArn) > 0 {
		input.AppInstanceArn = aws.String(_chimesdkidentityAppInstanceArn)
	}
	if len(_chimesdkidentityAppInstanceRetentionSettings) > 0 {
		if err := assignInputField(input, "AppInstanceRetentionSettings", _chimesdkidentityAppInstanceRetentionSettings); err != nil {
			log.Errorf("invalid --app-instance-retention-settings: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutAppInstanceRetentionSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the number of days before the AppInstanceUser is automatically deleted.
// A background process deletes expired AppInstanceUsers within 6 hours of
// expiration. Actual deletion times may vary.
//
// Expired AppInstanceUsers that have not yet been deleted appear as active, and
// you can update their expiration settings. The system honors the new settings.
func chimesdkidentity_PutAppInstanceUserExpirationSettings(cfg aws.Config, client *chimesdkidentity.Client) {
	input := &chimesdkidentity.PutAppInstanceUserExpirationSettingsInput{
		// AppInstanceUserArn: *string, // Required
	}

	if len(_chimesdkidentityAppInstanceUserArn) > 0 {
		input.AppInstanceUserArn = aws.String(_chimesdkidentityAppInstanceUserArn)
	}
	if len(_chimesdkidentityExpirationSettings) > 0 {
		if err := assignInputField(input, "ExpirationSettings", _chimesdkidentityExpirationSettings); err != nil {
			log.Errorf("invalid --expiration-settings: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutAppInstanceUserExpirationSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Registers an endpoint under an Amazon Chime AppInstanceUser . The endpoint
// receives messages for a user. For push notifications, the endpoint is a mobile
// device used to receive mobile push notifications for a user.
func chimesdkidentity_RegisterAppInstanceUserEndpoint(cfg aws.Config, client *chimesdkidentity.Client) {
	input := &chimesdkidentity.RegisterAppInstanceUserEndpointInput{
		// AppInstanceUserArn: *string, // Required
		// ClientRequestToken: *string, // Required
		// EndpointAttributes: *types.EndpointAttributes, // Required
		// ResourceArn: *string, // Required
		// Type: types.AppInstanceUserEndpointType, // Required
	}

	if len(_chimesdkidentityAppInstanceUserArn) > 0 {
		input.AppInstanceUserArn = aws.String(_chimesdkidentityAppInstanceUserArn)
	}
	if len(_chimesdkidentityClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_chimesdkidentityClientRequestToken)
	}
	if len(_chimesdkidentityEndpointAttributes) > 0 {
		if err := assignInputField(input, "EndpointAttributes", _chimesdkidentityEndpointAttributes); err != nil {
			log.Errorf("invalid --endpoint-attributes: %s", err.Error())
			return
		}
	}
	if len(_chimesdkidentityResourceARN) > 0 {
		input.ResourceArn = aws.String(_chimesdkidentityResourceARN)
	}
	if len(_chimesdkidentityType) > 0 {
		if err := assignInputField(input, "Type", _chimesdkidentityType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_chimesdkidentityAllowMessages) > 0 {
		if err := assignInputField(input, "AllowMessages", _chimesdkidentityAllowMessages); err != nil {
			log.Errorf("invalid --allow-messages: %s", err.Error())
			return
		}
	}
	if len(_chimesdkidentityName) > 0 {
		input.Name = aws.String(_chimesdkidentityName)
	}

	if resp, err := client.RegisterAppInstanceUserEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Applies the specified tags to the specified Amazon Chime SDK identity resource.
func chimesdkidentity_TagResource(cfg aws.Config, client *chimesdkidentity.Client) {
	input := &chimesdkidentity.TagResourceInput{
		// ResourceARN: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_chimesdkidentityResourceARN) > 0 {
		input.ResourceARN = aws.String(_chimesdkidentityResourceARN)
	}
	if len(_chimesdkidentityTags) > 0 {
		if err := assignInputField(input, "Tags", _chimesdkidentityTags); err != nil {
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

// Removes the specified tags from the specified Amazon Chime SDK identity
// resource.
func chimesdkidentity_UntagResource(cfg aws.Config, client *chimesdkidentity.Client) {
	input := &chimesdkidentity.UntagResourceInput{
		// ResourceARN: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_chimesdkidentityResourceARN) > 0 {
		input.ResourceARN = aws.String(_chimesdkidentityResourceARN)
	}
	if len(_chimesdkidentityTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _chimesdkidentityTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates AppInstance metadata.
func chimesdkidentity_UpdateAppInstance(cfg aws.Config, client *chimesdkidentity.Client) {
	input := &chimesdkidentity.UpdateAppInstanceInput{
		// AppInstanceArn: *string, // Required
		// Metadata: *string, // Required
		// Name: *string, // Required
	}

	if len(_chimesdkidentityAppInstanceArn) > 0 {
		input.AppInstanceArn = aws.String(_chimesdkidentityAppInstanceArn)
	}
	if len(_chimesdkidentityMetadata) > 0 {
		input.Metadata = aws.String(_chimesdkidentityMetadata)
	}
	if len(_chimesdkidentityName) > 0 {
		input.Name = aws.String(_chimesdkidentityName)
	}

	if resp, err := client.UpdateAppInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the name and metadata of an AppInstanceBot .
func chimesdkidentity_UpdateAppInstanceBot(cfg aws.Config, client *chimesdkidentity.Client) {
	input := &chimesdkidentity.UpdateAppInstanceBotInput{
		// AppInstanceBotArn: *string, // Required
		// Metadata: *string, // Required
		// Name: *string, // Required
	}

	if len(_chimesdkidentityAppInstanceBotArn) > 0 {
		input.AppInstanceBotArn = aws.String(_chimesdkidentityAppInstanceBotArn)
	}
	if len(_chimesdkidentityMetadata) > 0 {
		input.Metadata = aws.String(_chimesdkidentityMetadata)
	}
	if len(_chimesdkidentityName) > 0 {
		input.Name = aws.String(_chimesdkidentityName)
	}
	if len(_chimesdkidentityConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _chimesdkidentityConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAppInstanceBot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the details of an AppInstanceUser . You can update names and metadata.
func chimesdkidentity_UpdateAppInstanceUser(cfg aws.Config, client *chimesdkidentity.Client) {
	input := &chimesdkidentity.UpdateAppInstanceUserInput{
		// AppInstanceUserArn: *string, // Required
		// Metadata: *string, // Required
		// Name: *string, // Required
	}

	if len(_chimesdkidentityAppInstanceUserArn) > 0 {
		input.AppInstanceUserArn = aws.String(_chimesdkidentityAppInstanceUserArn)
	}
	if len(_chimesdkidentityMetadata) > 0 {
		input.Metadata = aws.String(_chimesdkidentityMetadata)
	}
	if len(_chimesdkidentityName) > 0 {
		input.Name = aws.String(_chimesdkidentityName)
	}

	if resp, err := client.UpdateAppInstanceUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the details of an AppInstanceUserEndpoint . You can update the name and
// AllowMessage values.
func chimesdkidentity_UpdateAppInstanceUserEndpoint(cfg aws.Config, client *chimesdkidentity.Client) {
	input := &chimesdkidentity.UpdateAppInstanceUserEndpointInput{
		// AppInstanceUserArn: *string, // Required
		// EndpointId: *string, // Required
	}

	if len(_chimesdkidentityAppInstanceUserArn) > 0 {
		input.AppInstanceUserArn = aws.String(_chimesdkidentityAppInstanceUserArn)
	}
	if len(_chimesdkidentityEndpointId) > 0 {
		input.EndpointId = aws.String(_chimesdkidentityEndpointId)
	}
	if len(_chimesdkidentityAllowMessages) > 0 {
		if err := assignInputField(input, "AllowMessages", _chimesdkidentityAllowMessages); err != nil {
			log.Errorf("invalid --allow-messages: %s", err.Error())
			return
		}
	}
	if len(_chimesdkidentityName) > 0 {
		input.Name = aws.String(_chimesdkidentityName)
	}

	if resp, err := client.UpdateAppInstanceUserEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_chimesdkidentityCmd)
	_chimesdkidentityCmd.Flags().SortFlags = false

	_chimesdkidentityCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_chimesdkidentityCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_chimesdkidentityCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_chimesdkidentityCmd.Flags().StringVarP(&_chimesdkidentityAllowMessages, "allow-messages", "", "", "Allow Messages")
	_chimesdkidentityCmd.Flags().StringVarP(&_chimesdkidentityAppInstanceAdminArn, "app-instance-admin-arn", "", "", "App Instance Admin ARN")
	_chimesdkidentityCmd.Flags().StringVarP(&_chimesdkidentityAppInstanceArn, "app-instance-arn", "", "", "App Instance ARN")
	_chimesdkidentityCmd.Flags().StringVarP(&_chimesdkidentityAppInstanceBotArn, "app-instance-bot-arn", "", "", "App Instance Bot ARN")
	_chimesdkidentityCmd.Flags().StringVarP(&_chimesdkidentityAppInstanceRetentionSettings, "app-instance-retention-settings", "", "", "App Instance Retention Settings")
	_chimesdkidentityCmd.Flags().StringVarP(&_chimesdkidentityAppInstanceUserArn, "app-instance-user-arn", "", "", "App Instance User ARN")
	_chimesdkidentityCmd.Flags().StringVarP(&_chimesdkidentityAppInstanceUserId, "app-instance-user-id", "", "", "App Instance User ID")
	_chimesdkidentityCmd.Flags().StringVarP(&_chimesdkidentityClientRequestToken, "client-request-token", "", "", "Client Request Token")
	_chimesdkidentityCmd.Flags().StringVarP(&_chimesdkidentityConfiguration, "configuration", "", "", "Configuration")
	_chimesdkidentityCmd.Flags().StringVarP(&_chimesdkidentityEndpointAttributes, "endpoint-attributes", "", "", "Endpoint Attributes")
	_chimesdkidentityCmd.Flags().StringVarP(&_chimesdkidentityEndpointId, "endpoint-id", "", "", "Endpoint ID")
	_chimesdkidentityCmd.Flags().StringVarP(&_chimesdkidentityExpirationSettings, "expiration-settings", "", "", "Expiration Settings")
	_chimesdkidentityCmd.Flags().StringVarP(&_chimesdkidentityMaxResults, "max-results", "", "", "Max Results")
	_chimesdkidentityCmd.Flags().StringVarP(&_chimesdkidentityMetadata, "metadata", "", "", "Metadata")
	_chimesdkidentityCmd.Flags().StringVarP(&_chimesdkidentityName, "name", "", "", "Name")
	_chimesdkidentityCmd.Flags().StringVarP(&_chimesdkidentityNextToken, "next-token", "", "", "Next Token")
	_chimesdkidentityCmd.Flags().StringVarP(&_chimesdkidentityResourceARN, "resource-arn", "", "", "Resource ARN")
	_chimesdkidentityCmd.Flags().StringSliceVarP(&_chimesdkidentityTagKeys, "tag-keys", "", nil, "Tag Keys")
	_chimesdkidentityCmd.Flags().StringVarP(&_chimesdkidentityTags, "tags", "", "", "Tags")
	_chimesdkidentityCmd.Flags().StringVarP(&_chimesdkidentityType, "type", "", "", "Type")

	_chimesdkidentityCmd.Flags().BoolVarP(&_chimesdkidentityCreateAppInstance, "create-app-instance", "", false, "Create App Instance")
	_chimesdkidentityCmd.Flags().BoolVarP(&_chimesdkidentityCreateAppInstanceAdmin, "create-app-instance-admin", "", false, "Create App Instance Admin")
	_chimesdkidentityCmd.Flags().BoolVarP(&_chimesdkidentityCreateAppInstanceBot, "create-app-instance-bot", "", false, "Create App Instance Bot")
	_chimesdkidentityCmd.Flags().BoolVarP(&_chimesdkidentityCreateAppInstanceUser, "create-app-instance-user", "", false, "Create App Instance User")
	_chimesdkidentityCmd.Flags().BoolVarP(&_chimesdkidentityDeleteAppInstance, "delete-app-instance", "", false, "Delete App Instance")
	_chimesdkidentityCmd.Flags().BoolVarP(&_chimesdkidentityDeleteAppInstanceAdmin, "delete-app-instance-admin", "", false, "Delete App Instance Admin")
	_chimesdkidentityCmd.Flags().BoolVarP(&_chimesdkidentityDeleteAppInstanceBot, "delete-app-instance-bot", "", false, "Delete App Instance Bot")
	_chimesdkidentityCmd.Flags().BoolVarP(&_chimesdkidentityDeleteAppInstanceUser, "delete-app-instance-user", "", false, "Delete App Instance User")
	_chimesdkidentityCmd.Flags().BoolVarP(&_chimesdkidentityDeregisterAppInstanceUserEndpoint, "deregister-app-instance-user-endpoint", "", false, "Deregister App Instance User Endpoint")
	_chimesdkidentityCmd.Flags().BoolVarP(&_chimesdkidentityDescribeAppInstance, "describe-app-instance", "", false, "Describe App Instance")
	_chimesdkidentityCmd.Flags().BoolVarP(&_chimesdkidentityDescribeAppInstanceAdmin, "describe-app-instance-admin", "", false, "Describe App Instance Admin")
	_chimesdkidentityCmd.Flags().BoolVarP(&_chimesdkidentityDescribeAppInstanceBot, "describe-app-instance-bot", "", false, "Describe App Instance Bot")
	_chimesdkidentityCmd.Flags().BoolVarP(&_chimesdkidentityDescribeAppInstanceUser, "describe-app-instance-user", "", false, "Describe App Instance User")
	_chimesdkidentityCmd.Flags().BoolVarP(&_chimesdkidentityDescribeAppInstanceUserEndpoint, "describe-app-instance-user-endpoint", "", false, "Describe App Instance User Endpoint")
	_chimesdkidentityCmd.Flags().BoolVarP(&_chimesdkidentityGetAppInstanceRetentionSettings, "get-app-instance-retention-settings", "", false, "Get App Instance Retention Settings")
	_chimesdkidentityCmd.Flags().BoolVarP(&_chimesdkidentityListAppInstanceAdmins, "list-app-instance-admins", "", false, "List App Instance Admins")
	_chimesdkidentityCmd.Flags().BoolVarP(&_chimesdkidentityListAppInstanceBots, "list-app-instance-bots", "", false, "List App Instance Bots")
	_chimesdkidentityCmd.Flags().BoolVarP(&_chimesdkidentityListAppInstanceUserEndpoints, "list-app-instance-user-endpoints", "", false, "List App Instance User Endpoints")
	_chimesdkidentityCmd.Flags().BoolVarP(&_chimesdkidentityListAppInstanceUsers, "list-app-instance-users", "", false, "List App Instance Users")
	_chimesdkidentityCmd.Flags().BoolVarP(&_chimesdkidentityListAppInstances, "list-app-instances", "", false, "List App Instances")
	_chimesdkidentityCmd.Flags().BoolVarP(&_chimesdkidentityListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_chimesdkidentityCmd.Flags().BoolVarP(&_chimesdkidentityPutAppInstanceRetentionSettings, "put-app-instance-retention-settings", "", false, "Put App Instance Retention Settings")
	_chimesdkidentityCmd.Flags().BoolVarP(&_chimesdkidentityPutAppInstanceUserExpirationSettings, "put-app-instance-user-expiration-settings", "", false, "Put App Instance User Expiration Settings")
	_chimesdkidentityCmd.Flags().BoolVarP(&_chimesdkidentityRegisterAppInstanceUserEndpoint, "register-app-instance-user-endpoint", "", false, "Register App Instance User Endpoint")
	_chimesdkidentityCmd.Flags().BoolVarP(&_chimesdkidentityTagResource, "tag-resource", "", false, "Tag Resource")
	_chimesdkidentityCmd.Flags().BoolVarP(&_chimesdkidentityUntagResource, "untag-resource", "", false, "Untag Resource")
	_chimesdkidentityCmd.Flags().BoolVarP(&_chimesdkidentityUpdateAppInstance, "update-app-instance", "", false, "Update App Instance")
	_chimesdkidentityCmd.Flags().BoolVarP(&_chimesdkidentityUpdateAppInstanceBot, "update-app-instance-bot", "", false, "Update App Instance Bot")
	_chimesdkidentityCmd.Flags().BoolVarP(&_chimesdkidentityUpdateAppInstanceUser, "update-app-instance-user", "", false, "Update App Instance User")
	_chimesdkidentityCmd.Flags().BoolVarP(&_chimesdkidentityUpdateAppInstanceUserEndpoint, "update-app-instance-user-endpoint", "", false, "Update App Instance User Endpoint")

}
