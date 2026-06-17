package comprehend

// StopTrainingDocumentClassifier is generated as a reference stub.
// Executable command wiring lives under cmd/comprehend.go.
//
// Stops a document classifier training job while in progress.
//
// If the training job state is TRAINING , the job is marked for termination and
// put into the STOP_REQUESTED state. If the training job completes before it can
// be stopped, it is put into the TRAINED ; otherwise the training job is stopped
// and put into the STOPPED state and the service sends back an HTTP 200 response
// with an empty HTTP body.
