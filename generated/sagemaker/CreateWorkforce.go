package sagemaker

// CreateWorkforce is generated as a reference stub.
// Executable command wiring lives under cmd/sagemaker.go.
//
// Use this operation to create a workforce. This operation will return an error
// if a workforce already exists in the Amazon Web Services Region that you
// specify. You can only create one workforce in each Amazon Web Services Region
// per Amazon Web Services account.
//
// If you want to create a new workforce in an Amazon Web Services Region where a
// workforce already exists, use the [DeleteWorkforce]API operation to delete the existing
// workforce and then use CreateWorkforce to create a new workforce.
//
// To create a private workforce using Amazon Cognito, you must specify a Cognito
// user pool in CognitoConfig . You can also create an Amazon Cognito workforce
// using the Amazon SageMaker console. For more information, see [Create a Private Workforce (Amazon Cognito)].
//
// To create a private workforce using your own OIDC Identity Provider (IdP),
// specify your IdP configuration in OidcConfig . Your OIDC IdP must support groups
// because groups are used by Ground Truth and Amazon A2I to create work teams. For
// more information, see [Create a Private Workforce (OIDC IdP)].
//
// [Create a Private Workforce (Amazon Cognito)]: https://docs.aws.amazon.com/sagemaker/latest/dg/sms-workforce-create-private.html
// [DeleteWorkforce]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DeleteWorkforce.html
// [Create a Private Workforce (OIDC IdP)]: https://docs.aws.amazon.com/sagemaker/latest/dg/sms-workforce-create-private-oidc.html
