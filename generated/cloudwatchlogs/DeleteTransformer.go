package cloudwatchlogs

// DeleteTransformer is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatchlogs.go.
//
// Deletes the log transformer for the specified log group. As soon as you do
// this, the transformation of incoming log events according to that transformer
// stops. If this account has an account-level transformer that applies to this log
// group, the log group begins using that account-level transformer when this
// log-group level transformer is deleted.
//
// After you delete a transformer, be sure to edit any metric filters or
// subscription filters that relied on the transformed versions of the log events.
