package codepipeline

// DeleteWebhook is generated as a reference stub.
// Executable command wiring lives under cmd/codepipeline.go.
//
// Deletes a previously created webhook by name. Deleting the webhook stops
// CodePipeline from starting a pipeline every time an external event occurs. The
// API returns successfully when trying to delete a webhook that is already
// deleted. If a deleted webhook is re-created by calling PutWebhook with the same
// name, it will have a different URL.
