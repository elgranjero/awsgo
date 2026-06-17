package sagemaker

// CreateModel is generated as a reference stub.
// Executable command wiring lives under cmd/sagemaker.go.
//
// Creates a model in SageMaker. In the request, you name the model and describe a
// primary container. For the primary container, you specify the Docker image that
// contains inference code, artifacts (from prior training), and a custom
// environment map that the inference code uses when you deploy the model for
// predictions.
//
// Use this API to create a model if you want to use SageMaker hosting services or
// run a batch transform job.
//
// To host your model, you create an endpoint configuration with the
// CreateEndpointConfig API, and then create an endpoint with the CreateEndpoint
// API. SageMaker then deploys all of the containers that you defined for the model
// in the hosting environment.
//
// To run a batch transform using your model, you start a job with the
// CreateTransformJob API. SageMaker uses your model and your dataset to get
// inferences which are then saved to a specified S3 location.
//
// In the request, you also provide an IAM role that SageMaker can assume to
// access model artifacts and docker image for deployment on ML compute hosting
// instances or for batch transform jobs. In addition, you also use the IAM role to
// manage permissions the inference code needs. For example, if the inference code
// access any other Amazon Web Services resources, you grant necessary permissions
// via this role.
