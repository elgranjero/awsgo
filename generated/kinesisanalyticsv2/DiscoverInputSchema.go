package kinesisanalyticsv2

// DiscoverInputSchema is generated as a reference stub.
// Executable command wiring lives under cmd/kinesisanalyticsv2.go.
//
// Infers a schema for a SQL-based Kinesis Data Analytics application by
// evaluating sample records on the specified streaming source (Kinesis data stream
// or Kinesis Data Firehose delivery stream) or Amazon S3 object. In the response,
// the operation returns the inferred schema and also the sample records that the
// operation used to infer the schema.
//
// You can use the inferred schema when configuring a streaming source for your
// application. When you create an application using the Kinesis Data Analytics
// console, the console uses this operation to infer a schema and show it in the
// console user interface.
