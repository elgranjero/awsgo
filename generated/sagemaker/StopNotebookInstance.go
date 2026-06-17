package sagemaker

// StopNotebookInstance is generated as a reference stub.
// Executable command wiring lives under cmd/sagemaker.go.
//
// Terminates the ML compute instance. Before terminating the instance, SageMaker
// AI disconnects the ML storage volume from it. SageMaker AI preserves the ML
// storage volume. SageMaker AI stops charging you for the ML compute instance when
// you call StopNotebookInstance .
//
// To access data on the ML storage volume for a notebook instance that has been
// terminated, call the StartNotebookInstance API. StartNotebookInstance launches
// another ML compute instance, configures it, and attaches the preserved ML
// storage volume so you can continue your work.
