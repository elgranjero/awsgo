package codepipeline

// RetryStageExecution is generated as a reference stub.
// Executable command wiring lives under cmd/codepipeline.go.
//
// You can retry a stage that has failed without having to run a pipeline again
// from the beginning. You do this by either retrying the failed actions in a stage
// or by retrying all actions in the stage starting from the first action in the
// stage. When you retry the failed actions in a stage, all actions that are still
// in progress continue working, and failed actions are triggered again. When you
// retry a failed stage from the first action in the stage, the stage cannot have
// any actions in progress. Before a stage can be retried, it must either have all
// actions failed or some actions failed and some succeeded.
