package sagemaker

// CreateInferenceComponent is generated as a reference stub.
// Executable command wiring lives under cmd/sagemaker.go.
//
// Creates an inference component, which is a SageMaker AI hosting object that you
// can use to deploy a model to an endpoint. In the inference component settings,
// you specify the model, the endpoint, and how the model utilizes the resources
// that the endpoint hosts. You can optimize resource utilization by tailoring how
// the required CPU cores, accelerators, and memory are allocated. You can deploy
// multiple inference components to an endpoint, where each inference component
// contains one model and the resource utilization needs for that individual model.
// After you deploy an inference component, you can directly invoke the associated
// model when you use the InvokeEndpoint API action.
