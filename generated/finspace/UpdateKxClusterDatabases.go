package finspace

// UpdateKxClusterDatabases is generated as a reference stub.
// Executable command wiring lives under cmd/finspace.go.
//
// Updates the databases mounted on a kdb cluster, which includes the changesetId
// and all the dbPaths to be cached. This API does not allow you to change a
// database name or add a database if you created a cluster without one.
//
// Using this API you can point a cluster to a different changeset and modify a
// list of partitions being cached.
