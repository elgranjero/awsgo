package rdsdata

// ExecuteStatement is generated as a reference stub.
// Executable command wiring lives under cmd/rdsdata.go.
//
// Runs a SQL statement against a database.
//
// If a call isn't part of a transaction because it doesn't include the
// transactionID parameter, changes that result from the call are committed
// automatically.
//
// If the binary response data from the database is more than 1 MB, the call is
// terminated.
