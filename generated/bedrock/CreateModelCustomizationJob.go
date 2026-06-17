package bedrock

// CreateModelCustomizationJob is generated as a reference stub.
// Executable command wiring lives under cmd/bedrock.go.
//
// Creates a fine-tuning job to customize a base model.
//
// You specify the base foundation model and the location of the training data.
// After the model-customization job completes successfully, your custom model
// resource will be ready to use. Amazon Bedrock returns validation loss metrics
// and output generations after the job completes.
//
// For information on the format of training and validation data, see [Prepare the datasets].
//
// Model-customization jobs are asynchronous and the completion time depends on
// the base model and the training/validation data size. To monitor a job, use the
// GetModelCustomizationJob operation to retrieve the job status.
//
// For more information, see [Custom models] in the [Amazon Bedrock User Guide].
//
// [Custom models]: https://docs.aws.amazon.com/bedrock/latest/userguide/custom-models.html
// [Amazon Bedrock User Guide]: https://docs.aws.amazon.com/bedrock/latest/userguide/what-is-service.html
// [Prepare the datasets]: https://docs.aws.amazon.com/bedrock/latest/userguide/model-customization-prepare.html
