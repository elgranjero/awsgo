package connectcases

// UpdateCase is generated as a reference stub.
// Executable command wiring lives under cmd/connectcases.go.
//
// If you provide a value for PerformedBy.UserArn you must also have [connect:DescribeUser] permission
// on the User ARN resource that you provide
//
// Updates the values of fields on a case. Fields to be updated are received as an
// array of id/value pairs identical to the CreateCase input .
//
// If the action is successful, the service sends back an HTTP 200 response with
// an empty HTTP body.
//
// [connect:DescribeUser]: https://docs.aws.amazon.com/connect/latest/APIReference/API_DescribeUser.html
