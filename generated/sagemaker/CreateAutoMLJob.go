package sagemaker

// CreateAutoMLJob is generated as a reference stub.
// Executable command wiring lives under cmd/sagemaker.go.
//
// Creates an Autopilot job also referred to as Autopilot experiment or AutoML job.
//
// An AutoML job in SageMaker AI is a fully automated process that allows you to
// build machine learning models with minimal effort and machine learning
// expertise. When initiating an AutoML job, you provide your data and optionally
// specify parameters tailored to your use case. SageMaker AI then automates the
// entire model development lifecycle, including data preprocessing, model
// training, tuning, and evaluation. AutoML jobs are designed to simplify and
// accelerate the model building process by automating various tasks and exploring
// different combinations of machine learning algorithms, data preprocessing
// techniques, and hyperparameter values. The output of an AutoML job comprises one
// or more trained models ready for deployment and inference. Additionally,
// SageMaker AI AutoML jobs generate a candidate model leaderboard, allowing you to
// select the best-performing model for deployment.
//
// For more information about AutoML jobs, see [https://docs.aws.amazon.com/sagemaker/latest/dg/autopilot-automate-model-development.html] in the SageMaker AI developer
// guide.
//
// We recommend using the new versions [CreateAutoMLJobV2] and [DescribeAutoMLJobV2], which offer backward compatibility.
//
// CreateAutoMLJobV2 can manage tabular problem types identical to those of its
// previous version CreateAutoMLJob , as well as time-series forecasting,
// non-tabular problem types such as image or text classification, and text
// generation (LLMs fine-tuning).
//
// Find guidelines about how to migrate a CreateAutoMLJob to CreateAutoMLJobV2 in [Migrate a CreateAutoMLJob to CreateAutoMLJobV2].
//
// You can find the best-performing model after you run an AutoML job by calling [DescribeAutoMLJobV2]
// (recommended) or [DescribeAutoMLJob].
//
// [DescribeAutoMLJob]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeAutoMLJob.html
// [DescribeAutoMLJobV2]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeAutoMLJobV2.html
// [https://docs.aws.amazon.com/sagemaker/latest/dg/autopilot-automate-model-development.html]: https://docs.aws.amazon.com/sagemaker/latest/dg/autopilot-automate-model-development.html
// [CreateAutoMLJobV2]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateAutoMLJobV2.html
// [Migrate a CreateAutoMLJob to CreateAutoMLJobV2]: https://docs.aws.amazon.com/sagemaker/latest/dg/autopilot-automate-model-development-create-experiment.html#autopilot-create-experiment-api-migrate-v1-v2
