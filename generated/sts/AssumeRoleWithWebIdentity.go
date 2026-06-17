package sts

// AssumeRoleWithWebIdentity is generated as a reference stub.
// Executable command wiring lives under cmd/sts.go.
//
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
