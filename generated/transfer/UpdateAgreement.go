package transfer

// UpdateAgreement is generated as a reference stub.
// Executable command wiring lives under cmd/transfer.go.
//
// Updates some of the parameters for an existing agreement. Provide the
// AgreementId and the ServerId for the agreement that you want to update, along
// with the new values for the parameters to update.
//
// Specify either BaseDirectory or CustomDirectories , but not both. Specifying
// both causes the command to fail.
//
// If you update an agreement from using base directory to custom directories, the
// base directory is no longer used. Similarly, if you change from custom
// directories to a base directory, the custom directories are no longer used.
