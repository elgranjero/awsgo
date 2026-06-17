package sagemaker

// ListTrainingJobs is generated as a reference stub.
// Executable command wiring lives under cmd/sagemaker.go.
//
// Lists training jobs.
//
// When StatusEquals and MaxResults are set at the same time, the MaxResults
// number of training jobs are first retrieved ignoring the StatusEquals parameter
// and then they are filtered by the StatusEquals parameter, which is returned as
// a response.
//
// For example, if ListTrainingJobs is invoked with the following parameters:
//
// { ... MaxResults: 100, StatusEquals: InProgress ... }
//
// First, 100 trainings jobs with any status, including those other than InProgress
// , are selected (sorted according to the creation time, from the most current to
// the oldest). Next, those with a status of InProgress are returned.
//
// You can quickly test the API using the following Amazon Web Services CLI code.
//
// aws sagemaker list-training-jobs --max-results 100 --status-equals InProgress
