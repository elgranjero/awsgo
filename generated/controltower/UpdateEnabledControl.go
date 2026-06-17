package controltower

// UpdateEnabledControl is generated as a reference stub.
// Executable command wiring lives under cmd/controltower.go.
//
// Updates the configuration of an already enabled control.
//
// If the enabled control shows an EnablementStatus of SUCCEEDED, supply
// parameters that are different from the currently configured parameters.
// Otherwise, Amazon Web Services Control Tower will not accept the request.
//
// If the enabled control shows an EnablementStatus of FAILED, Amazon Web Services
// Control Tower updates the control to match any valid parameters that you supply.
//
// If the DriftSummary status for the control shows as DRIFTED , you cannot call
// this API. Instead, you can update the control by calling the ResetEnabledControl
// API. Alternatively, you can call DisableControl and then call EnableControl
// again. Also, you can run an extending governance operation to repair drift. For
// usage examples, see the [Controls Reference Guide].
//
// [Controls Reference Guide]: https://docs.aws.amazon.com/controltower/latest/controlreference/control-api-examples-short.html
