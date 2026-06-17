package cognitoidentityprovider

// ForgotPassword is generated as a reference stub.
// Executable command wiring lives under cmd/cognitoidentityprovider.go.
//
// Sends a password-reset confirmation code to the email address or phone number
// of the requested username. The message delivery method is determined by the
// user's available attributes and the AccountRecoverySetting configuration of the
// user pool.
//
// For the Username parameter, you can use the username or an email, phone, or
// preferred username alias.
//
// If neither a verified phone number nor a verified email exists, Amazon Cognito
// responds with an InvalidParameterException error . If your app client has a
// client secret and you don't provide a SECRET_HASH parameter, this API returns
// NotAuthorizedException .
//
// Amazon Cognito doesn't evaluate Identity and Access Management (IAM) policies
// in requests for this API operation. For this operation, you can't use IAM
// credentials to authorize requests, and you can't grant IAM permissions in
// policies. For more information about authorization models in Amazon Cognito, see
// [Using the Amazon Cognito user pools API and user pool endpoints].
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
// [SMS message settings for Amazon Cognito user pools]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-sms-settings.html
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [sandbox mode]: https://docs.aws.amazon.com/sns/latest/dg/sns-sms-sandbox.html
// [Amazon Pinpoint]: https://console.aws.amazon.com/pinpoint/home/
