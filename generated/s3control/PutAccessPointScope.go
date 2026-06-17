package s3control

// PutAccessPointScope is generated as a reference stub.
// Executable command wiring lives under cmd/s3control.go.
//
// Creates or replaces the access point scope for a directory bucket. You can use
// the access point scope to restrict access to specific prefixes, API operations,
// or a combination of both.
//
// You can specify any amount of prefixes, but the total length of characters of
// all prefixes must be less than 256 bytes in size.
//
// To use this operation, you must have the permission to perform the
// s3express:PutAccessPointScope action.
//
// For information about REST API errors, see [REST error responses].
//
// [REST error responses]: https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html#RESTErrorResponses
