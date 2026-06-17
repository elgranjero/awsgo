package lambda

// AddPermission is generated as a reference stub.
// Executable command wiring lives under cmd/lambda.go.
//
// Grants a [principal] permission to use a function. You can apply the policy at the
// function level, or specify a qualifier to restrict access to a single version or
// alias. If you use a qualifier, the invoker must use the full Amazon Resource
// Name (ARN) of that version or alias to invoke the function. Note: Lambda does
// not support adding policies to version $LATEST.
//
// To grant permission to another account, specify the account ID as the Principal
// . To grant permission to an organization defined in Organizations, specify the
// organization ID as the PrincipalOrgID . For Amazon Web Services services, the
// principal is a domain-style identifier that the service defines, such as
// s3.amazonaws.com or sns.amazonaws.com . For Amazon Web Services services, you
// can also specify the ARN of the associated resource as the SourceArn . If you
// grant permission to a service principal without specifying the source, other
// accounts could potentially configure resources in their account to invoke your
// Lambda function.
//
// This operation adds a statement to a resource-based permissions policy for the
// function. For more information about function policies, see [Using resource-based policies for Lambda].
//
// [principal]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_principal.html#Principal_specifying
// [Using resource-based policies for Lambda]: https://docs.aws.amazon.com/lambda/latest/dg/access-control-resource-based.html
