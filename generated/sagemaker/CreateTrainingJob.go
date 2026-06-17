package sagemaker

// CreateTrainingJob is generated as a reference stub.
// Executable command wiring lives under cmd/sagemaker.go.
//
// Starts a model training job. After training completes, SageMaker saves the
// resulting model artifacts to an Amazon S3 location that you specify.
//
// If you choose to host your model using SageMaker hosting services, you can use
// the resulting model artifacts as part of the model. You can also use the
// artifacts in a machine learning service other than SageMaker, provided that you
// know how to use them for inference.
//
// In the request body, you provide the following:
//
// - AlgorithmSpecification - Identifies the training algorithm to use.
//
// - HyperParameters - Specify these algorithm-specific parameters to enable the
// estimation of model parameters during training. Hyperparameters can be tuned to
// optimize this learning process. For a list of hyperparameters for each training
// algorithm provided by SageMaker, see [Algorithms].
//
// Do not include any security-sensitive information including account access IDs,
//
// secrets, or tokens in any hyperparameter fields. As part of the shared
// responsibility model, you are responsible for any potential exposure,
// unauthorized access, or compromise of your sensitive data if caused by
// security-sensitive information included in the request hyperparameter variable
// or plain text fields.
//
// - InputDataConfig - Describes the input required by the training job and the
// Amazon S3, EFS, or FSx location where it is stored.
//
// - OutputDataConfig - Identifies the Amazon S3 bucket where you want SageMaker
// to save the results of model training.
//
// - ResourceConfig - Identifies the resources, ML compute instances, and ML
// storage volumes to deploy for model training. In distributed training, you
// specify more than one instance.
//
// - EnableManagedSpotTraining - Optimize the cost of training machine learning
// models by up to 80% by using Amazon EC2 Spot instances. For more information,
// see [Managed Spot Training].
//
// - RoleArn - The Amazon Resource Name (ARN) that SageMaker assumes to perform
// tasks on your behalf during model training. You must grant this role the
// necessary permissions so that SageMaker can successfully complete model
// training.
//
// - StoppingCondition - To help cap training costs, use MaxRuntimeInSeconds to
// set a time limit for training. Use MaxWaitTimeInSeconds to specify how long a
// managed spot training job has to complete.
//
// - Environment - The environment variables to set in the Docker container.
//
// Do not include any security-sensitive information including account access IDs,
//
// secrets, or tokens in any environment fields. As part of the shared
// responsibility model, you are responsible for any potential exposure,
// unauthorized access, or compromise of your sensitive data if caused by
// security-sensitive information included in the request environment variable or
// plain text fields.
//
// - RetryStrategy - The number of times to retry the job when the job fails due
// to an InternalServerError .
//
// For more information about SageMaker, see [How It Works].
//
// [Algorithms]: https://docs.aws.amazon.com/sagemaker/latest/dg/algos.html
// [How It Works]: https://docs.aws.amazon.com/sagemaker/latest/dg/how-it-works.html
//
// [Managed Spot Training]: https://docs.aws.amazon.com/sagemaker/latest/dg/model-managed-spot-training.html
