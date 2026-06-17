package connectcases

// DeleteCase is generated as a reference stub.
// Executable command wiring lives under cmd/connectcases.go.
//
// The DeleteCase API permanently deletes a case and all its associated resources
//
// from the cases data store. After a successful deletion, you cannot:
//
// - Retrieve related items
//
// - Access audit history
//
// - Perform any operations that require the CaseID
//
// This action is irreversible. After you delete a case, you cannot recover its
// data.
