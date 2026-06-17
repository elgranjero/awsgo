package sagemaker

// CreateLabelingJob is generated as a reference stub.
// Executable command wiring lives under cmd/sagemaker.go.
//
// Creates a job that uses workers to label the data objects in your input
// dataset. You can use the labeled data to train machine learning models.
//
// You can select your workforce from one of three providers:
//
// - A private workforce that you create. It can include employees, contractors,
// and outside experts. Use a private workforce when want the data to stay within
// your organization or when a specific set of skills is required.
//
// - One or more vendors that you select from the Amazon Web Services
// Marketplace. Vendors provide expertise in specific areas.
//
// - The Amazon Mechanical Turk workforce. This is the largest workforce, but it
// should only be used for public data or data that has been stripped of any
// personally identifiable information.
//
// You can also use automated data labeling to reduce the number of data objects
// that need to be labeled by a human. Automated data labeling uses active learning
// to determine if a data object can be labeled by machine or if it needs to be
// sent to a human worker. For more information, see [Using Automated Data Labeling].
//
// The data objects to be labeled are contained in an Amazon S3 bucket. You create
// a manifest file that describes the location of each object. For more
// information, see [Using Input and Output Data].
//
// The output can be used as the manifest file for another labeling job or as
// training data for your machine learning models.
//
// You can use this operation to create a static labeling job or a streaming
// labeling job. A static labeling job stops if all data objects in the input
// manifest file identified in ManifestS3Uri have been labeled. A streaming
// labeling job runs perpetually until it is manually stopped, or remains idle for
// 10 days. You can send new data objects to an active ( InProgress ) streaming
// labeling job in real time. To learn how to create a static labeling job, see [Create a Labeling Job (API)]in
// the Amazon SageMaker Developer Guide. To learn how to create a streaming
// labeling job, see [Create a Streaming Labeling Job].
//
// [Using Automated Data Labeling]: https://docs.aws.amazon.com/sagemaker/latest/dg/sms-automated-labeling.html
// [Create a Streaming Labeling Job]: https://docs.aws.amazon.com/sagemaker/latest/dg/sms-streaming-create-job.html
// [Create a Labeling Job (API)]: https://docs.aws.amazon.com/sagemaker/latest/dg/sms-create-labeling-job-api.html
// [Using Input and Output Data]: https://docs.aws.amazon.com/sagemaker/latest/dg/sms-data.html
