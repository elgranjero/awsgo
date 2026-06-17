package omics

// UploadReadSetPart is generated as a reference stub.
// Executable command wiring lives under cmd/omics.go.
//
// Uploads a specific part of a read set into a sequence store. When you a upload
// a read set part with a part number that already exists, the new part replaces
// the existing one. This operation returns a JSON formatted response containing a
// string identifier that is used to confirm that parts are being added to the
// intended upload.
//
// For more information, see [Direct upload to a sequence store] in the Amazon Web Services HealthOmics User Guide.
//
// [Direct upload to a sequence store]: https://docs.aws.amazon.com/omics/latest/dev/synchronous-uploads.html
