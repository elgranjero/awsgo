package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/connectcases"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// connectcasesCmd represents the connectcases command
var _connectcasesCmd = &cobra.Command{
	Use:   "connectcases",
	Short: "AWS connectcases CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := connectcases.NewFromConfig(cfg)
		if _connectcasesBatchGetCaseRule {
			connectcases_BatchGetCaseRule(cfg, client)
			return
		}
		if _connectcasesBatchGetField {
			connectcases_BatchGetField(cfg, client)
			return
		}
		if _connectcasesBatchPutFieldOptions {
			connectcases_BatchPutFieldOptions(cfg, client)
			return
		}
		if _connectcasesCreateCase {
			connectcases_CreateCase(cfg, client)
			return
		}
		if _connectcasesCreateCaseRule {
			connectcases_CreateCaseRule(cfg, client)
			return
		}
		if _connectcasesCreateDomain {
			connectcases_CreateDomain(cfg, client)
			return
		}
		if _connectcasesCreateField {
			connectcases_CreateField(cfg, client)
			return
		}
		if _connectcasesCreateLayout {
			connectcases_CreateLayout(cfg, client)
			return
		}
		if _connectcasesCreateRelatedItem {
			connectcases_CreateRelatedItem(cfg, client)
			return
		}
		if _connectcasesCreateTemplate {
			connectcases_CreateTemplate(cfg, client)
			return
		}
		if _connectcasesDeleteCase {
			connectcases_DeleteCase(cfg, client)
			return
		}
		if _connectcasesDeleteCaseRule {
			connectcases_DeleteCaseRule(cfg, client)
			return
		}
		if _connectcasesDeleteDomain {
			connectcases_DeleteDomain(cfg, client)
			return
		}
		if _connectcasesDeleteField {
			connectcases_DeleteField(cfg, client)
			return
		}
		if _connectcasesDeleteLayout {
			connectcases_DeleteLayout(cfg, client)
			return
		}
		if _connectcasesDeleteRelatedItem {
			connectcases_DeleteRelatedItem(cfg, client)
			return
		}
		if _connectcasesDeleteTemplate {
			connectcases_DeleteTemplate(cfg, client)
			return
		}
		if _connectcasesGetCase {
			connectcases_GetCase(cfg, client)
			return
		}
		if _connectcasesGetCaseAuditEvents {
			connectcases_GetCaseAuditEvents(cfg, client)
			return
		}
		if _connectcasesGetCaseEventConfiguration {
			connectcases_GetCaseEventConfiguration(cfg, client)
			return
		}
		if _connectcasesGetDomain {
			connectcases_GetDomain(cfg, client)
			return
		}
		if _connectcasesGetLayout {
			connectcases_GetLayout(cfg, client)
			return
		}
		if _connectcasesGetTemplate {
			connectcases_GetTemplate(cfg, client)
			return
		}
		if _connectcasesListCaseRules {
			connectcases_ListCaseRules(cfg, client)
			return
		}
		if _connectcasesListCasesForContact {
			connectcases_ListCasesForContact(cfg, client)
			return
		}
		if _connectcasesListDomains {
			connectcases_ListDomains(cfg, client)
			return
		}
		if _connectcasesListFieldOptions {
			connectcases_ListFieldOptions(cfg, client)
			return
		}
		if _connectcasesListFields {
			connectcases_ListFields(cfg, client)
			return
		}
		if _connectcasesListLayouts {
			connectcases_ListLayouts(cfg, client)
			return
		}
		if _connectcasesListTagsForResource {
			connectcases_ListTagsForResource(cfg, client)
			return
		}
		if _connectcasesListTemplates {
			connectcases_ListTemplates(cfg, client)
			return
		}
		if _connectcasesPutCaseEventConfiguration {
			connectcases_PutCaseEventConfiguration(cfg, client)
			return
		}
		if _connectcasesSearchAllRelatedItems {
			connectcases_SearchAllRelatedItems(cfg, client)
			return
		}
		if _connectcasesSearchCases {
			connectcases_SearchCases(cfg, client)
			return
		}
		if _connectcasesSearchRelatedItems {
			connectcases_SearchRelatedItems(cfg, client)
			return
		}
		if _connectcasesTagResource {
			connectcases_TagResource(cfg, client)
			return
		}
		if _connectcasesUntagResource {
			connectcases_UntagResource(cfg, client)
			return
		}
		if _connectcasesUpdateCase {
			connectcases_UpdateCase(cfg, client)
			return
		}
		if _connectcasesUpdateCaseRule {
			connectcases_UpdateCaseRule(cfg, client)
			return
		}
		if _connectcasesUpdateField {
			connectcases_UpdateField(cfg, client)
			return
		}
		if _connectcasesUpdateLayout {
			connectcases_UpdateLayout(cfg, client)
			return
		}
		if _connectcasesUpdateTemplate {
			connectcases_UpdateTemplate(cfg, client)
			return
		}

	},
}

var (
	_connectcasesBatchGetCaseRule          bool
	_connectcasesBatchGetField             bool
	_connectcasesBatchPutFieldOptions      bool
	_connectcasesCreateCase                bool
	_connectcasesCreateCaseRule            bool
	_connectcasesCreateDomain              bool
	_connectcasesCreateField               bool
	_connectcasesCreateLayout              bool
	_connectcasesCreateRelatedItem         bool
	_connectcasesCreateTemplate            bool
	_connectcasesDeleteCase                bool
	_connectcasesDeleteCaseRule            bool
	_connectcasesDeleteDomain              bool
	_connectcasesDeleteField               bool
	_connectcasesDeleteLayout              bool
	_connectcasesDeleteRelatedItem         bool
	_connectcasesDeleteTemplate            bool
	_connectcasesGetCase                   bool
	_connectcasesGetCaseAuditEvents        bool
	_connectcasesGetCaseEventConfiguration bool
	_connectcasesGetDomain                 bool
	_connectcasesGetLayout                 bool
	_connectcasesGetTemplate               bool
	_connectcasesListCaseRules             bool
	_connectcasesListCasesForContact       bool
	_connectcasesListDomains               bool
	_connectcasesListFieldOptions          bool
	_connectcasesListFields                bool
	_connectcasesListLayouts               bool
	_connectcasesListTagsForResource       bool
	_connectcasesListTemplates             bool
	_connectcasesPutCaseEventConfiguration bool
	_connectcasesSearchAllRelatedItems     bool
	_connectcasesSearchCases               bool
	_connectcasesSearchRelatedItems        bool
	_connectcasesTagResource               bool
	_connectcasesUntagResource             bool
	_connectcasesUpdateCase                bool
	_connectcasesUpdateCaseRule            bool
	_connectcasesUpdateField               bool
	_connectcasesUpdateLayout              bool
	_connectcasesUpdateTemplate            bool

	_connectcasesArn                          string
	_connectcasesAttributes                   string
	_connectcasesCaseId                       string
	_connectcasesCaseRuleId                   string
	_connectcasesCaseRules                    string
	_connectcasesClientToken                  string
	_connectcasesContactArn                   string
	_connectcasesContent                      string
	_connectcasesDescription                  string
	_connectcasesDomainId                     string
	_connectcasesEventBridge                  string
	_connectcasesFieldId                      string
	_connectcasesFields                       string
	_connectcasesFilter                       string
	_connectcasesFilters                      string
	_connectcasesLayoutConfiguration          string
	_connectcasesLayoutId                     string
	_connectcasesMaxResults                   string
	_connectcasesName                         string
	_connectcasesNextToken                    string
	_connectcasesOptions                      string
	_connectcasesPerformedBy                  string
	_connectcasesRelatedItemId                string
	_connectcasesRequiredFields               string
	_connectcasesRule                         string
	_connectcasesRules                        string
	_connectcasesSearchTerm                   string
	_connectcasesSorts                        string
	_connectcasesStatus                       string
	_connectcasesTagKeys                      []string
	_connectcasesTagPropagationConfigurations string
	_connectcasesTags                         string
	_connectcasesTemplateId                   string
	_connectcasesType                         string
	_connectcasesValues                       []string
)

// Gets a batch of case rules. In the Amazon Connect admin website, case rules are
// known as case field conditions. For more information about case field
// conditions, see [Add case field conditions to a case template].
//
// [Add case field conditions to a case template]: https://docs.aws.amazon.com/connect/latest/adminguide/case-field-conditions.html
func connectcases_BatchGetCaseRule(cfg aws.Config, client *connectcases.Client) {
	input := &connectcases.BatchGetCaseRuleInput{
		// CaseRules: []types.CaseRuleIdentifier, // Required
		// DomainId: *string, // Required
	}

	if len(_connectcasesCaseRules) > 0 {
		if err := assignInputField(input, "CaseRules", _connectcasesCaseRules); err != nil {
			log.Errorf("invalid --case-rules: %s", err.Error())
			return
		}
	}
	if len(_connectcasesDomainId) > 0 {
		input.DomainId = aws.String(_connectcasesDomainId)
	}

	if resp, err := client.BatchGetCaseRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the description for the list of fields in the request parameters.
func connectcases_BatchGetField(cfg aws.Config, client *connectcases.Client) {
	input := &connectcases.BatchGetFieldInput{
		// DomainId: *string, // Required
		// Fields: []types.FieldIdentifier, // Required
	}

	if len(_connectcasesDomainId) > 0 {
		input.DomainId = aws.String(_connectcasesDomainId)
	}
	if len(_connectcasesFields) > 0 {
		if err := assignInputField(input, "Fields", _connectcasesFields); err != nil {
			log.Errorf("invalid --fields: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchGetField(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates and updates a set of field options for a single select field in a Cases
// domain.
func connectcases_BatchPutFieldOptions(cfg aws.Config, client *connectcases.Client) {
	input := &connectcases.BatchPutFieldOptionsInput{
		// DomainId: *string, // Required
		// FieldId: *string, // Required
		// Options: []types.FieldOption, // Required
	}

	if len(_connectcasesDomainId) > 0 {
		input.DomainId = aws.String(_connectcasesDomainId)
	}
	if len(_connectcasesFieldId) > 0 {
		input.FieldId = aws.String(_connectcasesFieldId)
	}
	if len(_connectcasesOptions) > 0 {
		if err := assignInputField(input, "Options", _connectcasesOptions); err != nil {
			log.Errorf("invalid --options: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchPutFieldOptions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// If you provide a value for PerformedBy.UserArn you must also have [connect:DescribeUser] permission
// on the User ARN resource that you provide
//
// Creates a case in the specified Cases domain. Case system and custom fields are
// taken as an array id/value pairs with a declared data types.
//
// When creating a case from a template that has tag propagation configurations,
// the specified tags are automatically applied to the case.
//
// The following fields are required when creating a case:
//
// - customer_id - You must provide the full customer profile ARN in this format:
// arn:aws:profile:your_AWS_Region:your_AWS_account
// ID:domains/your_profiles_domain_name/profiles/profile_ID
//
// - title
//
// [connect:DescribeUser]: https://docs.aws.amazon.com/connect/latest/APIReference/API_DescribeUser.html
func connectcases_CreateCase(cfg aws.Config, client *connectcases.Client) {
	input := &connectcases.CreateCaseInput{
		// DomainId: *string, // Required
		// Fields: []types.FieldValue, // Required
		// TemplateId: *string, // Required
	}

	if len(_connectcasesDomainId) > 0 {
		input.DomainId = aws.String(_connectcasesDomainId)
	}
	if len(_connectcasesFields) > 0 {
		if err := assignInputField(input, "Fields", _connectcasesFields); err != nil {
			log.Errorf("invalid --fields: %s", err.Error())
			return
		}
	}
	if len(_connectcasesTemplateId) > 0 {
		input.TemplateId = aws.String(_connectcasesTemplateId)
	}
	if len(_connectcasesClientToken) > 0 {
		input.ClientToken = aws.String(_connectcasesClientToken)
	}
	if len(_connectcasesPerformedBy) > 0 {
		if err := assignInputField(input, "PerformedBy", _connectcasesPerformedBy); err != nil {
			log.Errorf("invalid --performed-by: %s", err.Error())
			return
		}
	}
	if len(_connectcasesTags) > 0 {
		if err := assignInputField(input, "Tags", _connectcasesTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new case rule. In the Amazon Connect admin website, case rules are
// known as case field conditions. For more information about case field
// conditions, see [Add case field conditions to a case template].
//
// [Add case field conditions to a case template]: https://docs.aws.amazon.com/connect/latest/adminguide/case-field-conditions.html
func connectcases_CreateCaseRule(cfg aws.Config, client *connectcases.Client) {
	input := &connectcases.CreateCaseRuleInput{
		// DomainId: *string, // Required
		// Name: *string, // Required
		// Rule: types.CaseRuleDetails, // Required
	}

	if len(_connectcasesDomainId) > 0 {
		input.DomainId = aws.String(_connectcasesDomainId)
	}
	if len(_connectcasesName) > 0 {
		input.Name = aws.String(_connectcasesName)
	}
	if len(_connectcasesRule) > 0 {
		if err := assignInputField(input, "Rule", _connectcasesRule); err != nil {
			log.Errorf("invalid --rule: %s", err.Error())
			return
		}
	}
	if len(_connectcasesDescription) > 0 {
		input.Description = aws.String(_connectcasesDescription)
	}

	if resp, err := client.CreateCaseRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a domain, which is a container for all case data, such as cases,
// fields, templates and layouts. Each Amazon Connect instance can be associated
// with only one Cases domain.
//
// This will not associate your connect instance to Cases domain. Instead, use the
// Amazon Connect [CreateIntegrationAssociation]API. You need specific IAM permissions to successfully associate
// the Cases domain. For more information, see [Onboard to Cases].
//
// [CreateIntegrationAssociation]: https://docs.aws.amazon.com/connect/latest/APIReference/API_CreateIntegrationAssociation.html
// [Onboard to Cases]: https://docs.aws.amazon.com/connect/latest/adminguide/required-permissions-iam-cases.html#onboard-cases-iam
func connectcases_CreateDomain(cfg aws.Config, client *connectcases.Client) {
	input := &connectcases.CreateDomainInput{
		// Name: *string, // Required
	}

	if len(_connectcasesName) > 0 {
		input.Name = aws.String(_connectcasesName)
	}

	if resp, err := client.CreateDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a field in the Cases domain. This field is used to define the case
// object model (that is, defines what data can be captured on cases) in a Cases
// domain.
func connectcases_CreateField(cfg aws.Config, client *connectcases.Client) {
	input := &connectcases.CreateFieldInput{
		// DomainId: *string, // Required
		// Name: *string, // Required
		// Type: types.FieldType, // Required
	}

	if len(_connectcasesDomainId) > 0 {
		input.DomainId = aws.String(_connectcasesDomainId)
	}
	if len(_connectcasesName) > 0 {
		input.Name = aws.String(_connectcasesName)
	}
	if len(_connectcasesType) > 0 {
		if err := assignInputField(input, "Type", _connectcasesType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_connectcasesAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _connectcasesAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_connectcasesDescription) > 0 {
		input.Description = aws.String(_connectcasesDescription)
	}

	if resp, err := client.CreateField(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a layout in the Cases domain. Layouts define the following
// configuration in the top section and More Info tab of the Cases user interface:
//
// - Fields to display to the users
//
// - Field ordering
//
// Title and Status fields cannot be part of layouts since they are not
// configurable.
func connectcases_CreateLayout(cfg aws.Config, client *connectcases.Client) {
	input := &connectcases.CreateLayoutInput{
		// Content: types.LayoutContent, // Required
		// DomainId: *string, // Required
		// Name: *string, // Required
	}

	if len(_connectcasesContent) > 0 {
		if err := assignInputField(input, "Content", _connectcasesContent); err != nil {
			log.Errorf("invalid --content: %s", err.Error())
			return
		}
	}
	if len(_connectcasesDomainId) > 0 {
		input.DomainId = aws.String(_connectcasesDomainId)
	}
	if len(_connectcasesName) > 0 {
		input.Name = aws.String(_connectcasesName)
	}

	if resp, err := client.CreateLayout(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a related item (comments, tasks, and contacts) and associates it with a
// case.
//
// There's a quota for the number of fields allowed in a Custom type related item.
// See [Amazon Connect Cases quotas].
//
// # Use cases
//
// Following are examples of related items that you may want to associate with a
// case:
//
// - Related contacts, such as calls, chats, emails tasks
//
// - Comments, for agent notes
//
// - SLAs, to capture target resolution goals
//
// - Cases, to capture related Amazon Connect Cases
//
// - Files, such as policy documentation or customer-provided attachments
//
// - Custom related items, which provide flexibility for you to define related
// items that such as bookings, orders, products, notices, and more
//
// # Important things to know
//
// - If you are associating a contact to a case by passing in Contact for a type
// , you must have [DescribeContact]permission on the ARN of the contact that you provide in
// content.contact.contactArn .
//
// - A Related Item is a resource that is associated with a case. It may or may
// not have an external identifier linking it to an external resource (for example,
// a contactArn ). All Related Items have their own internal identifier, the
// relatedItemArn . Examples of related items include comments and contacts .
//
// - If you provide a value for performedBy.userArn you must also have [DescribeUser]
// permission on the ARN of the user that you provide.
//
// - The type field is reserved for internal use only.
//
// Endpoints: See [Amazon Connect endpoints and quotas].
//
// [DescribeUser]: https://docs.aws.amazon.com/connect/latest/APIReference/API_DescribeUser.html
// [Amazon Connect Cases quotas]: https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-service-limits.html#cases-quotas
// [DescribeContact]: https://docs.aws.amazon.com/connect/latest/APIReference/API_DescribeContact.html
// [Amazon Connect endpoints and quotas]: https://docs.aws.amazon.com/general/latest/gr/connect_region.html
func connectcases_CreateRelatedItem(cfg aws.Config, client *connectcases.Client) {
	input := &connectcases.CreateRelatedItemInput{
		// CaseId: *string, // Required
		// Content: types.RelatedItemInputContent, // Required
		// DomainId: *string, // Required
		// Type: types.RelatedItemType, // Required
	}

	if len(_connectcasesCaseId) > 0 {
		input.CaseId = aws.String(_connectcasesCaseId)
	}
	if len(_connectcasesContent) > 0 {
		if err := assignInputField(input, "Content", _connectcasesContent); err != nil {
			log.Errorf("invalid --content: %s", err.Error())
			return
		}
	}
	if len(_connectcasesDomainId) > 0 {
		input.DomainId = aws.String(_connectcasesDomainId)
	}
	if len(_connectcasesType) > 0 {
		if err := assignInputField(input, "Type", _connectcasesType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_connectcasesPerformedBy) > 0 {
		if err := assignInputField(input, "PerformedBy", _connectcasesPerformedBy); err != nil {
			log.Errorf("invalid --performed-by: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRelatedItem(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a template in the Cases domain. This template is used to define the
// case object model (that is, to define what data can be captured on cases) in a
// Cases domain. A template must have a unique name within a domain, and it must
// reference existing field IDs and layout IDs. Additionally, multiple fields with
// same IDs are not allowed within the same Template. A template can be either
// Active or Inactive, as indicated by its status. Inactive templates cannot be
// used to create cases.
//
// Other template APIs are:
//
// [DeleteTemplate]
//
// [GetTemplate]
//
// [ListTemplates]
//
// [UpdateTemplate]
//
// [DeleteTemplate]: https://docs.aws.amazon.com/connect/latest/APIReference/API_connect-cases_DeleteTemplate.html
// [ListTemplates]: https://docs.aws.amazon.com/connect/latest/APIReference/API_connect-cases_ListTemplates.html
// [UpdateTemplate]: https://docs.aws.amazon.com/connect/latest/APIReference/API_connect-cases_UpdateTemplate.html
// [GetTemplate]: https://docs.aws.amazon.com/connect/latest/APIReference/API_connect-cases_GetTemplate.html
func connectcases_CreateTemplate(cfg aws.Config, client *connectcases.Client) {
	input := &connectcases.CreateTemplateInput{
		// DomainId: *string, // Required
		// Name: *string, // Required
	}

	if len(_connectcasesDomainId) > 0 {
		input.DomainId = aws.String(_connectcasesDomainId)
	}
	if len(_connectcasesName) > 0 {
		input.Name = aws.String(_connectcasesName)
	}
	if len(_connectcasesDescription) > 0 {
		input.Description = aws.String(_connectcasesDescription)
	}
	if len(_connectcasesLayoutConfiguration) > 0 {
		if err := assignInputField(input, "LayoutConfiguration", _connectcasesLayoutConfiguration); err != nil {
			log.Errorf("invalid --layout-configuration: %s", err.Error())
			return
		}
	}
	if len(_connectcasesRequiredFields) > 0 {
		if err := assignInputField(input, "RequiredFields", _connectcasesRequiredFields); err != nil {
			log.Errorf("invalid --required-fields: %s", err.Error())
			return
		}
	}
	if len(_connectcasesRules) > 0 {
		if err := assignInputField(input, "Rules", _connectcasesRules); err != nil {
			log.Errorf("invalid --rules: %s", err.Error())
			return
		}
	}
	if len(_connectcasesStatus) > 0 {
		if err := assignInputField(input, "Status", _connectcasesStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_connectcasesTagPropagationConfigurations) > 0 {
		if err := assignInputField(input, "TagPropagationConfigurations", _connectcasesTagPropagationConfigurations); err != nil {
			log.Errorf("invalid --tag-propagation-configurations: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The DeleteCase API permanently deletes a case and all its associated resources
// from the cases data store. After a successful deletion, you cannot:
//
// - Retrieve related items
//
// - Access audit history
//
// - Perform any operations that require the CaseID
//
// This action is irreversible. After you delete a case, you cannot recover its
// data.
func connectcases_DeleteCase(cfg aws.Config, client *connectcases.Client) {
	input := &connectcases.DeleteCaseInput{
		// CaseId: *string, // Required
		// DomainId: *string, // Required
	}

	if len(_connectcasesCaseId) > 0 {
		input.CaseId = aws.String(_connectcasesCaseId)
	}
	if len(_connectcasesDomainId) > 0 {
		input.DomainId = aws.String(_connectcasesDomainId)
	}

	if resp, err := client.DeleteCase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a case rule. In the Amazon Connect admin website, case rules are known
// as case field conditions. For more information about case field conditions, see [Add case field conditions to a case template]
// .
//
// [Add case field conditions to a case template]: https://docs.aws.amazon.com/connect/latest/adminguide/case-field-conditions.html
func connectcases_DeleteCaseRule(cfg aws.Config, client *connectcases.Client) {
	input := &connectcases.DeleteCaseRuleInput{
		// CaseRuleId: *string, // Required
		// DomainId: *string, // Required
	}

	if len(_connectcasesCaseRuleId) > 0 {
		input.CaseRuleId = aws.String(_connectcasesCaseRuleId)
	}
	if len(_connectcasesDomainId) > 0 {
		input.DomainId = aws.String(_connectcasesDomainId)
	}

	if resp, err := client.DeleteCaseRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Cases domain.
// After deleting your domain you must disassociate the deleted domain from your
// Amazon Connect instance with another API call before being able to use Cases
// again with this Amazon Connect instance. See [DeleteIntegrationAssociation].
//
// [DeleteIntegrationAssociation]: https://docs.aws.amazon.com/connect/latest/APIReference/API_DeleteIntegrationAssociation.html
func connectcases_DeleteDomain(cfg aws.Config, client *connectcases.Client) {
	input := &connectcases.DeleteDomainInput{
		// DomainId: *string, // Required
	}

	if len(_connectcasesDomainId) > 0 {
		input.DomainId = aws.String(_connectcasesDomainId)
	}

	if resp, err := client.DeleteDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a field from a cases template.
// After a field is deleted:
//
// - You can still retrieve the field by calling BatchGetField .
//
// - You cannot update a deleted field by calling UpdateField ; it throws a
// ValidationException .
//
// - Deleted fields are not included in the ListFields response.
//
// - Calling CreateCase with a deleted field throws a ValidationException
// denoting which field identifiers in the request have been deleted.
//
// - Calling GetCase with a deleted field identifier returns the deleted field's
// value if one exists.
//
// - Calling UpdateCase with a deleted field ID throws a ValidationException if
// the case does not already contain a value for the deleted field. Otherwise it
// succeeds, allowing you to update or remove (using emptyValue: {} ) the field's
// value from the case.
//
// - GetTemplate does not return field IDs for deleted fields.
//
// - GetLayout does not return field IDs for deleted fields.
//
// - Calling SearchCases with the deleted field ID as a filter returns any cases
// that have a value for the deleted field that matches the filter criteria.
//
// - Calling SearchCases with a searchTerm value that matches a deleted field's
// value on a case returns the case in the response.
//
// - Calling BatchPutFieldOptions with a deleted field ID throw a
// ValidationException .
//
// - Calling GetCaseEventConfiguration does not return field IDs for deleted
// fields.
func connectcases_DeleteField(cfg aws.Config, client *connectcases.Client) {
	input := &connectcases.DeleteFieldInput{
		// DomainId: *string, // Required
		// FieldId: *string, // Required
	}

	if len(_connectcasesDomainId) > 0 {
		input.DomainId = aws.String(_connectcasesDomainId)
	}
	if len(_connectcasesFieldId) > 0 {
		input.FieldId = aws.String(_connectcasesFieldId)
	}

	if resp, err := client.DeleteField(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a layout from a cases template. You can delete up to 100 layouts per
// domain.
//
// After a layout is deleted:
//
// - You can still retrieve the layout by calling GetLayout .
//
// - You cannot update a deleted layout by calling UpdateLayout ; it throws a
// ValidationException .
//
// - Deleted layouts are not included in the ListLayouts response.
func connectcases_DeleteLayout(cfg aws.Config, client *connectcases.Client) {
	input := &connectcases.DeleteLayoutInput{
		// DomainId: *string, // Required
		// LayoutId: *string, // Required
	}

	if len(_connectcasesDomainId) > 0 {
		input.DomainId = aws.String(_connectcasesDomainId)
	}
	if len(_connectcasesLayoutId) > 0 {
		input.LayoutId = aws.String(_connectcasesLayoutId)
	}

	if resp, err := client.DeleteLayout(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the related item resource under a case.
// This API cannot be used on a FILE type related attachment. To delete this type
// of file, use the [DeleteAttachedFile]API
//
// [DeleteAttachedFile]: https://docs.aws.amazon.com/connect/latest/APIReference/API_DeleteAttachedFile.html
func connectcases_DeleteRelatedItem(cfg aws.Config, client *connectcases.Client) {
	input := &connectcases.DeleteRelatedItemInput{
		// CaseId: *string, // Required
		// DomainId: *string, // Required
		// RelatedItemId: *string, // Required
	}

	if len(_connectcasesCaseId) > 0 {
		input.CaseId = aws.String(_connectcasesCaseId)
	}
	if len(_connectcasesDomainId) > 0 {
		input.DomainId = aws.String(_connectcasesDomainId)
	}
	if len(_connectcasesRelatedItemId) > 0 {
		input.RelatedItemId = aws.String(_connectcasesRelatedItemId)
	}

	if resp, err := client.DeleteRelatedItem(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a cases template. You can delete up to 100 templates per domain.
// After a cases template is deleted:
//
// - You can still retrieve the template by calling GetTemplate .
//
// - You cannot update the template.
//
// - You cannot create a case by using the deleted template.
//
// - Deleted templates are not included in the ListTemplates response.
func connectcases_DeleteTemplate(cfg aws.Config, client *connectcases.Client) {
	input := &connectcases.DeleteTemplateInput{
		// DomainId: *string, // Required
		// TemplateId: *string, // Required
	}

	if len(_connectcasesDomainId) > 0 {
		input.DomainId = aws.String(_connectcasesDomainId)
	}
	if len(_connectcasesTemplateId) > 0 {
		input.TemplateId = aws.String(_connectcasesTemplateId)
	}

	if resp, err := client.DeleteTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a specific case if it exists.
func connectcases_GetCase(cfg aws.Config, client *connectcases.Client) {
	input := &connectcases.GetCaseInput{
		// CaseId: *string, // Required
		// DomainId: *string, // Required
		// Fields: []types.FieldIdentifier, // Required
	}

	if len(_connectcasesCaseId) > 0 {
		input.CaseId = aws.String(_connectcasesCaseId)
	}
	if len(_connectcasesDomainId) > 0 {
		input.DomainId = aws.String(_connectcasesDomainId)
	}
	if len(_connectcasesFields) > 0 {
		if err := assignInputField(input, "Fields", _connectcasesFields); err != nil {
			log.Errorf("invalid --fields: %s", err.Error())
			return
		}
	}
	if len(_connectcasesNextToken) > 0 {
		input.NextToken = aws.String(_connectcasesNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetCase(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connectcases.GetCaseOutput
	p := connectcases.NewGetCasePaginator(client, input)
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

// Returns the audit history about a specific case if it exists.
func connectcases_GetCaseAuditEvents(cfg aws.Config, client *connectcases.Client) {
	input := &connectcases.GetCaseAuditEventsInput{
		// CaseId: *string, // Required
		// DomainId: *string, // Required
	}

	if len(_connectcasesCaseId) > 0 {
		input.CaseId = aws.String(_connectcasesCaseId)
	}
	if len(_connectcasesDomainId) > 0 {
		input.DomainId = aws.String(_connectcasesDomainId)
	}
	if len(_connectcasesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectcasesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectcasesNextToken) > 0 {
		input.NextToken = aws.String(_connectcasesNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetCaseAuditEvents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connectcases.GetCaseAuditEventsOutput
	p := connectcases.NewGetCaseAuditEventsPaginator(client, input)
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

// Returns the case event publishing configuration.
func connectcases_GetCaseEventConfiguration(cfg aws.Config, client *connectcases.Client) {
	input := &connectcases.GetCaseEventConfigurationInput{
		// DomainId: *string, // Required
	}

	if len(_connectcasesDomainId) > 0 {
		input.DomainId = aws.String(_connectcasesDomainId)
	}

	if resp, err := client.GetCaseEventConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a specific domain if it exists.
func connectcases_GetDomain(cfg aws.Config, client *connectcases.Client) {
	input := &connectcases.GetDomainInput{
		// DomainId: *string, // Required
	}

	if len(_connectcasesDomainId) > 0 {
		input.DomainId = aws.String(_connectcasesDomainId)
	}

	if resp, err := client.GetDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the details for the requested layout.
func connectcases_GetLayout(cfg aws.Config, client *connectcases.Client) {
	input := &connectcases.GetLayoutInput{
		// DomainId: *string, // Required
		// LayoutId: *string, // Required
	}

	if len(_connectcasesDomainId) > 0 {
		input.DomainId = aws.String(_connectcasesDomainId)
	}
	if len(_connectcasesLayoutId) > 0 {
		input.LayoutId = aws.String(_connectcasesLayoutId)
	}

	if resp, err := client.GetLayout(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the details for the requested template. Other template APIs are:
// [CreateTemplate]
//
// [DeleteTemplate]
//
// [ListTemplates]
//
// [UpdateTemplate]
//
// [DeleteTemplate]: https://docs.aws.amazon.com/connect/latest/APIReference/API_connect-cases_DeleteTemplate.html
// [CreateTemplate]: https://docs.aws.amazon.com/connect/latest/APIReference/API_connect-cases_CreateTemplate.html
// [ListTemplates]: https://docs.aws.amazon.com/connect/latest/APIReference/API_connect-cases_ListTemplates.html
// [UpdateTemplate]: https://docs.aws.amazon.com/connect/latest/APIReference/API_connect-cases_UpdateTemplate.html
func connectcases_GetTemplate(cfg aws.Config, client *connectcases.Client) {
	input := &connectcases.GetTemplateInput{
		// DomainId: *string, // Required
		// TemplateId: *string, // Required
	}

	if len(_connectcasesDomainId) > 0 {
		input.DomainId = aws.String(_connectcasesDomainId)
	}
	if len(_connectcasesTemplateId) > 0 {
		input.TemplateId = aws.String(_connectcasesTemplateId)
	}

	if resp, err := client.GetTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all case rules in a Cases domain. In the Amazon Connect admin website,
// case rules are known as case field conditions. For more information about case
// field conditions, see [Add case field conditions to a case template].
//
// [Add case field conditions to a case template]: https://docs.aws.amazon.com/connect/latest/adminguide/case-field-conditions.html
func connectcases_ListCaseRules(cfg aws.Config, client *connectcases.Client) {
	input := &connectcases.ListCaseRulesInput{
		// DomainId: *string, // Required
	}

	if len(_connectcasesDomainId) > 0 {
		input.DomainId = aws.String(_connectcasesDomainId)
	}
	if len(_connectcasesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectcasesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectcasesNextToken) > 0 {
		input.NextToken = aws.String(_connectcasesNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCaseRules(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connectcases.ListCaseRulesOutput
	p := connectcases.NewListCaseRulesPaginator(client, input)
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

// Lists cases for a given contact.
func connectcases_ListCasesForContact(cfg aws.Config, client *connectcases.Client) {
	input := &connectcases.ListCasesForContactInput{
		// ContactArn: *string, // Required
		// DomainId: *string, // Required
	}

	if len(_connectcasesContactArn) > 0 {
		input.ContactArn = aws.String(_connectcasesContactArn)
	}
	if len(_connectcasesDomainId) > 0 {
		input.DomainId = aws.String(_connectcasesDomainId)
	}
	if len(_connectcasesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectcasesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectcasesNextToken) > 0 {
		input.NextToken = aws.String(_connectcasesNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCasesForContact(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connectcases.ListCasesForContactOutput
	p := connectcases.NewListCasesForContactPaginator(client, input)
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

// Lists all cases domains in the Amazon Web Services account. Each list item is a
// condensed summary object of the domain.
func connectcases_ListDomains(cfg aws.Config, client *connectcases.Client) {
	input := &connectcases.ListDomainsInput{}

	if len(_connectcasesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectcasesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectcasesNextToken) > 0 {
		input.NextToken = aws.String(_connectcasesNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDomains(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connectcases.ListDomainsOutput
	p := connectcases.NewListDomainsPaginator(client, input)
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

// Lists all of the field options for a field identifier in the domain.
func connectcases_ListFieldOptions(cfg aws.Config, client *connectcases.Client) {
	input := &connectcases.ListFieldOptionsInput{
		// DomainId: *string, // Required
		// FieldId: *string, // Required
	}

	if len(_connectcasesDomainId) > 0 {
		input.DomainId = aws.String(_connectcasesDomainId)
	}
	if len(_connectcasesFieldId) > 0 {
		input.FieldId = aws.String(_connectcasesFieldId)
	}
	if len(_connectcasesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectcasesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectcasesNextToken) > 0 {
		input.NextToken = aws.String(_connectcasesNextToken)
	}
	if len(_connectcasesValues) > 0 {
		input.Values = append([]string(nil), _connectcasesValues...)
	}

	if disablePaginator() {
		if resp, err := client.ListFieldOptions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connectcases.ListFieldOptionsOutput
	p := connectcases.NewListFieldOptionsPaginator(client, input)
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

// Lists all fields in a Cases domain.
func connectcases_ListFields(cfg aws.Config, client *connectcases.Client) {
	input := &connectcases.ListFieldsInput{
		// DomainId: *string, // Required
	}

	if len(_connectcasesDomainId) > 0 {
		input.DomainId = aws.String(_connectcasesDomainId)
	}
	if len(_connectcasesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectcasesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectcasesNextToken) > 0 {
		input.NextToken = aws.String(_connectcasesNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListFields(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connectcases.ListFieldsOutput
	p := connectcases.NewListFieldsPaginator(client, input)
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

// Lists all layouts in the given cases domain. Each list item is a condensed
// summary object of the layout.
func connectcases_ListLayouts(cfg aws.Config, client *connectcases.Client) {
	input := &connectcases.ListLayoutsInput{
		// DomainId: *string, // Required
	}

	if len(_connectcasesDomainId) > 0 {
		input.DomainId = aws.String(_connectcasesDomainId)
	}
	if len(_connectcasesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectcasesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectcasesNextToken) > 0 {
		input.NextToken = aws.String(_connectcasesNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListLayouts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connectcases.ListLayoutsOutput
	p := connectcases.NewListLayoutsPaginator(client, input)
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

// Lists tags for a resource.
func connectcases_ListTagsForResource(cfg aws.Config, client *connectcases.Client) {
	input := &connectcases.ListTagsForResourceInput{
		// Arn: *string, // Required
	}

	if len(_connectcasesArn) > 0 {
		input.Arn = aws.String(_connectcasesArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all of the templates in a Cases domain. Each list item is a condensed
// summary object of the template.
//
// Other template APIs are:
//
// [CreateTemplate]
//
// [DeleteTemplate]
//
// [GetTemplate]
//
// [UpdateTemplate]
//
// [DeleteTemplate]: https://docs.aws.amazon.com/connect/latest/APIReference/API_connect-cases_DeleteTemplate.html
// [CreateTemplate]: https://docs.aws.amazon.com/connect/latest/APIReference/API_connect-cases_CreateTemplate.html
// [UpdateTemplate]: https://docs.aws.amazon.com/connect/latest/APIReference/API_connect-cases_UpdateTemplate.html
// [GetTemplate]: https://docs.aws.amazon.com/connect/latest/APIReference/API_connect-cases_GetTemplate.html
func connectcases_ListTemplates(cfg aws.Config, client *connectcases.Client) {
	input := &connectcases.ListTemplatesInput{
		// DomainId: *string, // Required
	}

	if len(_connectcasesDomainId) > 0 {
		input.DomainId = aws.String(_connectcasesDomainId)
	}
	if len(_connectcasesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectcasesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectcasesNextToken) > 0 {
		input.NextToken = aws.String(_connectcasesNextToken)
	}
	if len(_connectcasesStatus) > 0 {
		if err := assignInputField(input, "Status", _connectcasesStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListTemplates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connectcases.ListTemplatesOutput
	p := connectcases.NewListTemplatesPaginator(client, input)
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

// Adds case event publishing configuration. For a complete list of fields you can
// add to the event message, see [Create case fields]in the Amazon Connect Administrator Guide
//
// [Create case fields]: https://docs.aws.amazon.com/connect/latest/adminguide/case-fields.html
func connectcases_PutCaseEventConfiguration(cfg aws.Config, client *connectcases.Client) {
	input := &connectcases.PutCaseEventConfigurationInput{
		// DomainId: *string, // Required
		// EventBridge: *types.EventBridgeConfiguration, // Required
	}

	if len(_connectcasesDomainId) > 0 {
		input.DomainId = aws.String(_connectcasesDomainId)
	}
	if len(_connectcasesEventBridge) > 0 {
		if err := assignInputField(input, "EventBridge", _connectcasesEventBridge); err != nil {
			log.Errorf("invalid --event-bridge: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutCaseEventConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Searches for related items across all cases within a domain. This is a global
// search operation that returns related items from multiple cases, unlike the
// case-specific [SearchRelatedItems]API.
//
// # Use cases
//
// Following are common uses cases for this API:
//
// - Find cases with similar issues across the domain. For example, search for
// all cases containing comments about "product defect" to identify patterns and
// existing solutions.
//
// - Locate all cases associated with specific contacts or orders. For example,
// find all cases linked to a contactArn to understand the complete customer
// journey.
//
// - Monitor SLA compliance across cases. For example, search for all cases with
// "Active" SLA status to prioritize remediation efforts.
//
// # Important things to know
//
// - This API returns case identifiers, not complete case objects. To retrieve
// full case details, you must make additional calls to the [GetCase]API for each
// returned case ID.
//
// - This API searches across related items content, not case fields. Use the [SearchCases]
// API to search within case field values.
//
// Endpoints: See [Amazon Connect endpoints and quotas].
//
// [GetCase]: https://docs.aws.amazon.com/connect/latest/APIReference/API_connect-cases_GetCase.html
// [SearchCases]: https://docs.aws.amazon.com/connect/latest/APIReference/API_connect-cases_SearchCases.html
// [Amazon Connect endpoints and quotas]: https://docs.aws.amazon.com/general/latest/gr/connect_region.html
// [SearchRelatedItems]: https://docs.aws.amazon.com/connect/latest/APIReference/API_connect-cases_SearchRelatedItems.html
func connectcases_SearchAllRelatedItems(cfg aws.Config, client *connectcases.Client) {
	input := &connectcases.SearchAllRelatedItemsInput{
		// DomainId: *string, // Required
	}

	if len(_connectcasesDomainId) > 0 {
		input.DomainId = aws.String(_connectcasesDomainId)
	}
	if len(_connectcasesFilters) > 0 {
		if err := assignInputField(input, "Filters", _connectcasesFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_connectcasesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectcasesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectcasesNextToken) > 0 {
		input.NextToken = aws.String(_connectcasesNextToken)
	}
	if len(_connectcasesSorts) > 0 {
		if err := assignInputField(input, "Sorts", _connectcasesSorts); err != nil {
			log.Errorf("invalid --sorts: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.SearchAllRelatedItems(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connectcases.SearchAllRelatedItemsOutput
	p := connectcases.NewSearchAllRelatedItemsPaginator(client, input)
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

// Searches for cases within their associated Cases domain. Search results are
// returned as a paginated list of abridged case documents.
//
// For customer_id you must provide the full customer profile ARN in this format:
// arn:aws:profile:your AWS Region:your AWS account ID:domains/profiles domain
// name/profiles/profile ID .
func connectcases_SearchCases(cfg aws.Config, client *connectcases.Client) {
	input := &connectcases.SearchCasesInput{
		// DomainId: *string, // Required
	}

	if len(_connectcasesDomainId) > 0 {
		input.DomainId = aws.String(_connectcasesDomainId)
	}
	if len(_connectcasesFields) > 0 {
		if err := assignInputField(input, "Fields", _connectcasesFields); err != nil {
			log.Errorf("invalid --fields: %s", err.Error())
			return
		}
	}
	if len(_connectcasesFilter) > 0 {
		if err := assignInputField(input, "Filter", _connectcasesFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_connectcasesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectcasesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectcasesNextToken) > 0 {
		input.NextToken = aws.String(_connectcasesNextToken)
	}
	if len(_connectcasesSearchTerm) > 0 {
		input.SearchTerm = aws.String(_connectcasesSearchTerm)
	}
	if len(_connectcasesSorts) > 0 {
		if err := assignInputField(input, "Sorts", _connectcasesSorts); err != nil {
			log.Errorf("invalid --sorts: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.SearchCases(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connectcases.SearchCasesOutput
	p := connectcases.NewSearchCasesPaginator(client, input)
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

// Searches for related items that are associated with a case.
// If no filters are provided, this returns all related items associated with a
// case.
func connectcases_SearchRelatedItems(cfg aws.Config, client *connectcases.Client) {
	input := &connectcases.SearchRelatedItemsInput{
		// CaseId: *string, // Required
		// DomainId: *string, // Required
	}

	if len(_connectcasesCaseId) > 0 {
		input.CaseId = aws.String(_connectcasesCaseId)
	}
	if len(_connectcasesDomainId) > 0 {
		input.DomainId = aws.String(_connectcasesDomainId)
	}
	if len(_connectcasesFilters) > 0 {
		if err := assignInputField(input, "Filters", _connectcasesFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_connectcasesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectcasesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectcasesNextToken) > 0 {
		input.NextToken = aws.String(_connectcasesNextToken)
	}

	if disablePaginator() {
		if resp, err := client.SearchRelatedItems(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connectcases.SearchRelatedItemsOutput
	p := connectcases.NewSearchRelatedItemsPaginator(client, input)
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

// Adds tags to a resource.
func connectcases_TagResource(cfg aws.Config, client *connectcases.Client) {
	input := &connectcases.TagResourceInput{
		// Arn: *string, // Required
		// Tags: map[string]*string, // Required
	}

	if len(_connectcasesArn) > 0 {
		input.Arn = aws.String(_connectcasesArn)
	}
	if len(_connectcasesTags) > 0 {
		if err := assignInputField(input, "Tags", _connectcasesTags); err != nil {
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

// Untags a resource.
func connectcases_UntagResource(cfg aws.Config, client *connectcases.Client) {
	input := &connectcases.UntagResourceInput{
		// Arn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_connectcasesArn) > 0 {
		input.Arn = aws.String(_connectcasesArn)
	}
	if len(_connectcasesTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _connectcasesTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// If you provide a value for PerformedBy.UserArn you must also have [connect:DescribeUser] permission
// on the User ARN resource that you provide
//
// Updates the values of fields on a case. Fields to be updated are received as an
// array of id/value pairs identical to the CreateCase input .
//
// If the action is successful, the service sends back an HTTP 200 response with
// an empty HTTP body.
//
// [connect:DescribeUser]: https://docs.aws.amazon.com/connect/latest/APIReference/API_DescribeUser.html
func connectcases_UpdateCase(cfg aws.Config, client *connectcases.Client) {
	input := &connectcases.UpdateCaseInput{
		// CaseId: *string, // Required
		// DomainId: *string, // Required
		// Fields: []types.FieldValue, // Required
	}

	if len(_connectcasesCaseId) > 0 {
		input.CaseId = aws.String(_connectcasesCaseId)
	}
	if len(_connectcasesDomainId) > 0 {
		input.DomainId = aws.String(_connectcasesDomainId)
	}
	if len(_connectcasesFields) > 0 {
		if err := assignInputField(input, "Fields", _connectcasesFields); err != nil {
			log.Errorf("invalid --fields: %s", err.Error())
			return
		}
	}
	if len(_connectcasesPerformedBy) > 0 {
		if err := assignInputField(input, "PerformedBy", _connectcasesPerformedBy); err != nil {
			log.Errorf("invalid --performed-by: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateCase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a case rule. In the Amazon Connect admin website, case rules are known
// as case field conditions. For more information about case field conditions, see [Add case field conditions to a case template]
// .
//
// [Add case field conditions to a case template]: https://docs.aws.amazon.com/connect/latest/adminguide/case-field-conditions.html
func connectcases_UpdateCaseRule(cfg aws.Config, client *connectcases.Client) {
	input := &connectcases.UpdateCaseRuleInput{
		// CaseRuleId: *string, // Required
		// DomainId: *string, // Required
	}

	if len(_connectcasesCaseRuleId) > 0 {
		input.CaseRuleId = aws.String(_connectcasesCaseRuleId)
	}
	if len(_connectcasesDomainId) > 0 {
		input.DomainId = aws.String(_connectcasesDomainId)
	}
	if len(_connectcasesDescription) > 0 {
		input.Description = aws.String(_connectcasesDescription)
	}
	if len(_connectcasesName) > 0 {
		input.Name = aws.String(_connectcasesName)
	}
	if len(_connectcasesRule) > 0 {
		if err := assignInputField(input, "Rule", _connectcasesRule); err != nil {
			log.Errorf("invalid --rule: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateCaseRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the properties of an existing field.
func connectcases_UpdateField(cfg aws.Config, client *connectcases.Client) {
	input := &connectcases.UpdateFieldInput{
		// DomainId: *string, // Required
		// FieldId: *string, // Required
	}

	if len(_connectcasesDomainId) > 0 {
		input.DomainId = aws.String(_connectcasesDomainId)
	}
	if len(_connectcasesFieldId) > 0 {
		input.FieldId = aws.String(_connectcasesFieldId)
	}
	if len(_connectcasesAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _connectcasesAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_connectcasesDescription) > 0 {
		input.Description = aws.String(_connectcasesDescription)
	}
	if len(_connectcasesName) > 0 {
		input.Name = aws.String(_connectcasesName)
	}

	if resp, err := client.UpdateField(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the attributes of an existing layout.
// If the action is successful, the service sends back an HTTP 200 response with
// an empty HTTP body.
//
// A ValidationException is returned when you add non-existent fieldIds to a
// layout.
//
// Title and Status fields cannot be part of layouts because they are not
// configurable.
func connectcases_UpdateLayout(cfg aws.Config, client *connectcases.Client) {
	input := &connectcases.UpdateLayoutInput{
		// DomainId: *string, // Required
		// LayoutId: *string, // Required
	}

	if len(_connectcasesDomainId) > 0 {
		input.DomainId = aws.String(_connectcasesDomainId)
	}
	if len(_connectcasesLayoutId) > 0 {
		input.LayoutId = aws.String(_connectcasesLayoutId)
	}
	if len(_connectcasesContent) > 0 {
		if err := assignInputField(input, "Content", _connectcasesContent); err != nil {
			log.Errorf("invalid --content: %s", err.Error())
			return
		}
	}
	if len(_connectcasesName) > 0 {
		input.Name = aws.String(_connectcasesName)
	}

	if resp, err := client.UpdateLayout(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the attributes of an existing template. The template attributes that
// can be modified include name , description , layoutConfiguration ,
// requiredFields , and status . At least one of these attributes must not be null.
// If a null value is provided for a given attribute, that attribute is ignored and
// its current value is preserved.
//
// Other template APIs are:
//
// [CreateTemplate]
//
// [DeleteTemplate]
//
// [GetTemplate]
//
// [ListTemplates]
//
// [DeleteTemplate]: https://docs.aws.amazon.com/connect/latest/APIReference/API_connect-cases_DeleteTemplate.html
// [CreateTemplate]: https://docs.aws.amazon.com/connect/latest/APIReference/API_connect-cases_CreateTemplate.html
// [ListTemplates]: https://docs.aws.amazon.com/connect/latest/APIReference/API_connect-cases_ListTemplates.html
// [GetTemplate]: https://docs.aws.amazon.com/connect/latest/APIReference/API_connect-cases_GetTemplate.html
func connectcases_UpdateTemplate(cfg aws.Config, client *connectcases.Client) {
	input := &connectcases.UpdateTemplateInput{
		// DomainId: *string, // Required
		// TemplateId: *string, // Required
	}

	if len(_connectcasesDomainId) > 0 {
		input.DomainId = aws.String(_connectcasesDomainId)
	}
	if len(_connectcasesTemplateId) > 0 {
		input.TemplateId = aws.String(_connectcasesTemplateId)
	}
	if len(_connectcasesDescription) > 0 {
		input.Description = aws.String(_connectcasesDescription)
	}
	if len(_connectcasesLayoutConfiguration) > 0 {
		if err := assignInputField(input, "LayoutConfiguration", _connectcasesLayoutConfiguration); err != nil {
			log.Errorf("invalid --layout-configuration: %s", err.Error())
			return
		}
	}
	if len(_connectcasesName) > 0 {
		input.Name = aws.String(_connectcasesName)
	}
	if len(_connectcasesRequiredFields) > 0 {
		if err := assignInputField(input, "RequiredFields", _connectcasesRequiredFields); err != nil {
			log.Errorf("invalid --required-fields: %s", err.Error())
			return
		}
	}
	if len(_connectcasesRules) > 0 {
		if err := assignInputField(input, "Rules", _connectcasesRules); err != nil {
			log.Errorf("invalid --rules: %s", err.Error())
			return
		}
	}
	if len(_connectcasesStatus) > 0 {
		if err := assignInputField(input, "Status", _connectcasesStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_connectcasesTagPropagationConfigurations) > 0 {
		if err := assignInputField(input, "TagPropagationConfigurations", _connectcasesTagPropagationConfigurations); err != nil {
			log.Errorf("invalid --tag-propagation-configurations: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_connectcasesCmd)
	_connectcasesCmd.Flags().SortFlags = false

	_connectcasesCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_connectcasesCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_connectcasesCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_connectcasesCmd.Flags().StringVarP(&_connectcasesArn, "arn", "", "", "ARN")
	_connectcasesCmd.Flags().StringVarP(&_connectcasesAttributes, "attributes", "", "", "Attributes")
	_connectcasesCmd.Flags().StringVarP(&_connectcasesCaseId, "case-id", "", "", "Case ID")
	_connectcasesCmd.Flags().StringVarP(&_connectcasesCaseRuleId, "case-rule-id", "", "", "Case Rule ID")
	_connectcasesCmd.Flags().StringVarP(&_connectcasesCaseRules, "case-rules", "", "", "Case Rules")
	_connectcasesCmd.Flags().StringVarP(&_connectcasesClientToken, "client-token", "", "", "Client Token")
	_connectcasesCmd.Flags().StringVarP(&_connectcasesContactArn, "contact-arn", "", "", "Contact ARN")
	_connectcasesCmd.Flags().StringVarP(&_connectcasesContent, "content", "", "", "Content")
	_connectcasesCmd.Flags().StringVarP(&_connectcasesDescription, "description", "", "", "Description")
	_connectcasesCmd.Flags().StringVarP(&_connectcasesDomainId, "domain-id", "", "", "Domain ID")
	_connectcasesCmd.Flags().StringVarP(&_connectcasesEventBridge, "event-bridge", "", "", "Event Bridge")
	_connectcasesCmd.Flags().StringVarP(&_connectcasesFieldId, "field-id", "", "", "Field ID")
	_connectcasesCmd.Flags().StringVarP(&_connectcasesFields, "fields", "", "", "Fields")
	_connectcasesCmd.Flags().StringVarP(&_connectcasesFilter, "filter", "", "", "Filter")
	_connectcasesCmd.Flags().StringVarP(&_connectcasesFilters, "filters", "", "", "Filters")
	_connectcasesCmd.Flags().StringVarP(&_connectcasesLayoutConfiguration, "layout-configuration", "", "", "Layout Configuration")
	_connectcasesCmd.Flags().StringVarP(&_connectcasesLayoutId, "layout-id", "", "", "Layout ID")
	_connectcasesCmd.Flags().StringVarP(&_connectcasesMaxResults, "max-results", "", "", "Max Results")
	_connectcasesCmd.Flags().StringVarP(&_connectcasesName, "name", "", "", "Name")
	_connectcasesCmd.Flags().StringVarP(&_connectcasesNextToken, "next-token", "", "", "Next Token")
	_connectcasesCmd.Flags().StringVarP(&_connectcasesOptions, "options", "", "", "Options")
	_connectcasesCmd.Flags().StringVarP(&_connectcasesPerformedBy, "performed-by", "", "", "Performed By")
	_connectcasesCmd.Flags().StringVarP(&_connectcasesRelatedItemId, "related-item-id", "", "", "Related Item ID")
	_connectcasesCmd.Flags().StringVarP(&_connectcasesRequiredFields, "required-fields", "", "", "Required Fields")
	_connectcasesCmd.Flags().StringVarP(&_connectcasesRule, "rule", "", "", "Rule")
	_connectcasesCmd.Flags().StringVarP(&_connectcasesRules, "rules", "", "", "Rules")
	_connectcasesCmd.Flags().StringVarP(&_connectcasesSearchTerm, "search-term", "", "", "Search Term")
	_connectcasesCmd.Flags().StringVarP(&_connectcasesSorts, "sorts", "", "", "Sorts")
	_connectcasesCmd.Flags().StringVarP(&_connectcasesStatus, "status", "", "", "Status")
	_connectcasesCmd.Flags().StringSliceVarP(&_connectcasesTagKeys, "tag-keys", "", nil, "Tag Keys")
	_connectcasesCmd.Flags().StringVarP(&_connectcasesTagPropagationConfigurations, "tag-propagation-configurations", "", "", "Tag Propagation Configurations")
	_connectcasesCmd.Flags().StringVarP(&_connectcasesTags, "tags", "", "", "Tags")
	_connectcasesCmd.Flags().StringVarP(&_connectcasesTemplateId, "template-id", "", "", "Template ID")
	_connectcasesCmd.Flags().StringVarP(&_connectcasesType, "type", "", "", "Type")
	_connectcasesCmd.Flags().StringSliceVarP(&_connectcasesValues, "values", "", nil, "Values")

	_connectcasesCmd.Flags().BoolVarP(&_connectcasesBatchGetCaseRule, "batch-get-case-rule", "", false, "Batch Get Case Rule")
	_connectcasesCmd.Flags().BoolVarP(&_connectcasesBatchGetField, "batch-get-field", "", false, "Batch Get Field")
	_connectcasesCmd.Flags().BoolVarP(&_connectcasesBatchPutFieldOptions, "batch-put-field-options", "", false, "Batch Put Field Options")
	_connectcasesCmd.Flags().BoolVarP(&_connectcasesCreateCase, "create-case", "", false, "Create Case")
	_connectcasesCmd.Flags().BoolVarP(&_connectcasesCreateCaseRule, "create-case-rule", "", false, "Create Case Rule")
	_connectcasesCmd.Flags().BoolVarP(&_connectcasesCreateDomain, "create-domain", "", false, "Create Domain")
	_connectcasesCmd.Flags().BoolVarP(&_connectcasesCreateField, "create-field", "", false, "Create Field")
	_connectcasesCmd.Flags().BoolVarP(&_connectcasesCreateLayout, "create-layout", "", false, "Create Layout")
	_connectcasesCmd.Flags().BoolVarP(&_connectcasesCreateRelatedItem, "create-related-item", "", false, "Create Related Item")
	_connectcasesCmd.Flags().BoolVarP(&_connectcasesCreateTemplate, "create-template", "", false, "Create Template")
	_connectcasesCmd.Flags().BoolVarP(&_connectcasesDeleteCase, "delete-case", "", false, "Delete Case")
	_connectcasesCmd.Flags().BoolVarP(&_connectcasesDeleteCaseRule, "delete-case-rule", "", false, "Delete Case Rule")
	_connectcasesCmd.Flags().BoolVarP(&_connectcasesDeleteDomain, "delete-domain", "", false, "Delete Domain")
	_connectcasesCmd.Flags().BoolVarP(&_connectcasesDeleteField, "delete-field", "", false, "Delete Field")
	_connectcasesCmd.Flags().BoolVarP(&_connectcasesDeleteLayout, "delete-layout", "", false, "Delete Layout")
	_connectcasesCmd.Flags().BoolVarP(&_connectcasesDeleteRelatedItem, "delete-related-item", "", false, "Delete Related Item")
	_connectcasesCmd.Flags().BoolVarP(&_connectcasesDeleteTemplate, "delete-template", "", false, "Delete Template")
	_connectcasesCmd.Flags().BoolVarP(&_connectcasesGetCase, "get-case", "", false, "Get Case")
	_connectcasesCmd.Flags().BoolVarP(&_connectcasesGetCaseAuditEvents, "get-case-audit-events", "", false, "Get Case Audit Events")
	_connectcasesCmd.Flags().BoolVarP(&_connectcasesGetCaseEventConfiguration, "get-case-event-configuration", "", false, "Get Case Event Configuration")
	_connectcasesCmd.Flags().BoolVarP(&_connectcasesGetDomain, "get-domain", "", false, "Get Domain")
	_connectcasesCmd.Flags().BoolVarP(&_connectcasesGetLayout, "get-layout", "", false, "Get Layout")
	_connectcasesCmd.Flags().BoolVarP(&_connectcasesGetTemplate, "get-template", "", false, "Get Template")
	_connectcasesCmd.Flags().BoolVarP(&_connectcasesListCaseRules, "list-case-rules", "", false, "List Case Rules")
	_connectcasesCmd.Flags().BoolVarP(&_connectcasesListCasesForContact, "list-cases-for-contact", "", false, "List Cases For Contact")
	_connectcasesCmd.Flags().BoolVarP(&_connectcasesListDomains, "list-domains", "", false, "List Domains")
	_connectcasesCmd.Flags().BoolVarP(&_connectcasesListFieldOptions, "list-field-options", "", false, "List Field Options")
	_connectcasesCmd.Flags().BoolVarP(&_connectcasesListFields, "list-fields", "", false, "List Fields")
	_connectcasesCmd.Flags().BoolVarP(&_connectcasesListLayouts, "list-layouts", "", false, "List Layouts")
	_connectcasesCmd.Flags().BoolVarP(&_connectcasesListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_connectcasesCmd.Flags().BoolVarP(&_connectcasesListTemplates, "list-templates", "", false, "List Templates")
	_connectcasesCmd.Flags().BoolVarP(&_connectcasesPutCaseEventConfiguration, "put-case-event-configuration", "", false, "Put Case Event Configuration")
	_connectcasesCmd.Flags().BoolVarP(&_connectcasesSearchAllRelatedItems, "search-all-related-items", "", false, "Search All Related Items")
	_connectcasesCmd.Flags().BoolVarP(&_connectcasesSearchCases, "search-cases", "", false, "Search Cases")
	_connectcasesCmd.Flags().BoolVarP(&_connectcasesSearchRelatedItems, "search-related-items", "", false, "Search Related Items")
	_connectcasesCmd.Flags().BoolVarP(&_connectcasesTagResource, "tag-resource", "", false, "Tag Resource")
	_connectcasesCmd.Flags().BoolVarP(&_connectcasesUntagResource, "untag-resource", "", false, "Untag Resource")
	_connectcasesCmd.Flags().BoolVarP(&_connectcasesUpdateCase, "update-case", "", false, "Update Case")
	_connectcasesCmd.Flags().BoolVarP(&_connectcasesUpdateCaseRule, "update-case-rule", "", false, "Update Case Rule")
	_connectcasesCmd.Flags().BoolVarP(&_connectcasesUpdateField, "update-field", "", false, "Update Field")
	_connectcasesCmd.Flags().BoolVarP(&_connectcasesUpdateLayout, "update-layout", "", false, "Update Layout")
	_connectcasesCmd.Flags().BoolVarP(&_connectcasesUpdateTemplate, "update-template", "", false, "Update Template")

}
