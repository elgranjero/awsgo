package cognitoidentityprovider

// AdminRespondToAuthChallenge is generated as a reference stub.
// Executable command wiring lives under cmd/cognitoidentityprovider.go.
//
// Some API operations in a user pool generate a challenge, like a prompt for an
// MFA code, for device authentication that bypasses MFA, or for a custom
// authentication challenge. An AdminRespondToAuthChallenge API request provides
// the answer to that challenge, like a code or a secure remote password (SRP). The
// parameters of a response to an authentication challenge vary with the type of
// challenge.
//
// For more information about custom authentication challenges, see [Custom authentication challenge Lambda triggers].
//
// This action might generate an SMS text message. Starting June 1, 2021, US
// telecom carriers require you to register an origination phone number before you
// can send SMS messages to US phone numbers. If you use SMS text messages in
// Amazon Cognito, you must register a phone number with [Amazon Pinpoint]. Amazon Cognito uses the
// registered number automatically. Otherwise, Amazon Cognito users who must
// receive SMS messages might not be able to sign up, activate their accounts, or
// sign in.
//
// If you have never used SMS text messages with Amazon Cognito or any other
// Amazon Web Services service, Amazon Simple Notification Service might place your
// account in the SMS sandbox. In [sandbox mode], you can send messages only to verified phone
// numbers. After you test your app while in the sandbox environment, you can move
// out of the sandbox and into production. For more information, see [SMS message settings for Amazon Cognito user pools]in the Amazon
// Cognito Developer Guide.
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
// [SMS message settings for Amazon Cognito user pools]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-sms-settings.html
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [sandbox mode]: https://docs.aws.amazon.com/sns/latest/dg/sns-sms-sandbox.html
// [Custom authentication challenge Lambda triggers]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-challenge.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
// [Amazon Pinpoint]: https://console.aws.amazon.com/pinpoint/home/
