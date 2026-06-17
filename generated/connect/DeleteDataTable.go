package connect

// DeleteDataTable is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// Deletes a data table and all associated attributes, versions, audits, and
// values. Does not update any references to the data table, even from other data
// tables. This includes dynamic values and conditional validations. System managed
// data tables are not deletable by customers. API users may delete the table at
// any time. When deletion is requested from the admin website, a warning is shown
// alerting the user of the most recent time the table and its values were
// accessed.
