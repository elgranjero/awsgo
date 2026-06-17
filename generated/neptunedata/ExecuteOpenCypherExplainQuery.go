package neptunedata

// ExecuteOpenCypherExplainQuery is generated as a reference stub.
// Executable command wiring lives under cmd/neptunedata.go.
//
// Executes an openCypher explain request. See [The openCypher explain feature] for more information.
//
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:ReadDataViaQuery]IAM action in that cluster.
//
// Note that the [neptune-db:QueryLanguage:OpenCypher] IAM condition key can be used in the policy document to restrict
// the use of openCypher queries (see [Condition keys available in Neptune IAM data-access policy statements]).
//
// [Condition keys available in Neptune IAM data-access policy statements]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html
// [The openCypher explain feature]: https://docs.aws.amazon.com/neptune/latest/userguide/access-graph-opencypher-explain.html
// [neptune-db:ReadDataViaQuery]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#readdataviaquery
// [neptune-db:QueryLanguage:OpenCypher]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html#iam-neptune-condition-keys
