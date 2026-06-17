package ecr

// GetAuthorizationToken is generated as a reference stub.
// Executable command wiring lives under cmd/ecr.go.
//
// Retrieves an authorization token. An authorization token represents your IAM
// authentication credentials and can be used to access any Amazon ECR registry
// that your IAM principal has access to. The authorization token is valid for 12
// hours.
//
// The authorizationToken returned is a base64 encoded string that can be decoded
// and used in a docker login command to authenticate to a registry. The CLI
// offers an get-login-password command that simplifies the login process. For
// more information, see [Registry authentication]in the Amazon Elastic Container Registry User Guide.
//
// [Registry authentication]: https://docs.aws.amazon.com/AmazonECR/latest/userguide/Registries.html#registry_auth
