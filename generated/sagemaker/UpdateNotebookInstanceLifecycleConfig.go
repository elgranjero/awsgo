package sagemaker

// UpdateNotebookInstanceLifecycleConfig is generated as a reference stub.
// Executable command wiring lives under cmd/sagemaker.go.
//
// Updates a notebook instance lifecycle configuration created with the [CreateNotebookInstanceLifecycleConfig] API.
//
// Updates to lifecycle configurations affect all notebook instances using that
// configuration upon their next start. Lifecycle configuration scripts execute
// with root access and the notebook instance's IAM execution role privileges.
// Grant this permission only to trusted principals. See [Customize a Notebook Instance Using a Lifecycle Configuration Script]for security best
// practices.
//
// [Customize a Notebook Instance Using a Lifecycle Configuration Script]: https://docs.aws.amazon.com/sagemaker/latest/dg/notebook-lifecycle-config.html
// [CreateNotebookInstanceLifecycleConfig]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateNotebookInstanceLifecycleConfig.html
