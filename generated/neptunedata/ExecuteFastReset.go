package neptunedata

// ExecuteFastReset is generated as a reference stub.
// Executable command wiring lives under cmd/neptunedata.go.
//
// The fast reset REST API lets you reset a Neptune graph quicky and easily,
// removing all of its data.
//
// Neptune fast reset is a two-step process. First you call ExecuteFastReset with
// action set to initiateDatabaseReset . This returns a UUID token which you then
// include when calling ExecuteFastReset again with action set to
// performDatabaseReset . See [Empty an Amazon Neptune DB cluster using the fast reset API].
//
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:ResetDatabase]IAM action in that cluster.
//
// [Empty an Amazon Neptune DB cluster using the fast reset API]: https://docs.aws.amazon.com/neptune/latest/userguide/manage-console-fast-reset.html
// [neptune-db:ResetDatabase]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#resetdatabase
