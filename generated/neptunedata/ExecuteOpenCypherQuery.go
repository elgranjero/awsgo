package neptunedata

// ExecuteOpenCypherQuery is generated as a reference stub.
// Executable command wiring lives under cmd/neptunedata.go.
//
// Executes an openCypher query. See [Accessing the Neptune Graph with openCypher] for more information.
//
// Neptune supports building graph applications using openCypher, which is
// currently one of the most popular query languages among developers working with
// graph databases. Developers, business analysts, and data scientists like
// openCypher's declarative, SQL-inspired syntax because it provides a familiar
// structure in which to querying property graphs.
//
// The openCypher language was originally developed by Neo4j, then open-sourced in
// 2015 and contributed to the [openCypher project]under an Apache 2 open-source license.
//
// Note that when invoking this operation in a Neptune cluster that has IAM
// authentication enabled, the IAM user or role making the request must have a
// policy attached that allows one of the following IAM actions in that cluster,
// depending on the query:
//
// [neptune-db:ReadDataViaQuery]
//
// [neptune-db:WriteDataViaQuery]
//
// [neptune-db:DeleteDataViaQuery]
//
// Note also that the [neptune-db:QueryLanguage:OpenCypher] IAM condition key can be used in the policy document to
// restrict the use of openCypher queries (see [Condition keys available in Neptune IAM data-access policy statements]).
//
// [neptune-db:DeleteDataViaQuery]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#deletedataviaquery
// [Accessing the Neptune Graph with openCypher]: https://docs.aws.amazon.com/neptune/latest/userguide/access-graph-opencypher.html
// [Condition keys available in Neptune IAM data-access policy statements]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html
// [openCypher project]: https://opencypher.org/
// [neptune-db:ReadDataViaQuery]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#readdataviaquery
// [neptune-db:WriteDataViaQuery]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#writedataviaquery
// [neptune-db:QueryLanguage:OpenCypher]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html#iam-neptune-condition-keys
