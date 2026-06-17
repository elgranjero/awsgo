package sagemaker

// UpdateNotebookInstance is generated as a reference stub.
// Executable command wiring lives under cmd/sagemaker.go.
//
// Updates a notebook instance. NotebookInstance updates include upgrading or
// downgrading the ML compute instance used for your notebook instance to
// accommodate changes in your workload requirements.
//
// This API can attach lifecycle configurations to notebook instances. Lifecycle
// configuration scripts execute with root access and the notebook instance's IAM
// execution role privileges. Principals with this permission and access to
// lifecycle configurations can execute code with the execution role's credentials.
// See [Customize a Notebook Instance Using a Lifecycle Configuration Script]for security best practices.
//
// [Customize a Notebook Instance Using a Lifecycle Configuration Script]: https://docs.aws.amazon.com/sagemaker/latest/dg/notebook-lifecycle-config.html
