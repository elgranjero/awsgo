package elasticache

// RebootCacheCluster is generated as a reference stub.
// Executable command wiring lives under cmd/elasticache.go.
//
// Reboots some, or all, of the cache nodes within a provisioned cluster. This
// operation applies any modified cache parameter groups to the cluster. The reboot
// operation takes place as soon as possible, and results in a momentary outage to
// the cluster. During the reboot, the cluster status is set to REBOOTING.
//
// The reboot causes the contents of the cache (for each cache node being
// rebooted) to be lost.
//
// When the reboot is complete, a cluster event is created.
//
// Rebooting a cluster is currently supported on Memcached, Valkey and Redis OSS
// (cluster mode disabled) clusters. Rebooting is not supported on Valkey or Redis
// OSS (cluster mode enabled) clusters.
//
// If you make changes to parameters that require a Valkey or Redis OSS (cluster
// mode enabled) cluster reboot for the changes to be applied, see [Rebooting a Cluster]for an
// alternate process.
//
// [Rebooting a Cluster]: http://docs.aws.amazon.com/AmazonElastiCache/latest/dg/nodes.rebooting.html
