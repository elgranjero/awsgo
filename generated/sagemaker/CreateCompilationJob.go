package sagemaker

// CreateCompilationJob is generated as a reference stub.
// Executable command wiring lives under cmd/sagemaker.go.
//
// Starts a model compilation job. After the model has been compiled, Amazon
// SageMaker AI saves the resulting model artifacts to an Amazon Simple Storage
// Service (Amazon S3) bucket that you specify.
//
// If you choose to host your model using Amazon SageMaker AI hosting services,
// you can use the resulting model artifacts as part of the model. You can also use
// the artifacts with Amazon Web Services IoT Greengrass. In that case, deploy them
// as an ML resource.
//
// In the request body, you provide the following:
//
// - A name for the compilation job
//
// - Information about the input model artifacts
//
// - The output location for the compiled model and the device (target) that the
// model runs on
//
// - The Amazon Resource Name (ARN) of the IAM role that Amazon SageMaker AI
// assumes to perform the model compilation job.
//
// You can also provide a Tag to track the model compilation job's resource use
// and costs. The response body contains the CompilationJobArn for the compiled
// job.
//
// To stop a model compilation job, use [StopCompilationJob]. To get information about a particular
// model compilation job, use [DescribeCompilationJob]. To get information about multiple model
// compilation jobs, use [ListCompilationJobs].
//
// [StopCompilationJob]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_StopCompilationJob.html
// [DescribeCompilationJob]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeCompilationJob.html
// [ListCompilationJobs]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_ListCompilationJobs.html
