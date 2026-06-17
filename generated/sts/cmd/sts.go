package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// stsCmd represents the sts command
var _stsCmd = &cobra.Command{
	Use:   "sts",
	Short: "AWS sts CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := sts.NewFromConfig(cfg)
		if _stsAssumeRole {
			sts_AssumeRole(cfg, client)
			return
		}
		if _stsAssumeRoleWithSAML {
			sts_AssumeRoleWithSAML(cfg, client)
			return
		}
		if _stsAssumeRoleWithWebIdentity {
			sts_AssumeRoleWithWebIdentity(cfg, client)
			return
		}
		if _stsAssumeRoot {
			sts_AssumeRoot(cfg, client)
			return
		}
		if _stsDecodeAuthorizationMessage {
			sts_DecodeAuthorizationMessage(cfg, client)
			return
		}
		if _stsGetAccessKeyInfo {
			sts_GetAccessKeyInfo(cfg, client)
			return
		}
		if _stsGetCallerIdentity {
			sts_GetCallerIdentity(cfg, client)
			return
		}
		if _stsGetDelegatedAccessToken {
			sts_GetDelegatedAccessToken(cfg, client)
			return
		}
		if _stsGetFederationToken {
			sts_GetFederationToken(cfg, client)
			return
		}
		if _stsGetSessionToken {
			sts_GetSessionToken(cfg, client)
			return
		}
		if _stsGetWebIdentityToken {
			sts_GetWebIdentityToken(cfg, client)
			return
		}

	},
}

var (
	_stsAssumeRole                 bool
	_stsAssumeRoleWithSAML         bool
	_stsAssumeRoleWithWebIdentity  bool
	_stsAssumeRoot                 bool
	_stsDecodeAuthorizationMessage bool
	_stsGetAccessKeyInfo           bool
	_stsGetCallerIdentity          bool
	_stsGetDelegatedAccessToken    bool
	_stsGetFederationToken         bool
	_stsGetSessionToken            bool
	_stsGetWebIdentityToken        bool

	_stsAccessKeyId       string
	_stsAudience          []string
	_stsDurationSeconds   string
	_stsEncodedMessage    string
	_stsExternalId        string
	_stsName              string
	_stsPolicy            string
	_stsPolicyArns        string
	_stsPrincipalArn      string
	_stsProvidedContexts  string
	_stsProviderId        string
	_stsRoleArn           string
	_stsRoleSessionName   string
	_stsSAMLAssertion     string
	_stsSerialNumber      string
	_stsSigningAlgorithm  string
	_stsSourceIdentity    string
	_stsTags              string
	_stsTargetPrincipal   string
	_stsTaskPolicyArn     string
	_stsTokenCode         string
	_stsTradeInToken      string
	_stsTransitiveTagKeys []string
	_stsWebIdentityToken  string
)

// Returns a set of temporary security credentials that you can use to access
// Amazon Web Services resources. These temporary credentials consist of an access
// key ID, a secret access key, and a security token. Typically, you use AssumeRole
// within your account or for cross-account access. For a comparison of AssumeRole
// with other API operations that produce temporary credentials, see [Requesting Temporary Security Credentials]and [Compare STS credentials] in the
// IAM User Guide.
//
// # Permissions
//
// The temporary security credentials created by AssumeRole can be used to make
// API calls to any Amazon Web Services service with the following exception: You
// cannot call the Amazon Web Services STS GetFederationToken or GetSessionToken
// API operations.
//
// (Optional) You can pass inline or managed session policies to this operation.
// You can pass a single JSON policy document to use as an inline session policy.
// You can also specify up to 10 managed policy Amazon Resource Names (ARNs) to use
// as managed session policies. The plaintext that you use for both inline and
// managed session policies can't exceed 2,048 characters. Passing policies to this
// operation returns new temporary credentials. The resulting session's permissions
// are the intersection of the role's identity-based policy and the session
// policies. You can use the role's temporary credentials in subsequent Amazon Web
// Services API calls to access resources in the account that owns the role. You
// cannot use session policies to grant more permissions than those allowed by the
// identity-based policy of the role that is being assumed. For more information,
// see [Session Policies]in the IAM User Guide.
//
// When you create a role, you create two policies: a role trust policy that
// specifies who can assume the role, and a permissions policy that specifies what
// can be done with the role. You specify the trusted principal that is allowed to
// assume the role in the role trust policy.
//
// To assume a role from a different account, your Amazon Web Services account
// must be trusted by the role. The trust relationship is defined in the role's
// trust policy when the role is created. That trust policy states which accounts
// are allowed to delegate that access to users in the account.
//
// A user who wants to access a role in a different account must also have
// permissions that are delegated from the account administrator. The administrator
// must attach a policy that allows the user to call AssumeRole for the ARN of the
// role in the other account.
//
// To allow a user to assume a role in the same account, you can do either of the
// following:
//
// - Attach a policy to the user that allows the user to call AssumeRole (as long
// as the role's trust policy trusts the account).
//
// - Add the user as a principal directly in the role's trust policy.
//
// You can do either because the role’s trust policy acts as an IAM resource-based
// policy. When a resource-based policy grants access to a principal in the same
// account, no additional identity-based policy is required. For more information
// about trust policies and resource-based policies, see [IAM Policies]in the IAM User Guide.
//
// # Tags
//
// (Optional) You can pass tag key-value pairs to your session. These tags are
// called session tags. For more information about session tags, see [Passing Session Tags in STS]in the IAM
// User Guide.
//
// An administrator must grant you the permissions necessary to pass session tags.
// The administrator can also create granular permissions to allow you to pass only
// specific session tags. For more information, see [Tutorial: Using Tags for Attribute-Based Access Control]in the IAM User Guide.
//
// You can set the session tags as transitive. Transitive tags persist during role
// chaining. For more information, see [Chaining Roles with Session Tags]in the IAM User Guide.
//
// # Using MFA with AssumeRole
//
// (Optional) You can include multi-factor authentication (MFA) information when
// you call AssumeRole . This is useful for cross-account scenarios to ensure that
// the user that assumes the role has been authenticated with an Amazon Web
// Services MFA device. In that scenario, the trust policy of the role being
// assumed includes a condition that tests for MFA authentication. If the caller
// does not include valid MFA information, the request to assume the role is
// denied. The condition in a trust policy that tests for MFA authentication might
// look like the following example.
//
// "Condition": {"Bool": {"aws:MultiFactorAuthPresent": true}}
//
// For more information, see [Configuring MFA-Protected API Access] in the IAM User Guide guide.
//
// To use MFA with AssumeRole , you pass values for the SerialNumber and TokenCode
// parameters. The SerialNumber value identifies the user's hardware or virtual
// MFA device. The TokenCode is the time-based one-time password (TOTP) that the
// MFA device produces.
//
// [Configuring MFA-Protected API Access]: https://docs.aws.amazon.com/IAM/latest/UserGuide/MFAProtectedAPI.html
// [Session Policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#policies_session
// [Passing Session Tags in STS]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_session-tags.html
// [Chaining Roles with Session Tags]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_session-tags.html#id_session-tags_role-chaining
// [IAM Policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html
// [Requesting Temporary Security Credentials]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html
// [Compare STS credentials]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_sts-comparison.html
// [Tutorial: Using Tags for Attribute-Based Access Control]: https://docs.aws.amazon.com/IAM/latest/UserGuide/tutorial_attribute-based-access-control.html
func sts_AssumeRole(cfg aws.Config, client *sts.Client) {
	input := &sts.AssumeRoleInput{
		// RoleArn: *string, // Required
		// RoleSessionName: *string, // Required
	}

	if len(_stsRoleArn) > 0 {
		input.RoleArn = aws.String(_stsRoleArn)
	}
	if len(_stsRoleSessionName) > 0 {
		input.RoleSessionName = aws.String(_stsRoleSessionName)
	}
	if len(_stsDurationSeconds) > 0 {
		if err := assignInputField(input, "DurationSeconds", _stsDurationSeconds); err != nil {
			log.Errorf("invalid --duration-seconds: %s", err.Error())
			return
		}
	}
	if len(_stsExternalId) > 0 {
		input.ExternalId = aws.String(_stsExternalId)
	}
	if len(_stsPolicy) > 0 {
		input.Policy = aws.String(_stsPolicy)
	}
	if len(_stsPolicyArns) > 0 {
		if err := assignInputField(input, "PolicyArns", _stsPolicyArns); err != nil {
			log.Errorf("invalid --policy-arns: %s", err.Error())
			return
		}
	}
	if len(_stsProvidedContexts) > 0 {
		if err := assignInputField(input, "ProvidedContexts", _stsProvidedContexts); err != nil {
			log.Errorf("invalid --provided-contexts: %s", err.Error())
			return
		}
	}
	if len(_stsSerialNumber) > 0 {
		input.SerialNumber = aws.String(_stsSerialNumber)
	}
	if len(_stsSourceIdentity) > 0 {
		input.SourceIdentity = aws.String(_stsSourceIdentity)
	}
	if len(_stsTags) > 0 {
		if err := assignInputField(input, "Tags", _stsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_stsTokenCode) > 0 {
		input.TokenCode = aws.String(_stsTokenCode)
	}
	if len(_stsTransitiveTagKeys) > 0 {
		input.TransitiveTagKeys = append([]string(nil), _stsTransitiveTagKeys...)
	}

	if resp, err := client.AssumeRole(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a set of temporary security credentials for users who have been
// authenticated via a SAML authentication response. This operation provides a
// mechanism for tying an enterprise identity store or directory to role-based
// Amazon Web Services access without user-specific credentials or configuration.
// For a comparison of AssumeRoleWithSAML with the other API operations that
// produce temporary credentials, see [Requesting Temporary Security Credentials]and [Compare STS credentials] in the IAM User Guide.
//
// The temporary security credentials returned by this operation consist of an
// access key ID, a secret access key, and a security token. Applications can use
// these temporary security credentials to sign calls to Amazon Web Services
// services.
//
// AssumeRoleWithSAML will not work on IAM Identity Center managed roles. These
// roles' names start with AWSReservedSSO_ .
//
// # Session Duration
//
// By default, the temporary security credentials created by AssumeRoleWithSAML
// last for one hour. However, you can use the optional DurationSeconds parameter
// to specify the duration of your session. Your role session lasts for the
// duration that you specify, or until the time specified in the SAML
// authentication response's SessionNotOnOrAfter value, whichever is shorter. You
// can provide a DurationSeconds value from 900 seconds (15 minutes) up to the
// maximum session duration setting for the role. This setting can have a value
// from 1 hour to 12 hours. To learn how to view the maximum value for your role,
// see [View the Maximum Session Duration Setting for a Role]in the IAM User Guide. The maximum session duration limit applies when you
// use the AssumeRole* API operations or the assume-role* CLI commands. However
// the limit does not apply when you use those operations to create a console URL.
// For more information, see [Using IAM Roles]in the IAM User Guide.
//
// [Role chaining]limits your CLI or Amazon Web Services API role session to a maximum of one
// hour. When you use the AssumeRole API operation to assume a role, you can
// specify the duration of your role session with the DurationSeconds parameter.
// You can specify a parameter value of up to 43200 seconds (12 hours), depending
// on the maximum session duration setting for your role. However, if you assume a
// role using role chaining and provide a DurationSeconds parameter value greater
// than one hour, the operation fails.
//
// # Permissions
//
// The temporary security credentials created by AssumeRoleWithSAML can be used to
// make API calls to any Amazon Web Services service with the following exception:
// you cannot call the STS GetFederationToken or GetSessionToken API operations.
//
// (Optional) You can pass inline or managed [session policies] to this operation. You can pass a
// single JSON policy document to use as an inline session policy. You can also
// specify up to 10 managed policy Amazon Resource Names (ARNs) to use as managed
// session policies. The plaintext that you use for both inline and managed session
// policies can't exceed 2,048 characters. Passing policies to this operation
// returns new temporary credentials. The resulting session's permissions are the
// intersection of the role's identity-based policy and the session policies. You
// can use the role's temporary credentials in subsequent Amazon Web Services API
// calls to access resources in the account that owns the role. You cannot use
// session policies to grant more permissions than those allowed by the
// identity-based policy of the role that is being assumed. For more information,
// see [Session Policies]in the IAM User Guide.
//
// Calling AssumeRoleWithSAML does not require the use of Amazon Web Services
// security credentials. The identity of the caller is validated by using keys in
// the metadata document that is uploaded for the SAML provider entity for your
// identity provider.
//
// Calling AssumeRoleWithSAML can result in an entry in your CloudTrail logs. The
// entry includes the value in the NameID element of the SAML assertion. We
// recommend that you use a NameIDType that is not associated with any personally
// identifiable information (PII). For example, you could instead use the
// persistent identifier ( urn:oasis:names:tc:SAML:2.0:nameid-format:persistent ).
//
// # Tags
//
// (Optional) You can configure your IdP to pass attributes into your SAML
// assertion as session tags. Each session tag consists of a key name and an
// associated value. For more information about session tags, see [Passing Session Tags in STS]in the IAM User
// Guide.
//
// You can pass up to 50 session tags. The plaintext session tag keys can’t exceed
// 128 characters and the values can’t exceed 256 characters. For these and
// additional limits, see [IAM and STS Character Limits]in the IAM User Guide.
//
// An Amazon Web Services conversion compresses the passed inline session policy,
// managed policy ARNs, and session tags into a packed binary format that has a
// separate limit. Your request can fail for this limit even if your plaintext
// meets the other requirements. The PackedPolicySize response element indicates
// by percentage how close the policies and tags for your request are to the upper
// size limit.
//
// You can pass a session tag with the same key as a tag that is attached to the
// role. When you do, session tags override the role's tags with the same key.
//
// An administrator must grant you the permissions necessary to pass session tags.
// The administrator can also create granular permissions to allow you to pass only
// specific session tags. For more information, see [Tutorial: Using Tags for Attribute-Based Access Control]in the IAM User Guide.
//
// You can set the session tags as transitive. Transitive tags persist during role
// chaining. For more information, see [Chaining Roles with Session Tags]in the IAM User Guide.
//
// # SAML Configuration
//
// Before your application can call AssumeRoleWithSAML , you must configure your
// SAML identity provider (IdP) to issue the claims required by Amazon Web
// Services. Additionally, you must use Identity and Access Management (IAM) to
// create a SAML provider entity in your Amazon Web Services account that
// represents your identity provider. You must also create an IAM role that
// specifies this SAML provider in its trust policy.
//
// For more information, see the following resources:
//
// [About SAML 2.0-based Federation]
// - in the IAM User Guide.
//
// [Creating SAML Identity Providers]
// - in the IAM User Guide.
//
// [Configuring a Relying Party and Claims]
// - in the IAM User Guide.
//
// [Creating a Role for SAML 2.0 Federation]
// - in the IAM User Guide.
//
// [View the Maximum Session Duration Setting for a Role]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use.html#id_roles_use_view-role-max-session
// [Creating a Role for SAML 2.0 Federation]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-idp_saml.html
// [IAM and STS Character Limits]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-limits.html#reference_iam-limits-entity-length
// [Creating SAML Identity Providers]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_create_saml.html
// [session policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#policies_session
// [Requesting Temporary Security Credentials]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html
// [Compare STS credentials]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_sts-comparison.html
// [Tutorial: Using Tags for Attribute-Based Access Control]: https://docs.aws.amazon.com/IAM/latest/UserGuide/tutorial_attribute-based-access-control.html
// [Configuring a Relying Party and Claims]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_create_saml_relying-party.html
// [Role chaining]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_terms-and-concepts.html#iam-term-role-chaining
// [Using IAM Roles]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use.html
// [Session Policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#policies_session
// [Passing Session Tags in STS]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_session-tags.html
// [About SAML 2.0-based Federation]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_saml.html
// [Chaining Roles with Session Tags]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_session-tags.html#id_session-tags_role-chaining
func sts_AssumeRoleWithSAML(cfg aws.Config, client *sts.Client) {
	input := &sts.AssumeRoleWithSAMLInput{
		// PrincipalArn: *string, // Required
		// RoleArn: *string, // Required
		// SAMLAssertion: *string, // Required
	}

	if len(_stsPrincipalArn) > 0 {
		input.PrincipalArn = aws.String(_stsPrincipalArn)
	}
	if len(_stsRoleArn) > 0 {
		input.RoleArn = aws.String(_stsRoleArn)
	}
	if len(_stsSAMLAssertion) > 0 {
		input.SAMLAssertion = aws.String(_stsSAMLAssertion)
	}
	if len(_stsDurationSeconds) > 0 {
		if err := assignInputField(input, "DurationSeconds", _stsDurationSeconds); err != nil {
			log.Errorf("invalid --duration-seconds: %s", err.Error())
			return
		}
	}
	if len(_stsPolicy) > 0 {
		input.Policy = aws.String(_stsPolicy)
	}
	if len(_stsPolicyArns) > 0 {
		if err := assignInputField(input, "PolicyArns", _stsPolicyArns); err != nil {
			log.Errorf("invalid --policy-arns: %s", err.Error())
			return
		}
	}

	if resp, err := client.AssumeRoleWithSAML(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a set of temporary security credentials for users who have been
// authenticated in a mobile or web application with a web identity provider.
// Example providers include the OAuth 2.0 providers Login with Amazon and
// Facebook, or any OpenID Connect-compatible identity provider such as Google or [Amazon Cognito federated identities].
//
// For mobile applications, we recommend that you use Amazon Cognito. You can use
// Amazon Cognito with the [Amazon Web Services SDK for iOS Developer Guide]and the [Amazon Web Services SDK for Android Developer Guide] to uniquely identify a user. You can also
// supply the user with a consistent identity throughout the lifetime of an
// application.
//
// To learn more about Amazon Cognito, see [Amazon Cognito identity pools] in Amazon Cognito Developer Guide.
//
// Calling AssumeRoleWithWebIdentity does not require the use of Amazon Web
// Services security credentials. Therefore, you can distribute an application (for
// example, on mobile devices) that requests temporary security credentials without
// including long-term Amazon Web Services credentials in the application. You also
// don't need to deploy server-based proxy services that use long-term Amazon Web
// Services credentials. Instead, the identity of the caller is validated by using
// a token from the web identity provider. For a comparison of
// AssumeRoleWithWebIdentity with the other API operations that produce temporary
// credentials, see [Requesting Temporary Security Credentials]and [Compare STS credentials] in the IAM User Guide.
//
// The temporary security credentials returned by this API consist of an access
// key ID, a secret access key, and a security token. Applications can use these
// temporary security credentials to sign calls to Amazon Web Services service API
// operations.
//
// # Session Duration
//
// By default, the temporary security credentials created by
// AssumeRoleWithWebIdentity last for one hour. However, you can use the optional
// DurationSeconds parameter to specify the duration of your session. You can
// provide a value from 900 seconds (15 minutes) up to the maximum session duration
// setting for the role. This setting can have a value from 1 hour to 12 hours. To
// learn how to view the maximum value for your role, see [Update the maximum session duration for a role]in the IAM User Guide.
// The maximum session duration limit applies when you use the AssumeRole* API
// operations or the assume-role* CLI commands. However the limit does not apply
// when you use those operations to create a console URL. For more information, see
// [Using IAM Roles]in the IAM User Guide.
//
// # Permissions
//
// The temporary security credentials created by AssumeRoleWithWebIdentity can be
// used to make API calls to any Amazon Web Services service with the following
// exception: you cannot call the STS GetFederationToken or GetSessionToken API
// operations.
//
// (Optional) You can pass inline or managed [session policies] to this operation. You can pass a
// single JSON policy document to use as an inline session policy. You can also
// specify up to 10 managed policy Amazon Resource Names (ARNs) to use as managed
// session policies. The plaintext that you use for both inline and managed session
// policies can't exceed 2,048 characters. Passing policies to this operation
// returns new temporary credentials. The resulting session's permissions are the
// intersection of the role's identity-based policy and the session policies. You
// can use the role's temporary credentials in subsequent Amazon Web Services API
// calls to access resources in the account that owns the role. You cannot use
// session policies to grant more permissions than those allowed by the
// identity-based policy of the role that is being assumed. For more information,
// see [Session Policies]in the IAM User Guide.
//
// # Tags
//
// (Optional) You can configure your IdP to pass attributes into your web identity
// token as session tags. Each session tag consists of a key name and an associated
// value. For more information about session tags, see [Passing session tags using AssumeRoleWithWebIdentity]in the IAM User Guide.
//
// You can pass up to 50 session tags. The plaintext session tag keys can’t exceed
// 128 characters and the values can’t exceed 256 characters. For these and
// additional limits, see [IAM and STS Character Limits]in the IAM User Guide.
//
// An Amazon Web Services conversion compresses the passed inline session policy,
// managed policy ARNs, and session tags into a packed binary format that has a
// separate limit. Your request can fail for this limit even if your plaintext
// meets the other requirements. The PackedPolicySize response element indicates
// by percentage how close the policies and tags for your request are to the upper
// size limit.
//
// You can pass a session tag with the same key as a tag that is attached to the
// role. When you do, the session tag overrides the role tag with the same key.
//
// An administrator must grant you the permissions necessary to pass session tags.
// The administrator can also create granular permissions to allow you to pass only
// specific session tags. For more information, see [Tutorial: Using Tags for Attribute-Based Access Control]in the IAM User Guide.
//
// You can set the session tags as transitive. Transitive tags persist during role
// chaining. For more information, see [Chaining Roles with Session Tags]in the IAM User Guide.
//
// # Identities
//
// Before your application can call AssumeRoleWithWebIdentity , you must have an
// identity token from a supported identity provider and create a role that the
// application can assume. The role that your application assumes must trust the
// identity provider that is associated with the identity token. In other words,
// the identity provider must be specified in the role's trust policy.
//
// Calling AssumeRoleWithWebIdentity can result in an entry in your CloudTrail
// logs. The entry includes the [Subject]of the provided web identity token. We recommend
// that you avoid using any personally identifiable information (PII) in this
// field. For example, you could instead use a GUID or a pairwise identifier, as [suggested in the OIDC specification].
//
// For more information about how to use OIDC federation and the
// AssumeRoleWithWebIdentity API, see the following resources:
//
// [Using Web Identity Federation API Operations for Mobile Apps]
// - and [Federation Through a Web-based Identity Provider].
//
// [Amazon Web Services SDK for iOS Developer Guide]
// - and [Amazon Web Services SDK for Android Developer Guide]. These toolkits contain sample apps that show how to invoke the
// identity providers. The toolkits then show how to use the information from these
// providers to get and use temporary security credentials.
//
// [Amazon Web Services SDK for iOS Developer Guide]: http://aws.amazon.com/sdkforios/
// [Passing session tags using AssumeRoleWithWebIdentity]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_session-tags.html#id_session-tags_adding-assume-role-idp
// [Amazon Web Services SDK for Android Developer Guide]: http://aws.amazon.com/sdkforandroid/
// [IAM and STS Character Limits]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-limits.html#reference_iam-limits-entity-length
// [session policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#policies_session
// [Requesting Temporary Security Credentials]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html
// [Compare STS credentials]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_sts-comparison.html
// [Subject]: http://openid.net/specs/openid-connect-core-1_0.html#Claims
// [Tutorial: Using Tags for Attribute-Based Access Control]: https://docs.aws.amazon.com/IAM/latest/UserGuide/tutorial_attribute-based-access-control.html
// [Amazon Cognito identity pools]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-identity.html
// [Federation Through a Web-based Identity Provider]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html#api_assumerolewithwebidentity
// [Using IAM Roles]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use.html
// [Session Policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#policies_session
// [Amazon Cognito federated identities]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-identity.html
// [Chaining Roles with Session Tags]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_session-tags.html#id_session-tags_role-chaining
// [Update the maximum session duration for a role]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_update-role-settings.html#id_roles_update-session-duration
// [Using Web Identity Federation API Operations for Mobile Apps]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_oidc_manual.html
// [suggested in the OIDC specification]: http://openid.net/specs/openid-connect-core-1_0.html#SubjectIDTypes
func sts_AssumeRoleWithWebIdentity(cfg aws.Config, client *sts.Client) {
	input := &sts.AssumeRoleWithWebIdentityInput{
		// RoleArn: *string, // Required
		// RoleSessionName: *string, // Required
		// WebIdentityToken: *string, // Required
	}

	if len(_stsRoleArn) > 0 {
		input.RoleArn = aws.String(_stsRoleArn)
	}
	if len(_stsRoleSessionName) > 0 {
		input.RoleSessionName = aws.String(_stsRoleSessionName)
	}
	if len(_stsWebIdentityToken) > 0 {
		input.WebIdentityToken = aws.String(_stsWebIdentityToken)
	}
	if len(_stsDurationSeconds) > 0 {
		if err := assignInputField(input, "DurationSeconds", _stsDurationSeconds); err != nil {
			log.Errorf("invalid --duration-seconds: %s", err.Error())
			return
		}
	}
	if len(_stsPolicy) > 0 {
		input.Policy = aws.String(_stsPolicy)
	}
	if len(_stsPolicyArns) > 0 {
		if err := assignInputField(input, "PolicyArns", _stsPolicyArns); err != nil {
			log.Errorf("invalid --policy-arns: %s", err.Error())
			return
		}
	}
	if len(_stsProviderId) > 0 {
		input.ProviderId = aws.String(_stsProviderId)
	}

	if resp, err := client.AssumeRoleWithWebIdentity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a set of short term credentials you can use to perform privileged tasks
// on a member account in your organization. You must use credentials from an
// Organizations management account or a delegated administrator account for IAM to
// call AssumeRoot . You cannot use root user credentials to make this call.
//
// Before you can launch a privileged session, you must have centralized root
// access in your organization. For steps to enable this feature, see [Centralize root access for member accounts]in the IAM
// User Guide.
//
// The STS global endpoint is not supported for AssumeRoot. You must send this
// request to a Regional STS endpoint. For more information, see [Endpoints].
//
// You can track AssumeRoot in CloudTrail logs to determine what actions were
// performed in a session. For more information, see [Track privileged tasks in CloudTrail]in the IAM User Guide.
//
// When granting access to privileged tasks you should only grant the necessary
// permissions required to perform that task. For more information, see [Security best practices in IAM]. In
// addition, you can use [service control policies](SCPs) to manage and limit permissions in your
// organization. See [General examples]in the Organizations User Guide for more information on SCPs.
//
// [Endpoints]: https://docs.aws.amazon.com/STS/latest/APIReference/welcome.html#sts-endpoints
// [Security best practices in IAM]: https://docs.aws.amazon.com/IAM/latest/UserGuide/best-practices.html
// [Track privileged tasks in CloudTrail]: https://docs.aws.amazon.com/IAM/latest/UserGuide/cloudtrail-track-privileged-tasks.html
// [General examples]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_scps_examples_general.html
// [service control policies]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_scps.html
// [Centralize root access for member accounts]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_root-enable-root-access.html
func sts_AssumeRoot(cfg aws.Config, client *sts.Client) {
	input := &sts.AssumeRootInput{
		// TargetPrincipal: *string, // Required
		// TaskPolicyArn: *types.PolicyDescriptorType, // Required
	}

	if len(_stsTargetPrincipal) > 0 {
		input.TargetPrincipal = aws.String(_stsTargetPrincipal)
	}
	if len(_stsTaskPolicyArn) > 0 {
		if err := assignInputField(input, "TaskPolicyArn", _stsTaskPolicyArn); err != nil {
			log.Errorf("invalid --task-policy-arn: %s", err.Error())
			return
		}
	}
	if len(_stsDurationSeconds) > 0 {
		if err := assignInputField(input, "DurationSeconds", _stsDurationSeconds); err != nil {
			log.Errorf("invalid --duration-seconds: %s", err.Error())
			return
		}
	}

	if resp, err := client.AssumeRoot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Decodes additional information about the authorization status of a request from
// an encoded message returned in response to an Amazon Web Services request.
//
// For example, if a user is not authorized to perform an operation that he or she
// has requested, the request returns a Client.UnauthorizedOperation response (an
// HTTP 403 response). Some Amazon Web Services operations additionally return an
// encoded message that can provide details about this authorization failure.
//
// Only certain Amazon Web Services operations return an encoded authorization
// message. The documentation for an individual operation indicates whether that
// operation returns an encoded message in addition to returning an HTTP code.
//
// The message is encoded because the details of the authorization status can
// contain privileged information that the user who requested the operation should
// not see. To decode an authorization status message, a user must be granted
// permissions through an IAM [policy]to request the DecodeAuthorizationMessage (
// sts:DecodeAuthorizationMessage ) action.
//
// The decoded message includes the following type of information:
//
// - Whether the request was denied due to an explicit deny or due to the
// absence of an explicit allow. For more information, see [Determining Whether a Request is Allowed or Denied]in the IAM User
// Guide.
//
// - The principal who made the request.
//
// - The requested action.
//
// - The requested resource.
//
// - The values of condition keys in the context of the user's request.
//
// [Determining Whether a Request is Allowed or Denied]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html#policy-eval-denyallow
// [policy]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html
func sts_DecodeAuthorizationMessage(cfg aws.Config, client *sts.Client) {
	input := &sts.DecodeAuthorizationMessageInput{
		// EncodedMessage: *string, // Required
	}

	if len(_stsEncodedMessage) > 0 {
		input.EncodedMessage = aws.String(_stsEncodedMessage)
	}

	if resp, err := client.DecodeAuthorizationMessage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the account identifier for the specified access key ID.
// Access keys consist of two parts: an access key ID (for example,
// AKIAIOSFODNN7EXAMPLE ) and a secret access key (for example,
// wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY ). For more information about access
// keys, see [Managing Access Keys for IAM Users]in the IAM User Guide.
//
// When you pass an access key ID to this operation, it returns the ID of the
// Amazon Web Services account to which the keys belong. Access key IDs beginning
// with AKIA are long-term credentials for an IAM user or the Amazon Web Services
// account root user. Access key IDs beginning with ASIA are temporary credentials
// that are created using STS operations. If the account in the response belongs to
// you, you can sign in as the root user and review your root user access keys.
// Then, you can pull a [credentials report]to learn which IAM user owns the keys. To learn who
// requested the temporary credentials for an ASIA access key, view the STS events
// in your [CloudTrail logs]in the IAM User Guide.
//
// This operation does not indicate the state of the access key. The key might be
// active, inactive, or deleted. Active keys might not have permissions to perform
// an operation. Providing a deleted access key might return an error that the key
// doesn't exist.
//
// [credentials report]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_getting-report.html
// [CloudTrail logs]: https://docs.aws.amazon.com/IAM/latest/UserGuide/cloudtrail-integration.html
// [Managing Access Keys for IAM Users]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_access-keys.html
func sts_GetAccessKeyInfo(cfg aws.Config, client *sts.Client) {
	input := &sts.GetAccessKeyInfoInput{
		// AccessKeyId: *string, // Required
	}

	if len(_stsAccessKeyId) > 0 {
		input.AccessKeyId = aws.String(_stsAccessKeyId)
	}

	if resp, err := client.GetAccessKeyInfo(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns details about the IAM user or role whose credentials are used to call
// the operation.
//
// No permissions are required to perform this operation. If an administrator
// attaches a policy to your identity that explicitly denies access to the
// sts:GetCallerIdentity action, you can still perform this operation. Permissions
// are not required because the same information is returned when access is denied.
// To view an example response, see [I Am Not Authorized to Perform: iam:DeleteVirtualMFADevice]in the IAM User Guide.
//
// [I Am Not Authorized to Perform: iam:DeleteVirtualMFADevice]: https://docs.aws.amazon.com/IAM/latest/UserGuide/troubleshoot_general.html#troubleshoot_general_access-denied-delete-mfa
func sts_GetCallerIdentity(cfg aws.Config, client *sts.Client) {
	input := &sts.GetCallerIdentityInput{}

	if resp, err := client.GetCallerIdentity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Exchanges a trade-in token for temporary Amazon Web Services credentials with
// the permissions associated with the assumed principal. This operation allows you
// to obtain credentials for a specific principal based on a trade-in token,
// enabling delegation of access to Amazon Web Services resources.
func sts_GetDelegatedAccessToken(cfg aws.Config, client *sts.Client) {
	input := &sts.GetDelegatedAccessTokenInput{
		// TradeInToken: *string, // Required
	}

	if len(_stsTradeInToken) > 0 {
		input.TradeInToken = aws.String(_stsTradeInToken)
	}

	if resp, err := client.GetDelegatedAccessToken(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a set of temporary security credentials (consisting of an access key
// ID, a secret access key, and a security token) for a user. A typical use is in a
// proxy application that gets temporary security credentials on behalf of
// distributed applications inside a corporate network.
//
// You must call the GetFederationToken operation using the long-term security
// credentials of an IAM user. As a result, this call is appropriate in contexts
// where those credentials can be safeguarded, usually in a server-based
// application. For a comparison of GetFederationToken with the other API
// operations that produce temporary credentials, see [Requesting Temporary Security Credentials]and [Compare STS credentials] in the IAM User Guide.
//
// Although it is possible to call GetFederationToken using the security
// credentials of an Amazon Web Services account root user rather than an IAM user
// that you create for the purpose of a proxy application, we do not recommend it.
// For more information, see [Safeguard your root user credentials and don't use them for everyday tasks]in the IAM User Guide.
//
// You can create a mobile-based or browser-based app that can authenticate users
// using a web identity provider like Login with Amazon, Facebook, Google, or an
// OpenID Connect-compatible identity provider. In this case, we recommend that you
// use [Amazon Cognito]or AssumeRoleWithWebIdentity . For more information, see [Federation Through a Web-based Identity Provider] in the IAM User
// Guide.
//
// # Session duration
//
// The temporary credentials are valid for the specified duration, from 900
// seconds (15 minutes) up to a maximum of 129,600 seconds (36 hours). The default
// session duration is 43,200 seconds (12 hours). Temporary credentials obtained by
// using the root user credentials have a maximum duration of 3,600 seconds (1
// hour).
//
// # Permissions
//
// You can use the temporary credentials created by GetFederationToken in any
// Amazon Web Services service with the following exceptions:
//
// - You cannot call any IAM operations using the CLI or the Amazon Web Services
// API. This limitation does not apply to console sessions.
//
// - You cannot call any STS operations except GetCallerIdentity .
//
// You can use temporary credentials for single sign-on (SSO) to the console.
//
// You must pass an inline or managed [session policy] to this operation. You can pass a single
// JSON policy document to use as an inline session policy. You can also specify up
// to 10 managed policy Amazon Resource Names (ARNs) to use as managed session
// policies. The plaintext that you use for both inline and managed session
// policies can't exceed 2,048 characters.
//
// Though the session policy parameters are optional, if you do not pass a policy,
// then the resulting federated user session has no permissions. When you pass
// session policies, the session permissions are the intersection of the IAM user
// policies and the session policies that you pass. This gives you a way to further
// restrict the permissions for a federated user. You cannot use session policies
// to grant more permissions than those that are defined in the permissions policy
// of the IAM user. For more information, see [Session Policies]in the IAM User Guide. For
// information about using GetFederationToken to create temporary security
// credentials, see [GetFederationToken—Federation Through a Custom Identity Broker].
//
// You can use the credentials to access a resource that has a resource-based
// policy. If that policy specifically references the federated user session in the
// Principal element of the policy, the session has the permissions allowed by the
// policy. These permissions are granted in addition to the permissions granted by
// the session policies.
//
// # Tags
//
// (Optional) You can pass tag key-value pairs to your session. These are called
// session tags. For more information about session tags, see [Passing Session Tags in STS]in the IAM User
// Guide.
//
// You can create a mobile-based or browser-based app that can authenticate users
// using a web identity provider like Login with Amazon, Facebook, Google, or an
// OpenID Connect-compatible identity provider. In this case, we recommend that you
// use [Amazon Cognito]or AssumeRoleWithWebIdentity . For more information, see [Federation Through a Web-based Identity Provider] in the IAM User
// Guide.
//
// An administrator must grant you the permissions necessary to pass session tags.
// The administrator can also create granular permissions to allow you to pass only
// specific session tags. For more information, see [Tutorial: Using Tags for Attribute-Based Access Control]in the IAM User Guide.
//
// Tag key–value pairs are not case sensitive, but case is preserved. This means
// that you cannot have separate Department and department tag keys. Assume that
// the user that you are federating has the Department = Marketing tag and you
// pass the department = engineering session tag. Department and department are
// not saved as separate tags, and the session tag passed in the request takes
// precedence over the user tag.
//
// [Federation Through a Web-based Identity Provider]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html#api_assumerolewithwebidentity
// [session policy]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#policies_session
// [Amazon Cognito]: http://aws.amazon.com/cognito/
// [Session Policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#policies_session
// [Passing Session Tags in STS]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_session-tags.html
// [GetFederationToken—Federation Through a Custom Identity Broker]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html#api_getfederationtoken
// [Safeguard your root user credentials and don't use them for everyday tasks]: https://docs.aws.amazon.com/IAM/latest/UserGuide/best-practices.html#lock-away-credentials
// [Requesting Temporary Security Credentials]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html
// [Compare STS credentials]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_sts-comparison.html
// [Tutorial: Using Tags for Attribute-Based Access Control]: https://docs.aws.amazon.com/IAM/latest/UserGuide/tutorial_attribute-based-access-control.html
func sts_GetFederationToken(cfg aws.Config, client *sts.Client) {
	input := &sts.GetFederationTokenInput{
		// Name: *string, // Required
	}

	if len(_stsName) > 0 {
		input.Name = aws.String(_stsName)
	}
	if len(_stsDurationSeconds) > 0 {
		if err := assignInputField(input, "DurationSeconds", _stsDurationSeconds); err != nil {
			log.Errorf("invalid --duration-seconds: %s", err.Error())
			return
		}
	}
	if len(_stsPolicy) > 0 {
		input.Policy = aws.String(_stsPolicy)
	}
	if len(_stsPolicyArns) > 0 {
		if err := assignInputField(input, "PolicyArns", _stsPolicyArns); err != nil {
			log.Errorf("invalid --policy-arns: %s", err.Error())
			return
		}
	}
	if len(_stsTags) > 0 {
		if err := assignInputField(input, "Tags", _stsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetFederationToken(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a set of temporary credentials for an Amazon Web Services account or
// IAM user. The credentials consist of an access key ID, a secret access key, and
// a security token. Typically, you use GetSessionToken if you want to use MFA to
// protect programmatic calls to specific Amazon Web Services API operations like
// Amazon EC2 StopInstances .
//
// MFA-enabled IAM users must call GetSessionToken and submit an MFA code that is
// associated with their MFA device. Using the temporary security credentials that
// the call returns, IAM users can then make programmatic calls to API operations
// that require MFA authentication. An incorrect MFA code causes the API to return
// an access denied error. For a comparison of GetSessionToken with the other API
// operations that produce temporary credentials, see [Requesting Temporary Security Credentials]and [Compare STS credentials] in the IAM User Guide.
//
// No permissions are required for users to perform this operation. The purpose of
// the sts:GetSessionToken operation is to authenticate the user using MFA. You
// cannot use policies to control authentication operations. For more information,
// see [Permissions for GetSessionToken]in the IAM User Guide.
//
// # Session Duration
//
// The GetSessionToken operation must be called by using the long-term Amazon Web
// Services security credentials of an IAM user. Credentials that are created by
// IAM users are valid for the duration that you specify. This duration can range
// from 900 seconds (15 minutes) up to a maximum of 129,600 seconds (36 hours),
// with a default of 43,200 seconds (12 hours). Credentials based on account
// credentials can range from 900 seconds (15 minutes) up to 3,600 seconds (1
// hour), with a default of 1 hour.
//
// # Permissions
//
// The temporary security credentials created by GetSessionToken can be used to
// make API calls to any Amazon Web Services service with the following exceptions:
//
// - You cannot call any IAM API operations unless MFA authentication
// information is included in the request.
//
// - You cannot call any STS API except AssumeRole or GetCallerIdentity .
//
// The credentials that GetSessionToken returns are based on permissions
// associated with the IAM user whose credentials were used to call the operation.
// The temporary credentials have the same permissions as the IAM user.
//
// Although it is possible to call GetSessionToken using the security credentials
// of an Amazon Web Services account root user rather than an IAM user, we do not
// recommend it. If GetSessionToken is called using root user credentials, the
// temporary credentials have root user permissions. For more information, see [Safeguard your root user credentials and don't use them for everyday tasks]in
// the IAM User Guide
//
// For more information about using GetSessionToken to create temporary
// credentials, see [Temporary Credentials for Users in Untrusted Environments]in the IAM User Guide.
//
// [Permissions for GetSessionToken]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_control-access_getsessiontoken.html
// [Temporary Credentials for Users in Untrusted Environments]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html#api_getsessiontoken
// [Safeguard your root user credentials and don't use them for everyday tasks]: https://docs.aws.amazon.com/IAM/latest/UserGuide/best-practices.html#lock-away-credentials
// [Requesting Temporary Security Credentials]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html
// [Compare STS credentials]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_sts-comparison.html
func sts_GetSessionToken(cfg aws.Config, client *sts.Client) {
	input := &sts.GetSessionTokenInput{}

	if len(_stsDurationSeconds) > 0 {
		if err := assignInputField(input, "DurationSeconds", _stsDurationSeconds); err != nil {
			log.Errorf("invalid --duration-seconds: %s", err.Error())
			return
		}
	}
	if len(_stsSerialNumber) > 0 {
		input.SerialNumber = aws.String(_stsSerialNumber)
	}
	if len(_stsTokenCode) > 0 {
		input.TokenCode = aws.String(_stsTokenCode)
	}

	if resp, err := client.GetSessionToken(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a signed JSON Web Token (JWT) that represents the calling Amazon Web
// Services identity. The returned JWT can be used to authenticate with external
// services that support OIDC discovery. The token is signed by Amazon Web Services
// STS and can be publicly verified using the verification keys published at the
// issuer's JWKS endpoint.
func sts_GetWebIdentityToken(cfg aws.Config, client *sts.Client) {
	input := &sts.GetWebIdentityTokenInput{
		// Audience: []string, // Required
		// SigningAlgorithm: *string, // Required
	}

	if len(_stsAudience) > 0 {
		input.Audience = append([]string(nil), _stsAudience...)
	}
	if len(_stsSigningAlgorithm) > 0 {
		input.SigningAlgorithm = aws.String(_stsSigningAlgorithm)
	}
	if len(_stsDurationSeconds) > 0 {
		if err := assignInputField(input, "DurationSeconds", _stsDurationSeconds); err != nil {
			log.Errorf("invalid --duration-seconds: %s", err.Error())
			return
		}
	}
	if len(_stsTags) > 0 {
		if err := assignInputField(input, "Tags", _stsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetWebIdentityToken(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_stsCmd)
	_stsCmd.Flags().SortFlags = false

	_stsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_stsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_stsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_stsCmd.Flags().StringVarP(&_stsAccessKeyId, "access-key-id", "", "", "Access Key ID")
	_stsCmd.Flags().StringSliceVarP(&_stsAudience, "audience", "", nil, "Audience")
	_stsCmd.Flags().StringVarP(&_stsDurationSeconds, "duration-seconds", "", "", "Duration Seconds")
	_stsCmd.Flags().StringVarP(&_stsEncodedMessage, "encoded-message", "", "", "Encoded Message")
	_stsCmd.Flags().StringVarP(&_stsExternalId, "external-id", "", "", "External ID")
	_stsCmd.Flags().StringVarP(&_stsName, "name", "", "", "Name")
	_stsCmd.Flags().StringVarP(&_stsPolicy, "policy", "", "", "Policy")
	_stsCmd.Flags().StringVarP(&_stsPolicyArns, "policy-arns", "", "", "Policy Arns")
	_stsCmd.Flags().StringVarP(&_stsPrincipalArn, "principal-arn", "", "", "Principal ARN")
	_stsCmd.Flags().StringVarP(&_stsProvidedContexts, "provided-contexts", "", "", "Provided Contexts")
	_stsCmd.Flags().StringVarP(&_stsProviderId, "provider-id", "", "", "Provider ID")
	_stsCmd.Flags().StringVarP(&_stsRoleArn, "role-arn", "", "", "Role ARN")
	_stsCmd.Flags().StringVarP(&_stsRoleSessionName, "role-session-name", "", "", "Role Session Name")
	_stsCmd.Flags().StringVarP(&_stsSAMLAssertion, "saml-assertion", "", "", "Saml Assertion")
	_stsCmd.Flags().StringVarP(&_stsSerialNumber, "serial-number", "", "", "Serial Number")
	_stsCmd.Flags().StringVarP(&_stsSigningAlgorithm, "signing-algorithm", "", "", "Signing Algorithm")
	_stsCmd.Flags().StringVarP(&_stsSourceIdentity, "source-identity", "", "", "Source Identity")
	_stsCmd.Flags().StringVarP(&_stsTags, "tags", "", "", "Tags")
	_stsCmd.Flags().StringVarP(&_stsTargetPrincipal, "target-principal", "", "", "Target Principal")
	_stsCmd.Flags().StringVarP(&_stsTaskPolicyArn, "task-policy-arn", "", "", "Task Policy ARN")
	_stsCmd.Flags().StringVarP(&_stsTokenCode, "token-code", "", "", "Token Code")
	_stsCmd.Flags().StringVarP(&_stsTradeInToken, "trade-in-token", "", "", "Trade In Token")
	_stsCmd.Flags().StringSliceVarP(&_stsTransitiveTagKeys, "transitive-tag-keys", "", nil, "Transitive Tag Keys")
	_stsCmd.Flags().StringVarP(&_stsWebIdentityToken, "web-identity-token", "", "", "Web Identity Token")

	_stsCmd.Flags().BoolVarP(&_stsAssumeRole, "assume-role", "", false, "Assume Role")
	_stsCmd.Flags().BoolVarP(&_stsAssumeRoleWithSAML, "assume-role-with-saml", "", false, "Assume Role With Saml")
	_stsCmd.Flags().BoolVarP(&_stsAssumeRoleWithWebIdentity, "assume-role-with-web-identity", "", false, "Assume Role With Web Identity")
	_stsCmd.Flags().BoolVarP(&_stsAssumeRoot, "assume-root", "", false, "Assume Root")
	_stsCmd.Flags().BoolVarP(&_stsDecodeAuthorizationMessage, "decode-authorization-message", "", false, "Decode Authorization Message")
	_stsCmd.Flags().BoolVarP(&_stsGetAccessKeyInfo, "get-access-key-info", "", false, "Get Access Key Info")
	_stsCmd.Flags().BoolVarP(&_stsGetCallerIdentity, "get-caller-identity", "", false, "Get Caller Identity")
	_stsCmd.Flags().BoolVarP(&_stsGetDelegatedAccessToken, "get-delegated-access-token", "", false, "Get Delegated Access Token")
	_stsCmd.Flags().BoolVarP(&_stsGetFederationToken, "get-federation-token", "", false, "Get Federation Token")
	_stsCmd.Flags().BoolVarP(&_stsGetSessionToken, "get-session-token", "", false, "Get Session Token")
	_stsCmd.Flags().BoolVarP(&_stsGetWebIdentityToken, "get-web-identity-token", "", false, "Get Web Identity Token")

}
