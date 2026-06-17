package dynamodb

// CreateTable is generated as a reference stub.
// Executable command wiring lives under cmd/dynamodb.go.
//
// The CreateTable operation adds a new table to your account. In an Amazon Web
// Services account, table names must be unique within each Region. That is, you
// can have two tables with same name if you create the tables in different
// Regions.
//
// CreateTable is an asynchronous operation. Upon receiving a CreateTable request,
// DynamoDB immediately returns a response with a TableStatus of CREATING . After
// the table is created, DynamoDB sets the TableStatus to ACTIVE . You can perform
// read and write operations only on an ACTIVE table.
//
// You can optionally define secondary indexes on the new table, as part of the
// CreateTable operation. If you want to create multiple tables with secondary
// indexes on them, you must create the tables sequentially. Only one table with
// secondary indexes can be in the CREATING state at any given time.
//
// You can use the DescribeTable action to check the table status.
