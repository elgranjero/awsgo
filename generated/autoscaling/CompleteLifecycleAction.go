package autoscaling

// CompleteLifecycleAction is generated as a reference stub.
// Executable command wiring lives under cmd/autoscaling.go.
//
// Completes the lifecycle action for the specified token or instance with the
// specified result.
//
// This step is a part of the procedure for adding a lifecycle hook to an Auto
// Scaling group:
//
// - (Optional) Create a launch template or launch configuration with a user
// data script that runs while an instance is in a wait state due to a lifecycle
// hook.
//
// - (Optional) Create a Lambda function and a rule that allows Amazon
// EventBridge to invoke your Lambda function when an instance is put into a wait
// state due to a lifecycle hook.
//
// - (Optional) Create a notification target and an IAM role. The target can be
// either an Amazon SQS queue or an Amazon SNS topic. The role allows Amazon EC2
// Auto Scaling to publish lifecycle notifications to the target.
//
// - Create the lifecycle hook. Specify whether the hook is used when the
// instances launch or terminate.
//
// - If you need more time, record the lifecycle action heartbeat to keep the
// instance in a wait state.
//
// - If you finish before the timeout period ends, send a callback by using the [CompleteLifecycleAction]
// API call.
//
// For more information, see [Complete a lifecycle action] in the Amazon EC2 Auto Scaling User Guide.
//
// [CompleteLifecycleAction]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_CompleteLifecycleAction.html
// [Complete a lifecycle action]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/completing-lifecycle-hooks.html
