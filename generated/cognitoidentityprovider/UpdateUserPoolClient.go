package cognitoidentityprovider

// UpdateUserPoolClient is generated as a reference stub.
// Executable command wiring lives under cmd/cognitoidentityprovider.go.
//
// Given a user pool app client ID, updates the configuration. To avoid setting
// parameters to Amazon Cognito defaults, construct this API request to pass the
// existing configuration of your app client, modified to include the changes that
// you want to make.
//
// If you don't provide a value for an attribute, Amazon Cognito sets it to its
// default value.
//
// Unlike app clients created in the console, Amazon Cognito doesn't automatically
// assign a branding style to app clients that you configure with this API
// operation. Managed login and classic hosted UI pages aren't available for your
// client until after you apply a branding style.
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
