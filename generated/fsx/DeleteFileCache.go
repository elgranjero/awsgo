package fsx

// DeleteFileCache is generated as a reference stub.
// Executable command wiring lives under cmd/fsx.go.
//
// Deletes an Amazon File Cache resource. After deletion, the cache no longer
// exists, and its data is gone.
//
// The DeleteFileCache operation returns while the cache has the DELETING status.
// You can check the cache deletion status by calling the [DescribeFileCaches]operation, which returns
// a list of caches in your account. If you pass the cache ID for a deleted cache,
// the DescribeFileCaches operation returns a FileCacheNotFound error.
//
// The data in a deleted cache is also deleted and can't be recovered by any means.
//
// [DescribeFileCaches]: https://docs.aws.amazon.com/fsx/latest/APIReference/API_DescribeFileCaches.html
