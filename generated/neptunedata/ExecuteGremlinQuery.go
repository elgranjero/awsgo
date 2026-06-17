package neptunedata

// ExecuteGremlinQuery is generated as a reference stub.
// Executable command wiring lives under cmd/neptunedata.go.
//
// This commands executes a Gremlin query. Amazon Neptune is compatible with
// Apache TinkerPop3 and Gremlin, so you can use the Gremlin traversal language to
// query the graph, as described under [The Graph]in the Apache TinkerPop3 documentation.
// More details can also be found in [Accessing a Neptune graph with Gremlin].
//
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that enables one of the following IAM actions in that cluster, depending on the
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
// [The Graph]: https://tinkerpop.apache.org/docs/current/reference/#graph
// [Accessing a Neptune graph with Gremlin]: https://docs.aws.amazon.com/neptune/latest/userguide/access-graph-gremlin.html
// [neptune-db:ReadDataViaQuery]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#readdataviaquery
// [neptune-db:WriteDataViaQuery]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#writedataviaquery
// [neptune-db:QueryLanguage:Gremlin]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html#iam-neptune-condition-keys
