package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/repostspace"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// repostspaceCmd represents the repostspace command
var _repostspaceCmd = &cobra.Command{
	Use:   "repostspace",
	Short: "AWS repostspace CLI",
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
		client := repostspace.NewFromConfig(cfg)
		if _repostspaceBatchAddChannelRoleToAccessors {
			repostspace_BatchAddChannelRoleToAccessors(cfg, client)
			return
		}
		if _repostspaceBatchAddRole {
			repostspace_BatchAddRole(cfg, client)
			return
		}
		if _repostspaceBatchRemoveChannelRoleFromAccessors {
			repostspace_BatchRemoveChannelRoleFromAccessors(cfg, client)
			return
		}
		if _repostspaceBatchRemoveRole {
			repostspace_BatchRemoveRole(cfg, client)
			return
		}
		if _repostspaceCreateChannel {
			repostspace_CreateChannel(cfg, client)
			return
		}
		if _repostspaceCreateSpace {
			repostspace_CreateSpace(cfg, client)
			return
		}
		if _repostspaceDeleteSpace {
			repostspace_DeleteSpace(cfg, client)
			return
		}
		if _repostspaceDeregisterAdmin {
			repostspace_DeregisterAdmin(cfg, client)
			return
		}
		if _repostspaceGetChannel {
			repostspace_GetChannel(cfg, client)
			return
		}
		if _repostspaceGetSpace {
			repostspace_GetSpace(cfg, client)
			return
		}
		if _repostspaceListChannels {
			repostspace_ListChannels(cfg, client)
			return
		}
		if _repostspaceListSpaces {
			repostspace_ListSpaces(cfg, client)
			return
		}
		if _repostspaceListTagsForResource {
			repostspace_ListTagsForResource(cfg, client)
			return
		}
		if _repostspaceRegisterAdmin {
			repostspace_RegisterAdmin(cfg, client)
			return
		}
		if _repostspaceSendInvites {
			repostspace_SendInvites(cfg, client)
			return
		}
		if _repostspaceTagResource {
			repostspace_TagResource(cfg, client)
			return
		}
		if _repostspaceUntagResource {
			repostspace_UntagResource(cfg, client)
			return
		}
		if _repostspaceUpdateChannel {
			repostspace_UpdateChannel(cfg, client)
			return
		}
		if _repostspaceUpdateSpace {
			repostspace_UpdateSpace(cfg, client)
			return
		}

	},
}

var (
	_repostspaceBatchAddChannelRoleToAccessors      bool
	_repostspaceBatchAddRole                        bool
	_repostspaceBatchRemoveChannelRoleFromAccessors bool
	_repostspaceBatchRemoveRole                     bool
	_repostspaceCreateChannel                       bool
	_repostspaceCreateSpace                         bool
	_repostspaceDeleteSpace                         bool
	_repostspaceDeregisterAdmin                     bool
	_repostspaceGetChannel                          bool
	_repostspaceGetSpace                            bool
	_repostspaceListChannels                        bool
	_repostspaceListSpaces                          bool
	_repostspaceListTagsForResource                 bool
	_repostspaceRegisterAdmin                       bool
	_repostspaceSendInvites                         bool
	_repostspaceTagResource                         bool
	_repostspaceUntagResource                       bool
	_repostspaceUpdateChannel                       bool
	_repostspaceUpdateSpace                         bool

	_repostspaceAccessorIds           []string
	_repostspaceAdminId               string
	_repostspaceBody                  string
	_repostspaceChannelDescription    string
	_repostspaceChannelId             string
	_repostspaceChannelName           string
	_repostspaceChannelRole           string
	_repostspaceDescription           string
	_repostspaceMaxResults            string
	_repostspaceName                  string
	_repostspaceNextToken             string
	_repostspaceResourceArn           string
	_repostspaceRole                  string
	_repostspaceRoleArn               string
	_repostspaceSpaceId               string
	_repostspaceSubdomain             string
	_repostspaceSupportedEmailDomains string
	_repostspaceTagKeys               []string
	_repostspaceTags                  string
	_repostspaceTier                  string
	_repostspaceTitle                 string
	_repostspaceUserKMSKey            string
)

// Add role to multiple users or groups in a private re:Post channel.
func repostspace_BatchAddChannelRoleToAccessors(cfg aws.Config, client *repostspace.Client) {
	input := &repostspace.BatchAddChannelRoleToAccessorsInput{
		// AccessorIds: []string, // Required
		// ChannelId: *string, // Required
		// ChannelRole: types.ChannelRole, // Required
		// SpaceId: *string, // Required
	}

	if len(_repostspaceAccessorIds) > 0 {
		input.AccessorIds = append([]string(nil), _repostspaceAccessorIds...)
	}
	if len(_repostspaceChannelId) > 0 {
		input.ChannelId = aws.String(_repostspaceChannelId)
	}
	if len(_repostspaceChannelRole) > 0 {
		if err := assignInputField(input, "ChannelRole", _repostspaceChannelRole); err != nil {
			log.Errorf("invalid --channel-role: %s", err.Error())
			return
		}
	}
	if len(_repostspaceSpaceId) > 0 {
		input.SpaceId = aws.String(_repostspaceSpaceId)
	}

	if resp, err := client.BatchAddChannelRoleToAccessors(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Add a role to multiple users or groups in a private re:Post.
func repostspace_BatchAddRole(cfg aws.Config, client *repostspace.Client) {
	input := &repostspace.BatchAddRoleInput{
		// AccessorIds: []string, // Required
		// Role: types.Role, // Required
		// SpaceId: *string, // Required
	}

	if len(_repostspaceAccessorIds) > 0 {
		input.AccessorIds = append([]string(nil), _repostspaceAccessorIds...)
	}
	if len(_repostspaceRole) > 0 {
		if err := assignInputField(input, "Role", _repostspaceRole); err != nil {
			log.Errorf("invalid --role: %s", err.Error())
			return
		}
	}
	if len(_repostspaceSpaceId) > 0 {
		input.SpaceId = aws.String(_repostspaceSpaceId)
	}

	if resp, err := client.BatchAddRole(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Remove a role from multiple users or groups in a private re:Post channel.
func repostspace_BatchRemoveChannelRoleFromAccessors(cfg aws.Config, client *repostspace.Client) {
	input := &repostspace.BatchRemoveChannelRoleFromAccessorsInput{
		// AccessorIds: []string, // Required
		// ChannelId: *string, // Required
		// ChannelRole: types.ChannelRole, // Required
		// SpaceId: *string, // Required
	}

	if len(_repostspaceAccessorIds) > 0 {
		input.AccessorIds = append([]string(nil), _repostspaceAccessorIds...)
	}
	if len(_repostspaceChannelId) > 0 {
		input.ChannelId = aws.String(_repostspaceChannelId)
	}
	if len(_repostspaceChannelRole) > 0 {
		if err := assignInputField(input, "ChannelRole", _repostspaceChannelRole); err != nil {
			log.Errorf("invalid --channel-role: %s", err.Error())
			return
		}
	}
	if len(_repostspaceSpaceId) > 0 {
		input.SpaceId = aws.String(_repostspaceSpaceId)
	}

	if resp, err := client.BatchRemoveChannelRoleFromAccessors(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Remove a role from multiple users or groups in a private re:Post.
func repostspace_BatchRemoveRole(cfg aws.Config, client *repostspace.Client) {
	input := &repostspace.BatchRemoveRoleInput{
		// AccessorIds: []string, // Required
		// Role: types.Role, // Required
		// SpaceId: *string, // Required
	}

	if len(_repostspaceAccessorIds) > 0 {
		input.AccessorIds = append([]string(nil), _repostspaceAccessorIds...)
	}
	if len(_repostspaceRole) > 0 {
		if err := assignInputField(input, "Role", _repostspaceRole); err != nil {
			log.Errorf("invalid --role: %s", err.Error())
			return
		}
	}
	if len(_repostspaceSpaceId) > 0 {
		input.SpaceId = aws.String(_repostspaceSpaceId)
	}

	if resp, err := client.BatchRemoveRole(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a channel in an AWS re:Post Private private re:Post.
func repostspace_CreateChannel(cfg aws.Config, client *repostspace.Client) {
	input := &repostspace.CreateChannelInput{
		// ChannelName: *string, // Required
		// SpaceId: *string, // Required
	}

	if len(_repostspaceChannelName) > 0 {
		input.ChannelName = aws.String(_repostspaceChannelName)
	}
	if len(_repostspaceSpaceId) > 0 {
		input.SpaceId = aws.String(_repostspaceSpaceId)
	}
	if len(_repostspaceChannelDescription) > 0 {
		input.ChannelDescription = aws.String(_repostspaceChannelDescription)
	}

	if resp, err := client.CreateChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an AWS re:Post Private private re:Post.
func repostspace_CreateSpace(cfg aws.Config, client *repostspace.Client) {
	input := &repostspace.CreateSpaceInput{
		// Name: *string, // Required
		// Subdomain: *string, // Required
		// Tier: types.TierLevel, // Required
	}

	if len(_repostspaceName) > 0 {
		input.Name = aws.String(_repostspaceName)
	}
	if len(_repostspaceSubdomain) > 0 {
		input.Subdomain = aws.String(_repostspaceSubdomain)
	}
	if len(_repostspaceTier) > 0 {
		if err := assignInputField(input, "Tier", _repostspaceTier); err != nil {
			log.Errorf("invalid --tier: %s", err.Error())
			return
		}
	}
	if len(_repostspaceDescription) > 0 {
		input.Description = aws.String(_repostspaceDescription)
	}
	if len(_repostspaceRoleArn) > 0 {
		input.RoleArn = aws.String(_repostspaceRoleArn)
	}
	if len(_repostspaceSupportedEmailDomains) > 0 {
		if err := assignInputField(input, "SupportedEmailDomains", _repostspaceSupportedEmailDomains); err != nil {
			log.Errorf("invalid --supported-email-domains: %s", err.Error())
			return
		}
	}
	if len(_repostspaceTags) > 0 {
		if err := assignInputField(input, "Tags", _repostspaceTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_repostspaceUserKMSKey) > 0 {
		input.UserKMSKey = aws.String(_repostspaceUserKMSKey)
	}

	if resp, err := client.CreateSpace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an AWS re:Post Private private re:Post.
func repostspace_DeleteSpace(cfg aws.Config, client *repostspace.Client) {
	input := &repostspace.DeleteSpaceInput{
		// SpaceId: *string, // Required
	}

	if len(_repostspaceSpaceId) > 0 {
		input.SpaceId = aws.String(_repostspaceSpaceId)
	}

	if resp, err := client.DeleteSpace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the user or group from the list of administrators of the private
// re:Post.
func repostspace_DeregisterAdmin(cfg aws.Config, client *repostspace.Client) {
	input := &repostspace.DeregisterAdminInput{
		// AdminId: *string, // Required
		// SpaceId: *string, // Required
	}

	if len(_repostspaceAdminId) > 0 {
		input.AdminId = aws.String(_repostspaceAdminId)
	}
	if len(_repostspaceSpaceId) > 0 {
		input.SpaceId = aws.String(_repostspaceSpaceId)
	}

	if resp, err := client.DeregisterAdmin(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Displays information about a channel in a private re:Post.
func repostspace_GetChannel(cfg aws.Config, client *repostspace.Client) {
	input := &repostspace.GetChannelInput{
		// ChannelId: *string, // Required
		// SpaceId: *string, // Required
	}

	if len(_repostspaceChannelId) > 0 {
		input.ChannelId = aws.String(_repostspaceChannelId)
	}
	if len(_repostspaceSpaceId) > 0 {
		input.SpaceId = aws.String(_repostspaceSpaceId)
	}

	if resp, err := client.GetChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Displays information about the AWS re:Post Private private re:Post.
func repostspace_GetSpace(cfg aws.Config, client *repostspace.Client) {
	input := &repostspace.GetSpaceInput{
		// SpaceId: *string, // Required
	}

	if len(_repostspaceSpaceId) > 0 {
		input.SpaceId = aws.String(_repostspaceSpaceId)
	}

	if resp, err := client.GetSpace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the list of channel within a private re:Post with some information
// about each channel.
func repostspace_ListChannels(cfg aws.Config, client *repostspace.Client) {
	input := &repostspace.ListChannelsInput{
		// SpaceId: *string, // Required
	}

	if len(_repostspaceSpaceId) > 0 {
		input.SpaceId = aws.String(_repostspaceSpaceId)
	}
	if len(_repostspaceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _repostspaceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_repostspaceNextToken) > 0 {
		input.NextToken = aws.String(_repostspaceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListChannels(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*repostspace.ListChannelsOutput
	p := repostspace.NewListChannelsPaginator(client, input)
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

// Returns a list of AWS re:Post Private private re:Posts in the account with some
// information about each private re:Post.
func repostspace_ListSpaces(cfg aws.Config, client *repostspace.Client) {
	input := &repostspace.ListSpacesInput{}

	if len(_repostspaceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _repostspaceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_repostspaceNextToken) > 0 {
		input.NextToken = aws.String(_repostspaceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSpaces(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*repostspace.ListSpacesOutput
	p := repostspace.NewListSpacesPaginator(client, input)
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

// Returns the tags that are associated with the AWS re:Post Private resource
// specified by the resourceArn. The only resource that can be tagged is a private
// re:Post.
func repostspace_ListTagsForResource(cfg aws.Config, client *repostspace.Client) {
	input := &repostspace.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_repostspaceResourceArn) > 0 {
		input.ResourceArn = aws.String(_repostspaceResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a user or group to the list of administrators of the private re:Post.
func repostspace_RegisterAdmin(cfg aws.Config, client *repostspace.Client) {
	input := &repostspace.RegisterAdminInput{
		// AdminId: *string, // Required
		// SpaceId: *string, // Required
	}

	if len(_repostspaceAdminId) > 0 {
		input.AdminId = aws.String(_repostspaceAdminId)
	}
	if len(_repostspaceSpaceId) > 0 {
		input.SpaceId = aws.String(_repostspaceSpaceId)
	}

	if resp, err := client.RegisterAdmin(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sends an invitation email to selected users and groups.
func repostspace_SendInvites(cfg aws.Config, client *repostspace.Client) {
	input := &repostspace.SendInvitesInput{
		// AccessorIds: []string, // Required
		// Body: *string, // Required
		// SpaceId: *string, // Required
		// Title: *string, // Required
	}

	if len(_repostspaceAccessorIds) > 0 {
		input.AccessorIds = append([]string(nil), _repostspaceAccessorIds...)
	}
	if len(_repostspaceBody) > 0 {
		input.Body = aws.String(_repostspaceBody)
	}
	if len(_repostspaceSpaceId) > 0 {
		input.SpaceId = aws.String(_repostspaceSpaceId)
	}
	if len(_repostspaceTitle) > 0 {
		input.Title = aws.String(_repostspaceTitle)
	}

	if resp, err := client.SendInvites(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates tags with an AWS re:Post Private resource. Currently, the only
// resource that can be tagged is the private re:Post. If you specify a new tag key
// for the resource, the tag is appended to the list of tags that are associated
// with the resource. If you specify a tag key that’s already associated with the
// resource, the new tag value that you specify replaces the previous value for
// that tag.
func repostspace_TagResource(cfg aws.Config, client *repostspace.Client) {
	input := &repostspace.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_repostspaceResourceArn) > 0 {
		input.ResourceArn = aws.String(_repostspaceResourceArn)
	}
	if len(_repostspaceTags) > 0 {
		if err := assignInputField(input, "Tags", _repostspaceTags); err != nil {
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

// Removes the association of the tag with the AWS re:Post Private resource.
func repostspace_UntagResource(cfg aws.Config, client *repostspace.Client) {
	input := &repostspace.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_repostspaceResourceArn) > 0 {
		input.ResourceArn = aws.String(_repostspaceResourceArn)
	}
	if len(_repostspaceTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _repostspaceTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies an existing channel.
func repostspace_UpdateChannel(cfg aws.Config, client *repostspace.Client) {
	input := &repostspace.UpdateChannelInput{
		// ChannelId: *string, // Required
		// ChannelName: *string, // Required
		// SpaceId: *string, // Required
	}

	if len(_repostspaceChannelId) > 0 {
		input.ChannelId = aws.String(_repostspaceChannelId)
	}
	if len(_repostspaceChannelName) > 0 {
		input.ChannelName = aws.String(_repostspaceChannelName)
	}
	if len(_repostspaceSpaceId) > 0 {
		input.SpaceId = aws.String(_repostspaceSpaceId)
	}
	if len(_repostspaceChannelDescription) > 0 {
		input.ChannelDescription = aws.String(_repostspaceChannelDescription)
	}

	if resp, err := client.UpdateChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies an existing AWS re:Post Private private re:Post.
func repostspace_UpdateSpace(cfg aws.Config, client *repostspace.Client) {
	input := &repostspace.UpdateSpaceInput{
		// SpaceId: *string, // Required
	}

	if len(_repostspaceSpaceId) > 0 {
		input.SpaceId = aws.String(_repostspaceSpaceId)
	}
	if len(_repostspaceDescription) > 0 {
		input.Description = aws.String(_repostspaceDescription)
	}
	if len(_repostspaceRoleArn) > 0 {
		input.RoleArn = aws.String(_repostspaceRoleArn)
	}
	if len(_repostspaceSupportedEmailDomains) > 0 {
		if err := assignInputField(input, "SupportedEmailDomains", _repostspaceSupportedEmailDomains); err != nil {
			log.Errorf("invalid --supported-email-domains: %s", err.Error())
			return
		}
	}
	if len(_repostspaceTier) > 0 {
		if err := assignInputField(input, "Tier", _repostspaceTier); err != nil {
			log.Errorf("invalid --tier: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSpace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_repostspaceCmd)
	_repostspaceCmd.Flags().SortFlags = false

	_repostspaceCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_repostspaceCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_repostspaceCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_repostspaceCmd.Flags().StringSliceVarP(&_repostspaceAccessorIds, "accessor-ids", "", nil, "Accessor Ids")
	_repostspaceCmd.Flags().StringVarP(&_repostspaceAdminId, "admin-id", "", "", "Admin ID")
	_repostspaceCmd.Flags().StringVarP(&_repostspaceBody, "body", "", "", "Body")
	_repostspaceCmd.Flags().StringVarP(&_repostspaceChannelDescription, "channel-description", "", "", "Channel Description")
	_repostspaceCmd.Flags().StringVarP(&_repostspaceChannelId, "channel-id", "", "", "Channel ID")
	_repostspaceCmd.Flags().StringVarP(&_repostspaceChannelName, "channel-name", "", "", "Channel Name")
	_repostspaceCmd.Flags().StringVarP(&_repostspaceChannelRole, "channel-role", "", "", "Channel Role")
	_repostspaceCmd.Flags().StringVarP(&_repostspaceDescription, "description", "", "", "Description")
	_repostspaceCmd.Flags().StringVarP(&_repostspaceMaxResults, "max-results", "", "", "Max Results")
	_repostspaceCmd.Flags().StringVarP(&_repostspaceName, "name", "", "", "Name")
	_repostspaceCmd.Flags().StringVarP(&_repostspaceNextToken, "next-token", "", "", "Next Token")
	_repostspaceCmd.Flags().StringVarP(&_repostspaceResourceArn, "resource-arn", "", "", "Resource ARN")
	_repostspaceCmd.Flags().StringVarP(&_repostspaceRole, "role", "", "", "Role")
	_repostspaceCmd.Flags().StringVarP(&_repostspaceRoleArn, "role-arn", "", "", "Role ARN")
	_repostspaceCmd.Flags().StringVarP(&_repostspaceSpaceId, "space-id", "", "", "Space ID")
	_repostspaceCmd.Flags().StringVarP(&_repostspaceSubdomain, "subdomain", "", "", "Subdomain")
	_repostspaceCmd.Flags().StringVarP(&_repostspaceSupportedEmailDomains, "supported-email-domains", "", "", "Supported Email Domains")
	_repostspaceCmd.Flags().StringSliceVarP(&_repostspaceTagKeys, "tag-keys", "", nil, "Tag Keys")
	_repostspaceCmd.Flags().StringVarP(&_repostspaceTags, "tags", "", "", "Tags")
	_repostspaceCmd.Flags().StringVarP(&_repostspaceTier, "tier", "", "", "Tier")
	_repostspaceCmd.Flags().StringVarP(&_repostspaceTitle, "title", "", "", "Title")
	_repostspaceCmd.Flags().StringVarP(&_repostspaceUserKMSKey, "user-kms-key", "", "", "User KMS Key")

	_repostspaceCmd.Flags().BoolVarP(&_repostspaceBatchAddChannelRoleToAccessors, "batch-add-channel-role-to-accessors", "", false, "Batch Add Channel Role To Accessors")
	_repostspaceCmd.Flags().BoolVarP(&_repostspaceBatchAddRole, "batch-add-role", "", false, "Batch Add Role")
	_repostspaceCmd.Flags().BoolVarP(&_repostspaceBatchRemoveChannelRoleFromAccessors, "batch-remove-channel-role-from-accessors", "", false, "Batch Remove Channel Role From Accessors")
	_repostspaceCmd.Flags().BoolVarP(&_repostspaceBatchRemoveRole, "batch-remove-role", "", false, "Batch Remove Role")
	_repostspaceCmd.Flags().BoolVarP(&_repostspaceCreateChannel, "create-channel", "", false, "Create Channel")
	_repostspaceCmd.Flags().BoolVarP(&_repostspaceCreateSpace, "create-space", "", false, "Create Space")
	_repostspaceCmd.Flags().BoolVarP(&_repostspaceDeleteSpace, "delete-space", "", false, "Delete Space")
	_repostspaceCmd.Flags().BoolVarP(&_repostspaceDeregisterAdmin, "deregister-admin", "", false, "Deregister Admin")
	_repostspaceCmd.Flags().BoolVarP(&_repostspaceGetChannel, "get-channel", "", false, "Get Channel")
	_repostspaceCmd.Flags().BoolVarP(&_repostspaceGetSpace, "get-space", "", false, "Get Space")
	_repostspaceCmd.Flags().BoolVarP(&_repostspaceListChannels, "list-channels", "", false, "List Channels")
	_repostspaceCmd.Flags().BoolVarP(&_repostspaceListSpaces, "list-spaces", "", false, "List Spaces")
	_repostspaceCmd.Flags().BoolVarP(&_repostspaceListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_repostspaceCmd.Flags().BoolVarP(&_repostspaceRegisterAdmin, "register-admin", "", false, "Register Admin")
	_repostspaceCmd.Flags().BoolVarP(&_repostspaceSendInvites, "send-invites", "", false, "Send Invites")
	_repostspaceCmd.Flags().BoolVarP(&_repostspaceTagResource, "tag-resource", "", false, "Tag Resource")
	_repostspaceCmd.Flags().BoolVarP(&_repostspaceUntagResource, "untag-resource", "", false, "Untag Resource")
	_repostspaceCmd.Flags().BoolVarP(&_repostspaceUpdateChannel, "update-channel", "", false, "Update Channel")
	_repostspaceCmd.Flags().BoolVarP(&_repostspaceUpdateSpace, "update-space", "", false, "Update Space")

}
