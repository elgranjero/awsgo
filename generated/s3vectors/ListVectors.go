package s3vectors

// ListVectors is generated as a reference stub.
// Executable command wiring lives under cmd/s3vectors.go.
//
// List vectors in the specified vector index. To specify the vector index, you
// can either use both the vector bucket name and the vector index name, or use the
// vector index Amazon Resource Name (ARN).
//
// ListVectors operations proceed sequentially; however, for faster performance on
// a large number of vectors in a vector index, applications can request a parallel
// ListVectors operation by providing the segmentCount and segmentIndex parameters.
//
// Permissions You must have the s3vectors:ListVectors permission to use this
// operation. Additional permissions are required based on the request parameters
// you specify:
//
// - With only s3vectors:ListVectors permission, you can list vector keys when
// returnData and returnMetadata are both set to false or not specified..
//
// - If you set returnData or returnMetadata to true, you must have both
// s3vectors:ListVectors and s3vectors:GetVectors permissions. The request fails
// with a 403 Forbidden error if you request vector data or metadata without the
// s3vectors:GetVectors permission.
