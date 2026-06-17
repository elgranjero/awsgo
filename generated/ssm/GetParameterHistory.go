package ssm

// GetParameterHistory is generated as a reference stub.
// Executable command wiring lives under cmd/ssm.go.
//
// Retrieves the history of all changes to a parameter.
//
// Parameter names can't contain spaces. The service removes any spaces specified
// for the beginning or end of a parameter name. If the specified name for a
// parameter contains spaces between characters, the request fails with a
// ValidationException error.
//
// If you change the KMS key alias for the KMS key used to encrypt a parameter,
// then you must also update the key alias the parameter uses to reference KMS.
// Otherwise, GetParameterHistory retrieves whatever the original key alias was
// referencing.
