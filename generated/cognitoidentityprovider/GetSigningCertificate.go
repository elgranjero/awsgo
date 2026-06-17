package cognitoidentityprovider

// GetSigningCertificate is generated as a reference stub.
// Executable command wiring lives under cmd/cognitoidentityprovider.go.
//
// Given a user pool ID, returns the signing certificate for SAML 2.0 federation.
//
// Issued certificates are valid for 10 years from the date of issue. Amazon
// Cognito issues and assigns a new signing certificate annually. This renewal
// process returns a new value in the response to GetSigningCertificate , but
// doesn't invalidate the original certificate.
//
// For more information, see [Signing SAML requests].
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
// [Signing SAML requests]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-SAML-signing-encryption.html#cognito-user-pools-SAML-signing
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
