package comprehend

// DeleteEntityRecognizer is generated as a reference stub.
// Executable command wiring lives under cmd/comprehend.go.
//
// Deletes an entity recognizer.
//
// Only those recognizers that are in terminated states (IN_ERROR, TRAINED) will
// be deleted. If an active inference job is using the model, a
// ResourceInUseException will be returned.
//
// This is an asynchronous action that puts the recognizer into a DELETING state,
// and it is then removed by a background job. Once removed, the recognizer
// disappears from your account and is no longer available for use.
