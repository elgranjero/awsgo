package entityresolution

// GenerateMatchId is generated as a reference stub.
// Executable command wiring lives under cmd/entityresolution.go.
//
// Generates or retrieves Match IDs for records using a rule-based matching
// workflow. When you call this operation, it processes your records against the
// workflow's matching rules to identify potential matches. For existing records,
// it retrieves their Match IDs and associated rules. For records without matches,
// it generates new Match IDs. The operation saves results to Amazon S3.
//
// The processing type ( processingType ) you choose affects both the accuracy and
// response time of the operation. Additional charges apply for each API call,
// whether made through the Entity Resolution console or directly via the API. The
// rule-based matching workflow must exist and be active before calling this
// operation.
