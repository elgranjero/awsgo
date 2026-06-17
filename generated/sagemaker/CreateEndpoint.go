package sagemaker

// CreateEndpoint is generated as a reference stub.
// Executable command wiring lives under cmd/sagemaker.go.
//
// Creates an endpoint using the endpoint configuration specified in the request.
// SageMaker uses the endpoint to provision resources and deploy models. You create
// the endpoint configuration with the [CreateEndpointConfig]API.
//
// Use this API to deploy models using SageMaker hosting services.
//
// You must not delete an EndpointConfig that is in use by an endpoint that is
// live or while the UpdateEndpoint or CreateEndpoint operations are being
// performed on the endpoint. To update an endpoint, you must create a new
// EndpointConfig .
//
// The endpoint name must be unique within an Amazon Web Services Region in your
// Amazon Web Services account.
//
// When it receives the request, SageMaker creates the endpoint, launches the
// resources (ML compute instances), and deploys the model(s) on them.
//
// When you call [CreateEndpoint], a load call is made to DynamoDB to verify that your endpoint
// configuration exists. When you read data from a DynamoDB table supporting [Eventually Consistent Reads]
// Eventually Consistent Reads , the response might not reflect the results of a
// recently completed write operation. The response might include some stale data.
// If the dependent entities are not yet in DynamoDB, this causes a validation
// error. If you repeat your read request after a short time, the response should
// return the latest data. So retry logic is recommended to handle these possible
// issues. We also recommend that customers call [DescribeEndpointConfig]before calling [CreateEndpoint] to minimize the
// potential impact of a DynamoDB eventually consistent read.
//
// When SageMaker receives the request, it sets the endpoint status to Creating .
// After it creates the endpoint, it sets the status to InService . SageMaker can
// then process incoming requests for inferences. To check the status of an
// endpoint, use the [DescribeEndpoint]API.
//
// If any of the models hosted at this endpoint get model data from an Amazon S3
// location, SageMaker uses Amazon Web Services Security Token Service to download
// model artifacts from the S3 path you provided. Amazon Web Services STS is
// activated in your Amazon Web Services account by default. If you previously
// deactivated Amazon Web Services STS for a region, you need to reactivate Amazon
// Web Services STS for that region. For more information, see [Activating and Deactivating Amazon Web Services STS in an Amazon Web Services Region]in the Amazon Web
// Services Identity and Access Management User Guide.
//
// To add the IAM role policies for using this API operation, go to the [IAM console], and
// choose Roles in the left navigation pane. Search the IAM role that you want to
// grant access to use the [CreateEndpoint]and [CreateEndpointConfig] API operations, add the following policies to the
// role.
//
// - Option 1: For a full SageMaker access, search and attach the
// AmazonSageMakerFullAccess policy.
//
// - Option 2: For granting a limited access to an IAM role, paste the following
// Action elements manually into the JSON file of the IAM role:
//
// "Action": ["sagemaker:CreateEndpoint", "sagemaker:CreateEndpointConfig"]
//
// "Resource": [
//
// "arn:aws:sagemaker:region:account-id:endpoint/endpointName"
//
// "arn:aws:sagemaker:region:account-id:endpoint-config/endpointConfigName"
//
// ]
//
// For more information, see [SageMaker API Permissions: Actions, Permissions, and Resources Reference].
//
// [Eventually Consistent Reads]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadConsistency.html
// [IAM console]: https://console.aws.amazon.com/iam/
// [DescribeEndpoint]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeEndpoint.html
// [CreateEndpoint]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateEndpoint.html
// [CreateEndpointConfig]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateEndpointConfig.html
// [DescribeEndpointConfig]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeEndpointConfig.html
// [Activating and Deactivating Amazon Web Services STS in an Amazon Web Services Region]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_enable-regions.html
// [SageMaker API Permissions: Actions, Permissions, and Resources Reference]: https://docs.aws.amazon.com/sagemaker/latest/dg/api-permissions-reference.html
