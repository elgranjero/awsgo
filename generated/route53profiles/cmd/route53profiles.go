package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53profiles"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// route53profilesCmd represents the route53profiles command
var _route53profilesCmd = &cobra.Command{
	Use:   "route53profiles",
	Short: "AWS route53profiles CLI",
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
		client := route53profiles.NewFromConfig(cfg)
		if _route53profilesAssociateProfile {
			route53profiles_AssociateProfile(cfg, client)
			return
		}
		if _route53profilesAssociateResourceToProfile {
			route53profiles_AssociateResourceToProfile(cfg, client)
			return
		}
		if _route53profilesCreateProfile {
			route53profiles_CreateProfile(cfg, client)
			return
		}
		if _route53profilesDeleteProfile {
			route53profiles_DeleteProfile(cfg, client)
			return
		}
		if _route53profilesDisassociateProfile {
			route53profiles_DisassociateProfile(cfg, client)
			return
		}
		if _route53profilesDisassociateResourceFromProfile {
			route53profiles_DisassociateResourceFromProfile(cfg, client)
			return
		}
		if _route53profilesGetProfile {
			route53profiles_GetProfile(cfg, client)
			return
		}
		if _route53profilesGetProfileAssociation {
			route53profiles_GetProfileAssociation(cfg, client)
			return
		}
		if _route53profilesGetProfileResourceAssociation {
			route53profiles_GetProfileResourceAssociation(cfg, client)
			return
		}
		if _route53profilesListProfileAssociations {
			route53profiles_ListProfileAssociations(cfg, client)
			return
		}
		if _route53profilesListProfileResourceAssociations {
			route53profiles_ListProfileResourceAssociations(cfg, client)
			return
		}
		if _route53profilesListProfiles {
			route53profiles_ListProfiles(cfg, client)
			return
		}
		if _route53profilesListTagsForResource {
			route53profiles_ListTagsForResource(cfg, client)
			return
		}
		if _route53profilesTagResource {
			route53profiles_TagResource(cfg, client)
			return
		}
		if _route53profilesUntagResource {
			route53profiles_UntagResource(cfg, client)
			return
		}
		if _route53profilesUpdateProfileResourceAssociation {
			route53profiles_UpdateProfileResourceAssociation(cfg, client)
			return
		}

	},
}

var (
	_route53profilesAssociateProfile                 bool
	_route53profilesAssociateResourceToProfile       bool
	_route53profilesCreateProfile                    bool
	_route53profilesDeleteProfile                    bool
	_route53profilesDisassociateProfile              bool
	_route53profilesDisassociateResourceFromProfile  bool
	_route53profilesGetProfile                       bool
	_route53profilesGetProfileAssociation            bool
	_route53profilesGetProfileResourceAssociation    bool
	_route53profilesListProfileAssociations          bool
	_route53profilesListProfileResourceAssociations  bool
	_route53profilesListProfiles                     bool
	_route53profilesListTagsForResource              bool
	_route53profilesTagResource                      bool
	_route53profilesUntagResource                    bool
	_route53profilesUpdateProfileResourceAssociation bool

	_route53profilesClientToken                  string
	_route53profilesMaxResults                   string
	_route53profilesName                         string
	_route53profilesNextToken                    string
	_route53profilesProfileAssociationId         string
	_route53profilesProfileId                    string
	_route53profilesProfileResourceAssociationId string
	_route53profilesResourceArn                  string
	_route53profilesResourceId                   string
	_route53profilesResourceProperties           string
	_route53profilesResourceType                 string
	_route53profilesTagKeys                      []string
	_route53profilesTags                         string
)

// Associates a Route 53 Profiles profile with a VPC. A VPC can have only one
// Profile associated with it, but a Profile can be associated with 1000 of VPCs
// (and you can request a higher quota). For more information, see [https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DNSLimitations.html#limits-api-entities].
//
// [https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DNSLimitations.html#limits-api-entities]: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DNSLimitations.html#limits-api-entities
func route53profiles_AssociateProfile(cfg aws.Config, client *route53profiles.Client) {
	input := &route53profiles.AssociateProfileInput{
		// Name: *string, // Required
		// ProfileId: *string, // Required
		// ResourceId: *string, // Required
	}

	if len(_route53profilesName) > 0 {
		input.Name = aws.String(_route53profilesName)
	}
	if len(_route53profilesProfileId) > 0 {
		input.ProfileId = aws.String(_route53profilesProfileId)
	}
	if len(_route53profilesResourceId) > 0 {
		input.ResourceId = aws.String(_route53profilesResourceId)
	}
	if len(_route53profilesTags) > 0 {
		if err := assignInputField(input, "Tags", _route53profilesTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.AssociateProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a DNS reource configuration to a Route 53 Profile.
func route53profiles_AssociateResourceToProfile(cfg aws.Config, client *route53profiles.Client) {
	input := &route53profiles.AssociateResourceToProfileInput{
		// Name: *string, // Required
		// ProfileId: *string, // Required
		// ResourceArn: *string, // Required
	}

	if len(_route53profilesName) > 0 {
		input.Name = aws.String(_route53profilesName)
	}
	if len(_route53profilesProfileId) > 0 {
		input.ProfileId = aws.String(_route53profilesProfileId)
	}
	if len(_route53profilesResourceArn) > 0 {
		input.ResourceArn = aws.String(_route53profilesResourceArn)
	}
	if len(_route53profilesResourceProperties) > 0 {
		input.ResourceProperties = aws.String(_route53profilesResourceProperties)
	}

	if resp, err := client.AssociateResourceToProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an empty Route 53 Profile.
func route53profiles_CreateProfile(cfg aws.Config, client *route53profiles.Client) {
	input := &route53profiles.CreateProfileInput{
		// ClientToken: *string, // Required
		// Name: *string, // Required
	}

	if len(_route53profilesClientToken) > 0 {
		input.ClientToken = aws.String(_route53profilesClientToken)
	}
	if len(_route53profilesName) > 0 {
		input.Name = aws.String(_route53profilesName)
	}
	if len(_route53profilesTags) > 0 {
		if err := assignInputField(input, "Tags", _route53profilesTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified Route 53 Profile. Before you can delete a profile, you
// must first disassociate it from all VPCs.
func route53profiles_DeleteProfile(cfg aws.Config, client *route53profiles.Client) {
	input := &route53profiles.DeleteProfileInput{
		// ProfileId: *string, // Required
	}

	if len(_route53profilesProfileId) > 0 {
		input.ProfileId = aws.String(_route53profilesProfileId)
	}

	if resp, err := client.DeleteProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Dissociates a specified Route 53 Profile from the specified VPC.
func route53profiles_DisassociateProfile(cfg aws.Config, client *route53profiles.Client) {
	input := &route53profiles.DisassociateProfileInput{
		// ProfileId: *string, // Required
		// ResourceId: *string, // Required
	}

	if len(_route53profilesProfileId) > 0 {
		input.ProfileId = aws.String(_route53profilesProfileId)
	}
	if len(_route53profilesResourceId) > 0 {
		input.ResourceId = aws.String(_route53profilesResourceId)
	}

	if resp, err := client.DisassociateProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Dissoaciated a specified resource, from the Route 53 Profile.
func route53profiles_DisassociateResourceFromProfile(cfg aws.Config, client *route53profiles.Client) {
	input := &route53profiles.DisassociateResourceFromProfileInput{
		// ProfileId: *string, // Required
		// ResourceArn: *string, // Required
	}

	if len(_route53profilesProfileId) > 0 {
		input.ProfileId = aws.String(_route53profilesProfileId)
	}
	if len(_route53profilesResourceArn) > 0 {
		input.ResourceArn = aws.String(_route53profilesResourceArn)
	}

	if resp, err := client.DisassociateResourceFromProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a specified Route 53 Profile, such as whether
// whether the Profile is shared, and the current status of the Profile.
func route53profiles_GetProfile(cfg aws.Config, client *route53profiles.Client) {
	input := &route53profiles.GetProfileInput{
		// ProfileId: *string, // Required
	}

	if len(_route53profilesProfileId) > 0 {
		input.ProfileId = aws.String(_route53profilesProfileId)
	}

	if resp, err := client.GetProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a Route 53 Profile association for a VPC. A VPC can have only one
// Profile association, but a Profile can be associated with up to 5000 VPCs.
func route53profiles_GetProfileAssociation(cfg aws.Config, client *route53profiles.Client) {
	input := &route53profiles.GetProfileAssociationInput{
		// ProfileAssociationId: *string, // Required
	}

	if len(_route53profilesProfileAssociationId) > 0 {
		input.ProfileAssociationId = aws.String(_route53profilesProfileAssociationId)
	}

	if resp, err := client.GetProfileAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a specified Route 53 Profile resource association.
func route53profiles_GetProfileResourceAssociation(cfg aws.Config, client *route53profiles.Client) {
	input := &route53profiles.GetProfileResourceAssociationInput{
		// ProfileResourceAssociationId: *string, // Required
	}

	if len(_route53profilesProfileResourceAssociationId) > 0 {
		input.ProfileResourceAssociationId = aws.String(_route53profilesProfileResourceAssociationId)
	}

	if resp, err := client.GetProfileResourceAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all the VPCs that the specified Route 53 Profile is associated with.
func route53profiles_ListProfileAssociations(cfg aws.Config, client *route53profiles.Client) {
	input := &route53profiles.ListProfileAssociationsInput{}

	if len(_route53profilesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53profilesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53profilesNextToken) > 0 {
		input.NextToken = aws.String(_route53profilesNextToken)
	}
	if len(_route53profilesProfileId) > 0 {
		input.ProfileId = aws.String(_route53profilesProfileId)
	}
	if len(_route53profilesResourceId) > 0 {
		input.ResourceId = aws.String(_route53profilesResourceId)
	}

	if disablePaginator() {
		if resp, err := client.ListProfileAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53profiles.ListProfileAssociationsOutput
	p := route53profiles.NewListProfileAssociationsPaginator(client, input)
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

// Lists all the resource associations for the specified Route 53 Profile.
func route53profiles_ListProfileResourceAssociations(cfg aws.Config, client *route53profiles.Client) {
	input := &route53profiles.ListProfileResourceAssociationsInput{
		// ProfileId: *string, // Required
	}

	if len(_route53profilesProfileId) > 0 {
		input.ProfileId = aws.String(_route53profilesProfileId)
	}
	if len(_route53profilesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53profilesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53profilesNextToken) > 0 {
		input.NextToken = aws.String(_route53profilesNextToken)
	}
	if len(_route53profilesResourceType) > 0 {
		input.ResourceType = aws.String(_route53profilesResourceType)
	}

	if disablePaginator() {
		if resp, err := client.ListProfileResourceAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53profiles.ListProfileResourceAssociationsOutput
	p := route53profiles.NewListProfileResourceAssociationsPaginator(client, input)
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

// Lists all the Route 53 Profiles associated with your Amazon Web Services
// account.
func route53profiles_ListProfiles(cfg aws.Config, client *route53profiles.Client) {
	input := &route53profiles.ListProfilesInput{}

	if len(_route53profilesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53profilesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53profilesNextToken) > 0 {
		input.NextToken = aws.String(_route53profilesNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListProfiles(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53profiles.ListProfilesOutput
	p := route53profiles.NewListProfilesPaginator(client, input)
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

// Lists the tags that you associated with the specified resource.
func route53profiles_ListTagsForResource(cfg aws.Config, client *route53profiles.Client) {
	input := &route53profiles.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_route53profilesResourceArn) > 0 {
		input.ResourceArn = aws.String(_route53profilesResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds one or more tags to a specified resource.
func route53profiles_TagResource(cfg aws.Config, client *route53profiles.Client) {
	input := &route53profiles.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_route53profilesResourceArn) > 0 {
		input.ResourceArn = aws.String(_route53profilesResourceArn)
	}
	if len(_route53profilesTags) > 0 {
		if err := assignInputField(input, "Tags", _route53profilesTags); err != nil {
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

// Removes one or more tags from a specified resource.
func route53profiles_UntagResource(cfg aws.Config, client *route53profiles.Client) {
	input := &route53profiles.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_route53profilesResourceArn) > 0 {
		input.ResourceArn = aws.String(_route53profilesResourceArn)
	}
	if len(_route53profilesTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _route53profilesTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified Route 53 Profile resourse association.
func route53profiles_UpdateProfileResourceAssociation(cfg aws.Config, client *route53profiles.Client) {
	input := &route53profiles.UpdateProfileResourceAssociationInput{
		// ProfileResourceAssociationId: *string, // Required
	}

	if len(_route53profilesProfileResourceAssociationId) > 0 {
		input.ProfileResourceAssociationId = aws.String(_route53profilesProfileResourceAssociationId)
	}
	if len(_route53profilesName) > 0 {
		input.Name = aws.String(_route53profilesName)
	}
	if len(_route53profilesResourceProperties) > 0 {
		input.ResourceProperties = aws.String(_route53profilesResourceProperties)
	}

	if resp, err := client.UpdateProfileResourceAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_route53profilesCmd)
	_route53profilesCmd.Flags().SortFlags = false

	_route53profilesCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_route53profilesCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_route53profilesCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_route53profilesCmd.Flags().StringVarP(&_route53profilesClientToken, "client-token", "", "", "Client Token")
	_route53profilesCmd.Flags().StringVarP(&_route53profilesMaxResults, "max-results", "", "", "Max Results")
	_route53profilesCmd.Flags().StringVarP(&_route53profilesName, "name", "", "", "Name")
	_route53profilesCmd.Flags().StringVarP(&_route53profilesNextToken, "next-token", "", "", "Next Token")
	_route53profilesCmd.Flags().StringVarP(&_route53profilesProfileAssociationId, "profile-association-id", "", "", "Profile Association ID")
	_route53profilesCmd.Flags().StringVarP(&_route53profilesProfileId, "profile-id", "", "", "Profile ID")
	_route53profilesCmd.Flags().StringVarP(&_route53profilesProfileResourceAssociationId, "profile-resource-association-id", "", "", "Profile Resource Association ID")
	_route53profilesCmd.Flags().StringVarP(&_route53profilesResourceArn, "resource-arn", "", "", "Resource ARN")
	_route53profilesCmd.Flags().StringVarP(&_route53profilesResourceId, "resource-id", "", "", "Resource ID")
	_route53profilesCmd.Flags().StringVarP(&_route53profilesResourceProperties, "resource-properties", "", "", "Resource Properties")
	_route53profilesCmd.Flags().StringVarP(&_route53profilesResourceType, "resource-type", "", "", "Resource Type")
	_route53profilesCmd.Flags().StringSliceVarP(&_route53profilesTagKeys, "tag-keys", "", nil, "Tag Keys")
	_route53profilesCmd.Flags().StringVarP(&_route53profilesTags, "tags", "", "", "Tags")

	_route53profilesCmd.Flags().BoolVarP(&_route53profilesAssociateProfile, "associate-profile", "", false, "Associate Profile")
	_route53profilesCmd.Flags().BoolVarP(&_route53profilesAssociateResourceToProfile, "associate-resource-to-profile", "", false, "Associate Resource To Profile")
	_route53profilesCmd.Flags().BoolVarP(&_route53profilesCreateProfile, "create-profile", "", false, "Create Profile")
	_route53profilesCmd.Flags().BoolVarP(&_route53profilesDeleteProfile, "delete-profile", "", false, "Delete Profile")
	_route53profilesCmd.Flags().BoolVarP(&_route53profilesDisassociateProfile, "disassociate-profile", "", false, "Disassociate Profile")
	_route53profilesCmd.Flags().BoolVarP(&_route53profilesDisassociateResourceFromProfile, "disassociate-resource-from-profile", "", false, "Disassociate Resource From Profile")
	_route53profilesCmd.Flags().BoolVarP(&_route53profilesGetProfile, "get-profile", "", false, "Get Profile")
	_route53profilesCmd.Flags().BoolVarP(&_route53profilesGetProfileAssociation, "get-profile-association", "", false, "Get Profile Association")
	_route53profilesCmd.Flags().BoolVarP(&_route53profilesGetProfileResourceAssociation, "get-profile-resource-association", "", false, "Get Profile Resource Association")
	_route53profilesCmd.Flags().BoolVarP(&_route53profilesListProfileAssociations, "list-profile-associations", "", false, "List Profile Associations")
	_route53profilesCmd.Flags().BoolVarP(&_route53profilesListProfileResourceAssociations, "list-profile-resource-associations", "", false, "List Profile Resource Associations")
	_route53profilesCmd.Flags().BoolVarP(&_route53profilesListProfiles, "list-profiles", "", false, "List Profiles")
	_route53profilesCmd.Flags().BoolVarP(&_route53profilesListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_route53profilesCmd.Flags().BoolVarP(&_route53profilesTagResource, "tag-resource", "", false, "Tag Resource")
	_route53profilesCmd.Flags().BoolVarP(&_route53profilesUntagResource, "untag-resource", "", false, "Untag Resource")
	_route53profilesCmd.Flags().BoolVarP(&_route53profilesUpdateProfileResourceAssociation, "update-profile-resource-association", "", false, "Update Profile Resource Association")

}
