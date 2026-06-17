package pipes

// UpdatePipe is generated as a reference stub.
// Executable command wiring lives under cmd/pipes.go.
//
// Update an existing pipe. When you call UpdatePipe , EventBridge only the updates
// fields you have specified in the request; the rest remain unchanged. The
// exception to this is if you modify any Amazon Web Services-service specific
// fields in the SourceParameters , EnrichmentParameters , or TargetParameters
// objects. For example, DynamoDBStreamParameters or EventBridgeEventBusParameters
// . EventBridge updates the fields in these objects atomically as one and
// overrides existing values. This is by design, and means that if you don't
// specify an optional field in one of these Parameters objects, EventBridge sets
// that field to its system-default value during the update.
//
// For more information about pipes, see [Amazon EventBridge Pipes] in the Amazon EventBridge User Guide.
//
// [Amazon EventBridge Pipes]: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes.html
