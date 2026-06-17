package securitylake

// CreateCustomLogSource is generated as a reference stub.
// Executable command wiring lives under cmd/securitylake.go.
//
// Adds a third-party custom source in Amazon Security Lake, from the Amazon Web
// Services Region where you want to create a custom source. Security Lake can
// collect logs and events from third-party custom sources. After creating the
// appropriate IAM role to invoke Glue crawler, use this API to add a custom source
// name in Security Lake. This operation creates a partition in the Amazon S3
// bucket for Security Lake as the target location for log files from the custom
// source. In addition, this operation also creates an associated Glue table and an
// Glue crawler.
