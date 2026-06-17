package sagemaker

// DeleteEndpoint is generated as a reference stub.
// Executable command wiring lives under cmd/sagemaker.go.
//
// Deletes an endpoint. SageMaker frees up all of the resources that were deployed
// when the endpoint was created.
//
// SageMaker retires any custom KMS key grants associated with the endpoint,
// meaning you don't need to use the [RevokeGrant]API call.
//
// When you delete your endpoint, SageMaker asynchronously deletes associated
// endpoint resources such as KMS key grants. You might still see these resources
// in your account for a few minutes after deleting your endpoint. Do not delete or
// revoke the permissions for your [ExecutionRoleArn], otherwise SageMaker cannot delete these
// resources.
//
// [ExecutionRoleArn]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateModel.html#sagemaker-CreateModel-request-ExecutionRoleArn
// [RevokeGrant]: http://docs.aws.amazon.com/kms/latest/APIReference/API_RevokeGrant.html
