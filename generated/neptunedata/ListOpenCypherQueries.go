package neptunedata

// ListOpenCypherQueries is generated as a reference stub.
// Executable command wiring lives under cmd/neptunedata.go.
//
// Lists active openCypher queries. See [Neptune openCypher status endpoint] for more information.
//
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:GetQueryStatus]IAM action in that cluster.
//
// Note that the [neptune-db:QueryLanguage:OpenCypher] IAM condition key can be used in the policy document to restrict
// the use of openCypher queries (see [Condition keys available in Neptune IAM data-access policy statements]).
//
// [Condition keys available in Neptune IAM data-access policy statements]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html
// [Neptune openCypher status endpoint]: https://docs.aws.amazon.com/neptune/latest/userguide/access-graph-opencypher-status.html
// [neptune-db:GetQueryStatus]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#getquerystatus
// [neptune-db:QueryLanguage:OpenCypher]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html#iam-neptune-condition-keys
