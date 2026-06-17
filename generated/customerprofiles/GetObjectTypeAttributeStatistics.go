package customerprofiles

// GetObjectTypeAttributeStatistics is generated as a reference stub.
// Executable command wiring lives under cmd/customerprofiles.go.
//
// The GetObjectTypeAttributeValues API delivers statistical insights about
// attributes within a specific object type, but is exclusively available for
// domains with data store enabled. This API performs daily calculations to provide
// statistical information about your attribute values, helping you understand
// patterns and trends in your data. The statistical calculations are performed
// once per day, providing a consistent snapshot of your attribute data
// characteristics.
//
// You'll receive null values in two scenarios:
//
// During the first period after enabling data vault (unless a calculation cycle
// occurs, which happens once daily).
//
// For attributes that don't contain numeric values.
