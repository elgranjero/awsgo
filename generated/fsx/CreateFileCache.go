package fsx

// CreateFileCache is generated as a reference stub.
// Executable command wiring lives under cmd/fsx.go.
//
// Creates a new Amazon File Cache resource.
//
// You can use this operation with a client request token in the request that
// Amazon File Cache uses to ensure idempotent creation. If a cache with the
// specified client request token exists and the parameters match, CreateFileCache
// returns the description of the existing cache. If a cache with the specified
// client request token exists and the parameters don't match, this call returns
// IncompatibleParameterError . If a file cache with the specified client request
// token doesn't exist, CreateFileCache does the following:
//
// - Creates a new, empty Amazon File Cache resource with an assigned ID, and an
// initial lifecycle state of CREATING .
//
// - Returns the description of the cache in JSON format.
//
// The CreateFileCache call returns while the cache's lifecycle state is still
// CREATING . You can check the cache creation status by calling the [DescribeFileCaches] operation,
// which returns the cache state along with other information.
//
// [DescribeFileCaches]: https://docs.aws.amazon.com/fsx/latest/APIReference/API_DescribeFileCaches.html
