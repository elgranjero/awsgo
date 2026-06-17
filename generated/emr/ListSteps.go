package emr

// ListSteps is generated as a reference stub.
// Executable command wiring lives under cmd/emr.go.
//
// Provides a list of steps for the cluster in reverse order unless you specify
// stepIds with the request or filter by StepStates . You can specify a maximum of
// 10 stepIDs . The CLI automatically paginates results to return a list greater
// than 50 steps. To return more than 50 steps using the CLI, specify a Marker ,
// which is a pagination token that indicates the next set of steps to retrieve.
