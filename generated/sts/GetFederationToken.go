package sts

// GetFederationToken is generated as a reference stub.
// Executable command wiring lives under cmd/sts.go.
//
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
