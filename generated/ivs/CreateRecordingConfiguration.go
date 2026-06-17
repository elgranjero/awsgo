package ivs

// CreateRecordingConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/ivs.go.
//
// Creates a new recording configuration, used to enable recording to Amazon S3.
//
// Known issue: In the us-east-1 region, if you use the Amazon Web Services CLI to
// create a recording configuration, it returns success even if the S3 bucket is in
// a different region. In this case, the state of the recording configuration is
// CREATE_FAILED (instead of ACTIVE ). (In other regions, the CLI correctly returns
// failure if the bucket is in a different region.)
//
// Workaround: Ensure that your S3 bucket is in the same region as the recording
// configuration. If you create a recording configuration in a different region as
// your S3 bucket, delete that recording configuration and create a new one with an
// S3 bucket from the correct region.
