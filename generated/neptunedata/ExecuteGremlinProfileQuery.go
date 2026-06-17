package neptunedata

// ExecuteGremlinProfileQuery is generated as a reference stub.
// Executable command wiring lives under cmd/neptunedata.go.
//
// Executes a Gremlin Profile query, which runs a specified traversal, collects
// various metrics about the run, and produces a profile report as output. See [Gremlin profile API in Neptune]for
// details.
//
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:ReadDataViaQuery]IAM action in that cluster.
//
// Note that the [neptune-db:QueryLanguage:Gremlin] IAM condition key can be used in the policy document to restrict
// the use of Gremlin queries (see [Condition keys available in Neptune IAM data-access policy statements]).
//
// [Gremlin profile API in Neptune]: https://docs.aws.amazon.com/neptune/latest/userguide/gremlin-profile-api.html
// [Condition keys available in Neptune IAM data-access policy statements]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html
// [neptune-db:ReadDataViaQuery]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#readdataviaquery
// [neptune-db:QueryLanguage:Gremlin]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html#iam-neptune-condition-keys
