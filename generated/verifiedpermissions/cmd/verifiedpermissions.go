package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/verifiedpermissions"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// verifiedpermissionsCmd represents the verifiedpermissions command
var _verifiedpermissionsCmd = &cobra.Command{
	Use:   "verifiedpermissions",
	Short: "AWS verifiedpermissions CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := verifiedpermissions.NewFromConfig(cfg)
		if _verifiedpermissionsBatchGetPolicy {
			verifiedpermissions_BatchGetPolicy(cfg, client)
			return
		}
		if _verifiedpermissionsBatchIsAuthorized {
			verifiedpermissions_BatchIsAuthorized(cfg, client)
			return
		}
		if _verifiedpermissionsBatchIsAuthorizedWithToken {
			verifiedpermissions_BatchIsAuthorizedWithToken(cfg, client)
			return
		}
		if _verifiedpermissionsCreateIdentitySource {
			verifiedpermissions_CreateIdentitySource(cfg, client)
			return
		}
		if _verifiedpermissionsCreatePolicy {
			verifiedpermissions_CreatePolicy(cfg, client)
			return
		}
		if _verifiedpermissionsCreatePolicyStore {
			verifiedpermissions_CreatePolicyStore(cfg, client)
			return
		}
		if _verifiedpermissionsCreatePolicyTemplate {
			verifiedpermissions_CreatePolicyTemplate(cfg, client)
			return
		}
		if _verifiedpermissionsDeleteIdentitySource {
			verifiedpermissions_DeleteIdentitySource(cfg, client)
			return
		}
		if _verifiedpermissionsDeletePolicy {
			verifiedpermissions_DeletePolicy(cfg, client)
			return
		}
		if _verifiedpermissionsDeletePolicyStore {
			verifiedpermissions_DeletePolicyStore(cfg, client)
			return
		}
		if _verifiedpermissionsDeletePolicyTemplate {
			verifiedpermissions_DeletePolicyTemplate(cfg, client)
			return
		}
		if _verifiedpermissionsGetIdentitySource {
			verifiedpermissions_GetIdentitySource(cfg, client)
			return
		}
		if _verifiedpermissionsGetPolicy {
			verifiedpermissions_GetPolicy(cfg, client)
			return
		}
		if _verifiedpermissionsGetPolicyStore {
			verifiedpermissions_GetPolicyStore(cfg, client)
			return
		}
		if _verifiedpermissionsGetPolicyTemplate {
			verifiedpermissions_GetPolicyTemplate(cfg, client)
			return
		}
		if _verifiedpermissionsGetSchema {
			verifiedpermissions_GetSchema(cfg, client)
			return
		}
		if _verifiedpermissionsIsAuthorized {
			verifiedpermissions_IsAuthorized(cfg, client)
			return
		}
		if _verifiedpermissionsIsAuthorizedWithToken {
			verifiedpermissions_IsAuthorizedWithToken(cfg, client)
			return
		}
		if _verifiedpermissionsListIdentitySources {
			verifiedpermissions_ListIdentitySources(cfg, client)
			return
		}
		if _verifiedpermissionsListPolicies {
			verifiedpermissions_ListPolicies(cfg, client)
			return
		}
		if _verifiedpermissionsListPolicyStores {
			verifiedpermissions_ListPolicyStores(cfg, client)
			return
		}
		if _verifiedpermissionsListPolicyTemplates {
			verifiedpermissions_ListPolicyTemplates(cfg, client)
			return
		}
		if _verifiedpermissionsListTagsForResource {
			verifiedpermissions_ListTagsForResource(cfg, client)
			return
		}
		if _verifiedpermissionsPutSchema {
			verifiedpermissions_PutSchema(cfg, client)
			return
		}
		if _verifiedpermissionsTagResource {
			verifiedpermissions_TagResource(cfg, client)
			return
		}
		if _verifiedpermissionsUntagResource {
			verifiedpermissions_UntagResource(cfg, client)
			return
		}
		if _verifiedpermissionsUpdateIdentitySource {
			verifiedpermissions_UpdateIdentitySource(cfg, client)
			return
		}
		if _verifiedpermissionsUpdatePolicy {
			verifiedpermissions_UpdatePolicy(cfg, client)
			return
		}
		if _verifiedpermissionsUpdatePolicyStore {
			verifiedpermissions_UpdatePolicyStore(cfg, client)
			return
		}
		if _verifiedpermissionsUpdatePolicyTemplate {
			verifiedpermissions_UpdatePolicyTemplate(cfg, client)
			return
		}

	},
}

var (
	_verifiedpermissionsBatchGetPolicy             bool
	_verifiedpermissionsBatchIsAuthorized          bool
	_verifiedpermissionsBatchIsAuthorizedWithToken bool
	_verifiedpermissionsCreateIdentitySource       bool
	_verifiedpermissionsCreatePolicy               bool
	_verifiedpermissionsCreatePolicyStore          bool
	_verifiedpermissionsCreatePolicyTemplate       bool
	_verifiedpermissionsDeleteIdentitySource       bool
	_verifiedpermissionsDeletePolicy               bool
	_verifiedpermissionsDeletePolicyStore          bool
	_verifiedpermissionsDeletePolicyTemplate       bool
	_verifiedpermissionsGetIdentitySource          bool
	_verifiedpermissionsGetPolicy                  bool
	_verifiedpermissionsGetPolicyStore             bool
	_verifiedpermissionsGetPolicyTemplate          bool
	_verifiedpermissionsGetSchema                  bool
	_verifiedpermissionsIsAuthorized               bool
	_verifiedpermissionsIsAuthorizedWithToken      bool
	_verifiedpermissionsListIdentitySources        bool
	_verifiedpermissionsListPolicies               bool
	_verifiedpermissionsListPolicyStores           bool
	_verifiedpermissionsListPolicyTemplates        bool
	_verifiedpermissionsListTagsForResource        bool
	_verifiedpermissionsPutSchema                  bool
	_verifiedpermissionsTagResource                bool
	_verifiedpermissionsUntagResource              bool
	_verifiedpermissionsUpdateIdentitySource       bool
	_verifiedpermissionsUpdatePolicy               bool
	_verifiedpermissionsUpdatePolicyStore          bool
	_verifiedpermissionsUpdatePolicyTemplate       bool

	_verifiedpermissionsAccessToken         string
	_verifiedpermissionsAction              string
	_verifiedpermissionsClientToken         string
	_verifiedpermissionsConfiguration       string
	_verifiedpermissionsContext             string
	_verifiedpermissionsDefinition          string
	_verifiedpermissionsDeletionProtection  string
	_verifiedpermissionsDescription         string
	_verifiedpermissionsEncryptionSettings  string
	_verifiedpermissionsEntities            string
	_verifiedpermissionsFilter              string
	_verifiedpermissionsFilters             string
	_verifiedpermissionsIdentitySourceId    string
	_verifiedpermissionsIdentityToken       string
	_verifiedpermissionsMaxResults          string
	_verifiedpermissionsNextToken           string
	_verifiedpermissionsPolicyId            string
	_verifiedpermissionsPolicyStoreId       string
	_verifiedpermissionsPolicyTemplateId    string
	_verifiedpermissionsPrincipal           string
	_verifiedpermissionsPrincipalEntityType string
	_verifiedpermissionsRequests            string
	_verifiedpermissionsResource            string
	_verifiedpermissionsResourceArn         string
	_verifiedpermissionsStatement           string
	_verifiedpermissionsTagKeys             []string
	_verifiedpermissionsTags                string
	_verifiedpermissionsUpdateConfiguration string
	_verifiedpermissionsValidationSettings  string
)

// Retrieves information about a group (batch) of policies.
// The BatchGetPolicy operation doesn't have its own IAM permission. To authorize
// this operation for Amazon Web Services principals, include the permission
// verifiedpermissions:GetPolicy in their IAM policies.
func verifiedpermissions_BatchGetPolicy(cfg aws.Config, client *verifiedpermissions.Client) {
	input := &verifiedpermissions.BatchGetPolicyInput{
		// Requests: []types.BatchGetPolicyInputItem, // Required
	}

	if len(_verifiedpermissionsRequests) > 0 {
		if err := assignInputField(input, "Requests", _verifiedpermissionsRequests); err != nil {
			log.Errorf("invalid --requests: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchGetPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Makes a series of decisions about multiple authorization requests for one
// principal or resource. Each request contains the equivalent content of an
// IsAuthorized request: principal, action, resource, and context. Either the
// principal or the resource parameter must be identical across all requests. For
// example, Verified Permissions won't evaluate a pair of requests where bob views
// photo1 and alice views photo2 . Authorization of bob to view photo1 and photo2 ,
// or bob and alice to view photo1 , are valid batches.
//
// The request is evaluated against all policies in the specified policy store
// that match the entities that you declare. The result of the decisions is a
// series of Allow or Deny responses, along with the IDs of the policies that
// produced each decision.
//
// The entities of a BatchIsAuthorized API request can contain up to 100
// principals and up to 100 resources. The requests of a BatchIsAuthorized API
// request can contain up to 30 requests.
//
// The BatchIsAuthorized operation doesn't have its own IAM permission. To
// authorize this operation for Amazon Web Services principals, include the
// permission verifiedpermissions:IsAuthorized in their IAM policies.
func verifiedpermissions_BatchIsAuthorized(cfg aws.Config, client *verifiedpermissions.Client) {
	input := &verifiedpermissions.BatchIsAuthorizedInput{
		// PolicyStoreId: *string, // Required
		// Requests: []types.BatchIsAuthorizedInputItem, // Required
	}

	if len(_verifiedpermissionsPolicyStoreId) > 0 {
		input.PolicyStoreId = aws.String(_verifiedpermissionsPolicyStoreId)
	}
	if len(_verifiedpermissionsRequests) > 0 {
		if err := assignInputField(input, "Requests", _verifiedpermissionsRequests); err != nil {
			log.Errorf("invalid --requests: %s", err.Error())
			return
		}
	}
	if len(_verifiedpermissionsEntities) > 0 {
		if err := assignInputField(input, "Entities", _verifiedpermissionsEntities); err != nil {
			log.Errorf("invalid --entities: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchIsAuthorized(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Makes a series of decisions about multiple authorization requests for one
// token. The principal in this request comes from an external identity source in
// the form of an identity or access token, formatted as a [JSON web token (JWT)]. The information in
// the parameters can also define additional context that Verified Permissions can
// include in the evaluations.
//
// The request is evaluated against all policies in the specified policy store
// that match the entities that you provide in the entities declaration and in the
// token. The result of the decisions is a series of Allow or Deny responses,
// along with the IDs of the policies that produced each decision.
//
// The entities of a BatchIsAuthorizedWithToken API request can contain up to 100
// resources and up to 99 user groups. The requests of a BatchIsAuthorizedWithToken
// API request can contain up to 30 requests.
//
// The BatchIsAuthorizedWithToken operation doesn't have its own IAM permission.
// To authorize this operation for Amazon Web Services principals, include the
// permission verifiedpermissions:IsAuthorizedWithToken in their IAM policies.
//
// [JSON web token (JWT)]: https://wikipedia.org/wiki/JSON_Web_Token
func verifiedpermissions_BatchIsAuthorizedWithToken(cfg aws.Config, client *verifiedpermissions.Client) {
	input := &verifiedpermissions.BatchIsAuthorizedWithTokenInput{
		// PolicyStoreId: *string, // Required
		// Requests: []types.BatchIsAuthorizedWithTokenInputItem, // Required
	}

	if len(_verifiedpermissionsPolicyStoreId) > 0 {
		input.PolicyStoreId = aws.String(_verifiedpermissionsPolicyStoreId)
	}
	if len(_verifiedpermissionsRequests) > 0 {
		if err := assignInputField(input, "Requests", _verifiedpermissionsRequests); err != nil {
			log.Errorf("invalid --requests: %s", err.Error())
			return
		}
	}
	if len(_verifiedpermissionsAccessToken) > 0 {
		input.AccessToken = aws.String(_verifiedpermissionsAccessToken)
	}
	if len(_verifiedpermissionsEntities) > 0 {
		if err := assignInputField(input, "Entities", _verifiedpermissionsEntities); err != nil {
			log.Errorf("invalid --entities: %s", err.Error())
			return
		}
	}
	if len(_verifiedpermissionsIdentityToken) > 0 {
		input.IdentityToken = aws.String(_verifiedpermissionsIdentityToken)
	}

	if resp, err := client.BatchIsAuthorizedWithToken(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds an identity source to a policy store–an Amazon Cognito user pool or OpenID
// Connect (OIDC) identity provider (IdP).
//
// After you create an identity source, you can use the identities provided by the
// IdP as proxies for the principal in authorization queries that use the [IsAuthorizedWithToken]or [BatchIsAuthorizedWithToken] API
// operations. These identities take the form of tokens that contain claims about
// the user, such as IDs, attributes and group memberships. Identity sources
// provide identity (ID) tokens and access tokens. Verified Permissions derives
// information about your user and session from token claims. Access tokens provide
// action context to your policies, and ID tokens provide principal Attributes .
//
// Tokens from an identity source user continue to be usable until they expire.
// Token revocation and resource deletion have no effect on the validity of a token
// in your policy store
//
// To reference a user from this identity source in your Cedar policies, refer to
// the following syntax examples.
//
// - Amazon Cognito user pool: Namespace::[Entity type]::[User pool ID]|[user
// principal attribute] , for example
// MyCorp::User::us-east-1_EXAMPLE|a1b2c3d4-5678-90ab-cdef-EXAMPLE11111 .
//
// - OpenID Connect (OIDC) provider: Namespace::[Entity
// type]::[entityIdPrefix]|[user principal attribute] , for example
// MyCorp::User::MyOIDCProvider|a1b2c3d4-5678-90ab-cdef-EXAMPLE22222 .
//
// Verified Permissions is [eventually consistent] . It can take a few seconds for a new or changed
// element to propagate through the service and be visible in the results of other
// Verified Permissions operations.
//
// [IsAuthorizedWithToken]: https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_IsAuthorizedWithToken.html
// [eventually consistent]: https://wikipedia.org/wiki/Eventual_consistency
// [BatchIsAuthorizedWithToken]: https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_BatchIsAuthorizedWithToken.html
func verifiedpermissions_CreateIdentitySource(cfg aws.Config, client *verifiedpermissions.Client) {
	input := &verifiedpermissions.CreateIdentitySourceInput{
		// Configuration: types.Configuration, // Required
		// PolicyStoreId: *string, // Required
	}

	if len(_verifiedpermissionsConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _verifiedpermissionsConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_verifiedpermissionsPolicyStoreId) > 0 {
		input.PolicyStoreId = aws.String(_verifiedpermissionsPolicyStoreId)
	}
	if len(_verifiedpermissionsClientToken) > 0 {
		input.ClientToken = aws.String(_verifiedpermissionsClientToken)
	}
	if len(_verifiedpermissionsPrincipalEntityType) > 0 {
		input.PrincipalEntityType = aws.String(_verifiedpermissionsPrincipalEntityType)
	}

	if resp, err := client.CreateIdentitySource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a Cedar policy and saves it in the specified policy store. You can
// create either a static policy or a policy linked to a policy template.
//
// - To create a static policy, provide the Cedar policy text in the StaticPolicy
// section of the PolicyDefinition .
//
// - To create a policy that is dynamically linked to a policy template, specify
// the policy template ID and the principal and resource to associate with this
// policy in the templateLinked section of the PolicyDefinition . If the policy
// template is ever updated, any policies linked to the policy template
// automatically use the updated template.
//
// Creating a policy causes it to be validated against the schema in the policy
// store. If the policy doesn't pass validation, the operation fails and the policy
// isn't stored.
//
// Verified Permissions is [eventually consistent] . It can take a few seconds for a new or changed
// element to propagate through the service and be visible in the results of other
// Verified Permissions operations.
//
// [eventually consistent]: https://wikipedia.org/wiki/Eventual_consistency
func verifiedpermissions_CreatePolicy(cfg aws.Config, client *verifiedpermissions.Client) {
	input := &verifiedpermissions.CreatePolicyInput{
		// Definition: types.PolicyDefinition, // Required
		// PolicyStoreId: *string, // Required
	}

	if len(_verifiedpermissionsDefinition) > 0 {
		if err := assignInputField(input, "Definition", _verifiedpermissionsDefinition); err != nil {
			log.Errorf("invalid --definition: %s", err.Error())
			return
		}
	}
	if len(_verifiedpermissionsPolicyStoreId) > 0 {
		input.PolicyStoreId = aws.String(_verifiedpermissionsPolicyStoreId)
	}
	if len(_verifiedpermissionsClientToken) > 0 {
		input.ClientToken = aws.String(_verifiedpermissionsClientToken)
	}

	if resp, err := client.CreatePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a policy store. A policy store is a container for policy resources.
// Although [Cedar supports multiple namespaces], Verified Permissions currently supports only one namespace per
// policy store.
//
// Verified Permissions is [eventually consistent] . It can take a few seconds for a new or changed
// element to propagate through the service and be visible in the results of other
// Verified Permissions operations.
//
// [eventually consistent]: https://wikipedia.org/wiki/Eventual_consistency
// [Cedar supports multiple namespaces]: https://docs.cedarpolicy.com/schema/schema.html#namespace
func verifiedpermissions_CreatePolicyStore(cfg aws.Config, client *verifiedpermissions.Client) {
	input := &verifiedpermissions.CreatePolicyStoreInput{
		// ValidationSettings: *types.ValidationSettings, // Required
	}

	if len(_verifiedpermissionsValidationSettings) > 0 {
		if err := assignInputField(input, "ValidationSettings", _verifiedpermissionsValidationSettings); err != nil {
			log.Errorf("invalid --validation-settings: %s", err.Error())
			return
		}
	}
	if len(_verifiedpermissionsClientToken) > 0 {
		input.ClientToken = aws.String(_verifiedpermissionsClientToken)
	}
	if len(_verifiedpermissionsDeletionProtection) > 0 {
		if err := assignInputField(input, "DeletionProtection", _verifiedpermissionsDeletionProtection); err != nil {
			log.Errorf("invalid --deletion-protection: %s", err.Error())
			return
		}
	}
	if len(_verifiedpermissionsDescription) > 0 {
		input.Description = aws.String(_verifiedpermissionsDescription)
	}
	if len(_verifiedpermissionsEncryptionSettings) > 0 {
		if err := assignInputField(input, "EncryptionSettings", _verifiedpermissionsEncryptionSettings); err != nil {
			log.Errorf("invalid --encryption-settings: %s", err.Error())
			return
		}
	}
	if len(_verifiedpermissionsTags) > 0 {
		if err := assignInputField(input, "Tags", _verifiedpermissionsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePolicyStore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a policy template. A template can use placeholders for the principal
// and resource. A template must be instantiated into a policy by associating it
// with specific principals and resources to use for the placeholders. That
// instantiated policy can then be considered in authorization decisions. The
// instantiated policy works identically to any other policy, except that it is
// dynamically linked to the template. If the template changes, then any policies
// that are linked to that template are immediately updated as well.
//
// Verified Permissions is [eventually consistent] . It can take a few seconds for a new or changed
// element to propagate through the service and be visible in the results of other
// Verified Permissions operations.
//
// [eventually consistent]: https://wikipedia.org/wiki/Eventual_consistency
func verifiedpermissions_CreatePolicyTemplate(cfg aws.Config, client *verifiedpermissions.Client) {
	input := &verifiedpermissions.CreatePolicyTemplateInput{
		// PolicyStoreId: *string, // Required
		// Statement: *string, // Required
	}

	if len(_verifiedpermissionsPolicyStoreId) > 0 {
		input.PolicyStoreId = aws.String(_verifiedpermissionsPolicyStoreId)
	}
	if len(_verifiedpermissionsStatement) > 0 {
		input.Statement = aws.String(_verifiedpermissionsStatement)
	}
	if len(_verifiedpermissionsClientToken) > 0 {
		input.ClientToken = aws.String(_verifiedpermissionsClientToken)
	}
	if len(_verifiedpermissionsDescription) > 0 {
		input.Description = aws.String(_verifiedpermissionsDescription)
	}

	if resp, err := client.CreatePolicyTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an identity source that references an identity provider (IdP) such as
// Amazon Cognito. After you delete the identity source, you can no longer use
// tokens for identities from that identity source to represent principals in
// authorization queries made using [IsAuthorizedWithToken]. operations.
//
// [IsAuthorizedWithToken]: https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_IsAuthorizedWithToken.html
func verifiedpermissions_DeleteIdentitySource(cfg aws.Config, client *verifiedpermissions.Client) {
	input := &verifiedpermissions.DeleteIdentitySourceInput{
		// IdentitySourceId: *string, // Required
		// PolicyStoreId: *string, // Required
	}

	if len(_verifiedpermissionsIdentitySourceId) > 0 {
		input.IdentitySourceId = aws.String(_verifiedpermissionsIdentitySourceId)
	}
	if len(_verifiedpermissionsPolicyStoreId) > 0 {
		input.PolicyStoreId = aws.String(_verifiedpermissionsPolicyStoreId)
	}

	if resp, err := client.DeleteIdentitySource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified policy from the policy store.
// This operation is idempotent; if you specify a policy that doesn't exist, the
// request response returns a successful HTTP 200 status code.
func verifiedpermissions_DeletePolicy(cfg aws.Config, client *verifiedpermissions.Client) {
	input := &verifiedpermissions.DeletePolicyInput{
		// PolicyId: *string, // Required
		// PolicyStoreId: *string, // Required
	}

	if len(_verifiedpermissionsPolicyId) > 0 {
		input.PolicyId = aws.String(_verifiedpermissionsPolicyId)
	}
	if len(_verifiedpermissionsPolicyStoreId) > 0 {
		input.PolicyStoreId = aws.String(_verifiedpermissionsPolicyStoreId)
	}

	if resp, err := client.DeletePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified policy store.
// This operation is idempotent. If you specify a policy store that does not
// exist, the request response will still return a successful HTTP 200 status code.
func verifiedpermissions_DeletePolicyStore(cfg aws.Config, client *verifiedpermissions.Client) {
	input := &verifiedpermissions.DeletePolicyStoreInput{
		// PolicyStoreId: *string, // Required
	}

	if len(_verifiedpermissionsPolicyStoreId) > 0 {
		input.PolicyStoreId = aws.String(_verifiedpermissionsPolicyStoreId)
	}

	if resp, err := client.DeletePolicyStore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified policy template from the policy store.
// This operation also deletes any policies that were created from the specified
// policy template. Those policies are immediately removed from all future API
// responses, and are asynchronously deleted from the policy store.
func verifiedpermissions_DeletePolicyTemplate(cfg aws.Config, client *verifiedpermissions.Client) {
	input := &verifiedpermissions.DeletePolicyTemplateInput{
		// PolicyStoreId: *string, // Required
		// PolicyTemplateId: *string, // Required
	}

	if len(_verifiedpermissionsPolicyStoreId) > 0 {
		input.PolicyStoreId = aws.String(_verifiedpermissionsPolicyStoreId)
	}
	if len(_verifiedpermissionsPolicyTemplateId) > 0 {
		input.PolicyTemplateId = aws.String(_verifiedpermissionsPolicyTemplateId)
	}

	if resp, err := client.DeletePolicyTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the details about the specified identity source.
func verifiedpermissions_GetIdentitySource(cfg aws.Config, client *verifiedpermissions.Client) {
	input := &verifiedpermissions.GetIdentitySourceInput{
		// IdentitySourceId: *string, // Required
		// PolicyStoreId: *string, // Required
	}

	if len(_verifiedpermissionsIdentitySourceId) > 0 {
		input.IdentitySourceId = aws.String(_verifiedpermissionsIdentitySourceId)
	}
	if len(_verifiedpermissionsPolicyStoreId) > 0 {
		input.PolicyStoreId = aws.String(_verifiedpermissionsPolicyStoreId)
	}

	if resp, err := client.GetIdentitySource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the specified policy.
func verifiedpermissions_GetPolicy(cfg aws.Config, client *verifiedpermissions.Client) {
	input := &verifiedpermissions.GetPolicyInput{
		// PolicyId: *string, // Required
		// PolicyStoreId: *string, // Required
	}

	if len(_verifiedpermissionsPolicyId) > 0 {
		input.PolicyId = aws.String(_verifiedpermissionsPolicyId)
	}
	if len(_verifiedpermissionsPolicyStoreId) > 0 {
		input.PolicyStoreId = aws.String(_verifiedpermissionsPolicyStoreId)
	}

	if resp, err := client.GetPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details about a policy store.
func verifiedpermissions_GetPolicyStore(cfg aws.Config, client *verifiedpermissions.Client) {
	input := &verifiedpermissions.GetPolicyStoreInput{
		// PolicyStoreId: *string, // Required
	}

	if len(_verifiedpermissionsPolicyStoreId) > 0 {
		input.PolicyStoreId = aws.String(_verifiedpermissionsPolicyStoreId)
	}
	if len(_verifiedpermissionsTags) > 0 {
		if err := assignInputField(input, "Tags", _verifiedpermissionsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetPolicyStore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieve the details for the specified policy template in the specified policy
// store.
func verifiedpermissions_GetPolicyTemplate(cfg aws.Config, client *verifiedpermissions.Client) {
	input := &verifiedpermissions.GetPolicyTemplateInput{
		// PolicyStoreId: *string, // Required
		// PolicyTemplateId: *string, // Required
	}

	if len(_verifiedpermissionsPolicyStoreId) > 0 {
		input.PolicyStoreId = aws.String(_verifiedpermissionsPolicyStoreId)
	}
	if len(_verifiedpermissionsPolicyTemplateId) > 0 {
		input.PolicyTemplateId = aws.String(_verifiedpermissionsPolicyTemplateId)
	}

	if resp, err := client.GetPolicyTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieve the details for the specified schema in the specified policy store.
func verifiedpermissions_GetSchema(cfg aws.Config, client *verifiedpermissions.Client) {
	input := &verifiedpermissions.GetSchemaInput{
		// PolicyStoreId: *string, // Required
	}

	if len(_verifiedpermissionsPolicyStoreId) > 0 {
		input.PolicyStoreId = aws.String(_verifiedpermissionsPolicyStoreId)
	}

	if resp, err := client.GetSchema(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Makes an authorization decision about a service request described in the
// parameters. The information in the parameters can also define additional context
// that Verified Permissions can include in the evaluation. The request is
// evaluated against all matching policies in the specified policy store. The
// result of the decision is either Allow or Deny , along with a list of the
// policies that resulted in the decision.
func verifiedpermissions_IsAuthorized(cfg aws.Config, client *verifiedpermissions.Client) {
	input := &verifiedpermissions.IsAuthorizedInput{
		// PolicyStoreId: *string, // Required
	}

	if len(_verifiedpermissionsPolicyStoreId) > 0 {
		input.PolicyStoreId = aws.String(_verifiedpermissionsPolicyStoreId)
	}
	if len(_verifiedpermissionsAction) > 0 {
		if err := assignInputField(input, "Action", _verifiedpermissionsAction); err != nil {
			log.Errorf("invalid --action: %s", err.Error())
			return
		}
	}
	if len(_verifiedpermissionsContext) > 0 {
		if err := assignInputField(input, "Context", _verifiedpermissionsContext); err != nil {
			log.Errorf("invalid --context: %s", err.Error())
			return
		}
	}
	if len(_verifiedpermissionsEntities) > 0 {
		if err := assignInputField(input, "Entities", _verifiedpermissionsEntities); err != nil {
			log.Errorf("invalid --entities: %s", err.Error())
			return
		}
	}
	if len(_verifiedpermissionsPrincipal) > 0 {
		if err := assignInputField(input, "Principal", _verifiedpermissionsPrincipal); err != nil {
			log.Errorf("invalid --principal: %s", err.Error())
			return
		}
	}
	if len(_verifiedpermissionsResource) > 0 {
		if err := assignInputField(input, "Resource", _verifiedpermissionsResource); err != nil {
			log.Errorf("invalid --resource: %s", err.Error())
			return
		}
	}

	if resp, err := client.IsAuthorized(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Makes an authorization decision about a service request described in the
// parameters. The principal in this request comes from an external identity source
// in the form of an identity token formatted as a [JSON web token (JWT)]. The information in the
// parameters can also define additional context that Verified Permissions can
// include in the evaluation. The request is evaluated against all matching
// policies in the specified policy store. The result of the decision is either
// Allow or Deny , along with a list of the policies that resulted in the decision.
//
// Verified Permissions validates each token that is specified in a request by
// checking its expiration date and its signature.
//
// Tokens from an identity source user continue to be usable until they expire.
// Token revocation and resource deletion have no effect on the validity of a token
// in your policy store
//
// [JSON web token (JWT)]: https://wikipedia.org/wiki/JSON_Web_Token
func verifiedpermissions_IsAuthorizedWithToken(cfg aws.Config, client *verifiedpermissions.Client) {
	input := &verifiedpermissions.IsAuthorizedWithTokenInput{
		// PolicyStoreId: *string, // Required
	}

	if len(_verifiedpermissionsPolicyStoreId) > 0 {
		input.PolicyStoreId = aws.String(_verifiedpermissionsPolicyStoreId)
	}
	if len(_verifiedpermissionsAccessToken) > 0 {
		input.AccessToken = aws.String(_verifiedpermissionsAccessToken)
	}
	if len(_verifiedpermissionsAction) > 0 {
		if err := assignInputField(input, "Action", _verifiedpermissionsAction); err != nil {
			log.Errorf("invalid --action: %s", err.Error())
			return
		}
	}
	if len(_verifiedpermissionsContext) > 0 {
		if err := assignInputField(input, "Context", _verifiedpermissionsContext); err != nil {
			log.Errorf("invalid --context: %s", err.Error())
			return
		}
	}
	if len(_verifiedpermissionsEntities) > 0 {
		if err := assignInputField(input, "Entities", _verifiedpermissionsEntities); err != nil {
			log.Errorf("invalid --entities: %s", err.Error())
			return
		}
	}
	if len(_verifiedpermissionsIdentityToken) > 0 {
		input.IdentityToken = aws.String(_verifiedpermissionsIdentityToken)
	}
	if len(_verifiedpermissionsResource) > 0 {
		if err := assignInputField(input, "Resource", _verifiedpermissionsResource); err != nil {
			log.Errorf("invalid --resource: %s", err.Error())
			return
		}
	}

	if resp, err := client.IsAuthorizedWithToken(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a paginated list of all of the identity sources defined in the
// specified policy store.
func verifiedpermissions_ListIdentitySources(cfg aws.Config, client *verifiedpermissions.Client) {
	input := &verifiedpermissions.ListIdentitySourcesInput{
		// PolicyStoreId: *string, // Required
	}

	if len(_verifiedpermissionsPolicyStoreId) > 0 {
		input.PolicyStoreId = aws.String(_verifiedpermissionsPolicyStoreId)
	}
	if len(_verifiedpermissionsFilters) > 0 {
		if err := assignInputField(input, "Filters", _verifiedpermissionsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_verifiedpermissionsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _verifiedpermissionsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_verifiedpermissionsNextToken) > 0 {
		input.NextToken = aws.String(_verifiedpermissionsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListIdentitySources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*verifiedpermissions.ListIdentitySourcesOutput
	p := verifiedpermissions.NewListIdentitySourcesPaginator(client, input)
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

// Returns a paginated list of all policies stored in the specified policy store.
func verifiedpermissions_ListPolicies(cfg aws.Config, client *verifiedpermissions.Client) {
	input := &verifiedpermissions.ListPoliciesInput{
		// PolicyStoreId: *string, // Required
	}

	if len(_verifiedpermissionsPolicyStoreId) > 0 {
		input.PolicyStoreId = aws.String(_verifiedpermissionsPolicyStoreId)
	}
	if len(_verifiedpermissionsFilter) > 0 {
		if err := assignInputField(input, "Filter", _verifiedpermissionsFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_verifiedpermissionsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _verifiedpermissionsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_verifiedpermissionsNextToken) > 0 {
		input.NextToken = aws.String(_verifiedpermissionsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPolicies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*verifiedpermissions.ListPoliciesOutput
	p := verifiedpermissions.NewListPoliciesPaginator(client, input)
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

// Returns a paginated list of all policy stores in the calling Amazon Web
// Services account.
func verifiedpermissions_ListPolicyStores(cfg aws.Config, client *verifiedpermissions.Client) {
	input := &verifiedpermissions.ListPolicyStoresInput{}

	if len(_verifiedpermissionsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _verifiedpermissionsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_verifiedpermissionsNextToken) > 0 {
		input.NextToken = aws.String(_verifiedpermissionsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPolicyStores(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*verifiedpermissions.ListPolicyStoresOutput
	p := verifiedpermissions.NewListPolicyStoresPaginator(client, input)
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

// Returns a paginated list of all policy templates in the specified policy store.
func verifiedpermissions_ListPolicyTemplates(cfg aws.Config, client *verifiedpermissions.Client) {
	input := &verifiedpermissions.ListPolicyTemplatesInput{
		// PolicyStoreId: *string, // Required
	}

	if len(_verifiedpermissionsPolicyStoreId) > 0 {
		input.PolicyStoreId = aws.String(_verifiedpermissionsPolicyStoreId)
	}
	if len(_verifiedpermissionsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _verifiedpermissionsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_verifiedpermissionsNextToken) > 0 {
		input.NextToken = aws.String(_verifiedpermissionsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPolicyTemplates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*verifiedpermissions.ListPolicyTemplatesOutput
	p := verifiedpermissions.NewListPolicyTemplatesPaginator(client, input)
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

// Returns the tags associated with the specified Amazon Verified Permissions
// resource. In Verified Permissions, policy stores can be tagged.
func verifiedpermissions_ListTagsForResource(cfg aws.Config, client *verifiedpermissions.Client) {
	input := &verifiedpermissions.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_verifiedpermissionsResourceArn) > 0 {
		input.ResourceArn = aws.String(_verifiedpermissionsResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates the policy schema in the specified policy store. The schema
// is used to validate any Cedar policies and policy templates submitted to the
// policy store. Any changes to the schema validate only policies and templates
// submitted after the schema change. Existing policies and templates are not
// re-evaluated against the changed schema. If you later update a policy, then it
// is evaluated against the new schema at that time.
//
// Verified Permissions is [eventually consistent] . It can take a few seconds for a new or changed
// element to propagate through the service and be visible in the results of other
// Verified Permissions operations.
//
// [eventually consistent]: https://wikipedia.org/wiki/Eventual_consistency
func verifiedpermissions_PutSchema(cfg aws.Config, client *verifiedpermissions.Client) {
	input := &verifiedpermissions.PutSchemaInput{
		// Definition: types.SchemaDefinition, // Required
		// PolicyStoreId: *string, // Required
	}

	if len(_verifiedpermissionsDefinition) > 0 {
		if err := assignInputField(input, "Definition", _verifiedpermissionsDefinition); err != nil {
			log.Errorf("invalid --definition: %s", err.Error())
			return
		}
	}
	if len(_verifiedpermissionsPolicyStoreId) > 0 {
		input.PolicyStoreId = aws.String(_verifiedpermissionsPolicyStoreId)
	}

	if resp, err := client.PutSchema(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assigns one or more tags (key-value pairs) to the specified Amazon Verified
// Permissions resource. Tags can help you organize and categorize your resources.
// You can also use them to scope user permissions by granting a user permission to
// access or change only resources with certain tag values. In Verified
// Permissions, policy stores can be tagged.
//
// Tags don't have any semantic meaning to Amazon Web Services and are interpreted
// strictly as strings of characters.
//
// You can use the TagResource action with a resource that already has tags. If
// you specify a new tag key, this tag is appended to the list of tags associated
// with the resource. If you specify a tag key that is already associated with the
// resource, the new tag value that you specify replaces the previous value for
// that tag.
//
// You can associate as many as 50 tags with a resource.
func verifiedpermissions_TagResource(cfg aws.Config, client *verifiedpermissions.Client) {
	input := &verifiedpermissions.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_verifiedpermissionsResourceArn) > 0 {
		input.ResourceArn = aws.String(_verifiedpermissionsResourceArn)
	}
	if len(_verifiedpermissionsTags) > 0 {
		if err := assignInputField(input, "Tags", _verifiedpermissionsTags); err != nil {
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

// Removes one or more tags from the specified Amazon Verified Permissions
// resource. In Verified Permissions, policy stores can be tagged.
func verifiedpermissions_UntagResource(cfg aws.Config, client *verifiedpermissions.Client) {
	input := &verifiedpermissions.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_verifiedpermissionsResourceArn) > 0 {
		input.ResourceArn = aws.String(_verifiedpermissionsResourceArn)
	}
	if len(_verifiedpermissionsTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _verifiedpermissionsTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified identity source to use a new identity provider (IdP), or
// to change the mapping of identities from the IdP to a different principal entity
// type.
//
// Verified Permissions is [eventually consistent] . It can take a few seconds for a new or changed
// element to propagate through the service and be visible in the results of other
// Verified Permissions operations.
//
// [eventually consistent]: https://wikipedia.org/wiki/Eventual_consistency
func verifiedpermissions_UpdateIdentitySource(cfg aws.Config, client *verifiedpermissions.Client) {
	input := &verifiedpermissions.UpdateIdentitySourceInput{
		// IdentitySourceId: *string, // Required
		// PolicyStoreId: *string, // Required
		// UpdateConfiguration: types.UpdateConfiguration, // Required
	}

	if len(_verifiedpermissionsIdentitySourceId) > 0 {
		input.IdentitySourceId = aws.String(_verifiedpermissionsIdentitySourceId)
	}
	if len(_verifiedpermissionsPolicyStoreId) > 0 {
		input.PolicyStoreId = aws.String(_verifiedpermissionsPolicyStoreId)
	}
	if len(_verifiedpermissionsUpdateConfiguration) > 0 {
		if err := assignInputField(input, "UpdateConfiguration", _verifiedpermissionsUpdateConfiguration); err != nil {
			log.Errorf("invalid --update-configuration: %s", err.Error())
			return
		}
	}
	if len(_verifiedpermissionsPrincipalEntityType) > 0 {
		input.PrincipalEntityType = aws.String(_verifiedpermissionsPrincipalEntityType)
	}

	if resp, err := client.UpdateIdentitySource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies a Cedar static policy in the specified policy store. You can change
// only certain elements of the [UpdatePolicyDefinition]parameter. You can directly update only static
// policies. To change a template-linked policy, you must update the template
// instead, using [UpdatePolicyTemplate].
//
// - If policy validation is enabled in the policy store, then updating a static
// policy causes Verified Permissions to validate the policy against the schema in
// the policy store. If the updated static policy doesn't pass validation, the
// operation fails and the update isn't stored.
//
// - When you edit a static policy, you can change only certain elements of a
// static policy:
//
// - The action referenced by the policy.
//
// - A condition clause, such as when and unless.
//
// You can't change these elements of a static policy:
//
// - Changing a policy from a static policy to a template-linked policy.
//
// - Changing the effect of a static policy from permit or forbid.
//
// - The principal referenced by a static policy.
//
// - The resource referenced by a static policy.
//
// - To update a template-linked policy, you must update the template instead.
//
// Verified Permissions is [eventually consistent] . It can take a few seconds for a new or changed
// element to propagate through the service and be visible in the results of other
// Verified Permissions operations.
//
// [eventually consistent]: https://wikipedia.org/wiki/Eventual_consistency
// [UpdatePolicyTemplate]: https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_UpdatePolicyTemplate.html
// [UpdatePolicyDefinition]: https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_UpdatePolicyInput.html#amazonverifiedpermissions-UpdatePolicy-request-UpdatePolicyDefinition
func verifiedpermissions_UpdatePolicy(cfg aws.Config, client *verifiedpermissions.Client) {
	input := &verifiedpermissions.UpdatePolicyInput{
		// PolicyId: *string, // Required
		// PolicyStoreId: *string, // Required
	}

	if len(_verifiedpermissionsPolicyId) > 0 {
		input.PolicyId = aws.String(_verifiedpermissionsPolicyId)
	}
	if len(_verifiedpermissionsPolicyStoreId) > 0 {
		input.PolicyStoreId = aws.String(_verifiedpermissionsPolicyStoreId)
	}
	if len(_verifiedpermissionsDefinition) > 0 {
		if err := assignInputField(input, "Definition", _verifiedpermissionsDefinition); err != nil {
			log.Errorf("invalid --definition: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdatePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the validation setting for a policy store.
// Verified Permissions is [eventually consistent] . It can take a few seconds for a new or changed
// element to propagate through the service and be visible in the results of other
// Verified Permissions operations.
//
// [eventually consistent]: https://wikipedia.org/wiki/Eventual_consistency
func verifiedpermissions_UpdatePolicyStore(cfg aws.Config, client *verifiedpermissions.Client) {
	input := &verifiedpermissions.UpdatePolicyStoreInput{
		// PolicyStoreId: *string, // Required
		// ValidationSettings: *types.ValidationSettings, // Required
	}

	if len(_verifiedpermissionsPolicyStoreId) > 0 {
		input.PolicyStoreId = aws.String(_verifiedpermissionsPolicyStoreId)
	}
	if len(_verifiedpermissionsValidationSettings) > 0 {
		if err := assignInputField(input, "ValidationSettings", _verifiedpermissionsValidationSettings); err != nil {
			log.Errorf("invalid --validation-settings: %s", err.Error())
			return
		}
	}
	if len(_verifiedpermissionsDeletionProtection) > 0 {
		if err := assignInputField(input, "DeletionProtection", _verifiedpermissionsDeletionProtection); err != nil {
			log.Errorf("invalid --deletion-protection: %s", err.Error())
			return
		}
	}
	if len(_verifiedpermissionsDescription) > 0 {
		input.Description = aws.String(_verifiedpermissionsDescription)
	}

	if resp, err := client.UpdatePolicyStore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified policy template. You can update only the description and
// the some elements of the [policyBody].
//
// Changes you make to the policy template content are immediately (within the
// constraints of eventual consistency) reflected in authorization decisions that
// involve all template-linked policies instantiated from this template.
//
// Verified Permissions is [eventually consistent] . It can take a few seconds for a new or changed
// element to propagate through the service and be visible in the results of other
// Verified Permissions operations.
//
// [eventually consistent]: https://wikipedia.org/wiki/Eventual_consistency
// [policyBody]: https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_UpdatePolicyTemplate.html#amazonverifiedpermissions-UpdatePolicyTemplate-request-policyBody
func verifiedpermissions_UpdatePolicyTemplate(cfg aws.Config, client *verifiedpermissions.Client) {
	input := &verifiedpermissions.UpdatePolicyTemplateInput{
		// PolicyStoreId: *string, // Required
		// PolicyTemplateId: *string, // Required
		// Statement: *string, // Required
	}

	if len(_verifiedpermissionsPolicyStoreId) > 0 {
		input.PolicyStoreId = aws.String(_verifiedpermissionsPolicyStoreId)
	}
	if len(_verifiedpermissionsPolicyTemplateId) > 0 {
		input.PolicyTemplateId = aws.String(_verifiedpermissionsPolicyTemplateId)
	}
	if len(_verifiedpermissionsStatement) > 0 {
		input.Statement = aws.String(_verifiedpermissionsStatement)
	}
	if len(_verifiedpermissionsDescription) > 0 {
		input.Description = aws.String(_verifiedpermissionsDescription)
	}

	if resp, err := client.UpdatePolicyTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_verifiedpermissionsCmd)
	_verifiedpermissionsCmd.Flags().SortFlags = false

	_verifiedpermissionsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_verifiedpermissionsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_verifiedpermissionsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_verifiedpermissionsCmd.Flags().StringVarP(&_verifiedpermissionsAccessToken, "access-token", "", "", "Access Token")
	_verifiedpermissionsCmd.Flags().StringVarP(&_verifiedpermissionsAction, "action", "", "", "Action")
	_verifiedpermissionsCmd.Flags().StringVarP(&_verifiedpermissionsClientToken, "client-token", "", "", "Client Token")
	_verifiedpermissionsCmd.Flags().StringVarP(&_verifiedpermissionsConfiguration, "configuration", "", "", "Configuration")
	_verifiedpermissionsCmd.Flags().StringVarP(&_verifiedpermissionsContext, "context", "", "", "Context")
	_verifiedpermissionsCmd.Flags().StringVarP(&_verifiedpermissionsDefinition, "definition", "", "", "Definition")
	_verifiedpermissionsCmd.Flags().StringVarP(&_verifiedpermissionsDeletionProtection, "deletion-protection", "", "", "Deletion Protection")
	_verifiedpermissionsCmd.Flags().StringVarP(&_verifiedpermissionsDescription, "description", "", "", "Description")
	_verifiedpermissionsCmd.Flags().StringVarP(&_verifiedpermissionsEncryptionSettings, "encryption-settings", "", "", "Encryption Settings")
	_verifiedpermissionsCmd.Flags().StringVarP(&_verifiedpermissionsEntities, "entities", "", "", "Entities")
	_verifiedpermissionsCmd.Flags().StringVarP(&_verifiedpermissionsFilter, "filter", "", "", "Filter")
	_verifiedpermissionsCmd.Flags().StringVarP(&_verifiedpermissionsFilters, "filters", "", "", "Filters")
	_verifiedpermissionsCmd.Flags().StringVarP(&_verifiedpermissionsIdentitySourceId, "identity-source-id", "", "", "Identity Source ID")
	_verifiedpermissionsCmd.Flags().StringVarP(&_verifiedpermissionsIdentityToken, "identity-token", "", "", "Identity Token")
	_verifiedpermissionsCmd.Flags().StringVarP(&_verifiedpermissionsMaxResults, "max-results", "", "", "Max Results")
	_verifiedpermissionsCmd.Flags().StringVarP(&_verifiedpermissionsNextToken, "next-token", "", "", "Next Token")
	_verifiedpermissionsCmd.Flags().StringVarP(&_verifiedpermissionsPolicyId, "policy-id", "", "", "Policy ID")
	_verifiedpermissionsCmd.Flags().StringVarP(&_verifiedpermissionsPolicyStoreId, "policy-store-id", "", "", "Policy Store ID")
	_verifiedpermissionsCmd.Flags().StringVarP(&_verifiedpermissionsPolicyTemplateId, "policy-template-id", "", "", "Policy Template ID")
	_verifiedpermissionsCmd.Flags().StringVarP(&_verifiedpermissionsPrincipal, "principal", "", "", "Principal")
	_verifiedpermissionsCmd.Flags().StringVarP(&_verifiedpermissionsPrincipalEntityType, "principal-entity-type", "", "", "Principal Entity Type")
	_verifiedpermissionsCmd.Flags().StringVarP(&_verifiedpermissionsRequests, "requests", "", "", "Requests")
	_verifiedpermissionsCmd.Flags().StringVarP(&_verifiedpermissionsResource, "resource", "", "", "Resource")
	_verifiedpermissionsCmd.Flags().StringVarP(&_verifiedpermissionsResourceArn, "resource-arn", "", "", "Resource ARN")
	_verifiedpermissionsCmd.Flags().StringVarP(&_verifiedpermissionsStatement, "statement", "", "", "Statement")
	_verifiedpermissionsCmd.Flags().StringSliceVarP(&_verifiedpermissionsTagKeys, "tag-keys", "", nil, "Tag Keys")
	_verifiedpermissionsCmd.Flags().StringVarP(&_verifiedpermissionsTags, "tags", "", "", "Tags")
	_verifiedpermissionsCmd.Flags().StringVarP(&_verifiedpermissionsUpdateConfiguration, "update-configuration", "", "", "Update Configuration")
	_verifiedpermissionsCmd.Flags().StringVarP(&_verifiedpermissionsValidationSettings, "validation-settings", "", "", "Validation Settings")

	_verifiedpermissionsCmd.Flags().BoolVarP(&_verifiedpermissionsBatchGetPolicy, "batch-get-policy", "", false, "Batch Get Policy")
	_verifiedpermissionsCmd.Flags().BoolVarP(&_verifiedpermissionsBatchIsAuthorized, "batch-is-authorized", "", false, "Batch Is Authorized")
	_verifiedpermissionsCmd.Flags().BoolVarP(&_verifiedpermissionsBatchIsAuthorizedWithToken, "batch-is-authorized-with-token", "", false, "Batch Is Authorized With Token")
	_verifiedpermissionsCmd.Flags().BoolVarP(&_verifiedpermissionsCreateIdentitySource, "create-identity-source", "", false, "Create Identity Source")
	_verifiedpermissionsCmd.Flags().BoolVarP(&_verifiedpermissionsCreatePolicy, "create-policy", "", false, "Create Policy")
	_verifiedpermissionsCmd.Flags().BoolVarP(&_verifiedpermissionsCreatePolicyStore, "create-policy-store", "", false, "Create Policy Store")
	_verifiedpermissionsCmd.Flags().BoolVarP(&_verifiedpermissionsCreatePolicyTemplate, "create-policy-template", "", false, "Create Policy Template")
	_verifiedpermissionsCmd.Flags().BoolVarP(&_verifiedpermissionsDeleteIdentitySource, "delete-identity-source", "", false, "Delete Identity Source")
	_verifiedpermissionsCmd.Flags().BoolVarP(&_verifiedpermissionsDeletePolicy, "delete-policy", "", false, "Delete Policy")
	_verifiedpermissionsCmd.Flags().BoolVarP(&_verifiedpermissionsDeletePolicyStore, "delete-policy-store", "", false, "Delete Policy Store")
	_verifiedpermissionsCmd.Flags().BoolVarP(&_verifiedpermissionsDeletePolicyTemplate, "delete-policy-template", "", false, "Delete Policy Template")
	_verifiedpermissionsCmd.Flags().BoolVarP(&_verifiedpermissionsGetIdentitySource, "get-identity-source", "", false, "Get Identity Source")
	_verifiedpermissionsCmd.Flags().BoolVarP(&_verifiedpermissionsGetPolicy, "get-policy", "", false, "Get Policy")
	_verifiedpermissionsCmd.Flags().BoolVarP(&_verifiedpermissionsGetPolicyStore, "get-policy-store", "", false, "Get Policy Store")
	_verifiedpermissionsCmd.Flags().BoolVarP(&_verifiedpermissionsGetPolicyTemplate, "get-policy-template", "", false, "Get Policy Template")
	_verifiedpermissionsCmd.Flags().BoolVarP(&_verifiedpermissionsGetSchema, "get-schema", "", false, "Get Schema")
	_verifiedpermissionsCmd.Flags().BoolVarP(&_verifiedpermissionsIsAuthorized, "is-authorized", "", false, "Is Authorized")
	_verifiedpermissionsCmd.Flags().BoolVarP(&_verifiedpermissionsIsAuthorizedWithToken, "is-authorized-with-token", "", false, "Is Authorized With Token")
	_verifiedpermissionsCmd.Flags().BoolVarP(&_verifiedpermissionsListIdentitySources, "list-identity-sources", "", false, "List Identity Sources")
	_verifiedpermissionsCmd.Flags().BoolVarP(&_verifiedpermissionsListPolicies, "list-policies", "", false, "List Policies")
	_verifiedpermissionsCmd.Flags().BoolVarP(&_verifiedpermissionsListPolicyStores, "list-policy-stores", "", false, "List Policy Stores")
	_verifiedpermissionsCmd.Flags().BoolVarP(&_verifiedpermissionsListPolicyTemplates, "list-policy-templates", "", false, "List Policy Templates")
	_verifiedpermissionsCmd.Flags().BoolVarP(&_verifiedpermissionsListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_verifiedpermissionsCmd.Flags().BoolVarP(&_verifiedpermissionsPutSchema, "put-schema", "", false, "Put Schema")
	_verifiedpermissionsCmd.Flags().BoolVarP(&_verifiedpermissionsTagResource, "tag-resource", "", false, "Tag Resource")
	_verifiedpermissionsCmd.Flags().BoolVarP(&_verifiedpermissionsUntagResource, "untag-resource", "", false, "Untag Resource")
	_verifiedpermissionsCmd.Flags().BoolVarP(&_verifiedpermissionsUpdateIdentitySource, "update-identity-source", "", false, "Update Identity Source")
	_verifiedpermissionsCmd.Flags().BoolVarP(&_verifiedpermissionsUpdatePolicy, "update-policy", "", false, "Update Policy")
	_verifiedpermissionsCmd.Flags().BoolVarP(&_verifiedpermissionsUpdatePolicyStore, "update-policy-store", "", false, "Update Policy Store")
	_verifiedpermissionsCmd.Flags().BoolVarP(&_verifiedpermissionsUpdatePolicyTemplate, "update-policy-template", "", false, "Update Policy Template")

}
