package neptune

// ResetDBClusterParameterGroup is generated as a reference stub.
// Executable command wiring lives under cmd/neptune.go.
//
// Modifies the parameters of a DB cluster parameter group to the default value.
//
// To reset specific parameters submit a list of the following: ParameterName and
// ApplyMethod . To reset the entire DB cluster parameter group, specify the
// DBClusterParameterGroupName and ResetAllParameters parameters.
//
// When resetting the entire group, dynamic parameters are updated immediately and
// static parameters are set to pending-reboot to take effect on the next DB
// instance restart or RebootDBInstancerequest. You must call RebootDBInstance for every DB instance in your DB
// cluster that you want the updated static parameter to apply to.
