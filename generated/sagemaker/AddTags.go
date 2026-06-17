package sagemaker

// AddTags is generated as a reference stub.
// Executable command wiring lives under cmd/sagemaker.go.
//
// Adds or overwrites one or more tags for the specified SageMaker resource. You
// can add tags to notebook instances, training jobs, hyperparameter tuning jobs,
// batch transform jobs, models, labeling jobs, work teams, endpoint
// configurations, and endpoints.
//
// Each tag consists of a key and an optional value. Tag keys must be unique per
// resource. For more information about tags, see For more information, see [Amazon Web Services Tagging Strategies].
//
// Tags that you add to a hyperparameter tuning job by calling this API are also
// added to any training jobs that the hyperparameter tuning job launches after you
// call this API, but not to training jobs that the hyperparameter tuning job
// launched before you called this API. To make sure that the tags associated with
// a hyperparameter tuning job are also added to all training jobs that the
// hyperparameter tuning job launches, add the tags when you first create the
// tuning job by specifying them in the Tags parameter of [CreateHyperParameterTuningJob]
//
// Tags that you add to a SageMaker Domain or User Profile by calling this API are
// also added to any Apps that the Domain or User Profile launches after you call
// this API, but not to Apps that the Domain or User Profile launched before you
// called this API. To make sure that the tags associated with a Domain or User
// Profile are also added to all Apps that the Domain or User Profile launches, add
// the tags when you first create the Domain or User Profile by specifying them in
// the Tags parameter of [CreateDomain] or [CreateUserProfile].
//
// [CreateHyperParameterTuningJob]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateHyperParameterTuningJob.html
// [Amazon Web Services Tagging Strategies]: https://aws.amazon.com/answers/account-management/aws-tagging-strategies/
// [CreateUserProfile]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateUserProfile.html
// [CreateDomain]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateDomain.html
