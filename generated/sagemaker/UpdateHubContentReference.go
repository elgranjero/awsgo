package sagemaker

// UpdateHubContentReference is generated as a reference stub.
// Executable command wiring lives under cmd/sagemaker.go.
//
// Updates the contents of a SageMaker hub for a ModelReference resource. A
// ModelReference allows you to access public SageMaker JumpStart models from
// within your private hub.
//
// When using this API, you can update the MinVersion field for additional
// flexibility in the model version. You shouldn't update any additional fields
// when using this API, because the metadata in your private hub should match the
// public JumpStart model's metadata.
//
// If you want to update a Model or Notebook resource in your hub, use the
// UpdateHubContent API instead.
//
// For more information about adding model references to your hub, see [Add models to a private hub].
//
// [Add models to a private hub]: https://docs.aws.amazon.com/sagemaker/latest/dg/jumpstart-curated-hubs-admin-guide-add-models.html
