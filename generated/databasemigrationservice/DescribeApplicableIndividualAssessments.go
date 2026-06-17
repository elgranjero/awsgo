package databasemigrationservice

// DescribeApplicableIndividualAssessments is generated as a reference stub.
// Executable command wiring lives under cmd/databasemigrationservice.go.
//
// Provides a list of individual assessments that you can specify for a new
// premigration assessment run, given one or more parameters.
//
// If you specify an existing migration task, this operation provides the default
// individual assessments you can specify for that task. Otherwise, the specified
// parameters model elements of a possible migration task on which to base a
// premigration assessment run.
//
// To use these migration task modeling parameters, you must specify an existing
// replication instance, a source database engine, a target database engine, and a
// migration type. This combination of parameters potentially limits the default
// individual assessments available for an assessment run created for a
// corresponding migration task.
//
// If you specify no parameters, this operation provides a list of all possible
// individual assessments that you can specify for an assessment run. If you
// specify any one of the task modeling parameters, you must specify all of them or
// the operation cannot provide a list of individual assessments. The only
// parameter that you can specify alone is for an existing migration task. The
// specified task definition then determines the default list of individual
// assessments that you can specify in an assessment run for the task.
