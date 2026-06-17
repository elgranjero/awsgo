package sts

// AssumeRoleWithSAML is generated as a reference stub.
// Executable command wiring lives under cmd/sts.go.
//
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
