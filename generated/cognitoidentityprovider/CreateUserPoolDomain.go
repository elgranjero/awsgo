package cognitoidentityprovider

// CreateUserPoolDomain is generated as a reference stub.
// Executable command wiring lives under cmd/cognitoidentityprovider.go.
//
// A user pool domain hosts managed login, an authorization server and web server
// for authentication in your application. This operation creates a new user pool
// prefix domain or custom domain and sets the managed login branding version. Set
// the branding version to 1 for hosted UI (classic) or 2 for managed login. When
// you choose a custom domain, you must provide an SSL certificate in the US East
// (N. Virginia) Amazon Web Services Region in your request.
//
// Your prefix domain might take up to one minute to take effect. Your custom
// domain is online within five minutes, but it can take up to one hour to
// distribute your SSL certificate.
//
// For more information about adding a custom domain to your user pool, see [Configuring a user pool domain].
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
// [Configuring a user pool domain]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-add-custom-domain.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
