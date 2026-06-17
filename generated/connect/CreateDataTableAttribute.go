package connect

// CreateDataTableAttribute is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// Adds an attribute to an existing data table. Creating a new primary attribute
// uses the empty value for the specified value type for all existing records. This
// should not affect uniqueness of published data tables since the existing primary
// values will already be unique. Creating attributes does not create any values.
// System managed tables may not allow customers to create new attributes.
