package redshift

// RebootCluster is generated as a reference stub.
// Executable command wiring lives under cmd/redshift.go.
//
// Reboots a cluster. This action is taken as soon as possible. It results in a
// momentary outage to the cluster, during which the cluster status is set to
// rebooting . A cluster event is created when the reboot is completed. Any pending
// cluster modifications (see ModifyCluster) are applied at this reboot. For more information
// about managing clusters, go to [Amazon Redshift Clusters]in the Amazon Redshift Cluster Management Guide.
//
// [Amazon Redshift Clusters]: https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html
