package omics

// CreateMultipartReadSetUpload is generated as a reference stub.
// Executable command wiring lives under cmd/omics.go.
//
// Initiates a multipart read set upload for uploading partitioned source files
// into a sequence store. You can directly import source files from an EC2 instance
// and other local compute, or from an S3 bucket. To separate these source files
// into parts, use the split operation. Each part cannot be larger than 100 MB. If
// the operation is successful, it provides an uploadId which is required by the
// UploadReadSetPart API operation to upload parts into a sequence store.
//
// To continue uploading a multipart read set into your sequence store, you must
// use the UploadReadSetPart API operation to upload each part individually
// following the steps below:
//
// - Specify the uploadId obtained from the previous call to
// CreateMultipartReadSetUpload .
//
// - Upload parts for that uploadId .
//
// When you have finished uploading parts, use the CompleteMultipartReadSetUpload
// API to complete the multipart read set upload and to retrieve the final read set
// IDs in the response.
//
// To learn more about creating parts and the split operation, see [Direct upload to a sequence store] in the Amazon
// Web Services HealthOmics User Guide.
//
// [Direct upload to a sequence store]: https://docs.aws.amazon.com/omics/latest/dev/synchronous-uploads.html
