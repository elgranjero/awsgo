package backupsearch

// ListSearchJobBackups is generated as a reference stub.
// Executable command wiring lives under cmd/backupsearch.go.
//
// This operation returns a list of all backups (recovery points) in a paginated
// format that were included in the search job.
//
// If a search does not display an expected backup in the results, you can call
// this operation to display each backup included in the search. Any backups that
// were not included because they have a FAILED status from a permissions issue
// will be displayed, along with a status message.
//
// Only recovery points with a backup index that has a status of ACTIVE will be
// included in search results. If the index has any other status, its status will
// be displayed along with a status message.
