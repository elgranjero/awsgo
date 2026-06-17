package rds

// ModifyDBClusterParameterGroup is generated as a reference stub.
// Executable command wiring lives under cmd/rds.go.
//
// Modifies the parameters of a DB cluster parameter group. To modify more than
// one parameter, submit a list of the following: ParameterName , ParameterValue ,
// and ApplyMethod . A maximum of 20 parameters can be modified in a single request.
//
// There are two types of parameters - dynamic parameters and static parameters.
// Changes to dynamic parameters are applied to the DB cluster immediately without
// a reboot. Changes to static parameters are applied only after the DB cluster is
// rebooted, which can be done using RebootDBCluster operation. You can use the
// Parameter Groups option of the [Amazon RDS console]or the DescribeDBClusterParameters operation to
// verify that your DB cluster parameter group has been created or modified.
//
// For more information on Amazon Aurora DB clusters, see [What is Amazon Aurora?] in the Amazon Aurora
// User Guide.
//
// For more information on Multi-AZ DB clusters, see [Multi-AZ DB cluster deployments] in the Amazon RDS User
// Guide.
//
// [Amazon RDS console]: https://console.aws.amazon.com/rds/
// [What is Amazon Aurora?]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html
// [Multi-AZ DB cluster deployments]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html
