package docdb

// ModifyDBClusterParameterGroup is generated as a reference stub.
// Executable command wiring lives under cmd/docdb.go.
//
// Modifies the parameters of a cluster parameter group. To modify more than one
//
// parameter, submit a list of the following: ParameterName , ParameterValue , and
// ApplyMethod . A maximum of 20 parameters can be modified in a single request.
//
// Changes to dynamic parameters are applied immediately. Changes to static
// parameters require a reboot or maintenance window
//
// before the change can take effect.
//
// After you create a cluster parameter group, you should wait at least 5 minutes
// before creating your first cluster that uses that cluster parameter group as the
// default parameter group. This allows Amazon DocumentDB to fully complete the
// create action before the parameter group is used as the default for a new
// cluster. This step is especially important for parameters that are critical when
// creating the default database for a cluster, such as the character set for the
// default database defined by the character_set_database parameter.
