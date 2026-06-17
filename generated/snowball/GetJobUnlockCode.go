package snowball

// GetJobUnlockCode is generated as a reference stub.
// Executable command wiring lives under cmd/snowball.go.
//
// Returns the UnlockCode code value for the specified job. A particular UnlockCode
// value can be accessed for up to 360 days after the associated job has been
// created.
//
// The UnlockCode value is a 29-character code with 25 alphanumeric characters and
// 4 hyphens. This code is used to decrypt the manifest file when it is passed
// along with the manifest to the Snow device through the Snowball client when the
// client is started for the first time. The only valid status for calling this API
// is WithCustomer as the manifest and Unlock code values are used for securing
// your device and should only be used when you have the device.
//
// As a best practice, we recommend that you don't save a copy of the UnlockCode
// in the same location as the manifest file for that job. Saving these separately
// helps prevent unauthorized parties from gaining access to the Snow device
// associated with that job.
