package cognitoidentityprovider

// GetCSVHeader is generated as a reference stub.
// Executable command wiring lives under cmd/cognitoidentityprovider.go.
//
// Given a user pool ID, generates a comma-separated value (CSV) list populated
// with available user attributes in the user pool. This list is the header for the
// CSV file that determines the users in a user import job. Save the content of
// CSVHeader in the response as a .csv file and populate it with the usernames and
// attributes of users that you want to import. For more information about CSV user
// import, see [Importing users from a CSV file].
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
// [Importing users from a CSV file]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-using-import-tool.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
