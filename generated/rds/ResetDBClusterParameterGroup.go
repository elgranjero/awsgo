package rds

// ResetDBClusterParameterGroup is generated as a reference stub.
// Executable command wiring lives under cmd/rds.go.
//
// Modifies the parameters of a DB cluster parameter group to the default value.
// To reset specific parameters submit a list of the following: ParameterName and
// ApplyMethod . To reset the entire DB cluster parameter group, specify the
// DBClusterParameterGroupName and ResetAllParameters parameters.
//
// When resetting the entire group, dynamic parameters are updated immediately and
// static parameters are set to pending-reboot to take effect on the next DB
// instance restart or RebootDBInstance request. You must call RebootDBInstance
// for every DB instance in your DB cluster that you want the updated static
// parameter to apply to.
//
// For more information on Amazon Aurora DB clusters, see [What is Amazon Aurora?] in the Amazon Aurora
// User Guide.
//
// For more information on Multi-AZ DB clusters, see [Multi-AZ DB cluster deployments] in the Amazon RDS User
// Guide.
//
// [What is Amazon Aurora?]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html
// [Multi-AZ DB cluster deployments]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html
