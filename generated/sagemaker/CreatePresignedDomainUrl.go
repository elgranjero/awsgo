package sagemaker

// CreatePresignedDomainUrl is generated as a reference stub.
// Executable command wiring lives under cmd/sagemaker.go.
//
// Creates a URL for a specified UserProfile in a Domain. When accessed in a web
// browser, the user will be automatically signed in to the domain, and granted
// access to all of the Apps and files associated with the Domain's Amazon Elastic
// File System volume. This operation can only be called when the authentication
// mode equals IAM.
//
// The IAM role or user passed to this API defines the permissions to access the
// app. Once the presigned URL is created, no additional permission is required to
// access this URL. IAM authorization policies for this API are also enforced for
// every HTTP request and WebSocket frame that attempts to connect to the app.
//
// You can restrict access to this API and to the URL that it returns to a list of
// IP addresses, Amazon VPCs or Amazon VPC Endpoints that you specify. For more
// information, see [Connect to Amazon SageMaker AI Studio Through an Interface VPC Endpoint].
//
// - The URL that you get from a call to CreatePresignedDomainUrl has a default
// timeout of 5 minutes. You can configure this value using ExpiresInSeconds . If
// you try to use the URL after the timeout limit expires, you are directed to the
// Amazon Web Services console sign-in page.
//
// - The JupyterLab session default expiration time is 12 hours. You can
// configure this value using SessionExpirationDurationInSeconds.
//
// [Connect to Amazon SageMaker AI Studio Through an Interface VPC Endpoint]: https://docs.aws.amazon.com/sagemaker/latest/dg/studio-interface-endpoint.html
