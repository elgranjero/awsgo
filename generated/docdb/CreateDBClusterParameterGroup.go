package docdb

// CreateDBClusterParameterGroup is generated as a reference stub.
// Executable command wiring lives under cmd/docdb.go.
//
// Creates a new cluster parameter group.
//
// Parameters in a cluster parameter group apply to all of the instances in a
// cluster.
//
// A cluster parameter group is initially created with the default parameters for
// the database engine used by instances in the cluster. In Amazon DocumentDB, you
// cannot make modifications directly to the default.docdb3.6 cluster parameter
// group. If your Amazon DocumentDB cluster is using the default cluster parameter
// group and you want to modify a value in it, you must first [create a new parameter group]or [copy an existing parameter group], modify it, and
// then apply the modified parameter group to your cluster. For the new cluster
// parameter group and associated settings to take effect, you must then reboot the
// instances in the cluster without failover. For more information, see [Modifying Amazon DocumentDB Cluster Parameter Groups].
//
// [create a new parameter group]: https://docs.aws.amazon.com/documentdb/latest/developerguide/cluster_parameter_group-create.html
// [Modifying Amazon DocumentDB Cluster Parameter Groups]: https://docs.aws.amazon.com/documentdb/latest/developerguide/cluster_parameter_group-modify.html
// [copy an existing parameter group]: https://docs.aws.amazon.com/documentdb/latest/developerguide/cluster_parameter_group-copy.html
