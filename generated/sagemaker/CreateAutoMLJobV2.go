package sagemaker

// CreateAutoMLJobV2 is generated as a reference stub.
// Executable command wiring lives under cmd/sagemaker.go.
//
// Creates an Autopilot job also referred to as Autopilot experiment or AutoML job
// V2.
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
// AutoML jobs V2 support various problem types such as regression, binary, and
// multiclass classification with tabular data, text and image classification,
// time-series forecasting, and fine-tuning of large language models (LLMs) for
// text generation.
//
// [CreateAutoMLJobV2]and [DescribeAutoMLJobV2] are new versions of [CreateAutoMLJob] and [DescribeAutoMLJob] which offer backward compatibility.
//
// CreateAutoMLJobV2 can manage tabular problem types identical to those of its
// previous version CreateAutoMLJob , as well as time-series forecasting,
// non-tabular problem types such as image or text classification, and text
// generation (LLMs fine-tuning).
//
// Find guidelines about how to migrate a CreateAutoMLJob to CreateAutoMLJobV2 in [Migrate a CreateAutoMLJob to CreateAutoMLJobV2].
//
// For the list of available problem types supported by CreateAutoMLJobV2 , see [AutoMLProblemTypeConfig].
//
// You can find the best-performing model after you run an AutoML job V2 by
// calling [DescribeAutoMLJobV2].
//
// [CreateAutoMLJob]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateAutoMLJob.html
// [DescribeAutoMLJob]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeAutoMLJob.html
// [DescribeAutoMLJobV2]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeAutoMLJobV2.html
// [https://docs.aws.amazon.com/sagemaker/latest/dg/autopilot-automate-model-development.html]: https://docs.aws.amazon.com/sagemaker/latest/dg/autopilot-automate-model-development.html
// [CreateAutoMLJobV2]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateAutoMLJobV2.html
// [Migrate a CreateAutoMLJob to CreateAutoMLJobV2]: https://docs.aws.amazon.com/sagemaker/latest/dg/autopilot-automate-model-development-create-experiment.html#autopilot-create-experiment-api-migrate-v1-v2
// [AutoMLProblemTypeConfig]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_AutoMLProblemTypeConfig.html
