package cognitoidentityprovider

// UpdateUserPoolDomain is generated as a reference stub.
// Executable command wiring lives under cmd/cognitoidentityprovider.go.
//
// A user pool domain hosts managed login, an authorization server and web server
// for authentication in your application. This operation updates the branding
// version for user pool domains between 1 for hosted UI (classic) and 2 for
// managed login. It also updates the SSL certificate for user pool custom domains.
//
// Changes to the domain branding version take up to one minute to take effect for
// a prefix domain and up to five minutes for a custom domain.
//
// This operation doesn't change the name of your user pool domain. To change your
// domain, delete it with DeleteUserPoolDomain and create a new domain with
// CreateUserPoolDomain .
//
// You can pass the ARN of a new Certificate Manager certificate in this request.
// Typically, ACM certificates automatically renew and you user pool can continue
// to use the same ARN. But if you generate a new certificate for your custom
// domain name, replace the original configuration with the new ARN in this
// request.
//
// ACM certificates for custom domains must be in the US East (N. Virginia) Amazon
// Web Services Region. After you submit your request, Amazon Cognito requires up
// to 1 hour to distribute your new certificate to your custom domain.
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
