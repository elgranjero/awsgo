package cognitoidentityprovider

// UpdateManagedLoginBranding is generated as a reference stub.
// Executable command wiring lives under cmd/cognitoidentityprovider.go.
//
// Configures the branding settings for a user pool style. This operation is the
// programmatic option for the configuration of a style in the branding editor.
//
// Provides values for UI customization in a Settings JSON object and image files
// in an Assets array.
//
// This operation has a 2-megabyte request-size limit and include the CSS settings
// and image assets for your app client. Your branding settings might exceed 2MB in
// size. Amazon Cognito doesn't require that you pass all parameters in one request
// and preserves existing style settings that you don't specify. If your request is
// larger than 2MB, separate it into multiple requests, each with a size smaller
// than the limit.
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
