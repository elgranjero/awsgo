package docdb

// CreateGlobalCluster is generated as a reference stub.
// Executable command wiring lives under cmd/docdb.go.
//
// Creates an Amazon DocumentDB global cluster that can span multiple multiple
// Amazon Web Services Regions. The global cluster contains one primary cluster
// with read-write capability, and up-to 10 read-only secondary clusters. Global
// clusters uses storage-based fast replication across regions with latencies less
// than one second, using dedicated infrastructure with no impact to your
// workload’s performance.
//
// You can create a global cluster that is initially empty, and then add a primary
// and a secondary to it. Or you can specify an existing cluster during the create
// operation, and this cluster becomes the primary of the global cluster.
//
// This action only applies to Amazon DocumentDB clusters.
