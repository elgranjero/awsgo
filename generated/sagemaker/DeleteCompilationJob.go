package sagemaker

// DeleteCompilationJob is generated as a reference stub.
// Executable command wiring lives under cmd/sagemaker.go.
//
// Deletes the specified compilation job. This action deletes only the compilation
// job resource in Amazon SageMaker AI. It doesn't delete other resources that are
// related to that job, such as the model artifacts that the job creates, the
// compilation logs in CloudWatch, the compiled model, or the IAM role.
//
// You can delete a compilation job only if its current status is COMPLETED ,
// FAILED , or STOPPED . If the job status is STARTING or INPROGRESS , stop the
// job, and then delete it after its status becomes STOPPED .
