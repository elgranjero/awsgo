package rdsdata

// BeginTransaction is generated as a reference stub.
// Executable command wiring lives under cmd/rdsdata.go.
//
// Starts a SQL transaction.
//
// A transaction can run for a maximum of 24 hours. A transaction is terminated
// and rolled back automatically after 24 hours.
//
// A transaction times out if no calls use its transaction ID in three minutes. If
// a transaction times out before it's committed, it's rolled back automatically.
//
// For Aurora MySQL, DDL statements inside a transaction cause an implicit commit.
// We recommend that you run each MySQL DDL statement in a separate
// ExecuteStatement call with continueAfterTimeout enabled.
