package ssm

// GetMaintenanceWindowTask is generated as a reference stub.
// Executable command wiring lives under cmd/ssm.go.
//
// Retrieves the details of a maintenance window task.
//
// For maintenance window tasks without a specified target, you can't supply
// values for --max-errors and --max-concurrency . Instead, the system inserts a
// placeholder value of 1 , which may be reported in the response to this command.
// These values don't affect the running of your task and can be ignored.
//
// To retrieve a list of tasks in a maintenance window, instead use the DescribeMaintenanceWindowTasks command.
