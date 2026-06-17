package ssm

// DescribeMaintenanceWindowTasks is generated as a reference stub.
// Executable command wiring lives under cmd/ssm.go.
//
// Lists the tasks in a maintenance window.
//
// For maintenance window tasks without a specified target, you can't supply
// values for --max-errors and --max-concurrency . Instead, the system inserts a
// placeholder value of 1 , which may be reported in the response to this command.
// These values don't affect the running of your task and can be ignored.
