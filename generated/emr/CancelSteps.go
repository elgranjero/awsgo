package emr

// CancelSteps is generated as a reference stub.
// Executable command wiring lives under cmd/emr.go.
//
// Cancels a pending step or steps in a running cluster. Available only in Amazon
// EMR versions 4.8.0 and later, excluding version 5.0.0. A maximum of 256 steps
// are allowed in each CancelSteps request. CancelSteps is idempotent but
// asynchronous; it does not guarantee that a step will be canceled, even if the
// request is successfully submitted. When you use Amazon EMR releases 5.28.0 and
// later, you can cancel steps that are in a PENDING or RUNNING state. In earlier
// versions of Amazon EMR, you can only cancel steps that are in a PENDING state.
