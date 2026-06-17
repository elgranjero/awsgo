package connectcases

// CreateCase is generated as a reference stub.
// Executable command wiring lives under cmd/connectcases.go.
//
// If you provide a value for PerformedBy.UserArn you must also have [connect:DescribeUser] permission
// on the User ARN resource that you provide
//
// Creates a case in the specified Cases domain. Case system and custom fields are
// taken as an array id/value pairs with a declared data types.
//
// When creating a case from a template that has tag propagation configurations,
// the specified tags are automatically applied to the case.
//
// The following fields are required when creating a case:
//
// - customer_id - You must provide the full customer profile ARN in this format:
// arn:aws:profile:your_AWS_Region:your_AWS_account
// ID:domains/your_profiles_domain_name/profiles/profile_ID
//
// - title
//
// [connect:DescribeUser]: https://docs.aws.amazon.com/connect/latest/APIReference/API_DescribeUser.html
