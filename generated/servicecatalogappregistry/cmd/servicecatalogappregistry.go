package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// servicecatalogappregistryCmd represents the servicecatalogappregistry command
var _servicecatalogappregistryCmd = &cobra.Command{
	Use:   "servicecatalogappregistry",
	Short: "AWS servicecatalogappregistry CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := servicecatalogappregistry.NewFromConfig(cfg)
		if _servicecatalogappregistryAssociateAttributeGroup {
			servicecatalogappregistry_AssociateAttributeGroup(cfg, client)
			return
		}
		if _servicecatalogappregistryAssociateResource {
			servicecatalogappregistry_AssociateResource(cfg, client)
			return
		}
		if _servicecatalogappregistryCreateApplication {
			servicecatalogappregistry_CreateApplication(cfg, client)
			return
		}
		if _servicecatalogappregistryCreateAttributeGroup {
			servicecatalogappregistry_CreateAttributeGroup(cfg, client)
			return
		}
		if _servicecatalogappregistryDeleteApplication {
			servicecatalogappregistry_DeleteApplication(cfg, client)
			return
		}
		if _servicecatalogappregistryDeleteAttributeGroup {
			servicecatalogappregistry_DeleteAttributeGroup(cfg, client)
			return
		}
		if _servicecatalogappregistryDisassociateAttributeGroup {
			servicecatalogappregistry_DisassociateAttributeGroup(cfg, client)
			return
		}
		if _servicecatalogappregistryDisassociateResource {
			servicecatalogappregistry_DisassociateResource(cfg, client)
			return
		}
		if _servicecatalogappregistryGetApplication {
			servicecatalogappregistry_GetApplication(cfg, client)
			return
		}
		if _servicecatalogappregistryGetAssociatedResource {
			servicecatalogappregistry_GetAssociatedResource(cfg, client)
			return
		}
		if _servicecatalogappregistryGetAttributeGroup {
			servicecatalogappregistry_GetAttributeGroup(cfg, client)
			return
		}
		if _servicecatalogappregistryGetConfiguration {
			servicecatalogappregistry_GetConfiguration(cfg, client)
			return
		}
		if _servicecatalogappregistryListApplications {
			servicecatalogappregistry_ListApplications(cfg, client)
			return
		}
		if _servicecatalogappregistryListAssociatedAttributeGroups {
			servicecatalogappregistry_ListAssociatedAttributeGroups(cfg, client)
			return
		}
		if _servicecatalogappregistryListAssociatedResources {
			servicecatalogappregistry_ListAssociatedResources(cfg, client)
			return
		}
		if _servicecatalogappregistryListAttributeGroups {
			servicecatalogappregistry_ListAttributeGroups(cfg, client)
			return
		}
		if _servicecatalogappregistryListAttributeGroupsForApplication {
			servicecatalogappregistry_ListAttributeGroupsForApplication(cfg, client)
			return
		}
		if _servicecatalogappregistryListTagsForResource {
			servicecatalogappregistry_ListTagsForResource(cfg, client)
			return
		}
		if _servicecatalogappregistryPutConfiguration {
			servicecatalogappregistry_PutConfiguration(cfg, client)
			return
		}
		if _servicecatalogappregistrySyncResource {
			servicecatalogappregistry_SyncResource(cfg, client)
			return
		}
		if _servicecatalogappregistryTagResource {
			servicecatalogappregistry_TagResource(cfg, client)
			return
		}
		if _servicecatalogappregistryUntagResource {
			servicecatalogappregistry_UntagResource(cfg, client)
			return
		}
		if _servicecatalogappregistryUpdateApplication {
			servicecatalogappregistry_UpdateApplication(cfg, client)
			return
		}
		if _servicecatalogappregistryUpdateAttributeGroup {
			servicecatalogappregistry_UpdateAttributeGroup(cfg, client)
			return
		}

	},
}

var (
	_servicecatalogappregistryAssociateAttributeGroup           bool
	_servicecatalogappregistryAssociateResource                 bool
	_servicecatalogappregistryCreateApplication                 bool
	_servicecatalogappregistryCreateAttributeGroup              bool
	_servicecatalogappregistryDeleteApplication                 bool
	_servicecatalogappregistryDeleteAttributeGroup              bool
	_servicecatalogappregistryDisassociateAttributeGroup        bool
	_servicecatalogappregistryDisassociateResource              bool
	_servicecatalogappregistryGetApplication                    bool
	_servicecatalogappregistryGetAssociatedResource             bool
	_servicecatalogappregistryGetAttributeGroup                 bool
	_servicecatalogappregistryGetConfiguration                  bool
	_servicecatalogappregistryListApplications                  bool
	_servicecatalogappregistryListAssociatedAttributeGroups     bool
	_servicecatalogappregistryListAssociatedResources           bool
	_servicecatalogappregistryListAttributeGroups               bool
	_servicecatalogappregistryListAttributeGroupsForApplication bool
	_servicecatalogappregistryListTagsForResource               bool
	_servicecatalogappregistryPutConfiguration                  bool
	_servicecatalogappregistrySyncResource                      bool
	_servicecatalogappregistryTagResource                       bool
	_servicecatalogappregistryUntagResource                     bool
	_servicecatalogappregistryUpdateApplication                 bool
	_servicecatalogappregistryUpdateAttributeGroup              bool

	_servicecatalogappregistryApplication       string
	_servicecatalogappregistryAttributeGroup    string
	_servicecatalogappregistryAttributes        string
	_servicecatalogappregistryClientToken       string
	_servicecatalogappregistryConfiguration     string
	_servicecatalogappregistryDescription       string
	_servicecatalogappregistryMaxResults        string
	_servicecatalogappregistryName              string
	_servicecatalogappregistryNextToken         string
	_servicecatalogappregistryOptions           string
	_servicecatalogappregistryResource          string
	_servicecatalogappregistryResourceArn       string
	_servicecatalogappregistryResourceTagStatus string
	_servicecatalogappregistryResourceType      string
	_servicecatalogappregistryTagKeys           []string
	_servicecatalogappregistryTags              string
)

// Associates an attribute group with an application to augment the application's
// metadata with the group's attributes. This feature enables applications to be
// described with user-defined details that are machine-readable, such as
// third-party integrations.
func servicecatalogappregistry_AssociateAttributeGroup(cfg aws.Config, client *servicecatalogappregistry.Client) {
	input := &servicecatalogappregistry.AssociateAttributeGroupInput{
		// Application: *string, // Required
		// AttributeGroup: *string, // Required
	}

	if len(_servicecatalogappregistryApplication) > 0 {
		input.Application = aws.String(_servicecatalogappregistryApplication)
	}
	if len(_servicecatalogappregistryAttributeGroup) > 0 {
		input.AttributeGroup = aws.String(_servicecatalogappregistryAttributeGroup)
	}

	if resp, err := client.AssociateAttributeGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a resource with an application. The resource can be specified by
// its ARN or name. The application can be specified by ARN, ID, or name.
//
// # Minimum permissions
//
// You must have the following permissions to associate a resource using the
// OPTIONS parameter set to APPLY_APPLICATION_TAG .
//
// - tag:GetResources
//
// - tag:TagResources
//
// You must also have these additional permissions if you don't use the
// AWSServiceCatalogAppRegistryFullAccess policy. For more information, see [AWSServiceCatalogAppRegistryFullAccess] in
// the AppRegistry Administrator Guide.
//
// - resource-groups:AssociateResource
//
// - cloudformation:UpdateStack
//
// - cloudformation:DescribeStacks
//
// In addition, you must have the tagging permission defined by the Amazon Web
// Services service that creates the resource. For more information, see [TagResources]in the
// Resource Groups Tagging API Reference.
//
// [TagResources]: https://docs.aws.amazon.com/resourcegroupstagging/latest/APIReference/API_TagResources.html
// [AWSServiceCatalogAppRegistryFullAccess]: https://docs.aws.amazon.com/servicecatalog/latest/arguide/full.html
func servicecatalogappregistry_AssociateResource(cfg aws.Config, client *servicecatalogappregistry.Client) {
	input := &servicecatalogappregistry.AssociateResourceInput{
		// Application: *string, // Required
		// Resource: *string, // Required
		// ResourceType: types.ResourceType, // Required
	}

	if len(_servicecatalogappregistryApplication) > 0 {
		input.Application = aws.String(_servicecatalogappregistryApplication)
	}
	if len(_servicecatalogappregistryResource) > 0 {
		input.Resource = aws.String(_servicecatalogappregistryResource)
	}
	if len(_servicecatalogappregistryResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _servicecatalogappregistryResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogappregistryOptions) > 0 {
		if err := assignInputField(input, "Options", _servicecatalogappregistryOptions); err != nil {
			log.Errorf("invalid --options: %s", err.Error())
			return
		}
	}

	if resp, err := client.AssociateResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new application that is the top-level node in a hierarchy of related
// cloud resource abstractions.
func servicecatalogappregistry_CreateApplication(cfg aws.Config, client *servicecatalogappregistry.Client) {
	input := &servicecatalogappregistry.CreateApplicationInput{
		// ClientToken: *string, // Required
		// Name: *string, // Required
	}

	if len(_servicecatalogappregistryClientToken) > 0 {
		input.ClientToken = aws.String(_servicecatalogappregistryClientToken)
	}
	if len(_servicecatalogappregistryName) > 0 {
		input.Name = aws.String(_servicecatalogappregistryName)
	}
	if len(_servicecatalogappregistryDescription) > 0 {
		input.Description = aws.String(_servicecatalogappregistryDescription)
	}
	if len(_servicecatalogappregistryTags) > 0 {
		if err := assignInputField(input, "Tags", _servicecatalogappregistryTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new attribute group as a container for user-defined attributes. This
// feature enables users to have full control over their cloud application's
// metadata in a rich machine-readable format to facilitate integration with
// automated workflows and third-party tools.
func servicecatalogappregistry_CreateAttributeGroup(cfg aws.Config, client *servicecatalogappregistry.Client) {
	input := &servicecatalogappregistry.CreateAttributeGroupInput{
		// Attributes: *string, // Required
		// ClientToken: *string, // Required
		// Name: *string, // Required
	}

	if len(_servicecatalogappregistryAttributes) > 0 {
		input.Attributes = aws.String(_servicecatalogappregistryAttributes)
	}
	if len(_servicecatalogappregistryClientToken) > 0 {
		input.ClientToken = aws.String(_servicecatalogappregistryClientToken)
	}
	if len(_servicecatalogappregistryName) > 0 {
		input.Name = aws.String(_servicecatalogappregistryName)
	}
	if len(_servicecatalogappregistryDescription) > 0 {
		input.Description = aws.String(_servicecatalogappregistryDescription)
	}
	if len(_servicecatalogappregistryTags) > 0 {
		if err := assignInputField(input, "Tags", _servicecatalogappregistryTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAttributeGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an application that is specified either by its application ID, name, or
// ARN. All associated attribute groups and resources must be disassociated from it
// before deleting an application.
func servicecatalogappregistry_DeleteApplication(cfg aws.Config, client *servicecatalogappregistry.Client) {
	input := &servicecatalogappregistry.DeleteApplicationInput{
		// Application: *string, // Required
	}

	if len(_servicecatalogappregistryApplication) > 0 {
		input.Application = aws.String(_servicecatalogappregistryApplication)
	}

	if resp, err := client.DeleteApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an attribute group, specified either by its attribute group ID, name,
// or ARN.
func servicecatalogappregistry_DeleteAttributeGroup(cfg aws.Config, client *servicecatalogappregistry.Client) {
	input := &servicecatalogappregistry.DeleteAttributeGroupInput{
		// AttributeGroup: *string, // Required
	}

	if len(_servicecatalogappregistryAttributeGroup) > 0 {
		input.AttributeGroup = aws.String(_servicecatalogappregistryAttributeGroup)
	}

	if resp, err := client.DeleteAttributeGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates an attribute group from an application to remove the extra
// attributes contained in the attribute group from the application's metadata.
// This operation reverts AssociateAttributeGroup .
func servicecatalogappregistry_DisassociateAttributeGroup(cfg aws.Config, client *servicecatalogappregistry.Client) {
	input := &servicecatalogappregistry.DisassociateAttributeGroupInput{
		// Application: *string, // Required
		// AttributeGroup: *string, // Required
	}

	if len(_servicecatalogappregistryApplication) > 0 {
		input.Application = aws.String(_servicecatalogappregistryApplication)
	}
	if len(_servicecatalogappregistryAttributeGroup) > 0 {
		input.AttributeGroup = aws.String(_servicecatalogappregistryAttributeGroup)
	}

	if resp, err := client.DisassociateAttributeGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a resource from application. Both the resource and the
// application can be specified either by ID or name.
//
// # Minimum permissions
//
// You must have the following permissions to remove a resource that's been
// associated with an application using the APPLY_APPLICATION_TAG option for [AssociateResource].
//
// - tag:GetResources
//
// - tag:UntagResources
//
// You must also have the following permissions if you don't use the
// AWSServiceCatalogAppRegistryFullAccess policy. For more information, see [AWSServiceCatalogAppRegistryFullAccess] in
// the AppRegistry Administrator Guide.
//
// - resource-groups:DisassociateResource
//
// - cloudformation:UpdateStack
//
// - cloudformation:DescribeStacks
//
// In addition, you must have the tagging permission defined by the Amazon Web
// Services service that creates the resource. For more information, see [UntagResources]in the
// Resource Groups Tagging API Reference.
//
// [UntagResources]: https://docs.aws.amazon.com/resourcegroupstagging/latest/APIReference/API_UntTagResources.html
// [AWSServiceCatalogAppRegistryFullAccess]: https://docs.aws.amazon.com/servicecatalog/latest/arguide/full.html
// [AssociateResource]: https://docs.aws.amazon.com/servicecatalog/latest/dg/API_app-registry_AssociateResource.html
func servicecatalogappregistry_DisassociateResource(cfg aws.Config, client *servicecatalogappregistry.Client) {
	input := &servicecatalogappregistry.DisassociateResourceInput{
		// Application: *string, // Required
		// Resource: *string, // Required
		// ResourceType: types.ResourceType, // Required
	}

	if len(_servicecatalogappregistryApplication) > 0 {
		input.Application = aws.String(_servicecatalogappregistryApplication)
	}
	if len(_servicecatalogappregistryResource) > 0 {
		input.Resource = aws.String(_servicecatalogappregistryResource)
	}
	if len(_servicecatalogappregistryResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _servicecatalogappregistryResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.DisassociateResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves metadata information about one of your applications. The application
// can be specified by its ARN, ID, or name (which is unique within one account in
// one region at a given point in time). Specify by ARN or ID in automated
// workflows if you want to make sure that the exact same application is returned
// or a ResourceNotFoundException is thrown, avoiding the ABA addressing problem.
func servicecatalogappregistry_GetApplication(cfg aws.Config, client *servicecatalogappregistry.Client) {
	input := &servicecatalogappregistry.GetApplicationInput{
		// Application: *string, // Required
	}

	if len(_servicecatalogappregistryApplication) > 0 {
		input.Application = aws.String(_servicecatalogappregistryApplication)
	}

	if resp, err := client.GetApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the resource associated with the application.
func servicecatalogappregistry_GetAssociatedResource(cfg aws.Config, client *servicecatalogappregistry.Client) {
	input := &servicecatalogappregistry.GetAssociatedResourceInput{
		// Application: *string, // Required
		// Resource: *string, // Required
		// ResourceType: types.ResourceType, // Required
	}

	if len(_servicecatalogappregistryApplication) > 0 {
		input.Application = aws.String(_servicecatalogappregistryApplication)
	}
	if len(_servicecatalogappregistryResource) > 0 {
		input.Resource = aws.String(_servicecatalogappregistryResource)
	}
	if len(_servicecatalogappregistryResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _servicecatalogappregistryResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogappregistryMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _servicecatalogappregistryMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogappregistryNextToken) > 0 {
		input.NextToken = aws.String(_servicecatalogappregistryNextToken)
	}
	if len(_servicecatalogappregistryResourceTagStatus) > 0 {
		if err := assignInputField(input, "ResourceTagStatus", _servicecatalogappregistryResourceTagStatus); err != nil {
			log.Errorf("invalid --resource-tag-status: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetAssociatedResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves an attribute group by its ARN, ID, or name. The attribute group can
// be specified by its ARN, ID, or name.
func servicecatalogappregistry_GetAttributeGroup(cfg aws.Config, client *servicecatalogappregistry.Client) {
	input := &servicecatalogappregistry.GetAttributeGroupInput{
		// AttributeGroup: *string, // Required
	}

	if len(_servicecatalogappregistryAttributeGroup) > 0 {
		input.AttributeGroup = aws.String(_servicecatalogappregistryAttributeGroup)
	}

	if resp, err := client.GetAttributeGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a TagKey configuration from an account.
func servicecatalogappregistry_GetConfiguration(cfg aws.Config, client *servicecatalogappregistry.Client) {
	input := &servicecatalogappregistry.GetConfigurationInput{}

	if resp, err := client.GetConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of all of your applications. Results are paginated.
func servicecatalogappregistry_ListApplications(cfg aws.Config, client *servicecatalogappregistry.Client) {
	input := &servicecatalogappregistry.ListApplicationsInput{}

	if len(_servicecatalogappregistryMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _servicecatalogappregistryMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogappregistryNextToken) > 0 {
		input.NextToken = aws.String(_servicecatalogappregistryNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListApplications(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*servicecatalogappregistry.ListApplicationsOutput
	p := servicecatalogappregistry.NewListApplicationsPaginator(client, input)
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

// Lists all attribute groups that are associated with specified application.
// Results are paginated.
func servicecatalogappregistry_ListAssociatedAttributeGroups(cfg aws.Config, client *servicecatalogappregistry.Client) {
	input := &servicecatalogappregistry.ListAssociatedAttributeGroupsInput{
		// Application: *string, // Required
	}

	if len(_servicecatalogappregistryApplication) > 0 {
		input.Application = aws.String(_servicecatalogappregistryApplication)
	}
	if len(_servicecatalogappregistryMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _servicecatalogappregistryMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogappregistryNextToken) > 0 {
		input.NextToken = aws.String(_servicecatalogappregistryNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAssociatedAttributeGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*servicecatalogappregistry.ListAssociatedAttributeGroupsOutput
	p := servicecatalogappregistry.NewListAssociatedAttributeGroupsPaginator(client, input)
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

// Lists all of the resources that are associated with the specified application.
// Results are paginated.
//
// If you share an application, and a consumer account associates a tag query to
// the application, all of the users who can access the application can also view
// the tag values in all accounts that are associated with it using this API.
func servicecatalogappregistry_ListAssociatedResources(cfg aws.Config, client *servicecatalogappregistry.Client) {
	input := &servicecatalogappregistry.ListAssociatedResourcesInput{
		// Application: *string, // Required
	}

	if len(_servicecatalogappregistryApplication) > 0 {
		input.Application = aws.String(_servicecatalogappregistryApplication)
	}
	if len(_servicecatalogappregistryMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _servicecatalogappregistryMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogappregistryNextToken) > 0 {
		input.NextToken = aws.String(_servicecatalogappregistryNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAssociatedResources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*servicecatalogappregistry.ListAssociatedResourcesOutput
	p := servicecatalogappregistry.NewListAssociatedResourcesPaginator(client, input)
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

// Lists all attribute groups which you have access to. Results are paginated.
func servicecatalogappregistry_ListAttributeGroups(cfg aws.Config, client *servicecatalogappregistry.Client) {
	input := &servicecatalogappregistry.ListAttributeGroupsInput{}

	if len(_servicecatalogappregistryMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _servicecatalogappregistryMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogappregistryNextToken) > 0 {
		input.NextToken = aws.String(_servicecatalogappregistryNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAttributeGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*servicecatalogappregistry.ListAttributeGroupsOutput
	p := servicecatalogappregistry.NewListAttributeGroupsPaginator(client, input)
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

// Lists the details of all attribute groups associated with a specific
// application. The results display in pages.
func servicecatalogappregistry_ListAttributeGroupsForApplication(cfg aws.Config, client *servicecatalogappregistry.Client) {
	input := &servicecatalogappregistry.ListAttributeGroupsForApplicationInput{
		// Application: *string, // Required
	}

	if len(_servicecatalogappregistryApplication) > 0 {
		input.Application = aws.String(_servicecatalogappregistryApplication)
	}
	if len(_servicecatalogappregistryMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _servicecatalogappregistryMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogappregistryNextToken) > 0 {
		input.NextToken = aws.String(_servicecatalogappregistryNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAttributeGroupsForApplication(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*servicecatalogappregistry.ListAttributeGroupsForApplicationOutput
	p := servicecatalogappregistry.NewListAttributeGroupsForApplicationPaginator(client, input)
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

// Lists all of the tags on the resource.
func servicecatalogappregistry_ListTagsForResource(cfg aws.Config, client *servicecatalogappregistry.Client) {
	input := &servicecatalogappregistry.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_servicecatalogappregistryResourceArn) > 0 {
		input.ResourceArn = aws.String(_servicecatalogappregistryResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a TagKey configuration to an account.
func servicecatalogappregistry_PutConfiguration(cfg aws.Config, client *servicecatalogappregistry.Client) {
	input := &servicecatalogappregistry.PutConfigurationInput{
		// Configuration: *types.AppRegistryConfiguration, // Required
	}

	if len(_servicecatalogappregistryConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _servicecatalogappregistryConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Syncs the resource with current AppRegistry records.
// Specifically, the resource’s AppRegistry system tags sync with its associated
// application. We remove the resource's AppRegistry system tags if it does not
// associate with the application. The caller must have permissions to read and
// update the resource.
func servicecatalogappregistry_SyncResource(cfg aws.Config, client *servicecatalogappregistry.Client) {
	input := &servicecatalogappregistry.SyncResourceInput{
		// Resource: *string, // Required
		// ResourceType: types.ResourceType, // Required
	}

	if len(_servicecatalogappregistryResource) > 0 {
		input.Resource = aws.String(_servicecatalogappregistryResource)
	}
	if len(_servicecatalogappregistryResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _servicecatalogappregistryResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.SyncResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assigns one or more tags (key-value pairs) to the specified resource.
// Each tag consists of a key and an optional value. If a tag with the same key is
// already associated with the resource, this action updates its value.
//
// This operation returns an empty response if the call was successful.
func servicecatalogappregistry_TagResource(cfg aws.Config, client *servicecatalogappregistry.Client) {
	input := &servicecatalogappregistry.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_servicecatalogappregistryResourceArn) > 0 {
		input.ResourceArn = aws.String(_servicecatalogappregistryResourceArn)
	}
	if len(_servicecatalogappregistryTags) > 0 {
		if err := assignInputField(input, "Tags", _servicecatalogappregistryTags); err != nil {
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

// Removes tags from a resource.
// This operation returns an empty response if the call was successful.
func servicecatalogappregistry_UntagResource(cfg aws.Config, client *servicecatalogappregistry.Client) {
	input := &servicecatalogappregistry.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_servicecatalogappregistryResourceArn) > 0 {
		input.ResourceArn = aws.String(_servicecatalogappregistryResourceArn)
	}
	if len(_servicecatalogappregistryTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _servicecatalogappregistryTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing application with new attributes.
func servicecatalogappregistry_UpdateApplication(cfg aws.Config, client *servicecatalogappregistry.Client) {
	input := &servicecatalogappregistry.UpdateApplicationInput{
		// Application: *string, // Required
	}

	if len(_servicecatalogappregistryApplication) > 0 {
		input.Application = aws.String(_servicecatalogappregistryApplication)
	}
	if len(_servicecatalogappregistryDescription) > 0 {
		input.Description = aws.String(_servicecatalogappregistryDescription)
	}
	if len(_servicecatalogappregistryName) > 0 {
		input.Name = aws.String(_servicecatalogappregistryName)
	}

	if resp, err := client.UpdateApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing attribute group with new details.
func servicecatalogappregistry_UpdateAttributeGroup(cfg aws.Config, client *servicecatalogappregistry.Client) {
	input := &servicecatalogappregistry.UpdateAttributeGroupInput{
		// AttributeGroup: *string, // Required
	}

	if len(_servicecatalogappregistryAttributeGroup) > 0 {
		input.AttributeGroup = aws.String(_servicecatalogappregistryAttributeGroup)
	}
	if len(_servicecatalogappregistryAttributes) > 0 {
		input.Attributes = aws.String(_servicecatalogappregistryAttributes)
	}
	if len(_servicecatalogappregistryDescription) > 0 {
		input.Description = aws.String(_servicecatalogappregistryDescription)
	}
	if len(_servicecatalogappregistryName) > 0 {
		input.Name = aws.String(_servicecatalogappregistryName)
	}

	if resp, err := client.UpdateAttributeGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_servicecatalogappregistryCmd)
	_servicecatalogappregistryCmd.Flags().SortFlags = false

	_servicecatalogappregistryCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_servicecatalogappregistryCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_servicecatalogappregistryCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_servicecatalogappregistryCmd.Flags().StringVarP(&_servicecatalogappregistryApplication, "application", "", "", "Application")
	_servicecatalogappregistryCmd.Flags().StringVarP(&_servicecatalogappregistryAttributeGroup, "attribute-group", "", "", "Attribute Group")
	_servicecatalogappregistryCmd.Flags().StringVarP(&_servicecatalogappregistryAttributes, "attributes", "", "", "Attributes")
	_servicecatalogappregistryCmd.Flags().StringVarP(&_servicecatalogappregistryClientToken, "client-token", "", "", "Client Token")
	_servicecatalogappregistryCmd.Flags().StringVarP(&_servicecatalogappregistryConfiguration, "configuration", "", "", "Configuration")
	_servicecatalogappregistryCmd.Flags().StringVarP(&_servicecatalogappregistryDescription, "description", "", "", "Description")
	_servicecatalogappregistryCmd.Flags().StringVarP(&_servicecatalogappregistryMaxResults, "max-results", "", "", "Max Results")
	_servicecatalogappregistryCmd.Flags().StringVarP(&_servicecatalogappregistryName, "name", "", "", "Name")
	_servicecatalogappregistryCmd.Flags().StringVarP(&_servicecatalogappregistryNextToken, "next-token", "", "", "Next Token")
	_servicecatalogappregistryCmd.Flags().StringVarP(&_servicecatalogappregistryOptions, "options", "", "", "Options")
	_servicecatalogappregistryCmd.Flags().StringVarP(&_servicecatalogappregistryResource, "resource", "", "", "Resource")
	_servicecatalogappregistryCmd.Flags().StringVarP(&_servicecatalogappregistryResourceArn, "resource-arn", "", "", "Resource ARN")
	_servicecatalogappregistryCmd.Flags().StringVarP(&_servicecatalogappregistryResourceTagStatus, "resource-tag-status", "", "", "Resource Tag Status")
	_servicecatalogappregistryCmd.Flags().StringVarP(&_servicecatalogappregistryResourceType, "resource-type", "", "", "Resource Type")
	_servicecatalogappregistryCmd.Flags().StringSliceVarP(&_servicecatalogappregistryTagKeys, "tag-keys", "", nil, "Tag Keys")
	_servicecatalogappregistryCmd.Flags().StringVarP(&_servicecatalogappregistryTags, "tags", "", "", "Tags")

	_servicecatalogappregistryCmd.Flags().BoolVarP(&_servicecatalogappregistryAssociateAttributeGroup, "associate-attribute-group", "", false, "Associate Attribute Group")
	_servicecatalogappregistryCmd.Flags().BoolVarP(&_servicecatalogappregistryAssociateResource, "associate-resource", "", false, "Associate Resource")
	_servicecatalogappregistryCmd.Flags().BoolVarP(&_servicecatalogappregistryCreateApplication, "create-application", "", false, "Create Application")
	_servicecatalogappregistryCmd.Flags().BoolVarP(&_servicecatalogappregistryCreateAttributeGroup, "create-attribute-group", "", false, "Create Attribute Group")
	_servicecatalogappregistryCmd.Flags().BoolVarP(&_servicecatalogappregistryDeleteApplication, "delete-application", "", false, "Delete Application")
	_servicecatalogappregistryCmd.Flags().BoolVarP(&_servicecatalogappregistryDeleteAttributeGroup, "delete-attribute-group", "", false, "Delete Attribute Group")
	_servicecatalogappregistryCmd.Flags().BoolVarP(&_servicecatalogappregistryDisassociateAttributeGroup, "disassociate-attribute-group", "", false, "Disassociate Attribute Group")
	_servicecatalogappregistryCmd.Flags().BoolVarP(&_servicecatalogappregistryDisassociateResource, "disassociate-resource", "", false, "Disassociate Resource")
	_servicecatalogappregistryCmd.Flags().BoolVarP(&_servicecatalogappregistryGetApplication, "get-application", "", false, "Get Application")
	_servicecatalogappregistryCmd.Flags().BoolVarP(&_servicecatalogappregistryGetAssociatedResource, "get-associated-resource", "", false, "Get Associated Resource")
	_servicecatalogappregistryCmd.Flags().BoolVarP(&_servicecatalogappregistryGetAttributeGroup, "get-attribute-group", "", false, "Get Attribute Group")
	_servicecatalogappregistryCmd.Flags().BoolVarP(&_servicecatalogappregistryGetConfiguration, "get-configuration", "", false, "Get Configuration")
	_servicecatalogappregistryCmd.Flags().BoolVarP(&_servicecatalogappregistryListApplications, "list-applications", "", false, "List Applications")
	_servicecatalogappregistryCmd.Flags().BoolVarP(&_servicecatalogappregistryListAssociatedAttributeGroups, "list-associated-attribute-groups", "", false, "List Associated Attribute Groups")
	_servicecatalogappregistryCmd.Flags().BoolVarP(&_servicecatalogappregistryListAssociatedResources, "list-associated-resources", "", false, "List Associated Resources")
	_servicecatalogappregistryCmd.Flags().BoolVarP(&_servicecatalogappregistryListAttributeGroups, "list-attribute-groups", "", false, "List Attribute Groups")
	_servicecatalogappregistryCmd.Flags().BoolVarP(&_servicecatalogappregistryListAttributeGroupsForApplication, "list-attribute-groups-for-application", "", false, "List Attribute Groups For Application")
	_servicecatalogappregistryCmd.Flags().BoolVarP(&_servicecatalogappregistryListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_servicecatalogappregistryCmd.Flags().BoolVarP(&_servicecatalogappregistryPutConfiguration, "put-configuration", "", false, "Put Configuration")
	_servicecatalogappregistryCmd.Flags().BoolVarP(&_servicecatalogappregistrySyncResource, "sync-resource", "", false, "Sync Resource")
	_servicecatalogappregistryCmd.Flags().BoolVarP(&_servicecatalogappregistryTagResource, "tag-resource", "", false, "Tag Resource")
	_servicecatalogappregistryCmd.Flags().BoolVarP(&_servicecatalogappregistryUntagResource, "untag-resource", "", false, "Untag Resource")
	_servicecatalogappregistryCmd.Flags().BoolVarP(&_servicecatalogappregistryUpdateApplication, "update-application", "", false, "Update Application")
	_servicecatalogappregistryCmd.Flags().BoolVarP(&_servicecatalogappregistryUpdateAttributeGroup, "update-attribute-group", "", false, "Update Attribute Group")

}
