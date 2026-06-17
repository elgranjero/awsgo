package cognitoidentityprovider

// UpdateTerms is generated as a reference stub.
// Executable command wiring lives under cmd/cognitoidentityprovider.go.
//
// Modifies existing terms documents for the requested app client. When Terms and
// conditions and Privacy policy documents are configured, the app client displays
// links to them in the sign-up page of managed login for the app client.
//
// You can provide URLs for terms documents in the languages that are supported by [managed login localization]
// . Amazon Cognito directs users to the terms documents for their current
// language, with fallback to default if no document exists for the language.
//
// Each request accepts one type of terms document and a map of language-to-link
// for that document type. You must provide both types of terms documents in at
// least one language before Amazon Cognito displays your terms documents. Supply
// each type in separate requests.
//
// For more information, see [Terms documents].
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
// [Terms documents]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-managed-login.html#managed-login-terms-documents
// [managed login localization]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-managed-login.html#managed-login-localization
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
