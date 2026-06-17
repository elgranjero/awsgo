package mailmanager

// DeleteArchive is generated as a reference stub.
// Executable command wiring lives under cmd/mailmanager.go.
//
// Initiates deletion of an email archive. This changes the archive state to
// pending deletion. In this state, no new emails can be added, and existing
// archived emails become inaccessible (search, export, download). The archive and
// all of its contents will be permanently deleted 30 days after entering the
// pending deletion state, regardless of the configured retention period.
