package cognitoidentityprovider

// SetUserMFAPreference is generated as a reference stub.
// Executable command wiring lives under cmd/cognitoidentityprovider.go.
//
// Set the user's multi-factor authentication (MFA) method preference, including
// which MFA factors are activated and if any are preferred. Only one factor can be
// set as preferred. The preferred MFA factor will be used to authenticate a user
// if multiple factors are activated. If multiple options are activated and no
// preference is set, a challenge to choose an MFA option will be returned during
// sign-in. If an MFA type is activated for a user, the user will be prompted for
// MFA during all sign-in attempts unless device tracking is turned on and the
// device has been trusted. If you want MFA to be applied selectively based on the
// assessed risk level of sign-in attempts, deactivate MFA for users and turn on
// Adaptive Authentication for the user pool.
//
// Authorize this action with a signed-in user's access token. It must include the
// scope aws.cognito.signin.user.admin .
//
// Amazon Cognito doesn't evaluate Identity and Access Management (IAM) policies
// in requests for this API operation. For this operation, you can't use IAM
// credentials to authorize requests, and you can't grant IAM permissions in
// policies. For more information about authorization models in Amazon Cognito, see
// [Using the Amazon Cognito user pools API and user pool endpoints].
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
