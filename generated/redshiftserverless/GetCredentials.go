package redshiftserverless

// GetCredentials is generated as a reference stub.
// Executable command wiring lives under cmd/redshiftserverless.go.
//
// Returns a database user name and temporary password with temporary
// authorization to log in to Amazon Redshift Serverless.
//
// By default, the temporary credentials expire in 900 seconds. You can optionally
// specify a duration between 900 seconds (15 minutes) and 3600 seconds (60
// minutes).
//
// The Identity and Access Management (IAM) user or role that runs GetCredentials
// must have an IAM policy attached that allows access to all necessary actions and
// resources.
//
// If the DbName parameter is specified, the IAM policy must allow access to the
// resource dbname for the specified database name.
