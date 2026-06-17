package neptunedata

// ExecuteGremlinExplainQuery is generated as a reference stub.
// Executable command wiring lives under cmd/neptunedata.go.
//
// Executes a Gremlin Explain query.
//
// Amazon Neptune has added a Gremlin feature named explain that provides is a
// self-service tool for understanding the execution approach being taken by the
// Neptune engine for the query. You invoke it by adding an explain parameter to
// an HTTP call that submits a Gremlin query.
//
// The explain feature provides information about the logical structure of query
// execution plans. You can use this information to identify potential evaluation
// and execution bottlenecks and to tune your query, as explained in [Tuning Gremlin queries]. You can
// also use query hints to improve query execution plans.
//
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows one of the following IAM actions in that cluster, depending on the
// query:
//
// [neptune-db:ReadDataViaQuery]
//
// [neptune-db:WriteDataViaQuery]
//
// [neptune-db:DeleteDataViaQuery]
//
// Note that the [neptune-db:QueryLanguage:Gremlin] IAM condition key can be used in the policy document to restrict
// the use of Gremlin queries (see [Condition keys available in Neptune IAM data-access policy statements]).
//
// [neptune-db:DeleteDataViaQuery]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#deletedataviaquery
// [Condition keys available in Neptune IAM data-access policy statements]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html
// [neptune-db:ReadDataViaQuery]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#readdataviaquery
// [Tuning Gremlin queries]: https://docs.aws.amazon.com/neptune/latest/userguide/gremlin-traversal-tuning.html
// [neptune-db:WriteDataViaQuery]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#writedataviaquery
// [neptune-db:QueryLanguage:Gremlin]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html#iam-neptune-condition-keys
