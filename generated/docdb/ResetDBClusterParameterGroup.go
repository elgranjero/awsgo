package docdb

// ResetDBClusterParameterGroup is generated as a reference stub.
// Executable command wiring lives under cmd/docdb.go.
//
// Modifies the parameters of a cluster parameter group to the default value. To
//
// reset specific parameters, submit a list of the following: ParameterName and
// ApplyMethod . To reset the entire cluster parameter group, specify the
// DBClusterParameterGroupName and ResetAllParameters parameters.
//
// When you reset the entire group, dynamic parameters are updated immediately and
// static parameters are set to pending-reboot to take effect on the next DB
// instance reboot.
