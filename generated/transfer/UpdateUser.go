package transfer

// UpdateUser is generated as a reference stub.
// Executable command wiring lives under cmd/transfer.go.
//
// Assigns new properties to a user. Parameters you pass modify any or all of the
// following: the home directory, role, and policy for the UserName and ServerId
// you specify.
//
// The response returns the ServerId and the UserName for the updated user.
//
// In the console, you can select Restricted when you create or update a user.
// This ensures that the user can't access anything outside of their home
// directory. The programmatic way to configure this behavior is to update the
// user. Set their HomeDirectoryType to LOGICAL , and specify HomeDirectoryMappings
// with Entry as root ( / ) and Target as their home directory.
//
// For example, if the user's home directory is /test/admin-user , the following
// command updates the user so that their configuration in the console shows the
// Restricted flag as selected.
//
// aws transfer update-user --server-id <server-id> --user-name admin-user
// --home-directory-type LOGICAL --home-directory-mappings "[{\"Entry\":\"/\",
// \"Target\":\"/test/admin-user\"}]"
