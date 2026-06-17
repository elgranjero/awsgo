package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mq"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// mqCmd represents the mq command
var _mqCmd = &cobra.Command{
	Use:   "mq",
	Short: "AWS mq CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := mq.NewFromConfig(cfg)
		if _mqCreateBroker {
			mq_CreateBroker(cfg, client)
			return
		}
		if _mqCreateConfiguration {
			mq_CreateConfiguration(cfg, client)
			return
		}
		if _mqCreateTags {
			mq_CreateTags(cfg, client)
			return
		}
		if _mqCreateUser {
			mq_CreateUser(cfg, client)
			return
		}
		if _mqDeleteBroker {
			mq_DeleteBroker(cfg, client)
			return
		}
		if _mqDeleteConfiguration {
			mq_DeleteConfiguration(cfg, client)
			return
		}
		if _mqDeleteTags {
			mq_DeleteTags(cfg, client)
			return
		}
		if _mqDeleteUser {
			mq_DeleteUser(cfg, client)
			return
		}
		if _mqDescribeBroker {
			mq_DescribeBroker(cfg, client)
			return
		}
		if _mqDescribeBrokerEngineTypes {
			mq_DescribeBrokerEngineTypes(cfg, client)
			return
		}
		if _mqDescribeBrokerInstanceOptions {
			mq_DescribeBrokerInstanceOptions(cfg, client)
			return
		}
		if _mqDescribeConfiguration {
			mq_DescribeConfiguration(cfg, client)
			return
		}
		if _mqDescribeConfigurationRevision {
			mq_DescribeConfigurationRevision(cfg, client)
			return
		}
		if _mqDescribeUser {
			mq_DescribeUser(cfg, client)
			return
		}
		if _mqListBrokers {
			mq_ListBrokers(cfg, client)
			return
		}
		if _mqListConfigurationRevisions {
			mq_ListConfigurationRevisions(cfg, client)
			return
		}
		if _mqListConfigurations {
			mq_ListConfigurations(cfg, client)
			return
		}
		if _mqListTags {
			mq_ListTags(cfg, client)
			return
		}
		if _mqListUsers {
			mq_ListUsers(cfg, client)
			return
		}
		if _mqPromote {
			mq_Promote(cfg, client)
			return
		}
		if _mqRebootBroker {
			mq_RebootBroker(cfg, client)
			return
		}
		if _mqUpdateBroker {
			mq_UpdateBroker(cfg, client)
			return
		}
		if _mqUpdateConfiguration {
			mq_UpdateConfiguration(cfg, client)
			return
		}
		if _mqUpdateUser {
			mq_UpdateUser(cfg, client)
			return
		}

	},
}

var (
	_mqCreateBroker                  bool
	_mqCreateConfiguration           bool
	_mqCreateTags                    bool
	_mqCreateUser                    bool
	_mqDeleteBroker                  bool
	_mqDeleteConfiguration           bool
	_mqDeleteTags                    bool
	_mqDeleteUser                    bool
	_mqDescribeBroker                bool
	_mqDescribeBrokerEngineTypes     bool
	_mqDescribeBrokerInstanceOptions bool
	_mqDescribeConfiguration         bool
	_mqDescribeConfigurationRevision bool
	_mqDescribeUser                  bool
	_mqListBrokers                   bool
	_mqListConfigurationRevisions    bool
	_mqListConfigurations            bool
	_mqListTags                      bool
	_mqListUsers                     bool
	_mqPromote                       bool
	_mqRebootBroker                  bool
	_mqUpdateBroker                  bool
	_mqUpdateConfiguration           bool
	_mqUpdateUser                    bool

	_mqAuthenticationStrategy          string
	_mqAutoMinorVersionUpgrade         string
	_mqBrokerId                        string
	_mqBrokerName                      string
	_mqConfiguration                   string
	_mqConfigurationId                 string
	_mqConfigurationRevision           string
	_mqConsoleAccess                   string
	_mqCreatorRequestId                string
	_mqData                            string
	_mqDataReplicationMode             string
	_mqDataReplicationPrimaryBrokerArn string
	_mqDeploymentMode                  string
	_mqDescription                     string
	_mqEncryptionOptions               string
	_mqEngineType                      string
	_mqEngineVersion                   string
	_mqGroups                          []string
	_mqHostInstanceType                string
	_mqLdapServerMetadata              string
	_mqLogs                            string
	_mqMaintenanceWindowStartTime      string
	_mqMaxResults                      string
	_mqMode                            string
	_mqName                            string
	_mqNextToken                       string
	_mqPassword                        string
	_mqPubliclyAccessible              string
	_mqReplicationUser                 string
	_mqResourceArn                     string
	_mqSecurityGroups                  []string
	_mqStorageType                     string
	_mqSubnetIds                       []string
	_mqTagKeys                         []string
	_mqTags                            string
	_mqUsername                        string
	_mqUsers                           string
)

// Creates a broker. Note: This API is asynchronous.
// To create a broker, you must either use the AmazonMQFullAccess IAM policy or
// include the following EC2 permissions in your IAM policy.
//
// - ec2:CreateNetworkInterface
//
// # This permission is required to allow Amazon MQ to create an elastic network
//
// interface (ENI) on behalf of your account.
//
// - ec2:CreateNetworkInterfacePermission
//
// This permission is required to attach the ENI to the broker instance.
//
// - ec2:DeleteNetworkInterface
//
// - ec2:DeleteNetworkInterfacePermission
//
// - ec2:DetachNetworkInterface
//
// - ec2:DescribeInternetGateways
//
// - ec2:DescribeNetworkInterfaces
//
// - ec2:DescribeNetworkInterfacePermissions
//
// - ec2:DescribeRouteTables
//
// - ec2:DescribeSecurityGroups
//
// - ec2:DescribeSubnets
//
// - ec2:DescribeVpcs
//
// For more information, see [Create an IAM User and Get Your Amazon Web Services Credentials] and [Never Modify or Delete the Amazon MQ Elastic Network Interface] in the Amazon MQ Developer Guide.
//
// [Never Modify or Delete the Amazon MQ Elastic Network Interface]: https://docs.aws.amazon.com//amazon-mq/latest/developer-guide/connecting-to-amazon-mq.html#never-modify-delete-elastic-network-interface
// [Create an IAM User and Get Your Amazon Web Services Credentials]: https://docs.aws.amazon.com//amazon-mq/latest/developer-guide/amazon-mq-setting-up.html#create-iam-user
func mq_CreateBroker(cfg aws.Config, client *mq.Client) {
	input := &mq.CreateBrokerInput{
		// BrokerName: *string, // Required
		// DeploymentMode: types.DeploymentMode, // Required
		// EngineType: types.EngineType, // Required
		// HostInstanceType: *string, // Required
		// PubliclyAccessible: *bool, // Required
	}

	if len(_mqBrokerName) > 0 {
		input.BrokerName = aws.String(_mqBrokerName)
	}
	if len(_mqDeploymentMode) > 0 {
		if err := assignInputField(input, "DeploymentMode", _mqDeploymentMode); err != nil {
			log.Errorf("invalid --deployment-mode: %s", err.Error())
			return
		}
	}
	if len(_mqEngineType) > 0 {
		if err := assignInputField(input, "EngineType", _mqEngineType); err != nil {
			log.Errorf("invalid --engine-type: %s", err.Error())
			return
		}
	}
	if len(_mqHostInstanceType) > 0 {
		input.HostInstanceType = aws.String(_mqHostInstanceType)
	}
	if len(_mqPubliclyAccessible) > 0 {
		if err := assignInputField(input, "PubliclyAccessible", _mqPubliclyAccessible); err != nil {
			log.Errorf("invalid --publicly-accessible: %s", err.Error())
			return
		}
	}
	if len(_mqAuthenticationStrategy) > 0 {
		if err := assignInputField(input, "AuthenticationStrategy", _mqAuthenticationStrategy); err != nil {
			log.Errorf("invalid --authentication-strategy: %s", err.Error())
			return
		}
	}
	if len(_mqAutoMinorVersionUpgrade) > 0 {
		if err := assignInputField(input, "AutoMinorVersionUpgrade", _mqAutoMinorVersionUpgrade); err != nil {
			log.Errorf("invalid --auto-minor-version-upgrade: %s", err.Error())
			return
		}
	}
	if len(_mqConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _mqConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_mqCreatorRequestId) > 0 {
		input.CreatorRequestId = aws.String(_mqCreatorRequestId)
	}
	if len(_mqDataReplicationMode) > 0 {
		if err := assignInputField(input, "DataReplicationMode", _mqDataReplicationMode); err != nil {
			log.Errorf("invalid --data-replication-mode: %s", err.Error())
			return
		}
	}
	if len(_mqDataReplicationPrimaryBrokerArn) > 0 {
		input.DataReplicationPrimaryBrokerArn = aws.String(_mqDataReplicationPrimaryBrokerArn)
	}
	if len(_mqEncryptionOptions) > 0 {
		if err := assignInputField(input, "EncryptionOptions", _mqEncryptionOptions); err != nil {
			log.Errorf("invalid --encryption-options: %s", err.Error())
			return
		}
	}
	if len(_mqEngineVersion) > 0 {
		input.EngineVersion = aws.String(_mqEngineVersion)
	}
	if len(_mqLdapServerMetadata) > 0 {
		if err := assignInputField(input, "LdapServerMetadata", _mqLdapServerMetadata); err != nil {
			log.Errorf("invalid --ldap-server-metadata: %s", err.Error())
			return
		}
	}
	if len(_mqLogs) > 0 {
		if err := assignInputField(input, "Logs", _mqLogs); err != nil {
			log.Errorf("invalid --logs: %s", err.Error())
			return
		}
	}
	if len(_mqMaintenanceWindowStartTime) > 0 {
		if err := assignInputField(input, "MaintenanceWindowStartTime", _mqMaintenanceWindowStartTime); err != nil {
			log.Errorf("invalid --maintenance-window-start-time: %s", err.Error())
			return
		}
	}
	if len(_mqSecurityGroups) > 0 {
		input.SecurityGroups = append([]string(nil), _mqSecurityGroups...)
	}
	if len(_mqStorageType) > 0 {
		if err := assignInputField(input, "StorageType", _mqStorageType); err != nil {
			log.Errorf("invalid --storage-type: %s", err.Error())
			return
		}
	}
	if len(_mqSubnetIds) > 0 {
		input.SubnetIds = append([]string(nil), _mqSubnetIds...)
	}
	if len(_mqTags) > 0 {
		if err := assignInputField(input, "Tags", _mqTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_mqUsers) > 0 {
		if err := assignInputField(input, "Users", _mqUsers); err != nil {
			log.Errorf("invalid --users: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateBroker(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new configuration for the specified configuration name. Amazon MQ
// uses the default configuration (the engine type and version).
func mq_CreateConfiguration(cfg aws.Config, client *mq.Client) {
	input := &mq.CreateConfigurationInput{
		// EngineType: types.EngineType, // Required
		// Name: *string, // Required
	}

	if len(_mqEngineType) > 0 {
		if err := assignInputField(input, "EngineType", _mqEngineType); err != nil {
			log.Errorf("invalid --engine-type: %s", err.Error())
			return
		}
	}
	if len(_mqName) > 0 {
		input.Name = aws.String(_mqName)
	}
	if len(_mqAuthenticationStrategy) > 0 {
		if err := assignInputField(input, "AuthenticationStrategy", _mqAuthenticationStrategy); err != nil {
			log.Errorf("invalid --authentication-strategy: %s", err.Error())
			return
		}
	}
	if len(_mqEngineVersion) > 0 {
		input.EngineVersion = aws.String(_mqEngineVersion)
	}
	if len(_mqTags) > 0 {
		if err := assignInputField(input, "Tags", _mqTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Add a tag to a resource.
func mq_CreateTags(cfg aws.Config, client *mq.Client) {
	input := &mq.CreateTagsInput{
		// ResourceArn: *string, // Required
	}

	if len(_mqResourceArn) > 0 {
		input.ResourceArn = aws.String(_mqResourceArn)
	}
	if len(_mqTags) > 0 {
		if err := assignInputField(input, "Tags", _mqTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTags(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an ActiveMQ user.
// Do not add personally identifiable information (PII) or other confidential or
// sensitive information in broker usernames. Broker usernames are accessible to
// other Amazon Web Services services, including CloudWatch Logs. Broker usernames
// are not intended to be used for private or sensitive data.
func mq_CreateUser(cfg aws.Config, client *mq.Client) {
	input := &mq.CreateUserInput{
		// BrokerId: *string, // Required
		// Password: *string, // Required
		// Username: *string, // Required
	}

	if len(_mqBrokerId) > 0 {
		input.BrokerId = aws.String(_mqBrokerId)
	}
	if len(_mqPassword) > 0 {
		input.Password = aws.String(_mqPassword)
	}
	if len(_mqUsername) > 0 {
		input.Username = aws.String(_mqUsername)
	}
	if len(_mqConsoleAccess) > 0 {
		if err := assignInputField(input, "ConsoleAccess", _mqConsoleAccess); err != nil {
			log.Errorf("invalid --console-access: %s", err.Error())
			return
		}
	}
	if len(_mqGroups) > 0 {
		input.Groups = append([]string(nil), _mqGroups...)
	}
	if len(_mqReplicationUser) > 0 {
		if err := assignInputField(input, "ReplicationUser", _mqReplicationUser); err != nil {
			log.Errorf("invalid --replication-user: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a broker. Note: This API is asynchronous.
func mq_DeleteBroker(cfg aws.Config, client *mq.Client) {
	input := &mq.DeleteBrokerInput{
		// BrokerId: *string, // Required
	}

	if len(_mqBrokerId) > 0 {
		input.BrokerId = aws.String(_mqBrokerId)
	}

	if resp, err := client.DeleteBroker(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified configuration.
func mq_DeleteConfiguration(cfg aws.Config, client *mq.Client) {
	input := &mq.DeleteConfigurationInput{
		// ConfigurationId: *string, // Required
	}

	if len(_mqConfigurationId) > 0 {
		input.ConfigurationId = aws.String(_mqConfigurationId)
	}

	if resp, err := client.DeleteConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a tag from a resource.
func mq_DeleteTags(cfg aws.Config, client *mq.Client) {
	input := &mq.DeleteTagsInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_mqResourceArn) > 0 {
		input.ResourceArn = aws.String(_mqResourceArn)
	}
	if len(_mqTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _mqTagKeys...)
	}

	if resp, err := client.DeleteTags(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an ActiveMQ user.
func mq_DeleteUser(cfg aws.Config, client *mq.Client) {
	input := &mq.DeleteUserInput{
		// BrokerId: *string, // Required
		// Username: *string, // Required
	}

	if len(_mqBrokerId) > 0 {
		input.BrokerId = aws.String(_mqBrokerId)
	}
	if len(_mqUsername) > 0 {
		input.Username = aws.String(_mqUsername)
	}

	if resp, err := client.DeleteUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the specified broker.
func mq_DescribeBroker(cfg aws.Config, client *mq.Client) {
	input := &mq.DescribeBrokerInput{
		// BrokerId: *string, // Required
	}

	if len(_mqBrokerId) > 0 {
		input.BrokerId = aws.String(_mqBrokerId)
	}

	if resp, err := client.DescribeBroker(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describe available engine types and versions.
func mq_DescribeBrokerEngineTypes(cfg aws.Config, client *mq.Client) {
	input := &mq.DescribeBrokerEngineTypesInput{}

	if len(_mqEngineType) > 0 {
		input.EngineType = aws.String(_mqEngineType)
	}
	if len(_mqMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mqMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mqNextToken) > 0 {
		input.NextToken = aws.String(_mqNextToken)
	}

	if resp, err := client.DescribeBrokerEngineTypes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describe available broker instance options.
func mq_DescribeBrokerInstanceOptions(cfg aws.Config, client *mq.Client) {
	input := &mq.DescribeBrokerInstanceOptionsInput{}

	if len(_mqEngineType) > 0 {
		input.EngineType = aws.String(_mqEngineType)
	}
	if len(_mqHostInstanceType) > 0 {
		input.HostInstanceType = aws.String(_mqHostInstanceType)
	}
	if len(_mqMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mqMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mqNextToken) > 0 {
		input.NextToken = aws.String(_mqNextToken)
	}
	if len(_mqStorageType) > 0 {
		input.StorageType = aws.String(_mqStorageType)
	}

	if resp, err := client.DescribeBrokerInstanceOptions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the specified configuration.
func mq_DescribeConfiguration(cfg aws.Config, client *mq.Client) {
	input := &mq.DescribeConfigurationInput{
		// ConfigurationId: *string, // Required
	}

	if len(_mqConfigurationId) > 0 {
		input.ConfigurationId = aws.String(_mqConfigurationId)
	}

	if resp, err := client.DescribeConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the specified configuration revision for the specified configuration.
func mq_DescribeConfigurationRevision(cfg aws.Config, client *mq.Client) {
	input := &mq.DescribeConfigurationRevisionInput{
		// ConfigurationId: *string, // Required
		// ConfigurationRevision: *string, // Required
	}

	if len(_mqConfigurationId) > 0 {
		input.ConfigurationId = aws.String(_mqConfigurationId)
	}
	if len(_mqConfigurationRevision) > 0 {
		input.ConfigurationRevision = aws.String(_mqConfigurationRevision)
	}

	if resp, err := client.DescribeConfigurationRevision(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about an ActiveMQ user.
func mq_DescribeUser(cfg aws.Config, client *mq.Client) {
	input := &mq.DescribeUserInput{
		// BrokerId: *string, // Required
		// Username: *string, // Required
	}

	if len(_mqBrokerId) > 0 {
		input.BrokerId = aws.String(_mqBrokerId)
	}
	if len(_mqUsername) > 0 {
		input.Username = aws.String(_mqUsername)
	}

	if resp, err := client.DescribeUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of all brokers.
func mq_ListBrokers(cfg aws.Config, client *mq.Client) {
	input := &mq.ListBrokersInput{}

	if len(_mqMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mqMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mqNextToken) > 0 {
		input.NextToken = aws.String(_mqNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListBrokers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mq.ListBrokersOutput
	p := mq.NewListBrokersPaginator(client, input)
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

// Returns a list of all revisions for the specified configuration.
func mq_ListConfigurationRevisions(cfg aws.Config, client *mq.Client) {
	input := &mq.ListConfigurationRevisionsInput{
		// ConfigurationId: *string, // Required
	}

	if len(_mqConfigurationId) > 0 {
		input.ConfigurationId = aws.String(_mqConfigurationId)
	}
	if len(_mqMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mqMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mqNextToken) > 0 {
		input.NextToken = aws.String(_mqNextToken)
	}

	if resp, err := client.ListConfigurationRevisions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of all configurations.
func mq_ListConfigurations(cfg aws.Config, client *mq.Client) {
	input := &mq.ListConfigurationsInput{}

	if len(_mqMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mqMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mqNextToken) > 0 {
		input.NextToken = aws.String(_mqNextToken)
	}

	if resp, err := client.ListConfigurations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists tags for a resource.
func mq_ListTags(cfg aws.Config, client *mq.Client) {
	input := &mq.ListTagsInput{
		// ResourceArn: *string, // Required
	}

	if len(_mqResourceArn) > 0 {
		input.ResourceArn = aws.String(_mqResourceArn)
	}

	if resp, err := client.ListTags(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of all ActiveMQ users.
func mq_ListUsers(cfg aws.Config, client *mq.Client) {
	input := &mq.ListUsersInput{
		// BrokerId: *string, // Required
	}

	if len(_mqBrokerId) > 0 {
		input.BrokerId = aws.String(_mqBrokerId)
	}
	if len(_mqMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mqMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mqNextToken) > 0 {
		input.NextToken = aws.String(_mqNextToken)
	}

	if resp, err := client.ListUsers(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Promotes a data replication replica broker to the primary broker role.
func mq_Promote(cfg aws.Config, client *mq.Client) {
	input := &mq.PromoteInput{
		// BrokerId: *string, // Required
		// Mode: types.PromoteMode, // Required
	}

	if len(_mqBrokerId) > 0 {
		input.BrokerId = aws.String(_mqBrokerId)
	}
	if len(_mqMode) > 0 {
		if err := assignInputField(input, "Mode", _mqMode); err != nil {
			log.Errorf("invalid --mode: %s", err.Error())
			return
		}
	}

	if resp, err := client.Promote(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Reboots a broker. Note: This API is asynchronous.
func mq_RebootBroker(cfg aws.Config, client *mq.Client) {
	input := &mq.RebootBrokerInput{
		// BrokerId: *string, // Required
	}

	if len(_mqBrokerId) > 0 {
		input.BrokerId = aws.String(_mqBrokerId)
	}

	if resp, err := client.RebootBroker(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a pending configuration change to a broker.
func mq_UpdateBroker(cfg aws.Config, client *mq.Client) {
	input := &mq.UpdateBrokerInput{
		// BrokerId: *string, // Required
	}

	if len(_mqBrokerId) > 0 {
		input.BrokerId = aws.String(_mqBrokerId)
	}
	if len(_mqAuthenticationStrategy) > 0 {
		if err := assignInputField(input, "AuthenticationStrategy", _mqAuthenticationStrategy); err != nil {
			log.Errorf("invalid --authentication-strategy: %s", err.Error())
			return
		}
	}
	if len(_mqAutoMinorVersionUpgrade) > 0 {
		if err := assignInputField(input, "AutoMinorVersionUpgrade", _mqAutoMinorVersionUpgrade); err != nil {
			log.Errorf("invalid --auto-minor-version-upgrade: %s", err.Error())
			return
		}
	}
	if len(_mqConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _mqConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_mqDataReplicationMode) > 0 {
		if err := assignInputField(input, "DataReplicationMode", _mqDataReplicationMode); err != nil {
			log.Errorf("invalid --data-replication-mode: %s", err.Error())
			return
		}
	}
	if len(_mqEngineVersion) > 0 {
		input.EngineVersion = aws.String(_mqEngineVersion)
	}
	if len(_mqHostInstanceType) > 0 {
		input.HostInstanceType = aws.String(_mqHostInstanceType)
	}
	if len(_mqLdapServerMetadata) > 0 {
		if err := assignInputField(input, "LdapServerMetadata", _mqLdapServerMetadata); err != nil {
			log.Errorf("invalid --ldap-server-metadata: %s", err.Error())
			return
		}
	}
	if len(_mqLogs) > 0 {
		if err := assignInputField(input, "Logs", _mqLogs); err != nil {
			log.Errorf("invalid --logs: %s", err.Error())
			return
		}
	}
	if len(_mqMaintenanceWindowStartTime) > 0 {
		if err := assignInputField(input, "MaintenanceWindowStartTime", _mqMaintenanceWindowStartTime); err != nil {
			log.Errorf("invalid --maintenance-window-start-time: %s", err.Error())
			return
		}
	}
	if len(_mqSecurityGroups) > 0 {
		input.SecurityGroups = append([]string(nil), _mqSecurityGroups...)
	}

	if resp, err := client.UpdateBroker(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified configuration.
func mq_UpdateConfiguration(cfg aws.Config, client *mq.Client) {
	input := &mq.UpdateConfigurationInput{
		// ConfigurationId: *string, // Required
		// Data: *string, // Required
	}

	if len(_mqConfigurationId) > 0 {
		input.ConfigurationId = aws.String(_mqConfigurationId)
	}
	if len(_mqData) > 0 {
		input.Data = aws.String(_mqData)
	}
	if len(_mqDescription) > 0 {
		input.Description = aws.String(_mqDescription)
	}

	if resp, err := client.UpdateConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the information for an ActiveMQ user.
func mq_UpdateUser(cfg aws.Config, client *mq.Client) {
	input := &mq.UpdateUserInput{
		// BrokerId: *string, // Required
		// Username: *string, // Required
	}

	if len(_mqBrokerId) > 0 {
		input.BrokerId = aws.String(_mqBrokerId)
	}
	if len(_mqUsername) > 0 {
		input.Username = aws.String(_mqUsername)
	}
	if len(_mqConsoleAccess) > 0 {
		if err := assignInputField(input, "ConsoleAccess", _mqConsoleAccess); err != nil {
			log.Errorf("invalid --console-access: %s", err.Error())
			return
		}
	}
	if len(_mqGroups) > 0 {
		input.Groups = append([]string(nil), _mqGroups...)
	}
	if len(_mqPassword) > 0 {
		input.Password = aws.String(_mqPassword)
	}
	if len(_mqReplicationUser) > 0 {
		if err := assignInputField(input, "ReplicationUser", _mqReplicationUser); err != nil {
			log.Errorf("invalid --replication-user: %s", err.Error())
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
	_rootCmd.AddCommand(_mqCmd)
	_mqCmd.Flags().SortFlags = false

	_mqCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_mqCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_mqCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_mqCmd.Flags().StringVarP(&_mqAuthenticationStrategy, "authentication-strategy", "", "", "Authentication Strategy")
	_mqCmd.Flags().StringVarP(&_mqAutoMinorVersionUpgrade, "auto-minor-version-upgrade", "", "", "Auto Minor Version Upgrade")
	_mqCmd.Flags().StringVarP(&_mqBrokerId, "broker-id", "", "", "Broker ID")
	_mqCmd.Flags().StringVarP(&_mqBrokerName, "broker-name", "", "", "Broker Name")
	_mqCmd.Flags().StringVarP(&_mqConfiguration, "configuration", "", "", "Configuration")
	_mqCmd.Flags().StringVarP(&_mqConfigurationId, "configuration-id", "", "", "Configuration ID")
	_mqCmd.Flags().StringVarP(&_mqConfigurationRevision, "configuration-revision", "", "", "Configuration Revision")
	_mqCmd.Flags().StringVarP(&_mqConsoleAccess, "console-access", "", "", "Console Access")
	_mqCmd.Flags().StringVarP(&_mqCreatorRequestId, "creator-request-id", "", "", "Creator Request ID")
	_mqCmd.Flags().StringVarP(&_mqData, "data", "", "", "Data")
	_mqCmd.Flags().StringVarP(&_mqDataReplicationMode, "data-replication-mode", "", "", "Data Replication Mode")
	_mqCmd.Flags().StringVarP(&_mqDataReplicationPrimaryBrokerArn, "data-replication-primary-broker-arn", "", "", "Data Replication Primary Broker ARN")
	_mqCmd.Flags().StringVarP(&_mqDeploymentMode, "deployment-mode", "", "", "Deployment Mode")
	_mqCmd.Flags().StringVarP(&_mqDescription, "description", "", "", "Description")
	_mqCmd.Flags().StringVarP(&_mqEncryptionOptions, "encryption-options", "", "", "Encryption Options")
	_mqCmd.Flags().StringVarP(&_mqEngineType, "engine-type", "", "", "Engine Type")
	_mqCmd.Flags().StringVarP(&_mqEngineVersion, "engine-version", "", "", "Engine Version")
	_mqCmd.Flags().StringSliceVarP(&_mqGroups, "groups", "", nil, "Groups")
	_mqCmd.Flags().StringVarP(&_mqHostInstanceType, "host-instance-type", "", "", "Host Instance Type")
	_mqCmd.Flags().StringVarP(&_mqLdapServerMetadata, "ldap-server-metadata", "", "", "Ldap Server Metadata")
	_mqCmd.Flags().StringVarP(&_mqLogs, "logs", "", "", "Logs")
	_mqCmd.Flags().StringVarP(&_mqMaintenanceWindowStartTime, "maintenance-window-start-time", "", "", "Maintenance Window Start Time")
	_mqCmd.Flags().StringVarP(&_mqMaxResults, "max-results", "", "", "Max Results")
	_mqCmd.Flags().StringVarP(&_mqMode, "mode", "", "", "Mode")
	_mqCmd.Flags().StringVarP(&_mqName, "name", "", "", "Name")
	_mqCmd.Flags().StringVarP(&_mqNextToken, "next-token", "", "", "Next Token")
	_mqCmd.Flags().StringVarP(&_mqPassword, "password", "", "", "Password")
	_mqCmd.Flags().StringVarP(&_mqPubliclyAccessible, "publicly-accessible", "", "", "Publicly Accessible")
	_mqCmd.Flags().StringVarP(&_mqReplicationUser, "replication-user", "", "", "Replication User")
	_mqCmd.Flags().StringVarP(&_mqResourceArn, "resource-arn", "", "", "Resource ARN")
	_mqCmd.Flags().StringSliceVarP(&_mqSecurityGroups, "security-groups", "", nil, "Security Groups")
	_mqCmd.Flags().StringVarP(&_mqStorageType, "storage-type", "", "", "Storage Type")
	_mqCmd.Flags().StringSliceVarP(&_mqSubnetIds, "subnet-ids", "", nil, "Subnet Ids")
	_mqCmd.Flags().StringSliceVarP(&_mqTagKeys, "tag-keys", "", nil, "Tag Keys")
	_mqCmd.Flags().StringVarP(&_mqTags, "tags", "", "", "Tags")
	_mqCmd.Flags().StringVarP(&_mqUsername, "username", "", "", "Username")
	_mqCmd.Flags().StringVarP(&_mqUsers, "users", "", "", "Users")

	_mqCmd.Flags().BoolVarP(&_mqCreateBroker, "create-broker", "", false, "Create Broker")
	_mqCmd.Flags().BoolVarP(&_mqCreateConfiguration, "create-configuration", "", false, "Create Configuration")
	_mqCmd.Flags().BoolVarP(&_mqCreateTags, "create-tags", "", false, "Create Tags")
	_mqCmd.Flags().BoolVarP(&_mqCreateUser, "create-user", "", false, "Create User")
	_mqCmd.Flags().BoolVarP(&_mqDeleteBroker, "delete-broker", "", false, "Delete Broker")
	_mqCmd.Flags().BoolVarP(&_mqDeleteConfiguration, "delete-configuration", "", false, "Delete Configuration")
	_mqCmd.Flags().BoolVarP(&_mqDeleteTags, "delete-tags", "", false, "Delete Tags")
	_mqCmd.Flags().BoolVarP(&_mqDeleteUser, "delete-user", "", false, "Delete User")
	_mqCmd.Flags().BoolVarP(&_mqDescribeBroker, "describe-broker", "", false, "Describe Broker")
	_mqCmd.Flags().BoolVarP(&_mqDescribeBrokerEngineTypes, "describe-broker-engine-types", "", false, "Describe Broker Engine Types")
	_mqCmd.Flags().BoolVarP(&_mqDescribeBrokerInstanceOptions, "describe-broker-instance-options", "", false, "Describe Broker Instance Options")
	_mqCmd.Flags().BoolVarP(&_mqDescribeConfiguration, "describe-configuration", "", false, "Describe Configuration")
	_mqCmd.Flags().BoolVarP(&_mqDescribeConfigurationRevision, "describe-configuration-revision", "", false, "Describe Configuration Revision")
	_mqCmd.Flags().BoolVarP(&_mqDescribeUser, "describe-user", "", false, "Describe User")
	_mqCmd.Flags().BoolVarP(&_mqListBrokers, "list-brokers", "", false, "List Brokers")
	_mqCmd.Flags().BoolVarP(&_mqListConfigurationRevisions, "list-configuration-revisions", "", false, "List Configuration Revisions")
	_mqCmd.Flags().BoolVarP(&_mqListConfigurations, "list-configurations", "", false, "List Configurations")
	_mqCmd.Flags().BoolVarP(&_mqListTags, "list-tags", "", false, "List Tags")
	_mqCmd.Flags().BoolVarP(&_mqListUsers, "list-users", "", false, "List Users")
	_mqCmd.Flags().BoolVarP(&_mqPromote, "promote", "", false, "Promote")
	_mqCmd.Flags().BoolVarP(&_mqRebootBroker, "reboot-broker", "", false, "Reboot Broker")
	_mqCmd.Flags().BoolVarP(&_mqUpdateBroker, "update-broker", "", false, "Update Broker")
	_mqCmd.Flags().BoolVarP(&_mqUpdateConfiguration, "update-configuration", "", false, "Update Configuration")
	_mqCmd.Flags().BoolVarP(&_mqUpdateUser, "update-user", "", false, "Update User")

}
