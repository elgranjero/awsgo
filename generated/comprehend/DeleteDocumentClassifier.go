package comprehend

// DeleteDocumentClassifier is generated as a reference stub.
// Executable command wiring lives under cmd/comprehend.go.
//
// Deletes a previously created document classifier
//
// Only those classifiers that are in terminated states (IN_ERROR, TRAINED) will
// be deleted. If an active inference job is using the model, a
// ResourceInUseException will be returned.
//
// This is an asynchronous action that puts the classifier into a DELETING state,
// and it is then removed by a background job. Once removed, the classifier
// disappears from your account and is no longer available for use.
