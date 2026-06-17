package translate

// StopTextTranslationJob is generated as a reference stub.
// Executable command wiring lives under cmd/translate.go.
//
// Stops an asynchronous batch translation job that is in progress.
//
// If the job's state is IN_PROGRESS , the job will be marked for termination and
// put into the STOP_REQUESTED state. If the job completes before it can be
// stopped, it is put into the COMPLETED state. Otherwise, the job is put into the
// STOPPED state.
//
// Asynchronous batch translation jobs are started with the StartTextTranslationJob operation. You can
// use the DescribeTextTranslationJobor ListTextTranslationJobs operations to get a batch translation job's JobId .
