package snowball

// GetJobManifest is generated as a reference stub.
// Executable command wiring lives under cmd/snowball.go.
//
// Returns a link to an Amazon S3 presigned URL for the manifest file associated
// with the specified JobId value. You can access the manifest file for up to 60
// minutes after this request has been made. To access the manifest file after 60
// minutes have passed, you'll have to make another call to the GetJobManifest
// action.
//
// The manifest is an encrypted file that you can download after your job enters
// the WithCustomer status. This is the only valid status for calling this API as
// the manifest and UnlockCode code value are used for securing your device and
// should only be used when you have the device. The manifest is decrypted by using
// the UnlockCode code value, when you pass both values to the Snow device through
// the Snowball client when the client is started for the first time.
//
// As a best practice, we recommend that you don't save a copy of an UnlockCode
// value in the same location as the manifest file for that job. Saving these
// separately helps prevent unauthorized parties from gaining access to the Snow
// device associated with that job.
//
// The credentials of a given job, including its manifest file and unlock code,
// expire 360 days after the job is created.
