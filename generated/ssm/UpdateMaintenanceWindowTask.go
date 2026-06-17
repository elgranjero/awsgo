package ssm

// UpdateMaintenanceWindowTask is generated as a reference stub.
// Executable command wiring lives under cmd/ssm.go.
//
// Modifies a task assigned to a maintenance window. You can't change the task
// type, but you can change the following values:
//
// - TaskARN . For example, you can change a RUN_COMMAND task from
// AWS-RunPowerShellScript to AWS-RunShellScript .
//
// - ServiceRoleArn
//
// - TaskInvocationParameters
//
// - Priority
//
// - MaxConcurrency
//
// - MaxErrors
//
// One or more targets must be specified for maintenance window Run Command-type
// tasks. Depending on the task, targets are optional for other maintenance window
// task types (Automation, Lambda, and Step Functions). For more information about
// running tasks that don't specify targets, see [Registering maintenance window tasks without targets]in the Amazon Web Services
// Systems Manager User Guide.
//
// If the value for a parameter in UpdateMaintenanceWindowTask is null, then the
// corresponding field isn't modified. If you set Replace to true, then all fields
// required by the RegisterTaskWithMaintenanceWindowoperation are required for this request. Optional fields that
// aren't specified are set to null.
//
// When you update a maintenance window task that has options specified in
// TaskInvocationParameters , you must provide again all the
// TaskInvocationParameters values that you want to retain. The values you don't
// specify again are removed. For example, suppose that when you registered a Run
// Command task, you specified TaskInvocationParameters values for Comment ,
// NotificationConfig , and OutputS3BucketName . If you update the maintenance
// window task and specify only a different OutputS3BucketName value, the values
// for Comment and NotificationConfig are removed.
//
// [Registering maintenance window tasks without targets]: https://docs.aws.amazon.com/systems-manager/latest/userguide/maintenance-windows-targetless-tasks.html
