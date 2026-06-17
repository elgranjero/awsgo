package bedrock

// CreateCustomModel is generated as a reference stub.
// Executable command wiring lives under cmd/bedrock.go.
//
// Creates a new custom model in Amazon Bedrock. After the model is active, you
// can use it for inference.
//
// To use the model for inference, you must purchase Provisioned Throughput for
// it. You can't use On-demand inference with these custom models. For more
// information about Provisioned Throughput, see [Provisioned Throughput].
//
// The model appears in ListCustomModels with a customizationType of imported . To
// track the status of the new model, you use the GetCustomModel API operation.
// The model can be in the following states:
//
// - Creating - Initial state during validation and registration
//
// - Active - Model is ready for use in inference
//
// - Failed - Creation process encountered an error
//
// # Related APIs
//
// [GetCustomModel]
//
// [ListCustomModels]
//
// [DeleteCustomModel]
//
// [Provisioned Throughput]: https://docs.aws.amazon.com/bedrock/latest/userguide/prov-throughput.html
// [ListCustomModels]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_ListCustomModels.html
// [DeleteCustomModel]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_DeleteCustomModel.html
// [GetCustomModel]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_GetCustomModel.html
