package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// cognitoidentityCmd represents the cognitoidentity command
var _cognitoidentityCmd = &cobra.Command{
	Use:   "cognitoidentity",
	Short: "AWS cognitoidentity CLI",
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
		client := cognitoidentity.NewFromConfig(cfg)
		if _cognitoidentityCreateIdentityPool {
			cognitoidentity_CreateIdentityPool(cfg, client)
			return
		}
		if _cognitoidentityDeleteIdentities {
			cognitoidentity_DeleteIdentities(cfg, client)
			return
		}
		if _cognitoidentityDeleteIdentityPool {
			cognitoidentity_DeleteIdentityPool(cfg, client)
			return
		}
		if _cognitoidentityDescribeIdentity {
			cognitoidentity_DescribeIdentity(cfg, client)
			return
		}
		if _cognitoidentityDescribeIdentityPool {
			cognitoidentity_DescribeIdentityPool(cfg, client)
			return
		}
		if _cognitoidentityGetCredentialsForIdentity {
			cognitoidentity_GetCredentialsForIdentity(cfg, client)
			return
		}
		if _cognitoidentityGetId {
			cognitoidentity_GetId(cfg, client)
			return
		}
		if _cognitoidentityGetIdentityPoolRoles {
			cognitoidentity_GetIdentityPoolRoles(cfg, client)
			return
		}
		if _cognitoidentityGetOpenIdToken {
			cognitoidentity_GetOpenIdToken(cfg, client)
			return
		}
		if _cognitoidentityGetOpenIdTokenForDeveloperIdentity {
			cognitoidentity_GetOpenIdTokenForDeveloperIdentity(cfg, client)
			return
		}
		if _cognitoidentityGetPrincipalTagAttributeMap {
			cognitoidentity_GetPrincipalTagAttributeMap(cfg, client)
			return
		}
		if _cognitoidentityListIdentities {
			cognitoidentity_ListIdentities(cfg, client)
			return
		}
		if _cognitoidentityListIdentityPools {
			cognitoidentity_ListIdentityPools(cfg, client)
			return
		}
		if _cognitoidentityListTagsForResource {
			cognitoidentity_ListTagsForResource(cfg, client)
			return
		}
		if _cognitoidentityLookupDeveloperIdentity {
			cognitoidentity_LookupDeveloperIdentity(cfg, client)
			return
		}
		if _cognitoidentityMergeDeveloperIdentities {
			cognitoidentity_MergeDeveloperIdentities(cfg, client)
			return
		}
		if _cognitoidentitySetIdentityPoolRoles {
			cognitoidentity_SetIdentityPoolRoles(cfg, client)
			return
		}
		if _cognitoidentitySetPrincipalTagAttributeMap {
			cognitoidentity_SetPrincipalTagAttributeMap(cfg, client)
			return
		}
		if _cognitoidentityTagResource {
			cognitoidentity_TagResource(cfg, client)
			return
		}
		if _cognitoidentityUnlinkDeveloperIdentity {
			cognitoidentity_UnlinkDeveloperIdentity(cfg, client)
			return
		}
		if _cognitoidentityUnlinkIdentity {
			cognitoidentity_UnlinkIdentity(cfg, client)
			return
		}
		if _cognitoidentityUntagResource {
			cognitoidentity_UntagResource(cfg, client)
			return
		}
		if _cognitoidentityUpdateIdentityPool {
			cognitoidentity_UpdateIdentityPool(cfg, client)
			return
		}

	},
}

var (
	_cognitoidentityCreateIdentityPool                 bool
	_cognitoidentityDeleteIdentities                   bool
	_cognitoidentityDeleteIdentityPool                 bool
	_cognitoidentityDescribeIdentity                   bool
	_cognitoidentityDescribeIdentityPool               bool
	_cognitoidentityGetCredentialsForIdentity          bool
	_cognitoidentityGetId                              bool
	_cognitoidentityGetIdentityPoolRoles               bool
	_cognitoidentityGetOpenIdToken                     bool
	_cognitoidentityGetOpenIdTokenForDeveloperIdentity bool
	_cognitoidentityGetPrincipalTagAttributeMap        bool
	_cognitoidentityListIdentities                     bool
	_cognitoidentityListIdentityPools                  bool
	_cognitoidentityListTagsForResource                bool
	_cognitoidentityLookupDeveloperIdentity            bool
	_cognitoidentityMergeDeveloperIdentities           bool
	_cognitoidentitySetIdentityPoolRoles               bool
	_cognitoidentitySetPrincipalTagAttributeMap        bool
	_cognitoidentityTagResource                        bool
	_cognitoidentityUnlinkDeveloperIdentity            bool
	_cognitoidentityUnlinkIdentity                     bool
	_cognitoidentityUntagResource                      bool
	_cognitoidentityUpdateIdentityPool                 bool

	_cognitoidentityAccountId                      string
	_cognitoidentityAllowClassicFlow               string
	_cognitoidentityAllowUnauthenticatedIdentities string
	_cognitoidentityCognitoIdentityProviders       string
	_cognitoidentityCustomRoleArn                  string
	_cognitoidentityDestinationUserIdentifier      string
	_cognitoidentityDeveloperProviderName          string
	_cognitoidentityDeveloperUserIdentifier        string
	_cognitoidentityHideDisabled                   string
	_cognitoidentityIdentityId                     string
	_cognitoidentityIdentityIdsToDelete            []string
	_cognitoidentityIdentityPoolId                 string
	_cognitoidentityIdentityPoolName               string
	_cognitoidentityIdentityPoolTags               string
	_cognitoidentityIdentityProviderName           string
	_cognitoidentityLogins                         string
	_cognitoidentityLoginsToRemove                 []string
	_cognitoidentityMaxResults                     string
	_cognitoidentityNextToken                      string
	_cognitoidentityOpenIdConnectProviderARNs      []string
	_cognitoidentityPrincipalTags                  string
	_cognitoidentityResourceArn                    string
	_cognitoidentityRoleMappings                   string
	_cognitoidentityRoles                          string
	_cognitoidentitySamlProviderARNs               []string
	_cognitoidentitySourceUserIdentifier           string
	_cognitoidentitySupportedLoginProviders        string
	_cognitoidentityTagKeys                        []string
	_cognitoidentityTags                           string
	_cognitoidentityTokenDuration                  string
	_cognitoidentityUseDefaults                    string
)

// Creates a new identity pool. The identity pool is a store of user identity
// information that is specific to your Amazon Web Services account. The keys for
// SupportedLoginProviders are as follows:
//
// - Facebook: graph.facebook.com
//
// - Google: accounts.google.com
//
// - Sign in With Apple: appleid.apple.com
//
// - Amazon: www.amazon.com
//
// - Twitter: api.twitter.com
//
// - Digits: www.digits.com
//
// If you don't provide a value for a parameter, Amazon Cognito sets it to its
// default value.
//
// You must use Amazon Web Services developer credentials to call this operation.
func cognitoidentity_CreateIdentityPool(cfg aws.Config, client *cognitoidentity.Client) {
	input := &cognitoidentity.CreateIdentityPoolInput{
		// AllowUnauthenticatedIdentities: bool, // Required
		// IdentityPoolName: *string, // Required
	}

	if len(_cognitoidentityAllowUnauthenticatedIdentities) > 0 {
		if err := assignInputField(input, "AllowUnauthenticatedIdentities", _cognitoidentityAllowUnauthenticatedIdentities); err != nil {
			log.Errorf("invalid --allow-unauthenticated-identities: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityIdentityPoolName) > 0 {
		input.IdentityPoolName = aws.String(_cognitoidentityIdentityPoolName)
	}
	if len(_cognitoidentityAllowClassicFlow) > 0 {
		if err := assignInputField(input, "AllowClassicFlow", _cognitoidentityAllowClassicFlow); err != nil {
			log.Errorf("invalid --allow-classic-flow: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityCognitoIdentityProviders) > 0 {
		if err := assignInputField(input, "CognitoIdentityProviders", _cognitoidentityCognitoIdentityProviders); err != nil {
			log.Errorf("invalid --cognito-identity-providers: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityDeveloperProviderName) > 0 {
		input.DeveloperProviderName = aws.String(_cognitoidentityDeveloperProviderName)
	}
	if len(_cognitoidentityIdentityPoolTags) > 0 {
		if err := assignInputField(input, "IdentityPoolTags", _cognitoidentityIdentityPoolTags); err != nil {
			log.Errorf("invalid --identity-pool-tags: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityOpenIdConnectProviderARNs) > 0 {
		input.OpenIdConnectProviderARNs = append([]string(nil), _cognitoidentityOpenIdConnectProviderARNs...)
	}
	if len(_cognitoidentitySamlProviderARNs) > 0 {
		input.SamlProviderARNs = append([]string(nil), _cognitoidentitySamlProviderARNs...)
	}
	if len(_cognitoidentitySupportedLoginProviders) > 0 {
		if err := assignInputField(input, "SupportedLoginProviders", _cognitoidentitySupportedLoginProviders); err != nil {
			log.Errorf("invalid --supported-login-providers: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateIdentityPool(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes identities from an identity pool. You can specify a list of 1-60
// identities that you want to delete.
//
// You must use Amazon Web Services developer credentials to call this operation.
func cognitoidentity_DeleteIdentities(cfg aws.Config, client *cognitoidentity.Client) {
	input := &cognitoidentity.DeleteIdentitiesInput{
		// IdentityIdsToDelete: []string, // Required
	}

	if len(_cognitoidentityIdentityIdsToDelete) > 0 {
		input.IdentityIdsToDelete = append([]string(nil), _cognitoidentityIdentityIdsToDelete...)
	}

	if resp, err := client.DeleteIdentities(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an identity pool. Once a pool is deleted, users will not be able to
// authenticate with the pool.
//
// You must use Amazon Web Services developer credentials to call this operation.
func cognitoidentity_DeleteIdentityPool(cfg aws.Config, client *cognitoidentity.Client) {
	input := &cognitoidentity.DeleteIdentityPoolInput{
		// IdentityPoolId: *string, // Required
	}

	if len(_cognitoidentityIdentityPoolId) > 0 {
		input.IdentityPoolId = aws.String(_cognitoidentityIdentityPoolId)
	}

	if resp, err := client.DeleteIdentityPool(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns metadata related to the given identity, including when the identity was
// created and any associated linked logins.
//
// You must use Amazon Web Services developer credentials to call this operation.
func cognitoidentity_DescribeIdentity(cfg aws.Config, client *cognitoidentity.Client) {
	input := &cognitoidentity.DescribeIdentityInput{
		// IdentityId: *string, // Required
	}

	if len(_cognitoidentityIdentityId) > 0 {
		input.IdentityId = aws.String(_cognitoidentityIdentityId)
	}

	if resp, err := client.DescribeIdentity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets details about a particular identity pool, including the pool name, ID
// description, creation date, and current number of users.
//
// You must use Amazon Web Services developer credentials to call this operation.
func cognitoidentity_DescribeIdentityPool(cfg aws.Config, client *cognitoidentity.Client) {
	input := &cognitoidentity.DescribeIdentityPoolInput{
		// IdentityPoolId: *string, // Required
	}

	if len(_cognitoidentityIdentityPoolId) > 0 {
		input.IdentityPoolId = aws.String(_cognitoidentityIdentityPoolId)
	}

	if resp, err := client.DescribeIdentityPool(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns credentials for the provided identity ID. Any provided logins will be
// validated against supported login providers. If the token is for
// cognito-identity.amazonaws.com , it will be passed through to Security Token
// Service with the appropriate role for the token.
//
// This is a public API. You do not need any credentials to call this API.
func cognitoidentity_GetCredentialsForIdentity(cfg aws.Config, client *cognitoidentity.Client) {
	input := &cognitoidentity.GetCredentialsForIdentityInput{
		// IdentityId: *string, // Required
	}

	if len(_cognitoidentityIdentityId) > 0 {
		input.IdentityId = aws.String(_cognitoidentityIdentityId)
	}
	if len(_cognitoidentityCustomRoleArn) > 0 {
		input.CustomRoleArn = aws.String(_cognitoidentityCustomRoleArn)
	}
	if len(_cognitoidentityLogins) > 0 {
		if err := assignInputField(input, "Logins", _cognitoidentityLogins); err != nil {
			log.Errorf("invalid --logins: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetCredentialsForIdentity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Generates (or retrieves) IdentityID. Supplying multiple logins will create an
// implicit linked account.
//
// This is a public API. You do not need any credentials to call this API.
func cognitoidentity_GetId(cfg aws.Config, client *cognitoidentity.Client) {
	input := &cognitoidentity.GetIdInput{
		// IdentityPoolId: *string, // Required
	}

	if len(_cognitoidentityIdentityPoolId) > 0 {
		input.IdentityPoolId = aws.String(_cognitoidentityIdentityPoolId)
	}
	if len(_cognitoidentityAccountId) > 0 {
		input.AccountId = aws.String(_cognitoidentityAccountId)
	}
	if len(_cognitoidentityLogins) > 0 {
		if err := assignInputField(input, "Logins", _cognitoidentityLogins); err != nil {
			log.Errorf("invalid --logins: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetId(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the roles for an identity pool.
// You must use Amazon Web Services developer credentials to call this operation.
func cognitoidentity_GetIdentityPoolRoles(cfg aws.Config, client *cognitoidentity.Client) {
	input := &cognitoidentity.GetIdentityPoolRolesInput{
		// IdentityPoolId: *string, // Required
	}

	if len(_cognitoidentityIdentityPoolId) > 0 {
		input.IdentityPoolId = aws.String(_cognitoidentityIdentityPoolId)
	}

	if resp, err := client.GetIdentityPoolRoles(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets an OpenID token, using a known Cognito ID. This known Cognito ID is
// returned by GetId. You can optionally add additional logins for the identity.
// Supplying multiple logins creates an implicit link.
//
// The OpenID token is valid for 10 minutes.
//
// This is a public API. You do not need any credentials to call this API.
func cognitoidentity_GetOpenIdToken(cfg aws.Config, client *cognitoidentity.Client) {
	input := &cognitoidentity.GetOpenIdTokenInput{
		// IdentityId: *string, // Required
	}

	if len(_cognitoidentityIdentityId) > 0 {
		input.IdentityId = aws.String(_cognitoidentityIdentityId)
	}
	if len(_cognitoidentityLogins) > 0 {
		if err := assignInputField(input, "Logins", _cognitoidentityLogins); err != nil {
			log.Errorf("invalid --logins: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetOpenIdToken(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Registers (or retrieves) a Cognito IdentityId and an OpenID Connect token for a
// user authenticated by your backend authentication process. Supplying multiple
// logins will create an implicit linked account. You can only specify one
// developer provider as part of the Logins map, which is linked to the identity
// pool. The developer provider is the "domain" by which Cognito will refer to your
// users.
//
// You can use GetOpenIdTokenForDeveloperIdentity to create a new identity and to
// link new logins (that is, user credentials issued by a public provider or
// developer provider) to an existing identity. When you want to create a new
// identity, the IdentityId should be null. When you want to associate a new login
// with an existing authenticated/unauthenticated identity, you can do so by
// providing the existing IdentityId . This API will create the identity in the
// specified IdentityPoolId .
//
// You must use Amazon Web Services developer credentials to call this operation.
func cognitoidentity_GetOpenIdTokenForDeveloperIdentity(cfg aws.Config, client *cognitoidentity.Client) {
	input := &cognitoidentity.GetOpenIdTokenForDeveloperIdentityInput{
		// IdentityPoolId: *string, // Required
		// Logins: map[string]string, // Required
	}

	if len(_cognitoidentityIdentityPoolId) > 0 {
		input.IdentityPoolId = aws.String(_cognitoidentityIdentityPoolId)
	}
	if len(_cognitoidentityLogins) > 0 {
		if err := assignInputField(input, "Logins", _cognitoidentityLogins); err != nil {
			log.Errorf("invalid --logins: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityIdentityId) > 0 {
		input.IdentityId = aws.String(_cognitoidentityIdentityId)
	}
	if len(_cognitoidentityPrincipalTags) > 0 {
		if err := assignInputField(input, "PrincipalTags", _cognitoidentityPrincipalTags); err != nil {
			log.Errorf("invalid --principal-tags: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityTokenDuration) > 0 {
		if err := assignInputField(input, "TokenDuration", _cognitoidentityTokenDuration); err != nil {
			log.Errorf("invalid --token-duration: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetOpenIdTokenForDeveloperIdentity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use GetPrincipalTagAttributeMap to list all mappings between PrincipalTags and
// user attributes.
func cognitoidentity_GetPrincipalTagAttributeMap(cfg aws.Config, client *cognitoidentity.Client) {
	input := &cognitoidentity.GetPrincipalTagAttributeMapInput{
		// IdentityPoolId: *string, // Required
		// IdentityProviderName: *string, // Required
	}

	if len(_cognitoidentityIdentityPoolId) > 0 {
		input.IdentityPoolId = aws.String(_cognitoidentityIdentityPoolId)
	}
	if len(_cognitoidentityIdentityProviderName) > 0 {
		input.IdentityProviderName = aws.String(_cognitoidentityIdentityProviderName)
	}

	if resp, err := client.GetPrincipalTagAttributeMap(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the identities in an identity pool.
// You must use Amazon Web Services developer credentials to call this operation.
func cognitoidentity_ListIdentities(cfg aws.Config, client *cognitoidentity.Client) {
	input := &cognitoidentity.ListIdentitiesInput{
		// IdentityPoolId: *string, // Required
		// MaxResults: *int32, // Required
	}

	if len(_cognitoidentityIdentityPoolId) > 0 {
		input.IdentityPoolId = aws.String(_cognitoidentityIdentityPoolId)
	}
	if len(_cognitoidentityMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cognitoidentityMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityHideDisabled) > 0 {
		if err := assignInputField(input, "HideDisabled", _cognitoidentityHideDisabled); err != nil {
			log.Errorf("invalid --hide-disabled: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityNextToken) > 0 {
		input.NextToken = aws.String(_cognitoidentityNextToken)
	}

	if resp, err := client.ListIdentities(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all of the Cognito identity pools registered for your account.
// You must use Amazon Web Services developer credentials to call this operation.
func cognitoidentity_ListIdentityPools(cfg aws.Config, client *cognitoidentity.Client) {
	input := &cognitoidentity.ListIdentityPoolsInput{
		// MaxResults: *int32, // Required
	}

	if len(_cognitoidentityMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cognitoidentityMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityNextToken) > 0 {
		input.NextToken = aws.String(_cognitoidentityNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListIdentityPools(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cognitoidentity.ListIdentityPoolsOutput
	p := cognitoidentity.NewListIdentityPoolsPaginator(client, input)
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

// Lists the tags that are assigned to an Amazon Cognito identity pool.
// A tag is a label that you can apply to identity pools to categorize and manage
// them in different ways, such as by purpose, owner, environment, or other
// criteria.
//
// You can use this action up to 10 times per second, per account.
func cognitoidentity_ListTagsForResource(cfg aws.Config, client *cognitoidentity.Client) {
	input := &cognitoidentity.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_cognitoidentityResourceArn) > 0 {
		input.ResourceArn = aws.String(_cognitoidentityResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the IdentityID associated with a DeveloperUserIdentifier or the list
// of DeveloperUserIdentifier values associated with an IdentityId for an existing
// identity. Either IdentityID or DeveloperUserIdentifier must not be null. If you
// supply only one of these values, the other value will be searched in the
// database and returned as a part of the response. If you supply both,
// DeveloperUserIdentifier will be matched against IdentityID . If the values are
// verified against the database, the response returns both values and is the same
// as the request. Otherwise, a ResourceConflictException is thrown.
//
// LookupDeveloperIdentity is intended for low-throughput control plane
// operations: for example, to enable customer service to locate an identity ID by
// username. If you are using it for higher-volume operations such as user
// authentication, your requests are likely to be throttled. GetOpenIdTokenForDeveloperIdentityis a better option
// for higher-volume operations for user authentication.
//
// You must use Amazon Web Services developer credentials to call this operation.
func cognitoidentity_LookupDeveloperIdentity(cfg aws.Config, client *cognitoidentity.Client) {
	input := &cognitoidentity.LookupDeveloperIdentityInput{
		// IdentityPoolId: *string, // Required
	}

	if len(_cognitoidentityIdentityPoolId) > 0 {
		input.IdentityPoolId = aws.String(_cognitoidentityIdentityPoolId)
	}
	if len(_cognitoidentityDeveloperUserIdentifier) > 0 {
		input.DeveloperUserIdentifier = aws.String(_cognitoidentityDeveloperUserIdentifier)
	}
	if len(_cognitoidentityIdentityId) > 0 {
		input.IdentityId = aws.String(_cognitoidentityIdentityId)
	}
	if len(_cognitoidentityMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cognitoidentityMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityNextToken) > 0 {
		input.NextToken = aws.String(_cognitoidentityNextToken)
	}

	if resp, err := client.LookupDeveloperIdentity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Merges two users having different IdentityId s, existing in the same identity
// pool, and identified by the same developer provider. You can use this action to
// request that discrete users be merged and identified as a single user in the
// Cognito environment. Cognito associates the given source user (
// SourceUserIdentifier ) with the IdentityId of the DestinationUserIdentifier .
// Only developer-authenticated users can be merged. If the users to be merged are
// associated with the same public provider, but as two different users, an
// exception will be thrown.
//
// The number of linked logins is limited to 20. So, the number of linked logins
// for the source user, SourceUserIdentifier , and the destination user,
// DestinationUserIdentifier , together should not be larger than 20. Otherwise, an
// exception will be thrown.
//
// You must use Amazon Web Services developer credentials to call this operation.
func cognitoidentity_MergeDeveloperIdentities(cfg aws.Config, client *cognitoidentity.Client) {
	input := &cognitoidentity.MergeDeveloperIdentitiesInput{
		// DestinationUserIdentifier: *string, // Required
		// DeveloperProviderName: *string, // Required
		// IdentityPoolId: *string, // Required
		// SourceUserIdentifier: *string, // Required
	}

	if len(_cognitoidentityDestinationUserIdentifier) > 0 {
		input.DestinationUserIdentifier = aws.String(_cognitoidentityDestinationUserIdentifier)
	}
	if len(_cognitoidentityDeveloperProviderName) > 0 {
		input.DeveloperProviderName = aws.String(_cognitoidentityDeveloperProviderName)
	}
	if len(_cognitoidentityIdentityPoolId) > 0 {
		input.IdentityPoolId = aws.String(_cognitoidentityIdentityPoolId)
	}
	if len(_cognitoidentitySourceUserIdentifier) > 0 {
		input.SourceUserIdentifier = aws.String(_cognitoidentitySourceUserIdentifier)
	}

	if resp, err := client.MergeDeveloperIdentities(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the roles for an identity pool. These roles are used when making calls to GetCredentialsForIdentity
// action.
//
// You must use Amazon Web Services developer credentials to call this operation.
func cognitoidentity_SetIdentityPoolRoles(cfg aws.Config, client *cognitoidentity.Client) {
	input := &cognitoidentity.SetIdentityPoolRolesInput{
		// IdentityPoolId: *string, // Required
		// Roles: map[string]string, // Required
	}

	if len(_cognitoidentityIdentityPoolId) > 0 {
		input.IdentityPoolId = aws.String(_cognitoidentityIdentityPoolId)
	}
	if len(_cognitoidentityRoles) > 0 {
		if err := assignInputField(input, "Roles", _cognitoidentityRoles); err != nil {
			log.Errorf("invalid --roles: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityRoleMappings) > 0 {
		if err := assignInputField(input, "RoleMappings", _cognitoidentityRoleMappings); err != nil {
			log.Errorf("invalid --role-mappings: %s", err.Error())
			return
		}
	}

	if resp, err := client.SetIdentityPoolRoles(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// You can use this operation to use default (username and clientID) attribute or
// custom attribute mappings.
func cognitoidentity_SetPrincipalTagAttributeMap(cfg aws.Config, client *cognitoidentity.Client) {
	input := &cognitoidentity.SetPrincipalTagAttributeMapInput{
		// IdentityPoolId: *string, // Required
		// IdentityProviderName: *string, // Required
	}

	if len(_cognitoidentityIdentityPoolId) > 0 {
		input.IdentityPoolId = aws.String(_cognitoidentityIdentityPoolId)
	}
	if len(_cognitoidentityIdentityProviderName) > 0 {
		input.IdentityProviderName = aws.String(_cognitoidentityIdentityProviderName)
	}
	if len(_cognitoidentityPrincipalTags) > 0 {
		if err := assignInputField(input, "PrincipalTags", _cognitoidentityPrincipalTags); err != nil {
			log.Errorf("invalid --principal-tags: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityUseDefaults) > 0 {
		if err := assignInputField(input, "UseDefaults", _cognitoidentityUseDefaults); err != nil {
			log.Errorf("invalid --use-defaults: %s", err.Error())
			return
		}
	}

	if resp, err := client.SetPrincipalTagAttributeMap(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assigns a set of tags to the specified Amazon Cognito identity pool. A tag is a
// label that you can use to categorize and manage identity pools in different
// ways, such as by purpose, owner, environment, or other criteria.
//
// Each tag consists of a key and value, both of which you define. A key is a
// general category for more specific values. For example, if you have two versions
// of an identity pool, one for testing and another for production, you might
// assign an Environment tag key to both identity pools. The value of this key
// might be Test for one identity pool and Production for the other.
//
// Tags are useful for cost tracking and access control. You can activate your
// tags so that they appear on the Billing and Cost Management console, where you
// can track the costs associated with your identity pools. In an IAM policy, you
// can constrain permissions for identity pools based on specific tags or tag
// values.
//
// You can use this action up to 5 times per second, per account. An identity pool
// can have as many as 50 tags.
func cognitoidentity_TagResource(cfg aws.Config, client *cognitoidentity.Client) {
	input := &cognitoidentity.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_cognitoidentityResourceArn) > 0 {
		input.ResourceArn = aws.String(_cognitoidentityResourceArn)
	}
	if len(_cognitoidentityTags) > 0 {
		if err := assignInputField(input, "Tags", _cognitoidentityTags); err != nil {
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

// Unlinks a DeveloperUserIdentifier from an existing identity. Unlinked developer
// users will be considered new identities next time they are seen. If, for a given
// Cognito identity, you remove all federated identities as well as the developer
// user identifier, the Cognito identity becomes inaccessible.
//
// You must use Amazon Web Services developer credentials to call this operation.
func cognitoidentity_UnlinkDeveloperIdentity(cfg aws.Config, client *cognitoidentity.Client) {
	input := &cognitoidentity.UnlinkDeveloperIdentityInput{
		// DeveloperProviderName: *string, // Required
		// DeveloperUserIdentifier: *string, // Required
		// IdentityId: *string, // Required
		// IdentityPoolId: *string, // Required
	}

	if len(_cognitoidentityDeveloperProviderName) > 0 {
		input.DeveloperProviderName = aws.String(_cognitoidentityDeveloperProviderName)
	}
	if len(_cognitoidentityDeveloperUserIdentifier) > 0 {
		input.DeveloperUserIdentifier = aws.String(_cognitoidentityDeveloperUserIdentifier)
	}
	if len(_cognitoidentityIdentityId) > 0 {
		input.IdentityId = aws.String(_cognitoidentityIdentityId)
	}
	if len(_cognitoidentityIdentityPoolId) > 0 {
		input.IdentityPoolId = aws.String(_cognitoidentityIdentityPoolId)
	}

	if resp, err := client.UnlinkDeveloperIdentity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Unlinks a federated identity from an existing account. Unlinked logins will be
// considered new identities next time they are seen. Removing the last linked
// login will make this identity inaccessible.
//
// This is a public API. You do not need any credentials to call this API.
func cognitoidentity_UnlinkIdentity(cfg aws.Config, client *cognitoidentity.Client) {
	input := &cognitoidentity.UnlinkIdentityInput{
		// IdentityId: *string, // Required
		// Logins: map[string]string, // Required
		// LoginsToRemove: []string, // Required
	}

	if len(_cognitoidentityIdentityId) > 0 {
		input.IdentityId = aws.String(_cognitoidentityIdentityId)
	}
	if len(_cognitoidentityLogins) > 0 {
		if err := assignInputField(input, "Logins", _cognitoidentityLogins); err != nil {
			log.Errorf("invalid --logins: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityLoginsToRemove) > 0 {
		input.LoginsToRemove = append([]string(nil), _cognitoidentityLoginsToRemove...)
	}

	if resp, err := client.UnlinkIdentity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified tags from the specified Amazon Cognito identity pool. You
// can use this action up to 5 times per second, per account
func cognitoidentity_UntagResource(cfg aws.Config, client *cognitoidentity.Client) {
	input := &cognitoidentity.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_cognitoidentityResourceArn) > 0 {
		input.ResourceArn = aws.String(_cognitoidentityResourceArn)
	}
	if len(_cognitoidentityTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _cognitoidentityTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration of an identity pool.
// If you don't provide a value for a parameter, Amazon Cognito sets it to its
// default value.
//
// You must use Amazon Web Services developer credentials to call this operation.
func cognitoidentity_UpdateIdentityPool(cfg aws.Config, client *cognitoidentity.Client) {
	input := &cognitoidentity.UpdateIdentityPoolInput{
		// AllowUnauthenticatedIdentities: bool, // Required
		// IdentityPoolId: *string, // Required
		// IdentityPoolName: *string, // Required
	}

	if len(_cognitoidentityAllowUnauthenticatedIdentities) > 0 {
		if err := assignInputField(input, "AllowUnauthenticatedIdentities", _cognitoidentityAllowUnauthenticatedIdentities); err != nil {
			log.Errorf("invalid --allow-unauthenticated-identities: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityIdentityPoolId) > 0 {
		input.IdentityPoolId = aws.String(_cognitoidentityIdentityPoolId)
	}
	if len(_cognitoidentityIdentityPoolName) > 0 {
		input.IdentityPoolName = aws.String(_cognitoidentityIdentityPoolName)
	}
	if len(_cognitoidentityAllowClassicFlow) > 0 {
		if err := assignInputField(input, "AllowClassicFlow", _cognitoidentityAllowClassicFlow); err != nil {
			log.Errorf("invalid --allow-classic-flow: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityCognitoIdentityProviders) > 0 {
		if err := assignInputField(input, "CognitoIdentityProviders", _cognitoidentityCognitoIdentityProviders); err != nil {
			log.Errorf("invalid --cognito-identity-providers: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityDeveloperProviderName) > 0 {
		input.DeveloperProviderName = aws.String(_cognitoidentityDeveloperProviderName)
	}
	if len(_cognitoidentityIdentityPoolTags) > 0 {
		if err := assignInputField(input, "IdentityPoolTags", _cognitoidentityIdentityPoolTags); err != nil {
			log.Errorf("invalid --identity-pool-tags: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityOpenIdConnectProviderARNs) > 0 {
		input.OpenIdConnectProviderARNs = append([]string(nil), _cognitoidentityOpenIdConnectProviderARNs...)
	}
	if len(_cognitoidentitySamlProviderARNs) > 0 {
		input.SamlProviderARNs = append([]string(nil), _cognitoidentitySamlProviderARNs...)
	}
	if len(_cognitoidentitySupportedLoginProviders) > 0 {
		if err := assignInputField(input, "SupportedLoginProviders", _cognitoidentitySupportedLoginProviders); err != nil {
			log.Errorf("invalid --supported-login-providers: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateIdentityPool(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_cognitoidentityCmd)
	_cognitoidentityCmd.Flags().SortFlags = false

	_cognitoidentityCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_cognitoidentityCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_cognitoidentityCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_cognitoidentityCmd.Flags().StringVarP(&_cognitoidentityAccountId, "account-id", "", "", "Account ID")
	_cognitoidentityCmd.Flags().StringVarP(&_cognitoidentityAllowClassicFlow, "allow-classic-flow", "", "", "Allow Classic Flow")
	_cognitoidentityCmd.Flags().StringVarP(&_cognitoidentityAllowUnauthenticatedIdentities, "allow-unauthenticated-identities", "", "", "Allow Unauthenticated Identities")
	_cognitoidentityCmd.Flags().StringVarP(&_cognitoidentityCognitoIdentityProviders, "cognito-identity-providers", "", "", "Cognito Identity Providers")
	_cognitoidentityCmd.Flags().StringVarP(&_cognitoidentityCustomRoleArn, "custom-role-arn", "", "", "Custom Role ARN")
	_cognitoidentityCmd.Flags().StringVarP(&_cognitoidentityDestinationUserIdentifier, "destination-user-identifier", "", "", "Destination User Identifier")
	_cognitoidentityCmd.Flags().StringVarP(&_cognitoidentityDeveloperProviderName, "developer-provider-name", "", "", "Developer Provider Name")
	_cognitoidentityCmd.Flags().StringVarP(&_cognitoidentityDeveloperUserIdentifier, "developer-user-identifier", "", "", "Developer User Identifier")
	_cognitoidentityCmd.Flags().StringVarP(&_cognitoidentityHideDisabled, "hide-disabled", "", "", "Hide Disabled")
	_cognitoidentityCmd.Flags().StringVarP(&_cognitoidentityIdentityId, "identity-id", "", "", "Identity ID")
	_cognitoidentityCmd.Flags().StringSliceVarP(&_cognitoidentityIdentityIdsToDelete, "identity-ids-to-delete", "", nil, "Identity Ids To Delete")
	_cognitoidentityCmd.Flags().StringVarP(&_cognitoidentityIdentityPoolId, "identity-pool-id", "", "", "Identity Pool ID")
	_cognitoidentityCmd.Flags().StringVarP(&_cognitoidentityIdentityPoolName, "identity-pool-name", "", "", "Identity Pool Name")
	_cognitoidentityCmd.Flags().StringVarP(&_cognitoidentityIdentityPoolTags, "identity-pool-tags", "", "", "Identity Pool Tags")
	_cognitoidentityCmd.Flags().StringVarP(&_cognitoidentityIdentityProviderName, "identity-provider-name", "", "", "Identity Provider Name")
	_cognitoidentityCmd.Flags().StringVarP(&_cognitoidentityLogins, "logins", "", "", "Logins")
	_cognitoidentityCmd.Flags().StringSliceVarP(&_cognitoidentityLoginsToRemove, "logins-to-remove", "", nil, "Logins To Remove")
	_cognitoidentityCmd.Flags().StringVarP(&_cognitoidentityMaxResults, "max-results", "", "", "Max Results")
	_cognitoidentityCmd.Flags().StringVarP(&_cognitoidentityNextToken, "next-token", "", "", "Next Token")
	_cognitoidentityCmd.Flags().StringSliceVarP(&_cognitoidentityOpenIdConnectProviderARNs, "open-id-connect-provider-arns", "", nil, "Open ID Connect Provider Arns")
	_cognitoidentityCmd.Flags().StringVarP(&_cognitoidentityPrincipalTags, "principal-tags", "", "", "Principal Tags")
	_cognitoidentityCmd.Flags().StringVarP(&_cognitoidentityResourceArn, "resource-arn", "", "", "Resource ARN")
	_cognitoidentityCmd.Flags().StringVarP(&_cognitoidentityRoleMappings, "role-mappings", "", "", "Role Mappings")
	_cognitoidentityCmd.Flags().StringVarP(&_cognitoidentityRoles, "roles", "", "", "Roles")
	_cognitoidentityCmd.Flags().StringSliceVarP(&_cognitoidentitySamlProviderARNs, "saml-provider-arns", "", nil, "Saml Provider Arns")
	_cognitoidentityCmd.Flags().StringVarP(&_cognitoidentitySourceUserIdentifier, "source-user-identifier", "", "", "Source User Identifier")
	_cognitoidentityCmd.Flags().StringVarP(&_cognitoidentitySupportedLoginProviders, "supported-login-providers", "", "", "Supported Login Providers")
	_cognitoidentityCmd.Flags().StringSliceVarP(&_cognitoidentityTagKeys, "tag-keys", "", nil, "Tag Keys")
	_cognitoidentityCmd.Flags().StringVarP(&_cognitoidentityTags, "tags", "", "", "Tags")
	_cognitoidentityCmd.Flags().StringVarP(&_cognitoidentityTokenDuration, "token-duration", "", "", "Token Duration")
	_cognitoidentityCmd.Flags().StringVarP(&_cognitoidentityUseDefaults, "use-defaults", "", "", "Use Defaults")

	_cognitoidentityCmd.Flags().BoolVarP(&_cognitoidentityCreateIdentityPool, "create-identity-pool", "", false, "Create Identity Pool")
	_cognitoidentityCmd.Flags().BoolVarP(&_cognitoidentityDeleteIdentities, "delete-identities", "", false, "Delete Identities")
	_cognitoidentityCmd.Flags().BoolVarP(&_cognitoidentityDeleteIdentityPool, "delete-identity-pool", "", false, "Delete Identity Pool")
	_cognitoidentityCmd.Flags().BoolVarP(&_cognitoidentityDescribeIdentity, "describe-identity", "", false, "Describe Identity")
	_cognitoidentityCmd.Flags().BoolVarP(&_cognitoidentityDescribeIdentityPool, "describe-identity-pool", "", false, "Describe Identity Pool")
	_cognitoidentityCmd.Flags().BoolVarP(&_cognitoidentityGetCredentialsForIdentity, "get-credentials-for-identity", "", false, "Get Credentials For Identity")
	_cognitoidentityCmd.Flags().BoolVarP(&_cognitoidentityGetId, "get-id", "", false, "Get ID")
	_cognitoidentityCmd.Flags().BoolVarP(&_cognitoidentityGetIdentityPoolRoles, "get-identity-pool-roles", "", false, "Get Identity Pool Roles")
	_cognitoidentityCmd.Flags().BoolVarP(&_cognitoidentityGetOpenIdToken, "get-open-id-token", "", false, "Get Open ID Token")
	_cognitoidentityCmd.Flags().BoolVarP(&_cognitoidentityGetOpenIdTokenForDeveloperIdentity, "get-open-id-token-for-developer-identity", "", false, "Get Open ID Token For Developer Identity")
	_cognitoidentityCmd.Flags().BoolVarP(&_cognitoidentityGetPrincipalTagAttributeMap, "get-principal-tag-attribute-map", "", false, "Get Principal Tag Attribute Map")
	_cognitoidentityCmd.Flags().BoolVarP(&_cognitoidentityListIdentities, "list-identities", "", false, "List Identities")
	_cognitoidentityCmd.Flags().BoolVarP(&_cognitoidentityListIdentityPools, "list-identity-pools", "", false, "List Identity Pools")
	_cognitoidentityCmd.Flags().BoolVarP(&_cognitoidentityListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_cognitoidentityCmd.Flags().BoolVarP(&_cognitoidentityLookupDeveloperIdentity, "lookup-developer-identity", "", false, "Lookup Developer Identity")
	_cognitoidentityCmd.Flags().BoolVarP(&_cognitoidentityMergeDeveloperIdentities, "merge-developer-identities", "", false, "Merge Developer Identities")
	_cognitoidentityCmd.Flags().BoolVarP(&_cognitoidentitySetIdentityPoolRoles, "set-identity-pool-roles", "", false, "Set Identity Pool Roles")
	_cognitoidentityCmd.Flags().BoolVarP(&_cognitoidentitySetPrincipalTagAttributeMap, "set-principal-tag-attribute-map", "", false, "Set Principal Tag Attribute Map")
	_cognitoidentityCmd.Flags().BoolVarP(&_cognitoidentityTagResource, "tag-resource", "", false, "Tag Resource")
	_cognitoidentityCmd.Flags().BoolVarP(&_cognitoidentityUnlinkDeveloperIdentity, "unlink-developer-identity", "", false, "Unlink Developer Identity")
	_cognitoidentityCmd.Flags().BoolVarP(&_cognitoidentityUnlinkIdentity, "unlink-identity", "", false, "Unlink Identity")
	_cognitoidentityCmd.Flags().BoolVarP(&_cognitoidentityUntagResource, "untag-resource", "", false, "Untag Resource")
	_cognitoidentityCmd.Flags().BoolVarP(&_cognitoidentityUpdateIdentityPool, "update-identity-pool", "", false, "Update Identity Pool")

}
