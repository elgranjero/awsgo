package applicationdiscoveryservice

// BatchDeleteImportData is generated as a reference stub.
// Executable command wiring lives under cmd/applicationdiscoveryservice.go.
//
// Deletes one or more import tasks, each identified by their import ID. Each
// import task has a number of records that can identify servers or applications.
//
// Amazon Web Services Application Discovery Service has built-in matching logic
// that will identify when discovered servers match existing entries that you've
// previously discovered, the information for the already-existing discovered
// server is updated. When you delete an import task that contains records that
// were used to match, the information in those matched records that comes from the
// deleted records will also be deleted.
