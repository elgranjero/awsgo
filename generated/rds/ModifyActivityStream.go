package rds

// ModifyActivityStream is generated as a reference stub.
// Executable command wiring lives under cmd/rds.go.
//
// Changes the audit policy state of a database activity stream to either locked
// (default) or unlocked. A locked policy is read-only, whereas an unlocked policy
// is read/write. If your activity stream is started and locked, you can unlock it,
// customize your audit policy, and then lock your activity stream. Restarting the
// activity stream isn't required. For more information, see [Modifying a database activity stream]in the Amazon RDS
// User Guide.
//
// This operation is supported for RDS for Oracle and Microsoft SQL Server.
//
// [Modifying a database activity stream]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/DBActivityStreams.Modifying.html
