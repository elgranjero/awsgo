package emr

// DescribeJobFlows is generated as a reference stub.
// Executable command wiring lives under cmd/emr.go.
//
// This API is no longer supported and will eventually be removed. We recommend
// you use ListClusters, DescribeCluster, ListSteps, ListInstanceGroups and ListBootstrapActions instead.
//
// DescribeJobFlows returns a list of job flows that match all of the supplied
// parameters. The parameters can include a list of job flow IDs, job flow states,
// and restrictions on job flow creation date and time.
//
// Regardless of supplied parameters, only job flows created within the last two
// months are returned.
//
// If no parameters are supplied, then job flows matching either of the following
// criteria are returned:
//
// - Job flows created and completed in the last two weeks
//
// - Job flows created within the last two months that are in one of the
// following states: RUNNING , WAITING , SHUTTING_DOWN , STARTING
//
// Amazon EMR can return a maximum of 512 job flow descriptions.
//
// Deprecated: This operation has been deprecated.
