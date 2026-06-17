package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloud9"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// cloud9Cmd represents the cloud9 command
var _cloud9Cmd = &cobra.Command{
	Use:   "cloud9",
	Short: "AWS cloud9 CLI",
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
		client := cloud9.NewFromConfig(cfg)
		if _cloud9CreateEnvironmentEC2 {
			cloud9_CreateEnvironmentEC2(cfg, client)
			return
		}
		if _cloud9CreateEnvironmentMembership {
			cloud9_CreateEnvironmentMembership(cfg, client)
			return
		}
		if _cloud9DeleteEnvironment {
			cloud9_DeleteEnvironment(cfg, client)
			return
		}
		if _cloud9DeleteEnvironmentMembership {
			cloud9_DeleteEnvironmentMembership(cfg, client)
			return
		}
		if _cloud9DescribeEnvironmentMemberships {
			cloud9_DescribeEnvironmentMemberships(cfg, client)
			return
		}
		if _cloud9DescribeEnvironmentStatus {
			cloud9_DescribeEnvironmentStatus(cfg, client)
			return
		}
		if _cloud9DescribeEnvironments {
			cloud9_DescribeEnvironments(cfg, client)
			return
		}
		if _cloud9ListEnvironments {
			cloud9_ListEnvironments(cfg, client)
			return
		}
		if _cloud9ListTagsForResource {
			cloud9_ListTagsForResource(cfg, client)
			return
		}
		if _cloud9TagResource {
			cloud9_TagResource(cfg, client)
			return
		}
		if _cloud9UntagResource {
			cloud9_UntagResource(cfg, client)
			return
		}
		if _cloud9UpdateEnvironment {
			cloud9_UpdateEnvironment(cfg, client)
			return
		}
		if _cloud9UpdateEnvironmentMembership {
			cloud9_UpdateEnvironmentMembership(cfg, client)
			return
		}

	},
}

var (
	_cloud9CreateEnvironmentEC2           bool
	_cloud9CreateEnvironmentMembership    bool
	_cloud9DeleteEnvironment              bool
	_cloud9DeleteEnvironmentMembership    bool
	_cloud9DescribeEnvironmentMemberships bool
	_cloud9DescribeEnvironmentStatus      bool
	_cloud9DescribeEnvironments           bool
	_cloud9ListEnvironments               bool
	_cloud9ListTagsForResource            bool
	_cloud9TagResource                    bool
	_cloud9UntagResource                  bool
	_cloud9UpdateEnvironment              bool
	_cloud9UpdateEnvironmentMembership    bool

	_cloud9AutomaticStopTimeMinutes string
	_cloud9ClientRequestToken       string
	_cloud9ConnectionType           string
	_cloud9Description              string
	_cloud9DryRun                   string
	_cloud9EnvironmentId            string
	_cloud9EnvironmentIds           []string
	_cloud9ImageId                  string
	_cloud9InstanceType             string
	_cloud9ManagedCredentialsAction string
	_cloud9MaxResults               string
	_cloud9Name                     string
	_cloud9NextToken                string
	_cloud9OwnerArn                 string
	_cloud9Permissions              string
	_cloud9ResourceARN              string
	_cloud9SubnetId                 string
	_cloud9TagKeys                  []string
	_cloud9Tags                     string
	_cloud9UserArn                  string
)

// Creates an Cloud9 development environment, launches an Amazon Elastic Compute
// Cloud (Amazon EC2) instance, and then connects from the instance to the
// environment.
//
// Cloud9 is no longer available to new customers. Existing customers of Cloud9
// can continue to use the service as normal. [Learn more"]
//
// [Learn more"]: http://aws.amazon.com/blogs/devops/how-to-migrate-from-aws-cloud9-to-aws-ide-toolkits-or-aws-cloudshell/
func cloud9_CreateEnvironmentEC2(cfg aws.Config, client *cloud9.Client) {
	input := &cloud9.CreateEnvironmentEC2Input{
		// ImageId: *string, // Required
		// InstanceType: *string, // Required
		// Name: *string, // Required
	}

	if len(_cloud9ImageId) > 0 {
		input.ImageId = aws.String(_cloud9ImageId)
	}
	if len(_cloud9InstanceType) > 0 {
		input.InstanceType = aws.String(_cloud9InstanceType)
	}
	if len(_cloud9Name) > 0 {
		input.Name = aws.String(_cloud9Name)
	}
	if len(_cloud9AutomaticStopTimeMinutes) > 0 {
		if err := assignInputField(input, "AutomaticStopTimeMinutes", _cloud9AutomaticStopTimeMinutes); err != nil {
			log.Errorf("invalid --automatic-stop-time-minutes: %s", err.Error())
			return
		}
	}
	if len(_cloud9ClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_cloud9ClientRequestToken)
	}
	if len(_cloud9ConnectionType) > 0 {
		if err := assignInputField(input, "ConnectionType", _cloud9ConnectionType); err != nil {
			log.Errorf("invalid --connection-type: %s", err.Error())
			return
		}
	}
	if len(_cloud9Description) > 0 {
		input.Description = aws.String(_cloud9Description)
	}
	if len(_cloud9DryRun) > 0 {
		if err := assignInputField(input, "DryRun", _cloud9DryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}
	if len(_cloud9OwnerArn) > 0 {
		input.OwnerArn = aws.String(_cloud9OwnerArn)
	}
	if len(_cloud9SubnetId) > 0 {
		input.SubnetId = aws.String(_cloud9SubnetId)
	}
	if len(_cloud9Tags) > 0 {
		if err := assignInputField(input, "Tags", _cloud9Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateEnvironmentEC2(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds an environment member to an Cloud9 development environment.
// Cloud9 is no longer available to new customers. Existing customers of Cloud9
// can continue to use the service as normal. [Learn more"]
//
// [Learn more"]: http://aws.amazon.com/blogs/devops/how-to-migrate-from-aws-cloud9-to-aws-ide-toolkits-or-aws-cloudshell/
func cloud9_CreateEnvironmentMembership(cfg aws.Config, client *cloud9.Client) {
	input := &cloud9.CreateEnvironmentMembershipInput{
		// EnvironmentId: *string, // Required
		// Permissions: types.MemberPermissions, // Required
		// UserArn: *string, // Required
	}

	if len(_cloud9EnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_cloud9EnvironmentId)
	}
	if len(_cloud9Permissions) > 0 {
		if err := assignInputField(input, "Permissions", _cloud9Permissions); err != nil {
			log.Errorf("invalid --permissions: %s", err.Error())
			return
		}
	}
	if len(_cloud9UserArn) > 0 {
		input.UserArn = aws.String(_cloud9UserArn)
	}

	if resp, err := client.CreateEnvironmentMembership(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Cloud9 development environment. If an Amazon EC2 instance is
// connected to the environment, also terminates the instance.
//
// Cloud9 is no longer available to new customers. Existing customers of Cloud9
// can continue to use the service as normal. [Learn more"]
//
// [Learn more"]: http://aws.amazon.com/blogs/devops/how-to-migrate-from-aws-cloud9-to-aws-ide-toolkits-or-aws-cloudshell/
func cloud9_DeleteEnvironment(cfg aws.Config, client *cloud9.Client) {
	input := &cloud9.DeleteEnvironmentInput{
		// EnvironmentId: *string, // Required
	}

	if len(_cloud9EnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_cloud9EnvironmentId)
	}

	if resp, err := client.DeleteEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an environment member from a development environment.
// Cloud9 is no longer available to new customers. Existing customers of Cloud9
// can continue to use the service as normal. [Learn more"]
//
// [Learn more"]: http://aws.amazon.com/blogs/devops/how-to-migrate-from-aws-cloud9-to-aws-ide-toolkits-or-aws-cloudshell/
func cloud9_DeleteEnvironmentMembership(cfg aws.Config, client *cloud9.Client) {
	input := &cloud9.DeleteEnvironmentMembershipInput{
		// EnvironmentId: *string, // Required
		// UserArn: *string, // Required
	}

	if len(_cloud9EnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_cloud9EnvironmentId)
	}
	if len(_cloud9UserArn) > 0 {
		input.UserArn = aws.String(_cloud9UserArn)
	}

	if resp, err := client.DeleteEnvironmentMembership(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about environment members for an Cloud9 development
// environment.
//
// Cloud9 is no longer available to new customers. Existing customers of Cloud9
// can continue to use the service as normal. [Learn more"]
//
// [Learn more"]: http://aws.amazon.com/blogs/devops/how-to-migrate-from-aws-cloud9-to-aws-ide-toolkits-or-aws-cloudshell/
func cloud9_DescribeEnvironmentMemberships(cfg aws.Config, client *cloud9.Client) {
	input := &cloud9.DescribeEnvironmentMembershipsInput{}

	if len(_cloud9EnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_cloud9EnvironmentId)
	}
	if len(_cloud9MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cloud9MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cloud9NextToken) > 0 {
		input.NextToken = aws.String(_cloud9NextToken)
	}
	if len(_cloud9Permissions) > 0 {
		if err := assignInputField(input, "Permissions", _cloud9Permissions); err != nil {
			log.Errorf("invalid --permissions: %s", err.Error())
			return
		}
	}
	if len(_cloud9UserArn) > 0 {
		input.UserArn = aws.String(_cloud9UserArn)
	}

	if disablePaginator() {
		if resp, err := client.DescribeEnvironmentMemberships(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloud9.DescribeEnvironmentMembershipsOutput
	p := cloud9.NewDescribeEnvironmentMembershipsPaginator(client, input)
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

// Gets status information for an Cloud9 development environment.
// Cloud9 is no longer available to new customers. Existing customers of Cloud9
// can continue to use the service as normal. [Learn more"]
//
// [Learn more"]: http://aws.amazon.com/blogs/devops/how-to-migrate-from-aws-cloud9-to-aws-ide-toolkits-or-aws-cloudshell/
func cloud9_DescribeEnvironmentStatus(cfg aws.Config, client *cloud9.Client) {
	input := &cloud9.DescribeEnvironmentStatusInput{
		// EnvironmentId: *string, // Required
	}

	if len(_cloud9EnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_cloud9EnvironmentId)
	}

	if resp, err := client.DescribeEnvironmentStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about Cloud9 development environments.
// Cloud9 is no longer available to new customers. Existing customers of Cloud9
// can continue to use the service as normal. [Learn more"]
//
// [Learn more"]: http://aws.amazon.com/blogs/devops/how-to-migrate-from-aws-cloud9-to-aws-ide-toolkits-or-aws-cloudshell/
func cloud9_DescribeEnvironments(cfg aws.Config, client *cloud9.Client) {
	input := &cloud9.DescribeEnvironmentsInput{
		// EnvironmentIds: []string, // Required
	}

	if len(_cloud9EnvironmentIds) > 0 {
		input.EnvironmentIds = append([]string(nil), _cloud9EnvironmentIds...)
	}

	if resp, err := client.DescribeEnvironments(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a list of Cloud9 development environment identifiers.
// Cloud9 is no longer available to new customers. Existing customers of Cloud9
// can continue to use the service as normal. [Learn more"]
//
// Cloud9 is no longer available to new customers. Existing customers of Cloud9
// can continue to use the service as normal. [Learn more"]
//
// [Learn more"]: http://aws.amazon.com/blogs/devops/how-to-migrate-from-aws-cloud9-to-aws-ide-toolkits-or-aws-cloudshell/
func cloud9_ListEnvironments(cfg aws.Config, client *cloud9.Client) {
	input := &cloud9.ListEnvironmentsInput{}

	if len(_cloud9MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cloud9MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cloud9NextToken) > 0 {
		input.NextToken = aws.String(_cloud9NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEnvironments(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloud9.ListEnvironmentsOutput
	p := cloud9.NewListEnvironmentsPaginator(client, input)
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

// Gets a list of the tags associated with an Cloud9 development environment.
// Cloud9 is no longer available to new customers. Existing customers of Cloud9
// can continue to use the service as normal. [Learn more"]
//
// [Learn more"]: http://aws.amazon.com/blogs/devops/how-to-migrate-from-aws-cloud9-to-aws-ide-toolkits-or-aws-cloudshell/
func cloud9_ListTagsForResource(cfg aws.Config, client *cloud9.Client) {
	input := &cloud9.ListTagsForResourceInput{
		// ResourceARN: *string, // Required
	}

	if len(_cloud9ResourceARN) > 0 {
		input.ResourceARN = aws.String(_cloud9ResourceARN)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds tags to an Cloud9 development environment.
// Cloud9 is no longer available to new customers. Existing customers of Cloud9
// can continue to use the service as normal. [Learn more"]
//
// Tags that you add to an Cloud9 environment by using this method will NOT be
// automatically propagated to underlying resources.
//
// [Learn more"]: http://aws.amazon.com/blogs/devops/how-to-migrate-from-aws-cloud9-to-aws-ide-toolkits-or-aws-cloudshell/
func cloud9_TagResource(cfg aws.Config, client *cloud9.Client) {
	input := &cloud9.TagResourceInput{
		// ResourceARN: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_cloud9ResourceARN) > 0 {
		input.ResourceARN = aws.String(_cloud9ResourceARN)
	}
	if len(_cloud9Tags) > 0 {
		if err := assignInputField(input, "Tags", _cloud9Tags); err != nil {
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

// Removes tags from an Cloud9 development environment.
// Cloud9 is no longer available to new customers. Existing customers of Cloud9
// can continue to use the service as normal. [Learn more"]
//
// [Learn more"]: http://aws.amazon.com/blogs/devops/how-to-migrate-from-aws-cloud9-to-aws-ide-toolkits-or-aws-cloudshell/
func cloud9_UntagResource(cfg aws.Config, client *cloud9.Client) {
	input := &cloud9.UntagResourceInput{
		// ResourceARN: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_cloud9ResourceARN) > 0 {
		input.ResourceARN = aws.String(_cloud9ResourceARN)
	}
	if len(_cloud9TagKeys) > 0 {
		input.TagKeys = append([]string(nil), _cloud9TagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes the settings of an existing Cloud9 development environment.
// Cloud9 is no longer available to new customers. Existing customers of Cloud9
// can continue to use the service as normal. [Learn more"]
//
// [Learn more"]: http://aws.amazon.com/blogs/devops/how-to-migrate-from-aws-cloud9-to-aws-ide-toolkits-or-aws-cloudshell/
func cloud9_UpdateEnvironment(cfg aws.Config, client *cloud9.Client) {
	input := &cloud9.UpdateEnvironmentInput{
		// EnvironmentId: *string, // Required
	}

	if len(_cloud9EnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_cloud9EnvironmentId)
	}
	if len(_cloud9Description) > 0 {
		input.Description = aws.String(_cloud9Description)
	}
	if len(_cloud9ManagedCredentialsAction) > 0 {
		if err := assignInputField(input, "ManagedCredentialsAction", _cloud9ManagedCredentialsAction); err != nil {
			log.Errorf("invalid --managed-credentials-action: %s", err.Error())
			return
		}
	}
	if len(_cloud9Name) > 0 {
		input.Name = aws.String(_cloud9Name)
	}

	if resp, err := client.UpdateEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes the settings of an existing environment member for an Cloud9
// development environment.
//
// Cloud9 is no longer available to new customers. Existing customers of Cloud9
// can continue to use the service as normal. [Learn more"]
//
// [Learn more"]: http://aws.amazon.com/blogs/devops/how-to-migrate-from-aws-cloud9-to-aws-ide-toolkits-or-aws-cloudshell/
func cloud9_UpdateEnvironmentMembership(cfg aws.Config, client *cloud9.Client) {
	input := &cloud9.UpdateEnvironmentMembershipInput{
		// EnvironmentId: *string, // Required
		// Permissions: types.MemberPermissions, // Required
		// UserArn: *string, // Required
	}

	if len(_cloud9EnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_cloud9EnvironmentId)
	}
	if len(_cloud9Permissions) > 0 {
		if err := assignInputField(input, "Permissions", _cloud9Permissions); err != nil {
			log.Errorf("invalid --permissions: %s", err.Error())
			return
		}
	}
	if len(_cloud9UserArn) > 0 {
		input.UserArn = aws.String(_cloud9UserArn)
	}

	if resp, err := client.UpdateEnvironmentMembership(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_cloud9Cmd)
	_cloud9Cmd.Flags().SortFlags = false

	_cloud9Cmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_cloud9Cmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_cloud9Cmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_cloud9Cmd.Flags().StringVarP(&_cloud9AutomaticStopTimeMinutes, "automatic-stop-time-minutes", "", "", "Automatic Stop Time Minutes")
	_cloud9Cmd.Flags().StringVarP(&_cloud9ClientRequestToken, "client-request-token", "", "", "Client Request Token")
	_cloud9Cmd.Flags().StringVarP(&_cloud9ConnectionType, "connection-type", "", "", "Connection Type")
	_cloud9Cmd.Flags().StringVarP(&_cloud9Description, "description", "", "", "Description")
	_cloud9Cmd.Flags().StringVarP(&_cloud9DryRun, "dry-run", "", "", "Dry Run")
	_cloud9Cmd.Flags().StringVarP(&_cloud9EnvironmentId, "environment-id", "", "", "Environment ID")
	_cloud9Cmd.Flags().StringSliceVarP(&_cloud9EnvironmentIds, "environment-ids", "", nil, "Environment Ids")
	_cloud9Cmd.Flags().StringVarP(&_cloud9ImageId, "image-id", "", "", "Image ID")
	_cloud9Cmd.Flags().StringVarP(&_cloud9InstanceType, "instance-type", "", "", "Instance Type")
	_cloud9Cmd.Flags().StringVarP(&_cloud9ManagedCredentialsAction, "managed-credentials-action", "", "", "Managed Credentials Action")
	_cloud9Cmd.Flags().StringVarP(&_cloud9MaxResults, "max-results", "", "", "Max Results")
	_cloud9Cmd.Flags().StringVarP(&_cloud9Name, "name", "", "", "Name")
	_cloud9Cmd.Flags().StringVarP(&_cloud9NextToken, "next-token", "", "", "Next Token")
	_cloud9Cmd.Flags().StringVarP(&_cloud9OwnerArn, "owner-arn", "", "", "Owner ARN")
	_cloud9Cmd.Flags().StringVarP(&_cloud9Permissions, "permissions", "", "", "Permissions")
	_cloud9Cmd.Flags().StringVarP(&_cloud9ResourceARN, "resource-arn", "", "", "Resource ARN")
	_cloud9Cmd.Flags().StringVarP(&_cloud9SubnetId, "subnet-id", "", "", "Subnet ID")
	_cloud9Cmd.Flags().StringSliceVarP(&_cloud9TagKeys, "tag-keys", "", nil, "Tag Keys")
	_cloud9Cmd.Flags().StringVarP(&_cloud9Tags, "tags", "", "", "Tags")
	_cloud9Cmd.Flags().StringVarP(&_cloud9UserArn, "user-arn", "", "", "User ARN")

	_cloud9Cmd.Flags().BoolVarP(&_cloud9CreateEnvironmentEC2, "create-environment-ec2", "", false, "Create Environment EC2")
	_cloud9Cmd.Flags().BoolVarP(&_cloud9CreateEnvironmentMembership, "create-environment-membership", "", false, "Create Environment Membership")
	_cloud9Cmd.Flags().BoolVarP(&_cloud9DeleteEnvironment, "delete-environment", "", false, "Delete Environment")
	_cloud9Cmd.Flags().BoolVarP(&_cloud9DeleteEnvironmentMembership, "delete-environment-membership", "", false, "Delete Environment Membership")
	_cloud9Cmd.Flags().BoolVarP(&_cloud9DescribeEnvironmentMemberships, "describe-environment-memberships", "", false, "Describe Environment Memberships")
	_cloud9Cmd.Flags().BoolVarP(&_cloud9DescribeEnvironmentStatus, "describe-environment-status", "", false, "Describe Environment Status")
	_cloud9Cmd.Flags().BoolVarP(&_cloud9DescribeEnvironments, "describe-environments", "", false, "Describe Environments")
	_cloud9Cmd.Flags().BoolVarP(&_cloud9ListEnvironments, "list-environments", "", false, "List Environments")
	_cloud9Cmd.Flags().BoolVarP(&_cloud9ListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_cloud9Cmd.Flags().BoolVarP(&_cloud9TagResource, "tag-resource", "", false, "Tag Resource")
	_cloud9Cmd.Flags().BoolVarP(&_cloud9UntagResource, "untag-resource", "", false, "Untag Resource")
	_cloud9Cmd.Flags().BoolVarP(&_cloud9UpdateEnvironment, "update-environment", "", false, "Update Environment")
	_cloud9Cmd.Flags().BoolVarP(&_cloud9UpdateEnvironmentMembership, "update-environment-membership", "", false, "Update Environment Membership")

}
