package glue

// StartMLEvaluationTaskRun is generated as a reference stub.
// Executable command wiring lives under cmd/glue.go.
//
// Starts a task to estimate the quality of the transform.
//
// When you provide label sets as examples of truth, Glue machine learning uses
// some of those examples to learn from them. The rest of the labels are used as a
// test to estimate quality.
//
// Returns a unique identifier for the run. You can call GetMLTaskRun to get more
// information about the stats of the EvaluationTaskRun .
