package efs

// DescribeFileSystems is generated as a reference stub.
// Executable command wiring lives under cmd/efs.go.
//
// Returns the description of a specific Amazon EFS file system if either the file
// system CreationToken or the FileSystemId is provided. Otherwise, it returns
// descriptions of all file systems owned by the caller's Amazon Web Services
// account in the Amazon Web Services Region of the endpoint that you're calling.
//
// When retrieving all file system descriptions, you can optionally specify the
// MaxItems parameter to limit the number of descriptions in a response. This
// number is automatically set to 100. If more file system descriptions remain,
// Amazon EFS returns a NextMarker , an opaque token, in the response. In this
// case, you should send a subsequent request with the Marker request parameter
// set to the value of NextMarker .
//
// To retrieve a list of your file system descriptions, this operation is used in
// an iterative process, where DescribeFileSystems is called first without the
// Marker and then the operation continues to call it with the Marker parameter
// set to the value of the NextMarker from the previous response until the
// response has no NextMarker .
//
// The order of file systems returned in the response of one DescribeFileSystems
// call and the order of file systems returned across the responses of a multi-call
// iteration is unspecified.
//
// This operation requires permissions for the
// elasticfilesystem:DescribeFileSystems action.
