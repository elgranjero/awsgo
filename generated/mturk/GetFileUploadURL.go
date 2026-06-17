package mturk

// GetFileUploadURL is generated as a reference stub.
// Executable command wiring lives under cmd/mturk.go.
//
// The GetFileUploadURL operation generates and returns a temporary URL. You use
//
// the temporary URL to retrieve a file uploaded by a Worker as an answer to a
// FileUploadAnswer question for a HIT. The temporary URL is generated the instant
// the GetFileUploadURL operation is called, and is valid for 60 seconds. You can
// get a temporary file upload URL any time until the HIT is disposed. After the
// HIT is disposed, any uploaded files are deleted, and cannot be retrieved.
//
// Pending Deprecation on December 12, 2017. The Answer Specification
//
// structure will no longer support the FileUploadAnswer element to be used for
// the QuestionForm data structure. Instead, we recommend that Requesters who want
// to create HITs asking Workers to upload files to use Amazon S3.
