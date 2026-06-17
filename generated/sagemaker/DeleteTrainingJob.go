package sagemaker

// DeleteTrainingJob is generated as a reference stub.
// Executable command wiring lives under cmd/sagemaker.go.
//
// Deletes a training job. After SageMaker deletes a training job, all of the
// metadata for the training job is lost. You can delete only training jobs that
// are in a terminal state ( Stopped , Failed , or Completed ) and don't retain an
// Available[managed warm pool] . You cannot delete a job that is in the InProgress or Stopping
// state. After deleting the job, you can reuse its name to create another training
// job.
//
// [managed warm pool]: https://docs.aws.amazon.com/sagemaker/latest/dg/train-warm-pools.html
