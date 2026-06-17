package personalize

// CreateSchema is generated as a reference stub.
// Executable command wiring lives under cmd/personalize.go.
//
// Creates an Amazon Personalize schema from the specified schema string. The
// schema you create must be in Avro JSON format.
//
// Amazon Personalize recognizes three schema variants. Each schema is associated
// with a dataset type and has a set of required field and keywords. If you are
// creating a schema for a dataset in a Domain dataset group, you provide the
// domain of the Domain dataset group. You specify a schema when you call [CreateDataset].
//
// # Related APIs
//
// [ListSchemas]
//
// [DescribeSchema]
//
// [DeleteSchema]
//
// [CreateDataset]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateDataset.html
// [DescribeSchema]: https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeSchema.html
// [DeleteSchema]: https://docs.aws.amazon.com/personalize/latest/dg/API_DeleteSchema.html
// [ListSchemas]: https://docs.aws.amazon.com/personalize/latest/dg/API_ListSchemas.html
