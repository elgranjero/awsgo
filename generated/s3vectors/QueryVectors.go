package s3vectors

// QueryVectors is generated as a reference stub.
// Executable command wiring lives under cmd/s3vectors.go.
//
// Performs an approximate nearest neighbor search query in a vector index using a
// query vector. By default, it returns the keys of approximate nearest neighbors.
// You can optionally include the computed distance (between the query vector and
// each vector in the response), the vector data, and metadata of each vector in
// the response.
//
// To specify the vector index, you can either use both the vector bucket name and
// the vector index name, or use the vector index Amazon Resource Name (ARN).
//
// Permissions You must have the s3vectors:QueryVectors permission to use this
// operation. Additional permissions are required based on the request parameters
// you specify:
//
// - With only s3vectors:QueryVectors permission, you can retrieve vector keys of
// approximate nearest neighbors and computed distances between these vectors. This
// permission is sufficient only when you don't set any metadata filters and don't
// request vector data or metadata (by keeping the returnMetadata parameter set
// to false or not specified).
//
// - If you specify a metadata filter or set returnMetadata to true, you must
// have both s3vectors:QueryVectors and s3vectors:GetVectors permissions. The
// request fails with a 403 Forbidden error if you request metadata filtering,
// vector data, or metadata without the s3vectors:GetVectors permission.
