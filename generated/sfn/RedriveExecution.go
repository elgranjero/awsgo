package sfn

// RedriveExecution is generated as a reference stub.
// Executable command wiring lives under cmd/sfn.go.
//
// Restarts unsuccessful executions of Standard workflows that didn't complete
// successfully in the last 14 days. These include failed, aborted, or timed out
// executions. When you [redrive]an execution, it continues the failed execution from the
// unsuccessful step and uses the same input. Step Functions preserves the results
// and execution history of the successful steps, and doesn't rerun these steps
// when you redrive an execution. Redriven executions use the same state machine
// definition and execution ARN as the original execution attempt.
//
// For workflows that include an [Inline Map] or [Parallel] state, RedriveExecution API action
// reschedules and redrives only the iterations and branches that failed or
// aborted.
//
// To redrive a workflow that includes a Distributed Map state whose Map Run
// failed, you must redrive the [parent workflow]. The parent workflow redrives all the
// unsuccessful states, including a failed Map Run. If a Map Run was not started in
// the original execution attempt, the redriven parent workflow starts the Map Run.
//
// This API action is not supported by EXPRESS state machines.
//
// However, you can restart the unsuccessful executions of Express child workflows
// in a Distributed Map by redriving its Map Run. When you redrive a Map Run, the
// Express child workflows are rerun using the StartExecutionAPI action. For more information,
// see [Redriving Map Runs].
//
// You can redrive executions if your original execution meets the following
// conditions:
//
// - The execution status isn't SUCCEEDED .
//
// - Your workflow execution has not exceeded the redrivable period of 14 days.
// Redrivable period refers to the time during which you can redrive a given
// execution. This period starts from the day a state machine completes its
// execution.
//
// - The workflow execution has not exceeded the maximum open time of one year.
// For more information about state machine quotas, see [Quotas related to state machine executions].
//
// - The execution event history count is less than 24,999. Redriven executions
// append their event history to the existing event history. Make sure your
// workflow execution contains less than 24,999 events to accommodate the
// ExecutionRedriven history event and at least one other history event.
//
// [redrive]: https://docs.aws.amazon.com/step-functions/latest/dg/redrive-executions.html
// [parent workflow]: https://docs.aws.amazon.com/step-functions/latest/dg/use-dist-map-orchestrate-large-scale-parallel-workloads.html#dist-map-orchestrate-parallel-workloads-key-terms
// [Quotas related to state machine executions]: https://docs.aws.amazon.com/step-functions/latest/dg/limits-overview.html#service-limits-state-machine-executions
// [Parallel]: https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-parallel-state.html
// [Inline Map]: https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-map-state.html
// [Redriving Map Runs]: https://docs.aws.amazon.com/step-functions/latest/dg/redrive-map-run.html
