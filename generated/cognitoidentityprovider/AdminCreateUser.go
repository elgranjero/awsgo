package cognitoidentityprovider

// AdminCreateUser is generated as a reference stub.
// Executable command wiring lives under cmd/cognitoidentityprovider.go.
//
// Creates a new user in the specified user pool.
//
// If MessageAction isn't set, the default is to send a welcome message via email
// or phone (SMS).
//
// This message is based on a template that you configured in your call to create
// or update a user pool. This template includes your custom sign-up instructions
// and placeholders for user name and temporary password.
//
// Alternatively, you can call AdminCreateUser with SUPPRESS for the MessageAction
// parameter, and Amazon Cognito won't send any email.
//
// In either case, if the user has a password, they will be in the
// FORCE_CHANGE_PASSWORD state until they sign in and set their password. Your
// invitation message template must have the {####} password placeholder if your
// users have passwords. If your template doesn't have this placeholder, Amazon
// Cognito doesn't deliver the invitation message. In this case, you must update
// your message template and resend the password with a new AdminCreateUser
// request with a MessageAction value of RESEND .
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
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
// [Amazon Pinpoint]: https://console.aws.amazon.com/pinpoint/home/
