package connect

// BatchCreateDataTableValue is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// Creates values for attributes in a data table. The value may be a default or it
// may be associated with a primary value. The value must pass all customer defined
// validation as well as the default validation for the value type. The operation
// must conform to Batch Operation API Standards. Although the standard specifies
// that successful and failed entities are listed separately in the response,
// authorization fails if any primary values or attributes are unauthorized. The
// combination of primary values and the attribute name serve as the identifier for
// the individual item request.
