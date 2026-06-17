package s3vectors

// PutVectors is generated as a reference stub.
// Executable command wiring lives under cmd/s3vectors.go.
//
// Adds one or more vectors to a vector index. To specify the vector index, you
// can either use both the vector bucket name and the vector index name, or use the
// vector index Amazon Resource Name (ARN).
//
// For more information about limits, see [Limitations and restrictions] in the Amazon S3 User Guide.
//
// When inserting vector data into your vector index, you must provide the vector
// data as float32 (32-bit floating point) values. If you pass higher-precision
// values to an Amazon Web Services SDK, S3 Vectors converts the values to 32-bit
// floating point before storing them, and GetVectors , ListVectors , and
// QueryVectors operations return the float32 values. Different Amazon Web Services
// SDKs may have different default numeric types, so ensure your vectors are
// properly formatted as float32 values regardless of which SDK you're using. For
// example, in Python, use numpy.float32 or explicitly cast your values.
//
// Permissions You must have the s3vectors:PutVectors permission to use this
// operation.
//
// [Limitations and restrictions]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-vectors-limitations.html
