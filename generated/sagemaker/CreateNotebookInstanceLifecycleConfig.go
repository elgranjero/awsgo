package sagemaker

// CreateNotebookInstanceLifecycleConfig is generated as a reference stub.
// Executable command wiring lives under cmd/sagemaker.go.
//
// Creates a lifecycle configuration that you can associate with a notebook
// instance. A lifecycle configuration is a collection of shell scripts that run
// when you create or start a notebook instance.
//
// Each lifecycle configuration script has a limit of 16384 characters.
//
// The value of the $PATH environment variable that is available to both scripts
// is /sbin:bin:/usr/sbin:/usr/bin .
//
// View Amazon CloudWatch Logs for notebook instance lifecycle configurations in
// log group /aws/sagemaker/NotebookInstances in log stream
// [notebook-instance-name]/[LifecycleConfigHook] .
//
// Lifecycle configuration scripts cannot run for longer than 5 minutes. If a
// script runs for longer than 5 minutes, it fails and the notebook instance is not
// created or started.
//
// For information about notebook instance lifestyle configurations, see [Step 2.1: (Optional) Customize a Notebook Instance].
//
// Lifecycle configuration scripts execute with root access and the notebook
// instance's IAM execution role privileges. Grant this permission only to trusted
// principals. See [Customize a Notebook Instance Using a Lifecycle Configuration Script]for security best practices.
//
// [Customize a Notebook Instance Using a Lifecycle Configuration Script]: https://docs.aws.amazon.com/sagemaker/latest/dg/notebook-lifecycle-config.html
// [Step 2.1: (Optional) Customize a Notebook Instance]: https://docs.aws.amazon.com/sagemaker/latest/dg/notebook-lifecycle-config.html
