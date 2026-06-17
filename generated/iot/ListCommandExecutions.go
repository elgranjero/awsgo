package iot

// ListCommandExecutions is generated as a reference stub.
// Executable command wiring lives under cmd/iot.go.
//
// List all command executions.
//
// - You must provide only the startedTimeFilter or the completedTimeFilter
// information. If you provide both time filters, the API will generate an error.
// You can use this information to retrieve a list of command executions within a
// specific timeframe.
//
// - You must provide only the commandArn or the thingArn information depending
// on whether you want to list executions for a specific command or an IoT thing.
// If you provide both fields, the API will generate an error.
//
// For more information about considerations for using this API, see [List command executions in your account (CLI)].
//
// [List command executions in your account (CLI)]: https://docs.aws.amazon.com/iot/latest/developerguide/iot-remote-command-execution-start-monitor.html#iot-remote-command-execution-list-cli
