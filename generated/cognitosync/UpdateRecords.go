package cognitosync

// UpdateRecords is generated as a reference stub.
// Executable command wiring lives under cmd/cognitosync.go.
//
// Posts updates to records and adds and deletes records for a dataset and user.
//
// The sync count in the record patch is your last known sync count for that
// record. The server will reject an UpdateRecords request with a
// ResourceConflictException if you try to patch a record with a new value but a
// stale sync count.
//
// For example, if the sync count on the server is 5 for a key called highScore
// and you try and submit a new highScore with sync count of 4, the request will be
// rejected. To obtain the current sync count for a record, call ListRecords. On a
// successful update of the record, the response returns the new sync count for
// that record. You should present that sync count the next time you try to update
// that same record. When the record does not exist, specify the sync count as 0.
//
// This API can be called with temporary user credentials provided by Cognito
// Identity or with developer credentials.
